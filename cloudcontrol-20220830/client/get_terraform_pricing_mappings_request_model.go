// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTerraformPricingMappingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v map[string]interface{}) *GetTerraformPricingMappingsRequest
	GetBody() map[string]interface{}
}

type GetTerraformPricingMappingsRequest struct {
	// The request body. The resourceTypes field specifies a list of Terraform resource types, such as alicloud_instance. You can specify a maximum of 200 resource types in a single request.
	//
	// example:
	//
	// {
	//
	//   "resourceTypes": [
	//
	//     "alicloud_instance",
	//
	//     "alicloud_vpc"
	//
	//   ]
	//
	// }
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTerraformPricingMappingsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTerraformPricingMappingsRequest) GoString() string {
	return s.String()
}

func (s *GetTerraformPricingMappingsRequest) GetBody() map[string]interface{} {
	return s.Body
}

func (s *GetTerraformPricingMappingsRequest) SetBody(v map[string]interface{}) *GetTerraformPricingMappingsRequest {
	s.Body = v
	return s
}

func (s *GetTerraformPricingMappingsRequest) Validate() error {
	return dara.Validate(s)
}
