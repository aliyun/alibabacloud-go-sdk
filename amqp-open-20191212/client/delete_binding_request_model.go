// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBindingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindingKey(v string) *DeleteBindingRequest
	GetBindingKey() *string
	SetBindingType(v string) *DeleteBindingRequest
	GetBindingType() *string
	SetDestinationName(v string) *DeleteBindingRequest
	GetDestinationName() *string
	SetInstanceId(v string) *DeleteBindingRequest
	GetInstanceId() *string
	SetSourceExchange(v string) *DeleteBindingRequest
	GetSourceExchange() *string
	SetVirtualHost(v string) *DeleteBindingRequest
	GetVirtualHost() *string
}

type DeleteBindingRequest struct {
	// The binding key.
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
	return dara.Prettify(s)
}

func (s DeleteBindingRequest) GoString() string {
	return s.String()
}

func (s *DeleteBindingRequest) GetBindingKey() *string {
	return s.BindingKey
}

func (s *DeleteBindingRequest) GetBindingType() *string {
	return s.BindingType
}

func (s *DeleteBindingRequest) GetDestinationName() *string {
	return s.DestinationName
}

func (s *DeleteBindingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteBindingRequest) GetSourceExchange() *string {
	return s.SourceExchange
}

func (s *DeleteBindingRequest) GetVirtualHost() *string {
	return s.VirtualHost
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

func (s *DeleteBindingRequest) Validate() error {
	return dara.Validate(s)
}
