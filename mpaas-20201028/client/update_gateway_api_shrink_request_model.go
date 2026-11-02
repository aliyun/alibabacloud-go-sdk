// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayApiShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v int64) *UpdateGatewayApiShrinkRequest
	GetApiId() *int64
	SetApiName(v string) *UpdateGatewayApiShrinkRequest
	GetApiName() *string
	SetApiStatus(v string) *UpdateGatewayApiShrinkRequest
	GetApiStatus() *string
	SetApiType(v string) *UpdateGatewayApiShrinkRequest
	GetApiType() *string
	SetAppId(v string) *UpdateGatewayApiShrinkRequest
	GetAppId() *string
	SetAuthRuleName(v string) *UpdateGatewayApiShrinkRequest
	GetAuthRuleName() *string
	SetCacheRuleShrink(v string) *UpdateGatewayApiShrinkRequest
	GetCacheRuleShrink() *string
	SetCharset(v string) *UpdateGatewayApiShrinkRequest
	GetCharset() *string
	SetContentType(v string) *UpdateGatewayApiShrinkRequest
	GetContentType() *string
	SetDefaultLimitRuleShrink(v string) *UpdateGatewayApiShrinkRequest
	GetDefaultLimitRuleShrink() *string
	SetDescription(v string) *UpdateGatewayApiShrinkRequest
	GetDescription() *string
	SetGroupId(v int64) *UpdateGatewayApiShrinkRequest
	GetGroupId() *int64
	SetHeaderRulesShrink(v string) *UpdateGatewayApiShrinkRequest
	GetHeaderRulesShrink() *string
	SetInterfaceType(v string) *UpdateGatewayApiShrinkRequest
	GetInterfaceType() *string
	SetLimitRuleShrink(v string) *UpdateGatewayApiShrinkRequest
	GetLimitRuleShrink() *string
	SetMethod(v string) *UpdateGatewayApiShrinkRequest
	GetMethod() *string
	SetMethodName(v string) *UpdateGatewayApiShrinkRequest
	GetMethodName() *string
	SetMigrateRuleShrink(v string) *UpdateGatewayApiShrinkRequest
	GetMigrateRuleShrink() *string
	SetMockRuleShrink(v string) *UpdateGatewayApiShrinkRequest
	GetMockRuleShrink() *string
	SetNeedEncrypt(v bool) *UpdateGatewayApiShrinkRequest
	GetNeedEncrypt() *bool
	SetNeedJsonp(v bool) *UpdateGatewayApiShrinkRequest
	GetNeedJsonp() *bool
	SetNeedSign(v bool) *UpdateGatewayApiShrinkRequest
	GetNeedSign() *bool
	SetOperationType(v string) *UpdateGatewayApiShrinkRequest
	GetOperationType() *string
	SetPath(v string) *UpdateGatewayApiShrinkRequest
	GetPath() *string
	SetRequestBodyModel(v string) *UpdateGatewayApiShrinkRequest
	GetRequestBodyModel() *string
	SetRequestParamsShrink(v string) *UpdateGatewayApiShrinkRequest
	GetRequestParamsShrink() *string
	SetResponseBodyModel(v string) *UpdateGatewayApiShrinkRequest
	GetResponseBodyModel() *string
	SetRpcParameterTypesShrink(v string) *UpdateGatewayApiShrinkRequest
	GetRpcParameterTypesShrink() *string
	SetSerialization(v string) *UpdateGatewayApiShrinkRequest
	GetSerialization() *string
	SetTenantId(v string) *UpdateGatewayApiShrinkRequest
	GetTenantId() *string
	SetTimeout(v int32) *UpdateGatewayApiShrinkRequest
	GetTimeout() *int32
	SetWorkspaceId(v string) *UpdateGatewayApiShrinkRequest
	GetWorkspaceId() *string
}

