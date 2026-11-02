// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayApiShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *CreateGatewayApiShrinkRequest
	GetApiName() *string
	SetApiStatus(v string) *CreateGatewayApiShrinkRequest
	GetApiStatus() *string
	SetApiType(v string) *CreateGatewayApiShrinkRequest
	GetApiType() *string
	SetAppId(v string) *CreateGatewayApiShrinkRequest
	GetAppId() *string
	SetAuthRuleName(v string) *CreateGatewayApiShrinkRequest
	GetAuthRuleName() *string
	SetCacheRuleShrink(v string) *CreateGatewayApiShrinkRequest
	GetCacheRuleShrink() *string
	SetCharset(v string) *CreateGatewayApiShrinkRequest
	GetCharset() *string
	SetContentType(v string) *CreateGatewayApiShrinkRequest
	GetContentType() *string
	SetDefaultLimitRuleShrink(v string) *CreateGatewayApiShrinkRequest
	GetDefaultLimitRuleShrink() *string
	SetDescription(v string) *CreateGatewayApiShrinkRequest
	GetDescription() *string
	SetGroupId(v int64) *CreateGatewayApiShrinkRequest
	GetGroupId() *int64
	SetHeaderRulesShrink(v string) *CreateGatewayApiShrinkRequest
	GetHeaderRulesShrink() *string
	SetInterfaceType(v string) *CreateGatewayApiShrinkRequest
	GetInterfaceType() *string
	SetLimitRuleShrink(v string) *CreateGatewayApiShrinkRequest
	GetLimitRuleShrink() *string
	SetMethod(v string) *CreateGatewayApiShrinkRequest
	GetMethod() *string
	SetMethodName(v string) *CreateGatewayApiShrinkRequest
	GetMethodName() *string
	SetMigrateRuleShrink(v string) *CreateGatewayApiShrinkRequest
	GetMigrateRuleShrink() *string
	SetMockRuleShrink(v string) *CreateGatewayApiShrinkRequest
	GetMockRuleShrink() *string
	SetNeedEncrypt(v bool) *CreateGatewayApiShrinkRequest
	GetNeedEncrypt() *bool
	SetNeedJsonp(v bool) *CreateGatewayApiShrinkRequest
	GetNeedJsonp() *bool
	SetNeedSign(v bool) *CreateGatewayApiShrinkRequest
	GetNeedSign() *bool
	SetOperationType(v string) *CreateGatewayApiShrinkRequest
	GetOperationType() *string
	SetPath(v string) *CreateGatewayApiShrinkRequest
	GetPath() *string
	SetRequestBodyModel(v string) *CreateGatewayApiShrinkRequest
	GetRequestBodyModel() *string
	SetRequestParamsShrink(v string) *CreateGatewayApiShrinkRequest
	GetRequestParamsShrink() *string
	SetResponseBodyModel(v string) *CreateGatewayApiShrinkRequest
	GetResponseBodyModel() *string
	SetRpcParameterTypesShrink(v string) *CreateGatewayApiShrinkRequest
	GetRpcParameterTypesShrink() *string
	SetSerialization(v string) *CreateGatewayApiShrinkRequest
	GetSerialization() *string
	SetTenantId(v string) *CreateGatewayApiShrinkRequest
	GetTenantId() *string
	SetTimeout(v int32) *CreateGatewayApiShrinkRequest
	GetTimeout() *int32
	SetWorkspaceId(v string) *CreateGatewayApiShrinkRequest
	GetWorkspaceId() *string
}

