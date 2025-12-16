// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAppResponseBody
	GetRequestId() *string
	SetResult(v *DescribeAppResponseBodyResult) *DescribeAppResponseBody
	GetResult() *DescribeAppResponseBodyResult
}

type DescribeAppResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 33477D76-C380-1D84-A4AD-043F52876CB1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The version information.
	//
	// example:
	//
	// {}
	Result *DescribeAppResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppResponseBody) GetResult() *DescribeAppResponseBodyResult {
	return s.Result
}

func (s *DescribeAppResponseBody) SetRequestId(v string) *DescribeAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppResponseBody) SetResult(v *DescribeAppResponseBodyResult) *DescribeAppResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAppResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAppResponseBodyResult struct {
	// Indicates whether the version is automatically switched to an online version.
	//
	// example:
	//
	// true
	AutoSwitch *bool `json:"autoSwitch,omitempty" xml:"autoSwitch,omitempty"`
	// The capability opening configurations.
	Cluster *DescribeAppResponseBodyResultCluster `json:"cluster,omitempty" xml:"cluster,omitempty" type:"Struct"`
	// The cluster name.
	//
	// example:
	//
	// -
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	// The configurations of the data sources.
	DataSources []*DescribeAppResponseBodyResultDataSources `json:"dataSources,omitempty" xml:"dataSources,omitempty" type:"Repeated"`
	// The description of the version.
	//
	// example:
	//
	// -
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
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
	Domain *DescribeAppResponseBodyResultDomain `json:"domain,omitempty" xml:"domain,omitempty" type:"Struct"`
	// The default display fields.
	//
	// example:
	//
	// []
	FetchFields []*string `json:"fetchFields,omitempty" xml:"fetchFields,omitempty" type:"Repeated"`
	// The configurations of rough sort.
	FirstRanks []*DescribeAppResponseBodyResultFirstRanks `json:"firstRanks,omitempty" xml:"firstRanks,omitempty" type:"Repeated"`
	// The version ID.
	//
	// example:
	//
	// 100303063
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The industry model module.
	//
	// example:
	//
	// [ { "table": "table1", "fields": [ { "name": "field1", "interpretation": "Title" }, { "name": "field2", "interpretation": "Number" } ] } ]
	Interpretations map[string]interface{} `json:"interpretations,omitempty" xml:"interpretations,omitempty"`
	// Indices whether the version is an online version.
	//
	// example:
	//
	// True
	IsCurrent *bool `json:"isCurrent,omitempty" xml:"isCurrent,omitempty"`
	// The progress of data import, in percentage. For example, a value of 83 indicates 83%.
	//
	// example:
	//
	// 100
	ProgressPercent *int32 `json:"progressPercent,omitempty" xml:"progressPercent,omitempty"`
	// The prompt configurations.
	Prompts []map[string]interface{} `json:"prompts,omitempty" xml:"prompts,omitempty" type:"Repeated"`
	// The query intent understanding configurations.
	QueryProcessors []*DescribeAppResponseBodyResultQueryProcessors `json:"queryProcessors,omitempty" xml:"queryProcessors,omitempty" type:"Repeated"`
	// The quota information.
	//
	// example:
	//
	// {}
	Quota *DescribeAppResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The schema of the application.
	//
	// example:
	//
	// {}
	Schema *DescribeAppResponseBodyResultSchema `json:"schema,omitempty" xml:"schema,omitempty" type:"Struct"`
	// The multi-table schema.
	Schemas []*DescribeAppResponseBodyResultSchemas `json:"schemas,omitempty" xml:"schemas,omitempty" type:"Repeated"`
	// The configurations of fine sort.
	SecondRanks []*DescribeAppResponseBodyResultSecondRanks `json:"secondRanks,omitempty" xml:"secondRanks,omitempty" type:"Repeated"`
	// The status of the version. Valid values:
	//
	// 	- ok: The version is normal.
	//
	// 	- stopped: The version is suspended.
	//
	// 	- frozen: The version is frozen.
	//
	// 	- initializing: The version is being initialized.
	//
	// 	- unavailable: The version is invalid.
	//
	// 	- data_waiting: Data is to be initialized.
	//
	// 	- data_preparing: Data is being initialized.
	//
	// example:
	//
	// ok
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The summary configurations of search results.
	Summaries []*DescribeAppResponseBodyResultSummaries `json:"summaries,omitempty" xml:"summaries,omitempty" type:"Repeated"`
	// The edition type. Valid values:
	//
	// 	- standard: a standard edition application.
	//
	// 	- advance: an advanced edition application of an old version. New versions are not supported for this edition.
	//
	// 	- enhanced: an advanced edition application of a new version.
	//
	// example:
	//
	// enhanced
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeAppResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResult) GetAutoSwitch() *bool {
	return s.AutoSwitch
}

