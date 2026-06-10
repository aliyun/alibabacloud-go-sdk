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
	SetControlPlaneEndpointsConfig(v *CreateClusterRequestControlPlaneEndpointsConfig) *CreateClusterRequest
	GetControlPlaneEndpointsConfig() *CreateClusterRequestControlPlaneEndpointsConfig
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
	// Deprecated
	//
	// **[Deprecated]*	- Access control list for the registered cluster API Server SLB.
	AccessControlList []*string `json:"access_control_list,omitempty" xml:"access_control_list,omitempty" type:"Repeated"`
	// List of cluster components. Specify components to install during cluster creation using the `addons` parameter.
	//
	// **Network components**: Required. Choose between Flannel and Terway:
	//
	// - Flannel network: [{"name":"flannel","config":""}].
	//
	// - Terway network: [{"name": "terway-eniip","config": ""}].
	//
	// **Storage components**: Optional. Only `csi` is supported:
	//
	// `csi`: [{"name":"csi-plugin","config": ""},{"name": "csi-provisioner","config": ""}].
	//
	// **Logging components**: Optional. We recommend enabling this. Without SLS, you cannot use cluster audit features.
	//
	// - Use an existing `SLS Project`: [{"name": "loongcollector","config": "{"IngressDashboardEnabled":"true","sls_project_name":"your_sls_project_name"}"}].
	//
	// - Create a new `SLS Project`: [{"name": "loongcollector","config": "{"IngressDashboardEnabled":"true"}"}].
	//
	// **Ingress components**: Optional. ACK dedicated clusters install the `nginx-ingress-controller` by default.
	//
	// - Install Ingress with public network access: [{"name":"nginx-ingress-controller","config":"{"IngressSlbNetworkType":"internet"}"}].
	//
	// - Disable default Ingress installation: [{"name": "nginx-ingress-controller","config": "","disabled": true}].
	//
	// **Event Hub**: Optional. Enabled by default.
	//
	// Event Hub provides storage, querying, and alerting for Kubernetes events. The associated Logstore is free for 90 days. For more information about the free policy, see [Create and use Kubernetes Event Hub](https://help.aliyun.com/document_detail/150476.html).
	//
	// Example to enable Event Hub: [{"name":"ack-node-problem-detector","config":"{"sls_project_name":"your_sls_project_name"}"}].
	Addons []*Addon `json:"addons,omitempty" xml:"addons,omitempty" type:"Repeated"`
	// ServiceAccount is the access credential used by pods to communicate with the cluster API Server. The `api-audiences` specifies valid `token` identities for authenticating requests on the `apiserver` side. You can configure multiple `audience` values separated by commas (,).
	//
	// For more information about `ServiceAccount`, see [Deploy service account token volume projection](https://help.aliyun.com/document_detail/160384.html).
	//
	// example:
	//
	// kubernetes.default.svc
	ApiAudiences *string `json:"api_audiences,omitempty" xml:"api_audiences,omitempty"`
	// Cluster audit log configuration.
	AuditLogConfig *CreateClusterRequestAuditLogConfig `json:"audit_log_config,omitempty" xml:"audit_log_config,omitempty" type:"Struct"`
	// [Intelligent Managed Mode](https://help.aliyun.com/document_detail/2938898.html) configuration.
	AutoMode *CreateClusterRequestAutoMode `json:"auto_mode,omitempty" xml:"auto_mode,omitempty" type:"Struct"`
	// Deprecated
	//
	// **[Deprecated]**
	//
	// Whether to enable auto-renewal. Valid only when `charge_type` is `PrePaid`. Valid values:
	//
	// - `true`: Enable auto-renewal.
	//
	// - `false`: Disable auto-renewal.
	//
	// Default value: `false`.
	//
	// This parameter changed on October 15, 2024. For more information, see [Announcement on Changes to CreateCluster API Parameters](https://help.aliyun.com/document_detail/2849194.html).
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// Deprecated
	//
	// **[Deprecated]**
	//
	// Auto-renewal period in months. Valid only when prepaid and auto-renewal are enabled. When `PeriodUnit=Month`, valid values are {1, 2, 3, 6, 12}.
	//
	// Default value: 1.
	//
	// This parameter changed on October 15, 2024. For more information, see [Announcement on Changes to CreateCluster API Parameters](https://help.aliyun.com/document_detail/2849194.html).
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// Deprecated
	//
	// **[Deprecated]**
	//
	// Billing type for the CLB instance used by the API Server. Default value: PostPaid. Valid values:
	//
	// - PostPaid: Pay-as-you-go.
	//
	// - PrePaid: Subscription. New CLB instances no longer support subscription billing, but existing instances are unaffected.
	//
	// 	Notice:
	//
	// - This parameter changed on October 15, 2024. For more information, see [Announcement on Changes to CreateCluster API Parameters](https://help.aliyun.com/document_detail/2849194.html).
	//
	// - Starting December 1, 2024, new CLB instances will no longer support subscription billing and will incur an instance fee.
	//
	// <props="china">For details, see [[Product Announcement\\] Discontinuation of Subscription Billing for API Server CLB in New Clusters](~~2851191~~) and [Adjustment of Classic Load Balancer Billing Items](https://help.aliyun.com/document_detail/2839797.html).
	//
	// <props="intl">For details, see [Adjustment of Classic Load Balancer Billing Items](https://help.aliyun.com/document_detail/2839797.html).
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"charge_type,omitempty" xml:"charge_type,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `security_hardening_os` parameter under `control_plane_config`. For node pool configuration, use the `security_hardening_os` parameter under `scaling_group` in `nodepool`.
	//
	// example:
	//
	// false
	CisEnabled *bool `json:"cis_enabled,omitempty" xml:"cis_enabled,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane node configuration, use the cloud_monitor_flags parameter under `control_plane_config`. For node pool configuration, use the `cms_enabled` parameter under `kubernetes_config` in `nodepool`.
	//
	// Whether to install the CloudMonitor agent on the cluster. Valid values:
	//
	// - `true`: Install the CloudMonitor agent.
	//
	// - `false`: Do not install the CloudMonitor agent.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// Cluster local domain name.
	//
	// Naming rules: The domain name consists of one or more parts separated by dots (.). Each part can be up to 63 characters long and can contain lowercase letters, digits, and hyphens (-). Each part must start and end with a lowercase letter or digit.
	//
	// example:
	//
	// cluster.local
	ClusterDomain *string `json:"cluster_domain,omitempty" xml:"cluster_domain,omitempty"`
	// When you set `cluster_type` to `ManagedKubernetes` and configure `profile`, you can further specify the cluster specification. Valid values:
	//
	// - `ack.standard`: Basic Edition (default when this parameter is empty)
	//
	// - `ack.pro.small`: Pro Edition
	//
	// - `ack.pro.xlarge`: Pro XL
	//
	// - `ack.pro.2xlarge`: Pro 2XL
	//
	// - `ack.pro.4xlarge`: Pro 4XL (requires whitelist approval from customer service)
	//
	// Pro XL, Pro 2XL, and Pro 4XL are three tiers provided by <props="china">[ACK Pro Provisioned Control Plane](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane)<props="intl">[ACK Pro Provisioned Control Plane](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane). These tiers pre-allocate and dedicate control plane resources to ensure consistently high API concurrency and pod scheduling performance. They are suitable for AI training and inference, ultra-large-scale clusters, and mission-critical workloads.
	//
	// For cluster management fees for Pro Edition and provisioned control plane clusters, see <props="china">[Cluster management fees](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee)<props="intl">[Cluster management fees](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee).
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// - `Kubernetes`: ACK dedicated cluster.
	//
	// - `ManagedKubernetes`: ACK managed clusters, including ACK managed clusters (Pro Edition, Basic Edition), ACK serverless clusters (Pro Edition, Basic Edition), ACK Edge clusters (Pro Edition, Basic Edition), and ACK LINGJUN clusters (Pro Edition).
	//
	// - `ExternalKubernetes`: registered cluster.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// Pod network CIDR block. It must be a valid private CIDR block: 10.0.0.0/8, 172.16.0.0/12 to 172.31.0.0/16, or 192.168.0.0/16. It cannot overlap with the VPC or existing Kubernetes cluster CIDR blocks in the VPC. This cannot be modified after cluster creation.
	//
	// For cluster network planning, see [Network planning for ACK managed clusters](https://help.aliyun.com/document_detail/86500.html).
	//
	// > This field is required for Flannel clusters.
	//
	// example:
	//
	// 172.20.0.0/16
	ContainerCidr *string `json:"container_cidr,omitempty" xml:"container_cidr,omitempty"`
	// Control plane configuration for ACK dedicated clusters.
	ControlPlaneConfig *CreateClusterRequestControlPlaneConfig `json:"control_plane_config,omitempty" xml:"control_plane_config,omitempty" type:"Struct"`
	// Cluster connection configuration.
	ControlPlaneEndpointsConfig *CreateClusterRequestControlPlaneEndpointsConfig `json:"control_plane_endpoints_config,omitempty" xml:"control_plane_endpoints_config,omitempty" type:"Struct"`
	// List of component names specifying which control plane component logs to collect.
	//
	// By default, logs are collected for kube-apiserver, kube-controller-manager, kube-scheduler, and cloud-controller-manager.
	ControlplaneLogComponents []*string `json:"controlplane_log_components,omitempty" xml:"controlplane_log_components,omitempty" type:"Repeated"`
	// SLS Project for control plane component logs. You can use an existing Project for log storage or let the system automatically create one. If auto-created, the Project name will be `k8s-log-{ClusterID}`.
	//
	// example:
	//
	// k8s-log-xxx
	ControlplaneLogProject *string `json:"controlplane_log_project,omitempty" xml:"controlplane_log_project,omitempty"`
	// Number of days to retain control plane component logs.
	//
	// example:
	//
	// 30
	ControlplaneLogTtl *string `json:"controlplane_log_ttl,omitempty" xml:"controlplane_log_ttl,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `cpu_policy` parameter under `control_plane_config`. For node pool configuration, use the `cpu_policy` parameter under `kubernetes_config` in `nodepool`.
	//
	// Node CPU management policy. Supported for cluster versions 1.12.6 and later:
	//
	// - `static`: Enhances CPU affinity and exclusivity for pods with specific resource characteristics on the node.
	//
	// - `none`: Uses the default CPU affinity scheme.
	//
	// Default value: `none`.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use the `extra_sans` parameter instead.
	//
	// Custom certificate SAN. Separate multiple IP addresses or domain names with commas (,).
	//
	// example:
	//
	// cs.aliyun.com
	CustomSan *string `json:"custom_san,omitempty" xml:"custom_san,omitempty"`
	// Cluster deletion protection prevents accidental cluster deletion through the console or API. Valid values:
	//
	// - `true`: Enable deletion protection. You cannot delete the cluster through the console or API.
	//
	// - `false`: Disable deletion protection. You can delete the cluster through the console or API.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- By default, clusters are not rolled back on creation failure. You must manually clean up failed clusters.
	//
	// Whether to roll back on cluster creation failure. Valid values:
	//
	// - `true`: Roll back on failure.
	//
	// - `false`: Do not roll back on failure.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	DisableRollback *bool `json:"disable_rollback,omitempty" xml:"disable_rollback,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use the `rrsa_config` parameter instead.
	//
	// Whether to enable RRSA.
	//
	// - true: Enable.
	//
	// - false: Disable.
	//
	// example:
	//
	// false
	EnableRrsa *bool `json:"enable_rrsa,omitempty" xml:"enable_rrsa,omitempty"`
	// KMS key ID used to encrypt data disks. For more information, see [Key Management Service](https://help.aliyun.com/document_detail/28935.html).
	//
	// > This feature is available only for professional managed clusters (ACK Pro clusters).
	//
	// example:
	//
	// 0fe64791-55eb-4fc7-84c5-c6c7cdca****
	EncryptionProviderKey *string `json:"encryption_provider_key,omitempty" xml:"encryption_provider_key,omitempty"`
	// Whether to enable public network access. Expose the API Server through an EIP to allow public network access to the cluster.
	//
	// - `true`: Enable public network access.
	//
	// - `false`: Disable public network access. If disabled, you cannot access the cluster API Server from external networks.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	EndpointPublicAccess *bool `json:"endpoint_public_access,omitempty" xml:"endpoint_public_access,omitempty"`
	// Custom API Server certificate SAN (Subject Alternative Name).
	ExtraSans []*string `json:"extra_sans,omitempty" xml:"extra_sans,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- Adding existing nodes during cluster creation is not supported. To add existing nodes to a cluster, first create a node pool and then call the [AttachInstancesToNodePool](https://help.aliyun.com/document_detail/2667920.html) API.
	//
	// When using existing instances to create a cluster, whether to mount data disks. Valid values:
	//
	// - `true`: Store containers and images on the data disk. Existing data on the data disk will be lost. Back up your data.
	//
	// - `false`: Do not store containers and images on the data disk.
	//
	// Default value: `false`.
	//
	// Data disk mounting rules:
	//
	// - If the ECS instance has a mounted data disk and the file system of the last data disk is uninitialized, the system automatically formats it as ext4 to store /var/lib/docker and /var/lib/kubelet.
	//
	// - If the ECS instance has no mounted data disks, no new data disks are mounted.
	//
	// example:
	//
	// false
	FormatDisk *bool `json:"format_disk,omitempty" xml:"format_disk,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `image_id` parameter under `control_plane_config`. For node pool configuration, use the `image_id` parameter under `scaling_group` in `nodepool`.
	//
	// Custom node image. By default, the system image is used. When you select a custom image, it replaces the default system image. See [Custom images](https://help.aliyun.com/document_detail/146647.html).
	//
	// example:
	//
	// m-bp16z7xko3vvv8gt****
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `image_type` parameter under `control_plane_config`. For node pool configuration, use the `image_type` parameter under `scaling_group` in `nodepool`.
	//
	// Operating system distribution type. We recommend using this parameter to specify the node operating system. Valid values:
	//
	// - CentOS
	//
	// - AliyunLinux
	//
	// - AliyunLinux Qboot
	//
	// - AliyunLinuxUEFI
	//
	// - AliyunLinux3
	//
	// - Windows
	//
	// - WindowsCore
	//
	// - AliyunLinux3Arm64
	//
	// - ContainerOS
	//
	// Default value: `CentOS`.
	//
	// example:
	//
	// AliyunLinux
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Adding existing nodes during cluster creation is not supported. To add existing nodes to a cluster, first create a node pool and then call the [AttachInstancesToNodePool](https://help.aliyun.com/document_detail/2667920.html) API.
	//
	// When using existing nodes to create a cluster, specify the ECS instance list. These instances are added as worker nodes to the cluster.
	//
	// > This field is required when using existing instances to create a cluster.
	Instances []*string `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// Cluster IP stack.
	//
	// example:
	//
	// 可选值：
	//
	// ipv4（单栈）
	//
	// dual（双栈），默认值为ipv4。
	IpStack *string `json:"ip_stack,omitempty" xml:"ip_stack,omitempty"`
	// Automatically create an advanced security group. This takes effect when `security_group_id` is empty.
	//
	// > With basic security groups, the total number of nodes and Terway pods in a cluster cannot exceed 2,000. We recommend using advanced security groups for Terway network clusters.
	//
	// - `true`: Create and use an advanced security group.
	//
	// - `false`: Use a basic security group.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	IsEnterpriseSecurityGroup *bool `json:"is_enterprise_security_group,omitempty" xml:"is_enterprise_security_group,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Adding existing nodes during cluster creation is not supported. To add existing nodes to a cluster, first create a node pool and then call the [AttachInstancesToNodePool](https://help.aliyun.com/document_detail/2667920.html) API.
	//
	// When using existing instances to create a cluster, whether to keep the instance names.
	//
	// - `true`: Keep.
	//
	// - `false`: Replace with system-generated names.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	KeepInstanceName *bool `json:"keep_instance_name,omitempty" xml:"keep_instance_name,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `key_pair` parameter under `control_plane_config`. For node pool configuration, use the `key_pair` parameter under `scaling_group` in `nodepool`.
	//
	// Key pair name. Choose either this parameter or `login_password`.
	//
	// example:
	//
	// security-key
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// Cluster version, aligned with the Kubernetes community baseline version. We recommend selecting the latest version. If not specified, the latest version is used by default.
	//
	// You can create clusters using any of the three most recent versions. Use the [DescribeKubernetesVersionMetadata](https://help.aliyun.com/document_detail/2667899.html) API to query supported cluster versions.
	//
	// For more information about Kubernetes versions supported by ACK, see [Overview of Kubernetes version releases](https://help.aliyun.com/document_detail/185269.html).
	//
	// example:
	//
	// 1.32.1-aliyun.1
	KubernetesVersion *string `json:"kubernetes_version,omitempty" xml:"kubernetes_version,omitempty"`
	// Specify the CLB instance ID for API Server access. If specified, no new API Server CLB is automatically created.
	//
	// > Ensure the CLB instance has no dependencies (such as listeners or backend servers). Shared and public CLB instances are not supported.
	//
	// example:
	//
	// lb-wz9t256gqa3vbouk****
	LoadBalancerId *string `json:"load_balancer_id,omitempty" xml:"load_balancer_id,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- CLB is billed based on usage. This parameter has no effect.
	//
	// Load Balancer specification. Valid values:
	//
	// - slb.s1.small
	//
	// - slb.s2.small
	//
	// - slb.s2.medium
	//
	// - slb.s3.small
	//
	// - slb.s3.medium
	//
	// - slb.s3.large
	//
	// Default value: `slb.s2.small`.
	//
	// example:
	//
	// slb.s2.small
	LoadBalancerSpec *string `json:"load_balancer_spec,omitempty" xml:"load_balancer_spec,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Enables SLS for the cluster. Applies only to ACK Serverless clusters and must be set to `SLS`.
	//
	// example:
	//
	// SLS
	LoggingType *string `json:"logging_type,omitempty" xml:"logging_type,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `login_password` parameter under `control_plane_config`. For node pool configuration, use the `login_password` parameter under `scaling_group` in `nodepool`.
	//
	// SSH login password. Choose either this parameter or `key_pair`. Password rules: 8–30 characters, containing at least three of the following: uppercase letters, lowercase letters, digits, and special characters.
	//
	// example:
	//
	// null
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// Cluster maintenance window.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window,omitempty" xml:"maintenance_window,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `auto_renew` parameter under `control_plane_config`.
	//
	// Whether to enable auto-renewal for master nodes. Valid only when `master_instance_charge_type` is `PrePaid`. Valid values:
	//
	// - `true`: Enable auto-renewal.
	//
	// - `false`: Disable auto-renewal.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	MasterAutoRenew *bool `json:"master_auto_renew,omitempty" xml:"master_auto_renew,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `auto_renew_period` parameter under `control_plane_config`.
	//
	// Auto-renewal period for master nodes in months. Required and valid only when subscription billing is selected.
	//
	// Valid values: {1, 2, 3, 6, 12}.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	MasterAutoRenewPeriod *int64 `json:"master_auto_renew_period,omitempty" xml:"master_auto_renew_period,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `size` parameter under `control_plane_config`.
	//
	// Number of master nodes. Valid values: `3` or `5`.
	//
	// Default value: `3`.
	//
	// example:
	//
	// 3
	MasterCount *int64 `json:"master_count,omitempty" xml:"master_count,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `instance_charge_type` parameter under `control_plane_config`.
	//
	// Master node billing type. Valid values:
	//
	// - `PrePaid`: Subscription.
	//
	// - `PostPaid`: Pay-as-you-go.
	//
	// Default value: `PostPaid`.
	//
	// example:
	//
	// PrePaid
	MasterInstanceChargeType *string `json:"master_instance_charge_type,omitempty" xml:"master_instance_charge_type,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `instance_types` parameter under `control_plane_config`.
	//
	// Master node instance types. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	MasterInstanceTypes []*string `json:"master_instance_types,omitempty" xml:"master_instance_types,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `unit` parameter under `control_plane_config`.
	//
	// Subscription duration for master nodes in months. Required and valid only when `master_instance_charge_type` is `PrePaid`.
	//
	// Valid values: {1, 2, 3, 6, 12, 24, 36, 48, 60}.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	MasterPeriod *int64 `json:"master_period,omitempty" xml:"master_period,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `period_unit` parameter under `control_plane_config`.
	//
	// Master node billing cycle. Required when master_instance_charge_type is `PrePaid`.
	//
	// Valid value: `Month`. Only monthly billing is supported.
	//
	// example:
	//
	// Month
	MasterPeriodUnit *string `json:"master_period_unit,omitempty" xml:"master_period_unit,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `system_disk_category` parameter under `control_plane_config`.
	//
	// Master node system disk type. Valid values:
	//
	// - `cloud_efficiency`: Ultra disk.
	//
	// - `cloud_ssd`: Standard SSD.
	//
	// - `cloud_essd`: ESSD.
	//
	// Default value: `cloud_ssd`. The default may vary by zone.
	//
	// example:
	//
	// cloud_ssd
	MasterSystemDiskCategory *string `json:"master_system_disk_category,omitempty" xml:"master_system_disk_category,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `system_disk_performance_level` parameter under `control_plane_config`.
	//
	// Performance level for master node system disks. Applies only to ESSD disks. Disk performance levels depend on disk size. For more information, see [ESSD](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	MasterSystemDiskPerformanceLevel *string `json:"master_system_disk_performance_level,omitempty" xml:"master_system_disk_performance_level,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `system_disk_size` parameter under `control_plane_config`.
	//
	// Valid range: [40,500].
	//
	// Default value: `120`.
	//
	// example:
	//
	// 120
	MasterSystemDiskSize *int64 `json:"master_system_disk_size,omitempty" xml:"master_system_disk_size,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `system_disk_snapshot_policy_id` parameter under `control_plane_config`.
	//
	// Automatic snapshot policy ID for master node system disks.
	//
	// example:
	//
	// sp-2zej1nogjvovnz4z****
	MasterSystemDiskSnapshotPolicyId *string `json:"master_system_disk_snapshot_policy_id,omitempty" xml:"master_system_disk_snapshot_policy_id,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use the `vswitch_ids` parameter instead.
	//
	// List of master node vSwitch IDs. The number of vSwitches must be in the range [1,3]. For high availability, we recommend selecting three vSwitches in different zones.
	//
	// The number of specified instance types must match `master_count` and correspond one-to-one with elements in `master_vswitch_ids`.
	MasterVswitchIds []*string `json:"master_vswitch_ids,omitempty" xml:"master_vswitch_ids,omitempty" type:"Repeated"`
	// Custom cluster name. It can contain digits, letters, Chinese characters, or hyphens (-). The name must be 1 to 63 characters long and cannot start with a hyphen (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster-demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use the `snat_entry` parameter instead.
	//
	// example:
	//
	// true
	NatGateway *bool `json:"nat_gateway,omitempty" xml:"nat_gateway,omitempty"`
	// Number of node IP addresses, determined by the specified CIDR mask. This applies only to Flannel network clusters.
	//
	// Default value: `26`.
	//
	// example:
	//
	// 25
	NodeCidrMask *string `json:"node_cidr_mask,omitempty" xml:"node_cidr_mask,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use the `node_name_mode` parameter under `kubernetes_config` in `nodepool` instead.
	//
	// example:
	//
	// null
	NodeNameMode *string `json:"node_name_mode,omitempty" xml:"node_name_mode,omitempty"`
	// Deprecated
	//
	// Node service port range: [30000,65535].
	//
	// Default value: `30000-32767`.
	//
	// example:
	//
	// 30000~32767
	NodePortRange *string `json:"node_port_range,omitempty" xml:"node_port_range,omitempty"`
	// List of node pools.
	Nodepools []*Nodepool `json:"nodepools,omitempty" xml:"nodepools,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configuration, use the `desired_size` parameter under `scaling_group` in `nodepool`.
	//
	// Number of worker nodes. Range: [0,100].
	//
	// example:
	//
	// 3
	NumOfNodes *int64 `json:"num_of_nodes,omitempty" xml:"num_of_nodes,omitempty"`
	// Cluster automated operations policy.
	OperationPolicy *CreateClusterRequestOperationPolicy `json:"operation_policy,omitempty" xml:"operation_policy,omitempty" type:"Struct"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane node configuration, use the `image_type` parameter under `control_plane_config`. For node pool configuration, use the `image_type` parameter under `scaling_group` in `nodepool`.
	//
	// Operating system platform type. Valid values:
	//
	// - Windows
	//
	// - Linux
	//
	// Default value: `Linux`.
	//
	// example:
	//
	// Linux
	OsType *string `json:"os_type,omitempty" xml:"os_type,omitempty"`
	// Deprecated
	//
	// **[Deprecated]**
	//
	// Subscription duration in months. Required and valid only when charge_type is PrePaid.
	//
	// Valid values: {1, 2, 3, 6, 12, 24, 36, 48, 60}.
	//
	// Default value: 1.
	//
	// This parameter changed on October 15, 2024. For more information, see [Announcement on Changes to CreateCluster API Parameters](https://help.aliyun.com/document_detail/2849194.html).
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// Deprecated
	//
	// **[Deprecated]**
	//
	// Billing cycle. Required when the billing type is PrePaid.
	//
	// Valid value: Month. Only monthly billing is supported.
	//
	// This parameter changed on October 15, 2024. For more information, see [Announcement on Changes to CreateCluster API Parameters](https://help.aliyun.com/document_detail/2849194.html).
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configuration, use the `platform` parameter under `scaling_group` in `nodepool`.
	//
	// Operating system distribution. Valid values:
	//
	// - CentOS
	//
	// - AliyunLinux
	//
	// - QbootAliyunLinux
	//
	// - Qboot
	//
	// - Windows
	//
	// - WindowsCore
	//
	// Default value: `CentOS`.
	//
	// example:
	//
	// CentOS
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- When using the Terway network plugin, specify virtual switches to assign IP addresses to pods. Each pod virtual switch corresponds to a worker node virtual switch, and both must be in the same zone.
	//
	// > We recommend that the pod virtual switch CIDR mask not exceed /19 and must not exceed /25. Otherwise, the number of assignable pod IP addresses becomes very limited, affecting normal cluster operation.
	PodVswitchIds []*string `json:"pod_vswitch_ids,omitempty" xml:"pod_vswitch_ids,omitempty" type:"Repeated"`
	// When you set `cluster_type` to `ManagedKubernetes` (ACK managed cluster), you can further specify the cluster subtype.
	//
	// - `Default`: ACK managed cluster, including ACK clusters (Pro Edition, Basic Edition).
	//
	// - `Edge`: ACK Edge cluster, including ACK Edge clusters (Pro Edition, Basic Edition).
	//
	// - `Serverless`: ACK serverless cluster, including ACK serverless clusters (Pro Edition, Basic Edition).
	//
	// - `Lingjun`: ACK LINGJUN cluster, available only in Pro Edition.
	//
	// example:
	//
	// Default
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
	// kube-proxy proxy mode
	//
	// - `iptables`: A mature and stable kube-proxy mode. Kubernetes Service discovery and load balancing use iptables rules. Performance is moderate and scales poorly with large numbers of Services. Suitable for clusters with few Services.
	//
	// - `ipvs`: A high-performance kube-proxy mode. Kubernetes Service discovery and load balancing use the Linux IPVS module. Suitable for clusters with many Services requiring high-performance load balancing.
	//
	// - `nftables`: A next-generation kube-proxy mode based on Linux nftables for Service discovery and load balancing. It is the modern replacement for iptables. Compared to iptables, nftables offers better network performance, faster rule updates, and superior scalability for large numbers of Services.<br>
	//
	//   Supported only for clusters running Kubernetes version 1.35 or later. The Kubernetes community deprecated IPVS starting in version 1.35. We recommend using nftables for new clusters to ensure long-term community support.<br>
	//
	// Default value: `ipvs`.
	//
	// example:
	//
	// ipvs
	ProxyMode *string `json:"proxy_mode,omitempty" xml:"proxy_mode,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configuration, use the `rds_instances` parameter under `scaling_group` in `nodepool`.
	//
	// List of RDS instances to add to the whitelist. We recommend adding the pod and node CIDR blocks of your container cluster to the RDS whitelist. Setting RDS instances might fail if they are not in the Running state.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// The region ID where the cluster is deployed. For more information, see [Regions supported by Container Service for Kubernetes](https://help.aliyun.com/document_detail/216938.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// The resource group ID to which the cluster belongs, used to isolate different resources.
	//
	// example:
	//
	// rg-acfm3mkrure****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// RRSA feature configuration.
	RrsaConfig *CreateClusterRequestRrsaConfig `json:"rrsa_config,omitempty" xml:"rrsa_config,omitempty" type:"Struct"`
	// Deprecated
	//
	// Container runtime for the cluster. Supports containerd, sandboxed containers, and Docker.
	//
	// > Kubernetes 1.24 no longer supports Docker as a built-in container runtime.
	//
	// For more information, see [Comparison of Docker, containerd, and sandboxed container runtimes](https://help.aliyun.com/document_detail/160313.html).
	Runtime *Runtime `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// Specify an existing security group ID when creating a cluster. Choose either this parameter or `is_enterprise_security_group`. Cluster nodes are automatically added to this security group.
	//
	// example:
	//
	// sg-bp1bdue0qc1g7k****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configuration, use the `security_hardening_os` parameter under `control_plane_config`. For node pool configuration, use the `security_hardening_os` parameter under `scaling_group` in `nodepool`.
	//
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
	// ServiceAccount is the access credential used by pods to communicate with the cluster API Server. The `service-account-issuer` is the issuer identity in the `serviceaccount token`, represented by the `iss` field in the `token payload`.
	//
	// For more information about `ServiceAccount`, see [Deploy service account token volume projection](https://help.aliyun.com/document_detail/160384.html).
	//
	// example:
	//
	// kubernetes.default.svc
	ServiceAccountIssuer *string `json:"service_account_issuer,omitempty" xml:"service_account_issuer,omitempty"`
	// Service network CIDR block. Valid ranges: 10.0.0.0/16-24, 172.16.0.0/16-24 to 172.31.0.0/16-24, 192.168.0.0/16-24.
	//
	// It cannot overlap with the VPC CIDR block 10.1.0.0/21 or existing Kubernetes cluster CIDR blocks in the VPC. This cannot be modified after cluster creation.
	//
	// Default value: 172.19.0.0/20.
	//
	// example:
	//
	// 172.21.0.0/20
	ServiceCidr *string `json:"service_cidr,omitempty" xml:"service_cidr,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Service discovery type for the cluster, used to specify the service discovery method in `ACK Serverless` clusters.
	//
	// - `CoreDNS`: Uses the standard Kubernetes service discovery component CoreDNS. Requires deploying a set of containers for DNS resolution. Defaults to two ECI instances with 0.25 vCPU and 512 MiB memory each.
	//
	// - `PrivateZone`: Uses Alibaba Cloud PrivateZone for service discovery. Requires enabling the PrivateZone service.
	//
	// Default value: Disabled.
	ServiceDiscoveryTypes []*string `json:"service_discovery_types,omitempty" xml:"service_discovery_types,omitempty" type:"Repeated"`
	// Configure SNAT for the VPC. Valid values:
	//
	// - `true`: Automatically create a NAT Gateway and configure SNAT rules. Set this to `true` if nodes or applications in the cluster need public network access.
	//
	// - `false`: Do not create a NAT Gateway or SNAT rules. Nodes and applications in the cluster cannot access the public network.
	//
	// > If you do not enable this during cluster creation but later need public network access, you can [enable it manually](https://help.aliyun.com/document_detail/178480.html).
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SnatEntry *bool `json:"snat_entry,omitempty" xml:"snat_entry,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane node configuration, use the `soc_enabled` parameter under `control_plane_config`. For node pool configuration, use the `soc_enabled` parameter under `scaling_group` in `nodepool`.
	//
	// MLPS 2.0 security hardening. For more information, see [Using MLPS 2.0 security hardening in ACK](https://help.aliyun.com/document_detail/196148.html).
	//
	// Valid values:
	//
	// - `true`: Enable MLPS 2.0 security hardening.
	//
	// - `false`: Disable MLPS 2.0 security hardening.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// Whether to enable public SSH logon. Used to log on to master nodes of ACK dedicated clusters. This parameter does not take effect for managed clusters.
	//
	// - `true`: Enable.
	//
	// - `false`: Disable.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	SshFlags *bool `json:"ssh_flags,omitempty" xml:"ssh_flags,omitempty"`
	// Node labels. Label rules:
	//
	// - Labels consist of case-sensitive key-value pairs. You can add up to 20 tags.
	//
	// - Tag keys must be unique and up to 64 characters long. Tag values can be empty and up to 128 characters long. Neither tag keys nor tag values can start with "aliyun", "acs:", "https\\://", or "http\\://". For more information, see [Labels and Selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set).
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configuration, use the `taints` parameter under `kubernetes_config` in `nodepool`.
	//
	// Node taint information. Taints and tolerations work together to prevent pods from being scheduled onto unsuitable nodes. For more information, see [taint-and-toleration](https://kubernetes.io/zh/docs/concepts/scheduling-eviction/taint-and-toleration/).
	Taints []*Taint `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- By default, clusters are not rolled back on creation failure. You must manually clean up failed clusters.
	//
	// Cluster creation timeout in minutes.
	//
	// Default value: `60`.
	//
	// example:
	//
	// 60
	TimeoutMins *int64 `json:"timeout_mins,omitempty" xml:"timeout_mins,omitempty"`
	// Time zone used by the cluster. See [Supported time zones](https://help.aliyun.com/document_detail/354879.html).
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// Custom cluster CA.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----****
	UserCa *string `json:"user_ca,omitempty" xml:"user_ca,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Custom node data.
	//
	// example:
	//
	// IyEvdXNyL2Jpbi9iYXNoCmVjaG8gIkhlbGxvIEFD****
	UserData *string `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// The VPC used by the cluster. You must provide this when creating a cluster.
	//
	// example:
	//
	// vpc-2zeik9h3ahvv2zz95****
	Vpcid *string `json:"vpcid,omitempty" xml:"vpcid,omitempty"`
	// Virtual switches for cluster nodes. This field is required when creating a zero-node managed cluster.
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configuration, use the `auto_renew` parameter under `scaling_group` in `nodepool`.
	//
	// Whether to enable auto-renewal for worker nodes. Valid only when `worker_instance_charge_type` is `PrePaid`. Valid values:
	//
	// - `true`: Enable auto-renewal.
	//
	// - `false`: Disable auto-renewal.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	WorkerAutoRenew *bool `json:"worker_auto_renew,omitempty" xml:"worker_auto_renew,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configuration, use the `auto_renew_period` parameter under `scaling_group` in `nodepool`.
	//
	// Auto-renewal period for worker nodes in months. Required and valid only when subscription billing is selected.
	//
	// Valid values: {1, 2, 3, 6, 12}.
	//
	// example:
	//
	// 1
	WorkerAutoRenewPeriod *int64 `json:"worker_auto_renew_period,omitempty" xml:"worker_auto_renew_period,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configuration, use the `data_disks` parameter under `scaling_group` in `nodepool`.
	//
	// Configuration for worker node data disks, including type and size.
	WorkerDataDisks []*CreateClusterRequestWorkerDataDisks `json:"worker_data_disks,omitempty" xml:"worker_data_disks,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configuration, use the `instance_charge_type` parameter under `scaling_group` in `nodepool`.
	//
	// Worker node billing type. Valid values:
	//
	// - `PrePaid`: Subscription.
	//
	// - `PostPaid`: Pay-as-you-go.
	//
	// Default value: Pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	WorkerInstanceChargeType *string `json:"worker_instance_charge_type,omitempty" xml:"worker_instance_charge_type,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configuration, use the `instance_types` parameter under `scaling_group` in `nodepool`.
	//
	// Worker node instance configuration.
	WorkerInstanceTypes []*string `json:"worker_instance_types,omitempty" xml:"worker_instance_types,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configuration, use the `period` parameter under `scaling_group` in `nodepool`.
	//
	// Subscription duration for worker nodes in months. Required and valid only when `worker_instance_charge_type` is `PrePaid`.
	//
	// Valid values: {1, 2, 3, 6, 12, 24, 36, 48, 60}.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	WorkerPeriod *int64 `json:"worker_period,omitempty" xml:"worker_period,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configuration, use the `period_unit` parameter under `scaling_group` in `nodepool`.
	//
	// Worker node billing cycle. Required when worker_instance_charge_type is `PrePaid`.
	//
	// Valid value: `Month`. Only monthly billing is supported.
	//
	// example:
	//
	// Month
	WorkerPeriodUnit *string `json:"worker_period_unit,omitempty" xml:"worker_period_unit,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configuration, use the `system_disk_category` parameter under `scaling_group` in `nodepool`.
	//
	// Worker node system disk type. For more information, see [Overview of Elastic Block Storage](https://help.aliyun.com/document_detail/63136.html).
	//
	// Valid values:
	//
	// - `cloud_efficiency`: Ultra disk.
	//
	// - `cloud_ssd`: Standard SSD.
	//
	// Default value: `cloud_ssd`.
	//
	// example:
	//
	// cloud_efficiency
	WorkerSystemDiskCategory *string `json:"worker_system_disk_category,omitempty" xml:"worker_system_disk_category,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configuration, use the `system_disk_performance_level` parameter under `scaling_group` in `nodepool`.
	//
	// When the system disk is an ESSD, you can set the performance level (PL). For more information, see [ESSD](https://help.aliyun.com/document_detail/122389.html).
	//
	// Valid values:
	//
	// - PL0
	//
	// - PL1
	//
	// - PL2
	//
	// - PL3
	//
	// example:
	//
	// PL1
	WorkerSystemDiskPerformanceLevel *string `json:"worker_system_disk_performance_level,omitempty" xml:"worker_system_disk_performance_level,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configuration, use the `system_disk_size` parameter under `scaling_group` in `nodepool`.
	//
	// Worker node system disk size in GiB.
	//
	// Valid range: [40,500].
	//
	// This value must be greater than or equal to max{40, ImageSize}.
	//
	// Default value: `120`.
	//
	// example:
	//
	// 120
	WorkerSystemDiskSize *int64 `json:"worker_system_disk_size,omitempty" xml:"worker_system_disk_size,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configuration, use the `system_disk_snapshot_policy_id` parameter under `scaling_group` in `nodepool`.
	//
	// Automatic snapshot policy ID for worker node system disks.
	//
	// example:
	//
	// sp-2zej1nogjvovnz4z****
	WorkerSystemDiskSnapshotPolicyId *string `json:"worker_system_disk_snapshot_policy_id,omitempty" xml:"worker_system_disk_snapshot_policy_id,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configuration, use the `vswitch_ids` parameter under `scaling_group` in `nodepool`.
	//
	// List of vSwitches used by cluster nodes. One node corresponds to one value.
	//
	// When creating a zero-node managed cluster, `worker_vswitch_ids` is optional, but you must provide `vswitch_ids`.
	WorkerVswitchIds []*string `json:"worker_vswitch_ids,omitempty" xml:"worker_vswitch_ids,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- Use the `zone_ids` parameter instead.
	//
	// The zone ID in the cluster region. This parameter applies only to ACK managed clusters.
	//
	// When creating an ACK managed cluster, if you do not specify `vpc_id` and `vswitch_ids`, you must specify `zone_id` to automatically create VPC network resources in this zone. This parameter is ignored if you specify `vpc_id` and `vswitch_ids`.
	//
	// example:
	//
	// cn-beiji****
	ZoneId *string `json:"zone_id,omitempty" xml:"zone_id,omitempty"`
	// A list of zone IDs in the cluster region. This parameter applies only to ACK managed clusters.
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

func (s *CreateClusterRequest) GetControlPlaneEndpointsConfig() *CreateClusterRequestControlPlaneEndpointsConfig {
	return s.ControlPlaneEndpointsConfig
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

func (s *CreateClusterRequest) SetControlPlaneEndpointsConfig(v *CreateClusterRequestControlPlaneEndpointsConfig) *CreateClusterRequest {
	s.ControlPlaneEndpointsConfig = v
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
	if s.Addons != nil {
		for _, item := range s.Addons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AuditLogConfig != nil {
		if err := s.AuditLogConfig.Validate(); err != nil {
			return err
		}
	}
	if s.AutoMode != nil {
		if err := s.AutoMode.Validate(); err != nil {
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
	if s.Nodepools != nil {
		for _, item := range s.Nodepools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OperationPolicy != nil {
		if err := s.OperationPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.RrsaConfig != nil {
		if err := s.RrsaConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Taints != nil {
		for _, item := range s.Taints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WorkerDataDisks != nil {
		for _, item := range s.WorkerDataDisks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateClusterRequestAuditLogConfig struct {
	// Whether to enable cluster audit logging.
	//
	// - true: Enable.
	//
	// - false: Disable.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The [SLS Project](https://help.aliyun.com/document_detail/48873.html) containing the [Logstore](https://help.aliyun.com/document_detail/48873.html) for cluster audit logs.
	//
	// - Default value: `k8s-log-{clusterid}`.
	//
	// - When audit logging is enabled, a Logstore for audit logs is created in the specified SLS Project.
	//
	// example:
	//
	// k8s-log-c2345xxxxxxxxxxxx
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
	// Whether to enable Intelligent Managed Mode.
	//
	// - true: Enable.
	//
	// - false: Disable.
	//
	// example:
	//
	// true
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
	// Whether to enable auto-renewal for control plane nodes. Valid only when charge_type is `PrePaid`.
	//
	// - true: Enable auto-renewal.
	//
	// - false: Disable auto-renewal.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// Auto-renewal duration for control plane nodes in months.
	//
	// Valid values: {1, 2, 3, 6, 12}.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// Control plane node billing type.
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
	// Whether to install Cloud Monitor on nodes.
	//
	// - true: Install the CloudMonitor agent.
	//
	// - false: Do not install the CloudMonitor agent.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// Node CPU management policy.
	//
	// - static: Enhances CPU affinity and exclusivity for pods with specific resource characteristics on the node.
	//
	// - none: Uses the default CPU affinity scheme.
	//
	// Default value: none.
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
	// Image ID.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20240819.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// Operating system image type.
	//
	// example:
	//
	// AliyunLinux3
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// ECS instance metadata access configuration.
	InstanceMetadataOptions *InstanceMetadataOptions `json:"instance_metadata_options,omitempty" xml:"instance_metadata_options,omitempty"`
	// Node instance types.
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// Key pair name. Choose either this parameter or login_password.
	//
	// example:
	//
	// ack
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// SSH login password. Password rules: 8–30 characters, containing at least three of the following: uppercase letters, lowercase letters, digits, and special characters. Choose either this parameter or key_pair.
	//
	// example:
	//
	// ********
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// **[Deprecated]*	- Node service port range.
	//
	// example:
	//
	// 30000-32767
	NodePortRange *string `json:"node_port_range,omitempty" xml:"node_port_range,omitempty"`
	// Subscription duration for control plane nodes in months. Required and valid only when charge_type is `PrePaid`.
	//
	// Valid values: {1, 2, 3, 6, 12, 24, 36, 48, 60}.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// Billing cycle unit for control plane nodes. Required and valid only when charge_type is `PrePaid`.
	//
	// Valid value: `Month`. Only monthly billing is supported.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// **[Deprecated]*	- Control plane node runtime name. Valid value:
	//
	// containerd: Containerd runtime, supported for all cluster versions.
	//
	// example:
	//
	// containerd
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// Whether to enable Alibaba Cloud OS security hardening.
	//
	// - true: Enable Alibaba Cloud OS security hardening.
	//
	// - false: Disable Alibaba Cloud OS security hardening.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	SecurityHardeningOs *bool `json:"security_hardening_os,omitempty" xml:"security_hardening_os,omitempty"`
	// Number of control plane nodes.
	//
	// Valid values: `3` or `5`.
	//
	// example:
	//
	// 3
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// Whether to enable MLPS 2.0 security hardening.
	//
	// - true: Enable MLPS 2.0 security hardening.
	//
	// - false: Disable MLPS 2.0 security hardening.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// Whether to enable performance burst for node system disks.
	//
	// - true: Enable.
	//
	// - false: Disable.
	//
	// This parameter is supported only when `system_disk_category` is `cloud_auto`.
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// Node system disk type.
	//
	// - `cloud_efficiency`: Ultra disk.
	//
	// - `cloud_ssd`: Standard SSD.
	//
	// - `cloud_essd`: ESSD.
	//
	// - `cloud_auto`: ESSD AutoPL.
	//
	// - `cloud_essd_entry`: ESSD Entry.
	//
	// Default value: `cloud_ssd`. The default may vary by zone.
	//
	// example:
	//
	// cloud_ssd
	SystemDiskCategory *string `json:"system_disk_category,omitempty" xml:"system_disk_category,omitempty"`
	// Node system disk performance level. Applies only to ESSD disks.
	//
	// Disk performance levels depend on disk size. For more information, see [ESSD](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// Provisioned read/write IOPS for node system disks.
	//
	// Valid range: 0 to min{50,000, 1000 × capacity - baseline performance}. Baseline performance = min{1,800 + 50 × capacity, 50,000}.
	//
	// This parameter is supported only when `system_disk_category` is `cloud_auto`.
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// Node system disk size in GiB.
	//
	// Valid range: [40,500].
	//
	// Default value: `120`.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// Automatic snapshot backup policy for node system disks.
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
	if s.InstanceMetadataOptions != nil {
		if err := s.InstanceMetadataOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateClusterRequestControlPlaneEndpointsConfig struct {
	// Internal domain name configuration for the cluster, applicable to ACK managed clusters. Internal domain names allow node-side system components such as kubelet and kube-proxy to access the API Server. Without internal domain name access, node-side components access the API Server through the CLB IP address.
	InternalDnsConfig *CreateClusterRequestControlPlaneEndpointsConfigInternalDnsConfig `json:"internal_dns_config,omitempty" xml:"internal_dns_config,omitempty" type:"Struct"`
}

func (s CreateClusterRequestControlPlaneEndpointsConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestControlPlaneEndpointsConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestControlPlaneEndpointsConfig) GetInternalDnsConfig() *CreateClusterRequestControlPlaneEndpointsConfigInternalDnsConfig {
	return s.InternalDnsConfig
}

func (s *CreateClusterRequestControlPlaneEndpointsConfig) SetInternalDnsConfig(v *CreateClusterRequestControlPlaneEndpointsConfigInternalDnsConfig) *CreateClusterRequestControlPlaneEndpointsConfig {
	s.InternalDnsConfig = v
	return s
}

func (s *CreateClusterRequestControlPlaneEndpointsConfig) Validate() error {
	if s.InternalDnsConfig != nil {
		if err := s.InternalDnsConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateClusterRequestControlPlaneEndpointsConfigInternalDnsConfig struct {
	// VPCs where the internal domain name resolution takes effect.
	BindVpcs []*string `json:"bind_vpcs,omitempty" xml:"bind_vpcs,omitempty" type:"Repeated"`
}

func (s CreateClusterRequestControlPlaneEndpointsConfigInternalDnsConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestControlPlaneEndpointsConfigInternalDnsConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestControlPlaneEndpointsConfigInternalDnsConfig) GetBindVpcs() []*string {
	return s.BindVpcs
}

func (s *CreateClusterRequestControlPlaneEndpointsConfigInternalDnsConfig) SetBindVpcs(v []*string) *CreateClusterRequestControlPlaneEndpointsConfigInternalDnsConfig {
	s.BindVpcs = v
	return s
}

func (s *CreateClusterRequestControlPlaneEndpointsConfigInternalDnsConfig) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestOperationPolicy struct {
	// Cluster automatic upgrade.
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
	if s.ClusterAutoUpgrade != nil {
		if err := s.ClusterAutoUpgrade.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateClusterRequestOperationPolicyClusterAutoUpgrade struct {
	// Cluster automatic upgrade frequency. Valid values:
	//
	// - patch: Automatically upgrade to the latest patch version within the current minor version. New Kubernetes versions do not include breaking changes.
	//
	// - stable: Automatically upgrade to the latest patch version of the second-newest minor version. New Kubernetes versions may include API and feature changes but have undergone extensive stability validation.
	//
	// - rapid: Automatically upgrade to the latest patch version of the newest minor version to quickly access new Kubernetes community features.
	//
	// example:
	//
	// stable
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// Whether to enable cluster automatic upgrade.
	//
	// - true: Enable.
	//
	// - false: Disable.
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
	// Whether to enable RRSA.
	//
	// - true: Enable.
	//
	// - false: Disable.
	//
	// example:
	//
	// true
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
	// Data disk type.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// Whether to encrypt the data disk. Valid values:
	//
	// - `true`: Encrypt the data disk.
	//
	// - `false`: Do not encrypt the data disk.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	Encrypted *string `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// Data disk performance level. Applies only to [ESSD](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"performance_level,omitempty" xml:"performance_level,omitempty"`
	// Data disk size in GiB. Valid range: 40–32767.
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
