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
	// Access control list for the registered cluster API Server SLB.
	AccessControlList []*string `json:"access_control_list,omitempty" xml:"access_control_list,omitempty" type:"Repeated"`
	// Custom API Server certificate SAN (Subject Alternative Name).
	//
	// Used to add custom IPs or domain names to the SAN field of the cluster API Server server certificate for client access control.
	//
	// Only managed clusters support this parameter.
	ApiServerCustomCertSans *ModifyClusterRequestApiServerCustomCertSans `json:"api_server_custom_cert_sans,omitempty" xml:"api_server_custom_cert_sans,omitempty" type:"Struct"`
	// Whether to associate an EIP with the cluster for public access to API Server. Valid values:
	//
	// - `true`: Associate an EIP with the cluster.
	//
	// - `false`: Do not associate an EIP with the cluster.
	//
	// example:
	//
	// true
	ApiServerEip *bool `json:"api_server_eip,omitempty" xml:"api_server_eip,omitempty"`
	// The ID of the EIP instance associated with the cluster API Server. This parameter takes effect only when `api_server_eip` is set to `true`.
	//
	// example:
	//
	// eip-wz9fnasl6dsfhmvci****
	ApiServerEipId *string `json:"api_server_eip_id,omitempty" xml:"api_server_eip_id,omitempty"`
	// Custom cluster name. The name can contain digits, Chinese characters, English characters, or hyphens (-), must be 1 to 63 characters in length, and cannot start with a hyphen (-).
	//
	// example:
	//
	// cluster-new-name
	ClusterName *string `json:"cluster_name,omitempty" xml:"cluster_name,omitempty"`
	// When `cluster_type` is set to `ManagedKubernetes` and `profile` is configured, specifies the cluster specification. Valid values:
	//
	// - `ack.pro.small`: Pro Edition
	//
	// - `ack.pro.xlarge`: Pro XL
	//
	// - `ack.pro.2xlarge`: Pro 2XL
	//
	// - `ack.pro.4xlarge`: Pro 4XL (requires contacting customer service to enable allowlisting)
	//
	// Pro XL, Pro 2XL, and Pro 4XL are three tiers provided by <props="china">[ACK Pro Provisioned Control Plane](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane)<props="intl">[ACK Pro Provisioned Control Plane](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane). By pre-allocating and fixing control plane resources, it ensures that API concurrency and Pod scheduling capabilities are always at a determined high level, suitable for AI training and inference, ultra-large-scale clusters, and mission-critical workloads.
	//
	// For cluster management fees for Pro Edition and Provisioned Control Plane editions, see <props="china">[Cluster Management Fees](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee)<props="intl">[Cluster Management Fees](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee).
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// Dedicated cluster control plane configuration.
	ControlPlaneConfig *ModifyClusterRequestControlPlaneConfig `json:"control_plane_config,omitempty" xml:"control_plane_config,omitempty" type:"Struct"`
	// Cluster connection configuration.
	ControlPlaneEndpointsConfig *ModifyClusterRequestControlPlaneEndpointsConfig `json:"control_plane_endpoints_config,omitempty" xml:"control_plane_endpoints_config,omitempty" type:"Struct"`
	// Cluster deletion protection, which prevents accidental deletion of the cluster through the console or API. Valid values:
	//
	// - `true`: Enable cluster deletion protection. The cluster cannot be deleted through the console or API.
	//
	// - `false`: Disable cluster deletion protection. The cluster can be deleted through the console or API.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// Enable or disable the RRSA feature (only managed clusters support this parameter). Valid values:
	//
	// - `true`: Enable.
	//
	// - `false`: Disable.
	//
	// example:
	//
	// true
	EnableRrsa *bool `json:"enable_rrsa,omitempty" xml:"enable_rrsa,omitempty"`
	// Deprecated
	//
	// Rebind the cluster test domain. Valid values:
	//
	// - `true`: Rebind the cluster test domain.
	//
	// - `false`: Do not rebind the cluster test domain.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	IngressDomainRebinding *bool `json:"ingress_domain_rebinding,omitempty" xml:"ingress_domain_rebinding,omitempty"`
	// Deprecated
	//
	// SLB instance ID of the cluster to be modified.
	//
	// example:
	//
	// lb-wz97kes8tnndkpodw****
	IngressLoadbalancerId *string `json:"ingress_loadbalancer_id,omitempty" xml:"ingress_loadbalancer_id,omitempty"`
	// Deprecated
	//
	// Instance deletion protection to prevent accidental deletion and release of nodes through the console or API. Valid values:
	//
	// - `true`: Nodes cannot be accidentally deleted through the console or API.
	//
	// - `false`: Nodes can be deleted through the console or API.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	InstanceDeletionProtection *bool `json:"instance_deletion_protection,omitempty" xml:"instance_deletion_protection,omitempty"`
	// Cluster maintenance window. This feature only takes effect for ACK Pro managed clusters.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window,omitempty" xml:"maintenance_window,omitempty"`
	// Cluster automatic O&M policy.
	OperationPolicy *ModifyClusterRequestOperationPolicy `json:"operation_policy,omitempty" xml:"operation_policy,omitempty" type:"Struct"`
	// Cluster resource group ID.
	//
	// example:
	//
	// rg-acfmyvw3wjm****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// Control plane security group ID.
	//
	// - If you have configured blocking rules in the security group, ensure that the security group rules allow the protocols and ports required by the cluster. For recommended security group rules, see [Configure and Manage Cluster Security Groups](https://help.aliyun.com/document_detail/353191.html).
	//
	// - For non-ACK dedicated clusters, during the change process, the cluster control plane and installed managed components (such as terway-controlplane) will briefly restart. We recommend performing this operation during off-peak hours. After the control plane security group is changed, the ENIs used by the cluster control plane and installed managed components will be automatically added to the new security group.
	//
	// - For ACK dedicated clusters, after the control plane security group is changed, newly scaled-out Master nodes will automatically use the new control plane security group. Existing control plane nodes are not affected.
	//
	// example:
	//
	// sg-bp1h6rk3pgct2a08***
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// System event storage configuration.
	SystemEventsLogging *ModifyClusterRequestSystemEventsLogging `json:"system_events_logging,omitempty" xml:"system_events_logging,omitempty" type:"Struct"`
	// Cluster timezone. See [Supported Timezones](https://help.aliyun.com/document_detail/354879.html).
	//
	// - After changing the timezone, cluster inspection configurations will use the new timezone settings.
	//
	// - For managed clusters, during the change process, the cluster control plane and installed managed components (such as terway-controlplane) will briefly restart. We recommend performing this operation during off-peak hours. After changing the timezone, newly scaled-out nodes will automatically use the new timezone settings. Existing nodes are not affected. You can use the node pool node reset feature to apply the new settings to existing nodes.
	//
	// - For dedicated clusters, after changing the timezone, newly scaled-out nodes (including control plane nodes) will automatically use the new timezone settings. Existing nodes (including control plane nodes) are not affected. You can use the node pool node reset feature to apply the new settings to existing nodes. For control plane nodes, you need to scale out first and then scale in to apply the settings to all control plane nodes.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// Cluster control plane vSwitches. For dedicated clusters, this takes effect on newly scaled-out control plane nodes. When modifying control plane vSwitches for managed clusters, note the following:
	//
	// - This parameter performs a full overwrite update. You must specify the complete list of target vSwitches.
	//
	// - During the change, control plane components will briefly restart. Proceed with caution.
	//
	// - Ensure that all security groups of the cluster (including the control plane security group, all node pool security groups, and container network security groups) allow inbound and outbound traffic for the IP ranges of the new vSwitches to prevent nodes and containers from being unable to connect to the API Server.
	//
	// - If the new control plane vSwitches have ACL rules configured, ensure that the ACL rules allow communication with the cluster nodes, container network, and other IP ranges.
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
	// Overwrite or append SAN configuration. Valid values:
	//
	// - overwrite: Overwrite.
	//
	// - append: Append.
	//
	// example:
	//
	// append
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// SAN list.
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
	// Whether to enable automatic renewal for control plane node instances. This parameter takes effect only when `charge_type` is set to `PrePaid`. Valid values:
	//
	// - `true`: Enable automatic renewal.
	//
	// - `false`: Disable automatic renewal.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// Duration for each automatic renewal of control plane node instances.
	//
	// Valid values: {1, 2, 3, 6, 12}. Unit: months.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// Control plane node instance billing method. Valid values:
	//
	// - `PrePaid`: Subscription.
	//
	// - `PostPaid`: Pay-as-you-go.
	//
	// Default value: `PostPaid`.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"charge_type,omitempty" xml:"charge_type,omitempty"`
	// Whether to install the Cloud Monitor agent on control plane nodes. Valid values:
	//
	// - `true`: Install the Cloud Monitor agent.
	//
	// - `false`: Do not install the Cloud Monitor agent.
	//
	// example:
	//
	// true
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// Node CPU management policy. When the cluster version is 1.12.6 or later, the following two policies are supported:
	//
	// - `static`: Allows enhanced CPU affinity and exclusivity for Pods with certain resource characteristics on the node.
	//
	// - `none`: Uses the existing default CPU affinity scheme.
	//
	// Default value: `none`.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// Deployment set ID.
	//
	// example:
	//
	// ds-bp10b35imuam5amw****
	DeploymentsetId *string `json:"deploymentset_id,omitempty" xml:"deploymentset_id,omitempty"`
	// Custom image ID. Specified when using a custom image.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20240819.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// Operating system image type. Valid values:
	//
	// - `AliyunLinux3`: Alinux3 image.
	//
	// - `Custom`: Custom image.
	//
	// example:
	//
	// AliyunLinux3
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// Instance types. For more information, see [Instance Family](https://help.aliyun.com/document_detail/25378.html).
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// Key pair name. Mutually exclusive with `login_password`.
	//
	// example:
	//
	// ack
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// SSH login password. Mutually exclusive with `key_pair`. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. If you want to use password login, specify this parameter during scale-out.
	//
	// example:
	//
	// Ack@2000.
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// Node service port range.
	//
	// Available port range: [30000, 65535].
	//
	// Default value: 30000-32767.
	//
	// example:
	//
	// 30000-32767
	NodePortRange *string `json:"node_port_range,omitempty" xml:"node_port_range,omitempty"`
	// Control plane node instance subscription duration. This parameter takes effect and is required only when `charge_type` is set to `PrePaid`.
	//
	// When `period_unit=Month`, valid values: {1, 2, 3, 6, 12, 24, 36, 48, 60}.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// Control plane node instance billing period. This parameter takes effect only when `charge_type` is set to `PrePaid`.
	//
	// `Month`: Billed on a monthly basis. Currently, only monthly billing is supported.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// Container runtime name. Valid values:
	//
	// - `containerd`: Recommended. Supported by all cluster versions.
	//
	// Default value: containerd.
	//
	// example:
	//
	// containerd
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// Alibaba Cloud OS security hardening. Valid values:
	//
	// - `true`: Enable Alibaba Cloud OS security hardening.
	//
	// - `false`: Disable Alibaba Cloud OS security hardening.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SecurityHardeningOs *bool `json:"security_hardening_os,omitempty" xml:"security_hardening_os,omitempty"`
	// Number of control plane nodes. To scale out the dedicated cluster control plane, this parameter specifies the target number of control plane nodes and must be greater than the current number of control plane nodes.
	//
	// example:
	//
	// 5
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// Security hardening for compliance. For more information, see [ACK Security Hardening for Compliance](https://help.aliyun.com/document_detail/196148.html).
	//
	// Valid values:
	//
	// - `true`: Enable security hardening for compliance.
	//
	// - `false`: Disable security hardening for compliance.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// Whether to enable burst (performance bursting) for the node system disk. Valid values:
	//
	// - `true`: Enable.
	//
	// - `false`: Disable.
	//
	// This parameter is supported only when `system_disk_category` is set to `cloud_auto`. For more information, see [ESSD AutoPL](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// Node system disk type. Valid values:
	//
	// - `cloud_efficiency`: Ultra disk.
	//
	// - `cloud_ssd`: SSD disk.
	//
	// - `cloud_essd`: ESSD disk.
	//
	// - `cloud_auto`: ESSD AutoPL disk.
	//
	// - `cloud_essd_entry`: ESSD Entry disk.
	//
	// example:
	//
	// cloud_essd
	SystemDiskCategory *string `json:"system_disk_category,omitempty" xml:"system_disk_category,omitempty"`
	// Node system disk performance level. Only applicable to ESSD disks. The performance level is related to the disk size. For more information, see [ESSD](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// Provisioned read/write IOPS for the node system disk. Valid values: 0 to min{50,000, 1000*capacity - baseline performance}. Baseline performance = min{1,800 + 50*capacity, 50,000}.
	//
	// This parameter is supported only when `system_disk_category` is set to `cloud_auto`. For more information, see [ESSD AutoPL](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// Node system disk size. Valid values: [40, 500]. Unit: GiB.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// Automatic snapshot policy ID for the node system disk.
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
	// Cluster internal domain name configuration. Applicable to ACK managed clusters. The cluster internal domain name is used by node-side system components such as kubelet and kube-proxy to access the API Server. When the cluster internal domain name access is not enabled, node-side system components access via the CLB IP.
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
	// VPCs where the cluster internal domain name record resolution takes effect.
	BindVpcs []*string `json:"bind_vpcs,omitempty" xml:"bind_vpcs,omitempty" type:"Repeated"`
	// Whether to enable cluster internal domain name access. Valid values:
	//
	// - true: Enable cluster internal domain name access. Node-side components (kubelet, kube-proxy) will access the API Server through the cluster internal domain name.
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

type ModifyClusterRequestOperationPolicy struct {
	// Cluster automatic upgrade.
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
	// Cluster automatic upgrade frequency. For more information, see [Upgrade Frequency](https://help.aliyun.com/document_detail/2712866.html).
	//
	// Valid values:
	//
	// - patch: Latest patch version.
	//
	// - stable: Second latest minor version.
	//
	// - rapid: Latest minor version.
	//
	// example:
	//
	// patch
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// Whether to enable cluster automatic upgrade.
	//
	// - true: Enable automatic upgrade.
	//
	// - false: Disable automatic upgrade.
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
	// Whether to enable system event storage.
	//
	//
	// - true: Enable system event storage.
	//
	// - false: Disable system event storage.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// LogProject name for system event storage.
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
