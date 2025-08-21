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
	// The ID of the request.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return results.
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
	return dara.Validate(s)
}

type RestartInstanceResponseBodyResult struct {
	// The time when the instance was created.
	//
	// example:
	//
	// 2020-07-06T10:18:48.662Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// es-cn-abc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The configuration of the IK dictionaries.
	DictList []*RestartInstanceResponseBodyResultDictList `json:"dictList,omitempty" xml:"dictList,omitempty" type:"Repeated"`
	// The intranet access address of the instance.
	//
	// example:
	//
	// es-cn-nif1q8auz0003****.elasticsearch.aliyuncs.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The version of the instance.
	//
	// example:
	//
	// 6.7.0_with_X-Pack
	EsVersion *string `json:"esVersion,omitempty" xml:"esVersion,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// es-cn-n6w1o1x0w001c****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The configuration of Kibana nodes.
	KibanaConfiguration *RestartInstanceResponseBodyResultKibanaConfiguration `json:"kibanaConfiguration,omitempty" xml:"kibanaConfiguration,omitempty" type:"Struct"`
	// The public network access address of Kibana.
	//
	// example:
	//
	// es-cn-nif1q8auz0003****.kibana.elasticsearch.aliyuncs.com
	KibanaDomain *string `json:"kibanaDomain,omitempty" xml:"kibanaDomain,omitempty"`
	// The public port of the Kibana network.
	//
	// example:
	//
	// 5601
	KibanaPort *int32 `json:"kibanaPort,omitempty" xml:"kibanaPort,omitempty"`
	// The configuration of dedicated master nodes.
	MasterConfiguration *RestartInstanceResponseBodyResultMasterConfiguration `json:"masterConfiguration,omitempty" xml:"masterConfiguration,omitempty" type:"Struct"`
	// The network configuration.
	NetworkConfig *RestartInstanceResponseBodyResultNetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
	// The number of data nodes.
	//
	// example:
	//
	// 2
	NodeAmount *int32 `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	// The configuration of data nodes.
	NodeSpec *RestartInstanceResponseBodyResultNodeSpec `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	// The billing method of the created ECS instance.
	//
	// Valid values: prepaid and postpaid.
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
	// The public network port.
	//
	// example:
	//
	// 9200
	PublicPort *int32 `json:"publicPort,omitempty" xml:"publicPort,omitempty"`
	// The state of the cluster.
	//
	// Supported: active (normal), activating (initializing), inactive (blocked), and invalid (expired).
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The configuration of the synonym dictionaries.
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
	return dara.Validate(s)
}

type RestartInstanceResponseBodyResultDictList struct {
	// The size of the Dictionary File. Unit: bytes.
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
	// 	- OSS: OSS open storage (need to ensure that OSS storage space is public readable.)
	//
	// 	- ORIGIN: Open source Elasticsearch
	//
	// 	- UPLOAD: Uploaded files
	//
	// example:
	//
	// ORIGIN
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The type of the dictionary. Valid values:
	//
	// 	- STOP: The STOP word.
	//
	// 	- MAIN: MAIN Dictionary
	//
	// 	- SYNONYMS: SYNONYMS
	//
	// 	- ALI_WS: an Alibaba Dictionary.
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
	// The number of performance metrics.
	//
	// example:
	//
	// 1
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
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The specification of data nodes.
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
	// The number of nodes in the cluster.
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
	// The storage type of the node. This tool only supports cloud_ssd (cloud SSD) disks.
	//
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The node specifications of the cluster.
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
	// The ID of the VPC.
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
	// The ID of the vSwitch associated with the specified VPC.
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
	// The storage space size per data node. Unit: GB.
	//
	// example:
	//
	// 50
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The storage type of the node. Valid values: cloud_ssd and cloud_efficiency.
	//
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The specification of data nodes.
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
	// The size of the Dictionary File. Unit: bytes.
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
	// 	- OSS:OSS open storage (the OSS storage space must be publicly readable.)
	//
	// 	- ORIGIN: open-source Elasticsearch
	//
	// 	- UPLOAD
	//
	// example:
	//
	// ORIGIN
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The type of the dictionary. Valid values:
	//
	// 	- STOP: The STOP word.
	//
	// 	- MAIN: MAIN Dictionary
	//
	// 	- SYNONYMS: SYNONYMS
	//
	// 	- ALI_WS: an Alibaba Dictionary.
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
