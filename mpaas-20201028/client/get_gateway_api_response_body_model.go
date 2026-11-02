// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetGatewayApiResponseBody
	GetAccessDeniedDetail() *string
	SetApi(v *GetGatewayApiResponseBodyApi) *GetGatewayApiResponseBody
	GetApi() *GetGatewayApiResponseBodyApi
	SetErrorField(v string) *GetGatewayApiResponseBody
	GetErrorField() *string
	SetHttpStatusCode(v int32) *GetGatewayApiResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetGatewayApiResponseBody
	GetRequestId() *string
	SetResultCode(v string) *GetGatewayApiResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *GetGatewayApiResponseBody
	GetResultMessage() *string
	SetSuccess(v bool) *GetGatewayApiResponseBody
	GetSuccess() *bool
}

type GetGatewayApiResponseBody struct {
	AccessDeniedDetail *string                       `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Api                *GetGatewayApiResponseBodyApi `json:"Api,omitempty" xml:"Api,omitempty" type:"Struct"`
	ErrorField         *string                       `json:"ErrorField,omitempty" xml:"ErrorField,omitempty"`
	HttpStatusCode     *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId          *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode         *string                       `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string                       `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	Success            *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayApiResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayApiResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetGatewayApiResponseBody) GetApi() *GetGatewayApiResponseBodyApi {
	return s.Api
}

func (s *GetGatewayApiResponseBody) GetErrorField() *string {
	return s.ErrorField
}

func (s *GetGatewayApiResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetGatewayApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayApiResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetGatewayApiResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *GetGatewayApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGatewayApiResponseBody) SetAccessDeniedDetail(v string) *GetGatewayApiResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetGatewayApiResponseBody) SetApi(v *GetGatewayApiResponseBodyApi) *GetGatewayApiResponseBody {
	s.Api = v
	return s
}

func (s *GetGatewayApiResponseBody) SetErrorField(v string) *GetGatewayApiResponseBody {
	s.ErrorField = &v
	return s
}

func (s *GetGatewayApiResponseBody) SetHttpStatusCode(v int32) *GetGatewayApiResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGatewayApiResponseBody) SetRequestId(v string) *GetGatewayApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayApiResponseBody) SetResultCode(v string) *GetGatewayApiResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetGatewayApiResponseBody) SetResultMessage(v string) *GetGatewayApiResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *GetGatewayApiResponseBody) SetSuccess(v bool) *GetGatewayApiResponseBody {
	s.Success = &v
	return s
}

func (s *GetGatewayApiResponseBody) Validate() error {
	if s.Api != nil {
		if err := s.Api.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayApiResponseBodyApi struct {
	ApiId             *int64                                        `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName           *string                                       `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	ApiStatus         *string                                       `json:"ApiStatus,omitempty" xml:"ApiStatus,omitempty"`
	ApiType           *string                                       `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	AuthRuleName      *string                                       `json:"AuthRuleName,omitempty" xml:"AuthRuleName,omitempty"`
	CacheRule         *GetGatewayApiResponseBodyApiCacheRule        `json:"CacheRule,omitempty" xml:"CacheRule,omitempty" type:"Struct"`
	Charset           *string                                       `json:"Charset,omitempty" xml:"Charset,omitempty"`
	ContentType       *string                                       `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	DefaultLimitRule  *GetGatewayApiResponseBodyApiDefaultLimitRule `json:"DefaultLimitRule,omitempty" xml:"DefaultLimitRule,omitempty" type:"Struct"`
	Description       *string                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId           *int64                                        `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HeaderRules       []*GetGatewayApiResponseBodyApiHeaderRules    `json:"HeaderRules,omitempty" xml:"HeaderRules,omitempty" type:"Repeated"`
	InterfaceType     *string                                       `json:"InterfaceType,omitempty" xml:"InterfaceType,omitempty"`
	LimitRule         *GetGatewayApiResponseBodyApiLimitRule        `json:"LimitRule,omitempty" xml:"LimitRule,omitempty" type:"Struct"`
	Method            *string                                       `json:"Method,omitempty" xml:"Method,omitempty"`
	MethodName        *string                                       `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	MigrateRule       *GetGatewayApiResponseBodyApiMigrateRule      `json:"MigrateRule,omitempty" xml:"MigrateRule,omitempty" type:"Struct"`
	MockRule          *GetGatewayApiResponseBodyApiMockRule         `json:"MockRule,omitempty" xml:"MockRule,omitempty" type:"Struct"`
	NeedEncrypt       *bool                                         `json:"NeedEncrypt,omitempty" xml:"NeedEncrypt,omitempty"`
	NeedJsonp         *bool                                         `json:"NeedJsonp,omitempty" xml:"NeedJsonp,omitempty"`
	NeedSign          *bool                                         `json:"NeedSign,omitempty" xml:"NeedSign,omitempty"`
	OperationType     *string                                       `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	Path              *string                                       `json:"Path,omitempty" xml:"Path,omitempty"`
	RequestBodyModel  *string                                       `json:"RequestBodyModel,omitempty" xml:"RequestBodyModel,omitempty"`
	RequestParams     []*GetGatewayApiResponseBodyApiRequestParams  `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Repeated"`
	ResponseBodyModel *string                                       `json:"ResponseBodyModel,omitempty" xml:"ResponseBodyModel,omitempty"`
	RpcParameterTypes []*string                                     `json:"RpcParameterTypes,omitempty" xml:"RpcParameterTypes,omitempty" type:"Repeated"`
	Serialization     *string                                       `json:"Serialization,omitempty" xml:"Serialization,omitempty"`
	SysName           *string                                       `json:"SysName,omitempty" xml:"SysName,omitempty"`
	Timeout           *int32                                        `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s GetGatewayApiResponseBodyApi) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayApiResponseBodyApi) GoString() string {
	return s.String()
}

func (s *GetGatewayApiResponseBodyApi) GetApiId() *int64 {
	return s.ApiId
}

func (s *GetGatewayApiResponseBodyApi) GetApiName() *string {
	return s.ApiName
}

func (s *GetGatewayApiResponseBodyApi) GetApiStatus() *string {
	return s.ApiStatus
}

func (s *GetGatewayApiResponseBodyApi) GetApiType() *string {
	return s.ApiType
}

func (s *GetGatewayApiResponseBodyApi) GetAuthRuleName() *string {
	return s.AuthRuleName
}

func (s *GetGatewayApiResponseBodyApi) GetCacheRule() *GetGatewayApiResponseBodyApiCacheRule {
	return s.CacheRule
}

func (s *GetGatewayApiResponseBodyApi) GetCharset() *string {
	return s.Charset
}

func (s *GetGatewayApiResponseBodyApi) GetContentType() *string {
	return s.ContentType
}

func (s *GetGatewayApiResponseBodyApi) GetDefaultLimitRule() *GetGatewayApiResponseBodyApiDefaultLimitRule {
	return s.DefaultLimitRule
}

func (s *GetGatewayApiResponseBodyApi) GetDescription() *string {
	return s.Description
}

func (s *GetGatewayApiResponseBodyApi) GetGroupId() *int64 {
	return s.GroupId
}

func (s *GetGatewayApiResponseBodyApi) GetHeaderRules() []*GetGatewayApiResponseBodyApiHeaderRules {
	return s.HeaderRules
}

func (s *GetGatewayApiResponseBodyApi) GetInterfaceType() *string {
	return s.InterfaceType
}

func (s *GetGatewayApiResponseBodyApi) GetLimitRule() *GetGatewayApiResponseBodyApiLimitRule {
	return s.LimitRule
}

func (s *GetGatewayApiResponseBodyApi) GetMethod() *string {
	return s.Method
}

func (s *GetGatewayApiResponseBodyApi) GetMethodName() *string {
	return s.MethodName
}

func (s *GetGatewayApiResponseBodyApi) GetMigrateRule() *GetGatewayApiResponseBodyApiMigrateRule {
	return s.MigrateRule
}

func (s *GetGatewayApiResponseBodyApi) GetMockRule() *GetGatewayApiResponseBodyApiMockRule {
	return s.MockRule
}

func (s *GetGatewayApiResponseBodyApi) GetNeedEncrypt() *bool {
	return s.NeedEncrypt
}

func (s *GetGatewayApiResponseBodyApi) GetNeedJsonp() *bool {
	return s.NeedJsonp
}

func (s *GetGatewayApiResponseBodyApi) GetNeedSign() *bool {
	return s.NeedSign
}

func (s *GetGatewayApiResponseBodyApi) GetOperationType() *string {
	return s.OperationType
}

func (s *GetGatewayApiResponseBodyApi) GetPath() *string {
	return s.Path
}

func (s *GetGatewayApiResponseBodyApi) GetRequestBodyModel() *string {
	return s.RequestBodyModel
}

func (s *GetGatewayApiResponseBodyApi) GetRequestParams() []*GetGatewayApiResponseBodyApiRequestParams {
	return s.RequestParams
}

func (s *GetGatewayApiResponseBodyApi) GetResponseBodyModel() *string {
	return s.ResponseBodyModel
}

func (s *GetGatewayApiResponseBodyApi) GetRpcParameterTypes() []*string {
	return s.RpcParameterTypes
}

func (s *GetGatewayApiResponseBodyApi) GetSerialization() *string {
	return s.Serialization
}

func (s *GetGatewayApiResponseBodyApi) GetSysName() *string {
	return s.SysName
}

func (s *GetGatewayApiResponseBodyApi) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetGatewayApiResponseBodyApi) SetApiId(v int64) *GetGatewayApiResponseBodyApi {
	s.ApiId = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetApiName(v string) *GetGatewayApiResponseBodyApi {
	s.ApiName = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetApiStatus(v string) *GetGatewayApiResponseBodyApi {
	s.ApiStatus = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetApiType(v string) *GetGatewayApiResponseBodyApi {
	s.ApiType = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetAuthRuleName(v string) *GetGatewayApiResponseBodyApi {
	s.AuthRuleName = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetCacheRule(v *GetGatewayApiResponseBodyApiCacheRule) *GetGatewayApiResponseBodyApi {
	s.CacheRule = v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetCharset(v string) *GetGatewayApiResponseBodyApi {
	s.Charset = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetContentType(v string) *GetGatewayApiResponseBodyApi {
	s.ContentType = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetDefaultLimitRule(v *GetGatewayApiResponseBodyApiDefaultLimitRule) *GetGatewayApiResponseBodyApi {
	s.DefaultLimitRule = v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetDescription(v string) *GetGatewayApiResponseBodyApi {
	s.Description = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetGroupId(v int64) *GetGatewayApiResponseBodyApi {
	s.GroupId = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetHeaderRules(v []*GetGatewayApiResponseBodyApiHeaderRules) *GetGatewayApiResponseBodyApi {
	s.HeaderRules = v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetInterfaceType(v string) *GetGatewayApiResponseBodyApi {
	s.InterfaceType = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetLimitRule(v *GetGatewayApiResponseBodyApiLimitRule) *GetGatewayApiResponseBodyApi {
	s.LimitRule = v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetMethod(v string) *GetGatewayApiResponseBodyApi {
	s.Method = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetMethodName(v string) *GetGatewayApiResponseBodyApi {
	s.MethodName = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetMigrateRule(v *GetGatewayApiResponseBodyApiMigrateRule) *GetGatewayApiResponseBodyApi {
	s.MigrateRule = v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetMockRule(v *GetGatewayApiResponseBodyApiMockRule) *GetGatewayApiResponseBodyApi {
	s.MockRule = v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetNeedEncrypt(v bool) *GetGatewayApiResponseBodyApi {
	s.NeedEncrypt = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetNeedJsonp(v bool) *GetGatewayApiResponseBodyApi {
	s.NeedJsonp = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetNeedSign(v bool) *GetGatewayApiResponseBodyApi {
	s.NeedSign = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetOperationType(v string) *GetGatewayApiResponseBodyApi {
	s.OperationType = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetPath(v string) *GetGatewayApiResponseBodyApi {
	s.Path = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetRequestBodyModel(v string) *GetGatewayApiResponseBodyApi {
	s.RequestBodyModel = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetRequestParams(v []*GetGatewayApiResponseBodyApiRequestParams) *GetGatewayApiResponseBodyApi {
	s.RequestParams = v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetResponseBodyModel(v string) *GetGatewayApiResponseBodyApi {
	s.ResponseBodyModel = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetRpcParameterTypes(v []*string) *GetGatewayApiResponseBodyApi {
	s.RpcParameterTypes = v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetSerialization(v string) *GetGatewayApiResponseBodyApi {
	s.Serialization = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetSysName(v string) *GetGatewayApiResponseBodyApi {
	s.SysName = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) SetTimeout(v int32) *GetGatewayApiResponseBodyApi {
	s.Timeout = &v
	return s
}

func (s *GetGatewayApiResponseBodyApi) Validate() error {
	if s.CacheRule != nil {
		if err := s.CacheRule.Validate(); err != nil {
			return err
		}
	}
	if s.DefaultLimitRule != nil {
		if err := s.DefaultLimitRule.Validate(); err != nil {
			return err
		}
	}
	if s.HeaderRules != nil {
		for _, item := range s.HeaderRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LimitRule != nil {
		if err := s.LimitRule.Validate(); err != nil {
			return err
		}
	}
	if s.MigrateRule != nil {
		if err := s.MigrateRule.Validate(); err != nil {
			return err
		}
	}
	if s.MockRule != nil {
		if err := s.MockRule.Validate(); err != nil {
			return err
		}
	}
	if s.RequestParams != nil {
		for _, item := range s.RequestParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGatewayApiResponseBodyApiCacheRule struct {
	CacheKey  *string `json:"CacheKey,omitempty" xml:"CacheKey,omitempty"`
	NeedCache *bool   `json:"NeedCache,omitempty" xml:"NeedCache,omitempty"`
	Ttl       *int64  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s GetGatewayApiResponseBodyApiCacheRule) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayApiResponseBodyApiCacheRule) GoString() string {
	return s.String()
}

func (s *GetGatewayApiResponseBodyApiCacheRule) GetCacheKey() *string {
	return s.CacheKey
}

func (s *GetGatewayApiResponseBodyApiCacheRule) GetNeedCache() *bool {
	return s.NeedCache
}

func (s *GetGatewayApiResponseBodyApiCacheRule) GetTtl() *int64 {
	return s.Ttl
}

func (s *GetGatewayApiResponseBodyApiCacheRule) SetCacheKey(v string) *GetGatewayApiResponseBodyApiCacheRule {
	s.CacheKey = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiCacheRule) SetNeedCache(v bool) *GetGatewayApiResponseBodyApiCacheRule {
	s.NeedCache = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiCacheRule) SetTtl(v int64) *GetGatewayApiResponseBodyApiCacheRule {
	s.Ttl = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiCacheRule) Validate() error {
	return dara.Validate(s)
}

type GetGatewayApiResponseBodyApiDefaultLimitRule struct {
	DefaultLimit *bool `json:"DefaultLimit,omitempty" xml:"DefaultLimit,omitempty"`
}

func (s GetGatewayApiResponseBodyApiDefaultLimitRule) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayApiResponseBodyApiDefaultLimitRule) GoString() string {
	return s.String()
}

func (s *GetGatewayApiResponseBodyApiDefaultLimitRule) GetDefaultLimit() *bool {
	return s.DefaultLimit
}

func (s *GetGatewayApiResponseBodyApiDefaultLimitRule) SetDefaultLimit(v bool) *GetGatewayApiResponseBodyApiDefaultLimitRule {
	s.DefaultLimit = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiDefaultLimitRule) Validate() error {
	return dara.Validate(s)
}

type GetGatewayApiResponseBodyApiHeaderRules struct {
	HeaderKey *string `json:"HeaderKey,omitempty" xml:"HeaderKey,omitempty"`
	Location  *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetGatewayApiResponseBodyApiHeaderRules) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayApiResponseBodyApiHeaderRules) GoString() string {
	return s.String()
}

func (s *GetGatewayApiResponseBodyApiHeaderRules) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *GetGatewayApiResponseBodyApiHeaderRules) GetLocation() *string {
	return s.Location
}

func (s *GetGatewayApiResponseBodyApiHeaderRules) GetType() *string {
	return s.Type
}

func (s *GetGatewayApiResponseBodyApiHeaderRules) GetValue() *string {
	return s.Value
}

func (s *GetGatewayApiResponseBodyApiHeaderRules) SetHeaderKey(v string) *GetGatewayApiResponseBodyApiHeaderRules {
	s.HeaderKey = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiHeaderRules) SetLocation(v string) *GetGatewayApiResponseBodyApiHeaderRules {
	s.Location = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiHeaderRules) SetType(v string) *GetGatewayApiResponseBodyApiHeaderRules {
	s.Type = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiHeaderRules) SetValue(v string) *GetGatewayApiResponseBodyApiHeaderRules {
	s.Value = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiHeaderRules) Validate() error {
	return dara.Validate(s)
}

type GetGatewayApiResponseBodyApiLimitRule struct {
	DefaultResponse *string  `json:"DefaultResponse,omitempty" xml:"DefaultResponse,omitempty"`
	Limit           *float64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Mode            *string  `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s GetGatewayApiResponseBodyApiLimitRule) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayApiResponseBodyApiLimitRule) GoString() string {
	return s.String()
}

