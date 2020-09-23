// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AttachInstancesRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 待添加的实例列表。
	Instances []*string `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// 容器运行时。
	Runtime *AttachInstancesRequestRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
	// 自定义镜像ID。
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// 是否格式化数据盘。
	FormatDisk *bool `json:"format_disk,omitempty" xml:"format_disk,omitempty"`
	// 是否保留实例名称。
	KeepInstanceName *bool `json:"keep_instance_name,omitempty" xml:"keep_instance_name,omitempty"`
	// CPU策略。
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// key_pair名称，与login_password二选一
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// password，与key_pair二选一。
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// 是否为边缘节点。
	IsEdgeWorker *bool `json:"is_edge_worker,omitempty" xml:"is_edge_worker,omitempty"`
	// 用户自定义数据。
	UserData *string `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// 节点池ID，欲将节点添加到哪个节点池中。。
	NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	// RDS实例列表。
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// 节点标签。
	Tags []*AttachInstancesRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s AttachInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesRequest) GoString() string {
	return s.String()
}

func (s *AttachInstancesRequest) SetClusterId(v string) *AttachInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *AttachInstancesRequest) SetInstances(v []*string) *AttachInstancesRequest {
	s.Instances = v
	return s
}

func (s *AttachInstancesRequest) SetRuntime(v *AttachInstancesRequestRuntime) *AttachInstancesRequest {
	s.Runtime = v
	return s
}

func (s *AttachInstancesRequest) SetImageId(v string) *AttachInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *AttachInstancesRequest) SetFormatDisk(v bool) *AttachInstancesRequest {
	s.FormatDisk = &v
	return s
}

func (s *AttachInstancesRequest) SetKeepInstanceName(v bool) *AttachInstancesRequest {
	s.KeepInstanceName = &v
	return s
}

func (s *AttachInstancesRequest) SetCpuPolicy(v string) *AttachInstancesRequest {
	s.CpuPolicy = &v
	return s
}

func (s *AttachInstancesRequest) SetKeyPair(v string) *AttachInstancesRequest {
	s.KeyPair = &v
	return s
}

func (s *AttachInstancesRequest) SetPassword(v string) *AttachInstancesRequest {
	s.Password = &v
	return s
}

func (s *AttachInstancesRequest) SetIsEdgeWorker(v bool) *AttachInstancesRequest {
	s.IsEdgeWorker = &v
	return s
}

func (s *AttachInstancesRequest) SetUserData(v string) *AttachInstancesRequest {
	s.UserData = &v
	return s
}

func (s *AttachInstancesRequest) SetNodepoolId(v string) *AttachInstancesRequest {
	s.NodepoolId = &v
	return s
}

func (s *AttachInstancesRequest) SetRdsInstances(v []*string) *AttachInstancesRequest {
	s.RdsInstances = v
	return s
}

func (s *AttachInstancesRequest) SetTags(v []*AttachInstancesRequestTags) *AttachInstancesRequest {
	s.Tags = v
	return s
}

type AttachInstancesRequestRuntime struct {
	// 容器运行时名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 容器运行时版本。
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s AttachInstancesRequestRuntime) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesRequestRuntime) GoString() string {
	return s.String()
}

func (s *AttachInstancesRequestRuntime) SetName(v string) *AttachInstancesRequestRuntime {
	s.Name = &v
	return s
}

func (s *AttachInstancesRequestRuntime) SetVersion(v string) *AttachInstancesRequestRuntime {
	s.Version = &v
	return s
}

type AttachInstancesRequestTags struct {
	// 标签名。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 标签值。
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AttachInstancesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesRequestTags) GoString() string {
	return s.String()
}

func (s *AttachInstancesRequestTags) SetKey(v string) *AttachInstancesRequestTags {
	s.Key = &v
	return s
}

func (s *AttachInstancesRequestTags) SetValue(v string) *AttachInstancesRequestTags {
	s.Value = &v
	return s
}

type AttachInstancesResponseBody struct {
	// 节点信息列表。
	List []*AttachInstancesResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// 任务ID。
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s AttachInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponseBody) SetList(v []*AttachInstancesResponseBodyList) *AttachInstancesResponseBody {
	s.List = v
	return s
}

func (s *AttachInstancesResponseBody) SetTaskId(v string) *AttachInstancesResponseBody {
	s.TaskId = &v
	return s
}

type AttachInstancesResponseBodyList struct {
	// 状态码。
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 实例ID。
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 添加结果描述。
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s AttachInstancesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesResponseBodyList) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponseBodyList) SetCode(v string) *AttachInstancesResponseBodyList {
	s.Code = &v
	return s
}

func (s *AttachInstancesResponseBodyList) SetInstanceId(v string) *AttachInstancesResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *AttachInstancesResponseBodyList) SetMessage(v string) *AttachInstancesResponseBodyList {
	s.Message = &v
	return s
}

type AttachInstancesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesResponse) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponse) SetHeaders(v map[string]*string) *AttachInstancesResponse {
	s.Headers = v
	return s
}

func (s *AttachInstancesResponse) SetBody(v *AttachInstancesResponseBody) *AttachInstancesResponse {
	s.Body = v
	return s
}

type CancelClusterUpgradeRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s CancelClusterUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelClusterUpgradeRequest) GoString() string {
	return s.String()
}

func (s *CancelClusterUpgradeRequest) SetClusterId(v string) *CancelClusterUpgradeRequest {
	s.ClusterId = &v
	return s
}

type CancelClusterUpgradeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s CancelClusterUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelClusterUpgradeResponse) GoString() string {
	return s.String()
}

func (s *CancelClusterUpgradeResponse) SetHeaders(v map[string]*string) *CancelClusterUpgradeResponse {
	s.Headers = v
	return s
}

type CancelComponentUpgradeRequest struct {
	// 集群ID。
	Clusterid *string `json:"clusterid,omitempty" xml:"clusterid,omitempty"`
	// 组件ID。
	Componentid *string `json:"componentid,omitempty" xml:"componentid,omitempty"`
}

func (s CancelComponentUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelComponentUpgradeRequest) GoString() string {
	return s.String()
}

func (s *CancelComponentUpgradeRequest) SetClusterid(v string) *CancelComponentUpgradeRequest {
	s.Clusterid = &v
	return s
}

func (s *CancelComponentUpgradeRequest) SetComponentid(v string) *CancelComponentUpgradeRequest {
	s.Componentid = &v
	return s
}

type CancelComponentUpgradeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s CancelComponentUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelComponentUpgradeResponse) GoString() string {
	return s.String()
}

func (s *CancelComponentUpgradeResponse) SetHeaders(v map[string]*string) *CancelComponentUpgradeResponse {
	s.Headers = v
	return s
}

type CreateClusterRequest struct {
	// 集群名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 集群类型
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// 集群所属地域ID。
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// 集群所属地域内的可用区ID。
	ZoneId *string `json:"zone_id,omitempty" xml:"zone_id,omitempty"`
	// 集群版本好。
	KubernetesVersion *string `json:"kubernetes_version,omitempty" xml:"kubernetes_version,omitempty"`
	// 集群是否开启删除保护。
	DeletionProtection *string `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// 容器运行时。
	Runtime *CreateClusterRequestRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
	// 集群使用的VPC。
	Vpcid *string `json:"vpcid,omitempty" xml:"vpcid,omitempty"`
	// 集群使用的虚拟交换机。
	WorkerVswitchIds []*string `json:"worker_vswitch_ids,omitempty" xml:"worker_vswitch_ids,omitempty" type:"Repeated"`
	// POD网络地址段。
	ContainerCidr *string `json:"container_cidr,omitempty" xml:"container_cidr,omitempty"`
	// Service网络地址段。
	ServiceCidr *string `json:"service_cidr,omitempty" xml:"service_cidr,omitempty"`
	// 节点IP数量，这里通过CIDR来指定。
	NodeCidrMask *string `json:"node_cidr_mask,omitempty" xml:"node_cidr_mask,omitempty"`
	// 集群是否配置SNAT。
	SnatEntry *bool `json:"snat_entry,omitempty" xml:"snat_entry,omitempty"`
	// 集群是否运行公网访问。
	EndpointPublicAccess *bool `json:"endpoint_public_access,omitempty" xml:"endpoint_public_access,omitempty"`
	// 集群是否开启公网SSH登录。
	SshFlags *bool `json:"ssh_flags,omitempty" xml:"ssh_flags,omitempty"`
	// RDS列表，将该ECS加入到选择的RDS实例的白名单中。。
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// 自定义安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// 是否自动创建企业安全组，与security_group_id二选一。
	IsEnterpriseSecurityGroup *bool `json:"is_enterprise_security_group,omitempty" xml:"is_enterprise_security_group,omitempty"`
	// kube-proxy代理模式。
	ProxyMode *string `json:"proxy_mode,omitempty" xml:"proxy_mode,omitempty"`
	// 集群标签。
	Tags []*CreateClusterRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// 自定义镜像ID。
	ImagesId *string `json:"images_id,omitempty" xml:"images_id,omitempty"`
	// Master节点付费类型。
	MasterInstanceChargeType *string `json:"master_instance_charge_type,omitempty" xml:"master_instance_charge_type,omitempty"`
	// Master节点包年包月时长，当master_instance_charge_type取值为PrePaid时才生效且为必选值。
	MasterPeriod *int64 `json:"master_period,omitempty" xml:"master_period,omitempty"`
	// Master节点包年包月周期。
	MasterPeriodUnit *string `json:"master_period_unit,omitempty" xml:"master_period_unit,omitempty"`
	// Master节点是否自动续费。
	MasterAutoRenew *bool `json:"master_auto_renew,omitempty" xml:"master_auto_renew,omitempty"`
	// Master节点自动续费周期。
	MasterAutoRenewPeriod *int64 `json:"master_auto_renew_period,omitempty" xml:"master_auto_renew_period,omitempty"`
	// Master节点数量。
	MasterCount *int64 `json:"master_count,omitempty" xml:"master_count,omitempty"`
	// Master节点交换机ID列表。
	MasterVswitchIds []*string `json:"master_vswitch_ids,omitempty" xml:"master_vswitch_ids,omitempty" type:"Repeated"`
	// Master节点ECS规格类型。
	MasterInstanceTypes []*string `json:"master_instance_types,omitempty" xml:"master_instance_types,omitempty" type:"Repeated"`
	// Master节点系统盘类型。
	MasterSystemDiskCategory *string `json:"master_system_disk_category,omitempty" xml:"master_system_disk_category,omitempty"`
	// Master节点系统盘大小。
	MasterSystemDiskSize *int64 `json:"master_system_disk_size,omitempty" xml:"master_system_disk_size,omitempty"`
	// Worker节点付费类型。
	WorkerInstanceChargeType *string `json:"worker_instance_charge_type,omitempty" xml:"worker_instance_charge_type,omitempty"`
	// Worker节点包年包月时长。
	WorkerPeriod *int64 `json:"worker_period,omitempty" xml:"worker_period,omitempty"`
	// Worker节点包年包月周期。
	WorkerPeriodUnit *string `json:"worker_period_unit,omitempty" xml:"worker_period_unit,omitempty"`
	// Worker节点是否自动续费。
	WorkerAutoRenew *bool `json:"worker_auto_renew,omitempty" xml:"worker_auto_renew,omitempty"`
	// Worker节点自动续费周期。
	WorkerAutoRenewPeriod *int64 `json:"worker_auto_renew_period,omitempty" xml:"worker_auto_renew_period,omitempty"`
	// Worker节点数量。
	NumOfNodes *int64 `json:"num_of_nodes,omitempty" xml:"num_of_nodes,omitempty"`
	// Worker节点ECS实例类型。
	WorkerInstanceTypes []*string `json:"worker_instance_types,omitempty" xml:"worker_instance_types,omitempty" type:"Repeated"`
	// Worker节点系统盘类型。
	WorkerSystemDiskCategory *string `json:"worker_system_disk_category,omitempty" xml:"worker_system_disk_category,omitempty"`
	// Worker节点系统盘大小。
	WorkerSystemDiskSize *int64 `json:"worker_system_disk_size,omitempty" xml:"worker_system_disk_size,omitempty"`
	// Worker节点数据盘配置。
	WorkerDataDisks []*CreateClusterRequestWorkerDataDisks `json:"worker_data_disks,omitempty" xml:"worker_data_disks,omitempty" type:"Repeated"`
	// 操作系统。
	OsType *string `json:"os_type,omitempty" xml:"os_type,omitempty"`
	// key_pair名称，和login_password二选一。
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// SSH登录密码，与key_pair二选一。
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// 节点用户自定义数据。
	UserData *string `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// 节点服务端口范围。
	NodePortRange *string `json:"node_port_range,omitempty" xml:"node_port_range,omitempty"`
	// CPU管理策略。
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// 污点信息。
	Taints []*CreateClusterRequestTaints `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// 是否安装云监控插件。
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// 组件信息。
	Addons []*CreateClusterRequestAddons `json:"addons,omitempty" xml:"addons,omitempty" type:"Repeated"`
	// 操作系统发行版。
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// 虚拟交换机列表。List长度范围为[1，3]。当集群类型为托管版或标准serverless集群时，该参数必填。
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
	// 是否开启PrivateZone用于服务发现。
	PrivateZone *bool `json:"private_zone,omitempty" xml:"private_zone,omitempty"`
	// 边缘集群标识。
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
	// Pod的虚拟交换机列表，在ENI多网卡模式下，需要传额外的VSwitch ID给addon。
	PodVswitchIds []*string `json:"pod_vswitch_ids,omitempty" xml:"pod_vswitch_ids,omitempty" type:"Repeated"`
	// 集群创建失败后是否回滚。
	DisableRollback *bool `json:"disable_rollback,omitempty" xml:"disable_rollback,omitempty"`
	// 集群创建超时时间。
	TimeoutMins *int64 `json:"timeout_mins,omitempty" xml:"timeout_mins,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetName(v string) *CreateClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateClusterRequest) SetClusterType(v string) *CreateClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterRequest) SetRegionId(v string) *CreateClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateClusterRequest) SetZoneId(v string) *CreateClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateClusterRequest) SetKubernetesVersion(v string) *CreateClusterRequest {
	s.KubernetesVersion = &v
	return s
}

func (s *CreateClusterRequest) SetDeletionProtection(v string) *CreateClusterRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateClusterRequest) SetRuntime(v *CreateClusterRequestRuntime) *CreateClusterRequest {
	s.Runtime = v
	return s
}

func (s *CreateClusterRequest) SetVpcid(v string) *CreateClusterRequest {
	s.Vpcid = &v
	return s
}

func (s *CreateClusterRequest) SetWorkerVswitchIds(v []*string) *CreateClusterRequest {
	s.WorkerVswitchIds = v
	return s
}

func (s *CreateClusterRequest) SetContainerCidr(v string) *CreateClusterRequest {
	s.ContainerCidr = &v
	return s
}

func (s *CreateClusterRequest) SetServiceCidr(v string) *CreateClusterRequest {
	s.ServiceCidr = &v
	return s
}

func (s *CreateClusterRequest) SetNodeCidrMask(v string) *CreateClusterRequest {
	s.NodeCidrMask = &v
	return s
}

func (s *CreateClusterRequest) SetSnatEntry(v bool) *CreateClusterRequest {
	s.SnatEntry = &v
	return s
}

func (s *CreateClusterRequest) SetEndpointPublicAccess(v bool) *CreateClusterRequest {
	s.EndpointPublicAccess = &v
	return s
}

func (s *CreateClusterRequest) SetSshFlags(v bool) *CreateClusterRequest {
	s.SshFlags = &v
	return s
}

func (s *CreateClusterRequest) SetRdsInstances(v []*string) *CreateClusterRequest {
	s.RdsInstances = v
	return s
}

func (s *CreateClusterRequest) SetSecurityGroupId(v string) *CreateClusterRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetIsEnterpriseSecurityGroup(v bool) *CreateClusterRequest {
	s.IsEnterpriseSecurityGroup = &v
	return s
}

func (s *CreateClusterRequest) SetProxyMode(v string) *CreateClusterRequest {
	s.ProxyMode = &v
	return s
}

func (s *CreateClusterRequest) SetTags(v []*CreateClusterRequestTags) *CreateClusterRequest {
	s.Tags = v
	return s
}

func (s *CreateClusterRequest) SetImagesId(v string) *CreateClusterRequest {
	s.ImagesId = &v
	return s
}

