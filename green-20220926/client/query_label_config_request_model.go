// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLabelConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassify(v string) *QueryLabelConfigRequest
	GetClassify() *string
	SetRegionId(v string) *QueryLabelConfigRequest
	GetRegionId() *string
	SetResourceType(v string) *QueryLabelConfigRequest
	GetResourceType() *string
	SetServiceCode(v string) *QueryLabelConfigRequest
	GetServiceCode() *string
	SetType(v string) *QueryLabelConfigRequest
	GetType() *string
}

type QueryLabelConfigRequest struct {
	// The classification. Separate multiple values with commas.
	//
	// example:
	//
	// guard-scene
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource type. Separate multiple values with commas.
	//
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The service code. Separate multiple values with commas.
	//
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The type. Separate multiple values with commas.
	//
	// content_moderation
	//
	// example:
	//
	// content_moderation
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryLabelConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLabelConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryLabelConfigRequest) GetClassify() *string {
	return s.Classify
}

func (s *QueryLabelConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryLabelConfigRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *QueryLabelConfigRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *QueryLabelConfigRequest) GetType() *string {
	return s.Type
}

func (s *QueryLabelConfigRequest) SetClassify(v string) *QueryLabelConfigRequest {
	s.Classify = &v
	return s
}

func (s *QueryLabelConfigRequest) SetRegionId(v string) *QueryLabelConfigRequest {
	s.RegionId = &v
	return s
}

func (s *QueryLabelConfigRequest) SetResourceType(v string) *QueryLabelConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *QueryLabelConfigRequest) SetServiceCode(v string) *QueryLabelConfigRequest {
	s.ServiceCode = &v
	return s
}

func (s *QueryLabelConfigRequest) SetType(v string) *QueryLabelConfigRequest {
	s.Type = &v
	return s
}

func (s *QueryLabelConfigRequest) Validate() error {
	return dara.Validate(s)
}
