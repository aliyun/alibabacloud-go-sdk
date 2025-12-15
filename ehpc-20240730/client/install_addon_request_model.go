// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAddonRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonName(v string) *InstallAddonRequest
	GetAddonName() *string
	SetAddonVersion(v string) *InstallAddonRequest
	GetAddonVersion() *string
	SetClusterId(v string) *InstallAddonRequest
	GetClusterId() *string
	SetResourcesSpec(v string) *InstallAddonRequest
	GetResourcesSpec() *string
	SetServicesSpec(v string) *InstallAddonRequest
	GetServicesSpec() *string
}

type InstallAddonRequest struct {
	// The addon name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login
	AddonName *string `json:"AddonName,omitempty" xml:"AddonName,omitempty"`
	// The addon version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	AddonVersion *string `json:"AddonVersion,omitempty" xml:"AddonVersion,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The resource configurations of the addon.
	//
	// This parameter is required.
	//
	// example:
	//
	// `{"EipResource": {"AutoCreate": true}, "EcsResources": [{"InstanceType": "ecs.c7.xlarge", "ImageId": "centos_7_6_xxx.vhd", "SystemDisk": {"Category": "cloud_essd", "Size": 40, "Level": "PL0"}]}`
	ResourcesSpec *string `json:"ResourcesSpec,omitempty" xml:"ResourcesSpec,omitempty"`
	// The service configurations of the addon.
	//
	// This parameter is required.
	//
	// example:
	//
	// `[{"ServiceName": "SSH", "ServiceAccessType": null, "ServiceAccessUrl": null, "NetworkACL": [{"IpProtocol": "TCP", "Port": 22, "SourceCidrIp": "0.0.0.0/0"}]}, {"ServiceName": "VNC", "ServiceAccessType": null, "ServiceAccessUrl": null, "NetworkACL": [{"IpProtocol": "TCP", "Port": 12016, "SourceCidrIp": "0.0.0.0/0"}]}]`
	ServicesSpec *string `json:"ServicesSpec,omitempty" xml:"ServicesSpec,omitempty"`
}

func (s InstallAddonRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallAddonRequest) GoString() string {
	return s.String()
}

func (s *InstallAddonRequest) GetAddonName() *string {
	return s.AddonName
}

func (s *InstallAddonRequest) GetAddonVersion() *string {
	return s.AddonVersion
}

func (s *InstallAddonRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *InstallAddonRequest) GetResourcesSpec() *string {
	return s.ResourcesSpec
}

func (s *InstallAddonRequest) GetServicesSpec() *string {
	return s.ServicesSpec
}

func (s *InstallAddonRequest) SetAddonName(v string) *InstallAddonRequest {
	s.AddonName = &v
	return s
}

func (s *InstallAddonRequest) SetAddonVersion(v string) *InstallAddonRequest {
	s.AddonVersion = &v
	return s
}

func (s *InstallAddonRequest) SetClusterId(v string) *InstallAddonRequest {
	s.ClusterId = &v
	return s
}

func (s *InstallAddonRequest) SetResourcesSpec(v string) *InstallAddonRequest {
	s.ResourcesSpec = &v
	return s
}

func (s *InstallAddonRequest) SetServicesSpec(v string) *InstallAddonRequest {
	s.ServicesSpec = &v
	return s
}

func (s *InstallAddonRequest) Validate() error {
	return dara.Validate(s)
}
