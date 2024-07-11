// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DataValue struct {
	// The Alibaba Cloud account ID or Resource Access Management (RAM) user to which the AccessKey pair that is used to create the static username and password belongs.
	//
	// example:
	//
	// 1565*******973901
	MasterUid *int64 `json:"masterUid,omitempty" xml:"masterUid,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// amqp-cn-uqm******03
	CInstanceId *string `json:"cInstanceId,omitempty" xml:"cInstanceId,omitempty"`
	// The AccessKey ID that is used to create the static username and password.
	//
	// example:
	//
	// LTAI5***********eRZtEJ6vfo
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// The static username.
	//
	// example:
	//
	// MjphbXFwLWNuLXVxbTJ6cjc2djAwMzpMVEFJNX*******ZNMWVSWnRFSjZ2Zm8=
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// The static password.
	//
	// example:
	//
	// OUYwQzM2QjZBRkUxNDRFM***************MzZCNzdDQzoxNjcxNDMwMzkyODI1
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The timestamp that indicates when the static username and password were deleted. Unit: milliseconds.
	//
	// example:
	//
	// 1671175303522
	Deleted *int64 `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// The timestamp that indicates when the static username and password were created. Unit: milliseconds.
	//
	// example:
	//
	// 1671175303522
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
}

func (s DataValue) String() string {
	return tea.Prettify(s)
}

func (s DataValue) GoString() string {
	return s.String()
}

func (s *DataValue) SetMasterUid(v int64) *DataValue {
	s.MasterUid = &v
	return s
}

func (s *DataValue) SetCInstanceId(v string) *DataValue {
	s.CInstanceId = &v
	return s
}

func (s *DataValue) SetAccessKey(v string) *DataValue {
	s.AccessKey = &v
	return s
}

func (s *DataValue) SetUserName(v string) *DataValue {
	s.UserName = &v
	return s
}

func (s *DataValue) SetPassword(v string) *DataValue {
	s.Password = &v
	return s
}

func (s *DataValue) SetDeleted(v int64) *DataValue {
	s.Deleted = &v
	return s
}

func (s *DataValue) SetCreateTimestamp(v int64) *DataValue {
	s.CreateTimestamp = &v
	return s
}

type CreateAccountRequest struct {
	// The AccessKey ID of your Alibaba Cloud account or RAM user. For information about how to obtain an AccessKey pair, see [Create an AccessKey pair](https://help.aliyun.com/document_detail/116401.html).
	//
	// >  If you use the pair of static username and password that is created by using the Accesskey pair of a RAM user to access ApsaraMQ for RabbitMQ to send and receive messages, make sure that the RAM user is granted the required permissions. For more information, see [RAM policies](https://help.aliyun.com/document_detail/146559.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAI5t8be*******tEJ6vfo
	AccountAccessKey *string `json:"accountAccessKey,omitempty" xml:"accountAccessKey,omitempty"`
	// The timestamp that indicates when the password is created. Unit: milliseconds.
	//
	// >  This timestamp is specified by you and is used to generate a static password. The timestamp is not the timestamp that indicates when the system generates the password.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1671175303522
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// The ID of the instance for which you want to create a pair of static username and password.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-*********
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The AccessKey secret signature. The system generates a static password based on the signature in the request, the AccessKey secret signature, and the username.
	//
	// The system uses the HMAC-SHA1 algorithm to generate the AccessKey secret signature based on the timestamp that indicates when the username is created and the AccessKey ID. For more information, see the **"Sample code on how to generate a signature"*	- section of this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4c1a6367ce4c4255e9617326f9133ac6359533f6
	SecretSign *string `json:"secretSign,omitempty" xml:"secretSign,omitempty"`
	// The signature. The system generates a static password based on the signature in the request, the AccessKey secret signature, and the username.
	//
	// The system uses the HMAC-SHA1 algorithm to generate the signature based on the timestamp that indicates when the username is created and the AccessKey ID. For more information, see the **"Sample code on how to generate a signature"*	- section of this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22c2d7d1769cb53c5a6d9213248e2de524c4f799
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// The static username that you want to create.
	//
	// The value of this parameter is a Base64-encoded string that is generated based on the instance ID and AccessKey ID. For more information, see the "**Sample code on how to generate a username**" section of this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// MjphbXFwLWNuLXp***********************Q4YmVNbVZNMWVSWnRFSjZ2Zm8=
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) SetAccountAccessKey(v string) *CreateAccountRequest {
	s.AccountAccessKey = &v
	return s
}

func (s *CreateAccountRequest) SetCreateTimestamp(v int64) *CreateAccountRequest {
	s.CreateTimestamp = &v
	return s
}

func (s *CreateAccountRequest) SetInstanceId(v string) *CreateAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAccountRequest) SetSecretSign(v string) *CreateAccountRequest {
	s.SecretSign = &v
	return s
}

func (s *CreateAccountRequest) SetSignature(v string) *CreateAccountRequest {
	s.Signature = &v
	return s
}

func (s *CreateAccountRequest) SetUserName(v string) *CreateAccountRequest {
	s.UserName = &v
	return s
}

type CreateAccountResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 92385FD2-624A-48C9-8FB5-753F2AFA***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) SetCode(v int32) *CreateAccountResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAccountResponseBody) SetData(v *CreateAccountResponseBodyData) *CreateAccountResponseBody {
	s.Data = v
	return s
}

func (s *CreateAccountResponseBody) SetMessage(v string) *CreateAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccountResponseBody) SetSuccess(v bool) *CreateAccountResponseBody {
	s.Success = &v
	return s
}

type CreateAccountResponseBodyData struct {
	// The AccessKey ID that is used to create the password.
	//
	// example:
	//
	// LTAI5***********eRZtEJ6vfo
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The timestamp that indicates when the password was created. Unit: milliseconds.
	//
	// example:
	//
	// 1671175303522
	CreateTimeStamp *int64 `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// amqp-cn-*********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The Alibaba Cloud account ID or RAM user to which the AccessKey pair that is used to create the static username and password belongs.
	//
	// example:
	//
	// 15657*********01
	MasterUId *int64 `json:"MasterUId,omitempty" xml:"MasterUId,omitempty"`
	// The created static password.
	//
	// example:
	//
	// NEMxQTYzNjdDRTVDNDI1NUU5NjE3**************1MzNGODoxNjcxMTc1MzEzNTIy
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The created static username.
	//
	// example:
	//
	// MjphbXFwLWNuLXVxbTJ6cjc2djAwMzpMVEFJNX*******ZNMWVSWnRFSjZ2Zm8=
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateAccountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBodyData) SetAccessKey(v string) *CreateAccountResponseBodyData {
	s.AccessKey = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetCreateTimeStamp(v int64) *CreateAccountResponseBodyData {
	s.CreateTimeStamp = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetInstanceId(v string) *CreateAccountResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetMasterUId(v int64) *CreateAccountResponseBodyData {
	s.MasterUId = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetPassword(v string) *CreateAccountResponseBodyData {
	s.Password = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetUserName(v string) *CreateAccountResponseBodyData {
	s.UserName = &v
	return s
}

type CreateAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateAccountResponse) SetHeaders(v map[string]*string) *CreateAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateAccountResponse) SetStatusCode(v int32) *CreateAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccountResponse) SetBody(v *CreateAccountResponseBody) *CreateAccountResponse {
	s.Body = v
	return s
}

type CreateBindingRequest struct {
	// The key-value pairs that are configured for the headers attributes of a message. One or more key-value pairs can be concatenated to configure the headers attributes of a message. You must specify the x-match attribute as one of the valid values. You can specify custom values for other attributes. Valid values of the x-match attribute:
	//
	// 	- \\*\\*all: \\*\\*A headers exchange routes a message to a queue only if all binding attributes of the queue except for x-match match the headers attributes of the message. This value is the default value.
	//
	// 	- \\*\\*any: \\*\\*A headers exchange routes a message to a queue if one or more binding attributes of the queue except for x-match match the headers attributes of the message.
	//
	// Separate the attributes with semicolons (;). Separate the key and value of an attribute with a colon (:). Example: x-match:all;type:report;format:pdf. This parameter is available for only headers exchanges. You can set this parameter to an arbitrary value for other types of exchanges.
	//
	// example:
	//
	// x-match:all;type:report;format:pdf
	Argument *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	// The binding key.
	//
	// 	- If the source exchange is not a topic exchange, the binding key must meet the following conventions:
	//
	//     	- The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), forward slashes (/), and at signs (@).
	//
	//     	- The binding key must be 1 to 255 characters in length.
	//
	// 	- If the source exchange is a topic exchange, the binding key must meet the following conventions:
	//
	//     	- The binding key can contain letters, digits, hyphens (-), underscores (_), asterisks (\\*), periods (.), number signs (#), forward slashes (/), and at signs (@).
	//
	//     	- The binding key cannot start or end with a period (.). If a binding key starts with a number sign (#) or an asterisk (\\*), the number sign (#) or asterisk (\\*) must be followed by a period (.). If the binding key ends with a number sign (#) or an asterisk (\\*), the number sign (#) or asterisk (\\*) must be preceded by a period (.). If a number sign (#) or an asterisk (\\*) is used in the middle of a binding key, the number sign (#) or asterisk (\\*) must be preceded and followed by a period (.).
	//
	//     	- The binding key must be 1 to 255 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// .test
	BindingKey *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	// The type of the object that you want to bind to the source exchange. Valid values:
	//
	// 	- \\*\\*0: \\*\\*Queue
	//
	// 	- \\*\\*1: \\*\\*Exchange
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	BindingType *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	// The name of the object that you want to bind to the source exchange. You must create the object in the ApsaraMQ for RabbitMQ console in advance. The vhost of the object is the same as the vhost to which the source exchange specified by **SourceExchange*	- belongs. The vhost of the source exchange is the one specified by **VirtualHost**.
	//
	// This parameter is required.
	//
	// example:
	//
	// DemoQueue
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-v0h1kb9nu***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the source exchange. You must create the source exchange in the ApsaraMQ for RabbitMQ console in advance.
	//
	// This parameter is required.
	//
	// example:
	//
	// NormalEX
	SourceExchange *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
	// The virtual host (vhost) name. You must create the vhost in the ApsaraMQ for RabbitMQ console in advance. The object specified by **DestinationName*	- and the source exchange specified by **SourceExchange*	- must belong to the vhost.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s CreateBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBindingRequest) GoString() string {
	return s.String()
}

func (s *CreateBindingRequest) SetArgument(v string) *CreateBindingRequest {
	s.Argument = &v
	return s
}

func (s *CreateBindingRequest) SetBindingKey(v string) *CreateBindingRequest {
	s.BindingKey = &v
	return s
}

func (s *CreateBindingRequest) SetBindingType(v string) *CreateBindingRequest {
	s.BindingType = &v
	return s
}

func (s *CreateBindingRequest) SetDestinationName(v string) *CreateBindingRequest {
	s.DestinationName = &v
	return s
}

func (s *CreateBindingRequest) SetInstanceId(v string) *CreateBindingRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBindingRequest) SetSourceExchange(v string) *CreateBindingRequest {
	s.SourceExchange = &v
	return s
}

func (s *CreateBindingRequest) SetVirtualHost(v string) *CreateBindingRequest {
	s.VirtualHost = &v
	return s
}

type CreateBindingResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 09768C14-E51C-4F4A-8077-30810032C***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBindingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBindingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBindingResponseBody) SetRequestId(v string) *CreateBindingResponseBody {
	s.RequestId = &v
	return s
}

type CreateBindingResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBindingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBindingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBindingResponse) GoString() string {
	return s.String()
}

func (s *CreateBindingResponse) SetHeaders(v map[string]*string) *CreateBindingResponse {
	s.Headers = v
	return s
}

func (s *CreateBindingResponse) SetStatusCode(v int32) *CreateBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBindingResponse) SetBody(v *CreateBindingResponseBody) *CreateBindingResponse {
	s.Body = v
	return s
}

type CreateExchangeRequest struct {
	// The alternate exchange. An alternate exchange is used to receive messages that fail to be routed to queues from the current exchange.
	//
	// example:
	//
	// DemoAE
	AlternateExchange *string `json:"AlternateExchange,omitempty" xml:"AlternateExchange,omitempty"`
	// Specifies whether to automatically delete the exchange. Valid values:
	//
	// 	- **true**: If the last queue that is bound to the exchange is unbound, the exchange is automatically deleted.
	//
	// 	- **false**: If the last queue that is bound to the exchange is unbound, the exchange is not automatically deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	AutoDeleteState *bool `json:"AutoDeleteState,omitempty" xml:"AutoDeleteState,omitempty"`
	// The name of the exchange that you want to create. The exchange name must meet the following conventions:
	//
	// 	- The name must be 1 to 255 characters in length, and can contain only letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@).
	//
	// 	- After the exchange is created, you cannot change its name. If you want to change its name, delete the exchange and create another exchange.
	//
	// This parameter is required.
	//
	// example:
	//
	// DemoExchange
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	// The exchange type. Valid values:
	//
	// 	- **DIRECT**: An exchange of this type routes a message to the queue whose binding key is exactly the same as the routing key of the message.
	//
	// 	- **TOPIC**: This type of exchange is similar to direct exchanges. An exchange of this type routes a message to one or more queues based on the results of the fuzzy match or multi-condition match between the routing key of the message and the binding keys of the current exchange.
	//
	// 	- **FANOUT**: An exchange of this type routes all received messages to all queues bound to this exchange. You can use a fanout exchange to broadcast messages.
	//
	// 	- **HEADERS**: This type of exchange is similar to direct exchanges. The only difference is that a headers exchange routes messages based on the headers attributes instead of routing keys. When you bind a headers exchange to a queue, you must configure binding attributes in the key-value format for the binding. When you send a message to a headers exchange, you must configure the headers attributes in the key-value format for the message. After a headers exchange receives a message, the exchange routes the message based on the matching results between the headers attributes of the message and the binding attributes of the bound queues.
	//
	// 	- **X-CONSISTENT-HASH**: An exchange of this type allows you to perform hash calculations on routing keys or header values and use consistent hashing to route a message to different queues.
	//
	// This parameter is required.
	//
	// example:
	//
	// DIRECT
	ExchangeType *string `json:"ExchangeType,omitempty" xml:"ExchangeType,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ for which you want to create an exchange.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-v0h1kb9nu***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the exchange is an internal exchange. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Internal *bool `json:"Internal,omitempty" xml:"Internal,omitempty"`
	// The name of the vhost to which the exchange that you want to create belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost  *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
	XDelayedType *string `json:"XDelayedType,omitempty" xml:"XDelayedType,omitempty"`
}

