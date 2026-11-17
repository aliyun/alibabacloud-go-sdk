// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateKnowledgeObjectRequest
	GetClientToken() *string
	SetContent(v *CreateKnowledgeObjectRequestContent) *CreateKnowledgeObjectRequest
	GetContent() *CreateKnowledgeObjectRequestContent
	SetDescription(v string) *CreateKnowledgeObjectRequest
	GetDescription() *string
	SetKnowledgeBaseId(v string) *CreateKnowledgeObjectRequest
	GetKnowledgeBaseId() *string
	SetName(v string) *CreateKnowledgeObjectRequest
	GetName() *string
	SetType(v string) *CreateKnowledgeObjectRequest
	GetType() *string
	SetUploadId(v string) *CreateKnowledgeObjectRequest
	GetUploadId() *string
}

type CreateKnowledgeObjectRequest struct {
	// The idempotency token for the creation request. Keep the original request unchanged when retrying.
	//
	// example:
	//
	// create-document-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Only Content.Document is supported. It includes Text and ContentFormat (MARKDOWN or TEXT).
	Content *CreateKnowledgeObjectRequestContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The document description.
	//
	// example:
	//
	// Sales metrics and policy document
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the target knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The document name. Existing objects are not overwritten based on the name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Business description
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the object. For this operation, the value is fixed to DOCUMENT.
	//
	// This parameter is required.
	//
	// example:
	//
	// DOCUMENT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the completed PUT upload session for this knowledge base. This parameter is mutually exclusive with Content.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abd
	UploadId *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s CreateKnowledgeObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateKnowledgeObjectRequest) GetContent() *CreateKnowledgeObjectRequestContent {
	return s.Content
}

func (s *CreateKnowledgeObjectRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeObjectRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *CreateKnowledgeObjectRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeObjectRequest) GetType() *string {
	return s.Type
}

func (s *CreateKnowledgeObjectRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *CreateKnowledgeObjectRequest) SetClientToken(v string) *CreateKnowledgeObjectRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateKnowledgeObjectRequest) SetContent(v *CreateKnowledgeObjectRequestContent) *CreateKnowledgeObjectRequest {
	s.Content = v
	return s
}

func (s *CreateKnowledgeObjectRequest) SetDescription(v string) *CreateKnowledgeObjectRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeObjectRequest) SetKnowledgeBaseId(v string) *CreateKnowledgeObjectRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *CreateKnowledgeObjectRequest) SetName(v string) *CreateKnowledgeObjectRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeObjectRequest) SetType(v string) *CreateKnowledgeObjectRequest {
	s.Type = &v
	return s
}

func (s *CreateKnowledgeObjectRequest) SetUploadId(v string) *CreateKnowledgeObjectRequest {
	s.UploadId = &v
	return s
}

func (s *CreateKnowledgeObjectRequest) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKnowledgeObjectRequestContent struct {
	// This operation returns only the document content.
	Document *CreateKnowledgeObjectRequestContentDocument `json:"Document,omitempty" xml:"Document,omitempty" type:"Struct"`
	// This operation does not accept this branch. Use a build task to import YAML for new semantic models.
	SemanticModel *CreateKnowledgeObjectRequestContentSemanticModel `json:"SemanticModel,omitempty" xml:"SemanticModel,omitempty" type:"Struct"`
}

func (s CreateKnowledgeObjectRequestContent) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectRequestContent) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectRequestContent) GetDocument() *CreateKnowledgeObjectRequestContentDocument {
	return s.Document
}

func (s *CreateKnowledgeObjectRequestContent) GetSemanticModel() *CreateKnowledgeObjectRequestContentSemanticModel {
	return s.SemanticModel
}

func (s *CreateKnowledgeObjectRequestContent) SetDocument(v *CreateKnowledgeObjectRequestContentDocument) *CreateKnowledgeObjectRequestContent {
	s.Document = v
	return s
}

func (s *CreateKnowledgeObjectRequestContent) SetSemanticModel(v *CreateKnowledgeObjectRequestContentSemanticModel) *CreateKnowledgeObjectRequestContent {
	s.SemanticModel = v
	return s
}

func (s *CreateKnowledgeObjectRequestContent) Validate() error {
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

type CreateKnowledgeObjectRequestContentDocument struct {
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

func (s CreateKnowledgeObjectRequestContentDocument) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectRequestContentDocument) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectRequestContentDocument) GetContentFormat() *string {
	return s.ContentFormat
}

func (s *CreateKnowledgeObjectRequestContentDocument) GetText() *string {
	return s.Text
}

