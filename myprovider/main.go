package main

import (
	//"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return &schema.Provider{
				ResourcesMap: map[string]*schema.Resource{
					"myprovider_resource": ProvidersResource(),
				},
			}
		},
	})
}

func ProvidersResource() *schema.Resource {
	return &schema.Resource{
		Create: ProvidersResourceCreate,
		Read:   ProvidersResourceRead,
		Update: ProvidersResourceUpdate,
		Delete: ProvidersResourceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ProvidersResourceCreate(r *schema.ResourceData, my interface{}) error {
	name := r.Get("name").(string)
	r.SetId(name)
	return nil
}

func ProvidersResourceRead(r *schema.ResourceData, my interface{}) error {
	return nil
}

func ProvidersResourceUpdate(r *schema.ResourceData, my interface{}) error {
	return nil
}

func ProvidersResourceDelete(r *schema.ResourceData, my interface{}) error {
	return nil
}