func (s *CreateClusterRequest) SetMasterInstanceChargeType(v string) *CreateClusterRequest {
	s.MasterInstanceChargeType = &v
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

func (s *CreateClusterRequest) SetMasterVswitchIds(v []*string) *CreateClusterRequest {
	s.MasterVswitchIds = v
	return s
}

func (s *CreateClusterRequest) SetMasterInstanceTypes(v []*string) *CreateClusterRequest {
	s.MasterInstanceTypes = v
	return s
}

func (s *CreateClusterRequest) SetMasterSystemDiskCategory(v string) *CreateClusterRequest {
	s.MasterSystemDiskCategory = &v
	return s
}

func (s *CreateClusterRequest) SetMasterSystemDiskSize(v int64) *CreateClusterRequest {
	s.MasterSystemDiskSize = &v
	return s
}

func (s *CreateClusterRequest) SetWorkerInstanceChargeType(v string) *CreateClusterRequest {
	s.WorkerInstanceChargeType = &v
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

func (s *CreateClusterRequest) SetWorkerAutoRenew(v bool) *CreateClusterRequest {
	s.WorkerAutoRenew = &v
	return s
}

func (s *CreateClusterRequest) SetWorkerAutoRenewPeriod(v int64) *CreateClusterRequest {
	s.WorkerAutoRenewPeriod = &v
	return s
}

func (s *CreateClusterRequest) SetNumOfNodes(v int64) *CreateClusterRequest {
	s.NumOfNodes = &v
	return s
}

func (s *CreateClusterRequest) SetWorkerInstanceTypes(v []*string) *CreateClusterRequest {
	s.WorkerInstanceTypes = v
	return s
}

func (s *CreateClusterRequest) SetWorkerSystemDiskCategory(v string) *CreateClusterRequest {
	s.WorkerSystemDiskCategory = &v
	return s
}

func (s *CreateClusterRequest) SetWorkerSystemDiskSize(v int64) *CreateClusterRequest {
	s.WorkerSystemDiskSize = &v
	return s
}

func (s *CreateClusterRequest) SetWorkerDataDisks(v []*CreateClusterRequestWorkerDataDisks) *CreateClusterRequest {
	s.WorkerDataDisks = v
	return s
}

func (s *CreateClusterRequest) SetOsType(v string) *CreateClusterRequest {
	s.OsType = &v
	return s
}

func (s *CreateClusterRequest) SetKeyPair(v string) *CreateClusterRequest {
	s.KeyPair = &v
	return s
}

func (s *CreateClusterRequest) SetLoginPassword(v string) *CreateClusterRequest {
	s.LoginPassword = &v
	return s
}

func (s *CreateClusterRequest) SetUserData(v string) *CreateClusterRequest {
	s.UserData = &v
	return s
}

func (s *CreateClusterRequest) SetNodePortRange(v string) *CreateClusterRequest {
	s.NodePortRange = &v
	return s
}

func (s *CreateClusterRequest) SetCpuPolicy(v string) *CreateClusterRequest {
	s.CpuPolicy = &v
	return s
}

func (s *CreateClusterRequest) SetTaints(v []*CreateClusterRequestTaints) *CreateClusterRequest {
	s.Taints = v
	return s
}

func (s *CreateClusterRequest) SetCloudMonitorFlags(v bool) *CreateClusterRequest {
	s.CloudMonitorFlags = &v
	return s
}

func (s *CreateClusterRequest) SetAddons(v []*CreateClusterRequestAddons) *CreateClusterRequest {
	s.Addons = v
	return s
}

func (s *CreateClusterRequest) SetPlatform(v string) *CreateClusterRequest {
	s.Platform = &v
	return s
}

func (s *CreateClusterRequest) SetVswitchIds(v []*string) *CreateClusterRequest {
	s.VswitchIds = v
	return s
}

func (s *CreateClusterRequest) SetPrivateZone(v bool) *CreateClusterRequest {
	s.PrivateZone = &v
	return s
}

func (s *CreateClusterRequest) SetProfile(v string) *CreateClusterRequest {
	s.Profile = &v
	return s
}

func (s *CreateClusterRequest) SetPodVswitchIds(v []*string) *CreateClusterRequest {
	s.PodVswitchIds = v
	return s
}

func (s *CreateClusterRequest) SetDisableRollback(v bool) *CreateClusterRequest {
	s.DisableRollback = &v
	return s
}

func (s *CreateClusterRequest) SetTimeoutMins(v int64) *CreateClusterRequest {
	s.TimeoutMins = &v
	return s
}

type CreateClusterRequestRuntime struct {
	// 容器运行时名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 容器运行时版本。
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateClusterRequestRuntime) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestRuntime) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestRuntime) SetName(v string) *CreateClusterRequestRuntime {
	s.Name = &v
	return s
}

func (s *CreateClusterRequestRuntime) SetVersion(v string) *CreateClusterRequestRuntime {
	s.Version = &v
	return s
}

type CreateClusterRequestTags struct {
	// 标签key。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 标签值。
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateClusterRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestTags) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestTags) SetKey(v string) *CreateClusterRequestTags {
	s.Key = &v
	return s
}

func (s *CreateClusterRequestTags) SetValue(v string) *CreateClusterRequestTags {
	s.Value = &v
	return s
}

type CreateClusterRequestWorkerDataDisks struct {
	// 数据盘是否开启云盘备份。
	AutoSnapshotPolicyId *string `json:"auto_snapshot_policy_id,omitempty" xml:"auto_snapshot_policy_id,omitempty"`
	// 数据盘类型。
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 数据盘是否加密。
	Encrypted *string `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// 数据盘大小。
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
}

func (s CreateClusterRequestWorkerDataDisks) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestWorkerDataDisks) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestWorkerDataDisks) SetAutoSnapshotPolicyId(v string) *CreateClusterRequestWorkerDataDisks {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateClusterRequestWorkerDataDisks) SetCategory(v string) *CreateClusterRequestWorkerDataDisks {
	s.Category = &v
	return s
}

func (s *CreateClusterRequestWorkerDataDisks) SetEncrypted(v string) *CreateClusterRequestWorkerDataDisks {
	s.Encrypted = &v
	return s
}

func (s *CreateClusterRequestWorkerDataDisks) SetSize(v string) *CreateClusterRequestWorkerDataDisks {
	s.Size = &v
	return s
}

type CreateClusterRequestTaints struct {
	// 调度策略。
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	// 污点key。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 污点值。
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateClusterRequestTaints) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestTaints) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestTaints) SetEffect(v string) *CreateClusterRequestTaints {
	s.Effect = &v
	return s
}

func (s *CreateClusterRequestTaints) SetKey(v string) *CreateClusterRequestTaints {
	s.Key = &v
	return s
}

func (s *CreateClusterRequestTaints) SetValue(v string) *CreateClusterRequestTaints {
	s.Value = &v
	return s
}

type CreateClusterRequestAddons struct {
	// 组件需要的配置。
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// 组件名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateClusterRequestAddons) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestAddons) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestAddons) SetConfig(v string) *CreateClusterRequestAddons {
	s.Config = &v
	return s
}

func (s *CreateClusterRequestAddons) SetName(v string) *CreateClusterRequestAddons {
	s.Name = &v
	return s
}

type CreateClusterResponseBody struct {
	// 集群ID。
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// 请求ID。
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 任务ID。
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetTaskId(v string) *CreateClusterResponseBody {
	s.TaskId = &v
	return s
}

type CreateClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateKubernetesTriggerRequest struct {
	// 地域ID。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 项目名称。
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 触发器类型。
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateKubernetesTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKubernetesTriggerRequest) GoString() string {
	return s.String()
}

func (s *CreateKubernetesTriggerRequest) SetRegionId(v string) *CreateKubernetesTriggerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateKubernetesTriggerRequest) SetClusterId(v string) *CreateKubernetesTriggerRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateKubernetesTriggerRequest) SetProjectId(v string) *CreateKubernetesTriggerRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateKubernetesTriggerRequest) SetType(v string) *CreateKubernetesTriggerRequest {
	s.Type = &v
	return s
}

type CreateKubernetesTriggerResponseBody struct {
	// 触发器行为。
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 集群ID。
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// 触发器ID。
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 项目名称。
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty"`
}

func (s CreateKubernetesTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKubernetesTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKubernetesTriggerResponseBody) SetAction(v string) *CreateKubernetesTriggerResponseBody {
	s.Action = &v
	return s
}

func (s *CreateKubernetesTriggerResponseBody) SetClusterId(v string) *CreateKubernetesTriggerResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateKubernetesTriggerResponseBody) SetId(v string) *CreateKubernetesTriggerResponseBody {
	s.Id = &v
	return s
}

func (s *CreateKubernetesTriggerResponseBody) SetProjectId(v string) *CreateKubernetesTriggerResponseBody {
	s.ProjectId = &v
	return s
}

type CreateKubernetesTriggerResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateKubernetesTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateKubernetesTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateKubernetesTriggerResponse) GoString() string {
	return s.String()
}

func (s *CreateKubernetesTriggerResponse) SetHeaders(v map[string]*string) *CreateKubernetesTriggerResponse {
	s.Headers = v
	return s
}

func (s *CreateKubernetesTriggerResponse) SetBody(v *CreateKubernetesTriggerResponseBody) *CreateKubernetesTriggerResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 要保留的资源列表。
	RetainResources []*string `json:"retain_resources,omitempty" xml:"retain_resources,omitempty" type:"Repeated"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterRequest) SetRetainResources(v []*string) *DeleteClusterRequest {
	s.RetainResources = v
	return s
}

type DeleteClusterResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

type DeleteKubernetesTriggerRequest struct {
	// 触发器ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteKubernetesTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKubernetesTriggerRequest) GoString() string {
	return s.String()
}

func (s *DeleteKubernetesTriggerRequest) SetId(v string) *DeleteKubernetesTriggerRequest {
	s.Id = &v
	return s
}

type DeleteKubernetesTriggerResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteKubernetesTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKubernetesTriggerResponse) GoString() string {
	return s.String()
}

func (s *DeleteKubernetesTriggerResponse) SetHeaders(v map[string]*string) *DeleteKubernetesTriggerResponse {
	s.Headers = v
	return s
}

type DescribeAddonsRequest struct {
	// Region ID。
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// 集群类型，默认为kubernetes。
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
}

func (s DescribeAddonsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAddonsRequest) SetRegion(v string) *DescribeAddonsRequest {
	s.Region = &v
	return s
}

func (s *DescribeAddonsRequest) SetClusterType(v string) *DescribeAddonsRequest {
	s.ClusterType = &v
	return s
}

type DescribeAddonsResponseBody struct {
	// 组件分组信息，例如：存储类组件，网络组件等。
	ComponentGroups []*DescribeAddonsResponseBodyComponentGroups `json:"ComponentGroups,omitempty" xml:"ComponentGroups,omitempty" type:"Repeated"`
	// 标准组件信息，包含各个组件的描述信息。
	StandardComponents *DescribeAddonsResponseBodyStandardComponents `json:"StandardComponents,omitempty" xml:"StandardComponents,omitempty" type:"Struct"`
}

func (s DescribeAddonsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponseBody) SetComponentGroups(v []*DescribeAddonsResponseBodyComponentGroups) *DescribeAddonsResponseBody {
	s.ComponentGroups = v
	return s
}

func (s *DescribeAddonsResponseBody) SetStandardComponents(v *DescribeAddonsResponseBodyStandardComponents) *DescribeAddonsResponseBody {
	s.StandardComponents = v
	return s
}

type DescribeAddonsResponseBodyComponentGroups struct {
	// 默认组件组。
	Default []*string `json:"default,omitempty" xml:"default,omitempty" type:"Repeated"`
	// 组件组名称。
	GroupName *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	// 组件清单。
	Items []*DescribeAddonsResponseBodyComponentGroupsItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s DescribeAddonsResponseBodyComponentGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsResponseBodyComponentGroups) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponseBodyComponentGroups) SetDefault(v []*string) *DescribeAddonsResponseBodyComponentGroups {
	s.Default = v
	return s
}

func (s *DescribeAddonsResponseBodyComponentGroups) SetGroupName(v string) *DescribeAddonsResponseBodyComponentGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeAddonsResponseBodyComponentGroups) SetItems(v []*DescribeAddonsResponseBodyComponentGroupsItems) *DescribeAddonsResponseBodyComponentGroups {
	s.Items = v
	return s
}

type DescribeAddonsResponseBodyComponentGroupsItems struct {
	// 组件描述信息。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 是否禁止默认安装。
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// 组件名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 是否为必需组件。
	Required *string `json:"required,omitempty" xml:"required,omitempty"`
	// 组件版本。
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeAddonsResponseBodyComponentGroupsItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsResponseBodyComponentGroupsItems) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponseBodyComponentGroupsItems) SetDescription(v string) *DescribeAddonsResponseBodyComponentGroupsItems {
	s.Description = &v
	return s
}

func (s *DescribeAddonsResponseBodyComponentGroupsItems) SetDisabled(v bool) *DescribeAddonsResponseBodyComponentGroupsItems {
	s.Disabled = &v
	return s
}

func (s *DescribeAddonsResponseBodyComponentGroupsItems) SetName(v string) *DescribeAddonsResponseBodyComponentGroupsItems {
	s.Name = &v
	return s
}

func (s *DescribeAddonsResponseBodyComponentGroupsItems) SetRequired(v string) *DescribeAddonsResponseBodyComponentGroupsItems {
	s.Required = &v
	return s
}

func (s *DescribeAddonsResponseBodyComponentGroupsItems) SetVersion(v string) *DescribeAddonsResponseBodyComponentGroupsItems {
	s.Version = &v
	return s
}

type DescribeAddonsResponseBodyStandardComponents struct {
	// 组件名称。
	ComponentName map[string]map[string]interface{} `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
}

func (s DescribeAddonsResponseBodyStandardComponents) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsResponseBodyStandardComponents) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponseBodyStandardComponents) SetComponentName(v map[string]map[string]interface{}) *DescribeAddonsResponseBodyStandardComponents {
	s.ComponentName = v
	return s
}

type DescribeAddonsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAddonsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAddonsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponse) SetHeaders(v map[string]*string) *DescribeAddonsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAddonsResponse) SetBody(v *DescribeAddonsResponseBody) *DescribeAddonsResponse {
	s.Body = v
	return s
}

type DescribeClusterAddonUpgradeStatusRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 组件ID。
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
}

func (s DescribeClusterAddonUpgradeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonUpgradeStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonUpgradeStatusRequest) SetClusterId(v string) *DescribeClusterAddonUpgradeStatusRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusRequest) SetComponentId(v string) *DescribeClusterAddonUpgradeStatusRequest {
	s.ComponentId = &v
	return s
}

type DescribeClusterAddonUpgradeStatusResponseBody struct {
	// 组件信息。
	AddonInfo *DescribeClusterAddonUpgradeStatusResponseBodyAddonInfo `json:"addon_info,omitempty" xml:"addon_info,omitempty" type:"Struct"`
	// 是否能够升级。
	CanUpgrade *bool `json:"can_upgrade,omitempty" xml:"can_upgrade,omitempty"`
	// 模板信息，采用base64加密。
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
}

func (s DescribeClusterAddonUpgradeStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonUpgradeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonUpgradeStatusResponseBody) SetAddonInfo(v *DescribeClusterAddonUpgradeStatusResponseBodyAddonInfo) *DescribeClusterAddonUpgradeStatusResponseBody {
	s.AddonInfo = v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponseBody) SetCanUpgrade(v bool) *DescribeClusterAddonUpgradeStatusResponseBody {
	s.CanUpgrade = &v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponseBody) SetTemplate(v string) *DescribeClusterAddonUpgradeStatusResponseBody {
	s.Template = &v
	return s
}

type DescribeClusterAddonUpgradeStatusResponseBodyAddonInfo struct {
	// Addon类别。
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 组件名称。
	ComponentName *string `json:"component_name,omitempty" xml:"component_name,omitempty"`
	// 升级说明信息。
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 组件版本。
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// 组件配置文件。
	Yaml *string `json:"yaml,omitempty" xml:"yaml,omitempty"`
}

func (s DescribeClusterAddonUpgradeStatusResponseBodyAddonInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonUpgradeStatusResponseBodyAddonInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyAddonInfo) SetCategory(v string) *DescribeClusterAddonUpgradeStatusResponseBodyAddonInfo {
	s.Category = &v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyAddonInfo) SetComponentName(v string) *DescribeClusterAddonUpgradeStatusResponseBodyAddonInfo {
	s.ComponentName = &v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyAddonInfo) SetMessage(v string) *DescribeClusterAddonUpgradeStatusResponseBodyAddonInfo {
	s.Message = &v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyAddonInfo) SetVersion(v string) *DescribeClusterAddonUpgradeStatusResponseBodyAddonInfo {
	s.Version = &v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyAddonInfo) SetYaml(v string) *DescribeClusterAddonUpgradeStatusResponseBodyAddonInfo {
	s.Yaml = &v
	return s
}

type DescribeClusterAddonUpgradeStatusResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterAddonUpgradeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterAddonUpgradeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonUpgradeStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonUpgradeStatusResponse) SetHeaders(v map[string]*string) *DescribeClusterAddonUpgradeStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponse) SetBody(v *DescribeClusterAddonUpgradeStatusResponseBody) *DescribeClusterAddonUpgradeStatusResponse {
	s.Body = v
	return s
}

type DescribeClusterAddonsUpgradeStatusRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 组件列表。
	ComponentIds []*string `json:"componentIds,omitempty" xml:"componentIds,omitempty" type:"Repeated"`
}

func (s DescribeClusterAddonsUpgradeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonsUpgradeStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonsUpgradeStatusRequest) SetClusterId(v string) *DescribeClusterAddonsUpgradeStatusRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterAddonsUpgradeStatusRequest) SetComponentIds(v []*string) *DescribeClusterAddonsUpgradeStatusRequest {
	s.ComponentIds = v
	return s
}

type DescribeClusterAddonsUpgradeStatusResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterAddonsUpgradeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonsUpgradeStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonsUpgradeStatusResponse) SetHeaders(v map[string]*string) *DescribeClusterAddonsUpgradeStatusResponse {
	s.Headers = v
	return s
}

type DescribeClusterAddonsVersionRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterAddonsVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonsVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonsVersionRequest) SetClusterId(v string) *DescribeClusterAddonsVersionRequest {
	s.ClusterId = &v
	return s
}

type DescribeClusterAddonsVersionResponseBody struct {
	// 组件名称。
	AddonsName map[string]map[string]interface{} `json:"AddonsName,omitempty" xml:"AddonsName,omitempty"`
}

func (s DescribeClusterAddonsVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonsVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonsVersionResponseBody) SetAddonsName(v map[string]map[string]interface{}) *DescribeClusterAddonsVersionResponseBody {
	s.AddonsName = v
	return s
}

type DescribeClusterAddonsVersionResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterAddonsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterAddonsVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonsVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonsVersionResponse) SetHeaders(v map[string]*string) *DescribeClusterAddonsVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterAddonsVersionResponse) SetBody(v *DescribeClusterAddonsVersionResponseBody) *DescribeClusterAddonsVersionResponse {
	s.Body = v
	return s
}

type DescribeClusterAttachScriptsRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 节点CPU架构,支持amd64、arm、arm64。
	Arch *string `json:"arch,omitempty" xml:"arch,omitempty"`
	// 边缘托管版集群节点的接入配置。
	Options *DescribeClusterAttachScriptsRequestOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s DescribeClusterAttachScriptsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttachScriptsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttachScriptsRequest) SetClusterId(v string) *DescribeClusterAttachScriptsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterAttachScriptsRequest) SetArch(v string) *DescribeClusterAttachScriptsRequest {
	s.Arch = &v
	return s
}