func (s *CreateKnowledgeObjectRequestContentDocument) SetContentFormat(v string) *CreateKnowledgeObjectRequestContentDocument {
	s.ContentFormat = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentDocument) SetText(v string) *CreateKnowledgeObjectRequestContentDocument {
	s.Text = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentDocument) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeObjectRequestContentSemanticModel struct {
	// The logical datasets included in the model. If no public datasets exist, this is an empty array.
	Datasets []*CreateKnowledgeObjectRequestContentSemanticModelDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// The executable metrics in the model. If no metrics exist, this is an empty array. Metrics with only candidate formulas are not included.
	Metrics []*CreateKnowledgeObjectRequestContentSemanticModelMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The defined relationships in the model expressed by field pairs. If no relationships exist, this is an empty array. Relationships are not guessed from graph paths.
	Relationships []*CreateKnowledgeObjectRequestContentSemanticModelRelationships `json:"Relationships,omitempty" xml:"Relationships,omitempty" type:"Repeated"`
}

func (s CreateKnowledgeObjectRequestContentSemanticModel) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectRequestContentSemanticModel) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectRequestContentSemanticModel) GetDatasets() []*CreateKnowledgeObjectRequestContentSemanticModelDatasets {
	return s.Datasets
}

func (s *CreateKnowledgeObjectRequestContentSemanticModel) GetMetrics() []*CreateKnowledgeObjectRequestContentSemanticModelMetrics {
	return s.Metrics
}

func (s *CreateKnowledgeObjectRequestContentSemanticModel) GetRelationships() []*CreateKnowledgeObjectRequestContentSemanticModelRelationships {
	return s.Relationships
}

