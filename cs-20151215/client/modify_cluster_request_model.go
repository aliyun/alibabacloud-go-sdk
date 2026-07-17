// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessControlList(v []*string) *ModifyClusterRequest
	GetAccessControlList() []*string
	SetApiServerCustomCertSans(v *ModifyClusterRequestApiServerCustomCertSans) *ModifyClusterRequest
	GetApiServerCustomCertSans() *ModifyClusterRequestApiServerCustomCertSans
	SetApiServerEip(v bool) *ModifyClusterRequest
	GetApiServerEip() *bool
	SetApiServerEipId(v string) *ModifyClusterRequest
	GetApiServerEipId() *string
	SetClientToken(v string) *ModifyClusterRequest
	GetClientToken() *string
	SetClusterName(v string) *ModifyClusterRequest
	GetClusterName() *string
	SetClusterSpec(v string) *ModifyClusterRequest
	GetClusterSpec() *string
	SetControlPlaneConfig(v *ModifyClusterRequestControlPlaneConfig) *ModifyClusterRequest
	GetControlPlaneConfig() *ModifyClusterRequestControlPlaneConfig
	SetControlPlaneEndpointsConfig(v *ModifyClusterRequestControlPlaneEndpointsConfig) *ModifyClusterRequest
	GetControlPlaneEndpointsConfig() *ModifyClusterRequestControlPlaneEndpointsConfig
	SetDeletionProtection(v bool) *ModifyClusterRequest
	GetDeletionProtection() *bool
	SetEnableRrsa(v bool) *ModifyClusterRequest
	GetEnableRrsa() *bool
	SetIngressDomainRebinding(v bool) *ModifyClusterRequest
	GetIngressDomainRebinding() *bool
	SetIngressLoadbalancerId(v string) *ModifyClusterRequest
	GetIngressLoadbalancerId() *string
	SetInstanceDeletionProtection(v bool) *ModifyClusterRequest
	GetInstanceDeletionProtection() *bool
	SetMaintenanceWindow(v *MaintenanceWindow) *ModifyClusterRequest
	GetMaintenanceWindow() *MaintenanceWindow
	SetOperationPolicy(v *ModifyClusterRequestOperationPolicy) *ModifyClusterRequest
	GetOperationPolicy() *ModifyClusterRequestOperationPolicy
	SetResourceGroupId(v string) *ModifyClusterRequest
	GetResourceGroupId() *string
	SetSecurityGroupId(v string) *ModifyClusterRequest
	GetSecurityGroupId() *string
	SetSystemEventsLogging(v *ModifyClusterRequestSystemEventsLogging) *ModifyClusterRequest
	GetSystemEventsLogging() *ModifyClusterRequestSystemEventsLogging
	SetTimezone(v string) *ModifyClusterRequest
	GetTimezone() *string
	SetVswitchIds(v []*string) *ModifyClusterRequest
	GetVswitchIds() []*string
}

