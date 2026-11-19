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
	SetEnableSSE(v bool) *CreateTopicRequest
	GetEnableSSE() *bool
	SetKmsKeyId(v string) *CreateTopicRequest
	GetKmsKeyId() *string
	SetMaxMessageSize(v int64) *CreateTopicRequest
	GetMaxMessageSize() *int64
	SetSseAlgorithm(v string) *CreateTopicRequest
	GetSseAlgorithm() *string
	SetSseType(v string) *CreateTopicRequest
	GetSseType() *string
	SetTag(v []*CreateTopicRequestTag) *CreateTopicRequest
	GetTag() []*CreateTopicRequestTag
	SetTopicName(v string) *CreateTopicRequest
	GetTopicName() *string
	SetTopicType(v string) *CreateTopicRequest
	GetTopicType() *string
}

type CreateTopicRequest struct {
	// Specifies whether to enable the logging feature. Valid values:
	//
	// - true: enables the feature.
	//
	// - false: disables the feature.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	EnableLogging *bool `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	// Specifies whether to enable server-side encryption for the message body. Valid values: true or false.
	EnableSSE *bool `json:"EnableSSE,omitempty" xml:"EnableSSE,omitempty"`
	// The KMS key ID used when KMS encryption is enabled. The key must belong to the same account and region as the topic.
	//
	// example:
	//
	// key-xxx
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The maximum size of the message body that can be sent to the topic.
	//
	// Valid values: 1024 to 65536. Unit: bytes.
	//
	// Default value: 65536.
	//
	// example:
	//
	// 10240
	MaxMessageSize *int64 `json:"MaxMessageSize,omitempty" xml:"MaxMessageSize,omitempty"`
	// The server-side encryption algorithm. Currently, only AES-256-GCM is supported. If this parameter is not specified, the default algorithm is used.
	//
	// example:
	//
	// AES-256-GCM
	SseAlgorithm *string `json:"SseAlgorithm,omitempty" xml:"SseAlgorithm,omitempty"`
	// The server-side encryption method. Valid values: SMQ (service-hosted key) or KMS (user-managed KMS key).
	//
	// example:
	//
	// KMS
	SseType *string `json:"SseType,omitempty" xml:"SseType,omitempty"`
	// The list of resource tags.
	Tag []*CreateTopicRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the topic. The name can be up to 120 characters in length and must start with a letter. It can contain letters, digits, and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// The type of the topic. Valid values:
	//
	//    	- normal: standard topic.
	//
	//    	- fifo: FIFO topic.
	//
	// example:
	//
	// normal
	TopicType *string `json:"TopicType,omitempty" xml:"TopicType,omitempty"`
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

func (s *CreateTopicRequest) GetEnableSSE() *bool {
	return s.EnableSSE
}

func (s *CreateTopicRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *CreateTopicRequest) GetMaxMessageSize() *int64 {
	return s.MaxMessageSize
}

func (s *CreateTopicRequest) GetSseAlgorithm() *string {
	return s.SseAlgorithm
}

func (s *CreateTopicRequest) GetSseType() *string {
	return s.SseType
}

func (s *CreateTopicRequest) GetTag() []*CreateTopicRequestTag {
	return s.Tag
}

func (s *CreateTopicRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *CreateTopicRequest) GetTopicType() *string {
	return s.TopicType
}

func (s *CreateTopicRequest) SetEnableLogging(v bool) *CreateTopicRequest {
	s.EnableLogging = &v
	return s
}

func (s *CreateTopicRequest) SetEnableSSE(v bool) *CreateTopicRequest {
	s.EnableSSE = &v
	return s
}

func (s *CreateTopicRequest) SetKmsKeyId(v string) *CreateTopicRequest {
	s.KmsKeyId = &v
	return s
}

func (s *CreateTopicRequest) SetMaxMessageSize(v int64) *CreateTopicRequest {
	s.MaxMessageSize = &v
	return s
}

func (s *CreateTopicRequest) SetSseAlgorithm(v string) *CreateTopicRequest {
	s.SseAlgorithm = &v
	return s
}

func (s *CreateTopicRequest) SetSseType(v string) *CreateTopicRequest {
	s.SseType = &v
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

func (s *CreateTopicRequest) SetTopicType(v string) *CreateTopicRequest {
	s.TopicType = &v
	return s
}

func (s *CreateTopicRequest) Validate() error {
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

type CreateTopicRequestTag struct {
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
