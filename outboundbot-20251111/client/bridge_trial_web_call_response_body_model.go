// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBridgeTrialWebCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BridgeTrialWebCallResponseBody
	GetCode() *string
	SetData(v *BridgeTrialWebCallResponseBodyData) *BridgeTrialWebCallResponseBody
	GetData() *BridgeTrialWebCallResponseBodyData
	SetHttpStatusCode(v int32) *BridgeTrialWebCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BridgeTrialWebCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *BridgeTrialWebCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *BridgeTrialWebCallResponseBody
	GetRequestId() *string
}

type BridgeTrialWebCallResponseBody struct {
	// The business processing status code. OK indicates success. When an exception occurs, the corresponding business error code is returned. The error types and mapping rules follow the current API configuration, and no error codes are added or modified this time.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The session, channel, service address, and temporary authentication token for this web voice call. Use the actually returned connection parameters.
	//
	// example:
	//
	// {"InstanceId":"instance-example-001","SessionId":"ws-11111111-2222-4333-8444-555555555555","ChannelId":"ch-ws-10001-057100000000-1794560400000-ws-11111111-2222-4333-8444-555555555555","ExpirationTime":"1794564000000","ServerUrl":"wss://example.invalid/voice","Token":"example-web-call-token"}
	Data *BridgeTrialWebCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code recorded in the business response body. A normal return is 200. To determine whether the business is successful, you must also check Code. You cannot determine that the call is successful solely based on the existence of this field.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The business response or error description. It is usually empty or not returned upon success, and it is not required to return a fixed text upon success. Upon failure, it is an error description used to assist troubleshooting, and it is combined with Code and Params to understand the cause of the error. The example is an error message when the trial link visitor token expires (corresponding to the Code Expired.VisitorToken), and is for format reference only.
	//
	// example:
	//
	// VisitorToken: example-visitor-token is expired.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of parameters used for the localized display of error messages. It is not a scenario variable or business result. A normal successful response is usually an empty array. When an error occurs, parameters corresponding to the error template may be returned.
	//
	// example:
	//
	// []
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The ID used to associate and troubleshoot this POP request. The value in the response corresponds to this request. System request parameters are injected by the gateway, and the caller does not need to fill in the system fields.
	//
	// example:
	//
	// AAAAAAAA-BBBB-4CCC-8DDD-EEEEEEEEEEEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BridgeTrialWebCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BridgeTrialWebCallResponseBody) GoString() string {
	return s.String()
}

func (s *BridgeTrialWebCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *BridgeTrialWebCallResponseBody) GetData() *BridgeTrialWebCallResponseBodyData {
	return s.Data
}

func (s *BridgeTrialWebCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BridgeTrialWebCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BridgeTrialWebCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *BridgeTrialWebCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BridgeTrialWebCallResponseBody) SetCode(v string) *BridgeTrialWebCallResponseBody {
	s.Code = &v
	return s
}

func (s *BridgeTrialWebCallResponseBody) SetData(v *BridgeTrialWebCallResponseBodyData) *BridgeTrialWebCallResponseBody {
	s.Data = v
	return s
}

func (s *BridgeTrialWebCallResponseBody) SetHttpStatusCode(v int32) *BridgeTrialWebCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BridgeTrialWebCallResponseBody) SetMessage(v string) *BridgeTrialWebCallResponseBody {
	s.Message = &v
	return s
}

func (s *BridgeTrialWebCallResponseBody) SetParams(v []*string) *BridgeTrialWebCallResponseBody {
	s.Params = v
	return s
}

func (s *BridgeTrialWebCallResponseBody) SetRequestId(v string) *BridgeTrialWebCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *BridgeTrialWebCallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BridgeTrialWebCallResponseBodyData struct {
	// The web voice connection channel ID assigned or returned by the RTC gateway. It is used to associate and connect the voice channel for this call. It is different from the trial channel AccessChannelId. You must use the real value returned by this API.
	//
	// example:
	//
	// ch-ws-10001-057100000000-1794560400000-ws-11111111-2222-4333-8444-555555555555
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The expiration time of the web voice connection token for this call. The current POP type is string, and the backend WebCallContext type is long. The code does not mark the time unit. The format and unit must be confirmed according to the actual service provider agreement. The example is only a numeric string and does not represent a verified validity period unit.
	//
	// example:
	//
	// 1794564000000
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The instance identity of the associated product. You must use an instance that the invoker has access to and that actually exists. The instance must be consistent with the instance to which the scenario, channel, and session belong. The maximum length is subject to the current POP parameter definition.
	//
	// example:
	//
	// instance-example-001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The service address used to establish the web voice connection for this call. The format and protocol are determined by the service provider of the RTC gateway. You must use the actual returned value and cannot concatenate it yourself. The example is an unconnectable desensitized address and does not represent a production domain name or protocol guarantee.
	//
	// example:
	//
	// wss://example.invalid/voice
	ServerUrl *string `json:"ServerUrl,omitempty" xml:"ServerUrl,omitempty"`
	// The session ID used to uniquely associate this interaction. For a text session, it is the text- identifier returned when the session is started. For a web voice session, it is the ws- identifier. For a real number outbound call test, it is the test- identifier. Subsequent calls must use the actual returned value.
	//
	// example:
	//
	// ws-11111111-2222-4333-8444-555555555555
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The temporary token used to connect to the voice channel for this call. It must be used together with the ServerUrl and ChannelId returned this time. Do not disclose or record the actual token without desensitization. The example is not a valid credential, and the format and validity period are determined by the service provider.
	//
	// example:
	//
	// example-web-call-token
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s BridgeTrialWebCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BridgeTrialWebCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *BridgeTrialWebCallResponseBodyData) GetChannelId() *string {
	return s.ChannelId
}

func (s *BridgeTrialWebCallResponseBodyData) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *BridgeTrialWebCallResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BridgeTrialWebCallResponseBodyData) GetServerUrl() *string {
	return s.ServerUrl
}

func (s *BridgeTrialWebCallResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *BridgeTrialWebCallResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *BridgeTrialWebCallResponseBodyData) SetChannelId(v string) *BridgeTrialWebCallResponseBodyData {
	s.ChannelId = &v
	return s
}

func (s *BridgeTrialWebCallResponseBodyData) SetExpirationTime(v string) *BridgeTrialWebCallResponseBodyData {
	s.ExpirationTime = &v
	return s
}

func (s *BridgeTrialWebCallResponseBodyData) SetInstanceId(v string) *BridgeTrialWebCallResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *BridgeTrialWebCallResponseBodyData) SetServerUrl(v string) *BridgeTrialWebCallResponseBodyData {
	s.ServerUrl = &v
	return s
}

func (s *BridgeTrialWebCallResponseBodyData) SetSessionId(v string) *BridgeTrialWebCallResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *BridgeTrialWebCallResponseBodyData) SetToken(v string) *BridgeTrialWebCallResponseBodyData {
	s.Token = &v
	return s
}

func (s *BridgeTrialWebCallResponseBodyData) Validate() error {
	return dara.Validate(s)
}
