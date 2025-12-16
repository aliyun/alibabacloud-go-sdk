// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAppResponseBody
	GetRequestId() *string
	SetResult(v *CreateAppResponseBodyResult) *CreateAppResponseBody
	GetResult() *CreateAppResponseBodyResult
}

type CreateAppResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABCDEFG
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// {}
	Result *CreateAppResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppResponseBody) GetResult() *CreateAppResponseBodyResult {
	return s.Result
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetResult(v *CreateAppResponseBodyResult) *CreateAppResponseBody {
	s.Result = v
	return s
}

func (s *CreateAppResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppResponseBodyResult struct {
	// The capability opening configurations.
	Cluster *CreateAppResponseBodyResultCluster `json:"cluster,omitempty" xml:"cluster,omitempty" type:"Struct"`
	// Deprecated
	//
	// The name of the cluster.
	//
	// example:
	//
	// vpc_sh_domain_1
	ClusterName *string                  `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	ConfigItems []map[string]interface{} `json:"configItems,omitempty" xml:"configItems,omitempty" type:"Repeated"`
	Created     *int64                   `json:"created,omitempty" xml:"created,omitempty"`
	// The configurations of the data sources.
	DataSources []*CreateAppResponseBodyResultDataSources `json:"dataSources,omitempty" xml:"dataSources,omitempty" type:"Repeated"`
	// The description of the application.
	//
	// example:
	//
	// My application
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The industry model module.
	Domain *CreateAppResponseBodyResultDomain `json:"domain,omitempty" xml:"domain,omitempty" type:"Struct"`
	// The default display fields.
	FetchFields []*string `json:"fetchFields,omitempty" xml:"fetchFields,omitempty" type:"Repeated"`
	// The configurations of rough sort.
	FirstRanks []*CreateAppResponseBodyResultFirstRanks `json:"firstRanks,omitempty" xml:"firstRanks,omitempty" type:"Repeated"`
	// The application ID.
	//
	// example:
	//
	// 12888
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The descriptions of the LLM table fields.
	//
	// example:
	//
	// { "name": "longyu_llm_1", "schemas": [], "interpretations": [ { "table": "table1", "fields": [ { "name": "field1", "interpretation": "Title" }, { "name": "field2", "interpretation": "Number" } ] } ] }
	Interpretations []map[string]interface{} `json:"interpretations,omitempty" xml:"interpretations,omitempty" type:"Repeated"`
	// Indicates whether the version is an online version.
	//
	// example:
	//
	// 12333
	IsCurrent *bool `json:"isCurrent,omitempty" xml:"isCurrent,omitempty"`
	// The percentage for the data import progress.
	//
	// example:
	//
	// 100
	ProgressPercent *int32 `json:"progressPercent,omitempty" xml:"progressPercent,omitempty"`
	// The prompt configurations
	Prompts []map[string]interface{} `json:"prompts,omitempty" xml:"prompts,omitempty" type:"Repeated"`
	// The query intent understanding configurations.
	QueryProcessors []*CreateAppResponseBodyResultQueryProcessors `json:"queryProcessors,omitempty" xml:"queryProcessors,omitempty" type:"Repeated"`
	// The quota.
	Quota *CreateAppResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The single-table schema.
	Schema *CreateAppResponseBodyResultSchema `json:"schema,omitempty" xml:"schema,omitempty" type:"Struct"`
	// The multi-table schema.
	Schemas []*CreateAppResponseBodyResultSchemas `json:"schemas,omitempty" xml:"schemas,omitempty" type:"Repeated"`
	// The configurations of fine sort.
	SecondRanks []*CreateAppResponseBodyResultSecondRanks `json:"secondRanks,omitempty" xml:"secondRanks,omitempty" type:"Repeated"`
	// The status of the application. Valid values:
	//
	// 	- OK
	//
	// 	- STOPPED
	//
	// 	- FROZEN
	//
	// 	- INITIALIZING
	//
	// 	- UNAVAILABLE
	//
	// 	- DATA_WAITING
	//
	// 	- DATA_PREPARING
	//
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The summary configurations of search results.
	Summaries  []*CreateAppResponseBodyResultSummaries `json:"summaries,omitempty" xml:"summaries,omitempty" type:"Repeated"`
	SwitchTime *int64                                  `json:"switchTime,omitempty" xml:"switchTime,omitempty"`
	// The type of the application. Valid values:
	//
	// 	- standard
	//
	// 	- enhanced
	//
	// example:
	//
	// standard
	Type    *string `json:"type,omitempty" xml:"type,omitempty"`
	Updated *int64  `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateAppResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResult) GetCluster() *CreateAppResponseBodyResultCluster {
	return s.Cluster
}

func (s *CreateAppResponseBodyResult) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateAppResponseBodyResult) GetConfigItems() []map[string]interface{} {
	return s.ConfigItems
}

func (s *CreateAppResponseBodyResult) GetCreated() *int64 {
	return s.Created
}

func (s *CreateAppResponseBodyResult) GetDataSources() []*CreateAppResponseBodyResultDataSources {
	return s.DataSources
}

func (s *CreateAppResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *CreateAppResponseBodyResult) GetDomain() *CreateAppResponseBodyResultDomain {
	return s.Domain
}

func (s *CreateAppResponseBodyResult) GetFetchFields() []*string {
	return s.FetchFields
}

func (s *CreateAppResponseBodyResult) GetFirstRanks() []*CreateAppResponseBodyResultFirstRanks {
	return s.FirstRanks
}

func (s *CreateAppResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *CreateAppResponseBodyResult) GetInterpretations() []map[string]interface{} {
	return s.Interpretations
}

func (s *CreateAppResponseBodyResult) GetIsCurrent() *bool {
	return s.IsCurrent
}

func (s *CreateAppResponseBodyResult) GetProgressPercent() *int32 {
	return s.ProgressPercent
}

func (s *CreateAppResponseBodyResult) GetPrompts() []map[string]interface{} {
	return s.Prompts
}

func (s *CreateAppResponseBodyResult) GetQueryProcessors() []*CreateAppResponseBodyResultQueryProcessors {
	return s.QueryProcessors
}

func (s *CreateAppResponseBodyResult) GetQuota() *CreateAppResponseBodyResultQuota {
	return s.Quota
}

func (s *CreateAppResponseBodyResult) GetSchema() *CreateAppResponseBodyResultSchema {
	return s.Schema
}

func (s *CreateAppResponseBodyResult) GetSchemas() []*CreateAppResponseBodyResultSchemas {
	return s.Schemas
}

func (s *CreateAppResponseBodyResult) GetSecondRanks() []*CreateAppResponseBodyResultSecondRanks {
	return s.SecondRanks
}

func (s *CreateAppResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *CreateAppResponseBodyResult) GetSummaries() []*CreateAppResponseBodyResultSummaries {
	return s.Summaries
}

func (s *CreateAppResponseBodyResult) GetSwitchTime() *int64 {
	return s.SwitchTime
}

func (s *CreateAppResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *CreateAppResponseBodyResult) GetUpdated() *int64 {
	return s.Updated
}

func (s *CreateAppResponseBodyResult) SetCluster(v *CreateAppResponseBodyResultCluster) *CreateAppResponseBodyResult {
	s.Cluster = v
	return s
}

func (s *CreateAppResponseBodyResult) SetClusterName(v string) *CreateAppResponseBodyResult {
	s.ClusterName = &v
	return s
}

func (s *CreateAppResponseBodyResult) SetConfigItems(v []map[string]interface{}) *CreateAppResponseBodyResult {
	s.ConfigItems = v
	return s
}

func (s *CreateAppResponseBodyResult) SetCreated(v int64) *CreateAppResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateAppResponseBodyResult) SetDataSources(v []*CreateAppResponseBodyResultDataSources) *CreateAppResponseBodyResult {
	s.DataSources = v
	return s
}

