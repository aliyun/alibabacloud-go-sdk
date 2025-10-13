// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsumerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerId(v string) *ConsumerConfig
	GetConsumerId() *string
	SetName(v string) *ConsumerConfig
	GetName() *string
}

type ConsumerConfig struct {
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ConsumerConfig) String() string {
	return dara.Prettify(s)
}

func (s ConsumerConfig) GoString() string {
	return s.String()
}

func (s *ConsumerConfig) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ConsumerConfig) GetName() *string {
	return s.Name
}

func (s *ConsumerConfig) SetConsumerId(v string) *ConsumerConfig {
	s.ConsumerId = &v
	return s
}

func (s *ConsumerConfig) SetName(v string) *ConsumerConfig {
	s.Name = &v
	return s
}

func (s *ConsumerConfig) Validate() error {
	return dara.Validate(s)
}
