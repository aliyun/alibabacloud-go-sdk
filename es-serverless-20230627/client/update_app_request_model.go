// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyReason(v string) *UpdateAppRequest
	GetApplyReason() *string
	SetAuthentication(v *UpdateAppRequestAuthentication) *UpdateAppRequest
	GetAuthentication() *UpdateAppRequestAuthentication
	SetContactInfo(v string) *UpdateAppRequest
	GetContactInfo() *string
	SetDescription(v string) *UpdateAppRequest
	GetDescription() *string
	SetLimiterInfo(v *UpdateAppRequestLimiterInfo) *UpdateAppRequest
	GetLimiterInfo() *UpdateAppRequestLimiterInfo
	SetNetwork(v []*UpdateAppRequestNetwork) *UpdateAppRequest
	GetNetwork() []*UpdateAppRequestNetwork
	SetPrivateNetwork(v []*UpdateAppRequestPrivateNetwork) *UpdateAppRequest
	GetPrivateNetwork() []*UpdateAppRequestPrivateNetwork
}

type UpdateAppRequest struct {
	ApplyReason    *string                         `json:"applyReason,omitempty" xml:"applyReason,omitempty"`
	Authentication *UpdateAppRequestAuthentication `json:"authentication,omitempty" xml:"authentication,omitempty" type:"Struct"`
	ContactInfo    *string                         `json:"contactInfo,omitempty" xml:"contactInfo,omitempty"`
	// 应用备注
	Description    *string                           `json:"description,omitempty" xml:"description,omitempty"`
	LimiterInfo    *UpdateAppRequestLimiterInfo      `json:"limiterInfo,omitempty" xml:"limiterInfo,omitempty" type:"Struct"`
	Network        []*UpdateAppRequestNetwork        `json:"network,omitempty" xml:"network,omitempty" type:"Repeated"`
	PrivateNetwork []*UpdateAppRequestPrivateNetwork `json:"privateNetwork,omitempty" xml:"privateNetwork,omitempty" type:"Repeated"`
}

func (s UpdateAppRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppRequest) GetApplyReason() *string {
	return s.ApplyReason
}

func (s *UpdateAppRequest) GetAuthentication() *UpdateAppRequestAuthentication {
	return s.Authentication
}

func (s *UpdateAppRequest) GetContactInfo() *string {
	return s.ContactInfo
}

func (s *UpdateAppRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAppRequest) GetLimiterInfo() *UpdateAppRequestLimiterInfo {
	return s.LimiterInfo
}

func (s *UpdateAppRequest) GetNetwork() []*UpdateAppRequestNetwork {
	return s.Network
}

func (s *UpdateAppRequest) GetPrivateNetwork() []*UpdateAppRequestPrivateNetwork {
	return s.PrivateNetwork
}

func (s *UpdateAppRequest) SetApplyReason(v string) *UpdateAppRequest {
	s.ApplyReason = &v
	return s
}

func (s *UpdateAppRequest) SetAuthentication(v *UpdateAppRequestAuthentication) *UpdateAppRequest {
	s.Authentication = v
	return s
}

func (s *UpdateAppRequest) SetContactInfo(v string) *UpdateAppRequest {
	s.ContactInfo = &v
	return s
}

func (s *UpdateAppRequest) SetDescription(v string) *UpdateAppRequest {
	s.Description = &v
	return s
}

func (s *UpdateAppRequest) SetLimiterInfo(v *UpdateAppRequestLimiterInfo) *UpdateAppRequest {
	s.LimiterInfo = v
	return s
}

func (s *UpdateAppRequest) SetNetwork(v []*UpdateAppRequestNetwork) *UpdateAppRequest {
	s.Network = v
	return s
}

func (s *UpdateAppRequest) SetPrivateNetwork(v []*UpdateAppRequestPrivateNetwork) *UpdateAppRequest {
	s.PrivateNetwork = v
	return s
}

func (s *UpdateAppRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateAppRequestAuthentication struct {
	BasicAuth []*UpdateAppRequestAuthenticationBasicAuth `json:"basicAuth,omitempty" xml:"basicAuth,omitempty" type:"Repeated"`
}

func (s UpdateAppRequestAuthentication) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppRequestAuthentication) GoString() string {
	return s.String()
}

func (s *UpdateAppRequestAuthentication) GetBasicAuth() []*UpdateAppRequestAuthenticationBasicAuth {
	return s.BasicAuth
}

func (s *UpdateAppRequestAuthentication) SetBasicAuth(v []*UpdateAppRequestAuthenticationBasicAuth) *UpdateAppRequestAuthentication {
	s.BasicAuth = v
	return s
}

