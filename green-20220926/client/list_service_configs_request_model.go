// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassify(v string) *ListServiceConfigsRequest
	GetClassify() *string
	SetRegionId(v string) *ListServiceConfigsRequest
	GetRegionId() *string
	SetResourceType(v string) *ListServiceConfigsRequest
	GetResourceType() *string
	SetUseStatus(v string) *ListServiceConfigsRequest
	GetUseStatus() *string
}

type ListServiceConfigsRequest struct {
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	UseStatus    *string `json:"UseStatus,omitempty" xml:"UseStatus,omitempty"`
}

func (s ListServiceConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceConfigsRequest) GetClassify() *string {
	return s.Classify
}

func (s *ListServiceConfigsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListServiceConfigsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListServiceConfigsRequest) GetUseStatus() *string {
	return s.UseStatus
}

func (s *ListServiceConfigsRequest) SetClassify(v string) *ListServiceConfigsRequest {
	s.Classify = &v
	return s
}

func (s *ListServiceConfigsRequest) SetRegionId(v string) *ListServiceConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceConfigsRequest) SetResourceType(v string) *ListServiceConfigsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListServiceConfigsRequest) SetUseStatus(v string) *ListServiceConfigsRequest {
	s.UseStatus = &v
	return s
}

func (s *ListServiceConfigsRequest) Validate() error {
	return dara.Validate(s)
}