type CreateGatewayApiShrinkRequest struct {
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

func (s CreateGatewayApiShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiShrinkRequest) GetApiName() *string {
	return s.ApiName
}

func (s *CreateGatewayApiShrinkRequest) GetApiStatus() *string {
	return s.ApiStatus
}

func (s *CreateGatewayApiShrinkRequest) GetApiType() *string {
	return s.ApiType
}

func (s *CreateGatewayApiShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateGatewayApiShrinkRequest) GetAuthRuleName() *string {
	return s.AuthRuleName
}

func (s *CreateGatewayApiShrinkRequest) GetCacheRuleShrink() *string {
	return s.CacheRuleShrink
}

func (s *CreateGatewayApiShrinkRequest) GetCharset() *string {
	return s.Charset
}

func (s *CreateGatewayApiShrinkRequest) GetContentType() *string {
	return s.ContentType
}

func (s *CreateGatewayApiShrinkRequest) GetDefaultLimitRuleShrink() *string {
	return s.DefaultLimitRuleShrink
}

func (s *CreateGatewayApiShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGatewayApiShrinkRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateGatewayApiShrinkRequest) GetHeaderRulesShrink() *string {
	return s.HeaderRulesShrink
}

func (s *CreateGatewayApiShrinkRequest) GetInterfaceType() *string {
	return s.InterfaceType
}

func (s *CreateGatewayApiShrinkRequest) GetLimitRuleShrink() *string {
	return s.LimitRuleShrink
}

func (s *CreateGatewayApiShrinkRequest) GetMethod() *string {
	return s.Method
}

func (s *CreateGatewayApiShrinkRequest) GetMethodName() *string {
	return s.MethodName
}

func (s *CreateGatewayApiShrinkRequest) GetMigrateRuleShrink() *string {
	return s.MigrateRuleShrink
}

func (s *CreateGatewayApiShrinkRequest) GetMockRuleShrink() *string {
	return s.MockRuleShrink
}

func (s *CreateGatewayApiShrinkRequest) GetNeedEncrypt() *bool {
	return s.NeedEncrypt
}

func (s *CreateGatewayApiShrinkRequest) GetNeedJsonp() *bool {
	return s.NeedJsonp
}

func (s *CreateGatewayApiShrinkRequest) GetNeedSign() *bool {
	return s.NeedSign
}

func (s *CreateGatewayApiShrinkRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *CreateGatewayApiShrinkRequest) GetPath() *string {
	return s.Path
}

func (s *CreateGatewayApiShrinkRequest) GetRequestBodyModel() *string {
	return s.RequestBodyModel
}

func (s *CreateGatewayApiShrinkRequest) GetRequestParamsShrink() *string {
	return s.RequestParamsShrink
}

func (s *CreateGatewayApiShrinkRequest) GetResponseBodyModel() *string {
	return s.ResponseBodyModel
}

func (s *CreateGatewayApiShrinkRequest) GetRpcParameterTypesShrink() *string {
	return s.RpcParameterTypesShrink
}

func (s *CreateGatewayApiShrinkRequest) GetSerialization() *string {
	return s.Serialization
}

func (s *CreateGatewayApiShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGatewayApiShrinkRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateGatewayApiShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateGatewayApiShrinkRequest) SetApiName(v string) *CreateGatewayApiShrinkRequest {
	s.ApiName = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetApiStatus(v string) *CreateGatewayApiShrinkRequest {
	s.ApiStatus = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetApiType(v string) *CreateGatewayApiShrinkRequest {
	s.ApiType = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetAppId(v string) *CreateGatewayApiShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetAuthRuleName(v string) *CreateGatewayApiShrinkRequest {
	s.AuthRuleName = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetCacheRuleShrink(v string) *CreateGatewayApiShrinkRequest {
	s.CacheRuleShrink = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetCharset(v string) *CreateGatewayApiShrinkRequest {
	s.Charset = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetContentType(v string) *CreateGatewayApiShrinkRequest {
	s.ContentType = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetDefaultLimitRuleShrink(v string) *CreateGatewayApiShrinkRequest {
	s.DefaultLimitRuleShrink = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetDescription(v string) *CreateGatewayApiShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetGroupId(v int64) *CreateGatewayApiShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetHeaderRulesShrink(v string) *CreateGatewayApiShrinkRequest {
	s.HeaderRulesShrink = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetInterfaceType(v string) *CreateGatewayApiShrinkRequest {
	s.InterfaceType = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetLimitRuleShrink(v string) *CreateGatewayApiShrinkRequest {
	s.LimitRuleShrink = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetMethod(v string) *CreateGatewayApiShrinkRequest {
	s.Method = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetMethodName(v string) *CreateGatewayApiShrinkRequest {
	s.MethodName = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetMigrateRuleShrink(v string) *CreateGatewayApiShrinkRequest {
	s.MigrateRuleShrink = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetMockRuleShrink(v string) *CreateGatewayApiShrinkRequest {
	s.MockRuleShrink = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetNeedEncrypt(v bool) *CreateGatewayApiShrinkRequest {
	s.NeedEncrypt = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetNeedJsonp(v bool) *CreateGatewayApiShrinkRequest {
	s.NeedJsonp = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetNeedSign(v bool) *CreateGatewayApiShrinkRequest {
	s.NeedSign = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetOperationType(v string) *CreateGatewayApiShrinkRequest {
	s.OperationType = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetPath(v string) *CreateGatewayApiShrinkRequest {
	s.Path = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetRequestBodyModel(v string) *CreateGatewayApiShrinkRequest {
	s.RequestBodyModel = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetRequestParamsShrink(v string) *CreateGatewayApiShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetResponseBodyModel(v string) *CreateGatewayApiShrinkRequest {
	s.ResponseBodyModel = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetRpcParameterTypesShrink(v string) *CreateGatewayApiShrinkRequest {
	s.RpcParameterTypesShrink = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetSerialization(v string) *CreateGatewayApiShrinkRequest {
	s.Serialization = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetTenantId(v string) *CreateGatewayApiShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetTimeout(v int32) *CreateGatewayApiShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) SetWorkspaceId(v string) *CreateGatewayApiShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateGatewayApiShrinkRequest) Validate() error {
	return dara.Validate(s)
}