func (s *DescribeAppResponseBodyResult) GetCluster() *DescribeAppResponseBodyResultCluster {
	return s.Cluster
}

func (s *DescribeAppResponseBodyResult) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeAppResponseBodyResult) GetDataSources() []*DescribeAppResponseBodyResultDataSources {
	return s.DataSources
}

func (s *DescribeAppResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeAppResponseBodyResult) GetDomain() *DescribeAppResponseBodyResultDomain {
	return s.Domain
}

func (s *DescribeAppResponseBodyResult) GetFetchFields() []*string {
	return s.FetchFields
}

func (s *DescribeAppResponseBodyResult) GetFirstRanks() []*DescribeAppResponseBodyResultFirstRanks {
	return s.FirstRanks
}

func (s *DescribeAppResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *DescribeAppResponseBodyResult) GetInterpretations() map[string]interface{} {
	return s.Interpretations
}

func (s *DescribeAppResponseBodyResult) GetIsCurrent() *bool {
	return s.IsCurrent
}

func (s *DescribeAppResponseBodyResult) GetProgressPercent() *int32 {
	return s.ProgressPercent
}

func (s *DescribeAppResponseBodyResult) GetPrompts() []map[string]interface{} {
	return s.Prompts
}

func (s *DescribeAppResponseBodyResult) GetQueryProcessors() []*DescribeAppResponseBodyResultQueryProcessors {
	return s.QueryProcessors
}

func (s *DescribeAppResponseBodyResult) GetQuota() *DescribeAppResponseBodyResultQuota {
	return s.Quota
}

func (s *DescribeAppResponseBodyResult) GetSchema() *DescribeAppResponseBodyResultSchema {
	return s.Schema
}

func (s *DescribeAppResponseBodyResult) GetSchemas() []*DescribeAppResponseBodyResultSchemas {
	return s.Schemas
}

func (s *DescribeAppResponseBodyResult) GetSecondRanks() []*DescribeAppResponseBodyResultSecondRanks {
	return s.SecondRanks
}

func (s *DescribeAppResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *DescribeAppResponseBodyResult) GetSummaries() []*DescribeAppResponseBodyResultSummaries {
	return s.Summaries
}

