// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCImageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitecture(v string) *DescribeRCImageListRequest
	GetArchitecture() *string
	SetImageId(v string) *DescribeRCImageListRequest
	GetImageId() *string
	SetImageName(v string) *DescribeRCImageListRequest
	GetImageName() *string
	SetInstanceType(v string) *DescribeRCImageListRequest
	GetInstanceType() *string
	SetPageNumber(v int32) *DescribeRCImageListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRCImageListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeRCImageListRequest
	GetRegionId() *string
	SetType(v string) *DescribeRCImageListRequest
	GetType() *string
}

type DescribeRCImageListRequest struct {
	// The image architecture. Valid values:
	//
	// 	- x86_64
	//
	// 	- arm64
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName    *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The image type. Set the value to **self**.
	//
	// example:
	//
	// self
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeRCImageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCImageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCImageListRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeRCImageListRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeRCImageListRequest) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeRCImageListRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeRCImageListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRCImageListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRCImageListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCImageListRequest) GetType() *string {
	return s.Type
}

func (s *DescribeRCImageListRequest) SetArchitecture(v string) *DescribeRCImageListRequest {
	s.Architecture = &v
	return s
}

func (s *DescribeRCImageListRequest) SetImageId(v string) *DescribeRCImageListRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeRCImageListRequest) SetImageName(v string) *DescribeRCImageListRequest {
	s.ImageName = &v
	return s
}

func (s *DescribeRCImageListRequest) SetInstanceType(v string) *DescribeRCImageListRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeRCImageListRequest) SetPageNumber(v int32) *DescribeRCImageListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCImageListRequest) SetPageSize(v int32) *DescribeRCImageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRCImageListRequest) SetRegionId(v string) *DescribeRCImageListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCImageListRequest) SetType(v string) *DescribeRCImageListRequest {
	s.Type = &v
	return s
}

func (s *DescribeRCImageListRequest) Validate() error {
	return dara.Validate(s)
}
