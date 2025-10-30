// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataServiceApiDocumentResponseBody
	GetCode() *string
	SetData(v *GetDataServiceApiDocumentResponseBodyData) *GetDataServiceApiDocumentResponseBody
	GetData() *GetDataServiceApiDocumentResponseBodyData
	SetHttpStatusCode(v int32) *GetDataServiceApiDocumentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataServiceApiDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDataServiceApiDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataServiceApiDocumentResponseBody
	GetSuccess() *bool
}

type GetDataServiceApiDocumentResponseBody struct {
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDataServiceApiDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataServiceApiDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiDocumentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataServiceApiDocumentResponseBody) GetData() *GetDataServiceApiDocumentResponseBodyData {
	return s.Data
}

func (s *GetDataServiceApiDocumentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataServiceApiDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataServiceApiDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceApiDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataServiceApiDocumentResponseBody) SetCode(v string) *GetDataServiceApiDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBody) SetData(v *GetDataServiceApiDocumentResponseBodyData) *GetDataServiceApiDocumentResponseBody {
	s.Data = v
	return s
}

func (s *GetDataServiceApiDocumentResponseBody) SetHttpStatusCode(v int32) *GetDataServiceApiDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBody) SetMessage(v string) *GetDataServiceApiDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBody) SetRequestId(v string) *GetDataServiceApiDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBody) SetSuccess(v bool) *GetDataServiceApiDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataServiceApiDocumentResponseBodyData struct {
	// example:
	//
	// 102101
	ApiId           *int64                                                    `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiRegisterInfo *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo `json:"ApiRegisterInfo,omitempty" xml:"ApiRegisterInfo,omitempty" type:"Struct"`
	// example:
	//
	// 60
	ApiTimeout *int32 `json:"ApiTimeout,omitempty" xml:"ApiTimeout,omitempty"`
	// example:
	//
	// bizUnit_test
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// example:
	//
	// 600
	CacheTime *string `json:"CacheTime,omitempty" xml:"CacheTime,omitempty"`
	// example:
	//
	// 1
	CreateType *int32 `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	DirectDatasourceId *int64 `json:"DirectDatasourceId,omitempty" xml:"DirectDatasourceId,omitempty"`
	// example:
	//
	// test
	DirectDatasourceName *string `json:"DirectDatasourceName,omitempty" xml:"DirectDatasourceName,omitempty"`
	// example:
	//
	// 1
	Env *int32 `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 1011
	GroupId *int32 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 1011
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IsLogicalTable *bool   `json:"IsLogicalTable,omitempty" xml:"IsLogicalTable,omitempty"`
	IsPagedQuery   *bool   `json:"IsPagedQuery,omitempty" xml:"IsPagedQuery,omitempty"`
	IsSpecialSql   *bool   `json:"IsSpecialSql,omitempty" xml:"IsSpecialSql,omitempty"`
	// example:
	//
	// 1
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// 测试
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OpenCache *bool   `json:"OpenCache,omitempty" xml:"OpenCache,omitempty"`
	// example:
	//
	// 10201
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 1
	Protocol        *int32                                                      `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	PublicParamList []*GetDataServiceApiDocumentResponseBodyDataPublicParamList `json:"PublicParamList,omitempty" xml:"PublicParamList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	RequestMethod    *int32                                                       `json:"RequestMethod,omitempty" xml:"RequestMethod,omitempty"`
	RequestParamList []*GetDataServiceApiDocumentResponseBodyDataRequestParamList `json:"RequestParamList,omitempty" xml:"RequestParamList,omitempty" type:"Repeated"`
	// example:
	//
	// 10021
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// test
	ResourceGroupName *string                                                       `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	ResponseParamList []*GetDataServiceApiDocumentResponseBodyDataResponseParamList `json:"ResponseParamList,omitempty" xml:"ResponseParamList,omitempty" type:"Repeated"`
	// example:
	//
	// {"count": 88}
	ResultSample *string `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	// example:
	//
	// 100
	ReturnLimit *int32 `json:"ReturnLimit,omitempty" xml:"ReturnLimit,omitempty"`
	// example:
	//
	// 1
	ReturnType *int32 `json:"ReturnType,omitempty" xml:"ReturnType,omitempty"`
	// example:
	//
	// NORMAL_SQL
	ScriptType *string `json:"ScriptType,omitempty" xml:"ScriptType,omitempty"`
	// example:
	//
	// select col1 from table1;
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// example:
	//
	// t_logical_test1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 60
	Timeout *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// example:
	//
	// 1
	UpdateRate *int32 `json:"UpdateRate,omitempty" xml:"UpdateRate,omitempty"`
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetDataServiceApiDocumentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetApiId() *int64 {
	return s.ApiId
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetApiRegisterInfo() *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo {
	return s.ApiRegisterInfo
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetApiTimeout() *int32 {
	return s.ApiTimeout
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetCacheTime() *string {
	return s.CacheTime
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetCreateType() *int32 {
	return s.CreateType
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetDirectDatasourceId() *int64 {
	return s.DirectDatasourceId
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetDirectDatasourceName() *string {
	return s.DirectDatasourceName
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetEnv() *int32 {
	return s.Env
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetGroupId() *int32 {
	return s.GroupId
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetIsLogicalTable() *bool {
	return s.IsLogicalTable
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetIsPagedQuery() *bool {
	return s.IsPagedQuery
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetIsSpecialSql() *bool {
	return s.IsSpecialSql
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetMode() *int32 {
	return s.Mode
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetOpenCache() *bool {
	return s.OpenCache
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetProtocol() *int32 {
	return s.Protocol
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetPublicParamList() []*GetDataServiceApiDocumentResponseBodyDataPublicParamList {
	return s.PublicParamList
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetRequestMethod() *int32 {
	return s.RequestMethod
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetRequestParamList() []*GetDataServiceApiDocumentResponseBodyDataRequestParamList {
	return s.RequestParamList
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetResponseParamList() []*GetDataServiceApiDocumentResponseBodyDataResponseParamList {
	return s.ResponseParamList
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetResultSample() *string {
	return s.ResultSample
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetReturnLimit() *int32 {
	return s.ReturnLimit
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetReturnType() *int32 {
	return s.ReturnType
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetScriptType() *string {
	return s.ScriptType
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetSql() *string {
	return s.Sql
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetTableName() *string {
	return s.TableName
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetTimeout() *string {
	return s.Timeout
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetUpdateRate() *int32 {
	return s.UpdateRate
}

func (s *GetDataServiceApiDocumentResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetApiId(v int64) *GetDataServiceApiDocumentResponseBodyData {
	s.ApiId = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetApiRegisterInfo(v *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) *GetDataServiceApiDocumentResponseBodyData {
	s.ApiRegisterInfo = v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetApiTimeout(v int32) *GetDataServiceApiDocumentResponseBodyData {
	s.ApiTimeout = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetBizUnitName(v string) *GetDataServiceApiDocumentResponseBodyData {
	s.BizUnitName = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetCacheTime(v string) *GetDataServiceApiDocumentResponseBodyData {
	s.CacheTime = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetCreateType(v int32) *GetDataServiceApiDocumentResponseBodyData {
	s.CreateType = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetDescription(v string) *GetDataServiceApiDocumentResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetDirectDatasourceId(v int64) *GetDataServiceApiDocumentResponseBodyData {
	s.DirectDatasourceId = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetDirectDatasourceName(v string) *GetDataServiceApiDocumentResponseBodyData {
	s.DirectDatasourceName = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetEnv(v int32) *GetDataServiceApiDocumentResponseBodyData {
	s.Env = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetGroupId(v int32) *GetDataServiceApiDocumentResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetGroupName(v string) *GetDataServiceApiDocumentResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetIsLogicalTable(v bool) *GetDataServiceApiDocumentResponseBodyData {
	s.IsLogicalTable = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetIsPagedQuery(v bool) *GetDataServiceApiDocumentResponseBodyData {
	s.IsPagedQuery = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetIsSpecialSql(v bool) *GetDataServiceApiDocumentResponseBodyData {
	s.IsSpecialSql = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetMode(v int32) *GetDataServiceApiDocumentResponseBodyData {
	s.Mode = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetName(v string) *GetDataServiceApiDocumentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetOpenCache(v bool) *GetDataServiceApiDocumentResponseBodyData {
	s.OpenCache = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetProjectId(v int32) *GetDataServiceApiDocumentResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetProjectName(v string) *GetDataServiceApiDocumentResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetProtocol(v int32) *GetDataServiceApiDocumentResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetPublicParamList(v []*GetDataServiceApiDocumentResponseBodyDataPublicParamList) *GetDataServiceApiDocumentResponseBodyData {
	s.PublicParamList = v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetRequestMethod(v int32) *GetDataServiceApiDocumentResponseBodyData {
	s.RequestMethod = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetRequestParamList(v []*GetDataServiceApiDocumentResponseBodyDataRequestParamList) *GetDataServiceApiDocumentResponseBodyData {
	s.RequestParamList = v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetResourceGroupId(v string) *GetDataServiceApiDocumentResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetResourceGroupName(v string) *GetDataServiceApiDocumentResponseBodyData {
	s.ResourceGroupName = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetResponseParamList(v []*GetDataServiceApiDocumentResponseBodyDataResponseParamList) *GetDataServiceApiDocumentResponseBodyData {
	s.ResponseParamList = v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetResultSample(v string) *GetDataServiceApiDocumentResponseBodyData {
	s.ResultSample = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetReturnLimit(v int32) *GetDataServiceApiDocumentResponseBodyData {
	s.ReturnLimit = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetReturnType(v int32) *GetDataServiceApiDocumentResponseBodyData {
	s.ReturnType = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetScriptType(v string) *GetDataServiceApiDocumentResponseBodyData {
	s.ScriptType = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetSql(v string) *GetDataServiceApiDocumentResponseBodyData {
	s.Sql = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetTableName(v string) *GetDataServiceApiDocumentResponseBodyData {
	s.TableName = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetTimeout(v string) *GetDataServiceApiDocumentResponseBodyData {
	s.Timeout = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetUpdateRate(v int32) *GetDataServiceApiDocumentResponseBodyData {
	s.UpdateRate = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) SetVersion(v string) *GetDataServiceApiDocumentResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyData) Validate() error {
	if s.ApiRegisterInfo != nil {
		if err := s.ApiRegisterInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PublicParamList != nil {
		for _, item := range s.PublicParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RequestParamList != nil {
		for _, item := range s.RequestParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResponseParamList != nil {
		for _, item := range s.ResponseParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo struct {
	// example:
	//
	// 3
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// 102311
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// test
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// example:
	//
	// 2
	HttpMethod *int32 `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// example:
	//
	// 1
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// /api/test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// https
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// example:
	//
	// http://192.168.1.1:8080
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) GetAuthType() *string {
	return s.AuthType
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) GetDatasourceId() *string {
	return s.DatasourceId
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) GetHttpMethod() *int32 {
	return s.HttpMethod
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) GetMode() *int32 {
	return s.Mode
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) GetPath() *string {
	return s.Path
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) GetProtocol() *string {
	return s.Protocol
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) GetUrl() *string {
	return s.Url
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) SetAuthType(v string) *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo {
	s.AuthType = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) SetDatasourceId(v string) *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo {
	s.DatasourceId = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) SetDatasourceName(v string) *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo {
	s.DatasourceName = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) SetHttpMethod(v int32) *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo {
	s.HttpMethod = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) SetMode(v int32) *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo {
	s.Mode = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) SetPath(v string) *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo {
	s.Path = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) SetProtocol(v string) *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo {
	s.Protocol = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) SetTimeout(v int32) *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo {
	s.Timeout = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) SetUrl(v string) *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo {
	s.Url = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataApiRegisterInfo) Validate() error {
	return dara.Validate(s)
}

