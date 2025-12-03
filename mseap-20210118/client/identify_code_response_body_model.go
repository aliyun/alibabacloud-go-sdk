// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdentifyCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *IdentifyCodeResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *IdentifyCodeResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *IdentifyCodeResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *IdentifyCodeResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *IdentifyCodeResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *IdentifyCodeResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *IdentifyCodeResponseBody
	GetHttpStatusCode() *int32
	SetModule(v string) *IdentifyCodeResponseBody
	GetModule() *string
	SetRequestId(v string) *IdentifyCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IdentifyCodeResponseBody
	GetSuccess() *bool
}

type IdentifyCodeResponseBody struct {
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
	// baasamlservice
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
	// 1234567890
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
	// 230,94
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// 8F30A673-46F0-5774-9D25-B45A29DB626E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s IdentifyCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IdentifyCodeResponseBody) GoString() string {
	return s.String()
}

func (s *IdentifyCodeResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *IdentifyCodeResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *IdentifyCodeResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *IdentifyCodeResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *IdentifyCodeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *IdentifyCodeResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *IdentifyCodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *IdentifyCodeResponseBody) GetModule() *string {
	return s.Module
}

func (s *IdentifyCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IdentifyCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IdentifyCodeResponseBody) SetAllowRetry(v bool) *IdentifyCodeResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetAppName(v string) *IdentifyCodeResponseBody {
	s.AppName = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetDynamicCode(v string) *IdentifyCodeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetDynamicMessage(v string) *IdentifyCodeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetErrorCode(v string) *IdentifyCodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetErrorMsg(v string) *IdentifyCodeResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetHttpStatusCode(v int32) *IdentifyCodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetModule(v string) *IdentifyCodeResponseBody {
	s.Module = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetRequestId(v string) *IdentifyCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetSuccess(v bool) *IdentifyCodeResponseBody {
	s.Success = &v
	return s
}

func (s *IdentifyCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
