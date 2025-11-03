// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueueConsumersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListQueueConsumersResponseBodyData) *ListQueueConsumersResponseBody
	GetData() *ListQueueConsumersResponseBodyData
	SetRequestId(v string) *ListQueueConsumersResponseBody
	GetRequestId() *string
}

type ListQueueConsumersResponseBody struct {
	// The returned data.
	Data *ListQueueConsumersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4409B7D5-E4EC-4EB5-804A-385DCDFCD***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListQueueConsumersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQueueConsumersResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueueConsumersResponseBody) GetData() *ListQueueConsumersResponseBodyData {
	return s.Data
}

func (s *ListQueueConsumersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQueueConsumersResponseBody) SetData(v *ListQueueConsumersResponseBodyData) *ListQueueConsumersResponseBody {
	s.Data = v
	return s
}

func (s *ListQueueConsumersResponseBody) SetRequestId(v string) *ListQueueConsumersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueueConsumersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQueueConsumersResponseBodyData struct {
	// The consumers.
	Consumers []*ListQueueConsumersResponseBodyDataConsumers `json:"Consumers,omitempty" xml:"Consumers,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current returned page. If this parameter is empty, all data is retrieved.
	//
	// example:
	//
	// caebacccb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListQueueConsumersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListQueueConsumersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQueueConsumersResponseBodyData) GetConsumers() []*ListQueueConsumersResponseBodyDataConsumers {
	return s.Consumers
}

func (s *ListQueueConsumersResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQueueConsumersResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQueueConsumersResponseBodyData) SetConsumers(v []*ListQueueConsumersResponseBodyDataConsumers) *ListQueueConsumersResponseBodyData {
	s.Consumers = v
	return s
}

func (s *ListQueueConsumersResponseBodyData) SetMaxResults(v int32) *ListQueueConsumersResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListQueueConsumersResponseBodyData) SetNextToken(v string) *ListQueueConsumersResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListQueueConsumersResponseBodyData) Validate() error {
	if s.Consumers != nil {
		for _, item := range s.Consumers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQueueConsumersResponseBodyDataConsumers struct {
	// The consumer tag.
	//
	// example:
	//
	// sgen-1
	ConsumerTag *string `json:"ConsumerTag,omitempty" xml:"ConsumerTag,omitempty"`
}

func (s ListQueueConsumersResponseBodyDataConsumers) String() string {
	return dara.Prettify(s)
}

func (s ListQueueConsumersResponseBodyDataConsumers) GoString() string {
	return s.String()
}

func (s *ListQueueConsumersResponseBodyDataConsumers) GetConsumerTag() *string {
	return s.ConsumerTag
}

func (s *ListQueueConsumersResponseBodyDataConsumers) SetConsumerTag(v string) *ListQueueConsumersResponseBodyDataConsumers {
	s.ConsumerTag = &v
	return s
}

func (s *ListQueueConsumersResponseBodyDataConsumers) Validate() error {
	return dara.Validate(s)
}
