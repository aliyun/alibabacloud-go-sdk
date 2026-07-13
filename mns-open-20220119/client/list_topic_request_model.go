// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int64) *ListTopicRequest
	GetPageNum() *int64
	SetPageSize(v int64) *ListTopicRequest
	GetPageSize() *int64
	SetTag(v []*ListTopicRequestTag) *ListTopicRequest
	GetTag() []*ListTopicRequestTag
	SetTopicName(v string) *ListTopicRequest
	GetTopicName() *string
	SetTopicType(v string) *ListTopicRequest
	GetTopicType() *string
}

type ListTopicRequest struct {
	// The page number of the results to return.
	//
	// Valid values: 1 to 100000000.
	//
	// If you set this parameter to a value less than 1, the system uses 1 by default. If you set this parameter to a value greater than 100000000, the system uses 100000000 by default.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of results to return on each page.
	//
	// Valid values: 10 to 50.
	//
	// If you set this parameter to a value less than 10, the system uses 10 by default. If you set this parameter to a value greater than 50, the system uses 50 by default.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of tags.
	Tag []*ListTopicRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the topic.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// The type of the topic. Valid values:
	//
	//    	- normal: normal topic
	//
	//    	- fifo: FIFO topic
	//
	// example:
	//
	// normal
	TopicType *string `json:"TopicType,omitempty" xml:"TopicType,omitempty"`
}

func (s ListTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTopicRequest) GoString() string {
	return s.String()
}

func (s *ListTopicRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListTopicRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTopicRequest) GetTag() []*ListTopicRequestTag {
	return s.Tag
}

func (s *ListTopicRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *ListTopicRequest) GetTopicType() *string {
	return s.TopicType
}

func (s *ListTopicRequest) SetPageNum(v int64) *ListTopicRequest {
	s.PageNum = &v
	return s
}

func (s *ListTopicRequest) SetPageSize(v int64) *ListTopicRequest {
	s.PageSize = &v
	return s
}

func (s *ListTopicRequest) SetTag(v []*ListTopicRequestTag) *ListTopicRequest {
	s.Tag = v
	return s
}

func (s *ListTopicRequest) SetTopicName(v string) *ListTopicRequest {
	s.TopicName = &v
	return s
}

func (s *ListTopicRequest) SetTopicType(v string) *ListTopicRequest {
	s.TopicType = &v
	return s
}

func (s *ListTopicRequest) Validate() error {
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

type ListTopicRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTopicRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTopicRequestTag) GoString() string {
	return s.String()
}

func (s *ListTopicRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTopicRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTopicRequestTag) SetKey(v string) *ListTopicRequestTag {
	s.Key = &v
	return s
}

func (s *ListTopicRequestTag) SetValue(v string) *ListTopicRequestTag {
	s.Value = &v
	return s
}

func (s *ListTopicRequestTag) Validate() error {
	return dara.Validate(s)
}
