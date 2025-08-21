// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListInstanceResponseBodyHeaders) *ListInstanceResponseBody
	GetHeaders() *ListInstanceResponseBodyHeaders
	SetRequestId(v string) *ListInstanceResponseBody
	GetRequestId() *string
	SetResult(v []*ListInstanceResponseBodyResult) *ListInstanceResponseBody
	GetResult() []*ListInstanceResponseBodyResult
}

type ListInstanceResponseBody struct {
	// The status of the instance. Valid values:
	//
	// 	- active: normal
	//
	// 	- activating: taking effect
	//
	// 	- inactive: frozen
	//
	// 	- invalid: invalid
	Headers *ListInstanceResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The time when the node is created.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether it is a service VPC.
	Result []*ListInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBody) GetHeaders() *ListInstanceResponseBodyHeaders {
	return s.Headers
}

func (s *ListInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceResponseBody) GetResult() []*ListInstanceResponseBodyResult {
	return s.Result
}

func (s *ListInstanceResponseBody) SetHeaders(v *ListInstanceResponseBodyHeaders) *ListInstanceResponseBody {
	s.Headers = v
	return s
}

func (s *ListInstanceResponseBody) SetRequestId(v string) *ListInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceResponseBody) SetResult(v []*ListInstanceResponseBodyResult) *ListInstanceResponseBody {
	s.Result = v
	return s
}