type ModifyClusterRequest struct {
	// Deprecated
	//
	// The access control list for the registered cluster API Server SLB.
	AccessControlList []*string `json:"access_control_list,omitempty" xml:"access_control_list,omitempty" type:"Repeated"`
	// The custom API Server certificate Subject Alternative Name (SAN). This parameter adds custom IP addresses or domain names to the SAN field of the cluster API Server certificate for client access control.
	//
	// Only managed clusters support this parameter.
	ApiServerCustomCertSans *ModifyClusterRequestApiServerCustomCertSans `json:"api_server_custom_cert_sans,omitempty" xml:"api_server_custom_cert_sans,omitempty" type:"Struct"`
	// Indicates whether an Elastic IP Address (EIP) is attached to the cluster for public network access to the API server. Valid values:
	//
	// - `true`: An EIP is attached to the cluster.
	//
	// - `false`: No EIP is attached to the cluster.
	//
	// example:
	//
	// true
	ApiServerEip *bool `json:"api_server_eip,omitempty" xml:"api_server_eip,omitempty"`
	// The instance ID of the EIP attached to the cluster API Server. This parameter takes effect only when `api_server_eip` is set to `true`.
	//
	// example:
	//
	// eip-wz9fnasl6dsfhmvci****
	ApiServerEipId *string `json:"api_server_eip_id,omitempty" xml:"api_server_eip_id,omitempty"`
	// The client token.
	//
	// example:
	//
	// af31042c-6355-495b-b6e3-exxb9669
	ClientToken *string `json:"client_token,omitempty" xml:"client_token,omitempty"`
	// The custom cluster name. The name can contain digits, Chinese characters, letters, and hyphens (-). It must be 1 to 63 characters in length and cannot start with a hyphen (-).
	//
	// example:
	//
	// cluster-new-name
	ClusterName *string `json:"cluster_name,omitempty" xml:"cluster_name,omitempty"`
	// The cluster specification when `cluster_type` is set to `ManagedKubernetes` and `profile` is configured. Valid values:
	//
	// - `ack.pro.small`: Pro
	//
	// - `ack.pro.xlarge`: Pro XL
	//
	// - `ack.pro.2xlarge`: Pro 2XL
	//
	// - `ack.pro.4xlarge`: Pro 4XL (contact customer service to be added to the whitelist)
	//
	// Pro XL, Pro 2XL, and Pro 4XL are three tiers provided by <props="china">[ACK Pro Provisioned Control Plane](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane)<props="intl">[ACK Pro Provisioned Control Plane](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane). By pre-allocating and dedicating control plane resources, these tiers ensure that API concurrency and pod scheduling capabilities remain at a consistently high level. They are suitable for AI training and inference, ultra-large-scale clusters, and mission-critical workloads.
	//
	// For information about the cluster management fees for Pro and provisioned control plane editions, see <props="china">[Cluster management fees](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee)<props="intl">[Cluster management fees](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee).
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// The dedicated cluster control plane configuration.
	ControlPlaneConfig *ModifyClusterRequestControlPlaneConfig `json:"control_plane_config,omitempty" xml:"control_plane_config,omitempty" type:"Struct"`
	// The cluster connection configuration.
	ControlPlaneEndpointsConfig *ModifyClusterRequestControlPlaneEndpointsConfig `json:"control_plane_endpoints_config,omitempty" xml:"control_plane_endpoints_config,omitempty" type:"Struct"`
	// Specifies whether to enable deletion protection for the cluster to prevent accidental deletion through the console or API. Valid values:
	//
	// - `true`: Enables cluster deletion protection. The cluster cannot be deleted through the console or API.
	//
	// - `false`: Disables cluster deletion protection. The cluster can be deleted through the console or API.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// Specifies whether to enable or disable the RAM Roles for Service Accounts (RRSA) feature. Only managed clusters support this parameter. Valid values:
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
	//
	// example:
	//
	// true
	EnableRrsa *bool `json:"enable_rrsa,omitempty" xml:"enable_rrsa,omitempty"`
	// Deprecated
	//
	// Specifies whether to rebind the cluster test domain name. Valid values:
	//
	// - `true`: Rebinds the cluster test domain name.
	//
	// - `false`: Does not rebind the cluster test domain name.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	IngressDomainRebinding *bool `json:"ingress_domain_rebinding,omitempty" xml:"ingress_domain_rebinding,omitempty"`
	// Deprecated
	//
	// The SLB instance ID of the cluster to be modified.
	//
	// example:
	//
	// lb-wz97kes8tnndkpodw****
	IngressLoadbalancerId *string `json:"ingress_loadbalancer_id,omitempty" xml:"ingress_loadbalancer_id,omitempty"`
	// Deprecated
	//
	// Specifies whether to enable instance deletion protection to prevent accidental deletion of nodes through the console or API. Valid values:
	//
	// - `true`: Nodes cannot be accidentally deleted through the console or API.
	//
	// - `false`: Nodes can be accidentally deleted through the console or API.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	InstanceDeletionProtection *bool `json:"instance_deletion_protection,omitempty" xml:"instance_deletion_protection,omitempty"`
	// The maintenance window configuration for the cluster. This feature takes effect only for ACK Pro clusters.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window,omitempty" xml:"maintenance_window,omitempty"`
	// The cluster automatic O&M policy.
	OperationPolicy *ModifyClusterRequestOperationPolicy `json:"operation_policy,omitempty" xml:"operation_policy,omitempty" type:"Struct"`
	// The resource group ID of the cluster.
	//
	// example:
	//
	// rg-acfmyvw3wjm****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// The control plane security group ID.
	//
	// - If you have configured blocking rules in the security group, ensure that the security group rules allow the protocols and ports required by the cluster. For information about recommended security group rules, see [Configure and manage cluster security groups](https://help.aliyun.com/document_detail/353191.html).
	//
	// - For non-ACK dedicated clusters, the cluster control plane and installed managed components (such as terway-controlplane) briefly restart during the procedure. Perform this operation during off-peak hours. After the control plane security group is changed, the network interface controllers (NICs) used by the cluster control plane and installed managed components are automatically added to the new security group.
	//
	// - For ACK dedicated clusters, after the control plane security group is changed, newly scaled-out master nodes automatically use the new control plane security group. Existing control plane nodes are not affected.
	//
	// example:
	//
	// sg-bp1h6rk3pgct2a08***
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// The system event storage configuration.
	SystemEventsLogging *ModifyClusterRequestSystemEventsLogging `json:"system_events_logging,omitempty" xml:"system_events_logging,omitempty" type:"Struct"`
	// The cluster time zone. See [Supported time zones](https://help.aliyun.com/document_detail/354879.html).
	//
	// - After the time zone is changed, the cluster inspection configuration uses the new time zone settings.
	//
	// - For managed clusters, the cluster control plane and installed managed components (such as terway-controlplane) briefly restart during the change. Perform this operation during off-peak hours. After the time zone is changed, newly scaled-out nodes automatically use the new time zone settings. Existing nodes are not affected. You can use the node pool node reset feature to apply the new settings to existing nodes.
	//
	// - For dedicated clusters, after the time zone is changed, newly scaled-out nodes (including control plane nodes) automatically use the new time zone settings. Existing nodes (including control plane nodes) are not affected. You can use the node pool node reset feature to apply the new settings to existing nodes. For control plane nodes, scale out and then scale in to apply the new settings to all control plane nodes.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The vSwitches for the cluster control plane. For dedicated clusters, the change applies to newly scaled-out control plane nodes. When changing control plane vSwitches for managed clusters, note the following:
	//
	// - This parameter performs a full overwrite. Specify the complete list of target vSwitches.
	//
	// - Control plane components briefly restart during the change. Proceed with caution.
	//
	// - Ensure that all security groups of the cluster (including the control plane security group, security groups of all node pools, and security groups used by the container network) allow inbound and outbound traffic for the IP CIDR blocks of the new vSwitches to prevent nodes and containers from losing connectivity to the API server.
	//
	// - If the new control plane vSwitches have ACL rules configured, ensure that the ACL rules allow communication with the CIDR blocks of cluster nodes and the container network.
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
}

