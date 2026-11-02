// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v int64) *UpdateGatewayApiRequest
	GetApiId() *int64
	SetApiName(v string) *UpdateGatewayApiRequest
	GetApiName() *string
	SetApiStatus(v string) *UpdateGatewayApiRequest
	GetApiStatus() *string
	SetApiType(v string) *UpdateGatewayApiRequest
	GetApiType() *string
	SetAppId(v string) *UpdateGatewayApiRequest
	GetAppId() *string
	SetAuthRuleName(v string) *UpdateGatewayApiRequest
	GetAuthRuleName() *string
	SetCacheRule(v *UpdateGatewayApiRequestCacheRule) *UpdateGatewayApiRequest
	GetCacheRule() *UpdateGatewayApiRequestCacheRule
	SetCharset(v string) *UpdateGatewayApiRequest
	GetCharset() *string
	SetContentType(v string) *UpdateGatewayApiRequest
	GetContentType() *string
	SetDefaultLimitRule(v *UpdateGatewayApiRequestDefaultLimitRule) *UpdateGatewayApiRequest
	GetDefaultLimitRule() *UpdateGatewayApiRequestDefaultLimitRule
	SetDescription(v string) *UpdateGatewayApiRequest
	GetDescription() *string
	SetGroupId(v int64) *UpdateGatewayApiRequest
	GetGroupId() *int64
	SetHeaderRules(v []*UpdateGatewayApiRequestHeaderRules) *UpdateGatewayApiRequest
	GetHeaderRules() []*UpdateGatewayApiRequestHeaderRules
	SetInterfaceType(v string) *UpdateGatewayApiRequest
	GetInterfaceType() *string
	SetLimitRule(v *UpdateGatewayApiRequestLimitRule) *UpdateGatewayApiRequest
	GetLimitRule() *UpdateGatewayApiRequestLimitRule
	SetMethod(v string) *UpdateGatewayApiRequest
	GetMethod() *string
	SetMethodName(v string) *UpdateGatewayApiRequest
	GetMethodName() *string
	SetMigrateRule(v *UpdateGatewayApiRequestMigrateRule) *UpdateGatewayApiRequest
	GetMigrateRule() *UpdateGatewayApiRequestMigrateRule
	SetMockRule(v *UpdateGatewayApiRequestMockRule) *UpdateGatewayApiRequest
	GetMockRule() *UpdateGatewayApiRequestMockRule
	SetNeedEncrypt(v bool) *UpdateGatewayApiRequest
	GetNeedEncrypt() *bool
	SetNeedJsonp(v bool) *UpdateGatewayApiRequest
	GetNeedJsonp() *bool
	SetNeedSign(v bool) *UpdateGatewayApiRequest
	GetNeedSign() *bool
	SetOperationType(v string) *UpdateGatewayApiRequest
	GetOperationType() *string
	SetPath(v string) *UpdateGatewayApiRequest
	GetPath() *string
	SetRequestBodyModel(v string) *UpdateGatewayApiRequest
	GetRequestBodyModel() *string
	SetRequestParams(v []*UpdateGatewayApiRequestRequestParams) *UpdateGatewayApiRequest
	GetRequestParams() []*UpdateGatewayApiRequestRequestParams
	SetResponseBodyModel(v string) *UpdateGatewayApiRequest
	GetResponseBodyModel() *string
	SetRpcParameterTypes(v []*string) *UpdateGatewayApiRequest
	GetRpcParameterTypes() []*string
	SetSerialization(v string) *UpdateGatewayApiRequest
	GetSerialization() *string
	SetTenantId(v string) *UpdateGatewayApiRequest
	GetTenantId() *string
	SetTimeout(v int32) *UpdateGatewayApiRequest
	GetTimeout() *int32
	SetWorkspaceId(v string) *UpdateGatewayApiRequest
	GetWorkspaceId() *string
}

