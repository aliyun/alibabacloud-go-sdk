// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *PutRecordsRequest
	GetProjectName() *string
	SetRecords(v []*PutRecordsRequestRecords) *PutRecordsRequest
	GetRecords() []*PutRecordsRequestRecords
	SetShardId(v string) *PutRecordsRequest
	GetShardId() *string
	SetTopicName(v string) *PutRecordsRequest
	GetTopicName() *string
}

type PutRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// This parameter is required.
	Records []*PutRecordsRequestRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
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

func (s PutRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s PutRecordsRequest) GoString() string {
	return s.String()
}

func (s *PutRecordsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *PutRecordsRequest) GetRecords() []*PutRecordsRequestRecords {
	return s.Records
}

func (s *PutRecordsRequest) GetShardId() *string {
	return s.ShardId
}

func (s *PutRecordsRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *PutRecordsRequest) SetProjectName(v string) *PutRecordsRequest {
	s.ProjectName = &v
	return s
}

func (s *PutRecordsRequest) SetRecords(v []*PutRecordsRequestRecords) *PutRecordsRequest {
	s.Records = v
	return s
}

func (s *PutRecordsRequest) SetShardId(v string) *PutRecordsRequest {
	s.ShardId = &v
	return s
}

func (s *PutRecordsRequest) SetTopicName(v string) *PutRecordsRequest {
	s.TopicName = &v
	return s
}

func (s *PutRecordsRequest) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutRecordsRequestRecords struct {
	// example:
	//
	// {"key1":"val1","key2":"val2"}
	Attributes map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["aa", "bb", "12", "12.34"]
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s PutRecordsRequestRecords) String() string {
	return dara.Prettify(s)
}

func (s PutRecordsRequestRecords) GoString() string {
	return s.String()
}

func (s *PutRecordsRequestRecords) GetAttributes() map[string]*string {
	return s.Attributes
}

func (s *PutRecordsRequestRecords) GetData() []*string {
	return s.Data
}

func (s *PutRecordsRequestRecords) SetAttributes(v map[string]*string) *PutRecordsRequestRecords {
	s.Attributes = v
	return s
}

func (s *PutRecordsRequestRecords) SetData(v []*string) *PutRecordsRequestRecords {
	s.Data = v
	return s
}

func (s *PutRecordsRequestRecords) Validate() error {
	return dara.Validate(s)
}
