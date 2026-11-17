// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBridgeWebCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BridgeWebCallResponseBody
	GetCode() *string
	SetData(v *BridgeWebCallResponseBodyData) *BridgeWebCallResponseBody
	GetData() *BridgeWebCallResponseBodyData
	SetHttpStatusCode(v int32) *BridgeWebCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BridgeWebCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *BridgeWebCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *BridgeWebCallResponseBody
	GetRequestId() *string
}

type BridgeWebCallResponseBody struct {
	// The business processing status code. A value of OK indicates success. If an exception occurs, the corresponding business error code is returned. The error types and mapping rules follow the current API configuration. No error codes are added or modified in this operation.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The session, channel, service address, and temporary authentication token for this web voice call. Use the actual returned connection parameters.
	//
	// example:
	//
	// {"InstanceId":"instance-example-001","SessionId":"ws-11111111-2222-4333-8444-555555555555","ChannelId":"ch-ws-10001-057100000000-1794560400000-ws-11111111-2222-4333-8444-555555555555","ExpirationTime":"1794564000000","ServerUrl":"wss://example.invalid/voice","Token":"example-web-call-token"}
	Data *BridgeWebCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code recorded in the business response body. A normal return value is 200. To determine whether the business operation is successful, you must also check Code. You cannot determine the success of the call solely based on the presence of this field.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The business response or error description. This field is usually empty or not returned upon success. A fixed text is not returned upon success. Upon failure, this field contains an error description to help troubleshoot the issue. Combine this field with Code and Params to understand the cause of the error. The example shows an error message when the instance does not exist, where Code is NotExists.InstanceId. This is for format reference only.
	//
	// example:
	//
	// Instance instance-example-001 does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of parameters used for the localized display of error messages. This is not a scenario variable or business result. A normal success response is usually an empty array. If an error occurs, parameters corresponding to the error template may be returned.
	//
	// example:
	//
	// []
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The ID used to associate and troubleshoot this POP request. The value in the response corresponds to the current request. System request parameters are injected by the gateway, and the caller does not need to fill in system fields.
	//
	// example:
	//
	// AAAAAAAA-BBBB-4CCC-8DDD-EEEEEEEEEEEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BridgeWebCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BridgeWebCallResponseBody) GoString() string {
	return s.String()
}

func (s *BridgeWebCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *BridgeWebCallResponseBody) GetData() *BridgeWebCallResponseBodyData {
	return s.Data
}

func (s *BridgeWebCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BridgeWebCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BridgeWebCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *BridgeWebCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BridgeWebCallResponseBody) SetCode(v string) *BridgeWebCallResponseBody {
	s.Code = &v
	return s
}

func (s *BridgeWebCallResponseBody) SetData(v *BridgeWebCallResponseBodyData) *BridgeWebCallResponseBody {
	s.Data = v
	return s
}

func (s *BridgeWebCallResponseBody) SetHttpStatusCode(v int32) *BridgeWebCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BridgeWebCallResponseBody) SetMessage(v string) *BridgeWebCallResponseBody {
	s.Message = &v
	return s
}

func (s *BridgeWebCallResponseBody) SetParams(v []*string) *BridgeWebCallResponseBody {
	s.Params = v
	return s
}

func (s *BridgeWebCallResponseBody) SetRequestId(v string) *BridgeWebCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *BridgeWebCallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BridgeWebCallResponseBodyData struct {
	// The web voice connection channel ID assigned or returned by the RTC gateway. This ID is used to associate and connect the voice channel for this call. This is different from the trial channel AccessChannelId. You must use the actual value returned by this operation.
	//
	// example:
	//
	// ch-ws-10001-057100000000-1794560400000-ws-11111111-2222-4333-8444-555555555555
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The expiration time of the web voice connection token for this call. The current POP type is string, and the backend WebCallContext type is long. The code does not specify the time unit. Confirm the format and unit based on the actual service provider agreement. The example is only a numeric string and does not represent a verified validity period unit.
	//
	// example:
	//
	// 1794564000000
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The instance identity of the associated product. You must use an existing instance that the invoker has permission to access. The instance must be consistent with the instance to which the scenario, channel, and session belong. The maximum length is subject to the current POP parameter definition.
	//
	// example:
	//
	// instance-example-001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The service address used to establish the web voice connection for this call. The format and protocol are determined by the service provider of the RTC gateway. You must use the actual returned value and cannot concatenate it yourself. The example is an unconnectable masked address and does not represent a production domain name or protocol guarantee.
	//
	// example:
	//
	// wss://example.invalid/voice
	ServerUrl *string `json:"ServerUrl,omitempty" xml:"ServerUrl,omitempty"`
	// The session ID used to uniquely associate this interaction. The ID starts with text- for text sessions returned when starting a session, ws- for web voice sessions, and test- for real number outbound call tests. Subsequent calls must use the actual returned value.
	//
	// example:
	//
	// ws-11111111-2222-4333-8444-555555555555
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The temporary token used to connect to the voice channel for this call. This token must be used together with the ServerUrl and ChannelId returned in this response. Do not expose or record actual tokens without masking them. The example is not a valid credential. The format and validity period are determined by the service provider.
	//
	// example:
	//
	// example-web-call-token
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s BridgeWebCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BridgeWebCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *BridgeWebCallResponseBodyData) GetChannelId() *string {
	return s.ChannelId
}

func (s *BridgeWebCallResponseBodyData) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *BridgeWebCallResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BridgeWebCallResponseBodyData) GetServerUrl() *string {
	return s.ServerUrl
}

func (s *BridgeWebCallResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *BridgeWebCallResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *BridgeWebCallResponseBodyData) SetChannelId(v string) *BridgeWebCallResponseBodyData {
	s.ChannelId = &v
	return s
}

func (s *BridgeWebCallResponseBodyData) SetExpirationTime(v string) *BridgeWebCallResponseBodyData {
	s.ExpirationTime = &v
	return s
}

func (s *BridgeWebCallResponseBodyData) SetInstanceId(v string) *BridgeWebCallResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *BridgeWebCallResponseBodyData) SetServerUrl(v string) *BridgeWebCallResponseBodyData {
	s.ServerUrl = &v
	return s
}

func (s *BridgeWebCallResponseBodyData) SetSessionId(v string) *BridgeWebCallResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *BridgeWebCallResponseBodyData) SetToken(v string) *BridgeWebCallResponseBodyData {
	s.Token = &v
	return s
}

func (s *BridgeWebCallResponseBodyData) Validate() error {
	return dara.Validate(s)
}
