// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetKnowledgeBaseSourceResponseBody
	GetCode() *string
	SetDescription(v string) *GetKnowledgeBaseSourceResponseBody
	GetDescription() *string
	SetDirectoryId(v string) *GetKnowledgeBaseSourceResponseBody
	GetDirectoryId() *string
	SetDirectoryPath(v string) *GetKnowledgeBaseSourceResponseBody
	GetDirectoryPath() *string
	SetGmtCreate(v int64) *GetKnowledgeBaseSourceResponseBody
	GetGmtCreate() *int64
	SetGmtModified(v int64) *GetKnowledgeBaseSourceResponseBody
	GetGmtModified() *int64
	SetMessage(v string) *GetKnowledgeBaseSourceResponseBody
	GetMessage() *string
	SetName(v string) *GetKnowledgeBaseSourceResponseBody
	GetName() *string
	SetRequestId(v string) *GetKnowledgeBaseSourceResponseBody
	GetRequestId() *string
	SetSourceId(v string) *GetKnowledgeBaseSourceResponseBody
	GetSourceId() *string
	SetSourceKind(v string) *GetKnowledgeBaseSourceResponseBody
	GetSourceKind() *string
	SetSourceTags(v string) *GetKnowledgeBaseSourceResponseBody
	GetSourceTags() *string
	SetSourceType(v string) *GetKnowledgeBaseSourceResponseBody
	GetSourceType() *string
	SetStatus(v string) *GetKnowledgeBaseSourceResponseBody
	GetStatus() *string
	SetStatusMessage(v string) *GetKnowledgeBaseSourceResponseBody
	GetStatusMessage() *string
}

type GetKnowledgeBaseSourceResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the to-do card type.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The full path of the category to which the knowledge item belongs.
	//
	// example:
	//
	// string_value
	DirectoryPath *string `json:"directoryPath,omitempty" xml:"directoryPath,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 1
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The unique identifier on the business system side, that is, the business ID.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The knowledge base ownership type. Valid values: aliding_kb_doc (DingTalk knowledge base document) and normal (common knowledge).
	//
	// example:
	//
	// string_value
	SourceKind *string `json:"sourceKind,omitempty" xml:"sourceKind,omitempty"`
	// The resource tags. This is optional and is a JSON string list, such as ["tagA","tagB"].
	//
	// example:
	//
	// string_value
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The type of the resource source. Valid values:
	//
	// - ExportTaskId: the resource export ID.
	//
	// - TaskId: the Module execution task ID.
	//
	// - StatePath: the OSS path where the resource state is stored.
	//
	// example:
	//
	// string_value
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The data source status. Valid values:
	//
	// - **1**: online.
	//
	// - **0**: offline.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The status message.
	//
	// example:
	//
	// string_value
	StatusMessage *string `json:"statusMessage,omitempty" xml:"statusMessage,omitempty"`
}

func (s GetKnowledgeBaseSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetKnowledgeBaseSourceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetKnowledgeBaseSourceResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetKnowledgeBaseSourceResponseBody) GetDirectoryPath() *string {
	return s.DirectoryPath
}

func (s *GetKnowledgeBaseSourceResponseBody) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetKnowledgeBaseSourceResponseBody) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetKnowledgeBaseSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKnowledgeBaseSourceResponseBody) GetName() *string {
	return s.Name
}

func (s *GetKnowledgeBaseSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKnowledgeBaseSourceResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *GetKnowledgeBaseSourceResponseBody) GetSourceKind() *string {
	return s.SourceKind
}

func (s *GetKnowledgeBaseSourceResponseBody) GetSourceTags() *string {
	return s.SourceTags
}

func (s *GetKnowledgeBaseSourceResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *GetKnowledgeBaseSourceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetKnowledgeBaseSourceResponseBody) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *GetKnowledgeBaseSourceResponseBody) SetCode(v string) *GetKnowledgeBaseSourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetKnowledgeBaseSourceResponseBody) SetDescription(v string) *GetKnowledgeBaseSourceResponseBody {
	s.Description = &v
	return s
}

func (s *GetKnowledgeBaseSourceResponseBody) SetDirectoryId(v string) *GetKnowledgeBaseSourceResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *GetKnowledgeBaseSourceResponseBody) SetDirectoryPath(v string) *GetKnowledgeBaseSourceResponseBody {
	s.DirectoryPath = &v
	return s
}

func (s *GetKnowledgeBaseSourceResponseBody) SetGmtCreate(v int64) *GetKnowledgeBaseSourceResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetKnowledgeBaseSourceResponseBody) SetGmtModified(v int64) *GetKnowledgeBaseSourceResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetKnowledgeBaseSourceResponseBody) SetMessage(v string) *GetKnowledgeBaseSourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetKnowledgeBaseSourceResponseBody) SetName(v string) *GetKnowledgeBaseSourceResponseBody {
	s.Name = &v
	return s
}

func (s *GetKnowledgeBaseSourceResponseBody) SetRequestId(v string) *GetKnowledgeBaseSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKnowledgeBaseSourceResponseBody) SetSourceId(v string) *GetKnowledgeBaseSourceResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetKnowledgeBaseSourceResponseBody) SetSourceKind(v string) *GetKnowledgeBaseSourceResponseBody {
	s.SourceKind = &v
	return s
}

func (s *GetKnowledgeBaseSourceResponseBody) SetSourceTags(v string) *GetKnowledgeBaseSourceResponseBody {
	s.SourceTags = &v
	return s
}

func (s *GetKnowledgeBaseSourceResponseBody) SetSourceType(v string) *GetKnowledgeBaseSourceResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetKnowledgeBaseSourceResponseBody) SetStatus(v string) *GetKnowledgeBaseSourceResponseBody {
	s.Status = &v
	return s
}

func (s *GetKnowledgeBaseSourceResponseBody) SetStatusMessage(v string) *GetKnowledgeBaseSourceResponseBody {
	s.StatusMessage = &v
	return s
}

func (s *GetKnowledgeBaseSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
