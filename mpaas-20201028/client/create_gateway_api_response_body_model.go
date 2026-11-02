// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateGatewayApiResponseBody
	GetAccessDeniedDetail() *string
	SetApi(v *CreateGatewayApiResponseBodyApi) *CreateGatewayApiResponseBody
	GetApi() *CreateGatewayApiResponseBodyApi
	SetErrorField(v string) *CreateGatewayApiResponseBody
	GetErrorField() *string
	SetHttpStatusCode(v int32) *CreateGatewayApiResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateGatewayApiResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateGatewayApiResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateGatewayApiResponseBody
	GetResultMessage() *string
	SetSuccess(v bool) *CreateGatewayApiResponseBody
	GetSuccess() *bool
}

type CreateGatewayApiResponseBody struct {
	AccessDeniedDetail *string                          `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Api                *CreateGatewayApiResponseBodyApi `json:"Api,omitempty" xml:"Api,omitempty" type:"Struct"`
	ErrorField         *string                          `json:"ErrorField,omitempty" xml:"ErrorField,omitempty"`
	HttpStatusCode     *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId          *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode         *string                          `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string                          `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	Success            *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGatewayApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateGatewayApiResponseBody) GetApi() *CreateGatewayApiResponseBodyApi {
	return s.Api
}

func (s *CreateGatewayApiResponseBody) GetErrorField() *string {
	return s.ErrorField
}

func (s *CreateGatewayApiResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateGatewayApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGatewayApiResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateGatewayApiResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateGatewayApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateGatewayApiResponseBody) SetAccessDeniedDetail(v string) *CreateGatewayApiResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateGatewayApiResponseBody) SetApi(v *CreateGatewayApiResponseBodyApi) *CreateGatewayApiResponseBody {
	s.Api = v
	return s
}

func (s *CreateGatewayApiResponseBody) SetErrorField(v string) *CreateGatewayApiResponseBody {
	s.ErrorField = &v
	return s
}

func (s *CreateGatewayApiResponseBody) SetHttpStatusCode(v int32) *CreateGatewayApiResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateGatewayApiResponseBody) SetRequestId(v string) *CreateGatewayApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayApiResponseBody) SetResultCode(v string) *CreateGatewayApiResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateGatewayApiResponseBody) SetResultMessage(v string) *CreateGatewayApiResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateGatewayApiResponseBody) SetSuccess(v bool) *CreateGatewayApiResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGatewayApiResponseBody) Validate() error {
	if s.Api != nil {
		if err := s.Api.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGatewayApiResponseBodyApi struct {
	ApiId             *int64                                           `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName           *string                                          `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	ApiStatus         *string                                          `json:"ApiStatus,omitempty" xml:"ApiStatus,omitempty"`
	ApiType           *string                                          `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	AuthRuleName      *string                                          `json:"AuthRuleName,omitempty" xml:"AuthRuleName,omitempty"`
	CacheRule         *CreateGatewayApiResponseBodyApiCacheRule        `json:"CacheRule,omitempty" xml:"CacheRule,omitempty" type:"Struct"`
	Charset           *string                                          `json:"Charset,omitempty" xml:"Charset,omitempty"`
	ContentType       *string                                          `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	DefaultLimitRule  *CreateGatewayApiResponseBodyApiDefaultLimitRule `json:"DefaultLimitRule,omitempty" xml:"DefaultLimitRule,omitempty" type:"Struct"`
	Description       *string                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId           *int64                                           `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HeaderRules       []*CreateGatewayApiResponseBodyApiHeaderRules    `json:"HeaderRules,omitempty" xml:"HeaderRules,omitempty" type:"Repeated"`
	InterfaceType     *string                                          `json:"InterfaceType,omitempty" xml:"InterfaceType,omitempty"`
	LimitRule         *CreateGatewayApiResponseBodyApiLimitRule        `json:"LimitRule,omitempty" xml:"LimitRule,omitempty" type:"Struct"`
	Method            *string                                          `json:"Method,omitempty" xml:"Method,omitempty"`
	MethodName        *string                                          `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	MigrateRule       *CreateGatewayApiResponseBodyApiMigrateRule      `json:"MigrateRule,omitempty" xml:"MigrateRule,omitempty" type:"Struct"`
	MockRule          *CreateGatewayApiResponseBodyApiMockRule         `json:"MockRule,omitempty" xml:"MockRule,omitempty" type:"Struct"`
	NeedEncrypt       *bool                                            `json:"NeedEncrypt,omitempty" xml:"NeedEncrypt,omitempty"`
	NeedJsonp         *bool                                            `json:"NeedJsonp,omitempty" xml:"NeedJsonp,omitempty"`
	NeedSign          *bool                                            `json:"NeedSign,omitempty" xml:"NeedSign,omitempty"`
	OperationType     *string                                          `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	Path              *string                                          `json:"Path,omitempty" xml:"Path,omitempty"`
	RequestBodyModel  *string                                          `json:"RequestBodyModel,omitempty" xml:"RequestBodyModel,omitempty"`
	RequestParams     []*CreateGatewayApiResponseBodyApiRequestParams  `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Repeated"`
	ResponseBodyModel *string                                          `json:"ResponseBodyModel,omitempty" xml:"ResponseBodyModel,omitempty"`
	RpcParameterTypes []*string                                        `json:"RpcParameterTypes,omitempty" xml:"RpcParameterTypes,omitempty" type:"Repeated"`
	Serialization     *string                                          `json:"Serialization,omitempty" xml:"Serialization,omitempty"`
	SysName           *string                                          `json:"SysName,omitempty" xml:"SysName,omitempty"`
	Timeout           *int32                                           `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateGatewayApiResponseBodyApi) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiResponseBodyApi) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiResponseBodyApi) GetApiId() *int64 {
	return s.ApiId
}

func (s *CreateGatewayApiResponseBodyApi) GetApiName() *string {
	return s.ApiName
}

func (s *CreateGatewayApiResponseBodyApi) GetApiStatus() *string {
	return s.ApiStatus
}

func (s *CreateGatewayApiResponseBodyApi) GetApiType() *string {
	return s.ApiType
}

func (s *CreateGatewayApiResponseBodyApi) GetAuthRuleName() *string {
	return s.AuthRuleName
}

func (s *CreateGatewayApiResponseBodyApi) GetCacheRule() *CreateGatewayApiResponseBodyApiCacheRule {
	return s.CacheRule
}

func (s *CreateGatewayApiResponseBodyApi) GetCharset() *string {
	return s.Charset
}

func (s *CreateGatewayApiResponseBodyApi) GetContentType() *string {
	return s.ContentType
}

func (s *CreateGatewayApiResponseBodyApi) GetDefaultLimitRule() *CreateGatewayApiResponseBodyApiDefaultLimitRule {
	return s.DefaultLimitRule
}

func (s *CreateGatewayApiResponseBodyApi) GetDescription() *string {
	return s.Description
}

func (s *CreateGatewayApiResponseBodyApi) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateGatewayApiResponseBodyApi) GetHeaderRules() []*CreateGatewayApiResponseBodyApiHeaderRules {
	return s.HeaderRules
}

func (s *CreateGatewayApiResponseBodyApi) GetInterfaceType() *string {
	return s.InterfaceType
}

func (s *CreateGatewayApiResponseBodyApi) GetLimitRule() *CreateGatewayApiResponseBodyApiLimitRule {
	return s.LimitRule
}

func (s *CreateGatewayApiResponseBodyApi) GetMethod() *string {
	return s.Method
}

func (s *CreateGatewayApiResponseBodyApi) GetMethodName() *string {
	return s.MethodName
}

func (s *CreateGatewayApiResponseBodyApi) GetMigrateRule() *CreateGatewayApiResponseBodyApiMigrateRule {
	return s.MigrateRule
}

func (s *CreateGatewayApiResponseBodyApi) GetMockRule() *CreateGatewayApiResponseBodyApiMockRule {
	return s.MockRule
}

func (s *CreateGatewayApiResponseBodyApi) GetNeedEncrypt() *bool {
	return s.NeedEncrypt
}

func (s *CreateGatewayApiResponseBodyApi) GetNeedJsonp() *bool {
	return s.NeedJsonp
}

func (s *CreateGatewayApiResponseBodyApi) GetNeedSign() *bool {
	return s.NeedSign
}

func (s *CreateGatewayApiResponseBodyApi) GetOperationType() *string {
	return s.OperationType
}

func (s *CreateGatewayApiResponseBodyApi) GetPath() *string {
	return s.Path
}

func (s *CreateGatewayApiResponseBodyApi) GetRequestBodyModel() *string {
	return s.RequestBodyModel
}

func (s *CreateGatewayApiResponseBodyApi) GetRequestParams() []*CreateGatewayApiResponseBodyApiRequestParams {
	return s.RequestParams
}

func (s *CreateGatewayApiResponseBodyApi) GetResponseBodyModel() *string {
	return s.ResponseBodyModel
}

func (s *CreateGatewayApiResponseBodyApi) GetRpcParameterTypes() []*string {
	return s.RpcParameterTypes
}

func (s *CreateGatewayApiResponseBodyApi) GetSerialization() *string {
	return s.Serialization
}

func (s *CreateGatewayApiResponseBodyApi) GetSysName() *string {
	return s.SysName
}

func (s *CreateGatewayApiResponseBodyApi) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateGatewayApiResponseBodyApi) SetApiId(v int64) *CreateGatewayApiResponseBodyApi {
	s.ApiId = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetApiName(v string) *CreateGatewayApiResponseBodyApi {
	s.ApiName = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetApiStatus(v string) *CreateGatewayApiResponseBodyApi {
	s.ApiStatus = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetApiType(v string) *CreateGatewayApiResponseBodyApi {
	s.ApiType = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetAuthRuleName(v string) *CreateGatewayApiResponseBodyApi {
	s.AuthRuleName = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetCacheRule(v *CreateGatewayApiResponseBodyApiCacheRule) *CreateGatewayApiResponseBodyApi {
	s.CacheRule = v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetCharset(v string) *CreateGatewayApiResponseBodyApi {
	s.Charset = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetContentType(v string) *CreateGatewayApiResponseBodyApi {
	s.ContentType = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetDefaultLimitRule(v *CreateGatewayApiResponseBodyApiDefaultLimitRule) *CreateGatewayApiResponseBodyApi {
	s.DefaultLimitRule = v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetDescription(v string) *CreateGatewayApiResponseBodyApi {
	s.Description = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetGroupId(v int64) *CreateGatewayApiResponseBodyApi {
	s.GroupId = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetHeaderRules(v []*CreateGatewayApiResponseBodyApiHeaderRules) *CreateGatewayApiResponseBodyApi {
	s.HeaderRules = v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetInterfaceType(v string) *CreateGatewayApiResponseBodyApi {
	s.InterfaceType = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetLimitRule(v *CreateGatewayApiResponseBodyApiLimitRule) *CreateGatewayApiResponseBodyApi {
	s.LimitRule = v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetMethod(v string) *CreateGatewayApiResponseBodyApi {
	s.Method = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetMethodName(v string) *CreateGatewayApiResponseBodyApi {
	s.MethodName = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetMigrateRule(v *CreateGatewayApiResponseBodyApiMigrateRule) *CreateGatewayApiResponseBodyApi {
	s.MigrateRule = v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetMockRule(v *CreateGatewayApiResponseBodyApiMockRule) *CreateGatewayApiResponseBodyApi {
	s.MockRule = v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetNeedEncrypt(v bool) *CreateGatewayApiResponseBodyApi {
	s.NeedEncrypt = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetNeedJsonp(v bool) *CreateGatewayApiResponseBodyApi {
	s.NeedJsonp = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetNeedSign(v bool) *CreateGatewayApiResponseBodyApi {
	s.NeedSign = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetOperationType(v string) *CreateGatewayApiResponseBodyApi {
	s.OperationType = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetPath(v string) *CreateGatewayApiResponseBodyApi {
	s.Path = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetRequestBodyModel(v string) *CreateGatewayApiResponseBodyApi {
	s.RequestBodyModel = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetRequestParams(v []*CreateGatewayApiResponseBodyApiRequestParams) *CreateGatewayApiResponseBodyApi {
	s.RequestParams = v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetResponseBodyModel(v string) *CreateGatewayApiResponseBodyApi {
	s.ResponseBodyModel = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetRpcParameterTypes(v []*string) *CreateGatewayApiResponseBodyApi {
	s.RpcParameterTypes = v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetSerialization(v string) *CreateGatewayApiResponseBodyApi {
	s.Serialization = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetSysName(v string) *CreateGatewayApiResponseBodyApi {
	s.SysName = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) SetTimeout(v int32) *CreateGatewayApiResponseBodyApi {
	s.Timeout = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApi) Validate() error {
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

type CreateGatewayApiResponseBodyApiCacheRule struct {
	CacheKey  *string `json:"CacheKey,omitempty" xml:"CacheKey,omitempty"`
	NeedCache *bool   `json:"NeedCache,omitempty" xml:"NeedCache,omitempty"`
	Ttl       *int64  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s CreateGatewayApiResponseBodyApiCacheRule) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiResponseBodyApiCacheRule) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiResponseBodyApiCacheRule) GetCacheKey() *string {
	return s.CacheKey
}

func (s *CreateGatewayApiResponseBodyApiCacheRule) GetNeedCache() *bool {
	return s.NeedCache
}

func (s *CreateGatewayApiResponseBodyApiCacheRule) GetTtl() *int64 {
	return s.Ttl
}

func (s *CreateGatewayApiResponseBodyApiCacheRule) SetCacheKey(v string) *CreateGatewayApiResponseBodyApiCacheRule {
	s.CacheKey = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiCacheRule) SetNeedCache(v bool) *CreateGatewayApiResponseBodyApiCacheRule {
	s.NeedCache = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiCacheRule) SetTtl(v int64) *CreateGatewayApiResponseBodyApiCacheRule {
	s.Ttl = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiCacheRule) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayApiResponseBodyApiDefaultLimitRule struct {
	DefaultLimit *bool `json:"DefaultLimit,omitempty" xml:"DefaultLimit,omitempty"`
}

func (s CreateGatewayApiResponseBodyApiDefaultLimitRule) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiResponseBodyApiDefaultLimitRule) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiResponseBodyApiDefaultLimitRule) GetDefaultLimit() *bool {
	return s.DefaultLimit
}

func (s *CreateGatewayApiResponseBodyApiDefaultLimitRule) SetDefaultLimit(v bool) *CreateGatewayApiResponseBodyApiDefaultLimitRule {
	s.DefaultLimit = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiDefaultLimitRule) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayApiResponseBodyApiHeaderRules struct {
	HeaderKey *string `json:"HeaderKey,omitempty" xml:"HeaderKey,omitempty"`
	Location  *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateGatewayApiResponseBodyApiHeaderRules) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiResponseBodyApiHeaderRules) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiResponseBodyApiHeaderRules) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *CreateGatewayApiResponseBodyApiHeaderRules) GetLocation() *string {
	return s.Location
}

func (s *CreateGatewayApiResponseBodyApiHeaderRules) GetType() *string {
	return s.Type
}

func (s *CreateGatewayApiResponseBodyApiHeaderRules) GetValue() *string {
	return s.Value
}

func (s *CreateGatewayApiResponseBodyApiHeaderRules) SetHeaderKey(v string) *CreateGatewayApiResponseBodyApiHeaderRules {
	s.HeaderKey = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiHeaderRules) SetLocation(v string) *CreateGatewayApiResponseBodyApiHeaderRules {
	s.Location = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiHeaderRules) SetType(v string) *CreateGatewayApiResponseBodyApiHeaderRules {
	s.Type = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiHeaderRules) SetValue(v string) *CreateGatewayApiResponseBodyApiHeaderRules {
	s.Value = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiHeaderRules) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayApiResponseBodyApiLimitRule struct {
	DefaultResponse *string  `json:"DefaultResponse,omitempty" xml:"DefaultResponse,omitempty"`
	Limit           *float64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Mode            *string  `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s CreateGatewayApiResponseBodyApiLimitRule) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiResponseBodyApiLimitRule) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiResponseBodyApiLimitRule) GetDefaultResponse() *string {
	return s.DefaultResponse
}

