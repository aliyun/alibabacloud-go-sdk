// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallbackTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *CallbackTaskResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CallbackTaskResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CallbackTaskResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CallbackTaskResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *CallbackTaskResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CallbackTaskResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *CallbackTaskResponseBody
	GetHttpStatusCode() *int32
	SetModule(v bool) *CallbackTaskResponseBody
	GetModule() *bool
	SetRequestId(v string) *CallbackTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CallbackTaskResponseBody
	GetSuccess() *bool
}

type CallbackTaskResponseBody struct {
	// allowRetry
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// appName
	//
	// example:
	//
	// bohai-web-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamicCode
	//
	// example:
	//
	// 1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamicMessage
	//
	// example:
	//
	// can not find env: vpc-sg-pre
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// errorCode
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// errorMsg
	//
	// example:
	//
	// Success. Request Success.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// True
	Module *bool `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// 56B009EB-AAA5-52C9-B05F-AF425E3885E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CallbackTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CallbackTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CallbackTaskResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CallbackTaskResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CallbackTaskResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CallbackTaskResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CallbackTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CallbackTaskResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CallbackTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CallbackTaskResponseBody) GetModule() *bool {
	return s.Module
}

func (s *CallbackTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CallbackTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CallbackTaskResponseBody) SetAllowRetry(v bool) *CallbackTaskResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CallbackTaskResponseBody) SetAppName(v string) *CallbackTaskResponseBody {
	s.AppName = &v
	return s
}

func (s *CallbackTaskResponseBody) SetDynamicCode(v string) *CallbackTaskResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CallbackTaskResponseBody) SetDynamicMessage(v string) *CallbackTaskResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CallbackTaskResponseBody) SetErrorCode(v string) *CallbackTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CallbackTaskResponseBody) SetErrorMsg(v string) *CallbackTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CallbackTaskResponseBody) SetHttpStatusCode(v int32) *CallbackTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CallbackTaskResponseBody) SetModule(v bool) *CallbackTaskResponseBody {
	s.Module = &v
	return s
}

func (s *CallbackTaskResponseBody) SetRequestId(v string) *CallbackTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CallbackTaskResponseBody) SetSuccess(v bool) *CallbackTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CallbackTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
