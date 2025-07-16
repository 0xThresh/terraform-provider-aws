//Copyright © 2025, Oracle and/or its affiliates. All rights reserved.

package odb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/odb"
	odbtypes "github.com/aws/aws-sdk-go-v2/service/odb/types"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
	fwtypes "github.com/hashicorp/terraform-provider-aws/internal/framework/types"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/names"
	"time"
)

// Function annotations are used for datasource registration to the Provider. DO NOT EDIT.
// @FrameworkDataSource("aws_odb_network", name="Network")
func newDataSourceNetwork(context.Context) (datasource.DataSourceWithConfigure, error) {
	return &dataSourceNetwork{}, nil
}

const (
	DSNameNetwork = "Odb Network Data Source"
)

type dataSourceNetwork struct {
	framework.DataSourceWithModel[odbNetworkDataSourceModel]
}

var OdbNetworkDataSource dataSourceNetwork

func (d *dataSourceNetwork) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	statusType := fwtypes.StringEnumType[odbtypes.ResourceStatus]()
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			names.AttrARN: framework.ARNAttributeComputedOnly(),
			names.AttrID: schema.StringAttribute{
				Required: true,
			},
			"display_name": schema.StringAttribute{
				Computed: true,
			},
			"availability_zone_id": schema.StringAttribute{
				Computed: true,
			},
			"availability_zone": schema.StringAttribute{
				Computed: true,
			},
			"backup_subnet_cidr": schema.StringAttribute{
				Computed: true,
			},
			"client_subnet_cidr": schema.StringAttribute{
				Computed: true,
			},
			"custom_domain_name": schema.StringAttribute{
				Computed: true,
			},
			"default_dns_prefix": schema.StringAttribute{
				Computed: true,
			},
			"oci_network_anchor_id": schema.StringAttribute{
				Computed: true,
			},
			"oci_network_anchor_url": schema.StringAttribute{
				Computed: true,
			},
			"oci_resource_anchor_name": schema.StringAttribute{
				Computed: true,
			},
			"oci_vcn_id": schema.StringAttribute{
				Computed: true,
			},
			"oci_vcn_url": schema.StringAttribute{
				Computed: true,
			},
			"percent_progress": schema.Float64Attribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				CustomType: statusType,
				Computed:   true,
			},
			"status_reason": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"managed_services": schema.ObjectAttribute{
				Computed:   true,
				CustomType: fwtypes.NewObjectTypeOf[odbNetworkManagedServicesDataSourceModel](ctx),
				AttributeTypes: map[string]attr.Type{
					"service_network_arn":  types.StringType,
					"resource_gateway_arn": types.StringType,
					"managed_service_ipv4_cidrs": types.ListType{
						ElemType: types.StringType,
					},
					"service_network_endpoint": types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"vpc_endpoint_id":   types.StringType,
							"vpc_endpoint_type": fwtypes.StringEnumType[odbtypes.VpcEndpointType](),
						},
					},
					"managed_s3_backup_access": types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"status": fwtypes.StringEnumType[odbtypes.ResourceStatus](),
							"ipv4_addresses": types.ListType{
								ElemType: types.StringType,
							},
						},
					},
					"zero_tl_access": types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"status": fwtypes.StringEnumType[odbtypes.ManagedResourceStatus](),
							"cidr":   types.StringType,
						},
					},
					"s3_access": types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"status": fwtypes.StringEnumType[odbtypes.ManagedResourceStatus](),
							"ipv4_addresses": types.ListType{
								ElemType: types.StringType,
							},
							"domain_name":        types.StringType,
							"s3_policy_document": types.StringType,
						},
					},
				},
			},
			names.AttrTags: tftags.TagsAttributeComputedOnly(),
			"oci_dns_forwarding_configs": schema.ListAttribute{
				Computed:   true,
				CustomType: fwtypes.NewListNestedObjectTypeOf[odbNwkOciDnsForwardingConfigDataSourceModel](ctx),
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"domain_name":         types.StringType,
						"oci_dns_listener_ip": types.StringType,
					},
				},
			},
		},
	}
}

