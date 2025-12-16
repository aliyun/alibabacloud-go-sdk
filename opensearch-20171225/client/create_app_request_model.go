// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSwitch(v bool) *CreateAppRequest
	GetAutoSwitch() *bool
	SetCluster(v *CreateAppRequestCluster) *CreateAppRequest
	GetCluster() *CreateAppRequestCluster
	SetConfigItems(v []map[string]interface{}) *CreateAppRequest
	GetConfigItems() []map[string]interface{}
	SetDataSources(v []*CreateAppRequestDataSources) *CreateAppRequest
	GetDataSources() []*CreateAppRequestDataSources
	SetDescription(v string) *CreateAppRequest
	GetDescription() *string
	SetDomain(v *CreateAppRequestDomain) *CreateAppRequest
	GetDomain() *CreateAppRequestDomain
	SetFetchFields(v []*string) *CreateAppRequest
	GetFetchFields() []*string
	SetFirstRanks(v []*CreateAppRequestFirstRanks) *CreateAppRequest
	GetFirstRanks() []*CreateAppRequestFirstRanks
	SetInterpretations(v []map[string]interface{}) *CreateAppRequest
	GetInterpretations() []map[string]interface{}
	SetNetworkType(v string) *CreateAppRequest
	GetNetworkType() *string
	SetPrompts(v []map[string]interface{}) *CreateAppRequest
	GetPrompts() []map[string]interface{}
	SetQueryProcessors(v []*CreateAppRequestQueryProcessors) *CreateAppRequest
	GetQueryProcessors() []*CreateAppRequestQueryProcessors
	SetRealtimeShared(v bool) *CreateAppRequest
	GetRealtimeShared() *bool
	SetSchema(v *CreateAppRequestSchema) *CreateAppRequest
	GetSchema() *CreateAppRequestSchema
	SetSchemas(v []*CreateAppRequestSchemas) *CreateAppRequest
	GetSchemas() []*CreateAppRequestSchemas
	SetSecondRanks(v []*CreateAppRequestSecondRanks) *CreateAppRequest
	GetSecondRanks() []*CreateAppRequestSecondRanks
	SetSummaries(v []*CreateAppRequestSummaries) *CreateAppRequest
	GetSummaries() []*CreateAppRequestSummaries
	SetDryRun(v bool) *CreateAppRequest
	GetDryRun() *bool
}

type CreateAppRequest struct {
	// Specifies whether to automatically switch the created version to an online version. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	AutoSwitch *bool `json:"autoSwitch,omitempty" xml:"autoSwitch,omitempty"`
	// The capability opening configurations.
	Cluster     *CreateAppRequestCluster `json:"cluster,omitempty" xml:"cluster,omitempty" type:"Struct"`
	ConfigItems []map[string]interface{} `json:"configItems,omitempty" xml:"configItems,omitempty" type:"Repeated"`
	// The configurations of data sources.
	DataSources []*CreateAppRequestDataSources `json:"dataSources,omitempty" xml:"dataSources,omitempty" type:"Repeated"`
	// The version description.
	//
	// example:
	//
	// "Version description"
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The industry model module.
	Domain *CreateAppRequestDomain `json:"domain,omitempty" xml:"domain,omitempty" type:"Struct"`
	// The default display fields.
	FetchFields []*string `json:"fetchFields,omitempty" xml:"fetchFields,omitempty" type:"Repeated"`
	// The configurations of rough sort.
	FirstRanks      []*CreateAppRequestFirstRanks `json:"firstRanks,omitempty" xml:"firstRanks,omitempty" type:"Repeated"`
	Interpretations []map[string]interface{}      `json:"interpretations,omitempty" xml:"interpretations,omitempty" type:"Repeated"`
	// The zone identifier. Valid values:
	//
	// 	- vpc
	//
	// 	- oxs
	//
	// example:
	//
	// vpc
	NetworkType *string                  `json:"networkType,omitempty" xml:"networkType,omitempty"`
	Prompts     []map[string]interface{} `json:"prompts,omitempty" xml:"prompts,omitempty" type:"Repeated"`
	// The query intent understanding configurations.
	QueryProcessors []*CreateAppRequestQueryProcessors `json:"queryProcessors,omitempty" xml:"queryProcessors,omitempty" type:"Repeated"`
	RealtimeShared  *bool                              `json:"realtimeShared,omitempty" xml:"realtimeShared,omitempty"`
	// The single-table schema.
	Schema *CreateAppRequestSchema `json:"schema,omitempty" xml:"schema,omitempty" type:"Struct"`
	// The multi-table schema.
	Schemas []*CreateAppRequestSchemas `json:"schemas,omitempty" xml:"schemas,omitempty" type:"Repeated"`
	// The configurations of fine sort.
	SecondRanks []*CreateAppRequestSecondRanks `json:"secondRanks,omitempty" xml:"secondRanks,omitempty" type:"Repeated"`
	// The summary configurations of search results.
	Summaries []*CreateAppRequestSummaries `json:"summaries,omitempty" xml:"summaries,omitempty" type:"Repeated"`
	// Specifies whether to perform a dry run. This parameter is only used to check whether the data source is valid. Valid values: true and false.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) GetAutoSwitch() *bool {
	return s.AutoSwitch
}

func (s *CreateAppRequest) GetCluster() *CreateAppRequestCluster {
	return s.Cluster
}

func (s *CreateAppRequest) GetConfigItems() []map[string]interface{} {
	return s.ConfigItems
}

