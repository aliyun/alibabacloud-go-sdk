// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateTopicRequest
	GetComment() *string
	SetEnableSchemaRegistry(v bool) *CreateTopicRequest
	GetEnableSchemaRegistry() *bool
	SetExpandMode(v bool) *CreateTopicRequest
	GetExpandMode() *bool
	SetLifecycle(v int32) *CreateTopicRequest
	GetLifecycle() *int32
	SetProjectName(v string) *CreateTopicRequest
	GetProjectName() *string
	SetRecordSchema(v string) *CreateTopicRequest
	GetRecordSchema() *string
	SetRecordType(v string) *CreateTopicRequest
	GetRecordType() *string
	SetShardCount(v int32) *CreateTopicRequest
	GetShardCount() *int32
	SetTopicName(v string) *CreateTopicRequest
	GetTopicName() *string
}

type CreateTopicRequest struct {
	// The description of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Specifies whether to enable multi-version schema. After this feature is enabled, a topic can have multiple schemas. You can select one of the schemas for writing. The consumer automatically parses each record based on the version tag. If the schema for the corresponding version has been deleted, parsing fails.
	//
	// > Enabling multi-version schema has the following impacts:
	//
	// 1. You can no longer use the appendFields operation.
	//
	//  2. You can create, delete, modify, and query schemas.
	//
	//  3. Connectors are created by using the schema of the latest version.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	EnableSchemaRegistry *bool `json:"EnableSchemaRegistry,omitempty" xml:"EnableSchemaRegistry,omitempty"`
	// The expansion mode of the topic. After the expansion mode is enabled, shards support horizontal scaling and no longer support merge or split operations. The number of shards can only increase and cannot decrease. After this mode is enabled, you can consume the current topic by using Kafka.
	//
	// example:
	//
	// true
	ExpandMode *bool `json:"ExpandMode,omitempty" xml:"ExpandMode,omitempty"`
	// The lifecycle of the topic. Unit: days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Lifecycle *int32 `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// The project name.
	//
	// This parameter is required.
	//
	// example:
	//
	// xiaowutest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The schema table structure.
	//
	// example:
	//
	// "{\\"fields\\":[{\\"name\\":\\"field_init\\",\\"type\\":\\"STRING\\",\\"notnull\\":\\"false\\"}]}"
	RecordSchema *string `json:"RecordSchema,omitempty" xml:"RecordSchema,omitempty"`
	// The topic type. Valid values:
	//
	// 1. Blob: supports writing a block of binary data as a single record.
	//
	// 1. Tuple: supports database-like records where each record contains multiple columns. You must specify RecordSchema because data is transmitted over the network as strings and requires a schema to convert the data into the corresponding types.
	//
	// This parameter is required.
	//
	// example:
	//
	// TUPLE
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// The number of shards.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// The topic name.
	//
	// This parameter is required.
	//
	// example:
	//
	// adsb_cat021
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s CreateTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTopicRequest) GoString() string {
	return s.String()
}

func (s *CreateTopicRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateTopicRequest) GetEnableSchemaRegistry() *bool {
	return s.EnableSchemaRegistry
}

func (s *CreateTopicRequest) GetExpandMode() *bool {
	return s.ExpandMode
}

func (s *CreateTopicRequest) GetLifecycle() *int32 {
	return s.Lifecycle
}

func (s *CreateTopicRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateTopicRequest) GetRecordSchema() *string {
	return s.RecordSchema
}

func (s *CreateTopicRequest) GetRecordType() *string {
	return s.RecordType
}

func (s *CreateTopicRequest) GetShardCount() *int32 {
	return s.ShardCount
}

func (s *CreateTopicRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *CreateTopicRequest) SetComment(v string) *CreateTopicRequest {
	s.Comment = &v
	return s
}

func (s *CreateTopicRequest) SetEnableSchemaRegistry(v bool) *CreateTopicRequest {
	s.EnableSchemaRegistry = &v
	return s
}

func (s *CreateTopicRequest) SetExpandMode(v bool) *CreateTopicRequest {
	s.ExpandMode = &v
	return s
}

func (s *CreateTopicRequest) SetLifecycle(v int32) *CreateTopicRequest {
	s.Lifecycle = &v
	return s
}

func (s *CreateTopicRequest) SetProjectName(v string) *CreateTopicRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateTopicRequest) SetRecordSchema(v string) *CreateTopicRequest {
	s.RecordSchema = &v
	return s
}

func (s *CreateTopicRequest) SetRecordType(v string) *CreateTopicRequest {
	s.RecordType = &v
	return s
}

func (s *CreateTopicRequest) SetShardCount(v int32) *CreateTopicRequest {
	s.ShardCount = &v
	return s
}

func (s *CreateTopicRequest) SetTopicName(v string) *CreateTopicRequest {
	s.TopicName = &v
	return s
}

func (s *CreateTopicRequest) Validate() error {
	return dara.Validate(s)
}
