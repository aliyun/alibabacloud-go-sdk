// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatappTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteChatappTemplateResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteChatappTemplateResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteChatappTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteChatappTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteChatappTemplateResponseBody
	GetSuccess() *bool
}

type DeleteChatappTemplateResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteChatappTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatappTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChatappTemplateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteChatappTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteChatappTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteChatappTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChatappTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteChatappTemplateResponseBody) SetAccessDeniedDetail(v string) *DeleteChatappTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteChatappTemplateResponseBody) SetCode(v string) *DeleteChatappTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChatappTemplateResponseBody) SetMessage(v string) *DeleteChatappTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteChatappTemplateResponseBody) SetRequestId(v string) *DeleteChatappTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChatappTemplateResponseBody) SetSuccess(v bool) *DeleteChatappTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteChatappTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
