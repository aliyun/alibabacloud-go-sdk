// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFeatureConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuery(v string) *GetFeatureConfigRequest
	GetQuery() *string
	SetRegionId(v string) *GetFeatureConfigRequest
	GetRegionId() *string
	SetResourceType(v string) *GetFeatureConfigRequest
	GetResourceType() *string
	SetServiceCode(v string) *GetFeatureConfigRequest
	GetServiceCode() *string
	SetType(v string) *GetFeatureConfigRequest
	GetType() *string
}

type GetFeatureConfigRequest struct {
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// Region ID
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

func (s GetFeatureConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureConfigRequest) GoString() string {
	return s.String()
}

func (s *GetFeatureConfigRequest) GetQuery() *string {
	return s.Query
}

func (s *GetFeatureConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetFeatureConfigRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetFeatureConfigRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetFeatureConfigRequest) GetType() *string {
	return s.Type
}

func (s *GetFeatureConfigRequest) SetQuery(v string) *GetFeatureConfigRequest {
	s.Query = &v
	return s
}

func (s *GetFeatureConfigRequest) SetRegionId(v string) *GetFeatureConfigRequest {
	s.RegionId = &v
	return s
}

func (s *GetFeatureConfigRequest) SetResourceType(v string) *GetFeatureConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *GetFeatureConfigRequest) SetServiceCode(v string) *GetFeatureConfigRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetFeatureConfigRequest) SetType(v string) *GetFeatureConfigRequest {
	s.Type = &v
	return s
}

func (s *GetFeatureConfigRequest) Validate() error {
	return dara.Validate(s)
}