func (s *CreateAppRequest) GetDataSources() []*CreateAppRequestDataSources {
	return s.DataSources
}

func (s *CreateAppRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAppRequest) GetDomain() *CreateAppRequestDomain {
	return s.Domain
}

func (s *CreateAppRequest) GetFetchFields() []*string {
	return s.FetchFields
}

func (s *CreateAppRequest) GetFirstRanks() []*CreateAppRequestFirstRanks {
	return s.FirstRanks
}

func (s *CreateAppRequest) GetInterpretations() []map[string]interface{} {
	return s.Interpretations
}

func (s *CreateAppRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateAppRequest) GetPrompts() []map[string]interface{} {
	return s.Prompts
}

func (s *CreateAppRequest) GetQueryProcessors() []*CreateAppRequestQueryProcessors {
	return s.QueryProcessors
}

func (s *CreateAppRequest) GetRealtimeShared() *bool {
	return s.RealtimeShared
}

func (s *CreateAppRequest) GetSchema() *CreateAppRequestSchema {
	return s.Schema
}

func (s *CreateAppRequest) GetSchemas() []*CreateAppRequestSchemas {
	return s.Schemas
}

func (s *CreateAppRequest) GetSecondRanks() []*CreateAppRequestSecondRanks {
	return s.SecondRanks
}

func (s *CreateAppRequest) GetSummaries() []*CreateAppRequestSummaries {
	return s.Summaries
}

func (s *CreateAppRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateAppRequest) SetAutoSwitch(v bool) *CreateAppRequest {
	s.AutoSwitch = &v
	return s
}

func (s *CreateAppRequest) SetCluster(v *CreateAppRequestCluster) *CreateAppRequest {
	s.Cluster = v
	return s
}

func (s *CreateAppRequest) SetConfigItems(v []map[string]interface{}) *CreateAppRequest {
	s.ConfigItems = v
	return s
}

func (s *CreateAppRequest) SetDataSources(v []*CreateAppRequestDataSources) *CreateAppRequest {
	s.DataSources = v
	return s
}

func (s *CreateAppRequest) SetDescription(v string) *CreateAppRequest {
	s.Description = &v
	return s
}

func (s *CreateAppRequest) SetDomain(v *CreateAppRequestDomain) *CreateAppRequest {
	s.Domain = v
	return s
}

func (s *CreateAppRequest) SetFetchFields(v []*string) *CreateAppRequest {
	s.FetchFields = v
	return s
}

func (s *CreateAppRequest) SetFirstRanks(v []*CreateAppRequestFirstRanks) *CreateAppRequest {
	s.FirstRanks = v
	return s
}

func (s *CreateAppRequest) SetInterpretations(v []map[string]interface{}) *CreateAppRequest {
	s.Interpretations = v
	return s
}

func (s *CreateAppRequest) SetNetworkType(v string) *CreateAppRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateAppRequest) SetPrompts(v []map[string]interface{}) *CreateAppRequest {
	s.Prompts = v
	return s
}

func (s *CreateAppRequest) SetQueryProcessors(v []*CreateAppRequestQueryProcessors) *CreateAppRequest {
	s.QueryProcessors = v
	return s
}

func (s *CreateAppRequest) SetRealtimeShared(v bool) *CreateAppRequest {
	s.RealtimeShared = &v
	return s
}

func (s *CreateAppRequest) SetSchema(v *CreateAppRequestSchema) *CreateAppRequest {
	s.Schema = v
	return s
}

func (s *CreateAppRequest) SetSchemas(v []*CreateAppRequestSchemas) *CreateAppRequest {
	s.Schemas = v
	return s
}

func (s *CreateAppRequest) SetSecondRanks(v []*CreateAppRequestSecondRanks) *CreateAppRequest {
	s.SecondRanks = v
	return s
}

func (s *CreateAppRequest) SetSummaries(v []*CreateAppRequestSummaries) *CreateAppRequest {
	s.Summaries = v
	return s
}

