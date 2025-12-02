// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesFromLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListImagesFromLibResponseBody
	GetCode() *int32
	SetCurrentPage(v int32) *ListImagesFromLibResponseBody
	GetCurrentPage() *int32
	SetHttpStatusCode(v int32) *ListImagesFromLibResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListImagesFromLibResponseBodyItems) *ListImagesFromLibResponseBody
	GetItems() []*ListImagesFromLibResponseBodyItems
	SetMsg(v string) *ListImagesFromLibResponseBody
	GetMsg() *string
	SetPageSize(v int32) *ListImagesFromLibResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListImagesFromLibResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListImagesFromLibResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListImagesFromLibResponseBody
	GetTotalCount() *int32
}

type ListImagesFromLibResponseBody struct {
	// Error code, consistent with HTTP status.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Current page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Data of the current page.
	Items []*ListImagesFromLibResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Further description of the error code.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// ID assigned by the backend to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success indicator.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total number of images.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListImagesFromLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImagesFromLibResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesFromLibResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListImagesFromLibResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListImagesFromLibResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListImagesFromLibResponseBody) GetItems() []*ListImagesFromLibResponseBodyItems {
	return s.Items
}

func (s *ListImagesFromLibResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *ListImagesFromLibResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImagesFromLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImagesFromLibResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListImagesFromLibResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListImagesFromLibResponseBody) SetCode(v int32) *ListImagesFromLibResponseBody {
	s.Code = &v
	return s
}

func (s *ListImagesFromLibResponseBody) SetCurrentPage(v int32) *ListImagesFromLibResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListImagesFromLibResponseBody) SetHttpStatusCode(v int32) *ListImagesFromLibResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListImagesFromLibResponseBody) SetItems(v []*ListImagesFromLibResponseBodyItems) *ListImagesFromLibResponseBody {
	s.Items = v
	return s
}

func (s *ListImagesFromLibResponseBody) SetMsg(v string) *ListImagesFromLibResponseBody {
	s.Msg = &v
	return s
}

func (s *ListImagesFromLibResponseBody) SetPageSize(v int32) *ListImagesFromLibResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListImagesFromLibResponseBody) SetRequestId(v string) *ListImagesFromLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesFromLibResponseBody) SetSuccess(v bool) *ListImagesFromLibResponseBody {
	s.Success = &v
	return s
}

func (s *ListImagesFromLibResponseBody) SetTotalCount(v int32) *ListImagesFromLibResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListImagesFromLibResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListImagesFromLibResponseBodyItems struct {
	// Creation time.
	//
	// example:
	//
	// 2022-11-30 16:30:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Image ID.
	//
	// example:
	//
	// 112
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Image URL.
	//
	// example:
	//
	// https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com/image/upload/IMG_2123.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Thumbnail URL.
	//
	// example:
	//
	// https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com/image/upload/IMG_2123.jpg
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
}

func (s ListImagesFromLibResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListImagesFromLibResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListImagesFromLibResponseBodyItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListImagesFromLibResponseBodyItems) GetImageId() *string {
	return s.ImageId
}

func (s *ListImagesFromLibResponseBodyItems) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ListImagesFromLibResponseBodyItems) GetThumbnailUrl() *string {
	return s.ThumbnailUrl
}

func (s *ListImagesFromLibResponseBodyItems) SetGmtCreate(v string) *ListImagesFromLibResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListImagesFromLibResponseBodyItems) SetImageId(v string) *ListImagesFromLibResponseBodyItems {
	s.ImageId = &v
	return s
}

func (s *ListImagesFromLibResponseBodyItems) SetImageUrl(v string) *ListImagesFromLibResponseBodyItems {
	s.ImageUrl = &v
	return s
}

func (s *ListImagesFromLibResponseBodyItems) SetThumbnailUrl(v string) *ListImagesFromLibResponseBodyItems {
	s.ThumbnailUrl = &v
	return s
}

func (s *ListImagesFromLibResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
