// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWebshellTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWebshellTokenResponseBody
	GetCode() *string
	SetData(v *GetWebshellTokenResponseBodyData) *GetWebshellTokenResponseBody
	GetData() *GetWebshellTokenResponseBodyData
	SetErrorCode(v string) *GetWebshellTokenResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetWebshellTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWebshellTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWebshellTokenResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *GetWebshellTokenResponseBody
	GetTraceId() *string
}

type GetWebshellTokenResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The request was successful.
	//
	// 	- **3xx**: The request was redirected.
	//
	// 	- **4xx**: The request failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *GetWebshellTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the call failed. Value values:
	//
	// 	- If the request was successful, **ErrorCode*	- is not returned.
	//
	// 	- If the request failed, the **ErrorCode*	- parameter is returned. For more information, see **Error codes*	- section in this topic.
	//
	// example:
	//
	// 400
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message. Value description:
	//
	// 	- If the request was successful, a success message is returned.
	//
	// 	- An error code is returned if the request failed.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s GetWebshellTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWebshellTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebshellTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWebshellTokenResponseBody) GetData() *GetWebshellTokenResponseBodyData {
	return s.Data
}

func (s *GetWebshellTokenResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetWebshellTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWebshellTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWebshellTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWebshellTokenResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *GetWebshellTokenResponseBody) SetCode(v string) *GetWebshellTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetWebshellTokenResponseBody) SetData(v *GetWebshellTokenResponseBodyData) *GetWebshellTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetWebshellTokenResponseBody) SetErrorCode(v string) *GetWebshellTokenResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWebshellTokenResponseBody) SetMessage(v string) *GetWebshellTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetWebshellTokenResponseBody) SetRequestId(v string) *GetWebshellTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWebshellTokenResponseBody) SetSuccess(v bool) *GetWebshellTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GetWebshellTokenResponseBody) SetTraceId(v string) *GetWebshellTokenResponseBody {
	s.TraceId = &v
	return s
}

func (s *GetWebshellTokenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWebshellTokenResponseBodyData struct {
	// Webshell page address
	//
	// example:
	//
	// HttpUrl :  "https://saenext.console.aliyun.com/cn-shenzhen/app-list/app1/micro-app/shell/pod1?tokenId=xxx
	HttpUrl *string `json:"HttpUrl,omitempty" xml:"HttpUrl,omitempty"`
	// The information about the token.
	//
	// example:
	//
	// zWWpvRj_5pzof4hfo7-hGynM8oGMmO_7
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The WebSocket address.
	//
	// example:
	//
	// wss://sae-webshell.console.aliyun.com/websocket/eamWebshell?tokenId=xxx&region=cn-shenzhen
	WebSocketUrl *string `json:"WebSocketUrl,omitempty" xml:"WebSocketUrl,omitempty"`
}

func (s GetWebshellTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWebshellTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWebshellTokenResponseBodyData) GetHttpUrl() *string {
	return s.HttpUrl
}

func (s *GetWebshellTokenResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *GetWebshellTokenResponseBodyData) GetWebSocketUrl() *string {
	return s.WebSocketUrl
}

func (s *GetWebshellTokenResponseBodyData) SetHttpUrl(v string) *GetWebshellTokenResponseBodyData {
	s.HttpUrl = &v
	return s
}

func (s *GetWebshellTokenResponseBodyData) SetToken(v string) *GetWebshellTokenResponseBodyData {
	s.Token = &v
	return s
}

func (s *GetWebshellTokenResponseBodyData) SetWebSocketUrl(v string) *GetWebshellTokenResponseBodyData {
	s.WebSocketUrl = &v
	return s
}

func (s *GetWebshellTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