func (s *CreateAppResponseBodyResult) SetDescription(v string) *CreateAppResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateAppResponseBodyResult) SetDomain(v *CreateAppResponseBodyResultDomain) *CreateAppResponseBodyResult {
	s.Domain = v
	return s
}

func (s *CreateAppResponseBodyResult) SetFetchFields(v []*string) *CreateAppResponseBodyResult {
	s.FetchFields = v
	return s
}

func (s *CreateAppResponseBodyResult) SetFirstRanks(v []*CreateAppResponseBodyResultFirstRanks) *CreateAppResponseBodyResult {
	s.FirstRanks = v
	return s
}

func (s *CreateAppResponseBodyResult) SetId(v string) *CreateAppResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateAppResponseBodyResult) SetInterpretations(v []map[string]interface{}) *CreateAppResponseBodyResult {
	s.Interpretations = v
	return s
}

func (s *CreateAppResponseBodyResult) SetIsCurrent(v bool) *CreateAppResponseBodyResult {
	s.IsCurrent = &v
	return s
}

func (s *CreateAppResponseBodyResult) SetProgressPercent(v int32) *CreateAppResponseBodyResult {
	s.ProgressPercent = &v
	return s
}

func (s *CreateAppResponseBodyResult) SetPrompts(v []map[string]interface{}) *CreateAppResponseBodyResult {
	s.Prompts = v
	return s
}

func (s *CreateAppResponseBodyResult) SetQueryProcessors(v []*CreateAppResponseBodyResultQueryProcessors) *CreateAppResponseBodyResult {
	s.QueryProcessors = v
	return s
}

func (s *CreateAppResponseBodyResult) SetQuota(v *CreateAppResponseBodyResultQuota) *CreateAppResponseBodyResult {
	s.Quota = v
	return s
}

func (s *CreateAppResponseBodyResult) SetSchema(v *CreateAppResponseBodyResultSchema) *CreateAppResponseBodyResult {
	s.Schema = v
	return s
}

func (s *CreateAppResponseBodyResult) SetSchemas(v []*CreateAppResponseBodyResultSchemas) *CreateAppResponseBodyResult {
	s.Schemas = v
	return s
}

func (s *CreateAppResponseBodyResult) SetSecondRanks(v []*CreateAppResponseBodyResultSecondRanks) *CreateAppResponseBodyResult {
	s.SecondRanks = v
	return s
}