func (s CreateExchangeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExchangeRequest) GoString() string {
	return s.String()
}

func (s *CreateExchangeRequest) SetAlternateExchange(v string) *CreateExchangeRequest {
	s.AlternateExchange = &v
	return s
}

func (s *CreateExchangeRequest) SetAutoDeleteState(v bool) *CreateExchangeRequest {
	s.AutoDeleteState = &v
	return s
}

func (s *CreateExchangeRequest) SetExchangeName(v string) *CreateExchangeRequest {
	s.ExchangeName = &v
	return s
}

func (s *CreateExchangeRequest) SetExchangeType(v string) *CreateExchangeRequest {
	s.ExchangeType = &v
	return s
}

func (s *CreateExchangeRequest) SetInstanceId(v string) *CreateExchangeRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateExchangeRequest) SetInternal(v bool) *CreateExchangeRequest {
	s.Internal = &v
	return s
}

func (s *CreateExchangeRequest) SetVirtualHost(v string) *CreateExchangeRequest {
	s.VirtualHost = &v
	return s
}

func (s *CreateExchangeRequest) SetXDelayedType(v string) *CreateExchangeRequest {
	s.XDelayedType = &v
	return s
}

type CreateExchangeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 021788F6-E50C-4BD6-9F80-66B0A19A6***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExchangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExchangeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExchangeResponseBody) SetRequestId(v string) *CreateExchangeResponseBody {
	s.RequestId = &v
	return s
}

type CreateExchangeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExchangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExchangeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExchangeResponse) GoString() string {
	return s.String()
}

func (s *CreateExchangeResponse) SetHeaders(v map[string]*string) *CreateExchangeResponse {
	s.Headers = v
	return s
}

func (s *CreateExchangeResponse) SetStatusCode(v int32) *CreateExchangeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExchangeResponse) SetBody(v *CreateExchangeResponseBody) *CreateExchangeResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	// example:
	//
	// AutoRenewal
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// example:
	//
	// c2c5d1274axxxxxxxx
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// professional
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 50000
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// example:
	//
	// 128
	MaxEipTps *int64 `json:"MaxEipTps,omitempty" xml:"MaxEipTps,omitempty"`
	// example:
	//
	// 1000
	MaxPrivateTps *int64 `json:"MaxPrivateTps,omitempty" xml:"MaxPrivateTps,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// Month
	PeriodCycle *string `json:"PeriodCycle,omitempty" xml:"PeriodCycle,omitempty"`
	// example:
	//
	// 1000
	QueueCapacity *int32 `json:"QueueCapacity,omitempty" xml:"QueueCapacity,omitempty"`
	// autoRenew和renewStatus都是续费方式，当两个同时填写时，以renewStatus为准
	RenewStatus         *string `json:"RenewStatus,omitempty" xml:"RenewStatus,omitempty"`
	RenewalDurationUnit *string `json:"RenewalDurationUnit,omitempty" xml:"RenewalDurationUnit,omitempty"`
	// example:
	//
	// onDemand
	ServerlessChargeType *string `json:"ServerlessChargeType,omitempty" xml:"ServerlessChargeType,omitempty"`
	// example:
	//
	// 7
	StorageSize *int32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// example:
	//
	// true
	SupportEip *bool `json:"SupportEip,omitempty" xml:"SupportEip,omitempty"`
	// example:
	//
	// true
	SupportTracing *bool `json:"SupportTracing,omitempty" xml:"SupportTracing,omitempty"`
	// example:
	//
	// 3
	TracingStorageTime *int32 `json:"TracingStorageTime,omitempty" xml:"TracingStorageTime,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetAutoRenew(v bool) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenewPeriod(v int32) *CreateInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstanceRequest) SetClientToken(v string) *CreateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetMaxConnections(v int32) *CreateInstanceRequest {
	s.MaxConnections = &v
	return s
}

func (s *CreateInstanceRequest) SetMaxEipTps(v int64) *CreateInstanceRequest {
	s.MaxEipTps = &v
	return s
}

func (s *CreateInstanceRequest) SetMaxPrivateTps(v int64) *CreateInstanceRequest {
	s.MaxPrivateTps = &v
	return s
}

func (s *CreateInstanceRequest) SetPaymentType(v string) *CreateInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v int32) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriodCycle(v string) *CreateInstanceRequest {
	s.PeriodCycle = &v
	return s
}

func (s *CreateInstanceRequest) SetQueueCapacity(v int32) *CreateInstanceRequest {
	s.QueueCapacity = &v
	return s
}

func (s *CreateInstanceRequest) SetRenewStatus(v string) *CreateInstanceRequest {
	s.RenewStatus = &v
	return s
}

func (s *CreateInstanceRequest) SetRenewalDurationUnit(v string) *CreateInstanceRequest {
	s.RenewalDurationUnit = &v
	return s
}

func (s *CreateInstanceRequest) SetServerlessChargeType(v string) *CreateInstanceRequest {
	s.ServerlessChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetStorageSize(v int32) *CreateInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateInstanceRequest) SetSupportEip(v bool) *CreateInstanceRequest {
	s.SupportEip = &v
	return s
}

func (s *CreateInstanceRequest) SetSupportTracing(v bool) *CreateInstanceRequest {
	s.SupportTracing = &v
	return s
}

func (s *CreateInstanceRequest) SetTracingStorageTime(v int32) *CreateInstanceRequest {
	s.TracingStorageTime = &v
	return s
}

type CreateInstanceResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// amqp-cn-xxxxx
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// xxx failed,xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CCBB1225-C392-480E-8C7F-D09AB2CD2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetCode(v int32) *CreateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponseBody) SetData(v interface{}) *CreateInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateInstanceResponseBody) SetMessage(v string) *CreateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetSuccess(v bool) *CreateInstanceResponseBody {
	s.Success = &v
	return s
}

type CreateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetStatusCode(v int32) *CreateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type CreateQueueRequest struct {
	// Specifies whether to automatically delete the queue. Valid values:
	//
	// 	- true: The queue is automatically deleted. After the last consumer unsubscribes from the queue, the queue is automatically deleted.
	//
	// 	- false: The queue is not automatically deleted.
	//
	// example:
	//
	// false
	AutoDeleteState *bool `json:"AutoDeleteState,omitempty" xml:"AutoDeleteState,omitempty"`
	// The validity period after which the queue is automatically deleted. If the queue is not accessed within the specified period of time, the queue is automatically deleted.
	//
	// Unit: milliseconds.
	//
	// >  You can use the feature that corresponds to this parameter only after you enable the feature. To enable the feature, [submit a ticket](https://ticket-intl.console.aliyun.com/#/ticket/createIndex).
	//
	// example:
	//
	// 10000
	AutoExpireState *int64 `json:"AutoExpireState,omitempty" xml:"AutoExpireState,omitempty"`
	// The dead-letter exchange. A dead-letter exchange is used to receive rejected messages.
	//
	// If a consumer rejects a message that cannot be redelivered, ApsaraMQ for RabbitMQ routes the message to the specified dead-letter exchange. Then, the dead-letter exchange routes the message to the queue that is bound to the dead-letter exchange for storage.
	//
	// example:
	//
	// DLExchange
	DeadLetterExchange *string `json:"DeadLetterExchange,omitempty" xml:"DeadLetterExchange,omitempty"`
	// The dead-letter routing key. The key must be 1 to 255 characters in length, and can contain only letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@).
	//
	// example:
	//
	// test.dl
	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitempty" xml:"DeadLetterRoutingKey,omitempty"`
	// Specifies whether the exchange is an exclusive exchange. Valid values:
	//
	// 	- true: The exchange is an exclusive exchange. Only the connection that declares the exclusive exchange can use the exclusive exchange. After the connection is closed, the exclusive exchange is automatically deleted.
	//
	// 	- false: The exchange is not an exclusive exchange.
	//
	// example:
	//
	// false
	ExclusiveState *bool `json:"ExclusiveState,omitempty" xml:"ExclusiveState,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance on which you want to create a queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-v0h1kb9nu***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is unavailable in the current version of ApsaraMQ for RabbitMQ.
	//
	// The maximum number of messages that can be stored in the queue. If this threshold is exceeded, the earliest stored messages in the queue are deleted.
	//
	// example:
	//
	// 1000
	MaxLength *int64 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// Queue priorities are not supported. The value does not affect the call or return results.
	//
	// example:
	//
	// 10
	MaximumPriority *int32 `json:"MaximumPriority,omitempty" xml:"MaximumPriority,omitempty"`
	// The message time to live (TTL) of the queue.
	//
	// 	- If the retention period of a message in the queue exceeds the message TTL of the queue, the message expires.
	//
	// 	- The message TTL must be set to a non-negative integer. The maximum message TTL is one day. Unit: milliseconds. For example, if the message TTL is 1,000 milliseconds, the message can be retained for up to 1 second in the queue.
	//
	// example:
	//
	// 1000
	MessageTTL *int64 `json:"MessageTTL,omitempty" xml:"MessageTTL,omitempty"`
	// The name of the queue that you want to create.
	//
	// 	- The name must be 1 to 255 characters in length, and can contain only letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@).
	//
	// 	- After the queue is created, you cannot change the name of the queue. If you want to change the name of the queue, delete the queue and create another queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// DemoQueue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The name of the vhost to which the queue that you want to create belongs. The name must be 1 to 255 characters in length, and can contain only letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@).
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s CreateQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueRequest) GoString() string {
	return s.String()
}

func (s *CreateQueueRequest) SetAutoDeleteState(v bool) *CreateQueueRequest {
	s.AutoDeleteState = &v
	return s
}

func (s *CreateQueueRequest) SetAutoExpireState(v int64) *CreateQueueRequest {
	s.AutoExpireState = &v
	return s
}

func (s *CreateQueueRequest) SetDeadLetterExchange(v string) *CreateQueueRequest {
	s.DeadLetterExchange = &v
	return s
}

func (s *CreateQueueRequest) SetDeadLetterRoutingKey(v string) *CreateQueueRequest {
	s.DeadLetterRoutingKey = &v
	return s
}

func (s *CreateQueueRequest) SetExclusiveState(v bool) *CreateQueueRequest {
	s.ExclusiveState = &v
	return s
}

func (s *CreateQueueRequest) SetInstanceId(v string) *CreateQueueRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateQueueRequest) SetMaxLength(v int64) *CreateQueueRequest {
	s.MaxLength = &v
	return s
}

func (s *CreateQueueRequest) SetMaximumPriority(v int32) *CreateQueueRequest {
	s.MaximumPriority = &v
	return s
}

func (s *CreateQueueRequest) SetMessageTTL(v int64) *CreateQueueRequest {
	s.MessageTTL = &v
	return s
}

func (s *CreateQueueRequest) SetQueueName(v string) *CreateQueueRequest {
	s.QueueName = &v
	return s
}

func (s *CreateQueueRequest) SetVirtualHost(v string) *CreateQueueRequest {
	s.VirtualHost = &v
	return s
}

type CreateQueueResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 59B52E2C-0B8E-44EC-A314-D0314A50***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQueueResponseBody) SetRequestId(v string) *CreateQueueResponseBody {
	s.RequestId = &v
	return s
}

type CreateQueueResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueResponse) GoString() string {
	return s.String()
}

func (s *CreateQueueResponse) SetHeaders(v map[string]*string) *CreateQueueResponse {
	s.Headers = v
	return s
}

func (s *CreateQueueResponse) SetStatusCode(v int32) *CreateQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQueueResponse) SetBody(v *CreateQueueResponseBody) *CreateQueueResponse {
	s.Body = v
	return s
}

type CreateVirtualHostRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-v0h1kb9n***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the vhost that you want to create. Valid values:
	//
	// 	- The name can contain letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slash (/), and at signs (@).
	//
	// 	- The name must be 1 to 255 characters in length.
	//
	// 	- After the vhost is created, you cannot change its name. If you want to change the name of a vhost, delete the vhost and create another vhost.
	//
	// This parameter is required.
	//
	// example:
	//
	// Demo
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s CreateVirtualHostRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualHostRequest) GoString() string {
	return s.String()
}

func (s *CreateVirtualHostRequest) SetInstanceId(v string) *CreateVirtualHostRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateVirtualHostRequest) SetVirtualHost(v string) *CreateVirtualHostRequest {
	s.VirtualHost = &v
	return s
}

type CreateVirtualHostResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 628705FD-03EE-4ABE-BB21-E1672960***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVirtualHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualHostResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirtualHostResponseBody) SetRequestId(v string) *CreateVirtualHostResponseBody {
	s.RequestId = &v
	return s
}

type CreateVirtualHostResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVirtualHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVirtualHostResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualHostResponse) GoString() string {
	return s.String()
}

func (s *CreateVirtualHostResponse) SetHeaders(v map[string]*string) *CreateVirtualHostResponse {
	s.Headers = v
	return s
}

func (s *CreateVirtualHostResponse) SetStatusCode(v int32) *CreateVirtualHostResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVirtualHostResponse) SetBody(v *CreateVirtualHostResponseBody) *CreateVirtualHostResponse {
	s.Body = v
	return s
}

type DeleteAccountRequest struct {
	// The timestamp that indicates when the pair of static username and password that you want to delete was created. Unit: milliseconds.
	//
	// You can call the [ListAccounts](https://help.aliyun.com/document_detail/472730.html) operation to view the timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1671175303522
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The pair of username and password that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// MjphbXFwLWNuLXVxbTJ5cjc3djAwMzpMVEFJNXQ4YmVNbVZNMWVSWnRFSjZ2Zm1=
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) SetCreateTimestamp(v int64) *DeleteAccountRequest {
	s.CreateTimestamp = &v
	return s
}

func (s *DeleteAccountRequest) SetUserName(v string) *DeleteAccountRequest {
	s.UserName = &v
	return s
}

type DeleteAccountResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 021788F6-E50C-4BD6-9F80-66B0A19A6***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponseBody) SetCode(v int32) *DeleteAccountResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAccountResponseBody) SetData(v bool) *DeleteAccountResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAccountResponseBody) SetMessage(v string) *DeleteAccountResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAccountResponseBody) SetRequestId(v string) *DeleteAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccountResponseBody) SetSuccess(v bool) *DeleteAccountResponseBody {
	s.Success = &v
	return s
}

type DeleteAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponse) SetHeaders(v map[string]*string) *DeleteAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccountResponse) SetStatusCode(v int32) *DeleteAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccountResponse) SetBody(v *DeleteAccountResponseBody) *DeleteAccountResponse {
	s.Body = v
	return s
}

type DeleteBindingRequest struct {
	// The binding key.
	//
	// This parameter is required.
	//
	// example:
	//
	// .test.
	BindingKey *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	// The type of the object that you want to unbind from the source exchange. Valid values:
	//
	// 	- **QUEUE**
	//
	// 	- **EXCHANGE**
	//
	// This parameter is required.
	//
	// example:
	//
	// QUEUE
	BindingType *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	// The name of the object that you want to unbind from the source exchange.
	//
	// This parameter is required.
	//
	// example:
	//
	// DemoQueue
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-v0h1kb9nu***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the source exchange.
	//
	// This parameter is required.
	//
	// example:
	//
	// NormalEX
	SourceExchange *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
	// The vhost name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s DeleteBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBindingRequest) GoString() string {
	return s.String()
}

func (s *DeleteBindingRequest) SetBindingKey(v string) *DeleteBindingRequest {
	s.BindingKey = &v
	return s
}

func (s *DeleteBindingRequest) SetBindingType(v string) *DeleteBindingRequest {
	s.BindingType = &v
	return s
}

func (s *DeleteBindingRequest) SetDestinationName(v string) *DeleteBindingRequest {
	s.DestinationName = &v
	return s
}

func (s *DeleteBindingRequest) SetInstanceId(v string) *DeleteBindingRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteBindingRequest) SetSourceExchange(v string) *DeleteBindingRequest {
	s.SourceExchange = &v
	return s
}

func (s *DeleteBindingRequest) SetVirtualHost(v string) *DeleteBindingRequest {
	s.VirtualHost = &v
	return s
}

type DeleteBindingResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 021788F6-E50C-4BD6-9F80-66B0A19A6***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBindingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBindingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBindingResponseBody) SetRequestId(v string) *DeleteBindingResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBindingResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBindingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBindingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBindingResponse) GoString() string {
	return s.String()
}

func (s *DeleteBindingResponse) SetHeaders(v map[string]*string) *DeleteBindingResponse {
	s.Headers = v
	return s
}

func (s *DeleteBindingResponse) SetStatusCode(v int32) *DeleteBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBindingResponse) SetBody(v *DeleteBindingResponseBody) *DeleteBindingResponse {
	s.Body = v
	return s
}

type DeleteExchangeRequest struct {
	// The name of the exchange that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// DemoExchange
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance whose exchange you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-v0h1kb9nu***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The vhost to which the exchange that you want to delete belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s DeleteExchangeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExchangeRequest) GoString() string {
	return s.String()
}

func (s *DeleteExchangeRequest) SetExchangeName(v string) *DeleteExchangeRequest {
	s.ExchangeName = &v
	return s
}

func (s *DeleteExchangeRequest) SetInstanceId(v string) *DeleteExchangeRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteExchangeRequest) SetVirtualHost(v string) *DeleteExchangeRequest {
	s.VirtualHost = &v
	return s
}

type DeleteExchangeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6961FFB8-6358-4EDC-9E3C-4A0C56CE6***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExchangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExchangeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExchangeResponseBody) SetRequestId(v string) *DeleteExchangeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteExchangeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExchangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExchangeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExchangeResponse) GoString() string {
	return s.String()
}

func (s *DeleteExchangeResponse) SetHeaders(v map[string]*string) *DeleteExchangeResponse {
	s.Headers = v
	return s
}

func (s *DeleteExchangeResponse) SetStatusCode(v int32) *DeleteExchangeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExchangeResponse) SetBody(v *DeleteExchangeResponseBody) *DeleteExchangeResponse {
	s.Body = v
	return s
}

type DeleteQueueRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1880770869023***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The queue name.
	//
	// This parameter is required.
	//
	// example:
	//
	// DemoQueue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The vhost name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s DeleteQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueRequest) GoString() string {
	return s.String()
}

func (s *DeleteQueueRequest) SetInstanceId(v string) *DeleteQueueRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteQueueRequest) SetQueueName(v string) *DeleteQueueRequest {
	s.QueueName = &v
	return s
}

func (s *DeleteQueueRequest) SetVirtualHost(v string) *DeleteQueueRequest {
	s.VirtualHost = &v
	return s
}

type DeleteQueueResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 92385FD2-624A-48C9-8FB5-753F2AFA***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponseBody) SetRequestId(v string) *DeleteQueueResponseBody {
	s.RequestId = &v
	return s
}

type DeleteQueueResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueResponse) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponse) SetHeaders(v map[string]*string) *DeleteQueueResponse {
	s.Headers = v
	return s
}

func (s *DeleteQueueResponse) SetStatusCode(v int32) *DeleteQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQueueResponse) SetBody(v *DeleteQueueResponseBody) *DeleteQueueResponse {
	s.Body = v
	return s
}

type DeleteVirtualHostRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance to which the vhost you want to delete belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-v0h1kb9nu***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the vhost that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s DeleteVirtualHostRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualHostRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualHostRequest) SetInstanceId(v string) *DeleteVirtualHostRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteVirtualHostRequest) SetVirtualHost(v string) *DeleteVirtualHostRequest {
	s.VirtualHost = &v
	return s
}

type DeleteVirtualHostResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4311050D-BD63-48F9-822B-947A75A1***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVirtualHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualHostResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVirtualHostResponseBody) SetRequestId(v string) *DeleteVirtualHostResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVirtualHostResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVirtualHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVirtualHostResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualHostResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualHostResponse) SetHeaders(v map[string]*string) *DeleteVirtualHostResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirtualHostResponse) SetStatusCode(v int32) *DeleteVirtualHostResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVirtualHostResponse) SetBody(v *DeleteVirtualHostResponseBody) *DeleteVirtualHostResponse {
	s.Body = v
	return s
}

type GetMetadataAmountRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// amqp-cn-v0h1kb9n***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetMetadataAmountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMetadataAmountRequest) GoString() string {
	return s.String()
}

func (s *GetMetadataAmountRequest) SetInstanceId(v string) *GetMetadataAmountRequest {
	s.InstanceId = &v
	return s
}

type GetMetadataAmountResponseBody struct {
	// The returned data.
	Data *GetMetadataAmountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B75ACF23-2BEB-44AC-A0B6-AE14EDCA***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMetadataAmountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMetadataAmountResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetadataAmountResponseBody) SetData(v *GetMetadataAmountResponseBodyData) *GetMetadataAmountResponseBody {
	s.Data = v
	return s
}

func (s *GetMetadataAmountResponseBody) SetRequestId(v string) *GetMetadataAmountResponseBody {
	s.RequestId = &v
	return s
}

type GetMetadataAmountResponseBodyData struct {
	// The number of created exchanges on the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// 7
	CurrentExchanges *int32 `json:"CurrentExchanges,omitempty" xml:"CurrentExchanges,omitempty"`
	// The number of created queues on the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// 1
	CurrentQueues *int32 `json:"CurrentQueues,omitempty" xml:"CurrentQueues,omitempty"`
	// The number of created vhosts on the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// 1
	CurrentVirtualHosts *int32 `json:"CurrentVirtualHosts,omitempty" xml:"CurrentVirtualHosts,omitempty"`
	// The maximum number of exchanges that can be created on the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// 20
	MaxExchanges *int32 `json:"MaxExchanges,omitempty" xml:"MaxExchanges,omitempty"`
	// The maximum number of queues that can be created on the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// 20
	MaxQueues *int32 `json:"MaxQueues,omitempty" xml:"MaxQueues,omitempty"`
	// The maximum number of vhosts that can be created on the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// 10
	MaxVirtualHosts *int32 `json:"MaxVirtualHosts,omitempty" xml:"MaxVirtualHosts,omitempty"`
}

func (s GetMetadataAmountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMetadataAmountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetadataAmountResponseBodyData) SetCurrentExchanges(v int32) *GetMetadataAmountResponseBodyData {
	s.CurrentExchanges = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetCurrentQueues(v int32) *GetMetadataAmountResponseBodyData {
	s.CurrentQueues = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetCurrentVirtualHosts(v int32) *GetMetadataAmountResponseBodyData {
	s.CurrentVirtualHosts = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetMaxExchanges(v int32) *GetMetadataAmountResponseBodyData {
	s.MaxExchanges = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetMaxQueues(v int32) *GetMetadataAmountResponseBodyData {
	s.MaxQueues = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetMaxVirtualHosts(v int32) *GetMetadataAmountResponseBodyData {
	s.MaxVirtualHosts = &v
	return s
}

type GetMetadataAmountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetadataAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetadataAmountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMetadataAmountResponse) GoString() string {
	return s.String()
}

func (s *GetMetadataAmountResponse) SetHeaders(v map[string]*string) *GetMetadataAmountResponse {
	s.Headers = v
	return s
}

func (s *GetMetadataAmountResponse) SetStatusCode(v int32) *GetMetadataAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetadataAmountResponse) SetBody(v *GetMetadataAmountResponseBody) *GetMetadataAmountResponse {
	s.Body = v
	return s
}

type ListAccountsRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance for which you want to query the static username and password.
	//
	// example:
	//
	// amqp-cn-20p****04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListAccountsRequest) SetInstanceId(v string) *ListAccountsRequest {
	s.InstanceId = &v
	return s
}

type ListAccountsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data map[string][]*DataValue `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 549A5A97-FE61-5A23-8126-3A11929C1EC4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBody) SetCode(v int32) *ListAccountsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAccountsResponseBody) SetData(v map[string][]*DataValue) *ListAccountsResponseBody {
	s.Data = v
	return s
}

func (s *ListAccountsResponseBody) SetMessage(v string) *ListAccountsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAccountsResponseBody) SetRequestId(v string) *ListAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountsResponseBody) SetSuccess(v bool) *ListAccountsResponseBody {
	s.Success = &v
	return s
}

type ListAccountsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListAccountsResponse) SetHeaders(v map[string]*string) *ListAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListAccountsResponse) SetStatusCode(v int32) *ListAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountsResponse) SetBody(v *ListAccountsResponseBody) *ListAccountsResponse {
	s.Body = v
	return s
}

type ListBindingsRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1880770869023***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return. Valid values:
	//
	// **1 to 100**
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end position of the previous returned page. To obtain the next batch of data, call the operation again by using the value of NextToken returned by the previous request. If you call this operation for the first time or want to query all results, set NextToken to an empty string.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The vhost name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s ListBindingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListBindingsRequest) SetInstanceId(v string) *ListBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListBindingsRequest) SetMaxResults(v int32) *ListBindingsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBindingsRequest) SetNextToken(v string) *ListBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListBindingsRequest) SetVirtualHost(v string) *ListBindingsRequest {
	s.VirtualHost = &v
	return s
}

type ListBindingsResponseBody struct {
	// The returned data.
	Data *ListBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E0A71208-3E87-4732-81CC-B18E0B4B1***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBindingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBindingsResponseBody) SetData(v *ListBindingsResponseBodyData) *ListBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListBindingsResponseBody) SetRequestId(v string) *ListBindingsResponseBody {
	s.RequestId = &v
	return s
}

type ListBindingsResponseBodyData struct {
	// The bindings.
	Bindings []*ListBindingsResponseBodyDataBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current returned page. If this parameter is empty, all data is retrieved.
	//
	// example:
	//
	// caebacccb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListBindingsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBindingsResponseBodyData) SetBindings(v []*ListBindingsResponseBodyDataBindings) *ListBindingsResponseBodyData {
	s.Bindings = v
	return s
}

func (s *ListBindingsResponseBodyData) SetMaxResults(v int32) *ListBindingsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListBindingsResponseBodyData) SetNextToken(v string) *ListBindingsResponseBodyData {
	s.NextToken = &v
	return s
}

