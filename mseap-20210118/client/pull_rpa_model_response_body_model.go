// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPullRpaModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *PullRpaModelResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *PullRpaModelResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *PullRpaModelResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *PullRpaModelResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *PullRpaModelResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *PullRpaModelResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *PullRpaModelResponseBody
	GetHttpStatusCode() *int32
	SetModule(v string) *PullRpaModelResponseBody
	GetModule() *string
	SetRequestId(v string) *PullRpaModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PullRpaModelResponseBody
	GetSuccess() *bool
}

type PullRpaModelResponseBody struct {
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
	// gatewayprood
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
	// 100008
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
	// {}
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// 06055768-6BC0-5FE7-BDFF-BD4D79537035
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PullRpaModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PullRpaModelResponseBody) GoString() string {
	return s.String()
}

func (s *PullRpaModelResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *PullRpaModelResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *PullRpaModelResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *PullRpaModelResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *PullRpaModelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PullRpaModelResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *PullRpaModelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PullRpaModelResponseBody) GetModule() *string {
	return s.Module
}

func (s *PullRpaModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PullRpaModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PullRpaModelResponseBody) SetAllowRetry(v bool) *PullRpaModelResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *PullRpaModelResponseBody) SetAppName(v string) *PullRpaModelResponseBody {
	s.AppName = &v
	return s
}

func (s *PullRpaModelResponseBody) SetDynamicCode(v string) *PullRpaModelResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *PullRpaModelResponseBody) SetDynamicMessage(v string) *PullRpaModelResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PullRpaModelResponseBody) SetErrorCode(v string) *PullRpaModelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PullRpaModelResponseBody) SetErrorMsg(v string) *PullRpaModelResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *PullRpaModelResponseBody) SetHttpStatusCode(v int32) *PullRpaModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PullRpaModelResponseBody) SetModule(v string) *PullRpaModelResponseBody {
	s.Module = &v
	return s
}

func (s *PullRpaModelResponseBody) SetRequestId(v string) *PullRpaModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *PullRpaModelResponseBody) SetSuccess(v bool) *PullRpaModelResponseBody {
	s.Success = &v
	return s
}

func (s *PullRpaModelResponseBody) Validate() error {
	return dara.Validate(s)
}