func (s *CreateAppResponseBodyResult) SetStatus(v string) *CreateAppResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateAppResponseBodyResult) SetSummaries(v []*CreateAppResponseBodyResultSummaries) *CreateAppResponseBodyResult {
	s.Summaries = v
	return s
}

func (s *CreateAppResponseBodyResult) SetSwitchTime(v int64) *CreateAppResponseBodyResult {
	s.SwitchTime = &v
	return s
}

func (s *CreateAppResponseBodyResult) SetType(v string) *CreateAppResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateAppResponseBodyResult) SetUpdated(v int64) *CreateAppResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *CreateAppResponseBodyResult) Validate() error {
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
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
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

type CreateAppResponseBodyResultCluster struct {
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
	MaxTimeoutMS             *int32                   `json:"maxTimeoutMS,omitempty" xml:"maxTimeoutMS,omitempty"`
	TextEmbeddingModel       *string                  `json:"textEmbeddingModel,omitempty" xml:"textEmbeddingModel,omitempty"`
	TextSparseEmbeddingModel *string                  `json:"textSparseEmbeddingModel,omitempty" xml:"textSparseEmbeddingModel,omitempty"`
	VectorIndexConfigs       []map[string]interface{} `json:"vectorIndexConfigs,omitempty" xml:"vectorIndexConfigs,omitempty" type:"Repeated"`
}

func (s CreateAppResponseBodyResultCluster) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultCluster) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultCluster) GetChunkModels() []map[string]interface{} {
	return s.ChunkModels
}

func (s *CreateAppResponseBodyResultCluster) GetGraphRag() map[string]interface{} {
	return s.GraphRag
}

func (s *CreateAppResponseBodyResultCluster) GetImageContentRecognizerModels() []map[string]interface{} {
	return s.ImageContentRecognizerModels
}

func (s *CreateAppResponseBodyResultCluster) GetMaxQueryClauseLength() *int32 {
	return s.MaxQueryClauseLength
}

func (s *CreateAppResponseBodyResultCluster) GetMaxTimeoutMS() *int32 {
	return s.MaxTimeoutMS
}

func (s *CreateAppResponseBodyResultCluster) GetTextEmbeddingModel() *string {
	return s.TextEmbeddingModel
}

func (s *CreateAppResponseBodyResultCluster) GetTextSparseEmbeddingModel() *string {
	return s.TextSparseEmbeddingModel
}

func (s *CreateAppResponseBodyResultCluster) GetVectorIndexConfigs() []map[string]interface{} {
	return s.VectorIndexConfigs
}

func (s *CreateAppResponseBodyResultCluster) SetChunkModels(v []map[string]interface{}) *CreateAppResponseBodyResultCluster {
	s.ChunkModels = v
	return s
}

func (s *CreateAppResponseBodyResultCluster) SetGraphRag(v map[string]interface{}) *CreateAppResponseBodyResultCluster {
	s.GraphRag = v
	return s
}

func (s *CreateAppResponseBodyResultCluster) SetImageContentRecognizerModels(v []map[string]interface{}) *CreateAppResponseBodyResultCluster {
	s.ImageContentRecognizerModels = v
	return s
}

func (s *CreateAppResponseBodyResultCluster) SetMaxQueryClauseLength(v int32) *CreateAppResponseBodyResultCluster {
	s.MaxQueryClauseLength = &v
	return s
}

func (s *CreateAppResponseBodyResultCluster) SetMaxTimeoutMS(v int32) *CreateAppResponseBodyResultCluster {
	s.MaxTimeoutMS = &v
	return s
}

func (s *CreateAppResponseBodyResultCluster) SetTextEmbeddingModel(v string) *CreateAppResponseBodyResultCluster {
	s.TextEmbeddingModel = &v
	return s
}

func (s *CreateAppResponseBodyResultCluster) SetTextSparseEmbeddingModel(v string) *CreateAppResponseBodyResultCluster {
	s.TextSparseEmbeddingModel = &v
	return s
}

func (s *CreateAppResponseBodyResultCluster) SetVectorIndexConfigs(v []map[string]interface{}) *CreateAppResponseBodyResultCluster {
	s.VectorIndexConfigs = v
	return s
}

func (s *CreateAppResponseBodyResultCluster) Validate() error {
	return dara.Validate(s)
}

type CreateAppResponseBodyResultDataSources struct {
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

func (s CreateAppResponseBodyResultDataSources) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultDataSources) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultDataSources) GetFields() []map[string]interface{} {
	return s.Fields
}

func (s *CreateAppResponseBodyResultDataSources) GetKeyField() *string {
	return s.KeyField
}

func (s *CreateAppResponseBodyResultDataSources) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateAppResponseBodyResultDataSources) GetPlugins() map[string]interface{} {
	return s.Plugins
}

func (s *CreateAppResponseBodyResultDataSources) GetSchemaName() *string {
	return s.SchemaName
}

func (s *CreateAppResponseBodyResultDataSources) GetTableName() *string {
	return s.TableName
}

func (s *CreateAppResponseBodyResultDataSources) GetType() *string {
	return s.Type
}

func (s *CreateAppResponseBodyResultDataSources) SetFields(v []map[string]interface{}) *CreateAppResponseBodyResultDataSources {
	s.Fields = v
	return s
}

