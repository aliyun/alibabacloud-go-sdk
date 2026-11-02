// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *CreateGatewayApiRequest
	GetApiName() *string
	SetApiStatus(v string) *CreateGatewayApiRequest
	GetApiStatus() *string
	SetApiType(v string) *CreateGatewayApiRequest
	GetApiType() *string
	SetAppId(v string) *CreateGatewayApiRequest
	GetAppId() *string
	SetAuthRuleName(v string) *CreateGatewayApiRequest
	GetAuthRuleName() *string
	SetCacheRule(v *CreateGatewayApiRequestCacheRule) *CreateGatewayApiRequest
	GetCacheRule() *CreateGatewayApiRequestCacheRule
	SetCharset(v string) *CreateGatewayApiRequest
	GetCharset() *string
	SetContentType(v string) *CreateGatewayApiRequest
	GetContentType() *string
	SetDefaultLimitRule(v *CreateGatewayApiRequestDefaultLimitRule) *CreateGatewayApiRequest
	GetDefaultLimitRule() *CreateGatewayApiRequestDefaultLimitRule
	SetDescription(v string) *CreateGatewayApiRequest
	GetDescription() *string
	SetGroupId(v int64) *CreateGatewayApiRequest
	GetGroupId() *int64
	SetHeaderRules(v []*CreateGatewayApiRequestHeaderRules) *CreateGatewayApiRequest
	GetHeaderRules() []*CreateGatewayApiRequestHeaderRules
	SetInterfaceType(v string) *CreateGatewayApiRequest
	GetInterfaceType() *string
	SetLimitRule(v *CreateGatewayApiRequestLimitRule) *CreateGatewayApiRequest
	GetLimitRule() *CreateGatewayApiRequestLimitRule
	SetMethod(v string) *CreateGatewayApiRequest
	GetMethod() *string
	SetMethodName(v string) *CreateGatewayApiRequest
	GetMethodName() *string
	SetMigrateRule(v *CreateGatewayApiRequestMigrateRule) *CreateGatewayApiRequest
	GetMigrateRule() *CreateGatewayApiRequestMigrateRule
	SetMockRule(v *CreateGatewayApiRequestMockRule) *CreateGatewayApiRequest
	GetMockRule() *CreateGatewayApiRequestMockRule
	SetNeedEncrypt(v bool) *CreateGatewayApiRequest
	GetNeedEncrypt() *bool
	SetNeedJsonp(v bool) *CreateGatewayApiRequest
	GetNeedJsonp() *bool
	SetNeedSign(v bool) *CreateGatewayApiRequest
	GetNeedSign() *bool
	SetOperationType(v string) *CreateGatewayApiRequest
	GetOperationType() *string
	SetPath(v string) *CreateGatewayApiRequest
	GetPath() *string
	SetRequestBodyModel(v string) *CreateGatewayApiRequest
	GetRequestBodyModel() *string
	SetRequestParams(v []*CreateGatewayApiRequestRequestParams) *CreateGatewayApiRequest
	GetRequestParams() []*CreateGatewayApiRequestRequestParams
	SetResponseBodyModel(v string) *CreateGatewayApiRequest
	GetResponseBodyModel() *string
	SetRpcParameterTypes(v []*string) *CreateGatewayApiRequest
	GetRpcParameterTypes() []*string
	SetSerialization(v string) *CreateGatewayApiRequest
	GetSerialization() *string
	SetTenantId(v string) *CreateGatewayApiRequest
	GetTenantId() *string
	SetTimeout(v int32) *CreateGatewayApiRequest
	GetTimeout() *int32
	SetWorkspaceId(v string) *CreateGatewayApiRequest
	GetWorkspaceId() *string
}

