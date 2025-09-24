// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody
	GetImages() []*ListImagesResponseBodyImages
	SetRequestId(v string) *ListImagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListImagesResponseBody
	GetTotalCount() *int64
}

type ListImagesResponseBody struct {
	// The images.
	Images []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned images.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) GetImages() []*ListImagesResponseBodyImages {
	return s.Images
}

func (s *ListImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImagesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListImagesResponseBody) SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesResponseBody) SetTotalCount(v int64) *ListImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListImagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListImagesResponseBodyImages struct {
	// The accessibility of the image. Valid values:
	//
	// 	- PUBLIC: All members can access the image.
	//
	// 	- PRIVATE: Only the creator can access the image.
	//
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The image description.
	//
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the image is created, in UTC. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the image is modified, in UTC. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The image ID.
	//
	// example:
	//
	// image-tzi7f9******s45t
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image address, which includes the version number.
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	// The image tags.
	Labels []*ListImagesResponseBodyImagesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The image name.
	//
	// example:
	//
	// tensorflow_2.9
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 15577******82932
	ParentUserId *string `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	// The image size. Unit: GB.
	//
	// example:
	//
	// 2
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// 镜像来源 ID
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// 镜像来源类型
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 15577******82932
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 20******55
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListImagesResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListImagesResponseBodyImages) GetDescription() *string {
	return s.Description
}

func (s *ListImagesResponseBodyImages) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListImagesResponseBodyImages) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListImagesResponseBodyImages) GetImageId() *string {
	return s.ImageId
}

func (s *ListImagesResponseBodyImages) GetImageUri() *string {
	return s.ImageUri
}

func (s *ListImagesResponseBodyImages) GetLabels() []*ListImagesResponseBodyImagesLabels {
	return s.Labels
}

func (s *ListImagesResponseBodyImages) GetName() *string {
	return s.Name
}

func (s *ListImagesResponseBodyImages) GetParentUserId() *string {
	return s.ParentUserId
}

func (s *ListImagesResponseBodyImages) GetSize() *int32 {
	return s.Size
}

func (s *ListImagesResponseBodyImages) GetSourceId() *string {
	return s.SourceId
}

func (s *ListImagesResponseBodyImages) GetSourceType() *string {
	return s.SourceType
}

func (s *ListImagesResponseBodyImages) GetUserId() *string {
	return s.UserId
}

func (s *ListImagesResponseBodyImages) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListImagesResponseBodyImages) SetAccessibility(v string) *ListImagesResponseBodyImages {
	s.Accessibility = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetDescription(v string) *ListImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetGmtCreateTime(v string) *ListImagesResponseBodyImages {
	s.GmtCreateTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetGmtModifiedTime(v string) *ListImagesResponseBodyImages {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageId(v string) *ListImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageUri(v string) *ListImagesResponseBodyImages {
	s.ImageUri = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetLabels(v []*ListImagesResponseBodyImagesLabels) *ListImagesResponseBodyImages {
	s.Labels = v
	return s
}

func (s *ListImagesResponseBodyImages) SetName(v string) *ListImagesResponseBodyImages {
	s.Name = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetParentUserId(v string) *ListImagesResponseBodyImages {
	s.ParentUserId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSize(v int32) *ListImagesResponseBodyImages {
	s.Size = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSourceId(v string) *ListImagesResponseBodyImages {
	s.SourceId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSourceType(v string) *ListImagesResponseBodyImages {
	s.SourceType = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetUserId(v string) *ListImagesResponseBodyImages {
	s.UserId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetWorkspaceId(v string) *ListImagesResponseBodyImages {
	s.WorkspaceId = &v
	return s
}

func (s *ListImagesResponseBodyImages) Validate() error {
	return dara.Validate(s)
}

type ListImagesResponseBodyImagesLabels struct {
	// The tag key.
	//
	// example:
	//
	// system.chipType
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// GPU
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListImagesResponseBodyImagesLabels) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBodyImagesLabels) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesLabels) GetKey() *string {
	return s.Key
}

func (s *ListImagesResponseBodyImagesLabels) GetValue() *string {
	return s.Value
}

func (s *ListImagesResponseBodyImagesLabels) SetKey(v string) *ListImagesResponseBodyImagesLabels {
	s.Key = &v
	return s
}

func (s *ListImagesResponseBodyImagesLabels) SetValue(v string) *ListImagesResponseBodyImagesLabels {
	s.Value = &v
	return s
}

func (s *ListImagesResponseBodyImagesLabels) Validate() error {
	return dara.Validate(s)
}