func (s *CreateGatewayApiResponseBodyApiLimitRule) GetLimit() *float64 {
	return s.Limit
}

func (s *CreateGatewayApiResponseBodyApiLimitRule) GetMode() *string {
	return s.Mode
}

func (s *CreateGatewayApiResponseBodyApiLimitRule) SetDefaultResponse(v string) *CreateGatewayApiResponseBodyApiLimitRule {
	s.DefaultResponse = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiLimitRule) SetLimit(v float64) *CreateGatewayApiResponseBodyApiLimitRule {
	s.Limit = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiLimitRule) SetMode(v string) *CreateGatewayApiResponseBodyApiLimitRule {
	s.Mode = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiLimitRule) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayApiResponseBodyApiMigrateRule struct {
	FlowPercent          *int32  `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
	GroupId              *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	NeedMigrate          *bool   `json:"NeedMigrate,omitempty" xml:"NeedMigrate,omitempty"`
	NeedSwitchCompletely *bool   `json:"NeedSwitchCompletely,omitempty" xml:"NeedSwitchCompletely,omitempty"`
	UpstreamType         *string `json:"UpstreamType,omitempty" xml:"UpstreamType,omitempty"`
}

func (s CreateGatewayApiResponseBodyApiMigrateRule) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiResponseBodyApiMigrateRule) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiResponseBodyApiMigrateRule) GetFlowPercent() *int32 {
	return s.FlowPercent
}

func (s *CreateGatewayApiResponseBodyApiMigrateRule) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateGatewayApiResponseBodyApiMigrateRule) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateGatewayApiResponseBodyApiMigrateRule) GetNeedMigrate() *bool {
	return s.NeedMigrate
}

func (s *CreateGatewayApiResponseBodyApiMigrateRule) GetNeedSwitchCompletely() *bool {
	return s.NeedSwitchCompletely
}

func (s *CreateGatewayApiResponseBodyApiMigrateRule) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *CreateGatewayApiResponseBodyApiMigrateRule) SetFlowPercent(v int32) *CreateGatewayApiResponseBodyApiMigrateRule {
	s.FlowPercent = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiMigrateRule) SetGroupId(v int64) *CreateGatewayApiResponseBodyApiMigrateRule {
	s.GroupId = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiMigrateRule) SetGroupName(v string) *CreateGatewayApiResponseBodyApiMigrateRule {
	s.GroupName = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiMigrateRule) SetNeedMigrate(v bool) *CreateGatewayApiResponseBodyApiMigrateRule {
	s.NeedMigrate = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiMigrateRule) SetNeedSwitchCompletely(v bool) *CreateGatewayApiResponseBodyApiMigrateRule {
	s.NeedSwitchCompletely = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiMigrateRule) SetUpstreamType(v string) *CreateGatewayApiResponseBodyApiMigrateRule {
	s.UpstreamType = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiMigrateRule) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayApiResponseBodyApiMockRule struct {
	MockData   *string `json:"MockData,omitempty" xml:"MockData,omitempty"`
	NeedMock   *bool   `json:"NeedMock,omitempty" xml:"NeedMock,omitempty"`
	Percentage *int32  `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateGatewayApiResponseBodyApiMockRule) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiResponseBodyApiMockRule) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiResponseBodyApiMockRule) GetMockData() *string {
	return s.MockData
}

