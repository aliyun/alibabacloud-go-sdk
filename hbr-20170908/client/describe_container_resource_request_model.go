// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeContainerResourceRequest
	GetClusterId() *string
	SetPageNumber(v int32) *DescribeContainerResourceRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeContainerResourceRequest
	GetPageSize() *int32
	SetResourceId(v string) *DescribeContainerResourceRequest
	GetResourceId() *string
	SetResourceType(v string) *DescribeContainerResourceRequest
	GetResourceType() *string
}

type DescribeContainerResourceRequest struct {
	// example:
	//
	// cc-0005**********hhjw
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// a9ab843d-****-****-8e46-1d67a82128a7
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// PV
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeContainerResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerResourceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeContainerResourceRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeContainerResourceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeContainerResourceRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeContainerResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeContainerResourceRequest) SetClusterId(v string) *DescribeContainerResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeContainerResourceRequest) SetPageNumber(v int32) *DescribeContainerResourceRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeContainerResourceRequest) SetPageSize(v int32) *DescribeContainerResourceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeContainerResourceRequest) SetResourceId(v string) *DescribeContainerResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeContainerResourceRequest) SetResourceType(v string) *DescribeContainerResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeContainerResourceRequest) Validate() error {
	return dara.Validate(s)
}
