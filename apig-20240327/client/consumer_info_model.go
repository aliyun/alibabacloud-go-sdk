// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsumerInfo interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerId(v string) *ConsumerInfo
	GetConsumerId() *string
	SetEnable(v bool) *ConsumerInfo
	GetEnable() *bool
	SetName(v string) *ConsumerInfo
	GetName() *string
}

type ConsumerInfo struct {
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	Enable     *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ConsumerInfo) String() string {
	return dara.Prettify(s)
}

func (s ConsumerInfo) GoString() string {
	return s.String()
}

func (s *ConsumerInfo) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ConsumerInfo) GetEnable() *bool {
	return s.Enable
}

func (s *ConsumerInfo) GetName() *string {
	return s.Name
}

func (s *ConsumerInfo) SetConsumerId(v string) *ConsumerInfo {
	s.ConsumerId = &v
	return s
}

func (s *ConsumerInfo) SetEnable(v bool) *ConsumerInfo {
	s.Enable = &v
	return s
}

func (s *ConsumerInfo) SetName(v string) *ConsumerInfo {
	s.Name = &v
	return s
}

func (s *ConsumerInfo) Validate() error {
	return dara.Validate(s)
}
