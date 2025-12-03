// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRedisValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *GetRedisValueResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetRedisValueResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetRedisValueResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetRedisValueResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *GetRedisValueResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetRedisValueResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *GetRedisValueResponseBody
	GetHttpStatusCode() *int32
	SetModule(v string) *GetRedisValueResponseBody
	GetModule() *string
	SetRequestId(v string) *GetRedisValueResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRedisValueResponseBody
	GetSuccess() *bool
}

type GetRedisValueResponseBody struct {
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
	// can not find env: eleme-zb
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
	// zxdfghjklasdfghjkl
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
	// 107,72
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

func (s GetRedisValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRedisValueResponseBody) GoString() string {
	return s.String()
}

func (s *GetRedisValueResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetRedisValueResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetRedisValueResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetRedisValueResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetRedisValueResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetRedisValueResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetRedisValueResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetRedisValueResponseBody) GetModule() *string {
	return s.Module
}

func (s *GetRedisValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRedisValueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRedisValueResponseBody) SetAllowRetry(v bool) *GetRedisValueResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetRedisValueResponseBody) SetAppName(v string) *GetRedisValueResponseBody {
	s.AppName = &v
	return s
}

func (s *GetRedisValueResponseBody) SetDynamicCode(v string) *GetRedisValueResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetRedisValueResponseBody) SetDynamicMessage(v string) *GetRedisValueResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetRedisValueResponseBody) SetErrorCode(v string) *GetRedisValueResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRedisValueResponseBody) SetErrorMsg(v string) *GetRedisValueResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetRedisValueResponseBody) SetHttpStatusCode(v int32) *GetRedisValueResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRedisValueResponseBody) SetModule(v string) *GetRedisValueResponseBody {
	s.Module = &v
	return s
}

func (s *GetRedisValueResponseBody) SetRequestId(v string) *GetRedisValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRedisValueResponseBody) SetSuccess(v bool) *GetRedisValueResponseBody {
	s.Success = &v
	return s
}

func (s *GetRedisValueResponseBody) Validate() error {
	return dara.Validate(s)
}
