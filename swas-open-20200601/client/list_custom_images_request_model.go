// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListCustomImagesRequest
	GetClientToken() *string
	SetDataSnapshotId(v string) *ListCustomImagesRequest
	GetDataSnapshotId() *string
	SetImageIds(v string) *ListCustomImagesRequest
	GetImageIds() *string
	SetImageNames(v string) *ListCustomImagesRequest
	GetImageNames() *string
	SetInstanceId(v string) *ListCustomImagesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListCustomImagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomImagesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListCustomImagesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListCustomImagesRequest
	GetResourceGroupId() *string
	SetShare(v bool) *ListCustomImagesRequest
	GetShare() *bool
	SetSystemSnapshotId(v string) *ListCustomImagesRequest
	GetSystemSnapshotId() *string
	SetTag(v []*ListCustomImagesRequestTag) *ListCustomImagesRequest
	GetTag() []*ListCustomImagesRequestTag
}

type ListCustomImagesRequest struct {
	// The client token that is used to ensure the idempotence of the request. A client token is generated from your client and must be unique among different requests. **ClientToken*	- supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The data cloud disk snapshot ID.
	//
	// example:
	//
	// s-acscasca****
	DataSnapshotId *string `json:"DataSnapshotId,omitempty" xml:"DataSnapshotId,omitempty"`
	// The image IDs of the simple application servers. The value can be a JSON array that consists of up to 100 instance IDs. Separate multiple instance IDs with commas (,).
	//
	// example:
	//
	// ["fe9c66133a9d4688872869726b52****", "794c230fd3e64ea19f83f4d7a0ad****"]
	ImageIds *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	// The image names of the simple application servers. The value can be a JSON array that consists of up to 100 instance IDs. Separate multiple instance IDs with commas (,).
	//
	// example:
	//
	// ["test1****", "test2****"]
	ImageNames *string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty"`
	// The instance ID of the simple application server from which the image was created.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number of the snapshot list.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query.
	//
	// - Maximum value: 100.
	//
	// - Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application servers. You can call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to query the list of supported regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2bti7cf7yj2i
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to query shared images. Valid values:
	//
	// - False: The response does not include shared images.
	//
	// - True: Only shared images are returned.
	//
	// If this parameter is not specified, all images are returned by default.
	//
	// example:
	//
	// False
	Share *bool `json:"Share,omitempty" xml:"Share,omitempty"`
	// The system cloud disk snapshot ID.
	//
	// example:
	//
	// s-bp14m09pq8***0g6
	SystemSnapshotId *string `json:"SystemSnapshotId,omitempty" xml:"SystemSnapshotId,omitempty"`
	// The tags of the custom images.
	Tag []*ListCustomImagesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListCustomImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomImagesRequest) GoString() string {
	return s.String()
}

func (s *ListCustomImagesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListCustomImagesRequest) GetDataSnapshotId() *string {
	return s.DataSnapshotId
}

func (s *ListCustomImagesRequest) GetImageIds() *string {
	return s.ImageIds
}

func (s *ListCustomImagesRequest) GetImageNames() *string {
	return s.ImageNames
}

func (s *ListCustomImagesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCustomImagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomImagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomImagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCustomImagesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListCustomImagesRequest) GetShare() *bool {
	return s.Share
}

func (s *ListCustomImagesRequest) GetSystemSnapshotId() *string {
	return s.SystemSnapshotId
}

func (s *ListCustomImagesRequest) GetTag() []*ListCustomImagesRequestTag {
	return s.Tag
}

func (s *ListCustomImagesRequest) SetClientToken(v string) *ListCustomImagesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListCustomImagesRequest) SetDataSnapshotId(v string) *ListCustomImagesRequest {
	s.DataSnapshotId = &v
	return s
}

func (s *ListCustomImagesRequest) SetImageIds(v string) *ListCustomImagesRequest {
	s.ImageIds = &v
	return s
}

func (s *ListCustomImagesRequest) SetImageNames(v string) *ListCustomImagesRequest {
	s.ImageNames = &v
	return s
}

func (s *ListCustomImagesRequest) SetInstanceId(v string) *ListCustomImagesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCustomImagesRequest) SetPageNumber(v int32) *ListCustomImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomImagesRequest) SetPageSize(v int32) *ListCustomImagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomImagesRequest) SetRegionId(v string) *ListCustomImagesRequest {
	s.RegionId = &v
	return s
}

func (s *ListCustomImagesRequest) SetResourceGroupId(v string) *ListCustomImagesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListCustomImagesRequest) SetShare(v bool) *ListCustomImagesRequest {
	s.Share = &v
	return s
}

func (s *ListCustomImagesRequest) SetSystemSnapshotId(v string) *ListCustomImagesRequest {
	s.SystemSnapshotId = &v
	return s
}

func (s *ListCustomImagesRequest) SetTag(v []*ListCustomImagesRequestTag) *ListCustomImagesRequest {
	s.Tag = v
	return s
}

func (s *ListCustomImagesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomImagesRequestTag struct {
	// The tag key of the custom image. Valid values of the tag key length: 1 to 64. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the custom image. Valid values of the tag value length: 1 to 64. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListCustomImagesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListCustomImagesRequestTag) GoString() string {
	return s.String()
}

func (s *ListCustomImagesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListCustomImagesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListCustomImagesRequestTag) SetKey(v string) *ListCustomImagesRequestTag {
	s.Key = &v
	return s
}

func (s *ListCustomImagesRequestTag) SetValue(v string) *ListCustomImagesRequestTag {
	s.Value = &v
	return s
}

func (s *ListCustomImagesRequestTag) Validate() error {
	return dara.Validate(s)
}
