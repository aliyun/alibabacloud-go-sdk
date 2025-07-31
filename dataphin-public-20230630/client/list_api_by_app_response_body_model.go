// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiByAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListApiByAppResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListApiByAppResponseBody
	GetHttpStatusCode() *int32
	SetListResult(v *ListApiByAppResponseBodyListResult) *ListApiByAppResponseBody
	GetListResult() *ListApiByAppResponseBodyListResult
	SetMessage(v string) *ListApiByAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListApiByAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListApiByAppResponseBody
	GetSuccess() *bool
}

type ListApiByAppResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ListResult     *ListApiByAppResponseBodyListResult `json:"ListResult,omitempty" xml:"ListResult,omitempty" type:"Struct"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListApiByAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApiByAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiByAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListApiByAppResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListApiByAppResponseBody) GetListResult() *ListApiByAppResponseBodyListResult {
	return s.ListResult
}

func (s *ListApiByAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApiByAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApiByAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListApiByAppResponseBody) SetCode(v string) *ListApiByAppResponseBody {
	s.Code = &v
	return s
}

func (s *ListApiByAppResponseBody) SetHttpStatusCode(v int32) *ListApiByAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListApiByAppResponseBody) SetListResult(v *ListApiByAppResponseBodyListResult) *ListApiByAppResponseBody {
	s.ListResult = v
	return s
}

func (s *ListApiByAppResponseBody) SetMessage(v string) *ListApiByAppResponseBody {
	s.Message = &v
	return s
}

func (s *ListApiByAppResponseBody) SetRequestId(v string) *ListApiByAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApiByAppResponseBody) SetSuccess(v bool) *ListApiByAppResponseBody {
	s.Success = &v
	return s
}

func (s *ListApiByAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApiByAppResponseBodyListResult struct {
	Data []*ListApiByAppResponseBodyListResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApiByAppResponseBodyListResult) String() string {
	return dara.Prettify(s)
}

func (s ListApiByAppResponseBodyListResult) GoString() string {
	return s.String()
}

func (s *ListApiByAppResponseBodyListResult) GetData() []*ListApiByAppResponseBodyListResultData {
	return s.Data
}

func (s *ListApiByAppResponseBodyListResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApiByAppResponseBodyListResult) SetData(v []*ListApiByAppResponseBodyListResultData) *ListApiByAppResponseBodyListResult {
	s.Data = v
	return s
}

func (s *ListApiByAppResponseBodyListResult) SetTotalCount(v int32) *ListApiByAppResponseBodyListResult {
	s.TotalCount = &v
	return s
}

func (s *ListApiByAppResponseBodyListResult) Validate() error {
	return dara.Validate(s)
}

type ListApiByAppResponseBodyListResultData struct {
	// example:
	//
	// 987654321
	ApiNo *int64 `json:"ApiNo,omitempty" xml:"ApiNo,omitempty"`
	// example:
	//
	// 30
	ApiTimeout *int64 `json:"ApiTimeout,omitempty" xml:"ApiTimeout,omitempty"`
	// example:
	//
	// exampleApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// exampleModule
	BizModuleEnName *string `json:"BizModuleEnName,omitempty" xml:"BizModuleEnName,omitempty"`
	// example:
	//
	// 1
	CacheSwitch *string `json:"CacheSwitch,omitempty" xml:"CacheSwitch,omitempty"`
	// example:
	//
	// 60
	CacheTime *string `json:"CacheTime,omitempty" xml:"CacheTime,omitempty"`
	// example:
	//
	// 1
	CreateType *int64 `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	// example:
	//
	// 0
	DbEnv *int64 `json:"DbEnv,omitempty" xml:"DbEnv,omitempty"`
	// example:
	//
	// 这是一个示例API
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 54321
	DirectDatasourceId *int64 `json:"DirectDatasourceId,omitempty" xml:"DirectDatasourceId,omitempty"`
	// example:
	//
	// directDatasource
	DirectDatasourceName *string `json:"DirectDatasourceName,omitempty" xml:"DirectDatasourceName,omitempty"`
	// example:
	//
	// 67890
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// exampleGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 12345
	Id             *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	IsLogicalTable *bool  `json:"IsLogicalTable,omitempty" xml:"IsLogicalTable,omitempty"`
	// example:
	//
	// 0
	IsPagedQuery *int64 `json:"IsPagedQuery,omitempty" xml:"IsPagedQuery,omitempty"`
	// example:
	//
	// 100
	MaxReturnNum *int64 `json:"MaxReturnNum,omitempty" xml:"MaxReturnNum,omitempty"`
	// example:
	//
	// 0
	ModelType *int64 `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// example:
	//
	// exampleApi
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 112233
	ProjId *int64 `json:"ProjId,omitempty" xml:"ProjId,omitempty"`
	// example:
	//
	// exampleProject
	ProjName *string `json:"ProjName,omitempty" xml:"ProjName,omitempty"`
	// example:
	//
	// 3
	Protocol *int64 `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// HTTP和HTTPS
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	// -
	PublicParamList []*ListApiByAppResponseBodyListResultDataPublicParamList `json:"PublicParamList,omitempty" xml:"PublicParamList,omitempty" type:"Repeated"`
	RegisterApi     *ListApiByAppResponseBodyListResultDataRegisterApi       `json:"RegisterApi,omitempty" xml:"RegisterApi,omitempty" type:"Struct"`
	// example:
	//
	// 1
	RequestMethod *int64 `json:"RequestMethod,omitempty" xml:"RequestMethod,omitempty"`
	// example:
	//
	// GET
	RequestMethodName *string `json:"RequestMethodName,omitempty" xml:"RequestMethodName,omitempty"`
	// -
	RequestParamList []*ListApiByAppResponseBodyListResultDataRequestParamList `json:"RequestParamList,omitempty" xml:"RequestParamList,omitempty" type:"Repeated"`
	// example:
	//
	// exampleResourceGroup
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// -
	ResponseParamList []*ListApiByAppResponseBodyListResultDataResponseParamList `json:"ResponseParamList,omitempty" xml:"ResponseParamList,omitempty" type:"Repeated"`
	// example:
	//
	// {"status":"success","data":[]}
	ResultSample *string `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	// example:
	//
	// 3
	ReturnType *int64 `json:"ReturnType,omitempty" xml:"ReturnType,omitempty"`
	// example:
	//
	// JSON
	ReturnTypeName *string `json:"ReturnTypeName,omitempty" xml:"ReturnTypeName,omitempty"`
	// example:
	//
	// grp12345
	RsGrpId *string `json:"RsGrpId,omitempty" xml:"RsGrpId,omitempty"`
	// example:
	//
	// SQL
	ScriptType *string `json:"ScriptType,omitempty" xml:"ScriptType,omitempty"`
	// example:
	//
	// 0
	SpecialSql *int64 `json:"SpecialSql,omitempty" xml:"SpecialSql,omitempty"`
	// example:
	//
	// SELECT 	- FROM example_table
	SqlStatement *string `json:"SqlStatement,omitempty" xml:"SqlStatement,omitempty"`
	// example:
	//
	// exampleTable
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 10
	Timeout *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// example:
	//
	// 2
	UpdateRate *int64 `json:"UpdateRate,omitempty" xml:"UpdateRate,omitempty"`
	// example:
	//
	// 每日更新
	UpdateRateName *string `json:"UpdateRateName,omitempty" xml:"UpdateRateName,omitempty"`
	// example:
	//
	// v1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListApiByAppResponseBodyListResultData) String() string {
	return dara.Prettify(s)
}

func (s ListApiByAppResponseBodyListResultData) GoString() string {
	return s.String()
}

func (s *ListApiByAppResponseBodyListResultData) GetApiNo() *int64 {
	return s.ApiNo
}

func (s *ListApiByAppResponseBodyListResultData) GetApiTimeout() *int64 {
	return s.ApiTimeout
}

func (s *ListApiByAppResponseBodyListResultData) GetAppName() *string {
	return s.AppName
}

func (s *ListApiByAppResponseBodyListResultData) GetBizModuleEnName() *string {
	return s.BizModuleEnName
}

func (s *ListApiByAppResponseBodyListResultData) GetCacheSwitch() *string {
	return s.CacheSwitch
}

func (s *ListApiByAppResponseBodyListResultData) GetCacheTime() *string {
	return s.CacheTime
}

func (s *ListApiByAppResponseBodyListResultData) GetCreateType() *int64 {
	return s.CreateType
}

func (s *ListApiByAppResponseBodyListResultData) GetDbEnv() *int64 {
	return s.DbEnv
}

func (s *ListApiByAppResponseBodyListResultData) GetDescription() *string {
	return s.Description
}

func (s *ListApiByAppResponseBodyListResultData) GetDirectDatasourceId() *int64 {
	return s.DirectDatasourceId
}

func (s *ListApiByAppResponseBodyListResultData) GetDirectDatasourceName() *string {
	return s.DirectDatasourceName
}

func (s *ListApiByAppResponseBodyListResultData) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListApiByAppResponseBodyListResultData) GetGroupName() *string {
	return s.GroupName
}

func (s *ListApiByAppResponseBodyListResultData) GetId() *int64 {
	return s.Id
}

func (s *ListApiByAppResponseBodyListResultData) GetIsLogicalTable() *bool {
	return s.IsLogicalTable
}

func (s *ListApiByAppResponseBodyListResultData) GetIsPagedQuery() *int64 {
	return s.IsPagedQuery
}

func (s *ListApiByAppResponseBodyListResultData) GetMaxReturnNum() *int64 {
	return s.MaxReturnNum
}

func (s *ListApiByAppResponseBodyListResultData) GetModelType() *int64 {
	return s.ModelType
}

func (s *ListApiByAppResponseBodyListResultData) GetName() *string {
	return s.Name
}

func (s *ListApiByAppResponseBodyListResultData) GetProjId() *int64 {
	return s.ProjId
}

func (s *ListApiByAppResponseBodyListResultData) GetProjName() *string {
	return s.ProjName
}

func (s *ListApiByAppResponseBodyListResultData) GetProtocol() *int64 {
	return s.Protocol
}

func (s *ListApiByAppResponseBodyListResultData) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *ListApiByAppResponseBodyListResultData) GetPublicParamList() []*ListApiByAppResponseBodyListResultDataPublicParamList {
	return s.PublicParamList
}

func (s *ListApiByAppResponseBodyListResultData) GetRegisterApi() *ListApiByAppResponseBodyListResultDataRegisterApi {
	return s.RegisterApi
}

func (s *ListApiByAppResponseBodyListResultData) GetRequestMethod() *int64 {
	return s.RequestMethod
}

func (s *ListApiByAppResponseBodyListResultData) GetRequestMethodName() *string {
	return s.RequestMethodName
}

func (s *ListApiByAppResponseBodyListResultData) GetRequestParamList() []*ListApiByAppResponseBodyListResultDataRequestParamList {
	return s.RequestParamList
}

func (s *ListApiByAppResponseBodyListResultData) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ListApiByAppResponseBodyListResultData) GetResponseParamList() []*ListApiByAppResponseBodyListResultDataResponseParamList {
	return s.ResponseParamList
}

func (s *ListApiByAppResponseBodyListResultData) GetResultSample() *string {
	return s.ResultSample
}

func (s *ListApiByAppResponseBodyListResultData) GetReturnType() *int64 {
	return s.ReturnType
}

func (s *ListApiByAppResponseBodyListResultData) GetReturnTypeName() *string {
	return s.ReturnTypeName
}

func (s *ListApiByAppResponseBodyListResultData) GetRsGrpId() *string {
	return s.RsGrpId
}

func (s *ListApiByAppResponseBodyListResultData) GetScriptType() *string {
	return s.ScriptType
}

func (s *ListApiByAppResponseBodyListResultData) GetSpecialSql() *int64 {
	return s.SpecialSql
}

func (s *ListApiByAppResponseBodyListResultData) GetSqlStatement() *string {
	return s.SqlStatement
}

func (s *ListApiByAppResponseBodyListResultData) GetTableName() *string {
	return s.TableName
}

func (s *ListApiByAppResponseBodyListResultData) GetTimeout() *string {
	return s.Timeout
}

func (s *ListApiByAppResponseBodyListResultData) GetUpdateRate() *int64 {
	return s.UpdateRate
}

func (s *ListApiByAppResponseBodyListResultData) GetUpdateRateName() *string {
	return s.UpdateRateName
}

func (s *ListApiByAppResponseBodyListResultData) GetVersion() *string {
	return s.Version
}

func (s *ListApiByAppResponseBodyListResultData) SetApiNo(v int64) *ListApiByAppResponseBodyListResultData {
	s.ApiNo = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetApiTimeout(v int64) *ListApiByAppResponseBodyListResultData {
	s.ApiTimeout = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetAppName(v string) *ListApiByAppResponseBodyListResultData {
	s.AppName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetBizModuleEnName(v string) *ListApiByAppResponseBodyListResultData {
	s.BizModuleEnName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetCacheSwitch(v string) *ListApiByAppResponseBodyListResultData {
	s.CacheSwitch = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetCacheTime(v string) *ListApiByAppResponseBodyListResultData {
	s.CacheTime = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetCreateType(v int64) *ListApiByAppResponseBodyListResultData {
	s.CreateType = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetDbEnv(v int64) *ListApiByAppResponseBodyListResultData {
	s.DbEnv = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetDescription(v string) *ListApiByAppResponseBodyListResultData {
	s.Description = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetDirectDatasourceId(v int64) *ListApiByAppResponseBodyListResultData {
	s.DirectDatasourceId = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetDirectDatasourceName(v string) *ListApiByAppResponseBodyListResultData {
	s.DirectDatasourceName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetGroupId(v int64) *ListApiByAppResponseBodyListResultData {
	s.GroupId = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetGroupName(v string) *ListApiByAppResponseBodyListResultData {
	s.GroupName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetId(v int64) *ListApiByAppResponseBodyListResultData {
	s.Id = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetIsLogicalTable(v bool) *ListApiByAppResponseBodyListResultData {
	s.IsLogicalTable = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetIsPagedQuery(v int64) *ListApiByAppResponseBodyListResultData {
	s.IsPagedQuery = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetMaxReturnNum(v int64) *ListApiByAppResponseBodyListResultData {
	s.MaxReturnNum = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetModelType(v int64) *ListApiByAppResponseBodyListResultData {
	s.ModelType = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetName(v string) *ListApiByAppResponseBodyListResultData {
	s.Name = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetProjId(v int64) *ListApiByAppResponseBodyListResultData {
	s.ProjId = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetProjName(v string) *ListApiByAppResponseBodyListResultData {
	s.ProjName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetProtocol(v int64) *ListApiByAppResponseBodyListResultData {
	s.Protocol = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetProtocolName(v string) *ListApiByAppResponseBodyListResultData {
	s.ProtocolName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetPublicParamList(v []*ListApiByAppResponseBodyListResultDataPublicParamList) *ListApiByAppResponseBodyListResultData {
	s.PublicParamList = v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetRegisterApi(v *ListApiByAppResponseBodyListResultDataRegisterApi) *ListApiByAppResponseBodyListResultData {
	s.RegisterApi = v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetRequestMethod(v int64) *ListApiByAppResponseBodyListResultData {
	s.RequestMethod = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetRequestMethodName(v string) *ListApiByAppResponseBodyListResultData {
	s.RequestMethodName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetRequestParamList(v []*ListApiByAppResponseBodyListResultDataRequestParamList) *ListApiByAppResponseBodyListResultData {
	s.RequestParamList = v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetResourceGroupName(v string) *ListApiByAppResponseBodyListResultData {
	s.ResourceGroupName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetResponseParamList(v []*ListApiByAppResponseBodyListResultDataResponseParamList) *ListApiByAppResponseBodyListResultData {
	s.ResponseParamList = v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetResultSample(v string) *ListApiByAppResponseBodyListResultData {
	s.ResultSample = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetReturnType(v int64) *ListApiByAppResponseBodyListResultData {
	s.ReturnType = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetReturnTypeName(v string) *ListApiByAppResponseBodyListResultData {
	s.ReturnTypeName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetRsGrpId(v string) *ListApiByAppResponseBodyListResultData {
	s.RsGrpId = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetScriptType(v string) *ListApiByAppResponseBodyListResultData {
	s.ScriptType = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetSpecialSql(v int64) *ListApiByAppResponseBodyListResultData {
	s.SpecialSql = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetSqlStatement(v string) *ListApiByAppResponseBodyListResultData {
	s.SqlStatement = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetTableName(v string) *ListApiByAppResponseBodyListResultData {
	s.TableName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetTimeout(v string) *ListApiByAppResponseBodyListResultData {
	s.Timeout = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetUpdateRate(v int64) *ListApiByAppResponseBodyListResultData {
	s.UpdateRate = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetUpdateRateName(v string) *ListApiByAppResponseBodyListResultData {
	s.UpdateRateName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) SetVersion(v string) *ListApiByAppResponseBodyListResultData {
	s.Version = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultData) Validate() error {
	return dara.Validate(s)
}

type ListApiByAppResponseBodyListResultDataPublicParamList struct {
	// example:
	//
	// yyyy-MM-dd
	DateFormat *string `json:"DateFormat,omitempty" xml:"DateFormat,omitempty"`
	// example:
	//
	// default_public_value
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// 这是一个示例公共参数
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// CODE_003
	DescriptionCode *string `json:"DescriptionCode,omitempty" xml:"DescriptionCode,omitempty"`
	// example:
	//
	// 3001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// app_key_value
	InitialValue *string `json:"InitialValue,omitempty" xml:"InitialValue,omitempty"`
	// example:
	//
	// publicColumn1
	MappingColumn *string `json:"MappingColumn,omitempty" xml:"MappingColumn,omitempty"`
	// example:
	//
	// 1
	Must *int64 `json:"Must,omitempty" xml:"Must,omitempty"`
	// example:
	//
	// =
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 1
	Optional *int64 `json:"Optional,omitempty" xml:"Optional,omitempty"`
	// example:
	//
	// original_public_column
	OriginalColumn *string `json:"OriginalColumn,omitempty" xml:"OriginalColumn,omitempty"`
	// example:
	//
	// publicParam1
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// example:
	//
	// String
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// example:
	//
	// /path/to/public_param
	ParameterLocation *string `json:"ParameterLocation,omitempty" xml:"ParameterLocation,omitempty"`
	// example:
	//
	// publicValue1
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// -
	TableAndDsList []*ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList `json:"TableAndDsList,omitempty" xml:"TableAndDsList,omitempty" type:"Repeated"`
}

func (s ListApiByAppResponseBodyListResultDataPublicParamList) String() string {
	return dara.Prettify(s)
}

func (s ListApiByAppResponseBodyListResultDataPublicParamList) GoString() string {
	return s.String()
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) GetDateFormat() *string {
	return s.DateFormat
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) GetDescription() *string {
	return s.Description
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) GetDescriptionCode() *string {
	return s.DescriptionCode
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) GetId() *int64 {
	return s.Id
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) GetInitialValue() *string {
	return s.InitialValue
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) GetMappingColumn() *string {
	return s.MappingColumn
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) GetMust() *int64 {
	return s.Must
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) GetOperator() *string {
	return s.Operator
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) GetOptional() *int64 {
	return s.Optional
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) GetOriginalColumn() *string {
	return s.OriginalColumn
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) GetParamName() *string {
	return s.ParamName
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) GetParamType() *string {
	return s.ParamType
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) GetParameterLocation() *string {
	return s.ParameterLocation
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) GetSample() *string {
	return s.Sample
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) GetTableAndDsList() []*ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList {
	return s.TableAndDsList
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) SetDateFormat(v string) *ListApiByAppResponseBodyListResultDataPublicParamList {
	s.DateFormat = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) SetDefaultValue(v string) *ListApiByAppResponseBodyListResultDataPublicParamList {
	s.DefaultValue = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) SetDescription(v string) *ListApiByAppResponseBodyListResultDataPublicParamList {
	s.Description = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) SetDescriptionCode(v string) *ListApiByAppResponseBodyListResultDataPublicParamList {
	s.DescriptionCode = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) SetId(v int64) *ListApiByAppResponseBodyListResultDataPublicParamList {
	s.Id = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) SetInitialValue(v string) *ListApiByAppResponseBodyListResultDataPublicParamList {
	s.InitialValue = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) SetMappingColumn(v string) *ListApiByAppResponseBodyListResultDataPublicParamList {
	s.MappingColumn = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) SetMust(v int64) *ListApiByAppResponseBodyListResultDataPublicParamList {
	s.Must = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) SetOperator(v string) *ListApiByAppResponseBodyListResultDataPublicParamList {
	s.Operator = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) SetOptional(v int64) *ListApiByAppResponseBodyListResultDataPublicParamList {
	s.Optional = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) SetOriginalColumn(v string) *ListApiByAppResponseBodyListResultDataPublicParamList {
	s.OriginalColumn = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) SetParamName(v string) *ListApiByAppResponseBodyListResultDataPublicParamList {
	s.ParamName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) SetParamType(v string) *ListApiByAppResponseBodyListResultDataPublicParamList {
	s.ParamType = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) SetParameterLocation(v string) *ListApiByAppResponseBodyListResultDataPublicParamList {
	s.ParameterLocation = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) SetSample(v string) *ListApiByAppResponseBodyListResultDataPublicParamList {
	s.Sample = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) SetTableAndDsList(v []*ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList) *ListApiByAppResponseBodyListResultDataPublicParamList {
	s.TableAndDsList = v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamList) Validate() error {
	return dara.Validate(s)
}

type ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList struct {
	// example:
	//
	// ds54321
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// publicDatasource
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// example:
	//
	// 3
	DatasourceType *int64 `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	// example:
	//
	// https://example.com/public_datasource
	DatasourceUrl *string `json:"DatasourceUrl,omitempty" xml:"DatasourceUrl,omitempty"`
	// example:
	//
	// public_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList) String() string {
	return dara.Prettify(s)
}