type GetDataServiceApiDocumentResponseBodyDataPublicParamList struct {
	// example:
	//
	// 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IsRequired  *bool   `json:"IsRequired,omitempty" xml:"IsRequired,omitempty"`
	// example:
	//
	// col1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// test
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataServiceApiDocumentResponseBodyDataPublicParamList) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiDocumentResponseBodyDataPublicParamList) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiDocumentResponseBodyDataPublicParamList) GetDescription() *string {
	return s.Description
}

func (s *GetDataServiceApiDocumentResponseBodyDataPublicParamList) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *GetDataServiceApiDocumentResponseBodyDataPublicParamList) GetName() *string {
	return s.Name
}

func (s *GetDataServiceApiDocumentResponseBodyDataPublicParamList) GetSample() *string {
	return s.Sample
}

func (s *GetDataServiceApiDocumentResponseBodyDataPublicParamList) GetType() *string {
	return s.Type
}

func (s *GetDataServiceApiDocumentResponseBodyDataPublicParamList) SetDescription(v string) *GetDataServiceApiDocumentResponseBodyDataPublicParamList {
	s.Description = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataPublicParamList) SetIsRequired(v bool) *GetDataServiceApiDocumentResponseBodyDataPublicParamList {
	s.IsRequired = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataPublicParamList) SetName(v string) *GetDataServiceApiDocumentResponseBodyDataPublicParamList {
	s.Name = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataPublicParamList) SetSample(v string) *GetDataServiceApiDocumentResponseBodyDataPublicParamList {
	s.Sample = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataPublicParamList) SetType(v string) *GetDataServiceApiDocumentResponseBodyDataPublicParamList {
	s.Type = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataPublicParamList) Validate() error {
	return dara.Validate(s)
}

