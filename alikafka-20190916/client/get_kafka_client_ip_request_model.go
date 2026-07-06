// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKafkaClientIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetKafkaClientIpRequest
	GetEndTime() *int64
	SetGroup(v string) *GetKafkaClientIpRequest
	GetGroup() *string
	SetInstanceId(v string) *GetKafkaClientIpRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetKafkaClientIpRequest
	GetRegionId() *string
	SetStartTime(v int64) *GetKafkaClientIpRequest
	GetStartTime() *int64
	SetTopic(v string) *GetKafkaClientIpRequest
	GetTopic() *string
	SetType(v string) *GetKafkaClientIpRequest
	GetType() *string
}

type GetKafkaClientIpRequest struct {
	// End time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1716343502000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Consumer group ID.
	//
	// > This parameter is required when Type is set to byGroup.
	//
	// example:
	//
	// group_name
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1716343501000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Topic name.
	//
	// > - This parameter is required when Type is set to byTopic.
	//
	// example:
	//
	// topic_name
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The type of client IP query. You can choose from three methods.
	//
	// - byInstance: Query client IPs for the instance within the specified time range.
	//
	// - byTopic: Query client IPs for the topic within the specified time range.
	//
	// - byGroup: Query client IPs for the group within the specified time range.
	//
	// This parameter is required.
	//
	// example:
	//
	// byInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetKafkaClientIpRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKafkaClientIpRequest) GoString() string {
	return s.String()
}

func (s *GetKafkaClientIpRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetKafkaClientIpRequest) GetGroup() *string {
	return s.Group
}

func (s *GetKafkaClientIpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetKafkaClientIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetKafkaClientIpRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetKafkaClientIpRequest) GetTopic() *string {
	return s.Topic
}

func (s *GetKafkaClientIpRequest) GetType() *string {
	return s.Type
}

func (s *GetKafkaClientIpRequest) SetEndTime(v int64) *GetKafkaClientIpRequest {
	s.EndTime = &v
	return s
}

func (s *GetKafkaClientIpRequest) SetGroup(v string) *GetKafkaClientIpRequest {
	s.Group = &v
	return s
}

func (s *GetKafkaClientIpRequest) SetInstanceId(v string) *GetKafkaClientIpRequest {
	s.InstanceId = &v
	return s
}

func (s *GetKafkaClientIpRequest) SetRegionId(v string) *GetKafkaClientIpRequest {
	s.RegionId = &v
	return s
}

func (s *GetKafkaClientIpRequest) SetStartTime(v int64) *GetKafkaClientIpRequest {
	s.StartTime = &v
	return s
}

func (s *GetKafkaClientIpRequest) SetTopic(v string) *GetKafkaClientIpRequest {
	s.Topic = &v
	return s
}

func (s *GetKafkaClientIpRequest) SetType(v string) *GetKafkaClientIpRequest {
	s.Type = &v
	return s
}

func (s *GetKafkaClientIpRequest) Validate() error {
	return dara.Validate(s)
}
