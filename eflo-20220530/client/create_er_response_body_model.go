// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateErResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateErResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *CreateErResponseBody
	GetCode() *int32
	SetContent(v *CreateErResponseBodyContent) *CreateErResponseBody
	GetContent() *CreateErResponseBodyContent
	SetMessage(v string) *CreateErResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateErResponseBody
	GetRequestId() *string
}

type CreateErResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *CreateErResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateErResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateErResponseBody) GoString() string {
	return s.String()
}

func (s *CreateErResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateErResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateErResponseBody) GetContent() *CreateErResponseBodyContent {
	return s.Content
}

func (s *CreateErResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateErResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateErResponseBody) SetAccessDeniedDetail(v string) *CreateErResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateErResponseBody) SetCode(v int32) *CreateErResponseBody {
	s.Code = &v
	return s
}

func (s *CreateErResponseBody) SetContent(v *CreateErResponseBodyContent) *CreateErResponseBody {
	s.Content = v
	return s
}

func (s *CreateErResponseBody) SetMessage(v string) *CreateErResponseBody {
	s.Message = &v
	return s
}

func (s *CreateErResponseBody) SetRequestId(v string) *CreateErResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateErResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateErResponseBodyContent struct {
	// Lingjun HUB ID
	//
	// example:
	//
	// er-aueyxxsy
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
}

func (s CreateErResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s CreateErResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateErResponseBodyContent) GetErId() *string {
	return s.ErId
}

func (s *CreateErResponseBodyContent) SetErId(v string) *CreateErResponseBodyContent {
	s.ErId = &v
	return s
}

func (s *CreateErResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