func (s *CreateAppRequest) SetDryRun(v bool) *CreateAppRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAppRequest) Validate() error {
	if s.Cluster != nil {
		if err := s.Cluster.Validate(); err != nil {
			return err
		}
	}
	if s.DataSources != nil {
		for _, item := range s.DataSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Domain != nil {
		if err := s.Domain.Validate(); err != nil {
			return err
		}
	}
	if s.FirstRanks != nil {
		for _, item := range s.FirstRanks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QueryProcessors != nil {
		for _, item := range s.QueryProcessors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Schema != nil {
		if err := s.Schema.Validate(); err != nil {
			return err
		}
	}
	if s.Schemas != nil {
		for _, item := range s.Schemas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SecondRanks != nil {
		for _, item := range s.SecondRanks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Summaries != nil {
		for _, item := range s.Summaries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAppRequestCluster struct {
	ChunkModels                  []map[string]interface{} `json:"chunkModels,omitempty" xml:"chunkModels,omitempty" type:"Repeated"`
	GraphRag                     map[string]interface{}   `json:"graphRag,omitempty" xml:"graphRag,omitempty"`
	ImageContentRecognizerModels []map[string]interface{} `json:"imageContentRecognizerModels,omitempty" xml:"imageContentRecognizerModels,omitempty" type:"Repeated"`
	// The maximum length of the query clause.
	//
	// example:
	//
	// 1024
	MaxQueryClauseLength *int32 `json:"maxQueryClauseLength,omitempty" xml:"maxQueryClauseLength,omitempty"`
	// The timeout period. Unit: milliseconds.
	//
	// example:
	//
	// 750
	MaxTimeoutMS *int32 `json:"maxTimeoutMS,omitempty" xml:"maxTimeoutMS,omitempty"`
	// example:
	//
	// ops-text-embedding-002
	TextEmbeddingModel *string `json:"textEmbeddingModel,omitempty" xml:"textEmbeddingModel,omitempty"`
	// example:
	//
	// ops-text-sparse-embedding-001
	TextSparseEmbeddingModel *string                  `json:"textSparseEmbeddingModel,omitempty" xml:"textSparseEmbeddingModel,omitempty"`
	VectorIndexConfigs       []map[string]interface{} `json:"vectorIndexConfigs,omitempty" xml:"vectorIndexConfigs,omitempty" type:"Repeated"`
}

func (s CreateAppRequestCluster) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestCluster) GoString() string {
	return s.String()
}

func (s *CreateAppRequestCluster) GetChunkModels() []map[string]interface{} {
	return s.ChunkModels
}

func (s *CreateAppRequestCluster) GetGraphRag() map[string]interface{} {
	return s.GraphRag
}

func (s *CreateAppRequestCluster) GetImageContentRecognizerModels() []map[string]interface{} {
	return s.ImageContentRecognizerModels
}

func (s *CreateAppRequestCluster) GetMaxQueryClauseLength() *int32 {
	return s.MaxQueryClauseLength
}

func (s *CreateAppRequestCluster) GetMaxTimeoutMS() *int32 {
	return s.MaxTimeoutMS
}

func (s *CreateAppRequestCluster) GetTextEmbeddingModel() *string {
	return s.TextEmbeddingModel
}

func (s *CreateAppRequestCluster) GetTextSparseEmbeddingModel() *string {
	return s.TextSparseEmbeddingModel
}

func (s *CreateAppRequestCluster) GetVectorIndexConfigs() []map[string]interface{} {
	return s.VectorIndexConfigs
}

func (s *CreateAppRequestCluster) SetChunkModels(v []map[string]interface{}) *CreateAppRequestCluster {
	s.ChunkModels = v
	return s
}

func (s *CreateAppRequestCluster) SetGraphRag(v map[string]interface{}) *CreateAppRequestCluster {
	s.GraphRag = v
	return s
}

func (s *CreateAppRequestCluster) SetImageContentRecognizerModels(v []map[string]interface{}) *CreateAppRequestCluster {
	s.ImageContentRecognizerModels = v
	return s
}

func (s *CreateAppRequestCluster) SetMaxQueryClauseLength(v int32) *CreateAppRequestCluster {
	s.MaxQueryClauseLength = &v
	return s
}

func (s *CreateAppRequestCluster) SetMaxTimeoutMS(v int32) *CreateAppRequestCluster {
	s.MaxTimeoutMS = &v
	return s
}

func (s *CreateAppRequestCluster) SetTextEmbeddingModel(v string) *CreateAppRequestCluster {
	s.TextEmbeddingModel = &v
	return s
}

func (s *CreateAppRequestCluster) SetTextSparseEmbeddingModel(v string) *CreateAppRequestCluster {
	s.TextSparseEmbeddingModel = &v
	return s
}

func (s *CreateAppRequestCluster) SetVectorIndexConfigs(v []map[string]interface{}) *CreateAppRequestCluster {
	s.VectorIndexConfigs = v
	return s
}

func (s *CreateAppRequestCluster) Validate() error {
	return dara.Validate(s)
}

type CreateAppRequestDataSources struct {
	// The information about field mappings.
	Fields []map[string]interface{} `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// The primary key.
	//
	// example:
	//
	// id
	KeyField *string `json:"keyField,omitempty" xml:"keyField,omitempty"`
	// The information about the data source.
	//
	// example:
	//
	// {
	//
	//   "instanceId": "rds-instance-id",
	//
	//   "dbName": "my_db",
	//
	//   "dbTableName": "my_table",
	//
	//   "dbUser": "my",
	//
	//   "dbPassword": "my_passwd",
	//
	//   "filter":"",
	//
	//   "autoSync": true
	//
	// }
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The plug-ins that are used for data processing.
	//
	// name:
	//
	// 	- JsonKeyValueExtractor
	//
	// 	- MultiValueSpliter
	//
	// 	- KeyValueExtractor
	//
	// 	- StringCatenateExtractor
	//
	// 	- HTMLTagRemover
	//
	// parameters:
	//
	// 	- JsonKeyValueExtractor
	//
	// 	- MultiValueSpliter
	//
	// 	- KeyValueExtractor
	//
	// 	- StringCatenateExtractor
	//
	// 	- HTMLTagRemover
	//
	// example:
	//
	// {
	//
	//   "name": "JsonKeyValueExtractor",
	//
	//   "parameters": {
	//
	//   "key": "my_field"
	//
	// }
	//
	// }
	Plugins map[string]interface{} `json:"plugins,omitempty" xml:"plugins,omitempty"`
	// The name of the wide table.
	//
	// example:
	//
	// table_name
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	// The name of the table in the application.
	//
	// example:
	//
	// main
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- rds
	//
	// 	- odps
	//
	// 	- opensearch
	//
	// 	- polardb
	//
	// example:
	//
	// rds
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAppRequestDataSources) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestDataSources) GoString() string {
	return s.String()
}

func (s *CreateAppRequestDataSources) GetFields() []map[string]interface{} {
	return s.Fields
}

func (s *CreateAppRequestDataSources) GetKeyField() *string {
	return s.KeyField
}

func (s *CreateAppRequestDataSources) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateAppRequestDataSources) GetPlugins() map[string]interface{} {
	return s.Plugins
}

func (s *CreateAppRequestDataSources) GetSchemaName() *string {
	return s.SchemaName
}

func (s *CreateAppRequestDataSources) GetTableName() *string {
	return s.TableName
}

func (s *CreateAppRequestDataSources) GetType() *string {
	return s.Type
}

func (s *CreateAppRequestDataSources) SetFields(v []map[string]interface{}) *CreateAppRequestDataSources {
	s.Fields = v
	return s
}

func (s *CreateAppRequestDataSources) SetKeyField(v string) *CreateAppRequestDataSources {
	s.KeyField = &v
	return s
}

func (s *CreateAppRequestDataSources) SetParameters(v map[string]interface{}) *CreateAppRequestDataSources {
	s.Parameters = v
	return s
}

func (s *CreateAppRequestDataSources) SetPlugins(v map[string]interface{}) *CreateAppRequestDataSources {
	s.Plugins = v
	return s
}

func (s *CreateAppRequestDataSources) SetSchemaName(v string) *CreateAppRequestDataSources {
	s.SchemaName = &v
	return s
}

func (s *CreateAppRequestDataSources) SetTableName(v string) *CreateAppRequestDataSources {
	s.TableName = &v
	return s
}

func (s *CreateAppRequestDataSources) SetType(v string) *CreateAppRequestDataSources {
	s.Type = &v
	return s
}

func (s *CreateAppRequestDataSources) Validate() error {
	return dara.Validate(s)
}

type CreateAppRequestDomain struct {
	// The industry category.
	//
	// example:
	//
	// general
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The selected feature category. Valid values:
	//
	// 	- qp: query analysis
	//
	// 	- algo: sort policy
	//
	// 	- service: service
	//
	// example:
	//
	// {"qp":["spellcheck"],"algo":["pop"],"service":["suggest"]}
	Functions map[string]interface{} `json:"functions,omitempty" xml:"functions,omitempty"`
	// The industry type.
	//
	// example:
	//
	// ecommerce
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateAppRequestDomain) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestDomain) GoString() string {
	return s.String()
}

func (s *CreateAppRequestDomain) GetCategory() *string {
	return s.Category
}

func (s *CreateAppRequestDomain) GetFunctions() map[string]interface{} {
	return s.Functions
}

func (s *CreateAppRequestDomain) GetName() *string {
	return s.Name
}

func (s *CreateAppRequestDomain) SetCategory(v string) *CreateAppRequestDomain {
	s.Category = &v
	return s
}

func (s *CreateAppRequestDomain) SetFunctions(v map[string]interface{}) *CreateAppRequestDomain {
	s.Functions = v
	return s
}

func (s *CreateAppRequestDomain) SetName(v string) *CreateAppRequestDomain {
	s.Name = &v
	return s
}

func (s *CreateAppRequestDomain) Validate() error {
	return dara.Validate(s)
}

type CreateAppRequestFirstRanks struct {
	// Specifies whether the expression is the default one.
	//
	// example:
	//
	// False
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The description.
	//
	// example:
	//
	// Description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The information about the expression. The information can be of the array or string type.
	//
	// example:
	//
	// String :"random()*100+now()";
	//
	// Array: [
	//
	//     {
	//
	//       "attribute": "static_bm25()",
	//
	//       "arg": "",
	//
	//       "weight": 10
	//
	//     }
	//
	//   ]
	Meta interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	// The name of the rough sort expression.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The expression type. Valid values:
	//
	// 	- STRUCT: The content of the expression is a structure.
	//
	// 	- STRING (default): You can configure a custom formula.
	//
	// example:
	//
	// STRING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAppRequestFirstRanks) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestFirstRanks) GoString() string {
	return s.String()
}

func (s *CreateAppRequestFirstRanks) GetActive() *bool {
	return s.Active
}

func (s *CreateAppRequestFirstRanks) GetDescription() *string {
	return s.Description
}

func (s *CreateAppRequestFirstRanks) GetMeta() interface{} {
	return s.Meta
}

func (s *CreateAppRequestFirstRanks) GetName() *string {
	return s.Name
}

func (s *CreateAppRequestFirstRanks) GetType() *string {
	return s.Type
}

func (s *CreateAppRequestFirstRanks) SetActive(v bool) *CreateAppRequestFirstRanks {
	s.Active = &v
	return s
}

func (s *CreateAppRequestFirstRanks) SetDescription(v string) *CreateAppRequestFirstRanks {
	s.Description = &v
	return s
}

func (s *CreateAppRequestFirstRanks) SetMeta(v interface{}) *CreateAppRequestFirstRanks {
	s.Meta = v
	return s
}

func (s *CreateAppRequestFirstRanks) SetName(v string) *CreateAppRequestFirstRanks {
	s.Name = &v
	return s
}

func (s *CreateAppRequestFirstRanks) SetType(v string) *CreateAppRequestFirstRanks {
	s.Type = &v
	return s
}

func (s *CreateAppRequestFirstRanks) Validate() error {
	return dara.Validate(s)
}

type CreateAppRequestQueryProcessors struct {
	// Specifies whether the rule is the default one.
	//
	// example:
	//
	// True
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The industry category.
	//
	// example:
	//
	// ""
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The industry type. Valid values:
	//
	// 	- GENERAL
	//
	// 	- ECOMMERCE
	//
	// 	- IT_CONTENT
	//
	// example:
	//
	// ECOMMERCE
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The index range.
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The rule name.
	//
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The features.
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s CreateAppRequestQueryProcessors) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestQueryProcessors) GoString() string {
	return s.String()
}

func (s *CreateAppRequestQueryProcessors) GetActive() *bool {
	return s.Active
}

func (s *CreateAppRequestQueryProcessors) GetCategory() *string {
	return s.Category
}

func (s *CreateAppRequestQueryProcessors) GetDomain() *string {
	return s.Domain
}

func (s *CreateAppRequestQueryProcessors) GetIndexes() []*string {
	return s.Indexes
}

func (s *CreateAppRequestQueryProcessors) GetName() *string {
	return s.Name
}

func (s *CreateAppRequestQueryProcessors) GetProcessors() []map[string]interface{} {
	return s.Processors
}

func (s *CreateAppRequestQueryProcessors) SetActive(v bool) *CreateAppRequestQueryProcessors {
	s.Active = &v
	return s
}

func (s *CreateAppRequestQueryProcessors) SetCategory(v string) *CreateAppRequestQueryProcessors {
	s.Category = &v
	return s
}

func (s *CreateAppRequestQueryProcessors) SetDomain(v string) *CreateAppRequestQueryProcessors {
	s.Domain = &v
	return s
}

func (s *CreateAppRequestQueryProcessors) SetIndexes(v []*string) *CreateAppRequestQueryProcessors {
	s.Indexes = v
	return s
}

func (s *CreateAppRequestQueryProcessors) SetName(v string) *CreateAppRequestQueryProcessors {
	s.Name = &v
	return s
}

func (s *CreateAppRequestQueryProcessors) SetProcessors(v []map[string]interface{}) *CreateAppRequestQueryProcessors {
	s.Processors = v
	return s
}

func (s *CreateAppRequestQueryProcessors) Validate() error {
	return dara.Validate(s)
}

type CreateAppRequestSchema struct {
	// The sort configurations.
	IndexSortConfig []*CreateAppRequestSchemaIndexSortConfig `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	// The index schema.
	Indexes *CreateAppRequestSchemaIndexes `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
	// The name of the wide table.
	//
	// example:
	//
	// table_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the level-1 routing field.
	//
	// example:
	//
	// field1
	RouteField *string `json:"routeField,omitempty" xml:"routeField,omitempty"`
	// The hot values of the level-1 routing field. After you configure this parameter, level-2 routing is enabled.
	RouteFieldValues []*string `json:"routeFieldValues,omitempty" xml:"routeFieldValues,omitempty" type:"Repeated"`
	// The name of the level-2 routing field. This parameter takes effect only when the `routeFieldValues` parameter is configured. By default, the wide-table primary key field is used as the level-2 routing field.
	//
	// example:
	//
	// field2
	SecondRouteField *string `json:"secondRouteField,omitempty" xml:"secondRouteField,omitempty"`
	// The table schema.
	//
	// example:
	//
	// {
	//
	// 	"primaryTable": true,
	//
	// 	"name": "main",
	//
	// 	"fields": {
	//
	// 		"id": {
	//
	// 			"name": "id",
	//
	// 			"type": "LITERAL",
	//
	// 			"primaryKey": true
	//
	// 		},
	//
	// 		"title": {
	//
	// 			"name": "title",
	//
	// 			"type": "TEXT",
	//
	// 			"primaryKey": false
	//
	// 		},
	//
	// 		"buy": {
	//
	// 			"name": "buy",
	//
	// 			"type": "INT",
	//
	// 			"primaryKey": false
	//
	// 		},
	//
	// 		"cate_id": {
	//
	// 			"name": "cate_id",
	//
	// 			"type": "INT",
	//
	// 			"primaryKey": false
	//
	// 		},
	//
	// 		"cate_name": {
	//
	// 			"name": "cate_name",
	//
	// 			"type": "LITERAL",
	//
	// 			"primaryKey": false
	//
	// 		}
	//
	// 	}
	//
	// }
	Tables map[string]interface{} `json:"tables,omitempty" xml:"tables,omitempty"`
	// The document clearing configurations.
	TtlField *CreateAppRequestSchemaTtlField `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s CreateAppRequestSchema) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestSchema) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSchema) GetIndexSortConfig() []*CreateAppRequestSchemaIndexSortConfig {
	return s.IndexSortConfig
}

func (s *CreateAppRequestSchema) GetIndexes() *CreateAppRequestSchemaIndexes {
	return s.Indexes
}

func (s *CreateAppRequestSchema) GetName() *string {
	return s.Name
}

func (s *CreateAppRequestSchema) GetRouteField() *string {
	return s.RouteField
}

func (s *CreateAppRequestSchema) GetRouteFieldValues() []*string {
	return s.RouteFieldValues
}

func (s *CreateAppRequestSchema) GetSecondRouteField() *string {
	return s.SecondRouteField
}

func (s *CreateAppRequestSchema) GetTables() map[string]interface{} {
	return s.Tables
}

func (s *CreateAppRequestSchema) GetTtlField() *CreateAppRequestSchemaTtlField {
	return s.TtlField
}

func (s *CreateAppRequestSchema) SetIndexSortConfig(v []*CreateAppRequestSchemaIndexSortConfig) *CreateAppRequestSchema {
	s.IndexSortConfig = v
	return s
}

func (s *CreateAppRequestSchema) SetIndexes(v *CreateAppRequestSchemaIndexes) *CreateAppRequestSchema {
	s.Indexes = v
	return s
}

func (s *CreateAppRequestSchema) SetName(v string) *CreateAppRequestSchema {
	s.Name = &v
	return s
}

func (s *CreateAppRequestSchema) SetRouteField(v string) *CreateAppRequestSchema {
	s.RouteField = &v
	return s
}

func (s *CreateAppRequestSchema) SetRouteFieldValues(v []*string) *CreateAppRequestSchema {
	s.RouteFieldValues = v
	return s
}

func (s *CreateAppRequestSchema) SetSecondRouteField(v string) *CreateAppRequestSchema {
	s.SecondRouteField = &v
	return s
}

func (s *CreateAppRequestSchema) SetTables(v map[string]interface{}) *CreateAppRequestSchema {
	s.Tables = v
	return s
}

func (s *CreateAppRequestSchema) SetTtlField(v *CreateAppRequestSchemaTtlField) *CreateAppRequestSchema {
	s.TtlField = v
	return s
}

func (s *CreateAppRequestSchema) Validate() error {
	if s.IndexSortConfig != nil {
		for _, item := range s.IndexSortConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Indexes != nil {
		if err := s.Indexes.Validate(); err != nil {
			return err
		}
	}
	if s.TtlField != nil {
		if err := s.TtlField.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppRequestSchemaIndexSortConfig struct {
	// The sort method.
	//
	// example:
	//
	// ASC;
	//
	// DESC;
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// The sort field.
	//
	// example:
	//
	// field1
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
}

func (s CreateAppRequestSchemaIndexSortConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestSchemaIndexSortConfig) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSchemaIndexSortConfig) GetDirection() *string {
	return s.Direction
}

func (s *CreateAppRequestSchemaIndexSortConfig) GetField() *string {
	return s.Field
}

func (s *CreateAppRequestSchemaIndexSortConfig) SetDirection(v string) *CreateAppRequestSchemaIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *CreateAppRequestSchemaIndexSortConfig) SetField(v string) *CreateAppRequestSchemaIndexSortConfig {
	s.Field = &v
	return s
}

func (s *CreateAppRequestSchemaIndexSortConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAppRequestSchemaIndexes struct {
	// The attribute fields.
	FilterFields []*string `json:"filterFields,omitempty" xml:"filterFields,omitempty" type:"Repeated"`
	// The index fields.
	//
	// example:
	//
	// {
	//
	//   "fields": ["title"],
	//
	//   "analyzer": "chn_standard"
	//
	// }
	SearchFields map[string]interface{} `json:"searchFields,omitempty" xml:"searchFields,omitempty"`
}

func (s CreateAppRequestSchemaIndexes) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestSchemaIndexes) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSchemaIndexes) GetFilterFields() []*string {
	return s.FilterFields
}

func (s *CreateAppRequestSchemaIndexes) GetSearchFields() map[string]interface{} {
	return s.SearchFields
}

func (s *CreateAppRequestSchemaIndexes) SetFilterFields(v []*string) *CreateAppRequestSchemaIndexes {
	s.FilterFields = v
	return s
}

func (s *CreateAppRequestSchemaIndexes) SetSearchFields(v map[string]interface{}) *CreateAppRequestSchemaIndexes {
	s.SearchFields = v
	return s
}

func (s *CreateAppRequestSchemaIndexes) Validate() error {
	return dara.Validate(s)
}

type CreateAppRequestSchemaTtlField struct {
	// The name of the document time field.
	//
	// example:
	//
	// text1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The TTL. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	Ttl *int64 `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s CreateAppRequestSchemaTtlField) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestSchemaTtlField) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSchemaTtlField) GetName() *string {
	return s.Name
}

func (s *CreateAppRequestSchemaTtlField) GetTtl() *int64 {
	return s.Ttl
}

func (s *CreateAppRequestSchemaTtlField) SetName(v string) *CreateAppRequestSchemaTtlField {
	s.Name = &v
	return s
}

func (s *CreateAppRequestSchemaTtlField) SetTtl(v int64) *CreateAppRequestSchemaTtlField {
	s.Ttl = &v
	return s
}

func (s *CreateAppRequestSchemaTtlField) Validate() error {
	return dara.Validate(s)
}

type CreateAppRequestSchemas struct {
	// The sort configurations.
	IndexSortConfig []*CreateAppRequestSchemasIndexSortConfig `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	// The index schema.
	Indexes *CreateAppRequestSchemasIndexes `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
	// The name of the wide table.
	//
	// example:
	//
	// table_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the level-1 routing field.
	//
	// example:
	//
	// field1
	RouteField *string `json:"routeField,omitempty" xml:"routeField,omitempty"`
	// The hot values of the level-1 routing field. After you configure this parameter, level-2 routing is enabled.
	RouteFieldValues []*string `json:"routeFieldValues,omitempty" xml:"routeFieldValues,omitempty" type:"Repeated"`
	// The name of the level-2 routing field. This parameter takes effect only when the routeFieldValues parameter is configured. By default, the wide-table primary key field is used as the level-2 routing field.
	//
	// example:
	//
	// field2
	SecondRouteField *string `json:"secondRouteField,omitempty" xml:"secondRouteField,omitempty"`
	// The table schema.
	//
	// example:
	//
	// {
	//
	// 	"primaryTable": true,
	//
	// 	"name": "main",
	//
	// 	"fields": {
	//
	// 		"id": {
	//
	// 			"name": "id",
	//
	// 			"type": "LITERAL",
	//
	// 			"primaryKey": true
	//
	// 		},
	//
	// 		"title": {
	//
	// 			"name": "title",
	//
	// 			"type": "TEXT",
	//
	// 			"primaryKey": false
	//
	// 		},
	//
	// 		"buy": {
	//
	// 			"name": "buy",
	//
	// 			"type": "INT",
	//
	// 			"primaryKey": false
	//
	// 		},
	//
	// 		"cate_id": {
	//
	// 			"name": "cate_id",
	//
	// 			"type": "INT",
	//
	// 			"primaryKey": false
	//
	// 		},
	//
	// 		"cate_name": {
	//
	// 			"name": "cate_name",
	//
	// 			"type": "LITERAL",
	//
	// 			"primaryKey": false
	//
	// 		}
	//
	// 	}
	//
	// }
	Tables map[string]interface{} `json:"tables,omitempty" xml:"tables,omitempty"`
	// The document clearing configurations.
	TtlField *CreateAppRequestSchemasTtlField `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s CreateAppRequestSchemas) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestSchemas) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSchemas) GetIndexSortConfig() []*CreateAppRequestSchemasIndexSortConfig {
	return s.IndexSortConfig
}

func (s *CreateAppRequestSchemas) GetIndexes() *CreateAppRequestSchemasIndexes {
	return s.Indexes
}

func (s *CreateAppRequestSchemas) GetName() *string {
	return s.Name
}

func (s *CreateAppRequestSchemas) GetRouteField() *string {
	return s.RouteField
}

func (s *CreateAppRequestSchemas) GetRouteFieldValues() []*string {
	return s.RouteFieldValues
}

func (s *CreateAppRequestSchemas) GetSecondRouteField() *string {
	return s.SecondRouteField
}

func (s *CreateAppRequestSchemas) GetTables() map[string]interface{} {
	return s.Tables
}

func (s *CreateAppRequestSchemas) GetTtlField() *CreateAppRequestSchemasTtlField {
	return s.TtlField
}

func (s *CreateAppRequestSchemas) SetIndexSortConfig(v []*CreateAppRequestSchemasIndexSortConfig) *CreateAppRequestSchemas {
	s.IndexSortConfig = v
	return s
}

func (s *CreateAppRequestSchemas) SetIndexes(v *CreateAppRequestSchemasIndexes) *CreateAppRequestSchemas {
	s.Indexes = v
	return s
}

func (s *CreateAppRequestSchemas) SetName(v string) *CreateAppRequestSchemas {
	s.Name = &v
	return s
}

func (s *CreateAppRequestSchemas) SetRouteField(v string) *CreateAppRequestSchemas {
	s.RouteField = &v
	return s
}

func (s *CreateAppRequestSchemas) SetRouteFieldValues(v []*string) *CreateAppRequestSchemas {
	s.RouteFieldValues = v
	return s
}

func (s *CreateAppRequestSchemas) SetSecondRouteField(v string) *CreateAppRequestSchemas {
	s.SecondRouteField = &v
	return s
}

func (s *CreateAppRequestSchemas) SetTables(v map[string]interface{}) *CreateAppRequestSchemas {
	s.Tables = v
	return s
}

func (s *CreateAppRequestSchemas) SetTtlField(v *CreateAppRequestSchemasTtlField) *CreateAppRequestSchemas {
	s.TtlField = v
	return s
}

func (s *CreateAppRequestSchemas) Validate() error {
	if s.IndexSortConfig != nil {
		for _, item := range s.IndexSortConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Indexes != nil {
		if err := s.Indexes.Validate(); err != nil {
			return err
		}
	}
	if s.TtlField != nil {
		if err := s.TtlField.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppRequestSchemasIndexSortConfig struct {
	// The sort method.
	//
	// 	- ASC
	//
	// 	- DESC
	//
	// example:
	//
	// DESC
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// The sort field.
	//
	// example:
	//
	// field1
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
}

func (s CreateAppRequestSchemasIndexSortConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestSchemasIndexSortConfig) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSchemasIndexSortConfig) GetDirection() *string {
	return s.Direction
}

func (s *CreateAppRequestSchemasIndexSortConfig) GetField() *string {
	return s.Field
}

func (s *CreateAppRequestSchemasIndexSortConfig) SetDirection(v string) *CreateAppRequestSchemasIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *CreateAppRequestSchemasIndexSortConfig) SetField(v string) *CreateAppRequestSchemasIndexSortConfig {
	s.Field = &v
	return s
}

func (s *CreateAppRequestSchemasIndexSortConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAppRequestSchemasIndexes struct {
	// The attribute fields.
	FilterFields []*string `json:"filterFields,omitempty" xml:"filterFields,omitempty" type:"Repeated"`
	// The index fields.
	//
	// example:
	//
	// {
	//
	//   "fields": ["title"],
	//
	//   "analyzer": "chn_standard"
	//
	// }
	SearchFields map[string]interface{} `json:"searchFields,omitempty" xml:"searchFields,omitempty"`
}

func (s CreateAppRequestSchemasIndexes) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestSchemasIndexes) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSchemasIndexes) GetFilterFields() []*string {
	return s.FilterFields
}

func (s *CreateAppRequestSchemasIndexes) GetSearchFields() map[string]interface{} {
	return s.SearchFields
}

func (s *CreateAppRequestSchemasIndexes) SetFilterFields(v []*string) *CreateAppRequestSchemasIndexes {
	s.FilterFields = v
	return s
}

func (s *CreateAppRequestSchemasIndexes) SetSearchFields(v map[string]interface{}) *CreateAppRequestSchemasIndexes {
	s.SearchFields = v
	return s
}

func (s *CreateAppRequestSchemasIndexes) Validate() error {
	return dara.Validate(s)
}

type CreateAppRequestSchemasTtlField struct {
	// The name of the document time field.
	//
	// example:
	//
	// text1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The TTL. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	Ttl *int64 `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s CreateAppRequestSchemasTtlField) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestSchemasTtlField) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSchemasTtlField) GetName() *string {
	return s.Name
}

func (s *CreateAppRequestSchemasTtlField) GetTtl() *int64 {
	return s.Ttl
}

func (s *CreateAppRequestSchemasTtlField) SetName(v string) *CreateAppRequestSchemasTtlField {
	s.Name = &v
	return s
}

func (s *CreateAppRequestSchemasTtlField) SetTtl(v int64) *CreateAppRequestSchemasTtlField {
	s.Ttl = &v
	return s
}

func (s *CreateAppRequestSchemasTtlField) Validate() error {
	return dara.Validate(s)
}

type CreateAppRequestSecondRanks struct {
	// Specifies whether the expression is the default one.
	//
	// example:
	//
	// False
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The description.
	//
	// example:
	//
	// Description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The fine sort expression. You can define an expression that contains fields, feature functions, and mathematical functions to implement complex sort logic.
	//
	// example:
	//
	// "cate_id > 0 and cate_id < 1000"
	Meta interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	// The name of the fine sort expression.
	//
	// example:
	//
	// default
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateAppRequestSecondRanks) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestSecondRanks) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSecondRanks) GetActive() *bool {
	return s.Active
}

