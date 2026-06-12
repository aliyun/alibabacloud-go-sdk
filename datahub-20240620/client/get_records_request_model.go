// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *GetRecordsRequest
	GetProjectName() *string
	SetShardId(v string) *GetRecordsRequest
	GetShardId() *string
	SetStartTime(v int64) *GetRecordsRequest
	GetStartTime() *int64
	SetTopicName(v string) *GetRecordsRequest
	GetTopicName() *string
}

type GetRecordsRequest struct {
	// The project name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The shard ID from which to read data.
	//
	// example:
	//
	// 7
	ShardId *string `json:"ShardId,omitempty" xml:"ShardId,omitempty"`
	// The start time for reading data. By default, data is read starting from 5 minutes before the current time.
	//
	// example:
	//
	// 1769065251123
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The topic name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRecordsRequest) GoString() string {
	return s.String()
}

func (s *GetRecordsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetRecordsRequest) GetShardId() *string {
	return s.ShardId
}

func (s *GetRecordsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetRecordsRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *GetRecordsRequest) SetProjectName(v string) *GetRecordsRequest {
	s.ProjectName = &v
	return s
}

func (s *GetRecordsRequest) SetShardId(v string) *GetRecordsRequest {
	s.ShardId = &v
	return s
}

func (s *GetRecordsRequest) SetStartTime(v int64) *GetRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *GetRecordsRequest) SetTopicName(v string) *GetRecordsRequest {
	s.TopicName = &v
	return s
}

func (s *GetRecordsRequest) Validate() error {
	return dara.Validate(s)
}
