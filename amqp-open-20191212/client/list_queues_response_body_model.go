// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListQueuesResponseBodyData) *ListQueuesResponseBody
	GetData() *ListQueuesResponseBodyData
	SetRequestId(v string) *ListQueuesResponseBody
	GetRequestId() *string
}

type ListQueuesResponseBody struct {
	// The returned data.
	Data *ListQueuesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CE811989-9F02-42CE-97A6-2239CB5C2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListQueuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQueuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBody) GetData() *ListQueuesResponseBodyData {
	return s.Data
}

func (s *ListQueuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQueuesResponseBody) SetData(v *ListQueuesResponseBodyData) *ListQueuesResponseBody {
	s.Data = v
	return s
}

func (s *ListQueuesResponseBody) SetRequestId(v string) *ListQueuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueuesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQueuesResponseBodyData struct {
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
	// The queues.
	Queues []*ListQueuesResponseBodyDataQueues `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
}

func (s ListQueuesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListQueuesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQueuesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQueuesResponseBodyData) GetQueues() []*ListQueuesResponseBodyDataQueues {
	return s.Queues
}

func (s *ListQueuesResponseBodyData) SetMaxResults(v int32) *ListQueuesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListQueuesResponseBodyData) SetNextToken(v string) *ListQueuesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListQueuesResponseBodyData) SetQueues(v []*ListQueuesResponseBodyDataQueues) *ListQueuesResponseBodyData {
	s.Queues = v
	return s
}

func (s *ListQueuesResponseBodyData) Validate() error {
	if s.Queues != nil {
		for _, item := range s.Queues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQueuesResponseBodyDataQueues struct {
	// The attributes.
	//
	// example:
	//
	// test
	Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// Indicates whether the queue was automatically deleted.
	//
	// example:
	//
	// false
	AutoDeleteState *bool `json:"AutoDeleteState,omitempty" xml:"AutoDeleteState,omitempty"`
	// The time when the queue was created.
	//
	// example:
	//
	// 1580887085240
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the queue is an exclusive queue.
	//
	// example:
	//
	// false
	ExclusiveState *bool `json:"ExclusiveState,omitempty" xml:"ExclusiveState,omitempty"`
	// The time when messages in the queue were last consumed.
	//
	// example:
	//
	// 1680887085240
	LastConsumeTime *int64 `json:"LastConsumeTime,omitempty" xml:"LastConsumeTime,omitempty"`
	// The queue name.
	//
	// example:
	//
	// QueueTest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance to which the queue belongs.
	//
	// example:
	//
	// 1880770869023***
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The vhost name.
	//
	// example:
	//
	// test
	VHostName *string `json:"VHostName,omitempty" xml:"VHostName,omitempty"`
}

func (s ListQueuesResponseBodyDataQueues) String() string {
	return dara.Prettify(s)
}

func (s ListQueuesResponseBodyDataQueues) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyDataQueues) GetAttributes() map[string]interface{} {
	return s.Attributes
}

func (s *ListQueuesResponseBodyDataQueues) GetAutoDeleteState() *bool {
	return s.AutoDeleteState
}

func (s *ListQueuesResponseBodyDataQueues) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListQueuesResponseBodyDataQueues) GetExclusiveState() *bool {
	return s.ExclusiveState
}

func (s *ListQueuesResponseBodyDataQueues) GetLastConsumeTime() *int64 {
	return s.LastConsumeTime
}

func (s *ListQueuesResponseBodyDataQueues) GetName() *string {
	return s.Name
}

func (s *ListQueuesResponseBodyDataQueues) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListQueuesResponseBodyDataQueues) GetVHostName() *string {
	return s.VHostName
}

func (s *ListQueuesResponseBodyDataQueues) SetAttributes(v map[string]interface{}) *ListQueuesResponseBodyDataQueues {
	s.Attributes = v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetAutoDeleteState(v bool) *ListQueuesResponseBodyDataQueues {
	s.AutoDeleteState = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetCreateTime(v int64) *ListQueuesResponseBodyDataQueues {
	s.CreateTime = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetExclusiveState(v bool) *ListQueuesResponseBodyDataQueues {
	s.ExclusiveState = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetLastConsumeTime(v int64) *ListQueuesResponseBodyDataQueues {
	s.LastConsumeTime = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetName(v string) *ListQueuesResponseBodyDataQueues {
	s.Name = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetOwnerId(v string) *ListQueuesResponseBodyDataQueues {
	s.OwnerId = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetVHostName(v string) *ListQueuesResponseBodyDataQueues {
	s.VHostName = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) Validate() error {
	return dara.Validate(s)
}
