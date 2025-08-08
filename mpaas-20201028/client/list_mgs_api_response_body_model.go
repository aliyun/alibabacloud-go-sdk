// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMgsApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListMgsApiResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMgsApiResponseBody
	GetResultCode() *string
	SetResultContent(v *ListMgsApiResponseBodyResultContent) *ListMgsApiResponseBody
	GetResultContent() *ListMgsApiResponseBodyResultContent
	SetResultMessage(v string) *ListMgsApiResponseBody
	GetResultMessage() *string
}

type ListMgsApiResponseBody struct {
	RequestId     *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                              `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *ListMgsApiResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                              `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMgsApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMgsApiResponseBody) GoString() string {
	return s.String()
}

func (s *ListMgsApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMgsApiResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMgsApiResponseBody) GetResultContent() *ListMgsApiResponseBodyResultContent {
	return s.ResultContent
}

func (s *ListMgsApiResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMgsApiResponseBody) SetRequestId(v string) *ListMgsApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMgsApiResponseBody) SetResultCode(v string) *ListMgsApiResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMgsApiResponseBody) SetResultContent(v *ListMgsApiResponseBodyResultContent) *ListMgsApiResponseBody {
	s.ResultContent = v
	return s
}

func (s *ListMgsApiResponseBody) SetResultMessage(v string) *ListMgsApiResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMgsApiResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMgsApiResponseBodyResultContent struct {
	ErrorMessage *string                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	Value        []*ListMgsApiResponseBodyResultContentValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListMgsApiResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s ListMgsApiResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *ListMgsApiResponseBodyResultContent) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListMgsApiResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *ListMgsApiResponseBodyResultContent) GetValue() []*ListMgsApiResponseBodyResultContentValue {
	return s.Value
}

func (s *ListMgsApiResponseBodyResultContent) SetErrorMessage(v string) *ListMgsApiResponseBodyResultContent {
	s.ErrorMessage = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContent) SetSuccess(v bool) *ListMgsApiResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContent) SetValue(v []*ListMgsApiResponseBodyResultContentValue) *ListMgsApiResponseBodyResultContent {
	s.Value = v
	return s
}

func (s *ListMgsApiResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}

