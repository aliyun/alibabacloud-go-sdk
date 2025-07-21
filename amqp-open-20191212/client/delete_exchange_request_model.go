// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExchangeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExchangeName(v string) *DeleteExchangeRequest
	GetExchangeName() *string
	SetInstanceId(v string) *DeleteExchangeRequest
	GetInstanceId() *string
	SetVirtualHost(v string) *DeleteExchangeRequest
	GetVirtualHost() *string
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
	return dara.Prettify(s)
}

func (s DeleteExchangeRequest) GoString() string {
	return s.String()
}

func (s *DeleteExchangeRequest) GetExchangeName() *string {
	return s.ExchangeName
}

func (s *DeleteExchangeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteExchangeRequest) GetVirtualHost() *string {
	return s.VirtualHost
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

func (s *DeleteExchangeRequest) Validate() error {
	return dara.Validate(s)
}