func (s ModifyClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterRequest) GetAccessControlList() []*string {
	return s.AccessControlList
}

func (s *ModifyClusterRequest) GetApiServerCustomCertSans() *ModifyClusterRequestApiServerCustomCertSans {
	return s.ApiServerCustomCertSans
}

func (s *ModifyClusterRequest) GetApiServerEip() *bool {
	return s.ApiServerEip
}

func (s *ModifyClusterRequest) GetApiServerEipId() *string {
	return s.ApiServerEipId
}

func (s *ModifyClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ModifyClusterRequest) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *ModifyClusterRequest) GetControlPlaneConfig() *ModifyClusterRequestControlPlaneConfig {
	return s.ControlPlaneConfig
}

func (s *ModifyClusterRequest) GetControlPlaneEndpointsConfig() *ModifyClusterRequestControlPlaneEndpointsConfig {
	return s.ControlPlaneEndpointsConfig
}

func (s *ModifyClusterRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ModifyClusterRequest) GetEnableRrsa() *bool {
	return s.EnableRrsa
}

func (s *ModifyClusterRequest) GetIngressDomainRebinding() *bool {
	return s.IngressDomainRebinding
}

func (s *ModifyClusterRequest) GetIngressLoadbalancerId() *string {
	return s.IngressLoadbalancerId
}

func (s *ModifyClusterRequest) GetInstanceDeletionProtection() *bool {
	return s.InstanceDeletionProtection
}

func (s *ModifyClusterRequest) GetMaintenanceWindow() *MaintenanceWindow {
	return s.MaintenanceWindow
}

func (s *ModifyClusterRequest) GetOperationPolicy() *ModifyClusterRequestOperationPolicy {
	return s.OperationPolicy
}

func (s *ModifyClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyClusterRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ModifyClusterRequest) GetSystemEventsLogging() *ModifyClusterRequestSystemEventsLogging {
	return s.SystemEventsLogging
}

func (s *ModifyClusterRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *ModifyClusterRequest) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *ModifyClusterRequest) SetAccessControlList(v []*string) *ModifyClusterRequest {
	s.AccessControlList = v
	return s
}

func (s *ModifyClusterRequest) SetApiServerCustomCertSans(v *ModifyClusterRequestApiServerCustomCertSans) *ModifyClusterRequest {
	s.ApiServerCustomCertSans = v
	return s
}

func (s *ModifyClusterRequest) SetApiServerEip(v bool) *ModifyClusterRequest {
	s.ApiServerEip = &v
	return s
}

func (s *ModifyClusterRequest) SetApiServerEipId(v string) *ModifyClusterRequest {
	s.ApiServerEipId = &v
	return s
}

func (s *ModifyClusterRequest) SetClientToken(v string) *ModifyClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyClusterRequest) SetClusterName(v string) *ModifyClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *ModifyClusterRequest) SetClusterSpec(v string) *ModifyClusterRequest {
	s.ClusterSpec = &v
	return s
}

func (s *ModifyClusterRequest) SetControlPlaneConfig(v *ModifyClusterRequestControlPlaneConfig) *ModifyClusterRequest {
	s.ControlPlaneConfig = v
	return s
}

