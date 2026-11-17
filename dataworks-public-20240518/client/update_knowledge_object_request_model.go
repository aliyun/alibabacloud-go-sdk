// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateKnowledgeObjectRequest
	GetClientToken() *string
	SetContent(v *UpdateKnowledgeObjectRequestContent) *UpdateKnowledgeObjectRequest
	GetContent() *UpdateKnowledgeObjectRequestContent
	SetDescription(v string) *UpdateKnowledgeObjectRequest
	GetDescription() *string
	SetExpectedModifyTime(v int64) *UpdateKnowledgeObjectRequest
	GetExpectedModifyTime() *int64
	SetKnowledgeBaseId(v string) *UpdateKnowledgeObjectRequest
	GetKnowledgeBaseId() *string
	SetKnowledgeObjectId(v string) *UpdateKnowledgeObjectRequest
	GetKnowledgeObjectId() *string
	SetName(v string) *UpdateKnowledgeObjectRequest
	GetName() *string
	SetUploadId(v string) *UpdateKnowledgeObjectRequest
	GetUploadId() *string
	SetVersionNo(v int64) *UpdateKnowledgeObjectRequest
	GetVersionNo() *int64
}

type UpdateKnowledgeObjectRequest struct {
	// The client token. This parameter is supported only for DOCUMENT. Model updates do not guarantee exactly-once remote save notifications.
	//
	// example:
	//
	// edit-document-text-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The typed content update. Object properties are submitted as needed. Arrays are replaced in their entirety when submitted. Unspecified properties are retained.
	Content *UpdateKnowledgeObjectRequestContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The new description. An empty string clears the description. If this parameter is not specified, the original value is retained.
	//
	// example:
	//
	// Actual paid amount
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expected modification time. This parameter is supported only for DOCUMENT. It is required when modifying the body text and optional when modifying metadata.
	//
	// example:
	//
	// 1790200800000
	ExpectedModifyTime *int64 `json:"ExpectedModifyTime,omitempty" xml:"ExpectedModifyTime,omitempty"`
	// The ID of the knowledge base to which the object belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The stable knowledge object ID returned by the Get or List operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// obj-model-001
	KnowledgeObjectId *string `json:"KnowledgeObjectId,omitempty" xml:"KnowledgeObjectId,omitempty"`
	// The new name. This parameter cannot be empty. If this parameter is not specified, the original value is retained.
	//
	// example:
	//
	// orders
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The upload ID. This parameter is supported only for SEMANTIC_MODEL. It replaces the public model content with a complete YAML file and is mutually exclusive with Content.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abd
	UploadId *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
	// The version number. This parameter is required when updating Content or UploadId for SEMANTIC_MODEL. Use the VersionNo returned by the Get operation. This is not an optimistic lock.
	//
	// example:
	//
	// 1
	VersionNo *int64 `json:"VersionNo,omitempty" xml:"VersionNo,omitempty"`
}

func (s UpdateKnowledgeObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateKnowledgeObjectRequest) GetContent() *UpdateKnowledgeObjectRequestContent {
	return s.Content
}

func (s *UpdateKnowledgeObjectRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeObjectRequest) GetExpectedModifyTime() *int64 {
	return s.ExpectedModifyTime
}

func (s *UpdateKnowledgeObjectRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *UpdateKnowledgeObjectRequest) GetKnowledgeObjectId() *string {
	return s.KnowledgeObjectId
}

func (s *UpdateKnowledgeObjectRequest) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeObjectRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *UpdateKnowledgeObjectRequest) GetVersionNo() *int64 {
	return s.VersionNo
}

func (s *UpdateKnowledgeObjectRequest) SetClientToken(v string) *UpdateKnowledgeObjectRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateKnowledgeObjectRequest) SetContent(v *UpdateKnowledgeObjectRequestContent) *UpdateKnowledgeObjectRequest {
	s.Content = v
	return s
}

func (s *UpdateKnowledgeObjectRequest) SetDescription(v string) *UpdateKnowledgeObjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeObjectRequest) SetExpectedModifyTime(v int64) *UpdateKnowledgeObjectRequest {
	s.ExpectedModifyTime = &v
	return s
}

func (s *UpdateKnowledgeObjectRequest) SetKnowledgeBaseId(v string) *UpdateKnowledgeObjectRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *UpdateKnowledgeObjectRequest) SetKnowledgeObjectId(v string) *UpdateKnowledgeObjectRequest {
	s.KnowledgeObjectId = &v
	return s
}

func (s *UpdateKnowledgeObjectRequest) SetName(v string) *UpdateKnowledgeObjectRequest {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeObjectRequest) SetUploadId(v string) *UpdateKnowledgeObjectRequest {
	s.UploadId = &v
	return s
}

func (s *UpdateKnowledgeObjectRequest) SetVersionNo(v int64) *UpdateKnowledgeObjectRequest {
	s.VersionNo = &v
	return s
}

