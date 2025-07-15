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
	SetPageNumber(v int64) *ListImagesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListImagesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListImagesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListImagesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListImagesResponseBody
	GetTotalCount() *int32
}

type ListImagesResponseBody struct {
	Images []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListImagesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListImagesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListImagesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListImagesResponseBody) SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetPageNumber(v int64) *ListImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListImagesResponseBody) SetPageSize(v int64) *ListImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesResponseBody) SetSuccess(v bool) *ListImagesResponseBody {
	s.Success = &v
	return s
}

func (s *ListImagesResponseBody) SetTotalCount(v int32) *ListImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListImagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListImagesResponseBodyImages struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 2022-12-09T07:06:34Z
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocumentId  *int32  `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// m-bp181x855551ww5yq****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// VM
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// example:
	//
	// app-image
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OsTag      *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// v1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Weight  *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListImagesResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) GetAppId() *string {
	return s.AppId
}

func (s *ListImagesResponseBodyImages) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListImagesResponseBodyImages) GetDescription() *string {
	return s.Description
}

func (s *ListImagesResponseBodyImages) GetDocumentId() *int32 {
	return s.DocumentId
}

func (s *ListImagesResponseBodyImages) GetImageId() *string {
	return s.ImageId
}

func (s *ListImagesResponseBodyImages) GetImageType() *string {
	return s.ImageType
}

func (s *ListImagesResponseBodyImages) GetName() *string {
	return s.Name
}

func (s *ListImagesResponseBodyImages) GetOsTag() *string {
	return s.OsTag
}

func (s *ListImagesResponseBodyImages) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListImagesResponseBodyImages) GetVersion() *string {
	return s.Version
}

func (s *ListImagesResponseBodyImages) GetWeight() *int32 {
	return s.Weight
}

func (s *ListImagesResponseBodyImages) SetAppId(v string) *ListImagesResponseBodyImages {
	s.AppId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetCreateTime(v string) *ListImagesResponseBodyImages {
	s.CreateTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetDescription(v string) *ListImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetDocumentId(v int32) *ListImagesResponseBodyImages {
	s.DocumentId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageId(v string) *ListImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageType(v string) *ListImagesResponseBodyImages {
	s.ImageType = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetName(v string) *ListImagesResponseBodyImages {
	s.Name = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetOsTag(v string) *ListImagesResponseBodyImages {
	s.OsTag = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetUpdateTime(v string) *ListImagesResponseBodyImages {
	s.UpdateTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetVersion(v string) *ListImagesResponseBodyImages {
	s.Version = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetWeight(v int32) *ListImagesResponseBodyImages {
	s.Weight = &v
	return s
}

func (s *ListImagesResponseBodyImages) Validate() error {
	return dara.Validate(s)
}
