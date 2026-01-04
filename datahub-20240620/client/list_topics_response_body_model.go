// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v *ListTopicsResponseBodyList) *ListTopicsResponseBody
	GetList() *ListTopicsResponseBodyList
	SetMaxResults(v int32) *ListTopicsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTopicsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTopicsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTopicsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListTopicsResponseBody
	GetTotalCount() *int32
}

type ListTopicsResponseBody struct {
	List *ListTopicsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 9892074a2a89600ae4b0d5a34fb99a3f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// A20A7093-8FE0-058C-BE0C-3C8057D5F1A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTopicsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTopicsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTopicsResponseBody) GetList() *ListTopicsResponseBodyList {
	return s.List
}

func (s *ListTopicsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTopicsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTopicsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTopicsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTopicsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTopicsResponseBody) SetList(v *ListTopicsResponseBodyList) *ListTopicsResponseBody {
	s.List = v
	return s
}

func (s *ListTopicsResponseBody) SetMaxResults(v int32) *ListTopicsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTopicsResponseBody) SetNextToken(v string) *ListTopicsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTopicsResponseBody) SetRequestId(v string) *ListTopicsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTopicsResponseBody) SetSuccess(v bool) *ListTopicsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTopicsResponseBody) SetTotalCount(v int32) *ListTopicsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTopicsResponseBody) Validate() error {
	if s.List != nil {
		if err := s.List.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTopicsResponseBodyList struct {
	Topic []*ListTopicsResponseBodyListTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Repeated"`
}

func (s ListTopicsResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListTopicsResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListTopicsResponseBodyList) GetTopic() []*ListTopicsResponseBodyListTopic {
	return s.Topic
}

func (s *ListTopicsResponseBodyList) SetTopic(v []*ListTopicsResponseBodyListTopic) *ListTopicsResponseBodyList {
	s.Topic = v
	return s
}

func (s *ListTopicsResponseBodyList) Validate() error {
	if s.Topic != nil {
		for _, item := range s.Topic {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTopicsResponseBodyListTopic struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 1753346106000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 276887103073464052
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// false
	EnableSchemaRegistry *string `json:"EnableSchemaRegistry,omitempty" xml:"EnableSchemaRegistry,omitempty"`
	// example:
	//
	// true
	ExpandMode *string `json:"ExpandMode,omitempty" xml:"ExpandMode,omitempty"`
	// example:
	//
	// 3
	Lifecycle *int32 `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// example:
	//
	// poc_test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// [{\\"Type\\":\\"STRING\\",\\"AllowNull\\":true,\\"Name\\":\\"str\\"},{\\"Type\\":\\"STRING\\",\\"AllowNull\\":true,\\"Name\\":\\"dt\\"}]
	RecordSchema *string `json:"RecordSchema,omitempty" xml:"RecordSchema,omitempty"`
	// example:
	//
	// TUPLE
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// example:
	//
	// 1
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// example:
	//
	// 10000
	Storage *int64 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// example:
	//
	// ods_bio_safety_env_disinfection
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// example:
	//
	// 1753346106000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListTopicsResponseBodyListTopic) String() string {
	return dara.Prettify(s)
}

func (s ListTopicsResponseBodyListTopic) GoString() string {
	return s.String()
}

func (s *ListTopicsResponseBodyListTopic) GetComment() *string {
	return s.Comment
}

func (s *ListTopicsResponseBodyListTopic) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListTopicsResponseBodyListTopic) GetCreator() *string {
	return s.Creator
}

func (s *ListTopicsResponseBodyListTopic) GetEnableSchemaRegistry() *string {
	return s.EnableSchemaRegistry
}

func (s *ListTopicsResponseBodyListTopic) GetExpandMode() *string {
	return s.ExpandMode
}

func (s *ListTopicsResponseBodyListTopic) GetLifecycle() *int32 {
	return s.Lifecycle
}

func (s *ListTopicsResponseBodyListTopic) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListTopicsResponseBodyListTopic) GetRecordSchema() *string {
	return s.RecordSchema
}

func (s *ListTopicsResponseBodyListTopic) GetRecordType() *string {
	return s.RecordType
}

func (s *ListTopicsResponseBodyListTopic) GetShardCount() *int32 {
	return s.ShardCount
}

func (s *ListTopicsResponseBodyListTopic) GetStorage() *int64 {
	return s.Storage
}

func (s *ListTopicsResponseBodyListTopic) GetTopicName() *string {
	return s.TopicName
}

func (s *ListTopicsResponseBodyListTopic) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListTopicsResponseBodyListTopic) SetComment(v string) *ListTopicsResponseBodyListTopic {
	s.Comment = &v
	return s
}

func (s *ListTopicsResponseBodyListTopic) SetCreateTime(v int64) *ListTopicsResponseBodyListTopic {
	s.CreateTime = &v
	return s
}

func (s *ListTopicsResponseBodyListTopic) SetCreator(v string) *ListTopicsResponseBodyListTopic {
	s.Creator = &v
	return s
}

func (s *ListTopicsResponseBodyListTopic) SetEnableSchemaRegistry(v string) *ListTopicsResponseBodyListTopic {
	s.EnableSchemaRegistry = &v
	return s
}

func (s *ListTopicsResponseBodyListTopic) SetExpandMode(v string) *ListTopicsResponseBodyListTopic {
	s.ExpandMode = &v
	return s
}

func (s *ListTopicsResponseBodyListTopic) SetLifecycle(v int32) *ListTopicsResponseBodyListTopic {
	s.Lifecycle = &v
	return s
}

func (s *ListTopicsResponseBodyListTopic) SetProjectName(v string) *ListTopicsResponseBodyListTopic {
	s.ProjectName = &v
	return s
}

func (s *ListTopicsResponseBodyListTopic) SetRecordSchema(v string) *ListTopicsResponseBodyListTopic {
	s.RecordSchema = &v
	return s
}

func (s *ListTopicsResponseBodyListTopic) SetRecordType(v string) *ListTopicsResponseBodyListTopic {
	s.RecordType = &v
	return s
}

func (s *ListTopicsResponseBodyListTopic) SetShardCount(v int32) *ListTopicsResponseBodyListTopic {
	s.ShardCount = &v
	return s
}

func (s *ListTopicsResponseBodyListTopic) SetStorage(v int64) *ListTopicsResponseBodyListTopic {
	s.Storage = &v
	return s
}

func (s *ListTopicsResponseBodyListTopic) SetTopicName(v string) *ListTopicsResponseBodyListTopic {
	s.TopicName = &v
	return s
}

func (s *ListTopicsResponseBodyListTopic) SetUpdateTime(v int64) *ListTopicsResponseBodyListTopic {
	s.UpdateTime = &v
	return s
}

func (s *ListTopicsResponseBodyListTopic) Validate() error {
	return dara.Validate(s)
}