func (s *CreateAppRequestSecondRanks) GetDescription() *string {
	return s.Description
}

func (s *CreateAppRequestSecondRanks) GetMeta() interface{} {
	return s.Meta
}

func (s *CreateAppRequestSecondRanks) GetName() *string {
	return s.Name
}

func (s *CreateAppRequestSecondRanks) SetActive(v bool) *CreateAppRequestSecondRanks {
	s.Active = &v
	return s
}

func (s *CreateAppRequestSecondRanks) SetDescription(v string) *CreateAppRequestSecondRanks {
	s.Description = &v
	return s
}

func (s *CreateAppRequestSecondRanks) SetMeta(v interface{}) *CreateAppRequestSecondRanks {
	s.Meta = v
	return s
}

func (s *CreateAppRequestSecondRanks) SetName(v string) *CreateAppRequestSecondRanks {
	s.Name = &v
	return s
}

func (s *CreateAppRequestSecondRanks) Validate() error {
	return dara.Validate(s)
}

type CreateAppRequestSummaries struct {
	// The collection of summary configurations.
	Meta []*CreateAppRequestSummariesMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The group name.
	//
	// example:
	//
	// default
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateAppRequestSummaries) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestSummaries) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSummaries) GetMeta() []*CreateAppRequestSummariesMeta {
	return s.Meta
}

