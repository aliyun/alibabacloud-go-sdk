// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateErResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateErResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *UpdateErResponseBody
	GetCode() *int32
	SetContent(v map[string]interface{}) *UpdateErResponseBody
	GetContent() map[string]interface{}
	SetMessage(v string) *UpdateErResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateErResponseBody
	GetRequestId() *string
}

type UpdateErResponseBody struct {
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
	// The returned data.
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
	// 3D9D6E7B-365B-5200-BFA6-9B79E269058C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateErResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateErResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateErResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateErResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateErResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *UpdateErResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateErResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateErResponseBody) SetAccessDeniedDetail(v string) *UpdateErResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateErResponseBody) SetCode(v int32) *UpdateErResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateErResponseBody) SetContent(v map[string]interface{}) *UpdateErResponseBody {
	s.Content = v
	return s
}

func (s *UpdateErResponseBody) SetMessage(v string) *UpdateErResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateErResponseBody) SetRequestId(v string) *UpdateErResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateErResponseBody) Validate() error {
	return dara.Validate(s)
}