type UpdateGatewayApiRequest struct {
	// This parameter is required.
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// This parameter is required.
	ApiName   *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	ApiStatus *string `json:"ApiStatus,omitempty" xml:"ApiStatus,omitempty"`
	// This parameter is required.
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// This parameter is required.
	AppId            *string                                  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AuthRuleName     *string                                  `json:"AuthRuleName,omitempty" xml:"AuthRuleName,omitempty"`
	CacheRule        *UpdateGatewayApiRequestCacheRule        `json:"CacheRule,omitempty" xml:"CacheRule,omitempty" type:"Struct"`
	Charset          *string                                  `json:"Charset,omitempty" xml:"Charset,omitempty"`
	ContentType      *string                                  `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	DefaultLimitRule *UpdateGatewayApiRequestDefaultLimitRule `json:"DefaultLimitRule,omitempty" xml:"DefaultLimitRule,omitempty" type:"Struct"`
	Description      *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	GroupId       *int64                                `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HeaderRules   []*UpdateGatewayApiRequestHeaderRules `json:"HeaderRules,omitempty" xml:"HeaderRules,omitempty" type:"Repeated"`
	InterfaceType *string                               `json:"InterfaceType,omitempty" xml:"InterfaceType,omitempty"`
	LimitRule     *UpdateGatewayApiRequestLimitRule     `json:"LimitRule,omitempty" xml:"LimitRule,omitempty" type:"Struct"`
	Method        *string                               `json:"Method,omitempty" xml:"Method,omitempty"`
	MethodName    *string                               `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	MigrateRule   *UpdateGatewayApiRequestMigrateRule   `json:"MigrateRule,omitempty" xml:"MigrateRule,omitempty" type:"Struct"`
	MockRule      *UpdateGatewayApiRequestMockRule      `json:"MockRule,omitempty" xml:"MockRule,omitempty" type:"Struct"`
	NeedEncrypt   *bool                                 `json:"NeedEncrypt,omitempty" xml:"NeedEncrypt,omitempty"`
	NeedJsonp     *bool                                 `json:"NeedJsonp,omitempty" xml:"NeedJsonp,omitempty"`
	NeedSign      *bool                                 `json:"NeedSign,omitempty" xml:"NeedSign,omitempty"`
	// This parameter is required.
	OperationType     *string                                 `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	Path              *string                                 `json:"Path,omitempty" xml:"Path,omitempty"`
	RequestBodyModel  *string                                 `json:"RequestBodyModel,omitempty" xml:"RequestBodyModel,omitempty"`
	RequestParams     []*UpdateGatewayApiRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Repeated"`
	ResponseBodyModel *string                                 `json:"ResponseBodyModel,omitempty" xml:"ResponseBodyModel,omitempty"`
	RpcParameterTypes []*string                               `json:"RpcParameterTypes,omitempty" xml:"RpcParameterTypes,omitempty" type:"Repeated"`
	Serialization     *string                                 `json:"Serialization,omitempty" xml:"Serialization,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Timeout  *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateGatewayApiRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayApiRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayApiRequest) GetApiId() *int64 {
	return s.ApiId
}

func (s *UpdateGatewayApiRequest) GetApiName() *string {
	return s.ApiName
}

func (s *UpdateGatewayApiRequest) GetApiStatus() *string {
	return s.ApiStatus
}

func (s *UpdateGatewayApiRequest) GetApiType() *string {
	return s.ApiType
}

func (s *UpdateGatewayApiRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateGatewayApiRequest) GetAuthRuleName() *string {
	return s.AuthRuleName
}

func (s *UpdateGatewayApiRequest) GetCacheRule() *UpdateGatewayApiRequestCacheRule {
	return s.CacheRule
}

func (s *UpdateGatewayApiRequest) GetCharset() *string {
	return s.Charset
}

func (s *UpdateGatewayApiRequest) GetContentType() *string {
	return s.ContentType
}

func (s *UpdateGatewayApiRequest) GetDefaultLimitRule() *UpdateGatewayApiRequestDefaultLimitRule {
	return s.DefaultLimitRule
}

func (s *UpdateGatewayApiRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateGatewayApiRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *UpdateGatewayApiRequest) GetHeaderRules() []*UpdateGatewayApiRequestHeaderRules {
	return s.HeaderRules
}

func (s *UpdateGatewayApiRequest) GetInterfaceType() *string {
	return s.InterfaceType
}

func (s *UpdateGatewayApiRequest) GetLimitRule() *UpdateGatewayApiRequestLimitRule {
	return s.LimitRule
}

func (s *UpdateGatewayApiRequest) GetMethod() *string {
	return s.Method
}

func (s *UpdateGatewayApiRequest) GetMethodName() *string {
	return s.MethodName
}

func (s *UpdateGatewayApiRequest) GetMigrateRule() *UpdateGatewayApiRequestMigrateRule {
	return s.MigrateRule
}

func (s *UpdateGatewayApiRequest) GetMockRule() *UpdateGatewayApiRequestMockRule {
	return s.MockRule
}

func (s *UpdateGatewayApiRequest) GetNeedEncrypt() *bool {
	return s.NeedEncrypt
}

func (s *UpdateGatewayApiRequest) GetNeedJsonp() *bool {
	return s.NeedJsonp
}

func (s *UpdateGatewayApiRequest) GetNeedSign() *bool {
	return s.NeedSign
}

func (s *UpdateGatewayApiRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *UpdateGatewayApiRequest) GetPath() *string {
	return s.Path
}

func (s *UpdateGatewayApiRequest) GetRequestBodyModel() *string {
	return s.RequestBodyModel
}

func (s *UpdateGatewayApiRequest) GetRequestParams() []*UpdateGatewayApiRequestRequestParams {
	return s.RequestParams
}

func (s *UpdateGatewayApiRequest) GetResponseBodyModel() *string {
	return s.ResponseBodyModel
}

func (s *UpdateGatewayApiRequest) GetRpcParameterTypes() []*string {
	return s.RpcParameterTypes
}

func (s *UpdateGatewayApiRequest) GetSerialization() *string {
	return s.Serialization
}

func (s *UpdateGatewayApiRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateGatewayApiRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateGatewayApiRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateGatewayApiRequest) SetApiId(v int64) *UpdateGatewayApiRequest {
	s.ApiId = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetApiName(v string) *UpdateGatewayApiRequest {
	s.ApiName = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetApiStatus(v string) *UpdateGatewayApiRequest {
	s.ApiStatus = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetApiType(v string) *UpdateGatewayApiRequest {
	s.ApiType = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetAppId(v string) *UpdateGatewayApiRequest {
	s.AppId = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetAuthRuleName(v string) *UpdateGatewayApiRequest {
	s.AuthRuleName = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetCacheRule(v *UpdateGatewayApiRequestCacheRule) *UpdateGatewayApiRequest {
	s.CacheRule = v
	return s
}

func (s *UpdateGatewayApiRequest) SetCharset(v string) *UpdateGatewayApiRequest {
	s.Charset = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetContentType(v string) *UpdateGatewayApiRequest {
	s.ContentType = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetDefaultLimitRule(v *UpdateGatewayApiRequestDefaultLimitRule) *UpdateGatewayApiRequest {
	s.DefaultLimitRule = v
	return s
}

func (s *UpdateGatewayApiRequest) SetDescription(v string) *UpdateGatewayApiRequest {
	s.Description = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetGroupId(v int64) *UpdateGatewayApiRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetHeaderRules(v []*UpdateGatewayApiRequestHeaderRules) *UpdateGatewayApiRequest {
	s.HeaderRules = v
	return s
}

func (s *UpdateGatewayApiRequest) SetInterfaceType(v string) *UpdateGatewayApiRequest {
	s.InterfaceType = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetLimitRule(v *UpdateGatewayApiRequestLimitRule) *UpdateGatewayApiRequest {
	s.LimitRule = v
	return s
}

func (s *UpdateGatewayApiRequest) SetMethod(v string) *UpdateGatewayApiRequest {
	s.Method = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetMethodName(v string) *UpdateGatewayApiRequest {
	s.MethodName = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetMigrateRule(v *UpdateGatewayApiRequestMigrateRule) *UpdateGatewayApiRequest {
	s.MigrateRule = v
	return s
}

func (s *UpdateGatewayApiRequest) SetMockRule(v *UpdateGatewayApiRequestMockRule) *UpdateGatewayApiRequest {
	s.MockRule = v
	return s
}

func (s *UpdateGatewayApiRequest) SetNeedEncrypt(v bool) *UpdateGatewayApiRequest {
	s.NeedEncrypt = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetNeedJsonp(v bool) *UpdateGatewayApiRequest {
	s.NeedJsonp = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetNeedSign(v bool) *UpdateGatewayApiRequest {
	s.NeedSign = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetOperationType(v string) *UpdateGatewayApiRequest {
	s.OperationType = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetPath(v string) *UpdateGatewayApiRequest {
	s.Path = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetRequestBodyModel(v string) *UpdateGatewayApiRequest {
	s.RequestBodyModel = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetRequestParams(v []*UpdateGatewayApiRequestRequestParams) *UpdateGatewayApiRequest {
	s.RequestParams = v
	return s
}

func (s *UpdateGatewayApiRequest) SetResponseBodyModel(v string) *UpdateGatewayApiRequest {
	s.ResponseBodyModel = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetRpcParameterTypes(v []*string) *UpdateGatewayApiRequest {
	s.RpcParameterTypes = v
	return s
}

func (s *UpdateGatewayApiRequest) SetSerialization(v string) *UpdateGatewayApiRequest {
	s.Serialization = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetTenantId(v string) *UpdateGatewayApiRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetTimeout(v int32) *UpdateGatewayApiRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateGatewayApiRequest) SetWorkspaceId(v string) *UpdateGatewayApiRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateGatewayApiRequest) Validate() error {
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

type UpdateGatewayApiRequestCacheRule struct {
	CacheKey  *string `json:"CacheKey,omitempty" xml:"CacheKey,omitempty"`
	NeedCache *bool   `json:"NeedCache,omitempty" xml:"NeedCache,omitempty"`
	Ttl       *int64  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s UpdateGatewayApiRequestCacheRule) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayApiRequestCacheRule) GoString() string {
	return s.String()
}

func (s *UpdateGatewayApiRequestCacheRule) GetCacheKey() *string {
	return s.CacheKey
}

func (s *UpdateGatewayApiRequestCacheRule) GetNeedCache() *bool {
	return s.NeedCache
}

func (s *UpdateGatewayApiRequestCacheRule) GetTtl() *int64 {
	return s.Ttl
}

func (s *UpdateGatewayApiRequestCacheRule) SetCacheKey(v string) *UpdateGatewayApiRequestCacheRule {
	s.CacheKey = &v
	return s
}

func (s *UpdateGatewayApiRequestCacheRule) SetNeedCache(v bool) *UpdateGatewayApiRequestCacheRule {
	s.NeedCache = &v
	return s
}

func (s *UpdateGatewayApiRequestCacheRule) SetTtl(v int64) *UpdateGatewayApiRequestCacheRule {
	s.Ttl = &v
	return s
}

func (s *UpdateGatewayApiRequestCacheRule) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayApiRequestDefaultLimitRule struct {
	DefaultLimit *bool `json:"DefaultLimit,omitempty" xml:"DefaultLimit,omitempty"`
}

func (s UpdateGatewayApiRequestDefaultLimitRule) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayApiRequestDefaultLimitRule) GoString() string {
	return s.String()
}

func (s *UpdateGatewayApiRequestDefaultLimitRule) GetDefaultLimit() *bool {
	return s.DefaultLimit
}

func (s *UpdateGatewayApiRequestDefaultLimitRule) SetDefaultLimit(v bool) *UpdateGatewayApiRequestDefaultLimitRule {
	s.DefaultLimit = &v
	return s
}

func (s *UpdateGatewayApiRequestDefaultLimitRule) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayApiRequestHeaderRules struct {
	HeaderKey *string `json:"HeaderKey,omitempty" xml:"HeaderKey,omitempty"`
	Location  *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateGatewayApiRequestHeaderRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayApiRequestHeaderRules) GoString() string {
	return s.String()
}

func (s *UpdateGatewayApiRequestHeaderRules) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *UpdateGatewayApiRequestHeaderRules) GetLocation() *string {
	return s.Location
}

func (s *UpdateGatewayApiRequestHeaderRules) GetType() *string {
	return s.Type
}

func (s *UpdateGatewayApiRequestHeaderRules) GetValue() *string {
	return s.Value
}

func (s *UpdateGatewayApiRequestHeaderRules) SetHeaderKey(v string) *UpdateGatewayApiRequestHeaderRules {
	s.HeaderKey = &v
	return s
}

func (s *UpdateGatewayApiRequestHeaderRules) SetLocation(v string) *UpdateGatewayApiRequestHeaderRules {
	s.Location = &v
	return s
}

func (s *UpdateGatewayApiRequestHeaderRules) SetType(v string) *UpdateGatewayApiRequestHeaderRules {
	s.Type = &v
	return s
}

func (s *UpdateGatewayApiRequestHeaderRules) SetValue(v string) *UpdateGatewayApiRequestHeaderRules {
	s.Value = &v
	return s
}

func (s *UpdateGatewayApiRequestHeaderRules) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayApiRequestLimitRule struct {
	DefaultResponse *string  `json:"DefaultResponse,omitempty" xml:"DefaultResponse,omitempty"`
	Limit           *float64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Mode            *string  `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s UpdateGatewayApiRequestLimitRule) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayApiRequestLimitRule) GoString() string {
	return s.String()
}

