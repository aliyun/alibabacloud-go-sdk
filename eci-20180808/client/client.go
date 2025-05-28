// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CommitContainerRequest struct {
	// The access credential configurations of the Container Registry Enterprise Edition instance.
	//
	// >  If you use a Container Registry Personal Edition instance, you do not need to configure this parameter. If you use a Container Registry Enterprise Edition instance, you must configure this parameter.
	AcrRegistryInfo *CommitContainerRequestAcrRegistryInfo `json:"AcrRegistryInfo,omitempty" xml:"AcrRegistryInfo,omitempty" type:"Struct"`
	// The details about the ARN that is required for authorization.
	Arn *CommitContainerRequestArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Struct"`
	// The ID of the container group.
	//
	// This parameter is required.
	//
	// example:
	//
	// eci-bp1do4xz75fa5sd****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The name of the container.
	//
	// This parameter is required.
	//
	// example:
	//
	// container-1
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The image of the container.
	Image        *CommitContainerRequestImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	OwnerAccount *string                      `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	// The Alibaba Cloud Resource Access (ARN) of the RAM role that is assigned to the user (the authorized account) in cross-account authorization scenarios.
	//
	// example:
	//
	// acs:ram::1609982529******:role/role-test
	ArnService *string `json:"ArnService,omitempty" xml:"ArnService,omitempty"`
	// The ARN of the RAM role that is assigned to the authorizer in cross-account authorization scenarios.
	//
	// example:
	//
	// acs:ram::1298452580******:role/role-test
	ArnUser *string `json:"ArnUser,omitempty" xml:"ArnUser,omitempty"`
	// The ID of the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// cri-nwj395hgf6f3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CommitContainerRequestAcrRegistryInfo) String() string {
	return tea.Prettify(s)
}

func (s CommitContainerRequestAcrRegistryInfo) GoString() string {
	return s.String()
}

func (s *CommitContainerRequestAcrRegistryInfo) SetArnService(v string) *CommitContainerRequestAcrRegistryInfo {
	s.ArnService = &v
	return s
}

func (s *CommitContainerRequestAcrRegistryInfo) SetArnUser(v string) *CommitContainerRequestAcrRegistryInfo {
	s.ArnUser = &v
	return s
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
	// The ARN of the authorized role.
	//
	// example:
	//
	// acs:ram:xxx
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The authorization type. A value of service indicates that RAM roles are used for authorization.
	//
	// example:
	//
	// user
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
	//
	// example:
	//
	// ECI
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// The message about the image.
	//
	// example:
	//
	// test commit
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/eastest/registry-test
	Repository *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	// The tag of the image. This parameter is empty by default, which indicates that the tag is not modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.6
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
	// The ID of the request.
	//
	// example:
	//
	// CB8D2B22-D636-4182-****-1FC9DBDAD66F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// t-bp18cwvrb3v9tjk0tekz
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CommitContainerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CopyDataCacheRequest struct {
	// The bucket in which the DataCache is stored.
	//
	// example:
	//
	// default
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the DataCache in the source region.
	//
	// This parameter is required.
	//
	// example:
	//
	// edc-bp15l4vvys94oo******
	DataCacheId *string `json:"DataCacheId,omitempty" xml:"DataCacheId,omitempty"`
	// The destination region of the DataCache.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// The DataCache name.
	//
	// example:
	//
	// test
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The storage path of the data.
	//
	// example:
	//
	// /bucket-test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The source region of the DataCache.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group to which the DataCache belongs.
	//
	// example:
	//
	// rg-2df3isufhi38****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The number of days for which the DataCache is retained.
	//
	// example:
	//
	// 7
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The tags of the DataCache.
	Tag []*CopyDataCacheRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CopyDataCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyDataCacheRequest) GoString() string {
	return s.String()
}

func (s *CopyDataCacheRequest) SetBucket(v string) *CopyDataCacheRequest {
	s.Bucket = &v
	return s
}

func (s *CopyDataCacheRequest) SetClientToken(v string) *CopyDataCacheRequest {
	s.ClientToken = &v
	return s
}

func (s *CopyDataCacheRequest) SetDataCacheId(v string) *CopyDataCacheRequest {
	s.DataCacheId = &v
	return s
}

func (s *CopyDataCacheRequest) SetDestinationRegionId(v string) *CopyDataCacheRequest {
	s.DestinationRegionId = &v
	return s
}

func (s *CopyDataCacheRequest) SetName(v string) *CopyDataCacheRequest {
	s.Name = &v
	return s
}

func (s *CopyDataCacheRequest) SetOwnerAccount(v string) *CopyDataCacheRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CopyDataCacheRequest) SetOwnerId(v int64) *CopyDataCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *CopyDataCacheRequest) SetPath(v string) *CopyDataCacheRequest {
	s.Path = &v
	return s
}

func (s *CopyDataCacheRequest) SetRegionId(v string) *CopyDataCacheRequest {
	s.RegionId = &v
	return s
}

func (s *CopyDataCacheRequest) SetResourceGroupId(v string) *CopyDataCacheRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CopyDataCacheRequest) SetResourceOwnerAccount(v string) *CopyDataCacheRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CopyDataCacheRequest) SetResourceOwnerId(v int64) *CopyDataCacheRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CopyDataCacheRequest) SetRetentionDays(v int32) *CopyDataCacheRequest {
	s.RetentionDays = &v
	return s
}

func (s *CopyDataCacheRequest) SetTag(v []*CopyDataCacheRequestTag) *CopyDataCacheRequest {
	s.Tag = v
	return s
}

type CopyDataCacheRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// testkey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// testvalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CopyDataCacheRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CopyDataCacheRequestTag) GoString() string {
	return s.String()
}

func (s *CopyDataCacheRequestTag) SetKey(v string) *CopyDataCacheRequestTag {
	s.Key = &v
	return s
}

func (s *CopyDataCacheRequestTag) SetValue(v string) *CopyDataCacheRequestTag {
	s.Value = &v
	return s
}

type CopyDataCacheResponseBody struct {
	// The ID generated for the DataCache in the destination region.
	//
	// example:
	//
	// edc-bp1423y6d7v7l6ua****
	DataCacheId *string `json:"DataCacheId,omitempty" xml:"DataCacheId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 58EE0CB3-C864-5395-A4F7-24F425074839
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyDataCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyDataCacheResponseBody) GoString() string {
	return s.String()
}

func (s *CopyDataCacheResponseBody) SetDataCacheId(v string) *CopyDataCacheResponseBody {
	s.DataCacheId = &v
	return s
}

func (s *CopyDataCacheResponseBody) SetRequestId(v string) *CopyDataCacheResponseBody {
	s.RequestId = &v
	return s
}

type CopyDataCacheResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyDataCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyDataCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyDataCacheResponse) GoString() string {
	return s.String()
}

func (s *CopyDataCacheResponse) SetHeaders(v map[string]*string) *CopyDataCacheResponse {
	s.Headers = v
	return s
}

func (s *CopyDataCacheResponse) SetStatusCode(v int32) *CopyDataCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyDataCacheResponse) SetBody(v *CopyDataCacheResponseBody) *CopyDataCacheResponse {
	s.Body = v
	return s
}

type CreateContainerGroupRequest struct {
	DnsConfig           *CreateContainerGroupRequestDnsConfig           `json:"DnsConfig,omitempty" xml:"DnsConfig,omitempty" type:"Struct"`
	HostSecurityContext *CreateContainerGroupRequestHostSecurityContext `json:"HostSecurityContext,omitempty" xml:"HostSecurityContext,omitempty" type:"Struct"`
	SecurityContext     *CreateContainerGroupRequestSecurityContext     `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	// The information about the Container Registry Enterprise Edition instance that provides the image for the creation of the elastic container instance. For more information, see [Pull images from a Container Registry Enterprise Edition instance without using a secret](https://help.aliyun.com/document_detail/194250.html).
	AcrRegistryInfo []*CreateContainerGroupRequestAcrRegistryInfo `json:"AcrRegistryInfo,omitempty" xml:"AcrRegistryInfo,omitempty" type:"Repeated"`
	// The active period of the elastic container instance. After this period expires, the instance is forced to exit. Unit: seconds.
	//
	// example:
	//
	// 1000
	ActiveDeadlineSeconds *int64 `json:"ActiveDeadlineSeconds,omitempty" xml:"ActiveDeadlineSeconds,omitempty"`
	// Specifies whether to automatically create an EIP and associate it with the elastic container instance.
	//
	// example:
	//
	// true
	AutoCreateEip *bool `json:"AutoCreateEip,omitempty" xml:"AutoCreateEip,omitempty"`
	// Specifies whether to automatically match image caches. Default value: false.
	//
	// example:
	//
	// false
	AutoMatchImageCache *bool `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-xxxx-12d3-xxxx-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The compute category of the instance. For more information, see [Specify a compute category to create an elastic container instance](https://help.aliyun.com/document_detail/2638061.html).
	ComputeCategory []*string `json:"ComputeCategory,omitempty" xml:"ComputeCategory,omitempty" type:"Repeated"`
	// The information about the container.
	//
	// This parameter is required.
	Container []*CreateContainerGroupRequestContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Repeated"`
	// The name of the elastic container instance (container group). The name must meet the following requirements:
	//
	// 	- The name must be 2 to 128 characters in length.
	//
	// 	- The name can contain lowercase letters, digits, and hyphens (-). It cannot start or end with a hyphen (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// nginx-test
	ContainerGroupName *string `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	// Specifies whether to enable container resource view. Container resource view displays the actual container resource data instead of data of the host. If the specifications of the generated elastic container instance are larger than the specifications that you request for when you create the instance, you can enable the ContainerResourceView feature to ensure that the resources that you view in the container are the same as the resources that you request for.
	//
	// example:
	//
	// false
	ContainerResourceView *bool `json:"ContainerResourceView,omitempty" xml:"ContainerResourceView,omitempty"`
	// The path to core dump files. For more information, see [Save core files to volumes](https://help.aliyun.com/document_detail/167801.html).
	//
	// >  The path cannot start with |. You cannot use core dump files to configure executable programs.``
	//
	// example:
	//
	// /xx/xx/core
	CorePattern *string `json:"CorePattern,omitempty" xml:"CorePattern,omitempty"`
	// The number of vCPUs that you want to allocate to the instance.
	//
	// example:
	//
	// 1.0
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The CPU architecture of the instance. Default value: AMD64. Valid values:
	//
	// 	- AMD64
	//
	// 	- ARM64
	//
	// example:
	//
	// ARM64
	CpuArchitecture *string `json:"CpuArchitecture,omitempty" xml:"CpuArchitecture,omitempty"`
	// The number of physical CPU cores. You can specify this parameter for only specific ECS instance types.
	//
	// example:
	//
	// 2
	CpuOptionsCore *int32 `json:"CpuOptionsCore,omitempty" xml:"CpuOptionsCore,omitempty"`
	// This parameter is not available.
	//
	// example:
	//
	// 1
	CpuOptionsNuma *string `json:"CpuOptionsNuma,omitempty" xml:"CpuOptionsNuma,omitempty"`
	// The number of threads per core. You can specify this parameter for only specific ECS instance types. A value of 1 specifies that Hyper-Threading is disabled.
	//
	// example:
	//
	// 2
	CpuOptionsThreadsPerCore *int32 `json:"CpuOptionsThreadsPerCore,omitempty" xml:"CpuOptionsThreadsPerCore,omitempty"`
	// The bucket that stores the data cache.
	//
	// example:
	//
	// default
	DataCacheBucket *string `json:"DataCacheBucket,omitempty" xml:"DataCacheBucket,omitempty"`
	// Specifies whether to enable the performance burst feature when ESSDs AutoPL are used to store data caches. For more information, see [ESSDs AutoPL](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	DataCacheBurstingEnabled *bool `json:"DataCacheBurstingEnabled,omitempty" xml:"DataCacheBurstingEnabled,omitempty"`
	// The performance level (PL) of the disk that you want to use to store data caches.\\
	//
	// Enhanced SSDs (ESSDs) are preferentially used to store data caches. The default performance level is PL1.
	//
	// example:
	//
	// PL1
	DataCachePL *string `json:"DataCachePL,omitempty" xml:"DataCachePL,omitempty"`
	// The input/output operations per second (IOPS) provisioned for ESSDs AutoPL when ESSDs AutoPL are used to store data caches.\\
	//
	// Valid values: 0 to min{50000, 1000 × Storage capacity - Baseline IOPS}. Baseline IOPS = min{1,800 + 50 × Storage capacity, 50,000}.\\
	//
	// For more information, see [ESSDs AutoPL](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	DataCacheProvisionedIops *int64 `json:"DataCacheProvisionedIops,omitempty" xml:"DataCacheProvisionedIops,omitempty"`
	// The Domain Name System (DNS) policy. Valid values:
	//
	// 	- None: uses the DNS that is specified by DnsConfig.
	//
	// 	- Default: uses the DNS that is specified for the runtime environment.
	//
	// example:
	//
	// Default
	DnsPolicy *string `json:"DnsPolicy,omitempty" xml:"DnsPolicy,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, service limits, and available resources. If the request passes the dry run, the DryRunOperation error code is returned. Otherwise, an error message is returned.
	//
	// 	- false (default): performs a dry run and performs the actual request. If the request passes the dry run, the elastic container instance is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The maximum outbound bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 1024000
	EgressBandwidth *int64 `json:"EgressBandwidth,omitempty" xml:"EgressBandwidth,omitempty"`
	// The maximum bandwidth value for the EIP. Unit: Mbit/s. Default value: 5.\\
	//
	// This parameter is valid only when AutoCreateEip is set to true.
	//
	// example:
	//
	// 5
	EipBandwidth *int32 `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	// The EIP bandwidth plan that you want to associate with the instance.
	//
	// example:
	//
	// cbwp-2zeukbj916scmj51m****
	EipCommonBandwidthPackage *string `json:"EipCommonBandwidthPackage,omitempty" xml:"EipCommonBandwidthPackage,omitempty"`
	// The line type of the EIP. Default value: BGP. Valid values:
	//
	// 	- BGP: BGP (Multi-ISP) line
	//
	// 	- BGP_PRO: BGP (Multi-ISP) Pro line
	//
	// example:
	//
	// BPG
	EipISP *string `json:"EipISP,omitempty" xml:"EipISP,omitempty"`
	// The ID of the elastic IP address (EIP).
	//
	// example:
	//
	// eip-uf66jeqopgqa9hdn****
	EipInstanceId *string `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	// The increased capacity of the temporary storage space. Unit: GiB.\\
	//
	// For more information, see [Increase the size of the temporary storage space](https://help.aliyun.com/document_detail/204066.html).
	//
	// example:
	//
	// 20
	EphemeralStorage *int32 `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	// Specifies whether to configure the instance to use a fixed IP address. For more information, see [Configure an elastic container instance to use a fixed IP address](https://help.aliyun.com/document_detail/2381086.html).
	//
	// example:
	//
	// true
	FixedIp *string `json:"FixedIp,omitempty" xml:"FixedIp,omitempty"`
	// The retention period of the fixed IP address after the original instance is released and the fixed IP address becomes idle. Unit: hours. Default value: 48.
	//
	// example:
	//
	// 24
	FixedIpRetainHour *int32 `json:"FixedIpRetainHour,omitempty" xml:"FixedIpRetainHour,omitempty"`
	// The version of the GPU driver. Default value: tesla=470.82.01. Valid values:
	//
	// 	- tesla=470.82.01
	//
	// 	- tesla=525.85.12
	//
	// >  You can switch the GPU driver version only for a few Elastic Compute Service (ECS) instance types. For more information, see [Specify GPU-accelerated ECS instance types to create an elastic container instance](https://help.aliyun.com/document_detail/2579486.html).
	//
	// example:
	//
	// tesla=525.85.12
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// The alias of the elastic container instance.
	HostAliase []*CreateContainerGroupRequestHostAliase `json:"HostAliase,omitempty" xml:"HostAliase,omitempty" type:"Repeated"`
	// The hostname.
	//
	// example:
	//
	// test
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The image acceleration mode. Valid values:
	//
	// 	- nydus: uses Nydus to accelerate image pulling. The images must support Nydus.
	//
	// 	- dadi: uses DADI to accelerate image pulling. The images must support DADI.
	//
	// 	- p2p: uses P2P to accelerate image pulling. The images must support p2p.
	//
	// 	- imc: uses image caches to accelerate image pulling.
	//
	// example:
	//
	// imc
	ImageAccelerateMode *string `json:"ImageAccelerateMode,omitempty" xml:"ImageAccelerateMode,omitempty"`
	// The information about the logon credentials.
	ImageRegistryCredential []*CreateContainerGroupRequestImageRegistryCredential `json:"ImageRegistryCredential,omitempty" xml:"ImageRegistryCredential,omitempty" type:"Repeated"`
	// The ID of the image cache. For more information, see [Use image caches to accelerate the creation of instances](https://help.aliyun.com/document_detail/141281.html).
	//
	// example:
	//
	// imc-2zebxkiifuyzzlhl****
	ImageSnapshotId *string `json:"ImageSnapshotId,omitempty" xml:"ImageSnapshotId,omitempty"`
	// The maximum inbound bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 1024000
	IngressBandwidth *int64 `json:"IngressBandwidth,omitempty" xml:"IngressBandwidth,omitempty"`
	// The information about the init containers.
	InitContainer []*CreateContainerGroupRequestInitContainer `json:"InitContainer,omitempty" xml:"InitContainer,omitempty" type:"Repeated"`
	// The address of the self-managed image repository. When you create an elastic container instance by using an image in a self-managed image repository that uses a self-signed certificate, you must specify this parameter to skip the certificate authentication. This prevents image pull failures caused by certificate authentication failures.
	//
	// example:
	//
	// "harbor***.pre.com,192.168.XX.XX:5000,reg***.test.com:80"
	InsecureRegistry *string `json:"InsecureRegistry,omitempty" xml:"InsecureRegistry,omitempty"`
	// The ECS instance types that you specify to create the elastic container instance. Multiple instance types are supported. For more information, see [Specify ECS instance types to create an elastic container instance](https://help.aliyun.com/document_detail/114664.html).
	//
	// example:
	//
	// ecs.c5.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of IPv6 addresses that are assigned to the instance. Set the value to 1. You can assign only one IPv6 address to an elastic container instance.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The maximum IPv6 Internet bandwidth when you set Ipv6GatewayBandwidthEnable to true. Valid values:
	//
	// 	- If the billing method for IPv6 network usage is pay-by-bandwidth, the maximum IPv6 Internet bandwidth ranges from 1 to 2,000 Mbit/s.
	//
	// 	- If the billing method for IPv6 network usage is pay-by-traffic, the maximum IPv6 Internet bandwidth varies based on the edition of the IPv6 gateway.
	//
	//     	- If the IPv6 gateway is of Free Edition, the maximum IPv6 Internet bandwidth ranges from 1 to 200 Mbit/s.
	//
	//     	- If the IPv6 gateway is of Enterprise Edition, the maximum IPv6 Internet bandwidth ranges from 1 to 500 Mbit/s.
	//
	//     	- If the IPv6 gateway is of Enhanced Enterprise Edition, the maximum IPv6 Internet bandwidth ranges from 1 to 1000 Mbit/s.
	//
	// The default value is the maximum value in the Internet bandwidth range of the IPv6 gateway.
	//
	// example:
	//
	// 100
	Ipv6GatewayBandwidth *string `json:"Ipv6GatewayBandwidth,omitempty" xml:"Ipv6GatewayBandwidth,omitempty"`
	// Specifies whether to enable Internet access to the elastic container instance over IPv6 addresses.
	//
	// example:
	//
	// true
	Ipv6GatewayBandwidthEnable *bool `json:"Ipv6GatewayBandwidthEnable,omitempty" xml:"Ipv6GatewayBandwidthEnable,omitempty"`
	// The memory size that you want to allocate to the instance. Unit: GiB.
	//
	// example:
	//
	// 2.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The endpoints of the Network Time Protocol (NTP) servers.
	//
	// example:
	//
	// ntp.cloud.aliyuncs.com
	NtpServer []*string `json:"NtpServer,omitempty" xml:"NtpServer,omitempty" type:"Repeated"`
	// The operating system of the elastic container instance. Default value: Linux. Valid values:
	//
	// 	- Linux
	//
	// 	- Windows
	//
	// >  Windows instances are in invitational preview. To use the operating system, submit a ticket.
	//
	// example:
	//
	// Windows
	OsType       *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The address of the self-managed image repository. When you create an elastic container instance by using an image in a self-managed image repository that uses the HTTP protocol, you must specify this parameter. This way, Elastic Container Instance pulls the image over the HTTP protocol instead of the default HTTPS protocol. This prevents image pull failures caused by different protocols.
	//
	// example:
	//
	// "harbor***.pre.com,192.168.XX.XX:5000,reg***.test.com:80"
	PlainHttpRegistry *string `json:"PlainHttpRegistry,omitempty" xml:"PlainHttpRegistry,omitempty"`
	// The private IP address of the elastic container instance. Only IPv4 addresses are supported. Make sure that the IP address is idle.
	//
	// example:
	//
	// 172.16.0.1
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The name of the instance Resource Access Management (RAM) role. You can use the same RAM role to access elastic container instances and ECS instances. For more information, see [Use an instance RAM role by calling API operations](https://help.aliyun.com/document_detail/61178.html).
	//
	// example:
	//
	// RamTestRole
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-uf66jeqopgqa9hdn****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The restart policy of the instance. Valid values:
	//
	// 	- Always: Always restarts the instance if a container in the instance exits upon termination.
	//
	// 	- Never: Never restarts the instance if a container in the instance exits upon termination.
	//
	// 	- OnFailure: Restarts the instance only if a container in the instance exists upon failure with a status code of non-zero.
	//
	// Default value: Always.
	//
	// example:
	//
	// Always
	RestartPolicy *string `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	// The resource scheduling policy when you specify multiple zones to create an elastic container instance. To specify multiple zones, you can use the VSwitchId to specify multiple vSwitches. Valid values:
	//
	// 	- VSwitchOrdered: The system schedules resources in the sequence of the vSwitches.
	//
	// 	- VSwitchRandom: The system schedules resources at random.
	//
	// For more information, see [Specify multiple zones to create an elastic container instance](https://help.aliyun.com/document_detail/157290.html).
	//
	// example:
	//
	// VSwitchOrdered
	ScheduleStrategy *string `json:"ScheduleStrategy,omitempty" xml:"ScheduleStrategy,omitempty"`
	// The ID of the security group to which the instance belongs. Instances in the same security group can access each other.
	//
	// If you do not specify a security group, the system automatically uses the default security group in the region that you selected. Make sure that the inbound rules of the security group contain the container protocols and port numbers that you want to expose. If you do not have a default security group in the region, the system creates a default security group, and then adds the container protocols and port numbers that you specified to the inbound rules of the security group.
	//
	// example:
	//
	// sg-uf66jeqopgqa9hdn****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// Specifies whether to use a shared namespace. Default value: false.
	//
	// example:
	//
	// false
	ShareProcessNamespace *bool `json:"ShareProcessNamespace,omitempty" xml:"ShareProcessNamespace,omitempty"`
	// The protection period of the preemptible elastic container instance. Unit: hours. Default value: 1. A value of 0 indicates no protection period.
	//
	// example:
	//
	// 1
	SpotDuration *int64 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The maximum hourly price of the preemptible elastic container instance. The value can be accurate to three decimal places.
	//
	// If you set SpotStrategy to SpotWithPriceLimit, you must specify the SpotPriceLimit parameter.
	//
	// example:
	//
	// 0.025
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bid policy for the instance. Valid values:
	//
	// 	- NoSpot: The instance is created as a pay-as-you-go instance.
	//
	// 	- SpotWithPriceLimit: The instance is created as a preemptible instance for which you specify the maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instance is created as a preemptible instance for which the market price at the time of purchase is automatically used as the bid price.
	//
	// Default value: NoSpot.
	//
	// example:
	//
	// SpotWithPriceLimit
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// Specifies whether to enable periodical execution.
	//
	// 	- true: enables periodical execution.
	//
	// 	- false: disables periodical execution.
	//
	// example:
	//
	// true
	StrictSpot *bool `json:"StrictSpot,omitempty" xml:"StrictSpot,omitempty"`
	// The tags that you want to add to the instance. You can bind a maximum of 20 tags. For more information, see [Use tags to manage elastic container instances](https://help.aliyun.com/document_detail/146608.html).
	Tag []*CreateContainerGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The buffer period of time during which the program handles operations before the program is stopped. Unit: seconds.
	//
	// example:
	//
	// 60
	TerminationGracePeriodSeconds *int64 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	// The IDs of the vSwitches that connect to the instance. You can specify up to 10 vSwitch IDs at a time. Separate multiple vSwitch IDs with commas (,). Example: `vsw-***,vsw-***`.
	//
	// If you do not specify a vSwitch, the system automatically uses the default vSwitch in the default VPC in the region that you selected. If you do not have a default VPC or a default vSwitch in the region, the system automatically creates a default VPC and a default vSwitch.
	//
	// >  The number of IP addresses in the vSwitch CIDR block determines the maximum number of elastic container instances that you can create for the vSwitch. Before you create elastic container instances, you must plan the CIDR block of the vSwitch.
	//
	// example:
	//
	// vsw-bp1xpiowfm5vo8o3c****,vsw-bp1rkyjgr1xwoho6k****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The information about the volume that you want to mount to the container.
	Volume []*CreateContainerGroupRequestVolume `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Repeated"`
	// The zone ID of the instance. If you do not specify this parameter, the system selects a zone.
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// cn-hangzhou-b
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

func (s *CreateContainerGroupRequest) SetComputeCategory(v []*string) *CreateContainerGroupRequest {
	s.ComputeCategory = v
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

func (s *CreateContainerGroupRequest) SetCpuArchitecture(v string) *CreateContainerGroupRequest {
	s.CpuArchitecture = &v
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

func (s *CreateContainerGroupRequest) SetDataCacheBucket(v string) *CreateContainerGroupRequest {
	s.DataCacheBucket = &v
	return s
}

func (s *CreateContainerGroupRequest) SetDataCacheBurstingEnabled(v bool) *CreateContainerGroupRequest {
	s.DataCacheBurstingEnabled = &v
	return s
}

func (s *CreateContainerGroupRequest) SetDataCachePL(v string) *CreateContainerGroupRequest {
	s.DataCachePL = &v
	return s
}

func (s *CreateContainerGroupRequest) SetDataCacheProvisionedIops(v int64) *CreateContainerGroupRequest {
	s.DataCacheProvisionedIops = &v
	return s
}

func (s *CreateContainerGroupRequest) SetDnsPolicy(v string) *CreateContainerGroupRequest {
	s.DnsPolicy = &v
	return s
}

func (s *CreateContainerGroupRequest) SetDryRun(v bool) *CreateContainerGroupRequest {
	s.DryRun = &v
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

func (s *CreateContainerGroupRequest) SetFixedIp(v string) *CreateContainerGroupRequest {
	s.FixedIp = &v
	return s
}

func (s *CreateContainerGroupRequest) SetFixedIpRetainHour(v int32) *CreateContainerGroupRequest {
	s.FixedIpRetainHour = &v
	return s
}

func (s *CreateContainerGroupRequest) SetGpuDriverVersion(v string) *CreateContainerGroupRequest {
	s.GpuDriverVersion = &v
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

func (s *CreateContainerGroupRequest) SetOsType(v string) *CreateContainerGroupRequest {
	s.OsType = &v
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

func (s *CreateContainerGroupRequest) SetPrivateIpAddress(v string) *CreateContainerGroupRequest {
	s.PrivateIpAddress = &v
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
	// The IP addresses of DNS servers.
	//
	// example:
	//
	// 172.10.*.**
	NameServer []*string `json:"NameServer,omitempty" xml:"NameServer,omitempty" type:"Repeated"`
	// Configuration options of the DNS server.
	Option []*CreateContainerGroupRequestDnsConfigOption `json:"Option,omitempty" xml:"Option,omitempty" type:"Repeated"`
	// The search domains of DNS servers.
	//
	// example:
	//
	// svc.local.kubenetes
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
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the option.
	//
	// example:
	//
	// value
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
	// Configure a security context to modify unsafe sysctls. For more information, see [Configure a security context](https://help.aliyun.com/document_detail/462313.html).
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
	// The name of the unsafe sysctl when you modify sysctls by configuring a security context. Valid values:
	//
	// 	- kernel.shm \\	- (except kernel.shm_rmid_forced)
	//
	// 	- kernel.msg\\*
	//
	// 	- kernel.sem
	//
	// 	- fs.mqueue.\\*
	//
	// 	- net.\\*(except net.ipv4.tcp_syncookies, net.ipv4.ping_group_range, and net.ipv4.ip_unprivileged_port_start)
	//
	// example:
	//
	// kernel.msgmax
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the unsafe sysctl when you modify sysctls by configuring a security context.
	//
	// example:
	//
	// 65536
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
	// Configure a security context to modify safe sysctls. For more information, see [Configure a security context](https://help.aliyun.com/document_detail/462313.html).
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
	// The name of the safe sysctl when you modify sysctls by configuring a security context. Valid values:
	//
	// 	- net.ipv4.ping_group_range
	//
	// 	- net.ipv4.ip_unprivileged_port_start
	//
	// example:
	//
	// kernel.msgmax
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the safe sysctl when you modify sysctls by configuring a security context.
	//
	// example:
	//
	// 65536
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
	// The Alibaba Cloud Resource Name (ARN) of the RAM role in the Alibaba Cloud account to which the elastic container instance belongs.
	//
	// example:
	//
	// acs:ram::1609982529******:role/role-assume
	ArnService *string `json:"ArnService,omitempty" xml:"ArnService,omitempty"`
	// The ARN of the RAM role in the Alibaba Cloud account to which the Container Registry Enterprise Edition instance belongs.
	//
	// example:
	//
	// acs:ram::1298452580******:role/role-acr
	ArnUser *string `json:"ArnUser,omitempty" xml:"ArnUser,omitempty"`
	// The domain names of the Container Registry Enterprise Edition instance. By default, all domain names of the instance are displayed. You can specify multiple domain names. Separate multiple domain names with commas (,).
	//
	// example:
	//
	// *****-****-registry.cn-beijing.cr.aliyuncs.com
	Domain []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
	// The ID of the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// cri-nwj395hgf6f3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// acr-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID of the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateContainerGroupRequestAcrRegistryInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestAcrRegistryInfo) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestAcrRegistryInfo) SetArnService(v string) *CreateContainerGroupRequestAcrRegistryInfo {
	s.ArnService = &v
	return s
}

func (s *CreateContainerGroupRequestAcrRegistryInfo) SetArnUser(v string) *CreateContainerGroupRequestAcrRegistryInfo {
	s.ArnUser = &v
	return s
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
	LivenessProbe   *CreateContainerGroupRequestContainerLivenessProbe   `json:"LivenessProbe,omitempty" xml:"LivenessProbe,omitempty" type:"Struct"`
	ReadinessProbe  *CreateContainerGroupRequestContainerReadinessProbe  `json:"ReadinessProbe,omitempty" xml:"ReadinessProbe,omitempty" type:"Struct"`
	SecurityContext *CreateContainerGroupRequestContainerSecurityContext `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	// The arguments that are passed to the startup command of the container. You can specify up to 10 arguments.
	//
	// example:
	//
	// 100
	Arg []*string `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	// The commands to be executed in the container when you use a CLI to perform health checks.
	//
	// >  When you configure ReadinessProbe-related parameters, you can select only one of the HttpGet, Exec, and TcpSocket check methods.
	//
	// example:
	//
	// sleep
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// The number of vCPUs that you want to allocate to the container.
	//
	// example:
	//
	// 0.25
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The environment variables of the container.
	EnvironmentVar []*CreateContainerGroupRequestContainerEnvironmentVar `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	// Specifies whether to hide the information about environment variables when you query the details of an elastic container instance. Default value: false. Valid values:
	//
	// 	- false
	//
	// 	- true If environment variables contain sensitive information, you can set this parameter to true to improve security of the information.
	//
	// example:
	//
	// false
	EnvironmentVarHide *bool `json:"EnvironmentVarHide,omitempty" xml:"EnvironmentVarHide,omitempty"`
	// The number of GPUs that you want to allocate to the container.
	//
	// example:
	//
	// 1
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The image of the container.
	//
	// This parameter is required.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/eci_open/nginx:latest
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The policy that you want to use to pull images. Valid values:
	//
	// 	- Always: Each time instances are created, image pulling is performed.
	//
	// 	- IfNotPresent: On-premises images are preferentially used. If no on-premises images are available, image pulling is performed.
	//
	// 	- Never: On-premises images are always used. Image pulling is not performed.
	//
	// example:
	//
	// Always
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The commands to be executed in containers when you use a CLI to specify a postStart hook.
	//
	// example:
	//
	// ["/bin/sh", "-c", "echo Hello from the postStart handler > /usr/share/message"]
	LifecyclePostStartHandlerExec []*string `json:"LifecyclePostStartHandlerExec,omitempty" xml:"LifecyclePostStartHandlerExec,omitempty" type:"Repeated"`
	// The IP address of the host that receives the HTTP GET request when you use an HTTP request to specify a postStart hook.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePostStartHandlerHttpGetHost *string `json:"LifecyclePostStartHandlerHttpGetHost,omitempty" xml:"LifecyclePostStartHandlerHttpGetHost,omitempty"`
	// The HTTP GET request header.
	LifecyclePostStartHandlerHttpGetHttpHeader []*CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader `json:"LifecyclePostStartHandlerHttpGetHttpHeader,omitempty" xml:"LifecyclePostStartHandlerHttpGetHttpHeader,omitempty" type:"Repeated"`
	// The path to which the system sends an HTTP GET request for a health check when you use an HTTP request to specify a postStart hook.
	//
	// example:
	//
	// /healthyz
	LifecyclePostStartHandlerHttpGetPath *string `json:"LifecyclePostStartHandlerHttpGetPath,omitempty" xml:"LifecyclePostStartHandlerHttpGetPath,omitempty"`
	// The port to which the system sends an HTTP GET request when you use an HTTP request to specify a postStart hook.
	//
	// example:
	//
	// 5050
	LifecyclePostStartHandlerHttpGetPort *int32 `json:"LifecyclePostStartHandlerHttpGetPort,omitempty" xml:"LifecyclePostStartHandlerHttpGetPort,omitempty"`
	// The protocol type of HTTP GET requests when you use HTTP requests to specify a postStart hook. Valid values:
	//
	// 	- HTTP
	//
	// 	- HTTPS
	//
	// example:
	//
	// HTTPS
	LifecyclePostStartHandlerHttpGetScheme *string `json:"LifecyclePostStartHandlerHttpGetScheme,omitempty" xml:"LifecyclePostStartHandlerHttpGetScheme,omitempty"`
	// The IP address of the host that receives the TCP socket request when you use a TCP socket request to specify a postStart hook.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePostStartHandlerTcpSocketHost *string `json:"LifecyclePostStartHandlerTcpSocketHost,omitempty" xml:"LifecyclePostStartHandlerTcpSocketHost,omitempty"`
	// The port to which the system sends a TCP socket request for a health check when you use TCP sockets to specify a postStart hook.
	//
	// example:
	//
	// 80
	LifecyclePostStartHandlerTcpSocketPort *int32 `json:"LifecyclePostStartHandlerTcpSocketPort,omitempty" xml:"LifecyclePostStartHandlerTcpSocketPort,omitempty"`
	// The commands to be executed in containers when you use a CLI to specify a preStop hook.
	//
	// example:
	//
	// ["/bin/sh", "-c","echo Hello from the preStop handler > /usr/share/message"]
	LifecyclePreStopHandlerExec []*string `json:"LifecyclePreStopHandlerExec,omitempty" xml:"LifecyclePreStopHandlerExec,omitempty" type:"Repeated"`
	// The IP address of the host that receives the HTTP GET request when you use an HTTP request to specify a preStop hook.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePreStopHandlerHttpGetHost *string `json:"LifecyclePreStopHandlerHttpGetHost,omitempty" xml:"LifecyclePreStopHandlerHttpGetHost,omitempty"`
	// The HTTP GET request header.
	LifecyclePreStopHandlerHttpGetHttpHeader []*CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader `json:"LifecyclePreStopHandlerHttpGetHttpHeader,omitempty" xml:"LifecyclePreStopHandlerHttpGetHttpHeader,omitempty" type:"Repeated"`
	// The path to which the system sends an HTTP GET request for a health check when you use an HTTP request to specify a preSop hook.
	//
	// example:
	//
	// /healthyz
	LifecyclePreStopHandlerHttpGetPath *string `json:"LifecyclePreStopHandlerHttpGetPath,omitempty" xml:"LifecyclePreStopHandlerHttpGetPath,omitempty"`
	// The port to which the system sends an HTTP GET request for a health check when you use HTTP requests to specify a preStop hook.
	//
	// example:
	//
	// 88
	LifecyclePreStopHandlerHttpGetPort *int32 `json:"LifecyclePreStopHandlerHttpGetPort,omitempty" xml:"LifecyclePreStopHandlerHttpGetPort,omitempty"`
	// The protocol type of the HTTP GET request when you use an HTTP request to specify a preStop hook. Valid values:
	//
	// 	- HTTP
	//
	// 	- HTTPS
	//
	// example:
	//
	// HTTP
	LifecyclePreStopHandlerHttpGetScheme *string `json:"LifecyclePreStopHandlerHttpGetScheme,omitempty" xml:"LifecyclePreStopHandlerHttpGetScheme,omitempty"`
	// The IP address of the host that receives the TCP socket request when you use a TCP socket request to specify a preStop hook.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePreStopHandlerTcpSocketHost *string `json:"LifecyclePreStopHandlerTcpSocketHost,omitempty" xml:"LifecyclePreStopHandlerTcpSocketHost,omitempty"`
	// The port to which the system sends a TCP socket request for a health check when you use TCP sockets to specify a preStop hook.
	//
	// example:
	//
	// 90
	LifecyclePreStopHandlerTcpSocketPort *int32 `json:"LifecyclePreStopHandlerTcpSocketPort,omitempty" xml:"LifecyclePreStopHandlerTcpSocketPort,omitempty"`
	// The memory size of the container. Unit: GiB
	//
	// example:
	//
	// 0.5
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the container.
	//
	// This parameter is required.
	//
	// example:
	//
	// nginx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port to which the system sends an HTTP GET request for a health check when you use HTTP requests to perform health checks.
	//
	// >  When you configure LivenessProbe-related parameters, you can select only one of the HttpGet, Exec, and TcpSocket check methods.
	Port []*CreateContainerGroupRequestContainerPort `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
	// Specifies whether to enable privileged mode for the container. That is, whether to run the container in privileged mode. Default value: false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	SecurityContextPrivileged *bool `json:"SecurityContextPrivileged,omitempty" xml:"SecurityContextPrivileged,omitempty"`
	// The user group that runs the container.
	//
	// example:
	//
	// 3000
	SecurityContextRunAsGroup *int64 `json:"SecurityContextRunAsGroup,omitempty" xml:"SecurityContextRunAsGroup,omitempty"`
	// Specifies whether to run the container as a non-root user. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	SecurityContextRunAsNonRoot *bool `json:"SecurityContextRunAsNonRoot,omitempty" xml:"SecurityContextRunAsNonRoot,omitempty"`
	// Specifies whether the container allocates buffer resources to standard input streams when the container is running. If you do not specify this parameter, an end-of-file (EOF) error may occur when standard input streams in the container are read. Default value: false.
	//
	// example:
	//
	// false
	Stdin *bool `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	// Specifies whether standard input streams are disconnected from multiple sessions after a client is disconnected.\\
	//
	// If StdinOnce is set to true, standard input streams are connected after the container is started, and remain idle until a client is connected to receive data. After the client is disconnected, standard input streams are also disconnected, and remain disconnected until the container restarts.
	//
	// example:
	//
	// false
	StdinOnce *bool `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	// The path of the file from which the system retrieves termination messages of the container when the container exits.
	//
	// example:
	//
	// /tmp/termination-log
	TerminationMessagePath *string `json:"TerminationMessagePath,omitempty" xml:"TerminationMessagePath,omitempty"`
	// The message notification policy. This parameter is empty by default. Only Message Service (MNS) queue message notifications can be sent.
	//
	// example:
	//
	// FallbackToLogsOnError
	TerminationMessagePolicy *string `json:"TerminationMessagePolicy,omitempty" xml:"TerminationMessagePolicy,omitempty"`
	// Specifies whether to enable interaction. Default value: false.
	//
	// If the command is a /bin/bash command, set the value to true.
	//
	// example:
	//
	// false
	Tty *bool `json:"Tty,omitempty" xml:"Tty,omitempty"`
	// The information about the volume that you want to mount on the container.
	VolumeMount []*CreateContainerGroupRequestContainerVolumeMount `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
	// The working directory of the container.
	//
	// example:
	//
	// /usr/local/
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

func (s *CreateContainerGroupRequestContainer) SetEnvironmentVarHide(v bool) *CreateContainerGroupRequestContainer {
	s.EnvironmentVarHide = &v
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

func (s *CreateContainerGroupRequestContainer) SetSecurityContextPrivileged(v bool) *CreateContainerGroupRequestContainer {
	s.SecurityContextPrivileged = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetSecurityContextRunAsGroup(v int64) *CreateContainerGroupRequestContainer {
	s.SecurityContextRunAsGroup = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetSecurityContextRunAsNonRoot(v bool) *CreateContainerGroupRequestContainer {
	s.SecurityContextRunAsNonRoot = &v
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
	Exec                *CreateContainerGroupRequestContainerLivenessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" type:"Struct"`
	FailureThreshold    *int32                                                      `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *CreateContainerGroupRequestContainerLivenessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	InitialDelaySeconds *int32                                                      `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                      `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                      `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *CreateContainerGroupRequestContainerLivenessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
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
	Exec                *CreateContainerGroupRequestContainerReadinessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" type:"Struct"`
	FailureThreshold    *int32                                                       `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *CreateContainerGroupRequestContainerReadinessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	InitialDelaySeconds *int32                                                       `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                       `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                       `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *CreateContainerGroupRequestContainerReadinessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
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
	Capability             *CreateContainerGroupRequestContainerSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" type:"Struct"`
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
	FieldRef *CreateContainerGroupRequestContainerEnvironmentVarFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" type:"Struct"`
	// The name of the environment variable. The name must be 1 to 128 characters in length and can contain letters, digits, and underscores (_). It cannot start with a digit.``
	//
	// example:
	//
	// PATH
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable. The value can be up to 256 characters in length.
	//
	// example:
	//
	// /usr/local/bin
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
	// The name of the custom field in the HTTP GET request header when you use HTTP requests to specify a postStart hook.
	//
	// example:
	//
	// Xiao-Custom-Header
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the custom field in the HTTP GET request header when you use HTTP requests to specify a postStart hook.
	//
	// example:
	//
	// test-postStart
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
	// The name of the custom field in the HTTP GET request header when you use HTTP requests to specify a presto hook.
	//
	// example:
	//
	// Xiao-Custom-Header
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the custom field in the HTTP GET request header when you use HTTP requests to specify a preStop hook.
	//
	// example:
	//
	// test-preStop
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
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// example:
	//
	// TCP
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
	// >  The data stored in this directory is overwritten by the data on the volume. Specify this parameter with caution.
	//
	// example:
	//
	// /pod/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation settings of the volume. Mount propagation allows volumes that are mounted on one container to be shared with other containers in the same pod, or even with other pods on the same node. Valid values:
	//
	// 	- None: The volume mount does not receive subsequent mounts that are performed on this volume or subdirectories of this volume.
	//
	// 	- HostToCotainer: The volume mount receives subsequent mounts that are performed on this volume or the subdirectories of this volume.
	//
	// 	- Bidirectional: This value is similar to HostToContainer. The volume mount receives subsequent mounts that are performed on this volume or the subdirectories of this volume. In addition, all volume mounts that are mounted on the container are propagated back to the host and all containers of all pods that use the same volume.
	//
	// Default value: None.
	//
	// example:
	//
	// None
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The name of the volume. The value of this parameter is the same as the name of the volume that is mounted to containers.
	//
	// example:
	//
	// default-volume1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the volume is read-only. Default value: false.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The subdirectory of the volume.
	//
	// example:
	//
	// data2/
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
	// The hostnames of the elastic container instance.
	//
	// example:
	//
	// hehe.com
	Hostname []*string `json:"Hostname,omitempty" xml:"Hostname,omitempty" type:"Repeated"`
	// The IP address of the elastic container instance.
	//
	// example:
	//
	// 1.1.1.1
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
	// The password that you use to access the image repository.
	//
	// example:
	//
	// yourpassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The address of the image repository.
	//
	// example:
	//
	// registry-vpc.cn-shanghai.aliyuncs.com
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The username that you use to access the image repository.
	//
	// example:
	//
	// yourusername
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
	SecurityContext *CreateContainerGroupRequestInitContainerSecurityContext `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	// The arguments that are passed to the startup command of the init container.
	//
	// example:
	//
	// 10
	Arg []*string `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	// The startup commands of the init container.
	//
	// example:
	//
	// sleep
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// The number of vCPUs that you want to allocate to the init container.
	//
	// example:
	//
	// 0.5
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The environment variable of the init container.
	EnvironmentVar []*CreateContainerGroupRequestInitContainerEnvironmentVar `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	// The number of GPUs that you want to allocate to the init container.
	//
	// example:
	//
	// 1
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The image of the init container.
	//
	// example:
	//
	// nginx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The policy that you want to use to pull images. Valid values:
	//
	// 	- Always: Each time instances are created, image pulling is performed.
	//
	// 	- IfNotPresent: On-premises images are preferentially used. If no on-premises images are available, image pulling is performed.
	//
	// 	- Never: On-premises images are always used. Image pulling is not performed.
	//
	// example:
	//
	// Always
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The memory size that you want to allocate to the init container. Unit: GiB.
	//
	// example:
	//
	// 1.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the init container.
	//
	// example:
	//
	// test-init
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The information about the port.
	Port []*CreateContainerGroupRequestInitContainerPort `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
	// The path of the file from which the system retrieves termination messages of the init container when the init container exits.
	//
	// example:
	//
	// /tmp/termination-log
	TerminationMessagePath *string `json:"TerminationMessagePath,omitempty" xml:"TerminationMessagePath,omitempty"`
	// The message notification policy. This parameter is empty by default.
	//
	// example:
	//
	// *****
	TerminationMessagePolicy *string `json:"TerminationMessagePolicy,omitempty" xml:"TerminationMessagePolicy,omitempty"`
	// The information about the volumes that you want to mount to the init containers.
	VolumeMount []*CreateContainerGroupRequestInitContainerVolumeMount `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
	// The working directory of the init container.
	//
	// example:
	//
	// /usr/local
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
	Capability             *CreateContainerGroupRequestInitContainerSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" type:"Struct"`
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
	FieldRef *CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" type:"Struct"`
	// The name of the environment variable. The name must be 1 to 128 bits in length and can contain letters, digits, and underscores (_). It cannot start with a digit.``
	//
	// example:
	//
	// Path
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable. The value must be 0 to 256 bits in length.
	//
	// example:
	//
	// /usr/bin/
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
	//
	// example:
	//
	// 8888
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// example:
	//
	// TCP
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
	// The directory to which the volume is mounted. The data stored in this directory is overwritten by the data on the volume. Specify this parameter with caution.
	//
	// example:
	//
	// /usr/share/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation settings of the volume. Mount propagation allows volumes that are mounted on one init container to be shared with other init containers in the same pod, or even with other pods on the same node. Valid values:
	//
	// 	- None: The volume mount does not receive subsequent mounts that are performed on this volume or subdirectories of this volume.
	//
	// 	- HostToCotainer: The volume mount receives subsequent mounts that are performed on this volume or the subdirectories of this volume.
	//
	// 	- Bidirectional: This value is similar to HostToContainer. The volume mount receives subsequent mounts that are performed on this volume or the subdirectories of this volume. In addition, all volume mounts that are mounted on the init container are propagated back to the host and all init containers of all pods that use the same volume.
	//
	// Default value: None.
	//
	// example:
	//
	// None
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The name of the volume.
	//
	// example:
	//
	// test-empty
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the mount path is read-only. Default value: false.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The subdirectory of the volume. The pod can mount different directories of the same volume to different subdirectories of init containers.
	//
	// example:
	//
	// /usr/sub/
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
	// The tag key. The tag key cannot be an empty string and must be unique. The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// version
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be an empty string. The tag value can be up to 128 characters in length. It cannot start with `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// 3
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
	ConfigFileVolume *CreateContainerGroupRequestVolumeConfigFileVolume `json:"ConfigFileVolume,omitempty" xml:"ConfigFileVolume,omitempty" type:"Struct"`
	DiskVolume       *CreateContainerGroupRequestVolumeDiskVolume       `json:"DiskVolume,omitempty" xml:"DiskVolume,omitempty" type:"Struct"`
	EmptyDirVolume   *CreateContainerGroupRequestVolumeEmptyDirVolume   `json:"EmptyDirVolume,omitempty" xml:"EmptyDirVolume,omitempty" type:"Struct"`
	FlexVolume       *CreateContainerGroupRequestVolumeFlexVolume       `json:"FlexVolume,omitempty" xml:"FlexVolume,omitempty" type:"Struct"`
	HostPathVolume   *CreateContainerGroupRequestVolumeHostPathVolume   `json:"HostPathVolume,omitempty" xml:"HostPathVolume,omitempty" type:"Struct"`
	NFSVolume        *CreateContainerGroupRequestVolumeNFSVolume        `json:"NFSVolume,omitempty" xml:"NFSVolume,omitempty" type:"Struct"`
	// The name of the volume.
	//
	// example:
	//
	// default-volume1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the HostPath volume. Valid values:
	//
	// 	- Directory
	//
	// 	- File
	//
	// >  Only users in the whitelist can mount HostPath volumes.
	//
	// example:
	//
	// ConfigFileVolume
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
	// The ID of the instance.
	//
	// example:
	//
	// eci-uf6fonnghi50u374****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 89945DD3-9072-47D0-A318-353284CFC7B3
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContainerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateDataCacheRequest struct {
	// The bucket in which the data is stored. By default, the default bucket is used. You can use a custom bucket for business grouping and to prevent path conflicts.
	//
	// >  eci-system is the reserved bucket of the ECI and cannot be used.
	//
	// example:
	//
	// default
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-****-12d3-****-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The data source.
	DataSource *CreateDataCacheRequestDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The elastic IP address (EIP) to be created and associated. If no NAT gateway is configured for the virtual private cloud (VPC), you can associate an EIP to pull data from the Internet.
	EipCreateParam *CreateDataCacheRequestEipCreateParam `json:"EipCreateParam,omitempty" xml:"EipCreateParam,omitempty" type:"Struct"`
	// The existing elastic IP address (EIP) to be associated. If no NAT gateway is configured for the virtual private cloud (VPC), you can associate an EIP to pull data from the Internet.
	//
	// example:
	//
	// eip-uf66jeqopgqa9hdn****
	EipInstanceId *string `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	// The DataCache name.
	//
	// example:
	//
	// test
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The storage path of the data.
	//
	// example:
	//
	// /model/test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-2df3isufhi38****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The number of days for which the DataCache is retained. When the retention days end, the DataCache is deleted. By default, DataCaches do not expire.
	//
	// example:
	//
	// 7
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The ID of the security group to which the generated ECI belongs during the creation of the data cache.
	//
	// example:
	//
	// sg-2ze81zoc3yl7a3we****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The size of the data cache. Unit: GiB. Default value: 20. Evaluate the required size based on the actual data size.
	//
	// example:
	//
	// 30
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The tags to be bound to the data cache.
	Tag []*CreateDataCacheRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which the generated ECI belongs during the creation of the data cache.
	//
	// example:
	//
	// vsw-bp1jrgfqqy54kg5hc****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateDataCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCacheRequest) GoString() string {
	return s.String()
}

func (s *CreateDataCacheRequest) SetBucket(v string) *CreateDataCacheRequest {
	s.Bucket = &v
	return s
}

func (s *CreateDataCacheRequest) SetClientToken(v string) *CreateDataCacheRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDataCacheRequest) SetDataSource(v *CreateDataCacheRequestDataSource) *CreateDataCacheRequest {
	s.DataSource = v
	return s
}

func (s *CreateDataCacheRequest) SetEipCreateParam(v *CreateDataCacheRequestEipCreateParam) *CreateDataCacheRequest {
	s.EipCreateParam = v
	return s
}

func (s *CreateDataCacheRequest) SetEipInstanceId(v string) *CreateDataCacheRequest {
	s.EipInstanceId = &v
	return s
}

func (s *CreateDataCacheRequest) SetName(v string) *CreateDataCacheRequest {
	s.Name = &v
	return s
}

func (s *CreateDataCacheRequest) SetOwnerAccount(v string) *CreateDataCacheRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDataCacheRequest) SetOwnerId(v int64) *CreateDataCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDataCacheRequest) SetPath(v string) *CreateDataCacheRequest {
	s.Path = &v
	return s
}

func (s *CreateDataCacheRequest) SetRegionId(v string) *CreateDataCacheRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDataCacheRequest) SetResourceGroupId(v string) *CreateDataCacheRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDataCacheRequest) SetResourceOwnerAccount(v string) *CreateDataCacheRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDataCacheRequest) SetResourceOwnerId(v int64) *CreateDataCacheRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDataCacheRequest) SetRetentionDays(v int32) *CreateDataCacheRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateDataCacheRequest) SetSecurityGroupId(v string) *CreateDataCacheRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateDataCacheRequest) SetSize(v int32) *CreateDataCacheRequest {
	s.Size = &v
	return s
}

