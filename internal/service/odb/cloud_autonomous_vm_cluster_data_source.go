// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package odb

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/odb"
	odbtypes "github.com/aws/aws-sdk-go-v2/service/odb/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
	fwtypes "github.com/hashicorp/terraform-provider-aws/internal/framework/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// Function annotations are used for datasource registration to the Provider. DO NOT EDIT.
// @FrameworkDataSource("aws_odb_cloud_autonomous_vm_cluster", name="Cloud Autonomous Vm Cluster")
func newDataSourceCloudAutonomousVmCluster(context.Context) (datasource.DataSourceWithConfigure, error) {
	return &dataSourceCloudAutonomousVmCluster{}, nil
}

const (
	DSNameCloudAutonomousVmCluster = "Cloud Autonomous Vm Cluster Data Source"
)

type dataSourceCloudAutonomousVmCluster struct {
	framework.DataSourceWithModel[cloudAutonomousVmClusterDataSourceModel]
}

func (d *dataSourceCloudAutonomousVmCluster) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	status := fwtypes.StringEnumType[odbtypes.ResourceStatus]()
	licenseModel := fwtypes.StringEnumType[odbtypes.LicenseModel]()
	computeModel := fwtypes.StringEnumType[odbtypes.ComputeModel]()
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			names.AttrARN: framework.ARNAttributeComputedOnly(),

			names.AttrID: schema.StringAttribute{
				Required: true,
			},
			"cloud_exadata_infrastructure_id": schema.StringAttribute{
				Computed: true,
			},
			"autonomous_data_storage_percentage": schema.Float32Attribute{
				Computed: true,
			},
			"autonomous_data_storage_size_in_tbs": schema.Float64Attribute{
				Computed: true,
			},
			"available_autonomous_data_storage_size_in_tbs": schema.Float64Attribute{
				Computed: true,
			},
			"available_container_databases": schema.Int32Attribute{
				Computed: true,
			},
			"available_cpus": schema.Float32Attribute{
				Computed: true,
			},
			"compute_model": schema.StringAttribute{
				CustomType: computeModel,
				Computed:   true,
			},
			"cpu_core_count": schema.Int32Attribute{
				Computed: true,
			},
			"cpu_core_count_per_node": schema.Int32Attribute{
				Computed: true,
			},
			"cpu_percentage": schema.Float32Attribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"data_storage_size_in_gbs": schema.Float64Attribute{
				Computed: true,
			},
			"data_storage_size_in_tbs": schema.Float64Attribute{
				Computed: true,
			},
			"odb_node_storage_size_in_gbs": schema.Int32Attribute{
				Computed: true,
			},
			"db_servers": schema.SetAttribute{
				Computed:    true,
				CustomType:  fwtypes.SetOfStringType,
				ElementType: types.StringType,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"display_name": schema.StringAttribute{
				Computed: true,
			},
			"domain": schema.StringAttribute{
				Computed: true,
			},
			"exadata_storage_in_tbs_lowest_scaled_value": schema.Float64Attribute{
				Computed: true,
			},
			"hostname": schema.StringAttribute{
				Computed: true,
			},
			"is_mtls_enabled_vm_cluster": schema.BoolAttribute{
				Computed: true,
			},
			"license_model": schema.StringAttribute{
				CustomType: licenseModel,
				Computed:   true,
			},
			"max_acds_lowest_scaled_value": schema.Int32Attribute{
				Computed: true,
			},
			"memory_per_oracle_compute_unit_in_gbs": schema.Int32Attribute{
				Computed: true,
			},
			"memory_size_in_gbs": schema.Int32Attribute{
				Computed: true,
			},
			"node_count": schema.Int32Attribute{
				Computed: true,
			},
			"non_provisionable_autonomous_container_databases": schema.Int32Attribute{
				Computed: true,
			},
			"oci_resource_anchor_name": schema.StringAttribute{
				Computed: true,
			},
			"oci_url": schema.StringAttribute{
				Computed: true,
			},
			"ocid": schema.StringAttribute{
				Computed: true,
			},
			"odb_network_id": schema.StringAttribute{
				Computed: true,
			},
			"percent_progress": schema.Float32Attribute{
				Computed: true,
			},
			"provisionable_autonomous_container_databases": schema.Int32Attribute{
				Computed: true,
			},
			"provisioned_autonomous_container_databases": schema.Int32Attribute{
				Computed: true,
			},
			"provisioned_cpus": schema.Float32Attribute{
				Computed: true,
			},
			"reclaimable_cpus": schema.Float32Attribute{
				Computed: true,
			},
			"reserved_cpus": schema.Float32Attribute{
				Computed: true,
			},
			"scan_listener_port_non_tls": schema.Int32Attribute{
				Computed: true,
			},
			"scan_listener_port_tls": schema.Int32Attribute{
				Computed: true,
			},
			"shape": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				CustomType: status,
				Computed:   true,
			},
			"reason": schema.StringAttribute{
				Computed: true,
			},
			"time_database_ssl_certificate_expires": schema.StringAttribute{
				Computed: true,
			},
			"time_ords_certificate_expires": schema.StringAttribute{
				Computed: true,
			},
			"time_zone": schema.StringAttribute{
				Computed: true,
			},
			"total_autonomous_data_storage_in_tbs": schema.Float32Attribute{
				Computed: true,
			},
			"total_container_databases": schema.Int32Attribute{
				Computed: true,
			},
			"total_cpus": schema.Float32Attribute{
				Computed: true,
			},
			names.AttrTags: tftags.TagsAttributeComputedOnly(),
			"maintenance_window": schema.ObjectAttribute{
				Computed:   true,
				CustomType: fwtypes.NewObjectTypeOf[cloudAutonomousVmClusterMaintenanceWindowDataSourceModel](ctx),

				AttributeTypes: map[string]attr.Type{
					"days_of_week": types.SetType{
						ElemType: fwtypes.StringEnumType[odbtypes.DayOfWeekName](),
					},
					"hours_of_day": types.SetType{
						ElemType: types.Int32Type,
					},
					"lead_time_in_weeks": types.Int32Type,
					"months": types.SetType{
						ElemType: fwtypes.StringEnumType[odbtypes.MonthName](),
					},
					"preference": fwtypes.StringEnumType[odbtypes.PreferenceType](),
					"weeks_of_month": types.SetType{
						ElemType: types.Int32Type,
					},
				},
			},
		},
	}
}

