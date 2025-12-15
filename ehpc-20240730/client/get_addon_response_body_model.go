// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAddonResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddon(v *GetAddonResponseBodyAddon) *GetAddonResponseBody
	GetAddon() *GetAddonResponseBodyAddon
	SetRequestId(v string) *GetAddonResponseBody
	GetRequestId() *string
}

type GetAddonResponseBody struct {
	// The information about the addon.
	Addon *GetAddonResponseBodyAddon `json:"Addon,omitempty" xml:"Addon,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BBC2F93D-003A-49C4-850C-B826EECF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAddonResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAddonResponseBody) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBody) GetAddon() *GetAddonResponseBodyAddon {
	return s.Addon
}

func (s *GetAddonResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAddonResponseBody) SetAddon(v *GetAddonResponseBodyAddon) *GetAddonResponseBody {
	s.Addon = v
	return s
}

func (s *GetAddonResponseBody) SetRequestId(v string) *GetAddonResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAddonResponseBody) Validate() error {
	if s.Addon != nil {
		if err := s.Addon.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAddonResponseBodyAddon struct {
	// The addon ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login-1.0-W2g****
	AddonId *string `json:"AddonId,omitempty" xml:"AddonId,omitempty"`
	// The addon description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The addon icon.
	//
	// example:
	//
	// /assets/icons/your_icon.svg
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The time when the addon was installed.
	//
	// example:
	//
	// 2024-08-22 18:11:17
	InstallTime *string `json:"InstallTime,omitempty" xml:"InstallTime,omitempty"`
	// The addon label.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The addon name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource configurations of the addon.
	ResourcesSpec *GetAddonResponseBodyAddonResourcesSpec `json:"ResourcesSpec,omitempty" xml:"ResourcesSpec,omitempty" type:"Struct"`
	// The service configurations of the addon.
	ServicesSpec []*GetAddonResponseBodyAddonServicesSpec `json:"ServicesSpec,omitempty" xml:"ServicesSpec,omitempty" type:"Repeated"`
	// The addon status.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The addon version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAddonResponseBodyAddon) String() string {
	return dara.Prettify(s)
}

func (s GetAddonResponseBodyAddon) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyAddon) GetAddonId() *string {
	return s.AddonId
}

func (s *GetAddonResponseBodyAddon) GetDescription() *string {
	return s.Description
}

func (s *GetAddonResponseBodyAddon) GetIcon() *string {
	return s.Icon
}

func (s *GetAddonResponseBodyAddon) GetInstallTime() *string {
	return s.InstallTime
}

func (s *GetAddonResponseBodyAddon) GetLabel() *string {
	return s.Label
}

func (s *GetAddonResponseBodyAddon) GetName() *string {
	return s.Name
}

func (s *GetAddonResponseBodyAddon) GetResourcesSpec() *GetAddonResponseBodyAddonResourcesSpec {
	return s.ResourcesSpec
}

func (s *GetAddonResponseBodyAddon) GetServicesSpec() []*GetAddonResponseBodyAddonServicesSpec {
	return s.ServicesSpec
}

func (s *GetAddonResponseBodyAddon) GetStatus() *string {
	return s.Status
}

func (s *GetAddonResponseBodyAddon) GetVersion() *string {
	return s.Version
}

func (s *GetAddonResponseBodyAddon) SetAddonId(v string) *GetAddonResponseBodyAddon {
	s.AddonId = &v
	return s
}

func (s *GetAddonResponseBodyAddon) SetDescription(v string) *GetAddonResponseBodyAddon {
	s.Description = &v
	return s
}

func (s *GetAddonResponseBodyAddon) SetIcon(v string) *GetAddonResponseBodyAddon {
	s.Icon = &v
	return s
}

func (s *GetAddonResponseBodyAddon) SetInstallTime(v string) *GetAddonResponseBodyAddon {
	s.InstallTime = &v
	return s
}

func (s *GetAddonResponseBodyAddon) SetLabel(v string) *GetAddonResponseBodyAddon {
	s.Label = &v
	return s
}

func (s *GetAddonResponseBodyAddon) SetName(v string) *GetAddonResponseBodyAddon {
	s.Name = &v
	return s
}

func (s *GetAddonResponseBodyAddon) SetResourcesSpec(v *GetAddonResponseBodyAddonResourcesSpec) *GetAddonResponseBodyAddon {
	s.ResourcesSpec = v
	return s
}

func (s *GetAddonResponseBodyAddon) SetServicesSpec(v []*GetAddonResponseBodyAddonServicesSpec) *GetAddonResponseBodyAddon {
	s.ServicesSpec = v
	return s
}

func (s *GetAddonResponseBodyAddon) SetStatus(v string) *GetAddonResponseBodyAddon {
	s.Status = &v
	return s
}

func (s *GetAddonResponseBodyAddon) SetVersion(v string) *GetAddonResponseBodyAddon {
	s.Version = &v
	return s
}

func (s *GetAddonResponseBodyAddon) Validate() error {
	if s.ResourcesSpec != nil {
		if err := s.ResourcesSpec.Validate(); err != nil {
			return err
		}
	}
	if s.ServicesSpec != nil {
		for _, item := range s.ServicesSpec {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAddonResponseBodyAddonResourcesSpec struct {
	// The Elastic Compute Service (ECS) resource configurations of the addon.
	EcsResources []*AddonNodeTemplate `json:"EcsResources,omitempty" xml:"EcsResources,omitempty" type:"Repeated"`
	// The Elastic IP Address (EIP) configurations.
	EipResource *GetAddonResponseBodyAddonResourcesSpecEipResource `json:"EipResource,omitempty" xml:"EipResource,omitempty" type:"Struct"`
}

func (s GetAddonResponseBodyAddonResourcesSpec) String() string {
	return dara.Prettify(s)
}

func (s GetAddonResponseBodyAddonResourcesSpec) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyAddonResourcesSpec) GetEcsResources() []*AddonNodeTemplate {
	return s.EcsResources
}

func (s *GetAddonResponseBodyAddonResourcesSpec) GetEipResource() *GetAddonResponseBodyAddonResourcesSpecEipResource {
	return s.EipResource
}

func (s *GetAddonResponseBodyAddonResourcesSpec) SetEcsResources(v []*AddonNodeTemplate) *GetAddonResponseBodyAddonResourcesSpec {
	s.EcsResources = v
	return s
}

func (s *GetAddonResponseBodyAddonResourcesSpec) SetEipResource(v *GetAddonResponseBodyAddonResourcesSpecEipResource) *GetAddonResponseBodyAddonResourcesSpec {
	s.EipResource = v
	return s
}

func (s *GetAddonResponseBodyAddonResourcesSpec) Validate() error {
	if s.EcsResources != nil {
		for _, item := range s.EcsResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EipResource != nil {
		if err := s.EipResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAddonResponseBodyAddonResourcesSpecEipResource struct {
	// Indicates whether the EIP is automatically created.
	//
	// example:
	//
	// True
	AutoCreate *bool `json:"AutoCreate,omitempty" xml:"AutoCreate,omitempty"`
	// The maximum bandwidth of the EIP. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The EIP.
	//
	// example:
	//
	// 39.108.xx.xx
	EipAddress *string `json:"EipAddress,omitempty" xml:"EipAddress,omitempty"`
	// The EIP ID.
	//
	// example:
	//
	// eip-bp1vi9124tbx1g3kr****
	EipInstanceId *string `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	// The billing method of the EIP.
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// 	- PrePaid: subscription.
	//
	// Default value: PostPaid.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the EIP. Valid values:
	//
	// 	- PayByBandwidth: pay by bandwidth.
	//
	// 	- PayByTraffic: pay by data transfer.
	//
	// Valid values of N: 1 to 10.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
}

