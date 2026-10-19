// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int64) *ListQueueRequest
	GetPageNum() *int64
	SetPageSize(v int64) *ListQueueRequest
	GetPageSize() *int64
	SetQueueName(v string) *ListQueueRequest
	GetQueueName() *string
	SetQueueType(v string) *ListQueueRequest
	GetQueueType() *string
	SetTag(v []*ListQueueRequestTag) *ListQueueRequest
	GetTag() []*ListQueueRequestTag
}

type ListQueueRequest struct {
	// The page number of the returned results. Valid values: 1 to 100000000. If the specified value is less than 1, the system defaults to 1. If the specified value is greater than 100000000, the system defaults to 100000000.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Valid values: 10 to 50. If the specified value is less than 10, the system defaults to 10. If the specified value is greater than 50, the system defaults to 50.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the queue.
	//
	// example:
	//
	// demo-queue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The type of the queue. Valid values:
	//
	//    	- normal: standard queue
	//
	//    	- fifo: FIFO queue
	//
	// example:
	//
	// normal
	QueueType *string `json:"QueueType,omitempty" xml:"QueueType,omitempty"`
	// The list of resource tags.
	Tag []*ListQueueRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQueueRequest) GoString() string {
	return s.String()
}

func (s *ListQueueRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListQueueRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListQueueRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *ListQueueRequest) GetQueueType() *string {
	return s.QueueType
}

func (s *ListQueueRequest) GetTag() []*ListQueueRequestTag {
	return s.Tag
}

func (s *ListQueueRequest) SetPageNum(v int64) *ListQueueRequest {
	s.PageNum = &v
	return s
}

func (s *ListQueueRequest) SetPageSize(v int64) *ListQueueRequest {
	s.PageSize = &v
	return s
}

func (s *ListQueueRequest) SetQueueName(v string) *ListQueueRequest {
	s.QueueName = &v
	return s
}

func (s *ListQueueRequest) SetQueueType(v string) *ListQueueRequest {
	s.QueueType = &v
	return s
}

func (s *ListQueueRequest) SetTag(v []*ListQueueRequestTag) *ListQueueRequest {
	s.Tag = v
	return s
}

func (s *ListQueueRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQueueRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListQueueRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListQueueRequestTag) GoString() string {
	return s.String()
}

func (s *ListQueueRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListQueueRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListQueueRequestTag) SetKey(v string) *ListQueueRequestTag {
	s.Key = &v
	return s
}

func (s *ListQueueRequestTag) SetValue(v string) *ListQueueRequestTag {
	s.Value = &v
	return s
}

func (s *ListQueueRequestTag) Validate() error {
	return dara.Validate(s)
}
