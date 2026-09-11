// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTerraformPricingMappingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMappingVersion(v string) *GetTerraformPricingMappingsResponseBody
	GetMappingVersion() *string
	SetMappings(v []*GetTerraformPricingMappingsResponseBodyMappings) *GetTerraformPricingMappingsResponseBody
	GetMappings() []*GetTerraformPricingMappingsResponseBodyMappings
	SetMissingResourceTypes(v []*string) *GetTerraformPricingMappingsResponseBody
	GetMissingResourceTypes() []*string
	SetRequestId(v string) *GetTerraformPricingMappingsResponseBody
	GetRequestId() *string
	SetSchemaVersion(v string) *GetTerraformPricingMappingsResponseBody
	GetSchemaVersion() *string
	SetUnsupportedResourceTypes(v []*string) *GetTerraformPricingMappingsResponseBody
	GetUnsupportedResourceTypes() []*string
}

type GetTerraformPricingMappingsResponseBody struct {
	// The mapping content version, which is the timestamp of the most recent data change. Consumers can use this value for caching and auditing.
	//
	// example:
	//
	// 1786000000000
	MappingVersion *string `json:"mappingVersion,omitempty" xml:"mappingVersion,omitempty"`
	// The list of matched mappings. Each item contains a resourceType and pricingTargets, which include pricing targets and parameter extraction rules. The rules reference Terraform plan resource properties by using $after/$before.
	Mappings []*GetTerraformPricingMappingsResponseBodyMappings `json:"mappings,omitempty" xml:"mappings,omitempty" type:"Repeated"`
	// The resource types in the request that do not have registered mappings. Consumers must treat these as unknown cost. Do not assume they are free.
	MissingResourceTypes []*string `json:"missingResourceTypes,omitempty" xml:"missingResourceTypes,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 1AF0AD89-ED4F-5E9E-8B7B-9A3B27CE9E1B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The schema version of the mapping catalog. The current version is 1.0. Consumers use this value to determine compatibility.
	//
	// example:
	//
	// 1.0
	SchemaVersion *string `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
	// The resource types that are confirmed to not support pricing, such as free resources or resources without a pricing interface. These are different from missing resource types.
	UnsupportedResourceTypes []*string `json:"unsupportedResourceTypes,omitempty" xml:"unsupportedResourceTypes,omitempty" type:"Repeated"`
}

func (s GetTerraformPricingMappingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTerraformPricingMappingsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTerraformPricingMappingsResponseBody) GetMappingVersion() *string {
	return s.MappingVersion
}

func (s *GetTerraformPricingMappingsResponseBody) GetMappings() []*GetTerraformPricingMappingsResponseBodyMappings {
	return s.Mappings
}

func (s *GetTerraformPricingMappingsResponseBody) GetMissingResourceTypes() []*string {
	return s.MissingResourceTypes
}

func (s *GetTerraformPricingMappingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTerraformPricingMappingsResponseBody) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *GetTerraformPricingMappingsResponseBody) GetUnsupportedResourceTypes() []*string {
	return s.UnsupportedResourceTypes
}

func (s *GetTerraformPricingMappingsResponseBody) SetMappingVersion(v string) *GetTerraformPricingMappingsResponseBody {
	s.MappingVersion = &v
	return s
}

func (s *GetTerraformPricingMappingsResponseBody) SetMappings(v []*GetTerraformPricingMappingsResponseBodyMappings) *GetTerraformPricingMappingsResponseBody {
	s.Mappings = v
	return s
}

func (s *GetTerraformPricingMappingsResponseBody) SetMissingResourceTypes(v []*string) *GetTerraformPricingMappingsResponseBody {
	s.MissingResourceTypes = v
	return s
}

func (s *GetTerraformPricingMappingsResponseBody) SetRequestId(v string) *GetTerraformPricingMappingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTerraformPricingMappingsResponseBody) SetSchemaVersion(v string) *GetTerraformPricingMappingsResponseBody {
	s.SchemaVersion = &v
	return s
}

func (s *GetTerraformPricingMappingsResponseBody) SetUnsupportedResourceTypes(v []*string) *GetTerraformPricingMappingsResponseBody {
	s.UnsupportedResourceTypes = v
	return s
}

func (s *GetTerraformPricingMappingsResponseBody) Validate() error {
	if s.Mappings != nil {
		for _, item := range s.Mappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTerraformPricingMappingsResponseBodyMappings struct {
	// The list of pricing targets. Each item contains actions (create/update), popCode/popVersion/apiName, pricingUnit, params (parameter extraction rules: from=$after.xxx / const / default / expand), and when/whenChanged conditions.
	PricingTargets []map[string]interface{} `json:"pricingTargets,omitempty" xml:"pricingTargets,omitempty" type:"Repeated"`
	// The Terraform resource type, such as alicloud_instance.
	//
	// example:
	//
	// alicloud_instance
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetTerraformPricingMappingsResponseBodyMappings) String() string {
	return dara.Prettify(s)
}

func (s GetTerraformPricingMappingsResponseBodyMappings) GoString() string {
	return s.String()
}

func (s *GetTerraformPricingMappingsResponseBodyMappings) GetPricingTargets() []map[string]interface{} {
	return s.PricingTargets
}

func (s *GetTerraformPricingMappingsResponseBodyMappings) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetTerraformPricingMappingsResponseBodyMappings) SetPricingTargets(v []map[string]interface{}) *GetTerraformPricingMappingsResponseBodyMappings {
	s.PricingTargets = v
	return s
}

func (s *GetTerraformPricingMappingsResponseBodyMappings) SetResourceType(v string) *GetTerraformPricingMappingsResponseBodyMappings {
	s.ResourceType = &v
	return s
}

func (s *GetTerraformPricingMappingsResponseBodyMappings) Validate() error {
	return dara.Validate(s)
}
