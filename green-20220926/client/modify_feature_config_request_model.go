// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFeatureConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyFeatureConfigRequest
	GetConfig() *string
	SetDescription(v string) *ModifyFeatureConfigRequest
	GetDescription() *string
	SetField(v string) *ModifyFeatureConfigRequest
	GetField() *string
	SetRegionId(v string) *ModifyFeatureConfigRequest
	GetRegionId() *string
	SetResourceType(v string) *ModifyFeatureConfigRequest
	GetResourceType() *string
	SetServiceCode(v string) *ModifyFeatureConfigRequest
	GetServiceCode() *string
	SetType(v string) *ModifyFeatureConfigRequest
	GetType() *string
}

type ModifyFeatureConfigRequest struct {
	// Configuration, in JSON format
	//
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Label meaning
	//
	// example:
	//
	// 标签2
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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

func (s ModifyFeatureConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFeatureConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyFeatureConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyFeatureConfigRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyFeatureConfigRequest) GetField() *string {
	return s.Field
}

func (s *ModifyFeatureConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyFeatureConfigRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ModifyFeatureConfigRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *ModifyFeatureConfigRequest) GetType() *string {
	return s.Type
}

func (s *ModifyFeatureConfigRequest) SetConfig(v string) *ModifyFeatureConfigRequest {
	s.Config = &v
	return s
}

func (s *ModifyFeatureConfigRequest) SetDescription(v string) *ModifyFeatureConfigRequest {
	s.Description = &v
	return s
}

func (s *ModifyFeatureConfigRequest) SetField(v string) *ModifyFeatureConfigRequest {
	s.Field = &v
	return s
}

func (s *ModifyFeatureConfigRequest) SetRegionId(v string) *ModifyFeatureConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFeatureConfigRequest) SetResourceType(v string) *ModifyFeatureConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyFeatureConfigRequest) SetServiceCode(v string) *ModifyFeatureConfigRequest {
	s.ServiceCode = &v
	return s
}

func (s *ModifyFeatureConfigRequest) SetType(v string) *ModifyFeatureConfigRequest {
	s.Type = &v
	return s
}

func (s *ModifyFeatureConfigRequest) Validate() error {
	return dara.Validate(s)
}