func (s *DescribeClusterAttachScriptsRequest) SetOptions(v *DescribeClusterAttachScriptsRequestOptions) *DescribeClusterAttachScriptsRequest {
	s.Options = v
	return s
}

type DescribeClusterAttachScriptsRequestOptions struct {
	// 需要安装的组件列表，默认为空，不安装。
	AllowedClusterAddons []*string `json:"allowedClusterAddons,omitempty" xml:"allowedClusterAddons,omitempty" type:"Repeated"`
	// 是否开启iptables，默认值true。
	EnableIptables *bool `json:"enableIptables,omitempty" xml:"enableIptables,omitempty"`
	// flannel使用的网卡名。默认使用节点默认路由的网卡名。
	FlannelIface *string `json:"flannelIface,omitempty" xml:"flannelIface,omitempty"`
	// 表示要接入的节点是否为GPU节点，默认为空。当前支持的GPU版本是Nvidia_Tesla_T4、Nvidia_Tesla_P4、Nvidia_Tesla_P100。
	GpuVersion *string `json:"gpuVersion,omitempty" xml:"gpuVersion,omitempty"`
	// 是否由edgeadm安装并检测Runtime，默认false。
	ManageRuntime *bool `json:"manageRuntime,omitempty" xml:"manageRuntime,omitempty"`
	// 设置节点名。  - ""（默认值，表示使用主机名。） - "*"（表示随机生成6位的字符串。） - "*.XXX"（表示随机生成6位字符串+XXX后缀。）
	NodeNameOverride *string `json:"nodeNameOverride,omitempty" xml:"nodeNameOverride,omitempty"`
	// 是否使用静默模式安装。
	Quiet *string `json:"quiet,omitempty" xml:"quiet,omitempty"`
}

func (s DescribeClusterAttachScriptsRequestOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttachScriptsRequestOptions) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttachScriptsRequestOptions) SetAllowedClusterAddons(v []*string) *DescribeClusterAttachScriptsRequestOptions {
	s.AllowedClusterAddons = v
	return s
}

func (s *DescribeClusterAttachScriptsRequestOptions) SetEnableIptables(v bool) *DescribeClusterAttachScriptsRequestOptions {
	s.EnableIptables = &v
	return s
}

func (s *DescribeClusterAttachScriptsRequestOptions) SetFlannelIface(v string) *DescribeClusterAttachScriptsRequestOptions {
	s.FlannelIface = &v
	return s
}

func (s *DescribeClusterAttachScriptsRequestOptions) SetGpuVersion(v string) *DescribeClusterAttachScriptsRequestOptions {
	s.GpuVersion = &v
	return s
}

func (s *DescribeClusterAttachScriptsRequestOptions) SetManageRuntime(v bool) *DescribeClusterAttachScriptsRequestOptions {
	s.ManageRuntime = &v
	return s
}

func (s *DescribeClusterAttachScriptsRequestOptions) SetNodeNameOverride(v string) *DescribeClusterAttachScriptsRequestOptions {
	s.NodeNameOverride = &v
	return s
}

func (s *DescribeClusterAttachScriptsRequestOptions) SetQuiet(v string) *DescribeClusterAttachScriptsRequestOptions {
	s.Quiet = &v
	return s
}

type DescribeClusterAttachScriptsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterAttachScriptsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttachScriptsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttachScriptsResponse) SetHeaders(v map[string]*string) *DescribeClusterAttachScriptsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterAttachScriptsResponse) SetBody(v string) *DescribeClusterAttachScriptsResponse {
	s.Body = &v
	return s
}

type DescribeClusterDetailRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailRequest) SetClusterId(v string) *DescribeClusterDetailRequest {
	s.ClusterId = &v
	return s
}

type DescribeClusterDetailResponseBody struct {
	// 集群ID。
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// 集群类型。
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// 集群创建时间。
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// 集群当前K8S版本。
	CurrentVersion *string `json:"current_version,omitempty" xml:"current_version,omitempty"`
	// 集群是否开启删除保护。
	DeletionProtection *bool `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// 集群内Docker版本。
	DockerVersion *string `json:"docker_version,omitempty" xml:"docker_version,omitempty"`
	// 集群Ingress SLB ID。
	ExternalLoadbalancerId *string `json:"external_loadbalancer_id,omitempty" xml:"external_loadbalancer_id,omitempty"`
	// 集群实例类型。
	InstanceType *string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
	// 集群元数据。
	MetaData *string `json:"meta_data,omitempty" xml:"meta_data,omitempty"`
	// 集群名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 集群采用的网络类型，例如：VPC网络。
	NetworkMode *string `json:"network_mode,omitempty" xml:"network_mode,omitempty"`
	// 集群所在地域ID。
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// 集群资源组ID。
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// 集群安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// 集群节点数量。
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// 集群运行状态。
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// 集群标签。
	Tags []*DescribeClusterDetailResponseBodyTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// 集群更新时间。
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// 集群使用的VPC ID。
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id,omitempty"`
	// 集群使用的虚拟交换机ID。
	VswitchCidr *string `json:"vswitch_cidr,omitempty" xml:"vswitch_cidr,omitempty"`
	// 集群节点使用的虚拟交换机列表。
	VswitchId *string `json:"vswitch_id,omitempty" xml:"vswitch_id,omitempty"`
	// 集群所在地域内的可用区ID。
	ZoneId *string `json:"zone_id,omitempty" xml:"zone_id,omitempty"`
}

func (s DescribeClusterDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBody) SetClusterId(v string) *DescribeClusterDetailResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetClusterType(v string) *DescribeClusterDetailResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetCreated(v string) *DescribeClusterDetailResponseBody {
	s.Created = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetCurrentVersion(v string) *DescribeClusterDetailResponseBody {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetDeletionProtection(v bool) *DescribeClusterDetailResponseBody {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetDockerVersion(v string) *DescribeClusterDetailResponseBody {
	s.DockerVersion = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetExternalLoadbalancerId(v string) *DescribeClusterDetailResponseBody {
	s.ExternalLoadbalancerId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetInstanceType(v string) *DescribeClusterDetailResponseBody {
	s.InstanceType = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetMetaData(v string) *DescribeClusterDetailResponseBody {
	s.MetaData = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetName(v string) *DescribeClusterDetailResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetNetworkMode(v string) *DescribeClusterDetailResponseBody {
	s.NetworkMode = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetRegionId(v string) *DescribeClusterDetailResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetResourceGroupId(v string) *DescribeClusterDetailResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetSecurityGroupId(v string) *DescribeClusterDetailResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetSize(v int32) *DescribeClusterDetailResponseBody {
	s.Size = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetState(v string) *DescribeClusterDetailResponseBody {
	s.State = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetTags(v []*DescribeClusterDetailResponseBodyTags) *DescribeClusterDetailResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetUpdated(v string) *DescribeClusterDetailResponseBody {
	s.Updated = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetVpcId(v string) *DescribeClusterDetailResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetVswitchCidr(v string) *DescribeClusterDetailResponseBody {
	s.VswitchCidr = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetVswitchId(v string) *DescribeClusterDetailResponseBody {
	s.VswitchId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetZoneId(v string) *DescribeClusterDetailResponseBody {
	s.ZoneId = &v
	return s
}

type DescribeClusterDetailResponseBodyTags struct {
	// 标签名。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 标签值。
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeClusterDetailResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyTags) SetKey(v string) *DescribeClusterDetailResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyTags) SetValue(v string) *DescribeClusterDetailResponseBodyTags {
	s.Value = &v
	return s
}

type DescribeClusterDetailResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponse) SetHeaders(v map[string]*string) *DescribeClusterDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterDetailResponse) SetBody(v *DescribeClusterDetailResponseBody) *DescribeClusterDetailResponse {
	s.Body = v
	return s
}

type DescribeClusterLogsRequest struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterLogsRequest) SetClusterId(v string) *DescribeClusterLogsRequest {
	s.ClusterId = &v
	return s
}

type DescribeClusterLogsResponseBody struct {
	// 日志ID。
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
	// 集群ID。
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// 日志详情。
	ClusterLog *string `json:"cluster_log,omitempty" xml:"cluster_log,omitempty"`
	// 日志创建时间。
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// 日志级别。
	LogLevel *string `json:"log_level,omitempty" xml:"log_level,omitempty"`
	// 日志更新时间。
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeClusterLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterLogsResponseBody) SetID(v string) *DescribeClusterLogsResponseBody {
	s.ID = &v
	return s
}

func (s *DescribeClusterLogsResponseBody) SetClusterId(v string) *DescribeClusterLogsResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterLogsResponseBody) SetClusterLog(v string) *DescribeClusterLogsResponseBody {
	s.ClusterLog = &v
	return s
}

func (s *DescribeClusterLogsResponseBody) SetCreated(v string) *DescribeClusterLogsResponseBody {
	s.Created = &v
	return s
}

func (s *DescribeClusterLogsResponseBody) SetLogLevel(v string) *DescribeClusterLogsResponseBody {
	s.LogLevel = &v
	return s
}

func (s *DescribeClusterLogsResponseBody) SetUpdated(v string) *DescribeClusterLogsResponseBody {
	s.Updated = &v
	return s
}

type DescribeClusterLogsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterLogsResponse) SetHeaders(v map[string]*string) *DescribeClusterLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterLogsResponse) SetBody(v *DescribeClusterLogsResponseBody) *DescribeClusterLogsResponse {
	s.Body = v
	return s
}

type DescribeClusterNodesRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 每页展示结果数。
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 结果只展示几页。
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 节点池ID。
	NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	// 节点状态信息。
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s DescribeClusterNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesRequest) SetClusterId(v string) *DescribeClusterNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterNodesRequest) SetPageSize(v string) *DescribeClusterNodesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterNodesRequest) SetPageNumber(v string) *DescribeClusterNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterNodesRequest) SetNodepoolId(v string) *DescribeClusterNodesRequest {
	s.NodepoolId = &v
	return s
}

func (s *DescribeClusterNodesRequest) SetState(v string) *DescribeClusterNodesRequest {
	s.State = &v
	return s
}

type DescribeClusterNodesResponseBody struct {
	// 节点信息列表。
	Nodes []*DescribeClusterNodesResponseBodyNodes `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// 分页信息。
	Page *DescribeClusterNodesResponseBodyPage `json:"page,omitempty" xml:"page,omitempty" type:"Struct"`
}

func (s DescribeClusterNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesResponseBody) SetNodes(v []*DescribeClusterNodesResponseBodyNodes) *DescribeClusterNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *DescribeClusterNodesResponseBody) SetPage(v *DescribeClusterNodesResponseBodyPage) *DescribeClusterNodesResponseBody {
	s.Page = v
	return s
}

type DescribeClusterNodesResponseBodyNodes struct {
	// 节点创建时间。
	CreationTime *string `json:"creation_time,omitempty" xml:"creation_time,omitempty"`
	// 错误信息说明。
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 节点过期时间。
	ExpiredTime *string `json:"expired_time,omitempty" xml:"expired_time,omitempty"`
	// 节点主机名。
	HostName *string `json:"host_name,omitempty" xml:"host_name,omitempty"`
	// 节点使用的镜像ID。
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// 节点付费类型。
	InstanceChargeType *string `json:"instance_charge_type,omitempty" xml:"instance_charge_type,omitempty"`
	// 节点实例ID。
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// 节点实例名称。
	InstanceName *string `json:"instance_name,omitempty" xml:"instance_name,omitempty"`
	// 节点实例角色类型，Master或Worker。
	InstanceRole *string `json:"instance_role,omitempty" xml:"instance_role,omitempty"`
	// 节点实例状态，
	InstanceStatus *string `json:"instance_status,omitempty" xml:"instance_status,omitempty"`
	// 节点实例类型。
	InstanceType *string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
	// 节点实例所属ECS实例簇名称。
	InstanceTypeFamily *string `json:"instance_type_family,omitempty" xml:"instance_type_family,omitempty"`
	// 节点IP地址。
	IpAddress []*string `json:"ip_address,omitempty" xml:"ip_address,omitempty" type:"Repeated"`
	// 节点是否为aliyun实例。
	IsAliyunNode *bool `json:"is_aliyun_node,omitempty" xml:"is_aliyun_node,omitempty"`
	// 节点名称，该名称是k8s专用名称。
	NodeName *string `json:"node_name,omitempty" xml:"node_name,omitempty"`
	// 节点状态，是否Ready。
	NodeStatus *string `json:"node_status,omitempty" xml:"node_status,omitempty"`
	// 节点所属的节点池ID。
	NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	// 节点通过什么方式创建出来的，例如：ROS。
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// ECS运行状态，例如：running。
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s DescribeClusterNodesResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesResponseBodyNodes) SetCreationTime(v string) *DescribeClusterNodesResponseBodyNodes {
	s.CreationTime = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetErrorMessage(v string) *DescribeClusterNodesResponseBodyNodes {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetExpiredTime(v string) *DescribeClusterNodesResponseBodyNodes {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetHostName(v string) *DescribeClusterNodesResponseBodyNodes {
	s.HostName = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetImageId(v string) *DescribeClusterNodesResponseBodyNodes {
	s.ImageId = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceChargeType(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceId(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceId = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceName(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceName = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceRole(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceRole = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceStatus(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceType(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceType = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceTypeFamily(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetIpAddress(v []*string) *DescribeClusterNodesResponseBodyNodes {
	s.IpAddress = v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetIsAliyunNode(v bool) *DescribeClusterNodesResponseBodyNodes {
	s.IsAliyunNode = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetNodeName(v string) *DescribeClusterNodesResponseBodyNodes {
	s.NodeName = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetNodeStatus(v string) *DescribeClusterNodesResponseBodyNodes {
	s.NodeStatus = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetNodepoolId(v string) *DescribeClusterNodesResponseBodyNodes {
	s.NodepoolId = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetSource(v string) *DescribeClusterNodesResponseBodyNodes {
	s.Source = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetState(v string) *DescribeClusterNodesResponseBodyNodes {
	s.State = &v
	return s
}

type DescribeClusterNodesResponseBodyPage struct {
	// 总页数。
	PageNumber *int32 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// 单页展示结果数量。
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 结果总条数。
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s DescribeClusterNodesResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodesResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesResponseBodyPage) SetPageNumber(v int32) *DescribeClusterNodesResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyPage) SetPageSize(v int32) *DescribeClusterNodesResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyPage) SetTotalCount(v int32) *DescribeClusterNodesResponseBodyPage {
	s.TotalCount = &v
	return s
}

type DescribeClusterNodesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesResponse) SetHeaders(v map[string]*string) *DescribeClusterNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterNodesResponse) SetBody(v *DescribeClusterNodesResponseBody) *DescribeClusterNodesResponse {
	s.Body = v
	return s
}

type DescribeClusterResourcesRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourcesRequest) SetClusterId(v string) *DescribeClusterResourcesRequest {
	s.ClusterId = &v
	return s
}

type DescribeClusterResourcesResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    []*DescribeClusterResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeClusterResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourcesResponse) SetHeaders(v map[string]*string) *DescribeClusterResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterResourcesResponse) SetBody(v []*DescribeClusterResourcesResponseBody) *DescribeClusterResourcesResponse {
	s.Body = v
	return s
}

type DescribeClusterResourcesResponseBody struct {
	// 集群ID。
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// 资源创建时间。
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// 资源实例ID。
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// 资源元信息。
	ResourceInfo *string `json:"resource_info,omitempty" xml:"resource_info,omitempty"`
	// 资源类型。
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// 资源状态。
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s DescribeClusterResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourcesResponseBody) SetClusterId(v string) *DescribeClusterResourcesResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterResourcesResponseBody) SetCreated(v string) *DescribeClusterResourcesResponseBody {
	s.Created = &v
	return s
}

func (s *DescribeClusterResourcesResponseBody) SetInstanceId(v string) *DescribeClusterResourcesResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeClusterResourcesResponseBody) SetResourceInfo(v string) *DescribeClusterResourcesResponseBody {
	s.ResourceInfo = &v
	return s
}

func (s *DescribeClusterResourcesResponseBody) SetResourceType(v string) *DescribeClusterResourcesResponseBody {
	s.ResourceType = &v
	return s
}

func (s *DescribeClusterResourcesResponseBody) SetState(v string) *DescribeClusterResourcesResponseBody {
	s.State = &v
	return s
}

type DescribeClusterUserKubeconfigRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// ApiServer是否为内网地址。
	PrivateIpAddress *bool `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s DescribeClusterUserKubeconfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterUserKubeconfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterUserKubeconfigRequest) SetClusterId(v string) *DescribeClusterUserKubeconfigRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterUserKubeconfigRequest) SetPrivateIpAddress(v bool) *DescribeClusterUserKubeconfigRequest {
	s.PrivateIpAddress = &v
	return s
}

type DescribeClusterUserKubeconfigResponseBody struct {
	// kubeconfig内容。
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
}

func (s DescribeClusterUserKubeconfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterUserKubeconfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterUserKubeconfigResponseBody) SetConfig(v string) *DescribeClusterUserKubeconfigResponseBody {
	s.Config = &v
	return s
}

type DescribeClusterUserKubeconfigResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterUserKubeconfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterUserKubeconfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterUserKubeconfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterUserKubeconfigResponse) SetHeaders(v map[string]*string) *DescribeClusterUserKubeconfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterUserKubeconfigResponse) SetBody(v *DescribeClusterUserKubeconfigResponseBody) *DescribeClusterUserKubeconfigResponse {
	s.Body = v
	return s
}

