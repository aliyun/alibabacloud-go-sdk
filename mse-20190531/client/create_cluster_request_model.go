// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateClusterRequest
	GetAcceptLanguage() *string
	SetChargeType(v string) *CreateClusterRequest
	GetChargeType() *string
	SetClusterSpecification(v string) *CreateClusterRequest
	GetClusterSpecification() *string
	SetClusterType(v string) *CreateClusterRequest
	GetClusterType() *string
	SetClusterVersion(v string) *CreateClusterRequest
	GetClusterVersion() *string
	SetConnectionType(v string) *CreateClusterRequest
	GetConnectionType() *string
	SetDiskType(v string) *CreateClusterRequest
	GetDiskType() *string
	SetEipEnabled(v bool) *CreateClusterRequest
	GetEipEnabled() *bool
	SetInstanceCount(v int32) *CreateClusterRequest
	GetInstanceCount() *int32
	SetInstanceName(v string) *CreateClusterRequest
	GetInstanceName() *string
	SetMseVersion(v string) *CreateClusterRequest
	GetMseVersion() *string
	SetNetType(v string) *CreateClusterRequest
	GetNetType() *string
	SetPrivateSlbSpecification(v string) *CreateClusterRequest
	GetPrivateSlbSpecification() *string
	SetPubNetworkFlow(v string) *CreateClusterRequest
	GetPubNetworkFlow() *string
	SetPubSlbSpecification(v string) *CreateClusterRequest
	GetPubSlbSpecification() *string
	SetRegion(v string) *CreateClusterRequest
	GetRegion() *string
	SetRequestPars(v string) *CreateClusterRequest
	GetRequestPars() *string
	SetResourceGroupId(v string) *CreateClusterRequest
	GetResourceGroupId() *string
	SetSecurityGroupType(v string) *CreateClusterRequest
	GetSecurityGroupType() *string
	SetTag(v []*CreateClusterRequestTag) *CreateClusterRequest
	GetTag() []*CreateClusterRequestTag
	SetVSwitchId(v string) *CreateClusterRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateClusterRequest
	GetVpcId() *string
}

