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
	// This parameter is required.
	CreateCommand *CreateDataServiceApiRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type CreateDataServiceApiRequestCreateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 101
	ApiGroupId *int64 `json:"ApiGroupId,omitempty" xml:"ApiGroupId,omitempty"`
	// This parameter is required.
	ApiGroupName *string `json:"ApiGroupName,omitempty" xml:"ApiGroupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// API_01
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	ApiType *int32 `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// This parameter is required.
	BizProtocol []*int32 `json:"BizProtocol,omitempty" xml:"BizProtocol,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 600
	CacheTimeout *int32 `json:"CacheTimeout,omitempty" xml:"CacheTimeout,omitempty"`
	// example:
	//
	// 1
	CallMode         *int32  `json:"CallMode,omitempty" xml:"CallMode,omitempty"`
	CustomUpdateRate *string `json:"CustomUpdateRate,omitempty" xml:"CustomUpdateRate,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 30
	ExecutionTimeout *int32 `json:"ExecutionTimeout,omitempty" xml:"ExecutionTimeout,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RequestType *int32 `json:"RequestType,omitempty" xml:"RequestType,omitempty"`
	// This parameter is required.
	ScriptDetails *CreateDataServiceApiRequestCreateCommandScriptDetails `json:"ScriptDetails,omitempty" xml:"ScriptDetails,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// example:
	//
	// 1
	UpdateRate *int32 `json:"UpdateRate,omitempty" xml:"UpdateRate,omitempty"`
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
	return dara.Validate(s)
}

type CreateDataServiceApiRequestCreateCommandScriptDetails struct {
	// example:
	//
	// 6668888888888812345L
	DatasourceID *int64 `json:"DatasourceID,omitempty" xml:"DatasourceID,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DatasourceType *int32 `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	// example:
	//
	// false
	IsPaginated *bool `json:"IsPaginated,omitempty" xml:"IsPaginated,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// select a,b,c from table1 where d = ${d}
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// This parameter is required.
	ScriptRequestParameters []*CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters `json:"ScriptRequestParameters,omitempty" xml:"ScriptRequestParameters,omitempty" type:"Repeated"`
	// This parameter is required.
	ScriptResponseParameters []*CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters `json:"ScriptResponseParameters,omitempty" xml:"ScriptResponseParameters,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	SortPriority *int32 `json:"SortPriority,omitempty" xml:"SortPriority,omitempty"`
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
	return dara.Validate(s)
}

type CreateDataServiceApiRequestCreateCommandScriptDetailsScriptRequestParameters struct {
	// example:
	//
	// test
	ExampleValue *string `json:"ExampleValue,omitempty" xml:"ExampleValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	IsRequiredParameter *bool `json:"IsRequiredParameter,omitempty" xml:"IsRequiredParameter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// STRING
	ParameterDataType    *string `json:"ParameterDataType,omitempty" xml:"ParameterDataType,omitempty"`
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
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
	// example:
	//
	// amazing
	ExampleValue *string `json:"ExampleValue,omitempty" xml:"ExampleValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// STRING
	ParameterDataType    *string `json:"ParameterDataType,omitempty" xml:"ParameterDataType,omitempty"`
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
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

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters) SetParameterName(v string) *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters {
	s.ParameterName = &v
	return s
}

func (s *CreateDataServiceApiRequestCreateCommandScriptDetailsScriptResponseParameters) Validate() error {
	return dara.Validate(s)
}