type ListBindingsResponseBodyDataBindings struct {
	// The x-match attribute. Valid values:
	//
	// 	- **all:*	- A headers exchange routes a message to a queue only if all binding attributes of the queue except for x-match match the headers attributes of the message. This value is the default value.
	//
	// 	- **any:*	- A headers exchange routes a message to a queue if one or more binding attributes of the queue except for x-match match the headers attributes of the message.
	//
	// This parameter is available only for headers exchanges.
	//
	// example:
	//
	// all
	Argument *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	// The binding key.
	//
	// 	- If the source exchange is not a topic exchange, the binding key must meet the following conventions:
	//
	//     	- The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), forward slashes (/), and at signs (@).
	//
	//     	- The binding key must be 1 to 255 characters in length.
	//
	// 	- If the source exchange is a topic exchange, the binding key must meet the following conventions:
	//
	//     	- The binding key can contain letters, digits, hyphens (-), underscores (_), asterisks (\\*), periods (.), number signs (#), forward slashes (/), and at signs (@).
	//
	//     	- The binding key cannot start or end with a period (.). If a binding key starts with a number sign (#) or an asterisk (\\*), the number sign (#) or asterisk (\\*) must be followed by a period (.). If the binding key ends with a number sign (#) or an asterisk (\\*), the number sign (#) or asterisk (\\*) must be preceded by a period (.). If a number sign (#) or an asterisk (\\*) is used in the middle of a binding key, the number sign (#) or asterisk (\\*) must be preceded and followed by a period (.).
	//
	//     	- The binding key must be 1 to 255 characters in length.
	//
	// example:
	//
	// amq.test
	BindingKey *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	// The type of the object to which the source exchange is bound. Valid values:
	//
	// 	- **QUEUE**
	//
	// 	- **EXCHANGE**
	//
	// example:
	//
	// QUEUE
	BindingType *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	// The name of the object to which the source exchange is bound.
	//
	// example:
	//
	// QueueTest
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
	// The name of the source exchange.
	//
	// example:
	//
	// test
	SourceExchange *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
}

func (s ListBindingsResponseBodyDataBindings) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsResponseBodyDataBindings) GoString() string {
	return s.String()
}

func (s *ListBindingsResponseBodyDataBindings) SetArgument(v string) *ListBindingsResponseBodyDataBindings {
	s.Argument = &v
	return s
}

func (s *ListBindingsResponseBodyDataBindings) SetBindingKey(v string) *ListBindingsResponseBodyDataBindings {
	s.BindingKey = &v
	return s
}

func (s *ListBindingsResponseBodyDataBindings) SetBindingType(v string) *ListBindingsResponseBodyDataBindings {
	s.BindingType = &v
	return s
}

func (s *ListBindingsResponseBodyDataBindings) SetDestinationName(v string) *ListBindingsResponseBodyDataBindings {
	s.DestinationName = &v
	return s
}

func (s *ListBindingsResponseBodyDataBindings) SetSourceExchange(v string) *ListBindingsResponseBodyDataBindings {
	s.SourceExchange = &v
	return s
}

type ListBindingsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBindingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListBindingsResponse) SetHeaders(v map[string]*string) *ListBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListBindingsResponse) SetStatusCode(v int32) *ListBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBindingsResponse) SetBody(v *ListBindingsResponseBody) *ListBindingsResponse {
	s.Body = v
	return s
}

type ListDownStreamBindingsRequest struct {
	// The exchange name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance to which the exchange belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1880770869023***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end position of the previous returned page. To obtain the next batch of data, call the operation again by using the value of NextToken returned by the previous request. If you call this operation for the first time or want to query all results, set NextToken to an empty string.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the vhost to which the exchange belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s ListDownStreamBindingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDownStreamBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsRequest) SetExchangeName(v string) *ListDownStreamBindingsRequest {
	s.ExchangeName = &v
	return s
}

func (s *ListDownStreamBindingsRequest) SetInstanceId(v string) *ListDownStreamBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDownStreamBindingsRequest) SetMaxResults(v int32) *ListDownStreamBindingsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDownStreamBindingsRequest) SetNextToken(v string) *ListDownStreamBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDownStreamBindingsRequest) SetVirtualHost(v string) *ListDownStreamBindingsRequest {
	s.VirtualHost = &v
	return s
}

type ListDownStreamBindingsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListDownStreamBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9C1E0502-0790-4FDB-8C96-6D5C8D9B7***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDownStreamBindingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDownStreamBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsResponseBody) SetCode(v int32) *ListDownStreamBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDownStreamBindingsResponseBody) SetData(v *ListDownStreamBindingsResponseBodyData) *ListDownStreamBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListDownStreamBindingsResponseBody) SetMessage(v string) *ListDownStreamBindingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDownStreamBindingsResponseBody) SetRequestId(v string) *ListDownStreamBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDownStreamBindingsResponseBody) SetSuccess(v bool) *ListDownStreamBindingsResponseBody {
	s.Success = &v
	return s
}

type ListDownStreamBindingsResponseBodyData struct {
	// The bindings.
	Bindings []*ListDownStreamBindingsResponseBodyDataBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current returned page. If this parameter is empty, all data is retrieved.
	//
	// example:
	//
	// caebacccb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListDownStreamBindingsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDownStreamBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsResponseBodyData) SetBindings(v []*ListDownStreamBindingsResponseBodyDataBindings) *ListDownStreamBindingsResponseBodyData {
	s.Bindings = v
	return s
}

func (s *ListDownStreamBindingsResponseBodyData) SetMaxResults(v int32) *ListDownStreamBindingsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyData) SetNextToken(v string) *ListDownStreamBindingsResponseBodyData {
	s.NextToken = &v
	return s
}

type ListDownStreamBindingsResponseBodyDataBindings struct {
	// The x-match attribute. Valid values:
	//
	// 	- **all:*	- A headers exchange routes a message to a queue only if all binding attributes of the queue except for x-match match the headers attributes of the message. This value is the default value.
	//
	// 	- **any:*	- A headers exchange routes a message to a queue if one or more binding attributes of the queue except for x-match match the headers attributes of the message.
	//
	// This parameter is available only for headers exchanges.
	//
	// example:
	//
	// test
	Argument *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	// The binding key.
	//
	// 	- If the source exchange is not a topic exchange, the binding key must meet the following conventions:
	//
	//     	- The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), forward slashes (/), and at signs (@).
	//
	//     	- The binding key must be 1 to 255 characters in length.
	//
	// 	- If the source exchange is a topic exchange, the binding key must meet the following conventions:
	//
	//     	- The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@).
	//
	//     	- The binding key cannot start or end with a period (.). If a binding key starts with a number sign (#) or an asterisk (\\*), the number sign (#) or asterisk (\\*) must be followed by a period (.). If the binding key ends with a number sign (#) or an asterisk (\\*), the number sign (#) or asterisk (\\*) must be preceded by a period (.). If a number sign (#) or an asterisk (\\*) is used in the middle of a binding key, the number sign (#) or asterisk (\\*) must be preceded and followed by a period (.).
	//
	//     	- The binding key must be 1 to 255 characters in length.
	//
	// example:
	//
	// amq.test
	BindingKey *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	// The type of the object to which the source exchange is bound. Valid values:
	//
	// 	- **QUEUE**
	//
	// 	- **EXCHANGE**
	//
	// example:
	//
	// QUEUE
	BindingType *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	// The name of the object to which the source exchange is bound.
	//
	// example:
	//
	// QueueTest
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
	// The name of the source exchange.
	//
	// example:
	//
	// test
	SourceExchange *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
}

func (s ListDownStreamBindingsResponseBodyDataBindings) String() string {
	return tea.Prettify(s)
}

func (s ListDownStreamBindingsResponseBodyDataBindings) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetArgument(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.Argument = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetBindingKey(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.BindingKey = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetBindingType(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.BindingType = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetDestinationName(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.DestinationName = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetSourceExchange(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.SourceExchange = &v
	return s
}

type ListDownStreamBindingsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDownStreamBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDownStreamBindingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDownStreamBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsResponse) SetHeaders(v map[string]*string) *ListDownStreamBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListDownStreamBindingsResponse) SetStatusCode(v int32) *ListDownStreamBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDownStreamBindingsResponse) SetBody(v *ListDownStreamBindingsResponseBody) *ListDownStreamBindingsResponse {
	s.Body = v
	return s
}

type ListExchangeUpStreamBindingsRequest struct {
	// The exchange name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1880770869023***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end position of the previous returned page. To obtain the next batch of data, call the operation again by using the value of NextToken returned by the previous request. If you call this operation for the first time or want to query all results, set NextToken to an empty string.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The virtual host (vhost) name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s ListExchangeUpStreamBindingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExchangeUpStreamBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsRequest) SetExchangeName(v string) *ListExchangeUpStreamBindingsRequest {
	s.ExchangeName = &v
	return s
}

func (s *ListExchangeUpStreamBindingsRequest) SetInstanceId(v string) *ListExchangeUpStreamBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListExchangeUpStreamBindingsRequest) SetMaxResults(v int32) *ListExchangeUpStreamBindingsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExchangeUpStreamBindingsRequest) SetNextToken(v string) *ListExchangeUpStreamBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListExchangeUpStreamBindingsRequest) SetVirtualHost(v string) *ListExchangeUpStreamBindingsRequest {
	s.VirtualHost = &v
	return s
}

type ListExchangeUpStreamBindingsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListExchangeUpStreamBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2DCCCE88-BC82-4A4F-AF5E-9A759672B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListExchangeUpStreamBindingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExchangeUpStreamBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetCode(v int32) *ListExchangeUpStreamBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetData(v *ListExchangeUpStreamBindingsResponseBodyData) *ListExchangeUpStreamBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetMessage(v string) *ListExchangeUpStreamBindingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetRequestId(v string) *ListExchangeUpStreamBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetSuccess(v bool) *ListExchangeUpStreamBindingsResponseBody {
	s.Success = &v
	return s
}

type ListExchangeUpStreamBindingsResponseBodyData struct {
	// The bindings.
	Bindings []*ListExchangeUpStreamBindingsResponseBodyDataBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current returned page. If this parameter is empty, all data is retrieved.
	//
	// example:
	//
	// caebacccb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListExchangeUpStreamBindingsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListExchangeUpStreamBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsResponseBodyData) SetBindings(v []*ListExchangeUpStreamBindingsResponseBodyDataBindings) *ListExchangeUpStreamBindingsResponseBodyData {
	s.Bindings = v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyData) SetMaxResults(v int32) *ListExchangeUpStreamBindingsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyData) SetNextToken(v string) *ListExchangeUpStreamBindingsResponseBodyData {
	s.NextToken = &v
	return s
}

type ListExchangeUpStreamBindingsResponseBodyDataBindings struct {
	// The x-match attribute. Valid values:
	//
	// 	- **all:*	- A headers exchange routes a message to a queue only if all binding attributes of the queue except for x-match match the headers attributes of the message. This value is the default value.
	//
	// 	- **any:*	- A headers exchange routes a message to a queue if one or more binding attributes of the queue except for x-match match the headers attributes of the message.
	//
	// This parameter is available only for headers exchanges.
	//
	// example:
	//
	// all
	Argument *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	// The binding key.
	//
	// 	- If the source exchange is not a topic exchange, the binding key must meet the following conventions:
	//
	//     	- The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), forward slashes (/), and at signs (@).
	//
	//     	- The binding key must be 1 to 255 characters in length.
	//
	// 	- If the source exchange is a topic exchange, the binding key must meet the following conventions:
	//
	//     	- The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@).
	//
	//     	- The binding key cannot start or end with a period (.). If a binding key starts with a number sign (#) or an asterisk (\\*), the number sign (#) or asterisk (\\*) must be followed by a period (.). If the binding key ends with a number sign (#) or an asterisk (\\*), the number sign (#) or asterisk (\\*) must be preceded by a period (.). If a number sign (#) or an asterisk (\\*) is used in the middle of a binding key, the number sign (#) or asterisk (\\*) must be preceded and followed by a period (.).
	//
	//     	- The binding key must be 1 to 255 characters in length.
	//
	// example:
	//
	// amq.dle.test
	BindingKey *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	// The type of the object to which the source exchange is bound. Valid values:
	//
	// 	- **QUEUE**
	//
	// 	- **EXCHANGE**
	//
	// example:
	//
	// EXCHANGE
	BindingType *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	// The name of the object to which the source exchange is bound.
	//
	// example:
	//
	// test
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
	// The name of the source exchange.
	//
	// example:
	//
	// dle
	SourceExchange *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
}

func (s ListExchangeUpStreamBindingsResponseBodyDataBindings) String() string {
	return tea.Prettify(s)
}

func (s ListExchangeUpStreamBindingsResponseBodyDataBindings) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetArgument(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.Argument = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetBindingKey(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.BindingKey = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetBindingType(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.BindingType = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetDestinationName(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.DestinationName = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetSourceExchange(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.SourceExchange = &v
	return s
}

type ListExchangeUpStreamBindingsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExchangeUpStreamBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExchangeUpStreamBindingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExchangeUpStreamBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsResponse) SetHeaders(v map[string]*string) *ListExchangeUpStreamBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListExchangeUpStreamBindingsResponse) SetStatusCode(v int32) *ListExchangeUpStreamBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponse) SetBody(v *ListExchangeUpStreamBindingsResponseBody) *ListExchangeUpStreamBindingsResponse {
	s.Body = v
	return s
}

type ListExchangesRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-7pp2mwbc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return. Valid values: **1 to 100**
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If you call this operation for the first time or a next query is not required, leave this parameter empty.
	//
	// 	- If a next query is to be sent, set the value to the value of `NextToken` that is returned from the previous request.
	//
	// example:
	//
	// AAAANDQBYW1xcC1jbi03cHAybXdiY3AwMGEBdmhvc3QBAXNkZndhYWJhATE2NDkzMTM4OTU5NDIB4o3z1pPwWzk4aYuiRffi8R6-****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The vhost name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s ListExchangesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExchangesRequest) GoString() string {
	return s.String()
}

