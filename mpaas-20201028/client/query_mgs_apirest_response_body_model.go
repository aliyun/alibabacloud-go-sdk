// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMgsApirestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryMgsApirestResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryMgsApirestResponseBody
	GetResultCode() *string
	SetResultContent(v *QueryMgsApirestResponseBodyResultContent) *QueryMgsApirestResponseBody
	GetResultContent() *QueryMgsApirestResponseBodyResultContent
	SetResultMessage(v string) *QueryMgsApirestResponseBody
	GetResultMessage() *string
}

type QueryMgsApirestResponseBody struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                   `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *QueryMgsApirestResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                   `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMgsApirestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApirestResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMgsApirestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMgsApirestResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryMgsApirestResponseBody) GetResultContent() *QueryMgsApirestResponseBodyResultContent {
	return s.ResultContent
}

func (s *QueryMgsApirestResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryMgsApirestResponseBody) SetRequestId(v string) *QueryMgsApirestResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMgsApirestResponseBody) SetResultCode(v string) *QueryMgsApirestResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMgsApirestResponseBody) SetResultContent(v *QueryMgsApirestResponseBodyResultContent) *QueryMgsApirestResponseBody {
	s.ResultContent = v
	return s
}

func (s *QueryMgsApirestResponseBody) SetResultMessage(v string) *QueryMgsApirestResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryMgsApirestResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApirestResponseBodyResultContent struct {
	ErrorMessage *string                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	Value        *QueryMgsApirestResponseBodyResultContentValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s QueryMgsApirestResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApirestResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *QueryMgsApirestResponseBodyResultContent) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryMgsApirestResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMgsApirestResponseBodyResultContent) GetValue() *QueryMgsApirestResponseBodyResultContentValue {
	return s.Value
}

