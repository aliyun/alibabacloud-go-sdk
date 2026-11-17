// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetKnowledgeObjectResponseBody
	GetCode() *string
	SetData(v *GetKnowledgeObjectResponseBodyData) *GetKnowledgeObjectResponseBody
	GetData() *GetKnowledgeObjectResponseBodyData
	SetMessage(v string) *GetKnowledgeObjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetKnowledgeObjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKnowledgeObjectResponseBody
	GetSuccess() *bool
}

type GetKnowledgeObjectResponseBody struct {
	// The business error code. You can use this code to determine the error type when a request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *GetKnowledgeObjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetKnowledgeObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeObjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetKnowledgeObjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetKnowledgeObjectResponseBody) GetData() *GetKnowledgeObjectResponseBodyData {
	return s.Data
}

func (s *GetKnowledgeObjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKnowledgeObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKnowledgeObjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKnowledgeObjectResponseBody) SetCode(v string) *GetKnowledgeObjectResponseBody {
	s.Code = &v
	return s
}

func (s *GetKnowledgeObjectResponseBody) SetData(v *GetKnowledgeObjectResponseBodyData) *GetKnowledgeObjectResponseBody {
	s.Data = v
	return s
}

func (s *GetKnowledgeObjectResponseBody) SetMessage(v string) *GetKnowledgeObjectResponseBody {
	s.Message = &v
	return s
}

func (s *GetKnowledgeObjectResponseBody) SetRequestId(v string) *GetKnowledgeObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKnowledgeObjectResponseBody) SetSuccess(v bool) *GetKnowledgeObjectResponseBody {
	s.Success = &v
	return s
}

func (s *GetKnowledgeObjectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKnowledgeObjectResponseBodyData struct {
	// The current valid typed knowledge content. The structure returned is a document or semantic model based on the object type.
	Content *GetKnowledgeObjectResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The time when the object was first created. This field is not returned for existing objects whose creation time cannot be reliably determined.
	//
	// example:
	//
	// 1790200800000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The brief description of the object. This field is editable. Pass an empty string to clear the description.
	//
	// example:
	//
	// Sales order analysis model
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The complete shards returned for multi-file specification snapshots. Each item contains a FileName and a DownloadUrl. For single-file snapshots, use the DownloadUrl field.
	DownloadFiles []*GetKnowledgeObjectResponseBodyDataDownloadFiles `json:"DownloadFiles,omitempty" xml:"DownloadFiles,omitempty" type:"Repeated"`
	// The temporary download URL of the uploaded original file for documents. For semantic models, this is the temporary download URL of the complete specification YAML for the current version.
	//
	// example:
	//
	// https://example.invalid/current-model-yaml
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// Indicates whether the current caller has full content editing permissions. The value is true for OWNER and ADMIN and false for VIEWER. Authorization is checked again for write operations.
	Editable *bool `json:"Editable,omitempty" xml:"Editable,omitempty"`
	// The validity period of the download URL, in seconds. The value is fixed at 300.
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
	// The ID of the knowledge base.
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
	// The time when the object was last modified, in milliseconds. This value is used for conditional updates of the document body. This field does not indicate that the semantic model supports optimistic locking.
	//
	// example:
	//
	// 1790200800000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The user-facing name of the object. This field is editable. Renaming the object does not change the knowledge object ID.
	//
	// example:
	//
	// Sales subject model
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The size of the uploaded original file in bytes. This field is read-only and does not reflect the length of the edited text.
	//
	// example:
	//
	// 1024
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
	// The array of source references. An empty array is currently returned.
	SourceReferences []*GetKnowledgeObjectResponseBodyDataSourceReferences `json:"SourceReferences,omitempty" xml:"SourceReferences,omitempty" type:"Repeated"`
	// The object type. Current valid values are SEMANTIC_MODEL and DOCUMENT. This field may be extended in the future.
	//
	// example:
	//
	// SEMANTIC_MODEL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the semantic model read in this request. This version is used to specify the target version for subsequent saves and is not an optimistic lock.
	//
	// example:
	//
	// 1
	VersionNo *int64 `json:"VersionNo,omitempty" xml:"VersionNo,omitempty"`
}

func (s GetKnowledgeObjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKnowledgeObjectResponseBodyData) GetContent() *GetKnowledgeObjectResponseBodyDataContent {
	return s.Content
}

func (s *GetKnowledgeObjectResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetKnowledgeObjectResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetKnowledgeObjectResponseBodyData) GetDownloadFiles() []*GetKnowledgeObjectResponseBodyDataDownloadFiles {
	return s.DownloadFiles
}

func (s *GetKnowledgeObjectResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetKnowledgeObjectResponseBodyData) GetEditable() *bool {
	return s.Editable
}

func (s *GetKnowledgeObjectResponseBodyData) GetExpiresIn() *int32 {
	return s.ExpiresIn
}

func (s *GetKnowledgeObjectResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetKnowledgeObjectResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgeObjectResponseBodyData) GetKnowledgeObjectId() *string {
	return s.KnowledgeObjectId
}

func (s *GetKnowledgeObjectResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetKnowledgeObjectResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetKnowledgeObjectResponseBodyData) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *GetKnowledgeObjectResponseBodyData) GetSourceReferences() []*GetKnowledgeObjectResponseBodyDataSourceReferences {
	return s.SourceReferences
}

