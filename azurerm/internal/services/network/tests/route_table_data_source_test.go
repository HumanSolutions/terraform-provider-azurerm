package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance/check"
)

type RouteTableDataSource struct {
}

func TestAccDataSourceRouteTable_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_route_table", "test")
	r := RouteTableDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("route.#").HasValue("0"),
			),
		},
	})
}

func TestAccDataSourceRouteTable_singleRoute(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_route_table", "test")
	r := RouteTableDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.singleRoute(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("route.#").HasValue("1"),
				check.That(data.ResourceName).Key("route.0.name").HasValue("route1"),
				check.That(data.ResourceName).Key("route.0.address_prefix").HasValue("10.1.0.0/16"),
				check.That(data.ResourceName).Key("route.0.next_hop_type").HasValue("VnetLocal"),
			),
		},
	})
}

func TestAccDataSourceRouteTable_multipleRoutes(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_route_table", "test")
	r := RouteTableDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.multipleRoutes(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("route.#").HasValue("2"),
				check.That(data.ResourceName).Key("route.0.name").HasValue("route1"),
				check.That(data.ResourceName).Key("route.0.address_prefix").HasValue("10.1.0.0/16"),
				check.That(data.ResourceName).Key("route.0.next_hop_type").HasValue("VnetLocal"),
				check.That(data.ResourceName).Key("route.1.name").HasValue("route2"),
				check.That(data.ResourceName).Key("route.1.address_prefix").HasValue("10.2.0.0/16"),
				check.That(data.ResourceName).Key("route.1.next_hop_type").HasValue("VnetLocal"),
			),
		},
	})
}

func (RouteTableDataSource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

data "azurerm_route_table" "test" {
  name                = azurerm_route_table.test.name
  resource_group_name = azurerm_route_table.test.resource_group_name
}
`, RouteTableResource{}.basic(data))
}

func (RouteTableDataSource) singleRoute(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

data "azurerm_route_table" "test" {
  name                = azurerm_route_table.test.name
  resource_group_name = azurerm_route_table.test.resource_group_name
}
`, RouteTableResource{}.singleRoute(data))
}

func (RouteTableDataSource) multipleRoutes(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

data "azurerm_route_table" "test" {
  name                = azurerm_route_table.test.name
  resource_group_name = azurerm_route_table.test.resource_group_name
}
`, RouteTableResource{}.multipleRoutes(data))
}
