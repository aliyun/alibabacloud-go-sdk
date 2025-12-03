// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProxyByTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *GetProxyByTypeResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetProxyByTypeResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetProxyByTypeResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetProxyByTypeResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *GetProxyByTypeResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetProxyByTypeResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *GetProxyByTypeResponseBody
	GetHttpStatusCode() *int32
	SetModule(v string) *GetProxyByTypeResponseBody
	GetModule() *string
	SetRequestId(v string) *GetProxyByTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetProxyByTypeResponseBody
	GetSuccess() *bool
}

type GetProxyByTypeResponseBody struct {
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
	// The specified parameter is invalid.
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
	// 118.113.245.10:3128
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// CEC1731F-0DA9-5E7D-AE53-7E4D76201C48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetProxyByTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProxyByTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetProxyByTypeResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetProxyByTypeResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetProxyByTypeResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetProxyByTypeResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetProxyByTypeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetProxyByTypeResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetProxyByTypeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetProxyByTypeResponseBody) GetModule() *string {
	return s.Module
}

func (s *GetProxyByTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProxyByTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetProxyByTypeResponseBody) SetAllowRetry(v bool) *GetProxyByTypeResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetAppName(v string) *GetProxyByTypeResponseBody {
	s.AppName = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetDynamicCode(v string) *GetProxyByTypeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetDynamicMessage(v string) *GetProxyByTypeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetErrorCode(v string) *GetProxyByTypeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetErrorMsg(v string) *GetProxyByTypeResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetHttpStatusCode(v int32) *GetProxyByTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetModule(v string) *GetProxyByTypeResponseBody {
	s.Module = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetRequestId(v string) *GetProxyByTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetSuccess(v bool) *GetProxyByTypeResponseBody {
	s.Success = &v
	return s
}

func (s *GetProxyByTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
