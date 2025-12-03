// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushRpaTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *PushRpaTaskResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *PushRpaTaskResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *PushRpaTaskResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *PushRpaTaskResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *PushRpaTaskResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *PushRpaTaskResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *PushRpaTaskResponseBody
	GetHttpStatusCode() *int32
	SetModule(v string) *PushRpaTaskResponseBody
	GetModule() *string
	SetRequestId(v string) *PushRpaTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PushRpaTaskResponseBody
	GetSuccess() *bool
}

type PushRpaTaskResponseBody struct {
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
	// itl-material
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
	// can not find env: lazada-sg-pre
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
	// 11111111111111111111111
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 58.23.71.83:3128
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// ECE5E7EF-6898-5E24-97A1-B96C73BDE26C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushRpaTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushRpaTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PushRpaTaskResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *PushRpaTaskResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *PushRpaTaskResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *PushRpaTaskResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *PushRpaTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PushRpaTaskResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *PushRpaTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PushRpaTaskResponseBody) GetModule() *string {
	return s.Module
}

func (s *PushRpaTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushRpaTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PushRpaTaskResponseBody) SetAllowRetry(v bool) *PushRpaTaskResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetAppName(v string) *PushRpaTaskResponseBody {
	s.AppName = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetDynamicCode(v string) *PushRpaTaskResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetDynamicMessage(v string) *PushRpaTaskResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetErrorCode(v string) *PushRpaTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetErrorMsg(v string) *PushRpaTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetHttpStatusCode(v int32) *PushRpaTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetModule(v string) *PushRpaTaskResponseBody {
	s.Module = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetRequestId(v string) *PushRpaTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetSuccess(v bool) *PushRpaTaskResponseBody {
	s.Success = &v
	return s
}

func (s *PushRpaTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