func (s *CreateAppResponseBodyResultDataSources) SetKeyField(v string) *CreateAppResponseBodyResultDataSources {
	s.KeyField = &v
	return s
}

func (s *CreateAppResponseBodyResultDataSources) SetParameters(v map[string]interface{}) *CreateAppResponseBodyResultDataSources {
	s.Parameters = v
	return s
}

func (s *CreateAppResponseBodyResultDataSources) SetPlugins(v map[string]interface{}) *CreateAppResponseBodyResultDataSources {
	s.Plugins = v
	return s
}

func (s *CreateAppResponseBodyResultDataSources) SetSchemaName(v string) *CreateAppResponseBodyResultDataSources {
	s.SchemaName = &v
	return s
}

func (s *CreateAppResponseBodyResultDataSources) SetTableName(v string) *CreateAppResponseBodyResultDataSources {
	s.TableName = &v
	return s
}

func (s *CreateAppResponseBodyResultDataSources) SetType(v string) *CreateAppResponseBodyResultDataSources {
	s.Type = &v
	return s
}

func (s *CreateAppResponseBodyResultDataSources) Validate() error {
	return dara.Validate(s)
}

type CreateAppResponseBodyResultDomain struct {
	// The industry category.
	//
	// example:
	//
	// -
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The selected features.
	Functions *CreateAppResponseBodyResultDomainFunctions `json:"functions,omitempty" xml:"functions,omitempty" type:"Struct"`
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
	// GENERAL
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateAppResponseBodyResultDomain) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultDomain) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultDomain) GetCategory() *string {
	return s.Category
}

func (s *CreateAppResponseBodyResultDomain) GetFunctions() *CreateAppResponseBodyResultDomainFunctions {
	return s.Functions
}

func (s *CreateAppResponseBodyResultDomain) GetName() *string {
	return s.Name
}

func (s *CreateAppResponseBodyResultDomain) SetCategory(v string) *CreateAppResponseBodyResultDomain {
	s.Category = &v
	return s
}

func (s *CreateAppResponseBodyResultDomain) SetFunctions(v *CreateAppResponseBodyResultDomainFunctions) *CreateAppResponseBodyResultDomain {
	s.Functions = v
	return s
}

func (s *CreateAppResponseBodyResultDomain) SetName(v string) *CreateAppResponseBodyResultDomain {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyResultDomain) Validate() error {
	if s.Functions != nil {
		if err := s.Functions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppResponseBodyResultDomainFunctions struct {
	// The features of the sort policy category.
	Algo []*string `json:"algo,omitempty" xml:"algo,omitempty" type:"Repeated"`
	// The features of the query analysis category.
	Qp []*string `json:"qp,omitempty" xml:"qp,omitempty" type:"Repeated"`
	// The features of the service category.
	Service []*string `json:"service,omitempty" xml:"service,omitempty" type:"Repeated"`
}

func (s CreateAppResponseBodyResultDomainFunctions) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultDomainFunctions) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultDomainFunctions) GetAlgo() []*string {
	return s.Algo
}

func (s *CreateAppResponseBodyResultDomainFunctions) GetQp() []*string {
	return s.Qp
}

func (s *CreateAppResponseBodyResultDomainFunctions) GetService() []*string {
	return s.Service
}

func (s *CreateAppResponseBodyResultDomainFunctions) SetAlgo(v []*string) *CreateAppResponseBodyResultDomainFunctions {
	s.Algo = v
	return s
}

func (s *CreateAppResponseBodyResultDomainFunctions) SetQp(v []*string) *CreateAppResponseBodyResultDomainFunctions {
	s.Qp = v
	return s
}

func (s *CreateAppResponseBodyResultDomainFunctions) SetService(v []*string) *CreateAppResponseBodyResultDomainFunctions {
	s.Service = v
	return s
}

func (s *CreateAppResponseBodyResultDomainFunctions) Validate() error {
	return dara.Validate(s)
}

type CreateAppResponseBodyResultFirstRanks struct {
	// Indicates whether the expression is the default one.
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
	// STRUCT: The content of the expression is a structure. STRING (default): You can configure a custom formula.
	//
	// example:
	//
	// STRING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAppResponseBodyResultFirstRanks) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultFirstRanks) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultFirstRanks) GetActive() *bool {
	return s.Active
}

func (s *CreateAppResponseBodyResultFirstRanks) GetDescription() *string {
	return s.Description
}

func (s *CreateAppResponseBodyResultFirstRanks) GetMeta() interface{} {
	return s.Meta
}

func (s *CreateAppResponseBodyResultFirstRanks) GetName() *string {
	return s.Name
}

func (s *CreateAppResponseBodyResultFirstRanks) GetType() *string {
	return s.Type
}

func (s *CreateAppResponseBodyResultFirstRanks) SetActive(v bool) *CreateAppResponseBodyResultFirstRanks {
	s.Active = &v
	return s
}

func (s *CreateAppResponseBodyResultFirstRanks) SetDescription(v string) *CreateAppResponseBodyResultFirstRanks {
	s.Description = &v
	return s
}

func (s *CreateAppResponseBodyResultFirstRanks) SetMeta(v interface{}) *CreateAppResponseBodyResultFirstRanks {
	s.Meta = v
	return s
}

