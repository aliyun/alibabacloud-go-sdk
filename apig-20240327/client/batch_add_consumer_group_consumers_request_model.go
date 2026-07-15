// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddConsumerGroupConsumersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerIds(v []*string) *BatchAddConsumerGroupConsumersRequest
	GetConsumerIds() []*string
}

type BatchAddConsumerGroupConsumersRequest struct {
	// example:
	//
	// ["cs-8c13d2b4f8a1"]
	ConsumerIds []*string `json:"consumerIds,omitempty" xml:"consumerIds,omitempty" type:"Repeated"`
}

func (s BatchAddConsumerGroupConsumersRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchAddConsumerGroupConsumersRequest) GoString() string {
	return s.String()
}

func (s *BatchAddConsumerGroupConsumersRequest) GetConsumerIds() []*string {
	return s.ConsumerIds
}

func (s *BatchAddConsumerGroupConsumersRequest) SetConsumerIds(v []*string) *BatchAddConsumerGroupConsumersRequest {
	s.ConsumerIds = v
	return s
}

func (s *BatchAddConsumerGroupConsumersRequest) Validate() error {
	return dara.Validate(s)
}