type ListMgsApiResponseBodyResultContentValue struct {
	ApiInvoker         *ListMgsApiResponseBodyResultContentValueApiInvoker         `json:"ApiInvoker,omitempty" xml:"ApiInvoker,omitempty" type:"Struct"`
	ApiName            *string                                                     `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	ApiStatus          *string                                                     `json:"ApiStatus,omitempty" xml:"ApiStatus,omitempty"`
	ApiType            *string                                                     `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	AppId              *string                                                     `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AuthRuleName       *string                                                     `json:"AuthRuleName,omitempty" xml:"AuthRuleName,omitempty"`
	CacheRule          *ListMgsApiResponseBodyResultContentValueCacheRule          `json:"CacheRule,omitempty" xml:"CacheRule,omitempty" type:"Struct"`
	Charset            *string                                                     `json:"Charset,omitempty" xml:"Charset,omitempty"`
	CircuitBreakerRule *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule `json:"CircuitBreakerRule,omitempty" xml:"CircuitBreakerRule,omitempty" type:"Struct"`
	ContentType        *string                                                     `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	Description        *string                                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate          *string                                                     `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *string                                                     `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HeaderRule         []*ListMgsApiResponseBodyResultContentValueHeaderRule       `json:"HeaderRule,omitempty" xml:"HeaderRule,omitempty" type:"Repeated"`
	HeaderRules        []*ListMgsApiResponseBodyResultContentValueHeaderRules      `json:"HeaderRules,omitempty" xml:"HeaderRules,omitempty" type:"Repeated"`
	Host               *string                                                     `json:"Host,omitempty" xml:"Host,omitempty"`
	Id                 *int64                                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	InterfaceType      *string                                                     `json:"InterfaceType,omitempty" xml:"InterfaceType,omitempty"`
	LimitRule          *ListMgsApiResponseBodyResultContentValueLimitRule          `json:"LimitRule,omitempty" xml:"LimitRule,omitempty" type:"Struct"`
	Method             *string                                                     `json:"Method,omitempty" xml:"Method,omitempty"`
	MethodName         *string                                                     `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	MigrateRule        *ListMgsApiResponseBodyResultContentValueMigrateRule        `json:"MigrateRule,omitempty" xml:"MigrateRule,omitempty" type:"Struct"`
	MockRule           *ListMgsApiResponseBodyResultContentValueMockRule           `json:"MockRule,omitempty" xml:"MockRule,omitempty" type:"Struct"`
	NeedETag           *string                                                     `json:"NeedETag,omitempty" xml:"NeedETag,omitempty"`
	NeedEncrypt        *string                                                     `json:"NeedEncrypt,omitempty" xml:"NeedEncrypt,omitempty"`
	NeedJsonp          *string                                                     `json:"NeedJsonp,omitempty" xml:"NeedJsonp,omitempty"`
	NeedSign           *string                                                     `json:"NeedSign,omitempty" xml:"NeedSign,omitempty"`
	OperationType      *string                                                     `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	ParamGetMethod     *string                                                     `json:"ParamGetMethod,omitempty" xml:"ParamGetMethod,omitempty"`
	Path               *string                                                     `json:"Path,omitempty" xml:"Path,omitempty"`
	RequestBodyModel   *string                                                     `json:"RequestBodyModel,omitempty" xml:"RequestBodyModel,omitempty"`
	RequestParams      []*ListMgsApiResponseBodyResultContentValueRequestParams    `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Repeated"`
	ResponseBodyModel  *string                                                     `json:"ResponseBodyModel,omitempty" xml:"ResponseBodyModel,omitempty"`
	SysId              *int64                                                      `json:"SysId,omitempty" xml:"SysId,omitempty"`
	SysName            *string                                                     `json:"SysName,omitempty" xml:"SysName,omitempty"`
	Timeout            *string                                                     `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	WorkspaceId        *string                                                     `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMgsApiResponseBodyResultContentValue) String() string {
	return dara.Prettify(s)
}

func (s ListMgsApiResponseBodyResultContentValue) GoString() string {
	return s.String()
}

func (s *ListMgsApiResponseBodyResultContentValue) GetApiInvoker() *ListMgsApiResponseBodyResultContentValueApiInvoker {
	return s.ApiInvoker
}

func (s *ListMgsApiResponseBodyResultContentValue) GetApiName() *string {
	return s.ApiName
}

func (s *ListMgsApiResponseBodyResultContentValue) GetApiStatus() *string {
	return s.ApiStatus
}

func (s *ListMgsApiResponseBodyResultContentValue) GetApiType() *string {
	return s.ApiType
}

func (s *ListMgsApiResponseBodyResultContentValue) GetAppId() *string {
	return s.AppId
}

func (s *ListMgsApiResponseBodyResultContentValue) GetAuthRuleName() *string {
	return s.AuthRuleName
}

func (s *ListMgsApiResponseBodyResultContentValue) GetCacheRule() *ListMgsApiResponseBodyResultContentValueCacheRule {
	return s.CacheRule
}

func (s *ListMgsApiResponseBodyResultContentValue) GetCharset() *string {
	return s.Charset
}

func (s *ListMgsApiResponseBodyResultContentValue) GetCircuitBreakerRule() *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule {
	return s.CircuitBreakerRule
}

func (s *ListMgsApiResponseBodyResultContentValue) GetContentType() *string {
	return s.ContentType
}

func (s *ListMgsApiResponseBodyResultContentValue) GetDescription() *string {
	return s.Description
}

func (s *ListMgsApiResponseBodyResultContentValue) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMgsApiResponseBodyResultContentValue) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMgsApiResponseBodyResultContentValue) GetHeaderRule() []*ListMgsApiResponseBodyResultContentValueHeaderRule {
	return s.HeaderRule
}