func (s *CreateAppRequestSummaries) GetName() *string {
	return s.Name
}

func (s *CreateAppRequestSummaries) SetMeta(v []*CreateAppRequestSummariesMeta) *CreateAppRequestSummaries {
	s.Meta = v
	return s
}

func (s *CreateAppRequestSummaries) SetName(v string) *CreateAppRequestSummaries {
	s.Name = &v
	return s
}

func (s *CreateAppRequestSummaries) Validate() error {
	if s.Meta != nil {
		for _, item := range s.Meta {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAppRequestSummariesMeta struct {
	// The element that is used for highlighting.
	//
	// example:
	//
	// em
	Element *string `json:"element,omitempty" xml:"element,omitempty"`
	// The connector that is used to connect segments.
	//
	// example:
	//
	// ...
	Ellipsis *string `json:"ellipsis,omitempty" xml:"ellipsis,omitempty"`
	// The field.
	//
	// example:
	//
	// field1
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	// The length of the segment. Valid values: 1 to 300.
	//
	// example:
	//
	// 50
	Len *int32 `json:"len,omitempty" xml:"len,omitempty"`
	// The number of segments. Valid values: 1 to 5.
	//
	// example:
	//
	// 1
	Snippet *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
}

func (s CreateAppRequestSummariesMeta) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestSummariesMeta) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSummariesMeta) GetElement() *string {
	return s.Element
}

func (s *CreateAppRequestSummariesMeta) GetEllipsis() *string {
	return s.Ellipsis
}

func (s *CreateAppRequestSummariesMeta) GetField() *string {
	return s.Field
}

func (s *CreateAppRequestSummariesMeta) GetLen() *int32 {
	return s.Len
}

func (s *CreateAppRequestSummariesMeta) GetSnippet() *string {
	return s.Snippet
}

func (s *CreateAppRequestSummariesMeta) SetElement(v string) *CreateAppRequestSummariesMeta {
	s.Element = &v
	return s
}

func (s *CreateAppRequestSummariesMeta) SetEllipsis(v string) *CreateAppRequestSummariesMeta {
	s.Ellipsis = &v
	return s
}

func (s *CreateAppRequestSummariesMeta) SetField(v string) *CreateAppRequestSummariesMeta {
	s.Field = &v
	return s
}

func (s *CreateAppRequestSummariesMeta) SetLen(v int32) *CreateAppRequestSummariesMeta {
	s.Len = &v
	return s
}

func (s *CreateAppRequestSummariesMeta) SetSnippet(v string) *CreateAppRequestSummariesMeta {
	s.Snippet = &v
	return s
}

func (s *CreateAppRequestSummariesMeta) Validate() error {
	return dara.Validate(s)
}
