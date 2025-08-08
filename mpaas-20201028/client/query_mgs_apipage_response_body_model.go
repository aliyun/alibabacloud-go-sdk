// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMgsApipageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryMgsApipageResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryMgsApipageResponseBody
	GetResultCode() *string
	SetResultContent(v *QueryMgsApipageResponseBodyResultContent) *QueryMgsApipageResponseBody
	GetResultContent() *QueryMgsApipageResponseBodyResultContent
	SetResultMessage(v string) *QueryMgsApipageResponseBody
	GetResultMessage() *string
}

type QueryMgsApipageResponseBody struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                   `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *QueryMgsApipageResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                   `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMgsApipageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApipageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMgsApipageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMgsApipageResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryMgsApipageResponseBody) GetResultContent() *QueryMgsApipageResponseBodyResultContent {
	return s.ResultContent
}

func (s *QueryMgsApipageResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryMgsApipageResponseBody) SetRequestId(v string) *QueryMgsApipageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMgsApipageResponseBody) SetResultCode(v string) *QueryMgsApipageResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMgsApipageResponseBody) SetResultContent(v *QueryMgsApipageResponseBodyResultContent) *QueryMgsApipageResponseBody {
	s.ResultContent = v
	return s
}

func (s *QueryMgsApipageResponseBody) SetResultMessage(v string) *QueryMgsApipageResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryMgsApipageResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApipageResponseBodyResultContent struct {
	Current  *int64                                          `json:"Current,omitempty" xml:"Current,omitempty"`
	List     []*QueryMgsApipageResponseBodyResultContentList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageSize *int64                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryMgsApipageResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApipageResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *QueryMgsApipageResponseBodyResultContent) GetCurrent() *int64 {
	return s.Current
}

func (s *QueryMgsApipageResponseBodyResultContent) GetList() []*QueryMgsApipageResponseBodyResultContentList {
	return s.List
}

func (s *QueryMgsApipageResponseBodyResultContent) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryMgsApipageResponseBodyResultContent) GetTotal() *int64 {
	return s.Total
}