func (s *CreateAppResponseBodyResultFirstRanks) SetName(v string) *CreateAppResponseBodyResultFirstRanks {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyResultFirstRanks) SetType(v string) *CreateAppResponseBodyResultFirstRanks {
	s.Type = &v
	return s
}

func (s *CreateAppResponseBodyResultFirstRanks) Validate() error {
	return dara.Validate(s)
}

type CreateAppResponseBodyResultQueryProcessors struct {
	// Indicates whether the rule is the default one.
	//
	// example:
	//
	// False
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
	// GENERAL
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The index range.
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The rule name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The features.
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s CreateAppResponseBodyResultQueryProcessors) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultQueryProcessors) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultQueryProcessors) GetActive() *bool {
	return s.Active
}

func (s *CreateAppResponseBodyResultQueryProcessors) GetCategory() *string {
	return s.Category
}

func (s *CreateAppResponseBodyResultQueryProcessors) GetDomain() *string {
	return s.Domain
}

func (s *CreateAppResponseBodyResultQueryProcessors) GetIndexes() []*string {
	return s.Indexes
}

func (s *CreateAppResponseBodyResultQueryProcessors) GetName() *string {
	return s.Name
}

func (s *CreateAppResponseBodyResultQueryProcessors) GetProcessors() []map[string]interface{} {
	return s.Processors
}

func (s *CreateAppResponseBodyResultQueryProcessors) SetActive(v bool) *CreateAppResponseBodyResultQueryProcessors {
	s.Active = &v
	return s
}

func (s *CreateAppResponseBodyResultQueryProcessors) SetCategory(v string) *CreateAppResponseBodyResultQueryProcessors {
	s.Category = &v
	return s
}

func (s *CreateAppResponseBodyResultQueryProcessors) SetDomain(v string) *CreateAppResponseBodyResultQueryProcessors {
	s.Domain = &v
	return s
}

func (s *CreateAppResponseBodyResultQueryProcessors) SetIndexes(v []*string) *CreateAppResponseBodyResultQueryProcessors {
	s.Indexes = v
	return s
}

func (s *CreateAppResponseBodyResultQueryProcessors) SetName(v string) *CreateAppResponseBodyResultQueryProcessors {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyResultQueryProcessors) SetProcessors(v []map[string]interface{}) *CreateAppResponseBodyResultQueryProcessors {
	s.Processors = v
	return s
}

func (s *CreateAppResponseBodyResultQueryProcessors) Validate() error {
	return dara.Validate(s)
}

type CreateAppResponseBodyResultQuota struct {
	// The computing resources.
	//
	// example:
	//
	// 20
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity.
	//
	// example:
	//
	// 1
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The search request.
	//
	// example:
	//
	// 5
	Qps *int32 `json:"qps,omitempty" xml:"qps,omitempty"`
	// The specifications. Valid values:
	//
	// 	- opensearch.share.junior: basic
	//
	// 	- opensearch.share.common: shared general-purpose
	//
	// 	- opensearch.share.compute: shared computing
	//
	// 	- opensearch.share.storage: shared storage
	//
	// 	- opensearch.private.common: exclusive general-purpose
	//
	// 	- opensearch.private.compute: exclusive computing
	//
	// 	- opensearch.private.storage: exclusive storage
	//
	// example:
	//
	// opensearch.share.common
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
	// example:
	//
	// 100
	UsedComputeResource *int32 `json:"usedComputeResource,omitempty" xml:"usedComputeResource,omitempty"`
	// example:
	//
	// 1024
	UsedDocSize *float64 `json:"usedDocSize,omitempty" xml:"usedDocSize,omitempty"`
	// example:
	//
	// 100
	UsedQps *int32 `json:"usedQps,omitempty" xml:"usedQps,omitempty"`
}

func (s CreateAppResponseBodyResultQuota) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultQuota) GetComputeResource() *int32 {
	return s.ComputeResource
}

func (s *CreateAppResponseBodyResultQuota) GetDocSize() *int32 {
	return s.DocSize
}

func (s *CreateAppResponseBodyResultQuota) GetQps() *int32 {
	return s.Qps
}

func (s *CreateAppResponseBodyResultQuota) GetSpec() *string {
	return s.Spec
}

func (s *CreateAppResponseBodyResultQuota) GetUsedComputeResource() *int32 {
	return s.UsedComputeResource
}

func (s *CreateAppResponseBodyResultQuota) GetUsedDocSize() *float64 {
	return s.UsedDocSize
}

func (s *CreateAppResponseBodyResultQuota) GetUsedQps() *int32 {
	return s.UsedQps
}

func (s *CreateAppResponseBodyResultQuota) SetComputeResource(v int32) *CreateAppResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *CreateAppResponseBodyResultQuota) SetDocSize(v int32) *CreateAppResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *CreateAppResponseBodyResultQuota) SetQps(v int32) *CreateAppResponseBodyResultQuota {
	s.Qps = &v
	return s
}

func (s *CreateAppResponseBodyResultQuota) SetSpec(v string) *CreateAppResponseBodyResultQuota {
	s.Spec = &v
	return s
}

func (s *CreateAppResponseBodyResultQuota) SetUsedComputeResource(v int32) *CreateAppResponseBodyResultQuota {
	s.UsedComputeResource = &v
	return s
}

