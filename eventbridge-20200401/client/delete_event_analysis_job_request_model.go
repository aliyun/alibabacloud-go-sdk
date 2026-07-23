// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventAnalysisJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceResource(v *DeleteEventAnalysisJobRequestSourceResource) *DeleteEventAnalysisJobRequest
	GetSourceResource() *DeleteEventAnalysisJobRequestSourceResource
}

type DeleteEventAnalysisJobRequest struct {
	// The identifier of the source resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Kafka":{"RegionId":"cn-hangzhou","InstanceId":"alikafka_post-cn-xxx","Topic":"my_topic"}}
	SourceResource *DeleteEventAnalysisJobRequestSourceResource `json:"SourceResource,omitempty" xml:"SourceResource,omitempty" type:"Struct"`
}

func (s DeleteEventAnalysisJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventAnalysisJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventAnalysisJobRequest) GetSourceResource() *DeleteEventAnalysisJobRequestSourceResource {
	return s.SourceResource
}

func (s *DeleteEventAnalysisJobRequest) SetSourceResource(v *DeleteEventAnalysisJobRequestSourceResource) *DeleteEventAnalysisJobRequest {
	s.SourceResource = v
	return s
}

func (s *DeleteEventAnalysisJobRequest) Validate() error {
	if s.SourceResource != nil {
		if err := s.SourceResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteEventAnalysisJobRequestSourceResource struct {
	// The Kafka data source.
	Kafka *DeleteEventAnalysisJobRequestSourceResourceKafka `json:"Kafka,omitempty" xml:"Kafka,omitempty" type:"Struct"`
	// The RocketMQ data source.
	RocketMQ *DeleteEventAnalysisJobRequestSourceResourceRocketMQ `json:"RocketMQ,omitempty" xml:"RocketMQ,omitempty" type:"Struct"`
}

func (s DeleteEventAnalysisJobRequestSourceResource) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventAnalysisJobRequestSourceResource) GoString() string {
	return s.String()
}

func (s *DeleteEventAnalysisJobRequestSourceResource) GetKafka() *DeleteEventAnalysisJobRequestSourceResourceKafka {
	return s.Kafka
}

func (s *DeleteEventAnalysisJobRequestSourceResource) GetRocketMQ() *DeleteEventAnalysisJobRequestSourceResourceRocketMQ {
	return s.RocketMQ
}

func (s *DeleteEventAnalysisJobRequestSourceResource) SetKafka(v *DeleteEventAnalysisJobRequestSourceResourceKafka) *DeleteEventAnalysisJobRequestSourceResource {
	s.Kafka = v
	return s
}

func (s *DeleteEventAnalysisJobRequestSourceResource) SetRocketMQ(v *DeleteEventAnalysisJobRequestSourceResourceRocketMQ) *DeleteEventAnalysisJobRequestSourceResource {
	s.RocketMQ = v
	return s
}

func (s *DeleteEventAnalysisJobRequestSourceResource) Validate() error {
	if s.Kafka != nil {
		if err := s.Kafka.Validate(); err != nil {
			return err
		}
	}
	if s.RocketMQ != nil {
		if err := s.RocketMQ.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteEventAnalysisJobRequestSourceResourceKafka struct {
	// The instance ID of the Kafka instance.
	//
	// example:
	//
	// alikafka_post-cn-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region of the Kafka instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the Kafka topic.
	//
	// example:
	//
	// my_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s DeleteEventAnalysisJobRequestSourceResourceKafka) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventAnalysisJobRequestSourceResourceKafka) GoString() string {
	return s.String()
}

func (s *DeleteEventAnalysisJobRequestSourceResourceKafka) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteEventAnalysisJobRequestSourceResourceKafka) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEventAnalysisJobRequestSourceResourceKafka) GetTopic() *string {
	return s.Topic
}

func (s *DeleteEventAnalysisJobRequestSourceResourceKafka) SetInstanceId(v string) *DeleteEventAnalysisJobRequestSourceResourceKafka {
	s.InstanceId = &v
	return s
}

func (s *DeleteEventAnalysisJobRequestSourceResourceKafka) SetRegionId(v string) *DeleteEventAnalysisJobRequestSourceResourceKafka {
	s.RegionId = &v
	return s
}

func (s *DeleteEventAnalysisJobRequestSourceResourceKafka) SetTopic(v string) *DeleteEventAnalysisJobRequestSourceResourceKafka {
	s.Topic = &v
	return s
}

func (s *DeleteEventAnalysisJobRequestSourceResourceKafka) Validate() error {
	return dara.Validate(s)
}

type DeleteEventAnalysisJobRequestSourceResourceRocketMQ struct {
	// The instance ID of the RocketMQ instance.
	//
	// example:
	//
	// rmq-cn-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the RocketMQ instance.
	//
	// example:
	//
	// Cloud_5
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The region of the RocketMQ instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the RocketMQ topic.
	//
	// example:
	//
	// my_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s DeleteEventAnalysisJobRequestSourceResourceRocketMQ) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventAnalysisJobRequestSourceResourceRocketMQ) GoString() string {
	return s.String()
}

func (s *DeleteEventAnalysisJobRequestSourceResourceRocketMQ) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteEventAnalysisJobRequestSourceResourceRocketMQ) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DeleteEventAnalysisJobRequestSourceResourceRocketMQ) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEventAnalysisJobRequestSourceResourceRocketMQ) GetTopic() *string {
	return s.Topic
}

func (s *DeleteEventAnalysisJobRequestSourceResourceRocketMQ) SetInstanceId(v string) *DeleteEventAnalysisJobRequestSourceResourceRocketMQ {
	s.InstanceId = &v
	return s
}

func (s *DeleteEventAnalysisJobRequestSourceResourceRocketMQ) SetInstanceType(v string) *DeleteEventAnalysisJobRequestSourceResourceRocketMQ {
	s.InstanceType = &v
	return s
}

func (s *DeleteEventAnalysisJobRequestSourceResourceRocketMQ) SetRegionId(v string) *DeleteEventAnalysisJobRequestSourceResourceRocketMQ {
	s.RegionId = &v
	return s
}

func (s *DeleteEventAnalysisJobRequestSourceResourceRocketMQ) SetTopic(v string) *DeleteEventAnalysisJobRequestSourceResourceRocketMQ {
	s.Topic = &v
	return s
}

func (s *DeleteEventAnalysisJobRequestSourceResourceRocketMQ) Validate() error {
	return dara.Validate(s)
}
