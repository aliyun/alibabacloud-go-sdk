// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateChatFlowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateChatFlowResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *CreateChatFlowResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *CreateChatFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateChatFlowResponseBody
	GetRequestId() *string
	SetResponse(v map[string]interface{}) *CreateChatFlowResponseBody
	GetResponse() map[string]interface{}
	SetSuccess(v bool) *CreateChatFlowResponseBody
	GetSuccess() *bool
}

type CreateChatFlowResponseBody struct {
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
	// 示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data object.
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// Error message.
	//
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Unique request ID.
	//
	// example:
	//
	// 示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Response data
	Response map[string]interface{} `json:"Response,omitempty" xml:"Response,omitempty"`
	// Whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateChatFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChatFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatFlowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateChatFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateChatFlowResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *CreateChatFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateChatFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChatFlowResponseBody) GetResponse() map[string]interface{} {
	return s.Response
}

func (s *CreateChatFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateChatFlowResponseBody) SetAccessDeniedDetail(v string) *CreateChatFlowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateChatFlowResponseBody) SetCode(v string) *CreateChatFlowResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChatFlowResponseBody) SetData(v map[string]interface{}) *CreateChatFlowResponseBody {
	s.Data = v
	return s
}

func (s *CreateChatFlowResponseBody) SetMessage(v string) *CreateChatFlowResponseBody {
	s.Message = &v
	return s
}

func (s *CreateChatFlowResponseBody) SetRequestId(v string) *CreateChatFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChatFlowResponseBody) SetResponse(v map[string]interface{}) *CreateChatFlowResponseBody {
	s.Response = v
	return s
}

func (s *CreateChatFlowResponseBody) SetSuccess(v bool) *CreateChatFlowResponseBody {
	s.Success = &v
	return s
}

func (s *CreateChatFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
