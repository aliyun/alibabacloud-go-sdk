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
	SetSystemType(v string) *DescribeImageListRequest
	GetSystemType() *string
}

type DescribeImageListRequest struct {
	// An array of tag objects.
	ImageBizTags []*DescribeImageListRequestImageBizTags `json:"ImageBizTags,omitempty" xml:"ImageBizTags,omitempty" type:"Repeated"`
	// The image ID.
	//
	// example:
	//
	// imgc-075cllfeuazh0****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// Android 12 system image
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// > This parameter is not yet available.
	//
	// example:
	//
	// null
	ImagePackageType *string `json:"ImagePackageType,omitempty" xml:"ImagePackageType,omitempty"`
	// The image type.
	//
	// This parameter is required.
	//
	// example:
	//
	// System
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The instance type.
	//
	// example:
	//
	// cpm.gx7.10xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum number of entries to return on each page for a paged query. Valid values: 1 to 100. Default value: 20.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that specifies the position from which to start the query. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kw9dGL5jves2FS9RLq****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The status of the image.
	//
	// example:
	//
	// AVAILABLE
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemType *string `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
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

func (s *DescribeImageListRequest) GetSystemType() *string {
	return s.SystemType
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

func (s *DescribeImageListRequest) SetSystemType(v string) *DescribeImageListRequest {
	s.SystemType = &v
	return s
}

func (s *DescribeImageListRequest) Validate() error {
	if s.ImageBizTags != nil {
		for _, item := range s.ImageBizTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImageListRequestImageBizTags struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
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
