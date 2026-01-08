// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateAppRequest
	GetAppName() *string
	SetAuthentication(v *CreateAppRequestAuthentication) *CreateAppRequest
	GetAuthentication() *CreateAppRequestAuthentication
	SetChargeType(v string) *CreateAppRequest
	GetChargeType() *string
	SetDescription(v string) *CreateAppRequest
	GetDescription() *string
	SetNetwork(v []*CreateAppRequestNetwork) *CreateAppRequest
	GetNetwork() []*CreateAppRequestNetwork
	SetPrivateNetwork(v []*CreateAppRequestPrivateNetwork) *CreateAppRequest
	GetPrivateNetwork() []*CreateAppRequestPrivateNetwork
	SetQuotaInfo(v *CreateAppRequestQuotaInfo) *CreateAppRequest
	GetQuotaInfo() *CreateAppRequestQuotaInfo
	SetRegionId(v string) *CreateAppRequest
	GetRegionId() *string
	SetScenario(v string) *CreateAppRequest
	GetScenario() *string
	SetTags(v []*CreateAppRequestTags) *CreateAppRequest
	GetTags() []*CreateAppRequestTags
	SetVersion(v string) *CreateAppRequest
	GetVersion() *string
	SetClientToken(v string) *CreateAppRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateAppRequest
	GetDryRun() *bool
}

type CreateAppRequest struct {
	// 应用名
	//
	// This parameter is required.
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// This parameter is required.
	Authentication *CreateAppRequestAuthentication `json:"authentication,omitempty" xml:"authentication,omitempty" type:"Struct"`
	// This parameter is required.
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// 应用备注
	Description    *string                           `json:"description,omitempty" xml:"description,omitempty"`
	Network        []*CreateAppRequestNetwork        `json:"network,omitempty" xml:"network,omitempty" type:"Repeated"`
	PrivateNetwork []*CreateAppRequestPrivateNetwork `json:"privateNetwork,omitempty" xml:"privateNetwork,omitempty" type:"Repeated"`
	QuotaInfo      *CreateAppRequestQuotaInfo        `json:"quotaInfo,omitempty" xml:"quotaInfo,omitempty" type:"Struct"`
	RegionId       *string                           `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Scenario       *string                           `json:"scenario,omitempty" xml:"scenario,omitempty"`
	Tags           []*CreateAppRequestTags           `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	Version        *string                           `json:"version,omitempty" xml:"version,omitempty"`
	ClientToken    *string                           `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	DryRun         *bool                             `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppRequest) GetAuthentication() *CreateAppRequestAuthentication {
	return s.Authentication
}

func (s *CreateAppRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateAppRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAppRequest) GetNetwork() []*CreateAppRequestNetwork {
	return s.Network
}

func (s *CreateAppRequest) GetPrivateNetwork() []*CreateAppRequestPrivateNetwork {
	return s.PrivateNetwork
}

func (s *CreateAppRequest) GetQuotaInfo() *CreateAppRequestQuotaInfo {
	return s.QuotaInfo
}

func (s *CreateAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAppRequest) GetScenario() *string {
	return s.Scenario
}

func (s *CreateAppRequest) GetTags() []*CreateAppRequestTags {
	return s.Tags
}

func (s *CreateAppRequest) GetVersion() *string {
	return s.Version
}

func (s *CreateAppRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAppRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateAppRequest) SetAppName(v string) *CreateAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppRequest) SetAuthentication(v *CreateAppRequestAuthentication) *CreateAppRequest {
	s.Authentication = v
	return s
}

func (s *CreateAppRequest) SetChargeType(v string) *CreateAppRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateAppRequest) SetDescription(v string) *CreateAppRequest {
	s.Description = &v
	return s
}

func (s *CreateAppRequest) SetNetwork(v []*CreateAppRequestNetwork) *CreateAppRequest {
	s.Network = v
	return s
}

func (s *CreateAppRequest) SetPrivateNetwork(v []*CreateAppRequestPrivateNetwork) *CreateAppRequest {
	s.PrivateNetwork = v
	return s
}

func (s *CreateAppRequest) SetQuotaInfo(v *CreateAppRequestQuotaInfo) *CreateAppRequest {
	s.QuotaInfo = v
	return s
}

func (s *CreateAppRequest) SetRegionId(v string) *CreateAppRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAppRequest) SetScenario(v string) *CreateAppRequest {
	s.Scenario = &v
	return s
}

func (s *CreateAppRequest) SetTags(v []*CreateAppRequestTags) *CreateAppRequest {
	s.Tags = v
	return s
}

