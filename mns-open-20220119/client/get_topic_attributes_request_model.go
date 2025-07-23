// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTag(v []*GetTopicAttributesRequestTag) *GetTopicAttributesRequest
	GetTag() []*GetTopicAttributesRequestTag
	SetTopicName(v string) *GetTopicAttributesRequest
	GetTopicName() *string
}

type GetTopicAttributesRequest struct {
	// The tag.
	Tag []*GetTopicAttributesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo-topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetTopicAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTopicAttributesRequest) GoString() string {
	return s.String()
}

func (s *GetTopicAttributesRequest) GetTag() []*GetTopicAttributesRequestTag {
	return s.Tag
}

func (s *GetTopicAttributesRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *GetTopicAttributesRequest) SetTag(v []*GetTopicAttributesRequestTag) *GetTopicAttributesRequest {
	s.Tag = v
	return s
}

func (s *GetTopicAttributesRequest) SetTopicName(v string) *GetTopicAttributesRequest {
	s.TopicName = &v
	return s
}

func (s *GetTopicAttributesRequest) Validate() error {
	return dara.Validate(s)
}

type GetTopicAttributesRequestTag struct {
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

func (s GetTopicAttributesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s GetTopicAttributesRequestTag) GoString() string {
	return s.String()
}

func (s *GetTopicAttributesRequestTag) GetKey() *string {
	return s.Key
}

func (s *GetTopicAttributesRequestTag) GetValue() *string {
	return s.Value
}

func (s *GetTopicAttributesRequestTag) SetKey(v string) *GetTopicAttributesRequestTag {
	s.Key = &v
	return s
}

func (s *GetTopicAttributesRequestTag) SetValue(v string) *GetTopicAttributesRequestTag {
	s.Value = &v
	return s
}

func (s *GetTopicAttributesRequestTag) Validate() error {
	return dara.Validate(s)
}