type CreateClusterRequest struct {
	// The language type of the returned information:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Billing method, including PREPAY (Subscription) and POSTPAY (Pay-As-You-Go).
	//
	// This parameter is ignored for the Serverless edition.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Engine specifications, with the following values:
	//
	// [Professional Edition]
	//
	// - `MSE_SC_2_4_60_c`: 2 cores, 4GB
	//
	// - `MSE_SC_1_2_60_c`: 1 core, 2GB
	//
	// - `MSE_SC_4_8_60_c`: 4 cores, 8GB
	//
	// - `MSE_SC_8_16_60_c`: 8 cores, 16GB
	//
	// - `MSE_SC_16_32_60_c`: 16 cores, 32GB
	//
	// [Developer Edition]
	//
	// - `MSE_SC_1_2_60_c`: 1 core, 2GB
	//
	// - `MSE_SC_2_4_60_c`: 2 cores, 4GB
	//
	// [Serverless Edition]
	//
	// Ignore this parameter, or you can fill in `MSE_SC_SERVERLESS`.
	//
	// This parameter is required.
	//
	// example:
	//
	// MSE_SC_2_4_60_c
	ClusterSpecification *string `json:"ClusterSpecification,omitempty" xml:"ClusterSpecification,omitempty"`
	// Cluster type, including ZooKeeper, Nacos-Ans.
	//
	// This parameter is required.
	//
	// example:
	//
	// Nacos-Ans
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// Cluster version, with the following values:
	//
	// [Professional Edition]
	//
	// - `NACOS_2_0_0`: indicates Nacos 2.x.x version.
	//
	// - `ZooKeeper_3_8_0`: indicates ZooKeeper 3.8.x version.
	//
	// [Developer Edition]
	//
	// - `NACOS_2_0_0`: indicates Nacos 2.x version.
	//
	// - `ZooKeeper_3_8_0`: indicates ZooKeeper 3.8.x version.
	//
	// [Serverless Edition]
	//
	// - `NACOS_2_0_0`: indicates Nacos 2.x version.
	//
	// - `ZooKeeper_3_8_0`: indicates ZooKeeper 3.8.x version.
	//
	// This parameter is required.
	//
	// example:
	//
	// NACOS_2_0_0
	ClusterVersion *string `json:"ClusterVersion,omitempty" xml:"ClusterVersion,omitempty"`
	// Network access type, `slb` or `single_eni`; some regions\\" Developer Edition only support the `single_eni` type.
	//
	// example:
	//
	// slb
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// Deprecated
	//
	// No longer in use
	//
	// example:
	//
	// alicloud-disk-ssd
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// Valid when `ConnectionType` is `single_eni`, indicating whether to enable public network access (Elastic IP).
	//
	// if can be null:
	// false
	EipEnabled *bool `json:"EipEnabled,omitempty" xml:"EipEnabled,omitempty"`
	// Number of instance nodes, with a range limit of 1 to 9.
	//
	// [Professional Edition]
	//
	// - The number of instances must be 3 or more and must be an odd number.
	//
	// [Developer Edition]
	//
	// - The number of instances can only be 1.
	//
	// [Serverless Edition]
	//
	// Ignore this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// Custom instance name
	//
	// example:
	//
	// tanshuyingtest001
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Required unless under special circumstances, with the following values:
	//
	// - `mse_pro`: indicates Professional Edition.
	//
	// - `mse_dev`: indicates Developer Edition.
	//
	// - `mse_serverless`: indicates Serverless Edition.
	//
	// example:
	//
	// mse_pro
	MseVersion *string `json:"MseVersion,omitempty" xml:"MseVersion,omitempty"`
	// Network type, with the following values:
	//
	// - `privatenet`: indicates a private network.
	//
	// - `pubnet`: indicates a public network.
	//
	// This parameter is required.
	//
	// example:
	//
	// privatenet
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// Deprecated
	//
	// No longer in use
	//
	// example:
	//
	// slb.s1.small
	PrivateSlbSpecification *string `json:"PrivateSlbSpecification,omitempty" xml:"PrivateSlbSpecification,omitempty"`
	// Valid when `ConnectionType` is `slb`. 0 indicates no public network access SLB creation, and values above 1 indicate a fixed bandwidth for public network access SLB; unit: Mbps.
	//
	// Value range: 0~5000.
	//
	// example:
	//
	// 0
	PubNetworkFlow *string `json:"PubNetworkFlow,omitempty" xml:"PubNetworkFlow,omitempty"`
	// Deprecated
	//
	// No longer in use
	//
	// example:
	//
	// slb.s1.small
	PubSlbSpecification *string `json:"PubSlbSpecification,omitempty" xml:"PubSlbSpecification,omitempty"`
	// The region where the cluster is located, including but not limited to the following regions:
	//
	// - `cn-hangzhou`: Hangzhou
	//
	// - `cn-beijing`: Beijing
	//
	// - `cn-shanghai`: Shanghai
	//
	// - `cn-zhangjiakou`: Zhangjiakou
	//
	// - `cn-shenzhen`: Shenzhen
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Extended request parameters, in JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	// Resource group ID. For more details about the resource group, see [Basic Information of Resource Group](https://help.aliyun.com/document_detail/457230.html).
	//
	// example:
	//
	// rg-aekzcqmoay3dlyq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Valid when `ConnectionType` is `single_eni`, indicating the security group type of the instance.
	//
	// example:
	//
	// enterprise
	SecurityGroupType *string `json:"SecurityGroupType,omitempty" xml:"SecurityGroupType,omitempty"`
	// List of tags to be added. Contains up to 20 items.
	Tag []*CreateClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Switch ID.
	//
	// example:
	//
	// vsw-bp17opt4v18sto39k****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID.
	//
	// example:
	//
	// vpc-bp1t50e045b5g7i3p****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateClusterRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateClusterRequest) GetClusterSpecification() *string {
	return s.ClusterSpecification
}

