// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsumerGroupInfo interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerGroupId(v string) *ConsumerGroupInfo
	GetConsumerGroupId() *string
	SetGatewayType(v string) *ConsumerGroupInfo
	GetGatewayType() *string
	SetName(v string) *ConsumerGroupInfo
	GetName() *string
}

type ConsumerGroupInfo struct {
	// example:
	//
	// csg-8c13d2b4f8a1
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// example:
	//
	// API
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// example:
	//
	// api-consumer-group
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ConsumerGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s ConsumerGroupInfo) GoString() string {
	return s.String()
}

func (s *ConsumerGroupInfo) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *ConsumerGroupInfo) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ConsumerGroupInfo) GetName() *string {
	return s.Name
}

func (s *ConsumerGroupInfo) SetConsumerGroupId(v string) *ConsumerGroupInfo {
	s.ConsumerGroupId = &v
	return s
}

func (s *ConsumerGroupInfo) SetGatewayType(v string) *ConsumerGroupInfo {
	s.GatewayType = &v
	return s
}

func (s *ConsumerGroupInfo) SetName(v string) *ConsumerGroupInfo {
	s.Name = &v
	return s
}

func (s *ConsumerGroupInfo) Validate() error {
	return dara.Validate(s)
}
