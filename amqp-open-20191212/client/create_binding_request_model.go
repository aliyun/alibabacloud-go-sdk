// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBindingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArgument(v string) *CreateBindingRequest
	GetArgument() *string
	SetBindingKey(v string) *CreateBindingRequest
	GetBindingKey() *string
	SetBindingType(v string) *CreateBindingRequest
	GetBindingType() *string
	SetDestinationName(v string) *CreateBindingRequest
	GetDestinationName() *string
	SetInstanceId(v string) *CreateBindingRequest
	GetInstanceId() *string
	SetSourceExchange(v string) *CreateBindingRequest
	GetSourceExchange() *string
	SetVirtualHost(v string) *CreateBindingRequest
	GetVirtualHost() *string
}

type CreateBindingRequest struct {
	// The key-value pairs for the message header attributes. The message header attributes consist of one or more key-value pairs. The x-match attribute is required. Other attributes are custom. The x-match attribute supports the following values:
	//
	// - all: This is the default value. All key-value pairs in the message header must match.
	//
	// - any: At least one key-value pair in the message header must match.
	//
	// Separate attributes with semicolons (;) and separate keys from values with colons (:). Example: x-match:all;type:report;format:pdf
	//
	// This parameter is valid only for headers exchanges. For other types of exchanges, this parameter is ignored.
	//
	// example:
	//
	// x-match:all;type:report;format:pdf
	Argument *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	// The binding key.
	//
	// - If the source exchange is not a topic exchange:
	//
	//   - It can contain letters, digits, hyphens (-), underscores (_), periods (.), forward slashes (/), and at signs (@).
	//
	//   - The length must be 1 to 255 characters.
	//
	// - If the source exchange is a topic exchange:
	//
	//   - It can contain letters, digits, hyphens (-), underscores (_), asterisks (\\*), periods (.), number signs (#), forward slashes (/), and at signs (@).
	//
	//   - The key cannot start or end with a period (.). If the key starts with a number sign (#) or an asterisk (\\*), a period (.) must immediately follow. If the key ends with a number sign (#) or an asterisk (\\*), a period (.) must immediately precede it. If a number sign (#) or an asterisk (\\*) is in the middle of the key, it must have a period (.) on both sides.
	//
	//   - The length must be 1 to 255 characters.
	//
	// example:
	//
	// .test
	BindingKey *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	// The type of the destination object. Valid values:
	//
	// - 0: Queue
	//
	// - 1: Exchange
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	BindingType *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	// The name of the binding destination. The destination must be created in the console. It must belong to the same vhost as `SourceExchange`. The `VirtualHost` parameter specifies the vhost.
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
	// The name of the source exchange. This exchange must be created in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// NormalEX
	SourceExchange *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
	// The name of the vhost. The vhost must be created in the console. Both `DestinationName` and `SourceExchange` must belong to this vhost.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s CreateBindingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBindingRequest) GoString() string {
	return s.String()
}

func (s *CreateBindingRequest) GetArgument() *string {
	return s.Argument
}

func (s *CreateBindingRequest) GetBindingKey() *string {
	return s.BindingKey
}

func (s *CreateBindingRequest) GetBindingType() *string {
	return s.BindingType
}

func (s *CreateBindingRequest) GetDestinationName() *string {
	return s.DestinationName
}

func (s *CreateBindingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateBindingRequest) GetSourceExchange() *string {
	return s.SourceExchange
}

func (s *CreateBindingRequest) GetVirtualHost() *string {
	return s.VirtualHost
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

func (s *CreateBindingRequest) Validate() error {
	return dara.Validate(s)
}