func (d *dataSourceCloudAutonomousVmCluster) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	conn := d.Meta().ODBClient(ctx)

	var data cloudAutonomousVmClusterDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	input := odb.GetCloudAutonomousVmClusterInput{
		CloudAutonomousVmClusterId: data.CloudAutonomousVmClusterId.ValueStringPointer(),
	}

	out, err := conn.GetCloudAutonomousVmCluster(ctx, &input)
	if err != nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.ODB, create.ErrActionReading, DSNameCloudAutonomousVmCluster, data.CloudAutonomousVmClusterId.ValueString(), err),
			err.Error(),
		)
		return
	}
	data.CreatedAt = types.StringValue(out.CloudAutonomousVmCluster.CreatedAt.Format(time.RFC3339))
	if out.CloudAutonomousVmCluster.TimeOrdsCertificateExpires != nil {
		data.TimeOrdsCertificateExpires = types.StringValue(out.CloudAutonomousVmCluster.TimeOrdsCertificateExpires.Format(time.RFC3339))
	} else {
		data.TimeOrdsCertificateExpires = types.StringValue(NotAvailableValues)
	}
	if out.CloudAutonomousVmCluster.TimeDatabaseSslCertificateExpires != nil {
		data.TimeDatabaseSslCertificateExpires = types.StringValue(out.CloudAutonomousVmCluster.TimeDatabaseSslCertificateExpires.Format(time.RFC3339))
	} else {
		data.TimeDatabaseSslCertificateExpires = types.StringValue(NotAvailableValues)
	}
	tagsRead, err := listTags(ctx, conn, *out.CloudAutonomousVmCluster.CloudAutonomousVmClusterArn)
	if err != nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.ODB, create.ErrActionReading, DSNameCloudAutonomousVmCluster, data.CloudAutonomousVmClusterId.ValueString(), err),
			err.Error(),
		)
		return
	}
	if tagsRead != nil {
		data.Tags = tftags.FlattenStringValueMap(ctx, tagsRead.Map())
	}

	data.MaintenanceWindow = d.flattenMaintenanceWindow(ctx, out.CloudAutonomousVmCluster.MaintenanceWindow)

	resp.Diagnostics.Append(flex.Flatten(ctx, out.CloudAutonomousVmCluster, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *dataSourceCloudAutonomousVmCluster) flattenMaintenanceWindow(ctx context.Context, avmcMW *odbtypes.MaintenanceWindow) fwtypes.ObjectValueOf[cloudAutonomousVmClusterMaintenanceWindowDataSourceModel] {
	//days of week
	computedMW := cloudAutonomousVmClusterMaintenanceWindowDataSourceModel{}
	if avmcMW.DaysOfWeek != nil {
		daysOfWeek := make([]attr.Value, 0, len(avmcMW.DaysOfWeek))
		for _, dayOfWeek := range avmcMW.DaysOfWeek {
			dayOfWeekStringValue := fwtypes.StringEnumValue(dayOfWeek.Name).StringValue
			daysOfWeek = append(daysOfWeek, dayOfWeekStringValue)
		}

		setValueOfDaysOfWeek, _ := basetypes.NewSetValue(types.StringType, daysOfWeek)
		daysOfWeekRead := fwtypes.SetValueOf[fwtypes.StringEnum[odbtypes.DayOfWeekName]]{
			SetValue: setValueOfDaysOfWeek,
		}
		computedMW.DaysOfWeek = daysOfWeekRead
	}

	//hours of the day
	if avmcMW.HoursOfDay != nil {
		hoursOfTheDay := make([]attr.Value, 0, len(avmcMW.HoursOfDay))
		for _, hourOfTheDay := range avmcMW.HoursOfDay {
			daysOfWeekInt32Value := types.Int32Value(hourOfTheDay)
			hoursOfTheDay = append(hoursOfTheDay, daysOfWeekInt32Value)
		}
		setValuesOfHoursOfTheDay, _ := basetypes.NewSetValue(types.Int32Type, hoursOfTheDay)
		hoursOfTheDayRead := fwtypes.SetValueOf[types.Int32]{
			SetValue: setValuesOfHoursOfTheDay,
		}
		computedMW.HoursOfDay = hoursOfTheDayRead
	}

	//months
	if avmcMW.Months != nil {
		months := make([]attr.Value, 0, len(avmcMW.Months))
		for _, month := range avmcMW.Months {
			monthStringValue := fwtypes.StringEnumValue(month.Name).StringValue
			months = append(months, monthStringValue)
		}
		setValuesOfMonth, _ := basetypes.NewSetValue(types.StringType, months)
		monthsRead := fwtypes.SetValueOf[fwtypes.StringEnum[odbtypes.MonthName]]{
			SetValue: setValuesOfMonth,
		}
		computedMW.Months = monthsRead
	}

	//weeks of month
	if avmcMW.WeeksOfMonth != nil {
		weeksOfMonth := make([]attr.Value, 0, len(avmcMW.WeeksOfMonth))
		for _, weekOfMonth := range avmcMW.WeeksOfMonth {
			weeksOfMonthInt32Value := types.Int32Value(weekOfMonth)
			weeksOfMonth = append(weeksOfMonth, weeksOfMonthInt32Value)
		}
		setValuesOfWeekOfMonth, _ := basetypes.NewSetValue(types.Int32Type, weeksOfMonth)
		weeksOfMonthRead := fwtypes.SetValueOf[types.Int32]{
			SetValue: setValuesOfWeekOfMonth,
		}
		computedMW.WeeksOfMonth = weeksOfMonthRead
	}

	computedMW.LeadTimeInWeeks = types.Int32PointerValue(avmcMW.LeadTimeInWeeks)
	computedMW.Preference = fwtypes.StringEnumValue(avmcMW.Preference)

	result, _ := fwtypes.NewObjectValueOf[cloudAutonomousVmClusterMaintenanceWindowDataSourceModel](ctx, &computedMW)
	return result
}

type cloudAutonomousVmClusterDataSourceModel struct {
	framework.WithRegionModel
	CloudAutonomousVmClusterArn                  types.String                                                                    `tfsdk:"arn"`
	CloudAutonomousVmClusterId                   types.String                                                                    `tfsdk:"id"`
	CloudExadataInfrastructureId                 types.String                                                                    `tfsdk:"cloud_exadata_infrastructure_id"`
	AutonomousDataStoragePercentage              types.Float32                                                                   `tfsdk:"autonomous_data_storage_percentage"`
	AutonomousDataStorageSizeInTBs               types.Float64                                                                   `tfsdk:"autonomous_data_storage_size_in_tbs"`
	AvailableAutonomousDataStorageSizeInTBs      types.Float64                                                                   `tfsdk:"available_autonomous_data_storage_size_in_tbs"`
	AvailableContainerDatabases                  types.Int32                                                                     `tfsdk:"available_container_databases"`
	AvailableCpus                                types.Float32                                                                   `tfsdk:"available_cpus"`
	ComputeModel                                 fwtypes.StringEnum[odbtypes.ComputeModel]                                       `tfsdk:"compute_model"`
	CpuCoreCount                                 types.Int32                                                                     `tfsdk:"cpu_core_count"`
	CpuCoreCountPerNode                          types.Int32                                                                     `tfsdk:"cpu_core_count_per_node"`
	CpuPercentage                                types.Float32                                                                   `tfsdk:"cpu_percentage"`
	CreatedAt                                    types.String                                                                    `tfsdk:"created_at" autoflex:",noflatten"`
	DataStorageSizeInGBs                         types.Float64                                                                   `tfsdk:"data_storage_size_in_gbs"`
	DataStorageSizeInTBs                         types.Float64                                                                   `tfsdk:"data_storage_size_in_tbs"`
	DbNodeStorageSizeInGBs                       types.Int32                                                                     `tfsdk:"odb_node_storage_size_in_gbs"`
	DbServers                                    fwtypes.SetValueOf[types.String]                                                `tfsdk:"db_servers"`
	Description                                  types.String                                                                    `tfsdk:"description"`
	DisplayName                                  types.String                                                                    `tfsdk:"display_name"`
	Domain                                       types.String                                                                    `tfsdk:"domain"`
	ExadataStorageInTBsLowestScaledValue         types.Float64                                                                   `tfsdk:"exadata_storage_in_tbs_lowest_scaled_value"`
	Hostname                                     types.String                                                                    `tfsdk:"hostname"`
	IsMtlsEnabledVmCluster                       types.Bool                                                                      `tfsdk:"is_mtls_enabled_vm_cluster"`
	LicenseModel                                 fwtypes.StringEnum[odbtypes.LicenseModel]                                       `tfsdk:"license_model"`
	MaxAcdsLowestScaledValue                     types.Int32                                                                     `tfsdk:"max_acds_lowest_scaled_value"`
	MemoryPerOracleComputeUnitInGBs              types.Int32                                                                     `tfsdk:"memory_per_oracle_compute_unit_in_gbs"`
	MemorySizeInGBs                              types.Int32                                                                     `tfsdk:"memory_size_in_gbs"`
	NodeCount                                    types.Int32                                                                     `tfsdk:"node_count"`
	NonProvisionableAutonomousContainerDatabases types.Int32                                                                     `tfsdk:"non_provisionable_autonomous_container_databases"`
	OciResourceAnchorName                        types.String                                                                    `tfsdk:"oci_resource_anchor_name"`
	OciUrl                                       types.String                                                                    `tfsdk:"oci_url"`
	Ocid                                         types.String                                                                    `tfsdk:"ocid"`
	OdbNetworkId                                 types.String                                                                    `tfsdk:"odb_network_id"`
	PercentProgress                              types.Float32                                                                   `tfsdk:"percent_progress"`
	ProvisionableAutonomousContainerDatabases    types.Int32                                                                     `tfsdk:"provisionable_autonomous_container_databases"`
	ProvisionedAutonomousContainerDatabases      types.Int32                                                                     `tfsdk:"provisioned_autonomous_container_databases"`
	ProvisionedCpus                              types.Float32                                                                   `tfsdk:"provisioned_cpus"`
	ReclaimableCpus                              types.Float32                                                                   `tfsdk:"reclaimable_cpus"`
	ReservedCpus                                 types.Float32                                                                   `tfsdk:"reserved_cpus"`
	ScanListenerPortNonTls                       types.Int32                                                                     `tfsdk:"scan_listener_port_non_tls"`
	ScanListenerPortTls                          types.Int32                                                                     `tfsdk:"scan_listener_port_tls"`
	Shape                                        types.String                                                                    `tfsdk:"shape"`
	Status                                       fwtypes.StringEnum[odbtypes.ResourceStatus]                                     `tfsdk:"status"`
	StatusReason                                 types.String                                                                    `tfsdk:"reason"`
	TimeDatabaseSslCertificateExpires            types.String                                                                    `tfsdk:"time_database_ssl_certificate_expires" autoflex:",noflatten"`
	TimeOrdsCertificateExpires                   types.String                                                                    `tfsdk:"time_ords_certificate_expires" autoflex:",noflatten"`
	TimeZone                                     types.String                                                                    `tfsdk:"time_zone"`
	TotalAutonomousDataStorageInTBs              types.Float32                                                                   `tfsdk:"total_autonomous_data_storage_in_tbs"`
	TotalContainerDatabases                      types.Int32                                                                     `tfsdk:"total_container_databases"`
	TotalCpus                                    types.Float32                                                                   `tfsdk:"total_cpus"`
	MaintenanceWindow                            fwtypes.ObjectValueOf[cloudAutonomousVmClusterMaintenanceWindowDataSourceModel] `tfsdk:"maintenance_window" autoflex:",noflatten"`
	Tags                                         tftags.Map                                                                      `tfsdk:"tags"`
}
type cloudAutonomousVmClusterMaintenanceWindowDataSourceModel struct {
	DaysOfWeek      fwtypes.SetValueOf[fwtypes.StringEnum[odbtypes.DayOfWeekName]] `tfsdk:"days_of_week"`
	HoursOfDay      fwtypes.SetValueOf[types.Int32]                                `tfsdk:"hours_of_day"`
	LeadTimeInWeeks types.Int32                                                    `tfsdk:"lead_time_in_weeks"`
	Months          fwtypes.SetValueOf[fwtypes.StringEnum[odbtypes.MonthName]]     `tfsdk:"months"`
	Preference      fwtypes.StringEnum[odbtypes.PreferenceType]                    `tfsdk:"preference"`
	WeeksOfMonth    fwtypes.SetValueOf[types.Int32]                                `tfsdk:"weeks_of_month"`
}
