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
	SetSupportPolicy(v string) *ModifyClusterRequest
	GetSupportPolicy() *string
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
	// The access control list of the SLB instance for the API server of the registered cluster.
	AccessControlList []*string `json:"access_control_list,omitempty" xml:"access_control_list,omitempty" type:"Repeated"`
	// The custom Subject Alternative Names (SANs) for the API server certificate. You can add custom IP addresses or domain names to the SAN field of the API server certificate to implement access control for clients.
	//
	// Only managed clusters support this parameter.
	ApiServerCustomCertSans *ModifyClusterRequestApiServerCustomCertSans `json:"api_server_custom_cert_sans,omitempty" xml:"api_server_custom_cert_sans,omitempty" type:"Struct"`
	// Specifies whether an EIP is attached to the cluster for public network access to the API server. Valid values:
	//
	// - `true`: attaches an EIP to the cluster.
	//
	// - `false`: does not attach an EIP to the cluster.
	//
	// example:
	//
	// true
	ApiServerEip *bool `json:"api_server_eip,omitempty" xml:"api_server_eip,omitempty"`
	// The EIP instance ID attached to the API server of the cluster. This parameter takes effect only when `api_server_eip` is set to `true`.
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
	// The custom cluster name. The name must be 1 to 63 characters in length and can contain digits, Chinese characters, letters, and hyphens (-). It cannot start with a hyphen (-).
	//
	// example:
	//
	// cluster-new-name
	ClusterName *string `json:"cluster_name,omitempty" xml:"cluster_name,omitempty"`
	// The specification of the cluster when `cluster_type` is set to `ManagedKubernetes` and `profile` is configured. Valid values:
	//
	// - `ack.pro.small`: Pro Edition
	//
	// - `ack.pro.xlarge`: Pro XL
	//
	// - `ack.pro.2xlarge`: Pro 2XL
	//
	// - `ack.pro.4xlarge`: Pro 4XL (contact customer service to apply for whitelist access)
	//
	// Pro XL, Pro 2XL, and Pro 4XL are the three tiers provided by the <props="china">[ACK Pro provisioned control plane](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane)<props="intl">[ACK Pro provisioned control plane](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane). By pre-allocating and reserving control plane resources, these tiers ensure that API concurrency and pod scheduling capabilities are always maintained at a consistently high level. They are suitable for AI training and inference, ultra-large-scale clusters, and mission-critical workloads.
	//
	// For more information about the cluster management fees for the Pro Edition and provisioned control plane editions, see <props="china">[Cluster management fees](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee)<props="intl">[Cluster management fees](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee).
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// The control plane configuration for dedicated clusters.
	ControlPlaneConfig *ModifyClusterRequestControlPlaneConfig `json:"control_plane_config,omitempty" xml:"control_plane_config,omitempty" type:"Struct"`
	// The cluster connection configuration.
	ControlPlaneEndpointsConfig *ModifyClusterRequestControlPlaneEndpointsConfig `json:"control_plane_endpoints_config,omitempty" xml:"control_plane_endpoints_config,omitempty" type:"Struct"`
	// Specifies whether to enable deletion protection for the cluster to prevent accidental deletion through the console or API calls. Valid values:
	//
	// - `true`: Enables deletion protection. The cluster cannot be deleted through the console or API calls.
	//
	// - `false`: Disables deletion protection. The cluster can be deleted through the console or API calls.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// Specifies whether to enable the RAM Roles for Service Accounts (RRSA) feature. Only managed clusters support this parameter. Valid values:
	//
	// - `true`: Enables the feature.
	//
	// - `false`: Disables the feature.
	//
	// example:
	//
	// true
	EnableRrsa *bool `json:"enable_rrsa,omitempty" xml:"enable_rrsa,omitempty"`
	// Deprecated
	//
	// Specifies whether to rebind the test domain name of the cluster to the SLB address of the Ingress. Valid values:
	//
	// - `true`: Rebinds the test domain name of the cluster.
	//
	// - `false`: Does not rebind the test domain name of the cluster.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	IngressDomainRebinding *bool `json:"ingress_domain_rebinding,omitempty" xml:"ingress_domain_rebinding,omitempty"`
	// Deprecated
	//
	// The instance ID of the SLB for the Ingress of the cluster to be modified.
	//
	// example:
	//
	// lb-wz97kes8tnndkpodw****
	IngressLoadbalancerId *string `json:"ingress_loadbalancer_id,omitempty" xml:"ingress_loadbalancer_id,omitempty"`
	// Deprecated
	//
	// Specifies whether to enable deletion protection for instances to prevent accidental deletion or release of nodes through the console or API calls. Valid values:
	//
	// - `true`: Enables deletion protection. Nodes cannot be accidentally deleted or released through the console or API calls.
	//
	// - `false`: Disables deletion protection. Nodes can be accidentally deleted or released through the console or API calls.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	InstanceDeletionProtection *bool `json:"instance_deletion_protection,omitempty" xml:"instance_deletion_protection,omitempty"`
	// The maintenance window of the cluster. This feature takes effect only in ACK managed clusters of the Pro edition.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window,omitempty" xml:"maintenance_window,omitempty"`
	// The automatic O&M policy for the cluster.
	OperationPolicy *ModifyClusterRequestOperationPolicy `json:"operation_policy,omitempty" xml:"operation_policy,omitempty" type:"Struct"`
	// The ID of the resource group to which the cluster belongs.
	//
	// example:
	//
	// rg-acfmyvw3wjm****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// The ID of the control plane security group.
	//
	// - If you have configured blocking rules in the security group, make sure that the security group rules allow the protocols and ports on which the cluster depends. For more information about the recommended security group rules for cluster management, see [Configure and manage cluster security groups](https://help.aliyun.com/document_detail/353191.html).
	//
	// - For ACK clusters other than dedicated clusters, the cluster control plane and installed managed components such as terway-controlplane will be briefly restarted during the procedure. We recommend that you perform this operation during off-peak hours. After the control plane security group is changed, the elastic network interfaces (ENIs) used by the cluster control plane and installed managed components are automatically added to the new security group.
	//
	// - For ACK dedicated clusters, after the control plane security group is changed, newly scaled-out master nodes automatically apply the new control plane security group. Existing control plane nodes are not affected.
	//
	// example:
	//
	// sg-bp1h6rk3pgct2a08***
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// The version support policy for the cluster. The value is case-sensitive. Valid values:
	//
	// - Standard: Standard support.
	//
	// - Extended: Extended support. After the standard support period ends, the cluster enters the extended support period and continues to receive CVE patch versions.
	//
	// Only ACK managed clusters of the Pro edition (where cluster_spec is set to ack.pro.	- and profile is set to Default) support the Extended value. If you specify Extended for ACK standard clusters, ACS clusters, ACK managed clusters of the edge edition, or ACK dedicated clusters, the request will be rejected.
	//
	// If you do not specify this parameter, the current value of the cluster remains unchanged. After a cluster is upgraded to an extended support patch version, it cannot be changed back to Standard.
	//
	// example:
	//
	// Extended
	SupportPolicy *string `json:"support_policy,omitempty" xml:"support_policy,omitempty"`
	// The system event logging configuration.
	SystemEventsLogging *ModifyClusterRequestSystemEventsLogging `json:"system_events_logging,omitempty" xml:"system_events_logging,omitempty" type:"Struct"`
	// The time zone of the cluster. For more information, see [Supported time zones](https://help.aliyun.com/document_detail/354879.html).
	//
	// - After the time zone is changed, the cluster inspection configuration adopts the new time zone settings.
	//
	// - For managed clusters, the cluster control plane and installed managed components such as terway-controlplane will be briefly restarted during the modification. We recommend that you perform this operation during off-peak hours. After the time zone is changed, newly scaled-out nodes automatically apply the new time zone settings. Existing nodes are not affected. You can use the node reset feature in the node pool to apply the new settings to existing nodes.
	//
	// - For dedicated clusters, after the time zone is changed, newly scaled-out nodes, including control plane nodes, automatically apply the new time zone settings. Existing nodes, including control plane nodes, are not affected. You can use the node reset feature in the node pool to apply the new settings to existing nodes. For control plane nodes, you must scale out and then scale in the nodes to apply the new settings to all control plane nodes.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The vSwitches for the cluster control plane. For dedicated clusters, the change takes effect on newly scaled-out control plane nodes. For managed clusters, note the following when changing the control plane vSwitches:
	//
	// - This parameter overwrites the existing configuration. You must specify the complete list of target vSwitches.
	//
	// - Control plane components will be briefly restarted during the modification. Proceed with caution.
	//
	// - Make sure that the inbound and outbound rules of all security groups in the cluster, including the control plane security group, security groups of all node pools, and security groups used by the container network, allow the CIDR blocks of the new vSwitches. This prevents nodes and containers from failing to connect to the API server.
	//
	// - If ACL rules are configured for the new control plane vSwitches, make sure that the ACL rules allow communication with the CIDR blocks of cluster nodes and the container network.
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

