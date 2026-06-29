// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateDataServiceApiRequestCreateCommand) *CreateDataServiceApiRequest
	GetCreateCommand() *CreateDataServiceApiRequestCreateCommand
	SetOpTenantId(v int64) *CreateDataServiceApiRequest
	GetOpTenantId() *int64
}

type CreateDataServiceApiRequest struct {
	// The request for creating an API.
	//
	// This parameter is required.
	CreateCommand *CreateDataServiceApiRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateDataServiceApiRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceApiRequest) GoString() string {
	return s.String()
}

func (s *CreateDataServiceApiRequest) GetCreateCommand() *CreateDataServiceApiRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateDataServiceApiRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateDataServiceApiRequest) SetCreateCommand(v *CreateDataServiceApiRequestCreateCommand) *CreateDataServiceApiRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateDataServiceApiRequest) SetOpTenantId(v int64) *CreateDataServiceApiRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateDataServiceApiRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataServiceApiRequestCreateCommand struct {
	// The group ID of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 101
	ApiGroupId *int64 `json:"ApiGroupId,omitempty" xml:"ApiGroupId,omitempty"`
	// The group name of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 默认API分组
	ApiGroupName *string `json:"ApiGroupName,omitempty" xml:"ApiGroupName,omitempty"`
	// The name of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// API_01
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The type of the API. Valid values:
	//
	// - 3: datasource SQL mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ApiType *int32 `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// The protocol. Different gateway types support different protocols. For more information, see the documentation. Valid values:
	//
	// - 0: HTTP
	//
	// - 1: HTTPS.
	//
	// This parameter is required.
	BizProtocol []*int32 `json:"BizProtocol,omitempty" xml:"BizProtocol,omitempty" type:"Repeated"`
	// The cache timeout period, in seconds.
	//
	// example:
	//
	// 600
	CacheTimeout *int32 `json:"CacheTimeout,omitempty" xml:"CacheTimeout,omitempty"`
	// The call mode of the API. Default value: 1. Valid values:
	//
	// - 1: synchronous call
	//
	// - 2: asynchronous call.
	//
	// example:
	//
	// 1
	CallMode *int32 `json:"CallMode,omitempty" xml:"CallMode,omitempty"`
	// The custom update frequency. This parameter is required when the update frequency is set to custom.
	//
	// example:
	//
	// 每天8点
	CustomUpdateRate *string `json:"CustomUpdateRate,omitempty" xml:"CustomUpdateRate,omitempty"`
	// The description of the API.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The configuration of the operation-type API. This parameter is not required when creating a query-type API.
	DmlConfig *CreateDataServiceApiRequestCreateCommandDmlConfig `json:"DmlConfig,omitempty" xml:"DmlConfig,omitempty" type:"Struct"`
	// The execution timeout period for asynchronous API calls. This parameter takes effect only for asynchronous API calls and is required when the call mode is asynchronous.
	//
	// example:
	//
	// 30
	ExecutionTimeout *int32 `json:"ExecutionTimeout,omitempty" xml:"ExecutionTimeout,omitempty"`
	// The development mode of the API. Valid values:
	//
	// - 0: Basic mode
	//
	// - 1: Dev-Prod mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The ID of the data service project.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The request method of the API. Valid values:
	//
	// - 0 (GET): Returns a single record. The query result is unique.
	//
	// - 1 (LIST): Returns multiple records.
	//
	// - 2 (CREATE): Creates objects. Supports single or batch creation.
	//
	// - 3 (UPDATE): Updates objects. Supports single or batch updates.
	//
	// - 4 (DELETE): Deletes objects. Supports single or batch deletions.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RequestType *int32 `json:"RequestType,omitempty" xml:"RequestType,omitempty"`
	// Specifies whether to return the SQL in the result.
	//
	// example:
	//
	// true
	ReturnSqlSwitch *bool `json:"ReturnSqlSwitch,omitempty" xml:"ReturnSqlSwitch,omitempty"`
	// The list of row-level permission IDs.
	RowPermissionIds []*int64 `json:"RowPermissionIds,omitempty" xml:"RowPermissionIds,omitempty" type:"Repeated"`
	// The details of the script API.
	//
	// This parameter is required.
	ScriptDetails *CreateDataServiceApiRequestCreateCommandScriptDetails `json:"ScriptDetails,omitempty" xml:"ScriptDetails,omitempty" type:"Struct"`
	// The timeout period, in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The update frequency. Default value: 1. Valid values:
	//
	// - 0: custom
	//
	// - 1: day
	//
	// - 2: hour
	//
	// - 3: minute.
	//
	// example:
	//
	// 1
	UpdateRate *int32 `json:"UpdateRate,omitempty" xml:"UpdateRate,omitempty"`
	// The version of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// V1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateDataServiceApiRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceApiRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateDataServiceApiRequestCreateCommand) GetApiGroupId() *int64 {
	return s.ApiGroupId
}

