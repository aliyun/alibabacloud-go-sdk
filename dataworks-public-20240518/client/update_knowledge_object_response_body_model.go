// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateKnowledgeObjectResponseBody
	GetCode() *string
	SetData(v *UpdateKnowledgeObjectResponseBodyData) *UpdateKnowledgeObjectResponseBody
	GetData() *UpdateKnowledgeObjectResponseBodyData
	SetMessage(v string) *UpdateKnowledgeObjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateKnowledgeObjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateKnowledgeObjectResponseBody
	GetSuccess() *bool
}

type UpdateKnowledgeObjectResponseBody struct {
	// The business error code. You can use this code to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *UpdateKnowledgeObjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s UpdateKnowledgeObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateKnowledgeObjectResponseBody) GetData() *UpdateKnowledgeObjectResponseBodyData {
	return s.Data
}

func (s *UpdateKnowledgeObjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateKnowledgeObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKnowledgeObjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateKnowledgeObjectResponseBody) SetCode(v string) *UpdateKnowledgeObjectResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBody) SetData(v *UpdateKnowledgeObjectResponseBodyData) *UpdateKnowledgeObjectResponseBody {
	s.Data = v
	return s
}

func (s *UpdateKnowledgeObjectResponseBody) SetMessage(v string) *UpdateKnowledgeObjectResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBody) SetRequestId(v string) *UpdateKnowledgeObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBody) SetSuccess(v bool) *UpdateKnowledgeObjectResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeObjectResponseBodyData struct {
	// The current valid typed knowledge content. The document or semantic model structure is returned based on the type.
	Content *UpdateKnowledgeObjectResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The time when the object was first created.
	//
	// example:
	//
	// 1790200800000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the knowledge object. This field is editable. Pass an empty string to clear the description.
	//
	// example:
	//
	// Actual paid amount
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The complete specification shards for multi-file models. Each item includes FileName and DownloadUrl. This field is optional.
	DownloadFiles []*UpdateKnowledgeObjectResponseBodyDataDownloadFiles `json:"DownloadFiles,omitempty" xml:"DownloadFiles,omitempty" type:"Repeated"`
	// The temporary download URL for the single-file YAML specification. This field is optional.
	//
	// example:
	//
	// https://example.invalid/short-lived-result-download
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// Indicates whether the current caller has full content editing permissions. The value is true for OWNER and ADMIN and false for VIEWER. Authorization is checked again for write operations.
	Editable *bool `json:"Editable,omitempty" xml:"Editable,omitempty"`
	// The validity period of the download URL, in seconds. This field is optional.
	//
	// example:
	//
	// 300
	ExpiresIn *int32 `json:"ExpiresIn,omitempty" xml:"ExpiresIn,omitempty"`
	// The original file name of the uploaded document. This field is read-only. Renaming the object does not change this value. This field is not returned for external documents and semantic models.
	//
	// example:
	//
	// policy.md
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The ID of the associated knowledge base.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The ID of the knowledge object. For semantic models, use the model ID from the semantic layer. For documents, use the stable ID from the local document table. Callers must treat the ID as opaque and must not use a retrieval ID in its place.
	//
	// example:
	//
	// obj-model-001
	KnowledgeObjectId *string `json:"KnowledgeObjectId,omitempty" xml:"KnowledgeObjectId,omitempty"`
	// The time when the object was last modified, in milliseconds. This value is used for conditional updates of DOCUMENT content. This field does not indicate that optimistic locking updates are supported for semantic models.
	//
	// example:
	//
	// 1791789000000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The user-facing name of the knowledge object. This field is editable. Renaming the object does not change the KnowledgeObjectId.
	//
	// example:
	//
	// orders
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The size of the original uploaded file in bytes. This field is read-only and does not represent the length of the edited content.
	//
	// example:
	//
	// 1024
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
	// The array of source references. An empty array is currently returned.
	SourceReferences []*UpdateKnowledgeObjectResponseBodyDataSourceReferences `json:"SourceReferences,omitempty" xml:"SourceReferences,omitempty" type:"Repeated"`
	// The knowledge object type. The current values are SEMANTIC_MODEL and DOCUMENT. More types may be added in the future.
	//
	// example:
	//
	// SEMANTIC_MODEL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the model that is saved and read in the current operation.
	//
	// example:
	//
	// 1
	VersionNo *int64 `json:"VersionNo,omitempty" xml:"VersionNo,omitempty"`
}

func (s UpdateKnowledgeObjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectResponseBodyData) GetContent() *UpdateKnowledgeObjectResponseBodyDataContent {
	return s.Content
}

func (s *UpdateKnowledgeObjectResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *UpdateKnowledgeObjectResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeObjectResponseBodyData) GetDownloadFiles() []*UpdateKnowledgeObjectResponseBodyDataDownloadFiles {
	return s.DownloadFiles
}

func (s *UpdateKnowledgeObjectResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *UpdateKnowledgeObjectResponseBodyData) GetEditable() *bool {
	return s.Editable
}

func (s *UpdateKnowledgeObjectResponseBodyData) GetExpiresIn() *int32 {
	return s.ExpiresIn
}

func (s *UpdateKnowledgeObjectResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *UpdateKnowledgeObjectResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *UpdateKnowledgeObjectResponseBodyData) GetKnowledgeObjectId() *string {
	return s.KnowledgeObjectId
}

func (s *UpdateKnowledgeObjectResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *UpdateKnowledgeObjectResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeObjectResponseBodyData) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *UpdateKnowledgeObjectResponseBodyData) GetSourceReferences() []*UpdateKnowledgeObjectResponseBodyDataSourceReferences {
	return s.SourceReferences
}

func (s *UpdateKnowledgeObjectResponseBodyData) GetType() *string {
	return s.Type
}

func (s *UpdateKnowledgeObjectResponseBodyData) GetVersionNo() *int64 {
	return s.VersionNo
}

