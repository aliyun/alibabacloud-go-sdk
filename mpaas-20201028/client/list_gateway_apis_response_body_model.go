// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListGatewayApisResponseBody
	GetAccessDeniedDetail() *string
	SetApis(v []*ListGatewayApisResponseBodyApis) *ListGatewayApisResponseBody
	GetApis() []*ListGatewayApisResponseBodyApis
	SetErrorField(v string) *ListGatewayApisResponseBody
	GetErrorField() *string
	SetHttpStatusCode(v int32) *ListGatewayApisResponseBody
	GetHttpStatusCode() *int32
	SetPageNumber(v int32) *ListGatewayApisResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewayApisResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListGatewayApisResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListGatewayApisResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ListGatewayApisResponseBody
	GetResultMessage() *string
	SetSuccess(v bool) *ListGatewayApisResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListGatewayApisResponseBody
	GetTotal() *int64
}

type ListGatewayApisResponseBody struct {
	AccessDeniedDetail *string                            `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Apis               []*ListGatewayApisResponseBodyApis `json:"Apis,omitempty" xml:"Apis,omitempty" type:"Repeated"`
	ErrorField         *string                            `json:"ErrorField,omitempty" xml:"ErrorField,omitempty"`
	HttpStatusCode     *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	PageNumber         *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode         *string                            `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string                            `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	Success            *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Total              *int64                             `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListGatewayApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayApisResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayApisResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListGatewayApisResponseBody) GetApis() []*ListGatewayApisResponseBodyApis {
	return s.Apis
}

func (s *ListGatewayApisResponseBody) GetErrorField() *string {
	return s.ErrorField
}

func (s *ListGatewayApisResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGatewayApisResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayApisResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayApisResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListGatewayApisResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListGatewayApisResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGatewayApisResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListGatewayApisResponseBody) SetAccessDeniedDetail(v string) *ListGatewayApisResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListGatewayApisResponseBody) SetApis(v []*ListGatewayApisResponseBodyApis) *ListGatewayApisResponseBody {
	s.Apis = v
	return s
}

func (s *ListGatewayApisResponseBody) SetErrorField(v string) *ListGatewayApisResponseBody {
	s.ErrorField = &v
	return s
}

func (s *ListGatewayApisResponseBody) SetHttpStatusCode(v int32) *ListGatewayApisResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGatewayApisResponseBody) SetPageNumber(v int32) *ListGatewayApisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayApisResponseBody) SetPageSize(v int32) *ListGatewayApisResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListGatewayApisResponseBody) SetRequestId(v string) *ListGatewayApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayApisResponseBody) SetResultCode(v string) *ListGatewayApisResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListGatewayApisResponseBody) SetResultMessage(v string) *ListGatewayApisResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListGatewayApisResponseBody) SetSuccess(v bool) *ListGatewayApisResponseBody {
	s.Success = &v
	return s
}

func (s *ListGatewayApisResponseBody) SetTotal(v int64) *ListGatewayApisResponseBody {
	s.Total = &v
	return s
}

