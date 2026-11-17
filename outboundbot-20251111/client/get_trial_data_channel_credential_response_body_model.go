// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrialDataChannelCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTrialDataChannelCredentialResponseBody
	GetCode() *string
	SetData(v *GetTrialDataChannelCredentialResponseBodyData) *GetTrialDataChannelCredentialResponseBody
	GetData() *GetTrialDataChannelCredentialResponseBodyData
	SetHttpStatusCode(v int32) *GetTrialDataChannelCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTrialDataChannelCredentialResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetTrialDataChannelCredentialResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetTrialDataChannelCredentialResponseBody
	GetRequestId() *string
}

type GetTrialDataChannelCredentialResponseBody struct {
	// The business processing status code. OK indicates success. When an exception occurs, the corresponding business error code is returned. The error types and mapping rules follow the current API configuration. No error codes are added or modified in this operation.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data channel connection and authentication information for the specified device. Use the actual return values and do not deduce topics or credentials manually.
	//
	// example:
	//
	// {"DeviceId":"10001","Topic":"example-topic-10001","ClientId":"example-client-10001","Endpoint":"example.invalid","UserName":"example-user","Password":"example-password-not-a-credential","ExpirationTime":1794564000000}
	Data *GetTrialDataChannelCredentialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code recorded in the business response body. A normal return value is 200. To determine whether the business operation is successful, you must also check Code. You cannot determine the success of the call solely based on the presence of this field.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The business response or error description. It is usually empty or not returned upon success, and no fixed text is returned upon success. Upon failure, it provides an error description to assist in troubleshooting. Combine it with Code and Params to understand the cause of the error. The example is an error message when the trial link visitor token expires (corresponding to the Code Expired.VisitorToken) and is for format reference only.
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
	// The ID used to correlate and troubleshoot this POP request. The value in the response corresponds to this request. System request parameters are injected by the gateway, and the caller does not need to fill in system fields.
	//
	// example:
	//
	// AAAAAAAA-BBBB-4CCC-8DDD-EEEEEEEEEEEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTrialDataChannelCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTrialDataChannelCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrialDataChannelCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTrialDataChannelCredentialResponseBody) GetData() *GetTrialDataChannelCredentialResponseBodyData {
	return s.Data
}

func (s *GetTrialDataChannelCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTrialDataChannelCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTrialDataChannelCredentialResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetTrialDataChannelCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTrialDataChannelCredentialResponseBody) SetCode(v string) *GetTrialDataChannelCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *GetTrialDataChannelCredentialResponseBody) SetData(v *GetTrialDataChannelCredentialResponseBodyData) *GetTrialDataChannelCredentialResponseBody {
	s.Data = v
	return s
}

func (s *GetTrialDataChannelCredentialResponseBody) SetHttpStatusCode(v int32) *GetTrialDataChannelCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTrialDataChannelCredentialResponseBody) SetMessage(v string) *GetTrialDataChannelCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *GetTrialDataChannelCredentialResponseBody) SetParams(v []*string) *GetTrialDataChannelCredentialResponseBody {
	s.Params = v
	return s
}

func (s *GetTrialDataChannelCredentialResponseBody) SetRequestId(v string) *GetTrialDataChannelCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrialDataChannelCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTrialDataChannelCredentialResponseBodyData struct {
	// The client connection ID returned by the data channel service. It must be used together with the address, username, and password returned in this response. The example is a desensitized placeholder value and does not declare a fixed naming rule.
	//
	// example:
	//
	// example-client-10001
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the client device that initiates web voice calls or receives data channel events. The same client must use the same DeviceId when retrieving data channel credentials and initiating calls. The example is only a business device ID and not a real credential assigned by the platform.
	//
	// example:
	//
	// 10001
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The access endpoint of the data channel. The actual address, protocol, port, and authentication method are subject to the return values of the RTC gateway and client conventions. Local code only delegates to the gateway service and cannot be used to confirm a fixed production domain name or MQTT/WebSocket protocol.
	//
	// example:
	//
	// example.invalid
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The expiration time of the temporary credentials for the data channel. The current POP maps to expirationTime, while the local RTC second-party package field is expiredTime. The field mapping compatibility and time unit must be verified against the actual gateway version. Do not assume that the field is always returned or interpret it as milliseconds without authorization. The example only demonstrates the integer format.
	//
	// example:
	//
	// 1794564000000
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The temporary authentication password used to connect to the data channel. It must be used together with UserName and within the validity period. It is a sensitive credential and the real value must not be disclosed. The example is a placeholder value that cannot be used for authentication.
	//
	// example:
	//
	// example-password-not-a-credential
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The topic ID returned by the data channel service. The client subscribes to or receives events based on the actual protocol and this topic. The client must not deduce the topic based on DeviceId or treat the example as a real available topic.
	//
	// example:
	//
	// example-topic-10001
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The temporary authentication username used to connect to the data channel. It must be used together with the Password and Endpoint returned in the same response. The example placeholder value cannot be used for logon.
	//
	// example:
	//
	// example-user
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetTrialDataChannelCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTrialDataChannelCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTrialDataChannelCredentialResponseBodyData) GetClientId() *string {
	return s.ClientId
}

func (s *GetTrialDataChannelCredentialResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetTrialDataChannelCredentialResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetTrialDataChannelCredentialResponseBodyData) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *GetTrialDataChannelCredentialResponseBodyData) GetPassword() *string {
	return s.Password
}

func (s *GetTrialDataChannelCredentialResponseBodyData) GetTopic() *string {
	return s.Topic
}

func (s *GetTrialDataChannelCredentialResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *GetTrialDataChannelCredentialResponseBodyData) SetClientId(v string) *GetTrialDataChannelCredentialResponseBodyData {
	s.ClientId = &v
	return s
}

func (s *GetTrialDataChannelCredentialResponseBodyData) SetDeviceId(v string) *GetTrialDataChannelCredentialResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *GetTrialDataChannelCredentialResponseBodyData) SetEndpoint(v string) *GetTrialDataChannelCredentialResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetTrialDataChannelCredentialResponseBodyData) SetExpirationTime(v int64) *GetTrialDataChannelCredentialResponseBodyData {
	s.ExpirationTime = &v
	return s
}

func (s *GetTrialDataChannelCredentialResponseBodyData) SetPassword(v string) *GetTrialDataChannelCredentialResponseBodyData {
	s.Password = &v
	return s
}

func (s *GetTrialDataChannelCredentialResponseBodyData) SetTopic(v string) *GetTrialDataChannelCredentialResponseBodyData {
	s.Topic = &v
	return s
}

func (s *GetTrialDataChannelCredentialResponseBodyData) SetUserName(v string) *GetTrialDataChannelCredentialResponseBodyData {
	s.UserName = &v
	return s
}

func (s *GetTrialDataChannelCredentialResponseBodyData) Validate() error {
	return dara.Validate(s)
}
