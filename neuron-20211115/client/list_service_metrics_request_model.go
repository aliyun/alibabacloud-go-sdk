// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListServiceMetricsRequest
	GetEndTime() *string
	SetEnv(v string) *ListServiceMetricsRequest
	GetEnv() *string
	SetGroupId(v int64) *ListServiceMetricsRequest
	GetGroupId() *int64
	SetIntervalInSec(v int32) *ListServiceMetricsRequest
	GetIntervalInSec() *int32
	SetIp(v string) *ListServiceMetricsRequest
	GetIp() *string
	SetMeasures(v string) *ListServiceMetricsRequest
	GetMeasures() *string
	SetServiceGroupId(v int64) *ListServiceMetricsRequest
	GetServiceGroupId() *int64
	SetStartTime(v string) *ListServiceMetricsRequest
	GetStartTime() *string
	SetTopicId(v int64) *ListServiceMetricsRequest
	GetTopicId() *int64
	SetType(v string) *ListServiceMetricsRequest
	GetType() *string
}

type ListServiceMetricsRequest struct {
	// This parameter is required.
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	Env     *string `json:"env,omitempty" xml:"env,omitempty"`
	GroupId *int64  `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// This parameter is required.
	IntervalInSec  *int32  `json:"intervalInSec,omitempty" xml:"intervalInSec,omitempty"`
	Ip             *string `json:"ip,omitempty" xml:"ip,omitempty"`
	Measures       *string `json:"measures,omitempty" xml:"measures,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	// This parameter is required.
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TopicId   *int64  `json:"topicId,omitempty" xml:"topicId,omitempty"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListServiceMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceMetricsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListServiceMetricsRequest) GetEnv() *string {
	return s.Env
}

func (s *ListServiceMetricsRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListServiceMetricsRequest) GetIntervalInSec() *int32 {
	return s.IntervalInSec
}

func (s *ListServiceMetricsRequest) GetIp() *string {
	return s.Ip
}

func (s *ListServiceMetricsRequest) GetMeasures() *string {
	return s.Measures
}

func (s *ListServiceMetricsRequest) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *ListServiceMetricsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListServiceMetricsRequest) GetTopicId() *int64 {
	return s.TopicId
}

func (s *ListServiceMetricsRequest) GetType() *string {
	return s.Type
}

func (s *ListServiceMetricsRequest) SetEndTime(v string) *ListServiceMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ListServiceMetricsRequest) SetEnv(v string) *ListServiceMetricsRequest {
	s.Env = &v
	return s
}

func (s *ListServiceMetricsRequest) SetGroupId(v int64) *ListServiceMetricsRequest {
	s.GroupId = &v
	return s
}

func (s *ListServiceMetricsRequest) SetIntervalInSec(v int32) *ListServiceMetricsRequest {
	s.IntervalInSec = &v
	return s
}

func (s *ListServiceMetricsRequest) SetIp(v string) *ListServiceMetricsRequest {
	s.Ip = &v
	return s
}

func (s *ListServiceMetricsRequest) SetMeasures(v string) *ListServiceMetricsRequest {
	s.Measures = &v
	return s
}

func (s *ListServiceMetricsRequest) SetServiceGroupId(v int64) *ListServiceMetricsRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *ListServiceMetricsRequest) SetStartTime(v string) *ListServiceMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *ListServiceMetricsRequest) SetTopicId(v int64) *ListServiceMetricsRequest {
	s.TopicId = &v
	return s
}

func (s *ListServiceMetricsRequest) SetType(v string) *ListServiceMetricsRequest {
	s.Type = &v
	return s
}

func (s *ListServiceMetricsRequest) Validate() error {
	return dara.Validate(s)
}