type CreateGatewayApiRequest struct {
	// This parameter is required.
	ApiName   *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	ApiStatus *string `json:"ApiStatus,omitempty" xml:"ApiStatus,omitempty"`
	// This parameter is required.
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// This parameter is required.
	AppId            *string                                  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AuthRuleName     *string                                  `json:"AuthRuleName,omitempty" xml:"AuthRuleName,omitempty"`
	CacheRule        *CreateGatewayApiRequestCacheRule        `json:"CacheRule,omitempty" xml:"CacheRule,omitempty" type:"Struct"`
	Charset          *string                                  `json:"Charset,omitempty" xml:"Charset,omitempty"`
	ContentType      *string                                  `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	DefaultLimitRule *CreateGatewayApiRequestDefaultLimitRule `json:"DefaultLimitRule,omitempty" xml:"DefaultLimitRule,omitempty" type:"Struct"`
	Description      *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	GroupId       *int64                                `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HeaderRules   []*CreateGatewayApiRequestHeaderRules `json:"HeaderRules,omitempty" xml:"HeaderRules,omitempty" type:"Repeated"`
	InterfaceType *string                               `json:"InterfaceType,omitempty" xml:"InterfaceType,omitempty"`
	LimitRule     *CreateGatewayApiRequestLimitRule     `json:"LimitRule,omitempty" xml:"LimitRule,omitempty" type:"Struct"`
	Method        *string                               `json:"Method,omitempty" xml:"Method,omitempty"`
	MethodName    *string                               `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	MigrateRule   *CreateGatewayApiRequestMigrateRule   `json:"MigrateRule,omitempty" xml:"MigrateRule,omitempty" type:"Struct"`
	MockRule      *CreateGatewayApiRequestMockRule      `json:"MockRule,omitempty" xml:"MockRule,omitempty" type:"Struct"`
	NeedEncrypt   *bool                                 `json:"NeedEncrypt,omitempty" xml:"NeedEncrypt,omitempty"`
	NeedJsonp     *bool                                 `json:"NeedJsonp,omitempty" xml:"NeedJsonp,omitempty"`
	NeedSign      *bool                                 `json:"NeedSign,omitempty" xml:"NeedSign,omitempty"`
	// This parameter is required.
	OperationType     *string                                 `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	Path              *string                                 `json:"Path,omitempty" xml:"Path,omitempty"`
	RequestBodyModel  *string                                 `json:"RequestBodyModel,omitempty" xml:"RequestBodyModel,omitempty"`
	RequestParams     []*CreateGatewayApiRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Repeated"`
	ResponseBodyModel *string                                 `json:"ResponseBodyModel,omitempty" xml:"ResponseBodyModel,omitempty"`
	RpcParameterTypes []*string                               `json:"RpcParameterTypes,omitempty" xml:"RpcParameterTypes,omitempty" type:"Repeated"`
	Serialization     *string                                 `json:"Serialization,omitempty" xml:"Serialization,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Timeout  *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateGatewayApiRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiRequest) GetApiName() *string {
	return s.ApiName
}

func (s *CreateGatewayApiRequest) GetApiStatus() *string {
	return s.ApiStatus
}

func (s *CreateGatewayApiRequest) GetApiType() *string {
	return s.ApiType
}

func (s *CreateGatewayApiRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateGatewayApiRequest) GetAuthRuleName() *string {
	return s.AuthRuleName
}

func (s *CreateGatewayApiRequest) GetCacheRule() *CreateGatewayApiRequestCacheRule {
	return s.CacheRule
}

func (s *CreateGatewayApiRequest) GetCharset() *string {
	return s.Charset
}

func (s *CreateGatewayApiRequest) GetContentType() *string {
	return s.ContentType
}

func (s *CreateGatewayApiRequest) GetDefaultLimitRule() *CreateGatewayApiRequestDefaultLimitRule {
	return s.DefaultLimitRule
}

func (s *CreateGatewayApiRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGatewayApiRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateGatewayApiRequest) GetHeaderRules() []*CreateGatewayApiRequestHeaderRules {
	return s.HeaderRules
}

func (s *CreateGatewayApiRequest) GetInterfaceType() *string {
	return s.InterfaceType
}

func (s *CreateGatewayApiRequest) GetLimitRule() *CreateGatewayApiRequestLimitRule {
	return s.LimitRule
}

func (s *CreateGatewayApiRequest) GetMethod() *string {
	return s.Method
}

func (s *CreateGatewayApiRequest) GetMethodName() *string {
	return s.MethodName
}

func (s *CreateGatewayApiRequest) GetMigrateRule() *CreateGatewayApiRequestMigrateRule {
	return s.MigrateRule
}

func (s *CreateGatewayApiRequest) GetMockRule() *CreateGatewayApiRequestMockRule {
	return s.MockRule
}

func (s *CreateGatewayApiRequest) GetNeedEncrypt() *bool {
	return s.NeedEncrypt
}

func (s *CreateGatewayApiRequest) GetNeedJsonp() *bool {
	return s.NeedJsonp
}

func (s *CreateGatewayApiRequest) GetNeedSign() *bool {
	return s.NeedSign
}

func (s *CreateGatewayApiRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *CreateGatewayApiRequest) GetPath() *string {
	return s.Path
}

func (s *CreateGatewayApiRequest) GetRequestBodyModel() *string {
	return s.RequestBodyModel
}

func (s *CreateGatewayApiRequest) GetRequestParams() []*CreateGatewayApiRequestRequestParams {
	return s.RequestParams
}

func (s *CreateGatewayApiRequest) GetResponseBodyModel() *string {
	return s.ResponseBodyModel
}

func (s *CreateGatewayApiRequest) GetRpcParameterTypes() []*string {
	return s.RpcParameterTypes
}

func (s *CreateGatewayApiRequest) GetSerialization() *string {
	return s.Serialization
}

func (s *CreateGatewayApiRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGatewayApiRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateGatewayApiRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateGatewayApiRequest) SetApiName(v string) *CreateGatewayApiRequest {
	s.ApiName = &v
	return s
}

func (s *CreateGatewayApiRequest) SetApiStatus(v string) *CreateGatewayApiRequest {
	s.ApiStatus = &v
	return s
}

func (s *CreateGatewayApiRequest) SetApiType(v string) *CreateGatewayApiRequest {
	s.ApiType = &v
	return s
}

func (s *CreateGatewayApiRequest) SetAppId(v string) *CreateGatewayApiRequest {
	s.AppId = &v
	return s
}

func (s *CreateGatewayApiRequest) SetAuthRuleName(v string) *CreateGatewayApiRequest {
	s.AuthRuleName = &v
	return s
}

func (s *CreateGatewayApiRequest) SetCacheRule(v *CreateGatewayApiRequestCacheRule) *CreateGatewayApiRequest {
	s.CacheRule = v
	return s
}

func (s *CreateGatewayApiRequest) SetCharset(v string) *CreateGatewayApiRequest {
	s.Charset = &v
	return s
}

func (s *CreateGatewayApiRequest) SetContentType(v string) *CreateGatewayApiRequest {
	s.ContentType = &v
	return s
}

func (s *CreateGatewayApiRequest) SetDefaultLimitRule(v *CreateGatewayApiRequestDefaultLimitRule) *CreateGatewayApiRequest {
	s.DefaultLimitRule = v
	return s
}

func (s *CreateGatewayApiRequest) SetDescription(v string) *CreateGatewayApiRequest {
	s.Description = &v
	return s
}

func (s *CreateGatewayApiRequest) SetGroupId(v int64) *CreateGatewayApiRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGatewayApiRequest) SetHeaderRules(v []*CreateGatewayApiRequestHeaderRules) *CreateGatewayApiRequest {
	s.HeaderRules = v
	return s
}

func (s *CreateGatewayApiRequest) SetInterfaceType(v string) *CreateGatewayApiRequest {
	s.InterfaceType = &v
	return s
}

func (s *CreateGatewayApiRequest) SetLimitRule(v *CreateGatewayApiRequestLimitRule) *CreateGatewayApiRequest {
	s.LimitRule = v
	return s
}

func (s *CreateGatewayApiRequest) SetMethod(v string) *CreateGatewayApiRequest {
	s.Method = &v
	return s
}

func (s *CreateGatewayApiRequest) SetMethodName(v string) *CreateGatewayApiRequest {
	s.MethodName = &v
	return s
}

func (s *CreateGatewayApiRequest) SetMigrateRule(v *CreateGatewayApiRequestMigrateRule) *CreateGatewayApiRequest {
	s.MigrateRule = v
	return s
}

func (s *CreateGatewayApiRequest) SetMockRule(v *CreateGatewayApiRequestMockRule) *CreateGatewayApiRequest {
	s.MockRule = v
	return s
}

func (s *CreateGatewayApiRequest) SetNeedEncrypt(v bool) *CreateGatewayApiRequest {
	s.NeedEncrypt = &v
	return s
}

func (s *CreateGatewayApiRequest) SetNeedJsonp(v bool) *CreateGatewayApiRequest {
	s.NeedJsonp = &v
	return s
}

func (s *CreateGatewayApiRequest) SetNeedSign(v bool) *CreateGatewayApiRequest {
	s.NeedSign = &v
	return s
}

func (s *CreateGatewayApiRequest) SetOperationType(v string) *CreateGatewayApiRequest {
	s.OperationType = &v
	return s
}

func (s *CreateGatewayApiRequest) SetPath(v string) *CreateGatewayApiRequest {
	s.Path = &v
	return s
}

func (s *CreateGatewayApiRequest) SetRequestBodyModel(v string) *CreateGatewayApiRequest {
	s.RequestBodyModel = &v
	return s
}

func (s *CreateGatewayApiRequest) SetRequestParams(v []*CreateGatewayApiRequestRequestParams) *CreateGatewayApiRequest {
	s.RequestParams = v
	return s
}

func (s *CreateGatewayApiRequest) SetResponseBodyModel(v string) *CreateGatewayApiRequest {
	s.ResponseBodyModel = &v
	return s
}

func (s *CreateGatewayApiRequest) SetRpcParameterTypes(v []*string) *CreateGatewayApiRequest {
	s.RpcParameterTypes = v
	return s
}

func (s *CreateGatewayApiRequest) SetSerialization(v string) *CreateGatewayApiRequest {
	s.Serialization = &v
	return s
}

func (s *CreateGatewayApiRequest) SetTenantId(v string) *CreateGatewayApiRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGatewayApiRequest) SetTimeout(v int32) *CreateGatewayApiRequest {
	s.Timeout = &v
	return s
}

func (s *CreateGatewayApiRequest) SetWorkspaceId(v string) *CreateGatewayApiRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateGatewayApiRequest) Validate() error {
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

type CreateGatewayApiRequestCacheRule struct {
	CacheKey  *string `json:"CacheKey,omitempty" xml:"CacheKey,omitempty"`
	NeedCache *bool   `json:"NeedCache,omitempty" xml:"NeedCache,omitempty"`
	Ttl       *int64  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s CreateGatewayApiRequestCacheRule) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiRequestCacheRule) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiRequestCacheRule) GetCacheKey() *string {
	return s.CacheKey
}

func (s *CreateGatewayApiRequestCacheRule) GetNeedCache() *bool {
	return s.NeedCache
}

func (s *CreateGatewayApiRequestCacheRule) GetTtl() *int64 {
	return s.Ttl
}

func (s *CreateGatewayApiRequestCacheRule) SetCacheKey(v string) *CreateGatewayApiRequestCacheRule {
	s.CacheKey = &v
	return s
}

func (s *CreateGatewayApiRequestCacheRule) SetNeedCache(v bool) *CreateGatewayApiRequestCacheRule {
	s.NeedCache = &v
	return s
}

func (s *CreateGatewayApiRequestCacheRule) SetTtl(v int64) *CreateGatewayApiRequestCacheRule {
	s.Ttl = &v
	return s
}

func (s *CreateGatewayApiRequestCacheRule) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayApiRequestDefaultLimitRule struct {
	DefaultLimit *bool `json:"DefaultLimit,omitempty" xml:"DefaultLimit,omitempty"`
}

func (s CreateGatewayApiRequestDefaultLimitRule) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiRequestDefaultLimitRule) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiRequestDefaultLimitRule) GetDefaultLimit() *bool {
	return s.DefaultLimit
}

func (s *CreateGatewayApiRequestDefaultLimitRule) SetDefaultLimit(v bool) *CreateGatewayApiRequestDefaultLimitRule {
	s.DefaultLimit = &v
	return s
}

func (s *CreateGatewayApiRequestDefaultLimitRule) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayApiRequestHeaderRules struct {
	HeaderKey *string `json:"HeaderKey,omitempty" xml:"HeaderKey,omitempty"`
	Location  *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateGatewayApiRequestHeaderRules) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiRequestHeaderRules) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiRequestHeaderRules) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *CreateGatewayApiRequestHeaderRules) GetLocation() *string {
	return s.Location
}

func (s *CreateGatewayApiRequestHeaderRules) GetType() *string {
	return s.Type
}

func (s *CreateGatewayApiRequestHeaderRules) GetValue() *string {
	return s.Value
}

func (s *CreateGatewayApiRequestHeaderRules) SetHeaderKey(v string) *CreateGatewayApiRequestHeaderRules {
	s.HeaderKey = &v
	return s
}

func (s *CreateGatewayApiRequestHeaderRules) SetLocation(v string) *CreateGatewayApiRequestHeaderRules {
	s.Location = &v
	return s
}

func (s *CreateGatewayApiRequestHeaderRules) SetType(v string) *CreateGatewayApiRequestHeaderRules {
	s.Type = &v
	return s
}

func (s *CreateGatewayApiRequestHeaderRules) SetValue(v string) *CreateGatewayApiRequestHeaderRules {
	s.Value = &v
	return s
}

func (s *CreateGatewayApiRequestHeaderRules) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayApiRequestLimitRule struct {
	DefaultResponse *string  `json:"DefaultResponse,omitempty" xml:"DefaultResponse,omitempty"`
	Limit           *float64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Mode            *string  `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s CreateGatewayApiRequestLimitRule) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiRequestLimitRule) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiRequestLimitRule) GetDefaultResponse() *string {
	return s.DefaultResponse
}