func (s *ListInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstanceResponseBodyHeaders struct {
	// Specifies whether to include dedicated master nodes (obsolete). Valid values:
	//
	// 	- true: The files contain data that is dumped to the IA storage medium.
	//
	// 	- false: The files do not contain data that is dumped to the IA storage medium.
	//
	// example:
	//
	// 10
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListInstanceResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyHeaders) GetXTotalCount() *int32 {
	return s.XTotalCount
}

func (s *ListInstanceResponseBodyHeaders) SetXTotalCount(v int32) *ListInstanceResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

func (s *ListInstanceResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListInstanceResponseBodyResult struct {
	// The billing method of the instance. Valid values:
	//
	// 	- **prepaid**: subscription
	//
	// 	- **postpaid**: pay-as-you-go
	//
	// example:
	//
	// false
	AdvancedDedicateMaster *bool   `json:"advancedDedicateMaster,omitempty" xml:"advancedDedicateMaster,omitempty"`
	ArchType               *string `json:"archType,omitempty" xml:"archType,omitempty"`
	// The instance type of the node. For more information, see [Specifications](https://help.aliyun.com/document_detail/271718.html).
	ClientNodeConfiguration *ListInstanceResponseBodyResultClientNodeConfiguration `json:"clientNodeConfiguration,omitempty" xml:"clientNodeConfiguration,omitempty" type:"Struct"`
	// The status of the pay-as-you-go service that is overlaid on a subscription instance. Valid values:
	//
	// 	- **active**: normal
	//
	// 	- **closed**: Close
	//
	// 	- **indebt**: Overdue payments are frozen
	//
	// example:
	//
	// 2018-07-13T03:58:07.253Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The edition of the dedicated KMS instance.
	//
	// example:
	//
	// false
	DedicateMaster *bool `json:"dedicateMaster,omitempty" xml:"dedicateMaster,omitempty"`
	// The key of the tag.
	//
	// example:
	//
	// es-cn-abc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Domain      *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The configuration of Kibana nodes.
	ElasticDataNodeConfiguration *ListInstanceResponseBodyResultElasticDataNodeConfiguration `json:"elasticDataNodeConfiguration,omitempty" xml:"elasticDataNodeConfiguration,omitempty" type:"Struct"`
	EndTime                      *int64                                                      `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 6.7_with_X-Pack
	EsVersion *string `json:"esVersion,omitempty" xml:"esVersion,omitempty"`
	// The configurations of elastic data nodes.
	ExtendConfigs []map[string]interface{} `json:"extendConfigs,omitempty" xml:"extendConfigs,omitempty" type:"Repeated"`
	// The instance type of the node. For more information, see [Specifications](https://help.aliyun.com/document_detail/271718.html).
	//
	// example:
	//
	// es-cn-v641a0ta3000g****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The configuration of cluster extension parameters.
	//
	// example:
	//
	// true
	IsNewDeployment *string `json:"isNewDeployment,omitempty" xml:"isNewDeployment,omitempty"`
	// The instance type of the node. For more information, see [Specifications](https://help.aliyun.com/document_detail/271718.html).
	KibanaConfiguration      *ListInstanceResponseBodyResultKibanaConfiguration `json:"kibanaConfiguration,omitempty" xml:"kibanaConfiguration,omitempty" type:"Struct"`
	KibanaIPWhitelist        []*string                                          `json:"kibanaIPWhitelist,omitempty" xml:"kibanaIPWhitelist,omitempty" type:"Repeated"`
	KibanaPrivateIPWhitelist []*string                                          `json:"kibanaPrivateIPWhitelist,omitempty" xml:"kibanaPrivateIPWhitelist,omitempty" type:"Repeated"`
	// The VPC ID of the cluster.
	MasterConfiguration *ListInstanceResponseBodyResultMasterConfiguration `json:"masterConfiguration,omitempty" xml:"masterConfiguration,omitempty" type:"Struct"`
	// The instance type of the node. For more information, see [Specifications](https://help.aliyun.com/document_detail/271718.html).
	NetworkConfig *ListInstanceResponseBodyResultNetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
	// The ID of the resource group.
	//
	// example:
	//
	// 2
	NodeAmount *int32 `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	// The VPC ID of the cluster.
	NodeSpec *ListInstanceResponseBodyResultNodeSpec `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	// The time when the instance was last updated.
	//
	// example:
	//
	// postpaid
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	Port        *string `json:"port,omitempty" xml:"port,omitempty"`
	// The tags of the instance. Each tag is a key-value pair.
	//
	// example:
	//
	// active
	PostpaidServiceStatus     *string   `json:"postpaidServiceStatus,omitempty" xml:"postpaidServiceStatus,omitempty"`
	PrivateNetworkIpWhiteList []*string `json:"privateNetworkIpWhiteList,omitempty" xml:"privateNetworkIpWhiteList,omitempty" type:"Repeated"`
	Protocol                  *string   `json:"protocol,omitempty" xml:"protocol,omitempty"`
	PublicIpWhitelist         []*string `json:"publicIpWhitelist,omitempty" xml:"publicIpWhitelist,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// example:
	//
	// rg-aekzvowej3i****
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Specifies whether to deploy the new architecture.
	//
	// example:
	//
	// true
	ServiceVpc *bool `json:"serviceVpc,omitempty" xml:"serviceVpc,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The number of nodes.
	Tags []*ListInstanceResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// Coordination node configuration.
	//
	// example:
	//
	// 2018-07-18T10:10:04.484Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// vpc-bp1uag5jj38c****
	VpcInstanceId *string `json:"vpcInstanceId,omitempty" xml:"vpcInstanceId,omitempty"`
}

func (s ListInstanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResult) GetAdvancedDedicateMaster() *bool {
	return s.AdvancedDedicateMaster
}

func (s *ListInstanceResponseBodyResult) GetArchType() *string {
	return s.ArchType
}

func (s *ListInstanceResponseBodyResult) GetClientNodeConfiguration() *ListInstanceResponseBodyResultClientNodeConfiguration {
	return s.ClientNodeConfiguration
}

func (s *ListInstanceResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListInstanceResponseBodyResult) GetDedicateMaster() *bool {
	return s.DedicateMaster
}

func (s *ListInstanceResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListInstanceResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *ListInstanceResponseBodyResult) GetElasticDataNodeConfiguration() *ListInstanceResponseBodyResultElasticDataNodeConfiguration {
	return s.ElasticDataNodeConfiguration
}

func (s *ListInstanceResponseBodyResult) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListInstanceResponseBodyResult) GetEsVersion() *string {
	return s.EsVersion
}

func (s *ListInstanceResponseBodyResult) GetExtendConfigs() []map[string]interface{} {
	return s.ExtendConfigs
}

func (s *ListInstanceResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceResponseBodyResult) GetIsNewDeployment() *string {
	return s.IsNewDeployment
}

func (s *ListInstanceResponseBodyResult) GetKibanaConfiguration() *ListInstanceResponseBodyResultKibanaConfiguration {
	return s.KibanaConfiguration
}

func (s *ListInstanceResponseBodyResult) GetKibanaIPWhitelist() []*string {
	return s.KibanaIPWhitelist
}

func (s *ListInstanceResponseBodyResult) GetKibanaPrivateIPWhitelist() []*string {
	return s.KibanaPrivateIPWhitelist
}

func (s *ListInstanceResponseBodyResult) GetMasterConfiguration() *ListInstanceResponseBodyResultMasterConfiguration {
	return s.MasterConfiguration
}

func (s *ListInstanceResponseBodyResult) GetNetworkConfig() *ListInstanceResponseBodyResultNetworkConfig {
	return s.NetworkConfig
}

func (s *ListInstanceResponseBodyResult) GetNodeAmount() *int32 {
	return s.NodeAmount
}

func (s *ListInstanceResponseBodyResult) GetNodeSpec() *ListInstanceResponseBodyResultNodeSpec {
	return s.NodeSpec
}

func (s *ListInstanceResponseBodyResult) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListInstanceResponseBodyResult) GetPort() *string {
	return s.Port
}

func (s *ListInstanceResponseBodyResult) GetPostpaidServiceStatus() *string {
	return s.PostpaidServiceStatus
}

func (s *ListInstanceResponseBodyResult) GetPrivateNetworkIpWhiteList() []*string {
	return s.PrivateNetworkIpWhiteList
}

func (s *ListInstanceResponseBodyResult) GetProtocol() *string {
	return s.Protocol
}

func (s *ListInstanceResponseBodyResult) GetPublicIpWhitelist() []*string {
	return s.PublicIpWhitelist
}

func (s *ListInstanceResponseBodyResult) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstanceResponseBodyResult) GetServiceVpc() *bool {
	return s.ServiceVpc
}

func (s *ListInstanceResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceResponseBodyResult) GetTags() []*ListInstanceResponseBodyResultTags {
	return s.Tags
}

func (s *ListInstanceResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListInstanceResponseBodyResult) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *ListInstanceResponseBodyResult) SetAdvancedDedicateMaster(v bool) *ListInstanceResponseBodyResult {
	s.AdvancedDedicateMaster = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetArchType(v string) *ListInstanceResponseBodyResult {
	s.ArchType = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetClientNodeConfiguration(v *ListInstanceResponseBodyResultClientNodeConfiguration) *ListInstanceResponseBodyResult {
	s.ClientNodeConfiguration = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetCreatedAt(v string) *ListInstanceResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetDedicateMaster(v bool) *ListInstanceResponseBodyResult {
	s.DedicateMaster = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetDescription(v string) *ListInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetDomain(v string) *ListInstanceResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetElasticDataNodeConfiguration(v *ListInstanceResponseBodyResultElasticDataNodeConfiguration) *ListInstanceResponseBodyResult {
	s.ElasticDataNodeConfiguration = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetEndTime(v int64) *ListInstanceResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetEsVersion(v string) *ListInstanceResponseBodyResult {
	s.EsVersion = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetExtendConfigs(v []map[string]interface{}) *ListInstanceResponseBodyResult {
	s.ExtendConfigs = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetInstanceId(v string) *ListInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetIsNewDeployment(v string) *ListInstanceResponseBodyResult {
	s.IsNewDeployment = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetKibanaConfiguration(v *ListInstanceResponseBodyResultKibanaConfiguration) *ListInstanceResponseBodyResult {
	s.KibanaConfiguration = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetKibanaIPWhitelist(v []*string) *ListInstanceResponseBodyResult {
	s.KibanaIPWhitelist = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetKibanaPrivateIPWhitelist(v []*string) *ListInstanceResponseBodyResult {
	s.KibanaPrivateIPWhitelist = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetMasterConfiguration(v *ListInstanceResponseBodyResultMasterConfiguration) *ListInstanceResponseBodyResult {
	s.MasterConfiguration = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetNetworkConfig(v *ListInstanceResponseBodyResultNetworkConfig) *ListInstanceResponseBodyResult {
	s.NetworkConfig = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetNodeAmount(v int32) *ListInstanceResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetNodeSpec(v *ListInstanceResponseBodyResultNodeSpec) *ListInstanceResponseBodyResult {
	s.NodeSpec = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetPaymentType(v string) *ListInstanceResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetPort(v string) *ListInstanceResponseBodyResult {
	s.Port = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetPostpaidServiceStatus(v string) *ListInstanceResponseBodyResult {
	s.PostpaidServiceStatus = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetPrivateNetworkIpWhiteList(v []*string) *ListInstanceResponseBodyResult {
	s.PrivateNetworkIpWhiteList = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetProtocol(v string) *ListInstanceResponseBodyResult {
	s.Protocol = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetPublicIpWhitelist(v []*string) *ListInstanceResponseBodyResult {
	s.PublicIpWhitelist = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetResourceGroupId(v string) *ListInstanceResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetServiceVpc(v bool) *ListInstanceResponseBodyResult {
	s.ServiceVpc = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetStatus(v string) *ListInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetTags(v []*ListInstanceResponseBodyResultTags) *ListInstanceResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetUpdatedAt(v string) *ListInstanceResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetVpcInstanceId(v string) *ListInstanceResponseBodyResult {
	s.VpcInstanceId = &v
	return s
}

func (s *ListInstanceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListInstanceResponseBodyResultClientNodeConfiguration struct {
	// The size of the node storage space. Unit: GB.
	//
	// example:
	//
	// 3
	Amount *int32 `json:"amount,omitempty" xml:"amount,omitempty"`
	// Specifies whether to enable disk encryption for the node. Valid values:
	//
	// 	- true: enables instant image cache.
	//
	// 	- false: disables reuse of image cache layers.
	//
	// example:
	//
	// 20
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The storage type of the node. Valid values:
	//
	// 	- cloud_ssd: SSD.
	//
	// 	- cloud_essd: ESSD.
	//
	// 	- cloud_efficiency: ultra disk
	//
	// example:
	//
	// cloud_efficiency
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// elasticsearch.sn2ne.large
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
	SpecInfo *string `json:"specInfo,omitempty" xml:"specInfo,omitempty"`
}

func (s ListInstanceResponseBodyResultClientNodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyResultClientNodeConfiguration) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResultClientNodeConfiguration) GetAmount() *int32 {
	return s.Amount
}

func (s *ListInstanceResponseBodyResultClientNodeConfiguration) GetDisk() *int32 {
	return s.Disk
}

func (s *ListInstanceResponseBodyResultClientNodeConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *ListInstanceResponseBodyResultClientNodeConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *ListInstanceResponseBodyResultClientNodeConfiguration) GetSpecInfo() *string {
	return s.SpecInfo
}

func (s *ListInstanceResponseBodyResultClientNodeConfiguration) SetAmount(v int32) *ListInstanceResponseBodyResultClientNodeConfiguration {
	s.Amount = &v
	return s
}

func (s *ListInstanceResponseBodyResultClientNodeConfiguration) SetDisk(v int32) *ListInstanceResponseBodyResultClientNodeConfiguration {
	s.Disk = &v
	return s
}

func (s *ListInstanceResponseBodyResultClientNodeConfiguration) SetDiskType(v string) *ListInstanceResponseBodyResultClientNodeConfiguration {
	s.DiskType = &v
	return s
}

func (s *ListInstanceResponseBodyResultClientNodeConfiguration) SetSpec(v string) *ListInstanceResponseBodyResultClientNodeConfiguration {
	s.Spec = &v
	return s
}

func (s *ListInstanceResponseBodyResultClientNodeConfiguration) SetSpecInfo(v string) *ListInstanceResponseBodyResultClientNodeConfiguration {
	s.SpecInfo = &v
	return s
}

func (s *ListInstanceResponseBodyResultClientNodeConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListInstanceResponseBodyResultElasticDataNodeConfiguration struct {
	// The number of nodes.
	//
	// example:
	//
	// 3
	Amount *int32 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The size of the node storage space. Unit: GB.
	//
	// example:
	//
	// 20
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The storage type of the node.
	//
	// example:
	//
	// true
	DiskEncryption *bool `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	// The configuration of dedicated master nodes.
	//
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The instance type of the node. For more information, see [Specifications](https://help.aliyun.com/document_detail/271718.html).
	//
	// example:
	//
	// elasticsearch.sn2ne.large
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
	SpecInfo *string `json:"specInfo,omitempty" xml:"specInfo,omitempty"`
}

func (s ListInstanceResponseBodyResultElasticDataNodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyResultElasticDataNodeConfiguration) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) GetAmount() *int32 {
	return s.Amount
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) GetDisk() *int32 {
	return s.Disk
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) GetDiskEncryption() *bool {
	return s.DiskEncryption
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) GetSpecInfo() *string {
	return s.SpecInfo
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) SetAmount(v int32) *ListInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.Amount = &v
	return s
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) SetDisk(v int32) *ListInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.Disk = &v
	return s
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) SetDiskEncryption(v bool) *ListInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.DiskEncryption = &v
	return s
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) SetDiskType(v string) *ListInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.DiskType = &v
	return s
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) SetSpec(v string) *ListInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.Spec = &v
	return s
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) SetSpecInfo(v string) *ListInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.SpecInfo = &v
	return s
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListInstanceResponseBodyResultKibanaConfiguration struct {
	// The size of the node storage space. Unit: GB.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The storage type of the node. Only cloud_ssd(SSD cloud disk) is supported.
	//
	// example:
	//
	// 20
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The network configurations.
	//
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// elasticsearch.n4.small
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
	SpecInfo *string `json:"specInfo,omitempty" xml:"specInfo,omitempty"`
}

