// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchPipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *SearchPipelineResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *SearchPipelineResponseBody
	GetPageSize() *int64
	SetPipelineList(v *SearchPipelineResponseBodyPipelineList) *SearchPipelineResponseBody
	GetPipelineList() *SearchPipelineResponseBodyPipelineList
	SetRequestId(v string) *SearchPipelineResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *SearchPipelineResponseBody
	GetTotalCount() *int64
}

type SearchPipelineResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The MPS queues.
	PipelineList *SearchPipelineResponseBodyPipelineList `json:"PipelineList,omitempty" xml:"PipelineList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 338CA33A-AE83-5DF4-B6F2-C6D3ED8143F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchPipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *SearchPipelineResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *SearchPipelineResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *SearchPipelineResponseBody) GetPipelineList() *SearchPipelineResponseBodyPipelineList {
	return s.PipelineList
}

func (s *SearchPipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchPipelineResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SearchPipelineResponseBody) SetPageNumber(v int64) *SearchPipelineResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchPipelineResponseBody) SetPageSize(v int64) *SearchPipelineResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchPipelineResponseBody) SetPipelineList(v *SearchPipelineResponseBodyPipelineList) *SearchPipelineResponseBody {
	s.PipelineList = v
	return s
}

func (s *SearchPipelineResponseBody) SetRequestId(v string) *SearchPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchPipelineResponseBody) SetTotalCount(v int64) *SearchPipelineResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchPipelineResponseBody) Validate() error {
	if s.PipelineList != nil {
		if err := s.PipelineList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchPipelineResponseBodyPipelineList struct {
	Pipeline []*SearchPipelineResponseBodyPipelineListPipeline `json:"Pipeline,omitempty" xml:"Pipeline,omitempty" type:"Repeated"`
}

func (s SearchPipelineResponseBodyPipelineList) String() string {
	return dara.Prettify(s)
}

func (s SearchPipelineResponseBodyPipelineList) GoString() string {
	return s.String()
}

func (s *SearchPipelineResponseBodyPipelineList) GetPipeline() []*SearchPipelineResponseBodyPipelineListPipeline {
	return s.Pipeline
}

func (s *SearchPipelineResponseBodyPipelineList) SetPipeline(v []*SearchPipelineResponseBodyPipelineListPipeline) *SearchPipelineResponseBodyPipelineList {
	s.Pipeline = v
	return s
}

func (s *SearchPipelineResponseBodyPipelineList) Validate() error {
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

type SearchPipelineResponseBodyPipelineListPipeline struct {
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
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
	NotifyConfig *SearchPipelineResponseBodyPipelineListPipelineNotifyConfig `json:"NotifyConfig,omitempty" xml:"NotifyConfig,omitempty" type:"Struct"`
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
	// 1
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

func (s SearchPipelineResponseBodyPipelineListPipeline) String() string {
	return dara.Prettify(s)
}

func (s SearchPipelineResponseBodyPipelineListPipeline) GoString() string {
	return s.String()
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) GetId() *string {
	return s.Id
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) GetName() *string {
	return s.Name
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) GetNotifyConfig() *SearchPipelineResponseBodyPipelineListPipelineNotifyConfig {
	return s.NotifyConfig
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) GetQuotaAllocate() *int64 {
	return s.QuotaAllocate
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) GetRole() *string {
	return s.Role
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) GetSpeed() *string {
	return s.Speed
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) GetSpeedLevel() *int64 {
	return s.SpeedLevel
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) GetState() *string {
	return s.State
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) SetCreationTime(v string) *SearchPipelineResponseBodyPipelineListPipeline {
	s.CreationTime = &v
	return s
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) SetId(v string) *SearchPipelineResponseBodyPipelineListPipeline {
	s.Id = &v
	return s
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) SetName(v string) *SearchPipelineResponseBodyPipelineListPipeline {
	s.Name = &v
	return s
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) SetNotifyConfig(v *SearchPipelineResponseBodyPipelineListPipelineNotifyConfig) *SearchPipelineResponseBodyPipelineListPipeline {
	s.NotifyConfig = v
	return s
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) SetQuotaAllocate(v int64) *SearchPipelineResponseBodyPipelineListPipeline {
	s.QuotaAllocate = &v
	return s
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) SetRole(v string) *SearchPipelineResponseBodyPipelineListPipeline {
	s.Role = &v
	return s
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) SetSpeed(v string) *SearchPipelineResponseBodyPipelineListPipeline {
	s.Speed = &v
	return s
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) SetSpeedLevel(v int64) *SearchPipelineResponseBodyPipelineListPipeline {
	s.SpeedLevel = &v
	return s
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) SetState(v string) *SearchPipelineResponseBodyPipelineListPipeline {
	s.State = &v
	return s
}

func (s *SearchPipelineResponseBodyPipelineListPipeline) Validate() error {
	if s.NotifyConfig != nil {
		if err := s.NotifyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchPipelineResponseBodyPipelineListPipelineNotifyConfig struct {
	// The tags.
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

func (s SearchPipelineResponseBodyPipelineListPipelineNotifyConfig) String() string {
	return dara.Prettify(s)
}

func (s SearchPipelineResponseBodyPipelineListPipelineNotifyConfig) GoString() string {
	return s.String()
}

func (s *SearchPipelineResponseBodyPipelineListPipelineNotifyConfig) GetMqTag() *string {
	return s.MqTag
}

func (s *SearchPipelineResponseBodyPipelineListPipelineNotifyConfig) GetMqTopic() *string {
	return s.MqTopic
}

func (s *SearchPipelineResponseBodyPipelineListPipelineNotifyConfig) GetQueueName() *string {
	return s.QueueName
}

func (s *SearchPipelineResponseBodyPipelineListPipelineNotifyConfig) GetTopic() *string {
	return s.Topic
}

func (s *SearchPipelineResponseBodyPipelineListPipelineNotifyConfig) SetMqTag(v string) *SearchPipelineResponseBodyPipelineListPipelineNotifyConfig {
	s.MqTag = &v
	return s
}

func (s *SearchPipelineResponseBodyPipelineListPipelineNotifyConfig) SetMqTopic(v string) *SearchPipelineResponseBodyPipelineListPipelineNotifyConfig {
	s.MqTopic = &v
	return s
}

func (s *SearchPipelineResponseBodyPipelineListPipelineNotifyConfig) SetQueueName(v string) *SearchPipelineResponseBodyPipelineListPipelineNotifyConfig {
	s.QueueName = &v
	return s
}

func (s *SearchPipelineResponseBodyPipelineListPipelineNotifyConfig) SetTopic(v string) *SearchPipelineResponseBodyPipelineListPipelineNotifyConfig {
	s.Topic = &v
	return s
}

func (s *SearchPipelineResponseBodyPipelineListPipelineNotifyConfig) Validate() error {
	return dara.Validate(s)
}