func (s *CreateDataServiceApiRequestCreateCommand) GetApiGroupName() *string {
	return s.ApiGroupName
}

func (s *CreateDataServiceApiRequestCreateCommand) GetApiName() *string {
	return s.ApiName
}

func (s *CreateDataServiceApiRequestCreateCommand) GetApiType() *int32 {
	return s.ApiType
}

func (s *CreateDataServiceApiRequestCreateCommand) GetBizProtocol() []*int32 {
	return s.BizProtocol
}

func (s *CreateDataServiceApiRequestCreateCommand) GetCacheTimeout() *int32 {
	return s.CacheTimeout
}

func (s *CreateDataServiceApiRequestCreateCommand) GetCallMode() *int32 {
	return s.CallMode
}

func (s *CreateDataServiceApiRequestCreateCommand) GetCustomUpdateRate() *string {
	return s.CustomUpdateRate
}

func (s *CreateDataServiceApiRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateDataServiceApiRequestCreateCommand) GetDmlConfig() *CreateDataServiceApiRequestCreateCommandDmlConfig {
	return s.DmlConfig
}

func (s *CreateDataServiceApiRequestCreateCommand) GetExecutionTimeout() *int32 {
	return s.ExecutionTimeout
}

func (s *CreateDataServiceApiRequestCreateCommand) GetMode() *int32 {
	return s.Mode
}

func (s *CreateDataServiceApiRequestCreateCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataServiceApiRequestCreateCommand) GetRequestType() *int32 {
	return s.RequestType
}

func (s *CreateDataServiceApiRequestCreateCommand) GetReturnSqlSwitch() *bool {
	return s.ReturnSqlSwitch
}

func (s *CreateDataServiceApiRequestCreateCommand) GetRowPermissionIds() []*int64 {
	return s.RowPermissionIds
}

func (s *CreateDataServiceApiRequestCreateCommand) GetScriptDetails() *CreateDataServiceApiRequestCreateCommandScriptDetails {
	return s.ScriptDetails
}

func (s *CreateDataServiceApiRequestCreateCommand) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateDataServiceApiRequestCreateCommand) GetUpdateRate() *int32 {
	return s.UpdateRate
}

func (s *CreateDataServiceApiRequestCreateCommand) GetVersion() *string {
	return s.Version
}

