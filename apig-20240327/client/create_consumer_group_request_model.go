// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerGroupId(v string) *CreateConsumerGroupRequest
	GetConsumerGroupId() *string
	SetDescription(v string) *CreateConsumerGroupRequest
	GetDescription() *string
	SetGatewayType(v string) *CreateConsumerGroupRequest
	GetGatewayType() *string
	SetName(v string) *CreateConsumerGroupRequest
	GetName() *string
}

type CreateConsumerGroupRequest struct {
	// example:
	//
	// csg-8c13d2b4f8a1
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// example:
	//
	// 用于线上 API 调用方分组
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// API
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// example:
	//
	// api-consumer-group
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateConsumerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupRequest) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *CreateConsumerGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateConsumerGroupRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *CreateConsumerGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateConsumerGroupRequest) SetConsumerGroupId(v string) *CreateConsumerGroupRequest {
	s.ConsumerGroupId = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetDescription(v string) *CreateConsumerGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetGatewayType(v string) *CreateConsumerGroupRequest {
	s.GatewayType = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetName(v string) *CreateConsumerGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateConsumerGroupRequest) Validate() error {
	return dara.Validate(s)
}
