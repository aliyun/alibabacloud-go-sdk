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
	SetClusterName(v string) *ModifyClusterRequest
	GetClusterName() *string
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
	// The network access control list (ACL) of the SLB instance associated with the API server if the cluster is a registered cluster.
	AccessControlList []*string `json:"access_control_list,omitempty" xml:"access_control_list,omitempty" type:"Repeated"`
	// The custom subject alternative names (SANs) for the API server certificate to accept requests from specified IP addresses or domain names. This parameter is available only for ACK managed clusters.
	ApiServerCustomCertSans *ModifyClusterRequestApiServerCustomCertSans `json:"api_server_custom_cert_sans,omitempty" xml:"api_server_custom_cert_sans,omitempty" type:"Struct"`
	// Specifies whether to associate an elastic IP address (EIP) with the cluster. This EIP is used to enable access to the API server over the Internet. Valid values:
	//
	// 	- `true`: associates an EIP with the cluster.
	//
	// 	- `false`: does not associate an EIP with the cluster.
	//
	// example:
	//
	// true
	ApiServerEip *bool `json:"api_server_eip,omitempty" xml:"api_server_eip,omitempty"`
	// The ID of the EIP that you want to associate with the API server of the cluster. This parameter takes effect when `api_server_eip` is set to `true`.
	//
	// example:
	//
	// eip-wz9fnasl6dsfhmvci****
	ApiServerEipId *string `json:"api_server_eip_id,omitempty" xml:"api_server_eip_id,omitempty"`
	// The cluster name.
	//
	// The cluster name must be 1 to 63 characters in length, and can contain digits, letters, and hyphens (-). The cluster name cannot start with a hyphen (-).
	//
	// example:
	//
	// cluster-new-name
	ClusterName *string `json:"cluster_name,omitempty" xml:"cluster_name,omitempty"`
	// The control plane configurations of an ACK dedicated cluster.
	ControlPlaneConfig          *ModifyClusterRequestControlPlaneConfig          `json:"control_plane_config,omitempty" xml:"control_plane_config,omitempty" type:"Struct"`
	ControlPlaneEndpointsConfig *ModifyClusterRequestControlPlaneEndpointsConfig `json:"control_plane_endpoints_config,omitempty" xml:"control_plane_endpoints_config,omitempty" type:"Struct"`
	// Specifies whether to enable cluster deletion protection. If you enable this option, the cluster cannot be deleted in the console or by calling API operations. Valid values:
	//
	// 	- `true`: enables cluster deletion protection.
	//
	// 	- `false`: disables cluster deletion protection.
	//
	// Default value: `false`
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// Specifies whether to enable the RAM Roles for Service Accounts (RRSA) feature. This parameter is available only for ACK managed clusters. Valid values:
	//
	// 	- `true`: enables the RRSA feature.
	//
	// 	- `false`: disables the RRSA feature.
	//
	// example:
	//
	// true
	EnableRrsa *bool `json:"enable_rrsa,omitempty" xml:"enable_rrsa,omitempty"`
	// Deprecated
	//
	// Specifies whether to remap the test domain name of the cluster. Valid values:
	//
	// 	- `true`: remaps the test domain name of the cluster.
	//
	// 	- `false`: does not remap the test domain name of the cluster.
	//
	// Default value: `false`
	//
	// example:
	//
	// true
	IngressDomainRebinding *bool `json:"ingress_domain_rebinding,omitempty" xml:"ingress_domain_rebinding,omitempty"`
	// Deprecated
	//
	// The ID of the Server Load Balancer (SLB) instance of the cluster to be modified.
	//
	// example:
	//
	// lb-wz97kes8tnndkpodw****
	IngressLoadbalancerId *string `json:"ingress_loadbalancer_id,omitempty" xml:"ingress_loadbalancer_id,omitempty"`
	// Deprecated
	//
	// Specifies whether to enable instance deletion protection. If you enable this option, the instance cannot be deleted in the console or by calling API operations. Valid values:
	//
	// 	- `true`: enables instance deletion protection.
	//
	// 	- `false`: disables instance deletion protection.
	//
	// Default value: `false`
	//
	// example:
	//
	// true
	InstanceDeletionProtection *bool `json:"instance_deletion_protection,omitempty" xml:"instance_deletion_protection,omitempty"`
	// The cluster maintenance window. This feature takes effect only for ACK Pro clusters.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window,omitempty" xml:"maintenance_window,omitempty"`
	// The automatic O\\&M policy of the cluster.
	OperationPolicy *ModifyClusterRequestOperationPolicy `json:"operation_policy,omitempty" xml:"operation_policy,omitempty" type:"Struct"`
	// The resource group ID of the cluster.
	//
	// example:
	//
	// rg-acfmyvw3wjm****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// The ID of the security group for the control plane.
	//
	// - If block rules are configured in the security group, ensure the security group rules allow traffic for protocols and ports required by the cluster. For recommended security group rules, see [Configure and manage security groups for an ACK cluster](https://www.alibabacloud.com/help/en/ack/ack-managed-and-ack-dedicated/user-guide/configure-security-group-rules-to-enforce-access-control-on-ack-clusters?spm=a2c63.p38356.help-menu-85222.d_2_0_4_3.43e35d09s8oSlR).
	//
	// - For non-ACK dedicated clusters:
	//
	//   - During security group updates, the cluster control plane and managed components (e.g., terway-controlplane) will restart briefly. Perform this operation during off-peak hours.
	//
	//   - After updating the control plane security group, the Elastic Network Interfaces (ENIs) used by the control plane and managed components will automatically join the new security group.
	//
	// - For ACK dedicated clusters:
	//
	//    - After updating the control plane security group, newly scaled-out master nodes will automatically apply the new security group. Existing control plane nodes remain unaffected.
	//
	// example:
	//
	// sg-bp1h6rk3pgct2a08***
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// The storage configurations of system events.
	SystemEventsLogging *ModifyClusterRequestSystemEventsLogging `json:"system_events_logging,omitempty" xml:"system_events_logging,omitempty" type:"Struct"`
	// The time zone configuration for the cluster.
	//
	// - After modifying the time zone, cluster inspection configurations will adopt the new time zone.
	//
	// - For ACK managed clusters:
	//
	//    - During time zone updates, the cluster control plane and managed components (e.g., terway-controlplane) will restart briefly. Perform this operation during off-peak hours.
	//
	//    - After updating the time zone:
	//
	//       - Newly scaled-out nodes will automatically apply the new time zone.
	//
	//       - Existing nodes remain unaffected. Reset the node to apply changes to existing nodes.
	//
	// - For ACK dedicated clusters:
	//
	//    - After updating the time zone:
	//
	//       - Newly scaled-out nodes (including control plane nodes) automatically apply the new time zone.
	//
	//       - Existing nodes (including control plane nodes) remain unaffected. Reset the node to apply changes to existing nodes.
	//
	//       - For control plane nodes, perform a scale-out followed by a scale-in to apply the new time zone to all control plane nodes.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The vSwitches of the control plane. This parameter can be used to change the vSwitches of the control plane in an ACK managed cluster. Take note of the following items:
	//
	// 	- This parameter overwrites the existing configuration. You must specify all vSwitches of the control plane.
	//
	// 	- The control plane components restarts during the change process. Exercise caution when you perform this operation.
	//
	// 	- Ensure that all security groups of the cluster, including the security groups of the control plane, all node pools, and container network, are allowed to access the CIDR blocks of the new vSwitches. This ensures that the nodes and containers can connect to the API server.
	//
	// 	- If the new vSwitches of the control plane are configured with an ACL, ensure that the ACL allows communication between the new vSwitches and CIDR blocks such as those of the cluster nodes and the container network.
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