func (s *CreateGatewayApiRequestLimitRule) GetLimit() *float64 {
	return s.Limit
}

func (s *CreateGatewayApiRequestLimitRule) GetMode() *string {
	return s.Mode
}

func (s *CreateGatewayApiRequestLimitRule) SetDefaultResponse(v string) *CreateGatewayApiRequestLimitRule {
	s.DefaultResponse = &v
	return s
}

func (s *CreateGatewayApiRequestLimitRule) SetLimit(v float64) *CreateGatewayApiRequestLimitRule {
	s.Limit = &v
	return s
}

func (s *CreateGatewayApiRequestLimitRule) SetMode(v string) *CreateGatewayApiRequestLimitRule {
	s.Mode = &v
	return s
}

func (s *CreateGatewayApiRequestLimitRule) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayApiRequestMigrateRule struct {
	FlowPercent          *int32 `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
	GroupId              *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	NeedMigrate          *bool  `json:"NeedMigrate,omitempty" xml:"NeedMigrate,omitempty"`
	NeedSwitchCompletely *bool  `json:"NeedSwitchCompletely,omitempty" xml:"NeedSwitchCompletely,omitempty"`
}

func (s CreateGatewayApiRequestMigrateRule) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiRequestMigrateRule) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiRequestMigrateRule) GetFlowPercent() *int32 {
	return s.FlowPercent
}

func (s *CreateGatewayApiRequestMigrateRule) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateGatewayApiRequestMigrateRule) GetNeedMigrate() *bool {
	return s.NeedMigrate
}

func (s *CreateGatewayApiRequestMigrateRule) GetNeedSwitchCompletely() *bool {
	return s.NeedSwitchCompletely
}

func (s *CreateGatewayApiRequestMigrateRule) SetFlowPercent(v int32) *CreateGatewayApiRequestMigrateRule {
	s.FlowPercent = &v
	return s
}

func (s *CreateGatewayApiRequestMigrateRule) SetGroupId(v int64) *CreateGatewayApiRequestMigrateRule {
	s.GroupId = &v
	return s
}

func (s *CreateGatewayApiRequestMigrateRule) SetNeedMigrate(v bool) *CreateGatewayApiRequestMigrateRule {
	s.NeedMigrate = &v
	return s
}

func (s *CreateGatewayApiRequestMigrateRule) SetNeedSwitchCompletely(v bool) *CreateGatewayApiRequestMigrateRule {
	s.NeedSwitchCompletely = &v
	return s
}

func (s *CreateGatewayApiRequestMigrateRule) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayApiRequestMockRule struct {
	MockData   *string `json:"MockData,omitempty" xml:"MockData,omitempty"`
	NeedMock   *bool   `json:"NeedMock,omitempty" xml:"NeedMock,omitempty"`
	Percentage *int32  `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateGatewayApiRequestMockRule) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiRequestMockRule) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiRequestMockRule) GetMockData() *string {
	return s.MockData
}