func (s GetAddonResponseBodyAddonResourcesSpecEipResource) String() string {
	return dara.Prettify(s)
}

func (s GetAddonResponseBodyAddonResourcesSpecEipResource) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) GetAutoCreate() *bool {
	return s.AutoCreate
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) GetEipAddress() *string {
	return s.EipAddress
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) GetEipInstanceId() *string {
	return s.EipInstanceId
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) SetAutoCreate(v bool) *GetAddonResponseBodyAddonResourcesSpecEipResource {
	s.AutoCreate = &v
	return s
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) SetBandwidth(v string) *GetAddonResponseBodyAddonResourcesSpecEipResource {
	s.Bandwidth = &v
	return s
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) SetEipAddress(v string) *GetAddonResponseBodyAddonResourcesSpecEipResource {
	s.EipAddress = &v
	return s
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) SetEipInstanceId(v string) *GetAddonResponseBodyAddonResourcesSpecEipResource {
	s.EipInstanceId = &v
	return s
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) SetInstanceChargeType(v string) *GetAddonResponseBodyAddonResourcesSpecEipResource {
	s.InstanceChargeType = &v
	return s
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) SetInternetChargeType(v string) *GetAddonResponseBodyAddonResourcesSpecEipResource {
	s.InternetChargeType = &v
	return s
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) Validate() error {
	return dara.Validate(s)
}

type GetAddonResponseBodyAddonServicesSpec struct {
	// The parameter configurations of the service.
	InputParams []*GetAddonResponseBodyAddonServicesSpecInputParams `json:"InputParams,omitempty" xml:"InputParams,omitempty" type:"Repeated"`
	// The security group configurations of the service.
	NetworkACL []*GetAddonResponseBodyAddonServicesSpecNetworkACL `json:"NetworkACL,omitempty" xml:"NetworkACL,omitempty" type:"Repeated"`
	// The service access type.
	//
	// example:
	//
	// URL
	ServiceAccessType *string `json:"ServiceAccessType,omitempty" xml:"ServiceAccessType,omitempty"`
	// The service access URL.
	//
	// example:
	//
	// https://47.96.xx.xx:12008
	ServiceAccessUrl *string `json:"ServiceAccessUrl,omitempty" xml:"ServiceAccessUrl,omitempty"`
	// The service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Web Portal
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetAddonResponseBodyAddonServicesSpec) String() string {
	return dara.Prettify(s)
}

