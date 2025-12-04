// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateErAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateErAttachmentResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *UpdateErAttachmentResponseBody
	GetCode() *int32
	SetContent(v map[string]interface{}) *UpdateErAttachmentResponseBody
	GetContent() map[string]interface{}
	SetMessage(v string) *UpdateErAttachmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateErAttachmentResponseBody
	GetRequestId() *string
}

type UpdateErAttachmentResponseBody struct {
	// The details about the access denial.
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
	//
	// example:
	//
	// {}
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
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
	// 7F9082CC-3D94-560F-A575-8E8EF6CE2CB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateErAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateErAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateErAttachmentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateErAttachmentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateErAttachmentResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *UpdateErAttachmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateErAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateErAttachmentResponseBody) SetAccessDeniedDetail(v string) *UpdateErAttachmentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateErAttachmentResponseBody) SetCode(v int32) *UpdateErAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateErAttachmentResponseBody) SetContent(v map[string]interface{}) *UpdateErAttachmentResponseBody {
	s.Content = v
	return s
}

func (s *UpdateErAttachmentResponseBody) SetMessage(v string) *UpdateErAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateErAttachmentResponseBody) SetRequestId(v string) *UpdateErAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateErAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
