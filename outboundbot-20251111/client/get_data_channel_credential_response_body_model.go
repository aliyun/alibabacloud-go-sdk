// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataChannelCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataChannelCredentialResponseBody
	GetCode() *string
	SetData(v *GetDataChannelCredentialResponseBodyData) *GetDataChannelCredentialResponseBody
	GetData() *GetDataChannelCredentialResponseBodyData
	SetHttpStatusCode(v int32) *GetDataChannelCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataChannelCredentialResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetDataChannelCredentialResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetDataChannelCredentialResponseBody
	GetRequestId() *string
}

type GetDataChannelCredentialResponseBody struct {
	// The business processing status code. OK indicates success. When an exception occurs, the corresponding business error code is returned. Error types and mapping rules follow the current API configuration, and no error codes are added or modified in this operation.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data channel connection and authentication information for the specified device. Use the actual return values and do not infer topics or credentials.
	//
	// example:
	//
	// {"DeviceId":"10001","Topic":"example-topic-10001","ClientId":"example-client-10001","Endpoint":"example.invalid","UserName":"example-user","Password":"example-password-not-a-credential","ExpirationTime":1794564000000}
	Data *GetDataChannelCredentialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code recorded in the business response body. A normal return value is 200. To determine whether the business operation is successful, you must also check Code. You cannot determine the success of the call solely based on the presence of this field.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The business response or error description. It is usually empty or not returned upon success, and you should not expect a fixed text to be returned upon success. Upon failure, it provides an error description to assist in troubleshooting, and should be used in conjunction with Code and Params to understand the cause of the error. The example shows an error message when the instance does not exist (corresponding to the Code NotExists.InstanceId) and is provided for format reference only.
	//
	// example:
	//
	// Instance instance-example-001 does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of parameters used for the localized display of error messages, which are not scenario variables or business results. A normal successful response is usually an empty array. When an error occurs, parameters corresponding to the error template may be returned.
	//
	// example:
	//
	// []
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The ID used to associate and troubleshoot this POP request. The value in the response corresponds to this request. System request parameters are injected by the gateway, and the caller does not need to fill in system fields.
	//
	// example:
	//
	// AAAAAAAA-BBBB-4CCC-8DDD-EEEEEEEEEEEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataChannelCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataChannelCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataChannelCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataChannelCredentialResponseBody) GetData() *GetDataChannelCredentialResponseBodyData {
	return s.Data
}

func (s *GetDataChannelCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataChannelCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataChannelCredentialResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetDataChannelCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataChannelCredentialResponseBody) SetCode(v string) *GetDataChannelCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataChannelCredentialResponseBody) SetData(v *GetDataChannelCredentialResponseBodyData) *GetDataChannelCredentialResponseBody {
	s.Data = v
	return s
}

func (s *GetDataChannelCredentialResponseBody) SetHttpStatusCode(v int32) *GetDataChannelCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataChannelCredentialResponseBody) SetMessage(v string) *GetDataChannelCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataChannelCredentialResponseBody) SetParams(v []*string) *GetDataChannelCredentialResponseBody {
	s.Params = v
	return s
}

func (s *GetDataChannelCredentialResponseBody) SetRequestId(v string) *GetDataChannelCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataChannelCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataChannelCredentialResponseBodyData struct {
	// The client connection identifier returned by the data channel service, which should be used in conjunction with the address, username, and password returned in this response. The example is a desensitized placeholder value and does not declare a fixed naming rule.
	//
	// example:
	//
	// example-client-10001
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The device ID corresponding to the current data channel credentials, which should match the requested DeviceId. Use it in conjunction with the Topic, ClientId, and authentication information returned in the same response, and do not mix it with the ChannelId of the voice channel.
	//
	// example:
	//
	// 10001
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The access endpoint for the data channel. The actual address, protocol, port, and authentication method are subject to the RTC gateway return and client conventions. Local code only delegates to the gateway service, and you cannot use this to confirm a fixed production domain name or MQTT/WebSocket protocol.
	//
	// example:
	//
	// example.invalid
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The expiration time of the temporary credentials for the data channel. The current POP maps this to expirationTime, whereas the field in the local RTC SDK is expiredTime. You must verify the actual gateway version for field mapping compatibility and time units. Do not assume that this field is always returned or arbitrarily interpret it as milliseconds. The example only demonstrates the integer format.
	//
	// example:
	//
	// 1794564000000
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The temporary authentication password used to connect to the data channel, which should be used in conjunction with UserName and within its validity period. It is a sensitive credential, and the real value must not be disclosed. The example is a placeholder value that cannot be used for authentication.
	//
	// example:
	//
	// example-password-not-a-credential
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The topic identifier returned by the data channel service. The client subscribes to or receives events based on the actual protocol and this topic. Do not infer the topic based on DeviceId, and do not treat the example as a real usable topic.
	//
	// example:
	//
	// example-topic-10001
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The temporary authentication username used to connect to the data channel. It should be used in conjunction with the Password and Endpoint returned in the same response. Do not use the example placeholder value for logon.
	//
	// example:
	//
	// example-user
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetDataChannelCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataChannelCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataChannelCredentialResponseBodyData) GetClientId() *string {
	return s.ClientId
}

func (s *GetDataChannelCredentialResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetDataChannelCredentialResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetDataChannelCredentialResponseBodyData) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *GetDataChannelCredentialResponseBodyData) GetPassword() *string {
	return s.Password
}

func (s *GetDataChannelCredentialResponseBodyData) GetTopic() *string {
	return s.Topic
}

func (s *GetDataChannelCredentialResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *GetDataChannelCredentialResponseBodyData) SetClientId(v string) *GetDataChannelCredentialResponseBodyData {
	s.ClientId = &v
	return s
}

func (s *GetDataChannelCredentialResponseBodyData) SetDeviceId(v string) *GetDataChannelCredentialResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *GetDataChannelCredentialResponseBodyData) SetEndpoint(v string) *GetDataChannelCredentialResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetDataChannelCredentialResponseBodyData) SetExpirationTime(v int64) *GetDataChannelCredentialResponseBodyData {
	s.ExpirationTime = &v
	return s
}

func (s *GetDataChannelCredentialResponseBodyData) SetPassword(v string) *GetDataChannelCredentialResponseBodyData {
	s.Password = &v
	return s
}

func (s *GetDataChannelCredentialResponseBodyData) SetTopic(v string) *GetDataChannelCredentialResponseBodyData {
	s.Topic = &v
	return s
}

func (s *GetDataChannelCredentialResponseBodyData) SetUserName(v string) *GetDataChannelCredentialResponseBodyData {
	s.UserName = &v
	return s
}

func (s *GetDataChannelCredentialResponseBodyData) Validate() error {
	return dara.Validate(s)
}