func (s *QueryMgsApirestResponseBodyResultContent) SetErrorMessage(v string) *QueryMgsApirestResponseBodyResultContent {
	s.ErrorMessage = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContent) SetSuccess(v bool) *QueryMgsApirestResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContent) SetValue(v *QueryMgsApirestResponseBodyResultContentValue) *QueryMgsApirestResponseBodyResultContent {
	s.Value = v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApirestResponseBodyResultContentValue struct {
	ApiInvoker         *QueryMgsApirestResponseBodyResultContentValueApiInvoker         `json:"ApiInvoker,omitempty" xml:"ApiInvoker,omitempty" type:"Struct"`
	ApiName            *string                                                          `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	ApiStatus          *string                                                          `json:"ApiStatus,omitempty" xml:"ApiStatus,omitempty"`
	ApiType            *string                                                          `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	AppId              *string                                                          `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AuthRuleName       *string                                                          `json:"AuthRuleName,omitempty" xml:"AuthRuleName,omitempty"`
	CacheRule          *QueryMgsApirestResponseBodyResultContentValueCacheRule          `json:"CacheRule,omitempty" xml:"CacheRule,omitempty" type:"Struct"`
	Charset            *string                                                          `json:"Charset,omitempty" xml:"Charset,omitempty"`
	CircuitBreakerRule *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule `json:"CircuitBreakerRule,omitempty" xml:"CircuitBreakerRule,omitempty" type:"Struct"`
	ContentType        *string                                                          `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	DefaultLimitRule   *QueryMgsApirestResponseBodyResultContentValueDefaultLimitRule   `json:"DefaultLimitRule,omitempty" xml:"DefaultLimitRule,omitempty" type:"Struct"`
	Description        *string                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate          *string                                                          `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *string                                                          `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HeaderRule         []*QueryMgsApirestResponseBodyResultContentValueHeaderRule       `json:"HeaderRule,omitempty" xml:"HeaderRule,omitempty" type:"Repeated"`
	HeaderRules        []*QueryMgsApirestResponseBodyResultContentValueHeaderRules      `json:"HeaderRules,omitempty" xml:"HeaderRules,omitempty" type:"Repeated"`
	Host               *string                                                          `json:"Host,omitempty" xml:"Host,omitempty"`
	Id                 *int64                                                           `json:"Id,omitempty" xml:"Id,omitempty"`
	InterfaceType      *string                                                          `json:"InterfaceType,omitempty" xml:"InterfaceType,omitempty"`
	LimitRule          *QueryMgsApirestResponseBodyResultContentValueLimitRule          `json:"LimitRule,omitempty" xml:"LimitRule,omitempty" type:"Struct"`
	Method             *string                                                          `json:"Method,omitempty" xml:"Method,omitempty"`
	MethodName         *string                                                          `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	MigrateRule        *QueryMgsApirestResponseBodyResultContentValueMigrateRule        `json:"MigrateRule,omitempty" xml:"MigrateRule,omitempty" type:"Struct"`
	MockRule           *QueryMgsApirestResponseBodyResultContentValueMockRule           `json:"MockRule,omitempty" xml:"MockRule,omitempty" type:"Struct"`
	NeedETag           *string                                                          `json:"NeedETag,omitempty" xml:"NeedETag,omitempty"`
	NeedEncrypt        *string                                                          `json:"NeedEncrypt,omitempty" xml:"NeedEncrypt,omitempty"`
	NeedJsonp          *string                                                          `json:"NeedJsonp,omitempty" xml:"NeedJsonp,omitempty"`
	NeedSign           *string                                                          `json:"NeedSign,omitempty" xml:"NeedSign,omitempty"`
	OperationType      *string                                                          `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	ParamGetMethod     *string                                                          `json:"ParamGetMethod,omitempty" xml:"ParamGetMethod,omitempty"`
	Path               *string                                                          `json:"Path,omitempty" xml:"Path,omitempty"`
	RequestBodyModel   *string                                                          `json:"RequestBodyModel,omitempty" xml:"RequestBodyModel,omitempty"`
	RequestParams      []*QueryMgsApirestResponseBodyResultContentValueRequestParams    `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Repeated"`
	ResponseBodyModel  *string                                                          `json:"ResponseBodyModel,omitempty" xml:"ResponseBodyModel,omitempty"`
	SysId              *int64                                                           `json:"SysId,omitempty" xml:"SysId,omitempty"`
	SysName            *string                                                          `json:"SysName,omitempty" xml:"SysName,omitempty"`
	Timeout            *string                                                          `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	WorkspaceId        *string                                                          `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMgsApirestResponseBodyResultContentValue) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApirestResponseBodyResultContentValue) GoString() string {
	return s.String()
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetApiInvoker() *QueryMgsApirestResponseBodyResultContentValueApiInvoker {
	return s.ApiInvoker
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetApiName() *string {
	return s.ApiName
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetApiStatus() *string {
	return s.ApiStatus
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetApiType() *string {
	return s.ApiType
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetAppId() *string {
	return s.AppId
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetAuthRuleName() *string {
	return s.AuthRuleName
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetCacheRule() *QueryMgsApirestResponseBodyResultContentValueCacheRule {
	return s.CacheRule
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetCharset() *string {
	return s.Charset
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetCircuitBreakerRule() *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule {
	return s.CircuitBreakerRule
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetContentType() *string {
	return s.ContentType
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetDefaultLimitRule() *QueryMgsApirestResponseBodyResultContentValueDefaultLimitRule {
	return s.DefaultLimitRule
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetDescription() *string {
	return s.Description
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetHeaderRule() []*QueryMgsApirestResponseBodyResultContentValueHeaderRule {
	return s.HeaderRule
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetHeaderRules() []*QueryMgsApirestResponseBodyResultContentValueHeaderRules {
	return s.HeaderRules
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetHost() *string {
	return s.Host
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetId() *int64 {
	return s.Id
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetInterfaceType() *string {
	return s.InterfaceType
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetLimitRule() *QueryMgsApirestResponseBodyResultContentValueLimitRule {
	return s.LimitRule
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetMethod() *string {
	return s.Method
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetMethodName() *string {
	return s.MethodName
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetMigrateRule() *QueryMgsApirestResponseBodyResultContentValueMigrateRule {
	return s.MigrateRule
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetMockRule() *QueryMgsApirestResponseBodyResultContentValueMockRule {
	return s.MockRule
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetNeedETag() *string {
	return s.NeedETag
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetNeedEncrypt() *string {
	return s.NeedEncrypt
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetNeedJsonp() *string {
	return s.NeedJsonp
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetNeedSign() *string {
	return s.NeedSign
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetOperationType() *string {
	return s.OperationType
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetParamGetMethod() *string {
	return s.ParamGetMethod
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetPath() *string {
	return s.Path
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetRequestBodyModel() *string {
	return s.RequestBodyModel
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetRequestParams() []*QueryMgsApirestResponseBodyResultContentValueRequestParams {
	return s.RequestParams
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetResponseBodyModel() *string {
	return s.ResponseBodyModel
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetSysId() *int64 {
	return s.SysId
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetSysName() *string {
	return s.SysName
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetTimeout() *string {
	return s.Timeout
}

func (s *QueryMgsApirestResponseBodyResultContentValue) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetApiInvoker(v *QueryMgsApirestResponseBodyResultContentValueApiInvoker) *QueryMgsApirestResponseBodyResultContentValue {
	s.ApiInvoker = v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetApiName(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.ApiName = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetApiStatus(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.ApiStatus = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetApiType(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.ApiType = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetAppId(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.AppId = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetAuthRuleName(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.AuthRuleName = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetCacheRule(v *QueryMgsApirestResponseBodyResultContentValueCacheRule) *QueryMgsApirestResponseBodyResultContentValue {
	s.CacheRule = v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetCharset(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.Charset = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetCircuitBreakerRule(v *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) *QueryMgsApirestResponseBodyResultContentValue {
	s.CircuitBreakerRule = v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetContentType(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.ContentType = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetDefaultLimitRule(v *QueryMgsApirestResponseBodyResultContentValueDefaultLimitRule) *QueryMgsApirestResponseBodyResultContentValue {
	s.DefaultLimitRule = v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetDescription(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.Description = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetGmtCreate(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.GmtCreate = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetGmtModified(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.GmtModified = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetHeaderRule(v []*QueryMgsApirestResponseBodyResultContentValueHeaderRule) *QueryMgsApirestResponseBodyResultContentValue {
	s.HeaderRule = v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetHeaderRules(v []*QueryMgsApirestResponseBodyResultContentValueHeaderRules) *QueryMgsApirestResponseBodyResultContentValue {
	s.HeaderRules = v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetHost(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.Host = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetId(v int64) *QueryMgsApirestResponseBodyResultContentValue {
	s.Id = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetInterfaceType(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.InterfaceType = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetLimitRule(v *QueryMgsApirestResponseBodyResultContentValueLimitRule) *QueryMgsApirestResponseBodyResultContentValue {
	s.LimitRule = v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetMethod(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.Method = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetMethodName(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.MethodName = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetMigrateRule(v *QueryMgsApirestResponseBodyResultContentValueMigrateRule) *QueryMgsApirestResponseBodyResultContentValue {
	s.MigrateRule = v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetMockRule(v *QueryMgsApirestResponseBodyResultContentValueMockRule) *QueryMgsApirestResponseBodyResultContentValue {
	s.MockRule = v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetNeedETag(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.NeedETag = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetNeedEncrypt(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.NeedEncrypt = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetNeedJsonp(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.NeedJsonp = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetNeedSign(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.NeedSign = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetOperationType(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.OperationType = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetParamGetMethod(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.ParamGetMethod = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetPath(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.Path = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetRequestBodyModel(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.RequestBodyModel = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetRequestParams(v []*QueryMgsApirestResponseBodyResultContentValueRequestParams) *QueryMgsApirestResponseBodyResultContentValue {
	s.RequestParams = v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetResponseBodyModel(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.ResponseBodyModel = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetSysId(v int64) *QueryMgsApirestResponseBodyResultContentValue {
	s.SysId = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetSysName(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.SysName = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetTimeout(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.Timeout = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) SetWorkspaceId(v string) *QueryMgsApirestResponseBodyResultContentValue {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValue) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApirestResponseBodyResultContentValueApiInvoker struct {
	HttpInvoker *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker `json:"HttpInvoker,omitempty" xml:"HttpInvoker,omitempty" type:"Struct"`
	RpcInvoker  *string                                                             `json:"RpcInvoker,omitempty" xml:"RpcInvoker,omitempty"`
}

func (s QueryMgsApirestResponseBodyResultContentValueApiInvoker) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApirestResponseBodyResultContentValueApiInvoker) GoString() string {
	return s.String()
}

func (s *QueryMgsApirestResponseBodyResultContentValueApiInvoker) GetHttpInvoker() *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker {
	return s.HttpInvoker
}

func (s *QueryMgsApirestResponseBodyResultContentValueApiInvoker) GetRpcInvoker() *string {
	return s.RpcInvoker
}

func (s *QueryMgsApirestResponseBodyResultContentValueApiInvoker) SetHttpInvoker(v *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker) *QueryMgsApirestResponseBodyResultContentValueApiInvoker {
	s.HttpInvoker = v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueApiInvoker) SetRpcInvoker(v string) *QueryMgsApirestResponseBodyResultContentValueApiInvoker {
	s.RpcInvoker = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueApiInvoker) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker struct {
	Charset     *string `json:"Charset,omitempty" xml:"Charset,omitempty"`
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	Host        *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Method      *string `json:"Method,omitempty" xml:"Method,omitempty"`
	Path        *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker) GoString() string {
	return s.String()
}

func (s *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker) GetCharset() *string {
	return s.Charset
}

func (s *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker) GetContentType() *string {
	return s.ContentType
}

func (s *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker) GetHost() *string {
	return s.Host
}

func (s *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker) GetMethod() *string {
	return s.Method
}

func (s *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker) GetPath() *string {
	return s.Path
}

func (s *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker) SetCharset(v string) *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker {
	s.Charset = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker) SetContentType(v string) *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker {
	s.ContentType = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker) SetHost(v string) *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker {
	s.Host = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker) SetMethod(v string) *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker {
	s.Method = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker) SetPath(v string) *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker {
	s.Path = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueApiInvokerHttpInvoker) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApirestResponseBodyResultContentValueCacheRule struct {
	CacheKey  *string `json:"CacheKey,omitempty" xml:"CacheKey,omitempty"`
	NeedCache *bool   `json:"NeedCache,omitempty" xml:"NeedCache,omitempty"`
	Ttl       *int64  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s QueryMgsApirestResponseBodyResultContentValueCacheRule) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApirestResponseBodyResultContentValueCacheRule) GoString() string {
	return s.String()
}

func (s *QueryMgsApirestResponseBodyResultContentValueCacheRule) GetCacheKey() *string {
	return s.CacheKey
}

func (s *QueryMgsApirestResponseBodyResultContentValueCacheRule) GetNeedCache() *bool {
	return s.NeedCache
}

func (s *QueryMgsApirestResponseBodyResultContentValueCacheRule) GetTtl() *int64 {
	return s.Ttl
}

func (s *QueryMgsApirestResponseBodyResultContentValueCacheRule) SetCacheKey(v string) *QueryMgsApirestResponseBodyResultContentValueCacheRule {
	s.CacheKey = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueCacheRule) SetNeedCache(v bool) *QueryMgsApirestResponseBodyResultContentValueCacheRule {
	s.NeedCache = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueCacheRule) SetTtl(v int64) *QueryMgsApirestResponseBodyResultContentValueCacheRule {
	s.Ttl = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueCacheRule) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule struct {
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

func (s QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) GoString() string {
	return s.String()
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) GetAppId() *string {
	return s.AppId
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) GetDefaultResponse() *string {
	return s.DefaultResponse
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) GetErrorThreshold() *int64 {
	return s.ErrorThreshold
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) GetId() *int64 {
	return s.Id
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) GetModel() *string {
	return s.Model
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) GetOpenTimeoutSeconds() *int64 {
	return s.OpenTimeoutSeconds
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) GetSlowRatioThreshold() *float64 {
	return s.SlowRatioThreshold
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) GetSwitchStatus() *string {
	return s.SwitchStatus
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) GetWindowsInSeconds() *int64 {
	return s.WindowsInSeconds
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) SetAppId(v string) *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule {
	s.AppId = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) SetDefaultResponse(v string) *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule {
	s.DefaultResponse = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) SetErrorThreshold(v int64) *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule {
	s.ErrorThreshold = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) SetId(v int64) *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule {
	s.Id = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) SetModel(v string) *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule {
	s.Model = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) SetOpenTimeoutSeconds(v int64) *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule {
	s.OpenTimeoutSeconds = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) SetSlowRatioThreshold(v float64) *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule {
	s.SlowRatioThreshold = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) SetSwitchStatus(v string) *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule {
	s.SwitchStatus = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) SetWindowsInSeconds(v int64) *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule {
	s.WindowsInSeconds = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) SetWorkspaceId(v string) *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueCircuitBreakerRule) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApirestResponseBodyResultContentValueDefaultLimitRule struct {
	ConfigId     *int32 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	DefaultLimit *bool  `json:"DefaultLimit,omitempty" xml:"DefaultLimit,omitempty"`
}

func (s QueryMgsApirestResponseBodyResultContentValueDefaultLimitRule) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApirestResponseBodyResultContentValueDefaultLimitRule) GoString() string {
	return s.String()
}

func (s *QueryMgsApirestResponseBodyResultContentValueDefaultLimitRule) GetConfigId() *int32 {
	return s.ConfigId
}

func (s *QueryMgsApirestResponseBodyResultContentValueDefaultLimitRule) GetDefaultLimit() *bool {
	return s.DefaultLimit
}

func (s *QueryMgsApirestResponseBodyResultContentValueDefaultLimitRule) SetConfigId(v int32) *QueryMgsApirestResponseBodyResultContentValueDefaultLimitRule {
	s.ConfigId = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueDefaultLimitRule) SetDefaultLimit(v bool) *QueryMgsApirestResponseBodyResultContentValueDefaultLimitRule {
	s.DefaultLimit = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueDefaultLimitRule) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApirestResponseBodyResultContentValueHeaderRule struct {
	HeaderKey *string `json:"HeaderKey,omitempty" xml:"HeaderKey,omitempty"`
	Location  *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryMgsApirestResponseBodyResultContentValueHeaderRule) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApirestResponseBodyResultContentValueHeaderRule) GoString() string {
	return s.String()
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRule) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRule) GetLocation() *string {
	return s.Location
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRule) GetType() *string {
	return s.Type
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRule) GetValue() *string {
	return s.Value
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRule) SetHeaderKey(v string) *QueryMgsApirestResponseBodyResultContentValueHeaderRule {
	s.HeaderKey = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRule) SetLocation(v string) *QueryMgsApirestResponseBodyResultContentValueHeaderRule {
	s.Location = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRule) SetType(v string) *QueryMgsApirestResponseBodyResultContentValueHeaderRule {
	s.Type = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRule) SetValue(v string) *QueryMgsApirestResponseBodyResultContentValueHeaderRule {
	s.Value = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRule) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApirestResponseBodyResultContentValueHeaderRules struct {
	HeaderKey *string `json:"HeaderKey,omitempty" xml:"HeaderKey,omitempty"`
	Location  *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryMgsApirestResponseBodyResultContentValueHeaderRules) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApirestResponseBodyResultContentValueHeaderRules) GoString() string {
	return s.String()
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRules) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRules) GetLocation() *string {
	return s.Location
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRules) GetType() *string {
	return s.Type
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRules) GetValue() *string {
	return s.Value
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRules) SetHeaderKey(v string) *QueryMgsApirestResponseBodyResultContentValueHeaderRules {
	s.HeaderKey = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRules) SetLocation(v string) *QueryMgsApirestResponseBodyResultContentValueHeaderRules {
	s.Location = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRules) SetType(v string) *QueryMgsApirestResponseBodyResultContentValueHeaderRules {
	s.Type = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRules) SetValue(v string) *QueryMgsApirestResponseBodyResultContentValueHeaderRules {
	s.Value = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueHeaderRules) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApirestResponseBodyResultContentValueLimitRule struct {
	DefaultResponse *string `json:"DefaultResponse,omitempty" xml:"DefaultResponse,omitempty"`
	I18nResponse    *string `json:"I18nResponse,omitempty" xml:"I18nResponse,omitempty"`
	Interval        *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Limit           *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s QueryMgsApirestResponseBodyResultContentValueLimitRule) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApirestResponseBodyResultContentValueLimitRule) GoString() string {
	return s.String()
}

func (s *QueryMgsApirestResponseBodyResultContentValueLimitRule) GetDefaultResponse() *string {
	return s.DefaultResponse
}

func (s *QueryMgsApirestResponseBodyResultContentValueLimitRule) GetI18nResponse() *string {
	return s.I18nResponse
}

func (s *QueryMgsApirestResponseBodyResultContentValueLimitRule) GetInterval() *int64 {
	return s.Interval
}

func (s *QueryMgsApirestResponseBodyResultContentValueLimitRule) GetLimit() *int64 {
	return s.Limit
}

func (s *QueryMgsApirestResponseBodyResultContentValueLimitRule) GetMode() *string {
	return s.Mode
}

func (s *QueryMgsApirestResponseBodyResultContentValueLimitRule) SetDefaultResponse(v string) *QueryMgsApirestResponseBodyResultContentValueLimitRule {
	s.DefaultResponse = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueLimitRule) SetI18nResponse(v string) *QueryMgsApirestResponseBodyResultContentValueLimitRule {
	s.I18nResponse = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueLimitRule) SetInterval(v int64) *QueryMgsApirestResponseBodyResultContentValueLimitRule {
	s.Interval = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueLimitRule) SetLimit(v int64) *QueryMgsApirestResponseBodyResultContentValueLimitRule {
	s.Limit = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueLimitRule) SetMode(v string) *QueryMgsApirestResponseBodyResultContentValueLimitRule {
	s.Mode = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueLimitRule) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApirestResponseBodyResultContentValueMigrateRule struct {
	FlowPercent          *int64  `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
	NeedMigrate          *bool   `json:"NeedMigrate,omitempty" xml:"NeedMigrate,omitempty"`
	NeedSwitchCompletely *bool   `json:"NeedSwitchCompletely,omitempty" xml:"NeedSwitchCompletely,omitempty"`
	SysId                *int64  `json:"SysId,omitempty" xml:"SysId,omitempty"`
	SysName              *string `json:"SysName,omitempty" xml:"SysName,omitempty"`
	UpstreamType         *string `json:"UpstreamType,omitempty" xml:"UpstreamType,omitempty"`
}

func (s QueryMgsApirestResponseBodyResultContentValueMigrateRule) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApirestResponseBodyResultContentValueMigrateRule) GoString() string {
	return s.String()
}

func (s *QueryMgsApirestResponseBodyResultContentValueMigrateRule) GetFlowPercent() *int64 {
	return s.FlowPercent
}

func (s *QueryMgsApirestResponseBodyResultContentValueMigrateRule) GetNeedMigrate() *bool {
	return s.NeedMigrate
}

func (s *QueryMgsApirestResponseBodyResultContentValueMigrateRule) GetNeedSwitchCompletely() *bool {
	return s.NeedSwitchCompletely
}

func (s *QueryMgsApirestResponseBodyResultContentValueMigrateRule) GetSysId() *int64 {
	return s.SysId
}

func (s *QueryMgsApirestResponseBodyResultContentValueMigrateRule) GetSysName() *string {
	return s.SysName
}

func (s *QueryMgsApirestResponseBodyResultContentValueMigrateRule) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *QueryMgsApirestResponseBodyResultContentValueMigrateRule) SetFlowPercent(v int64) *QueryMgsApirestResponseBodyResultContentValueMigrateRule {
	s.FlowPercent = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueMigrateRule) SetNeedMigrate(v bool) *QueryMgsApirestResponseBodyResultContentValueMigrateRule {
	s.NeedMigrate = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueMigrateRule) SetNeedSwitchCompletely(v bool) *QueryMgsApirestResponseBodyResultContentValueMigrateRule {
	s.NeedSwitchCompletely = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueMigrateRule) SetSysId(v int64) *QueryMgsApirestResponseBodyResultContentValueMigrateRule {
	s.SysId = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueMigrateRule) SetSysName(v string) *QueryMgsApirestResponseBodyResultContentValueMigrateRule {
	s.SysName = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueMigrateRule) SetUpstreamType(v string) *QueryMgsApirestResponseBodyResultContentValueMigrateRule {
	s.UpstreamType = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueMigrateRule) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApirestResponseBodyResultContentValueMockRule struct {
	MockData   *string `json:"MockData,omitempty" xml:"MockData,omitempty"`
	NeedMock   *bool   `json:"NeedMock,omitempty" xml:"NeedMock,omitempty"`
	Percentage *int64  `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryMgsApirestResponseBodyResultContentValueMockRule) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApirestResponseBodyResultContentValueMockRule) GoString() string {
	return s.String()
}

func (s *QueryMgsApirestResponseBodyResultContentValueMockRule) GetMockData() *string {
	return s.MockData
}

func (s *QueryMgsApirestResponseBodyResultContentValueMockRule) GetNeedMock() *bool {
	return s.NeedMock
}

func (s *QueryMgsApirestResponseBodyResultContentValueMockRule) GetPercentage() *int64 {
	return s.Percentage
}

func (s *QueryMgsApirestResponseBodyResultContentValueMockRule) GetType() *string {
	return s.Type
}

func (s *QueryMgsApirestResponseBodyResultContentValueMockRule) SetMockData(v string) *QueryMgsApirestResponseBodyResultContentValueMockRule {
	s.MockData = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueMockRule) SetNeedMock(v bool) *QueryMgsApirestResponseBodyResultContentValueMockRule {
	s.NeedMock = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueMockRule) SetPercentage(v int64) *QueryMgsApirestResponseBodyResultContentValueMockRule {
	s.Percentage = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueMockRule) SetType(v string) *QueryMgsApirestResponseBodyResultContentValueMockRule {
	s.Type = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueMockRule) Validate() error {
	return dara.Validate(s)
}

type QueryMgsApirestResponseBodyResultContentValueRequestParams struct {
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

func (s QueryMgsApirestResponseBodyResultContentValueRequestParams) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApirestResponseBodyResultContentValueRequestParams) GoString() string {
	return s.String()
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) GetApiId() *string {
	return s.ApiId
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) GetAppId() *string {
	return s.AppId
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) GetDescription() *string {
	return s.Description
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) GetId() *int64 {
	return s.Id
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) GetLocation() *string {
	return s.Location
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) GetName() *string {
	return s.Name
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) GetRefType() *string {
	return s.RefType
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) GetType() *string {
	return s.Type
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) SetApiId(v string) *QueryMgsApirestResponseBodyResultContentValueRequestParams {
	s.ApiId = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) SetAppId(v string) *QueryMgsApirestResponseBodyResultContentValueRequestParams {
	s.AppId = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) SetDefaultValue(v string) *QueryMgsApirestResponseBodyResultContentValueRequestParams {
	s.DefaultValue = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) SetDescription(v string) *QueryMgsApirestResponseBodyResultContentValueRequestParams {
	s.Description = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) SetId(v int64) *QueryMgsApirestResponseBodyResultContentValueRequestParams {
	s.Id = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) SetLocation(v string) *QueryMgsApirestResponseBodyResultContentValueRequestParams {
	s.Location = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) SetName(v string) *QueryMgsApirestResponseBodyResultContentValueRequestParams {
	s.Name = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) SetRefType(v string) *QueryMgsApirestResponseBodyResultContentValueRequestParams {
	s.RefType = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) SetType(v string) *QueryMgsApirestResponseBodyResultContentValueRequestParams {
	s.Type = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) SetWorkspaceId(v string) *QueryMgsApirestResponseBodyResultContentValueRequestParams {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMgsApirestResponseBodyResultContentValueRequestParams) Validate() error {
	return dara.Validate(s)
}
