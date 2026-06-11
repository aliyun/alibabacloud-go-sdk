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
	// Additional information about the deployment status, such as error details.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The target deployment platform. For example, set this to `SAE` to deploy on Serverless App Engine.
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// Configuration settings for deploying to Serverless App Engine (SAE). This object is required when the `platform` is `SAE`.
	SaeConfig *HiMarketPortalDeployConfigSaeConfig `json:"saeConfig,omitempty" xml:"saeConfig,omitempty" type:"Struct"`
	// The current status of the deployment. Possible values include `succeeded`, `failed`, and `in_progress`.
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
	// The ID of the application in Serverless App Engine.
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The ID of the namespace that logically isolates the application.
	NamespaceId *string `json:"namespaceId,omitempty" xml:"namespaceId,omitempty"`
	// The name of the OIDC role that grants permissions to the application.
	OidcRoleName *string `json:"oidcRoleName,omitempty" xml:"oidcRoleName,omitempty"`
	// The ID of the region in which to deploy the application.
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The desired number of application replicas.
	Replicas *string `json:"replicas,omitempty" xml:"replicas,omitempty"`
	// The ID of the security group to apply to the application instances. A security group acts as a virtual firewall.
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The ID of the vSwitch within the specified VPC. Serverless App Engine launches application instances in the vSwitch\\"s zone.
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The ID of the VPC to connect the application to.
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
