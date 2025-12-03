// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeByTemplateIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *GetNodeByTemplateIdResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetNodeByTemplateIdResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetNodeByTemplateIdResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetNodeByTemplateIdResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *GetNodeByTemplateIdResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetNodeByTemplateIdResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *GetNodeByTemplateIdResponseBody
	GetHttpStatusCode() *int32
	SetModule(v string) *GetNodeByTemplateIdResponseBody
	GetModule() *string
	SetRequestId(v string) *GetNodeByTemplateIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNodeByTemplateIdResponseBody
	GetSuccess() *bool
}

type GetNodeByTemplateIdResponseBody struct {
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
	// qdhxgcagent01
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
	// 220,116
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// 53D045B1-466F-5165-B3BB-42E36F02BA86
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNodeByTemplateIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeByTemplateIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeByTemplateIdResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetNodeByTemplateIdResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetNodeByTemplateIdResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetNodeByTemplateIdResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetNodeByTemplateIdResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetNodeByTemplateIdResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetNodeByTemplateIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetNodeByTemplateIdResponseBody) GetModule() *string {
	return s.Module
}

func (s *GetNodeByTemplateIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeByTemplateIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNodeByTemplateIdResponseBody) SetAllowRetry(v bool) *GetNodeByTemplateIdResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetAppName(v string) *GetNodeByTemplateIdResponseBody {
	s.AppName = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetDynamicCode(v string) *GetNodeByTemplateIdResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetDynamicMessage(v string) *GetNodeByTemplateIdResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetErrorCode(v string) *GetNodeByTemplateIdResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetErrorMsg(v string) *GetNodeByTemplateIdResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetHttpStatusCode(v int32) *GetNodeByTemplateIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetModule(v string) *GetNodeByTemplateIdResponseBody {
	s.Module = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetRequestId(v string) *GetNodeByTemplateIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetSuccess(v bool) *GetNodeByTemplateIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) Validate() error {
	return dara.Validate(s)
}