func (s *CreateDataCacheRequest) SetTag(v []*CreateDataCacheRequestTag) *CreateDataCacheRequest {
	s.Tag = v
	return s
}

func (s *CreateDataCacheRequest) SetVSwitchId(v string) *CreateDataCacheRequest {
	s.VSwitchId = &v
	return s
}

type CreateDataCacheRequestDataSource struct {
	// The parameters that are configured for the data source.
	Options map[string]*string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- URL
	//
	// 	- NAS
	//
	// 	- OSS
	//
	// 	- SNAPSHOT
	//
	// example:
	//
	// NAS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataCacheRequestDataSource) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCacheRequestDataSource) GoString() string {
	return s.String()
}

func (s *CreateDataCacheRequestDataSource) SetOptions(v map[string]*string) *CreateDataCacheRequestDataSource {
	s.Options = v
	return s
}

func (s *CreateDataCacheRequestDataSource) SetType(v string) *CreateDataCacheRequestDataSource {
	s.Type = &v
	return s
}

type CreateDataCacheRequestEipCreateParam struct {
	// The bandwidth of the EIP. Unit: Mbit/s. Default value: 5.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The EIP bandwidth plan to be associated.
	//
	// example:
	//
	// cbwp-2zeukbj916scmj51m****
	CommonBandwidthPackage *string `json:"CommonBandwidthPackage,omitempty" xml:"CommonBandwidthPackage,omitempty"`
	// The line type of the EIP. Valid values:
	//
	// 	- BGP: BGP (Multi-ISP) line
	//
	// 	- BGP_PRO: BGP (Multi-ISP) Pro line
	//
	// example:
	//
	// BGP
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The metering method of the EIP. Valid values:
	//
	// 	- PayByBandwidth: pay-by-bandwidth
	//
	// 	- PayByTraffic: pay-by-data-transfer
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The ID of the IP address pool. The EIP is allocated from the IP address pool. You cannot use the IP address pool feature by default. To use this feature, you must apply for the privilege in the Quota Center console.
	//
	// example:
	//
	// pippool-2vc4xdtjhwl4t0go5****
	PublicIpAddressPoolId *string `json:"PublicIpAddressPoolId,omitempty" xml:"PublicIpAddressPoolId,omitempty"`
}

func (s CreateDataCacheRequestEipCreateParam) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCacheRequestEipCreateParam) GoString() string {
	return s.String()
}

func (s *CreateDataCacheRequestEipCreateParam) SetBandwidth(v int32) *CreateDataCacheRequestEipCreateParam {
	s.Bandwidth = &v
	return s
}

func (s *CreateDataCacheRequestEipCreateParam) SetCommonBandwidthPackage(v string) *CreateDataCacheRequestEipCreateParam {
	s.CommonBandwidthPackage = &v
	return s
}

func (s *CreateDataCacheRequestEipCreateParam) SetISP(v string) *CreateDataCacheRequestEipCreateParam {
	s.ISP = &v
	return s
}

func (s *CreateDataCacheRequestEipCreateParam) SetInternetChargeType(v string) *CreateDataCacheRequestEipCreateParam {
	s.InternetChargeType = &v
	return s
}

func (s *CreateDataCacheRequestEipCreateParam) SetPublicIpAddressPoolId(v string) *CreateDataCacheRequestEipCreateParam {
	s.PublicIpAddressPoolId = &v
	return s
}

type CreateDataCacheRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// testkey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// testvalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDataCacheRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCacheRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDataCacheRequestTag) SetKey(v string) *CreateDataCacheRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDataCacheRequestTag) SetValue(v string) *CreateDataCacheRequestTag {
	s.Value = &v
	return s
}