func (s *UpdateKnowledgeObjectResponseBodyData) SetContent(v *UpdateKnowledgeObjectResponseBodyDataContent) *UpdateKnowledgeObjectResponseBodyData {
	s.Content = v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyData) SetCreateTime(v int64) *UpdateKnowledgeObjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyData) SetDescription(v string) *UpdateKnowledgeObjectResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyData) SetDownloadFiles(v []*UpdateKnowledgeObjectResponseBodyDataDownloadFiles) *UpdateKnowledgeObjectResponseBodyData {
	s.DownloadFiles = v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyData) SetDownloadUrl(v string) *UpdateKnowledgeObjectResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyData) SetEditable(v bool) *UpdateKnowledgeObjectResponseBodyData {
	s.Editable = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyData) SetExpiresIn(v int32) *UpdateKnowledgeObjectResponseBodyData {
	s.ExpiresIn = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyData) SetFileName(v string) *UpdateKnowledgeObjectResponseBodyData {
	s.FileName = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyData) SetKnowledgeBaseId(v string) *UpdateKnowledgeObjectResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyData) SetKnowledgeObjectId(v string) *UpdateKnowledgeObjectResponseBodyData {
	s.KnowledgeObjectId = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyData) SetModifyTime(v int64) *UpdateKnowledgeObjectResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyData) SetName(v string) *UpdateKnowledgeObjectResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyData) SetSizeBytes(v int64) *UpdateKnowledgeObjectResponseBodyData {
	s.SizeBytes = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyData) SetSourceReferences(v []*UpdateKnowledgeObjectResponseBodyDataSourceReferences) *UpdateKnowledgeObjectResponseBodyData {
	s.SourceReferences = v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyData) SetType(v string) *UpdateKnowledgeObjectResponseBodyData {
	s.Type = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyData) SetVersionNo(v int64) *UpdateKnowledgeObjectResponseBodyData {
	s.VersionNo = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyData) Validate() error {
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

type UpdateKnowledgeObjectResponseBodyDataContent struct {
	// The document content. This property exists only when Type is set to DOCUMENT.
	Document *UpdateKnowledgeObjectResponseBodyDataContentDocument `json:"Document,omitempty" xml:"Document,omitempty" type:"Struct"`
	// The semantic model content. This property exists only when Type is set to SEMANTIC_MODEL.
	SemanticModel *UpdateKnowledgeObjectResponseBodyDataContentSemanticModel `json:"SemanticModel,omitempty" xml:"SemanticModel,omitempty" type:"Struct"`
}

func (s UpdateKnowledgeObjectResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectResponseBodyDataContent) GetDocument() *UpdateKnowledgeObjectResponseBodyDataContentDocument {
	return s.Document
}

func (s *UpdateKnowledgeObjectResponseBodyDataContent) GetSemanticModel() *UpdateKnowledgeObjectResponseBodyDataContentSemanticModel {
	return s.SemanticModel
}

func (s *UpdateKnowledgeObjectResponseBodyDataContent) SetDocument(v *UpdateKnowledgeObjectResponseBodyDataContentDocument) *UpdateKnowledgeObjectResponseBodyDataContent {
	s.Document = v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContent) SetSemanticModel(v *UpdateKnowledgeObjectResponseBodyDataContentSemanticModel) *UpdateKnowledgeObjectResponseBodyDataContent {
	s.SemanticModel = v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContent) Validate() error {
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

type UpdateKnowledgeObjectResponseBodyDataContentDocument struct {
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

func (s UpdateKnowledgeObjectResponseBodyDataContentDocument) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectResponseBodyDataContentDocument) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentDocument) GetContentFormat() *string {
	return s.ContentFormat
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentDocument) GetText() *string {
	return s.Text
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentDocument) SetContentFormat(v string) *UpdateKnowledgeObjectResponseBodyDataContentDocument {
	s.ContentFormat = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentDocument) SetText(v string) *UpdateKnowledgeObjectResponseBodyDataContentDocument {
	s.Text = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentDocument) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeObjectResponseBodyDataContentSemanticModel struct {
	// The logical datasets included in the model. An empty array is returned if no public datasets are available.
	Datasets []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// The executable metrics in the model. This is an empty array if no metrics are available. Metrics with only candidate formulas are not included.
	Metrics []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The defined relationships in the model expressed by field pairs. This is an empty array if no relationships are available. Relationships are not inferred from graph paths.
	Relationships []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships `json:"Relationships,omitempty" xml:"Relationships,omitempty" type:"Repeated"`
}

func (s UpdateKnowledgeObjectResponseBodyDataContentSemanticModel) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectResponseBodyDataContentSemanticModel) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModel) GetDatasets() []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	return s.Datasets
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModel) GetMetrics() []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	return s.Metrics
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModel) GetRelationships() []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships {
	return s.Relationships
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModel) SetDatasets(v []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModel {
	s.Datasets = v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModel) SetMetrics(v []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModel {
	s.Metrics = v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModel) SetRelationships(v []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModel {
	s.Relationships = v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModel) Validate() error {
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

type UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets struct {
	// The business description of the dataset. This parameter is optional.
	//
	// example:
	//
	// Actual paid amount
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The DataWorks Data Map table entity ID. This field is returned when a reliable association exists. If the ID is omitted for an existing dataset of the same name, the existing ID is retained. Submitted values are processed according to the semantic layer rules for saving data.
	//
	// example:
	//
	// maxcompute-table:::project_name:schema_name:table_name
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The defined and publicly available fields of the dataset. This is an empty array if no field definitions are provided.
	Fields []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The name of the dataset within the model, used for relationship positioning within the model.
	//
	// example:
	//
	// orders
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The explicitly declared primary key field name. This parameter is optional. Each item must exist in the Fields of this dataset.
	PrimaryKey []*string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty" type:"Repeated"`
}

func (s UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GetEntityId() *string {
	return s.EntityId
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GetFields() []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields {
	return s.Fields
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GetPrimaryKey() []*string {
	return s.PrimaryKey
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) SetDescription(v string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) SetEntityId(v string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	s.EntityId = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) SetFields(v []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	s.Fields = v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) SetName(v string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) SetPrimaryKey(v []*string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	s.PrimaryKey = v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) Validate() error {
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

type UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields struct {
	// The data type of the field. This parameter is required for publicly available model fields and is not inferred by the server.
	//
	// example:
	//
	// DECIMAL
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The business description of the field. This parameter is optional.
	//
	// example:
	//
	// Actual paid amount
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// pay_amount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) GetDataType() *string {
	return s.DataType
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) SetDataType(v string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields {
	s.DataType = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) SetDescription(v string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) SetName(v string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics struct {
	// The name of the associated SemanticDataset. This must be locatable within the Datasets of this model.
	//
	// example:
	//
	// orders
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The description of the metric calculation logic. This parameter is optional.
	//
	// example:
	//
	// Actual paid amount minus refund amount
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The executable expression. This must include at least one dialect version and is not auto-generated by the model.
	Expression *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression `json:"Expression,omitempty" xml:"Expression,omitempty" type:"Struct"`
	// The name of the metric.
	//
	// example:
	//
	// Net sales
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The explicitly declared unit of the metric. This parameter is optional and is not returned if no reliable source is available.
	//
	// example:
	//
	// CNY
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GetExpression() *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression {
	return s.Expression
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GetUnit() *string {
	return s.Unit
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) SetDatasetName(v string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	s.DatasetName = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) SetDescription(v string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) SetExpression(v *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	s.Expression = v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) SetName(v string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) SetUnit(v string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	s.Unit = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) Validate() error {
	if s.Expression != nil {
		if err := s.Expression.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression struct {
	// The expressions for the same metric in different SQL dialects. At least one item is required. Dialects cannot be duplicated.
	Dialects []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects `json:"Dialects,omitempty" xml:"Dialects,omitempty" type:"Repeated"`
}

func (s UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) GetDialects() []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects {
	return s.Dialects
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) SetDialects(v []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression {
	s.Dialects = v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) Validate() error {
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

type UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects struct {
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
	// SUM(pay_amount - refund_amount)
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) GetDialect() *string {
	return s.Dialect
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) GetSql() *string {
	return s.Sql
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) SetDialect(v string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects {
	s.Dialect = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) SetSql(v string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects {
	s.Sql = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships struct {
	// The name of the source SemanticDataset.
	//
	// example:
	//
	// orders
	FromDataset *string `json:"FromDataset,omitempty" xml:"FromDataset,omitempty"`
	// The equi-join key pairs for the fields at both ends. At least one item is required. Composite joins are listed item by item.
	JoinKeys []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys `json:"JoinKeys,omitempty" xml:"JoinKeys,omitempty" type:"Repeated"`
	// The name of the relationship within the model.
	//
	// example:
	//
	// orders
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the target SemanticDataset.
	//
	// example:
	//
	// customers
	ToDataset *string `json:"ToDataset,omitempty" xml:"ToDataset,omitempty"`
}

func (s UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) GetFromDataset() *string {
	return s.FromDataset
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) GetJoinKeys() []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys {
	return s.JoinKeys
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) GetToDataset() *string {
	return s.ToDataset
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) SetFromDataset(v string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships {
	s.FromDataset = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) SetJoinKeys(v []*UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships {
	s.JoinKeys = v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) SetName(v string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) SetToDataset(v string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships {
	s.ToDataset = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) Validate() error {
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

type UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys struct {
	// The field name in FromDataset.Fields.
	//
	// example:
	//
	// customer_id
	FromField *string `json:"FromField,omitempty" xml:"FromField,omitempty"`
	// The field name in ToDataset.Fields.
	//
	// example:
	//
	// id
	ToField *string `json:"ToField,omitempty" xml:"ToField,omitempty"`
}

func (s UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) GetFromField() *string {
	return s.FromField
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) GetToField() *string {
	return s.ToField
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) SetFromField(v string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys {
	s.FromField = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) SetToField(v string) *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys {
	s.ToField = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeObjectResponseBodyDataDownloadFiles struct {
	// The temporary download URL for the single-file YAML specification. This field is optional.
	//
	// example:
	//
	// https://example.invalid/short-lived-result-download
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The original file name of the uploaded document. This field is read-only. Renaming the object does not change this value. This field is not returned for external documents and semantic models.
	//
	// example:
	//
	// policy.md
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s UpdateKnowledgeObjectResponseBodyDataDownloadFiles) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectResponseBodyDataDownloadFiles) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectResponseBodyDataDownloadFiles) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *UpdateKnowledgeObjectResponseBodyDataDownloadFiles) GetFileName() *string {
	return s.FileName
}

func (s *UpdateKnowledgeObjectResponseBodyDataDownloadFiles) SetDownloadUrl(v string) *UpdateKnowledgeObjectResponseBodyDataDownloadFiles {
	s.DownloadUrl = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataDownloadFiles) SetFileName(v string) *UpdateKnowledgeObjectResponseBodyDataDownloadFiles {
	s.FileName = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataDownloadFiles) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeObjectResponseBodyDataSourceReferences struct {
	// The source ID of the registered KnowledgeSource in the current knowledge base. This field exists only when the type is KNOWLEDGE_SOURCE.
	//
	// example:
	//
	// src-001
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The category of the source reference. The structure retains KNOWLEDGE_SOURCE. The SourceReferences array is currently empty.
	//
	// example:
	//
	// SEMANTIC_MODEL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateKnowledgeObjectResponseBodyDataSourceReferences) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectResponseBodyDataSourceReferences) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectResponseBodyDataSourceReferences) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateKnowledgeObjectResponseBodyDataSourceReferences) GetType() *string {
	return s.Type
}

func (s *UpdateKnowledgeObjectResponseBodyDataSourceReferences) SetSourceId(v string) *UpdateKnowledgeObjectResponseBodyDataSourceReferences {
	s.SourceId = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataSourceReferences) SetType(v string) *UpdateKnowledgeObjectResponseBodyDataSourceReferences {
	s.Type = &v
	return s
}

func (s *UpdateKnowledgeObjectResponseBodyDataSourceReferences) Validate() error {
	return dara.Validate(s)
}