func (s *UpdateAppRequestAuthentication) Validate() error {
	return dara.Validate(s)
}

type UpdateAppRequestAuthenticationBasicAuth struct {
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s UpdateAppRequestAuthenticationBasicAuth) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppRequestAuthenticationBasicAuth) GoString() string {
	return s.String()
}

func (s *UpdateAppRequestAuthenticationBasicAuth) GetPassword() *string {
	return s.Password
}

func (s *UpdateAppRequestAuthenticationBasicAuth) GetUsername() *string {
	return s.Username
}

func (s *UpdateAppRequestAuthenticationBasicAuth) SetPassword(v string) *UpdateAppRequestAuthenticationBasicAuth {
	s.Password = &v
	return s
}

func (s *UpdateAppRequestAuthenticationBasicAuth) SetUsername(v string) *UpdateAppRequestAuthenticationBasicAuth {
	s.Username = &v
	return s
}

func (s *UpdateAppRequestAuthenticationBasicAuth) Validate() error {
	return dara.Validate(s)
}

type UpdateAppRequestLimiterInfo struct {
	Limiters []*UpdateAppRequestLimiterInfoLimiters `json:"limiters,omitempty" xml:"limiters,omitempty" type:"Repeated"`
}

func (s UpdateAppRequestLimiterInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppRequestLimiterInfo) GoString() string {
	return s.String()
}

func (s *UpdateAppRequestLimiterInfo) GetLimiters() []*UpdateAppRequestLimiterInfoLimiters {
	return s.Limiters
}

func (s *UpdateAppRequestLimiterInfo) SetLimiters(v []*UpdateAppRequestLimiterInfoLimiters) *UpdateAppRequestLimiterInfo {
	s.Limiters = v
	return s
}

func (s *UpdateAppRequestLimiterInfo) Validate() error {
	return dara.Validate(s)
}

type UpdateAppRequestLimiterInfoLimiters struct {
	MaxValue *int32    `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue *int32    `json:"minValue,omitempty" xml:"minValue,omitempty"`
	Type     *string   `json:"type,omitempty" xml:"type,omitempty"`
	Values   []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s UpdateAppRequestLimiterInfoLimiters) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppRequestLimiterInfoLimiters) GoString() string {
	return s.String()
}

func (s *UpdateAppRequestLimiterInfoLimiters) GetMaxValue() *int32 {
	return s.MaxValue
}

func (s *UpdateAppRequestLimiterInfoLimiters) GetMinValue() *int32 {
	return s.MinValue
}

func (s *UpdateAppRequestLimiterInfoLimiters) GetType() *string {
	return s.Type
}

func (s *UpdateAppRequestLimiterInfoLimiters) GetValues() []*string {
	return s.Values
}

func (s *UpdateAppRequestLimiterInfoLimiters) SetMaxValue(v int32) *UpdateAppRequestLimiterInfoLimiters {
	s.MaxValue = &v
	return s
}

func (s *UpdateAppRequestLimiterInfoLimiters) SetMinValue(v int32) *UpdateAppRequestLimiterInfoLimiters {
	s.MinValue = &v
	return s
}

func (s *UpdateAppRequestLimiterInfoLimiters) SetType(v string) *UpdateAppRequestLimiterInfoLimiters {
	s.Type = &v
	return s
}

func (s *UpdateAppRequestLimiterInfoLimiters) SetValues(v []*string) *UpdateAppRequestLimiterInfoLimiters {
	s.Values = v
	return s
}

func (s *UpdateAppRequestLimiterInfoLimiters) Validate() error {
	return dara.Validate(s)
}

type UpdateAppRequestNetwork struct {
	Domain       *string                                `json:"domain,omitempty" xml:"domain,omitempty"`
	Enabled      *bool                                  `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Port         *int32                                 `json:"port,omitempty" xml:"port,omitempty"`
	Type         *string                                `json:"type,omitempty" xml:"type,omitempty"`
	WhiteIpGroup []*UpdateAppRequestNetworkWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s UpdateAppRequestNetwork) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppRequestNetwork) GoString() string {
	return s.String()
}

func (s *UpdateAppRequestNetwork) GetDomain() *string {
	return s.Domain
}

func (s *UpdateAppRequestNetwork) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateAppRequestNetwork) GetPort() *int32 {
	return s.Port
}

func (s *UpdateAppRequestNetwork) GetType() *string {
	return s.Type
}