type CreateDataCacheResponseBody struct {
	// The DataCache ID.
	//
	// example:
	//
	// edc-bp15l4vvys94oo******
	DataCacheId *string `json:"DataCacheId,omitempty" xml:"DataCacheId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D81A4A13-6DCC-4579-AC62-90A6C3EC7BBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCacheResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataCacheResponseBody) SetDataCacheId(v string) *CreateDataCacheResponseBody {
	s.DataCacheId = &v
	return s
}

func (s *CreateDataCacheResponseBody) SetRequestId(v string) *CreateDataCacheResponseBody {
	s.RequestId = &v
	return s
}

type CreateDataCacheResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCacheResponse) GoString() string {
	return s.String()
}

func (s *CreateDataCacheResponse) SetHeaders(v map[string]*string) *CreateDataCacheResponse {
	s.Headers = v
	return s
}

func (s *CreateDataCacheResponse) SetStatusCode(v int32) *CreateDataCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataCacheResponse) SetBody(v *CreateDataCacheResponseBody) *CreateDataCacheResponse {
	s.Body = v
	return s
}

type CreateImageCacheRequest struct {
	// Information about the Container Registry Enterprise Edition instance. For more information, see [Pull images from a Container Registry Enterprise Edition instance without using secrets](https://help.aliyun.com/document_detail/194250.html).
	AcrRegistryInfo []*CreateImageCacheRequestAcrRegistryInfo `json:"AcrRegistryInfo,omitempty" xml:"AcrRegistryInfo,omitempty" type:"Repeated"`
	// Comments.
	//
	// example:
	//
	// hide
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// Specifies whether to enable reuse of image cache layers. If you enable this feature, and the image cache that you want to create and an existing image cache contain duplicate image layers, the system reuses the duplicate image layers to create the new image cache. This accelerates the creation of the image cache. Valid values:
	//
	// 	- true: enables reuse of image cache layers.
	//
	// 	- false: disables reuse of image cache layers.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	AutoMatchImageCache *bool `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure the idempotence of a request](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-xxx-xxx-xxxx-42665544xxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the elastic IP address (EIP). If you want to pull images over the Internet, make sure that the elastic container instance can access the Internet. You can configure an EIP or a NAT gateway for the elastic container instance to access the Internet.
	//
	// example:
	//
	// eip-2zedsm5mfl3uhdj2d****
	EipInstanceId *string `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	// The elimination policy of the image cache. This parameter is empty by default, which indicates that the image cache is always retained.
	//
	// You can set this parameter to LRU, which indicates that the image cache can be automatically deleted. When the number of image caches reaches the quota, the system automatically deletes the image caches whose EliminationStrategy parameter is set to LRU and that are least commonly used.
	//
	// example:
	//
	// LRU
	EliminationStrategy *string `json:"EliminationStrategy,omitempty" xml:"EliminationStrategy,omitempty"`
	// Specifies whether to enable the instant image cache feature. The feature can accelerate the creation of image caches. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// true
	Flash *bool `json:"Flash,omitempty" xml:"Flash,omitempty"`
	// The number of temporary local snapshots. By default, the system creates one snapshot for each image cache. If an image cache is used to create multiple elastic container instances at a time, we recommend that you set this parameter to create multiple snapshots for the image cache. We recommend that you create one snapshot for creation of every 1,000 elastic container instances.
	//
	// >  If you set the Flash parameter to true, instant image cache is enabled. During the creation of the image cache, the system first creates a temporary local snapshot for you to instantly use the snapshot. After the temporary local snapshot is created, the system begins to create a regular snapshot. After the regular snapshot is created, the temporary local snapshot is automatically deleted.
	//
	// example:
	//
	// 7
	FlashCopyCount *int32 `json:"FlashCopyCount,omitempty" xml:"FlashCopyCount,omitempty"`
	// Container image N that is used to create the image cache.
	//
	// This parameter is required.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/eci_open/nginx:1.15.10-perl
	Image []*string `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
	// The name of the image cache.
	//
	// This parameter is required.
	//
	// example:
	//
	// testcache
	ImageCacheName *string `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	// The size of the image cache. Unit: GiB. Default value: 20.
	//
	// example:
	//
	// 20
	ImageCacheSize *int32 `json:"ImageCacheSize,omitempty" xml:"ImageCacheSize,omitempty"`
	// The image repository.
	ImageRegistryCredential []*CreateImageCacheRequestImageRegistryCredential `json:"ImageRegistryCredential,omitempty" xml:"ImageRegistryCredential,omitempty" type:"Repeated"`
	// The address of the self-managed image repository.
	//
	// When you create an image cache by using an image in a self-managed image repository that uses a self-signed certificate, you must specify this parameter to skip the certificate authentication. This can prevent the image from failing to pull due to certificate authentication failures.
	//
	// example:
	//
	// "harbor***.pre.com,192.168.XX.XX:5000,reg***.test.com:80"
	InsecureRegistry *string `json:"InsecureRegistry,omitempty" xml:"InsecureRegistry,omitempty"`
	// The operating system of the image. Default value: Linux. Valid values:
	//
	// - Linux
	//
	// - Windows
	//
	// > Windows instances are in invitational preview. To use the operating system, submit a ticket.
	//
	// example:
	//
	// Linux
	OsType       *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The address of the self-managed image repository. When you create an image cache by using an image in a self-managed image repository that uses the HTTP protocol, you must specify this parameter. This way, Elastic Container Instance uses the HTTP protocol instead of the default HTTPS protocol to pull the image. This can prevent the image from failing to pull due to different protocols.
	//
	// example:
	//
	// "harbor***.pre.com,192.168.XX.XX:5000,reg***.test.com:80"
	PlainHttpRegistry *string `json:"PlainHttpRegistry,omitempty" xml:"PlainHttpRegistry,omitempty"`
	// The region ID of the image cache.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzh43v*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The retention period of the image cache. Unit: days. When the retention period ends, the image cache expires and is deleted. By default, image caches never expire.
	//
	// >  The image caches that fail to be created are only retained for one day.
	//
	// example:
	//
	// 7
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-uf66jeqopgqa9hdn****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The number of regular snapshots. By default, the system creates one snapshot for each image cache. If an image cache is used to create multiple elastic container instances at a time, we recommend that you set this parameter to create multiple snapshots for the image cache. We recommend that you create one snapshot for creation of every 1,000 elastic container instances.
	//
	// >  If you set the Flash parameter to false, instant image cache is disabled. In this case, only regular snapshots are generated during the creation of the image cache.
	//
	// example:
	//
	// 7
	StandardCopyCount *int32 `json:"StandardCopyCount,omitempty" xml:"StandardCopyCount,omitempty"`
	// The tag of the image cache.
	Tag []*CreateImageCacheRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch. You can specify up to 10 vSwitch IDs. Separate multiple vSwitch IDs with commas (,). Example: `vsw-***,vsw-***`.
	//
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the image cache.
	//
	// example:
	//
	// cn-hangzhou-g
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

func (s *CreateImageCacheRequest) SetOsType(v string) *CreateImageCacheRequest {
	s.OsType = &v
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
	// The Alibaba Cloud Resource Name (ARN) of the RAM roles in the Alibaba Cloud account to which the elastic container instance belongs.
	//
	// example:
	//
	// acs:ram::1609982529******:role/role-assume
	ArnService *string `json:"ArnService,omitempty" xml:"ArnService,omitempty"`
	// The ARN of the RAM roles in the Alibaba Cloud account to which the Container Registry Enterprise Edition instance belongs.
	//
	// example:
	//
	// acs:ram::1298452580******:role/role-acr
	ArnUser *string `json:"ArnUser,omitempty" xml:"ArnUser,omitempty"`
	// The domain names of the Container Registry Enterprise Edition instance. By default, all domain names of the instance are displayed. You can specify multiple domain names. Separate multiple domain names with commas (,).
	//
	// example:
	//
	// test****-registry.cn-beijing.cr.aliyuncs.com
	Domain []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
	// The ID of Container Registry Enterprise Edition instance N.
	//
	// example:
	//
	// cri-nwj395hgf6f3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of Container Registry Enterprise Edition instance N.
	//
	// example:
	//
	// test****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID of Container Registry Enterprise Edition instance N.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateImageCacheRequestAcrRegistryInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheRequestAcrRegistryInfo) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequestAcrRegistryInfo) SetArnService(v string) *CreateImageCacheRequestAcrRegistryInfo {
	s.ArnService = &v
	return s
}

func (s *CreateImageCacheRequestAcrRegistryInfo) SetArnUser(v string) *CreateImageCacheRequestAcrRegistryInfo {
	s.ArnUser = &v
	return s
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
	//
	// example:
	//
	// password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The address of the image repository without the `http://` or `https://` prefix.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The username that is used to log on to image repository N.
	//
	// example:
	//
	// username
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
	//
	// example:
	//
	// imc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the image cache. Valid values of N: 1 to 20.
	//
	// example:
	//
	// test
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
	//
	// example:
	//
	// eci-2zebxkiifuyzzlhl****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The ID of the image cache.
	//
	// example:
	//
	// imc-2zebxkiifuyzzlhl****
	ImageCacheId *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0E234675-3465-4CC3-9D0F-9A864BC391DD
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// This parameter is required.
	//
	// example:
	//
	// eci-bp1dvysdafbh00t7****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The type of the O&M task. Valid values:
	//
	// 	- coredump
	//
	// 	- tcpdump
	//
	// This parameter is required.
	//
	// example:
	//
	// coredump
	OpsType *string `json:"OpsType,omitempty" xml:"OpsType,omitempty"`
	// The value of the O\\&M task. You can set this parameter based on the value of OpsType.
	//
	// 	- If OpsType is set to coredump, the valid values of OpsValue are:
	//
	//     	- enable: enables coredump.
	//
	//     	- disable: disables coredump.
	//
	// 	- If OpsType is set to tcpdump, the value of OpsValue is a JSON-formatted parameter string. Example: `{"Enable":true, "IfDeviceName":"eth0"}`. The value may contain the following parameters:
	//
	//     	- Enable: specifies whether to enable tcpdump. You must specify this parameter. Valid values: true and false.
	//
	//     	- Protocol: the network protocol. Valid values: tcp, udp, and icmpv4.
	//
	//     	- SourceIp: the source IP address.
	//
	//     	- SourceCidr: the source CIDR block. If you specify both an IP address and a CIDR block, the IP address is ignored if the CIDR block is valid.
	//
	//     	- SourcePort: the source port. Valid values: 1 to 65535.
	//
	//     	- DestIp: the destination IP address.
	//
	//     	- DestCidr: the destination CIDR block. If you specify both an IP address and a CIDR block, the IP address is ignored if the CIDR block is valid.
	//
	//     	- DestPort: the destination port. Valid values: 1 to 65535.
	//
	//     	- IfDeviceName: the destination network interface controller. Default value: eth0.
	//
	//     	- Snaplen: the length to be captured. Default value: 65535. Unit: bytes.
	//
	//     	- Duration: the captured period. Unit: seconds.
	//
	//     	- PacketNum: the number of packets to be captured.
	//
	//     	- FileSize: the size of the destination files on packets. Unit: bytes. Maximum value: 1073741824. 1073741824 bytes is equal to 1 GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// enable
	OpsValue     *string `json:"OpsValue,omitempty" xml:"OpsValue,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the O&M task.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	//
	// example:
	//
	// 89945DD3-9072-47D0-A318-353284CFC7B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The creation result of the O&M task.
	//
	// example:
	//
	// Coredump settings has been configured successfully.
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceOpsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IP address of the DNS server. If dnsPolicy=ClusterFirst is configured for the Elastic Container Instance pod, Elastic Container Instance uses the configuration to provide DNS services to containers. You can configure multiple IP addresses. Separate multiple IP addresses with commas (,).
	//
	// example:
	//
	// 100.1.XX.XX
	ClusterDNS *string `json:"ClusterDNS,omitempty" xml:"ClusterDNS,omitempty"`
	// The domain name of the cluster. If this parameter is specified, in addition to the search domain of the host, Kubelet configures all containers to search for the specified domain name.
	//
	// example:
	//
	// example.com
	ClusterDomain *string `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	// The custom resources that are supported by the virtual node. If a custom resource is specified in the request of an Elastic Container Instance pod, the pod is scheduled to run on the virtual node that supports the custom resource. You can use the Resource name = Number of resources format to specify custom resources. Separate multiple resources with commas (,).
	//
	// example:
	//
	// example1.com=100,example2.com=200
	CustomResources *string `json:"CustomResources,omitempty" xml:"CustomResources,omitempty"`
	// The ID of the elastic IP address (EIP).
	//
	// example:
	//
	// eip-uf66jeqopgqa9hdn****
	EipInstanceId *string `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	// Specifies whether to enable Internet access for the VNode. Default value: false.
	//
	// If the value of this parameter is true, the VNode exposes a public IP address to external services.
	//
	// example:
	//
	// false
	EnablePublicNetwork *bool `json:"EnablePublicNetwork,omitempty" xml:"EnablePublicNetwork,omitempty"`
	// The KubeConfig of the Kubernetes cluster with which the VNode is connected. The value must be Base64-encoded.
	//
	// example:
	//
	// JTVDbmFwaVZlcnNpb24lM0ElMjB2MSU1Q25jbHVzdGVycyUzQSU1Q24tJTIwY2x1c3RlciUzQSU1Q24uLi******
	KubeConfig   *string `json:"KubeConfig,omitempty" xml:"KubeConfig,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the virtual node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-uf66jeqopgqa9hdn****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the security group. The VNode and the elastic container instances in the VNode are added to the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-2ze81zoc3yl7a3we****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The tags to add to the VNode. You can add up to 20 tags.
	Tag []*CreateVirtualNodeRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The taints of the VNode. You can configure up to 20 taints.
	Taint []*CreateVirtualNodeRequestTaint `json:"Taint,omitempty" xml:"Taint,omitempty" type:"Repeated"`
	// Specifies whether to enable TLS bootstrapping. If you set this parameter to true, use the KubeConfig certificate for TLS bootstrapping. Valid values:
	//
	// - true
	//
	// - false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	TlsBootstrapEnabled *bool `json:"TlsBootstrapEnabled,omitempty" xml:"TlsBootstrapEnabled,omitempty"`
	// The ID of the vSwitch. The vSwitch is connected to the VNode and the elastic container instances in the VNode.
	//
	// You can specify 1 to 10 vSwitches for a VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-2ze23nqzig8inprou****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the VNode. The name must be 2 to 128 characters in length, and can contain lowercase letters, digits, periods (.), and hyphens (-).
	//
	// example:
	//
	// testNode
	VirtualNodeName *string `json:"VirtualNodeName,omitempty" xml:"VirtualNodeName,omitempty"`
	// The zone ID of the VNode.
	//
	// example:
	//
	// cn-hangzhou-b
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
	// The tag key.
	//
	// example:
	//
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
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
	// The effect of the taint. Valid values:
	//
	// 	- NoSchedule: does not schedule pods to run on the VNodes that have the taint.
	//
	// 	- NoExecute: evicts existing pods on the VNodes that have the taint while not scheduling pods to run on the VNodes.
	//
	// 	- PreferNoSchedule: avoids scheduling pods to run on the VNodes that have the taint.
	//
	// example:
	//
	// NoSchedule
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// The key of the taint.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the taint.
	//
	// example:
	//
	// testValue
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
	//
	// example:
	//
	// 89164E78-FC82-4684-BE97-DCDD85D26546
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the VNode.
	//
	// example:
	//
	// vnd-2ze960zkdqrldeaw****
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVirtualNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure the idempotence of a request?](https://help.aliyun.com/document_detail/25693.html)
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// eci-2zelg8vwnlzdhf8hv****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// Specifies whether to forcibly delete resources. Default value: false. Valid values:
	//
	// 	- true: forcibly deletes resources without waiting for the timeout period for graceful terminations.
	//
	// 	- false: waits for the timeout period for graceful terminations to delete resources.
	//
	// example:
	//
	// true
	Force        *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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

func (s *DeleteContainerGroupRequest) SetForce(v bool) *DeleteContainerGroupRequest {
	s.Force = &v
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
	// The request ID.
	//
	// example:
	//
	// 45D5B0AD-3B00-4A9B-9911-6D5303B06712
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContainerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DeleteDataCacheRequest struct {
	// The bucket that stores the DataCache. By default, the bucket is named default.
	//
	// example:
	//
	// default
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure the idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 5e74729e-ce21-4c3e-9163-cbaf0f02768e
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the DataCache.
	//
	// example:
	//
	// edc-bp1a7n7uawwwol******
	DataCacheId  *string `json:"DataCacheId,omitempty" xml:"DataCacheId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The directory in which the virtual host of the DataCache is located.
	//
	// example:
	//
	// /data/models/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The region ID of the DataCache.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDataCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataCacheRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataCacheRequest) SetBucket(v string) *DeleteDataCacheRequest {
	s.Bucket = &v
	return s
}

func (s *DeleteDataCacheRequest) SetClientToken(v string) *DeleteDataCacheRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDataCacheRequest) SetDataCacheId(v string) *DeleteDataCacheRequest {
	s.DataCacheId = &v
	return s
}

func (s *DeleteDataCacheRequest) SetOwnerAccount(v string) *DeleteDataCacheRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDataCacheRequest) SetOwnerId(v int64) *DeleteDataCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDataCacheRequest) SetPath(v string) *DeleteDataCacheRequest {
	s.Path = &v
	return s
}

func (s *DeleteDataCacheRequest) SetRegionId(v string) *DeleteDataCacheRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDataCacheRequest) SetResourceOwnerAccount(v string) *DeleteDataCacheRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDataCacheRequest) SetResourceOwnerId(v int64) *DeleteDataCacheRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteDataCacheResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C0F8A012-1426-5EB2-96F5-B4416DEB5B78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataCacheResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataCacheResponseBody) SetRequestId(v string) *DeleteDataCacheResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDataCacheResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataCacheResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataCacheResponse) SetHeaders(v map[string]*string) *DeleteDataCacheResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataCacheResponse) SetStatusCode(v int32) *DeleteDataCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataCacheResponse) SetBody(v *DeleteDataCacheResponseBody) *DeleteDataCacheResponse {
	s.Body = v
	return s
}

type DeleteImageCacheRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure the idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the image cache.
	//
	// This parameter is required.
	//
	// example:
	//
	// imc-2zebxkiifuyzzlhl****
	ImageCacheId *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the image cache.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	// The request ID.
	//
	// example:
	//
	// 89945DD3-9072-47D0-A218-354284CFC7A2
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteImageCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency of requests](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the virtual node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the virtual node.
	//
	// This parameter is required.
	//
	// example:
	//
	// vnd-2ze960zkdqrldeaw****
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
	//
	// example:
	//
	// 89945DD3-9072-47D0-A318-353284CFC7B3
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVirtualNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// This parameter is required.
	DestinationResource *DescribeAvailableResourceRequestDestinationResource `json:"DestinationResource,omitempty" xml:"DestinationResource,omitempty" type:"Struct"`
	OwnerAccount        *string                                              `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64                                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the ECS instance families.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/146965.html) operation to query the most recent list of regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The information about the preemptible instances that you want to query.
	SpotResource *DescribeAvailableResourceRequestSpotResource `json:"SpotResource,omitempty" xml:"SpotResource,omitempty" type:"Struct"`
	// The zone ID of the ECS instance families.
	//
	// This parameter is empty by default, which indicates that ECS instance families available in all zones in the specified region are queried.
	//
	// example:
	//
	// cn-hangzhou-e
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
	// 	- InstanceTypeFamily: queries instance families. If you use this parameter value, you must also specify the Value parameter.
	//
	// 	- InstanceType: queries instance types. If you use this parameter value, you must also specify the Value, Cores, and Memory parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// InstanceTypeFamily
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The number of vCPUs. This parameter is available only when the Category parameter is set to InstanceType.
	//
	// example:
	//
	// 2
	Cores *float32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The size of the memory. Unit: GiB. This parameter is available only when the Category parameter is set to InstanceType.
	//
	// example:
	//
	// 4
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// Instance families or instance types.
	//
	// 	- If you set Category to InstanceTypeFamily, you must set this parameter to instance families such as ecs.c5.
	//
	// 	- If you set Category to InstanceType, you must set this parameter to instance types such as ecs.c5.large.
	//
	// example:
	//
	// ecs.c6
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
	// The protection period of the preemptible instance. Unit: hours. Default value: 1. The value of 0 indicates no protection period.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The maximum hourly price of the preemptible elastic container instance. The value can be accurate to three decimal places. If you set SpotStrategy to SpotWithPriceLimit, you must specify the SpotPriceLimit parameter.
	//
	// example:
	//
	// 0.2
	SpotPriceLimit *float64 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding policy for the elastic container instance. Valid values:
	//
	// 	- NoSpot: The instance is created as a regular pay-as-you-go instance.
	//
	// 	- SpotWithPriceLimit: The instance is created as a preemptible instance with a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instance is created as a preemptible instance for which the market price at the time of purchase is automatically used as the bid price.
	//
	// Default value: NoSpot.
	//
	// > If you set this parameter to SpotWithPriceLimit or SpotAsPriceGo to query preemptible instances, you must set Category to InstanceType. You must also use the Value parameter to specify instance types, or use the Cores and Memory parameters to specify the number of vCPUs and memory size.
	//
	// example:
	//
	// SpotAsPriceGo
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
	// The zones in which the specified resources are available.
	AvailableZones *DescribeAvailableResourceResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6829735A-FF76-58C6-AECB-27CBF135A7AA
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
	// The resources that are available in the specified zone.
	AvailableResources *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Struct"`
	// The region ID of the resources.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The zone ID of the resources.
	//
	// example:
	//
	// cn-hangzhou-e
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
	// The information about the resources that are available in the zones.
	SupportedResources *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources `json:"SupportedResources,omitempty" xml:"SupportedResources,omitempty" type:"Struct"`
	// The type of the resource. Valid values:
	//
	// 	- InstanceTypeFamily: instance families.
	//
	// 	- InstanceType: instance types.
	//
	// example:
	//
	// InstanceTypeFamily
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
	// The category of resources based on stock status. Valid values:
	//
	// 	- WithStock: Resources are in sufficient stock.
	//
	// 	- ClosedWithStock: Resources are insufficient. We recommend that you use instance types that are in sufficient stock.
	//
	// 	- WithoutStock: Resources are sold out and will be replenished. We recommend that you use instance types that are in sufficient stock.
	//
	// 	- ClosedWithoutStock: Resources are sold out and will not be replenished. We recommend that you use instance types that are in sufficient stock.
	//
	// example:
	//
	// WithStock
	StatusCategory *string `json:"StatusCategory,omitempty" xml:"StatusCategory,omitempty"`
	// The ECS instance types or instance families that are available in the zones.
	//
	// 	- If the return value of the Type parameter is InstanceTypeFamily, this parameter indicates instance families that are returned.
	//
	// 	- If the return value of the Type parameter is InstanceType, this parameter indicates instance types that are returned.
	//
	// example:
	//
	// ecs.c6
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the elastic container instance on which the CommitContainer task is executed.\\
	//
	// You must enter the instance ID, the task ID, or both for the request.
	//
	// example:
	//
	// eci-2zelg8vwnlzdhf8hv****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The number of entries to return on each page.\\
	//
	// Maximum value: 50.\\
	//
	// Default value: 10.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Set the value to the value of NextToken that is returned from the last request.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the task.
	TaskId []*string `json:"TaskId,omitempty" xml:"TaskId,omitempty" type:"Repeated"`
	// The status of the task. Valid values:
	//
	// 	- Running
	//
	// 	- Succeeded
	//
	// 	- Failed
	//
	// example:
	//
	// Running
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
	// Details of the task.
	CommitTasks []*DescribeCommitContainerTaskResponseBodyCommitTasks `json:"CommitTasks,omitempty" xml:"CommitTasks,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token that is returned in this request.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 45D5B0AD-3B00-4A9B-9911-6D5303B06712
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
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
	// The information about the phase that the task arrives.
	CommitPhaseInfos []*DescribeCommitContainerTaskResponseBodyCommitTasksCommitPhaseInfos `json:"CommitPhaseInfos,omitempty" xml:"CommitPhaseInfos,omitempty" type:"Repeated"`
	// The container name.
	//
	// example:
	//
	// worker0
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The message about the state.
	//
	// example:
	//
	// Commit container -xxxxxx
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The time when the task was started.
	//
	// example:
	//
	// 2023-03-06T08:22:40Z
	TaskCreationTime *string `json:"TaskCreationTime,omitempty" xml:"TaskCreationTime,omitempty"`
	// The time when the task was complete.
	//
	// example:
	//
	// 2023-03-06T08:23:40Z
	TaskFinishedTime *string `json:"TaskFinishedTime,omitempty" xml:"TaskFinishedTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-2zej6nstkg744qc3****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The progress of the task in percentage.
	//
	// example:
	//
	// 50%
	TaskProgress *string `json:"TaskProgress,omitempty" xml:"TaskProgress,omitempty"`
	// The state of the task. Valid values:
	//
	// 	- Running
	//
	// 	- Succeeded
	//
	// 	- Failed
	//
	// example:
	//
	// Running
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

func (s *DescribeCommitContainerTaskResponseBodyCommitTasks) SetStatusMessage(v string) *DescribeCommitContainerTaskResponseBodyCommitTasks {
	s.StatusMessage = &v
	return s
}

func (s *DescribeCommitContainerTaskResponseBodyCommitTasks) SetTaskCreationTime(v string) *DescribeCommitContainerTaskResponseBodyCommitTasks {
	s.TaskCreationTime = &v
	return s
}

func (s *DescribeCommitContainerTaskResponseBodyCommitTasks) SetTaskFinishedTime(v string) *DescribeCommitContainerTaskResponseBodyCommitTasks {
	s.TaskFinishedTime = &v
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
	// The message about the phase.
	//
	// example:
	//
	// Pull base image for container container-1 successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The phase name. Valid values:
	//
	// 	- PullBaseImage: Pull the original container image.
	//
	// 	- CommitContainer: Commit the container to generate an image.
	//
	// 	- PushCommittedImage: Push the image to Container Registry.
	//
	// example:
	//
	// PullBaseImage
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The record time of the phase.
	//
	// example:
	//
	// 2023-01-05T14:06:40.920005316+08:00
	RecordTime *string `json:"RecordTime,omitempty" xml:"RecordTime,omitempty"`
	// The state of the phase.
	//
	// example:
	//
	// Success
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCommitContainerTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The IDs of the elastic container instances. You can specify up to 20 IDs. Each ID must be a JSON string.
	//
	// example:
	//
	// ["eci-bp17gw49eu09yiwm****", "eci-bp19aq49du01abcm****", "eci-2zegym1qhbmdfr1s****"]
	ContainerGroupIds *string `json:"ContainerGroupIds,omitempty" xml:"ContainerGroupIds,omitempty"`
	// The event source. Valid values:
	//
	// 	- EciService
	//
	// 	- K8sAgent
	//
	// This parameter is empty by default. This indicates that all events are queried.
	//
	// example:
	//
	// K8sAgent
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The maximum number of elastic container instances to be returned for this request. Default value: 200.
	//
	// >  The number of elastic container instances to be returned is no greater than this parameter value.
	//
	// example:
	//
	// 200
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// d78f2dd8-5979-42fe-xaee-b16db43be5bc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-uf66jeqopgqa9hdn****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// A relative time in seconds before the current time from which to show event information. This parameter is used to poll incremental events.
	//
	// example:
	//
	// 60
	SinceSecond *int32 `json:"SinceSecond,omitempty" xml:"SinceSecond,omitempty"`
	// The tag that is added to the elastic container instances.
	Tag []*DescribeContainerGroupEventsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-k
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
	// The key of the tag.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// name
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
	// Details of the events.
	Data []*DescribeContainerGroupEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1340C38D-6189-54D1-86F6-7D5ECF3E0088
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries of returned events.
	//
	// example:
	//
	// 10
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
	// The annotations of the elastic container instance.
	//
	// example:
	//
	// "{\\"tenancy.x-k8s.io/namespace\\":\\"redis\\"}"
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// The ID of the elastic container instance.
	//
	// example:
	//
	// eci-uf6fonnghi50v6nq****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The events.
	Events []*DescribeContainerGroupEventsResponseBodyDataEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The name of the elastic container instance.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace where the elastic container instance resides.
	//
	// example:
	//
	// redis
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The UUID of the elastic container instance.
	//
	// example:
	//
	// 3fc6b309-****-****
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s DescribeContainerGroupEventsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsResponseBodyData) SetAnnotations(v string) *DescribeContainerGroupEventsResponseBodyData {
	s.Annotations = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyData) SetContainerGroupId(v string) *DescribeContainerGroupEventsResponseBodyData {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyData) SetEvents(v []*DescribeContainerGroupEventsResponseBodyDataEvents) *DescribeContainerGroupEventsResponseBodyData {
	s.Events = v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyData) SetName(v string) *DescribeContainerGroupEventsResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyData) SetNamespace(v string) *DescribeContainerGroupEventsResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyData) SetUuid(v string) *DescribeContainerGroupEventsResponseBodyData {
	s.Uuid = &v
	return s
}

