// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRedisValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *SetRedisValueResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *SetRedisValueResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *SetRedisValueResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *SetRedisValueResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *SetRedisValueResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *SetRedisValueResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *SetRedisValueResponseBody
	GetHttpStatusCode() *int32
	SetModule(v bool) *SetRedisValueResponseBody
	GetModule() *bool
	SetRequestId(v string) *SetRedisValueResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetRedisValueResponseBody
	GetSuccess() *bool
}

type SetRedisValueResponseBody struct {
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
	// cloudquery
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
	// 10.151.12.0/24,100.104.36.0/26,47.102.181.0/24,100.104.52.0/24,47.101.109.0/24,120.55.129.0/24,11.115.103.0/24,47.102.234.0/24
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// errorCode
	//
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// errorMsg
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
	// 71,135
	Module *bool `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// 195BABE2-7105-5C16-ABCE-2D0997CCE2E3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetRedisValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetRedisValueResponseBody) GoString() string {
	return s.String()
}

func (s *SetRedisValueResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *SetRedisValueResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *SetRedisValueResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *SetRedisValueResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *SetRedisValueResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SetRedisValueResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *SetRedisValueResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SetRedisValueResponseBody) GetModule() *bool {
	return s.Module
}

func (s *SetRedisValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetRedisValueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetRedisValueResponseBody) SetAllowRetry(v bool) *SetRedisValueResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *SetRedisValueResponseBody) SetAppName(v string) *SetRedisValueResponseBody {
	s.AppName = &v
	return s
}

func (s *SetRedisValueResponseBody) SetDynamicCode(v string) *SetRedisValueResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SetRedisValueResponseBody) SetDynamicMessage(v string) *SetRedisValueResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SetRedisValueResponseBody) SetErrorCode(v string) *SetRedisValueResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetRedisValueResponseBody) SetErrorMsg(v string) *SetRedisValueResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SetRedisValueResponseBody) SetHttpStatusCode(v int32) *SetRedisValueResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SetRedisValueResponseBody) SetModule(v bool) *SetRedisValueResponseBody {
	s.Module = &v
	return s
}

func (s *SetRedisValueResponseBody) SetRequestId(v string) *SetRedisValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetRedisValueResponseBody) SetSuccess(v bool) *SetRedisValueResponseBody {
	s.Success = &v
	return s
}

func (s *SetRedisValueResponseBody) Validate() error {
	return dara.Validate(s)
}
