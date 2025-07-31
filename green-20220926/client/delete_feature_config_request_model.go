// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFeatureConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetField(v string) *DeleteFeatureConfigRequest
	GetField() *string
	SetRegionId(v string) *DeleteFeatureConfigRequest
	GetRegionId() *string
	SetResourceType(v string) *DeleteFeatureConfigRequest
	GetResourceType() *string
	SetServiceCode(v string) *DeleteFeatureConfigRequest
	GetServiceCode() *string
	SetType(v string) *DeleteFeatureConfigRequest
	GetType() *string
}

type DeleteFeatureConfigRequest struct {
	// Label value, customer-defined
	//
	// example:
	//
	// __config__
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// Region
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource type.
	//
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Service code.
	//
	// example:
	//
	// llm_query_moderation
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Type
	//
	// example:
	//
	// custom_llm_template
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteFeatureConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFeatureConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteFeatureConfigRequest) GetField() *string {
	return s.Field
}

func (s *DeleteFeatureConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteFeatureConfigRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeleteFeatureConfigRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DeleteFeatureConfigRequest) GetType() *string {
	return s.Type
}

func (s *DeleteFeatureConfigRequest) SetField(v string) *DeleteFeatureConfigRequest {
	s.Field = &v
	return s
}

func (s *DeleteFeatureConfigRequest) SetRegionId(v string) *DeleteFeatureConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFeatureConfigRequest) SetResourceType(v string) *DeleteFeatureConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *DeleteFeatureConfigRequest) SetServiceCode(v string) *DeleteFeatureConfigRequest {
	s.ServiceCode = &v
	return s
}

func (s *DeleteFeatureConfigRequest) SetType(v string) *DeleteFeatureConfigRequest {
	s.Type = &v
	return s
}

func (s *DeleteFeatureConfigRequest) Validate() error {
	return dara.Validate(s)
}
