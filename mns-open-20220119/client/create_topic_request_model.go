// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableLogging(v bool) *CreateTopicRequest
	GetEnableLogging() *bool
	SetMaxMessageSize(v int64) *CreateTopicRequest
	GetMaxMessageSize() *int64
	SetTag(v []*CreateTopicRequestTag) *CreateTopicRequest
	GetTag() []*CreateTopicRequestTag
	SetTopicName(v string) *CreateTopicRequest
	GetTopicName() *string
}

type CreateTopicRequest struct {
	// Specifies whether to enable the log management feature. Valid values:
	//
	// 	- true: enabled.
	//
	// 	- false: disabled. Default value: false.
	//
	// example:
	//
	// true
	EnableLogging *bool `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	// The maximum length of the message that is sent to the topic. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	//
	// example:
	//
	// 10240
	MaxMessageSize *int64 `json:"MaxMessageSize,omitempty" xml:"MaxMessageSize,omitempty"`
	// The tags.
	Tag []*CreateTopicRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the topic that you want to create.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s CreateTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTopicRequest) GoString() string {
	return s.String()
}

func (s *CreateTopicRequest) GetEnableLogging() *bool {
	return s.EnableLogging
}

func (s *CreateTopicRequest) GetMaxMessageSize() *int64 {
	return s.MaxMessageSize
}

func (s *CreateTopicRequest) GetTag() []*CreateTopicRequestTag {
	return s.Tag
}

func (s *CreateTopicRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *CreateTopicRequest) SetEnableLogging(v bool) *CreateTopicRequest {
	s.EnableLogging = &v
	return s
}

func (s *CreateTopicRequest) SetMaxMessageSize(v int64) *CreateTopicRequest {
	s.MaxMessageSize = &v
	return s
}

func (s *CreateTopicRequest) SetTag(v []*CreateTopicRequestTag) *CreateTopicRequest {
	s.Tag = v
	return s
}

func (s *CreateTopicRequest) SetTopicName(v string) *CreateTopicRequest {
	s.TopicName = &v
	return s
}

func (s *CreateTopicRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTopicRequestTag struct {
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
	// joyce.wang
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTopicRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTopicRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTopicRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTopicRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTopicRequestTag) SetKey(v string) *CreateTopicRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTopicRequestTag) SetValue(v string) *CreateTopicRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTopicRequestTag) Validate() error {
	return dara.Validate(s)
}