func (s *ModifyClusterRequest) SetControlPlaneEndpointsConfig(v *ModifyClusterRequestControlPlaneEndpointsConfig) *ModifyClusterRequest {
	s.ControlPlaneEndpointsConfig = v
	return s
}

func (s *ModifyClusterRequest) SetDeletionProtection(v bool) *ModifyClusterRequest {
	s.DeletionProtection = &v
	return s
}

func (s *ModifyClusterRequest) SetEnableRrsa(v bool) *ModifyClusterRequest {
	s.EnableRrsa = &v
	return s
}

func (s *ModifyClusterRequest) SetIngressDomainRebinding(v bool) *ModifyClusterRequest {
	s.IngressDomainRebinding = &v
	return s
}

func (s *ModifyClusterRequest) SetIngressLoadbalancerId(v string) *ModifyClusterRequest {
	s.IngressLoadbalancerId = &v
	return s
}

func (s *ModifyClusterRequest) SetInstanceDeletionProtection(v bool) *ModifyClusterRequest {
	s.InstanceDeletionProtection = &v
	return s
}

func (s *ModifyClusterRequest) SetMaintenanceWindow(v *MaintenanceWindow) *ModifyClusterRequest {
	s.MaintenanceWindow = v
	return s
}

func (s *ModifyClusterRequest) SetOperationPolicy(v *ModifyClusterRequestOperationPolicy) *ModifyClusterRequest {
	s.OperationPolicy = v
	return s
}

func (s *ModifyClusterRequest) SetResourceGroupId(v string) *ModifyClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyClusterRequest) SetSecurityGroupId(v string) *ModifyClusterRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifyClusterRequest) SetSystemEventsLogging(v *ModifyClusterRequestSystemEventsLogging) *ModifyClusterRequest {
	s.SystemEventsLogging = v
	return s
}

func (s *ModifyClusterRequest) SetTimezone(v string) *ModifyClusterRequest {
	s.Timezone = &v
	return s
}

func (s *ModifyClusterRequest) SetVswitchIds(v []*string) *ModifyClusterRequest {
	s.VswitchIds = v
	return s
}

