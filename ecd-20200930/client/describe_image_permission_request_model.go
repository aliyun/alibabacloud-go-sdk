// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImagePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *DescribeImagePermissionRequest
	GetImageId() *string
	SetRegionId(v string) *DescribeImagePermissionRequest
	GetRegionId() *string
}

type DescribeImagePermissionRequest struct {
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-4zfb6zj728hhr****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeImagePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePermissionRequest) GoString() string {
	return s.String()
}

func (s *DescribeImagePermissionRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImagePermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImagePermissionRequest) SetImageId(v string) *DescribeImagePermissionRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeImagePermissionRequest) SetRegionId(v string) *DescribeImagePermissionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImagePermissionRequest) Validate() error {
	return dara.Validate(s)
}