func (s *DescribeAppResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *DescribeAppResponseBodyResult) SetAutoSwitch(v bool) *DescribeAppResponseBodyResult {
	s.AutoSwitch = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetCluster(v *DescribeAppResponseBodyResultCluster) *DescribeAppResponseBodyResult {
	s.Cluster = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetClusterName(v string) *DescribeAppResponseBodyResult {
	s.ClusterName = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetDataSources(v []*DescribeAppResponseBodyResultDataSources) *DescribeAppResponseBodyResult {
	s.DataSources = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetDescription(v string) *DescribeAppResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetDomain(v *DescribeAppResponseBodyResultDomain) *DescribeAppResponseBodyResult {
	s.Domain = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetFetchFields(v []*string) *DescribeAppResponseBodyResult {
	s.FetchFields = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetFirstRanks(v []*DescribeAppResponseBodyResultFirstRanks) *DescribeAppResponseBodyResult {
	s.FirstRanks = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetId(v string) *DescribeAppResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetInterpretations(v map[string]interface{}) *DescribeAppResponseBodyResult {
	s.Interpretations = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetIsCurrent(v bool) *DescribeAppResponseBodyResult {
	s.IsCurrent = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetProgressPercent(v int32) *DescribeAppResponseBodyResult {
	s.ProgressPercent = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetPrompts(v []map[string]interface{}) *DescribeAppResponseBodyResult {
	s.Prompts = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetQueryProcessors(v []*DescribeAppResponseBodyResultQueryProcessors) *DescribeAppResponseBodyResult {
	s.QueryProcessors = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetQuota(v *DescribeAppResponseBodyResultQuota) *DescribeAppResponseBodyResult {
	s.Quota = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetSchema(v *DescribeAppResponseBodyResultSchema) *DescribeAppResponseBodyResult {
	s.Schema = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetSchemas(v []*DescribeAppResponseBodyResultSchemas) *DescribeAppResponseBodyResult {
	s.Schemas = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetSecondRanks(v []*DescribeAppResponseBodyResultSecondRanks) *DescribeAppResponseBodyResult {
	s.SecondRanks = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetStatus(v string) *DescribeAppResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetSummaries(v []*DescribeAppResponseBodyResultSummaries) *DescribeAppResponseBodyResult {
	s.Summaries = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetType(v string) *DescribeAppResponseBodyResult {
	s.Type = &v
	return s
}

func (s *DescribeAppResponseBodyResult) Validate() error {
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

type DescribeAppResponseBodyResultCluster struct {
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
}

func (s DescribeAppResponseBodyResultCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultCluster) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultCluster) GetMaxQueryClauseLength() *int32 {
	return s.MaxQueryClauseLength
}

func (s *DescribeAppResponseBodyResultCluster) GetMaxTimeoutMS() *int32 {
	return s.MaxTimeoutMS
}

func (s *DescribeAppResponseBodyResultCluster) SetMaxQueryClauseLength(v int32) *DescribeAppResponseBodyResultCluster {
	s.MaxQueryClauseLength = &v
	return s
}

func (s *DescribeAppResponseBodyResultCluster) SetMaxTimeoutMS(v int32) *DescribeAppResponseBodyResultCluster {
	s.MaxTimeoutMS = &v
	return s
}

func (s *DescribeAppResponseBodyResultCluster) Validate() error {
	return dara.Validate(s)
}

type DescribeAppResponseBodyResultDataSources struct {
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
	// name
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

func (s DescribeAppResponseBodyResultDataSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultDataSources) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultDataSources) GetFields() []map[string]interface{} {
	return s.Fields
}

func (s *DescribeAppResponseBodyResultDataSources) GetKeyField() *string {
	return s.KeyField
}

func (s *DescribeAppResponseBodyResultDataSources) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *DescribeAppResponseBodyResultDataSources) GetPlugins() map[string]interface{} {
	return s.Plugins
}

func (s *DescribeAppResponseBodyResultDataSources) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeAppResponseBodyResultDataSources) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAppResponseBodyResultDataSources) GetType() *string {
	return s.Type
}

func (s *DescribeAppResponseBodyResultDataSources) SetFields(v []map[string]interface{}) *DescribeAppResponseBodyResultDataSources {
	s.Fields = v
	return s
}

func (s *DescribeAppResponseBodyResultDataSources) SetKeyField(v string) *DescribeAppResponseBodyResultDataSources {
	s.KeyField = &v
	return s
}

func (s *DescribeAppResponseBodyResultDataSources) SetParameters(v map[string]interface{}) *DescribeAppResponseBodyResultDataSources {
	s.Parameters = v
	return s
}

func (s *DescribeAppResponseBodyResultDataSources) SetPlugins(v map[string]interface{}) *DescribeAppResponseBodyResultDataSources {
	s.Plugins = v
	return s
}

func (s *DescribeAppResponseBodyResultDataSources) SetSchemaName(v string) *DescribeAppResponseBodyResultDataSources {
	s.SchemaName = &v
	return s
}

func (s *DescribeAppResponseBodyResultDataSources) SetTableName(v string) *DescribeAppResponseBodyResultDataSources {
	s.TableName = &v
	return s
}

func (s *DescribeAppResponseBodyResultDataSources) SetType(v string) *DescribeAppResponseBodyResultDataSources {
	s.Type = &v
	return s
}

func (s *DescribeAppResponseBodyResultDataSources) Validate() error {
	return dara.Validate(s)
}

type DescribeAppResponseBodyResultDomain struct {
	// The type of the edition. Valid values: standard, advance, and enhanced. A value of standard indicates a standard edition. A value of advance indicates an advanced edition which is of an old version. New version is not supported for this edition. A value of enhanced indicates an advanced edition which is of a new version.
	//
	// example:
	//
	// -
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The search results.
	//
	// example:
	//
	// {}
	Functions *DescribeAppResponseBodyResultDomainFunctions `json:"functions,omitempty" xml:"functions,omitempty" type:"Struct"`
	// The name (in English).
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeAppResponseBodyResultDomain) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultDomain) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultDomain) GetCategory() *string {
	return s.Category
}

func (s *DescribeAppResponseBodyResultDomain) GetFunctions() *DescribeAppResponseBodyResultDomainFunctions {
	return s.Functions
}

func (s *DescribeAppResponseBodyResultDomain) GetName() *string {
	return s.Name
}

func (s *DescribeAppResponseBodyResultDomain) SetCategory(v string) *DescribeAppResponseBodyResultDomain {
	s.Category = &v
	return s
}

func (s *DescribeAppResponseBodyResultDomain) SetFunctions(v *DescribeAppResponseBodyResultDomainFunctions) *DescribeAppResponseBodyResultDomain {
	s.Functions = v
	return s
}

func (s *DescribeAppResponseBodyResultDomain) SetName(v string) *DescribeAppResponseBodyResultDomain {
	s.Name = &v
	return s
}

func (s *DescribeAppResponseBodyResultDomain) Validate() error {
	if s.Functions != nil {
		if err := s.Functions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAppResponseBodyResultDomainFunctions struct {
	// The structure of the algorithm.
	//
	// example:
	//
	// []
	Algo []*string `json:"algo,omitempty" xml:"algo,omitempty" type:"Repeated"`
	// The information about the layout.
	//
	// example:
	//
	// []
	Qp []*string `json:"qp,omitempty" xml:"qp,omitempty" type:"Repeated"`
	// The description of each feature.
	//
	// example:
	//
	// []
	Service []*string `json:"service,omitempty" xml:"service,omitempty" type:"Repeated"`
}

func (s DescribeAppResponseBodyResultDomainFunctions) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultDomainFunctions) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultDomainFunctions) GetAlgo() []*string {
	return s.Algo
}

func (s *DescribeAppResponseBodyResultDomainFunctions) GetQp() []*string {
	return s.Qp
}

func (s *DescribeAppResponseBodyResultDomainFunctions) GetService() []*string {
	return s.Service
}

func (s *DescribeAppResponseBodyResultDomainFunctions) SetAlgo(v []*string) *DescribeAppResponseBodyResultDomainFunctions {
	s.Algo = v
	return s
}

func (s *DescribeAppResponseBodyResultDomainFunctions) SetQp(v []*string) *DescribeAppResponseBodyResultDomainFunctions {
	s.Qp = v
	return s
}

func (s *DescribeAppResponseBodyResultDomainFunctions) SetService(v []*string) *DescribeAppResponseBodyResultDomainFunctions {
	s.Service = v
	return s
}

func (s *DescribeAppResponseBodyResultDomainFunctions) Validate() error {
	return dara.Validate(s)
}

type DescribeAppResponseBodyResultFirstRanks struct {
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
	// abc
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
	// STRUCT: The content of the expression is a structure. STRING (default): custom formula.
	//
	// example:
	//
	// STRING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeAppResponseBodyResultFirstRanks) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultFirstRanks) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultFirstRanks) GetActive() *bool {
	return s.Active
}

func (s *DescribeAppResponseBodyResultFirstRanks) GetDescription() *string {
	return s.Description
}

func (s *DescribeAppResponseBodyResultFirstRanks) GetMeta() interface{} {
	return s.Meta
}

func (s *DescribeAppResponseBodyResultFirstRanks) GetName() *string {
	return s.Name
}

func (s *DescribeAppResponseBodyResultFirstRanks) GetType() *string {
	return s.Type
}

func (s *DescribeAppResponseBodyResultFirstRanks) SetActive(v bool) *DescribeAppResponseBodyResultFirstRanks {
	s.Active = &v
	return s
}

func (s *DescribeAppResponseBodyResultFirstRanks) SetDescription(v string) *DescribeAppResponseBodyResultFirstRanks {
	s.Description = &v
	return s
}

func (s *DescribeAppResponseBodyResultFirstRanks) SetMeta(v interface{}) *DescribeAppResponseBodyResultFirstRanks {
	s.Meta = v
	return s
}

func (s *DescribeAppResponseBodyResultFirstRanks) SetName(v string) *DescribeAppResponseBodyResultFirstRanks {
	s.Name = &v
	return s
}

func (s *DescribeAppResponseBodyResultFirstRanks) SetType(v string) *DescribeAppResponseBodyResultFirstRanks {
	s.Type = &v
	return s
}

func (s *DescribeAppResponseBodyResultFirstRanks) Validate() error {
	return dara.Validate(s)
}

type DescribeAppResponseBodyResultQueryProcessors struct {
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
	// Then index range.
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

func (s DescribeAppResponseBodyResultQueryProcessors) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultQueryProcessors) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultQueryProcessors) GetActive() *bool {
	return s.Active
}

func (s *DescribeAppResponseBodyResultQueryProcessors) GetCategory() *string {
	return s.Category
}

func (s *DescribeAppResponseBodyResultQueryProcessors) GetDomain() *string {
	return s.Domain
}

func (s *DescribeAppResponseBodyResultQueryProcessors) GetIndexes() []*string {
	return s.Indexes
}

func (s *DescribeAppResponseBodyResultQueryProcessors) GetName() *string {
	return s.Name
}

func (s *DescribeAppResponseBodyResultQueryProcessors) GetProcessors() []map[string]interface{} {
	return s.Processors
}

func (s *DescribeAppResponseBodyResultQueryProcessors) SetActive(v bool) *DescribeAppResponseBodyResultQueryProcessors {
	s.Active = &v
	return s
}

func (s *DescribeAppResponseBodyResultQueryProcessors) SetCategory(v string) *DescribeAppResponseBodyResultQueryProcessors {
	s.Category = &v
	return s
}

func (s *DescribeAppResponseBodyResultQueryProcessors) SetDomain(v string) *DescribeAppResponseBodyResultQueryProcessors {
	s.Domain = &v
	return s
}

func (s *DescribeAppResponseBodyResultQueryProcessors) SetIndexes(v []*string) *DescribeAppResponseBodyResultQueryProcessors {
	s.Indexes = v
	return s
}

func (s *DescribeAppResponseBodyResultQueryProcessors) SetName(v string) *DescribeAppResponseBodyResultQueryProcessors {
	s.Name = &v
	return s
}

func (s *DescribeAppResponseBodyResultQueryProcessors) SetProcessors(v []map[string]interface{}) *DescribeAppResponseBodyResultQueryProcessors {
	s.Processors = v
	return s
}

func (s *DescribeAppResponseBodyResultQueryProcessors) Validate() error {
	return dara.Validate(s)
}

type DescribeAppResponseBodyResultQuota struct {
	// The computing resources. Unit: logical computing unit (LCU).
	//
	// example:
	//
	// 20
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity. Unit: GB.
	//
	// example:
	//
	// 1
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The number of search requests per second. You are charged based on the number of search requests per second in the earlier billing model.
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
}

func (s DescribeAppResponseBodyResultQuota) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultQuota) GetComputeResource() *int32 {
	return s.ComputeResource
}

func (s *DescribeAppResponseBodyResultQuota) GetDocSize() *int32 {
	return s.DocSize
}

func (s *DescribeAppResponseBodyResultQuota) GetQps() *int32 {
	return s.Qps
}

func (s *DescribeAppResponseBodyResultQuota) GetSpec() *string {
	return s.Spec
}

func (s *DescribeAppResponseBodyResultQuota) SetComputeResource(v int32) *DescribeAppResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *DescribeAppResponseBodyResultQuota) SetDocSize(v int32) *DescribeAppResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *DescribeAppResponseBodyResultQuota) SetQps(v int32) *DescribeAppResponseBodyResultQuota {
	s.Qps = &v
	return s
}

func (s *DescribeAppResponseBodyResultQuota) SetSpec(v string) *DescribeAppResponseBodyResultQuota {
	s.Spec = &v
	return s
}

func (s *DescribeAppResponseBodyResultQuota) Validate() error {
	return dara.Validate(s)
}

type DescribeAppResponseBodyResultSchema struct {
	// The sort configurations.
	IndexSortConfig []*DescribeAppResponseBodyResultSchemaIndexSortConfig `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	// The index schema.
	Indexes *DescribeAppResponseBodyResultSchemaIndexes `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
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
	// The hot values of the level-1 routing field.
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
	TtlField *DescribeAppResponseBodyResultSchemaTtlField `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s DescribeAppResponseBodyResultSchema) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultSchema) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSchema) GetIndexSortConfig() []*DescribeAppResponseBodyResultSchemaIndexSortConfig {
	return s.IndexSortConfig
}

func (s *DescribeAppResponseBodyResultSchema) GetIndexes() *DescribeAppResponseBodyResultSchemaIndexes {
	return s.Indexes
}

func (s *DescribeAppResponseBodyResultSchema) GetName() *string {
	return s.Name
}

func (s *DescribeAppResponseBodyResultSchema) GetRouteField() *string {
	return s.RouteField
}

func (s *DescribeAppResponseBodyResultSchema) GetRouteFieldValues() []*string {
	return s.RouteFieldValues
}

func (s *DescribeAppResponseBodyResultSchema) GetSecondRouteField() *string {
	return s.SecondRouteField
}

func (s *DescribeAppResponseBodyResultSchema) GetTables() map[string]interface{} {
	return s.Tables
}

func (s *DescribeAppResponseBodyResultSchema) GetTtlField() *DescribeAppResponseBodyResultSchemaTtlField {
	return s.TtlField
}

func (s *DescribeAppResponseBodyResultSchema) SetIndexSortConfig(v []*DescribeAppResponseBodyResultSchemaIndexSortConfig) *DescribeAppResponseBodyResultSchema {
	s.IndexSortConfig = v
	return s
}

func (s *DescribeAppResponseBodyResultSchema) SetIndexes(v *DescribeAppResponseBodyResultSchemaIndexes) *DescribeAppResponseBodyResultSchema {
	s.Indexes = v
	return s
}

func (s *DescribeAppResponseBodyResultSchema) SetName(v string) *DescribeAppResponseBodyResultSchema {
	s.Name = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchema) SetRouteField(v string) *DescribeAppResponseBodyResultSchema {
	s.RouteField = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchema) SetRouteFieldValues(v []*string) *DescribeAppResponseBodyResultSchema {
	s.RouteFieldValues = v
	return s
}

func (s *DescribeAppResponseBodyResultSchema) SetSecondRouteField(v string) *DescribeAppResponseBodyResultSchema {
	s.SecondRouteField = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchema) SetTables(v map[string]interface{}) *DescribeAppResponseBodyResultSchema {
	s.Tables = v
	return s
}

func (s *DescribeAppResponseBodyResultSchema) SetTtlField(v *DescribeAppResponseBodyResultSchemaTtlField) *DescribeAppResponseBodyResultSchema {
	s.TtlField = v
	return s
}

func (s *DescribeAppResponseBodyResultSchema) Validate() error {
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

type DescribeAppResponseBodyResultSchemaIndexSortConfig struct {
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

func (s DescribeAppResponseBodyResultSchemaIndexSortConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultSchemaIndexSortConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSchemaIndexSortConfig) GetDirection() *string {
	return s.Direction
}

func (s *DescribeAppResponseBodyResultSchemaIndexSortConfig) GetField() *string {
	return s.Field
}

func (s *DescribeAppResponseBodyResultSchemaIndexSortConfig) SetDirection(v string) *DescribeAppResponseBodyResultSchemaIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemaIndexSortConfig) SetField(v string) *DescribeAppResponseBodyResultSchemaIndexSortConfig {
	s.Field = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemaIndexSortConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeAppResponseBodyResultSchemaIndexes struct {
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

func (s DescribeAppResponseBodyResultSchemaIndexes) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultSchemaIndexes) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSchemaIndexes) GetFilterFields() []*string {
	return s.FilterFields
}

func (s *DescribeAppResponseBodyResultSchemaIndexes) GetSearchFields() map[string]interface{} {
	return s.SearchFields
}

func (s *DescribeAppResponseBodyResultSchemaIndexes) SetFilterFields(v []*string) *DescribeAppResponseBodyResultSchemaIndexes {
	s.FilterFields = v
	return s
}

func (s *DescribeAppResponseBodyResultSchemaIndexes) SetSearchFields(v map[string]interface{}) *DescribeAppResponseBodyResultSchemaIndexes {
	s.SearchFields = v
	return s
}

func (s *DescribeAppResponseBodyResultSchemaIndexes) Validate() error {
	return dara.Validate(s)
}

type DescribeAppResponseBodyResultSchemaTtlField struct {
	// The document time field.
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

func (s DescribeAppResponseBodyResultSchemaTtlField) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultSchemaTtlField) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSchemaTtlField) GetName() *string {
	return s.Name
}

func (s *DescribeAppResponseBodyResultSchemaTtlField) GetTtl() *int64 {
	return s.Ttl
}

func (s *DescribeAppResponseBodyResultSchemaTtlField) SetName(v string) *DescribeAppResponseBodyResultSchemaTtlField {
	s.Name = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemaTtlField) SetTtl(v int64) *DescribeAppResponseBodyResultSchemaTtlField {
	s.Ttl = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemaTtlField) Validate() error {
	return dara.Validate(s)
}

type DescribeAppResponseBodyResultSchemas struct {
	// The sort configurations.
	IndexSortConfig []*DescribeAppResponseBodyResultSchemasIndexSortConfig `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	// The index schema.
	Indexes *DescribeAppResponseBodyResultSchemasIndexes `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
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
	// The hot values of the level-1 routing field.
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
	TtlField *DescribeAppResponseBodyResultSchemasTtlField `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s DescribeAppResponseBodyResultSchemas) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultSchemas) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSchemas) GetIndexSortConfig() []*DescribeAppResponseBodyResultSchemasIndexSortConfig {
	return s.IndexSortConfig
}

func (s *DescribeAppResponseBodyResultSchemas) GetIndexes() *DescribeAppResponseBodyResultSchemasIndexes {
	return s.Indexes
}

func (s *DescribeAppResponseBodyResultSchemas) GetName() *string {
	return s.Name
}

func (s *DescribeAppResponseBodyResultSchemas) GetRouteField() *string {
	return s.RouteField
}

func (s *DescribeAppResponseBodyResultSchemas) GetRouteFieldValues() []*string {
	return s.RouteFieldValues
}

func (s *DescribeAppResponseBodyResultSchemas) GetSecondRouteField() *string {
	return s.SecondRouteField
}

func (s *DescribeAppResponseBodyResultSchemas) GetTables() map[string]interface{} {
	return s.Tables
}

func (s *DescribeAppResponseBodyResultSchemas) GetTtlField() *DescribeAppResponseBodyResultSchemasTtlField {
	return s.TtlField
}

func (s *DescribeAppResponseBodyResultSchemas) SetIndexSortConfig(v []*DescribeAppResponseBodyResultSchemasIndexSortConfig) *DescribeAppResponseBodyResultSchemas {
	s.IndexSortConfig = v
	return s
}

func (s *DescribeAppResponseBodyResultSchemas) SetIndexes(v *DescribeAppResponseBodyResultSchemasIndexes) *DescribeAppResponseBodyResultSchemas {
	s.Indexes = v
	return s
}

func (s *DescribeAppResponseBodyResultSchemas) SetName(v string) *DescribeAppResponseBodyResultSchemas {
	s.Name = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemas) SetRouteField(v string) *DescribeAppResponseBodyResultSchemas {
	s.RouteField = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemas) SetRouteFieldValues(v []*string) *DescribeAppResponseBodyResultSchemas {
	s.RouteFieldValues = v
	return s
}

func (s *DescribeAppResponseBodyResultSchemas) SetSecondRouteField(v string) *DescribeAppResponseBodyResultSchemas {
	s.SecondRouteField = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemas) SetTables(v map[string]interface{}) *DescribeAppResponseBodyResultSchemas {
	s.Tables = v
	return s
}

func (s *DescribeAppResponseBodyResultSchemas) SetTtlField(v *DescribeAppResponseBodyResultSchemasTtlField) *DescribeAppResponseBodyResultSchemas {
	s.TtlField = v
	return s
}

func (s *DescribeAppResponseBodyResultSchemas) Validate() error {
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

type DescribeAppResponseBodyResultSchemasIndexSortConfig struct {
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

func (s DescribeAppResponseBodyResultSchemasIndexSortConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultSchemasIndexSortConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSchemasIndexSortConfig) GetDirection() *string {
	return s.Direction
}

func (s *DescribeAppResponseBodyResultSchemasIndexSortConfig) GetField() *string {
	return s.Field
}

func (s *DescribeAppResponseBodyResultSchemasIndexSortConfig) SetDirection(v string) *DescribeAppResponseBodyResultSchemasIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemasIndexSortConfig) SetField(v string) *DescribeAppResponseBodyResultSchemasIndexSortConfig {
	s.Field = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemasIndexSortConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeAppResponseBodyResultSchemasIndexes struct {
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

func (s DescribeAppResponseBodyResultSchemasIndexes) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultSchemasIndexes) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSchemasIndexes) GetFilterFields() []*string {
	return s.FilterFields
}

func (s *DescribeAppResponseBodyResultSchemasIndexes) GetSearchFields() map[string]interface{} {
	return s.SearchFields
}

func (s *DescribeAppResponseBodyResultSchemasIndexes) SetFilterFields(v []*string) *DescribeAppResponseBodyResultSchemasIndexes {
	s.FilterFields = v
	return s
}

func (s *DescribeAppResponseBodyResultSchemasIndexes) SetSearchFields(v map[string]interface{}) *DescribeAppResponseBodyResultSchemasIndexes {
	s.SearchFields = v
	return s
}

func (s *DescribeAppResponseBodyResultSchemasIndexes) Validate() error {
	return dara.Validate(s)
}

type DescribeAppResponseBodyResultSchemasTtlField struct {
	// The document time field.
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

func (s DescribeAppResponseBodyResultSchemasTtlField) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultSchemasTtlField) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSchemasTtlField) GetName() *string {
	return s.Name
}

func (s *DescribeAppResponseBodyResultSchemasTtlField) GetTtl() *int64 {
	return s.Ttl
}

func (s *DescribeAppResponseBodyResultSchemasTtlField) SetName(v string) *DescribeAppResponseBodyResultSchemasTtlField {
	s.Name = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemasTtlField) SetTtl(v int64) *DescribeAppResponseBodyResultSchemasTtlField {
	s.Ttl = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemasTtlField) Validate() error {
	return dara.Validate(s)
}

type DescribeAppResponseBodyResultSecondRanks struct {
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

func (s DescribeAppResponseBodyResultSecondRanks) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultSecondRanks) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSecondRanks) GetActive() *bool {
	return s.Active
}

func (s *DescribeAppResponseBodyResultSecondRanks) GetDescription() *string {
	return s.Description
}

func (s *DescribeAppResponseBodyResultSecondRanks) GetMeta() interface{} {
	return s.Meta
}

func (s *DescribeAppResponseBodyResultSecondRanks) GetName() *string {
	return s.Name
}

func (s *DescribeAppResponseBodyResultSecondRanks) SetActive(v bool) *DescribeAppResponseBodyResultSecondRanks {
	s.Active = &v
	return s
}

func (s *DescribeAppResponseBodyResultSecondRanks) SetDescription(v string) *DescribeAppResponseBodyResultSecondRanks {
	s.Description = &v
	return s
}

func (s *DescribeAppResponseBodyResultSecondRanks) SetMeta(v interface{}) *DescribeAppResponseBodyResultSecondRanks {
	s.Meta = v
	return s
}

func (s *DescribeAppResponseBodyResultSecondRanks) SetName(v string) *DescribeAppResponseBodyResultSecondRanks {
	s.Name = &v
	return s
}

func (s *DescribeAppResponseBodyResultSecondRanks) Validate() error {
	return dara.Validate(s)
}

type DescribeAppResponseBodyResultSummaries struct {
	// The summary configurations.
	Meta []*DescribeAppResponseBodyResultSummariesMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The group name.
	//
	// example:
	//
	// default
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeAppResponseBodyResultSummaries) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultSummaries) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSummaries) GetMeta() []*DescribeAppResponseBodyResultSummariesMeta {
	return s.Meta
}

func (s *DescribeAppResponseBodyResultSummaries) GetName() *string {
	return s.Name
}

func (s *DescribeAppResponseBodyResultSummaries) SetMeta(v []*DescribeAppResponseBodyResultSummariesMeta) *DescribeAppResponseBodyResultSummaries {
	s.Meta = v
	return s
}

func (s *DescribeAppResponseBodyResultSummaries) SetName(v string) *DescribeAppResponseBodyResultSummaries {
	s.Name = &v
	return s
}

func (s *DescribeAppResponseBodyResultSummaries) Validate() error {
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

type DescribeAppResponseBodyResultSummariesMeta struct {
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

func (s DescribeAppResponseBodyResultSummariesMeta) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBodyResultSummariesMeta) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSummariesMeta) GetElement() *string {
	return s.Element
}

func (s *DescribeAppResponseBodyResultSummariesMeta) GetEllipsis() *string {
	return s.Ellipsis
}

func (s *DescribeAppResponseBodyResultSummariesMeta) GetField() *string {
	return s.Field
}

func (s *DescribeAppResponseBodyResultSummariesMeta) GetLen() *int32 {
	return s.Len
}

func (s *DescribeAppResponseBodyResultSummariesMeta) GetSnippet() *string {
	return s.Snippet
}

func (s *DescribeAppResponseBodyResultSummariesMeta) SetElement(v string) *DescribeAppResponseBodyResultSummariesMeta {
	s.Element = &v
	return s
}

func (s *DescribeAppResponseBodyResultSummariesMeta) SetEllipsis(v string) *DescribeAppResponseBodyResultSummariesMeta {
	s.Ellipsis = &v
	return s
}

func (s *DescribeAppResponseBodyResultSummariesMeta) SetField(v string) *DescribeAppResponseBodyResultSummariesMeta {
	s.Field = &v
	return s
}

func (s *DescribeAppResponseBodyResultSummariesMeta) SetLen(v int32) *DescribeAppResponseBodyResultSummariesMeta {
	s.Len = &v
	return s
}

func (s *DescribeAppResponseBodyResultSummariesMeta) SetSnippet(v string) *DescribeAppResponseBodyResultSummariesMeta {
	s.Snippet = &v
	return s
}

func (s *DescribeAppResponseBodyResultSummariesMeta) Validate() error {
	return dara.Validate(s)
}