func (s *UpdateGatewayApiRequestLimitRule) GetDefaultResponse() *string {
	return s.DefaultResponse
}

func (s *UpdateGatewayApiRequestLimitRule) GetLimit() *float64 {
	return s.Limit
}

func (s *UpdateGatewayApiRequestLimitRule) GetMode() *string {
	return s.Mode
}

func (s *UpdateGatewayApiRequestLimitRule) SetDefaultResponse(v string) *UpdateGatewayApiRequestLimitRule {
	s.DefaultResponse = &v
	return s
}

func (s *UpdateGatewayApiRequestLimitRule) SetLimit(v float64) *UpdateGatewayApiRequestLimitRule {
	s.Limit = &v
	return s
}

func (s *UpdateGatewayApiRequestLimitRule) SetMode(v string) *UpdateGatewayApiRequestLimitRule {
	s.Mode = &v
	return s
}

func (s *UpdateGatewayApiRequestLimitRule) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayApiRequestMigrateRule struct {
	FlowPercent          *int32 `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
	GroupId              *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	NeedMigrate          *bool  `json:"NeedMigrate,omitempty" xml:"NeedMigrate,omitempty"`
	NeedSwitchCompletely *bool  `json:"NeedSwitchCompletely,omitempty" xml:"NeedSwitchCompletely,omitempty"`
}

func (s UpdateGatewayApiRequestMigrateRule) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayApiRequestMigrateRule) GoString() string {
	return s.String()
}

func (s *UpdateGatewayApiRequestMigrateRule) GetFlowPercent() *int32 {
	return s.FlowPercent
}

func (s *UpdateGatewayApiRequestMigrateRule) GetGroupId() *int64 {
	return s.GroupId
}

func (s *UpdateGatewayApiRequestMigrateRule) GetNeedMigrate() *bool {
	return s.NeedMigrate
}

func (s *UpdateGatewayApiRequestMigrateRule) GetNeedSwitchCompletely() *bool {
	return s.NeedSwitchCompletely
}

func (s *UpdateGatewayApiRequestMigrateRule) SetFlowPercent(v int32) *UpdateGatewayApiRequestMigrateRule {
	s.FlowPercent = &v
	return s
}

func (s *UpdateGatewayApiRequestMigrateRule) SetGroupId(v int64) *UpdateGatewayApiRequestMigrateRule {
	s.GroupId = &v
	return s
}

func (s *UpdateGatewayApiRequestMigrateRule) SetNeedMigrate(v bool) *UpdateGatewayApiRequestMigrateRule {
	s.NeedMigrate = &v
	return s
}

func (s *UpdateGatewayApiRequestMigrateRule) SetNeedSwitchCompletely(v bool) *UpdateGatewayApiRequestMigrateRule {
	s.NeedSwitchCompletely = &v
	return s
}

func (s *UpdateGatewayApiRequestMigrateRule) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayApiRequestMockRule struct {
	MockData   *string `json:"MockData,omitempty" xml:"MockData,omitempty"`
	NeedMock   *bool   `json:"NeedMock,omitempty" xml:"NeedMock,omitempty"`
	Percentage *int32  `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateGatewayApiRequestMockRule) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayApiRequestMockRule) GoString() string {
	return s.String()
}

