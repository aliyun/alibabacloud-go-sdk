// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *QueryMessageRequest
	GetBeginTime() *int64
	SetInstanceId(v string) *QueryMessageRequest
	GetInstanceId() *string
	SetOffset(v string) *QueryMessageRequest
	GetOffset() *string
	SetPartition(v string) *QueryMessageRequest
	GetPartition() *string
	SetQueryType(v string) *QueryMessageRequest
	GetQueryType() *string
	SetRegionId(v string) *QueryMessageRequest
	GetRegionId() *string
	SetTopic(v string) *QueryMessageRequest
	GetTopic() *string
}

type QueryMessageRequest struct {
	// The beginning of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1672410180000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The consumer offset of the partition.
	//
	// example:
	//
	// 100
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The partition ID.
	//
	// example:
	//
	// 0
	Partition *string `json:"Partition,omitempty" xml:"Partition,omitempty"`
	// The query type. Valid values:
	//
	// 	- byOffset: queries messages by offset. If you select this value, you must configure Partition and Offset.
	//
	// 	- byTimestamp: queries messages by time. If you select this value, you must configure BeginTime.
	//
	// This parameter is required.
	//
	// example:
	//
	// byTimestamp
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The ID of the region where the resource resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The topic name.
	//
	// This parameter is required.
	//
	// example:
	//
	// testkafka
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s QueryMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageRequest) GoString() string {
	return s.String()
}

func (s *QueryMessageRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *QueryMessageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryMessageRequest) GetOffset() *string {
	return s.Offset
}

func (s *QueryMessageRequest) GetPartition() *string {
	return s.Partition
}

func (s *QueryMessageRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *QueryMessageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryMessageRequest) GetTopic() *string {
	return s.Topic
}

func (s *QueryMessageRequest) SetBeginTime(v int64) *QueryMessageRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMessageRequest) SetInstanceId(v string) *QueryMessageRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMessageRequest) SetOffset(v string) *QueryMessageRequest {
	s.Offset = &v
	return s
}

func (s *QueryMessageRequest) SetPartition(v string) *QueryMessageRequest {
	s.Partition = &v
	return s
}

func (s *QueryMessageRequest) SetQueryType(v string) *QueryMessageRequest {
	s.QueryType = &v
	return s
}

func (s *QueryMessageRequest) SetRegionId(v string) *QueryMessageRequest {
	s.RegionId = &v
	return s
}

func (s *QueryMessageRequest) SetTopic(v string) *QueryMessageRequest {
	s.Topic = &v
	return s
}

func (s *QueryMessageRequest) Validate() error {
	return dara.Validate(s)
}
