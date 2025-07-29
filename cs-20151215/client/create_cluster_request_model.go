// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessControlList(v []*string) *CreateClusterRequest
	GetAccessControlList() []*string
	SetAddons(v []*Addon) *CreateClusterRequest
	GetAddons() []*Addon
	SetApiAudiences(v string) *CreateClusterRequest
	GetApiAudiences() *string
	SetAuditLogConfig(v *CreateClusterRequestAuditLogConfig) *CreateClusterRequest
	GetAuditLogConfig() *CreateClusterRequestAuditLogConfig
	SetAutoMode(v *CreateClusterRequestAutoMode) *CreateClusterRequest
	GetAutoMode() *CreateClusterRequestAutoMode
	SetAutoRenew(v bool) *CreateClusterRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int64) *CreateClusterRequest
	GetAutoRenewPeriod() *int64
	SetChargeType(v string) *CreateClusterRequest
	GetChargeType() *string
	SetCisEnabled(v bool) *CreateClusterRequest
	GetCisEnabled() *bool
	SetCloudMonitorFlags(v bool) *CreateClusterRequest
	GetCloudMonitorFlags() *bool
	SetClusterDomain(v string) *CreateClusterRequest
	GetClusterDomain() *string
	SetClusterSpec(v string) *CreateClusterRequest
	GetClusterSpec() *string
	SetClusterType(v string) *CreateClusterRequest
	GetClusterType() *string
	SetContainerCidr(v string) *CreateClusterRequest
	GetContainerCidr() *string
	SetControlPlaneConfig(v *CreateClusterRequestControlPlaneConfig) *CreateClusterRequest
	GetControlPlaneConfig() *CreateClusterRequestControlPlaneConfig
	SetControlplaneLogComponents(v []*string) *CreateClusterRequest
	GetControlplaneLogComponents() []*string
	SetControlplaneLogProject(v string) *CreateClusterRequest
	GetControlplaneLogProject() *string
	SetControlplaneLogTtl(v string) *CreateClusterRequest
	GetControlplaneLogTtl() *string
	SetCpuPolicy(v string) *CreateClusterRequest
	GetCpuPolicy() *string
	SetCustomSan(v string) *CreateClusterRequest
	GetCustomSan() *string
	SetDeletionProtection(v bool) *CreateClusterRequest
	GetDeletionProtection() *bool
	SetDisableRollback(v bool) *CreateClusterRequest
	GetDisableRollback() *bool
	SetEnableRrsa(v bool) *CreateClusterRequest
	GetEnableRrsa() *bool
	SetEncryptionProviderKey(v string) *CreateClusterRequest
	GetEncryptionProviderKey() *string
	SetEndpointPublicAccess(v bool) *CreateClusterRequest
	GetEndpointPublicAccess() *bool
	SetExtraSans(v []*string) *CreateClusterRequest
	GetExtraSans() []*string
	SetFormatDisk(v bool) *CreateClusterRequest
	GetFormatDisk() *bool
	SetImageId(v string) *CreateClusterRequest
	GetImageId() *string
	SetImageType(v string) *CreateClusterRequest
	GetImageType() *string
	SetInstances(v []*string) *CreateClusterRequest
	GetInstances() []*string
	SetIpStack(v string) *CreateClusterRequest
	GetIpStack() *string
	SetIsEnterpriseSecurityGroup(v bool) *CreateClusterRequest
	GetIsEnterpriseSecurityGroup() *bool
	SetKeepInstanceName(v bool) *CreateClusterRequest
	GetKeepInstanceName() *bool
	SetKeyPair(v string) *CreateClusterRequest
	GetKeyPair() *string
	SetKubernetesVersion(v string) *CreateClusterRequest
	GetKubernetesVersion() *string
	SetLoadBalancerId(v string) *CreateClusterRequest
	GetLoadBalancerId() *string
	SetLoadBalancerSpec(v string) *CreateClusterRequest
	GetLoadBalancerSpec() *string
	SetLoggingType(v string) *CreateClusterRequest
	GetLoggingType() *string
	SetLoginPassword(v string) *CreateClusterRequest
	GetLoginPassword() *string
	SetMaintenanceWindow(v *MaintenanceWindow) *CreateClusterRequest
	GetMaintenanceWindow() *MaintenanceWindow
	SetMasterAutoRenew(v bool) *CreateClusterRequest
	GetMasterAutoRenew() *bool
	SetMasterAutoRenewPeriod(v int64) *CreateClusterRequest
	GetMasterAutoRenewPeriod() *int64
	SetMasterCount(v int64) *CreateClusterRequest
	GetMasterCount() *int64
	SetMasterInstanceChargeType(v string) *CreateClusterRequest
	GetMasterInstanceChargeType() *string
	SetMasterInstanceTypes(v []*string) *CreateClusterRequest
	GetMasterInstanceTypes() []*string
	SetMasterPeriod(v int64) *CreateClusterRequest
	GetMasterPeriod() *int64
	SetMasterPeriodUnit(v string) *CreateClusterRequest
	GetMasterPeriodUnit() *string
	SetMasterSystemDiskCategory(v string) *CreateClusterRequest
	GetMasterSystemDiskCategory() *string
	SetMasterSystemDiskPerformanceLevel(v string) *CreateClusterRequest
	GetMasterSystemDiskPerformanceLevel() *string
	SetMasterSystemDiskSize(v int64) *CreateClusterRequest
	GetMasterSystemDiskSize() *int64
	SetMasterSystemDiskSnapshotPolicyId(v string) *CreateClusterRequest
	GetMasterSystemDiskSnapshotPolicyId() *string
	SetMasterVswitchIds(v []*string) *CreateClusterRequest
	GetMasterVswitchIds() []*string
	SetName(v string) *CreateClusterRequest
	GetName() *string
	SetNatGateway(v bool) *CreateClusterRequest
	GetNatGateway() *bool
	SetNodeCidrMask(v string) *CreateClusterRequest
	GetNodeCidrMask() *string
	SetNodeNameMode(v string) *CreateClusterRequest
	GetNodeNameMode() *string
	SetNodePortRange(v string) *CreateClusterRequest
	GetNodePortRange() *string
	SetNodepools(v []*Nodepool) *CreateClusterRequest
	GetNodepools() []*Nodepool
	SetNumOfNodes(v int64) *CreateClusterRequest
	GetNumOfNodes() *int64
	SetOperationPolicy(v *CreateClusterRequestOperationPolicy) *CreateClusterRequest
	GetOperationPolicy() *CreateClusterRequestOperationPolicy
	SetOsType(v string) *CreateClusterRequest
	GetOsType() *string
	SetPeriod(v int64) *CreateClusterRequest
	GetPeriod() *int64
	SetPeriodUnit(v string) *CreateClusterRequest
	GetPeriodUnit() *string
	SetPlatform(v string) *CreateClusterRequest
	GetPlatform() *string
	SetPodVswitchIds(v []*string) *CreateClusterRequest
	GetPodVswitchIds() []*string
	SetProfile(v string) *CreateClusterRequest
	GetProfile() *string
	SetProxyMode(v string) *CreateClusterRequest
	GetProxyMode() *string
	SetRdsInstances(v []*string) *CreateClusterRequest
	GetRdsInstances() []*string
	SetRegionId(v string) *CreateClusterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateClusterRequest
	GetResourceGroupId() *string
	SetRrsaConfig(v *CreateClusterRequestRrsaConfig) *CreateClusterRequest
	GetRrsaConfig() *CreateClusterRequestRrsaConfig
	SetRuntime(v *Runtime) *CreateClusterRequest
	GetRuntime() *Runtime
	SetSecurityGroupId(v string) *CreateClusterRequest
	GetSecurityGroupId() *string
	SetSecurityHardeningOs(v bool) *CreateClusterRequest
	GetSecurityHardeningOs() *bool
	SetServiceAccountIssuer(v string) *CreateClusterRequest
	GetServiceAccountIssuer() *string
	SetServiceCidr(v string) *CreateClusterRequest
	GetServiceCidr() *string
	SetServiceDiscoveryTypes(v []*string) *CreateClusterRequest
	GetServiceDiscoveryTypes() []*string
	SetSnatEntry(v bool) *CreateClusterRequest
	GetSnatEntry() *bool
	SetSocEnabled(v bool) *CreateClusterRequest
	GetSocEnabled() *bool
	SetSshFlags(v bool) *CreateClusterRequest
	GetSshFlags() *bool
	SetTags(v []*Tag) *CreateClusterRequest
	GetTags() []*Tag
	SetTaints(v []*Taint) *CreateClusterRequest
	GetTaints() []*Taint
	SetTimeoutMins(v int64) *CreateClusterRequest
	GetTimeoutMins() *int64
	SetTimezone(v string) *CreateClusterRequest
	GetTimezone() *string
	SetUserCa(v string) *CreateClusterRequest
	GetUserCa() *string
	SetUserData(v string) *CreateClusterRequest
	GetUserData() *string
	SetVpcid(v string) *CreateClusterRequest
	GetVpcid() *string
	SetVswitchIds(v []*string) *CreateClusterRequest
	GetVswitchIds() []*string
	SetWorkerAutoRenew(v bool) *CreateClusterRequest
	GetWorkerAutoRenew() *bool
	SetWorkerAutoRenewPeriod(v int64) *CreateClusterRequest
	GetWorkerAutoRenewPeriod() *int64
	SetWorkerDataDisks(v []*CreateClusterRequestWorkerDataDisks) *CreateClusterRequest
	GetWorkerDataDisks() []*CreateClusterRequestWorkerDataDisks
	SetWorkerInstanceChargeType(v string) *CreateClusterRequest
	GetWorkerInstanceChargeType() *string
	SetWorkerInstanceTypes(v []*string) *CreateClusterRequest
	GetWorkerInstanceTypes() []*string
	SetWorkerPeriod(v int64) *CreateClusterRequest
	GetWorkerPeriod() *int64
	SetWorkerPeriodUnit(v string) *CreateClusterRequest
	GetWorkerPeriodUnit() *string
	SetWorkerSystemDiskCategory(v string) *CreateClusterRequest
	GetWorkerSystemDiskCategory() *string
	SetWorkerSystemDiskPerformanceLevel(v string) *CreateClusterRequest
	GetWorkerSystemDiskPerformanceLevel() *string
	SetWorkerSystemDiskSize(v int64) *CreateClusterRequest
	GetWorkerSystemDiskSize() *int64
	SetWorkerSystemDiskSnapshotPolicyId(v string) *CreateClusterRequest
	GetWorkerSystemDiskSnapshotPolicyId() *string
	SetWorkerVswitchIds(v []*string) *CreateClusterRequest
	GetWorkerVswitchIds() []*string
	SetZoneId(v string) *CreateClusterRequest
	GetZoneId() *string
	SetZoneIds(v []*string) *CreateClusterRequest
	GetZoneIds() []*string
}

