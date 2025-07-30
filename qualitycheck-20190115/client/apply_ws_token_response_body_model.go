// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyWsTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyWsTokenResponseBody
	GetCode() *string
	SetData(v *ApplyWsTokenResponseBodyData) *ApplyWsTokenResponseBody
	GetData() *ApplyWsTokenResponseBodyData
	SetHttpStatusCode(v int32) *ApplyWsTokenResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ApplyWsTokenResponseBody
	GetRequestId() *string
}

type ApplyWsTokenResponseBody struct {
	// example:
	//
	// 200
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ApplyWsTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyWsTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyWsTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyWsTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyWsTokenResponseBody) GetData() *ApplyWsTokenResponseBodyData {
	return s.Data
}

func (s *ApplyWsTokenResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ApplyWsTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyWsTokenResponseBody) SetCode(v string) *ApplyWsTokenResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyWsTokenResponseBody) SetData(v *ApplyWsTokenResponseBodyData) *ApplyWsTokenResponseBody {
	s.Data = v
	return s
}

func (s *ApplyWsTokenResponseBody) SetHttpStatusCode(v int32) *ApplyWsTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ApplyWsTokenResponseBody) SetRequestId(v string) *ApplyWsTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyWsTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type ApplyWsTokenResponseBodyData struct {
	// example:
	//
	// 42e59bcd-7206-44c5-ad34-525d364687c4
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 5607b21d1728700640
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// wss://ws-gateway-real-voice.aliyuncs.com
	WsEndpoint *string `json:"WsEndpoint,omitempty" xml:"WsEndpoint,omitempty"`
}

func (s ApplyWsTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ApplyWsTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *ApplyWsTokenResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *ApplyWsTokenResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *ApplyWsTokenResponseBodyData) GetWsEndpoint() *string {
	return s.WsEndpoint
}

func (s *ApplyWsTokenResponseBodyData) SetSessionId(v string) *ApplyWsTokenResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *ApplyWsTokenResponseBodyData) SetToken(v string) *ApplyWsTokenResponseBodyData {
	s.Token = &v
	return s
}

func (s *ApplyWsTokenResponseBodyData) SetWsEndpoint(v string) *ApplyWsTokenResponseBodyData {
	s.WsEndpoint = &v
	return s
}

func (s *ApplyWsTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
