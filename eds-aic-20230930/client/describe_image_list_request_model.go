// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *DescribeImageListRequest
	GetImageId() *string
	SetImageName(v string) *DescribeImageListRequest
	GetImageName() *string
	SetImagePackageType(v string) *DescribeImageListRequest
	GetImagePackageType() *string
	SetImageType(v string) *DescribeImageListRequest
	GetImageType() *string
	SetMaxResults(v int32) *DescribeImageListRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeImageListRequest
	GetNextToken() *string
	SetStatus(v string) *DescribeImageListRequest
	GetStatus() *string
}

type DescribeImageListRequest struct {
	// The ID of the image.
	//
	// example:
	//
	// imgc-075cllfeuazh0****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// Android 12 image
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// Image package type.
	//
	// example:
	//
	// VM
	ImagePackageType *string `json:"ImagePackageType,omitempty" xml:"ImagePackageType,omitempty"`
	// The type of the image.
	//
	// Valid values:
	//
	// 	- User: custom images.
	//
	// 	- System: system images.
	//
	// This parameter is required.
	//
	// example:
	//
	// System
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If the parameter is left empty, the data is queried from the first entry.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kw9dGL5jves2FS9RLq****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The state of the image.
	//
	// Valid values:
	//
	// 	- AVAILABLE: The image is available.
	//
	// 	- DELETE: The image is deleted.
	//
	// 	- INIT: The image is being initialized.
	//
	// 	- CREATE_FAILED: The image failed to be created.
	//
	// 	- CREATING: The image is being created.
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeImageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageListRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageListRequest) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeImageListRequest) GetImagePackageType() *string {
	return s.ImagePackageType
}

func (s *DescribeImageListRequest) GetImageType() *string {
	return s.ImageType
}

func (s *DescribeImageListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeImageListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeImageListRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeImageListRequest) SetImageId(v string) *DescribeImageListRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeImageListRequest) SetImageName(v string) *DescribeImageListRequest {
	s.ImageName = &v
	return s
}

func (s *DescribeImageListRequest) SetImagePackageType(v string) *DescribeImageListRequest {
	s.ImagePackageType = &v
	return s
}

func (s *DescribeImageListRequest) SetImageType(v string) *DescribeImageListRequest {
	s.ImageType = &v
	return s
}

func (s *DescribeImageListRequest) SetMaxResults(v int32) *DescribeImageListRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeImageListRequest) SetNextToken(v string) *DescribeImageListRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeImageListRequest) SetStatus(v string) *DescribeImageListRequest {
	s.Status = &v
	return s
}

func (s *DescribeImageListRequest) Validate() error {
	return dara.Validate(s)
}
