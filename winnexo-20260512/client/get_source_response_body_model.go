// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSourceResponseBody
	GetCode() *string
	SetCompletionTime(v string) *GetSourceResponseBody
	GetCompletionTime() *string
	SetDescription(v string) *GetSourceResponseBody
	GetDescription() *string
	SetGmtCreate(v string) *GetSourceResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *GetSourceResponseBody
	GetGmtModified() *string
	SetHasNotes(v bool) *GetSourceResponseBody
	GetHasNotes() *bool
	SetHasSettings(v bool) *GetSourceResponseBody
	GetHasSettings() *bool
	SetHasStructuredTables(v bool) *GetSourceResponseBody
	GetHasStructuredTables() *bool
	SetHasUnstructuredDocs(v bool) *GetSourceResponseBody
	GetHasUnstructuredDocs() *bool
	SetMessage(v string) *GetSourceResponseBody
	GetMessage() *string
	SetName(v string) *GetSourceResponseBody
	GetName() *string
	SetNotes(v string) *GetSourceResponseBody
	GetNotes() *string
	SetObjectBindings(v []*GetSourceResponseBodyObjectBindings) *GetSourceResponseBody
	GetObjectBindings() []*GetSourceResponseBodyObjectBindings
	SetObjectId(v string) *GetSourceResponseBody
	GetObjectId() *string
	SetObjectType(v string) *GetSourceResponseBody
	GetObjectType() *string
	SetOperatingObjectName(v string) *GetSourceResponseBody
	GetOperatingObjectName() *string
	SetRequestId(v string) *GetSourceResponseBody
	GetRequestId() *string
	SetScope(v string) *GetSourceResponseBody
	GetScope() *string
	SetSettings(v map[string]interface{}) *GetSourceResponseBody
	GetSettings() map[string]interface{}
	SetSkillOutputId(v string) *GetSourceResponseBody
	GetSkillOutputId() *string
	SetSourceId(v string) *GetSourceResponseBody
	GetSourceId() *string
	SetSourceKind(v string) *GetSourceResponseBody
	GetSourceKind() *string
	SetSourceTags(v string) *GetSourceResponseBody
	GetSourceTags() *string
	SetSourceType(v string) *GetSourceResponseBody
	GetSourceType() *string
	SetStatus(v string) *GetSourceResponseBody
	GetStatus() *string
	SetStructuredTables(v []*string) *GetSourceResponseBody
	GetStructuredTables() []*string
	SetUnstructuredDocs(v []*GetSourceResponseBodyUnstructuredDocs) *GetSourceResponseBody
	GetUnstructuredDocs() []*GetSourceResponseBodyUnstructuredDocs
}

type GetSourceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The completion time, in milliseconds.
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	CompletionTime *string `json:"completionTime,omitempty" xml:"completionTime,omitempty"`
	// The description of the to-do card type.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The creation time.
	//
	// example:
	//
	// string_value
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// string_value
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Indicates whether notes exist.
	//
	// example:
	//
	// true
	HasNotes *bool `json:"hasNotes,omitempty" xml:"hasNotes,omitempty"`
	// Indicates whether settings configuration exists.
	//
	// example:
	//
	// true
	HasSettings *bool `json:"hasSettings,omitempty" xml:"hasSettings,omitempty"`
	// Indicates whether structured tables exist.
	//
	// example:
	//
	// true
	HasStructuredTables *bool `json:"hasStructuredTables,omitempty" xml:"hasStructuredTables,omitempty"`
	// Indicates whether unstructured documents exist.
	//
	// example:
	//
	// true
	HasUnstructuredDocs *bool `json:"hasUnstructuredDocs,omitempty" xml:"hasUnstructuredDocs,omitempty"`
	// The status code description.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The username.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The meeting notes content (optional). The notes are used for auxiliary analysis.
	//
	// example:
	//
	// string_value
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// The object bindings.
	ObjectBindings []*GetSourceResponseBodyObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
	// The associated variable ID.
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The object type, such as customer. This field has a value when type is set to mention.
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The digital employee name (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The task scope.
	//
	// example:
	//
	// PERSONAL
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The user profile settings.
	Settings map[string]interface{} `json:"settings,omitempty" xml:"settings,omitempty"`
	// The skill output ID (carried when the output is saved as a resource).
	//
	// example:
	//
	// exampleSkillOutputId
	SkillOutputId *string `json:"skillOutputId,omitempty" xml:"skillOutputId,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The knowledge base (KB) ownership type. Valid values:
	//
	// - aliding_kb_doc: DingTalk knowledge base document.
	//
	// - normal: common knowledge.
	//
	// example:
	//
	// string_value
	SourceKind *string `json:"sourceKind,omitempty" xml:"sourceKind,omitempty"`
	// The resource tags (optional, a JSON string list such as ["tagA","tagB"]).
	//
	// example:
	//
	// string_value
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The resource type.
	//
	// example:
	//
	// string_value
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The final status of the message.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of structured tables.
	//
	// example:
	//
	// string_value
	StructuredTables []*string `json:"structuredTables,omitempty" xml:"structuredTables,omitempty" type:"Repeated"`
	// The unstructured documents.
	UnstructuredDocs []*GetSourceResponseBodyUnstructuredDocs `json:"unstructuredDocs,omitempty" xml:"unstructuredDocs,omitempty" type:"Repeated"`
}

func (s GetSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSourceResponseBody) GetCompletionTime() *string {
	return s.CompletionTime
}

func (s *GetSourceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetSourceResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetSourceResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetSourceResponseBody) GetHasNotes() *bool {
	return s.HasNotes
}

func (s *GetSourceResponseBody) GetHasSettings() *bool {
	return s.HasSettings
}

func (s *GetSourceResponseBody) GetHasStructuredTables() *bool {
	return s.HasStructuredTables
}

func (s *GetSourceResponseBody) GetHasUnstructuredDocs() *bool {
	return s.HasUnstructuredDocs
}

func (s *GetSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSourceResponseBody) GetName() *string {
	return s.Name
}

func (s *GetSourceResponseBody) GetNotes() *string {
	return s.Notes
}

func (s *GetSourceResponseBody) GetObjectBindings() []*GetSourceResponseBodyObjectBindings {
	return s.ObjectBindings
}

func (s *GetSourceResponseBody) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetSourceResponseBody) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetSourceResponseBody) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *GetSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSourceResponseBody) GetScope() *string {
	return s.Scope
}

func (s *GetSourceResponseBody) GetSettings() map[string]interface{} {
	return s.Settings
}

func (s *GetSourceResponseBody) GetSkillOutputId() *string {
	return s.SkillOutputId
}

func (s *GetSourceResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *GetSourceResponseBody) GetSourceKind() *string {
	return s.SourceKind
}

func (s *GetSourceResponseBody) GetSourceTags() *string {
	return s.SourceTags
}

func (s *GetSourceResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *GetSourceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSourceResponseBody) GetStructuredTables() []*string {
	return s.StructuredTables
}

func (s *GetSourceResponseBody) GetUnstructuredDocs() []*GetSourceResponseBodyUnstructuredDocs {
	return s.UnstructuredDocs
}

func (s *GetSourceResponseBody) SetCode(v string) *GetSourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetSourceResponseBody) SetCompletionTime(v string) *GetSourceResponseBody {
	s.CompletionTime = &v
	return s
}

func (s *GetSourceResponseBody) SetDescription(v string) *GetSourceResponseBody {
	s.Description = &v
	return s
}

func (s *GetSourceResponseBody) SetGmtCreate(v string) *GetSourceResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetSourceResponseBody) SetGmtModified(v string) *GetSourceResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetSourceResponseBody) SetHasNotes(v bool) *GetSourceResponseBody {
	s.HasNotes = &v
	return s
}

func (s *GetSourceResponseBody) SetHasSettings(v bool) *GetSourceResponseBody {
	s.HasSettings = &v
	return s
}

func (s *GetSourceResponseBody) SetHasStructuredTables(v bool) *GetSourceResponseBody {
	s.HasStructuredTables = &v
	return s
}

func (s *GetSourceResponseBody) SetHasUnstructuredDocs(v bool) *GetSourceResponseBody {
	s.HasUnstructuredDocs = &v
	return s
}

func (s *GetSourceResponseBody) SetMessage(v string) *GetSourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetSourceResponseBody) SetName(v string) *GetSourceResponseBody {
	s.Name = &v
	return s
}

func (s *GetSourceResponseBody) SetNotes(v string) *GetSourceResponseBody {
	s.Notes = &v
	return s
}

func (s *GetSourceResponseBody) SetObjectBindings(v []*GetSourceResponseBodyObjectBindings) *GetSourceResponseBody {
	s.ObjectBindings = v
	return s
}

func (s *GetSourceResponseBody) SetObjectId(v string) *GetSourceResponseBody {
	s.ObjectId = &v
	return s
}

func (s *GetSourceResponseBody) SetObjectType(v string) *GetSourceResponseBody {
	s.ObjectType = &v
	return s
}

func (s *GetSourceResponseBody) SetOperatingObjectName(v string) *GetSourceResponseBody {
	s.OperatingObjectName = &v
	return s
}

func (s *GetSourceResponseBody) SetRequestId(v string) *GetSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSourceResponseBody) SetScope(v string) *GetSourceResponseBody {
	s.Scope = &v
	return s
}

func (s *GetSourceResponseBody) SetSettings(v map[string]interface{}) *GetSourceResponseBody {
	s.Settings = v
	return s
}

func (s *GetSourceResponseBody) SetSkillOutputId(v string) *GetSourceResponseBody {
	s.SkillOutputId = &v
	return s
}

func (s *GetSourceResponseBody) SetSourceId(v string) *GetSourceResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetSourceResponseBody) SetSourceKind(v string) *GetSourceResponseBody {
	s.SourceKind = &v
	return s
}

func (s *GetSourceResponseBody) SetSourceTags(v string) *GetSourceResponseBody {
	s.SourceTags = &v
	return s
}

func (s *GetSourceResponseBody) SetSourceType(v string) *GetSourceResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetSourceResponseBody) SetStatus(v string) *GetSourceResponseBody {
	s.Status = &v
	return s
}

func (s *GetSourceResponseBody) SetStructuredTables(v []*string) *GetSourceResponseBody {
	s.StructuredTables = v
	return s
}

func (s *GetSourceResponseBody) SetUnstructuredDocs(v []*GetSourceResponseBodyUnstructuredDocs) *GetSourceResponseBody {
	s.UnstructuredDocs = v
	return s
}

func (s *GetSourceResponseBody) Validate() error {
	if s.ObjectBindings != nil {
		for _, item := range s.ObjectBindings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UnstructuredDocs != nil {
		for _, item := range s.UnstructuredDocs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSourceResponseBodyObjectBindings struct {
	// The semantic graph name to which the object belongs.
	//
	// example:
	//
	// string_value
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The ID of the recommended item, which can be a **feedId*	- or a micro-application ID.
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The object type, such as customer. This field has a value when type is set to mention.
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s GetSourceResponseBodyObjectBindings) String() string {
	return dara.Prettify(s)
}

func (s GetSourceResponseBodyObjectBindings) GoString() string {
	return s.String()
}

func (s *GetSourceResponseBodyObjectBindings) GetGraphName() *string {
	return s.GraphName
}

func (s *GetSourceResponseBodyObjectBindings) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetSourceResponseBodyObjectBindings) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetSourceResponseBodyObjectBindings) SetGraphName(v string) *GetSourceResponseBodyObjectBindings {
	s.GraphName = &v
	return s
}

func (s *GetSourceResponseBodyObjectBindings) SetObjectId(v string) *GetSourceResponseBodyObjectBindings {
	s.ObjectId = &v
	return s
}

func (s *GetSourceResponseBodyObjectBindings) SetObjectType(v string) *GetSourceResponseBodyObjectBindings {
	s.ObjectType = &v
	return s
}

func (s *GetSourceResponseBodyObjectBindings) Validate() error {
	return dara.Validate(s)
}

type GetSourceResponseBodyUnstructuredDocs struct {
	// The completion time, in milliseconds.
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	CompletionTime *string `json:"completionTime,omitempty" xml:"completionTime,omitempty"`
	// The file name, including the file name extension.
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The file record ID (optional, corresponding to settings.file_record_id).
	//
	// example:
	//
	// exampleFileRecordId
	FileRecordId *string `json:"fileRecordId,omitempty" xml:"fileRecordId,omitempty"`
	// The file type. Valid values:
	//
	// - **file**: file.
	//
	// - **folder**: folder.
	//
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// The session analysis result in OSS URL format. The URL expires in one hour.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	OssUrl *string `json:"ossUrl,omitempty" xml:"ossUrl,omitempty"`
	// The DocumentAgent semantic ID.
	//
	// example:
	//
	// exampleSemanticsId
	SemanticsId *string `json:"semanticsId,omitempty" xml:"semanticsId,omitempty"`
}

func (s GetSourceResponseBodyUnstructuredDocs) String() string {
	return dara.Prettify(s)
}

func (s GetSourceResponseBodyUnstructuredDocs) GoString() string {
	return s.String()
}

func (s *GetSourceResponseBodyUnstructuredDocs) GetCompletionTime() *string {
	return s.CompletionTime
}

func (s *GetSourceResponseBodyUnstructuredDocs) GetFileName() *string {
	return s.FileName
}

func (s *GetSourceResponseBodyUnstructuredDocs) GetFileRecordId() *string {
	return s.FileRecordId
}

func (s *GetSourceResponseBodyUnstructuredDocs) GetFileType() *string {
	return s.FileType
}

func (s *GetSourceResponseBodyUnstructuredDocs) GetOssUrl() *string {
	return s.OssUrl
}

func (s *GetSourceResponseBodyUnstructuredDocs) GetSemanticsId() *string {
	return s.SemanticsId
}

func (s *GetSourceResponseBodyUnstructuredDocs) SetCompletionTime(v string) *GetSourceResponseBodyUnstructuredDocs {
	s.CompletionTime = &v
	return s
}

func (s *GetSourceResponseBodyUnstructuredDocs) SetFileName(v string) *GetSourceResponseBodyUnstructuredDocs {
	s.FileName = &v
	return s
}

func (s *GetSourceResponseBodyUnstructuredDocs) SetFileRecordId(v string) *GetSourceResponseBodyUnstructuredDocs {
	s.FileRecordId = &v
	return s
}

func (s *GetSourceResponseBodyUnstructuredDocs) SetFileType(v string) *GetSourceResponseBodyUnstructuredDocs {
	s.FileType = &v
	return s
}

func (s *GetSourceResponseBodyUnstructuredDocs) SetOssUrl(v string) *GetSourceResponseBodyUnstructuredDocs {
	s.OssUrl = &v
	return s
}

func (s *GetSourceResponseBodyUnstructuredDocs) SetSemanticsId(v string) *GetSourceResponseBodyUnstructuredDocs {
	s.SemanticsId = &v
	return s
}

func (s *GetSourceResponseBodyUnstructuredDocs) Validate() error {
	return dara.Validate(s)
}
