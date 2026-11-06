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
	SetClientToken(v string) *CreateClusterRequest
	GetClientToken() *string
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
	SetSupportPolicy(v string) *CreateClusterRequest
	GetSupportPolicy() *string
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
	// **[Deprecated]*	- The access control list (ACL) for the Server Load Balancer (SLB) instance of the API server in a registered cluster.
	AccessControlList []*string `json:"access_control_list,omitempty" xml:"access_control_list,omitempty" type:"Repeated"`
	// The list of cluster components. You can use `addons` to specify the cluster components that you want to install when you create a cluster.
	//
	// **Network components**: Required. Includes two options for the network type: Flannel and Terway. You must select one of them when you create a cluster:
	//
	// - Flannel: [{"name":"flannel","config":""}\\].
	//
	// - Terway: [{"name": "terway-eniip","config": ""}\\].
	//
	// **Storage components**: Optional. Only `csi` is supported:
	//
	// `csi`: [{"name":"csi-plugin","config": ""},{"name": "csi-provisioner","config": ""}\\].
	//
	// **Log components**: Optional. We recommend that you enable this component. If you do not enable Log Service, you cannot use the cluster audit feature.
	//
	// - Use an existing `SLS Project`: [{"name": "loongcollector","config": "{\\"IngressDashboardEnabled\\":\\"true\\",\\"sls_project_name\\":\\"your_sls_project_name\\"}"}\\].
	//
	// - Create a new `SLS Project`: [{"name": "loongcollector","config": "{\\"IngressDashboardEnabled\\":\\"true\\"}"}\\].
	//
	// **Ingress components**: Optional. By default, the `nginx-ingress-controller` Ingress component is installed in ACK dedicated clusters.
	//
	// - Install the Ingress component and enable Internet access: [{"name":"nginx-ingress-controller","config":"{\\"IngressSlbNetworkType\\":\\"internet\\"}"}\\].
	//
	// - Disable the default installation of the Ingress component: [{"name": "nginx-ingress-controller","config": "","disabled": true}\\].
	//
	// **Event center**: Optional. Enabled by default.
	//
	// The event center provides capabilities such as storage, query, and alerting for Kubernetes events. The Logstore associated with the Kubernetes event center is free of charge for 90 days. For more information about the free quota policy, see [Create and use a Kubernetes event center](https://help.aliyun.com/document_detail/150476.html).
	//
	// Example of enabling the event center: [{"name":"ack-node-problem-detector","config":"{\\"sls_project_name\\":\\"your_sls_project_name\\"}"}\\].
	Addons []*Addon `json:"addons,omitempty" xml:"addons,omitempty" type:"Repeated"`
	// A ServiceAccount is the access credential for communication between pods and the cluster API server. `api-audiences` specifies the valid identities for request `tokens` and is used by the `apiserver` to authenticate whether a request `token` is valid. You can configure multiple `audiences` and separate them with commas (,).
	//
	// For more information about `ServiceAccount`, see [Deploy service account token volume projection](https://help.aliyun.com/document_detail/160384.html).
	//
	// example:
	//
	// kubernetes.default.svc
	ApiAudiences *string `json:"api_audiences,omitempty" xml:"api_audiences,omitempty"`
	// The cluster audit log configurations.
	AuditLogConfig *CreateClusterRequestAuditLogConfig `json:"audit_log_config,omitempty" xml:"audit_log_config,omitempty" type:"Struct"`
	// The [intelligent managed mode](https://help.aliyun.com/document_detail/2938898.html) configurations.
	AutoMode *CreateClusterRequestAutoMode `json:"auto_mode,omitempty" xml:"auto_mode,omitempty" type:"Struct"`
	// Deprecated
	//
	// **[Deprecated]*	- Specifies whether to enable auto-renewal. This parameter takes effect only when `charge_type` is set to `PrePaid`. Valid values:
	//
	// - `true`: Enables auto-renewal.
	//
	// - `false`: Disables auto-renewal.
	//
	// Default value: `false`.
	//
	// This parameter was changed on October 15, 2024. For more information, see [Announcement on parameter behavior changes of the CreateCluster operation for cluster creation](https://help.aliyun.com/document_detail/2849194.html).
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The auto-renewal cycle. This parameter takes effect only when you select subscription and auto-renewal. When `PeriodUnit` is set to `Month`, valid values are 1, 2, 3, 6, and 12.
	//
	// Default value: 1.
	//
	// This parameter was changed on October 15, 2024. For more information, see [Announcement on parameter behavior changes of the CreateCluster operation for cluster creation](https://help.aliyun.com/document_detail/2849194.html).
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The billing method of the Classic Load Balancer (CLB) instance used by the API server. Default value: PostPaid. Valid values:
	//
	// - PostPaid: Pay-as-you-go.
	//
	// - PrePaid: Subscription. This billing method is no longer supported for new CLB instances, but existing instances are not affected.
	//
	// 	Notice:
	//
	// - This parameter was changed on October 15, 2024. For more information, see [Announcement on parameter behavior changes of the CreateCluster operation for cluster creation](https://help.aliyun.com/document_detail/2849194.html).
	//
	// - Starting December 1, 2024, newly created CLB instances no longer support the subscription billing method, and an instance fee is charged.
	//
	// </notice>
	//
	// <props="china">For more information, see [Service notices: Deprecation of the subscription billing method for the CLB instance of the API server in newly created clusters](https://help.aliyun.com/document_detail/2851191.html) and [Announcement on billing item adjustments for Classic Load Balancer (CLB)](https://help.aliyun.com/document_detail/2839797.html).
	//
	// <props="intl">For more information, see [Announcement on billing item adjustments for Classic Load Balancer (CLB)](https://help.aliyun.com/document_detail/2839797.html).
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"charge_type,omitempty" xml:"charge_type,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- To configure the cluster control plane, use the `security_hardening_os` parameter under `control_plane_config` instead. To configure node pools, use the `security_hardening_os` parameter under `scaling_group` in `nodepool` instead.
	//
	// example:
	//
	// false
	CisEnabled *bool `json:"cis_enabled,omitempty" xml:"cis_enabled,omitempty"`
	// The client token.
	//
	// example:
	//
	// 0a287651-1277-4c0d-80d2-623a4*****
	ClientToken *string `json:"client_token,omitempty" xml:"client_token,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- To configure control plane nodes of a cluster, use the `cloud_monitor_flags` parameter under `control_plane_config` instead. To configure node pools, use the `cms_enabled` parameter under `kubernetes_config` in `nodepool` instead.
	//
	// Specifies whether to install the CloudMonitor agent on the cluster. Valid values:
	//
	// - `true`: Installs the CloudMonitor agent.
	//
	// - `false`: Does not install the CloudMonitor agent.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// The local domain name of the cluster.
	//
	// Naming rules: The domain name consists of one or more parts separated by periods (.). Each part can be up to 63 characters in length and can contain lowercase letters, digits, and hyphens (-). It must start and end with a lowercase letter or a digit.
	//
	// example:
	//
	// cluster.local
	ClusterDomain *string `json:"cluster_domain,omitempty" xml:"cluster_domain,omitempty"`
	// After you set `cluster_type` to `ManagedKubernetes` and configure `profile`, you can further specify the cluster edition. Valid values:
	//
	// - `ack.standard`: Basic Edition. This is the default value if this parameter is left empty.
	//
	// - `ack.pro.small`: Pro Edition.
	//
	// - `ack.pro.xlarge`: Pro XL.
	//
	// - `ack.pro.2xlarge`: Pro 2XL.
	//
	// - `ack.pro.4xlarge`: Pro 4XL. You must contact customer service to apply for whitelist access.
	//
	// Pro XL, Pro 2XL, and Pro 4XL are three tiers provided by <props="china">[ACK Pro Provisioned Control Plane](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane)<props="intl">[ACK Pro Provisioned Control Plane](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane). By pre-allocating and reserving control plane resources, these editions ensure that API concurrency and pod scheduling capabilities are always maintained at a consistently high level. They are suitable for AI training and inference, ultra-large-scale clusters, and mission-critical workloads.
	//
	// For more information about the cluster management fees for Pro Edition and provisioned control plane editions, see <props="china">[Cluster management fees](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee)<props="intl">[Cluster management fees](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee).
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// Valid values:
	//
	// - `Kubernetes`: ACK dedicated cluster.
	//
	// - `ManagedKubernetes`: ACK managed cluster, including ACK managed clusters (Pro Edition and Basic Edition), ACK Serverless clusters (Pro Edition and Basic Edition), ACK Edge clusters (Pro Edition and Basic Edition), and ACK Lingjun clusters (Pro Edition).
	//
	// - `ExternalKubernetes`: Registered cluster.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// The CIDR block for pods. It must be a valid private CIDR block, specifically one of the following CIDR blocks or their subnets: 10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16. It cannot overlap with the CIDR block of the VPC or the CIDR blocks used by existing Kubernetes clusters in the VPC. It cannot be modified after the cluster is created.
	//
	// For more information about cluster network planning, see [Network planning for ACK managed clusters](https://help.aliyun.com/document_detail/86500.html).
	//
	// > This parameter is required for Flannel clusters.
	//
	// example:
	//
	// 172.20.0.0/16
	ContainerCidr *string `json:"container_cidr,omitempty" xml:"container_cidr,omitempty"`
	// The control plane configuration for an ACK dedicated cluster.
	ControlPlaneConfig *CreateClusterRequestControlPlaneConfig `json:"control_plane_config,omitempty" xml:"control_plane_config,omitempty" type:"Struct"`
	// The connection configuration of the cluster. If you specify this parameter, the endpoint_public_access and load_balancer_id parameters do not take effect. To specify a CLB or NLB instance, use load_balancers_config to specify the corresponding instance ID.
	ControlPlaneEndpointsConfig *CreateClusterRequestControlPlaneEndpointsConfig `json:"control_plane_endpoints_config,omitempty" xml:"control_plane_endpoints_config,omitempty" type:"Struct"`
	// The list of component names. This parameter specifies the control plane components whose logs are collected. By default, logs of kube-apiserver, kube-controller-manager, kube-scheduler, and cloud-controller-manager are collected.
	ControlplaneLogComponents []*string `json:"controlplane_log_components,omitempty" xml:"controlplane_log_components,omitempty" type:"Repeated"`
	// The Log Service project for control plane component logs. You can use an existing project for log storage, or use the automatic creation feature to create a project. If you choose automatic creation, a Log Service project named k8s-log-{ClusterID} is created.
	//
	// example:
	//
	// k8s-log-xxx
	ControlplaneLogProject *string `json:"controlplane_log_project,omitempty" xml:"controlplane_log_project,omitempty"`
	// The log retention period for control plane components. Unit: days.
	//
	// example:
	//
	// 30
	ControlplaneLogTtl *string `json:"controlplane_log_ttl,omitempty" xml:"controlplane_log_ttl,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use the cpu_policy parameter under control_plane_config for cluster control plane configuration, and use the cpu_policy parameter under kubernetes_config in nodepool for node pool configuration.
	//
	// The CPU management policy for nodes. This parameter supports the following policies when the cluster version is 1.12.6 or later:
	//
	// - static: allows you to enhance CPU affinity and exclusivity for pods with specific resource characteristics on the node.
	//
	// - none: uses the default CPU affinity policy.
	//
	// Default value: none.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use the extra_sans parameter instead.
	//
	// The SANs of the custom certificate. Separate multiple IP addresses or domain names with commas (,).
	//
	// example:
	//
	// cs.aliyun.com
	CustomSan *string `json:"custom_san,omitempty" xml:"custom_san,omitempty"`
	// Specifies whether to enable deletion protection for the cluster to prevent accidental deletion through the console or API calls. Valid values:
	//
	// - `true`: Enables cluster deletion protection. The cluster cannot be deleted through the console or API.
	//
	// - `false`: Disables cluster deletion protection. The cluster can be deleted through the console or API.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- By default, no rollback is performed when cluster creation fails. You must manually clean up the failed cluster.
	//
	// Specifies whether to perform a rollback when cluster creation fails. Valid values:
	//
	// - `true`: Performs a rollback when cluster creation fails.
	//
	// - `false`: Does not perform a rollback when cluster creation fails.
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
	// Specifies whether to enable RRSA. Valid values:
	//
	// - true: Enables RRSA.
	//
	// - false: Disables RRSA.
	//
	// example:
	//
	// false
	EnableRrsa *bool `json:"enable_rrsa,omitempty" xml:"enable_rrsa,omitempty"`
	// The ID of the KMS key used to encrypt data cloud disks. For more information, see [Key Management Service](https://help.aliyun.com/document_detail/28935.html).
	//
	// > This feature takes effect only in ACK Pro clusters.
	//
	// example:
	//
	// 0fe64791-55eb-4fc7-84c5-c6c7cdca****
	EncryptionProviderKey *string `json:"encryption_provider_key,omitempty" xml:"encryption_provider_key,omitempty"`
	// Specifies whether to enable public network access. An elastic IP address (EIP) is used to expose the API Server for public network access to the cluster. Valid values:
	//
	// - `true`: Enables public network access.
	//
	// - `false`: Disables public network access. If you select this value, the cluster API Server cannot be accessed over the Internet.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	EndpointPublicAccess *bool `json:"endpoint_public_access,omitempty" xml:"endpoint_public_access,omitempty"`
	// The custom Subject Alternative Name (SAN) of the API Server certificate.
	ExtraSans []*string `json:"extra_sans,omitempty" xml:"extra_sans,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- You cannot select existing nodes when creating a cluster. To add existing nodes to a cluster, create a node pool first and call the [AttachInstancesToNodePool](https://help.aliyun.com/document_detail/2667920.html) operation.
	//
	// Specifies whether to mount data cloud disks to existing instances when creating a cluster. Valid values:
	//
	// - `true`: Stores containers and images on data cloud disks. Existing data on the data cloud disks will be lost. Back up your data first.
	//
	// - `false`: Does not store containers and images on data cloud disks.
	//
	// Default value: `false`.
	//
	// Data cloud disk mounting rules:
	//
	// - If the ECS instance has data cloud disks attached and the file system of the last data cloud disk is not initialized, the system automatically formats the data cloud disk to ext4 to store /var/lib/docker and /var/lib/kubelet.
	//
	// - If the ECS instance does not have data cloud disks attached, no new data cloud disk will be attached.
	//
	// example:
	//
	// false
	FormatDisk *bool `json:"format_disk,omitempty" xml:"format_disk,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For cluster control plane configurations, use the `image_id` parameter under `control_plane_config` instead. For node pool configurations, use the `image_id` parameter under `scaling_group` in `nodepool` instead.
	//
	// The custom image for nodes. By default, the system image is used. If you select a custom image, it replaces the default system image. For more information, see [Custom images](https://help.aliyun.com/document_detail/146647.html).
	//
	// example:
	//
	// m-bp16z7xko3vvv8gt****
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For cluster control plane configurations, use the `image_type` parameter under `control_plane_config` instead. For node pool configurations, use the `image_type` parameter under `scaling_group` in `nodepool` instead.
	//
	// The operating system distribution type. We recommend that you use this field to specify the node operating system. Valid values:
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
	// **[Deprecated]*	- You cannot select existing nodes when creating a cluster. To add existing nodes to a cluster, create a node pool first and call the [AttachInstancesToNodePool](https://help.aliyun.com/document_detail/2667920.html) operation.
	//
	// The list of existing ECS instances to be added as worker nodes when creating a cluster with existing nodes.
	//
	// > This field is required when you create a cluster with existing instances.
	Instances []*string `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// The IP stack of the cluster.
	//
	// example:
	//
	// Valid values:
	//
	// ipv4 (single-stack)
	//
	// dual (dual-stack). Default value: ipv4
	IpStack *string `json:"ip_stack,omitempty" xml:"ip_stack,omitempty"`
	// Specifies whether to enable automatic creation of an advanced security group. This parameter takes effect when `security_group_id` is empty.
	//
	// > If you use a basic security group, the total number of nodes and Terway pods in the cluster cannot exceed 2,000. Therefore, we recommend that you use an advanced security group when creating a cluster of the Terway network type.
	//
	// Valid values:
	//
	// - `true`: Creates and uses an advanced security group.
	//
	// - `false`: Uses a basic security group.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	IsEnterpriseSecurityGroup *bool `json:"is_enterprise_security_group,omitempty" xml:"is_enterprise_security_group,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- You cannot select existing nodes when creating a cluster. To add existing nodes to a cluster, create a node pool first and call the [AttachInstancesToNodePool](https://help.aliyun.com/document_detail/2667920.html) operation.
	//
	// Specifies whether to retain the instance names when creating a cluster with existing instances. Valid values:
	//
	// - `true`: Retains the instance names.
	//
	// - `false`: Does not retain the instance names. The names are replaced based on system rules.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	KeepInstanceName *bool `json:"keep_instance_name,omitempty" xml:"keep_instance_name,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For cluster control plane configurations, use the `key_pair` parameter under `control_plane_config` instead. For node pool configurations, use the `key_pair` parameter under `scaling_group` in `nodepool` instead.
	//
	// The name of the key pair. You must specify either this parameter or `login_password`.
	//
	// example:
	//
	// security-key
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// The cluster version, which is consistent with the Kubernetes community baseline version. We recommend that you select the latest version. If you do not specify this parameter, the latest version is used by default.
	//
	// You can create clusters of the latest three versions. You can call the [DescribeKubernetesVersionMetadata](https://help.aliyun.com/document_detail/2667899.html) operation to query the supported cluster versions.
	//
	// For more information about the Kubernetes versions supported by ACK, see [Kubernetes version release overview](https://help.aliyun.com/document_detail/185269.html).
	//
	// example:
	//
	// 1.32.1-aliyun.1
	KubernetesVersion *string `json:"kubernetes_version,omitempty" xml:"kubernetes_version,omitempty"`
	// The instance ID of the Classic Load Balancer (CLB) used for API server access. If you specify this parameter, automatic creation of a CLB instance for the API server is disabled.
	//
	// > Make sure that the CLB instance has no other dependencies, such as listeners or backend servers. Shared-resource and Internet-facing CLB instances are not supported.
	//
	// example:
	//
	// lb-wz9t256gqa3vbouk****
	LoadBalancerId *string `json:"load_balancer_id,omitempty" xml:"load_balancer_id,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For CLB instances, billing is based on usage. This parameter does not take effect.
	//
	// The load balancing specification. Valid values:
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
	// **[Deprecated]*	- This parameter enables Simple Log Service for the cluster. It takes effect only for ACK Serverless clusters and must be set to `SLS`.
	//
	// example:
	//
	// SLS
	LoggingType *string `json:"logging_type,omitempty" xml:"logging_type,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For cluster control plane configurations, use the `login_password` parameter under `control_plane_config` instead. For node pool configurations, use the `login_password` parameter under `scaling_group` in `nodepool` instead.
	//
	// The SSH logon password. You must specify either this parameter or `key_pair`. The password must be 8 to 30 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// example:
	//
	// null
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The maintenance window of the cluster.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window,omitempty" xml:"maintenance_window,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For cluster control plane configurations, use the `auto_renew` parameter under `control_plane_config` instead.
	//
	// Specifies whether to enable auto-renewal for master nodes. This parameter takes effect only when `master_instance_charge_type` is set to `PrePaid`. Valid values:
	//
	// - `true`: Enables auto-renewal.
	//
	// - `false`: Disables auto-renewal.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	MasterAutoRenew *bool `json:"master_auto_renew,omitempty" xml:"master_auto_renew,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For cluster control plane configurations, use the `auto_renew_period` parameter under `control_plane_config` instead.
	//
	// The auto-renewal epoch for master nodes. This parameter takes effect and is required only when the subscription billing method is selected.
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
	// **[Deprecated]*	- For cluster control plane configurations, use the `size` parameter under `control_plane_config` instead.
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
	// **[Deprecated]*	- For cluster control plane configurations, use the `instance_charge_type` parameter under `control_plane_config` instead.
	//
	// The billing method of master nodes. Valid values:
	//
	// - `PrePaid`: subscription.
	//
	// - `PostPaid`: pay-as-you-go.
	//
	// Default value: `PostPaid`.
	//
	// example:
	//
	// PrePaid
	MasterInstanceChargeType *string `json:"master_instance_charge_type,omitempty" xml:"master_instance_charge_type,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For cluster control plane configurations, use the `instance_types` parameter under `control_plane_config` instead.
	//
	// The instance types of master nodes. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	MasterInstanceTypes []*string `json:"master_instance_types,omitempty" xml:"master_instance_types,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- For cluster control plane configurations, use the `unit` parameter under `control_plane_config` instead.
	//
	// The subscription period of master nodes. This parameter takes effect and is required only when `master_instance_charge_type` is set to `PrePaid`.
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
	// **[Deprecated]*	- For cluster control plane configurations, use the `period_unit` parameter under `control_plane_config` instead.
	//
	// The subscription period unit of master nodes. This parameter is required when the billing method is `PrePaid`.
	//
	// Valid value: `Month`. Currently, only month is supported.
	//
	// example:
	//
	// Month
	MasterPeriodUnit *string `json:"master_period_unit,omitempty" xml:"master_period_unit,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For cluster control plane configurations, use the `system_disk_category` parameter under `control_plane_config` instead.
	//
	// The system cloud disk category of master nodes. Valid values:
	//
	// - `cloud_efficiency`: ultra disk.
	//
	// - `cloud_ssd`: standard SSD.
	//
	// - `cloud_essd`: ESSD.
	//
	// Default value: `cloud_ssd`. The default value may vary by zone.
	//
	// example:
	//
	// cloud_ssd
	MasterSystemDiskCategory *string `json:"master_system_disk_category,omitempty" xml:"master_system_disk_category,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For cluster control plane configurations, use the `system_disk_performance_level` parameter under `control_plane_config` instead.
	//
	// The performance level of the system cloud disk for master nodes. This parameter takes effect only for ESSDs. The performance level is related to the disk size. For more information, see [ESSD standard SSD](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	MasterSystemDiskPerformanceLevel *string `json:"master_system_disk_performance_level,omitempty" xml:"master_system_disk_performance_level,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For cluster control plane configurations, use the `system_disk_size` parameter under `control_plane_config` instead.
	//
	// The system cloud disk size of master nodes. Valid values: 40 to 500. Unit: GiB.
	//
	// Default value: `120`.
	//
	// example:
	//
	// 120
	MasterSystemDiskSize *int64 `json:"master_system_disk_size,omitempty" xml:"master_system_disk_size,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For cluster control plane configurations, use the `system_disk_snapshot_policy_id` parameter under `control_plane_config` instead.
	//
	// The ID of the automatic snapshot policy applied to the system cloud disk of master nodes.
	//
	// example:
	//
	// sp-2zej1nogjvovnz4z****
	MasterSystemDiskSnapshotPolicyId *string `json:"master_system_disk_snapshot_policy_id,omitempty" xml:"master_system_disk_snapshot_policy_id,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use the `vswitch_ids` parameter instead.
	//
	// The list of vSwitch IDs for master nodes. You can specify 1 to 3 vSwitches. To ensure high availability of the cluster, we recommend that you select three vSwitches in different zones.
	//
	// The number of specified vSwitches must be consistent with the value of `master_count` and correspond to the elements in `master_vswitch_ids` one by one.
	MasterVswitchIds []*string `json:"master_vswitch_ids,omitempty" xml:"master_vswitch_ids,omitempty" type:"Repeated"`
	// The custom name of the cluster. The name must be 1 to 63 characters in length and can contain digits, Chinese characters, English letters, and hyphens (-). It cannot start with a hyphen (-).
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
	// The maximum number of IP addresses that can be assigned to nodes. This parameter determines the number of IP addresses by specifying the CIDR block of the network. It takes effect only for clusters of the Flannel network type.
	//
	// Default value: `26`.
	//
	// example:
	//
	// 25
	NodeCidrMask *string `json:"node_cidr_mask,omitempty" xml:"node_cidr_mask,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configurations, use the `node_name_mode` parameter under `kubernetes_config` in `nodepool` instead.
	//
	// example:
	//
	// null
	NodeNameMode *string `json:"node_name_mode,omitempty" xml:"node_name_mode,omitempty"`
	// Deprecated
	//
	// The node service port. Valid port range: 30000 to 65535.
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
	// **[Deprecated]*	- For node pool configurations, use the `desired_size` parameter under `scaling_group` in `nodepool` instead. The number of worker nodes. Valid values: 0 to 100.
	//
	// example:
	//
	// 3
	NumOfNodes *int64 `json:"num_of_nodes,omitempty" xml:"num_of_nodes,omitempty"`
	// The automatic O&M policy for the cluster.
	OperationPolicy *CreateClusterRequestOperationPolicy `json:"operation_policy,omitempty" xml:"operation_policy,omitempty" type:"Struct"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane node configurations, use the `image_type` parameter under `control_plane_config` instead. For node pool configurations, use the `image_type` parameter under `scaling_group` in `nodepool` instead. The operating system platform type. Valid values:
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
	// **[Deprecated]*	- The subscription duration. This parameter takes effect and is required only when `charge_type` is set to `PrePaid`. Valid values: 1, 2, 3, 6, 12, 24, 36, 48, and 60. Default value: 1. This field was changed on October 15, 2024. For more information, see [Announcement on parameter behavior changes of the CreateCluster API operation](https://help.aliyun.com/document_detail/2849194.html).
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The billing cycle. This parameter is required when the billing method is `PrePaid`. Valid value: `Month`. Currently, only monthly billing is supported. This field was changed on October 15, 2024. For more information, see [Announcement on parameter behavior changes of the CreateCluster API operation](https://help.aliyun.com/document_detail/2849194.html).
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configurations, use the `platform` parameter under `scaling_group` in `nodepool` instead. The operating system distribution. Valid values:
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
	// **[Deprecated]*	- When the Terway network plug-in is selected, you must specify the vSwitches used to assign IP addresses to pods. Each pod vSwitch corresponds to a worker node vSwitch, and they must be in the same zone.
	//
	// > We recommend that the subnet mask length of the pod vSwitch does not exceed 19, and the maximum subnet mask length does not exceed 25. Otherwise, the number of pod IP addresses that can be allocated in the cluster network is very limited, which affects the normal use of the cluster.
	PodVswitchIds []*string `json:"pod_vswitch_ids,omitempty" xml:"pod_vswitch_ids,omitempty" type:"Repeated"`
	// The cluster subtype when `cluster_type` is set to `ManagedKubernetes`, which indicates an ACK managed cluster. Valid values:
	//
	// - `Default`: ACK managed clusters, including ACK Pro and ACK ACK clusters.
	//
	// - `Edge`: ACK Edge clusters, including ACK Edge Pro and ACK Edge ACK clusters.
	//
	// - `Serverless`: ACK Serverless clusters, including ACK Serverless Pro and ACK Serverless ACK clusters.
	//
	// - `Lingjun`: ACK Lingjun clusters, which provide the Pro edition.
	//
	// example:
	//
	// Default
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
	// The kube-proxy mode. Valid values:
	//
	// - `iptables`: A mature and stable kube-proxy mode. Service discovery and load balancing for Kubernetes Services are configured by using iptables rules. However, the performance is average and greatly affected by the scale. This mode is suitable for clusters with a small number of Services.
	//
	// - `ipvs`: A high-performance kube-proxy mode. Service discovery and load balancing for Kubernetes Services are configured by using the Linux IPVS module. This mode is suitable for clusters with a large number of Services and scenarios that require high-performance load balancing.
	//
	// - `nftables`: A next-generation kube-proxy mode. Service discovery and load balancing are implemented based on Linux nftables, which is a modern alternative to iptables. Compared with iptables, nftables provides better network performance, rule update efficiency, and performance in large-scale Service scenarios.
	//
	// This mode is supported only in clusters of Kubernetes 1.35 and later. The Kubernetes community deprecated IPVS in version 1.35. We recommend that you use nftables for new clusters to obtain long-term community support.
	//
	// Default value: `ipvs`.
	//
	// example:
	//
	// ipvs
	ProxyMode *string `json:"proxy_mode,omitempty" xml:"proxy_mode,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configurations, use the `rds_instances` parameter under `scaling_group` in `nodepool` instead. The list of ApsaraDB RDS instances to which you want to add the cluster to the whitelist. We recommend that you add the pod CIDR block and node CIDR block to the RDS whitelist. If you specify RDS instances here, nodes may fail to be created because the instances are not in the running state.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// The ID of the region where the cluster is deployed. For details, see [Regions supported by container service](https://help.aliyun.com/document_detail/216938.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// The ID of the resource group to which the cluster belongs. This parameter is used to isolate different resources.
	//
	// example:
	//
	// rg-acfm3mkrure****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// The configuration of the RAM Roles for Service Accounts (RRSA) feature.
	RrsaConfig *CreateClusterRequestRrsaConfig `json:"rrsa_config,omitempty" xml:"rrsa_config,omitempty" type:"Struct"`
	// Deprecated
	//
	// The container runtime used in the cluster. Valid values: containerd, sandboxed container, and Docker.
	//
	// > Kubernetes 1.24 and later no longer support Docker as a built-in container runtime.
	//
	// For more information, see [Comparison among Docker, containerd, and sandboxed container runtimes](https://help.aliyun.com/document_detail/160313.html).
	Runtime *Runtime `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The ID of the security group. Required when you create a cluster by using an existing security group. This parameter and is_enterprise_security_group are mutually exclusive. Cluster nodes are automatically added to this security group.
	//
	// example:
	//
	// sg-bp1bdue0qc1g7k****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane configurations, use the `security_hardening_os` parameter under `control_plane_config` instead. For node pool configurations, use the `security_hardening_os` parameter under `scaling_group` in `nodepool` instead. Specifies whether to enable Alibaba Cloud OS security hardening. Valid values:
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SecurityHardeningOs *bool `json:"security_hardening_os,omitempty" xml:"security_hardening_os,omitempty"`
	// A ServiceAccount is an access credential for communication between pods and the cluster API server. The `service-account-issuer` is the issuer identity in the `serviceaccount token`, which is the `iss` field in the `token payload`. For more information about `ServiceAccount`, see [Deploy service account token volume projection](https://help.aliyun.com/document_detail/160384.html).
	//
	// example:
	//
	// kubernetes.default.svc
	ServiceAccountIssuer *string `json:"service_account_issuer,omitempty" xml:"service_account_issuer,omitempty"`
	// The CIDR block of the Service network. Valid values: 10.0.0.0/16 to 10.0.0.0/24, 172.16.0.0/16 to 172.31.0.0/24, and 192.168.0.0/16 to 192.168.0.0/24. The CIDR block cannot overlap with the VPC CIDR block 10.1.0.0/21 or the CIDR blocks used by existing Kubernetes clusters in the VPC. The CIDR block cannot be modified after the cluster is created. Default value: 172.19.0.0/20.
	//
	// example:
	//
	// 172.21.0.0/20
	ServiceCidr *string `json:"service_cidr,omitempty" xml:"service_cidr,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The service discovery type used in the cluster. This parameter is used to specify the service discovery method in an `ACK Serverless` cluster. Valid values:
	//
	// - `CoreDNS`: Uses the standard Kubernetes service discovery component CoreDNS. You must deploy a group of containers in the cluster for DNS resolution. By default, two ECI instances with 0.25 vCPUs and 512 MiB of memory are used.
	//
	// - `PrivateZone`: Uses Alibaba Cloud PrivateZone to provide service discovery capabilities. You must activate the PrivateZone service.
	//
	// Default value: None (disabled).
	ServiceDiscoveryTypes []*string `json:"service_discovery_types,omitempty" xml:"service_discovery_types,omitempty" type:"Repeated"`
	// Specifies whether to configure SNAT for the VPC. Valid values:
	//
	// - `true`: Enables automatic creation of a NAT gateway and configures SNAT rules. Settings: `true` is required if nodes and applications in the cluster need to access the Internet.
	//
	// - `false`: Does not create a NAT gateway or configure SNAT rules. Nodes and applications in the cluster cannot access the Internet.
	//
	// > If you do not enable this feature when you create the cluster, you can [manually enable it](https://help.aliyun.com/document_detail/178480.html) later if your business needs to access the Internet.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SnatEntry *bool `json:"snat_entry,omitempty" xml:"snat_entry,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For control plane node configurations, use the `soc_enabled` parameter under `control_plane_config` instead. For node pool configurations, use the `soc_enabled` parameter under `scaling_group` in `nodepool` instead. Specifies whether to enable mlps 2.0 security hardening. For more information, see [ACK mlps 2.0 security hardening instructions](https://help.aliyun.com/document_detail/196148.html). Valid values:
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// Specifies whether to enable public SSH access. This parameter is used to log on to the master nodes of an ACK dedicated cluster. This parameter does not take effect in managed clusters. Valid values:
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	SshFlags *bool `json:"ssh_flags,omitempty" xml:"ssh_flags,omitempty"`
	// The support policy for the cluster version. The value is case-sensitive. Valid values:
	//
	// - Standard: Standard support.
	//
	// - Extended: Extended support. After the standard support period ends, the cluster enters the extended support period and continues to receive CVE patch versions.
	//
	// Only ACK Pro managed clusters (where `cluster_spec` is set to `ack.pro.*` and `profile` is set to `Default`) support the `Extended` value. Specifying `Extended` for ACK Basic, ACS, ACK Edge managed, and ACK dedicated clusters will be rejected. Patch versions in the extended support period are also applicable only to ACK Pro managed clusters. Specifying such versions for other cluster types will be rejected. If this parameter is not specified, the default value is determined by the server configuration.
	//
	// example:
	//
	// Extended
	SupportPolicy *string `json:"support_policy,omitempty" xml:"support_policy,omitempty"`
	// The node tags. Tag definition rules:
	//
	// - A tag is a case-sensitive key-value pair. You can configure up to 20 tags.
	//
	// - Tag keys must be unique and can be up to 64 characters in length. Tag values can be empty and can be up to 128 characters in length. Tag keys and values cannot start with "aliyun", "acs:", "https://", or "http://". For more information, see [Labels and Selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set).
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configurations, use the `taints` parameter under `kubernetes_config` in `nodepool` instead. The taints of the nodes. Taints and tolerations work together to prevent pods from being scheduled to inappropriate nodes. For more information, see [Taints and Tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/).
	//
	// [taint-and-toleration](https://kubernetes.io/zh/docs/concepts/scheduling-eviction/taint-and-toleration/)
	Taints []*Taint `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- By default, no rollback is performed when cluster creation fails. You must manually clean up the failed cluster. The timeout period for cluster creation, in minutes. Default value: `60`.
	//
	// example:
	//
	// 60
	TimeoutMins *int64 `json:"timeout_mins,omitempty" xml:"timeout_mins,omitempty"`
	// The time zone used by the cluster. For more information, see [Supported time zones](https://help.aliyun.com/document_detail/354879.html).
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The custom cluster CA certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----****
	UserCa *string `json:"user_ca,omitempty" xml:"user_ca,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The custom node data.
	//
	// example:
	//
	// IyEvdXNyL2Jpbi9iYXNoCmVjaG8gIkhlbGxvIEFD****
	UserData *string `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// The ID of the VPC used by the cluster. This parameter is required when you create the cluster.
	//
	// example:
	//
	// vpc-2zeik9h3ahvv2zz95****
	Vpcid *string `json:"vpcid,omitempty" xml:"vpcid,omitempty"`
	// The vSwitches where the cluster nodes are deployed. This parameter is required when you create a managed cluster with zero nodes.
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configurations, use the `auto_renew` parameter under `scaling_group` in `nodepool` instead. Specifies whether to enable auto-renewal for worker nodes. This parameter takes effect only when `worker_instance_charge_type` is set to `PrePaid`. Valid values:
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	WorkerAutoRenew *bool `json:"worker_auto_renew,omitempty" xml:"worker_auto_renew,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configurations, use the `auto_renew_period` parameter under `scaling_group` in `nodepool` instead.
	//
	// The auto-renewal epoch for worker nodes. This parameter takes effect and is required only when the subscription billing method is selected.
	//
	// Valid values: 1, 2, 3, 6, and 12.
	//
	// example:
	//
	// 1
	WorkerAutoRenewPeriod *int64 `json:"worker_auto_renew_period,omitempty" xml:"worker_auto_renew_period,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configurations, use the `data_disks` parameter under `scaling_group` in `nodepool` instead.
	//
	// The combination of configurations such as the type and size of the data cloud disks for worker nodes.
	WorkerDataDisks []*CreateClusterRequestWorkerDataDisks `json:"worker_data_disks,omitempty" xml:"worker_data_disks,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configurations, use the `instance_charge_type` parameter under `scaling_group` in `nodepool` instead.
	//
	// The billing method for worker nodes. Valid values:
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
	WorkerInstanceChargeType *string `json:"worker_instance_charge_type,omitempty" xml:"worker_instance_charge_type,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configurations, use the `instance_types` parameter under `scaling_group` in `nodepool` instead.
	//
	// The instance configurations for worker nodes.
	WorkerInstanceTypes []*string `json:"worker_instance_types,omitempty" xml:"worker_instance_types,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configurations, use the `period` parameter under `scaling_group` in `nodepool` instead.
	//
	// The subscription duration for worker nodes. This parameter takes effect and is required only when `worker_instance_charge_type` is set to `PrePaid`.
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
	// **[Deprecated]*	- For node pool configurations, use the `period_unit` parameter under `scaling_group` in `nodepool` instead.
	//
	// The subscription period unit for worker nodes. This parameter is required when the billing method is set to `PrePaid`.
	//
	// Valid value: `Month`. Currently, only month is supported.
	//
	// example:
	//
	// Month
	WorkerPeriodUnit *string `json:"worker_period_unit,omitempty" xml:"worker_period_unit,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configurations, use the `system_disk_category` parameter under `scaling_group` in `nodepool` instead.
	//
	// The type of the system cloud disk for worker nodes. For more information, see [Block Storage overview](https://help.aliyun.com/document_detail/63136.html).
	//
	// Valid values:
	//
	// - `cloud_efficiency`: Ultra cloud disk.
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
	// **[Deprecated]*	- For node pool configurations, use the `system_disk_performance_level` parameter under `scaling_group` in `nodepool` instead.
	//
	// When the system cloud disk is an ESSD, you can configure the performance level (PL). For more information, see [ESSD](https://help.aliyun.com/document_detail/122389.html).
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
	// **[Deprecated]*	- For node pool configurations, use the `system_disk_size` parameter under `scaling_group` in `nodepool` instead.
	//
	// The size of the system cloud disk for worker nodes. Unit: GiB.
	//
	// Valid values: 40 to 500.
	//
	// The value of this parameter must be greater than or equal to the maximum of 40 and the image size.
	//
	// Default value: `120`.
	//
	// example:
	//
	// 120
	WorkerSystemDiskSize *int64 `json:"worker_system_disk_size,omitempty" xml:"worker_system_disk_size,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configurations, use the `system_disk_snapshot_policy_id` parameter under `scaling_group` in `nodepool` instead.
	//
	// The ID of the automatic snapshot policy applied to the system cloud disk of worker nodes.
	//
	// example:
	//
	// sp-2zej1nogjvovnz4z****
	WorkerSystemDiskSnapshotPolicyId *string `json:"worker_system_disk_snapshot_policy_id,omitempty" xml:"worker_system_disk_snapshot_policy_id,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- For node pool configurations, use the `vswitch_ids` parameter under `scaling_group` in `nodepool` instead.
	//
	// The list of vSwitches used by cluster nodes. Each node corresponds to one vSwitch.
	//
	// When you create a managed cluster with zero nodes, the `worker_vswitch_ids` parameter is not required, but you must specify `vswitch_ids`.
	WorkerVswitchIds []*string `json:"worker_vswitch_ids,omitempty" xml:"worker_vswitch_ids,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- Use the `zone_ids` parameter instead.
	//
	// The zone ID of the region where the cluster resides. This parameter is specific to managed ACK clusters.
	//
	// When you create a managed ACK cluster, if `vpc_id` and `vswitch_ids` are not specified, you must specify `zone_id` for the cluster to automatically create a VPC and network resources in this zone. If `vpc_id` and `vswitch_ids` are specified, this parameter does not take effect.
	//
	// example:
	//
	// cn-beiji****
	ZoneId *string `json:"zone_id,omitempty" xml:"zone_id,omitempty"`
	// The zone IDs of the region where the cluster resides. This parameter is specific to managed ACK clusters.
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

func (s *CreateClusterRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *CreateClusterRequest) GetSupportPolicy() *string {
	return s.SupportPolicy
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

func (s *CreateClusterRequest) SetClientToken(v string) *CreateClusterRequest {
	s.ClientToken = &v
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

func (s *CreateClusterRequest) SetSupportPolicy(v string) *CreateClusterRequest {
	s.SupportPolicy = &v
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
	// Specifies whether to enable the cluster audit log feature. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The [SLS Project](https://help.aliyun.com/document_detail/48873.html) where the [Logstore](https://help.aliyun.com/document_detail/48873.html) for cluster audit logs is located.
	//
	// - Default value: `k8s-log-{clusterid}`.
	//
	// - After you enable the cluster audit log feature, a Logstore for cluster audit logs is created in the specified SLS Project.
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
	// Specifies whether to enable the intelligent managed mode. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
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
	// Specifies whether to enable auto-renewal for control plane nodes. This parameter takes effect only when the billing method is set to PrePaid. Valid values:
	//
	// - true: enables auto-renewal.
	//
	// - false: disables auto-renewal.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// The auto-renewal period for control plane nodes. Valid values: 1, 2, 3, 6, and 12. Unit: months.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// The billing method for control plane nodes. Valid values:
	//
	// - PrePaid: subscription.
	//
	// - PostPaid: pay-as-you-go.
	//
	// Default value: PostPaid.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"charge_type,omitempty" xml:"charge_type,omitempty"`
	// Specifies whether to install the CloudMonitor agent on the nodes. Valid values:
	//
	// - true: installs the CloudMonitor agent.
	//
	// - false: does not install the CloudMonitor agent.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// The CPU management policy for nodes. Valid values:
	//
	// - static: allows you to enhance CPU affinity and exclusivity for pods with specific resource characteristics on the node.
	//
	// - none: uses the default CPU affinity policy.
	//
	// Default value: none.
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
	// The ID of the image.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20240819.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The type of the operating system image.
	//
	// example:
	//
	// AliyunLinux3
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// The instance metadata access configuration for ECS instances. This parameter is available only to users in the whitelist. To use this parameter, submit a ticket.
	InstanceMetadataOptions *InstanceMetadataOptions `json:"instance_metadata_options,omitempty" xml:"instance_metadata_options,omitempty"`
	// The instance types of the nodes.
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// The name of the key pair. You must specify this parameter or the login_password parameter.
	//
	// example:
	//
	// ack
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// The SSH logon password. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. You must specify this parameter or the key_pair parameter.
	//
	// example:
	//
	// ********
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// **[Deprecated]*	- The service port range of the node.
	//
	// example:
	//
	// 30000-32767
	NodePortRange *string `json:"node_port_range,omitempty" xml:"node_port_range,omitempty"`
	// The subscription period for control plane nodes. This parameter is required and takes effect only when the billing method is set to PrePaid. Valid values: 1, 2, 3, 6, 12, 24, 36, 48, and 60. Unit: months.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The unit of the subscription period for control plane nodes. This parameter is required and takes effect only when the billing method is set to PrePaid. Valid value: Month. Currently, only months are supported.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// **[Deprecated]*	- The name of the runtime for control plane nodes. Valid value:
	//
	// containerd: the containerd runtime, which supports all cluster versions.
	//
	// example:
	//
	// containerd
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// Specifies whether to enable security hardening for Alibaba Cloud OS. Valid values:
	//
	// - true: enables security hardening for Alibaba Cloud OS.
	//
	// - false: disables security hardening for Alibaba Cloud OS.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	SecurityHardeningOs *bool `json:"security_hardening_os,omitempty" xml:"security_hardening_os,omitempty"`
	// The number of control plane nodes. Valid values: 3 and 5.
	//
	// example:
	//
	// 3
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// Specifies whether to enable MLPS 2.0 security hardening. Valid values:
	//
	// - true: enables MLPS 2.0 security hardening.
	//
	// - false: disables MLPS 2.0 security hardening.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// Specifies whether to enable the performance burst feature for the system cloud disk of the node. Valid values:
	//
	// - true: enables the feature.
	//
	// - false: disables the feature.
	//
	// This parameter is supported only when system_disk_category is set to cloud_auto.
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// The type of the system cloud disk for the node. Valid values:
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: ESSD.
	//
	// - cloud_auto: ESSD AutoPL cloud disk.
	//
	// - cloud_essd_entry: ESSD Entry cloud disk.
	//
	// Default value: cloud_ssd. The default value may vary by zone.
	//
	// example:
	//
	// cloud_ssd
	SystemDiskCategory *string `json:"system_disk_category,omitempty" xml:"system_disk_category,omitempty"`
	// The performance level of the system cloud disk for the node. This parameter takes effect only for ESSDs. The performance level varies based on the disk size. For more information, see [ESSD (standard SSD)](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// The provisioned read and write IOPS for the system cloud disk of the node. Valid values: 0 to min{50,000, 1000 × capacity - baseline performance}. Baseline performance = min{1,800 + 50 × capacity, 50,000}.
	//
	// This parameter is supported only when system_disk_category is set to cloud_auto.
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// The size of the system cloud disk for the node. Valid values: 40 to 500. Unit: GiB.
	//
	// Default value: 120.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// The ID of the automatic snapshot policy for the node.
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
	// The internal DNS configuration of the cluster, which applies to ACK managed clusters. The internal domain name is used by node-side system components such as kubelet and kube-proxy to access the API server. If internal domain name access is not enabled, node-side system components access the API server by using the IP address of the CLB instance.
	InternalDnsConfig *CreateClusterRequestControlPlaneEndpointsConfigInternalDnsConfig `json:"internal_dns_config,omitempty" xml:"internal_dns_config,omitempty" type:"Struct"`
	// The connection configuration of the cluster. If you specify this parameter, the endpoint_public_access and load_balancer_id parameters do not take effect. ACK supports only automatic creation of NLB instances. To specify a CLB or NLB instance, use load_balancers_config to specify the corresponding instance ID.
	LoadBalancersConfig []*CreateClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig `json:"load_balancers_config,omitempty" xml:"load_balancers_config,omitempty" type:"Repeated"`
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

func (s *CreateClusterRequestControlPlaneEndpointsConfig) GetLoadBalancersConfig() []*CreateClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig {
	return s.LoadBalancersConfig
}

func (s *CreateClusterRequestControlPlaneEndpointsConfig) SetInternalDnsConfig(v *CreateClusterRequestControlPlaneEndpointsConfigInternalDnsConfig) *CreateClusterRequestControlPlaneEndpointsConfig {
	s.InternalDnsConfig = v
	return s
}

func (s *CreateClusterRequestControlPlaneEndpointsConfig) SetLoadBalancersConfig(v []*CreateClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig) *CreateClusterRequestControlPlaneEndpointsConfig {
	s.LoadBalancersConfig = v
	return s
}

func (s *CreateClusterRequestControlPlaneEndpointsConfig) Validate() error {
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

type CreateClusterRequestControlPlaneEndpointsConfigInternalDnsConfig struct {
	// The VPCs where the internal DNS records of the cluster take effect.
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

type CreateClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig struct {
	// The type of the endpoint.
	//
	// example:
	//
	// private
	EndpointType *string `json:"endpoint_type,omitempty" xml:"endpoint_type,omitempty"`
	// The ID of the NLB instance.
	//
	// example:
	//
	// nlb-0ogk9aaxxxxxxx
	LoadBalancerId *string `json:"load_balancer_id,omitempty" xml:"load_balancer_id,omitempty"`
}

func (s CreateClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig) GetEndpointType() *string {
	return s.EndpointType
}

func (s *CreateClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig) SetEndpointType(v string) *CreateClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig {
	s.EndpointType = &v
	return s
}

func (s *CreateClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig) SetLoadBalancerId(v string) *CreateClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateClusterRequestControlPlaneEndpointsConfigLoadBalancersConfig) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestOperationPolicy struct {
	// The automatic cluster upgrade configuration.
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
	// The automatic cluster upgrade channel. Valid values:
	//
	// - patch: Automatically upgrades to the patch version when a patch version is available for the current minor version. The new Kubernetes version does not include breaking changes.
	//
	// - stable: Automatically upgrades to the latest patch version of the second most recent minor version. The new Kubernetes version may include API and feature changes, but its stability has been widely verified.
	//
	// - rapid: Automatically upgrades to the latest patch version of the latest minor version to obtain new features from the Kubernetes community faster.
	//
	// example:
	//
	// stable
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// Specifies whether to enable automatic cluster upgrade. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
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
	// Specifies whether to enable the RRSA feature. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
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
	// The type of the data cloud disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// Specifies whether to encrypt the data cloud disk. Valid values:
	//
	// - `true`: Encrypts the data cloud disk.
	//
	// - `false`: Does not encrypt the data cloud disk.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	Encrypted *string `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// The performance level of the data cloud disk. This parameter takes effect only for [ESSD](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"performance_level,omitempty" xml:"performance_level,omitempty"`
	// The size of the data cloud disk. Valid values: 40 to 32767. Unit: GiB.
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
