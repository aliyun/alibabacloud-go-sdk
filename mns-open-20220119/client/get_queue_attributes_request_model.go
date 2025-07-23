// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQueueName(v string) *GetQueueAttributesRequest
	GetQueueName() *string
	SetTag(v []*GetQueueAttributesRequestTag) *GetQueueAttributesRequest
	GetTag() []*GetQueueAttributesRequestTag
}

type GetQueueAttributesRequest struct {
	// The name of the queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo-queue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The tags.
	Tag []*GetQueueAttributesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetQueueAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueueAttributesRequest) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *GetQueueAttributesRequest) GetTag() []*GetQueueAttributesRequestTag {
	return s.Tag
}

func (s *GetQueueAttributesRequest) SetQueueName(v string) *GetQueueAttributesRequest {
	s.QueueName = &v
	return s
}

func (s *GetQueueAttributesRequest) SetTag(v []*GetQueueAttributesRequestTag) *GetQueueAttributesRequest {
	s.Tag = v
	return s
}

func (s *GetQueueAttributesRequest) Validate() error {
	return dara.Validate(s)
}

type GetQueueAttributesRequestTag struct {
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

func (s GetQueueAttributesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s GetQueueAttributesRequestTag) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesRequestTag) GetKey() *string {
	return s.Key
}

func (s *GetQueueAttributesRequestTag) GetValue() *string {
	return s.Value
}

func (s *GetQueueAttributesRequestTag) SetKey(v string) *GetQueueAttributesRequestTag {
	s.Key = &v
	return s
}

func (s *GetQueueAttributesRequestTag) SetValue(v string) *GetQueueAttributesRequestTag {
	s.Value = &v
	return s
}

func (s *GetQueueAttributesRequestTag) Validate() error {
	return dara.Validate(s)
}
