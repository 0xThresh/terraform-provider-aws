//Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.

package odb_test

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/odb"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	tfodb "github.com/hashicorp/terraform-provider-aws/internal/service/odb"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type listOdbNetworkPeering struct {
}

func TestAccODBNetworkPeeringBasic(t *testing.T) {
	ctx := acctest.Context(t)
	var listOfPeeredNwks = listOdbNetworkPeering{}
	var output odb.ListOdbPeeringConnectionsOutput

	dataSourceName := "data.aws_odb_network_peering_connections_list.test"
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			listOfPeeredNwks.testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.ODBServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: listOfPeeredNwks.basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.ComposeTestCheckFunc(func(s *terraform.State) error {
						listOfPeeredNwks.count(ctx, dataSourceName, &output)
						resource.TestCheckResourceAttr(dataSourceName, "odb_peering_connections.#", strconv.Itoa(len(output.OdbPeeringConnections)))
						return nil
					},
					),
				),
			},
		},
	})
}

func (listOdbNetworkPeering) basic() string {
	config := fmt.Sprintf(`

data "aws_odb_network_peering_connections_list" "test" {

}
`)
	return config
}

func (listOdbNetworkPeering) count(ctx context.Context, name string, list *odb.ListOdbPeeringConnectionsOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return create.Error(names.ODB, create.ErrActionCheckingExistence, tfodb.DSNameNetworkPeeringConnectionsList, name, errors.New("not found"))
		}
		conn := acctest.Provider.Meta().(*conns.AWSClient).ODBClient(ctx)
		resp, err := tfodb.ListOdbPeeringConnections(ctx, conn)
		if err != nil {
			return create.Error(names.ODB, create.ErrActionCheckingExistence, tfodb.DSNameNetworkPeeringConnectionsList, rs.Primary.ID, err)
		}
		list.OdbPeeringConnections = resp.OdbPeeringConnections
		return nil
	}
}
func (listOdbNetworkPeering) testAccPreCheck(ctx context.Context, t *testing.T) {
	conn := acctest.Provider.Meta().(*conns.AWSClient).ODBClient(ctx)
	input := &odb.ListOdbPeeringConnectionsInput{}
	_, err := conn.ListOdbPeeringConnections(ctx, input)
	if acctest.PreCheckSkipError(err) {
		t.Skipf("skipping acceptance testing: %s", err)
	}
	if err != nil {
		t.Fatalf("unexpected PreCheck error: %s", err)
	}
}
