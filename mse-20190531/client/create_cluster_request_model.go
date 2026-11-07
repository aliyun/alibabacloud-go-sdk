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
	// The language of the response. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The billing method. Valid values: PREPAY (subscription) and POSTPAY (pay-as-you-go). This parameter is ignored for Serverless Edition.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The engine specifications. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// MSE_SC_2_4_60_c
	ClusterSpecification *string `json:"ClusterSpecification,omitempty" xml:"ClusterSpecification,omitempty"`
	// The cluster type. Valid values: ZooKeeper and Nacos-Ans.
	//
	// This parameter is required.
	//
	// example:
	//
	// Nacos-Ans
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The cluster version. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// NACOS_2_0_0
	ClusterVersion *string `json:"ClusterVersion,omitempty" xml:"ClusterVersion,omitempty"`
	// The network access type. Valid values: `slb` and `single_eni`. In some regions, Developer Edition supports only `single_eni`.
	//
	// example:
	//
	// slb
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated.
	//
	// example:
	//
	// alicloud-disk-ssd
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// Specifies whether to enable public network access by using an elastic IP address (EIP). This parameter takes effect only when ConnectionType is set to `single_eni`.
	//
	// if can be null:
	// false
	EipEnabled *bool `json:"EipEnabled,omitempty" xml:"EipEnabled,omitempty"`
	// The number of nodes in the instance. Valid values: 1 to 9.
	//
	// [Professional Edition]
	//
	// -  `NACOS_2_0_0`: Nacos 2.x.x.
	//
	// - `ZooKeeper_3_8_0`: ZooKeeper 3.8.x.
	//
	// [Developer Edition]
	//
	// -  `NACOS_2_0_0`: Nacos 2.x.
	//
	// - `ZooKeeper_3_8_0`: ZooKeeper 3.8.x.
	//
	// [Serverless Edition]
	//
	// -  `NACOS_2_0_0`: Nacos 2.x.
	//
	// - `ZooKeeper_3_8_0`: ZooKeeper 3.8.x.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The custom name of the instance.
	//
	// example:
	//
	// tanshuyingtest001
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance edition. This parameter is required unless otherwise specified. Valid values:
	//
	// - `mse_pro`: Professional Edition.
	//
	// - `mse_dev`: Developer Edition.
	//
	// - `mse_serverless`: Serverless Edition.
	//
	// example:
	//
	// mse_pro
	MseVersion *string `json:"MseVersion,omitempty" xml:"MseVersion,omitempty"`
	// The network type. Valid values:
	//
	// - `privatenet`: virtual private cloud (VPC).
	//
	// - `pubnet`: Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// privatenet
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated.
	//
	// example:
	//
	// slb.s1.small
	PrivateSlbSpecification *string `json:"PrivateSlbSpecification,omitempty" xml:"PrivateSlbSpecification,omitempty"`
	// The bandwidth of the Internet-facing Server Load Balancer (SLB) instance. This parameter takes effect only when ConnectionType is set to `slb`. A value of 0 indicates that no Internet-facing SLB instance is created. A value greater than 1 indicates the fixed bandwidth of the Internet-facing SLB instance. Unit: Mbit/s. Valid values: 0 to 5000.
	//
	// example:
	//
	// 0
	PubNetworkFlow *string `json:"PubNetworkFlow,omitempty" xml:"PubNetworkFlow,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated.
	//
	// example:
	//
	// slb.s1.small
	PubSlbSpecification *string `json:"PubSlbSpecification,omitempty" xml:"PubSlbSpecification,omitempty"`
	// The region where the cluster is deployed. Valid values include but are not limited to:
	//
	// - `cn-hangzhou`: China (Hangzhou)
	//
	// - `cn-beijing`: China (Beijing)
	//
	// - `cn-shanghai`: China (Shanghai)
	//
	// - `cn-zhangjiakou`: China (Zhangjiakou)
	//
	// - `cn-shenzhen`: China (Shenzhen)
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The extended request parameters in JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	// The resource group ID. For more information, see [View the basic information of a resource group](https://help.aliyun.com/document_detail/457230.html).
	//
	// example:
	//
	// rg-aekzcqmoay3dlyq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the security group to which the instance belongs. This parameter takes effect only when ConnectionType is set to `single_eni`. Valid values:
	//
	// example:
	//
	// enterprise
	SecurityGroupType *string `json:"SecurityGroupType,omitempty" xml:"SecurityGroupType,omitempty"`
	// The tags to be added. You can add up to 20 tags.
	Tag []*CreateClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp17opt4v18sto39k****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
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
	// The tag key.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