func (s *CreateAppRequest) SetVersion(v string) *CreateAppRequest {
	s.Version = &v
	return s
}

func (s *CreateAppRequest) SetClientToken(v string) *CreateAppRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppRequest) SetDryRun(v bool) *CreateAppRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAppRequest) Validate() error {
	if s.Authentication != nil {
		if err := s.Authentication.Validate(); err != nil {
			return err
		}
	}
	if s.Network != nil {
		for _, item := range s.Network {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PrivateNetwork != nil {
		for _, item := range s.PrivateNetwork {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QuotaInfo != nil {
		if err := s.QuotaInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAppRequestAuthentication struct {
	BasicAuth []*CreateAppRequestAuthenticationBasicAuth `json:"basicAuth,omitempty" xml:"basicAuth,omitempty" type:"Repeated"`
}

func (s CreateAppRequestAuthentication) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestAuthentication) GoString() string {
	return s.String()
}

func (s *CreateAppRequestAuthentication) GetBasicAuth() []*CreateAppRequestAuthenticationBasicAuth {
	return s.BasicAuth
}

func (s *CreateAppRequestAuthentication) SetBasicAuth(v []*CreateAppRequestAuthenticationBasicAuth) *CreateAppRequestAuthentication {
	s.BasicAuth = v
	return s
}

func (s *CreateAppRequestAuthentication) Validate() error {
	if s.BasicAuth != nil {
		for _, item := range s.BasicAuth {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAppRequestAuthenticationBasicAuth struct {
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s CreateAppRequestAuthenticationBasicAuth) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestAuthenticationBasicAuth) GoString() string {
	return s.String()
}

func (s *CreateAppRequestAuthenticationBasicAuth) GetPassword() *string {
	return s.Password
}

func (s *CreateAppRequestAuthenticationBasicAuth) GetUsername() *string {
	return s.Username
}

func (s *CreateAppRequestAuthenticationBasicAuth) SetPassword(v string) *CreateAppRequestAuthenticationBasicAuth {
	s.Password = &v
	return s
}

func (s *CreateAppRequestAuthenticationBasicAuth) SetUsername(v string) *CreateAppRequestAuthenticationBasicAuth {
	s.Username = &v
	return s
}

func (s *CreateAppRequestAuthenticationBasicAuth) Validate() error {
	return dara.Validate(s)
}

type CreateAppRequestNetwork struct {
	Domain       *string                                `json:"domain,omitempty" xml:"domain,omitempty"`
	Enabled      *bool                                  `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Port         *int32                                 `json:"port,omitempty" xml:"port,omitempty"`
	Type         *string                                `json:"type,omitempty" xml:"type,omitempty"`
	WhiteIpGroup []*CreateAppRequestNetworkWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s CreateAppRequestNetwork) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestNetwork) GoString() string {
	return s.String()
}

func (s *CreateAppRequestNetwork) GetDomain() *string {
	return s.Domain
}

func (s *CreateAppRequestNetwork) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateAppRequestNetwork) GetPort() *int32 {
	return s.Port
}

func (s *CreateAppRequestNetwork) GetType() *string {
	return s.Type
}

func (s *CreateAppRequestNetwork) GetWhiteIpGroup() []*CreateAppRequestNetworkWhiteIpGroup {
	return s.WhiteIpGroup
}

func (s *CreateAppRequestNetwork) SetDomain(v string) *CreateAppRequestNetwork {
	s.Domain = &v
	return s
}

func (s *CreateAppRequestNetwork) SetEnabled(v bool) *CreateAppRequestNetwork {
	s.Enabled = &v
	return s
}

func (s *CreateAppRequestNetwork) SetPort(v int32) *CreateAppRequestNetwork {
	s.Port = &v
	return s
}

func (s *CreateAppRequestNetwork) SetType(v string) *CreateAppRequestNetwork {
	s.Type = &v
	return s
}

func (s *CreateAppRequestNetwork) SetWhiteIpGroup(v []*CreateAppRequestNetworkWhiteIpGroup) *CreateAppRequestNetwork {
	s.WhiteIpGroup = v
	return s
}

func (s *CreateAppRequestNetwork) Validate() error {
	if s.WhiteIpGroup != nil {
		for _, item := range s.WhiteIpGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAppRequestNetworkWhiteIpGroup struct {
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s CreateAppRequestNetworkWhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestNetworkWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *CreateAppRequestNetworkWhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateAppRequestNetworkWhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *CreateAppRequestNetworkWhiteIpGroup) SetGroupName(v string) *CreateAppRequestNetworkWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *CreateAppRequestNetworkWhiteIpGroup) SetIps(v []*string) *CreateAppRequestNetworkWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *CreateAppRequestNetworkWhiteIpGroup) Validate() error {
	return dara.Validate(s)
}

type CreateAppRequestPrivateNetwork struct {
	Enabled       *bool                                         `json:"enabled,omitempty" xml:"enabled,omitempty"`
	PvlEndpointId *string                                       `json:"pvlEndpointId,omitempty" xml:"pvlEndpointId,omitempty"`
	Type          *string                                       `json:"type,omitempty" xml:"type,omitempty"`
	VpcId         *string                                       `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	WhiteIpGroup  []*CreateAppRequestPrivateNetworkWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s CreateAppRequestPrivateNetwork) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestPrivateNetwork) GoString() string {
	return s.String()
}

func (s *CreateAppRequestPrivateNetwork) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateAppRequestPrivateNetwork) GetPvlEndpointId() *string {
	return s.PvlEndpointId
}

func (s *CreateAppRequestPrivateNetwork) GetType() *string {
	return s.Type
}

func (s *CreateAppRequestPrivateNetwork) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateAppRequestPrivateNetwork) GetWhiteIpGroup() []*CreateAppRequestPrivateNetworkWhiteIpGroup {
	return s.WhiteIpGroup
}

func (s *CreateAppRequestPrivateNetwork) SetEnabled(v bool) *CreateAppRequestPrivateNetwork {
	s.Enabled = &v
	return s
}

func (s *CreateAppRequestPrivateNetwork) SetPvlEndpointId(v string) *CreateAppRequestPrivateNetwork {
	s.PvlEndpointId = &v
	return s
}

func (s *CreateAppRequestPrivateNetwork) SetType(v string) *CreateAppRequestPrivateNetwork {
	s.Type = &v
	return s
}

func (s *CreateAppRequestPrivateNetwork) SetVpcId(v string) *CreateAppRequestPrivateNetwork {
	s.VpcId = &v
	return s
}

func (s *CreateAppRequestPrivateNetwork) SetWhiteIpGroup(v []*CreateAppRequestPrivateNetworkWhiteIpGroup) *CreateAppRequestPrivateNetwork {
	s.WhiteIpGroup = v
	return s
}

func (s *CreateAppRequestPrivateNetwork) Validate() error {
	if s.WhiteIpGroup != nil {
		for _, item := range s.WhiteIpGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAppRequestPrivateNetworkWhiteIpGroup struct {
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s CreateAppRequestPrivateNetworkWhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestPrivateNetworkWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *CreateAppRequestPrivateNetworkWhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateAppRequestPrivateNetworkWhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *CreateAppRequestPrivateNetworkWhiteIpGroup) SetGroupName(v string) *CreateAppRequestPrivateNetworkWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *CreateAppRequestPrivateNetworkWhiteIpGroup) SetIps(v []*string) *CreateAppRequestPrivateNetworkWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *CreateAppRequestPrivateNetworkWhiteIpGroup) Validate() error {
	return dara.Validate(s)
}

type CreateAppRequestQuotaInfo struct {
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	Cu      *int32  `json:"cu,omitempty" xml:"cu,omitempty"`
	Storage *int32  `json:"storage,omitempty" xml:"storage,omitempty"`
}

func (s CreateAppRequestQuotaInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestQuotaInfo) GoString() string {
	return s.String()
}

func (s *CreateAppRequestQuotaInfo) GetAppType() *string {
	return s.AppType
}

func (s *CreateAppRequestQuotaInfo) GetCu() *int32 {
	return s.Cu
}

func (s *CreateAppRequestQuotaInfo) GetStorage() *int32 {
	return s.Storage
}

func (s *CreateAppRequestQuotaInfo) SetAppType(v string) *CreateAppRequestQuotaInfo {
	s.AppType = &v
	return s
}

func (s *CreateAppRequestQuotaInfo) SetCu(v int32) *CreateAppRequestQuotaInfo {
	s.Cu = &v
	return s
}

func (s *CreateAppRequestQuotaInfo) SetStorage(v int32) *CreateAppRequestQuotaInfo {
	s.Storage = &v
	return s
}

func (s *CreateAppRequestQuotaInfo) Validate() error {
	return dara.Validate(s)
}

type CreateAppRequestTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateAppRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestTags) GoString() string {
	return s.String()
}

func (s *CreateAppRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateAppRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateAppRequestTags) SetKey(v string) *CreateAppRequestTags {
	s.Key = &v
	return s
}

func (s *CreateAppRequestTags) SetValue(v string) *CreateAppRequestTags {
	s.Value = &v
	return s
}

func (s *CreateAppRequestTags) Validate() error {
	return dara.Validate(s)
}