type CreateClusterRequest struct {
	// The network access control list (ACL) rule of the SLB instance associated with the API server if the cluster is a registered cluster.
	AccessControlList []*string `json:"access_control_list,omitempty" xml:"access_control_list,omitempty" type:"Repeated"`
	// The components that you want to install in the cluster. When you create a cluster, you can configure the `addons` parameter to specify the components that you want to install.
	//
	// **Network plug-in**: required. The Flannel and Terway plug-ins are supported. Select one of the plug-ins for the cluster.
	//
	// 	- If you want to use the Terway component, specify the network plug-in in the [{"name":"flannel","config":""}] format.
	//
	// 	- If you want to use the Terway component, specify the value network plug-in in the [{"name": "terway-eniip","Config": ""}] format.
	//
	// **Volume plug-in**: optional. Only the `Container Storage Interface (CSI)` plug-in is supported.
	//
	// Specify the `CSI` plug-in in the following format: [{"name":"csi-plugin","config": ""},{"name": "csi-provisioner","config": ""}].
	//
	// **Simple Log Service component**: optional. We recommend that you enable Simple Log Service. If Simple Log Service is disabled, you cannot use the cluster auditing feature.
	//
	// 	- Specify an existing `Simple Log Service project` in the following format: [{"name": "logtail-ds","config": "{"IngressDashboardEnabled":"true","sls_project_name":"your_sls_project_name"}"}].
	//
	// 	- To create a `Simple Log Service project`, specify the component in the following format: [{"name": "logtail-ds","config": "{"IngressDashboardEnabled":"true"}"}].
	//
	// **Ingress controller**: optional. By default, the `nginx-ingress-controller` component is installed in ACK dedicated clusters.
	//
	// 	- To install nginx-ingress-controller and enable Internet access, specify the Ingress controller in the following format: [{"name":"nginx-ingress-controller","config":"{"IngressSlbNetworkType":"internet"}"}].
	//
	// 	- To disable the automatic installation of nginx-ingress-controller, specify the Ingress controller in the following format: [{"name": "nginx-ingress-controller","config": "","disabled": true}].
	//
	// **Event center**: optional. By default, the event center feature is enabled.
	//
	// You can use ACK event centers to store and query events and configure alerts. You can use the Logstores that are associated with ACK event centers free of charge within 90 days. For more information, see [Create and use an event center](https://help.aliyun.com/document_detail/150476.html).
	//
	// To enable the event center feature, specify the event center component in the following format: [{"name":"ack-node-problem-detector","config":"{"sls_project_name":"your_sls_project_name"}"}].
	Addons []*Addon `json:"addons,omitempty" xml:"addons,omitempty" type:"Repeated"`
	// Service accounts provide identities for pods when pods communicate with the `API server` of the cluster. The `api-audiences` parameter validates `tokens` and is used by the `API server` to check whether the `tokens` of requests are valid. Separate multiple values with commas (,).``
	//
	// For more information about `service accounts`, see [Enable service account token volume projection](https://help.aliyun.com/document_detail/160384.html).
	//
	// example:
	//
	// kubernetes.default.svc
	ApiAudiences   *string                             `json:"api_audiences,omitempty" xml:"api_audiences,omitempty"`
	AuditLogConfig *CreateClusterRequestAuditLogConfig `json:"audit_log_config,omitempty" xml:"audit_log_config,omitempty" type:"Struct"`
	AutoMode       *CreateClusterRequestAutoMode       `json:"auto_mode,omitempty" xml:"auto_mode,omitempty" type:"Struct"`
	// Deprecated
	//
	// [**Deprecated**]
	//
	// Specifies whether to enable auto-renewal. This parameter takes effect only when `charge_type` is set to `PrePaid`. Valid values:
	//
	// 	- `true`: enables auto-renewal.
	//
	// 	- `false`: disables auto-renewal.
	//
	// Default value: `false`.
	//
	// This parameter was changed on October 15, 2024. For more information, see [Announcement on changes to the parameter behavior of the CreateCluster operation](https://help.aliyun.com/document_detail/2849194.html).
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// Deprecated
	//
	// [**Deprecated**]
	//
	// The auto-renewal duration. This parameter takes effect only if charge_type is set to PrePaid and auto_renew is set to true. If you set `period_unit` to Month, the valid values of auto_renew_period are 1, 2, 3, 6, and 12.
	//
	// Default value: 1.
	//
	// This parameter was changed on October 15, 2024. For more information, see [Announcement on changes to the parameter behavior of the CreateCluster operation](https://help.aliyun.com/document_detail/2849194.html).
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// Deprecated
	//
	// [**Deprecated**]
	//
	// The billing method of the CLB instance that is used by the API server. Default value: PostPaid. Valid values:
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// 	- PrePaid: subscription. This billing method is not supported by newly created CLB instances. Existing CLB instances are not affected.
	//
	// >
	//
	// 	- This parameter was changed on October 15, 2024. For more information, see [Announcement on changes to the parameter behavior of the CreateCluster operation](https://help.aliyun.com/document_detail/2849194.html).
	//
	// 	- Starting from December 1, 2024, newly created CLB instances no longer support the subscription billing method, and an instance fee will be charged for newly created CLB instances
	//
	// For more information, see [CLB billing adjustments](https://help.aliyun.com/document_detail/2839797.html).
	//
	// example:
	//
	// 1
	ChargeType *string `json:"charge_type,omitempty" xml:"charge_type,omitempty"`
	// Deprecated
	//
	// [Deprecated] When you configure the control plane, use the `security_hardening_os` parameter in the `control_plane_config` section instead. When you configure a node pool, use the `security_hardening_os` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// example:
	//
	// false
	CisEnabled *bool `json:"cis_enabled,omitempty" xml:"cis_enabled,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `cloud_monitor_flags` parameter in the `control_plane_config` section instead. When you configure a node pool, use the `cms_enabled` parameter of the `kubernetes_config` field in the nodepool section instead.
	//
	// Specifies whether to install the CloudMonitor agent. Valid values:
	//
	// 	- `true`: installs the CloudMonitor agent.
	//
	// 	- `false`: does not install the CloudMonitor agent.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// The domain name of the cluster.
	//
	// The domain name can contain one or more parts that are separated by periods (.). Each part cannot exceed 63 characters in length, and can contain lowercase letters, digits, and hyphens (-). Each part must start and end with a lowercase letter or digit.
	//
	// example:
	//
	// cluster.local
	ClusterDomain *string `json:"cluster_domain,omitempty" xml:"cluster_domain,omitempty"`
	// If you set `cluster_type` to `ManagedKubernetes` and specify `profile`, you can further specify the edition of the cluster. Valid values:
	//
	// 	- `ack.pro.small`: Pro Edition.
	//
	// 	- `ack.standard`: Basic Edition. If you leave the parameter empty, an ACK Basic cluster is created.
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// 	- `Kubernetes`: ACK dedicated cluster.
	//
	// 	- `ManagedKubernetes`: ACK managed cluster. ACK managed clusters include ACK Basic clusters, ACK Pro clusters, ACK Serverless clusters (Basic Edition and Pro Edition), ACK Edge clusters (Basic Edition and Pro Edition), and ACK Lingjun clusters (Pro Edition).
	//
	// 	- `ExternalKubernetes`: registered cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// The pod CIDR block. You can specify 10.0.0.0/8, 172.16-31.0.0/12-16, 192.168.0.0/16, or their subnets as the pod CIDR block. The pod CIDR block cannot overlap with the CIDR block of the VPC in which the cluster is deployed and the CIDR blocks of existing clusters in the VPC. You cannot modify the pod CIDR block after you create the cluster.
	//
	// For more information about how to plan the network of an ACK cluster, see [Plan the network of an ACK cluster](https://help.aliyun.com/document_detail/86500.html).
	//
	// >  This parameter is required if the cluster uses the Flannel plug-in.
	//
	// example:
	//
	// 172.20.0.0/16
	ContainerCidr *string `json:"container_cidr,omitempty" xml:"container_cidr,omitempty"`
	// The control plane configurations of an ACK dedicated cluster.
	ControlPlaneConfig *CreateClusterRequestControlPlaneConfig `json:"control_plane_config,omitempty" xml:"control_plane_config,omitempty" type:"Struct"`
	// The control plane components for which you want to enable log collection.
	//
	// By default, the logs of kube-apiserver, kube-controller-manager, and kube-scheduler are collected.
	ControlplaneLogComponents []*string `json:"controlplane_log_components,omitempty" xml:"controlplane_log_components,omitempty" type:"Repeated"`
	// The Simple Log Service project that is used to store the logs of control plane components. You can use an existing project or create one. If you choose to create a Simple Log Service project, the created project is named in the `k8s-log-{ClusterID}` format.
	//
	// example:
	//
	// k8s-log-xxx
	ControlplaneLogProject *string `json:"controlplane_log_project,omitempty" xml:"controlplane_log_project,omitempty"`
	// The retention period of control plane logs in days.
	//
	// example:
	//
	// 30
	ControlplaneLogTtl *string `json:"controlplane_log_ttl,omitempty" xml:"controlplane_log_ttl,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `cpu_policy` parameter in the `control_plane_config` section instead. When you configure a node pool, use the `cpu_policy` parameter of the `kubernetes_config` field in the `nodepool` section instead.
	//
	// The CPU management policy of the node. The following policies are supported if the Kubernetes version of the cluster is 1.12.6 or later:
	//
	// 	- `static`: allows pods with specific resource characteristics on the node to be granted enhanced CPU affinity and exclusivity.
	//
	// 	- `none`: specifies that the default CPU affinity is used.
	//
	// Default value: `none`.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// Deprecated
	//
	// The custom subject alternative names (SANs) for the API server certificate to accept requests from specified IP addresses or domain names. Separate multiple IP addresses and domain names with commas (,).
	//
	// example:
	//
	// cs.aliyun.com
	CustomSan *string `json:"custom_san,omitempty" xml:"custom_san,omitempty"`
	// Specifies whether to enable cluster deletion protection. If you enable this option, the cluster cannot be deleted in the console or by calling API operations. Valid values:
	//
	// 	- `true`: enables cluster deletion protection.
	//
	// 	- `false`: disables cluster deletion protection.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] By default, the system does not perform a rollback when the cluster fails to be created. You must manually delete the cluster that fails to be created.
	//
	// Specifies whether to perform a rollback when the cluster fails to be created. Valid values:
	//
	// 	- `true`: performs a rollback when the cluster fails to be created.
	//
	// 	- `false`: does not perform a rollback when the cluster fails to be created.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	DisableRollback *bool `json:"disable_rollback,omitempty" xml:"disable_rollback,omitempty"`
	// Deprecated
	//
	// Specifies whether to enable the RAM Roles for Service Accounts (RRSA) feature.
	//
	// example:
	//
	// true
	EnableRrsa *bool `json:"enable_rrsa,omitempty" xml:"enable_rrsa,omitempty"`
	// The ID of the Key Management Service (KMS) key that is used to encrypt the system disk. For more information, see [What is KMS?](https://help.aliyun.com/document_detail/28935.html)
	//
	// >  The key can be used only in ACK Pro clusters.
	//
	// example:
	//
	// 0fe64791-55eb-4fc7-84c5-c6c7cdca****
	EncryptionProviderKey *string `json:"encryption_provider_key,omitempty" xml:"encryption_provider_key,omitempty"`
	// Specifies whether to enable Internet access for the cluster. You can use an elastic IP address (EIP) to expose the API server. This way, you can access the cluster over the Internet. Valid values:
	//
	// 	- `true`: enables Internet access for the cluster.
	//
	// 	- `false`: disables Internet access for the cluster. If you set the value to false, the API server cannot be accessed over the Internet.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	EndpointPublicAccess *bool     `json:"endpoint_public_access,omitempty" xml:"endpoint_public_access,omitempty"`
	ExtraSans            []*string `json:"extra_sans,omitempty" xml:"extra_sans,omitempty" type:"Repeated"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, you cannot add existing nodes to the cluster. If you want to add existing nodes, you must first create a node pool and then call the [AttachInstancesToNodePool](https://help.aliyun.com/document_detail/2667920.html) operation.
	//
	// Specifies whether to mount a data disk to a node that is created based on an existing ECS instance. Valid values:
	//
	// 	- `true`: stores the data of containers and images on a data disk. The existing data stored in the data disk is lost. Back up the existing data first.
	//
	// 	- `false`: does not store the data of containers and images on a data disk.
	//
	// Default value: `false`.
	//
	// How data disks are mounted:
	//
	// 	- If an ECS instance has data disks mounted and the file system of the last data disk is not initialized, the system automatically formats the data disk to ext4. Then, the system mounts the data disk to /var/lib/docker and /var/lib/kubelet.
	//
	// 	- If no data disk is mounted to the ECS instance, the system does not purchase a new data disk.
	//
	// example:
	//
	// false
	FormatDisk *bool `json:"format_disk,omitempty" xml:"format_disk,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `image_id` parameter in the `control_plane_config` section instead. When you configure a node pool, use the `image_id` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The custom image for nodes. By default, the image provided by ACK is used. You can select a custom image to replace the default image. For more information, see [Use a custom image to create an ACK cluster](https://help.aliyun.com/document_detail/146647.html).
	//
	// example:
	//
	// m-bp16z7xko3vvv8gt****
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `image_type` parameter in the `control_plane_config` section instead. When you configure a node pool, use the `image_type` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The type of OS distribution that you want to use. To specify the node OS, we recommend that you use this parameter. Valid values:
	//
	// 	- CentOS
	//
	// 	- AliyunLinux
	//
	// 	- AliyunLinux Qboot
	//
	// 	- AliyunLinuxUEFI
	//
	// 	- AliyunLinux3
	//
	// 	- Windows
	//
	// 	- WindowsCore
	//
	// 	- AliyunLinux3Arm64
	//
	// 	- ContainerOS
	//
	// Default value: `CentOS`.
	//
	// example:
	//
	// AliyunLinux
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, you cannot add existing nodes to the cluster. If you want to add existing nodes, you must first create a node pool and then call the [AttachInstancesToNodePool](https://help.aliyun.com/document_detail/2667920.html) operation.
	//
	// The existing ECS instances that are specified as worker nodes for the cluster.
	//
	// >  This parameter is required if you create worker nodes on existing ECS instances.
	Instances []*string `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// The IP stack of the cluster.
	//
	// example:
	//
	// Optional value: ipv4 (Single stack) or ipv6 (Dual Stack)
	//
	// Default value: IPV4
	IpStack *string `json:"ip_stack,omitempty" xml:"ip_stack,omitempty"`
	// Specifies whether to create an advanced security group. This parameter takes effect only if `security_group_id` is left empty.
	//
	// >  To use a basic security group, make sure that the sum of the number of nodes in the cluster and the number of pods that use Terway does not exceed 2,000. Therefore, we recommend that you specify an advanced security group for a cluster that uses Terway.
	//
	// 	- `true`: creates an advanced security group.
	//
	// 	- `false`: does not create an advanced security group.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	IsEnterpriseSecurityGroup *bool `json:"is_enterprise_security_group,omitempty" xml:"is_enterprise_security_group,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, you cannot add existing nodes to the cluster. If you want to add existing nodes, you must first create a node pool and then call the [AttachInstancesToNodePool](https://help.aliyun.com/document_detail/2667920.html) operation.
	//
	// Specifies whether to retain the names of existing ECS instances that are used in the cluster. Valid values:
	//
	// 	- `true`: retains the names.
	//
	// 	- `false`: does not retain the names. The system assigns new names.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	KeepInstanceName *bool `json:"keep_instance_name,omitempty" xml:"keep_instance_name,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `key_pair` parameter in the `control_plane_config` section instead. When you configure a node pool, use the `key_pair` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The name of the key pair. You must configure this parameter or `login_password`.
	//
	// example:
	//
	// secrity-key
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// The Kubernetes version of the cluster. The Kubernetes versions supported by ACK are the same as the Kubernetes versions supported by open source Kubernetes. We recommend that you specify the latest Kubernetes version. If you do not specify this parameter, the latest Kubernetes version is used.
	//
	// You can create ACK clusters of the latest three Kubernetes versions in the ACK console. If you want to create clusters that run earlier Kubernetes versions, use the ACK API. For more information about the Kubernetes versions supported by ACK, see [Support for Kubernetes versions](https://help.aliyun.com/document_detail/185269.html).
	//
	// example:
	//
	// 1.16.9-aliyun.1
	KubernetesVersion *string `json:"kubernetes_version,omitempty" xml:"kubernetes_version,omitempty"`
	// Specifies the ID of the CLB instance for accessing the API server. If this parameter is specified, the system does not automatically create a CLB instance for the API server.
	//
	// >  Make sure that the CLB instance does not have other dependencies, such as listeners and backend servers. You cannot specify shared-resource or Internet-facing CLB instances.
	//
	// example:
	//
	// lb-wz9t256gqa3vbouk****
	LoadBalancerId *string `json:"load_balancer_id,omitempty" xml:"load_balancer_id,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] The pay-as-you-go billing method is used by Classic Load Balancer (CLB) instances. This parameter does not take effect.
	//
	// The specification of the Server Load Balancer (SLB) instance. Valid values:
	//
	// 	- slb.s1.small
	//
	// 	- slb.s2.small
	//
	// 	- slb.s2.medium
	//
	// 	- slb.s3.small
	//
	// 	- slb.s3.medium
	//
	// 	- slb.s3.large
	//
	// Default value: `slb.s2.small`.
	//
	// example:
	//
	// slb.s2.small
	LoadBalancerSpec *string `json:"load_balancer_spec,omitempty" xml:"load_balancer_spec,omitempty"`
	// Enables Simple Log Service for the cluster. This parameter takes effect only for ACK Serverless clusters. Set the value to `SLS`.
	//
	// example:
	//
	// SLS
	LoggingType *string `json:"logging_type,omitempty" xml:"logging_type,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `login_password` parameter in the `control_plane_config` section instead. When you configure a node pool, use the `login_password` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The password for SSH logon. You must set this parameter or `key_pair`. The password must be 8 to 30 characters in length, and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// example:
	//
	// Hello@1234
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The configurations of the cluster maintenance window.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window,omitempty" xml:"maintenance_window,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `auto-renew` parameter in the `control_plane_config` section instead.
	//
	// Specifies whether to enable auto-renewal for master nodes. This parameter takes effect only when `master_instance_charge_type` is set to `PrePaid`. Valid values:
	//
	// 	- `true`: enables auto-renewal.
	//
	// 	- `false`: disables auto-renewal.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	MasterAutoRenew *bool `json:"master_auto_renew,omitempty" xml:"master_auto_renew,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `auto-renew_period` parameter in the `control_plane_config` section instead.
	//
	// The auto-renewal duration. This parameter takes effect and is required only when the subscription billing method is selected for master nodes.
	//
	// Valid values: 1, 2, 3, 6, and 12.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	MasterAutoRenewPeriod *int64 `json:"master_auto_renew_period,omitempty" xml:"master_auto_renew_period,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `size` parameter in the `control_plane_config` section instead.
	//
	// The number of master nodes. Valid values: `3` and `5`.
	//
	// Default value: `3`.
	//
	// example:
	//
	// 3
	MasterCount *int64 `json:"master_count,omitempty" xml:"master_count,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `instance_charge_type` parameter in the `control_plane_config` section instead.
	//
	// The billing method of master nodes. Valid values:
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
	MasterInstanceChargeType *string `json:"master_instance_charge_type,omitempty" xml:"master_instance_charge_type,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `instance_types` parameter in the `control_plane_config` section instead.
	//
	// The instance types of master nodes. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	MasterInstanceTypes []*string `json:"master_instance_types,omitempty" xml:"master_instance_types,omitempty" type:"Repeated"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `unit` parameter in the `control_plane_config` section instead.
	//
	// The subscription duration of master nodes. This parameter takes effect and is required only when `master_instance_charge_type` is set to `PrePaid`.
	//
	// Valid values: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	MasterPeriod *int64 `json:"master_period,omitempty" xml:"master_period,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `period_unit` parameter in the `control_plane_config` section instead.
	//
	// The billing cycle of the master nodes in the cluster. This parameter is required if master_instance_charge_type is set to `PrePaid`.
	//
	// Valid value: `Month`, which indicates that master nodes are billed only on a monthly basis.
	//
	// example:
	//
	// Month
	MasterPeriodUnit *string `json:"master_period_unit,omitempty" xml:"master_period_unit,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `system_disk_category` parameter in the `control_plane_config` section instead.
	//
	// The system disk category of master nodes. Valid values:
	//
	// 	- `cloud_efficiency`: ultra disk.
	//
	// 	- `cloud_ssd`: standard SSD.
	//
	// 	- `cloud_essd`: Enterprise SSD (ESSD).
	//
	// Default value: `cloud_ssd`. The default value may vary in different zones.
	//
	// example:
	//
	// cloud_ssd
	MasterSystemDiskCategory *string `json:"master_system_disk_category,omitempty" xml:"master_system_disk_category,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `system_disk_performance_level` parameter in the `control_plane_config` section instead.
	//
	// The performance level (PL) of the system disk that you want to use for master nodes. This parameter takes effect only for ESSDs. For more information about the relationship between disk PLs and disk sizes, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	MasterSystemDiskPerformanceLevel *string `json:"master_system_disk_performance_level,omitempty" xml:"master_system_disk_performance_level,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `system_disk_disk` parameter in the `control_plane_config` section instead.
	//
	// The system disk size of master nodes. Valid values: 40 to 500. Unit: GiB.
	//
	// Default value: `120`.
	//
	// example:
	//
	// 120
	MasterSystemDiskSize *int64 `json:"master_system_disk_size,omitempty" xml:"master_system_disk_size,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `system_disk_snapshot_policy_id` parameter in the `control_plane_config` section instead.
	//
	// The ID of the automatic snapshot policy that is used by the system disk specified for master nodes.
	//
	// example:
	//
	// sp-2zej1nogjvovnz4z****
	MasterSystemDiskSnapshotPolicyId *string `json:"master_system_disk_snapshot_policy_id,omitempty" xml:"master_system_disk_snapshot_policy_id,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] Use the `vswitch_ids` parameter instead.
	//
	// The IDs of the vSwitches that are specified for master nodes. You can specify up to three vSwitches. We recommend that you specify three vSwitches in different zones to ensure high availability.
	//
	// The number of vSwitches must be the same as the value of the `master_count` parameter and also the same as the number of vSwitches specified in the `master_vswitch_ids` parameter.
	MasterVswitchIds []*string `json:"master_vswitch_ids,omitempty" xml:"master_vswitch_ids,omitempty" type:"Repeated"`
	// The cluster name.
	//
	// The name must be 1 to 63 characters in length, and can contain digits, letters, and hyphens (-). The name cannot start with a hyphen (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster-demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// [Deprecated] Use the `snat_entry` parameter instead.
	//
	// example:
	//
	// true
	NatGateway *bool `json:"nat_gateway,omitempty" xml:"nat_gateway,omitempty"`
	// The maximum number of IP addresses that can be assigned to each node. This number is determined by the subnet mask of the specified CIDR block. This parameter takes effect only if the cluster uses the Flannel plug-in.
	//
	// Default value: `26`.
	//
	// example:
	//
	// 25
	NodeCidrMask *string `json:"node_cidr_mask,omitempty" xml:"node_cidr_mask,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, use the `node_name_mode` parameter of the `kubernetes_config` field in the `nodepool` section instead.
	//
	// The custom node name.
	//
	// A custom node name consists of a prefix, a node IP address, and a suffix.
	//
	// 	- The prefix and suffix can contain multiple parts that are separated by periods (.). Each part can contain lowercase letters, digits, and hyphens (-), and must start and end with a lowercase letter or digit.
	//
	// 	- The IP substring length specifies the number of digits to be truncated from the end of the node IP address. The IP substring length ranges from 5 to 12.
	//
	// For example, if the node IP address is 192.168.0.55, the prefix is aliyun.com, the IP substring length is 5, and the suffix is test, the node name will aliyun.com00055test.
	//
	// example:
	//
	// aliyun.com00055test
	NodeNameMode *string `json:"node_name_mode,omitempty" xml:"node_name_mode,omitempty"`
	// The node port range. Valid values: 30000 to 65535.
	//
	// Default value: `30000-32767`.
	//
	// example:
	//
	// 30000~32767
	NodePortRange *string `json:"node_port_range,omitempty" xml:"node_port_range,omitempty"`
	// The list of node pools.
	Nodepools []*Nodepool `json:"nodepools,omitempty" xml:"nodepools,omitempty" type:"Repeated"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, use the `desired_size` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The number of worker nodes. Valid values: 0 to 100.
	//
	// example:
	//
	// 3
	NumOfNodes *int64 `json:"num_of_nodes,omitempty" xml:"num_of_nodes,omitempty"`
	// The automatic O\\&M policy of the cluster.
	OperationPolicy *CreateClusterRequestOperationPolicy `json:"operation_policy,omitempty" xml:"operation_policy,omitempty" type:"Struct"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `image_type` parameter in the `control_plane_config` section instead. When you configure a node pool, use the `image_type` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The type of OS. Valid values:
	//
	// 	- Windows
	//
	// 	- Linux
	//
	// Default value: `Linux`.
	//
	// example:
	//
	// Linux
	OsType *string `json:"os_type,omitempty" xml:"os_type,omitempty"`
	// Deprecated
	//
	// [**Deprecated**]
	//
	// The subscription duration. This parameter takes effect and is required only when you set charge_type to PrePaid.
	//
	// Valid values: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// Default value: 1.
	//
	// This parameter was changed on October 15, 2024. For more information, see [Announcement on changes to the parameter behavior of the CreateCluster operation](https://help.aliyun.com/document_detail/2849194.html).
	//
	// example:
	//
	// FY2023
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// Deprecated
	//
	// [**Deprecated**]
	//
	// The billing cycle. This parameter is required if charge_type is set to PrePaid.
	//
	// Valid value: Month, which indicates that resources are billed only on a monthly basis.
	//
	// This parameter was changed on October 15, 2024. For more information, see [Announcement on changes to the parameter behavior of the CreateCluster operation](https://help.aliyun.com/document_detail/2849194.html).
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, use the `platform` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The OS distribution that is used. Valid values:
	//
	// 	- CentOS
	//
	// 	- AliyunLinux
	//
	// 	- QbootAliyunLinux
	//
	// 	- Qboot
	//
	// 	- Windows
	//
	// 	- WindowsCore
	//
	// Default value: `CentOS`.
	//
	// example:
	//
	// CentOS
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// If you select Terway as the network plug-in, you must allocate vSwitches to pods. For each vSwitch that allocates IP addresses to worker nodes, you must select a vSwitch in the same zone to allocate IP addresses to pods.
	//
	// >  We recommend that you select pod vSwitches whose subnet masks do not exceed 19 bits in length. The maximum subnet mask length of a pod vSwitch is 25 bits. If you select a pod vSwitch whose subnet mask exceeds 25 bits in length, the IP addresses that can be allocated to pods may be insufficient.
	PodVswitchIds []*string `json:"pod_vswitch_ids,omitempty" xml:"pod_vswitch_ids,omitempty" type:"Repeated"`
	// If you set `cluster_type` to `ManagedKubernetes`, an ACK managed cluster is created. In this case, you can further specify the cluster edition. Valid values:
	//
	// 	- `Default`: ACK managed cluster. ACK managed clusters include ACK Basic clusters and ACK Pro clusters.
	//
	// 	- `Edge`: ACK Edge cluster. ACK Edge clusters include ACK Edge Basic clusters and ACK Edge Pro clusters.
	//
	// 	- `Serverless`: ACK Serverless cluster. ACK Serverless clusters include ACK Serverless Basic clusters and ACK Serverless Pro clusters.
	//
	// 	- `Lingjun`: ACK Lingjun Pro cluster.
	//
	// example:
	//
	// Default
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
	// The kube-proxy mode. Valid values:
	//
	// 	- `iptables`: a mature and stable mode that uses iptables rules to conduct service discovery and load balancing. The performance of this mode is limited by the size of the cluster. This mode is suitable for clusters that run a small number of Services.
	//
	// 	- `ipvs`: a mode that provides high performance and uses IP Virtual Server (IPVS) to conduct service discovery and load balancing. This mode is suitable for clusters that run a large number of Services. We recommend that you use this mode in scenarios that require high-performance load balancing.
	//
	// Default value: `ipvs`.
	//
	// example:
	//
	// ipvs
	ProxyMode *string `json:"proxy_mode,omitempty" xml:"proxy_mode,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, use the `rds_instances` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The ApsaraDB RDS instances. The pod CIDR block and node CIDR block are added to the whitelists of the ApsaraDB RDS instances. We recommend that you add the pod CIDR block and node CIDR block to the whitelists of the ApsaraDB RDS instances in the ApsaraDB RDS console. If the RDS instances are not in the Running state, new nodes cannot be added to the cluster.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// The ID of the region in which the cluster is deployed. For more information about the regions supported by ACK, see [Regions supported by ACK](https://help.aliyun.com/document_detail/216938.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// The ID of the resource group to which the cluster belongs. You can use resource groups to isolate clusters.
	//
	// example:
	//
	// rg-acfm3mkrure****
	ResourceGroupId *string                         `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	RrsaConfig      *CreateClusterRequestRrsaConfig `json:"rrsa_config,omitempty" xml:"rrsa_config,omitempty" type:"Struct"`
	// The container runtime. The default container runtime is Docker. containerd and Sandboxed-Container are also supported.
	//
	// For more information about how to select a proper container runtime, see [Comparison among Docker, containerd, and Sandboxed-Container](https://help.aliyun.com/document_detail/160313.html).
	Runtime *Runtime `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The ID of an existing security group. You must specify this parameter or `is_enterprise_security_group`. Cluster nodes are automatically added to the security group.
	//
	// example:
	//
	// sg-bp1bdue0qc1g7k****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `security_hardening_os` parameter in the `control_plane_config` section instead. When you configure a node pool, use the `security_hardening_os` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// Specifies whether to enable Alibaba Cloud Linux Security Hardening. Valid values:
	//
	// 	- `true`: enables Alibaba Cloud Linux Security Hardening.
	//
	// 	- `false`: disables Alibaba Cloud Linux Security Hardening.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SecurityHardeningOs *bool `json:"security_hardening_os,omitempty" xml:"security_hardening_os,omitempty"`
	// Service accounts provide identities for pods when pods communicate with the `API server` of the cluster. `service-account-issuer` specifies the issuer of the `serviceaccount token`, which is specified by using the `iss` field in the `token payload`.
	//
	// For more information about `ServiceAccount`, see [Enable service account token volume projection](https://help.aliyun.com/document_detail/160384.html).
	//
	// example:
	//
	// kubernetes.default.svc
	ServiceAccountIssuer *string `json:"service_account_issuer,omitempty" xml:"service_account_issuer,omitempty"`
	// The Service CIDR block. Valid values: 10.0.0.0/16-24, 172.16-31.0.0/16-24, and 192.168.0.0/16-24. The Service CIDR block cannot overlap with the VPC CIDR block (10.1.0.0/21) or the CIDR blocks of existing clusters in the VPC. You cannot modify the Service CIDR block after the cluster is created.
	//
	// By default, the Service CIDR block is set to 172.19.0.0/20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.21.0.0/20
	ServiceCidr *string `json:"service_cidr,omitempty" xml:"service_cidr,omitempty"`
	// The methods for implementing service discovery in `ACK Serverless` clusters.
	//
	// 	- `CoreDNS`: a standard service discovery plug-in that is provided by open source Kubernetes. To use DNS resolution, you must provision pods. By default, two elastic container instances are used. The specification of each instance is 0.25 vCores and 512 MiB of memory.
	//
	// 	- `PrivateZone`: a DNS resolution service provided by Alibaba Cloud. You must activate Alibaba Cloud DNS PrivateZone before you can use it for service discovery.
	//
	// By default, this parameter is not specified.
	ServiceDiscoveryTypes []*string `json:"service_discovery_types,omitempty" xml:"service_discovery_types,omitempty" type:"Repeated"`
	// Specifies whether to configure SNAT rules for the VPC in which your cluster is deployed. Valid values:
	//
	// 	- `true`: automatically creates a NAT gateway and configures SNAT rules. Set the value to `true` if nodes and applications in the cluster need to access the Internet.
	//
	// 	- `false`: does not create a NAT gateway or configure SNAT rules. In this case, nodes and applications in the cluster cannot access the Internet.
	//
	// >  If this feature is disabled when you create the cluster, you can also manually enable this feature after you create the cluster. For more information, see [Enable an existing ACK cluster to access the Internet](https://help.aliyun.com/document_detail/178480.html).
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	SnatEntry *bool `json:"snat_entry,omitempty" xml:"snat_entry,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure the control plane, use the `soc_enabled` parameter in the `control_plane_config` section instead. When you configure a node pool, use the `soc_enabled` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
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
	// Specifies whether to enable SSH logon. If this parameter is set to true, you can log on to master nodes in an ACK dedicated cluster over the Internet. This parameter does not take effect for ACK managed clusters. Valid values:
	//
	// 	- `true`: enables SSH logon.
	//
	// 	- `false`: disables SSH logon.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	SshFlags *bool `json:"ssh_flags,omitempty" xml:"ssh_flags,omitempty"`
	// The labels that you want to add to nodes. You must add labels based on the following rules:
	//
	// 	- A label is a case-sensitive key-value pair. You can add up to 20 labels.
	//
	// 	- When you add a label, you must specify a unique key, but you can leave the value empty. A key cannot exceed 64 characters in length, and a value cannot exceed 128 characters in length. Keys and values cannot start with aliyun, acs:, https://, or http://. For more information, see [Labels and Selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set).
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, use the `taints` parameter of the `kubernetes_config` field in the `nodepool` section instead.
	//
	// The taints that you want to add to nodes. Taints can be used together with tolerations to avoid scheduling pods to specific nodes. For more information, see [taint-and-toleration](https://kubernetes.io/zh/docs/concepts/scheduling-eviction/taint-and-toleration/).
	Taints []*Taint `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// Deprecated
	//
	// [**Deprecated**] By default, the system does not perform a rollback when the cluster fails to be created. You must manually delete the cluster that fails to be created.
	//
	// Specifies the timeout period of cluster creation. Unit: minutes.
	//
	// Default value: `60`.
	//
	// example:
	//
	// 60
	TimeoutMins *int64 `json:"timeout_mins,omitempty" xml:"timeout_mins,omitempty"`
	// The time zone of the cluster.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The custom Certificate Authority (CA) certificate used by the cluster.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----****
	UserCa *string `json:"user_ca,omitempty" xml:"user_ca,omitempty"`
	// The user data of nodes.
	//
	// example:
	//
	// IyEvdXNyL2Jpbi9iYXNoCmVjaG8gIkhlbGxvIEFD****
	UserData *string `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// The virtual private cloud (VPC) in which you want to deploy the cluster. This parameter is required.
	//
	// example:
	//
	// vpc-2zeik9h3ahvv2zz95****
	Vpcid *string `json:"vpcid,omitempty" xml:"vpcid,omitempty"`
	// The vSwitches for nodes in the cluster. This parameter is required if you create an ACK managed cluster that does not contain nodes.
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, use the `auto_renew` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// Specifies whether to enable auto-renewal for worker nodes. This parameter takes effect and is required only when `worker_instance_charge_type` is set to `PrePaid`. Valid values:
	//
	// 	- `true`: enables auto-renewal.
	//
	// 	- `false`: disables auto-renewal.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	WorkerAutoRenew *bool `json:"worker_auto_renew,omitempty" xml:"worker_auto_renew,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, use the `auto_renew_period` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The auto-renewal duration of worker nodes. This parameter takes effect and is required only if the subscription billing method is selected for worker nodes.
	//
	// Valid values: 1, 2, 3, 6, and 12.
	//
	// example:
	//
	// 1
	WorkerAutoRenewPeriod *int64 `json:"worker_auto_renew_period,omitempty" xml:"worker_auto_renew_period,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, use the `data_disks` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The configurations of the data disks that you want to mount to worker nodes. The configurations include the disk category and disk size.
	WorkerDataDisks []*CreateClusterRequestWorkerDataDisks `json:"worker_data_disks,omitempty" xml:"worker_data_disks,omitempty" type:"Repeated"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, use the `instance_charge_type` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The billing method of worker nodes. Valid values:
	//
	// 	- `PrePaid`: subscription.
	//
	// 	- `PostPaid`: pay-as-you-go.
	//
	// Default value: PostPaid.
	//
	// example:
	//
	// PrePaid
	WorkerInstanceChargeType *string `json:"worker_instance_charge_type,omitempty" xml:"worker_instance_charge_type,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, use the `instance_types` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The instance configurations of worker nodes.
	WorkerInstanceTypes []*string `json:"worker_instance_types,omitempty" xml:"worker_instance_types,omitempty" type:"Repeated"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, use the `period` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The subscription duration of worker nodes. This parameter takes effect and is required only when `worker_instance_charge_type` is set to `PrePaid`.
	//
	// Valid values: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	WorkerPeriod *int64 `json:"worker_period,omitempty" xml:"worker_period,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, use the `period_unit` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The billing cycle of worker nodes. This parameter is required if worker_instance_charge_type is set to `PrePaid`.
	//
	// Valid value: `Month`, which indicates that worker nodes are billed only on a monthly basis.
	//
	// example:
	//
	// Month
	WorkerPeriodUnit *string `json:"worker_period_unit,omitempty" xml:"worker_period_unit,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, use the `system_disk_category` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The system disk category of worker nodes. For more information, see [Elastic Block Storage devices](https://help.aliyun.com/document_detail/63136.html).
	//
	// Valid values:
	//
	// 	- `cloud_efficiency`: ultra disk.
	//
	// 	- `cloud_ssd`: standard SSD.
	//
	// Default value: `cloud_ssd`.
	//
	// example:
	//
	// cloud_efficiency
	WorkerSystemDiskCategory *string `json:"worker_system_disk_category,omitempty" xml:"worker_system_disk_category,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, use the `system_disk_performance_level` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// If the system disk is an ESSD, you can specify the PL of the ESSD. For more information, see [Enterprise SSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// Valid values:
	//
	// 	- PL0
	//
	// 	- PL1
	//
	// 	- PL2
	//
	// 	- PL3
	//
	// example:
	//
	// PL1
	WorkerSystemDiskPerformanceLevel *string `json:"worker_system_disk_performance_level,omitempty" xml:"worker_system_disk_performance_level,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, use the `system_disk_size` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The system disk size of worker nodes. Unit: GiB.
	//
	// Valid values: 40 to 500.
	//
	// The value of this parameter must be at least 40 and greater than or equal to the image size.
	//
	// Default value: `120`.
	//
	// example:
	//
	// 120
	WorkerSystemDiskSize *int64 `json:"worker_system_disk_size,omitempty" xml:"worker_system_disk_size,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, use the `system_disk_snapshot_policy_id` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The ID of the automatic snapshot policy that is used by the system disk specified for worker nodes.
	//
	// example:
	//
	// sp-2zej1nogjvovnz4z****
	WorkerSystemDiskSnapshotPolicyId *string `json:"worker_system_disk_snapshot_policy_id,omitempty" xml:"worker_system_disk_snapshot_policy_id,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] When you configure a node pool, use the `vswitch_ids` parameter of the `scaling_group` field in the `nodepool` section instead.
	//
	// The vSwitches for worker nodes. Each worker node is allocated a vSwitch.
	//
	// `worker_vswitch_ids` is optional, but `vswitch_ids` is required if you create an ACK managed cluster that does not contain nodes.
	WorkerVswitchIds []*string `json:"worker_vswitch_ids,omitempty" xml:"worker_vswitch_ids,omitempty" type:"Repeated"`
	// Deprecated
	//
	// [Deprecated] Use the `zone_ids` parameter instead.
	//
	// The ID of the zone to which the cluster belongs. This parameter is specific to ACK Serverless clusters.
	//
	// When you create an ACK managed cluster, you must set the `zone_id` parameter if `vpc_id` and `vswitch_ids` are not specified. This way, the system automatically creates a VPC in the specified zone. This parameter is invalid if you specify the `vpc_id` and `vswitch_ids` parameters.
	//
	// example:
	//
	// cn-beiji****
	ZoneId *string `json:"zone_id,omitempty" xml:"zone_id,omitempty"`
	// The IDs of the zone in which the cluster is deployed. This parameter is specific to ACK managed clusters.
	//
	// When you create an ACK managed cluster, you must set the `zone_id` parameter if `vpc_id` and `vswitch_ids` are not specified. This way, the system automatically creates a VPC in the specified zone. This parameter is invalid if you specify the `vpc_id` and `vswitch_ids` parameters.
	ZoneIds []*string `json:"zone_ids,omitempty" xml:"zone_ids,omitempty" type:"Repeated"`
}

func (s CreateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) GetAccessControlList() []*string {
	return s.AccessControlList
}

func (s *CreateClusterRequest) GetAddons() []*Addon {
	return s.Addons
}

func (s *CreateClusterRequest) GetApiAudiences() *string {
	return s.ApiAudiences
}

func (s *CreateClusterRequest) GetAuditLogConfig() *CreateClusterRequestAuditLogConfig {
	return s.AuditLogConfig
}

func (s *CreateClusterRequest) GetAutoMode() *CreateClusterRequestAutoMode {
	return s.AutoMode
}

func (s *CreateClusterRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateClusterRequest) GetAutoRenewPeriod() *int64 {
	return s.AutoRenewPeriod
}

func (s *CreateClusterRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateClusterRequest) GetCisEnabled() *bool {
	return s.CisEnabled
}

func (s *CreateClusterRequest) GetCloudMonitorFlags() *bool {
	return s.CloudMonitorFlags
}

func (s *CreateClusterRequest) GetClusterDomain() *string {
	return s.ClusterDomain
}

func (s *CreateClusterRequest) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *CreateClusterRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *CreateClusterRequest) GetContainerCidr() *string {
	return s.ContainerCidr
}

func (s *CreateClusterRequest) GetControlPlaneConfig() *CreateClusterRequestControlPlaneConfig {
	return s.ControlPlaneConfig
}

func (s *CreateClusterRequest) GetControlplaneLogComponents() []*string {
	return s.ControlplaneLogComponents
}

func (s *CreateClusterRequest) GetControlplaneLogProject() *string {
	return s.ControlplaneLogProject
}

func (s *CreateClusterRequest) GetControlplaneLogTtl() *string {
	return s.ControlplaneLogTtl
}

func (s *CreateClusterRequest) GetCpuPolicy() *string {
	return s.CpuPolicy
}

func (s *CreateClusterRequest) GetCustomSan() *string {
	return s.CustomSan
}

func (s *CreateClusterRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateClusterRequest) GetDisableRollback() *bool {
	return s.DisableRollback
}

func (s *CreateClusterRequest) GetEnableRrsa() *bool {
	return s.EnableRrsa
}

func (s *CreateClusterRequest) GetEncryptionProviderKey() *string {
	return s.EncryptionProviderKey
}

func (s *CreateClusterRequest) GetEndpointPublicAccess() *bool {
	return s.EndpointPublicAccess
}

func (s *CreateClusterRequest) GetExtraSans() []*string {
	return s.ExtraSans
}

func (s *CreateClusterRequest) GetFormatDisk() *bool {
	return s.FormatDisk
}

func (s *CreateClusterRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateClusterRequest) GetImageType() *string {
	return s.ImageType
}

func (s *CreateClusterRequest) GetInstances() []*string {
	return s.Instances
}

func (s *CreateClusterRequest) GetIpStack() *string {
	return s.IpStack
}

func (s *CreateClusterRequest) GetIsEnterpriseSecurityGroup() *bool {
	return s.IsEnterpriseSecurityGroup
}

func (s *CreateClusterRequest) GetKeepInstanceName() *bool {
	return s.KeepInstanceName
}

func (s *CreateClusterRequest) GetKeyPair() *string {
	return s.KeyPair
}

func (s *CreateClusterRequest) GetKubernetesVersion() *string {
	return s.KubernetesVersion
}

func (s *CreateClusterRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateClusterRequest) GetLoadBalancerSpec() *string {
	return s.LoadBalancerSpec
}

func (s *CreateClusterRequest) GetLoggingType() *string {
	return s.LoggingType
}

func (s *CreateClusterRequest) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *CreateClusterRequest) GetMaintenanceWindow() *MaintenanceWindow {
	return s.MaintenanceWindow
}

func (s *CreateClusterRequest) GetMasterAutoRenew() *bool {
	return s.MasterAutoRenew
}

func (s *CreateClusterRequest) GetMasterAutoRenewPeriod() *int64 {
	return s.MasterAutoRenewPeriod
}

func (s *CreateClusterRequest) GetMasterCount() *int64 {
	return s.MasterCount
}

func (s *CreateClusterRequest) GetMasterInstanceChargeType() *string {
	return s.MasterInstanceChargeType
}

func (s *CreateClusterRequest) GetMasterInstanceTypes() []*string {
	return s.MasterInstanceTypes
}

func (s *CreateClusterRequest) GetMasterPeriod() *int64 {
	return s.MasterPeriod
}

func (s *CreateClusterRequest) GetMasterPeriodUnit() *string {
	return s.MasterPeriodUnit
}

func (s *CreateClusterRequest) GetMasterSystemDiskCategory() *string {
	return s.MasterSystemDiskCategory
}

func (s *CreateClusterRequest) GetMasterSystemDiskPerformanceLevel() *string {
	return s.MasterSystemDiskPerformanceLevel
}

func (s *CreateClusterRequest) GetMasterSystemDiskSize() *int64 {
	return s.MasterSystemDiskSize
}

func (s *CreateClusterRequest) GetMasterSystemDiskSnapshotPolicyId() *string {
	return s.MasterSystemDiskSnapshotPolicyId
}

func (s *CreateClusterRequest) GetMasterVswitchIds() []*string {
	return s.MasterVswitchIds
}

func (s *CreateClusterRequest) GetName() *string {
	return s.Name
}

func (s *CreateClusterRequest) GetNatGateway() *bool {
	return s.NatGateway
}

func (s *CreateClusterRequest) GetNodeCidrMask() *string {
	return s.NodeCidrMask
}

func (s *CreateClusterRequest) GetNodeNameMode() *string {
	return s.NodeNameMode
}

func (s *CreateClusterRequest) GetNodePortRange() *string {
	return s.NodePortRange
}

func (s *CreateClusterRequest) GetNodepools() []*Nodepool {
	return s.Nodepools
}

func (s *CreateClusterRequest) GetNumOfNodes() *int64 {
	return s.NumOfNodes
}

func (s *CreateClusterRequest) GetOperationPolicy() *CreateClusterRequestOperationPolicy {
	return s.OperationPolicy
}

func (s *CreateClusterRequest) GetOsType() *string {
	return s.OsType
}

func (s *CreateClusterRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *CreateClusterRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateClusterRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateClusterRequest) GetPodVswitchIds() []*string {
	return s.PodVswitchIds
}

func (s *CreateClusterRequest) GetProfile() *string {
	return s.Profile
}

func (s *CreateClusterRequest) GetProxyMode() *string {
	return s.ProxyMode
}

func (s *CreateClusterRequest) GetRdsInstances() []*string {
	return s.RdsInstances
}

func (s *CreateClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateClusterRequest) GetRrsaConfig() *CreateClusterRequestRrsaConfig {
	return s.RrsaConfig
}

func (s *CreateClusterRequest) GetRuntime() *Runtime {
	return s.Runtime
}

func (s *CreateClusterRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateClusterRequest) GetSecurityHardeningOs() *bool {
	return s.SecurityHardeningOs
}

func (s *CreateClusterRequest) GetServiceAccountIssuer() *string {
	return s.ServiceAccountIssuer
}

func (s *CreateClusterRequest) GetServiceCidr() *string {
	return s.ServiceCidr
}

func (s *CreateClusterRequest) GetServiceDiscoveryTypes() []*string {
	return s.ServiceDiscoveryTypes
}

func (s *CreateClusterRequest) GetSnatEntry() *bool {
	return s.SnatEntry
}

func (s *CreateClusterRequest) GetSocEnabled() *bool {
	return s.SocEnabled
}

func (s *CreateClusterRequest) GetSshFlags() *bool {
	return s.SshFlags
}

func (s *CreateClusterRequest) GetTags() []*Tag {
	return s.Tags
}

func (s *CreateClusterRequest) GetTaints() []*Taint {
	return s.Taints
}

func (s *CreateClusterRequest) GetTimeoutMins() *int64 {
	return s.TimeoutMins
}

func (s *CreateClusterRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *CreateClusterRequest) GetUserCa() *string {
	return s.UserCa
}

func (s *CreateClusterRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateClusterRequest) GetVpcid() *string {
	return s.Vpcid
}

func (s *CreateClusterRequest) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *CreateClusterRequest) GetWorkerAutoRenew() *bool {
	return s.WorkerAutoRenew
}

func (s *CreateClusterRequest) GetWorkerAutoRenewPeriod() *int64 {
	return s.WorkerAutoRenewPeriod
}

func (s *CreateClusterRequest) GetWorkerDataDisks() []*CreateClusterRequestWorkerDataDisks {
	return s.WorkerDataDisks
}

func (s *CreateClusterRequest) GetWorkerInstanceChargeType() *string {
	return s.WorkerInstanceChargeType
}

func (s *CreateClusterRequest) GetWorkerInstanceTypes() []*string {
	return s.WorkerInstanceTypes
}

func (s *CreateClusterRequest) GetWorkerPeriod() *int64 {
	return s.WorkerPeriod
}

func (s *CreateClusterRequest) GetWorkerPeriodUnit() *string {
	return s.WorkerPeriodUnit
}

func (s *CreateClusterRequest) GetWorkerSystemDiskCategory() *string {
	return s.WorkerSystemDiskCategory
}

func (s *CreateClusterRequest) GetWorkerSystemDiskPerformanceLevel() *string {
	return s.WorkerSystemDiskPerformanceLevel
}

func (s *CreateClusterRequest) GetWorkerSystemDiskSize() *int64 {
	return s.WorkerSystemDiskSize
}

func (s *CreateClusterRequest) GetWorkerSystemDiskSnapshotPolicyId() *string {
	return s.WorkerSystemDiskSnapshotPolicyId
}

func (s *CreateClusterRequest) GetWorkerVswitchIds() []*string {
	return s.WorkerVswitchIds
}

func (s *CreateClusterRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateClusterRequest) GetZoneIds() []*string {
	return s.ZoneIds
}

func (s *CreateClusterRequest) SetAccessControlList(v []*string) *CreateClusterRequest {
	s.AccessControlList = v
	return s
}

func (s *CreateClusterRequest) SetAddons(v []*Addon) *CreateClusterRequest {
	s.Addons = v
	return s
}

func (s *CreateClusterRequest) SetApiAudiences(v string) *CreateClusterRequest {
	s.ApiAudiences = &v
	return s
}

func (s *CreateClusterRequest) SetAuditLogConfig(v *CreateClusterRequestAuditLogConfig) *CreateClusterRequest {
	s.AuditLogConfig = v
	return s
}

func (s *CreateClusterRequest) SetAutoMode(v *CreateClusterRequestAutoMode) *CreateClusterRequest {
	s.AutoMode = v
	return s
}

func (s *CreateClusterRequest) SetAutoRenew(v bool) *CreateClusterRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateClusterRequest) SetAutoRenewPeriod(v int64) *CreateClusterRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateClusterRequest) SetChargeType(v string) *CreateClusterRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateClusterRequest) SetCisEnabled(v bool) *CreateClusterRequest {
	s.CisEnabled = &v
	return s
}

func (s *CreateClusterRequest) SetCloudMonitorFlags(v bool) *CreateClusterRequest {
	s.CloudMonitorFlags = &v
	return s
}

func (s *CreateClusterRequest) SetClusterDomain(v string) *CreateClusterRequest {
	s.ClusterDomain = &v
	return s
}

func (s *CreateClusterRequest) SetClusterSpec(v string) *CreateClusterRequest {
	s.ClusterSpec = &v
	return s
}

func (s *CreateClusterRequest) SetClusterType(v string) *CreateClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterRequest) SetContainerCidr(v string) *CreateClusterRequest {
	s.ContainerCidr = &v
	return s
}

func (s *CreateClusterRequest) SetControlPlaneConfig(v *CreateClusterRequestControlPlaneConfig) *CreateClusterRequest {
	s.ControlPlaneConfig = v
	return s
}

func (s *CreateClusterRequest) SetControlplaneLogComponents(v []*string) *CreateClusterRequest {
	s.ControlplaneLogComponents = v
	return s
}

func (s *CreateClusterRequest) SetControlplaneLogProject(v string) *CreateClusterRequest {
	s.ControlplaneLogProject = &v
	return s
}

func (s *CreateClusterRequest) SetControlplaneLogTtl(v string) *CreateClusterRequest {
	s.ControlplaneLogTtl = &v
	return s
}

func (s *CreateClusterRequest) SetCpuPolicy(v string) *CreateClusterRequest {
	s.CpuPolicy = &v
	return s
}

func (s *CreateClusterRequest) SetCustomSan(v string) *CreateClusterRequest {
	s.CustomSan = &v
	return s
}

func (s *CreateClusterRequest) SetDeletionProtection(v bool) *CreateClusterRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateClusterRequest) SetDisableRollback(v bool) *CreateClusterRequest {
	s.DisableRollback = &v
	return s
}

func (s *CreateClusterRequest) SetEnableRrsa(v bool) *CreateClusterRequest {
	s.EnableRrsa = &v
	return s
}

func (s *CreateClusterRequest) SetEncryptionProviderKey(v string) *CreateClusterRequest {
	s.EncryptionProviderKey = &v
	return s
}

func (s *CreateClusterRequest) SetEndpointPublicAccess(v bool) *CreateClusterRequest {
	s.EndpointPublicAccess = &v
	return s
}

func (s *CreateClusterRequest) SetExtraSans(v []*string) *CreateClusterRequest {
	s.ExtraSans = v
	return s
}

func (s *CreateClusterRequest) SetFormatDisk(v bool) *CreateClusterRequest {
	s.FormatDisk = &v
	return s
}

func (s *CreateClusterRequest) SetImageId(v string) *CreateClusterRequest {
	s.ImageId = &v
	return s
}

func (s *CreateClusterRequest) SetImageType(v string) *CreateClusterRequest {
	s.ImageType = &v
	return s
}

func (s *CreateClusterRequest) SetInstances(v []*string) *CreateClusterRequest {
	s.Instances = v
	return s
}

func (s *CreateClusterRequest) SetIpStack(v string) *CreateClusterRequest {
	s.IpStack = &v
	return s
}

func (s *CreateClusterRequest) SetIsEnterpriseSecurityGroup(v bool) *CreateClusterRequest {
	s.IsEnterpriseSecurityGroup = &v
	return s
}

func (s *CreateClusterRequest) SetKeepInstanceName(v bool) *CreateClusterRequest {
	s.KeepInstanceName = &v
	return s
}

func (s *CreateClusterRequest) SetKeyPair(v string) *CreateClusterRequest {
	s.KeyPair = &v
	return s
}

func (s *CreateClusterRequest) SetKubernetesVersion(v string) *CreateClusterRequest {
	s.KubernetesVersion = &v
	return s
}

func (s *CreateClusterRequest) SetLoadBalancerId(v string) *CreateClusterRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateClusterRequest) SetLoadBalancerSpec(v string) *CreateClusterRequest {
	s.LoadBalancerSpec = &v
	return s
}

func (s *CreateClusterRequest) SetLoggingType(v string) *CreateClusterRequest {
	s.LoggingType = &v
	return s
}

func (s *CreateClusterRequest) SetLoginPassword(v string) *CreateClusterRequest {
	s.LoginPassword = &v
	return s
}

func (s *CreateClusterRequest) SetMaintenanceWindow(v *MaintenanceWindow) *CreateClusterRequest {
	s.MaintenanceWindow = v
	return s
}

func (s *CreateClusterRequest) SetMasterAutoRenew(v bool) *CreateClusterRequest {
	s.MasterAutoRenew = &v
	return s
}

func (s *CreateClusterRequest) SetMasterAutoRenewPeriod(v int64) *CreateClusterRequest {
	s.MasterAutoRenewPeriod = &v
	return s
}

func (s *CreateClusterRequest) SetMasterCount(v int64) *CreateClusterRequest {
	s.MasterCount = &v
	return s
}

func (s *CreateClusterRequest) SetMasterInstanceChargeType(v string) *CreateClusterRequest {
	s.MasterInstanceChargeType = &v
	return s
}

func (s *CreateClusterRequest) SetMasterInstanceTypes(v []*string) *CreateClusterRequest {
	s.MasterInstanceTypes = v
	return s
}

func (s *CreateClusterRequest) SetMasterPeriod(v int64) *CreateClusterRequest {
	s.MasterPeriod = &v
	return s
}

func (s *CreateClusterRequest) SetMasterPeriodUnit(v string) *CreateClusterRequest {
	s.MasterPeriodUnit = &v
	return s
}

func (s *CreateClusterRequest) SetMasterSystemDiskCategory(v string) *CreateClusterRequest {
	s.MasterSystemDiskCategory = &v
	return s
}

func (s *CreateClusterRequest) SetMasterSystemDiskPerformanceLevel(v string) *CreateClusterRequest {
	s.MasterSystemDiskPerformanceLevel = &v
	return s
}

func (s *CreateClusterRequest) SetMasterSystemDiskSize(v int64) *CreateClusterRequest {
	s.MasterSystemDiskSize = &v
	return s
}

func (s *CreateClusterRequest) SetMasterSystemDiskSnapshotPolicyId(v string) *CreateClusterRequest {
	s.MasterSystemDiskSnapshotPolicyId = &v
	return s
}

func (s *CreateClusterRequest) SetMasterVswitchIds(v []*string) *CreateClusterRequest {
	s.MasterVswitchIds = v
	return s
}

func (s *CreateClusterRequest) SetName(v string) *CreateClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateClusterRequest) SetNatGateway(v bool) *CreateClusterRequest {
	s.NatGateway = &v
	return s
}

func (s *CreateClusterRequest) SetNodeCidrMask(v string) *CreateClusterRequest {
	s.NodeCidrMask = &v
	return s
}

func (s *CreateClusterRequest) SetNodeNameMode(v string) *CreateClusterRequest {
	s.NodeNameMode = &v
	return s
}

func (s *CreateClusterRequest) SetNodePortRange(v string) *CreateClusterRequest {
	s.NodePortRange = &v
	return s
}

func (s *CreateClusterRequest) SetNodepools(v []*Nodepool) *CreateClusterRequest {
	s.Nodepools = v
	return s
}

func (s *CreateClusterRequest) SetNumOfNodes(v int64) *CreateClusterRequest {
	s.NumOfNodes = &v
	return s
}

func (s *CreateClusterRequest) SetOperationPolicy(v *CreateClusterRequestOperationPolicy) *CreateClusterRequest {
	s.OperationPolicy = v
	return s
}

func (s *CreateClusterRequest) SetOsType(v string) *CreateClusterRequest {
	s.OsType = &v
	return s
}

func (s *CreateClusterRequest) SetPeriod(v int64) *CreateClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateClusterRequest) SetPeriodUnit(v string) *CreateClusterRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateClusterRequest) SetPlatform(v string) *CreateClusterRequest {
	s.Platform = &v
	return s
}

func (s *CreateClusterRequest) SetPodVswitchIds(v []*string) *CreateClusterRequest {
	s.PodVswitchIds = v
	return s
}

func (s *CreateClusterRequest) SetProfile(v string) *CreateClusterRequest {
	s.Profile = &v
	return s
}

func (s *CreateClusterRequest) SetProxyMode(v string) *CreateClusterRequest {
	s.ProxyMode = &v
	return s
}

func (s *CreateClusterRequest) SetRdsInstances(v []*string) *CreateClusterRequest {
	s.RdsInstances = v
	return s
}

func (s *CreateClusterRequest) SetRegionId(v string) *CreateClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateClusterRequest) SetResourceGroupId(v string) *CreateClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetRrsaConfig(v *CreateClusterRequestRrsaConfig) *CreateClusterRequest {
	s.RrsaConfig = v
	return s
}

func (s *CreateClusterRequest) SetRuntime(v *Runtime) *CreateClusterRequest {
	s.Runtime = v
	return s
}

func (s *CreateClusterRequest) SetSecurityGroupId(v string) *CreateClusterRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityHardeningOs(v bool) *CreateClusterRequest {
	s.SecurityHardeningOs = &v
	return s
}

func (s *CreateClusterRequest) SetServiceAccountIssuer(v string) *CreateClusterRequest {
	s.ServiceAccountIssuer = &v
	return s
}

func (s *CreateClusterRequest) SetServiceCidr(v string) *CreateClusterRequest {
	s.ServiceCidr = &v
	return s
}

func (s *CreateClusterRequest) SetServiceDiscoveryTypes(v []*string) *CreateClusterRequest {
	s.ServiceDiscoveryTypes = v
	return s
}

func (s *CreateClusterRequest) SetSnatEntry(v bool) *CreateClusterRequest {
	s.SnatEntry = &v
	return s
}

func (s *CreateClusterRequest) SetSocEnabled(v bool) *CreateClusterRequest {
	s.SocEnabled = &v
	return s
}

func (s *CreateClusterRequest) SetSshFlags(v bool) *CreateClusterRequest {
	s.SshFlags = &v
	return s
}

func (s *CreateClusterRequest) SetTags(v []*Tag) *CreateClusterRequest {
	s.Tags = v
	return s
}

func (s *CreateClusterRequest) SetTaints(v []*Taint) *CreateClusterRequest {
	s.Taints = v
	return s
}

func (s *CreateClusterRequest) SetTimeoutMins(v int64) *CreateClusterRequest {
	s.TimeoutMins = &v
	return s
}

func (s *CreateClusterRequest) SetTimezone(v string) *CreateClusterRequest {
	s.Timezone = &v
	return s
}

func (s *CreateClusterRequest) SetUserCa(v string) *CreateClusterRequest {
	s.UserCa = &v
	return s
}

func (s *CreateClusterRequest) SetUserData(v string) *CreateClusterRequest {
	s.UserData = &v
	return s
}

func (s *CreateClusterRequest) SetVpcid(v string) *CreateClusterRequest {
	s.Vpcid = &v
	return s
}

func (s *CreateClusterRequest) SetVswitchIds(v []*string) *CreateClusterRequest {
	s.VswitchIds = v
	return s
}

func (s *CreateClusterRequest) SetWorkerAutoRenew(v bool) *CreateClusterRequest {
	s.WorkerAutoRenew = &v
	return s
}

func (s *CreateClusterRequest) SetWorkerAutoRenewPeriod(v int64) *CreateClusterRequest {
	s.WorkerAutoRenewPeriod = &v
	return s
}

func (s *CreateClusterRequest) SetWorkerDataDisks(v []*CreateClusterRequestWorkerDataDisks) *CreateClusterRequest {
	s.WorkerDataDisks = v
	return s
}

func (s *CreateClusterRequest) SetWorkerInstanceChargeType(v string) *CreateClusterRequest {
	s.WorkerInstanceChargeType = &v
	return s
}

func (s *CreateClusterRequest) SetWorkerInstanceTypes(v []*string) *CreateClusterRequest {
	s.WorkerInstanceTypes = v
	return s
}

func (s *CreateClusterRequest) SetWorkerPeriod(v int64) *CreateClusterRequest {
	s.WorkerPeriod = &v
	return s
}

func (s *CreateClusterRequest) SetWorkerPeriodUnit(v string) *CreateClusterRequest {
	s.WorkerPeriodUnit = &v
	return s
}

func (s *CreateClusterRequest) SetWorkerSystemDiskCategory(v string) *CreateClusterRequest {
	s.WorkerSystemDiskCategory = &v
	return s
}

func (s *CreateClusterRequest) SetWorkerSystemDiskPerformanceLevel(v string) *CreateClusterRequest {
	s.WorkerSystemDiskPerformanceLevel = &v
	return s
}

func (s *CreateClusterRequest) SetWorkerSystemDiskSize(v int64) *CreateClusterRequest {
	s.WorkerSystemDiskSize = &v
	return s
}

func (s *CreateClusterRequest) SetWorkerSystemDiskSnapshotPolicyId(v string) *CreateClusterRequest {
	s.WorkerSystemDiskSnapshotPolicyId = &v
	return s
}

func (s *CreateClusterRequest) SetWorkerVswitchIds(v []*string) *CreateClusterRequest {
	s.WorkerVswitchIds = v
	return s
}

func (s *CreateClusterRequest) SetZoneId(v string) *CreateClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateClusterRequest) SetZoneIds(v []*string) *CreateClusterRequest {
	s.ZoneIds = v
	return s
}

func (s *CreateClusterRequest) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestAuditLogConfig struct {
	Enabled        *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	SlsProjectName *string `json:"sls_project_name,omitempty" xml:"sls_project_name,omitempty"`
}

func (s CreateClusterRequestAuditLogConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestAuditLogConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestAuditLogConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateClusterRequestAuditLogConfig) GetSlsProjectName() *string {
	return s.SlsProjectName
}

func (s *CreateClusterRequestAuditLogConfig) SetEnabled(v bool) *CreateClusterRequestAuditLogConfig {
	s.Enabled = &v
	return s
}

func (s *CreateClusterRequestAuditLogConfig) SetSlsProjectName(v string) *CreateClusterRequestAuditLogConfig {
	s.SlsProjectName = &v
	return s
}

func (s *CreateClusterRequestAuditLogConfig) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestAutoMode struct {
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s CreateClusterRequestAutoMode) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestAutoMode) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestAutoMode) GetEnable() *bool {
	return s.Enable
}

func (s *CreateClusterRequestAutoMode) SetEnable(v bool) *CreateClusterRequestAutoMode {
	s.Enable = &v
	return s
}

func (s *CreateClusterRequestAutoMode) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestControlPlaneConfig struct {
	// Specifies whether to enable auto-renewal for the node.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// The auto-renewal duration for the node.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// The billing method of the node.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"charge_type,omitempty" xml:"charge_type,omitempty"`
	// Specifies whether to install CloudMonitor on the node.
	//
	// example:
	//
	// true
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// The CPU management policy of the node.
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
	// The image ID.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20240819.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The type of the OS image.
	//
	// example:
	//
	// AliyunLinux3
	ImageType               *string                  `json:"image_type,omitempty" xml:"image_type,omitempty"`
	InstanceMetadataOptions *InstanceMetadataOptions `json:"instance_metadata_options,omitempty" xml:"instance_metadata_options,omitempty"`
	// The instance types of the nodes.
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// The name of the key pair. You must set this parameter or login_password.
	//
	// example:
	//
	// ack
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// The SSH logon password. The password must be 8 to 30 characters in length and contain a minimum of three of the following character types: uppercase letters, lowercase letters, digits, and special characters. You must set this parameter or key_pair.
	//
	// example:
	//
	// ack@Test
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The node port range.
	//
	// example:
	//
	// 30000-32767
	NodePortRange *string `json:"node_port_range,omitempty" xml:"node_port_range,omitempty"`
	// The subscription duration of the node.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The unit of the subscription duration of the node.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// The container runtime.
	//
	// example:
	//
	// containerd
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// Specifies whether to enable Alibaba Cloud Linux Security Hardening.
	//
	// example:
	//
	// true
	SecurityHardeningOs *bool `json:"security_hardening_os,omitempty" xml:"security_hardening_os,omitempty"`
	// The number of control plane nodes.
	//
	// example:
	//
	// 3
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// Specifies whether to enable MLPS security hardening.
	//
	// example:
	//
	// true
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// Specifies whether to enable the burst feature for the system disk.
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// The system disk category for the node.
	//
	// example:
	//
	// cloud_essd
	SystemDiskCategory *string `json:"system_disk_category,omitempty" xml:"system_disk_category,omitempty"`
	// The PL of the system disk that you want to use for the node. This parameter takes effect only for ESSDs.
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// The preset read/write IOPS of the system disk.
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// The system disk size of the node. The value must be at least 40 GB.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// The automatic snapshot policy of the node.
	//
	// example:
	//
	// sp-2zej1nogjvovnz4z****
	SystemDiskSnapshotPolicyId *string `json:"system_disk_snapshot_policy_id,omitempty" xml:"system_disk_snapshot_policy_id,omitempty"`
}

func (s CreateClusterRequestControlPlaneConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestControlPlaneConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestControlPlaneConfig) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateClusterRequestControlPlaneConfig) GetAutoRenewPeriod() *int64 {
	return s.AutoRenewPeriod
}