func (s *ListGatewayApisResponseBody) Validate() error {
	if s.Apis != nil {
		for _, item := range s.Apis {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGatewayApisResponseBodyApis struct {
	ApiId             *int64                                           `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName           *string                                          `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	ApiStatus         *string                                          `json:"ApiStatus,omitempty" xml:"ApiStatus,omitempty"`
	ApiType           *string                                          `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	AuthRuleName      *string                                          `json:"AuthRuleName,omitempty" xml:"AuthRuleName,omitempty"`
	CacheRule         *ListGatewayApisResponseBodyApisCacheRule        `json:"CacheRule,omitempty" xml:"CacheRule,omitempty" type:"Struct"`
	Charset           *string                                          `json:"Charset,omitempty" xml:"Charset,omitempty"`
	ContentType       *string                                          `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	DefaultLimitRule  *ListGatewayApisResponseBodyApisDefaultLimitRule `json:"DefaultLimitRule,omitempty" xml:"DefaultLimitRule,omitempty" type:"Struct"`
	Description       *string                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId           *int64                                           `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HeaderRules       []*ListGatewayApisResponseBodyApisHeaderRules    `json:"HeaderRules,omitempty" xml:"HeaderRules,omitempty" type:"Repeated"`
	InterfaceType     *string                                          `json:"InterfaceType,omitempty" xml:"InterfaceType,omitempty"`
	LimitRule         *ListGatewayApisResponseBodyApisLimitRule        `json:"LimitRule,omitempty" xml:"LimitRule,omitempty" type:"Struct"`
	Method            *string                                          `json:"Method,omitempty" xml:"Method,omitempty"`
	MethodName        *string                                          `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	MigrateRule       *ListGatewayApisResponseBodyApisMigrateRule      `json:"MigrateRule,omitempty" xml:"MigrateRule,omitempty" type:"Struct"`
	MockRule          *ListGatewayApisResponseBodyApisMockRule         `json:"MockRule,omitempty" xml:"MockRule,omitempty" type:"Struct"`
	NeedEncrypt       *bool                                            `json:"NeedEncrypt,omitempty" xml:"NeedEncrypt,omitempty"`
	NeedJsonp         *bool                                            `json:"NeedJsonp,omitempty" xml:"NeedJsonp,omitempty"`
	NeedSign          *bool                                            `json:"NeedSign,omitempty" xml:"NeedSign,omitempty"`
	OperationType     *string                                          `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	Path              *string                                          `json:"Path,omitempty" xml:"Path,omitempty"`
	RequestBodyModel  *string                                          `json:"RequestBodyModel,omitempty" xml:"RequestBodyModel,omitempty"`
	RequestParams     []*ListGatewayApisResponseBodyApisRequestParams  `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Repeated"`
	ResponseBodyModel *string                                          `json:"ResponseBodyModel,omitempty" xml:"ResponseBodyModel,omitempty"`
	RpcParameterTypes []*string                                        `json:"RpcParameterTypes,omitempty" xml:"RpcParameterTypes,omitempty" type:"Repeated"`
	Serialization     *string                                          `json:"Serialization,omitempty" xml:"Serialization,omitempty"`
	SysName           *string                                          `json:"SysName,omitempty" xml:"SysName,omitempty"`
	Timeout           *int32                                           `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ListGatewayApisResponseBodyApis) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayApisResponseBodyApis) GoString() string {
	return s.String()
}

func (s *ListGatewayApisResponseBodyApis) GetApiId() *int64 {
	return s.ApiId
}

func (s *ListGatewayApisResponseBodyApis) GetApiName() *string {
	return s.ApiName
}

func (s *ListGatewayApisResponseBodyApis) GetApiStatus() *string {
	return s.ApiStatus
}

func (s *ListGatewayApisResponseBodyApis) GetApiType() *string {
	return s.ApiType
}

func (s *ListGatewayApisResponseBodyApis) GetAuthRuleName() *string {
	return s.AuthRuleName
}

func (s *ListGatewayApisResponseBodyApis) GetCacheRule() *ListGatewayApisResponseBodyApisCacheRule {
	return s.CacheRule
}

func (s *ListGatewayApisResponseBodyApis) GetCharset() *string {
	return s.Charset
}

func (s *ListGatewayApisResponseBodyApis) GetContentType() *string {
	return s.ContentType
}

func (s *ListGatewayApisResponseBodyApis) GetDefaultLimitRule() *ListGatewayApisResponseBodyApisDefaultLimitRule {
	return s.DefaultLimitRule
}

func (s *ListGatewayApisResponseBodyApis) GetDescription() *string {
	return s.Description
}

func (s *ListGatewayApisResponseBodyApis) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListGatewayApisResponseBodyApis) GetHeaderRules() []*ListGatewayApisResponseBodyApisHeaderRules {
	return s.HeaderRules
}

func (s *ListGatewayApisResponseBodyApis) GetInterfaceType() *string {
	return s.InterfaceType
}

func (s *ListGatewayApisResponseBodyApis) GetLimitRule() *ListGatewayApisResponseBodyApisLimitRule {
	return s.LimitRule
}

func (s *ListGatewayApisResponseBodyApis) GetMethod() *string {
	return s.Method
}

func (s *ListGatewayApisResponseBodyApis) GetMethodName() *string {
	return s.MethodName
}

func (s *ListGatewayApisResponseBodyApis) GetMigrateRule() *ListGatewayApisResponseBodyApisMigrateRule {
	return s.MigrateRule
}

func (s *ListGatewayApisResponseBodyApis) GetMockRule() *ListGatewayApisResponseBodyApisMockRule {
	return s.MockRule
}

func (s *ListGatewayApisResponseBodyApis) GetNeedEncrypt() *bool {
	return s.NeedEncrypt
}

func (s *ListGatewayApisResponseBodyApis) GetNeedJsonp() *bool {
	return s.NeedJsonp
}

func (s *ListGatewayApisResponseBodyApis) GetNeedSign() *bool {
	return s.NeedSign
}

func (s *ListGatewayApisResponseBodyApis) GetOperationType() *string {
	return s.OperationType
}

func (s *ListGatewayApisResponseBodyApis) GetPath() *string {
	return s.Path
}

func (s *ListGatewayApisResponseBodyApis) GetRequestBodyModel() *string {
	return s.RequestBodyModel
}

func (s *ListGatewayApisResponseBodyApis) GetRequestParams() []*ListGatewayApisResponseBodyApisRequestParams {
	return s.RequestParams
}

func (s *ListGatewayApisResponseBodyApis) GetResponseBodyModel() *string {
	return s.ResponseBodyModel
}

func (s *ListGatewayApisResponseBodyApis) GetRpcParameterTypes() []*string {
	return s.RpcParameterTypes
}

func (s *ListGatewayApisResponseBodyApis) GetSerialization() *string {
	return s.Serialization
}

func (s *ListGatewayApisResponseBodyApis) GetSysName() *string {
	return s.SysName
}

func (s *ListGatewayApisResponseBodyApis) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListGatewayApisResponseBodyApis) SetApiId(v int64) *ListGatewayApisResponseBodyApis {
	s.ApiId = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetApiName(v string) *ListGatewayApisResponseBodyApis {
	s.ApiName = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetApiStatus(v string) *ListGatewayApisResponseBodyApis {
	s.ApiStatus = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetApiType(v string) *ListGatewayApisResponseBodyApis {
	s.ApiType = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetAuthRuleName(v string) *ListGatewayApisResponseBodyApis {
	s.AuthRuleName = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetCacheRule(v *ListGatewayApisResponseBodyApisCacheRule) *ListGatewayApisResponseBodyApis {
	s.CacheRule = v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetCharset(v string) *ListGatewayApisResponseBodyApis {
	s.Charset = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetContentType(v string) *ListGatewayApisResponseBodyApis {
	s.ContentType = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetDefaultLimitRule(v *ListGatewayApisResponseBodyApisDefaultLimitRule) *ListGatewayApisResponseBodyApis {
	s.DefaultLimitRule = v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetDescription(v string) *ListGatewayApisResponseBodyApis {
	s.Description = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetGroupId(v int64) *ListGatewayApisResponseBodyApis {
	s.GroupId = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetHeaderRules(v []*ListGatewayApisResponseBodyApisHeaderRules) *ListGatewayApisResponseBodyApis {
	s.HeaderRules = v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetInterfaceType(v string) *ListGatewayApisResponseBodyApis {
	s.InterfaceType = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetLimitRule(v *ListGatewayApisResponseBodyApisLimitRule) *ListGatewayApisResponseBodyApis {
	s.LimitRule = v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetMethod(v string) *ListGatewayApisResponseBodyApis {
	s.Method = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetMethodName(v string) *ListGatewayApisResponseBodyApis {
	s.MethodName = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetMigrateRule(v *ListGatewayApisResponseBodyApisMigrateRule) *ListGatewayApisResponseBodyApis {
	s.MigrateRule = v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetMockRule(v *ListGatewayApisResponseBodyApisMockRule) *ListGatewayApisResponseBodyApis {
	s.MockRule = v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetNeedEncrypt(v bool) *ListGatewayApisResponseBodyApis {
	s.NeedEncrypt = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetNeedJsonp(v bool) *ListGatewayApisResponseBodyApis {
	s.NeedJsonp = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetNeedSign(v bool) *ListGatewayApisResponseBodyApis {
	s.NeedSign = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetOperationType(v string) *ListGatewayApisResponseBodyApis {
	s.OperationType = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetPath(v string) *ListGatewayApisResponseBodyApis {
	s.Path = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetRequestBodyModel(v string) *ListGatewayApisResponseBodyApis {
	s.RequestBodyModel = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetRequestParams(v []*ListGatewayApisResponseBodyApisRequestParams) *ListGatewayApisResponseBodyApis {
	s.RequestParams = v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetResponseBodyModel(v string) *ListGatewayApisResponseBodyApis {
	s.ResponseBodyModel = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetRpcParameterTypes(v []*string) *ListGatewayApisResponseBodyApis {
	s.RpcParameterTypes = v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetSerialization(v string) *ListGatewayApisResponseBodyApis {
	s.Serialization = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetSysName(v string) *ListGatewayApisResponseBodyApis {
	s.SysName = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) SetTimeout(v int32) *ListGatewayApisResponseBodyApis {
	s.Timeout = &v
	return s
}

func (s *ListGatewayApisResponseBodyApis) Validate() error {
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

type ListGatewayApisResponseBodyApisCacheRule struct {
	CacheKey  *string `json:"CacheKey,omitempty" xml:"CacheKey,omitempty"`
	NeedCache *bool   `json:"NeedCache,omitempty" xml:"NeedCache,omitempty"`
	Ttl       *int64  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s ListGatewayApisResponseBodyApisCacheRule) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayApisResponseBodyApisCacheRule) GoString() string {
	return s.String()
}

func (s *ListGatewayApisResponseBodyApisCacheRule) GetCacheKey() *string {
	return s.CacheKey
}

func (s *ListGatewayApisResponseBodyApisCacheRule) GetNeedCache() *bool {
	return s.NeedCache
}

func (s *ListGatewayApisResponseBodyApisCacheRule) GetTtl() *int64 {
	return s.Ttl
}

func (s *ListGatewayApisResponseBodyApisCacheRule) SetCacheKey(v string) *ListGatewayApisResponseBodyApisCacheRule {
	s.CacheKey = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisCacheRule) SetNeedCache(v bool) *ListGatewayApisResponseBodyApisCacheRule {
	s.NeedCache = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisCacheRule) SetTtl(v int64) *ListGatewayApisResponseBodyApisCacheRule {
	s.Ttl = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisCacheRule) Validate() error {
	return dara.Validate(s)
}

type ListGatewayApisResponseBodyApisDefaultLimitRule struct {
	DefaultLimit *bool `json:"DefaultLimit,omitempty" xml:"DefaultLimit,omitempty"`
}

func (s ListGatewayApisResponseBodyApisDefaultLimitRule) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayApisResponseBodyApisDefaultLimitRule) GoString() string {
	return s.String()
}

func (s *ListGatewayApisResponseBodyApisDefaultLimitRule) GetDefaultLimit() *bool {
	return s.DefaultLimit
}

func (s *ListGatewayApisResponseBodyApisDefaultLimitRule) SetDefaultLimit(v bool) *ListGatewayApisResponseBodyApisDefaultLimitRule {
	s.DefaultLimit = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisDefaultLimitRule) Validate() error {
	return dara.Validate(s)
}

type ListGatewayApisResponseBodyApisHeaderRules struct {
	HeaderKey *string `json:"HeaderKey,omitempty" xml:"HeaderKey,omitempty"`
	Location  *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListGatewayApisResponseBodyApisHeaderRules) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayApisResponseBodyApisHeaderRules) GoString() string {
	return s.String()
}

func (s *ListGatewayApisResponseBodyApisHeaderRules) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *ListGatewayApisResponseBodyApisHeaderRules) GetLocation() *string {
	return s.Location
}

func (s *ListGatewayApisResponseBodyApisHeaderRules) GetType() *string {
	return s.Type
}

func (s *ListGatewayApisResponseBodyApisHeaderRules) GetValue() *string {
	return s.Value
}

func (s *ListGatewayApisResponseBodyApisHeaderRules) SetHeaderKey(v string) *ListGatewayApisResponseBodyApisHeaderRules {
	s.HeaderKey = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisHeaderRules) SetLocation(v string) *ListGatewayApisResponseBodyApisHeaderRules {
	s.Location = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisHeaderRules) SetType(v string) *ListGatewayApisResponseBodyApisHeaderRules {
	s.Type = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisHeaderRules) SetValue(v string) *ListGatewayApisResponseBodyApisHeaderRules {
	s.Value = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisHeaderRules) Validate() error {
	return dara.Validate(s)
}

type ListGatewayApisResponseBodyApisLimitRule struct {
	DefaultResponse *string  `json:"DefaultResponse,omitempty" xml:"DefaultResponse,omitempty"`
	Limit           *float64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Mode            *string  `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ListGatewayApisResponseBodyApisLimitRule) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayApisResponseBodyApisLimitRule) GoString() string {
	return s.String()
}

func (s *ListGatewayApisResponseBodyApisLimitRule) GetDefaultResponse() *string {
	return s.DefaultResponse
}

func (s *ListGatewayApisResponseBodyApisLimitRule) GetLimit() *float64 {
	return s.Limit
}

func (s *ListGatewayApisResponseBodyApisLimitRule) GetMode() *string {
	return s.Mode
}

func (s *ListGatewayApisResponseBodyApisLimitRule) SetDefaultResponse(v string) *ListGatewayApisResponseBodyApisLimitRule {
	s.DefaultResponse = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisLimitRule) SetLimit(v float64) *ListGatewayApisResponseBodyApisLimitRule {
	s.Limit = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisLimitRule) SetMode(v string) *ListGatewayApisResponseBodyApisLimitRule {
	s.Mode = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisLimitRule) Validate() error {
	return dara.Validate(s)
}

type ListGatewayApisResponseBodyApisMigrateRule struct {
	FlowPercent          *int32  `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
	GroupId              *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	NeedMigrate          *bool   `json:"NeedMigrate,omitempty" xml:"NeedMigrate,omitempty"`
	NeedSwitchCompletely *bool   `json:"NeedSwitchCompletely,omitempty" xml:"NeedSwitchCompletely,omitempty"`
	UpstreamType         *string `json:"UpstreamType,omitempty" xml:"UpstreamType,omitempty"`
}

func (s ListGatewayApisResponseBodyApisMigrateRule) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayApisResponseBodyApisMigrateRule) GoString() string {
	return s.String()
}

func (s *ListGatewayApisResponseBodyApisMigrateRule) GetFlowPercent() *int32 {
	return s.FlowPercent
}

func (s *ListGatewayApisResponseBodyApisMigrateRule) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListGatewayApisResponseBodyApisMigrateRule) GetGroupName() *string {
	return s.GroupName
}

func (s *ListGatewayApisResponseBodyApisMigrateRule) GetNeedMigrate() *bool {
	return s.NeedMigrate
}

func (s *ListGatewayApisResponseBodyApisMigrateRule) GetNeedSwitchCompletely() *bool {
	return s.NeedSwitchCompletely
}

func (s *ListGatewayApisResponseBodyApisMigrateRule) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *ListGatewayApisResponseBodyApisMigrateRule) SetFlowPercent(v int32) *ListGatewayApisResponseBodyApisMigrateRule {
	s.FlowPercent = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisMigrateRule) SetGroupId(v int64) *ListGatewayApisResponseBodyApisMigrateRule {
	s.GroupId = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisMigrateRule) SetGroupName(v string) *ListGatewayApisResponseBodyApisMigrateRule {
	s.GroupName = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisMigrateRule) SetNeedMigrate(v bool) *ListGatewayApisResponseBodyApisMigrateRule {
	s.NeedMigrate = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisMigrateRule) SetNeedSwitchCompletely(v bool) *ListGatewayApisResponseBodyApisMigrateRule {
	s.NeedSwitchCompletely = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisMigrateRule) SetUpstreamType(v string) *ListGatewayApisResponseBodyApisMigrateRule {
	s.UpstreamType = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisMigrateRule) Validate() error {
	return dara.Validate(s)
}

type ListGatewayApisResponseBodyApisMockRule struct {
	MockData   *string `json:"MockData,omitempty" xml:"MockData,omitempty"`
	NeedMock   *bool   `json:"NeedMock,omitempty" xml:"NeedMock,omitempty"`
	Percentage *int32  `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGatewayApisResponseBodyApisMockRule) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayApisResponseBodyApisMockRule) GoString() string {
	return s.String()
}

func (s *ListGatewayApisResponseBodyApisMockRule) GetMockData() *string {
	return s.MockData
}

func (s *ListGatewayApisResponseBodyApisMockRule) GetNeedMock() *bool {
	return s.NeedMock
}

func (s *ListGatewayApisResponseBodyApisMockRule) GetPercentage() *int32 {
	return s.Percentage
}

func (s *ListGatewayApisResponseBodyApisMockRule) GetType() *string {
	return s.Type
}

func (s *ListGatewayApisResponseBodyApisMockRule) SetMockData(v string) *ListGatewayApisResponseBodyApisMockRule {
	s.MockData = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisMockRule) SetNeedMock(v bool) *ListGatewayApisResponseBodyApisMockRule {
	s.NeedMock = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisMockRule) SetPercentage(v int32) *ListGatewayApisResponseBodyApisMockRule {
	s.Percentage = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisMockRule) SetType(v string) *ListGatewayApisResponseBodyApisMockRule {
	s.Type = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisMockRule) Validate() error {
	return dara.Validate(s)
}

type ListGatewayApisResponseBodyApisRequestParams struct {
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RefType      *string `json:"RefType,omitempty" xml:"RefType,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGatewayApisResponseBodyApisRequestParams) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayApisResponseBodyApisRequestParams) GoString() string {
	return s.String()
}

func (s *ListGatewayApisResponseBodyApisRequestParams) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListGatewayApisResponseBodyApisRequestParams) GetDescription() *string {
	return s.Description
}

func (s *ListGatewayApisResponseBodyApisRequestParams) GetLocation() *string {
	return s.Location
}

func (s *ListGatewayApisResponseBodyApisRequestParams) GetName() *string {
	return s.Name
}

func (s *ListGatewayApisResponseBodyApisRequestParams) GetRefType() *string {
	return s.RefType
}

func (s *ListGatewayApisResponseBodyApisRequestParams) GetType() *string {
	return s.Type
}

func (s *ListGatewayApisResponseBodyApisRequestParams) SetDefaultValue(v string) *ListGatewayApisResponseBodyApisRequestParams {
	s.DefaultValue = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisRequestParams) SetDescription(v string) *ListGatewayApisResponseBodyApisRequestParams {
	s.Description = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisRequestParams) SetLocation(v string) *ListGatewayApisResponseBodyApisRequestParams {
	s.Location = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisRequestParams) SetName(v string) *ListGatewayApisResponseBodyApisRequestParams {
	s.Name = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisRequestParams) SetRefType(v string) *ListGatewayApisResponseBodyApisRequestParams {
	s.RefType = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisRequestParams) SetType(v string) *ListGatewayApisResponseBodyApisRequestParams {
	s.Type = &v
	return s
}

func (s *ListGatewayApisResponseBodyApisRequestParams) Validate() error {
	return dara.Validate(s)
}