type DescribeClusterV2UserKubeconfigRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 是否为内网访问。
	PrivateIpAddress *bool `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s DescribeClusterV2UserKubeconfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2UserKubeconfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2UserKubeconfigRequest) SetClusterId(v string) *DescribeClusterV2UserKubeconfigRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterV2UserKubeconfigRequest) SetPrivateIpAddress(v bool) *DescribeClusterV2UserKubeconfigRequest {
	s.PrivateIpAddress = &v
	return s
}

type DescribeClusterV2UserKubeconfigResponseBody struct {
	// kubeconfig内容。
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
}

func (s DescribeClusterV2UserKubeconfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2UserKubeconfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2UserKubeconfigResponseBody) SetConfig(v string) *DescribeClusterV2UserKubeconfigResponseBody {
	s.Config = &v
	return s
}

type DescribeClusterV2UserKubeconfigResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterV2UserKubeconfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterV2UserKubeconfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2UserKubeconfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2UserKubeconfigResponse) SetHeaders(v map[string]*string) *DescribeClusterV2UserKubeconfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterV2UserKubeconfigResponse) SetBody(v *DescribeClusterV2UserKubeconfigResponseBody) *DescribeClusterV2UserKubeconfigResponse {
	s.Body = v
	return s
}

type DescribeClustersRequest struct {
	// 集群名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 集群类型。
	ClusterType *string `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
}

func (s DescribeClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeClustersRequest) SetName(v string) *DescribeClustersRequest {
	s.Name = &v
	return s
}

func (s *DescribeClustersRequest) SetClusterType(v string) *DescribeClustersRequest {
	s.ClusterType = &v
	return s
}

type DescribeClustersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    []*DescribeClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponse) SetHeaders(v map[string]*string) *DescribeClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeClustersResponse) SetBody(v []*DescribeClustersResponseBody) *DescribeClustersResponse {
	s.Body = v
	return s
}

type DescribeClustersResponseBody struct {
	// 集群ID。
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// 集群类型。
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// 集群创建时间。
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// 集群当前版本。
	CurrentVersion *string `json:"current_version,omitempty" xml:"current_version,omitempty"`
	// 节点系统盘类型。
	DataDiskCategory *string `json:"data_disk_category,omitempty" xml:"data_disk_category,omitempty"`
	// 节点系统盘大小。
	DataDiskSize *int64 `json:"data_disk_size,omitempty" xml:"data_disk_size,omitempty"`
	// 集群是否开启删除保护。
	DeletionProtection *bool `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// 容器运行时版本。
	DockerVersion *string `json:"docker_version,omitempty" xml:"docker_version,omitempty"`
	// 集群Ingerss SLB实例的ID。
	ExternalLoadbalancerId *string `json:"external_loadbalancer_id,omitempty" xml:"external_loadbalancer_id,omitempty"`
	// 集群创建时版本。
	InitVersion *string `json:"init_version,omitempty" xml:"init_version,omitempty"`
	// 集群的endpoint地址。
	MasterUrl *string `json:"master_url,omitempty" xml:"master_url,omitempty"`
	// 集群元数据。
	MetaData *string `json:"meta_data,omitempty" xml:"meta_data,omitempty"`
	// 集群名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 集群使用的网络类型。
	NetworkMode *string `json:"network_mode,omitempty" xml:"network_mode,omitempty"`
	// 集群是否开启Private Zone，默认false。
	PrivateZone *bool `json:"private_zone,omitempty" xml:"private_zone,omitempty"`
	// 集群标识，区分是否为边缘托管版。
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
	// 集群所在地域ID。
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// 集群资源组ID。
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// 集群安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// 集群内实例数量。
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 集群运行状态。
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// POD网络。
	SubnetCidr *string `json:"subnet_cidr,omitempty" xml:"subnet_cidr,omitempty"`
	// 集群标签。
	Tags []*DescribeClustersResponseBodyTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// 集群更新时间。
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// 集群使用的VPC ID。
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id,omitempty"`
	// 虚拟交换机网络ID。
	VswitchCidr *string `json:"vswitch_cidr,omitempty" xml:"vswitch_cidr,omitempty"`
	// 节点使用的Vswitch ID。
	VswitchId *string `json:"vswitch_id,omitempty" xml:"vswitch_id,omitempty"`
	// 集群Worker节点RAM角色名称。
	WorkerRamRoleName *string `json:"worker_ram_role_name,omitempty" xml:"worker_ram_role_name,omitempty"`
	// 集群所在Region内的区域ID。
	ZoneId *string `json:"zone_id,omitempty" xml:"zone_id,omitempty"`
}

func (s DescribeClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBody) SetClusterId(v string) *DescribeClustersResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClustersResponseBody) SetClusterType(v string) *DescribeClustersResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeClustersResponseBody) SetCreated(v string) *DescribeClustersResponseBody {
	s.Created = &v
	return s
}

func (s *DescribeClustersResponseBody) SetCurrentVersion(v string) *DescribeClustersResponseBody {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeClustersResponseBody) SetDataDiskCategory(v string) *DescribeClustersResponseBody {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeClustersResponseBody) SetDataDiskSize(v int64) *DescribeClustersResponseBody {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeClustersResponseBody) SetDeletionProtection(v bool) *DescribeClustersResponseBody {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeClustersResponseBody) SetDockerVersion(v string) *DescribeClustersResponseBody {
	s.DockerVersion = &v
	return s
}

func (s *DescribeClustersResponseBody) SetExternalLoadbalancerId(v string) *DescribeClustersResponseBody {
	s.ExternalLoadbalancerId = &v
	return s
}

func (s *DescribeClustersResponseBody) SetInitVersion(v string) *DescribeClustersResponseBody {
	s.InitVersion = &v
	return s
}

func (s *DescribeClustersResponseBody) SetMasterUrl(v string) *DescribeClustersResponseBody {
	s.MasterUrl = &v
	return s
}

func (s *DescribeClustersResponseBody) SetMetaData(v string) *DescribeClustersResponseBody {
	s.MetaData = &v
	return s
}

func (s *DescribeClustersResponseBody) SetName(v string) *DescribeClustersResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeClustersResponseBody) SetNetworkMode(v string) *DescribeClustersResponseBody {
	s.NetworkMode = &v
	return s
}

func (s *DescribeClustersResponseBody) SetPrivateZone(v bool) *DescribeClustersResponseBody {
	s.PrivateZone = &v
	return s
}

func (s *DescribeClustersResponseBody) SetProfile(v string) *DescribeClustersResponseBody {
	s.Profile = &v
	return s
}

func (s *DescribeClustersResponseBody) SetRegionId(v string) *DescribeClustersResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeClustersResponseBody) SetResourceGroupId(v string) *DescribeClustersResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClustersResponseBody) SetSecurityGroupId(v string) *DescribeClustersResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClustersResponseBody) SetSize(v int64) *DescribeClustersResponseBody {
	s.Size = &v
	return s
}

func (s *DescribeClustersResponseBody) SetState(v string) *DescribeClustersResponseBody {
	s.State = &v
	return s
}

func (s *DescribeClustersResponseBody) SetSubnetCidr(v string) *DescribeClustersResponseBody {
	s.SubnetCidr = &v
	return s
}

func (s *DescribeClustersResponseBody) SetTags(v []*DescribeClustersResponseBodyTags) *DescribeClustersResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeClustersResponseBody) SetUpdated(v string) *DescribeClustersResponseBody {
	s.Updated = &v
	return s
}

func (s *DescribeClustersResponseBody) SetVpcId(v string) *DescribeClustersResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeClustersResponseBody) SetVswitchCidr(v string) *DescribeClustersResponseBody {
	s.VswitchCidr = &v
	return s
}

func (s *DescribeClustersResponseBody) SetVswitchId(v string) *DescribeClustersResponseBody {
	s.VswitchId = &v
	return s
}

func (s *DescribeClustersResponseBody) SetWorkerRamRoleName(v string) *DescribeClustersResponseBody {
	s.WorkerRamRoleName = &v
	return s
}

func (s *DescribeClustersResponseBody) SetZoneId(v string) *DescribeClustersResponseBody {
	s.ZoneId = &v
	return s
}

type DescribeClustersResponseBodyTags struct {
	// 标签名。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 标签值。
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeClustersResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBodyTags) SetKey(v string) *DescribeClustersResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeClustersResponseBodyTags) SetValue(v string) *DescribeClustersResponseBodyTags {
	s.Value = &v
	return s
}

type DescribeClustersV1Request struct {
	// 集群名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 集群类型。
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// 单页大小。
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 分页数。
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
}

func (s DescribeClustersV1Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersV1Request) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1Request) SetName(v string) *DescribeClustersV1Request {
	s.Name = &v
	return s
}

func (s *DescribeClustersV1Request) SetClusterType(v string) *DescribeClustersV1Request {
	s.ClusterType = &v
	return s
}

func (s *DescribeClustersV1Request) SetPageSize(v int64) *DescribeClustersV1Request {
	s.PageSize = &v
	return s
}

func (s *DescribeClustersV1Request) SetPageNumber(v int64) *DescribeClustersV1Request {
	s.PageNumber = &v
	return s
}

type DescribeClustersV1ResponseBody struct {
	// 集群详情列表。
	Clusters []*DescribeClustersV1ResponseBodyClusters `json:"clusters,omitempty" xml:"clusters,omitempty" type:"Repeated"`
	// 分页信息。
	PageInfo *DescribeClustersV1ResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
}

func (s DescribeClustersV1ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersV1ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1ResponseBody) SetClusters(v []*DescribeClustersV1ResponseBodyClusters) *DescribeClustersV1ResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeClustersV1ResponseBody) SetPageInfo(v *DescribeClustersV1ResponseBodyPageInfo) *DescribeClustersV1ResponseBody {
	s.PageInfo = v
	return s
}

type DescribeClustersV1ResponseBodyClusters struct {
	// 集群健康状态。
	ClusterHealthy *string `json:"cluster_healthy,omitempty" xml:"cluster_healthy,omitempty"`
	// 集群ID。
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// 集群类型。
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// 集群创建时间。
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// 集群当前版本。
	CurrentVersion *string `json:"current_version,omitempty" xml:"current_version,omitempty"`
	// 数据盘类型。
	DataDiskCategory *string `json:"data_disk_category,omitempty" xml:"data_disk_category,omitempty"`
	// 数据盘大小。
	DataDiskSize *int64 `json:"data_disk_size,omitempty" xml:"data_disk_size,omitempty"`
	// 集群是否开启删除保护。
	DeletionProtection *bool `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// 集群使用的Docker版本。
	DockerVersion *string `json:"docker_version,omitempty" xml:"docker_version,omitempty"`
	// 集群Ingress对应的SLB的地址。
	ExternalLoadbalancerId *string `json:"external_loadbalancer_id,omitempty" xml:"external_loadbalancer_id,omitempty"`
	// 集群初始化版本。
	InitVersion *string `json:"init_version,omitempty" xml:"init_version,omitempty"`
	// 集群访问的端点信息。
	MasterUrl *string `json:"master_url,omitempty" xml:"master_url,omitempty"`
	// 集群元数据信息。
	MetaData *string `json:"meta_data,omitempty" xml:"meta_data,omitempty"`
	// 集群名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 集群使用的网络类型，例如：VPC网络。
	NetworkMode *string `json:"network_mode,omitempty" xml:"network_mode,omitempty"`
	// 节点状态。
	NodeStatus *string `json:"node_status,omitempty" xml:"node_status,omitempty"`
	// 集群是否开启Private Zone。
	PrivateZone *bool `json:"private_zone,omitempty" xml:"private_zone,omitempty"`
	// 边缘集群表示，用于区分边缘托管版集群。
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
	// 地域ID。
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// 集群资源组ID。
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// 集群安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// 集群节点数。
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 集群运行状态。
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// POD网段地址。
	SubnetCidr *string `json:"subnet_cidr,omitempty" xml:"subnet_cidr,omitempty"`
	// 集群标签。
	Tags []*DescribeClustersV1ResponseBodyClustersTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// 集群更新时间。
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// 集群所在的VPC ID。
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id,omitempty"`
	// 虚拟交换机CIDR。
	VswitchCidr *string `json:"vswitch_cidr,omitempty" xml:"vswitch_cidr,omitempty"`
	// 集群使用的虚拟交换ID。
	VswitchId *string `json:"vswitch_id,omitempty" xml:"vswitch_id,omitempty"`
	// 集群Worker RAM角色。
	WorkerRamRoleName *string `json:"worker_ram_role_name,omitempty" xml:"worker_ram_role_name,omitempty"`
	// 可用区ID。
	ZoneId *string `json:"zone_id,omitempty" xml:"zone_id,omitempty"`
}

