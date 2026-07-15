// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketPortalDeployConfig interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *HiMarketPortalDeployConfig
	GetMessage() *string
	SetPlatform(v string) *HiMarketPortalDeployConfig
	GetPlatform() *string
	SetSaeConfig(v *HiMarketPortalDeployConfigSaeConfig) *HiMarketPortalDeployConfig
	GetSaeConfig() *HiMarketPortalDeployConfigSaeConfig
	SetStatus(v string) *HiMarketPortalDeployConfig
	GetStatus() *string
}

type HiMarketPortalDeployConfig struct {
	// The deployment status message.
	//
	// example:
	//
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The deployment platform type.
	//
	// example:
	//
	// SAE
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// The SAE (Serverless App Engine) deployment configuration.
	SaeConfig *HiMarketPortalDeployConfigSaeConfig `json:"saeConfig,omitempty" xml:"saeConfig,omitempty" type:"Struct"`
	// The deployment status.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s HiMarketPortalDeployConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketPortalDeployConfig) GoString() string {
	return s.String()
}

func (s *HiMarketPortalDeployConfig) GetMessage() *string {
	return s.Message
}

func (s *HiMarketPortalDeployConfig) GetPlatform() *string {
	return s.Platform
}

func (s *HiMarketPortalDeployConfig) GetSaeConfig() *HiMarketPortalDeployConfigSaeConfig {
	return s.SaeConfig
}

func (s *HiMarketPortalDeployConfig) GetStatus() *string {
	return s.Status
}

func (s *HiMarketPortalDeployConfig) SetMessage(v string) *HiMarketPortalDeployConfig {
	s.Message = &v
	return s
}

func (s *HiMarketPortalDeployConfig) SetPlatform(v string) *HiMarketPortalDeployConfig {
	s.Platform = &v
	return s
}

func (s *HiMarketPortalDeployConfig) SetSaeConfig(v *HiMarketPortalDeployConfigSaeConfig) *HiMarketPortalDeployConfig {
	s.SaeConfig = v
	return s
}

func (s *HiMarketPortalDeployConfig) SetStatus(v string) *HiMarketPortalDeployConfig {
	s.Status = &v
	return s
}

func (s *HiMarketPortalDeployConfig) Validate() error {
	if s.SaeConfig != nil {
		if err := s.SaeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HiMarketPortalDeployConfigSaeConfig struct {
	// The SAE application ID.
	//
	// example:
	//
	// app-xxx
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The SAE namespace ID.
	//
	// example:
	//
	// cn-hangzhou:default
	NamespaceId *string `json:"namespaceId,omitempty" xml:"namespaceId,omitempty"`
	// The OIDC role name used for SAE to access other cloud resources.
	//
	// example:
	//
	// apig-portal-oidc-role
	OidcRoleName *string `json:"oidcRoleName,omitempty" xml:"oidcRoleName,omitempty"`
	// The region ID of SAE.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The number of instance replicas.
	//
	// example:
	//
	// 2
	Replicas *string `json:"replicas,omitempty" xml:"replicas,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-xxx
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-xxx
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-xxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s HiMarketPortalDeployConfigSaeConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketPortalDeployConfigSaeConfig) GoString() string {
	return s.String()
}

func (s *HiMarketPortalDeployConfigSaeConfig) GetAppId() *string {
	return s.AppId
}

func (s *HiMarketPortalDeployConfigSaeConfig) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *HiMarketPortalDeployConfigSaeConfig) GetOidcRoleName() *string {
	return s.OidcRoleName
}

func (s *HiMarketPortalDeployConfigSaeConfig) GetRegionId() *string {
	return s.RegionId
}

func (s *HiMarketPortalDeployConfigSaeConfig) GetReplicas() *string {
	return s.Replicas
}

func (s *HiMarketPortalDeployConfigSaeConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *HiMarketPortalDeployConfigSaeConfig) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *HiMarketPortalDeployConfigSaeConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *HiMarketPortalDeployConfigSaeConfig) SetAppId(v string) *HiMarketPortalDeployConfigSaeConfig {
	s.AppId = &v
	return s
}

func (s *HiMarketPortalDeployConfigSaeConfig) SetNamespaceId(v string) *HiMarketPortalDeployConfigSaeConfig {
	s.NamespaceId = &v
	return s
}

func (s *HiMarketPortalDeployConfigSaeConfig) SetOidcRoleName(v string) *HiMarketPortalDeployConfigSaeConfig {
	s.OidcRoleName = &v
	return s
}

func (s *HiMarketPortalDeployConfigSaeConfig) SetRegionId(v string) *HiMarketPortalDeployConfigSaeConfig {
	s.RegionId = &v
	return s
}

func (s *HiMarketPortalDeployConfigSaeConfig) SetReplicas(v string) *HiMarketPortalDeployConfigSaeConfig {
	s.Replicas = &v
	return s
}

func (s *HiMarketPortalDeployConfigSaeConfig) SetSecurityGroupId(v string) *HiMarketPortalDeployConfigSaeConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *HiMarketPortalDeployConfigSaeConfig) SetVSwitchId(v string) *HiMarketPortalDeployConfigSaeConfig {
	s.VSwitchId = &v
	return s
}

func (s *HiMarketPortalDeployConfigSaeConfig) SetVpcId(v string) *HiMarketPortalDeployConfigSaeConfig {
	s.VpcId = &v
	return s
}

func (s *HiMarketPortalDeployConfigSaeConfig) Validate() error {
	return dara.Validate(s)
}
