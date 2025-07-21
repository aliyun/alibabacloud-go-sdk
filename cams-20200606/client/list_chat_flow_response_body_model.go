// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListChatFlowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListChatFlowResponseBody
	GetCode() *string
	SetMessage(v string) *ListChatFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListChatFlowResponseBody
	GetRequestId() *string
	SetResponse(v map[string]interface{}) *ListChatFlowResponseBody
	GetResponse() map[string]interface{}
	SetSuccess(v bool) *ListChatFlowResponseBody
	GetSuccess() *bool
}

type ListChatFlowResponseBody struct {
	// Access denied details, this field is returned only when RAM verification fails.
	//
	// example:
	//
	// 无
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// System error code. For more details on error codes, please refer to the error code documentation.
	//
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error message.
	//
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Response data
	//
	// example:
	//
	// 无
	Response map[string]interface{} `json:"Response,omitempty" xml:"Response,omitempty"`
	// Whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListChatFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChatFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatFlowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListChatFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChatFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListChatFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChatFlowResponseBody) GetResponse() map[string]interface{} {
	return s.Response
}

func (s *ListChatFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListChatFlowResponseBody) SetAccessDeniedDetail(v string) *ListChatFlowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListChatFlowResponseBody) SetCode(v string) *ListChatFlowResponseBody {
	s.Code = &v
	return s
}

func (s *ListChatFlowResponseBody) SetMessage(v string) *ListChatFlowResponseBody {
	s.Message = &v
	return s
}

func (s *ListChatFlowResponseBody) SetRequestId(v string) *ListChatFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChatFlowResponseBody) SetResponse(v map[string]interface{}) *ListChatFlowResponseBody {
	s.Response = v
	return s
}

func (s *ListChatFlowResponseBody) SetSuccess(v bool) *ListChatFlowResponseBody {
	s.Success = &v
	return s
}

func (s *ListChatFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
