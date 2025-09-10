// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemindQuotaApplicationApprovalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *RemindQuotaApplicationApprovalResponseBody
	GetAllowRetry() *bool
	SetDynamicCode(v string) *RemindQuotaApplicationApprovalResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *RemindQuotaApplicationApprovalResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *RemindQuotaApplicationApprovalResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *RemindQuotaApplicationApprovalResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *RemindQuotaApplicationApprovalResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *RemindQuotaApplicationApprovalResponseBody
	GetHttpStatusCode() *int32
	SetModule(v string) *RemindQuotaApplicationApprovalResponseBody
	GetModule() *string
	SetRequestId(v string) *RemindQuotaApplicationApprovalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemindQuotaApplicationApprovalResponseBody
	GetSuccess() *bool
}

type RemindQuotaApplicationApprovalResponseBody struct {
	// Indicates whether retries are allowed. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// PARAMETER.ILLEGALL
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// Parameter[%s] is required.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The parameters whose values are invalid.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// RAM.PERMISSION.DENIED
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You are not authorized to do this action or the API input parameter is invalid.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The quota application ID.
	//
	// example:
	//
	// 219f1ff6-6205-495f-bda7-90d2fe945e****
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemindQuotaApplicationApprovalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemindQuotaApplicationApprovalResponseBody) GoString() string {
	return s.String()
}

func (s *RemindQuotaApplicationApprovalResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *RemindQuotaApplicationApprovalResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *RemindQuotaApplicationApprovalResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *RemindQuotaApplicationApprovalResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *RemindQuotaApplicationApprovalResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RemindQuotaApplicationApprovalResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *RemindQuotaApplicationApprovalResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemindQuotaApplicationApprovalResponseBody) GetModule() *string {
	return s.Module
}

func (s *RemindQuotaApplicationApprovalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemindQuotaApplicationApprovalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetAllowRetry(v bool) *RemindQuotaApplicationApprovalResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetDynamicCode(v string) *RemindQuotaApplicationApprovalResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetDynamicMessage(v string) *RemindQuotaApplicationApprovalResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetErrorArgs(v []interface{}) *RemindQuotaApplicationApprovalResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetErrorCode(v string) *RemindQuotaApplicationApprovalResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetErrorMsg(v string) *RemindQuotaApplicationApprovalResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetHttpStatusCode(v int32) *RemindQuotaApplicationApprovalResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetModule(v string) *RemindQuotaApplicationApprovalResponseBody {
	s.Module = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetRequestId(v string) *RemindQuotaApplicationApprovalResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) SetSuccess(v bool) *RemindQuotaApplicationApprovalResponseBody {
	s.Success = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponseBody) Validate() error {
	return dara.Validate(s)
}