func (s *CreateAppResponseBodyResultQuota) SetUsedDocSize(v float64) *CreateAppResponseBodyResultQuota {
	s.UsedDocSize = &v
	return s
}

func (s *CreateAppResponseBodyResultQuota) SetUsedQps(v int32) *CreateAppResponseBodyResultQuota {
	s.UsedQps = &v
	return s
}

func (s *CreateAppResponseBodyResultQuota) Validate() error {
	return dara.Validate(s)
}

type CreateAppResponseBodyResultSchema struct {
	// The sort configurations.
	IndexSortConfig []*CreateAppResponseBodyResultSchemaIndexSortConfig `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	// The index schema.
	Indexes *CreateAppResponseBodyResultSchemaIndexes `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
	// The name of the wide table.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the level-1 routing field.
	//
	// example:
	//
	// field1
	RouteField *string `json:"routeField,omitempty" xml:"routeField,omitempty"`
	// The name of the level-2 routing field. This parameter takes effect only when the routeFieldValues parameter is configured. By default, the wide-table primary key field is used as the level-2 routing field.
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
	TtlField *CreateAppResponseBodyResultSchemaTtlField `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s CreateAppResponseBodyResultSchema) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultSchema) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSchema) GetIndexSortConfig() []*CreateAppResponseBodyResultSchemaIndexSortConfig {
	return s.IndexSortConfig
}

func (s *CreateAppResponseBodyResultSchema) GetIndexes() *CreateAppResponseBodyResultSchemaIndexes {
	return s.Indexes
}

func (s *CreateAppResponseBodyResultSchema) GetName() *string {
	return s.Name
}

func (s *CreateAppResponseBodyResultSchema) GetRouteField() *string {
	return s.RouteField
}

func (s *CreateAppResponseBodyResultSchema) GetRouteFieldValues() []*string {
	return s.RouteFieldValues
}

func (s *CreateAppResponseBodyResultSchema) GetSecondRouteField() *string {
	return s.SecondRouteField
}

func (s *CreateAppResponseBodyResultSchema) GetTables() map[string]interface{} {
	return s.Tables
}

func (s *CreateAppResponseBodyResultSchema) GetTtlField() *CreateAppResponseBodyResultSchemaTtlField {
	return s.TtlField
}

func (s *CreateAppResponseBodyResultSchema) SetIndexSortConfig(v []*CreateAppResponseBodyResultSchemaIndexSortConfig) *CreateAppResponseBodyResultSchema {
	s.IndexSortConfig = v
	return s
}

func (s *CreateAppResponseBodyResultSchema) SetIndexes(v *CreateAppResponseBodyResultSchemaIndexes) *CreateAppResponseBodyResultSchema {
	s.Indexes = v
	return s
}

func (s *CreateAppResponseBodyResultSchema) SetName(v string) *CreateAppResponseBodyResultSchema {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyResultSchema) SetRouteField(v string) *CreateAppResponseBodyResultSchema {
	s.RouteField = &v
	return s
}

func (s *CreateAppResponseBodyResultSchema) SetRouteFieldValues(v []*string) *CreateAppResponseBodyResultSchema {
	s.RouteFieldValues = v
	return s
}

func (s *CreateAppResponseBodyResultSchema) SetSecondRouteField(v string) *CreateAppResponseBodyResultSchema {
	s.SecondRouteField = &v
	return s
}

func (s *CreateAppResponseBodyResultSchema) SetTables(v map[string]interface{}) *CreateAppResponseBodyResultSchema {
	s.Tables = v
	return s
}

func (s *CreateAppResponseBodyResultSchema) SetTtlField(v *CreateAppResponseBodyResultSchemaTtlField) *CreateAppResponseBodyResultSchema {
	s.TtlField = v
	return s
}

func (s *CreateAppResponseBodyResultSchema) Validate() error {
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

type CreateAppResponseBodyResultSchemaIndexSortConfig struct {
	// The sort method. Valid values:
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

func (s CreateAppResponseBodyResultSchemaIndexSortConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultSchemaIndexSortConfig) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSchemaIndexSortConfig) GetDirection() *string {
	return s.Direction
}

func (s *CreateAppResponseBodyResultSchemaIndexSortConfig) GetField() *string {
	return s.Field
}

func (s *CreateAppResponseBodyResultSchemaIndexSortConfig) SetDirection(v string) *CreateAppResponseBodyResultSchemaIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemaIndexSortConfig) SetField(v string) *CreateAppResponseBodyResultSchemaIndexSortConfig {
	s.Field = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemaIndexSortConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAppResponseBodyResultSchemaIndexes struct {
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

func (s CreateAppResponseBodyResultSchemaIndexes) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultSchemaIndexes) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSchemaIndexes) GetFilterFields() []*string {
	return s.FilterFields
}

func (s *CreateAppResponseBodyResultSchemaIndexes) GetSearchFields() map[string]interface{} {
	return s.SearchFields
}

func (s *CreateAppResponseBodyResultSchemaIndexes) SetFilterFields(v []*string) *CreateAppResponseBodyResultSchemaIndexes {
	s.FilterFields = v
	return s
}

func (s *CreateAppResponseBodyResultSchemaIndexes) SetSearchFields(v map[string]interface{}) *CreateAppResponseBodyResultSchemaIndexes {
	s.SearchFields = v
	return s
}

func (s *CreateAppResponseBodyResultSchemaIndexes) Validate() error {
	return dara.Validate(s)
}

type CreateAppResponseBodyResultSchemaTtlField struct {
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

func (s CreateAppResponseBodyResultSchemaTtlField) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultSchemaTtlField) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSchemaTtlField) GetName() *string {
	return s.Name
}

func (s *CreateAppResponseBodyResultSchemaTtlField) GetTtl() *int64 {
	return s.Ttl
}

func (s *CreateAppResponseBodyResultSchemaTtlField) SetName(v string) *CreateAppResponseBodyResultSchemaTtlField {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemaTtlField) SetTtl(v int64) *CreateAppResponseBodyResultSchemaTtlField {
	s.Ttl = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemaTtlField) Validate() error {
	return dara.Validate(s)
}

type CreateAppResponseBodyResultSchemas struct {
	// The sort configurations.
	IndexSortConfig []*CreateAppResponseBodyResultSchemasIndexSortConfig `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	// The index schema.
	Indexes *CreateAppResponseBodyResultSchemasIndexes `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
	// The name of the wide table.
	//
	// example:
	//
	// main
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
	TtlField *CreateAppResponseBodyResultSchemasTtlField `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s CreateAppResponseBodyResultSchemas) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultSchemas) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSchemas) GetIndexSortConfig() []*CreateAppResponseBodyResultSchemasIndexSortConfig {
	return s.IndexSortConfig
}