func (s ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList) GoString() string {
	return s.String()
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList) GetDatasourceId() *string {
	return s.DatasourceId
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList) GetDatasourceType() *int64 {
	return s.DatasourceType
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList) GetDatasourceUrl() *string {
	return s.DatasourceUrl
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList) GetTableName() *string {
	return s.TableName
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList) SetDatasourceId(v string) *ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList {
	s.DatasourceId = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList) SetDatasourceName(v string) *ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList {
	s.DatasourceName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList) SetDatasourceType(v int64) *ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList {
	s.DatasourceType = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList) SetDatasourceUrl(v string) *ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList {
	s.DatasourceUrl = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList) SetTableName(v string) *ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList {
	s.TableName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataPublicParamListTableAndDsList) Validate() error {
	return dara.Validate(s)
}

type ListApiByAppResponseBodyListResultDataRegisterApi struct {
	// example:
	//
	// 67890
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// example:
	//
	// OAuth2
	AuthenticationMode *string `json:"AuthenticationMode,omitempty" xml:"AuthenticationMode,omitempty"`
	// example:
	//
	// ds67890
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// registerDatasource
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// example:
	//
	// {"status":"error","message":"Invalid request"}
	FailExample *string `json:"FailExample,omitempty" xml:"FailExample,omitempty"`
	// example:
	//
	// 1
	HttpMethod *int64 `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// example:
	//
	// 0
	ModelType *int64 `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// example:
	//
	// /api/v1/register
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// HTTPS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// {"status":"success","data":[]}
	SuccessExample *string `json:"SuccessExample,omitempty" xml:"SuccessExample,omitempty"`
	// example:
	//
	// 30
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// example:
	//
	// https://example.com/register_api
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListApiByAppResponseBodyListResultDataRegisterApi) String() string {
	return dara.Prettify(s)
}