func (s ListInstanceResponseBodyResultKibanaConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyResultKibanaConfiguration) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResultKibanaConfiguration) GetAmount() *int32 {
	return s.Amount
}

func (s *ListInstanceResponseBodyResultKibanaConfiguration) GetDisk() *int32 {
	return s.Disk
}

func (s *ListInstanceResponseBodyResultKibanaConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *ListInstanceResponseBodyResultKibanaConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *ListInstanceResponseBodyResultKibanaConfiguration) GetSpecInfo() *string {
	return s.SpecInfo
}

func (s *ListInstanceResponseBodyResultKibanaConfiguration) SetAmount(v int32) *ListInstanceResponseBodyResultKibanaConfiguration {
	s.Amount = &v
	return s
}

func (s *ListInstanceResponseBodyResultKibanaConfiguration) SetDisk(v int32) *ListInstanceResponseBodyResultKibanaConfiguration {
	s.Disk = &v
	return s
}

func (s *ListInstanceResponseBodyResultKibanaConfiguration) SetDiskType(v string) *ListInstanceResponseBodyResultKibanaConfiguration {
	s.DiskType = &v
	return s
}

func (s *ListInstanceResponseBodyResultKibanaConfiguration) SetSpec(v string) *ListInstanceResponseBodyResultKibanaConfiguration {
	s.Spec = &v
	return s
}