func (s *ListExchangesRequest) SetInstanceId(v string) *ListExchangesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListExchangesRequest) SetMaxResults(v int32) *ListExchangesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExchangesRequest) SetNextToken(v string) *ListExchangesRequest {
	s.NextToken = &v
	return s
}

func (s *ListExchangesRequest) SetVirtualHost(v string) *ListExchangesRequest {
	s.VirtualHost = &v
	return s
}

type ListExchangesResponseBody struct {
	// The returned data.
	Data *ListExchangesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// FEBA5E0C-50D0-4FA6-A794-4901E5465***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListExchangesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExchangesResponseBody) GoString() string {
	return s.String()
}

func (s *ListExchangesResponseBody) SetData(v *ListExchangesResponseBodyData) *ListExchangesResponseBody {
	s.Data = v
	return s
}

func (s *ListExchangesResponseBody) SetRequestId(v string) *ListExchangesResponseBody {
	s.RequestId = &v
	return s
}

type ListExchangesResponseBodyData struct {
	// The exchanges.
	Exchanges []*ListExchangesResponseBodyDataExchanges `json:"Exchanges,omitempty" xml:"Exchanges,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current returned page.``
	//
	// 	- If the value of this parameter is empty, the next query is not required and the token used to start the next query is unavailable.``
	//
	// 	- If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.``
	//
	// example:
	//
	// AAAANDQBYW1xcC1jbi03cHAybXdiY3AwMGEBdmhvc3QBAXNkZndhYWJhATE2NDkzMTM4OTU5NDIB4o3z1pPwWzk4aYuiRffi8R6-****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListExchangesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListExchangesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExchangesResponseBodyData) SetExchanges(v []*ListExchangesResponseBodyDataExchanges) *ListExchangesResponseBodyData {
	s.Exchanges = v
	return s
}

func (s *ListExchangesResponseBodyData) SetMaxResults(v int32) *ListExchangesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListExchangesResponseBodyData) SetNextToken(v string) *ListExchangesResponseBodyData {
	s.NextToken = &v
	return s
}

type ListExchangesResponseBodyDataExchanges struct {
	// The attributes. This parameter is unavailable in the current version.
	//
	// example:
	//
	// test
	Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// Indicates whether the exchange was automatically deleted.
	//
	// example:
	//
	// false
	AutoDeleteState *bool `json:"AutoDeleteState,omitempty" xml:"AutoDeleteState,omitempty"`
	// The timestamp that indicates when the exchange was created. Unit: milliseconds.
	//
	// example:
	//
	// 1580886216000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The exchange type.
	//
	// example:
	//
	// DIRECT
	ExchangeType *string `json:"ExchangeType,omitempty" xml:"ExchangeType,omitempty"`
	// The exchange name.
	//
	// example:
	//
	// amq.direct
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The vhost name.
	//
	// example:
	//
	// test
	VHostName *string `json:"VHostName,omitempty" xml:"VHostName,omitempty"`
}

func (s ListExchangesResponseBodyDataExchanges) String() string {
	return tea.Prettify(s)
}

func (s ListExchangesResponseBodyDataExchanges) GoString() string {
	return s.String()
}

func (s *ListExchangesResponseBodyDataExchanges) SetAttributes(v map[string]interface{}) *ListExchangesResponseBodyDataExchanges {
	s.Attributes = v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetAutoDeleteState(v bool) *ListExchangesResponseBodyDataExchanges {
	s.AutoDeleteState = &v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetCreateTime(v int64) *ListExchangesResponseBodyDataExchanges {
	s.CreateTime = &v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetExchangeType(v string) *ListExchangesResponseBodyDataExchanges {
	s.ExchangeType = &v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetName(v string) *ListExchangesResponseBodyDataExchanges {
	s.Name = &v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetVHostName(v string) *ListExchangesResponseBodyDataExchanges {
	s.VHostName = &v
	return s
}

type ListExchangesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExchangesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExchangesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExchangesResponse) GoString() string {
	return s.String()
}

func (s *ListExchangesResponse) SetHeaders(v map[string]*string) *ListExchangesResponse {
	s.Headers = v
	return s
}

func (s *ListExchangesResponse) SetStatusCode(v int32) *ListExchangesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExchangesResponse) SetBody(v *ListExchangesResponseBody) *ListExchangesResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// The maximum number of entries to return. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end position of the previous returned page. To obtain the next batch of data, call the operation again by using the value of NextToken returned by the previous request. If you call this operation for the first time or want to query all results, set NextToken to an empty string.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetMaxResults(v int32) *ListInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesRequest) SetNextToken(v string) *ListInstancesRequest {
	s.NextToken = &v
	return s
}

type ListInstancesResponseBody struct {
	// The returned data.
	Data *ListInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CCBB1225-C392-480E-8C7F-D09AB2CD2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetData(v *ListInstancesResponseBodyData) *ListInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

type ListInstancesResponseBodyData struct {
	// The instances.
	Instances []*ListInstancesResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current returned page. If this parameter is empty, all data is retrieved.
	//
	// example:
	//
	// caebacccb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyData) SetInstances(v []*ListInstancesResponseBodyDataInstances) *ListInstancesResponseBodyData {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBodyData) SetMaxResults(v int32) *ListInstancesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetNextToken(v string) *ListInstancesResponseBodyData {
	s.NextToken = &v
	return s
}

type ListInstancesResponseBodyDataInstances struct {
	// Indicates whether the instance is automatically renewed.
	//
	// example:
	//
	// false
	AutoRenewInstance *bool `json:"AutoRenewInstance,omitempty" xml:"AutoRenewInstance,omitempty"`
	// The endpoint that is used to access the instance over the classic network. This parameter is no longer available.
	//
	// example:
	//
	// amqp-cn-st21x7kv****.not-support
	ClassicEndpoint *string `json:"ClassicEndpoint,omitempty" xml:"ClassicEndpoint,omitempty"`
	// The timestamp that indicates when the instance expires. Unit: milliseconds.
	//
	// example:
	//
	// 1651507200000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The instance ID
	//
	// example:
	//
	// amqp-cn-st21x7kv****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// amqp-cn-st21x7kv****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type.
	//
	// 	- PROFESSIONAL: Professional Edition
	//
	// 	- ENTERPRISE: Enterprise Edition
	//
	// 	- VIP: Enterprise Platinum Edition
	//
	// example:
	//
	// professional
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum number of Internet-based transactions per second (TPS) for the instance.
	//
	// example:
	//
	// 24832
	MaxEipTps *int32 `json:"MaxEipTps,omitempty" xml:"MaxEipTps,omitempty"`
	// The maximum number of queues on the instance.
	//
	// example:
	//
	// 50
	MaxQueue *int32 `json:"MaxQueue,omitempty" xml:"MaxQueue,omitempty"`
	// The maximum number of VPC-based TPS for the instance.
	//
	// example:
	//
	// 5000
	MaxTps *int32 `json:"MaxTps,omitempty" xml:"MaxTps,omitempty"`
	// The maximum number of vhosts on the instance.
	//
	// example:
	//
	// 50
	MaxVhost *int32 `json:"MaxVhost,omitempty" xml:"MaxVhost,omitempty"`
	// The timestamp that indicates when the order was created. Unit: milliseconds.
	//
	// example:
	//
	// 1572441939000
	OrderCreateTime *int64 `json:"OrderCreateTime,omitempty" xml:"OrderCreateTime,omitempty"`
	// The billing method. Valid values:
	//
	// 	- PrePaid: the subscription billing method.
	//
	// 	- POST_PAID: the pay-as-you-go billing method.
	//
	// example:
	//
	// PRE_PAID
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The virtual private cloud (VPC) endpoint of the instance.
	//
	// example:
	//
	// amqp-cn-st21x7kv****.mq-amqp.cn-hangzhou-a.aliyuncs.com
	PrivateEndpoint *string `json:"PrivateEndpoint,omitempty" xml:"PrivateEndpoint,omitempty"`
	// The public endpoint of the instance.
	//
	// example:
	//
	// amqp-cn-st21x7kv****.mq-amqp.cn-hangzhou-a.aliyuncs.com
	PublicEndpoint *string `json:"PublicEndpoint,omitempty" xml:"PublicEndpoint,omitempty"`
	// The instance status. Valid values:
	//
	// 	- DEPLOYING: The instance is being deployed.
	//
	// 	- EXPIRED: The instance is expired.
	//
	// 	- SERVING: The instance is running.
	//
	// 	- RELEASED: The instance is released.
	//
	// example:
	//
	// SERVING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The disk size. Unit: GB.
	//
	// >  For Professional Edition instances and Enterprise Edition instances, this parameter is unavailable and \\*\\*-1\\*\\	- is returned.
	//
	// example:
	//
	// 200
	StorageSize *int32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// Indicates whether the instance supports elastic IP addresses (EIPs).
	//
	// example:
	//
	// true
	SupportEIP *bool `json:"SupportEIP,omitempty" xml:"SupportEIP,omitempty"`
	// 标签列表。
	Tags []*ListInstancesResponseBodyDataInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyDataInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataInstances) SetAutoRenewInstance(v bool) *ListInstancesResponseBodyDataInstances {
	s.AutoRenewInstance = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetClassicEndpoint(v string) *ListInstancesResponseBodyDataInstances {
	s.ClassicEndpoint = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetExpireTime(v int64) *ListInstancesResponseBodyDataInstances {
	s.ExpireTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetInstanceId(v string) *ListInstancesResponseBodyDataInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetInstanceName(v string) *ListInstancesResponseBodyDataInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetInstanceType(v string) *ListInstancesResponseBodyDataInstances {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetMaxEipTps(v int32) *ListInstancesResponseBodyDataInstances {
	s.MaxEipTps = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetMaxQueue(v int32) *ListInstancesResponseBodyDataInstances {
	s.MaxQueue = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetMaxTps(v int32) *ListInstancesResponseBodyDataInstances {
	s.MaxTps = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetMaxVhost(v int32) *ListInstancesResponseBodyDataInstances {
	s.MaxVhost = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetOrderCreateTime(v int64) *ListInstancesResponseBodyDataInstances {
	s.OrderCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetOrderType(v string) *ListInstancesResponseBodyDataInstances {
	s.OrderType = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetPrivateEndpoint(v string) *ListInstancesResponseBodyDataInstances {
	s.PrivateEndpoint = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetPublicEndpoint(v string) *ListInstancesResponseBodyDataInstances {
	s.PublicEndpoint = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetStatus(v string) *ListInstancesResponseBodyDataInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetStorageSize(v int32) *ListInstancesResponseBodyDataInstances {
	s.StorageSize = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetSupportEIP(v bool) *ListInstancesResponseBodyDataInstances {
	s.SupportEIP = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetTags(v []*ListInstancesResponseBodyDataInstancesTags) *ListInstancesResponseBodyDataInstances {
	s.Tags = v
	return s
}

type ListInstancesResponseBodyDataInstancesTags struct {
	// 标签键。
	//
	// example:
	//
	// region
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值。
	//
	// example:
	//
	// hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesResponseBodyDataInstancesTags) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyDataInstancesTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataInstancesTags) SetKey(v string) *ListInstancesResponseBodyDataInstancesTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstancesTags) SetValue(v string) *ListInstancesResponseBodyDataInstancesTags {
	s.Value = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ListQueueConsumersRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 188077086902***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The token that marks the end position of the previous returned page. To obtain the next batch of data, call the operation again by using the value of NextToken returned by the previous request. If you call this operation for the first time or want to query all results, set NextToken to an empty string.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of data entries to return. If you do not configure this parameter, the default value 1 is used.
	//
	// Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	QueryCount *int32 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// The name of the queue for which you want to query online consumers.
	//
	// This parameter is required.
	//
	// example:
	//
	// queue-rabbit-springboot-advance5
	Queue *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	// The virtual host (vhost) name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s ListQueueConsumersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueueConsumersRequest) GoString() string {
	return s.String()
}

func (s *ListQueueConsumersRequest) SetInstanceId(v string) *ListQueueConsumersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListQueueConsumersRequest) SetNextToken(v string) *ListQueueConsumersRequest {
	s.NextToken = &v
	return s
}

func (s *ListQueueConsumersRequest) SetQueryCount(v int32) *ListQueueConsumersRequest {
	s.QueryCount = &v
	return s
}

func (s *ListQueueConsumersRequest) SetQueue(v string) *ListQueueConsumersRequest {
	s.Queue = &v
	return s
}

func (s *ListQueueConsumersRequest) SetVirtualHost(v string) *ListQueueConsumersRequest {
	s.VirtualHost = &v
	return s
}

type ListQueueConsumersResponseBody struct {
	// The returned data.
	Data *ListQueueConsumersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4409B7D5-E4EC-4EB5-804A-385DCDFCD***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListQueueConsumersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueueConsumersResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueueConsumersResponseBody) SetData(v *ListQueueConsumersResponseBodyData) *ListQueueConsumersResponseBody {
	s.Data = v
	return s
}

func (s *ListQueueConsumersResponseBody) SetRequestId(v string) *ListQueueConsumersResponseBody {
	s.RequestId = &v
	return s
}

type ListQueueConsumersResponseBodyData struct {
	// The consumers.
	Consumers []*ListQueueConsumersResponseBodyDataConsumers `json:"Consumers,omitempty" xml:"Consumers,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current returned page. If this parameter is empty, all data is retrieved.
	//
	// example:
	//
	// caebacccb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListQueueConsumersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQueueConsumersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQueueConsumersResponseBodyData) SetConsumers(v []*ListQueueConsumersResponseBodyDataConsumers) *ListQueueConsumersResponseBodyData {
	s.Consumers = v
	return s
}