func (s DescribeClustersV1ResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersV1ResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1ResponseBodyClusters) SetClusterHealthy(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ClusterHealthy = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetClusterId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetClusterType(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ClusterType = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetCreated(v string) *DescribeClustersV1ResponseBodyClusters {
	s.Created = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetCurrentVersion(v string) *DescribeClustersV1ResponseBodyClusters {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetDataDiskCategory(v string) *DescribeClustersV1ResponseBodyClusters {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetDataDiskSize(v int64) *DescribeClustersV1ResponseBodyClusters {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetDeletionProtection(v bool) *DescribeClustersV1ResponseBodyClusters {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetDockerVersion(v string) *DescribeClustersV1ResponseBodyClusters {
	s.DockerVersion = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetExternalLoadbalancerId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ExternalLoadbalancerId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetInitVersion(v string) *DescribeClustersV1ResponseBodyClusters {
	s.InitVersion = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetMasterUrl(v string) *DescribeClustersV1ResponseBodyClusters {
	s.MasterUrl = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetMetaData(v string) *DescribeClustersV1ResponseBodyClusters {
	s.MetaData = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetName(v string) *DescribeClustersV1ResponseBodyClusters {
	s.Name = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetNetworkMode(v string) *DescribeClustersV1ResponseBodyClusters {
	s.NetworkMode = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetNodeStatus(v string) *DescribeClustersV1ResponseBodyClusters {
	s.NodeStatus = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetPrivateZone(v bool) *DescribeClustersV1ResponseBodyClusters {
	s.PrivateZone = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetProfile(v string) *DescribeClustersV1ResponseBodyClusters {
	s.Profile = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetRegionId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.RegionId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetResourceGroupId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetSecurityGroupId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetSize(v int64) *DescribeClustersV1ResponseBodyClusters {
	s.Size = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetState(v string) *DescribeClustersV1ResponseBodyClusters {
	s.State = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetSubnetCidr(v string) *DescribeClustersV1ResponseBodyClusters {
	s.SubnetCidr = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetTags(v []*DescribeClustersV1ResponseBodyClustersTags) *DescribeClustersV1ResponseBodyClusters {
	s.Tags = v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetUpdated(v string) *DescribeClustersV1ResponseBodyClusters {
	s.Updated = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetVpcId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.VpcId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetVswitchCidr(v string) *DescribeClustersV1ResponseBodyClusters {
	s.VswitchCidr = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetVswitchId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.VswitchId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetWorkerRamRoleName(v string) *DescribeClustersV1ResponseBodyClusters {
	s.WorkerRamRoleName = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetZoneId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ZoneId = &v
	return s
}

type DescribeClustersV1ResponseBodyClustersTags struct {
	// 标签键。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 标签值。
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeClustersV1ResponseBodyClustersTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersV1ResponseBodyClustersTags) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1ResponseBodyClustersTags) SetKey(v string) *DescribeClustersV1ResponseBodyClustersTags {
	s.Key = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClustersTags) SetValue(v string) *DescribeClustersV1ResponseBodyClustersTags {
	s.Value = &v
	return s
}

type DescribeClustersV1ResponseBodyPageInfo struct {
	// 分页总数。
	PageNumber *int32 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// 单页大小。
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 结果总条目数。
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s DescribeClustersV1ResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersV1ResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1ResponseBodyPageInfo) SetPageNumber(v int32) *DescribeClustersV1ResponseBodyPageInfo {
	s.PageNumber = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyPageInfo) SetPageSize(v int32) *DescribeClustersV1ResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyPageInfo) SetTotalCount(v int32) *DescribeClustersV1ResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeClustersV1Response struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClustersV1ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClustersV1Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersV1Response) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1Response) SetHeaders(v map[string]*string) *DescribeClustersV1Response {
	s.Headers = v
	return s
}

func (s *DescribeClustersV1Response) SetBody(v *DescribeClustersV1ResponseBody) *DescribeClustersV1Response {
	s.Body = v
	return s
}

type DescribeExternalAgentRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeExternalAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExternalAgentRequest) GoString() string {
	return s.String()
}

func (s *DescribeExternalAgentRequest) SetClusterId(v string) *DescribeExternalAgentRequest {
	s.ClusterId = &v
	return s
}

type DescribeExternalAgentResponseBody struct {
	// 代理配置。
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
}

func (s DescribeExternalAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExternalAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExternalAgentResponseBody) SetConfig(v string) *DescribeExternalAgentResponseBody {
	s.Config = &v
	return s
}

type DescribeExternalAgentResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExternalAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExternalAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExternalAgentResponse) GoString() string {
	return s.String()
}

func (s *DescribeExternalAgentResponse) SetHeaders(v map[string]*string) *DescribeExternalAgentResponse {
	s.Headers = v
	return s
}

func (s *DescribeExternalAgentResponse) SetBody(v *DescribeExternalAgentResponseBody) *DescribeExternalAgentResponse {
	s.Body = v
	return s
}

type DescribeTemplatesRequest struct {
	// 模板类型，部署模板类型，目前一共有2种类型，取值为：kubernetes或compose。
	TemplateType *string `json:"template_type,omitempty" xml:"template_type,omitempty"`
}

func (s DescribeTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesRequest) SetTemplateType(v string) *DescribeTemplatesRequest {
	s.TemplateType = &v
	return s
}

type DescribeTemplatesResponseBody struct {
	// 分页信息。
	PageInfo *DescribeTemplatesResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
	// 模板列表。
	Templates []*DescribeTemplatesResponseBodyTemplates `json:"templates,omitempty" xml:"templates,omitempty" type:"Repeated"`
}

func (s DescribeTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBody) SetPageInfo(v *DescribeTemplatesResponseBodyPageInfo) *DescribeTemplatesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeTemplatesResponseBody) SetTemplates(v []*DescribeTemplatesResponseBodyTemplates) *DescribeTemplatesResponseBody {
	s.Templates = v
	return s
}

type DescribeTemplatesResponseBodyPageInfo struct {
	// 分页数。
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// 单页大小。
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 结果总数。
	TotalCount *int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s DescribeTemplatesResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBodyPageInfo) SetPageNumber(v int64) *DescribeTemplatesResponseBodyPageInfo {
	s.PageNumber = &v
	return s
}

func (s *DescribeTemplatesResponseBodyPageInfo) SetPageSize(v int64) *DescribeTemplatesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeTemplatesResponseBodyPageInfo) SetTotalCount(v int64) *DescribeTemplatesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeTemplatesResponseBodyTemplates struct {
	// 模板访问权限，取值为：private、pubilc或shared。。
	Acl *string `json:"acl,omitempty" xml:"acl,omitempty"`
	// 模板创建时间。
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// 模板描述信息。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 模板ID。
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 模板名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 模板标签，如果不显式指定了，默认为模板名称。
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// 模板详情。
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// 部署模板类型，目前只有kubernetes一种生效。
	TemplateType *string `json:"template_type,omitempty" xml:"template_type,omitempty"`
	// 模板修改时间。
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBodyTemplates) SetAcl(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Acl = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetCreated(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Created = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetDescription(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetId(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Id = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetName(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetTags(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Tags = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetTemplate(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Template = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetTemplateType(v string) *DescribeTemplatesResponseBodyTemplates {
	s.TemplateType = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetUpdated(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Updated = &v
	return s
}

type DescribeTemplatesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponse) SetHeaders(v map[string]*string) *DescribeTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplatesResponse) SetBody(v *DescribeTemplatesResponseBody) *DescribeTemplatesResponse {
	s.Body = v
	return s
}

type DescribeUserQuotaResponseBody struct {
	// 托管版集群配额。
	AmkClusterQuota *int64 `json:"amk_cluster_quota,omitempty" xml:"amk_cluster_quota,omitempty"`
	// Serverless集群配额。
	AskClusterQuota *int64 `json:"ask_cluster_quota,omitempty" xml:"ask_cluster_quota,omitempty"`
	// 集群节点池配额。
	ClusterNodepoolQuota *int64 `json:"cluster_nodepool_quota,omitempty" xml:"cluster_nodepool_quota,omitempty"`
	// 专有版集群托管版集群的总配额。
	ClusterQuota *int64 `json:"cluster_quota,omitempty" xml:"cluster_quota,omitempty"`
	// 单集群的节点配额。
	NodeQuota *int64 `json:"node_quota,omitempty" xml:"node_quota,omitempty"`
}

func (s DescribeUserQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserQuotaResponseBody) SetAmkClusterQuota(v int64) *DescribeUserQuotaResponseBody {
	s.AmkClusterQuota = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetAskClusterQuota(v int64) *DescribeUserQuotaResponseBody {
	s.AskClusterQuota = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetClusterNodepoolQuota(v int64) *DescribeUserQuotaResponseBody {
	s.ClusterNodepoolQuota = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetClusterQuota(v int64) *DescribeUserQuotaResponseBody {
	s.ClusterQuota = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetNodeQuota(v int64) *DescribeUserQuotaResponseBody {
	s.NodeQuota = &v
	return s
}

type DescribeUserQuotaResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserQuotaResponse) SetHeaders(v map[string]*string) *DescribeUserQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserQuotaResponse) SetBody(v *DescribeUserQuotaResponseBody) *DescribeUserQuotaResponse {
	s.Body = v
	return s
}

type GetKubernetesTriggerRequest struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 应用所属命名空间。
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 应用类型。
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 应用名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetKubernetesTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s GetKubernetesTriggerRequest) GoString() string {
	return s.String()
}

func (s *GetKubernetesTriggerRequest) SetClusterId(v string) *GetKubernetesTriggerRequest {
	s.ClusterId = &v
	return s
}

func (s *GetKubernetesTriggerRequest) SetNamespace(v string) *GetKubernetesTriggerRequest {
	s.Namespace = &v
	return s
}

func (s *GetKubernetesTriggerRequest) SetType(v string) *GetKubernetesTriggerRequest {
	s.Type = &v
	return s
}

func (s *GetKubernetesTriggerRequest) SetName(v string) *GetKubernetesTriggerRequest {
	s.Name = &v
	return s
}

type GetKubernetesTriggerResponseBody struct {
	// 触发器详情。
	Triggers []*GetKubernetesTriggerResponseBodyTriggers `json:"triggers,omitempty" xml:"triggers,omitempty" type:"Repeated"`
}

func (s GetKubernetesTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetKubernetesTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *GetKubernetesTriggerResponseBody) SetTriggers(v []*GetKubernetesTriggerResponseBodyTriggers) *GetKubernetesTriggerResponseBody {
	s.Triggers = v
	return s
}

type GetKubernetesTriggerResponseBodyTriggers struct {
	// 触发器行为。
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 集群ID。
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// 触发器ID。
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 项目ID，格式为：${namepsce}/${应用名}，
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// 触发器Token。
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GetKubernetesTriggerResponseBodyTriggers) String() string {
	return tea.Prettify(s)
}

func (s GetKubernetesTriggerResponseBodyTriggers) GoString() string {
	return s.String()
}

func (s *GetKubernetesTriggerResponseBodyTriggers) SetAction(v string) *GetKubernetesTriggerResponseBodyTriggers {
	s.Action = &v
	return s
}

func (s *GetKubernetesTriggerResponseBodyTriggers) SetClusterId(v string) *GetKubernetesTriggerResponseBodyTriggers {
	s.ClusterId = &v
	return s
}

func (s *GetKubernetesTriggerResponseBodyTriggers) SetId(v string) *GetKubernetesTriggerResponseBodyTriggers {
	s.Id = &v
	return s
}

func (s *GetKubernetesTriggerResponseBodyTriggers) SetProjectId(v string) *GetKubernetesTriggerResponseBodyTriggers {
	s.ProjectId = &v
	return s
}

func (s *GetKubernetesTriggerResponseBodyTriggers) SetToken(v string) *GetKubernetesTriggerResponseBodyTriggers {
	s.Token = &v
	return s
}

type GetKubernetesTriggerResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetKubernetesTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetKubernetesTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetKubernetesTriggerResponse) GoString() string {
	return s.String()
}

func (s *GetKubernetesTriggerResponse) SetHeaders(v map[string]*string) *GetKubernetesTriggerResponse {
	s.Headers = v
	return s
}

func (s *GetKubernetesTriggerResponse) SetBody(v *GetKubernetesTriggerResponseBody) *GetKubernetesTriggerResponse {
	s.Body = v
	return s
}

type GetUpgradeStatusRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetUpgradeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUpgradeStatusRequest) GoString() string {
	return s.String()
}

func (s *GetUpgradeStatusRequest) SetClusterId(v string) *GetUpgradeStatusRequest {
	s.ClusterId = &v
	return s
}

type GetUpgradeStatusResponseBody struct {
	// 错误信息描述。
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 预检查返回ID。
	PrecheckReportId *string `json:"precheck_report_id,omitempty" xml:"precheck_report_id,omitempty"`
	// 升级状态。
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 升级任务执行到哪一步了。
	UpgradeStep *string `json:"upgrade_step,omitempty" xml:"upgrade_step,omitempty"`
}

func (s GetUpgradeStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUpgradeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetUpgradeStatusResponseBody) SetErrorMessage(v string) *GetUpgradeStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUpgradeStatusResponseBody) SetPrecheckReportId(v string) *GetUpgradeStatusResponseBody {
	s.PrecheckReportId = &v
	return s
}

func (s *GetUpgradeStatusResponseBody) SetStatus(v string) *GetUpgradeStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetUpgradeStatusResponseBody) SetUpgradeStep(v string) *GetUpgradeStatusResponseBody {
	s.UpgradeStep = &v
	return s
}

type GetUpgradeStatusResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUpgradeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUpgradeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUpgradeStatusResponse) GoString() string {
	return s.String()
}

func (s *GetUpgradeStatusResponse) SetHeaders(v map[string]*string) *GetUpgradeStatusResponse {
	s.Headers = v
	return s
}

func (s *GetUpgradeStatusResponse) SetBody(v *GetUpgradeStatusResponseBody) *GetUpgradeStatusResponse {
	s.Body = v
	return s
}

type InstallClusterAddonsRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Addon列表。
	Body []*InstallClusterAddonsRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s InstallClusterAddonsRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallClusterAddonsRequest) GoString() string {
	return s.String()
}

func (s *InstallClusterAddonsRequest) SetClusterId(v string) *InstallClusterAddonsRequest {
	s.ClusterId = &v
	return s
}

func (s *InstallClusterAddonsRequest) SetBody(v []*InstallClusterAddonsRequestBody) *InstallClusterAddonsRequest {
	s.Body = v
	return s
}

type InstallClusterAddonsRequestBody struct {
	// Addon配置信息。
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// 是否禁止默认安装。
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// Addon名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 是否默认安装。
	Required *string `json:"required,omitempty" xml:"required,omitempty"`
	// Addon版本号。
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s InstallClusterAddonsRequestBody) String() string {
	return tea.Prettify(s)
}

func (s InstallClusterAddonsRequestBody) GoString() string {
	return s.String()
}

func (s *InstallClusterAddonsRequestBody) SetConfig(v string) *InstallClusterAddonsRequestBody {
	s.Config = &v
	return s
}

func (s *InstallClusterAddonsRequestBody) SetDisabled(v bool) *InstallClusterAddonsRequestBody {
	s.Disabled = &v
	return s
}

func (s *InstallClusterAddonsRequestBody) SetName(v string) *InstallClusterAddonsRequestBody {
	s.Name = &v
	return s
}

func (s *InstallClusterAddonsRequestBody) SetRequired(v string) *InstallClusterAddonsRequestBody {
	s.Required = &v
	return s
}

func (s *InstallClusterAddonsRequestBody) SetVersion(v string) *InstallClusterAddonsRequestBody {
	s.Version = &v
	return s
}

type InstallClusterAddonsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s InstallClusterAddonsResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallClusterAddonsResponse) GoString() string {
	return s.String()
}

func (s *InstallClusterAddonsResponse) SetHeaders(v map[string]*string) *InstallClusterAddonsResponse {
	s.Headers = v
	return s
}

type ModifyClusterRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 集群是否开启删除保护。
	DeletionProtection *bool `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// 集群的Ingress SLB的ID。
	IngressLoadbalancerId *string `json:"ingress_loadbalancer_id,omitempty" xml:"ingress_loadbalancer_id,omitempty"`
	// 集群是否开启EIP。
	ApiServerEip *bool `json:"api_server_eip,omitempty" xml:"api_server_eip,omitempty"`
	// 集群的API Server的EIP ID。
	ApiServerEipId *string `json:"api_server_eip_id,omitempty" xml:"api_server_eip_id,omitempty"`
	// 集群资源组ID。
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// 域名是否重新绑定到Ingress的SLB地址。
	IngressDomainRebinding *string `json:"ingress_domain_rebinding,omitempty" xml:"ingress_domain_rebinding,omitempty"`
}

func (s ModifyClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterRequest) SetClusterId(v string) *ModifyClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterRequest) SetDeletionProtection(v bool) *ModifyClusterRequest {
	s.DeletionProtection = &v
	return s
}

func (s *ModifyClusterRequest) SetIngressLoadbalancerId(v string) *ModifyClusterRequest {
	s.IngressLoadbalancerId = &v
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

func (s *ModifyClusterRequest) SetResourceGroupId(v string) *ModifyClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyClusterRequest) SetIngressDomainRebinding(v string) *ModifyClusterRequest {
	s.IngressDomainRebinding = &v
	return s
}

type ModifyClusterResponseBody struct {
	// 集群ID。
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// 请求ID。
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 任务ID。
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s ModifyClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterResponseBody) SetClusterId(v string) *ModifyClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterResponseBody) SetRequestId(v string) *ModifyClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterResponseBody) SetTaskId(v string) *ModifyClusterResponseBody {
	s.TaskId = &v
	return s
}

type ModifyClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterResponse) SetHeaders(v map[string]*string) *ModifyClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterResponse) SetBody(v *ModifyClusterResponseBody) *ModifyClusterResponse {
	s.Body = v
	return s
}

type ModifyClusterConfigurationRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 自定配置。
	CustomizeConfig *ModifyClusterConfigurationRequestCustomizeConfig `json:"customize_config,omitempty" xml:"customize_config,omitempty" type:"Struct"`
}

func (s ModifyClusterConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterConfigurationRequest) SetClusterId(v string) *ModifyClusterConfigurationRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterConfigurationRequest) SetCustomizeConfig(v *ModifyClusterConfigurationRequestCustomizeConfig) *ModifyClusterConfigurationRequest {
	s.CustomizeConfig = v
	return s
}

type ModifyClusterConfigurationRequestCustomizeConfig struct {
	// 配置集合。
	Configs *ModifyClusterConfigurationRequestCustomizeConfigConfigs `json:"configs,omitempty" xml:"configs,omitempty" type:"Struct"`
	// 配置名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ModifyClusterConfigurationRequestCustomizeConfig) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConfigurationRequestCustomizeConfig) GoString() string {
	return s.String()
}

func (s *ModifyClusterConfigurationRequestCustomizeConfig) SetConfigs(v *ModifyClusterConfigurationRequestCustomizeConfigConfigs) *ModifyClusterConfigurationRequestCustomizeConfig {
	s.Configs = v
	return s
}

func (s *ModifyClusterConfigurationRequestCustomizeConfig) SetName(v string) *ModifyClusterConfigurationRequestCustomizeConfig {
	s.Name = &v
	return s
}

type ModifyClusterConfigurationRequestCustomizeConfigConfigs struct {
	// key。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// value。
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ModifyClusterConfigurationRequestCustomizeConfigConfigs) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConfigurationRequestCustomizeConfigConfigs) GoString() string {
	return s.String()
}

func (s *ModifyClusterConfigurationRequestCustomizeConfigConfigs) SetKey(v string) *ModifyClusterConfigurationRequestCustomizeConfigConfigs {
	s.Key = &v
	return s
}

func (s *ModifyClusterConfigurationRequestCustomizeConfigConfigs) SetValue(v string) *ModifyClusterConfigurationRequestCustomizeConfigConfigs {
	s.Value = &v
	return s
}

type ModifyClusterConfigurationResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ModifyClusterConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterConfigurationResponse) SetHeaders(v map[string]*string) *ModifyClusterConfigurationResponse {
	s.Headers = v
	return s
}

type ModifyClusterTagsRequest struct {
	// 汲取ID
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 标签列表。
	Body []*ModifyClusterTagsRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ModifyClusterTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterTagsRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterTagsRequest) SetClusterId(v string) *ModifyClusterTagsRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterTagsRequest) SetBody(v []*ModifyClusterTagsRequestBody) *ModifyClusterTagsRequest {
	s.Body = v
	return s
}

type ModifyClusterTagsRequestBody struct {
	// 标签名。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 标签值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ModifyClusterTagsRequestBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterTagsRequestBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterTagsRequestBody) SetKey(v string) *ModifyClusterTagsRequestBody {
	s.Key = &v
	return s
}

func (s *ModifyClusterTagsRequestBody) SetValue(v string) *ModifyClusterTagsRequestBody {
	s.Value = &v
	return s
}

type ModifyClusterTagsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ModifyClusterTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterTagsResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterTagsResponse) SetHeaders(v map[string]*string) *ModifyClusterTagsResponse {
	s.Headers = v
	return s
}

type PauseComponentUpgradeRequest struct {
	// 集群ID。
	Clusterid *string `json:"clusterid,omitempty" xml:"clusterid,omitempty"`
	// 组件ID。
	Componentid *string `json:"componentid,omitempty" xml:"componentid,omitempty"`
}

func (s PauseComponentUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s PauseComponentUpgradeRequest) GoString() string {
	return s.String()
}

func (s *PauseComponentUpgradeRequest) SetClusterid(v string) *PauseComponentUpgradeRequest {
	s.Clusterid = &v
	return s
}

func (s *PauseComponentUpgradeRequest) SetComponentid(v string) *PauseComponentUpgradeRequest {
	s.Componentid = &v
	return s
}

type PauseComponentUpgradeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s PauseComponentUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s PauseComponentUpgradeResponse) GoString() string {
	return s.String()
}

func (s *PauseComponentUpgradeResponse) SetHeaders(v map[string]*string) *PauseComponentUpgradeResponse {
	s.Headers = v
	return s
}

type RemoveClusterNodesRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 是否同时释放ECS。
	ReleaseNode *bool `json:"release_node,omitempty" xml:"release_node,omitempty"`
	// 是否排空节点上的Pod。
	DrainNode *bool `json:"drain_node,omitempty" xml:"drain_node,omitempty"`
	// 要移除的Node列表。
	Nodes []*string `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
}

func (s RemoveClusterNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *RemoveClusterNodesRequest) SetClusterId(v string) *RemoveClusterNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *RemoveClusterNodesRequest) SetReleaseNode(v bool) *RemoveClusterNodesRequest {
	s.ReleaseNode = &v
	return s
}

func (s *RemoveClusterNodesRequest) SetDrainNode(v bool) *RemoveClusterNodesRequest {
	s.DrainNode = &v
	return s
}

func (s *RemoveClusterNodesRequest) SetNodes(v []*string) *RemoveClusterNodesRequest {
	s.Nodes = v
	return s
}

type RemoveClusterNodesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s RemoveClusterNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveClusterNodesResponse) GoString() string {
	return s.String()
}

func (s *RemoveClusterNodesResponse) SetHeaders(v map[string]*string) *RemoveClusterNodesResponse {
	s.Headers = v
	return s
}

type ResumeComponentUpgradeRequest struct {
	// 集群ID。
	Clusterid *string `json:"clusterid,omitempty" xml:"clusterid,omitempty"`
	// 组件ID。
	Componentid *string `json:"componentid,omitempty" xml:"componentid,omitempty"`
}

func (s ResumeComponentUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeComponentUpgradeRequest) GoString() string {
	return s.String()
}

func (s *ResumeComponentUpgradeRequest) SetClusterid(v string) *ResumeComponentUpgradeRequest {
	s.Clusterid = &v
	return s
}

func (s *ResumeComponentUpgradeRequest) SetComponentid(v string) *ResumeComponentUpgradeRequest {
	s.Componentid = &v
	return s
}

type ResumeComponentUpgradeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ResumeComponentUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeComponentUpgradeResponse) GoString() string {
	return s.String()
}

func (s *ResumeComponentUpgradeResponse) SetHeaders(v map[string]*string) *ResumeComponentUpgradeResponse {
	s.Headers = v
	return s
}

type ScaleClusterRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 扩容节点数。
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// keypair名称，和login_password二选一。
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// SSH登录密码。和keypair二选一。
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// 是否挂载数据盘。
	WorkerDataDisk *bool `json:"worker_data_disk,omitempty" xml:"worker_data_disk,omitempty"`
	// Worker节点ECS规格类型。
	WorkerInstanceTypes []*string `json:"worker_instance_types,omitempty" xml:"worker_instance_types,omitempty" type:"Repeated"`
	// 节点付费类型。
	WorkerInstanceChargeType *string `json:"worker_instance_charge_type,omitempty" xml:"worker_instance_charge_type,omitempty"`
	// 节点包年包月时长。
	WorkerPeriod *int64 `json:"worker_period,omitempty" xml:"worker_period,omitempty"`
	// 当指定为PrePaid的时候需要指定周期。
	WorkerPeriodUnit *string `json:"worker_period_unit,omitempty" xml:"worker_period_unit,omitempty"`
	// 节点是否开启Worker节点自动续费。
	WorkerAutoRenew *bool `json:"worker_auto_renew,omitempty" xml:"worker_auto_renew,omitempty"`
	// 自动续费周期。
	WorkerAutoRenewPeriod *int64 `json:"worker_auto_renew_period,omitempty" xml:"worker_auto_renew_period,omitempty"`
	// 节点系统盘类型。
	WorkerSystemDiskCategory *string `json:"worker_system_disk_category,omitempty" xml:"worker_system_disk_category,omitempty"`
	// 节点系统盘大小
	WorkerSystemDiskSize *int64 `json:"worker_system_disk_size,omitempty" xml:"worker_system_disk_size,omitempty"`
	// 节点是否安装云监控插件。
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// 节点CPU策略。
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// 失败是否回滚。
	DisableRollback *bool `json:"disable_rollback,omitempty" xml:"disable_rollback,omitempty"`
	// 节点交换机ID列表。
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
	// Worker数据盘类型、大小等配置的组合。
	WorkerDataDisks []*ScaleClusterRequestWorkerDataDisks `json:"worker_data_disks,omitempty" xml:"worker_data_disks,omitempty" type:"Repeated"`
	// 集群标签。
	Tags []*ScaleClusterRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// 节点污点标记。
	Taints []*ScaleClusterRequestTaints `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
}

func (s ScaleClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterRequest) GoString() string {
	return s.String()
}

func (s *ScaleClusterRequest) SetClusterId(v string) *ScaleClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ScaleClusterRequest) SetCount(v int64) *ScaleClusterRequest {
	s.Count = &v
	return s
}

func (s *ScaleClusterRequest) SetKeyPair(v string) *ScaleClusterRequest {
	s.KeyPair = &v
	return s
}

func (s *ScaleClusterRequest) SetLoginPassword(v string) *ScaleClusterRequest {
	s.LoginPassword = &v
	return s
}

func (s *ScaleClusterRequest) SetWorkerDataDisk(v bool) *ScaleClusterRequest {
	s.WorkerDataDisk = &v
	return s
}

func (s *ScaleClusterRequest) SetWorkerInstanceTypes(v []*string) *ScaleClusterRequest {
	s.WorkerInstanceTypes = v
	return s
}

func (s *ScaleClusterRequest) SetWorkerInstanceChargeType(v string) *ScaleClusterRequest {
	s.WorkerInstanceChargeType = &v
	return s
}

func (s *ScaleClusterRequest) SetWorkerPeriod(v int64) *ScaleClusterRequest {
	s.WorkerPeriod = &v
	return s
}

func (s *ScaleClusterRequest) SetWorkerPeriodUnit(v string) *ScaleClusterRequest {
	s.WorkerPeriodUnit = &v
	return s
}

func (s *ScaleClusterRequest) SetWorkerAutoRenew(v bool) *ScaleClusterRequest {
	s.WorkerAutoRenew = &v
	return s
}

func (s *ScaleClusterRequest) SetWorkerAutoRenewPeriod(v int64) *ScaleClusterRequest {
	s.WorkerAutoRenewPeriod = &v
	return s
}

func (s *ScaleClusterRequest) SetWorkerSystemDiskCategory(v string) *ScaleClusterRequest {
	s.WorkerSystemDiskCategory = &v
	return s
}

func (s *ScaleClusterRequest) SetWorkerSystemDiskSize(v int64) *ScaleClusterRequest {
	s.WorkerSystemDiskSize = &v
	return s
}

func (s *ScaleClusterRequest) SetCloudMonitorFlags(v bool) *ScaleClusterRequest {
	s.CloudMonitorFlags = &v
	return s
}

func (s *ScaleClusterRequest) SetCpuPolicy(v string) *ScaleClusterRequest {
	s.CpuPolicy = &v
	return s
}

func (s *ScaleClusterRequest) SetDisableRollback(v bool) *ScaleClusterRequest {
	s.DisableRollback = &v
	return s
}

func (s *ScaleClusterRequest) SetVswitchIds(v []*string) *ScaleClusterRequest {
	s.VswitchIds = v
	return s
}

func (s *ScaleClusterRequest) SetWorkerDataDisks(v []*ScaleClusterRequestWorkerDataDisks) *ScaleClusterRequest {
	s.WorkerDataDisks = v
	return s
}

func (s *ScaleClusterRequest) SetTags(v []*ScaleClusterRequestTags) *ScaleClusterRequest {
	s.Tags = v
	return s
}

func (s *ScaleClusterRequest) SetTaints(v []*ScaleClusterRequestTaints) *ScaleClusterRequest {
	s.Taints = v
	return s
}

type ScaleClusterRequestWorkerDataDisks struct {
	// 数据盘类型。
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 是否对数据盘加密。
	Encrypted *string `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// 数据盘大小。
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ScaleClusterRequestWorkerDataDisks) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterRequestWorkerDataDisks) GoString() string {
	return s.String()
}

func (s *ScaleClusterRequestWorkerDataDisks) SetCategory(v string) *ScaleClusterRequestWorkerDataDisks {
	s.Category = &v
	return s
}

func (s *ScaleClusterRequestWorkerDataDisks) SetEncrypted(v string) *ScaleClusterRequestWorkerDataDisks {
	s.Encrypted = &v
	return s
}

func (s *ScaleClusterRequestWorkerDataDisks) SetSize(v string) *ScaleClusterRequestWorkerDataDisks {
	s.Size = &v
	return s
}

type ScaleClusterRequestTags struct {
	// 标签值。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s ScaleClusterRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterRequestTags) GoString() string {
	return s.String()
}

func (s *ScaleClusterRequestTags) SetKey(v string) *ScaleClusterRequestTags {
	s.Key = &v
	return s
}

type ScaleClusterRequestTaints struct {
	// 污点生效策略。
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	// 污点键。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 污点值。
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ScaleClusterRequestTaints) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterRequestTaints) GoString() string {
	return s.String()
}

func (s *ScaleClusterRequestTaints) SetEffect(v string) *ScaleClusterRequestTaints {
	s.Effect = &v
	return s
}

func (s *ScaleClusterRequestTaints) SetKey(v string) *ScaleClusterRequestTaints {
	s.Key = &v
	return s
}

func (s *ScaleClusterRequestTaints) SetValue(v string) *ScaleClusterRequestTaints {
	s.Value = &v
	return s
}

type ScaleClusterShrinkRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 扩容节点数。
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// keypair名称，和login_password二选一。
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// SSH登录密码。和keypair二选一。
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// 是否挂载数据盘。
	WorkerDataDisk *bool `json:"worker_data_disk,omitempty" xml:"worker_data_disk,omitempty"`
	// Worker节点ECS规格类型。
	WorkerInstanceTypes []*string `json:"worker_instance_types,omitempty" xml:"worker_instance_types,omitempty" type:"Repeated"`
	// 节点付费类型。
	WorkerInstanceChargeType *string `json:"worker_instance_charge_type,omitempty" xml:"worker_instance_charge_type,omitempty"`
	// 节点包年包月时长。
	WorkerPeriod *int64 `json:"worker_period,omitempty" xml:"worker_period,omitempty"`
	// 当指定为PrePaid的时候需要指定周期。
	WorkerPeriodUnit *string `json:"worker_period_unit,omitempty" xml:"worker_period_unit,omitempty"`
	// 节点是否开启Worker节点自动续费。
	WorkerAutoRenew *bool `json:"worker_auto_renew,omitempty" xml:"worker_auto_renew,omitempty"`
	// 自动续费周期。
	WorkerAutoRenewPeriod *int64 `json:"worker_auto_renew_period,omitempty" xml:"worker_auto_renew_period,omitempty"`
	// 节点系统盘类型。
	WorkerSystemDiskCategory *string `json:"worker_system_disk_category,omitempty" xml:"worker_system_disk_category,omitempty"`
	// 节点系统盘大小
	WorkerSystemDiskSize *int64 `json:"worker_system_disk_size,omitempty" xml:"worker_system_disk_size,omitempty"`
	// 节点是否安装云监控插件。
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// 节点CPU策略。
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// 失败是否回滚。
	DisableRollback *bool `json:"disable_rollback,omitempty" xml:"disable_rollback,omitempty"`
	// 节点交换机ID列表。
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
	// Worker数据盘类型、大小等配置的组合。
	WorkerDataDisks []*ScaleClusterShrinkRequestWorkerDataDisks `json:"worker_data_disks,omitempty" xml:"worker_data_disks,omitempty" type:"Repeated"`
	// 集群标签。
	Tags []*ScaleClusterShrinkRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// 节点污点标记。
	TaintsShrink *string `json:"taints,omitempty" xml:"taints,omitempty"`
}

func (s ScaleClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *ScaleClusterShrinkRequest) SetClusterId(v string) *ScaleClusterShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ScaleClusterShrinkRequest) SetCount(v int64) *ScaleClusterShrinkRequest {
	s.Count = &v
	return s
}

func (s *ScaleClusterShrinkRequest) SetKeyPair(v string) *ScaleClusterShrinkRequest {
	s.KeyPair = &v
	return s
}

func (s *ScaleClusterShrinkRequest) SetLoginPassword(v string) *ScaleClusterShrinkRequest {
	s.LoginPassword = &v
	return s
}

func (s *ScaleClusterShrinkRequest) SetWorkerDataDisk(v bool) *ScaleClusterShrinkRequest {
	s.WorkerDataDisk = &v
	return s
}

func (s *ScaleClusterShrinkRequest) SetWorkerInstanceTypes(v []*string) *ScaleClusterShrinkRequest {
	s.WorkerInstanceTypes = v
	return s
}

func (s *ScaleClusterShrinkRequest) SetWorkerInstanceChargeType(v string) *ScaleClusterShrinkRequest {
	s.WorkerInstanceChargeType = &v
	return s
}

func (s *ScaleClusterShrinkRequest) SetWorkerPeriod(v int64) *ScaleClusterShrinkRequest {
	s.WorkerPeriod = &v
	return s
}

func (s *ScaleClusterShrinkRequest) SetWorkerPeriodUnit(v string) *ScaleClusterShrinkRequest {
	s.WorkerPeriodUnit = &v
	return s
}

func (s *ScaleClusterShrinkRequest) SetWorkerAutoRenew(v bool) *ScaleClusterShrinkRequest {
	s.WorkerAutoRenew = &v
	return s
}

func (s *ScaleClusterShrinkRequest) SetWorkerAutoRenewPeriod(v int64) *ScaleClusterShrinkRequest {
	s.WorkerAutoRenewPeriod = &v
	return s
}

func (s *ScaleClusterShrinkRequest) SetWorkerSystemDiskCategory(v string) *ScaleClusterShrinkRequest {
	s.WorkerSystemDiskCategory = &v
	return s
}

func (s *ScaleClusterShrinkRequest) SetWorkerSystemDiskSize(v int64) *ScaleClusterShrinkRequest {
	s.WorkerSystemDiskSize = &v
	return s
}

func (s *ScaleClusterShrinkRequest) SetCloudMonitorFlags(v bool) *ScaleClusterShrinkRequest {
	s.CloudMonitorFlags = &v
	return s
}

func (s *ScaleClusterShrinkRequest) SetCpuPolicy(v string) *ScaleClusterShrinkRequest {
	s.CpuPolicy = &v
	return s
}

func (s *ScaleClusterShrinkRequest) SetDisableRollback(v bool) *ScaleClusterShrinkRequest {
	s.DisableRollback = &v
	return s
}

func (s *ScaleClusterShrinkRequest) SetVswitchIds(v []*string) *ScaleClusterShrinkRequest {
	s.VswitchIds = v
	return s
}

func (s *ScaleClusterShrinkRequest) SetWorkerDataDisks(v []*ScaleClusterShrinkRequestWorkerDataDisks) *ScaleClusterShrinkRequest {
	s.WorkerDataDisks = v
	return s
}

func (s *ScaleClusterShrinkRequest) SetTags(v []*ScaleClusterShrinkRequestTags) *ScaleClusterShrinkRequest {
	s.Tags = v
	return s
}

func (s *ScaleClusterShrinkRequest) SetTaintsShrink(v string) *ScaleClusterShrinkRequest {
	s.TaintsShrink = &v
	return s
}

type ScaleClusterShrinkRequestWorkerDataDisks struct {
	// 数据盘类型。
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 是否对数据盘加密。
	Encrypted *string `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// 数据盘大小。
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ScaleClusterShrinkRequestWorkerDataDisks) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterShrinkRequestWorkerDataDisks) GoString() string {
	return s.String()
}

func (s *ScaleClusterShrinkRequestWorkerDataDisks) SetCategory(v string) *ScaleClusterShrinkRequestWorkerDataDisks {
	s.Category = &v
	return s
}

func (s *ScaleClusterShrinkRequestWorkerDataDisks) SetEncrypted(v string) *ScaleClusterShrinkRequestWorkerDataDisks {
	s.Encrypted = &v
	return s
}

func (s *ScaleClusterShrinkRequestWorkerDataDisks) SetSize(v string) *ScaleClusterShrinkRequestWorkerDataDisks {
	s.Size = &v
	return s
}

type ScaleClusterShrinkRequestTags struct {
	// 标签值。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s ScaleClusterShrinkRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterShrinkRequestTags) GoString() string {
	return s.String()
}

func (s *ScaleClusterShrinkRequestTags) SetKey(v string) *ScaleClusterShrinkRequestTags {
	s.Key = &v
	return s
}

type ScaleClusterResponseBody struct {
	// 集群ID。
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// 请求ID。
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 任务ID。
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s ScaleClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleClusterResponseBody) SetClusterId(v string) *ScaleClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ScaleClusterResponseBody) SetRequestId(v string) *ScaleClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleClusterResponseBody) SetTaskId(v string) *ScaleClusterResponseBody {
	s.TaskId = &v
	return s
}

type ScaleClusterResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ScaleClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScaleClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterResponse) GoString() string {
	return s.String()
}

func (s *ScaleClusterResponse) SetHeaders(v map[string]*string) *ScaleClusterResponse {
	s.Headers = v
	return s
}

func (s *ScaleClusterResponse) SetBody(v *ScaleClusterResponseBody) *ScaleClusterResponse {
	s.Body = v
	return s
}

type ScaleOutClusterRequest struct {
	// 扩容目标集群的集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 扩容实例数量。
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// Worker节点付费类型。
	WorkerInstanceChargeType *string `json:"worker_instance_charge_type,omitempty" xml:"worker_instance_charge_type,omitempty"`
	// Worker节点包年包月时长。
	WorkerPeriod *int32 `json:"worker_period,omitempty" xml:"worker_period,omitempty"`
	// Worker节点预付费周期。
	WorkerPeriodUnit *string `json:"worker_period_unit,omitempty" xml:"worker_period_unit,omitempty"`
	// Worker节点是否开启自动续费。
	WorkerAutoRenew *bool `json:"worker_auto_renew,omitempty" xml:"worker_auto_renew,omitempty"`
	// Worker节点自动续费周期。
	WorkerAutoRenewPeriod *int32 `json:"worker_auto_renew_period,omitempty" xml:"worker_auto_renew_period,omitempty"`
	// Worker节点系统盘类型。
	WorkerSystemDiskCategory *string `json:"worker_system_disk_category,omitempty" xml:"worker_system_disk_category,omitempty"`
	// Worker节点系统盘大小。
	WorkerSystemDiskSize *int32 `json:"worker_system_disk_size,omitempty" xml:"worker_system_disk_size,omitempty"`
	// Worker节点是否挂载数据盘。
	WorkerDataDisk *bool `json:"worker_data_disk,omitempty" xml:"worker_data_disk,omitempty"`
	// keypair名称，和login_password二选一。
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// SSH登录密码，和key_pair二选一。
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// 是否安装云监控插件。
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// CPU策略，取值static或者none。
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// 失败是否回滚。
	DisableRollback *bool `json:"disable_rollback,omitempty" xml:"disable_rollback,omitempty"`
	// 自定义镜像ID。
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// 用户自定义数据。
	UserData *string `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// 容器引擎。
	Runtime *ScaleOutClusterRequestRuntime `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Struct"`
	// 节点交换机ID列表，交换机个数取值范围为1~3。
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
	// Worker节点ECS规格类型代码。
	WorkerInstanceTypes []*string `json:"worker_instance_types,omitempty" xml:"worker_instance_types,omitempty" type:"Repeated"`
	// RDS白名单实例列表。
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// Worker数据盘类型、大小等配置的组合。
	WorkerDataDisks []*ScaleOutClusterRequestWorkerDataDisks `json:"worker_data_disks,omitempty" xml:"worker_data_disks,omitempty" type:"Repeated"`
	// 节点标签。
	Tags []*ScaleOutClusterRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// 节点污点信息。
	Taints []*ScaleOutClusterRequestTaints `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
}

func (s ScaleOutClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutClusterRequest) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterRequest) SetClusterId(v string) *ScaleOutClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ScaleOutClusterRequest) SetCount(v int32) *ScaleOutClusterRequest {
	s.Count = &v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerInstanceChargeType(v string) *ScaleOutClusterRequest {
	s.WorkerInstanceChargeType = &v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerPeriod(v int32) *ScaleOutClusterRequest {
	s.WorkerPeriod = &v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerPeriodUnit(v string) *ScaleOutClusterRequest {
	s.WorkerPeriodUnit = &v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerAutoRenew(v bool) *ScaleOutClusterRequest {
	s.WorkerAutoRenew = &v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerAutoRenewPeriod(v int32) *ScaleOutClusterRequest {
	s.WorkerAutoRenewPeriod = &v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerSystemDiskCategory(v string) *ScaleOutClusterRequest {
	s.WorkerSystemDiskCategory = &v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerSystemDiskSize(v int32) *ScaleOutClusterRequest {
	s.WorkerSystemDiskSize = &v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerDataDisk(v bool) *ScaleOutClusterRequest {
	s.WorkerDataDisk = &v
	return s
}

func (s *ScaleOutClusterRequest) SetKeyPair(v string) *ScaleOutClusterRequest {
	s.KeyPair = &v
	return s
}

func (s *ScaleOutClusterRequest) SetLoginPassword(v string) *ScaleOutClusterRequest {
	s.LoginPassword = &v
	return s
}

func (s *ScaleOutClusterRequest) SetCloudMonitorFlags(v bool) *ScaleOutClusterRequest {
	s.CloudMonitorFlags = &v
	return s
}

func (s *ScaleOutClusterRequest) SetCpuPolicy(v string) *ScaleOutClusterRequest {
	s.CpuPolicy = &v
	return s
}

func (s *ScaleOutClusterRequest) SetDisableRollback(v bool) *ScaleOutClusterRequest {
	s.DisableRollback = &v
	return s
}

func (s *ScaleOutClusterRequest) SetImageId(v string) *ScaleOutClusterRequest {
	s.ImageId = &v
	return s
}

func (s *ScaleOutClusterRequest) SetUserData(v string) *ScaleOutClusterRequest {
	s.UserData = &v
	return s
}

func (s *ScaleOutClusterRequest) SetRuntime(v *ScaleOutClusterRequestRuntime) *ScaleOutClusterRequest {
	s.Runtime = v
	return s
}

func (s *ScaleOutClusterRequest) SetVswitchIds(v []*string) *ScaleOutClusterRequest {
	s.VswitchIds = v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerInstanceTypes(v []*string) *ScaleOutClusterRequest {
	s.WorkerInstanceTypes = v
	return s
}

func (s *ScaleOutClusterRequest) SetRdsInstances(v []*string) *ScaleOutClusterRequest {
	s.RdsInstances = v
	return s
}

func (s *ScaleOutClusterRequest) SetWorkerDataDisks(v []*ScaleOutClusterRequestWorkerDataDisks) *ScaleOutClusterRequest {
	s.WorkerDataDisks = v
	return s
}

func (s *ScaleOutClusterRequest) SetTags(v []*ScaleOutClusterRequestTags) *ScaleOutClusterRequest {
	s.Tags = v
	return s
}

func (s *ScaleOutClusterRequest) SetTaints(v []*ScaleOutClusterRequestTaints) *ScaleOutClusterRequest {
	s.Taints = v
	return s
}

type ScaleOutClusterRequestRuntime struct {
	// 容器引擎名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 容器引擎版本。
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ScaleOutClusterRequestRuntime) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutClusterRequestRuntime) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterRequestRuntime) SetName(v string) *ScaleOutClusterRequestRuntime {
	s.Name = &v
	return s
}

func (s *ScaleOutClusterRequestRuntime) SetVersion(v string) *ScaleOutClusterRequestRuntime {
	s.Version = &v
	return s
}

type ScaleOutClusterRequestWorkerDataDisks struct {
	// 数据盘类型。
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 是否对数据盘加密。
	Encrypted *string `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// 数据盘大小。
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ScaleOutClusterRequestWorkerDataDisks) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutClusterRequestWorkerDataDisks) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterRequestWorkerDataDisks) SetCategory(v string) *ScaleOutClusterRequestWorkerDataDisks {
	s.Category = &v
	return s
}

func (s *ScaleOutClusterRequestWorkerDataDisks) SetEncrypted(v string) *ScaleOutClusterRequestWorkerDataDisks {
	s.Encrypted = &v
	return s
}

func (s *ScaleOutClusterRequestWorkerDataDisks) SetSize(v string) *ScaleOutClusterRequestWorkerDataDisks {
	s.Size = &v
	return s
}

type ScaleOutClusterRequestTags struct {
	// 标签名。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 标签值。
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ScaleOutClusterRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutClusterRequestTags) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterRequestTags) SetKey(v string) *ScaleOutClusterRequestTags {
	s.Key = &v
	return s
}

func (s *ScaleOutClusterRequestTags) SetValue(v string) *ScaleOutClusterRequestTags {
	s.Value = &v
	return s
}

type ScaleOutClusterRequestTaints struct {
	// 污点生效策略。
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	// 污点名。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 污点值。
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ScaleOutClusterRequestTaints) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutClusterRequestTaints) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterRequestTaints) SetEffect(v string) *ScaleOutClusterRequestTaints {
	s.Effect = &v
	return s
}