func (s *ListInstanceResponseBodyResultKibanaConfiguration) SetSpecInfo(v string) *ListInstanceResponseBodyResultKibanaConfiguration {
	s.SpecInfo = &v
	return s
}

func (s *ListInstanceResponseBodyResultKibanaConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListInstanceResponseBodyResultMasterConfiguration struct {
	// The network type. Only Virtual Private Cloud (VPC) is supported.
	//
	// example:
	//
	// 3
	Amount *int32 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The vSwitch ID of the cluster.
	//
	// example:
	//
	// 20
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The configuration of data nodes.
	//
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The zone where the cluster resides.
	//
	// example:
	//
	// elasticsearch.sn2ne.large
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
	SpecInfo *string `json:"specInfo,omitempty" xml:"specInfo,omitempty"`
}

func (s ListInstanceResponseBodyResultMasterConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyResultMasterConfiguration) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResultMasterConfiguration) GetAmount() *int32 {
	return s.Amount
}

func (s *ListInstanceResponseBodyResultMasterConfiguration) GetDisk() *int32 {
	return s.Disk
}

func (s *ListInstanceResponseBodyResultMasterConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *ListInstanceResponseBodyResultMasterConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *ListInstanceResponseBodyResultMasterConfiguration) GetSpecInfo() *string {
	return s.SpecInfo
}

