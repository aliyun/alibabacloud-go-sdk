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
	// The end of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1716343502000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the consumer group.
	//
	// >  This parameter is required only if you set Type to byGroup.
	//
	// example:
	//
	// group_name
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1716343501000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The topic name.
	//
	// >
	//
	// 	- This parameter is required only if you set Type to byTopic.
	//
	// example:
	//
	// topic_name
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The query method that you want to use to query the client IP addresses. Valid values:
	//
	// 	- byInstance: queries the IP addresses of the clients that are connected to the instance within a specific period of time.
	//
	// 	- byTopic: queries the IP addresses of the clients that are connected to a specific topic on the instance within a specific period of time.
	//
	// 	- byGroup: queries the IP addresses of the clients that are connected to a specific group on the instance within a specific period of time.
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
