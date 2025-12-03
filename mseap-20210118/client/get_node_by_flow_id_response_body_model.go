// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeByFlowIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *GetNodeByFlowIdResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetNodeByFlowIdResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetNodeByFlowIdResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetNodeByFlowIdResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *GetNodeByFlowIdResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetNodeByFlowIdResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *GetNodeByFlowIdResponseBody
	GetHttpStatusCode() *int32
	SetModule(v string) *GetNodeByFlowIdResponseBody
	GetModule() *string
	SetRequestId(v string) *GetNodeByFlowIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNodeByFlowIdResponseBody
	GetSuccess() *bool
}

type GetNodeByFlowIdResponseBody struct {
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
	// 200,131
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

func (s GetNodeByFlowIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeByFlowIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeByFlowIdResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetNodeByFlowIdResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetNodeByFlowIdResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetNodeByFlowIdResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetNodeByFlowIdResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetNodeByFlowIdResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetNodeByFlowIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetNodeByFlowIdResponseBody) GetModule() *string {
	return s.Module
}

func (s *GetNodeByFlowIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeByFlowIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNodeByFlowIdResponseBody) SetAllowRetry(v bool) *GetNodeByFlowIdResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetAppName(v string) *GetNodeByFlowIdResponseBody {
	s.AppName = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetDynamicCode(v string) *GetNodeByFlowIdResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetDynamicMessage(v string) *GetNodeByFlowIdResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetErrorCode(v string) *GetNodeByFlowIdResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetErrorMsg(v string) *GetNodeByFlowIdResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetHttpStatusCode(v int32) *GetNodeByFlowIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetModule(v string) *GetNodeByFlowIdResponseBody {
	s.Module = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetRequestId(v string) *GetNodeByFlowIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetSuccess(v bool) *GetNodeByFlowIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) Validate() error {
	return dara.Validate(s)
}
