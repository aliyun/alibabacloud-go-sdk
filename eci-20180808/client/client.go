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

type CommitContainerRequest struct {
	// The configurations of the access credential for the Container Registry Enterprise Edition instance.
	AcrRegistryInfo *CommitContainerRequestAcrRegistryInfo `json:"AcrRegistryInfo,omitempty" xml:"AcrRegistryInfo,omitempty" type:"Struct"`
	// The ARN that is required for authorization.
	Arn *CommitContainerRequestArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Struct"`
	// The ID of the container group.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The name of the container.
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The image of the container.
	Image        *CommitContainerRequestImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	OwnerAccount *string                      `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CommitContainerRequest) String() string {
	return tea.Prettify(s)
}

func (s CommitContainerRequest) GoString() string {
	return s.String()
}

func (s *CommitContainerRequest) SetAcrRegistryInfo(v *CommitContainerRequestAcrRegistryInfo) *CommitContainerRequest {
	s.AcrRegistryInfo = v
	return s
}

func (s *CommitContainerRequest) SetArn(v *CommitContainerRequestArn) *CommitContainerRequest {
	s.Arn = v
	return s
}

func (s *CommitContainerRequest) SetContainerGroupId(v string) *CommitContainerRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *CommitContainerRequest) SetContainerName(v string) *CommitContainerRequest {
	s.ContainerName = &v
	return s
}

func (s *CommitContainerRequest) SetImage(v *CommitContainerRequestImage) *CommitContainerRequest {
	s.Image = v
	return s
}

func (s *CommitContainerRequest) SetOwnerAccount(v string) *CommitContainerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CommitContainerRequest) SetOwnerId(v int64) *CommitContainerRequest {
	s.OwnerId = &v
	return s
}

func (s *CommitContainerRequest) SetRegionId(v string) *CommitContainerRequest {
	s.RegionId = &v
	return s
}

func (s *CommitContainerRequest) SetResourceOwnerAccount(v string) *CommitContainerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CommitContainerRequest) SetResourceOwnerId(v int64) *CommitContainerRequest {
	s.ResourceOwnerId = &v
	return s
}

type CommitContainerRequestAcrRegistryInfo struct {
	// The configurations of the access credential for the Container Registry Enterprise Edition instance: the ID of the Enterprise Edition instance. This is a required parameter.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CommitContainerRequestAcrRegistryInfo) String() string {
	return tea.Prettify(s)
}

func (s CommitContainerRequestAcrRegistryInfo) GoString() string {
	return s.String()
}

func (s *CommitContainerRequestAcrRegistryInfo) SetInstanceId(v string) *CommitContainerRequestAcrRegistryInfo {
	s.InstanceId = &v
	return s
}

func (s *CommitContainerRequestAcrRegistryInfo) SetRegionId(v string) *CommitContainerRequestAcrRegistryInfo {
	s.RegionId = &v
	return s
}

type CommitContainerRequestArn struct {
	// The ARN of the RAM role.
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The type of the authorization.
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s CommitContainerRequestArn) String() string {
	return tea.Prettify(s)
}

func (s CommitContainerRequestArn) GoString() string {
	return s.String()
}

func (s *CommitContainerRequestArn) SetRoleArn(v string) *CommitContainerRequestArn {
	s.RoleArn = &v
	return s
}

func (s *CommitContainerRequestArn) SetRoleType(v string) *CommitContainerRequestArn {
	s.RoleType = &v
	return s
}

type CommitContainerRequestImage struct {
	// The authorization of the image.
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// The message about the image.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The image repository.
	Repository *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	// The tag of the image. This parameter is empty by default, which indicates that the tag is not modified.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CommitContainerRequestImage) String() string {
	return tea.Prettify(s)
}

func (s CommitContainerRequestImage) GoString() string {
	return s.String()
}

func (s *CommitContainerRequestImage) SetAuthor(v string) *CommitContainerRequestImage {
	s.Author = &v
	return s
}

func (s *CommitContainerRequestImage) SetMessage(v string) *CommitContainerRequestImage {
	s.Message = &v
	return s
}

func (s *CommitContainerRequestImage) SetRepository(v string) *CommitContainerRequestImage {
	s.Repository = &v
	return s
}

func (s *CommitContainerRequestImage) SetTag(v string) *CommitContainerRequestImage {
	s.Tag = &v
	return s
}

type CommitContainerResponseBody struct {
	// The unique ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CommitContainerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CommitContainerResponseBody) GoString() string {
	return s.String()
}

func (s *CommitContainerResponseBody) SetRequestId(v string) *CommitContainerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CommitContainerResponseBody) SetTaskId(v string) *CommitContainerResponseBody {
	s.TaskId = &v
	return s
}

type CommitContainerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CommitContainerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CommitContainerResponse) String() string {
	return tea.Prettify(s)
}

func (s CommitContainerResponse) GoString() string {
	return s.String()
}

func (s *CommitContainerResponse) SetHeaders(v map[string]*string) *CommitContainerResponse {
	s.Headers = v
	return s
}

func (s *CommitContainerResponse) SetStatusCode(v int32) *CommitContainerResponse {
	s.StatusCode = &v
	return s
}

func (s *CommitContainerResponse) SetBody(v *CommitContainerResponseBody) *CommitContainerResponse {
	s.Body = v
	return s
}

type CreateContainerGroupRequest struct {
	DnsConfig           *CreateContainerGroupRequestDnsConfig           `json:"DnsConfig,omitempty" xml:"DnsConfig,omitempty" type:"Struct"`
	HostSecurityContext *CreateContainerGroupRequestHostSecurityContext `json:"HostSecurityContext,omitempty" xml:"HostSecurityContext,omitempty" type:"Struct"`
	SecurityContext     *CreateContainerGroupRequestSecurityContext     `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	// Details of the Container Registry Enterprise Edition instances.
	AcrRegistryInfo []*CreateContainerGroupRequestAcrRegistryInfo `json:"AcrRegistryInfo,omitempty" xml:"AcrRegistryInfo,omitempty" type:"Repeated"`
	// The validity period of the elastic container instance. When this period expires, the instance is forced to exit. Unit: seconds.
	ActiveDeadlineSeconds *int64 `json:"ActiveDeadlineSeconds,omitempty" xml:"ActiveDeadlineSeconds,omitempty"`
	// Specifies whether to automatically create an EIP and associate it with the elastic container instance.
	AutoCreateEip *bool `json:"AutoCreateEip,omitempty" xml:"AutoCreateEip,omitempty"`
	// Specifies whether to automatically match image caches. Default value: false.
	AutoMatchImageCache *bool `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency of requests?](~~25693~~)
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The information about containers.
	Container []*CreateContainerGroupRequestContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Repeated"`
	// The name of the elastic container instance. Format requirements:
	//
	// *   The name must be 2 to 128 characters in length
	// *   The name can contain lowercase letters, digits, and hyphens (-). It cannot start or end with a hyphen (-).
	ContainerGroupName *string `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	// Specifies whether to enable container resource view. Container resource view displays the actual container resource data instead of data of the host. If the specifications of the generated elastic container instance are larger than the specifications that you request for when you create the instance, you can enable the ContainerResourceView feature to ensure that the resources that you view in the container are the same as the resources that you request for.
	ContainerResourceView *bool `json:"ContainerResourceView,omitempty" xml:"ContainerResourceView,omitempty"`
	// The path to store core dump files. For more information, see [Save core files to volumes](~~167801~~).
	//
	// > The path cannot start with a vertical bar (`|`). You cannot use core dump files to configure executable programs.
	CorePattern *string `json:"CorePattern,omitempty" xml:"CorePattern,omitempty"`
	// The number of vCPUs that you want to allocate to the elastic container instance.
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of physical CPU cores. You can specify this parameter only for specific instance types. For more information, see [Specify custom CPU options](~~197781~~).
	CpuOptionsCore *int32 `json:"CpuOptionsCore,omitempty" xml:"CpuOptionsCore,omitempty"`
	// This parameter is not available.
	CpuOptionsNuma *string `json:"CpuOptionsNuma,omitempty" xml:"CpuOptionsNuma,omitempty"`
	// The number of threads per core. You can specify this parameter only for specific instance types. If you set this parameter to 1, Hyper-Threading is disabled. For more information, see [Specify custom CPU options](~~197781~~).
	CpuOptionsThreadsPerCore *int32 `json:"CpuOptionsThreadsPerCore,omitempty" xml:"CpuOptionsThreadsPerCore,omitempty"`
	// The Domain Name System (DNS) policy. Valid values:
	//
	// *   None: uses the DNS that is specified for DnsConfig-related parameters.
	// *   Default: uses the DNS that is specified for the runtime environment.
	DnsPolicy *string `json:"DnsPolicy,omitempty" xml:"DnsPolicy,omitempty"`
	// The maximum outbound bandwidth. Unit: bytes.
	EgressBandwidth *int64 `json:"EgressBandwidth,omitempty" xml:"EgressBandwidth,omitempty"`
	// The bandwidth of the EIP. Unit: Mbit/s. Default value: 5.\
	// You can specify this parameter when you set AutoCreateEip to true.
	EipBandwidth *int32 `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	// Specifies the EIP bandwidth plan that you want to use.
	EipCommonBandwidthPackage *string `json:"EipCommonBandwidthPackage,omitempty" xml:"EipCommonBandwidthPackage,omitempty"`
	// Specifies the line type of the EIP. Valid values:
	//
	// *   BPG: BGP (Multi-ISP) line
	// *   BGP_PRO: BGP (Multi-ISP) Pro line
	EipISP *string `json:"EipISP,omitempty" xml:"EipISP,omitempty"`
	// The ID of the elastic IP address (EIP).
	EipInstanceId *string `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	// The increased storage capacity of the temporary storage space. Unit: GiB.\
	// For more information, see [Increase the storage capacity of the temporary storage space](~~204066~~).
	EphemeralStorage *int32 `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	// The alias of the elastic container instance.
	HostAliase []*CreateContainerGroupRequestHostAliase `json:"HostAliase,omitempty" xml:"HostAliase,omitempty" type:"Repeated"`
	// The hostname of the instance.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The image acceleration mode. Valid values:
	//
	// *   nydus: Nydus is used to accelerate image pulling. The images must support Nydus.
	// *   dadi: DADI is used to accelerate image pulling. The images must support DADI.
	// *   p2p: P2P is used to accelerate image pulling. The images must support P2P.
	// *   imc: Image caches are used to accelerate image pulling.
	ImageAccelerateMode *string `json:"ImageAccelerateMode,omitempty" xml:"ImageAccelerateMode,omitempty"`
	// The information about the image repository.
	ImageRegistryCredential []*CreateContainerGroupRequestImageRegistryCredential `json:"ImageRegistryCredential,omitempty" xml:"ImageRegistryCredential,omitempty" type:"Repeated"`
	// The ID of the image cache. For more information, see [Use image caches to accelerate the creation of instances](~~141281~~).
	ImageSnapshotId *string `json:"ImageSnapshotId,omitempty" xml:"ImageSnapshotId,omitempty"`
	// The maximum inbound bandwidth. Unit: bytes.
	IngressBandwidth *int64 `json:"IngressBandwidth,omitempty" xml:"IngressBandwidth,omitempty"`
	// The init containers.
	InitContainer []*CreateContainerGroupRequestInitContainer `json:"InitContainer,omitempty" xml:"InitContainer,omitempty" type:"Repeated"`
	// The address of the self-managed image repository. When you create an elastic container instance by using an image in a self-managed image repository that uses a self-signed certificate, you must specify this parameter to skip the certificate authentication. This prevents image pull failures caused by certificate authentication failures.
	InsecureRegistry *string `json:"InsecureRegistry,omitempty" xml:"InsecureRegistry,omitempty"`
	// The ECS instance type. Different instance types are supported. For more information, see [Specify an ECS instance type to create an elastic container instance](~~114664~~).
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of IPv6 addresses. Set the value to 1. You can assign only one IPv6 address to an elastic container instance.
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The peak Internet bandwidth of the IPv6 address when the Ipv6GatewayBandwidthEnable parameter is set to true. Valid values:
	//
	// *   If the billing method for the Internet bandwidth of the IPv6 gateway is pay-by-bandwidth, the Internet bandwidth of the IPv6 address ranges from 1 to 2,000 Mbit/s.
	//
	// *   If the billing method for the Internet bandwidth of the IPv6 gateway is pay-by-traffic, the Internet bandwidth range of the IPv6 address is based on the edition of the IPv6 gateway.
	//
	//     *   If the IPv6 gateway is of Free Edition, the Internet bandwidth of the IPv6 address ranges from 1 to 200 Mbit/s.
	//     *   If the IPv6 gateway is of Enterprise Edition, the Internet bandwidth of the IPv6 address ranges from 1 to 500 Mbit/s.
	//     *   If the IPv6 gateway is of Enhanced Enterprise Edition, the Internet bandwidth of the IPv6 address ranges from 1 to 1,000 Mbit/s.
	//
	// The default value is the maximum value in the Internet bandwidth range of the IPv6 gateway.
	Ipv6GatewayBandwidth *string `json:"Ipv6GatewayBandwidth,omitempty" xml:"Ipv6GatewayBandwidth,omitempty"`
	// Specifies whether to enable IPv6 Internet access for the elastic container instance.
	Ipv6GatewayBandwidthEnable *bool `json:"Ipv6GatewayBandwidthEnable,omitempty" xml:"Ipv6GatewayBandwidthEnable,omitempty"`
	// The memory size that you want to allocate to the elastic container instance. Unit: GiB.
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The domain names of the NTP server.
	NtpServer    []*string `json:"NtpServer,omitempty" xml:"NtpServer,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The address of the self-managed image repository. When you create an elastic container instance by using an image in a self-managed image repository that uses the HTTP protocol, you must specify this parameter. This allows Elastic Container Instance to pull the image over the HTTP protocol instead over the default HTTPS protocol. This prevents image pull failures caused by different protocols.
	PlainHttpRegistry *string `json:"PlainHttpRegistry,omitempty" xml:"PlainHttpRegistry,omitempty"`
	// The name of the RAM role that you want to associate with the elastic container instance. You can use the RAM role to access elastic container instances and ECS instances. For more information, see [Use an instance RAM role by calling API operations](~~61178~~).
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The restart policy of the elastic container instance. Valid values:
	//
	// *   Always: Always restarts the instance.
	// *   Never: Never restarts the instance.
	// *   OnFailure: Restarts the instance when the last start failed.
	//
	// Default value: Always.
	RestartPolicy *string `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	// The resource scheduling policy when you specify multiple zones to create an elastic container instance. To specify multiple zones, you can use the VSwitchId to specify multiple vSwitches. Valid values:
	//
	// *   VSwitchOrdered: The system schedules resources in the sequence of the vSwitches.
	// *   VSwitchRandom: The system schedules resources at random.
	//
	// For more information, see [Specify multiple zones to create an elastic container instance](~~157290~~).
	ScheduleStrategy *string `json:"ScheduleStrategy,omitempty" xml:"ScheduleStrategy,omitempty"`
	// The ID of the security group to which the instance is assigned. Instances within the same security group can access each other.
	//
	// If you do not specify a security group, the system automatically uses the default security group in the region that you selected. Make sure that the inbound rules of the security group contain the container protocols and port numbers that you want to expose. If you do not have a default security group in the region, the system creates a default security group, and then adds the container protocols and port numbers that you specified to the inbound rules of the security group.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// Specifies whether to use a shared namespace. Default value: false.
	ShareProcessNamespace *bool `json:"ShareProcessNamespace,omitempty" xml:"ShareProcessNamespace,omitempty"`
	// The protection period of the preemptible instance. Unit: hours. Default value: 1. Valid values: 0 to 6.
	SpotDuration *int64 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The maximum hourly price of the preemptible elastic container instance. The value can contain up to three decimal places.
	//
	// If you set SpotStrategy to SpotWithPriceLimit, you must specify SpotPriceLimit.
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding policy for the elastic container instance. Valid values:
	//
	// *   NoSpot: The instance is created as a regular pay-as-you-go instance.
	// *   SpotWithPriceLimit: The instance is created as a preemptible instance with a user-defined maximum hourly price.
	// *   SpotAsPriceGo: The instance is created as a preemptible instance for which the market price at the time of purchase is automatically used as the bid price.
	//
	// Default value: NoSpot.
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// Specifies whether to enable periodical execution.
	//
	// *   true: enables periodical execution.
	// *   false: disables periodical execution.
	StrictSpot *bool `json:"StrictSpot,omitempty" xml:"StrictSpot,omitempty"`
	// The tags that you want to bind with the instance. You can bind a maximum of 20 tags. For more information, see [Use tags to manage elastic container instances](~~146608~~).
	Tag []*CreateContainerGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The buffer time during which the program handles operations before the program stops.
	TerminationGracePeriodSeconds *int64 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	// The ID of the vSwitch to which the instance is connected. You can specify up to 10 vSwitch IDs. Separate multiple vSwitch IDs with commas (,). Example: `vsw-***,vsw-***`.
	//
	// If no vSwitch is specified, the system automatically uses the default vSwitch in the default VPC in the region that you selected. If you do not have a default VPC or a default vSwitch in the region, the system automatically creates a default VPC and a default vSwitch.
	//
	// > The number of IP addresses in the vSwitch CIDR block determines the maximum number of elastic container instances that can be created for the vSwitch. Before you create elastic container instances, plan the CIDR block of the vSwitch.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Details about the volumes.
	Volume []*CreateContainerGroupRequestVolume `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Repeated"`
	// The ID of the zone in which the elastic container instance is deployed. If you do not specify this parameter, the system selects a zone.
	//
	// This parameter is empty by default.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateContainerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequest) SetDnsConfig(v *CreateContainerGroupRequestDnsConfig) *CreateContainerGroupRequest {
	s.DnsConfig = v
	return s
}

func (s *CreateContainerGroupRequest) SetHostSecurityContext(v *CreateContainerGroupRequestHostSecurityContext) *CreateContainerGroupRequest {
	s.HostSecurityContext = v
	return s
}

func (s *CreateContainerGroupRequest) SetSecurityContext(v *CreateContainerGroupRequestSecurityContext) *CreateContainerGroupRequest {
	s.SecurityContext = v
	return s
}

func (s *CreateContainerGroupRequest) SetAcrRegistryInfo(v []*CreateContainerGroupRequestAcrRegistryInfo) *CreateContainerGroupRequest {
	s.AcrRegistryInfo = v
	return s
}

func (s *CreateContainerGroupRequest) SetActiveDeadlineSeconds(v int64) *CreateContainerGroupRequest {
	s.ActiveDeadlineSeconds = &v
	return s
}

func (s *CreateContainerGroupRequest) SetAutoCreateEip(v bool) *CreateContainerGroupRequest {
	s.AutoCreateEip = &v
	return s
}

func (s *CreateContainerGroupRequest) SetAutoMatchImageCache(v bool) *CreateContainerGroupRequest {
	s.AutoMatchImageCache = &v
	return s
}

func (s *CreateContainerGroupRequest) SetClientToken(v string) *CreateContainerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateContainerGroupRequest) SetContainer(v []*CreateContainerGroupRequestContainer) *CreateContainerGroupRequest {
	s.Container = v
	return s
}

func (s *CreateContainerGroupRequest) SetContainerGroupName(v string) *CreateContainerGroupRequest {
	s.ContainerGroupName = &v
	return s
}

func (s *CreateContainerGroupRequest) SetContainerResourceView(v bool) *CreateContainerGroupRequest {
	s.ContainerResourceView = &v
	return s
}

func (s *CreateContainerGroupRequest) SetCorePattern(v string) *CreateContainerGroupRequest {
	s.CorePattern = &v
	return s
}

func (s *CreateContainerGroupRequest) SetCpu(v float32) *CreateContainerGroupRequest {
	s.Cpu = &v
	return s
}

func (s *CreateContainerGroupRequest) SetCpuOptionsCore(v int32) *CreateContainerGroupRequest {
	s.CpuOptionsCore = &v
	return s
}

func (s *CreateContainerGroupRequest) SetCpuOptionsNuma(v string) *CreateContainerGroupRequest {
	s.CpuOptionsNuma = &v
	return s
}

func (s *CreateContainerGroupRequest) SetCpuOptionsThreadsPerCore(v int32) *CreateContainerGroupRequest {
	s.CpuOptionsThreadsPerCore = &v
	return s
}

func (s *CreateContainerGroupRequest) SetDnsPolicy(v string) *CreateContainerGroupRequest {
	s.DnsPolicy = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEgressBandwidth(v int64) *CreateContainerGroupRequest {
	s.EgressBandwidth = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEipBandwidth(v int32) *CreateContainerGroupRequest {
	s.EipBandwidth = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEipCommonBandwidthPackage(v string) *CreateContainerGroupRequest {
	s.EipCommonBandwidthPackage = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEipISP(v string) *CreateContainerGroupRequest {
	s.EipISP = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEipInstanceId(v string) *CreateContainerGroupRequest {
	s.EipInstanceId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEphemeralStorage(v int32) *CreateContainerGroupRequest {
	s.EphemeralStorage = &v
	return s
}

func (s *CreateContainerGroupRequest) SetHostAliase(v []*CreateContainerGroupRequestHostAliase) *CreateContainerGroupRequest {
	s.HostAliase = v
	return s
}

func (s *CreateContainerGroupRequest) SetHostName(v string) *CreateContainerGroupRequest {
	s.HostName = &v
	return s
}

func (s *CreateContainerGroupRequest) SetImageAccelerateMode(v string) *CreateContainerGroupRequest {
	s.ImageAccelerateMode = &v
	return s
}

func (s *CreateContainerGroupRequest) SetImageRegistryCredential(v []*CreateContainerGroupRequestImageRegistryCredential) *CreateContainerGroupRequest {
	s.ImageRegistryCredential = v
	return s
}

func (s *CreateContainerGroupRequest) SetImageSnapshotId(v string) *CreateContainerGroupRequest {
	s.ImageSnapshotId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetIngressBandwidth(v int64) *CreateContainerGroupRequest {
	s.IngressBandwidth = &v
	return s
}

func (s *CreateContainerGroupRequest) SetInitContainer(v []*CreateContainerGroupRequestInitContainer) *CreateContainerGroupRequest {
	s.InitContainer = v
	return s
}

func (s *CreateContainerGroupRequest) SetInsecureRegistry(v string) *CreateContainerGroupRequest {
	s.InsecureRegistry = &v
	return s
}

func (s *CreateContainerGroupRequest) SetInstanceType(v string) *CreateContainerGroupRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateContainerGroupRequest) SetIpv6AddressCount(v int32) *CreateContainerGroupRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateContainerGroupRequest) SetIpv6GatewayBandwidth(v string) *CreateContainerGroupRequest {
	s.Ipv6GatewayBandwidth = &v
	return s
}

func (s *CreateContainerGroupRequest) SetIpv6GatewayBandwidthEnable(v bool) *CreateContainerGroupRequest {
	s.Ipv6GatewayBandwidthEnable = &v
	return s
}

func (s *CreateContainerGroupRequest) SetMemory(v float32) *CreateContainerGroupRequest {
	s.Memory = &v
	return s
}

func (s *CreateContainerGroupRequest) SetNtpServer(v []*string) *CreateContainerGroupRequest {
	s.NtpServer = v
	return s
}

func (s *CreateContainerGroupRequest) SetOwnerAccount(v string) *CreateContainerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateContainerGroupRequest) SetOwnerId(v int64) *CreateContainerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetPlainHttpRegistry(v string) *CreateContainerGroupRequest {
	s.PlainHttpRegistry = &v
	return s
}

func (s *CreateContainerGroupRequest) SetRamRoleName(v string) *CreateContainerGroupRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateContainerGroupRequest) SetRegionId(v string) *CreateContainerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetResourceGroupId(v string) *CreateContainerGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetResourceOwnerAccount(v string) *CreateContainerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateContainerGroupRequest) SetResourceOwnerId(v int64) *CreateContainerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetRestartPolicy(v string) *CreateContainerGroupRequest {
	s.RestartPolicy = &v
	return s
}

func (s *CreateContainerGroupRequest) SetScheduleStrategy(v string) *CreateContainerGroupRequest {
	s.ScheduleStrategy = &v
	return s
}

func (s *CreateContainerGroupRequest) SetSecurityGroupId(v string) *CreateContainerGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetShareProcessNamespace(v bool) *CreateContainerGroupRequest {
	s.ShareProcessNamespace = &v
	return s
}

func (s *CreateContainerGroupRequest) SetSpotDuration(v int64) *CreateContainerGroupRequest {
	s.SpotDuration = &v
	return s
}

func (s *CreateContainerGroupRequest) SetSpotPriceLimit(v float32) *CreateContainerGroupRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *CreateContainerGroupRequest) SetSpotStrategy(v string) *CreateContainerGroupRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateContainerGroupRequest) SetStrictSpot(v bool) *CreateContainerGroupRequest {
	s.StrictSpot = &v
	return s
}

func (s *CreateContainerGroupRequest) SetTag(v []*CreateContainerGroupRequestTag) *CreateContainerGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateContainerGroupRequest) SetTerminationGracePeriodSeconds(v int64) *CreateContainerGroupRequest {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *CreateContainerGroupRequest) SetVSwitchId(v string) *CreateContainerGroupRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetVolume(v []*CreateContainerGroupRequestVolume) *CreateContainerGroupRequest {
	s.Volume = v
	return s
}

func (s *CreateContainerGroupRequest) SetZoneId(v string) *CreateContainerGroupRequest {
	s.ZoneId = &v
	return s
}

type CreateContainerGroupRequestDnsConfig struct {
	// The IP addresses of the DNS servers.
	NameServer []*string `json:"NameServer,omitempty" xml:"NameServer,omitempty" type:"Repeated"`
	// Configuration options of the DNS server.
	Option []*CreateContainerGroupRequestDnsConfigOption `json:"Option,omitempty" xml:"Option,omitempty" type:"Repeated"`
	// The search domains of the DNS server.
	Search []*string `json:"Search,omitempty" xml:"Search,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestDnsConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestDnsConfig) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestDnsConfig) SetNameServer(v []*string) *CreateContainerGroupRequestDnsConfig {
	s.NameServer = v
	return s
}

func (s *CreateContainerGroupRequestDnsConfig) SetOption(v []*CreateContainerGroupRequestDnsConfigOption) *CreateContainerGroupRequestDnsConfig {
	s.Option = v
	return s
}

func (s *CreateContainerGroupRequestDnsConfig) SetSearch(v []*string) *CreateContainerGroupRequestDnsConfig {
	s.Search = v
	return s
}

type CreateContainerGroupRequestDnsConfigOption struct {
	// The name of the option.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the option.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestDnsConfigOption) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestDnsConfigOption) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestDnsConfigOption) SetName(v string) *CreateContainerGroupRequestDnsConfigOption {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestDnsConfigOption) SetValue(v string) *CreateContainerGroupRequestDnsConfigOption {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestHostSecurityContext struct {
	// Configure a security context to modify unsafe sysctls. For more information, see [Configure a security context](~~462313~~).
	Sysctl []*CreateContainerGroupRequestHostSecurityContextSysctl `json:"Sysctl,omitempty" xml:"Sysctl,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestHostSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestHostSecurityContext) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestHostSecurityContext) SetSysctl(v []*CreateContainerGroupRequestHostSecurityContextSysctl) *CreateContainerGroupRequestHostSecurityContext {
	s.Sysctl = v
	return s
}