func (s *ScaleOutClusterRequestTaints) SetKey(v string) *ScaleOutClusterRequestTaints {
	s.Key = &v
	return s
}

func (s *ScaleOutClusterRequestTaints) SetValue(v string) *ScaleOutClusterRequestTaints {
	s.Value = &v
	return s
}

type ScaleOutClusterResponseBody struct {
	// 集群ID。
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// 请求ID。
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 任务ID。
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s ScaleOutClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterResponseBody) SetClusterId(v string) *ScaleOutClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ScaleOutClusterResponseBody) SetRequestId(v string) *ScaleOutClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleOutClusterResponseBody) SetTaskId(v string) *ScaleOutClusterResponseBody {
	s.TaskId = &v
	return s
}

type ScaleOutClusterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ScaleOutClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScaleOutClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutClusterResponse) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterResponse) SetHeaders(v map[string]*string) *ScaleOutClusterResponse {
	s.Headers = v
	return s
}

func (s *ScaleOutClusterResponse) SetBody(v *ScaleOutClusterResponseBody) *ScaleOutClusterResponse {
	s.Body = v
	return s
}

type UnInstallClusterAddonsRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 卸载组件列表。
	Addons []*UnInstallClusterAddonsRequestAddons `json:"addons,omitempty" xml:"addons,omitempty" type:"Repeated"`
}