type DescribeContainerGroupEventsResponseBodyDataEvents struct {
	// The number of events.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The first occurrence time of the event.
	//
	// example:
	//
	// 2021-10-04T09:08:04Z
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	// The most recent occurrence time of the event.
	//
	// example:
	//
	// 2021-10-04T09:08:04Z
	LastTimestamp *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The message about the event.
	//
	// example:
	//
	// Started container
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The metadata of the event.
	Metadata *DescribeContainerGroupEventsResponseBodyDataEventsMetadata `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Struct"`
	// The cause of the event.
	//
	// example:
	//
	// Started
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The component from which the event is reported.
	//
	// example:
	//
	// test
	ReportingComponent *string `json:"ReportingComponent,omitempty" xml:"ReportingComponent,omitempty"`
	// The instance from which the event is reported.
	//
	// example:
	//
	// test
	ReportingInstance *string `json:"ReportingInstance,omitempty" xml:"ReportingInstance,omitempty"`
	// The source.
	Source *DescribeContainerGroupEventsResponseBodyDataEventsSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	// The event type. Valid values:
	//
	// 	- Normal
	//
	// 	- Warning
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The resource object that the event is about.
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
	// The event name.
	//
	// example:
	//
	// eci-uto-created-eci-for-ubuntu.167e3fb73cc7f9cb
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
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
	// The component.
	//
	// example:
	//
	// kubelet
	Component *string `json:"Component,omitempty" xml:"Component,omitempty"`
	// The host.
	//
	// example:
	//
	// eci
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
	//
	// example:
	//
	// v1
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	// The resource type.
	//
	// example:
	//
	// Pod
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The resource name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace where the resource resides.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// eci-8vb5nkcq3a5cu0p3****
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerGroupEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// eci-bp17da5olmkuagfr****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The end of the time range to query. The default value is the current time.
	//
	// Specify the time in RFC 3339 format.
	//
	// example:
	//
	// 2019-09-12T01:05:00.000Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The data aggregation period. Unit: seconds. Valid values: 15, 30, 60, and 600. Default value: 60.
	//
	// >  If the StartTime and EndTime parameters are not specified, the system returns the monitoring data generated in the last 5 minutes with a data aggregation period of 15s. The Period parameter is ignored.
	//
	// example:
	//
	// 15
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. The beginning of the time range must be a time point in the last 30 days. The default value is 5 minutes before the value of EndTime.
	//
	// Specify the time in RFC 3339 format. For example, to query the data starting from March 12, 2019, 09:00 UTC+8, you can set the parameter to 2019-03-12T09:00:00.000+08:00 or 2019-03-12T01:00:00.000Z.
	//
	// example:
	//
	// 2019-03-12T01:00:00.000Z
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
	// The instance ID.
	//
	// example:
	//
	// eci-bp1dqpye03ke5s6p****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The monitoring data of the elastic container instance.
	Records []*DescribeContainerGroupMetricResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D81A4A13-6DCC-4579-AC62-90A6C3EC7BBC
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
	// The time when the monitoring data entry was collected. The time follows the RFC 3339 format.
	//
	// example:
	//
	// 2021-08-25T16:22:00.000+08:00
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
	// The upper limit of vCPU usage. The calculation formula for this parameter: The number of vCPUs × 1000.
	//
	// example:
	//
	// 2000
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The average load in the last 10 seconds.
	//
	// example:
	//
	// 0
	Load *int64 `json:"Load,omitempty" xml:"Load,omitempty"`
	// The cumulative usage of vCPUs.
	//
	// example:
	//
	// 70769883
	UsageCoreNanoSeconds *int64 `json:"UsageCoreNanoSeconds,omitempty" xml:"UsageCoreNanoSeconds,omitempty"`
	// The vCPU usage in the sampling window. Unit for the sampling window: nanoseconds.
	//
	// example:
	//
	// 0
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
	//
	// example:
	//
	// u1
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
	// The upper limit of vCPU usage. The calculation formula for this parameter: The number of vCPUs × 1000.
	//
	// example:
	//
	// 1024
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The average load in the last 10 seconds.
	//
	// example:
	//
	// 0
	Load *int64 `json:"Load,omitempty" xml:"Load,omitempty"`
	// The cumulative usage of vCPUs.
	//
	// example:
	//
	// 40876694
	UsageCoreNanoSeconds *int64 `json:"UsageCoreNanoSeconds,omitempty" xml:"UsageCoreNanoSeconds,omitempty"`
	// The vCPU usage in the sampling window. Unit for the sampling window: nanoseconds.
	//
	// example:
	//
	// 0
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
	// The size of the available memory. Unit: bytes.
	//
	// example:
	//
	// 4285091840
	AvailableBytes *int64 `json:"AvailableBytes,omitempty" xml:"AvailableBytes,omitempty"`
	// The size of the cache. Unit: bytes.
	//
	// example:
	//
	// 9289728
	Cache *int64 `json:"Cache,omitempty" xml:"Cache,omitempty"`
	// The size of the resident memory set, which indicates the size of the physical memory that is actually used. Unit: bytes.
	//
	// example:
	//
	// 9289728
	Rss *int64 `json:"Rss,omitempty" xml:"Rss,omitempty"`
	// The size of the used memory. Unit: bytes.
	//
	// example:
	//
	// 9289728
	UsageBytes *int64 `json:"UsageBytes,omitempty" xml:"UsageBytes,omitempty"`
	// The usage of the working set. Unit: bytes.
	//
	// example:
	//
	// 9289728
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
	//
	// example:
	//
	// /dev/vda
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The amount of data that was read from the disk. Unit: bytes.
	//
	// example:
	//
	// 210714112
	ReadBytes *int64 `json:"ReadBytes,omitempty" xml:"ReadBytes,omitempty"`
	// This parameter is unavailable for public use.
	//
	// example:
	//
	// 123
	ReadIO *int64 `json:"ReadIO,omitempty" xml:"ReadIO,omitempty"`
	// The amount of data that was written to the disk. Unit: bytes.
	//
	// example:
	//
	// 1375109120
	WriteBytes *int64 `json:"WriteBytes,omitempty" xml:"WriteBytes,omitempty"`
	// This parameter is unavailable for public use.
	//
	// example:
	//
	// 123
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
	//
	// example:
	//
	// 35319836672
	Available *int64 `json:"Available,omitempty" xml:"Available,omitempty"`
	// The total file system space.
	//
	// example:
	//
	// 41610981376
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The type of the partition. Valid values:
	//
	// 	- System
	//
	// 	- Volume
	//
	// 	- Other
	//
	// example:
	//
	// System
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The name of the partition.
	//
	// example:
	//
	// /dev/root
	FsName *string `json:"FsName,omitempty" xml:"FsName,omitempty"`
	// The size of used space.
	//
	// example:
	//
	// 4368744448
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
	// The size of the available memory. Unit: bytes.
	//
	// example:
	//
	// 4285091840
	AvailableBytes *int64 `json:"AvailableBytes,omitempty" xml:"AvailableBytes,omitempty"`
	// The size of the cache. Unit: bytes.
	//
	// example:
	//
	// 9289728
	Cache *int64 `json:"Cache,omitempty" xml:"Cache,omitempty"`
	// The size of the resident memory set, which indicates the size of the physical memory that is actually used. Unit: bytes.
	//
	// example:
	//
	// 9289728
	Rss *int64 `json:"Rss,omitempty" xml:"Rss,omitempty"`
	// The size of the used memory. Unit: bytes.
	//
	// example:
	//
	// 9289728
	UsageBytes *int64 `json:"UsageBytes,omitempty" xml:"UsageBytes,omitempty"`
	// The usage of the working set. Unit: bytes.
	//
	// example:
	//
	// 9289728
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
	//
	// example:
	//
	// eth0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of bytes received by the NIC.
	//
	// example:
	//
	// 505001954
	RxBytes *int64 `json:"RxBytes,omitempty" xml:"RxBytes,omitempty"`
	// The number of received packets dropped on the NIC.
	//
	// example:
	//
	// 0
	RxDrops *int64 `json:"RxDrops,omitempty" xml:"RxDrops,omitempty"`
	// The number of received packet errors on the NIC.
	//
	// example:
	//
	// 0
	RxErrors *int64 `json:"RxErrors,omitempty" xml:"RxErrors,omitempty"`
	// The number of packets received by the NIC.
	//
	// example:
	//
	// 4800583
	RxPackets *int64 `json:"RxPackets,omitempty" xml:"RxPackets,omitempty"`
	// The number of bytes transmitted by the NIC.
	//
	// example:
	//
	// 1381805699
	TxBytes *int64 `json:"TxBytes,omitempty" xml:"TxBytes,omitempty"`
	// The number of transmitted packets dropped on the NIC.
	//
	// example:
	//
	// 0
	TxDrops *int64 `json:"TxDrops,omitempty" xml:"TxDrops,omitempty"`
	// The number of transmitted packet errors on the NIC.
	//
	// example:
	//
	// 0
	TxErrors *int64 `json:"TxErrors,omitempty" xml:"TxErrors,omitempty"`
	// The number of packets transmitted by the NIC.
	//
	// example:
	//
	// 5158427
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerGroupMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The computing power type. A value of economy specifies economic instances.
	//
	// example:
	//
	// economy
	ComputeCategory *string `json:"ComputeCategory,omitempty" xml:"ComputeCategory,omitempty"`
	// The number of vCPUs. For information about the vCPU and memory specifications that are supported by Elastic Container Instance, see [vCPU and memory specifications](https://help.aliyun.com/document_detail/114662.html).
	//
	// example:
	//
	// 2.0
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The storage size of the temporary storage space. Unit: GiB.
	//
	// example:
	//
	// 20
	EphemeralStorage *int32 `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	// The instance type of the Elastic Compute Service (ECS) instance that is used to create the elastic container instance. For information about the ECS instance types that are supported by Elastic Container Instance, see [ECS instance types that are supported by Elastic Container Instance](https://help.aliyun.com/document_detail/114664.html).
	//
	// > If you specify this parameter, the specified specifications of vCPUs and memory are ignored. Only the price of the ECS instance type is returned.
	//
	// example:
	//
	// ecs.c5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The size of the memory. Unit: GiB. For information about the vCPU and memory specifications that are supported by Elastic Container Instance, see [vCPU and memory specifications](https://help.aliyun.com/document_detail/114662.html).
	//
	// example:
	//
	// 4.0
	Memory       *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	OwnerAccount *string  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/146965.html) operation to query the most recent region and zone list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The protection period of the preemptible instance. Unit: hours. Default value: 1. The value of 0 indicates no protection period.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The maximum hourly price of the preemptible elastic container instance. The value can contain up to three decimal places. If you set SpotStrategy to SpotWithPriceLimit, you must specify SpotPriceLimit.
	//
	// example:
	//
	// 0.2
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding policy for the elastic container instance. Valid values:
	//
	// 	- NoSpot: The instance is a regular pay-as-you-go instance.
	//
	// 	- SpotWithPriceLimit: The instance is a preemptible instance with a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instance is a preemptible instance for which the market price at the time of purchase is automatically used as the bid price.
	//
	// Default value: NoSpot.
	//
	// example:
	//
	// SpotWithPriceLimit
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The zone ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/146965.html) operation to query the most recent region and zone list.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeContainerGroupPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceRequest) SetComputeCategory(v string) *DescribeContainerGroupPriceRequest {
	s.ComputeCategory = &v
	return s
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
	// The information about the prices and discount rules.
	PriceInfo *DescribeContainerGroupPriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7A872585-33C7-4D69-AB8E-412E81EBA387
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
	// The information about the prices of preemptible elastic container instances.
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
	// 	- CNY: This value only applies to the China site (aliyun.com).
	//
	// 	- USD: This value only applies to the International site (alibabacloud.com).
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The information about the price.
	DetailInfos *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos `json:"DetailInfos,omitempty" xml:"DetailInfos,omitempty" type:"Struct"`
	// The discount.
	//
	// example:
	//
	// 0
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 0.00012252
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The transaction price, which is equal to the original price minus the discount.
	//
	// example:
	//
	// 0.00012252
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
	//
	// example:
	//
	// 0
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 0.000098
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// cpu
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// Details about the pricing rules.
	Rules *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	// The transaction price.
	//
	// example:
	//
	// 0.000098
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
	// The description of the rule.
	//
	// example:
	//
	// Receive a 15% discount on a 1-year subscription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 123456
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
	//
	// example:
	//
	// Receive a 15% discount on a 1-year subscription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the promotion rule.
	//
	// example:
	//
	// 123456
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
	// The ECS instance type.
	//
	// example:
	//
	// ecs.c5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The original price.
	//
	// example:
	//
	// 0.622
	OriginPrice *float32 `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	// The prices of preemptible elastic container instances.
	//
	// example:
	//
	// 0.056
	SpotPrice *float32 `json:"SpotPrice,omitempty" xml:"SpotPrice,omitempty"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-hangzhou-i
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerGroupPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The IDs of the instances. You can specify up to 20 IDs. Each ID must be a string in the JSON format.
	//
	// example:
	//
	// ["eci-bp17gw49eu09yiwm****", "eci-bp19aq49du01abcm****", "eci-2zegym1qhbmdfr1s****"]
	ContainerGroupIds *string `json:"ContainerGroupIds,omitempty" xml:"ContainerGroupIds,omitempty"`
	// Specifies the maximum number of elastic container instances to be returned for this request. Default value: 200.
	//
	// > The number of returned resources can be less than or equal to the value of this parameter.
	//
	// example:
	//
	// 200
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.\\
	//
	// You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// d78f2dd8-5979-42fe-xaee-b16db43be5bc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instances belong.
	//
	// example:
	//
	// rg-uf66jeqopgqa9hdn****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// A relative time in seconds before the current time from which to show elastic container instances whose status changes. This parameter is used to poll status of elastic container instances.
	//
	// example:
	//
	// 60
	SinceSecond *int32 `json:"SinceSecond,omitempty" xml:"SinceSecond,omitempty"`
	// The tag that is bound to the instances.
	Tag []*DescribeContainerGroupStatusRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the instances.
	//
	// example:
	//
	// cn-hangzhou-k
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
	// The key of the tag.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// testValue
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
	// The collection of the statuses of the container groups.
	Data []*DescribeContainerGroupStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// d78f2dd8-5979-42fe-****-b16db43be5bc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1340C38D-6189-54D1-86F6-7D5ECF3E0088
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
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
	// Annotations that are added to the container groups.
	//
	// example:
	//
	// "{\\"tenancy.x-k8s.io/namespace\\":\\"redis\\"}"
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// The ID of the container group.
	//
	// example:
	//
	// eci-bp1jrgfqqy54kg5hc****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The name of the container group.
	//
	// example:
	//
	// nginx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace in which the container group resides.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The status of the container group.
	PodStatus *DescribeContainerGroupStatusResponseBodyDataPodStatus `json:"PodStatus,omitempty" xml:"PodStatus,omitempty" type:"Struct"`
	// The status of the container group.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The universally unique identifier (UUID) of the container group, which is similar to the unique identifier (UID) of the Kubernetes pod in terms of the concept and usage.
	//
	// example:
	//
	// 78ee0657-987g-b8b2-1f507dic4****
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyData) SetAnnotations(v string) *DescribeContainerGroupStatusResponseBodyData {
	s.Annotations = &v
	return s
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
	// The conditions of the container group.
	Conditions []*DescribeContainerGroupStatusResponseBodyDataPodStatusConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The statuses about the containers.
	ContainerStatuses []*DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses `json:"ContainerStatuses,omitempty" xml:"ContainerStatuses,omitempty" type:"Repeated"`
	// The IP address of the host.
	//
	// example:
	//
	// 192.168.XX.XX
	HostIp *string `json:"HostIp,omitempty" xml:"HostIp,omitempty"`
	// The lifecycle phase of the container group.
	//
	// example:
	//
	// Running
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The IP address of the container group.
	//
	// example:
	//
	// 192.168.XX.XX
	PodIp *string `json:"PodIp,omitempty" xml:"PodIp,omitempty"`
	// The IP addresses of the container groups.
	PodIps []*DescribeContainerGroupStatusResponseBodyDataPodStatusPodIps `json:"PodIps,omitempty" xml:"PodIps,omitempty" type:"Repeated"`
	// The quality of service (QoS) of the container group.
	//
	// example:
	//
	// Guaranteed
	QosClass *string `json:"QosClass,omitempty" xml:"QosClass,omitempty"`
	// The time when the container started to run.
	//
	// example:
	//
	// 2021-05-12T07:02:47Z
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
	// The message about the event.
	//
	// example:
	//
	// Completed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason for the transition into the current status of the event.
	//
	// example:
	//
	// Completed
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The time when the status last changed.
	//
	// example:
	//
	// 2021-05-12T07:02:47Z
	LastTransitionTime *string `json:"lastTransitionTime,omitempty" xml:"lastTransitionTime,omitempty"`
	// The status of the condition.
	//
	// example:
	//
	// True
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the condition. Valid values:
	//
	// 	- PodReadyToStartContainers
	//
	// 	- Initialized
	//
	// 	- Ready
	//
	// 	- ContainersReady
	//
	// 	- PodScheduled
	//
	// 	- ContainerHasSufficientDisk
	//
	// 	- ContainerInstanceCreated
	//
	// 	- Unschedulable
	//
	// example:
	//
	// Ready
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
	//
	// example:
	//
	// registry-vpc.cn-zhangjiakou.aliyuncs.com/eci_open/ubuntu:18.04
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image ID.
	//
	// example:
	//
	// registry-vpc.cn-zhangjiakou.aliyuncs.com/eci_open/ubuntu@sha256:134c7fe821b9d359490cd009ce7ca322453f4f2d018623f849e580a89a685e5d
	ImageID *string `json:"ImageID,omitempty" xml:"ImageID,omitempty"`
	// The last status of the container.
	LastState *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState `json:"LastState,omitempty" xml:"LastState,omitempty" type:"Struct"`
	// The name of the container.
	//
	// example:
	//
	// ubuntu
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the container is ready.
	//
	// example:
	//
	// true
	Ready *bool `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// The number of times that the container restarted.
	//
	// example:
	//
	// 1
	RestartCount *int32 `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	// Indicates whether the container is started.
	//
	// example:
	//
	// true
	Started *bool `json:"Started,omitempty" xml:"Started,omitempty"`
	// The status of the container. Valid values:
	//
	// 	- Waiting
	//
	// 	- Running
	//
	// 	- Terminated
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
	// The container is terminated and exits after a successful or failed running.
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
	// The time when the container started to run.
	//
	// example:
	//
	// 2021-05-23T20:49:31Z
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
	// The container ID.
	//
	// example:
	//
	// containerd://3ff993933bea366c4719e43a1b067d89bc7f01f1f573981659a44ff17a******
	ContainerID *string `json:"ContainerID,omitempty" xml:"ContainerID,omitempty"`
	// The exit code.
	//
	// example:
	//
	// 0
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time when the container ends running.
	//
	// example:
	//
	// 2021-05-23T20:49:31Z
	FinishedAt *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	// The message about the event.
	//
	// example:
	//
	// Completed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason for the transition into the current status of the event.
	//
	// example:
	//
	// Completed
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The signal code.
	//
	// example:
	//
	// 1
	Signal *int32 `json:"Signal,omitempty" xml:"Signal,omitempty"`
	// The time when the container started to run.
	//
	// example:
	//
	// 2021-05-12T07:02:52Z
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
	// The message about the event.
	//
	// example:
	//
	// Back-off
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason for the transition into the current status of the event.
	//
	// example:
	//
	// Started
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
	// The container is terminated and exits after a successful or failed running.
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
	// The time when the container started to run.
	//
	// example:
	//
	// 2021-05-23T20:49:31Z
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
	// The container ID.
	//
	// example:
	//
	// containerd://3ff993933bea366c4719e43a1b067d89bc7f01f1f573981659a44ff17a******
	ContainerID *string `json:"ContainerID,omitempty" xml:"ContainerID,omitempty"`
	// The exit code.
	//
	// example:
	//
	// 0
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time when the container ends running.
	//
	// example:
	//
	// 2021-05-23T20:55:31Z
	FinishedAt *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	// The message about the event.
	//
	// example:
	//
	// Completed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason for the transition into the current status of the event.
	//
	// example:
	//
	// Completed
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The signal code.
	//
	// example:
	//
	// 1
	Signal *int32 `json:"Signal,omitempty" xml:"Signal,omitempty"`
	// The time when the container started to run.
	//
	// example:
	//
	// 2021-05-23T20:49:31Z
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
	// The message about the event.
	//
	// example:
	//
	// Back-off
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason for the transition into the current status of the event.
	//
	// example:
	//
	// Started
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
	// The IP address of the container group.
	//
	// example:
	//
	// 192.168.XX.XX
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerGroupStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The compute category of the instance. Valid values:
	//
	// 	- economy
	//
	// 	- general
	//
	// example:
	//
	// economy
	ComputeCategory *string `json:"ComputeCategory,omitempty" xml:"ComputeCategory,omitempty"`
	// The IDs of the elastic container instances in JSON format. You can specify up to 20 IDs.
	//
	// example:
	//
	// ["eci-bp17gw49eu09yiwm****", "eci-bp19aq49du01abcm****", "eci-2zegym1qhbmdfr1s****"]
	ContainerGroupIds *string `json:"ContainerGroupIds,omitempty" xml:"ContainerGroupIds,omitempty"`
	// The name of the elastic container instance.
	//
	// example:
	//
	// test
	ContainerGroupName *string `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	// The maximum number of resources to return. Default value: 20. Maximum value: 20.
	//
	// >  The number of returned resources is less than or equal to the specified number.
	//
	// example:
	//
	// 20
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The token that determines the start point of the query. If this parameter is left empty, all results have been returned.
	//
	// > You do not need to specify this parameter in the first request. Starting from the second request, you can obtain the token from the result returned by the previous request.
	//
	// example:
	//
	// d78f2dd8-5979-42fe-****-b16db43be5bc
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aekzh43v*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the security group to which the instance belongs.
	//
	// example:
	//
	// sg-uf66jeqopgqa9hdn****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The status of the elastic container instance. Valid values:
	//
	// 	- Pending: The instance is being started.
	//
	// 	- Running: The instance is running.
	//
	// 	- Succeeded: The instance runs successfully.
	//
	// 	- Failed: The instance fails to run.
	//
	// 	- Scheduling: The instance is being created.
	//
	// 	- ScheduleFailed: The instance fails to be created.
	//
	// 	- Restarting: The instance is being restarted.
	//
	// 	- Updating: The instance is being updated.
	//
	// 	- Terminating: The instance is being terminated.
	//
	// 	- Expired: The instance expires.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag of the instances.
	Tag []*DescribeContainerGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which the elastic container instances are connected.
	//
	// example:
	//
	// vsw-bp1jrgfqqy54kg5hc****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Specify whether to return event information.
	//
	// example:
	//
	// true
	WithEvent *bool `json:"WithEvent,omitempty" xml:"WithEvent,omitempty"`
	// The ID of the zone in which the elastic container instances are deployed. If you do not specify this parameter, the system selects a zone.
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeContainerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsRequest) SetComputeCategory(v string) *DescribeContainerGroupsRequest {
	s.ComputeCategory = &v
	return s
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

func (s *DescribeContainerGroupsRequest) SetSecurityGroupId(v string) *DescribeContainerGroupsRequest {
	s.SecurityGroupId = &v
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
	//
	// example:
	//
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instances.
	//
	// example:
	//
	// test
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
	// Details about the queried elastic container instances.
	ContainerGroups []*DescribeContainerGroupsResponseBodyContainerGroups `json:"ContainerGroups,omitempty" xml:"ContainerGroups,omitempty" type:"Repeated"`
	// The token that determines the start point of the query.
	//
	// example:
	//
	// d78f2dd8-5979-42fe-****-b16db43be5bc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request. The value is unique.
	//
	// example:
	//
	// 89945DD3-9072-47D0-A318-353284CFC7B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of queried instances.
	//
	// example:
	//
	// 1
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
	// The computing power type of the elastic container instance. Valid values:
	//
	// 	- economy: economic computing power.
	//
	// 	- general: general-purpose computing power.
	//
	// example:
	//
	// economy
	ComputeCategory *string `json:"ComputeCategory,omitempty" xml:"ComputeCategory,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// eci-bp1jrgfqqy54kg5hc****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// test
	ContainerGroupName *string `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	// The information about containers in the elastic container instance.
	Containers []*DescribeContainerGroupsResponseBodyContainerGroupsContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// The number of vCPUs that are allocated to the elastic container instance.
	//
	// example:
	//
	// 2.0
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The time when the instance was created. The time follows the RFC 3339 standard. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-02T15:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The discount.
	//
	// example:
	//
	// 9
	Discount *int32 `json:"Discount,omitempty" xml:"Discount,omitempty"`
	// The Domain Name System (DNS) settings.
	DnsConfig *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig `json:"DnsConfig,omitempty" xml:"DnsConfig,omitempty" type:"Struct"`
	// The Domain Name System (DNS) policy. Valid values:
	//
	// - None: uses the DNS that is specified by DnsConfig.
	//
	// - Default: uses the DNS that is specified for the runtime environment.
	//
	// example:
	//
	// None
	DnsPolicy *string `json:"DnsPolicy,omitempty" xml:"DnsPolicy,omitempty"`
	// The security context of the elastic container instance.
	EciSecurityContext *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext `json:"EciSecurityContext,omitempty" xml:"EciSecurityContext,omitempty" type:"Struct"`
	// The ID of the elastic network interface (ENI).
	//
	// example:
	//
	// eni-bp14rzqi6fd8satm****
	EniInstanceId *string `json:"EniInstanceId,omitempty" xml:"EniInstanceId,omitempty"`
	// The size of the temporary storage space. Unit: GiB.
	//
	// example:
	//
	// 20
	EphemeralStorage *int32 `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	// The events of the elastic container instance. A maximum of 50 events can be returned.
	Events []*DescribeContainerGroupsResponseBodyContainerGroupsEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The time when the elastic container instance failed to run due to overdue payments. The time follows the RFC 3339 standard. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-02T15:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The time when the instance failed to run. The time follows the RFC 3339 standard. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-02T15:00:00Z
	FailedTime *string `json:"FailedTime,omitempty" xml:"FailedTime,omitempty"`
	// The hostnames and IP addresses for a container that are added to the hosts file of the elastic container instance.
	HostAliases []*DescribeContainerGroupsResponseBodyContainerGroupsHostAliases `json:"HostAliases,omitempty" xml:"HostAliases,omitempty" type:"Repeated"`
	// The information about the init containers.
	InitContainers []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainers `json:"InitContainers,omitempty" xml:"InitContainers,omitempty" type:"Repeated"`
	// The instance type of the specified Elastic Compute Service (ECS) instance.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 192.0.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 172.16.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The IPv6 address of the instance.
	//
	// example:
	//
	// 2001:XXXX:4:4:4:4:4:4
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	// The memory size of the instance. Unit: GiB.
	//
	// example:
	//
	// 4.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the instance RAM role. The elastic container instance and the ECS instance share a RAM role. For more information, see [Use an instance RAM role by calling API operations](https://help.aliyun.com/document_detail/61178.html).
	//
	// example:
	//
	// ram:PassRole
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-8db03793gfrz****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The restart policy of the elastic container instance. Valid values:
	//
	// 	- Never: Never restarts the instance if a container in the instance exits upon termination.
	//
	// 	- Always: Always restarts the instance if a container in the instance exits upon termination.
	//
	// 	- OnFailure: Restarts the instance only if a container in the instance exists upon failure with a status code of non-zero.
	//
	// example:
	//
	// Never
	RestartPolicy *string `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	// The ID of the security group to which the instances belong.
	//
	// example:
	//
	// sg-bp12ujq5zpxuyjfo3o8r
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The maximum hourly price for the preemptible elastic container instance.
	//
	// This parameter is returned only when SpotStrategy is set to SpotWithPriceLimit.
	//
	// example:
	//
	// 0.025
	SpotPriceLimit *float64 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bid policy for the instance. Default value: NoSpot. Valid values:
	//
	// 	- NoSpot: The instance is a regular pay-as-you-go instance.
	//
	// 	- SpotWithPriceLimit: The instance is a preemptible instance that has a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instance is a preemptible instance for which the market price at the time of purchase is automatically used as the bid price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- Pending: The instance is being started.
	//
	// 	- Running: The instance is running.
	//
	// 	- Succeeded: The instance successfully runs.
	//
	// 	- Failed: The instance fails to run.
	//
	// 	- Scheduling: The instance is being created.
	//
	// 	- ScheduleFailed: The instance fails to be created.
	//
	// 	- Restarting: The instance is being restarted.
	//
	// 	- Updating: The instance is being updated.
	//
	// 	- Terminating: The instance is being terminated.
	//
	// 	- Expired: The instance is expired.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when all containers exited on success. The time follows the RFC 3339 standard. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-12-11T04:33:42Z
	SucceededTime *string `json:"SucceededTime,omitempty" xml:"SucceededTime,omitempty"`
	// The tags that are added to the instance.
	Tags []*DescribeContainerGroupsResponseBodyContainerGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// i-xxx
	TenantEniInstanceId *string `json:"TenantEniInstanceId,omitempty" xml:"TenantEniInstanceId,omitempty"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// 10.10.XX.XX
	TenantEniIp *string `json:"TenantEniIp,omitempty" xml:"TenantEniIp,omitempty"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// sg-xxx
	TenantSecurityGroupId *string `json:"TenantSecurityGroupId,omitempty" xml:"TenantSecurityGroupId,omitempty"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// vsw-xxx
	TenantVSwitchId *string `json:"TenantVSwitchId,omitempty" xml:"TenantVSwitchId,omitempty"`
	// The ID of the vSwitch to which the instance is connected.
	//
	// example:
	//
	// vsw-bp1vzjjflab6wvjox****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The information about the volumes.
	Volumes []*DescribeContainerGroupsResponseBodyContainerGroupsVolumes `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
	// The ID of the VPC to which the instance belongs.
	//
	// example:
	//
	// vpc-1vzjjflab6wvjox****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone to which the instance belongs.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroups) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetComputeCategory(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.ComputeCategory = &v
	return s
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

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetDnsPolicy(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.DnsPolicy = &v
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
	// The arguments that are passed to the startup commands of the container.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The startup commands of the container.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The number of vCPUs that are allocated to the container.
	//
	// example:
	//
	// 2.0
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The current container status.
	CurrentState *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState `json:"CurrentState,omitempty" xml:"CurrentState,omitempty" type:"Struct"`
	// The environment variables of the container.
	EnvironmentVars []*DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The number of GPUs.
	//
	// example:
	//
	// 1
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The image in the container.
	//
	// example:
	//
	// mysql
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image pulling policy. Valid values:
	//
	// 	- Always: Each time the instance is updated, image pulling is performed.
	//
	// 	- IfNotPresent: On-premises images are preferentially used. If no on-premises images are available, image pulling is performed.
	//
	// 	- Never: On-premises images are always used. Image pulling is not performed.
	//
	// example:
	//
	// Always
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The liveness probe of the container.
	LivenessProbe *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe `json:"LivenessProbe,omitempty" xml:"LivenessProbe,omitempty" type:"Struct"`
	// The memory size of the container. Unit: GiB.
	//
	// example:
	//
	// 2.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the container.
	//
	// example:
	//
	// nginx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The exposed port and protocol of the container.
	Ports []*DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The previous state of the container.
	PreviousState *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState `json:"PreviousState,omitempty" xml:"PreviousState,omitempty" type:"Struct"`
	// The readiness probe that is used to check whether the container is ready to serve a request.
	ReadinessProbe *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe `json:"ReadinessProbe,omitempty" xml:"ReadinessProbe,omitempty" type:"Struct"`
	// Indicates whether the container passed the readiness probe.
	//
	// example:
	//
	// true
	Ready *bool `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// The number of times that the container restarted.
	//
	// example:
	//
	// 0
	RestartCount *int32 `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	// The security context of the elastic container instance.
	SecurityContext *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	// Indicates whether the container allocates buffer resources to standard input streams when the container is running. If you do not specify this parameter, an end-of-file (EOF) error may occur when standard input streams in the container are read. Default value: false.
	//
	// example:
	//
	// true
	Stdin *bool `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	// Indicates whether standard input streams are disconnected after a client is disconnected. If Stdin is set to true, standard input streams remain connected among multiple sessions. If StdinOnce is set to true, standard input streams are connected after the container is started, and remain idle until a client is connected to receive data. After the client is disconnected, streams are also disconnected, and remain disconnected until the container restarts.
	//
	// example:
	//
	// true
	StdinOnce *bool `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	// Indicates whether interaction is enabled. Default value: false. If the value of the Command parameter is `/bin/bash`, the value of this parameter is true.
	//
	// example:
	//
	// false
	Tty *bool `json:"Tty,omitempty" xml:"Tty,omitempty"`
	// Information about the mounted volumes.
	VolumeMounts []*DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts `json:"VolumeMounts,omitempty" xml:"VolumeMounts,omitempty" type:"Repeated"`
	// The working directory of the container.
	//
	// example:
	//
	// /usr/local/nginx
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
	//
	// example:
	//
	// working
	DetailStatus *string `json:"DetailStatus,omitempty" xml:"DetailStatus,omitempty"`
	// The exit code of the container.
	//
	// example:
	//
	// 0
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time when the container stopped running.
	//
	// example:
	//
	// 2018-08-02T15:00:00Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The message about the container status.
	//
	// example:
	//
	// Back-off 5m0s restarting failed container=test pod=test_eci-xxx(xxx)
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason why the container is in this status.
	//
	// example:
	//
	// Started
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The code of the container status.
	//
	// example:
	//
	// 1
	Signal *int32 `json:"Signal,omitempty" xml:"Signal,omitempty"`
	// The time when the container started to run.
	//
	// example:
	//
	// 2019-12-19T12:05:02Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The container status. Valid values:
	//
	// 	- Waiting
	//
	// 	- Running
	//
	// 	- Terminated
	//
	// example:
	//
	// Terminated
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
	//
	// example:
	//
	// PATH
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable.
	//
	// example:
	//
	// /usr/bin/
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The source of the environment variable value. This parameter has a value only when the Value parameter is left empty.
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
	//
	// example:
	//
	// status.podIP
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
	// The commands that are used to check the containers.
	Execs []*string `json:"Execs,omitempty" xml:"Execs,omitempty" type:"Repeated"`
	// The minimum number of consecutive failures that must occur for the check to be considered failed. Default value: 3.
	//
	// example:
	//
	// 3
	FailureThreshold *int32 `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	// The HTTP GET method that is used to check the container.
	HttpGet *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	// The number of seconds between the time when the startup of the container ends and the time when the probe starts.
	//
	// example:
	//
	// 10
	InitialDelaySeconds *int32 `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	// The interval at which the health check is performed. Default value: 10. Minimum value: 1. Unit: seconds.
	//
	// example:
	//
	// 5
	PeriodSeconds *int32 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	// The minimum number of consecutive successes that must occur for the check to be considered successful. Default value: 1. The value must be 1.
	//
	// example:
	//
	// 1
	SuccessThreshold *int32 `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	// The TCP socket method that is used to check the container.
	TcpSocket *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
	// The timeout period of the check. Default value: 1. Minimum value: 1. Unit: seconds.
	//
	// example:
	//
	// 10
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
	// The path to which the system sends an HTTP GET request for a health check.
	//
	// example:
	//
	// /usr/nginx/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The port to which the system sends an HTTP GET request for a health check.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type supported by the method. Valid values: HTTP and HTTPS.
	//
	// example:
	//
	// HTTP
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
	//
	// example:
	//
	// 1.2.3.4
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port number.
	//
	// example:
	//
	// 80
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
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// TCP
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
	//
	// example:
	//
	// working
	DetailStatus *string `json:"DetailStatus,omitempty" xml:"DetailStatus,omitempty"`
	// The exit code of the container.
	//
	// example:
	//
	// 0
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time when the container stopped running.
	//
	// example:
	//
	// 2018-08-02T15:00:00Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The message about the container status.
	//
	// example:
	//
	// Back-off 5m0s restarting failed container=test pod=test_eci-xxx(xxx)
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason why the container is in this status.
	//
	// example:
	//
	// Completed
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The code of the container status.
	//
	// example:
	//
	// 1
	Signal *int32 `json:"Signal,omitempty" xml:"Signal,omitempty"`
	// The time when the container started to run.
	//
	// example:
	//
	// 2018-08-02T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The container status. Valid values:
	//
	// 	- Waiting: The container is being started.
	//
	// 	- Running: The container is running.
	//
	// 	- Terminated: The container stops running.
	//
	// example:
	//
	// Terminated
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
	// The commands that are run in the container when you use a CLI to perform health checks.
	Execs []*string `json:"Execs,omitempty" xml:"Execs,omitempty" type:"Repeated"`
	// The minimum number of consecutive failures that must occur for the check to be considered failed. Default value: 3.
	//
	// example:
	//
	// 3
	FailureThreshold *int32 `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	// The HTTP GET method that is used to check the container.
	HttpGet *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	// The number of seconds between the time when the startup of the container ends and the time when the probe starts.
	//
	// example:
	//
	// 5
	InitialDelaySeconds *int32 `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	// The interval at which the health check is performed. Default value: 10. Minimum value: 1. Unit: seconds.
	//
	// example:
	//
	// 1
	PeriodSeconds *int32 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	// The minimum number of consecutive successes that must occur for the check to be considered successful. Default value: 1. The value must be 1.
	//
	// example:
	//
	// 1
	SuccessThreshold *int32 `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	// The TCP socket method that is used to check the container.
	TcpSocket *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
	// The timeout period of the check. Default value: 1. Minimum value: 1. Unit: seconds.
	//
	// example:
	//
	// 1
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
	// The path to which the system sends an HTTP GET request for a health check.
	//
	// example:
	//
	// /usr/local
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The port to which the system sends an HTTP GET request for a health check.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type supported by the method. Valid values: HTTP and HTTPS.
	//
	// example:
	//
	// HTTP
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
	// The hostname.
	//
	// example:
	//
	// 1.2.3.4
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port number.
	//
	// example:
	//
	// 8888
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
	// Indicates whether permissions on the root file system are read-only. Valid value: true.
	//
	// example:
	//
	// true
	ReadOnlyRootFilesystem *bool `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	// The user ID (UID) that is used to run the container.
	//
	// example:
	//
	// 1000
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
	// The permissions specific to the process in the container.
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
	// The directory of the volume that is mounted to the container. The data in this directory is overwritten by the data on the volume.
	//
	// example:
	//
	// /usr/share/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation setting of the volume. Mount propagation allows volumes that are mounted on one container to be shared with other containers in the same pod, or even with other pods on the same node. Valid values:
	//
	// 	- None: The volume mount does not receive subsequent mounts that are performed on this volume or on the subdirectories of this volume.
	//
	// 	- HostToCotainer: The volume mount receives subsequent mounts that are performed on this volume or on the subdirectories of this volume.
	//
	// 	- Bidirectional: This value is similar to HostToCotainer. The volume mount receives subsequent mounts that are performed on this volume or on the subdirectories of this volume. In addition, all volume mounts that are performed on the container are propagated back to the host and all containers of all pods that use the same volume.
	//
	// example:
	//
	// None
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The volume name.
	//
	// example:
	//
	// test-empty
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the volume is read-only.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The subdirectory of the volume. You can use this parameter to mount the same volume to different subdirectories of the container.
	//
	// example:
	//
	// /usr/sub
	SubPath *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
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

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) SetSubPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts {
	s.SubPath = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig struct {
	// The IP addresses of DNS servers.
	NameServers []*string `json:"NameServers,omitempty" xml:"NameServers,omitempty" type:"Repeated"`
	// The options. Each option is a name-value pair. The value in the name-value pair is optional.
	Options []*DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
	// The search domains of DNS servers.
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
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The variable value of the option.
	//
	// example:
	//
	// value
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
	// sysctl parameters.
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
	//
	// example:
	//
	// kernel.msgmax
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the Sysctl parameter.
	//
	// example:
	//
	// 65536
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
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The start time of the event.
	//
	// example:
	//
	// 2018-08-02T15:00:00Z
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	// The end time of the event.
	//
	// example:
	//
	// 2018-08-02T15:00:00Z
	LastTimestamp *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The event message.
	//
	// example:
	//
	// Started container
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The category to which the event belongs.
	//
	// example:
	//
	// test-xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The event name.
	//
	// example:
	//
	// Created
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The type of the event. Valid values:
	//
	// 	- Normal
	//
	// 	- Warning
	//
	// example:
	//
	// Normal
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
	// The information about the hosts.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// The IP address of the instance.
	//
	// example:
	//
	// 192.0.XX.XX
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
	// The arguments that are passed to the startup commands of the container.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The startup commands of the containers.
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// The number of vCPUs that are allocated to the container.
	//
	// example:
	//
	// 1.0
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The current container status.
	CurrentState *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState `json:"CurrentState,omitempty" xml:"CurrentState,omitempty" type:"Struct"`
	// The environment variables of the init container.
	EnvironmentVars []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The number of GPUs.
	//
	// example:
	//
	// 1
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The image of the container.
	//
	// example:
	//
	// nginx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image pulling policy. Valid values:
	//
	// 	- Always: Each time the instance is updated, image pulling is performed.
	//
	// 	- IfNotPresent: On-premises images are preferentially used. If no on-premises images are available, image pulling is performed.
	//
	// 	- Never: On-premises images are always used. Image pulling is not performed.
	//
	// example:
	//
	// Always
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The memory size of the init container. Unit: GiB.
	//
	// example:
	//
	// 2.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the init container.
	//
	// example:
	//
	// Init-xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The exposed port and protocol of the container.
	Ports []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The previous state of the container.
	PreviousState *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState `json:"PreviousState,omitempty" xml:"PreviousState,omitempty" type:"Struct"`
	// Indicates whether the container passed the readiness probe.
	//
	// example:
	//
	// true
	Ready *bool `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// The number of times that the container restarted.
	//
	// example:
	//
	// 5
	RestartCount *int32 `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	// The security context of the container.
	SecurityContext *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	// The information about the volumes that are mounted to the init container.
	VolumeMounts []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts `json:"VolumeMounts,omitempty" xml:"VolumeMounts,omitempty" type:"Repeated"`
	// The working directory of the container.
	//
	// example:
	//
	// /usr/test
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
	//
	// example:
	//
	// working
	DetailStatus *string `json:"DetailStatus,omitempty" xml:"DetailStatus,omitempty"`
	// The exit code of the container.
	//
	// example:
	//
	// 0
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time when the container stopped running.
	//
	// example:
	//
	// 2018-08-02T15:00:00Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The event message.
	//
	// example:
	//
	// Created container
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason why the container is in this status.
	//
	// example:
	//
	// Started
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The code of the container status.
	//
	// example:
	//
	// 1
	Signal *int32 `json:"Signal,omitempty" xml:"Signal,omitempty"`
	// The time when the container started to run.
	//
	// example:
	//
	// 2018-08-02T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The container status. Valid values:
	//
	// 	- Waiting
	//
	// 	- Running
	//
	// 	- Terminated
	//
	// example:
	//
	// Running
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
	//
	// example:
	//
	// PATH
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable.
	//
	// example:
	//
	// /usr/local/bin
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The source of the environment variable value. This parameter has a value only when the Value parameter is left empty.
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
	// The specified fields.
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
	// The path of the field. Only `status.podIP` is supported.
	//
	// example:
	//
	// status.podIP
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
	//
	// example:
	//
	// 8888
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// TCP
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
	//
	// example:
	//
	// working
	DetailStatus *string `json:"DetailStatus,omitempty" xml:"DetailStatus,omitempty"`
	// The exit code of the container.
	//
	// example:
	//
	// 0
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time when the container stopped running.
	//
	// example:
	//
	// 2019-12-24T05:35:44Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The message about the container status.
	//
	// example:
	//
	// Started container
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason why the container is in this status.
	//
	// example:
	//
	// completed
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The code of the container status.
	//
	// example:
	//
	// 1
	Signal *int32 `json:"Signal,omitempty" xml:"Signal,omitempty"`
	// The time when the container started to run.
	//
	// example:
	//
	// 2019-12-24T05:35:44Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The container status. Valid values: Waiting, Running, and Terminated.
	//
	// example:
	//
	// Terminated
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
	// Indicates whether permissions on the root file system are read-only.
	//
	// example:
	//
	// true
	ReadOnlyRootFilesystem *bool `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	// The UID that is used to run the entry point of the container process.
	//
	// example:
	//
	// 1557
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
	// The directory of the volume that is mounted to the container. The data in this directory is overwritten by the data on the volume.
	//
	// example:
	//
	// /usr/share/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation setting of the volume. Mount propagation allows volumes that are mounted on one container to be shared with other containers in the same pod, or even with other pods on the same node. Valid values:
	//
	// 	- None: The volume mount does not receive subsequent mounts that are performed on this volume or on the subdirectories of this volume.
	//
	// 	- HostToCotainer: The volume mount receives subsequent mounts that are performed on this volume or on the subdirectories of this volume.
	//
	// 	- Bidirectional: This value is similar to HostToCotainer. The volume mount receives subsequent mounts that are performed on this volume or on the subdirectories of this volume. In addition, all volume mounts that are performed on the container are propagated back to the host and all containers of all pods that use the same volume.
	//
	// example:
	//
	// None
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The name of the volume. The value of this parameter is the same as the name of the volume that you selected when you purchased the container.
	//
	// example:
	//
	// test-empty
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the volume is read-only.
	//
	// example:
	//
	// false
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
	//
	// example:
	//
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test-xh
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
	// The path of the ConfigFile volume.
	ConfigFileVolumeConfigFileToPaths []*DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths `json:"ConfigFileVolumeConfigFileToPaths,omitempty" xml:"ConfigFileVolumeConfigFileToPaths,omitempty" type:"Repeated"`
	// The ID of the disk when you set Type to DiskVolume.
	//
	// example:
	//
	// d-xxxx
	DiskVolumeDiskId *string `json:"DiskVolumeDiskId,omitempty" xml:"DiskVolumeDiskId,omitempty"`
	// The file system type of the disk volume.
	//
	// example:
	//
	// xfs
	DiskVolumeFsType *string `json:"DiskVolumeFsType,omitempty" xml:"DiskVolumeFsType,omitempty"`
	// The storage media for the emptyDir volume. This parameter is empty by default, indicating that the node file system is used as the storage media. Valid values:
	//
	// 	- Memory: Memory is used as the storage media.
	//
	// 	- LocalRaid0: Local disks are formed into RAID 0. This value is valid only if an elastic container instance that has local disks mounted is created. For more information, see [Create an elastic container instance that has local disks mounted](https://help.aliyun.com/document_detail/114664.html).
	//
	// example:
	//
	// Memory
	EmptyDirVolumeMedium *string `json:"EmptyDirVolumeMedium,omitempty" xml:"EmptyDirVolumeMedium,omitempty"`
	// The storage size of the emptyDir volume.
	//
	// example:
	//
	// 256Mi
	EmptyDirVolumeSizeLimit *string `json:"EmptyDirVolumeSizeLimit,omitempty" xml:"EmptyDirVolumeSizeLimit,omitempty"`
	// The name of the driver when you set the Type parameter to FlexVolume.
	//
	// example:
	//
	// flexvolume
	FlexVolumeDriver *string `json:"FlexVolumeDriver,omitempty" xml:"FlexVolumeDriver,omitempty"`
	// The file system type when you set the Type parameter to FlexVolume. The default value varies based on the script of the FlexVolume plug-in.
	//
	// example:
	//
	// ntfs
	FlexVolumeFsType *string `json:"FlexVolumeFsType,omitempty" xml:"FlexVolumeFsType,omitempty"`
	// The options when you set the Type parameter to FlexVolume.
	//
	// example:
	//
	// [nolock,tcp,noresvport]
	FlexVolumeOptions *string `json:"FlexVolumeOptions,omitempty" xml:"FlexVolumeOptions,omitempty"`
	// The path of the NFS volume.
	//
	// example:
	//
	// /eci
	NFSVolumePath *string `json:"NFSVolumePath,omitempty" xml:"NFSVolumePath,omitempty"`
	// Indicates whether the NFS volume is read-only.
	//
	// example:
	//
	// false
	NFSVolumeReadOnly *bool `json:"NFSVolumeReadOnly,omitempty" xml:"NFSVolumeReadOnly,omitempty"`
	// The address of the NFS server.
	//
	// example:
	//
	// 0eafxxxxx-xxxxx.cn-hangzhou.nas.aliyuncs.com
	NFSVolumeServer *string `json:"NFSVolumeServer,omitempty" xml:"NFSVolumeServer,omitempty"`
	// The volume name.
	//
	// example:
	//
	// test-empty
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the volume. Valid values:
	//
	// 	- EmptyDirVolume
	//
	// 	- NFSVolume
	//
	// 	- ConfigFileVolume
	//
	// 	- FlexVolume
	//
	// example:
	//
	// EmptyDirVolume
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

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetEmptyDirVolumeMedium(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.EmptyDirVolumeMedium = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetEmptyDirVolumeSizeLimit(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.EmptyDirVolumeSizeLimit = &v
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
	//
	// example:
	//
	// hello world
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The relative path of the ConfigFile volume.
	//
	// example:
	//
	// /usr/bin/
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// eci-uf6fonnghi50v6nq****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The name of the container.
	//
	// This parameter is required.
	//
	// example:
	//
	// nginx
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// Specifies whether to query the logs of the previous container if the container exits and restarts. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	LastTime *bool `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The limit on the total size of logs. Unit: bytes. Valid values: 1 to 1048576(1 MB).
	//
	// example:
	//
	// 123798
	LimitBytes   *int64  `json:"LimitBytes,omitempty" xml:"LimitBytes,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the elastic container instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// A relative time in seconds before the current time from which to show logs. Examples: 10, 20, and 30.
	//
	// example:
	//
	// 10
	SinceSeconds *int32 `json:"SinceSeconds,omitempty" xml:"SinceSeconds,omitempty"`
	// The beginning of the time range to query. Specify the time in the RFC 3339 standard. The time must be in UTC.
	//
	// example:
	//
	// 2018-08-02T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The number of the most recent log entries that you want to query. Default value: 500. Maximum value: 2000.\\
	//
	// A maximum of 1 MB log data can be returned.
	//
	// example:
	//
	// 500
	Tail *int32 `json:"Tail,omitempty" xml:"Tail,omitempty"`
	// Specifies whether to return the timestamps of logs. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
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
	// The container name.
	//
	// example:
	//
	// nginx
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The content of the log.
	//
	// example:
	//
	// {}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CA850ADA-****-4AC8-****-5B5990EAB7D0
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeDataCachesRequest struct {
	// The bucket that stores the data cache. Default value: default.
	//
	// example:
	//
	// default
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The data cache IDs.
	DataCacheId []*string `json:"DataCacheId,omitempty" xml:"DataCacheId,omitempty" type:"Repeated"`
	// The maximum entries of query results that are allowed to be displayed. Valid values: 1 to 20. Default value: 20.
	//
	// example:
	//
	// 20
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The query token. Set the value to the NextToken value that is returned in the previous call.
	//
	// example:
	//
	// d78f2dd8-5979-42fe-****-b16db43be5bc
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The virtual host (vHost) directory in which the data cache resides.
	//
	// example:
	//
	// /data/models/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The region ID of the data caches that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the data cache belongs.
	//
	// example:
	//
	// rg-acfmzw2jz2z****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags that are attached to the data cache.
	Tag []*DescribeDataCachesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDataCachesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCachesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataCachesRequest) SetBucket(v string) *DescribeDataCachesRequest {
	s.Bucket = &v
	return s
}

func (s *DescribeDataCachesRequest) SetDataCacheId(v []*string) *DescribeDataCachesRequest {
	s.DataCacheId = v
	return s
}

func (s *DescribeDataCachesRequest) SetLimit(v int32) *DescribeDataCachesRequest {
	s.Limit = &v
	return s
}

func (s *DescribeDataCachesRequest) SetNextToken(v string) *DescribeDataCachesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDataCachesRequest) SetOwnerAccount(v string) *DescribeDataCachesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDataCachesRequest) SetOwnerId(v int64) *DescribeDataCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDataCachesRequest) SetPath(v string) *DescribeDataCachesRequest {
	s.Path = &v
	return s
}

func (s *DescribeDataCachesRequest) SetRegionId(v string) *DescribeDataCachesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDataCachesRequest) SetResourceGroupId(v string) *DescribeDataCachesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDataCachesRequest) SetResourceOwnerAccount(v string) *DescribeDataCachesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDataCachesRequest) SetResourceOwnerId(v int64) *DescribeDataCachesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDataCachesRequest) SetTag(v []*DescribeDataCachesRequestTag) *DescribeDataCachesRequest {
	s.Tag = v
	return s
}

type DescribeDataCachesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// testkey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testvalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDataCachesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCachesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDataCachesRequestTag) SetKey(v string) *DescribeDataCachesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDataCachesRequestTag) SetValue(v string) *DescribeDataCachesRequestTag {
	s.Value = &v
	return s
}

type DescribeDataCachesResponseBody struct {
	// The information about the data caches.
	DataCaches []*DescribeDataCachesResponseBodyDataCaches `json:"DataCaches,omitempty" xml:"DataCaches,omitempty" type:"Repeated"`
	// The query token. Set the value to the NextToken value that is returned in the previous call.
	//
	// example:
	//
	// d78f2dd8-5979-42fe-xaee-b16db43be5bc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B8756BA0-6452-419C-9727-37A6209C85E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataCachesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCachesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataCachesResponseBody) SetDataCaches(v []*DescribeDataCachesResponseBodyDataCaches) *DescribeDataCachesResponseBody {
	s.DataCaches = v
	return s
}

func (s *DescribeDataCachesResponseBody) SetNextToken(v string) *DescribeDataCachesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDataCachesResponseBody) SetRequestId(v string) *DescribeDataCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataCachesResponseBody) SetTotalCount(v int32) *DescribeDataCachesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDataCachesResponseBodyDataCaches struct {
	// The bucket that stores the data cache.
	//
	// example:
	//
	// default
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The ID of the elastic container instance that was generated when the data cache was created.
	//
	// example:
	//
	// eci-8vb1y2w1dv7zeirn****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The time when the data cache was created.
	//
	// example:
	//
	// 2023-06-16T02:43Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the data cache.
	//
	// example:
	//
	// edc-bp15l4vvys94oo******
	DataCacheId *string `json:"DataCacheId,omitempty" xml:"DataCacheId,omitempty"`
	// The information about the data source.
	DataSource *DescribeDataCachesResponseBodyDataCachesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The events.
	Events []*DescribeDataCachesResponseBodyDataCachesEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The time when the data cache expires.
	//
	// example:
	//
	// 2023-06-26T02:43Z
	ExpireDateTime *string `json:"ExpireDateTime,omitempty" xml:"ExpireDateTime,omitempty"`
	// The ID of the on-premises snapshot.
	//
	// example:
	//
	// s-bp12w3v37sit96t6****
	FlashSnapshotId *string `json:"FlashSnapshotId,omitempty" xml:"FlashSnapshotId,omitempty"`
	// The time when the data cache was last matched.
	//
	// example:
	//
	// 2023-06-18T02:43Z
	LastMatchedTime *string `json:"LastMatchedTime,omitempty" xml:"LastMatchedTime,omitempty"`
	// The name of the data cache.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The directory in which the virtual host of the data cache resides.
	//
	// example:
	//
	// /data/models/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The creation progress of the data cache.
	//
	// example:
	//
	// 100%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2hlt3ux4****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The size of the data cache. Unit: GiB.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The snapshot ID.
	//
	// example:
	//
	// s-2zec5oj8e1yhxijt****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The status of the data cache. Valid values:
	//
	// 	- Loading: The data cache is loading data.
	//
	// 	- Creating: The data cache is being created.
	//
	// 	- Available: The data cache is created.
	//
	// 	- Failed: The data cache failed to be created.
	//
	// 	- Updating: The data cache is being updated.
	//
	// 	- UpdateFailed: The data cache failed to be updated.
	//
	// If the data cache is in the Available state, the data cache can be used.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are attached to the data cache.
	Tags []*DescribeDataCachesResponseBodyDataCachesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeDataCachesResponseBodyDataCaches) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCachesResponseBodyDataCaches) GoString() string {
	return s.String()
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetBucket(v string) *DescribeDataCachesResponseBodyDataCaches {
	s.Bucket = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetContainerGroupId(v string) *DescribeDataCachesResponseBodyDataCaches {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetCreationTime(v string) *DescribeDataCachesResponseBodyDataCaches {
	s.CreationTime = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetDataCacheId(v string) *DescribeDataCachesResponseBodyDataCaches {
	s.DataCacheId = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetDataSource(v *DescribeDataCachesResponseBodyDataCachesDataSource) *DescribeDataCachesResponseBodyDataCaches {
	s.DataSource = v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetEvents(v []*DescribeDataCachesResponseBodyDataCachesEvents) *DescribeDataCachesResponseBodyDataCaches {
	s.Events = v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetExpireDateTime(v string) *DescribeDataCachesResponseBodyDataCaches {
	s.ExpireDateTime = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetFlashSnapshotId(v string) *DescribeDataCachesResponseBodyDataCaches {
	s.FlashSnapshotId = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetLastMatchedTime(v string) *DescribeDataCachesResponseBodyDataCaches {
	s.LastMatchedTime = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetName(v string) *DescribeDataCachesResponseBodyDataCaches {
	s.Name = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetPath(v string) *DescribeDataCachesResponseBodyDataCaches {
	s.Path = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetProgress(v string) *DescribeDataCachesResponseBodyDataCaches {
	s.Progress = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetRegionId(v string) *DescribeDataCachesResponseBodyDataCaches {
	s.RegionId = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetResourceGroupId(v string) *DescribeDataCachesResponseBodyDataCaches {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetSize(v int32) *DescribeDataCachesResponseBodyDataCaches {
	s.Size = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetSnapshotId(v string) *DescribeDataCachesResponseBodyDataCaches {
	s.SnapshotId = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetStatus(v string) *DescribeDataCachesResponseBodyDataCaches {
	s.Status = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCaches) SetTags(v []*DescribeDataCachesResponseBodyDataCachesTags) *DescribeDataCachesResponseBodyDataCaches {
	s.Tags = v
	return s
}

type DescribeDataCachesResponseBodyDataCachesDataSource struct {
	// The parameters that are configured for the data source.
	//
	// example:
	//
	// {
	//
	//       "options": {
	//
	//             "path": "/",
	//
	//             "server": "0ce1f4****-or***.cn-hangzhou.nas.aliyuncs.com",
	//
	//             "vers": "3",
	//
	//             "options": "nolock,proto=tcp,rsize=1048576,wsize=1048576,hard,timeo=600,retrans=2,noresvport"
	//
	//       },
	//
	//       "type": "NAS"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- NAS
	//
	// 	- OSS
	//
	// 	- URL
	//
	// 	- SNAPSHOT
	//
	// example:
	//
	// URL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDataCachesResponseBodyDataCachesDataSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCachesResponseBodyDataCachesDataSource) GoString() string {
	return s.String()
}

func (s *DescribeDataCachesResponseBodyDataCachesDataSource) SetOptions(v string) *DescribeDataCachesResponseBodyDataCachesDataSource {
	s.Options = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCachesDataSource) SetType(v string) *DescribeDataCachesResponseBodyDataCachesDataSource {
	s.Type = &v
	return s
}

type DescribeDataCachesResponseBodyDataCachesEvents struct {
	// The number of times that the event occurred.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The time when the event started.
	//
	// example:
	//
	// 2021-06-16T02:14:48Z
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	// The time when the event ended.
	//
	// example:
	//
	// 2021-06-16T02:24:48Z
	LastTimestamp *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The message about the event.
	//
	// example:
	//
	// Successfully check data cache resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The event name.
	//
	// example:
	//
	// datacache
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The reason for the transition into the current status of the event.
	//
	// example:
	//
	// Started
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The type of the event. Valid values:
	//
	// 	- Normal
	//
	// 	- Warning
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDataCachesResponseBodyDataCachesEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCachesResponseBodyDataCachesEvents) GoString() string {
	return s.String()
}

func (s *DescribeDataCachesResponseBodyDataCachesEvents) SetCount(v int32) *DescribeDataCachesResponseBodyDataCachesEvents {
	s.Count = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCachesEvents) SetFirstTimestamp(v string) *DescribeDataCachesResponseBodyDataCachesEvents {
	s.FirstTimestamp = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCachesEvents) SetLastTimestamp(v string) *DescribeDataCachesResponseBodyDataCachesEvents {
	s.LastTimestamp = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCachesEvents) SetMessage(v string) *DescribeDataCachesResponseBodyDataCachesEvents {
	s.Message = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCachesEvents) SetName(v string) *DescribeDataCachesResponseBodyDataCachesEvents {
	s.Name = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCachesEvents) SetReason(v string) *DescribeDataCachesResponseBodyDataCachesEvents {
	s.Reason = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCachesEvents) SetType(v string) *DescribeDataCachesResponseBodyDataCachesEvents {
	s.Type = &v
	return s
}

type DescribeDataCachesResponseBodyDataCachesTags struct {
	// The tag key.
	//
	// example:
	//
	// testkey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testvalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDataCachesResponseBodyDataCachesTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCachesResponseBodyDataCachesTags) GoString() string {
	return s.String()
}

func (s *DescribeDataCachesResponseBodyDataCachesTags) SetKey(v string) *DescribeDataCachesResponseBodyDataCachesTags {
	s.Key = &v
	return s
}

func (s *DescribeDataCachesResponseBodyDataCachesTags) SetValue(v string) *DescribeDataCachesResponseBodyDataCachesTags {
	s.Value = &v
	return s
}

type DescribeDataCachesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataCachesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataCachesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCachesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataCachesResponse) SetHeaders(v map[string]*string) *DescribeDataCachesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataCachesResponse) SetStatusCode(v int32) *DescribeDataCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataCachesResponse) SetBody(v *DescribeDataCachesResponseBody) *DescribeDataCachesResponse {
	s.Body = v
	return s
}

type DescribeImageCachesRequest struct {
	// The container images.
	//
	// example:
	//
	// nginx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The IDs of the image caches.
	//
	// example:
	//
	// imc-bp195erqe9o2pb09****
	ImageCacheId *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	// The names of the image caches.
	//
	// example:
	//
	// testcache
	ImageCacheName *string `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	// Specifies whether the image layers of the image caches must contain all image layers of the container image.\\
	//
	// If you configure MatchImage, you can configure this parameter to further filter query results.
	//
	// example:
	//
	// true
	ImageFullMatch *bool `json:"ImageFullMatch,omitempty" xml:"ImageFullMatch,omitempty"`
	// The quantity of image caches whose image layers contain all image layers of the container image.\\
	//
	// If you configure MatchImage, you can configure this parameter to further filter query results.
	//
	// example:
	//
	// 3
	ImageMatchCountRequest *int32 `json:"ImageMatchCountRequest,omitempty" xml:"ImageMatchCountRequest,omitempty"`
	// The maximum entries of query results that are allowed to be displayed.
	//
	// example:
	//
	// 20
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The container images used to match the image caches that you want to query. You can specify a maximum of 100 container images.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/eci_open/nginx:1.15.10-perl
	MatchImage []*string `json:"MatchImage,omitempty" xml:"MatchImage,omitempty" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the image caches.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the image caches belong.
	//
	// example:
	//
	// rg-2df3isufhi38****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IDs of the snapshots that correspond to the image caches.
	//
	// example:
	//
	// s-2zec5oj8e1yhxijt****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The tags to add to the image caches.
	Tag []*DescribeImageCachesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *DescribeImageCachesRequest) SetImageFullMatch(v bool) *DescribeImageCachesRequest {
	s.ImageFullMatch = &v
	return s
}

func (s *DescribeImageCachesRequest) SetImageMatchCountRequest(v int32) *DescribeImageCachesRequest {
	s.ImageMatchCountRequest = &v
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
	// The key of tag N of the image cache.
	//
	// example:
	//
	// imc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the image cache.
	//
	// example:
	//
	// test
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
	// The information about image caches.
	ImageCaches []*DescribeImageCachesResponseBodyImageCaches `json:"ImageCaches,omitempty" xml:"ImageCaches,omitempty" type:"Repeated"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39FC2E43-3DD7-4CEF-9EF4-E4AD6E635301
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
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
	// The ID of the elastic container instance.
	//
	// example:
	//
	// eci-bp18oq3m15prd9jb****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The time when the image cache was created.
	//
	// example:
	//
	// 2021-02-09T02:24:07Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The elimination policy of the image cache. This parameter is empty by default, which indicates that the image cache is always retained.
	//
	// You can set this parameter to LRU, which indicates that the image cache can be automatically deleted. When the number of image caches reaches the quota, the system automatically deletes the image caches whose EliminationStrategy parameter is set to LRU and that are least regularly used.
	//
	// example:
	//
	// LRU
	EliminationStrategy *string `json:"EliminationStrategy,omitempty" xml:"EliminationStrategy,omitempty"`
	// The events of pulling an image to create the image cache.
	Events []*DescribeImageCachesResponseBodyImageCachesEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The time when the image cache expires.
	//
	// example:
	//
	// 2019-11-10T09:00:48Z
	ExpireDateTime *string `json:"ExpireDateTime,omitempty" xml:"ExpireDateTime,omitempty"`
	// The ID of the local snapshot. A temporary local snapshot is created if the instant image cache feature is enabled.
	//
	// example:
	//
	// s-bp12w3v37sit96t6****
	FlashSnapshotId *string `json:"FlashSnapshotId,omitempty" xml:"FlashSnapshotId,omitempty"`
	// The ID of the image cache.
	//
	// example:
	//
	// imc-bp195erqe9o2pb09****
	ImageCacheId *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	// The name of the image cache.
	//
	// example:
	//
	// imagetest
	ImageCacheName *string `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	// The size of the image cache. Unit: GiB.
	//
	// example:
	//
	// 20
	ImageCacheSize *int32 `json:"ImageCacheSize,omitempty" xml:"ImageCacheSize,omitempty"`
	// The images contained in the image cache.
	Images []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The time when the image cache was last matched.
	//
	// example:
	//
	// 2021-08-18T03:48:10Z
	LastMatchedTime *string `json:"LastMatchedTime,omitempty" xml:"LastMatchedTime,omitempty"`
	// The progress of creating the snapshot that is used to create the image cache.
	//
	// >  If the instant image cache feature is enabled, the system creates a temporary local snapshot and then a regular snapshot. In this case, this parameter indicates the progress of creating the regular snapshot.
	//
	// example:
	//
	// 100%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The region ID of the image cache.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the image cache belongs.
	//
	// example:
	//
	// rg-2df3isufhi38****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the snapshot that corresponds to the image cache.
	//
	// example:
	//
	// s-2zec5oj8e1yhxijt****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The status of the image cache. Valid values:
	//
	// 	- Preparing: The image cache is being prepared.
	//
	// 	- Creating: The image is being created.
	//
	// 	- Ready: The image cache is created.
	//
	// 	- Failed: The image cache failed to be created.
	//
	// 	- Updating: The image cache is being updated.
	//
	// 	- UpdateFailed: The image cache failed to be updated.
	//
	// The image cache is ready for use when it is in the Ready state.
	//
	// example:
	//
	// Ready
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
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The time when the event started.
	//
	// example:
	//
	// 2021-02-09T02:24:48Z
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	// The time when the event ended.
	//
	// example:
	//
	// 2021-02-09T02:24:48Z
	LastTimestamp *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The message about the event.
	//
	// example:
	//
	// Successfully check image cache resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the event.
	//
	// example:
	//
	// imagetest.1661f31f851a****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cause of the event.
	//
	// example:
	//
	// Started
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The type of the event. Valid values:
	//
	// example:
	//
	// Normal
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

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetReason(v string) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.Reason = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetType(v string) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.Type = &v
	return s
}

type DescribeImageCachesResponseBodyImageCachesTags struct {
	// The tag key.
	//
	// example:
	//
	// imc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageCachesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the elastic container instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// eci-bp1dvysdafbh00t7****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The type of the O\\&M task. Valid values:
	//
	// 	- coredump
	//
	// 	- tcpdump
	//
	// This parameter is required.
	//
	// example:
	//
	// coredump
	OpsType      *string `json:"OpsType,omitempty" xml:"OpsType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	// The details of the queried O\\&M tasks.
	Records []*DescribeInstanceOpsRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 89945DD3-9072-47D0-A318-353284CFC7B3
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
	// The time when the O\\&M task was created.
	//
	// example:
	//
	// 2021-12-29T15:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the O\\&M task expires.
	//
	// example:
	//
	// 2099-12-29T15:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The status of the O\\&M task.
	//
	// - Ready
	//
	// - Failed
	//
	// - Expired
	//
	// - Closed
	//
	// - Success
	//
	// example:
	//
	// Ready
	OpsStatus *string `json:"OpsStatus,omitempty" xml:"OpsStatus,omitempty"`
	// The type of the O\\&M task.
	//
	// example:
	//
	// coredump
	OpsType *string `json:"OpsType,omitempty" xml:"OpsType,omitempty"`
	// The content of the O\\&M result. The value is the download URL of the files that are generated for the O\\&M task. This parameter is returned only when the value of the OpsStatus parameter is Success.
	//
	// example:
	//
	// https://******
	ResultContent *string `json:"ResultContent,omitempty" xml:"ResultContent,omitempty"`
	// The type of the O\\&M result. This parameter is returned only when the value of the OpsStatus parameter is Success.\\
	//
	// The only value of the parameter is Oss. The value indicates that the files generated for the O\\&M task are saved to Object Storage Service (OSS) buckets.
	//
	// example:
	//
	// OSS
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceOpsRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The instance ID. The value is a JSON array. You can specify up to 20 instance IDs at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["eci-2zegym1qhbmdfr1s****","eci-2ze39w5svzj5ic34****"]
	ContainerGroupIds *string `json:"ContainerGroupIds,omitempty" xml:"ContainerGroupIds,omitempty"`
	// The type of the monitoring data. Set the value to summary. This value indicates that records are returned.
	//
	// example:
	//
	// summary
	MetricType   *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the elastic container instances belong.
	//
	// example:
	//
	// rg-acfmzw2jz2z****
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
	// The monitoring data of the elastic container instances.
	MonitorDatas []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatas `json:"MonitorDatas,omitempty" xml:"MonitorDatas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B8756BA0-6452-419C-9727-37A6209C85E0
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
	//
	// example:
	//
	// eci-bp15oolgorelvhfb****
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
	//
	// example:
	//
	// 2019-12-19T02:31:54Z
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
	//
	// example:
	//
	// 2000
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The average load in the last 10 seconds.
	//
	// example:
	//
	// 0
	Load *int64 `json:"Load,omitempty" xml:"Load,omitempty"`
	// The cumulative usage of vCPUs.
	//
	// example:
	//
	// 70769883
	UsageCoreNanoSeconds *int64 `json:"UsageCoreNanoSeconds,omitempty" xml:"UsageCoreNanoSeconds,omitempty"`
	// The vCPU usage in the sampling window. Unit for the sampling window: nanoseconds.
	//
	// example:
	//
	// 0
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
	// The name.
	//
	// example:
	//
	// test1
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
	//
	// example:
	//
	// 321234
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The average load in the last 10 seconds.
	//
	// example:
	//
	// 12341
	Load *int64 `json:"Load,omitempty" xml:"Load,omitempty"`
	// The cumulative usage of vCPUs.
	//
	// example:
	//
	// 40876694
	UsageCoreNanoSeconds *int64 `json:"UsageCoreNanoSeconds,omitempty" xml:"UsageCoreNanoSeconds,omitempty"`
	// The vCPU usage in the sampling window. Unit for the sampling window: nanoseconds.
	//
	// example:
	//
	// 0
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
	//
	// example:
	//
	// 4285091840
	AvailableBytes *int64 `json:"AvailableBytes,omitempty" xml:"AvailableBytes,omitempty"`
	// The size of the cache. Unit: bytes.
	//
	// example:
	//
	// 9289728
	Cache *int64 `json:"Cache,omitempty" xml:"Cache,omitempty"`
	// The size of the resident memory set, which indicates the size of the physical memory that is actually used. Unit: bytes.
	//
	// example:
	//
	// 9289728
	Rss *int64 `json:"Rss,omitempty" xml:"Rss,omitempty"`
	// The size of the used memory. Unit: bytes.
	//
	// example:
	//
	// 9289728
	UsageBytes *int64 `json:"UsageBytes,omitempty" xml:"UsageBytes,omitempty"`
	// The usage of the working set. Unit: bytes.
	//
	// example:
	//
	// 9289728
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
	//
	// example:
	//
	// /dev/vda
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The amount of data that was read from the disk. Unit: bytes.
	//
	// example:
	//
	// 210714112
	ReadBytes *int64 `json:"ReadBytes,omitempty" xml:"ReadBytes,omitempty"`
	// This parameter is unavailable for public use.
	//
	// example:
	//
	// 123
	ReadIo *int64 `json:"ReadIo,omitempty" xml:"ReadIo,omitempty"`
	// The amount of data that was written to the disk. Unit: bytes.
	//
	// example:
	//
	// 1375109120
	WriteBytes *int64 `json:"WriteBytes,omitempty" xml:"WriteBytes,omitempty"`
	// This parameter is unavailable for public use.
	//
	// example:
	//
	// 123
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
	//
	// example:
	//
	// 35319836672
	Available *int64 `json:"Available,omitempty" xml:"Available,omitempty"`
	// The total file system space.
	//
	// example:
	//
	// 41610981376
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The name of the partition.
	//
	// example:
	//
	// /dev/root
	FsName *string `json:"FsName,omitempty" xml:"FsName,omitempty"`
	// The size of used space.
	//
	// example:
	//
	// 4368744448
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
	//
	// example:
	//
	// 4289445888
	AvailableBytes *int64 `json:"AvailableBytes,omitempty" xml:"AvailableBytes,omitempty"`
	// The size of the cache. Unit: bytes.
	//
	// example:
	//
	// 7028736
	Cache *int64 `json:"Cache,omitempty" xml:"Cache,omitempty"`
	// The size of the resident memory set, which indicates the size of the physical memory that is actually used. Unit: bytes.
	//
	// example:
	//
	// 1593344
	Rss *int64 `json:"Rss,omitempty" xml:"Rss,omitempty"`
	// The size of the used memory. Unit: bytes.
	//
	// example:
	//
	// 11153408
	UsageBytes *int64 `json:"UsageBytes,omitempty" xml:"UsageBytes,omitempty"`
	// The usage of the working set. Unit: bytes.
	//
	// example:
	//
	// 5521408
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
	//
	// example:
	//
	// eth0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The total number of bytes received.
	//
	// example:
	//
	// 505001954
	RxBytes *int64 `json:"RxBytes,omitempty" xml:"RxBytes,omitempty"`
	// The number of packets that fail to be received.
	//
	// example:
	//
	// 0
	RxDrops *int64 `json:"RxDrops,omitempty" xml:"RxDrops,omitempty"`
	// The number of receipt errors.
	//
	// example:
	//
	// 0
	RxErrors *int64 `json:"RxErrors,omitempty" xml:"RxErrors,omitempty"`
	// The total number of packets received.
	//
	// example:
	//
	// 4800583
	RxPackets *int64 `json:"RxPackets,omitempty" xml:"RxPackets,omitempty"`
	// The total number of bytes sent.
	//
	// example:
	//
	// 1381805699
	TxBytes *int64 `json:"TxBytes,omitempty" xml:"TxBytes,omitempty"`
	// The number of packets that fail to arrive at their destination.
	//
	// example:
	//
	// 0
	TxDrops *int64 `json:"TxDrops,omitempty" xml:"TxDrops,omitempty"`
	// The total number of sending errors.
	//
	// example:
	//
	// 0
	TxErrors *int64 `json:"TxErrors,omitempty" xml:"TxErrors,omitempty"`
	// The total number of packets sent.
	//
	// example:
	//
	// 5158427
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMultiContainerGroupMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
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
	// The list of regions.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 89945DD3-9072-47D0-A318-353284CFC7B3
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
	// The list of recommended zones.
	RecommendZones []*string `json:"RecommendZones,omitempty" xml:"RecommendZones,omitempty" type:"Repeated"`
	// The endpoint for the region.
	//
	// example:
	//
	// eci.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of unavailable zones.
	UnavailableZones []*string `json:"UnavailableZones,omitempty" xml:"UnavailableZones,omitempty" type:"Repeated"`
	// The queried zones.
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

func (s *DescribeRegionsResponseBodyRegions) SetUnavailableZones(v []*string) *DescribeRegionsResponseBodyRegions {
	s.UnavailableZones = v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetZones(v []*string) *DescribeRegionsResponseBodyRegions {
	s.Zones = v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency of requests?](https://help.aliyun.com/document_detail/25693.html)
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The maximum number of resources that are allowed to return for this request. Default value: 20. Maximum value: 20.
	//
	// >  The number of returned resources is less than or equal to the specified number.
	//
	// example:
	//
	// 20
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The token that determines the start point of the next query. If this parameter is empty, all results have been returned.
	//
	// You do not need to specify this parameter in the first request. From the second request, you can obtain the token from the result returned by the previous request.
	//
	// example:
	//
	// d78f2dd8-5979-42fe-****-b16db43be5bc
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the virtual nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-uf66jeqopgqa9hdn****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the virtual node. Valid values:
	//
	// 	- Pending
	//
	// 	- Ready
	//
	// 	- Failed
	//
	// example:
	//
	// Ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are bound to the virtual node.
	Tag []*DescribeVirtualNodesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The IDs of the virtual nodes. You can specify up to 20 IDs. Each ID must be a string in the JSON format.
	//
	// example:
	//
	// ["vnd-2ze960zkdqrldeaw****","vnd-3ebzcviqbwt25dsz****"]
	VirtualNodeIds *string `json:"VirtualNodeIds,omitempty" xml:"VirtualNodeIds,omitempty"`
	// The names of the virtual nodes.
	//
	// example:
	//
	// testNode
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
	//
	// example:
	//
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// example:
	//
	// test
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
	// The token that determines the start point of the next query.
	//
	// example:
	//
	// d78f2dd8-5979-42fe-****-b16db43be5bc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9C9B9917-ED22-50D5-ADE6-9FA9D58AD05F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of virtual nodes that were queried.
	//
	// example:
	//
	// 1
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
	// The time when the virtual node was created. The time follows the RFC 3339 standard and is displayed in UTC.
	//
	// example:
	//
	// 2021-09-08T15:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The events about the virtual node.
	Events []*DescribeVirtualNodesResponseBodyVirtualNodesEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The public IP address of the virtual node.
	//
	// example:
	//
	// 192.0.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the virtual node.
	//
	// example:
	//
	// 10.19.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The ID of the region in which the virtual node resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the virtual node belongs.
	//
	// example:
	//
	// rg-uf66jeqopgqa9hdn****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the virtual node. Valid values:
	//
	// 	- Pending
	//
	// 	- Ready
	//
	// 	- Failed
	//
	// example:
	//
	// Ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are bound to the virtual node.
	Tags []*DescribeVirtualNodesResponseBodyVirtualNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the virtual node.
	//
	// example:
	//
	// vnd-2ze960zkdqrldeaw****
	VirtualNodeId *string `json:"VirtualNodeId,omitempty" xml:"VirtualNodeId,omitempty"`
	// The name of the virtual node.
	//
	// example:
	//
	// testNode
	VirtualNodeName *string `json:"VirtualNodeName,omitempty" xml:"VirtualNodeName,omitempty"`
	// The ID of the security group to which the virtual node belongs.
	//
	// example:
	//
	// sg-2zeeyaaxlkq9sppl****
	VirtualNodeSecurityGroupId *string `json:"VirtualNodeSecurityGroupId,omitempty" xml:"VirtualNodeSecurityGroupId,omitempty"`
	// The ID of the vSwitch with which the virtual node is associated.
	//
	// example:
	//
	// vsw-2ze23nqzig8inprou****
	VirtualNodeVSwitchId *string `json:"VirtualNodeVSwitchId,omitempty" xml:"VirtualNodeVSwitchId,omitempty"`
	// The ID of the VPC to which the virtual node belongs.
	//
	// example:
	//
	// vpc-1vzjjflab6wvjox****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVirtualNodesResponseBodyVirtualNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualNodesResponseBodyVirtualNodes) GoString() string {
	return s.String()
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

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetRegionId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.RegionId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetResourceGroupId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.ResourceGroupId = &v
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

type DescribeVirtualNodesResponseBodyVirtualNodesEvents struct {
	// The number of events.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The time when the event started.
	//
	// example:
	//
	// 2021-09-08T02:24:48Z
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	// The time when the event ended.
	//
	// example:
	//
	// 2021-09-08T02:24:52Z
	LastTimestamp *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The message of the event.
	//
	// example:
	//
	// Successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the object to which the event belongs.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the event.
	//
	// example:
	//
	// Created
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The type of the event. Valid values:
	//
	// 	- Normal
	//
	// 	- Warning
	//
	// example:
	//
	// Normal
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
	// The key of the tag.
	//
	// example:
	//
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// test
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVirtualNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The commands to run in the container. You can specify up to 20 commands. Each command can be up to 256 characters in length.\\
	//
	// The strings must be in the JSON format. Example: `["/bin/sh", "-c", "ls -a"]`.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["/bin/sh", "-c", "ls -a"]
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The ID of the elastic container instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// eci-2zebxkiifulhl****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The name of the container.
	//
	// This parameter is required.
	//
	// example:
	//
	// nginx
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to read the commands from standard input (stdin). Default value: true.
	//
	// example:
	//
	// true
	Stdin *bool `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	// Specifies whether to immediately run the command and synchronously return the result. Default value: false.\\
	//
	// If this parameter is set to true, TTY must be set to false. Command cannot be set to `/bin/bash`.
	//
	// example:
	//
	// false
	Sync *bool `json:"Sync,omitempty" xml:"Sync,omitempty"`
	// Specifies whether to enable interaction. Default value: false.\\
	//
	// If the Command parameter is set to `/bin/bash`, set this parameter to true.
	//
	// example:
	//
	// false
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
	// The HTTP URL. You can use this URL to enter the container within 30 seconds after this operation is called. For more information, see [Use and integrate the Elastic Container Instance terminal](https://help.aliyun.com/document_detail/202846.html).
	//
	// example:
	//
	// https://eci.console.aliyun.com/terminal?param=X32a****
	HttpUrl *string `json:"HttpUrl,omitempty" xml:"HttpUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 28B583A8-313D-4383-7817-B5A9F7E0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The output of the command. This parameter is returned only if Sync is set to true.
	//
	// example:
	//
	// Hello
	SyncResponse *string `json:"SyncResponse,omitempty" xml:"SyncResponse,omitempty"`
	// The WebSocket URL. You can use this URL to establish a WebSocket connection with the container.
	//
	// example:
	//
	// wss://eci-cn-shanghai.aliyun.com/exec/?s=ktHPx****
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecContainerCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListTagResourcesRequest struct {
	// The maximum number of entries to return.
	//
	// example:
	//
	// 20
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- ContainerGroup: elastic container instance
	//
	// 	- ImageCache: image cache
	//
	// 	- DataCache: data cache
	//
	// 	- VirtualNode: virtual node
	//
	// This parameter is required.
	//
	// example:
	//
	// ContainerGroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags that are added to the resource.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetLimit(v string) *ListTagResourcesRequest {
	s.Limit = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
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

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerId(v int64) *ListTagResourcesRequest {
	s.ResourceOwnerId = &v
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
	// The tag key.
	//
	// example:
	//
	// testkey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testvalue
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
	// The returned pagination token which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1340C38D-6189-54D1-86F6-7D5ECF3E0088
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried resources.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// The resource ID.
	//
	// example:
	//
	// eci-2zelg8vwnlzdhf8hv****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- ContainerGroup: elastic container instance
	//
	// 	- ImageCache: image cache
	//
	// 	- DataCache: data cache
	//
	// 	- VirtualNode: virtual node
	//
	// example:
	//
	// ContainerGroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag that is added to the resource.
	//
	// example:
	//
	// testkey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag that is added to the resource.
	//
	// example:
	//
	// testvalue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListUsageRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	// The information about the used amounts and upper limits of privileges and quotas that you have in the specified region. The information contains the following items:
	//
	// 	- UsedCpu: the number of existing vCPUs.
	//
	// 	- MaxCpu: the upper limit of vCPUs.
	//
	// 	- MaxImageCacheCount: the upper limit of manually created image caches.
	//
	// 	- UsedImageCacheCount: the number of existing image caches that are manually created.
	//
	// 	- MaxAutoImageCacheCount: the upper limit of automatically created image caches.
	//
	// 	- UsedAutoImageCacheCount: the number of existing image caches that are automatically created.
	//
	// 	- MaxDataCacheCount: the upper limit of DataCaches.
	//
	// 	- UsedDataCacheCount: the number of existing DataCaches.
	//
	// example:
	//
	// {"UsedCpu": 11,"MaxCpu": 1000,"MaxImageCacheCount": 50,"UsedImageCacheCount": 0}
	Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 89164E78-FC82-4684-BE97-DCDD85D26546
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the elastic container instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// eci-2zelg8vwnlzdhf8hv****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The size of the volume after the volume is scaled up. Unit: GiB. Valid values:
	//
	// 	- Ultra disk (cloud_efficiency): 20 to 32768
	//
	// 	- Standard SSD (cloud_ssd): 20 to 32768
	//
	// 	- Enhanced SSD (cloud_essd): 20 to 32768
	//
	// 	- Basic disk (cloud): 5 to 2000
	//
	// >  The capacity of the volume after the volume is scaled up must be greater than the original capacity of the volume. If the new capacity is equal to the original capacity of the volume, only the file system is scaled up.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	NewSize      *int64  `json:"NewSize,omitempty" xml:"NewSize,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the volume that is mounted to the elastic container instance. Only disk volumes can be scaled up.
	//
	// This parameter is required.
	//
	// example:
	//
	// default-volume1
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
	// The request ID.
	//
	// example:
	//
	// 89945DD3-9072-47D0-A318-353284CFC7B3
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeContainerGroupVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure the idempotence of a request?](https://help.aliyun.com/document_detail/25693.html)
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// eci-2zelg8vwnlzdhf8hv****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	// The request ID.
	//
	// example:
	//
	// 45D5B0AD-3B00-4A9B-9911-6D5303B06712
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartContainerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type TagResourcesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources.
	//
	// This parameter is required.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- ContainerGroup: elastic container instance
	//
	// 	- ImageCache: image cache
	//
	// 	- DataCache: data cache
	//
	// 	- VirtualNode: virtual node
	//
	// This parameter is required.
	//
	// example:
	//
	// ContainerGroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags that you want to add to the resource. A maximum of 20 tags can be added to a resource.
	//
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerId(v int64) *TagResourcesRequest {
	s.ResourceOwnerId = &v
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
	// The tag key. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// testkey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// testvalue
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
	// The request ID.
	//
	// example:
	//
	// 0E234675-3465-4CC3-9D0F-9A864BC391DD
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the resource. This parameter is valid only when the `TagKey` parameter is not specified in the request. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-xxx-xxx-xxxx-42665544xxxx
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources.
	//
	// This parameter is required.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- ContainerGroup: elastic container instance
	//
	// 	- ImageCache: image cache
	//
	// 	- DataCache: data cache
	//
	// 	- VirtualNode: virtual node
	//
	// This parameter is required.
	//
	// example:
	//
	// ContainerGroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The keys of the tags that you want to remove from the resource. You can enter up to 20 tag keys.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetClientToken(v string) *UntagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerId(v int64) *UntagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 89945DD3-9072-47D0-A318-353284CFC7B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateContainerGroupRequest struct {
	DnsConfig *UpdateContainerGroupRequestDnsConfig `json:"DnsConfig,omitempty" xml:"DnsConfig,omitempty" type:"Struct"`
	// Details of the Container Registry Enterprise Edition instance that hosts the image of the init container.
	AcrRegistryInfo []*UpdateContainerGroupRequestAcrRegistryInfo `json:"AcrRegistryInfo,omitempty" xml:"AcrRegistryInfo,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new configurations of the container group.
	Container []*UpdateContainerGroupRequestContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Repeated"`
	// The ID of the elastic container instance that you want to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// eci-2zelg8vwnlzdhf8hv****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The number of vCPUs that are allocated to the elastic container instance.
	//
	// example:
	//
	// 2.0
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The information about the credentials of the image repository.
	ImageRegistryCredential []*UpdateContainerGroupRequestImageRegistryCredential `json:"ImageRegistryCredential,omitempty" xml:"ImageRegistryCredential,omitempty" type:"Repeated"`
	// The information about the new init container.
	InitContainer []*UpdateContainerGroupRequestInitContainer `json:"InitContainer,omitempty" xml:"InitContainer,omitempty" type:"Repeated"`
	// The size of the memory that is allocated to the elastic container instance. Unit: GiB.
	//
	// example:
	//
	// 4.0
	Memory       *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	OwnerAccount *string  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-2df3isufhi38****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The restart policy of the elastic container instance. Valid values:
	//
	// 	- Always: Always restarts the instance if a container in the instance exits upon termination.
	//
	// 	- Never: Never restarts the instance if a container in the instance exits upon termination.
	//
	// 	- OnFailure: Restarts the instance only if a container in the instance exists upon failure with a status code of non-zero.
	//
	// example:
	//
	// Always
	RestartPolicy *string `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	// The tags that are bound to the instance.
	Tag []*UpdateContainerGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The update type. Valid values:
	//
	// 	- RenewUpdate: full updates. You must specify all relevant parameters to update the elastic container instance. For a parameter of the list type, you must specify all the items contained in the parameter even if you want to update only some of the items. For a parameter of the struct type, you must specify all the members even if you want to update only some of the members.
	//
	// 	- IncrementalUpdate: incremental updates. You may specify only the parameter that you want to update. Other related parameters remain unchanged.
	//
	// Default value: RenewUpdate.
	//
	// example:
	//
	// RenewUpdate
	UpdateType *string `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
	// The volumes that are mounted to the instance.
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
	// The IP addresses of DNS servers.
	//
	// example:
	//
	// 1.2.3.4
	NameServer []*string `json:"NameServer,omitempty" xml:"NameServer,omitempty" type:"Repeated"`
	// The configurations of DNS.
	Option []*UpdateContainerGroupRequestDnsConfigOption `json:"Option,omitempty" xml:"Option,omitempty" type:"Repeated"`
	// The search domains of the Domain Name System (DNS) server.
	//
	// example:
	//
	// my.dns.search.suffix
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
	//
	// example:
	//
	// ndots
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The option value of DNS configurations.
	//
	// example:
	//
	// 2
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
	// The domain names of the Container Registry Enterprise Edition instance. By default, all domain names of the instance are displayed. You can specify specific domain names. Separate multiple domain names with commas (,).
	Domain []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
	// The ID of the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// cri-nwj395hgf6f3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// acr-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the region where the Container Registry Enterprise Edition instance resides.
	//
	// example:
	//
	// cn-beijing
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
	LivenessProbe   *UpdateContainerGroupRequestContainerLivenessProbe   `json:"LivenessProbe,omitempty" xml:"LivenessProbe,omitempty" type:"Struct"`
	ReadinessProbe  *UpdateContainerGroupRequestContainerReadinessProbe  `json:"ReadinessProbe,omitempty" xml:"ReadinessProbe,omitempty" type:"Struct"`
	SecurityContext *UpdateContainerGroupRequestContainerSecurityContext `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	// The arguments that you want to pass to the startup command of the container. You can specify up to 10 arguments.
	//
	// example:
	//
	// hello
	Arg []*string `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	// The commands that you want to run to perform the health check.
	//
	// example:
	//
	// echo
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// The number of vCPUs that you want to allocate to the container
	//
	// example:
	//
	// 1.0
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The environment variables for the container.
	EnvironmentVar []*UpdateContainerGroupRequestContainerEnvironmentVar `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	// The number of GPUs that you want to allocate to the container.
	//
	// example:
	//
	// 1
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The image of the container.
	//
	// example:
	//
	// jenkins
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image pulling policy. Valid values:
	//
	// 	- Always: Each time the instance is updated, image pulling is performed.
	//
	// 	- IfNotPresent: On-premises images are used first. If no on-premises images are available, image pulling is performed.
	//
	// 	- Never: On-premises images are always used. Image pulling is not performed.
	//
	// example:
	//
	// Never
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The commands to be executed in the container when you use the CLI to specify the postStart callback function.
	//
	// example:
	//
	// hide
	LifecyclePostStartHandlerExec []*string `json:"LifecyclePostStartHandlerExec,omitempty" xml:"LifecyclePostStartHandlerExec,omitempty" type:"Repeated"`
	// The IP address of the host that receives the HTTP GET request when you use an HTTP request to specify the postStart callback function.
	//
	// example:
	//
	// hide
	LifecyclePostStartHandlerHttpGetHost *string `json:"LifecyclePostStartHandlerHttpGetHost,omitempty" xml:"LifecyclePostStartHandlerHttpGetHost,omitempty"`
	// The information about the valid HTTP request headers among the generated HTTP request headers.
	LifecyclePostStartHandlerHttpGetHttpHeaders []*UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders `json:"LifecyclePostStartHandlerHttpGetHttpHeaders,omitempty" xml:"LifecyclePostStartHandlerHttpGetHttpHeaders,omitempty" type:"Repeated"`
	// The path to which the system sends an HTTP GET request for a health check when you use an HTTP request to specify the postStart callback function.
	//
	// example:
	//
	// /healthyz
	LifecyclePostStartHandlerHttpGetPath *string `json:"LifecyclePostStartHandlerHttpGetPath,omitempty" xml:"LifecyclePostStartHandlerHttpGetPath,omitempty"`
	// The port to which the system sends the HTTP GET request when you use an HTTP request to specify the postStart callback function.
	//
	// example:
	//
	// 1
	LifecyclePostStartHandlerHttpGetPort *int32 `json:"LifecyclePostStartHandlerHttpGetPort,omitempty" xml:"LifecyclePostStartHandlerHttpGetPort,omitempty"`
	// The path to which the system sends an HTTP GET request for a health check when you use an HTTP request to specify the postStart callback function.
	//
	// example:
	//
	// /healthyz
	LifecyclePostStartHandlerHttpGetScheme *string `json:"LifecyclePostStartHandlerHttpGetScheme,omitempty" xml:"LifecyclePostStartHandlerHttpGetScheme,omitempty"`
	// The IP address of the host that receives the TCP socket request when you use a TCP socket request to specify the postStart callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePostStartHandlerTcpSocketHost *string `json:"LifecyclePostStartHandlerTcpSocketHost,omitempty" xml:"LifecyclePostStartHandlerTcpSocketHost,omitempty"`
	// The port to which the system sends a TCP socket request for a health check when you use TCP sockets to specify the postStart callback function.
	//
	// example:
	//
	// 1
	LifecyclePostStartHandlerTcpSocketPort *int32 `json:"LifecyclePostStartHandlerTcpSocketPort,omitempty" xml:"LifecyclePostStartHandlerTcpSocketPort,omitempty"`
	// The commands to be executed in the container when you use the CLI to specify the preStop callback function.
	//
	// example:
	//
	// hide
	LifecyclePreStopHandlerExec []*string `json:"LifecyclePreStopHandlerExec,omitempty" xml:"LifecyclePreStopHandlerExec,omitempty" type:"Repeated"`
	// The IP address of the host that receives the HTTP GET request when you use an HTTP request to specify the preStop callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePreStopHandlerHttpGetHost *string `json:"LifecyclePreStopHandlerHttpGetHost,omitempty" xml:"LifecyclePreStopHandlerHttpGetHost,omitempty"`
	// The information about the generated HTTP request header.
	LifecyclePreStopHandlerHttpGetHttpHeader []*UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader `json:"LifecyclePreStopHandlerHttpGetHttpHeader,omitempty" xml:"LifecyclePreStopHandlerHttpGetHttpHeader,omitempty" type:"Repeated"`
	// The path to which the system sends an HTTP GET request for a health check when you use an HTTP request to specify the preSop callback function.
	//
	// example:
	//
	// /healthyz
	LifecyclePreStopHandlerHttpGetPath *string `json:"LifecyclePreStopHandlerHttpGetPath,omitempty" xml:"LifecyclePreStopHandlerHttpGetPath,omitempty"`
	// The port to which the system sends the HTTP GET request for a health check when you use an HTTP request to specify the preStop callback function.
	//
	// example:
	//
	// 1
	LifecyclePreStopHandlerHttpGetPort *int32 `json:"LifecyclePreStopHandlerHttpGetPort,omitempty" xml:"LifecyclePreStopHandlerHttpGetPort,omitempty"`
	// The protocol type of the HTTP GET request when you use an HTTP request to specify the preStop callback function. Valid values:
	//
	// 	- HTTP
	//
	// 	- HTTPS
	//
	// example:
	//
	// HTTP
	LifecyclePreStopHandlerHttpGetScheme *string `json:"LifecyclePreStopHandlerHttpGetScheme,omitempty" xml:"LifecyclePreStopHandlerHttpGetScheme,omitempty"`
	// The IP address of the host that receives the TCP socket request when you use a TCP socket request to specify the preStop callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePreStopHandlerTcpSocketHost *string `json:"LifecyclePreStopHandlerTcpSocketHost,omitempty" xml:"LifecyclePreStopHandlerTcpSocketHost,omitempty"`
	// The port to which the system sends a TCP socket request for a health check when you use TCP sockets to specify the preStop callback function.
	//
	// example:
	//
	// 80
	LifecyclePreStopHandlerTcpSocketPort *int32 `json:"LifecyclePreStopHandlerTcpSocketPort,omitempty" xml:"LifecyclePreStopHandlerTcpSocketPort,omitempty"`
	// The memory size of the container.
	//
	// example:
	//
	// 2.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the container.
	//
	// example:
	//
	// jenkins
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port to which the system sends an HTTP GET request for a health check.
	Port []*UpdateContainerGroupRequestContainerPort `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
	// Specifies whether the container allocates buffer resources to standard input streams when the container is running. If you do not specify this parameter, an end-of-file (EOF) error may occur when standard input streams in the container are read. Default value: false.
	//
	// example:
	//
	// false
	Stdin *bool `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	// Specifies whether standard input streams are disconnected after a client is disconnected. If Stdin is set to true, standard input streams remain connected among multiple sessions. If StdinOnce is set to true, standard input streams are connected after the container is started, and remain idle until a client is connected to receive data. After the client is disconnected, streams are also disconnected, and remain disconnected until the container restarts.
	//
	// example:
	//
	// true
	StdinOnce *bool `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	// Specifies whether to enable interaction. Default value: false. If the command is a /bin/bash command, set the value to true.
	//
	// example:
	//
	// false
	Tty *bool `json:"Tty,omitempty" xml:"Tty,omitempty"`
	// Pod volumes that you want to mount into the filesystem of the container.
	VolumeMount []*UpdateContainerGroupRequestContainerVolumeMount `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
	// The working directory of the container.
	//
	// example:
	//
	// /usr/share/
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
	Exec                *UpdateContainerGroupRequestContainerLivenessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" type:"Struct"`
	FailureThreshold    *int32                                                      `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *UpdateContainerGroupRequestContainerLivenessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	InitialDelaySeconds *int32                                                      `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                      `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                      `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *UpdateContainerGroupRequestContainerLivenessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
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
	Exec                *UpdateContainerGroupRequestContainerReadinessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" type:"Struct"`
	FailureThreshold    *int32                                                       `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *UpdateContainerGroupRequestContainerReadinessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	InitialDelaySeconds *int32                                                       `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                       `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                       `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *UpdateContainerGroupRequestContainerReadinessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
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
	Capability             *UpdateContainerGroupRequestContainerSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" type:"Struct"`
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
	FieldRef *UpdateContainerGroupRequestContainerEnvironmentVarFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" type:"Struct"`
	// The name of the environment variable for the container.
	//
	// example:
	//
	// PATH
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable for the container.
	//
	// example:
	//
	// /usr/bin/local/
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
	// The request parameter of the HTTP GET request when you use an HTTP request to specify the postStart callback function.
	//
	// example:
	//
	// testValue
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request parameter value of the HTTP GET request when you use an HTTP request to specify the postStart callback function.
	//
	// example:
	//
	// test
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
	// The request parameter of the HTTP GET request when you use an HTTP request to specify the preStop callback function.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request parameter value of the HTTP GET request when you use an HTTP request to specify the preStop callback function.
	//
	// example:
	//
	// testValue
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
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol of the container. Valid values: TCP and UDP.
	//
	// example:
	//
	// TCP
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
	// The directory of the volume that is mounted to the container. The data in this directory is overwritten by the data on the volume. Specify this parameter with caution.
	//
	// example:
	//
	// /usr/share/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation settings of the volume. Mount propagation allows volumes that are mounted on one container to be shared with other containers in the same pod, or even with other pods on the same node. Valid values:
	//
	// 	- None: This volume mount does not receive subsequent mounts that are performed on this volume or subdirectories of this volume.
	//
	// 	- HostToCotainer: The volume mount receives subsequent mounts that are performed on this volume or the subdirectories of this volume.
	//
	// 	- Bidirectional: The volume mount behaves the same as the HostToContainer mount. The volume mount receives subsequent mounts that are performed on the volume or the subdirectories of the volume. In addition, all volume mounts that are mounted on the container are propagated back to the host and all containers of all pods that use the same volume.
	//
	// Default value: None.
	//
	// example:
	//
	// None
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The name of the volume that is mounted to the container. Valid values: the values of Volume.N.Name, which are the names of volumes that are mounted to the elastic container instance.
	//
	// example:
	//
	// test-empty
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the volume is read-only. Default value: false.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The subdirectory of the volume that is mounted to the container. You can use this parameter to mount the same volume to different subdirectories of the container.
	//
	// example:
	//
	// /usr/share/sub/
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
	// The password that you use to access the image repository.
	//
	// example:
	//
	// yourpassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The address of the image repository. This address does not contain `http://` or `https://`.
	//
	// example:
	//
	// registry.cn-shanghai.aliyuncs.com/ecitest/nginx:alpine
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The username that you use to access the image repository.
	//
	// example:
	//
	// yourname
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
	SecurityContext *UpdateContainerGroupRequestInitContainerSecurityContext `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	// The arguments that you want to pass to the startup command of the init container.
	//
	// example:
	//
	// 10
	Arg []*string `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	// The commands that are used to start the init container.
	//
	// example:
	//
	// /bin/sh sleep
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// The number of vCPUs that you want to allocate to the init container.
	//
	// example:
	//
	// 2.0
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The environment variable of the init container.
	EnvironmentVar []*UpdateContainerGroupRequestInitContainerEnvironmentVar `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	// The number of GPUs you want to allocate to the init container.
	//
	// example:
	//
	// 1
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The image of the init container.
	//
	// example:
	//
	// nginx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image pulling policy. Valid values:
	//
	// 	- Always: Each time the instance is updated, image pulling is performed.
	//
	// 	- IfNotPresent: On-premises images are used first. If no on-premises images are available, image pulling is performed.
	//
	// 	- Never: On-premises images are always used. Image pulling is not performed.
	//
	// example:
	//
	// Onfailure
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The memory size of the init container.
	//
	// example:
	//
	// 4.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the init container.
	//
	// example:
	//
	// init-nginx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port number. Valid values: 1 to 65535.
	Port []*UpdateContainerGroupRequestInitContainerPort `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
	// Specifies whether the init container allocates buffer resources to standard input streams when the init container is running. If you do not specify this parameter, an EOF error may occur when standard input streams in the init container are read. Default value: false.
	//
	// example:
	//
	// false
	Stdin *bool `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	// Specifies whether standard input streams are disconnected after a client is disconnected. If Stdin is set to true, standard input streams remain connected among multiple sessions. If StdinOnce is set to true, standard input streams are connected after the init container is started, and remain idle until a client is connected to receive data. After the client is disconnected, streams are also disconnected, and remain disconnected until the init container restarts.
	//
	// example:
	//
	// true
	StdinOnce *bool `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	// Specifies whether to enable interaction. Default value: false. If the command is a /bin/bash command, set the value to true.
	//
	// example:
	//
	// true
	Tty *bool `json:"Tty,omitempty" xml:"Tty,omitempty"`
	// The information about the volume that you want to mount on the init container.
	VolumeMount []*UpdateContainerGroupRequestInitContainerVolumeMount `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
	// The working directory of the init container.
	//
	// example:
	//
	// /bin/local/
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
	Capability             *UpdateContainerGroupRequestInitContainerSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" type:"Struct"`
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
	FieldRef *UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" type:"Struct"`
	// The name of the environment variable for the init container.
	//
	// example:
	//
	// PATH
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable for the init container.
	//
	// example:
	//
	// /usr/local/bin
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
	// The port number of the init container. Valid values: 1 to 65535.
	//
	// example:
	//
	// 9000
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol of the init container. Valid values: TCP and UDP.
	//
	// example:
	//
	// TCP
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
	// The mount directory of the init container. The data in this directory is overwritten by the data on the volume. Specify this parameter with caution.
	//
	// example:
	//
	// /pod/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation settings of the volume. Mount propagation allows volumes that are mounted on one container to be shared with other containers in the same pod, or even with other pods on the same node. Valid values:
	//
	// 	- None: The volume mount does not receive subsequent mounts that are performed on this volume or subdirectories of this volume.
	//
	// 	- HostToContainer: The volume mount receives all subsequent mounts that are performed on this volume or subdirectories of this volume.
	//
	// 	- Bidirectional: The volume mount behaves the same as the HostToContainer mount. The volume mount receives subsequent mounts that are performed on the volume or the subdirectories of the volume. In addition, all volume mounts that are mounted on the container are propagated back to the host and all containers of all pods that use the same volume.
	//
	// Default value: None.
	//
	// example:
	//
	// None
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The name of the volume that is mounted to the init container. Valid values: the values of Volume.N.Name, which are the names of volumes that are mounted to the elastic container instance.
	//
	// example:
	//
	// default-volume1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the volume is read-only. Default value: false.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The subdirectory of the volume that is mounted to the init container. You can use this parameter to mount the same volume to different subdirectories of the init container.
	//
	// example:
	//
	// data2/
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
	// The tag key.
	//
	// example:
	//
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// hxh
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
	ConfigFileVolume *UpdateContainerGroupRequestVolumeConfigFileVolume `json:"ConfigFileVolume,omitempty" xml:"ConfigFileVolume,omitempty" type:"Struct"`
	EmptyDirVolume   *UpdateContainerGroupRequestVolumeEmptyDirVolume   `json:"EmptyDirVolume,omitempty" xml:"EmptyDirVolume,omitempty" type:"Struct"`
	FlexVolume       *UpdateContainerGroupRequestVolumeFlexVolume       `json:"FlexVolume,omitempty" xml:"FlexVolume,omitempty" type:"Struct"`
	HostPathVolume   *UpdateContainerGroupRequestVolumeHostPathVolume   `json:"HostPathVolume,omitempty" xml:"HostPathVolume,omitempty" type:"Struct"`
	NFSVolume        *UpdateContainerGroupRequestVolumeNFSVolume        `json:"NFSVolume,omitempty" xml:"NFSVolume,omitempty" type:"Struct"`
	// The volume name.
	//
	// example:
	//
	// test-empty
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the HostPath volume. Valid values:
	//
	// 	- Directory
	//
	// 	- File
	//
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// EmptyDirVolume
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
	// The request ID.
	//
	// example:
	//
	// CB8D2B22-D636-4182-****-1FC9DBDAD66F
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContainerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UpdateDataCacheRequest struct {
	// The bucket in which the data cache is stored. Default value: default.
	//
	// example:
	//
	// default
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate a token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure the idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 2f22dc16-cad0-4d3f-87e5-91e604756547
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The data cache ID.
	//
	// example:
	//
	// edc-bp1a7n7uawwwol******
	DataCacheId *string `json:"DataCacheId,omitempty" xml:"DataCacheId,omitempty"`
	// The information about the data source.
	DataSource *UpdateDataCacheRequestDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The elastic IP address (EIP) to be created and associated. If no NAT gateway is configured for the virtual private cloud (VPC), you can associate an EIP to pull data from the Internet.
	EipCreateParam *UpdateDataCacheRequestEipCreateParam `json:"EipCreateParam,omitempty" xml:"EipCreateParam,omitempty" type:"Struct"`
	// The ID of the elastic IP address (EIP). If no NAT gateway is configured for the virtual private cloud (VPC), you can bind an EIP to the elastic container instance to pull data from the Internet.
	//
	// example:
	//
	// eip-2zedsm5mfl3uhdj2d****
	EipInstanceId *string `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	// The data cache name.
	//
	// example:
	//
	// test
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzw2jz2z****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The retention period for the data cache. The data cache is deleted after the retention period expires. By default, the data cache does not expire.
	//
	// example:
	//
	// 7
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-2ze81zoc3yl7a3we****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The data cache size.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The tags that are added to the data cache.
	Tag []*UpdateDataCacheRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2ze23nqzig8inprou****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s UpdateDataCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataCacheRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataCacheRequest) SetBucket(v string) *UpdateDataCacheRequest {
	s.Bucket = &v
	return s
}

func (s *UpdateDataCacheRequest) SetClientToken(v string) *UpdateDataCacheRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateDataCacheRequest) SetDataCacheId(v string) *UpdateDataCacheRequest {
	s.DataCacheId = &v
	return s
}

func (s *UpdateDataCacheRequest) SetDataSource(v *UpdateDataCacheRequestDataSource) *UpdateDataCacheRequest {
	s.DataSource = v
	return s
}

func (s *UpdateDataCacheRequest) SetEipCreateParam(v *UpdateDataCacheRequestEipCreateParam) *UpdateDataCacheRequest {
	s.EipCreateParam = v
	return s
}

func (s *UpdateDataCacheRequest) SetEipInstanceId(v string) *UpdateDataCacheRequest {
	s.EipInstanceId = &v
	return s
}

func (s *UpdateDataCacheRequest) SetName(v string) *UpdateDataCacheRequest {
	s.Name = &v
	return s
}

func (s *UpdateDataCacheRequest) SetOwnerAccount(v string) *UpdateDataCacheRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateDataCacheRequest) SetOwnerId(v int64) *UpdateDataCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateDataCacheRequest) SetRegionId(v string) *UpdateDataCacheRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataCacheRequest) SetResourceGroupId(v string) *UpdateDataCacheRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateDataCacheRequest) SetResourceOwnerAccount(v string) *UpdateDataCacheRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateDataCacheRequest) SetResourceOwnerId(v int64) *UpdateDataCacheRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateDataCacheRequest) SetRetentionDays(v int32) *UpdateDataCacheRequest {
	s.RetentionDays = &v
	return s
}

func (s *UpdateDataCacheRequest) SetSecurityGroupId(v string) *UpdateDataCacheRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateDataCacheRequest) SetSize(v int32) *UpdateDataCacheRequest {
	s.Size = &v
	return s
}

func (s *UpdateDataCacheRequest) SetTag(v []*UpdateDataCacheRequestTag) *UpdateDataCacheRequest {
	s.Tag = v
	return s
}

func (s *UpdateDataCacheRequest) SetVSwitchId(v string) *UpdateDataCacheRequest {
	s.VSwitchId = &v
	return s
}

type UpdateDataCacheRequestDataSource struct {
	// The parameters that are configured for the data source.
	Options map[string]*string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- URL
	//
	// 	- NAS
	//
	// 	- OSS
	//
	// 	- SNAPSHOT
	//
	// example:
	//
	// NAS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateDataCacheRequestDataSource) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataCacheRequestDataSource) GoString() string {
	return s.String()
}

func (s *UpdateDataCacheRequestDataSource) SetOptions(v map[string]*string) *UpdateDataCacheRequestDataSource {
	s.Options = v
	return s
}

func (s *UpdateDataCacheRequestDataSource) SetType(v string) *UpdateDataCacheRequestDataSource {
	s.Type = &v
	return s
}

type UpdateDataCacheRequestEipCreateParam struct {
	// The bandwidth of the EIP. Unit: Mbit/s. Default value: 5.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The EIP bandwidth plan to be associated.
	//
	// example:
	//
	// cbwp-2zeukbj916scmj51m****
	CommonBandwidthPackage *string `json:"CommonBandwidthPackage,omitempty" xml:"CommonBandwidthPackage,omitempty"`
	// The line type of the EIP. Valid values:
	//
	// 	- BGP (default): BGP (Multi-ISP) line
	//
	// 	- BGP_PRO: BGP (Multi-ISP) Pro line
	//
	// example:
	//
	// BGP
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The metering method of the EIP. Valid values:
	//
	// 	- PayByBandwidth: pay-by-bandwidth
	//
	// 	- PayByTraffic: pay-by-data-transfer
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The ID of the IP address pool. The EIP is allocated from the IP address pool. You cannot use the IP address pool feature by default. To use this feature, you must apply for the privilege in the Quota Center console.
	//
	// example:
	//
	// pippool-bp187arfugi543y1s****
	PublicIpAddressPoolId *string `json:"PublicIpAddressPoolId,omitempty" xml:"PublicIpAddressPoolId,omitempty"`
}

func (s UpdateDataCacheRequestEipCreateParam) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataCacheRequestEipCreateParam) GoString() string {
	return s.String()
}

func (s *UpdateDataCacheRequestEipCreateParam) SetBandwidth(v int32) *UpdateDataCacheRequestEipCreateParam {
	s.Bandwidth = &v
	return s
}

func (s *UpdateDataCacheRequestEipCreateParam) SetCommonBandwidthPackage(v string) *UpdateDataCacheRequestEipCreateParam {
	s.CommonBandwidthPackage = &v
	return s
}

func (s *UpdateDataCacheRequestEipCreateParam) SetISP(v string) *UpdateDataCacheRequestEipCreateParam {
	s.ISP = &v
	return s
}

func (s *UpdateDataCacheRequestEipCreateParam) SetInternetChargeType(v string) *UpdateDataCacheRequestEipCreateParam {
	s.InternetChargeType = &v
	return s
}

func (s *UpdateDataCacheRequestEipCreateParam) SetPublicIpAddressPoolId(v string) *UpdateDataCacheRequestEipCreateParam {
	s.PublicIpAddressPoolId = &v
	return s
}

type UpdateDataCacheRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// testkey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testvalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateDataCacheRequestTag) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataCacheRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateDataCacheRequestTag) SetKey(v string) *UpdateDataCacheRequestTag {
	s.Key = &v
	return s
}

func (s *UpdateDataCacheRequestTag) SetValue(v string) *UpdateDataCacheRequestTag {
	s.Value = &v
	return s
}

type UpdateDataCacheResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD8BBB43-8E05-5F46-89A9-2512D8A324A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDataCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataCacheResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataCacheResponseBody) SetRequestId(v string) *UpdateDataCacheResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDataCacheResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataCacheResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataCacheResponse) SetHeaders(v map[string]*string) *UpdateDataCacheResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataCacheResponse) SetStatusCode(v int32) *UpdateDataCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataCacheResponse) SetBody(v *UpdateDataCacheResponseBody) *UpdateDataCacheResponse {
	s.Body = v
	return s
}

type UpdateImageCacheRequest struct {
	// The information about the Container Registry Enterprise Edition instance.
	AcrRegistryInfo []*UpdateImageCacheRequestAcrRegistryInfo `json:"AcrRegistryInfo,omitempty" xml:"AcrRegistryInfo,omitempty" type:"Repeated"`
	// Specifies whether to enable reuse of image cache layers. If you enable this feature and the image cache that you want to create and an existing image cache contain duplicate image layers, the system reuses the duplicate image layers to create the new image cache. This accelerates the creation of the image cache. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// true
	AutoMatchImageCache *bool `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure the idempotence of a request?](https://help.aliyun.com/document_detail/25693.html)
	//
	// example:
	//
	// 123e4567-xxx-xxx-xxxx-42665544xxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the elastic IP address (EIP). If you want to pull public images, you must make sure that the elastic container instance can access the Internet. To enable Internet access, you can configure an EIP or a NAT gateway for the instance.
	//
	// example:
	//
	// eip-2zedsm5mfl3uhdj2d****
	EipInstanceId *string `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	// The elimination policy for the image cache. This parameter is empty by default, which indicates that the image cache is always retained.
	//
	// You can set this parameter to LRU, which indicates that the image cache can be automatically deleted. When the number of image caches reaches the quota, the system automatically deletes the image caches whose EliminationStrategy parameter is set to LRU and that are least recently used.
	//
	// example:
	//
	// LRU
	EliminationStrategy *string `json:"EliminationStrategy,omitempty" xml:"EliminationStrategy,omitempty"`
	// Specifies whether to enable the instant image cache feature. The feature can accelerate the creation of image caches. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// true
	Flash *bool `json:"Flash,omitempty" xml:"Flash,omitempty"`
	// The number of duplicates of temporary local snapshots. By default, the system creates one snapshot for each image cache. If you use the image cache to create multiple elastic container instances at a time, we recommend that you configure this parameter to create multiple snapshot duplicates for the image cache. We recommend that you create one snapshot duplicate for creation of every 1,000 elastic container instances.
	//
	// > If you enable the instant image cache feature by setting Flash to true, a local snapshot is first created during the creation of the image cache. After the local snapshot is created, regular snapshots start to be created. After the regular snapshots are created, the local snapshot is automatically deleted.
	//
	// example:
	//
	// 7
	FlashCopyCount *int32 `json:"FlashCopyCount,omitempty" xml:"FlashCopyCount,omitempty"`
	// Container images that are used to create the image cache.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/eci_open/nginx:1.15.10-perl
	Image []*string `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
	// The ID of the image cache.
	//
	// This parameter is required.
	//
	// example:
	//
	// imc-bp195erqe9o2pb09****
	ImageCacheId *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	// The name of the image cache.
	//
	// example:
	//
	// testcache
	ImageCacheName *string `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	// The size of the image cache. Unit: GiB. Default value: 20.
	//
	// example:
	//
	// 20
	ImageCacheSize *int32 `json:"ImageCacheSize,omitempty" xml:"ImageCacheSize,omitempty"`
	// The information about the image repository.
	ImageRegistryCredential []*UpdateImageCacheRequestImageRegistryCredential `json:"ImageRegistryCredential,omitempty" xml:"ImageRegistryCredential,omitempty" type:"Repeated"`
	OwnerAccount            *string                                           `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64                                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the image cache.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the image cache belongs.
	//
	// example:
	//
	// rg-aekzh43v*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The retention period of the image cache. Unit: days. When the retention period elapses, the image cache expires and is deleted. By default, image caches never expire.
	//
	// > The image caches that fail to be created are retained for only 1 day.
	//
	// example:
	//
	// 7
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The ID of the security group to which the image cache belongs.
	//
	// example:
	//
	// sg-uf66jeqopgqa9hdn****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The number of duplicates of regular snapshots. By default, the system creates one snapshot for each image cache. If you use the image cache to create multiple elastic container instances at a time, we recommend that you configure this parameter to create multiple snapshot duplicates for the image cache. We recommend that you create one snapshot duplicate for creation of every 1,000 elastic container instances.
	//
	// > If you disable the instant image cache feature by setting Flash to false, only regular snapshots are generated when you create an image cache.
	//
	// example:
	//
	// 7
	StandardCopyCount *int32 `json:"StandardCopyCount,omitempty" xml:"StandardCopyCount,omitempty"`
	// The tags to add to the image cache. A maximum of 20 tags can be added to the image cache.
	Tag []*UpdateImageCacheRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which the image cache is connected.
	//
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
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
	// The domain names of the Container Registry Enterprise Edition instance. By default, all domain names of the instance are displayed. You can specify multiple domain names. Separate multiple domain names with commas (,).
	//
	// example:
	//
	// test****-registry.cn-hangzhou.cr.aliyuncs.com
	Domain []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
	// The ID of the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// cri-nwj395hgf6f3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// test****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID of the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// cn-hangzhou
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
	// The password that is used to access the image repository.
	//
	// example:
	//
	// password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The image repository address without `http://` or `https://` as a prefix.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The username that is used to access the image repository.
	//
	// example:
	//
	// username
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
	// The key of tag N to add to the image cache.
	//
	// example:
	//
	// imc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the image cache.
	//
	// example:
	//
	// test
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
	// The request ID.
	//
	// example:
	//
	// 0E234675-3465-4CC3-9D0F-9A864BC391DD
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateImageCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency of requests?](https://help.aliyun.com/document_detail/25693.html)
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IP address of the DNS server. If `dnsPolicy=ClusterFirst` is configured for the Elastic Container Instance pod, Elastic Container Instance uses the configuration to provide DNS services to containers. You can configure multiple IP addresses. Separate multiple IP addresses with commas (,).
	//
	// example:
	//
	// 100.1.XX.XX
	ClusterDNS *string `json:"ClusterDNS,omitempty" xml:"ClusterDNS,omitempty"`
	// The domain name of the cluster. If this parameter is specified, in addition to the search domain of the host, Kubelet configures all containers to search for the specified domain name.
	//
	// example:
	//
	// example.com
	ClusterDomain *string `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	// The custom resources that are supported by the virtual node. If a custom resource is specified in the request of an Elastic Container Instance pod, the pod is scheduled to run on the virtual node that supports the custom resource. You can use the `Resource name = Number of resources` format to specify custom resources. Separate multiple resources with commas (,).
	//
	// example:
	//
	// example1.com=100,example2.com=200
	CustomResources *string `json:"CustomResources,omitempty" xml:"CustomResources,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the virtual node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags to add to the virtual node.
	Tag []*UpdateVirtualNodeRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the virtual node.
	//
	// This parameter is required.
	//
	// example:
	//
	// vnd-2ze960zkdqrldeaw****
	VirtualNodeId *string `json:"VirtualNodeId,omitempty" xml:"VirtualNodeId,omitempty"`
	// The name of the virtual node.
	//
	// example:
	//
	// testNode
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

func (s *UpdateVirtualNodeRequest) SetResourceOwnerAccount(v string) *UpdateVirtualNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetResourceOwnerId(v int64) *UpdateVirtualNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetTag(v []*UpdateVirtualNodeRequestTag) *UpdateVirtualNodeRequest {
	s.Tag = v
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
	// The key of tag N to add to the virtual node.
	//
	// example:
	//
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the virtual node.
	//
	// example:
	//
	// test
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
	// The request ID.
	//
	// example:
	//
	// 89945DD3-9072-47D0-A318-353284CFC7B3
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVirtualNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

// Summary:
//
// Creates an asynchronous task to store a specified container in an elastic container instance as an image and pushes the image to an image repository of Alibaba Cloud Container Registry.
//
// Description:
//
// You must specify the Alibaba Cloud Resource Name (ARN) of the RAM role of the Container Registry Enterprise Edition instance to grant the elastic container instance to assume the RAM role to push images.
//
// @param request - CommitContainerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CommitContainerResponse
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

// Summary:
//
// Creates an asynchronous task to store a specified container in an elastic container instance as an image and pushes the image to an image repository of Alibaba Cloud Container Registry.
//
// Description:
//
// You must specify the Alibaba Cloud Resource Name (ARN) of the RAM role of the Container Registry Enterprise Edition instance to grant the elastic container instance to assume the RAM role to push images.
//
// @param request - CommitContainerRequest
//
// @return CommitContainerResponse
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

// Summary:
//
// Copies a DataCache from one region to another region.
//
// Description:
//
// You cannot directly use a DataCache across regions. You can call this operation to copy a DataCache from one region to another region. This operation is suitable for the following scenarios:
//
//   - If you want to use a DataCache across regions and the DataCache exists in Region A, you can call this operation to quickly copy the DataCache to Region B.
//
//   - If you directly pull data from a region outside China to a region inside the Chinese mainland when you create a DataCache, the data may be pulled at a slow speed due to network limits. In this case, you can create a DataCache in a region outside the Chinese mainland but inside China, such as the China (Hong Kong) region, and call this operation to copy the data to the region inside the Chinese mainland.
//
// > The process of copying a DataCache is equivalent to copying a snapshot. You are charged for the traffic generated during the copy process and the storage of the generated DataCache.
//
// @param request - CopyDataCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyDataCacheResponse
func (client *Client) CopyDataCacheWithOptions(request *CopyDataCacheRequest, runtime *util.RuntimeOptions) (_result *CopyDataCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bucket)) {
		query["Bucket"] = request.Bucket
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataCacheId)) {
		query["DataCacheId"] = request.DataCacheId
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationRegionId)) {
		query["DestinationRegionId"] = request.DestinationRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
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

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CopyDataCache"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyDataCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Copies a DataCache from one region to another region.
//
// Description:
//
// You cannot directly use a DataCache across regions. You can call this operation to copy a DataCache from one region to another region. This operation is suitable for the following scenarios:
//
//   - If you want to use a DataCache across regions and the DataCache exists in Region A, you can call this operation to quickly copy the DataCache to Region B.
//
//   - If you directly pull data from a region outside China to a region inside the Chinese mainland when you create a DataCache, the data may be pulled at a slow speed due to network limits. In this case, you can create a DataCache in a region outside the Chinese mainland but inside China, such as the China (Hong Kong) region, and call this operation to copy the data to the region inside the Chinese mainland.
//
// > The process of copying a DataCache is equivalent to copying a snapshot. You are charged for the traffic generated during the copy process and the storage of the generated DataCache.
//
// @param request - CopyDataCacheRequest
//
// @return CopyDataCacheResponse
func (client *Client) CopyDataCache(request *CopyDataCacheRequest) (_result *CopyDataCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CopyDataCacheResponse{}
	_body, _err := client.CopyDataCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an elastic container instance.
//
// Description:
//
// When you call the CreateContainerGroup operation to create an elastic container instance, the system automatically creates a service-linked role named AliyunServiceRoleForECI. You can assume the service-linked role to access relevant cloud services such as Elastic Compute Service (ECS) and Virtual Private Cloud (VPC). For more information, see [Elastic Container Instance service-linked role](https://help.aliyun.com/document_detail/212914.html).
//
// When you create an elastic container instance, you can configure features that are related to instances, images, and storage based on your business requirements. For information about parameters configured for the features and the description of the parameters, see the following documents:
//
// **Instances**
//
// You can use one of the following methods to create an elastic container instance:
//
//   - [Specify the number of vCPUs and memory size to create an elastic container instance](https://help.aliyun.com/document_detail/114662.html)
//
//   - [Specify ECS instance types to create an elastic container instance](https://help.aliyun.com/document_detail/114664.html)
//
// Both the preceding creation methods support the following features:
//
//   - [Create a preemptible elastic container instance](https://help.aliyun.com/document_detail/157759.html)
//
//   - [Configure multiple zones](https://help.aliyun.com/document_detail/157290.html)
//
//   - [Configure multiple specifications](https://help.aliyun.com/document_detail/146468.html)
//
//   - [Use tags to manage elastic container instances](https://help.aliyun.com/document_detail/146608.html)
//
// **Images**
//
//   - [Configure a container image](https://help.aliyun.com/document_detail/461311.html)
//
//   - [Use the image cache feature to accelerate the creation of an elastic container instance](https://help.aliyun.com/document_detail/141281.html)
//
//   - [Specify a Container Registry Enterprise Edition instance](https://help.aliyun.com/document_detail/194250.html)
//
//   - [Use self-managed image repositories](https://help.aliyun.com/document_detail/378059.html)
//
// **Networking**
//
//   - [Create and Associate an EIP](https://help.aliyun.com/document_detail/99146.html)
//
//   - [Assign a security group](https://help.aliyun.com/document_detail/176237.html)
//
//   - [Assign an IPv6 address to an elastic container instance](https://help.aliyun.com/document_detail/451282.html)
//
//   - [Configure maximum bandwidth](https://help.aliyun.com/document_detail/190635.html)
//
// **Storage**
//
//   - [Mount a disk volume](https://help.aliyun.com/document_detail/144571.html)
//
//   - [Mount a NAS volume](https://help.aliyun.com/document_detail/464075.html)
//
//   - [Mount an OSS bucket to an elastic container instance as a volume](https://help.aliyun.com/document_detail/464076.html)
//
//   - [Mount an emptyDir volume](https://help.aliyun.com/document_detail/464078.html)
//
//   - [Mount a ConfigFile volume](https://help.aliyun.com/document_detail/464080.html)
//
//   - [Increase the size of the temporary storage space](https://help.aliyun.com/document_detail/204066.html)
//
// **Container configuration**
//
//   - [Configure startup commands and arguments for a container](https://help.aliyun.com/document_detail/94593.html)
//
//   - [Use probes to perform health checks on a container](https://help.aliyun.com/document_detail/99053.html)
//
//   - [Obtain metadata by using environment variables](https://help.aliyun.com/document_detail/141788.html)
//
//   - [Configure a security context for an elastic container instance or a container](https://help.aliyun.com/document_detail/462313.html)
//
//   - [Configure the NTP service](https://help.aliyun.com/document_detail/462768.html)
//
// **Logging and O\\&M**
//
//   - [Use environment variables to configure log collection](https://help.aliyun.com/document_detail/121973.html)
//
//   - [Save core files to volumes](https://help.aliyun.com/document_detail/167801.html)
//
// @param request - CreateContainerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateContainerGroupResponse
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

	if !tea.BoolValue(util.IsUnset(request.ComputeCategory)) {
		query["ComputeCategory"] = request.ComputeCategory
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

	if !tea.BoolValue(util.IsUnset(request.CpuArchitecture)) {
		query["CpuArchitecture"] = request.CpuArchitecture
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

	if !tea.BoolValue(util.IsUnset(request.DataCacheBucket)) {
		query["DataCacheBucket"] = request.DataCacheBucket
	}

	if !tea.BoolValue(util.IsUnset(request.DataCacheBurstingEnabled)) {
		query["DataCacheBurstingEnabled"] = request.DataCacheBurstingEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.DataCachePL)) {
		query["DataCachePL"] = request.DataCachePL
	}

	if !tea.BoolValue(util.IsUnset(request.DataCacheProvisionedIops)) {
		query["DataCacheProvisionedIops"] = request.DataCacheProvisionedIops
	}

	if !tea.BoolValue(util.IsUnset(request.DnsPolicy)) {
		query["DnsPolicy"] = request.DnsPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
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

	if !tea.BoolValue(util.IsUnset(request.FixedIp)) {
		query["FixedIp"] = request.FixedIp
	}

	if !tea.BoolValue(util.IsUnset(request.FixedIpRetainHour)) {
		query["FixedIpRetainHour"] = request.FixedIpRetainHour
	}

	if !tea.BoolValue(util.IsUnset(request.GpuDriverVersion)) {
		query["GpuDriverVersion"] = request.GpuDriverVersion
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

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
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

	if !tea.BoolValue(util.IsUnset(request.PrivateIpAddress)) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
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

// Summary:
//
// Creates an elastic container instance.
//
// Description:
//
// When you call the CreateContainerGroup operation to create an elastic container instance, the system automatically creates a service-linked role named AliyunServiceRoleForECI. You can assume the service-linked role to access relevant cloud services such as Elastic Compute Service (ECS) and Virtual Private Cloud (VPC). For more information, see [Elastic Container Instance service-linked role](https://help.aliyun.com/document_detail/212914.html).
//
// When you create an elastic container instance, you can configure features that are related to instances, images, and storage based on your business requirements. For information about parameters configured for the features and the description of the parameters, see the following documents:
//
// **Instances**
//
// You can use one of the following methods to create an elastic container instance:
//
//   - [Specify the number of vCPUs and memory size to create an elastic container instance](https://help.aliyun.com/document_detail/114662.html)
//
//   - [Specify ECS instance types to create an elastic container instance](https://help.aliyun.com/document_detail/114664.html)
//
// Both the preceding creation methods support the following features:
//
//   - [Create a preemptible elastic container instance](https://help.aliyun.com/document_detail/157759.html)
//
//   - [Configure multiple zones](https://help.aliyun.com/document_detail/157290.html)
//
//   - [Configure multiple specifications](https://help.aliyun.com/document_detail/146468.html)
//
//   - [Use tags to manage elastic container instances](https://help.aliyun.com/document_detail/146608.html)
//
// **Images**
//
//   - [Configure a container image](https://help.aliyun.com/document_detail/461311.html)
//
//   - [Use the image cache feature to accelerate the creation of an elastic container instance](https://help.aliyun.com/document_detail/141281.html)
//
//   - [Specify a Container Registry Enterprise Edition instance](https://help.aliyun.com/document_detail/194250.html)
//
//   - [Use self-managed image repositories](https://help.aliyun.com/document_detail/378059.html)
//
// **Networking**
//
//   - [Create and Associate an EIP](https://help.aliyun.com/document_detail/99146.html)
//
//   - [Assign a security group](https://help.aliyun.com/document_detail/176237.html)
//
//   - [Assign an IPv6 address to an elastic container instance](https://help.aliyun.com/document_detail/451282.html)
//
//   - [Configure maximum bandwidth](https://help.aliyun.com/document_detail/190635.html)
//
// **Storage**
//
//   - [Mount a disk volume](https://help.aliyun.com/document_detail/144571.html)
//
//   - [Mount a NAS volume](https://help.aliyun.com/document_detail/464075.html)
//
//   - [Mount an OSS bucket to an elastic container instance as a volume](https://help.aliyun.com/document_detail/464076.html)
//
//   - [Mount an emptyDir volume](https://help.aliyun.com/document_detail/464078.html)
//
//   - [Mount a ConfigFile volume](https://help.aliyun.com/document_detail/464080.html)
//
//   - [Increase the size of the temporary storage space](https://help.aliyun.com/document_detail/204066.html)
//
// **Container configuration**
//
//   - [Configure startup commands and arguments for a container](https://help.aliyun.com/document_detail/94593.html)
//
//   - [Use probes to perform health checks on a container](https://help.aliyun.com/document_detail/99053.html)
//
//   - [Obtain metadata by using environment variables](https://help.aliyun.com/document_detail/141788.html)
//
//   - [Configure a security context for an elastic container instance or a container](https://help.aliyun.com/document_detail/462313.html)
//
//   - [Configure the NTP service](https://help.aliyun.com/document_detail/462768.html)
//
// **Logging and O\\&M**
//
//   - [Use environment variables to configure log collection](https://help.aliyun.com/document_detail/121973.html)
//
//   - [Save core files to volumes](https://help.aliyun.com/document_detail/167801.html)
//
// @param request - CreateContainerGroupRequest
//
// @return CreateContainerGroupResponse
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

// Summary:
//
// Creates a DataCache.
//
// Description:
//
//	  You are charged for the creation of image caches. We recommend that you learn the relevant billing information in advance. For more information, see [DataCaches](https://help.aliyun.com/document_detail/2503093.html).
//
//		- Before you create an image cache, you must evaluate the size of the data to be cached. If the size of the data exceeds the specified cache size, the image cache fails to be created.
//
//		- When a data cache is being created, the system automatically creates a temporary elastic container instance (ECI) and an enhanced SSD (ESSD) for the data cache. During the creation, do not delete the ECI and ESSD. Otherwise, the data cache fails to be created.
//
//		- When a data cache is being created, a snapshot is generated for the data cache. Do not delete the snapshot. Otherwise, the data cache becomes invalid.
//
// @param request - CreateDataCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataCacheResponse
func (client *Client) CreateDataCacheWithOptions(request *CreateDataCacheRequest, runtime *util.RuntimeOptions) (_result *CreateDataCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bucket)) {
		query["Bucket"] = request.Bucket
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataSource)) {
		query["DataSource"] = request.DataSource
	}

	if !tea.BoolValue(util.IsUnset(request.EipCreateParam)) {
		query["EipCreateParam"] = request.EipCreateParam
	}

	if !tea.BoolValue(util.IsUnset(request.EipInstanceId)) {
		query["EipInstanceId"] = request.EipInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
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

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
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
		Action:      tea.String("CreateDataCache"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a DataCache.
//
// Description:
//
//	  You are charged for the creation of image caches. We recommend that you learn the relevant billing information in advance. For more information, see [DataCaches](https://help.aliyun.com/document_detail/2503093.html).
//
//		- Before you create an image cache, you must evaluate the size of the data to be cached. If the size of the data exceeds the specified cache size, the image cache fails to be created.
//
//		- When a data cache is being created, the system automatically creates a temporary elastic container instance (ECI) and an enhanced SSD (ESSD) for the data cache. During the creation, do not delete the ECI and ESSD. Otherwise, the data cache fails to be created.
//
//		- When a data cache is being created, a snapshot is generated for the data cache. Do not delete the snapshot. Otherwise, the data cache becomes invalid.
//
// @param request - CreateDataCacheRequest
//
// @return CreateDataCacheResponse
func (client *Client) CreateDataCache(request *CreateDataCacheRequest) (_result *CreateDataCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDataCacheResponse{}
	_body, _err := client.CreateDataCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an image cache. The image cache can accelerate image pulling and reduce the instance startup time when you create an elastic container instance later.
//
// Description:
//
//	  **Precautions**
//
//	    	- You are charged for creation of image caches. We recommend that you learn the relevant billing information in advance. For more information about billing of image caches, see [Image caches](https://help.aliyun.com/document_detail/447682.html).
//
//	    	- Before you create an image cache, you must estimate the total size of the images that you want to cache. If the total size of the images exceeds the specified cache size, the image cache cannot be created.
//
//	    	- When an image cache is being created, the system creates an intermediate elastic container instance and an intermediate enhanced SSD (ESSD) at performance level 1 (PL1). Do not delete the intermediate instance and the ESSD while the image cache is being created. If you delete the intermediate instance or the ESSD, the image cache cannot be created.
//
//	    	- A temporary local snapshot and a specific number of regular snapshots are generated during the creation of the image cache. Do not delete these snapshots. If you delete these snapshots, the image cache becomes invalid.
//
//	    	- If you use SDKs, SDK for Java 1.0.10 or later and SDK for Python 1.0.7 or later are supported.
//
//		- **Usage notes**
//
//	    	- For images that are created based on Container Registry Enterprise Edition instances and use custom domain names, if you want to configure password-free access to the image caches, you must use AcrRegistryInfo-related parameters to specify Container Registry instances. When you configure AcrRegistryInfo-related parameters, you must set the AcrRegistryInfo.N.InstanceId parameter.
//
//	    	- If the image cache that you created will be used to create more than 1,000 elastic container instances at a time, we recommend that you use the StandardCopyCount and FlashCopyCount parameters to create multiple temporary local snapshots and regular snapshots of the image. The multiple snapshots are billed based on incremental data. If no incremental data exists on the multiple snapshots, you are not charged for the multiple snapshots.
//
// >  When you call the CreateImageCache operation to create an image cache, the system automatically creates a service-linked role named AliyunServiceRoleForECI. The role is used to access other Alibaba Cloud services such as Elastic Compute Service (ECS) and Virtual Private Cloud (VPC). For more information, see [Elastic Container Instance service-linked role](https://help.aliyun.com/document_detail/212914.html).
//
// @param request - CreateImageCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageCacheResponse
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

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
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

// Summary:
//
// Creates an image cache. The image cache can accelerate image pulling and reduce the instance startup time when you create an elastic container instance later.
//
// Description:
//
//	  **Precautions**
//
//	    	- You are charged for creation of image caches. We recommend that you learn the relevant billing information in advance. For more information about billing of image caches, see [Image caches](https://help.aliyun.com/document_detail/447682.html).
//
//	    	- Before you create an image cache, you must estimate the total size of the images that you want to cache. If the total size of the images exceeds the specified cache size, the image cache cannot be created.
//
//	    	- When an image cache is being created, the system creates an intermediate elastic container instance and an intermediate enhanced SSD (ESSD) at performance level 1 (PL1). Do not delete the intermediate instance and the ESSD while the image cache is being created. If you delete the intermediate instance or the ESSD, the image cache cannot be created.
//
//	    	- A temporary local snapshot and a specific number of regular snapshots are generated during the creation of the image cache. Do not delete these snapshots. If you delete these snapshots, the image cache becomes invalid.
//
//	    	- If you use SDKs, SDK for Java 1.0.10 or later and SDK for Python 1.0.7 or later are supported.
//
//		- **Usage notes**
//
//	    	- For images that are created based on Container Registry Enterprise Edition instances and use custom domain names, if you want to configure password-free access to the image caches, you must use AcrRegistryInfo-related parameters to specify Container Registry instances. When you configure AcrRegistryInfo-related parameters, you must set the AcrRegistryInfo.N.InstanceId parameter.
//
//	    	- If the image cache that you created will be used to create more than 1,000 elastic container instances at a time, we recommend that you use the StandardCopyCount and FlashCopyCount parameters to create multiple temporary local snapshots and regular snapshots of the image. The multiple snapshots are billed based on incremental data. If no incremental data exists on the multiple snapshots, you are not charged for the multiple snapshots.
//
// >  When you call the CreateImageCache operation to create an image cache, the system automatically creates a service-linked role named AliyunServiceRoleForECI. The role is used to access other Alibaba Cloud services such as Elastic Compute Service (ECS) and Virtual Private Cloud (VPC). For more information, see [Elastic Container Instance service-linked role](https://help.aliyun.com/document_detail/212914.html).
//
// @param request - CreateImageCacheRequest
//
// @return CreateImageCacheResponse
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

// Summary:
//
// Creates an O&M task.
//
// Description:
//
// O&M tasks are classified into:
//
//   - coredump: After you enable coredump, the system generates a core dump file when a container unexpectedly stops. You can use the core dump file to analyze the exception and find out the cause of the problem. For more information, see [Enable coredump](https://help.aliyun.com/document_detail/167801.html).
//
//   - tcpdump: After you enable tcpdump, the system captures network packets when a container unexpectedly stops. You can analyze the packets and locate network problems. For more information, see Enable [tcpdump](https://help.aliyun.com/document_detail/429749.html).
//
// @param request - CreateInstanceOpsTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceOpsTaskResponse
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

// Summary:
//
// Creates an O&M task.
//
// Description:
//
// O&M tasks are classified into:
//
//   - coredump: After you enable coredump, the system generates a core dump file when a container unexpectedly stops. You can use the core dump file to analyze the exception and find out the cause of the problem. For more information, see [Enable coredump](https://help.aliyun.com/document_detail/167801.html).
//
//   - tcpdump: After you enable tcpdump, the system captures network packets when a container unexpectedly stops. You can analyze the packets and locate network problems. For more information, see Enable [tcpdump](https://help.aliyun.com/document_detail/429749.html).
//
// @param request - CreateInstanceOpsTaskRequest
//
// @return CreateInstanceOpsTaskResponse
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

// Summary:
//
// Creates a VNode to connect self-managed Kubernetes clusters to elastic container instances.
//
// Description:
//
//	  When you call this operation to create a virtual node, the system automatically creates a service-linked role AliyunServiceRoleForECIVnode. This way, you can use the service-linked role to access relevant cloud services such as Elastic Container Instance, Elastic Compute Service (ECS), and Virtual Private Cloud (VPC). For more information, see [Service-linked role for virtual nodes](https://help.aliyun.com/document_detail/311014.html).
//
//		- You are charged for virtual nodes based on number of virtual nodes that you use. Each virtual node has a resident node, which is equivalent to an ECI instance with 2 vCPU cores and 8 GiB memory. You are charged for virtual nodes based on elastic container instances.
//
// @param request - CreateVirtualNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVirtualNodeResponse
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

// Summary:
//
// Creates a VNode to connect self-managed Kubernetes clusters to elastic container instances.
//
// Description:
//
//	  When you call this operation to create a virtual node, the system automatically creates a service-linked role AliyunServiceRoleForECIVnode. This way, you can use the service-linked role to access relevant cloud services such as Elastic Container Instance, Elastic Compute Service (ECS), and Virtual Private Cloud (VPC). For more information, see [Service-linked role for virtual nodes](https://help.aliyun.com/document_detail/311014.html).
//
//		- You are charged for virtual nodes based on number of virtual nodes that you use. Each virtual node has a resident node, which is equivalent to an ECI instance with 2 vCPU cores and 8 GiB memory. You are charged for virtual nodes based on elastic container instances.
//
// @param request - CreateVirtualNodeRequest
//
// @return CreateVirtualNodeResponse
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

// Summary:
//
// Deletes a container group.
//
// Description:
//
// You can delete a container group that you no longer need. Before you delete a container group, make sure that you understand the lifecycle of container groups. For more information, see [Lifecycle of an elastic container instance](https://help.aliyun.com/document_detail/122385.html).
//
// @param request - DeleteContainerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContainerGroupResponse
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

	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
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

// Summary:
//
// Deletes a container group.
//
// Description:
//
// You can delete a container group that you no longer need. Before you delete a container group, make sure that you understand the lifecycle of container groups. For more information, see [Lifecycle of an elastic container instance](https://help.aliyun.com/document_detail/122385.html).
//
// @param request - DeleteContainerGroupRequest
//
// @return DeleteContainerGroupResponse
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

// Summary:
//
// Deletes a DataCache.
//
// @param request - DeleteDataCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataCacheResponse
func (client *Client) DeleteDataCacheWithOptions(request *DeleteDataCacheRequest, runtime *util.RuntimeOptions) (_result *DeleteDataCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bucket)) {
		query["Bucket"] = request.Bucket
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataCacheId)) {
		query["DataCacheId"] = request.DataCacheId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
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
		Action:      tea.String("DeleteDataCache"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a DataCache.
//
// @param request - DeleteDataCacheRequest
//
// @return DeleteDataCacheResponse
func (client *Client) DeleteDataCache(request *DeleteDataCacheRequest) (_result *DeleteDataCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataCacheResponse{}
	_body, _err := client.DeleteDataCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an image cache.
//
// @param request - DeleteImageCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteImageCacheResponse
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

// Summary:
//
// Deletes an image cache.
//
// @param request - DeleteImageCacheRequest
//
// @return DeleteImageCacheResponse
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

// Summary:
//
// Deletes a virtual node.
//
// @param request - DeleteVirtualNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVirtualNodeResponse
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

// Summary:
//
// Deletes a virtual node.
//
// @param request - DeleteVirtualNodeRequest
//
// @return DeleteVirtualNodeResponse
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

// Summary:
//
// Queries the Elastic Compute Service (ECS) instance families that are available in a specified region and zone.
//
// Description:
//
// When you call the CreateContainerGroup operation to create an elastic container instance, you can use the InstanceType parameter to specify ECS instance types that fit your specific needs. To ensure that the elastic container instance can be created, you can call the DescribeAvailableResource operation to query which ECS instance types and instance families are available in the specified region and zone before you create the elastic container instance.
//
// @param request - DescribeAvailableResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableResourceResponse
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

// Summary:
//
// Queries the Elastic Compute Service (ECS) instance families that are available in a specified region and zone.
//
// Description:
//
// When you call the CreateContainerGroup operation to create an elastic container instance, you can use the InstanceType parameter to specify ECS instance types that fit your specific needs. To ensure that the elastic container instance can be created, you can call the DescribeAvailableResource operation to query which ECS instance types and instance families are available in the specified region and zone before you create the elastic container instance.
//
// @param request - DescribeAvailableResourceRequest
//
// @return DescribeAvailableResourceResponse
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

// Summary:
//
// Queries the details of a CommitContainer task.
//
// @param request - DescribeCommitContainerTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCommitContainerTaskResponse
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

// Summary:
//
// Queries the details of a CommitContainer task.
//
// @param request - DescribeCommitContainerTaskRequest
//
// @return DescribeCommitContainerTaskResponse
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

// Summary:
//
// Queries event information about multiple container groups at a time.
//
// Description:
//
// You can call this operation to query the event information about multiple elastic container instances at a time. By default, the most recent 50 entries of events of each elastic container instance are returned.
//
// @param request - DescribeContainerGroupEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeContainerGroupEventsResponse
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

// Summary:
//
// Queries event information about multiple container groups at a time.
//
// Description:
//
// You can call this operation to query the event information about multiple elastic container instances at a time. By default, the most recent 50 entries of events of each elastic container instance are returned.
//
// @param request - DescribeContainerGroupEventsRequest
//
// @return DescribeContainerGroupEventsResponse
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

// Summary:
//
// Queries the monitoring data of an elastic container instance.
//
// Description:
//
//	  A maximum of 50 monitoring data entries can be returned. If the number of monitoring data entries exceeds this limit, an error message is returned.
//
//		- You can query real-time monitoring data (data generated within the last 5 minutes) and historical data (data generated more than 5 minutes ago). If the time range to query starts or ends later than the current time, historical monitoring data generated more than 5 minutes ago is returned.
//
//		- The elastic container instance whose monitoring data you want to query must be created after April 3, 2019, 15:00 UTC+8.
//
// @param request - DescribeContainerGroupMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeContainerGroupMetricResponse
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

// Summary:
//
// Queries the monitoring data of an elastic container instance.
//
// Description:
//
//	  A maximum of 50 monitoring data entries can be returned. If the number of monitoring data entries exceeds this limit, an error message is returned.
//
//		- You can query real-time monitoring data (data generated within the last 5 minutes) and historical data (data generated more than 5 minutes ago). If the time range to query starts or ends later than the current time, historical monitoring data generated more than 5 minutes ago is returned.
//
//		- The elastic container instance whose monitoring data you want to query must be created after April 3, 2019, 15:00 UTC+8.
//
// @param request - DescribeContainerGroupMetricRequest
//
// @return DescribeContainerGroupMetricResponse
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

// Summary:
//
// Queries the price of an elastic container instance.
//
// Description:
//
//	  When you call this operation, you cannot use resource groups to control the permissions of a RAM user.
//
//		- You can create an elastic container instance by specifying vCPU and memory resource specifications or by specifying ECS instance types. When you call this operation to query the prices of elastic container instances, pass in specifications of the elastic container instances.
//
//	    	- [vCPU and memory specifications](https://help.aliyun.com/document_detail/114662.html).
//
//	    	- [ECS instance types that are supported by Elastic Container Instance](https://help.aliyun.com/document_detail/114664.html).
//
// @param request - DescribeContainerGroupPriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeContainerGroupPriceResponse
func (client *Client) DescribeContainerGroupPriceWithOptions(request *DescribeContainerGroupPriceRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerGroupPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComputeCategory)) {
		query["ComputeCategory"] = request.ComputeCategory
	}

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

// Summary:
//
// Queries the price of an elastic container instance.
//
// Description:
//
//	  When you call this operation, you cannot use resource groups to control the permissions of a RAM user.
//
//		- You can create an elastic container instance by specifying vCPU and memory resource specifications or by specifying ECS instance types. When you call this operation to query the prices of elastic container instances, pass in specifications of the elastic container instances.
//
//	    	- [vCPU and memory specifications](https://help.aliyun.com/document_detail/114662.html).
//
//	    	- [ECS instance types that are supported by Elastic Container Instance](https://help.aliyun.com/document_detail/114664.html).
//
// @param request - DescribeContainerGroupPriceRequest
//
// @return DescribeContainerGroupPriceResponse
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

// Summary:
//
// Queries the statuses of multiple container groups at a time.
//
// @param request - DescribeContainerGroupStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeContainerGroupStatusResponse
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

// Summary:
//
// Queries the statuses of multiple container groups at a time.
//
// @param request - DescribeContainerGroupStatusRequest
//
// @return DescribeContainerGroupStatusResponse
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

// Summary:
//
// Queries information about multiple elastic container instances at a time.
//
// Description:
//
//	  After an elastic container instance is terminated, its underlying computing resources are recycled. By default, other resources, such as elastic IP addresses (EIPs), that are created together with the instance are released together with the instance.
//
//		- The metadata of an instance in the final status (Failed, Succeeded, or Expired) is retained based on the following rules:
//
//	    	- All metadata information is retained within 1 hour since the instance enters the final status.
//
//	    	- One hour after the instance enters the final status, only the latest 100 entries of metadata information in each region are retained.
//
// @param request - DescribeContainerGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeContainerGroupsResponse
func (client *Client) DescribeContainerGroupsWithOptions(request *DescribeContainerGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComputeCategory)) {
		query["ComputeCategory"] = request.ComputeCategory
	}

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

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
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

// Summary:
//
// Queries information about multiple elastic container instances at a time.
//
// Description:
//
//	  After an elastic container instance is terminated, its underlying computing resources are recycled. By default, other resources, such as elastic IP addresses (EIPs), that are created together with the instance are released together with the instance.
//
//		- The metadata of an instance in the final status (Failed, Succeeded, or Expired) is retained based on the following rules:
//
//	    	- All metadata information is retained within 1 hour since the instance enters the final status.
//
//	    	- One hour after the instance enters the final status, only the latest 100 entries of metadata information in each region are retained.
//
// @param request - DescribeContainerGroupsRequest
//
// @return DescribeContainerGroupsResponse
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

// Summary:
//
// Queries the logs of a container in a container group.
//
// @param request - DescribeContainerLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeContainerLogResponse
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

// Summary:
//
// Queries the logs of a container in a container group.
//
// @param request - DescribeContainerLogRequest
//
// @return DescribeContainerLogResponse
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

// Summary:
//
// Queries the information about data caches.
//
// @param request - DescribeDataCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataCachesResponse
func (client *Client) DescribeDataCachesWithOptions(request *DescribeDataCachesRequest, runtime *util.RuntimeOptions) (_result *DescribeDataCachesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bucket)) {
		query["Bucket"] = request.Bucket
	}

	if !tea.BoolValue(util.IsUnset(request.DataCacheId)) {
		query["DataCacheId"] = request.DataCacheId
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

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
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

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataCaches"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about data caches.
//
// @param request - DescribeDataCachesRequest
//
// @return DescribeDataCachesResponse
func (client *Client) DescribeDataCaches(request *DescribeDataCachesRequest) (_result *DescribeDataCachesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataCachesResponse{}
	_body, _err := client.DescribeDataCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about image caches.
//
// @param request - DescribeImageCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImageCachesResponse
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

	if !tea.BoolValue(util.IsUnset(request.ImageFullMatch)) {
		query["ImageFullMatch"] = request.ImageFullMatch
	}

	if !tea.BoolValue(util.IsUnset(request.ImageMatchCountRequest)) {
		query["ImageMatchCountRequest"] = request.ImageMatchCountRequest
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

// Summary:
//
// Queries information about image caches.
//
// @param request - DescribeImageCachesRequest
//
// @return DescribeImageCachesResponse
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

// Summary:
//
// Queries the information about operations and maintenance tasks of an elastic container instance.
//
// @param request - DescribeInstanceOpsRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceOpsRecordsResponse
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

// Summary:
//
// Queries the information about operations and maintenance tasks of an elastic container instance.
//
// @param request - DescribeInstanceOpsRecordsRequest
//
// @return DescribeInstanceOpsRecordsResponse
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

// Summary:
//
// Queries the monitoring data of elastic container instances.
//
// Description:
//
//	  Only the latest entry of monitoring data of each elastic container instance is returned.
//
//		- You can query only the monitoring data of elastic container instances that are created after April 3, 2019 15:00:00 UTC+8.
//
// @param request - DescribeMultiContainerGroupMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMultiContainerGroupMetricResponse
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

// Summary:
//
// Queries the monitoring data of elastic container instances.
//
// Description:
//
//	  Only the latest entry of monitoring data of each elastic container instance is returned.
//
//		- You can query only the monitoring data of elastic container instances that are created after April 3, 2019 15:00:00 UTC+8.
//
// @param request - DescribeMultiContainerGroupMetricRequest
//
// @return DescribeMultiContainerGroupMetricResponse
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

// Summary:
//
// Queries the regions and zones in which Elastic Container Instance is available.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
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

// Summary:
//
// Queries the regions and zones in which Elastic Container Instance is available.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
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

// Summary:
//
// Queries information about virtual nodes.
//
// @param request - DescribeVirtualNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVirtualNodesResponse
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

// Summary:
//
// Queries information about virtual nodes.
//
// @param request - DescribeVirtualNodesRequest
//
// @return DescribeVirtualNodesResponse
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

// Summary:
//
// Runs commands in a container.
//
// @param request - ExecContainerCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecContainerCommandResponse
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

// Summary:
//
// Runs commands in a container.
//
// @param request - ExecContainerCommandRequest
//
// @return ExecContainerCommandResponse
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

// Summary:
//
// Queries the tags that are added to an Elastic Container Instance resource.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Version:     tea.String("2018-08-08"),
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

// Summary:
//
// Queries the tags that are added to an Elastic Container Instance resource.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
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

// Summary:
//
// Queries the used amounts and upper limits of privileges and quotas that you have in a specified region.
//
// Description:
//
// This operation does not support resource group authentication.
//
// @param request - ListUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsageResponse
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

// Summary:
//
// Queries the used amounts and upper limits of privileges and quotas that you have in a specified region.
//
// Description:
//
// This operation does not support resource group authentication.
//
// @param request - ListUsageRequest
//
// @return ListUsageResponse
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

// Summary:
//
// Scales up a disk volume that is mounted to an elastic container instance.
//
// Description:
//
// ## [](#)Usage notes
//
//   - You can scale up volumes by calling this operation. You cannot scale down volumes by calling this operation.
//
//   - Only volumes of the disk type can be scaled up. Volumes of other types cannot be scaled up.
//
// @param request - ResizeContainerGroupVolumeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResizeContainerGroupVolumeResponse
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

// Summary:
//
// Scales up a disk volume that is mounted to an elastic container instance.
//
// Description:
//
// ## [](#)Usage notes
//
//   - You can scale up volumes by calling this operation. You cannot scale down volumes by calling this operation.
//
//   - Only volumes of the disk type can be scaled up. Volumes of other types cannot be scaled up.
//
// @param request - ResizeContainerGroupVolumeRequest
//
// @return ResizeContainerGroupVolumeResponse
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

// Summary:
//
// Restarts an elastic container instance.
//
// Description:
//
//	  Only elastic container instances that are in the Pending or Running state can be restarted. Instances that are in the Succeeded or Failed state cannot be restarted.
//
//		- Elastic container instances that were created before 15:00:00 on March 7, 2019 cannot be restarted.
//
//		- When an elastic container instance is being restarted, its status changes into Restarting.
//
// @param request - RestartContainerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartContainerGroupResponse
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

// Summary:
//
// Restarts an elastic container instance.
//
// Description:
//
//	  Only elastic container instances that are in the Pending or Running state can be restarted. Instances that are in the Succeeded or Failed state cannot be restarted.
//
//		- Elastic container instances that were created before 15:00:00 on March 7, 2019 cannot be restarted.
//
//		- When an elastic container instance is being restarted, its status changes into Restarting.
//
// @param request - RestartContainerGroupRequest
//
// @return RestartContainerGroupResponse
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

// Summary:
//
// Adds tags to an Elastic Container Instance resource.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Version:     tea.String("2018-08-08"),
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

// Summary:
//
// Adds tags to an Elastic Container Instance resource.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
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

// Summary:
//
// Removes tags from an Elastic Container Instance resource.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

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

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from an Elastic Container Instance resource.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an elastic container instance.
//
// Description:
//
//	  Only elastic container instances that are in the Pending or Running state can be updated. After you call this operation to update an elastic container instance, the instance enters the Updating state.
//
//		- If the RestartPolicy parameter is set to Never for the elastic container instance that you are updating, the containers of the instance may fail. Exercise caution if you want to update the kind of instances.
//
// @param request - UpdateContainerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContainerGroupResponse
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

// Summary:
//
// Updates an elastic container instance.
//
// Description:
//
//	  Only elastic container instances that are in the Pending or Running state can be updated. After you call this operation to update an elastic container instance, the instance enters the Updating state.
//
//		- If the RestartPolicy parameter is set to Never for the elastic container instance that you are updating, the containers of the instance may fail. Exercise caution if you want to update the kind of instances.
//
// @param request - UpdateContainerGroupRequest
//
// @return UpdateContainerGroupResponse
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

// Summary:
//
// Updates a data cache.
//
// @param request - UpdateDataCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataCacheResponse
func (client *Client) UpdateDataCacheWithOptions(request *UpdateDataCacheRequest, runtime *util.RuntimeOptions) (_result *UpdateDataCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bucket)) {
		query["Bucket"] = request.Bucket
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataCacheId)) {
		query["DataCacheId"] = request.DataCacheId
	}

	if !tea.BoolValue(util.IsUnset(request.DataSource)) {
		query["DataSource"] = request.DataSource
	}

	if !tea.BoolValue(util.IsUnset(request.EipCreateParam)) {
		query["EipCreateParam"] = request.EipCreateParam
	}

	if !tea.BoolValue(util.IsUnset(request.EipInstanceId)) {
		query["EipInstanceId"] = request.EipInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
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

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
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
		Action:      tea.String("UpdateDataCache"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDataCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data cache.
//
// @param request - UpdateDataCacheRequest
//
// @return UpdateDataCacheResponse
func (client *Client) UpdateDataCache(request *UpdateDataCacheRequest) (_result *UpdateDataCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDataCacheResponse{}
	_body, _err := client.UpdateDataCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an image cache.
//
// Description:
//
// Only image caches that are in the Ready or UpdateFailed state can be updated.
//
// @param request - UpdateImageCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateImageCacheResponse
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

// Summary:
//
// Updates an image cache.
//
// Description:
//
// Only image caches that are in the Ready or UpdateFailed state can be updated.
//
// @param request - UpdateImageCacheRequest
//
// @return UpdateImageCacheResponse
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

// Summary:
//
// Updates attributes of a virtual node.
//
// Description:
//
// ## Usage notes
//
// Only virtual nodes that are in the Ready state can be updated.
//
// @param request - UpdateVirtualNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVirtualNodeResponse
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

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
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

// Summary:
//
// Updates attributes of a virtual node.
//
// Description:
//
// ## Usage notes
//
// Only virtual nodes that are in the Ready state can be updated.
//
// @param request - UpdateVirtualNodeRequest
//
// @return UpdateVirtualNodeResponse
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