func (s *ListQueueConsumersResponseBodyData) SetMaxResults(v int32) *ListQueueConsumersResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListQueueConsumersResponseBodyData) SetNextToken(v string) *ListQueueConsumersResponseBodyData {
	s.NextToken = &v
	return s
}

type ListQueueConsumersResponseBodyDataConsumers struct {
	// The consumer tag.
	//
	// example:
	//
	// sgen-1
	ConsumerTag *string `json:"ConsumerTag,omitempty" xml:"ConsumerTag,omitempty"`
}

func (s ListQueueConsumersResponseBodyDataConsumers) String() string {
	return tea.Prettify(s)
}

func (s ListQueueConsumersResponseBodyDataConsumers) GoString() string {
	return s.String()
}

func (s *ListQueueConsumersResponseBodyDataConsumers) SetConsumerTag(v string) *ListQueueConsumersResponseBodyDataConsumers {
	s.ConsumerTag = &v
	return s
}

type ListQueueConsumersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueueConsumersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQueueConsumersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueueConsumersResponse) GoString() string {
	return s.String()
}

func (s *ListQueueConsumersResponse) SetHeaders(v map[string]*string) *ListQueueConsumersResponse {
	s.Headers = v
	return s
}

func (s *ListQueueConsumersResponse) SetStatusCode(v int32) *ListQueueConsumersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueueConsumersResponse) SetBody(v *ListQueueConsumersResponseBody) *ListQueueConsumersResponse {
	s.Body = v
	return s
}

type ListQueueUpStreamBindingsRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1880770869023***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end position of the previous returned page. To obtain the next batch of data, call the operation again by using the value of NextToken returned by the previous request. If you call this operation for the first time or want to query all results, set NextToken to an empty string.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The queue name.
	//
	// This parameter is required.
	//
	// example:
	//
	// QueueTest
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The virtual host (vhost) name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s ListQueueUpStreamBindingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueueUpStreamBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsRequest) SetInstanceId(v string) *ListQueueUpStreamBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListQueueUpStreamBindingsRequest) SetMaxResults(v int32) *ListQueueUpStreamBindingsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQueueUpStreamBindingsRequest) SetNextToken(v string) *ListQueueUpStreamBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListQueueUpStreamBindingsRequest) SetQueueName(v string) *ListQueueUpStreamBindingsRequest {
	s.QueueName = &v
	return s
}

func (s *ListQueueUpStreamBindingsRequest) SetVirtualHost(v string) *ListQueueUpStreamBindingsRequest {
	s.VirtualHost = &v
	return s
}

type ListQueueUpStreamBindingsResponseBody struct {
	// The returned data.
	Data *ListQueueUpStreamBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8BFB1C9D-08A2-4859-A47C-403C9EFA2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListQueueUpStreamBindingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueueUpStreamBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsResponseBody) SetData(v *ListQueueUpStreamBindingsResponseBodyData) *ListQueueUpStreamBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBody) SetRequestId(v string) *ListQueueUpStreamBindingsResponseBody {
	s.RequestId = &v
	return s
}

type ListQueueUpStreamBindingsResponseBodyData struct {
	// The bindings.
	Bindings []*ListQueueUpStreamBindingsResponseBodyDataBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 1
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current returned page. If this parameter is empty, all data is retrieved.
	//
	// example:
	//
	// caebacccb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListQueueUpStreamBindingsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQueueUpStreamBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsResponseBodyData) SetBindings(v []*ListQueueUpStreamBindingsResponseBodyDataBindings) *ListQueueUpStreamBindingsResponseBodyData {
	s.Bindings = v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyData) SetMaxResults(v string) *ListQueueUpStreamBindingsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyData) SetNextToken(v string) *ListQueueUpStreamBindingsResponseBodyData {
	s.NextToken = &v
	return s
}

type ListQueueUpStreamBindingsResponseBodyDataBindings struct {
	// The x-match attribute. Valid values:
	//
	// 	- **all:*	- A headers exchange routes a message to a queue only if all binding attributes of the queue except for x-match match the headers attributes of the message. This value is the default value.
	//
	// 	- **any:*	- A headers exchange routes a message to a queue if one or more binding attributes of the queue except for x-match match the headers attributes of the message.
	//
	// This parameter is available for only headers exchanges.
	//
	// example:
	//
	// all
	Argument *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	// The binding key.
	//
	// 	- If the source exchange is not a topic exchange, the binding key must meet the following conventions:
	//
	//     	- The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), forward slashes (/), and at signs (@).
	//
	//     	- The binding key must be 1 to 255 characters in length.
	//
	// 	- If the source exchange is a topic exchange, the binding key must meet the following conventions:
	//
	//     	- The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@).
	//
	//     	- The binding key cannot start or end with a period (.). If a binding key starts with a number sign (#) or an asterisk (\\*), the number sign (#) or asterisk (\\*) must be followed by a period (.). If the binding key ends with a number sign (#) or an asterisk (\\*), the number sign (#) or asterisk (\\*) must be preceded by a period (.). If a number sign (#) or an asterisk (\\*) is used in the middle of a binding key, the number sign (#) or asterisk (\\*) must be preceded and followed by a period (.).
	//
	//     	- The binding key must be 1 to 255 characters in length.
	//
	// example:
	//
	// amq.test
	BindingKey *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	// The type of the object to which the source exchange is bound. Valid values:
	//
	// 	- **QUEUE**
	//
	// 	- **EXCHANGE**
	//
	// example:
	//
	// QUEUE
	BindingType *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	// The name of the object to which the source exchange is bound.
	//
	// example:
	//
	// QueueTest
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
	// The name of the source exchange.
	//
	// example:
	//
	// test
	SourceExchange *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
}

func (s ListQueueUpStreamBindingsResponseBodyDataBindings) String() string {
	return tea.Prettify(s)
}

func (s ListQueueUpStreamBindingsResponseBodyDataBindings) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetArgument(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.Argument = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetBindingKey(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.BindingKey = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetBindingType(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.BindingType = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetDestinationName(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.DestinationName = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetSourceExchange(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.SourceExchange = &v
	return s
}

type ListQueueUpStreamBindingsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueueUpStreamBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQueueUpStreamBindingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueueUpStreamBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsResponse) SetHeaders(v map[string]*string) *ListQueueUpStreamBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListQueueUpStreamBindingsResponse) SetStatusCode(v int32) *ListQueueUpStreamBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponse) SetBody(v *ListQueueUpStreamBindingsResponseBody) *ListQueueUpStreamBindingsResponse {
	s.Body = v
	return s
}

type ListQueuesRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1880770869023***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end position of the previous returned page. To obtain the next batch of data, call the operation again by using the value of NextToken returned by the previous request. If you call this operation for the first time or want to query all results, set NextToken to an empty string.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The virtual host (vhost) name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s ListQueuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesRequest) GoString() string {
	return s.String()
}

func (s *ListQueuesRequest) SetInstanceId(v string) *ListQueuesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListQueuesRequest) SetMaxResults(v int32) *ListQueuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQueuesRequest) SetNextToken(v string) *ListQueuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListQueuesRequest) SetVirtualHost(v string) *ListQueuesRequest {
	s.VirtualHost = &v
	return s
}

type ListQueuesResponseBody struct {
	// The returned data.
	Data *ListQueuesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CE811989-9F02-42CE-97A6-2239CB5C2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListQueuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBody) SetData(v *ListQueuesResponseBodyData) *ListQueuesResponseBody {
	s.Data = v
	return s
}

func (s *ListQueuesResponseBody) SetRequestId(v string) *ListQueuesResponseBody {
	s.RequestId = &v
	return s
}