func (s *ModifyClusterRequest) GetClusterName() *string {
	return s.ClusterName
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

func (s *ModifyClusterRequest) SetClusterName(v string) *ModifyClusterRequest {
	s.ClusterName = &v
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
	// Specifies whether to overwrite or add SANs. Valid values:
	//
	// 	- overwrite: overwrites SANs.
	//
	// 	- append: adds SANs.
	//
	// example:
	//
	// append
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The list of SANs.
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
	// Specifies whether to enable auto-renewal for control plane nodes. This parameter takes effect only when `charge_type` is set to `PrePaid`. Valid values:
	//
	// 	- `true`: enables auto-renewal.
	//
	// 	- `false`: disables auto-renewal.
	//
	// Default value: `false`
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// The auto-renewal period of control plane nodes. Valid values: 1, 2, 3, 6, and 12.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// The billing method of control plane nodes. Valid values:
	//
	// 	- `PrePaid`: subscription.
	//
	// 	- `PostPaid`: pay-as-you-go.
	//
	// Default value: `PostPaid`.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"charge_type,omitempty" xml:"charge_type,omitempty"`
	// Specifies whether to install the CloudMonitor agent. Valid values:
	//
	// 	- `true`: installs the CloudMonitor agent.
	//
	// 	- `false`: does not install the CloudMonitor agent.
	//
	// example:
	//
	// true
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// The CPU management policy of nodes in the node pool. The following policies are supported if the Kubernetes version of the cluster is 1.12.6 or later:
	//
	// 	- `static`: allows pods with specific resource characteristics on the node to be granted with enhanced CPU affinity and exclusivity.
	//
	// 	- `none`: specifies that the default CPU affinity is used.
	//
	// Default value: `none`.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// The ID of the deployment set.
	//
	// example:
	//
	// ds-bp10b35imuam5amw****
	DeploymentsetId *string `json:"deploymentset_id,omitempty" xml:"deploymentset_id,omitempty"`
	// The custom image ID. You must configure this parameter if you use a custom image.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20240819.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The type of the OS image. Valid values:
	//
	// 	- `AliyunLinux3`: Alibaba Cloud Linux 3.
	//
	// 	- `Custom`: the custom image.
	//
	// example:
	//
	// AliyunLinux3
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// The type of instance. For more information, see [Overview of ECS instance families](https://help.aliyun.com/document_detail/25378.html).
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// The name of the key pair. You must configure either this parameter or the `login_password` parameter.
	//
	// example:
	//
	// ack
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// The password for SSH logon. You must configure either this parameter or the `key_pair` parameter. The password must be 8 to 30 characters in length, and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. To log on with a password, you must specify this parameter during the scale-out.
	//
	// example:
	//
	// Ack@2000.
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The node port range.
	//
	// example:
	//
	// 30000-32767
	NodePortRange *string `json:"node_port_range,omitempty" xml:"node_port_range,omitempty"`
	// The subscription duration of the instance. This parameter takes effect and is required only when `charge_type` is set to `PrePaid`.
	//
	// If `PeriodUnit=Month` is specified, the valid values are 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The billing cycle of control plane nodes. This parameter takes effect only when `instance_charge_type` is set to `PrePaid`.
	//
	// Set the value to `Month`.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// The type of the container runtime. Valid values:
	//
	// 	- `containerd`: supports all Kubernetes versions. We recommend that you set the parameter to this value.
	//
	// Default value: containerd.
	//
	// example:
	//
	// containerd
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// Specifies whether to enable Alibaba Cloud Linux Security Hardening. Valid values:
	//
	// 	- `true`: enables Alibaba Cloud Linux Security Hardening.
	//
	// 	- `false`: disables Alibaba Cloud Linux Security Hardening.
	//
	// Default value: `false`
	//
	// example:
	//
	// true
	SecurityHardeningOs *bool `json:"security_hardening_os,omitempty" xml:"security_hardening_os,omitempty"`
	// The number of control plane nodes. If you want to scale out the control plane in an ACK dedicated cluster, set this parameter to the desired number of nodes. This parameter must be greater than the current number of nodes.
	//
	// example:
	//
	// 5
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// Specifies whether to enable Multi-Level Protection Scheme (MLPS) security hardening. For more information, see [ACK security hardening based on MLPS](https://help.aliyun.com/document_detail/196148.html).
	//
	// Valid values:
	//
	// 	- `true`: enables MLPS security hardening.
	//
	// 	- `false`: disables MLPS security hardening.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// Specifies whether to enable the burst feature for the system disk. Valid values:
	//
	// 	- `true`: enables the burst feature.
	//
	// 	- `false`: disables the burst feature.
	//
	// This parameter is effective only when `system_disk_category` is set to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// The category of the system disk for nodes. Valid values:
	//
	// 	- `cloud`: basic disk.
	//
	// 	- `cloud_efficiency`: ultra disk.
	//
	// 	- `cloud_ssd`: standard SSD.
	//
	// 	- `cloud_essd`: Enterprise ESSD (ESSD).
	//
	// 	- `cloud_auto`: ESSD AutoPL disk.
	//
	// 	- `cloud_essd_entry`: ESSD Entry disk.
	//
	// example:
	//
	// cloud_essd
	SystemDiskCategory *string `json:"system_disk_category,omitempty" xml:"system_disk_category,omitempty"`
	// The performance level (PL) of the system disk that you want to use for the node. This parameter is effective only for ESSDs. This parameter is related to the disk size. For more information, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// The preset read/write input/output operations per second (IOPS) of the system disk. Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}. Baseline IOPS = min{1,800 + 50 × Capacity, 50,000}.
	//
	// This parameter is effective only when `system_disk_category` is set to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// The type of the system disk. Valid values: [40,500]. Unit: GiB.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// The ID of the automatic snapshot policy applied to the node system disk.
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
	InternalDnsConfig *ModifyClusterRequestControlPlaneEndpointsConfigInternalDnsConfig `json:"internal_dns_config,omitempty" xml:"internal_dns_config,omitempty" type:"Struct"`
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

func (s *ModifyClusterRequestControlPlaneEndpointsConfig) SetInternalDnsConfig(v *ModifyClusterRequestControlPlaneEndpointsConfigInternalDnsConfig) *ModifyClusterRequestControlPlaneEndpointsConfig {
	s.InternalDnsConfig = v
	return s
}

func (s *ModifyClusterRequestControlPlaneEndpointsConfig) Validate() error {
	if s.InternalDnsConfig != nil {
		if err := s.InternalDnsConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyClusterRequestControlPlaneEndpointsConfigInternalDnsConfig struct {
	BindVpcs []*string `json:"bind_vpcs,omitempty" xml:"bind_vpcs,omitempty" type:"Repeated"`
	Enabled  *bool     `json:"enabled,omitempty" xml:"enabled,omitempty"`
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

type ModifyClusterRequestOperationPolicy struct {
	// The configurations of automatic update.
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
	// The frequency of auto cluster update. For more information, see [Update frequency](https://help.aliyun.com/document_detail/2712866.html).
	//
	// Valid values:
	//
	// 	- patch: the latest patch version.
	//
	// 	- stables: the second-latest minor version.
	//
	// 	- rapid: the latest minor version.
	//
	// example:
	//
	// patch
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// Specifies whether to enable automatic update.
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
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The name of the Simple Log Service project that stores system events.
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
