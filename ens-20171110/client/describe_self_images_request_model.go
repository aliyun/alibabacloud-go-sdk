// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSelfImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *DescribeSelfImagesRequest
	GetImageId() *string
	SetImageName(v string) *DescribeSelfImagesRequest
	GetImageName() *string
	SetPageNumber(v int32) *DescribeSelfImagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSelfImagesRequest
	GetPageSize() *int32
	SetSnapshotId(v string) *DescribeSelfImagesRequest
	GetSnapshotId() *string
}

type DescribeSelfImagesRequest struct {
	// The ID of the image. Fuzzy search is supported.
	//
	// example:
	//
	// centos_6_08_64_20G_a****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image. Fuzzy search is supported.
	//
	// example:
	//
	// centos_6_08_64_20G_a****
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The page number to return. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the snapshot.
	//
	// example:
	//
	// sp-5xg63dmojc1oaa3pk****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DescribeSelfImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSelfImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSelfImagesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeSelfImagesRequest) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeSelfImagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSelfImagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSelfImagesRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeSelfImagesRequest) SetImageId(v string) *DescribeSelfImagesRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeSelfImagesRequest) SetImageName(v string) *DescribeSelfImagesRequest {
	s.ImageName = &v
	return s
}

func (s *DescribeSelfImagesRequest) SetPageNumber(v int32) *DescribeSelfImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSelfImagesRequest) SetPageSize(v int32) *DescribeSelfImagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSelfImagesRequest) SetSnapshotId(v string) *DescribeSelfImagesRequest {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSelfImagesRequest) Validate() error {
	return dara.Validate(s)
}
