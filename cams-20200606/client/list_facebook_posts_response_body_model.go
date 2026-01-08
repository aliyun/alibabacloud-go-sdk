// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFacebookPostsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListFacebookPostsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListFacebookPostsResponseBody
	GetCode() *string
	SetData(v []*ListFacebookPostsResponseBodyData) *ListFacebookPostsResponseBody
	GetData() []*ListFacebookPostsResponseBodyData
	SetMessage(v string) *ListFacebookPostsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFacebookPostsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListFacebookPostsResponseBody
	GetSuccess() *bool
}

type ListFacebookPostsResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListFacebookPostsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFacebookPostsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFacebookPostsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFacebookPostsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListFacebookPostsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFacebookPostsResponseBody) GetData() []*ListFacebookPostsResponseBodyData {
	return s.Data
}

func (s *ListFacebookPostsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFacebookPostsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFacebookPostsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListFacebookPostsResponseBody) SetAccessDeniedDetail(v string) *ListFacebookPostsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListFacebookPostsResponseBody) SetCode(v string) *ListFacebookPostsResponseBody {
	s.Code = &v
	return s
}

func (s *ListFacebookPostsResponseBody) SetData(v []*ListFacebookPostsResponseBodyData) *ListFacebookPostsResponseBody {
	s.Data = v
	return s
}

func (s *ListFacebookPostsResponseBody) SetMessage(v string) *ListFacebookPostsResponseBody {
	s.Message = &v
	return s
}

func (s *ListFacebookPostsResponseBody) SetRequestId(v string) *ListFacebookPostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFacebookPostsResponseBody) SetSuccess(v bool) *ListFacebookPostsResponseBody {
	s.Success = &v
	return s
}

func (s *ListFacebookPostsResponseBody) Validate() error {
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

type ListFacebookPostsResponseBodyData struct {
	// example:
	//
	// 2025-04-08T10:44:48
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// N/A
	FullPicture *string `json:"FullPicture,omitempty" xml:"FullPicture,omitempty"`
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// https://xxxxxx07
	MediaUrl *string `json:"MediaUrl,omitempty" xml:"MediaUrl,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// https://xxxxxxx
	PermalinkUrl *string `json:"PermalinkUrl,omitempty" xml:"PermalinkUrl,omitempty"`
	// example:
	//
	// 453645465***
	RequestNo *string `json:"RequestNo,omitempty" xml:"RequestNo,omitempty"`
}

func (s ListFacebookPostsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFacebookPostsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFacebookPostsResponseBodyData) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListFacebookPostsResponseBodyData) GetFullPicture() *string {
	return s.FullPicture
}

func (s *ListFacebookPostsResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListFacebookPostsResponseBodyData) GetMediaType() *string {
	return s.MediaType
}

func (s *ListFacebookPostsResponseBodyData) GetMediaUrl() *string {
	return s.MediaUrl
}

func (s *ListFacebookPostsResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ListFacebookPostsResponseBodyData) GetPermalinkUrl() *string {
	return s.PermalinkUrl
}

func (s *ListFacebookPostsResponseBodyData) GetRequestNo() *string {
	return s.RequestNo
}

func (s *ListFacebookPostsResponseBodyData) SetCreatedTime(v string) *ListFacebookPostsResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *ListFacebookPostsResponseBodyData) SetFullPicture(v string) *ListFacebookPostsResponseBodyData {
	s.FullPicture = &v
	return s
}

func (s *ListFacebookPostsResponseBodyData) SetId(v string) *ListFacebookPostsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListFacebookPostsResponseBodyData) SetMediaType(v string) *ListFacebookPostsResponseBodyData {
	s.MediaType = &v
	return s
}

func (s *ListFacebookPostsResponseBodyData) SetMediaUrl(v string) *ListFacebookPostsResponseBodyData {
	s.MediaUrl = &v
	return s
}

func (s *ListFacebookPostsResponseBodyData) SetMessage(v string) *ListFacebookPostsResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListFacebookPostsResponseBodyData) SetPermalinkUrl(v string) *ListFacebookPostsResponseBodyData {
	s.PermalinkUrl = &v
	return s
}

func (s *ListFacebookPostsResponseBodyData) SetRequestNo(v string) *ListFacebookPostsResponseBodyData {
	s.RequestNo = &v
	return s
}

func (s *ListFacebookPostsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
