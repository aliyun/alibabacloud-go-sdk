// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPipelineListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNonExistPids(v *QueryPipelineListResponseBodyNonExistPids) *QueryPipelineListResponseBody
	GetNonExistPids() *QueryPipelineListResponseBodyNonExistPids
	SetPipelineList(v *QueryPipelineListResponseBodyPipelineList) *QueryPipelineListResponseBody
	GetPipelineList() *QueryPipelineListResponseBodyPipelineList
	SetRequestId(v string) *QueryPipelineListResponseBody
	GetRequestId() *string
}

type QueryPipelineListResponseBody struct {
	// The IDs of MPS queues that do not exist.
	NonExistPids *QueryPipelineListResponseBodyNonExistPids `json:"NonExistPids,omitempty" xml:"NonExistPids,omitempty" type:"Struct"`
	// The MPS queues.
	PipelineList *QueryPipelineListResponseBodyPipelineList `json:"PipelineList,omitempty" xml:"PipelineList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1C538EAA-ACAF-5AD8-B091-A72C63007149
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPipelineListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPipelineListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPipelineListResponseBody) GetNonExistPids() *QueryPipelineListResponseBodyNonExistPids {
	return s.NonExistPids
}

func (s *QueryPipelineListResponseBody) GetPipelineList() *QueryPipelineListResponseBodyPipelineList {
	return s.PipelineList
}

func (s *QueryPipelineListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPipelineListResponseBody) SetNonExistPids(v *QueryPipelineListResponseBodyNonExistPids) *QueryPipelineListResponseBody {
	s.NonExistPids = v
	return s
}

func (s *QueryPipelineListResponseBody) SetPipelineList(v *QueryPipelineListResponseBodyPipelineList) *QueryPipelineListResponseBody {
	s.PipelineList = v
	return s
}

func (s *QueryPipelineListResponseBody) SetRequestId(v string) *QueryPipelineListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPipelineListResponseBody) Validate() error {
	if s.NonExistPids != nil {
		if err := s.NonExistPids.Validate(); err != nil {
			return err
		}
	}
	if s.PipelineList != nil {
		if err := s.PipelineList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPipelineListResponseBodyNonExistPids struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s QueryPipelineListResponseBodyNonExistPids) String() string {
	return dara.Prettify(s)
}

func (s QueryPipelineListResponseBodyNonExistPids) GoString() string {
	return s.String()
}

func (s *QueryPipelineListResponseBodyNonExistPids) GetString_() []*string {
	return s.String_
}

func (s *QueryPipelineListResponseBodyNonExistPids) SetString_(v []*string) *QueryPipelineListResponseBodyNonExistPids {
	s.String_ = v
	return s
}

func (s *QueryPipelineListResponseBodyNonExistPids) Validate() error {
	return dara.Validate(s)
}

type QueryPipelineListResponseBodyPipelineList struct {
	Pipeline []*QueryPipelineListResponseBodyPipelineListPipeline `json:"Pipeline,omitempty" xml:"Pipeline,omitempty" type:"Repeated"`
}

func (s QueryPipelineListResponseBodyPipelineList) String() string {
	return dara.Prettify(s)
}

func (s QueryPipelineListResponseBodyPipelineList) GoString() string {
	return s.String()
}

func (s *QueryPipelineListResponseBodyPipelineList) GetPipeline() []*QueryPipelineListResponseBodyPipelineListPipeline {
	return s.Pipeline
}

func (s *QueryPipelineListResponseBodyPipelineList) SetPipeline(v []*QueryPipelineListResponseBodyPipelineListPipeline) *QueryPipelineListResponseBodyPipelineList {
	s.Pipeline = v
	return s
}

func (s *QueryPipelineListResponseBodyPipelineList) Validate() error {
	if s.Pipeline != nil {
		for _, item := range s.Pipeline {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryPipelineListResponseBodyPipelineListPipeline struct {
	ExtendConfig *QueryPipelineListResponseBodyPipelineListPipelineExtendConfig `json:"ExtendConfig,omitempty" xml:"ExtendConfig,omitempty" type:"Struct"`
	// The ID of the MPS queue.
	//
	// example:
	//
	// d1ce4d3efcb549419193f50f1fcd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the MPS queue.
	//
	// example:
	//
	// example-pipeline-****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Message Service (MNS) configuration.
	NotifyConfig *QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig `json:"NotifyConfig,omitempty" xml:"NotifyConfig,omitempty" type:"Struct"`
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
	// AliyunMTSDefaultRole
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
	// The level of the MPS queue.
	//
	// example:
	//
	// 2
	SpeedLevel *int64 `json:"SpeedLevel,omitempty" xml:"SpeedLevel,omitempty"`
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

func (s QueryPipelineListResponseBodyPipelineListPipeline) String() string {
	return dara.Prettify(s)
}

func (s QueryPipelineListResponseBodyPipelineListPipeline) GoString() string {
	return s.String()
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) GetExtendConfig() *QueryPipelineListResponseBodyPipelineListPipelineExtendConfig {
	return s.ExtendConfig
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) GetId() *string {
	return s.Id
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) GetName() *string {
	return s.Name
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) GetNotifyConfig() *QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig {
	return s.NotifyConfig
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) GetQuotaAllocate() *int64 {
	return s.QuotaAllocate
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) GetRole() *string {
	return s.Role
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) GetSpeed() *string {
	return s.Speed
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) GetSpeedLevel() *int64 {
	return s.SpeedLevel
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) GetState() *string {
	return s.State
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) SetExtendConfig(v *QueryPipelineListResponseBodyPipelineListPipelineExtendConfig) *QueryPipelineListResponseBodyPipelineListPipeline {
	s.ExtendConfig = v
	return s
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) SetId(v string) *QueryPipelineListResponseBodyPipelineListPipeline {
	s.Id = &v
	return s
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) SetName(v string) *QueryPipelineListResponseBodyPipelineListPipeline {
	s.Name = &v
	return s
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) SetNotifyConfig(v *QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig) *QueryPipelineListResponseBodyPipelineListPipeline {
	s.NotifyConfig = v
	return s
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) SetQuotaAllocate(v int64) *QueryPipelineListResponseBodyPipelineListPipeline {
	s.QuotaAllocate = &v
	return s
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) SetRole(v string) *QueryPipelineListResponseBodyPipelineListPipeline {
	s.Role = &v
	return s
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) SetSpeed(v string) *QueryPipelineListResponseBodyPipelineListPipeline {
	s.Speed = &v
	return s
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) SetSpeedLevel(v int64) *QueryPipelineListResponseBodyPipelineListPipeline {
	s.SpeedLevel = &v
	return s
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) SetState(v string) *QueryPipelineListResponseBodyPipelineListPipeline {
	s.State = &v
	return s
}

func (s *QueryPipelineListResponseBodyPipelineListPipeline) Validate() error {
	if s.ExtendConfig != nil {
		if err := s.ExtendConfig.Validate(); err != nil {
			return err
		}
	}
	if s.NotifyConfig != nil {
		if err := s.NotifyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPipelineListResponseBodyPipelineListPipelineExtendConfig struct {
	IsBoostNew                *bool   `json:"IsBoostNew,omitempty" xml:"IsBoostNew,omitempty"`
	MaxMultiSpeed             *int32  `json:"MaxMultiSpeed,omitempty" xml:"MaxMultiSpeed,omitempty"`
	MultiSpeedDowngradePolicy *string `json:"MultiSpeedDowngradePolicy,omitempty" xml:"MultiSpeedDowngradePolicy,omitempty"`
}

func (s QueryPipelineListResponseBodyPipelineListPipelineExtendConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryPipelineListResponseBodyPipelineListPipelineExtendConfig) GoString() string {
	return s.String()
}

func (s *QueryPipelineListResponseBodyPipelineListPipelineExtendConfig) GetIsBoostNew() *bool {
	return s.IsBoostNew
}

func (s *QueryPipelineListResponseBodyPipelineListPipelineExtendConfig) GetMaxMultiSpeed() *int32 {
	return s.MaxMultiSpeed
}

func (s *QueryPipelineListResponseBodyPipelineListPipelineExtendConfig) GetMultiSpeedDowngradePolicy() *string {
	return s.MultiSpeedDowngradePolicy
}

func (s *QueryPipelineListResponseBodyPipelineListPipelineExtendConfig) SetIsBoostNew(v bool) *QueryPipelineListResponseBodyPipelineListPipelineExtendConfig {
	s.IsBoostNew = &v
	return s
}

func (s *QueryPipelineListResponseBodyPipelineListPipelineExtendConfig) SetMaxMultiSpeed(v int32) *QueryPipelineListResponseBodyPipelineListPipelineExtendConfig {
	s.MaxMultiSpeed = &v
	return s
}

func (s *QueryPipelineListResponseBodyPipelineListPipelineExtendConfig) SetMultiSpeedDowngradePolicy(v string) *QueryPipelineListResponseBodyPipelineListPipelineExtendConfig {
	s.MultiSpeedDowngradePolicy = &v
	return s
}

func (s *QueryPipelineListResponseBodyPipelineListPipelineExtendConfig) Validate() error {
	return dara.Validate(s)
}

type QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig struct {
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
	// example1,example2
	MqTopic *string `json:"MqTopic,omitempty" xml:"MqTopic,omitempty"`
	// The name of the queue that is created in MNS.
	//
	// example:
	//
	// example-queue-****
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The name of the topic that is created in MNS.
	//
	// example:
	//
	// example-topic-****
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig) GoString() string {
	return s.String()
}

func (s *QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig) GetMqTag() *string {
	return s.MqTag
}

func (s *QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig) GetMqTopic() *string {
	return s.MqTopic
}

func (s *QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig) GetQueueName() *string {
	return s.QueueName
}

func (s *QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig) GetTopic() *string {
	return s.Topic
}

func (s *QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig) SetMqTag(v string) *QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig {
	s.MqTag = &v
	return s
}

func (s *QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig) SetMqTopic(v string) *QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig {
	s.MqTopic = &v
	return s
}

func (s *QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig) SetQueueName(v string) *QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig {
	s.QueueName = &v
	return s
}

func (s *QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig) SetTopic(v string) *QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig {
	s.Topic = &v
	return s
}

func (s *QueryPipelineListResponseBodyPipelineListPipelineNotifyConfig) Validate() error {
	return dara.Validate(s)
}
