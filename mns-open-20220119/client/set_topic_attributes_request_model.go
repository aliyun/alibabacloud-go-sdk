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
	SetEnableSSE(v bool) *SetTopicAttributesRequest
	GetEnableSSE() *bool
	SetKmsKeyId(v string) *SetTopicAttributesRequest
	GetKmsKeyId() *string
	SetMaxMessageSize(v int64) *SetTopicAttributesRequest
	GetMaxMessageSize() *int64
	SetSseAlgorithm(v string) *SetTopicAttributesRequest
	GetSseAlgorithm() *string
	SetSseType(v string) *SetTopicAttributesRequest
	GetSseType() *string
	SetTopicName(v string) *SetTopicAttributesRequest
	GetTopicName() *string
}

type SetTopicAttributesRequest struct {
	// Specifies whether to enable the log management feature. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	EnableLogging *bool   `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	EnableSSE     *bool   `json:"EnableSSE,omitempty" xml:"EnableSSE,omitempty"`
	KmsKeyId      *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The maximum length of the message body sent to the topic. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	//
	// example:
	//
	// 65536
	MaxMessageSize *int64  `json:"MaxMessageSize,omitempty" xml:"MaxMessageSize,omitempty"`
	SseAlgorithm   *string `json:"SseAlgorithm,omitempty" xml:"SseAlgorithm,omitempty"`
	SseType        *string `json:"SseType,omitempty" xml:"SseType,omitempty"`
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

func (s *SetTopicAttributesRequest) GetEnableSSE() *bool {
	return s.EnableSSE
}

func (s *SetTopicAttributesRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *SetTopicAttributesRequest) GetMaxMessageSize() *int64 {
	return s.MaxMessageSize
}

func (s *SetTopicAttributesRequest) GetSseAlgorithm() *string {
	return s.SseAlgorithm
}

func (s *SetTopicAttributesRequest) GetSseType() *string {
	return s.SseType
}

func (s *SetTopicAttributesRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *SetTopicAttributesRequest) SetEnableLogging(v bool) *SetTopicAttributesRequest {
	s.EnableLogging = &v
	return s
}

func (s *SetTopicAttributesRequest) SetEnableSSE(v bool) *SetTopicAttributesRequest {
	s.EnableSSE = &v
	return s
}

func (s *SetTopicAttributesRequest) SetKmsKeyId(v string) *SetTopicAttributesRequest {
	s.KmsKeyId = &v
	return s
}

func (s *SetTopicAttributesRequest) SetMaxMessageSize(v int64) *SetTopicAttributesRequest {
	s.MaxMessageSize = &v
	return s
}

func (s *SetTopicAttributesRequest) SetSseAlgorithm(v string) *SetTopicAttributesRequest {
	s.SseAlgorithm = &v
	return s
}

func (s *SetTopicAttributesRequest) SetSseType(v string) *SetTopicAttributesRequest {
	s.SseType = &v
	return s
}

func (s *SetTopicAttributesRequest) SetTopicName(v string) *SetTopicAttributesRequest {
	s.TopicName = &v
	return s
}

func (s *SetTopicAttributesRequest) Validate() error {
	return dara.Validate(s)
}