func (s *ModifyClusterRequest) Validate() error {
	if s.ApiServerCustomCertSans != nil {
		if err := s.ApiServerCustomCertSans.Validate(); err != nil {
			return err
		}
	}
	if s.ControlPlaneConfig != nil {
		if err := s.ControlPlaneConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ControlPlaneEndpointsConfig != nil {
		if err := s.ControlPlaneEndpointsConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MaintenanceWindow != nil {
		if err := s.MaintenanceWindow.Validate(); err != nil {
			return err
		}
	}
	if s.OperationPolicy != nil {
		if err := s.OperationPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.SystemEventsLogging != nil {
		if err := s.SystemEventsLogging.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyClusterRequestApiServerCustomCertSans struct {
	// Specifies whether to overwrite or append the SAN configuration. Valid values:
	//
	// - overwrite: overwrites the existing configuration.
	//
	// - append: appends to the existing configuration.
	//
	// example:
	//
	// append
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The SAN list.
	SubjectAlternativeNames []*string `json:"subject_alternative_names,omitempty" xml:"subject_alternative_names,omitempty" type:"Repeated"`
}

func (s ModifyClusterRequestApiServerCustomCertSans) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterRequestApiServerCustomCertSans) GoString() string {
	return s.String()
}

func (s *ModifyClusterRequestApiServerCustomCertSans) GetAction() *string {
	return s.Action
}

func (s *ModifyClusterRequestApiServerCustomCertSans) GetSubjectAlternativeNames() []*string {
	return s.SubjectAlternativeNames
}

func (s *ModifyClusterRequestApiServerCustomCertSans) SetAction(v string) *ModifyClusterRequestApiServerCustomCertSans {
	s.Action = &v
	return s
}

func (s *ModifyClusterRequestApiServerCustomCertSans) SetSubjectAlternativeNames(v []*string) *ModifyClusterRequestApiServerCustomCertSans {
	s.SubjectAlternativeNames = v
	return s
}

func (s *ModifyClusterRequestApiServerCustomCertSans) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterRequestControlPlaneConfig struct {
	// Specifies whether to enable auto-renewal for control plane node instances. This parameter takes effect only when `charge_type` is set to `PrePaid`. Valid values:
	//
	// - `true`: Enables auto-renewal.
	//
	// - `false`: Disables auto-renewal.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// The auto-renewal duration for each renewal cycle of control plane node instances.
	//
	// Valid values: {1, 2, 3, 6, 12}. Unit: months.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// The billing method for control plane node instances. Valid values:
	//
	// - `PrePaid`: subscription.
	//
	// - `PostPaid`: pay-as-you-go.
	//
	// Default value: `PostPaid`.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"charge_type,omitempty" xml:"charge_type,omitempty"`
	// Specifies whether to install the CloudMonitor agent on control plane nodes. Valid values:
	//
	// - `true`: Installs the CloudMonitor agent.
	//
	// - `false`: Does not install the CloudMonitor agent.
	//
	// example:
	//
	// true
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// The node CPU management policy. The following policies are supported for clusters running version 1.12.6 or later:
	//
	// - `static`: Allows pods with certain resource characteristics on the node to be granted enhanced CPU affinity and exclusivity.
	//
	// - `none`: Uses the existing default CPU affinity scheme.
	//
	// Default value: `none`.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// The deployment set ID.
	//
	// example:
	//
	// ds-bp10b35imuam5amw****
	DeploymentsetId *string `json:"deploymentset_id,omitempty" xml:"deploymentset_id,omitempty"`
	// The custom image ID. Specify this parameter when using a custom image.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20240819.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The operating system image type. Valid values:
	//
	// - `AliyunLinux3`: Alinux3 image.
	//
	// - `Custom`: custom image.
	//
	// example:
	//
	// AliyunLinux3
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// The instance types. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// The key pair name. Mutually exclusive with `login_password`.
	//
	// example:
	//
	// ack
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// The SSH logon password. Mutually exclusive with `key_pair`. The password must be 8 to 30 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. To use password-based logon, specify this parameter during scale-out.
	//
	// example:
	//
	// Ack@2000.
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The node service port range. Available port range: [30000, 65535].
	//
	// Default value: 30000-32767.
	//
	// example:
	//
	// 30000-32767
	NodePortRange *string `json:"node_port_range,omitempty" xml:"node_port_range,omitempty"`
	// The subscription duration for control plane node instances. This parameter takes effect and is required only when `charge_type` is set to `PrePaid`.
	//
	// When `period_unit=Month`, valid values are {1, 2, 3, 6, 12, 24, 36, 48, 60}.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The billing cycle unit for control plane node instances. This parameter takes effect only when `charge_type` is set to `PrePaid`.
	//
	// `Month`: The billing cycle is measured in months. Currently, only months are supported.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// The container runtime name. Valid values:
	//
	// - `containerd`: Recommended. Supported by all cluster versions.
	//
	// Default value: containerd.
	//
	// example:
	//
	// containerd
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// Specifies whether to enable Alibaba Cloud OS security hardening. Valid values:
	//
	// - `true`: Enables Alibaba Cloud OS security hardening.
	//
	// - `false`: Disables Alibaba Cloud OS security hardening.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SecurityHardeningOs *bool `json:"security_hardening_os,omitempty" xml:"security_hardening_os,omitempty"`
	// The number of control plane nodes. To scale out the control plane of a dedicated cluster, set this parameter to the target number of control plane nodes, which must be greater than the current number.
	//
	// example:
	//
	// 5
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// Specifies whether to enable MLPS 2.0 security hardening. For more information, see [ACK MLPS 2.0 security hardening](https://help.aliyun.com/document_detail/196148.html).
	//
	// Valid values:
	//
	// - `true`: Enables MLPS 2.0 security hardening.
	//
	// - `false`: Disables MLPS 2.0 security hardening.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// Specifies whether to enable burst (performance burst) for the node system cloud disk. Valid values:
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
	//
	// This parameter is supported only when `system_disk_category` is set to `cloud_auto`. For more information, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// The node system cloud disk type. Valid values:
	//
	// - `cloud_efficiency`: ultra cloud disk.
	//
	// - `cloud_ssd`: standard SSD.
	//
	// - `cloud_essd`: Enterprise SSD (ESSD).
	//
	// - `cloud_auto`: ESSD AutoPL cloud disk.
	//
	// - `cloud_essd_entry`: ESSD Entry cloud disk.
	//
	// example:
	//
	// cloud_essd
	SystemDiskCategory *string `json:"system_disk_category,omitempty" xml:"system_disk_category,omitempty"`
	// The performance level of the node system cloud disk. This parameter takes effect only for ESSD cloud disks. The performance level varies based on the cloud disk size. For more information, see [standard SSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// The provisioned read/write IOPS for the node system cloud disk. Valid values: 0 to min{50,000, 1000 × capacity - baseline performance}. Baseline performance = min{1,800 + 50 × capacity, 50000}.
	//
	// This parameter is supported only when `system_disk_category` is set to `cloud_auto`. For more information, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// The node system cloud disk size. Valid values: [40, 500]. Unit: GiB.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// The ID of the automatic snapshot policy applied to the node system cloud disk.
	//
	// example:
	//
	// sp-2zej1nogjvovnz4z****
	SystemDiskSnapshotPolicyId *string `json:"system_disk_snapshot_policy_id,omitempty" xml:"system_disk_snapshot_policy_id,omitempty"`
}

func (s ModifyClusterRequestControlPlaneConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterRequestControlPlaneConfig) GoString() string {
	return s.String()
}

func (s *ModifyClusterRequestControlPlaneConfig) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ModifyClusterRequestControlPlaneConfig) GetAutoRenewPeriod() *int64 {
	return s.AutoRenewPeriod
}

func (s *ModifyClusterRequestControlPlaneConfig) GetChargeType() *string {
	return s.ChargeType
}

func (s *ModifyClusterRequestControlPlaneConfig) GetCloudMonitorFlags() *bool {
	return s.CloudMonitorFlags
}

func (s *ModifyClusterRequestControlPlaneConfig) GetCpuPolicy() *string {
	return s.CpuPolicy
}

func (s *ModifyClusterRequestControlPlaneConfig) GetDeploymentsetId() *string {
	return s.DeploymentsetId
}

func (s *ModifyClusterRequestControlPlaneConfig) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyClusterRequestControlPlaneConfig) GetImageType() *string {
	return s.ImageType
}

func (s *ModifyClusterRequestControlPlaneConfig) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *ModifyClusterRequestControlPlaneConfig) GetKeyPair() *string {
	return s.KeyPair
}

func (s *ModifyClusterRequestControlPlaneConfig) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *ModifyClusterRequestControlPlaneConfig) GetNodePortRange() *string {
	return s.NodePortRange
}