func (s ListApiByAppResponseBodyListResultDataRegisterApi) GoString() string {
	return s.String()
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) GetApiId() *int64 {
	return s.ApiId
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) GetAuthenticationMode() *string {
	return s.AuthenticationMode
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) GetDatasourceId() *string {
	return s.DatasourceId
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) GetFailExample() *string {
	return s.FailExample
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) GetHttpMethod() *int64 {
	return s.HttpMethod
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) GetModelType() *int64 {
	return s.ModelType
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) GetPath() *string {
	return s.Path
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) GetProtocol() *string {
	return s.Protocol
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) GetSuccessExample() *string {
	return s.SuccessExample
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) GetTimeout() *int64 {
	return s.Timeout
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) GetUrl() *string {
	return s.Url
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) SetApiId(v int64) *ListApiByAppResponseBodyListResultDataRegisterApi {
	s.ApiId = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) SetAuthenticationMode(v string) *ListApiByAppResponseBodyListResultDataRegisterApi {
	s.AuthenticationMode = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) SetDatasourceId(v string) *ListApiByAppResponseBodyListResultDataRegisterApi {
	s.DatasourceId = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) SetDatasourceName(v string) *ListApiByAppResponseBodyListResultDataRegisterApi {
	s.DatasourceName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) SetFailExample(v string) *ListApiByAppResponseBodyListResultDataRegisterApi {
	s.FailExample = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) SetHttpMethod(v int64) *ListApiByAppResponseBodyListResultDataRegisterApi {
	s.HttpMethod = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) SetModelType(v int64) *ListApiByAppResponseBodyListResultDataRegisterApi {
	s.ModelType = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) SetPath(v string) *ListApiByAppResponseBodyListResultDataRegisterApi {
	s.Path = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) SetProtocol(v string) *ListApiByAppResponseBodyListResultDataRegisterApi {
	s.Protocol = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) SetSuccessExample(v string) *ListApiByAppResponseBodyListResultDataRegisterApi {
	s.SuccessExample = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) SetTimeout(v int64) *ListApiByAppResponseBodyListResultDataRegisterApi {
	s.Timeout = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) SetUrl(v string) *ListApiByAppResponseBodyListResultDataRegisterApi {
	s.Url = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRegisterApi) Validate() error {
	return dara.Validate(s)
}

