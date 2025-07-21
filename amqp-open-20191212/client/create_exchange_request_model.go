// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExchangeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlternateExchange(v string) *CreateExchangeRequest
	GetAlternateExchange() *string
	SetAutoDeleteState(v bool) *CreateExchangeRequest
	GetAutoDeleteState() *bool
	SetExchangeName(v string) *CreateExchangeRequest
	GetExchangeName() *string
	SetExchangeType(v string) *CreateExchangeRequest
	GetExchangeType() *string
	SetInstanceId(v string) *CreateExchangeRequest
	GetInstanceId() *string
	SetInternal(v bool) *CreateExchangeRequest
	GetInternal() *bool
	SetVirtualHost(v string) *CreateExchangeRequest
	GetVirtualHost() *string
	SetXDelayedType(v string) *CreateExchangeRequest
	GetXDelayedType() *string
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
	return dara.Prettify(s)
}

func (s CreateExchangeRequest) GoString() string {
	return s.String()
}

func (s *CreateExchangeRequest) GetAlternateExchange() *string {
	return s.AlternateExchange
}

func (s *CreateExchangeRequest) GetAutoDeleteState() *bool {
	return s.AutoDeleteState
}

func (s *CreateExchangeRequest) GetExchangeName() *string {
	return s.ExchangeName
}

func (s *CreateExchangeRequest) GetExchangeType() *string {
	return s.ExchangeType
}

func (s *CreateExchangeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateExchangeRequest) GetInternal() *bool {
	return s.Internal
}

func (s *CreateExchangeRequest) GetVirtualHost() *string {
	return s.VirtualHost
}

func (s *CreateExchangeRequest) GetXDelayedType() *string {
	return s.XDelayedType
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

func (s *CreateExchangeRequest) Validate() error {
	return dara.Validate(s)
}