type UpdateGatewayApiShrinkRequest struct {
	// This parameter is required.
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// This parameter is required.
	ApiName   *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	ApiStatus *string `json:"ApiStatus,omitempty" xml:"ApiStatus,omitempty"`
	// This parameter is required.
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// This parameter is required.
	AppId                  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AuthRuleName           *string `json:"AuthRuleName,omitempty" xml:"AuthRuleName,omitempty"`
	CacheRuleShrink        *string `json:"CacheRule,omitempty" xml:"CacheRule,omitempty"`
	Charset                *string `json:"Charset,omitempty" xml:"Charset,omitempty"`
	ContentType            *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	DefaultLimitRuleShrink *string `json:"DefaultLimitRule,omitempty" xml:"DefaultLimitRule,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	GroupId           *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HeaderRulesShrink *string `json:"HeaderRules,omitempty" xml:"HeaderRules,omitempty"`
	InterfaceType     *string `json:"InterfaceType,omitempty" xml:"InterfaceType,omitempty"`
	LimitRuleShrink   *string `json:"LimitRule,omitempty" xml:"LimitRule,omitempty"`
	Method            *string `json:"Method,omitempty" xml:"Method,omitempty"`
	MethodName        *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	MigrateRuleShrink *string `json:"MigrateRule,omitempty" xml:"MigrateRule,omitempty"`
	MockRuleShrink    *string `json:"MockRule,omitempty" xml:"MockRule,omitempty"`
	NeedEncrypt       *bool   `json:"NeedEncrypt,omitempty" xml:"NeedEncrypt,omitempty"`
	NeedJsonp         *bool   `json:"NeedJsonp,omitempty" xml:"NeedJsonp,omitempty"`
	NeedSign          *bool   `json:"NeedSign,omitempty" xml:"NeedSign,omitempty"`
	// This parameter is required.
	OperationType           *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	Path                    *string `json:"Path,omitempty" xml:"Path,omitempty"`
	RequestBodyModel        *string `json:"RequestBodyModel,omitempty" xml:"RequestBodyModel,omitempty"`
	RequestParamsShrink     *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
	ResponseBodyModel       *string `json:"ResponseBodyModel,omitempty" xml:"ResponseBodyModel,omitempty"`
	RpcParameterTypesShrink *string `json:"RpcParameterTypes,omitempty" xml:"RpcParameterTypes,omitempty"`
	Serialization           *string `json:"Serialization,omitempty" xml:"Serialization,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Timeout  *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateGatewayApiShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayApiShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayApiShrinkRequest) GetApiId() *int64 {
	return s.ApiId
}

func (s *UpdateGatewayApiShrinkRequest) GetApiName() *string {
	return s.ApiName
}

func (s *UpdateGatewayApiShrinkRequest) GetApiStatus() *string {
	return s.ApiStatus
}

func (s *UpdateGatewayApiShrinkRequest) GetApiType() *string {
	return s.ApiType
}

func (s *UpdateGatewayApiShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateGatewayApiShrinkRequest) GetAuthRuleName() *string {
	return s.AuthRuleName
}

func (s *UpdateGatewayApiShrinkRequest) GetCacheRuleShrink() *string {
	return s.CacheRuleShrink
}

func (s *UpdateGatewayApiShrinkRequest) GetCharset() *string {
	return s.Charset
}

func (s *UpdateGatewayApiShrinkRequest) GetContentType() *string {
	return s.ContentType
}

func (s *UpdateGatewayApiShrinkRequest) GetDefaultLimitRuleShrink() *string {
	return s.DefaultLimitRuleShrink
}

func (s *UpdateGatewayApiShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateGatewayApiShrinkRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *UpdateGatewayApiShrinkRequest) GetHeaderRulesShrink() *string {
	return s.HeaderRulesShrink
}

func (s *UpdateGatewayApiShrinkRequest) GetInterfaceType() *string {
	return s.InterfaceType
}

func (s *UpdateGatewayApiShrinkRequest) GetLimitRuleShrink() *string {
	return s.LimitRuleShrink
}

func (s *UpdateGatewayApiShrinkRequest) GetMethod() *string {
	return s.Method
}

func (s *UpdateGatewayApiShrinkRequest) GetMethodName() *string {
	return s.MethodName
}

func (s *UpdateGatewayApiShrinkRequest) GetMigrateRuleShrink() *string {
	return s.MigrateRuleShrink
}

func (s *UpdateGatewayApiShrinkRequest) GetMockRuleShrink() *string {
	return s.MockRuleShrink
}

func (s *UpdateGatewayApiShrinkRequest) GetNeedEncrypt() *bool {
	return s.NeedEncrypt
}

func (s *UpdateGatewayApiShrinkRequest) GetNeedJsonp() *bool {
	return s.NeedJsonp
}