func (s UnInstallClusterAddonsRequest) String() string {
	return tea.Prettify(s)
}

func (s UnInstallClusterAddonsRequest) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsRequest) SetClusterId(v string) *UnInstallClusterAddonsRequest {
	s.ClusterId = &v
	return s
}

func (s *UnInstallClusterAddonsRequest) SetAddons(v []*UnInstallClusterAddonsRequestAddons) *UnInstallClusterAddonsRequest {
	s.Addons = v
	return s
}

type UnInstallClusterAddonsRequestAddons struct {
	// 组件名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UnInstallClusterAddonsRequestAddons) String() string {
	return tea.Prettify(s)
}

func (s UnInstallClusterAddonsRequestAddons) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsRequestAddons) SetName(v string) *UnInstallClusterAddonsRequestAddons {
	s.Name = &v
	return s
}

type UnInstallClusterAddonsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UnInstallClusterAddonsResponse) String() string {
	return tea.Prettify(s)
}

func (s UnInstallClusterAddonsResponse) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsResponse) SetHeaders(v map[string]*string) *UnInstallClusterAddonsResponse {
	s.Headers = v
	return s
}

type UpdateTemplateRequest struct {
	// 部署模板ID。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 部署模板名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 部署模板yaml。
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// 部署模板标签
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// 部署模板描述信息。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 部署模板类型。
	TemplateType *string `json:"template_type,omitempty" xml:"template_type,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) SetTemplateId(v string) *UpdateTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateTemplateRequest) SetName(v string) *UpdateTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplate(v string) *UpdateTemplateRequest {
	s.Template = &v
	return s
}

func (s *UpdateTemplateRequest) SetTags(v string) *UpdateTemplateRequest {
	s.Tags = &v
	return s
}

func (s *UpdateTemplateRequest) SetDescription(v string) *UpdateTemplateRequest {
	s.Description = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateType(v string) *UpdateTemplateRequest {
	s.TemplateType = &v
	return s
}

type UpdateTemplateResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponse) SetHeaders(v map[string]*string) *UpdateTemplateResponse {
	s.Headers = v
	return s
}

type UpgradeClusterRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 组件名称，集群升级时取值"k8s"。
	ComponentName *string `json:"component_name,omitempty" xml:"component_name,omitempty"`
	// 当前版本。
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// 目标版本。
	NextVersion *string `json:"next_version,omitempty" xml:"next_version,omitempty"`
}

func (s UpgradeClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClusterRequest) SetClusterId(v string) *UpgradeClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *UpgradeClusterRequest) SetComponentName(v string) *UpgradeClusterRequest {
	s.ComponentName = &v
	return s
}

func (s *UpgradeClusterRequest) SetVersion(v string) *UpgradeClusterRequest {
	s.Version = &v
	return s
}

func (s *UpgradeClusterRequest) SetNextVersion(v string) *UpgradeClusterRequest {
	s.NextVersion = &v
	return s
}

type UpgradeClusterResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpgradeClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterResponse) GoString() string {
	return s.String()
}

func (s *UpgradeClusterResponse) SetHeaders(v map[string]*string) *UpgradeClusterResponse {
	s.Headers = v
	return s
}

type UpgradeClusterAddonsRequest struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Request body，类型是对象数组。
	Body []*UpgradeClusterAddonsRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s UpgradeClusterAddonsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterAddonsRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClusterAddonsRequest) SetClusterId(v string) *UpgradeClusterAddonsRequest {
	s.ClusterId = &v
	return s
}

func (s *UpgradeClusterAddonsRequest) SetBody(v []*UpgradeClusterAddonsRequestBody) *UpgradeClusterAddonsRequest {
	s.Body = v
	return s
}

type UpgradeClusterAddonsRequestBody struct {
	// 组件名称。
	ComponentName *string `json:"component_name,omitempty" xml:"component_name,omitempty"`
	// 可升级版本。
	NextVersion *string `json:"next_version,omitempty" xml:"next_version,omitempty"`
	// 当前版本。
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpgradeClusterAddonsRequestBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterAddonsRequestBody) GoString() string {
	return s.String()
}

func (s *UpgradeClusterAddonsRequestBody) SetComponentName(v string) *UpgradeClusterAddonsRequestBody {
	s.ComponentName = &v
	return s
}

func (s *UpgradeClusterAddonsRequestBody) SetNextVersion(v string) *UpgradeClusterAddonsRequestBody {
	s.NextVersion = &v
	return s
}

func (s *UpgradeClusterAddonsRequestBody) SetVersion(v string) *UpgradeClusterAddonsRequestBody {
	s.Version = &v
	return s
}

type UpgradeClusterAddonsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpgradeClusterAddonsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterAddonsResponse) GoString() string {
	return s.String()
}

func (s *UpgradeClusterAddonsResponse) SetHeaders(v map[string]*string) *UpgradeClusterAddonsResponse {
	s.Headers = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          tea.String("cs.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("cs.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("cs.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("cs.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("cs.aliyuncs.com"),
		"cn-edge-1":                   tea.String("cs.aliyuncs.com"),
		"cn-fujian":                   tea.String("cs.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("cs.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("cs.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("cs.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("cs.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("cs.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("cs.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("cs.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("cs.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("cs.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("cs.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("cs.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("cs.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("cs.aliyuncs.com"),
		"cn-wuhan":                    tea.String("cs.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("cs.aliyuncs.com"),
		"cn-yushanfang":               tea.String("cs.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("cs.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("cs.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("cs.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("cs.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("cs.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachInstances(ClusterId *string, request *AttachInstancesRequest) (_result *AttachInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AttachInstancesResponse{}
	_body, _err := client.AttachInstancesWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachInstancesWithOptions(ClusterId *string, request *AttachInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AttachInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	body["instances"] = request.Instances
	body["runtime"] = request.Runtime
	body["image_id"] = request.ImageId
	body["format_disk"] = request.FormatDisk
	body["keep_instance_name"] = request.KeepInstanceName
	body["cpu_policy"] = request.CpuPolicy
	body["key_pair"] = request.KeyPair
	body["password"] = request.Password
	body["is_edge_worker"] = request.IsEdgeWorker
	body["user_data"] = request.UserData
	body["nodepool_id"] = request.NodepoolId
	body["rds_instances"] = request.RdsInstances
	body["tags"] = request.Tags
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    body,
	}
	_result = &AttachInstancesResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("AttachInstances"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/attach"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelClusterUpgrade(ClusterId *string, request *CancelClusterUpgradeRequest) (_result *CancelClusterUpgradeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelClusterUpgradeResponse{}
	_body, _err := client.CancelClusterUpgradeWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelClusterUpgradeWithOptions(ClusterId *string, request *CancelClusterUpgradeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelClusterUpgradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &CancelClusterUpgradeResponse{}
	_body, _err := client.DoROARequest(tea.String("CancelClusterUpgrade"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v2/clusters/"+tea.StringValue(ClusterId)+"/upgrade/cancel"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelComponentUpgrade(clusterid *string, componentid *string, request *CancelComponentUpgradeRequest) (_result *CancelComponentUpgradeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelComponentUpgradeResponse{}
	_body, _err := client.CancelComponentUpgradeWithOptions(clusterid, componentid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelComponentUpgradeWithOptions(clusterid *string, componentid *string, request *CancelComponentUpgradeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelComponentUpgradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &CancelComponentUpgradeResponse{}
	_body, _err := client.DoROARequest(tea.String("CancelComponentUpgrade"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterid)+"/components/{componentid}/cancel"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	body["name"] = request.Name
	body["cluster_type"] = request.ClusterType
	body["region_id"] = request.RegionId
	body["zone_id"] = request.ZoneId
	body["kubernetes_version"] = request.KubernetesVersion
	body["deletion_protection"] = request.DeletionProtection
	body["runtime"] = request.Runtime
	body["vpcid"] = request.Vpcid
	body["worker_vswitch_ids"] = request.WorkerVswitchIds
	body["container_cidr"] = request.ContainerCidr
	body["service_cidr"] = request.ServiceCidr
	body["node_cidr_mask"] = request.NodeCidrMask
	body["snat_entry"] = request.SnatEntry
	body["endpoint_public_access"] = request.EndpointPublicAccess
	body["ssh_flags"] = request.SshFlags
	body["rds_instances"] = request.RdsInstances
	body["security_group_id"] = request.SecurityGroupId
	body["is_enterprise_security_group"] = request.IsEnterpriseSecurityGroup
	body["proxy_mode"] = request.ProxyMode
	body["tags"] = request.Tags
	body["images_id"] = request.ImagesId
	body["master_instance_charge_type"] = request.MasterInstanceChargeType
	body["master_period"] = request.MasterPeriod
	body["master_period_unit"] = request.MasterPeriodUnit
	body["master_auto_renew"] = request.MasterAutoRenew
	body["master_auto_renew_period"] = request.MasterAutoRenewPeriod
	body["master_count"] = request.MasterCount
	body["master_vswitch_ids"] = request.MasterVswitchIds
	body["master_instance_types"] = request.MasterInstanceTypes
	body["master_system_disk_category"] = request.MasterSystemDiskCategory
	body["master_system_disk_size"] = request.MasterSystemDiskSize
	body["worker_instance_charge_type"] = request.WorkerInstanceChargeType
	body["worker_period"] = request.WorkerPeriod
	body["worker_period_unit"] = request.WorkerPeriodUnit
	body["worker_auto_renew"] = request.WorkerAutoRenew
	body["worker_auto_renew_period"] = request.WorkerAutoRenewPeriod
	body["num_of_nodes"] = request.NumOfNodes
	body["worker_instance_types"] = request.WorkerInstanceTypes
	body["worker_system_disk_category"] = request.WorkerSystemDiskCategory
	body["worker_system_disk_size"] = request.WorkerSystemDiskSize
	body["worker_data_disks"] = request.WorkerDataDisks
	body["os_type"] = request.OsType
	body["key_pair"] = request.KeyPair
	body["login_password"] = request.LoginPassword
	body["user_data"] = request.UserData
	body["node_port_range"] = request.NodePortRange
	body["cpu_policy"] = request.CpuPolicy
	body["taints"] = request.Taints
	body["cloud_monitor_flags"] = request.CloudMonitorFlags
	body["addons"] = request.Addons
	body["platform"] = request.Platform
	body["vswitch_ids"] = request.VswitchIds
	body["private_zone"] = request.PrivateZone
	body["profile"] = request.Profile
	body["pod_vswitch_ids"] = request.PodVswitchIds
	body["disable_rollback"] = request.DisableRollback
	body["timeout_mins"] = request.TimeoutMins
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    body,
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("CreateCluster"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateKubernetesTrigger(request *CreateKubernetesTriggerRequest) (_result *CreateKubernetesTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateKubernetesTriggerResponse{}
	_body, _err := client.CreateKubernetesTriggerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateKubernetesTriggerWithOptions(request *CreateKubernetesTriggerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateKubernetesTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	body["RegionId"] = request.RegionId
	body["ClusterId"] = request.ClusterId
	body["ProjectId"] = request.ProjectId
	body["Type"] = request.Type
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    body,
	}
	_result = &CreateKubernetesTriggerResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("CreateKubernetesTrigger"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/triggers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCluster(ClusterId *string, request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterWithOptions(ClusterId *string, request *DeleteClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["retain_resources"] = request.RetainResources
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteCluster"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteKubernetesTrigger(Id *string, request *DeleteKubernetesTriggerRequest) (_result *DeleteKubernetesTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteKubernetesTriggerResponse{}
	_body, _err := client.DeleteKubernetesTriggerWithOptions(Id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteKubernetesTriggerWithOptions(Id *string, request *DeleteKubernetesTriggerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteKubernetesTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteKubernetesTriggerResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteKubernetesTrigger"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/triggers/revoke/"+tea.StringValue(Id)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAddons(request *DescribeAddonsRequest) (_result *DescribeAddonsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAddonsResponse{}
	_body, _err := client.DescribeAddonsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAddonsWithOptions(request *DescribeAddonsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAddonsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["region"] = request.Region
	query["cluster_type"] = request.ClusterType
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeAddonsResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeAddons"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/components/metadata"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterAddonUpgradeStatus(ClusterId *string, ComponentId *string, request *DescribeClusterAddonUpgradeStatusRequest) (_result *DescribeClusterAddonUpgradeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterAddonUpgradeStatusResponse{}
	_body, _err := client.DescribeClusterAddonUpgradeStatusWithOptions(ClusterId, ComponentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterAddonUpgradeStatusWithOptions(ClusterId *string, ComponentId *string, request *DescribeClusterAddonUpgradeStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeClusterAddonUpgradeStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DescribeClusterAddonUpgradeStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeClusterAddonUpgradeStatus"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/components/{ComponentId}/upgradestatus"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 批量查询集群Addon升级状态

 */
func (client *Client) DescribeClusterAddonsUpgradeStatus(ClusterId *string, request *DescribeClusterAddonsUpgradeStatusRequest) (_result *DescribeClusterAddonsUpgradeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterAddonsUpgradeStatusResponse{}
	_body, _err := client.DescribeClusterAddonsUpgradeStatusWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterAddonsUpgradeStatusWithOptions(ClusterId *string, request *DescribeClusterAddonsUpgradeStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeClusterAddonsUpgradeStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["componentIds"] = request.ComponentIds
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeClusterAddonsUpgradeStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeClusterAddonsUpgradeStatus"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/[ClusterId]/components/upgradestatus"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterAddonsVersion(ClusterId *string, request *DescribeClusterAddonsVersionRequest) (_result *DescribeClusterAddonsVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterAddonsVersionResponse{}
	_body, _err := client.DescribeClusterAddonsVersionWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterAddonsVersionWithOptions(ClusterId *string, request *DescribeClusterAddonsVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeClusterAddonsVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DescribeClusterAddonsVersionResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeClusterAddonsVersion"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/components/version"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterAttachScripts(ClusterId *string, request *DescribeClusterAttachScriptsRequest) (_result *DescribeClusterAttachScriptsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterAttachScriptsResponse{}
	_body, _err := client.DescribeClusterAttachScriptsWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterAttachScriptsWithOptions(ClusterId *string, request *DescribeClusterAttachScriptsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeClusterAttachScriptsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	body["arch"] = request.Arch
	body["options"] = request.Options
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    body,
	}
	_result = &DescribeClusterAttachScriptsResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("DescribeClusterAttachScripts"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/attachscript"), tea.String("string"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterDetail(ClusterId *string, request *DescribeClusterDetailRequest) (_result *DescribeClusterDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterDetailResponse{}
	_body, _err := client.DescribeClusterDetailWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterDetailWithOptions(ClusterId *string, request *DescribeClusterDetailRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeClusterDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DescribeClusterDetailResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeClusterDetail"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterLogs(ClusterId *string, request *DescribeClusterLogsRequest) (_result *DescribeClusterLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterLogsResponse{}
	_body, _err := client.DescribeClusterLogsWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterLogsWithOptions(ClusterId *string, request *DescribeClusterLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeClusterLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DescribeClusterLogsResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeClusterLogs"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/logs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterNodes(ClusterId *string, request *DescribeClusterNodesRequest) (_result *DescribeClusterNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterNodesResponse{}
	_body, _err := client.DescribeClusterNodesWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterNodesWithOptions(ClusterId *string, request *DescribeClusterNodesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeClusterNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["pageSize"] = request.PageSize
	query["pageNumber"] = request.PageNumber
	query["nodepool_id"] = request.NodepoolId
	query["state"] = request.State
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeClusterNodesResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeClusterNodes"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/nodes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterResources(ClusterId *string, request *DescribeClusterResourcesRequest) (_result *DescribeClusterResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterResourcesResponse{}
	_body, _err := client.DescribeClusterResourcesWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterResourcesWithOptions(ClusterId *string, request *DescribeClusterResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeClusterResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DescribeClusterResourcesResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeClusterResources"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/resources"), tea.String("array"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterUserKubeconfig(ClusterId *string, request *DescribeClusterUserKubeconfigRequest) (_result *DescribeClusterUserKubeconfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterUserKubeconfigResponse{}
	_body, _err := client.DescribeClusterUserKubeconfigWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterUserKubeconfigWithOptions(ClusterId *string, request *DescribeClusterUserKubeconfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeClusterUserKubeconfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["PrivateIpAddress"] = request.PrivateIpAddress
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeClusterUserKubeconfigResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeClusterUserKubeconfig"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/k8s/"+tea.StringValue(ClusterId)+"/user_config"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterV2UserKubeconfig(ClusterId *string, request *DescribeClusterV2UserKubeconfigRequest) (_result *DescribeClusterV2UserKubeconfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClusterV2UserKubeconfigResponse{}
	_body, _err := client.DescribeClusterV2UserKubeconfigWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterV2UserKubeconfigWithOptions(ClusterId *string, request *DescribeClusterV2UserKubeconfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeClusterV2UserKubeconfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["PrivateIpAddress"] = request.PrivateIpAddress
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeClusterV2UserKubeconfigResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeClusterV2UserKubeconfig"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v2/k8s/"+tea.StringValue(ClusterId)+"/user_config"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusters(request *DescribeClustersRequest) (_result *DescribeClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClustersResponse{}
	_body, _err := client.DescribeClustersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClustersWithOptions(request *DescribeClustersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["name"] = request.Name
	query["clusterType"] = request.ClusterType
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeClustersResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeClusters"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters"), tea.String("array"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClustersV1(request *DescribeClustersV1Request) (_result *DescribeClustersV1Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeClustersV1Response{}
	_body, _err := client.DescribeClustersV1WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClustersV1WithOptions(request *DescribeClustersV1Request, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeClustersV1Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Name"] = request.Name
	query["ClusterType"] = request.ClusterType
	query["page_size"] = request.PageSize
	query["page_number"] = request.PageNumber
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeClustersV1Response{}
	_body, _err := client.DoROARequest(tea.String("DescribeClustersV1"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/clusters"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExternalAgent(ClusterId *string, request *DescribeExternalAgentRequest) (_result *DescribeExternalAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeExternalAgentResponse{}
	_body, _err := client.DescribeExternalAgentWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExternalAgentWithOptions(ClusterId *string, request *DescribeExternalAgentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeExternalAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DescribeExternalAgentResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeExternalAgent"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/k8s/"+tea.StringValue(ClusterId)+"/external/agent/deployment"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTemplates(request *DescribeTemplatesRequest) (_result *DescribeTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTemplatesResponse{}
	_body, _err := client.DescribeTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTemplatesWithOptions(request *DescribeTemplatesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["template_type"] = request.TemplateType
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeTemplatesResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeTemplates"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/templates"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserQuota() (_result *DescribeUserQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeUserQuotaResponse{}
	_body, _err := client.DescribeUserQuotaWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserQuotaWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeUserQuotaResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DescribeUserQuotaResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeUserQuota"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/quota"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetKubernetesTrigger(ClusterId *string, request *GetKubernetesTriggerRequest) (_result *GetKubernetesTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetKubernetesTriggerResponse{}
	_body, _err := client.GetKubernetesTriggerWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetKubernetesTriggerWithOptions(ClusterId *string, request *GetKubernetesTriggerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetKubernetesTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Namespace"] = request.Namespace
	query["Type"] = request.Type
	query["Name"] = request.Name
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetKubernetesTriggerResponse{}
	_body, _err := client.DoROARequest(tea.String("GetKubernetesTrigger"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/triggers/"+tea.StringValue(ClusterId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUpgradeStatus(ClusterId *string, request *GetUpgradeStatusRequest) (_result *GetUpgradeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUpgradeStatusResponse{}
	_body, _err := client.GetUpgradeStatusWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUpgradeStatusWithOptions(ClusterId *string, request *GetUpgradeStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUpgradeStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetUpgradeStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUpgradeStatus"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v2/clusters/"+tea.StringValue(ClusterId)+"/upgrade/status"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallClusterAddons(ClusterId *string, request *InstallClusterAddonsRequest) (_result *InstallClusterAddonsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallClusterAddonsResponse{}
	_body, _err := client.InstallClusterAddonsWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallClusterAddonsWithOptions(ClusterId *string, request *InstallClusterAddonsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InstallClusterAddonsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    util.ToArray(request.Body),
	}
	_result = &InstallClusterAddonsResponse{}
	_body, _err := client.DoROARequest(tea.String("InstallClusterAddons"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/components/install"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCluster(ClusterId *string, request *ModifyClusterRequest) (_result *ModifyClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyClusterResponse{}
	_body, _err := client.ModifyClusterWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterWithOptions(ClusterId *string, request *ModifyClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	body["deletion_protection"] = request.DeletionProtection
	body["ingress_loadbalancer_id"] = request.IngressLoadbalancerId
	body["api_server_eip"] = request.ApiServerEip
	body["api_server_eip_id"] = request.ApiServerEipId
	body["resource_group_id"] = request.ResourceGroupId
	body["ingress_domain_rebinding"] = request.IngressDomainRebinding
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    body,
	}
	_result = &ModifyClusterResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("ModifyCluster"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v2/clusters/"+tea.StringValue(ClusterId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterConfiguration(ClusterId *string, request *ModifyClusterConfigurationRequest) (_result *ModifyClusterConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyClusterConfigurationResponse{}
	_body, _err := client.ModifyClusterConfigurationWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterConfigurationWithOptions(ClusterId *string, request *ModifyClusterConfigurationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyClusterConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	body["customize_config"] = request.CustomizeConfig
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    body,
	}
	_result = &ModifyClusterConfigurationResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("ModifyClusterConfiguration"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/configuration"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterTags(ClusterId *string, request *ModifyClusterTagsRequest) (_result *ModifyClusterTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyClusterTagsResponse{}
	_body, _err := client.ModifyClusterTagsWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterTagsWithOptions(ClusterId *string, request *ModifyClusterTagsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyClusterTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    util.ToArray(request.Body),
	}
	_result = &ModifyClusterTagsResponse{}
	_body, _err := client.DoROARequest(tea.String("ModifyClusterTags"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/tags"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PauseComponentUpgrade(clusterid *string, componentid *string, request *PauseComponentUpgradeRequest) (_result *PauseComponentUpgradeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PauseComponentUpgradeResponse{}
	_body, _err := client.PauseComponentUpgradeWithOptions(clusterid, componentid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PauseComponentUpgradeWithOptions(clusterid *string, componentid *string, request *PauseComponentUpgradeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PauseComponentUpgradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &PauseComponentUpgradeResponse{}
	_body, _err := client.DoROARequest(tea.String("PauseComponentUpgrade"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterid)+"/components/{componentid}/pause"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveClusterNodes(ClusterId *string, request *RemoveClusterNodesRequest) (_result *RemoveClusterNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveClusterNodesResponse{}
	_body, _err := client.RemoveClusterNodesWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveClusterNodesWithOptions(ClusterId *string, request *RemoveClusterNodesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveClusterNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	body["release_node"] = request.ReleaseNode
	body["drain_node"] = request.DrainNode
	body["nodes"] = request.Nodes
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    body,
	}
	_result = &RemoveClusterNodesResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("RemoveClusterNodes"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v2/clusters/"+tea.StringValue(ClusterId)+"/nodes/remove"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeComponentUpgrade(clusterid *string, componentid *string, request *ResumeComponentUpgradeRequest) (_result *ResumeComponentUpgradeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeComponentUpgradeResponse{}
	_body, _err := client.ResumeComponentUpgradeWithOptions(clusterid, componentid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeComponentUpgradeWithOptions(clusterid *string, componentid *string, request *ResumeComponentUpgradeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ResumeComponentUpgradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ResumeComponentUpgradeResponse{}
	_body, _err := client.DoROARequest(tea.String("ResumeComponentUpgrade"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterid)+"/components/{componentid}/resume"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScaleCluster(ClusterId *string, request *ScaleClusterRequest) (_result *ScaleClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScaleClusterResponse{}
	_body, _err := client.ScaleClusterWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScaleClusterWithOptions(ClusterId *string, tmpReq *ScaleClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ScaleClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ScaleClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Taints)) {
		request.TaintsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Taints, tea.String("taints"), tea.String("json"))
	}

	body := map[string]interface{}{}
	body["count"] = request.Count
	body["key_pair"] = request.KeyPair
	body["login_password"] = request.LoginPassword
	body["worker_data_disk"] = request.WorkerDataDisk
	body["worker_instance_types"] = request.WorkerInstanceTypes
	body["worker_instance_charge_type"] = request.WorkerInstanceChargeType
	body["worker_period"] = request.WorkerPeriod
	body["worker_period_unit"] = request.WorkerPeriodUnit
	body["worker_auto_renew"] = request.WorkerAutoRenew
	body["worker_auto_renew_period"] = request.WorkerAutoRenewPeriod
	body["worker_system_disk_category"] = request.WorkerSystemDiskCategory
	body["worker_system_disk_size"] = request.WorkerSystemDiskSize
	body["cloud_monitor_flags"] = request.CloudMonitorFlags
	body["cpu_policy"] = request.CpuPolicy
	body["disable_rollback"] = request.DisableRollback
	body["vswitch_ids"] = request.VswitchIds
	body["worker_data_disks"] = request.WorkerDataDisks
	body["tags"] = request.Tags
	body["taints"] = request.TaintsShrink
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    body,
	}
	_result = &ScaleClusterResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("ScaleCluster"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScaleOutCluster(ClusterId *string, request *ScaleOutClusterRequest) (_result *ScaleOutClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScaleOutClusterResponse{}
	_body, _err := client.ScaleOutClusterWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScaleOutClusterWithOptions(ClusterId *string, request *ScaleOutClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ScaleOutClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	body["count"] = request.Count
	body["worker_instance_charge_type"] = request.WorkerInstanceChargeType
	body["worker_period"] = request.WorkerPeriod
	body["worker_period_unit"] = request.WorkerPeriodUnit
	body["worker_auto_renew"] = request.WorkerAutoRenew
	body["worker_auto_renew_period"] = request.WorkerAutoRenewPeriod
	body["worker_system_disk_category"] = request.WorkerSystemDiskCategory
	body["worker_system_disk_size"] = request.WorkerSystemDiskSize
	body["worker_data_disk"] = request.WorkerDataDisk
	body["key_pair"] = request.KeyPair
	body["login_password"] = request.LoginPassword
	body["cloud_monitor_flags"] = request.CloudMonitorFlags
	body["cpu_policy"] = request.CpuPolicy
	body["disable_rollback"] = request.DisableRollback
	body["image_id"] = request.ImageId
	body["user_data"] = request.UserData
	body["runtime"] = request.Runtime
	body["vswitch_ids"] = request.VswitchIds
	body["worker_instance_types"] = request.WorkerInstanceTypes
	body["rds_instances"] = request.RdsInstances
	body["worker_data_disks"] = request.WorkerDataDisks
	body["tags"] = request.Tags
	body["taints"] = request.Taints
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    body,
	}
	_result = &ScaleOutClusterResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("ScaleOutCluster"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v2/clusters/"+tea.StringValue(ClusterId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnInstallClusterAddons(ClusterId *string, request *UnInstallClusterAddonsRequest) (_result *UnInstallClusterAddonsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnInstallClusterAddonsResponse{}
	_body, _err := client.UnInstallClusterAddonsWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnInstallClusterAddonsWithOptions(ClusterId *string, request *UnInstallClusterAddonsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UnInstallClusterAddonsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	body["addons"] = request.Addons
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    body,
	}
	_result = &UnInstallClusterAddonsResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("UnInstallClusterAddons"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/components/uninstall"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTemplate(TemplateId *string, request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(TemplateId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTemplateWithOptions(TemplateId *string, request *UpdateTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	body["name"] = request.Name
	body["template"] = request.Template
	body["tags"] = request.Tags
	body["description"] = request.Description
	body["template_type"] = request.TemplateType
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    body,
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("UpdateTemplate"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/templates/"+tea.StringValue(TemplateId)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeCluster(ClusterId *string, request *UpgradeClusterRequest) (_result *UpgradeClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeClusterResponse{}
	_body, _err := client.UpgradeClusterWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeClusterWithOptions(ClusterId *string, request *UpgradeClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpgradeClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	body["component_name"] = request.ComponentName
	body["version"] = request.Version
	body["next_version"] = request.NextVersion
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    body,
	}
	_result = &UpgradeClusterResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("UpgradeCluster"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v2/clusters/"+tea.StringValue(ClusterId)+"/upgrade"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeClusterAddons(ClusterId *string, request *UpgradeClusterAddonsRequest) (_result *UpgradeClusterAddonsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeClusterAddonsResponse{}
	_body, _err := client.UpgradeClusterAddonsWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeClusterAddonsWithOptions(ClusterId *string, request *UpgradeClusterAddonsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpgradeClusterAddonsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    util.ToArray(request.Body),
	}
	_result = &UpgradeClusterAddonsResponse{}
	_body, _err := client.DoROARequest(tea.String("UpgradeClusterAddons"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/components/upgrade"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