func (s *GetGatewayApiResponseBodyApiLimitRule) GetDefaultResponse() *string {
	return s.DefaultResponse
}

func (s *GetGatewayApiResponseBodyApiLimitRule) GetLimit() *float64 {
	return s.Limit
}

func (s *GetGatewayApiResponseBodyApiLimitRule) GetMode() *string {
	return s.Mode
}

func (s *GetGatewayApiResponseBodyApiLimitRule) SetDefaultResponse(v string) *GetGatewayApiResponseBodyApiLimitRule {
	s.DefaultResponse = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiLimitRule) SetLimit(v float64) *GetGatewayApiResponseBodyApiLimitRule {
	s.Limit = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiLimitRule) SetMode(v string) *GetGatewayApiResponseBodyApiLimitRule {
	s.Mode = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiLimitRule) Validate() error {
	return dara.Validate(s)
}

type GetGatewayApiResponseBodyApiMigrateRule struct {
	FlowPercent          *int32  `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
	GroupId              *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	NeedMigrate          *bool   `json:"NeedMigrate,omitempty" xml:"NeedMigrate,omitempty"`
	NeedSwitchCompletely *bool   `json:"NeedSwitchCompletely,omitempty" xml:"NeedSwitchCompletely,omitempty"`
	UpstreamType         *string `json:"UpstreamType,omitempty" xml:"UpstreamType,omitempty"`
}

func (s GetGatewayApiResponseBodyApiMigrateRule) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayApiResponseBodyApiMigrateRule) GoString() string {
	return s.String()
}

func (s *GetGatewayApiResponseBodyApiMigrateRule) GetFlowPercent() *int32 {
	return s.FlowPercent
}

func (s *GetGatewayApiResponseBodyApiMigrateRule) GetGroupId() *int64 {
	return s.GroupId
}

func (s *GetGatewayApiResponseBodyApiMigrateRule) GetGroupName() *string {
	return s.GroupName
}

func (s *GetGatewayApiResponseBodyApiMigrateRule) GetNeedMigrate() *bool {
	return s.NeedMigrate
}

func (s *GetGatewayApiResponseBodyApiMigrateRule) GetNeedSwitchCompletely() *bool {
	return s.NeedSwitchCompletely
}

func (s *GetGatewayApiResponseBodyApiMigrateRule) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *GetGatewayApiResponseBodyApiMigrateRule) SetFlowPercent(v int32) *GetGatewayApiResponseBodyApiMigrateRule {
	s.FlowPercent = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiMigrateRule) SetGroupId(v int64) *GetGatewayApiResponseBodyApiMigrateRule {
	s.GroupId = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiMigrateRule) SetGroupName(v string) *GetGatewayApiResponseBodyApiMigrateRule {
	s.GroupName = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiMigrateRule) SetNeedMigrate(v bool) *GetGatewayApiResponseBodyApiMigrateRule {
	s.NeedMigrate = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiMigrateRule) SetNeedSwitchCompletely(v bool) *GetGatewayApiResponseBodyApiMigrateRule {
	s.NeedSwitchCompletely = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiMigrateRule) SetUpstreamType(v string) *GetGatewayApiResponseBodyApiMigrateRule {
	s.UpstreamType = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiMigrateRule) Validate() error {
	return dara.Validate(s)
}

type GetGatewayApiResponseBodyApiMockRule struct {
	MockData   *string `json:"MockData,omitempty" xml:"MockData,omitempty"`
	NeedMock   *bool   `json:"NeedMock,omitempty" xml:"NeedMock,omitempty"`
	Percentage *int32  `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetGatewayApiResponseBodyApiMockRule) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayApiResponseBodyApiMockRule) GoString() string {
	return s.String()
}

