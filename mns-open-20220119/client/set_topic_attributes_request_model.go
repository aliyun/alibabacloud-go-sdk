// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTopicAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableLogging(v bool) *SetTopicAttributesRequest
	GetEnableLogging() *bool
	SetMaxMessageSize(v int64) *SetTopicAttributesRequest
	GetMaxMessageSize() *int64
	SetTopicName(v string) *SetTopicAttributesRequest
	GetTopicName() *string
}

type SetTopicAttributesRequest struct {
	// Specifies whether to enable the log management feature. Valid values:
	//
	// 	- true: enabled.
	//
	// 	- false: disabled. Default value: false.
	//
	// example:
	//
	// True
	EnableLogging *bool `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	// The maximum length of the message that is sent to the topic. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	//
	// example:
	//
	// 65536
	MaxMessageSize *int64 `json:"MaxMessageSize,omitempty" xml:"MaxMessageSize,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s SetTopicAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s SetTopicAttributesRequest) GoString() string {
	return s.String()
}

func (s *SetTopicAttributesRequest) GetEnableLogging() *bool {
	return s.EnableLogging
}

func (s *SetTopicAttributesRequest) GetMaxMessageSize() *int64 {
	return s.MaxMessageSize
}

func (s *SetTopicAttributesRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *SetTopicAttributesRequest) SetEnableLogging(v bool) *SetTopicAttributesRequest {
	s.EnableLogging = &v
	return s
}

func (s *SetTopicAttributesRequest) SetMaxMessageSize(v int64) *SetTopicAttributesRequest {
	s.MaxMessageSize = &v
	return s
}

func (s *SetTopicAttributesRequest) SetTopicName(v string) *SetTopicAttributesRequest {
	s.TopicName = &v
	return s
}

func (s *SetTopicAttributesRequest) Validate() error {
	return dara.Validate(s)
}
