// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKnowledgeObjectResponseBody
	GetCode() *string
	SetData(v *CreateKnowledgeObjectResponseBodyData) *CreateKnowledgeObjectResponseBody
	GetData() *CreateKnowledgeObjectResponseBodyData
	SetMessage(v string) *CreateKnowledgeObjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateKnowledgeObjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateKnowledgeObjectResponseBody
	GetSuccess() *bool
}

type CreateKnowledgeObjectResponseBody struct {
	// The business error code. You can use this code to determine the error type when a request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *CreateKnowledgeObjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// The request parameters are invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateKnowledgeObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKnowledgeObjectResponseBody) GetData() *CreateKnowledgeObjectResponseBodyData {
	return s.Data
}

func (s *CreateKnowledgeObjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKnowledgeObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeObjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateKnowledgeObjectResponseBody) SetCode(v string) *CreateKnowledgeObjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBody) SetData(v *CreateKnowledgeObjectResponseBodyData) *CreateKnowledgeObjectResponseBody {
	s.Data = v
	return s
}

func (s *CreateKnowledgeObjectResponseBody) SetMessage(v string) *CreateKnowledgeObjectResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBody) SetRequestId(v string) *CreateKnowledgeObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBody) SetSuccess(v bool) *CreateKnowledgeObjectResponseBody {
	s.Success = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKnowledgeObjectResponseBodyData struct {
	// The current valid typed knowledge content. Returns the document or semantic model structure based on the Type.
	Content *CreateKnowledgeObjectResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The time when the object was first created.
	//
	// example:
	//
	// 1790200800000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description. This field is editable. Pass an empty string to clear the description.
	//
	// example:
	//
	// Sales metrics and policy document
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The download items for multi-file specification YAMLs, including file names and temporary download URLs.
	DownloadFiles []*CreateKnowledgeObjectResponseBodyDataDownloadFiles `json:"DownloadFiles,omitempty" xml:"DownloadFiles,omitempty" type:"Repeated"`
	// The temporary download URL for the original uploaded file of the DOCUMENT or the current specification YAML of the SEMANTIC_MODEL returned by GetKnowledgeObject. This is not a persistent property.
	//
	// example:
	//
	// https://example.invalid/short-lived-result-download
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// Indicates whether the current caller has full content editing permissions. The value is true for OWNER and ADMIN and false for VIEWER. Authorization is checked again for write operations.
	Editable *bool `json:"Editable,omitempty" xml:"Editable,omitempty"`
	// The validity period of the download URL, in seconds. This field is returned together with DownloadUrl.
	//
	// example:
	//
	// 300
	ExpiresIn *int32 `json:"ExpiresIn,omitempty" xml:"ExpiresIn,omitempty"`
	// The original file name of the uploaded document. This field is read-only. Renaming the object does not change this value. This field is not returned for external documents and semantic models.
	//
	// example:
	//
	// Business_Description.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The ID of the knowledge base to which the object belongs.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The ID of the knowledge object. For semantic models, use the model ID from the semantic layer. For documents, use the stable ID from the local document table. Callers must treat the ID as opaque and must not use a retrieval ID in its place.
	//
	// example:
	//
	// obj-document-001
	KnowledgeObjectId *string `json:"KnowledgeObjectId,omitempty" xml:"KnowledgeObjectId,omitempty"`
	// The time when the object was last modified, in milliseconds. This value is used for conditional updates of the DOCUMENT body. This field does not indicate that optimistic locking updates are supported for semantic models.
	//
	// example:
	//
	// 1790200800000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The user-facing name. This field is editable. Renaming the object does not change the KnowledgeObjectId.
	//
	// example:
	//
	// Business description
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The size of the uploaded original file in bytes. This field is read-only and does not represent the length of the edited body.
	//
	// example:
	//
	// 10240
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
	// The array of source references. Currently, an empty array is returned.
	SourceReferences []*CreateKnowledgeObjectResponseBodyDataSourceReferences `json:"SourceReferences,omitempty" xml:"SourceReferences,omitempty" type:"Repeated"`
	// The object type. The value is fixed to DOCUMENT.
	//
	// example:
	//
	// DOCUMENT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The current version of the SEMANTIC_MODEL being read. Pass this value when updating the content to edit the version in place.
	//
	// example:
	//
	// 1
	VersionNo *int64 `json:"VersionNo,omitempty" xml:"VersionNo,omitempty"`
}

func (s CreateKnowledgeObjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectResponseBodyData) GetContent() *CreateKnowledgeObjectResponseBodyDataContent {
	return s.Content
}

func (s *CreateKnowledgeObjectResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateKnowledgeObjectResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeObjectResponseBodyData) GetDownloadFiles() []*CreateKnowledgeObjectResponseBodyDataDownloadFiles {
	return s.DownloadFiles
}

func (s *CreateKnowledgeObjectResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *CreateKnowledgeObjectResponseBodyData) GetEditable() *bool {
	return s.Editable
}

func (s *CreateKnowledgeObjectResponseBodyData) GetExpiresIn() *int32 {
	return s.ExpiresIn
}

func (s *CreateKnowledgeObjectResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *CreateKnowledgeObjectResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *CreateKnowledgeObjectResponseBodyData) GetKnowledgeObjectId() *string {
	return s.KnowledgeObjectId
}

func (s *CreateKnowledgeObjectResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *CreateKnowledgeObjectResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeObjectResponseBodyData) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *CreateKnowledgeObjectResponseBodyData) GetSourceReferences() []*CreateKnowledgeObjectResponseBodyDataSourceReferences {
	return s.SourceReferences
}

