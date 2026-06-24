// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MoveResourceGroupResponseBody
	GetRequestId() *string
	SetResult(v *MoveResourceGroupResponseBodyResult) *MoveResourceGroupResponseBody
	GetResult() *MoveResourceGroupResponseBodyResult
}

type MoveResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *MoveResourceGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveResourceGroupResponseBody) GetResult() *MoveResourceGroupResponseBodyResult {
	return s.Result
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveResourceGroupResponseBody) SetResult(v *MoveResourceGroupResponseBodyResult) *MoveResourceGroupResponseBody {
	s.Result = v
	return s
}

func (s *MoveResourceGroupResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MoveResourceGroupResponseBodyResult struct {
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
	DictList []*MoveResourceGroupResponseBodyResultDictList `json:"dictList,omitempty" xml:"dictList,omitempty" type:"Repeated"`
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
	// The configuration of Kibana nodes.
	KibanaConfiguration *MoveResourceGroupResponseBodyResultKibanaConfiguration `json:"kibanaConfiguration,omitempty" xml:"kibanaConfiguration,omitempty" type:"Struct"`
	// The public network access address for Kibana.
	//
	// example:
	//
	// es-cn-nif1q8auz0003****.kibana.elasticsearch.aliyuncs.com
	KibanaDomain *string `json:"kibanaDomain,omitempty" xml:"kibanaDomain,omitempty"`
	// The public port for Kibana.
	//
	// example:
	//
	// 5601
	KibanaPort *int32 `json:"kibanaPort,omitempty" xml:"kibanaPort,omitempty"`
	// The configuration of master nodes.
	MasterConfiguration *MoveResourceGroupResponseBodyResultMasterConfiguration `json:"masterConfiguration,omitempty" xml:"masterConfiguration,omitempty" type:"Struct"`
	// The network configuration.
	NetworkConfig *MoveResourceGroupResponseBodyResultNetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
	// The number of data nodes of the instance.
	//
	// example:
	//
	// 2
	NodeAmount *int32 `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	// The configuration of data nodes.
	NodeSpec *MoveResourceGroupResponseBodyResultNodeSpec `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	// The billing method of the instance. Valid values:
	//
	// - prepaid: subscription.
	//
	// - postpaid: pay-as-you-go.
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
	// The status of the instance. Valid values:
	//
	// - active: Normal.
	//
	// - activating: Taking effect.
	//
	// - inactive: Frozen.
	//
	// - invalid: Expired.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The synonym dictionary configuration.
	SynonymsDicts []*MoveResourceGroupResponseBodyResultSynonymsDicts `json:"synonymsDicts,omitempty" xml:"synonymsDicts,omitempty" type:"Repeated"`
	// The time when the instance was last updated.
	//
	// example:
	//
	// 2018-07-18T10:10:04.484Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s MoveResourceGroupResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *MoveResourceGroupResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *MoveResourceGroupResponseBodyResult) GetDictList() []*MoveResourceGroupResponseBodyResultDictList {
	return s.DictList
}

func (s *MoveResourceGroupResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *MoveResourceGroupResponseBodyResult) GetEsVersion() *string {
	return s.EsVersion
}

func (s *MoveResourceGroupResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MoveResourceGroupResponseBodyResult) GetKibanaConfiguration() *MoveResourceGroupResponseBodyResultKibanaConfiguration {
	return s.KibanaConfiguration
}

func (s *MoveResourceGroupResponseBodyResult) GetKibanaDomain() *string {
	return s.KibanaDomain
}

func (s *MoveResourceGroupResponseBodyResult) GetKibanaPort() *int32 {
	return s.KibanaPort
}

func (s *MoveResourceGroupResponseBodyResult) GetMasterConfiguration() *MoveResourceGroupResponseBodyResultMasterConfiguration {
	return s.MasterConfiguration
}

func (s *MoveResourceGroupResponseBodyResult) GetNetworkConfig() *MoveResourceGroupResponseBodyResultNetworkConfig {
	return s.NetworkConfig
}

func (s *MoveResourceGroupResponseBodyResult) GetNodeAmount() *int32 {
	return s.NodeAmount
}

func (s *MoveResourceGroupResponseBodyResult) GetNodeSpec() *MoveResourceGroupResponseBodyResultNodeSpec {
	return s.NodeSpec
}

func (s *MoveResourceGroupResponseBodyResult) GetPaymentType() *string {
	return s.PaymentType
}

func (s *MoveResourceGroupResponseBodyResult) GetPublicDomain() *string {
	return s.PublicDomain
}

func (s *MoveResourceGroupResponseBodyResult) GetPublicPort() *int32 {
	return s.PublicPort
}

func (s *MoveResourceGroupResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *MoveResourceGroupResponseBodyResult) GetSynonymsDicts() []*MoveResourceGroupResponseBodyResultSynonymsDicts {
	return s.SynonymsDicts
}

func (s *MoveResourceGroupResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *MoveResourceGroupResponseBodyResult) SetCreatedAt(v string) *MoveResourceGroupResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetDescription(v string) *MoveResourceGroupResponseBodyResult {
	s.Description = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetDictList(v []*MoveResourceGroupResponseBodyResultDictList) *MoveResourceGroupResponseBodyResult {
	s.DictList = v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetDomain(v string) *MoveResourceGroupResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetEsVersion(v string) *MoveResourceGroupResponseBodyResult {
	s.EsVersion = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetInstanceId(v string) *MoveResourceGroupResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetKibanaConfiguration(v *MoveResourceGroupResponseBodyResultKibanaConfiguration) *MoveResourceGroupResponseBodyResult {
	s.KibanaConfiguration = v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetKibanaDomain(v string) *MoveResourceGroupResponseBodyResult {
	s.KibanaDomain = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetKibanaPort(v int32) *MoveResourceGroupResponseBodyResult {
	s.KibanaPort = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetMasterConfiguration(v *MoveResourceGroupResponseBodyResultMasterConfiguration) *MoveResourceGroupResponseBodyResult {
	s.MasterConfiguration = v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetNetworkConfig(v *MoveResourceGroupResponseBodyResultNetworkConfig) *MoveResourceGroupResponseBodyResult {
	s.NetworkConfig = v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetNodeAmount(v int32) *MoveResourceGroupResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetNodeSpec(v *MoveResourceGroupResponseBodyResultNodeSpec) *MoveResourceGroupResponseBodyResult {
	s.NodeSpec = v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetPaymentType(v string) *MoveResourceGroupResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetPublicDomain(v string) *MoveResourceGroupResponseBodyResult {
	s.PublicDomain = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetPublicPort(v int32) *MoveResourceGroupResponseBodyResult {
	s.PublicPort = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetStatus(v string) *MoveResourceGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetSynonymsDicts(v []*MoveResourceGroupResponseBodyResultSynonymsDicts) *MoveResourceGroupResponseBodyResult {
	s.SynonymsDicts = v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetUpdatedAt(v string) *MoveResourceGroupResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) Validate() error {
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

type MoveResourceGroupResponseBodyResultDictList struct {
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
	// - OSS: OSS open storage. Make sure that the OSS bucket has public-read permission.
	//
	// - ORIGIN: retains the previously uploaded dictionary.
	//
	// example:
	//
	// ORIGIN
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The dictionary type. Valid values:
	//
	// - STOP: stopword.
	//
	// - MAIN: main dictionary.
	//
	// - SYNONYMS: synonym dictionary.
	//
	// - ALI_WS: Alibaba dictionary.
	//
	// example:
	//
	// MAIN
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MoveResourceGroupResponseBodyResultDictList) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceGroupResponseBodyResultDictList) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBodyResultDictList) GetFileSize() *int64 {
	return s.FileSize
}

func (s *MoveResourceGroupResponseBodyResultDictList) GetName() *string {
	return s.Name
}

func (s *MoveResourceGroupResponseBodyResultDictList) GetSourceType() *string {
	return s.SourceType
}

func (s *MoveResourceGroupResponseBodyResultDictList) GetType() *string {
	return s.Type
}

func (s *MoveResourceGroupResponseBodyResultDictList) SetFileSize(v int64) *MoveResourceGroupResponseBodyResultDictList {
	s.FileSize = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultDictList) SetName(v string) *MoveResourceGroupResponseBodyResultDictList {
	s.Name = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultDictList) SetSourceType(v string) *MoveResourceGroupResponseBodyResultDictList {
	s.SourceType = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultDictList) SetType(v string) *MoveResourceGroupResponseBodyResultDictList {
	s.Type = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultDictList) Validate() error {
	return dara.Validate(s)
}

type MoveResourceGroupResponseBodyResultKibanaConfiguration struct {
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

func (s MoveResourceGroupResponseBodyResultKibanaConfiguration) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceGroupResponseBodyResultKibanaConfiguration) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBodyResultKibanaConfiguration) GetAmount() *int32 {
	return s.Amount
}

func (s *MoveResourceGroupResponseBodyResultKibanaConfiguration) GetDisk() *int32 {
	return s.Disk
}

func (s *MoveResourceGroupResponseBodyResultKibanaConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *MoveResourceGroupResponseBodyResultKibanaConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *MoveResourceGroupResponseBodyResultKibanaConfiguration) SetAmount(v int32) *MoveResourceGroupResponseBodyResultKibanaConfiguration {
	s.Amount = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultKibanaConfiguration) SetDisk(v int32) *MoveResourceGroupResponseBodyResultKibanaConfiguration {
	s.Disk = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultKibanaConfiguration) SetDiskType(v string) *MoveResourceGroupResponseBodyResultKibanaConfiguration {
	s.DiskType = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultKibanaConfiguration) SetSpec(v string) *MoveResourceGroupResponseBodyResultKibanaConfiguration {
	s.Spec = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultKibanaConfiguration) Validate() error {
	return dara.Validate(s)
}

type MoveResourceGroupResponseBodyResultMasterConfiguration struct {
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
	// elasticsearch.sn2ne.large
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s MoveResourceGroupResponseBodyResultMasterConfiguration) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceGroupResponseBodyResultMasterConfiguration) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBodyResultMasterConfiguration) GetAmount() *int32 {
	return s.Amount
}

func (s *MoveResourceGroupResponseBodyResultMasterConfiguration) GetDisk() *int32 {
	return s.Disk
}

func (s *MoveResourceGroupResponseBodyResultMasterConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *MoveResourceGroupResponseBodyResultMasterConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *MoveResourceGroupResponseBodyResultMasterConfiguration) SetAmount(v int32) *MoveResourceGroupResponseBodyResultMasterConfiguration {
	s.Amount = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultMasterConfiguration) SetDisk(v int32) *MoveResourceGroupResponseBodyResultMasterConfiguration {
	s.Disk = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultMasterConfiguration) SetDiskType(v string) *MoveResourceGroupResponseBodyResultMasterConfiguration {
	s.DiskType = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultMasterConfiguration) SetSpec(v string) *MoveResourceGroupResponseBodyResultMasterConfiguration {
	s.Spec = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultMasterConfiguration) Validate() error {
	return dara.Validate(s)
}

type MoveResourceGroupResponseBodyResultNetworkConfig struct {
	// The network type. Only VPC is supported.
	//
	// example:
	//
	// vpc
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-bp16k1dvzxtmagcva****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// The zone where the instance resides.
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

func (s MoveResourceGroupResponseBodyResultNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceGroupResponseBodyResultNetworkConfig) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBodyResultNetworkConfig) GetType() *string {
	return s.Type
}

func (s *MoveResourceGroupResponseBodyResultNetworkConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *MoveResourceGroupResponseBodyResultNetworkConfig) GetVsArea() *string {
	return s.VsArea
}

func (s *MoveResourceGroupResponseBodyResultNetworkConfig) GetVswitchId() *string {
	return s.VswitchId
}

func (s *MoveResourceGroupResponseBodyResultNetworkConfig) SetType(v string) *MoveResourceGroupResponseBodyResultNetworkConfig {
	s.Type = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultNetworkConfig) SetVpcId(v string) *MoveResourceGroupResponseBodyResultNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultNetworkConfig) SetVsArea(v string) *MoveResourceGroupResponseBodyResultNetworkConfig {
	s.VsArea = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultNetworkConfig) SetVswitchId(v string) *MoveResourceGroupResponseBodyResultNetworkConfig {
	s.VswitchId = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultNetworkConfig) Validate() error {
	return dara.Validate(s)
}

type MoveResourceGroupResponseBodyResultNodeSpec struct {
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

func (s MoveResourceGroupResponseBodyResultNodeSpec) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceGroupResponseBodyResultNodeSpec) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBodyResultNodeSpec) GetDisk() *int32 {
	return s.Disk
}

func (s *MoveResourceGroupResponseBodyResultNodeSpec) GetDiskType() *string {
	return s.DiskType
}

func (s *MoveResourceGroupResponseBodyResultNodeSpec) GetSpec() *string {
	return s.Spec
}

func (s *MoveResourceGroupResponseBodyResultNodeSpec) SetDisk(v int32) *MoveResourceGroupResponseBodyResultNodeSpec {
	s.Disk = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultNodeSpec) SetDiskType(v string) *MoveResourceGroupResponseBodyResultNodeSpec {
	s.DiskType = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultNodeSpec) SetSpec(v string) *MoveResourceGroupResponseBodyResultNodeSpec {
	s.Spec = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultNodeSpec) Validate() error {
	return dara.Validate(s)
}

type MoveResourceGroupResponseBodyResultSynonymsDicts struct {
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
	// - OSS: OSS open storage. Make sure that the OSS bucket has public-read permission.
	//
	// - ORIGIN: retains the previously uploaded dictionary.
	//
	// example:
	//
	// ORIGIN
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The dictionary type. Valid values:
	//
	// - STOP: stopword.
	//
	// - MAIN: main dictionary.
	//
	// - SYNONYMS: synonym dictionary.
	//
	// - ALI_WS: Alibaba dictionary.
	//
	// example:
	//
	// STOP
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MoveResourceGroupResponseBodyResultSynonymsDicts) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceGroupResponseBodyResultSynonymsDicts) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBodyResultSynonymsDicts) GetFileSize() *int64 {
	return s.FileSize
}

func (s *MoveResourceGroupResponseBodyResultSynonymsDicts) GetName() *string {
	return s.Name
}

func (s *MoveResourceGroupResponseBodyResultSynonymsDicts) GetSourceType() *string {
	return s.SourceType
}

func (s *MoveResourceGroupResponseBodyResultSynonymsDicts) GetType() *string {
	return s.Type
}

func (s *MoveResourceGroupResponseBodyResultSynonymsDicts) SetFileSize(v int64) *MoveResourceGroupResponseBodyResultSynonymsDicts {
	s.FileSize = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultSynonymsDicts) SetName(v string) *MoveResourceGroupResponseBodyResultSynonymsDicts {
	s.Name = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultSynonymsDicts) SetSourceType(v string) *MoveResourceGroupResponseBodyResultSynonymsDicts {
	s.SourceType = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultSynonymsDicts) SetType(v string) *MoveResourceGroupResponseBodyResultSynonymsDicts {
	s.Type = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultSynonymsDicts) Validate() error {
	return dara.Validate(s)
}