func (s *CreateClusterRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *CreateClusterRequest) GetClusterVersion() *string {
	return s.ClusterVersion
}

func (s *CreateClusterRequest) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *CreateClusterRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *CreateClusterRequest) GetEipEnabled() *bool {
	return s.EipEnabled
}

func (s *CreateClusterRequest) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *CreateClusterRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateClusterRequest) GetMseVersion() *string {
	return s.MseVersion
}

func (s *CreateClusterRequest) GetNetType() *string {
	return s.NetType
}

func (s *CreateClusterRequest) GetPrivateSlbSpecification() *string {
	return s.PrivateSlbSpecification
}

func (s *CreateClusterRequest) GetPubNetworkFlow() *string {
	return s.PubNetworkFlow
}

func (s *CreateClusterRequest) GetPubSlbSpecification() *string {
	return s.PubSlbSpecification
}

func (s *CreateClusterRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateClusterRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *CreateClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateClusterRequest) GetSecurityGroupType() *string {
	return s.SecurityGroupType
}

func (s *CreateClusterRequest) GetTag() []*CreateClusterRequestTag {
	return s.Tag
}

func (s *CreateClusterRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateClusterRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateClusterRequest) SetAcceptLanguage(v string) *CreateClusterRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateClusterRequest) SetChargeType(v string) *CreateClusterRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateClusterRequest) SetClusterSpecification(v string) *CreateClusterRequest {
	s.ClusterSpecification = &v
	return s
}

func (s *CreateClusterRequest) SetClusterType(v string) *CreateClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterRequest) SetClusterVersion(v string) *CreateClusterRequest {
	s.ClusterVersion = &v
	return s
}

func (s *CreateClusterRequest) SetConnectionType(v string) *CreateClusterRequest {
	s.ConnectionType = &v
	return s
}

func (s *CreateClusterRequest) SetDiskType(v string) *CreateClusterRequest {
	s.DiskType = &v
	return s
}

func (s *CreateClusterRequest) SetEipEnabled(v bool) *CreateClusterRequest {
	s.EipEnabled = &v
	return s
}

func (s *CreateClusterRequest) SetInstanceCount(v int32) *CreateClusterRequest {
	s.InstanceCount = &v
	return s
}

func (s *CreateClusterRequest) SetInstanceName(v string) *CreateClusterRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateClusterRequest) SetMseVersion(v string) *CreateClusterRequest {
	s.MseVersion = &v
	return s
}

func (s *CreateClusterRequest) SetNetType(v string) *CreateClusterRequest {
	s.NetType = &v
	return s
}

func (s *CreateClusterRequest) SetPrivateSlbSpecification(v string) *CreateClusterRequest {
	s.PrivateSlbSpecification = &v
	return s
}

func (s *CreateClusterRequest) SetPubNetworkFlow(v string) *CreateClusterRequest {
	s.PubNetworkFlow = &v
	return s
}

func (s *CreateClusterRequest) SetPubSlbSpecification(v string) *CreateClusterRequest {
	s.PubSlbSpecification = &v
	return s
}

func (s *CreateClusterRequest) SetRegion(v string) *CreateClusterRequest {
	s.Region = &v
	return s
}

func (s *CreateClusterRequest) SetRequestPars(v string) *CreateClusterRequest {
	s.RequestPars = &v
	return s
}

func (s *CreateClusterRequest) SetResourceGroupId(v string) *CreateClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityGroupType(v string) *CreateClusterRequest {
	s.SecurityGroupType = &v
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

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateClusterRequestTag struct {
	// Tag key.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag value.
	//
	// example:
	//
	// prd
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateClusterRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateClusterRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateClusterRequestTag) SetKey(v string) *CreateClusterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateClusterRequestTag) SetValue(v string) *CreateClusterRequestTag {
	s.Value = &v
	return s
}

func (s *CreateClusterRequestTag) Validate() error {
	return dara.Validate(s)
}
