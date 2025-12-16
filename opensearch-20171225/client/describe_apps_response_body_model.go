// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAppsResponseBody
	GetRequestId() *string
	SetResult(v []*DescribeAppsResponseBodyResult) *DescribeAppsResponseBody
	GetResult() []*DescribeAppsResponseBodyResult
}

type DescribeAppsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 77CAA411-0010-4DB9-82E2-1C384E590AFF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The response parameters.
	Result []*DescribeAppsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s DescribeAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppsResponseBody) GetResult() []*DescribeAppsResponseBodyResult {
	return s.Result
}

func (s *DescribeAppsResponseBody) SetRequestId(v string) *DescribeAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppsResponseBody) SetResult(v []*DescribeAppsResponseBodyResult) *DescribeAppsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAppsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAppsResponseBodyResult struct {
	// Indicates whether the version is automatically switched to an online version.
	//
	// example:
	//
	// true
	AutoSwitch *bool `json:"autoSwitch,omitempty" xml:"autoSwitch,omitempty"`
	// The capability opening configurations.
	Cluster *DescribeAppsResponseBodyResultCluster `json:"cluster,omitempty" xml:"cluster,omitempty" type:"Struct"`
	// The cluster name.
	//
	// example:
	//
	// vpc_sh_domain_1
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	// The configurations of data sources.
	DataSources []*DescribeAppsResponseBodyResultDataSources `json:"dataSources,omitempty" xml:"dataSources,omitempty" type:"Repeated"`
	// The description.
	//
	// example:
	//
	// ""
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The type of the industry. Valid values:
	//
	// 	- GENERAL
	//
	// 	- ECOMMERCE
	//
	// 	- IT_CONTENT
	Domain *DescribeAppsResponseBodyResultDomain `json:"domain,omitempty" xml:"domain,omitempty" type:"Struct"`
	// The default display fields.
	FetchFields []*string `json:"fetchFields,omitempty" xml:"fetchFields,omitempty" type:"Repeated"`
	// The configurations of rough sort.
	FirstRanks []*DescribeAppsResponseBodyResultFirstRanks `json:"firstRanks,omitempty" xml:"firstRanks,omitempty" type:"Repeated"`
	// The group ID.
	//
	// example:
	//
	// 100302881
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The field model.
	Interpretations map[string]interface{} `json:"interpretations,omitempty" xml:"interpretations,omitempty"`
	// Indicates whether the version is an online version.
	//
	// example:
	//
	// 12333
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
	QueryProcessors []*DescribeAppsResponseBodyResultQueryProcessors `json:"queryProcessors,omitempty" xml:"queryProcessors,omitempty" type:"Repeated"`
	// The quota information.
	Quota *DescribeAppsResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The application schema.
	Schema *DescribeAppsResponseBodyResultSchema `json:"schema,omitempty" xml:"schema,omitempty" type:"Struct"`
	// The single-table schema.
	Schemas []*DescribeAppsResponseBodyResultSchemas `json:"schemas,omitempty" xml:"schemas,omitempty" type:"Repeated"`
	// The configurations of fine sort.
	SecondRanks []*DescribeAppsResponseBodyResultSecondRanks `json:"secondRanks,omitempty" xml:"secondRanks,omitempty" type:"Repeated"`
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
	// normal
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The search result summary configurations.
	Summaries []*DescribeAppsResponseBodyResultSummaries `json:"summaries,omitempty" xml:"summaries,omitempty" type:"Repeated"`
	// The type of the application. Valid values:
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

func (s DescribeAppsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResult) GetAutoSwitch() *bool {
	return s.AutoSwitch
}

func (s *DescribeAppsResponseBodyResult) GetCluster() *DescribeAppsResponseBodyResultCluster {
	return s.Cluster
}

func (s *DescribeAppsResponseBodyResult) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeAppsResponseBodyResult) GetDataSources() []*DescribeAppsResponseBodyResultDataSources {
	return s.DataSources
}

func (s *DescribeAppsResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeAppsResponseBodyResult) GetDomain() *DescribeAppsResponseBodyResultDomain {
	return s.Domain
}

func (s *DescribeAppsResponseBodyResult) GetFetchFields() []*string {
	return s.FetchFields
}

func (s *DescribeAppsResponseBodyResult) GetFirstRanks() []*DescribeAppsResponseBodyResultFirstRanks {
	return s.FirstRanks
}

