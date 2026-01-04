// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *GetTopicResponseBody
	GetComment() *string
	SetCreateTime(v string) *GetTopicResponseBody
	GetCreateTime() *string
	SetCreator(v string) *GetTopicResponseBody
	GetCreator() *string
	SetEnableSchemaRegistry(v bool) *GetTopicResponseBody
	GetEnableSchemaRegistry() *bool
	SetExpandMode(v bool) *GetTopicResponseBody
	GetExpandMode() *bool
	SetLifecycle(v int32) *GetTopicResponseBody
	GetLifecycle() *int32
	SetProjectName(v string) *GetTopicResponseBody
	GetProjectName() *string
	SetRecordSchema(v string) *GetTopicResponseBody
	GetRecordSchema() *string
	SetRecordType(v string) *GetTopicResponseBody
	GetRecordType() *string
	SetRequestId(v string) *GetTopicResponseBody
	GetRequestId() *string
	SetShardCount(v int32) *GetTopicResponseBody
	GetShardCount() *int32
	SetStorage(v int64) *GetTopicResponseBody
	GetStorage() *int64
	SetSuccess(v bool) *GetTopicResponseBody
	GetSuccess() *bool
	SetTopicName(v string) *GetTopicResponseBody
	GetTopicName() *string
	SetUpdateTime(v string) *GetTopicResponseBody
	GetUpdateTime() *string
}

type GetTopicResponseBody struct {
	// example:
	//
	// test_comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 1724041098000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1397493986831962
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// false
	EnableSchemaRegistry *bool `json:"EnableSchemaRegistry,omitempty" xml:"EnableSchemaRegistry,omitempty"`
	// example:
	//
	// true
	ExpandMode *bool `json:"ExpandMode,omitempty" xml:"ExpandMode,omitempty"`
	// example:
	//
	// 3
	Lifecycle *int32 `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// example:
	//
	// test_project
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
	// A20A7093-8FE0-058C-BE0C-3C8057D5F1A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// example:
	//
	// 12252454
	Storage *int64 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// test_topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// example:
	//
	// 1724041098000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTopicResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicResponseBody) GetComment() *string {
	return s.Comment
}

func (s *GetTopicResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTopicResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *GetTopicResponseBody) GetEnableSchemaRegistry() *bool {
	return s.EnableSchemaRegistry
}

func (s *GetTopicResponseBody) GetExpandMode() *bool {
	return s.ExpandMode
}

func (s *GetTopicResponseBody) GetLifecycle() *int32 {
	return s.Lifecycle
}

func (s *GetTopicResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetTopicResponseBody) GetRecordSchema() *string {
	return s.RecordSchema
}

func (s *GetTopicResponseBody) GetRecordType() *string {
	return s.RecordType
}

func (s *GetTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTopicResponseBody) GetShardCount() *int32 {
	return s.ShardCount
}

func (s *GetTopicResponseBody) GetStorage() *int64 {
	return s.Storage
}

func (s *GetTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTopicResponseBody) GetTopicName() *string {
	return s.TopicName
}

func (s *GetTopicResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetTopicResponseBody) SetComment(v string) *GetTopicResponseBody {
	s.Comment = &v
	return s
}

func (s *GetTopicResponseBody) SetCreateTime(v string) *GetTopicResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetTopicResponseBody) SetCreator(v string) *GetTopicResponseBody {
	s.Creator = &v
	return s
}

func (s *GetTopicResponseBody) SetEnableSchemaRegistry(v bool) *GetTopicResponseBody {
	s.EnableSchemaRegistry = &v
	return s
}

func (s *GetTopicResponseBody) SetExpandMode(v bool) *GetTopicResponseBody {
	s.ExpandMode = &v
	return s
}

func (s *GetTopicResponseBody) SetLifecycle(v int32) *GetTopicResponseBody {
	s.Lifecycle = &v
	return s
}

func (s *GetTopicResponseBody) SetProjectName(v string) *GetTopicResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetTopicResponseBody) SetRecordSchema(v string) *GetTopicResponseBody {
	s.RecordSchema = &v
	return s
}

func (s *GetTopicResponseBody) SetRecordType(v string) *GetTopicResponseBody {
	s.RecordType = &v
	return s
}

func (s *GetTopicResponseBody) SetRequestId(v string) *GetTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicResponseBody) SetShardCount(v int32) *GetTopicResponseBody {
	s.ShardCount = &v
	return s
}

func (s *GetTopicResponseBody) SetStorage(v int64) *GetTopicResponseBody {
	s.Storage = &v
	return s
}

func (s *GetTopicResponseBody) SetSuccess(v bool) *GetTopicResponseBody {
	s.Success = &v
	return s
}

func (s *GetTopicResponseBody) SetTopicName(v string) *GetTopicResponseBody {
	s.TopicName = &v
	return s
}

func (s *GetTopicResponseBody) SetUpdateTime(v string) *GetTopicResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetTopicResponseBody) Validate() error {
	return dara.Validate(s)
}