func (s *CreateAppResponseBodyResultSchemas) GetIndexes() *CreateAppResponseBodyResultSchemasIndexes {
	return s.Indexes
}

func (s *CreateAppResponseBodyResultSchemas) GetName() *string {
	return s.Name
}

func (s *CreateAppResponseBodyResultSchemas) GetRouteField() *string {
	return s.RouteField
}

func (s *CreateAppResponseBodyResultSchemas) GetRouteFieldValues() []*string {
	return s.RouteFieldValues
}

func (s *CreateAppResponseBodyResultSchemas) GetSecondRouteField() *string {
	return s.SecondRouteField
}

func (s *CreateAppResponseBodyResultSchemas) GetTables() map[string]interface{} {
	return s.Tables
}

func (s *CreateAppResponseBodyResultSchemas) GetTtlField() *CreateAppResponseBodyResultSchemasTtlField {
	return s.TtlField
}

func (s *CreateAppResponseBodyResultSchemas) SetIndexSortConfig(v []*CreateAppResponseBodyResultSchemasIndexSortConfig) *CreateAppResponseBodyResultSchemas {
	s.IndexSortConfig = v
	return s
}

func (s *CreateAppResponseBodyResultSchemas) SetIndexes(v *CreateAppResponseBodyResultSchemasIndexes) *CreateAppResponseBodyResultSchemas {
	s.Indexes = v
	return s
}

func (s *CreateAppResponseBodyResultSchemas) SetName(v string) *CreateAppResponseBodyResultSchemas {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemas) SetRouteField(v string) *CreateAppResponseBodyResultSchemas {
	s.RouteField = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemas) SetRouteFieldValues(v []*string) *CreateAppResponseBodyResultSchemas {
	s.RouteFieldValues = v
	return s
}

func (s *CreateAppResponseBodyResultSchemas) SetSecondRouteField(v string) *CreateAppResponseBodyResultSchemas {
	s.SecondRouteField = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemas) SetTables(v map[string]interface{}) *CreateAppResponseBodyResultSchemas {
	s.Tables = v
	return s
}

func (s *CreateAppResponseBodyResultSchemas) SetTtlField(v *CreateAppResponseBodyResultSchemasTtlField) *CreateAppResponseBodyResultSchemas {
	s.TtlField = v
	return s
}

func (s *CreateAppResponseBodyResultSchemas) Validate() error {
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

type CreateAppResponseBodyResultSchemasIndexSortConfig struct {
	// The sort method. Valid values:
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
	// fIeld1
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
}

func (s CreateAppResponseBodyResultSchemasIndexSortConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultSchemasIndexSortConfig) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSchemasIndexSortConfig) GetDirection() *string {
	return s.Direction
}

func (s *CreateAppResponseBodyResultSchemasIndexSortConfig) GetField() *string {
	return s.Field
}

func (s *CreateAppResponseBodyResultSchemasIndexSortConfig) SetDirection(v string) *CreateAppResponseBodyResultSchemasIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemasIndexSortConfig) SetField(v string) *CreateAppResponseBodyResultSchemasIndexSortConfig {
	s.Field = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemasIndexSortConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAppResponseBodyResultSchemasIndexes struct {
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

func (s CreateAppResponseBodyResultSchemasIndexes) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultSchemasIndexes) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSchemasIndexes) GetFilterFields() []*string {
	return s.FilterFields
}

func (s *CreateAppResponseBodyResultSchemasIndexes) GetSearchFields() map[string]interface{} {
	return s.SearchFields
}

func (s *CreateAppResponseBodyResultSchemasIndexes) SetFilterFields(v []*string) *CreateAppResponseBodyResultSchemasIndexes {
	s.FilterFields = v
	return s
}