func (s *GetKnowledgeObjectResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetKnowledgeObjectResponseBodyData) GetVersionNo() *int64 {
	return s.VersionNo
}

func (s *GetKnowledgeObjectResponseBodyData) SetContent(v *GetKnowledgeObjectResponseBodyDataContent) *GetKnowledgeObjectResponseBodyData {
	s.Content = v
	return s
}

func (s *GetKnowledgeObjectResponseBodyData) SetCreateTime(v int64) *GetKnowledgeObjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyData) SetDescription(v string) *GetKnowledgeObjectResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyData) SetDownloadFiles(v []*GetKnowledgeObjectResponseBodyDataDownloadFiles) *GetKnowledgeObjectResponseBodyData {
	s.DownloadFiles = v
	return s
}

func (s *GetKnowledgeObjectResponseBodyData) SetDownloadUrl(v string) *GetKnowledgeObjectResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyData) SetEditable(v bool) *GetKnowledgeObjectResponseBodyData {
	s.Editable = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyData) SetExpiresIn(v int32) *GetKnowledgeObjectResponseBodyData {
	s.ExpiresIn = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyData) SetFileName(v string) *GetKnowledgeObjectResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyData) SetKnowledgeBaseId(v string) *GetKnowledgeObjectResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyData) SetKnowledgeObjectId(v string) *GetKnowledgeObjectResponseBodyData {
	s.KnowledgeObjectId = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyData) SetModifyTime(v int64) *GetKnowledgeObjectResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyData) SetName(v string) *GetKnowledgeObjectResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyData) SetSizeBytes(v int64) *GetKnowledgeObjectResponseBodyData {
	s.SizeBytes = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyData) SetSourceReferences(v []*GetKnowledgeObjectResponseBodyDataSourceReferences) *GetKnowledgeObjectResponseBodyData {
	s.SourceReferences = v
	return s
}