func (s *ModifyClusterRequestControlPlaneConfig) GetPeriod() *int64 {
	return s.Period
}

func (s *ModifyClusterRequestControlPlaneConfig) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyClusterRequestControlPlaneConfig) GetRuntime() *string {
	return s.Runtime
}

func (s *ModifyClusterRequestControlPlaneConfig) GetSecurityHardeningOs() *bool {
	return s.SecurityHardeningOs
}

func (s *ModifyClusterRequestControlPlaneConfig) GetSize() *int64 {
	return s.Size
}

func (s *ModifyClusterRequestControlPlaneConfig) GetSocEnabled() *bool {
	return s.SocEnabled
}

func (s *ModifyClusterRequestControlPlaneConfig) GetSystemDiskBurstingEnabled() *bool {
	return s.SystemDiskBurstingEnabled
}

func (s *ModifyClusterRequestControlPlaneConfig) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *ModifyClusterRequestControlPlaneConfig) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *ModifyClusterRequestControlPlaneConfig) GetSystemDiskProvisionedIops() *int64 {
	return s.SystemDiskProvisionedIops
}

func (s *ModifyClusterRequestControlPlaneConfig) GetSystemDiskSize() *int64 {
	return s.SystemDiskSize
}

func (s *ModifyClusterRequestControlPlaneConfig) GetSystemDiskSnapshotPolicyId() *string {
	return s.SystemDiskSnapshotPolicyId
}