func (s *GetGatewayApiResponseBodyApiMockRule) GetMockData() *string {
	return s.MockData
}

func (s *GetGatewayApiResponseBodyApiMockRule) GetNeedMock() *bool {
	return s.NeedMock
}

func (s *GetGatewayApiResponseBodyApiMockRule) GetPercentage() *int32 {
	return s.Percentage
}

func (s *GetGatewayApiResponseBodyApiMockRule) GetType() *string {
	return s.Type
}

func (s *GetGatewayApiResponseBodyApiMockRule) SetMockData(v string) *GetGatewayApiResponseBodyApiMockRule {
	s.MockData = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiMockRule) SetNeedMock(v bool) *GetGatewayApiResponseBodyApiMockRule {
	s.NeedMock = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiMockRule) SetPercentage(v int32) *GetGatewayApiResponseBodyApiMockRule {
	s.Percentage = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiMockRule) SetType(v string) *GetGatewayApiResponseBodyApiMockRule {
	s.Type = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiMockRule) Validate() error {
	return dara.Validate(s)
}

type GetGatewayApiResponseBodyApiRequestParams struct {
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RefType      *string `json:"RefType,omitempty" xml:"RefType,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetGatewayApiResponseBodyApiRequestParams) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayApiResponseBodyApiRequestParams) GoString() string {
	return s.String()
}

func (s *GetGatewayApiResponseBodyApiRequestParams) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetGatewayApiResponseBodyApiRequestParams) GetDescription() *string {
	return s.Description
}

func (s *GetGatewayApiResponseBodyApiRequestParams) GetLocation() *string {
	return s.Location
}

func (s *GetGatewayApiResponseBodyApiRequestParams) GetName() *string {
	return s.Name
}

func (s *GetGatewayApiResponseBodyApiRequestParams) GetRefType() *string {
	return s.RefType
}

func (s *GetGatewayApiResponseBodyApiRequestParams) GetType() *string {
	return s.Type
}

func (s *GetGatewayApiResponseBodyApiRequestParams) SetDefaultValue(v string) *GetGatewayApiResponseBodyApiRequestParams {
	s.DefaultValue = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiRequestParams) SetDescription(v string) *GetGatewayApiResponseBodyApiRequestParams {
	s.Description = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiRequestParams) SetLocation(v string) *GetGatewayApiResponseBodyApiRequestParams {
	s.Location = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiRequestParams) SetName(v string) *GetGatewayApiResponseBodyApiRequestParams {
	s.Name = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiRequestParams) SetRefType(v string) *GetGatewayApiResponseBodyApiRequestParams {
	s.RefType = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiRequestParams) SetType(v string) *GetGatewayApiResponseBodyApiRequestParams {
	s.Type = &v
	return s
}

func (s *GetGatewayApiResponseBodyApiRequestParams) Validate() error {
	return dara.Validate(s)
}
