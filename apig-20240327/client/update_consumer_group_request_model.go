// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConsumerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateConsumerGroupRequest
	GetDescription() *string
	SetName(v string) *UpdateConsumerGroupRequest
	GetName() *string
}

type UpdateConsumerGroupRequest struct {
	// example:
	//
	// 用于线上 API 调用方分组
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// api-consumer-group
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateConsumerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateConsumerGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateConsumerGroupRequest) SetDescription(v string) *UpdateConsumerGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateConsumerGroupRequest) SetName(v string) *UpdateConsumerGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateConsumerGroupRequest) Validate() error {
	return dara.Validate(s)
}
