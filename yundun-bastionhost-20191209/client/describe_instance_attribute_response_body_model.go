// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceAttribute(v *DescribeInstanceAttributeResponseBodyInstanceAttribute) *DescribeInstanceAttributeResponseBody
	GetInstanceAttribute() *DescribeInstanceAttributeResponseBodyInstanceAttribute
	SetRequestId(v string) *DescribeInstanceAttributeResponseBody
	GetRequestId() *string
}

type DescribeInstanceAttributeResponseBody struct {
	// The instance attribute information.
	InstanceAttribute *DescribeInstanceAttributeResponseBodyInstanceAttribute `json:"InstanceAttribute,omitempty" xml:"InstanceAttribute,omitempty" type:"Struct"`
	// The ID of the request. Alibaba Cloud generates a unique identifier for each request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 082FAB35-6AB9-4FD5-8750-D36673548E76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBody) GetInstanceAttribute() *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	return s.InstanceAttribute
}

func (s *DescribeInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceAttributeResponseBody) SetInstanceAttribute(v *DescribeInstanceAttributeResponseBodyInstanceAttribute) *DescribeInstanceAttributeResponseBody {
	s.InstanceAttribute = v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetRequestId(v string) *DescribeInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) Validate() error {
	if s.InstanceAttribute != nil {
		if err := s.InstanceAttribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceAttributeResponseBodyInstanceAttribute struct {
	AiCreditStatus *string `json:"AiCreditStatus,omitempty" xml:"AiCreditStatus,omitempty"`
	AiOpsModule    *string `json:"AiOpsModule,omitempty" xml:"AiOpsModule,omitempty"`
	// The application O&M module. Valid values: Enable (enabled) and Disable (disabled).
	//
	// example:
	//
	// Enable
	AppOperationModule *string `json:"AppOperationModule,omitempty" xml:"AppOperationModule,omitempty"`
	// The list of authorized security group IDs.
	AuthorizedSecurityGroups []*string `json:"AuthorizedSecurityGroups,omitempty" xml:"AuthorizedSecurityGroups,omitempty" type:"Repeated"`
	// The total bandwidth of the bastion host instance.
	//
	// example:
	//
	// 30
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The extended bandwidth package of the bastion host.
	//
	// example:
	//
	// 5
	BandwidthPackage *string `json:"BandwidthPackage,omitempty" xml:"BandwidthPackage,omitempty"`
	// The status of the database O&M feature.
	//
	// - **Enable**: Database O&M is supported.
	//
	// - **Disable**: Database O&M is not supported.
	//
	// example:
	//
	// Disable
	DbOperationModule *string `json:"DbOperationModule,omitempty" xml:"DbOperationModule,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// Test API
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the elastic network interface (ENI). An ENI is a virtual network interface controller (NIC) that can be attached to the bastion host instance.
	//
	// example:
	//
	// eni-bp1455jrzwm7moaxxxxx
	EniInstanceId *string `json:"EniInstanceId,omitempty" xml:"EniInstanceId,omitempty"`
	// The timestamp when the bastion host instance expires. Unit: milliseconds.
	//
	// example:
	//
	// 1578326400000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The status of the HSM hardware encryption module. Indicates whether the bastion host is integrated with HSM.
	HSMModule *string `json:"HSMModule,omitempty" xml:"HSMModule,omitempty"`
	// The IDaaS integration module. Valid values: Enable (enabled) and Disable (disabled).
	//
	// example:
	//
	// Enable
	IDaaSModule *string `json:"IDaaSModule,omitempty" xml:"IDaaSModule,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// bastionhost-cn-78v1ghxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance status. Valid values:
	//
	// - **PENDING**: Not initialized.
	//
	// - **CREATING**: Being created.
	//
	// - **RUNNING**: Running.
	//
	// - **EXPIRED**: Expired.
	//
	// - **CREATE_FAILED**: Creation failed.
	//
	// - **UPGRADING**: Being upgraded.
	//
	// - **UPGRADE_FAILED**: Upgrade failed.
	//
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The public domain name.
	//
	// example:
	//
	// ******lwb-public.bastionhost.aliyuncs.com
	InternetEndpoint *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	// The internal domain name.
	//
	// example:
	//
	// ******xalwb.bastionhost.aliyuncs.com
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	// The KMS Secrets Manager integration module. Valid values: Enable (enabled) and Disable (disabled).
	//
	// example:
	//
	// Enable
	KmsSecretModule *string `json:"KmsSecretModule,omitempty" xml:"KmsSecretModule,omitempty"`
	// The license code.
	//
	// example:
	//
	// bhah_ent_50_asset
	LicenseCode *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	// The status of the password change task feature.
	//
	// - **Enable**: Enabled.
	//
	// - **Disable**: Disabled.
	//
	// example:
	//
	// Enable
	ModifyPasswordModule *string `json:"ModifyPasswordModule,omitempty" xml:"ModifyPasswordModule,omitempty"`
	// The status of the network domain proxy feature.
	//
	// - **Enable**: The network domain proxy mode is supported.
	//
	// - **Disable**: The network domain proxy mode is not supported.
	//
	// example:
	//
	// Enable
	NetworkProxyModule *string `json:"NetworkProxyModule,omitempty" xml:"NetworkProxyModule,omitempty"`
	// The O&M ports of the bastion host.
	Ports []*DescribeInstanceAttributeResponseBodyInstanceAttributePorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The list of internal egress IP addresses of the bastion host.
	PrivateExportIps []*string `json:"PrivateExportIps,omitempty" xml:"PrivateExportIps,omitempty" type:"Repeated"`
	// The list of IP addresses in the internal whitelist.
	PrivateWhiteList []*string `json:"PrivateWhiteList,omitempty" xml:"PrivateWhiteList,omitempty" type:"Repeated"`
	// The list of public egress IP addresses of the bastion host.
	PublicExportIps []*string `json:"PublicExportIps,omitempty" xml:"PublicExportIps,omitempty" type:"Repeated"`
	// The list of public IP addresses of the bastion host.
	PublicIps []*string `json:"PublicIps,omitempty" xml:"PublicIps,omitempty" type:"Repeated"`
	// Indicates whether the bastion host instance is accessible over the Internet. Valid values:
	//
	// - **true**: The bastion host is accessible over the Internet.
	//
	// - **false**: The bastion host is not accessible over the Internet.
	//
	// example:
	//
	// true
	PublicNetworkAccess *bool `json:"PublicNetworkAccess,omitempty" xml:"PublicNetworkAccess,omitempty"`
	// The public whitelist of the bastion host.
	PublicWhiteList []*string `json:"PublicWhiteList,omitempty" xml:"PublicWhiteList,omitempty" type:"Repeated"`
	// The multi-account module. Valid values: Enable (enabled) and Disable (disabled).
	//
	// example:
	//
	// Enable
	RDModule *string `json:"RDModule,omitempty" xml:"RDModule,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aekzc427db******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of rules for the bastion host instance.
	RouterRules []*string `json:"RouterRules,omitempty" xml:"RouterRules,omitempty" type:"Repeated"`
	// The script O&M module. Valid values: Enable (enabled) and Disable (disabled).
	//
	// example:
	//
	// Enable
	ScriptDeliverModule *string `json:"ScriptDeliverModule,omitempty" xml:"ScriptDeliverModule,omitempty"`
	// The list of security group IDs to which the instance belongs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The ID of the secondary vSwitch associated with the bastion host instance.
	//
	// example:
	//
	// vsw-uf6cmnae7hu5****
	SlaveVswitchId *string `json:"SlaveVswitchId,omitempty" xml:"SlaveVswitchId,omitempty"`
	// The timestamp when the bastion host instance was purchased or renewed. Unit: milliseconds.
	//
	// example:
	//
	// 1577681345000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total storage capacity of the purchased bastion host. Unit: bytes.
	//
	// example:
	//
	// 2199023255552
	Storage *int64 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// The VPC ID associated with the instance.
	//
	// example:
	//
	// vpc-bp1c85tzgqu1bf5bxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID associated with the instance.
	//
	// example:
	//
	// vsw-bp1xfwzzfti0kjbfxxxxx
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The status of the Web Terminal feature.
	//
	// - **Enable**: Web remote connection is supported.
	//
	// - **Disable**: Web remote connection is not supported.
	//
	// example:
	//
	// Enable
	WebTerminalModule *string `json:"WebTerminalModule,omitempty" xml:"WebTerminalModule,omitempty"`
	// The IP address whitelist to configure.
	WhiteListPolicies []*DescribeInstanceAttributeResponseBodyInstanceAttributeWhiteListPolicies `json:"WhiteListPolicies,omitempty" xml:"WhiteListPolicies,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetAiCreditStatus() *string {
	return s.AiCreditStatus
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetAiOpsModule() *string {
	return s.AiOpsModule
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetAppOperationModule() *string {
	return s.AppOperationModule
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetAuthorizedSecurityGroups() []*string {
	return s.AuthorizedSecurityGroups
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetBandwidthPackage() *string {
	return s.BandwidthPackage
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetDbOperationModule() *string {
	return s.DbOperationModule
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetEniInstanceId() *string {
	return s.EniInstanceId
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetHSMModule() *string {
	return s.HSMModule
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetIDaaSModule() *string {
	return s.IDaaSModule
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetInternetEndpoint() *string {
	return s.InternetEndpoint
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetIntranetEndpoint() *string {
	return s.IntranetEndpoint
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetKmsSecretModule() *string {
	return s.KmsSecretModule
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetLicenseCode() *string {
	return s.LicenseCode
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetModifyPasswordModule() *string {
	return s.ModifyPasswordModule
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetNetworkProxyModule() *string {
	return s.NetworkProxyModule
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetPorts() []*DescribeInstanceAttributeResponseBodyInstanceAttributePorts {
	return s.Ports
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetPrivateExportIps() []*string {
	return s.PrivateExportIps
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetPrivateWhiteList() []*string {
	return s.PrivateWhiteList
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetPublicExportIps() []*string {
	return s.PublicExportIps
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetPublicIps() []*string {
	return s.PublicIps
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetPublicNetworkAccess() *bool {
	return s.PublicNetworkAccess
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetPublicWhiteList() []*string {
	return s.PublicWhiteList
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetRDModule() *string {
	return s.RDModule
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetRouterRules() []*string {
	return s.RouterRules
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetScriptDeliverModule() *string {
	return s.ScriptDeliverModule
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetSlaveVswitchId() *string {
	return s.SlaveVswitchId
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetStorage() *int64 {
	return s.Storage
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetWebTerminalModule() *string {
	return s.WebTerminalModule
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetWhiteListPolicies() []*DescribeInstanceAttributeResponseBodyInstanceAttributeWhiteListPolicies {
	return s.WhiteListPolicies
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetAiCreditStatus(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.AiCreditStatus = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetAiOpsModule(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.AiOpsModule = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetAppOperationModule(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.AppOperationModule = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetAuthorizedSecurityGroups(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.AuthorizedSecurityGroups = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetBandwidth(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Bandwidth = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetBandwidthPackage(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.BandwidthPackage = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetDbOperationModule(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.DbOperationModule = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetDescription(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Description = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetEniInstanceId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.EniInstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetExpireTime(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetHSMModule(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.HSMModule = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetIDaaSModule(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.IDaaSModule = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceStatus(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInternetEndpoint(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InternetEndpoint = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetIntranetEndpoint(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.IntranetEndpoint = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetKmsSecretModule(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.KmsSecretModule = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetLicenseCode(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.LicenseCode = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetModifyPasswordModule(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ModifyPasswordModule = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetNetworkProxyModule(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.NetworkProxyModule = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPorts(v []*DescribeInstanceAttributeResponseBodyInstanceAttributePorts) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Ports = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPrivateExportIps(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PrivateExportIps = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPrivateWhiteList(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PrivateWhiteList = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPublicExportIps(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PublicExportIps = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPublicIps(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PublicIps = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPublicNetworkAccess(v bool) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PublicNetworkAccess = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPublicWhiteList(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PublicWhiteList = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetRDModule(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.RDModule = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetRegionId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetResourceGroupId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetRouterRules(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.RouterRules = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetScriptDeliverModule(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ScriptDeliverModule = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetSecurityGroupIds(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetSlaveVswitchId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.SlaveVswitchId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetStartTime(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetStorage(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Storage = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetVpcId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetVswitchId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetWebTerminalModule(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.WebTerminalModule = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetWhiteListPolicies(v []*DescribeInstanceAttributeResponseBodyInstanceAttributeWhiteListPolicies) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.WhiteListPolicies = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) Validate() error {
	if s.Ports != nil {
		for _, item := range s.Ports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WhiteListPolicies != nil {
		for _, item := range s.WhiteListPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceAttributeResponseBodyInstanceAttributePorts struct {
	// The custom port defined by the user.
	//
	// > Only SSH and RDP ports can be modified. If no custom O&M port is configured for the bastion host, the value is the same as the standard port.
	//
	// example:
	//
	// 600xx
	CustomPort *int32 `json:"CustomPort,omitempty" xml:"CustomPort,omitempty"`
	// The standard port of the bastion host. Valid values:
	//
	// - **SSH**: 60022
	//
	// - **RDP**: 63389
	//
	// - **HTTPS**: 443
	//
	// example:
	//
	// 60022
	StandardPort *int32 `json:"StandardPort,omitempty" xml:"StandardPort,omitempty"`
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttributePorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttributePorts) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributePorts) GetCustomPort() *int32 {
	return s.CustomPort
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributePorts) GetStandardPort() *int32 {
	return s.StandardPort
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributePorts) SetCustomPort(v int32) *DescribeInstanceAttributeResponseBodyInstanceAttributePorts {
	s.CustomPort = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributePorts) SetStandardPort(v int32) *DescribeInstanceAttributeResponseBodyInstanceAttributePorts {
	s.StandardPort = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributePorts) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAttributeResponseBodyInstanceAttributeWhiteListPolicies struct {
	// The description of the whitelist rule.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IP address whitelist to configure. A maximum of 50 IP addresses are supported. Separate multiple IP addresses with commas (,).
	//
	// example:
	//
	// 94.74.xx.xx/32
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttributeWhiteListPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttributeWhiteListPolicies) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributeWhiteListPolicies) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributeWhiteListPolicies) GetEntry() *string {
	return s.Entry
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributeWhiteListPolicies) SetDescription(v string) *DescribeInstanceAttributeResponseBodyInstanceAttributeWhiteListPolicies {
	s.Description = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributeWhiteListPolicies) SetEntry(v string) *DescribeInstanceAttributeResponseBodyInstanceAttributeWhiteListPolicies {
	s.Entry = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributeWhiteListPolicies) Validate() error {
	return dara.Validate(s)
}