func (s *CreateClusterRequestControlPlaneConfig) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateClusterRequestControlPlaneConfig) GetCloudMonitorFlags() *bool {
	return s.CloudMonitorFlags
}

func (s *CreateClusterRequestControlPlaneConfig) GetCpuPolicy() *string {
	return s.CpuPolicy
}

func (s *CreateClusterRequestControlPlaneConfig) GetDeploymentsetId() *string {
	return s.DeploymentsetId
}

func (s *CreateClusterRequestControlPlaneConfig) GetImageId() *string {
	return s.ImageId
}

func (s *CreateClusterRequestControlPlaneConfig) GetImageType() *string {
	return s.ImageType
}

func (s *CreateClusterRequestControlPlaneConfig) GetInstanceMetadataOptions() *InstanceMetadataOptions {
	return s.InstanceMetadataOptions
}

func (s *CreateClusterRequestControlPlaneConfig) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *CreateClusterRequestControlPlaneConfig) GetKeyPair() *string {
	return s.KeyPair
}

func (s *CreateClusterRequestControlPlaneConfig) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *CreateClusterRequestControlPlaneConfig) GetNodePortRange() *string {
	return s.NodePortRange
}

func (s *CreateClusterRequestControlPlaneConfig) GetPeriod() *int64 {
	return s.Period
}

