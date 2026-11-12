// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMqttTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMqttTicketResponseBody
	GetCode() *string
	SetExpireTimeStamp(v string) *GetMqttTicketResponseBody
	GetExpireTimeStamp() *string
	SetMessage(v string) *GetMqttTicketResponseBody
	GetMessage() *string
	SetMqttEndpoint(v string) *GetMqttTicketResponseBody
	GetMqttEndpoint() *string
	SetMqttGroupId(v string) *GetMqttTicketResponseBody
	GetMqttGroupId() *string
	SetMqttInstanceId(v string) *GetMqttTicketResponseBody
	GetMqttInstanceId() *string
	SetMqttPassword(v string) *GetMqttTicketResponseBody
	GetMqttPassword() *string
	SetMqttUsername(v string) *GetMqttTicketResponseBody
	GetMqttUsername() *string
	SetMqttVPCEndpoint(v string) *GetMqttTicketResponseBody
	GetMqttVPCEndpoint() *string
	SetRequestId(v string) *GetMqttTicketResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetMqttTicketResponseBody
	GetSuccess() *string
	SetTopicList(v string) *GetMqttTicketResponseBody
	GetTopicList() *string
}

type GetMqttTicketResponseBody struct {
	// The error code returned when the request fails. For example, Client.Forbidden indicates that the client is forbidden from establishing a connection.
	//
	// example:
	//
	// Client.Forbidden
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The expiration time of the MQTT connection password. It is a UNIX timestamp in milliseconds and is returned as a string.
	//
	// example:
	//
	// 1794278400000
	ExpireTimeStamp *string `json:"ExpireTimeStamp,omitempty" xml:"ExpireTimeStamp,omitempty"`
	// The error message returned when the request fails, which corresponds to the error code.
	//
	// example:
	//
	// Client forbidden connect
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The access endpoint for the MQTT client. Use the endpoint returned in this response to establish a connection.
	//
	// example:
	//
	// mqtt.example.com
	MqttEndpoint *string `json:"MqttEndpoint,omitempty" xml:"MqttEndpoint,omitempty"`
	// The MQTT group ID, used to concatenate the MQTT ClientId.
	//
	// example:
	//
	// GID_EXAMPLE
	MqttGroupId *string `json:"MqttGroupId,omitempty" xml:"MqttGroupId,omitempty"`
	// The MQTT instance ID, used for token authentication.
	//
	// example:
	//
	// post-cn-7mz2****
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	// The password used to establish an MQTT connection. It consists of publish and subscribe tokens in the format of W|PublishToken|R|SubscribeToken. Use the complete password returned in the response. The desensitized examples in the documentation cannot be used for connections.
	//
	// example:
	//
	// W|pubTokenExample****|R|subTokenExample****
	MqttPassword *string `json:"MqttPassword,omitempty" xml:"MqttPassword,omitempty"`
	// The username used to establish an MQTT connection. The format is Token|AccessKeyId|InstanceId. AccessKeyId and InstanceId are identifiers used by the server to generate credentials. Use the complete username returned in the response.
	//
	// example:
	//
	// Token|LTAI5tExample****|post-cn-7mz2****
	MqttUsername *string `json:"MqttUsername,omitempty" xml:"MqttUsername,omitempty"`
	// The VPC access endpoint for MQTT. It is used by clients that meet the VPC network access requirements.
	//
	// example:
	//
	// mqtt-vpc.example.com
	MqttVPCEndpoint *string `json:"MqttVPCEndpoint,omitempty" xml:"MqttVPCEndpoint,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B9E3F7B2-5B13-4A7B-8C36-6D1B0A8F0000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The topic list. Each element in the list contains topic (the topic name), type (the topic type), and permission (the permission). The list returned by this operation includes topics for publishing, tenant subscription, device subscription, and system subscription. The publish permission is W, and the subscribe permission is R.
	//
	// example:
	//
	// [{"topic":"pub/550e8400-e29b-41d4-a716-446655440000","type":"pub","permission":"W"},{"topic":"sub/ten/123456789012****","type":"sub-ten","permission":"R"},{"topic":"sub/uuid/550e8400-e29b-41d4-a716-446655440000","type":"sub-uuid","permission":"R"},{"topic":"sub/sys","type":"sub-sys","permission":"R"}]
	TopicList *string `json:"TopicList,omitempty" xml:"TopicList,omitempty"`
}

func (s GetMqttTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMqttTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetMqttTicketResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMqttTicketResponseBody) GetExpireTimeStamp() *string {
	return s.ExpireTimeStamp
}

func (s *GetMqttTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMqttTicketResponseBody) GetMqttEndpoint() *string {
	return s.MqttEndpoint
}

func (s *GetMqttTicketResponseBody) GetMqttGroupId() *string {
	return s.MqttGroupId
}

func (s *GetMqttTicketResponseBody) GetMqttInstanceId() *string {
	return s.MqttInstanceId
}

func (s *GetMqttTicketResponseBody) GetMqttPassword() *string {
	return s.MqttPassword
}

func (s *GetMqttTicketResponseBody) GetMqttUsername() *string {
	return s.MqttUsername
}

func (s *GetMqttTicketResponseBody) GetMqttVPCEndpoint() *string {
	return s.MqttVPCEndpoint
}

func (s *GetMqttTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMqttTicketResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetMqttTicketResponseBody) GetTopicList() *string {
	return s.TopicList
}

func (s *GetMqttTicketResponseBody) SetCode(v string) *GetMqttTicketResponseBody {
	s.Code = &v
	return s
}

func (s *GetMqttTicketResponseBody) SetExpireTimeStamp(v string) *GetMqttTicketResponseBody {
	s.ExpireTimeStamp = &v
	return s
}

func (s *GetMqttTicketResponseBody) SetMessage(v string) *GetMqttTicketResponseBody {
	s.Message = &v
	return s
}

func (s *GetMqttTicketResponseBody) SetMqttEndpoint(v string) *GetMqttTicketResponseBody {
	s.MqttEndpoint = &v
	return s
}

func (s *GetMqttTicketResponseBody) SetMqttGroupId(v string) *GetMqttTicketResponseBody {
	s.MqttGroupId = &v
	return s
}

func (s *GetMqttTicketResponseBody) SetMqttInstanceId(v string) *GetMqttTicketResponseBody {
	s.MqttInstanceId = &v
	return s
}

func (s *GetMqttTicketResponseBody) SetMqttPassword(v string) *GetMqttTicketResponseBody {
	s.MqttPassword = &v
	return s
}

func (s *GetMqttTicketResponseBody) SetMqttUsername(v string) *GetMqttTicketResponseBody {
	s.MqttUsername = &v
	return s
}

func (s *GetMqttTicketResponseBody) SetMqttVPCEndpoint(v string) *GetMqttTicketResponseBody {
	s.MqttVPCEndpoint = &v
	return s
}

func (s *GetMqttTicketResponseBody) SetRequestId(v string) *GetMqttTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMqttTicketResponseBody) SetSuccess(v string) *GetMqttTicketResponseBody {
	s.Success = &v
	return s
}

func (s *GetMqttTicketResponseBody) SetTopicList(v string) *GetMqttTicketResponseBody {
	s.TopicList = &v
	return s
}

func (s *GetMqttTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
