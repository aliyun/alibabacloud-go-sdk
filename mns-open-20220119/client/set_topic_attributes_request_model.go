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
	EnableLogging *bool `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	// Specifies whether to enable server-side encryption for message bodies. Valid values: true or false.
	EnableSSE *bool `json:"EnableSSE,omitempty" xml:"EnableSSE,omitempty"`
	// The KMS key ID. This parameter is required when you use KMS for encryption. The key must belong to the same Alibaba Cloud account and region as the topic.
	//
	// example:
	//
	// key-xxx
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The maximum size of a message body sent to the topic. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	//
	// example:
	//
	// 65536
	MaxMessageSize *int64 `json:"MaxMessageSize,omitempty" xml:"MaxMessageSize,omitempty"`
	// The server-side encryption algorithm. Currently, only AES-256-GCM is supported. If you do not specify this parameter, the default algorithm is used.
	//
	// example:
	//
	// AES-256-GCM
	SseAlgorithm *string `json:"SseAlgorithm,omitempty" xml:"SseAlgorithm,omitempty"`
	// The server-side encryption method. Valid values: SMQ (service-managed key) or KMS (customer-managed KMS key).
	//
	// example:
	//
	// KMS
	SseType *string `json:"SseType,omitempty" xml:"SseType,omitempty"`
	// The name of the topic. The name can be up to 120 characters in length. It must start with a letter and can contain letters, digits, and hyphens.
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