func (s *CreateAppResponseBodyResultSchemasIndexes) SetSearchFields(v map[string]interface{}) *CreateAppResponseBodyResultSchemasIndexes {
	s.SearchFields = v
	return s
}

func (s *CreateAppResponseBodyResultSchemasIndexes) Validate() error {
	return dara.Validate(s)
}

type CreateAppResponseBodyResultSchemasTtlField struct {
	// The name of the document time field.
	//
	// example:
	//
	// fIeld1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The TTL. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	Ttl *int64 `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s CreateAppResponseBodyResultSchemasTtlField) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultSchemasTtlField) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSchemasTtlField) GetName() *string {
	return s.Name
}

func (s *CreateAppResponseBodyResultSchemasTtlField) GetTtl() *int64 {
	return s.Ttl
}

func (s *CreateAppResponseBodyResultSchemasTtlField) SetName(v string) *CreateAppResponseBodyResultSchemasTtlField {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemasTtlField) SetTtl(v int64) *CreateAppResponseBodyResultSchemasTtlField {
	s.Ttl = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemasTtlField) Validate() error {
	return dara.Validate(s)
}

type CreateAppResponseBodyResultSecondRanks struct {
	// Indicates whether the expression is the default one.
	//
	// example:
	//
	// False
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The description.
	//
	// example:
	//
	// default
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

func (s CreateAppResponseBodyResultSecondRanks) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultSecondRanks) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSecondRanks) GetActive() *bool {
	return s.Active
}

func (s *CreateAppResponseBodyResultSecondRanks) GetDescription() *string {
	return s.Description
}

func (s *CreateAppResponseBodyResultSecondRanks) GetMeta() interface{} {
	return s.Meta
}

func (s *CreateAppResponseBodyResultSecondRanks) GetName() *string {
	return s.Name
}

func (s *CreateAppResponseBodyResultSecondRanks) SetActive(v bool) *CreateAppResponseBodyResultSecondRanks {
	s.Active = &v
	return s
}

func (s *CreateAppResponseBodyResultSecondRanks) SetDescription(v string) *CreateAppResponseBodyResultSecondRanks {
	s.Description = &v
	return s
}

func (s *CreateAppResponseBodyResultSecondRanks) SetMeta(v interface{}) *CreateAppResponseBodyResultSecondRanks {
	s.Meta = v
	return s
}

func (s *CreateAppResponseBodyResultSecondRanks) SetName(v string) *CreateAppResponseBodyResultSecondRanks {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyResultSecondRanks) Validate() error {
	return dara.Validate(s)
}

type CreateAppResponseBodyResultSummaries struct {
	// The collection of summary configurations.
	Meta []*CreateAppResponseBodyResultSummariesMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The group name.
	//
	// example:
	//
	// fefault
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateAppResponseBodyResultSummaries) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultSummaries) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSummaries) GetMeta() []*CreateAppResponseBodyResultSummariesMeta {
	return s.Meta
}

func (s *CreateAppResponseBodyResultSummaries) GetName() *string {
	return s.Name
}

func (s *CreateAppResponseBodyResultSummaries) SetMeta(v []*CreateAppResponseBodyResultSummariesMeta) *CreateAppResponseBodyResultSummaries {
	s.Meta = v
	return s
}

func (s *CreateAppResponseBodyResultSummaries) SetName(v string) *CreateAppResponseBodyResultSummaries {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyResultSummaries) Validate() error {
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

type CreateAppResponseBodyResultSummariesMeta struct {
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
	// 5
	Snippet *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
}

func (s CreateAppResponseBodyResultSummariesMeta) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResultSummariesMeta) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSummariesMeta) GetElement() *string {
	return s.Element
}

func (s *CreateAppResponseBodyResultSummariesMeta) GetEllipsis() *string {
	return s.Ellipsis
}

func (s *CreateAppResponseBodyResultSummariesMeta) GetField() *string {
	return s.Field
}

func (s *CreateAppResponseBodyResultSummariesMeta) GetLen() *int32 {
	return s.Len
}

func (s *CreateAppResponseBodyResultSummariesMeta) GetSnippet() *string {
	return s.Snippet
}

func (s *CreateAppResponseBodyResultSummariesMeta) SetElement(v string) *CreateAppResponseBodyResultSummariesMeta {
	s.Element = &v
	return s
}

func (s *CreateAppResponseBodyResultSummariesMeta) SetEllipsis(v string) *CreateAppResponseBodyResultSummariesMeta {
	s.Ellipsis = &v
	return s
}

func (s *CreateAppResponseBodyResultSummariesMeta) SetField(v string) *CreateAppResponseBodyResultSummariesMeta {
	s.Field = &v
	return s
}

func (s *CreateAppResponseBodyResultSummariesMeta) SetLen(v int32) *CreateAppResponseBodyResultSummariesMeta {
	s.Len = &v
	return s
}

func (s *CreateAppResponseBodyResultSummariesMeta) SetSnippet(v string) *CreateAppResponseBodyResultSummariesMeta {
	s.Snippet = &v
	return s
}

func (s *CreateAppResponseBodyResultSummariesMeta) Validate() error {
	return dara.Validate(s)
}