func (s *ModifyClusterRequestControlPlaneConfig) SetAutoRenew(v bool) *ModifyClusterRequestControlPlaneConfig {
	s.AutoRenew = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetAutoRenewPeriod(v int64) *ModifyClusterRequestControlPlaneConfig {
	s.AutoRenewPeriod = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetChargeType(v string) *ModifyClusterRequestControlPlaneConfig {
	s.ChargeType = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetCloudMonitorFlags(v bool) *ModifyClusterRequestControlPlaneConfig {
	s.CloudMonitorFlags = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetCpuPolicy(v string) *ModifyClusterRequestControlPlaneConfig {
	s.CpuPolicy = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetDeploymentsetId(v string) *ModifyClusterRequestControlPlaneConfig {
	s.DeploymentsetId = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetImageId(v string) *ModifyClusterRequestControlPlaneConfig {
	s.ImageId = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetImageType(v string) *ModifyClusterRequestControlPlaneConfig {
	s.ImageType = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetInstanceTypes(v []*string) *ModifyClusterRequestControlPlaneConfig {
	s.InstanceTypes = v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetKeyPair(v string) *ModifyClusterRequestControlPlaneConfig {
	s.KeyPair = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetLoginPassword(v string) *ModifyClusterRequestControlPlaneConfig {
	s.LoginPassword = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetNodePortRange(v string) *ModifyClusterRequestControlPlaneConfig {
	s.NodePortRange = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetPeriod(v int64) *ModifyClusterRequestControlPlaneConfig {
	s.Period = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetPeriodUnit(v string) *ModifyClusterRequestControlPlaneConfig {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetRuntime(v string) *ModifyClusterRequestControlPlaneConfig {
	s.Runtime = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetSecurityHardeningOs(v bool) *ModifyClusterRequestControlPlaneConfig {
	s.SecurityHardeningOs = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetSize(v int64) *ModifyClusterRequestControlPlaneConfig {
	s.Size = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetSocEnabled(v bool) *ModifyClusterRequestControlPlaneConfig {
	s.SocEnabled = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetSystemDiskBurstingEnabled(v bool) *ModifyClusterRequestControlPlaneConfig {
	s.SystemDiskBurstingEnabled = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetSystemDiskCategory(v string) *ModifyClusterRequestControlPlaneConfig {
	s.SystemDiskCategory = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetSystemDiskPerformanceLevel(v string) *ModifyClusterRequestControlPlaneConfig {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetSystemDiskProvisionedIops(v int64) *ModifyClusterRequestControlPlaneConfig {
	s.SystemDiskProvisionedIops = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetSystemDiskSize(v int64) *ModifyClusterRequestControlPlaneConfig {
	s.SystemDiskSize = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) SetSystemDiskSnapshotPolicyId(v string) *ModifyClusterRequestControlPlaneConfig {
	s.SystemDiskSnapshotPolicyId = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterRequestControlPlaneEndpointsConfig struct {
	// The internal DNS configuration for the cluster. Applicable to ACK managed clusters. The internal domain name is used by node-side system components such as kubelet and kube-proxy to access the API Server. When internal domain name access is not enabled, node-side system components access the API Server through the CLB IP address.
	InternalDnsConfig *ModifyClusterRequestControlPlaneEndpointsConfigInternalDnsConfig `json:"internal_dns_config,omitempty" xml:"internal_dns_config,omitempty" type:"Struct"`
	// The cluster access load balancing configuration.
	LoadBalancersConfig []*ModifyClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig `json:"load_balancers_config,omitempty" xml:"load_balancers_config,omitempty" type:"Repeated"`
}

func (s ModifyClusterRequestControlPlaneEndpointsConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterRequestControlPlaneEndpointsConfig) GoString() string {
	return s.String()
}

func (s *ModifyClusterRequestControlPlaneEndpointsConfig) GetInternalDnsConfig() *ModifyClusterRequestControlPlaneEndpointsConfigInternalDnsConfig {
	return s.InternalDnsConfig
}

func (s *ModifyClusterRequestControlPlaneEndpointsConfig) GetLoadBalancersConfig() []*ModifyClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig {
	return s.LoadBalancersConfig
}

func (s *ModifyClusterRequestControlPlaneEndpointsConfig) SetInternalDnsConfig(v *ModifyClusterRequestControlPlaneEndpointsConfigInternalDnsConfig) *ModifyClusterRequestControlPlaneEndpointsConfig {
	s.InternalDnsConfig = v
	return s
}

func (s *ModifyClusterRequestControlPlaneEndpointsConfig) SetLoadBalancersConfig(v []*ModifyClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig) *ModifyClusterRequestControlPlaneEndpointsConfig {
	s.LoadBalancersConfig = v
	return s
}

func (s *ModifyClusterRequestControlPlaneEndpointsConfig) Validate() error {
	if s.InternalDnsConfig != nil {
		if err := s.InternalDnsConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LoadBalancersConfig != nil {
		for _, item := range s.LoadBalancersConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyClusterRequestControlPlaneEndpointsConfigInternalDnsConfig struct {
	// The VPCs where the internal domain name resolution takes effect.
	BindVpcs []*string `json:"bind_vpcs,omitempty" xml:"bind_vpcs,omitempty" type:"Repeated"`
	// Specifies whether to enable internal domain name access for the cluster. Valid values:
	//
	// - true: Enables internal domain name access. Node-side components (kubelet, kube-proxy) access the API Server through the internal domain name.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s ModifyClusterRequestControlPlaneEndpointsConfigInternalDnsConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterRequestControlPlaneEndpointsConfigInternalDnsConfig) GoString() string {
	return s.String()
}

func (s *ModifyClusterRequestControlPlaneEndpointsConfigInternalDnsConfig) GetBindVpcs() []*string {
	return s.BindVpcs
}

func (s *ModifyClusterRequestControlPlaneEndpointsConfigInternalDnsConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyClusterRequestControlPlaneEndpointsConfigInternalDnsConfig) SetBindVpcs(v []*string) *ModifyClusterRequestControlPlaneEndpointsConfigInternalDnsConfig {
	s.BindVpcs = v
	return s
}

func (s *ModifyClusterRequestControlPlaneEndpointsConfigInternalDnsConfig) SetEnabled(v bool) *ModifyClusterRequestControlPlaneEndpointsConfigInternalDnsConfig {
	s.Enabled = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneEndpointsConfigInternalDnsConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig struct {
	// The endpoint type.
	//
	// example:
	//
	// public
	EndpointType *string `json:"endpoint_type,omitempty" xml:"endpoint_type,omitempty"`
	// The SLB instance associated with the endpoint.
	//
	// example:
	//
	// nlb-xxxx
	LoadBalancerId *string `json:"load_balancer_id,omitempty" xml:"load_balancer_id,omitempty"`
}

func (s ModifyClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig) GoString() string {
	return s.String()
}

func (s *ModifyClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig) GetEndpointType() *string {
	return s.EndpointType
}

func (s *ModifyClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ModifyClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig) SetEndpointType(v string) *ModifyClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig {
	s.EndpointType = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig) SetLoadBalancerId(v string) *ModifyClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig {
	s.LoadBalancerId = &v
	return s
}

func (s *ModifyClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterRequestOperationPolicy struct {
	// The cluster auto upgrade configuration.
	ClusterAutoUpgrade *ModifyClusterRequestOperationPolicyClusterAutoUpgrade `json:"cluster_auto_upgrade,omitempty" xml:"cluster_auto_upgrade,omitempty" type:"Struct"`
}

func (s ModifyClusterRequestOperationPolicy) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterRequestOperationPolicy) GoString() string {
	return s.String()
}

func (s *ModifyClusterRequestOperationPolicy) GetClusterAutoUpgrade() *ModifyClusterRequestOperationPolicyClusterAutoUpgrade {
	return s.ClusterAutoUpgrade
}

func (s *ModifyClusterRequestOperationPolicy) SetClusterAutoUpgrade(v *ModifyClusterRequestOperationPolicyClusterAutoUpgrade) *ModifyClusterRequestOperationPolicy {
	s.ClusterAutoUpgrade = v
	return s
}

func (s *ModifyClusterRequestOperationPolicy) Validate() error {
	if s.ClusterAutoUpgrade != nil {
		if err := s.ClusterAutoUpgrade.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyClusterRequestOperationPolicyClusterAutoUpgrade struct {
	// The cluster auto upgrade frequency. For more information, see [Upgrade frequency](https://help.aliyun.com/document_detail/2712866.html).
	//
	// Valid values:
	//
	// - patch: latest patch version.
	//
	// - stable: second-latest minor version.
	//
	// - rapid: latest minor version.
	//
	// example:
	//
	// patch
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// Specifies whether to enable cluster auto upgrade.
	//
	// - true: Enables auto upgrade.
	//
	// - false: Disables auto upgrade.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s ModifyClusterRequestOperationPolicyClusterAutoUpgrade) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterRequestOperationPolicyClusterAutoUpgrade) GoString() string {
	return s.String()
}

func (s *ModifyClusterRequestOperationPolicyClusterAutoUpgrade) GetChannel() *string {
	return s.Channel
}

func (s *ModifyClusterRequestOperationPolicyClusterAutoUpgrade) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyClusterRequestOperationPolicyClusterAutoUpgrade) SetChannel(v string) *ModifyClusterRequestOperationPolicyClusterAutoUpgrade {
	s.Channel = &v
	return s
}

func (s *ModifyClusterRequestOperationPolicyClusterAutoUpgrade) SetEnabled(v bool) *ModifyClusterRequestOperationPolicyClusterAutoUpgrade {
	s.Enabled = &v
	return s
}

func (s *ModifyClusterRequestOperationPolicyClusterAutoUpgrade) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterRequestSystemEventsLogging struct {
	// Specifies whether to enable system event storage.
	//
	//
	// - true: Enables system event storage.
	//
	// - false: Disables system event storage.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The name of the LogProject used for system event storage.
	//
	// example:
	//
	// k8s-log-cb95aa626a47740afbf6aa099b65****
	LoggingProject *string `json:"logging_project,omitempty" xml:"logging_project,omitempty"`
}

func (s ModifyClusterRequestSystemEventsLogging) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterRequestSystemEventsLogging) GoString() string {
	return s.String()
}

func (s *ModifyClusterRequestSystemEventsLogging) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyClusterRequestSystemEventsLogging) GetLoggingProject() *string {
	return s.LoggingProject
}

func (s *ModifyClusterRequestSystemEventsLogging) SetEnabled(v bool) *ModifyClusterRequestSystemEventsLogging {
	s.Enabled = &v
	return s
}

func (s *ModifyClusterRequestSystemEventsLogging) SetLoggingProject(v string) *ModifyClusterRequestSystemEventsLogging {
	s.LoggingProject = &v
	return s
}

func (s *ModifyClusterRequestSystemEventsLogging) Validate() error {
	return dara.Validate(s)
}
