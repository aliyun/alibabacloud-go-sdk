// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutRecordsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *PutRecordsShrinkRequest
	GetProjectName() *string
	SetRecordsShrink(v string) *PutRecordsShrinkRequest
	GetRecordsShrink() *string
	SetShardId(v string) *PutRecordsShrinkRequest
	GetShardId() *string
	SetTopicName(v string) *PutRecordsShrinkRequest
	GetTopicName() *string
}

type PutRecordsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// This parameter is required.
	RecordsShrink *string `json:"Records,omitempty" xml:"Records,omitempty"`
	// example:
	//
	// 7
	ShardId *string `json:"ShardId,omitempty" xml:"ShardId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s PutRecordsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PutRecordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *PutRecordsShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *PutRecordsShrinkRequest) GetRecordsShrink() *string {
	return s.RecordsShrink
}

func (s *PutRecordsShrinkRequest) GetShardId() *string {
	return s.ShardId
}

func (s *PutRecordsShrinkRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *PutRecordsShrinkRequest) SetProjectName(v string) *PutRecordsShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *PutRecordsShrinkRequest) SetRecordsShrink(v string) *PutRecordsShrinkRequest {
	s.RecordsShrink = &v
	return s
}

func (s *PutRecordsShrinkRequest) SetShardId(v string) *PutRecordsShrinkRequest {
	s.ShardId = &v
	return s
}

func (s *PutRecordsShrinkRequest) SetTopicName(v string) *PutRecordsShrinkRequest {
	s.TopicName = &v
	return s
}

func (s *PutRecordsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
