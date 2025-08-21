// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSharePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunId(v int64) *DescribeImageSharePermissionRequest
	GetAliyunId() *int64
	SetImageId(v string) *DescribeImageSharePermissionRequest
	GetImageId() *string
	SetPageNumber(v string) *DescribeImageSharePermissionRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeImageSharePermissionRequest
	GetPageSize() *string
}

type DescribeImageSharePermissionRequest struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 171710408091****
	AliyunId *int64 `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	// The ID of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-5s7qotzavwbrnzaqh4unm****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The page number. Pages start from page **1**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **100**.
	//
	// Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeImageSharePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSharePermissionRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionRequest) GetAliyunId() *int64 {
	return s.AliyunId
}

func (s *DescribeImageSharePermissionRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageSharePermissionRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeImageSharePermissionRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeImageSharePermissionRequest) SetAliyunId(v int64) *DescribeImageSharePermissionRequest {
	s.AliyunId = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) SetImageId(v string) *DescribeImageSharePermissionRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) SetPageNumber(v string) *DescribeImageSharePermissionRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) SetPageSize(v string) *DescribeImageSharePermissionRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) Validate() error {
	return dara.Validate(s)
}
