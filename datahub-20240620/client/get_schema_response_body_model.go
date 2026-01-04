// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *GetSchemaResponseBody
	GetCreateTime() *int64
	SetCreator(v string) *GetSchemaResponseBody
	GetCreator() *string
	SetProjectName(v string) *GetSchemaResponseBody
	GetProjectName() *string
	SetRecordSchema(v string) *GetSchemaResponseBody
	GetRecordSchema() *string
	SetRequestId(v string) *GetSchemaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSchemaResponseBody
	GetSuccess() *bool
	SetTopicName(v string) *GetSchemaResponseBody
	GetTopicName() *string
	SetVersionId(v int32) *GetSchemaResponseBody
	GetVersionId() *int32
}

type GetSchemaResponseBody struct {
	// example:
	//
	// 1724041098000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1559031978056215
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// [{\\"Type\\":\\"STRING\\",\\"AllowNull\\":true,\\"Name\\":\\"context\\"}]
	RecordSchema *string `json:"RecordSchema,omitempty" xml:"RecordSchema,omitempty"`
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
	// test_topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// example:
	//
	// 0
	VersionId *int32 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetSchemaResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetSchemaResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *GetSchemaResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetSchemaResponseBody) GetRecordSchema() *string {
	return s.RecordSchema
}

func (s *GetSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSchemaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSchemaResponseBody) GetTopicName() *string {
	return s.TopicName
}

func (s *GetSchemaResponseBody) GetVersionId() *int32 {
	return s.VersionId
}

func (s *GetSchemaResponseBody) SetCreateTime(v int64) *GetSchemaResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetSchemaResponseBody) SetCreator(v string) *GetSchemaResponseBody {
	s.Creator = &v
	return s
}

func (s *GetSchemaResponseBody) SetProjectName(v string) *GetSchemaResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetSchemaResponseBody) SetRecordSchema(v string) *GetSchemaResponseBody {
	s.RecordSchema = &v
	return s
}

func (s *GetSchemaResponseBody) SetRequestId(v string) *GetSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSchemaResponseBody) SetSuccess(v bool) *GetSchemaResponseBody {
	s.Success = &v
	return s
}

func (s *GetSchemaResponseBody) SetTopicName(v string) *GetSchemaResponseBody {
	s.TopicName = &v
	return s
}

func (s *GetSchemaResponseBody) SetVersionId(v int32) *GetSchemaResponseBody {
	s.VersionId = &v
	return s
}

func (s *GetSchemaResponseBody) Validate() error {
	return dara.Validate(s)
}