func (s *ListInstanceResponseBodyResultMasterConfiguration) SetAmount(v int32) *ListInstanceResponseBodyResultMasterConfiguration {
	s.Amount = &v
	return s
}

func (s *ListInstanceResponseBodyResultMasterConfiguration) SetDisk(v int32) *ListInstanceResponseBodyResultMasterConfiguration {
	s.Disk = &v
	return s
}

func (s *ListInstanceResponseBodyResultMasterConfiguration) SetDiskType(v string) *ListInstanceResponseBodyResultMasterConfiguration {
	s.DiskType = &v
	return s
}

func (s *ListInstanceResponseBodyResultMasterConfiguration) SetSpec(v string) *ListInstanceResponseBodyResultMasterConfiguration {
	s.Spec = &v
	return s
}

func (s *ListInstanceResponseBodyResultMasterConfiguration) SetSpecInfo(v string) *ListInstanceResponseBodyResultMasterConfiguration {
	s.SpecInfo = &v
	return s
}

func (s *ListInstanceResponseBodyResultMasterConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListInstanceResponseBodyResultNetworkConfig struct {
	// The storage type of the node. Valid values:
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_efficiency: ultra disk
	//
	// example:
	//
	// vpc
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The storage space of the node. Unit: GB.
	//
	// example:
	//
	// vpc-abc
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// Specifies whether to use disk encryption. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// cn-hangzhou-e
	VsArea *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	// The performance level of the ESSD. This parameter is required when the diskType parameter is set to cloud_essd. Valid values: PL1, PL2, and PL3.
	//
	// example:
	//
	// vsw-def
	VswitchId        *string                                                        `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
	WhiteIpGroupList []*ListInstanceResponseBodyResultNetworkConfigWhiteIpGroupList `json:"whiteIpGroupList,omitempty" xml:"whiteIpGroupList,omitempty" type:"Repeated"`
}

func (s ListInstanceResponseBodyResultNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyResultNetworkConfig) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResultNetworkConfig) GetType() *string {
	return s.Type
}

func (s *ListInstanceResponseBodyResultNetworkConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *ListInstanceResponseBodyResultNetworkConfig) GetVsArea() *string {
	return s.VsArea
}

func (s *ListInstanceResponseBodyResultNetworkConfig) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ListInstanceResponseBodyResultNetworkConfig) GetWhiteIpGroupList() []*ListInstanceResponseBodyResultNetworkConfigWhiteIpGroupList {
	return s.WhiteIpGroupList
}

func (s *ListInstanceResponseBodyResultNetworkConfig) SetType(v string) *ListInstanceResponseBodyResultNetworkConfig {
	s.Type = &v
	return s
}

func (s *ListInstanceResponseBodyResultNetworkConfig) SetVpcId(v string) *ListInstanceResponseBodyResultNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *ListInstanceResponseBodyResultNetworkConfig) SetVsArea(v string) *ListInstanceResponseBodyResultNetworkConfig {
	s.VsArea = &v
	return s
}

func (s *ListInstanceResponseBodyResultNetworkConfig) SetVswitchId(v string) *ListInstanceResponseBodyResultNetworkConfig {
	s.VswitchId = &v
	return s
}

func (s *ListInstanceResponseBodyResultNetworkConfig) SetWhiteIpGroupList(v []*ListInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) *ListInstanceResponseBodyResultNetworkConfig {
	s.WhiteIpGroupList = v
	return s
}

func (s *ListInstanceResponseBodyResultNetworkConfig) Validate() error {
	return dara.Validate(s)
}

type ListInstanceResponseBodyResultNetworkConfigWhiteIpGroupList struct {
	// example:
	//
	// default
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
	// example:
	//
	// PUBLIC_KIBANA
	WhiteIpType *string `json:"whiteIpType,omitempty" xml:"whiteIpType,omitempty"`
}

func (s ListInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) GetGroupName() *string {
	return s.GroupName
}

func (s *ListInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) GetIps() []*string {
	return s.Ips
}

func (s *ListInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) GetWhiteIpType() *string {
	return s.WhiteIpType
}

func (s *ListInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) SetGroupName(v string) *ListInstanceResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.GroupName = &v
	return s
}

func (s *ListInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) SetIps(v []*string) *ListInstanceResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.Ips = v
	return s
}

func (s *ListInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) SetWhiteIpType(v string) *ListInstanceResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.WhiteIpType = &v
	return s
}

func (s *ListInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) Validate() error {
	return dara.Validate(s)
}

type ListInstanceResponseBodyResultNodeSpec struct {
	// example:
	//
	// 50
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// example:
	//
	// false
	DiskEncryption *bool `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
	// example:
	//
	// elasticsearch.n4.small
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
	SpecInfo *string `json:"specInfo,omitempty" xml:"specInfo,omitempty"`
}

func (s ListInstanceResponseBodyResultNodeSpec) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyResultNodeSpec) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResultNodeSpec) GetDisk() *int32 {
	return s.Disk
}