func (s *UpdateKnowledgeObjectRequest) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeObjectRequestContent struct {
	// The document content. This property exists only when Type is set to DOCUMENT.
	Document *UpdateKnowledgeObjectRequestContentDocument `json:"Document,omitempty" xml:"Document,omitempty" type:"Struct"`
	// The semantic model content. This property exists only when Type is set to SEMANTIC_MODEL.
	SemanticModel *UpdateKnowledgeObjectRequestContentSemanticModel `json:"SemanticModel,omitempty" xml:"SemanticModel,omitempty" type:"Struct"`
}

func (s UpdateKnowledgeObjectRequestContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectRequestContent) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectRequestContent) GetDocument() *UpdateKnowledgeObjectRequestContentDocument {
	return s.Document
}

func (s *UpdateKnowledgeObjectRequestContent) GetSemanticModel() *UpdateKnowledgeObjectRequestContentSemanticModel {
	return s.SemanticModel
}

func (s *UpdateKnowledgeObjectRequestContent) SetDocument(v *UpdateKnowledgeObjectRequestContentDocument) *UpdateKnowledgeObjectRequestContent {
	s.Document = v
	return s
}

func (s *UpdateKnowledgeObjectRequestContent) SetSemanticModel(v *UpdateKnowledgeObjectRequestContentSemanticModel) *UpdateKnowledgeObjectRequestContent {
	s.SemanticModel = v
	return s
}

func (s *UpdateKnowledgeObjectRequestContent) Validate() error {
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

type UpdateKnowledgeObjectRequestContentDocument struct {
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

func (s UpdateKnowledgeObjectRequestContentDocument) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectRequestContentDocument) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectRequestContentDocument) GetContentFormat() *string {
	return s.ContentFormat
}

func (s *UpdateKnowledgeObjectRequestContentDocument) GetText() *string {
	return s.Text
}

func (s *UpdateKnowledgeObjectRequestContentDocument) SetContentFormat(v string) *UpdateKnowledgeObjectRequestContentDocument {
	s.ContentFormat = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentDocument) SetText(v string) *UpdateKnowledgeObjectRequestContentDocument {
	s.Text = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentDocument) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeObjectRequestContentSemanticModel struct {
	// The logical datasets contained in the model. This is an empty array if no public datasets are available.
	Datasets []*UpdateKnowledgeObjectRequestContentSemanticModelDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// The executable metrics in the model. This is an empty array if no metrics are available. Metrics with only candidate formulas are not included.
	Metrics []*UpdateKnowledgeObjectRequestContentSemanticModelMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The defined relationships in the model expressed by field pairs. This is an empty array if no relationships are available. Relationships are not inferred from graph paths.
	Relationships []*UpdateKnowledgeObjectRequestContentSemanticModelRelationships `json:"Relationships,omitempty" xml:"Relationships,omitempty" type:"Repeated"`
}