func (s *CreateClusterRequestControlPlaneConfig) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateClusterRequestControlPlaneConfig) GetRuntime() *string {
	return s.Runtime
}

func (s *CreateClusterRequestControlPlaneConfig) GetSecurityHardeningOs() *bool {
	return s.SecurityHardeningOs
}

func (s *CreateClusterRequestControlPlaneConfig) GetSize() *int64 {
	return s.Size
}

func (s *CreateClusterRequestControlPlaneConfig) GetSocEnabled() *bool {
	return s.SocEnabled
}

func (s *CreateClusterRequestControlPlaneConfig) GetSystemDiskBurstingEnabled() *bool {
	return s.SystemDiskBurstingEnabled
}

func (s *CreateClusterRequestControlPlaneConfig) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *CreateClusterRequestControlPlaneConfig) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *CreateClusterRequestControlPlaneConfig) GetSystemDiskProvisionedIops() *int64 {
	return s.SystemDiskProvisionedIops
}

func (s *CreateClusterRequestControlPlaneConfig) GetSystemDiskSize() *int64 {
	return s.SystemDiskSize
}

func (s *CreateClusterRequestControlPlaneConfig) GetSystemDiskSnapshotPolicyId() *string {
	return s.SystemDiskSnapshotPolicyId
}