func (s *CreateKnowledgeObjectResponseBodyData) GetType() *string {
	return s.Type
}

func (s *CreateKnowledgeObjectResponseBodyData) GetVersionNo() *int64 {
	return s.VersionNo
}

func (s *CreateKnowledgeObjectResponseBodyData) SetContent(v *CreateKnowledgeObjectResponseBodyDataContent) *CreateKnowledgeObjectResponseBodyData {
	s.Content = v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyData) SetCreateTime(v int64) *CreateKnowledgeObjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyData) SetDescription(v string) *CreateKnowledgeObjectResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyData) SetDownloadFiles(v []*CreateKnowledgeObjectResponseBodyDataDownloadFiles) *CreateKnowledgeObjectResponseBodyData {
	s.DownloadFiles = v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyData) SetDownloadUrl(v string) *CreateKnowledgeObjectResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyData) SetEditable(v bool) *CreateKnowledgeObjectResponseBodyData {
	s.Editable = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyData) SetExpiresIn(v int32) *CreateKnowledgeObjectResponseBodyData {
	s.ExpiresIn = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyData) SetFileName(v string) *CreateKnowledgeObjectResponseBodyData {
	s.FileName = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyData) SetKnowledgeBaseId(v string) *CreateKnowledgeObjectResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyData) SetKnowledgeObjectId(v string) *CreateKnowledgeObjectResponseBodyData {
	s.KnowledgeObjectId = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyData) SetModifyTime(v int64) *CreateKnowledgeObjectResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyData) SetName(v string) *CreateKnowledgeObjectResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyData) SetSizeBytes(v int64) *CreateKnowledgeObjectResponseBodyData {
	s.SizeBytes = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyData) SetSourceReferences(v []*CreateKnowledgeObjectResponseBodyDataSourceReferences) *CreateKnowledgeObjectResponseBodyData {
	s.SourceReferences = v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyData) SetType(v string) *CreateKnowledgeObjectResponseBodyData {
	s.Type = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyData) SetVersionNo(v int64) *CreateKnowledgeObjectResponseBodyData {
	s.VersionNo = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyData) Validate() error {
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

type CreateKnowledgeObjectResponseBodyDataContent struct {
	// This operation returns only the document content.
	Document *CreateKnowledgeObjectResponseBodyDataContentDocument `json:"Document,omitempty" xml:"Document,omitempty" type:"Struct"`
	// The semantic model. This field exists only when Type is set to SEMANTIC_MODEL.
	SemanticModel *CreateKnowledgeObjectResponseBodyDataContentSemanticModel `json:"SemanticModel,omitempty" xml:"SemanticModel,omitempty" type:"Struct"`
}

func (s CreateKnowledgeObjectResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectResponseBodyDataContent) GetDocument() *CreateKnowledgeObjectResponseBodyDataContentDocument {
	return s.Document
}

func (s *CreateKnowledgeObjectResponseBodyDataContent) GetSemanticModel() *CreateKnowledgeObjectResponseBodyDataContentSemanticModel {
	return s.SemanticModel
}

func (s *CreateKnowledgeObjectResponseBodyDataContent) SetDocument(v *CreateKnowledgeObjectResponseBodyDataContentDocument) *CreateKnowledgeObjectResponseBodyDataContent {
	s.Document = v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContent) SetSemanticModel(v *CreateKnowledgeObjectResponseBodyDataContentSemanticModel) *CreateKnowledgeObjectResponseBodyDataContent {
	s.SemanticModel = v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContent) Validate() error {
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

type CreateKnowledgeObjectResponseBodyDataContentDocument struct {
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
	// Updated full text
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s CreateKnowledgeObjectResponseBodyDataContentDocument) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectResponseBodyDataContentDocument) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectResponseBodyDataContentDocument) GetContentFormat() *string {
	return s.ContentFormat
}

