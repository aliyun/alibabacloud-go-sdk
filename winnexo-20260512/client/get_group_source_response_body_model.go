// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGroupSourceResponseBody
	GetCode() *string
	SetDescription(v string) *GetGroupSourceResponseBody
	GetDescription() *string
	SetGmtCreate(v string) *GetGroupSourceResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *GetGroupSourceResponseBody
	GetGmtModified() *string
	SetGroupId(v string) *GetGroupSourceResponseBody
	GetGroupId() *string
	SetMessage(v string) *GetGroupSourceResponseBody
	GetMessage() *string
	SetName(v string) *GetGroupSourceResponseBody
	GetName() *string
	SetRequestId(v string) *GetGroupSourceResponseBody
	GetRequestId() *string
	SetScope(v string) *GetGroupSourceResponseBody
	GetScope() *string
	SetSourceId(v string) *GetGroupSourceResponseBody
	GetSourceId() *string
	SetSourceKind(v string) *GetGroupSourceResponseBody
	GetSourceKind() *string
	SetSourceTags(v string) *GetGroupSourceResponseBody
	GetSourceTags() *string
	SetSourceType(v string) *GetGroupSourceResponseBody
	GetSourceType() *string
	SetStatus(v string) *GetGroupSourceResponseBody
	GetStatus() *string
}

type GetGroupSourceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The pipeline description.
	//
	// example:
	//
	// recorder function
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2026-08-26T10:00:00+08:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The time when the resource was last modified, in ISO 8601 format.
	//
	// example:
	//
	// 2026-08-20T14:00:00+08:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The project group ID.
	//
	// example:
	//
	// exampleGroupId
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The permission scope.
	//
	// example:
	//
	// GROUP
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The knowledge base ownership type. Valid values:
	//
	// - aliding_kb_doc: DingTalk knowledge base document.
	//
	// - normal: Common knowledge.
	//
	// example:
	//
	// string_value
	SourceKind *string `json:"sourceKind,omitempty" xml:"sourceKind,omitempty"`
	// The resource tags. This parameter is optional. The value is a JSON string list, such as ["tagA","tagB"].
	//
	// example:
	//
	// ["Important","Document"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The type of the resource source. Valid values:
	//
	// - ExportTaskId: The resource export ID.
	//
	// - TaskId: The module execution task ID.
	//
	// - StatePath: The OSS path where the resource state is stored.
	//
	// example:
	//
	// string_value
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The resource status. The initial status during the creation process is typically PENDING. If the on_create operation fails, the status is FAILED.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetGroupSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGroupSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGroupSourceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetGroupSourceResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetGroupSourceResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetGroupSourceResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *GetGroupSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGroupSourceResponseBody) GetName() *string {
	return s.Name
}

func (s *GetGroupSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGroupSourceResponseBody) GetScope() *string {
	return s.Scope
}

func (s *GetGroupSourceResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *GetGroupSourceResponseBody) GetSourceKind() *string {
	return s.SourceKind
}

func (s *GetGroupSourceResponseBody) GetSourceTags() *string {
	return s.SourceTags
}

func (s *GetGroupSourceResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *GetGroupSourceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetGroupSourceResponseBody) SetCode(v string) *GetGroupSourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetDescription(v string) *GetGroupSourceResponseBody {
	s.Description = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetGmtCreate(v string) *GetGroupSourceResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetGmtModified(v string) *GetGroupSourceResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetGroupId(v string) *GetGroupSourceResponseBody {
	s.GroupId = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetMessage(v string) *GetGroupSourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetName(v string) *GetGroupSourceResponseBody {
	s.Name = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetRequestId(v string) *GetGroupSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetScope(v string) *GetGroupSourceResponseBody {
	s.Scope = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetSourceId(v string) *GetGroupSourceResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetSourceKind(v string) *GetGroupSourceResponseBody {
	s.SourceKind = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetSourceTags(v string) *GetGroupSourceResponseBody {
	s.SourceTags = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetSourceType(v string) *GetGroupSourceResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetGroupSourceResponseBody) SetStatus(v string) *GetGroupSourceResponseBody {
	s.Status = &v
	return s
}

func (s *GetGroupSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
