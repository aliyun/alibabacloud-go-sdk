// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPipeline(v *UpdatePipelineResponseBodyPipeline) *UpdatePipelineResponseBody
	GetPipeline() *UpdatePipelineResponseBodyPipeline
	SetRequestId(v string) *UpdatePipelineResponseBody
	GetRequestId() *string
}

type UpdatePipelineResponseBody struct {
	// The details of the MPS queue.
	Pipeline *UpdatePipelineResponseBodyPipeline `json:"Pipeline,omitempty" xml:"Pipeline,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1FE0F96B-544D-4244-9D83-DFCFB0E5A231
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelineResponseBody) GetPipeline() *UpdatePipelineResponseBodyPipeline {
	return s.Pipeline
}

func (s *UpdatePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePipelineResponseBody) SetPipeline(v *UpdatePipelineResponseBodyPipeline) *UpdatePipelineResponseBody {
	s.Pipeline = v
	return s
}

func (s *UpdatePipelineResponseBody) SetRequestId(v string) *UpdatePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePipelineResponseBody) Validate() error {
	if s.Pipeline != nil {
		if err := s.Pipeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePipelineResponseBodyPipeline struct {
	// The ID of the MPS queue.
	//
	// example:
	//
	// d1ce4d3efcb549419193f50f1fcd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The new name of the MPS queue.
	//
	// example:
	//
	// example-pipeline-****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The MNS configuration.
	NotifyConfig *UpdatePipelineResponseBodyPipelineNotifyConfig `json:"NotifyConfig,omitempty" xml:"NotifyConfig,omitempty" type:"Struct"`
	// The quota that is allocated to the MPS queue.
	//
	// example:
	//
	// 10
	QuotaAllocate *int64 `json:"QuotaAllocate,omitempty" xml:"QuotaAllocate,omitempty"`
	// The role that is assigned to the current RAM user.
	//
	// example:
	//
	// AliyunMTSExampleRole
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The type of the MPS queue. Default value: **Standard**. Valid values:
	//
	// 	- **Boost**: MPS queue with transcoding speed boosted
	//
	// 	- **Standard**: standard MPS queue
	//
	// 	- **NarrowBandHDV2**: MPS queue that supports Narrowband HD 2.0
	//
	// 	- **AIVideoCover**: MPS queue for intelligent snapshot capture
	//
	// 	- **AIVideoFPShot**: MPS queue for media fingerprinting
	//
	// 	- **AIVideoCensor**: MPS queue for automated review
	//
	// 	- **AIVideoMCU**: MPS queue for smart tagging
	//
	// 	- **AIVideoSummary**: MPS queue for video synopsis
	//
	// 	- **AIVideoPorn**: MPS queue for pornography detection in videos
	//
	// 	- **AIAudioKWS**: MPS queue for keyword recognition in audio
	//
	// 	- **AIAudioASR**: MPS queue for speech-to-text conversion
	//
	// example:
	//
	// Standard
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// The state of the MPS queue. Valid values:
	//
	// 	- **Active**: The MPS queue is active.
	//
	// 	- **Paused**: The MPS queue is paused.
	//
	// example:
	//
	// Paused
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s UpdatePipelineResponseBodyPipeline) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineResponseBodyPipeline) GoString() string {
	return s.String()
}

func (s *UpdatePipelineResponseBodyPipeline) GetId() *string {
	return s.Id
}

func (s *UpdatePipelineResponseBodyPipeline) GetName() *string {
	return s.Name
}

func (s *UpdatePipelineResponseBodyPipeline) GetNotifyConfig() *UpdatePipelineResponseBodyPipelineNotifyConfig {
	return s.NotifyConfig
}

func (s *UpdatePipelineResponseBodyPipeline) GetQuotaAllocate() *int64 {
	return s.QuotaAllocate
}

func (s *UpdatePipelineResponseBodyPipeline) GetRole() *string {
	return s.Role
}

func (s *UpdatePipelineResponseBodyPipeline) GetSpeed() *string {
	return s.Speed
}

func (s *UpdatePipelineResponseBodyPipeline) GetState() *string {
	return s.State
}

func (s *UpdatePipelineResponseBodyPipeline) SetId(v string) *UpdatePipelineResponseBodyPipeline {
	s.Id = &v
	return s
}

func (s *UpdatePipelineResponseBodyPipeline) SetName(v string) *UpdatePipelineResponseBodyPipeline {
	s.Name = &v
	return s
}

func (s *UpdatePipelineResponseBodyPipeline) SetNotifyConfig(v *UpdatePipelineResponseBodyPipelineNotifyConfig) *UpdatePipelineResponseBodyPipeline {
	s.NotifyConfig = v
	return s
}

func (s *UpdatePipelineResponseBodyPipeline) SetQuotaAllocate(v int64) *UpdatePipelineResponseBodyPipeline {
	s.QuotaAllocate = &v
	return s
}

func (s *UpdatePipelineResponseBodyPipeline) SetRole(v string) *UpdatePipelineResponseBodyPipeline {
	s.Role = &v
	return s
}

func (s *UpdatePipelineResponseBodyPipeline) SetSpeed(v string) *UpdatePipelineResponseBodyPipeline {
	s.Speed = &v
	return s
}

func (s *UpdatePipelineResponseBodyPipeline) SetState(v string) *UpdatePipelineResponseBodyPipeline {
	s.State = &v
	return s
}

func (s *UpdatePipelineResponseBodyPipeline) Validate() error {
	if s.NotifyConfig != nil {
		if err := s.NotifyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePipelineResponseBodyPipelineNotifyConfig struct {
	// The tags of the messages.
	//
	// example:
	//
	// mts-test
	MqTag *string `json:"MqTag,omitempty" xml:"MqTag,omitempty"`
	// The queue of messages that are received.
	//
	// example:
	//
	// example1,example2
	MqTopic *string `json:"MqTopic,omitempty" xml:"MqTopic,omitempty"`
	// The queue that is created in MNS.
	//
	// example:
	//
	// example-queue-****
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The topic that is created in MNS.
	//
	// example:
	//
	// example-topic-****
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s UpdatePipelineResponseBodyPipelineNotifyConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineResponseBodyPipelineNotifyConfig) GoString() string {
	return s.String()
}

func (s *UpdatePipelineResponseBodyPipelineNotifyConfig) GetMqTag() *string {
	return s.MqTag
}

func (s *UpdatePipelineResponseBodyPipelineNotifyConfig) GetMqTopic() *string {
	return s.MqTopic
}

func (s *UpdatePipelineResponseBodyPipelineNotifyConfig) GetQueueName() *string {
	return s.QueueName
}

func (s *UpdatePipelineResponseBodyPipelineNotifyConfig) GetTopic() *string {
	return s.Topic
}

func (s *UpdatePipelineResponseBodyPipelineNotifyConfig) SetMqTag(v string) *UpdatePipelineResponseBodyPipelineNotifyConfig {
	s.MqTag = &v
	return s
}

func (s *UpdatePipelineResponseBodyPipelineNotifyConfig) SetMqTopic(v string) *UpdatePipelineResponseBodyPipelineNotifyConfig {
	s.MqTopic = &v
	return s
}

func (s *UpdatePipelineResponseBodyPipelineNotifyConfig) SetQueueName(v string) *UpdatePipelineResponseBodyPipelineNotifyConfig {
	s.QueueName = &v
	return s
}

func (s *UpdatePipelineResponseBodyPipelineNotifyConfig) SetTopic(v string) *UpdatePipelineResponseBodyPipelineNotifyConfig {
	s.Topic = &v
	return s
}

func (s *UpdatePipelineResponseBodyPipelineNotifyConfig) Validate() error {
	return dara.Validate(s)
}