type ListApiByAppResponseBodyListResultDataRequestParamList struct {
	// example:
	//
	// yyyy-MM-dd
	DateFormat *string `json:"DateFormat,omitempty" xml:"DateFormat,omitempty"`
	// example:
	//
	// default_value
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// 这是一个示例参数
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// CODE_001
	DescriptionCode *string `json:"DescriptionCode,omitempty" xml:"DescriptionCode,omitempty"`
	// example:
	//
	// 1001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// app_key_value
	InitialValue *string `json:"InitialValue,omitempty" xml:"InitialValue,omitempty"`
	// example:
	//
	// column1
	MappingColumn *string `json:"MappingColumn,omitempty" xml:"MappingColumn,omitempty"`
	// example:
	//
	// 1
	Must *int64 `json:"Must,omitempty" xml:"Must,omitempty"`
	// example:
	//
	// =
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 1
	Optional *int64 `json:"Optional,omitempty" xml:"Optional,omitempty"`
	// example:
	//
	// original_column
	OriginalColumn *string `json:"OriginalColumn,omitempty" xml:"OriginalColumn,omitempty"`
	// example:
	//
	// param1
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// example:
	//
	// String
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// example:
	//
	// /path/to/param
	ParameterLocation *string `json:"ParameterLocation,omitempty" xml:"ParameterLocation,omitempty"`
	// example:
	//
	// value1
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// -
	TableAndDsList []*ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList `json:"TableAndDsList,omitempty" xml:"TableAndDsList,omitempty" type:"Repeated"`
}