func (s *DescribeAppsResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *DescribeAppsResponseBodyResult) GetInterpretations() map[string]interface{} {
	return s.Interpretations
}

func (s *DescribeAppsResponseBodyResult) GetIsCurrent() *bool {
	return s.IsCurrent
}

func (s *DescribeAppsResponseBodyResult) GetProgressPercent() *int32 {
	return s.ProgressPercent
}

func (s *DescribeAppsResponseBodyResult) GetPrompts() []map[string]interface{} {
	return s.Prompts
}

func (s *DescribeAppsResponseBodyResult) GetQueryProcessors() []*DescribeAppsResponseBodyResultQueryProcessors {
	return s.QueryProcessors
}

func (s *DescribeAppsResponseBodyResult) GetQuota() *DescribeAppsResponseBodyResultQuota {
	return s.Quota
}

func (s *DescribeAppsResponseBodyResult) GetSchema() *DescribeAppsResponseBodyResultSchema {
	return s.Schema
}

func (s *DescribeAppsResponseBodyResult) GetSchemas() []*DescribeAppsResponseBodyResultSchemas {
	return s.Schemas
}

func (s *DescribeAppsResponseBodyResult) GetSecondRanks() []*DescribeAppsResponseBodyResultSecondRanks {
	return s.SecondRanks
}

func (s *DescribeAppsResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *DescribeAppsResponseBodyResult) GetSummaries() []*DescribeAppsResponseBodyResultSummaries {
	return s.Summaries
}

