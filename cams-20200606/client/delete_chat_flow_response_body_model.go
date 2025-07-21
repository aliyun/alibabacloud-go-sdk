// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteChatFlowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteChatFlowResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteChatFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteChatFlowResponseBody
	GetRequestId() *string
	SetResponse(v map[string]interface{}) *DeleteChatFlowResponseBody
	GetResponse() map[string]interface{}
	SetSuccess(v bool) *DeleteChatFlowResponseBody
	GetSuccess() *bool
}

type DeleteChatFlowResponseBody struct {
	// Access denied details, this field is returned only when RAM verification fails.
	//
	// example:
	//
	// 无
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Error code
	//
	// example:
	//
	// 示例值示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error message.
	//
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Unique request ID.
	//
	// example:
	//
	// 示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Response data
	//
	// example:
	//
	// 无
	Response map[string]interface{} `json:"Response,omitempty" xml:"Response,omitempty"`
	// Whether the call was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteChatFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChatFlowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteChatFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteChatFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteChatFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChatFlowResponseBody) GetResponse() map[string]interface{} {
	return s.Response
}

func (s *DeleteChatFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteChatFlowResponseBody) SetAccessDeniedDetail(v string) *DeleteChatFlowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteChatFlowResponseBody) SetCode(v string) *DeleteChatFlowResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChatFlowResponseBody) SetMessage(v string) *DeleteChatFlowResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteChatFlowResponseBody) SetRequestId(v string) *DeleteChatFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChatFlowResponseBody) SetResponse(v map[string]interface{}) *DeleteChatFlowResponseBody {
	s.Response = v
	return s
}

func (s *DeleteChatFlowResponseBody) SetSuccess(v bool) *DeleteChatFlowResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteChatFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
