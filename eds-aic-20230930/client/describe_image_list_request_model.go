// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageBizTags(v []*DescribeImageListRequestImageBizTags) *DescribeImageListRequest
	GetImageBizTags() []*DescribeImageListRequestImageBizTags
	SetImageId(v string) *DescribeImageListRequest
	GetImageId() *string
	SetImageName(v string) *DescribeImageListRequest
	GetImageName() *string
	SetImagePackageType(v string) *DescribeImageListRequest
	GetImagePackageType() *string
	SetImageType(v string) *DescribeImageListRequest
	GetImageType() *string
	SetInstanceType(v string) *DescribeImageListRequest
	GetInstanceType() *string
	SetMaxResults(v int32) *DescribeImageListRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeImageListRequest
	GetNextToken() *string
	SetStatus(v string) *DescribeImageListRequest
	GetStatus() *string
}

type DescribeImageListRequest struct {
	ImageBizTags []*DescribeImageListRequestImageBizTags `json:"ImageBizTags,omitempty" xml:"ImageBizTags,omitempty" type:"Repeated"`
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
	ImageType    *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
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

func (s *DescribeImageListRequest) GetImageBizTags() []*DescribeImageListRequestImageBizTags {
	return s.ImageBizTags
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

func (s *DescribeImageListRequest) GetInstanceType() *string {
	return s.InstanceType
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

func (s *DescribeImageListRequest) SetImageBizTags(v []*DescribeImageListRequestImageBizTags) *DescribeImageListRequest {
	s.ImageBizTags = v
	return s
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

func (s *DescribeImageListRequest) SetInstanceType(v string) *DescribeImageListRequest {
	s.InstanceType = &v
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

type DescribeImageListRequestImageBizTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImageListRequestImageBizTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListRequestImageBizTags) GoString() string {
	return s.String()
}

func (s *DescribeImageListRequestImageBizTags) GetKey() *string {
	return s.Key
}

func (s *DescribeImageListRequestImageBizTags) GetValue() *string {
	return s.Value
}

func (s *DescribeImageListRequestImageBizTags) SetKey(v string) *DescribeImageListRequestImageBizTags {
	s.Key = &v
	return s
}

func (s *DescribeImageListRequestImageBizTags) SetValue(v string) *DescribeImageListRequestImageBizTags {
	s.Value = &v
	return s
}

func (s *DescribeImageListRequestImageBizTags) Validate() error {
	return dara.Validate(s)
}
