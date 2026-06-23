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
	// The alternate exchange. Configure an alternate exchange to receive messages that fail to be routed.
	//
	// example:
	//
	// DemoAE
	AlternateExchange *string `json:"AlternateExchange,omitempty" xml:"AlternateExchange,omitempty"`
	// Specifies whether to automatically delete the exchange. Valid values:
	//
	// - **true**: Yes. The exchange is automatically deleted after the last queue is unbound from it.
	//
	// - **false**: No. The exchange is not automatically deleted after the last queue is unbound from it.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	AutoDeleteState *bool `json:"AutoDeleteState,omitempty" xml:"AutoDeleteState,omitempty"`
	// The name of the exchange. Note:
	//
	// - The name can contain only letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@). The name must be 1 to 255 characters in length.
	//
	// - The name of an exchange cannot be changed after the exchange is created. To change the name, delete the exchange and create a new one.
	//
	// This parameter is required.
	//
	// example:
	//
	// DemoExchange
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	// The type of the exchange. Valid values:
	//
	// - **DIRECT**: This routing rule type routes messages to a queue whose binding key exactly matches the routing key of the message.
	//
	// - **TOPIC**: This type is similar to the DIRECT type. It routes messages to bound queues using routing key pattern matching and string comparison.
	//
	// - **FANOUT**: This routing rule type is simple. It routes all messages sent to the exchange to all queues that are bound to the exchange. This works like a broadcast feature.
	//
	// - **HEADERS**: This type is similar to the DIRECT type. It uses header properties instead of a routing key for routing. When a queue is bound to a headers exchange, key-value pairs are defined for the binding. When a message is sent to the exchange, key-value pairs are defined in the message header. The exchange routes the message by comparing the key-value pairs in the header with the key-value pairs of the binding.
	//
	// This parameter is required.
	//
	// example:
	//
	// DIRECT
	ExchangeType *string `json:"ExchangeType,omitempty" xml:"ExchangeType,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-v0h1kb9nu***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the exchange is an internal exchange. Valid values:
	//
	// - **false**: No
	//
	// - **true**: Yes
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Internal *bool `json:"Internal,omitempty" xml:"Internal,omitempty"`
	// The name of the vhost to which the exchange belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
	// An x-delayed-message exchange lets you use the x-delay header property to specify a delivery delay for a message in milliseconds. The routing rule for the delayed message is determined by the exchange type that you specify for the XDelayedType parameter. This parameter sets the actual exchange type to which the message is delivered after the delay. Valid values:
	//
	// - **DIRECT**: Delivers the delayed message to the specified queue that is bound to a DIRECT exchange.
	//
	// - **TOPIC**: Delivers the delayed message to queues that are bound to a TOPIC exchange.
	//
	// - **FANOUT**: Delivers the delayed message to queues that are bound to a FANOUT exchange.
	//
	// - **HEADERS**: Delivers the delayed message to queues that are bound to a HEADERS exchange.
	//
	// - **X-JMS-TOPIC**: Delivers the delayed message to queues that are bound to an X-JMS-TOPIC exchange.
	//
	// example:
	//
	// DIRECT
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
