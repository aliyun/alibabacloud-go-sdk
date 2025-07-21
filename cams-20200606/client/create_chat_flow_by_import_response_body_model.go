// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatFlowByImportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateChatFlowByImportResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateChatFlowByImportResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *CreateChatFlowByImportResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *CreateChatFlowByImportResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateChatFlowByImportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateChatFlowByImportResponseBody
	GetSuccess() *bool
}

type CreateChatFlowByImportResponseBody struct {
	// Access denied details.
	//
	// example:
	//
	// 无
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Request status code.
	//
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data object.
	//
	// example:
	//
	// 无
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// Error message.
	//
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the request was successful
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateChatFlowByImportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChatFlowByImportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatFlowByImportResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateChatFlowByImportResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateChatFlowByImportResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *CreateChatFlowByImportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateChatFlowByImportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChatFlowByImportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateChatFlowByImportResponseBody) SetAccessDeniedDetail(v string) *CreateChatFlowByImportResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateChatFlowByImportResponseBody) SetCode(v string) *CreateChatFlowByImportResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChatFlowByImportResponseBody) SetData(v map[string]interface{}) *CreateChatFlowByImportResponseBody {
	s.Data = v
	return s
}

func (s *CreateChatFlowByImportResponseBody) SetMessage(v string) *CreateChatFlowByImportResponseBody {
	s.Message = &v
	return s
}

func (s *CreateChatFlowByImportResponseBody) SetRequestId(v string) *CreateChatFlowByImportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChatFlowByImportResponseBody) SetSuccess(v bool) *CreateChatFlowByImportResponseBody {
	s.Success = &v
	return s
}

func (s *CreateChatFlowByImportResponseBody) Validate() error {
	return dara.Validate(s)
}