func (s *DescribeAppsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *DescribeAppsResponseBodyResult) SetAutoSwitch(v bool) *DescribeAppsResponseBodyResult {
	s.AutoSwitch = &v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetCluster(v *DescribeAppsResponseBodyResultCluster) *DescribeAppsResponseBodyResult {
	s.Cluster = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetClusterName(v string) *DescribeAppsResponseBodyResult {
	s.ClusterName = &v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetDataSources(v []*DescribeAppsResponseBodyResultDataSources) *DescribeAppsResponseBodyResult {
	s.DataSources = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetDescription(v string) *DescribeAppsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetDomain(v *DescribeAppsResponseBodyResultDomain) *DescribeAppsResponseBodyResult {
	s.Domain = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetFetchFields(v []*string) *DescribeAppsResponseBodyResult {
	s.FetchFields = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetFirstRanks(v []*DescribeAppsResponseBodyResultFirstRanks) *DescribeAppsResponseBodyResult {
	s.FirstRanks = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetId(v string) *DescribeAppsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetInterpretations(v map[string]interface{}) *DescribeAppsResponseBodyResult {
	s.Interpretations = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetIsCurrent(v bool) *DescribeAppsResponseBodyResult {
	s.IsCurrent = &v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetProgressPercent(v int32) *DescribeAppsResponseBodyResult {
	s.ProgressPercent = &v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetPrompts(v []map[string]interface{}) *DescribeAppsResponseBodyResult {
	s.Prompts = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetQueryProcessors(v []*DescribeAppsResponseBodyResultQueryProcessors) *DescribeAppsResponseBodyResult {
	s.QueryProcessors = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetQuota(v *DescribeAppsResponseBodyResultQuota) *DescribeAppsResponseBodyResult {
	s.Quota = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetSchema(v *DescribeAppsResponseBodyResultSchema) *DescribeAppsResponseBodyResult {
	s.Schema = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetSchemas(v []*DescribeAppsResponseBodyResultSchemas) *DescribeAppsResponseBodyResult {
	s.Schemas = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetSecondRanks(v []*DescribeAppsResponseBodyResultSecondRanks) *DescribeAppsResponseBodyResult {
	s.SecondRanks = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetStatus(v string) *DescribeAppsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetSummaries(v []*DescribeAppsResponseBodyResultSummaries) *DescribeAppsResponseBodyResult {
	s.Summaries = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetType(v string) *DescribeAppsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *DescribeAppsResponseBodyResult) Validate() error {
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

type DescribeAppsResponseBodyResultCluster struct {
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

func (s DescribeAppsResponseBodyResultCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultCluster) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultCluster) GetMaxQueryClauseLength() *int32 {
	return s.MaxQueryClauseLength
}

func (s *DescribeAppsResponseBodyResultCluster) GetMaxTimeoutMS() *int32 {
	return s.MaxTimeoutMS
}

func (s *DescribeAppsResponseBodyResultCluster) SetMaxQueryClauseLength(v int32) *DescribeAppsResponseBodyResultCluster {
	s.MaxQueryClauseLength = &v
	return s
}

func (s *DescribeAppsResponseBodyResultCluster) SetMaxTimeoutMS(v int32) *DescribeAppsResponseBodyResultCluster {
	s.MaxTimeoutMS = &v
	return s
}

func (s *DescribeAppsResponseBodyResultCluster) Validate() error {
	return dara.Validate(s)
}

type DescribeAppsResponseBodyResultDataSources struct {
	// The information about field mappings.
	Fields []map[string]interface{} `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// The primary key.
	//
	// example:
	//
	// id
	KeyField *string `json:"keyField,omitempty" xml:"keyField,omitempty"`
	// The information about the data source.
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

func (s DescribeAppsResponseBodyResultDataSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultDataSources) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultDataSources) GetFields() []map[string]interface{} {
	return s.Fields
}

func (s *DescribeAppsResponseBodyResultDataSources) GetKeyField() *string {
	return s.KeyField
}

func (s *DescribeAppsResponseBodyResultDataSources) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *DescribeAppsResponseBodyResultDataSources) GetPlugins() map[string]interface{} {
	return s.Plugins
}

func (s *DescribeAppsResponseBodyResultDataSources) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeAppsResponseBodyResultDataSources) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAppsResponseBodyResultDataSources) GetType() *string {
	return s.Type
}

func (s *DescribeAppsResponseBodyResultDataSources) SetFields(v []map[string]interface{}) *DescribeAppsResponseBodyResultDataSources {
	s.Fields = v
	return s
}

func (s *DescribeAppsResponseBodyResultDataSources) SetKeyField(v string) *DescribeAppsResponseBodyResultDataSources {
	s.KeyField = &v
	return s
}

func (s *DescribeAppsResponseBodyResultDataSources) SetParameters(v map[string]interface{}) *DescribeAppsResponseBodyResultDataSources {
	s.Parameters = v
	return s
}

func (s *DescribeAppsResponseBodyResultDataSources) SetPlugins(v map[string]interface{}) *DescribeAppsResponseBodyResultDataSources {
	s.Plugins = v
	return s
}

func (s *DescribeAppsResponseBodyResultDataSources) SetSchemaName(v string) *DescribeAppsResponseBodyResultDataSources {
	s.SchemaName = &v
	return s
}

func (s *DescribeAppsResponseBodyResultDataSources) SetTableName(v string) *DescribeAppsResponseBodyResultDataSources {
	s.TableName = &v
	return s
}

func (s *DescribeAppsResponseBodyResultDataSources) SetType(v string) *DescribeAppsResponseBodyResultDataSources {
	s.Type = &v
	return s
}

func (s *DescribeAppsResponseBodyResultDataSources) Validate() error {
	return dara.Validate(s)
}

type DescribeAppsResponseBodyResultDomain struct {
	// The type of the edition. Valid values: standard, advance, and enhanced. A value of standard indicates a standard edition. A value of advance indicates an advanced edition which is of an old version. New version is not supported for this edition. A value of enhanced indicates an advanced edition which is of a new version.
	//
	// example:
	//
	// -
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The search results.
	Functions *DescribeAppsResponseBodyResultDomainFunctions `json:"functions,omitempty" xml:"functions,omitempty" type:"Struct"`
	// The name (in English).
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeAppsResponseBodyResultDomain) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultDomain) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultDomain) GetCategory() *string {
	return s.Category
}

func (s *DescribeAppsResponseBodyResultDomain) GetFunctions() *DescribeAppsResponseBodyResultDomainFunctions {
	return s.Functions
}

func (s *DescribeAppsResponseBodyResultDomain) GetName() *string {
	return s.Name
}

func (s *DescribeAppsResponseBodyResultDomain) SetCategory(v string) *DescribeAppsResponseBodyResultDomain {
	s.Category = &v
	return s
}

func (s *DescribeAppsResponseBodyResultDomain) SetFunctions(v *DescribeAppsResponseBodyResultDomainFunctions) *DescribeAppsResponseBodyResultDomain {
	s.Functions = v
	return s
}

func (s *DescribeAppsResponseBodyResultDomain) SetName(v string) *DescribeAppsResponseBodyResultDomain {
	s.Name = &v
	return s
}

func (s *DescribeAppsResponseBodyResultDomain) Validate() error {
	if s.Functions != nil {
		if err := s.Functions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAppsResponseBodyResultDomainFunctions struct {
	// Structure 1.
	Algo []*string `json:"algo,omitempty" xml:"algo,omitempty" type:"Repeated"`
	// Information 1.
	Qp []*string `json:"qp,omitempty" xml:"qp,omitempty" type:"Repeated"`
	// Feature 1.
	Service []*string `json:"service,omitempty" xml:"service,omitempty" type:"Repeated"`
}

func (s DescribeAppsResponseBodyResultDomainFunctions) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultDomainFunctions) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultDomainFunctions) GetAlgo() []*string {
	return s.Algo
}

func (s *DescribeAppsResponseBodyResultDomainFunctions) GetQp() []*string {
	return s.Qp
}

func (s *DescribeAppsResponseBodyResultDomainFunctions) GetService() []*string {
	return s.Service
}

func (s *DescribeAppsResponseBodyResultDomainFunctions) SetAlgo(v []*string) *DescribeAppsResponseBodyResultDomainFunctions {
	s.Algo = v
	return s
}

func (s *DescribeAppsResponseBodyResultDomainFunctions) SetQp(v []*string) *DescribeAppsResponseBodyResultDomainFunctions {
	s.Qp = v
	return s
}

func (s *DescribeAppsResponseBodyResultDomainFunctions) SetService(v []*string) *DescribeAppsResponseBodyResultDomainFunctions {
	s.Service = v
	return s
}

func (s *DescribeAppsResponseBodyResultDomainFunctions) Validate() error {
	return dara.Validate(s)
}

type DescribeAppsResponseBodyResultFirstRanks struct {
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
	// The information about the expression. The information is displayed in the array or string format.
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
	// STRUCT: The content of the expression is a structure. STRING (default): a custom formula.
	//
	// example:
	//
	// STRING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeAppsResponseBodyResultFirstRanks) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultFirstRanks) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultFirstRanks) GetActive() *bool {
	return s.Active
}

func (s *DescribeAppsResponseBodyResultFirstRanks) GetDescription() *string {
	return s.Description
}

func (s *DescribeAppsResponseBodyResultFirstRanks) GetMeta() interface{} {
	return s.Meta
}

func (s *DescribeAppsResponseBodyResultFirstRanks) GetName() *string {
	return s.Name
}

func (s *DescribeAppsResponseBodyResultFirstRanks) GetType() *string {
	return s.Type
}

func (s *DescribeAppsResponseBodyResultFirstRanks) SetActive(v bool) *DescribeAppsResponseBodyResultFirstRanks {
	s.Active = &v
	return s
}

func (s *DescribeAppsResponseBodyResultFirstRanks) SetDescription(v string) *DescribeAppsResponseBodyResultFirstRanks {
	s.Description = &v
	return s
}

func (s *DescribeAppsResponseBodyResultFirstRanks) SetMeta(v interface{}) *DescribeAppsResponseBodyResultFirstRanks {
	s.Meta = v
	return s
}

func (s *DescribeAppsResponseBodyResultFirstRanks) SetName(v string) *DescribeAppsResponseBodyResultFirstRanks {
	s.Name = &v
	return s
}

func (s *DescribeAppsResponseBodyResultFirstRanks) SetType(v string) *DescribeAppsResponseBodyResultFirstRanks {
	s.Type = &v
	return s
}

func (s *DescribeAppsResponseBodyResultFirstRanks) Validate() error {
	return dara.Validate(s)
}

type DescribeAppsResponseBodyResultQueryProcessors struct {
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
	// The type of the industry. Valid values:
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
	// The indexes.
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

func (s DescribeAppsResponseBodyResultQueryProcessors) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultQueryProcessors) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) GetActive() *bool {
	return s.Active
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) GetCategory() *string {
	return s.Category
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) GetDomain() *string {
	return s.Domain
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) GetIndexes() []*string {
	return s.Indexes
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) GetName() *string {
	return s.Name
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) GetProcessors() []map[string]interface{} {
	return s.Processors
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) SetActive(v bool) *DescribeAppsResponseBodyResultQueryProcessors {
	s.Active = &v
	return s
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) SetCategory(v string) *DescribeAppsResponseBodyResultQueryProcessors {
	s.Category = &v
	return s
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) SetDomain(v string) *DescribeAppsResponseBodyResultQueryProcessors {
	s.Domain = &v
	return s
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) SetIndexes(v []*string) *DescribeAppsResponseBodyResultQueryProcessors {
	s.Indexes = v
	return s
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) SetName(v string) *DescribeAppsResponseBodyResultQueryProcessors {
	s.Name = &v
	return s
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) SetProcessors(v []map[string]interface{}) *DescribeAppsResponseBodyResultQueryProcessors {
	s.Processors = v
	return s
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) Validate() error {
	return dara.Validate(s)
}

type DescribeAppsResponseBodyResultQuota struct {
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

func (s DescribeAppsResponseBodyResultQuota) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultQuota) GetComputeResource() *int32 {
	return s.ComputeResource
}

func (s *DescribeAppsResponseBodyResultQuota) GetDocSize() *int32 {
	return s.DocSize
}

func (s *DescribeAppsResponseBodyResultQuota) GetQps() *int32 {
	return s.Qps
}

func (s *DescribeAppsResponseBodyResultQuota) GetSpec() *string {
	return s.Spec
}

func (s *DescribeAppsResponseBodyResultQuota) SetComputeResource(v int32) *DescribeAppsResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *DescribeAppsResponseBodyResultQuota) SetDocSize(v int32) *DescribeAppsResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *DescribeAppsResponseBodyResultQuota) SetQps(v int32) *DescribeAppsResponseBodyResultQuota {
	s.Qps = &v
	return s
}

func (s *DescribeAppsResponseBodyResultQuota) SetSpec(v string) *DescribeAppsResponseBodyResultQuota {
	s.Spec = &v
	return s
}

func (s *DescribeAppsResponseBodyResultQuota) Validate() error {
	return dara.Validate(s)
}

type DescribeAppsResponseBodyResultSchema struct {
	// The sort configurations.
	IndexSortConfig []*DescribeAppsResponseBodyResultSchemaIndexSortConfig `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	// The index schema.
	Indexes *DescribeAppsResponseBodyResultSchemaIndexes `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
	// The name of the wide table.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The level-1 routing field.
	//
	// example:
	//
	// field1
	RouteField *string `json:"routeField,omitempty" xml:"routeField,omitempty"`
	// The hot values of the level-1 routing field.
	RouteFieldValues []*string `json:"routeFieldValues,omitempty" xml:"routeFieldValues,omitempty" type:"Repeated"`
	// The level-2 routing field. This parameter is returned if the routeFieldValues parameter is returned. By default, the wide-table primary key field is used as the level-2 routing field.
	//
	// example:
	//
	// field2
	SecondRouteField *string `json:"secondRouteField,omitempty" xml:"secondRouteField,omitempty"`
	// The table schema.
	Tables map[string]interface{} `json:"tables,omitempty" xml:"tables,omitempty"`
	// The document clearing configurations.
	TtlField *DescribeAppsResponseBodyResultSchemaTtlField `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s DescribeAppsResponseBodyResultSchema) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSchema) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSchema) GetIndexSortConfig() []*DescribeAppsResponseBodyResultSchemaIndexSortConfig {
	return s.IndexSortConfig
}

func (s *DescribeAppsResponseBodyResultSchema) GetIndexes() *DescribeAppsResponseBodyResultSchemaIndexes {
	return s.Indexes
}

func (s *DescribeAppsResponseBodyResultSchema) GetName() *string {
	return s.Name
}

func (s *DescribeAppsResponseBodyResultSchema) GetRouteField() *string {
	return s.RouteField
}

func (s *DescribeAppsResponseBodyResultSchema) GetRouteFieldValues() []*string {
	return s.RouteFieldValues
}

func (s *DescribeAppsResponseBodyResultSchema) GetSecondRouteField() *string {
	return s.SecondRouteField
}

func (s *DescribeAppsResponseBodyResultSchema) GetTables() map[string]interface{} {
	return s.Tables
}

func (s *DescribeAppsResponseBodyResultSchema) GetTtlField() *DescribeAppsResponseBodyResultSchemaTtlField {
	return s.TtlField
}

func (s *DescribeAppsResponseBodyResultSchema) SetIndexSortConfig(v []*DescribeAppsResponseBodyResultSchemaIndexSortConfig) *DescribeAppsResponseBodyResultSchema {
	s.IndexSortConfig = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchema) SetIndexes(v *DescribeAppsResponseBodyResultSchemaIndexes) *DescribeAppsResponseBodyResultSchema {
	s.Indexes = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchema) SetName(v string) *DescribeAppsResponseBodyResultSchema {
	s.Name = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchema) SetRouteField(v string) *DescribeAppsResponseBodyResultSchema {
	s.RouteField = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchema) SetRouteFieldValues(v []*string) *DescribeAppsResponseBodyResultSchema {
	s.RouteFieldValues = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchema) SetSecondRouteField(v string) *DescribeAppsResponseBodyResultSchema {
	s.SecondRouteField = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchema) SetTables(v map[string]interface{}) *DescribeAppsResponseBodyResultSchema {
	s.Tables = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchema) SetTtlField(v *DescribeAppsResponseBodyResultSchemaTtlField) *DescribeAppsResponseBodyResultSchema {
	s.TtlField = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchema) Validate() error {
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

type DescribeAppsResponseBodyResultSchemaIndexSortConfig struct {
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

func (s DescribeAppsResponseBodyResultSchemaIndexSortConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSchemaIndexSortConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSchemaIndexSortConfig) GetDirection() *string {
	return s.Direction
}

func (s *DescribeAppsResponseBodyResultSchemaIndexSortConfig) GetField() *string {
	return s.Field
}

func (s *DescribeAppsResponseBodyResultSchemaIndexSortConfig) SetDirection(v string) *DescribeAppsResponseBodyResultSchemaIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemaIndexSortConfig) SetField(v string) *DescribeAppsResponseBodyResultSchemaIndexSortConfig {
	s.Field = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemaIndexSortConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeAppsResponseBodyResultSchemaIndexes struct {
	// The attribute fields.
	FilterFields []*string `json:"filterFields,omitempty" xml:"filterFields,omitempty" type:"Repeated"`
	// The index fields.
	SearchFields map[string]interface{} `json:"searchFields,omitempty" xml:"searchFields,omitempty"`
}

func (s DescribeAppsResponseBodyResultSchemaIndexes) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSchemaIndexes) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSchemaIndexes) GetFilterFields() []*string {
	return s.FilterFields
}

func (s *DescribeAppsResponseBodyResultSchemaIndexes) GetSearchFields() map[string]interface{} {
	return s.SearchFields
}

func (s *DescribeAppsResponseBodyResultSchemaIndexes) SetFilterFields(v []*string) *DescribeAppsResponseBodyResultSchemaIndexes {
	s.FilterFields = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemaIndexes) SetSearchFields(v map[string]interface{}) *DescribeAppsResponseBodyResultSchemaIndexes {
	s.SearchFields = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemaIndexes) Validate() error {
	return dara.Validate(s)
}

type DescribeAppsResponseBodyResultSchemaTtlField struct {
	// The document clearing field.
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

func (s DescribeAppsResponseBodyResultSchemaTtlField) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSchemaTtlField) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSchemaTtlField) GetName() *string {
	return s.Name
}

func (s *DescribeAppsResponseBodyResultSchemaTtlField) GetTtl() *int64 {
	return s.Ttl
}

func (s *DescribeAppsResponseBodyResultSchemaTtlField) SetName(v string) *DescribeAppsResponseBodyResultSchemaTtlField {
	s.Name = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemaTtlField) SetTtl(v int64) *DescribeAppsResponseBodyResultSchemaTtlField {
	s.Ttl = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemaTtlField) Validate() error {
	return dara.Validate(s)
}

type DescribeAppsResponseBodyResultSchemas struct {
	// The sort configurations.
	IndexSortConfig []*DescribeAppsResponseBodyResultSchemasIndexSortConfig `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	// The index schema.
	Indexes *DescribeAppsResponseBodyResultSchemasIndexes `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
	// The name of the wide table.
	//
	// example:
	//
	// main
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The level-1 routing field.
	//
	// example:
	//
	// field1
	RouteField *string `json:"routeField,omitempty" xml:"routeField,omitempty"`
	// The hot values of the level-1 routing field.
	RouteFieldValues []*string `json:"routeFieldValues,omitempty" xml:"routeFieldValues,omitempty" type:"Repeated"`
	// The level-2 routing field. This parameter is returned if the routeFieldValues parameter is returned. By default, the wide-table primary key field is used as the level-2 routing field.
	//
	// example:
	//
	// field2
	SecondRouteField *string `json:"secondRouteField,omitempty" xml:"secondRouteField,omitempty"`
	// The table schema.
	Tables map[string]interface{} `json:"tables,omitempty" xml:"tables,omitempty"`
	// The document clearing configurations.
	TtlField *DescribeAppsResponseBodyResultSchemasTtlField `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s DescribeAppsResponseBodyResultSchemas) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSchemas) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSchemas) GetIndexSortConfig() []*DescribeAppsResponseBodyResultSchemasIndexSortConfig {
	return s.IndexSortConfig
}

func (s *DescribeAppsResponseBodyResultSchemas) GetIndexes() *DescribeAppsResponseBodyResultSchemasIndexes {
	return s.Indexes
}

func (s *DescribeAppsResponseBodyResultSchemas) GetName() *string {
	return s.Name
}

func (s *DescribeAppsResponseBodyResultSchemas) GetRouteField() *string {
	return s.RouteField
}

func (s *DescribeAppsResponseBodyResultSchemas) GetRouteFieldValues() []*string {
	return s.RouteFieldValues
}

func (s *DescribeAppsResponseBodyResultSchemas) GetSecondRouteField() *string {
	return s.SecondRouteField
}

func (s *DescribeAppsResponseBodyResultSchemas) GetTables() map[string]interface{} {
	return s.Tables
}

func (s *DescribeAppsResponseBodyResultSchemas) GetTtlField() *DescribeAppsResponseBodyResultSchemasTtlField {
	return s.TtlField
}

func (s *DescribeAppsResponseBodyResultSchemas) SetIndexSortConfig(v []*DescribeAppsResponseBodyResultSchemasIndexSortConfig) *DescribeAppsResponseBodyResultSchemas {
	s.IndexSortConfig = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemas) SetIndexes(v *DescribeAppsResponseBodyResultSchemasIndexes) *DescribeAppsResponseBodyResultSchemas {
	s.Indexes = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemas) SetName(v string) *DescribeAppsResponseBodyResultSchemas {
	s.Name = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemas) SetRouteField(v string) *DescribeAppsResponseBodyResultSchemas {
	s.RouteField = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemas) SetRouteFieldValues(v []*string) *DescribeAppsResponseBodyResultSchemas {
	s.RouteFieldValues = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemas) SetSecondRouteField(v string) *DescribeAppsResponseBodyResultSchemas {
	s.SecondRouteField = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemas) SetTables(v map[string]interface{}) *DescribeAppsResponseBodyResultSchemas {
	s.Tables = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemas) SetTtlField(v *DescribeAppsResponseBodyResultSchemasTtlField) *DescribeAppsResponseBodyResultSchemas {
	s.TtlField = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemas) Validate() error {
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

type DescribeAppsResponseBodyResultSchemasIndexSortConfig struct {
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

func (s DescribeAppsResponseBodyResultSchemasIndexSortConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSchemasIndexSortConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSchemasIndexSortConfig) GetDirection() *string {
	return s.Direction
}

func (s *DescribeAppsResponseBodyResultSchemasIndexSortConfig) GetField() *string {
	return s.Field
}

func (s *DescribeAppsResponseBodyResultSchemasIndexSortConfig) SetDirection(v string) *DescribeAppsResponseBodyResultSchemasIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemasIndexSortConfig) SetField(v string) *DescribeAppsResponseBodyResultSchemasIndexSortConfig {
	s.Field = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemasIndexSortConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeAppsResponseBodyResultSchemasIndexes struct {
	// The attribute fields.
	FilterFields []*string `json:"filterFields,omitempty" xml:"filterFields,omitempty" type:"Repeated"`
	// The index fields.
	SearchFields map[string]interface{} `json:"searchFields,omitempty" xml:"searchFields,omitempty"`
}

func (s DescribeAppsResponseBodyResultSchemasIndexes) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSchemasIndexes) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSchemasIndexes) GetFilterFields() []*string {
	return s.FilterFields
}

func (s *DescribeAppsResponseBodyResultSchemasIndexes) GetSearchFields() map[string]interface{} {
	return s.SearchFields
}

func (s *DescribeAppsResponseBodyResultSchemasIndexes) SetFilterFields(v []*string) *DescribeAppsResponseBodyResultSchemasIndexes {
	s.FilterFields = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemasIndexes) SetSearchFields(v map[string]interface{}) *DescribeAppsResponseBodyResultSchemasIndexes {
	s.SearchFields = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemasIndexes) Validate() error {
	return dara.Validate(s)
}

type DescribeAppsResponseBodyResultSchemasTtlField struct {
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

func (s DescribeAppsResponseBodyResultSchemasTtlField) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSchemasTtlField) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSchemasTtlField) GetName() *string {
	return s.Name
}

func (s *DescribeAppsResponseBodyResultSchemasTtlField) GetTtl() *int64 {
	return s.Ttl
}

func (s *DescribeAppsResponseBodyResultSchemasTtlField) SetName(v string) *DescribeAppsResponseBodyResultSchemasTtlField {
	s.Name = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemasTtlField) SetTtl(v int64) *DescribeAppsResponseBodyResultSchemasTtlField {
	s.Ttl = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemasTtlField) Validate() error {
	return dara.Validate(s)
}

type DescribeAppsResponseBodyResultSecondRanks struct {
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
	// The fine sort expression. You can define an expression that consists of fields, feature functions, and mathematical functions to implement complex sort logic.
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

func (s DescribeAppsResponseBodyResultSecondRanks) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSecondRanks) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSecondRanks) GetActive() *bool {
	return s.Active
}

func (s *DescribeAppsResponseBodyResultSecondRanks) GetDescription() *string {
	return s.Description
}

func (s *DescribeAppsResponseBodyResultSecondRanks) GetMeta() interface{} {
	return s.Meta
}

func (s *DescribeAppsResponseBodyResultSecondRanks) GetName() *string {
	return s.Name
}

func (s *DescribeAppsResponseBodyResultSecondRanks) SetActive(v bool) *DescribeAppsResponseBodyResultSecondRanks {
	s.Active = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSecondRanks) SetDescription(v string) *DescribeAppsResponseBodyResultSecondRanks {
	s.Description = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSecondRanks) SetMeta(v interface{}) *DescribeAppsResponseBodyResultSecondRanks {
	s.Meta = v
	return s
}

func (s *DescribeAppsResponseBodyResultSecondRanks) SetName(v string) *DescribeAppsResponseBodyResultSecondRanks {
	s.Name = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSecondRanks) Validate() error {
	return dara.Validate(s)
}

type DescribeAppsResponseBodyResultSummaries struct {
	// The summary configurations.
	Meta []*DescribeAppsResponseBodyResultSummariesMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The group name.
	//
	// example:
	//
	// fefault
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeAppsResponseBodyResultSummaries) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSummaries) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSummaries) GetMeta() []*DescribeAppsResponseBodyResultSummariesMeta {
	return s.Meta
}

func (s *DescribeAppsResponseBodyResultSummaries) GetName() *string {
	return s.Name
}

func (s *DescribeAppsResponseBodyResultSummaries) SetMeta(v []*DescribeAppsResponseBodyResultSummariesMeta) *DescribeAppsResponseBodyResultSummaries {
	s.Meta = v
	return s
}

func (s *DescribeAppsResponseBodyResultSummaries) SetName(v string) *DescribeAppsResponseBodyResultSummaries {
	s.Name = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSummaries) Validate() error {
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

type DescribeAppsResponseBodyResultSummariesMeta struct {
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

func (s DescribeAppsResponseBodyResultSummariesMeta) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSummariesMeta) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSummariesMeta) GetElement() *string {
	return s.Element
}

func (s *DescribeAppsResponseBodyResultSummariesMeta) GetEllipsis() *string {
	return s.Ellipsis
}

func (s *DescribeAppsResponseBodyResultSummariesMeta) GetField() *string {
	return s.Field
}

func (s *DescribeAppsResponseBodyResultSummariesMeta) GetLen() *int32 {
	return s.Len
}

func (s *DescribeAppsResponseBodyResultSummariesMeta) GetSnippet() *string {
	return s.Snippet
}

func (s *DescribeAppsResponseBodyResultSummariesMeta) SetElement(v string) *DescribeAppsResponseBodyResultSummariesMeta {
	s.Element = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSummariesMeta) SetEllipsis(v string) *DescribeAppsResponseBodyResultSummariesMeta {
	s.Ellipsis = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSummariesMeta) SetField(v string) *DescribeAppsResponseBodyResultSummariesMeta {
	s.Field = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSummariesMeta) SetLen(v int32) *DescribeAppsResponseBodyResultSummariesMeta {
	s.Len = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSummariesMeta) SetSnippet(v string) *DescribeAppsResponseBodyResultSummariesMeta {
	s.Snippet = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSummariesMeta) Validate() error {
	return dara.Validate(s)
}