func (s *GetKnowledgeObjectResponseBodyData) SetType(v string) *GetKnowledgeObjectResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyData) SetVersionNo(v int64) *GetKnowledgeObjectResponseBodyData {
	s.VersionNo = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyData) Validate() error {
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

type GetKnowledgeObjectResponseBodyDataContent struct {
	// The document content. This field is present only when the type is DOCUMENT.
	Document *GetKnowledgeObjectResponseBodyDataContentDocument `json:"Document,omitempty" xml:"Document,omitempty" type:"Struct"`
	// The semantic model content. This field is present only when the type is SEMANTIC_MODEL.
	SemanticModel *GetKnowledgeObjectResponseBodyDataContentSemanticModel `json:"SemanticModel,omitempty" xml:"SemanticModel,omitempty" type:"Struct"`
}

func (s GetKnowledgeObjectResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeObjectResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *GetKnowledgeObjectResponseBodyDataContent) GetDocument() *GetKnowledgeObjectResponseBodyDataContentDocument {
	return s.Document
}

func (s *GetKnowledgeObjectResponseBodyDataContent) GetSemanticModel() *GetKnowledgeObjectResponseBodyDataContentSemanticModel {
	return s.SemanticModel
}

func (s *GetKnowledgeObjectResponseBodyDataContent) SetDocument(v *GetKnowledgeObjectResponseBodyDataContentDocument) *GetKnowledgeObjectResponseBodyDataContent {
	s.Document = v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContent) SetSemanticModel(v *GetKnowledgeObjectResponseBodyDataContentSemanticModel) *GetKnowledgeObjectResponseBodyDataContent {
	s.SemanticModel = v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContent) Validate() error {
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

type GetKnowledgeObjectResponseBodyDataContentDocument struct {
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

func (s GetKnowledgeObjectResponseBodyDataContentDocument) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeObjectResponseBodyDataContentDocument) GoString() string {
	return s.String()
}

func (s *GetKnowledgeObjectResponseBodyDataContentDocument) GetContentFormat() *string {
	return s.ContentFormat
}

func (s *GetKnowledgeObjectResponseBodyDataContentDocument) GetText() *string {
	return s.Text
}

func (s *GetKnowledgeObjectResponseBodyDataContentDocument) SetContentFormat(v string) *GetKnowledgeObjectResponseBodyDataContentDocument {
	s.ContentFormat = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentDocument) SetText(v string) *GetKnowledgeObjectResponseBodyDataContentDocument {
	s.Text = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentDocument) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeObjectResponseBodyDataContentSemanticModel struct {
	// The logical datasets included in the model. An empty array is returned if no public datasets are available.
	Datasets []*GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// The executable metrics in the model. An empty array is returned if no metrics are available. Metrics with only candidate formulas are not included.
	Metrics []*GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The defined relationships expressed by field pairs in the model. An empty array is returned if no relationships are defined. Relationships are not inferred from graph paths.
	Relationships []*GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships `json:"Relationships,omitempty" xml:"Relationships,omitempty" type:"Repeated"`
}

func (s GetKnowledgeObjectResponseBodyDataContentSemanticModel) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeObjectResponseBodyDataContentSemanticModel) GoString() string {
	return s.String()
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModel) GetDatasets() []*GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	return s.Datasets
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModel) GetMetrics() []*GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	return s.Metrics
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModel) GetRelationships() []*GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships {
	return s.Relationships
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModel) SetDatasets(v []*GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) *GetKnowledgeObjectResponseBodyDataContentSemanticModel {
	s.Datasets = v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModel) SetMetrics(v []*GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) *GetKnowledgeObjectResponseBodyDataContentSemanticModel {
	s.Metrics = v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModel) SetRelationships(v []*GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) *GetKnowledgeObjectResponseBodyDataContentSemanticModel {
	s.Relationships = v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModel) Validate() error {
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

type GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets struct {
	// The business description of the dataset. This field is optional.
	//
	// example:
	//
	// Actual payment amount
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The DataWorks Data Map table entity ID. This field is returned when a reliable association exists. If the ID is omitted for an existing dataset of the same name, the existing ID is retained. Submitted values are processed according to the semantic layer rules for saving data.
	//
	// example:
	//
	// maxcompute-table:::project_name:schema_name:table_name
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The defined and publicly accessible fields of the dataset. An empty array is returned if no field definitions are provided.
	Fields []*GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The name of the dataset within the model. This name is used to locate relationships within the model.
	//
	// example:
	//
	// orders
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The explicitly declared primary key field names. This field is optional. Each field must exist in the Fields array of the dataset.
	PrimaryKey []*string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty" type:"Repeated"`
}

func (s GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GoString() string {
	return s.String()
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GetDescription() *string {
	return s.Description
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GetEntityId() *string {
	return s.EntityId
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GetFields() []*GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields {
	return s.Fields
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GetName() *string {
	return s.Name
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) GetPrimaryKey() []*string {
	return s.PrimaryKey
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) SetDescription(v string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	s.Description = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) SetEntityId(v string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	s.EntityId = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) SetFields(v []*GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	s.Fields = v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) SetName(v string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	s.Name = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) SetPrimaryKey(v []*string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets {
	s.PrimaryKey = v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasets) Validate() error {
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

type GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields struct {
	// The data type of the field. This field is required for publicly accessible model fields and is not inferred by the server.
	//
	// example:
	//
	// BIGINT
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The business description of the field. This field is optional.
	//
	// example:
	//
	// Actual payment amount
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// customer_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) GoString() string {
	return s.String()
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) GetDataType() *string {
	return s.DataType
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) GetDescription() *string {
	return s.Description
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) GetName() *string {
	return s.Name
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) SetDataType(v string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields {
	s.DataType = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) SetDescription(v string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields {
	s.Description = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) SetName(v string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields {
	s.Name = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelDatasetsFields) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics struct {
	// The name of the dataset to which the metric belongs. This name must match a dataset in the Datasets array of the model.
	//
	// example:
	//
	// orders
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The definition of the metric. This field is optional.
	//
	// example:
	//
	// Sales order analysis model
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The executable expression. This field must contain at least one dialect version and is not auto-generated by the model.
	Expression *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression `json:"Expression,omitempty" xml:"Expression,omitempty" type:"Struct"`
	// The name of the metric.
	//
	// example:
	//
	// Sales amount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The explicitly declared unit of the metric. This field is optional and is not returned if no reliable source is available.
	//
	// example:
	//
	// CNY
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GoString() string {
	return s.String()
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GetDescription() *string {
	return s.Description
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GetExpression() *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression {
	return s.Expression
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GetName() *string {
	return s.Name
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) GetUnit() *string {
	return s.Unit
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) SetDatasetName(v string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	s.DatasetName = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) SetDescription(v string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	s.Description = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) SetExpression(v *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	s.Expression = v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) SetName(v string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	s.Name = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) SetUnit(v string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics {
	s.Unit = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetrics) Validate() error {
	if s.Expression != nil {
		if err := s.Expression.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression struct {
	// The expressions of the metric in different SQL dialects. At least one item is required. Duplicate dialects are not allowed.
	Dialects []*GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects `json:"Dialects,omitempty" xml:"Dialects,omitempty" type:"Repeated"`
}

func (s GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) GoString() string {
	return s.String()
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) GetDialects() []*GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects {
	return s.Dialects
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) SetDialects(v []*GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression {
	s.Dialects = v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpression) Validate() error {
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

type GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects struct {
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

func (s GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) GoString() string {
	return s.String()
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) GetDialect() *string {
	return s.Dialect
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) GetSql() *string {
	return s.Sql
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) SetDialect(v string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects {
	s.Dialect = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) SetSql(v string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects {
	s.Sql = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelMetricsExpressionDialects) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships struct {
	// The name of the source dataset.
	//
	// example:
	//
	// orders
	FromDataset *string `json:"FromDataset,omitempty" xml:"FromDataset,omitempty"`
	// The equi-join key pairs for the fields at both ends. At least one pair is required. Composite joins are listed as individual pairs.
	JoinKeys []*GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys `json:"JoinKeys,omitempty" xml:"JoinKeys,omitempty" type:"Repeated"`
	// The name of the relationship within the model.
	//
	// example:
	//
	// orders_to_customers
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the target dataset.
	//
	// example:
	//
	// customers
	ToDataset *string `json:"ToDataset,omitempty" xml:"ToDataset,omitempty"`
}

func (s GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) GoString() string {
	return s.String()
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) GetFromDataset() *string {
	return s.FromDataset
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) GetJoinKeys() []*GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys {
	return s.JoinKeys
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) GetName() *string {
	return s.Name
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) GetToDataset() *string {
	return s.ToDataset
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) SetFromDataset(v string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships {
	s.FromDataset = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) SetJoinKeys(v []*GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships {
	s.JoinKeys = v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) SetName(v string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships {
	s.Name = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) SetToDataset(v string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships {
	s.ToDataset = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationships) Validate() error {
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

type GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys struct {
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

func (s GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) GoString() string {
	return s.String()
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) GetFromField() *string {
	return s.FromField
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) GetToField() *string {
	return s.ToField
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) SetFromField(v string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys {
	s.FromField = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) SetToField(v string) *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys {
	s.ToField = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataContentSemanticModelRelationshipsJoinKeys) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeObjectResponseBodyDataDownloadFiles struct {
	// The temporary HTTPS download URL of the shard. The validity period is specified by the ExpiresIn field.
	//
	// example:
	//
	// https://example.invalid/current-model-yaml
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The file name of the specification YAML shard. Multiple shards form the complete model and do not represent the object identity.
	//
	// example:
	//
	// policy.md
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetKnowledgeObjectResponseBodyDataDownloadFiles) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeObjectResponseBodyDataDownloadFiles) GoString() string {
	return s.String()
}

func (s *GetKnowledgeObjectResponseBodyDataDownloadFiles) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetKnowledgeObjectResponseBodyDataDownloadFiles) GetFileName() *string {
	return s.FileName
}

func (s *GetKnowledgeObjectResponseBodyDataDownloadFiles) SetDownloadUrl(v string) *GetKnowledgeObjectResponseBodyDataDownloadFiles {
	s.DownloadUrl = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataDownloadFiles) SetFileName(v string) *GetKnowledgeObjectResponseBodyDataDownloadFiles {
	s.FileName = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataDownloadFiles) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeObjectResponseBodyDataSourceReferences struct {
	// The source ID registered in the knowledge base. This field is present only when the type is KNOWLEDGE_SOURCE.
	//
	// example:
	//
	// src-001
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The type of the source reference. The structure reserves KNOWLEDGE_SOURCE. Currently, the SourceReferences array is empty.
	//
	// example:
	//
	// SEMANTIC_MODEL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetKnowledgeObjectResponseBodyDataSourceReferences) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeObjectResponseBodyDataSourceReferences) GoString() string {
	return s.String()
}

func (s *GetKnowledgeObjectResponseBodyDataSourceReferences) GetSourceId() *string {
	return s.SourceId
}

func (s *GetKnowledgeObjectResponseBodyDataSourceReferences) GetType() *string {
	return s.Type
}

func (s *GetKnowledgeObjectResponseBodyDataSourceReferences) SetSourceId(v string) *GetKnowledgeObjectResponseBodyDataSourceReferences {
	s.SourceId = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataSourceReferences) SetType(v string) *GetKnowledgeObjectResponseBodyDataSourceReferences {
	s.Type = &v
	return s
}

func (s *GetKnowledgeObjectResponseBodyDataSourceReferences) Validate() error {
	return dara.Validate(s)
}