func (s *QueryMgsApipageResponseBodyResultContent) SetCurrent(v int64) *QueryMgsApipageResponseBodyResultContent {
	s.Current = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContent) SetList(v []*QueryMgsApipageResponseBodyResultContentList) *QueryMgsApipageResponseBodyResultContent {
	s.List = v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContent) SetPageSize(v int64) *QueryMgsApipageResponseBodyResultContent {
	s.PageSize = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContent) SetTotal(v int64) *QueryMgsApipageResponseBodyResultContent {
	s.Total = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApipageResponseBodyResultContentList struct {
	ApiInvoker         *QueryMgsApipageResponseBodyResultContentListApiInvoker         `json:"ApiInvoker,omitempty" xml:"ApiInvoker,omitempty" type:"Struct"`
	ApiName            *string                                                         `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	ApiStatus          *string                                                         `json:"ApiStatus,omitempty" xml:"ApiStatus,omitempty"`
	ApiType            *string                                                         `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	AppId              *string                                                         `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AuthRuleName       *string                                                         `json:"AuthRuleName,omitempty" xml:"AuthRuleName,omitempty"`
	CacheRule          *QueryMgsApipageResponseBodyResultContentListCacheRule          `json:"CacheRule,omitempty" xml:"CacheRule,omitempty" type:"Struct"`
	Charset            *string                                                         `json:"Charset,omitempty" xml:"Charset,omitempty"`
	CircuitBreakerRule *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule `json:"CircuitBreakerRule,omitempty" xml:"CircuitBreakerRule,omitempty" type:"Struct"`
	ContentType        *string                                                         `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	Description        *string                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate          *string                                                         `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *string                                                         `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HeaderRule         []*QueryMgsApipageResponseBodyResultContentListHeaderRule       `json:"HeaderRule,omitempty" xml:"HeaderRule,omitempty" type:"Repeated"`
	HeaderRules        []*QueryMgsApipageResponseBodyResultContentListHeaderRules      `json:"HeaderRules,omitempty" xml:"HeaderRules,omitempty" type:"Repeated"`
	Host               *string                                                         `json:"Host,omitempty" xml:"Host,omitempty"`
	Id                 *int64                                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	InterfaceType      *string                                                         `json:"InterfaceType,omitempty" xml:"InterfaceType,omitempty"`
	LimitRule          *QueryMgsApipageResponseBodyResultContentListLimitRule          `json:"LimitRule,omitempty" xml:"LimitRule,omitempty" type:"Struct"`
	Method             *string                                                         `json:"Method,omitempty" xml:"Method,omitempty"`
	MethodName         *string                                                         `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	MigrateRule        *QueryMgsApipageResponseBodyResultContentListMigrateRule        `json:"MigrateRule,omitempty" xml:"MigrateRule,omitempty" type:"Struct"`
	MockRule           *QueryMgsApipageResponseBodyResultContentListMockRule           `json:"MockRule,omitempty" xml:"MockRule,omitempty" type:"Struct"`
	NeedETag           *string                                                         `json:"NeedETag,omitempty" xml:"NeedETag,omitempty"`
	NeedEncrypt        *string                                                         `json:"NeedEncrypt,omitempty" xml:"NeedEncrypt,omitempty"`
	NeedJsonp          *string                                                         `json:"NeedJsonp,omitempty" xml:"NeedJsonp,omitempty"`
	NeedSign           *string                                                         `json:"NeedSign,omitempty" xml:"NeedSign,omitempty"`
	OperationType      *string                                                         `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	ParamGetMethod     *string                                                         `json:"ParamGetMethod,omitempty" xml:"ParamGetMethod,omitempty"`
	Path               *string                                                         `json:"Path,omitempty" xml:"Path,omitempty"`
	RequestBodyModel   *string                                                         `json:"RequestBodyModel,omitempty" xml:"RequestBodyModel,omitempty"`
	RequestParams      []*QueryMgsApipageResponseBodyResultContentListRequestParams    `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Repeated"`
	ResponseBodyModel  *string                                                         `json:"ResponseBodyModel,omitempty" xml:"ResponseBodyModel,omitempty"`
	SysId              *int64                                                          `json:"SysId,omitempty" xml:"SysId,omitempty"`
	SysName            *string                                                         `json:"SysName,omitempty" xml:"SysName,omitempty"`
	Timeout            *string                                                         `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	WorkspaceId        *string                                                         `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMgsApipageResponseBodyResultContentList) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApipageResponseBodyResultContentList) GoString() string {
	return s.String()
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetApiInvoker() *QueryMgsApipageResponseBodyResultContentListApiInvoker {
	return s.ApiInvoker
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetApiName() *string {
	return s.ApiName
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetApiStatus() *string {
	return s.ApiStatus
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetApiType() *string {
	return s.ApiType
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetAppId() *string {
	return s.AppId
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetAuthRuleName() *string {
	return s.AuthRuleName
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetCacheRule() *QueryMgsApipageResponseBodyResultContentListCacheRule {
	return s.CacheRule
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetCharset() *string {
	return s.Charset
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetCircuitBreakerRule() *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule {
	return s.CircuitBreakerRule
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetContentType() *string {
	return s.ContentType
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetDescription() *string {
	return s.Description
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetHeaderRule() []*QueryMgsApipageResponseBodyResultContentListHeaderRule {
	return s.HeaderRule
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetHeaderRules() []*QueryMgsApipageResponseBodyResultContentListHeaderRules {
	return s.HeaderRules
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetHost() *string {
	return s.Host
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetId() *int64 {
	return s.Id
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetInterfaceType() *string {
	return s.InterfaceType
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetLimitRule() *QueryMgsApipageResponseBodyResultContentListLimitRule {
	return s.LimitRule
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetMethod() *string {
	return s.Method
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetMethodName() *string {
	return s.MethodName
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetMigrateRule() *QueryMgsApipageResponseBodyResultContentListMigrateRule {
	return s.MigrateRule
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetMockRule() *QueryMgsApipageResponseBodyResultContentListMockRule {
	return s.MockRule
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetNeedETag() *string {
	return s.NeedETag
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetNeedEncrypt() *string {
	return s.NeedEncrypt
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetNeedJsonp() *string {
	return s.NeedJsonp
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetNeedSign() *string {
	return s.NeedSign
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetOperationType() *string {
	return s.OperationType
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetParamGetMethod() *string {
	return s.ParamGetMethod
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetPath() *string {
	return s.Path
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetRequestBodyModel() *string {
	return s.RequestBodyModel
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetRequestParams() []*QueryMgsApipageResponseBodyResultContentListRequestParams {
	return s.RequestParams
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetResponseBodyModel() *string {
	return s.ResponseBodyModel
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetSysId() *int64 {
	return s.SysId
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetSysName() *string {
	return s.SysName
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetTimeout() *string {
	return s.Timeout
}

func (s *QueryMgsApipageResponseBodyResultContentList) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetApiInvoker(v *QueryMgsApipageResponseBodyResultContentListApiInvoker) *QueryMgsApipageResponseBodyResultContentList {
	s.ApiInvoker = v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetApiName(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.ApiName = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetApiStatus(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.ApiStatus = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetApiType(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.ApiType = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetAppId(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.AppId = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetAuthRuleName(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.AuthRuleName = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetCacheRule(v *QueryMgsApipageResponseBodyResultContentListCacheRule) *QueryMgsApipageResponseBodyResultContentList {
	s.CacheRule = v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetCharset(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.Charset = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetCircuitBreakerRule(v *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) *QueryMgsApipageResponseBodyResultContentList {
	s.CircuitBreakerRule = v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetContentType(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.ContentType = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetDescription(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.Description = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetGmtCreate(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.GmtCreate = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetGmtModified(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.GmtModified = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetHeaderRule(v []*QueryMgsApipageResponseBodyResultContentListHeaderRule) *QueryMgsApipageResponseBodyResultContentList {
	s.HeaderRule = v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetHeaderRules(v []*QueryMgsApipageResponseBodyResultContentListHeaderRules) *QueryMgsApipageResponseBodyResultContentList {
	s.HeaderRules = v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetHost(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.Host = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetId(v int64) *QueryMgsApipageResponseBodyResultContentList {
	s.Id = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetInterfaceType(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.InterfaceType = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetLimitRule(v *QueryMgsApipageResponseBodyResultContentListLimitRule) *QueryMgsApipageResponseBodyResultContentList {
	s.LimitRule = v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetMethod(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.Method = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetMethodName(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.MethodName = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetMigrateRule(v *QueryMgsApipageResponseBodyResultContentListMigrateRule) *QueryMgsApipageResponseBodyResultContentList {
	s.MigrateRule = v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetMockRule(v *QueryMgsApipageResponseBodyResultContentListMockRule) *QueryMgsApipageResponseBodyResultContentList {
	s.MockRule = v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetNeedETag(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.NeedETag = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetNeedEncrypt(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.NeedEncrypt = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetNeedJsonp(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.NeedJsonp = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetNeedSign(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.NeedSign = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetOperationType(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.OperationType = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetParamGetMethod(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.ParamGetMethod = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetPath(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.Path = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetRequestBodyModel(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.RequestBodyModel = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetRequestParams(v []*QueryMgsApipageResponseBodyResultContentListRequestParams) *QueryMgsApipageResponseBodyResultContentList {
	s.RequestParams = v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetResponseBodyModel(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.ResponseBodyModel = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetSysId(v int64) *QueryMgsApipageResponseBodyResultContentList {
	s.SysId = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetSysName(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.SysName = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetTimeout(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.Timeout = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) SetWorkspaceId(v string) *QueryMgsApipageResponseBodyResultContentList {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentList) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApipageResponseBodyResultContentListApiInvoker struct {
	HttpInvoker *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker `json:"HttpInvoker,omitempty" xml:"HttpInvoker,omitempty" type:"Struct"`
	RpcInvoker  *string                                                            `json:"RpcInvoker,omitempty" xml:"RpcInvoker,omitempty"`
}

func (s QueryMgsApipageResponseBodyResultContentListApiInvoker) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApipageResponseBodyResultContentListApiInvoker) GoString() string {
	return s.String()
}

func (s *QueryMgsApipageResponseBodyResultContentListApiInvoker) GetHttpInvoker() *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker {
	return s.HttpInvoker
}

func (s *QueryMgsApipageResponseBodyResultContentListApiInvoker) GetRpcInvoker() *string {
	return s.RpcInvoker
}

func (s *QueryMgsApipageResponseBodyResultContentListApiInvoker) SetHttpInvoker(v *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker) *QueryMgsApipageResponseBodyResultContentListApiInvoker {
	s.HttpInvoker = v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListApiInvoker) SetRpcInvoker(v string) *QueryMgsApipageResponseBodyResultContentListApiInvoker {
	s.RpcInvoker = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListApiInvoker) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker struct {
	Charset     *string `json:"Charset,omitempty" xml:"Charset,omitempty"`
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	Host        *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Method      *string `json:"Method,omitempty" xml:"Method,omitempty"`
	Path        *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker) GoString() string {
	return s.String()
}

func (s *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker) GetCharset() *string {
	return s.Charset
}

func (s *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker) GetContentType() *string {
	return s.ContentType
}

func (s *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker) GetHost() *string {
	return s.Host
}

func (s *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker) GetMethod() *string {
	return s.Method
}

func (s *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker) GetPath() *string {
	return s.Path
}

func (s *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker) SetCharset(v string) *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker {
	s.Charset = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker) SetContentType(v string) *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker {
	s.ContentType = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker) SetHost(v string) *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker {
	s.Host = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker) SetMethod(v string) *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker {
	s.Method = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker) SetPath(v string) *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker {
	s.Path = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListApiInvokerHttpInvoker) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApipageResponseBodyResultContentListCacheRule struct {
	CacheKey  *string `json:"CacheKey,omitempty" xml:"CacheKey,omitempty"`
	NeedCache *bool   `json:"NeedCache,omitempty" xml:"NeedCache,omitempty"`
	Ttl       *int64  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s QueryMgsApipageResponseBodyResultContentListCacheRule) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApipageResponseBodyResultContentListCacheRule) GoString() string {
	return s.String()
}

func (s *QueryMgsApipageResponseBodyResultContentListCacheRule) GetCacheKey() *string {
	return s.CacheKey
}

func (s *QueryMgsApipageResponseBodyResultContentListCacheRule) GetNeedCache() *bool {
	return s.NeedCache
}

func (s *QueryMgsApipageResponseBodyResultContentListCacheRule) GetTtl() *int64 {
	return s.Ttl
}

func (s *QueryMgsApipageResponseBodyResultContentListCacheRule) SetCacheKey(v string) *QueryMgsApipageResponseBodyResultContentListCacheRule {
	s.CacheKey = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListCacheRule) SetNeedCache(v bool) *QueryMgsApipageResponseBodyResultContentListCacheRule {
	s.NeedCache = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListCacheRule) SetTtl(v int64) *QueryMgsApipageResponseBodyResultContentListCacheRule {
	s.Ttl = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListCacheRule) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule struct {
	AppId              *string  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DefaultResponse    *string  `json:"DefaultResponse,omitempty" xml:"DefaultResponse,omitempty"`
	ErrorThreshold     *int64   `json:"ErrorThreshold,omitempty" xml:"ErrorThreshold,omitempty"`
	Id                 *int64   `json:"Id,omitempty" xml:"Id,omitempty"`
	Model              *string  `json:"Model,omitempty" xml:"Model,omitempty"`
	OpenTimeoutSeconds *int64   `json:"OpenTimeoutSeconds,omitempty" xml:"OpenTimeoutSeconds,omitempty"`
	SlowRatioThreshold *float64 `json:"SlowRatioThreshold,omitempty" xml:"SlowRatioThreshold,omitempty"`
	SwitchStatus       *string  `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
	WindowsInSeconds   *int64   `json:"WindowsInSeconds,omitempty" xml:"WindowsInSeconds,omitempty"`
	WorkspaceId        *string  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) GoString() string {
	return s.String()
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) GetAppId() *string {
	return s.AppId
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) GetDefaultResponse() *string {
	return s.DefaultResponse
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) GetErrorThreshold() *int64 {
	return s.ErrorThreshold
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) GetId() *int64 {
	return s.Id
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) GetModel() *string {
	return s.Model
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) GetOpenTimeoutSeconds() *int64 {
	return s.OpenTimeoutSeconds
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) GetSlowRatioThreshold() *float64 {
	return s.SlowRatioThreshold
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) GetSwitchStatus() *string {
	return s.SwitchStatus
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) GetWindowsInSeconds() *int64 {
	return s.WindowsInSeconds
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) SetAppId(v string) *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule {
	s.AppId = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) SetDefaultResponse(v string) *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule {
	s.DefaultResponse = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) SetErrorThreshold(v int64) *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule {
	s.ErrorThreshold = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) SetId(v int64) *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule {
	s.Id = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) SetModel(v string) *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule {
	s.Model = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) SetOpenTimeoutSeconds(v int64) *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule {
	s.OpenTimeoutSeconds = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) SetSlowRatioThreshold(v float64) *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule {
	s.SlowRatioThreshold = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) SetSwitchStatus(v string) *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule {
	s.SwitchStatus = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) SetWindowsInSeconds(v int64) *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule {
	s.WindowsInSeconds = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) SetWorkspaceId(v string) *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListCircuitBreakerRule) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApipageResponseBodyResultContentListHeaderRule struct {
	HeaderKey *string `json:"HeaderKey,omitempty" xml:"HeaderKey,omitempty"`
	Location  *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryMgsApipageResponseBodyResultContentListHeaderRule) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApipageResponseBodyResultContentListHeaderRule) GoString() string {
	return s.String()
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRule) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRule) GetLocation() *string {
	return s.Location
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRule) GetType() *string {
	return s.Type
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRule) GetValue() *string {
	return s.Value
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRule) SetHeaderKey(v string) *QueryMgsApipageResponseBodyResultContentListHeaderRule {
	s.HeaderKey = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRule) SetLocation(v string) *QueryMgsApipageResponseBodyResultContentListHeaderRule {
	s.Location = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRule) SetType(v string) *QueryMgsApipageResponseBodyResultContentListHeaderRule {
	s.Type = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRule) SetValue(v string) *QueryMgsApipageResponseBodyResultContentListHeaderRule {
	s.Value = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRule) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApipageResponseBodyResultContentListHeaderRules struct {
	HeaderKey *string `json:"HeaderKey,omitempty" xml:"HeaderKey,omitempty"`
	Location  *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryMgsApipageResponseBodyResultContentListHeaderRules) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApipageResponseBodyResultContentListHeaderRules) GoString() string {
	return s.String()
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRules) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRules) GetLocation() *string {
	return s.Location
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRules) GetType() *string {
	return s.Type
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRules) GetValue() *string {
	return s.Value
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRules) SetHeaderKey(v string) *QueryMgsApipageResponseBodyResultContentListHeaderRules {
	s.HeaderKey = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRules) SetLocation(v string) *QueryMgsApipageResponseBodyResultContentListHeaderRules {
	s.Location = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRules) SetType(v string) *QueryMgsApipageResponseBodyResultContentListHeaderRules {
	s.Type = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRules) SetValue(v string) *QueryMgsApipageResponseBodyResultContentListHeaderRules {
	s.Value = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListHeaderRules) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApipageResponseBodyResultContentListLimitRule struct {
	DefaultResponse *string `json:"DefaultResponse,omitempty" xml:"DefaultResponse,omitempty"`
	I18nResponse    *string `json:"I18nResponse,omitempty" xml:"I18nResponse,omitempty"`
	Interval        *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Limit           *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s QueryMgsApipageResponseBodyResultContentListLimitRule) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApipageResponseBodyResultContentListLimitRule) GoString() string {
	return s.String()
}

func (s *QueryMgsApipageResponseBodyResultContentListLimitRule) GetDefaultResponse() *string {
	return s.DefaultResponse
}

func (s *QueryMgsApipageResponseBodyResultContentListLimitRule) GetI18nResponse() *string {
	return s.I18nResponse
}

func (s *QueryMgsApipageResponseBodyResultContentListLimitRule) GetInterval() *int64 {
	return s.Interval
}

func (s *QueryMgsApipageResponseBodyResultContentListLimitRule) GetLimit() *int64 {
	return s.Limit
}

func (s *QueryMgsApipageResponseBodyResultContentListLimitRule) GetMode() *string {
	return s.Mode
}

func (s *QueryMgsApipageResponseBodyResultContentListLimitRule) SetDefaultResponse(v string) *QueryMgsApipageResponseBodyResultContentListLimitRule {
	s.DefaultResponse = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListLimitRule) SetI18nResponse(v string) *QueryMgsApipageResponseBodyResultContentListLimitRule {
	s.I18nResponse = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListLimitRule) SetInterval(v int64) *QueryMgsApipageResponseBodyResultContentListLimitRule {
	s.Interval = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListLimitRule) SetLimit(v int64) *QueryMgsApipageResponseBodyResultContentListLimitRule {
	s.Limit = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListLimitRule) SetMode(v string) *QueryMgsApipageResponseBodyResultContentListLimitRule {
	s.Mode = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListLimitRule) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApipageResponseBodyResultContentListMigrateRule struct {
	FlowPercent          *int64  `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
	NeedMigrate          *bool   `json:"NeedMigrate,omitempty" xml:"NeedMigrate,omitempty"`
	NeedSwitchCompletely *bool   `json:"NeedSwitchCompletely,omitempty" xml:"NeedSwitchCompletely,omitempty"`
	SysId                *int64  `json:"SysId,omitempty" xml:"SysId,omitempty"`
	SysName              *string `json:"SysName,omitempty" xml:"SysName,omitempty"`
	UpstreamType         *string `json:"UpstreamType,omitempty" xml:"UpstreamType,omitempty"`
}

func (s QueryMgsApipageResponseBodyResultContentListMigrateRule) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApipageResponseBodyResultContentListMigrateRule) GoString() string {
	return s.String()
}

func (s *QueryMgsApipageResponseBodyResultContentListMigrateRule) GetFlowPercent() *int64 {
	return s.FlowPercent
}

func (s *QueryMgsApipageResponseBodyResultContentListMigrateRule) GetNeedMigrate() *bool {
	return s.NeedMigrate
}

func (s *QueryMgsApipageResponseBodyResultContentListMigrateRule) GetNeedSwitchCompletely() *bool {
	return s.NeedSwitchCompletely
}

func (s *QueryMgsApipageResponseBodyResultContentListMigrateRule) GetSysId() *int64 {
	return s.SysId
}

func (s *QueryMgsApipageResponseBodyResultContentListMigrateRule) GetSysName() *string {
	return s.SysName
}

func (s *QueryMgsApipageResponseBodyResultContentListMigrateRule) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *QueryMgsApipageResponseBodyResultContentListMigrateRule) SetFlowPercent(v int64) *QueryMgsApipageResponseBodyResultContentListMigrateRule {
	s.FlowPercent = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListMigrateRule) SetNeedMigrate(v bool) *QueryMgsApipageResponseBodyResultContentListMigrateRule {
	s.NeedMigrate = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListMigrateRule) SetNeedSwitchCompletely(v bool) *QueryMgsApipageResponseBodyResultContentListMigrateRule {
	s.NeedSwitchCompletely = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListMigrateRule) SetSysId(v int64) *QueryMgsApipageResponseBodyResultContentListMigrateRule {
	s.SysId = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListMigrateRule) SetSysName(v string) *QueryMgsApipageResponseBodyResultContentListMigrateRule {
	s.SysName = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListMigrateRule) SetUpstreamType(v string) *QueryMgsApipageResponseBodyResultContentListMigrateRule {
	s.UpstreamType = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListMigrateRule) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApipageResponseBodyResultContentListMockRule struct {
	MockData   *string `json:"MockData,omitempty" xml:"MockData,omitempty"`
	NeedMock   *bool   `json:"NeedMock,omitempty" xml:"NeedMock,omitempty"`
	Percentage *int64  `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryMgsApipageResponseBodyResultContentListMockRule) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApipageResponseBodyResultContentListMockRule) GoString() string {
	return s.String()
}

func (s *QueryMgsApipageResponseBodyResultContentListMockRule) GetMockData() *string {
	return s.MockData
}

func (s *QueryMgsApipageResponseBodyResultContentListMockRule) GetNeedMock() *bool {
	return s.NeedMock
}

func (s *QueryMgsApipageResponseBodyResultContentListMockRule) GetPercentage() *int64 {
	return s.Percentage
}

func (s *QueryMgsApipageResponseBodyResultContentListMockRule) GetType() *string {
	return s.Type
}

func (s *QueryMgsApipageResponseBodyResultContentListMockRule) SetMockData(v string) *QueryMgsApipageResponseBodyResultContentListMockRule {
	s.MockData = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListMockRule) SetNeedMock(v bool) *QueryMgsApipageResponseBodyResultContentListMockRule {
	s.NeedMock = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListMockRule) SetPercentage(v int64) *QueryMgsApipageResponseBodyResultContentListMockRule {
	s.Percentage = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListMockRule) SetType(v string) *QueryMgsApipageResponseBodyResultContentListMockRule {
	s.Type = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListMockRule) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApipageResponseBodyResultContentListRequestParams struct {
	ApiId        *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RefType      *string `json:"RefType,omitempty" xml:"RefType,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WorkspaceId  *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMgsApipageResponseBodyResultContentListRequestParams) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApipageResponseBodyResultContentListRequestParams) GoString() string {
	return s.String()
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) GetApiId() *string {
	return s.ApiId
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) GetAppId() *string {
	return s.AppId
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) GetDescription() *string {
	return s.Description
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) GetId() *int64 {
	return s.Id
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) GetLocation() *string {
	return s.Location
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) GetName() *string {
	return s.Name
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) GetRefType() *string {
	return s.RefType
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) GetType() *string {
	return s.Type
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) SetApiId(v string) *QueryMgsApipageResponseBodyResultContentListRequestParams {
	s.ApiId = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) SetAppId(v string) *QueryMgsApipageResponseBodyResultContentListRequestParams {
	s.AppId = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) SetDefaultValue(v string) *QueryMgsApipageResponseBodyResultContentListRequestParams {
	s.DefaultValue = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) SetDescription(v string) *QueryMgsApipageResponseBodyResultContentListRequestParams {
	s.Description = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) SetId(v int64) *QueryMgsApipageResponseBodyResultContentListRequestParams {
	s.Id = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) SetLocation(v string) *QueryMgsApipageResponseBodyResultContentListRequestParams {
	s.Location = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) SetName(v string) *QueryMgsApipageResponseBodyResultContentListRequestParams {
	s.Name = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) SetRefType(v string) *QueryMgsApipageResponseBodyResultContentListRequestParams {
	s.RefType = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) SetType(v string) *QueryMgsApipageResponseBodyResultContentListRequestParams {
	s.Type = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) SetWorkspaceId(v string) *QueryMgsApipageResponseBodyResultContentListRequestParams {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMgsApipageResponseBodyResultContentListRequestParams) Validate() error {
	return dara.Validate(s)
}
