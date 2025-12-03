// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushRpaTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *PushRpaTaskDetailResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *PushRpaTaskDetailResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *PushRpaTaskDetailResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *PushRpaTaskDetailResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *PushRpaTaskDetailResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *PushRpaTaskDetailResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *PushRpaTaskDetailResponseBody
	GetHttpStatusCode() *int32
	SetModule(v string) *PushRpaTaskDetailResponseBody
	GetModule() *string
	SetRequestId(v string) *PushRpaTaskDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PushRpaTaskDetailResponseBody
	GetSuccess() *bool
}

type PushRpaTaskDetailResponseBody struct {
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
	// 207,155
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// 0320C9F4-5EDC-5355-9D7E-DF4CF6C2B3BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushRpaTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushRpaTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *PushRpaTaskDetailResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *PushRpaTaskDetailResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *PushRpaTaskDetailResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *PushRpaTaskDetailResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *PushRpaTaskDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PushRpaTaskDetailResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *PushRpaTaskDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PushRpaTaskDetailResponseBody) GetModule() *string {
	return s.Module
}

func (s *PushRpaTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushRpaTaskDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PushRpaTaskDetailResponseBody) SetAllowRetry(v bool) *PushRpaTaskDetailResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetAppName(v string) *PushRpaTaskDetailResponseBody {
	s.AppName = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetDynamicCode(v string) *PushRpaTaskDetailResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetDynamicMessage(v string) *PushRpaTaskDetailResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetErrorCode(v string) *PushRpaTaskDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetErrorMsg(v string) *PushRpaTaskDetailResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetHttpStatusCode(v int32) *PushRpaTaskDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetModule(v string) *PushRpaTaskDetailResponseBody {
	s.Module = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetRequestId(v string) *PushRpaTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetSuccess(v bool) *PushRpaTaskDetailResponseBody {
	s.Success = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