func (s *ListMgsApiResponseBodyResultContentValue) GetHeaderRules() []*ListMgsApiResponseBodyResultContentValueHeaderRules {
	return s.HeaderRules
}

func (s *ListMgsApiResponseBodyResultContentValue) GetHost() *string {
	return s.Host
}

func (s *ListMgsApiResponseBodyResultContentValue) GetId() *int64 {
	return s.Id
}

func (s *ListMgsApiResponseBodyResultContentValue) GetInterfaceType() *string {
	return s.InterfaceType
}

func (s *ListMgsApiResponseBodyResultContentValue) GetLimitRule() *ListMgsApiResponseBodyResultContentValueLimitRule {
	return s.LimitRule
}

func (s *ListMgsApiResponseBodyResultContentValue) GetMethod() *string {
	return s.Method
}

func (s *ListMgsApiResponseBodyResultContentValue) GetMethodName() *string {
	return s.MethodName
}

func (s *ListMgsApiResponseBodyResultContentValue) GetMigrateRule() *ListMgsApiResponseBodyResultContentValueMigrateRule {
	return s.MigrateRule
}

func (s *ListMgsApiResponseBodyResultContentValue) GetMockRule() *ListMgsApiResponseBodyResultContentValueMockRule {
	return s.MockRule
}

func (s *ListMgsApiResponseBodyResultContentValue) GetNeedETag() *string {
	return s.NeedETag
}

func (s *ListMgsApiResponseBodyResultContentValue) GetNeedEncrypt() *string {
	return s.NeedEncrypt
}

func (s *ListMgsApiResponseBodyResultContentValue) GetNeedJsonp() *string {
	return s.NeedJsonp
}

func (s *ListMgsApiResponseBodyResultContentValue) GetNeedSign() *string {
	return s.NeedSign
}

func (s *ListMgsApiResponseBodyResultContentValue) GetOperationType() *string {
	return s.OperationType
}

func (s *ListMgsApiResponseBodyResultContentValue) GetParamGetMethod() *string {
	return s.ParamGetMethod
}

func (s *ListMgsApiResponseBodyResultContentValue) GetPath() *string {
	return s.Path
}

func (s *ListMgsApiResponseBodyResultContentValue) GetRequestBodyModel() *string {
	return s.RequestBodyModel
}

func (s *ListMgsApiResponseBodyResultContentValue) GetRequestParams() []*ListMgsApiResponseBodyResultContentValueRequestParams {
	return s.RequestParams
}

func (s *ListMgsApiResponseBodyResultContentValue) GetResponseBodyModel() *string {
	return s.ResponseBodyModel
}

func (s *ListMgsApiResponseBodyResultContentValue) GetSysId() *int64 {
	return s.SysId
}

func (s *ListMgsApiResponseBodyResultContentValue) GetSysName() *string {
	return s.SysName
}

func (s *ListMgsApiResponseBodyResultContentValue) GetTimeout() *string {
	return s.Timeout
}