func (s *UpdateAppRequestNetwork) GetWhiteIpGroup() []*UpdateAppRequestNetworkWhiteIpGroup {
	return s.WhiteIpGroup
}

func (s *UpdateAppRequestNetwork) SetDomain(v string) *UpdateAppRequestNetwork {
	s.Domain = &v
	return s
}

func (s *UpdateAppRequestNetwork) SetEnabled(v bool) *UpdateAppRequestNetwork {
	s.Enabled = &v
	return s
}

func (s *UpdateAppRequestNetwork) SetPort(v int32) *UpdateAppRequestNetwork {
	s.Port = &v
	return s
}

func (s *UpdateAppRequestNetwork) SetType(v string) *UpdateAppRequestNetwork {
	s.Type = &v
	return s
}

func (s *UpdateAppRequestNetwork) SetWhiteIpGroup(v []*UpdateAppRequestNetworkWhiteIpGroup) *UpdateAppRequestNetwork {
	s.WhiteIpGroup = v
	return s
}

func (s *UpdateAppRequestNetwork) Validate() error {
	return dara.Validate(s)
}

type UpdateAppRequestNetworkWhiteIpGroup struct {
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s UpdateAppRequestNetworkWhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppRequestNetworkWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *UpdateAppRequestNetworkWhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateAppRequestNetworkWhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *UpdateAppRequestNetworkWhiteIpGroup) SetGroupName(v string) *UpdateAppRequestNetworkWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *UpdateAppRequestNetworkWhiteIpGroup) SetIps(v []*string) *UpdateAppRequestNetworkWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *UpdateAppRequestNetworkWhiteIpGroup) Validate() error {
	return dara.Validate(s)
}

type UpdateAppRequestPrivateNetwork struct {
	Enabled       *bool                                         `json:"enabled,omitempty" xml:"enabled,omitempty"`
	PvlEndpointId *string                                       `json:"pvlEndpointId,omitempty" xml:"pvlEndpointId,omitempty"`
	Type          *string                                       `json:"type,omitempty" xml:"type,omitempty"`
	VpcId         *string                                       `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	WhiteIpGroup  []*UpdateAppRequestPrivateNetworkWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s UpdateAppRequestPrivateNetwork) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppRequestPrivateNetwork) GoString() string {
	return s.String()
}

func (s *UpdateAppRequestPrivateNetwork) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateAppRequestPrivateNetwork) GetPvlEndpointId() *string {
	return s.PvlEndpointId
}

func (s *UpdateAppRequestPrivateNetwork) GetType() *string {
	return s.Type
}

func (s *UpdateAppRequestPrivateNetwork) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateAppRequestPrivateNetwork) GetWhiteIpGroup() []*UpdateAppRequestPrivateNetworkWhiteIpGroup {
	return s.WhiteIpGroup
}

func (s *UpdateAppRequestPrivateNetwork) SetEnabled(v bool) *UpdateAppRequestPrivateNetwork {
	s.Enabled = &v
	return s
}

func (s *UpdateAppRequestPrivateNetwork) SetPvlEndpointId(v string) *UpdateAppRequestPrivateNetwork {
	s.PvlEndpointId = &v
	return s
}

func (s *UpdateAppRequestPrivateNetwork) SetType(v string) *UpdateAppRequestPrivateNetwork {
	s.Type = &v
	return s
}

func (s *UpdateAppRequestPrivateNetwork) SetVpcId(v string) *UpdateAppRequestPrivateNetwork {
	s.VpcId = &v
	return s
}

func (s *UpdateAppRequestPrivateNetwork) SetWhiteIpGroup(v []*UpdateAppRequestPrivateNetworkWhiteIpGroup) *UpdateAppRequestPrivateNetwork {
	s.WhiteIpGroup = v
	return s
}

func (s *UpdateAppRequestPrivateNetwork) Validate() error {
	return dara.Validate(s)
}

type UpdateAppRequestPrivateNetworkWhiteIpGroup struct {
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s UpdateAppRequestPrivateNetworkWhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppRequestPrivateNetworkWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *UpdateAppRequestPrivateNetworkWhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateAppRequestPrivateNetworkWhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *UpdateAppRequestPrivateNetworkWhiteIpGroup) SetGroupName(v string) *UpdateAppRequestPrivateNetworkWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *UpdateAppRequestPrivateNetworkWhiteIpGroup) SetIps(v []*string) *UpdateAppRequestPrivateNetworkWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *UpdateAppRequestPrivateNetworkWhiteIpGroup) Validate() error {
	return dara.Validate(s)
}