func (s *UpdateGatewayApiRequestMockRule) GetMockData() *string {
	return s.MockData
}

func (s *UpdateGatewayApiRequestMockRule) GetNeedMock() *bool {
	return s.NeedMock
}

func (s *UpdateGatewayApiRequestMockRule) GetPercentage() *int32 {
	return s.Percentage
}

func (s *UpdateGatewayApiRequestMockRule) GetType() *string {
	return s.Type
}

func (s *UpdateGatewayApiRequestMockRule) SetMockData(v string) *UpdateGatewayApiRequestMockRule {
	s.MockData = &v
	return s
}

func (s *UpdateGatewayApiRequestMockRule) SetNeedMock(v bool) *UpdateGatewayApiRequestMockRule {
	s.NeedMock = &v
	return s
}

func (s *UpdateGatewayApiRequestMockRule) SetPercentage(v int32) *UpdateGatewayApiRequestMockRule {
	s.Percentage = &v
	return s
}

func (s *UpdateGatewayApiRequestMockRule) SetType(v string) *UpdateGatewayApiRequestMockRule {
	s.Type = &v
	return s
}

func (s *UpdateGatewayApiRequestMockRule) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayApiRequestRequestParams struct {
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RefType      *string `json:"RefType,omitempty" xml:"RefType,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateGatewayApiRequestRequestParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayApiRequestRequestParams) GoString() string {
	return s.String()
}

func (s *UpdateGatewayApiRequestRequestParams) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *UpdateGatewayApiRequestRequestParams) GetDescription() *string {
	return s.Description
}

func (s *UpdateGatewayApiRequestRequestParams) GetLocation() *string {
	return s.Location
}

func (s *UpdateGatewayApiRequestRequestParams) GetName() *string {
	return s.Name
}

func (s *UpdateGatewayApiRequestRequestParams) GetRefType() *string {
	return s.RefType
}

func (s *UpdateGatewayApiRequestRequestParams) GetType() *string {
	return s.Type
}

func (s *UpdateGatewayApiRequestRequestParams) SetDefaultValue(v string) *UpdateGatewayApiRequestRequestParams {
	s.DefaultValue = &v
	return s
}

func (s *UpdateGatewayApiRequestRequestParams) SetDescription(v string) *UpdateGatewayApiRequestRequestParams {
	s.Description = &v
	return s
}

func (s *UpdateGatewayApiRequestRequestParams) SetLocation(v string) *UpdateGatewayApiRequestRequestParams {
	s.Location = &v
	return s
}

func (s *UpdateGatewayApiRequestRequestParams) SetName(v string) *UpdateGatewayApiRequestRequestParams {
	s.Name = &v
	return s
}

func (s *UpdateGatewayApiRequestRequestParams) SetRefType(v string) *UpdateGatewayApiRequestRequestParams {
	s.RefType = &v
	return s
}

func (s *UpdateGatewayApiRequestRequestParams) SetType(v string) *UpdateGatewayApiRequestRequestParams {
	s.Type = &v
	return s
}

func (s *UpdateGatewayApiRequestRequestParams) Validate() error {
	return dara.Validate(s)
}