func (s *CreateGatewayApiResponseBodyApiMockRule) GetNeedMock() *bool {
	return s.NeedMock
}

func (s *CreateGatewayApiResponseBodyApiMockRule) GetPercentage() *int32 {
	return s.Percentage
}

func (s *CreateGatewayApiResponseBodyApiMockRule) GetType() *string {
	return s.Type
}

func (s *CreateGatewayApiResponseBodyApiMockRule) SetMockData(v string) *CreateGatewayApiResponseBodyApiMockRule {
	s.MockData = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiMockRule) SetNeedMock(v bool) *CreateGatewayApiResponseBodyApiMockRule {
	s.NeedMock = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiMockRule) SetPercentage(v int32) *CreateGatewayApiResponseBodyApiMockRule {
	s.Percentage = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiMockRule) SetType(v string) *CreateGatewayApiResponseBodyApiMockRule {
	s.Type = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiMockRule) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayApiResponseBodyApiRequestParams struct {
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RefType      *string `json:"RefType,omitempty" xml:"RefType,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateGatewayApiResponseBodyApiRequestParams) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiResponseBodyApiRequestParams) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiResponseBodyApiRequestParams) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *CreateGatewayApiResponseBodyApiRequestParams) GetDescription() *string {
	return s.Description
}

func (s *CreateGatewayApiResponseBodyApiRequestParams) GetLocation() *string {
	return s.Location
}

func (s *CreateGatewayApiResponseBodyApiRequestParams) GetName() *string {
	return s.Name
}

func (s *CreateGatewayApiResponseBodyApiRequestParams) GetRefType() *string {
	return s.RefType
}

func (s *CreateGatewayApiResponseBodyApiRequestParams) GetType() *string {
	return s.Type
}

func (s *CreateGatewayApiResponseBodyApiRequestParams) SetDefaultValue(v string) *CreateGatewayApiResponseBodyApiRequestParams {
	s.DefaultValue = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiRequestParams) SetDescription(v string) *CreateGatewayApiResponseBodyApiRequestParams {
	s.Description = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiRequestParams) SetLocation(v string) *CreateGatewayApiResponseBodyApiRequestParams {
	s.Location = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiRequestParams) SetName(v string) *CreateGatewayApiResponseBodyApiRequestParams {
	s.Name = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiRequestParams) SetRefType(v string) *CreateGatewayApiResponseBodyApiRequestParams {
	s.RefType = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiRequestParams) SetType(v string) *CreateGatewayApiResponseBodyApiRequestParams {
	s.Type = &v
	return s
}

func (s *CreateGatewayApiResponseBodyApiRequestParams) Validate() error {
	return dara.Validate(s)
}