func (s *ListInstanceResponseBodyResultNodeSpec) GetDiskEncryption() *bool {
	return s.DiskEncryption
}

func (s *ListInstanceResponseBodyResultNodeSpec) GetDiskType() *string {
	return s.DiskType
}

func (s *ListInstanceResponseBodyResultNodeSpec) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *ListInstanceResponseBodyResultNodeSpec) GetSpec() *string {
	return s.Spec
}

func (s *ListInstanceResponseBodyResultNodeSpec) GetSpecInfo() *string {
	return s.SpecInfo
}

func (s *ListInstanceResponseBodyResultNodeSpec) SetDisk(v int32) *ListInstanceResponseBodyResultNodeSpec {
	s.Disk = &v
	return s
}

func (s *ListInstanceResponseBodyResultNodeSpec) SetDiskEncryption(v bool) *ListInstanceResponseBodyResultNodeSpec {
	s.DiskEncryption = &v
	return s
}

func (s *ListInstanceResponseBodyResultNodeSpec) SetDiskType(v string) *ListInstanceResponseBodyResultNodeSpec {
	s.DiskType = &v
	return s
}

func (s *ListInstanceResponseBodyResultNodeSpec) SetPerformanceLevel(v string) *ListInstanceResponseBodyResultNodeSpec {
	s.PerformanceLevel = &v
	return s
}

