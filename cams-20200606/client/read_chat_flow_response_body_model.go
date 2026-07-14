// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadChatFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ReadChatFlowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ReadChatFlowResponseBody
	GetCode() *string
	SetMessage(v string) *ReadChatFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadChatFlowResponseBody
	GetRequestId() *string
	SetResponse(v map[string]interface{}) *ReadChatFlowResponseBody
	GetResponse() map[string]interface{}
	SetSuccess(v bool) *ReadChatFlowResponseBody
	GetSuccess() *bool
}

type ReadChatFlowResponseBody struct {
	// The detailed reason why the access was denied.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The error code. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned data.
	Response map[string]interface{} `json:"Response,omitempty" xml:"Response,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// - true: The operation is successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadChatFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadChatFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ReadChatFlowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ReadChatFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadChatFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadChatFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadChatFlowResponseBody) GetResponse() map[string]interface{} {
	return s.Response
}

func (s *ReadChatFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadChatFlowResponseBody) SetAccessDeniedDetail(v string) *ReadChatFlowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ReadChatFlowResponseBody) SetCode(v string) *ReadChatFlowResponseBody {
	s.Code = &v
	return s
}

func (s *ReadChatFlowResponseBody) SetMessage(v string) *ReadChatFlowResponseBody {
	s.Message = &v
	return s
}

func (s *ReadChatFlowResponseBody) SetRequestId(v string) *ReadChatFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadChatFlowResponseBody) SetResponse(v map[string]interface{}) *ReadChatFlowResponseBody {
	s.Response = v
	return s
}

func (s *ReadChatFlowResponseBody) SetSuccess(v bool) *ReadChatFlowResponseBody {
	s.Success = &v
	return s
}

func (s *ReadChatFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
