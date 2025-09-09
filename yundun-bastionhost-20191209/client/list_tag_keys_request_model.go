// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTagKeysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTagKeysRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListTagKeysRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListTagKeysRequest
	GetResourceGroupId() *string
	SetResourceType(v string) *ListTagKeysRequest
	GetResourceType() *string
}

type ListTagKeysRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the resource.
	//
	// Set the value to INSTANCE, which indicates that the resource is a bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTagKeysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagKeysRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTagKeysRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagKeysRequest) SetPageNumber(v int32) *ListTagKeysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceGroupId(v string) *ListTagKeysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) Validate() error {
	return dara.Validate(s)
}