func (s *ListInstanceResponseBodyResultNodeSpec) SetSpec(v string) *ListInstanceResponseBodyResultNodeSpec {
	s.Spec = &v
	return s
}

func (s *ListInstanceResponseBodyResultNodeSpec) SetSpecInfo(v string) *ListInstanceResponseBodyResultNodeSpec {
	s.SpecInfo = &v
	return s
}

func (s *ListInstanceResponseBodyResultNodeSpec) Validate() error {
	return dara.Validate(s)
}

type ListInstanceResponseBodyResultTags struct {
	// The size of the node storage space. Unit: GB.
	//
	// example:
	//
	// env
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The storage type of the node. Only ultra disks (cloud_efficiency) are supported.
	//
	// example:
	//
	// dev
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListInstanceResponseBodyResultTags) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResultTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListInstanceResponseBodyResultTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListInstanceResponseBodyResultTags) SetTagKey(v string) *ListInstanceResponseBodyResultTags {
	s.TagKey = &v
	return s
}

func (s *ListInstanceResponseBodyResultTags) SetTagValue(v string) *ListInstanceResponseBodyResultTags {
	s.TagValue = &v
	return s
}

func (s *ListInstanceResponseBodyResultTags) Validate() error {
	return dara.Validate(s)
}