func (s *CreateClusterRequestControlPlaneConfig) SetAutoRenew(v bool) *CreateClusterRequestControlPlaneConfig {
	s.AutoRenew = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetAutoRenewPeriod(v int64) *CreateClusterRequestControlPlaneConfig {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetChargeType(v string) *CreateClusterRequestControlPlaneConfig {
	s.ChargeType = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetCloudMonitorFlags(v bool) *CreateClusterRequestControlPlaneConfig {
	s.CloudMonitorFlags = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetCpuPolicy(v string) *CreateClusterRequestControlPlaneConfig {
	s.CpuPolicy = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetDeploymentsetId(v string) *CreateClusterRequestControlPlaneConfig {
	s.DeploymentsetId = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetImageId(v string) *CreateClusterRequestControlPlaneConfig {
	s.ImageId = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetImageType(v string) *CreateClusterRequestControlPlaneConfig {
	s.ImageType = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetInstanceMetadataOptions(v *InstanceMetadataOptions) *CreateClusterRequestControlPlaneConfig {
	s.InstanceMetadataOptions = v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetInstanceTypes(v []*string) *CreateClusterRequestControlPlaneConfig {
	s.InstanceTypes = v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetKeyPair(v string) *CreateClusterRequestControlPlaneConfig {
	s.KeyPair = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetLoginPassword(v string) *CreateClusterRequestControlPlaneConfig {
	s.LoginPassword = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetNodePortRange(v string) *CreateClusterRequestControlPlaneConfig {
	s.NodePortRange = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetPeriod(v int64) *CreateClusterRequestControlPlaneConfig {
	s.Period = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetPeriodUnit(v string) *CreateClusterRequestControlPlaneConfig {
	s.PeriodUnit = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetRuntime(v string) *CreateClusterRequestControlPlaneConfig {
	s.Runtime = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetSecurityHardeningOs(v bool) *CreateClusterRequestControlPlaneConfig {
	s.SecurityHardeningOs = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetSize(v int64) *CreateClusterRequestControlPlaneConfig {
	s.Size = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetSocEnabled(v bool) *CreateClusterRequestControlPlaneConfig {
	s.SocEnabled = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetSystemDiskBurstingEnabled(v bool) *CreateClusterRequestControlPlaneConfig {
	s.SystemDiskBurstingEnabled = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetSystemDiskCategory(v string) *CreateClusterRequestControlPlaneConfig {
	s.SystemDiskCategory = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetSystemDiskPerformanceLevel(v string) *CreateClusterRequestControlPlaneConfig {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetSystemDiskProvisionedIops(v int64) *CreateClusterRequestControlPlaneConfig {
	s.SystemDiskProvisionedIops = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetSystemDiskSize(v int64) *CreateClusterRequestControlPlaneConfig {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetSystemDiskSnapshotPolicyId(v string) *CreateClusterRequestControlPlaneConfig {
	s.SystemDiskSnapshotPolicyId = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestOperationPolicy struct {
	// The configurations of auto cluster upgrade.
	ClusterAutoUpgrade *CreateClusterRequestOperationPolicyClusterAutoUpgrade `json:"cluster_auto_upgrade,omitempty" xml:"cluster_auto_upgrade,omitempty" type:"Struct"`
}

func (s CreateClusterRequestOperationPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestOperationPolicy) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestOperationPolicy) GetClusterAutoUpgrade() *CreateClusterRequestOperationPolicyClusterAutoUpgrade {
	return s.ClusterAutoUpgrade
}

func (s *CreateClusterRequestOperationPolicy) SetClusterAutoUpgrade(v *CreateClusterRequestOperationPolicyClusterAutoUpgrade) *CreateClusterRequestOperationPolicy {
	s.ClusterAutoUpgrade = v
	return s
}

func (s *CreateClusterRequestOperationPolicy) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestOperationPolicyClusterAutoUpgrade struct {
	// The automatic update frequency. Valid values:
	//
	// 	- patch
	//
	// 	- stable
	//
	// 	- rapid
	//
	// example:
	//
	// patch
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// Specifies whether to enable auto cluster update.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateClusterRequestOperationPolicyClusterAutoUpgrade) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestOperationPolicyClusterAutoUpgrade) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestOperationPolicyClusterAutoUpgrade) GetChannel() *string {
	return s.Channel
}

func (s *CreateClusterRequestOperationPolicyClusterAutoUpgrade) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateClusterRequestOperationPolicyClusterAutoUpgrade) SetChannel(v string) *CreateClusterRequestOperationPolicyClusterAutoUpgrade {
	s.Channel = &v
	return s
}

func (s *CreateClusterRequestOperationPolicyClusterAutoUpgrade) SetEnabled(v bool) *CreateClusterRequestOperationPolicyClusterAutoUpgrade {
	s.Enabled = &v
	return s
}

func (s *CreateClusterRequestOperationPolicyClusterAutoUpgrade) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestRrsaConfig struct {
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateClusterRequestRrsaConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestRrsaConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestRrsaConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateClusterRequestRrsaConfig) SetEnabled(v bool) *CreateClusterRequestRrsaConfig {
	s.Enabled = &v
	return s
}

func (s *CreateClusterRequestRrsaConfig) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestWorkerDataDisks struct {
	// The data disk category.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// Specifies whether to encrypt the data disk. Valid values:
	//
	// 	- `true`: encrypts the data disk.
	//
	// 	- `false`: does not encrypt the data disk.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	Encrypted *string `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// The PL of the data disk. This parameter takes effect only for ESSDs. You can specify a higher PL if you increase the size of a data disk. For more information, see [Enterprise SSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"performance_level,omitempty" xml:"performance_level,omitempty"`
	// The data disk size. Valid values: 40 to 32767. Unit: GiB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
}

func (s CreateClusterRequestWorkerDataDisks) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestWorkerDataDisks) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestWorkerDataDisks) GetCategory() *string {
	return s.Category
}

func (s *CreateClusterRequestWorkerDataDisks) GetEncrypted() *string {
	return s.Encrypted
}

func (s *CreateClusterRequestWorkerDataDisks) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateClusterRequestWorkerDataDisks) GetSize() *string {
	return s.Size
}

func (s *CreateClusterRequestWorkerDataDisks) SetCategory(v string) *CreateClusterRequestWorkerDataDisks {
	s.Category = &v
	return s
}

func (s *CreateClusterRequestWorkerDataDisks) SetEncrypted(v string) *CreateClusterRequestWorkerDataDisks {
	s.Encrypted = &v
	return s
}

func (s *CreateClusterRequestWorkerDataDisks) SetPerformanceLevel(v string) *CreateClusterRequestWorkerDataDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateClusterRequestWorkerDataDisks) SetSize(v string) *CreateClusterRequestWorkerDataDisks {
	s.Size = &v
	return s
}

func (s *CreateClusterRequestWorkerDataDisks) Validate() error {
	return dara.Validate(s)
}