type ListQueuesResponseBodyData struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current returned page. If this parameter is empty, all data is retrieved.
	//
	// example:
	//
	// caebacccb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The queues.
	Queues []*ListQueuesResponseBodyDataQueues `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
}

func (s ListQueuesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyData) SetMaxResults(v int32) *ListQueuesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListQueuesResponseBodyData) SetNextToken(v string) *ListQueuesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListQueuesResponseBodyData) SetQueues(v []*ListQueuesResponseBodyDataQueues) *ListQueuesResponseBodyData {
	s.Queues = v
	return s
}

type ListQueuesResponseBodyDataQueues struct {
	// The attributes.
	//
	// example:
	//
	// test
	Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// Indicates whether the queue was automatically deleted.
	//
	// example:
	//
	// false
	AutoDeleteState *bool `json:"AutoDeleteState,omitempty" xml:"AutoDeleteState,omitempty"`
	// The time when the queue was created.
	//
	// example:
	//
	// 1580887085240
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the queue is an exclusive queue.
	//
	// example:
	//
	// false
	ExclusiveState *bool `json:"ExclusiveState,omitempty" xml:"ExclusiveState,omitempty"`
	// The time when messages in the queue were last consumed.
	//
	// example:
	//
	// 1680887085240
	LastConsumeTime *int64 `json:"LastConsumeTime,omitempty" xml:"LastConsumeTime,omitempty"`
	// The queue name.
	//
	// example:
	//
	// QueueTest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance to which the queue belongs.
	//
	// example:
	//
	// 1880770869023***
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The vhost name.
	//
	// example:
	//
	// test
	VHostName *string `json:"VHostName,omitempty" xml:"VHostName,omitempty"`
}

func (s ListQueuesResponseBodyDataQueues) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyDataQueues) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyDataQueues) SetAttributes(v map[string]interface{}) *ListQueuesResponseBodyDataQueues {
	s.Attributes = v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetAutoDeleteState(v bool) *ListQueuesResponseBodyDataQueues {
	s.AutoDeleteState = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetCreateTime(v int64) *ListQueuesResponseBodyDataQueues {
	s.CreateTime = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetExclusiveState(v bool) *ListQueuesResponseBodyDataQueues {
	s.ExclusiveState = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetLastConsumeTime(v int64) *ListQueuesResponseBodyDataQueues {
	s.LastConsumeTime = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetName(v string) *ListQueuesResponseBodyDataQueues {
	s.Name = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetOwnerId(v string) *ListQueuesResponseBodyDataQueues {
	s.OwnerId = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetVHostName(v string) *ListQueuesResponseBodyDataQueues {
	s.VHostName = &v
	return s
}

type ListQueuesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQueuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponse) GoString() string {
	return s.String()
}

func (s *ListQueuesResponse) SetHeaders(v map[string]*string) *ListQueuesResponse {
	s.Headers = v
	return s
}

func (s *ListQueuesResponse) SetStatusCode(v int32) *ListQueuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueuesResponse) SetBody(v *ListQueuesResponseBody) *ListQueuesResponse {
	s.Body = v
	return s
}

type ListVirtualHostsRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1880770869023***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return. Valid values: **1 to 100**
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end position of the previous returned page. To obtain the next batch of data, call the operation again by using the value of NextToken returned by the previous request. If you call this operation for the first time or want to query all results, set NextToken to an empty string.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListVirtualHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualHostsRequest) GoString() string {
	return s.String()
}

func (s *ListVirtualHostsRequest) SetInstanceId(v string) *ListVirtualHostsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVirtualHostsRequest) SetMaxResults(v int32) *ListVirtualHostsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVirtualHostsRequest) SetNextToken(v string) *ListVirtualHostsRequest {
	s.NextToken = &v
	return s
}

type ListVirtualHostsResponseBody struct {
	// The returned data.
	Data *ListVirtualHostsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EF4DB019-DA4A-4CE3-B220-223BBC93F***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVirtualHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualHostsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirtualHostsResponseBody) SetData(v *ListVirtualHostsResponseBodyData) *ListVirtualHostsResponseBody {
	s.Data = v
	return s
}

func (s *ListVirtualHostsResponseBody) SetRequestId(v string) *ListVirtualHostsResponseBody {
	s.RequestId = &v
	return s
}

type ListVirtualHostsResponseBodyData struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 2
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current returned page. If this parameter is empty, all data is retrieved.
	//
	// example:
	//
	// caebacccb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The vhosts.
	VirtualHosts []*ListVirtualHostsResponseBodyDataVirtualHosts `json:"VirtualHosts,omitempty" xml:"VirtualHosts,omitempty" type:"Repeated"`
}

func (s ListVirtualHostsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualHostsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVirtualHostsResponseBodyData) SetMaxResults(v int32) *ListVirtualHostsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListVirtualHostsResponseBodyData) SetNextToken(v string) *ListVirtualHostsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListVirtualHostsResponseBodyData) SetVirtualHosts(v []*ListVirtualHostsResponseBodyDataVirtualHosts) *ListVirtualHostsResponseBodyData {
	s.VirtualHosts = v
	return s
}

type ListVirtualHostsResponseBodyDataVirtualHosts struct {
	// The vhost name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListVirtualHostsResponseBodyDataVirtualHosts) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualHostsResponseBodyDataVirtualHosts) GoString() string {
	return s.String()
}

func (s *ListVirtualHostsResponseBodyDataVirtualHosts) SetName(v string) *ListVirtualHostsResponseBodyDataVirtualHosts {
	s.Name = &v
	return s
}

type ListVirtualHostsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirtualHostsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirtualHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualHostsResponse) GoString() string {
	return s.String()
}

func (s *ListVirtualHostsResponse) SetHeaders(v map[string]*string) *ListVirtualHostsResponse {
	s.Headers = v
	return s
}

func (s *ListVirtualHostsResponse) SetStatusCode(v int32) *ListVirtualHostsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirtualHostsResponse) SetBody(v *ListVirtualHostsResponseBody) *ListVirtualHostsResponse {
	s.Body = v
	return s
}

type UpdateInstanceRequest struct {
	// example:
	//
	// c2c5d1274axxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-jtexxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// professional
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 1000
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// example:
	//
	// 128
	MaxEipTps *int64 `json:"MaxEipTps,omitempty" xml:"MaxEipTps,omitempty"`
	// example:
	//
	// 1000
	MaxPrivateTps *int64 `json:"MaxPrivateTps,omitempty" xml:"MaxPrivateTps,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UPGRADE
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// example:
	//
	// 1000
	QueueCapacity *int32 `json:"QueueCapacity,omitempty" xml:"QueueCapacity,omitempty"`
	// example:
	//
	// onDemand
	ServerlessChargeType *string `json:"ServerlessChargeType,omitempty" xml:"ServerlessChargeType,omitempty"`
	// example:
	//
	// 7
	StorageSize *int32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// example:
	//
	// false
	SupportEip *bool `json:"SupportEip,omitempty" xml:"SupportEip,omitempty"`
	// example:
	//
	// false
	SupportTracing *bool `json:"SupportTracing,omitempty" xml:"SupportTracing,omitempty"`
	// example:
	//
	// 3
	TracingStorageTime *int32 `json:"TracingStorageTime,omitempty" xml:"TracingStorageTime,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) SetClientToken(v string) *UpdateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceId(v string) *UpdateInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceType(v string) *UpdateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *UpdateInstanceRequest) SetMaxConnections(v int32) *UpdateInstanceRequest {
	s.MaxConnections = &v
	return s
}

func (s *UpdateInstanceRequest) SetMaxEipTps(v int64) *UpdateInstanceRequest {
	s.MaxEipTps = &v
	return s
}

func (s *UpdateInstanceRequest) SetMaxPrivateTps(v int64) *UpdateInstanceRequest {
	s.MaxPrivateTps = &v
	return s
}

func (s *UpdateInstanceRequest) SetModifyType(v string) *UpdateInstanceRequest {
	s.ModifyType = &v
	return s
}

func (s *UpdateInstanceRequest) SetQueueCapacity(v int32) *UpdateInstanceRequest {
	s.QueueCapacity = &v
	return s
}

func (s *UpdateInstanceRequest) SetServerlessChargeType(v string) *UpdateInstanceRequest {
	s.ServerlessChargeType = &v
	return s
}

func (s *UpdateInstanceRequest) SetStorageSize(v int32) *UpdateInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *UpdateInstanceRequest) SetSupportEip(v bool) *UpdateInstanceRequest {
	s.SupportEip = &v
	return s
}

func (s *UpdateInstanceRequest) SetSupportTracing(v bool) *UpdateInstanceRequest {
	s.SupportTracing = &v
	return s
}

func (s *UpdateInstanceRequest) SetTracingStorageTime(v int32) *UpdateInstanceRequest {
	s.TracingStorageTime = &v
	return s
}

type UpdateInstanceResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {“instanceId”: “amqp-cn-jtexxxxx”, “orderId”: 2222222}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// InstanceNotExist
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 628705FD-03EE-4ABE-BB21-E1672960***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) SetCode(v int32) *UpdateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetData(v interface{}) *UpdateInstanceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateInstanceResponseBody) SetMessage(v string) *UpdateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetStatusCode(v string) *UpdateInstanceResponseBody {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetSuccess(v string) *UpdateInstanceResponseBody {
	s.Success = &v
	return s
}

type UpdateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponse) SetHeaders(v map[string]*string) *UpdateInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceResponse) SetStatusCode(v int32) *UpdateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceResponse) SetBody(v *UpdateInstanceResponseBody) *UpdateInstanceResponse {
	s.Body = v
	return s
}

type UpdateInstanceNameRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance for which you want to update the name.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-zvp2ajsj****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new name of the instance. No limits are imposed on the value. We recommend that you set this parameter to a maximum of 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-ZVp2ajsj****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s UpdateInstanceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameRequest) SetInstanceId(v string) *UpdateInstanceNameRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceNameRequest) SetInstanceName(v string) *UpdateInstanceNameRequest {
	s.InstanceName = &v
	return s
}

type UpdateInstanceNameResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message that is returned when an error occurs during the update of the instance name.
	//
	// example:
	//
	// InstanceNotExist
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6DC68EC9-0E76-5435-B8C0-FF9492B4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned message that indicates the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameResponseBody) SetCode(v int32) *UpdateInstanceNameResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetData(v string) *UpdateInstanceNameResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetMessage(v string) *UpdateInstanceNameResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetRequestId(v string) *UpdateInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetSuccess(v bool) *UpdateInstanceNameResponseBody {
	s.Success = &v
	return s
}

type UpdateInstanceNameResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameResponse) SetHeaders(v map[string]*string) *UpdateInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceNameResponse) SetStatusCode(v int32) *UpdateInstanceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceNameResponse) SetBody(v *UpdateInstanceNameResponseBody) *UpdateInstanceNameResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("amqp-open"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a pair of static username and password. If you access an ApsaraMQ for RabbitMQ broker from an open source RabbitMQ client, you must use a pair of username and password for authentication. You can access the ApsaraMQ for RabbitMQ broker only after the authentication is passed. ApsaraMQ for RabbitMQ allows you to generate usernames and passwords by using AccessKey pairs provided by Alibaba Cloud Resource Access Management (RAM).
//
// @param request - CreateAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccountResponse
func (client *Client) CreateAccountWithOptions(request *CreateAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountAccessKey)) {
		query["accountAccessKey"] = request.AccountAccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimestamp)) {
		query["createTimestamp"] = request.CreateTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretSign)) {
		query["secretSign"] = request.SecretSign
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["userName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccount"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a pair of static username and password. If you access an ApsaraMQ for RabbitMQ broker from an open source RabbitMQ client, you must use a pair of username and password for authentication. You can access the ApsaraMQ for RabbitMQ broker only after the authentication is passed. ApsaraMQ for RabbitMQ allows you to generate usernames and passwords by using AccessKey pairs provided by Alibaba Cloud Resource Access Management (RAM).
//
// @param request - CreateAccountRequest
//
// @return CreateAccountResponse
func (client *Client) CreateAccount(request *CreateAccountRequest) (_result *CreateAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccountResponse{}
	_body, _err := client.CreateAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a binding. In ApsaraMQ for RabbitMQ, after a producer sends a message to an exchange, the exchange routes the message to a queue or another exchange based on the binding relationship and the routing rule.
//
// @param request - CreateBindingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBindingResponse
func (client *Client) CreateBindingWithOptions(request *CreateBindingRequest, runtime *util.RuntimeOptions) (_result *CreateBindingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Argument)) {
		body["Argument"] = request.Argument
	}

	if !tea.BoolValue(util.IsUnset(request.BindingKey)) {
		body["BindingKey"] = request.BindingKey
	}

	if !tea.BoolValue(util.IsUnset(request.BindingType)) {
		body["BindingType"] = request.BindingType
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationName)) {
		body["DestinationName"] = request.DestinationName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceExchange)) {
		body["SourceExchange"] = request.SourceExchange
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualHost)) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBinding"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a binding. In ApsaraMQ for RabbitMQ, after a producer sends a message to an exchange, the exchange routes the message to a queue or another exchange based on the binding relationship and the routing rule.
//
// @param request - CreateBindingRequest
//
// @return CreateBindingResponse
func (client *Client) CreateBinding(request *CreateBindingRequest) (_result *CreateBindingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBindingResponse{}
	_body, _err := client.CreateBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an exchange. In ApsaraMQ for RabbitMQ, an exchange is used to route a message that is received from a producer to one or more queues or to discard the message. An exchange routes a message to queues by using the routing key and binding keys.
//
// @param request - CreateExchangeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExchangeResponse
func (client *Client) CreateExchangeWithOptions(request *CreateExchangeRequest, runtime *util.RuntimeOptions) (_result *CreateExchangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlternateExchange)) {
		body["AlternateExchange"] = request.AlternateExchange
	}

	if !tea.BoolValue(util.IsUnset(request.AutoDeleteState)) {
		body["AutoDeleteState"] = request.AutoDeleteState
	}

	if !tea.BoolValue(util.IsUnset(request.ExchangeName)) {
		body["ExchangeName"] = request.ExchangeName
	}

	if !tea.BoolValue(util.IsUnset(request.ExchangeType)) {
		body["ExchangeType"] = request.ExchangeType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Internal)) {
		body["Internal"] = request.Internal
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualHost)) {
		body["VirtualHost"] = request.VirtualHost
	}

	if !tea.BoolValue(util.IsUnset(request.XDelayedType)) {
		body["XDelayedType"] = request.XDelayedType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExchange"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateExchangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an exchange. In ApsaraMQ for RabbitMQ, an exchange is used to route a message that is received from a producer to one or more queues or to discard the message. An exchange routes a message to queues by using the routing key and binding keys.
//
// @param request - CreateExchangeRequest
//
// @return CreateExchangeResponse
func (client *Client) CreateExchange(request *CreateExchangeRequest) (_result *CreateExchangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateExchangeResponse{}
	_body, _err := client.CreateExchangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建实例-基于 openAPI 构建南天门购买工单信息数据
//
// @param request - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewPeriod)) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxConnections)) {
		query["MaxConnections"] = request.MaxConnections
	}

	if !tea.BoolValue(util.IsUnset(request.MaxEipTps)) {
		query["MaxEipTps"] = request.MaxEipTps
	}

	if !tea.BoolValue(util.IsUnset(request.MaxPrivateTps)) {
		query["MaxPrivateTps"] = request.MaxPrivateTps
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		query["PaymentType"] = request.PaymentType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodCycle)) {
		query["PeriodCycle"] = request.PeriodCycle
	}

	if !tea.BoolValue(util.IsUnset(request.QueueCapacity)) {
		query["QueueCapacity"] = request.QueueCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.RenewStatus)) {
		query["RenewStatus"] = request.RenewStatus
	}

	if !tea.BoolValue(util.IsUnset(request.RenewalDurationUnit)) {
		query["RenewalDurationUnit"] = request.RenewalDurationUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ServerlessChargeType)) {
		query["ServerlessChargeType"] = request.ServerlessChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.StorageSize)) {
		query["StorageSize"] = request.StorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SupportEip)) {
		query["SupportEip"] = request.SupportEip
	}

	if !tea.BoolValue(util.IsUnset(request.SupportTracing)) {
		query["SupportTracing"] = request.SupportTracing
	}

	if !tea.BoolValue(util.IsUnset(request.TracingStorageTime)) {
		query["TracingStorageTime"] = request.TracingStorageTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实例-基于 openAPI 构建南天门购买工单信息数据
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a queue. In ApsaraMQ for RabbitMQ, a queue is a message queue. All messages in ApsaraMQ for RabbitMQ are sent to a specific exchange and then routed to a bound queue by the exchange.
//
// @param request - CreateQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQueueResponse
func (client *Client) CreateQueueWithOptions(request *CreateQueueRequest, runtime *util.RuntimeOptions) (_result *CreateQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoDeleteState)) {
		body["AutoDeleteState"] = request.AutoDeleteState
	}

	if !tea.BoolValue(util.IsUnset(request.AutoExpireState)) {
		body["AutoExpireState"] = request.AutoExpireState
	}

	if !tea.BoolValue(util.IsUnset(request.DeadLetterExchange)) {
		body["DeadLetterExchange"] = request.DeadLetterExchange
	}

	if !tea.BoolValue(util.IsUnset(request.DeadLetterRoutingKey)) {
		body["DeadLetterRoutingKey"] = request.DeadLetterRoutingKey
	}

	if !tea.BoolValue(util.IsUnset(request.ExclusiveState)) {
		body["ExclusiveState"] = request.ExclusiveState
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxLength)) {
		body["MaxLength"] = request.MaxLength
	}

	if !tea.BoolValue(util.IsUnset(request.MaximumPriority)) {
		body["MaximumPriority"] = request.MaximumPriority
	}

	if !tea.BoolValue(util.IsUnset(request.MessageTTL)) {
		body["MessageTTL"] = request.MessageTTL
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		body["QueueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualHost)) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateQueue"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a queue. In ApsaraMQ for RabbitMQ, a queue is a message queue. All messages in ApsaraMQ for RabbitMQ are sent to a specific exchange and then routed to a bound queue by the exchange.
//
// @param request - CreateQueueRequest
//
// @return CreateQueueResponse
func (client *Client) CreateQueue(request *CreateQueueRequest) (_result *CreateQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateQueueResponse{}
	_body, _err := client.CreateQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a vhost. A vhost is used to logically isolate resources. Each vhost manages its own exchanges, queues, and bindings. Applications can run on independent vhosts in a secure manner. This way, the business of an application is not affected by other applications. Before you connect producers and consumers to an ApsaraMQ for RabbitMQ instance, you must specify vhosts for the producers and consumers.
//
// @param request - CreateVirtualHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVirtualHostResponse
func (client *Client) CreateVirtualHostWithOptions(request *CreateVirtualHostRequest, runtime *util.RuntimeOptions) (_result *CreateVirtualHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualHost)) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVirtualHost"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVirtualHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a vhost. A vhost is used to logically isolate resources. Each vhost manages its own exchanges, queues, and bindings. Applications can run on independent vhosts in a secure manner. This way, the business of an application is not affected by other applications. Before you connect producers and consumers to an ApsaraMQ for RabbitMQ instance, you must specify vhosts for the producers and consumers.
//
// @param request - CreateVirtualHostRequest
//
// @return CreateVirtualHostResponse
func (client *Client) CreateVirtualHost(request *CreateVirtualHostRequest) (_result *CreateVirtualHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVirtualHostResponse{}
	_body, _err := client.CreateVirtualHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a pair of username and password.
//
// @param request - DeleteAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccountWithOptions(request *DeleteAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateTimestamp)) {
		query["CreateTimestamp"] = request.CreateTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccount"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a pair of username and password.
//
// @param request - DeleteAccountRequest
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccount(request *DeleteAccountRequest) (_result *DeleteAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DeleteAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a binding to unbind a queue or an exchange from a source exchange.
//
// @param request - DeleteBindingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBindingResponse
func (client *Client) DeleteBindingWithOptions(request *DeleteBindingRequest, runtime *util.RuntimeOptions) (_result *DeleteBindingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindingKey)) {
		body["BindingKey"] = request.BindingKey
	}

	if !tea.BoolValue(util.IsUnset(request.BindingType)) {
		body["BindingType"] = request.BindingType
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationName)) {
		body["DestinationName"] = request.DestinationName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceExchange)) {
		body["SourceExchange"] = request.SourceExchange
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualHost)) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBinding"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a binding to unbind a queue or an exchange from a source exchange.
//
// @param request - DeleteBindingRequest
//
// @return DeleteBindingResponse
func (client *Client) DeleteBinding(request *DeleteBindingRequest) (_result *DeleteBindingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBindingResponse{}
	_body, _err := client.DeleteBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an exchange.
//
// Description:
//
// ## [](#)Usage notes
//
// 	- You cannot delete exchanges of the **headers*	- and **x-jms-topic*	- types.
//
// 	- You cannot delete built-in exchanges in a vhost. These exchanges are amq.direct, amq.topic, and amq.fanout.
//
// @param request - DeleteExchangeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExchangeResponse
func (client *Client) DeleteExchangeWithOptions(request *DeleteExchangeRequest, runtime *util.RuntimeOptions) (_result *DeleteExchangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExchangeName)) {
		body["ExchangeName"] = request.ExchangeName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualHost)) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExchange"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteExchangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an exchange.
//
// Description:
//
// ## [](#)Usage notes
//
// 	- You cannot delete exchanges of the **headers*	- and **x-jms-topic*	- types.
//
// 	- You cannot delete built-in exchanges in a vhost. These exchanges are amq.direct, amq.topic, and amq.fanout.
//
// @param request - DeleteExchangeRequest
//
// @return DeleteExchangeResponse
func (client *Client) DeleteExchange(request *DeleteExchangeRequest) (_result *DeleteExchangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteExchangeResponse{}
	_body, _err := client.DeleteExchangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a queue.
//
// @param request - DeleteQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQueueResponse
func (client *Client) DeleteQueueWithOptions(request *DeleteQueueRequest, runtime *util.RuntimeOptions) (_result *DeleteQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		body["QueueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualHost)) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteQueue"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a queue.
//
// @param request - DeleteQueueRequest
//
// @return DeleteQueueResponse
func (client *Client) DeleteQueue(request *DeleteQueueRequest) (_result *DeleteQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteQueueResponse{}
	_body, _err := client.DeleteQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a virtual host (vhost).
//
// Description:
//
// Before you delete a vhost, make sure that all exchanges and queues in the vhost are deleted.
//
// @param request - DeleteVirtualHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVirtualHostResponse
func (client *Client) DeleteVirtualHostWithOptions(request *DeleteVirtualHostRequest, runtime *util.RuntimeOptions) (_result *DeleteVirtualHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualHost)) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVirtualHost"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVirtualHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a virtual host (vhost).
//
// Description:
//
// Before you delete a vhost, make sure that all exchanges and queues in the vhost are deleted.
//
// @param request - DeleteVirtualHostRequest
//
// @return DeleteVirtualHostResponse
func (client *Client) DeleteVirtualHost(request *DeleteVirtualHostRequest) (_result *DeleteVirtualHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVirtualHostResponse{}
	_body, _err := client.DeleteVirtualHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the maximum number of vhosts, exchanges, and queues that you can create and the number of created vhosts, exchanges, and queues on an ApsaraMQ for RabbitMQ instance.
//
// @param request - GetMetadataAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetadataAmountResponse
func (client *Client) GetMetadataAmountWithOptions(request *GetMetadataAmountRequest, runtime *util.RuntimeOptions) (_result *GetMetadataAmountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMetadataAmount"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMetadataAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the maximum number of vhosts, exchanges, and queues that you can create and the number of created vhosts, exchanges, and queues on an ApsaraMQ for RabbitMQ instance.
//
// @param request - GetMetadataAmountRequest
//
// @return GetMetadataAmountResponse
func (client *Client) GetMetadataAmount(request *GetMetadataAmountRequest) (_result *GetMetadataAmountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMetadataAmountResponse{}
	_body, _err := client.GetMetadataAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the static username and password of an ApsaraMQ for RabbitMQ.
//
// @param request - ListAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountsResponse
func (client *Client) ListAccountsWithOptions(request *ListAccountsRequest, runtime *util.RuntimeOptions) (_result *ListAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccounts"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the static username and password of an ApsaraMQ for RabbitMQ.
//
// @param request - ListAccountsRequest
//
// @return ListAccountsResponse
func (client *Client) ListAccounts(request *ListAccountsRequest) (_result *ListAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccountsResponse{}
	_body, _err := client.ListAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all bindings of a virtual host (vhost) on an ApsaraMQ for RabbitMQ instance.
//
// @param request - ListBindingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBindingsResponse
func (client *Client) ListBindingsWithOptions(request *ListBindingsRequest, runtime *util.RuntimeOptions) (_result *ListBindingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBindings"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all bindings of a virtual host (vhost) on an ApsaraMQ for RabbitMQ instance.
//
// @param request - ListBindingsRequest
//
// @return ListBindingsResponse
func (client *Client) ListBindings(request *ListBindingsRequest) (_result *ListBindingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBindingsResponse{}
	_body, _err := client.ListBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all exchanges or queues to which an exchange is bound.
//
// @param request - ListDownStreamBindingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDownStreamBindingsResponse
func (client *Client) ListDownStreamBindingsWithOptions(request *ListDownStreamBindingsRequest, runtime *util.RuntimeOptions) (_result *ListDownStreamBindingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDownStreamBindings"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDownStreamBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all exchanges or queues to which an exchange is bound.
//
// @param request - ListDownStreamBindingsRequest
//
// @return ListDownStreamBindingsResponse
func (client *Client) ListDownStreamBindings(request *ListDownStreamBindingsRequest) (_result *ListDownStreamBindingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDownStreamBindingsResponse{}
	_body, _err := client.ListDownStreamBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all queues or exchanges that are bound to an exchange.
//
// @param request - ListExchangeUpStreamBindingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExchangeUpStreamBindingsResponse
func (client *Client) ListExchangeUpStreamBindingsWithOptions(request *ListExchangeUpStreamBindingsRequest, runtime *util.RuntimeOptions) (_result *ListExchangeUpStreamBindingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExchangeUpStreamBindings"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExchangeUpStreamBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all queues or exchanges that are bound to an exchange.
//
// @param request - ListExchangeUpStreamBindingsRequest
//
// @return ListExchangeUpStreamBindingsResponse
func (client *Client) ListExchangeUpStreamBindings(request *ListExchangeUpStreamBindingsRequest) (_result *ListExchangeUpStreamBindingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExchangeUpStreamBindingsResponse{}
	_body, _err := client.ListExchangeUpStreamBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all exchanges that are created in a virtual host (vhost).
//
// @param request - ListExchangesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExchangesResponse
func (client *Client) ListExchangesWithOptions(request *ListExchangesRequest, runtime *util.RuntimeOptions) (_result *ListExchangesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExchanges"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExchangesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all exchanges that are created in a virtual host (vhost).
//
// @param request - ListExchangesRequest
//
// @return ListExchangesResponse
func (client *Client) ListExchanges(request *ListExchangesRequest) (_result *ListExchangesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExchangesResponse{}
	_body, _err := client.ListExchangesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all AparaMQ for RabbitMQ instances in a region. The returned data includes the basic information, endpoint, and specification limits of each instance.
//
// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all AparaMQ for RabbitMQ instances in a region. The returned data includes the basic information, endpoint, and specification limits of each instance.
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the online consumers of a queue.
//
// Description:
//
// ApsaraMQ for RabbitMQ allows you to query only online consumers.
//
// @param request - ListQueueConsumersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueueConsumersResponse
func (client *Client) ListQueueConsumersWithOptions(request *ListQueueConsumersRequest, runtime *util.RuntimeOptions) (_result *ListQueueConsumersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueueConsumers"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueueConsumersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the online consumers of a queue.
//
// Description:
//
// ApsaraMQ for RabbitMQ allows you to query only online consumers.
//
// @param request - ListQueueConsumersRequest
//
// @return ListQueueConsumersResponse
func (client *Client) ListQueueConsumers(request *ListQueueConsumersRequest) (_result *ListQueueConsumersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQueueConsumersResponse{}
	_body, _err := client.ListQueueConsumersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the exchanges that are bound to a queue.
//
// @param request - ListQueueUpStreamBindingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueueUpStreamBindingsResponse
func (client *Client) ListQueueUpStreamBindingsWithOptions(request *ListQueueUpStreamBindingsRequest, runtime *util.RuntimeOptions) (_result *ListQueueUpStreamBindingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueueUpStreamBindings"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueueUpStreamBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the exchanges that are bound to a queue.
//
// @param request - ListQueueUpStreamBindingsRequest
//
// @return ListQueueUpStreamBindingsResponse
func (client *Client) ListQueueUpStreamBindings(request *ListQueueUpStreamBindingsRequest) (_result *ListQueueUpStreamBindingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQueueUpStreamBindingsResponse{}
	_body, _err := client.ListQueueUpStreamBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all queues in a vhost of an ApsaraMQ for RabbitMQ instance.
//
// @param request - ListQueuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueuesResponse
func (client *Client) ListQueuesWithOptions(request *ListQueuesRequest, runtime *util.RuntimeOptions) (_result *ListQueuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueues"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all queues in a vhost of an ApsaraMQ for RabbitMQ instance.
//
// @param request - ListQueuesRequest
//
// @return ListQueuesResponse
func (client *Client) ListQueues(request *ListQueuesRequest) (_result *ListQueuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQueuesResponse{}
	_body, _err := client.ListQueuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all virtual hosts (vhosts) on an ApsaraMQ for RabbitMQ instance.
//
// @param request - ListVirtualHostsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVirtualHostsResponse
func (client *Client) ListVirtualHostsWithOptions(request *ListVirtualHostsRequest, runtime *util.RuntimeOptions) (_result *ListVirtualHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVirtualHosts"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVirtualHostsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all virtual hosts (vhosts) on an ApsaraMQ for RabbitMQ instance.
//
// @param request - ListVirtualHostsRequest
//
// @return ListVirtualHostsResponse
func (client *Client) ListVirtualHosts(request *ListVirtualHostsRequest) (_result *ListVirtualHostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVirtualHostsResponse{}
	_body, _err := client.ListVirtualHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变更实例，升降配
//
// @param request - UpdateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithOptions(request *UpdateInstanceRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxConnections)) {
		query["MaxConnections"] = request.MaxConnections
	}

	if !tea.BoolValue(util.IsUnset(request.MaxEipTps)) {
		query["MaxEipTps"] = request.MaxEipTps
	}

	if !tea.BoolValue(util.IsUnset(request.MaxPrivateTps)) {
		query["MaxPrivateTps"] = request.MaxPrivateTps
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyType)) {
		query["ModifyType"] = request.ModifyType
	}

	if !tea.BoolValue(util.IsUnset(request.QueueCapacity)) {
		query["QueueCapacity"] = request.QueueCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.ServerlessChargeType)) {
		query["ServerlessChargeType"] = request.ServerlessChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.StorageSize)) {
		query["StorageSize"] = request.StorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SupportEip)) {
		query["SupportEip"] = request.SupportEip
	}

	if !tea.BoolValue(util.IsUnset(request.SupportTracing)) {
		query["SupportTracing"] = request.SupportTracing
	}

	if !tea.BoolValue(util.IsUnset(request.TracingStorageTime)) {
		query["TracingStorageTime"] = request.TracingStorageTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstance"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更实例，升降配
//
// @param request - UpdateInstanceRequest
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstance(request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the name of an ApsaraMQ for RabbitMQ instance. After an ApsaraMQ for RabbitMQ instance is created, the ID of the instance is used as its name by default. You can specify a custom name for an instance to facilitate instance identification.
//
// @param request - UpdateInstanceNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceNameResponse
func (client *Client) UpdateInstanceNameWithOptions(request *UpdateInstanceNameRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceName"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name of an ApsaraMQ for RabbitMQ instance. After an ApsaraMQ for RabbitMQ instance is created, the ID of the instance is used as its name by default. You can specify a custom name for an instance to facilitate instance identification.
//
// @param request - UpdateInstanceNameRequest
//
// @return UpdateInstanceNameResponse
func (client *Client) UpdateInstanceName(request *UpdateInstanceNameRequest) (_result *UpdateInstanceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceNameResponse{}
	_body, _err := client.UpdateInstanceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