func (s *UpdateGatewayApiShrinkRequest) GetNeedSign() *bool {
	return s.NeedSign
}

func (s *UpdateGatewayApiShrinkRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *UpdateGatewayApiShrinkRequest) GetPath() *string {
	return s.Path
}

func (s *UpdateGatewayApiShrinkRequest) GetRequestBodyModel() *string {
	return s.RequestBodyModel
}

func (s *UpdateGatewayApiShrinkRequest) GetRequestParamsShrink() *string {
	return s.RequestParamsShrink
}

func (s *UpdateGatewayApiShrinkRequest) GetResponseBodyModel() *string {
	return s.ResponseBodyModel
}

func (s *UpdateGatewayApiShrinkRequest) GetRpcParameterTypesShrink() *string {
	return s.RpcParameterTypesShrink
}

func (s *UpdateGatewayApiShrinkRequest) GetSerialization() *string {
	return s.Serialization
}

func (s *UpdateGatewayApiShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateGatewayApiShrinkRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateGatewayApiShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateGatewayApiShrinkRequest) SetApiId(v int64) *UpdateGatewayApiShrinkRequest {
	s.ApiId = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetApiName(v string) *UpdateGatewayApiShrinkRequest {
	s.ApiName = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetApiStatus(v string) *UpdateGatewayApiShrinkRequest {
	s.ApiStatus = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetApiType(v string) *UpdateGatewayApiShrinkRequest {
	s.ApiType = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetAppId(v string) *UpdateGatewayApiShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetAuthRuleName(v string) *UpdateGatewayApiShrinkRequest {
	s.AuthRuleName = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetCacheRuleShrink(v string) *UpdateGatewayApiShrinkRequest {
	s.CacheRuleShrink = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetCharset(v string) *UpdateGatewayApiShrinkRequest {
	s.Charset = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetContentType(v string) *UpdateGatewayApiShrinkRequest {
	s.ContentType = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetDefaultLimitRuleShrink(v string) *UpdateGatewayApiShrinkRequest {
	s.DefaultLimitRuleShrink = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetDescription(v string) *UpdateGatewayApiShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetGroupId(v int64) *UpdateGatewayApiShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetHeaderRulesShrink(v string) *UpdateGatewayApiShrinkRequest {
	s.HeaderRulesShrink = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetInterfaceType(v string) *UpdateGatewayApiShrinkRequest {
	s.InterfaceType = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetLimitRuleShrink(v string) *UpdateGatewayApiShrinkRequest {
	s.LimitRuleShrink = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetMethod(v string) *UpdateGatewayApiShrinkRequest {
	s.Method = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetMethodName(v string) *UpdateGatewayApiShrinkRequest {
	s.MethodName = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetMigrateRuleShrink(v string) *UpdateGatewayApiShrinkRequest {
	s.MigrateRuleShrink = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetMockRuleShrink(v string) *UpdateGatewayApiShrinkRequest {
	s.MockRuleShrink = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetNeedEncrypt(v bool) *UpdateGatewayApiShrinkRequest {
	s.NeedEncrypt = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetNeedJsonp(v bool) *UpdateGatewayApiShrinkRequest {
	s.NeedJsonp = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetNeedSign(v bool) *UpdateGatewayApiShrinkRequest {
	s.NeedSign = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetOperationType(v string) *UpdateGatewayApiShrinkRequest {
	s.OperationType = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetPath(v string) *UpdateGatewayApiShrinkRequest {
	s.Path = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetRequestBodyModel(v string) *UpdateGatewayApiShrinkRequest {
	s.RequestBodyModel = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetRequestParamsShrink(v string) *UpdateGatewayApiShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetResponseBodyModel(v string) *UpdateGatewayApiShrinkRequest {
	s.ResponseBodyModel = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetRpcParameterTypesShrink(v string) *UpdateGatewayApiShrinkRequest {
	s.RpcParameterTypesShrink = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetSerialization(v string) *UpdateGatewayApiShrinkRequest {
	s.Serialization = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetTenantId(v string) *UpdateGatewayApiShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetTimeout(v int32) *UpdateGatewayApiShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) SetWorkspaceId(v string) *UpdateGatewayApiShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateGatewayApiShrinkRequest) Validate() error {
	return dara.Validate(s)
}
