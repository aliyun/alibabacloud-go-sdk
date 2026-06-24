// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartInstanceResponseBody
	GetRequestId() *string
	SetResult(v *RestartInstanceResponseBodyResult) *RestartInstanceResponseBody
	GetResult() *RestartInstanceResponseBodyResult
}

type RestartInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *RestartInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s RestartInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartInstanceResponseBody) GetResult() *RestartInstanceResponseBodyResult {
	return s.Result
}

func (s *RestartInstanceResponseBody) SetRequestId(v string) *RestartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartInstanceResponseBody) SetResult(v *RestartInstanceResponseBodyResult) *RestartInstanceResponseBody {
	s.Result = v
	return s
}

func (s *RestartInstanceResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RestartInstanceResponseBodyResult struct {
	// The time when the instance was created.
	//
	// example:
	//
	// 2020-07-06T10:18:48.662Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The instance name.
	//
	// example:
	//
	// es-cn-abc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The IK dictionary configuration.
	DictList []*RestartInstanceResponseBodyResultDictList `json:"dictList,omitempty" xml:"dictList,omitempty" type:"Repeated"`
	// The internal endpoint of the instance.
	//
	// example:
	//
	// es-cn-nif1q8auz0003****.elasticsearch.aliyuncs.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The instance version.
	//
	// example:
	//
	// 6.7.0_with_X-Pack
	EsVersion *string `json:"esVersion,omitempty" xml:"esVersion,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// es-cn-n6w1o1x0w001c****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The Kibana node configuration.
	KibanaConfiguration *RestartInstanceResponseBodyResultKibanaConfiguration `json:"kibanaConfiguration,omitempty" xml:"kibanaConfiguration,omitempty" type:"Struct"`
	// The Kibana public network access address.
	//
	// example:
	//
	// es-cn-nif1q8auz0003****.kibana.elasticsearch.aliyuncs.com
	KibanaDomain *string `json:"kibanaDomain,omitempty" xml:"kibanaDomain,omitempty"`
	// The public port of Kibana.
	//
	// example:
	//
	// 5601
	KibanaPort *int32 `json:"kibanaPort,omitempty" xml:"kibanaPort,omitempty"`
	// The master node configuration.
	MasterConfiguration *RestartInstanceResponseBodyResultMasterConfiguration `json:"masterConfiguration,omitempty" xml:"masterConfiguration,omitempty" type:"Struct"`
	// The network configuration.
	NetworkConfig *RestartInstanceResponseBodyResultNetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
	// The number of data nodes in the instance.
	//
	// example:
	//
	// 2
	NodeAmount *int32 `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	// The data node configuration.
	NodeSpec *RestartInstanceResponseBodyResultNodeSpec `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	// The billing method of the instance.
	//
	// Valid values: prepaid (subscription) and postpaid (pay-as-you-go).
	//
	// example:
	//
	// postpaid
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The public network access address.
	//
	// example:
	//
	// es-cn-n6w1o1x0w001c****.public.elasticsearch.aliyuncs.com
	PublicDomain *string `json:"publicDomain,omitempty" xml:"publicDomain,omitempty"`
	// The public port.
	//
	// example:
	//
	// 9200
	PublicPort *int32 `json:"publicPort,omitempty" xml:"publicPort,omitempty"`
	// The status of the instance.
	//
	// Valid values: active, activating, inactive, and invalid.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The synonym dictionary configuration.
	SynonymsDicts []*RestartInstanceResponseBodyResultSynonymsDicts `json:"synonymsDicts,omitempty" xml:"synonymsDicts,omitempty" type:"Repeated"`
	// The time when the instance was last updated.
	//
	// example:
	//
	// 2018-07-18T10:10:04.484Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s RestartInstanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *RestartInstanceResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *RestartInstanceResponseBodyResult) GetDictList() []*RestartInstanceResponseBodyResultDictList {
	return s.DictList
}

func (s *RestartInstanceResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *RestartInstanceResponseBodyResult) GetEsVersion() *string {
	return s.EsVersion
}

func (s *RestartInstanceResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RestartInstanceResponseBodyResult) GetKibanaConfiguration() *RestartInstanceResponseBodyResultKibanaConfiguration {
	return s.KibanaConfiguration
}

func (s *RestartInstanceResponseBodyResult) GetKibanaDomain() *string {
	return s.KibanaDomain
}

func (s *RestartInstanceResponseBodyResult) GetKibanaPort() *int32 {
	return s.KibanaPort
}

func (s *RestartInstanceResponseBodyResult) GetMasterConfiguration() *RestartInstanceResponseBodyResultMasterConfiguration {
	return s.MasterConfiguration
}

func (s *RestartInstanceResponseBodyResult) GetNetworkConfig() *RestartInstanceResponseBodyResultNetworkConfig {
	return s.NetworkConfig
}

func (s *RestartInstanceResponseBodyResult) GetNodeAmount() *int32 {
	return s.NodeAmount
}

func (s *RestartInstanceResponseBodyResult) GetNodeSpec() *RestartInstanceResponseBodyResultNodeSpec {
	return s.NodeSpec
}

func (s *RestartInstanceResponseBodyResult) GetPaymentType() *string {
	return s.PaymentType
}

func (s *RestartInstanceResponseBodyResult) GetPublicDomain() *string {
	return s.PublicDomain
}

func (s *RestartInstanceResponseBodyResult) GetPublicPort() *int32 {
	return s.PublicPort
}

func (s *RestartInstanceResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *RestartInstanceResponseBodyResult) GetSynonymsDicts() []*RestartInstanceResponseBodyResultSynonymsDicts {
	return s.SynonymsDicts
}

func (s *RestartInstanceResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *RestartInstanceResponseBodyResult) SetCreatedAt(v string) *RestartInstanceResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetDescription(v string) *RestartInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetDictList(v []*RestartInstanceResponseBodyResultDictList) *RestartInstanceResponseBodyResult {
	s.DictList = v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetDomain(v string) *RestartInstanceResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetEsVersion(v string) *RestartInstanceResponseBodyResult {
	s.EsVersion = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetInstanceId(v string) *RestartInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetKibanaConfiguration(v *RestartInstanceResponseBodyResultKibanaConfiguration) *RestartInstanceResponseBodyResult {
	s.KibanaConfiguration = v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetKibanaDomain(v string) *RestartInstanceResponseBodyResult {
	s.KibanaDomain = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetKibanaPort(v int32) *RestartInstanceResponseBodyResult {
	s.KibanaPort = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetMasterConfiguration(v *RestartInstanceResponseBodyResultMasterConfiguration) *RestartInstanceResponseBodyResult {
	s.MasterConfiguration = v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetNetworkConfig(v *RestartInstanceResponseBodyResultNetworkConfig) *RestartInstanceResponseBodyResult {
	s.NetworkConfig = v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetNodeAmount(v int32) *RestartInstanceResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetNodeSpec(v *RestartInstanceResponseBodyResultNodeSpec) *RestartInstanceResponseBodyResult {
	s.NodeSpec = v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetPaymentType(v string) *RestartInstanceResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetPublicDomain(v string) *RestartInstanceResponseBodyResult {
	s.PublicDomain = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetPublicPort(v int32) *RestartInstanceResponseBodyResult {
	s.PublicPort = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetStatus(v string) *RestartInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetSynonymsDicts(v []*RestartInstanceResponseBodyResultSynonymsDicts) *RestartInstanceResponseBodyResult {
	s.SynonymsDicts = v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetUpdatedAt(v string) *RestartInstanceResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) Validate() error {
	if s.DictList != nil {
		for _, item := range s.DictList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
	if s.SynonymsDicts != nil {
		for _, item := range s.SynonymsDicts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RestartInstanceResponseBodyResultDictList struct {
	// The size of the dictionary file. Unit: bytes.
	//
	// example:
	//
	// 2782602
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// The name of the dictionary file.
	//
	// example:
	//
	// SYSTEM_MAIN.dic
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The source type. Valid values:
	//
	// - OSS: Object Storage Service (OSS). The OSS bucket must have public-read permissions.
	//
	// - ORIGIN: open-source Elasticsearch
	//
	// - UPLOAD: uploaded file.
	//
	// example:
	//
	// ORIGIN
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The dictionary type. Valid values:
	//
	// - STOP: stopword dictionary
	//
	// - MAIN: main dictionary
	//
	// - SYNONYMS: synonym dictionary
	//
	// - ALI_WS: Alibaba dictionary.
	//
	// example:
	//
	// MAIN
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RestartInstanceResponseBodyResultDictList) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceResponseBodyResultDictList) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBodyResultDictList) GetFileSize() *int64 {
	return s.FileSize
}

func (s *RestartInstanceResponseBodyResultDictList) GetName() *string {
	return s.Name
}

func (s *RestartInstanceResponseBodyResultDictList) GetSourceType() *string {
	return s.SourceType
}

func (s *RestartInstanceResponseBodyResultDictList) GetType() *string {
	return s.Type
}

func (s *RestartInstanceResponseBodyResultDictList) SetFileSize(v int64) *RestartInstanceResponseBodyResultDictList {
	s.FileSize = &v
	return s
}

func (s *RestartInstanceResponseBodyResultDictList) SetName(v string) *RestartInstanceResponseBodyResultDictList {
	s.Name = &v
	return s
}

func (s *RestartInstanceResponseBodyResultDictList) SetSourceType(v string) *RestartInstanceResponseBodyResultDictList {
	s.SourceType = &v
	return s
}

func (s *RestartInstanceResponseBodyResultDictList) SetType(v string) *RestartInstanceResponseBodyResultDictList {
	s.Type = &v
	return s
}

func (s *RestartInstanceResponseBodyResultDictList) Validate() error {
	return dara.Validate(s)
}

type RestartInstanceResponseBodyResultKibanaConfiguration struct {
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
	// The node specifications.
	//
	// example:
	//
	// elasticsearch.n4.small
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s RestartInstanceResponseBodyResultKibanaConfiguration) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceResponseBodyResultKibanaConfiguration) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBodyResultKibanaConfiguration) GetAmount() *int32 {
	return s.Amount
}

func (s *RestartInstanceResponseBodyResultKibanaConfiguration) GetDisk() *int32 {
	return s.Disk
}

func (s *RestartInstanceResponseBodyResultKibanaConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *RestartInstanceResponseBodyResultKibanaConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *RestartInstanceResponseBodyResultKibanaConfiguration) SetAmount(v int32) *RestartInstanceResponseBodyResultKibanaConfiguration {
	s.Amount = &v
	return s
}

func (s *RestartInstanceResponseBodyResultKibanaConfiguration) SetDisk(v int32) *RestartInstanceResponseBodyResultKibanaConfiguration {
	s.Disk = &v
	return s
}

func (s *RestartInstanceResponseBodyResultKibanaConfiguration) SetDiskType(v string) *RestartInstanceResponseBodyResultKibanaConfiguration {
	s.DiskType = &v
	return s
}

func (s *RestartInstanceResponseBodyResultKibanaConfiguration) SetSpec(v string) *RestartInstanceResponseBodyResultKibanaConfiguration {
	s.Spec = &v
	return s
}

func (s *RestartInstanceResponseBodyResultKibanaConfiguration) Validate() error {
	return dara.Validate(s)
}

type RestartInstanceResponseBodyResultMasterConfiguration struct {
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
	// The node specifications.
	//
	// example:
	//
	// elasticsearch.sn2ne.large
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s RestartInstanceResponseBodyResultMasterConfiguration) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceResponseBodyResultMasterConfiguration) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBodyResultMasterConfiguration) GetAmount() *int32 {
	return s.Amount
}

func (s *RestartInstanceResponseBodyResultMasterConfiguration) GetDisk() *int32 {
	return s.Disk
}

func (s *RestartInstanceResponseBodyResultMasterConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *RestartInstanceResponseBodyResultMasterConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *RestartInstanceResponseBodyResultMasterConfiguration) SetAmount(v int32) *RestartInstanceResponseBodyResultMasterConfiguration {
	s.Amount = &v
	return s
}

func (s *RestartInstanceResponseBodyResultMasterConfiguration) SetDisk(v int32) *RestartInstanceResponseBodyResultMasterConfiguration {
	s.Disk = &v
	return s
}

func (s *RestartInstanceResponseBodyResultMasterConfiguration) SetDiskType(v string) *RestartInstanceResponseBodyResultMasterConfiguration {
	s.DiskType = &v
	return s
}

func (s *RestartInstanceResponseBodyResultMasterConfiguration) SetSpec(v string) *RestartInstanceResponseBodyResultMasterConfiguration {
	s.Spec = &v
	return s
}

func (s *RestartInstanceResponseBodyResultMasterConfiguration) Validate() error {
	return dara.Validate(s)
}

type RestartInstanceResponseBodyResultNetworkConfig struct {
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
	// vpc-bp16k1dvzxtmagcva****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// The zone where the instance is deployed.
	//
	// example:
	//
	// cn-hangzhou-i
	VsArea *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1k4ec6s7sjdbudw****
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s RestartInstanceResponseBodyResultNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceResponseBodyResultNetworkConfig) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBodyResultNetworkConfig) GetType() *string {
	return s.Type
}

func (s *RestartInstanceResponseBodyResultNetworkConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *RestartInstanceResponseBodyResultNetworkConfig) GetVsArea() *string {
	return s.VsArea
}

func (s *RestartInstanceResponseBodyResultNetworkConfig) GetVswitchId() *string {
	return s.VswitchId
}

func (s *RestartInstanceResponseBodyResultNetworkConfig) SetType(v string) *RestartInstanceResponseBodyResultNetworkConfig {
	s.Type = &v
	return s
}

func (s *RestartInstanceResponseBodyResultNetworkConfig) SetVpcId(v string) *RestartInstanceResponseBodyResultNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *RestartInstanceResponseBodyResultNetworkConfig) SetVsArea(v string) *RestartInstanceResponseBodyResultNetworkConfig {
	s.VsArea = &v
	return s
}

func (s *RestartInstanceResponseBodyResultNetworkConfig) SetVswitchId(v string) *RestartInstanceResponseBodyResultNetworkConfig {
	s.VswitchId = &v
	return s
}

func (s *RestartInstanceResponseBodyResultNetworkConfig) Validate() error {
	return dara.Validate(s)
}

type RestartInstanceResponseBodyResultNodeSpec struct {
	// The storage size of the node. Unit: GB.
	//
	// example:
	//
	// 50
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The storage type of the node.
	//
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The node specifications.
	//
	// example:
	//
	// elasticsearch.n4.small
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s RestartInstanceResponseBodyResultNodeSpec) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceResponseBodyResultNodeSpec) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBodyResultNodeSpec) GetDisk() *int32 {
	return s.Disk
}

func (s *RestartInstanceResponseBodyResultNodeSpec) GetDiskType() *string {
	return s.DiskType
}

func (s *RestartInstanceResponseBodyResultNodeSpec) GetSpec() *string {
	return s.Spec
}

func (s *RestartInstanceResponseBodyResultNodeSpec) SetDisk(v int32) *RestartInstanceResponseBodyResultNodeSpec {
	s.Disk = &v
	return s
}

func (s *RestartInstanceResponseBodyResultNodeSpec) SetDiskType(v string) *RestartInstanceResponseBodyResultNodeSpec {
	s.DiskType = &v
	return s
}

func (s *RestartInstanceResponseBodyResultNodeSpec) SetSpec(v string) *RestartInstanceResponseBodyResultNodeSpec {
	s.Spec = &v
	return s
}

func (s *RestartInstanceResponseBodyResultNodeSpec) Validate() error {
	return dara.Validate(s)
}

type RestartInstanceResponseBodyResultSynonymsDicts struct {
	// The size of the dictionary file. Unit: bytes.
	//
	// example:
	//
	// 2782602
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// The name of the dictionary file.
	//
	// example:
	//
	// SYSTEM_MAIN.dic
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The source type. Valid values:
	//
	// - OSS: Object Storage Service (OSS). The OSS bucket must have public-read permissions.
	//
	// - ORIGIN: open-source Elasticsearch
	//
	// - UPLOAD: uploaded file.
	//
	// example:
	//
	// ORIGIN
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The dictionary type. Valid values:
	//
	// - STOP: stopword dictionary
	//
	// - MAIN: main dictionary
	//
	// - SYNONYMS: synonym dictionary
	//
	// - ALI_WS: Alibaba dictionary.
	//
	// example:
	//
	// STOP
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RestartInstanceResponseBodyResultSynonymsDicts) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceResponseBodyResultSynonymsDicts) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBodyResultSynonymsDicts) GetFileSize() *int64 {
	return s.FileSize
}

func (s *RestartInstanceResponseBodyResultSynonymsDicts) GetName() *string {
	return s.Name
}

func (s *RestartInstanceResponseBodyResultSynonymsDicts) GetSourceType() *string {
	return s.SourceType
}

func (s *RestartInstanceResponseBodyResultSynonymsDicts) GetType() *string {
	return s.Type
}

func (s *RestartInstanceResponseBodyResultSynonymsDicts) SetFileSize(v int64) *RestartInstanceResponseBodyResultSynonymsDicts {
	s.FileSize = &v
	return s
}

func (s *RestartInstanceResponseBodyResultSynonymsDicts) SetName(v string) *RestartInstanceResponseBodyResultSynonymsDicts {
	s.Name = &v
	return s
}

func (s *RestartInstanceResponseBodyResultSynonymsDicts) SetSourceType(v string) *RestartInstanceResponseBodyResultSynonymsDicts {
	s.SourceType = &v
	return s
}

func (s *RestartInstanceResponseBodyResultSynonymsDicts) SetType(v string) *RestartInstanceResponseBodyResultSynonymsDicts {
	s.Type = &v
	return s
}

func (s *RestartInstanceResponseBodyResultSynonymsDicts) Validate() error {
	return dara.Validate(s)
}
