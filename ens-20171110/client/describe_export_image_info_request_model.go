// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExportImageInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *DescribeExportImageInfoRequest
	GetImageId() *string
	SetImageName(v string) *DescribeExportImageInfoRequest
	GetImageName() *string
	SetPageNumber(v int32) *DescribeExportImageInfoRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeExportImageInfoRequest
	GetPageSize() *int32
}

type DescribeExportImageInfoRequest struct {
	// The ID of the image.
	//
	// example:
	//
	// m-xxxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// testImageName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeExportImageInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportImageInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeExportImageInfoRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeExportImageInfoRequest) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeExportImageInfoRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeExportImageInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeExportImageInfoRequest) SetImageId(v string) *DescribeExportImageInfoRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeExportImageInfoRequest) SetImageName(v string) *DescribeExportImageInfoRequest {
	s.ImageName = &v
	return s
}

func (s *DescribeExportImageInfoRequest) SetPageNumber(v int32) *DescribeExportImageInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeExportImageInfoRequest) SetPageSize(v int32) *DescribeExportImageInfoRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeExportImageInfoRequest) Validate() error {
	return dara.Validate(s)
}