type CreateContainerGroupRequestHostSecurityContextSysctl struct {
	// The name of the unsafe sysctl when you configure a security context to modify sysctls. Valid values:
	//
	// *   kernel.shm \* (except for kernel.shm_rmid_forced)
	// *   kernel.msg\*kernel.sem
	// *   fs.mqueue.\*
	// *   net.\* (except for net.ipv4.ip_local_port_range, net.ipv4.tcp_syncookies, net.ipv4.ping_group_range, and net.ipv4.ip_unprivileged_port_start)
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the unsafe sysctl when you configure a security context to modify sysctls.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestHostSecurityContextSysctl) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestHostSecurityContextSysctl) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestHostSecurityContextSysctl) SetName(v string) *CreateContainerGroupRequestHostSecurityContextSysctl {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestHostSecurityContextSysctl) SetValue(v string) *CreateContainerGroupRequestHostSecurityContextSysctl {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestSecurityContext struct {
	// Configure a security context to modify sysctls. For more information, see [Configure a security context](~~462313~~)
	Sysctl []*CreateContainerGroupRequestSecurityContextSysctl `json:"Sysctl,omitempty" xml:"Sysctl,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestSecurityContext) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestSecurityContext) SetSysctl(v []*CreateContainerGroupRequestSecurityContextSysctl) *CreateContainerGroupRequestSecurityContext {
	s.Sysctl = v
	return s
}

type CreateContainerGroupRequestSecurityContextSysctl struct {
	// The name of the safe sysctl when you configure a security context to modify sysctls. Valid values:
	//
	// *   net.ipv4.ping_group_range
	// *   net.ipv4.ip_unprivileged_port_start
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the safe sysctl when you configure a security context to modify sysctls.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestSecurityContextSysctl) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestSecurityContextSysctl) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestSecurityContextSysctl) SetName(v string) *CreateContainerGroupRequestSecurityContextSysctl {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestSecurityContextSysctl) SetValue(v string) *CreateContainerGroupRequestSecurityContextSysctl {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestAcrRegistryInfo struct {
	// The domain names of the Container Registry Enterprise Edition instance. By default, all domain names of the instance are displayed. You can specify one or more domain names. Separate multiple domain names with commas (,).
	Domain []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
	// The ID of the Container Registry Enterprise Edition instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Container Registry Enterprise Edition instance.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID of the Container Registry Enterprise Edition instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateContainerGroupRequestAcrRegistryInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestAcrRegistryInfo) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestAcrRegistryInfo) SetDomain(v []*string) *CreateContainerGroupRequestAcrRegistryInfo {
	s.Domain = v
	return s
}

func (s *CreateContainerGroupRequestAcrRegistryInfo) SetInstanceId(v string) *CreateContainerGroupRequestAcrRegistryInfo {
	s.InstanceId = &v
	return s
}

func (s *CreateContainerGroupRequestAcrRegistryInfo) SetInstanceName(v string) *CreateContainerGroupRequestAcrRegistryInfo {
	s.InstanceName = &v
	return s
}

func (s *CreateContainerGroupRequestAcrRegistryInfo) SetRegionId(v string) *CreateContainerGroupRequestAcrRegistryInfo {
	s.RegionId = &v
	return s
}

type CreateContainerGroupRequestContainer struct {
	LivenessProbe   *CreateContainerGroupRequestContainerLivenessProbe   `json:"LivenessProbe,omitempty" xml:"LivenessProbe,omitempty" require:"true" type:"Struct"`
	ReadinessProbe  *CreateContainerGroupRequestContainerReadinessProbe  `json:"ReadinessProbe,omitempty" xml:"ReadinessProbe,omitempty" require:"true" type:"Struct"`
	SecurityContext *CreateContainerGroupRequestContainerSecurityContext `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" require:"true" type:"Struct"`
	// The arguments that are passed to the startup command of the container. You can specify up to 10 arguments.
	Arg []*string `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	// The command that is to run in the container.
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// The number of vCPUs that you want to allocate to the container.
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The value of the environment variable for the container.
	EnvironmentVar []*CreateContainerGroupRequestContainerEnvironmentVar `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	// The number of GPUs to be allocated to the container.
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The image of the container.
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The policy that you want to use to pull an image. Valid values:
	//
	// *   Always: Each time instances are created, image pulling is performed.
	// *   IfNotPresent: Image pulling is performed as needed. On-premises images are preferentially used. If no on-premises images are available, image pulling is performed.
	// *   Never: On-premises images are always used. Image pulling is not performed.
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The command that you want to run in the container when you use a CLI to specify the postStart callback function.
	LifecyclePostStartHandlerExec []*string `json:"LifecyclePostStartHandlerExec,omitempty" xml:"LifecyclePostStartHandlerExec,omitempty" type:"Repeated"`
	// The IP address of the host that receives HTTP GET requests when you use HTTP requests to specify the postStart callback function.
	LifecyclePostStartHandlerHttpGetHost *string `json:"LifecyclePostStartHandlerHttpGetHost,omitempty" xml:"LifecyclePostStartHandlerHttpGetHost,omitempty"`
	// The HTTP GET request header.
	LifecyclePostStartHandlerHttpGetHttpHeader []*CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader `json:"LifecyclePostStartHandlerHttpGetHttpHeader,omitempty" xml:"LifecyclePostStartHandlerHttpGetHttpHeader,omitempty" type:"Repeated"`
	// The path to which HTTP GET requests are sent when you use HTTP requests to specify the postStart callback function.
	LifecyclePostStartHandlerHttpGetPath *string `json:"LifecyclePostStartHandlerHttpGetPath,omitempty" xml:"LifecyclePostStartHandlerHttpGetPath,omitempty"`
	// The port to which HTTP GET requests are sent when you use HTTP requests to specify the postStart callback function.
	LifecyclePostStartHandlerHttpGetPort *int32 `json:"LifecyclePostStartHandlerHttpGetPort,omitempty" xml:"LifecyclePostStartHandlerHttpGetPort,omitempty"`
	// The protocol type of HTTP GET requests when you use HTTP requests to specify the postStart callback function. Valid values:
	//
	// *   HTTP
	// *   HTTPS
	LifecyclePostStartHandlerHttpGetScheme *string `json:"LifecyclePostStartHandlerHttpGetScheme,omitempty" xml:"LifecyclePostStartHandlerHttpGetScheme,omitempty"`
	// The port that is detected by TCP sockets when you use TCP sockets to specify the postStart callback function.
	LifecyclePostStartHandlerTcpSocketHost *string `json:"LifecyclePostStartHandlerTcpSocketHost,omitempty" xml:"LifecyclePostStartHandlerTcpSocketHost,omitempty"`
	// The port that is detected by TCP sockets when you use TCP sockets to specify the postStart callback function.
	LifecyclePostStartHandlerTcpSocketPort *int32 `json:"LifecyclePostStartHandlerTcpSocketPort,omitempty" xml:"LifecyclePostStartHandlerTcpSocketPort,omitempty"`
	// The command that you want to run in the container when you use a CLI to specify the preStop callback function.
	LifecyclePreStopHandlerExec []*string `json:"LifecyclePreStopHandlerExec,omitempty" xml:"LifecyclePreStopHandlerExec,omitempty" type:"Repeated"`
	// The IP address of the host that receives HTTP GET requests when you use HTTP requests to specify the preStop callback function.
	LifecyclePreStopHandlerHttpGetHost *string `json:"LifecyclePreStopHandlerHttpGetHost,omitempty" xml:"LifecyclePreStopHandlerHttpGetHost,omitempty"`
	// The HTTP GET request header.
	LifecyclePreStopHandlerHttpGetHttpHeader []*CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader `json:"LifecyclePreStopHandlerHttpGetHttpHeader,omitempty" xml:"LifecyclePreStopHandlerHttpGetHttpHeader,omitempty" type:"Repeated"`
	// The path to which HTTP GET requests are sent when you use HTTP requests to specify the preStop callback function.
	LifecyclePreStopHandlerHttpGetPath *string `json:"LifecyclePreStopHandlerHttpGetPath,omitempty" xml:"LifecyclePreStopHandlerHttpGetPath,omitempty"`
	// The port to which HTTP GET requests are sent when you use HTTP requests to specify the preStop callback function.
	LifecyclePreStopHandlerHttpGetPort *int32 `json:"LifecyclePreStopHandlerHttpGetPort,omitempty" xml:"LifecyclePreStopHandlerHttpGetPort,omitempty"`
	// The protocol type of HTTP GET requests when you use HTTP requests to specify the preStop callback function. Valid values:
	//
	// *   HTTP
	// *   HTTPS
	LifecyclePreStopHandlerHttpGetScheme *string `json:"LifecyclePreStopHandlerHttpGetScheme,omitempty" xml:"LifecyclePreStopHandlerHttpGetScheme,omitempty"`
	// The host IP address that is detected by TCP sockets when you use TCP sockets to specify the preStop callback function.
	LifecyclePreStopHandlerTcpSocketHost *string `json:"LifecyclePreStopHandlerTcpSocketHost,omitempty" xml:"LifecyclePreStopHandlerTcpSocketHost,omitempty"`
	// The port that is detected by TCP sockets when you use TCP sockets to specify the preStop callback function.
	LifecyclePreStopHandlerTcpSocketPort *int32 `json:"LifecyclePreStopHandlerTcpSocketPort,omitempty" xml:"LifecyclePreStopHandlerTcpSocketPort,omitempty"`
	// The memory size of the container. Unit: GiB.
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the container.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port to which HTTP GET requests are sent when you use HTTP requests to perform health checks.
	Port []*CreateContainerGroupRequestContainerPort `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
	// Specifies whether the container allocates buffer resources to standard input streams when the container is running. If you do not specify this parameter, an end-of-file (EOF) error may occur when standard input streams in the container are read. Default value: false.
	Stdin *bool `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	// Specifies whether the container runtime closes the stdin channel after the stdin channel has been opened by a single attach session. If stdin is true, the stdin stream remains open across multiple attach sessions.\
	// If StdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and receive data until the client disconnects. When the client disconnects, stdin is closed and remains closed until the container is restarted.
	StdinOnce *bool `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	// The path of the file from which the system retrieves termination messages of the container when the container exits.
	TerminationMessagePath *string `json:"TerminationMessagePath,omitempty" xml:"TerminationMessagePath,omitempty"`
	// The message notification policy. This parameter is empty by default. You can configure only Message Service (MNS) queues to receive notifications.
	TerminationMessagePolicy *string `json:"TerminationMessagePolicy,omitempty" xml:"TerminationMessagePolicy,omitempty"`
	// Specifies whether to enable interaction. Default value: false.
	//
	// If the command is a /bin/bash command, set the value to true.
	Tty *bool `json:"Tty,omitempty" xml:"Tty,omitempty"`
	// Details about the volumes.
	VolumeMount []*CreateContainerGroupRequestContainerVolumeMount `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
	// The working directory of the container.
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s CreateContainerGroupRequestContainer) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainer) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainer) SetLivenessProbe(v *CreateContainerGroupRequestContainerLivenessProbe) *CreateContainerGroupRequestContainer {
	s.LivenessProbe = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetReadinessProbe(v *CreateContainerGroupRequestContainerReadinessProbe) *CreateContainerGroupRequestContainer {
	s.ReadinessProbe = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetSecurityContext(v *CreateContainerGroupRequestContainerSecurityContext) *CreateContainerGroupRequestContainer {
	s.SecurityContext = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetArg(v []*string) *CreateContainerGroupRequestContainer {
	s.Arg = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetCommand(v []*string) *CreateContainerGroupRequestContainer {
	s.Command = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetCpu(v float32) *CreateContainerGroupRequestContainer {
	s.Cpu = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetEnvironmentVar(v []*CreateContainerGroupRequestContainerEnvironmentVar) *CreateContainerGroupRequestContainer {
	s.EnvironmentVar = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetGpu(v int32) *CreateContainerGroupRequestContainer {
	s.Gpu = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetImage(v string) *CreateContainerGroupRequestContainer {
	s.Image = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetImagePullPolicy(v string) *CreateContainerGroupRequestContainer {
	s.ImagePullPolicy = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerExec(v []*string) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerExec = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetHost(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetHost = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetHttpHeader(v []*CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetHttpHeader = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetPath(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetPath = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetPort(v int32) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetPort = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetScheme(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetScheme = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerTcpSocketHost(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerTcpSocketHost = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerTcpSocketPort(v int32) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerTcpSocketPort = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerExec(v []*string) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerExec = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetHost(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetHost = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetHttpHeader(v []*CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetHttpHeader = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetPath(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetPath = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetPort(v int32) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetPort = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetScheme(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetScheme = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerTcpSocketHost(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerTcpSocketHost = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerTcpSocketPort(v int32) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerTcpSocketPort = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetMemory(v float32) *CreateContainerGroupRequestContainer {
	s.Memory = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetName(v string) *CreateContainerGroupRequestContainer {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetPort(v []*CreateContainerGroupRequestContainerPort) *CreateContainerGroupRequestContainer {
	s.Port = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetStdin(v bool) *CreateContainerGroupRequestContainer {
	s.Stdin = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetStdinOnce(v bool) *CreateContainerGroupRequestContainer {
	s.StdinOnce = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetTerminationMessagePath(v string) *CreateContainerGroupRequestContainer {
	s.TerminationMessagePath = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetTerminationMessagePolicy(v string) *CreateContainerGroupRequestContainer {
	s.TerminationMessagePolicy = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetTty(v bool) *CreateContainerGroupRequestContainer {
	s.Tty = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetVolumeMount(v []*CreateContainerGroupRequestContainerVolumeMount) *CreateContainerGroupRequestContainer {
	s.VolumeMount = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetWorkingDir(v string) *CreateContainerGroupRequestContainer {
	s.WorkingDir = &v
	return s
}

type CreateContainerGroupRequestContainerLivenessProbe struct {
	Exec                *CreateContainerGroupRequestContainerLivenessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" require:"true" type:"Struct"`
	FailureThreshold    *int32                                                      `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *CreateContainerGroupRequestContainerLivenessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" require:"true" type:"Struct"`
	InitialDelaySeconds *int32                                                      `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                      `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                      `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *CreateContainerGroupRequestContainerLivenessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" require:"true" type:"Struct"`
	TimeoutSeconds      *int32                                                      `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s CreateContainerGroupRequestContainerLivenessProbe) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLivenessProbe) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetExec(v *CreateContainerGroupRequestContainerLivenessProbeExec) *CreateContainerGroupRequestContainerLivenessProbe {
	s.Exec = v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetFailureThreshold(v int32) *CreateContainerGroupRequestContainerLivenessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetHttpGet(v *CreateContainerGroupRequestContainerLivenessProbeHttpGet) *CreateContainerGroupRequestContainerLivenessProbe {
	s.HttpGet = v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetInitialDelaySeconds(v int32) *CreateContainerGroupRequestContainerLivenessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetPeriodSeconds(v int32) *CreateContainerGroupRequestContainerLivenessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetSuccessThreshold(v int32) *CreateContainerGroupRequestContainerLivenessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetTcpSocket(v *CreateContainerGroupRequestContainerLivenessProbeTcpSocket) *CreateContainerGroupRequestContainerLivenessProbe {
	s.TcpSocket = v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetTimeoutSeconds(v int32) *CreateContainerGroupRequestContainerLivenessProbe {
	s.TimeoutSeconds = &v
	return s
}

type CreateContainerGroupRequestContainerLivenessProbeExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestContainerLivenessProbeExec) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLivenessProbeExec) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLivenessProbeExec) SetCommand(v []*string) *CreateContainerGroupRequestContainerLivenessProbeExec {
	s.Command = v
	return s
}

type CreateContainerGroupRequestContainerLivenessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s CreateContainerGroupRequestContainerLivenessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLivenessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLivenessProbeHttpGet) SetPath(v string) *CreateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbeHttpGet) SetPort(v int32) *CreateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbeHttpGet) SetScheme(v string) *CreateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Scheme = &v
	return s
}

type CreateContainerGroupRequestContainerLivenessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s CreateContainerGroupRequestContainerLivenessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLivenessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLivenessProbeTcpSocket) SetPort(v int32) *CreateContainerGroupRequestContainerLivenessProbeTcpSocket {
	s.Port = &v
	return s
}

type CreateContainerGroupRequestContainerReadinessProbe struct {
	Exec                *CreateContainerGroupRequestContainerReadinessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" require:"true" type:"Struct"`
	FailureThreshold    *int32                                                       `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *CreateContainerGroupRequestContainerReadinessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" require:"true" type:"Struct"`
	InitialDelaySeconds *int32                                                       `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                       `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                       `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *CreateContainerGroupRequestContainerReadinessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" require:"true" type:"Struct"`
	TimeoutSeconds      *int32                                                       `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s CreateContainerGroupRequestContainerReadinessProbe) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerReadinessProbe) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetExec(v *CreateContainerGroupRequestContainerReadinessProbeExec) *CreateContainerGroupRequestContainerReadinessProbe {
	s.Exec = v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetFailureThreshold(v int32) *CreateContainerGroupRequestContainerReadinessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetHttpGet(v *CreateContainerGroupRequestContainerReadinessProbeHttpGet) *CreateContainerGroupRequestContainerReadinessProbe {
	s.HttpGet = v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetInitialDelaySeconds(v int32) *CreateContainerGroupRequestContainerReadinessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetPeriodSeconds(v int32) *CreateContainerGroupRequestContainerReadinessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetSuccessThreshold(v int32) *CreateContainerGroupRequestContainerReadinessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetTcpSocket(v *CreateContainerGroupRequestContainerReadinessProbeTcpSocket) *CreateContainerGroupRequestContainerReadinessProbe {
	s.TcpSocket = v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetTimeoutSeconds(v int32) *CreateContainerGroupRequestContainerReadinessProbe {
	s.TimeoutSeconds = &v
	return s
}

type CreateContainerGroupRequestContainerReadinessProbeExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestContainerReadinessProbeExec) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerReadinessProbeExec) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerReadinessProbeExec) SetCommand(v []*string) *CreateContainerGroupRequestContainerReadinessProbeExec {
	s.Command = v
	return s
}

type CreateContainerGroupRequestContainerReadinessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s CreateContainerGroupRequestContainerReadinessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerReadinessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerReadinessProbeHttpGet) SetPath(v string) *CreateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbeHttpGet) SetPort(v int32) *CreateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbeHttpGet) SetScheme(v string) *CreateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Scheme = &v
	return s
}

type CreateContainerGroupRequestContainerReadinessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s CreateContainerGroupRequestContainerReadinessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerReadinessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerReadinessProbeTcpSocket) SetPort(v int32) *CreateContainerGroupRequestContainerReadinessProbeTcpSocket {
	s.Port = &v
	return s
}

type CreateContainerGroupRequestContainerSecurityContext struct {
	Capability             *CreateContainerGroupRequestContainerSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" require:"true" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                          `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                         `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s CreateContainerGroupRequestContainerSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerSecurityContext) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerSecurityContext) SetCapability(v *CreateContainerGroupRequestContainerSecurityContextCapability) *CreateContainerGroupRequestContainerSecurityContext {
	s.Capability = v
	return s
}

func (s *CreateContainerGroupRequestContainerSecurityContext) SetReadOnlyRootFilesystem(v bool) *CreateContainerGroupRequestContainerSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *CreateContainerGroupRequestContainerSecurityContext) SetRunAsUser(v int64) *CreateContainerGroupRequestContainerSecurityContext {
	s.RunAsUser = &v
	return s
}

type CreateContainerGroupRequestContainerSecurityContextCapability struct {
	Add []*string `json:"Add,omitempty" xml:"Add,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestContainerSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerSecurityContextCapability) SetAdd(v []*string) *CreateContainerGroupRequestContainerSecurityContextCapability {
	s.Add = v
	return s
}

type CreateContainerGroupRequestContainerEnvironmentVar struct {
	FieldRef *CreateContainerGroupRequestContainerEnvironmentVarFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" require:"true" type:"Struct"`
	// The name of the environment variable for the container. The name must be 1 to 128 characters in length. The name can contain letters, digits, and underscores (\_), and cannot start with a digit.``
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable for the container. The value must be 0 to 256 characters in length.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestContainerEnvironmentVar) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerEnvironmentVar) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerEnvironmentVar) SetFieldRef(v *CreateContainerGroupRequestContainerEnvironmentVarFieldRef) *CreateContainerGroupRequestContainerEnvironmentVar {
	s.FieldRef = v
	return s
}

func (s *CreateContainerGroupRequestContainerEnvironmentVar) SetKey(v string) *CreateContainerGroupRequestContainerEnvironmentVar {
	s.Key = &v
	return s
}

func (s *CreateContainerGroupRequestContainerEnvironmentVar) SetValue(v string) *CreateContainerGroupRequestContainerEnvironmentVar {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestContainerEnvironmentVarFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s CreateContainerGroupRequestContainerEnvironmentVarFieldRef) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerEnvironmentVarFieldRef) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerEnvironmentVarFieldRef) SetFieldPath(v string) *CreateContainerGroupRequestContainerEnvironmentVarFieldRef {
	s.FieldPath = &v
	return s
}

type CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader struct {
	// The custom field name in the HTTP GET request header when you use HTTP requests to specify the postStart callback function.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The custom field value in the HTTP GET request header when you use HTTP requests to specify the postStart callback function.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader) SetName(v string) *CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader) SetValue(v string) *CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader struct {
	// The custom field value in the HTTP GET request header when you use HTTP requests to specify the preStop callback function.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The custom field value in the HTTP GET request header when you use HTTP requests to specify the preStop callback function.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) SetName(v string) *CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) SetValue(v string) *CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestContainerPort struct {
	// The port number. Valid values: 1 to 65535.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the protocol. Valid values:
	//
	// *   TCP
	// *   UDP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s CreateContainerGroupRequestContainerPort) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerPort) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerPort) SetPort(v int32) *CreateContainerGroupRequestContainerPort {
	s.Port = &v
	return s
}

func (s *CreateContainerGroupRequestContainerPort) SetProtocol(v string) *CreateContainerGroupRequestContainerPort {
	s.Protocol = &v
	return s
}

type CreateContainerGroupRequestContainerVolumeMount struct {
	// The directory to which the volume is mounted.
	//
	// > The data stored in this directory is overwritten by the data on the volume. Specify this parameter with caution.
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation setting of the volume. Mount propagation allows volumes that are mounted on one container to be shared with other containers in the same pod, or even with other pods on the same node. Valid values:
	//
	// *   None: This volume mount does not receive subsequent mounts that are mounted to this volume or subdirectories of this volume by the host.
	// *   HostToCotainer: This volume mount receives all subsequent mounts that are mounted to this volume or subdirectories of this volume.
	// *   Bidirectional: The volume mount behaves the same as the HostToCotainer mount. The volume mount receives all subsequent mounts that are mounted to this volume or subdirectories of this volume. In addition, all volume mounts created by the container are propagated back to the host and to all containers of all pods that use the same volume.
	//
	// Default value: None.
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The name of the volume. The value of this parameter must be the same as the value of the Name parameter in the volume.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the volume is read-only. Default value: false.
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The subdirectory of the volume.
	SubPath *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s CreateContainerGroupRequestContainerVolumeMount) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerVolumeMount) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerVolumeMount) SetMountPath(v string) *CreateContainerGroupRequestContainerVolumeMount {
	s.MountPath = &v
	return s
}

func (s *CreateContainerGroupRequestContainerVolumeMount) SetMountPropagation(v string) *CreateContainerGroupRequestContainerVolumeMount {
	s.MountPropagation = &v
	return s
}

func (s *CreateContainerGroupRequestContainerVolumeMount) SetName(v string) *CreateContainerGroupRequestContainerVolumeMount {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestContainerVolumeMount) SetReadOnly(v bool) *CreateContainerGroupRequestContainerVolumeMount {
	s.ReadOnly = &v
	return s
}

func (s *CreateContainerGroupRequestContainerVolumeMount) SetSubPath(v string) *CreateContainerGroupRequestContainerVolumeMount {
	s.SubPath = &v
	return s
}

type CreateContainerGroupRequestHostAliase struct {
	// The hostname of the elastic container instance.
	Hostname []*string `json:"Hostname,omitempty" xml:"Hostname,omitempty" type:"Repeated"`
	// The IP address of the host.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s CreateContainerGroupRequestHostAliase) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestHostAliase) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestHostAliase) SetHostname(v []*string) *CreateContainerGroupRequestHostAliase {
	s.Hostname = v
	return s
}

func (s *CreateContainerGroupRequestHostAliase) SetIp(v string) *CreateContainerGroupRequestHostAliase {
	s.Ip = &v
	return s
}

type CreateContainerGroupRequestImageRegistryCredential struct {
	// The password that is used to access the image repository.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The address of the image repository.
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The username that is used to access the image repository.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateContainerGroupRequestImageRegistryCredential) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestImageRegistryCredential) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestImageRegistryCredential) SetPassword(v string) *CreateContainerGroupRequestImageRegistryCredential {
	s.Password = &v
	return s
}

func (s *CreateContainerGroupRequestImageRegistryCredential) SetServer(v string) *CreateContainerGroupRequestImageRegistryCredential {
	s.Server = &v
	return s
}

func (s *CreateContainerGroupRequestImageRegistryCredential) SetUserName(v string) *CreateContainerGroupRequestImageRegistryCredential {
	s.UserName = &v
	return s
}

type CreateContainerGroupRequestInitContainer struct {
	SecurityContext *CreateContainerGroupRequestInitContainerSecurityContext `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" require:"true" type:"Struct"`
	// The arguments that are passed to the container startup command.
	Arg []*string `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	// The command that you want to run to start the container.
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// The number of vCPUs that you want to allocate to the container.
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The environment variables for the container.
	EnvironmentVar []*CreateContainerGroupRequestInitContainerEnvironmentVar `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	// The number of GPUs that you want to allocate to the container.
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The container image.
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image pulling policy. Valid values:
	//
	// *   Always: Image pulling is always performed.
	// *   IfNotPresent: Image pulling is performed as needed. On-premises images are preferentially used. If no on-premises images are available, image pulling is performed.
	// *   Never: On-premises images are always used.
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The memory size of the container. Unit: GiB.
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the container.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port number of the init container.
	Port []*CreateContainerGroupRequestInitContainerPort `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
	// The path of the file from which the system retrieves termination messages of the container when the container exits.
	TerminationMessagePath *string `json:"TerminationMessagePath,omitempty" xml:"TerminationMessagePath,omitempty"`
	// The message notification policy. This parameter is empty by default.
	TerminationMessagePolicy *string `json:"TerminationMessagePolicy,omitempty" xml:"TerminationMessagePolicy,omitempty"`
	// The details about the volume mount.
	VolumeMount []*CreateContainerGroupRequestInitContainerVolumeMount `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
	// The working directory of the container.
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s CreateContainerGroupRequestInitContainer) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainer) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainer) SetSecurityContext(v *CreateContainerGroupRequestInitContainerSecurityContext) *CreateContainerGroupRequestInitContainer {
	s.SecurityContext = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetArg(v []*string) *CreateContainerGroupRequestInitContainer {
	s.Arg = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetCommand(v []*string) *CreateContainerGroupRequestInitContainer {
	s.Command = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetCpu(v float32) *CreateContainerGroupRequestInitContainer {
	s.Cpu = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetEnvironmentVar(v []*CreateContainerGroupRequestInitContainerEnvironmentVar) *CreateContainerGroupRequestInitContainer {
	s.EnvironmentVar = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetGpu(v int32) *CreateContainerGroupRequestInitContainer {
	s.Gpu = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetImage(v string) *CreateContainerGroupRequestInitContainer {
	s.Image = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetImagePullPolicy(v string) *CreateContainerGroupRequestInitContainer {
	s.ImagePullPolicy = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetMemory(v float32) *CreateContainerGroupRequestInitContainer {
	s.Memory = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetName(v string) *CreateContainerGroupRequestInitContainer {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetPort(v []*CreateContainerGroupRequestInitContainerPort) *CreateContainerGroupRequestInitContainer {
	s.Port = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetTerminationMessagePath(v string) *CreateContainerGroupRequestInitContainer {
	s.TerminationMessagePath = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetTerminationMessagePolicy(v string) *CreateContainerGroupRequestInitContainer {
	s.TerminationMessagePolicy = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetVolumeMount(v []*CreateContainerGroupRequestInitContainerVolumeMount) *CreateContainerGroupRequestInitContainer {
	s.VolumeMount = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetWorkingDir(v string) *CreateContainerGroupRequestInitContainer {
	s.WorkingDir = &v
	return s
}

type CreateContainerGroupRequestInitContainerSecurityContext struct {
	Capability             *CreateContainerGroupRequestInitContainerSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" require:"true" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                              `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                             `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s CreateContainerGroupRequestInitContainerSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerSecurityContext) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerSecurityContext) SetCapability(v *CreateContainerGroupRequestInitContainerSecurityContextCapability) *CreateContainerGroupRequestInitContainerSecurityContext {
	s.Capability = v
	return s
}

func (s *CreateContainerGroupRequestInitContainerSecurityContext) SetReadOnlyRootFilesystem(v bool) *CreateContainerGroupRequestInitContainerSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerSecurityContext) SetRunAsUser(v int64) *CreateContainerGroupRequestInitContainerSecurityContext {
	s.RunAsUser = &v
	return s
}

type CreateContainerGroupRequestInitContainerSecurityContextCapability struct {
	Add []*string `json:"Add,omitempty" xml:"Add,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestInitContainerSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerSecurityContextCapability) SetAdd(v []*string) *CreateContainerGroupRequestInitContainerSecurityContextCapability {
	s.Add = v
	return s
}

type CreateContainerGroupRequestInitContainerEnvironmentVar struct {
	FieldRef *CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" require:"true" type:"Struct"`
	// The name of the environment variable. The name must be 1 to 128 characters in length. The name can contain letters, digits, and underscores (\_), and cannot start with a digit.``
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable. The value must be 0 to 256 characters in length.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestInitContainerEnvironmentVar) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerEnvironmentVar) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerEnvironmentVar) SetFieldRef(v *CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef) *CreateContainerGroupRequestInitContainerEnvironmentVar {
	s.FieldRef = v
	return s
}

func (s *CreateContainerGroupRequestInitContainerEnvironmentVar) SetKey(v string) *CreateContainerGroupRequestInitContainerEnvironmentVar {
	s.Key = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerEnvironmentVar) SetValue(v string) *CreateContainerGroupRequestInitContainerEnvironmentVar {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef) SetFieldPath(v string) *CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef {
	s.FieldPath = &v
	return s
}

type CreateContainerGroupRequestInitContainerPort struct {
	// The port number. Valid values: 1 to 65535.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the protocol. Valid values:
	//
	// *   TCP
	// *   UDP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s CreateContainerGroupRequestInitContainerPort) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerPort) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerPort) SetPort(v int32) *CreateContainerGroupRequestInitContainerPort {
	s.Port = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerPort) SetProtocol(v string) *CreateContainerGroupRequestInitContainerPort {
	s.Protocol = &v
	return s
}

type CreateContainerGroupRequestInitContainerVolumeMount struct {
	// The data stored in this directory is overwritten by the data on the volume. Specify this parameter with caution.
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation setting of the volume. Mount propagation allows volumes that are mounted on one container to be shared with other containers in the same pod, or even with other pods on the same node. Valid values:
	//
	// *   None: This volume mount does not receive subsequent mounts that are mounted to this volume or subdirectories of this volume by the host.
	// *   HostToCotainer: This volume mount receives all subsequent mounts that are mounted to this volume or subdirectories of this volume.
	// *   Bidirectional: The volume mount behaves the same as the HostToCotainer mount. The volume mount receives all subsequent mounts that are mounted to this volume or subdirectories of this volume. In addition, all volume mounts created by the container are propagated back to the host and to all containers of all pods that use the same volume.
	//
	// Default value: None.
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The name of the volume that you want to mount.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the mount path is read-only. Default value: false.
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The subdirectory of the volume. The elastic container instance can mount different directories of the same volume to different subdirectories of containers.
	SubPath *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s CreateContainerGroupRequestInitContainerVolumeMount) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerVolumeMount) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerVolumeMount) SetMountPath(v string) *CreateContainerGroupRequestInitContainerVolumeMount {
	s.MountPath = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerVolumeMount) SetMountPropagation(v string) *CreateContainerGroupRequestInitContainerVolumeMount {
	s.MountPropagation = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerVolumeMount) SetName(v string) *CreateContainerGroupRequestInitContainerVolumeMount {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerVolumeMount) SetReadOnly(v bool) *CreateContainerGroupRequestInitContainerVolumeMount {
	s.ReadOnly = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerVolumeMount) SetSubPath(v string) *CreateContainerGroupRequestInitContainerVolumeMount {
	s.SubPath = &v
	return s
}

type CreateContainerGroupRequestTag struct {
	// The key of a tag. The tag key cannot be an empty string and must be unique. The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of a tag. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `acs:`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestTag) SetKey(v string) *CreateContainerGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateContainerGroupRequestTag) SetValue(v string) *CreateContainerGroupRequestTag {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestVolume struct {
	ConfigFileVolume *CreateContainerGroupRequestVolumeConfigFileVolume `json:"ConfigFileVolume,omitempty" xml:"ConfigFileVolume,omitempty" require:"true" type:"Struct"`
	DiskVolume       *CreateContainerGroupRequestVolumeDiskVolume       `json:"DiskVolume,omitempty" xml:"DiskVolume,omitempty" require:"true" type:"Struct"`
	EmptyDirVolume   *CreateContainerGroupRequestVolumeEmptyDirVolume   `json:"EmptyDirVolume,omitempty" xml:"EmptyDirVolume,omitempty" require:"true" type:"Struct"`
	FlexVolume       *CreateContainerGroupRequestVolumeFlexVolume       `json:"FlexVolume,omitempty" xml:"FlexVolume,omitempty" require:"true" type:"Struct"`
	HostPathVolume   *CreateContainerGroupRequestVolumeHostPathVolume   `json:"HostPathVolume,omitempty" xml:"HostPathVolume,omitempty" require:"true" type:"Struct"`
	NFSVolume        *CreateContainerGroupRequestVolumeNFSVolume        `json:"NFSVolume,omitempty" xml:"NFSVolume,omitempty" require:"true" type:"Struct"`
	// The name of the volume.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The category of the HostPath volume. Valid values:
	//
	// *   Directory
	// *   File
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateContainerGroupRequestVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolume) SetConfigFileVolume(v *CreateContainerGroupRequestVolumeConfigFileVolume) *CreateContainerGroupRequestVolume {
	s.ConfigFileVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetDiskVolume(v *CreateContainerGroupRequestVolumeDiskVolume) *CreateContainerGroupRequestVolume {
	s.DiskVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetEmptyDirVolume(v *CreateContainerGroupRequestVolumeEmptyDirVolume) *CreateContainerGroupRequestVolume {
	s.EmptyDirVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetFlexVolume(v *CreateContainerGroupRequestVolumeFlexVolume) *CreateContainerGroupRequestVolume {
	s.FlexVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetHostPathVolume(v *CreateContainerGroupRequestVolumeHostPathVolume) *CreateContainerGroupRequestVolume {
	s.HostPathVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetNFSVolume(v *CreateContainerGroupRequestVolumeNFSVolume) *CreateContainerGroupRequestVolume {
	s.NFSVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetName(v string) *CreateContainerGroupRequestVolume {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetType(v string) *CreateContainerGroupRequestVolume {
	s.Type = &v
	return s
}

type CreateContainerGroupRequestVolumeConfigFileVolume struct {
	ConfigFileToPath []*CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath `json:"ConfigFileToPath,omitempty" xml:"ConfigFileToPath,omitempty" type:"Repeated"`
	DefaultMode      *int32                                                               `json:"DefaultMode,omitempty" xml:"DefaultMode,omitempty"`
}

func (s CreateContainerGroupRequestVolumeConfigFileVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeConfigFileVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeConfigFileVolume) SetConfigFileToPath(v []*CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) *CreateContainerGroupRequestVolumeConfigFileVolume {
	s.ConfigFileToPath = v
	return s
}

func (s *CreateContainerGroupRequestVolumeConfigFileVolume) SetDefaultMode(v int32) *CreateContainerGroupRequestVolumeConfigFileVolume {
	s.DefaultMode = &v
	return s
}

type CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Mode    *int32  `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Path    *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) SetContent(v string) *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath {
	s.Content = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) SetMode(v int32) *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath {
	s.Mode = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) SetPath(v string) *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath {
	s.Path = &v
	return s
}

type CreateContainerGroupRequestVolumeDiskVolume struct {
	DiskId   *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskSize *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	FsType   *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
}

func (s CreateContainerGroupRequestVolumeDiskVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeDiskVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeDiskVolume) SetDiskId(v string) *CreateContainerGroupRequestVolumeDiskVolume {
	s.DiskId = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeDiskVolume) SetDiskSize(v int32) *CreateContainerGroupRequestVolumeDiskVolume {
	s.DiskSize = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeDiskVolume) SetFsType(v string) *CreateContainerGroupRequestVolumeDiskVolume {
	s.FsType = &v
	return s
}

type CreateContainerGroupRequestVolumeEmptyDirVolume struct {
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	SizeLimit *string `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s CreateContainerGroupRequestVolumeEmptyDirVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeEmptyDirVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeEmptyDirVolume) SetMedium(v string) *CreateContainerGroupRequestVolumeEmptyDirVolume {
	s.Medium = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeEmptyDirVolume) SetSizeLimit(v string) *CreateContainerGroupRequestVolumeEmptyDirVolume {
	s.SizeLimit = &v
	return s
}

type CreateContainerGroupRequestVolumeFlexVolume struct {
	Driver  *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	FsType  *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s CreateContainerGroupRequestVolumeFlexVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeFlexVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeFlexVolume) SetDriver(v string) *CreateContainerGroupRequestVolumeFlexVolume {
	s.Driver = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeFlexVolume) SetFsType(v string) *CreateContainerGroupRequestVolumeFlexVolume {
	s.FsType = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeFlexVolume) SetOptions(v string) *CreateContainerGroupRequestVolumeFlexVolume {
	s.Options = &v
	return s
}

type CreateContainerGroupRequestVolumeHostPathVolume struct {
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateContainerGroupRequestVolumeHostPathVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeHostPathVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeHostPathVolume) SetPath(v string) *CreateContainerGroupRequestVolumeHostPathVolume {
	s.Path = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeHostPathVolume) SetType(v string) *CreateContainerGroupRequestVolumeHostPathVolume {
	s.Type = &v
	return s
}

type CreateContainerGroupRequestVolumeNFSVolume struct {
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ReadOnly *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
}

func (s CreateContainerGroupRequestVolumeNFSVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeNFSVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeNFSVolume) SetPath(v string) *CreateContainerGroupRequestVolumeNFSVolume {
	s.Path = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeNFSVolume) SetReadOnly(v bool) *CreateContainerGroupRequestVolumeNFSVolume {
	s.ReadOnly = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeNFSVolume) SetServer(v string) *CreateContainerGroupRequestVolumeNFSVolume {
	s.Server = &v
	return s
}

type CreateContainerGroupResponseBody struct {
	// The ID of the elastic container instance.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateContainerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupResponseBody) SetContainerGroupId(v string) *CreateContainerGroupResponseBody {
	s.ContainerGroupId = &v
	return s
}

func (s *CreateContainerGroupResponseBody) SetRequestId(v string) *CreateContainerGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateContainerGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateContainerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateContainerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupResponse) SetHeaders(v map[string]*string) *CreateContainerGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateContainerGroupResponse) SetStatusCode(v int32) *CreateContainerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContainerGroupResponse) SetBody(v *CreateContainerGroupResponseBody) *CreateContainerGroupResponse {
	s.Body = v
	return s
}

type CreateImageCacheRequest struct {
	AcrRegistryInfo []*CreateImageCacheRequestAcrRegistryInfo `json:"AcrRegistryInfo,omitempty" xml:"AcrRegistryInfo,omitempty" type:"Repeated"`
	// Comments.
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// Specifies whether to enable reuse of image cache layers. If you enable this feature, and the image cache that you want to create and an existing image cache contain duplicate image layers, the system reuses the duplicate image layers to create the new image cache. This accelerates the creation of the image cache. Valid values:
	//
	// *   true: enables reuse of image cache layers.
	// *   false: disables reuse of image cache layers.
	//
	// Default value: false.
	AutoMatchImageCache *bool `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the elastic IP address (EIP). If you want to pull images over the Internet, make sure that the elastic container instance can access the Internet. You can configure an EIP or a NAT gateway for the elastic container instance to access the Internet.
	EipInstanceId *string `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	// The elimination policy of the image cache. This parameter is empty by default, which indicates that the image cache is always retained.
	//
	// You can set this parameter to LRU, which indicates that the image cache can be automatically deleted. When the number of image caches reaches the specified quota, the system automatically deletes the image caches whose EliminationStrategy parameter is set to LRU and that are least commonly used.
	EliminationStrategy *string `json:"EliminationStrategy,omitempty" xml:"EliminationStrategy,omitempty"`
	// Specifies whether to enable instant image cache. After you enable the feature, image caches can be created in an accelerated manner. Valid values:
	//
	// *   true: enables instant image cache.
	// *   false: disables instant image cache.
	//
	// Default value: false.
	//
	// >  The system automatically generates a temporary local snapshot for the image cache during the use of the instant image cache feature. You are charged for the instant use of the snapshot.
	Flash *bool `json:"Flash,omitempty" xml:"Flash,omitempty"`
	// The number of temporary local snapshots. By default, the system creates one snapshot for each image cache. If an image cache is used to create multiple elastic container instances at a time, we recommend that you set this parameter to create multiple snapshots for the image cache. We recommend that you create one snapshot for creation of every 1,000 elastic container instances.
	//
	// >  If you set the Flash parameter to true, instant image cache is enabled. During the creation of the image cache, the system first creates a temporary local snapshot for you to instantly use the snapshot. After the temporary local snapshot is created, the system begins to create a regular snapshot. After the regular snapshot is created, the temporary local snapshot is automatically deleted.
	FlashCopyCount *int32    `json:"FlashCopyCount,omitempty" xml:"FlashCopyCount,omitempty"`
	Image          []*string `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
	// The name of the image cache.
	ImageCacheName *string `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	// The size of the image cache. Unit: GiB. Default value: 20.
	ImageCacheSize          *int32                                            `json:"ImageCacheSize,omitempty" xml:"ImageCacheSize,omitempty"`
	ImageRegistryCredential []*CreateImageCacheRequestImageRegistryCredential `json:"ImageRegistryCredential,omitempty" xml:"ImageRegistryCredential,omitempty" type:"Repeated"`
	// The domain name of the self-managed image repository.
	//
	// When you create an image cache by using an image in a self-managed image repository that uses a self-signed certificate, you must specify this parameter to skip the certificate authentication. This can prevent the image from failing to pull due to certificate authentication failures.
	InsecureRegistry *string `json:"InsecureRegistry,omitempty" xml:"InsecureRegistry,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The domain name of the self-managed image repository. When you create an image cache by using an image in a self-managed image repository that uses the HTTP Protocol, you must specify this parameter. This way, Elastic Container Instance uses the HTTP protocol to pull the image. This can prevent the image from failing to pull due to different protocols.
	PlainHttpRegistry *string `json:"PlainHttpRegistry,omitempty" xml:"PlainHttpRegistry,omitempty"`
	// The region ID of the image cache.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The retention period of the image cache. Unit: days. When the retention period ends, the image cache expires and is deleted. By default, image caches never expire.
	//
	// >  The image caches that fail to be created are only retained for one day.
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The ID of the security group.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The number of regular snapshots. By default, the system creates one snapshot for each image cache. If an image cache is used to create multiple elastic container instances at a time, we recommend that you set this parameter to create multiple snapshots for the image cache. We recommend that you create one snapshot for creation of every 1,000 elastic container instances.
	//
	// >  If you set the Flash parameter to false, instant image cache is disabled. In this case, only regular snapshots are generated during the creation of the image cache.
	StandardCopyCount *int32                        `json:"StandardCopyCount,omitempty" xml:"StandardCopyCount,omitempty"`
	Tag               []*CreateImageCacheRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch. You can specify up to 10 vSwitch IDs. Separate multiple vSwitch IDs with commas (,). Example: `vsw-***,vsw-***`.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the image cache.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateImageCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheRequest) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequest) SetAcrRegistryInfo(v []*CreateImageCacheRequestAcrRegistryInfo) *CreateImageCacheRequest {
	s.AcrRegistryInfo = v
	return s
}

func (s *CreateImageCacheRequest) SetAnnotations(v string) *CreateImageCacheRequest {
	s.Annotations = &v
	return s
}

func (s *CreateImageCacheRequest) SetAutoMatchImageCache(v bool) *CreateImageCacheRequest {
	s.AutoMatchImageCache = &v
	return s
}

func (s *CreateImageCacheRequest) SetClientToken(v string) *CreateImageCacheRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageCacheRequest) SetEipInstanceId(v string) *CreateImageCacheRequest {
	s.EipInstanceId = &v
	return s
}

func (s *CreateImageCacheRequest) SetEliminationStrategy(v string) *CreateImageCacheRequest {
	s.EliminationStrategy = &v
	return s
}

func (s *CreateImageCacheRequest) SetFlash(v bool) *CreateImageCacheRequest {
	s.Flash = &v
	return s
}

func (s *CreateImageCacheRequest) SetFlashCopyCount(v int32) *CreateImageCacheRequest {
	s.FlashCopyCount = &v
	return s
}

func (s *CreateImageCacheRequest) SetImage(v []*string) *CreateImageCacheRequest {
	s.Image = v
	return s
}

func (s *CreateImageCacheRequest) SetImageCacheName(v string) *CreateImageCacheRequest {
	s.ImageCacheName = &v
	return s
}

func (s *CreateImageCacheRequest) SetImageCacheSize(v int32) *CreateImageCacheRequest {
	s.ImageCacheSize = &v
	return s
}

func (s *CreateImageCacheRequest) SetImageRegistryCredential(v []*CreateImageCacheRequestImageRegistryCredential) *CreateImageCacheRequest {
	s.ImageRegistryCredential = v
	return s
}

func (s *CreateImageCacheRequest) SetInsecureRegistry(v string) *CreateImageCacheRequest {
	s.InsecureRegistry = &v
	return s
}

func (s *CreateImageCacheRequest) SetOwnerAccount(v string) *CreateImageCacheRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateImageCacheRequest) SetOwnerId(v int64) *CreateImageCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateImageCacheRequest) SetPlainHttpRegistry(v string) *CreateImageCacheRequest {
	s.PlainHttpRegistry = &v
	return s
}

func (s *CreateImageCacheRequest) SetRegionId(v string) *CreateImageCacheRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageCacheRequest) SetResourceGroupId(v string) *CreateImageCacheRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateImageCacheRequest) SetResourceOwnerAccount(v string) *CreateImageCacheRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateImageCacheRequest) SetResourceOwnerId(v int64) *CreateImageCacheRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateImageCacheRequest) SetRetentionDays(v int32) *CreateImageCacheRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateImageCacheRequest) SetSecurityGroupId(v string) *CreateImageCacheRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateImageCacheRequest) SetStandardCopyCount(v int32) *CreateImageCacheRequest {
	s.StandardCopyCount = &v
	return s
}

func (s *CreateImageCacheRequest) SetTag(v []*CreateImageCacheRequestTag) *CreateImageCacheRequest {
	s.Tag = v
	return s
}

func (s *CreateImageCacheRequest) SetVSwitchId(v string) *CreateImageCacheRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateImageCacheRequest) SetZoneId(v string) *CreateImageCacheRequest {
	s.ZoneId = &v
	return s
}

type CreateImageCacheRequestAcrRegistryInfo struct {
	Domain []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
	// The ID of Container Registry Enterprise Edition instance N.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of Container Registry Enterprise Edition instance N.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID of Container Registry Enterprise Edition instance N.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateImageCacheRequestAcrRegistryInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheRequestAcrRegistryInfo) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequestAcrRegistryInfo) SetDomain(v []*string) *CreateImageCacheRequestAcrRegistryInfo {
	s.Domain = v
	return s
}

func (s *CreateImageCacheRequestAcrRegistryInfo) SetInstanceId(v string) *CreateImageCacheRequestAcrRegistryInfo {
	s.InstanceId = &v
	return s
}

func (s *CreateImageCacheRequestAcrRegistryInfo) SetInstanceName(v string) *CreateImageCacheRequestAcrRegistryInfo {
	s.InstanceName = &v
	return s
}

func (s *CreateImageCacheRequestAcrRegistryInfo) SetRegionId(v string) *CreateImageCacheRequestAcrRegistryInfo {
	s.RegionId = &v
	return s
}

type CreateImageCacheRequestImageRegistryCredential struct {
	// The password that is used to log on to image repository N.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The endpoint of the image repository without the `http://` or `https://` prefix.
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The username that is used to log on to image repository N.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateImageCacheRequestImageRegistryCredential) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheRequestImageRegistryCredential) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequestImageRegistryCredential) SetPassword(v string) *CreateImageCacheRequestImageRegistryCredential {
	s.Password = &v
	return s
}

func (s *CreateImageCacheRequestImageRegistryCredential) SetServer(v string) *CreateImageCacheRequestImageRegistryCredential {
	s.Server = &v
	return s
}

func (s *CreateImageCacheRequestImageRegistryCredential) SetUserName(v string) *CreateImageCacheRequestImageRegistryCredential {
	s.UserName = &v
	return s
}

type CreateImageCacheRequestTag struct {
	// The key of tag N of the image cache. Valid values of N: 1 to 20.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the image cache. Valid values of N: 1 to 20.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateImageCacheRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheRequestTag) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequestTag) SetKey(v string) *CreateImageCacheRequestTag {
	s.Key = &v
	return s
}

func (s *CreateImageCacheRequestTag) SetValue(v string) *CreateImageCacheRequestTag {
	s.Value = &v
	return s
}

type CreateImageCacheResponseBody struct {
	// The ID of the intermediate elastic container instance that is used to create the image cache.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The ID of the image cache.
	ImageCacheId *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageCacheResponseBody) SetContainerGroupId(v string) *CreateImageCacheResponseBody {
	s.ContainerGroupId = &v
	return s
}

func (s *CreateImageCacheResponseBody) SetImageCacheId(v string) *CreateImageCacheResponseBody {
	s.ImageCacheId = &v
	return s
}

func (s *CreateImageCacheResponseBody) SetRequestId(v string) *CreateImageCacheResponseBody {
	s.RequestId = &v
	return s
}

type CreateImageCacheResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateImageCacheResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateImageCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheResponse) GoString() string {
	return s.String()
}

func (s *CreateImageCacheResponse) SetHeaders(v map[string]*string) *CreateImageCacheResponse {
	s.Headers = v
	return s
}

func (s *CreateImageCacheResponse) SetStatusCode(v int32) *CreateImageCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageCacheResponse) SetBody(v *CreateImageCacheResponseBody) *CreateImageCacheResponse {
	s.Body = v
	return s
}

type CreateInstanceOpsTaskRequest struct {
	// The ID of the container group.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The type of the O&M task. Valid values:
	//
	// *   coredump
	// *   tcpdump
	OpsType *string `json:"OpsType,omitempty" xml:"OpsType,omitempty"`
	// The value of the O\&M task. You can set this parameter based on the value of OpsType.
	//
	// *   If OpsType is set to coredump, the valid values of OpsValue are:
	//
	//     *   enable: enables coredump.
	//     *   disable: disables coredump.
	//
	// *   If OpsType is set to tcpdump, the value of OpsValue is a JSON-formatted parameter string. Example: `{"Enable":true, "IfDeviceName":"eth0"}`. The value may contain the following parameters:
	//
	//     *   Enable: specifies whether to enable tcpdump. You must specify this parameter. Valid values: true and false.
	//     *   Protocol: the network protocol. Valid values: tcp, udp, and icmpv4.
	//     *   SourceIp: the source IP address.
	//     *   SourceCidr: the source CIDR block. If you specify both an IP address and a CIDR block, the IP address is ignored if the CIDR block is valid.
	//     *   SourcePort: the source port. Valid values: 1 to 65535.
	//     *   DestIp: the destination IP address.
	//     *   DestCidr: the destination CIDR block. If you specify both an IP address and a CIDR block, the IP address is ignored if the CIDR block is valid.
	//     *   DestPort: the destination port. Valid values: 1 to 65535.
	//     *   IfDeviceName: the destination network interface controller. Default value: eth0.
	//     *   Snaplen: the length to be captured. Default value: 65535. Unit: bytes.
	//     *   Duration: the captured period. Unit: seconds.
	//     *   PacketNum: the number of packets to be captured.
	//     *   FileSize: the size of the destination files on packets. Unit: bytes. Maximum value: 1073741824. 1073741824 bytes is equal to 1 GB.
	OpsValue     *string `json:"OpsValue,omitempty" xml:"OpsValue,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the O&M task.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateInstanceOpsTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceOpsTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceOpsTaskRequest) SetContainerGroupId(v string) *CreateInstanceOpsTaskRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *CreateInstanceOpsTaskRequest) SetOpsType(v string) *CreateInstanceOpsTaskRequest {
	s.OpsType = &v
	return s
}

func (s *CreateInstanceOpsTaskRequest) SetOpsValue(v string) *CreateInstanceOpsTaskRequest {
	s.OpsValue = &v
	return s
}

func (s *CreateInstanceOpsTaskRequest) SetOwnerAccount(v string) *CreateInstanceOpsTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateInstanceOpsTaskRequest) SetOwnerId(v int64) *CreateInstanceOpsTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateInstanceOpsTaskRequest) SetRegionId(v string) *CreateInstanceOpsTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceOpsTaskRequest) SetResourceOwnerAccount(v string) *CreateInstanceOpsTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateInstanceOpsTaskRequest) SetResourceOwnerId(v int64) *CreateInstanceOpsTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateInstanceOpsTaskResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The creation result of the O&M task.
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateInstanceOpsTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceOpsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceOpsTaskResponseBody) SetRequestId(v string) *CreateInstanceOpsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceOpsTaskResponseBody) SetResult(v string) *CreateInstanceOpsTaskResponseBody {
	s.Result = &v
	return s
}

type CreateInstanceOpsTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceOpsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceOpsTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceOpsTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceOpsTaskResponse) SetHeaders(v map[string]*string) *CreateInstanceOpsTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceOpsTaskResponse) SetStatusCode(v int32) *CreateInstanceOpsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceOpsTaskResponse) SetBody(v *CreateInstanceOpsTaskResponseBody) *CreateInstanceOpsTaskResponse {
	s.Body = v
	return s
}

type CreateVirtualNodeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterDNS      *string `json:"ClusterDNS,omitempty" xml:"ClusterDNS,omitempty"`
	ClusterDomain   *string `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	CustomResources *string `json:"CustomResources,omitempty" xml:"CustomResources,omitempty"`
	// The ID of the elastic IP address (EIP).
	EipInstanceId *string `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	// Specifies whether to enable Internet access for the VNode. Default value: false.
	//
	// If the value of this parameter is true, the VNode exposes a public IP address to external services.
	EnablePublicNetwork *bool `json:"EnablePublicNetwork,omitempty" xml:"EnablePublicNetwork,omitempty"`
	// KubeConfig of the Kubernetes cluster to which the VNode is to be connected. The value must be Base64-encoded.
	KubeConfig   *string `json:"KubeConfig,omitempty" xml:"KubeConfig,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VNode.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable certificate rotation. If you set the value to true, the system automatically applies for a new certificate for the VNode when the current certificate is about to expire. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false.
	RotateCertificateEnabled *bool `json:"RotateCertificateEnabled,omitempty" xml:"RotateCertificateEnabled,omitempty"`
	// The ID of the security group. The VNode and the elastic container instances in the VNode are added to the security group.
	SecurityGroupId *string                          `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag             []*CreateVirtualNodeRequestTag   `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	Taint           []*CreateVirtualNodeRequestTaint `json:"Taint,omitempty" xml:"Taint,omitempty" type:"Repeated"`
	// Specifies whether to enable TLS bootstrapping. If you set this parameter to true, use the KubeConfig certificate for TLS bootstrapping. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false.
	TlsBootstrapEnabled *bool `json:"TlsBootstrapEnabled,omitempty" xml:"TlsBootstrapEnabled,omitempty"`
	// The ID of the vSwitch. The vSwitch is connected to the VNode and the elastic container instances in the VNode.
	//
	// You can specify 1 to 10 vSwitches for a VPC.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the VNode. The name must be 2 to 128 characters in length, and can contain lowercase letters, digits, periods (.), and hyphens (-).
	VirtualNodeName *string `json:"VirtualNodeName,omitempty" xml:"VirtualNodeName,omitempty"`
	// The zone ID of the VNode.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateVirtualNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualNodeRequest) GoString() string {
	return s.String()
}

func (s *CreateVirtualNodeRequest) SetClientToken(v string) *CreateVirtualNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetClusterDNS(v string) *CreateVirtualNodeRequest {
	s.ClusterDNS = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetClusterDomain(v string) *CreateVirtualNodeRequest {
	s.ClusterDomain = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetCustomResources(v string) *CreateVirtualNodeRequest {
	s.CustomResources = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetEipInstanceId(v string) *CreateVirtualNodeRequest {
	s.EipInstanceId = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetEnablePublicNetwork(v bool) *CreateVirtualNodeRequest {
	s.EnablePublicNetwork = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetKubeConfig(v string) *CreateVirtualNodeRequest {
	s.KubeConfig = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetOwnerAccount(v string) *CreateVirtualNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetOwnerId(v int64) *CreateVirtualNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetRegionId(v string) *CreateVirtualNodeRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetResourceGroupId(v string) *CreateVirtualNodeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetResourceOwnerAccount(v string) *CreateVirtualNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetResourceOwnerId(v int64) *CreateVirtualNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetRotateCertificateEnabled(v bool) *CreateVirtualNodeRequest {
	s.RotateCertificateEnabled = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetSecurityGroupId(v string) *CreateVirtualNodeRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetTag(v []*CreateVirtualNodeRequestTag) *CreateVirtualNodeRequest {
	s.Tag = v
	return s
}

func (s *CreateVirtualNodeRequest) SetTaint(v []*CreateVirtualNodeRequestTaint) *CreateVirtualNodeRequest {
	s.Taint = v
	return s
}

func (s *CreateVirtualNodeRequest) SetTlsBootstrapEnabled(v bool) *CreateVirtualNodeRequest {
	s.TlsBootstrapEnabled = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetVSwitchId(v string) *CreateVirtualNodeRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetVirtualNodeName(v string) *CreateVirtualNodeRequest {
	s.VirtualNodeName = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetZoneId(v string) *CreateVirtualNodeRequest {
	s.ZoneId = &v
	return s
}

type CreateVirtualNodeRequestTag struct {
	// The key of the tag. Valid values of N: 1 to 20.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. Valid values of N: 1 to 20.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVirtualNodeRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualNodeRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVirtualNodeRequestTag) SetKey(v string) *CreateVirtualNodeRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVirtualNodeRequestTag) SetValue(v string) *CreateVirtualNodeRequestTag {
	s.Value = &v
	return s
}

type CreateVirtualNodeRequestTaint struct {
	// The effect of taint N. Valid values of N: 1 to 20. Valid values:
	//
	// *   NoSchedule: No pods are scheduled to the nodes that have the taint.
	// *   NoExecute: Existing pods in the node are evicted while no pods are scheduled to the nodes that have the taint.
	// *   PreferNoSchedule: Pods are preferentially not scheduled to the nodes that have the taint.
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// The key of taint N. Valid values of N: 1 to 20.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of taint N. Valid values of N: 1 to 20.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVirtualNodeRequestTaint) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualNodeRequestTaint) GoString() string {
	return s.String()
}

func (s *CreateVirtualNodeRequestTaint) SetEffect(v string) *CreateVirtualNodeRequestTaint {
	s.Effect = &v
	return s
}

func (s *CreateVirtualNodeRequestTaint) SetKey(v string) *CreateVirtualNodeRequestTaint {
	s.Key = &v
	return s
}

func (s *CreateVirtualNodeRequestTaint) SetValue(v string) *CreateVirtualNodeRequestTaint {
	s.Value = &v
	return s
}

type CreateVirtualNodeResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the VNode.
	VirtualNodeId *string `json:"VirtualNodeId,omitempty" xml:"VirtualNodeId,omitempty"`
}

func (s CreateVirtualNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirtualNodeResponseBody) SetRequestId(v string) *CreateVirtualNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirtualNodeResponseBody) SetVirtualNodeId(v string) *CreateVirtualNodeResponseBody {
	s.VirtualNodeId = &v
	return s
}

type CreateVirtualNodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVirtualNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVirtualNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualNodeResponse) GoString() string {
	return s.String()
}

func (s *CreateVirtualNodeResponse) SetHeaders(v map[string]*string) *CreateVirtualNodeResponse {
	s.Headers = v
	return s
}

func (s *CreateVirtualNodeResponse) SetStatusCode(v int32) *CreateVirtualNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVirtualNodeResponse) SetBody(v *CreateVirtualNodeResponseBody) *CreateVirtualNodeResponse {
	s.Body = v
	return s
}

type DeleteContainerGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the elastic container instance.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the elastic container instance.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteContainerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteContainerGroupRequest) SetClientToken(v string) *DeleteContainerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetContainerGroupId(v string) *DeleteContainerGroupRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetOwnerAccount(v string) *DeleteContainerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetOwnerId(v int64) *DeleteContainerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetRegionId(v string) *DeleteContainerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetResourceOwnerAccount(v string) *DeleteContainerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetResourceOwnerId(v int64) *DeleteContainerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteContainerGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContainerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContainerGroupResponseBody) SetRequestId(v string) *DeleteContainerGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteContainerGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteContainerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteContainerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteContainerGroupResponse) SetHeaders(v map[string]*string) *DeleteContainerGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteContainerGroupResponse) SetStatusCode(v int32) *DeleteContainerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContainerGroupResponse) SetBody(v *DeleteContainerGroupResponseBody) *DeleteContainerGroupResponse {
	s.Body = v
	return s
}

type DeleteImageCacheRequest struct {
	// The client token that is used to guarantee the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the image cache.
	ImageCacheId *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the image cache.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteImageCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageCacheRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageCacheRequest) SetClientToken(v string) *DeleteImageCacheRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteImageCacheRequest) SetImageCacheId(v string) *DeleteImageCacheRequest {
	s.ImageCacheId = &v
	return s
}

func (s *DeleteImageCacheRequest) SetOwnerAccount(v string) *DeleteImageCacheRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteImageCacheRequest) SetOwnerId(v int64) *DeleteImageCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteImageCacheRequest) SetRegionId(v string) *DeleteImageCacheRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteImageCacheRequest) SetResourceOwnerAccount(v string) *DeleteImageCacheRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteImageCacheRequest) SetResourceOwnerId(v int64) *DeleteImageCacheRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteImageCacheResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImageCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageCacheResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageCacheResponseBody) SetRequestId(v string) *DeleteImageCacheResponseBody {
	s.RequestId = &v
	return s
}

type DeleteImageCacheResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteImageCacheResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteImageCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageCacheResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageCacheResponse) SetHeaders(v map[string]*string) *DeleteImageCacheResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageCacheResponse) SetStatusCode(v int32) *DeleteImageCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImageCacheResponse) SetBody(v *DeleteImageCacheResponseBody) *DeleteImageCacheResponse {
	s.Body = v
	return s
}

type DeleteVirtualNodeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency of requests?](~~25693~~)
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the virtual node.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the virtual node.
	VirtualNodeId *string `json:"VirtualNodeId,omitempty" xml:"VirtualNodeId,omitempty"`
}

func (s DeleteVirtualNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualNodeRequest) SetClientToken(v string) *DeleteVirtualNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVirtualNodeRequest) SetOwnerAccount(v string) *DeleteVirtualNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVirtualNodeRequest) SetOwnerId(v int64) *DeleteVirtualNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVirtualNodeRequest) SetRegionId(v string) *DeleteVirtualNodeRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVirtualNodeRequest) SetResourceOwnerAccount(v string) *DeleteVirtualNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVirtualNodeRequest) SetResourceOwnerId(v int64) *DeleteVirtualNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVirtualNodeRequest) SetVirtualNodeId(v string) *DeleteVirtualNodeRequest {
	s.VirtualNodeId = &v
	return s
}

type DeleteVirtualNodeResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVirtualNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVirtualNodeResponseBody) SetRequestId(v string) *DeleteVirtualNodeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVirtualNodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVirtualNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVirtualNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualNodeResponse) SetHeaders(v map[string]*string) *DeleteVirtualNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirtualNodeResponse) SetStatusCode(v int32) *DeleteVirtualNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVirtualNodeResponse) SetBody(v *DeleteVirtualNodeResponseBody) *DeleteVirtualNodeResponse {
	s.Body = v
	return s
}

type DescribeAvailableResourceRequest struct {
	// The information about the resource that you want to query.
	DestinationResource *DescribeAvailableResourceRequestDestinationResource `json:"DestinationResource,omitempty" xml:"DestinationResource,omitempty" type:"Struct"`
	OwnerAccount        *string                                              `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64                                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// You can call the [DescribeRegions](~~146965~~) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The information about the preemptible instances that you want to query.
	SpotResource *DescribeAvailableResourceRequestSpotResource `json:"SpotResource,omitempty" xml:"SpotResource,omitempty" type:"Struct"`
	// The ID of the zone to which the instance belongs.
	//
	// This parameter is empty by default, which indicates that ECS instance families available in all zones in the specified region are queried.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequest) SetDestinationResource(v *DescribeAvailableResourceRequestDestinationResource) *DescribeAvailableResourceRequest {
	s.DestinationResource = v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetRegionId(v string) *DescribeAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetSpotResource(v *DescribeAvailableResourceRequestSpotResource) *DescribeAvailableResourceRequest {
	s.SpotResource = v
	return s
}

func (s *DescribeAvailableResourceRequest) SetZoneId(v string) *DescribeAvailableResourceRequest {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourceRequestDestinationResource struct {
	// The type of the resource. Valid values:
	//
	// *   InstanceTypeFamily: queries instance families. If you use this parameter value, you must also specify the Value parameter.
	// *   InstanceType: queries instance types. If you use this parameter value, you must also specify the Value, Cores, and Memory parameters.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The number of vCPUs. This parameter is available only when the Category parameter is set to InstanceType.
	Cores *float32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The size of the memory. Unit: GiB. This parameter is available only when the Category parameter is set to InstanceType.
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// Instance families or instance types.
	//
	// *   If you set Category to InstanceTypeFamily, you must set this parameter to instance families such as ecs.c5.
	// *   If you set Category to InstanceType, you must set this parameter to instance types such as ecs.c5.large.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAvailableResourceRequestDestinationResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceRequestDestinationResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequestDestinationResource) SetCategory(v string) *DescribeAvailableResourceRequestDestinationResource {
	s.Category = &v
	return s
}

func (s *DescribeAvailableResourceRequestDestinationResource) SetCores(v float32) *DescribeAvailableResourceRequestDestinationResource {
	s.Cores = &v
	return s
}

func (s *DescribeAvailableResourceRequestDestinationResource) SetMemory(v float32) *DescribeAvailableResourceRequestDestinationResource {
	s.Memory = &v
	return s
}

func (s *DescribeAvailableResourceRequestDestinationResource) SetValue(v string) *DescribeAvailableResourceRequestDestinationResource {
	s.Value = &v
	return s
}

type DescribeAvailableResourceRequestSpotResource struct {
	// The protection period of the preemptible instance. Unit: hours. Default value: 1. Valid values: 0 to 6.
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The maximum hourly price of the preemptible elastic container instance. The value can be accurate to three decimal places. If you set SpotStrategy to SpotWithPriceLimit, you must specify the SpotPriceLimit parameter.
	SpotPriceLimit *float64 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding policy for the elastic container instance. Valid values:
	//
	// *   NoSpot: The instance is a pay-as-you-go instance.
	// *   SpotWithPriceLimit: The instance is a preemptible instance with a user-defined maximum hourly price.
	// *   SpotAsPriceGo: The instance is a preemptible instance for which the market price at the time of purchase is automatically used as the bid price.
	//
	// Default value: NoSpot.
	//
	// >  If you set this parameter to SpotWithPriceLimit or SpotAsPriceGo to query preemptible instances, you must set Category to InstanceType. In addition, you must use the Value parameter to specify instance types, or use the Cores and Memory parameters to specify the number of vCPUs and memory size.
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s DescribeAvailableResourceRequestSpotResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceRequestSpotResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequestSpotResource) SetSpotDuration(v int32) *DescribeAvailableResourceRequestSpotResource {
	s.SpotDuration = &v
	return s
}

func (s *DescribeAvailableResourceRequestSpotResource) SetSpotPriceLimit(v float64) *DescribeAvailableResourceRequestSpotResource {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeAvailableResourceRequestSpotResource) SetSpotStrategy(v string) *DescribeAvailableResourceRequestSpotResource {
	s.SpotStrategy = &v
	return s
}

type DescribeAvailableResourceResponseBody struct {
	// An array of zones in which the specified resources are available.
	AvailableZones *DescribeAvailableResourceResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBody) SetAvailableZones(v *DescribeAvailableResourceResponseBodyAvailableZones) *DescribeAvailableResourceResponseBody {
	s.AvailableZones = v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetRequestId(v string) *DescribeAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZones struct {
	AvailableZone []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZones) SetAvailableZone(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) *DescribeAvailableResourceResponseBodyAvailableZones {
	s.AvailableZone = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone struct {
	// An array of resources available in the specified zone.
	AvailableResources *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Struct"`
	// The ID of the region where the instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the zone to which the instance belongs.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetAvailableResources(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.AvailableResources = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetRegionId(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetZoneId(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources struct {
	AvailableResource []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource `json:"AvailableResource,omitempty" xml:"AvailableResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) SetAvailableResource(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources {
	s.AvailableResource = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource struct {
	// An array of resources available in the zones.
	SupportedResources *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources `json:"SupportedResources,omitempty" xml:"SupportedResources,omitempty" type:"Struct"`
	// The resource type. Valid values:
	//
	// *   InstanceTypeFamily: instance families.
	// *   InstanceType: instance types.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) SetSupportedResources(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource {
	s.SupportedResources = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) SetType(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource {
	s.Type = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources struct {
	SupportedResource []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource `json:"SupportedResource,omitempty" xml:"SupportedResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) SetSupportedResource(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources {
	s.SupportedResource = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource struct {
	// The resource category based on the stock. Valid values:
	//
	// *   WithStock: Resources are in sufficient stock.
	// *   ClosedWithStock: Resources are insufficient. We recommend that you use other instance types that are in sufficient stock.
	// *   WithoutStock: Resources are sold out and will be replenished. We recommend that you use other instance types that are in sufficient stock.
	// *   ClosedWithoutStock: Resources are sold out and will not be replenished. We recommend that you use other instance types that are in sufficient stock.
	StatusCategory *string `json:"StatusCategory,omitempty" xml:"StatusCategory,omitempty"`
	// The ECS instance types or instance families that are supported in the zones.
	//
	// *   If the return value of the Type parameter is InstanceTypeFamily, this parameter indicates instance families that are returned.
	// *   If the return value of the Type parameter is InstanceType, this parameter indicates instance types that are returned.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetStatusCategory(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.StatusCategory = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetValue(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.Value = &v
	return s
}

type DescribeAvailableResourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponse) SetHeaders(v map[string]*string) *DescribeAvailableResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableResourceResponse) SetStatusCode(v int32) *DescribeAvailableResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableResourceResponse) SetBody(v *DescribeAvailableResourceResponseBody) *DescribeAvailableResourceResponse {
	s.Body = v
	return s
}

type DescribeCommitContainerTaskRequest struct {
	// The ID of the elastic container instance on which the CommitContainer task is executed.\
	// You must enter the instance ID, the task ID, or both for the request.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The number of entries to return on each page.\
	// Maximum value: 50.\
	// Default value: 10.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Set the value to the value of NextToken that is returned from the last call.
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the task.
	TaskId []*string `json:"TaskId,omitempty" xml:"TaskId,omitempty" type:"Repeated"`
	// The status of the task. Valid values:
	//
	// *   Running: The task is being executed.
	// *   Succeeded: The task is successfully executed.
	// *   Failed
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s DescribeCommitContainerTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCommitContainerTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeCommitContainerTaskRequest) SetContainerGroupId(v string) *DescribeCommitContainerTaskRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeCommitContainerTaskRequest) SetMaxResults(v int32) *DescribeCommitContainerTaskRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCommitContainerTaskRequest) SetNextToken(v string) *DescribeCommitContainerTaskRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCommitContainerTaskRequest) SetOwnerAccount(v string) *DescribeCommitContainerTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCommitContainerTaskRequest) SetOwnerId(v int64) *DescribeCommitContainerTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCommitContainerTaskRequest) SetRegionId(v string) *DescribeCommitContainerTaskRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCommitContainerTaskRequest) SetResourceOwnerAccount(v string) *DescribeCommitContainerTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCommitContainerTaskRequest) SetResourceOwnerId(v int64) *DescribeCommitContainerTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCommitContainerTaskRequest) SetTaskId(v []*string) *DescribeCommitContainerTaskRequest {
	s.TaskId = v
	return s
}