func (s GetAddonResponseBodyAddonServicesSpec) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyAddonServicesSpec) GetInputParams() []*GetAddonResponseBodyAddonServicesSpecInputParams {
	return s.InputParams
}

func (s *GetAddonResponseBodyAddonServicesSpec) GetNetworkACL() []*GetAddonResponseBodyAddonServicesSpecNetworkACL {
	return s.NetworkACL
}

func (s *GetAddonResponseBodyAddonServicesSpec) GetServiceAccessType() *string {
	return s.ServiceAccessType
}

func (s *GetAddonResponseBodyAddonServicesSpec) GetServiceAccessUrl() *string {
	return s.ServiceAccessUrl
}

func (s *GetAddonResponseBodyAddonServicesSpec) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetAddonResponseBodyAddonServicesSpec) SetInputParams(v []*GetAddonResponseBodyAddonServicesSpecInputParams) *GetAddonResponseBodyAddonServicesSpec {
	s.InputParams = v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpec) SetNetworkACL(v []*GetAddonResponseBodyAddonServicesSpecNetworkACL) *GetAddonResponseBodyAddonServicesSpec {
	s.NetworkACL = v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpec) SetServiceAccessType(v string) *GetAddonResponseBodyAddonServicesSpec {
	s.ServiceAccessType = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpec) SetServiceAccessUrl(v string) *GetAddonResponseBodyAddonServicesSpec {
	s.ServiceAccessUrl = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpec) SetServiceName(v string) *GetAddonResponseBodyAddonServicesSpec {
	s.ServiceName = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpec) Validate() error {
	if s.InputParams != nil {
		for _, item := range s.InputParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NetworkACL != nil {
		for _, item := range s.NetworkACL {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAddonResponseBodyAddonServicesSpecInputParams struct {
	// The help information of the parameter.
	HelpText *string `json:"HelpText,omitempty" xml:"HelpText,omitempty"`
	// The parameter label.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The parameter name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MYSQL_HOME
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter type.
	//
	// This parameter is required.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The parameter value.
	//
	// This parameter is required.
	//
	// example:
	//
	// usr/local/mysql
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAddonResponseBodyAddonServicesSpecInputParams) String() string {
	return dara.Prettify(s)
}

func (s GetAddonResponseBodyAddonServicesSpecInputParams) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyAddonServicesSpecInputParams) GetHelpText() *string {
	return s.HelpText
}

func (s *GetAddonResponseBodyAddonServicesSpecInputParams) GetLabel() *string {
	return s.Label
}

func (s *GetAddonResponseBodyAddonServicesSpecInputParams) GetName() *string {
	return s.Name
}

func (s *GetAddonResponseBodyAddonServicesSpecInputParams) GetType() *string {
	return s.Type
}

func (s *GetAddonResponseBodyAddonServicesSpecInputParams) GetValue() *string {
	return s.Value
}

func (s *GetAddonResponseBodyAddonServicesSpecInputParams) SetHelpText(v string) *GetAddonResponseBodyAddonServicesSpecInputParams {
	s.HelpText = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpecInputParams) SetLabel(v string) *GetAddonResponseBodyAddonServicesSpecInputParams {
	s.Label = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpecInputParams) SetName(v string) *GetAddonResponseBodyAddonServicesSpecInputParams {
	s.Name = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpecInputParams) SetType(v string) *GetAddonResponseBodyAddonServicesSpecInputParams {
	s.Type = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpecInputParams) SetValue(v string) *GetAddonResponseBodyAddonServicesSpecInputParams {
	s.Value = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpecInputParams) Validate() error {
	return dara.Validate(s)
}

type GetAddonResponseBodyAddonServicesSpecNetworkACL struct {
	// The protocol type. Valid values:
	//
	// TCP: forwards TCP packets.
	//
	// UDP: forwards UDP packets.
	//
	// Any: forwards all packets.
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The port number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *float32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The source CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.0.0/12
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s GetAddonResponseBodyAddonServicesSpecNetworkACL) String() string {
	return dara.Prettify(s)
}

func (s GetAddonResponseBodyAddonServicesSpecNetworkACL) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyAddonServicesSpecNetworkACL) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *GetAddonResponseBodyAddonServicesSpecNetworkACL) GetPort() *float32 {
	return s.Port
}

func (s *GetAddonResponseBodyAddonServicesSpecNetworkACL) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *GetAddonResponseBodyAddonServicesSpecNetworkACL) SetIpProtocol(v string) *GetAddonResponseBodyAddonServicesSpecNetworkACL {
	s.IpProtocol = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpecNetworkACL) SetPort(v float32) *GetAddonResponseBodyAddonServicesSpecNetworkACL {
	s.Port = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpecNetworkACL) SetSourceCidrIp(v string) *GetAddonResponseBodyAddonServicesSpecNetworkACL {
	s.SourceCidrIp = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpecNetworkACL) Validate() error {
	return dara.Validate(s)
}
