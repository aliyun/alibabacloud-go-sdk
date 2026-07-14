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
	// The details about the access denial.
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
	// The returned data object.
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// true
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