func (s UpdateKnowledgeObjectRequestContentSemanticModel) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectRequestContentSemanticModel) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModel) GetDatasets() []*UpdateKnowledgeObjectRequestContentSemanticModelDatasets {
	return s.Datasets
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModel) GetMetrics() []*UpdateKnowledgeObjectRequestContentSemanticModelMetrics {
	return s.Metrics
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModel) GetRelationships() []*UpdateKnowledgeObjectRequestContentSemanticModelRelationships {
	return s.Relationships
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModel) SetDatasets(v []*UpdateKnowledgeObjectRequestContentSemanticModelDatasets) *UpdateKnowledgeObjectRequestContentSemanticModel {
	s.Datasets = v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModel) SetMetrics(v []*UpdateKnowledgeObjectRequestContentSemanticModelMetrics) *UpdateKnowledgeObjectRequestContentSemanticModel {
	s.Metrics = v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModel) SetRelationships(v []*UpdateKnowledgeObjectRequestContentSemanticModelRelationships) *UpdateKnowledgeObjectRequestContentSemanticModel {
	s.Relationships = v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModel) Validate() error {
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

type UpdateKnowledgeObjectRequestContentSemanticModelDatasets struct {
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
	Fields []*UpdateKnowledgeObjectRequestContentSemanticModelDatasetsFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The name of the dataset within the model, used for relationship positioning within the model.
	//
	// example:
	//
	// orders
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The explicitly declared primary key field name. This parameter is optional. Each item must exist in the Fields of this dataset.
	PrimaryKey []*string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty" type:"Repeated"`
}

func (s UpdateKnowledgeObjectRequestContentSemanticModelDatasets) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectRequestContentSemanticModelDatasets) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasets) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasets) GetEntityId() *string {
	return s.EntityId
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasets) GetFields() []*UpdateKnowledgeObjectRequestContentSemanticModelDatasetsFields {
	return s.Fields
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasets) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasets) GetPrimaryKey() []*string {
	return s.PrimaryKey
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasets) SetDescription(v string) *UpdateKnowledgeObjectRequestContentSemanticModelDatasets {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasets) SetEntityId(v string) *UpdateKnowledgeObjectRequestContentSemanticModelDatasets {
	s.EntityId = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasets) SetFields(v []*UpdateKnowledgeObjectRequestContentSemanticModelDatasetsFields) *UpdateKnowledgeObjectRequestContentSemanticModelDatasets {
	s.Fields = v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasets) SetName(v string) *UpdateKnowledgeObjectRequestContentSemanticModelDatasets {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasets) SetPrimaryKey(v []*string) *UpdateKnowledgeObjectRequestContentSemanticModelDatasets {
	s.PrimaryKey = v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasets) Validate() error {
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

type UpdateKnowledgeObjectRequestContentSemanticModelDatasetsFields struct {
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

func (s UpdateKnowledgeObjectRequestContentSemanticModelDatasetsFields) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectRequestContentSemanticModelDatasetsFields) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasetsFields) GetDataType() *string {
	return s.DataType
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasetsFields) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasetsFields) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasetsFields) SetDataType(v string) *UpdateKnowledgeObjectRequestContentSemanticModelDatasetsFields {
	s.DataType = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasetsFields) SetDescription(v string) *UpdateKnowledgeObjectRequestContentSemanticModelDatasetsFields {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasetsFields) SetName(v string) *UpdateKnowledgeObjectRequestContentSemanticModelDatasetsFields {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelDatasetsFields) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeObjectRequestContentSemanticModelMetrics struct {
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
	Expression *UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpression `json:"Expression,omitempty" xml:"Expression,omitempty" type:"Struct"`
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

func (s UpdateKnowledgeObjectRequestContentSemanticModelMetrics) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectRequestContentSemanticModelMetrics) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetrics) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetrics) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetrics) GetExpression() *UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpression {
	return s.Expression
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetrics) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetrics) GetUnit() *string {
	return s.Unit
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetrics) SetDatasetName(v string) *UpdateKnowledgeObjectRequestContentSemanticModelMetrics {
	s.DatasetName = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetrics) SetDescription(v string) *UpdateKnowledgeObjectRequestContentSemanticModelMetrics {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetrics) SetExpression(v *UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpression) *UpdateKnowledgeObjectRequestContentSemanticModelMetrics {
	s.Expression = v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetrics) SetName(v string) *UpdateKnowledgeObjectRequestContentSemanticModelMetrics {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetrics) SetUnit(v string) *UpdateKnowledgeObjectRequestContentSemanticModelMetrics {
	s.Unit = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetrics) Validate() error {
	if s.Expression != nil {
		if err := s.Expression.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpression struct {
	// The expressions for the same metric in different SQL dialects. At least one item is required. Dialects cannot be duplicated.
	Dialects []*UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects `json:"Dialects,omitempty" xml:"Dialects,omitempty" type:"Repeated"`
}

func (s UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpression) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpression) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpression) GetDialects() []*UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects {
	return s.Dialects
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpression) SetDialects(v []*UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects) *UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpression {
	s.Dialects = v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpression) Validate() error {
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

type UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects struct {
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

func (s UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects) GetDialect() *string {
	return s.Dialect
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects) GetSql() *string {
	return s.Sql
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects) SetDialect(v string) *UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects {
	s.Dialect = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects) SetSql(v string) *UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects {
	s.Sql = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeObjectRequestContentSemanticModelRelationships struct {
	// The name of the source SemanticDataset.
	//
	// example:
	//
	// orders
	FromDataset *string `json:"FromDataset,omitempty" xml:"FromDataset,omitempty"`
	// The equi-join key pairs for the fields at both ends. At least one item is required. Composite joins are listed item by item.
	JoinKeys []*UpdateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys `json:"JoinKeys,omitempty" xml:"JoinKeys,omitempty" type:"Repeated"`
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

func (s UpdateKnowledgeObjectRequestContentSemanticModelRelationships) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectRequestContentSemanticModelRelationships) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelRelationships) GetFromDataset() *string {
	return s.FromDataset
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelRelationships) GetJoinKeys() []*UpdateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys {
	return s.JoinKeys
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelRelationships) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelRelationships) GetToDataset() *string {
	return s.ToDataset
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelRelationships) SetFromDataset(v string) *UpdateKnowledgeObjectRequestContentSemanticModelRelationships {
	s.FromDataset = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelRelationships) SetJoinKeys(v []*UpdateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys) *UpdateKnowledgeObjectRequestContentSemanticModelRelationships {
	s.JoinKeys = v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelRelationships) SetName(v string) *UpdateKnowledgeObjectRequestContentSemanticModelRelationships {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelRelationships) SetToDataset(v string) *UpdateKnowledgeObjectRequestContentSemanticModelRelationships {
	s.ToDataset = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelRelationships) Validate() error {
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

type UpdateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys struct {
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

func (s UpdateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys) GetFromField() *string {
	return s.FromField
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys) GetToField() *string {
	return s.ToField
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys) SetFromField(v string) *UpdateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys {
	s.FromField = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys) SetToField(v string) *UpdateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys {
	s.ToField = &v
	return s
}

func (s *UpdateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys) Validate() error {
	return dara.Validate(s)
}
