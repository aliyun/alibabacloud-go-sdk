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
	SetTag(v []*ListQueueRequestTag) *ListQueueRequest
	GetTag() []*ListQueueRequestTag
}

type ListQueueRequest struct {
	// The page number. Valid values: 1 to 100000000. If you set this parameter to a value smaller than 1, the value of this parameter is 1 by default. If you set this parameter to a value greater than 100000000, the value of this parameter is 100000000 by default.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Value values: 10 to 50. If you set this parameter to a value smaller than 10, the value of this parameter is 10 by default. If you set this parameter to a value greater than 50, the value of this parameter is 50 by default.
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
	// The tags.
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

func (s *ListQueueRequest) SetTag(v []*ListQueueRequestTag) *ListQueueRequest {
	s.Tag = v
	return s
}

func (s *ListQueueRequest) Validate() error {
	return dara.Validate(s)
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