type GetDataServiceApiDocumentResponseBodyDataRequestParamList struct {
	// example:
	//
	// 1
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IsRequired  *bool   `json:"IsRequired,omitempty" xml:"IsRequired,omitempty"`
	// example:
	//
	// col1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// test
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataServiceApiDocumentResponseBodyDataRequestParamList) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiDocumentResponseBodyDataRequestParamList) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiDocumentResponseBodyDataRequestParamList) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetDataServiceApiDocumentResponseBodyDataRequestParamList) GetDescription() *string {
	return s.Description
}

func (s *GetDataServiceApiDocumentResponseBodyDataRequestParamList) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *GetDataServiceApiDocumentResponseBodyDataRequestParamList) GetName() *string {
	return s.Name
}

func (s *GetDataServiceApiDocumentResponseBodyDataRequestParamList) GetSample() *string {
	return s.Sample
}

func (s *GetDataServiceApiDocumentResponseBodyDataRequestParamList) GetType() *string {
	return s.Type
}

func (s *GetDataServiceApiDocumentResponseBodyDataRequestParamList) SetDefaultValue(v string) *GetDataServiceApiDocumentResponseBodyDataRequestParamList {
	s.DefaultValue = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataRequestParamList) SetDescription(v string) *GetDataServiceApiDocumentResponseBodyDataRequestParamList {
	s.Description = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataRequestParamList) SetIsRequired(v bool) *GetDataServiceApiDocumentResponseBodyDataRequestParamList {
	s.IsRequired = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataRequestParamList) SetName(v string) *GetDataServiceApiDocumentResponseBodyDataRequestParamList {
	s.Name = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataRequestParamList) SetSample(v string) *GetDataServiceApiDocumentResponseBodyDataRequestParamList {
	s.Sample = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataRequestParamList) SetType(v string) *GetDataServiceApiDocumentResponseBodyDataRequestParamList {
	s.Type = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataRequestParamList) Validate() error {
	return dara.Validate(s)
}

