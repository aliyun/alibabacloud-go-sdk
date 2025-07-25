// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBasePath(v string) *ModifyApiGroupRequest
	GetBasePath() *string
	SetCompatibleFlags(v string) *ModifyApiGroupRequest
	GetCompatibleFlags() *string
	SetCustomAppCodeConfig(v string) *ModifyApiGroupRequest
	GetCustomAppCodeConfig() *string
	SetCustomTraceConfig(v string) *ModifyApiGroupRequest
	GetCustomTraceConfig() *string
	SetCustomerConfigs(v string) *ModifyApiGroupRequest
	GetCustomerConfigs() *string
	SetDefaultDomain(v string) *ModifyApiGroupRequest
	GetDefaultDomain() *string
	SetDescription(v string) *ModifyApiGroupRequest
	GetDescription() *string
	SetFilterAppCodeForBackend(v string) *ModifyApiGroupRequest
	GetFilterAppCodeForBackend() *string
	SetGroupId(v string) *ModifyApiGroupRequest
	GetGroupId() *string
	SetGroupName(v string) *ModifyApiGroupRequest
	GetGroupName() *string
	SetPassthroughHeaders(v string) *ModifyApiGroupRequest
	GetPassthroughHeaders() *string
	SetRpcPattern(v string) *ModifyApiGroupRequest
	GetRpcPattern() *string
	SetRpsLimitForServerless(v string) *ModifyApiGroupRequest
	GetRpsLimitForServerless() *string
	SetSecurityToken(v string) *ModifyApiGroupRequest
	GetSecurityToken() *string
	SetSupportSSE(v string) *ModifyApiGroupRequest
	GetSupportSSE() *string
	SetTag(v []*ModifyApiGroupRequestTag) *ModifyApiGroupRequest
	GetTag() []*ModifyApiGroupRequestTag
	SetUserLogConfig(v string) *ModifyApiGroupRequest
	GetUserLogConfig() *string
}