func (s ListApiByAppResponseBodyListResultDataRequestParamList) String() string {
	return dara.Prettify(s)
}

func (s ListApiByAppResponseBodyListResultDataRequestParamList) GoString() string {
	return s.String()
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) GetDateFormat() *string {
	return s.DateFormat
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) GetDescription() *string {
	return s.Description
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) GetDescriptionCode() *string {
	return s.DescriptionCode
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) GetId() *int64 {
	return s.Id
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) GetInitialValue() *string {
	return s.InitialValue
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) GetMappingColumn() *string {
	return s.MappingColumn
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) GetMust() *int64 {
	return s.Must
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) GetOperator() *string {
	return s.Operator
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) GetOptional() *int64 {
	return s.Optional
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) GetOriginalColumn() *string {
	return s.OriginalColumn
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) GetParamName() *string {
	return s.ParamName
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) GetParamType() *string {
	return s.ParamType
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) GetParameterLocation() *string {
	return s.ParameterLocation
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) GetSample() *string {
	return s.Sample
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) GetTableAndDsList() []*ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList {
	return s.TableAndDsList
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) SetDateFormat(v string) *ListApiByAppResponseBodyListResultDataRequestParamList {
	s.DateFormat = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) SetDefaultValue(v string) *ListApiByAppResponseBodyListResultDataRequestParamList {
	s.DefaultValue = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) SetDescription(v string) *ListApiByAppResponseBodyListResultDataRequestParamList {
	s.Description = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) SetDescriptionCode(v string) *ListApiByAppResponseBodyListResultDataRequestParamList {
	s.DescriptionCode = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) SetId(v int64) *ListApiByAppResponseBodyListResultDataRequestParamList {
	s.Id = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) SetInitialValue(v string) *ListApiByAppResponseBodyListResultDataRequestParamList {
	s.InitialValue = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) SetMappingColumn(v string) *ListApiByAppResponseBodyListResultDataRequestParamList {
	s.MappingColumn = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) SetMust(v int64) *ListApiByAppResponseBodyListResultDataRequestParamList {
	s.Must = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) SetOperator(v string) *ListApiByAppResponseBodyListResultDataRequestParamList {
	s.Operator = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) SetOptional(v int64) *ListApiByAppResponseBodyListResultDataRequestParamList {
	s.Optional = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) SetOriginalColumn(v string) *ListApiByAppResponseBodyListResultDataRequestParamList {
	s.OriginalColumn = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) SetParamName(v string) *ListApiByAppResponseBodyListResultDataRequestParamList {
	s.ParamName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) SetParamType(v string) *ListApiByAppResponseBodyListResultDataRequestParamList {
	s.ParamType = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) SetParameterLocation(v string) *ListApiByAppResponseBodyListResultDataRequestParamList {
	s.ParameterLocation = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) SetSample(v string) *ListApiByAppResponseBodyListResultDataRequestParamList {
	s.Sample = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) SetTableAndDsList(v []*ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList) *ListApiByAppResponseBodyListResultDataRequestParamList {
	s.TableAndDsList = v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamList) Validate() error {
	return dara.Validate(s)
}

type ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList struct {
	// example:
	//
	// ds12345
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// exampleDatasource
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// example:
	//
	// 1
	DatasourceType *int64 `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	// example:
	//
	// https://example.com/datasource
	DatasourceUrl *string `json:"DatasourceUrl,omitempty" xml:"DatasourceUrl,omitempty"`
	// example:
	//
	// example_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList) String() string {
	return dara.Prettify(s)
}

func (s ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList) GoString() string {
	return s.String()
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList) GetDatasourceId() *string {
	return s.DatasourceId
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList) GetDatasourceType() *int64 {
	return s.DatasourceType
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList) GetDatasourceUrl() *string {
	return s.DatasourceUrl
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList) GetTableName() *string {
	return s.TableName
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList) SetDatasourceId(v string) *ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList {
	s.DatasourceId = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList) SetDatasourceName(v string) *ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList {
	s.DatasourceName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList) SetDatasourceType(v int64) *ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList {
	s.DatasourceType = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList) SetDatasourceUrl(v string) *ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList {
	s.DatasourceUrl = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList) SetTableName(v string) *ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList {
	s.TableName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataRequestParamListTableAndDsList) Validate() error {
	return dara.Validate(s)
}

type ListApiByAppResponseBodyListResultDataResponseParamList struct {
	// example:
	//
	// yyyy-MM-dd
	DateFormat *string `json:"DateFormat,omitempty" xml:"DateFormat,omitempty"`
	// example:
	//
	// default_response_value
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// 这是一个示例响应参数
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// CODE_002
	DescriptionCode *string `json:"DescriptionCode,omitempty" xml:"DescriptionCode,omitempty"`
	// example:
	//
	// 2001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// app_key_value
	InitialValue *string `json:"InitialValue,omitempty" xml:"InitialValue,omitempty"`
	// example:
	//
	// responseColumn1
	MappingColumn *string `json:"MappingColumn,omitempty" xml:"MappingColumn,omitempty"`
	// example:
	//
	// 1
	Must *int64 `json:"Must,omitempty" xml:"Must,omitempty"`
	// example:
	//
	// =
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 1
	Optional *int64 `json:"Optional,omitempty" xml:"Optional,omitempty"`
	// example:
	//
	// original_response_column
	OriginalColumn *string `json:"OriginalColumn,omitempty" xml:"OriginalColumn,omitempty"`
	// example:
	//
	// responseParam1
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// example:
	//
	// String
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// example:
	//
	// /path/to/response_param
	ParameterLocation *string `json:"ParameterLocation,omitempty" xml:"ParameterLocation,omitempty"`
	// example:
	//
	// responseValue1
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// -
	TableAndDsList []*ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList `json:"TableAndDsList,omitempty" xml:"TableAndDsList,omitempty" type:"Repeated"`
}

func (s ListApiByAppResponseBodyListResultDataResponseParamList) String() string {
	return dara.Prettify(s)
}

func (s ListApiByAppResponseBodyListResultDataResponseParamList) GoString() string {
	return s.String()
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) GetDateFormat() *string {
	return s.DateFormat
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) GetDescription() *string {
	return s.Description
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) GetDescriptionCode() *string {
	return s.DescriptionCode
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) GetId() *int64 {
	return s.Id
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) GetInitialValue() *string {
	return s.InitialValue
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) GetMappingColumn() *string {
	return s.MappingColumn
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) GetMust() *int64 {
	return s.Must
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) GetOperator() *string {
	return s.Operator
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) GetOptional() *int64 {
	return s.Optional
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) GetOriginalColumn() *string {
	return s.OriginalColumn
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) GetParamName() *string {
	return s.ParamName
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) GetParamType() *string {
	return s.ParamType
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) GetParameterLocation() *string {
	return s.ParameterLocation
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) GetSample() *string {
	return s.Sample
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) GetTableAndDsList() []*ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList {
	return s.TableAndDsList
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) SetDateFormat(v string) *ListApiByAppResponseBodyListResultDataResponseParamList {
	s.DateFormat = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) SetDefaultValue(v string) *ListApiByAppResponseBodyListResultDataResponseParamList {
	s.DefaultValue = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) SetDescription(v string) *ListApiByAppResponseBodyListResultDataResponseParamList {
	s.Description = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) SetDescriptionCode(v string) *ListApiByAppResponseBodyListResultDataResponseParamList {
	s.DescriptionCode = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) SetId(v int64) *ListApiByAppResponseBodyListResultDataResponseParamList {
	s.Id = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) SetInitialValue(v string) *ListApiByAppResponseBodyListResultDataResponseParamList {
	s.InitialValue = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) SetMappingColumn(v string) *ListApiByAppResponseBodyListResultDataResponseParamList {
	s.MappingColumn = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) SetMust(v int64) *ListApiByAppResponseBodyListResultDataResponseParamList {
	s.Must = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) SetOperator(v string) *ListApiByAppResponseBodyListResultDataResponseParamList {
	s.Operator = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) SetOptional(v int64) *ListApiByAppResponseBodyListResultDataResponseParamList {
	s.Optional = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) SetOriginalColumn(v string) *ListApiByAppResponseBodyListResultDataResponseParamList {
	s.OriginalColumn = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) SetParamName(v string) *ListApiByAppResponseBodyListResultDataResponseParamList {
	s.ParamName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) SetParamType(v string) *ListApiByAppResponseBodyListResultDataResponseParamList {
	s.ParamType = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) SetParameterLocation(v string) *ListApiByAppResponseBodyListResultDataResponseParamList {
	s.ParameterLocation = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) SetSample(v string) *ListApiByAppResponseBodyListResultDataResponseParamList {
	s.Sample = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) SetTableAndDsList(v []*ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList) *ListApiByAppResponseBodyListResultDataResponseParamList {
	s.TableAndDsList = v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamList) Validate() error {
	return dara.Validate(s)
}

type ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList struct {
	// example:
	//
	// ds67890
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// responseDatasource
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// example:
	//
	// 2
	DatasourceType *int64 `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	// example:
	//
	// https://example.com/response_datasource
	DatasourceUrl *string `json:"DatasourceUrl,omitempty" xml:"DatasourceUrl,omitempty"`
	// example:
	//
	// response_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList) String() string {
	return dara.Prettify(s)
}

func (s ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList) GoString() string {
	return s.String()
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList) GetDatasourceId() *string {
	return s.DatasourceId
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList) GetDatasourceType() *int64 {
	return s.DatasourceType
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList) GetDatasourceUrl() *string {
	return s.DatasourceUrl
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList) GetTableName() *string {
	return s.TableName
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList) SetDatasourceId(v string) *ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList {
	s.DatasourceId = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList) SetDatasourceName(v string) *ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList {
	s.DatasourceName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList) SetDatasourceType(v int64) *ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList {
	s.DatasourceType = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList) SetDatasourceUrl(v string) *ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList {
	s.DatasourceUrl = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList) SetTableName(v string) *ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList {
	s.TableName = &v
	return s
}

func (s *ListApiByAppResponseBodyListResultDataResponseParamListTableAndDsList) Validate() error {
	return dara.Validate(s)
}