func (s *CreateGatewayApiRequestMockRule) GetNeedMock() *bool {
	return s.NeedMock
}

func (s *CreateGatewayApiRequestMockRule) GetPercentage() *int32 {
	return s.Percentage
}

func (s *CreateGatewayApiRequestMockRule) GetType() *string {
	return s.Type
}

func (s *CreateGatewayApiRequestMockRule) SetMockData(v string) *CreateGatewayApiRequestMockRule {
	s.MockData = &v
	return s
}

func (s *CreateGatewayApiRequestMockRule) SetNeedMock(v bool) *CreateGatewayApiRequestMockRule {
	s.NeedMock = &v
	return s
}

func (s *CreateGatewayApiRequestMockRule) SetPercentage(v int32) *CreateGatewayApiRequestMockRule {
	s.Percentage = &v
	return s
}

func (s *CreateGatewayApiRequestMockRule) SetType(v string) *CreateGatewayApiRequestMockRule {
	s.Type = &v
	return s
}

func (s *CreateGatewayApiRequestMockRule) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayApiRequestRequestParams struct {
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RefType      *string `json:"RefType,omitempty" xml:"RefType,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateGatewayApiRequestRequestParams) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiRequestRequestParams) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiRequestRequestParams) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *CreateGatewayApiRequestRequestParams) GetDescription() *string {
	return s.Description
}

func (s *CreateGatewayApiRequestRequestParams) GetLocation() *string {
	return s.Location
}

func (s *CreateGatewayApiRequestRequestParams) GetName() *string {
	return s.Name
}

func (s *CreateGatewayApiRequestRequestParams) GetRefType() *string {
	return s.RefType
}

func (s *CreateGatewayApiRequestRequestParams) GetType() *string {
	return s.Type
}

func (s *CreateGatewayApiRequestRequestParams) SetDefaultValue(v string) *CreateGatewayApiRequestRequestParams {
	s.DefaultValue = &v
	return s
}

func (s *CreateGatewayApiRequestRequestParams) SetDescription(v string) *CreateGatewayApiRequestRequestParams {
	s.Description = &v
	return s
}

func (s *CreateGatewayApiRequestRequestParams) SetLocation(v string) *CreateGatewayApiRequestRequestParams {
	s.Location = &v
	return s
}

func (s *CreateGatewayApiRequestRequestParams) SetName(v string) *CreateGatewayApiRequestRequestParams {
	s.Name = &v
	return s
}

func (s *CreateGatewayApiRequestRequestParams) SetRefType(v string) *CreateGatewayApiRequestRequestParams {
	s.RefType = &v
	return s
}

func (s *CreateGatewayApiRequestRequestParams) SetType(v string) *CreateGatewayApiRequestRequestParams {
	s.Type = &v
	return s
}

func (s *CreateGatewayApiRequestRequestParams) Validate() error {
	return dara.Validate(s)
}