func (s *ModifyClusterRequest) GetSupportPolicy() *string {
	return s.SupportPolicy
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

func (s *ModifyClusterRequest) SetSupportPolicy(v string) *ModifyClusterRequest {
	s.SupportPolicy = &v
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
	// Specifies whether to enable auto-renewal for the control plane node instances. This parameter takes effect only when `charge_type` is set to `PrePaid`. Valid values:
	//
	// - `true`: enables auto-renewal.
	//
	// - `false`: disables auto-renewal.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// The auto-renewal period for the control plane node instances.
	//
	// Valid values: 1, 2, 3, 6, and 12. Unit: months.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// The billing method of the control plane node instances. Valid values:
	//
	// - `PrePaid`: subscription
	//
	// - `PostPaid`: pay-as-you-go
	//
	// Default value: `PostPaid`.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"charge_type,omitempty" xml:"charge_type,omitempty"`
	// Specifies whether to install the CloudMonitor agent on the control plane nodes. Valid values:
	//
	// - `true`: installs the CloudMonitor agent.
	//
	// - `false`: does not install the CloudMonitor agent.
	//
	// example:
	//
	// true
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// The CPU management policy for nodes. The following two policies are supported when the cluster version is 1.12.6 or later:
	//
	// - `static`: allows pods with specific resource characteristics on the node to be granted enhanced CPU affinity and exclusivity.
	//
	// - `none`: indicates that the default CPU affinity scheme is used.
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
	// The custom image ID. Specify this parameter when you use a custom image.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20240819.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The type of the operating system image. Valid values:
	//
	// - `AliyunLinux3`: Alibaba Cloud Linux 3 image.
	//
	// - `Custom`: custom image.
	//
	// example:
	//
	// AliyunLinux3
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// The instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// The name of the key pair. You must specify this parameter or `login_password`.
	//
	// example:
	//
	// ack
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// The SSH logon password. You must specify this parameter or `key_pair`. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. If you want to use a password for logon, you must specify this parameter during scale-out.
	//
	// example:
	//
	// Ack@2000.
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The service port range of the node.
	//
	// Valid values: 30000 to 65535.
	//
	// Default value: 30000-32767.
	//
	// example:
	//
	// 30000-32767
	NodePortRange *string `json:"node_port_range,omitempty" xml:"node_port_range,omitempty"`
	// The subscription duration of the control plane node instances. This parameter is required and takes effect only when `charge_type` is set to `PrePaid`.
	//
	// When `period_unit` is set to `Month`, valid values are 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The billing cycle of the control plane node instances. This parameter takes effect only when `charge_type` is set to `PrePaid`.
	//
	// `Month`: months. Currently, only months are supported.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// The name of the container runtime. Valid values:
	//
	// - `containerd`: recommended. Supported by all cluster versions.
	//
	// Default value: containerd.
	//
	// example:
	//
	// containerd
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// Specifies whether to enable Alibaba Cloud OS security hardening. Valid values:
	//
	// - `true`: enables Alibaba Cloud OS security hardening.
	//
	// - `false`: disables Alibaba Cloud OS security hardening.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SecurityHardeningOs *bool `json:"security_hardening_os,omitempty" xml:"security_hardening_os,omitempty"`
	// The number of control plane nodes. If you want to scale out the control plane of a dedicated cluster, set this parameter to the target number of control plane nodes, which must be greater than the current number of control plane nodes.
	//
	// example:
	//
	// 5
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// Specifies whether to enable MLPS 2.0 security hardening. For more information, see [ACK MLPS 2.0 security hardening instructions](https://help.aliyun.com/document_detail/196148.html).
	//
	// Valid values:
	//
	// - `true`: enables MLPS 2.0 security hardening.
	//
	// - `false`: disables MLPS 2.0 security hardening.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// Specifies whether to enable the burst mode for the system cloud disk of the node. Valid values:
	//
	// - `true`: yes
	//
	// - `false`: no
	//
	// This parameter is supported only when `system_disk_category` is set to `cloud_auto`. For more information, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// The type of the system cloud disk. Valid values:
	//
	// - `cloud_efficiency`: ultra disk
	//
	// - `cloud_ssd`: standard SSD
	//
	// - `cloud_essd`: ESSD
	//
	// - `cloud_auto`: ESSD AutoPL cloud disk
	//
	// - `cloud_essd_entry`: ESSD Entry cloud disk
	//
	// example:
	//
	// cloud_essd
	SystemDiskCategory *string `json:"system_disk_category,omitempty" xml:"system_disk_category,omitempty"`
	// The performance level of the system cloud disk. This parameter takes effect only for ESSDs. The performance level is related to the disk size. For more information, see [ESSD](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// The provisioned read and write IOPS of the system cloud disk. Valid values: 0 to min{50,000, 1000 × capacity - baseline performance}. Baseline performance = min{1,800 + 50 × capacity, 50000}.
	//
	// This parameter is supported only when `system_disk_category` is set to `cloud_auto`. For more information, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// The size of the system cloud disk. Valid values: 40 to 500. Unit: GiB.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// The ID of the automatic snapshot policy applied to the system cloud disk of the node.
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
	// The internal DNS configuration for ACK managed clusters. Internal domain names are used by node-side system components such as kubelet and kube-proxy to access the API server. If internal DNS access is not enabled, node-side system components access the API server by using the CLB IP address.
	InternalDnsConfig *ModifyClusterRequestControlPlaneEndpointsConfigInternalDnsConfig `json:"internal_dns_config,omitempty" xml:"internal_dns_config,omitempty" type:"Struct"`
	// The load balancing configuration for cluster access.
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
	// The VPCs where internal DNS resolution takes effect.
	BindVpcs []*string `json:"bind_vpcs,omitempty" xml:"bind_vpcs,omitempty" type:"Repeated"`
	// Specifies whether to enable internal DNS access for the cluster. Valid values:
	//
	// - true: Enables internal DNS access. Node-side components such as kubelet and kube-proxy access the API server by using the internal domain name.
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
	// The ID of the Server Load Balancer (SLB) instance associated with the endpoint.
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
	// The automatic cluster upgrade configuration.
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
	// The channel for automatic cluster upgrades. For more information, see [Upgrade channels](https://help.aliyun.com/document_detail/2712866.html).
	//
	// Valid values:
	//
	// - patch: The latest patch version.
	//
	// - stable: The second latest minor version.
	//
	// - rapid: The latest minor version.
	//
	// example:
	//
	// patch
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// Specifies whether to enable automatic cluster upgrades. Valid values:
	//
	// - true: Enables automatic upgrades.
	//
	// - false: Disables automatic upgrades.
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
	// Specifies whether to enable system event logging. Valid values:
	//
	// - true: Enables system event logging.
	//
	// - false: Disables system event logging.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The name of the LogProject for system event logging.
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
