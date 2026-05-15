// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConsumerAPIKeyInput interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *UpdateConsumerAPIKeyInput
	GetActive() *bool
	SetDescription(v string) *UpdateConsumerAPIKeyInput
	GetDescription() *string
}

type UpdateConsumerAPIKeyInput struct {
	// 是否启用该密钥
	//
	// example:
	//
	// false
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// 更新后的描述信息
	//
	// example:
	//
	// 更新后的密钥描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateConsumerAPIKeyInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerAPIKeyInput) GoString() string {
	return s.String()
}

func (s *UpdateConsumerAPIKeyInput) GetActive() *bool {
	return s.Active
}

func (s *UpdateConsumerAPIKeyInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateConsumerAPIKeyInput) SetActive(v bool) *UpdateConsumerAPIKeyInput {
	s.Active = &v
	return s
}

func (s *UpdateConsumerAPIKeyInput) SetDescription(v string) *UpdateConsumerAPIKeyInput {
	s.Description = &v
	return s
}

func (s *UpdateConsumerAPIKeyInput) Validate() error {
	return dara.Validate(s)
}