func (s *DescribeCommitContainerTaskRequest) SetTaskStatus(v string) *DescribeCommitContainerTaskRequest {
	s.TaskStatus = &v
	return s
}

type DescribeCommitContainerTaskResponseBody struct {
	// The details of the task.
	CommitTasks []*DescribeCommitContainerTaskResponseBodyCommitTasks `json:"CommitTasks,omitempty" xml:"CommitTasks,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token that is returned in this request.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCommitContainerTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCommitContainerTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCommitContainerTaskResponseBody) SetCommitTasks(v []*DescribeCommitContainerTaskResponseBodyCommitTasks) *DescribeCommitContainerTaskResponseBody {
	s.CommitTasks = v
	return s
}

func (s *DescribeCommitContainerTaskResponseBody) SetMaxResults(v string) *DescribeCommitContainerTaskResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeCommitContainerTaskResponseBody) SetNextToken(v string) *DescribeCommitContainerTaskResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeCommitContainerTaskResponseBody) SetRequestId(v string) *DescribeCommitContainerTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCommitContainerTaskResponseBody) SetTotalCount(v int32) *DescribeCommitContainerTaskResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCommitContainerTaskResponseBodyCommitTasks struct {
	// The stage of the CommitContainer task.
	CommitPhaseInfos []*DescribeCommitContainerTaskResponseBodyCommitTasksCommitPhaseInfos `json:"CommitPhaseInfos,omitempty" xml:"CommitPhaseInfos,omitempty" type:"Repeated"`
	// The name of the container.
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The time when the task was created.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The time when the task ended.
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The message about the status of the task.
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The ID of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The progress of the task in percentage.
	TaskProgress *string `json:"TaskProgress,omitempty" xml:"TaskProgress,omitempty"`
	// The status of the task. Valid values:
	//
	// *   Running: The task is being executed.
	// *   Succeeded: The task is successfully executed.
	// *   Failed
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s DescribeCommitContainerTaskResponseBodyCommitTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeCommitContainerTaskResponseBodyCommitTasks) GoString() string {
	return s.String()
}

func (s *DescribeCommitContainerTaskResponseBodyCommitTasks) SetCommitPhaseInfos(v []*DescribeCommitContainerTaskResponseBodyCommitTasksCommitPhaseInfos) *DescribeCommitContainerTaskResponseBodyCommitTasks {
	s.CommitPhaseInfos = v
	return s
}

func (s *DescribeCommitContainerTaskResponseBodyCommitTasks) SetContainerName(v string) *DescribeCommitContainerTaskResponseBodyCommitTasks {
	s.ContainerName = &v
	return s
}

func (s *DescribeCommitContainerTaskResponseBodyCommitTasks) SetCreationTime(v string) *DescribeCommitContainerTaskResponseBodyCommitTasks {
	s.CreationTime = &v
	return s
}

func (s *DescribeCommitContainerTaskResponseBodyCommitTasks) SetFinishedTime(v string) *DescribeCommitContainerTaskResponseBodyCommitTasks {
	s.FinishedTime = &v
	return s
}

func (s *DescribeCommitContainerTaskResponseBodyCommitTasks) SetStatusMessage(v string) *DescribeCommitContainerTaskResponseBodyCommitTasks {
	s.StatusMessage = &v
	return s
}

func (s *DescribeCommitContainerTaskResponseBodyCommitTasks) SetTaskId(v string) *DescribeCommitContainerTaskResponseBodyCommitTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeCommitContainerTaskResponseBodyCommitTasks) SetTaskProgress(v string) *DescribeCommitContainerTaskResponseBodyCommitTasks {
	s.TaskProgress = &v
	return s
}

func (s *DescribeCommitContainerTaskResponseBodyCommitTasks) SetTaskStatus(v string) *DescribeCommitContainerTaskResponseBodyCommitTasks {
	s.TaskStatus = &v
	return s
}

type DescribeCommitContainerTaskResponseBodyCommitTasksCommitPhaseInfos struct {
	// The information about the stage.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the stage.
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The time of the stage.
	RecordTime *string `json:"RecordTime,omitempty" xml:"RecordTime,omitempty"`
	// The status of the stage. Valid values:
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCommitContainerTaskResponseBodyCommitTasksCommitPhaseInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeCommitContainerTaskResponseBodyCommitTasksCommitPhaseInfos) GoString() string {
	return s.String()
}

func (s *DescribeCommitContainerTaskResponseBodyCommitTasksCommitPhaseInfos) SetMessage(v string) *DescribeCommitContainerTaskResponseBodyCommitTasksCommitPhaseInfos {
	s.Message = &v
	return s
}

func (s *DescribeCommitContainerTaskResponseBodyCommitTasksCommitPhaseInfos) SetPhase(v string) *DescribeCommitContainerTaskResponseBodyCommitTasksCommitPhaseInfos {
	s.Phase = &v
	return s
}

func (s *DescribeCommitContainerTaskResponseBodyCommitTasksCommitPhaseInfos) SetRecordTime(v string) *DescribeCommitContainerTaskResponseBodyCommitTasksCommitPhaseInfos {
	s.RecordTime = &v
	return s
}

func (s *DescribeCommitContainerTaskResponseBodyCommitTasksCommitPhaseInfos) SetStatus(v string) *DescribeCommitContainerTaskResponseBodyCommitTasksCommitPhaseInfos {
	s.Status = &v
	return s
}

type DescribeCommitContainerTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCommitContainerTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCommitContainerTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCommitContainerTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeCommitContainerTaskResponse) SetHeaders(v map[string]*string) *DescribeCommitContainerTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeCommitContainerTaskResponse) SetStatusCode(v int32) *DescribeCommitContainerTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCommitContainerTaskResponse) SetBody(v *DescribeCommitContainerTaskResponseBody) *DescribeCommitContainerTaskResponse {
	s.Body = v
	return s
}

type DescribeContainerGroupEventsRequest struct {
	// The IDs of the container groups. You can specify up to 20 IDs. Each ID must be a string in JSON format.
	ContainerGroupIds *string `json:"ContainerGroupIds,omitempty" xml:"ContainerGroupIds,omitempty"`
	// The component reporting this event. Valid values:
	//
	// *   EciService
	// *   K8sAgent
	//
	// This parameter is empty by default, which indicates that events reported by all types of components are queried.
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// Specifies the maximum number of container groups to be returned for this request. Default value: 200.
	//
	// >  The number of container groups to be returned is no greater than this parameter value.
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The token that is used for the next query. If this parameter is empty, all results have been returned.
	//
	// You do not need to specify this parameter in the first request. You can obtain the token from the result returned by the previous request.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the container groups.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// A relative time in seconds before the current time from which to show event information. This parameter is used to poll incremental events.
	SinceSecond *int32                                    `json:"SinceSecond,omitempty" xml:"SinceSecond,omitempty"`
	Tag         []*DescribeContainerGroupEventsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the container groups.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeContainerGroupEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsRequest) SetContainerGroupIds(v string) *DescribeContainerGroupEventsRequest {
	s.ContainerGroupIds = &v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetEventSource(v string) *DescribeContainerGroupEventsRequest {
	s.EventSource = &v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetLimit(v int32) *DescribeContainerGroupEventsRequest {
	s.Limit = &v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetNextToken(v string) *DescribeContainerGroupEventsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetRegionId(v string) *DescribeContainerGroupEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetResourceGroupId(v string) *DescribeContainerGroupEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetSinceSecond(v int32) *DescribeContainerGroupEventsRequest {
	s.SinceSecond = &v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetTag(v []*DescribeContainerGroupEventsRequestTag) *DescribeContainerGroupEventsRequest {
	s.Tag = v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetVSwitchId(v string) *DescribeContainerGroupEventsRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetZoneId(v string) *DescribeContainerGroupEventsRequest {
	s.ZoneId = &v
	return s
}

type DescribeContainerGroupEventsRequestTag struct {
	// The key of tag N.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeContainerGroupEventsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsRequestTag) SetKey(v string) *DescribeContainerGroupEventsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeContainerGroupEventsRequestTag) SetValue(v string) *DescribeContainerGroupEventsRequestTag {
	s.Value = &v
	return s
}

type DescribeContainerGroupEventsResponseBody struct {
	// An array of event details.
	Data []*DescribeContainerGroupEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries of returned events.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeContainerGroupEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsResponseBody) SetData(v []*DescribeContainerGroupEventsResponseBodyData) *DescribeContainerGroupEventsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeContainerGroupEventsResponseBody) SetRequestId(v string) *DescribeContainerGroupEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBody) SetTotalCount(v int32) *DescribeContainerGroupEventsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeContainerGroupEventsResponseBodyData struct {
	// The ID of the container group.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// An array of events.
	Events []*DescribeContainerGroupEventsResponseBodyDataEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupEventsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsResponseBodyData) SetContainerGroupId(v string) *DescribeContainerGroupEventsResponseBodyData {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyData) SetEvents(v []*DescribeContainerGroupEventsResponseBodyDataEvents) *DescribeContainerGroupEventsResponseBodyData {
	s.Events = v
	return s
}

type DescribeContainerGroupEventsResponseBodyDataEvents struct {
	// The number of events.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The time at which the event was first recorded.
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	// The time at which the most recent occurrence of this event was recorded.
	LastTimestamp *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// A description of the status of this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The metadata of the event.
	Metadata *DescribeContainerGroupEventsResponseBodyDataEventsMetadata `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Struct"`
	// The reason for the transition into the current status of the event.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// Name of the controller that emitted the event.
	ReportingComponent *string `json:"ReportingComponent,omitempty" xml:"ReportingComponent,omitempty"`
	// ID of the controller instance.
	ReportingInstance *string `json:"ReportingInstance,omitempty" xml:"ReportingInstance,omitempty"`
	// The component reporting this event.
	Source *DescribeContainerGroupEventsResponseBodyDataEventsSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	// The type of the event. Valid values:
	//
	// *   Normal
	// *   Warning
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The object that this event is about.
	InvolvedObject *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject `json:"involvedObject,omitempty" xml:"involvedObject,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupEventsResponseBodyDataEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsResponseBodyDataEvents) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetCount(v int32) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.Count = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetFirstTimestamp(v string) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.FirstTimestamp = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetLastTimestamp(v string) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.LastTimestamp = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetMessage(v string) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetMetadata(v *DescribeContainerGroupEventsResponseBodyDataEventsMetadata) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.Metadata = v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetReason(v string) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetReportingComponent(v string) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.ReportingComponent = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetReportingInstance(v string) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.ReportingInstance = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetSource(v *DescribeContainerGroupEventsResponseBodyDataEventsSource) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.Source = v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetType(v string) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.Type = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetInvolvedObject(v *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.InvolvedObject = v
	return s
}

type DescribeContainerGroupEventsResponseBodyDataEventsMetadata struct {
	// The name of the event.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DescribeContainerGroupEventsResponseBodyDataEventsMetadata) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsResponseBodyDataEventsMetadata) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsMetadata) SetName(v string) *DescribeContainerGroupEventsResponseBodyDataEventsMetadata {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsMetadata) SetNamespace(v string) *DescribeContainerGroupEventsResponseBodyDataEventsMetadata {
	s.Namespace = &v
	return s
}

type DescribeContainerGroupEventsResponseBodyDataEventsSource struct {
	// The component for the operation.
	Component *string `json:"Component,omitempty" xml:"Component,omitempty"`
	// The type of the host.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
}

func (s DescribeContainerGroupEventsResponseBodyDataEventsSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsResponseBodyDataEventsSource) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsSource) SetComponent(v string) *DescribeContainerGroupEventsResponseBodyDataEventsSource {
	s.Component = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsSource) SetHost(v string) *DescribeContainerGroupEventsResponseBodyDataEventsSource {
	s.Host = &v
	return s
}

type DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject struct {
	// The version of Kubernetes.
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	// The type of the resource.
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The name of the resource.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace where the resource resides.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the resource.
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject) SetApiVersion(v string) *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject {
	s.ApiVersion = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject) SetKind(v string) *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject {
	s.Kind = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject) SetName(v string) *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject) SetNamespace(v string) *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject {
	s.Namespace = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject) SetUid(v string) *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject {
	s.Uid = &v
	return s
}

type DescribeContainerGroupEventsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeContainerGroupEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerGroupEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsResponse) SetHeaders(v map[string]*string) *DescribeContainerGroupEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerGroupEventsResponse) SetStatusCode(v int32) *DescribeContainerGroupEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerGroupEventsResponse) SetBody(v *DescribeContainerGroupEventsResponseBody) *DescribeContainerGroupEventsResponse {
	s.Body = v
	return s
}

type DescribeContainerGroupMetricRequest struct {
	// The ID of the elastic container instance.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The end of the time range to query. The default value is the current time.
	//
	// Specify the time in RFC 3339 format.
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The data aggregation period. Unit: seconds. Valid values: 15, 30, 60, and 600. Default value: 60.
	//
	// >  If StartTime and EndTime are not specified, the system returns the monitoring data generated in the last 5 minutes with a data aggregation period of 15s. The Period parameter is ignored.
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the region.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. The beginning of the time range must be a time point in the past 30 days. The default value is 5 minutes before the value of EndTime.
	//
	// Specify the time in RFC 3339 format. To query the data starting from March 12, 2019 09:00 UTC+8, you can set the parameter to 2019-03-12T09:00:00.000+08:00 or 2019-03-12T01:00:00.000Z.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeContainerGroupMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricRequest) SetContainerGroupId(v string) *DescribeContainerGroupMetricRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetEndTime(v string) *DescribeContainerGroupMetricRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetOwnerAccount(v string) *DescribeContainerGroupMetricRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetOwnerId(v int64) *DescribeContainerGroupMetricRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetPeriod(v string) *DescribeContainerGroupMetricRequest {
	s.Period = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetRegionId(v string) *DescribeContainerGroupMetricRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetResourceOwnerAccount(v string) *DescribeContainerGroupMetricRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetResourceOwnerId(v int64) *DescribeContainerGroupMetricRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetStartTime(v string) *DescribeContainerGroupMetricRequest {
	s.StartTime = &v
	return s
}

type DescribeContainerGroupMetricResponseBody struct {
	// The ID of the elastic container instance.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The monitoring data of the elastic container instance.
	Records []*DescribeContainerGroupMetricResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBody) SetContainerGroupId(v string) *DescribeContainerGroupMetricResponseBody {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBody) SetRecords(v []*DescribeContainerGroupMetricResponseBodyRecords) *DescribeContainerGroupMetricResponseBody {
	s.Records = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBody) SetRequestId(v string) *DescribeContainerGroupMetricResponseBody {
	s.RequestId = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecords struct {
	// The monitoring data of vCPUs.
	CPU *DescribeContainerGroupMetricResponseBodyRecordsCPU `json:"CPU,omitempty" xml:"CPU,omitempty" type:"Struct"`
	// The monitoring data of containers.
	Containers []*DescribeContainerGroupMetricResponseBodyRecordsContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// The monitoring data of disks.
	Disk []*DescribeContainerGroupMetricResponseBodyRecordsDisk `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
	// The monitoring data of file system partitions.
	Filesystem []*DescribeContainerGroupMetricResponseBodyRecordsFilesystem `json:"Filesystem,omitempty" xml:"Filesystem,omitempty" type:"Repeated"`
	// The monitoring data of the memory.
	Memory *DescribeContainerGroupMetricResponseBodyRecordsMemory `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	// The monitoring data of the network.
	Network *DescribeContainerGroupMetricResponseBodyRecordsNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The time when the monitoring data entry was collected. The time is indicated in RFC 3339 format.
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetCPU(v *DescribeContainerGroupMetricResponseBodyRecordsCPU) *DescribeContainerGroupMetricResponseBodyRecords {
	s.CPU = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetContainers(v []*DescribeContainerGroupMetricResponseBodyRecordsContainers) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Containers = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetDisk(v []*DescribeContainerGroupMetricResponseBodyRecordsDisk) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Disk = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetFilesystem(v []*DescribeContainerGroupMetricResponseBodyRecordsFilesystem) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Filesystem = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetMemory(v *DescribeContainerGroupMetricResponseBodyRecordsMemory) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Memory = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetNetwork(v *DescribeContainerGroupMetricResponseBodyRecordsNetwork) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Network = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetTimestamp(v string) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Timestamp = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsCPU struct {
	// The upper limit of vCPU usage (the maximum number of vCPUs × 1000).
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The average load in the last 10 seconds.
	Load *int64 `json:"Load,omitempty" xml:"Load,omitempty"`
	// The cumulative usage of vCPUs.
	UsageCoreNanoSeconds *int64 `json:"UsageCoreNanoSeconds,omitempty" xml:"UsageCoreNanoSeconds,omitempty"`
	// The vCPU usage in the sampling window. Unit for the sampling window: nanoseconds.
	UsageNanoCores *int64 `json:"UsageNanoCores,omitempty" xml:"UsageNanoCores,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsCPU) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsCPU) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsCPU) SetLimit(v int64) *DescribeContainerGroupMetricResponseBodyRecordsCPU {
	s.Limit = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsCPU) SetLoad(v int64) *DescribeContainerGroupMetricResponseBodyRecordsCPU {
	s.Load = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsCPU) SetUsageCoreNanoSeconds(v int64) *DescribeContainerGroupMetricResponseBodyRecordsCPU {
	s.UsageCoreNanoSeconds = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsCPU) SetUsageNanoCores(v int64) *DescribeContainerGroupMetricResponseBodyRecordsCPU {
	s.UsageNanoCores = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsContainers struct {
	// The vCPU monitoring data of the container.
	CPU *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU `json:"CPU,omitempty" xml:"CPU,omitempty" type:"Struct"`
	// The memory monitoring data of the container.
	Memory *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	// The name of the container.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainers) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainers) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainers) SetCPU(v *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) *DescribeContainerGroupMetricResponseBodyRecordsContainers {
	s.CPU = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainers) SetMemory(v *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) *DescribeContainerGroupMetricResponseBodyRecordsContainers {
	s.Memory = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainers) SetName(v string) *DescribeContainerGroupMetricResponseBodyRecordsContainers {
	s.Name = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsContainersCPU struct {
	// The upper limit of vCPU usage (the maximum number of vCPUs × 1000).
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The average load in the last 10 seconds.
	Load *int64 `json:"Load,omitempty" xml:"Load,omitempty"`
	// The cumulative usage of vCPUs.
	UsageCoreNanoSeconds *int64 `json:"UsageCoreNanoSeconds,omitempty" xml:"UsageCoreNanoSeconds,omitempty"`
	// The vCPU usage in the sampling window. Unit for the sampling window: nanoseconds.
	UsageNanoCores *int64 `json:"UsageNanoCores,omitempty" xml:"UsageNanoCores,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) SetLimit(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU {
	s.Limit = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) SetLoad(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU {
	s.Load = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) SetUsageCoreNanoSeconds(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU {
	s.UsageCoreNanoSeconds = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) SetUsageNanoCores(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU {
	s.UsageNanoCores = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsContainersMemory struct {
	// The amount of available memory.
	AvailableBytes *int64 `json:"AvailableBytes,omitempty" xml:"AvailableBytes,omitempty"`
	// The size of the cache.
	Cache *int64 `json:"Cache,omitempty" xml:"Cache,omitempty"`
	// The resident set size, which indicates the amount of physical memory that is actually used.
	Rss *int64 `json:"Rss,omitempty" xml:"Rss,omitempty"`
	// The amount of used memory.
	UsageBytes *int64 `json:"UsageBytes,omitempty" xml:"UsageBytes,omitempty"`
	// The usage of the working set.
	WorkingSet *int64 `json:"WorkingSet,omitempty" xml:"WorkingSet,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) SetAvailableBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory {
	s.AvailableBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) SetCache(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory {
	s.Cache = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) SetRss(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory {
	s.Rss = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) SetUsageBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory {
	s.UsageBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) SetWorkingSet(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory {
	s.WorkingSet = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsDisk struct {
	// The name of the disk.
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The amount of data that was read from the disk. Unit: bytes.
	ReadBytes *int64 `json:"ReadBytes,omitempty" xml:"ReadBytes,omitempty"`
	// This parameter is unavailable.
	ReadIO *int64 `json:"ReadIO,omitempty" xml:"ReadIO,omitempty"`
	// The amount of data that was written to the disk. Unit: bytes.
	WriteBytes *int64 `json:"WriteBytes,omitempty" xml:"WriteBytes,omitempty"`
	// This parameter is unavailable.
	WriteIO *int64 `json:"WriteIO,omitempty" xml:"WriteIO,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsDisk) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsDisk) SetDevice(v string) *DescribeContainerGroupMetricResponseBodyRecordsDisk {
	s.Device = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsDisk) SetReadBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsDisk {
	s.ReadBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsDisk) SetReadIO(v int64) *DescribeContainerGroupMetricResponseBodyRecordsDisk {
	s.ReadIO = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsDisk) SetWriteBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsDisk {
	s.WriteBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsDisk) SetWriteIO(v int64) *DescribeContainerGroupMetricResponseBodyRecordsDisk {
	s.WriteIO = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsFilesystem struct {
	// The size of the available space.
	Available *int64 `json:"Available,omitempty" xml:"Available,omitempty"`
	// The total size of space.
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The type of the partition. Valid values:
	//
	// *   System
	// *   Volume
	// *   Other
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The name of the partition.
	FsName *string `json:"FsName,omitempty" xml:"FsName,omitempty"`
	// The size of used space.
	Usage *int64 `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsFilesystem) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsFilesystem) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsFilesystem) SetAvailable(v int64) *DescribeContainerGroupMetricResponseBodyRecordsFilesystem {
	s.Available = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsFilesystem) SetCapacity(v int64) *DescribeContainerGroupMetricResponseBodyRecordsFilesystem {
	s.Capacity = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsFilesystem) SetCategory(v string) *DescribeContainerGroupMetricResponseBodyRecordsFilesystem {
	s.Category = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsFilesystem) SetFsName(v string) *DescribeContainerGroupMetricResponseBodyRecordsFilesystem {
	s.FsName = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsFilesystem) SetUsage(v int64) *DescribeContainerGroupMetricResponseBodyRecordsFilesystem {
	s.Usage = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsMemory struct {
	// The amount of available memory.
	AvailableBytes *int64 `json:"AvailableBytes,omitempty" xml:"AvailableBytes,omitempty"`
	// The size of the cache.
	Cache *int64 `json:"Cache,omitempty" xml:"Cache,omitempty"`
	// The resident set size, which indicates the amount of physical memory that is actually used.
	Rss *int64 `json:"Rss,omitempty" xml:"Rss,omitempty"`
	// The amount of used memory.
	UsageBytes *int64 `json:"UsageBytes,omitempty" xml:"UsageBytes,omitempty"`
	// The usage of the working set.
	WorkingSet *int64 `json:"WorkingSet,omitempty" xml:"WorkingSet,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsMemory) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsMemory) SetAvailableBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsMemory {
	s.AvailableBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsMemory) SetCache(v int64) *DescribeContainerGroupMetricResponseBodyRecordsMemory {
	s.Cache = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsMemory) SetRss(v int64) *DescribeContainerGroupMetricResponseBodyRecordsMemory {
	s.Rss = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsMemory) SetUsageBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsMemory {
	s.UsageBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsMemory) SetWorkingSet(v int64) *DescribeContainerGroupMetricResponseBodyRecordsMemory {
	s.WorkingSet = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsNetwork struct {
	// The monitoring data of network interface controllers (NICs).
	Interfaces []*DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces `json:"Interfaces,omitempty" xml:"Interfaces,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsNetwork) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsNetwork) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetwork) SetInterfaces(v []*DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) *DescribeContainerGroupMetricResponseBodyRecordsNetwork {
	s.Interfaces = v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces struct {
	// The name of the NIC.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of bytes received by the NIC.
	RxBytes *int64 `json:"RxBytes,omitempty" xml:"RxBytes,omitempty"`
	// The number of packets that the NIC failed to receive.
	RxDrops *int64 `json:"RxDrops,omitempty" xml:"RxDrops,omitempty"`
	// The number of bytes that the NIC failed to receive.
	RxErrors *int64 `json:"RxErrors,omitempty" xml:"RxErrors,omitempty"`
	// The number of packets received by the NIC.
	RxPackets *int64 `json:"RxPackets,omitempty" xml:"RxPackets,omitempty"`
	// The number of bytes sent by the NIC.
	TxBytes *int64 `json:"TxBytes,omitempty" xml:"TxBytes,omitempty"`
	// The number of packets that the NIC failed to send.
	TxDrops *int64 `json:"TxDrops,omitempty" xml:"TxDrops,omitempty"`
	// The number of bytes that the NIC failed to send.
	TxErrors *int64 `json:"TxErrors,omitempty" xml:"TxErrors,omitempty"`
	// The number of packets sent by the NIC.
	TxPackets *int64 `json:"TxPackets,omitempty" xml:"TxPackets,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetName(v string) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetRxBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.RxBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetRxDrops(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.RxDrops = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetRxErrors(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.RxErrors = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetRxPackets(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.RxPackets = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetTxBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.TxBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetTxDrops(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.TxDrops = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetTxErrors(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.TxErrors = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetTxPackets(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.TxPackets = &v
	return s
}

type DescribeContainerGroupMetricResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeContainerGroupMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerGroupMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponse) SetHeaders(v map[string]*string) *DescribeContainerGroupMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerGroupMetricResponse) SetStatusCode(v int32) *DescribeContainerGroupMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerGroupMetricResponse) SetBody(v *DescribeContainerGroupMetricResponseBody) *DescribeContainerGroupMetricResponse {
	s.Body = v
	return s
}

type DescribeContainerGroupPriceRequest struct {
	// The number of vCPUs.
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The storage size of the temporary storage space. Unit: GiB.
	EphemeralStorage *int32 `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	// The instance type of the Elastic Compute Service (ECS) instance used to create the elastic container instance. For more information about ECS instance types that can be used to create elastic container instances, see [Instance types](~~114664~~).
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The size of the memory. Unit: GiB.
	Memory       *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	OwnerAccount *string  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the instance resides. You can call the [DescribeRegions](~~146965~~) operation to query the most recent region and zone list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SpotDuration         *int32  `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The maximum hourly price of the preemptible elastic container instance. The value can contain up to three decimal places. If you set the SpotStrategy parameter to SpotWithPriceLimit, you must specify the SpotPriceLimit parameter.
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding policy for the elastic container instance. Valid values:
	//
	// *   NoSpot: The instance is a regular pay-as-you-go instance.
	// *   SpotWithPriceLimit: The instance is a preemptible instance with a user-defined maximum hourly price.
	// *   SpotAsPriceGo: The instance is a preemptible instance for which the market price at the time of purchase is automatically used as the bid price.
	//
	// Default value: NoSpot.
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The zone ID of the instance. You can call the [DescribeRegions](~~146965~~) operation to query the most recent region and zone list.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeContainerGroupPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceRequest) SetCpu(v float32) *DescribeContainerGroupPriceRequest {
	s.Cpu = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetEphemeralStorage(v int32) *DescribeContainerGroupPriceRequest {
	s.EphemeralStorage = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetInstanceType(v string) *DescribeContainerGroupPriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetMemory(v float32) *DescribeContainerGroupPriceRequest {
	s.Memory = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetOwnerAccount(v string) *DescribeContainerGroupPriceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetOwnerId(v int64) *DescribeContainerGroupPriceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetRegionId(v string) *DescribeContainerGroupPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetResourceOwnerAccount(v string) *DescribeContainerGroupPriceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetResourceOwnerId(v int64) *DescribeContainerGroupPriceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetSpotDuration(v int32) *DescribeContainerGroupPriceRequest {
	s.SpotDuration = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetSpotPriceLimit(v float32) *DescribeContainerGroupPriceRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetSpotStrategy(v string) *DescribeContainerGroupPriceRequest {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetZoneId(v string) *DescribeContainerGroupPriceRequest {
	s.ZoneId = &v
	return s
}

type DescribeContainerGroupPriceResponseBody struct {
	// The information about the price and discount rule.
	PriceInfo *DescribeContainerGroupPriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerGroupPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBody) SetPriceInfo(v *DescribeContainerGroupPriceResponseBodyPriceInfo) *DescribeContainerGroupPriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribeContainerGroupPriceResponseBody) SetRequestId(v string) *DescribeContainerGroupPriceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfo struct {
	// The price.
	Price *DescribeContainerGroupPriceResponseBodyPriceInfoPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// Details about the promotion rules.
	Rules *DescribeContainerGroupPriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	// The details about the price of the preemptible elastic container instance.
	SpotPrices *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices `json:"SpotPrices,omitempty" xml:"SpotPrices,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfo) SetPrice(v *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) *DescribeContainerGroupPriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfo) SetRules(v *DescribeContainerGroupPriceResponseBodyPriceInfoRules) *DescribeContainerGroupPriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfo) SetSpotPrices(v *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices) *DescribeContainerGroupPriceResponseBodyPriceInfo {
	s.SpotPrices = v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoPrice struct {
	// The currency unit. Valid values:
	//
	// - CNY: The value is available only on the China site (aliyun.com).
	// - USD: The value is available only on the international site (alibabacloud.com).
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The details about the price.
	DetailInfos *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos `json:"DetailInfos,omitempty" xml:"DetailInfos,omitempty" type:"Struct"`
	// The discount.
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The transaction price, which is equal to the original price minus the discount.
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPrice) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) SetDetailInfos(v *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos) *DescribeContainerGroupPriceResponseBodyPriceInfoPrice {
	s.DetailInfos = v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos struct {
	DetailInfo []*DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos) SetDetailInfo(v []*DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos {
	s.DetailInfo = v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo struct {
	// The discount.
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The name of the resource.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// Details about the pricing rules.
	Rules *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	// The transaction price.
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetDiscountPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetOriginalPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetResource(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.Resource = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetRules(v *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.Rules = v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetTradePrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.TradePrice = &v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules struct {
	Rule []*DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules) SetRule(v []*DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules {
	s.Rule = v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule struct {
	// The description of the promotion rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the promotion rule.
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule) SetDescription(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule {
	s.Description = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule) SetRuleId(v int64) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule {
	s.RuleId = &v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoRules struct {
	Rule []*DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoRules) SetRule(v []*DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule) *DescribeContainerGroupPriceResponseBodyPriceInfoRules {
	s.Rule = v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule struct {
	// The description of the promotion rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the promotion rule.
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule) SetDescription(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule {
	s.Description = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule) SetRuleId(v int64) *DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule {
	s.RuleId = &v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices struct {
	SpotPrice []*DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice `json:"SpotPrice,omitempty" xml:"SpotPrice,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices) SetSpotPrice(v []*DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices {
	s.SpotPrice = v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice struct {
	// The instance type of the ECS instance.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The original price.
	OriginPrice *float32 `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	// The price of the preemptible elastic container instance.
	SpotPrice *float32 `json:"SpotPrice,omitempty" xml:"SpotPrice,omitempty"`
	// The zone ID of the instance.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) SetInstanceType(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice {
	s.InstanceType = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) SetOriginPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice {
	s.OriginPrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) SetSpotPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice {
	s.SpotPrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) SetZoneId(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice {
	s.ZoneId = &v
	return s
}

type DescribeContainerGroupPriceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeContainerGroupPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerGroupPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponse) SetHeaders(v map[string]*string) *DescribeContainerGroupPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerGroupPriceResponse) SetStatusCode(v int32) *DescribeContainerGroupPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerGroupPriceResponse) SetBody(v *DescribeContainerGroupPriceResponseBody) *DescribeContainerGroupPriceResponse {
	s.Body = v
	return s
}

type DescribeContainerGroupStatusRequest struct {
	// The IDs of the container groups. You can specify up to 20 IDs. Each ID must be a string in the JSON format.
	ContainerGroupIds *string `json:"ContainerGroupIds,omitempty" xml:"ContainerGroupIds,omitempty"`
	// Specifies the maximum number of container groups to be returned for this request. Default value: 200.
	//
	// >  The number of container groups to be returned is no greater than this parameter value.
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The token that is used for the next query. If this parameter is empty, all results have been returned.
	//
	// You do not need to specify this parameter in the first request. You can obtain the token from the result returned by the previous request.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the container groups.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the container groups belong.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// A relative time in seconds before the current time from which to show container groups whose status changes. This parameter is used to poll status of container groups.
	SinceSecond *int32                                    `json:"SinceSecond,omitempty" xml:"SinceSecond,omitempty"`
	Tag         []*DescribeContainerGroupStatusRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone to which the container groups belong.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeContainerGroupStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusRequest) SetContainerGroupIds(v string) *DescribeContainerGroupStatusRequest {
	s.ContainerGroupIds = &v
	return s
}

func (s *DescribeContainerGroupStatusRequest) SetLimit(v int32) *DescribeContainerGroupStatusRequest {
	s.Limit = &v
	return s
}

func (s *DescribeContainerGroupStatusRequest) SetNextToken(v string) *DescribeContainerGroupStatusRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeContainerGroupStatusRequest) SetRegionId(v string) *DescribeContainerGroupStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerGroupStatusRequest) SetResourceGroupId(v string) *DescribeContainerGroupStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeContainerGroupStatusRequest) SetSinceSecond(v int32) *DescribeContainerGroupStatusRequest {
	s.SinceSecond = &v
	return s
}

func (s *DescribeContainerGroupStatusRequest) SetTag(v []*DescribeContainerGroupStatusRequestTag) *DescribeContainerGroupStatusRequest {
	s.Tag = v
	return s
}

func (s *DescribeContainerGroupStatusRequest) SetVSwitchId(v string) *DescribeContainerGroupStatusRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeContainerGroupStatusRequest) SetZoneId(v string) *DescribeContainerGroupStatusRequest {
	s.ZoneId = &v
	return s
}

type DescribeContainerGroupStatusRequestTag struct {
	// The key of tag N.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeContainerGroupStatusRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusRequestTag) SetKey(v string) *DescribeContainerGroupStatusRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeContainerGroupStatusRequestTag) SetValue(v string) *DescribeContainerGroupStatusRequestTag {
	s.Value = &v
	return s
}

type DescribeContainerGroupStatusResponseBody struct {
	// The collection of the status of the container groups.
	Data []*DescribeContainerGroupStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The token that was returned for the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBody) SetData(v []*DescribeContainerGroupStatusResponseBodyData) *DescribeContainerGroupStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBody) SetNextToken(v string) *DescribeContainerGroupStatusResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBody) SetRequestId(v string) *DescribeContainerGroupStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBody) SetTotalCount(v int32) *DescribeContainerGroupStatusResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyData struct {
	// The ID of the container group.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The name of the container group.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace in which the container group resides.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The status of the pod.
	PodStatus *DescribeContainerGroupStatusResponseBodyDataPodStatus `json:"PodStatus,omitempty" xml:"PodStatus,omitempty" type:"Struct"`
	// The status of the container group.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The universally unique identifier (UUID) of the container group, which is similar to the unique identifier (UID) of the Kubernetes pod in terms of the concept and usage.
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyData) SetContainerGroupId(v string) *DescribeContainerGroupStatusResponseBodyData {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyData) SetName(v string) *DescribeContainerGroupStatusResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyData) SetNamespace(v string) *DescribeContainerGroupStatusResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyData) SetPodStatus(v *DescribeContainerGroupStatusResponseBodyDataPodStatus) *DescribeContainerGroupStatusResponseBodyData {
	s.PodStatus = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyData) SetStatus(v string) *DescribeContainerGroupStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyData) SetUuid(v string) *DescribeContainerGroupStatusResponseBodyData {
	s.Uuid = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatus struct {
	// The conditions of the pod.
	Conditions []*DescribeContainerGroupStatusResponseBodyDataPodStatusConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The collection of the status about the container groups.
	ContainerStatuses []*DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses `json:"ContainerStatuses,omitempty" xml:"ContainerStatuses,omitempty" type:"Repeated"`
	// The IP address of the server.
	HostIp *string `json:"HostIp,omitempty" xml:"HostIp,omitempty"`
	// The lifecycle phase of the pod.
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The IP address of the pod.
	PodIp *string `json:"PodIp,omitempty" xml:"PodIp,omitempty"`
	// The collection of pod IP addresses.
	PodIps []*DescribeContainerGroupStatusResponseBodyDataPodStatusPodIps `json:"PodIps,omitempty" xml:"PodIps,omitempty" type:"Repeated"`
	// The quality of service (QoS) of the pod.
	QosClass *string `json:"QosClass,omitempty" xml:"QosClass,omitempty"`
	// The time when the container began to run.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatus) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatus) SetConditions(v []*DescribeContainerGroupStatusResponseBodyDataPodStatusConditions) *DescribeContainerGroupStatusResponseBodyDataPodStatus {
	s.Conditions = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatus) SetContainerStatuses(v []*DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) *DescribeContainerGroupStatusResponseBodyDataPodStatus {
	s.ContainerStatuses = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatus) SetHostIp(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatus {
	s.HostIp = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatus) SetPhase(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatus {
	s.Phase = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatus) SetPodIp(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatus {
	s.PodIp = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatus) SetPodIps(v []*DescribeContainerGroupStatusResponseBodyDataPodStatusPodIps) *DescribeContainerGroupStatusResponseBodyDataPodStatus {
	s.PodIps = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatus) SetQosClass(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatus {
	s.QosClass = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatus) SetStartTime(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatus {
	s.StartTime = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusConditions struct {
	// The event message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason for the transition into the current status of the event.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The time when the status last changed.
	LastTransitionTime *string `json:"lastTransitionTime,omitempty" xml:"lastTransitionTime,omitempty"`
	// The status of the pod condition.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the pod condition. Valid values:
	//
	// *   PodScheduled
	// *   Ready
	// *   Initialized
	// *   Unschedulable
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusConditions) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusConditions) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions) SetMessage(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions) SetReason(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions) SetLastTransitionTime(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions {
	s.LastTransitionTime = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions) SetStatus(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions {
	s.Status = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions) SetType(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions {
	s.Type = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses struct {
	// The image of the container.
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The ID of the image.
	ImageID *string `json:"ImageID,omitempty" xml:"ImageID,omitempty"`
	// The latest status of the container.
	LastState *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState `json:"LastState,omitempty" xml:"LastState,omitempty" type:"Struct"`
	// The name of the container.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the container is ready for use.
	Ready *bool `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// The number of times that the container restarted.
	RestartCount *int32 `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	// Indicates whether the container is started.
	Started *bool `json:"Started,omitempty" xml:"Started,omitempty"`
	// The state of the container. Valid values:
	//
	// *   Waiting
	// *   Running
	// *   Terminated
	State *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState `json:"State,omitempty" xml:"State,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) SetImage(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses {
	s.Image = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) SetImageID(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses {
	s.ImageID = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) SetLastState(v *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses {
	s.LastState = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) SetName(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) SetReady(v bool) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses {
	s.Ready = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) SetRestartCount(v int32) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses {
	s.RestartCount = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) SetStarted(v bool) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses {
	s.Started = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) SetState(v *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses {
	s.State = v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState struct {
	// The container is created and running.
	Running *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateRunning `json:"Running,omitempty" xml:"Running,omitempty" type:"Struct"`
	// The container is terminated and exits after a successful or failed run.
	Terminated *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated `json:"Terminated,omitempty" xml:"Terminated,omitempty" type:"Struct"`
	// The container is waiting for being created.
	Waiting *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting `json:"Waiting,omitempty" xml:"Waiting,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState) SetRunning(v *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateRunning) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState {
	s.Running = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState) SetTerminated(v *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState {
	s.Terminated = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState) SetWaiting(v *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState {
	s.Waiting = v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateRunning struct {
	// The beginning of the time range that was queried.
	StartedAtstartedAt *string `json:"StartedAtstartedAt,omitempty" xml:"StartedAtstartedAt,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateRunning) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateRunning) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateRunning) SetStartedAtstartedAt(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateRunning {
	s.StartedAtstartedAt = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated struct {
	// The ID of the container.
	ContainerID *string `json:"ContainerID,omitempty" xml:"ContainerID,omitempty"`
	// The exit code.
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The end of the time range that was queried.
	FinishedAt *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	// The event message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason for the transition into the current status of the event.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The signal code.
	Signal *int32 `json:"Signal,omitempty" xml:"Signal,omitempty"`
	// The beginning of the time range that was queried.
	StartedAt *string `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) SetContainerID(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated {
	s.ContainerID = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) SetExitCode(v int32) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated {
	s.ExitCode = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) SetFinishedAt(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated {
	s.FinishedAt = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) SetMessage(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) SetReason(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) SetSignal(v int32) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated {
	s.Signal = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) SetStartedAt(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated {
	s.StartedAt = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting struct {
	// The event message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason for the transition into the current status of the event.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting) SetMessage(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting) SetReason(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting {
	s.Reason = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState struct {
	// The container is created and running.
	Running *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateRunning `json:"Running,omitempty" xml:"Running,omitempty" type:"Struct"`
	// The container is terminated and exits after a successful or failed run.
	Terminated *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated `json:"Terminated,omitempty" xml:"Terminated,omitempty" type:"Struct"`
	// The container is waiting for being created.
	Waiting *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting `json:"Waiting,omitempty" xml:"Waiting,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState) SetRunning(v *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateRunning) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState {
	s.Running = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState) SetTerminated(v *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState {
	s.Terminated = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState) SetWaiting(v *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState {
	s.Waiting = v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateRunning struct {
	// The beginning of the time range that was queried.
	StartedAtstartedAt *string `json:"StartedAtstartedAt,omitempty" xml:"StartedAtstartedAt,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateRunning) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateRunning) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateRunning) SetStartedAtstartedAt(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateRunning {
	s.StartedAtstartedAt = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated struct {
	// The ID of the container.
	ContainerID *string `json:"ContainerID,omitempty" xml:"ContainerID,omitempty"`
	// The exit code.
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The end of the time range that was queried.
	FinishedAt *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	// The event message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason for the transition into the current status of the event.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The signal code.
	Signal *int32 `json:"Signal,omitempty" xml:"Signal,omitempty"`
	// The beginning of the time range that was queried.
	StartedAt *string `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) SetContainerID(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated {
	s.ContainerID = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) SetExitCode(v int32) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated {
	s.ExitCode = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) SetFinishedAt(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated {
	s.FinishedAt = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) SetMessage(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) SetReason(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) SetSignal(v int32) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated {
	s.Signal = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) SetStartedAt(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated {
	s.StartedAt = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting struct {
	// The event message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason for the transition into the current status of the event.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting) SetMessage(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting) SetReason(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting {
	s.Reason = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusPodIps struct {
	// The IP address of the pod.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusPodIps) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusPodIps) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusPodIps) SetIp(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusPodIps {
	s.Ip = &v
	return s
}

type DescribeContainerGroupStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeContainerGroupStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerGroupStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponse) SetHeaders(v map[string]*string) *DescribeContainerGroupStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerGroupStatusResponse) SetStatusCode(v int32) *DescribeContainerGroupStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerGroupStatusResponse) SetBody(v *DescribeContainerGroupStatusResponseBody) *DescribeContainerGroupStatusResponse {
	s.Body = v
	return s
}

type DescribeContainerGroupsRequest struct {
	// The IDs of the elastic container instances in JSON format. You can specify up to 20 IDs.
	ContainerGroupIds *string `json:"ContainerGroupIds,omitempty" xml:"ContainerGroupIds,omitempty"`
	// The name of the elastic container instance.
	ContainerGroupName *string `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	// The maximum number of resources to return. Default value: 20. Maximum value: 20.
	//
	// >  The number of returned resources is less than or equal to the specified number.
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The token that determines the start point of the query. If this parameter is left empty, all results have been returned.
	//
	// > You do not need to specify this parameter in the first request. Starting from the second request, you can obtain the token from the result returned by the previous request.
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the elastic container instances belong. If you do not specify a resource group when you create an elastic container instance, the system automatically adds the instance to the default resource group in your account.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the elastic container instance. Valid values:
	//
	// *   Pending: The instance is being started.
	// *   Running: The instance is running.
	// *   Succeeded: The instance runs successfully.
	// *   Failed: The instance fails to run.
	// *   Scheduling: The instance is being created.
	// *   ScheduleFailed: The instance fails to be created.
	// *   Restarting: The instance is being restarted.
	// *   Updating: The instance is being updated.
	// *   Terminating: The instance is being terminated.
	// *   Expired: The instance expires.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag of the instances.
	Tag []*DescribeContainerGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which the elastic container instances are connected.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Specifies whether to return event information.
	WithEvent *bool `json:"WithEvent,omitempty" xml:"WithEvent,omitempty"`
	// The ID of the zone in which the elastic container instances are deployed. If you do not specify this parameter, the system selects a zone.
	//
	// This parameter is empty by default.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeContainerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsRequest) SetContainerGroupIds(v string) *DescribeContainerGroupsRequest {
	s.ContainerGroupIds = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetContainerGroupName(v string) *DescribeContainerGroupsRequest {
	s.ContainerGroupName = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetLimit(v int32) *DescribeContainerGroupsRequest {
	s.Limit = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetNextToken(v string) *DescribeContainerGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetOwnerAccount(v string) *DescribeContainerGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetOwnerId(v int64) *DescribeContainerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetRegionId(v string) *DescribeContainerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetResourceGroupId(v string) *DescribeContainerGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetResourceOwnerAccount(v string) *DescribeContainerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetResourceOwnerId(v int64) *DescribeContainerGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetStatus(v string) *DescribeContainerGroupsRequest {
	s.Status = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetTag(v []*DescribeContainerGroupsRequestTag) *DescribeContainerGroupsRequest {
	s.Tag = v
	return s
}

func (s *DescribeContainerGroupsRequest) SetVSwitchId(v string) *DescribeContainerGroupsRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetWithEvent(v bool) *DescribeContainerGroupsRequest {
	s.WithEvent = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetZoneId(v string) *DescribeContainerGroupsRequest {
	s.ZoneId = &v
	return s
}

type DescribeContainerGroupsRequestTag struct {
	// The tag key of the instances.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instances.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeContainerGroupsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsRequestTag) SetKey(v string) *DescribeContainerGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeContainerGroupsRequestTag) SetValue(v string) *DescribeContainerGroupsRequestTag {
	s.Value = &v
	return s
}

type DescribeContainerGroupsResponseBody struct {
	// Details of the instances.
	ContainerGroups []*DescribeContainerGroupsResponseBodyContainerGroups `json:"ContainerGroups,omitempty" xml:"ContainerGroups,omitempty" type:"Repeated"`
	// The token that determines the start point of the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request. The value is unique.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of queried instances.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeContainerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBody) SetContainerGroups(v []*DescribeContainerGroupsResponseBodyContainerGroups) *DescribeContainerGroupsResponseBody {
	s.ContainerGroups = v
	return s
}

func (s *DescribeContainerGroupsResponseBody) SetNextToken(v string) *DescribeContainerGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeContainerGroupsResponseBody) SetRequestId(v string) *DescribeContainerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBody) SetTotalCount(v int32) *DescribeContainerGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroups struct {
	// The ID of the instance.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The name of the instance.
	ContainerGroupName *string `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	// The containers in the elastic container instance.
	Containers []*DescribeContainerGroupsResponseBodyContainerGroupsContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// The number of vCPUs that are allocated to the elastic container instance.
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The time when the instance was created. The time follows the RFC 3339 standard and must be in UTC.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The discount.
	Discount *int32 `json:"Discount,omitempty" xml:"Discount,omitempty"`
	// The Domain Name System (DNS) settings.
	DnsConfig *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig `json:"DnsConfig,omitempty" xml:"DnsConfig,omitempty" type:"Struct"`
	// The security context of the elastic container instance.
	EciSecurityContext *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext `json:"EciSecurityContext,omitempty" xml:"EciSecurityContext,omitempty" type:"Struct"`
	// The ID of the elastic network interface (ENI).
	EniInstanceId *string `json:"EniInstanceId,omitempty" xml:"EniInstanceId,omitempty"`
	// The size of the temporary storage space. Unit: GiB.
	EphemeralStorage *int32 `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	// The events of the elastic container instance. A maximum of 50 events can be returned.
	Events []*DescribeContainerGroupsResponseBodyContainerGroupsEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The time when the elastic container instance failed to run due to overdue payments. The time follows the RFC 3339 standard and must be in UTC.
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The time when the instance failed to run. The time follows the RFC 3339 standard and must be in UTC.
	FailedTime *string `json:"FailedTime,omitempty" xml:"FailedTime,omitempty"`
	// The hostnames and IP addresses of a container that are added to the hosts file of the elastic container instance.
	HostAliases []*DescribeContainerGroupsResponseBodyContainerGroupsHostAliases `json:"HostAliases,omitempty" xml:"HostAliases,omitempty" type:"Repeated"`
	// The init containers.
	InitContainers []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainers `json:"InitContainers,omitempty" xml:"InitContainers,omitempty" type:"Repeated"`
	// The instance type of the Elastic Compute Service (ECS) instance used to create the elastic container instance.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The public IP address.
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address.
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The IPv6 address of the instance.
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	// The memory size of the elastic container instance. Unit: GiB.
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the instance RAM role. You can use an instance RAM role to access both elastic container instances and ECS instances. For more information, see [Use the instance RAM role by calling APIs](~~61178~~).
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The restart policy of the elastic container instance. Valid values:
	//
	// *   Never: The instance never restarts.
	// *   Always: The instance always restarts.
	// *   OnFailure: The instance restarts if it fails to run.
	RestartPolicy *string `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	// The ID of the security group to which the instances belong.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The maximum hourly price for the preemptible instance.
	//
	// This parameter is returned only if you set the SpotStrategy parameter to SpotWithPriceLimit.
	SpotPriceLimit *float64 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding policy of the elastic container instance. Default value: NoSpot. Valid values:
	//
	// *   NoSpot: The instance is a regular pay-as-you-go instance.
	// *   SpotWithPriceLimit: The instance is a preemptible instance with a user-defined maximum hourly price.
	// *   SpotAsPriceGo: The instance is a preemptible instance for which the market price at the time of purchase is automatically used as the bid price.
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The status of the instance. Valid values:
	//
	// *   Pending: The instance is being started.
	// *   Running: The instance is running.
	// *   Succeeded: The instance runs successfully.
	// *   Failed: The instance fails to run.
	// *   Scheduling: The instance is being created.
	// *   ScheduleFailed: The instance fails to be created.
	// *   Restarting: The instance is being restarted.
	// *   Updating: The instance is being updated.
	// *   Terminating: The instance is being terminated.
	// *   Expired: The instance expires.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when all containers in the elastic container instance exited on success. The time follows the RFC 3339 standard and must be in UTC.
	SucceededTime *string `json:"SucceededTime,omitempty" xml:"SucceededTime,omitempty"`
	// The tags of the instances.
	Tags []*DescribeContainerGroupsResponseBodyContainerGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// This parameter is unavailable.
	TenantEniInstanceId *string `json:"TenantEniInstanceId,omitempty" xml:"TenantEniInstanceId,omitempty"`
	// This parameter is unavailable.
	TenantEniIp *string `json:"TenantEniIp,omitempty" xml:"TenantEniIp,omitempty"`
	// This parameter is unavailable.
	TenantSecurityGroupId *string `json:"TenantSecurityGroupId,omitempty" xml:"TenantSecurityGroupId,omitempty"`
	// This parameter is unavailable.
	TenantVSwitchId *string `json:"TenantVSwitchId,omitempty" xml:"TenantVSwitchId,omitempty"`
	// The ID of the vSwitch to which the instance is connected.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Information about the volumes.
	Volumes []*DescribeContainerGroupsResponseBodyContainerGroupsVolumes `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the elastic container instances belong.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone in which the elastic container instances are deployed.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroups) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetContainerGroupId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetContainerGroupName(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.ContainerGroupName = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetContainers(v []*DescribeContainerGroupsResponseBodyContainerGroupsContainers) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Containers = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetCpu(v float32) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Cpu = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetCreationTime(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.CreationTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetDiscount(v int32) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Discount = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetDnsConfig(v *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.DnsConfig = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetEciSecurityContext(v *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.EciSecurityContext = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetEniInstanceId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.EniInstanceId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetEphemeralStorage(v int32) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.EphemeralStorage = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetEvents(v []*DescribeContainerGroupsResponseBodyContainerGroupsEvents) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Events = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetExpiredTime(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetFailedTime(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.FailedTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetHostAliases(v []*DescribeContainerGroupsResponseBodyContainerGroupsHostAliases) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.HostAliases = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetInitContainers(v []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.InitContainers = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetInstanceType(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.InstanceType = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetInternetIp(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.InternetIp = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetIntranetIp(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.IntranetIp = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetIpv6Address(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Ipv6Address = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetMemory(v float32) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Memory = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetRamRoleName(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.RamRoleName = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetRegionId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetResourceGroupId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetRestartPolicy(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.RestartPolicy = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetSecurityGroupId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetSpotPriceLimit(v float64) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetSpotStrategy(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetStatus(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Status = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetSucceededTime(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.SucceededTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetTags(v []*DescribeContainerGroupsResponseBodyContainerGroupsTags) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Tags = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetTenantEniInstanceId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.TenantEniInstanceId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetTenantEniIp(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.TenantEniIp = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetTenantSecurityGroupId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.TenantSecurityGroupId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetTenantVSwitchId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.TenantVSwitchId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetVSwitchId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.VSwitchId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetVolumes(v []*DescribeContainerGroupsResponseBodyContainerGroupsVolumes) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Volumes = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetVpcId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.VpcId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetZoneId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.ZoneId = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainers struct {
	// Arguments that are used to start the container.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The command that is used to start the container.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The number of vCPUs.
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The current status of the container.
	CurrentState *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState `json:"CurrentState,omitempty" xml:"CurrentState,omitempty" type:"Struct"`
	// The environment variables of the container.
	EnvironmentVars []*DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The number of GPUs.
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The image of the container.
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The policy to pull images. Valid values:
	//
	// *   Always: Each time instances are created, image pulling is performed.
	// *   IfNotPresent: Image pulling is performed as needed. On-premises images are preferentially used. If no on-premises images are available, image pulling is performed.
	// *   Never: On-premises images are always used. Image pulling is not performed.
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The liveness probe of the container.
	LivenessProbe *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe `json:"LivenessProbe,omitempty" xml:"LivenessProbe,omitempty" type:"Struct"`
	// The memory size of the container. Unit: GiB.
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the container.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The exposed ports and protocols of the container.
	Ports []*DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The previous status of the container.
	PreviousState *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState `json:"PreviousState,omitempty" xml:"PreviousState,omitempty" type:"Struct"`
	// The readiness probe that is used to check whether the container is ready to serve a request.
	ReadinessProbe *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe `json:"ReadinessProbe,omitempty" xml:"ReadinessProbe,omitempty" type:"Struct"`
	// Indicates whether the container passed the readiness probe.
	Ready *bool `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// The number of times that the container restarted.
	RestartCount *int32 `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	// The security context of the elastic container instance.
	SecurityContext *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	// Indicates whether the container allocates buffer resources to stdin when the container runs. If this parameter is not configured, an EOF error may occur. Default value: false.
	Stdin *bool `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	// Indicates whether the container runtime closes the stdin channel after the stdin channel has been opened by a single attach session. If stdin is true, the stdin stream remains open across multiple attach sessions. If StdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and receive data until the client disconnects. When the client disconnects, stdin is closed and remains closed until the container is restarted.
	StdinOnce *bool `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	// Indicates whether interaction was enabled. Default value: false. If the Command parameter is a `/bin/bash` command, set this parameter to true.
	Tty *bool `json:"Tty,omitempty" xml:"Tty,omitempty"`
	// Information about the mounted volumes.
	VolumeMounts []*DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts `json:"VolumeMounts,omitempty" xml:"VolumeMounts,omitempty" type:"Repeated"`
	// The working directory of the container.
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainers) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainers) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetArgs(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Args = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetCommands(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Commands = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetCpu(v float32) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Cpu = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetCurrentState(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.CurrentState = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetEnvironmentVars(v []*DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.EnvironmentVars = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetGpu(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Gpu = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetImage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Image = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetImagePullPolicy(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetLivenessProbe(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.LivenessProbe = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetMemory(v float32) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Memory = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetPorts(v []*DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Ports = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetPreviousState(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.PreviousState = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetReadinessProbe(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.ReadinessProbe = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetReady(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Ready = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetRestartCount(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.RestartCount = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetSecurityContext(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.SecurityContext = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetStdin(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Stdin = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetStdinOnce(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.StdinOnce = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetTty(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Tty = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetVolumeMounts(v []*DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.VolumeMounts = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetWorkingDir(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.WorkingDir = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState struct {
	// The details of the container status.
	DetailStatus *string `json:"DetailStatus,omitempty" xml:"DetailStatus,omitempty"`
	// The exit code of the container.
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time when the container stopped running.
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The message about the container status.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason why the container is in this state.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The code of the container status.
	Signal *int32 `json:"Signal,omitempty" xml:"Signal,omitempty"`
	// The time when the container started to run.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the container. Valid values:
	//
	// *   Waiting
	// *   Running
	// *   Terminated
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetDetailStatus(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.DetailStatus = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetExitCode(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.ExitCode = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetFinishTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.FinishTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetMessage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetReason(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetSignal(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.Signal = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetStartTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.StartTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetState(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.State = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars struct {
	// The name of the environment variable.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The source of the environment variable value. This parameter has a value only when the Value parameter is not empty.
	ValueFrom *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom `json:"ValueFrom,omitempty" xml:"ValueFrom,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) SetKey(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars {
	s.Key = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) SetValue(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars {
	s.Value = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) SetValueFrom(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom) *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars {
	s.ValueFrom = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom struct {
	// The specified field.
	FieldRef *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom) SetFieldRef(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef) *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom {
	s.FieldRef = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef struct {
	// The path of the field.
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef) SetFieldPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef {
	s.FieldPath = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe struct {
	// The commands that were run.
	Execs []*string `json:"Execs,omitempty" xml:"Execs,omitempty" type:"Repeated"`
	// The minimum number of consecutive failures that must occur for the check to be considered failure. Default value: 3.
	FailureThreshold *int32 `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	// The HTTP GET method that is used to check the container.
	HttpGet *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	// The number of seconds between the time when the startup of the container ends and the time when the check starts.
	InitialDelaySeconds *int32 `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	// The interval at which the check is performed. Unit: seconds. Default value: 10. Minimum value: 1.
	PeriodSeconds *int32 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	// The minimum number of consecutive successes that must occur for the check to be considered successful. Default value: 1. Set the value to 1.
	SuccessThreshold *int32 `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	// The TCP socket method that is used to check the container.
	TcpSocket *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
	// The timeout period of the health check. Unit: seconds. Default value: 1. Minimum value: 1.
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetExecs(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.Execs = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetFailureThreshold(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetHttpGet(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.HttpGet = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetInitialDelaySeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetPeriodSeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetSuccessThreshold(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetTcpSocket(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.TcpSocket = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetTimeoutSeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.TimeoutSeconds = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet struct {
	// The path to which HTTP GET requests were sent.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The port to which HTTP GET requests were sent.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type of the HTTP GET requests.
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) SetPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) SetScheme(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet {
	s.Scheme = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket struct {
	// The hostname.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port number.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket) SetHost(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket {
	s.Host = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket {
	s.Port = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts struct {
	// The port number. Valid values: 1 to 65535.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the protocol.
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts {
	s.Port = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts) SetProtocol(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts {
	s.Protocol = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState struct {
	// The details of the container status.
	DetailStatus *string `json:"DetailStatus,omitempty" xml:"DetailStatus,omitempty"`
	// The exit code of the container.
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time when the container stopped running.
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The message about the container status.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason why the container is in this state.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The code of the container status.
	Signal *int32 `json:"Signal,omitempty" xml:"Signal,omitempty"`
	// The time when the container started to run.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the container. Valid values:
	//
	// *   Waiting: The container is being started.
	// *   Running: The container is running.
	// *   Terminated: The container terminates running.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetDetailStatus(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.DetailStatus = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetExitCode(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.ExitCode = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetFinishTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.FinishTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetMessage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetReason(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetSignal(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.Signal = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetStartTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.StartTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetState(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.State = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe struct {
	// The commands that are run in the container when you use the command line interface (CLI) to perform health checks.
	Execs []*string `json:"Execs,omitempty" xml:"Execs,omitempty" type:"Repeated"`
	// The minimum number of consecutive failures that must occur for the check to be considered failure. Default value: 3.
	FailureThreshold *int32 `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	// The HTTP GET method that is used to check the container.
	HttpGet *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	// The number of seconds between the time when the startup of the container ends and the time when the check starts.
	InitialDelaySeconds *int32 `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	// The interval at which the check is performed. Unit: seconds. Default value: 10. Minimum value: 1.
	PeriodSeconds *int32 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	// The minimum number of consecutive successes that must occur for the check to be considered successful. Default value: 1. Set the value to 1.
	SuccessThreshold *int32 `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	// The TCP socket method that is used to check the container.
	TcpSocket *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
	// The timeout period of the health check. Unit: seconds. Default value: 1. Minimum value: 1.
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetExecs(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.Execs = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetFailureThreshold(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetHttpGet(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.HttpGet = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetInitialDelaySeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetPeriodSeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetSuccessThreshold(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetTcpSocket(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.TcpSocket = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetTimeoutSeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.TimeoutSeconds = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet struct {
	// The path to which HTTP GET requests were sent.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The port to which the system sends an HTTP GET request for a health check.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type of the HTTP GET requests.
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) SetPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) SetScheme(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet {
	s.Scheme = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket struct {
	// The IP address of the host.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port number.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket) SetHost(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket {
	s.Host = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket {
	s.Port = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext struct {
	// The permissions specific to the processes in the container.
	Capability *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" type:"Struct"`
	// Indicates whether the root file system is set to the read-only mode. The only valid value is true.
	ReadOnlyRootFilesystem *bool `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	// The ID of the user that runs the container.
	RunAsUser *int64 `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) SetCapability(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability) *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext {
	s.Capability = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) SetReadOnlyRootFilesystem(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) SetRunAsUser(v int64) *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext {
	s.RunAsUser = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability struct {
	// The permissions specific to the processes in the container.
	Adds []*string `json:"Adds,omitempty" xml:"Adds,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability) SetAdds(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability {
	s.Adds = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts struct {
	// The directory to which the volume is mounted. Data under this directory is overwritten by the data on the volume.
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation setting of the volume. Mount propagation allows you to share volumes that are mounted on a container to other containers in the same pod, or even to other pods on the same node. Valid values:
	//
	// *   None: The volume mount does not receive subsequent mounts that are mounted to this volume or its subdirectories.
	// *   HostToCotainer: The volume mount receives all subsequent mounts that are mounted to this volume or its subdirectories.
	// *   Bidirectional: This value has a similar effect as HostToContainer. The volume mount receives all subsequent mounts that are mounted to this volume or its subdirectories. In addition, all volume mounts created by the container are propagated back to the host and to all containers of all pods that use the same volume.
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The name of the volume.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the volumes are read-only.
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) SetMountPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) SetMountPropagation(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) SetReadOnly(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts {
	s.ReadOnly = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig struct {
	// The IP addresses of DNS servers.
	NameServers []*string `json:"NameServers,omitempty" xml:"NameServers,omitempty" type:"Repeated"`
	// The options. Each option is a name-value pair. The value in the name-value pair is optional.
	Options []*DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
	// The search domains of the DNS server.
	Searches []*string `json:"Searches,omitempty" xml:"Searches,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) SetNameServers(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig {
	s.NameServers = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) SetOptions(v []*DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions) *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig {
	s.Options = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) SetSearches(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig {
	s.Searches = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions struct {
	// The variable name of the option.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The variable value of the option.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions) SetValue(v string) *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions {
	s.Value = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext struct {
	// Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch.
	Sysctls []*DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls `json:"Sysctls,omitempty" xml:"Sysctls,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext) SetSysctls(v []*DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls) *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext {
	s.Sysctls = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls struct {
	// The name of the Sysctl parameter.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the Sysctl parameter.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls) SetValue(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls {
	s.Value = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsEvents struct {
	// The number of the events.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The time when the event started.
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	// The time when the event ended.
	LastTimestamp *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The message about the event.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the category to which the event belongs.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the event.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The type of the event. Valid values:
	//
	// *   Normal
	// *   Warning
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEvents) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetCount(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.Count = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetFirstTimestamp(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.FirstTimestamp = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetLastTimestamp(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.LastTimestamp = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetMessage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetReason(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetType(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.Type = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsHostAliases struct {
	// The information about the host.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// The IP address of the host.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsHostAliases) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsHostAliases) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsHostAliases) SetHostnames(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsHostAliases {
	s.Hostnames = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsHostAliases) SetIp(v string) *DescribeContainerGroupsResponseBodyContainerGroupsHostAliases {
	s.Ip = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainers struct {
	// The arguments that are used to start the container.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The command that is used to start the container.
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// The number of vCPUs.
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The current status of the container.
	CurrentState *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState `json:"CurrentState,omitempty" xml:"CurrentState,omitempty" type:"Struct"`
	// The environment variables of the container.
	EnvironmentVars []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The number of GPUs.
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The image of the container.
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The policy to pull an image.
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The size of memory that is allocated to the init container. Unit: GiB.
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the init container.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The exposed ports and protocols of the container.
	Ports []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The previous status of the container.
	PreviousState *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState `json:"PreviousState,omitempty" xml:"PreviousState,omitempty" type:"Struct"`
	// Indicates whether the container passed the readiness probe.
	Ready *bool `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// The number of times that the container restarted.
	RestartCount *int32 `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	// The security context of the container.
	SecurityContext *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	// The information about the volumes that are mounted to the init container.
	VolumeMounts []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts `json:"VolumeMounts,omitempty" xml:"VolumeMounts,omitempty" type:"Repeated"`
	// The working directory of the container.
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetArgs(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Args = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetCommand(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Command = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetCpu(v float32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Cpu = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetCurrentState(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.CurrentState = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetEnvironmentVars(v []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.EnvironmentVars = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetGpu(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Gpu = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetImage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Image = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetImagePullPolicy(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetMemory(v float32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Memory = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetPorts(v []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Ports = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetPreviousState(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.PreviousState = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetReady(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Ready = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetRestartCount(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.RestartCount = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetSecurityContext(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.SecurityContext = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetVolumeMounts(v []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.VolumeMounts = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetWorkingDir(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.WorkingDir = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState struct {
	// The details of the container status.
	DetailStatus *string `json:"DetailStatus,omitempty" xml:"DetailStatus,omitempty"`
	// The exit code of the container.
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time when the container stopped running.
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The message of the event.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason why the container is in this state.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The code of the container status.
	Signal *int32 `json:"Signal,omitempty" xml:"Signal,omitempty"`
	// The time when the container started to run.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the container. Valid values:
	//
	// *   Waiting
	// *   Running
	// *   Terminated
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetDetailStatus(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.DetailStatus = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetExitCode(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.ExitCode = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetFinishTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.FinishTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetMessage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetReason(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetSignal(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.Signal = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetStartTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.StartTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetState(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.State = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars struct {
	// The name of the environment variable.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The source of the environment variable value. This parameter has a value only when the Value parameter is not empty.
	ValueFrom *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom `json:"ValueFrom,omitempty" xml:"ValueFrom,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) SetKey(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars {
	s.Key = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) SetValue(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars {
	s.Value = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) SetValueFrom(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars {
	s.ValueFrom = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom struct {
	// The specified field.
	FieldRef *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom) SetFieldRef(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom {
	s.FieldRef = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef struct {
	// The path from which the field is selected in the specified version. Only `status.podIP` is supported.
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef) SetFieldPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef {
	s.FieldPath = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts struct {
	// The port number. Valid values: 1 to 65535.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the protocol.
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts {
	s.Port = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts) SetProtocol(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts {
	s.Protocol = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState struct {
	// The details of the container status.
	DetailStatus *string `json:"DetailStatus,omitempty" xml:"DetailStatus,omitempty"`
	// The exit code of the container.
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time when the container stopped running.
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The message about the container status.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason why the container is in this state.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The code of the container status.
	Signal *int32 `json:"Signal,omitempty" xml:"Signal,omitempty"`
	// The time when the container started to run.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the container. Valid values: Waiting, Running, and Terminated.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetDetailStatus(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.DetailStatus = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetExitCode(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.ExitCode = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetFinishTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.FinishTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetMessage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetReason(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetSignal(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.Signal = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetStartTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.StartTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetState(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.State = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext struct {
	// The permissions specific to the processes in the container.
	Capability *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" type:"Struct"`
	// Indicates whether the root file system is set to the read-only mode. The only valid value is true.
	ReadOnlyRootFilesystem *bool `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	// The UID that is used to run the entry point of the container process.
	RunAsUser *int64 `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) SetCapability(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext {
	s.Capability = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) SetReadOnlyRootFilesystem(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) SetRunAsUser(v int64) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext {
	s.RunAsUser = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability struct {
	// The permissions specific to the processes in the container.
	Adds []*string `json:"Adds,omitempty" xml:"Adds,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability) SetAdds(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability {
	s.Adds = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts struct {
	// The directory to which the volume is mounted. Data under this directory is overwritten by the data on the volume.
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation setting of the volume. Mount propagation allows you to share volumes that are mounted on a container to other containers in the same pod, or even to other pods on the same node. Valid values:
	//
	// *   None: The volume mount does not receive subsequent mounts that are mounted to this volume or its subdirectories.
	// *   HostToCotainer: The volume mount receives all subsequent mounts that are mounted to this volume or its subdirectories.
	// *   Bidirectional: This value has a similar effect as HostToContainer. The volume mount receives all subsequent mounts that are mounted to this volume or its subdirectories. In addition, all volume mounts created by the container are propagated back to the host and to all containers of all pods that use the same volume.
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The name of the volume. The name is the same as the volume you selected when you purchased the container.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Default value: False.
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) SetMountPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) SetMountPropagation(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) SetReadOnly(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts {
	s.ReadOnly = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsTags struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsTags) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsTags) SetKey(v string) *DescribeContainerGroupsResponseBodyContainerGroupsTags {
	s.Key = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsTags) SetValue(v string) *DescribeContainerGroupsResponseBodyContainerGroupsTags {
	s.Value = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsVolumes struct {
	// The paths of the ConfigFile volume.
	ConfigFileVolumeConfigFileToPaths []*DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths `json:"ConfigFileVolumeConfigFileToPaths,omitempty" xml:"ConfigFileVolumeConfigFileToPaths,omitempty" type:"Repeated"`
	// The ID of the disk volume.
	DiskVolumeDiskId *string `json:"DiskVolumeDiskId,omitempty" xml:"DiskVolumeDiskId,omitempty"`
	// The file system type of the disk volume.
	DiskVolumeFsType *string `json:"DiskVolumeFsType,omitempty" xml:"DiskVolumeFsType,omitempty"`
	// The name of the driver when you set the Type parameter to FlexVolume.
	FlexVolumeDriver *string `json:"FlexVolumeDriver,omitempty" xml:"FlexVolumeDriver,omitempty"`
	// The file system type when you set the Type parameter to FlexVolume. The default value is determined by the script of the FlexVolume plug-in.
	FlexVolumeFsType *string `json:"FlexVolumeFsType,omitempty" xml:"FlexVolumeFsType,omitempty"`
	// The options when you set the Type parameter to FlexVolume.
	FlexVolumeOptions *string `json:"FlexVolumeOptions,omitempty" xml:"FlexVolumeOptions,omitempty"`
	// The path of the NFS volume.
	NFSVolumePath *string `json:"NFSVolumePath,omitempty" xml:"NFSVolumePath,omitempty"`
	// Indicates whether the NFS volume is read-only.
	NFSVolumeReadOnly *bool `json:"NFSVolumeReadOnly,omitempty" xml:"NFSVolumeReadOnly,omitempty"`
	// The endpoint of the server when you set the Type parameter to NFSVolume.
	NFSVolumeServer *string `json:"NFSVolumeServer,omitempty" xml:"NFSVolumeServer,omitempty"`
	// The name of the volume.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the volume. Valid values:
	//
	// *   EmptyDirVolume
	// *   NFSVolume
	// *   ConfigFileVolume
	// *   FlexVolume
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsVolumes) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsVolumes) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetConfigFileVolumeConfigFileToPaths(v []*DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.ConfigFileVolumeConfigFileToPaths = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetDiskVolumeDiskId(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.DiskVolumeDiskId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetDiskVolumeFsType(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.DiskVolumeFsType = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetFlexVolumeDriver(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.FlexVolumeDriver = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetFlexVolumeFsType(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.FlexVolumeFsType = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetFlexVolumeOptions(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.FlexVolumeOptions = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetNFSVolumePath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.NFSVolumePath = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetNFSVolumeReadOnly(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.NFSVolumeReadOnly = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetNFSVolumeServer(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.NFSVolumeServer = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetType(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.Type = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths struct {
	// The content of the ConfigFile volume. Maximum size: 32 KB.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The relative path of the ConfigFile volume.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths) SetContent(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths {
	s.Content = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths) SetPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths {
	s.Path = &v
	return s
}

type DescribeContainerGroupsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeContainerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponse) SetHeaders(v map[string]*string) *DescribeContainerGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerGroupsResponse) SetStatusCode(v int32) *DescribeContainerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerGroupsResponse) SetBody(v *DescribeContainerGroupsResponseBody) *DescribeContainerGroupsResponse {
	s.Body = v
	return s
}

type DescribeContainerLogRequest struct {
	// The ID of the elastic container instance.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The name of the container.
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// Specifies whether to query the previous container if the container exits and restarts. Valid values:
	//
	// *   true: queries the previous container.
	// *   false: does not query the previous container.
	//
	// Default value: false.
	LastTime *bool `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The limit on the total size of logs.
	LimitBytes   *int64  `json:"LimitBytes,omitempty" xml:"LimitBytes,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the elastic container instance is deployed.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// A relative time in seconds before the current time from which to show logs. Example: 10, 20, and 30.
	SinceSeconds *int32 `json:"SinceSeconds,omitempty" xml:"SinceSeconds,omitempty"`
	// The beginning of the time range to query. Specify the time in the RFC 3339 standard. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The number of log entries that you want to query. Default value: 500. Maximum value: 2000. A maximum of 1 MB of logs can be returned.
	Tail *int32 `json:"Tail,omitempty" xml:"Tail,omitempty"`
	// Specifies whether to return the timestamps of logs. Valid values:
	//
	// *   true: returns the timestamps of logs.
	// *   false: does not return the timestamps of logs.
	//
	// Default value: false.
	Timestamps *bool `json:"Timestamps,omitempty" xml:"Timestamps,omitempty"`
}

func (s DescribeContainerLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerLogRequest) SetContainerGroupId(v string) *DescribeContainerLogRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeContainerLogRequest) SetContainerName(v string) *DescribeContainerLogRequest {
	s.ContainerName = &v
	return s
}

func (s *DescribeContainerLogRequest) SetLastTime(v bool) *DescribeContainerLogRequest {
	s.LastTime = &v
	return s
}

func (s *DescribeContainerLogRequest) SetLimitBytes(v int64) *DescribeContainerLogRequest {
	s.LimitBytes = &v
	return s
}

func (s *DescribeContainerLogRequest) SetOwnerAccount(v string) *DescribeContainerLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeContainerLogRequest) SetOwnerId(v int64) *DescribeContainerLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeContainerLogRequest) SetRegionId(v string) *DescribeContainerLogRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerLogRequest) SetResourceOwnerAccount(v string) *DescribeContainerLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeContainerLogRequest) SetResourceOwnerId(v int64) *DescribeContainerLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeContainerLogRequest) SetSinceSeconds(v int32) *DescribeContainerLogRequest {
	s.SinceSeconds = &v
	return s
}

func (s *DescribeContainerLogRequest) SetStartTime(v string) *DescribeContainerLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeContainerLogRequest) SetTail(v int32) *DescribeContainerLogRequest {
	s.Tail = &v
	return s
}

func (s *DescribeContainerLogRequest) SetTimestamps(v bool) *DescribeContainerLogRequest {
	s.Timestamps = &v
	return s
}

type DescribeContainerLogResponseBody struct {
	// The name of the container.
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The content of the log.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerLogResponseBody) SetContainerName(v string) *DescribeContainerLogResponseBody {
	s.ContainerName = &v
	return s
}

func (s *DescribeContainerLogResponseBody) SetContent(v string) *DescribeContainerLogResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeContainerLogResponseBody) SetRequestId(v string) *DescribeContainerLogResponseBody {
	s.RequestId = &v
	return s
}

type DescribeContainerLogResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeContainerLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerLogResponse) SetHeaders(v map[string]*string) *DescribeContainerLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerLogResponse) SetStatusCode(v int32) *DescribeContainerLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerLogResponse) SetBody(v *DescribeContainerLogResponseBody) *DescribeContainerLogResponse {
	s.Body = v
	return s
}

type DescribeImageCachesRequest struct {
	// The image of the container.
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The ID of the image cache.
	ImageCacheId *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	// The name of the image cache.
	ImageCacheName *string `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	// The maximum number of query results that you want to be displayed.
	Limit      *int32    `json:"Limit,omitempty" xml:"Limit,omitempty"`
	MatchImage []*string `json:"MatchImage,omitempty" xml:"MatchImage,omitempty" type:"Repeated"`
	// The token that determines the start point of the query. Set the value to the NextToken value that is returned from the last query.
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the image cache.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the snapshot that is used to create the image cache.
	SnapshotId *string                          `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	Tag        []*DescribeImageCachesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeImageCachesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesRequest) SetImage(v string) *DescribeImageCachesRequest {
	s.Image = &v
	return s
}

func (s *DescribeImageCachesRequest) SetImageCacheId(v string) *DescribeImageCachesRequest {
	s.ImageCacheId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetImageCacheName(v string) *DescribeImageCachesRequest {
	s.ImageCacheName = &v
	return s
}

func (s *DescribeImageCachesRequest) SetLimit(v int32) *DescribeImageCachesRequest {
	s.Limit = &v
	return s
}

func (s *DescribeImageCachesRequest) SetMatchImage(v []*string) *DescribeImageCachesRequest {
	s.MatchImage = v
	return s
}

func (s *DescribeImageCachesRequest) SetNextToken(v string) *DescribeImageCachesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeImageCachesRequest) SetOwnerAccount(v string) *DescribeImageCachesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeImageCachesRequest) SetOwnerId(v int64) *DescribeImageCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetRegionId(v string) *DescribeImageCachesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetResourceGroupId(v string) *DescribeImageCachesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetResourceOwnerAccount(v string) *DescribeImageCachesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeImageCachesRequest) SetResourceOwnerId(v int64) *DescribeImageCachesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetSnapshotId(v string) *DescribeImageCachesRequest {
	s.SnapshotId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetTag(v []*DescribeImageCachesRequestTag) *DescribeImageCachesRequest {
	s.Tag = v
	return s
}

type DescribeImageCachesRequestTag struct {
	// The key of tag N of the image cache. Valid values of N: 1 to 20.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the image cache. Valid values of N: 1 to 20.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImageCachesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesRequestTag) SetKey(v string) *DescribeImageCachesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeImageCachesRequestTag) SetValue(v string) *DescribeImageCachesRequestTag {
	s.Value = &v
	return s
}

type DescribeImageCachesResponseBody struct {
	// The image caches.
	ImageCaches []*DescribeImageCachesResponseBodyImageCaches `json:"ImageCaches,omitempty" xml:"ImageCaches,omitempty" type:"Repeated"`
	// The query token that is returned in this query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageCachesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesResponseBody) SetImageCaches(v []*DescribeImageCachesResponseBodyImageCaches) *DescribeImageCachesResponseBody {
	s.ImageCaches = v
	return s
}

func (s *DescribeImageCachesResponseBody) SetNextToken(v string) *DescribeImageCachesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeImageCachesResponseBody) SetRequestId(v string) *DescribeImageCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageCachesResponseBody) SetTotalCount(v int32) *DescribeImageCachesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeImageCachesResponseBodyImageCaches struct {
	// The ID of the container group.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The time when the image cache was created.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The elimination policy of the image cache. This parameter is empty by default, which indicates that the image cache is always retained.
	//
	// You can set this parameter to LRU, which indicates that the image cache can be automatically deleted. When the number of image caches reaches the quota, the system automatically deletes the image caches whose EliminationStrategy parameter is set to LRU and that are least commonly used.
	EliminationStrategy *string `json:"EliminationStrategy,omitempty" xml:"EliminationStrategy,omitempty"`
	// Details about the events of pulling images to create the image cache.
	Events []*DescribeImageCachesResponseBodyImageCachesEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The time when the image cache expires.
	ExpireDateTime *string `json:"ExpireDateTime,omitempty" xml:"ExpireDateTime,omitempty"`
	// The ID of the local snapshot. A temporary local snapshot is created after the instant image cache feature is enabled.
	FlashSnapshotId *string `json:"FlashSnapshotId,omitempty" xml:"FlashSnapshotId,omitempty"`
	// The ID of the image cache.
	ImageCacheId *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	// The name of the image cache.
	ImageCacheName *string `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	// The size of the image cache. Unit: GiB.
	ImageCacheSize *int32 `json:"ImageCacheSize,omitempty" xml:"ImageCacheSize,omitempty"`
	// The images that are contained in the image cache.
	Images []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The time when the image cache was last matched.
	LastMatchedTime *string `json:"LastMatchedTime,omitempty" xml:"LastMatchedTime,omitempty"`
	// The progress of creating the snapshot that is used to create the image cache.
	//
	// >  If the instant image cache feature is enabled, the system creates a temporary local snapshot and then a regular snapshot. This parameter indicates the progress of creating the regular snapshot.
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The region ID of the image cache.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the snapshot that is used to create the image cache.
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The status of the image cache. Valid values:
	//
	// *   Preparing: The image cache is being prepared.
	// *   Creating: The image is being created.
	// *   Ready: The image cache is ready for use.
	// *   Failed: The image cache failed to be created.
	// *   Updating: The image cache is being updated.
	// *   UpdateFailed: The image cache failed to be updated.
	//
	// The image cache is ready for use when it is in the Ready state.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the image cache.
	Tags []*DescribeImageCachesResponseBodyImageCachesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeImageCachesResponseBodyImageCaches) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesResponseBodyImageCaches) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetContainerGroupId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetCreationTime(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.CreationTime = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetEliminationStrategy(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.EliminationStrategy = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetEvents(v []*DescribeImageCachesResponseBodyImageCachesEvents) *DescribeImageCachesResponseBodyImageCaches {
	s.Events = v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetExpireDateTime(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.ExpireDateTime = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetFlashSnapshotId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.FlashSnapshotId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetImageCacheId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.ImageCacheId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetImageCacheName(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.ImageCacheName = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetImageCacheSize(v int32) *DescribeImageCachesResponseBodyImageCaches {
	s.ImageCacheSize = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetImages(v []*string) *DescribeImageCachesResponseBodyImageCaches {
	s.Images = v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetLastMatchedTime(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.LastMatchedTime = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetProgress(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.Progress = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetRegionId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.RegionId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetResourceGroupId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetSnapshotId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.SnapshotId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetStatus(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.Status = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetTags(v []*DescribeImageCachesResponseBodyImageCachesTags) *DescribeImageCachesResponseBodyImageCaches {
	s.Tags = v
	return s
}

type DescribeImageCachesResponseBodyImageCachesEvents struct {
	// The number of events.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The time when the event started.
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	// The time when the event ended.
	LastTimestamp *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The message of the event.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the event.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the event.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeImageCachesResponseBodyImageCachesEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesResponseBodyImageCachesEvents) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetCount(v int32) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.Count = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetFirstTimestamp(v string) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.FirstTimestamp = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetLastTimestamp(v string) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.LastTimestamp = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetMessage(v string) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.Message = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetName(v string) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.Name = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetType(v string) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.Type = &v
	return s
}

type DescribeImageCachesResponseBodyImageCachesTags struct {
	// The tag key of the image cache.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the image cache.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImageCachesResponseBodyImageCachesTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesResponseBodyImageCachesTags) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesResponseBodyImageCachesTags) SetKey(v string) *DescribeImageCachesResponseBodyImageCachesTags {
	s.Key = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesTags) SetValue(v string) *DescribeImageCachesResponseBodyImageCachesTags {
	s.Value = &v
	return s
}

type DescribeImageCachesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeImageCachesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImageCachesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesResponse) SetHeaders(v map[string]*string) *DescribeImageCachesResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageCachesResponse) SetStatusCode(v int32) *DescribeImageCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageCachesResponse) SetBody(v *DescribeImageCachesResponseBody) *DescribeImageCachesResponse {
	s.Body = v
	return s
}

type DescribeInstanceOpsRecordsRequest struct {
	// The ID of the container group.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The type of the O\&M task. Valid values:
	//
	// *   coredump
	// *   tcpdump
	OpsType      *string `json:"OpsType,omitempty" xml:"OpsType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the O\&M task.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeInstanceOpsRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceOpsRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceOpsRecordsRequest) SetContainerGroupId(v string) *DescribeInstanceOpsRecordsRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeInstanceOpsRecordsRequest) SetOpsType(v string) *DescribeInstanceOpsRecordsRequest {
	s.OpsType = &v
	return s
}

func (s *DescribeInstanceOpsRecordsRequest) SetOwnerAccount(v string) *DescribeInstanceOpsRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceOpsRecordsRequest) SetOwnerId(v int64) *DescribeInstanceOpsRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceOpsRecordsRequest) SetRegionId(v string) *DescribeInstanceOpsRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceOpsRecordsRequest) SetResourceOwnerAccount(v string) *DescribeInstanceOpsRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceOpsRecordsRequest) SetResourceOwnerId(v int64) *DescribeInstanceOpsRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeInstanceOpsRecordsResponseBody struct {
	// The details of the O\&M task.
	Records []*DescribeInstanceOpsRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceOpsRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceOpsRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceOpsRecordsResponseBody) SetRecords(v []*DescribeInstanceOpsRecordsResponseBodyRecords) *DescribeInstanceOpsRecordsResponseBody {
	s.Records = v
	return s
}

func (s *DescribeInstanceOpsRecordsResponseBody) SetRequestId(v string) *DescribeInstanceOpsRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceOpsRecordsResponseBodyRecords struct {
	// The time when the O\&M task was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the O\&M task expires.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The status of the O\&M task.
	OpsStatus *string `json:"OpsStatus,omitempty" xml:"OpsStatus,omitempty"`
	// The type of the O\&M task.
	OpsType *string `json:"OpsType,omitempty" xml:"OpsType,omitempty"`
	// The content of the O\&M result. The content is the download URL of the files that are generated for the O\&M task.
	ResultContent *string `json:"ResultContent,omitempty" xml:"ResultContent,omitempty"`
	// The type of the O\&M result. Valid value: OSS. This value indicates that the files generated for the O\&M task are saved to Object Storage Service (OSS) buckets.
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
}

func (s DescribeInstanceOpsRecordsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceOpsRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeInstanceOpsRecordsResponseBodyRecords) SetCreateTime(v string) *DescribeInstanceOpsRecordsResponseBodyRecords {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstanceOpsRecordsResponseBodyRecords) SetExpireTime(v string) *DescribeInstanceOpsRecordsResponseBodyRecords {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceOpsRecordsResponseBodyRecords) SetOpsStatus(v string) *DescribeInstanceOpsRecordsResponseBodyRecords {
	s.OpsStatus = &v
	return s
}

func (s *DescribeInstanceOpsRecordsResponseBodyRecords) SetOpsType(v string) *DescribeInstanceOpsRecordsResponseBodyRecords {
	s.OpsType = &v
	return s
}

func (s *DescribeInstanceOpsRecordsResponseBodyRecords) SetResultContent(v string) *DescribeInstanceOpsRecordsResponseBodyRecords {
	s.ResultContent = &v
	return s
}

func (s *DescribeInstanceOpsRecordsResponseBodyRecords) SetResultType(v string) *DescribeInstanceOpsRecordsResponseBodyRecords {
	s.ResultType = &v
	return s
}

type DescribeInstanceOpsRecordsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceOpsRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceOpsRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceOpsRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceOpsRecordsResponse) SetHeaders(v map[string]*string) *DescribeInstanceOpsRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceOpsRecordsResponse) SetStatusCode(v int32) *DescribeInstanceOpsRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceOpsRecordsResponse) SetBody(v *DescribeInstanceOpsRecordsResponseBody) *DescribeInstanceOpsRecordsResponse {
	s.Body = v
	return s
}

type DescribeMultiContainerGroupMetricRequest struct {
	// A JSON array that consists of the IDs of the elastic container instances. A maximum of 20 IDs are supported.
	ContainerGroupIds *string `json:"ContainerGroupIds,omitempty" xml:"ContainerGroupIds,omitempty"`
	// The type of the monitoring data. Set the value to summary, which indicates that records are returned.
	MetricType   *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the elastic container instances.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instances belong.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeMultiContainerGroupMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricRequest) SetContainerGroupIds(v string) *DescribeMultiContainerGroupMetricRequest {
	s.ContainerGroupIds = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetMetricType(v string) *DescribeMultiContainerGroupMetricRequest {
	s.MetricType = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetOwnerAccount(v string) *DescribeMultiContainerGroupMetricRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetOwnerId(v int64) *DescribeMultiContainerGroupMetricRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetRegionId(v string) *DescribeMultiContainerGroupMetricRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetResourceGroupId(v string) *DescribeMultiContainerGroupMetricRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetResourceOwnerAccount(v string) *DescribeMultiContainerGroupMetricRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetResourceOwnerId(v int64) *DescribeMultiContainerGroupMetricRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBody struct {
	// The collection of monitoring data of the elastic container instances.
	MonitorDatas []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatas `json:"MonitorDatas,omitempty" xml:"MonitorDatas,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBody) SetMonitorDatas(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatas) *DescribeMultiContainerGroupMetricResponseBody {
	s.MonitorDatas = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBody) SetRequestId(v string) *DescribeMultiContainerGroupMetricResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatas struct {
	// The ID of the elastic container instance.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The details about monitoring data.
	Records []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatas) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatas) SetContainerGroupId(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatas {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatas) SetRecords(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatas {
	s.Records = v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords struct {
	// The monitoring data of vCPUs.
	CPU *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU `json:"CPU,omitempty" xml:"CPU,omitempty" type:"Struct"`
	// The monitoring data of containers.
	Containers []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// The monitoring data of disks.
	Disk []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
	// The monitoring data of file system partitions.
	Filesystem []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem `json:"Filesystem,omitempty" xml:"Filesystem,omitempty" type:"Repeated"`
	// The monitoring data of the memory.
	Memory *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	// The monitoring data of the network.
	Network *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The time when the entry of monitoring data was collected. The time follows the RFC 3339 format.
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetCPU(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.CPU = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetContainers(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Containers = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetDisk(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Disk = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetFilesystem(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Filesystem = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetMemory(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Memory = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetNetwork(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Network = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetTimestamp(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Timestamp = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU struct {
	// The upper limit of vCPU usage. The calculation formula for this parameter: The number of vCPUs × 1000.
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The average load in the last 10 seconds.
	Load *int64 `json:"Load,omitempty" xml:"Load,omitempty"`
	// The cumulative usage of vCPUs.
	UsageCoreNanoSeconds *int64 `json:"UsageCoreNanoSeconds,omitempty" xml:"UsageCoreNanoSeconds,omitempty"`
	// The vCPU usage in the sampling window. Unit for the sampling window: nanoseconds.
	UsageNanoCores *int64 `json:"UsageNanoCores,omitempty" xml:"UsageNanoCores,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) SetLimit(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU {
	s.Limit = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) SetLoad(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU {
	s.Load = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) SetUsageCoreNanoSeconds(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU {
	s.UsageCoreNanoSeconds = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) SetUsageNanoCores(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU {
	s.UsageNanoCores = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers struct {
	// The vCPU monitoring data of the container.
	CPU *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU `json:"CPU,omitempty" xml:"CPU,omitempty" type:"Struct"`
	// The memory monitoring data of the container.
	Memory *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	// The name of the container.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) SetCPU(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers {
	s.CPU = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) SetMemory(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers {
	s.Memory = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) SetName(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers {
	s.Name = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU struct {
	// The upper limit of vCPU usage. The calculation formula for this parameter: The number of vCPUs × 1000.
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The average load in the last 10 seconds.
	Load *int64 `json:"Load,omitempty" xml:"Load,omitempty"`
	// The cumulative usage of vCPUs.
	UsageCoreNanoSeconds *int64 `json:"UsageCoreNanoSeconds,omitempty" xml:"UsageCoreNanoSeconds,omitempty"`
	// The vCPU usage in the sampling window. Unit for the sampling window: nanoseconds.
	UsageNanoCores *int64 `json:"UsageNanoCores,omitempty" xml:"UsageNanoCores,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) SetLimit(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU {
	s.Limit = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) SetLoad(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU {
	s.Load = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) SetUsageCoreNanoSeconds(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU {
	s.UsageCoreNanoSeconds = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) SetUsageNanoCores(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU {
	s.UsageNanoCores = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory struct {
	// The size of the available memory. Unit: bytes.
	AvailableBytes *int64 `json:"AvailableBytes,omitempty" xml:"AvailableBytes,omitempty"`
	// The size of the cache. Unit: bytes.
	Cache *int64 `json:"Cache,omitempty" xml:"Cache,omitempty"`
	// The size of the resident memory set, which indicates the size of the physical memory that is actually used. Unit: bytes.
	Rss *int64 `json:"Rss,omitempty" xml:"Rss,omitempty"`
	// The size of the used memory. Unit: bytes.
	UsageBytes *int64 `json:"UsageBytes,omitempty" xml:"UsageBytes,omitempty"`
	// The usage of the working set. Unit: bytes.
	WorkingSet *int64 `json:"WorkingSet,omitempty" xml:"WorkingSet,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) SetAvailableBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory {
	s.AvailableBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) SetCache(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory {
	s.Cache = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) SetRss(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory {
	s.Rss = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) SetUsageBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory {
	s.UsageBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) SetWorkingSet(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory {
	s.WorkingSet = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk struct {
	// The name of the disk.
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The amount of data that was read from the disk. Unit: bytes.
	ReadBytes *int64 `json:"ReadBytes,omitempty" xml:"ReadBytes,omitempty"`
	// This parameter is unavailable.
	ReadIo *int64 `json:"ReadIo,omitempty" xml:"ReadIo,omitempty"`
	// The amount of data that was written to the disk. Unit: bytes.
	WriteBytes *int64 `json:"WriteBytes,omitempty" xml:"WriteBytes,omitempty"`
	// This parameter is unavailable.
	WriteIo *int64 `json:"WriteIo,omitempty" xml:"WriteIo,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) SetDevice(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk {
	s.Device = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) SetReadBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk {
	s.ReadBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) SetReadIo(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk {
	s.ReadIo = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) SetWriteBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk {
	s.WriteBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) SetWriteIo(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk {
	s.WriteIo = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem struct {
	// The size of the available space.
	Available *int64 `json:"Available,omitempty" xml:"Available,omitempty"`
	// The total size of space.
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The name of the partition.
	FsName *string `json:"FsName,omitempty" xml:"FsName,omitempty"`
	// The size of used space.
	Usage *int64 `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) SetAvailable(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem {
	s.Available = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) SetCapacity(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem {
	s.Capacity = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) SetFsName(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem {
	s.FsName = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) SetUsage(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem {
	s.Usage = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory struct {
	// The size of the available memory. Unit: bytes.
	AvailableBytes *int64 `json:"AvailableBytes,omitempty" xml:"AvailableBytes,omitempty"`
	// The size of the cache. Unit: bytes.
	Cache *int64 `json:"Cache,omitempty" xml:"Cache,omitempty"`
	// The size of the resident memory set, which indicates the size of the physical memory that is actually used. Unit: bytes.
	Rss *int64 `json:"Rss,omitempty" xml:"Rss,omitempty"`
	// The size of the used memory. Unit: bytes.
	UsageBytes *int64 `json:"UsageBytes,omitempty" xml:"UsageBytes,omitempty"`
	// The usage of the working set. Unit: bytes.
	WorkingSet *int64 `json:"WorkingSet,omitempty" xml:"WorkingSet,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) SetAvailableBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory {
	s.AvailableBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) SetCache(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory {
	s.Cache = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) SetRss(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory {
	s.Rss = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) SetUsageBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory {
	s.UsageBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) SetWorkingSet(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory {
	s.WorkingSet = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork struct {
	// The monitoring data of network interface controllers (NICs).
	Interfaces []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces `json:"Interfaces,omitempty" xml:"Interfaces,omitempty" type:"Repeated"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork) SetInterfaces(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork {
	s.Interfaces = v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces struct {
	// The name of the NIC.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The total number of bytes received.
	RxBytes *int64 `json:"RxBytes,omitempty" xml:"RxBytes,omitempty"`
	// The number of packets that fail to be received.
	RxDrops *int64 `json:"RxDrops,omitempty" xml:"RxDrops,omitempty"`
	// The number of receipt errors.
	RxErrors *int64 `json:"RxErrors,omitempty" xml:"RxErrors,omitempty"`
	// The total number of packets received.
	RxPackets *int64 `json:"RxPackets,omitempty" xml:"RxPackets,omitempty"`
	// The total number of bytes sent.
	TxBytes *int64 `json:"TxBytes,omitempty" xml:"TxBytes,omitempty"`
	// The number of packets that fail to arrive at their destination.
	TxDrops *int64 `json:"TxDrops,omitempty" xml:"TxDrops,omitempty"`
	// The total number of sending errors.
	TxErrors *int64 `json:"TxErrors,omitempty" xml:"TxErrors,omitempty"`
	// The total number of packets sent.
	TxPackets *int64 `json:"TxPackets,omitempty" xml:"TxPackets,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetName(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.Name = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetRxBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.RxBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetRxDrops(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.RxDrops = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetRxErrors(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.RxErrors = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetRxPackets(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.RxPackets = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetTxBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.TxBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetTxDrops(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.TxDrops = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetTxErrors(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.TxErrors = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetTxPackets(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.TxPackets = &v
	return s
}

type DescribeMultiContainerGroupMetricResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMultiContainerGroupMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMultiContainerGroupMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponse) SetHeaders(v map[string]*string) *DescribeMultiContainerGroupMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponse) SetStatusCode(v int32) *DescribeMultiContainerGroupMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponse) SetBody(v *DescribeMultiContainerGroupMetricResponseBody) *DescribeMultiContainerGroupMetricResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetOwnerAccount(v string) *DescribeRegionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerAccount(v string) *DescribeRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// The details about the regions.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	// The recommended zones. Recommended zones are zones that have relatively sufficient resources in the current region.
	RecommendZones []*string `json:"RecommendZones,omitempty" xml:"RecommendZones,omitempty" type:"Repeated"`
	// The endpoint of the region.
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the zones.
	Zones []*string `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRecommendZones(v []*string) *DescribeRegionsResponseBodyRegions {
	s.RecommendZones = v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetZones(v []*string) *DescribeRegionsResponseBodyRegions {
	s.Zones = v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeVirtualNodesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The maximum number of resources that are allowed to return for this request. Default value: 20. Maximum value: 20.
	//
	// >  The number of resources to be returned is no greater than this parameter value.
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The token that is used for the next query. If this parameter is empty, all results have been returned.
	//
	// You do not need to specify this parameter in the first request. You can obtain the token from the result returned by the previous request.
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the virtual nodes.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the virtual node. Valid values:
	//
	// *   Pending
	// *   Ready
	// *   Failed
	Status *string                           `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag    []*DescribeVirtualNodesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The IDs of the virtual nodes. You can specify up to 20 IDs. Each ID must be a string in JSON format.
	VirtualNodeIds *string `json:"VirtualNodeIds,omitempty" xml:"VirtualNodeIds,omitempty"`
	// The name of the virtual node.
	VirtualNodeName *string `json:"VirtualNodeName,omitempty" xml:"VirtualNodeName,omitempty"`
}

func (s DescribeVirtualNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVirtualNodesRequest) SetClientToken(v string) *DescribeVirtualNodesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetLimit(v int64) *DescribeVirtualNodesRequest {
	s.Limit = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetNextToken(v string) *DescribeVirtualNodesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetOwnerAccount(v string) *DescribeVirtualNodesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetOwnerId(v int64) *DescribeVirtualNodesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetRegionId(v string) *DescribeVirtualNodesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetResourceGroupId(v string) *DescribeVirtualNodesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetResourceOwnerAccount(v string) *DescribeVirtualNodesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetResourceOwnerId(v int64) *DescribeVirtualNodesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetStatus(v string) *DescribeVirtualNodesRequest {
	s.Status = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetTag(v []*DescribeVirtualNodesRequestTag) *DescribeVirtualNodesRequest {
	s.Tag = v
	return s
}

func (s *DescribeVirtualNodesRequest) SetVirtualNodeIds(v string) *DescribeVirtualNodesRequest {
	s.VirtualNodeIds = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetVirtualNodeName(v string) *DescribeVirtualNodesRequest {
	s.VirtualNodeName = &v
	return s
}

type DescribeVirtualNodesRequestTag struct {
	// The key of tag N.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVirtualNodesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualNodesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeVirtualNodesRequestTag) SetKey(v string) *DescribeVirtualNodesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeVirtualNodesRequestTag) SetValue(v string) *DescribeVirtualNodesRequestTag {
	s.Value = &v
	return s
}

type DescribeVirtualNodesResponseBody struct {
	// The token that was returned for the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of virtual nodes that were queried.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The virtual nodes that were queried.
	VirtualNodes []*DescribeVirtualNodesResponseBodyVirtualNodes `json:"VirtualNodes,omitempty" xml:"VirtualNodes,omitempty" type:"Repeated"`
}

func (s DescribeVirtualNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVirtualNodesResponseBody) SetNextToken(v string) *DescribeVirtualNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVirtualNodesResponseBody) SetRequestId(v string) *DescribeVirtualNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBody) SetTotalCount(v int32) *DescribeVirtualNodesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVirtualNodesResponseBody) SetVirtualNodes(v []*DescribeVirtualNodesResponseBodyVirtualNodes) *DescribeVirtualNodesResponseBody {
	s.VirtualNodes = v
	return s
}

type DescribeVirtualNodesResponseBodyVirtualNodes struct {
	// The ID of the Kubernetes cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of vCPUs. This parameter is unavailable.
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The time when the virtual node was created. The time follows the RFC 3339 standard and is displayed in UTC.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// An array of the events.
	Events []*DescribeVirtualNodesResponseBodyVirtualNodesEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The public IP address.
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The internal IP address.
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The memory size. This parameter is unavailable.
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The region ID of the virtual node.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the security group.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The status of the virtual node. Valid values:
	//
	// *   Pending
	// *   Ready
	// *   Failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// An array of tags.
	Tags       []*DescribeVirtualNodesResponseBodyVirtualNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	VSwitchIds []*string                                           `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the virtual node.
	VirtualNodeId *string `json:"VirtualNodeId,omitempty" xml:"VirtualNodeId,omitempty"`
	// The name of the virtual node.
	VirtualNodeName            *string `json:"VirtualNodeName,omitempty" xml:"VirtualNodeName,omitempty"`
	VirtualNodeSecurityGroupId *string `json:"VirtualNodeSecurityGroupId,omitempty" xml:"VirtualNodeSecurityGroupId,omitempty"`
	VirtualNodeVSwitchId       *string `json:"VirtualNodeVSwitchId,omitempty" xml:"VirtualNodeVSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone.
	ZoneId  *string   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneIds []*string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty" type:"Repeated"`
}

func (s DescribeVirtualNodesResponseBodyVirtualNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualNodesResponseBodyVirtualNodes) GoString() string {
	return s.String()
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetClusterId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.ClusterId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetCpu(v float32) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.Cpu = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetCreationTime(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.CreationTime = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetEvents(v []*DescribeVirtualNodesResponseBodyVirtualNodesEvents) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.Events = v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetInternetIp(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.InternetIp = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetIntranetIp(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.IntranetIp = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetMemory(v float32) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.Memory = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetRegionId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.RegionId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetResourceGroupId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetSecurityGroupId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetStatus(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.Status = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetTags(v []*DescribeVirtualNodesResponseBodyVirtualNodesTags) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.Tags = v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetVSwitchIds(v []*string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.VSwitchIds = v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetVirtualNodeId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.VirtualNodeId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetVirtualNodeName(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.VirtualNodeName = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetVirtualNodeSecurityGroupId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.VirtualNodeSecurityGroupId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetVirtualNodeVSwitchId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.VirtualNodeVSwitchId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetVpcId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.VpcId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetZoneId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.ZoneId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetZoneIds(v []*string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.ZoneIds = v
	return s
}

type DescribeVirtualNodesResponseBodyVirtualNodesEvents struct {
	// The number of events.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The time when the event started.
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	// The time when the event ended.
	LastTimestamp *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The message of the event.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the object to which the event belongs.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the event.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The type of the event. Valid values:
	//
	// *   Normal
	// *   Warning
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVirtualNodesResponseBodyVirtualNodesEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualNodesResponseBodyVirtualNodesEvents) GoString() string {
	return s.String()
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesEvents) SetCount(v int32) *DescribeVirtualNodesResponseBodyVirtualNodesEvents {
	s.Count = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesEvents) SetFirstTimestamp(v string) *DescribeVirtualNodesResponseBodyVirtualNodesEvents {
	s.FirstTimestamp = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesEvents) SetLastTimestamp(v string) *DescribeVirtualNodesResponseBodyVirtualNodesEvents {
	s.LastTimestamp = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesEvents) SetMessage(v string) *DescribeVirtualNodesResponseBodyVirtualNodesEvents {
	s.Message = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesEvents) SetName(v string) *DescribeVirtualNodesResponseBodyVirtualNodesEvents {
	s.Name = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesEvents) SetReason(v string) *DescribeVirtualNodesResponseBodyVirtualNodesEvents {
	s.Reason = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesEvents) SetType(v string) *DescribeVirtualNodesResponseBodyVirtualNodesEvents {
	s.Type = &v
	return s
}

type DescribeVirtualNodesResponseBodyVirtualNodesTags struct {
	// The tag key of the virtual node.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the virtual node.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVirtualNodesResponseBodyVirtualNodesTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualNodesResponseBodyVirtualNodesTags) GoString() string {
	return s.String()
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesTags) SetKey(v string) *DescribeVirtualNodesResponseBodyVirtualNodesTags {
	s.Key = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesTags) SetValue(v string) *DescribeVirtualNodesResponseBodyVirtualNodesTags {
	s.Value = &v
	return s
}

type DescribeVirtualNodesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVirtualNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVirtualNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVirtualNodesResponse) SetHeaders(v map[string]*string) *DescribeVirtualNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVirtualNodesResponse) SetStatusCode(v int32) *DescribeVirtualNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVirtualNodesResponse) SetBody(v *DescribeVirtualNodesResponseBody) *DescribeVirtualNodesResponse {
	s.Body = v
	return s
}

type ExecContainerCommandRequest struct {
	// The commands to run in the container. You can specify up to 20 commands. Each command can be up to 256 characters in length.
	//
	// The strings must be in the JSON format. Example: `["/bin/sh", "-c", "ls -a"]`.
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The ID of the elastic container instance.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The name of the container.
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where an elastic container instance is created.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to read the commands from standard input (stdin). Default value: true.
	Stdin *bool `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	// Specifies whether to run the command immediately and return the result. Default value: false.
	//
	// If you set this parameter to true, set the value of TTY to false.
	Sync *bool `json:"Sync,omitempty" xml:"Sync,omitempty"`
	// Specifies whether to enable interaction. Default value: false.
	//
	// If the command is a /bin/bash command, set the value to true.
	TTY *bool `json:"TTY,omitempty" xml:"TTY,omitempty"`
}

func (s ExecContainerCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecContainerCommandRequest) GoString() string {
	return s.String()
}

func (s *ExecContainerCommandRequest) SetCommand(v string) *ExecContainerCommandRequest {
	s.Command = &v
	return s
}

func (s *ExecContainerCommandRequest) SetContainerGroupId(v string) *ExecContainerCommandRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *ExecContainerCommandRequest) SetContainerName(v string) *ExecContainerCommandRequest {
	s.ContainerName = &v
	return s
}

func (s *ExecContainerCommandRequest) SetOwnerAccount(v string) *ExecContainerCommandRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ExecContainerCommandRequest) SetOwnerId(v int64) *ExecContainerCommandRequest {
	s.OwnerId = &v
	return s
}

func (s *ExecContainerCommandRequest) SetRegionId(v string) *ExecContainerCommandRequest {
	s.RegionId = &v
	return s
}

func (s *ExecContainerCommandRequest) SetResourceOwnerAccount(v string) *ExecContainerCommandRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ExecContainerCommandRequest) SetResourceOwnerId(v int64) *ExecContainerCommandRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ExecContainerCommandRequest) SetStdin(v bool) *ExecContainerCommandRequest {
	s.Stdin = &v
	return s
}

func (s *ExecContainerCommandRequest) SetSync(v bool) *ExecContainerCommandRequest {
	s.Sync = &v
	return s
}

func (s *ExecContainerCommandRequest) SetTTY(v bool) *ExecContainerCommandRequest {
	s.TTY = &v
	return s
}

type ExecContainerCommandResponseBody struct {
	// The HTTP URL. You can use this URL to enter the container within 30 seconds after this operation is called. For more information, see [Use and integrate the Elastic Container Instance terminal](~~202846~~).
	HttpUrl *string `json:"HttpUrl,omitempty" xml:"HttpUrl,omitempty"`
	// The unique ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The output of the command. This parameter is returned when Sync is set to true.
	SyncResponse *string `json:"SyncResponse,omitempty" xml:"SyncResponse,omitempty"`
	// The WebSocket URL. You can use this URL to establish a WebSocket connection with the container.
	WebSocketUri *string `json:"WebSocketUri,omitempty" xml:"WebSocketUri,omitempty"`
}

func (s ExecContainerCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecContainerCommandResponseBody) GoString() string {
	return s.String()
}

func (s *ExecContainerCommandResponseBody) SetHttpUrl(v string) *ExecContainerCommandResponseBody {
	s.HttpUrl = &v
	return s
}

func (s *ExecContainerCommandResponseBody) SetRequestId(v string) *ExecContainerCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecContainerCommandResponseBody) SetSyncResponse(v string) *ExecContainerCommandResponseBody {
	s.SyncResponse = &v
	return s
}

func (s *ExecContainerCommandResponseBody) SetWebSocketUri(v string) *ExecContainerCommandResponseBody {
	s.WebSocketUri = &v
	return s
}

type ExecContainerCommandResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExecContainerCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecContainerCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecContainerCommandResponse) GoString() string {
	return s.String()
}

func (s *ExecContainerCommandResponse) SetHeaders(v map[string]*string) *ExecContainerCommandResponse {
	s.Headers = v
	return s
}

func (s *ExecContainerCommandResponse) SetStatusCode(v int32) *ExecContainerCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecContainerCommandResponse) SetBody(v *ExecContainerCommandResponseBody) *ExecContainerCommandResponse {
	s.Body = v
	return s
}

type ListUsageRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsageRequest) GoString() string {
	return s.String()
}

func (s *ListUsageRequest) SetOwnerAccount(v string) *ListUsageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListUsageRequest) SetOwnerId(v int64) *ListUsageRequest {
	s.OwnerId = &v
	return s
}

func (s *ListUsageRequest) SetRegionId(v string) *ListUsageRequest {
	s.RegionId = &v
	return s
}

func (s *ListUsageRequest) SetResourceOwnerAccount(v string) *ListUsageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListUsageRequest) SetResourceOwnerId(v int64) *ListUsageRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListUsageResponseBody struct {
	// The information about the used amounts and upper limits of privileges and quotas that you have in a specified region. The information contains the following items:
	//
	// *   UsedCpu: the number of existing vCPUs.
	// *   MaxCpu: the maximum number of vCPUs.
	// *   UsedImageCacheCount: the number of existing image caches.
	// *   MaxImageCacheCount: the maximum number of image caches.
	Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsageResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsageResponseBody) SetAttributes(v map[string]interface{}) *ListUsageResponseBody {
	s.Attributes = v
	return s
}

func (s *ListUsageResponseBody) SetRequestId(v string) *ListUsageResponseBody {
	s.RequestId = &v
	return s
}

type ListUsageResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsageResponse) GoString() string {
	return s.String()
}

func (s *ListUsageResponse) SetHeaders(v map[string]*string) *ListUsageResponse {
	s.Headers = v
	return s
}

func (s *ListUsageResponse) SetStatusCode(v int32) *ListUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsageResponse) SetBody(v *ListUsageResponseBody) *ListUsageResponse {
	s.Body = v
	return s
}

type ResizeContainerGroupVolumeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure the idempotence of a request](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the elastic container instance for which you want to scale out the volume.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The size of the disk after the scale-out activity is complete. Unit: GiB. Valid values:
	//
	// *   20 to 32768 for ultra disks (cloud_efficiency)
	// *   20 to 32768 for standard SSD (cloud_ssd)
	// *   20 to 32768 for enhanced SSD (cloud_essd)
	// *   5 to 2000 for basic disk (cloud)
	//
	// >  The size of the disk after the scale-out activity is complete must be greater than the size of the disk before the scale-out activity is complete. If the size that you specify for the disk after the scale-out activity is complete is equal to the size of the disk before the scale-out activity is complete, only the file system is scaled out.
	NewSize      *int64  `json:"NewSize,omitempty" xml:"NewSize,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the elastic container instance.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the volume. Only volumes of Alibaba Cloud disk type can be scaled out.
	VolumeName *string `json:"VolumeName,omitempty" xml:"VolumeName,omitempty"`
}

func (s ResizeContainerGroupVolumeRequest) String() string {
	return tea.Prettify(s)
}

func (s ResizeContainerGroupVolumeRequest) GoString() string {
	return s.String()
}

func (s *ResizeContainerGroupVolumeRequest) SetClientToken(v string) *ResizeContainerGroupVolumeRequest {
	s.ClientToken = &v
	return s
}

func (s *ResizeContainerGroupVolumeRequest) SetContainerGroupId(v string) *ResizeContainerGroupVolumeRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *ResizeContainerGroupVolumeRequest) SetNewSize(v int64) *ResizeContainerGroupVolumeRequest {
	s.NewSize = &v
	return s
}

func (s *ResizeContainerGroupVolumeRequest) SetOwnerAccount(v string) *ResizeContainerGroupVolumeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResizeContainerGroupVolumeRequest) SetOwnerId(v int64) *ResizeContainerGroupVolumeRequest {
	s.OwnerId = &v
	return s
}

func (s *ResizeContainerGroupVolumeRequest) SetRegionId(v string) *ResizeContainerGroupVolumeRequest {
	s.RegionId = &v
	return s
}

func (s *ResizeContainerGroupVolumeRequest) SetResourceOwnerAccount(v string) *ResizeContainerGroupVolumeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResizeContainerGroupVolumeRequest) SetResourceOwnerId(v int64) *ResizeContainerGroupVolumeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResizeContainerGroupVolumeRequest) SetVolumeName(v string) *ResizeContainerGroupVolumeRequest {
	s.VolumeName = &v
	return s
}

type ResizeContainerGroupVolumeResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeContainerGroupVolumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeContainerGroupVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeContainerGroupVolumeResponseBody) SetRequestId(v string) *ResizeContainerGroupVolumeResponseBody {
	s.RequestId = &v
	return s
}

type ResizeContainerGroupVolumeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResizeContainerGroupVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResizeContainerGroupVolumeResponse) String() string {
	return tea.Prettify(s)
}

func (s ResizeContainerGroupVolumeResponse) GoString() string {
	return s.String()
}

func (s *ResizeContainerGroupVolumeResponse) SetHeaders(v map[string]*string) *ResizeContainerGroupVolumeResponse {
	s.Headers = v
	return s
}

func (s *ResizeContainerGroupVolumeResponse) SetStatusCode(v int32) *ResizeContainerGroupVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeContainerGroupVolumeResponse) SetBody(v *ResizeContainerGroupVolumeResponseBody) *ResizeContainerGroupVolumeResponse {
	s.Body = v
	return s
}

type RestartContainerGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure the idempotence of a request?](~~25693~~)
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the elastic container instance.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RestartContainerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartContainerGroupRequest) GoString() string {
	return s.String()
}

func (s *RestartContainerGroupRequest) SetClientToken(v string) *RestartContainerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartContainerGroupRequest) SetContainerGroupId(v string) *RestartContainerGroupRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *RestartContainerGroupRequest) SetOwnerAccount(v string) *RestartContainerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartContainerGroupRequest) SetOwnerId(v int64) *RestartContainerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartContainerGroupRequest) SetRegionId(v string) *RestartContainerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RestartContainerGroupRequest) SetResourceOwnerAccount(v string) *RestartContainerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartContainerGroupRequest) SetResourceOwnerId(v int64) *RestartContainerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

type RestartContainerGroupResponseBody struct {
	// The ID of the request. The value is unique.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartContainerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartContainerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RestartContainerGroupResponseBody) SetRequestId(v string) *RestartContainerGroupResponseBody {
	s.RequestId = &v
	return s
}

type RestartContainerGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestartContainerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartContainerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartContainerGroupResponse) GoString() string {
	return s.String()
}

func (s *RestartContainerGroupResponse) SetHeaders(v map[string]*string) *RestartContainerGroupResponse {
	s.Headers = v
	return s
}

func (s *RestartContainerGroupResponse) SetStatusCode(v int32) *RestartContainerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartContainerGroupResponse) SetBody(v *RestartContainerGroupResponseBody) *RestartContainerGroupResponse {
	s.Body = v
	return s
}

type UpdateContainerGroupRequest struct {
	DnsConfig       *UpdateContainerGroupRequestDnsConfig         `json:"DnsConfig,omitempty" xml:"DnsConfig,omitempty" type:"Struct"`
	AcrRegistryInfo []*UpdateContainerGroupRequestAcrRegistryInfo `json:"AcrRegistryInfo,omitempty" xml:"AcrRegistryInfo,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency](~~25693~~).
	ClientToken *string                                 `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Container   []*UpdateContainerGroupRequestContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Repeated"`
	// The ID of the elastic container instance that you want to update.
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The number of vCPUs allocated to the elastic container instance.
	Cpu                     *float32                                              `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	ImageRegistryCredential []*UpdateContainerGroupRequestImageRegistryCredential `json:"ImageRegistryCredential,omitempty" xml:"ImageRegistryCredential,omitempty" type:"Repeated"`
	InitContainer           []*UpdateContainerGroupRequestInitContainer           `json:"InitContainer,omitempty" xml:"InitContainer,omitempty" type:"Repeated"`
	// The memory size allocated to the elastic container instance. Unit: GiB.
	Memory       *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	OwnerAccount *string  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The restart policy of the elastic container instance. Valid values:
	//
	// *   Always: Always restarts the instance.
	// *   Never: Never restarts the instance.
	// *   OnFailure: Restarts the instance when the last start of the instance failed.
	RestartPolicy *string `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	// The list of tags that bound to the instance.
	Tag []*UpdateContainerGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The update type. Valid values:
	//
	// - RenewUpdate: full update. You must specify all relevant parameters to update the instance. For a parameter of the List type, you must specify all the items contained in the parameter even if you want to update only some of the items. For a parameter of the struct type, you must specify all the members if you want to update only some of the members.
	// - IncrementalUpdate: incremental update. You can specify only the parameter that needs to be updated. Other related parameters remain unchanged.
	//
	//  Default value: RenewUpdate.
	UpdateType *string `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
	// The list of data volumes.
	Volume []*UpdateContainerGroupRequestVolume `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequest) SetDnsConfig(v *UpdateContainerGroupRequestDnsConfig) *UpdateContainerGroupRequest {
	s.DnsConfig = v
	return s
}

func (s *UpdateContainerGroupRequest) SetAcrRegistryInfo(v []*UpdateContainerGroupRequestAcrRegistryInfo) *UpdateContainerGroupRequest {
	s.AcrRegistryInfo = v
	return s
}

func (s *UpdateContainerGroupRequest) SetClientToken(v string) *UpdateContainerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetContainer(v []*UpdateContainerGroupRequestContainer) *UpdateContainerGroupRequest {
	s.Container = v
	return s
}

func (s *UpdateContainerGroupRequest) SetContainerGroupId(v string) *UpdateContainerGroupRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetCpu(v float32) *UpdateContainerGroupRequest {
	s.Cpu = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetImageRegistryCredential(v []*UpdateContainerGroupRequestImageRegistryCredential) *UpdateContainerGroupRequest {
	s.ImageRegistryCredential = v
	return s
}

func (s *UpdateContainerGroupRequest) SetInitContainer(v []*UpdateContainerGroupRequestInitContainer) *UpdateContainerGroupRequest {
	s.InitContainer = v
	return s
}

func (s *UpdateContainerGroupRequest) SetMemory(v float32) *UpdateContainerGroupRequest {
	s.Memory = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetOwnerAccount(v string) *UpdateContainerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetOwnerId(v int64) *UpdateContainerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetRegionId(v string) *UpdateContainerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetResourceGroupId(v string) *UpdateContainerGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetResourceOwnerAccount(v string) *UpdateContainerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetResourceOwnerId(v int64) *UpdateContainerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetRestartPolicy(v string) *UpdateContainerGroupRequest {
	s.RestartPolicy = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetTag(v []*UpdateContainerGroupRequestTag) *UpdateContainerGroupRequest {
	s.Tag = v
	return s
}

func (s *UpdateContainerGroupRequest) SetUpdateType(v string) *UpdateContainerGroupRequest {
	s.UpdateType = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetVolume(v []*UpdateContainerGroupRequestVolume) *UpdateContainerGroupRequest {
	s.Volume = v
	return s
}

type UpdateContainerGroupRequestDnsConfig struct {
	// The list of IP addresses of DNS servers.
	NameServer []*string                                     `json:"NameServer,omitempty" xml:"NameServer,omitempty" type:"Repeated"`
	Option     []*UpdateContainerGroupRequestDnsConfigOption `json:"Option,omitempty" xml:"Option,omitempty" type:"Repeated"`
	// The list of DNS search domains.
	Search []*string `json:"Search,omitempty" xml:"Search,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestDnsConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestDnsConfig) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestDnsConfig) SetNameServer(v []*string) *UpdateContainerGroupRequestDnsConfig {
	s.NameServer = v
	return s
}

func (s *UpdateContainerGroupRequestDnsConfig) SetOption(v []*UpdateContainerGroupRequestDnsConfigOption) *UpdateContainerGroupRequestDnsConfig {
	s.Option = v
	return s
}

func (s *UpdateContainerGroupRequestDnsConfig) SetSearch(v []*string) *UpdateContainerGroupRequestDnsConfig {
	s.Search = v
	return s
}

type UpdateContainerGroupRequestDnsConfigOption struct {
	// The option name of DNS configurations.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The option value of DNS configurations.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateContainerGroupRequestDnsConfigOption) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestDnsConfigOption) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestDnsConfigOption) SetName(v string) *UpdateContainerGroupRequestDnsConfigOption {
	s.Name = &v
	return s
}

func (s *UpdateContainerGroupRequestDnsConfigOption) SetValue(v string) *UpdateContainerGroupRequestDnsConfigOption {
	s.Value = &v
	return s
}

type UpdateContainerGroupRequestAcrRegistryInfo struct {
	Domain []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
	// The ID of Container Registry Enterprise Edition instance N.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of Container Registry Enterprise Edition instance N.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID of Container Registry Enterprise Edition instance N.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateContainerGroupRequestAcrRegistryInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestAcrRegistryInfo) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestAcrRegistryInfo) SetDomain(v []*string) *UpdateContainerGroupRequestAcrRegistryInfo {
	s.Domain = v
	return s
}

func (s *UpdateContainerGroupRequestAcrRegistryInfo) SetInstanceId(v string) *UpdateContainerGroupRequestAcrRegistryInfo {
	s.InstanceId = &v
	return s
}

func (s *UpdateContainerGroupRequestAcrRegistryInfo) SetInstanceName(v string) *UpdateContainerGroupRequestAcrRegistryInfo {
	s.InstanceName = &v
	return s
}

func (s *UpdateContainerGroupRequestAcrRegistryInfo) SetRegionId(v string) *UpdateContainerGroupRequestAcrRegistryInfo {
	s.RegionId = &v
	return s
}

type UpdateContainerGroupRequestContainer struct {
	LivenessProbe   *UpdateContainerGroupRequestContainerLivenessProbe   `json:"LivenessProbe,omitempty" xml:"LivenessProbe,omitempty" require:"true" type:"Struct"`
	ReadinessProbe  *UpdateContainerGroupRequestContainerReadinessProbe  `json:"ReadinessProbe,omitempty" xml:"ReadinessProbe,omitempty" require:"true" type:"Struct"`
	SecurityContext *UpdateContainerGroupRequestContainerSecurityContext `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" require:"true" type:"Struct"`
	Arg             []*string                                            `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	Command         []*string                                            `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// The number of vCPUs that you want to allocate to container N.
	Cpu            *float32                                              `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	EnvironmentVar []*UpdateContainerGroupRequestContainerEnvironmentVar `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	// The number of GPUs that you want to allocate to container N.
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The image of container N.
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image pull policy of container N.
	ImagePullPolicy               *string   `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	LifecyclePostStartHandlerExec []*string `json:"LifecyclePostStartHandlerExec,omitempty" xml:"LifecyclePostStartHandlerExec,omitempty" type:"Repeated"`
	// The IP address of the host that receives HTTP GET requests when you use HTTP requests to specify the postStart callback function.
	LifecyclePostStartHandlerHttpGetHost        *string                                                                            `json:"LifecyclePostStartHandlerHttpGetHost,omitempty" xml:"LifecyclePostStartHandlerHttpGetHost,omitempty"`
	LifecyclePostStartHandlerHttpGetHttpHeaders []*UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders `json:"LifecyclePostStartHandlerHttpGetHttpHeaders,omitempty" xml:"LifecyclePostStartHandlerHttpGetHttpHeaders,omitempty" type:"Repeated"`
	// The port to which HTTP GET requests are sent when you use HTTP requests to specify the postStart callback function.
	LifecyclePostStartHandlerHttpGetPath *string `json:"LifecyclePostStartHandlerHttpGetPath,omitempty" xml:"LifecyclePostStartHandlerHttpGetPath,omitempty"`
	// The port to which HTTP GET requests are sent when you use HTTP requests to specify the postStart callback function.
	LifecyclePostStartHandlerHttpGetPort *int32 `json:"LifecyclePostStartHandlerHttpGetPort,omitempty" xml:"LifecyclePostStartHandlerHttpGetPort,omitempty"`
	// The path to which HTTP GET requests are sent when you use HTTP requests to specify the postStart callback function.
	LifecyclePostStartHandlerHttpGetScheme *string `json:"LifecyclePostStartHandlerHttpGetScheme,omitempty" xml:"LifecyclePostStartHandlerHttpGetScheme,omitempty"`
	// The host IP address of TCP socket probes when you use TCP sockets to specify the postStart callback function.
	LifecyclePostStartHandlerTcpSocketHost *string `json:"LifecyclePostStartHandlerTcpSocketHost,omitempty" xml:"LifecyclePostStartHandlerTcpSocketHost,omitempty"`
	// The port of TCP socket probes when you use TCP sockets to specify the postStart callback function.
	LifecyclePostStartHandlerTcpSocketPort *int32    `json:"LifecyclePostStartHandlerTcpSocketPort,omitempty" xml:"LifecyclePostStartHandlerTcpSocketPort,omitempty"`
	LifecyclePreStopHandlerExec            []*string `json:"LifecyclePreStopHandlerExec,omitempty" xml:"LifecyclePreStopHandlerExec,omitempty" type:"Repeated"`
	// The IP address of the host that receives HTTP GET requests when you use HTTP requests to specify the preStop callback function.
	LifecyclePreStopHandlerHttpGetHost       *string                                                                         `json:"LifecyclePreStopHandlerHttpGetHost,omitempty" xml:"LifecyclePreStopHandlerHttpGetHost,omitempty"`
	LifecyclePreStopHandlerHttpGetHttpHeader []*UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader `json:"LifecyclePreStopHandlerHttpGetHttpHeader,omitempty" xml:"LifecyclePreStopHandlerHttpGetHttpHeader,omitempty" type:"Repeated"`
	// The path to which HTTP GET requests are sent when you use HTTP requests to specify the preStop callback function.
	LifecyclePreStopHandlerHttpGetPath *string `json:"LifecyclePreStopHandlerHttpGetPath,omitempty" xml:"LifecyclePreStopHandlerHttpGetPath,omitempty"`
	// The port to which HTTP GET requests are sent when you use HTTP requests to specify the preStop callback function.
	LifecyclePreStopHandlerHttpGetPort *int32 `json:"LifecyclePreStopHandlerHttpGetPort,omitempty" xml:"LifecyclePreStopHandlerHttpGetPort,omitempty"`
	// The protocol type of HTTP GET requests when you use HTTP requests to specify the preStop callback function. Valid values:
	//
	// - HTTP
	// - HTTPS
	LifecyclePreStopHandlerHttpGetScheme *string `json:"LifecyclePreStopHandlerHttpGetScheme,omitempty" xml:"LifecyclePreStopHandlerHttpGetScheme,omitempty"`
	// The host IP address of TCP socket probes when you use TCP sockets to specify the preStop callback function.
	LifecyclePreStopHandlerTcpSocketHost *string `json:"LifecyclePreStopHandlerTcpSocketHost,omitempty" xml:"LifecyclePreStopHandlerTcpSocketHost,omitempty"`
	// The port of TCP socket probes when you use TCP sockets to specify the preStop callback function.
	LifecyclePreStopHandlerTcpSocketPort *int32 `json:"LifecyclePreStopHandlerTcpSocketPort,omitempty" xml:"LifecyclePreStopHandlerTcpSocketPort,omitempty"`
	// The memory size of container N.
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of container N.
	Name *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	Port []*UpdateContainerGroupRequestContainerPort `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
	// Specifies whether container N allocates buffer resources to standard input streams when the container runs. If you do not specify this parameter, an end-of-file (EOF) error may occur. Default value: false.
	Stdin *bool `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	// Specifies whether to maintain standard input streams connected after a client is disconnected. If Stdin is set to true, standard input streams remain connected during multiple sessions. If Container.N.StdinOnce is set to true, standard input streams are connected after the container is started and remain idle until a client is connected to receive data. After the client is disconnected, streams are also disconnected and remain in the disconnected state until the container is started again.
	StdinOnce *bool `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	// Specifies whether to enable interaction. Default value: false. If the command is a /bin/bash command, set the value to true.
	Tty         *bool                                              `json:"Tty,omitempty" xml:"Tty,omitempty"`
	VolumeMount []*UpdateContainerGroupRequestContainerVolumeMount `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
	// The working directory of container N.
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s UpdateContainerGroupRequestContainer) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainer) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainer) SetLivenessProbe(v *UpdateContainerGroupRequestContainerLivenessProbe) *UpdateContainerGroupRequestContainer {
	s.LivenessProbe = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetReadinessProbe(v *UpdateContainerGroupRequestContainerReadinessProbe) *UpdateContainerGroupRequestContainer {
	s.ReadinessProbe = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetSecurityContext(v *UpdateContainerGroupRequestContainerSecurityContext) *UpdateContainerGroupRequestContainer {
	s.SecurityContext = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetArg(v []*string) *UpdateContainerGroupRequestContainer {
	s.Arg = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetCommand(v []*string) *UpdateContainerGroupRequestContainer {
	s.Command = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetCpu(v float32) *UpdateContainerGroupRequestContainer {
	s.Cpu = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetEnvironmentVar(v []*UpdateContainerGroupRequestContainerEnvironmentVar) *UpdateContainerGroupRequestContainer {
	s.EnvironmentVar = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetGpu(v int32) *UpdateContainerGroupRequestContainer {
	s.Gpu = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetImage(v string) *UpdateContainerGroupRequestContainer {
	s.Image = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetImagePullPolicy(v string) *UpdateContainerGroupRequestContainer {
	s.ImagePullPolicy = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerExec(v []*string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerExec = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetHost(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetHost = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetHttpHeaders(v []*UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetHttpHeaders = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetPath(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetPath = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetPort(v int32) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetPort = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetScheme(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetScheme = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerTcpSocketHost(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerTcpSocketHost = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerTcpSocketPort(v int32) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerTcpSocketPort = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerExec(v []*string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerExec = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetHost(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetHost = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetHttpHeader(v []*UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetHttpHeader = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetPath(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetPath = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetPort(v int32) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetPort = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetScheme(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetScheme = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerTcpSocketHost(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerTcpSocketHost = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerTcpSocketPort(v int32) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerTcpSocketPort = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetMemory(v float32) *UpdateContainerGroupRequestContainer {
	s.Memory = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetName(v string) *UpdateContainerGroupRequestContainer {
	s.Name = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetPort(v []*UpdateContainerGroupRequestContainerPort) *UpdateContainerGroupRequestContainer {
	s.Port = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetStdin(v bool) *UpdateContainerGroupRequestContainer {
	s.Stdin = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetStdinOnce(v bool) *UpdateContainerGroupRequestContainer {
	s.StdinOnce = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetTty(v bool) *UpdateContainerGroupRequestContainer {
	s.Tty = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetVolumeMount(v []*UpdateContainerGroupRequestContainerVolumeMount) *UpdateContainerGroupRequestContainer {
	s.VolumeMount = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetWorkingDir(v string) *UpdateContainerGroupRequestContainer {
	s.WorkingDir = &v
	return s
}

type UpdateContainerGroupRequestContainerLivenessProbe struct {
	Exec                *UpdateContainerGroupRequestContainerLivenessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" require:"true" type:"Struct"`
	FailureThreshold    *int32                                                      `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *UpdateContainerGroupRequestContainerLivenessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" require:"true" type:"Struct"`
	InitialDelaySeconds *int32                                                      `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                      `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                      `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *UpdateContainerGroupRequestContainerLivenessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" require:"true" type:"Struct"`
	TimeoutSeconds      *int32                                                      `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s UpdateContainerGroupRequestContainerLivenessProbe) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLivenessProbe) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetExec(v *UpdateContainerGroupRequestContainerLivenessProbeExec) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.Exec = v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetFailureThreshold(v int32) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetHttpGet(v *UpdateContainerGroupRequestContainerLivenessProbeHttpGet) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.HttpGet = v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetInitialDelaySeconds(v int32) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetPeriodSeconds(v int32) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetSuccessThreshold(v int32) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetTcpSocket(v *UpdateContainerGroupRequestContainerLivenessProbeTcpSocket) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.TcpSocket = v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetTimeoutSeconds(v int32) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.TimeoutSeconds = &v
	return s
}

type UpdateContainerGroupRequestContainerLivenessProbeExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestContainerLivenessProbeExec) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLivenessProbeExec) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLivenessProbeExec) SetCommand(v []*string) *UpdateContainerGroupRequestContainerLivenessProbeExec {
	s.Command = v
	return s
}

type UpdateContainerGroupRequestContainerLivenessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s UpdateContainerGroupRequestContainerLivenessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLivenessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLivenessProbeHttpGet) SetPath(v string) *UpdateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbeHttpGet) SetPort(v int32) *UpdateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbeHttpGet) SetScheme(v string) *UpdateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Scheme = &v
	return s
}

type UpdateContainerGroupRequestContainerLivenessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s UpdateContainerGroupRequestContainerLivenessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLivenessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLivenessProbeTcpSocket) SetPort(v int32) *UpdateContainerGroupRequestContainerLivenessProbeTcpSocket {
	s.Port = &v
	return s
}

type UpdateContainerGroupRequestContainerReadinessProbe struct {
	Exec                *UpdateContainerGroupRequestContainerReadinessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" require:"true" type:"Struct"`
	FailureThreshold    *int32                                                       `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *UpdateContainerGroupRequestContainerReadinessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" require:"true" type:"Struct"`
	InitialDelaySeconds *int32                                                       `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                       `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                       `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *UpdateContainerGroupRequestContainerReadinessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" require:"true" type:"Struct"`
	TimeoutSeconds      *int32                                                       `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s UpdateContainerGroupRequestContainerReadinessProbe) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerReadinessProbe) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetExec(v *UpdateContainerGroupRequestContainerReadinessProbeExec) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.Exec = v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetFailureThreshold(v int32) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetHttpGet(v *UpdateContainerGroupRequestContainerReadinessProbeHttpGet) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.HttpGet = v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetInitialDelaySeconds(v int32) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetPeriodSeconds(v int32) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetSuccessThreshold(v int32) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetTcpSocket(v *UpdateContainerGroupRequestContainerReadinessProbeTcpSocket) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.TcpSocket = v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetTimeoutSeconds(v int32) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.TimeoutSeconds = &v
	return s
}

type UpdateContainerGroupRequestContainerReadinessProbeExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestContainerReadinessProbeExec) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerReadinessProbeExec) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerReadinessProbeExec) SetCommand(v []*string) *UpdateContainerGroupRequestContainerReadinessProbeExec {
	s.Command = v
	return s
}

type UpdateContainerGroupRequestContainerReadinessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s UpdateContainerGroupRequestContainerReadinessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerReadinessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerReadinessProbeHttpGet) SetPath(v string) *UpdateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbeHttpGet) SetPort(v int32) *UpdateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbeHttpGet) SetScheme(v string) *UpdateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Scheme = &v
	return s
}

type UpdateContainerGroupRequestContainerReadinessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s UpdateContainerGroupRequestContainerReadinessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerReadinessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerReadinessProbeTcpSocket) SetPort(v int32) *UpdateContainerGroupRequestContainerReadinessProbeTcpSocket {
	s.Port = &v
	return s
}

type UpdateContainerGroupRequestContainerSecurityContext struct {
	Capability             *UpdateContainerGroupRequestContainerSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" require:"true" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                          `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                         `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s UpdateContainerGroupRequestContainerSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerSecurityContext) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerSecurityContext) SetCapability(v *UpdateContainerGroupRequestContainerSecurityContextCapability) *UpdateContainerGroupRequestContainerSecurityContext {
	s.Capability = v
	return s
}

func (s *UpdateContainerGroupRequestContainerSecurityContext) SetReadOnlyRootFilesystem(v bool) *UpdateContainerGroupRequestContainerSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerSecurityContext) SetRunAsUser(v int64) *UpdateContainerGroupRequestContainerSecurityContext {
	s.RunAsUser = &v
	return s
}

type UpdateContainerGroupRequestContainerSecurityContextCapability struct {
	Add []*string `json:"Add,omitempty" xml:"Add,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestContainerSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerSecurityContextCapability) SetAdd(v []*string) *UpdateContainerGroupRequestContainerSecurityContextCapability {
	s.Add = v
	return s
}

type UpdateContainerGroupRequestContainerEnvironmentVar struct {
	FieldRef *UpdateContainerGroupRequestContainerEnvironmentVarFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" require:"true" type:"Struct"`
	// The name of the environment variable for container N.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable for container N.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateContainerGroupRequestContainerEnvironmentVar) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerEnvironmentVar) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerEnvironmentVar) SetFieldRef(v *UpdateContainerGroupRequestContainerEnvironmentVarFieldRef) *UpdateContainerGroupRequestContainerEnvironmentVar {
	s.FieldRef = v
	return s
}

func (s *UpdateContainerGroupRequestContainerEnvironmentVar) SetKey(v string) *UpdateContainerGroupRequestContainerEnvironmentVar {
	s.Key = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerEnvironmentVar) SetValue(v string) *UpdateContainerGroupRequestContainerEnvironmentVar {
	s.Value = &v
	return s
}

type UpdateContainerGroupRequestContainerEnvironmentVarFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s UpdateContainerGroupRequestContainerEnvironmentVarFieldRef) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerEnvironmentVarFieldRef) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerEnvironmentVarFieldRef) SetFieldPath(v string) *UpdateContainerGroupRequestContainerEnvironmentVarFieldRef {
	s.FieldPath = &v
	return s
}

type UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders struct {
	// The quest parameter of HTTP GET requests when you use HTTP requests to specify the postStart callback function.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The quest parameter value of HTTP GET requests when you use HTTP requests to specify the postStart callback function.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders) SetName(v string) *UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders {
	s.Name = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders) SetValue(v string) *UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders {
	s.Value = &v
	return s
}

type UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader struct {
	// The quest parameter of HTTP GET requests when you use HTTP requests to specify the preStop callback function.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The quest parameter value of HTTP GET requests when you use HTTP requests to specify the preStop callback function.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) SetName(v string) *UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader {
	s.Name = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) SetValue(v string) *UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader {
	s.Value = &v
	return s
}

type UpdateContainerGroupRequestContainerPort struct {
	// The port number. Valid values: 1 to 65535.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol of container N. Valid values: TCP and UDP.
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s UpdateContainerGroupRequestContainerPort) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerPort) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerPort) SetPort(v int32) *UpdateContainerGroupRequestContainerPort {
	s.Port = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerPort) SetProtocol(v string) *UpdateContainerGroupRequestContainerPort {
	s.Protocol = &v
	return s
}

type UpdateContainerGroupRequestContainerVolumeMount struct {
	// The directory of the volume mounted to container N. The data in this directory is overwritten by the data on the volume. Specify this parameter with caution.
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation settings of volume N. Mount propagation allows volumes that are mounted on one container to be shared with other containers in the same pod, or even with other pods on the same node. Valid values:
	//
	// - None: The volume mount does not receive subsequent mounts that are mounted to this volume or its subdirectories.
	// - HostToContainer: The volume mount receives all subsequent mounts that are mounted to this volume or its subdirectories.
	// - Bidirectional: This value is similar to HostToContainer. The volume mount receives all subsequent mounts that are mounted to this volume or its subdirectories. In addition, all volume mounts created by the container are propagated back to the host and to all containers of all pods that use the same volume.
	//
	//  Default value: None.
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The name of the volume mounted to container N. Valid values: the values of Volume.N.Name, which are the names of volumes mounted to the elastic container instance.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the volume is read-only. Default value: false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The subdirectory of the volume mounted to container N. The pod can mount different directories of the same volume to different subdirectories of containers.
	SubPath *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s UpdateContainerGroupRequestContainerVolumeMount) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerVolumeMount) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerVolumeMount) SetMountPath(v string) *UpdateContainerGroupRequestContainerVolumeMount {
	s.MountPath = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerVolumeMount) SetMountPropagation(v string) *UpdateContainerGroupRequestContainerVolumeMount {
	s.MountPropagation = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerVolumeMount) SetName(v string) *UpdateContainerGroupRequestContainerVolumeMount {
	s.Name = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerVolumeMount) SetReadOnly(v bool) *UpdateContainerGroupRequestContainerVolumeMount {
	s.ReadOnly = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerVolumeMount) SetSubPath(v string) *UpdateContainerGroupRequestContainerVolumeMount {
	s.SubPath = &v
	return s
}

type UpdateContainerGroupRequestImageRegistryCredential struct {
	// The password that is used to log on to the image repository.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The endpoint of the image repository. This endpoint does not contain `http://` or `https://`.
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The username that is used to log on to the image repository.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateContainerGroupRequestImageRegistryCredential) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestImageRegistryCredential) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestImageRegistryCredential) SetPassword(v string) *UpdateContainerGroupRequestImageRegistryCredential {
	s.Password = &v
	return s
}

func (s *UpdateContainerGroupRequestImageRegistryCredential) SetServer(v string) *UpdateContainerGroupRequestImageRegistryCredential {
	s.Server = &v
	return s
}

func (s *UpdateContainerGroupRequestImageRegistryCredential) SetUserName(v string) *UpdateContainerGroupRequestImageRegistryCredential {
	s.UserName = &v
	return s
}

type UpdateContainerGroupRequestInitContainer struct {
	SecurityContext *UpdateContainerGroupRequestInitContainerSecurityContext `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" require:"true" type:"Struct"`
	Arg             []*string                                                `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	Command         []*string                                                `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// The number of vCPUs that you want to allocate to init container N.
	Cpu            *float32                                                  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	EnvironmentVar []*UpdateContainerGroupRequestInitContainerEnvironmentVar `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	// The number of GPUs that you want to allocate to init container N.
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The image of init container N.
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image pull policy of init container N.
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The size of memory that you want to allocate to init container N.
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of init container N.
	Name *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	Port []*UpdateContainerGroupRequestInitContainerPort `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
	// Specifies whether init container N allocates buffer resources to standard input streams when the container runs. If you do not specify this parameter, an EOF error may occur. Default value: false.
	Stdin *bool `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	// Specifies whether to maintain standard input streams connected after a client is disconnected. If Stdin is set to true, standard input streams remain connected during multiple sessions. If InitContainer.N.StdinOnce is set to true, standard input streams are connected after the container is started and remain idle until a client is connected to receive data. After the client is disconnected, streams are also disconnected and remain in the disconnected state until the container is started again.
	StdinOnce *bool `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	// Specifies whether to enable interaction. Default value: false. If the command is a /bin/bash command, set the value to true.
	Tty         *bool                                                  `json:"Tty,omitempty" xml:"Tty,omitempty"`
	VolumeMount []*UpdateContainerGroupRequestInitContainerVolumeMount `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
	// The working directory of init container N.
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainer) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainer) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainer) SetSecurityContext(v *UpdateContainerGroupRequestInitContainerSecurityContext) *UpdateContainerGroupRequestInitContainer {
	s.SecurityContext = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetArg(v []*string) *UpdateContainerGroupRequestInitContainer {
	s.Arg = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetCommand(v []*string) *UpdateContainerGroupRequestInitContainer {
	s.Command = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetCpu(v float32) *UpdateContainerGroupRequestInitContainer {
	s.Cpu = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetEnvironmentVar(v []*UpdateContainerGroupRequestInitContainerEnvironmentVar) *UpdateContainerGroupRequestInitContainer {
	s.EnvironmentVar = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetGpu(v int32) *UpdateContainerGroupRequestInitContainer {
	s.Gpu = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetImage(v string) *UpdateContainerGroupRequestInitContainer {
	s.Image = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetImagePullPolicy(v string) *UpdateContainerGroupRequestInitContainer {
	s.ImagePullPolicy = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetMemory(v float32) *UpdateContainerGroupRequestInitContainer {
	s.Memory = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetName(v string) *UpdateContainerGroupRequestInitContainer {
	s.Name = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetPort(v []*UpdateContainerGroupRequestInitContainerPort) *UpdateContainerGroupRequestInitContainer {
	s.Port = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetStdin(v bool) *UpdateContainerGroupRequestInitContainer {
	s.Stdin = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetStdinOnce(v bool) *UpdateContainerGroupRequestInitContainer {
	s.StdinOnce = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetTty(v bool) *UpdateContainerGroupRequestInitContainer {
	s.Tty = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetVolumeMount(v []*UpdateContainerGroupRequestInitContainerVolumeMount) *UpdateContainerGroupRequestInitContainer {
	s.VolumeMount = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetWorkingDir(v string) *UpdateContainerGroupRequestInitContainer {
	s.WorkingDir = &v
	return s
}

type UpdateContainerGroupRequestInitContainerSecurityContext struct {
	Capability             *UpdateContainerGroupRequestInitContainerSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" require:"true" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                              `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                             `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainerSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerSecurityContext) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerSecurityContext) SetCapability(v *UpdateContainerGroupRequestInitContainerSecurityContextCapability) *UpdateContainerGroupRequestInitContainerSecurityContext {
	s.Capability = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerSecurityContext) SetReadOnlyRootFilesystem(v bool) *UpdateContainerGroupRequestInitContainerSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerSecurityContext) SetRunAsUser(v int64) *UpdateContainerGroupRequestInitContainerSecurityContext {
	s.RunAsUser = &v
	return s
}

type UpdateContainerGroupRequestInitContainerSecurityContextCapability struct {
	Add []*string `json:"Add,omitempty" xml:"Add,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestInitContainerSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerSecurityContextCapability) SetAdd(v []*string) *UpdateContainerGroupRequestInitContainerSecurityContextCapability {
	s.Add = v
	return s
}

type UpdateContainerGroupRequestInitContainerEnvironmentVar struct {
	FieldRef *UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" require:"true" type:"Struct"`
	// The name of the environment variable for init container N.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable for init container N.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainerEnvironmentVar) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerEnvironmentVar) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerEnvironmentVar) SetFieldRef(v *UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef) *UpdateContainerGroupRequestInitContainerEnvironmentVar {
	s.FieldRef = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerEnvironmentVar) SetKey(v string) *UpdateContainerGroupRequestInitContainerEnvironmentVar {
	s.Key = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerEnvironmentVar) SetValue(v string) *UpdateContainerGroupRequestInitContainerEnvironmentVar {
	s.Value = &v
	return s
}

type UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef) SetFieldPath(v string) *UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef {
	s.FieldPath = &v
	return s
}

type UpdateContainerGroupRequestInitContainerPort struct {
	// The port number of init container N. Valid values: 1 to 65535.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol of init container N. Valid values: TCP and UDP.
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainerPort) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerPort) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerPort) SetPort(v int32) *UpdateContainerGroupRequestInitContainerPort {
	s.Port = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerPort) SetProtocol(v string) *UpdateContainerGroupRequestInitContainerPort {
	s.Protocol = &v
	return s
}

type UpdateContainerGroupRequestInitContainerVolumeMount struct {
	// The directory of the volume mounted to init container N. The data in this directory is overwritten by the data on the volume. Specify this parameter with caution.
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation setting of the volume. Mount propagation allows volumes that are mounted on one init container to be shared with other init containers in the same pod, or even with other pods on the same node. Valid values:
	//
	// - None: The volume mount does not receive subsequent mounts that are mounted to this volume or its subdirectories.
	// - HostToContainer: The volume mount receives all subsequent mounts that are mounted to this volume or its subdirectories.
	// - Bidirectional: This value is similar to HostToContainer. The volume mount receives all subsequent mounts that are mounted to this volume or its subdirectories. In addition, all volume mounts created by the container are propagated back to the host and to all containers of all pods that use the same volume.
	//
	//  Default value: None.
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The name of the volume mounted to init container N. Valid values: the values of Volume.N.Name, which are the names of volumes mounted to the elastic container instance.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the volume is read-only. Default value: false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The subdirectory of the volume mounted to init container N. The pod can mount different directories of the same volume to different subdirectories of containers.
	SubPath *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainerVolumeMount) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerVolumeMount) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerVolumeMount) SetMountPath(v string) *UpdateContainerGroupRequestInitContainerVolumeMount {
	s.MountPath = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerVolumeMount) SetMountPropagation(v string) *UpdateContainerGroupRequestInitContainerVolumeMount {
	s.MountPropagation = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerVolumeMount) SetName(v string) *UpdateContainerGroupRequestInitContainerVolumeMount {
	s.Name = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerVolumeMount) SetReadOnly(v bool) *UpdateContainerGroupRequestInitContainerVolumeMount {
	s.ReadOnly = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerVolumeMount) SetSubPath(v string) *UpdateContainerGroupRequestInitContainerVolumeMount {
	s.SubPath = &v
	return s
}

type UpdateContainerGroupRequestTag struct {
	// The key of tag N.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateContainerGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestTag) SetKey(v string) *UpdateContainerGroupRequestTag {
	s.Key = &v
	return s
}

func (s *UpdateContainerGroupRequestTag) SetValue(v string) *UpdateContainerGroupRequestTag {
	s.Value = &v
	return s
}

type UpdateContainerGroupRequestVolume struct {
	ConfigFileVolume *UpdateContainerGroupRequestVolumeConfigFileVolume `json:"ConfigFileVolume,omitempty" xml:"ConfigFileVolume,omitempty" require:"true" type:"Struct"`
	EmptyDirVolume   *UpdateContainerGroupRequestVolumeEmptyDirVolume   `json:"EmptyDirVolume,omitempty" xml:"EmptyDirVolume,omitempty" require:"true" type:"Struct"`
	FlexVolume       *UpdateContainerGroupRequestVolumeFlexVolume       `json:"FlexVolume,omitempty" xml:"FlexVolume,omitempty" require:"true" type:"Struct"`
	HostPathVolume   *UpdateContainerGroupRequestVolumeHostPathVolume   `json:"HostPathVolume,omitempty" xml:"HostPathVolume,omitempty" require:"true" type:"Struct"`
	NFSVolume        *UpdateContainerGroupRequestVolumeNFSVolume        `json:"NFSVolume,omitempty" xml:"NFSVolume,omitempty" require:"true" type:"Struct"`
	// The name of volume N.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of volume N. Valid values:
	//
	// *   EmptyDirVolume: an emptyDir volume, which indicates a temporary directory.
	// *   ConfigFileVolume: a ConfigFile volume, which indicates a configuration file.
	// *   NFSVolume: an NFS volume, which indicates a network file system, such as a NAS file system.
	// *   FlexVolume: a volume that is mounted by using the FlexVolume plug-in. These volumes include disks, NAS file systems, and OSS buckets.
	// *   HostPathVolume: a HostPath volume, which indicates a file or directory on the host. This value is unavailable.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateContainerGroupRequestVolume) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolume) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolume) SetConfigFileVolume(v *UpdateContainerGroupRequestVolumeConfigFileVolume) *UpdateContainerGroupRequestVolume {
	s.ConfigFileVolume = v
	return s
}

func (s *UpdateContainerGroupRequestVolume) SetEmptyDirVolume(v *UpdateContainerGroupRequestVolumeEmptyDirVolume) *UpdateContainerGroupRequestVolume {
	s.EmptyDirVolume = v
	return s
}

func (s *UpdateContainerGroupRequestVolume) SetFlexVolume(v *UpdateContainerGroupRequestVolumeFlexVolume) *UpdateContainerGroupRequestVolume {
	s.FlexVolume = v
	return s
}

func (s *UpdateContainerGroupRequestVolume) SetHostPathVolume(v *UpdateContainerGroupRequestVolumeHostPathVolume) *UpdateContainerGroupRequestVolume {
	s.HostPathVolume = v
	return s
}

func (s *UpdateContainerGroupRequestVolume) SetNFSVolume(v *UpdateContainerGroupRequestVolumeNFSVolume) *UpdateContainerGroupRequestVolume {
	s.NFSVolume = v
	return s
}

func (s *UpdateContainerGroupRequestVolume) SetName(v string) *UpdateContainerGroupRequestVolume {
	s.Name = &v
	return s
}

func (s *UpdateContainerGroupRequestVolume) SetType(v string) *UpdateContainerGroupRequestVolume {
	s.Type = &v
	return s
}

type UpdateContainerGroupRequestVolumeConfigFileVolume struct {
	ConfigFileToPath []*UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath `json:"ConfigFileToPath,omitempty" xml:"ConfigFileToPath,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestVolumeConfigFileVolume) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolumeConfigFileVolume) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolumeConfigFileVolume) SetConfigFileToPath(v []*UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) *UpdateContainerGroupRequestVolumeConfigFileVolume {
	s.ConfigFileToPath = v
	return s
}

type UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Path    *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) SetContent(v string) *UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath {
	s.Content = &v
	return s
}

func (s *UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) SetPath(v string) *UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath {
	s.Path = &v
	return s
}

type UpdateContainerGroupRequestVolumeEmptyDirVolume struct {
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	SizeLimit *string `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s UpdateContainerGroupRequestVolumeEmptyDirVolume) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolumeEmptyDirVolume) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolumeEmptyDirVolume) SetMedium(v string) *UpdateContainerGroupRequestVolumeEmptyDirVolume {
	s.Medium = &v
	return s
}

func (s *UpdateContainerGroupRequestVolumeEmptyDirVolume) SetSizeLimit(v string) *UpdateContainerGroupRequestVolumeEmptyDirVolume {
	s.SizeLimit = &v
	return s
}

type UpdateContainerGroupRequestVolumeFlexVolume struct {
	Driver  *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	FsType  *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s UpdateContainerGroupRequestVolumeFlexVolume) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolumeFlexVolume) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolumeFlexVolume) SetDriver(v string) *UpdateContainerGroupRequestVolumeFlexVolume {
	s.Driver = &v
	return s
}

func (s *UpdateContainerGroupRequestVolumeFlexVolume) SetFsType(v string) *UpdateContainerGroupRequestVolumeFlexVolume {
	s.FsType = &v
	return s
}

func (s *UpdateContainerGroupRequestVolumeFlexVolume) SetOptions(v string) *UpdateContainerGroupRequestVolumeFlexVolume {
	s.Options = &v
	return s
}

type UpdateContainerGroupRequestVolumeHostPathVolume struct {
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateContainerGroupRequestVolumeHostPathVolume) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolumeHostPathVolume) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolumeHostPathVolume) SetPath(v string) *UpdateContainerGroupRequestVolumeHostPathVolume {
	s.Path = &v
	return s
}

func (s *UpdateContainerGroupRequestVolumeHostPathVolume) SetType(v string) *UpdateContainerGroupRequestVolumeHostPathVolume {
	s.Type = &v
	return s
}

type UpdateContainerGroupRequestVolumeNFSVolume struct {
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ReadOnly *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
}

func (s UpdateContainerGroupRequestVolumeNFSVolume) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolumeNFSVolume) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolumeNFSVolume) SetPath(v string) *UpdateContainerGroupRequestVolumeNFSVolume {
	s.Path = &v
	return s
}

func (s *UpdateContainerGroupRequestVolumeNFSVolume) SetReadOnly(v bool) *UpdateContainerGroupRequestVolumeNFSVolume {
	s.ReadOnly = &v
	return s
}

func (s *UpdateContainerGroupRequestVolumeNFSVolume) SetServer(v string) *UpdateContainerGroupRequestVolumeNFSVolume {
	s.Server = &v
	return s
}

type UpdateContainerGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateContainerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupResponseBody) SetRequestId(v string) *UpdateContainerGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateContainerGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateContainerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateContainerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupResponse) SetHeaders(v map[string]*string) *UpdateContainerGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateContainerGroupResponse) SetStatusCode(v int32) *UpdateContainerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContainerGroupResponse) SetBody(v *UpdateContainerGroupResponseBody) *UpdateContainerGroupResponse {
	s.Body = v
	return s
}

type UpdateImageCacheRequest struct {
	AcrRegistryInfo []*UpdateImageCacheRequestAcrRegistryInfo `json:"AcrRegistryInfo,omitempty" xml:"AcrRegistryInfo,omitempty" type:"Repeated"`
	// Specifies whether to enable reuse of image cache layers. If you enable this feature, and the image cache that you want to create and an existing image cache contain duplicate image layers, the system reuses the duplicate image layers to create the new image cache. This accelerates the creation of the image cache. Valid values:
	//
	// *   true: enables reuse of image cache layers.
	// *   false: does not enable reuse of image cache layers.
	//
	// Default value: false.
	AutoMatchImageCache *bool `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the elastic IP address (EIP). If you want to pull images over the Internet, make sure that the elastic container instance can access the Internet. You can configure an EIP or a NAT gateway for the elastic container instance to access the Internet.
	EipInstanceId       *string `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	EliminationStrategy *string `json:"EliminationStrategy,omitempty" xml:"EliminationStrategy,omitempty"`
	// Specifies whether to enable instant image cache. After you enable the feature, image caches can be created in an accelerated manner. Valid values:
	//
	// *   true: enables instant image cache.
	// *   false: does not enable instant image cache.
	//
	// Default value: false.
	//
	// >  The system automatically generates a temporary local snapshot for the image cache during the use of the instant image cache feature. You are charged for the instant use of the snapshot.
	Flash *bool `json:"Flash,omitempty" xml:"Flash,omitempty"`
	// The number of temporary local snapshots. By default, the system creates one snapshot for each image cache. If an image cache is used to create multiple elastic container instances at a time, we recommend that you set this parameter to create multiple snapshots for the image cache. We recommend that you create one snapshot for creation of every 1,000 elastic container instances.
	//
	// >  If you set the Flash parameter to true, instant image cache is enabled. During the creation of the image cache, the system first creates a temporary local snapshot for you to instantly use the snapshot. After the temporary local snapshot is created, the system begins to create a regular snapshot. After the regular snapshot is created, the temporary local snapshot is automatically deleted.
	FlashCopyCount *int32    `json:"FlashCopyCount,omitempty" xml:"FlashCopyCount,omitempty"`
	Image          []*string `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
	// The ID of the image cache.
	ImageCacheId *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	// The name of the image cache.
	ImageCacheName *string `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	// The size of the image cache. Unit: GiB. Default value: 20.
	ImageCacheSize          *int32                                            `json:"ImageCacheSize,omitempty" xml:"ImageCacheSize,omitempty"`
	ImageRegistryCredential []*UpdateImageCacheRequestImageRegistryCredential `json:"ImageRegistryCredential,omitempty" xml:"ImageRegistryCredential,omitempty" type:"Repeated"`
	OwnerAccount            *string                                           `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64                                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the image cache.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The retention period of the image cache. Unit: days. When the retention period ends, the image cache expires and is deleted. By default, image caches do not expire.
	//
	// >  The image caches that fail to be created are only retained for one day.
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The ID of the security group.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The number of regular snapshots. By default, the system creates one snapshot for each image cache. If an image cache is used to create multiple elastic container instances at a time, we recommend that you set this parameter to create multiple snapshots for the image cache. We recommend that you create one snapshot for creation of every 1,000 elastic container instances.
	//
	// >  If you set the Flash parameter to false, instant image cache is disabled. In this case, only regular snapshots are generated during the creation of the image cache.
	StandardCopyCount *int32                        `json:"StandardCopyCount,omitempty" xml:"StandardCopyCount,omitempty"`
	Tag               []*UpdateImageCacheRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s UpdateImageCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageCacheRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageCacheRequest) SetAcrRegistryInfo(v []*UpdateImageCacheRequestAcrRegistryInfo) *UpdateImageCacheRequest {
	s.AcrRegistryInfo = v
	return s
}

func (s *UpdateImageCacheRequest) SetAutoMatchImageCache(v bool) *UpdateImageCacheRequest {
	s.AutoMatchImageCache = &v
	return s
}

func (s *UpdateImageCacheRequest) SetClientToken(v string) *UpdateImageCacheRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateImageCacheRequest) SetEipInstanceId(v string) *UpdateImageCacheRequest {
	s.EipInstanceId = &v
	return s
}

func (s *UpdateImageCacheRequest) SetEliminationStrategy(v string) *UpdateImageCacheRequest {
	s.EliminationStrategy = &v
	return s
}

func (s *UpdateImageCacheRequest) SetFlash(v bool) *UpdateImageCacheRequest {
	s.Flash = &v
	return s
}

func (s *UpdateImageCacheRequest) SetFlashCopyCount(v int32) *UpdateImageCacheRequest {
	s.FlashCopyCount = &v
	return s
}

func (s *UpdateImageCacheRequest) SetImage(v []*string) *UpdateImageCacheRequest {
	s.Image = v
	return s
}

func (s *UpdateImageCacheRequest) SetImageCacheId(v string) *UpdateImageCacheRequest {
	s.ImageCacheId = &v
	return s
}

func (s *UpdateImageCacheRequest) SetImageCacheName(v string) *UpdateImageCacheRequest {
	s.ImageCacheName = &v
	return s
}

func (s *UpdateImageCacheRequest) SetImageCacheSize(v int32) *UpdateImageCacheRequest {
	s.ImageCacheSize = &v
	return s
}

func (s *UpdateImageCacheRequest) SetImageRegistryCredential(v []*UpdateImageCacheRequestImageRegistryCredential) *UpdateImageCacheRequest {
	s.ImageRegistryCredential = v
	return s
}

func (s *UpdateImageCacheRequest) SetOwnerAccount(v string) *UpdateImageCacheRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateImageCacheRequest) SetOwnerId(v int64) *UpdateImageCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateImageCacheRequest) SetRegionId(v string) *UpdateImageCacheRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateImageCacheRequest) SetResourceGroupId(v string) *UpdateImageCacheRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateImageCacheRequest) SetResourceOwnerAccount(v string) *UpdateImageCacheRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateImageCacheRequest) SetResourceOwnerId(v int64) *UpdateImageCacheRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateImageCacheRequest) SetRetentionDays(v int32) *UpdateImageCacheRequest {
	s.RetentionDays = &v
	return s
}

func (s *UpdateImageCacheRequest) SetSecurityGroupId(v string) *UpdateImageCacheRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateImageCacheRequest) SetStandardCopyCount(v int32) *UpdateImageCacheRequest {
	s.StandardCopyCount = &v
	return s
}

func (s *UpdateImageCacheRequest) SetTag(v []*UpdateImageCacheRequestTag) *UpdateImageCacheRequest {
	s.Tag = v
	return s
}

func (s *UpdateImageCacheRequest) SetVSwitchId(v string) *UpdateImageCacheRequest {
	s.VSwitchId = &v
	return s
}

type UpdateImageCacheRequestAcrRegistryInfo struct {
	Domain []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
	// The ID of the Container Registry Enterprise Edition instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Container Registry Enterprise Edition instance.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID of Container Registry Enterprise Edition instance N.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateImageCacheRequestAcrRegistryInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageCacheRequestAcrRegistryInfo) GoString() string {
	return s.String()
}

func (s *UpdateImageCacheRequestAcrRegistryInfo) SetDomain(v []*string) *UpdateImageCacheRequestAcrRegistryInfo {
	s.Domain = v
	return s
}

func (s *UpdateImageCacheRequestAcrRegistryInfo) SetInstanceId(v string) *UpdateImageCacheRequestAcrRegistryInfo {
	s.InstanceId = &v
	return s
}

func (s *UpdateImageCacheRequestAcrRegistryInfo) SetInstanceName(v string) *UpdateImageCacheRequestAcrRegistryInfo {
	s.InstanceName = &v
	return s
}

func (s *UpdateImageCacheRequestAcrRegistryInfo) SetRegionId(v string) *UpdateImageCacheRequestAcrRegistryInfo {
	s.RegionId = &v
	return s
}

type UpdateImageCacheRequestImageRegistryCredential struct {
	// The password that is used to log on to image repository N.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The endpoint of the image repository.
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The username that is used to log on to image repository N.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateImageCacheRequestImageRegistryCredential) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageCacheRequestImageRegistryCredential) GoString() string {
	return s.String()
}

func (s *UpdateImageCacheRequestImageRegistryCredential) SetPassword(v string) *UpdateImageCacheRequestImageRegistryCredential {
	s.Password = &v
	return s
}

func (s *UpdateImageCacheRequestImageRegistryCredential) SetServer(v string) *UpdateImageCacheRequestImageRegistryCredential {
	s.Server = &v
	return s
}

func (s *UpdateImageCacheRequestImageRegistryCredential) SetUserName(v string) *UpdateImageCacheRequestImageRegistryCredential {
	s.UserName = &v
	return s
}

type UpdateImageCacheRequestTag struct {
	// The key of tag N of the image cache. Valid values of N: 1 to 20.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the image cache. Valid values of N: 1 to 20.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateImageCacheRequestTag) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageCacheRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateImageCacheRequestTag) SetKey(v string) *UpdateImageCacheRequestTag {
	s.Key = &v
	return s
}

func (s *UpdateImageCacheRequestTag) SetValue(v string) *UpdateImageCacheRequestTag {
	s.Value = &v
	return s
}

type UpdateImageCacheResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateImageCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageCacheResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageCacheResponseBody) SetRequestId(v string) *UpdateImageCacheResponseBody {
	s.RequestId = &v
	return s
}

type UpdateImageCacheResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateImageCacheResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateImageCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageCacheResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageCacheResponse) SetHeaders(v map[string]*string) *UpdateImageCacheResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageCacheResponse) SetStatusCode(v int32) *UpdateImageCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImageCacheResponse) SetBody(v *UpdateImageCacheResponseBody) *UpdateImageCacheResponse {
	s.Body = v
	return s
}

type UpdateVirtualNodeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate this value, but you must ensure that it is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency](~~25693~~).
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterDNS      *string `json:"ClusterDNS,omitempty" xml:"ClusterDNS,omitempty"`
	ClusterDomain   *string `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	CustomResources *string `json:"CustomResources,omitempty" xml:"CustomResources,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the virtual node.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the virtual node belongs.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the security group.
	SecurityGroupId *string                        `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag             []*UpdateVirtualNodeRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual node.
	VirtualNodeId *string `json:"VirtualNodeId,omitempty" xml:"VirtualNodeId,omitempty"`
	// The name of the virtual node.
	VirtualNodeName *string `json:"VirtualNodeName,omitempty" xml:"VirtualNodeName,omitempty"`
}

func (s UpdateVirtualNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVirtualNodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateVirtualNodeRequest) SetClientToken(v string) *UpdateVirtualNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetClusterDNS(v string) *UpdateVirtualNodeRequest {
	s.ClusterDNS = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetClusterDomain(v string) *UpdateVirtualNodeRequest {
	s.ClusterDomain = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetCustomResources(v string) *UpdateVirtualNodeRequest {
	s.CustomResources = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetOwnerAccount(v string) *UpdateVirtualNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetOwnerId(v int64) *UpdateVirtualNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetRegionId(v string) *UpdateVirtualNodeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetResourceGroupId(v string) *UpdateVirtualNodeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetResourceOwnerAccount(v string) *UpdateVirtualNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetResourceOwnerId(v int64) *UpdateVirtualNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetSecurityGroupId(v string) *UpdateVirtualNodeRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetTag(v []*UpdateVirtualNodeRequestTag) *UpdateVirtualNodeRequest {
	s.Tag = v
	return s
}

func (s *UpdateVirtualNodeRequest) SetVSwitchId(v string) *UpdateVirtualNodeRequest {
	s.VSwitchId = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetVirtualNodeId(v string) *UpdateVirtualNodeRequest {
	s.VirtualNodeId = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetVirtualNodeName(v string) *UpdateVirtualNodeRequest {
	s.VirtualNodeName = &v
	return s
}

type UpdateVirtualNodeRequestTag struct {
	// The key of Tag N.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of Tag N.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateVirtualNodeRequestTag) String() string {
	return tea.Prettify(s)
}

func (s UpdateVirtualNodeRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateVirtualNodeRequestTag) SetKey(v string) *UpdateVirtualNodeRequestTag {
	s.Key = &v
	return s
}

func (s *UpdateVirtualNodeRequestTag) SetValue(v string) *UpdateVirtualNodeRequestTag {
	s.Value = &v
	return s
}

type UpdateVirtualNodeResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVirtualNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVirtualNodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVirtualNodeResponseBody) SetRequestId(v string) *UpdateVirtualNodeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateVirtualNodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateVirtualNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVirtualNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVirtualNodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateVirtualNodeResponse) SetHeaders(v map[string]*string) *UpdateVirtualNodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateVirtualNodeResponse) SetStatusCode(v int32) *UpdateVirtualNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVirtualNodeResponse) SetBody(v *UpdateVirtualNodeResponseBody) *UpdateVirtualNodeResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("eci"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
 * You must enter the Alibaba Cloud Resource Name (ARN) to obtain the permissions to push images.
 *
 * @param request CommitContainerRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CommitContainerResponse
 */
func (client *Client) CommitContainerWithOptions(request *CommitContainerRequest, runtime *util.RuntimeOptions) (_result *CommitContainerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcrRegistryInfo)) {
		query["AcrRegistryInfo"] = request.AcrRegistryInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Arn)) {
		query["Arn"] = request.Arn
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerName)) {
		query["ContainerName"] = request.ContainerName
	}

	if !tea.BoolValue(util.IsUnset(request.Image)) {
		query["Image"] = request.Image
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CommitContainer"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CommitContainerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You must enter the Alibaba Cloud Resource Name (ARN) to obtain the permissions to push images.
 *
 * @param request CommitContainerRequest
 * @return CommitContainerResponse
 */
func (client *Client) CommitContainer(request *CommitContainerRequest) (_result *CommitContainerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CommitContainerResponse{}
	_body, _err := client.CommitContainerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Usage description
 * When you call the CreateContainerGroup operation to create an elastic container instance, the system creates a service-linked role named AliyunServiceRoleForECI. This role is used to access other Alibaba Cloud services such as Elastic Compute Service (ECS) and Virtual Private Cloud (VPC). For more information, see [Elastic Container Instance service-linked role](~~212914~~).
 * ## Parameters configured for features
 * When you create an elastic container instance, you can configure features such as instances, images, and storage based on your requirements. For information about parameters configured for the features and the description of the parameters, see the following documents: **Instances** You can use one of the following methods to create an elastic container instance:
 * *   [Specify the number of vCPUs and memory size to create an elastic container instance](~~114662~~)
 *     *   [Create job-optimized elastic container instances](~~324246~~)
 *     *   [Ignore special containers during resource adjustment](~~446853~~)
 * *   [Specify ECS instance types to create an elastic container instance](~~114664~~)
 *     *   [Specify GPU-accelerated ECS instance types](~~114581~~)
 *     *   [Specify ECS instance types with local disks](~~164011~~)
 *     *   [Specify AMD-based ECS instance types](~~187411~~)
 * Both the preceding creation methods support the following features:
 * *   [Specify custom CPU options](~~197781~~)
 * *   [Create a preemptible elastic container instance](~~157759~~)
 * *   [Configure multiple zones](~~157290~~)
 * *   [Configure multiple specifications](~~146468~~)
 * *   [Use tags to manage elastic container instances](~~146608~~)
 * **Images**
 * *   [Configure a container image](~~461311~~)
 * *   [Use the image cache feature to accelerate the creation of an elastic container instance](~~141281~~)
 * *   [Specify a Container Registry Enterprise Edition instance](~~194250.~~)
 * *   [Use self-managed image repositories](~~378059~~)
 * **Networking**
 * *   [Create and Associate an EIP](~~99146~~)
 * *   [Assign a security group](~~176237~~)
 * *   [Assign an IPv6 address to an elastic container instance](~~451282~~)
 * *   [Configure maximum bandwidth](~~190635~~)
 * **Storage**
 * *   [Mount a disk volume](~~144571~~)
 * *   [Mount a NAS volume](~~464075~~)
 * *   [Mount an OSS bucket to an elastic container instance as a volume](~~464076~~)
 * *   [Mount an emptyDir volume](~~464078~~)
 * *   [Mount a ConfigFile volume](~~464080~~)
 * *   [Increase the storage capacity of the temporary storage space](~~204066~~)
 * **Container configuration**
 * *   [Configure startup commands and arguments for a container](~~94593~~)
 * *   [Use probes to perform health checks on a container](~~99053~~)
 * *   [Obtain metadata by using environment variables](~~141788~~)
 * *   [Configure a security context](~~462313~~)
 * *   [Configure the NTP service](~~462768~~)
 * **Logging and O\\&M**
 * *   [Use environment variables to configure log collection](~~121973~~)
 * *   [Save core files to volumes](~~167801~~)
 *
 * @param request CreateContainerGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateContainerGroupResponse
 */
func (client *Client) CreateContainerGroupWithOptions(request *CreateContainerGroupRequest, runtime *util.RuntimeOptions) (_result *CreateContainerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcrRegistryInfo)) {
		query["AcrRegistryInfo"] = request.AcrRegistryInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ActiveDeadlineSeconds)) {
		query["ActiveDeadlineSeconds"] = request.ActiveDeadlineSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.AutoCreateEip)) {
		query["AutoCreateEip"] = request.AutoCreateEip
	}

	if !tea.BoolValue(util.IsUnset(request.AutoMatchImageCache)) {
		query["AutoMatchImageCache"] = request.AutoMatchImageCache
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Container)) {
		query["Container"] = request.Container
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupName)) {
		query["ContainerGroupName"] = request.ContainerGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerResourceView)) {
		query["ContainerResourceView"] = request.ContainerResourceView
	}

	if !tea.BoolValue(util.IsUnset(request.CorePattern)) {
		query["CorePattern"] = request.CorePattern
	}

	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		query["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.CpuOptionsCore)) {
		query["CpuOptionsCore"] = request.CpuOptionsCore
	}

	if !tea.BoolValue(util.IsUnset(request.CpuOptionsNuma)) {
		query["CpuOptionsNuma"] = request.CpuOptionsNuma
	}

	if !tea.BoolValue(util.IsUnset(request.CpuOptionsThreadsPerCore)) {
		query["CpuOptionsThreadsPerCore"] = request.CpuOptionsThreadsPerCore
	}

	if !tea.BoolValue(util.IsUnset(request.DnsPolicy)) {
		query["DnsPolicy"] = request.DnsPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.EgressBandwidth)) {
		query["EgressBandwidth"] = request.EgressBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.EipBandwidth)) {
		query["EipBandwidth"] = request.EipBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.EipCommonBandwidthPackage)) {
		query["EipCommonBandwidthPackage"] = request.EipCommonBandwidthPackage
	}

	if !tea.BoolValue(util.IsUnset(request.EipISP)) {
		query["EipISP"] = request.EipISP
	}

	if !tea.BoolValue(util.IsUnset(request.EipInstanceId)) {
		query["EipInstanceId"] = request.EipInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EphemeralStorage)) {
		query["EphemeralStorage"] = request.EphemeralStorage
	}

	if !tea.BoolValue(util.IsUnset(request.HostAliase)) {
		query["HostAliase"] = request.HostAliase
	}

	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		query["HostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageAccelerateMode)) {
		query["ImageAccelerateMode"] = request.ImageAccelerateMode
	}

	if !tea.BoolValue(util.IsUnset(request.ImageRegistryCredential)) {
		query["ImageRegistryCredential"] = request.ImageRegistryCredential
	}

	if !tea.BoolValue(util.IsUnset(request.ImageSnapshotId)) {
		query["ImageSnapshotId"] = request.ImageSnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.IngressBandwidth)) {
		query["IngressBandwidth"] = request.IngressBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.InitContainer)) {
		query["InitContainer"] = request.InitContainer
	}

	if !tea.BoolValue(util.IsUnset(request.InsecureRegistry)) {
		query["InsecureRegistry"] = request.InsecureRegistry
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.Ipv6AddressCount)) {
		query["Ipv6AddressCount"] = request.Ipv6AddressCount
	}

	if !tea.BoolValue(util.IsUnset(request.Ipv6GatewayBandwidth)) {
		query["Ipv6GatewayBandwidth"] = request.Ipv6GatewayBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.Ipv6GatewayBandwidthEnable)) {
		query["Ipv6GatewayBandwidthEnable"] = request.Ipv6GatewayBandwidthEnable
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		query["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.NtpServer)) {
		query["NtpServer"] = request.NtpServer
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlainHttpRegistry)) {
		query["PlainHttpRegistry"] = request.PlainHttpRegistry
	}

	if !tea.BoolValue(util.IsUnset(request.RamRoleName)) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RestartPolicy)) {
		query["RestartPolicy"] = request.RestartPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleStrategy)) {
		query["ScheduleStrategy"] = request.ScheduleStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareProcessNamespace)) {
		query["ShareProcessNamespace"] = request.ShareProcessNamespace
	}

	if !tea.BoolValue(util.IsUnset(request.SpotDuration)) {
		query["SpotDuration"] = request.SpotDuration
	}

	if !tea.BoolValue(util.IsUnset(request.SpotPriceLimit)) {
		query["SpotPriceLimit"] = request.SpotPriceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.SpotStrategy)) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.StrictSpot)) {
		query["StrictSpot"] = request.StrictSpot
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TerminationGracePeriodSeconds)) {
		query["TerminationGracePeriodSeconds"] = request.TerminationGracePeriodSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		query["Volume"] = request.Volume
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.DnsConfig)) {
		query["DnsConfig"] = request.DnsConfig
	}

	if !tea.BoolValue(util.IsUnset(request.HostSecurityContext)) {
		query["HostSecurityContext"] = request.HostSecurityContext
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityContext)) {
		query["SecurityContext"] = request.SecurityContext
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateContainerGroup"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateContainerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Usage description
 * When you call the CreateContainerGroup operation to create an elastic container instance, the system creates a service-linked role named AliyunServiceRoleForECI. This role is used to access other Alibaba Cloud services such as Elastic Compute Service (ECS) and Virtual Private Cloud (VPC). For more information, see [Elastic Container Instance service-linked role](~~212914~~).
 * ## Parameters configured for features
 * When you create an elastic container instance, you can configure features such as instances, images, and storage based on your requirements. For information about parameters configured for the features and the description of the parameters, see the following documents: **Instances** You can use one of the following methods to create an elastic container instance:
 * *   [Specify the number of vCPUs and memory size to create an elastic container instance](~~114662~~)
 *     *   [Create job-optimized elastic container instances](~~324246~~)
 *     *   [Ignore special containers during resource adjustment](~~446853~~)
 * *   [Specify ECS instance types to create an elastic container instance](~~114664~~)
 *     *   [Specify GPU-accelerated ECS instance types](~~114581~~)
 *     *   [Specify ECS instance types with local disks](~~164011~~)
 *     *   [Specify AMD-based ECS instance types](~~187411~~)
 * Both the preceding creation methods support the following features:
 * *   [Specify custom CPU options](~~197781~~)
 * *   [Create a preemptible elastic container instance](~~157759~~)
 * *   [Configure multiple zones](~~157290~~)
 * *   [Configure multiple specifications](~~146468~~)
 * *   [Use tags to manage elastic container instances](~~146608~~)
 * **Images**
 * *   [Configure a container image](~~461311~~)
 * *   [Use the image cache feature to accelerate the creation of an elastic container instance](~~141281~~)
 * *   [Specify a Container Registry Enterprise Edition instance](~~194250.~~)
 * *   [Use self-managed image repositories](~~378059~~)
 * **Networking**
 * *   [Create and Associate an EIP](~~99146~~)
 * *   [Assign a security group](~~176237~~)
 * *   [Assign an IPv6 address to an elastic container instance](~~451282~~)
 * *   [Configure maximum bandwidth](~~190635~~)
 * **Storage**
 * *   [Mount a disk volume](~~144571~~)
 * *   [Mount a NAS volume](~~464075~~)
 * *   [Mount an OSS bucket to an elastic container instance as a volume](~~464076~~)
 * *   [Mount an emptyDir volume](~~464078~~)
 * *   [Mount a ConfigFile volume](~~464080~~)
 * *   [Increase the storage capacity of the temporary storage space](~~204066~~)
 * **Container configuration**
 * *   [Configure startup commands and arguments for a container](~~94593~~)
 * *   [Use probes to perform health checks on a container](~~99053~~)
 * *   [Obtain metadata by using environment variables](~~141788~~)
 * *   [Configure a security context](~~462313~~)
 * *   [Configure the NTP service](~~462768~~)
 * **Logging and O\\&M**
 * *   [Use environment variables to configure log collection](~~121973~~)
 * *   [Save core files to volumes](~~167801~~)
 *
 * @param request CreateContainerGroupRequest
 * @return CreateContainerGroupResponse
 */
func (client *Client) CreateContainerGroup(request *CreateContainerGroupRequest) (_result *CreateContainerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateContainerGroupResponse{}
	_body, _err := client.CreateContainerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Precautions
 * - You are charged for creation of image caches. We recommend that you learn the relevant billing information in advance. For more information, see [Billing](~~89142~~).
 * - Before you create an image cache, you must estimate the total size of the images that you want to cache. If the total size of the images exceeds the specified cache size, the image cache cannot be created.
 * - When an image cache is being created, the system creates an intermediate elastic container instance and an intermediate enhanced SSD (ESSD) at performance level 1 (PL1). Do not delete the intermediate instance and the ESSD while the image cache is being created. If you delete the intermediate instance or the ESSD, the image cache cannot be created.
 * - A temporary local snapshot and a specific number of regular snapshots are generated during the creation of the image cache. Do not delete these snapshots. If you delete these snapshots, the image cache becomes invalid.
 * - If you use SDKs, SDK for Java 1.0.10 or later and SDK for Python 1.0.7 or later are supported.
 * ### Usage notes
 * - You can configure AcrRegistryInfo-related parameters to pull images from Container Registry Enterprise Edition instances without using a password. When you use AcrRegistryInfo-related parameters to pull images from a Container Registry Enterprise Edition instance without using a password, you must specify the AcrRegistryInfo.N.InstanceId parameter.
 * - If the image cache that you created will be used to create more than 1,000 elastic container instances at a time, we recommend that you use the StandardCopyCount and FlashCopyCount parameters to create multiple temporary local snapshots and regular snapshots for the image cache. The multiple snapshots are billed based on incremental data. If no incremental data exists on the multiple snapshots, you are not charged for the multiple snapshots.
 * >  When you call the CreateImageCache operation to create an image cache, the system automatically creates a service-linked role named AliyunServiceRoleForECI. The role is used to access other Alibaba Cloud services such as Elastic Compute Service (ECS) and Virtual Private Cloud (VPC). For more information, see [Elastic Container Instance service-linked role](~~212914~~).
 *
 * @param request CreateImageCacheRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateImageCacheResponse
 */
func (client *Client) CreateImageCacheWithOptions(request *CreateImageCacheRequest, runtime *util.RuntimeOptions) (_result *CreateImageCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcrRegistryInfo)) {
		query["AcrRegistryInfo"] = request.AcrRegistryInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Annotations)) {
		query["Annotations"] = request.Annotations
	}

	if !tea.BoolValue(util.IsUnset(request.AutoMatchImageCache)) {
		query["AutoMatchImageCache"] = request.AutoMatchImageCache
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EipInstanceId)) {
		query["EipInstanceId"] = request.EipInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EliminationStrategy)) {
		query["EliminationStrategy"] = request.EliminationStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.Flash)) {
		query["Flash"] = request.Flash
	}

	if !tea.BoolValue(util.IsUnset(request.FlashCopyCount)) {
		query["FlashCopyCount"] = request.FlashCopyCount
	}

	if !tea.BoolValue(util.IsUnset(request.Image)) {
		query["Image"] = request.Image
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCacheName)) {
		query["ImageCacheName"] = request.ImageCacheName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCacheSize)) {
		query["ImageCacheSize"] = request.ImageCacheSize
	}

	if !tea.BoolValue(util.IsUnset(request.ImageRegistryCredential)) {
		query["ImageRegistryCredential"] = request.ImageRegistryCredential
	}

	if !tea.BoolValue(util.IsUnset(request.InsecureRegistry)) {
		query["InsecureRegistry"] = request.InsecureRegistry
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlainHttpRegistry)) {
		query["PlainHttpRegistry"] = request.PlainHttpRegistry
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RetentionDays)) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StandardCopyCount)) {
		query["StandardCopyCount"] = request.StandardCopyCount
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateImageCache"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateImageCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Precautions
 * - You are charged for creation of image caches. We recommend that you learn the relevant billing information in advance. For more information, see [Billing](~~89142~~).
 * - Before you create an image cache, you must estimate the total size of the images that you want to cache. If the total size of the images exceeds the specified cache size, the image cache cannot be created.
 * - When an image cache is being created, the system creates an intermediate elastic container instance and an intermediate enhanced SSD (ESSD) at performance level 1 (PL1). Do not delete the intermediate instance and the ESSD while the image cache is being created. If you delete the intermediate instance or the ESSD, the image cache cannot be created.
 * - A temporary local snapshot and a specific number of regular snapshots are generated during the creation of the image cache. Do not delete these snapshots. If you delete these snapshots, the image cache becomes invalid.
 * - If you use SDKs, SDK for Java 1.0.10 or later and SDK for Python 1.0.7 or later are supported.
 * ### Usage notes
 * - You can configure AcrRegistryInfo-related parameters to pull images from Container Registry Enterprise Edition instances without using a password. When you use AcrRegistryInfo-related parameters to pull images from a Container Registry Enterprise Edition instance without using a password, you must specify the AcrRegistryInfo.N.InstanceId parameter.
 * - If the image cache that you created will be used to create more than 1,000 elastic container instances at a time, we recommend that you use the StandardCopyCount and FlashCopyCount parameters to create multiple temporary local snapshots and regular snapshots for the image cache. The multiple snapshots are billed based on incremental data. If no incremental data exists on the multiple snapshots, you are not charged for the multiple snapshots.
 * >  When you call the CreateImageCache operation to create an image cache, the system automatically creates a service-linked role named AliyunServiceRoleForECI. The role is used to access other Alibaba Cloud services such as Elastic Compute Service (ECS) and Virtual Private Cloud (VPC). For more information, see [Elastic Container Instance service-linked role](~~212914~~).
 *
 * @param request CreateImageCacheRequest
 * @return CreateImageCacheResponse
 */
func (client *Client) CreateImageCache(request *CreateImageCacheRequest) (_result *CreateImageCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageCacheResponse{}
	_body, _err := client.CreateImageCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * O&M tasks are classified into:
 * *   coredump: After you enable coredump, the system generates a core dump file when a container unexpectedly stops. You can use the core dump file to analyze the exception and find out the cause of the problem. For more information, see [Enable coredump](~~167801~~).
 * *   tcpdump: After you enable tcpdump, the system captures network packets when a container unexpectedly stops. You can analyze the packets and locate network problems. For more information, see Enable [tcpdump](~~429749~~).
 *
 * @param request CreateInstanceOpsTaskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateInstanceOpsTaskResponse
 */
func (client *Client) CreateInstanceOpsTaskWithOptions(request *CreateInstanceOpsTaskRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceOpsTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OpsType)) {
		query["OpsType"] = request.OpsType
	}

	if !tea.BoolValue(util.IsUnset(request.OpsValue)) {
		query["OpsValue"] = request.OpsValue
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceOpsTask"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceOpsTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * O&M tasks are classified into:
 * *   coredump: After you enable coredump, the system generates a core dump file when a container unexpectedly stops. You can use the core dump file to analyze the exception and find out the cause of the problem. For more information, see [Enable coredump](~~167801~~).
 * *   tcpdump: After you enable tcpdump, the system captures network packets when a container unexpectedly stops. You can analyze the packets and locate network problems. For more information, see Enable [tcpdump](~~429749~~).
 *
 * @param request CreateInstanceOpsTaskRequest
 * @return CreateInstanceOpsTaskResponse
 */
func (client *Client) CreateInstanceOpsTask(request *CreateInstanceOpsTaskRequest) (_result *CreateInstanceOpsTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceOpsTaskResponse{}
	_body, _err := client.CreateInstanceOpsTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * When you call this operation to create a VNode, the system automatically creates a service-linked role AliyunServiceRoleForECIVnode for you to access relevant cloud services such as Elastic Container Instance, Elastic Compute Service (ECS), and Virtual Private Cloud (VPC). For more information, see [Virtual Node Service-linked role](~~311014~~).
 *
 * @param request CreateVirtualNodeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateVirtualNodeResponse
 */
func (client *Client) CreateVirtualNodeWithOptions(request *CreateVirtualNodeRequest, runtime *util.RuntimeOptions) (_result *CreateVirtualNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterDNS)) {
		query["ClusterDNS"] = request.ClusterDNS
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterDomain)) {
		query["ClusterDomain"] = request.ClusterDomain
	}

	if !tea.BoolValue(util.IsUnset(request.CustomResources)) {
		query["CustomResources"] = request.CustomResources
	}

	if !tea.BoolValue(util.IsUnset(request.EipInstanceId)) {
		query["EipInstanceId"] = request.EipInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePublicNetwork)) {
		query["EnablePublicNetwork"] = request.EnablePublicNetwork
	}

	if !tea.BoolValue(util.IsUnset(request.KubeConfig)) {
		query["KubeConfig"] = request.KubeConfig
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RotateCertificateEnabled)) {
		query["RotateCertificateEnabled"] = request.RotateCertificateEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.Taint)) {
		query["Taint"] = request.Taint
	}

	if !tea.BoolValue(util.IsUnset(request.TlsBootstrapEnabled)) {
		query["TlsBootstrapEnabled"] = request.TlsBootstrapEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualNodeName)) {
		query["VirtualNodeName"] = request.VirtualNodeName
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVirtualNode"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVirtualNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * When you call this operation to create a VNode, the system automatically creates a service-linked role AliyunServiceRoleForECIVnode for you to access relevant cloud services such as Elastic Container Instance, Elastic Compute Service (ECS), and Virtual Private Cloud (VPC). For more information, see [Virtual Node Service-linked role](~~311014~~).
 *
 * @param request CreateVirtualNodeRequest
 * @return CreateVirtualNodeResponse
 */
func (client *Client) CreateVirtualNode(request *CreateVirtualNodeRequest) (_result *CreateVirtualNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVirtualNodeResponse{}
	_body, _err := client.CreateVirtualNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteContainerGroupWithOptions(request *DeleteContainerGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteContainerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteContainerGroup"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteContainerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteContainerGroup(request *DeleteContainerGroupRequest) (_result *DeleteContainerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteContainerGroupResponse{}
	_body, _err := client.DeleteContainerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageCacheWithOptions(request *DeleteImageCacheRequest, runtime *util.RuntimeOptions) (_result *DeleteImageCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCacheId)) {
		query["ImageCacheId"] = request.ImageCacheId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteImageCache"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteImageCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImageCache(request *DeleteImageCacheRequest) (_result *DeleteImageCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImageCacheResponse{}
	_body, _err := client.DeleteImageCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Make sure that no elastic container instance exists on the virtual node before you call this operation to delete a virtual node.
 *
 * @param request DeleteVirtualNodeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteVirtualNodeResponse
 */
func (client *Client) DeleteVirtualNodeWithOptions(request *DeleteVirtualNodeRequest, runtime *util.RuntimeOptions) (_result *DeleteVirtualNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualNodeId)) {
		query["VirtualNodeId"] = request.VirtualNodeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVirtualNode"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVirtualNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Make sure that no elastic container instance exists on the virtual node before you call this operation to delete a virtual node.
 *
 * @param request DeleteVirtualNodeRequest
 * @return DeleteVirtualNodeResponse
 */
func (client *Client) DeleteVirtualNode(request *DeleteVirtualNodeRequest) (_result *DeleteVirtualNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVirtualNodeResponse{}
	_body, _err := client.DeleteVirtualNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * When you call the CreateContainerGroup operation to create an elastic container instance, you can use the InstanceType parameter to specify one or more ECS instance types that fit your specific needs. To ensure that the elastic container instance can be created, you can call the DescribeAvailableResource operation to query which ECS instance types and instance families are available in the specified region and zone before you create the elastic container instance.
 *
 * @param request DescribeAvailableResourceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeAvailableResourceResponse
 */
func (client *Client) DescribeAvailableResourceWithOptions(request *DescribeAvailableResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationResource)) {
		query["DestinationResource"] = request.DestinationResource
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SpotResource)) {
		query["SpotResource"] = request.SpotResource
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAvailableResource"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * When you call the CreateContainerGroup operation to create an elastic container instance, you can use the InstanceType parameter to specify one or more ECS instance types that fit your specific needs. To ensure that the elastic container instance can be created, you can call the DescribeAvailableResource operation to query which ECS instance types and instance families are available in the specified region and zone before you create the elastic container instance.
 *
 * @param request DescribeAvailableResourceRequest
 * @return DescribeAvailableResourceResponse
 */
func (client *Client) DescribeAvailableResource(request *DescribeAvailableResourceRequest) (_result *DescribeAvailableResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DescribeAvailableResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCommitContainerTaskWithOptions(request *DescribeCommitContainerTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeCommitContainerTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskStatus)) {
		query["TaskStatus"] = request.TaskStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCommitContainerTask"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCommitContainerTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCommitContainerTask(request *DescribeCommitContainerTaskRequest) (_result *DescribeCommitContainerTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCommitContainerTaskResponse{}
	_body, _err := client.DescribeCommitContainerTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the event information of multiple container groups at a time. By default, the latest 50 entries of events of each container group are returned.
 *
 * @param request DescribeContainerGroupEventsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeContainerGroupEventsResponse
 */
func (client *Client) DescribeContainerGroupEventsWithOptions(request *DescribeContainerGroupEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerGroupEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerGroupIds)) {
		query["ContainerGroupIds"] = request.ContainerGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.EventSource)) {
		query["EventSource"] = request.EventSource
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SinceSecond)) {
		query["SinceSecond"] = request.SinceSecond
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContainerGroupEvents"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContainerGroupEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the event information of multiple container groups at a time. By default, the latest 50 entries of events of each container group are returned.
 *
 * @param request DescribeContainerGroupEventsRequest
 * @return DescribeContainerGroupEventsResponse
 */
func (client *Client) DescribeContainerGroupEvents(request *DescribeContainerGroupEventsRequest) (_result *DescribeContainerGroupEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerGroupEventsResponse{}
	_body, _err := client.DescribeContainerGroupEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   A maximum of 50 monitoring data entries can be returned. If the number of monitoring data entries exceeds this limit, an error message is returned.
 * *   You can query real-time monitoring data (data generated within the last 5 minutes) and historical data (data generated more than 5 minutes ago). If the time range to query starts or ends later than the current time, historical monitoring data generated more than 5 minutes ago is returned.
 * *   You can query only the monitoring data of elastic container instances that are created after April 3, 2019 15:00 UTC+8.
 *
 * @param request DescribeContainerGroupMetricRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeContainerGroupMetricResponse
 */
func (client *Client) DescribeContainerGroupMetricWithOptions(request *DescribeContainerGroupMetricRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerGroupMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContainerGroupMetric"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContainerGroupMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   A maximum of 50 monitoring data entries can be returned. If the number of monitoring data entries exceeds this limit, an error message is returned.
 * *   You can query real-time monitoring data (data generated within the last 5 minutes) and historical data (data generated more than 5 minutes ago). If the time range to query starts or ends later than the current time, historical monitoring data generated more than 5 minutes ago is returned.
 * *   You can query only the monitoring data of elastic container instances that are created after April 3, 2019 15:00 UTC+8.
 *
 * @param request DescribeContainerGroupMetricRequest
 * @return DescribeContainerGroupMetricResponse
 */
func (client *Client) DescribeContainerGroupMetric(request *DescribeContainerGroupMetricRequest) (_result *DescribeContainerGroupMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerGroupMetricResponse{}
	_body, _err := client.DescribeContainerGroupMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This operation does not support resource group authentication.
 *
 * @param request DescribeContainerGroupPriceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeContainerGroupPriceResponse
 */
func (client *Client) DescribeContainerGroupPriceWithOptions(request *DescribeContainerGroupPriceRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerGroupPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		query["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.EphemeralStorage)) {
		query["EphemeralStorage"] = request.EphemeralStorage
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		query["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SpotDuration)) {
		query["SpotDuration"] = request.SpotDuration
	}

	if !tea.BoolValue(util.IsUnset(request.SpotPriceLimit)) {
		query["SpotPriceLimit"] = request.SpotPriceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.SpotStrategy)) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContainerGroupPrice"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContainerGroupPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation does not support resource group authentication.
 *
 * @param request DescribeContainerGroupPriceRequest
 * @return DescribeContainerGroupPriceResponse
 */
func (client *Client) DescribeContainerGroupPrice(request *DescribeContainerGroupPriceRequest) (_result *DescribeContainerGroupPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerGroupPriceResponse{}
	_body, _err := client.DescribeContainerGroupPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerGroupStatusWithOptions(request *DescribeContainerGroupStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerGroupStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerGroupIds)) {
		query["ContainerGroupIds"] = request.ContainerGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SinceSecond)) {
		query["SinceSecond"] = request.SinceSecond
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContainerGroupStatus"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContainerGroupStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerGroupStatus(request *DescribeContainerGroupStatusRequest) (_result *DescribeContainerGroupStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerGroupStatusResponse{}
	_body, _err := client.DescribeContainerGroupStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After elastic container instances stop running, the underlying computing resources are reclaimed. The instance information is retained based on the following rules:
 * *   For elastic container instances that stop on success, only the latest 100 entries of success information about instances in all regions is retained.
 * *   For elastic container instances that stop on failure, the instance information is retained for only 24 hours.
 *
 * @param request DescribeContainerGroupsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeContainerGroupsResponse
 */
func (client *Client) DescribeContainerGroupsWithOptions(request *DescribeContainerGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerGroupIds)) {
		query["ContainerGroupIds"] = request.ContainerGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupName)) {
		query["ContainerGroupName"] = request.ContainerGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.WithEvent)) {
		query["WithEvent"] = request.WithEvent
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContainerGroups"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContainerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After elastic container instances stop running, the underlying computing resources are reclaimed. The instance information is retained based on the following rules:
 * *   For elastic container instances that stop on success, only the latest 100 entries of success information about instances in all regions is retained.
 * *   For elastic container instances that stop on failure, the instance information is retained for only 24 hours.
 *
 * @param request DescribeContainerGroupsRequest
 * @return DescribeContainerGroupsResponse
 */
func (client *Client) DescribeContainerGroups(request *DescribeContainerGroupsRequest) (_result *DescribeContainerGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerGroupsResponse{}
	_body, _err := client.DescribeContainerGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerLogWithOptions(request *DescribeContainerLogRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerName)) {
		query["ContainerName"] = request.ContainerName
	}

	if !tea.BoolValue(util.IsUnset(request.LastTime)) {
		query["LastTime"] = request.LastTime
	}

	if !tea.BoolValue(util.IsUnset(request.LimitBytes)) {
		query["LimitBytes"] = request.LimitBytes
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SinceSeconds)) {
		query["SinceSeconds"] = request.SinceSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Tail)) {
		query["Tail"] = request.Tail
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamps)) {
		query["Timestamps"] = request.Timestamps
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContainerLog"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContainerLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerLog(request *DescribeContainerLogRequest) (_result *DescribeContainerLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerLogResponse{}
	_body, _err := client.DescribeContainerLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageCachesWithOptions(request *DescribeImageCachesRequest, runtime *util.RuntimeOptions) (_result *DescribeImageCachesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Image)) {
		query["Image"] = request.Image
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCacheId)) {
		query["ImageCacheId"] = request.ImageCacheId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCacheName)) {
		query["ImageCacheName"] = request.ImageCacheName
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.MatchImage)) {
		query["MatchImage"] = request.MatchImage
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImageCaches"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeImageCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImageCaches(request *DescribeImageCachesRequest) (_result *DescribeImageCachesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageCachesResponse{}
	_body, _err := client.DescribeImageCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceOpsRecordsWithOptions(request *DescribeInstanceOpsRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceOpsRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OpsType)) {
		query["OpsType"] = request.OpsType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceOpsRecords"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceOpsRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceOpsRecords(request *DescribeInstanceOpsRecordsRequest) (_result *DescribeInstanceOpsRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceOpsRecordsResponse{}
	_body, _err := client.DescribeInstanceOpsRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * - Only the latest entry of monitoring data of each elastic container instance is returned.
 * - You can query only the monitoring data of elastic container instances that are created after April 3, 2019 15:00:00 UTC+8.
 *
 * @param request DescribeMultiContainerGroupMetricRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeMultiContainerGroupMetricResponse
 */
func (client *Client) DescribeMultiContainerGroupMetricWithOptions(request *DescribeMultiContainerGroupMetricRequest, runtime *util.RuntimeOptions) (_result *DescribeMultiContainerGroupMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerGroupIds)) {
		query["ContainerGroupIds"] = request.ContainerGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMultiContainerGroupMetric"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMultiContainerGroupMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * - Only the latest entry of monitoring data of each elastic container instance is returned.
 * - You can query only the monitoring data of elastic container instances that are created after April 3, 2019 15:00:00 UTC+8.
 *
 * @param request DescribeMultiContainerGroupMetricRequest
 * @return DescribeMultiContainerGroupMetricResponse
 */
func (client *Client) DescribeMultiContainerGroupMetric(request *DescribeMultiContainerGroupMetricRequest) (_result *DescribeMultiContainerGroupMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMultiContainerGroupMetricResponse{}
	_body, _err := client.DescribeMultiContainerGroupMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVirtualNodesWithOptions(request *DescribeVirtualNodesRequest, runtime *util.RuntimeOptions) (_result *DescribeVirtualNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualNodeIds)) {
		query["VirtualNodeIds"] = request.VirtualNodeIds
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualNodeName)) {
		query["VirtualNodeName"] = request.VirtualNodeName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVirtualNodes"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVirtualNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVirtualNodes(request *DescribeVirtualNodesRequest) (_result *DescribeVirtualNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVirtualNodesResponse{}
	_body, _err := client.DescribeVirtualNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecContainerCommandWithOptions(request *ExecContainerCommandRequest, runtime *util.RuntimeOptions) (_result *ExecContainerCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Command)) {
		query["Command"] = request.Command
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerName)) {
		query["ContainerName"] = request.ContainerName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Stdin)) {
		query["Stdin"] = request.Stdin
	}

	if !tea.BoolValue(util.IsUnset(request.Sync)) {
		query["Sync"] = request.Sync
	}

	if !tea.BoolValue(util.IsUnset(request.TTY)) {
		query["TTY"] = request.TTY
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecContainerCommand"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecContainerCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecContainerCommand(request *ExecContainerCommandRequest) (_result *ExecContainerCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecContainerCommandResponse{}
	_body, _err := client.ExecContainerCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This operation cannot be used for resource group authentication.
 *
 * @param request ListUsageRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListUsageResponse
 */
func (client *Client) ListUsageWithOptions(request *ListUsageRequest, runtime *util.RuntimeOptions) (_result *ListUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsage"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation cannot be used for resource group authentication.
 *
 * @param request ListUsageRequest
 * @return ListUsageResponse
 */
func (client *Client) ListUsage(request *ListUsageRequest) (_result *ListUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsageResponse{}
	_body, _err := client.ListUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation only to scale out volumes, not to scale in volumes. You can scale out only volumes of the Alibaba Cloud disk type. Volumes of other types cannot be scaled out.
 *
 * @param request ResizeContainerGroupVolumeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ResizeContainerGroupVolumeResponse
 */
func (client *Client) ResizeContainerGroupVolumeWithOptions(request *ResizeContainerGroupVolumeRequest, runtime *util.RuntimeOptions) (_result *ResizeContainerGroupVolumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.NewSize)) {
		query["NewSize"] = request.NewSize
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VolumeName)) {
		query["VolumeName"] = request.VolumeName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResizeContainerGroupVolume"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResizeContainerGroupVolumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation only to scale out volumes, not to scale in volumes. You can scale out only volumes of the Alibaba Cloud disk type. Volumes of other types cannot be scaled out.
 *
 * @param request ResizeContainerGroupVolumeRequest
 * @return ResizeContainerGroupVolumeResponse
 */
func (client *Client) ResizeContainerGroupVolume(request *ResizeContainerGroupVolumeRequest) (_result *ResizeContainerGroupVolumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResizeContainerGroupVolumeResponse{}
	_body, _err := client.ResizeContainerGroupVolumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Only elastic container instances that are in the Pending or Running state can be restarted. Instances that are in the Succeeded or Failed state cannot be restarted.
 * *   Elastic container instances that were created before 15:00:00 on March 7, 2019 cannot be restarted.
 * *   When an elastic container instance is being restarted, its status changes into Restarting.
 *
 * @param request RestartContainerGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RestartContainerGroupResponse
 */
func (client *Client) RestartContainerGroupWithOptions(request *RestartContainerGroupRequest, runtime *util.RuntimeOptions) (_result *RestartContainerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartContainerGroup"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartContainerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Only elastic container instances that are in the Pending or Running state can be restarted. Instances that are in the Succeeded or Failed state cannot be restarted.
 * *   Elastic container instances that were created before 15:00:00 on March 7, 2019 cannot be restarted.
 * *   When an elastic container instance is being restarted, its status changes into Restarting.
 *
 * @param request RestartContainerGroupRequest
 * @return RestartContainerGroupResponse
 */
func (client *Client) RestartContainerGroup(request *RestartContainerGroupRequest) (_result *RestartContainerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartContainerGroupResponse{}
	_body, _err := client.RestartContainerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * * You can update only elastic container instances that are in the Pending or Running state. After you call this operation to update an elastic container instance, the instance enters the Updating state.
 * * You cannot update elastic container instances that were created before 15:00:00 March 7, 2019.
 *
 * @param request UpdateContainerGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateContainerGroupResponse
 */
func (client *Client) UpdateContainerGroupWithOptions(request *UpdateContainerGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateContainerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcrRegistryInfo)) {
		query["AcrRegistryInfo"] = request.AcrRegistryInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Container)) {
		query["Container"] = request.Container
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		query["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.ImageRegistryCredential)) {
		query["ImageRegistryCredential"] = request.ImageRegistryCredential
	}

	if !tea.BoolValue(util.IsUnset(request.InitContainer)) {
		query["InitContainer"] = request.InitContainer
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		query["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RestartPolicy)) {
		query["RestartPolicy"] = request.RestartPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateType)) {
		query["UpdateType"] = request.UpdateType
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		query["Volume"] = request.Volume
	}

	if !tea.BoolValue(util.IsUnset(request.DnsConfig)) {
		query["DnsConfig"] = request.DnsConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateContainerGroup"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateContainerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * * You can update only elastic container instances that are in the Pending or Running state. After you call this operation to update an elastic container instance, the instance enters the Updating state.
 * * You cannot update elastic container instances that were created before 15:00:00 March 7, 2019.
 *
 * @param request UpdateContainerGroupRequest
 * @return UpdateContainerGroupResponse
 */
func (client *Client) UpdateContainerGroup(request *UpdateContainerGroupRequest) (_result *UpdateContainerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateContainerGroupResponse{}
	_body, _err := client.UpdateContainerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Only image caches that are in the Ready or UpdateFailed state can be updated.
 *
 * @param request UpdateImageCacheRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateImageCacheResponse
 */
func (client *Client) UpdateImageCacheWithOptions(request *UpdateImageCacheRequest, runtime *util.RuntimeOptions) (_result *UpdateImageCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcrRegistryInfo)) {
		query["AcrRegistryInfo"] = request.AcrRegistryInfo
	}

	if !tea.BoolValue(util.IsUnset(request.AutoMatchImageCache)) {
		query["AutoMatchImageCache"] = request.AutoMatchImageCache
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EipInstanceId)) {
		query["EipInstanceId"] = request.EipInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EliminationStrategy)) {
		query["EliminationStrategy"] = request.EliminationStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.Flash)) {
		query["Flash"] = request.Flash
	}

	if !tea.BoolValue(util.IsUnset(request.FlashCopyCount)) {
		query["FlashCopyCount"] = request.FlashCopyCount
	}

	if !tea.BoolValue(util.IsUnset(request.Image)) {
		query["Image"] = request.Image
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCacheId)) {
		query["ImageCacheId"] = request.ImageCacheId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCacheName)) {
		query["ImageCacheName"] = request.ImageCacheName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCacheSize)) {
		query["ImageCacheSize"] = request.ImageCacheSize
	}

	if !tea.BoolValue(util.IsUnset(request.ImageRegistryCredential)) {
		query["ImageRegistryCredential"] = request.ImageRegistryCredential
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RetentionDays)) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StandardCopyCount)) {
		query["StandardCopyCount"] = request.StandardCopyCount
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateImageCache"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateImageCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Only image caches that are in the Ready or UpdateFailed state can be updated.
 *
 * @param request UpdateImageCacheRequest
 * @return UpdateImageCacheResponse
 */
func (client *Client) UpdateImageCache(request *UpdateImageCacheRequest) (_result *UpdateImageCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateImageCacheResponse{}
	_body, _err := client.UpdateImageCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Only virtual nodes that are in the Ready state can be updated.
 *
 * @param request UpdateVirtualNodeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateVirtualNodeResponse
 */
func (client *Client) UpdateVirtualNodeWithOptions(request *UpdateVirtualNodeRequest, runtime *util.RuntimeOptions) (_result *UpdateVirtualNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterDNS)) {
		query["ClusterDNS"] = request.ClusterDNS
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterDomain)) {
		query["ClusterDomain"] = request.ClusterDomain
	}

	if !tea.BoolValue(util.IsUnset(request.CustomResources)) {
		query["CustomResources"] = request.CustomResources
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualNodeId)) {
		query["VirtualNodeId"] = request.VirtualNodeId
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualNodeName)) {
		query["VirtualNodeName"] = request.VirtualNodeName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateVirtualNode"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVirtualNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Only virtual nodes that are in the Ready state can be updated.
 *
 * @param request UpdateVirtualNodeRequest
 * @return UpdateVirtualNodeResponse
 */
func (client *Client) UpdateVirtualNode(request *UpdateVirtualNodeRequest) (_result *UpdateVirtualNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVirtualNodeResponse{}
	_body, _err := client.UpdateVirtualNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
