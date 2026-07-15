// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRemoveConsumerGroupConsumersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerIds(v []*string) *BatchRemoveConsumerGroupConsumersRequest
	GetConsumerIds() []*string
}

type BatchRemoveConsumerGroupConsumersRequest struct {
	// example:
	//
	// ["cs-8c13d2b4f8a1"]
	ConsumerIds []*string `json:"consumerIds,omitempty" xml:"consumerIds,omitempty" type:"Repeated"`
}

func (s BatchRemoveConsumerGroupConsumersRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchRemoveConsumerGroupConsumersRequest) GoString() string {
	return s.String()
}

func (s *BatchRemoveConsumerGroupConsumersRequest) GetConsumerIds() []*string {
	return s.ConsumerIds
}

func (s *BatchRemoveConsumerGroupConsumersRequest) SetConsumerIds(v []*string) *BatchRemoveConsumerGroupConsumersRequest {
	s.ConsumerIds = v
	return s
}

func (s *BatchRemoveConsumerGroupConsumersRequest) Validate() error {
	return dara.Validate(s)
}
