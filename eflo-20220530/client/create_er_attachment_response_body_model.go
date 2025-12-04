// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateErAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateErAttachmentResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *CreateErAttachmentResponseBody
	GetCode() *int32
	SetContent(v *CreateErAttachmentResponseBodyContent) *CreateErAttachmentResponseBody
	GetContent() *CreateErAttachmentResponseBodyContent
	SetMessage(v string) *CreateErAttachmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateErAttachmentResponseBody
	GetRequestId() *string
}

type CreateErAttachmentResponseBody struct {
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
	// The response data.
	Content *CreateErAttachmentResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is displayed.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DBAD15D6-3F47-5B36-8A92-57C2919D13D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateErAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateErAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateErAttachmentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateErAttachmentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateErAttachmentResponseBody) GetContent() *CreateErAttachmentResponseBodyContent {
	return s.Content
}

func (s *CreateErAttachmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateErAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateErAttachmentResponseBody) SetAccessDeniedDetail(v string) *CreateErAttachmentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateErAttachmentResponseBody) SetCode(v int32) *CreateErAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateErAttachmentResponseBody) SetContent(v *CreateErAttachmentResponseBodyContent) *CreateErAttachmentResponseBody {
	s.Content = v
	return s
}

func (s *CreateErAttachmentResponseBody) SetMessage(v string) *CreateErAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateErAttachmentResponseBody) SetRequestId(v string) *CreateErAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateErAttachmentResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateErAttachmentResponseBodyContent struct {
	// The ID of the network connection instance.
	//
	// example:
	//
	// er-attachment-ggjbfhqv
	ErAttachmentId *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
}

func (s CreateErAttachmentResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s CreateErAttachmentResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateErAttachmentResponseBodyContent) GetErAttachmentId() *string {
	return s.ErAttachmentId
}

func (s *CreateErAttachmentResponseBodyContent) SetErAttachmentId(v string) *CreateErAttachmentResponseBodyContent {
	s.ErAttachmentId = &v
	return s
}

func (s *CreateErAttachmentResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
