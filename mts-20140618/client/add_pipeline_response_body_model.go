// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPipeline(v *AddPipelineResponseBodyPipeline) *AddPipelineResponseBody
	GetPipeline() *AddPipelineResponseBodyPipeline
	SetRequestId(v string) *AddPipelineResponseBody
	GetRequestId() *string
}

type AddPipelineResponseBody struct {
	// The MPS queue.
	Pipeline *AddPipelineResponseBodyPipeline `json:"Pipeline,omitempty" xml:"Pipeline,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CFEA608A-5A1C-4C83-A54B-6197BC250D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *AddPipelineResponseBody) GetPipeline() *AddPipelineResponseBodyPipeline {
	return s.Pipeline
}

func (s *AddPipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPipelineResponseBody) SetPipeline(v *AddPipelineResponseBodyPipeline) *AddPipelineResponseBody {
	s.Pipeline = v
	return s
}

func (s *AddPipelineResponseBody) SetRequestId(v string) *AddPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPipelineResponseBody) Validate() error {
	if s.Pipeline != nil {
		if err := s.Pipeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddPipelineResponseBodyPipeline struct {
	// The ID of the MPS queue.
	//
	// example:
	//
	// ed450ea0bfbd41e29f80a401fb4d****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the MPS queue.
	//
	// example:
	//
	// Media Fingerprint
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The MNS configuration.
	NotifyConfig *AddPipelineResponseBodyPipelineNotifyConfig `json:"NotifyConfig,omitempty" xml:"NotifyConfig,omitempty" type:"Struct"`
	// The quota that is allocated to the MPS queue.
	//
	// example:
	//
	// 10
	QuotaAllocate *int64 `json:"QuotaAllocate,omitempty" xml:"QuotaAllocate,omitempty"`
	// The role.
	//
	// example:
	//
	// AliyunMTSDefaultRole
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The type of the MPS queue.
	//
	// example:
	//
	// Standard
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// The level of the MPS queue.
	//
	// example:
	//
	// 1
	SpeedLevel *int64 `json:"SpeedLevel,omitempty" xml:"SpeedLevel,omitempty"`
	// The state of the MPS queue.
	//
	// 	- Active: The MPS queue is active. The jobs in the MPS queue are scheduled and transcoded by MPS.
	//
	// 	- Paused: The MPS queue is paused. Jobs in the MPS queue are no longer scheduled for transcoding by MPS. All of the jobs in the MPS queue remain in the Submitted state. Jobs that are being transcoded are not affected.
	//
	// example:
	//
	// Active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s AddPipelineResponseBodyPipeline) String() string {
	return dara.Prettify(s)
}

func (s AddPipelineResponseBodyPipeline) GoString() string {
	return s.String()
}

func (s *AddPipelineResponseBodyPipeline) GetId() *string {
	return s.Id
}

func (s *AddPipelineResponseBodyPipeline) GetName() *string {
	return s.Name
}

func (s *AddPipelineResponseBodyPipeline) GetNotifyConfig() *AddPipelineResponseBodyPipelineNotifyConfig {
	return s.NotifyConfig
}

func (s *AddPipelineResponseBodyPipeline) GetQuotaAllocate() *int64 {
	return s.QuotaAllocate
}

func (s *AddPipelineResponseBodyPipeline) GetRole() *string {
	return s.Role
}

func (s *AddPipelineResponseBodyPipeline) GetSpeed() *string {
	return s.Speed
}

func (s *AddPipelineResponseBodyPipeline) GetSpeedLevel() *int64 {
	return s.SpeedLevel
}

func (s *AddPipelineResponseBodyPipeline) GetState() *string {
	return s.State
}

func (s *AddPipelineResponseBodyPipeline) SetId(v string) *AddPipelineResponseBodyPipeline {
	s.Id = &v
	return s
}

func (s *AddPipelineResponseBodyPipeline) SetName(v string) *AddPipelineResponseBodyPipeline {
	s.Name = &v
	return s
}

func (s *AddPipelineResponseBodyPipeline) SetNotifyConfig(v *AddPipelineResponseBodyPipelineNotifyConfig) *AddPipelineResponseBodyPipeline {
	s.NotifyConfig = v
	return s
}

func (s *AddPipelineResponseBodyPipeline) SetQuotaAllocate(v int64) *AddPipelineResponseBodyPipeline {
	s.QuotaAllocate = &v
	return s
}

func (s *AddPipelineResponseBodyPipeline) SetRole(v string) *AddPipelineResponseBodyPipeline {
	s.Role = &v
	return s
}

func (s *AddPipelineResponseBodyPipeline) SetSpeed(v string) *AddPipelineResponseBodyPipeline {
	s.Speed = &v
	return s
}

func (s *AddPipelineResponseBodyPipeline) SetSpeedLevel(v int64) *AddPipelineResponseBodyPipeline {
	s.SpeedLevel = &v
	return s
}

func (s *AddPipelineResponseBodyPipeline) SetState(v string) *AddPipelineResponseBodyPipeline {
	s.State = &v
	return s
}

func (s *AddPipelineResponseBodyPipeline) Validate() error {
	if s.NotifyConfig != nil {
		if err := s.NotifyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddPipelineResponseBodyPipelineNotifyConfig struct {
	// The tag string.
	//
	// example:
	//
	// mts-test
	MqTag *string `json:"MqTag,omitempty" xml:"MqTag,omitempty"`
	// The queue of messages that are received.
	//
	// example:
	//
	// example1
	MqTopic *string `json:"MqTopic,omitempty" xml:"MqTopic,omitempty"`
	// The name of the queue.
	//
	// example:
	//
	// mts-queue-1
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// mts-topic-1
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s AddPipelineResponseBodyPipelineNotifyConfig) String() string {
	return dara.Prettify(s)
}

func (s AddPipelineResponseBodyPipelineNotifyConfig) GoString() string {
	return s.String()
}

func (s *AddPipelineResponseBodyPipelineNotifyConfig) GetMqTag() *string {
	return s.MqTag
}

func (s *AddPipelineResponseBodyPipelineNotifyConfig) GetMqTopic() *string {
	return s.MqTopic
}

func (s *AddPipelineResponseBodyPipelineNotifyConfig) GetQueueName() *string {
	return s.QueueName
}

func (s *AddPipelineResponseBodyPipelineNotifyConfig) GetTopic() *string {
	return s.Topic
}

func (s *AddPipelineResponseBodyPipelineNotifyConfig) SetMqTag(v string) *AddPipelineResponseBodyPipelineNotifyConfig {
	s.MqTag = &v
	return s
}

func (s *AddPipelineResponseBodyPipelineNotifyConfig) SetMqTopic(v string) *AddPipelineResponseBodyPipelineNotifyConfig {
	s.MqTopic = &v
	return s
}

func (s *AddPipelineResponseBodyPipelineNotifyConfig) SetQueueName(v string) *AddPipelineResponseBodyPipelineNotifyConfig {
	s.QueueName = &v
	return s
}

func (s *AddPipelineResponseBodyPipelineNotifyConfig) SetTopic(v string) *AddPipelineResponseBodyPipelineNotifyConfig {
	s.Topic = &v
	return s
}

func (s *AddPipelineResponseBodyPipelineNotifyConfig) Validate() error {
	return dara.Validate(s)
}
