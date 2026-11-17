// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeObjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListKnowledgeObjectsResponseBody
	GetCode() *string
	SetData(v *ListKnowledgeObjectsResponseBodyData) *ListKnowledgeObjectsResponseBody
	GetData() *ListKnowledgeObjectsResponseBodyData
	SetMaxResults(v int32) *ListKnowledgeObjectsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListKnowledgeObjectsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListKnowledgeObjectsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListKnowledgeObjectsResponseBody
	GetSuccess() *bool
}

type ListKnowledgeObjectsResponseBody struct {
	// The business error code, which is used to determine the error type when a request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *ListKnowledgeObjectsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The secure error message returned when a request fails.
	//
	// example:
	//
	// The request parameters are invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListKnowledgeObjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgeObjectsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListKnowledgeObjectsResponseBody) GetData() *ListKnowledgeObjectsResponseBodyData {
	return s.Data
}

func (s *ListKnowledgeObjectsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeObjectsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListKnowledgeObjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeObjectsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListKnowledgeObjectsResponseBody) SetCode(v string) *ListKnowledgeObjectsResponseBody {
	s.Code = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBody) SetData(v *ListKnowledgeObjectsResponseBodyData) *ListKnowledgeObjectsResponseBody {
	s.Data = v
	return s
}

func (s *ListKnowledgeObjectsResponseBody) SetMaxResults(v int32) *ListKnowledgeObjectsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBody) SetMessage(v string) *ListKnowledgeObjectsResponseBody {
	s.Message = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBody) SetRequestId(v string) *ListKnowledgeObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBody) SetSuccess(v bool) *ListKnowledgeObjectsResponseBody {
	s.Success = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKnowledgeObjectsResponseBodyData struct {
	// The cursor for the next page. The value is null if there is no next page.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The knowledge objects that the caller has permission to read on the current page.
	Objects []*ListKnowledgeObjectsResponseBodyDataObjects `json:"Objects,omitempty" xml:"Objects,omitempty" type:"Repeated"`
}

func (s ListKnowledgeObjectsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeObjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKnowledgeObjectsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeObjectsResponseBodyData) GetObjects() []*ListKnowledgeObjectsResponseBodyDataObjects {
	return s.Objects
}

func (s *ListKnowledgeObjectsResponseBodyData) SetNextToken(v string) *ListKnowledgeObjectsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyData) SetObjects(v []*ListKnowledgeObjectsResponseBodyDataObjects) *ListKnowledgeObjectsResponseBodyData {
	s.Objects = v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyData) Validate() error {
	if s.Objects != nil {
		for _, item := range s.Objects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgeObjectsResponseBodyDataObjects struct {
	// The typed content of the knowledge object. The semantic model content or document content is returned based on the Type.
	Content *ListKnowledgeObjectsResponseBodyDataObjectsContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The time when the object was first created.
	//
	// example:
	//
	// 1790200800000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The brief description of the object. This field is editable. Pass an empty string to clear the description.
	//
	// example:
	//
	// Sales metrics and policy documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The download items for multi-file specification YAMLs, including file names and temporary download URLs.
	DownloadFiles []*ListKnowledgeObjectsResponseBodyDataObjectsDownloadFiles `json:"DownloadFiles,omitempty" xml:"DownloadFiles,omitempty" type:"Repeated"`
	// The temporary download URL of the specification YAML file. The validity period is subject to ExpiresIn.
	//
	// example:
	//
	// https://example.invalid/short-lived-result-download
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// Indicates whether the current caller has full content editing permissions. The value is true for OWNER and ADMIN and false for VIEWER. Authorization is checked again for write operations.
	Editable *bool `json:"Editable,omitempty" xml:"Editable,omitempty"`
	// The validity period of the knowledge object download URL in seconds. The current value is 300 seconds.
	//
	// example:
	//
	// 300
	ExpiresIn *int32 `json:"ExpiresIn,omitempty" xml:"ExpiresIn,omitempty"`
	// The original file name of the uploaded document. This field is read-only. Renaming the Name field does not change this value. This field is not returned for external documents and semantic models.
	//
	// example:
	//
	// policy.md
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The knowledge base to which the object belongs.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The ID of the knowledge object. For semantic models, use the model ID from the semantic layer. For documents, use the stable ID from the local document table. Callers must treat the ID as opaque and must not use a retrieval ID in its place.
	//
	// example:
	//
	// obj-001
	KnowledgeObjectId *string `json:"KnowledgeObjectId,omitempty" xml:"KnowledgeObjectId,omitempty"`
	// The time when the current content was last updated.
	//
	// example:
	//
	// 1790200800000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The user-facing name of the object. This field is editable. Renaming the object does not change the KnowledgeObjectId.
	//
	// example:
	//
	// Sales policy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The size of the uploaded original file in bytes. This field is read-only and does not represent the length of the edited body text.
	//
	// example:
	//
	// 1024
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
	// The array of source references. Currently, an empty array is returned.
	SourceReferences []*ListKnowledgeObjectsResponseBodyDataObjectsSourceReferences `json:"SourceReferences,omitempty" xml:"SourceReferences,omitempty" type:"Repeated"`
	// The knowledge object type. The current values are SEMANTIC_MODEL and DOCUMENT. More types may be added in the future.
	//
	// example:
	//
	// DOCUMENT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The current version of the SEMANTIC_MODEL being read. Pass this value when updating content. The version can be edited in place.
	//
	// example:
	//
	// 1
	VersionNo *int64 `json:"VersionNo,omitempty" xml:"VersionNo,omitempty"`
}

func (s ListKnowledgeObjectsResponseBodyDataObjects) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeObjectsResponseBodyDataObjects) GoString() string {
	return s.String()
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) GetContent() *ListKnowledgeObjectsResponseBodyDataObjectsContent {
	return s.Content
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) GetDescription() *string {
	return s.Description
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) GetDownloadFiles() []*ListKnowledgeObjectsResponseBodyDataObjectsDownloadFiles {
	return s.DownloadFiles
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) GetEditable() *bool {
	return s.Editable
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) GetExpiresIn() *int32 {
	return s.ExpiresIn
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) GetFileName() *string {
	return s.FileName
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) GetKnowledgeObjectId() *string {
	return s.KnowledgeObjectId
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) GetName() *string {
	return s.Name
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) GetSourceReferences() []*ListKnowledgeObjectsResponseBodyDataObjectsSourceReferences {
	return s.SourceReferences
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) GetType() *string {
	return s.Type
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) GetVersionNo() *int64 {
	return s.VersionNo
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) SetContent(v *ListKnowledgeObjectsResponseBodyDataObjectsContent) *ListKnowledgeObjectsResponseBodyDataObjects {
	s.Content = v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) SetCreateTime(v int64) *ListKnowledgeObjectsResponseBodyDataObjects {
	s.CreateTime = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) SetDescription(v string) *ListKnowledgeObjectsResponseBodyDataObjects {
	s.Description = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) SetDownloadFiles(v []*ListKnowledgeObjectsResponseBodyDataObjectsDownloadFiles) *ListKnowledgeObjectsResponseBodyDataObjects {
	s.DownloadFiles = v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) SetDownloadUrl(v string) *ListKnowledgeObjectsResponseBodyDataObjects {
	s.DownloadUrl = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) SetEditable(v bool) *ListKnowledgeObjectsResponseBodyDataObjects {
	s.Editable = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) SetExpiresIn(v int32) *ListKnowledgeObjectsResponseBodyDataObjects {
	s.ExpiresIn = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) SetFileName(v string) *ListKnowledgeObjectsResponseBodyDataObjects {
	s.FileName = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) SetKnowledgeBaseId(v string) *ListKnowledgeObjectsResponseBodyDataObjects {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) SetKnowledgeObjectId(v string) *ListKnowledgeObjectsResponseBodyDataObjects {
	s.KnowledgeObjectId = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) SetModifyTime(v int64) *ListKnowledgeObjectsResponseBodyDataObjects {
	s.ModifyTime = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) SetName(v string) *ListKnowledgeObjectsResponseBodyDataObjects {
	s.Name = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) SetSizeBytes(v int64) *ListKnowledgeObjectsResponseBodyDataObjects {
	s.SizeBytes = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) SetSourceReferences(v []*ListKnowledgeObjectsResponseBodyDataObjectsSourceReferences) *ListKnowledgeObjectsResponseBodyDataObjects {
	s.SourceReferences = v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) SetType(v string) *ListKnowledgeObjectsResponseBodyDataObjects {
	s.Type = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) SetVersionNo(v int64) *ListKnowledgeObjectsResponseBodyDataObjects {
	s.VersionNo = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjects) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	if s.DownloadFiles != nil {
		for _, item := range s.DownloadFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SourceReferences != nil {
		for _, item := range s.SourceReferences {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgeObjectsResponseBodyDataObjectsContent struct {
	// The document content. This field exists only when Type is set to DOCUMENT.
	Document *ListKnowledgeObjectsResponseBodyDataObjectsContentDocument `json:"Document,omitempty" xml:"Document,omitempty" type:"Struct"`
	// The semantic model content. This field exists only when Type is set to SEMANTIC_MODEL.
	SemanticModel *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModel `json:"SemanticModel,omitempty" xml:"SemanticModel,omitempty" type:"Struct"`
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContent) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContent) GoString() string {
	return s.String()
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContent) GetDocument() *ListKnowledgeObjectsResponseBodyDataObjectsContentDocument {
	return s.Document
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContent) GetSemanticModel() *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModel {
	return s.SemanticModel
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContent) SetDocument(v *ListKnowledgeObjectsResponseBodyDataObjectsContentDocument) *ListKnowledgeObjectsResponseBodyDataObjectsContent {
	s.Document = v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContent) SetSemanticModel(v *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModel) *ListKnowledgeObjectsResponseBodyDataObjectsContent {
	s.SemanticModel = v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContent) Validate() error {
	if s.Document != nil {
		if err := s.Document.Validate(); err != nil {
			return err
		}
	}
	if s.SemanticModel != nil {
		if err := s.SemanticModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKnowledgeObjectsResponseBodyDataObjectsContentDocument struct {
	// The content format. This field is required when Text is provided. Initially, the supported values are MARKDOWN and TEXT.
	//
	// example:
	//
	// MARKDOWN
	ContentFormat *string `json:"ContentFormat,omitempty" xml:"ContentFormat,omitempty"`
	// The current document body that can be provided to the caller. This field is optional. It is omitted if parsing is incomplete, access is not authorized, or the full document cannot be provided securely. A single retrieved snippet does not stand in for the entire document.
	//
	// example:
	//
	// Updated full body text
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentDocument) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentDocument) GoString() string {
	return s.String()
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentDocument) GetContentFormat() *string {
	return s.ContentFormat
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentDocument) GetText() *string {
	return s.Text
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentDocument) SetContentFormat(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentDocument {
	s.ContentFormat = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentDocument) SetText(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentDocument {
	s.Text = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentDocument) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModel struct {
	// The logical datasets included in the model. An empty array is returned if no public datasets are available.
	Datasets []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// The executable metrics in the model. An empty array is returned if no metrics are available. Metrics with only candidate formulas are not included.
	Metrics []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The defined relationships in the model expressed by field pairs. An empty array is returned if no relationships are available. Relationships are not inferred from graph paths.
	Relationships []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships `json:"Relationships,omitempty" xml:"Relationships,omitempty" type:"Repeated"`
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModel) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModel) GoString() string {
	return s.String()
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModel) GetDatasets() []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets {
	return s.Datasets
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModel) GetMetrics() []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics {
	return s.Metrics
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModel) GetRelationships() []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships {
	return s.Relationships
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModel) SetDatasets(v []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModel {
	s.Datasets = v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModel) SetMetrics(v []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModel {
	s.Metrics = v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModel) SetRelationships(v []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModel {
	s.Relationships = v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModel) Validate() error {
	if s.Datasets != nil {
		for _, item := range s.Datasets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Relationships != nil {
		for _, item := range s.Relationships {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets struct {
	// The brief description of the object. This field is editable. Pass an empty string to clear the description.
	//
	// example:
	//
	// Sales metrics and policy documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The metadata entity ID returned by Data Map. Clients must pass this value as is and must not concatenate it.
	//
	// example:
	//
	// maxcompute-table:::project_name:schema_name:table_name
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The defined and publicly accessible fields of the dataset. An empty array is returned if no field definitions are provided.
	Fields []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasetsFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The user-facing name of the object. This field is editable. Renaming the object does not change the KnowledgeObjectId.
	//
	// example:
	//
	// Sales policy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The explicitly declared primary key field name. This field is optional. Each item must exist in the Fields of the current dataset.
	PrimaryKey []*string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty" type:"Repeated"`
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets) GoString() string {
	return s.String()
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets) GetDescription() *string {
	return s.Description
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets) GetEntityId() *string {
	return s.EntityId
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets) GetFields() []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasetsFields {
	return s.Fields
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets) GetName() *string {
	return s.Name
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets) GetPrimaryKey() []*string {
	return s.PrimaryKey
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets) SetDescription(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets {
	s.Description = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets) SetEntityId(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets {
	s.EntityId = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets) SetFields(v []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasetsFields) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets {
	s.Fields = v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets) SetName(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets {
	s.Name = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets) SetPrimaryKey(v []*string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets {
	s.PrimaryKey = v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasets) Validate() error {
	if s.Fields != nil {
		for _, item := range s.Fields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasetsFields struct {
	// The data type of the field. This field must be provided for publicly accessible model fields and is not inferred by the server.
	//
	// example:
	//
	// BIGINT
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The brief description of the object. This field is editable. Pass an empty string to clear the description.
	//
	// example:
	//
	// Sales metrics and policy documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user-facing name of the object. This field is editable. Renaming the object does not change the KnowledgeObjectId.
	//
	// example:
	//
	// Sales policy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasetsFields) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasetsFields) GoString() string {
	return s.String()
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasetsFields) GetDataType() *string {
	return s.DataType
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasetsFields) GetDescription() *string {
	return s.Description
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasetsFields) GetName() *string {
	return s.Name
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasetsFields) SetDataType(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasetsFields {
	s.DataType = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasetsFields) SetDescription(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasetsFields {
	s.Description = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasetsFields) SetName(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasetsFields {
	s.Name = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelDatasetsFields) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics struct {
	// The name of the semantic dataset to which the metric belongs. This value must be locatable in the Datasets of the current model.
	//
	// example:
	//
	// orders
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The brief description of the object. This field is editable. Pass an empty string to clear the description.
	//
	// example:
	//
	// Sales metrics and policy documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The executable expression. It must contain at least one dialect version and is not auto-generated by the model.
	Expression *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpression `json:"Expression,omitempty" xml:"Expression,omitempty" type:"Struct"`
	// The user-facing name of the object. This field is editable. Renaming the object does not change the KnowledgeObjectId.
	//
	// example:
	//
	// Sales policy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The explicitly declared unit of the metric. This field is optional and is not returned if no reliable source is available.
	//
	// example:
	//
	// CNY
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics) GoString() string {
	return s.String()
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics) GetDatasetName() *string {
	return s.DatasetName
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics) GetDescription() *string {
	return s.Description
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics) GetExpression() *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpression {
	return s.Expression
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics) GetName() *string {
	return s.Name
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics) GetUnit() *string {
	return s.Unit
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics) SetDatasetName(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics {
	s.DatasetName = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics) SetDescription(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics {
	s.Description = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics) SetExpression(v *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpression) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics {
	s.Expression = v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics) SetName(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics {
	s.Name = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics) SetUnit(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics {
	s.Unit = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetrics) Validate() error {
	if s.Expression != nil {
		if err := s.Expression.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpression struct {
	// The expressions of the same metric in different SQL dialects. At least one item is required. Dialects are not duplicated.
	Dialects []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpressionDialects `json:"Dialects,omitempty" xml:"Dialects,omitempty" type:"Repeated"`
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpression) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpression) GoString() string {
	return s.String()
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpression) GetDialects() []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpressionDialects {
	return s.Dialects
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpression) SetDialects(v []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpressionDialects) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpression {
	s.Dialects = v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpression) Validate() error {
	if s.Dialects != nil {
		for _, item := range s.Dialects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpressionDialects struct {
	// The SQL dialect, such as MaxCompute or ANSI_SQL.
	//
	// example:
	//
	// MaxCompute
	Dialect *string `json:"Dialect,omitempty" xml:"Dialect,omitempty"`
	// The non-empty calculation expression in the specified dialect. This is not a complete query statement.
	//
	// example:
	//
	// SUM(pay_amount)
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpressionDialects) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpressionDialects) GoString() string {
	return s.String()
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpressionDialects) GetDialect() *string {
	return s.Dialect
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpressionDialects) GetSql() *string {
	return s.Sql
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpressionDialects) SetDialect(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpressionDialects {
	s.Dialect = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpressionDialects) SetSql(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpressionDialects {
	s.Sql = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelMetricsExpressionDialects) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships struct {
	// The name of the starting semantic dataset.
	//
	// example:
	//
	// orders
	FromDataset *string `json:"FromDataset,omitempty" xml:"FromDataset,omitempty"`
	// The equi-join pairs of fields at both ends of the relationship. For composite relationships, fields are paired one by one.
	JoinKeys []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationshipsJoinKeys `json:"JoinKeys,omitempty" xml:"JoinKeys,omitempty" type:"Repeated"`
	// The user-facing name of the object. This field is editable. Renaming the object does not change the KnowledgeObjectId.
	//
	// example:
	//
	// Sales policy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the ending semantic dataset.
	//
	// example:
	//
	// customers
	ToDataset *string `json:"ToDataset,omitempty" xml:"ToDataset,omitempty"`
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships) GoString() string {
	return s.String()
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships) GetFromDataset() *string {
	return s.FromDataset
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships) GetJoinKeys() []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationshipsJoinKeys {
	return s.JoinKeys
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships) GetName() *string {
	return s.Name
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships) GetToDataset() *string {
	return s.ToDataset
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships) SetFromDataset(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships {
	s.FromDataset = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships) SetJoinKeys(v []*ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationshipsJoinKeys) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships {
	s.JoinKeys = v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships) SetName(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships {
	s.Name = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships) SetToDataset(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships {
	s.ToDataset = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationships) Validate() error {
	if s.JoinKeys != nil {
		for _, item := range s.JoinKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationshipsJoinKeys struct {
	// The name of the field in the starting dataset that participates in the relationship. This field must be defined in the Fields of the dataset corresponding to FromDataset.
	//
	// example:
	//
	// customer_id
	FromField *string `json:"FromField,omitempty" xml:"FromField,omitempty"`
	// The name of the field in the ending dataset that participates in the relationship. This field must be defined in the Fields of the dataset corresponding to ToDataset.
	//
	// example:
	//
	// id
	ToField *string `json:"ToField,omitempty" xml:"ToField,omitempty"`
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationshipsJoinKeys) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationshipsJoinKeys) GoString() string {
	return s.String()
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationshipsJoinKeys) GetFromField() *string {
	return s.FromField
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationshipsJoinKeys) GetToField() *string {
	return s.ToField
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationshipsJoinKeys) SetFromField(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationshipsJoinKeys {
	s.FromField = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationshipsJoinKeys) SetToField(v string) *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationshipsJoinKeys {
	s.ToField = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsContentSemanticModelRelationshipsJoinKeys) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeObjectsResponseBodyDataObjectsDownloadFiles struct {
	// The temporary download URL of the specification YAML file. The validity period is subject to ExpiresIn.
	//
	// example:
	//
	// https://example.invalid/short-lived-result-download
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The original file name of the uploaded document. This field is read-only. Renaming the Name field does not change this value. This field is not returned for external documents and semantic models.
	//
	// example:
	//
	// policy.md
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsDownloadFiles) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsDownloadFiles) GoString() string {
	return s.String()
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsDownloadFiles) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsDownloadFiles) GetFileName() *string {
	return s.FileName
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsDownloadFiles) SetDownloadUrl(v string) *ListKnowledgeObjectsResponseBodyDataObjectsDownloadFiles {
	s.DownloadUrl = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsDownloadFiles) SetFileName(v string) *ListKnowledgeObjectsResponseBodyDataObjectsDownloadFiles {
	s.FileName = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsDownloadFiles) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeObjectsResponseBodyDataObjectsSourceReferences struct {
	// The source ID. This field exists only when the type is KNOWLEDGE_SOURCE. It represents the SourceId of the KnowledgeSource registered in the current knowledge base.
	//
	// example:
	//
	// src-001
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The category of the source reference. The structure retains KNOWLEDGE_SOURCE. Currently, SourceReferences is an empty array.
	//
	// example:
	//
	// DOCUMENT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsSourceReferences) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeObjectsResponseBodyDataObjectsSourceReferences) GoString() string {
	return s.String()
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsSourceReferences) GetSourceId() *string {
	return s.SourceId
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsSourceReferences) GetType() *string {
	return s.Type
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsSourceReferences) SetSourceId(v string) *ListKnowledgeObjectsResponseBodyDataObjectsSourceReferences {
	s.SourceId = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsSourceReferences) SetType(v string) *ListKnowledgeObjectsResponseBodyDataObjectsSourceReferences {
	s.Type = &v
	return s
}

func (s *ListKnowledgeObjectsResponseBodyDataObjectsSourceReferences) Validate() error {
	return dara.Validate(s)
}
