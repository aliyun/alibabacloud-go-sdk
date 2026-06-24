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
	// The response headers.
	Headers *ListInstanceResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned results.
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
	if s.Headers != nil {
		if err := s.Headers.Validate(); err != nil {
			return err
		}
	}
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceResponseBodyHeaders struct {
	// The total number of instances.
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
	// Indicates whether the instance contains dedicated master nodes. Valid values:
	//
	// - true: The instance contains dedicated master nodes.
	//
	// - false: The instance does not contain dedicated master nodes.
	//
	// example:
	//
	// false
	AdvancedDedicateMaster *bool `json:"advancedDedicateMaster,omitempty" xml:"advancedDedicateMaster,omitempty"`
	// The deployment mode and architecture type:
	//
	// exclusive: basic management and control
	//
	// public: cloud-native management and control
	//
	// example:
	//
	// public
	ArchType *string `json:"archType,omitempty" xml:"archType,omitempty"`
	// The configuration of client nodes.
	ClientNodeConfiguration *ListInstanceResponseBodyResultClientNodeConfiguration `json:"clientNodeConfiguration,omitempty" xml:"clientNodeConfiguration,omitempty" type:"Struct"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2018-07-13T03:58:07.253Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Indicates whether the instance contains dedicated master nodes (deprecated). Valid values:
	//
	// - true: The instance contains dedicated master nodes.
	//
	// - false: The instance does not contain dedicated master nodes.
	//
	// example:
	//
	// false
	DedicateMaster *bool `json:"dedicateMaster,omitempty" xml:"dedicateMaster,omitempty"`
	// The instance name.
	//
	// example:
	//
	// es-cn-abc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The internal endpoint of the instance.
	//
	// example:
	//
	// es-cn-nif1q8auz0005****.elasticsearch.aliyuncs.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The configuration of elastic data nodes.
	ElasticDataNodeConfiguration *ListInstanceResponseBodyResultElasticDataNodeConfiguration `json:"elasticDataNodeConfiguration,omitempty" xml:"elasticDataNodeConfiguration,omitempty" type:"Struct"`
	// The expiration time of the instance.
	//
	// example:
	//
	// 1715826092044
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The instance version.
	//
	// example:
	//
	// 6.7_with_X-Pack
	EsVersion *string `json:"esVersion,omitempty" xml:"esVersion,omitempty"`
	// The extended configurations of the cluster.
	ExtendConfigs []map[string]interface{} `json:"extendConfigs,omitempty" xml:"extendConfigs,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// es-cn-v641a0ta3000g****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Indicates whether the instance uses the new deployment architecture.
	//
	// example:
	//
	// true
	IsNewDeployment *string `json:"isNewDeployment,omitempty" xml:"isNewDeployment,omitempty"`
	// The configuration of Kibana nodes.
	KibanaConfiguration *ListInstanceResponseBodyResultKibanaConfiguration `json:"kibanaConfiguration,omitempty" xml:"kibanaConfiguration,omitempty" type:"Struct"`
	// The public network access whitelist for Kibana nodes of the cluster.
	KibanaIPWhitelist []*string `json:"kibanaIPWhitelist,omitempty" xml:"kibanaIPWhitelist,omitempty" type:"Repeated"`
	// The private network access whitelist for Kibana nodes of the cluster.
	KibanaPrivateIPWhitelist []*string `json:"kibanaPrivateIPWhitelist,omitempty" xml:"kibanaPrivateIPWhitelist,omitempty" type:"Repeated"`
	// The configuration of master nodes.
	MasterConfiguration *ListInstanceResponseBodyResultMasterConfiguration `json:"masterConfiguration,omitempty" xml:"masterConfiguration,omitempty" type:"Struct"`
	// The network configuration.
	NetworkConfig *ListInstanceResponseBodyResultNetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
	// The number of data nodes in the instance.
	//
	// example:
	//
	// 2
	NodeAmount *int32 `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	// The configuration of data nodes.
	NodeSpec *ListInstanceResponseBodyResultNodeSpec `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	// The billing method of the instance. Valid values:
	//
	// - **prepaid**: subscription
	//
	// - **postpaid**: pay-as-you-go
	//
	// example:
	//
	// postpaid
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The access port of the instance.
	//
	// 	Notice: When the instance is being created or the instance status is abnormal, this value may be empty or 0.
	//
	// example:
	//
	// 9200
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// The status of the pay-as-you-go service that is overlaid on a subscription instance. Valid values:
	//
	// - **active**: normal
	//
	// - **closed**: closed
	//
	// - **indebt**: frozen due to overdue payment
	//
	// example:
	//
	// active
	PostpaidServiceStatus *string `json:"postpaidServiceStatus,omitempty" xml:"postpaidServiceStatus,omitempty"`
	// The private network access whitelist for the Elasticsearch cluster.
	PrivateNetworkIpWhiteList []*string `json:"privateNetworkIpWhiteList,omitempty" xml:"privateNetworkIpWhiteList,omitempty" type:"Repeated"`
	// The access protocol. Valid values: HTTP and HTTPS.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The public network access whitelist for the Elasticsearch cluster.
	PublicIpWhitelist []*string `json:"publicIpWhitelist,omitempty" xml:"publicIpWhitelist,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzvowej3i****
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Indicates whether the instance is a service VPC.
	//
	// example:
	//
	// true
	ServiceVpc *bool `json:"serviceVpc,omitempty" xml:"serviceVpc,omitempty"`
	// The status of the instance. Valid values:
	//
	// - active: normal
	//
	// - activating: taking effect
	//
	// - inactive: frozen
	//
	// - invalid: invalid. The cluster does not exist or is inaccessible. In this case, some fields in the API response may be missing, such as domain and kibanaDomain.
	//
	// - unknown: unknown. The cluster does not exist or is inaccessible. In this case, some fields in the API response may be missing, such as domain and kibanaDomain.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The instance tags.
	Tags []*ListInstanceResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The time when the instance was last updated.
	//
	// example:
	//
	// 2018-07-18T10:10:04.484Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1uag5jj38c****
	VpcInstanceId *string `json:"vpcInstanceId,omitempty" xml:"vpcInstanceId,omitempty"`
	// example:
	//
	// 1
	ZoneCount *int32                                     `json:"zoneCount,omitempty" xml:"zoneCount,omitempty"`
	ZoneInfos []*ListInstanceResponseBodyResultZoneInfos `json:"zoneInfos,omitempty" xml:"zoneInfos,omitempty" type:"Repeated"`
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

func (s *ListInstanceResponseBodyResult) GetZoneCount() *int32 {
	return s.ZoneCount
}

func (s *ListInstanceResponseBodyResult) GetZoneInfos() []*ListInstanceResponseBodyResultZoneInfos {
	return s.ZoneInfos
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

func (s *ListInstanceResponseBodyResult) SetZoneCount(v int32) *ListInstanceResponseBodyResult {
	s.ZoneCount = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetZoneInfos(v []*ListInstanceResponseBodyResultZoneInfos) *ListInstanceResponseBodyResult {
	s.ZoneInfos = v
	return s
}

func (s *ListInstanceResponseBodyResult) Validate() error {
	if s.ClientNodeConfiguration != nil {
		if err := s.ClientNodeConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.ElasticDataNodeConfiguration != nil {
		if err := s.ElasticDataNodeConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.KibanaConfiguration != nil {
		if err := s.KibanaConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.MasterConfiguration != nil {
		if err := s.MasterConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkConfig != nil {
		if err := s.NetworkConfig.Validate(); err != nil {
			return err
		}
	}
	if s.NodeSpec != nil {
		if err := s.NodeSpec.Validate(); err != nil {
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
	if s.ZoneInfos != nil {
		for _, item := range s.ZoneInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceResponseBodyResultClientNodeConfiguration struct {
	// The number of nodes.
	//
	// example:
	//
	// 3
	Amount *int32 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The storage size of the node. Unit: GB.
	//
	// example:
	//
	// 20
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The storage type of the node. Only ultra disks (cloud_efficiency) are supported.
	//
	// example:
	//
	// cloud_efficiency
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The node specifications. For more information about the specifications, see [Product specifications](https://help.aliyun.com/document_detail/271718.html).
	//
	// example:
	//
	// elasticsearch.sn2ne.large
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
	// The description of node specifications.
	//
	// example:
	//
	// 1C 2G
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
	// The storage size of the node. Unit: GB.
	//
	// example:
	//
	// 20
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// Indicates whether disk encryption is enabled for the node. Valid values:
	//
	// - true: Disk encryption is enabled.
	//
	// - false: Disk encryption is not enabled.
	//
	// example:
	//
	// true
	DiskEncryption *bool `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	// The storage type of the node. Valid values:
	//
	// - cloud_ssd: standard SSD
	//
	// - cloud_essd: enhanced SSD (ESSD)
	//
	// - cloud_efficiency: ultra disk
	//
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The node specifications. For more information about the specifications, see [Product specifications](https://help.aliyun.com/document_detail/271718.html).
	//
	// example:
	//
	// elasticsearch.sn2ne.large
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
	// The description of node specifications.
	//
	// example:
	//
	// 1C 2G
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
	// The number of nodes.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The storage size of the node. Unit: GB.
	//
	// example:
	//
	// 20
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The storage type of the node.
	//
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The node specifications. For more information about the specifications, see [Product specifications](https://help.aliyun.com/document_detail/271718.html).
	//
	// example:
	//
	// elasticsearch.n4.small
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
	// The description of node specifications.
	//
	// example:
	//
	// 1C 2G
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
	// The number of nodes.
	//
	// example:
	//
	// 3
	Amount *int32 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The storage size of the node. Unit: GB.
	//
	// example:
	//
	// 20
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The storage type of the node. Only cloud_ssd (standard SSD) is supported.
	//
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The node specifications. For more information about the specifications, see [Product specifications](https://help.aliyun.com/document_detail/271718.html).
	//
	// example:
	//
	// elasticsearch.sn2ne.large
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
	// The description of node specifications.
	//
	// example:
	//
	// 1C 2G
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
	// The network type. Only Virtual Private Cloud (VPC) is supported.
	//
	// example:
	//
	// vpc
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-abc
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// The zone where the instance is deployed.
	//
	// example:
	//
	// cn-hangzhou-e
	VsArea *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-def
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
	// The whitelist group settings for the public and private networks of the cluster.
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
	if s.WhiteIpGroupList != nil {
		for _, item := range s.WhiteIpGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceResponseBodyResultNetworkConfigWhiteIpGroupList struct {
	// The group name.
	//
	// example:
	//
	// default
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The IP address whitelist.
	Ips []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
	// The network type. PRIVATE_ES: Elasticsearch private network. PUBLIC_KIBANA: Kibana public network. PUBLIC_ES: Elasticsearch public network. PRIVATE_KIBANA: Kibana private network.
	//
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
	// The storage size of the node. Unit: GB.
	//
	// example:
	//
	// 50
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// Indicates whether disk encryption is used. Valid values:
	//
	// - true: Disk encryption is used.
	//
	// - false: Disk encryption is not used.
	//
	// example:
	//
	// false
	DiskEncryption *bool `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	// The storage type of the node. Valid values:
	//
	// - cloud_ssd: standard SSD
	//
	// - cloud_efficiency: ultra disk
	//
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The performance level of the ESSD. This parameter is required when diskType is cloud_essd. Valid values: PL1, PL2, and PL3.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
	// The node specifications. For more information about the specifications, see [Product specifications](https://help.aliyun.com/document_detail/271718.html).
	//
	// example:
	//
	// elasticsearch.n4.small
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
	// The description of node specifications.
	//
	// example:
	//
	// 1C 2G
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
	// The tag key.
	//
	// example:
	//
	// env
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value.
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

type ListInstanceResponseBodyResultZoneInfos struct {
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ListInstanceResponseBodyResultZoneInfos) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyResultZoneInfos) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResultZoneInfos) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceResponseBodyResultZoneInfos) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListInstanceResponseBodyResultZoneInfos) SetStatus(v string) *ListInstanceResponseBodyResultZoneInfos {
	s.Status = &v
	return s
}

func (s *ListInstanceResponseBodyResultZoneInfos) SetZoneId(v string) *ListInstanceResponseBodyResultZoneInfos {
	s.ZoneId = &v
	return s
}

func (s *ListInstanceResponseBodyResultZoneInfos) Validate() error {
	return dara.Validate(s)
}
