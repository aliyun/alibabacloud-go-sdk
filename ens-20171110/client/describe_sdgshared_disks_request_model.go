// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDGSharedDisksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *DescribeSDGSharedDisksRequest
	GetNamespace() *string
	SetPageNumber(v int32) *DescribeSDGSharedDisksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSDGSharedDisksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSDGSharedDisksRequest
	GetRegionId() *string
	SetSdgId(v string) *DescribeSDGSharedDisksRequest
	GetSdgId() *string
}

type DescribeSDGSharedDisksRequest struct {
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
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
	// cn-hangzhou-xx
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sdg-xxxx
	SdgId *string `json:"SdgId,omitempty" xml:"SdgId,omitempty"`
}

func (s DescribeSDGSharedDisksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGSharedDisksRequest) GoString() string {
	return s.String()
}

func (s *DescribeSDGSharedDisksRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeSDGSharedDisksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSDGSharedDisksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSDGSharedDisksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSDGSharedDisksRequest) GetSdgId() *string {
	return s.SdgId
}

func (s *DescribeSDGSharedDisksRequest) SetNamespace(v string) *DescribeSDGSharedDisksRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeSDGSharedDisksRequest) SetPageNumber(v int32) *DescribeSDGSharedDisksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSDGSharedDisksRequest) SetPageSize(v int32) *DescribeSDGSharedDisksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSDGSharedDisksRequest) SetRegionId(v string) *DescribeSDGSharedDisksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSDGSharedDisksRequest) SetSdgId(v string) *DescribeSDGSharedDisksRequest {
	s.SdgId = &v
	return s
}

func (s *DescribeSDGSharedDisksRequest) Validate() error {
	return dara.Validate(s)
}