func (s *ListMgsApiResponseBodyResultContentValue) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMgsApiResponseBodyResultContentValue) SetApiInvoker(v *ListMgsApiResponseBodyResultContentValueApiInvoker) *ListMgsApiResponseBodyResultContentValue {
	s.ApiInvoker = v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetApiName(v string) *ListMgsApiResponseBodyResultContentValue {
	s.ApiName = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetApiStatus(v string) *ListMgsApiResponseBodyResultContentValue {
	s.ApiStatus = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetApiType(v string) *ListMgsApiResponseBodyResultContentValue {
	s.ApiType = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetAppId(v string) *ListMgsApiResponseBodyResultContentValue {
	s.AppId = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetAuthRuleName(v string) *ListMgsApiResponseBodyResultContentValue {
	s.AuthRuleName = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetCacheRule(v *ListMgsApiResponseBodyResultContentValueCacheRule) *ListMgsApiResponseBodyResultContentValue {
	s.CacheRule = v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetCharset(v string) *ListMgsApiResponseBodyResultContentValue {
	s.Charset = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetCircuitBreakerRule(v *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) *ListMgsApiResponseBodyResultContentValue {
	s.CircuitBreakerRule = v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetContentType(v string) *ListMgsApiResponseBodyResultContentValue {
	s.ContentType = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetDescription(v string) *ListMgsApiResponseBodyResultContentValue {
	s.Description = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetGmtCreate(v string) *ListMgsApiResponseBodyResultContentValue {
	s.GmtCreate = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetGmtModified(v string) *ListMgsApiResponseBodyResultContentValue {
	s.GmtModified = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetHeaderRule(v []*ListMgsApiResponseBodyResultContentValueHeaderRule) *ListMgsApiResponseBodyResultContentValue {
	s.HeaderRule = v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetHeaderRules(v []*ListMgsApiResponseBodyResultContentValueHeaderRules) *ListMgsApiResponseBodyResultContentValue {
	s.HeaderRules = v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetHost(v string) *ListMgsApiResponseBodyResultContentValue {
	s.Host = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetId(v int64) *ListMgsApiResponseBodyResultContentValue {
	s.Id = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetInterfaceType(v string) *ListMgsApiResponseBodyResultContentValue {
	s.InterfaceType = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetLimitRule(v *ListMgsApiResponseBodyResultContentValueLimitRule) *ListMgsApiResponseBodyResultContentValue {
	s.LimitRule = v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetMethod(v string) *ListMgsApiResponseBodyResultContentValue {
	s.Method = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetMethodName(v string) *ListMgsApiResponseBodyResultContentValue {
	s.MethodName = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetMigrateRule(v *ListMgsApiResponseBodyResultContentValueMigrateRule) *ListMgsApiResponseBodyResultContentValue {
	s.MigrateRule = v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetMockRule(v *ListMgsApiResponseBodyResultContentValueMockRule) *ListMgsApiResponseBodyResultContentValue {
	s.MockRule = v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetNeedETag(v string) *ListMgsApiResponseBodyResultContentValue {
	s.NeedETag = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetNeedEncrypt(v string) *ListMgsApiResponseBodyResultContentValue {
	s.NeedEncrypt = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetNeedJsonp(v string) *ListMgsApiResponseBodyResultContentValue {
	s.NeedJsonp = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetNeedSign(v string) *ListMgsApiResponseBodyResultContentValue {
	s.NeedSign = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetOperationType(v string) *ListMgsApiResponseBodyResultContentValue {
	s.OperationType = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetParamGetMethod(v string) *ListMgsApiResponseBodyResultContentValue {
	s.ParamGetMethod = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetPath(v string) *ListMgsApiResponseBodyResultContentValue {
	s.Path = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetRequestBodyModel(v string) *ListMgsApiResponseBodyResultContentValue {
	s.RequestBodyModel = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetRequestParams(v []*ListMgsApiResponseBodyResultContentValueRequestParams) *ListMgsApiResponseBodyResultContentValue {
	s.RequestParams = v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetResponseBodyModel(v string) *ListMgsApiResponseBodyResultContentValue {
	s.ResponseBodyModel = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetSysId(v int64) *ListMgsApiResponseBodyResultContentValue {
	s.SysId = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetSysName(v string) *ListMgsApiResponseBodyResultContentValue {
	s.SysName = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetTimeout(v string) *ListMgsApiResponseBodyResultContentValue {
	s.Timeout = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) SetWorkspaceId(v string) *ListMgsApiResponseBodyResultContentValue {
	s.WorkspaceId = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValue) Validate() error {
	return dara.Validate(s)
}

type ListMgsApiResponseBodyResultContentValueApiInvoker struct {
	HttpInvoker *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker `json:"HttpInvoker,omitempty" xml:"HttpInvoker,omitempty" type:"Struct"`
	RpcInvoker  *string                                                        `json:"RpcInvoker,omitempty" xml:"RpcInvoker,omitempty"`
}

func (s ListMgsApiResponseBodyResultContentValueApiInvoker) String() string {
	return dara.Prettify(s)
}

func (s ListMgsApiResponseBodyResultContentValueApiInvoker) GoString() string {
	return s.String()
}

func (s *ListMgsApiResponseBodyResultContentValueApiInvoker) GetHttpInvoker() *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker {
	return s.HttpInvoker
}

func (s *ListMgsApiResponseBodyResultContentValueApiInvoker) GetRpcInvoker() *string {
	return s.RpcInvoker
}

func (s *ListMgsApiResponseBodyResultContentValueApiInvoker) SetHttpInvoker(v *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker) *ListMgsApiResponseBodyResultContentValueApiInvoker {
	s.HttpInvoker = v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueApiInvoker) SetRpcInvoker(v string) *ListMgsApiResponseBodyResultContentValueApiInvoker {
	s.RpcInvoker = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueApiInvoker) Validate() error {
	return dara.Validate(s)
}

type ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker struct {
	Charset     *string `json:"Charset,omitempty" xml:"Charset,omitempty"`
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	Host        *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Method      *string `json:"Method,omitempty" xml:"Method,omitempty"`
	Path        *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker) String() string {
	return dara.Prettify(s)
}

func (s ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker) GoString() string {
	return s.String()
}

func (s *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker) GetCharset() *string {
	return s.Charset
}

func (s *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker) GetContentType() *string {
	return s.ContentType
}

func (s *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker) GetHost() *string {
	return s.Host
}

func (s *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker) GetMethod() *string {
	return s.Method
}

func (s *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker) GetPath() *string {
	return s.Path
}

func (s *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker) SetCharset(v string) *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker {
	s.Charset = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker) SetContentType(v string) *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker {
	s.ContentType = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker) SetHost(v string) *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker {
	s.Host = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker) SetMethod(v string) *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker {
	s.Method = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker) SetPath(v string) *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker {
	s.Path = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueApiInvokerHttpInvoker) Validate() error {
	return dara.Validate(s)
}

type ListMgsApiResponseBodyResultContentValueCacheRule struct {
	CacheKey  *string `json:"CacheKey,omitempty" xml:"CacheKey,omitempty"`
	NeedCache *bool   `json:"NeedCache,omitempty" xml:"NeedCache,omitempty"`
	Ttl       *int64  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s ListMgsApiResponseBodyResultContentValueCacheRule) String() string {
	return dara.Prettify(s)
}

func (s ListMgsApiResponseBodyResultContentValueCacheRule) GoString() string {
	return s.String()
}

func (s *ListMgsApiResponseBodyResultContentValueCacheRule) GetCacheKey() *string {
	return s.CacheKey
}

func (s *ListMgsApiResponseBodyResultContentValueCacheRule) GetNeedCache() *bool {
	return s.NeedCache
}

func (s *ListMgsApiResponseBodyResultContentValueCacheRule) GetTtl() *int64 {
	return s.Ttl
}

func (s *ListMgsApiResponseBodyResultContentValueCacheRule) SetCacheKey(v string) *ListMgsApiResponseBodyResultContentValueCacheRule {
	s.CacheKey = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueCacheRule) SetNeedCache(v bool) *ListMgsApiResponseBodyResultContentValueCacheRule {
	s.NeedCache = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueCacheRule) SetTtl(v int64) *ListMgsApiResponseBodyResultContentValueCacheRule {
	s.Ttl = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueCacheRule) Validate() error {
	return dara.Validate(s)
}

type ListMgsApiResponseBodyResultContentValueCircuitBreakerRule struct {
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

func (s ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) String() string {
	return dara.Prettify(s)
}

func (s ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) GoString() string {
	return s.String()
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) GetAppId() *string {
	return s.AppId
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) GetDefaultResponse() *string {
	return s.DefaultResponse
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) GetErrorThreshold() *int64 {
	return s.ErrorThreshold
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) GetId() *int64 {
	return s.Id
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) GetModel() *string {
	return s.Model
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) GetOpenTimeoutSeconds() *int64 {
	return s.OpenTimeoutSeconds
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) GetSlowRatioThreshold() *float64 {
	return s.SlowRatioThreshold
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) GetSwitchStatus() *string {
	return s.SwitchStatus
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) GetWindowsInSeconds() *int64 {
	return s.WindowsInSeconds
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) SetAppId(v string) *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule {
	s.AppId = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) SetDefaultResponse(v string) *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule {
	s.DefaultResponse = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) SetErrorThreshold(v int64) *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule {
	s.ErrorThreshold = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) SetId(v int64) *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule {
	s.Id = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) SetModel(v string) *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule {
	s.Model = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) SetOpenTimeoutSeconds(v int64) *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule {
	s.OpenTimeoutSeconds = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) SetSlowRatioThreshold(v float64) *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule {
	s.SlowRatioThreshold = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) SetSwitchStatus(v string) *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule {
	s.SwitchStatus = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) SetWindowsInSeconds(v int64) *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule {
	s.WindowsInSeconds = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) SetWorkspaceId(v string) *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule {
	s.WorkspaceId = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueCircuitBreakerRule) Validate() error {
	return dara.Validate(s)
}

type ListMgsApiResponseBodyResultContentValueHeaderRule struct {
	HeaderKey *string `json:"HeaderKey,omitempty" xml:"HeaderKey,omitempty"`
	Location  *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListMgsApiResponseBodyResultContentValueHeaderRule) String() string {
	return dara.Prettify(s)
}

func (s ListMgsApiResponseBodyResultContentValueHeaderRule) GoString() string {
	return s.String()
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRule) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRule) GetLocation() *string {
	return s.Location
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRule) GetType() *string {
	return s.Type
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRule) GetValue() *string {
	return s.Value
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRule) SetHeaderKey(v string) *ListMgsApiResponseBodyResultContentValueHeaderRule {
	s.HeaderKey = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRule) SetLocation(v string) *ListMgsApiResponseBodyResultContentValueHeaderRule {
	s.Location = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRule) SetType(v string) *ListMgsApiResponseBodyResultContentValueHeaderRule {
	s.Type = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRule) SetValue(v string) *ListMgsApiResponseBodyResultContentValueHeaderRule {
	s.Value = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRule) Validate() error {
	return dara.Validate(s)
}

type ListMgsApiResponseBodyResultContentValueHeaderRules struct {
	HeaderKey *string `json:"HeaderKey,omitempty" xml:"HeaderKey,omitempty"`
	Location  *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListMgsApiResponseBodyResultContentValueHeaderRules) String() string {
	return dara.Prettify(s)
}

func (s ListMgsApiResponseBodyResultContentValueHeaderRules) GoString() string {
	return s.String()
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRules) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRules) GetLocation() *string {
	return s.Location
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRules) GetType() *string {
	return s.Type
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRules) GetValue() *string {
	return s.Value
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRules) SetHeaderKey(v string) *ListMgsApiResponseBodyResultContentValueHeaderRules {
	s.HeaderKey = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRules) SetLocation(v string) *ListMgsApiResponseBodyResultContentValueHeaderRules {
	s.Location = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRules) SetType(v string) *ListMgsApiResponseBodyResultContentValueHeaderRules {
	s.Type = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRules) SetValue(v string) *ListMgsApiResponseBodyResultContentValueHeaderRules {
	s.Value = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueHeaderRules) Validate() error {
	return dara.Validate(s)
}

type ListMgsApiResponseBodyResultContentValueLimitRule struct {
	DefaultResponse *string `json:"DefaultResponse,omitempty" xml:"DefaultResponse,omitempty"`
	I18nResponse    *string `json:"I18nResponse,omitempty" xml:"I18nResponse,omitempty"`
	Interval        *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Limit           *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ListMgsApiResponseBodyResultContentValueLimitRule) String() string {
	return dara.Prettify(s)
}

func (s ListMgsApiResponseBodyResultContentValueLimitRule) GoString() string {
	return s.String()
}

func (s *ListMgsApiResponseBodyResultContentValueLimitRule) GetDefaultResponse() *string {
	return s.DefaultResponse
}

func (s *ListMgsApiResponseBodyResultContentValueLimitRule) GetI18nResponse() *string {
	return s.I18nResponse
}

func (s *ListMgsApiResponseBodyResultContentValueLimitRule) GetInterval() *int64 {
	return s.Interval
}

func (s *ListMgsApiResponseBodyResultContentValueLimitRule) GetLimit() *int64 {
	return s.Limit
}

func (s *ListMgsApiResponseBodyResultContentValueLimitRule) GetMode() *string {
	return s.Mode
}

func (s *ListMgsApiResponseBodyResultContentValueLimitRule) SetDefaultResponse(v string) *ListMgsApiResponseBodyResultContentValueLimitRule {
	s.DefaultResponse = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueLimitRule) SetI18nResponse(v string) *ListMgsApiResponseBodyResultContentValueLimitRule {
	s.I18nResponse = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueLimitRule) SetInterval(v int64) *ListMgsApiResponseBodyResultContentValueLimitRule {
	s.Interval = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueLimitRule) SetLimit(v int64) *ListMgsApiResponseBodyResultContentValueLimitRule {
	s.Limit = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueLimitRule) SetMode(v string) *ListMgsApiResponseBodyResultContentValueLimitRule {
	s.Mode = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueLimitRule) Validate() error {
	return dara.Validate(s)
}

type ListMgsApiResponseBodyResultContentValueMigrateRule struct {
	FlowPercent          *int64  `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
	NeedMigrate          *bool   `json:"NeedMigrate,omitempty" xml:"NeedMigrate,omitempty"`
	NeedSwitchCompletely *bool   `json:"NeedSwitchCompletely,omitempty" xml:"NeedSwitchCompletely,omitempty"`
	SysId                *int64  `json:"SysId,omitempty" xml:"SysId,omitempty"`
	SysName              *string `json:"SysName,omitempty" xml:"SysName,omitempty"`
	UpstreamType         *string `json:"UpstreamType,omitempty" xml:"UpstreamType,omitempty"`
}

func (s ListMgsApiResponseBodyResultContentValueMigrateRule) String() string {
	return dara.Prettify(s)
}

func (s ListMgsApiResponseBodyResultContentValueMigrateRule) GoString() string {
	return s.String()
}

func (s *ListMgsApiResponseBodyResultContentValueMigrateRule) GetFlowPercent() *int64 {
	return s.FlowPercent
}

func (s *ListMgsApiResponseBodyResultContentValueMigrateRule) GetNeedMigrate() *bool {
	return s.NeedMigrate
}

func (s *ListMgsApiResponseBodyResultContentValueMigrateRule) GetNeedSwitchCompletely() *bool {
	return s.NeedSwitchCompletely
}

func (s *ListMgsApiResponseBodyResultContentValueMigrateRule) GetSysId() *int64 {
	return s.SysId
}

func (s *ListMgsApiResponseBodyResultContentValueMigrateRule) GetSysName() *string {
	return s.SysName
}

func (s *ListMgsApiResponseBodyResultContentValueMigrateRule) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *ListMgsApiResponseBodyResultContentValueMigrateRule) SetFlowPercent(v int64) *ListMgsApiResponseBodyResultContentValueMigrateRule {
	s.FlowPercent = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueMigrateRule) SetNeedMigrate(v bool) *ListMgsApiResponseBodyResultContentValueMigrateRule {
	s.NeedMigrate = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueMigrateRule) SetNeedSwitchCompletely(v bool) *ListMgsApiResponseBodyResultContentValueMigrateRule {
	s.NeedSwitchCompletely = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueMigrateRule) SetSysId(v int64) *ListMgsApiResponseBodyResultContentValueMigrateRule {
	s.SysId = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueMigrateRule) SetSysName(v string) *ListMgsApiResponseBodyResultContentValueMigrateRule {
	s.SysName = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueMigrateRule) SetUpstreamType(v string) *ListMgsApiResponseBodyResultContentValueMigrateRule {
	s.UpstreamType = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueMigrateRule) Validate() error {
	return dara.Validate(s)
}

type ListMgsApiResponseBodyResultContentValueMockRule struct {
	MockData   *string `json:"MockData,omitempty" xml:"MockData,omitempty"`
	NeedMock   *bool   `json:"NeedMock,omitempty" xml:"NeedMock,omitempty"`
	Percentage *int64  `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListMgsApiResponseBodyResultContentValueMockRule) String() string {
	return dara.Prettify(s)
}

func (s ListMgsApiResponseBodyResultContentValueMockRule) GoString() string {
	return s.String()
}

func (s *ListMgsApiResponseBodyResultContentValueMockRule) GetMockData() *string {
	return s.MockData
}

func (s *ListMgsApiResponseBodyResultContentValueMockRule) GetNeedMock() *bool {
	return s.NeedMock
}

func (s *ListMgsApiResponseBodyResultContentValueMockRule) GetPercentage() *int64 {
	return s.Percentage
}

func (s *ListMgsApiResponseBodyResultContentValueMockRule) GetType() *string {
	return s.Type
}

func (s *ListMgsApiResponseBodyResultContentValueMockRule) SetMockData(v string) *ListMgsApiResponseBodyResultContentValueMockRule {
	s.MockData = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueMockRule) SetNeedMock(v bool) *ListMgsApiResponseBodyResultContentValueMockRule {
	s.NeedMock = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueMockRule) SetPercentage(v int64) *ListMgsApiResponseBodyResultContentValueMockRule {
	s.Percentage = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueMockRule) SetType(v string) *ListMgsApiResponseBodyResultContentValueMockRule {
	s.Type = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueMockRule) Validate() error {
	return dara.Validate(s)
}

type ListMgsApiResponseBodyResultContentValueRequestParams struct {
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

func (s ListMgsApiResponseBodyResultContentValueRequestParams) String() string {
	return dara.Prettify(s)
}

func (s ListMgsApiResponseBodyResultContentValueRequestParams) GoString() string {
	return s.String()
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) GetApiId() *string {
	return s.ApiId
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) GetAppId() *string {
	return s.AppId
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) GetDescription() *string {
	return s.Description
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) GetId() *int64 {
	return s.Id
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) GetLocation() *string {
	return s.Location
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) GetName() *string {
	return s.Name
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) GetRefType() *string {
	return s.RefType
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) GetType() *string {
	return s.Type
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) SetApiId(v string) *ListMgsApiResponseBodyResultContentValueRequestParams {
	s.ApiId = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) SetAppId(v string) *ListMgsApiResponseBodyResultContentValueRequestParams {
	s.AppId = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) SetDefaultValue(v string) *ListMgsApiResponseBodyResultContentValueRequestParams {
	s.DefaultValue = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) SetDescription(v string) *ListMgsApiResponseBodyResultContentValueRequestParams {
	s.Description = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) SetId(v int64) *ListMgsApiResponseBodyResultContentValueRequestParams {
	s.Id = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) SetLocation(v string) *ListMgsApiResponseBodyResultContentValueRequestParams {
	s.Location = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) SetName(v string) *ListMgsApiResponseBodyResultContentValueRequestParams {
	s.Name = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) SetRefType(v string) *ListMgsApiResponseBodyResultContentValueRequestParams {
	s.RefType = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) SetType(v string) *ListMgsApiResponseBodyResultContentValueRequestParams {
	s.Type = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) SetWorkspaceId(v string) *ListMgsApiResponseBodyResultContentValueRequestParams {
	s.WorkspaceId = &v
	return s
}

func (s *ListMgsApiResponseBodyResultContentValueRequestParams) Validate() error {
	return dara.Validate(s)
}
