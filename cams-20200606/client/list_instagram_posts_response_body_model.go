// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstagramPostsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListInstagramPostsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListInstagramPostsResponseBody
	GetCode() *string
	SetData(v []*ListInstagramPostsResponseBodyData) *ListInstagramPostsResponseBody
	GetData() []*ListInstagramPostsResponseBodyData
	SetMessage(v string) *ListInstagramPostsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstagramPostsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstagramPostsResponseBody
	GetSuccess() *bool
}

type ListInstagramPostsResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListInstagramPostsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A***66411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstagramPostsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstagramPostsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstagramPostsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListInstagramPostsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstagramPostsResponseBody) GetData() []*ListInstagramPostsResponseBodyData {
	return s.Data
}

func (s *ListInstagramPostsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstagramPostsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstagramPostsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstagramPostsResponseBody) SetAccessDeniedDetail(v string) *ListInstagramPostsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListInstagramPostsResponseBody) SetCode(v string) *ListInstagramPostsResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstagramPostsResponseBody) SetData(v []*ListInstagramPostsResponseBodyData) *ListInstagramPostsResponseBody {
	s.Data = v
	return s
}

func (s *ListInstagramPostsResponseBody) SetMessage(v string) *ListInstagramPostsResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstagramPostsResponseBody) SetRequestId(v string) *ListInstagramPostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstagramPostsResponseBody) SetSuccess(v bool) *ListInstagramPostsResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstagramPostsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstagramPostsResponseBodyData struct {
	// example:
	//
	// 2024-09-30T15:53:01+08:00
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// https://xxxxxxxx
	FullPicture *string `json:"FullPicture,omitempty" xml:"FullPicture,omitempty"`
	// example:
	//
	// 111
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// image
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// https://qjck.oss-cn-shanghai.aliyuncs.com/xcx/2025-08-29/0bfa1c8d-d35f-2dd5-e082-7b00***ee9.jpg?x-oss-process=image/resize,m_lfit,w_1200
	MediaUrl *string `json:"MediaUrl,omitempty" xml:"MediaUrl,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// https://xxxxxxxx
	PermalinkUrl *string `json:"PermalinkUrl,omitempty" xml:"PermalinkUrl,omitempty"`
}

func (s ListInstagramPostsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstagramPostsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstagramPostsResponseBodyData) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListInstagramPostsResponseBodyData) GetFullPicture() *string {
	return s.FullPicture
}

func (s *ListInstagramPostsResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListInstagramPostsResponseBodyData) GetMediaType() *string {
	return s.MediaType
}

func (s *ListInstagramPostsResponseBodyData) GetMediaUrl() *string {
	return s.MediaUrl
}

func (s *ListInstagramPostsResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ListInstagramPostsResponseBodyData) GetPermalinkUrl() *string {
	return s.PermalinkUrl
}

func (s *ListInstagramPostsResponseBodyData) SetCreatedTime(v string) *ListInstagramPostsResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *ListInstagramPostsResponseBodyData) SetFullPicture(v string) *ListInstagramPostsResponseBodyData {
	s.FullPicture = &v
	return s
}

func (s *ListInstagramPostsResponseBodyData) SetId(v string) *ListInstagramPostsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListInstagramPostsResponseBodyData) SetMediaType(v string) *ListInstagramPostsResponseBodyData {
	s.MediaType = &v
	return s
}

func (s *ListInstagramPostsResponseBodyData) SetMediaUrl(v string) *ListInstagramPostsResponseBodyData {
	s.MediaUrl = &v
	return s
}

func (s *ListInstagramPostsResponseBodyData) SetMessage(v string) *ListInstagramPostsResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListInstagramPostsResponseBodyData) SetPermalinkUrl(v string) *ListInstagramPostsResponseBodyData {
	s.PermalinkUrl = &v
	return s
}

func (s *ListInstagramPostsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