func (s *CreateKnowledgeObjectRequestContentSemanticModel) SetDatasets(v []*CreateKnowledgeObjectRequestContentSemanticModelDatasets) *CreateKnowledgeObjectRequestContentSemanticModel {
	s.Datasets = v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModel) SetMetrics(v []*CreateKnowledgeObjectRequestContentSemanticModelMetrics) *CreateKnowledgeObjectRequestContentSemanticModel {
	s.Metrics = v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModel) SetRelationships(v []*CreateKnowledgeObjectRequestContentSemanticModelRelationships) *CreateKnowledgeObjectRequestContentSemanticModel {
	s.Relationships = v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModel) Validate() error {
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

type CreateKnowledgeObjectRequestContentSemanticModelDatasets struct {
	// The document description.
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
	Fields []*CreateKnowledgeObjectRequestContentSemanticModelDatasetsFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The document name. Existing objects are not overwritten based on the name.
	//
	// example:
	//
	// Business description
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The explicitly declared primary key field name. This parameter is optional. Each item must exist in the Fields of this dataset.
	PrimaryKey []*string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty" type:"Repeated"`
}

func (s CreateKnowledgeObjectRequestContentSemanticModelDatasets) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectRequestContentSemanticModelDatasets) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasets) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasets) GetEntityId() *string {
	return s.EntityId
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasets) GetFields() []*CreateKnowledgeObjectRequestContentSemanticModelDatasetsFields {
	return s.Fields
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasets) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasets) GetPrimaryKey() []*string {
	return s.PrimaryKey
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasets) SetDescription(v string) *CreateKnowledgeObjectRequestContentSemanticModelDatasets {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasets) SetEntityId(v string) *CreateKnowledgeObjectRequestContentSemanticModelDatasets {
	s.EntityId = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasets) SetFields(v []*CreateKnowledgeObjectRequestContentSemanticModelDatasetsFields) *CreateKnowledgeObjectRequestContentSemanticModelDatasets {
	s.Fields = v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasets) SetName(v string) *CreateKnowledgeObjectRequestContentSemanticModelDatasets {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasets) SetPrimaryKey(v []*string) *CreateKnowledgeObjectRequestContentSemanticModelDatasets {
	s.PrimaryKey = v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasets) Validate() error {
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

type CreateKnowledgeObjectRequestContentSemanticModelDatasetsFields struct {
	// The data type of the field. This must be provided for publicly available model fields and is not guessed by the server.
	//
	// example:
	//
	// BIGINT
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The document description.
	//
	// example:
	//
	// Sales metrics and policy document
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The document name. Existing objects are not overwritten based on the name.
	//
	// example:
	//
	// Business description
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateKnowledgeObjectRequestContentSemanticModelDatasetsFields) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectRequestContentSemanticModelDatasetsFields) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasetsFields) GetDataType() *string {
	return s.DataType
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasetsFields) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasetsFields) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasetsFields) SetDataType(v string) *CreateKnowledgeObjectRequestContentSemanticModelDatasetsFields {
	s.DataType = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasetsFields) SetDescription(v string) *CreateKnowledgeObjectRequestContentSemanticModelDatasetsFields {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasetsFields) SetName(v string) *CreateKnowledgeObjectRequestContentSemanticModelDatasetsFields {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelDatasetsFields) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeObjectRequestContentSemanticModelMetrics struct {
	// The SemanticDataset.Name to which the metric belongs. It must be locatable in the Datasets of this model.
	//
	// example:
	//
	// orders
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The document description.
	//
	// example:
	//
	// Sales metrics and policy document
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The executable expression. It must include at least one dialect version and is not auto-completed by the model.
	Expression *CreateKnowledgeObjectRequestContentSemanticModelMetricsExpression `json:"Expression,omitempty" xml:"Expression,omitempty" type:"Struct"`
	// The document name. Existing objects are not overwritten based on the name.
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

func (s CreateKnowledgeObjectRequestContentSemanticModelMetrics) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectRequestContentSemanticModelMetrics) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetrics) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetrics) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetrics) GetExpression() *CreateKnowledgeObjectRequestContentSemanticModelMetricsExpression {
	return s.Expression
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetrics) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetrics) GetUnit() *string {
	return s.Unit
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetrics) SetDatasetName(v string) *CreateKnowledgeObjectRequestContentSemanticModelMetrics {
	s.DatasetName = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetrics) SetDescription(v string) *CreateKnowledgeObjectRequestContentSemanticModelMetrics {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetrics) SetExpression(v *CreateKnowledgeObjectRequestContentSemanticModelMetricsExpression) *CreateKnowledgeObjectRequestContentSemanticModelMetrics {
	s.Expression = v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetrics) SetName(v string) *CreateKnowledgeObjectRequestContentSemanticModelMetrics {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetrics) SetUnit(v string) *CreateKnowledgeObjectRequestContentSemanticModelMetrics {
	s.Unit = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetrics) Validate() error {
	if s.Expression != nil {
		if err := s.Expression.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKnowledgeObjectRequestContentSemanticModelMetricsExpression struct {
	// The expressions of the same metric in different SQL dialects. At least one item is required. The same dialect is not repeated.
	Dialects []*CreateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects `json:"Dialects,omitempty" xml:"Dialects,omitempty" type:"Repeated"`
}

func (s CreateKnowledgeObjectRequestContentSemanticModelMetricsExpression) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectRequestContentSemanticModelMetricsExpression) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetricsExpression) GetDialects() []*CreateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects {
	return s.Dialects
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetricsExpression) SetDialects(v []*CreateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects) *CreateKnowledgeObjectRequestContentSemanticModelMetricsExpression {
	s.Dialects = v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetricsExpression) Validate() error {
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

type CreateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects struct {
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

func (s CreateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects) GetDialect() *string {
	return s.Dialect
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects) GetSql() *string {
	return s.Sql
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects) SetDialect(v string) *CreateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects {
	s.Dialect = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects) SetSql(v string) *CreateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects {
	s.Sql = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelMetricsExpressionDialects) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeObjectRequestContentSemanticModelRelationships struct {
	// The SemanticDataset.Name of the starting point.
	//
	// example:
	//
	// orders
	FromDataset *string `json:"FromDataset,omitempty" xml:"FromDataset,omitempty"`
	// The equi-join pairs of fields at both ends of the relationship. For composite relationships, fields are paired item by item.
	JoinKeys []*CreateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys `json:"JoinKeys,omitempty" xml:"JoinKeys,omitempty" type:"Repeated"`
	// The document name. Existing objects are not overwritten based on the name.
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

func (s CreateKnowledgeObjectRequestContentSemanticModelRelationships) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectRequestContentSemanticModelRelationships) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelRelationships) GetFromDataset() *string {
	return s.FromDataset
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelRelationships) GetJoinKeys() []*CreateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys {
	return s.JoinKeys
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelRelationships) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelRelationships) GetToDataset() *string {
	return s.ToDataset
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelRelationships) SetFromDataset(v string) *CreateKnowledgeObjectRequestContentSemanticModelRelationships {
	s.FromDataset = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelRelationships) SetJoinKeys(v []*CreateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys) *CreateKnowledgeObjectRequestContentSemanticModelRelationships {
	s.JoinKeys = v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelRelationships) SetName(v string) *CreateKnowledgeObjectRequestContentSemanticModelRelationships {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelRelationships) SetToDataset(v string) *CreateKnowledgeObjectRequestContentSemanticModelRelationships {
	s.ToDataset = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelRelationships) Validate() error {
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

type CreateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys struct {
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

func (s CreateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys) GetFromField() *string {
	return s.FromField
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys) GetToField() *string {
	return s.ToField
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys) SetFromField(v string) *CreateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys {
	s.FromField = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys) SetToField(v string) *CreateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys {
	s.ToField = &v
	return s
}

func (s *CreateKnowledgeObjectRequestContentSemanticModelRelationshipsJoinKeys) Validate() error {
	return dara.Validate(s)
}