func (s *CreateDataServiceApiRequestCreateCommand) SetApiGroupId(v int64) *CreateDataServiceApiRequestCreateCommand {
	s.ApiGroupId = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetApiGroupName(v string) *CreateDataServiceApiRequestCreateCommand {
	s.ApiGroupName = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetApiName(v string) *CreateDataServiceApiRequestCreateCommand {
	s.ApiName = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetApiType(v int32) *CreateDataServiceApiRequestCreateCommand {
	s.ApiType = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetBizProtocol(v []*int32) *CreateDataServiceApiRequestCreateCommand {
	s.BizProtocol = v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetCacheTimeout(v int32) *CreateDataServiceApiRequestCreateCommand {
	s.CacheTimeout = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetCallMode(v int32) *CreateDataServiceApiRequestCreateCommand {
	s.CallMode = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetCustomUpdateRate(v string) *CreateDataServiceApiRequestCreateCommand {
	s.CustomUpdateRate = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetDescription(v string) *CreateDataServiceApiRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetDmlConfig(v *CreateDataServiceApiRequestCreateCommandDmlConfig) *CreateDataServiceApiRequestCreateCommand {
	s.DmlConfig = v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetExecutionTimeout(v int32) *CreateDataServiceApiRequestCreateCommand {
	s.ExecutionTimeout = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetMode(v int32) *CreateDataServiceApiRequestCreateCommand {
	s.Mode = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetProjectId(v int64) *CreateDataServiceApiRequestCreateCommand {
	s.ProjectId = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetRequestType(v int32) *CreateDataServiceApiRequestCreateCommand {
	s.RequestType = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetReturnSqlSwitch(v bool) *CreateDataServiceApiRequestCreateCommand {
	s.ReturnSqlSwitch = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetRowPermissionIds(v []*int64) *CreateDataServiceApiRequestCreateCommand {
	s.RowPermissionIds = v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetScriptDetails(v *CreateDataServiceApiRequestCreateCommandScriptDetails) *CreateDataServiceApiRequestCreateCommand {
	s.ScriptDetails = v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetTimeout(v int32) *CreateDataServiceApiRequestCreateCommand {
	s.Timeout = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetUpdateRate(v int32) *CreateDataServiceApiRequestCreateCommand {
	s.UpdateRate = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) SetVersion(v string) *CreateDataServiceApiRequestCreateCommand {
	s.Version = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommand) Validate() error {
	if s.DmlConfig != nil {
		if err := s.DmlConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ScriptDetails != nil {
		if err := s.ScriptDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataServiceApiRequestCreateCommandDmlConfig struct {
	// The data volume per batch. Valid values:
	//
	// - When the data volume type is single record, this parameter cannot be set.
	//
	// - When the data volume type is batch:
	//
	//   - If the transaction processing mode is 1, this parameter cannot be set.
	//
	//   - If the transaction processing mode is 2, the value ranges from 1 to 1000000.
	//
	// example:
	//
	// 1000
	BatchInputDataSize *int32 `json:"BatchInputDataSize,omitempty" xml:"BatchInputDataSize,omitempty"`
	// The data volume type. Valid values:
	//
	// - 1: single record
	//
	// - 2: batch.
	//
	// example:
	//
	// 2
	DataVolumeType *int32 `json:"DataVolumeType,omitempty" xml:"DataVolumeType,omitempty"`
	// The error handling method. Valid values:
	//
	// - 1: partial success allowed
	//
	// - 2: all must succeed
	//
	// Parameter rules:
	//
	// - When the data volume type is single record, this parameter cannot be set.
	//
	// - When the data volume type is batch, the value is 1 or 2.
	//
	// example:
	//
	// 1
	ErrorHandlingType *int32 `json:"ErrorHandlingType,omitempty" xml:"ErrorHandlingType,omitempty"`
	// The maximum number of input records. Valid values:
	//
	// - When the data volume type is single record, this parameter cannot be set.
	//
	// - When the data volume type is batch, the value ranges from 1 to 1000000.
	//
	// example:
	//
	// 10000
	MaxInputDataSize *int32 `json:"MaxInputDataSize,omitempty" xml:"MaxInputDataSize,omitempty"`
	// The degree of parallelism. Valid values:
	//
	// - When the data volume type is single record, this parameter cannot be set.
	//
	// - When the data volume type is batch:
	//
	//   - If the transaction processing mode is 1, this parameter cannot be set.
	//
	//   - If the transaction processing mode is 2, the value ranges from 1 to 5.
	//
	// example:
	//
	// 1
	ParallelNum *int32 `json:"ParallelNum,omitempty" xml:"ParallelNum,omitempty"`
	// The transaction processing mode. Valid values:
	//
	// - 0: no transaction
	//
	// - 1: no batching
	//
	// - 2: batch processing
	//
	// Parameter rules:
	//
	// - When the data volume type is single record, the transaction processing mode is 0.
	//
	// - When the data volume type is batch:
	//
	//   - If the error handling method is 1, the transaction processing mode is 1 or 2.
	//
	//   - If the error handling method is 2, the transaction processing mode can only be 1.
	//
	// example:
	//
	// 2
	TransactionType *int32 `json:"TransactionType,omitempty" xml:"TransactionType,omitempty"`
}

func (s CreateDataServiceApiRequestCreateCommandDmlConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceApiRequestCreateCommandDmlConfig) GoString() string {
	return s.String()
}

func (s *CreateDataServiceApiRequestCreateCommandDmlConfig) GetBatchInputDataSize() *int32 {
	return s.BatchInputDataSize
}

func (s *CreateDataServiceApiRequestCreateCommandDmlConfig) GetDataVolumeType() *int32 {
	return s.DataVolumeType
}

func (s *CreateDataServiceApiRequestCreateCommandDmlConfig) GetErrorHandlingType() *int32 {
	return s.ErrorHandlingType
}

func (s *CreateDataServiceApiRequestCreateCommandDmlConfig) GetMaxInputDataSize() *int32 {
	return s.MaxInputDataSize
}

func (s *CreateDataServiceApiRequestCreateCommandDmlConfig) GetParallelNum() *int32 {
	return s.ParallelNum
}

func (s *CreateDataServiceApiRequestCreateCommandDmlConfig) GetTransactionType() *int32 {
	return s.TransactionType
}

func (s *CreateDataServiceApiRequestCreateCommandDmlConfig) SetBatchInputDataSize(v int32) *CreateDataServiceApiRequestCreateCommandDmlConfig {
	s.BatchInputDataSize = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandDmlConfig) SetDataVolumeType(v int32) *CreateDataServiceApiRequestCreateCommandDmlConfig {
	s.DataVolumeType = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandDmlConfig) SetErrorHandlingType(v int32) *CreateDataServiceApiRequestCreateCommandDmlConfig {
	s.ErrorHandlingType = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandDmlConfig) SetMaxInputDataSize(v int32) *CreateDataServiceApiRequestCreateCommandDmlConfig {
	s.MaxInputDataSize = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandDmlConfig) SetParallelNum(v int32) *CreateDataServiceApiRequestCreateCommandDmlConfig {
	s.ParallelNum = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandDmlConfig) SetTransactionType(v int32) *CreateDataServiceApiRequestCreateCommandDmlConfig {
	s.TransactionType = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandDmlConfig) Validate() error {
	return dara.Validate(s)
}

type CreateDataServiceApiRequestCreateCommandScriptDetails struct {
	// The ID of the datasource. This parameter is required when the API mode is direct datasource connection.
	//
	// example:
	//
	// 6668888888888812345L
	DatasourceID *int64 `json:"DatasourceID,omitempty" xml:"DatasourceID,omitempty"`
	// The data type on which the API is based. Valid values:
	//
	// - 1: datasource.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DatasourceType *int32 `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	// Specifies whether to paginate the results. This parameter is required only when RequestType is set to List. Default value: false. Pagination is not supported in asynchronous call mode.
	//
	// example:
	//
	// false
	IsPaginated *bool `json:"IsPaginated,omitempty" xml:"IsPaginated,omitempty"`
	// The SQL script.
	//
	// This parameter is required.
	//
	// example:
	//
	// select a,b,c from table1 where d = ${d}
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// The list of request parameters for the script API.
	ScriptRequestParameters []*CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters `json:"ScriptRequestParameters,omitempty" xml:"ScriptRequestParameters,omitempty" type:"Repeated"`
	// The list of response parameters for the script API.
	ScriptResponseParameters []*CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters `json:"ScriptResponseParameters,omitempty" xml:"ScriptResponseParameters,omitempty" type:"Repeated"`
	// The sorting priority. This parameter takes effect only when the SQL mode is basic mode. Default value: 2. Valid values:
	//
	// - 1: SQL script
	//
	// - 2: OrderByList request parameter.
	//
	// example:
	//
	// 2
	SortPriority *int32 `json:"SortPriority,omitempty" xml:"SortPriority,omitempty"`
	// The SQL mode. Valid values:
	//
	// - 1: basic mode
	//
	// - 2: advanced mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SqlMode *int32 `json:"SqlMode,omitempty" xml:"SqlMode,omitempty"`
}

func (s CreateDataServiceApiRequestCreateCommandScriptDetails) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceApiRequestCreateCommandScriptDetails) GoString() string {
	return s.String()
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetails) GetDatasourceID() *int64 {
	return s.DatasourceID
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetails) GetDatasourceType() *int32 {
	return s.DatasourceType
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetails) GetIsPaginated() *bool {
	return s.IsPaginated
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetails) GetScript() *string {
	return s.Script
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetails) GetScriptRequestParameters() []*CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters {
	return s.ScriptRequestParameters
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetails) GetScriptResponseParameters() []*CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters {
	return s.ScriptResponseParameters
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetails) GetSortPriority() *int32 {
	return s.SortPriority
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetails) GetSqlMode() *int32 {
	return s.SqlMode
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetails) SetDatasourceID(v int64) *CreateDataServiceApiRequestCreateCommandScriptDetails {
	s.DatasourceID = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetails) SetDatasourceType(v int32) *CreateDataServiceApiRequestCreateCommandScriptDetails {
	s.DatasourceType = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetails) SetIsPaginated(v bool) *CreateDataServiceApiRequestCreateCommandScriptDetails {
	s.IsPaginated = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetails) SetScript(v string) *CreateDataServiceApiRequestCreateCommandScriptDetails {
	s.Script = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetails) SetScriptRequestParameters(v []*CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) *CreateDataServiceApiRequestCreateCommandScriptDetails {
	s.ScriptRequestParameters = v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetails) SetScriptResponseParameters(v []*CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters) *CreateDataServiceApiRequestCreateCommandScriptDetails {
	s.ScriptResponseParameters = v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetails) SetSortPriority(v int32) *CreateDataServiceApiRequestCreateCommandScriptDetails {
	s.SortPriority = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetails) SetSqlMode(v int32) *CreateDataServiceApiRequestCreateCommandScriptDetails {
	s.SqlMode = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetails) Validate() error {
	if s.ScriptRequestParameters != nil {
		for _, item := range s.ScriptRequestParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScriptResponseParameters != nil {
		for _, item := range s.ScriptResponseParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters struct {
	// The default value of the input parameter for operation-type APIs. This parameter takes effect when the parameter is not required. If not specified, the value is null.
	//
	// example:
	//
	// test
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The example value.
	//
	// example:
	//
	// test
	ExampleValue *string `json:"ExampleValue,omitempty" xml:"ExampleValue,omitempty"`
	// Specifies whether the parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	IsRequiredParameter *bool `json:"IsRequiredParameter,omitempty" xml:"IsRequiredParameter,omitempty"`
	// The data type. Valid values:
	//
	// - "STRING"
	//
	// - "DOUBLE"
	//
	// - "INT"
	//
	// - "DATE"
	//
	// - "LONG"
	//
	// - "FLOAT"
	//
	// - "BOOLEAN"
	//
	// - "SHORT"
	//
	// - "BYTE"
	//
	// - "BIGDECIMAL"
	//
	// - "BINARY"
	//
	// - "ARRAY"
	//
	// - "Array(int)"
	//
	// - "Array(string)".
	//
	// This parameter is required.
	//
	// example:
	//
	// STRING
	ParameterDataType *string `json:"ParameterDataType,omitempty" xml:"ParameterDataType,omitempty"`
	// The parameter description.
	//
	// example:
	//
	// 字段d
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The parameter name.
	//
	// This parameter is required.
	//
	// example:
	//
	// d
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The value type of the parameter. Valid values:
	//
	// - 1 (single value): A fixed value used for operators such as =, >=, <=, >, <, !=, and between.
	//
	// - 2 (multiple values): The input parameter contains multiple values separated by commas (,). Used for In and Not In operators.
	//
	// This parameter is required.
	//
	// example:
	//
	// =
	ParameterValueType *string `json:"ParameterValueType,omitempty" xml:"ParameterValueType,omitempty"`
}

func (s CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) GetExampleValue() *string {
	return s.ExampleValue
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) GetIsRequiredParameter() *bool {
	return s.IsRequiredParameter
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) GetParameterDataType() *string {
	return s.ParameterDataType
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) GetParameterValueType() *string {
	return s.ParameterValueType
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) SetDefaultValue(v string) *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters {
	s.DefaultValue = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) SetExampleValue(v string) *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters {
	s.ExampleValue = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) SetIsRequiredParameter(v bool) *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters {
	s.IsRequiredParameter = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) SetParameterDataType(v string) *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters {
	s.ParameterDataType = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) SetParameterDescription(v string) *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters {
	s.ParameterDescription = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) SetParameterName(v string) *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters {
	s.ParameterName = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) SetParameterValueType(v string) *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters {
	s.ParameterValueType = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters) Validate() error {
	return dara.Validate(s)
}

type CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters struct {
	// The example value.
	//
	// example:
	//
	// amazing
	ExampleValue *string `json:"ExampleValue,omitempty" xml:"ExampleValue,omitempty"`
	// The data type. Valid values:
	//
	// - "STRING"
	//
	// - "DOUBLE"
	//
	// - "INT"
	//
	// - "DATE"
	//
	// - "LONG"
	//
	// - "FLOAT"
	//
	// - "BOOLEAN"
	//
	// - "SHORT"
	//
	// - "BYTE"
	//
	// - "BIGDECIMAL"
	//
	// - "BINARY"
	//
	// - "ARRAY"
	//
	// - "Array(int)"
	//
	// - "Array(string)".
	//
	// This parameter is required.
	//
	// example:
	//
	// STRING
	ParameterDataType *string `json:"ParameterDataType,omitempty" xml:"ParameterDataType,omitempty"`
	// The parameter description.
	//
	// example:
	//
	// 字段a
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The location of the response parameter for operation-type APIs. This parameter must be set when the API is an operation-type API with batch data volume. Valid values:
	//
	// - success: the response data of a successful operation
	//
	// - failed: the response data of a failed operation.
	//
	// example:
	//
	// success
	ParameterLocation *string `json:"ParameterLocation,omitempty" xml:"ParameterLocation,omitempty"`
	// The parameter name.
	//
	// This parameter is required.
	//
	// example:
	//
	// a
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
}

func (s CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters) GoString() string {
	return s.String()
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters) GetExampleValue() *string {
	return s.ExampleValue
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters) GetParameterDataType() *string {
	return s.ParameterDataType
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters) GetParameterLocation() *string {
	return s.ParameterLocation
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters) SetExampleValue(v string) *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters {
	s.ExampleValue = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters) SetParameterDataType(v string) *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters {
	s.ParameterDataType = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters) SetParameterDescription(v string) *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters {
	s.ParameterDescription = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters) SetParameterLocation(v string) *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters {
	s.ParameterLocation = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters) SetParameterName(v string) *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters {
	s.ParameterName = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters) Validate() error {
	return dara.Validate(s)
}