type GetDataServiceApiDocumentResponseBodyDataResponseParamList struct {
	// example:
	//
	// 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// col1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// test
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataServiceApiDocumentResponseBodyDataResponseParamList) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiDocumentResponseBodyDataResponseParamList) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiDocumentResponseBodyDataResponseParamList) GetDescription() *string {
	return s.Description
}

func (s *GetDataServiceApiDocumentResponseBodyDataResponseParamList) GetName() *string {
	return s.Name
}

func (s *GetDataServiceApiDocumentResponseBodyDataResponseParamList) GetSample() *string {
	return s.Sample
}

func (s *GetDataServiceApiDocumentResponseBodyDataResponseParamList) GetType() *string {
	return s.Type
}

func (s *GetDataServiceApiDocumentResponseBodyDataResponseParamList) SetDescription(v string) *GetDataServiceApiDocumentResponseBodyDataResponseParamList {
	s.Description = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataResponseParamList) SetName(v string) *GetDataServiceApiDocumentResponseBodyDataResponseParamList {
	s.Name = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataResponseParamList) SetSample(v string) *GetDataServiceApiDocumentResponseBodyDataResponseParamList {
	s.Sample = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataResponseParamList) SetType(v string) *GetDataServiceApiDocumentResponseBodyDataResponseParamList {
	s.Type = &v
	return s
}

func (s *GetDataServiceApiDocumentResponseBodyDataResponseParamList) Validate() error {
	return dara.Validate(s)
}