func (s *CreateKnowledgeObjectResponseBodyDataContentDocument) GetText() *string {
	return s.Text
}

func (s *CreateKnowledgeObjectResponseBodyDataContentDocument) SetContentFormat(v string) *CreateKnowledgeObjectResponseBodyDataContentDocument {
	s.ContentFormat = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentDocument) SetText(v string) *CreateKnowledgeObjectResponseBodyDataContentDocument {
	s.Text = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentDocument) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeObjectResponseBodyDataContentSemanticModel struct {
	// The logical datasets included in the model. An empty array is returned if no public datasets are available.
	Datasets []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// The executable metrics in the model. If no metrics exist, this is an empty array. Metrics with only candidate formulas are not included.
	Metrics []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The defined relationships in the model expressed by field pairs. If no relationships exist, this is an empty array. Relationships are not guessed from graph paths.
	Relationships []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships `json:"Relationships,omitempty" xml:"Relationships,omitempty" type:"Repeated"`
}

func (s CreateKnowledgeObjectResponseBodyDataContentSemanticModel) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectResponseBodyDataContentSemanticModel) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModel) GetDatasets() []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	return s.Datasets
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModel) GetMetrics() []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	return s.Metrics
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModel) GetRelationships() []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships {
	return s.Relationships
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModel) SetDatasets(v []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) *CreateKnowledgeObjectResponseBodyDataContentSemanticModel {
	s.Datasets = v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModel) SetMetrics(v []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) *CreateKnowledgeObjectResponseBodyDataContentSemanticModel {
	s.Metrics = v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModel) SetRelationships(v []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) *CreateKnowledgeObjectResponseBodyDataContentSemanticModel {
	s.Relationships = v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModel) Validate() error {
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

type CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets struct {
	// The description. This field is editable. Pass an empty string to clear the description.
	//
	// example:
	//
	// Sales metrics and policy document
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The metadata entity ID returned by Data Map. The client passes this ID as-is without concatenating it.
	//
	// example:
	//
	// maxcompute-table:::project_name:schema_name:table_name
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The defined and publicly available fields of the dataset. If no field definitions are provided, this is an empty array.
	Fields []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The user-facing name. This field is editable. Renaming the object does not change the KnowledgeObjectId.
	//
	// example:
	//
	// Business description
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The explicitly declared primary key field name. This parameter is optional. Each item must exist in the Fields of this dataset.
	PrimaryKey []*string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty" type:"Repeated"`
}

func (s CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GetEntityId() *string {
	return s.EntityId
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GetFields() []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields {
	return s.Fields
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GetPrimaryKey() []*string {
	return s.PrimaryKey
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) SetDescription(v string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) SetEntityId(v string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	s.EntityId = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) SetFields(v []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	s.Fields = v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) SetName(v string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) SetPrimaryKey(v []*string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	s.PrimaryKey = v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) Validate() error {
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

type CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields struct {
	// The data type of the field. This must be provided for publicly available model fields and is not guessed by the server.
	//
	// example:
	//
	// BIGINT
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The description. This field is editable. Pass an empty string to clear the description.
	//
	// example:
	//
	// Sales metrics and policy document
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user-facing name. This field is editable. Renaming the object does not change the KnowledgeObjectId.
	//
	// example:
	//
	// Business description
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) GetDataType() *string {
	return s.DataType
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) SetDataType(v string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields {
	s.DataType = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) SetDescription(v string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) SetName(v string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics struct {
	// The SemanticDataset.Name to which the metric belongs. It must be locatable in the Datasets of this model.
	//
	// example:
	//
	// orders
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The description. This field is editable. Pass an empty string to clear the description.
	//
	// example:
	//
	// Sales metrics and policy document
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The executable expression. It must include at least one dialect version and is not auto-completed by the model.
	Expression *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression `json:"Expression,omitempty" xml:"Expression,omitempty" type:"Struct"`
	// The user-facing name. This field is editable. Renaming the object does not change the KnowledgeObjectId.
	//
	// example:
	//
	// Business description
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The explicitly declared unit of the metric. This parameter is optional. It is not returned when no reliable source is available.
	//
	// example:
	//
	// CNY
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GetExpression() *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression {
	return s.Expression
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GetUnit() *string {
	return s.Unit
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) SetDatasetName(v string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	s.DatasetName = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) SetDescription(v string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) SetExpression(v *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	s.Expression = v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) SetName(v string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) SetUnit(v string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	s.Unit = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) Validate() error {
	if s.Expression != nil {
		if err := s.Expression.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression struct {
	// The expressions of the same metric in different SQL dialects. At least one item is required. The same dialect is not repeated.
	Dialects []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects `json:"Dialects,omitempty" xml:"Dialects,omitempty" type:"Repeated"`
}

func (s CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) GetDialects() []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects {
	return s.Dialects
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) SetDialects(v []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression {
	s.Dialects = v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) Validate() error {
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

type CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects struct {
	// The SQL dialect, such as MaxCompute or ANSI_SQL.
	//
	// example:
	//
	// MaxCompute
	Dialect *string `json:"Dialect,omitempty" xml:"Dialect,omitempty"`
	// The non-empty calculation expression in the specified dialect. It is not a complete query statement.
	//
	// example:
	//
	// SUM(pay_amount)
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) GetDialect() *string {
	return s.Dialect
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) GetSql() *string {
	return s.Sql
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) SetDialect(v string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects {
	s.Dialect = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) SetSql(v string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects {
	s.Sql = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships struct {
	// The SemanticDataset.Name of the starting point.
	//
	// example:
	//
	// orders
	FromDataset *string `json:"FromDataset,omitempty" xml:"FromDataset,omitempty"`
	// The equi-join pairs of fields at both ends of the relationship. For composite relationships, fields are paired item by item.
	JoinKeys []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys `json:"JoinKeys,omitempty" xml:"JoinKeys,omitempty" type:"Repeated"`
	// The user-facing name. This field is editable. Renaming the object does not change the KnowledgeObjectId.
	//
	// example:
	//
	// Business description
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The SemanticDataset.Name of the ending point.
	//
	// example:
	//
	// customers
	ToDataset *string `json:"ToDataset,omitempty" xml:"ToDataset,omitempty"`
}

func (s CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) GetFromDataset() *string {
	return s.FromDataset
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) GetJoinKeys() []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys {
	return s.JoinKeys
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) GetToDataset() *string {
	return s.ToDataset
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) SetFromDataset(v string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships {
	s.FromDataset = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) SetJoinKeys(v []*CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships {
	s.JoinKeys = v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) SetName(v string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) SetToDataset(v string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships {
	s.ToDataset = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) Validate() error {
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

type CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys struct {
	// The name of the field in the starting dataset that participates in the relationship. It must be defined in the Fields of the dataset corresponding to FromDataset.
	//
	// example:
	//
	// customer_id
	FromField *string `json:"FromField,omitempty" xml:"FromField,omitempty"`
	// The name of the field in the ending dataset that participates in the relationship. It must be defined in the Fields of the dataset corresponding to ToDataset.
	//
	// example:
	//
	// id
	ToField *string `json:"ToField,omitempty" xml:"ToField,omitempty"`
}

func (s CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) GetFromField() *string {
	return s.FromField
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) GetToField() *string {
	return s.ToField
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) SetFromField(v string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys {
	s.FromField = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) SetToField(v string) *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys {
	s.ToField = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeObjectResponseBodyDataDownloadFiles struct {
	// The temporary download URL for the specification YAML file. The validity period is specified by ExpiresIn.
	//
	// example:
	//
	// https://example.invalid/short-lived-result-download
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The original file name of the uploaded document. This field is read-only. Renaming the object does not change this value. This field is not returned for external documents and semantic models.
	//
	// example:
	//
	// Business_Description.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s CreateKnowledgeObjectResponseBodyDataDownloadFiles) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectResponseBodyDataDownloadFiles) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectResponseBodyDataDownloadFiles) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *CreateKnowledgeObjectResponseBodyDataDownloadFiles) GetFileName() *string {
	return s.FileName
}

func (s *CreateKnowledgeObjectResponseBodyDataDownloadFiles) SetDownloadUrl(v string) *CreateKnowledgeObjectResponseBodyDataDownloadFiles {
	s.DownloadUrl = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataDownloadFiles) SetFileName(v string) *CreateKnowledgeObjectResponseBodyDataDownloadFiles {
	s.FileName = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataDownloadFiles) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeObjectResponseBodyDataSourceReferences struct {
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

func (s CreateKnowledgeObjectResponseBodyDataSourceReferences) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectResponseBodyDataSourceReferences) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectResponseBodyDataSourceReferences) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateKnowledgeObjectResponseBodyDataSourceReferences) GetType() *string {
	return s.Type
}

func (s *CreateKnowledgeObjectResponseBodyDataSourceReferences) SetSourceId(v string) *CreateKnowledgeObjectResponseBodyDataSourceReferences {
	s.SourceId = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataSourceReferences) SetType(v string) *CreateKnowledgeObjectResponseBodyDataSourceReferences {
	s.Type = &v
	return s
}

func (s *CreateKnowledgeObjectResponseBodyDataSourceReferences) Validate() error {
	return dara.Validate(s)
}
