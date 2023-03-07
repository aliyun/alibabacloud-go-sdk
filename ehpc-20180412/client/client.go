// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddContainerAppRequest struct {
	// The type of the container. Set the value to singularity.
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
	// The description of the container.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tags of the image.
	//
	// The repository stores a type of images such as Ubuntu images. Tags are used to identify the images. Examples: 16.04, 17.04, and latest.
	//
	// Default value: latest
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The name of the container. The name must be 2 to 64 characters in length. It must start with a letter and can contain letters, digits, hyphens (-), and underscores (\_).
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the repository. The image that has the same name as the repository is pulled.
	//
	// For information about image names, visit [Docker Hub official website](https://hub.docker.com/search?q=\&type=image).
	Repository *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
}

func (s AddContainerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s AddContainerAppRequest) GoString() string {
	return s.String()
}

func (s *AddContainerAppRequest) SetContainerType(v string) *AddContainerAppRequest {
	s.ContainerType = &v
	return s
}

func (s *AddContainerAppRequest) SetDescription(v string) *AddContainerAppRequest {
	s.Description = &v
	return s
}

func (s *AddContainerAppRequest) SetImageTag(v string) *AddContainerAppRequest {
	s.ImageTag = &v
	return s
}

func (s *AddContainerAppRequest) SetName(v string) *AddContainerAppRequest {
	s.Name = &v
	return s
}

func (s *AddContainerAppRequest) SetRepository(v string) *AddContainerAppRequest {
	s.Repository = &v
	return s
}

type AddContainerAppResponseBody struct {
	// The ID of the container.
	ContainerId *AddContainerAppResponseBodyContainerId `json:"ContainerId,omitempty" xml:"ContainerId,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddContainerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddContainerAppResponseBody) GoString() string {
	return s.String()
}

func (s *AddContainerAppResponseBody) SetContainerId(v *AddContainerAppResponseBodyContainerId) *AddContainerAppResponseBody {
	s.ContainerId = v
	return s
}

func (s *AddContainerAppResponseBody) SetRequestId(v string) *AddContainerAppResponseBody {
	s.RequestId = &v
	return s
}

type AddContainerAppResponseBodyContainerId struct {
	ContainerId []*string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty" type:"Repeated"`
}

func (s AddContainerAppResponseBodyContainerId) String() string {
	return tea.Prettify(s)
}

func (s AddContainerAppResponseBodyContainerId) GoString() string {
	return s.String()
}

func (s *AddContainerAppResponseBodyContainerId) SetContainerId(v []*string) *AddContainerAppResponseBodyContainerId {
	s.ContainerId = v
	return s
}

type AddContainerAppResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddContainerAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddContainerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s AddContainerAppResponse) GoString() string {
	return s.String()
}

func (s *AddContainerAppResponse) SetHeaders(v map[string]*string) *AddContainerAppResponse {
	s.Headers = v
	return s
}

func (s *AddContainerAppResponse) SetStatusCode(v int32) *AddContainerAppResponse {
	s.StatusCode = &v
	return s
}

func (s *AddContainerAppResponse) SetBody(v *AddContainerAppResponseBody) *AddContainerAppResponse {
	s.Body = v
	return s
}

type AddExistedNodesRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the image that is specified for the compute nodes. The image must meet the following requirements:
	//
	// *   The operating system that is specified by the image must be the same as that of the existing cluster nodes. For example, if the operating system of the cluster nodes is CentOS, you can select only a CentOS image.
	//
	// > If you add nodes to a hybrid cloud cluster that supports multiple operating systems, you can select a Windows Server image or a CentOS image when the operating system of the cluster nodes is Windows.
	//
	// *   The major version of the image specified for the compute nodes that you want to add is the same as that of the image of the cluster. For example, if the version of the cluster image is CentOS 7.x, the version of the image specified for the compute nodes must be CentOS 7.x.
	//
	// You can call the [ListImages](~~87213~~) and [ListCustomImages](~~87215~~) operations to query the image ID.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The type of the image. Valid values:
	//
	// *   system: public image
	// *   self: custom image
	// *   others: shared image
	// *   marketplace: Alibaba Cloud Marketplace image
	//
	// Default value: system
	ImageOwnerAlias *string                           `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	Instance        []*AddExistedNodesRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
	// The queue in the cluster to which the node is to be added.
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
}

func (s AddExistedNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddExistedNodesRequest) GoString() string {
	return s.String()
}

func (s *AddExistedNodesRequest) SetClusterId(v string) *AddExistedNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *AddExistedNodesRequest) SetImageId(v string) *AddExistedNodesRequest {
	s.ImageId = &v
	return s
}

func (s *AddExistedNodesRequest) SetImageOwnerAlias(v string) *AddExistedNodesRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *AddExistedNodesRequest) SetInstance(v []*AddExistedNodesRequestInstance) *AddExistedNodesRequest {
	s.Instance = v
	return s
}

func (s *AddExistedNodesRequest) SetJobQueue(v string) *AddExistedNodesRequest {
	s.JobQueue = &v
	return s
}

type AddExistedNodesRequestInstance struct {
	// The Nth node ID. N starts from 1. Valid values: 1 to 100.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AddExistedNodesRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s AddExistedNodesRequestInstance) GoString() string {
	return s.String()
}

func (s *AddExistedNodesRequestInstance) SetId(v string) *AddExistedNodesRequestInstance {
	s.Id = &v
	return s
}

type AddExistedNodesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddExistedNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddExistedNodesResponseBody) GoString() string {
	return s.String()
}

func (s *AddExistedNodesResponseBody) SetRequestId(v string) *AddExistedNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddExistedNodesResponseBody) SetTaskId(v string) *AddExistedNodesResponseBody {
	s.TaskId = &v
	return s
}

type AddExistedNodesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddExistedNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddExistedNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddExistedNodesResponse) GoString() string {
	return s.String()
}

func (s *AddExistedNodesResponse) SetHeaders(v map[string]*string) *AddExistedNodesResponse {
	s.Headers = v
	return s
}

func (s *AddExistedNodesResponse) SetStatusCode(v int32) *AddExistedNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddExistedNodesResponse) SetBody(v *AddExistedNodesResponseBody) *AddExistedNodesResponse {
	s.Body = v
	return s
}

type AddLocalNodesRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information of the local node. A JSON string that contains the HostName, IpAddress, CpuCores, and Memory (Unit: MB) of the local node.
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// The queue to which to add the local node.
	Queue *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
}

func (s AddLocalNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLocalNodesRequest) GoString() string {
	return s.String()
}

func (s *AddLocalNodesRequest) SetClusterId(v string) *AddLocalNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *AddLocalNodesRequest) SetNodes(v string) *AddLocalNodesRequest {
	s.Nodes = &v
	return s
}

func (s *AddLocalNodesRequest) SetQueue(v string) *AddLocalNodesRequest {
	s.Queue = &v
	return s
}

type AddLocalNodesResponseBody struct {
	// The local nodes in the cluster.
	InstanceIds *AddLocalNodesResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLocalNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddLocalNodesResponseBody) GoString() string {
	return s.String()
}

func (s *AddLocalNodesResponseBody) SetInstanceIds(v *AddLocalNodesResponseBodyInstanceIds) *AddLocalNodesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *AddLocalNodesResponseBody) SetRequestId(v string) *AddLocalNodesResponseBody {
	s.RequestId = &v
	return s
}

type AddLocalNodesResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s AddLocalNodesResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s AddLocalNodesResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *AddLocalNodesResponseBodyInstanceIds) SetInstanceId(v []*string) *AddLocalNodesResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

type AddLocalNodesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddLocalNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddLocalNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLocalNodesResponse) GoString() string {
	return s.String()
}

func (s *AddLocalNodesResponse) SetHeaders(v map[string]*string) *AddLocalNodesResponse {
	s.Headers = v
	return s
}

func (s *AddLocalNodesResponse) SetStatusCode(v int32) *AddLocalNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLocalNodesResponse) SetBody(v *AddLocalNodesResponseBody) *AddLocalNodesResponse {
	s.Body = v
	return s
}

type AddNodesRequest struct {
	// Specifies whether to allocate a public IP address to the compute nodes. Valid values:
	//
	// *   true: A public IP address is allocated to the compute nodes.
	// *   false: A public IP address is not allocated to the compute nodes.
	//
	// Default value: false
	AllocatePublicAddress *bool `json:"AllocatePublicAddress,omitempty" xml:"AllocatePublicAddress,omitempty"`
	// Specifies whether to enable auto-renewal. The parameter takes effect only when EcsChargeType is set to PrePaid. Valid values:
	//
	// *   true: enables auto-renewal
	// *   false: disables auto-renewal
	//
	// Default value: true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period of the subscription compute nodes. The parameter takes effect when AutoRenew is set to true.
	//
	// *   If PeriodUnit is set to Week, the valid values of the AutoRenewPeriod parameter are 1, 2, and 3.
	// *   If PeriodUnit is set to Month, the valid values of the AutoRenewPeriod parameter are 1, 2, 3, 6, and 12.
	//
	// Default value: 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence of a request?](~~25693~~)
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the E-HPC cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether the compute nodes support hyper-threading. Valid values:
	//
	// *   true: Hyper-threading is supported.
	// *   false: Hyper-threading is not supported.
	//
	// Default value: true
	ComputeEnableHt *bool `json:"ComputeEnableHt,omitempty" xml:"ComputeEnableHt,omitempty"`
	// The protection period of the preemptible instance. Unit: hours. Valid values: 0 and 1. A value of 0 indicates that the preemptible instance has no protection period.
	ComputeSpotDuration *int32 `json:"ComputeSpotDuration,omitempty" xml:"ComputeSpotDuration,omitempty"`
	// The interruption mode of the preemptible instance. Default value: Terminate. Set the value to Terminate, which indicates that the instance is released.
	ComputeSpotInterruptionBehavior *string `json:"ComputeSpotInterruptionBehavior,omitempty" xml:"ComputeSpotInterruptionBehavior,omitempty"`
	// The maximum hourly price of the compute nodes. The value can be accurate to three decimal places. The parameter only takes effect when SpotStrategy is set to SpotWithPriceLimit.
	ComputeSpotPriceLimit *string `json:"ComputeSpotPriceLimit,omitempty" xml:"ComputeSpotPriceLimit,omitempty"`
	// The preemption policy of the compute nodes. The parameter only takes effect when EcsChargeType is set to PostPaid. Valid values:
	//
	// *   NoSpot: The compute nodes are pay-as-you-go instances.
	// *   SpotWithPriceLimit: The instance is a preemptible instance that has a user-defined maximum hourly price.
	// *   SpotAsPriceGo: The compute nodes are preemptible instances for which the market price at the time of purchase is used as the bid price.
	//
	// Default value: NoSpot
	ComputeSpotStrategy *string `json:"ComputeSpotStrategy,omitempty" xml:"ComputeSpotStrategy,omitempty"`
	// The number of compute nodes that you want to add. Valid values: 1 to 99. The value of this parameter is greater than that of the MinCount parameter.
	//
	// *   If the number of available ECS instances is less than the value of the MinCount parameter, the compute nodes cannot be added.
	// *   If the number of available ECS instances is greater than the value of the MinCount parameter and less than that of the Count parameter, the compute nodes are added based on the value of the MinCount parameter.
	// *   If the number of available ECS instances is greater than the value of the Count parameter, the compute nodes are added based on the value of the Count parameter.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The mode in which the compute nodes are added. Valid values:
	//
	// *   manual: The compute nodes are manually added.
	// *   autoscale: The compute nodes are automatically added.
	//
	// Default value: manual
	CreateMode *string                     `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	DataDisks  []*AddNodesRequestDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// The billing method of the compute nodes. Valid values:
	//
	// *   PostPaid: pay-as-you-go
	// *   PrePaid: subscription
	//
	// Default value: PostPaid
	//
	// If the parameter is set to PrePaid, auto-renewal is enabled by default. After the E-HPC cluster is released, auto-renewal is disabled.
	EcsChargeType *string `json:"EcsChargeType,omitempty" xml:"EcsChargeType,omitempty"`
	// The prefix of the hostname. You can specify the parameter to manage the compute nodes in an efficient manner.
	HostNamePrefix *string `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	// The suffix of the hostname. You can specify the parameter to manage the compute nodes in an efficient manner.
	HostNameSuffix *string `json:"HostNameSuffix,omitempty" xml:"HostNameSuffix,omitempty"`
	// The ID of the image that is specified for the compute nodes. The image must meet the following requirements:
	//
	// *   The operating system that is specified by the image must be the same as that of the existing cluster nodes. For example, if the operating system of the cluster nodes is CentOS, you can select only a CentOS image.
	//
	// > If you add nodes to a hybrid cloud cluster that supports multiple operating systems, you can select a Windows Server image or a CentOS image when the operating system of the cluster nodes is Windows.
	//
	// *   The major version of the image specified for the compute nodes that you want to add is the same as that of the image of the cluster. For example, if the version of the cluster image is CentOS 7.x, the version of the image specified for the compute nodes must be CentOS 7.x.
	//
	// You can call the [ListImages](~~87213~~) and [ListCustomImages](~~87215~~) operations to query the image ID.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The type of the image. Valid values:
	//
	// *   system: public image
	// *   self: custom image
	// *   others: shared image
	// *   marketplace: Alibaba Cloud Marketplace image
	//
	// Default value: system
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The instance type of the compute nodes. The default value is the instance type that was specified when you created the E-HPC cluster or the last time when you added compute nodes.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The billing method of the elastic IP address (EIP). Valid values:
	//
	// *   PayByBandwidth: pay-by-bandwidth
	// *   PayByTraffic: pay-by-traffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth. Unit: Mbit/s. Valid values:
	//
	// *   If the purchased outbound public bandwidth is less than or equal to 10 Mbit/s, the valid values of the parameter are 1 to 10 and the default value is 10.
	// *   If the purchased outbound public bandwidth is greater than 10 Mbit/s, the valid values of this parameter are 1 to the amount of the outbound bandwidth that is purchased.
	InternetMaxBandWidthIn *int32 `json:"InternetMaxBandWidthIn,omitempty" xml:"InternetMaxBandWidthIn,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s. Valid values: 0 to 100.
	//
	// Default value: 0
	InternetMaxBandWidthOut *int32 `json:"InternetMaxBandWidthOut,omitempty" xml:"InternetMaxBandWidthOut,omitempty"`
	// The queue to which the compute nodes are added.
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	// The minimum number of the compute nodes that you want to add. Valid values: 1 to 99. The value of the parameter is less than that of the Count parameter.
	//
	// *   If the number of available ECS instances is less than the value of the MinCount parameter, the compute nodes cannot be added.
	// *   If the number of available ECS instances is greater than the value of the MinCount parameter and less than that of the Count parameter, the compute nodes are added based on the value of the MinCount parameter.
	// *   If the number of available ECS instances is greater than the value of the Count parameter, the compute nodes are added based on the value of the Count parameter.
	//
	// Default value: 1
	MinCount                    *int32  `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// The duration of the subscription. The unit of the duration is specified by the PeriodUnit parameter. The parameter only takes effect when InstanceChargeType is set to PrePaid. Valid values:
	//
	// *   If PeriodUnit is set to Week, the valid values of the Period parameter are 1, 2, 3, and 4.
	// *   Valid values when PeriodUnit is set to Month: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// Default value: 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period. Valid values:
	//
	// *   Week
	// *   Month
	//
	// Default value: Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// Specifies whether to set the API operation as a synchronous operation. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false
	Sync *bool `json:"Sync,omitempty" xml:"Sync,omitempty"`
	// The performance level of the ESSD that is used as the system disk. Valid values:
	//
	// *   PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	// *   PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	// *   PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	// *   PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL1
	//
	// For more information about ESSD performance parameters, see [ESSD](~~122389~~).
	SystemDiskLevel *string `json:"SystemDiskLevel,omitempty" xml:"SystemDiskLevel,omitempty"`
	// The size of the system disk. Unit: GiB
	//
	// Valid values: 40 to 500
	//
	// Default value: 40
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The type of the system disk. Valid values:
	//
	// *   cloud_efficiency: ultra disk.
	// *   cloud_ssd: SSD.
	// *   cloud_essd: ESSD.
	// *   cloud: basic disk. Disks of this type are retired.
	//
	// Default value: cloud_efficiency
	SystemDiskType *string `json:"SystemDiskType,omitempty" xml:"SystemDiskType,omitempty"`
	// The ID of the vSwitch.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AddNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddNodesRequest) GoString() string {
	return s.String()
}

func (s *AddNodesRequest) SetAllocatePublicAddress(v bool) *AddNodesRequest {
	s.AllocatePublicAddress = &v
	return s
}

func (s *AddNodesRequest) SetAutoRenew(v string) *AddNodesRequest {
	s.AutoRenew = &v
	return s
}

func (s *AddNodesRequest) SetAutoRenewPeriod(v int32) *AddNodesRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *AddNodesRequest) SetClientToken(v string) *AddNodesRequest {
	s.ClientToken = &v
	return s
}

func (s *AddNodesRequest) SetClusterId(v string) *AddNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *AddNodesRequest) SetComputeEnableHt(v bool) *AddNodesRequest {
	s.ComputeEnableHt = &v
	return s
}

func (s *AddNodesRequest) SetComputeSpotDuration(v int32) *AddNodesRequest {
	s.ComputeSpotDuration = &v
	return s
}

func (s *AddNodesRequest) SetComputeSpotInterruptionBehavior(v string) *AddNodesRequest {
	s.ComputeSpotInterruptionBehavior = &v
	return s
}

func (s *AddNodesRequest) SetComputeSpotPriceLimit(v string) *AddNodesRequest {
	s.ComputeSpotPriceLimit = &v
	return s
}

func (s *AddNodesRequest) SetComputeSpotStrategy(v string) *AddNodesRequest {
	s.ComputeSpotStrategy = &v
	return s
}

func (s *AddNodesRequest) SetCount(v int32) *AddNodesRequest {
	s.Count = &v
	return s
}

func (s *AddNodesRequest) SetCreateMode(v string) *AddNodesRequest {
	s.CreateMode = &v
	return s
}

func (s *AddNodesRequest) SetDataDisks(v []*AddNodesRequestDataDisks) *AddNodesRequest {
	s.DataDisks = v
	return s
}

func (s *AddNodesRequest) SetEcsChargeType(v string) *AddNodesRequest {
	s.EcsChargeType = &v
	return s
}

func (s *AddNodesRequest) SetHostNamePrefix(v string) *AddNodesRequest {
	s.HostNamePrefix = &v
	return s
}

func (s *AddNodesRequest) SetHostNameSuffix(v string) *AddNodesRequest {
	s.HostNameSuffix = &v
	return s
}

func (s *AddNodesRequest) SetImageId(v string) *AddNodesRequest {
	s.ImageId = &v
	return s
}

func (s *AddNodesRequest) SetImageOwnerAlias(v string) *AddNodesRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *AddNodesRequest) SetInstanceType(v string) *AddNodesRequest {
	s.InstanceType = &v
	return s
}

func (s *AddNodesRequest) SetInternetChargeType(v string) *AddNodesRequest {
	s.InternetChargeType = &v
	return s
}

func (s *AddNodesRequest) SetInternetMaxBandWidthIn(v int32) *AddNodesRequest {
	s.InternetMaxBandWidthIn = &v
	return s
}

func (s *AddNodesRequest) SetInternetMaxBandWidthOut(v int32) *AddNodesRequest {
	s.InternetMaxBandWidthOut = &v
	return s
}

func (s *AddNodesRequest) SetJobQueue(v string) *AddNodesRequest {
	s.JobQueue = &v
	return s
}

func (s *AddNodesRequest) SetMinCount(v int32) *AddNodesRequest {
	s.MinCount = &v
	return s
}

func (s *AddNodesRequest) SetNetworkInterfaceTrafficMode(v string) *AddNodesRequest {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *AddNodesRequest) SetPeriod(v int32) *AddNodesRequest {
	s.Period = &v
	return s
}

func (s *AddNodesRequest) SetPeriodUnit(v string) *AddNodesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *AddNodesRequest) SetSync(v bool) *AddNodesRequest {
	s.Sync = &v
	return s
}

func (s *AddNodesRequest) SetSystemDiskLevel(v string) *AddNodesRequest {
	s.SystemDiskLevel = &v
	return s
}

func (s *AddNodesRequest) SetSystemDiskSize(v int32) *AddNodesRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *AddNodesRequest) SetSystemDiskType(v string) *AddNodesRequest {
	s.SystemDiskType = &v
	return s
}

func (s *AddNodesRequest) SetVSwitchId(v string) *AddNodesRequest {
	s.VSwitchId = &v
	return s
}

func (s *AddNodesRequest) SetZoneId(v string) *AddNodesRequest {
	s.ZoneId = &v
	return s
}

type AddNodesRequestDataDisks struct {
	// The type of the data disk. Valid values:
	//
	// *   cloud_efficiency: ultra disk
	// *   cloud_ssd: SSD
	// *   cloud_essd: ESSD
	// *   cloud: basic disk
	//
	// Default value: cloud_efficiency
	//
	// Valid values of N: 0 to 16
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// Specifies whether the data disk is released when the node is released. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: true
	//
	// Valid values of N: 0 to 16
	DataDiskDeleteWithInstance *bool `json:"DataDiskDeleteWithInstance,omitempty" xml:"DataDiskDeleteWithInstance,omitempty"`
	// Specifies whether to encrypt the data disk. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false
	//
	// Valid values of N: 0 to 16
	DataDiskEncrypted *bool `json:"DataDiskEncrypted,omitempty" xml:"DataDiskEncrypted,omitempty"`
	// The KMS key ID of the data disk.
	//
	// Valid values of N: 0 to 16
	DataDiskKMSKeyId *string `json:"DataDiskKMSKeyId,omitempty" xml:"DataDiskKMSKeyId,omitempty"`
	// The performance level of the ESSD used as the data disk. The parameter only takes effect only when the DataDisks.N.DataDiskCategory parameter is set to cloud_essd. Valid values:
	//
	// *   PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	// *   PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	// *   PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	// *   PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL1
	//
	// Valid values of N: 0 to 16
	DataDiskPerformanceLevel *string `json:"DataDiskPerformanceLevel,omitempty" xml:"DataDiskPerformanceLevel,omitempty"`
	// The size of the data disk. Unit: GB
	//
	// Valid values: 40 to 500
	//
	// Default value: 40
	//
	// Valid values of N: 0 to 16
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
}

func (s AddNodesRequestDataDisks) String() string {
	return tea.Prettify(s)
}

func (s AddNodesRequestDataDisks) GoString() string {
	return s.String()
}

func (s *AddNodesRequestDataDisks) SetDataDiskCategory(v string) *AddNodesRequestDataDisks {
	s.DataDiskCategory = &v
	return s
}

func (s *AddNodesRequestDataDisks) SetDataDiskDeleteWithInstance(v bool) *AddNodesRequestDataDisks {
	s.DataDiskDeleteWithInstance = &v
	return s
}

func (s *AddNodesRequestDataDisks) SetDataDiskEncrypted(v bool) *AddNodesRequestDataDisks {
	s.DataDiskEncrypted = &v
	return s
}

func (s *AddNodesRequestDataDisks) SetDataDiskKMSKeyId(v string) *AddNodesRequestDataDisks {
	s.DataDiskKMSKeyId = &v
	return s
}

func (s *AddNodesRequestDataDisks) SetDataDiskPerformanceLevel(v string) *AddNodesRequestDataDisks {
	s.DataDiskPerformanceLevel = &v
	return s
}

func (s *AddNodesRequestDataDisks) SetDataDiskSize(v int32) *AddNodesRequestDataDisks {
	s.DataDiskSize = &v
	return s
}

type AddNodesResponseBody struct {
	// The ID of the instance.
	//
	// >  AddNodes is an asynchronous API operation. If a request succeeds, a response is immediately generated before ECS instances are created. Therefore, the value of the parameter is null. You can call the [ListNodes](~~87161~~) operation to obtain the IDs of the ECS instances.
	InstanceIds *AddNodesResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddNodesResponseBody) GoString() string {
	return s.String()
}

func (s *AddNodesResponseBody) SetInstanceIds(v *AddNodesResponseBodyInstanceIds) *AddNodesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *AddNodesResponseBody) SetRequestId(v string) *AddNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddNodesResponseBody) SetTaskId(v string) *AddNodesResponseBody {
	s.TaskId = &v
	return s
}

type AddNodesResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s AddNodesResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s AddNodesResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *AddNodesResponseBodyInstanceIds) SetInstanceId(v []*string) *AddNodesResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

type AddNodesResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddNodesResponse) GoString() string {
	return s.String()
}

func (s *AddNodesResponse) SetHeaders(v map[string]*string) *AddNodesResponse {
	s.Headers = v
	return s
}

func (s *AddNodesResponse) SetStatusCode(v int32) *AddNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddNodesResponse) SetBody(v *AddNodesResponseBody) *AddNodesResponse {
	s.Body = v
	return s
}

type AddQueueRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the queue. The name must be 1 to 63 characters in length and start with a letter. It can contain letters, digits, and underscores (\_).
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s AddQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s AddQueueRequest) GoString() string {
	return s.String()
}

func (s *AddQueueRequest) SetClusterId(v string) *AddQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *AddQueueRequest) SetQueueName(v string) *AddQueueRequest {
	s.QueueName = &v
	return s
}

type AddQueueResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddQueueResponseBody) GoString() string {
	return s.String()
}

func (s *AddQueueResponseBody) SetRequestId(v string) *AddQueueResponseBody {
	s.RequestId = &v
	return s
}

type AddQueueResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s AddQueueResponse) GoString() string {
	return s.String()
}

func (s *AddQueueResponse) SetHeaders(v map[string]*string) *AddQueueResponse {
	s.Headers = v
	return s
}

func (s *AddQueueResponse) SetStatusCode(v int32) *AddQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *AddQueueResponse) SetBody(v *AddQueueResponseBody) *AddQueueResponse {
	s.Body = v
	return s
}

type AddSecurityGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure the idempotence of a request](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the security group.
	//
	// You can call the [DescribeSecurityGroups](~~25556~~) operation to query available security groups in the current region.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s AddSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *AddSecurityGroupRequest) SetClientToken(v string) *AddSecurityGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AddSecurityGroupRequest) SetClusterId(v string) *AddSecurityGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *AddSecurityGroupRequest) SetSecurityGroupId(v string) *AddSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

type AddSecurityGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSecurityGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddSecurityGroupResponseBody) SetRequestId(v string) *AddSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

type AddSecurityGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *AddSecurityGroupResponse) SetHeaders(v map[string]*string) *AddSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *AddSecurityGroupResponse) SetStatusCode(v int32) *AddSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSecurityGroupResponse) SetBody(v *AddSecurityGroupResponseBody) *AddSecurityGroupResponse {
	s.Body = v
	return s
}

type AddUsersRequest struct {
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	User      []*AddUsersRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s AddUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUsersRequest) GoString() string {
	return s.String()
}

func (s *AddUsersRequest) SetAsync(v bool) *AddUsersRequest {
	s.Async = &v
	return s
}

func (s *AddUsersRequest) SetClusterId(v string) *AddUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *AddUsersRequest) SetUser(v []*AddUsersRequestUser) *AddUsersRequest {
	s.User = v
	return s
}

type AddUsersRequestUser struct {
	// The permission group to which the user belongs. Valid values:
	//
	// *   users: an ordinary permission group. It is applicable to ordinary users that need only to submit and debug jobs.
	// *   wheel: a sudo permission group. It is applicable to the administrator who needs to manage the cluster. In addition to submitting and debugging jobs, users who have sudo permissions can run sudo commands to install software and restart nodes.
	//
	// Valid values of N: 1 to 100
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The name of the user that you want to add. The name must be 6 to 30 characters in length and can contain letters, digits, and periods (.). It must start with a letter.
	//
	// Valid values of N: 1 to 100
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The password of the user. The password must be 8 to 30 characters in length and contain three of the following items:
	//
	// *   Uppercase letter
	// *   Lowercase letter
	// *   Digit
	// *   Special character: `()~!@#$%^&*-_+=|{}[]:;\"/<>,.?/`
	//
	// Valid values of N: 1 to 100
	//
	// >  We recommend that you use HTTPS to call the AddUsers operation to ensure that the password remains confidential.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s AddUsersRequestUser) String() string {
	return tea.Prettify(s)
}

func (s AddUsersRequestUser) GoString() string {
	return s.String()
}

func (s *AddUsersRequestUser) SetGroup(v string) *AddUsersRequestUser {
	s.Group = &v
	return s
}

func (s *AddUsersRequestUser) SetName(v string) *AddUsersRequestUser {
	s.Name = &v
	return s
}

func (s *AddUsersRequestUser) SetPassword(v string) *AddUsersRequestUser {
	s.Password = &v
	return s
}

type AddUsersResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUsersResponseBody) GoString() string {
	return s.String()
}

func (s *AddUsersResponseBody) SetRequestId(v string) *AddUsersResponseBody {
	s.RequestId = &v
	return s
}

type AddUsersResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUsersResponse) GoString() string {
	return s.String()
}

func (s *AddUsersResponse) SetHeaders(v map[string]*string) *AddUsersResponse {
	s.Headers = v
	return s
}

func (s *AddUsersResponse) SetStatusCode(v int32) *AddUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUsersResponse) SetBody(v *AddUsersResponseBody) *AddUsersResponse {
	s.Body = v
	return s
}

type ApplyNodesRequest struct {
	// Specifies whether to allocate a public IP address to the compute nodes. Valid values:
	//
	// *   true: A public IP address is allocated to the compute nodes.
	// *   false: A public IP address is not allocated to the compute nodes.
	//
	// Default value: false
	AllocatePublicAddress *bool `json:"AllocatePublicAddress,omitempty" xml:"AllocatePublicAddress,omitempty"`
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87126~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The maximum hourly price of the compute nodes. The value is a floating-point number that supports up to three decimal places. The parameter takes effect only when ComputeSpotStrategy is set to SpotWithPriceLimit.
	//
	// If ComputeSpotPriceLimit and InstanceTypeModel.N.MaxPrice are specified at the same time, compute nodes are created based on the smaller value of these parameters.
	ComputeSpotPriceLimit *float32 `json:"ComputeSpotPriceLimit,omitempty" xml:"ComputeSpotPriceLimit,omitempty"`
	// The preemption policy of the compute nodes. Valid values:
	//
	// *   NoSpot: The compute nodes use the pay-as-you-go billing method.
	// *   SpotWithPriceLimit: The compute nodes are preemptible instances that have a user-defined maximum hourly price.
	// *   SpotAsPriceGo: The compute nodes are preemptible instances for which the market price at the time of purchase is used as the bid price.
	//
	// Default value: NoSpot
	ComputeSpotStrategy *string `json:"ComputeSpotStrategy,omitempty" xml:"ComputeSpotStrategy,omitempty"`
	// The number of vCPUs. The parameter is required when the ResourceAmountType parameter is set to Cores.
	//
	// You can set Cores, vCPU, and Memory to query node specifications. For example, you can query the available compute nodes that have 2 vCPUs and 16 GB of memory by setting vCPU to 2 and Memory to 16. You can also query compute nodes by zone. Query results are sorted by price.
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The prefix of the hostname. You can specify the parameter to manage the compute nodes in an efficient manner.
	HostNamePrefix *string `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	// The suffix of the hostname. You can specify the parameter to manage the compute nodes in an efficient manner.
	HostNameSuffix *string `json:"HostNameSuffix,omitempty" xml:"HostNameSuffix,omitempty"`
	// The image ID of the compute nodes to be added. The parameter takes effect only when the TargetImageId parameter is not specified.
	//
	// You can call the [ListImages](~~87213~~) and [ListCustomImages](~~87215~~) operations to query the image ID.
	//
	// >  If you add multiple compute nodes, the TargetImageId parameter takes effect only on the nodes for which the TargetImageId parameter is specified.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The level of the instance family. The parameter takes effect only when Cores and Memory are specified. Valid values:
	//
	// *   EntryLevel.
	// *   EnterpriseLevel.
	// *   CreditEntryLevel. For more information, see [What are burstable instances?](~~59977~~)
	//
	// Default value: EnterpriseLevel
	InstanceFamilyLevel *string                               `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	InstanceTypeModel   []*ApplyNodesRequestInstanceTypeModel `json:"InstanceTypeModel,omitempty" xml:"InstanceTypeModel,omitempty" type:"Repeated"`
	// The billing method of the elastic IP address (EIP). Valid values:
	//
	// *   PayByBandwidth: pay-by-bandwidth
	// *   PayByTraffic: pay-by-traffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth. Unit: Mbit/s. Valid values:
	//
	// *   If the purchased outbound public bandwidth is less than or equal to 10 Mbit/s, the valid values of the parameter are 1 to 10 and the default value is 10.
	// *   If the purchased outbound public bandwidth is greater than 10 Mbit/s, the valid values of this parameter are 1 to the amount of the outbound bandwidth that is purchased.
	InternetMaxBandWidthIn *int32 `json:"InternetMaxBandWidthIn,omitempty" xml:"InternetMaxBandWidthIn,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s. Valid values: 0 to 100.
	//
	// Default value: 0
	InternetMaxBandWidthOut *int32 `json:"InternetMaxBandWidthOut,omitempty" xml:"InternetMaxBandWidthOut,omitempty"`
	// The interval between two consecutive batches. Valid values: 60 to 600. Unit: seconds.
	//
	// Default value: 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The queue to which the compute nodes are added.
	//
	// You can call the [ListQueues](~~92176~~) operation to query the queue name.
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	// The memory capacity. The parameter is required when the ResourceAmountType parameter is set to Cores. Unit: GB.
	//
	// You can set Cores, vCPU, and Memory to query node specifications. For example, you can query the available compute nodes that have 2 vCPUs and 16 GB of memory by setting vCPU to 2 and Memory to 16. You can also query compute nodes by zone. Query results are sorted by price.
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The application policy of the preemptible nodes. Valid values:
	//
	// *   LowPriceResourcePlanning: Preemptible nodes are created based on the unit prices of vCPUs in ascending order. Preemptible nodes are created first when preemptible instance types are specified.
	// *   CapacityOptResourcePlanning: Preemptible nodes are created based on the prices and release rates in ascending order.
	// *   CustomizedResourcePlanning: Nodes are added based on the predefined value of the ZoneIds.N parameter. Instances of a zone that has a higher priority are used first.
	PriorityStrategy *string `json:"PriorityStrategy,omitempty" xml:"PriorityStrategy,omitempty"`
	// The type of the resource to be added. Valid values:
	//
	// *   Instances: compute node
	// *   Cores: vCPU and memory
	//
	// Default value: Instances
	ResourceAmountType *string `json:"ResourceAmountType,omitempty" xml:"ResourceAmountType,omitempty"`
	// The total number of batches to create nodes. Valid values: 1 to 10.
	//
	// Default value: 1
	Round *int32 `json:"Round,omitempty" xml:"Round,omitempty"`
	// Specifies whether to strictly meet the requirements of the TargetCapacity parameter. The parameter takes effect only when StrictSatisfiedTargetCapacity is set to true. Valid values:
	//
	// *   true: Check the inventory of the resources. Compute nodes are created based on the value of the TargetCapacity parameter only when the available resources are sufficient. Otherwise, no compute nodes are created.
	// *   false: Check the inventory of the resources. Compute nodes are created only when the available resources are sufficient. However, some compute nodes may fail to be created because resources become insufficient after the inventory is checked.
	//
	// Default value: false
	StrictResourceProvision *bool `json:"StrictResourceProvision,omitempty" xml:"StrictResourceProvision,omitempty"`
	// Specifies whether to meet the requirements of the TargetCapacity parameter. Valid values:
	//
	// *   true: If the available resources are fewer than the resources that you want to add, no compute nodes are created and an error is returned. If the available resources are more than the resources that you want to add, the following cases may occur:
	//
	//     *   If StrictResourceProvision is set to true, check the inventory of the resources. Compute nodes are created based on the value of the TargetCapacity parameter only when the available resources are sufficient. Otherwise, no compute nodes are created.
	//     *   If StrictResourceProvision is set to false, check the inventory of the resources. Compute nodes are created only when the available resources are sufficient. However, some compute nodes may fail to be created because resources become insufficient after the inventory is checked.
	//
	// *   false: If the available resources are insufficient, compute nodes are created based on the inventory of the resources.
	//
	// Default value: true
	StrictSatisfiedTargetCapacity *bool `json:"StrictSatisfiedTargetCapacity,omitempty" xml:"StrictSatisfiedTargetCapacity,omitempty"`
	// The performance level of the ESSD used as the system disk. Valid values:
	//
	// *   PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	// *   PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	// *   PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	// *   PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL0
	//
	// For more information, see [ESSDs](~~122389~~).
	SystemDiskLevel *string `json:"SystemDiskLevel,omitempty" xml:"SystemDiskLevel,omitempty"`
	// The size of the system disk. Unit: GB.
	//
	// Valid values: 40 to 500
	//
	// Default value: 40
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The type of the system disk. Valid values:
	//
	// *   cloud_efficiency: ultra disk.
	// *   cloud_ssd: SSD.
	// *   cloud_essd: ESSD.
	// *   cloud: basic disk. Disks of this type are retired.
	SystemDiskType *string                 `json:"SystemDiskType,omitempty" xml:"SystemDiskType,omitempty"`
	Tag            []*ApplyNodesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of the resource that you want to add. The specific number depends on the value of the ResourceAmountType parameter:
	//
	// *   If ResourceAmountType is set to Instance, the value range of TargetCapacity is 1 to 200.
	// *   If ResourceAmountType is set to Cores, the value range of TargetCapacity is 1 to 1,000.
	TargetCapacity *int32                        `json:"TargetCapacity,omitempty" xml:"TargetCapacity,omitempty"`
	ZoneInfos      []*ApplyNodesRequestZoneInfos `json:"ZoneInfos,omitempty" xml:"ZoneInfos,omitempty" type:"Repeated"`
}

func (s ApplyNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesRequest) GoString() string {
	return s.String()
}

func (s *ApplyNodesRequest) SetAllocatePublicAddress(v bool) *ApplyNodesRequest {
	s.AllocatePublicAddress = &v
	return s
}

func (s *ApplyNodesRequest) SetClusterId(v string) *ApplyNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ApplyNodesRequest) SetComputeSpotPriceLimit(v float32) *ApplyNodesRequest {
	s.ComputeSpotPriceLimit = &v
	return s
}

func (s *ApplyNodesRequest) SetComputeSpotStrategy(v string) *ApplyNodesRequest {
	s.ComputeSpotStrategy = &v
	return s
}

func (s *ApplyNodesRequest) SetCores(v int32) *ApplyNodesRequest {
	s.Cores = &v
	return s
}

func (s *ApplyNodesRequest) SetHostNamePrefix(v string) *ApplyNodesRequest {
	s.HostNamePrefix = &v
	return s
}

func (s *ApplyNodesRequest) SetHostNameSuffix(v string) *ApplyNodesRequest {
	s.HostNameSuffix = &v
	return s
}

func (s *ApplyNodesRequest) SetImageId(v string) *ApplyNodesRequest {
	s.ImageId = &v
	return s
}

func (s *ApplyNodesRequest) SetInstanceFamilyLevel(v string) *ApplyNodesRequest {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *ApplyNodesRequest) SetInstanceTypeModel(v []*ApplyNodesRequestInstanceTypeModel) *ApplyNodesRequest {
	s.InstanceTypeModel = v
	return s
}

func (s *ApplyNodesRequest) SetInternetChargeType(v string) *ApplyNodesRequest {
	s.InternetChargeType = &v
	return s
}

func (s *ApplyNodesRequest) SetInternetMaxBandWidthIn(v int32) *ApplyNodesRequest {
	s.InternetMaxBandWidthIn = &v
	return s
}

func (s *ApplyNodesRequest) SetInternetMaxBandWidthOut(v int32) *ApplyNodesRequest {
	s.InternetMaxBandWidthOut = &v
	return s
}

func (s *ApplyNodesRequest) SetInterval(v int32) *ApplyNodesRequest {
	s.Interval = &v
	return s
}

func (s *ApplyNodesRequest) SetJobQueue(v string) *ApplyNodesRequest {
	s.JobQueue = &v
	return s
}

func (s *ApplyNodesRequest) SetMemory(v int32) *ApplyNodesRequest {
	s.Memory = &v
	return s
}

func (s *ApplyNodesRequest) SetPriorityStrategy(v string) *ApplyNodesRequest {
	s.PriorityStrategy = &v
	return s
}

func (s *ApplyNodesRequest) SetResourceAmountType(v string) *ApplyNodesRequest {
	s.ResourceAmountType = &v
	return s
}

func (s *ApplyNodesRequest) SetRound(v int32) *ApplyNodesRequest {
	s.Round = &v
	return s
}

func (s *ApplyNodesRequest) SetStrictResourceProvision(v bool) *ApplyNodesRequest {
	s.StrictResourceProvision = &v
	return s
}

func (s *ApplyNodesRequest) SetStrictSatisfiedTargetCapacity(v bool) *ApplyNodesRequest {
	s.StrictSatisfiedTargetCapacity = &v
	return s
}

func (s *ApplyNodesRequest) SetSystemDiskLevel(v string) *ApplyNodesRequest {
	s.SystemDiskLevel = &v
	return s
}

func (s *ApplyNodesRequest) SetSystemDiskSize(v int32) *ApplyNodesRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *ApplyNodesRequest) SetSystemDiskType(v string) *ApplyNodesRequest {
	s.SystemDiskType = &v
	return s
}

func (s *ApplyNodesRequest) SetTag(v []*ApplyNodesRequestTag) *ApplyNodesRequest {
	s.Tag = v
	return s
}

func (s *ApplyNodesRequest) SetTargetCapacity(v int32) *ApplyNodesRequest {
	s.TargetCapacity = &v
	return s
}

func (s *ApplyNodesRequest) SetZoneInfos(v []*ApplyNodesRequestZoneInfos) *ApplyNodesRequest {
	s.ZoneInfos = v
	return s
}

type ApplyNodesRequestInstanceTypeModel struct {
	// The instance type of the compute node. The default value is the instance type that was specified when you created the cluster or the last time when you added compute nodes.
	//
	// Valid values of N: 1 to 10
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum hourly price that you can pay for the preemptible node. The value is a floating-point number that supports up to three decimal places.
	//
	// The parameter takes effect only when ComputeSpotStrategy is set to SpotWithPriceLimit.
	//
	// Valid values of N: 1 to 10
	MaxPrice *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	// The image ID of the compute node. You must select a Windows image.
	//
	// Valid values of N: 1 to 10
	TargetImageId *string `json:"TargetImageId,omitempty" xml:"TargetImageId,omitempty"`
}

func (s ApplyNodesRequestInstanceTypeModel) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesRequestInstanceTypeModel) GoString() string {
	return s.String()
}

func (s *ApplyNodesRequestInstanceTypeModel) SetInstanceType(v string) *ApplyNodesRequestInstanceTypeModel {
	s.InstanceType = &v
	return s
}

func (s *ApplyNodesRequestInstanceTypeModel) SetMaxPrice(v float32) *ApplyNodesRequestInstanceTypeModel {
	s.MaxPrice = &v
	return s
}

func (s *ApplyNodesRequestInstanceTypeModel) SetTargetImageId(v string) *ApplyNodesRequestInstanceTypeModel {
	s.TargetImageId = &v
	return s
}

type ApplyNodesRequestTag struct {
	// The tag key of the compute node that you want to attach. Valid values of N: 1 to 20. The tag key cannot be an empty string. It can be up to 128 characters in length and cannot start with acs: or aliyun. It cannot contain http:// or https://.
	//
	// Valid values of N: 1 to 10
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the compute node that you want to add. Valid values of N: 1 to 20. The tag value can be an empty string. It can be up to 128 characters in length and cannot start with acs: or contain http:// or https://.
	//
	// Valid values of N: 1 to 10
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ApplyNodesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesRequestTag) GoString() string {
	return s.String()
}

func (s *ApplyNodesRequestTag) SetKey(v string) *ApplyNodesRequestTag {
	s.Key = &v
	return s
}

func (s *ApplyNodesRequestTag) SetValue(v string) *ApplyNodesRequestTag {
	s.Value = &v
	return s
}

type ApplyNodesRequestZoneInfos struct {
	// The ID of the vSwitch. Valid values of N: 1 to 10.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone to which the cluster belongs. Valid values of N: 1 to 10.
	//
	// >  Each zone ID must be unique.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ApplyNodesRequestZoneInfos) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesRequestZoneInfos) GoString() string {
	return s.String()
}

func (s *ApplyNodesRequestZoneInfos) SetVSwitchId(v string) *ApplyNodesRequestZoneInfos {
	s.VSwitchId = &v
	return s
}

func (s *ApplyNodesRequestZoneInfos) SetZoneId(v string) *ApplyNodesRequestZoneInfos {
	s.ZoneId = &v
	return s
}

type ApplyNodesResponseBody struct {
	// The detailed result of the request.
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The ID of the Elastic Compute Service (ECS) instance.
	//
	// >  AddNodes is an asynchronous API operation. If a request succeeds, a response is immediately generated before ECS instances are created. Therefore, the value of the parameter is null. You can call the [ListNodes](~~87161~~) operation to query the ID of the ECS instance.
	InstanceIds *ApplyNodesResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	// The ID of the task.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of the compute nodes that were created.
	SatisfiedAmount *int32 `json:"SatisfiedAmount,omitempty" xml:"SatisfiedAmount,omitempty"`
	// The ID of the request.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ApplyNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyNodesResponseBody) SetDetail(v string) *ApplyNodesResponseBody {
	s.Detail = &v
	return s
}

func (s *ApplyNodesResponseBody) SetInstanceIds(v *ApplyNodesResponseBodyInstanceIds) *ApplyNodesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *ApplyNodesResponseBody) SetRequestId(v string) *ApplyNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyNodesResponseBody) SetSatisfiedAmount(v int32) *ApplyNodesResponseBody {
	s.SatisfiedAmount = &v
	return s
}

func (s *ApplyNodesResponseBody) SetTaskId(v string) *ApplyNodesResponseBody {
	s.TaskId = &v
	return s
}

type ApplyNodesResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s ApplyNodesResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *ApplyNodesResponseBodyInstanceIds) SetInstanceId(v []*string) *ApplyNodesResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

type ApplyNodesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesResponse) GoString() string {
	return s.String()
}

func (s *ApplyNodesResponse) SetHeaders(v map[string]*string) *ApplyNodesResponse {
	s.Headers = v
	return s
}

func (s *ApplyNodesResponse) SetStatusCode(v int32) *ApplyNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyNodesResponse) SetBody(v *ApplyNodesResponseBody) *ApplyNodesResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	EcsOrder *CreateClusterRequestEcsOrder `json:"EcsOrder,omitempty" xml:"EcsOrder,omitempty" type:"Struct"`
	// The type of the domain account service. Valid values:
	//
	// *   nis
	// *   ldap
	//
	// Default value: nis
	AccountType       *string                                  `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AdditionalVolumes []*CreateClusterRequestAdditionalVolumes `json:"AdditionalVolumes,omitempty" xml:"AdditionalVolumes,omitempty" type:"Repeated"`
	Application       []*CreateClusterRequestApplication       `json:"Application,omitempty" xml:"Application,omitempty" type:"Repeated"`
	// Specifies whether to enable auto-renewal for the subscription. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period of the subscription compute nodes. The parameter takes effect when AutoRenew is set to true.
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence of a request?](~~25693~~)
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The version of the E-HPC client. By default, the parameter is set to the latest version number.
	//
	// You can call the [ListCurrentClientVersion](~~87223~~) operation to query the latest version of the E-HPC client.
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The version of the E-HPC cluster.
	//
	// Default value: 1.0
	ClusterVersion *string `json:"ClusterVersion,omitempty" xml:"ClusterVersion,omitempty"`
	// Specifies whether the compute nodes support hyper-threading. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: true
	ComputeEnableHt *bool `json:"ComputeEnableHt,omitempty" xml:"ComputeEnableHt,omitempty"`
	// The maximum hourly price of the compute nodes. A maximum of three decimal places can be used in the value of the parameter. The parameter is valid only when the ComputeSpotStrategy parameter is set to SpotWithPriceLimit.
	ComputeSpotPriceLimit *string `json:"ComputeSpotPriceLimit,omitempty" xml:"ComputeSpotPriceLimit,omitempty"`
	// The bidding method of the compute nodes. Valid values:
	//
	// *   NoSpot: The compute nodes are pay-as-you-go instances.
	// *   SpotWithPriceLimit: The compute nodes are preemptible instances that have a user-defined maximum hourly price.
	// *   SpotAsPriceGo: The compute nodes are preemptible instances for which the market price at the time of purchase is used as the bid price.
	//
	// Default value: NoSpot
	ComputeSpotStrategy *string `json:"ComputeSpotStrategy,omitempty" xml:"ComputeSpotStrategy,omitempty"`
	// The mode in which the E-HPC cluster is deployed. Valid values:
	//
	// *   Standard: An account node, a scheduling node, a logon node, and multiple compute nodes are separately deployed.
	// *   Simple: A management node, a logon node, and multiple compute nodes are deployed. The management node consists of an account node and a scheduling node. The logon node and compute nodes are separately deployed.
	// *   Tiny: A management node and multiple compute nodes are deployed. The management node consists of an account node, a scheduling node, and a logon node. The compute nodes are separately deployed.
	//
	// Default value: Standard
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// The description of the E-HPC cluster. The description must be 2 to 256 characters in length. It cannot start with http:// or https://.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The domain name of the on-premises E-HPC cluster.
	//
	// This parameter takes effect only when the AccoutType parameter is set to Idap.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The billing method of the nodes. Valid values:
	//
	// *   PostPaid: pay-as-you-go
	// *   PrePaid: subscription
	//
	// If you set the parameter to PrePaid, auto-renewal is enabled by default.
	EcsChargeType *string `json:"EcsChargeType,omitempty" xml:"EcsChargeType,omitempty"`
	// The version of E-HPC. By default, the parameter is set to the latest version number.
	EhpcVersion *string `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	// Specifies whether to enable the high availability feature. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false
	//
	// >  If high availability is enabled, a primary management node and a secondary management node are used.
	HaEnable *bool `json:"HaEnable,omitempty" xml:"HaEnable,omitempty"`
	// The ID of the image.
	//
	// You can call the [ListImages](~~87213~~) and [ListCustomImages](~~87215~~) operations to query the images that are supported by E-HPC.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The type of the image. Valid values:
	//
	// *   system: public image
	// *   self: custom image
	// *   others: shared image
	//
	// Default value: system
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The URL of the job files that are uploaded to an Object Storage Service (OSS) bucket.
	InputFileUrl *string `json:"InputFileUrl,omitempty" xml:"InputFileUrl,omitempty"`
	// Specifies whether to enable auto scaling. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false
	IsComputeEss *bool `json:"IsComputeEss,omitempty" xml:"IsComputeEss,omitempty"`
	// The queue to which the compute nodes are added.
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	// The name of the AccessKey pair.
	//
	// >  For more information, see [Create an SSH key pair](~~51793~~).
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The name of the E-HPC cluster. The name must be 2 to 64 characters in length.
	Name                        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// The operating system tag of the image.
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	// The root password of the logon node. The password must be 8 to 30 characters in length and contain at least three of the following items: uppercase letters, lowercase letters, digits, and special characters. The password can contain the following special characters:
	//
	// `( ) ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ‘ < > , . ? /`
	//
	// You must specify either Password or KeyPairName. If both are specified, the Password parameter prevails.
	//
	// >  We recommend that you use HTTPS to call the API operation to prevent password leakages.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The duration of the subscription. The unit of the duration is specified by the `PeriodUnit` parameter.
	//
	// *   If you set PriceUnit to Year, the valid values of the Period parameter are 1, 2, and 3.
	// *   If you set PriceUnit to Month, the valid values of the Period parameter are 1, 2, 3, 4, 5, 6, 7, 8, and 9.
	// *   If you set PriceUnit to Hour, the valid value of the Period parameter is 1.
	//
	// Default value: 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// *   Year
	// *   Month
	// *   Hour
	//
	// Default value: Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The mode configurations of the plug-in. This parameter takes effect only when the SchedulerType parameter is set to custom.
	//
	// The value must be a JSON string. The parameter contains the following parameters: pluginMod, pluginLocalPath, and pluginOssPath.
	//
	// *   pluginMod: the mode of the plug-in. The following modes are supported:
	//
	//     *   oss: The plug-in is downloaded and decompressed from OSS to a local path. The local path is specified by the pluginLocalPath parameter.
	//     *   image: By default, the plug-in is stored in a pre-defined local path. The local path is specified by the pluginLocalPath parameter.
	//
	// *   pluginLocalPath: the local path where the plug-in is stored. We recommend that you select a shared directory in oss mode and a non-shared directory in image mode.
	//
	// *   pluginOssPath: the remote path where the plug-in is stored in OSS. This parameter takes effect only when the pluginMod parameter is set to oss.
	Plugin            *string                                  `json:"Plugin,omitempty" xml:"Plugin,omitempty"`
	PostInstallScript []*CreateClusterRequestPostInstallScript `json:"PostInstallScript,omitempty" xml:"PostInstallScript,omitempty" type:"Repeated"`
	RamNodeTypes      []*string                                `json:"RamNodeTypes,omitempty" xml:"RamNodeTypes,omitempty" type:"Repeated"`
	// The name of the Resource Access Management (RAM) role.
	//
	// You can call the [ListRoles](~~28713~~) operation provided by RAM to query the created RAM roles.
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The remote directory on which the file system is mounted.
	RemoteDirectory *string `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	// Specifies whether to enable Virtual Network Computing (VNC). Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false
	RemoteVisEnable *string `json:"RemoteVisEnable,omitempty" xml:"RemoteVisEnable,omitempty"`
	// The ID of the resource group.
	//
	// You can call the [ListResourceGroups](~~158855~~) operation to obtain the ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the Super Computing Cluster (SCC) instance.
	//
	// If you specify the parameter, the SCC instance is moved to a new SCC cluster.
	SccClusterId *string `json:"SccClusterId,omitempty" xml:"SccClusterId,omitempty"`
	// The type of the scheduler. Valid values:
	//
	// *   pbs
	// *   slurm
	// *   opengridscheduler
	// *   deadline
	//
	// Default value: pbs
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	// The ID of the security group to which the E-HPC cluster belongs.
	//
	// You can call the [DescribeSecurityGroups](~~25556~~) operation to query available security groups in the current region.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// If you do not use an existing security group, set the parameter to the name of a new security group. A default policy is applied to the new security group.
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	// The performance level of the ESSD that is used as the system disk. Valid values:
	//
	// *   PL0: An ESSD can deliver up to 10,000 random read/write IOPS.
	// *   PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	// *   PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	// *   PL3: An ESSD delivers up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL1
	//
	// For more information, see [ESSDs](~~122389~~).
	SystemDiskLevel *string `json:"SystemDiskLevel,omitempty" xml:"SystemDiskLevel,omitempty"`
	// The size of the system disk. Unit: GB.
	//
	// Valid values: 40 to 500
	//
	// Default value: 40
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The type of the system disk. Valid values:
	//
	// *   cloud_efficiency: ultra disk.
	// *   cloud_ssd: standard SSD.
	// *   cloud_essd: enhanced SSD (ESSD).
	// *   cloud: basic disk. Disks of this type are retired.
	//
	// Default value: cloud_ssd
	SystemDiskType *string                    `json:"SystemDiskType,omitempty" xml:"SystemDiskType,omitempty"`
	Tag            []*CreateClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch. E-HPC supports only VPC networks.
	//
	// You can call the [DescribeVSwitches](~~35748~~) operation to query available vSwitches.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the file system. If you leave the parameter empty, a Performance NAS file system is created by default.
	//
	// You can call the [ListFileSystemWithMountTargets](~~204364~~) operation to query available mount targets.
	VolumeId *string `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	// The mount options of the NFS file system that you want to mount by running the mount command.
	//
	// For more information, see [Mount an NFS file system on a Linux ECS instance](https://www.alibabacloud.com/help/en/nas/latest/mount-an-nfs-file-system-on-a-linux-ecs-instance#section-jyi-hyd-hbr).
	VolumeMountOption *string `json:"VolumeMountOption,omitempty" xml:"VolumeMountOption,omitempty"`
	// The mount target of the file system. Take note of the following information:
	//
	// *   If you do not specify the VolumeId parameter, you can leave the VolumeMountpoint parameter empty. A mount target is created by default.
	// *   If you specify the VolumeId parameter, the VolumeMountpoint parameter is required. You can call the [ListFileSystemWithMountTargets](~~204364~~) operation to query available mount targets.
	VolumeMountpoint *string `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	// The type of the protocol that is used by the file system. Valid values:
	//
	// *   NFS
	// *   SMB
	//
	// Default value: NFS
	VolumeProtocol *string `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	// The type of the shared storage. Set the value to `nas`, which indicates a NAS file system.
	VolumeType *string `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the E-HPC cluster belongs.
	//
	// You can call the [DescribeVpcs](~~35739~~) operation to query available VPCs.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Specifies whether not to install the agent.
	//
	// *   true: The agent is not installed.
	// *   false: The agent is installed.
	//
	// Default value: false
	WithoutAgent *bool `json:"WithoutAgent,omitempty" xml:"WithoutAgent,omitempty"`
	// Specifies whether the logon node uses an elastic IP address (EIP). Default value: false
	WithoutElasticIp *bool `json:"WithoutElasticIp,omitempty" xml:"WithoutElasticIp,omitempty"`
	// The ID of the zone.
	//
	// You can call the [ListRegions](~~188593~~) and [DescribeZones](~~25610~~) operations to query IDs of the zones where E-HPC is supported.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetEcsOrder(v *CreateClusterRequestEcsOrder) *CreateClusterRequest {
	s.EcsOrder = v
	return s
}

func (s *CreateClusterRequest) SetAccountType(v string) *CreateClusterRequest {
	s.AccountType = &v
	return s
}

func (s *CreateClusterRequest) SetAdditionalVolumes(v []*CreateClusterRequestAdditionalVolumes) *CreateClusterRequest {
	s.AdditionalVolumes = v
	return s
}

func (s *CreateClusterRequest) SetApplication(v []*CreateClusterRequestApplication) *CreateClusterRequest {
	s.Application = v
	return s
}

func (s *CreateClusterRequest) SetAutoRenew(v string) *CreateClusterRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateClusterRequest) SetAutoRenewPeriod(v int32) *CreateClusterRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateClusterRequest) SetClientToken(v string) *CreateClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateClusterRequest) SetClientVersion(v string) *CreateClusterRequest {
	s.ClientVersion = &v
	return s
}

func (s *CreateClusterRequest) SetClusterVersion(v string) *CreateClusterRequest {
	s.ClusterVersion = &v
	return s
}

func (s *CreateClusterRequest) SetComputeEnableHt(v bool) *CreateClusterRequest {
	s.ComputeEnableHt = &v
	return s
}

func (s *CreateClusterRequest) SetComputeSpotPriceLimit(v string) *CreateClusterRequest {
	s.ComputeSpotPriceLimit = &v
	return s
}

func (s *CreateClusterRequest) SetComputeSpotStrategy(v string) *CreateClusterRequest {
	s.ComputeSpotStrategy = &v
	return s
}

func (s *CreateClusterRequest) SetDeployMode(v string) *CreateClusterRequest {
	s.DeployMode = &v
	return s
}

func (s *CreateClusterRequest) SetDescription(v string) *CreateClusterRequest {
	s.Description = &v
	return s
}

func (s *CreateClusterRequest) SetDomain(v string) *CreateClusterRequest {
	s.Domain = &v
	return s
}

func (s *CreateClusterRequest) SetEcsChargeType(v string) *CreateClusterRequest {
	s.EcsChargeType = &v
	return s
}

func (s *CreateClusterRequest) SetEhpcVersion(v string) *CreateClusterRequest {
	s.EhpcVersion = &v
	return s
}

func (s *CreateClusterRequest) SetHaEnable(v bool) *CreateClusterRequest {
	s.HaEnable = &v
	return s
}

func (s *CreateClusterRequest) SetImageId(v string) *CreateClusterRequest {
	s.ImageId = &v
	return s
}

func (s *CreateClusterRequest) SetImageOwnerAlias(v string) *CreateClusterRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *CreateClusterRequest) SetInputFileUrl(v string) *CreateClusterRequest {
	s.InputFileUrl = &v
	return s
}

func (s *CreateClusterRequest) SetIsComputeEss(v bool) *CreateClusterRequest {
	s.IsComputeEss = &v
	return s
}

func (s *CreateClusterRequest) SetJobQueue(v string) *CreateClusterRequest {
	s.JobQueue = &v
	return s
}

func (s *CreateClusterRequest) SetKeyPairName(v string) *CreateClusterRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateClusterRequest) SetName(v string) *CreateClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateClusterRequest) SetNetworkInterfaceTrafficMode(v string) *CreateClusterRequest {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *CreateClusterRequest) SetOsTag(v string) *CreateClusterRequest {
	s.OsTag = &v
	return s
}

func (s *CreateClusterRequest) SetPassword(v string) *CreateClusterRequest {
	s.Password = &v
	return s
}

func (s *CreateClusterRequest) SetPeriod(v int32) *CreateClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateClusterRequest) SetPeriodUnit(v string) *CreateClusterRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateClusterRequest) SetPlugin(v string) *CreateClusterRequest {
	s.Plugin = &v
	return s
}

func (s *CreateClusterRequest) SetPostInstallScript(v []*CreateClusterRequestPostInstallScript) *CreateClusterRequest {
	s.PostInstallScript = v
	return s
}

func (s *CreateClusterRequest) SetRamNodeTypes(v []*string) *CreateClusterRequest {
	s.RamNodeTypes = v
	return s
}

func (s *CreateClusterRequest) SetRamRoleName(v string) *CreateClusterRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateClusterRequest) SetRemoteDirectory(v string) *CreateClusterRequest {
	s.RemoteDirectory = &v
	return s
}

func (s *CreateClusterRequest) SetRemoteVisEnable(v string) *CreateClusterRequest {
	s.RemoteVisEnable = &v
	return s
}

func (s *CreateClusterRequest) SetResourceGroupId(v string) *CreateClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetSccClusterId(v string) *CreateClusterRequest {
	s.SccClusterId = &v
	return s
}

func (s *CreateClusterRequest) SetSchedulerType(v string) *CreateClusterRequest {
	s.SchedulerType = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityGroupId(v string) *CreateClusterRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityGroupName(v string) *CreateClusterRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *CreateClusterRequest) SetSystemDiskLevel(v string) *CreateClusterRequest {
	s.SystemDiskLevel = &v
	return s
}

func (s *CreateClusterRequest) SetSystemDiskSize(v int32) *CreateClusterRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateClusterRequest) SetSystemDiskType(v string) *CreateClusterRequest {
	s.SystemDiskType = &v
	return s
}

func (s *CreateClusterRequest) SetTag(v []*CreateClusterRequestTag) *CreateClusterRequest {
	s.Tag = v
	return s
}

func (s *CreateClusterRequest) SetVSwitchId(v string) *CreateClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequest) SetVolumeId(v string) *CreateClusterRequest {
	s.VolumeId = &v
	return s
}

func (s *CreateClusterRequest) SetVolumeMountOption(v string) *CreateClusterRequest {
	s.VolumeMountOption = &v
	return s
}

func (s *CreateClusterRequest) SetVolumeMountpoint(v string) *CreateClusterRequest {
	s.VolumeMountpoint = &v
	return s
}

func (s *CreateClusterRequest) SetVolumeProtocol(v string) *CreateClusterRequest {
	s.VolumeProtocol = &v
	return s
}

func (s *CreateClusterRequest) SetVolumeType(v string) *CreateClusterRequest {
	s.VolumeType = &v
	return s
}

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequest) SetWithoutAgent(v bool) *CreateClusterRequest {
	s.WithoutAgent = &v
	return s
}

func (s *CreateClusterRequest) SetWithoutElasticIp(v bool) *CreateClusterRequest {
	s.WithoutElasticIp = &v
	return s
}

func (s *CreateClusterRequest) SetZoneId(v string) *CreateClusterRequest {
	s.ZoneId = &v
	return s
}

type CreateClusterRequestEcsOrder struct {
	Compute *CreateClusterRequestEcsOrderCompute `json:"Compute,omitempty" xml:"Compute,omitempty" require:"true" type:"Struct"`
	Login   *CreateClusterRequestEcsOrderLogin   `json:"Login,omitempty" xml:"Login,omitempty" require:"true" type:"Struct"`
	Manager *CreateClusterRequestEcsOrderManager `json:"Manager,omitempty" xml:"Manager,omitempty" require:"true" type:"Struct"`
}

func (s CreateClusterRequestEcsOrder) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestEcsOrder) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestEcsOrder) SetCompute(v *CreateClusterRequestEcsOrderCompute) *CreateClusterRequestEcsOrder {
	s.Compute = v
	return s
}

func (s *CreateClusterRequestEcsOrder) SetLogin(v *CreateClusterRequestEcsOrderLogin) *CreateClusterRequestEcsOrder {
	s.Login = v
	return s
}

func (s *CreateClusterRequestEcsOrder) SetManager(v *CreateClusterRequestEcsOrderManager) *CreateClusterRequestEcsOrder {
	s.Manager = v
	return s
}

type CreateClusterRequestEcsOrderCompute struct {
	// The number of the compute nodes. Valid values: 1 to 99.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The instance type of the compute nodes.
	//
	// You can call the [ListPreferredEcsTypes](~~188592~~) operation to query the recommended instance types.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateClusterRequestEcsOrderCompute) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestEcsOrderCompute) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestEcsOrderCompute) SetCount(v int32) *CreateClusterRequestEcsOrderCompute {
	s.Count = &v
	return s
}

func (s *CreateClusterRequestEcsOrderCompute) SetInstanceType(v string) *CreateClusterRequestEcsOrderCompute {
	s.InstanceType = &v
	return s
}

type CreateClusterRequestEcsOrderLogin struct {
	// The number of the logon nodes. Valid value: 1.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The instance type of the logon nodes.
	//
	// You can call the [ListPreferredEcsTypes](~~188592~~) operation to query the recommended instance types.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateClusterRequestEcsOrderLogin) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestEcsOrderLogin) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestEcsOrderLogin) SetCount(v int32) *CreateClusterRequestEcsOrderLogin {
	s.Count = &v
	return s
}

func (s *CreateClusterRequestEcsOrderLogin) SetInstanceType(v string) *CreateClusterRequestEcsOrderLogin {
	s.InstanceType = &v
	return s
}

type CreateClusterRequestEcsOrderManager struct {
	// The number of the management nodes. Valid values: 1 and 2.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The instance type of the management nodes.
	//
	// You can call the [ListPreferredEcsTypes](~~188592~~) operation to query the recommended instance types.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateClusterRequestEcsOrderManager) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestEcsOrderManager) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestEcsOrderManager) SetCount(v int32) *CreateClusterRequestEcsOrderManager {
	s.Count = &v
	return s
}

func (s *CreateClusterRequestEcsOrderManager) SetInstanceType(v string) *CreateClusterRequestEcsOrderManager {
	s.InstanceType = &v
	return s
}

type CreateClusterRequestAdditionalVolumes struct {
	// The queue of the nodes to which the additional file system is attached.
	//
	// Valid values of N: 1 to 10
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	// The local directory on which the additional file system is mounted.
	//
	// Valid values of N: 1 to 10
	LocalDirectory *string `json:"LocalDirectory,omitempty" xml:"LocalDirectory,omitempty"`
	// The type of the E-HPC cluster. Set the value to PublicCloud.
	//
	// Valid values of N: 1 to 10
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The remote directory on which the additional file system is mounted.
	//
	// Valid values of N: 1 to 10
	RemoteDirectory *string                                       `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	Roles           []*CreateClusterRequestAdditionalVolumesRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// The ID of the additional file system.
	//
	// Valid values of N: 1 to 10
	VolumeId *string `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	// The mount options of the additional file system.
	//
	// Valid values of N: 1 to 10
	VolumeMountOption *string `json:"VolumeMountOption,omitempty" xml:"VolumeMountOption,omitempty"`
	// The mount target of the additional file system.
	//
	// Valid values of N: 1 to 10
	VolumeMountpoint *string `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	// The type of the protocol that is used by the additional file system. Valid values:
	//
	// *   NFS
	// *   SMB
	//
	// Valid values of N: 1 to 10
	//
	// Default value: NFS
	VolumeProtocol *string `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	// The type of the additional shared storage. Only NAS file systems are supported.
	//
	// Valid values of N: 1 to 10
	VolumeType *string `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
}

func (s CreateClusterRequestAdditionalVolumes) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestAdditionalVolumes) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestAdditionalVolumes) SetJobQueue(v string) *CreateClusterRequestAdditionalVolumes {
	s.JobQueue = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetLocalDirectory(v string) *CreateClusterRequestAdditionalVolumes {
	s.LocalDirectory = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetLocation(v string) *CreateClusterRequestAdditionalVolumes {
	s.Location = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetRemoteDirectory(v string) *CreateClusterRequestAdditionalVolumes {
	s.RemoteDirectory = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetRoles(v []*CreateClusterRequestAdditionalVolumesRoles) *CreateClusterRequestAdditionalVolumes {
	s.Roles = v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetVolumeId(v string) *CreateClusterRequestAdditionalVolumes {
	s.VolumeId = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetVolumeMountOption(v string) *CreateClusterRequestAdditionalVolumes {
	s.VolumeMountOption = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetVolumeMountpoint(v string) *CreateClusterRequestAdditionalVolumes {
	s.VolumeMountpoint = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetVolumeProtocol(v string) *CreateClusterRequestAdditionalVolumes {
	s.VolumeProtocol = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetVolumeType(v string) *CreateClusterRequestAdditionalVolumes {
	s.VolumeType = &v
	return s
}

type CreateClusterRequestAdditionalVolumesRoles struct {
	// The type of the nodes to which the additional file system is attached.
	//
	// Valid values of N in AdditionalVolumes.N.Roles: 1 to 10
	//
	// Valid values of N in Roles.N.Name: 0 to 8
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateClusterRequestAdditionalVolumesRoles) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestAdditionalVolumesRoles) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestAdditionalVolumesRoles) SetName(v string) *CreateClusterRequestAdditionalVolumesRoles {
	s.Name = &v
	return s
}

type CreateClusterRequestApplication struct {
	// The tag of the software.
	//
	// Valid values of N: 0 to 100
	//
	// You can call the [ListSoftwares](~~87216~~) operation to query the tag of the software.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateClusterRequestApplication) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestApplication) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestApplication) SetTag(v string) *CreateClusterRequestApplication {
	s.Tag = &v
	return s
}

type CreateClusterRequestPostInstallScript struct {
	// The parameter that is used to run the script after the E-HPC cluster is created.
	//
	// Valid values of N: 0 to 16
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The URL that is used to download the script after the E-HPC cluster is created.
	//
	// Valid values of N: 0 to 16
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateClusterRequestPostInstallScript) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestPostInstallScript) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestPostInstallScript) SetArgs(v string) *CreateClusterRequestPostInstallScript {
	s.Args = &v
	return s
}

func (s *CreateClusterRequestPostInstallScript) SetUrl(v string) *CreateClusterRequestPostInstallScript {
	s.Url = &v
	return s
}

type CreateClusterRequestTag struct {
	// The key of the tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateClusterRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestTag) SetKey(v string) *CreateClusterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateClusterRequestTag) SetValue(v string) *CreateClusterRequestTag {
	s.Value = &v
	return s
}

type CreateClusterResponseBody struct {
	// The ID of the E-HPC cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// >  CreateCluster is an asynchronous API operation. If a request succeeds, a response is immediately generated before nodes are created. You can call the [ListTasks](~~268225~~) operation to query the result of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateClusterResponse) SetStatusCode(v int32) *CreateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateGWSClusterRequest struct {
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateGWSClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateGWSClusterRequest) SetClusterType(v string) *CreateGWSClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateGWSClusterRequest) SetName(v string) *CreateGWSClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateGWSClusterRequest) SetVSwitchId(v string) *CreateGWSClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateGWSClusterRequest) SetVpcId(v string) *CreateGWSClusterRequest {
	s.VpcId = &v
	return s
}

type CreateGWSClusterResponseBody struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGWSClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGWSClusterResponseBody) SetClusterId(v string) *CreateGWSClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateGWSClusterResponseBody) SetRequestId(v string) *CreateGWSClusterResponseBody {
	s.RequestId = &v
	return s
}

type CreateGWSClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGWSClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGWSClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateGWSClusterResponse) SetHeaders(v map[string]*string) *CreateGWSClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateGWSClusterResponse) SetStatusCode(v int32) *CreateGWSClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGWSClusterResponse) SetBody(v *CreateGWSClusterResponseBody) *CreateGWSClusterResponse {
	s.Body = v
	return s
}

type CreateGWSImageRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateGWSImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSImageRequest) GoString() string {
	return s.String()
}

func (s *CreateGWSImageRequest) SetInstanceId(v string) *CreateGWSImageRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateGWSImageRequest) SetName(v string) *CreateGWSImageRequest {
	s.Name = &v
	return s
}

type CreateGWSImageResponseBody struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGWSImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGWSImageResponseBody) SetImageId(v string) *CreateGWSImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateGWSImageResponseBody) SetRequestId(v string) *CreateGWSImageResponseBody {
	s.RequestId = &v
	return s
}

type CreateGWSImageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGWSImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGWSImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSImageResponse) GoString() string {
	return s.String()
}

func (s *CreateGWSImageResponse) SetHeaders(v map[string]*string) *CreateGWSImageResponse {
	s.Headers = v
	return s
}

func (s *CreateGWSImageResponse) SetStatusCode(v int32) *CreateGWSImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGWSImageResponse) SetBody(v *CreateGWSImageResponseBody) *CreateGWSImageResponse {
	s.Body = v
	return s
}

type CreateGWSInstanceRequest struct {
	AllocatePublicAddress   *bool   `json:"AllocatePublicAddress,omitempty" xml:"AllocatePublicAddress,omitempty"`
	AppList                 *string `json:"AppList,omitempty" xml:"AppList,omitempty"`
	AutoRenew               *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ImageId                 *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceChargeType      *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceType            *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetChargeType      *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthIn  *int32  `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut *int32  `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Period                  *string `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit              *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	SystemDiskCategory      *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	SystemDiskSize          *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	VSwitchId               *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	WorkMode                *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s CreateGWSInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateGWSInstanceRequest) SetAllocatePublicAddress(v bool) *CreateGWSInstanceRequest {
	s.AllocatePublicAddress = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetAppList(v string) *CreateGWSInstanceRequest {
	s.AppList = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetAutoRenew(v bool) *CreateGWSInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetClusterId(v string) *CreateGWSInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetImageId(v string) *CreateGWSInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetInstanceChargeType(v string) *CreateGWSInstanceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetInstanceType(v string) *CreateGWSInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetInternetChargeType(v string) *CreateGWSInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetInternetMaxBandwidthIn(v int32) *CreateGWSInstanceRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetInternetMaxBandwidthOut(v int32) *CreateGWSInstanceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetName(v string) *CreateGWSInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetPeriod(v string) *CreateGWSInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetPeriodUnit(v string) *CreateGWSInstanceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetSystemDiskCategory(v string) *CreateGWSInstanceRequest {
	s.SystemDiskCategory = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetSystemDiskSize(v int32) *CreateGWSInstanceRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetVSwitchId(v string) *CreateGWSInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetWorkMode(v string) *CreateGWSInstanceRequest {
	s.WorkMode = &v
	return s
}

type CreateGWSInstanceResponseBody struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGWSInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGWSInstanceResponseBody) SetInstanceId(v string) *CreateGWSInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateGWSInstanceResponseBody) SetRequestId(v string) *CreateGWSInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateGWSInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGWSInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGWSInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateGWSInstanceResponse) SetHeaders(v map[string]*string) *CreateGWSInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateGWSInstanceResponse) SetStatusCode(v int32) *CreateGWSInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGWSInstanceResponse) SetBody(v *CreateGWSInstanceResponseBody) *CreateGWSInstanceResponse {
	s.Body = v
	return s
}

type CreateHybridClusterRequest struct {
	EcsOrder    *CreateHybridClusterRequestEcsOrder      `json:"EcsOrder,omitempty" xml:"EcsOrder,omitempty" type:"Struct"`
	Application []*CreateHybridClusterRequestApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure the idempotence of a request](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The version of the client. By default, the latest version is used.
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The maximum hourly price for the ECS instance under the compute node. A maximum of three decimal places can be used in the value of the parameter. The parameter is valid only when the ComputeSpotStrategy parameter is set to SpotWithPriceLimit.
	ComputeSpotPriceLimit *float32 `json:"ComputeSpotPriceLimit,omitempty" xml:"ComputeSpotPriceLimit,omitempty"`
	// The preemption policy of the compute nodes. Valid values:
	//
	// *   NoSpot: The compute nodes are pay-as-you-go instances.
	// *   SpotWithPriceLimit: The instances of the compute node are preemptible instances. These types of instances have a specified maximum hourly price.
	// *   SpotAsPriceGo: The instances of the compute node are preemptible instances. The price of these instances is based on the current market price.
	//
	// Default value: NoSpot
	ComputeSpotStrategy *string `json:"ComputeSpotStrategy,omitempty" xml:"ComputeSpotStrategy,omitempty"`
	// The description of the cluster. The description must be 2 to 256 characters in length. It cannot start with http:// or https://.
	//
	// Default value: null
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the domain account service. Valid values:
	//
	// *   nis
	// *   ldap
	//
	// Default value: nis
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The version of E-HPC. By default, the latest version is used.
	EhpcVersion *string `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	// The ID of the image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The type of the image. Valid values:
	//
	// *   system: public image
	// *   self: custom image
	// *   others: shared image
	// *   marketplace: Alibaba Cloud Marketplace image
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The default queue of the scale-out nodes.
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	// The name of the AccessKey pair. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (\_), and hyphens (-).
	//
	// >  For more information, see [Create an SSH key pair](~~51793~~).
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The location where the cluster resides. Set the value to OnPremise.
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// Specifies whether the cluster supports multiple operating systems. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false
	MultiOs *bool `json:"MultiOs,omitempty" xml:"MultiOs,omitempty"`
	// The name of the cluster. The name must be 2 to 64 characters in length, and can contain only letters, digits, hyphens (-), and underscores (\_). It must start with a letter.
	Name  *string                            `json:"Name,omitempty" xml:"Name,omitempty"`
	Nodes []*CreateHybridClusterRequestNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The path in which the on-premises file system is mounted on the nodes on the cloud.
	OnPremiseVolumeLocalPath *string `json:"OnPremiseVolumeLocalPath,omitempty" xml:"OnPremiseVolumeLocalPath,omitempty"`
	// The mount target of the on-premises file system.
	OnPremiseVolumeMountPoint *string `json:"OnPremiseVolumeMountPoint,omitempty" xml:"OnPremiseVolumeMountPoint,omitempty"`
	// The type of the protocol that is used by the on-premises file system. Only NFS is supported.
	OnPremiseVolumeProtocol *string `json:"OnPremiseVolumeProtocol,omitempty" xml:"OnPremiseVolumeProtocol,omitempty"`
	// The mount path of the on-premises file system.
	OnPremiseVolumeRemotePath *string                                `json:"OnPremiseVolumeRemotePath,omitempty" xml:"OnPremiseVolumeRemotePath,omitempty"`
	OpenldapPar               *CreateHybridClusterRequestOpenldapPar `json:"OpenldapPar,omitempty" xml:"OpenldapPar,omitempty" type:"Struct"`
	// The image tag of the operating system. You can call the [ListImages](~~87213~~) operation to query the image tag.
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	// The root password of the logon node. The password must be 8 to 30 characters in length and contain at least three of the following items: uppercase letters, lowercase letters, digits, and special characters. The password can contain the following special characters:
	//
	// `() ~ ! @ # $ % ^ & * - = + | { } [ ] : ; ‘ < > , . ? /`
	//
	// >  We recommend that you use HTTPS to call the API operation to prevent password leakage.
	Password          *string                                        `json:"Password,omitempty" xml:"Password,omitempty"`
	Plugin            *string                                        `json:"Plugin,omitempty" xml:"Plugin,omitempty"`
	PostInstallScript []*CreateHybridClusterRequestPostInstallScript `json:"PostInstallScript,omitempty" xml:"PostInstallScript,omitempty" type:"Repeated"`
	// The remote directory to which the file system is mounted.
	RemoteDirectory *string `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether the scheduler is preinstalled for the image. Valid values:
	//
	// *   true: The scheduler is preinstalled. When you create or add a node, you do not need to install the scheduler.
	// *   false: The scheduler is not preinstalled. When you create or add a cluster, you must install the scheduler.
	SchedulerPreInstall *bool `json:"SchedulerPreInstall,omitempty" xml:"SchedulerPreInstall,omitempty"`
	// You can select an existing security group.
	//
	// >  If you specify this parameter, you cannot specify the `SecurityGroupName` parameter at the same time.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// If you do not use an existing security group, set the parameter to the name of a new security group. A default policy is applied to the new security group.
	//
	// >  If you specify this parameter, you cannot specify the `SecurityGroupId` parameter at the same time.
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	// The ID of the vSwitch.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the file system. NAS file systems cannot be automatically created.
	VolumeId *string `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	// The mount target of the file system. Mount targets cannot be automatically created for NAS file systems.
	VolumeMountpoint *string `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	// The type of the protocol that is used by the file system. Only NFS is supported.
	VolumeProtocol *string `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	// The type of the file system. Only NAS file systems are supported.
	VolumeType *string `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the cluster belongs.
	VpcId    *string                             `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	WinAdPar *CreateHybridClusterRequestWinAdPar `json:"WinAdPar,omitempty" xml:"WinAdPar,omitempty" type:"Struct"`
	// The ID of the zone.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateHybridClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequest) SetEcsOrder(v *CreateHybridClusterRequestEcsOrder) *CreateHybridClusterRequest {
	s.EcsOrder = v
	return s
}

func (s *CreateHybridClusterRequest) SetApplication(v []*CreateHybridClusterRequestApplication) *CreateHybridClusterRequest {
	s.Application = v
	return s
}

func (s *CreateHybridClusterRequest) SetClientToken(v string) *CreateHybridClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateHybridClusterRequest) SetClientVersion(v string) *CreateHybridClusterRequest {
	s.ClientVersion = &v
	return s
}

func (s *CreateHybridClusterRequest) SetComputeSpotPriceLimit(v float32) *CreateHybridClusterRequest {
	s.ComputeSpotPriceLimit = &v
	return s
}

func (s *CreateHybridClusterRequest) SetComputeSpotStrategy(v string) *CreateHybridClusterRequest {
	s.ComputeSpotStrategy = &v
	return s
}

func (s *CreateHybridClusterRequest) SetDescription(v string) *CreateHybridClusterRequest {
	s.Description = &v
	return s
}

func (s *CreateHybridClusterRequest) SetDomain(v string) *CreateHybridClusterRequest {
	s.Domain = &v
	return s
}

func (s *CreateHybridClusterRequest) SetEhpcVersion(v string) *CreateHybridClusterRequest {
	s.EhpcVersion = &v
	return s
}

func (s *CreateHybridClusterRequest) SetImageId(v string) *CreateHybridClusterRequest {
	s.ImageId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetImageOwnerAlias(v string) *CreateHybridClusterRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *CreateHybridClusterRequest) SetJobQueue(v string) *CreateHybridClusterRequest {
	s.JobQueue = &v
	return s
}

func (s *CreateHybridClusterRequest) SetKeyPairName(v string) *CreateHybridClusterRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateHybridClusterRequest) SetLocation(v string) *CreateHybridClusterRequest {
	s.Location = &v
	return s
}

func (s *CreateHybridClusterRequest) SetMultiOs(v bool) *CreateHybridClusterRequest {
	s.MultiOs = &v
	return s
}

func (s *CreateHybridClusterRequest) SetName(v string) *CreateHybridClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateHybridClusterRequest) SetNodes(v []*CreateHybridClusterRequestNodes) *CreateHybridClusterRequest {
	s.Nodes = v
	return s
}

func (s *CreateHybridClusterRequest) SetOnPremiseVolumeLocalPath(v string) *CreateHybridClusterRequest {
	s.OnPremiseVolumeLocalPath = &v
	return s
}

func (s *CreateHybridClusterRequest) SetOnPremiseVolumeMountPoint(v string) *CreateHybridClusterRequest {
	s.OnPremiseVolumeMountPoint = &v
	return s
}

func (s *CreateHybridClusterRequest) SetOnPremiseVolumeProtocol(v string) *CreateHybridClusterRequest {
	s.OnPremiseVolumeProtocol = &v
	return s
}

func (s *CreateHybridClusterRequest) SetOnPremiseVolumeRemotePath(v string) *CreateHybridClusterRequest {
	s.OnPremiseVolumeRemotePath = &v
	return s
}

func (s *CreateHybridClusterRequest) SetOpenldapPar(v *CreateHybridClusterRequestOpenldapPar) *CreateHybridClusterRequest {
	s.OpenldapPar = v
	return s
}

func (s *CreateHybridClusterRequest) SetOsTag(v string) *CreateHybridClusterRequest {
	s.OsTag = &v
	return s
}

func (s *CreateHybridClusterRequest) SetPassword(v string) *CreateHybridClusterRequest {
	s.Password = &v
	return s
}

func (s *CreateHybridClusterRequest) SetPlugin(v string) *CreateHybridClusterRequest {
	s.Plugin = &v
	return s
}

func (s *CreateHybridClusterRequest) SetPostInstallScript(v []*CreateHybridClusterRequestPostInstallScript) *CreateHybridClusterRequest {
	s.PostInstallScript = v
	return s
}

func (s *CreateHybridClusterRequest) SetRemoteDirectory(v string) *CreateHybridClusterRequest {
	s.RemoteDirectory = &v
	return s
}

func (s *CreateHybridClusterRequest) SetResourceGroupId(v string) *CreateHybridClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetSchedulerPreInstall(v bool) *CreateHybridClusterRequest {
	s.SchedulerPreInstall = &v
	return s
}

func (s *CreateHybridClusterRequest) SetSecurityGroupId(v string) *CreateHybridClusterRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetSecurityGroupName(v string) *CreateHybridClusterRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVSwitchId(v string) *CreateHybridClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVolumeId(v string) *CreateHybridClusterRequest {
	s.VolumeId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVolumeMountpoint(v string) *CreateHybridClusterRequest {
	s.VolumeMountpoint = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVolumeProtocol(v string) *CreateHybridClusterRequest {
	s.VolumeProtocol = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVolumeType(v string) *CreateHybridClusterRequest {
	s.VolumeType = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVpcId(v string) *CreateHybridClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetWinAdPar(v *CreateHybridClusterRequestWinAdPar) *CreateHybridClusterRequest {
	s.WinAdPar = v
	return s
}

func (s *CreateHybridClusterRequest) SetZoneId(v string) *CreateHybridClusterRequest {
	s.ZoneId = &v
	return s
}

type CreateHybridClusterRequestEcsOrder struct {
	Compute *CreateHybridClusterRequestEcsOrderCompute `json:"Compute,omitempty" xml:"Compute,omitempty" require:"true" type:"Struct"`
	Manager *CreateHybridClusterRequestEcsOrderManager `json:"Manager,omitempty" xml:"Manager,omitempty" require:"true" type:"Struct"`
}

func (s CreateHybridClusterRequestEcsOrder) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestEcsOrder) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestEcsOrder) SetCompute(v *CreateHybridClusterRequestEcsOrderCompute) *CreateHybridClusterRequestEcsOrder {
	s.Compute = v
	return s
}

func (s *CreateHybridClusterRequestEcsOrder) SetManager(v *CreateHybridClusterRequestEcsOrderManager) *CreateHybridClusterRequestEcsOrder {
	s.Manager = v
	return s
}

type CreateHybridClusterRequestEcsOrderCompute struct {
	// The instance type of the compute nodes.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateHybridClusterRequestEcsOrderCompute) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestEcsOrderCompute) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestEcsOrderCompute) SetInstanceType(v string) *CreateHybridClusterRequestEcsOrderCompute {
	s.InstanceType = &v
	return s
}

type CreateHybridClusterRequestEcsOrderManager struct {
	// The instance type of the management node on the cloud. Only Proxy Mode is supported.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateHybridClusterRequestEcsOrderManager) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestEcsOrderManager) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestEcsOrderManager) SetInstanceType(v string) *CreateHybridClusterRequestEcsOrderManager {
	s.InstanceType = &v
	return s
}

type CreateHybridClusterRequestApplication struct {
	// The tag of the application. Valid values of N: 1 to 5.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateHybridClusterRequestApplication) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestApplication) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestApplication) SetTag(v string) *CreateHybridClusterRequestApplication {
	s.Tag = &v
	return s
}

type CreateHybridClusterRequestNodes struct {
	// The service type of the domain account to which the on-premises node in the cluster belongs. Valid values:
	//
	// *   nis
	// *   ldap
	//
	// Default value: nis
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The directory of the on-premises node in the cluster.
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// The hostname of the on-premises node in the cluster.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The IP address of the on-premises node in the cluster.
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The role of the on-premises node in the cluster. Valid values:
	//
	// *   Manager: management node
	// *   Login: logon node
	// *   Compute: compute node
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The scheduler type of the on-premises node in the cluster. Valid values:
	//
	// *   pbs
	// *   slurm
	// *   opengridscheduler
	// *   deadline
	//
	// Default value: pbs
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
}

func (s CreateHybridClusterRequestNodes) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestNodes) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestNodes) SetAccountType(v string) *CreateHybridClusterRequestNodes {
	s.AccountType = &v
	return s
}

func (s *CreateHybridClusterRequestNodes) SetDir(v string) *CreateHybridClusterRequestNodes {
	s.Dir = &v
	return s
}

func (s *CreateHybridClusterRequestNodes) SetHostName(v string) *CreateHybridClusterRequestNodes {
	s.HostName = &v
	return s
}

func (s *CreateHybridClusterRequestNodes) SetIpAddress(v string) *CreateHybridClusterRequestNodes {
	s.IpAddress = &v
	return s
}

func (s *CreateHybridClusterRequestNodes) SetRole(v string) *CreateHybridClusterRequestNodes {
	s.Role = &v
	return s
}

func (s *CreateHybridClusterRequestNodes) SetSchedulerType(v string) *CreateHybridClusterRequestNodes {
	s.SchedulerType = &v
	return s
}

type CreateHybridClusterRequestOpenldapPar struct {
	BaseDn       *string `json:"BaseDn,omitempty" xml:"BaseDn,omitempty"`
	LdapServerIp *string `json:"LdapServerIp,omitempty" xml:"LdapServerIp,omitempty"`
}

func (s CreateHybridClusterRequestOpenldapPar) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestOpenldapPar) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestOpenldapPar) SetBaseDn(v string) *CreateHybridClusterRequestOpenldapPar {
	s.BaseDn = &v
	return s
}

func (s *CreateHybridClusterRequestOpenldapPar) SetLdapServerIp(v string) *CreateHybridClusterRequestOpenldapPar {
	s.LdapServerIp = &v
	return s
}

type CreateHybridClusterRequestPostInstallScript struct {
	// The parameters that are used to run the post-installation script. Valid values of N: 1 to 16.
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The full path of the post-installation script. Valid values of N: 1 to 16.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateHybridClusterRequestPostInstallScript) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestPostInstallScript) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestPostInstallScript) SetArgs(v string) *CreateHybridClusterRequestPostInstallScript {
	s.Args = &v
	return s
}

func (s *CreateHybridClusterRequestPostInstallScript) SetUrl(v string) *CreateHybridClusterRequestPostInstallScript {
	s.Url = &v
	return s
}

type CreateHybridClusterRequestWinAdPar struct {
	AdDc         *string `json:"AdDc,omitempty" xml:"AdDc,omitempty"`
	AdIp         *string `json:"AdIp,omitempty" xml:"AdIp,omitempty"`
	AdUser       *string `json:"AdUser,omitempty" xml:"AdUser,omitempty"`
	AdUserPasswd *string `json:"AdUserPasswd,omitempty" xml:"AdUserPasswd,omitempty"`
}

func (s CreateHybridClusterRequestWinAdPar) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestWinAdPar) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestWinAdPar) SetAdDc(v string) *CreateHybridClusterRequestWinAdPar {
	s.AdDc = &v
	return s
}

func (s *CreateHybridClusterRequestWinAdPar) SetAdIp(v string) *CreateHybridClusterRequestWinAdPar {
	s.AdIp = &v
	return s
}

func (s *CreateHybridClusterRequestWinAdPar) SetAdUser(v string) *CreateHybridClusterRequestWinAdPar {
	s.AdUser = &v
	return s
}

func (s *CreateHybridClusterRequestWinAdPar) SetAdUserPasswd(v string) *CreateHybridClusterRequestWinAdPar {
	s.AdUserPasswd = &v
	return s
}

type CreateHybridClusterResponseBody struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateHybridClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterResponseBody) SetClusterId(v string) *CreateHybridClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateHybridClusterResponseBody) SetRequestId(v string) *CreateHybridClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHybridClusterResponseBody) SetTaskId(v string) *CreateHybridClusterResponseBody {
	s.TaskId = &v
	return s
}

type CreateHybridClusterResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateHybridClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHybridClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterResponse) SetHeaders(v map[string]*string) *CreateHybridClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateHybridClusterResponse) SetStatusCode(v int32) *CreateHybridClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHybridClusterResponse) SetBody(v *CreateHybridClusterResponseBody) *CreateHybridClusterResponse {
	s.Body = v
	return s
}

type CreateJobFileRequest struct {
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The content of the job file. The content is encoded in Base64.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The user to which the job belongs.
	//
	// You can call the [ListUsers](~~188572~~) operation to query the users of the cluster.
	RunasUser *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	// The password of the user.
	RunasUserPassword *string `json:"RunasUserPassword,omitempty" xml:"RunasUserPassword,omitempty"`
	// The name of the job file.
	TargetFile *string `json:"TargetFile,omitempty" xml:"TargetFile,omitempty"`
}

func (s CreateJobFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobFileRequest) GoString() string {
	return s.String()
}

func (s *CreateJobFileRequest) SetAsync(v bool) *CreateJobFileRequest {
	s.Async = &v
	return s
}

func (s *CreateJobFileRequest) SetClusterId(v string) *CreateJobFileRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateJobFileRequest) SetContent(v string) *CreateJobFileRequest {
	s.Content = &v
	return s
}

func (s *CreateJobFileRequest) SetRunasUser(v string) *CreateJobFileRequest {
	s.RunasUser = &v
	return s
}

func (s *CreateJobFileRequest) SetRunasUserPassword(v string) *CreateJobFileRequest {
	s.RunasUserPassword = &v
	return s
}

func (s *CreateJobFileRequest) SetTargetFile(v string) *CreateJobFileRequest {
	s.TargetFile = &v
	return s
}

type CreateJobFileResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateJobFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobFileResponseBody) SetRequestId(v string) *CreateJobFileResponseBody {
	s.RequestId = &v
	return s
}

type CreateJobFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateJobFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateJobFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobFileResponse) GoString() string {
	return s.String()
}

func (s *CreateJobFileResponse) SetHeaders(v map[string]*string) *CreateJobFileResponse {
	s.Headers = v
	return s
}

func (s *CreateJobFileResponse) SetStatusCode(v int32) *CreateJobFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobFileResponse) SetBody(v *CreateJobFileResponseBody) *CreateJobFileResponse {
	s.Body = v
	return s
}

type CreateJobTemplateRequest struct {
	// The job array.
	//
	// Format: X-Y:Z. X is the minimum index value. Y is the maximum index value. Z is the step size. For example, 2-7:2 indicates that three jobs need to be run and their index values are 2, 4, and 6.
	ArrayRequest *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	// The maximum running time of the job. Valid formats:
	//
	// *   hh:mm:ss
	// *   mm:ss
	// *   ss
	//
	// We recommend that you use the hh:mm:ss format. If the maximum running time is 12 hours, set the value to 12:00:00.
	ClockTime *string `json:"ClockTime,omitempty" xml:"ClockTime,omitempty"`
	// The command that is used to run the job.
	CommandLine *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	// The maximum GPU usage required by a single compute node. Valid values: 1 to 8.
	//
	// The parameter takes effect only when the cluster uses PBS and a compute node is a GPU-accelerated instance.
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The URL of the job files that are uploaded to an Object Storage Service (OSS) bucket.
	InputFileUrl *string `json:"InputFileUrl,omitempty" xml:"InputFileUrl,omitempty"`
	// The maximum memory usage required by a single compute node. Unit: GB, MB, or KB. The unit is case-insensitive.
	Mem *string `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The name of the job template. The name must be 2 to 64 characters in length. It must start with a letter and can contain letters, digits, hyphens (-), and underscores (\_).
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of compute nodes. Valid values: 1 to 500.
	//
	// >  If the parameter is not specified, the Task, Thread, Mem, and Gpu parameters become invalid.
	Node *int32 `json:"Node,omitempty" xml:"Node,omitempty"`
	// The path that is used to run the job.
	PackagePath *string `json:"PackagePath,omitempty" xml:"PackagePath,omitempty"`
	// The priority of the job. Valid values: 0 to 9. A large value indicates a high priority.
	//
	// Default value: 0
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The name of the queue in which the job is run.
	//
	// You can call the [ListQueues](~~92176~~) operation to query the queue name.
	Queue *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	// Specifies whether the job can be rerun. Valid values:
	//
	// *   true: The job can be rerun.
	// *   false: The job cannot be rerun.
	ReRunable *bool `json:"ReRunable,omitempty" xml:"ReRunable,omitempty"`
	// The name of the user that runs the job.
	//
	// You can call the [ListUsers](~~188572~~) operation to query the users of the cluster.
	RunasUser *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	// The output file path of stderr.
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" xml:"StderrRedirectPath,omitempty"`
	// The output file path of stdout.
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" xml:"StdoutRedirectPath,omitempty"`
	// The number of tasks required by a single compute node. Valid values: 1 to 1000.
	Task *int32 `json:"Task,omitempty" xml:"Task,omitempty"`
	// The number of threads required by a single compute node. Valid values: 1 to 1000.
	Thread *int32 `json:"Thread,omitempty" xml:"Thread,omitempty"`
	// The command that is used to decompress the job files downloaded from an OSS bucket. The parameter takes effect only when WithUnzipCmd is set to true. Valid values:
	//
	// *   tar xzf: decompresses GZIP files.
	// *   tar xf: decompresses TAR files.
	// *   unzip: decompresses ZIP files.
	UnzipCmd *string `json:"UnzipCmd,omitempty" xml:"UnzipCmd,omitempty"`
	// The runtime variables passed to the job. They can be accessed by using environment variables in the executable file.
	Variables *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
	// Specifies whether to decompress the job files downloaded from an OSS bucket. Valid values:
	//
	// *   true: The job files are decompressed.
	// *   false: The job files are not decompressed.
	WithUnzipCmd *bool `json:"WithUnzipCmd,omitempty" xml:"WithUnzipCmd,omitempty"`
}

func (s CreateJobTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequest) SetArrayRequest(v string) *CreateJobTemplateRequest {
	s.ArrayRequest = &v
	return s
}

func (s *CreateJobTemplateRequest) SetClockTime(v string) *CreateJobTemplateRequest {
	s.ClockTime = &v
	return s
}

func (s *CreateJobTemplateRequest) SetCommandLine(v string) *CreateJobTemplateRequest {
	s.CommandLine = &v
	return s
}

func (s *CreateJobTemplateRequest) SetGpu(v int32) *CreateJobTemplateRequest {
	s.Gpu = &v
	return s
}

func (s *CreateJobTemplateRequest) SetInputFileUrl(v string) *CreateJobTemplateRequest {
	s.InputFileUrl = &v
	return s
}

func (s *CreateJobTemplateRequest) SetMem(v string) *CreateJobTemplateRequest {
	s.Mem = &v
	return s
}

func (s *CreateJobTemplateRequest) SetName(v string) *CreateJobTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateJobTemplateRequest) SetNode(v int32) *CreateJobTemplateRequest {
	s.Node = &v
	return s
}

func (s *CreateJobTemplateRequest) SetPackagePath(v string) *CreateJobTemplateRequest {
	s.PackagePath = &v
	return s
}

func (s *CreateJobTemplateRequest) SetPriority(v int32) *CreateJobTemplateRequest {
	s.Priority = &v
	return s
}

func (s *CreateJobTemplateRequest) SetQueue(v string) *CreateJobTemplateRequest {
	s.Queue = &v
	return s
}

func (s *CreateJobTemplateRequest) SetReRunable(v bool) *CreateJobTemplateRequest {
	s.ReRunable = &v
	return s
}

func (s *CreateJobTemplateRequest) SetRunasUser(v string) *CreateJobTemplateRequest {
	s.RunasUser = &v
	return s
}

func (s *CreateJobTemplateRequest) SetStderrRedirectPath(v string) *CreateJobTemplateRequest {
	s.StderrRedirectPath = &v
	return s
}

func (s *CreateJobTemplateRequest) SetStdoutRedirectPath(v string) *CreateJobTemplateRequest {
	s.StdoutRedirectPath = &v
	return s
}

func (s *CreateJobTemplateRequest) SetTask(v int32) *CreateJobTemplateRequest {
	s.Task = &v
	return s
}

func (s *CreateJobTemplateRequest) SetThread(v int32) *CreateJobTemplateRequest {
	s.Thread = &v
	return s
}

func (s *CreateJobTemplateRequest) SetUnzipCmd(v string) *CreateJobTemplateRequest {
	s.UnzipCmd = &v
	return s
}

func (s *CreateJobTemplateRequest) SetVariables(v string) *CreateJobTemplateRequest {
	s.Variables = &v
	return s
}

func (s *CreateJobTemplateRequest) SetWithUnzipCmd(v bool) *CreateJobTemplateRequest {
	s.WithUnzipCmd = &v
	return s
}

type CreateJobTemplateResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the job template.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateJobTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateResponseBody) SetRequestId(v string) *CreateJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobTemplateResponseBody) SetTemplateId(v string) *CreateJobTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type CreateJobTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateJobTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateJobTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateResponse) SetHeaders(v map[string]*string) *CreateJobTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateJobTemplateResponse) SetStatusCode(v int32) *CreateJobTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobTemplateResponse) SetBody(v *CreateJobTemplateResponseBody) *CreateJobTemplateResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to release Elastic Compute Service (ECS) instances that are created by using Elastic High Performance Computing (E-HPC).
	//
	// Default value: true
	ReleaseInstance *string `json:"ReleaseInstance,omitempty" xml:"ReleaseInstance,omitempty"`
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

func (s *DeleteClusterRequest) SetReleaseInstance(v string) *DeleteClusterRequest {
	s.ReleaseInstance = &v
	return s
}

type DeleteClusterResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClusterResponseBody) SetTaskId(v string) *DeleteClusterResponseBody {
	s.TaskId = &v
	return s
}

type DeleteClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteClusterResponse) SetStatusCode(v int32) *DeleteClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type DeleteContainerAppsRequest struct {
	ContainerApp []*DeleteContainerAppsRequestContainerApp `json:"ContainerApp,omitempty" xml:"ContainerApp,omitempty" type:"Repeated"`
}

func (s DeleteContainerAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerAppsRequest) GoString() string {
	return s.String()
}

func (s *DeleteContainerAppsRequest) SetContainerApp(v []*DeleteContainerAppsRequestContainerApp) *DeleteContainerAppsRequest {
	s.ContainerApp = v
	return s
}

type DeleteContainerAppsRequestContainerApp struct {
	// The ID of the containerized application that you want to delete. Valid values of N: 1 to 100.
	//
	// You can call the [ListContainerApps](~~87333~~) operation to query the ID of the containerized application.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteContainerAppsRequestContainerApp) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerAppsRequestContainerApp) GoString() string {
	return s.String()
}

func (s *DeleteContainerAppsRequestContainerApp) SetId(v string) *DeleteContainerAppsRequestContainerApp {
	s.Id = &v
	return s
}

type DeleteContainerAppsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContainerAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContainerAppsResponseBody) SetRequestId(v string) *DeleteContainerAppsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteContainerAppsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteContainerAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteContainerAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerAppsResponse) GoString() string {
	return s.String()
}

func (s *DeleteContainerAppsResponse) SetHeaders(v map[string]*string) *DeleteContainerAppsResponse {
	s.Headers = v
	return s
}

func (s *DeleteContainerAppsResponse) SetStatusCode(v int32) *DeleteContainerAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContainerAppsResponse) SetBody(v *DeleteContainerAppsResponseBody) *DeleteContainerAppsResponse {
	s.Body = v
	return s
}

type DeleteGWSClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DeleteGWSClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteGWSClusterRequest) SetClusterId(v string) *DeleteGWSClusterRequest {
	s.ClusterId = &v
	return s
}

type DeleteGWSClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGWSClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGWSClusterResponseBody) SetRequestId(v string) *DeleteGWSClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGWSClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGWSClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGWSClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteGWSClusterResponse) SetHeaders(v map[string]*string) *DeleteGWSClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteGWSClusterResponse) SetStatusCode(v int32) *DeleteGWSClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGWSClusterResponse) SetBody(v *DeleteGWSClusterResponseBody) *DeleteGWSClusterResponse {
	s.Body = v
	return s
}

type DeleteGWSInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteGWSInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGWSInstanceRequest) SetInstanceId(v string) *DeleteGWSInstanceRequest {
	s.InstanceId = &v
	return s
}

type DeleteGWSInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGWSInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGWSInstanceResponseBody) SetRequestId(v string) *DeleteGWSInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGWSInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGWSInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGWSInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGWSInstanceResponse) SetHeaders(v map[string]*string) *DeleteGWSInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGWSInstanceResponse) SetStatusCode(v int32) *DeleteGWSInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGWSInstanceResponse) SetBody(v *DeleteGWSInstanceResponseBody) *DeleteGWSInstanceResponse {
	s.Body = v
	return s
}

type DeleteImageRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the container. Set the value to singularity.
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
	// The tags of the image.
	//
	// Default value: latest
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The name of the repository.
	//
	// You can call the [ListContainerImages](~~87348~~) operation to query the name of the repository.
	Repository *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
}

func (s DeleteImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) SetClusterId(v string) *DeleteImageRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteImageRequest) SetContainerType(v string) *DeleteImageRequest {
	s.ContainerType = &v
	return s
}

func (s *DeleteImageRequest) SetImageTag(v string) *DeleteImageRequest {
	s.ImageTag = &v
	return s
}

func (s *DeleteImageRequest) SetRepository(v string) *DeleteImageRequest {
	s.Repository = &v
	return s
}

type DeleteImageResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBody) SetRequestId(v string) *DeleteImageResponseBody {
	s.RequestId = &v
	return s
}

type DeleteImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageResponse) SetHeaders(v map[string]*string) *DeleteImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageResponse) SetStatusCode(v int32) *DeleteImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImageResponse) SetBody(v *DeleteImageResponseBody) *DeleteImageResponse {
	s.Body = v
	return s
}

type DeleteJobTemplatesRequest struct {
	// The list of job templates. A maximum of 20 job templates can be deleted.
	//
	// Format: `[{"Id": "0.sched****"},{"Id": "1.sched****"}]`. Separate multiple job templates with commas (,).
	//
	// You can call the [ListJobTemplates](~~87248~~) operation to obtain the job template ID.
	Templates *string `json:"Templates,omitempty" xml:"Templates,omitempty"`
}

func (s DeleteJobTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobTemplatesRequest) SetTemplates(v string) *DeleteJobTemplatesRequest {
	s.Templates = &v
	return s
}

type DeleteJobTemplatesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobTemplatesResponseBody) SetRequestId(v string) *DeleteJobTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteJobTemplatesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteJobTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteJobTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobTemplatesResponse) SetHeaders(v map[string]*string) *DeleteJobTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobTemplatesResponse) SetStatusCode(v int32) *DeleteJobTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobTemplatesResponse) SetBody(v *DeleteJobTemplatesResponseBody) *DeleteJobTemplatesResponse {
	s.Body = v
	return s
}

type DeleteJobsRequest struct {
	// Specifies whether to use an asynchronous link to delete the jobs.
	//
	// Default value: false
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The list of jobs that you want to delete. Maximum number of jobs: 100. Minimum number of jobs: 1.
	//
	// Format: `[{"Id": "0.sched****"},{"Id": "1.sched****"}]`. Separate multiple jobs with commas (,).
	//
	// You can call the [ListJobs](~~87251~~) operation to query the job ID.
	Jobs *string `json:"Jobs,omitempty" xml:"Jobs,omitempty"`
}

func (s DeleteJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobsRequest) SetAsync(v bool) *DeleteJobsRequest {
	s.Async = &v
	return s
}

func (s *DeleteJobsRequest) SetClusterId(v string) *DeleteJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteJobsRequest) SetJobs(v string) *DeleteJobsRequest {
	s.Jobs = &v
	return s
}

type DeleteJobsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobsResponseBody) SetRequestId(v string) *DeleteJobsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteJobsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobsResponse) SetHeaders(v map[string]*string) *DeleteJobsResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobsResponse) SetStatusCode(v int32) *DeleteJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobsResponse) SetBody(v *DeleteJobsResponseBody) *DeleteJobsResponse {
	s.Body = v
	return s
}

type DeleteLocalImageRequest struct {
	// The ID of the cluster from which that you want to delete the image.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the image. Set the value to singularity.
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
	// The name of the image that you want to delete.
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
}

func (s DeleteLocalImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLocalImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteLocalImageRequest) SetClusterId(v string) *DeleteLocalImageRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteLocalImageRequest) SetContainerType(v string) *DeleteLocalImageRequest {
	s.ContainerType = &v
	return s
}

func (s *DeleteLocalImageRequest) SetImageName(v string) *DeleteLocalImageRequest {
	s.ImageName = &v
	return s
}

type DeleteLocalImageResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLocalImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLocalImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLocalImageResponseBody) SetRequestId(v string) *DeleteLocalImageResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLocalImageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLocalImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLocalImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLocalImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteLocalImageResponse) SetHeaders(v map[string]*string) *DeleteLocalImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteLocalImageResponse) SetStatusCode(v int32) *DeleteLocalImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLocalImageResponse) SetBody(v *DeleteLocalImageResponseBody) *DeleteLocalImageResponse {
	s.Body = v
	return s
}

type DeleteNodesRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string                       `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Instance  []*DeleteNodesRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
	// Specifies whether to release the instances that are created by using E-HPC.
	//
	// Default value: true
	ReleaseInstance *bool `json:"ReleaseInstance,omitempty" xml:"ReleaseInstance,omitempty"`
	// Specifies whether to directly delete the node. Valid values:
	//
	// *   true
	// *   false
	Sync *bool `json:"Sync,omitempty" xml:"Sync,omitempty"`
}

func (s DeleteNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodesRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodesRequest) SetClusterId(v string) *DeleteNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteNodesRequest) SetInstance(v []*DeleteNodesRequestInstance) *DeleteNodesRequest {
	s.Instance = v
	return s
}

func (s *DeleteNodesRequest) SetReleaseInstance(v bool) *DeleteNodesRequest {
	s.ReleaseInstance = &v
	return s
}

func (s *DeleteNodesRequest) SetSync(v bool) *DeleteNodesRequest {
	s.Sync = &v
	return s
}

type DeleteNodesRequestInstance struct {
	// The ID of the compute node that you want to delete. Valid values of N: 1 to 100.
	//
	// You can call the [DescribeCluster](~~87126~~) operation to query the IDs of the nodes in the cluster.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteNodesRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodesRequestInstance) GoString() string {
	return s.String()
}

func (s *DeleteNodesRequestInstance) SetId(v string) *DeleteNodesRequestInstance {
	s.Id = &v
	return s
}

type DeleteNodesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// *   If you set the Sync parameter to true, the DeleteNodes operation is synchronous. Valid value: Not Available.
	// *   If you set the Sync parameter to false, the DeleteNodes operation is asynchronous. You can call the [ListTasks](~~268225~~) operation to query the result of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodesResponseBody) SetRequestId(v string) *DeleteNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNodesResponseBody) SetTaskId(v string) *DeleteNodesResponseBody {
	s.TaskId = &v
	return s
}

type DeleteNodesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodesResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodesResponse) SetHeaders(v map[string]*string) *DeleteNodesResponse {
	s.Headers = v
	return s
}

func (s *DeleteNodesResponse) SetStatusCode(v int32) *DeleteNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNodesResponse) SetBody(v *DeleteNodesResponseBody) *DeleteNodesResponse {
	s.Body = v
	return s
}

type DeleteQueueRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the queue that you want to delete.
	//
	// You can call the [ListQueues](~~92176~~) operation to query the name of the queue.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s DeleteQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueRequest) GoString() string {
	return s.String()
}

func (s *DeleteQueueRequest) SetClusterId(v string) *DeleteQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteQueueRequest) SetQueueName(v string) *DeleteQueueRequest {
	s.QueueName = &v
	return s
}

type DeleteQueueResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponseBody) SetRequestId(v string) *DeleteQueueResponseBody {
	s.RequestId = &v
	return s
}

type DeleteQueueResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueResponse) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponse) SetHeaders(v map[string]*string) *DeleteQueueResponse {
	s.Headers = v
	return s
}

func (s *DeleteQueueResponse) SetStatusCode(v int32) *DeleteQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQueueResponse) SetBody(v *DeleteQueueResponseBody) *DeleteQueueResponse {
	s.Body = v
	return s
}

type DeleteSecurityGroupRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the security group.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DeleteSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupRequest) SetClusterId(v string) *DeleteSecurityGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteSecurityGroupRequest) SetSecurityGroupId(v string) *DeleteSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

type DeleteSecurityGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecurityGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupResponseBody) SetRequestId(v string) *DeleteSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSecurityGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupResponse) SetHeaders(v map[string]*string) *DeleteSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityGroupResponse) SetStatusCode(v int32) *DeleteSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecurityGroupResponse) SetBody(v *DeleteSecurityGroupResponseBody) *DeleteSecurityGroupResponse {
	s.Body = v
	return s
}

type DeleteUsersRequest struct {
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string                   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	User      []*DeleteUsersRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s DeleteUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersRequest) GoString() string {
	return s.String()
}

func (s *DeleteUsersRequest) SetAsync(v bool) *DeleteUsersRequest {
	s.Async = &v
	return s
}

func (s *DeleteUsersRequest) SetClusterId(v string) *DeleteUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteUsersRequest) SetUser(v []*DeleteUsersRequestUser) *DeleteUsersRequest {
	s.User = v
	return s
}

type DeleteUsersRequestUser struct {
	// The name of the user that you want to delete. Valid values of N: 1 to 100.
	//
	// You can call the [ListUsers](~~188572~~) operation to query the users of the cluster.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteUsersRequestUser) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersRequestUser) GoString() string {
	return s.String()
}

func (s *DeleteUsersRequestUser) SetName(v string) *DeleteUsersRequestUser {
	s.Name = &v
	return s
}

type DeleteUsersResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUsersResponseBody) SetRequestId(v string) *DeleteUsersResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersResponse) GoString() string {
	return s.String()
}

func (s *DeleteUsersResponse) SetHeaders(v map[string]*string) *DeleteUsersResponse {
	s.Headers = v
	return s
}

func (s *DeleteUsersResponse) SetStatusCode(v int32) *DeleteUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUsersResponse) SetBody(v *DeleteUsersResponseBody) *DeleteUsersResponse {
	s.Body = v
	return s
}

type DescribeAutoScaleConfigRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeAutoScaleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoScaleConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoScaleConfigRequest) SetClusterId(v string) *DescribeAutoScaleConfigRequest {
	s.ClusterId = &v
	return s
}

type DescribeAutoScaleConfigResponseBody struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the scheduler. Valid values:
	//
	// *   pbs
	// *   slurm
	// *   opengridscheduler
	// *   deadline
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// Indicates whether auto scale-out is enabled.
	EnableAutoGrow *bool `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	// Indicates whether auto scale-in is enabled.
	EnableAutoShrink *bool `json:"EnableAutoShrink,omitempty" xml:"EnableAutoShrink,omitempty"`
	// The list of nodes on which auto scaling is not enabled.
	ExcludeNodes *string `json:"ExcludeNodes,omitempty" xml:"ExcludeNodes,omitempty"`
	// The ratio of added nodes to the original ones. Valid values: 0 to 100.
	ExtraNodesGrowRatio *int32 `json:"ExtraNodesGrowRatio,omitempty" xml:"ExtraNodesGrowRatio,omitempty"`
	// The scale-out interval. The interval at which the compute nodes were scaled out. Valid values: 2 to 10.
	GrowIntervalInMinutes *int32 `json:"GrowIntervalInMinutes,omitempty" xml:"GrowIntervalInMinutes,omitempty"`
	// The percentage of the added nodes. Valid values: 1 to 100.
	GrowRatio *int32 `json:"GrowRatio,omitempty" xml:"GrowRatio,omitempty"`
	// The timeout period before the node was started. Valid values: 10 to 60.
	GrowTimeoutInMinutes *int32 `json:"GrowTimeoutInMinutes,omitempty" xml:"GrowTimeoutInMinutes,omitempty"`
	// The maximum number of compute nodes in the cluster. This parameter indicates the largest number of nodes that can be added to the cluster.
	MaxNodesInCluster *int32 `json:"MaxNodesInCluster,omitempty" xml:"MaxNodesInCluster,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of consecutive idle times of a node during a scale-in check. Valid values: 2 to 5.
	ShrinkIdleTimes *int32 `json:"ShrinkIdleTimes,omitempty" xml:"ShrinkIdleTimes,omitempty"`
	// The scale-in interval. The interval at which the compute nodes were scaled in. Valid values: 2 to 10.
	ShrinkIntervalInMinutes *int32 `json:"ShrinkIntervalInMinutes,omitempty" xml:"ShrinkIntervalInMinutes,omitempty"`
	// The maximum hourly rate of the instance. The value is accurate to three decimal places. It takes effect only when SpotStrategy is set to SpotWithPriceLimit.
	SpotPriceLimit *string `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding policy for the compute nodes. Valid values:
	//
	// *   NoSpot: The instance is created as a regular pay-as-you-go instance.
	// *   SpotWithPriceLimit: The instance is a preemptible one with a user-defined maximum hourly rate.
	// *   SpotAsPriceGo: The instance is created as a pay-as-you-go instance that is automatically priced based on the Alibaba Cloud Marketplace.
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The ID of the user.
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeAutoScaleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoScaleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoScaleConfigResponseBody) SetClusterId(v string) *DescribeAutoScaleConfigResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetClusterType(v string) *DescribeAutoScaleConfigResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetEnableAutoGrow(v bool) *DescribeAutoScaleConfigResponseBody {
	s.EnableAutoGrow = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetEnableAutoShrink(v bool) *DescribeAutoScaleConfigResponseBody {
	s.EnableAutoShrink = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetExcludeNodes(v string) *DescribeAutoScaleConfigResponseBody {
	s.ExcludeNodes = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetExtraNodesGrowRatio(v int32) *DescribeAutoScaleConfigResponseBody {
	s.ExtraNodesGrowRatio = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetGrowIntervalInMinutes(v int32) *DescribeAutoScaleConfigResponseBody {
	s.GrowIntervalInMinutes = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetGrowRatio(v int32) *DescribeAutoScaleConfigResponseBody {
	s.GrowRatio = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetGrowTimeoutInMinutes(v int32) *DescribeAutoScaleConfigResponseBody {
	s.GrowTimeoutInMinutes = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetMaxNodesInCluster(v int32) *DescribeAutoScaleConfigResponseBody {
	s.MaxNodesInCluster = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetRequestId(v string) *DescribeAutoScaleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetShrinkIdleTimes(v int32) *DescribeAutoScaleConfigResponseBody {
	s.ShrinkIdleTimes = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetShrinkIntervalInMinutes(v int32) *DescribeAutoScaleConfigResponseBody {
	s.ShrinkIntervalInMinutes = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetSpotPriceLimit(v string) *DescribeAutoScaleConfigResponseBody {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetSpotStrategy(v string) *DescribeAutoScaleConfigResponseBody {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetUid(v string) *DescribeAutoScaleConfigResponseBody {
	s.Uid = &v
	return s
}

type DescribeAutoScaleConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAutoScaleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutoScaleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoScaleConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoScaleConfigResponse) SetHeaders(v map[string]*string) *DescribeAutoScaleConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoScaleConfigResponse) SetStatusCode(v int32) *DescribeAutoScaleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoScaleConfigResponse) SetBody(v *DescribeAutoScaleConfigResponseBody) *DescribeAutoScaleConfigResponse {
	s.Body = v
	return s
}

type DescribeClusterRequest struct {
	// The ID of the cluster. You can call the [ListClusters](~~87116~~) operation to query the list of clusters in a region.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterRequest) SetClusterId(v string) *DescribeClusterRequest {
	s.ClusterId = &v
	return s
}

type DescribeClusterResponseBody struct {
	// The information about the cluster.
	ClusterInfo *DescribeClusterResponseBodyClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBody) SetClusterInfo(v *DescribeClusterResponseBodyClusterInfo) *DescribeClusterResponseBody {
	s.ClusterInfo = v
	return s
}

func (s *DescribeClusterResponseBody) SetRequestId(v string) *DescribeClusterResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClusterResponseBodyClusterInfo struct {
	// The service type of the domain account. Valid values:
	//
	// *   nis
	// *   ldap
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The array of the software in the cluster. The array contains the name and version of the software.
	Applications *DescribeClusterResponseBodyClusterInfoApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	// The image of the cluster.
	BaseOsTag *string `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	// The version of the E-HPC client.
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The time when the cluster was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The mode in which the cluster is deployed. Valid values:
	//
	// *   Standard: An account node, a scheduling node, a logon node, and multiple compute nodes are separately deployed.
	// *   Advanced: Two high availability (HA) account nodes, two HA scheduler nodes, one logon node, and multiple compute nodes are separately deployed.
	// *   Simple: A management node, a logon node, and multiple compute nodes are deployed. The management node consists of an account node and a scheduling node. The logon node and compute nodes are separately deployed.
	// *   Tiny: A management node and multiple compute nodes are deployed. The management node consists of an account node, a scheduling node, and a logon node. The compute nodes are separately deployed.
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// The description of the cluster.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The billing method of the nodes in the cluster. Valid values:
	//
	// *   PostPaid: pay-as-you-go
	// *   PrePaid: subscription
	EcsChargeType *string `json:"EcsChargeType,omitempty" xml:"EcsChargeType,omitempty"`
	// The list of ECS instance specifications and quantity.
	EcsInfo *DescribeClusterResponseBodyClusterInfoEcsInfo `json:"EcsInfo,omitempty" xml:"EcsInfo,omitempty" type:"Struct"`
	// Indicates whether the high availability feature is enabled.
	//
	// >  If high availability is enabled, a primary management node and a secondary management node are used.
	HaEnable *bool `json:"HaEnable,omitempty" xml:"HaEnable,omitempty"`
	// The ID of the Elastic Compute Service (ECS) instance.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The type of the image. Valid values:
	//
	// *   system: public image
	// *   self: custom image
	// *   others: shared image
	// *   marketplace: Alibaba Cloud Marketplace image
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The name of the AccessKey pair.
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The location where the cluster is deployed. Valid values:
	//
	// *   OnPremise: The cluster is deployed on a hybrid cloud.
	// *   PublicCloud: The node is deployed on a public cloud.
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the cluster.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of on-premises management nodes.
	//
	// This parameter is returned only when the cluster is deployed across hybrid environments and the hybrid-cloud proxy mode is enabled for the cluster.
	OnPremiseInfo *DescribeClusterResponseBodyClusterInfoOnPremiseInfo `json:"OnPremiseInfo,omitempty" xml:"OnPremiseInfo,omitempty" type:"Struct"`
	// The image tag of the operating system.
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	// The list of scripts downloaded after the cluster was created.
	PostInstallScripts *DescribeClusterResponseBodyClusterInfoPostInstallScripts `json:"PostInstallScripts,omitempty" xml:"PostInstallScripts,omitempty" type:"Struct"`
	RamNodeTypes       *string                                                   `json:"RamNodeTypes,omitempty" xml:"RamNodeTypes,omitempty"`
	RamRoleName        *string                                                   `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the security group.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remote directory on which the file system is mounted.
	RemoteDirectory *string `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	// The ID of the Super Computing Cluster (SCC) instance. If the cluster is not an SCC instance, a null string is returned.
	SccClusterId *string `json:"SccClusterId,omitempty" xml:"SccClusterId,omitempty"`
	// The type of the scheduler. Valid values:
	//
	// *   pbs
	// *   slurm
	// *   opengridscheduler
	// *   deadline
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	// The ID of the security group.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The status of the cluster. Valid values:
	//
	// *   uninit: The cluster is not initialized.
	// *   creating: The cluster is being created.
	// *   init: The cluster is being initialized.
	// *   running: The cluster is running.
	// *   exception: The cluster encounters an exception.
	// *   releasing: The cluster is being released.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the vSwitch. E-HPC can be deployed only in VPCs.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the Apsara File Storage NAS file system. NAS file systems cannot be automatically created.
	VolumeId *string `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	// The mount target of the file system. Mount targets cannot be automatically created for NAS file systems.
	VolumeMountpoint *string `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	// The type of the protocol that is used by the file system. Valid values:
	//
	// *   nfs
	// *   smb
	VolumeProtocol *string `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	// The type of the network shared storage. Valid value: NAS.
	VolumeType *string `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfo) SetAccountType(v string) *DescribeClusterResponseBodyClusterInfo {
	s.AccountType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetApplications(v *DescribeClusterResponseBodyClusterInfoApplications) *DescribeClusterResponseBodyClusterInfo {
	s.Applications = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetBaseOsTag(v string) *DescribeClusterResponseBodyClusterInfo {
	s.BaseOsTag = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetClientVersion(v string) *DescribeClusterResponseBodyClusterInfo {
	s.ClientVersion = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetCreateTime(v string) *DescribeClusterResponseBodyClusterInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetDeployMode(v string) *DescribeClusterResponseBodyClusterInfo {
	s.DeployMode = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetDescription(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Description = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetEcsChargeType(v string) *DescribeClusterResponseBodyClusterInfo {
	s.EcsChargeType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetEcsInfo(v *DescribeClusterResponseBodyClusterInfoEcsInfo) *DescribeClusterResponseBodyClusterInfo {
	s.EcsInfo = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetHaEnable(v bool) *DescribeClusterResponseBodyClusterInfo {
	s.HaEnable = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Id = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetImageId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.ImageId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetImageName(v string) *DescribeClusterResponseBodyClusterInfo {
	s.ImageName = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetImageOwnerAlias(v string) *DescribeClusterResponseBodyClusterInfo {
	s.ImageOwnerAlias = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetKeyPairName(v string) *DescribeClusterResponseBodyClusterInfo {
	s.KeyPairName = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetLocation(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Location = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetName(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Name = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetOnPremiseInfo(v *DescribeClusterResponseBodyClusterInfoOnPremiseInfo) *DescribeClusterResponseBodyClusterInfo {
	s.OnPremiseInfo = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetOsTag(v string) *DescribeClusterResponseBodyClusterInfo {
	s.OsTag = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetPostInstallScripts(v *DescribeClusterResponseBodyClusterInfoPostInstallScripts) *DescribeClusterResponseBodyClusterInfo {
	s.PostInstallScripts = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetRamNodeTypes(v string) *DescribeClusterResponseBodyClusterInfo {
	s.RamNodeTypes = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetRamRoleName(v string) *DescribeClusterResponseBodyClusterInfo {
	s.RamRoleName = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetRegionId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetRemoteDirectory(v string) *DescribeClusterResponseBodyClusterInfo {
	s.RemoteDirectory = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetSccClusterId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.SccClusterId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetSchedulerType(v string) *DescribeClusterResponseBodyClusterInfo {
	s.SchedulerType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetSecurityGroupId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetStatus(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Status = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVSwitchId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVolumeId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VolumeId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVolumeMountpoint(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VolumeMountpoint = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVolumeProtocol(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VolumeProtocol = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVolumeType(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VolumeType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVpcId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VpcId = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoApplications struct {
	ApplicationInfo []*DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo `json:"ApplicationInfo,omitempty" xml:"ApplicationInfo,omitempty" type:"Repeated"`
}

func (s DescribeClusterResponseBodyClusterInfoApplications) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoApplications) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoApplications) SetApplicationInfo(v []*DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) *DescribeClusterResponseBodyClusterInfoApplications {
	s.ApplicationInfo = v
	return s
}

type DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo struct {
	// The name of the software.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tag of the software.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The version of the software.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) SetName(v string) *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo {
	s.Name = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) SetTag(v string) *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo {
	s.Tag = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) SetVersion(v string) *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo {
	s.Version = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfo struct {
	// The list of compute nodes.
	Compute *DescribeClusterResponseBodyClusterInfoEcsInfoCompute `json:"Compute,omitempty" xml:"Compute,omitempty" type:"Struct"`
	// The list of logon nodes.
	Login *DescribeClusterResponseBodyClusterInfoEcsInfoLogin `json:"Login,omitempty" xml:"Login,omitempty" type:"Struct"`
	// The list of management nodes.
	Manager *DescribeClusterResponseBodyClusterInfoEcsInfoManager `json:"Manager,omitempty" xml:"Manager,omitempty" type:"Struct"`
	// The list of proxy nodes on the cloud.
	//
	// This parameter is returned only when the cluster is deployed across hybrid environments and the hybrid-cloud proxy mode is enabled for the cluster.
	ProxyMgr *DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr `json:"ProxyMgr,omitempty" xml:"ProxyMgr,omitempty" type:"Struct"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfo) SetCompute(v *DescribeClusterResponseBodyClusterInfoEcsInfoCompute) *DescribeClusterResponseBodyClusterInfoEcsInfo {
	s.Compute = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfo) SetLogin(v *DescribeClusterResponseBodyClusterInfoEcsInfoLogin) *DescribeClusterResponseBodyClusterInfoEcsInfo {
	s.Login = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfo) SetManager(v *DescribeClusterResponseBodyClusterInfoEcsInfoManager) *DescribeClusterResponseBodyClusterInfoEcsInfo {
	s.Manager = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfo) SetProxyMgr(v *DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr) *DescribeClusterResponseBodyClusterInfoEcsInfo {
	s.ProxyMgr = v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfoCompute struct {
	// The number of compute nodes.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The instance type of the compute nodes.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoCompute) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoCompute) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoCompute) SetCount(v int32) *DescribeClusterResponseBodyClusterInfoEcsInfoCompute {
	s.Count = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoCompute) SetInstanceType(v string) *DescribeClusterResponseBodyClusterInfoEcsInfoCompute {
	s.InstanceType = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfoLogin struct {
	// The number of logon nodes.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The instance type of the logon nodes.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoLogin) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoLogin) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoLogin) SetCount(v int32) *DescribeClusterResponseBodyClusterInfoEcsInfoLogin {
	s.Count = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoLogin) SetInstanceType(v string) *DescribeClusterResponseBodyClusterInfoEcsInfoLogin {
	s.InstanceType = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfoManager struct {
	// The number of management nodes.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The instance type of the management nodes.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoManager) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoManager) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoManager) SetCount(v int32) *DescribeClusterResponseBodyClusterInfoEcsInfoManager {
	s.Count = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoManager) SetInstanceType(v string) *DescribeClusterResponseBodyClusterInfoEcsInfoManager {
	s.InstanceType = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr struct {
	// The number of proxy nodes.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The instance type of the proxy node.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr) SetCount(v int32) *DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr {
	s.Count = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr) SetInstanceType(v string) *DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr {
	s.InstanceType = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoOnPremiseInfo struct {
	OnPremiseInfo []*DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo `json:"OnPremiseInfo,omitempty" xml:"OnPremiseInfo,omitempty" type:"Repeated"`
}

func (s DescribeClusterResponseBodyClusterInfoOnPremiseInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoOnPremiseInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoOnPremiseInfo) SetOnPremiseInfo(v []*DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo) *DescribeClusterResponseBodyClusterInfoOnPremiseInfo {
	s.OnPremiseInfo = v
	return s
}

type DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo struct {
	// The hostname of the on-premises management nodes.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The IP address of the on-premises management nodes.
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The type of on-premises management nodes. Valid values:
	//
	// - scheduler
	// - account
	// - account, scheduler
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo) SetHostName(v string) *DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo {
	s.HostName = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo) SetIP(v string) *DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo {
	s.IP = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo) SetType(v string) *DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo {
	s.Type = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoPostInstallScripts struct {
	PostInstallScriptInfo []*DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo `json:"PostInstallScriptInfo,omitempty" xml:"PostInstallScriptInfo,omitempty" type:"Repeated"`
}

func (s DescribeClusterResponseBodyClusterInfoPostInstallScripts) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoPostInstallScripts) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoPostInstallScripts) SetPostInstallScriptInfo(v []*DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo) *DescribeClusterResponseBodyClusterInfoPostInstallScripts {
	s.PostInstallScriptInfo = v
	return s
}

type DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo struct {
	// The runtime parameter of the script.
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The URL that was used to download the script.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo) SetArgs(v string) *DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo {
	s.Args = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo) SetUrl(v string) *DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo {
	s.Url = &v
	return s
}

type DescribeClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponse) SetHeaders(v map[string]*string) *DescribeClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterResponse) SetStatusCode(v int32) *DescribeClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterResponse) SetBody(v *DescribeClusterResponseBody) *DescribeClusterResponse {
	s.Body = v
	return s
}

type DescribeContainerAppRequest struct {
	// The ID of the containerized application.
	//
	// You can call the [ListContainerApps](~~87333~~) operation to query the ID of the containerized application.
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
}

func (s DescribeContainerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerAppRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerAppRequest) SetContainerId(v string) *DescribeContainerAppRequest {
	s.ContainerId = &v
	return s
}

type DescribeContainerAppResponseBody struct {
	// The information of the containerized application.
	ContainerAppInfo *DescribeContainerAppResponseBodyContainerAppInfo `json:"ContainerAppInfo,omitempty" xml:"ContainerAppInfo,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerAppResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerAppResponseBody) SetContainerAppInfo(v *DescribeContainerAppResponseBodyContainerAppInfo) *DescribeContainerAppResponseBody {
	s.ContainerAppInfo = v
	return s
}

func (s *DescribeContainerAppResponseBody) SetRequestId(v string) *DescribeContainerAppResponseBody {
	s.RequestId = &v
	return s
}

type DescribeContainerAppResponseBodyContainerAppInfo struct {
	// The time when the containerized application was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the containerized application.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the containerized application.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The tags of the image.
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The name of the containerized application.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the repository.
	Repository *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	// The type of the container. Set the value to singularity.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeContainerAppResponseBodyContainerAppInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerAppResponseBodyContainerAppInfo) GoString() string {
	return s.String()
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetCreateTime(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetDescription(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.Description = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetId(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.Id = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetImageTag(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.ImageTag = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetName(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.Name = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetRepository(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.Repository = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetType(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.Type = &v
	return s
}

type DescribeContainerAppResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeContainerAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerAppResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerAppResponse) SetHeaders(v map[string]*string) *DescribeContainerAppResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerAppResponse) SetStatusCode(v int32) *DescribeContainerAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerAppResponse) SetBody(v *DescribeContainerAppResponseBody) *DescribeContainerAppResponse {
	s.Body = v
	return s
}

type DescribeEstackImageRequest struct {
	// The number of the page to return.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeEstackImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEstackImageRequest) GoString() string {
	return s.String()
}

func (s *DescribeEstackImageRequest) SetPageNumber(v int32) *DescribeEstackImageRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEstackImageRequest) SetPageSize(v int32) *DescribeEstackImageRequest {
	s.PageSize = &v
	return s
}

type DescribeEstackImageResponseBody struct {
	// The array of base images.
	ImageList *DescribeEstackImageResponseBodyImageList `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of images.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEstackImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEstackImageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEstackImageResponseBody) SetImageList(v *DescribeEstackImageResponseBodyImageList) *DescribeEstackImageResponseBody {
	s.ImageList = v
	return s
}

func (s *DescribeEstackImageResponseBody) SetPageNumber(v int32) *DescribeEstackImageResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEstackImageResponseBody) SetPageSize(v int32) *DescribeEstackImageResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEstackImageResponseBody) SetRequestId(v string) *DescribeEstackImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEstackImageResponseBody) SetTotalCount(v int32) *DescribeEstackImageResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeEstackImageResponseBodyImageList struct {
	ImageListInfo []*DescribeEstackImageResponseBodyImageListImageListInfo `json:"ImageListInfo,omitempty" xml:"ImageListInfo,omitempty" type:"Repeated"`
}

func (s DescribeEstackImageResponseBodyImageList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEstackImageResponseBodyImageList) GoString() string {
	return s.String()
}

func (s *DescribeEstackImageResponseBodyImageList) SetImageListInfo(v []*DescribeEstackImageResponseBodyImageListImageListInfo) *DescribeEstackImageResponseBodyImageList {
	s.ImageListInfo = v
	return s
}

type DescribeEstackImageResponseBodyImageListImageListInfo struct {
	// The name of the image.
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The size of the image.
	ImageSize *int32 `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The type of the image.
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The download URL of the image.
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The time when the image was last modified.
	RecentUpdateTime *string `json:"RecentUpdateTime,omitempty" xml:"RecentUpdateTime,omitempty"`
}

func (s DescribeEstackImageResponseBodyImageListImageListInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeEstackImageResponseBodyImageListImageListInfo) GoString() string {
	return s.String()
}

func (s *DescribeEstackImageResponseBodyImageListImageListInfo) SetImageName(v string) *DescribeEstackImageResponseBodyImageListImageListInfo {
	s.ImageName = &v
	return s
}

func (s *DescribeEstackImageResponseBodyImageListImageListInfo) SetImageSize(v int32) *DescribeEstackImageResponseBodyImageListImageListInfo {
	s.ImageSize = &v
	return s
}

func (s *DescribeEstackImageResponseBodyImageListImageListInfo) SetImageType(v string) *DescribeEstackImageResponseBodyImageListImageListInfo {
	s.ImageType = &v
	return s
}

func (s *DescribeEstackImageResponseBodyImageListImageListInfo) SetImageUrl(v string) *DescribeEstackImageResponseBodyImageListImageListInfo {
	s.ImageUrl = &v
	return s
}

func (s *DescribeEstackImageResponseBodyImageListImageListInfo) SetRecentUpdateTime(v string) *DescribeEstackImageResponseBodyImageListImageListInfo {
	s.RecentUpdateTime = &v
	return s
}

type DescribeEstackImageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEstackImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEstackImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEstackImageResponse) GoString() string {
	return s.String()
}

func (s *DescribeEstackImageResponse) SetHeaders(v map[string]*string) *DescribeEstackImageResponse {
	s.Headers = v
	return s
}

func (s *DescribeEstackImageResponse) SetStatusCode(v int32) *DescribeEstackImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEstackImageResponse) SetBody(v *DescribeEstackImageResponseBody) *DescribeEstackImageResponse {
	s.Body = v
	return s
}

type DescribeGWSClusterPolicyRequest struct {
	AsyncMode *bool   `json:"AsyncMode,omitempty" xml:"AsyncMode,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeGWSClusterPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClusterPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeGWSClusterPolicyRequest) SetAsyncMode(v bool) *DescribeGWSClusterPolicyRequest {
	s.AsyncMode = &v
	return s
}

func (s *DescribeGWSClusterPolicyRequest) SetClusterId(v string) *DescribeGWSClusterPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeGWSClusterPolicyRequest) SetTaskId(v string) *DescribeGWSClusterPolicyRequest {
	s.TaskId = &v
	return s
}

type DescribeGWSClusterPolicyResponseBody struct {
	Clipboard   *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	LocalDrive  *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	Watermark   *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s DescribeGWSClusterPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClusterPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGWSClusterPolicyResponseBody) SetClipboard(v string) *DescribeGWSClusterPolicyResponseBody {
	s.Clipboard = &v
	return s
}

func (s *DescribeGWSClusterPolicyResponseBody) SetLocalDrive(v string) *DescribeGWSClusterPolicyResponseBody {
	s.LocalDrive = &v
	return s
}

func (s *DescribeGWSClusterPolicyResponseBody) SetRequestId(v string) *DescribeGWSClusterPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGWSClusterPolicyResponseBody) SetUsbRedirect(v string) *DescribeGWSClusterPolicyResponseBody {
	s.UsbRedirect = &v
	return s
}

func (s *DescribeGWSClusterPolicyResponseBody) SetWatermark(v string) *DescribeGWSClusterPolicyResponseBody {
	s.Watermark = &v
	return s
}

type DescribeGWSClusterPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGWSClusterPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGWSClusterPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClusterPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeGWSClusterPolicyResponse) SetHeaders(v map[string]*string) *DescribeGWSClusterPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeGWSClusterPolicyResponse) SetStatusCode(v int32) *DescribeGWSClusterPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGWSClusterPolicyResponse) SetBody(v *DescribeGWSClusterPolicyResponseBody) *DescribeGWSClusterPolicyResponse {
	s.Body = v
	return s
}

type DescribeGWSClustersRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGWSClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeGWSClustersRequest) SetClusterId(v string) *DescribeGWSClustersRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeGWSClustersRequest) SetPageNumber(v int32) *DescribeGWSClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGWSClustersRequest) SetPageSize(v int32) *DescribeGWSClustersRequest {
	s.PageSize = &v
	return s
}

type DescribeGWSClustersResponseBody struct {
	CallerType *string                                  `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
	Clusters   *DescribeGWSClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Struct"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGWSClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGWSClustersResponseBody) SetCallerType(v string) *DescribeGWSClustersResponseBody {
	s.CallerType = &v
	return s
}

func (s *DescribeGWSClustersResponseBody) SetClusters(v *DescribeGWSClustersResponseBodyClusters) *DescribeGWSClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeGWSClustersResponseBody) SetPageNumber(v int32) *DescribeGWSClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGWSClustersResponseBody) SetPageSize(v int32) *DescribeGWSClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGWSClustersResponseBody) SetRequestId(v string) *DescribeGWSClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGWSClustersResponseBody) SetTotalCount(v int32) *DescribeGWSClustersResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeGWSClustersResponseBodyClusters struct {
	ClusterInfo []*DescribeGWSClustersResponseBodyClustersClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Repeated"`
}

func (s DescribeGWSClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeGWSClustersResponseBodyClusters) SetClusterInfo(v []*DescribeGWSClustersResponseBodyClustersClusterInfo) *DescribeGWSClustersResponseBodyClusters {
	s.ClusterInfo = v
	return s
}

type DescribeGWSClustersResponseBodyClustersClusterInfo struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceCount *int32  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// VPC ID。
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeGWSClustersResponseBodyClustersClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClustersResponseBodyClustersClusterInfo) GoString() string {
	return s.String()
}

func (s *DescribeGWSClustersResponseBodyClustersClusterInfo) SetClusterId(v string) *DescribeGWSClustersResponseBodyClustersClusterInfo {
	s.ClusterId = &v
	return s
}

func (s *DescribeGWSClustersResponseBodyClustersClusterInfo) SetCreateTime(v string) *DescribeGWSClustersResponseBodyClustersClusterInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeGWSClustersResponseBodyClustersClusterInfo) SetInstanceCount(v int32) *DescribeGWSClustersResponseBodyClustersClusterInfo {
	s.InstanceCount = &v
	return s
}

func (s *DescribeGWSClustersResponseBodyClustersClusterInfo) SetStatus(v string) *DescribeGWSClustersResponseBodyClustersClusterInfo {
	s.Status = &v
	return s
}

func (s *DescribeGWSClustersResponseBodyClustersClusterInfo) SetVpcId(v string) *DescribeGWSClustersResponseBodyClustersClusterInfo {
	s.VpcId = &v
	return s
}

type DescribeGWSClustersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGWSClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGWSClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeGWSClustersResponse) SetHeaders(v map[string]*string) *DescribeGWSClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeGWSClustersResponse) SetStatusCode(v int32) *DescribeGWSClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGWSClustersResponse) SetBody(v *DescribeGWSClustersResponseBody) *DescribeGWSClustersResponse {
	s.Body = v
	return s
}

type DescribeGWSImagesRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGWSImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGWSImagesRequest) SetPageNumber(v int32) *DescribeGWSImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGWSImagesRequest) SetPageSize(v int32) *DescribeGWSImagesRequest {
	s.PageSize = &v
	return s
}

type DescribeGWSImagesResponseBody struct {
	Images     *DescribeGWSImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGWSImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGWSImagesResponseBody) SetImages(v *DescribeGWSImagesResponseBodyImages) *DescribeGWSImagesResponseBody {
	s.Images = v
	return s
}

func (s *DescribeGWSImagesResponseBody) SetPageNumber(v int32) *DescribeGWSImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGWSImagesResponseBody) SetPageSize(v int32) *DescribeGWSImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGWSImagesResponseBody) SetRequestId(v string) *DescribeGWSImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGWSImagesResponseBody) SetTotalCount(v int32) *DescribeGWSImagesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeGWSImagesResponseBodyImages struct {
	ImageInfo []*DescribeGWSImagesResponseBodyImagesImageInfo `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" type:"Repeated"`
}

func (s DescribeGWSImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeGWSImagesResponseBodyImages) SetImageInfo(v []*DescribeGWSImagesResponseBodyImagesImageInfo) *DescribeGWSImagesResponseBodyImages {
	s.ImageInfo = v
	return s
}

type DescribeGWSImagesResponseBodyImagesImageInfo struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ImageId    *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageType  *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Progress   *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGWSImagesResponseBodyImagesImageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSImagesResponseBodyImagesImageInfo) GoString() string {
	return s.String()
}

func (s *DescribeGWSImagesResponseBodyImagesImageInfo) SetCreateTime(v string) *DescribeGWSImagesResponseBodyImagesImageInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImagesImageInfo) SetImageId(v string) *DescribeGWSImagesResponseBodyImagesImageInfo {
	s.ImageId = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImagesImageInfo) SetImageType(v string) *DescribeGWSImagesResponseBodyImagesImageInfo {
	s.ImageType = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImagesImageInfo) SetName(v string) *DescribeGWSImagesResponseBodyImagesImageInfo {
	s.Name = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImagesImageInfo) SetProgress(v string) *DescribeGWSImagesResponseBodyImagesImageInfo {
	s.Progress = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImagesImageInfo) SetSize(v int32) *DescribeGWSImagesResponseBodyImagesImageInfo {
	s.Size = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImagesImageInfo) SetStatus(v string) *DescribeGWSImagesResponseBodyImagesImageInfo {
	s.Status = &v
	return s
}

type DescribeGWSImagesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGWSImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGWSImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGWSImagesResponse) SetHeaders(v map[string]*string) *DescribeGWSImagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGWSImagesResponse) SetStatusCode(v int32) *DescribeGWSImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGWSImagesResponse) SetBody(v *DescribeGWSImagesResponseBody) *DescribeGWSImagesResponse {
	s.Body = v
	return s
}

type DescribeGWSInstancesRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserUid    *int64  `json:"UserUid,omitempty" xml:"UserUid,omitempty"`
}

func (s DescribeGWSInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesRequest) SetClusterId(v string) *DescribeGWSInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeGWSInstancesRequest) SetInstanceId(v string) *DescribeGWSInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGWSInstancesRequest) SetPageNumber(v int32) *DescribeGWSInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGWSInstancesRequest) SetPageSize(v int32) *DescribeGWSInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGWSInstancesRequest) SetUserName(v string) *DescribeGWSInstancesRequest {
	s.UserName = &v
	return s
}

func (s *DescribeGWSInstancesRequest) SetUserUid(v int64) *DescribeGWSInstancesRequest {
	s.UserUid = &v
	return s
}

type DescribeGWSInstancesResponseBody struct {
	Instances  *DescribeGWSInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGWSInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesResponseBody) SetInstances(v *DescribeGWSInstancesResponseBodyInstances) *DescribeGWSInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeGWSInstancesResponseBody) SetPageNumber(v int32) *DescribeGWSInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGWSInstancesResponseBody) SetPageSize(v int32) *DescribeGWSInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGWSInstancesResponseBody) SetRequestId(v string) *DescribeGWSInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGWSInstancesResponseBody) SetTotalCount(v int32) *DescribeGWSInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeGWSInstancesResponseBodyInstances struct {
	InstanceInfo []*DescribeGWSInstancesResponseBodyInstancesInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Repeated"`
}

func (s DescribeGWSInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesResponseBodyInstances) SetInstanceInfo(v []*DescribeGWSInstancesResponseBodyInstancesInstanceInfo) *DescribeGWSInstancesResponseBodyInstances {
	s.InstanceInfo = v
	return s
}

type DescribeGWSInstancesResponseBodyInstancesInstanceInfo struct {
	AppList      *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Struct"`
	ClusterId    *string                                                       `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateTime   *string                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime   *string                                                       `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId   *string                                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string                                                       `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Name         *string                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	Status       *string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	UserName     *string                                                       `json:"UserName,omitempty" xml:"UserName,omitempty"`
	WorkMode     *string                                                       `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s DescribeGWSInstancesResponseBodyInstancesInstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesResponseBodyInstancesInstanceInfo) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetAppList(v *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppList) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.AppList = v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetClusterId(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.ClusterId = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetCreateTime(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetExpireTime(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.ExpireTime = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetInstanceId(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetInstanceType(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.InstanceType = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetName(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.Name = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetStatus(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.Status = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetUserName(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.UserName = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetWorkMode(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.WorkMode = &v
	return s
}

type DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppList struct {
	AppInfo []*DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo `json:"AppInfo,omitempty" xml:"AppInfo,omitempty" type:"Repeated"`
}

func (s DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppList) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppList) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppList) SetAppInfo(v []*DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo) *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppList {
	s.AppInfo = v
	return s
}

type DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo struct {
	AppArgs *string `json:"AppArgs,omitempty" xml:"AppArgs,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppPath *string `json:"AppPath,omitempty" xml:"AppPath,omitempty"`
}

func (s DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo) SetAppArgs(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo {
	s.AppArgs = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo) SetAppName(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo {
	s.AppName = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo) SetAppPath(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo {
	s.AppPath = &v
	return s
}

type DescribeGWSInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGWSInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGWSInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesResponse) SetHeaders(v map[string]*string) *DescribeGWSInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGWSInstancesResponse) SetStatusCode(v int32) *DescribeGWSInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGWSInstancesResponse) SetBody(v *DescribeGWSInstancesResponseBody) *DescribeGWSInstancesResponse {
	s.Body = v
	return s
}

type DescribeImageRequest struct {
	// The ID of the cluster that you want to manage.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the container. Set the value to singularity.
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
	// The tag of the image. Default value: latest.
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The name of the repository.
	Repository *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
}

func (s DescribeImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageRequest) SetClusterId(v string) *DescribeImageRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeImageRequest) SetContainerType(v string) *DescribeImageRequest {
	s.ContainerType = &v
	return s
}

func (s *DescribeImageRequest) SetImageTag(v string) *DescribeImageRequest {
	s.ImageTag = &v
	return s
}

func (s *DescribeImageRequest) SetRepository(v string) *DescribeImageRequest {
	s.Repository = &v
	return s
}

type DescribeImageResponseBody struct {
	// The information of the image.
	ImageInfo *DescribeImageResponseBodyImageInfo `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageResponseBody) SetImageInfo(v *DescribeImageResponseBodyImageInfo) *DescribeImageResponseBody {
	s.ImageInfo = v
	return s
}

func (s *DescribeImageResponseBody) SetRequestId(v string) *DescribeImageResponseBody {
	s.RequestId = &v
	return s
}

type DescribeImageResponseBodyImageInfo struct {
	// The ID of the image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the repository.
	Repository *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	// The status of the image.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The container system.
	System *string `json:"System,omitempty" xml:"System,omitempty"`
	// The tag of the image.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The type of the image. Valid values:
	//
	// *   shifter
	// *   docker
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the image was last updated.
	UpdateDateTime *string `json:"UpdateDateTime,omitempty" xml:"UpdateDateTime,omitempty"`
}

func (s DescribeImageResponseBodyImageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResponseBodyImageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageResponseBodyImageInfo) SetImageId(v string) *DescribeImageResponseBodyImageInfo {
	s.ImageId = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetRepository(v string) *DescribeImageResponseBodyImageInfo {
	s.Repository = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetStatus(v string) *DescribeImageResponseBodyImageInfo {
	s.Status = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetSystem(v string) *DescribeImageResponseBodyImageInfo {
	s.System = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetTag(v string) *DescribeImageResponseBodyImageInfo {
	s.Tag = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetType(v string) *DescribeImageResponseBodyImageInfo {
	s.Type = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetUpdateDateTime(v string) *DescribeImageResponseBodyImageInfo {
	s.UpdateDateTime = &v
	return s
}

type DescribeImageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageResponse) SetHeaders(v map[string]*string) *DescribeImageResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageResponse) SetStatusCode(v int32) *DescribeImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageResponse) SetBody(v *DescribeImageResponseBody) *DescribeImageResponse {
	s.Body = v
	return s
}

type DescribeImageGatewayConfigRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeImageGatewayConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGatewayConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageGatewayConfigRequest) SetClusterId(v string) *DescribeImageGatewayConfigRequest {
	s.ClusterId = &v
	return s
}

type DescribeImageGatewayConfigResponseBody struct {
	Imagegw   *DescribeImageGatewayConfigResponseBodyImagegw `json:"Imagegw,omitempty" xml:"Imagegw,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageGatewayConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGatewayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageGatewayConfigResponseBody) SetImagegw(v *DescribeImageGatewayConfigResponseBodyImagegw) *DescribeImageGatewayConfigResponseBody {
	s.Imagegw = v
	return s
}

func (s *DescribeImageGatewayConfigResponseBody) SetRequestId(v string) *DescribeImageGatewayConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeImageGatewayConfigResponseBodyImagegw struct {
	DefaultImageLocation   *string                                                 `json:"DefaultImageLocation,omitempty" xml:"DefaultImageLocation,omitempty"`
	ImageExpirationTimeout *string                                                 `json:"ImageExpirationTimeout,omitempty" xml:"ImageExpirationTimeout,omitempty"`
	Locations              *DescribeImageGatewayConfigResponseBodyImagegwLocations `json:"Locations,omitempty" xml:"Locations,omitempty" type:"Struct"`
	MongoDBURI             *string                                                 `json:"MongoDBURI,omitempty" xml:"MongoDBURI,omitempty"`
	PullUpdateTimeout      *int64                                                  `json:"PullUpdateTimeout,omitempty" xml:"PullUpdateTimeout,omitempty"`
	UpdateDateTime         *string                                                 `json:"UpdateDateTime,omitempty" xml:"UpdateDateTime,omitempty"`
}

func (s DescribeImageGatewayConfigResponseBodyImagegw) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGatewayConfigResponseBodyImagegw) GoString() string {
	return s.String()
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetDefaultImageLocation(v string) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.DefaultImageLocation = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetImageExpirationTimeout(v string) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.ImageExpirationTimeout = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetLocations(v *DescribeImageGatewayConfigResponseBodyImagegwLocations) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.Locations = v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetMongoDBURI(v string) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.MongoDBURI = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetPullUpdateTimeout(v int64) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.PullUpdateTimeout = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetUpdateDateTime(v string) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.UpdateDateTime = &v
	return s
}

type DescribeImageGatewayConfigResponseBodyImagegwLocations struct {
	LocationInfo []*DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo `json:"LocationInfo,omitempty" xml:"LocationInfo,omitempty" type:"Repeated"`
}

func (s DescribeImageGatewayConfigResponseBodyImagegwLocations) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGatewayConfigResponseBodyImagegwLocations) GoString() string {
	return s.String()
}

func (s *DescribeImageGatewayConfigResponseBodyImagegwLocations) SetLocationInfo(v []*DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo) *DescribeImageGatewayConfigResponseBodyImagegwLocations {
	s.LocationInfo = v
	return s
}

type DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo struct {
	Authentication *string `json:"Authentication,omitempty" xml:"Authentication,omitempty"`
	Location       *string `json:"Location,omitempty" xml:"Location,omitempty"`
	RemoteType     *string `json:"RemoteType,omitempty" xml:"RemoteType,omitempty"`
	URL            *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo) SetAuthentication(v string) *DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo {
	s.Authentication = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo) SetLocation(v string) *DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo {
	s.Location = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo) SetRemoteType(v string) *DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo {
	s.RemoteType = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo) SetURL(v string) *DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo {
	s.URL = &v
	return s
}

type DescribeImageGatewayConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeImageGatewayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImageGatewayConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGatewayConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageGatewayConfigResponse) SetHeaders(v map[string]*string) *DescribeImageGatewayConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageGatewayConfigResponse) SetStatusCode(v int32) *DescribeImageGatewayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageGatewayConfigResponse) SetBody(v *DescribeImageGatewayConfigResponseBody) *DescribeImageGatewayConfigResponse {
	s.Body = v
	return s
}

type DescribeImagePriceRequest struct {
	// The number of images that you want to purchase. Valid values: 1 to 1000.
	//
	// Default value: 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The ID of the image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The type of the order. The order can be set only as a purchase order. Valid value: INSTANCE-BUY.
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The subscription duration. Valid values:
	//
	// *   If PriceUnit is set to Day, the valid values of the Period parameter are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, and 30.
	// *   If PriceUnit is set to Month, the valid values of the Period parameter are 1, 2, 3, 4, 5, 6, 7, 8, and 9.
	// *   If PriceUnit is set to Year, the valid values of the Period parameter are 1, 2, and 3.
	//
	// Default value: 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// *   Day
	// *   Month
	// *   Year
	//
	// Default value: Day
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// The stock keeping unit (SKU) of the image. Valid value: package.
	SkuCode *string `json:"SkuCode,omitempty" xml:"SkuCode,omitempty"`
}

func (s DescribeImagePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeImagePriceRequest) SetAmount(v int32) *DescribeImagePriceRequest {
	s.Amount = &v
	return s
}

func (s *DescribeImagePriceRequest) SetImageId(v string) *DescribeImagePriceRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeImagePriceRequest) SetOrderType(v string) *DescribeImagePriceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeImagePriceRequest) SetPeriod(v int32) *DescribeImagePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribeImagePriceRequest) SetPriceUnit(v string) *DescribeImagePriceRequest {
	s.PriceUnit = &v
	return s
}

func (s *DescribeImagePriceRequest) SetSkuCode(v string) *DescribeImagePriceRequest {
	s.SkuCode = &v
	return s
}

type DescribeImagePriceResponseBody struct {
	// The number of images that you want to purchase.
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The discount that is applied.
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The ID of the image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The original price of the image.
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The final price of the image.
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeImagePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImagePriceResponseBody) SetAmount(v int32) *DescribeImagePriceResponseBody {
	s.Amount = &v
	return s
}

func (s *DescribeImagePriceResponseBody) SetDiscountPrice(v float32) *DescribeImagePriceResponseBody {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeImagePriceResponseBody) SetImageId(v string) *DescribeImagePriceResponseBody {
	s.ImageId = &v
	return s
}

func (s *DescribeImagePriceResponseBody) SetOriginalPrice(v float32) *DescribeImagePriceResponseBody {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeImagePriceResponseBody) SetRequestId(v string) *DescribeImagePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImagePriceResponseBody) SetTradePrice(v float32) *DescribeImagePriceResponseBody {
	s.TradePrice = &v
	return s
}

type DescribeImagePriceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeImagePriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImagePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeImagePriceResponse) SetHeaders(v map[string]*string) *DescribeImagePriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeImagePriceResponse) SetStatusCode(v int32) *DescribeImagePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImagePriceResponse) SetBody(v *DescribeImagePriceResponseBody) *DescribeImagePriceResponse {
	s.Body = v
	return s
}

type DescribeJobRequest struct {
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the job.
	//
	// You can call the [ListJobs](~~87251~~) operation to query the job ID.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DescribeJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobRequest) SetAsync(v bool) *DescribeJobRequest {
	s.Async = &v
	return s
}

func (s *DescribeJobRequest) SetClusterId(v string) *DescribeJobRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeJobRequest) SetJobId(v string) *DescribeJobRequest {
	s.JobId = &v
	return s
}

type DescribeJobResponseBody struct {
	// The list of returned job information.
	Message *DescribeJobResponseBodyMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBody) SetMessage(v *DescribeJobResponseBodyMessage) *DescribeJobResponseBody {
	s.Message = v
	return s
}

func (s *DescribeJobResponseBody) SetRequestId(v string) *DescribeJobResponseBody {
	s.RequestId = &v
	return s
}

type DescribeJobResponseBodyMessage struct {
	// The details of the job.
	JobInfo *string `json:"JobInfo,omitempty" xml:"JobInfo,omitempty"`
}

func (s DescribeJobResponseBodyMessage) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobResponseBodyMessage) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBodyMessage) SetJobInfo(v string) *DescribeJobResponseBodyMessage {
	s.JobInfo = &v
	return s
}

type DescribeJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobResponse) SetHeaders(v map[string]*string) *DescribeJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobResponse) SetStatusCode(v int32) *DescribeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobResponse) SetBody(v *DescribeJobResponseBody) *DescribeJobResponse {
	s.Body = v
	return s
}

type DescribeNFSClientStatusRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeNFSClientStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNFSClientStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeNFSClientStatusRequest) SetInstanceId(v string) *DescribeNFSClientStatusRequest {
	s.InstanceId = &v
	return s
}

type DescribeNFSClientStatusResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeNFSClientStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Status    *string                                    `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeNFSClientStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNFSClientStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNFSClientStatusResponseBody) SetRequestId(v string) *DescribeNFSClientStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNFSClientStatusResponseBody) SetResult(v *DescribeNFSClientStatusResponseBodyResult) *DescribeNFSClientStatusResponseBody {
	s.Result = v
	return s
}

func (s *DescribeNFSClientStatusResponseBody) SetStatus(v string) *DescribeNFSClientStatusResponseBody {
	s.Status = &v
	return s
}

type DescribeNFSClientStatusResponseBodyResult struct {
	ExitCode           *int32  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	Output             *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s DescribeNFSClientStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeNFSClientStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeNFSClientStatusResponseBodyResult) SetExitCode(v int32) *DescribeNFSClientStatusResponseBodyResult {
	s.ExitCode = &v
	return s
}

func (s *DescribeNFSClientStatusResponseBodyResult) SetInvokeRecordStatus(v string) *DescribeNFSClientStatusResponseBodyResult {
	s.InvokeRecordStatus = &v
	return s
}

func (s *DescribeNFSClientStatusResponseBodyResult) SetOutput(v string) *DescribeNFSClientStatusResponseBodyResult {
	s.Output = &v
	return s
}

type DescribeNFSClientStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNFSClientStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNFSClientStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNFSClientStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeNFSClientStatusResponse) SetHeaders(v map[string]*string) *DescribeNFSClientStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeNFSClientStatusResponse) SetStatusCode(v int32) *DescribeNFSClientStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNFSClientStatusResponse) SetBody(v *DescribeNFSClientStatusResponseBody) *DescribeNFSClientStatusResponse {
	s.Body = v
	return s
}

type DescribePriceRequest struct {
	// The billing method of the ECS instances. Valid values:
	//
	// *   PostPaid: pay-as-you-go
	// *   PrePaid: subscription
	//
	// Default value: PostPaid
	ChargeType  *string                            `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Commodities []*DescribePriceRequestCommodities `json:"Commodities,omitempty" xml:"Commodities,omitempty" type:"Repeated"`
	// The type of the order. The order can be set only as a purchase order. Valid value: INSTANCE-BUY.
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The billing cycle of the Elastic Compute Service (ECS) instances. This parameter takes effect only when the ChargeType parameter is set to PrePaid. Valid values:
	//
	// *   Month: pay-by-month
	// *   Year: pay-by-year
	// *   Hour: pay-by-hour
	//
	// Default value: Hour
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
}

func (s DescribePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) SetChargeType(v string) *DescribePriceRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribePriceRequest) SetCommodities(v []*DescribePriceRequestCommodities) *DescribePriceRequest {
	s.Commodities = v
	return s
}

func (s *DescribePriceRequest) SetOrderType(v string) *DescribePriceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribePriceRequest) SetPriceUnit(v string) *DescribePriceRequest {
	s.PriceUnit = &v
	return s
}

type DescribePriceRequestCommodities struct {
	// The node quantity of the type. Valid values: 1 to 1000.
	//
	// Default value: 1
	//
	// Valid values of N: 1 to 10
	Amount    *int32                                      `json:"Amount,omitempty" xml:"Amount,omitempty"`
	DataDisks []*DescribePriceRequestCommoditiesDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// The instance type of the node.
	//
	// Valid values of N: 1 to 10
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The EIP billing method of the node. Valid values:
	//
	// *   PayByBandwidth: pay-by-bandwidth
	// *   PayByTraffic: pay-by-traffic
	//
	// Valid values of N: 1 to 10
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum outbound public bandwidth of the node. Unit: Mbit/s.
	//
	// Valid values: 0 to 100
	//
	// Default value: 0
	//
	// Valid values of N: 1 to 10
	InternetMaxBandWidthOut *int32 `json:"InternetMaxBandWidthOut,omitempty" xml:"InternetMaxBandWidthOut,omitempty"`
	// The network type of the node. Valid value: VPC.
	//
	// Valid values of N: 1 to 10
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The type of the node. Valid values:
	//
	// *   Compute: compute node
	// *   Manager: management node
	// *   Login: logon node
	//
	// Valid values of N: 1 to 10
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The subscription duration of the node. Valid values:
	//
	// *   If PriceUnit is set to Year, the valid values of the Period parameter are 1, 2, and 3.
	// *   If PriceUnit is set to Month, the valid values of the Period parameter are 1, 2, 3, 4, 5, 6, 7, 8, and 9.
	// *   If PriceUnit is set to Hour, the valid value of the Period parameter is 1.
	//
	// Default value: 1
	//
	// Valid values of N: 1 to 10
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The system disk type of the node. Valid values:
	//
	// *   cloud_efficiency: ultra disk
	// *   cloud_ssd: SSD
	// *   cloud_essd: ESSD
	// *   cloud: basic disk
	//
	// Default value: cloud_efficiency
	//
	// Valid values of N: 1 to 10
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The performance level of the ESSD used as the system disk. This parameter takes effect only when the Commodities.N.SystemDiskCategory parameter is set to cloud_essd. Default value: PL1. Valid values:
	//
	// *   PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	// *   PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	// *   PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	// *   PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL1
	//
	// Valid values of N: 1 to 10
	SystemDiskPerformanceLevel *string `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	// The system disk size of the node. Unit: GB.
	//
	// Valid values: 40 to 500
	//
	// Default value: 40
	//
	// Valid values of N: 1 to 10
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribePriceRequestCommodities) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequestCommodities) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestCommodities) SetAmount(v int32) *DescribePriceRequestCommodities {
	s.Amount = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetDataDisks(v []*DescribePriceRequestCommoditiesDataDisks) *DescribePriceRequestCommodities {
	s.DataDisks = v
	return s
}

func (s *DescribePriceRequestCommodities) SetInstanceType(v string) *DescribePriceRequestCommodities {
	s.InstanceType = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetInternetChargeType(v string) *DescribePriceRequestCommodities {
	s.InternetChargeType = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetInternetMaxBandWidthOut(v int32) *DescribePriceRequestCommodities {
	s.InternetMaxBandWidthOut = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetNetworkType(v string) *DescribePriceRequestCommodities {
	s.NetworkType = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetNodeType(v string) *DescribePriceRequestCommodities {
	s.NodeType = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetPeriod(v int32) *DescribePriceRequestCommodities {
	s.Period = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetSystemDiskCategory(v string) *DescribePriceRequestCommodities {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetSystemDiskPerformanceLevel(v string) *DescribePriceRequestCommodities {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetSystemDiskSize(v int32) *DescribePriceRequestCommodities {
	s.SystemDiskSize = &v
	return s
}

type DescribePriceRequestCommoditiesDataDisks struct {
	// The type of the data disk. Valid values:
	//
	// *   cloud_efficiency: ultra disk
	// *   cloud_ssd: SSD
	// *   cloud_essd: ESSD
	// *   cloud: basic disk
	//
	// Default value: cloud_efficiency
	//
	// Valid values of N: 0 to 4
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// Specifies whether the data disk is released when the node is released. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: true
	//
	// Valid values of N: 0 to 4
	DeleteWithInstance *bool `json:"deleteWithInstance,omitempty" xml:"deleteWithInstance,omitempty"`
	// Specifies whether to encrypt the data disk. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false
	//
	// Valid values of N: 0 to 4
	Encrypted *bool `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// The performance level of the ESSD used as the data disk. This parameter takes effect only when the Commodities.N.DataDisks.N.category parameter is set to cloud_essd. Default value: PL1. Valid values:
	//
	// *   PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	// *   PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	// *   PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	// *   PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL1
	//
	// Valid values of N: 0 to 4
	PerformanceLevel *string `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
	// The size of the data disk. Unit: GB.
	//
	// Valid values: 40 to 500
	//
	// Default value: 40
	//
	// Valid values of N: 0 to 4
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s DescribePriceRequestCommoditiesDataDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequestCommoditiesDataDisks) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestCommoditiesDataDisks) SetCategory(v string) *DescribePriceRequestCommoditiesDataDisks {
	s.Category = &v
	return s
}

func (s *DescribePriceRequestCommoditiesDataDisks) SetDeleteWithInstance(v bool) *DescribePriceRequestCommoditiesDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *DescribePriceRequestCommoditiesDataDisks) SetEncrypted(v bool) *DescribePriceRequestCommoditiesDataDisks {
	s.Encrypted = &v
	return s
}

func (s *DescribePriceRequestCommoditiesDataDisks) SetPerformanceLevel(v string) *DescribePriceRequestCommoditiesDataDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribePriceRequestCommoditiesDataDisks) SetSize(v int32) *DescribePriceRequestCommoditiesDataDisks {
	s.Size = &v
	return s
}

type DescribePriceResponseBody struct {
	// The array of cluster prices. If you query the prices of multiple nodes in the cluster, the sequence of the prices in the returned value of PriceInfo is the same as that of the nodes in the request parameters. For example, the first price in the value of PriceInfo is the price of the first node specified in the request parameters.
	Prices *DescribePriceResponseBodyPrices `json:"Prices,omitempty" xml:"Prices,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total price.
	//
	// Unit: USD
	TotalTradePrice *float32 `json:"TotalTradePrice,omitempty" xml:"TotalTradePrice,omitempty"`
}

func (s DescribePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBody) SetPrices(v *DescribePriceResponseBodyPrices) *DescribePriceResponseBody {
	s.Prices = v
	return s
}

func (s *DescribePriceResponseBody) SetRequestId(v string) *DescribePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePriceResponseBody) SetTotalTradePrice(v float32) *DescribePriceResponseBody {
	s.TotalTradePrice = &v
	return s
}

type DescribePriceResponseBodyPrices struct {
	PriceInfo []*DescribePriceResponseBodyPricesPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyPrices) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyPrices) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPrices) SetPriceInfo(v []*DescribePriceResponseBodyPricesPriceInfo) *DescribePriceResponseBodyPrices {
	s.PriceInfo = v
	return s
}

type DescribePriceResponseBodyPricesPriceInfo struct {
	// The currency that is used to measure the price. Valid values:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The type of the node. Valid values:
	//
	// *   Manager: management node
	// *   Login: logon node
	// *   Compute: compute node
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The original price of the image.
	//
	// Unit: USD
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The final price.
	//
	// Unit: USD
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribePriceResponseBodyPricesPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyPricesPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPricesPriceInfo) SetCurrency(v string) *DescribePriceResponseBodyPricesPriceInfo {
	s.Currency = &v
	return s
}

func (s *DescribePriceResponseBodyPricesPriceInfo) SetNodeType(v string) *DescribePriceResponseBodyPricesPriceInfo {
	s.NodeType = &v
	return s
}

func (s *DescribePriceResponseBodyPricesPriceInfo) SetOriginalPrice(v float32) *DescribePriceResponseBodyPricesPriceInfo {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPricesPriceInfo) SetTradePrice(v float32) *DescribePriceResponseBodyPricesPriceInfo {
	s.TradePrice = &v
	return s
}

type DescribePriceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribePriceResponse) SetHeaders(v map[string]*string) *DescribePriceResponse {
	s.Headers = v
	return s
}

func (s *DescribePriceResponse) SetStatusCode(v int32) *DescribePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePriceResponse) SetBody(v *DescribePriceResponseBody) *DescribePriceResponse {
	s.Body = v
	return s
}

type EditJobTemplateRequest struct {
	// The job array.
	//
	// Format: X-Y:Z. X is the minimum index value. Y is the maximum index value. Z is the step size. For example, 2-7:2 indicates that three jobs need to be run and their index values are 2, 4, and 6.
	ArrayRequest *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	// The maximum running time of the job. Valid formats:
	//
	// *   hh:mm:ss
	// *   mm:ss
	// *   ss
	//
	// We recommend that you use the hh:mm:ss format. If the maximum running time is 12 hours, set the value to 12:00:00.
	ClockTime *string `json:"ClockTime,omitempty" xml:"ClockTime,omitempty"`
	// The command that is used to run the job.
	CommandLine *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	// The maximum GPU usage required by a single compute node. Valid values: 1 to 8.
	//
	// The parameter takes effect only when the cluster uses PBS and a compute node is a GPU-accelerated instance.
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The URL of the job files that are uploaded to an Object Storage Service (OSS) bucket.
	InputFileUrl *string `json:"InputFileUrl,omitempty" xml:"InputFileUrl,omitempty"`
	// The maximum memory usage required by a single compute node. Unit: GB, MB, or KB. The unit is case-insensitive.
	Mem *string `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The name of the job template.
	//
	// You can call the [ListJobTemplates](~~87248~~) operation to obtain the job template name.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of the compute nodes. Valid values: 1 to 500.
	//
	// >  If the parameter is not specified, the Task, Thread, Mem, and Gpu parameters become invalid.
	Node *int32 `json:"Node,omitempty" xml:"Node,omitempty"`
	// The path that is used to run the job.
	PackagePath *string `json:"PackagePath,omitempty" xml:"PackagePath,omitempty"`
	// The priority of the job. Valid values: 0 to 9. A large value indicates a high priority.
	//
	// Default value: 0
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The name of the queue.
	Queue *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	// Specifies whether the job can be rerun. Valid values:
	//
	// *   true: The job can be rerun.
	// *   false: The job cannot be rerun.
	ReRunable *bool `json:"ReRunable,omitempty" xml:"ReRunable,omitempty"`
	// The name of the user that runs the job.
	//
	// You can call the [ListUsers](~~188572~~) operation to query the users of the cluster.
	RunasUser *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	// The output file path of stderr.
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" xml:"StderrRedirectPath,omitempty"`
	// The output file path of stdout.
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" xml:"StdoutRedirectPath,omitempty"`
	// The number of tasks required by a single compute node. Valid values: 1 to 1000.
	Task *int32 `json:"Task,omitempty" xml:"Task,omitempty"`
	// The ID of the job template.
	//
	// You can call the [ListJobTemplates](~~87248~~) operation to obtain the job template ID.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The number of threads required by a single compute node. Valid values: 1 to 1000.
	Thread *int32 `json:"Thread,omitempty" xml:"Thread,omitempty"`
	// The command that is used to decompress the job files downloaded from an OSS bucket. The parameter takes effect only when WithUnzipCmd is set to true. Valid values:
	//
	// *   tar xzf: decompresses GZIP files.
	// *   tar xf: decompresses TAR files.
	// *   unzip: decompresses ZIP files.
	UnzipCmd *string `json:"UnzipCmd,omitempty" xml:"UnzipCmd,omitempty"`
	// The runtime variables passed to the job. They can be accessed by using environment variables in the executable file.
	Variables *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
	// Specifies whether to decompress the job files downloaded from an OSS bucket. Valid values:
	//
	// *   true: The job files are decompressed.
	// *   false: The job files are not decompressed.
	WithUnzipCmd *bool `json:"WithUnzipCmd,omitempty" xml:"WithUnzipCmd,omitempty"`
}

func (s EditJobTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s EditJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *EditJobTemplateRequest) SetArrayRequest(v string) *EditJobTemplateRequest {
	s.ArrayRequest = &v
	return s
}

func (s *EditJobTemplateRequest) SetClockTime(v string) *EditJobTemplateRequest {
	s.ClockTime = &v
	return s
}

func (s *EditJobTemplateRequest) SetCommandLine(v string) *EditJobTemplateRequest {
	s.CommandLine = &v
	return s
}

func (s *EditJobTemplateRequest) SetGpu(v int32) *EditJobTemplateRequest {
	s.Gpu = &v
	return s
}

func (s *EditJobTemplateRequest) SetInputFileUrl(v string) *EditJobTemplateRequest {
	s.InputFileUrl = &v
	return s
}

func (s *EditJobTemplateRequest) SetMem(v string) *EditJobTemplateRequest {
	s.Mem = &v
	return s
}

func (s *EditJobTemplateRequest) SetName(v string) *EditJobTemplateRequest {
	s.Name = &v
	return s
}

func (s *EditJobTemplateRequest) SetNode(v int32) *EditJobTemplateRequest {
	s.Node = &v
	return s
}

func (s *EditJobTemplateRequest) SetPackagePath(v string) *EditJobTemplateRequest {
	s.PackagePath = &v
	return s
}

func (s *EditJobTemplateRequest) SetPriority(v int32) *EditJobTemplateRequest {
	s.Priority = &v
	return s
}

func (s *EditJobTemplateRequest) SetQueue(v string) *EditJobTemplateRequest {
	s.Queue = &v
	return s
}

func (s *EditJobTemplateRequest) SetReRunable(v bool) *EditJobTemplateRequest {
	s.ReRunable = &v
	return s
}

func (s *EditJobTemplateRequest) SetRunasUser(v string) *EditJobTemplateRequest {
	s.RunasUser = &v
	return s
}

func (s *EditJobTemplateRequest) SetStderrRedirectPath(v string) *EditJobTemplateRequest {
	s.StderrRedirectPath = &v
	return s
}

func (s *EditJobTemplateRequest) SetStdoutRedirectPath(v string) *EditJobTemplateRequest {
	s.StdoutRedirectPath = &v
	return s
}

func (s *EditJobTemplateRequest) SetTask(v int32) *EditJobTemplateRequest {
	s.Task = &v
	return s
}

func (s *EditJobTemplateRequest) SetTemplateId(v string) *EditJobTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *EditJobTemplateRequest) SetThread(v int32) *EditJobTemplateRequest {
	s.Thread = &v
	return s
}

func (s *EditJobTemplateRequest) SetUnzipCmd(v string) *EditJobTemplateRequest {
	s.UnzipCmd = &v
	return s
}

func (s *EditJobTemplateRequest) SetVariables(v string) *EditJobTemplateRequest {
	s.Variables = &v
	return s
}

func (s *EditJobTemplateRequest) SetWithUnzipCmd(v bool) *EditJobTemplateRequest {
	s.WithUnzipCmd = &v
	return s
}

type EditJobTemplateResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the job template.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s EditJobTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *EditJobTemplateResponseBody) SetRequestId(v string) *EditJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *EditJobTemplateResponseBody) SetTemplateId(v string) *EditJobTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type EditJobTemplateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EditJobTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditJobTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s EditJobTemplateResponse) GoString() string {
	return s.String()
}

func (s *EditJobTemplateResponse) SetHeaders(v map[string]*string) *EditJobTemplateResponse {
	s.Headers = v
	return s
}

func (s *EditJobTemplateResponse) SetStatusCode(v int32) *EditJobTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *EditJobTemplateResponse) SetBody(v *EditJobTemplateResponseBody) *EditJobTemplateResponse {
	s.Body = v
	return s
}

type GetAccountingReportRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The layers at which you want to query the bandwidth and traffic data. Valid values:
	//
	// *   user: Query by user.
	// *   queue: Query by queue.
	// *   instance: Query by instance.
	Dim *string `json:"Dim,omitempty" xml:"Dim,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The actual name of the dimension to be queried. Valid values:
	//
	// *   If you set the value of the parameter Dim to user, the value of FilterValue is the name of the specified user.
	// *   If you set the value of the parameter Dim to queue, the value of FilterValue is the name of the specified queue.
	// *   If you set the value of the parameter Dim to instance, the value of FilterValue is the instance name.
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	// The ID of the job.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50.
	//
	// Default value: 10.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number of the returned page.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query type. Valid values:
	//
	// *   total_report: Queries the number of CPU cores in different dimensions.
	// *   job_report: Collects the historical node data of a node.
	// *   number_report: Queries job information in different dimensions.
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAccountingReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountingReportRequest) GoString() string {
	return s.String()
}

func (s *GetAccountingReportRequest) SetClusterId(v string) *GetAccountingReportRequest {
	s.ClusterId = &v
	return s
}

func (s *GetAccountingReportRequest) SetDim(v string) *GetAccountingReportRequest {
	s.Dim = &v
	return s
}

func (s *GetAccountingReportRequest) SetEndTime(v int32) *GetAccountingReportRequest {
	s.EndTime = &v
	return s
}

func (s *GetAccountingReportRequest) SetFilterValue(v string) *GetAccountingReportRequest {
	s.FilterValue = &v
	return s
}

func (s *GetAccountingReportRequest) SetJobId(v string) *GetAccountingReportRequest {
	s.JobId = &v
	return s
}

func (s *GetAccountingReportRequest) SetPageNumber(v int32) *GetAccountingReportRequest {
	s.PageNumber = &v
	return s
}

func (s *GetAccountingReportRequest) SetPageSize(v int32) *GetAccountingReportRequest {
	s.PageSize = &v
	return s
}

func (s *GetAccountingReportRequest) SetReportType(v string) *GetAccountingReportRequest {
	s.ReportType = &v
	return s
}

func (s *GetAccountingReportRequest) SetStartTime(v int32) *GetAccountingReportRequest {
	s.StartTime = &v
	return s
}

type GetAccountingReportResponseBody struct {
	// The list serialized in the JSON format. The list contains multiple records.
	Data *GetAccountingReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The list serialized in the JSON format. The list contains the column names of each record in the Data.
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of CPU cores in the queried cluster.
	TotalCoreTime *int32 `json:"TotalCoreTime,omitempty" xml:"TotalCoreTime,omitempty"`
	// The total number of entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAccountingReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountingReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountingReportResponseBody) SetData(v *GetAccountingReportResponseBodyData) *GetAccountingReportResponseBody {
	s.Data = v
	return s
}

func (s *GetAccountingReportResponseBody) SetMetrics(v string) *GetAccountingReportResponseBody {
	s.Metrics = &v
	return s
}

func (s *GetAccountingReportResponseBody) SetPageNumber(v int32) *GetAccountingReportResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetAccountingReportResponseBody) SetPageSize(v int32) *GetAccountingReportResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetAccountingReportResponseBody) SetRequestId(v string) *GetAccountingReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountingReportResponseBody) SetTotalCoreTime(v int32) *GetAccountingReportResponseBody {
	s.TotalCoreTime = &v
	return s
}

func (s *GetAccountingReportResponseBody) SetTotalCount(v int32) *GetAccountingReportResponseBody {
	s.TotalCount = &v
	return s
}

type GetAccountingReportResponseBodyData struct {
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetAccountingReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAccountingReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAccountingReportResponseBodyData) SetData(v []*string) *GetAccountingReportResponseBodyData {
	s.Data = v
	return s
}

type GetAccountingReportResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAccountingReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccountingReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountingReportResponse) GoString() string {
	return s.String()
}

func (s *GetAccountingReportResponse) SetHeaders(v map[string]*string) *GetAccountingReportResponse {
	s.Headers = v
	return s
}

func (s *GetAccountingReportResponse) SetStatusCode(v int32) *GetAccountingReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountingReportResponse) SetBody(v *GetAccountingReportResponseBody) *GetAccountingReportResponse {
	s.Body = v
	return s
}

type GetAutoScaleConfigRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetAutoScaleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigRequest) SetClusterId(v string) *GetAutoScaleConfigRequest {
	s.ClusterId = &v
	return s
}

type GetAutoScaleConfigResponseBody struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the scheduler. Valid values:
	//
	// *   slurm
	// *   pbs
	// *   opengridscheduler
	// *   deadline
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// Indicates whether the cluster enabled auto scale-out. Valid values:
	//
	// *   true
	// *   false
	EnableAutoGrow *bool `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	// Indicates whether the cluster enabled auto scale-in. Valid values:
	//
	// *   true
	// *   false
	EnableAutoShrink *bool `json:"EnableAutoShrink,omitempty" xml:"EnableAutoShrink,omitempty"`
	// The compute nodes that were excluded from the list of auto scaling nodes. Multiple compute nodes were separated with commas (,).
	ExcludeNodes *string `json:"ExcludeNodes,omitempty" xml:"ExcludeNodes,omitempty"`
	// The percentage of extra compute nodes. Valid values: 0 to 100.
	//
	// If you need to add 100 compute nodes and the value of the ExtraNodesGrowRatio parameter is 2, 102 compute nodes are added.
	ExtraNodesGrowRatio *int32 `json:"ExtraNodesGrowRatio,omitempty" xml:"ExtraNodesGrowRatio,omitempty"`
	// The interval between two consecutive rounds of scale-in. Unit: minutes. Valid values: 2 to 10.
	//
	// >  An interval may exist during multiple rounds of a scale-out task or between two consecutive scale-out tasks.
	GrowIntervalInMinutes *int32 `json:"GrowIntervalInMinutes,omitempty" xml:"GrowIntervalInMinutes,omitempty"`
	// The percentage of each round of scale-out. Valid values: 1 to 100.
	//
	// If you set GrowRatio to 50, the scale-out has two rounds. Each round completes half of the scale-out.
	GrowRatio *int32 `json:"GrowRatio,omitempty" xml:"GrowRatio,omitempty"`
	// The timeout period before the scale-out nodes were started. Unit: minutes. Valid values: 10 to 60.
	//
	// If the scale-out timeout period has been reached but the scale-out nodes still do not reach the Running state, the system resets them.
	GrowTimeoutInMinutes *int32 `json:"GrowTimeoutInMinutes,omitempty" xml:"GrowTimeoutInMinutes,omitempty"`
	// The image ID of the compute nodes in the queue.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The maximum number of compute nodes that can be added in the cluster. Valid values: 0 to 500.
	MaxNodesInCluster *int32 `json:"MaxNodesInCluster,omitempty" xml:"MaxNodesInCluster,omitempty"`
	// The auto scaling configuration of the queue.
	//
	// >  If auto scaling is enabled for the cluster and queue at the same time, the queue settings prevail.
	Queues *GetAutoScaleConfigResponseBodyQueues `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of consecutive times that a compute node is idle during the resource scale-in check. Valid values: 2 to 5.
	//
	// If the parameter is set to 3, a compute node is idle for more than three consecutive times. In this case, the node is released.
	ShrinkIdleTimes *int32 `json:"ShrinkIdleTimes,omitempty" xml:"ShrinkIdleTimes,omitempty"`
	// The interval between two consecutive rounds of scale-out. Unit: minutes. Valid values: 2 to 10.
	ShrinkIntervalInMinutes *int32 `json:"ShrinkIntervalInMinutes,omitempty" xml:"ShrinkIntervalInMinutes,omitempty"`
	// The maximum hourly price of the compute nodes. The value can be accurate to three decimal places. The parameter takes effect only when SpotStrategy is set to SpotWithPriceLimit.
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The preemption policy of the compute nodes. Valid values:
	//
	// *   NoSpot: The compute nodes are pay-as-you-go instances.
	// *   SpotWithPriceLimit: The compute nodes are preemptible instances that have a user-defined maximum hourly price.
	// *   SpotAsPriceGo: The compute nodes are preemptible instances for which the market price at the time of purchase is used as the bid price.
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The ID of the Alibaba Cloud account.
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetAutoScaleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBody) SetClusterId(v string) *GetAutoScaleConfigResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetClusterType(v string) *GetAutoScaleConfigResponseBody {
	s.ClusterType = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetEnableAutoGrow(v bool) *GetAutoScaleConfigResponseBody {
	s.EnableAutoGrow = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetEnableAutoShrink(v bool) *GetAutoScaleConfigResponseBody {
	s.EnableAutoShrink = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetExcludeNodes(v string) *GetAutoScaleConfigResponseBody {
	s.ExcludeNodes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetExtraNodesGrowRatio(v int32) *GetAutoScaleConfigResponseBody {
	s.ExtraNodesGrowRatio = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetGrowIntervalInMinutes(v int32) *GetAutoScaleConfigResponseBody {
	s.GrowIntervalInMinutes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetGrowRatio(v int32) *GetAutoScaleConfigResponseBody {
	s.GrowRatio = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetGrowTimeoutInMinutes(v int32) *GetAutoScaleConfigResponseBody {
	s.GrowTimeoutInMinutes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetImageId(v string) *GetAutoScaleConfigResponseBody {
	s.ImageId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetMaxNodesInCluster(v int32) *GetAutoScaleConfigResponseBody {
	s.MaxNodesInCluster = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetQueues(v *GetAutoScaleConfigResponseBodyQueues) *GetAutoScaleConfigResponseBody {
	s.Queues = v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetRequestId(v string) *GetAutoScaleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetShrinkIdleTimes(v int32) *GetAutoScaleConfigResponseBody {
	s.ShrinkIdleTimes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetShrinkIntervalInMinutes(v int32) *GetAutoScaleConfigResponseBody {
	s.ShrinkIntervalInMinutes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetSpotPriceLimit(v float32) *GetAutoScaleConfigResponseBody {
	s.SpotPriceLimit = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetSpotStrategy(v string) *GetAutoScaleConfigResponseBody {
	s.SpotStrategy = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetUid(v string) *GetAutoScaleConfigResponseBody {
	s.Uid = &v
	return s
}

type GetAutoScaleConfigResponseBodyQueues struct {
	QueueInfo []*GetAutoScaleConfigResponseBodyQueuesQueueInfo `json:"QueueInfo,omitempty" xml:"QueueInfo,omitempty" type:"Repeated"`
}

func (s GetAutoScaleConfigResponseBodyQueues) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBodyQueues) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBodyQueues) SetQueueInfo(v []*GetAutoScaleConfigResponseBodyQueuesQueueInfo) *GetAutoScaleConfigResponseBodyQueues {
	s.QueueInfo = v
	return s
}

type GetAutoScaleConfigResponseBodyQueuesQueueInfo struct {
	// The list of data disks.
	DataDisks *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Struct"`
	// Indicates whether the queue enabled auto scale-out. Valid values:
	//
	// *   true
	// *   false
	EnableAutoGrow *bool `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	// Indicates whether the queue enabled auto scale-in. Valid values:
	//
	// *   true
	// *   false
	EnableAutoShrink *bool `json:"EnableAutoShrink,omitempty" xml:"EnableAutoShrink,omitempty"`
	// The prefix of the queue name. You can query queues that have a specified prefix.
	HostNamePrefix *string `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	// The suffix of the queue name. You can query queues that have a specified suffix.
	HostNameSuffix *string `json:"HostNameSuffix,omitempty" xml:"HostNameSuffix,omitempty"`
	// The instance type of the compute nodes that were automatically added in the queue.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The specification information of the compute nodes.
	InstanceTypes *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Struct"`
	// The maximum number of compute nodes that can be added in a queue. Valid values: 0 to 500.
	MaxNodesInQueue *int32 `json:"MaxNodesInQueue,omitempty" xml:"MaxNodesInQueue,omitempty"`
	// The maximum number of compute nodes that can be added in each round of scale-out. Valid values: 0 to 99.
	//
	// Default value: 0.
	MaxNodesPerCycle *int64 `json:"MaxNodesPerCycle,omitempty" xml:"MaxNodesPerCycle,omitempty"`
	// The minimum number of compute nodes that can be retained in a queue. Valid values: 0 to 50.
	MinNodesInQueue *int32 `json:"MinNodesInQueue,omitempty" xml:"MinNodesInQueue,omitempty"`
	// The minimum number of compute nodes that can be added in each round of scale-out. Valid values: 1 to 99.
	//
	// Default value: 1.
	//
	// If the compute nodes that you want to add in a round is less than the minimum compute nodes that can be added, the value of this parameter is automatically changed to the number of compute nodes that you want to add. This ensures that compute nodes can be added as expected.
	//
	// >  The configuration takes effect only for the minimum compute nodes that can be added in the current round.
	MinNodesPerCycle *int64 `json:"MinNodesPerCycle,omitempty" xml:"MinNodesPerCycle,omitempty"`
	// The image ID of the compute nodes in the queue.
	QueueImageId *string `json:"QueueImageId,omitempty" xml:"QueueImageId,omitempty"`
	// The name of the queue.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the resource group to which the compute nodes belong.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The maximum hourly price of the compute nodes. The value can be accurate to three decimal places. The parameter takes effect only when SpotStrategy is set to SpotWithPriceLimit.
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The preemption policy of the compute nodes. Valid values:
	//
	// *   NoSpot: The compute nodes are pay-as-you-go instances.
	// *   SpotWithPriceLimit: The compute nodes are preemptible instances that have a user-defined maximum hourly price.
	// *   SpotAsPriceGo: The compute nodes are preemptible instances for which the market price at the time of purchase is used as the bid price.
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The type of the system disk. Valid values:
	//
	// *   cloud_efficiency: ultra disk
	// *   cloud_ssd: SSD
	// *   cloud_essd: ESSD
	// *   cloud: basic disk
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The performance level of the system disk. Valid values:
	//
	// *   PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	// *   PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	// *   PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	// *   PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	SystemDiskLevel *string `json:"SystemDiskLevel,omitempty" xml:"SystemDiskLevel,omitempty"`
	// The size of the system disk. Unit: GB. Valid values: 40 to 500.
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfo) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetDataDisks(v *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisks) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.DataDisks = v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetEnableAutoGrow(v bool) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.EnableAutoGrow = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetEnableAutoShrink(v bool) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.EnableAutoShrink = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetHostNamePrefix(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.HostNamePrefix = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetHostNameSuffix(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.HostNameSuffix = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetInstanceType(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.InstanceType = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetInstanceTypes(v *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypes) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.InstanceTypes = v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetMaxNodesInQueue(v int32) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.MaxNodesInQueue = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetMaxNodesPerCycle(v int64) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.MaxNodesPerCycle = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetMinNodesInQueue(v int32) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.MinNodesInQueue = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetMinNodesPerCycle(v int64) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.MinNodesPerCycle = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetQueueImageId(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.QueueImageId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetQueueName(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.QueueName = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetResourceGroupId(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetSpotPriceLimit(v float32) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.SpotPriceLimit = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetSpotStrategy(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.SpotStrategy = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetSystemDiskCategory(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.SystemDiskCategory = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetSystemDiskLevel(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.SystemDiskLevel = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetSystemDiskSize(v int32) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.SystemDiskSize = &v
	return s
}

type GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisks struct {
	DataDisksInfo []*GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo `json:"DataDisksInfo,omitempty" xml:"DataDisksInfo,omitempty" type:"Repeated"`
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisks) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisks) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisks) SetDataDisksInfo(v []*GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisks {
	s.DataDisksInfo = v
	return s
}

type GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo struct {
	// The type of the data disk. Valid values:
	//
	// - cloud_efficiency: ultra disk
	// - cloud_ssd: SSD
	// - cloud_essd: ESSD
	// - cloud: basic disk
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// Indicates whether the data disk is released when the node is released. Valid values:
	//
	// - true
	// - false
	DataDiskDeleteWithInstance *bool `json:"DataDiskDeleteWithInstance,omitempty" xml:"DataDiskDeleteWithInstance,omitempty"`
	// Indicates whether the data disk is encrypted. Valid values:
	//
	// - true
	// - false
	DataDiskEncrypted *bool `json:"DataDiskEncrypted,omitempty" xml:"DataDiskEncrypted,omitempty"`
	// The KMS key ID of the data disk.
	DataDiskKMSKeyId *string `json:"DataDiskKMSKeyId,omitempty" xml:"DataDiskKMSKeyId,omitempty"`
	// The performance level of the ESSD used as the data disk. The parameter takes effect only when the DataDisks.N.DataDiskCategory parameter is set to cloud_essd. Valid values:
	//
	// - PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	// - PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	// - PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	// - PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	DataDiskPerformanceLevel *string `json:"DataDiskPerformanceLevel,omitempty" xml:"DataDiskPerformanceLevel,omitempty"`
	// The capacity of the data disk. Unit: GB.
	//
	// Valid values: 40 to 500
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) SetDataDiskCategory(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo {
	s.DataDiskCategory = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) SetDataDiskDeleteWithInstance(v bool) *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo {
	s.DataDiskDeleteWithInstance = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) SetDataDiskEncrypted(v bool) *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo {
	s.DataDiskEncrypted = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) SetDataDiskKMSKeyId(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo {
	s.DataDiskKMSKeyId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) SetDataDiskPerformanceLevel(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo {
	s.DataDiskPerformanceLevel = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) SetDataDiskSize(v int32) *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo {
	s.DataDiskSize = &v
	return s
}

type GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypes struct {
	InstanceTypeInfo []*GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo `json:"InstanceTypeInfo,omitempty" xml:"InstanceTypeInfo,omitempty" type:"Repeated"`
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypes) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypes) SetInstanceTypeInfo(v []*GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypes {
	s.InstanceTypeInfo = v
	return s
}

type GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo struct {
	// The prefix of the hostname. You can query compute nodes that have a specified prefix.
	HostNamePrefix *string `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	// The instance type of the node.
	InstanceType             *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	SpotDuration             *int32  `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	// The maximum hourly price of the compute nodes. The value can be accurate to three decimal places. The parameter takes effect only when SpotStrategy is set to SpotWithPriceLimit.
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding method of the compute nodes. Valid values:
	//
	// *   NoSpot: The compute nodes are pay-as-you-go instances.
	// *   SpotWithPriceLimit: The compute nodes are preemptible instances that have a user-defined maximum hourly price.
	// *   SpotAsPriceGo: The compute nodes are preemptible instances for which the market price at the time of purchase is used as the bid price.
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The ID of the vSwitch.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) SetHostNamePrefix(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo {
	s.HostNamePrefix = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) SetInstanceType(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo {
	s.InstanceType = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) SetSpotDuration(v int32) *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo {
	s.SpotDuration = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) SetSpotInterruptionBehavior(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) SetSpotPriceLimit(v float32) *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo {
	s.SpotPriceLimit = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) SetSpotStrategy(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo {
	s.SpotStrategy = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) SetVSwitchId(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo {
	s.VSwitchId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) SetZoneId(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo {
	s.ZoneId = &v
	return s
}

type GetAutoScaleConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAutoScaleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAutoScaleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponse) SetHeaders(v map[string]*string) *GetAutoScaleConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAutoScaleConfigResponse) SetStatusCode(v int32) *GetAutoScaleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoScaleConfigResponse) SetBody(v *GetAutoScaleConfigResponseBody) *GetAutoScaleConfigResponse {
	s.Body = v
	return s
}

type GetCloudMetricLogsRequest struct {
	// The data aggregation interval. Unit: seconds.
	//
	// Valid values: 1, 10, 60, 600, and 3600.
	//
	// Default value: 1
	AggregationInterval *int32 `json:"AggregationInterval,omitempty" xml:"AggregationInterval,omitempty"`
	// The data aggregation type. Valid values:
	//
	// *   sum: the sum of the data
	// *   avg: the average value
	// *   max: the maximum value
	// *   min: the minimum value
	//
	// Aggregation is disabled by default.
	AggregationType *string `json:"AggregationType,omitempty" xml:"AggregationType,omitempty"`
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The filter conditions. A JSON string consisting of one or more key:value pairs. Value range of key:
	//
	// *   InstanceId: the ID of the node
	// *   Hostname: the hostname of the node
	// *   NetworkInterface: the name of the network interface
	// *   DiskDevice: the name of the disk
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// The category of the output performance metrics. Separate multiple metrics with commas (,). Valid values:
	//
	// *   cpu
	// *   memory
	MetricCategories *string `json:"MetricCategories,omitempty" xml:"MetricCategories,omitempty"`
	// The dimensions of the performance metric. Valid values:
	//
	// *   machine
	// *   process
	// *   network
	// *   disk
	MetricScope *string `json:"MetricScope,omitempty" xml:"MetricScope,omitempty"`
	// Logs are returned in reverse order of timestamps.
	//
	// Default value: false
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	To *int32 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetCloudMetricLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricLogsRequest) GoString() string {
	return s.String()
}

func (s *GetCloudMetricLogsRequest) SetAggregationInterval(v int32) *GetCloudMetricLogsRequest {
	s.AggregationInterval = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetAggregationType(v string) *GetCloudMetricLogsRequest {
	s.AggregationType = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetClusterId(v string) *GetCloudMetricLogsRequest {
	s.ClusterId = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetFilter(v string) *GetCloudMetricLogsRequest {
	s.Filter = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetFrom(v int32) *GetCloudMetricLogsRequest {
	s.From = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetMetricCategories(v string) *GetCloudMetricLogsRequest {
	s.MetricCategories = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetMetricScope(v string) *GetCloudMetricLogsRequest {
	s.MetricScope = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetReverse(v bool) *GetCloudMetricLogsRequest {
	s.Reverse = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetTo(v int32) *GetCloudMetricLogsRequest {
	s.To = &v
	return s
}

type GetCloudMetricLogsResponseBody struct {
	// The list of the performance data.
	MetricLogs *GetCloudMetricLogsResponseBodyMetricLogs `json:"MetricLogs,omitempty" xml:"MetricLogs,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCloudMetricLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricLogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudMetricLogsResponseBody) SetMetricLogs(v *GetCloudMetricLogsResponseBodyMetricLogs) *GetCloudMetricLogsResponseBody {
	s.MetricLogs = v
	return s
}

func (s *GetCloudMetricLogsResponseBody) SetRequestId(v string) *GetCloudMetricLogsResponseBody {
	s.RequestId = &v
	return s
}

type GetCloudMetricLogsResponseBodyMetricLogs struct {
	MetricLog []*GetCloudMetricLogsResponseBodyMetricLogsMetricLog `json:"MetricLog,omitempty" xml:"MetricLog,omitempty" type:"Repeated"`
}

func (s GetCloudMetricLogsResponseBodyMetricLogs) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricLogsResponseBodyMetricLogs) GoString() string {
	return s.String()
}

func (s *GetCloudMetricLogsResponseBodyMetricLogs) SetMetricLog(v []*GetCloudMetricLogsResponseBodyMetricLogsMetricLog) *GetCloudMetricLogsResponseBodyMetricLogs {
	s.MetricLog = v
	return s
}

type GetCloudMetricLogsResponseBodyMetricLogsMetricLog struct {
	// The name of the disk.
	DiskDevice *string `json:"DiskDevice,omitempty" xml:"DiskDevice,omitempty"`
	// The hostname of the node.
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The ID of the node.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// A JSON-serialized string that contains values for multiple performance metrics.
	MetricData *string `json:"MetricData,omitempty" xml:"MetricData,omitempty"`
	// The name of the network interface.
	NetworkInterface *string `json:"NetworkInterface,omitempty" xml:"NetworkInterface,omitempty"`
	// The timestamp of the log. This value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	Time *int32 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetCloudMetricLogsResponseBodyMetricLogsMetricLog) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricLogsResponseBodyMetricLogsMetricLog) GoString() string {
	return s.String()
}

func (s *GetCloudMetricLogsResponseBodyMetricLogsMetricLog) SetDiskDevice(v string) *GetCloudMetricLogsResponseBodyMetricLogsMetricLog {
	s.DiskDevice = &v
	return s
}

func (s *GetCloudMetricLogsResponseBodyMetricLogsMetricLog) SetHostname(v string) *GetCloudMetricLogsResponseBodyMetricLogsMetricLog {
	s.Hostname = &v
	return s
}

func (s *GetCloudMetricLogsResponseBodyMetricLogsMetricLog) SetInstanceId(v string) *GetCloudMetricLogsResponseBodyMetricLogsMetricLog {
	s.InstanceId = &v
	return s
}

func (s *GetCloudMetricLogsResponseBodyMetricLogsMetricLog) SetMetricData(v string) *GetCloudMetricLogsResponseBodyMetricLogsMetricLog {
	s.MetricData = &v
	return s
}

func (s *GetCloudMetricLogsResponseBodyMetricLogsMetricLog) SetNetworkInterface(v string) *GetCloudMetricLogsResponseBodyMetricLogsMetricLog {
	s.NetworkInterface = &v
	return s
}

func (s *GetCloudMetricLogsResponseBodyMetricLogsMetricLog) SetTime(v int32) *GetCloudMetricLogsResponseBodyMetricLogsMetricLog {
	s.Time = &v
	return s
}

type GetCloudMetricLogsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCloudMetricLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCloudMetricLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricLogsResponse) GoString() string {
	return s.String()
}

func (s *GetCloudMetricLogsResponse) SetHeaders(v map[string]*string) *GetCloudMetricLogsResponse {
	s.Headers = v
	return s
}

func (s *GetCloudMetricLogsResponse) SetStatusCode(v int32) *GetCloudMetricLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCloudMetricLogsResponse) SetBody(v *GetCloudMetricLogsResponseBody) *GetCloudMetricLogsResponse {
	s.Body = v
	return s
}

type GetCloudMetricProfilingRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The profiling ID. You can call the [ListCloudMetricProfilings](~~188711~~) operation to obtain the profiling ID.
	ProfilingId *string `json:"ProfilingId,omitempty" xml:"ProfilingId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetCloudMetricProfilingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricProfilingRequest) GoString() string {
	return s.String()
}

func (s *GetCloudMetricProfilingRequest) SetClusterId(v string) *GetCloudMetricProfilingRequest {
	s.ClusterId = &v
	return s
}

func (s *GetCloudMetricProfilingRequest) SetProfilingId(v string) *GetCloudMetricProfilingRequest {
	s.ProfilingId = &v
	return s
}

func (s *GetCloudMetricProfilingRequest) SetRegionId(v string) *GetCloudMetricProfilingRequest {
	s.RegionId = &v
	return s
}

type GetCloudMetricProfilingResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of profiling results.
	SvgUrls *GetCloudMetricProfilingResponseBodySvgUrls `json:"SvgUrls,omitempty" xml:"SvgUrls,omitempty" type:"Struct"`
}

func (s GetCloudMetricProfilingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricProfilingResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudMetricProfilingResponseBody) SetRequestId(v string) *GetCloudMetricProfilingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudMetricProfilingResponseBody) SetSvgUrls(v *GetCloudMetricProfilingResponseBodySvgUrls) *GetCloudMetricProfilingResponseBody {
	s.SvgUrls = v
	return s
}

type GetCloudMetricProfilingResponseBodySvgUrls struct {
	SvgInfo []*GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo `json:"SvgInfo,omitempty" xml:"SvgInfo,omitempty" type:"Repeated"`
}

func (s GetCloudMetricProfilingResponseBodySvgUrls) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricProfilingResponseBodySvgUrls) GoString() string {
	return s.String()
}

func (s *GetCloudMetricProfilingResponseBodySvgUrls) SetSvgInfo(v []*GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo) *GetCloudMetricProfilingResponseBodySvgUrls {
	s.SvgInfo = v
	return s
}

type GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo struct {
	// The name of the SVG file that contains the profiling results.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The size of the SVG file. Unit: bytes.
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The type of the SVG file.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the Object Storage Service (OSS) bucket where the scalable vector graphics (SVG) file is stored.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo) GoString() string {
	return s.String()
}

func (s *GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo) SetName(v string) *GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo {
	s.Name = &v
	return s
}

func (s *GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo) SetSize(v int32) *GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo {
	s.Size = &v
	return s
}

func (s *GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo) SetType(v string) *GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo {
	s.Type = &v
	return s
}

func (s *GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo) SetUrl(v string) *GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo {
	s.Url = &v
	return s
}

type GetCloudMetricProfilingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCloudMetricProfilingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCloudMetricProfilingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricProfilingResponse) GoString() string {
	return s.String()
}

func (s *GetCloudMetricProfilingResponse) SetHeaders(v map[string]*string) *GetCloudMetricProfilingResponse {
	s.Headers = v
	return s
}

func (s *GetCloudMetricProfilingResponse) SetStatusCode(v int32) *GetCloudMetricProfilingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCloudMetricProfilingResponse) SetBody(v *GetCloudMetricProfilingResponseBody) *GetCloudMetricProfilingResponse {
	s.Body = v
	return s
}

type GetClusterVolumesRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetClusterVolumesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesRequest) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesRequest) SetClusterId(v string) *GetClusterVolumesRequest {
	s.ClusterId = &v
	return s
}

type GetClusterVolumesResponseBody struct {
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the file system mounted to the E-HPC cluster.
	Volumes *GetClusterVolumesResponseBodyVolumes `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Struct"`
}

func (s GetClusterVolumesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesResponseBody) SetRegionId(v string) *GetClusterVolumesResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetClusterVolumesResponseBody) SetRequestId(v string) *GetClusterVolumesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterVolumesResponseBody) SetVolumes(v *GetClusterVolumesResponseBodyVolumes) *GetClusterVolumesResponseBody {
	s.Volumes = v
	return s
}

type GetClusterVolumesResponseBodyVolumes struct {
	VolumeInfo []*GetClusterVolumesResponseBodyVolumesVolumeInfo `json:"VolumeInfo,omitempty" xml:"VolumeInfo,omitempty" type:"Repeated"`
}

func (s GetClusterVolumesResponseBodyVolumes) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesResponseBodyVolumes) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesResponseBodyVolumes) SetVolumeInfo(v []*GetClusterVolumesResponseBodyVolumesVolumeInfo) *GetClusterVolumesResponseBodyVolumes {
	s.VolumeInfo = v
	return s
}

type GetClusterVolumesResponseBodyVolumesVolumeInfo struct {
	// The queue of the job.
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	// The local mount directory.
	LocalDirectory *string `json:"LocalDirectory,omitempty" xml:"LocalDirectory,omitempty"`
	// The type of cluster. Valid values:
	//
	// *   OnPremise: The cluster is deployed on a hybrid cloud.
	// *   PublicCloud: The cluster is deployed on a public cloud.
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// Indicates whether the resource can be unmounted.
	MustKeep *bool `json:"MustKeep,omitempty" xml:"MustKeep,omitempty"`
	// The remote mount directory.
	RemoteDirectory *string `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	// The array of the node on which the file system is mounted.
	Roles *GetClusterVolumesResponseBodyVolumesVolumeInfoRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
	// The ID of the file system.
	VolumeId *string `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	// The address of the mount target.
	VolumeMountpoint *string `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	// The storage protocol type of the file system. Valid values:
	//
	// *   NFS
	// *   SMB
	VolumeProtocol *string `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	// The type of the file system. Valid values:
	//
	// *   NAS
	VolumeType *string `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
}

func (s GetClusterVolumesResponseBodyVolumesVolumeInfo) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesResponseBodyVolumesVolumeInfo) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetJobQueue(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.JobQueue = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetLocalDirectory(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.LocalDirectory = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetLocation(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.Location = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetMustKeep(v bool) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.MustKeep = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetRemoteDirectory(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.RemoteDirectory = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetRoles(v *GetClusterVolumesResponseBodyVolumesVolumeInfoRoles) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.Roles = v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetVolumeId(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeId = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetVolumeMountpoint(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeMountpoint = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetVolumeProtocol(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeProtocol = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetVolumeType(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeType = &v
	return s
}

type GetClusterVolumesResponseBodyVolumesVolumeInfoRoles struct {
	RoleInfo []*GetClusterVolumesResponseBodyVolumesVolumeInfoRolesRoleInfo `json:"RoleInfo,omitempty" xml:"RoleInfo,omitempty" type:"Repeated"`
}

func (s GetClusterVolumesResponseBodyVolumesVolumeInfoRoles) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesResponseBodyVolumesVolumeInfoRoles) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfoRoles) SetRoleInfo(v []*GetClusterVolumesResponseBodyVolumesVolumeInfoRolesRoleInfo) *GetClusterVolumesResponseBodyVolumesVolumeInfoRoles {
	s.RoleInfo = v
	return s
}

type GetClusterVolumesResponseBodyVolumesVolumeInfoRolesRoleInfo struct {
	// The type of the node on which the file system is mounted. Valid values:
	//
	// *   Compute: compute node
	// *   Manager: management node
	// *   Login: logon node
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetClusterVolumesResponseBodyVolumesVolumeInfoRolesRoleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesResponseBodyVolumesVolumeInfoRolesRoleInfo) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfoRolesRoleInfo) SetName(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfoRolesRoleInfo {
	s.Name = &v
	return s
}

type GetClusterVolumesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetClusterVolumesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClusterVolumesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesResponse) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesResponse) SetHeaders(v map[string]*string) *GetClusterVolumesResponse {
	s.Headers = v
	return s
}

func (s *GetClusterVolumesResponse) SetStatusCode(v int32) *GetClusterVolumesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterVolumesResponse) SetBody(v *GetClusterVolumesResponseBody) *GetClusterVolumesResponse {
	s.Body = v
	return s
}

type GetCommonImageRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ContainType *string `json:"ContainType,omitempty" xml:"ContainType,omitempty"`
	ImageName   *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetCommonImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCommonImageRequest) GoString() string {
	return s.String()
}

func (s *GetCommonImageRequest) SetClusterId(v string) *GetCommonImageRequest {
	s.ClusterId = &v
	return s
}

func (s *GetCommonImageRequest) SetContainType(v string) *GetCommonImageRequest {
	s.ContainType = &v
	return s
}

func (s *GetCommonImageRequest) SetImageName(v string) *GetCommonImageRequest {
	s.ImageName = &v
	return s
}

func (s *GetCommonImageRequest) SetRegionId(v string) *GetCommonImageRequest {
	s.RegionId = &v
	return s
}

type GetCommonImageResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCommonImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCommonImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetCommonImageResponseBody) SetRequestId(v string) *GetCommonImageResponseBody {
	s.RequestId = &v
	return s
}

type GetCommonImageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCommonImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCommonImageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCommonImageResponse) GoString() string {
	return s.String()
}

func (s *GetCommonImageResponse) SetHeaders(v map[string]*string) *GetCommonImageResponse {
	s.Headers = v
	return s
}

func (s *GetCommonImageResponse) SetStatusCode(v int32) *GetCommonImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCommonImageResponse) SetBody(v *GetCommonImageResponseBody) *GetCommonImageResponse {
	s.Body = v
	return s
}

type GetGWSConnectTicketRequest struct {
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetGWSConnectTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGWSConnectTicketRequest) GoString() string {
	return s.String()
}

func (s *GetGWSConnectTicketRequest) SetAppName(v string) *GetGWSConnectTicketRequest {
	s.AppName = &v
	return s
}

func (s *GetGWSConnectTicketRequest) SetInstanceId(v string) *GetGWSConnectTicketRequest {
	s.InstanceId = &v
	return s
}

type GetGWSConnectTicketResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Ticket    *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s GetGWSConnectTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGWSConnectTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetGWSConnectTicketResponseBody) SetRequestId(v string) *GetGWSConnectTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGWSConnectTicketResponseBody) SetTicket(v string) *GetGWSConnectTicketResponseBody {
	s.Ticket = &v
	return s
}

type GetGWSConnectTicketResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGWSConnectTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGWSConnectTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGWSConnectTicketResponse) GoString() string {
	return s.String()
}

func (s *GetGWSConnectTicketResponse) SetHeaders(v map[string]*string) *GetGWSConnectTicketResponse {
	s.Headers = v
	return s
}

func (s *GetGWSConnectTicketResponse) SetStatusCode(v int32) *GetGWSConnectTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGWSConnectTicketResponse) SetBody(v *GetGWSConnectTicketResponseBody) *GetGWSConnectTicketResponse {
	s.Body = v
	return s
}

type GetHybridClusterConfigRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the on-premises compute node. You can call this operation to query the configurations of the on-premises compute node.
	//
	// By default, the operation queries the configurations of a cluster.
	Node *string `json:"Node,omitempty" xml:"Node,omitempty"`
}

func (s GetHybridClusterConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHybridClusterConfigRequest) GoString() string {
	return s.String()
}

func (s *GetHybridClusterConfigRequest) SetClusterId(v string) *GetHybridClusterConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *GetHybridClusterConfigRequest) SetNode(v string) *GetHybridClusterConfigRequest {
	s.Node = &v
	return s
}

type GetHybridClusterConfigResponseBody struct {
	// The configurations returned.
	//
	// *   If the parameter Node is null, you can obtain the configurations of the hybrid cloud cluster.
	// *   If the parameter Node is a specified on-premises compute node, you can obtain the configurations of the on-premises compute node.
	//
	// This parameter is returned in the ini format. You can use this parameter to configure on-premises cluster nodes.
	ClusterConfig *string `json:"ClusterConfig,omitempty" xml:"ClusterConfig,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHybridClusterConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHybridClusterConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetHybridClusterConfigResponseBody) SetClusterConfig(v string) *GetHybridClusterConfigResponseBody {
	s.ClusterConfig = &v
	return s
}

func (s *GetHybridClusterConfigResponseBody) SetRequestId(v string) *GetHybridClusterConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetHybridClusterConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHybridClusterConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHybridClusterConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHybridClusterConfigResponse) GoString() string {
	return s.String()
}

func (s *GetHybridClusterConfigResponse) SetHeaders(v map[string]*string) *GetHybridClusterConfigResponse {
	s.Headers = v
	return s
}

func (s *GetHybridClusterConfigResponse) SetStatusCode(v int32) *GetHybridClusterConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHybridClusterConfigResponse) SetBody(v *GetHybridClusterConfigResponseBody) *GetHybridClusterConfigResponse {
	s.Body = v
	return s
}

type GetIfEcsTypeSupportHtConfigRequest struct {
	// The instance type of the ECS instance.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s GetIfEcsTypeSupportHtConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIfEcsTypeSupportHtConfigRequest) GoString() string {
	return s.String()
}

func (s *GetIfEcsTypeSupportHtConfigRequest) SetInstanceType(v string) *GetIfEcsTypeSupportHtConfigRequest {
	s.InstanceType = &v
	return s
}

type GetIfEcsTypeSupportHtConfigResponseBody struct {
	// Indicates whether Hyper-Threading is enabled by default. Valid values:
	//
	// *   true: Hyper-Threading is enabled by default.
	//
	// *   false: Hyper-Threading is disabled by default
	//
	// > By default, Hyper-Threading is not enabled for the SCC specification family, while Hyper-Threading is enabled for other specification families by default.
	DefaultHtEnabled *bool `json:"DefaultHtEnabled,omitempty" xml:"DefaultHtEnabled,omitempty"`
	// The instance type of the ECS instance.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether hyper-threading is supported. Valid values:
	//
	// *   true: Hyper-Threading is supported.
	// *   false: Hyper-Threading is not supported.
	SupportHtConfig *bool `json:"SupportHtConfig,omitempty" xml:"SupportHtConfig,omitempty"`
}

func (s GetIfEcsTypeSupportHtConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIfEcsTypeSupportHtConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetIfEcsTypeSupportHtConfigResponseBody) SetDefaultHtEnabled(v bool) *GetIfEcsTypeSupportHtConfigResponseBody {
	s.DefaultHtEnabled = &v
	return s
}

func (s *GetIfEcsTypeSupportHtConfigResponseBody) SetInstanceType(v string) *GetIfEcsTypeSupportHtConfigResponseBody {
	s.InstanceType = &v
	return s
}

func (s *GetIfEcsTypeSupportHtConfigResponseBody) SetRequestId(v string) *GetIfEcsTypeSupportHtConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIfEcsTypeSupportHtConfigResponseBody) SetSupportHtConfig(v bool) *GetIfEcsTypeSupportHtConfigResponseBody {
	s.SupportHtConfig = &v
	return s
}

type GetIfEcsTypeSupportHtConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIfEcsTypeSupportHtConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIfEcsTypeSupportHtConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIfEcsTypeSupportHtConfigResponse) GoString() string {
	return s.String()
}

func (s *GetIfEcsTypeSupportHtConfigResponse) SetHeaders(v map[string]*string) *GetIfEcsTypeSupportHtConfigResponse {
	s.Headers = v
	return s
}

func (s *GetIfEcsTypeSupportHtConfigResponse) SetStatusCode(v int32) *GetIfEcsTypeSupportHtConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIfEcsTypeSupportHtConfigResponse) SetBody(v *GetIfEcsTypeSupportHtConfigResponseBody) *GetIfEcsTypeSupportHtConfigResponse {
	s.Body = v
	return s
}

type GetJobLogRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The node on which the job runs.
	//
	// *   If the job is completed, you do not need to specify the parameter.
	// *   If the job is running, you must specify the parameter.
	ExecHost *string `json:"ExecHost,omitempty" xml:"ExecHost,omitempty"`
	// The ID of the job.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The position where logs start to be read.
	//
	// Unit: bits
	//
	// Default value: 0
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The maximum size of logs that you can read in a single request.
	//
	// Unit: bits
	//
	// Default value: 1024
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s GetJobLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobLogRequest) GoString() string {
	return s.String()
}

func (s *GetJobLogRequest) SetClusterId(v string) *GetJobLogRequest {
	s.ClusterId = &v
	return s
}

func (s *GetJobLogRequest) SetExecHost(v string) *GetJobLogRequest {
	s.ExecHost = &v
	return s
}

func (s *GetJobLogRequest) SetJobId(v string) *GetJobLogRequest {
	s.JobId = &v
	return s
}

func (s *GetJobLogRequest) SetOffset(v int64) *GetJobLogRequest {
	s.Offset = &v
	return s
}

func (s *GetJobLogRequest) SetSize(v int32) *GetJobLogRequest {
	s.Size = &v
	return s
}

type GetJobLogResponseBody struct {
	// The content of the error logs. The content is encoded in Base64.
	ErrorLog *string `json:"ErrorLog,omitempty" xml:"ErrorLog,omitempty"`
	// The ID of the job.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The content of the output logs. The content is encoded in Base64.
	OutputLog *string `json:"OutputLog,omitempty" xml:"OutputLog,omitempty"`
	// The ID of the task.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobLogResponseBody) SetErrorLog(v string) *GetJobLogResponseBody {
	s.ErrorLog = &v
	return s
}

func (s *GetJobLogResponseBody) SetJobId(v string) *GetJobLogResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobLogResponseBody) SetOutputLog(v string) *GetJobLogResponseBody {
	s.OutputLog = &v
	return s
}

func (s *GetJobLogResponseBody) SetRequestId(v string) *GetJobLogResponseBody {
	s.RequestId = &v
	return s
}

type GetJobLogResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetJobLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobLogResponse) GoString() string {
	return s.String()
}

func (s *GetJobLogResponse) SetHeaders(v map[string]*string) *GetJobLogResponse {
	s.Headers = v
	return s
}

func (s *GetJobLogResponse) SetStatusCode(v int32) *GetJobLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobLogResponse) SetBody(v *GetJobLogResponseBody) *GetJobLogResponse {
	s.Body = v
	return s
}

type GetPostScriptsRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the region.
	//
	// You can call the [ListRegions](~~188593~~) operation to query the latest region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPostScriptsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPostScriptsRequest) GoString() string {
	return s.String()
}

func (s *GetPostScriptsRequest) SetClusterId(v string) *GetPostScriptsRequest {
	s.ClusterId = &v
	return s
}

func (s *GetPostScriptsRequest) SetRegionId(v string) *GetPostScriptsRequest {
	s.RegionId = &v
	return s
}

type GetPostScriptsResponseBody struct {
	// The post-installation scripts.
	PostInstallScripts []*GetPostScriptsResponseBodyPostInstallScripts `json:"PostInstallScripts,omitempty" xml:"PostInstallScripts,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPostScriptsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPostScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPostScriptsResponseBody) SetPostInstallScripts(v []*GetPostScriptsResponseBodyPostInstallScripts) *GetPostScriptsResponseBody {
	s.PostInstallScripts = v
	return s
}

func (s *GetPostScriptsResponseBody) SetRequestId(v string) *GetPostScriptsResponseBody {
	s.RequestId = &v
	return s
}

type GetPostScriptsResponseBodyPostInstallScripts struct {
	// The parameter that is used to run the Nth post-installation script. Valid values of N: 1 to 16.
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The URL that is used to download the Nth post-installation script. Valid values of N: 1 to 16.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetPostScriptsResponseBodyPostInstallScripts) String() string {
	return tea.Prettify(s)
}

func (s GetPostScriptsResponseBodyPostInstallScripts) GoString() string {
	return s.String()
}

func (s *GetPostScriptsResponseBodyPostInstallScripts) SetArgs(v string) *GetPostScriptsResponseBodyPostInstallScripts {
	s.Args = &v
	return s
}

func (s *GetPostScriptsResponseBodyPostInstallScripts) SetUrl(v string) *GetPostScriptsResponseBodyPostInstallScripts {
	s.Url = &v
	return s
}

type GetPostScriptsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPostScriptsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPostScriptsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPostScriptsResponse) GoString() string {
	return s.String()
}

func (s *GetPostScriptsResponse) SetHeaders(v map[string]*string) *GetPostScriptsResponse {
	s.Headers = v
	return s
}

func (s *GetPostScriptsResponse) SetStatusCode(v int32) *GetPostScriptsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPostScriptsResponse) SetBody(v *GetPostScriptsResponseBody) *GetPostScriptsResponse {
	s.Body = v
	return s
}

type GetSchedulerInfoRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the region.
	RegionId  *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Scheduler []*GetSchedulerInfoRequestScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Repeated"`
}

func (s GetSchedulerInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSchedulerInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSchedulerInfoRequest) SetClusterId(v string) *GetSchedulerInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *GetSchedulerInfoRequest) SetRegionId(v string) *GetSchedulerInfoRequest {
	s.RegionId = &v
	return s
}

func (s *GetSchedulerInfoRequest) SetScheduler(v []*GetSchedulerInfoRequestScheduler) *GetSchedulerInfoRequest {
	s.Scheduler = v
	return s
}

type GetSchedulerInfoRequestScheduler struct {
	// The name of the scheduler. Valid values:
	//
	// *   pbs
	// *   pbs19
	// *   slurm
	// *   slurm19
	// *   slurm20
	//
	// Valid values of N: 0 to 100
	SchedName *string `json:"SchedName,omitempty" xml:"SchedName,omitempty"`
}

func (s GetSchedulerInfoRequestScheduler) String() string {
	return tea.Prettify(s)
}

func (s GetSchedulerInfoRequestScheduler) GoString() string {
	return s.String()
}

func (s *GetSchedulerInfoRequestScheduler) SetSchedName(v string) *GetSchedulerInfoRequestScheduler {
	s.SchedName = &v
	return s
}

type GetSchedulerInfoResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The settings of the scheduler.
	SchedInfo []*GetSchedulerInfoResponseBodySchedInfo `json:"SchedInfo,omitempty" xml:"SchedInfo,omitempty" type:"Repeated"`
}

func (s GetSchedulerInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSchedulerInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSchedulerInfoResponseBody) SetRequestId(v string) *GetSchedulerInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSchedulerInfoResponseBody) SetSchedInfo(v []*GetSchedulerInfoResponseBodySchedInfo) *GetSchedulerInfoResponseBody {
	s.SchedInfo = v
	return s
}

type GetSchedulerInfoResponseBodySchedInfo struct {
	// The detailed settings of the scheduler.
	Configuration *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The type of the scheduler.
	SchedName *string `json:"SchedName,omitempty" xml:"SchedName,omitempty"`
}

func (s GetSchedulerInfoResponseBodySchedInfo) String() string {
	return tea.Prettify(s)
}

func (s GetSchedulerInfoResponseBodySchedInfo) GoString() string {
	return s.String()
}

func (s *GetSchedulerInfoResponseBodySchedInfo) SetConfiguration(v string) *GetSchedulerInfoResponseBodySchedInfo {
	s.Configuration = &v
	return s
}

func (s *GetSchedulerInfoResponseBodySchedInfo) SetSchedName(v string) *GetSchedulerInfoResponseBodySchedInfo {
	s.SchedName = &v
	return s
}

type GetSchedulerInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSchedulerInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSchedulerInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSchedulerInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSchedulerInfoResponse) SetHeaders(v map[string]*string) *GetSchedulerInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSchedulerInfoResponse) SetStatusCode(v int32) *GetSchedulerInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSchedulerInfoResponse) SetBody(v *GetSchedulerInfoResponseBody) *GetSchedulerInfoResponse {
	s.Body = v
	return s
}

type GetUserImageRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the image. Set the value to singularity.
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
	// The name of the image.
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The path where the image is stored in the OSS bucket.
	ImagePath *string `json:"ImagePath,omitempty" xml:"ImagePath,omitempty"`
	// The OSS bucket.
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	// The endpoint of OSS.
	OSSEndPoint *string `json:"OSSEndPoint,omitempty" xml:"OSSEndPoint,omitempty"`
}

func (s GetUserImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserImageRequest) GoString() string {
	return s.String()
}

func (s *GetUserImageRequest) SetClusterId(v string) *GetUserImageRequest {
	s.ClusterId = &v
	return s
}

func (s *GetUserImageRequest) SetContainerType(v string) *GetUserImageRequest {
	s.ContainerType = &v
	return s
}

func (s *GetUserImageRequest) SetImageName(v string) *GetUserImageRequest {
	s.ImageName = &v
	return s
}

func (s *GetUserImageRequest) SetImagePath(v string) *GetUserImageRequest {
	s.ImagePath = &v
	return s
}

func (s *GetUserImageRequest) SetOSSBucket(v string) *GetUserImageRequest {
	s.OSSBucket = &v
	return s
}

func (s *GetUserImageRequest) SetOSSEndPoint(v string) *GetUserImageRequest {
	s.OSSEndPoint = &v
	return s
}

type GetUserImageResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserImageResponseBody) SetRequestId(v string) *GetUserImageResponseBody {
	s.RequestId = &v
	return s
}

type GetUserImageResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserImageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserImageResponse) GoString() string {
	return s.String()
}

func (s *GetUserImageResponse) SetHeaders(v map[string]*string) *GetUserImageResponse {
	s.Headers = v
	return s
}

func (s *GetUserImageResponse) SetStatusCode(v int32) *GetUserImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserImageResponse) SetBody(v *GetUserImageResponseBody) *GetUserImageResponse {
	s.Body = v
	return s
}

type GetVisualServiceStatusRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetVisualServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVisualServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetVisualServiceStatusRequest) SetClusterId(v string) *GetVisualServiceStatusRequest {
	s.ClusterId = &v
	return s
}

type GetVisualServiceStatusResponseBody struct {
	// The response message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVisualServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVisualServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetVisualServiceStatusResponseBody) SetMessage(v string) *GetVisualServiceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetVisualServiceStatusResponseBody) SetRequestId(v string) *GetVisualServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetVisualServiceStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVisualServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVisualServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVisualServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetVisualServiceStatusResponse) SetHeaders(v map[string]*string) *GetVisualServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetVisualServiceStatusResponse) SetStatusCode(v int32) *GetVisualServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVisualServiceStatusResponse) SetBody(v *GetVisualServiceStatusResponseBody) *GetVisualServiceStatusResponse {
	s.Body = v
	return s
}

type InitializeEHPCRequest struct {
	// The ID of the region where the service-linked role is created.
	//
	// You can call the [ListRegions](~~188593~~) operation to obtain the IDs of regions supported by E-HPC.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InitializeEHPCRequest) String() string {
	return tea.Prettify(s)
}

func (s InitializeEHPCRequest) GoString() string {
	return s.String()
}

func (s *InitializeEHPCRequest) SetRegionId(v string) *InitializeEHPCRequest {
	s.RegionId = &v
	return s
}

type InitializeEHPCResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitializeEHPCResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitializeEHPCResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeEHPCResponseBody) SetRequestId(v string) *InitializeEHPCResponseBody {
	s.RequestId = &v
	return s
}

type InitializeEHPCResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InitializeEHPCResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitializeEHPCResponse) String() string {
	return tea.Prettify(s)
}

func (s InitializeEHPCResponse) GoString() string {
	return s.String()
}

func (s *InitializeEHPCResponse) SetHeaders(v map[string]*string) *InitializeEHPCResponse {
	s.Headers = v
	return s
}

func (s *InitializeEHPCResponse) SetStatusCode(v int32) *InitializeEHPCResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeEHPCResponse) SetBody(v *InitializeEHPCResponseBody) *InitializeEHPCResponse {
	s.Body = v
	return s
}

type InspectImageRequest struct {
	// The ID of the E-HPC cluster where the image whose Inspect information you want to view resides.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The container type of the image. Set the value to singularity.
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
	// The name of the image whose Inspect information you want to view.
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
}

func (s InspectImageRequest) String() string {
	return tea.Prettify(s)
}

func (s InspectImageRequest) GoString() string {
	return s.String()
}

func (s *InspectImageRequest) SetClusterId(v string) *InspectImageRequest {
	s.ClusterId = &v
	return s
}

func (s *InspectImageRequest) SetContainerType(v string) *InspectImageRequest {
	s.ContainerType = &v
	return s
}

func (s *InspectImageRequest) SetImageName(v string) *InspectImageRequest {
	s.ImageName = &v
	return s
}

type InspectImageResponseBody struct {
	// The status of the image.
	ImageStatus *InspectImageResponseBodyImageStatus `json:"ImageStatus,omitempty" xml:"ImageStatus,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InspectImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InspectImageResponseBody) GoString() string {
	return s.String()
}

func (s *InspectImageResponseBody) SetImageStatus(v *InspectImageResponseBodyImageStatus) *InspectImageResponseBody {
	s.ImageStatus = v
	return s
}

func (s *InspectImageResponseBody) SetRequestId(v string) *InspectImageResponseBody {
	s.RequestId = &v
	return s
}

type InspectImageResponseBodyImageStatus struct {
	// The list of Inspect information about the image.
	ImageInspectInfo *InspectImageResponseBodyImageStatusImageInspectInfo `json:"ImageInspectInfo,omitempty" xml:"ImageInspectInfo,omitempty" type:"Struct"`
}

func (s InspectImageResponseBodyImageStatus) String() string {
	return tea.Prettify(s)
}

func (s InspectImageResponseBodyImageStatus) GoString() string {
	return s.String()
}

func (s *InspectImageResponseBodyImageStatus) SetImageInspectInfo(v *InspectImageResponseBodyImageStatusImageInspectInfo) *InspectImageResponseBodyImageStatus {
	s.ImageInspectInfo = v
	return s
}

type InspectImageResponseBodyImageStatusImageInspectInfo struct {
	// The version of the bootstrapper used by the container image.
	BootStrap *string `json:"BootStrap,omitempty" xml:"BootStrap,omitempty"`
	// The architecture used to build the image.
	BuildArch *string `json:"BuildArch,omitempty" xml:"BuildArch,omitempty"`
	// The date on which the image was built.
	BuildDate *string `json:"BuildDate,omitempty" xml:"BuildDate,omitempty"`
	// The container version of the image.
	ContainerVersion *string `json:"ContainerVersion,omitempty" xml:"ContainerVersion,omitempty"`
	// The mode in which the image was built.
	DefFrom *string `json:"DefFrom,omitempty" xml:"DefFrom,omitempty"`
	// The singularity version and kernel version of the image.
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s InspectImageResponseBodyImageStatusImageInspectInfo) String() string {
	return tea.Prettify(s)
}

func (s InspectImageResponseBodyImageStatusImageInspectInfo) GoString() string {
	return s.String()
}

func (s *InspectImageResponseBodyImageStatusImageInspectInfo) SetBootStrap(v string) *InspectImageResponseBodyImageStatusImageInspectInfo {
	s.BootStrap = &v
	return s
}

func (s *InspectImageResponseBodyImageStatusImageInspectInfo) SetBuildArch(v string) *InspectImageResponseBodyImageStatusImageInspectInfo {
	s.BuildArch = &v
	return s
}

func (s *InspectImageResponseBodyImageStatusImageInspectInfo) SetBuildDate(v string) *InspectImageResponseBodyImageStatusImageInspectInfo {
	s.BuildDate = &v
	return s
}

func (s *InspectImageResponseBodyImageStatusImageInspectInfo) SetContainerVersion(v string) *InspectImageResponseBodyImageStatusImageInspectInfo {
	s.ContainerVersion = &v
	return s
}

func (s *InspectImageResponseBodyImageStatusImageInspectInfo) SetDefFrom(v string) *InspectImageResponseBodyImageStatusImageInspectInfo {
	s.DefFrom = &v
	return s
}

func (s *InspectImageResponseBodyImageStatusImageInspectInfo) SetSchemaVersion(v string) *InspectImageResponseBodyImageStatusImageInspectInfo {
	s.SchemaVersion = &v
	return s
}

type InspectImageResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InspectImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InspectImageResponse) String() string {
	return tea.Prettify(s)
}

func (s InspectImageResponse) GoString() string {
	return s.String()
}

func (s *InspectImageResponse) SetHeaders(v map[string]*string) *InspectImageResponse {
	s.Headers = v
	return s
}

func (s *InspectImageResponse) SetStatusCode(v int32) *InspectImageResponse {
	s.StatusCode = &v
	return s
}

func (s *InspectImageResponse) SetBody(v *InspectImageResponseBody) *InspectImageResponse {
	s.Body = v
	return s
}

type InstallSoftwareRequest struct {
	// The name of the software that you want to install.
	//
	// You can call the [ListSoftwares](~~87216~~) operation to query the software that can be installed.
	Application *string `json:"Application,omitempty" xml:"Application,omitempty"`
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s InstallSoftwareRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallSoftwareRequest) GoString() string {
	return s.String()
}

func (s *InstallSoftwareRequest) SetApplication(v string) *InstallSoftwareRequest {
	s.Application = &v
	return s
}

func (s *InstallSoftwareRequest) SetClusterId(v string) *InstallSoftwareRequest {
	s.ClusterId = &v
	return s
}

type InstallSoftwareResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallSoftwareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallSoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *InstallSoftwareResponseBody) SetRequestId(v string) *InstallSoftwareResponseBody {
	s.RequestId = &v
	return s
}

type InstallSoftwareResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InstallSoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallSoftwareResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallSoftwareResponse) GoString() string {
	return s.String()
}

func (s *InstallSoftwareResponse) SetHeaders(v map[string]*string) *InstallSoftwareResponse {
	s.Headers = v
	return s
}

func (s *InstallSoftwareResponse) SetStatusCode(v int32) *InstallSoftwareResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallSoftwareResponse) SetBody(v *InstallSoftwareResponseBody) *InstallSoftwareResponse {
	s.Body = v
	return s
}

type InvokeShellCommandRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The content of the command. The value must be 2 to 2,048 characters in length.
	Command  *string                              `json:"Command,omitempty" xml:"Command,omitempty"`
	Instance []*InvokeShellCommandRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
	// The timeout period. If a command times out, the command process is terminated. Unit: seconds.
	//
	// Default value: 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The working directory of the command. Default value: /root.
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s InvokeShellCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeShellCommandRequest) GoString() string {
	return s.String()
}

func (s *InvokeShellCommandRequest) SetClusterId(v string) *InvokeShellCommandRequest {
	s.ClusterId = &v
	return s
}

func (s *InvokeShellCommandRequest) SetCommand(v string) *InvokeShellCommandRequest {
	s.Command = &v
	return s
}

func (s *InvokeShellCommandRequest) SetInstance(v []*InvokeShellCommandRequestInstance) *InvokeShellCommandRequest {
	s.Instance = v
	return s
}

func (s *InvokeShellCommandRequest) SetTimeout(v int32) *InvokeShellCommandRequest {
	s.Timeout = &v
	return s
}

func (s *InvokeShellCommandRequest) SetWorkingDir(v string) *InvokeShellCommandRequest {
	s.WorkingDir = &v
	return s
}

type InvokeShellCommandRequestInstance struct {
	// The ID of the node on which the command is run.
	//
	// >  The Instance.N.Id parameter specifies the node on which the command is run. If it is not specified, the command is run on all nodes of the cluster.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s InvokeShellCommandRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s InvokeShellCommandRequestInstance) GoString() string {
	return s.String()
}

func (s *InvokeShellCommandRequestInstance) SetId(v string) *InvokeShellCommandRequestInstance {
	s.Id = &v
	return s
}

type InvokeShellCommandResponseBody struct {
	// The ID of the command. It is used to query the running status of the command.
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The list of IDs of the instances on which you want to run the command.
	InstanceIds *InvokeShellCommandResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InvokeShellCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeShellCommandResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeShellCommandResponseBody) SetCommandId(v string) *InvokeShellCommandResponseBody {
	s.CommandId = &v
	return s
}

func (s *InvokeShellCommandResponseBody) SetInstanceIds(v *InvokeShellCommandResponseBodyInstanceIds) *InvokeShellCommandResponseBody {
	s.InstanceIds = v
	return s
}

func (s *InvokeShellCommandResponseBody) SetRequestId(v string) *InvokeShellCommandResponseBody {
	s.RequestId = &v
	return s
}

type InvokeShellCommandResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s InvokeShellCommandResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s InvokeShellCommandResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *InvokeShellCommandResponseBodyInstanceIds) SetInstanceId(v []*string) *InvokeShellCommandResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

type InvokeShellCommandResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InvokeShellCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvokeShellCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeShellCommandResponse) GoString() string {
	return s.String()
}

func (s *InvokeShellCommandResponse) SetHeaders(v map[string]*string) *InvokeShellCommandResponse {
	s.Headers = v
	return s
}

func (s *InvokeShellCommandResponse) SetStatusCode(v int32) *InvokeShellCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeShellCommandResponse) SetBody(v *InvokeShellCommandResponseBody) *InvokeShellCommandResponse {
	s.Body = v
	return s
}

type ListAvailableEcsTypesRequest struct {
	// The billing method of the ECS instances. Valid values:
	//
	// *   PostPaid: pay-as-you-go
	// *   PrePaid: subscription
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// Specifies whether the ECS instances are sold out. Valid values:
	//
	// *   false: available
	// *   true: sold out
	//
	// Default value: false
	ShowSoldOut *bool `json:"ShowSoldOut,omitempty" xml:"ShowSoldOut,omitempty"`
	// The preemption policy of the ECS instances. Valid values:
	//
	// *   NoSpot: The ECS instances are pay-as-you-go instances.
	// *   SpotWithPriceLimit: The ECS instances are preemptible instances that have a user-defined maximum hourly price.
	// *   SpotAsPriceGo: The ECS instances are preemptible instances for which the market price at the time of purchase is used as the bid price.
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The zone ID.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListAvailableEcsTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesRequest) SetInstanceChargeType(v string) *ListAvailableEcsTypesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *ListAvailableEcsTypesRequest) SetShowSoldOut(v bool) *ListAvailableEcsTypesRequest {
	s.ShowSoldOut = &v
	return s
}

func (s *ListAvailableEcsTypesRequest) SetSpotStrategy(v string) *ListAvailableEcsTypesRequest {
	s.SpotStrategy = &v
	return s
}

func (s *ListAvailableEcsTypesRequest) SetZoneId(v string) *ListAvailableEcsTypesRequest {
	s.ZoneId = &v
	return s
}

type ListAvailableEcsTypesResponseBody struct {
	// The instance family to which the instance type belongs.
	InstanceTypeFamilies *ListAvailableEcsTypesResponseBodyInstanceTypeFamilies `json:"InstanceTypeFamilies,omitempty" xml:"InstanceTypeFamilies,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specifies whether preemptible instances are supported. Valid values:
	//
	// *   false: not supported
	// *   true: supported
	SupportSpotInstance *bool `json:"SupportSpotInstance,omitempty" xml:"SupportSpotInstance,omitempty"`
}

func (s ListAvailableEcsTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponseBody) SetInstanceTypeFamilies(v *ListAvailableEcsTypesResponseBodyInstanceTypeFamilies) *ListAvailableEcsTypesResponseBody {
	s.InstanceTypeFamilies = v
	return s
}

func (s *ListAvailableEcsTypesResponseBody) SetRequestId(v string) *ListAvailableEcsTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBody) SetSupportSpotInstance(v bool) *ListAvailableEcsTypesResponseBody {
	s.SupportSpotInstance = &v
	return s
}

type ListAvailableEcsTypesResponseBodyInstanceTypeFamilies struct {
	InstanceTypeFamilyInfo []*ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo `json:"InstanceTypeFamilyInfo,omitempty" xml:"InstanceTypeFamilyInfo,omitempty" type:"Repeated"`
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamilies) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamilies) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamilies) SetInstanceTypeFamilyInfo(v []*ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo) *ListAvailableEcsTypesResponseBodyInstanceTypeFamilies {
	s.InstanceTypeFamilyInfo = v
	return s
}

type ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo struct {
	// The instance family.
	Generation *string `json:"Generation,omitempty" xml:"Generation,omitempty"`
	// The ID of the instance family. For more information, see [Instance families](~~25378~~).
	InstanceTypeFamilyId *string `json:"InstanceTypeFamilyId,omitempty" xml:"InstanceTypeFamilyId,omitempty"`
	// The list of instance types.
	Types *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypes `json:"Types,omitempty" xml:"Types,omitempty" type:"Struct"`
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo) SetGeneration(v string) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo {
	s.Generation = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo) SetInstanceTypeFamilyId(v string) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo {
	s.InstanceTypeFamilyId = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo) SetTypes(v *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypes) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo {
	s.Types = v
	return s
}

type ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypes struct {
	TypesInfo []*ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo `json:"TypesInfo,omitempty" xml:"TypesInfo,omitempty" type:"Repeated"`
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypes) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypes) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypes) SetTypesInfo(v []*ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypes {
	s.TypesInfo = v
	return s
}

type ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo struct {
	// The number of vCPUs.
	CpuCoreCount *int32 `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty"`
	// The maximum number of elastic network interfaces (ENIs) that can be bound to an ECS instance.
	EniQuantity *int32 `json:"EniQuantity,omitempty" xml:"EniQuantity,omitempty"`
	// The number of GPUs of an ECS instance.
	GPUAmount *int32 `json:"GPUAmount,omitempty" xml:"GPUAmount,omitempty"`
	// The GPU type of the ECS instance.
	GPUSpec *string `json:"GPUSpec,omitempty" xml:"GPUSpec,omitempty"`
	// The maximum inbound internal bandwidth. Unit: Kbit/s.
	InstanceBandwidthRx *int32 `json:"InstanceBandwidthRx,omitempty" xml:"InstanceBandwidthRx,omitempty"`
	// The maximum outbound internal bandwidth. Unit: Kbit/s.
	InstanceBandwidthTx *int32 `json:"InstanceBandwidthTx,omitempty" xml:"InstanceBandwidthTx,omitempty"`
	// The inbound packet forwarding rate over the internal network. Unit: pps
	InstancePpsRx *int32 `json:"InstancePpsRx,omitempty" xml:"InstancePpsRx,omitempty"`
	// The outbound packet forwarding rate over the internal network. Unit: pps
	InstancePpsTx *int32 `json:"InstancePpsTx,omitempty" xml:"InstancePpsTx,omitempty"`
	// The ID of the ECS instance type.
	InstanceTypeId *string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty"`
	// The memory size of the ECS instance. Unit: GiB
	MemorySize *int32 `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	// The status of the ECS instance. Valid values:
	//
	// *   SoldOut
	// *   Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of zone IDs.
	ZoneIds *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfoZoneIds `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty" type:"Struct"`
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetCpuCoreCount(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.CpuCoreCount = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetEniQuantity(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.EniQuantity = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetGPUAmount(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.GPUAmount = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetGPUSpec(v string) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.GPUSpec = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetInstanceBandwidthRx(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.InstanceBandwidthRx = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetInstanceBandwidthTx(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.InstanceBandwidthTx = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetInstancePpsRx(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.InstancePpsRx = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetInstancePpsTx(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.InstancePpsTx = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetInstanceTypeId(v string) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.InstanceTypeId = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetMemorySize(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.MemorySize = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetStatus(v string) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.Status = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetZoneIds(v *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfoZoneIds) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.ZoneIds = v
	return s
}

type ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfoZoneIds struct {
	ZoneId []*string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" type:"Repeated"`
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfoZoneIds) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfoZoneIds) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfoZoneIds) SetZoneId(v []*string) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfoZoneIds {
	s.ZoneId = v
	return s
}

type ListAvailableEcsTypesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAvailableEcsTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAvailableEcsTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponse) SetHeaders(v map[string]*string) *ListAvailableEcsTypesResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableEcsTypesResponse) SetStatusCode(v int32) *ListAvailableEcsTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableEcsTypesResponse) SetBody(v *ListAvailableEcsTypesResponseBody) *ListAvailableEcsTypesResponse {
	s.Body = v
	return s
}

type ListCloudMetricProfilingsRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50
	//
	// Default value: 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCloudMetricProfilingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCloudMetricProfilingsRequest) GoString() string {
	return s.String()
}

func (s *ListCloudMetricProfilingsRequest) SetClusterId(v string) *ListCloudMetricProfilingsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCloudMetricProfilingsRequest) SetPageNumber(v int32) *ListCloudMetricProfilingsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCloudMetricProfilingsRequest) SetPageSize(v int32) *ListCloudMetricProfilingsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudMetricProfilingsRequest) SetRegionId(v string) *ListCloudMetricProfilingsRequest {
	s.RegionId = &v
	return s
}

type ListCloudMetricProfilingsResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries that are returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The profiling information of a specified cluster.
	Profilings *ListCloudMetricProfilingsResponseBodyProfilings `json:"Profilings,omitempty" xml:"Profilings,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudMetricProfilingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCloudMetricProfilingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudMetricProfilingsResponseBody) SetPageNumber(v int32) *ListCloudMetricProfilingsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBody) SetPageSize(v int32) *ListCloudMetricProfilingsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBody) SetProfilings(v *ListCloudMetricProfilingsResponseBodyProfilings) *ListCloudMetricProfilingsResponseBody {
	s.Profilings = v
	return s
}

func (s *ListCloudMetricProfilingsResponseBody) SetRequestId(v string) *ListCloudMetricProfilingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBody) SetTotalCount(v int32) *ListCloudMetricProfilingsResponseBody {
	s.TotalCount = &v
	return s
}

type ListCloudMetricProfilingsResponseBodyProfilings struct {
	ProfilingInfo []*ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo `json:"ProfilingInfo,omitempty" xml:"ProfilingInfo,omitempty" type:"Repeated"`
}

func (s ListCloudMetricProfilingsResponseBodyProfilings) String() string {
	return tea.Prettify(s)
}

func (s ListCloudMetricProfilingsResponseBodyProfilings) GoString() string {
	return s.String()
}

func (s *ListCloudMetricProfilingsResponseBodyProfilings) SetProfilingInfo(v []*ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) *ListCloudMetricProfilingsResponseBodyProfilings {
	s.ProfilingInfo = v
	return s
}

type ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo struct {
	// The duration of the profiling process. Unit: seconds
	//
	// Valid values: 10 to 300
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The frequency of the profiling process. Unit: Hz
	//
	// Valid values: 1 to 2000
	Freq *int32 `json:"Freq,omitempty" xml:"Freq,omitempty"`
	// The name of the host.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the node.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the profiling process.
	Pid *int32 `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The ID of the profiling process.
	ProfilingId *string `json:"ProfilingId,omitempty" xml:"ProfilingId,omitempty"`
	// The time when the profiling process is triggered.
	TriggerTime *string `json:"TriggerTime,omitempty" xml:"TriggerTime,omitempty"`
}

func (s ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) GoString() string {
	return s.String()
}

func (s *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) SetDuration(v int32) *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo {
	s.Duration = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) SetFreq(v int32) *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo {
	s.Freq = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) SetHostName(v string) *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo {
	s.HostName = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) SetInstanceId(v string) *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo {
	s.InstanceId = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) SetPid(v int32) *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo {
	s.Pid = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) SetProfilingId(v string) *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo {
	s.ProfilingId = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) SetTriggerTime(v string) *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo {
	s.TriggerTime = &v
	return s
}

type ListCloudMetricProfilingsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCloudMetricProfilingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCloudMetricProfilingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCloudMetricProfilingsResponse) GoString() string {
	return s.String()
}

func (s *ListCloudMetricProfilingsResponse) SetHeaders(v map[string]*string) *ListCloudMetricProfilingsResponse {
	s.Headers = v
	return s
}

func (s *ListCloudMetricProfilingsResponse) SetStatusCode(v int32) *ListCloudMetricProfilingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudMetricProfilingsResponse) SetBody(v *ListCloudMetricProfilingsResponseBody) *ListCloudMetricProfilingsResponse {
	s.Body = v
	return s
}

type ListClusterLogsRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of the page to return. Pages start from page 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100
	//
	// Default: 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClusterLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsRequest) GoString() string {
	return s.String()
}

func (s *ListClusterLogsRequest) SetClusterId(v string) *ListClusterLogsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterLogsRequest) SetPageNumber(v int32) *ListClusterLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClusterLogsRequest) SetPageSize(v int32) *ListClusterLogsRequest {
	s.PageSize = &v
	return s
}

type ListClusterLogsResponseBody struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The details about operations logs.
	Logs *ListClusterLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Struct"`
	// The number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries that are returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClusterLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterLogsResponseBody) SetClusterId(v string) *ListClusterLogsResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ListClusterLogsResponseBody) SetLogs(v *ListClusterLogsResponseBodyLogs) *ListClusterLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListClusterLogsResponseBody) SetPageNumber(v int32) *ListClusterLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClusterLogsResponseBody) SetPageSize(v int32) *ListClusterLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClusterLogsResponseBody) SetRequestId(v string) *ListClusterLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterLogsResponseBody) SetTotalCount(v int32) *ListClusterLogsResponseBody {
	s.TotalCount = &v
	return s
}

type ListClusterLogsResponseBodyLogs struct {
	LogInfo []*ListClusterLogsResponseBodyLogsLogInfo `json:"LogInfo,omitempty" xml:"LogInfo,omitempty" type:"Repeated"`
}

func (s ListClusterLogsResponseBodyLogs) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListClusterLogsResponseBodyLogs) SetLogInfo(v []*ListClusterLogsResponseBodyLogsLogInfo) *ListClusterLogsResponseBodyLogs {
	s.LogInfo = v
	return s
}

type ListClusterLogsResponseBodyLogsLogInfo struct {
	// The time when the log was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The severity level of the log entry.
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The content of the log entry.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the operation.
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s ListClusterLogsResponseBodyLogsLogInfo) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsResponseBodyLogsLogInfo) GoString() string {
	return s.String()
}

func (s *ListClusterLogsResponseBodyLogsLogInfo) SetCreateTime(v string) *ListClusterLogsResponseBodyLogsLogInfo {
	s.CreateTime = &v
	return s
}

func (s *ListClusterLogsResponseBodyLogsLogInfo) SetLevel(v string) *ListClusterLogsResponseBodyLogsLogInfo {
	s.Level = &v
	return s
}

func (s *ListClusterLogsResponseBodyLogsLogInfo) SetMessage(v string) *ListClusterLogsResponseBodyLogsLogInfo {
	s.Message = &v
	return s
}

func (s *ListClusterLogsResponseBodyLogsLogInfo) SetOperation(v string) *ListClusterLogsResponseBodyLogsLogInfo {
	s.Operation = &v
	return s
}

type ListClusterLogsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClusterLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsResponse) GoString() string {
	return s.String()
}

func (s *ListClusterLogsResponse) SetHeaders(v map[string]*string) *ListClusterLogsResponse {
	s.Headers = v
	return s
}

func (s *ListClusterLogsResponse) SetStatusCode(v int32) *ListClusterLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterLogsResponse) SetBody(v *ListClusterLogsResponseBody) *ListClusterLogsResponse {
	s.Body = v
	return s
}

type ListClustersRequest struct {
	// The number of the page to return. Pages start from page 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50.
	//
	// Default value: 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) SetPageNumber(v int32) *ListClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClustersRequest) SetPageSize(v int32) *ListClustersRequest {
	s.PageSize = &v
	return s
}

type ListClustersResponseBody struct {
	// The list of clusters.
	Clusters *ListClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Struct"`
	// The number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) SetClusters(v *ListClustersResponseBodyClusters) *ListClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *ListClustersResponseBody) SetPageNumber(v int32) *ListClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClustersResponseBody) SetPageSize(v int32) *ListClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetTotalCount(v int32) *ListClustersResponseBody {
	s.TotalCount = &v
	return s
}

type ListClustersResponseBodyClusters struct {
	ClusterInfoSimple []*ListClustersResponseBodyClustersClusterInfoSimple `json:"ClusterInfoSimple,omitempty" xml:"ClusterInfoSimple,omitempty" type:"Repeated"`
}

func (s ListClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClusters) SetClusterInfoSimple(v []*ListClustersResponseBodyClustersClusterInfoSimple) *ListClustersResponseBodyClusters {
	s.ClusterInfoSimple = v
	return s
}

type ListClustersResponseBodyClustersClusterInfoSimple struct {
	// The server type of the account. Valid values:
	//
	// *   nis
	// *   ldap
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The operating system tag of the base image. The tag was used only by the management node.
	BaseOsTag *string `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	// The version of the client.
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The maximum hourly price for the ECS instance under the compute node. The return value can be accurate to three decimal places.
	ComputeSpotPriceLimit *float32 `json:"ComputeSpotPriceLimit,omitempty" xml:"ComputeSpotPriceLimit,omitempty"`
	// The bidding method of the compute nodes. Valid values:
	//
	// *   NoSpot: The instances of the compute node are pay-as-you-go instances.
	// *   SpotWithPriceLimit: The instances of the compute node are preemptible instances. These types of instances have a specified maximum hourly price.
	// *   SpotAsPriceGo: The instances of the compute node are preemptible instances. The price of these instances is based on the current market price.
	ComputeSpotStrategy *string `json:"ComputeSpotStrategy,omitempty" xml:"ComputeSpotStrategy,omitempty"`
	// The information about compute nodes.
	Computes *ListClustersResponseBodyClustersClusterInfoSimpleComputes `json:"Computes,omitempty" xml:"Computes,omitempty" type:"Struct"`
	// The number of compute nodes in the cluster.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The time when the instance was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The mode in which the cluster was deployed. Valid values:
	//
	// *   Standard: An account node, a scheduling node, a logon node, and multiple compute nodes are separately deployed.
	// *   Advanced: Two high availability (HA) account nodes, two HA scheduler nodes, one logon node, and multiple compute nodes are separately deployed.
	// *   Simple: A management node, a logon node, and multiple compute nodes are deployed. The management node consists of an account node and a scheduling node. The logon node and compute nodes are separately deployed.
	// *   Tiny: A management node and multiple compute nodes are deployed. The management node consists of an account node, a scheduling node, and a logon node. The compute nodes are separately deployed.
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// The description of the cluster.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The version of E-HPC.
	EhpcVersion *string `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	// Indicates whether plug-ins were used in the cluster. Valid values:
	//
	// *   true: Plug-ins are used.
	// *   false: Plug-ins are not used.
	//
	// Default value: false
	HasPlugin *bool `json:"HasPlugin,omitempty" xml:"HasPlugin,omitempty"`
	// The ID of the cluster.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The type of the image. Valid values:
	//
	// *   system: public image
	// *   self: custom image
	// *   others: shared image
	// *   marketplace: Alibaba Cloud Marketplace image
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The billing method of the nodes in the cluster. Valid values:
	//
	// *   PostPaid: pay-as-you-go
	// *   PrePaid: subscription
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance type of the compute nodes.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Indicates whether a scaling group was enabled. Valid values:
	//
	// *   true: A scaling group is enabled.
	// *   false: No scaling group is enabled.
	IsComputeEss *bool `json:"IsComputeEss,omitempty" xml:"IsComputeEss,omitempty"`
	// The location where the cluster was deployed. Valid values:
	//
	// *   OnPremise: The cluster is deployed on a hybrid cloud.
	// *   PublicCloud: The cluster is deployed on a public cloud.
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The list of logon nodes.
	LoginNodes *string `json:"LoginNodes,omitempty" xml:"LoginNodes,omitempty"`
	// The list of management nodes.
	Managers *ListClustersResponseBodyClustersClusterInfoSimpleManagers `json:"Managers,omitempty" xml:"Managers,omitempty" type:"Struct"`
	// The name of the cluster.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The prefix of the node.
	NodePrefix *string `json:"NodePrefix,omitempty" xml:"NodePrefix,omitempty"`
	// The suffix of the node.
	NodeSuffix *string `json:"NodeSuffix,omitempty" xml:"NodeSuffix,omitempty"`
	// The operating system tag of the image.
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the scheduler. Valid values:
	//
	// *   pbs
	// *   slurm
	// *   opengridscheduler
	// *   deadline
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	// The status of the cluster. Valid values:
	//
	// *   uninit: The cluster is not initialized.
	// *   creating: The cluster is being created.
	// *   init: The cluster is being initialized.
	// *   running: The cluster is running.
	// *   exception: The cluster encounters an exception.
	// *   releasing: The cluster is being released.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The statistics of all resources in the cluster.
	TotalResources *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty" type:"Struct"`
	// The number of consumed resources in the cluster.
	UsedResources *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources `json:"UsedResources,omitempty" xml:"UsedResources,omitempty" type:"Struct"`
	// The ID of the vSwitch.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoSimple) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoSimple) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetAccountType(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.AccountType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetBaseOsTag(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.BaseOsTag = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetClientVersion(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.ClientVersion = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetComputeSpotPriceLimit(v float32) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.ComputeSpotPriceLimit = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetComputeSpotStrategy(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.ComputeSpotStrategy = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetComputes(v *ListClustersResponseBodyClustersClusterInfoSimpleComputes) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Computes = v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Count = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetCreateTime(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.CreateTime = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetDeployMode(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.DeployMode = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetDescription(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Description = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetEhpcVersion(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.EhpcVersion = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetHasPlugin(v bool) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.HasPlugin = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetId(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Id = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetImageId(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.ImageId = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetImageOwnerAlias(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetInstanceChargeType(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.InstanceChargeType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetInstanceType(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.InstanceType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetIsComputeEss(v bool) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.IsComputeEss = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetLocation(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Location = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetLoginNodes(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.LoginNodes = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetManagers(v *ListClustersResponseBodyClustersClusterInfoSimpleManagers) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Managers = v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetName(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetNodePrefix(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.NodePrefix = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetNodeSuffix(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.NodeSuffix = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetOsTag(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.OsTag = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetRegionId(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.RegionId = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetSchedulerType(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.SchedulerType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetStatus(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Status = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetTotalResources(v *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.TotalResources = v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetUsedResources(v *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.UsedResources = v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetVSwitchId(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.VSwitchId = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetVpcId(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.VpcId = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetZoneId(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.ZoneId = &v
	return s
}

type ListClustersResponseBodyClustersClusterInfoSimpleComputes struct {
	// The number of abnormal nodes.
	ExceptionCount *int32 `json:"ExceptionCount,omitempty" xml:"ExceptionCount,omitempty"`
	// The number of normal nodes.
	NormalCount *int32 `json:"NormalCount,omitempty" xml:"NormalCount,omitempty"`
	// The number of nodes that are being used in the queue. This includes those that are being initialized, installed, or released.
	OperatingCount *int32 `json:"OperatingCount,omitempty" xml:"OperatingCount,omitempty"`
	// The number of stopped nodes.
	StoppedCount *int32 `json:"StoppedCount,omitempty" xml:"StoppedCount,omitempty"`
	// The total number of nodes.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleComputes) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleComputes) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleComputes) SetExceptionCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleComputes {
	s.ExceptionCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleComputes) SetNormalCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleComputes {
	s.NormalCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleComputes) SetOperatingCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleComputes {
	s.OperatingCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleComputes) SetStoppedCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleComputes {
	s.StoppedCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleComputes) SetTotal(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleComputes {
	s.Total = &v
	return s
}

type ListClustersResponseBodyClustersClusterInfoSimpleManagers struct {
	// The number of abnormal nodes.
	ExceptionCount *int32 `json:"ExceptionCount,omitempty" xml:"ExceptionCount,omitempty"`
	// The number of normal nodes.
	NormalCount *int32 `json:"NormalCount,omitempty" xml:"NormalCount,omitempty"`
	// The number of nodes that are being used in the queue. This includes those that are being initialized, installed, or released.
	OperatingCount *int32 `json:"OperatingCount,omitempty" xml:"OperatingCount,omitempty"`
	// The number of stopped nodes.
	StoppedCount *int32 `json:"StoppedCount,omitempty" xml:"StoppedCount,omitempty"`
	// The total number of management nodes.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleManagers) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleManagers) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleManagers) SetExceptionCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleManagers {
	s.ExceptionCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleManagers) SetNormalCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleManagers {
	s.NormalCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleManagers) SetOperatingCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleManagers {
	s.OperatingCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleManagers) SetStoppedCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleManagers {
	s.StoppedCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleManagers) SetTotal(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleManagers {
	s.Total = &v
	return s
}

type ListClustersResponseBodyClustersClusterInfoSimpleTotalResources struct {
	// The number of CPU cores. Unit: cores.
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of GPU cards. Unit: cards.
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The memory size. Unit: MiB.
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) SetCpu(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources {
	s.Cpu = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) SetGpu(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources {
	s.Gpu = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) SetMemory(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources {
	s.Memory = &v
	return s
}

type ListClustersResponseBodyClustersClusterInfoSimpleUsedResources struct {
	// The number of CPU cores. Unit: cores.
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of GPU cards. Unit: cards.
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The memory size. Unit: MiB.
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) SetCpu(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources {
	s.Cpu = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) SetGpu(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources {
	s.Gpu = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) SetMemory(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources {
	s.Memory = &v
	return s
}

type ListClustersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponse) GoString() string {
	return s.String()
}

func (s *ListClustersResponse) SetHeaders(v map[string]*string) *ListClustersResponse {
	s.Headers = v
	return s
}

func (s *ListClustersResponse) SetStatusCode(v int32) *ListClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClustersResponse) SetBody(v *ListClustersResponseBody) *ListClustersResponse {
	s.Body = v
	return s
}

type ListClustersMetaRequest struct {
	// The number of the page to return. Pages start from page 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50.
	//
	// Default value: 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClustersMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersMetaRequest) GoString() string {
	return s.String()
}

func (s *ListClustersMetaRequest) SetPageNumber(v int32) *ListClustersMetaRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClustersMetaRequest) SetPageSize(v int32) *ListClustersMetaRequest {
	s.PageSize = &v
	return s
}

type ListClustersMetaResponseBody struct {
	// The list of clusters.
	Clusters *ListClustersMetaResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Struct"`
	// The number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values: 1 to 50.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClustersMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersMetaResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersMetaResponseBody) SetClusters(v *ListClustersMetaResponseBodyClusters) *ListClustersMetaResponseBody {
	s.Clusters = v
	return s
}

func (s *ListClustersMetaResponseBody) SetPageNumber(v int32) *ListClustersMetaResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClustersMetaResponseBody) SetPageSize(v int32) *ListClustersMetaResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClustersMetaResponseBody) SetRequestId(v string) *ListClustersMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersMetaResponseBody) SetTotalCount(v int32) *ListClustersMetaResponseBody {
	s.TotalCount = &v
	return s
}

type ListClustersMetaResponseBodyClusters struct {
	ClusterInfoSimple []*ListClustersMetaResponseBodyClustersClusterInfoSimple `json:"ClusterInfoSimple,omitempty" xml:"ClusterInfoSimple,omitempty" type:"Repeated"`
}

func (s ListClustersMetaResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListClustersMetaResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersMetaResponseBodyClusters) SetClusterInfoSimple(v []*ListClustersMetaResponseBodyClustersClusterInfoSimple) *ListClustersMetaResponseBodyClusters {
	s.ClusterInfoSimple = v
	return s
}

type ListClustersMetaResponseBodyClustersClusterInfoSimple struct {
	// The server type of the account. Valid values:
	//
	// *   nis
	// *   ldap
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The version of the client.
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The mode in which the cluster is deployed. Valid values:
	//
	// *   Standard: An account node, a scheduling node, a logon node, and multiple compute nodes are separately deployed.
	// *   Advanced: Two high availability (HA) account nodes, two HA scheduler nodes, one logon node, and multiple compute nodes are separately deployed.
	// *   Simple: A management node, a logon node, and multiple compute nodes are deployed. The management node consists of an account node and a scheduling node. The logon node and compute nodes are separately deployed.
	// *   Tiny: A management node and multiple compute nodes are deployed. The management node consists of an account node, a scheduling node, and a logon node. The compute nodes are separately deployed.
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// The description of the cluster.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the cluster uses a plug-in. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false
	HasPlugin *bool `json:"HasPlugin,omitempty" xml:"HasPlugin,omitempty"`
	// The ID of the cluster.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether a scaling group is enabled. Valid values:
	//
	// *   true: A scaling group is enabled.
	// *   false: No scaling group is enabled.
	IsComputeEss *bool `json:"IsComputeEss,omitempty" xml:"IsComputeEss,omitempty"`
	// The location where the cluster is deployed. Valid values:
	//
	// *   OnPremise: The cluster is deployed on a hybrid cloud.
	// *   PublicCloud: The cluster is deployed on a public cloud.
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the cluster.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The image tag of the operating system.
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	// The type of the scheduler. Valid values:
	//
	// *   pbs
	// *   slurm
	// *   opengridscheduler
	// *   deadline
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	// The status of the cluster. Valid values:
	//
	// *   uninit: The cluster is not initialized.
	// *   creating: The cluster is being created.
	// *   init: The cluster is being initialized.
	// *   running: The cluster is running.
	// *   exception: The cluster encounters an exception.
	// *   releasing: The cluster is being released.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the virtual private cloud (VPC).
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListClustersMetaResponseBodyClustersClusterInfoSimple) String() string {
	return tea.Prettify(s)
}

func (s ListClustersMetaResponseBodyClustersClusterInfoSimple) GoString() string {
	return s.String()
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetAccountType(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.AccountType = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetClientVersion(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.ClientVersion = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetDeployMode(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.DeployMode = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetDescription(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.Description = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetHasPlugin(v bool) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.HasPlugin = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetId(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.Id = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetIsComputeEss(v bool) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.IsComputeEss = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetLocation(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.Location = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetName(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.Name = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetOsTag(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.OsTag = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetSchedulerType(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.SchedulerType = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetStatus(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.Status = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetVpcId(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.VpcId = &v
	return s
}

type ListClustersMetaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClustersMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClustersMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersMetaResponse) GoString() string {
	return s.String()
}

func (s *ListClustersMetaResponse) SetHeaders(v map[string]*string) *ListClustersMetaResponse {
	s.Headers = v
	return s
}

func (s *ListClustersMetaResponse) SetStatusCode(v int32) *ListClustersMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClustersMetaResponse) SetBody(v *ListClustersMetaResponseBody) *ListClustersMetaResponse {
	s.Body = v
	return s
}

type ListCommandsRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the command.
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The page number of the returned page.
	//
	// Page number starts from page 1.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 50.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCommandsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCommandsRequest) GoString() string {
	return s.String()
}

func (s *ListCommandsRequest) SetClusterId(v string) *ListCommandsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCommandsRequest) SetCommandId(v string) *ListCommandsRequest {
	s.CommandId = &v
	return s
}

func (s *ListCommandsRequest) SetPageNumber(v int32) *ListCommandsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCommandsRequest) SetPageSize(v int32) *ListCommandsRequest {
	s.PageSize = &v
	return s
}

type ListCommandsResponseBody struct {
	// The list of commands.
	Commands *ListCommandsResponseBodyCommands `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on the current page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCommandsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCommandsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommandsResponseBody) SetCommands(v *ListCommandsResponseBodyCommands) *ListCommandsResponseBody {
	s.Commands = v
	return s
}

func (s *ListCommandsResponseBody) SetPageNumber(v int32) *ListCommandsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCommandsResponseBody) SetPageSize(v int32) *ListCommandsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCommandsResponseBody) SetRequestId(v string) *ListCommandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommandsResponseBody) SetTotalCount(v int32) *ListCommandsResponseBody {
	s.TotalCount = &v
	return s
}

type ListCommandsResponseBodyCommands struct {
	Command []*ListCommandsResponseBodyCommandsCommand `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s ListCommandsResponseBodyCommands) String() string {
	return tea.Prettify(s)
}

func (s ListCommandsResponseBodyCommands) GoString() string {
	return s.String()
}

func (s *ListCommandsResponseBodyCommands) SetCommand(v []*ListCommandsResponseBodyCommandsCommand) *ListCommandsResponseBodyCommands {
	s.Command = v
	return s
}

type ListCommandsResponseBodyCommandsCommand struct {
	// The content of the command.
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The ID of the command.
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The timeout period. Unit: seconds.
	Timeout *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The working directory of the command.
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s ListCommandsResponseBodyCommandsCommand) String() string {
	return tea.Prettify(s)
}

func (s ListCommandsResponseBodyCommandsCommand) GoString() string {
	return s.String()
}

func (s *ListCommandsResponseBodyCommandsCommand) SetCommandContent(v string) *ListCommandsResponseBodyCommandsCommand {
	s.CommandContent = &v
	return s
}

func (s *ListCommandsResponseBodyCommandsCommand) SetCommandId(v string) *ListCommandsResponseBodyCommandsCommand {
	s.CommandId = &v
	return s
}

func (s *ListCommandsResponseBodyCommandsCommand) SetTimeout(v string) *ListCommandsResponseBodyCommandsCommand {
	s.Timeout = &v
	return s
}

func (s *ListCommandsResponseBodyCommandsCommand) SetWorkingDir(v string) *ListCommandsResponseBodyCommandsCommand {
	s.WorkingDir = &v
	return s
}

type ListCommandsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCommandsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCommandsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCommandsResponse) GoString() string {
	return s.String()
}

func (s *ListCommandsResponse) SetHeaders(v map[string]*string) *ListCommandsResponse {
	s.Headers = v
	return s
}

func (s *ListCommandsResponse) SetStatusCode(v int32) *ListCommandsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCommandsResponse) SetBody(v *ListCommandsResponseBody) *ListCommandsResponse {
	s.Body = v
	return s
}

type ListCommunityImagesRequest struct {
	// The tag of the base operating system (BOS).
	BaseOsTag *string `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	// The ID of the cluster. If the cluster supports multiple operating systems, all community images in the region where the cluster resides are queried.
	//
	// If you do not specify the cluster ID, the community images that are supported by all clusters are queried.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the Elastic Compute Service (ECS) instance. If you do not specify the instance type, the community images that are supported by all instance types are queried.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ListCommunityImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCommunityImagesRequest) GoString() string {
	return s.String()
}

func (s *ListCommunityImagesRequest) SetBaseOsTag(v string) *ListCommunityImagesRequest {
	s.BaseOsTag = &v
	return s
}

func (s *ListCommunityImagesRequest) SetClusterId(v string) *ListCommunityImagesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCommunityImagesRequest) SetInstanceType(v string) *ListCommunityImagesRequest {
	s.InstanceType = &v
	return s
}

type ListCommunityImagesResponseBody struct {
	// The list of community images, including custom images and shared images.
	Images *ListCommunityImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCommunityImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCommunityImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommunityImagesResponseBody) SetImages(v *ListCommunityImagesResponseBodyImages) *ListCommunityImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListCommunityImagesResponseBody) SetRequestId(v string) *ListCommunityImagesResponseBody {
	s.RequestId = &v
	return s
}

type ListCommunityImagesResponseBodyImages struct {
	ImageInfo []*ListCommunityImagesResponseBodyImagesImageInfo `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" type:"Repeated"`
}

func (s ListCommunityImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListCommunityImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListCommunityImagesResponseBodyImages) SetImageInfo(v []*ListCommunityImagesResponseBodyImagesImageInfo) *ListCommunityImagesResponseBodyImages {
	s.ImageInfo = v
	return s
}

type ListCommunityImagesResponseBodyImagesImageInfo struct {
	// The tag of the BOS image.
	BaseOsTag *ListCommunityImagesResponseBodyImagesImageInfoBaseOsTag `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty" type:"Struct"`
	// The description of the image.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The type of the image. Valid values:
	//
	// *   self: custom image
	// *   others: shared image
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// An array of OS images that are supported by E-HPC.
	OsTag *ListCommunityImagesResponseBodyImagesImageInfoOsTag `json:"OsTag,omitempty" xml:"OsTag,omitempty" type:"Struct"`
	// The script that is run after the image is installed.
	PostInstallScript *string `json:"PostInstallScript,omitempty" xml:"PostInstallScript,omitempty"`
	// The billing unit of the image. Valid values:
	//
	// *   Hour
	// *   Month
	// *   Year
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The Alibaba Cloud Marketplace product code of the image.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The size of the image. Unit: GiB.
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The stock keeping unit (SKU) of the image. Valid values:
	//
	// *   ECS: pay-as-you-go
	// *   package: subscription
	SkuCode *string `json:"SkuCode,omitempty" xml:"SkuCode,omitempty"`
	// The status of the image. Valid values:
	//
	// *   UnAvailable: The image is unavailable.
	// *   Available: The image is available.
	// *   Creating: The image is being created.
	// *   CreateFailed: The image failed to be created.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The owner of the image.
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListCommunityImagesResponseBodyImagesImageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListCommunityImagesResponseBodyImagesImageInfo) GoString() string {
	return s.String()
}

func (s *ListCommunityImagesResponseBodyImagesImageInfo) SetBaseOsTag(v *ListCommunityImagesResponseBodyImagesImageInfoBaseOsTag) *ListCommunityImagesResponseBodyImagesImageInfo {
	s.BaseOsTag = v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfo) SetDescription(v string) *ListCommunityImagesResponseBodyImagesImageInfo {
	s.Description = &v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfo) SetImageId(v string) *ListCommunityImagesResponseBodyImagesImageInfo {
	s.ImageId = &v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfo) SetImageName(v string) *ListCommunityImagesResponseBodyImagesImageInfo {
	s.ImageName = &v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfo) SetImageOwnerAlias(v string) *ListCommunityImagesResponseBodyImagesImageInfo {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfo) SetOsTag(v *ListCommunityImagesResponseBodyImagesImageInfoOsTag) *ListCommunityImagesResponseBodyImagesImageInfo {
	s.OsTag = v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfo) SetPostInstallScript(v string) *ListCommunityImagesResponseBodyImagesImageInfo {
	s.PostInstallScript = &v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfo) SetPricingCycle(v string) *ListCommunityImagesResponseBodyImagesImageInfo {
	s.PricingCycle = &v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfo) SetProductCode(v string) *ListCommunityImagesResponseBodyImagesImageInfo {
	s.ProductCode = &v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfo) SetSize(v int32) *ListCommunityImagesResponseBodyImagesImageInfo {
	s.Size = &v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfo) SetSkuCode(v string) *ListCommunityImagesResponseBodyImagesImageInfo {
	s.SkuCode = &v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfo) SetStatus(v string) *ListCommunityImagesResponseBodyImagesImageInfo {
	s.Status = &v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfo) SetUid(v string) *ListCommunityImagesResponseBodyImagesImageInfo {
	s.Uid = &v
	return s
}

type ListCommunityImagesResponseBodyImagesImageInfoBaseOsTag struct {
	// The architecture of the operating system. Valid values:
	//
	// *   i386
	// *   x86\_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The tag of the OS image.
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	// The distribution of the operating system. Valid values:
	//
	// *   CentOS
	// *   Ubuntu
	// *   SUSE
	// *   OpenSUSE
	// *   Debian
	// *   CoreOS
	// *   Aliyun
	// *   Windows Server 2003
	// *   Windows Server 2008
	// *   Windows Server 2012
	// *   Others Linux
	// *   Customized Linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The version of the operating system.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListCommunityImagesResponseBodyImagesImageInfoBaseOsTag) String() string {
	return tea.Prettify(s)
}

func (s ListCommunityImagesResponseBodyImagesImageInfoBaseOsTag) GoString() string {
	return s.String()
}

func (s *ListCommunityImagesResponseBodyImagesImageInfoBaseOsTag) SetArchitecture(v string) *ListCommunityImagesResponseBodyImagesImageInfoBaseOsTag {
	s.Architecture = &v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfoBaseOsTag) SetOsTag(v string) *ListCommunityImagesResponseBodyImagesImageInfoBaseOsTag {
	s.OsTag = &v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfoBaseOsTag) SetPlatform(v string) *ListCommunityImagesResponseBodyImagesImageInfoBaseOsTag {
	s.Platform = &v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfoBaseOsTag) SetVersion(v string) *ListCommunityImagesResponseBodyImagesImageInfoBaseOsTag {
	s.Version = &v
	return s
}

type ListCommunityImagesResponseBodyImagesImageInfoOsTag struct {
	// The architecture of the operating system. Valid values:
	//
	// *   i386
	// *   x86\_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The tag of the BOS image.
	BaseOsTag *string `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	// The tag of the OS image.
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	// The operating system.
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The version of the operating system.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListCommunityImagesResponseBodyImagesImageInfoOsTag) String() string {
	return tea.Prettify(s)
}

func (s ListCommunityImagesResponseBodyImagesImageInfoOsTag) GoString() string {
	return s.String()
}

func (s *ListCommunityImagesResponseBodyImagesImageInfoOsTag) SetArchitecture(v string) *ListCommunityImagesResponseBodyImagesImageInfoOsTag {
	s.Architecture = &v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfoOsTag) SetBaseOsTag(v string) *ListCommunityImagesResponseBodyImagesImageInfoOsTag {
	s.BaseOsTag = &v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfoOsTag) SetOsTag(v string) *ListCommunityImagesResponseBodyImagesImageInfoOsTag {
	s.OsTag = &v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfoOsTag) SetPlatform(v string) *ListCommunityImagesResponseBodyImagesImageInfoOsTag {
	s.Platform = &v
	return s
}

func (s *ListCommunityImagesResponseBodyImagesImageInfoOsTag) SetVersion(v string) *ListCommunityImagesResponseBodyImagesImageInfoOsTag {
	s.Version = &v
	return s
}

type ListCommunityImagesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCommunityImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCommunityImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCommunityImagesResponse) GoString() string {
	return s.String()
}

func (s *ListCommunityImagesResponse) SetHeaders(v map[string]*string) *ListCommunityImagesResponse {
	s.Headers = v
	return s
}

func (s *ListCommunityImagesResponse) SetStatusCode(v int32) *ListCommunityImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCommunityImagesResponse) SetBody(v *ListCommunityImagesResponseBody) *ListCommunityImagesResponse {
	s.Body = v
	return s
}

type ListContainerAppsRequest struct {
	// The page number of the returned page.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListContainerAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListContainerAppsRequest) GoString() string {
	return s.String()
}

func (s *ListContainerAppsRequest) SetPageNumber(v int32) *ListContainerAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListContainerAppsRequest) SetPageSize(v int32) *ListContainerAppsRequest {
	s.PageSize = &v
	return s
}

type ListContainerAppsResponseBody struct {
	// The array of containerized applications.
	ContainerApps *ListContainerAppsResponseBodyContainerApps `json:"ContainerApps,omitempty" xml:"ContainerApps,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of containerized applications.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListContainerAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListContainerAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContainerAppsResponseBody) SetContainerApps(v *ListContainerAppsResponseBodyContainerApps) *ListContainerAppsResponseBody {
	s.ContainerApps = v
	return s
}

func (s *ListContainerAppsResponseBody) SetPageNumber(v int32) *ListContainerAppsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListContainerAppsResponseBody) SetPageSize(v int32) *ListContainerAppsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListContainerAppsResponseBody) SetRequestId(v string) *ListContainerAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContainerAppsResponseBody) SetTotalCount(v int32) *ListContainerAppsResponseBody {
	s.TotalCount = &v
	return s
}

type ListContainerAppsResponseBodyContainerApps struct {
	ContainerApps []*ListContainerAppsResponseBodyContainerAppsContainerApps `json:"ContainerApps,omitempty" xml:"ContainerApps,omitempty" type:"Repeated"`
}

func (s ListContainerAppsResponseBodyContainerApps) String() string {
	return tea.Prettify(s)
}

func (s ListContainerAppsResponseBodyContainerApps) GoString() string {
	return s.String()
}

func (s *ListContainerAppsResponseBodyContainerApps) SetContainerApps(v []*ListContainerAppsResponseBodyContainerAppsContainerApps) *ListContainerAppsResponseBodyContainerApps {
	s.ContainerApps = v
	return s
}

type ListContainerAppsResponseBodyContainerAppsContainerApps struct {
	// The time when the containerized application was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the containerized application.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the containerized application.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The tags of the image.
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The name of the containerized application.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the repository.
	Repository *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	// The type of the container. Set the value to singularity.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListContainerAppsResponseBodyContainerAppsContainerApps) String() string {
	return tea.Prettify(s)
}

func (s ListContainerAppsResponseBodyContainerAppsContainerApps) GoString() string {
	return s.String()
}

func (s *ListContainerAppsResponseBodyContainerAppsContainerApps) SetCreateTime(v string) *ListContainerAppsResponseBodyContainerAppsContainerApps {
	s.CreateTime = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerAppsContainerApps) SetDescription(v string) *ListContainerAppsResponseBodyContainerAppsContainerApps {
	s.Description = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerAppsContainerApps) SetId(v string) *ListContainerAppsResponseBodyContainerAppsContainerApps {
	s.Id = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerAppsContainerApps) SetImageTag(v string) *ListContainerAppsResponseBodyContainerAppsContainerApps {
	s.ImageTag = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerAppsContainerApps) SetName(v string) *ListContainerAppsResponseBodyContainerAppsContainerApps {
	s.Name = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerAppsContainerApps) SetRepository(v string) *ListContainerAppsResponseBodyContainerAppsContainerApps {
	s.Repository = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerAppsContainerApps) SetType(v string) *ListContainerAppsResponseBodyContainerAppsContainerApps {
	s.Type = &v
	return s
}

type ListContainerAppsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListContainerAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListContainerAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListContainerAppsResponse) GoString() string {
	return s.String()
}

func (s *ListContainerAppsResponse) SetHeaders(v map[string]*string) *ListContainerAppsResponse {
	s.Headers = v
	return s
}

func (s *ListContainerAppsResponse) SetStatusCode(v int32) *ListContainerAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContainerAppsResponse) SetBody(v *ListContainerAppsResponseBody) *ListContainerAppsResponse {
	s.Body = v
	return s
}

type ListContainerImagesRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the container. Set the value to singularity.
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50.
	//
	// Default value: 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListContainerImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesRequest) GoString() string {
	return s.String()
}

func (s *ListContainerImagesRequest) SetClusterId(v string) *ListContainerImagesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListContainerImagesRequest) SetContainerType(v string) *ListContainerImagesRequest {
	s.ContainerType = &v
	return s
}

func (s *ListContainerImagesRequest) SetPageNumber(v int32) *ListContainerImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListContainerImagesRequest) SetPageSize(v int32) *ListContainerImagesRequest {
	s.PageSize = &v
	return s
}

type ListContainerImagesResponseBody struct {
	// The information of the database.
	DBInfo *string `json:"DBInfo,omitempty" xml:"DBInfo,omitempty"`
	// The array of local images.
	Images *ListContainerImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListContainerImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListContainerImagesResponseBody) SetDBInfo(v string) *ListContainerImagesResponseBody {
	s.DBInfo = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetImages(v *ListContainerImagesResponseBodyImages) *ListContainerImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListContainerImagesResponseBody) SetPageNumber(v int32) *ListContainerImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetPageSize(v int32) *ListContainerImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetRequestId(v string) *ListContainerImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetTotalCount(v int32) *ListContainerImagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListContainerImagesResponseBodyImages struct {
	Images []*ListContainerImagesResponseBodyImagesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
}

func (s ListContainerImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListContainerImagesResponseBodyImages) SetImages(v []*ListContainerImagesResponseBodyImagesImages) *ListContainerImagesResponseBodyImages {
	s.Images = v
	return s
}

type ListContainerImagesResponseBodyImagesImages struct {
	// The ID of the image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the repository.
	Repository *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	// The status of the image.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The container system of the image.
	System *string `json:"System,omitempty" xml:"System,omitempty"`
	// The tags of the image.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The type of the container. Set the value to singularity.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the image was updated.
	UpdateDateTime *string `json:"UpdateDateTime,omitempty" xml:"UpdateDateTime,omitempty"`
}

func (s ListContainerImagesResponseBodyImagesImages) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesResponseBodyImagesImages) GoString() string {
	return s.String()
}

func (s *ListContainerImagesResponseBodyImagesImages) SetImageId(v string) *ListContainerImagesResponseBodyImagesImages {
	s.ImageId = &v
	return s
}

func (s *ListContainerImagesResponseBodyImagesImages) SetRepository(v string) *ListContainerImagesResponseBodyImagesImages {
	s.Repository = &v
	return s
}

func (s *ListContainerImagesResponseBodyImagesImages) SetStatus(v string) *ListContainerImagesResponseBodyImagesImages {
	s.Status = &v
	return s
}

func (s *ListContainerImagesResponseBodyImagesImages) SetSystem(v string) *ListContainerImagesResponseBodyImagesImages {
	s.System = &v
	return s
}

func (s *ListContainerImagesResponseBodyImagesImages) SetTag(v string) *ListContainerImagesResponseBodyImagesImages {
	s.Tag = &v
	return s
}

func (s *ListContainerImagesResponseBodyImagesImages) SetType(v string) *ListContainerImagesResponseBodyImagesImages {
	s.Type = &v
	return s
}

func (s *ListContainerImagesResponseBodyImagesImages) SetUpdateDateTime(v string) *ListContainerImagesResponseBodyImagesImages {
	s.UpdateDateTime = &v
	return s
}

type ListContainerImagesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListContainerImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListContainerImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesResponse) GoString() string {
	return s.String()
}

func (s *ListContainerImagesResponse) SetHeaders(v map[string]*string) *ListContainerImagesResponse {
	s.Headers = v
	return s
}

func (s *ListContainerImagesResponse) SetStatusCode(v int32) *ListContainerImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContainerImagesResponse) SetBody(v *ListContainerImagesResponseBody) *ListContainerImagesResponse {
	s.Body = v
	return s
}

type ListCpfsFileSystemsRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCpfsFileSystemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsRequest) SetFileSystemId(v string) *ListCpfsFileSystemsRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListCpfsFileSystemsRequest) SetPageNumber(v int32) *ListCpfsFileSystemsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCpfsFileSystemsRequest) SetPageSize(v int32) *ListCpfsFileSystemsRequest {
	s.PageSize = &v
	return s
}

type ListCpfsFileSystemsResponseBody struct {
	FileSystemList *ListCpfsFileSystemsResponseBodyFileSystemList `json:"FileSystemList,omitempty" xml:"FileSystemList,omitempty" type:"Struct"`
	PageNumber     *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCpfsFileSystemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsResponseBody) SetFileSystemList(v *ListCpfsFileSystemsResponseBodyFileSystemList) *ListCpfsFileSystemsResponseBody {
	s.FileSystemList = v
	return s
}

func (s *ListCpfsFileSystemsResponseBody) SetPageNumber(v int32) *ListCpfsFileSystemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBody) SetPageSize(v int32) *ListCpfsFileSystemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBody) SetRequestId(v string) *ListCpfsFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBody) SetTotalCount(v int32) *ListCpfsFileSystemsResponseBody {
	s.TotalCount = &v
	return s
}

type ListCpfsFileSystemsResponseBodyFileSystemList struct {
	FileSystems []*ListCpfsFileSystemsResponseBodyFileSystemListFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Repeated"`
}

func (s ListCpfsFileSystemsResponseBodyFileSystemList) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsResponseBodyFileSystemList) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemList) SetFileSystems(v []*ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) *ListCpfsFileSystemsResponseBodyFileSystemList {
	s.FileSystems = v
	return s
}

type ListCpfsFileSystemsResponseBodyFileSystemListFileSystems struct {
	Capacity        *string                                                                  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	CreateTime      *string                                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Destription     *string                                                                  `json:"Destription,omitempty" xml:"Destription,omitempty"`
	FileSystemId    *string                                                                  `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountTargetList *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetList `json:"MountTargetList,omitempty" xml:"MountTargetList,omitempty" type:"Struct"`
	ProtocolType    *string                                                                  `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RegionId        *string                                                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId          *string                                                                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) SetCapacity(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems {
	s.Capacity = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) SetCreateTime(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems {
	s.CreateTime = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) SetDestription(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems {
	s.Destription = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) SetFileSystemId(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems {
	s.FileSystemId = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) SetMountTargetList(v *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetList) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems {
	s.MountTargetList = v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) SetProtocolType(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems {
	s.ProtocolType = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) SetRegionId(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems {
	s.RegionId = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) SetZoneId(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems {
	s.ZoneId = &v
	return s
}

type ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetList struct {
	MountTargets []*ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets `json:"MountTargets,omitempty" xml:"MountTargets,omitempty" type:"Repeated"`
}

func (s ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetList) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetList) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetList) SetMountTargets(v []*ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetList {
	s.MountTargets = v
	return s
}

type ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets struct {
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	NetworkType       *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId             *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswId             *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
}

func (s ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetMountTargetDomain(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.MountTargetDomain = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetNetworkType(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.NetworkType = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetStatus(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.Status = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetVpcId(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.VpcId = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetVswId(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.VswId = &v
	return s
}

type ListCpfsFileSystemsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCpfsFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCpfsFileSystemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsResponse) SetHeaders(v map[string]*string) *ListCpfsFileSystemsResponse {
	s.Headers = v
	return s
}

func (s *ListCpfsFileSystemsResponse) SetStatusCode(v int32) *ListCpfsFileSystemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCpfsFileSystemsResponse) SetBody(v *ListCpfsFileSystemsResponseBody) *ListCpfsFileSystemsResponse {
	s.Body = v
	return s
}

type ListCurrentClientVersionResponseBody struct {
	// The latest version number of the E-HPC client.
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCurrentClientVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCurrentClientVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ListCurrentClientVersionResponseBody) SetClientVersion(v string) *ListCurrentClientVersionResponseBody {
	s.ClientVersion = &v
	return s
}

func (s *ListCurrentClientVersionResponseBody) SetRequestId(v string) *ListCurrentClientVersionResponseBody {
	s.RequestId = &v
	return s
}

type ListCurrentClientVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCurrentClientVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCurrentClientVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCurrentClientVersionResponse) GoString() string {
	return s.String()
}

func (s *ListCurrentClientVersionResponse) SetHeaders(v map[string]*string) *ListCurrentClientVersionResponse {
	s.Headers = v
	return s
}

func (s *ListCurrentClientVersionResponse) SetStatusCode(v int32) *ListCurrentClientVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCurrentClientVersionResponse) SetBody(v *ListCurrentClientVersionResponseBody) *ListCurrentClientVersionResponse {
	s.Body = v
	return s
}

type ListCustomImagesRequest struct {
	// The image tag of the base operating system. The tag is used only by the management node.
	BaseOsTag *string `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	// The ID of the cluster where the application resides. If the cluster supports multiple operating systems, all the images in the region where the cluster resides are queried.
	//
	// By default, if you do not specify the cluster ID, the images that are supported by all the clusters is queried.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The source of the image. Valid values:
	//
	// *   self: custom image
	// *   others: shared image
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// Specify the type of the instance. By default, if you do not specify the type of the instance, the list of images that are supported by all the instance types are queried.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ListCustomImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesRequest) GoString() string {
	return s.String()
}

func (s *ListCustomImagesRequest) SetBaseOsTag(v string) *ListCustomImagesRequest {
	s.BaseOsTag = &v
	return s
}

func (s *ListCustomImagesRequest) SetClusterId(v string) *ListCustomImagesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCustomImagesRequest) SetImageOwnerAlias(v string) *ListCustomImagesRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListCustomImagesRequest) SetInstanceType(v string) *ListCustomImagesRequest {
	s.InstanceType = &v
	return s
}

type ListCustomImagesResponseBody struct {
	// The list of custom images and shared images that are supported by the E-HPC.
	Images *ListCustomImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCustomImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBody) SetImages(v *ListCustomImagesResponseBodyImages) *ListCustomImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListCustomImagesResponseBody) SetRequestId(v string) *ListCustomImagesResponseBody {
	s.RequestId = &v
	return s
}

type ListCustomImagesResponseBodyImages struct {
	ImageInfo []*ListCustomImagesResponseBodyImagesImageInfo `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" type:"Repeated"`
}

func (s ListCustomImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyImages) SetImageInfo(v []*ListCustomImagesResponseBodyImagesImageInfo) *ListCustomImagesResponseBodyImages {
	s.ImageInfo = v
	return s
}

type ListCustomImagesResponseBodyImagesImageInfo struct {
	// The image tag of the base operating system.
	BaseOsTag *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty" type:"Struct"`
	// The description of the image.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The type of image. Valid values:
	//
	// *   self: custom image
	// *   others: shared image
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// An array of system images that are supported by E-HPC.
	OsTag *ListCustomImagesResponseBodyImagesImageInfoOsTag `json:"OsTag,omitempty" xml:"OsTag,omitempty" type:"Struct"`
	// The script that is run after the image is installed.
	PostInstallScript *string `json:"PostInstallScript,omitempty" xml:"PostInstallScript,omitempty"`
	// The billing unit of the image. Valid values:
	//
	// *   Hour
	// *   Month
	// *   Year
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The product code on Alibaba Cloud Marketplace.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The size of the image. Unit: GiB
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The stock keeping unit (SKU) of the image. Valid values:
	//
	// \-ECS: pay-as-you-go
	//
	// \-package: subscription
	SkuCode *string `json:"SkuCode,omitempty" xml:"SkuCode,omitempty"`
	// The status of the image. Valid values:
	//
	// *   UnAvailable: The image is unavailable.
	// *   Available: The image is available.
	// *   Creating: The image is being created.
	// *   CreateFailed: The image has failed to be created.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The owner of the image.
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListCustomImagesResponseBodyImagesImageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBodyImagesImageInfo) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetBaseOsTag(v *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) *ListCustomImagesResponseBodyImagesImageInfo {
	s.BaseOsTag = v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetDescription(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.Description = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetImageId(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.ImageId = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetImageName(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.ImageName = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetImageOwnerAlias(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetOsTag(v *ListCustomImagesResponseBodyImagesImageInfoOsTag) *ListCustomImagesResponseBodyImagesImageInfo {
	s.OsTag = v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetPostInstallScript(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.PostInstallScript = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetPricingCycle(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.PricingCycle = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetProductCode(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.ProductCode = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetSize(v int32) *ListCustomImagesResponseBodyImagesImageInfo {
	s.Size = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetSkuCode(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.SkuCode = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetStatus(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.Status = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetUid(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.Uid = &v
	return s
}

type ListCustomImagesResponseBodyImagesImageInfoBaseOsTag struct {
	// The architecture of the operating system. Valid values:
	//
	// *   i386
	// *   x86\_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The image tag of the operating system.
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	// The release version of the operating system. Valid values:
	//
	// *   CentOS
	// *   Ubuntu
	// *   SUSE
	// *   OpenSUSE
	// *   Debian
	// *   CoreOS
	// *   Aliyun
	// *   Windows Server 2003
	// *   Windows Server 2008
	// *   Windows Server 2012
	// *   Others Linux
	// *   Customized Linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The version number of the operating system.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) SetArchitecture(v string) *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag {
	s.Architecture = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) SetOsTag(v string) *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag {
	s.OsTag = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) SetPlatform(v string) *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag {
	s.Platform = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) SetVersion(v string) *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag {
	s.Version = &v
	return s
}

type ListCustomImagesResponseBodyImagesImageInfoOsTag struct {
	// The architecture of the operating system. Valid values:
	//
	// *   i386
	// *   x86\_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The image tag of the base operating system.
	BaseOsTag *string `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	// The image tag of the operating system.
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	// The platform of the operating system.
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The version of the operating system.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListCustomImagesResponseBodyImagesImageInfoOsTag) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBodyImagesImageInfoOsTag) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyImagesImageInfoOsTag) SetArchitecture(v string) *ListCustomImagesResponseBodyImagesImageInfoOsTag {
	s.Architecture = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfoOsTag) SetBaseOsTag(v string) *ListCustomImagesResponseBodyImagesImageInfoOsTag {
	s.BaseOsTag = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfoOsTag) SetOsTag(v string) *ListCustomImagesResponseBodyImagesImageInfoOsTag {
	s.OsTag = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfoOsTag) SetPlatform(v string) *ListCustomImagesResponseBodyImagesImageInfoOsTag {
	s.Platform = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfoOsTag) SetVersion(v string) *ListCustomImagesResponseBodyImagesImageInfoOsTag {
	s.Version = &v
	return s
}

type ListCustomImagesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCustomImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCustomImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponse) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponse) SetHeaders(v map[string]*string) *ListCustomImagesResponse {
	s.Headers = v
	return s
}

func (s *ListCustomImagesResponse) SetStatusCode(v int32) *ListCustomImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomImagesResponse) SetBody(v *ListCustomImagesResponseBody) *ListCustomImagesResponse {
	s.Body = v
	return s
}

type ListFileSystemWithMountTargetsRequest struct {
	// The page number of the page to return.
	//
	// Page numbers start from 1.
	//
	// Default value: 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50.
	//
	// Default value: 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFileSystemWithMountTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsRequest) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsRequest) SetPageNumber(v int32) *ListFileSystemWithMountTargetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFileSystemWithMountTargetsRequest) SetPageSize(v int32) *ListFileSystemWithMountTargetsRequest {
	s.PageSize = &v
	return s
}

type ListFileSystemWithMountTargetsResponseBody struct {
	// The list of file systems.
	FileSystemList *ListFileSystemWithMountTargetsResponseBodyFileSystemList `json:"FileSystemList,omitempty" xml:"FileSystemList,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFileSystemWithMountTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBody) SetFileSystemList(v *ListFileSystemWithMountTargetsResponseBodyFileSystemList) *ListFileSystemWithMountTargetsResponseBody {
	s.FileSystemList = v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBody) SetPageNumber(v int32) *ListFileSystemWithMountTargetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBody) SetPageSize(v int32) *ListFileSystemWithMountTargetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBody) SetRequestId(v string) *ListFileSystemWithMountTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBody) SetTotalCount(v int32) *ListFileSystemWithMountTargetsResponseBody {
	s.TotalCount = &v
	return s
}

type ListFileSystemWithMountTargetsResponseBodyFileSystemList struct {
	FileSystems []*ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Repeated"`
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemList) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemList) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemList) SetFileSystems(v []*ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) *ListFileSystemWithMountTargetsResponseBodyFileSystemList {
	s.FileSystems = v
	return s
}

type ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems struct {
	// The bandwidth of the file system. Unit: MB/s.
	BandWidth *int32 `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	// The capacity of the file system. Unit: GiB.
	Capacity *int32 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The time when the file system was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the file system.
	Destription *string `json:"Destription,omitempty" xml:"Destription,omitempty"`
	// Indicates whether the file system is encrypted. Valid values:
	//
	// *   0: The file system is not encrypted.
	// *   1: The file system is encrypted.
	EncryptType *int32 `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The type of the file system. Valid values:
	//
	// *   standard: General-purpose NAS file system
	// *   extreme: Extreme NAS file system
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The used capacity of the file system. Unit: bytes.
	MeteredSize *int32 `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	// The list of mount targets.
	MountTargetList *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetList `json:"MountTargetList,omitempty" xml:"MountTargetList,omitempty" type:"Struct"`
	// The list of storage plans.
	PackageList *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageList `json:"PackageList,omitempty" xml:"PackageList,omitempty" type:"Struct"`
	// The protocol type of the file system. Valid values:
	//
	// - NFS
	// - SMB
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the file system. Valid values:
	//
	// - Pending: The file system is being created or modified.
	// - Running: The file system is available.
	// - Stopped: The file system is stopped.
	// - Extending: The file system is being scaled out.
	// - Stopping: The file system is being stopped.
	// - Deleting: The file system is being deleted.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage type of the file system.
	//
	// - If FileSystemType is set to standard, the StorageType parameter has the following valid values: Capacity and Performance.
	// - If FileSystemType is set to extreme, the StorageType parameter has the following valid values: standard and advance.
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetBandWidth(v int32) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.BandWidth = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetCapacity(v int32) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.Capacity = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetCreateTime(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.CreateTime = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetDestription(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.Destription = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetEncryptType(v int32) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.EncryptType = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetFileSystemId(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.FileSystemId = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetFileSystemType(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.FileSystemType = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetMeteredSize(v int32) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.MeteredSize = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetMountTargetList(v *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetList) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.MountTargetList = v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetPackageList(v *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageList) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.PackageList = v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetProtocolType(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.ProtocolType = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetRegionId(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.RegionId = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetStatus(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.Status = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetStorageType(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.StorageType = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetVpcId(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.VpcId = &v
	return s
}

type ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetList struct {
	MountTargets []*ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets `json:"MountTargets,omitempty" xml:"MountTargets,omitempty" type:"Repeated"`
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetList) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetList) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetList) SetMountTargets(v []*ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetList {
	s.MountTargets = v
	return s
}

type ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets struct {
	// The name of the permission group that applied to the mount target.
	AccessGroup *string `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	// The domain name of the mount target.
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The network type of the mount target. Valid values:
	//
	// *   Vpc: virtual private cloud (VPC)
	// *   Classic: the classic network
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The status of the mount target. Valid values:
	//
	// *   Active: The mount target is available.
	// *   Inactive: The mount target is inactive.
	// *   Pending: The mount target is being created or modified.
	// *   Deleting: The mount target is being deleted.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch.
	VswId *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetAccessGroup(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.AccessGroup = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetMountTargetDomain(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.MountTargetDomain = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetNetworkType(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.NetworkType = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetStatus(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.Status = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetVpcId(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.VpcId = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetVswId(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.VswId = &v
	return s
}

type ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageList struct {
	Packages []*ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageListPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Repeated"`
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageList) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageList) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageList) SetPackages(v []*ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageListPackages) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageList {
	s.Packages = v
	return s
}

type ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageListPackages struct {
	// The ID of the storage plan.
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageListPackages) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageListPackages) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageListPackages) SetPackageId(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageListPackages {
	s.PackageId = &v
	return s
}

type ListFileSystemWithMountTargetsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFileSystemWithMountTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFileSystemWithMountTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponse) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponse) SetHeaders(v map[string]*string) *ListFileSystemWithMountTargetsResponse {
	s.Headers = v
	return s
}

func (s *ListFileSystemWithMountTargetsResponse) SetStatusCode(v int32) *ListFileSystemWithMountTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponse) SetBody(v *ListFileSystemWithMountTargetsResponseBody) *ListFileSystemWithMountTargetsResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	// The image tag of the operating system. The tag is used only for management nodes.
	BaseOsTag *string `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	// The instance type of the node.
	//
	// *   If a value is passed to the parameter, the list of images that are supported by the specified instance type is queried.
	// *   If no value is passed to the parameter, the list of images that are supported by all instance types is queried.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetBaseOsTag(v string) *ListImagesRequest {
	s.BaseOsTag = &v
	return s
}

func (s *ListImagesRequest) SetInstanceType(v string) *ListImagesRequest {
	s.InstanceType = &v
	return s
}

type ListImagesResponseBody struct {
	// The list of images that are supported by E-HPC.
	OsTags *ListImagesResponseBodyOsTags `json:"OsTags,omitempty" xml:"OsTags,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetOsTags(v *ListImagesResponseBodyOsTags) *ListImagesResponseBody {
	s.OsTags = v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

type ListImagesResponseBodyOsTags struct {
	OsInfo []*ListImagesResponseBodyOsTagsOsInfo `json:"OsInfo,omitempty" xml:"OsInfo,omitempty" type:"Repeated"`
}

func (s ListImagesResponseBodyOsTags) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyOsTags) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyOsTags) SetOsInfo(v []*ListImagesResponseBodyOsTagsOsInfo) *ListImagesResponseBodyOsTags {
	s.OsInfo = v
	return s
}

type ListImagesResponseBodyOsTagsOsInfo struct {
	// The architecture of the operating system. Valid values:
	//
	// *   i386
	// *   x86\_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The image tag of the operating system. The tag is used only for management nodes.
	BaseOsTag *string `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	// The ID of the image.
	ImageId  *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	OSName   *string `json:"OSName,omitempty" xml:"OSName,omitempty"`
	OSNameEn *string `json:"OSNameEn,omitempty" xml:"OSNameEn,omitempty"`
	// The image tag of the cluster.
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	// The operating system. Valid values:
	//
	// *   CentOS
	// *   windows
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The version of the operating system.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListImagesResponseBodyOsTagsOsInfo) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyOsTagsOsInfo) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetArchitecture(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.Architecture = &v
	return s
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetBaseOsTag(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.BaseOsTag = &v
	return s
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetImageId(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetOSName(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.OSName = &v
	return s
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetOSNameEn(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.OSNameEn = &v
	return s
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetOsTag(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.OsTag = &v
	return s
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetPlatform(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.Platform = &v
	return s
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetVersion(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.Version = &v
	return s
}

type ListImagesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetStatusCode(v int32) *ListImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type ListInstalledSoftwareRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListInstalledSoftwareRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwareRequest) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwareRequest) SetClusterId(v string) *ListInstalledSoftwareRequest {
	s.ClusterId = &v
	return s
}

type ListInstalledSoftwareResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of installed software.
	SoftwareList *ListInstalledSoftwareResponseBodySoftwareList `json:"SoftwareList,omitempty" xml:"SoftwareList,omitempty" type:"Struct"`
}

func (s ListInstalledSoftwareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwareResponseBody) SetRequestId(v string) *ListInstalledSoftwareResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstalledSoftwareResponseBody) SetSoftwareList(v *ListInstalledSoftwareResponseBodySoftwareList) *ListInstalledSoftwareResponseBody {
	s.SoftwareList = v
	return s
}

type ListInstalledSoftwareResponseBodySoftwareList struct {
	SoftwareList []*ListInstalledSoftwareResponseBodySoftwareListSoftwareList `json:"SoftwareList,omitempty" xml:"SoftwareList,omitempty" type:"Repeated"`
}

func (s ListInstalledSoftwareResponseBodySoftwareList) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwareResponseBodySoftwareList) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwareResponseBodySoftwareList) SetSoftwareList(v []*ListInstalledSoftwareResponseBodySoftwareListSoftwareList) *ListInstalledSoftwareResponseBodySoftwareList {
	s.SoftwareList = v
	return s
}

type ListInstalledSoftwareResponseBodySoftwareListSoftwareList struct {
	// The ID of the software.
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	// The name of the software.
	SoftwareName *string `json:"SoftwareName,omitempty" xml:"SoftwareName,omitempty"`
	// The status of the software. Valid values:
	//
	// *   Installing: The software is being installed.
	// *   Installed: The software is installed.
	SoftwareStatus *string `json:"SoftwareStatus,omitempty" xml:"SoftwareStatus,omitempty"`
	// The version of the software.
	SoftwareVersion *string `json:"SoftwareVersion,omitempty" xml:"SoftwareVersion,omitempty"`
}

func (s ListInstalledSoftwareResponseBodySoftwareListSoftwareList) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwareResponseBodySoftwareListSoftwareList) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwareResponseBodySoftwareListSoftwareList) SetSoftwareId(v string) *ListInstalledSoftwareResponseBodySoftwareListSoftwareList {
	s.SoftwareId = &v
	return s
}

func (s *ListInstalledSoftwareResponseBodySoftwareListSoftwareList) SetSoftwareName(v string) *ListInstalledSoftwareResponseBodySoftwareListSoftwareList {
	s.SoftwareName = &v
	return s
}

func (s *ListInstalledSoftwareResponseBodySoftwareListSoftwareList) SetSoftwareStatus(v string) *ListInstalledSoftwareResponseBodySoftwareListSoftwareList {
	s.SoftwareStatus = &v
	return s
}

func (s *ListInstalledSoftwareResponseBodySoftwareListSoftwareList) SetSoftwareVersion(v string) *ListInstalledSoftwareResponseBodySoftwareListSoftwareList {
	s.SoftwareVersion = &v
	return s
}

type ListInstalledSoftwareResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstalledSoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstalledSoftwareResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwareResponse) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwareResponse) SetHeaders(v map[string]*string) *ListInstalledSoftwareResponse {
	s.Headers = v
	return s
}

func (s *ListInstalledSoftwareResponse) SetStatusCode(v int32) *ListInstalledSoftwareResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstalledSoftwareResponse) SetBody(v *ListInstalledSoftwareResponseBody) *ListInstalledSoftwareResponse {
	s.Body = v
	return s
}

type ListInvocationResultsRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the command.
	//
	// You can call the [ListCommands](~~87388~~) operation to query the command ID.
	CommandId *string                                 `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	Instance  []*ListInvocationResultsRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
	// The status of the command that you want to query. Valid values:
	//
	// *   Finished
	// *   Running
	// *   Failed
	// *   Stopped
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	// The number of the page to return.
	//
	// Page numbers start from 1.
	//
	// Default value: 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50.
	//
	// Default value: 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInvocationResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationResultsRequest) GoString() string {
	return s.String()
}

func (s *ListInvocationResultsRequest) SetClusterId(v string) *ListInvocationResultsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInvocationResultsRequest) SetCommandId(v string) *ListInvocationResultsRequest {
	s.CommandId = &v
	return s
}

func (s *ListInvocationResultsRequest) SetInstance(v []*ListInvocationResultsRequestInstance) *ListInvocationResultsRequest {
	s.Instance = v
	return s
}

func (s *ListInvocationResultsRequest) SetInvokeRecordStatus(v string) *ListInvocationResultsRequest {
	s.InvokeRecordStatus = &v
	return s
}

func (s *ListInvocationResultsRequest) SetPageNumber(v int32) *ListInvocationResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInvocationResultsRequest) SetPageSize(v int32) *ListInvocationResultsRequest {
	s.PageSize = &v
	return s
}

type ListInvocationResultsRequestInstance struct {
	// The ID of the node on which the command is run.
	//
	// >  The Instance.N.Id parameter specifies the node on which the command is run. If it is not specified, the command is run on all nodes of the cluster.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListInvocationResultsRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationResultsRequestInstance) GoString() string {
	return s.String()
}

func (s *ListInvocationResultsRequestInstance) SetId(v string) *ListInvocationResultsRequestInstance {
	s.Id = &v
	return s
}

type ListInvocationResultsResponseBody struct {
	// The result of the command.
	InvocationResults *ListInvocationResultsResponseBodyInvocationResults `json:"InvocationResults,omitempty" xml:"InvocationResults,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInvocationResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInvocationResultsResponseBody) SetInvocationResults(v *ListInvocationResultsResponseBodyInvocationResults) *ListInvocationResultsResponseBody {
	s.InvocationResults = v
	return s
}

func (s *ListInvocationResultsResponseBody) SetPageNumber(v int32) *ListInvocationResultsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInvocationResultsResponseBody) SetPageSize(v int32) *ListInvocationResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInvocationResultsResponseBody) SetRequestId(v string) *ListInvocationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInvocationResultsResponseBody) SetTotalCount(v int32) *ListInvocationResultsResponseBody {
	s.TotalCount = &v
	return s
}

type ListInvocationResultsResponseBodyInvocationResults struct {
	InvocationResult []*ListInvocationResultsResponseBodyInvocationResultsInvocationResult `json:"InvocationResult,omitempty" xml:"InvocationResult,omitempty" type:"Repeated"`
}

func (s ListInvocationResultsResponseBodyInvocationResults) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationResultsResponseBodyInvocationResults) GoString() string {
	return s.String()
}

func (s *ListInvocationResultsResponseBodyInvocationResults) SetInvocationResult(v []*ListInvocationResultsResponseBodyInvocationResultsInvocationResult) *ListInvocationResultsResponseBodyInvocationResults {
	s.InvocationResult = v
	return s
}

type ListInvocationResultsResponseBodyInvocationResultsInvocationResult struct {
	// The ID of the command.
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The exit code.
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time at which the command entered the Finished state.
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The ID of the node on which the command was run.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the command. Valid values:
	//
	// *   Finished
	// *   Running
	// *   Failed
	// *   Stopped
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	// The output result.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the command was run and its result was obtained.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInvocationResultsResponseBodyInvocationResultsInvocationResult) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationResultsResponseBodyInvocationResultsInvocationResult) GoString() string {
	return s.String()
}

func (s *ListInvocationResultsResponseBodyInvocationResultsInvocationResult) SetCommandId(v string) *ListInvocationResultsResponseBodyInvocationResultsInvocationResult {
	s.CommandId = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResultsInvocationResult) SetExitCode(v int32) *ListInvocationResultsResponseBodyInvocationResultsInvocationResult {
	s.ExitCode = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResultsInvocationResult) SetFinishedTime(v string) *ListInvocationResultsResponseBodyInvocationResultsInvocationResult {
	s.FinishedTime = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResultsInvocationResult) SetInstanceId(v string) *ListInvocationResultsResponseBodyInvocationResultsInvocationResult {
	s.InstanceId = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResultsInvocationResult) SetInvokeRecordStatus(v string) *ListInvocationResultsResponseBodyInvocationResultsInvocationResult {
	s.InvokeRecordStatus = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResultsInvocationResult) SetMessage(v string) *ListInvocationResultsResponseBodyInvocationResultsInvocationResult {
	s.Message = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResultsInvocationResult) SetSuccess(v bool) *ListInvocationResultsResponseBodyInvocationResultsInvocationResult {
	s.Success = &v
	return s
}

type ListInvocationResultsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInvocationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInvocationResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationResultsResponse) GoString() string {
	return s.String()
}

func (s *ListInvocationResultsResponse) SetHeaders(v map[string]*string) *ListInvocationResultsResponse {
	s.Headers = v
	return s
}

func (s *ListInvocationResultsResponse) SetStatusCode(v int32) *ListInvocationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInvocationResultsResponse) SetBody(v *ListInvocationResultsResponseBody) *ListInvocationResultsResponse {
	s.Body = v
	return s
}

type ListInvocationStatusRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the command.
	//
	// You can call the [ListCommands](~~87388~~) operation to query the command ID.
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
}

func (s ListInvocationStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationStatusRequest) GoString() string {
	return s.String()
}

func (s *ListInvocationStatusRequest) SetClusterId(v string) *ListInvocationStatusRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInvocationStatusRequest) SetCommandId(v string) *ListInvocationStatusRequest {
	s.CommandId = &v
	return s
}

type ListInvocationStatusResponseBody struct {
	// The ID of the command.
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The list of statuses. A list is returned for each node.
	InvokeInstances *ListInvocationStatusResponseBodyInvokeInstances `json:"InvokeInstances,omitempty" xml:"InvokeInstances,omitempty" type:"Struct"`
	// The overall status of all nodes in the cluster. Valid values:
	//
	// - Finished
	// - Running
	// - Failed
	// - Stopped
	InvokeStatus *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInvocationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListInvocationStatusResponseBody) SetCommandId(v string) *ListInvocationStatusResponseBody {
	s.CommandId = &v
	return s
}

func (s *ListInvocationStatusResponseBody) SetInvokeInstances(v *ListInvocationStatusResponseBodyInvokeInstances) *ListInvocationStatusResponseBody {
	s.InvokeInstances = v
	return s
}

func (s *ListInvocationStatusResponseBody) SetInvokeStatus(v string) *ListInvocationStatusResponseBody {
	s.InvokeStatus = &v
	return s
}

func (s *ListInvocationStatusResponseBody) SetRequestId(v string) *ListInvocationStatusResponseBody {
	s.RequestId = &v
	return s
}

type ListInvocationStatusResponseBodyInvokeInstances struct {
	InvokeInstance []*ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance `json:"InvokeInstance,omitempty" xml:"InvokeInstance,omitempty" type:"Repeated"`
}

func (s ListInvocationStatusResponseBodyInvokeInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationStatusResponseBodyInvokeInstances) GoString() string {
	return s.String()
}

func (s *ListInvocationStatusResponseBodyInvokeInstances) SetInvokeInstance(v []*ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance) *ListInvocationStatusResponseBodyInvokeInstances {
	s.InvokeInstance = v
	return s
}

type ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance struct {
	// The ID of the node.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the node. Valid values:
	//
	// *   Finished
	// *   Running
	// *   Failed
	// *   Stopped
	InstanceInvokeStatus *string `json:"InstanceInvokeStatus,omitempty" xml:"InstanceInvokeStatus,omitempty"`
}

func (s ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance) GoString() string {
	return s.String()
}

func (s *ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance) SetInstanceId(v string) *ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance {
	s.InstanceId = &v
	return s
}

func (s *ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance) SetInstanceInvokeStatus(v string) *ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance {
	s.InstanceInvokeStatus = &v
	return s
}

type ListInvocationStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInvocationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInvocationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationStatusResponse) GoString() string {
	return s.String()
}

func (s *ListInvocationStatusResponse) SetHeaders(v map[string]*string) *ListInvocationStatusResponse {
	s.Headers = v
	return s
}

func (s *ListInvocationStatusResponse) SetStatusCode(v int32) *ListInvocationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInvocationStatusResponse) SetBody(v *ListInvocationStatusResponseBody) *ListInvocationStatusResponse {
	s.Body = v
	return s
}

type ListJobTemplatesRequest struct {
	// The name of the job template.
	//
	// You can call the [ListJobTemplates](~~87248~~) operation to obtain the job template name.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of the page to return. Page numbers start from 1.
	//
	// Default value: 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 50.
	//
	// Default value: 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListJobTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesRequest) SetName(v string) *ListJobTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListJobTemplatesRequest) SetPageNumber(v int32) *ListJobTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobTemplatesRequest) SetPageSize(v int32) *ListJobTemplatesRequest {
	s.PageSize = &v
	return s
}

type ListJobTemplatesResponseBody struct {
	// The number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of job templates.
	Templates *ListJobTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Struct"`
	// The total number of returned entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponseBody) SetPageNumber(v int32) *ListJobTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobTemplatesResponseBody) SetPageSize(v int32) *ListJobTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobTemplatesResponseBody) SetRequestId(v string) *ListJobTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobTemplatesResponseBody) SetTemplates(v *ListJobTemplatesResponseBodyTemplates) *ListJobTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *ListJobTemplatesResponseBody) SetTotalCount(v int32) *ListJobTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

type ListJobTemplatesResponseBodyTemplates struct {
	JobTemplates []*ListJobTemplatesResponseBodyTemplatesJobTemplates `json:"JobTemplates,omitempty" xml:"JobTemplates,omitempty" type:"Repeated"`
}

func (s ListJobTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponseBodyTemplates) SetJobTemplates(v []*ListJobTemplatesResponseBodyTemplatesJobTemplates) *ListJobTemplatesResponseBodyTemplates {
	s.JobTemplates = v
	return s
}

type ListJobTemplatesResponseBodyTemplatesJobTemplates struct {
	// The job array.
	//
	// Format: X-Y:Z. X is the minimum index value. Y is the maximum index value. Z is the step size. For example, 2-7:2 indicates that three jobs need to be run and their index values are 2, 4, and 6.
	ArrayRequest *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	// The maximum running time of the job. Valid formats:
	//
	// *   hh:mm:ss
	// *   mm:ss
	// *   ss
	ClockTime *string `json:"ClockTime,omitempty" xml:"ClockTime,omitempty"`
	// The command that was used to run the job.
	CommandLine *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	// The maximum GPU usage required by a single compute node. Valid values: 1 to 8.
	//
	// The parameter takes effect only when the cluster uses PBS and a compute node is a GPU-accelerated instance.
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The ID of the job template.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The URL of the job files that were uploaded to an Object Storage Service (OSS) bucket.
	InputFileUrl *string `json:"InputFileUrl,omitempty" xml:"InputFileUrl,omitempty"`
	// The maximum memory usage of a single compute node. The unit can be GB, MB, or KB, and is case-insensitive.
	Mem *string `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The name of the job template.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of the compute nodes. Valid values: 1 to 500.
	Node *int32 `json:"Node,omitempty" xml:"Node,omitempty"`
	// The path that was used to run the job.
	PackagePath *string `json:"PackagePath,omitempty" xml:"PackagePath,omitempty"`
	// The priority of the job. Valid values: 0 to 9. A large value indicates a high priority.
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The queue of the job.
	Queue *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	// Indicates whether the job can be rerun. Valid values:
	//
	// *   true: The job can be rerun.
	// *   false: The job cannot be rerun.
	ReRunable *bool `json:"ReRunable,omitempty" xml:"ReRunable,omitempty"`
	// The name of the user that ran the job.
	RunasUser *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	// The output file path of stderr.
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" xml:"StderrRedirectPath,omitempty"`
	// The output file path of stdout.
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" xml:"StdoutRedirectPath,omitempty"`
	// The number of tasks required by a single compute node. Valid values: 1 to 1000.
	Task *int32 `json:"Task,omitempty" xml:"Task,omitempty"`
	// The number of threads required by a single compute node. Valid values: 1 to 1000.
	Thread *int32 `json:"Thread,omitempty" xml:"Thread,omitempty"`
	// The command that was used to decompress the job files downloaded from an OSS bucket. The parameter takes effect only when WithUnzipCmd is set to true. Valid values:
	//
	// *   tar xzf: decompresses GZIP files.
	// *   tar xf: decompresses TAR files.
	// *   unzip: decompresses ZIP files.
	UnzipCmd *string `json:"UnzipCmd,omitempty" xml:"UnzipCmd,omitempty"`
	// The environment variables of the job.
	Variables *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
	// Indicates whether to decompress the job files downloaded from an OSS bucket. Valid values:
	//
	// *   true: The job files are decompressed.
	// *   false: The job files are not decompressed.
	WithUnzipCmd *bool `json:"WithUnzipCmd,omitempty" xml:"WithUnzipCmd,omitempty"`
}

func (s ListJobTemplatesResponseBodyTemplatesJobTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesResponseBodyTemplatesJobTemplates) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetArrayRequest(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.ArrayRequest = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetClockTime(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.ClockTime = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetCommandLine(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.CommandLine = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetGpu(v int32) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Gpu = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetId(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Id = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetInputFileUrl(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.InputFileUrl = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetMem(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Mem = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetName(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Name = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetNode(v int32) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Node = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetPackagePath(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.PackagePath = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetPriority(v int32) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Priority = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetQueue(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Queue = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetReRunable(v bool) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.ReRunable = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetRunasUser(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.RunasUser = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetStderrRedirectPath(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.StderrRedirectPath = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetStdoutRedirectPath(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.StdoutRedirectPath = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetTask(v int32) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Task = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetThread(v int32) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Thread = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetUnzipCmd(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.UnzipCmd = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetVariables(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Variables = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetWithUnzipCmd(v bool) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.WithUnzipCmd = &v
	return s
}

type ListJobTemplatesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListJobTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponse) SetHeaders(v map[string]*string) *ListJobTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListJobTemplatesResponse) SetStatusCode(v int32) *ListJobTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobTemplatesResponse) SetBody(v *ListJobTemplatesResponseBody) *ListJobTemplatesResponse {
	s.Body = v
	return s
}

type ListJobsRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the user that runs the job.
	//
	// You can call the [ListUsers](~~188572~~) operation to query the users in the cluster.
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 50.
	//
	// Default value: 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether the job can be rerun. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false
	Rerunable *string `json:"Rerunable,omitempty" xml:"Rerunable,omitempty"`
	// The status of the job. Valid values:
	//
	// *   all
	// *   finished
	// *   notfinish
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) SetClusterId(v string) *ListJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListJobsRequest) SetOwner(v string) *ListJobsRequest {
	s.Owner = &v
	return s
}

func (s *ListJobsRequest) SetPageNumber(v int32) *ListJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) SetRerunable(v string) *ListJobsRequest {
	s.Rerunable = &v
	return s
}

func (s *ListJobsRequest) SetState(v string) *ListJobsRequest {
	s.State = &v
	return s
}

type ListJobsResponseBody struct {
	// The list of jobs.
	Jobs *ListJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) SetJobs(v *ListJobsResponseBodyJobs) *ListJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListJobsResponseBody) SetPageNumber(v int32) *ListJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobsResponseBody) SetPageSize(v int32) *ListJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalCount(v int32) *ListJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListJobsResponseBodyJobs struct {
	JobInfo []*ListJobsResponseBodyJobsJobInfo `json:"JobInfo,omitempty" xml:"JobInfo,omitempty" type:"Repeated"`
}

func (s ListJobsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobs) SetJobInfo(v []*ListJobsResponseBodyJobsJobInfo) *ListJobsResponseBodyJobs {
	s.JobInfo = v
	return s
}

type ListJobsResponseBodyJobsJobInfo struct {
	// The job array. If the job is not in a queue, the output is empty.
	//
	// Format: X-Y:Z. X is the minimum index value. Y is the maximum index value. Z is the step size. For example, 2-7:2 indicates that three jobs need to be run and their index values are 2, 4, and 6.
	ArrayRequest *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	// The description of the job.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the job.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the job was last modified.
	LastModifyTime *string `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The name of the job.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of compute nodes that were used to run the job.
	NodeList *string `json:"NodeList,omitempty" xml:"NodeList,omitempty"`
	// The name of the user that runs the job.
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The priority of the job. Valid values: 0 to 9. A large value indicates a high priority.
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The resources that were used to run the job.
	Resources *ListJobsResponseBodyJobsJobInfoResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// The path that was used to run the job.
	ShellPath *string `json:"ShellPath,omitempty" xml:"ShellPath,omitempty"`
	// The time when the job started to run.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the job. Valid values: Valid values:
	//
	// *   FINISHED: The job is completed
	// *   RUNNING: The job is running.
	// *   QUEUED: The job is pending in a queue.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The output file path of stderr.
	Stderr *string `json:"Stderr,omitempty" xml:"Stderr,omitempty"`
	// The output file path of stdout.
	Stdout *string `json:"Stdout,omitempty" xml:"Stdout,omitempty"`
	// The time when the job was submitted.
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s ListJobsResponseBodyJobsJobInfo) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobsJobInfo) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsJobInfo) SetArrayRequest(v string) *ListJobsResponseBodyJobsJobInfo {
	s.ArrayRequest = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetComment(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Comment = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetId(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Id = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetLastModifyTime(v string) *ListJobsResponseBodyJobsJobInfo {
	s.LastModifyTime = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetName(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Name = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetNodeList(v string) *ListJobsResponseBodyJobsJobInfo {
	s.NodeList = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetOwner(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Owner = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetPriority(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Priority = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetResources(v *ListJobsResponseBodyJobsJobInfoResources) *ListJobsResponseBodyJobsJobInfo {
	s.Resources = v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetShellPath(v string) *ListJobsResponseBodyJobsJobInfo {
	s.ShellPath = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetStartTime(v string) *ListJobsResponseBodyJobsJobInfo {
	s.StartTime = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetState(v string) *ListJobsResponseBodyJobsJobInfo {
	s.State = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetStderr(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Stderr = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetStdout(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Stdout = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetSubmitTime(v string) *ListJobsResponseBodyJobsJobInfo {
	s.SubmitTime = &v
	return s
}

type ListJobsResponseBodyJobsJobInfoResources struct {
	// The number of CPUs that were used to run the job.
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The number of nodes that were used to run the job.
	Nodes *int32 `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s ListJobsResponseBodyJobsJobInfoResources) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobsJobInfoResources) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsJobInfoResources) SetCores(v int32) *ListJobsResponseBodyJobsJobInfoResources {
	s.Cores = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfoResources) SetNodes(v int32) *ListJobsResponseBodyJobsJobInfoResources {
	s.Nodes = &v
	return s
}

type ListJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponse) GoString() string {
	return s.String()
}

func (s *ListJobsResponse) SetHeaders(v map[string]*string) *ListJobsResponse {
	s.Headers = v
	return s
}

func (s *ListJobsResponse) SetStatusCode(v int32) *ListJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobsResponse) SetBody(v *ListJobsResponseBody) *ListJobsResponse {
	s.Body = v
	return s
}

type ListJobsWithFiltersRequest struct {
	// Specifies whether to enable asynchronous query.
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The latest time when a job is submitted. The value is a UNIX timestamp, which represents the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// The earliest time when a job is submitted. The value is a UNIX timestamp, which represents the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// The order in which jobs are sorted based on the execution time. Valid values:
	//
	// *   asc: ascending order
	// *   desc: descending order
	ExecuteOrder *string `json:"ExecuteOrder,omitempty" xml:"ExecuteOrder,omitempty"`
	// The name of the job. Fuzzy match is supported.
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The status of the job. Valid values:
	//
	// *   all
	// *   finished
	// *   notfinish
	//
	// Default value: all
	JobStatus *string   `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	Nodes     []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 50.
	//
	// Default value: 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The order in which jobs are sorted based on the time when they queue. Valid values:
	//
	// *   asc: ascending order
	// *   desc: descending order
	PendOrder *string   `json:"PendOrder,omitempty" xml:"PendOrder,omitempty"`
	Queues    []*string `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
	// The ID of the region.
	//
	// You can call the [ListRegions](~~188593~~) operation to query the list of regions where E-HPC is supported.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The order in which jobs are sorted based on the time when they are submitted. Valid values:
	//
	// *   asc: ascending order
	// *   desc: descending order
	SubmitOrder *string   `json:"SubmitOrder,omitempty" xml:"SubmitOrder,omitempty"`
	Users       []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListJobsWithFiltersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsWithFiltersRequest) GoString() string {
	return s.String()
}

func (s *ListJobsWithFiltersRequest) SetAsync(v bool) *ListJobsWithFiltersRequest {
	s.Async = &v
	return s
}

func (s *ListJobsWithFiltersRequest) SetClusterId(v string) *ListJobsWithFiltersRequest {
	s.ClusterId = &v
	return s
}

func (s *ListJobsWithFiltersRequest) SetCreateTimeEnd(v string) *ListJobsWithFiltersRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListJobsWithFiltersRequest) SetCreateTimeStart(v string) *ListJobsWithFiltersRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListJobsWithFiltersRequest) SetExecuteOrder(v string) *ListJobsWithFiltersRequest {
	s.ExecuteOrder = &v
	return s
}

func (s *ListJobsWithFiltersRequest) SetJobName(v string) *ListJobsWithFiltersRequest {
	s.JobName = &v
	return s
}

func (s *ListJobsWithFiltersRequest) SetJobStatus(v string) *ListJobsWithFiltersRequest {
	s.JobStatus = &v
	return s
}

func (s *ListJobsWithFiltersRequest) SetNodes(v []*string) *ListJobsWithFiltersRequest {
	s.Nodes = v
	return s
}

func (s *ListJobsWithFiltersRequest) SetPageNumber(v int64) *ListJobsWithFiltersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsWithFiltersRequest) SetPageSize(v int64) *ListJobsWithFiltersRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsWithFiltersRequest) SetPendOrder(v string) *ListJobsWithFiltersRequest {
	s.PendOrder = &v
	return s
}

func (s *ListJobsWithFiltersRequest) SetQueues(v []*string) *ListJobsWithFiltersRequest {
	s.Queues = v
	return s
}

func (s *ListJobsWithFiltersRequest) SetRegionId(v string) *ListJobsWithFiltersRequest {
	s.RegionId = &v
	return s
}

func (s *ListJobsWithFiltersRequest) SetSubmitOrder(v string) *ListJobsWithFiltersRequest {
	s.SubmitOrder = &v
	return s
}

func (s *ListJobsWithFiltersRequest) SetUsers(v []*string) *ListJobsWithFiltersRequest {
	s.Users = v
	return s
}

type ListJobsWithFiltersResponseBody struct {
	// The list of jobs.
	Jobs []*ListJobsWithFiltersResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   true: The call was successful.
	// *   false: The request failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobsWithFiltersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobsWithFiltersResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsWithFiltersResponseBody) SetJobs(v []*ListJobsWithFiltersResponseBodyJobs) *ListJobsWithFiltersResponseBody {
	s.Jobs = v
	return s
}

func (s *ListJobsWithFiltersResponseBody) SetPageNumber(v int64) *ListJobsWithFiltersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobsWithFiltersResponseBody) SetPageSize(v int64) *ListJobsWithFiltersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobsWithFiltersResponseBody) SetRequestId(v string) *ListJobsWithFiltersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsWithFiltersResponseBody) SetSuccess(v bool) *ListJobsWithFiltersResponseBody {
	s.Success = &v
	return s
}

func (s *ListJobsWithFiltersResponseBody) SetTotalCount(v int32) *ListJobsWithFiltersResponseBody {
	s.TotalCount = &v
	return s
}

type ListJobsWithFiltersResponseBodyJobs struct {
	// The job array. If the job is not in a queue, the output is empty.
	//
	// Format: X-Y:Z. X is the minimum index value. Y is the maximum index value. Z is the step size. For example, 2-7:2 indicates that three jobs need to be run and their index values are 2, 4, and 6.
	ArrayRequest *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	// The description of the job.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the job.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the job was last modified.
	LastModifyTime *string `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The name of the job.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of compute nodes that were used to run the job.
	NodeList *string `json:"NodeList,omitempty" xml:"NodeList,omitempty"`
	// The name of the user that ran the job.
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The priority of the job. Valid values: 0 to 9. A large value indicates a high priority.
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The number of queues that ran the job.
	Queue *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	// Indicates whether the job can be run again. Valid values:
	//
	// *   true
	// *   false
	Rerunable *bool `json:"Rerunable,omitempty" xml:"Rerunable,omitempty"`
	// The resources that were used to run the job.
	Resources *ListJobsWithFiltersResponseBodyJobsResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// The path that was used to run the job.
	ShellPath *string `json:"ShellPath,omitempty" xml:"ShellPath,omitempty"`
	// The time when the job started to run.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the job. Valid values:
	//
	// *   FINISHED: The job is completed.
	// *   RUNNING: The job connector is running.
	// *   QUEUED: The job is pending in a queue.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The output file path of stderr.
	Stderr *string `json:"Stderr,omitempty" xml:"Stderr,omitempty"`
	// The output file path of stdout.
	Stdout *string `json:"Stdout,omitempty" xml:"Stdout,omitempty"`
	// The time when the job was submitted.
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// The list of variables of the job.
	VariableList *string `json:"VariableList,omitempty" xml:"VariableList,omitempty"`
}

func (s ListJobsWithFiltersResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListJobsWithFiltersResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetArrayRequest(v string) *ListJobsWithFiltersResponseBodyJobs {
	s.ArrayRequest = &v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetComment(v string) *ListJobsWithFiltersResponseBodyJobs {
	s.Comment = &v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetId(v string) *ListJobsWithFiltersResponseBodyJobs {
	s.Id = &v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetLastModifyTime(v string) *ListJobsWithFiltersResponseBodyJobs {
	s.LastModifyTime = &v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetName(v string) *ListJobsWithFiltersResponseBodyJobs {
	s.Name = &v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetNodeList(v string) *ListJobsWithFiltersResponseBodyJobs {
	s.NodeList = &v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetOwner(v string) *ListJobsWithFiltersResponseBodyJobs {
	s.Owner = &v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetPriority(v string) *ListJobsWithFiltersResponseBodyJobs {
	s.Priority = &v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetQueue(v string) *ListJobsWithFiltersResponseBodyJobs {
	s.Queue = &v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetRerunable(v bool) *ListJobsWithFiltersResponseBodyJobs {
	s.Rerunable = &v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetResources(v *ListJobsWithFiltersResponseBodyJobsResources) *ListJobsWithFiltersResponseBodyJobs {
	s.Resources = v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetShellPath(v string) *ListJobsWithFiltersResponseBodyJobs {
	s.ShellPath = &v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetStartTime(v string) *ListJobsWithFiltersResponseBodyJobs {
	s.StartTime = &v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetState(v string) *ListJobsWithFiltersResponseBodyJobs {
	s.State = &v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetStderr(v string) *ListJobsWithFiltersResponseBodyJobs {
	s.Stderr = &v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetStdout(v string) *ListJobsWithFiltersResponseBodyJobs {
	s.Stdout = &v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetSubmitTime(v string) *ListJobsWithFiltersResponseBodyJobs {
	s.SubmitTime = &v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobs) SetVariableList(v string) *ListJobsWithFiltersResponseBodyJobs {
	s.VariableList = &v
	return s
}

type ListJobsWithFiltersResponseBodyJobsResources struct {
	// The number of CPUs that were used to run the job.
	Cores *int64 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The number of nodes that were used to run the job.
	Nodes *int64 `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s ListJobsWithFiltersResponseBodyJobsResources) String() string {
	return tea.Prettify(s)
}

func (s ListJobsWithFiltersResponseBodyJobsResources) GoString() string {
	return s.String()
}

func (s *ListJobsWithFiltersResponseBodyJobsResources) SetCores(v int64) *ListJobsWithFiltersResponseBodyJobsResources {
	s.Cores = &v
	return s
}

func (s *ListJobsWithFiltersResponseBodyJobsResources) SetNodes(v int64) *ListJobsWithFiltersResponseBodyJobsResources {
	s.Nodes = &v
	return s
}

type ListJobsWithFiltersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListJobsWithFiltersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobsWithFiltersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobsWithFiltersResponse) GoString() string {
	return s.String()
}

func (s *ListJobsWithFiltersResponse) SetHeaders(v map[string]*string) *ListJobsWithFiltersResponse {
	s.Headers = v
	return s
}

func (s *ListJobsWithFiltersResponse) SetStatusCode(v int32) *ListJobsWithFiltersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobsWithFiltersResponse) SetBody(v *ListJobsWithFiltersResponseBody) *ListJobsWithFiltersResponse {
	s.Body = v
	return s
}

type ListNodesRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The filter options of the node list.
	//
	// Format: {"status":"node_status"}. Replace node_status with the node status. Valid values of node_status:
	//
	// *   uninit: The node is being installed.
	// *   exception: An exception has occurred on the node.
	// *   running: The node is running.
	// *   initing: The node is being initialized.
	// *   releasing: The node is being released.
	// *   untracking: The node is not added to the cluster.
	// *   stopped: The node is stopped.
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The name of the node. You can perform a fuzzy search. MySQL regular expressions are supported.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The prefix of the hostname. You can query nodes that have a specified prefix.
	HostNamePrefix *string `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	// The suffix of the hostname. You can query nodes that have a specified suffix.
	HostNameSuffix *string `json:"HostNameSuffix,omitempty" xml:"HostNameSuffix,omitempty"`
	// The number of the page to return. Pages start from page 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100.
	//
	// Default value: 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The private IP address of the node.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The type of the node. Valid values:
	//
	// *   Manager: management node
	// *   Login: logon node
	// *   Compute: compute node
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The sorting method of the node list. Valid values:
	//
	// *   Forward: sorts the nodes in chronological order.
	// *   Backward: sorts the nodes in reverse chronological order.
	//
	// Default value: Forward
	//
	// >  Sequence is used in combination with SortBy. If SortBy is set to AddedTime and Sequence is set to Forward, nodes are sorted by the time that they were added in chronological order.
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The sorting method of the node list. Valid values:
	//
	// *   AddedTime: sorts the nodes by the time that they were added.
	// *   HostName: sorts the nodes by their host names.
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) SetClusterId(v string) *ListNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodesRequest) SetFilter(v string) *ListNodesRequest {
	s.Filter = &v
	return s
}

func (s *ListNodesRequest) SetHostName(v string) *ListNodesRequest {
	s.HostName = &v
	return s
}

func (s *ListNodesRequest) SetHostNamePrefix(v string) *ListNodesRequest {
	s.HostNamePrefix = &v
	return s
}

func (s *ListNodesRequest) SetHostNameSuffix(v string) *ListNodesRequest {
	s.HostNameSuffix = &v
	return s
}

func (s *ListNodesRequest) SetPageNumber(v int32) *ListNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesRequest) SetPageSize(v int32) *ListNodesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesRequest) SetPrivateIpAddress(v string) *ListNodesRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *ListNodesRequest) SetRole(v string) *ListNodesRequest {
	s.Role = &v
	return s
}

func (s *ListNodesRequest) SetSequence(v string) *ListNodesRequest {
	s.Sequence = &v
	return s
}

func (s *ListNodesRequest) SetSortBy(v string) *ListNodesRequest {
	s.SortBy = &v
	return s
}

type ListNodesResponseBody struct {
	// The information about nodes.
	Nodes *ListNodesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) SetNodes(v *ListNodesResponseBodyNodes) *ListNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *ListNodesResponseBody) SetPageNumber(v int32) *ListNodesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNodesResponseBody) SetPageSize(v int32) *ListNodesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesResponseBody) SetTotalCount(v int32) *ListNodesResponseBody {
	s.TotalCount = &v
	return s
}

type ListNodesResponseBodyNodes struct {
	NodeInfo []*ListNodesResponseBodyNodesNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodes) SetNodeInfo(v []*ListNodesResponseBodyNodesNodeInfo) *ListNodesResponseBodyNodes {
	s.NodeInfo = v
	return s
}

type ListNodesResponseBodyNodesNodeInfo struct {
	// The time when the node was added to the cluster.
	AddTime *string `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	// The mode in which the compute nodes are added. Valid values:
	//
	// *   manual: The compute nodes are manually added.
	// *   autoscale: The compute nodes are automatically added.
	CreateMode *string `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	// Indicates whether the node was created by using E-HPC.
	//
	// *   true: The node is created by using E-HPC.
	// *   false: The node is not created by using E-HPC.
	CreatedByEhpc *bool `json:"CreatedByEhpc,omitempty" xml:"CreatedByEhpc,omitempty"`
	// Indicates whether the subscription node expired. For a pay-as-you-go node, false is returned.
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The time when the subscription node expires. For a pay-as-you-go node, a null value is returned.
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The name of the node.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// Indicates whether hyper-threading is enabled.
	HtEnabled *bool `json:"HtEnabled,omitempty" xml:"HtEnabled,omitempty"`
	// The ID of the node.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The type of the image. Valid values:
	//
	// *   system: public image
	// *   self: custom image
	// *   others: shared image
	// *   marketplace: Alibaba Cloud Marketplace image
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The instance types of the node.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The IP address of the node.
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The location where the node was deployed. Valid values:
	//
	// *   OnPremise: The node is deployed on your data center.
	// *   PublicCloud: The node is deployed on the public cloud.
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The reason why the node was locked. Valid values:
	//
	// *   financial: The node is locked due to overdue payments.
	// *   security: The node is locked for security reasons.
	// *   recycling: The preemptible node is locked and pending release.
	// *   dedicatedhostfinancial: The node is locked due to the overdue payments of the dedicated host.
	//
	// By default, an empty string is returned.
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The public IP address of the node.
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the node. Valid values:
	//
	// *   Scheduler: primary scheduling node
	// *   SchedulerBackup: secondary scheduling node
	// *   Account: primary domain server node
	// *   AccountBackup: secondary domain server node
	// *   Login: logon node
	// *   Compute: compute node
	//
	// Scheduling nodes and domain server nodes are management nodes.
	Roles *ListNodesResponseBodyNodesNodeInfoRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
	// The bidding method of the compute nodes.
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The status of the node displayed on the scheduler. The status varies with the scheduler.
	StateInSched *string `json:"StateInSched,omitempty" xml:"StateInSched,omitempty"`
	// The status of the node. Valid values:
	//
	// *   uninit: The node is being installed.
	// *   exception: An exception has occurred on the node.
	// *   running: The node is running.
	// *   initing: The node is being initialized.
	// *   releasing: The node is being released.
	// *   untracking: The node is not added to the cluster.
	// *   stopped: The node is stopped.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The statistics of the resources used by the node.
	TotalResources *ListNodesResponseBodyNodesNodeInfoTotalResources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty" type:"Struct"`
	// The usage of the compute nodes in the cluster. For other types of nodes, an empty value is returned.
	UsedResources *ListNodesResponseBodyNodesNodeInfoUsedResources `json:"UsedResources,omitempty" xml:"UsedResources,omitempty" type:"Struct"`
	// The ID of the vSwitch.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The version of the client.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The ID of the virtual private cloud (VPC).
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListNodesResponseBodyNodesNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodesNodeInfo) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetAddTime(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.AddTime = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetCreateMode(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.CreateMode = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetCreatedByEhpc(v bool) *ListNodesResponseBodyNodesNodeInfo {
	s.CreatedByEhpc = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetExpired(v bool) *ListNodesResponseBodyNodesNodeInfo {
	s.Expired = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetExpiredTime(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.ExpiredTime = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetHostName(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.HostName = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetHtEnabled(v bool) *ListNodesResponseBodyNodesNodeInfo {
	s.HtEnabled = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetId(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetImageId(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.ImageId = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetImageOwnerAlias(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetInstanceType(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.InstanceType = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetIpAddress(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.IpAddress = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetLocation(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.Location = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetLockReason(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.LockReason = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetPublicIpAddress(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.PublicIpAddress = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetRegionId(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.RegionId = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetRoles(v *ListNodesResponseBodyNodesNodeInfoRoles) *ListNodesResponseBodyNodesNodeInfo {
	s.Roles = v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetSpotStrategy(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.SpotStrategy = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetStateInSched(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.StateInSched = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetStatus(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.Status = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetTotalResources(v *ListNodesResponseBodyNodesNodeInfoTotalResources) *ListNodesResponseBodyNodesNodeInfo {
	s.TotalResources = v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetUsedResources(v *ListNodesResponseBodyNodesNodeInfoUsedResources) *ListNodesResponseBodyNodesNodeInfo {
	s.UsedResources = v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetVSwitchId(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.VSwitchId = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetVersion(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.Version = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetVpcId(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.VpcId = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetZoneId(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.ZoneId = &v
	return s
}

type ListNodesResponseBodyNodesNodeInfoRoles struct {
	Role []*string `json:"Role,omitempty" xml:"Role,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBodyNodesNodeInfoRoles) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodesNodeInfoRoles) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodesNodeInfoRoles) SetRole(v []*string) *ListNodesResponseBodyNodesNodeInfoRoles {
	s.Role = v
	return s
}

type ListNodesResponseBodyNodesNodeInfoTotalResources struct {
	// The number of vCPUs.
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of GPUs.
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The memory capacity. Unit: GB
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesResponseBodyNodesNodeInfoTotalResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodesNodeInfoTotalResources) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodesNodeInfoTotalResources) SetCpu(v int32) *ListNodesResponseBodyNodesNodeInfoTotalResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfoTotalResources) SetGpu(v int32) *ListNodesResponseBodyNodesNodeInfoTotalResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfoTotalResources) SetMemory(v int32) *ListNodesResponseBodyNodesNodeInfoTotalResources {
	s.Memory = &v
	return s
}

type ListNodesResponseBodyNodesNodeInfoUsedResources struct {
	// The number of vCPUs.
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of GPUs.
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The memory capacity. Unit: GB
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesResponseBodyNodesNodeInfoUsedResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodesNodeInfoUsedResources) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodesNodeInfoUsedResources) SetCpu(v int32) *ListNodesResponseBodyNodesNodeInfoUsedResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfoUsedResources) SetGpu(v int32) *ListNodesResponseBodyNodesNodeInfoUsedResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfoUsedResources) SetMemory(v int32) *ListNodesResponseBodyNodesNodeInfoUsedResources {
	s.Memory = &v
	return s
}

type ListNodesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponse) GoString() string {
	return s.String()
}

func (s *ListNodesResponse) SetHeaders(v map[string]*string) *ListNodesResponse {
	s.Headers = v
	return s
}

func (s *ListNodesResponse) SetStatusCode(v int32) *ListNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesResponse) SetBody(v *ListNodesResponseBody) *ListNodesResponse {
	s.Body = v
	return s
}

type ListNodesByQueueRequest struct {
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of the page to return. Pages start from page 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50.
	//
	// Default value: 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the queue.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s ListNodesByQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueRequest) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueRequest) SetAsync(v bool) *ListNodesByQueueRequest {
	s.Async = &v
	return s
}

func (s *ListNodesByQueueRequest) SetClusterId(v string) *ListNodesByQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodesByQueueRequest) SetPageNumber(v int32) *ListNodesByQueueRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesByQueueRequest) SetPageSize(v int32) *ListNodesByQueueRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesByQueueRequest) SetQueueName(v string) *ListNodesByQueueRequest {
	s.QueueName = &v
	return s
}

type ListNodesByQueueResponseBody struct {
	// The list of nodes.
	Nodes *ListNodesByQueueResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
	// The number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesByQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueResponseBody) SetNodes(v *ListNodesByQueueResponseBodyNodes) *ListNodesByQueueResponseBody {
	s.Nodes = v
	return s
}

func (s *ListNodesByQueueResponseBody) SetPageNumber(v int32) *ListNodesByQueueResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNodesByQueueResponseBody) SetPageSize(v int32) *ListNodesByQueueResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNodesByQueueResponseBody) SetRequestId(v string) *ListNodesByQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesByQueueResponseBody) SetTotalCount(v int32) *ListNodesByQueueResponseBody {
	s.TotalCount = &v
	return s
}

type ListNodesByQueueResponseBodyNodes struct {
	NodeInfo []*ListNodesByQueueResponseBodyNodesNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
}

func (s ListNodesByQueueResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueResponseBodyNodes) SetNodeInfo(v []*ListNodesByQueueResponseBodyNodesNodeInfo) *ListNodesByQueueResponseBodyNodes {
	s.NodeInfo = v
	return s
}

type ListNodesByQueueResponseBodyNodesNodeInfo struct {
	// The time when the node was added to the cluster.
	AddTime *string `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	// The mode in which the node was added. Valid values:
	//
	// *   manual: The node was manually added.
	// *   autoscale: The node is automatically added.
	CreateMode *string `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	// Indicates whether the node was created by using E-HPC.
	CreatedByEhpc *bool `json:"CreatedByEhpc,omitempty" xml:"CreatedByEhpc,omitempty"`
	// Indicates whether the subscription node has expired. If the node is a pay-as-you-go node, false is returned.
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The time when the subscription instance expires. If the node is a pay-as-you-go node, a null value is returned.
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The name of the node.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// Indicates whether Hyper-Threading (HT) is enabled.
	HtEnabled *bool `json:"HtEnabled,omitempty" xml:"HtEnabled,omitempty"`
	// The ID of the ECS instance.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The type of the image. Valid values:
	//
	// *   system: public image
	// *   self: custom image
	// *   others: shared image
	// *   marketplace: Alibaba Cloud Marketplace image
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The private IP address of the node.
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The location where the node is deployed. Valid values:
	//
	// *   OnPremise: The node is deployed on a hybrid cloud.
	// *   PublicCloud: The node is deployed on a public cloud.
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The reason why the node is locked. Valid values:
	//
	// *   financial: The node is locked due to overdue payments.
	// *   security: The node is locked due to security reasons.
	// *   recycling: The preemptible node is locked and pending release.
	// *   dedicatedhostfinancial: The node is locked due to the overdue payments of the dedicated host.
	//
	// By default, an empty string is returned.
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The public IP address of the node.
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The preemption policy for the Elastic Compute Service (ECS) instance. Valid values:
	//
	// *   NoSpot: applies to regular pay-as-you-go instances.
	// *   SpotWithPriceLimit: The instances of the compute node are preemptible instances. These types of instances have a specified maximum hourly price.
	// *   SpotAsPriceGo: The instances of the compute node are preemptible instances. The price of these instances is based on the current market price.
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The status of the node displayed on the scheduler. The status varies with the scheduler.
	StateInSched *string `json:"StateInSched,omitempty" xml:"StateInSched,omitempty"`
	// The status of the node. Valid values:
	//
	// *   uninit: The node is not initialized.
	// *   init: The node is being initialized.
	// *   ready: The node is ready.
	// *   running: The node is running.
	// *   exception: An exception has occurred on the node.
	// *   untracking: The node is not added to the cluster.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of all resources in the cluster.
	TotalResources *ListNodesByQueueResponseBodyNodesNodeInfoTotalResources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty" type:"Struct"`
	// The usage of the compute nodes in the cluster. For other types of nodes, an empty value is returned.
	UsedResources *ListNodesByQueueResponseBodyNodesNodeInfoUsedResources `json:"UsedResources,omitempty" xml:"UsedResources,omitempty" type:"Struct"`
	// The ID of the vSwitch.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The version of the E-HPC client.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The ID of the virtual private cloud (VPC).
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListNodesByQueueResponseBodyNodesNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueResponseBodyNodesNodeInfo) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetAddTime(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.AddTime = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetCreateMode(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.CreateMode = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetCreatedByEhpc(v bool) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.CreatedByEhpc = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetExpired(v bool) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.Expired = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetExpiredTime(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.ExpiredTime = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetHostName(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.HostName = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetHtEnabled(v bool) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.HtEnabled = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetId(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.Id = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetImageId(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.ImageId = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetImageOwnerAlias(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetIpAddress(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.IpAddress = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetLocation(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.Location = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetLockReason(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.LockReason = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetPublicIpAddress(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.PublicIpAddress = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetRegionId(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.RegionId = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetSpotStrategy(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.SpotStrategy = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetStateInSched(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.StateInSched = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetStatus(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.Status = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetTotalResources(v *ListNodesByQueueResponseBodyNodesNodeInfoTotalResources) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.TotalResources = v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetUsedResources(v *ListNodesByQueueResponseBodyNodesNodeInfoUsedResources) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.UsedResources = v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetVSwitchId(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.VSwitchId = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetVersion(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.Version = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetVpcId(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.VpcId = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetZoneId(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.ZoneId = &v
	return s
}

type ListNodesByQueueResponseBodyNodesNodeInfoTotalResources struct {
	// The number of CPU cores. Unit: cores.
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The total number of GPU cards. Unit: cards.
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The memory capacity. Unit: GB.
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesByQueueResponseBodyNodesNodeInfoTotalResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueResponseBodyNodesNodeInfoTotalResources) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfoTotalResources) SetCpu(v int32) *ListNodesByQueueResponseBodyNodesNodeInfoTotalResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfoTotalResources) SetGpu(v int32) *ListNodesByQueueResponseBodyNodesNodeInfoTotalResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfoTotalResources) SetMemory(v int32) *ListNodesByQueueResponseBodyNodesNodeInfoTotalResources {
	s.Memory = &v
	return s
}

type ListNodesByQueueResponseBodyNodesNodeInfoUsedResources struct {
	// The number of CPU cores. Unit: cores.
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The total number of GPU cards. Unit: cards.
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The memory capacity. Unit: GB.
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesByQueueResponseBodyNodesNodeInfoUsedResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueResponseBodyNodesNodeInfoUsedResources) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfoUsedResources) SetCpu(v int32) *ListNodesByQueueResponseBodyNodesNodeInfoUsedResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfoUsedResources) SetGpu(v int32) *ListNodesByQueueResponseBodyNodesNodeInfoUsedResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfoUsedResources) SetMemory(v int32) *ListNodesByQueueResponseBodyNodesNodeInfoUsedResources {
	s.Memory = &v
	return s
}

type ListNodesByQueueResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodesByQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodesByQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueResponse) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueResponse) SetHeaders(v map[string]*string) *ListNodesByQueueResponse {
	s.Headers = v
	return s
}

func (s *ListNodesByQueueResponse) SetStatusCode(v int32) *ListNodesByQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesByQueueResponse) SetBody(v *ListNodesByQueueResponseBody) *ListNodesByQueueResponse {
	s.Body = v
	return s
}

type ListNodesNoPagingRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the node. You can perform a fuzzy search. MySQL regular expressions are supported.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The type of the node. Valid values:
	//
	// *   Manager: management node
	// *   Login: logon node
	// *   Compute: compute node
	//
	// Default value: Compute
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The sorting method of the node list. Valid values:
	//
	// *   Forward: sorts the nodes in chronological order.
	// *   Backward: sorts the nodes in reverse chronological order.
	//
	// Default value: Forward
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListNodesNoPagingRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingRequest) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingRequest) SetClusterId(v string) *ListNodesNoPagingRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodesNoPagingRequest) SetHostName(v string) *ListNodesNoPagingRequest {
	s.HostName = &v
	return s
}

func (s *ListNodesNoPagingRequest) SetRole(v string) *ListNodesNoPagingRequest {
	s.Role = &v
	return s
}

func (s *ListNodesNoPagingRequest) SetSequence(v string) *ListNodesNoPagingRequest {
	s.Sequence = &v
	return s
}

type ListNodesNoPagingResponseBody struct {
	// The information about nodes.
	Nodes *ListNodesNoPagingResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodesNoPagingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponseBody) SetNodes(v *ListNodesNoPagingResponseBodyNodes) *ListNodesNoPagingResponseBody {
	s.Nodes = v
	return s
}

func (s *ListNodesNoPagingResponseBody) SetRequestId(v string) *ListNodesNoPagingResponseBody {
	s.RequestId = &v
	return s
}

type ListNodesNoPagingResponseBodyNodes struct {
	NodeInfo []*ListNodesNoPagingResponseBodyNodesNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
}

func (s ListNodesNoPagingResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponseBodyNodes) SetNodeInfo(v []*ListNodesNoPagingResponseBodyNodesNodeInfo) *ListNodesNoPagingResponseBodyNodes {
	s.NodeInfo = v
	return s
}

type ListNodesNoPagingResponseBodyNodesNodeInfo struct {
	// The name of the node.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the node.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance type of the node.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The status of the node. Valid values:
	//
	// *   uninit: The node is being installed.
	// *   exception: An exception has occurred on the node.
	// *   running: The node is running.
	// *   initing: The node is being initialized.
	// *   releasing: The node is being released.
	// *   untracking: The node is not added to the cluster.
	// *   stopped: The node is stopped.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListNodesNoPagingResponseBodyNodesNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponseBodyNodesNodeInfo) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetHostName(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.HostName = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetId(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.Id = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetImageId(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.ImageId = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetInstanceType(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.InstanceType = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetStatus(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.Status = &v
	return s
}

type ListNodesNoPagingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodesNoPagingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodesNoPagingResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponse) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponse) SetHeaders(v map[string]*string) *ListNodesNoPagingResponse {
	s.Headers = v
	return s
}

func (s *ListNodesNoPagingResponse) SetStatusCode(v int32) *ListNodesNoPagingResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesNoPagingResponse) SetBody(v *ListNodesNoPagingResponseBody) *ListNodesNoPagingResponse {
	s.Body = v
	return s
}

type ListPreferredEcsTypesRequest struct {
	// The billing method of the ECS instance. Valid values:
	//
	// *   PostPaid: pay-as-you-go
	// *   PrePaid: subscription
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The bidding policy of the ECS instance. Valid values:
	//
	// *   NoSpot: The instance is created as a regular pay-as-you-go instance.
	// *   SpotWithPriceLimit: The instance to be created is a preemptible instance with a user-defined maximum hourly price.
	// *   SpotAsPriceGo: The instance is a preemptible instance whose price is based on the current market price.
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The ID of the zone.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListPreferredEcsTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesRequest) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesRequest) SetInstanceChargeType(v string) *ListPreferredEcsTypesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *ListPreferredEcsTypesRequest) SetSpotStrategy(v string) *ListPreferredEcsTypesRequest {
	s.SpotStrategy = &v
	return s
}

func (s *ListPreferredEcsTypesRequest) SetZoneId(v string) *ListPreferredEcsTypesRequest {
	s.ZoneId = &v
	return s
}

type ListPreferredEcsTypesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of recommended ECS instances. Each SeriesInfo element contains the recommended ECS instance types for various nodes of the E-HPC cluster.
	Series *ListPreferredEcsTypesResponseBodySeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Struct"`
	// Indicates whether spot instances are supported.
	SupportSpotInstance *bool `json:"SupportSpotInstance,omitempty" xml:"SupportSpotInstance,omitempty"`
}

func (s ListPreferredEcsTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBody) SetRequestId(v string) *ListPreferredEcsTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPreferredEcsTypesResponseBody) SetSeries(v *ListPreferredEcsTypesResponseBodySeries) *ListPreferredEcsTypesResponseBody {
	s.Series = v
	return s
}

func (s *ListPreferredEcsTypesResponseBody) SetSupportSpotInstance(v bool) *ListPreferredEcsTypesResponseBody {
	s.SupportSpotInstance = &v
	return s
}

type ListPreferredEcsTypesResponseBodySeries struct {
	SeriesInfo []*ListPreferredEcsTypesResponseBodySeriesSeriesInfo `json:"SeriesInfo,omitempty" xml:"SeriesInfo,omitempty" type:"Repeated"`
}

func (s ListPreferredEcsTypesResponseBodySeries) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeries) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeries) SetSeriesInfo(v []*ListPreferredEcsTypesResponseBodySeriesSeriesInfo) *ListPreferredEcsTypesResponseBodySeries {
	s.SeriesInfo = v
	return s
}

type ListPreferredEcsTypesResponseBodySeriesSeriesInfo struct {
	// The recommended ECS instance types for various nodes of the E-HPC cluster.
	Roles *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
	// The ID of the ECS instance series. Valid values:
	//
	// *   HighCompute: high computing
	// *   HighMem: high storage
	// *   GPU
	// *   All: all options.
	SeriesId *string `json:"SeriesId,omitempty" xml:"SeriesId,omitempty"`
	// The name of the instance series. Valid values:
	//
	// *   SeriesHighCompute
	// *   SeriesHighMem
	// *   SeriesGPU
	// *   SeriesAll
	SeriesName *string `json:"SeriesName,omitempty" xml:"SeriesName,omitempty"`
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfo) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfo) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfo) SetRoles(v *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) *ListPreferredEcsTypesResponseBodySeriesSeriesInfo {
	s.Roles = v
	return s
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfo) SetSeriesId(v string) *ListPreferredEcsTypesResponseBodySeriesSeriesInfo {
	s.SeriesId = &v
	return s
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfo) SetSeriesName(v string) *ListPreferredEcsTypesResponseBodySeriesSeriesInfo {
	s.SeriesName = &v
	return s
}

type ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles struct {
	// The list of recommended ECS instance types for compute nodes.
	Compute *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute `json:"Compute,omitempty" xml:"Compute,omitempty" type:"Struct"`
	// The list of recommended ECS instance types for logon nodes.
	Login *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin `json:"Login,omitempty" xml:"Login,omitempty" type:"Struct"`
	// The list of recommended ECS instance types for management nodes.
	Manager *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager `json:"Manager,omitempty" xml:"Manager,omitempty" type:"Struct"`
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) SetCompute(v *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles {
	s.Compute = v
	return s
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) SetLogin(v *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles {
	s.Login = v
	return s
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) SetManager(v *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles {
	s.Manager = v
	return s
}

type ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute struct {
	InstanceTypeId []*string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty" type:"Repeated"`
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute) SetInstanceTypeId(v []*string) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute {
	s.InstanceTypeId = v
	return s
}

type ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin struct {
	InstanceTypeId []*string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty" type:"Repeated"`
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin) SetInstanceTypeId(v []*string) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin {
	s.InstanceTypeId = v
	return s
}

type ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager struct {
	InstanceTypeId []*string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty" type:"Repeated"`
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager) SetInstanceTypeId(v []*string) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager {
	s.InstanceTypeId = v
	return s
}

type ListPreferredEcsTypesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPreferredEcsTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPreferredEcsTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponse) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponse) SetHeaders(v map[string]*string) *ListPreferredEcsTypesResponse {
	s.Headers = v
	return s
}

func (s *ListPreferredEcsTypesResponse) SetStatusCode(v int32) *ListPreferredEcsTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPreferredEcsTypesResponse) SetBody(v *ListPreferredEcsTypesResponseBody) *ListPreferredEcsTypesResponse {
	s.Body = v
	return s
}

type ListQueuesRequest struct {
	// Specifies whether to enable asynchronous query.
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListQueuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesRequest) GoString() string {
	return s.String()
}

func (s *ListQueuesRequest) SetAsync(v bool) *ListQueuesRequest {
	s.Async = &v
	return s
}

func (s *ListQueuesRequest) SetClusterId(v string) *ListQueuesRequest {
	s.ClusterId = &v
	return s
}

type ListQueuesResponseBody struct {
	Queues *ListQueuesResponseBodyQueues `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListQueuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBody) SetQueues(v *ListQueuesResponseBodyQueues) *ListQueuesResponseBody {
	s.Queues = v
	return s
}

func (s *ListQueuesResponseBody) SetRequestId(v string) *ListQueuesResponseBody {
	s.RequestId = &v
	return s
}

type ListQueuesResponseBodyQueues struct {
	QueueInfo []*ListQueuesResponseBodyQueuesQueueInfo `json:"QueueInfo,omitempty" xml:"QueueInfo,omitempty" type:"Repeated"`
}

func (s ListQueuesResponseBodyQueues) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyQueues) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyQueues) SetQueueInfo(v []*ListQueuesResponseBodyQueuesQueueInfo) *ListQueuesResponseBodyQueues {
	s.QueueInfo = v
	return s
}

type ListQueuesResponseBodyQueuesQueueInfo struct {
	// The instance type of the compute nodes.
	ComputeInstanceType *ListQueuesResponseBodyQueuesQueueInfoComputeInstanceType `json:"ComputeInstanceType,omitempty" xml:"ComputeInstanceType,omitempty" type:"Struct"`
	// Indicates whether the queue enabled auto scale-out. Valid values:
	//
	// true: The queue enabled auto scale-out.
	//
	// false: The queue disabled auto scale-out.
	EnableAutoGrow *bool `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	// The prefix of the host name.
	HostNamePrefix *string `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	// The suffix of the host name.
	HostNameSuffix *string `json:"HostNameSuffix,omitempty" xml:"HostNameSuffix,omitempty"`
	// The ID of the image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the queue.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the resource group to which the queue belongs.
	ResourceGroupId   *string                                                 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SpotInstanceTypes *ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypes `json:"SpotInstanceTypes,omitempty" xml:"SpotInstanceTypes,omitempty" type:"Struct"`
	// The preemption policy of the compute nodes. Valid values:
	//
	// NoSpot: The instances of the compute node are pay-as-you-go instances.
	//
	// SpotWithPriceLimit: The instances of the compute node are preemptible instances. These types of instances have a specified maximum hourly price.
	//
	// SpotAsPriceGo: The instances of the compute node are preemptible instances. The price of these instances is based on the current market price.
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The type of the queue. Valid values:
	//
	// Execution: Queues in which jobs can be executed.
	//
	// Router: Queues in which jobs cannot be executed but are forwarded to the bounded Execution queue for processing.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListQueuesResponseBodyQueuesQueueInfo) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyQueuesQueueInfo) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetComputeInstanceType(v *ListQueuesResponseBodyQueuesQueueInfoComputeInstanceType) *ListQueuesResponseBodyQueuesQueueInfo {
	s.ComputeInstanceType = v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetEnableAutoGrow(v bool) *ListQueuesResponseBodyQueuesQueueInfo {
	s.EnableAutoGrow = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetHostNamePrefix(v string) *ListQueuesResponseBodyQueuesQueueInfo {
	s.HostNamePrefix = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetHostNameSuffix(v string) *ListQueuesResponseBodyQueuesQueueInfo {
	s.HostNameSuffix = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetImageId(v string) *ListQueuesResponseBodyQueuesQueueInfo {
	s.ImageId = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetQueueName(v string) *ListQueuesResponseBodyQueuesQueueInfo {
	s.QueueName = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetResourceGroupId(v string) *ListQueuesResponseBodyQueuesQueueInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetSpotInstanceTypes(v *ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypes) *ListQueuesResponseBodyQueuesQueueInfo {
	s.SpotInstanceTypes = v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetSpotStrategy(v string) *ListQueuesResponseBodyQueuesQueueInfo {
	s.SpotStrategy = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetType(v string) *ListQueuesResponseBodyQueuesQueueInfo {
	s.Type = &v
	return s
}

type ListQueuesResponseBodyQueuesQueueInfoComputeInstanceType struct {
	InstanceType []*string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Repeated"`
}

func (s ListQueuesResponseBodyQueuesQueueInfoComputeInstanceType) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyQueuesQueueInfoComputeInstanceType) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyQueuesQueueInfoComputeInstanceType) SetInstanceType(v []*string) *ListQueuesResponseBodyQueuesQueueInfoComputeInstanceType {
	s.InstanceType = v
	return s
}

type ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypes struct {
	Instance []*ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypes) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypes) SetInstance(v []*ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance) *ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypes {
	s.Instance = v
	return s
}

type ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance struct {
	// The specifications of the ECS instance.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum hourly price of the preemptible instance. The value can be accurate to three decimal places. The parameter takes effect only when SpotStrategy is set to SpotWithPriceLimit.
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
}

func (s ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance) SetInstanceType(v string) *ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance {
	s.InstanceType = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance) SetSpotPriceLimit(v float32) *ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance {
	s.SpotPriceLimit = &v
	return s
}

type ListQueuesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQueuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQueuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponse) GoString() string {
	return s.String()
}

func (s *ListQueuesResponse) SetHeaders(v map[string]*string) *ListQueuesResponse {
	s.Headers = v
	return s
}

func (s *ListQueuesResponse) SetStatusCode(v int32) *ListQueuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueuesResponse) SetBody(v *ListQueuesResponseBody) *ListQueuesResponse {
	s.Body = v
	return s
}

type ListRegionsResponseBody struct {
	// The array of regions.
	Regions *ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetRegions(v *ListRegionsResponseBodyRegions) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListRegionsResponseBodyRegions struct {
	RegionInfo []*ListRegionsResponseBodyRegionsRegionInfo `json:"RegionInfo,omitempty" xml:"RegionInfo,omitempty" type:"Repeated"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) SetRegionInfo(v []*ListRegionsResponseBodyRegionsRegionInfo) *ListRegionsResponseBodyRegions {
	s.RegionInfo = v
	return s
}

type ListRegionsResponseBodyRegionsRegionInfo struct {
	// The region name.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsResponseBodyRegionsRegionInfo) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegionsRegionInfo) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegionsRegionInfo) SetLocalName(v string) *ListRegionsResponseBodyRegionsRegionInfo {
	s.LocalName = &v
	return s
}

func (s *ListRegionsResponseBodyRegionsRegionInfo) SetRegionId(v string) *ListRegionsResponseBodyRegionsRegionInfo {
	s.RegionId = &v
	return s
}

type ListRegionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetStatusCode(v int32) *ListRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListSecurityGroupsRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListSecurityGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupsRequest) SetClusterId(v string) *ListSecurityGroupsRequest {
	s.ClusterId = &v
	return s
}

type ListSecurityGroupsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the security group.
	SecurityGroups *ListSecurityGroupsResponseBodySecurityGroups `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Struct"`
	// The number of security groups.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSecurityGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupsResponseBody) SetRequestId(v string) *ListSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecurityGroupsResponseBody) SetSecurityGroups(v *ListSecurityGroupsResponseBodySecurityGroups) *ListSecurityGroupsResponseBody {
	s.SecurityGroups = v
	return s
}

func (s *ListSecurityGroupsResponseBody) SetTotalCount(v int32) *ListSecurityGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSecurityGroupsResponseBodySecurityGroups struct {
	SecurityGroup []*string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty" type:"Repeated"`
}

func (s ListSecurityGroupsResponseBodySecurityGroups) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityGroupsResponseBodySecurityGroups) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupsResponseBodySecurityGroups) SetSecurityGroup(v []*string) *ListSecurityGroupsResponseBodySecurityGroups {
	s.SecurityGroup = v
	return s
}

type ListSecurityGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSecurityGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupsResponse) SetHeaders(v map[string]*string) *ListSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityGroupsResponse) SetStatusCode(v int32) *ListSecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecurityGroupsResponse) SetBody(v *ListSecurityGroupsResponseBody) *ListSecurityGroupsResponse {
	s.Body = v
	return s
}

type ListSoftwaresRequest struct {
	// The version of the E-HPC client.
	//
	// You can call the [ListCurrentClientVersion](~~87223~~) operation to query the E-HPC client version.
	EhpcVersion *string `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	// The image tag of the cluster.
	//
	// You can use the [ListImages](~~87213~~) to query the image tag of the cluster.
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
}

func (s ListSoftwaresRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresRequest) GoString() string {
	return s.String()
}

func (s *ListSoftwaresRequest) SetEhpcVersion(v string) *ListSoftwaresRequest {
	s.EhpcVersion = &v
	return s
}

func (s *ListSoftwaresRequest) SetOsTag(v string) *ListSoftwaresRequest {
	s.OsTag = &v
	return s
}

type ListSoftwaresResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of the information about the software installed in the cluster.
	Softwares *ListSoftwaresResponseBodySoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Struct"`
}

func (s ListSoftwaresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBody) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBody) SetRequestId(v string) *ListSoftwaresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSoftwaresResponseBody) SetSoftwares(v *ListSoftwaresResponseBodySoftwares) *ListSoftwaresResponseBody {
	s.Softwares = v
	return s
}

type ListSoftwaresResponseBodySoftwares struct {
	SoftwareInfo []*ListSoftwaresResponseBodySoftwaresSoftwareInfo `json:"SoftwareInfo,omitempty" xml:"SoftwareInfo,omitempty" type:"Repeated"`
}

func (s ListSoftwaresResponseBodySoftwares) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodySoftwares) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodySoftwares) SetSoftwareInfo(v []*ListSoftwaresResponseBodySoftwaresSoftwareInfo) *ListSoftwaresResponseBodySoftwares {
	s.SoftwareInfo = v
	return s
}

type ListSoftwaresResponseBodySoftwaresSoftwareInfo struct {
	// The service type of the domain account. Valid values:
	//
	// *   nis
	// *   ldap
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The version of the domain account service.
	AccountVersion *string `json:"AccountVersion,omitempty" xml:"AccountVersion,omitempty"`
	// The list of the software in the cluster.
	Applications *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	// The version of the E-HPC client.
	EhpcVersion *string `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	// The image tag of the cluster.
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	// The type of the scheduler. Valid values:
	//
	// *   pbs
	// *   pbs19
	// *   slurm
	// *   slurm19
	// *   slurm20
	// *   opengridscheduler
	// *   deadline
	// *   gridengine
	// *   cube
	// *   custom
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	// The version of the scheduler.
	SchedulerVersion *string `json:"SchedulerVersion,omitempty" xml:"SchedulerVersion,omitempty"`
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfo) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetAccountType(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.AccountType = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetAccountVersion(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.AccountVersion = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetApplications(v *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.Applications = v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetEhpcVersion(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.EhpcVersion = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetOsTag(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.OsTag = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetSchedulerType(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.SchedulerType = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetSchedulerVersion(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.SchedulerVersion = &v
	return s
}

type ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications struct {
	ApplicationInfo []*ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo `json:"ApplicationInfo,omitempty" xml:"ApplicationInfo,omitempty" type:"Repeated"`
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications) SetApplicationInfo(v []*ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications {
	s.ApplicationInfo = v
	return s
}

type ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo struct {
	// The name of the software.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the software is required. Valid values:
	//
	// *   false: optional
	// *   true: required
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// The tag of the software.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The version of the software.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) SetName(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo {
	s.Name = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) SetRequired(v bool) *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo {
	s.Required = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) SetTag(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo {
	s.Tag = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) SetVersion(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo {
	s.Version = &v
	return s
}

type ListSoftwaresResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSoftwaresResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSoftwaresResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponse) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponse) SetHeaders(v map[string]*string) *ListSoftwaresResponse {
	s.Headers = v
	return s
}

func (s *ListSoftwaresResponse) SetStatusCode(v int32) *ListSoftwaresResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSoftwaresResponse) SetBody(v *ListSoftwaresResponseBody) *ListSoftwaresResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The token used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region to which the resource belongs.
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Set the value to cluster, which indicates E-HPC clusters.
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The key of the tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// The token that is required for the next query. If the NextToken parameter is empty, no subsequent query will be sent.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of tags.
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	// The ID of the resource. Set the value to the ID of the cluster.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Set the value to cluster, which indicates E-HPC clusters.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTasksRequest struct {
	// Specifies whether to display the response history of the asynchronous API operation. Valid values:
	//
	// *   true: displays the current response and response history of the asynchronous API operation.
	// *   false: displays only the current response of the asynchronous API operation. If no tasks are running, `[]` is returned.
	//
	// Default value: false
	//
	// >  If you specify the TaskId parameter, the Archived parameter is invalid.
	Archived *bool `json:"Archived,omitempty" xml:"Archived,omitempty"`
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to obtain the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of the page to return. Pages start from page 1. Valid values: 1 to 999.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50.
	//
	// Default value: 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the task. You can call the following asynchronous API operations to obtain the task ID.
	//
	// *   [CreateCluster](~~87100~~)
	// *   [StartCluster](~~200345~~)
	// *   [StopCluster](~~200346~~)
	// *   [DeleteCluster](~~87110~~)
	// *   [AddNodes](~~87147~~)
	// *   [StartNodes](~~87159~~)
	// *   [ResetNodes](~~87158~~)
	// *   [StopNodes](~~87160~~)
	// *   [DeleteNodes](~~87155~~)
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) SetArchived(v bool) *ListTasksRequest {
	s.Archived = &v
	return s
}

func (s *ListTasksRequest) SetClusterId(v string) *ListTasksRequest {
	s.ClusterId = &v
	return s
}

func (s *ListTasksRequest) SetPageNumber(v int32) *ListTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTasksRequest) SetPageSize(v int32) *ListTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksRequest) SetTaskId(v string) *ListTasksRequest {
	s.TaskId = &v
	return s
}

type ListTasksResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of task information.
	Tasks []*ListTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of entries of the task.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) SetPageNumber(v int32) *ListTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTasksResponseBody) SetPageSize(v int32) *ListTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetTasks(v []*ListTasksResponseBodyTasks) *ListTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListTasksResponseBody) SetTotalCount(v int32) *ListTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListTasksResponseBodyTasks struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The current step of the task.
	CurrentStep *int32 `json:"CurrentStep,omitempty" xml:"CurrentStep,omitempty"`
	// The list of error messages returned for the task.
	//
	// For information about error messages and their solutions, visit the [API Error Center](https://error-center.alibabacloud.com/status/product/EHPC).
	Errors *string `json:"Errors,omitempty" xml:"Errors,omitempty"`
	// The request parameters of the task. The value is a JSON string.
	Request *string `json:"Request,omitempty" xml:"Request,omitempty"`
	// The result of the task. Valid values:
	//
	// *   If TaskType is set to CreateCluster and AddComputes, the value is in the `{\"Instances\":[]}` format, which indicates the information of the nodes added to the cluster.
	// *   If TaskType is set to a value other than CreateCluster and AddComputes, the value is in the `{}` format.
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The status of the task. Valid values:
	//
	// *   Processing: The task is running.
	// *   Success: The task is completed.
	// *   Fail: The task failed.
	// *   PartialFail: The task partially failed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the task. Valid values:
	//
	// *   CreateCluster: creates a cluster by calling the [CreateCluster](~~87100~~) operation.
	// *   StartCluster: starts a cluster by calling the [StartCluster](~~200345~~) operation.
	// *   StopCluster: stops a cluster by calling the [StopCluster](~~200346~~) operation.
	// *   DeleteCluster: releases a cluster by calling the [DeleteCluster](~~87110~~) operation.
	// *   AddComputes: adds nodes to a cluster by calling the [AddNodes](~~87147~~) operation.
	// *   StartComputes: starts nodes by calling the [StartNodes](~~87159~~) operation.
	// *   ResetCompute: resets nodes by calling the [ResetNodes](~~87158~~) operation.
	// *   StopComputes: stops nodes by calling the [StopNodes](~~87160~~) operation.
	// *   DeleteComputes: deletes nodes by calling the [DeleteNodes](~~87155~~) operation.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The total number of steps of the task.
	TotalSteps *int32 `json:"TotalSteps,omitempty" xml:"TotalSteps,omitempty"`
}

func (s ListTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTasks) SetClusterId(v string) *ListTasksResponseBodyTasks {
	s.ClusterId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetCurrentStep(v int32) *ListTasksResponseBodyTasks {
	s.CurrentStep = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetErrors(v string) *ListTasksResponseBodyTasks {
	s.Errors = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetRequest(v string) *ListTasksResponseBodyTasks {
	s.Request = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetResult(v string) *ListTasksResponseBodyTasks {
	s.Result = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetStatus(v string) *ListTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTaskId(v string) *ListTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTaskType(v string) *ListTasksResponseBodyTasks {
	s.TaskType = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTotalSteps(v int32) *ListTasksResponseBodyTasks {
	s.TotalSteps = &v
	return s
}

type ListTasksResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTasksResponse) SetHeaders(v map[string]*string) *ListTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTasksResponse) SetStatusCode(v int32) *ListTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTasksResponse) SetBody(v *ListTasksResponseBody) *ListTasksResponse {
	s.Body = v
	return s
}

type ListUpgradeClientsRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListUpgradeClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUpgradeClientsRequest) GoString() string {
	return s.String()
}

func (s *ListUpgradeClientsRequest) SetClusterId(v string) *ListUpgradeClientsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListUpgradeClientsRequest) SetRegionId(v string) *ListUpgradeClientsRequest {
	s.RegionId = &v
	return s
}

type ListUpgradeClientsResponseBody struct {
	// The upgrade records of the cluster.
	ClientRecords []*ListUpgradeClientsResponseBodyClientRecords `json:"ClientRecords,omitempty" xml:"ClientRecords,omitempty" type:"Repeated"`
	// The current version of the E-HPC client.
	CurrentVersion *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The latest version of the E-HPC client.
	LatestVersion *string `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUpgradeClientsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUpgradeClientsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUpgradeClientsResponseBody) SetClientRecords(v []*ListUpgradeClientsResponseBodyClientRecords) *ListUpgradeClientsResponseBody {
	s.ClientRecords = v
	return s
}

func (s *ListUpgradeClientsResponseBody) SetCurrentVersion(v string) *ListUpgradeClientsResponseBody {
	s.CurrentVersion = &v
	return s
}

func (s *ListUpgradeClientsResponseBody) SetLatestVersion(v string) *ListUpgradeClientsResponseBody {
	s.LatestVersion = &v
	return s
}

func (s *ListUpgradeClientsResponseBody) SetRequestId(v string) *ListUpgradeClientsResponseBody {
	s.RequestId = &v
	return s
}

type ListUpgradeClientsResponseBodyClientRecords struct {
	// The version of the E-HPC client after the upgrade.
	NewVersion *string `json:"NewVersion,omitempty" xml:"NewVersion,omitempty"`
	// The version of the E-HPC client before the upgrade.
	OldVersion *string `json:"OldVersion,omitempty" xml:"OldVersion,omitempty"`
	// The ID of the user that upgraded the E-HPC client.
	SubUid *string `json:"SubUid,omitempty" xml:"SubUid,omitempty"`
	// The time when the operation was performed.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListUpgradeClientsResponseBodyClientRecords) String() string {
	return tea.Prettify(s)
}

func (s ListUpgradeClientsResponseBodyClientRecords) GoString() string {
	return s.String()
}

func (s *ListUpgradeClientsResponseBodyClientRecords) SetNewVersion(v string) *ListUpgradeClientsResponseBodyClientRecords {
	s.NewVersion = &v
	return s
}

func (s *ListUpgradeClientsResponseBodyClientRecords) SetOldVersion(v string) *ListUpgradeClientsResponseBodyClientRecords {
	s.OldVersion = &v
	return s
}

func (s *ListUpgradeClientsResponseBodyClientRecords) SetSubUid(v string) *ListUpgradeClientsResponseBodyClientRecords {
	s.SubUid = &v
	return s
}

func (s *ListUpgradeClientsResponseBodyClientRecords) SetUpdateTime(v string) *ListUpgradeClientsResponseBodyClientRecords {
	s.UpdateTime = &v
	return s
}

type ListUpgradeClientsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUpgradeClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUpgradeClientsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUpgradeClientsResponse) GoString() string {
	return s.String()
}

func (s *ListUpgradeClientsResponse) SetHeaders(v map[string]*string) *ListUpgradeClientsResponse {
	s.Headers = v
	return s
}

func (s *ListUpgradeClientsResponse) SetStatusCode(v int32) *ListUpgradeClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUpgradeClientsResponse) SetBody(v *ListUpgradeClientsResponseBody) *ListUpgradeClientsResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50.
	//
	// Default value: 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetClusterId(v string) *ListUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *ListUsersRequest) SetPageNumber(v int32) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

type ListUsersResponseBody struct {
	// The number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of users.
	Users *ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetPageNumber(v int32) *ListUsersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBody) SetPageSize(v int32) *ListUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int32) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v *ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

type ListUsersResponseBodyUsers struct {
	UserInfo []*ListUsersResponseBodyUsersUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) SetUserInfo(v []*ListUsersResponseBodyUsersUserInfo) *ListUsersResponseBodyUsers {
	s.UserInfo = v
	return s
}

type ListUsersResponseBodyUsersUserInfo struct {
	// The time when the user was created.
	AddTime *string `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	// The name of the permission group. Valid values:
	//
	// *   users: an ordinary permission group. It is applicable to ordinary users that need only to submit and debug jobs.
	// *   wheel: a sudo permission group. It is applicable to the administrator who needs to manage the cluster. In addition to submitting and debugging jobs, users who have sudo permissions can run sudo commands to install software and restart nodes.
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The username.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListUsersResponseBodyUsersUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsersUserInfo) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsersUserInfo) SetAddTime(v string) *ListUsersResponseBodyUsersUserInfo {
	s.AddTime = &v
	return s
}

func (s *ListUsersResponseBodyUsersUserInfo) SetGroup(v string) *ListUsersResponseBodyUsersUserInfo {
	s.Group = &v
	return s
}

func (s *ListUsersResponseBodyUsersUserInfo) SetName(v string) *ListUsersResponseBodyUsersUserInfo {
	s.Name = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ListUsersAsyncRequest struct {
	AsyncId    *string `json:"AsyncId,omitempty" xml:"AsyncId,omitempty"`
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUsersAsyncRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersAsyncRequest) GoString() string {
	return s.String()
}

func (s *ListUsersAsyncRequest) SetAsyncId(v string) *ListUsersAsyncRequest {
	s.AsyncId = &v
	return s
}

func (s *ListUsersAsyncRequest) SetClusterId(v string) *ListUsersAsyncRequest {
	s.ClusterId = &v
	return s
}

func (s *ListUsersAsyncRequest) SetPageNumber(v int32) *ListUsersAsyncRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersAsyncRequest) SetPageSize(v int32) *ListUsersAsyncRequest {
	s.PageSize = &v
	return s
}

type ListUsersAsyncResponseBody struct {
	AsyncId     *string                          `json:"AsyncId,omitempty" xml:"AsyncId,omitempty"`
	AsyncStatus *string                          `json:"AsyncStatus,omitempty" xml:"AsyncStatus,omitempty"`
	PageNumber  *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Users       *ListUsersAsyncResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListUsersAsyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersAsyncResponseBody) SetAsyncId(v string) *ListUsersAsyncResponseBody {
	s.AsyncId = &v
	return s
}

func (s *ListUsersAsyncResponseBody) SetAsyncStatus(v string) *ListUsersAsyncResponseBody {
	s.AsyncStatus = &v
	return s
}

func (s *ListUsersAsyncResponseBody) SetPageNumber(v int32) *ListUsersAsyncResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUsersAsyncResponseBody) SetPageSize(v int32) *ListUsersAsyncResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUsersAsyncResponseBody) SetRequestId(v string) *ListUsersAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersAsyncResponseBody) SetTotalCount(v int32) *ListUsersAsyncResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersAsyncResponseBody) SetUsers(v *ListUsersAsyncResponseBodyUsers) *ListUsersAsyncResponseBody {
	s.Users = v
	return s
}

type ListUsersAsyncResponseBodyUsers struct {
	UserInfo []*ListUsersAsyncResponseBodyUsersUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Repeated"`
}

func (s ListUsersAsyncResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersAsyncResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersAsyncResponseBodyUsers) SetUserInfo(v []*ListUsersAsyncResponseBodyUsersUserInfo) *ListUsersAsyncResponseBodyUsers {
	s.UserInfo = v
	return s
}

type ListUsersAsyncResponseBodyUsersUserInfo struct {
	AddTime *string `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	Group   *string `json:"Group,omitempty" xml:"Group,omitempty"`
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UserId  *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUsersAsyncResponseBodyUsersUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListUsersAsyncResponseBodyUsersUserInfo) GoString() string {
	return s.String()
}

func (s *ListUsersAsyncResponseBodyUsersUserInfo) SetAddTime(v string) *ListUsersAsyncResponseBodyUsersUserInfo {
	s.AddTime = &v
	return s
}

func (s *ListUsersAsyncResponseBodyUsersUserInfo) SetGroup(v string) *ListUsersAsyncResponseBodyUsersUserInfo {
	s.Group = &v
	return s
}

func (s *ListUsersAsyncResponseBodyUsersUserInfo) SetGroupId(v string) *ListUsersAsyncResponseBodyUsersUserInfo {
	s.GroupId = &v
	return s
}

func (s *ListUsersAsyncResponseBodyUsersUserInfo) SetName(v string) *ListUsersAsyncResponseBodyUsersUserInfo {
	s.Name = &v
	return s
}

func (s *ListUsersAsyncResponseBodyUsersUserInfo) SetUserId(v string) *ListUsersAsyncResponseBodyUsersUserInfo {
	s.UserId = &v
	return s
}

type ListUsersAsyncResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUsersAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsersAsyncResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersAsyncResponse) GoString() string {
	return s.String()
}

func (s *ListUsersAsyncResponse) SetHeaders(v map[string]*string) *ListUsersAsyncResponse {
	s.Headers = v
	return s
}

func (s *ListUsersAsyncResponse) SetStatusCode(v int32) *ListUsersAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersAsyncResponse) SetBody(v *ListUsersAsyncResponseBody) *ListUsersAsyncResponse {
	s.Body = v
	return s
}

type ListVolumesRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50.
	//
	// Default value: 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVolumesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesRequest) GoString() string {
	return s.String()
}

func (s *ListVolumesRequest) SetPageNumber(v int32) *ListVolumesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVolumesRequest) SetPageSize(v int32) *ListVolumesRequest {
	s.PageSize = &v
	return s
}

type ListVolumesResponseBody struct {
	// The number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information of file systems that are mounted on E-HPC clusters.
	Volumes *ListVolumesResponseBodyVolumes `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Struct"`
}

func (s ListVolumesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVolumesResponseBody) SetPageNumber(v int32) *ListVolumesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListVolumesResponseBody) SetPageSize(v int32) *ListVolumesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVolumesResponseBody) SetRequestId(v string) *ListVolumesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVolumesResponseBody) SetTotalCount(v int32) *ListVolumesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVolumesResponseBody) SetVolumes(v *ListVolumesResponseBodyVolumes) *ListVolumesResponseBody {
	s.Volumes = v
	return s
}

type ListVolumesResponseBodyVolumes struct {
	VolumeInfo []*ListVolumesResponseBodyVolumesVolumeInfo `json:"VolumeInfo,omitempty" xml:"VolumeInfo,omitempty" type:"Repeated"`
}

func (s ListVolumesResponseBodyVolumes) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponseBodyVolumes) GoString() string {
	return s.String()
}

func (s *ListVolumesResponseBodyVolumes) SetVolumeInfo(v []*ListVolumesResponseBodyVolumesVolumeInfo) *ListVolumesResponseBodyVolumes {
	s.VolumeInfo = v
	return s
}

type ListVolumesResponseBodyVolumesVolumeInfo struct {
	// The information of additional file systems mounted on E-HPC clusters.
	AdditionalVolumes *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumes `json:"AdditionalVolumes,omitempty" xml:"AdditionalVolumes,omitempty" type:"Struct"`
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remote directory on which the file system is mounted.
	RemoteDirectory *string `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	// The ID of the file system.
	VolumeId *string `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	// The domain name of the mount target.
	VolumeMountpoint *string `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	// The type of the storage protocol. Valid values:
	//
	// *   NFS
	// *   SMB
	VolumeProtocol *string `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	// The type of the file system that is mounted on the cluster. Only NAS is supported.
	VolumeType *string `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
}

func (s ListVolumesResponseBodyVolumesVolumeInfo) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponseBodyVolumesVolumeInfo) GoString() string {
	return s.String()
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetAdditionalVolumes(v *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumes) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.AdditionalVolumes = v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetClusterId(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.ClusterId = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetClusterName(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.ClusterName = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetRegionId(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.RegionId = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetRemoteDirectory(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.RemoteDirectory = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetVolumeId(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeId = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetVolumeMountpoint(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeMountpoint = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetVolumeProtocol(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeProtocol = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetVolumeType(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeType = &v
	return s
}

type ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumes struct {
	VolumeInfo []*ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo `json:"VolumeInfo,omitempty" xml:"VolumeInfo,omitempty" type:"Repeated"`
}

func (s ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumes) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumes) GoString() string {
	return s.String()
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumes) SetVolumeInfo(v []*ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumes {
	s.VolumeInfo = v
	return s
}

type ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo struct {
	// The queue to which the job belongs.
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	// The local mount directory.
	LocalDirectory *string `json:"LocalDirectory,omitempty" xml:"LocalDirectory,omitempty"`
	// The location where the cluster was deployed. Valid values:
	//
	// *   OnPremise: The cluster is deployed on a hybrid cloud.
	// *   PublicCloud: The cluster is deployed on a public cloud.
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The remote directory on which the file system is mounted.
	RemoteDirectory *string `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	// The type of the node on which the file system is mounted. Valid values:
	//
	// *   Manager: management node
	// *   Login: logon node
	// *   Compute: compute node
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The ID of the file system.
	VolumeId *string `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	// The domain name of the mount target.
	VolumeMountpoint *string `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	// The type of the storage protocol. Valid values:
	//
	// *   NFS
	// *   SMB
	VolumeProtocol *string `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	// The type of the additional file system. Only NAS is supported.
	VolumeType *string `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
}

func (s ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) GoString() string {
	return s.String()
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetJobQueue(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.JobQueue = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetLocalDirectory(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.LocalDirectory = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetLocation(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.Location = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetRemoteDirectory(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.RemoteDirectory = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetRole(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.Role = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetVolumeId(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.VolumeId = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetVolumeMountpoint(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.VolumeMountpoint = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetVolumeProtocol(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.VolumeProtocol = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetVolumeType(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.VolumeType = &v
	return s
}

type ListVolumesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVolumesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVolumesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponse) GoString() string {
	return s.String()
}

func (s *ListVolumesResponse) SetHeaders(v map[string]*string) *ListVolumesResponse {
	s.Headers = v
	return s
}

func (s *ListVolumesResponse) SetStatusCode(v int32) *ListVolumesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVolumesResponse) SetBody(v *ListVolumesResponseBody) *ListVolumesResponse {
	s.Body = v
	return s
}

type ModifyClusterAttributesRequest struct {
	// The ID of the cluster that you want to modify.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The new cluster description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The new image type of the cluster. Valid values:
	//
	// *   system: public image
	// *   self: custom image
	// *   others: shared image
	// *   marketplace: Alibaba Cloud Marketplace image
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The new cluster name.
	Name         *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	RamNodeTypes []*string `json:"RamNodeTypes,omitempty" xml:"RamNodeTypes,omitempty" type:"Repeated"`
	RamRoleName  *string   `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
}

func (s ModifyClusterAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterAttributesRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterAttributesRequest) SetClusterId(v string) *ModifyClusterAttributesRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterAttributesRequest) SetDescription(v string) *ModifyClusterAttributesRequest {
	s.Description = &v
	return s
}

func (s *ModifyClusterAttributesRequest) SetImageId(v string) *ModifyClusterAttributesRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyClusterAttributesRequest) SetImageOwnerAlias(v string) *ModifyClusterAttributesRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ModifyClusterAttributesRequest) SetName(v string) *ModifyClusterAttributesRequest {
	s.Name = &v
	return s
}

func (s *ModifyClusterAttributesRequest) SetRamNodeTypes(v []*string) *ModifyClusterAttributesRequest {
	s.RamNodeTypes = v
	return s
}

func (s *ModifyClusterAttributesRequest) SetRamRoleName(v string) *ModifyClusterAttributesRequest {
	s.RamRoleName = &v
	return s
}

type ModifyClusterAttributesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterAttributesResponseBody) SetRequestId(v string) *ModifyClusterAttributesResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterAttributesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyClusterAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyClusterAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterAttributesResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterAttributesResponse) SetHeaders(v map[string]*string) *ModifyClusterAttributesResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterAttributesResponse) SetStatusCode(v int32) *ModifyClusterAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterAttributesResponse) SetBody(v *ModifyClusterAttributesResponseBody) *ModifyClusterAttributesResponse {
	s.Body = v
	return s
}

type ModifyContainerAppAttributesRequest struct {
	// The ID of the container.
	//
	// You can call the [ListContainerApps](~~87333~~) operation to query the ID of the containerized application.
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The new description of the containerized application.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyContainerAppAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyContainerAppAttributesRequest) GoString() string {
	return s.String()
}

func (s *ModifyContainerAppAttributesRequest) SetContainerId(v string) *ModifyContainerAppAttributesRequest {
	s.ContainerId = &v
	return s
}

func (s *ModifyContainerAppAttributesRequest) SetDescription(v string) *ModifyContainerAppAttributesRequest {
	s.Description = &v
	return s
}

type ModifyContainerAppAttributesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyContainerAppAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyContainerAppAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyContainerAppAttributesResponseBody) SetRequestId(v string) *ModifyContainerAppAttributesResponseBody {
	s.RequestId = &v
	return s
}

type ModifyContainerAppAttributesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyContainerAppAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyContainerAppAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyContainerAppAttributesResponse) GoString() string {
	return s.String()
}

func (s *ModifyContainerAppAttributesResponse) SetHeaders(v map[string]*string) *ModifyContainerAppAttributesResponse {
	s.Headers = v
	return s
}

func (s *ModifyContainerAppAttributesResponse) SetStatusCode(v int32) *ModifyContainerAppAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyContainerAppAttributesResponse) SetBody(v *ModifyContainerAppAttributesResponseBody) *ModifyContainerAppAttributesResponse {
	s.Body = v
	return s
}

type ModifyImageGatewayConfigRequest struct {
	ClusterId              *string                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DBPassword             *string                                `json:"DBPassword,omitempty" xml:"DBPassword,omitempty"`
	DBServerInfo           *string                                `json:"DBServerInfo,omitempty" xml:"DBServerInfo,omitempty"`
	DBType                 *string                                `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBUsername             *string                                `json:"DBUsername,omitempty" xml:"DBUsername,omitempty"`
	DefaultRepoLocation    *string                                `json:"DefaultRepoLocation,omitempty" xml:"DefaultRepoLocation,omitempty"`
	ImageExpirationTimeout *string                                `json:"ImageExpirationTimeout,omitempty" xml:"ImageExpirationTimeout,omitempty"`
	PullUpdateTimeout      *int32                                 `json:"PullUpdateTimeout,omitempty" xml:"PullUpdateTimeout,omitempty"`
	Repo                   []*ModifyImageGatewayConfigRequestRepo `json:"Repo,omitempty" xml:"Repo,omitempty" type:"Repeated"`
}

func (s ModifyImageGatewayConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageGatewayConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageGatewayConfigRequest) SetClusterId(v string) *ModifyImageGatewayConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetDBPassword(v string) *ModifyImageGatewayConfigRequest {
	s.DBPassword = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetDBServerInfo(v string) *ModifyImageGatewayConfigRequest {
	s.DBServerInfo = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetDBType(v string) *ModifyImageGatewayConfigRequest {
	s.DBType = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetDBUsername(v string) *ModifyImageGatewayConfigRequest {
	s.DBUsername = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetDefaultRepoLocation(v string) *ModifyImageGatewayConfigRequest {
	s.DefaultRepoLocation = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetImageExpirationTimeout(v string) *ModifyImageGatewayConfigRequest {
	s.ImageExpirationTimeout = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetPullUpdateTimeout(v int32) *ModifyImageGatewayConfigRequest {
	s.PullUpdateTimeout = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetRepo(v []*ModifyImageGatewayConfigRequestRepo) *ModifyImageGatewayConfigRequest {
	s.Repo = v
	return s
}

type ModifyImageGatewayConfigRequestRepo struct {
	Auth     *string `json:"Auth,omitempty" xml:"Auth,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	URL      *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s ModifyImageGatewayConfigRequestRepo) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageGatewayConfigRequestRepo) GoString() string {
	return s.String()
}

func (s *ModifyImageGatewayConfigRequestRepo) SetAuth(v string) *ModifyImageGatewayConfigRequestRepo {
	s.Auth = &v
	return s
}

func (s *ModifyImageGatewayConfigRequestRepo) SetLocation(v string) *ModifyImageGatewayConfigRequestRepo {
	s.Location = &v
	return s
}

func (s *ModifyImageGatewayConfigRequestRepo) SetURL(v string) *ModifyImageGatewayConfigRequestRepo {
	s.URL = &v
	return s
}

type ModifyImageGatewayConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyImageGatewayConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageGatewayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImageGatewayConfigResponseBody) SetRequestId(v string) *ModifyImageGatewayConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyImageGatewayConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyImageGatewayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyImageGatewayConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageGatewayConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageGatewayConfigResponse) SetHeaders(v map[string]*string) *ModifyImageGatewayConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyImageGatewayConfigResponse) SetStatusCode(v int32) *ModifyImageGatewayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyImageGatewayConfigResponse) SetBody(v *ModifyImageGatewayConfigResponseBody) *ModifyImageGatewayConfigResponse {
	s.Body = v
	return s
}

type ModifyUserGroupsRequest struct {
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string                        `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	User      []*ModifyUserGroupsRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ModifyUserGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupsRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupsRequest) SetAsync(v bool) *ModifyUserGroupsRequest {
	s.Async = &v
	return s
}

func (s *ModifyUserGroupsRequest) SetClusterId(v string) *ModifyUserGroupsRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyUserGroupsRequest) SetUser(v []*ModifyUserGroupsRequestUser) *ModifyUserGroupsRequest {
	s.User = v
	return s
}

type ModifyUserGroupsRequestUser struct {
	// The new permission group of the user. Valid values:
	//
	// *   users: an ordinary permission group. It is applicable to ordinary users that need only to submit and debug jobs.
	// *   wheel: a sudo permission group. It is applicable to the administrator who needs to manage the cluster. In addition to submitting and debugging jobs, users who have sudo permissions can run sudo commands to install software and restart nodes.
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The name of the user whose permissions you want to modify. Valid values of N: 1 to 100.
	//
	// You can call the [ListUsers](~~188572~~) operation to query the users of the cluster.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyUserGroupsRequestUser) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupsRequestUser) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupsRequestUser) SetGroup(v string) *ModifyUserGroupsRequestUser {
	s.Group = &v
	return s
}

func (s *ModifyUserGroupsRequestUser) SetName(v string) *ModifyUserGroupsRequestUser {
	s.Name = &v
	return s
}

type ModifyUserGroupsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupsResponseBody) SetRequestId(v string) *ModifyUserGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserGroupsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyUserGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUserGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupsResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupsResponse) SetHeaders(v map[string]*string) *ModifyUserGroupsResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserGroupsResponse) SetStatusCode(v int32) *ModifyUserGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserGroupsResponse) SetBody(v *ModifyUserGroupsResponseBody) *ModifyUserGroupsResponse {
	s.Body = v
	return s
}

type ModifyUserPasswordsRequest struct {
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The ID of the E-HPC cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string                           `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	User      []*ModifyUserPasswordsRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ModifyUserPasswordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserPasswordsRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordsRequest) SetAsync(v bool) *ModifyUserPasswordsRequest {
	s.Async = &v
	return s
}

func (s *ModifyUserPasswordsRequest) SetClusterId(v string) *ModifyUserPasswordsRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyUserPasswordsRequest) SetUser(v []*ModifyUserPasswordsRequestUser) *ModifyUserPasswordsRequest {
	s.User = v
	return s
}

type ModifyUserPasswordsRequestUser struct {
	// The name of the Nth user whose password you want to modify. Valid values of N: 1 to 100.
	//
	// You can call the [ListUsers](~~188572~~) operation to query the users of the cluster.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The password of the Nth user. The password must be 8 to 30 characters in length and contain three of the following items:
	//
	// *   Uppercase letter
	// *   Lowercase letter
	// *   Digit
	// *   Special character: `()~!@#$%^&*-_+=|{}[]:;\"/<>,.?/`
	//
	// Valid values of N: 1 to 100
	//
	// >  We recommend that you use HTTPS to call the AddUsers operation to ensure that the password remains confidential.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s ModifyUserPasswordsRequestUser) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserPasswordsRequestUser) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordsRequestUser) SetName(v string) *ModifyUserPasswordsRequestUser {
	s.Name = &v
	return s
}

func (s *ModifyUserPasswordsRequestUser) SetPassword(v string) *ModifyUserPasswordsRequestUser {
	s.Password = &v
	return s
}

type ModifyUserPasswordsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserPasswordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserPasswordsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordsResponseBody) SetRequestId(v string) *ModifyUserPasswordsResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserPasswordsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyUserPasswordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUserPasswordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserPasswordsResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordsResponse) SetHeaders(v map[string]*string) *ModifyUserPasswordsResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserPasswordsResponse) SetStatusCode(v int32) *ModifyUserPasswordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserPasswordsResponse) SetBody(v *ModifyUserPasswordsResponseBody) *ModifyUserPasswordsResponse {
	s.Body = v
	return s
}

type ModifyVisualServicePasswdRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The connection password of the VNC remote visualization service. The password must be 8 to 30 characters in length and include at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include:
	//
	// `()~! @#$%^&*-_+=|{}[]:;\"/<>,.? /`
	//
	// >  You must use HTTPS to call the API to ensure that the password remains confidential.
	Passwd *string `json:"Passwd,omitempty" xml:"Passwd,omitempty"`
	// The username of the cluster. Default value: root user. You can call the [ListUsers](~~188572~~) operation to query all users in a cluster.
	RunasUser *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	// The user password of the cluster.
	RunasUserPassword *string `json:"RunasUserPassword,omitempty" xml:"RunasUserPassword,omitempty"`
}

func (s ModifyVisualServicePasswdRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVisualServicePasswdRequest) GoString() string {
	return s.String()
}

func (s *ModifyVisualServicePasswdRequest) SetClusterId(v string) *ModifyVisualServicePasswdRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyVisualServicePasswdRequest) SetPasswd(v string) *ModifyVisualServicePasswdRequest {
	s.Passwd = &v
	return s
}

func (s *ModifyVisualServicePasswdRequest) SetRunasUser(v string) *ModifyVisualServicePasswdRequest {
	s.RunasUser = &v
	return s
}

func (s *ModifyVisualServicePasswdRequest) SetRunasUserPassword(v string) *ModifyVisualServicePasswdRequest {
	s.RunasUserPassword = &v
	return s
}

type ModifyVisualServicePasswdResponseBody struct {
	// The status of the VNC Remote visualization service. Valid values:
	//
	// *   Service started.: started
	// *   Service stopped.: stopped
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVisualServicePasswdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVisualServicePasswdResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVisualServicePasswdResponseBody) SetMessage(v string) *ModifyVisualServicePasswdResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyVisualServicePasswdResponseBody) SetRequestId(v string) *ModifyVisualServicePasswdResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVisualServicePasswdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVisualServicePasswdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyVisualServicePasswdResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVisualServicePasswdResponse) GoString() string {
	return s.String()
}

func (s *ModifyVisualServicePasswdResponse) SetHeaders(v map[string]*string) *ModifyVisualServicePasswdResponse {
	s.Headers = v
	return s
}

func (s *ModifyVisualServicePasswdResponse) SetStatusCode(v int32) *ModifyVisualServicePasswdResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVisualServicePasswdResponse) SetBody(v *ModifyVisualServicePasswdResponseBody) *ModifyVisualServicePasswdResponse {
	s.Body = v
	return s
}

type MountNFSRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MountDir     *string `json:"MountDir,omitempty" xml:"MountDir,omitempty"`
	NfsDir       *string `json:"NfsDir,omitempty" xml:"NfsDir,omitempty"`
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RemoteDir    *string `json:"RemoteDir,omitempty" xml:"RemoteDir,omitempty"`
}

func (s MountNFSRequest) String() string {
	return tea.Prettify(s)
}

func (s MountNFSRequest) GoString() string {
	return s.String()
}

func (s *MountNFSRequest) SetInstanceId(v string) *MountNFSRequest {
	s.InstanceId = &v
	return s
}

func (s *MountNFSRequest) SetMountDir(v string) *MountNFSRequest {
	s.MountDir = &v
	return s
}

func (s *MountNFSRequest) SetNfsDir(v string) *MountNFSRequest {
	s.NfsDir = &v
	return s
}

func (s *MountNFSRequest) SetProtocolType(v string) *MountNFSRequest {
	s.ProtocolType = &v
	return s
}

func (s *MountNFSRequest) SetRemoteDir(v string) *MountNFSRequest {
	s.RemoteDir = &v
	return s
}

type MountNFSResponseBody struct {
	InvokeId  *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MountNFSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MountNFSResponseBody) GoString() string {
	return s.String()
}

func (s *MountNFSResponseBody) SetInvokeId(v string) *MountNFSResponseBody {
	s.InvokeId = &v
	return s
}

func (s *MountNFSResponseBody) SetRequestId(v string) *MountNFSResponseBody {
	s.RequestId = &v
	return s
}

type MountNFSResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MountNFSResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MountNFSResponse) String() string {
	return tea.Prettify(s)
}

func (s MountNFSResponse) GoString() string {
	return s.String()
}

func (s *MountNFSResponse) SetHeaders(v map[string]*string) *MountNFSResponse {
	s.Headers = v
	return s
}

func (s *MountNFSResponse) SetStatusCode(v int32) *MountNFSResponse {
	s.StatusCode = &v
	return s
}

func (s *MountNFSResponse) SetBody(v *MountNFSResponseBody) *MountNFSResponse {
	s.Body = v
	return s
}

type PullImageRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the image. Default value: shifter.
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
	// The tag of the image. Default value: latest.
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The name of the repository.
	Repository *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
}

func (s PullImageRequest) String() string {
	return tea.Prettify(s)
}

func (s PullImageRequest) GoString() string {
	return s.String()
}

func (s *PullImageRequest) SetClusterId(v string) *PullImageRequest {
	s.ClusterId = &v
	return s
}

func (s *PullImageRequest) SetContainerType(v string) *PullImageRequest {
	s.ContainerType = &v
	return s
}

func (s *PullImageRequest) SetImageTag(v string) *PullImageRequest {
	s.ImageTag = &v
	return s
}

func (s *PullImageRequest) SetRepository(v string) *PullImageRequest {
	s.Repository = &v
	return s
}

type PullImageResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PullImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PullImageResponseBody) GoString() string {
	return s.String()
}

func (s *PullImageResponseBody) SetRequestId(v string) *PullImageResponseBody {
	s.RequestId = &v
	return s
}

type PullImageResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PullImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PullImageResponse) String() string {
	return tea.Prettify(s)
}

func (s PullImageResponse) GoString() string {
	return s.String()
}

func (s *PullImageResponse) SetHeaders(v map[string]*string) *PullImageResponse {
	s.Headers = v
	return s
}

func (s *PullImageResponse) SetStatusCode(v int32) *PullImageResponse {
	s.StatusCode = &v
	return s
}

func (s *PullImageResponse) SetBody(v *PullImageResponseBody) *PullImageResponse {
	s.Body = v
	return s
}

type QueryServicePackAndPriceResponseBody struct {
	ChargeAmount   *int32                                           `json:"ChargeAmount,omitempty" xml:"ChargeAmount,omitempty"`
	Currency       *string                                          `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice  *float32                                         `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalAmount *int32                                           `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	OriginalPrice  *float32                                         `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	RegionId       *string                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServicePack    *QueryServicePackAndPriceResponseBodyServicePack `json:"ServicePack,omitempty" xml:"ServicePack,omitempty" type:"Struct"`
	TradePrice     *float32                                         `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s QueryServicePackAndPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryServicePackAndPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryServicePackAndPriceResponseBody) SetChargeAmount(v int32) *QueryServicePackAndPriceResponseBody {
	s.ChargeAmount = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetCurrency(v string) *QueryServicePackAndPriceResponseBody {
	s.Currency = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetDiscountPrice(v float32) *QueryServicePackAndPriceResponseBody {
	s.DiscountPrice = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetOriginalAmount(v int32) *QueryServicePackAndPriceResponseBody {
	s.OriginalAmount = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetOriginalPrice(v float32) *QueryServicePackAndPriceResponseBody {
	s.OriginalPrice = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetRegionId(v string) *QueryServicePackAndPriceResponseBody {
	s.RegionId = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetRequestId(v string) *QueryServicePackAndPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetServicePack(v *QueryServicePackAndPriceResponseBodyServicePack) *QueryServicePackAndPriceResponseBody {
	s.ServicePack = v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetTradePrice(v float32) *QueryServicePackAndPriceResponseBody {
	s.TradePrice = &v
	return s
}

type QueryServicePackAndPriceResponseBodyServicePack struct {
	ServicePackInfo []*QueryServicePackAndPriceResponseBodyServicePackServicePackInfo `json:"ServicePackInfo,omitempty" xml:"ServicePackInfo,omitempty" type:"Repeated"`
}

func (s QueryServicePackAndPriceResponseBodyServicePack) String() string {
	return tea.Prettify(s)
}

func (s QueryServicePackAndPriceResponseBodyServicePack) GoString() string {
	return s.String()
}

func (s *QueryServicePackAndPriceResponseBodyServicePack) SetServicePackInfo(v []*QueryServicePackAndPriceResponseBodyServicePackServicePackInfo) *QueryServicePackAndPriceResponseBodyServicePack {
	s.ServicePackInfo = v
	return s
}

type QueryServicePackAndPriceResponseBodyServicePackServicePackInfo struct {
	Capacity     *int32  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	EndTime      *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	StartTime    *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryServicePackAndPriceResponseBodyServicePackServicePackInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryServicePackAndPriceResponseBodyServicePackServicePackInfo) GoString() string {
	return s.String()
}

func (s *QueryServicePackAndPriceResponseBodyServicePackServicePackInfo) SetCapacity(v int32) *QueryServicePackAndPriceResponseBodyServicePackServicePackInfo {
	s.Capacity = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBodyServicePackServicePackInfo) SetEndTime(v int32) *QueryServicePackAndPriceResponseBodyServicePackServicePackInfo {
	s.EndTime = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBodyServicePackServicePackInfo) SetInstanceName(v string) *QueryServicePackAndPriceResponseBodyServicePackServicePackInfo {
	s.InstanceName = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBodyServicePackServicePackInfo) SetStartTime(v int32) *QueryServicePackAndPriceResponseBodyServicePackServicePackInfo {
	s.StartTime = &v
	return s
}

type QueryServicePackAndPriceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryServicePackAndPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryServicePackAndPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryServicePackAndPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryServicePackAndPriceResponse) SetHeaders(v map[string]*string) *QueryServicePackAndPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryServicePackAndPriceResponse) SetStatusCode(v int32) *QueryServicePackAndPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryServicePackAndPriceResponse) SetBody(v *QueryServicePackAndPriceResponseBody) *QueryServicePackAndPriceResponse {
	s.Body = v
	return s
}

type RecoverClusterRequest struct {
	// The service type of the domain account. Valid values:
	//
	// *   nis
	// *   ldap
	//
	// Default value: nis
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The version of the E-HPC client. The default value is the latest version of the client.
	//
	// You can call the [ListCurrentClientVersion](~~87223~~) operation to query the current version of the E-HPC client.
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The ID of the cluster. The cluster must be in the Exception state.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID and status.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the image.
	//
	// You can call the [ListImages](~~87213~~) and [ListCustomImages](~~87215~~) operations to query the images that are supported by E-HPC.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The type of the image. Valid values:
	//
	// *   system: public image
	// *   self: custom image
	// *   others: shared image
	//
	// Default value: system
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The image tag of the operating system.
	//
	// You can call the [ListImages](~~87213~~) and [ListCustomImages](~~87215~~) operations to query the image tags supported by Elastic High Performance Computing (E-HPC).
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	// The type of the scheduler. Valid values:
	//
	// *   pbs
	// *   slurm
	// *   opengridscheduler
	// *   deadline
	//
	// Default value: pbs
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
}

func (s RecoverClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s RecoverClusterRequest) GoString() string {
	return s.String()
}

func (s *RecoverClusterRequest) SetAccountType(v string) *RecoverClusterRequest {
	s.AccountType = &v
	return s
}

func (s *RecoverClusterRequest) SetClientVersion(v string) *RecoverClusterRequest {
	s.ClientVersion = &v
	return s
}

func (s *RecoverClusterRequest) SetClusterId(v string) *RecoverClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *RecoverClusterRequest) SetImageId(v string) *RecoverClusterRequest {
	s.ImageId = &v
	return s
}

func (s *RecoverClusterRequest) SetImageOwnerAlias(v string) *RecoverClusterRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *RecoverClusterRequest) SetOsTag(v string) *RecoverClusterRequest {
	s.OsTag = &v
	return s
}

func (s *RecoverClusterRequest) SetSchedulerType(v string) *RecoverClusterRequest {
	s.SchedulerType = &v
	return s
}

type RecoverClusterResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RecoverClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecoverClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverClusterResponseBody) SetRequestId(v string) *RecoverClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoverClusterResponseBody) SetTaskId(v string) *RecoverClusterResponseBody {
	s.TaskId = &v
	return s
}

type RecoverClusterResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecoverClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecoverClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s RecoverClusterResponse) GoString() string {
	return s.String()
}

func (s *RecoverClusterResponse) SetHeaders(v map[string]*string) *RecoverClusterResponse {
	s.Headers = v
	return s
}

func (s *RecoverClusterResponse) SetStatusCode(v int32) *RecoverClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverClusterResponse) SetBody(v *RecoverClusterResponseBody) *RecoverClusterResponse {
	s.Body = v
	return s
}

type RerunJobsRequest struct {
	// Specifies whether to use an asynchronous link to rerun the job.
	//
	// Default value: false
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The list of jobs that you want to run. Maximum number of jobs: 100. Minimum number of jobs: 1.
	//
	// Format: `[{"Id": "0.sched****"},{"Id": "1.sched****"}]`. Separate multiple jobs with commas (,).
	//
	// You can call the [ListJobs](~~87251~~) operation to query the job ID.
	//
	// >  You can rerun only jobs that are in the RUNNING or QUEUED state.
	Jobs *string `json:"Jobs,omitempty" xml:"Jobs,omitempty"`
}

func (s RerunJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s RerunJobsRequest) GoString() string {
	return s.String()
}

func (s *RerunJobsRequest) SetAsync(v bool) *RerunJobsRequest {
	s.Async = &v
	return s
}

func (s *RerunJobsRequest) SetClusterId(v string) *RerunJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *RerunJobsRequest) SetJobs(v string) *RerunJobsRequest {
	s.Jobs = &v
	return s
}

type RerunJobsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RerunJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RerunJobsResponseBody) GoString() string {
	return s.String()
}

func (s *RerunJobsResponseBody) SetRequestId(v string) *RerunJobsResponseBody {
	s.RequestId = &v
	return s
}

type RerunJobsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RerunJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RerunJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s RerunJobsResponse) GoString() string {
	return s.String()
}

func (s *RerunJobsResponse) SetHeaders(v map[string]*string) *RerunJobsResponse {
	s.Headers = v
	return s
}

func (s *RerunJobsResponse) SetStatusCode(v int32) *RerunJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *RerunJobsResponse) SetBody(v *RerunJobsResponseBody) *RerunJobsResponse {
	s.Body = v
	return s
}

type ResetNodesRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string                      `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Instance  []*ResetNodesRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s ResetNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetNodesRequest) GoString() string {
	return s.String()
}

func (s *ResetNodesRequest) SetClusterId(v string) *ResetNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ResetNodesRequest) SetInstance(v []*ResetNodesRequestInstance) *ResetNodesRequest {
	s.Instance = v
	return s
}

type ResetNodesRequestInstance struct {
	// The ID of the compute node that you want to reset. Valid values of N: 1 to 100
	//
	// You can call the [ListNodes](~~87161~~) operation to query the IDs of the compute nodes.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ResetNodesRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s ResetNodesRequestInstance) GoString() string {
	return s.String()
}

func (s *ResetNodesRequestInstance) SetId(v string) *ResetNodesRequestInstance {
	s.Id = &v
	return s
}

type ResetNodesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ResetNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ResetNodesResponseBody) SetRequestId(v string) *ResetNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetNodesResponseBody) SetTaskId(v string) *ResetNodesResponseBody {
	s.TaskId = &v
	return s
}

type ResetNodesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetNodesResponse) GoString() string {
	return s.String()
}

func (s *ResetNodesResponse) SetHeaders(v map[string]*string) *ResetNodesResponse {
	s.Headers = v
	return s
}

func (s *ResetNodesResponse) SetStatusCode(v int32) *ResetNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetNodesResponse) SetBody(v *ResetNodesResponseBody) *ResetNodesResponse {
	s.Body = v
	return s
}

type RunCloudMetricProfilingRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The duration of the profiling process. Unit: seconds.
	//
	// Value values: 10 to 300
	//
	// Default value: 30
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The sampling frequency. Unit: Hz
	//
	// Valid values: 1 to 2000
	//
	// Default value: 2000
	Freq *int32 `json:"Freq,omitempty" xml:"Freq,omitempty"`
	// The name of the host.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the profiling process.
	ProcessId *int32 `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The ID of the region where the cluster resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RunCloudMetricProfilingRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCloudMetricProfilingRequest) GoString() string {
	return s.String()
}

func (s *RunCloudMetricProfilingRequest) SetClusterId(v string) *RunCloudMetricProfilingRequest {
	s.ClusterId = &v
	return s
}

func (s *RunCloudMetricProfilingRequest) SetDuration(v int32) *RunCloudMetricProfilingRequest {
	s.Duration = &v
	return s
}

func (s *RunCloudMetricProfilingRequest) SetFreq(v int32) *RunCloudMetricProfilingRequest {
	s.Freq = &v
	return s
}

func (s *RunCloudMetricProfilingRequest) SetHostName(v string) *RunCloudMetricProfilingRequest {
	s.HostName = &v
	return s
}

func (s *RunCloudMetricProfilingRequest) SetProcessId(v int32) *RunCloudMetricProfilingRequest {
	s.ProcessId = &v
	return s
}

func (s *RunCloudMetricProfilingRequest) SetRegionId(v string) *RunCloudMetricProfilingRequest {
	s.RegionId = &v
	return s
}

type RunCloudMetricProfilingResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunCloudMetricProfilingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCloudMetricProfilingResponseBody) GoString() string {
	return s.String()
}

func (s *RunCloudMetricProfilingResponseBody) SetRequestId(v string) *RunCloudMetricProfilingResponseBody {
	s.RequestId = &v
	return s
}

type RunCloudMetricProfilingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RunCloudMetricProfilingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunCloudMetricProfilingResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCloudMetricProfilingResponse) GoString() string {
	return s.String()
}

func (s *RunCloudMetricProfilingResponse) SetHeaders(v map[string]*string) *RunCloudMetricProfilingResponse {
	s.Headers = v
	return s
}

func (s *RunCloudMetricProfilingResponse) SetStatusCode(v int32) *RunCloudMetricProfilingResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCloudMetricProfilingResponse) SetBody(v *RunCloudMetricProfilingResponseBody) *RunCloudMetricProfilingResponse {
	s.Body = v
	return s
}

type SetAutoScaleConfigRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to enable auto scale-out. Valid values:
	//
	// *   true: enables auto scale-out.
	// *   false: disables auto scale-out
	//
	// Default value: false
	EnableAutoGrow *bool `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	// Specifies whether to enable auto scale-in. Valid values:
	//
	// *   true: enables auto scale-in.
	// *   false: disables auto scale-in
	//
	// Default value: false
	EnableAutoShrink *bool `json:"EnableAutoShrink,omitempty" xml:"EnableAutoShrink,omitempty"`
	// The compute nodes that are excluded from the list of auto scaling nodes. Separate multiple compute nodes with commas (,).
	//
	// If you want to retain a compute node, you can set the node as an exceptional node. Then, the node is not released if it is idle.
	ExcludeNodes *string `json:"ExcludeNodes,omitempty" xml:"ExcludeNodes,omitempty"`
	// The percentage of extra compute nodes. Valid values: 0 to 100
	//
	// Default value: 0
	//
	// If you need to add 100 compute nodes and the value of the ExtraNodesGrowRatio parameter is 2, 102 compute nodes are added.
	ExtraNodesGrowRatio *int32 `json:"ExtraNodesGrowRatio,omitempty" xml:"ExtraNodesGrowRatio,omitempty"`
	// The interval between two consecutive rounds of scale-out. Unit: minutes.
	//
	// Valid values: 2 to 10
	//
	// Default value: 2
	//
	// >  An interval may exist during multiple rounds of a scale-out task or between two consecutive scale-out tasks.
	GrowIntervalInMinutes *int32 `json:"GrowIntervalInMinutes,omitempty" xml:"GrowIntervalInMinutes,omitempty"`
	// The percentage of each round of scale-out. Valid values: 1 to 100.
	//
	// Default value: 100
	//
	// If you set GrowRatio to 50, the scale-out has two rounds. Each round completes half of the scale-out.
	GrowRatio *int32 `json:"GrowRatio,omitempty" xml:"GrowRatio,omitempty"`
	// The scale-out timeout period. Unit: minutes.
	//
	// Valid values: 10 to 60
	//
	// Default value: 20
	//
	// If the scale-out timeout period has been reached but the scale-out nodes still do not reach the Running state, the system resets them.
	GrowTimeoutInMinutes *int32 `json:"GrowTimeoutInMinutes,omitempty" xml:"GrowTimeoutInMinutes,omitempty"`
	// The ID of the image.
	//
	// >
	// *   If you set both `Queues.N.QueueImageId` and `ImageId`, `Queues.N.QueueImageId` prevails.
	// *   If you set `Queues.N.QueueImageId` or `ImageId`, the parameter that you set takes effect.
	// *   If you leave both `Queues.N.QueueImageId` and `ImageId` empty, the image that was specified when you created the cluster or the last time when you scaled out the cluster is used by default.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The maximum number of compute nodes that can be added in the cluster. Valid values: 0 to 500.
	//
	// Default value: 100
	MaxNodesInCluster *int32                             `json:"MaxNodesInCluster,omitempty" xml:"MaxNodesInCluster,omitempty"`
	Queues            []*SetAutoScaleConfigRequestQueues `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
	// The number of consecutive times that a compute node is idle during the resource scale-in check.
	//
	// Valid values: 2 to 5
	//
	// Default value: 3
	//
	// If the parameter is set to 3, a compute node is idle for more than three consecutive times. In this case, the node is released. If a compute node is idle for more than 6 minutes in a row, it is released by default. This is because the default value of the ShrinkIntervalInMinutes parameter is 2.
	ShrinkIdleTimes *int32 `json:"ShrinkIdleTimes,omitempty" xml:"ShrinkIdleTimes,omitempty"`
	// The interval between two consecutive rounds of scale-in. Unit: minutes.
	//
	// Valid values: 2 to 10
	//
	// Default value: 2
	ShrinkIntervalInMinutes *int32 `json:"ShrinkIntervalInMinutes,omitempty" xml:"ShrinkIntervalInMinutes,omitempty"`
	// The maximum hourly price of the compute nodes. The value can be accurate to three decimal places. The parameter takes effect only when `SpotStrategy` is set to `SpotWithPriceLimit`.
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The preemption policy of the compute nodes. Valid values:
	//
	// *   NoSpot: The compute nodes are pay-as-you-go instances.
	// *   SpotWithPriceLimit: The compute nodes are preemptible instances that have a user-defined maximum hourly price.
	// *   SpotAsPriceGo: The compute nodes are preemptible instances for which the market price at the time of purchase is used as the bid price.
	//
	// Default value: NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s SetAutoScaleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigRequest) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigRequest) SetClusterId(v string) *SetAutoScaleConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetEnableAutoGrow(v bool) *SetAutoScaleConfigRequest {
	s.EnableAutoGrow = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetEnableAutoShrink(v bool) *SetAutoScaleConfigRequest {
	s.EnableAutoShrink = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetExcludeNodes(v string) *SetAutoScaleConfigRequest {
	s.ExcludeNodes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetExtraNodesGrowRatio(v int32) *SetAutoScaleConfigRequest {
	s.ExtraNodesGrowRatio = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetGrowIntervalInMinutes(v int32) *SetAutoScaleConfigRequest {
	s.GrowIntervalInMinutes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetGrowRatio(v int32) *SetAutoScaleConfigRequest {
	s.GrowRatio = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetGrowTimeoutInMinutes(v int32) *SetAutoScaleConfigRequest {
	s.GrowTimeoutInMinutes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetImageId(v string) *SetAutoScaleConfigRequest {
	s.ImageId = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetMaxNodesInCluster(v int32) *SetAutoScaleConfigRequest {
	s.MaxNodesInCluster = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetQueues(v []*SetAutoScaleConfigRequestQueues) *SetAutoScaleConfigRequest {
	s.Queues = v
	return s
}

func (s *SetAutoScaleConfigRequest) SetShrinkIdleTimes(v int32) *SetAutoScaleConfigRequest {
	s.ShrinkIdleTimes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetShrinkIntervalInMinutes(v int32) *SetAutoScaleConfigRequest {
	s.ShrinkIntervalInMinutes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetSpotPriceLimit(v float32) *SetAutoScaleConfigRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetSpotStrategy(v string) *SetAutoScaleConfigRequest {
	s.SpotStrategy = &v
	return s
}

type SetAutoScaleConfigRequestQueues struct {
	DataDisks []*SetAutoScaleConfigRequestQueuesDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// Specifies whether the queue enables auto scale-out. Valid values:
	//
	// *   true: enables auto scale-out.
	// *   false: disables auto scale-out
	//
	// Valid values of N: 1 to 8
	//
	// Default value: false
	EnableAutoGrow *bool `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	// Specifies whether the queue enables auto scale-in. Valid values:
	//
	// *   true: enables auto scale-in.
	// *   false: disables auto scale-in
	//
	// Valid values of N: 1 to 8
	//
	// Default value: false
	EnableAutoShrink *bool `json:"EnableAutoShrink,omitempty" xml:"EnableAutoShrink,omitempty"`
	// The hostname prefix of the host that is used to perform scale-out for the queue. You can manage compute nodes that have a specified hostname prefix.
	//
	// Valid values of N: 1 to 8
	HostNamePrefix *string `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	// The hostname suffix of the host that is used to perform scale-out for the queue. You can manage nodes that have a specified hostname suffix.
	//
	// Valid values of N: 1 to 8
	HostNameSuffix *string `json:"HostNameSuffix,omitempty" xml:"HostNameSuffix,omitempty"`
	// The instance type of the compute nodes that are automatically added in the queue. Valid values of N: 1 to 8
	InstanceType  *string                                         `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceTypes []*SetAutoScaleConfigRequestQueuesInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The maximum number of the compute nodes that can be added in the queue. Valid values: 0 to 500.
	//
	// Valid values of N: 1 to 8
	//
	// Default value: 100
	MaxNodesInQueue *int32 `json:"MaxNodesInQueue,omitempty" xml:"MaxNodesInQueue,omitempty"`
	// The maximum number of compute nodes that can be added in each round of scale-out. Valid values: 0 to 99.
	//
	// Default value: 0
	MaxNodesPerCycle *int64 `json:"MaxNodesPerCycle,omitempty" xml:"MaxNodesPerCycle,omitempty"`
	// The minimum number of the compute nodes that can be removed in the queue. Valid values: 0 to 50.
	//
	// Valid values of N: 1 to 8
	//
	// Default value: 0
	MinNodesInQueue *int32 `json:"MinNodesInQueue,omitempty" xml:"MinNodesInQueue,omitempty"`
	// The minimum number of compute nodes that can be added in each round of scale-out. Valid values: 1 to 99.
	//
	// Default value: 1
	//
	// If the compute nodes that you want to add in a round is less than the minimum compute nodes that can be added, the value of this parameter is automatically changed to the number of compute nodes that you want to add. This ensures that compute nodes can be added as expected.
	//
	// >  The configuration takes effect only for the minimum compute nodes that can be added in the current round.
	MinNodesPerCycle *int64 `json:"MinNodesPerCycle,omitempty" xml:"MinNodesPerCycle,omitempty"`
	// The image ID of the queue where scale-out is performed. Valid values of N: 1 to 8.
	//
	// >
	// *   If you set both `Queues.N.QueueImageId` and `ImageId`, `Queues.N.QueueImageId` prevails.
	// *   If you set `Queues.N.QueueImageId` or `ImageId`, the parameter that you set takes effect.
	// *   If you leave both `Queues.N.QueueImageId` and `ImageId` empty, the image that was specified when you created the cluster or the last time when you scaled out the cluster is used by default.
	QueueImageId *string `json:"QueueImageId,omitempty" xml:"QueueImageId,omitempty"`
	// The name of the queue. N queue names can be set at the same time. Valid values of N: 1 to 8.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The maximum hourly price of the compute nodes that are automatically added in the queue. The value can be accurate to three decimal places. The parameter takes effect only when `Queues.N.SpotStrategy` is set to `SpotWithPriceLimit`.
	//
	// Valid values of N: 1 to 8
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding method of the compute nodes that are automatically added in the queue. Valid values of N: 1 to 8
	//
	// Valid values:
	//
	// *   NoSpot: The compute nodes are pay-as-you-go instances.
	// *   SpotWithPriceLimit: The compute nodes are preemptible instances that have a user-defined maximum hourly price.
	// *   SpotAsPriceGo: The compute nodes are preemptible instances for which the market price at the time of purchase is used as the bid price.
	//
	// Default value: NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The type of the system disk specified for the compute nodes that are added in the queue. Valid values:
	//
	// *   cloud_efficiency: ultra disk.
	// *   cloud_ssd: SSD.
	// *   cloud_essd: ESSD.
	// *   cloud: basic disk. Disks of this type are retired.
	//
	// Valid values of N: 1 to 8
	//
	// Default value: cloud_efficiency
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The performance level of the system disk specified for the compute nodes that are added in the queue. Valid values:
	//
	// *   PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	// *   PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	// *   PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	// *   PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// Valid values of N: 1 to 8
	//
	// Default value: PL1
	SystemDiskLevel *string `json:"SystemDiskLevel,omitempty" xml:"SystemDiskLevel,omitempty"`
	// The size of the system disk specified for the compute nodes that are added in the queue. Unit: GB.
	//
	// Valid values: 40 to 500
	//
	// Valid values of N: 1 to 8
	//
	// Default value: 40
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s SetAutoScaleConfigRequestQueues) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigRequestQueues) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigRequestQueues) SetDataDisks(v []*SetAutoScaleConfigRequestQueuesDataDisks) *SetAutoScaleConfigRequestQueues {
	s.DataDisks = v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetEnableAutoGrow(v bool) *SetAutoScaleConfigRequestQueues {
	s.EnableAutoGrow = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetEnableAutoShrink(v bool) *SetAutoScaleConfigRequestQueues {
	s.EnableAutoShrink = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetHostNamePrefix(v string) *SetAutoScaleConfigRequestQueues {
	s.HostNamePrefix = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetHostNameSuffix(v string) *SetAutoScaleConfigRequestQueues {
	s.HostNameSuffix = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetInstanceType(v string) *SetAutoScaleConfigRequestQueues {
	s.InstanceType = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetInstanceTypes(v []*SetAutoScaleConfigRequestQueuesInstanceTypes) *SetAutoScaleConfigRequestQueues {
	s.InstanceTypes = v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetMaxNodesInQueue(v int32) *SetAutoScaleConfigRequestQueues {
	s.MaxNodesInQueue = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetMaxNodesPerCycle(v int64) *SetAutoScaleConfigRequestQueues {
	s.MaxNodesPerCycle = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetMinNodesInQueue(v int32) *SetAutoScaleConfigRequestQueues {
	s.MinNodesInQueue = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetMinNodesPerCycle(v int64) *SetAutoScaleConfigRequestQueues {
	s.MinNodesPerCycle = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetQueueImageId(v string) *SetAutoScaleConfigRequestQueues {
	s.QueueImageId = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetQueueName(v string) *SetAutoScaleConfigRequestQueues {
	s.QueueName = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetSpotPriceLimit(v float32) *SetAutoScaleConfigRequestQueues {
	s.SpotPriceLimit = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetSpotStrategy(v string) *SetAutoScaleConfigRequestQueues {
	s.SpotStrategy = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetSystemDiskCategory(v string) *SetAutoScaleConfigRequestQueues {
	s.SystemDiskCategory = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetSystemDiskLevel(v string) *SetAutoScaleConfigRequestQueues {
	s.SystemDiskLevel = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetSystemDiskSize(v int32) *SetAutoScaleConfigRequestQueues {
	s.SystemDiskSize = &v
	return s
}

type SetAutoScaleConfigRequestQueuesDataDisks struct {
	// The type of the data disk. Valid values:
	//
	// *   cloud_efficiency: ultra disk
	// *   cloud_ssd: SSD
	// *   cloud_essd: ESSD
	// *   cloud: basic disk
	//
	// Default value: cloud_efficiency
	//
	// Valid values of N: 0 to 16
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// Specifies whether the data disk is released when the node is released. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: true
	//
	// Valid values of N: 0 to 16
	DataDiskDeleteWithInstance *bool `json:"DataDiskDeleteWithInstance,omitempty" xml:"DataDiskDeleteWithInstance,omitempty"`
	// Specifies whether to encrypt the data disk. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false
	//
	// Valid values of N: 0 to 16
	DataDiskEncrypted *bool `json:"DataDiskEncrypted,omitempty" xml:"DataDiskEncrypted,omitempty"`
	// The KMS key ID of the data disk.
	//
	// Valid values of N: 0 to 16
	DataDiskKMSKeyId *string `json:"DataDiskKMSKeyId,omitempty" xml:"DataDiskKMSKeyId,omitempty"`
	// The performance level of the ESSD used as the data disk. The parameter takes effect only when the Queues.N.DataDisks.N.DataDiskCategory parameter is set to cloud_essd. Valid values:
	//
	// *   PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	// *   PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	// *   PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	// *   PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL1
	//
	// Valid values of N: 0 to 16
	DataDiskPerformanceLevel *string `json:"DataDiskPerformanceLevel,omitempty" xml:"DataDiskPerformanceLevel,omitempty"`
	// The size of the data disk. Unit: GB.
	//
	// Valid values: 40 to 500
	//
	// Default value: 40
	//
	// Valid values of N: 0 to 16
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
}

func (s SetAutoScaleConfigRequestQueuesDataDisks) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigRequestQueuesDataDisks) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigRequestQueuesDataDisks) SetDataDiskCategory(v string) *SetAutoScaleConfigRequestQueuesDataDisks {
	s.DataDiskCategory = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesDataDisks) SetDataDiskDeleteWithInstance(v bool) *SetAutoScaleConfigRequestQueuesDataDisks {
	s.DataDiskDeleteWithInstance = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesDataDisks) SetDataDiskEncrypted(v bool) *SetAutoScaleConfigRequestQueuesDataDisks {
	s.DataDiskEncrypted = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesDataDisks) SetDataDiskKMSKeyId(v string) *SetAutoScaleConfigRequestQueuesDataDisks {
	s.DataDiskKMSKeyId = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesDataDisks) SetDataDiskPerformanceLevel(v string) *SetAutoScaleConfigRequestQueuesDataDisks {
	s.DataDiskPerformanceLevel = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesDataDisks) SetDataDiskSize(v int32) *SetAutoScaleConfigRequestQueuesDataDisks {
	s.DataDiskSize = &v
	return s
}

type SetAutoScaleConfigRequestQueuesInstanceTypes struct {
	// The instance type of the compute nodes that are automatically added in the queue.
	//
	// N queue names can be set at the same time. Valid values of N: 1 to 8.
	//
	// The instance types of N compute nodes in the queue can be set at the same time when auto scaling is performed in the queue. Valid values of N: 0 to 500.
	InstanceType             *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	SpotDuration             *int32  `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	// The maximum hourly price of the compute nodes that are automatically added in the queue. The value can be accurate to three decimal places. The parameter takes effect only when `Queues.N.InstanceTypes.N.SpotStrategy` is set to `Queues.N.InstanceTypes.N.SpotStrategy`.
	//
	// The maximum hourly prices of the compute nodes that are automatically added in N queues can be set the same time. Valid values of N: 1 to 8.
	//
	// The maximum hourly prices of N compute nodes in the queue can be set at the same time when auto scaling is performed in the queue. Valid values of N: 0 to 500.
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding method of the compute nodes that are automatically added in the queue. Valid values:
	//
	// *   NoSpot: The compute nodes are pay-as-you-go instances.
	// *   SpotWithPriceLimit: The compute nodes are preemptible instances that have a user-defined maximum hourly price.
	// *   SpotAsPriceGo: The compute nodes are preemptible instances for which the market price at the time of purchase is used as the bid price.
	//
	// Default value: NoSpot
	//
	// N queue names can be set at the same time. Valid values of N: 1 to 8.
	//
	// The bidding methods of N compute nodes in the queue can be set at the same time when auto scaling is performed in the queue. Valid values of N: 0 to 500.
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The vSwitch ID of the compute nodes that are automatically added in the queue.
	//
	// N queue names can be set at the same time. Valid values of N: 1 to 8.
	//
	// The vSwitch IDs of N compute nodes in the queue can be set at the same time when auto scaling is performed in the queue. Valid values of N: 0 to 500.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the compute nodes that are automatically added in the queue belongs.
	//
	// N queue names can be set at the same time. Valid values of N: 1 to 8.
	//
	// The zone IDs of N compute nodes in the queue can be set at the same time when auto scaling is performed in the queue. Valid values of N: 0 to 500.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SetAutoScaleConfigRequestQueuesInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigRequestQueuesInstanceTypes) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetInstanceType(v string) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.InstanceType = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetSpotDuration(v int32) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.SpotDuration = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetSpotInterruptionBehavior(v string) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetSpotPriceLimit(v float32) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.SpotPriceLimit = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetSpotStrategy(v string) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.SpotStrategy = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetVSwitchId(v string) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.VSwitchId = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetZoneId(v string) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.ZoneId = &v
	return s
}

type SetAutoScaleConfigResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAutoScaleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigResponseBody) SetRequestId(v string) *SetAutoScaleConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetAutoScaleConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetAutoScaleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAutoScaleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigResponse) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigResponse) SetHeaders(v map[string]*string) *SetAutoScaleConfigResponse {
	s.Headers = v
	return s
}

func (s *SetAutoScaleConfigResponse) SetStatusCode(v int32) *SetAutoScaleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAutoScaleConfigResponse) SetBody(v *SetAutoScaleConfigResponseBody) *SetAutoScaleConfigResponse {
	s.Body = v
	return s
}

type SetGWSClusterPolicyRequest struct {
	AsyncMode   *bool   `json:"AsyncMode,omitempty" xml:"AsyncMode,omitempty"`
	Clipboard   *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	LocalDrive  *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	UdpPort     *string `json:"UdpPort,omitempty" xml:"UdpPort,omitempty"`
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	Watermark   *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s SetGWSClusterPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGWSClusterPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetGWSClusterPolicyRequest) SetAsyncMode(v bool) *SetGWSClusterPolicyRequest {
	s.AsyncMode = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetClipboard(v string) *SetGWSClusterPolicyRequest {
	s.Clipboard = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetClusterId(v string) *SetGWSClusterPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetLocalDrive(v string) *SetGWSClusterPolicyRequest {
	s.LocalDrive = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetUdpPort(v string) *SetGWSClusterPolicyRequest {
	s.UdpPort = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetUsbRedirect(v string) *SetGWSClusterPolicyRequest {
	s.UsbRedirect = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetWatermark(v string) *SetGWSClusterPolicyRequest {
	s.Watermark = &v
	return s
}

type SetGWSClusterPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetGWSClusterPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGWSClusterPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetGWSClusterPolicyResponseBody) SetRequestId(v string) *SetGWSClusterPolicyResponseBody {
	s.RequestId = &v
	return s
}

type SetGWSClusterPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetGWSClusterPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGWSClusterPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGWSClusterPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetGWSClusterPolicyResponse) SetHeaders(v map[string]*string) *SetGWSClusterPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetGWSClusterPolicyResponse) SetStatusCode(v int32) *SetGWSClusterPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *SetGWSClusterPolicyResponse) SetBody(v *SetGWSClusterPolicyResponseBody) *SetGWSClusterPolicyResponse {
	s.Body = v
	return s
}

type SetGWSInstanceNameRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SetGWSInstanceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceNameRequest) SetInstanceId(v string) *SetGWSInstanceNameRequest {
	s.InstanceId = &v
	return s
}

func (s *SetGWSInstanceNameRequest) SetName(v string) *SetGWSInstanceNameRequest {
	s.Name = &v
	return s
}

type SetGWSInstanceNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetGWSInstanceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceNameResponseBody) SetRequestId(v string) *SetGWSInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

type SetGWSInstanceNameResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetGWSInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGWSInstanceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceNameResponse) SetHeaders(v map[string]*string) *SetGWSInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *SetGWSInstanceNameResponse) SetStatusCode(v int32) *SetGWSInstanceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *SetGWSInstanceNameResponse) SetBody(v *SetGWSInstanceNameResponseBody) *SetGWSInstanceNameResponse {
	s.Body = v
	return s
}

type SetGWSInstanceUserRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserUid    *string `json:"UserUid,omitempty" xml:"UserUid,omitempty"`
}

func (s SetGWSInstanceUserRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceUserRequest) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceUserRequest) SetInstanceId(v string) *SetGWSInstanceUserRequest {
	s.InstanceId = &v
	return s
}

func (s *SetGWSInstanceUserRequest) SetUserName(v string) *SetGWSInstanceUserRequest {
	s.UserName = &v
	return s
}

func (s *SetGWSInstanceUserRequest) SetUserUid(v string) *SetGWSInstanceUserRequest {
	s.UserUid = &v
	return s
}

type SetGWSInstanceUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetGWSInstanceUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceUserResponseBody) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceUserResponseBody) SetRequestId(v string) *SetGWSInstanceUserResponseBody {
	s.RequestId = &v
	return s
}

type SetGWSInstanceUserResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetGWSInstanceUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGWSInstanceUserResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceUserResponse) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceUserResponse) SetHeaders(v map[string]*string) *SetGWSInstanceUserResponse {
	s.Headers = v
	return s
}

func (s *SetGWSInstanceUserResponse) SetStatusCode(v int32) *SetGWSInstanceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *SetGWSInstanceUserResponse) SetBody(v *SetGWSInstanceUserResponseBody) *SetGWSInstanceUserResponse {
	s.Body = v
	return s
}

type SetPostScriptsRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId          *string                                    `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PostInstallScripts []*SetPostScriptsRequestPostInstallScripts `json:"PostInstallScripts,omitempty" xml:"PostInstallScripts,omitempty" type:"Repeated"`
	// The ID of the region where the cluster resides. You can call the [ListRegions](~~188593~~) operation to query the latest region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetPostScriptsRequest) String() string {
	return tea.Prettify(s)
}

func (s SetPostScriptsRequest) GoString() string {
	return s.String()
}

func (s *SetPostScriptsRequest) SetClusterId(v string) *SetPostScriptsRequest {
	s.ClusterId = &v
	return s
}

func (s *SetPostScriptsRequest) SetPostInstallScripts(v []*SetPostScriptsRequestPostInstallScripts) *SetPostScriptsRequest {
	s.PostInstallScripts = v
	return s
}

func (s *SetPostScriptsRequest) SetRegionId(v string) *SetPostScriptsRequest {
	s.RegionId = &v
	return s
}

type SetPostScriptsRequestPostInstallScripts struct {
	// The parameter that is used to run the Nth post-installation script. Valid values of N: 1 to 16.
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The URL that is used to download the Nth post-installation script. Valid values of N: 1 to 16.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SetPostScriptsRequestPostInstallScripts) String() string {
	return tea.Prettify(s)
}

func (s SetPostScriptsRequestPostInstallScripts) GoString() string {
	return s.String()
}

func (s *SetPostScriptsRequestPostInstallScripts) SetArgs(v string) *SetPostScriptsRequestPostInstallScripts {
	s.Args = &v
	return s
}

func (s *SetPostScriptsRequestPostInstallScripts) SetUrl(v string) *SetPostScriptsRequestPostInstallScripts {
	s.Url = &v
	return s
}

type SetPostScriptsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPostScriptsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetPostScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *SetPostScriptsResponseBody) SetRequestId(v string) *SetPostScriptsResponseBody {
	s.RequestId = &v
	return s
}

type SetPostScriptsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetPostScriptsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetPostScriptsResponse) String() string {
	return tea.Prettify(s)
}

func (s SetPostScriptsResponse) GoString() string {
	return s.String()
}

func (s *SetPostScriptsResponse) SetHeaders(v map[string]*string) *SetPostScriptsResponse {
	s.Headers = v
	return s
}

func (s *SetPostScriptsResponse) SetStatusCode(v int32) *SetPostScriptsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPostScriptsResponse) SetBody(v *SetPostScriptsResponseBody) *SetPostScriptsResponse {
	s.Body = v
	return s
}

type SetQueueRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Node      []*SetQueueRequestNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Repeated"`
	// The name of the destination queue.
	//
	// You can call the [ListQueues](~~92176~~) operation to query the queue name.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s SetQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s SetQueueRequest) GoString() string {
	return s.String()
}

func (s *SetQueueRequest) SetClusterId(v string) *SetQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *SetQueueRequest) SetNode(v []*SetQueueRequestNode) *SetQueueRequest {
	s.Node = v
	return s
}

func (s *SetQueueRequest) SetQueueName(v string) *SetQueueRequest {
	s.QueueName = &v
	return s
}

type SetQueueRequestNode struct {
	// The name of the compute node that you want to move. Valid values of N: 1 to 100.
	//
	// You can call the [ListNodes](~~87161~~) operation to query the names of the compute nodes.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SetQueueRequestNode) String() string {
	return tea.Prettify(s)
}

func (s SetQueueRequestNode) GoString() string {
	return s.String()
}

func (s *SetQueueRequestNode) SetName(v string) *SetQueueRequestNode {
	s.Name = &v
	return s
}

type SetQueueResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetQueueResponseBody) GoString() string {
	return s.String()
}

func (s *SetQueueResponseBody) SetRequestId(v string) *SetQueueResponseBody {
	s.RequestId = &v
	return s
}

type SetQueueResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s SetQueueResponse) GoString() string {
	return s.String()
}

func (s *SetQueueResponse) SetHeaders(v map[string]*string) *SetQueueResponse {
	s.Headers = v
	return s
}

func (s *SetQueueResponse) SetStatusCode(v int32) *SetQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *SetQueueResponse) SetBody(v *SetQueueResponseBody) *SetQueueResponse {
	s.Body = v
	return s
}

type SetSchedulerInfoRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string                           `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PbsInfo   []*SetSchedulerInfoRequestPbsInfo `json:"PbsInfo,omitempty" xml:"PbsInfo,omitempty" type:"Repeated"`
	// The ID of the region.
	//
	// You can call the [ListRegions](~~188593~~) operation to obtain the IDs of regions supported by Elastic High Performance Computing (E-HPC).
	RegionId  *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Scheduler []*SetSchedulerInfoRequestScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Repeated"`
	SlurmInfo []*SetSchedulerInfoRequestSlurmInfo `json:"SlurmInfo,omitempty" xml:"SlurmInfo,omitempty" type:"Repeated"`
}

func (s SetSchedulerInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSchedulerInfoRequest) GoString() string {
	return s.String()
}

func (s *SetSchedulerInfoRequest) SetClusterId(v string) *SetSchedulerInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *SetSchedulerInfoRequest) SetPbsInfo(v []*SetSchedulerInfoRequestPbsInfo) *SetSchedulerInfoRequest {
	s.PbsInfo = v
	return s
}

func (s *SetSchedulerInfoRequest) SetRegionId(v string) *SetSchedulerInfoRequest {
	s.RegionId = &v
	return s
}

func (s *SetSchedulerInfoRequest) SetScheduler(v []*SetSchedulerInfoRequestScheduler) *SetSchedulerInfoRequest {
	s.Scheduler = v
	return s
}

func (s *SetSchedulerInfoRequest) SetSlurmInfo(v []*SetSchedulerInfoRequestSlurmInfo) *SetSchedulerInfoRequest {
	s.SlurmInfo = v
	return s
}

type SetSchedulerInfoRequestPbsInfo struct {
	AclLimit []*SetSchedulerInfoRequestPbsInfoAclLimit `json:"AclLimit,omitempty" xml:"AclLimit,omitempty" type:"Repeated"`
	// The retention period of jobs. After the retention period is exceeded, job data is deleted. Unit: days.
	//
	// Valid values: 1 to 30
	//
	// Default value: 14
	JobHistoryDuration *int32                                         `json:"JobHistoryDuration,omitempty" xml:"JobHistoryDuration,omitempty"`
	ResourceLimit      []*SetSchedulerInfoRequestPbsInfoResourceLimit `json:"ResourceLimit,omitempty" xml:"ResourceLimit,omitempty" type:"Repeated"`
	// PbsInfo specifies the number of PBS schedulers that can be configured in the cluster. Valid values of N: 0 to 100.
	//
	// SchedInterval specifies the scheduling period. Unit: seconds.
	//
	// A scheduling period is the interval between two consecutive running jobs. If you set SchedInterval to 60, another job can be run 60 seconds after a job starts running.
	//
	// Default value: 60
	SchedInterval *int32 `json:"SchedInterval,omitempty" xml:"SchedInterval,omitempty"`
	// The maximum number of jobs that can be scheduled in the cluster. If the total number of running jobs and queuing jobs exceeds the value, no more jobs can be submitted. Default value: 20000.
	SchedMaxJobs *int32 `json:"SchedMaxJobs,omitempty" xml:"SchedMaxJobs,omitempty"`
	// The maximum number of queuing jobs that can be scheduled in the cluster. If the number of queuing jobs exceeds the value, no more jobs can be submitted. Default value: 10000.
	SchedMaxQueuedJobs *int32 `json:"SchedMaxQueuedJobs,omitempty" xml:"SchedMaxQueuedJobs,omitempty"`
}

func (s SetSchedulerInfoRequestPbsInfo) String() string {
	return tea.Prettify(s)
}

func (s SetSchedulerInfoRequestPbsInfo) GoString() string {
	return s.String()
}

func (s *SetSchedulerInfoRequestPbsInfo) SetAclLimit(v []*SetSchedulerInfoRequestPbsInfoAclLimit) *SetSchedulerInfoRequestPbsInfo {
	s.AclLimit = v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfo) SetJobHistoryDuration(v int32) *SetSchedulerInfoRequestPbsInfo {
	s.JobHistoryDuration = &v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfo) SetResourceLimit(v []*SetSchedulerInfoRequestPbsInfoResourceLimit) *SetSchedulerInfoRequestPbsInfo {
	s.ResourceLimit = v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfo) SetSchedInterval(v int32) *SetSchedulerInfoRequestPbsInfo {
	s.SchedInterval = &v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfo) SetSchedMaxJobs(v int32) *SetSchedulerInfoRequestPbsInfo {
	s.SchedMaxJobs = &v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfo) SetSchedMaxQueuedJobs(v int32) *SetSchedulerInfoRequestPbsInfo {
	s.SchedMaxQueuedJobs = &v
	return s
}

type SetSchedulerInfoRequestPbsInfoAclLimit struct {
	// The user that can use the queue. Separate multiple users with commas (`,`).
	//
	// If you specify users, you must specify the PbsInfo.N.AclLimit.N.Queue parameter.
	AclUsers *string `json:"AclUsers,omitempty" xml:"AclUsers,omitempty"`
	// AclLimit specifies the queue that has limits when it is used. Valid values of N: 0 to 100.
	//
	// If you set `PbsInfo.N.AclLimit.N.Queue` to `workq` and `PbsInfo.N.AclLimit.N.AclUsers` to `user1,user2`, workq can be used only by user1 and user2.
	Queue *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
}

func (s SetSchedulerInfoRequestPbsInfoAclLimit) String() string {
	return tea.Prettify(s)
}

func (s SetSchedulerInfoRequestPbsInfoAclLimit) GoString() string {
	return s.String()
}

func (s *SetSchedulerInfoRequestPbsInfoAclLimit) SetAclUsers(v string) *SetSchedulerInfoRequestPbsInfoAclLimit {
	s.AclUsers = &v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfoAclLimit) SetQueue(v string) *SetSchedulerInfoRequestPbsInfoAclLimit {
	s.Queue = &v
	return s
}

type SetSchedulerInfoRequestPbsInfoResourceLimit struct {
	// The maximum number of vCPUs that can be used for nodes in a queue.
	Cpus *int32 `json:"Cpus,omitempty" xml:"Cpus,omitempty"`
	// The maximum number of jobs that can be submitted to the cluster. If the total number of running jobs and queuing jobs exceeds the value, no more jobs can be submitted.
	MaxJobs *int32 `json:"MaxJobs,omitempty" xml:"MaxJobs,omitempty"`
	// The maximum memory resources that can be used in a queue. Units:
	//
	// *   gb
	// *   mb
	// *   kb
	Mem *string `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The maximum number of nodes that can be used in a queue.
	Nodes *int32 `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// PbsInfo specifies the number of PBS schedulers that can be configured in the cluster. Valid values of N: 0 to 100.
	//
	// ResourceLimit specifies the maximum number of queue resources that can be used. Valid values of N: 0 to 100.
	//
	// Queue specifies the name of the queue that is used to run jobs.
	//
	// If one of the User, Cpus, Nodes, and Mem parameters is set in ResourceLimit, you must specify the Queue parameter.
	Queue *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	// The name of the user that runs jobs.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s SetSchedulerInfoRequestPbsInfoResourceLimit) String() string {
	return tea.Prettify(s)
}

func (s SetSchedulerInfoRequestPbsInfoResourceLimit) GoString() string {
	return s.String()
}

func (s *SetSchedulerInfoRequestPbsInfoResourceLimit) SetCpus(v int32) *SetSchedulerInfoRequestPbsInfoResourceLimit {
	s.Cpus = &v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfoResourceLimit) SetMaxJobs(v int32) *SetSchedulerInfoRequestPbsInfoResourceLimit {
	s.MaxJobs = &v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfoResourceLimit) SetMem(v string) *SetSchedulerInfoRequestPbsInfoResourceLimit {
	s.Mem = &v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfoResourceLimit) SetNodes(v int32) *SetSchedulerInfoRequestPbsInfoResourceLimit {
	s.Nodes = &v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfoResourceLimit) SetQueue(v string) *SetSchedulerInfoRequestPbsInfoResourceLimit {
	s.Queue = &v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfoResourceLimit) SetUser(v string) *SetSchedulerInfoRequestPbsInfoResourceLimit {
	s.User = &v
	return s
}

type SetSchedulerInfoRequestScheduler struct {
	// The name of the scheduler. Valid values:
	//
	// *   pbs
	// *   pbs19
	// *   slurm
	// *   slurm19
	// *   slurm20
	//
	// >  If you set Scheduler.N.SchedName to pbs or pbs19, you must specify at least one of the PbsInfo.N.SchedInterval, PbsInfo.N.JobHistoryDuration, and PbsInfo.N.AclLimit parameters. If you set Scheduler.N.SchedName to slurm, slurm19, or slurm20, you must specify at least one of the SlurmInfo.N.SchedInterval and SlurmInfo.N.BackfillInterval parameters.
	SchedName *string `json:"SchedName,omitempty" xml:"SchedName,omitempty"`
}

func (s SetSchedulerInfoRequestScheduler) String() string {
	return tea.Prettify(s)
}

func (s SetSchedulerInfoRequestScheduler) GoString() string {
	return s.String()
}

func (s *SetSchedulerInfoRequestScheduler) SetSchedName(v string) *SetSchedulerInfoRequestScheduler {
	s.SchedName = &v
	return s
}

type SetSchedulerInfoRequestSlurmInfo struct {
	// The backfill scheduling period. Unit: seconds.
	//
	// Default value: 60
	BackfillInterval *int32 `json:"BackfillInterval,omitempty" xml:"BackfillInterval,omitempty"`
	// SlurmInfo specifies the number of Slurm schedulers that can be configured in the cluster. Valid values of N: 0 to 100.
	//
	// SchedInterval specifies the scheduling period. Unit: seconds.
	//
	// Default value: 60
	SchedInterval *int32 `json:"SchedInterval,omitempty" xml:"SchedInterval,omitempty"`
}

func (s SetSchedulerInfoRequestSlurmInfo) String() string {
	return tea.Prettify(s)
}

func (s SetSchedulerInfoRequestSlurmInfo) GoString() string {
	return s.String()
}

func (s *SetSchedulerInfoRequestSlurmInfo) SetBackfillInterval(v int32) *SetSchedulerInfoRequestSlurmInfo {
	s.BackfillInterval = &v
	return s
}

func (s *SetSchedulerInfoRequestSlurmInfo) SetSchedInterval(v int32) *SetSchedulerInfoRequestSlurmInfo {
	s.SchedInterval = &v
	return s
}

type SetSchedulerInfoResponseBody struct {
	// The response message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetSchedulerInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetSchedulerInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SetSchedulerInfoResponseBody) SetMessage(v string) *SetSchedulerInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SetSchedulerInfoResponseBody) SetRequestId(v string) *SetSchedulerInfoResponseBody {
	s.RequestId = &v
	return s
}

type SetSchedulerInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetSchedulerInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetSchedulerInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SetSchedulerInfoResponse) GoString() string {
	return s.String()
}

func (s *SetSchedulerInfoResponse) SetHeaders(v map[string]*string) *SetSchedulerInfoResponse {
	s.Headers = v
	return s
}

func (s *SetSchedulerInfoResponse) SetStatusCode(v int32) *SetSchedulerInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSchedulerInfoResponse) SetBody(v *SetSchedulerInfoResponseBody) *SetSchedulerInfoResponse {
	s.Body = v
	return s
}

type StartClusterRequest struct {
	// The ID of the cluster that you want to start.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s StartClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s StartClusterRequest) GoString() string {
	return s.String()
}

func (s *StartClusterRequest) SetClusterId(v string) *StartClusterRequest {
	s.ClusterId = &v
	return s
}

type StartClusterResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StartClusterResponseBody) SetRequestId(v string) *StartClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartClusterResponseBody) SetTaskId(v string) *StartClusterResponseBody {
	s.TaskId = &v
	return s
}

type StartClusterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s StartClusterResponse) GoString() string {
	return s.String()
}

func (s *StartClusterResponse) SetHeaders(v map[string]*string) *StartClusterResponse {
	s.Headers = v
	return s
}

func (s *StartClusterResponse) SetStatusCode(v int32) *StartClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *StartClusterResponse) SetBody(v *StartClusterResponseBody) *StartClusterResponse {
	s.Body = v
	return s
}

type StartGWSInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartGWSInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartGWSInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartGWSInstanceRequest) SetInstanceId(v string) *StartGWSInstanceRequest {
	s.InstanceId = &v
	return s
}

type StartGWSInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartGWSInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartGWSInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartGWSInstanceResponseBody) SetRequestId(v string) *StartGWSInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartGWSInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartGWSInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartGWSInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartGWSInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartGWSInstanceResponse) SetHeaders(v map[string]*string) *StartGWSInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartGWSInstanceResponse) SetStatusCode(v int32) *StartGWSInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartGWSInstanceResponse) SetBody(v *StartGWSInstanceResponseBody) *StartGWSInstanceResponse {
	s.Body = v
	return s
}

type StartNodesRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string                      `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Instance  []*StartNodesRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
	// The role of the node. Valid values:
	//
	// *   Manager: management node
	// *   Login: logon node
	// *   Compute: compute node
	//
	// Default value: Compute
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s StartNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s StartNodesRequest) GoString() string {
	return s.String()
}

func (s *StartNodesRequest) SetClusterId(v string) *StartNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *StartNodesRequest) SetInstance(v []*StartNodesRequestInstance) *StartNodesRequest {
	s.Instance = v
	return s
}

func (s *StartNodesRequest) SetRole(v string) *StartNodesRequest {
	s.Role = &v
	return s
}

type StartNodesRequestInstance struct {
	// The ID of the Nth node. Valid values of N: 1 to 100.
	//
	// Make sure that the node is in the Stopped state. You can call the [ListNodes](~~87161~~) operation to query the status of the node.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s StartNodesRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s StartNodesRequestInstance) GoString() string {
	return s.String()
}

func (s *StartNodesRequestInstance) SetId(v string) *StartNodesRequestInstance {
	s.Id = &v
	return s
}

type StartNodesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartNodesResponseBody) GoString() string {
	return s.String()
}

func (s *StartNodesResponseBody) SetRequestId(v string) *StartNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartNodesResponseBody) SetTaskId(v string) *StartNodesResponseBody {
	s.TaskId = &v
	return s
}

type StartNodesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s StartNodesResponse) GoString() string {
	return s.String()
}

func (s *StartNodesResponse) SetHeaders(v map[string]*string) *StartNodesResponse {
	s.Headers = v
	return s
}

func (s *StartNodesResponse) SetStatusCode(v int32) *StartNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *StartNodesResponse) SetBody(v *StartNodesResponseBody) *StartNodesResponse {
	s.Body = v
	return s
}

type StartVisualServiceRequest struct {
	// A public IP address of logon nodes in the cluster.
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The fixed port. Set the value to 12016
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s StartVisualServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartVisualServiceRequest) GoString() string {
	return s.String()
}

func (s *StartVisualServiceRequest) SetCidrIp(v string) *StartVisualServiceRequest {
	s.CidrIp = &v
	return s
}

func (s *StartVisualServiceRequest) SetClusterId(v string) *StartVisualServiceRequest {
	s.ClusterId = &v
	return s
}

func (s *StartVisualServiceRequest) SetPort(v int32) *StartVisualServiceRequest {
	s.Port = &v
	return s
}

type StartVisualServiceResponseBody struct {
	// The status of the VNC Remote visualization service. Valid values:
	//
	// *   Service started
	// *   Service stopped
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartVisualServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartVisualServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StartVisualServiceResponseBody) SetMessage(v string) *StartVisualServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StartVisualServiceResponseBody) SetRequestId(v string) *StartVisualServiceResponseBody {
	s.RequestId = &v
	return s
}

type StartVisualServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartVisualServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartVisualServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartVisualServiceResponse) GoString() string {
	return s.String()
}

func (s *StartVisualServiceResponse) SetHeaders(v map[string]*string) *StartVisualServiceResponse {
	s.Headers = v
	return s
}

func (s *StartVisualServiceResponse) SetStatusCode(v int32) *StartVisualServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartVisualServiceResponse) SetBody(v *StartVisualServiceResponseBody) *StartVisualServiceResponse {
	s.Body = v
	return s
}

type StopClusterRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s StopClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s StopClusterRequest) GoString() string {
	return s.String()
}

func (s *StopClusterRequest) SetClusterId(v string) *StopClusterRequest {
	s.ClusterId = &v
	return s
}

type StopClusterResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StopClusterResponseBody) SetRequestId(v string) *StopClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopClusterResponseBody) SetTaskId(v string) *StopClusterResponseBody {
	s.TaskId = &v
	return s
}

type StopClusterResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s StopClusterResponse) GoString() string {
	return s.String()
}

func (s *StopClusterResponse) SetHeaders(v map[string]*string) *StopClusterResponse {
	s.Headers = v
	return s
}

func (s *StopClusterResponse) SetStatusCode(v int32) *StopClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *StopClusterResponse) SetBody(v *StopClusterResponseBody) *StopClusterResponse {
	s.Body = v
	return s
}

type StopGWSInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopGWSInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopGWSInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopGWSInstanceRequest) SetInstanceId(v string) *StopGWSInstanceRequest {
	s.InstanceId = &v
	return s
}

type StopGWSInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopGWSInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopGWSInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopGWSInstanceResponseBody) SetRequestId(v string) *StopGWSInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StopGWSInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopGWSInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopGWSInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopGWSInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopGWSInstanceResponse) SetHeaders(v map[string]*string) *StopGWSInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopGWSInstanceResponse) SetStatusCode(v int32) *StopGWSInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopGWSInstanceResponse) SetBody(v *StopGWSInstanceResponseBody) *StopGWSInstanceResponse {
	s.Body = v
	return s
}

type StopJobsRequest struct {
	// Specifies whether to use an asynchronous link to stop the job.
	//
	// Default value: false
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The list of jobs that you want to stop. Maximum number of jobs: 100. Minimum number of jobs: 1.
	//
	// Format: `[{"Id": "0.sched****"},{"Id": "1.sched****"}]`. Separate multiple jobs with commas (,).
	//
	// You can call the [ListJobs](~~87251~~) operation to query the job ID.
	//
	// >  You can stop only jobs that are in the RUNNING or QUEUED state.
	Jobs *string `json:"Jobs,omitempty" xml:"Jobs,omitempty"`
}

func (s StopJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s StopJobsRequest) GoString() string {
	return s.String()
}

func (s *StopJobsRequest) SetAsync(v bool) *StopJobsRequest {
	s.Async = &v
	return s
}

func (s *StopJobsRequest) SetClusterId(v string) *StopJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *StopJobsRequest) SetJobs(v string) *StopJobsRequest {
	s.Jobs = &v
	return s
}

type StopJobsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopJobsResponseBody) GoString() string {
	return s.String()
}

func (s *StopJobsResponseBody) SetRequestId(v string) *StopJobsResponseBody {
	s.RequestId = &v
	return s
}

type StopJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s StopJobsResponse) GoString() string {
	return s.String()
}

func (s *StopJobsResponse) SetHeaders(v map[string]*string) *StopJobsResponse {
	s.Headers = v
	return s
}

func (s *StopJobsResponse) SetStatusCode(v int32) *StopJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *StopJobsResponse) SetBody(v *StopJobsResponseBody) *StopJobsResponse {
	s.Body = v
	return s
}

type StopNodesRequest struct {
	// The ID of the E-HPC cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string                     `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Instance  []*StopNodesRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
	// The role of the node. Valid values:
	//
	// *   Manager: management node
	// *   Login: logon node
	// *   Compute: compute node
	//
	// Default value: Compute
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s StopNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s StopNodesRequest) GoString() string {
	return s.String()
}

func (s *StopNodesRequest) SetClusterId(v string) *StopNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *StopNodesRequest) SetInstance(v []*StopNodesRequestInstance) *StopNodesRequest {
	s.Instance = v
	return s
}

func (s *StopNodesRequest) SetRole(v string) *StopNodesRequest {
	s.Role = &v
	return s
}

type StopNodesRequestInstance struct {
	// The ID of the Nth node that you want to stop. Valid values of N: 1 to 100
	//
	// You can call the [ListNodes](~~87161~~) operation to query the IDs of the compute nodes.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s StopNodesRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s StopNodesRequestInstance) GoString() string {
	return s.String()
}

func (s *StopNodesRequestInstance) SetId(v string) *StopNodesRequestInstance {
	s.Id = &v
	return s
}

type StopNodesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopNodesResponseBody) GoString() string {
	return s.String()
}

func (s *StopNodesResponseBody) SetRequestId(v string) *StopNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopNodesResponseBody) SetTaskId(v string) *StopNodesResponseBody {
	s.TaskId = &v
	return s
}

type StopNodesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s StopNodesResponse) GoString() string {
	return s.String()
}

func (s *StopNodesResponse) SetHeaders(v map[string]*string) *StopNodesResponse {
	s.Headers = v
	return s
}

func (s *StopNodesResponse) SetStatusCode(v int32) *StopNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *StopNodesResponse) SetBody(v *StopNodesResponseBody) *StopNodesResponse {
	s.Body = v
	return s
}

type StopVisualServiceRequest struct {
	// A public IP address of login nodes in the cluster.
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The fixed port. Set the value to 12016.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s StopVisualServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopVisualServiceRequest) GoString() string {
	return s.String()
}

func (s *StopVisualServiceRequest) SetCidrIp(v string) *StopVisualServiceRequest {
	s.CidrIp = &v
	return s
}

func (s *StopVisualServiceRequest) SetClusterId(v string) *StopVisualServiceRequest {
	s.ClusterId = &v
	return s
}

func (s *StopVisualServiceRequest) SetPort(v int32) *StopVisualServiceRequest {
	s.Port = &v
	return s
}

type StopVisualServiceResponseBody struct {
	// The status of the VNC Remote Service. Valid values:
	//
	// *   Service started
	// *   Service stopped
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopVisualServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopVisualServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StopVisualServiceResponseBody) SetMessage(v string) *StopVisualServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StopVisualServiceResponseBody) SetRequestId(v string) *StopVisualServiceResponseBody {
	s.RequestId = &v
	return s
}

type StopVisualServiceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopVisualServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopVisualServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopVisualServiceResponse) GoString() string {
	return s.String()
}

func (s *StopVisualServiceResponse) SetHeaders(v map[string]*string) *StopVisualServiceResponse {
	s.Headers = v
	return s
}

func (s *StopVisualServiceResponse) SetStatusCode(v int32) *StopVisualServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopVisualServiceResponse) SetBody(v *StopVisualServiceResponseBody) *StopVisualServiceResponse {
	s.Body = v
	return s
}

type SubmitJobRequest struct {
	JobRetry *SubmitJobRequestJobRetry `json:"JobRetry,omitempty" xml:"JobRetry,omitempty" type:"Struct"`
	// The job array.
	//
	// Format: X-Y:Z. The minimum index value X is the first index. The maximum index value Y is the last index. Z is the step size. For example, 2-7:2 indicates that three jobs need to be run and their index values are 2, 4, and 6.
	ArrayRequest *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	// Specifies whether to use an asynchronous link to submit the job.
	//
	// Default value: false
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The maximum running time of the job. Valid formats:
	//
	// *   hh:mm:ss
	// *   mm:ss
	// *   ss
	//
	// We recommend that you use the hh:mm:ss format. If the maximum running time is 12 hours, set the value to 12:00:00.
	ClockTime *string `json:"ClockTime,omitempty" xml:"ClockTime,omitempty"`
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The command that is used to run the job.
	CommandLine *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	// The ID of the containerized application. If you want to use a container application, you must specify its ID.
	//
	// You can call the [ListContainerApps](~~87333~~) operation to query the container application ID.
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The number of CPU cores required by a single compute node.
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The maximum GPU usage required by a single compute node.
	//
	// The parameter takes effect only when the cluster uses PBS and a compute node is a GPU-accelerated instance.
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The URL of the job files that are uploaded to an Object Storage Service (OSS) bucket.
	InputFileUrl *string `json:"InputFileUrl,omitempty" xml:"InputFileUrl,omitempty"`
	// The name of the queue in which the job is run.
	//
	// You can call the [ListQueues](~~92176~~) operation to query the name of the queue.
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	// The maximum memory usage required by a single compute node. Unit: GB, MB, or KB. The unit is case-insensitive.
	Mem *string `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The name of the job. The name must be 6 to 30 characters in length and start with a letter. It can contain letters, digits, and periods (.).
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of compute nodes required to run the job.
	//
	// >  If the parameter is not specified, the Task, Thread, Mem, and Gpu parameters become invalid.
	Node *int32 `json:"Node,omitempty" xml:"Node,omitempty"`
	// The path that is used to run the job.
	PackagePath *string `json:"PackagePath,omitempty" xml:"PackagePath,omitempty"`
	// The command to perform on the job after the job is submitted.
	PostCmdLine *string `json:"PostCmdLine,omitempty" xml:"PostCmdLine,omitempty"`
	// The priority of the job. Valid values: 0 to 9. A large value indicates a high priority.
	//
	// Default value: 0
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Specifies whether the job can be rerun. Valid values:
	//
	// *   true: The job can be rerun.
	// *   false: The job cannot be rerun.
	ReRunable *bool `json:"ReRunable,omitempty" xml:"ReRunable,omitempty"`
	// The name of the user that runs the job.
	//
	// You can call the [ListUsers](~~188572~~) operation to query the users of the cluster.
	RunasUser *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	// The user password.
	RunasUserPassword *string `json:"RunasUserPassword,omitempty" xml:"RunasUserPassword,omitempty"`
	// The output file path of stderr.
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" xml:"StderrRedirectPath,omitempty"`
	// The output file path of stdout.
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" xml:"StdoutRedirectPath,omitempty"`
	// The number of processes created for a single compute node.
	//
	// The parameter is applicable to Message Passing Interface (MPI) jobs.
	Task *int32 `json:"Task,omitempty" xml:"Task,omitempty"`
	// The number of threads created for a single compute node.
	//
	// The parameter is applicable to OpenMP jobs.
	Thread *int32 `json:"Thread,omitempty" xml:"Thread,omitempty"`
	// The command for file decompression. The command that is used to decompress the job files downloaded from an OSS bucket. Valid values:
	//
	// *   tar xzf: Decompresses GZIP files.
	// *   tar xf: Decompresses TAR files.
	// *   unzip: Decompresses ZIP files.
	UnzipCmd *string `json:"UnzipCmd,omitempty" xml:"UnzipCmd,omitempty"`
	// The runtime variables passed to the job. They can be accessed by using environment variables in the executable file.
	Variables *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s SubmitJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitJobRequest) SetJobRetry(v *SubmitJobRequestJobRetry) *SubmitJobRequest {
	s.JobRetry = v
	return s
}

func (s *SubmitJobRequest) SetArrayRequest(v string) *SubmitJobRequest {
	s.ArrayRequest = &v
	return s
}

func (s *SubmitJobRequest) SetAsync(v bool) *SubmitJobRequest {
	s.Async = &v
	return s
}

func (s *SubmitJobRequest) SetClockTime(v string) *SubmitJobRequest {
	s.ClockTime = &v
	return s
}

func (s *SubmitJobRequest) SetClusterId(v string) *SubmitJobRequest {
	s.ClusterId = &v
	return s
}

func (s *SubmitJobRequest) SetCommandLine(v string) *SubmitJobRequest {
	s.CommandLine = &v
	return s
}

func (s *SubmitJobRequest) SetContainerId(v string) *SubmitJobRequest {
	s.ContainerId = &v
	return s
}

func (s *SubmitJobRequest) SetCpu(v int32) *SubmitJobRequest {
	s.Cpu = &v
	return s
}

func (s *SubmitJobRequest) SetGpu(v int32) *SubmitJobRequest {
	s.Gpu = &v
	return s
}

func (s *SubmitJobRequest) SetInputFileUrl(v string) *SubmitJobRequest {
	s.InputFileUrl = &v
	return s
}

func (s *SubmitJobRequest) SetJobQueue(v string) *SubmitJobRequest {
	s.JobQueue = &v
	return s
}

func (s *SubmitJobRequest) SetMem(v string) *SubmitJobRequest {
	s.Mem = &v
	return s
}

func (s *SubmitJobRequest) SetName(v string) *SubmitJobRequest {
	s.Name = &v
	return s
}

func (s *SubmitJobRequest) SetNode(v int32) *SubmitJobRequest {
	s.Node = &v
	return s
}

func (s *SubmitJobRequest) SetPackagePath(v string) *SubmitJobRequest {
	s.PackagePath = &v
	return s
}

func (s *SubmitJobRequest) SetPostCmdLine(v string) *SubmitJobRequest {
	s.PostCmdLine = &v
	return s
}

func (s *SubmitJobRequest) SetPriority(v int32) *SubmitJobRequest {
	s.Priority = &v
	return s
}

func (s *SubmitJobRequest) SetReRunable(v bool) *SubmitJobRequest {
	s.ReRunable = &v
	return s
}

func (s *SubmitJobRequest) SetRunasUser(v string) *SubmitJobRequest {
	s.RunasUser = &v
	return s
}

func (s *SubmitJobRequest) SetRunasUserPassword(v string) *SubmitJobRequest {
	s.RunasUserPassword = &v
	return s
}

func (s *SubmitJobRequest) SetStderrRedirectPath(v string) *SubmitJobRequest {
	s.StderrRedirectPath = &v
	return s
}

func (s *SubmitJobRequest) SetStdoutRedirectPath(v string) *SubmitJobRequest {
	s.StdoutRedirectPath = &v
	return s
}

func (s *SubmitJobRequest) SetTask(v int32) *SubmitJobRequest {
	s.Task = &v
	return s
}

func (s *SubmitJobRequest) SetThread(v int32) *SubmitJobRequest {
	s.Thread = &v
	return s
}

func (s *SubmitJobRequest) SetUnzipCmd(v string) *SubmitJobRequest {
	s.UnzipCmd = &v
	return s
}

func (s *SubmitJobRequest) SetVariables(v string) *SubmitJobRequest {
	s.Variables = &v
	return s
}

type SubmitJobRequestJobRetry struct {
	Count      *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	OnExitCode *int32 `json:"OnExitCode,omitempty" xml:"OnExitCode,omitempty"`
	Priority   *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s SubmitJobRequestJobRetry) String() string {
	return tea.Prettify(s)
}

func (s SubmitJobRequestJobRetry) GoString() string {
	return s.String()
}

func (s *SubmitJobRequestJobRetry) SetCount(v int32) *SubmitJobRequestJobRetry {
	s.Count = &v
	return s
}

func (s *SubmitJobRequestJobRetry) SetOnExitCode(v int32) *SubmitJobRequestJobRetry {
	s.OnExitCode = &v
	return s
}

func (s *SubmitJobRequestJobRetry) SetPriority(v int32) *SubmitJobRequestJobRetry {
	s.Priority = &v
	return s
}

type SubmitJobResponseBody struct {
	// The ID of the job.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitJobResponseBody) SetJobId(v string) *SubmitJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitJobResponseBody) SetRequestId(v string) *SubmitJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitJobResponse) SetHeaders(v map[string]*string) *SubmitJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitJobResponse) SetStatusCode(v int32) *SubmitJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitJobResponse) SetBody(v *SubmitJobResponseBody) *SubmitJobResponse {
	s.Body = v
	return s
}

type SummaryImagesRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the image. Set the value to singularity.
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
}

func (s SummaryImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s SummaryImagesRequest) GoString() string {
	return s.String()
}

func (s *SummaryImagesRequest) SetClusterId(v string) *SummaryImagesRequest {
	s.ClusterId = &v
	return s
}

func (s *SummaryImagesRequest) SetContainerType(v string) *SummaryImagesRequest {
	s.ContainerType = &v
	return s
}

type SummaryImagesResponseBody struct {
	// The names of all images in the cluster.
	ImagesName *string `json:"ImagesName,omitempty" xml:"ImagesName,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SummaryImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SummaryImagesResponseBody) GoString() string {
	return s.String()
}

func (s *SummaryImagesResponseBody) SetImagesName(v string) *SummaryImagesResponseBody {
	s.ImagesName = &v
	return s
}

func (s *SummaryImagesResponseBody) SetRequestId(v string) *SummaryImagesResponseBody {
	s.RequestId = &v
	return s
}

type SummaryImagesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SummaryImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SummaryImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s SummaryImagesResponse) GoString() string {
	return s.String()
}

func (s *SummaryImagesResponse) SetHeaders(v map[string]*string) *SummaryImagesResponse {
	s.Headers = v
	return s
}

func (s *SummaryImagesResponse) SetStatusCode(v int32) *SummaryImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *SummaryImagesResponse) SetBody(v *SummaryImagesResponseBody) *SummaryImagesResponse {
	s.Body = v
	return s
}

type SummaryImagesInfoRequest struct {
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The container type of the image. Set the value to singularity.
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
	// The name of the image. You can call the [SummaryImages](~~440783~~) operation to query the names of all images in a cluster.
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
}

func (s SummaryImagesInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SummaryImagesInfoRequest) GoString() string {
	return s.String()
}

func (s *SummaryImagesInfoRequest) SetClusterId(v string) *SummaryImagesInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *SummaryImagesInfoRequest) SetContainerType(v string) *SummaryImagesInfoRequest {
	s.ContainerType = &v
	return s
}

func (s *SummaryImagesInfoRequest) SetImageName(v string) *SummaryImagesInfoRequest {
	s.ImageName = &v
	return s
}

type SummaryImagesInfoResponseBody struct {
	// The detailed information about the image.
	ImagesInfo *string `json:"ImagesInfo,omitempty" xml:"ImagesInfo,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SummaryImagesInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SummaryImagesInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SummaryImagesInfoResponseBody) SetImagesInfo(v string) *SummaryImagesInfoResponseBody {
	s.ImagesInfo = &v
	return s
}

func (s *SummaryImagesInfoResponseBody) SetRequestId(v string) *SummaryImagesInfoResponseBody {
	s.RequestId = &v
	return s
}

type SummaryImagesInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SummaryImagesInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SummaryImagesInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SummaryImagesInfoResponse) GoString() string {
	return s.String()
}

func (s *SummaryImagesInfoResponse) SetHeaders(v map[string]*string) *SummaryImagesInfoResponse {
	s.Headers = v
	return s
}

func (s *SummaryImagesInfoResponse) SetStatusCode(v int32) *SummaryImagesInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SummaryImagesInfoResponse) SetBody(v *SummaryImagesInfoResponseBody) *SummaryImagesInfoResponse {
	s.Body = v
	return s
}

type SyncUsersRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the list of E-HPC clusters.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the region where the cluster resides.
	//
	// You can call the [ListRegions](~~188593~~) operation to query the list of regions where E-HPC is available.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SyncUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncUsersRequest) GoString() string {
	return s.String()
}

func (s *SyncUsersRequest) SetClusterId(v string) *SyncUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *SyncUsersRequest) SetRegionId(v string) *SyncUsersRequest {
	s.RegionId = &v
	return s
}

type SyncUsersResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncUsersResponseBody) GoString() string {
	return s.String()
}

func (s *SyncUsersResponseBody) SetRequestId(v string) *SyncUsersResponseBody {
	s.RequestId = &v
	return s
}

type SyncUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SyncUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncUsersResponse) GoString() string {
	return s.String()
}

func (s *SyncUsersResponse) SetHeaders(v map[string]*string) *SyncUsersResponse {
	s.Headers = v
	return s
}

func (s *SyncUsersResponse) SetStatusCode(v int32) *SyncUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncUsersResponse) SetBody(v *SyncUsersResponseBody) *SyncUsersResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The ID of the region to which the resource belongs.
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Set the value to cluster, which indicates E-HPC clusters.
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The key of the tag. The tag key cannot be an empty string. It can be up to 128 characters in length and cannot contain `http://` or `https://`. It must not start with `acs:` or `aliyun`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. The tag key cannot be an empty string. It can be up to 128 characters in length and cannot contain `http://` or `https://`. It must not start with `acs:` or `aliyun`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnTagResourcesRequest struct {
	// Specifies whether to remove all tags from the resource. This parameter is valid only when the TagKey.N parameter is not specified. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The ID of the region to which the resource belongs.
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Set the value to cluster, which indicates E-HPC clusters.
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UnTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequest) SetAll(v bool) *UnTagResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagResourcesRequest) SetRegionId(v string) *UnTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceId(v []*string) *UnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagResourcesRequest) SetResourceType(v string) *UnTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagResourcesRequest) SetTagKey(v []*string) *UnTagResourcesRequest {
	s.TagKey = v
	return s
}

type UnTagResourcesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponseBody) SetRequestId(v string) *UnTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UnTagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponse) SetHeaders(v map[string]*string) *UnTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UnTagResourcesResponse) SetStatusCode(v int32) *UnTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UnTagResourcesResponse) SetBody(v *UnTagResourcesResponseBody) *UnTagResourcesResponse {
	s.Body = v
	return s
}

type UninstallSoftwareRequest struct {
	// The name of the software that you want to uninstall.
	//
	// You can call the [ListInstalledSoftware](~~188591~~) operation to query the software that is installed in the cluster.
	Application *string `json:"Application,omitempty" xml:"Application,omitempty"`
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UninstallSoftwareRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallSoftwareRequest) GoString() string {
	return s.String()
}

func (s *UninstallSoftwareRequest) SetApplication(v string) *UninstallSoftwareRequest {
	s.Application = &v
	return s
}

func (s *UninstallSoftwareRequest) SetClusterId(v string) *UninstallSoftwareRequest {
	s.ClusterId = &v
	return s
}

type UninstallSoftwareResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UninstallSoftwareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallSoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallSoftwareResponseBody) SetRequestId(v string) *UninstallSoftwareResponseBody {
	s.RequestId = &v
	return s
}

type UninstallSoftwareResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UninstallSoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UninstallSoftwareResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallSoftwareResponse) GoString() string {
	return s.String()
}

func (s *UninstallSoftwareResponse) SetHeaders(v map[string]*string) *UninstallSoftwareResponse {
	s.Headers = v
	return s
}

func (s *UninstallSoftwareResponse) SetStatusCode(v int32) *UninstallSoftwareResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallSoftwareResponse) SetBody(v *UninstallSoftwareResponseBody) *UninstallSoftwareResponse {
	s.Body = v
	return s
}

type UpdateClusterVolumesRequest struct {
	AdditionalVolumes []*UpdateClusterVolumesRequestAdditionalVolumes `json:"AdditionalVolumes,omitempty" xml:"AdditionalVolumes,omitempty" type:"Repeated"`
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UpdateClusterVolumesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterVolumesRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterVolumesRequest) SetAdditionalVolumes(v []*UpdateClusterVolumesRequestAdditionalVolumes) *UpdateClusterVolumesRequest {
	s.AdditionalVolumes = v
	return s
}

func (s *UpdateClusterVolumesRequest) SetClusterId(v string) *UpdateClusterVolumesRequest {
	s.ClusterId = &v
	return s
}

type UpdateClusterVolumesRequestAdditionalVolumes struct {
	// The queue name of the nth attached mounted filesystem.
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	// The on-premises mount directory for the nth additional mounted file system.
	LocalDirectory *string `json:"LocalDirectory,omitempty" xml:"LocalDirectory,omitempty"`
	// The storage location of the nth attached mounted file system. Valid values:
	//
	// *   OnPremise: hybrid cloud cluster
	// *   PublicCloud: public cloud cluster
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The remote directory to be mounted by the nth additional mounted file system.
	RemoteDirectory *string                                              `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	Roles           []*UpdateClusterVolumesRequestAdditionalVolumesRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// The ID of the nth additional mounted file system.
	VolumeId *string `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	// The domain name of the mount target for the nth additional mounted file system.
	VolumeMountpoint *string `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	// The protocol type of the nth additional mounted file system. Valid values:
	//
	// *   NFS
	// *   SMB
	VolumeProtocol *string `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	// The type of the nth additional mounted file system. Currently, only NAS is supported.
	//
	// Valid values of N: 1 to 10.
	VolumeType *string `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
}

func (s UpdateClusterVolumesRequestAdditionalVolumes) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterVolumesRequestAdditionalVolumes) GoString() string {
	return s.String()
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetJobQueue(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.JobQueue = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetLocalDirectory(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.LocalDirectory = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetLocation(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.Location = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetRemoteDirectory(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.RemoteDirectory = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetRoles(v []*UpdateClusterVolumesRequestAdditionalVolumesRoles) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.Roles = v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetVolumeId(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.VolumeId = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetVolumeMountpoint(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.VolumeMountpoint = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetVolumeProtocol(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.VolumeProtocol = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetVolumeType(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.VolumeType = &v
	return s
}

type UpdateClusterVolumesRequestAdditionalVolumesRoles struct {
	// The node type on which the nth additional mounted file system is mounted. Valid values:
	//
	// *   Manager: management node
	// *   Login: logon node
	// *   Compute: compute node
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateClusterVolumesRequestAdditionalVolumesRoles) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterVolumesRequestAdditionalVolumesRoles) GoString() string {
	return s.String()
}

func (s *UpdateClusterVolumesRequestAdditionalVolumesRoles) SetName(v string) *UpdateClusterVolumesRequestAdditionalVolumesRoles {
	s.Name = &v
	return s
}

type UpdateClusterVolumesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateClusterVolumesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterVolumesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterVolumesResponseBody) SetRequestId(v string) *UpdateClusterVolumesResponseBody {
	s.RequestId = &v
	return s
}

type UpdateClusterVolumesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateClusterVolumesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateClusterVolumesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterVolumesResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterVolumesResponse) SetHeaders(v map[string]*string) *UpdateClusterVolumesResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterVolumesResponse) SetStatusCode(v int32) *UpdateClusterVolumesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClusterVolumesResponse) SetBody(v *UpdateClusterVolumesResponseBody) *UpdateClusterVolumesResponse {
	s.Body = v
	return s
}

type UpdateQueueConfigRequest struct {
	// The ID of the cluster.
	//
	// You can call the [ListClusters](~~87116~~) operation to query the cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The instance type of the node.
	//
	// You can call the [ListPreferredEcsTypes](~~188592~~) operation to query the recommended instance types.
	ComputeInstanceType *string `json:"ComputeInstanceType,omitempty" xml:"ComputeInstanceType,omitempty"`
	// The name of the queue.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the resource group.
	//
	// You can call the [ListResourceGroups](~~158855~~) operation to query the IDs of resource groups.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UpdateQueueConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQueueConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateQueueConfigRequest) SetClusterId(v string) *UpdateQueueConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateQueueConfigRequest) SetComputeInstanceType(v string) *UpdateQueueConfigRequest {
	s.ComputeInstanceType = &v
	return s
}

func (s *UpdateQueueConfigRequest) SetQueueName(v string) *UpdateQueueConfigRequest {
	s.QueueName = &v
	return s
}

func (s *UpdateQueueConfigRequest) SetResourceGroupId(v string) *UpdateQueueConfigRequest {
	s.ResourceGroupId = &v
	return s
}

type UpdateQueueConfigResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateQueueConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateQueueConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQueueConfigResponseBody) SetRequestId(v string) *UpdateQueueConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateQueueConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateQueueConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateQueueConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateQueueConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateQueueConfigResponse) SetHeaders(v map[string]*string) *UpdateQueueConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateQueueConfigResponse) SetStatusCode(v int32) *UpdateQueueConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQueueConfigResponse) SetBody(v *UpdateQueueConfigResponseBody) *UpdateQueueConfigResponse {
	s.Body = v
	return s
}

type UpgradeClientRequest struct {
	// The version to which the client will be upgraded. By default, the client is upgraded to the latest version. You can call the [ListCurrentClientVersion](~~87223~~) operation to query the latest version number of the Elastic High Performance Computing (E-HPC) client.
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UpgradeClientRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClientRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClientRequest) SetClientVersion(v string) *UpgradeClientRequest {
	s.ClientVersion = &v
	return s
}

func (s *UpgradeClientRequest) SetClusterId(v string) *UpgradeClientRequest {
	s.ClusterId = &v
	return s
}

type UpgradeClientResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeClientResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClientResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeClientResponseBody) SetRequestId(v string) *UpgradeClientResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeClientResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeClientResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeClientResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClientResponse) GoString() string {
	return s.String()
}

func (s *UpgradeClientResponse) SetHeaders(v map[string]*string) *UpgradeClientResponse {
	s.Headers = v
	return s
}

func (s *UpgradeClientResponse) SetStatusCode(v int32) *UpgradeClientResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeClientResponse) SetBody(v *UpgradeClientResponseBody) *UpgradeClientResponse {
	s.Body = v
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ehpc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

/**
 * If you select an image for a new containerized application, the image is pulled from Docker Hub by default. However, the version of the image may not be up to date. You can call the [PullImage](~~159052~~) operation to pull the latest image.
 *
 * @param request AddContainerAppRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddContainerAppResponse
 */
func (client *Client) AddContainerAppWithOptions(request *AddContainerAppRequest, runtime *util.RuntimeOptions) (_result *AddContainerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddContainerApp"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddContainerAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you select an image for a new containerized application, the image is pulled from Docker Hub by default. However, the version of the image may not be up to date. You can call the [PullImage](~~159052~~) operation to pull the latest image.
 *
 * @param request AddContainerAppRequest
 * @return AddContainerAppResponse
 */
func (client *Client) AddContainerApp(request *AddContainerAppRequest) (_result *AddContainerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddContainerAppResponse{}
	_body, _err := client.AddContainerAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   The compute nodes to be added are in the Stopped state.
 * *   After the compute nodes are added to the cluster, the operating systems of the nodes are replaced with the operating system specified by the ImageId parameter.
 * *   The hosts of the compute nodes must be different from those of the existing compute nodes in the cluster. Otherwise, the add operation fails.
 *
 * @param request AddExistedNodesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddExistedNodesResponse
 */
func (client *Client) AddExistedNodesWithOptions(request *AddExistedNodesRequest, runtime *util.RuntimeOptions) (_result *AddExistedNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddExistedNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddExistedNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   The compute nodes to be added are in the Stopped state.
 * *   After the compute nodes are added to the cluster, the operating systems of the nodes are replaced with the operating system specified by the ImageId parameter.
 * *   The hosts of the compute nodes must be different from those of the existing compute nodes in the cluster. Otherwise, the add operation fails.
 *
 * @param request AddExistedNodesRequest
 * @return AddExistedNodesResponse
 */
func (client *Client) AddExistedNodes(request *AddExistedNodesRequest) (_result *AddExistedNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddExistedNodesResponse{}
	_body, _err := client.AddExistedNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLocalNodesWithOptions(request *AddLocalNodesRequest, runtime *util.RuntimeOptions) (_result *AddLocalNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddLocalNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddLocalNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLocalNodes(request *AddLocalNodesRequest) (_result *AddLocalNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLocalNodesResponse{}
	_body, _err := client.AddLocalNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddNodesWithOptions(request *AddNodesRequest, runtime *util.RuntimeOptions) (_result *AddNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddNodes(request *AddNodesRequest) (_result *AddNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddNodesResponse{}
	_body, _err := client.AddNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddQueueWithOptions(request *AddQueueRequest, runtime *util.RuntimeOptions) (_result *AddQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddQueue"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddQueue(request *AddQueueRequest) (_result *AddQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddQueueResponse{}
	_body, _err := client.AddQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSecurityGroupWithOptions(request *AddSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *AddSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddSecurityGroup"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddSecurityGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSecurityGroup(request *AddSecurityGroupRequest) (_result *AddSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSecurityGroupResponse{}
	_body, _err := client.AddSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUsersWithOptions(request *AddUsersRequest, runtime *util.RuntimeOptions) (_result *AddUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUsers"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUsers(request *AddUsersRequest) (_result *AddUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUsersResponse{}
	_body, _err := client.AddUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ApplyNodes operation to specify the number of compute nodes, the number of vCPUs, and the memory size when you add nodes to a cluster.
 *
 * @param request ApplyNodesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ApplyNodesResponse
 */
func (client *Client) ApplyNodesWithOptions(request *ApplyNodesRequest, runtime *util.RuntimeOptions) (_result *ApplyNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ApplyNodes operation to specify the number of compute nodes, the number of vCPUs, and the memory size when you add nodes to a cluster.
 *
 * @param request ApplyNodesRequest
 * @return ApplyNodesResponse
 */
func (client *Client) ApplyNodes(request *ApplyNodesRequest) (_result *ApplyNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyNodesResponse{}
	_body, _err := client.ApplyNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After you create an Elastic High Performance Computing (E-HPC) cluster, you are charged for the cluster resources that you use. We recommend that you learn about the billing methods of E-HPC in advance. For more information, see [Billing overview](~~57844~~).
 *
 * @param request CreateClusterRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateClusterResponse
 */
func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After you create an Elastic High Performance Computing (E-HPC) cluster, you are charged for the cluster resources that you use. We recommend that you learn about the billing methods of E-HPC in advance. For more information, see [Billing overview](~~57844~~).
 *
 * @param request CreateClusterRequest
 * @return CreateClusterResponse
 */
func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGWSClusterWithOptions(request *CreateGWSClusterRequest, runtime *util.RuntimeOptions) (_result *CreateGWSClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGWSCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGWSClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGWSCluster(request *CreateGWSClusterRequest) (_result *CreateGWSClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGWSClusterResponse{}
	_body, _err := client.CreateGWSClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGWSImageWithOptions(request *CreateGWSImageRequest, runtime *util.RuntimeOptions) (_result *CreateGWSImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGWSImage"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGWSImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGWSImage(request *CreateGWSImageRequest) (_result *CreateGWSImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGWSImageResponse{}
	_body, _err := client.CreateGWSImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGWSInstanceWithOptions(request *CreateGWSInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateGWSInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGWSInstance"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGWSInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGWSInstance(request *CreateGWSInstanceRequest) (_result *CreateGWSInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGWSInstanceResponse{}
	_body, _err := client.CreateGWSInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHybridClusterWithOptions(request *CreateHybridClusterRequest, runtime *util.RuntimeOptions) (_result *CreateHybridClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHybridCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHybridClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHybridCluster(request *CreateHybridClusterRequest) (_result *CreateHybridClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHybridClusterResponse{}
	_body, _err := client.CreateHybridClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateJobFileWithOptions(request *CreateJobFileRequest, runtime *util.RuntimeOptions) (_result *CreateJobFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateJobFile"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateJobFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJobFile(request *CreateJobFileRequest) (_result *CreateJobFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateJobFileResponse{}
	_body, _err := client.CreateJobFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateJobTemplateWithOptions(request *CreateJobTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateJobTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateJobTemplate"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateJobTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJobTemplate(request *CreateJobTemplateRequest) (_result *CreateJobTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateJobTemplateResponse{}
	_body, _err := client.CreateJobTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After a cluster is released, the pay-as-you-go nodes and the subscription nodes that have expired are automatically released. The subscription nodes that have not expired are retained. If you need to release the subscription nodes that have not expired, change their billing method to pay-as-you-go. Before you release a cluster, make sure that you will no longer use the cluster.
 *
 * @param request DeleteClusterRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteClusterResponse
 */
func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After a cluster is released, the pay-as-you-go nodes and the subscription nodes that have expired are automatically released. The subscription nodes that have not expired are retained. If you need to release the subscription nodes that have not expired, change their billing method to pay-as-you-go. Before you release a cluster, make sure that you will no longer use the cluster.
 *
 * @param request DeleteClusterRequest
 * @return DeleteClusterResponse
 */
func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteContainerAppsWithOptions(request *DeleteContainerAppsRequest, runtime *util.RuntimeOptions) (_result *DeleteContainerAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteContainerApps"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteContainerAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteContainerApps(request *DeleteContainerAppsRequest) (_result *DeleteContainerAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteContainerAppsResponse{}
	_body, _err := client.DeleteContainerAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGWSClusterWithOptions(request *DeleteGWSClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteGWSClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGWSCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGWSClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGWSCluster(request *DeleteGWSClusterRequest) (_result *DeleteGWSClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGWSClusterResponse{}
	_body, _err := client.DeleteGWSClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGWSInstanceWithOptions(request *DeleteGWSInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteGWSInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGWSInstance"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGWSInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGWSInstance(request *DeleteGWSInstanceRequest) (_result *DeleteGWSInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGWSInstanceResponse{}
	_body, _err := client.DeleteGWSInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageWithOptions(request *DeleteImageRequest, runtime *util.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteImage"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImage(request *DeleteImageRequest) (_result *DeleteImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImageResponse{}
	_body, _err := client.DeleteImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobTemplatesWithOptions(request *DeleteJobTemplatesRequest, runtime *util.RuntimeOptions) (_result *DeleteJobTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteJobTemplates"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteJobTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJobTemplates(request *DeleteJobTemplatesRequest) (_result *DeleteJobTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteJobTemplatesResponse{}
	_body, _err := client.DeleteJobTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobsWithOptions(request *DeleteJobsRequest, runtime *util.RuntimeOptions) (_result *DeleteJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteJobs"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJobs(request *DeleteJobsRequest) (_result *DeleteJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteJobsResponse{}
	_body, _err := client.DeleteJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLocalImageWithOptions(request *DeleteLocalImageRequest, runtime *util.RuntimeOptions) (_result *DeleteLocalImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLocalImage"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLocalImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLocalImage(request *DeleteLocalImageRequest) (_result *DeleteLocalImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLocalImageResponse{}
	_body, _err := client.DeleteLocalImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you delete a compute node, we recommend that you export all job data from the node to prevent data loss.
 *
 * @param request DeleteNodesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteNodesResponse
 */
func (client *Client) DeleteNodesWithOptions(request *DeleteNodesRequest, runtime *util.RuntimeOptions) (_result *DeleteNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you delete a compute node, we recommend that you export all job data from the node to prevent data loss.
 *
 * @param request DeleteNodesRequest
 * @return DeleteNodesResponse
 */
func (client *Client) DeleteNodes(request *DeleteNodesRequest) (_result *DeleteNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNodesResponse{}
	_body, _err := client.DeleteNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteQueueWithOptions(request *DeleteQueueRequest, runtime *util.RuntimeOptions) (_result *DeleteQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteQueue"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteQueue(request *DeleteQueueRequest) (_result *DeleteQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteQueueResponse{}
	_body, _err := client.DeleteQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSecurityGroupWithOptions(request *DeleteSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSecurityGroup"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSecurityGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) (_result *DeleteSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecurityGroupResponse{}
	_body, _err := client.DeleteSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If you delete a user, only its information is deleted. The files stored in the /home directory for the user are still retained. For example, if you delete a user named user1, the files in the `/home/user1/` directory of the cluster are not deleted. However, a deleted user cannot be recovered. Even if you create another user that has the same name, the data that was retained for the deleted user is not reused.
 *
 * @param request DeleteUsersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteUsersResponse
 */
func (client *Client) DeleteUsersWithOptions(request *DeleteUsersRequest, runtime *util.RuntimeOptions) (_result *DeleteUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUsers"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you delete a user, only its information is deleted. The files stored in the /home directory for the user are still retained. For example, if you delete a user named user1, the files in the `/home/user1/` directory of the cluster are not deleted. However, a deleted user cannot be recovered. Even if you create another user that has the same name, the data that was retained for the deleted user is not reused.
 *
 * @param request DeleteUsersRequest
 * @return DeleteUsersResponse
 */
func (client *Client) DeleteUsers(request *DeleteUsersRequest) (_result *DeleteUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUsersResponse{}
	_body, _err := client.DeleteUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutoScaleConfigWithOptions(request *DescribeAutoScaleConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeAutoScaleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAutoScaleConfig"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAutoScaleConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutoScaleConfig(request *DescribeAutoScaleConfigRequest) (_result *DescribeAutoScaleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutoScaleConfigResponse{}
	_body, _err := client.DescribeAutoScaleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterWithOptions(request *DescribeClusterRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCluster(request *DescribeClusterRequest) (_result *DescribeClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterResponse{}
	_body, _err := client.DescribeClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerAppWithOptions(request *DescribeContainerAppRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContainerApp"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContainerAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerApp(request *DescribeContainerAppRequest) (_result *DescribeContainerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerAppResponse{}
	_body, _err := client.DescribeContainerAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEstackImageWithOptions(request *DescribeEstackImageRequest, runtime *util.RuntimeOptions) (_result *DescribeEstackImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEstackImage"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEstackImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEstackImage(request *DescribeEstackImageRequest) (_result *DescribeEstackImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEstackImageResponse{}
	_body, _err := client.DescribeEstackImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGWSClusterPolicyWithOptions(request *DescribeGWSClusterPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeGWSClusterPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AsyncMode)) {
		query["AsyncMode"] = request.AsyncMode
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGWSClusterPolicy"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGWSClusterPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGWSClusterPolicy(request *DescribeGWSClusterPolicyRequest) (_result *DescribeGWSClusterPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGWSClusterPolicyResponse{}
	_body, _err := client.DescribeGWSClusterPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGWSClustersWithOptions(request *DescribeGWSClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeGWSClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGWSClusters"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGWSClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGWSClusters(request *DescribeGWSClustersRequest) (_result *DescribeGWSClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGWSClustersResponse{}
	_body, _err := client.DescribeGWSClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGWSImagesWithOptions(request *DescribeGWSImagesRequest, runtime *util.RuntimeOptions) (_result *DescribeGWSImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGWSImages"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGWSImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGWSImages(request *DescribeGWSImagesRequest) (_result *DescribeGWSImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGWSImagesResponse{}
	_body, _err := client.DescribeGWSImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGWSInstancesWithOptions(request *DescribeGWSInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeGWSInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGWSInstances"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGWSInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGWSInstances(request *DescribeGWSInstancesRequest) (_result *DescribeGWSInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGWSInstancesResponse{}
	_body, _err := client.DescribeGWSInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageWithOptions(request *DescribeImageRequest, runtime *util.RuntimeOptions) (_result *DescribeImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImage"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImage(request *DescribeImageRequest) (_result *DescribeImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageResponse{}
	_body, _err := client.DescribeImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageGatewayConfigWithOptions(request *DescribeImageGatewayConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeImageGatewayConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImageGatewayConfig"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeImageGatewayConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImageGatewayConfig(request *DescribeImageGatewayConfigRequest) (_result *DescribeImageGatewayConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageGatewayConfigResponse{}
	_body, _err := client.DescribeImageGatewayConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImagePriceWithOptions(request *DescribeImagePriceRequest, runtime *util.RuntimeOptions) (_result *DescribeImagePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImagePrice"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeImagePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImagePrice(request *DescribeImagePriceRequest) (_result *DescribeImagePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImagePriceResponse{}
	_body, _err := client.DescribeImagePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeJobWithOptions(request *DescribeJobRequest, runtime *util.RuntimeOptions) (_result *DescribeJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeJob"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeJob(request *DescribeJobRequest) (_result *DescribeJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeJobResponse{}
	_body, _err := client.DescribeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNFSClientStatusWithOptions(request *DescribeNFSClientStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeNFSClientStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNFSClientStatus"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNFSClientStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNFSClientStatus(request *DescribeNFSClientStatusRequest) (_result *DescribeNFSClientStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNFSClientStatusResponse{}
	_body, _err := client.DescribeNFSClientStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ****
 *
 * @param request DescribePriceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribePriceResponse
 */
func (client *Client) DescribePriceWithOptions(request *DescribePriceRequest, runtime *util.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePrice"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ****
 *
 * @param request DescribePriceRequest
 * @return DescribePriceResponse
 */
func (client *Client) DescribePrice(request *DescribePriceRequest) (_result *DescribePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePriceResponse{}
	_body, _err := client.DescribePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditJobTemplateWithOptions(request *EditJobTemplateRequest, runtime *util.RuntimeOptions) (_result *EditJobTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EditJobTemplate"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EditJobTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditJobTemplate(request *EditJobTemplateRequest) (_result *EditJobTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditJobTemplateResponse{}
	_body, _err := client.EditJobTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccountingReportWithOptions(request *GetAccountingReportRequest, runtime *util.RuntimeOptions) (_result *GetAccountingReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccountingReport"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountingReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccountingReport(request *GetAccountingReportRequest) (_result *GetAccountingReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountingReportResponse{}
	_body, _err := client.GetAccountingReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAutoScaleConfigWithOptions(request *GetAutoScaleConfigRequest, runtime *util.RuntimeOptions) (_result *GetAutoScaleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAutoScaleConfig"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAutoScaleConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAutoScaleConfig(request *GetAutoScaleConfigRequest) (_result *GetAutoScaleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAutoScaleConfigResponse{}
	_body, _err := client.GetAutoScaleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCloudMetricLogsWithOptions(request *GetCloudMetricLogsRequest, runtime *util.RuntimeOptions) (_result *GetCloudMetricLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCloudMetricLogs"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCloudMetricLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCloudMetricLogs(request *GetCloudMetricLogsRequest) (_result *GetCloudMetricLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCloudMetricLogsResponse{}
	_body, _err := client.GetCloudMetricLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCloudMetricProfilingWithOptions(request *GetCloudMetricProfilingRequest, runtime *util.RuntimeOptions) (_result *GetCloudMetricProfilingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCloudMetricProfiling"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCloudMetricProfilingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCloudMetricProfiling(request *GetCloudMetricProfilingRequest) (_result *GetCloudMetricProfilingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCloudMetricProfilingResponse{}
	_body, _err := client.GetCloudMetricProfilingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClusterVolumesWithOptions(request *GetClusterVolumesRequest, runtime *util.RuntimeOptions) (_result *GetClusterVolumesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetClusterVolumes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClusterVolumesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClusterVolumes(request *GetClusterVolumesRequest) (_result *GetClusterVolumesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClusterVolumesResponse{}
	_body, _err := client.GetClusterVolumesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCommonImageWithOptions(request *GetCommonImageRequest, runtime *util.RuntimeOptions) (_result *GetCommonImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCommonImage"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCommonImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCommonImage(request *GetCommonImageRequest) (_result *GetCommonImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCommonImageResponse{}
	_body, _err := client.GetCommonImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGWSConnectTicketWithOptions(request *GetGWSConnectTicketRequest, runtime *util.RuntimeOptions) (_result *GetGWSConnectTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGWSConnectTicket"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGWSConnectTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGWSConnectTicket(request *GetGWSConnectTicketRequest) (_result *GetGWSConnectTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGWSConnectTicketResponse{}
	_body, _err := client.GetGWSConnectTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHybridClusterConfigWithOptions(request *GetHybridClusterConfigRequest, runtime *util.RuntimeOptions) (_result *GetHybridClusterConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHybridClusterConfig"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHybridClusterConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHybridClusterConfig(request *GetHybridClusterConfigRequest) (_result *GetHybridClusterConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHybridClusterConfigResponse{}
	_body, _err := client.GetHybridClusterConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIfEcsTypeSupportHtConfigWithOptions(request *GetIfEcsTypeSupportHtConfigRequest, runtime *util.RuntimeOptions) (_result *GetIfEcsTypeSupportHtConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIfEcsTypeSupportHtConfig"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIfEcsTypeSupportHtConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIfEcsTypeSupportHtConfig(request *GetIfEcsTypeSupportHtConfigRequest) (_result *GetIfEcsTypeSupportHtConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIfEcsTypeSupportHtConfigResponse{}
	_body, _err := client.GetIfEcsTypeSupportHtConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobLogWithOptions(request *GetJobLogRequest, runtime *util.RuntimeOptions) (_result *GetJobLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobLog"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobLog(request *GetJobLogRequest) (_result *GetJobLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobLogResponse{}
	_body, _err := client.GetJobLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPostScriptsWithOptions(request *GetPostScriptsRequest, runtime *util.RuntimeOptions) (_result *GetPostScriptsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPostScripts"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPostScriptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPostScripts(request *GetPostScriptsRequest) (_result *GetPostScriptsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPostScriptsResponse{}
	_body, _err := client.GetPostScriptsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSchedulerInfoWithOptions(request *GetSchedulerInfoRequest, runtime *util.RuntimeOptions) (_result *GetSchedulerInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSchedulerInfo"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSchedulerInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSchedulerInfo(request *GetSchedulerInfoRequest) (_result *GetSchedulerInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSchedulerInfoResponse{}
	_body, _err := client.GetSchedulerInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserImageWithOptions(request *GetUserImageRequest, runtime *util.RuntimeOptions) (_result *GetUserImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserImage"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserImage(request *GetUserImageRequest) (_result *GetUserImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserImageResponse{}
	_body, _err := client.GetUserImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVisualServiceStatusWithOptions(request *GetVisualServiceStatusRequest, runtime *util.RuntimeOptions) (_result *GetVisualServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVisualServiceStatus"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVisualServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVisualServiceStatus(request *GetVisualServiceStatusRequest) (_result *GetVisualServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVisualServiceStatusResponse{}
	_body, _err := client.GetVisualServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitializeEHPCWithOptions(request *InitializeEHPCRequest, runtime *util.RuntimeOptions) (_result *InitializeEHPCResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InitializeEHPC"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitializeEHPCResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitializeEHPC(request *InitializeEHPCRequest) (_result *InitializeEHPCResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitializeEHPCResponse{}
	_body, _err := client.InitializeEHPCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InspectImageWithOptions(request *InspectImageRequest, runtime *util.RuntimeOptions) (_result *InspectImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InspectImage"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InspectImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InspectImage(request *InspectImageRequest) (_result *InspectImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InspectImageResponse{}
	_body, _err := client.InspectImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallSoftwareWithOptions(request *InstallSoftwareRequest, runtime *util.RuntimeOptions) (_result *InstallSoftwareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallSoftware"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallSoftwareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallSoftware(request *InstallSoftwareRequest) (_result *InstallSoftwareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallSoftwareResponse{}
	_body, _err := client.InstallSoftwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvokeShellCommandWithOptions(request *InvokeShellCommandRequest, runtime *util.RuntimeOptions) (_result *InvokeShellCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InvokeShellCommand"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InvokeShellCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvokeShellCommand(request *InvokeShellCommandRequest) (_result *InvokeShellCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InvokeShellCommandResponse{}
	_body, _err := client.InvokeShellCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAvailableEcsTypesWithOptions(request *ListAvailableEcsTypesRequest, runtime *util.RuntimeOptions) (_result *ListAvailableEcsTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAvailableEcsTypes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAvailableEcsTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAvailableEcsTypes(request *ListAvailableEcsTypesRequest) (_result *ListAvailableEcsTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAvailableEcsTypesResponse{}
	_body, _err := client.ListAvailableEcsTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCloudMetricProfilingsWithOptions(request *ListCloudMetricProfilingsRequest, runtime *util.RuntimeOptions) (_result *ListCloudMetricProfilingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCloudMetricProfilings"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCloudMetricProfilingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCloudMetricProfilings(request *ListCloudMetricProfilingsRequest) (_result *ListCloudMetricProfilingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCloudMetricProfilingsResponse{}
	_body, _err := client.ListCloudMetricProfilingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterLogsWithOptions(request *ListClusterLogsRequest, runtime *util.RuntimeOptions) (_result *ListClusterLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterLogs"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterLogs(request *ListClusterLogsRequest) (_result *ListClusterLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterLogsResponse{}
	_body, _err := client.ListClusterLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusters"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClustersMetaWithOptions(request *ListClustersMetaRequest, runtime *util.RuntimeOptions) (_result *ListClustersMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClustersMeta"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClustersMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClustersMeta(request *ListClustersMetaRequest) (_result *ListClustersMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClustersMetaResponse{}
	_body, _err := client.ListClustersMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCommandsWithOptions(request *ListCommandsRequest, runtime *util.RuntimeOptions) (_result *ListCommandsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCommands"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCommandsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCommands(request *ListCommandsRequest) (_result *ListCommandsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCommandsResponse{}
	_body, _err := client.ListCommandsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCommunityImagesWithOptions(request *ListCommunityImagesRequest, runtime *util.RuntimeOptions) (_result *ListCommunityImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCommunityImages"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCommunityImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCommunityImages(request *ListCommunityImagesRequest) (_result *ListCommunityImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCommunityImagesResponse{}
	_body, _err := client.ListCommunityImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListContainerAppsWithOptions(request *ListContainerAppsRequest, runtime *util.RuntimeOptions) (_result *ListContainerAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListContainerApps"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListContainerAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListContainerApps(request *ListContainerAppsRequest) (_result *ListContainerAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListContainerAppsResponse{}
	_body, _err := client.ListContainerAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListContainerImagesWithOptions(request *ListContainerImagesRequest, runtime *util.RuntimeOptions) (_result *ListContainerImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListContainerImages"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListContainerImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListContainerImages(request *ListContainerImagesRequest) (_result *ListContainerImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListContainerImagesResponse{}
	_body, _err := client.ListContainerImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCpfsFileSystemsWithOptions(request *ListCpfsFileSystemsRequest, runtime *util.RuntimeOptions) (_result *ListCpfsFileSystemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCpfsFileSystems"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCpfsFileSystemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCpfsFileSystems(request *ListCpfsFileSystemsRequest) (_result *ListCpfsFileSystemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCpfsFileSystemsResponse{}
	_body, _err := client.ListCpfsFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCurrentClientVersionWithOptions(runtime *util.RuntimeOptions) (_result *ListCurrentClientVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListCurrentClientVersion"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCurrentClientVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCurrentClientVersion() (_result *ListCurrentClientVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCurrentClientVersionResponse{}
	_body, _err := client.ListCurrentClientVersionWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCustomImagesWithOptions(request *ListCustomImagesRequest, runtime *util.RuntimeOptions) (_result *ListCustomImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomImages"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCustomImages(request *ListCustomImagesRequest) (_result *ListCustomImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCustomImagesResponse{}
	_body, _err := client.ListCustomImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFileSystemWithMountTargetsWithOptions(request *ListFileSystemWithMountTargetsRequest, runtime *util.RuntimeOptions) (_result *ListFileSystemWithMountTargetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFileSystemWithMountTargets"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFileSystemWithMountTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFileSystemWithMountTargets(request *ListFileSystemWithMountTargetsRequest) (_result *ListFileSystemWithMountTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFileSystemWithMountTargetsResponse{}
	_body, _err := client.ListFileSystemWithMountTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImagesWithOptions(request *ListImagesRequest, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImages"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstalledSoftwareWithOptions(request *ListInstalledSoftwareRequest, runtime *util.RuntimeOptions) (_result *ListInstalledSoftwareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstalledSoftware"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstalledSoftwareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstalledSoftware(request *ListInstalledSoftwareRequest) (_result *ListInstalledSoftwareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstalledSoftwareResponse{}
	_body, _err := client.ListInstalledSoftwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInvocationResultsWithOptions(request *ListInvocationResultsRequest, runtime *util.RuntimeOptions) (_result *ListInvocationResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInvocationResults"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInvocationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInvocationResults(request *ListInvocationResultsRequest) (_result *ListInvocationResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInvocationResultsResponse{}
	_body, _err := client.ListInvocationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInvocationStatusWithOptions(request *ListInvocationStatusRequest, runtime *util.RuntimeOptions) (_result *ListInvocationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInvocationStatus"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInvocationStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInvocationStatus(request *ListInvocationStatusRequest) (_result *ListInvocationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInvocationStatusResponse{}
	_body, _err := client.ListInvocationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobTemplatesWithOptions(request *ListJobTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListJobTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobTemplates"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobTemplates(request *ListJobTemplatesRequest) (_result *ListJobTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJobTemplatesResponse{}
	_body, _err := client.ListJobTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobsWithOptions(request *ListJobsRequest, runtime *util.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobs"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobs(request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobsWithFiltersWithOptions(request *ListJobsWithFiltersRequest, runtime *util.RuntimeOptions) (_result *ListJobsWithFiltersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobsWithFilters"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobsWithFiltersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobsWithFilters(request *ListJobsWithFiltersRequest) (_result *ListJobsWithFiltersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJobsWithFiltersResponse{}
	_body, _err := client.ListJobsWithFiltersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodesWithOptions(request *ListNodesRequest, runtime *util.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodesByQueueWithOptions(request *ListNodesByQueueRequest, runtime *util.RuntimeOptions) (_result *ListNodesByQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodesByQueue"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesByQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodesByQueue(request *ListNodesByQueueRequest) (_result *ListNodesByQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesByQueueResponse{}
	_body, _err := client.ListNodesByQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodesNoPagingWithOptions(request *ListNodesNoPagingRequest, runtime *util.RuntimeOptions) (_result *ListNodesNoPagingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodesNoPaging"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesNoPagingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodesNoPaging(request *ListNodesNoPagingRequest) (_result *ListNodesNoPagingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesNoPagingResponse{}
	_body, _err := client.ListNodesNoPagingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPreferredEcsTypesWithOptions(request *ListPreferredEcsTypesRequest, runtime *util.RuntimeOptions) (_result *ListPreferredEcsTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPreferredEcsTypes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPreferredEcsTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPreferredEcsTypes(request *ListPreferredEcsTypesRequest) (_result *ListPreferredEcsTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPreferredEcsTypesResponse{}
	_body, _err := client.ListPreferredEcsTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQueuesWithOptions(request *ListQueuesRequest, runtime *util.RuntimeOptions) (_result *ListQueuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueues"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQueues(request *ListQueuesRequest) (_result *ListQueuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQueuesResponse{}
	_body, _err := client.ListQueuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegionsWithOptions(runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSecurityGroupsWithOptions(request *ListSecurityGroupsRequest, runtime *util.RuntimeOptions) (_result *ListSecurityGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSecurityGroups"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSecurityGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSecurityGroups(request *ListSecurityGroupsRequest) (_result *ListSecurityGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSecurityGroupsResponse{}
	_body, _err := client.ListSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSoftwaresWithOptions(request *ListSoftwaresRequest, runtime *util.RuntimeOptions) (_result *ListSoftwaresResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSoftwares"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSoftwaresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSoftwares(request *ListSoftwaresRequest) (_result *ListSoftwaresResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSoftwaresResponse{}
	_body, _err := client.ListSoftwaresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If you succeed in calling an asynchronous API operation, a response is generated before a resulting task is completed. Therefore, to query the result of the task, you can use the TaskId parameter returned by the API operation.
 *
 * @param request ListTasksRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTasksResponse
 */
func (client *Client) ListTasksWithOptions(request *ListTasksRequest, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTasks"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you succeed in calling an asynchronous API operation, a response is generated before a resulting task is completed. Therefore, to query the result of the task, you can use the TaskId parameter returned by the API operation.
 *
 * @param request ListTasksRequest
 * @return ListTasksResponse
 */
func (client *Client) ListTasks(request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUpgradeClientsWithOptions(request *ListUpgradeClientsRequest, runtime *util.RuntimeOptions) (_result *ListUpgradeClientsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUpgradeClients"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUpgradeClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUpgradeClients(request *ListUpgradeClientsRequest) (_result *ListUpgradeClientsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUpgradeClientsResponse{}
	_body, _err := client.ListUpgradeClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersAsyncWithOptions(request *ListUsersAsyncRequest, runtime *util.RuntimeOptions) (_result *ListUsersAsyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsersAsync"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsersAsync(request *ListUsersAsyncRequest) (_result *ListUsersAsyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersAsyncResponse{}
	_body, _err := client.ListUsersAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVolumesWithOptions(request *ListVolumesRequest, runtime *util.RuntimeOptions) (_result *ListVolumesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVolumes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVolumesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVolumes(request *ListVolumesRequest) (_result *ListVolumesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVolumesResponse{}
	_body, _err := client.ListVolumesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you modify the basic information of a cluster, you can call the [DescribeCluster](~~87126~~) operation to query details of the selected cluster.
 *
 * @param request ModifyClusterAttributesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyClusterAttributesResponse
 */
func (client *Client) ModifyClusterAttributesWithOptions(request *ModifyClusterAttributesRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterAttributes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you modify the basic information of a cluster, you can call the [DescribeCluster](~~87126~~) operation to query details of the selected cluster.
 *
 * @param request ModifyClusterAttributesRequest
 * @return ModifyClusterAttributesResponse
 */
func (client *Client) ModifyClusterAttributes(request *ModifyClusterAttributesRequest) (_result *ModifyClusterAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterAttributesResponse{}
	_body, _err := client.ModifyClusterAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyContainerAppAttributesWithOptions(request *ModifyContainerAppAttributesRequest, runtime *util.RuntimeOptions) (_result *ModifyContainerAppAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyContainerAppAttributes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyContainerAppAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyContainerAppAttributes(request *ModifyContainerAppAttributesRequest) (_result *ModifyContainerAppAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyContainerAppAttributesResponse{}
	_body, _err := client.ModifyContainerAppAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyImageGatewayConfigWithOptions(request *ModifyImageGatewayConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyImageGatewayConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyImageGatewayConfig"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyImageGatewayConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyImageGatewayConfig(request *ModifyImageGatewayConfigRequest) (_result *ModifyImageGatewayConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyImageGatewayConfigResponse{}
	_body, _err := client.ModifyImageGatewayConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserGroupsWithOptions(request *ModifyUserGroupsRequest, runtime *util.RuntimeOptions) (_result *ModifyUserGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUserGroups"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUserGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserGroups(request *ModifyUserGroupsRequest) (_result *ModifyUserGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserGroupsResponse{}
	_body, _err := client.ModifyUserGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserPasswordsWithOptions(request *ModifyUserPasswordsRequest, runtime *util.RuntimeOptions) (_result *ModifyUserPasswordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUserPasswords"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUserPasswordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserPasswords(request *ModifyUserPasswordsRequest) (_result *ModifyUserPasswordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserPasswordsResponse{}
	_body, _err := client.ModifyUserPasswordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyVisualServicePasswdWithOptions(request *ModifyVisualServicePasswdRequest, runtime *util.RuntimeOptions) (_result *ModifyVisualServicePasswdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVisualServicePasswd"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVisualServicePasswdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyVisualServicePasswd(request *ModifyVisualServicePasswdRequest) (_result *ModifyVisualServicePasswdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVisualServicePasswdResponse{}
	_body, _err := client.ModifyVisualServicePasswdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MountNFSWithOptions(request *MountNFSRequest, runtime *util.RuntimeOptions) (_result *MountNFSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MountNFS"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MountNFSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MountNFS(request *MountNFSRequest) (_result *MountNFSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MountNFSResponse{}
	_body, _err := client.MountNFSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PullImageWithOptions(request *PullImageRequest, runtime *util.RuntimeOptions) (_result *PullImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PullImage"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PullImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PullImage(request *PullImageRequest) (_result *PullImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PullImageResponse{}
	_body, _err := client.PullImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryServicePackAndPriceWithOptions(runtime *util.RuntimeOptions) (_result *QueryServicePackAndPriceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QueryServicePackAndPrice"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryServicePackAndPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryServicePackAndPrice() (_result *QueryServicePackAndPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryServicePackAndPriceResponse{}
	_body, _err := client.QueryServicePackAndPriceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the operation to reset and restore a cluster only when the cluster is in the Exception state. You can call the [ListClusters](~~87116~~) operation to query the ID and status of a cluster.
 * We recommend that you export all job data before you restore a cluster. When you reset and restore a cluster, take note of the following impacts:
 * *   The system disks of all nodes are changed. By default, new system disks are configured based on the settings that you specified when the cluster was created.
 * *   The data on the system disks and data disks of all cluster nodes is lost. The data includes user information, job information, scheduler queue information, and configuration data of auto-scaling queues. However, the data on Apsara File Storage NAS file systems is retained.
 * *   The self-managed queues in the cluster are deleted. All nodes are retained and migrated to the default queue of the cluster.
 *
 * @param request RecoverClusterRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RecoverClusterResponse
 */
func (client *Client) RecoverClusterWithOptions(request *RecoverClusterRequest, runtime *util.RuntimeOptions) (_result *RecoverClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecoverCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecoverClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the operation to reset and restore a cluster only when the cluster is in the Exception state. You can call the [ListClusters](~~87116~~) operation to query the ID and status of a cluster.
 * We recommend that you export all job data before you restore a cluster. When you reset and restore a cluster, take note of the following impacts:
 * *   The system disks of all nodes are changed. By default, new system disks are configured based on the settings that you specified when the cluster was created.
 * *   The data on the system disks and data disks of all cluster nodes is lost. The data includes user information, job information, scheduler queue information, and configuration data of auto-scaling queues. However, the data on Apsara File Storage NAS file systems is retained.
 * *   The self-managed queues in the cluster are deleted. All nodes are retained and migrated to the default queue of the cluster.
 *
 * @param request RecoverClusterRequest
 * @return RecoverClusterResponse
 */
func (client *Client) RecoverCluster(request *RecoverClusterRequest) (_result *RecoverClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecoverClusterResponse{}
	_body, _err := client.RecoverClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RerunJobsWithOptions(request *RerunJobsRequest, runtime *util.RuntimeOptions) (_result *RerunJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RerunJobs"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RerunJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RerunJobs(request *RerunJobsRequest) (_result *RerunJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RerunJobsResponse{}
	_body, _err := client.RerunJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After a node is reset, the operating system and software return to their initial states. To ensure that jobs run as expected, we recommend that you do not reset running nodes unless you need to perform crash recovery.
 *
 * @param request ResetNodesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ResetNodesResponse
 */
func (client *Client) ResetNodesWithOptions(request *ResetNodesRequest, runtime *util.RuntimeOptions) (_result *ResetNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After a node is reset, the operating system and software return to their initial states. To ensure that jobs run as expected, we recommend that you do not reset running nodes unless you need to perform crash recovery.
 *
 * @param request ResetNodesRequest
 * @return ResetNodesResponse
 */
func (client *Client) ResetNodes(request *ResetNodesRequest) (_result *ResetNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetNodesResponse{}
	_body, _err := client.ResetNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunCloudMetricProfilingWithOptions(request *RunCloudMetricProfilingRequest, runtime *util.RuntimeOptions) (_result *RunCloudMetricProfilingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunCloudMetricProfiling"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunCloudMetricProfilingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunCloudMetricProfiling(request *RunCloudMetricProfilingRequest) (_result *RunCloudMetricProfilingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCloudMetricProfilingResponse{}
	_body, _err := client.RunCloudMetricProfilingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If you specify different auto scaling settings in the Queue Configuration section and Global Configurations section on the Auto Scale page, the settings in the Queue Configuration section prevail.
 *
 * @param request SetAutoScaleConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SetAutoScaleConfigResponse
 */
func (client *Client) SetAutoScaleConfigWithOptions(request *SetAutoScaleConfigRequest, runtime *util.RuntimeOptions) (_result *SetAutoScaleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetAutoScaleConfig"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetAutoScaleConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you specify different auto scaling settings in the Queue Configuration section and Global Configurations section on the Auto Scale page, the settings in the Queue Configuration section prevail.
 *
 * @param request SetAutoScaleConfigRequest
 * @return SetAutoScaleConfigResponse
 */
func (client *Client) SetAutoScaleConfig(request *SetAutoScaleConfigRequest) (_result *SetAutoScaleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAutoScaleConfigResponse{}
	_body, _err := client.SetAutoScaleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGWSClusterPolicyWithOptions(request *SetGWSClusterPolicyRequest, runtime *util.RuntimeOptions) (_result *SetGWSClusterPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AsyncMode)) {
		query["AsyncMode"] = request.AsyncMode
	}

	if !tea.BoolValue(util.IsUnset(request.Clipboard)) {
		query["Clipboard"] = request.Clipboard
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.LocalDrive)) {
		query["LocalDrive"] = request.LocalDrive
	}

	if !tea.BoolValue(util.IsUnset(request.UdpPort)) {
		query["UdpPort"] = request.UdpPort
	}

	if !tea.BoolValue(util.IsUnset(request.UsbRedirect)) {
		query["UsbRedirect"] = request.UsbRedirect
	}

	if !tea.BoolValue(util.IsUnset(request.Watermark)) {
		query["Watermark"] = request.Watermark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetGWSClusterPolicy"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetGWSClusterPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGWSClusterPolicy(request *SetGWSClusterPolicyRequest) (_result *SetGWSClusterPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGWSClusterPolicyResponse{}
	_body, _err := client.SetGWSClusterPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGWSInstanceNameWithOptions(request *SetGWSInstanceNameRequest, runtime *util.RuntimeOptions) (_result *SetGWSInstanceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetGWSInstanceName"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetGWSInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGWSInstanceName(request *SetGWSInstanceNameRequest) (_result *SetGWSInstanceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGWSInstanceNameResponse{}
	_body, _err := client.SetGWSInstanceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGWSInstanceUserWithOptions(request *SetGWSInstanceUserRequest, runtime *util.RuntimeOptions) (_result *SetGWSInstanceUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetGWSInstanceUser"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetGWSInstanceUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGWSInstanceUser(request *SetGWSInstanceUserRequest) (_result *SetGWSInstanceUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGWSInstanceUserResponse{}
	_body, _err := client.SetGWSInstanceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetPostScriptsWithOptions(request *SetPostScriptsRequest, runtime *util.RuntimeOptions) (_result *SetPostScriptsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetPostScripts"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetPostScriptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetPostScripts(request *SetPostScriptsRequest) (_result *SetPostScriptsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetPostScriptsResponse{}
	_body, _err := client.SetPostScriptsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetQueueWithOptions(request *SetQueueRequest, runtime *util.RuntimeOptions) (_result *SetQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetQueue"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetQueue(request *SetQueueRequest) (_result *SetQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetQueueResponse{}
	_body, _err := client.SetQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetSchedulerInfoWithOptions(request *SetSchedulerInfoRequest, runtime *util.RuntimeOptions) (_result *SetSchedulerInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetSchedulerInfo"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetSchedulerInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetSchedulerInfo(request *SetSchedulerInfoRequest) (_result *SetSchedulerInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetSchedulerInfoResponse{}
	_body, _err := client.SetSchedulerInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartClusterWithOptions(request *StartClusterRequest, runtime *util.RuntimeOptions) (_result *StartClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCluster(request *StartClusterRequest) (_result *StartClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartClusterResponse{}
	_body, _err := client.StartClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartGWSInstanceWithOptions(request *StartGWSInstanceRequest, runtime *util.RuntimeOptions) (_result *StartGWSInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartGWSInstance"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartGWSInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartGWSInstance(request *StartGWSInstanceRequest) (_result *StartGWSInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartGWSInstanceResponse{}
	_body, _err := client.StartGWSInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartNodesWithOptions(request *StartNodesRequest, runtime *util.RuntimeOptions) (_result *StartNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartNodes(request *StartNodesRequest) (_result *StartNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartNodesResponse{}
	_body, _err := client.StartNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartVisualServiceWithOptions(request *StartVisualServiceRequest, runtime *util.RuntimeOptions) (_result *StartVisualServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartVisualService"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartVisualServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartVisualService(request *StartVisualServiceRequest) (_result *StartVisualServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartVisualServiceResponse{}
	_body, _err := client.StartVisualServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If you stop a subscription compute node, its billing is not affected. If you stop a pay-as-you-go compute node for which you have enabled the *economical mode*, you are no longer charged for its computing resources. For more information, see [Economical mode](~~63353~~).
 *
 * @param request StopClusterRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StopClusterResponse
 */
func (client *Client) StopClusterWithOptions(request *StopClusterRequest, runtime *util.RuntimeOptions) (_result *StopClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you stop a subscription compute node, its billing is not affected. If you stop a pay-as-you-go compute node for which you have enabled the *economical mode*, you are no longer charged for its computing resources. For more information, see [Economical mode](~~63353~~).
 *
 * @param request StopClusterRequest
 * @return StopClusterResponse
 */
func (client *Client) StopCluster(request *StopClusterRequest) (_result *StopClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopClusterResponse{}
	_body, _err := client.StopClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopGWSInstanceWithOptions(request *StopGWSInstanceRequest, runtime *util.RuntimeOptions) (_result *StopGWSInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopGWSInstance"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopGWSInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopGWSInstance(request *StopGWSInstanceRequest) (_result *StopGWSInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopGWSInstanceResponse{}
	_body, _err := client.StopGWSInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopJobsWithOptions(request *StopJobsRequest, runtime *util.RuntimeOptions) (_result *StopJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopJobs"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopJobs(request *StopJobsRequest) (_result *StopJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopJobsResponse{}
	_body, _err := client.StopJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopNodesWithOptions(request *StopNodesRequest, runtime *util.RuntimeOptions) (_result *StopNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopNodes(request *StopNodesRequest) (_result *StopNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopNodesResponse{}
	_body, _err := client.StopNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopVisualServiceWithOptions(request *StopVisualServiceRequest, runtime *util.RuntimeOptions) (_result *StopVisualServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopVisualService"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopVisualServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopVisualService(request *StopVisualServiceRequest) (_result *StopVisualServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopVisualServiceResponse{}
	_body, _err := client.StopVisualServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you submit a job in a cluster, you must upload a job file to the cluster, for example, job.sh. For more information, see [CreateJobFile](~~159049~~).
 *
 * @param request SubmitJobRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SubmitJobResponse
 */
func (client *Client) SubmitJobWithOptions(request *SubmitJobRequest, runtime *util.RuntimeOptions) (_result *SubmitJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitJob"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you submit a job in a cluster, you must upload a job file to the cluster, for example, job.sh. For more information, see [CreateJobFile](~~159049~~).
 *
 * @param request SubmitJobRequest
 * @return SubmitJobResponse
 */
func (client *Client) SubmitJob(request *SubmitJobRequest) (_result *SubmitJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitJobResponse{}
	_body, _err := client.SubmitJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SummaryImagesWithOptions(request *SummaryImagesRequest, runtime *util.RuntimeOptions) (_result *SummaryImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SummaryImages"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SummaryImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SummaryImages(request *SummaryImagesRequest) (_result *SummaryImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SummaryImagesResponse{}
	_body, _err := client.SummaryImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SummaryImagesInfoWithOptions(request *SummaryImagesInfoRequest, runtime *util.RuntimeOptions) (_result *SummaryImagesInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SummaryImagesInfo"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SummaryImagesInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SummaryImagesInfo(request *SummaryImagesInfoRequest) (_result *SummaryImagesInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SummaryImagesInfoResponse{}
	_body, _err := client.SummaryImagesInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncUsersWithOptions(request *SyncUsersRequest, runtime *util.RuntimeOptions) (_result *SyncUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncUsers"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncUsers(request *SyncUsersRequest) (_result *SyncUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncUsersResponse{}
	_body, _err := client.SyncUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnTagResourcesWithOptions(request *UnTagResourcesRequest, runtime *util.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnTagResources"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnTagResources(request *UnTagResourcesRequest) (_result *UnTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.UnTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UninstallSoftwareWithOptions(request *UninstallSoftwareRequest, runtime *util.RuntimeOptions) (_result *UninstallSoftwareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UninstallSoftware"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UninstallSoftwareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UninstallSoftware(request *UninstallSoftwareRequest) (_result *UninstallSoftwareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UninstallSoftwareResponse{}
	_body, _err := client.UninstallSoftwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateClusterVolumesWithOptions(request *UpdateClusterVolumesRequest, runtime *util.RuntimeOptions) (_result *UpdateClusterVolumesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateClusterVolumes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateClusterVolumesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateClusterVolumes(request *UpdateClusterVolumesRequest) (_result *UpdateClusterVolumesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateClusterVolumesResponse{}
	_body, _err := client.UpdateClusterVolumesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After you update the instance types of a resource group, the nodes that you add by scaling out the cluster are automatically included in the resource group.
 *
 * @param request UpdateQueueConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateQueueConfigResponse
 */
func (client *Client) UpdateQueueConfigWithOptions(request *UpdateQueueConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateQueueConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateQueueConfig"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateQueueConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After you update the instance types of a resource group, the nodes that you add by scaling out the cluster are automatically included in the resource group.
 *
 * @param request UpdateQueueConfigRequest
 * @return UpdateQueueConfigResponse
 */
func (client *Client) UpdateQueueConfig(request *UpdateQueueConfigRequest) (_result *UpdateQueueConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateQueueConfigResponse{}
	_body, _err := client.UpdateQueueConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeClientWithOptions(request *UpgradeClientRequest, runtime *util.RuntimeOptions) (_result *UpgradeClientResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeClient"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeClient(request *UpgradeClientRequest) (_result *UpgradeClientResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeClientResponse{}
	_body, _err := client.UpgradeClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