func (d *dataSourceNetwork) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	conn := d.Meta().ODBClient(ctx)
	var data odbNetworkDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := odb.GetOdbNetworkInput{
		OdbNetworkId: data.OdbNetworkId.ValueStringPointer(),
	}

	out, err := conn.GetOdbNetwork(ctx, &input)
	if err != nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.ODB, create.ErrActionReading, DSNameNetwork, data.OdbNetworkId.String(), err),
			err.Error(),
		)
		return
	}

	data.CreatedAt = types.StringValue(out.OdbNetwork.CreatedAt.Format(time.RFC3339))
	resp.Diagnostics.Append(flex.Flatten(ctx, out.OdbNetwork, &data, flex.WithIgnoredFieldNamesAppend("CreatedAt"))...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

type odbNetworkDataSourceModel struct {
	framework.WithRegionModel
	AvailabilityZone        types.String                                                                 `tfsdk:"availability_zone"`
	AvailabilityZoneId      types.String                                                                 `tfsdk:"availability_zone_id"`
	BackupSubnetCidr        types.String                                                                 `tfsdk:"backup_subnet_cidr"`
	ClientSubnetCidr        types.String                                                                 `tfsdk:"client_subnet_cidr"`
	CustomDomainName        types.String                                                                 `tfsdk:"custom_domain_name"`
	DefaultDnsPrefix        types.String                                                                 `tfsdk:"default_dns_prefix"`
	DisplayName             types.String                                                                 `tfsdk:"display_name"`
	OciDnsForwardingConfigs fwtypes.ListNestedObjectValueOf[odbNwkOciDnsForwardingConfigDataSourceModel] `tfsdk:"oci_dns_forwarding_configs"`
	OciNetworkAnchorId      types.String                                                                 `tfsdk:"oci_network_anchor_id"`
	OciNetworkAnchorUrl     types.String                                                                 `tfsdk:"oci_network_anchor_url"`
	OciResourceAnchorName   types.String                                                                 `tfsdk:"oci_resource_anchor_name"`
	OciVcnId                types.String                                                                 `tfsdk:"oci_vcn_id"`
	OciVcnUrl               types.String                                                                 `tfsdk:"oci_vcn_url"`
	OdbNetworkArn           types.String                                                                 `tfsdk:"arn"`
	OdbNetworkId            types.String                                                                 `tfsdk:"id"`
	PercentProgress         types.Float64                                                                `tfsdk:"percent_progress"`
	Status                  fwtypes.StringEnum[odbtypes.ResourceStatus]                                  `tfsdk:"status"`
	StatusReason            types.String                                                                 `tfsdk:"status_reason"`
	CreatedAt               types.String                                                                 `tfsdk:"created_at"`
	ManagedServices         fwtypes.ObjectValueOf[odbNetworkManagedServicesDataSourceModel]              `tfsdk:"managed_services"`
	Tags                    tftags.Map                                                                   `tfsdk:"tags"`
}

type odbNwkOciDnsForwardingConfigDataSourceModel struct {
	DomainName       types.String `tfsdk:"domain_name"`
	OciDnsListenerIp types.String `tfsdk:"oci_dns_listener_ip"`
}

type odbNetworkManagedServicesDataSourceModel struct {
	ServiceNetworkArn        types.String                                                           `tfsdk:"service_network_arn"`
	ResourceGatewayArn       types.String                                                           `tfsdk:"resource_gateway_arn"`
	ManagedServicesIpv4Cidrs fwtypes.ListOfString                                                   `tfsdk:"managed_service_ipv4_cidrs"`
	ServiceNetworkEndpoint   fwtypes.ObjectValueOf[serviceNetworkEndpointOdbNetworkDataSourceModel] `tfsdk:"service_network_endpoint"`
	ManagedS3BackupAccess    fwtypes.ObjectValueOf[managedS3BackupAccessOdbNetworkDataSourceModel]  `tfsdk:"managed_s3_backup_access"`
	ZeroEtlAccess            fwtypes.ObjectValueOf[zeroEtlAccessOdbNetworkDataSourceModel]          `tfsdk:"zero_tl_access"`
	S3Access                 fwtypes.ObjectValueOf[s3AccessOdbNetworkDataSourceModel]               `tfsdk:"s3_access"`
}

type serviceNetworkEndpointOdbNetworkDataSourceModel struct {
	VpcEndpointId   types.String                                 `tfsdk:"vpc_endpoint_id"`
	VpcEndpointType fwtypes.StringEnum[odbtypes.VpcEndpointType] `tfsdk:"vpc_endpoint_type"`
}

type managedS3BackupAccessOdbNetworkDataSourceModel struct {
	Status        fwtypes.StringEnum[odbtypes.ManagedResourceStatus] `tfsdk:"status"`
	Ipv4Addresses fwtypes.ListOfString                               `tfsdk:"ipv4_addresses"`
}

type zeroEtlAccessOdbNetworkDataSourceModel struct {
	Status fwtypes.StringEnum[odbtypes.ManagedResourceStatus] `tfsdk:"status"`
	Cidr   types.String                                       `tfsdk:"cidr"`
}

type s3AccessOdbNetworkDataSourceModel struct {
	Status           fwtypes.StringEnum[odbtypes.ManagedResourceStatus] `tfsdk:"status"`
	Ipv4Addresses    fwtypes.ListOfString                               `tfsdk:"ipv4_addresses"`
	DomainName       types.String                                       `tfsdk:"domain_name"`
	S3PolicyDocument types.String                                       `tfsdk:"s3_policy_document"`
}
