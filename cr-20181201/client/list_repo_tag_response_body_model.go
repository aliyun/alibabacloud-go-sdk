// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRepoTagResponseBody
	GetCode() *string
	SetImages(v []*ListRepoTagResponseBodyImages) *ListRepoTagResponseBody
	GetImages() []*ListRepoTagResponseBodyImages
	SetIsSuccess(v bool) *ListRepoTagResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListRepoTagResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListRepoTagResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRepoTagResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListRepoTagResponseBody
	GetTotalCount() *string
}

type ListRepoTagResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The images.
	Images []*ListRepoTagResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 031572FA-7D8F-4C05-B790-1071E0E05DE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRepoTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepoTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepoTagResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRepoTagResponseBody) GetImages() []*ListRepoTagResponseBodyImages {
	return s.Images
}

func (s *ListRepoTagResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListRepoTagResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListRepoTagResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRepoTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepoTagResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListRepoTagResponseBody) SetCode(v string) *ListRepoTagResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepoTagResponseBody) SetImages(v []*ListRepoTagResponseBodyImages) *ListRepoTagResponseBody {
	s.Images = v
	return s
}

func (s *ListRepoTagResponseBody) SetIsSuccess(v bool) *ListRepoTagResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepoTagResponseBody) SetPageNo(v int32) *ListRepoTagResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListRepoTagResponseBody) SetPageSize(v int32) *ListRepoTagResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRepoTagResponseBody) SetRequestId(v string) *ListRepoTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepoTagResponseBody) SetTotalCount(v string) *ListRepoTagResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRepoTagResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRepoTagResponseBodyImages struct {
	// The digest of the image.
	//
	// example:
	//
	// 67bfbcc12b67936ec7f867927817cbb071832b873dbcaed312a1930ba5f1****
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The time when the image was created.
	//
	// example:
	//
	// 1572839125000
	ImageCreate *string `json:"ImageCreate,omitempty" xml:"ImageCreate,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// 45023655bf39c382e26a8607d057c27871dee163c1ecf48cc1ebf2a1****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The size of the image.
	//
	// example:
	//
	// 27107966
	ImageSize *int64 `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The time when the image was last updated.
	//
	// example:
	//
	// 1572875608000
	ImageUpdate *string `json:"ImageUpdate,omitempty" xml:"ImageUpdate,omitempty"`
	// The status of the image.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag of the image.
	//
	// example:
	//
	// v0.1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListRepoTagResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s ListRepoTagResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListRepoTagResponseBodyImages) GetDigest() *string {
	return s.Digest
}

func (s *ListRepoTagResponseBodyImages) GetImageCreate() *string {
	return s.ImageCreate
}

func (s *ListRepoTagResponseBodyImages) GetImageId() *string {
	return s.ImageId
}

func (s *ListRepoTagResponseBodyImages) GetImageSize() *int64 {
	return s.ImageSize
}

func (s *ListRepoTagResponseBodyImages) GetImageUpdate() *string {
	return s.ImageUpdate
}

func (s *ListRepoTagResponseBodyImages) GetStatus() *string {
	return s.Status
}

func (s *ListRepoTagResponseBodyImages) GetTag() *string {
	return s.Tag
}

func (s *ListRepoTagResponseBodyImages) SetDigest(v string) *ListRepoTagResponseBodyImages {
	s.Digest = &v
	return s
}

func (s *ListRepoTagResponseBodyImages) SetImageCreate(v string) *ListRepoTagResponseBodyImages {
	s.ImageCreate = &v
	return s
}

func (s *ListRepoTagResponseBodyImages) SetImageId(v string) *ListRepoTagResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListRepoTagResponseBodyImages) SetImageSize(v int64) *ListRepoTagResponseBodyImages {
	s.ImageSize = &v
	return s
}

func (s *ListRepoTagResponseBodyImages) SetImageUpdate(v string) *ListRepoTagResponseBodyImages {
	s.ImageUpdate = &v
	return s
}

func (s *ListRepoTagResponseBodyImages) SetStatus(v string) *ListRepoTagResponseBodyImages {
	s.Status = &v
	return s
}

func (s *ListRepoTagResponseBodyImages) SetTag(v string) *ListRepoTagResponseBodyImages {
	s.Tag = &v
	return s
}

func (s *ListRepoTagResponseBodyImages) Validate() error {
	return dara.Validate(s)
}