type ModifyApiGroupRequest struct {
	// The root path of the API.
	//
	// example:
	//
	// /eeee
	BasePath *string `json:"BasePath,omitempty" xml:"BasePath,omitempty"`
	// The list of associated tags. Separate multiple tags with commas (,).
	//
	// example:
	//
	// depart:dep1
	CompatibleFlags *string `json:"CompatibleFlags,omitempty" xml:"CompatibleFlags,omitempty"`
	// The custom appcode configuration.
	//
	// example:
	//
	// {"location":"HEADER","name":"myAppCodeHeader"}
	CustomAppCodeConfig *string `json:"CustomAppCodeConfig,omitempty" xml:"CustomAppCodeConfig,omitempty"`
	// The custom trace configuration.
	//
	// example:
	//
	// {\\"parameterLocation\\":\\"HEADER\\",\\"parameterName\\":\\"traceId\\"}
	CustomTraceConfig *string `json:"CustomTraceConfig,omitempty" xml:"CustomTraceConfig,omitempty"`
	// The data of custom configuration items.
	//
	// example:
	//
	// removeResponseServerHeader
	CustomerConfigs *string `json:"CustomerConfigs,omitempty" xml:"CustomerConfigs,omitempty"`
	// The default domain name.
	//
	// example:
	//
	// mkt.api.gaore.com
	DefaultDomain *string `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	// The API group description that you want to specify, which cannot exceed 180 characters. If this parameter is not specified, the group description is not modified.
	//
	// example:
	//
	// New weather informations.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// If filter AppCode for backend.
	//
	// example:
	//
	// false
	FilterAppCodeForBackend *string `json:"FilterAppCodeForBackend,omitempty" xml:"FilterAppCodeForBackend,omitempty"`
	// The ID of the API group. This ID is generated by the system and globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 523e8dc7bbe04613b5b1d726c2a7889d
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The API group name must be globally unique. The name must be 4 to 50 characters in length. It must start with a letter and can contain letters, digits, and underscores (_). If this parameter is not specified, the group name is not modified.
	//
	// example:
	//
	// NewWeather
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Specifies whether to pass headers.
	//
	// example:
	//
	// eagleeye-rpcid,x-b3-traceid,host
	PassthroughHeaders *string `json:"PassthroughHeaders,omitempty" xml:"PassthroughHeaders,omitempty"`
	// The RPC mode.
	//
	// example:
	//
	// {}
	RpcPattern            *string `json:"RpcPattern,omitempty" xml:"RpcPattern,omitempty"`
	RpsLimitForServerless *string `json:"RpsLimitForServerless,omitempty" xml:"RpsLimitForServerless,omitempty"`
	SecurityToken         *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// If support SSE.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	SupportSSE *string `json:"SupportSSE,omitempty" xml:"SupportSSE,omitempty"`
	// The object tags that match the lifecycle rule. You can specify multiple tags.
	//
	// example:
	//
	// Key， Value
	Tag []*ModifyApiGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The user log configuration.
	//
	// example:
	//
	// {\\"requestBody\\":false,\\"responseBody\\":false,\\"queryString\\":\\"\\",\\"requestHeaders\\":\\"\\",\\"responseHeaders\\":\\"\\",\\"jwtClaims\\":\\"\\"}
	UserLogConfig *string `json:"UserLogConfig,omitempty" xml:"UserLogConfig,omitempty"`
}

func (s ModifyApiGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupRequest) GetBasePath() *string {
	return s.BasePath
}

func (s *ModifyApiGroupRequest) GetCompatibleFlags() *string {
	return s.CompatibleFlags
}

func (s *ModifyApiGroupRequest) GetCustomAppCodeConfig() *string {
	return s.CustomAppCodeConfig
}

func (s *ModifyApiGroupRequest) GetCustomTraceConfig() *string {
	return s.CustomTraceConfig
}

func (s *ModifyApiGroupRequest) GetCustomerConfigs() *string {
	return s.CustomerConfigs
}

func (s *ModifyApiGroupRequest) GetDefaultDomain() *string {
	return s.DefaultDomain
}

func (s *ModifyApiGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyApiGroupRequest) GetFilterAppCodeForBackend() *string {
	return s.FilterAppCodeForBackend
}

func (s *ModifyApiGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyApiGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyApiGroupRequest) GetPassthroughHeaders() *string {
	return s.PassthroughHeaders
}

func (s *ModifyApiGroupRequest) GetRpcPattern() *string {
	return s.RpcPattern
}

func (s *ModifyApiGroupRequest) GetRpsLimitForServerless() *string {
	return s.RpsLimitForServerless
}

func (s *ModifyApiGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyApiGroupRequest) GetSupportSSE() *string {
	return s.SupportSSE
}

func (s *ModifyApiGroupRequest) GetTag() []*ModifyApiGroupRequestTag {
	return s.Tag
}

func (s *ModifyApiGroupRequest) GetUserLogConfig() *string {
	return s.UserLogConfig
}

func (s *ModifyApiGroupRequest) SetBasePath(v string) *ModifyApiGroupRequest {
	s.BasePath = &v
	return s
}

func (s *ModifyApiGroupRequest) SetCompatibleFlags(v string) *ModifyApiGroupRequest {
	s.CompatibleFlags = &v
	return s
}

func (s *ModifyApiGroupRequest) SetCustomAppCodeConfig(v string) *ModifyApiGroupRequest {
	s.CustomAppCodeConfig = &v
	return s
}

func (s *ModifyApiGroupRequest) SetCustomTraceConfig(v string) *ModifyApiGroupRequest {
	s.CustomTraceConfig = &v
	return s
}

func (s *ModifyApiGroupRequest) SetCustomerConfigs(v string) *ModifyApiGroupRequest {
	s.CustomerConfigs = &v
	return s
}

func (s *ModifyApiGroupRequest) SetDefaultDomain(v string) *ModifyApiGroupRequest {
	s.DefaultDomain = &v
	return s
}

func (s *ModifyApiGroupRequest) SetDescription(v string) *ModifyApiGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyApiGroupRequest) SetFilterAppCodeForBackend(v string) *ModifyApiGroupRequest {
	s.FilterAppCodeForBackend = &v
	return s
}

func (s *ModifyApiGroupRequest) SetGroupId(v string) *ModifyApiGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyApiGroupRequest) SetGroupName(v string) *ModifyApiGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyApiGroupRequest) SetPassthroughHeaders(v string) *ModifyApiGroupRequest {
	s.PassthroughHeaders = &v
	return s
}

func (s *ModifyApiGroupRequest) SetRpcPattern(v string) *ModifyApiGroupRequest {
	s.RpcPattern = &v
	return s
}

func (s *ModifyApiGroupRequest) SetRpsLimitForServerless(v string) *ModifyApiGroupRequest {
	s.RpsLimitForServerless = &v
	return s
}

func (s *ModifyApiGroupRequest) SetSecurityToken(v string) *ModifyApiGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyApiGroupRequest) SetSupportSSE(v string) *ModifyApiGroupRequest {
	s.SupportSSE = &v
	return s
}

func (s *ModifyApiGroupRequest) SetTag(v []*ModifyApiGroupRequestTag) *ModifyApiGroupRequest {
	s.Tag = v
	return s
}

func (s *ModifyApiGroupRequest) SetUserLogConfig(v string) *ModifyApiGroupRequest {
	s.UserLogConfig = &v
	return s
}

func (s *ModifyApiGroupRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyApiGroupRequestTag struct {
	// The key of the tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// uat
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyApiGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiGroupRequestTag) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *ModifyApiGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *ModifyApiGroupRequestTag) SetKey(v string) *ModifyApiGroupRequestTag {
	s.Key = &v
	return s
}

func (s *ModifyApiGroupRequestTag) SetValue(v string) *ModifyApiGroupRequestTag {
	s.Value = &v
	return s
}

func (s *ModifyApiGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
