// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *GetVariableResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetVariableResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetVariableResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetVariableResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *GetVariableResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetVariableResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *GetVariableResponseBody
	GetHttpStatusCode() *int32
	SetModule(v string) *GetVariableResponseBody
	GetModule() *string
	SetRequestId(v string) *GetVariableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetVariableResponseBody
	GetSuccess() *bool
}

type GetVariableResponseBody struct {
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
	// voldemort-aliyun-com
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
	// can not find env: eleme-zb-pre
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
	// module
	//
	// example:
	//
	// 207,155
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// D28419C9-335E-50A7-BD7D-ACF250A825E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVariableResponseBody) GoString() string {
	return s.String()
}

func (s *GetVariableResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetVariableResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetVariableResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetVariableResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetVariableResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetVariableResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetVariableResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetVariableResponseBody) GetModule() *string {
	return s.Module
}

func (s *GetVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVariableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetVariableResponseBody) SetAllowRetry(v bool) *GetVariableResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetVariableResponseBody) SetAppName(v string) *GetVariableResponseBody {
	s.AppName = &v
	return s
}

func (s *GetVariableResponseBody) SetDynamicCode(v string) *GetVariableResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetVariableResponseBody) SetDynamicMessage(v string) *GetVariableResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetVariableResponseBody) SetErrorCode(v string) *GetVariableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetVariableResponseBody) SetErrorMsg(v string) *GetVariableResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetVariableResponseBody) SetHttpStatusCode(v int32) *GetVariableResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVariableResponseBody) SetModule(v string) *GetVariableResponseBody {
	s.Module = &v
	return s
}

func (s *GetVariableResponseBody) SetRequestId(v string) *GetVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVariableResponseBody) SetSuccess(v bool) *GetVariableResponseBody {
	s.Success = &v
	return s
}

func (s *GetVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
