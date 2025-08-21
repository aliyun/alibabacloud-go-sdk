// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressType(v string) *ModifyInstanceRequest
	GetAddressType() *string
	SetBandwidth(v string) *ModifyInstanceRequest
	GetBandwidth() *string
	SetBaseBandwidth(v string) *ModifyInstanceRequest
	GetBaseBandwidth() *string
	SetDomainCount(v string) *ModifyInstanceRequest
	GetDomainCount() *string
	SetEditionSale(v string) *ModifyInstanceRequest
	GetEditionSale() *string
	SetFunctionVersion(v string) *ModifyInstanceRequest
	GetFunctionVersion() *string
	SetInstanceId(v string) *ModifyInstanceRequest
	GetInstanceId() *string
	SetModifyType(v string) *ModifyInstanceRequest
	GetModifyType() *string
	SetNormalBandwidth(v string) *ModifyInstanceRequest
	GetNormalBandwidth() *string
	SetNormalQps(v string) *ModifyInstanceRequest
	GetNormalQps() *string
	SetPortCount(v string) *ModifyInstanceRequest
	GetPortCount() *string
	SetProductPlan(v string) *ModifyInstanceRequest
	GetProductPlan() *string
	SetProductType(v string) *ModifyInstanceRequest
	GetProductType() *string
	SetServiceBandwidth(v string) *ModifyInstanceRequest
	GetServiceBandwidth() *string
	SetServicePartner(v string) *ModifyInstanceRequest
	GetServicePartner() *string
}

type ModifyInstanceRequest struct {
	// Address type. Values:
	//
	// - **Ipv4**: IPv4.
	//
	// - **Ipv6**: IPv6.
	//
	// example:
	//
	// Ipv4
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// Elastic protection bandwidth (Mainland China). Unit: Gbps.
	//
	// example:
	//
	// 30
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// Guaranteed protection bandwidth (Mainland China). Unit: Gbps.
	//
	// example:
	//
	// 30
	BaseBandwidth *string `json:"BaseBandwidth,omitempty" xml:"BaseBandwidth,omitempty"`
	// Number of protected domains.
	//
	// example:
	//
	// 50
	DomainCount *string `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	// Protection package (Mainland China). Values:
	//
	// - **coop**: Indicates the DDoS High Defense (Mainland China) Professional Edition.
	//
	// - **advance**: Indicates the DDoS High Defense (Mainland China) Professional Edition.
	//
	// example:
	//
	// coop
	EditionSale *string `json:"EditionSale,omitempty" xml:"EditionSale,omitempty"`
	// Function version, with values:
	//
	// - **0**: Standard function.
	//
	// - **1**: Enhanced function.
	//
	// example:
	//
	// 0
	FunctionVersion *string `json:"FunctionVersion,omitempty" xml:"FunctionVersion,omitempty"`
	// The ID of the DDoS High Defense instance.
	//
	// > You can call [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) to query the ID information of all DDoS High Defense instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-6ja1y6p5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Adjustment type, with values
	//
	// - UPGRADE: Upgrade.
	//
	// - DOWNGRADE: Downgrade.
	//
	// This parameter is required.
	//
	// example:
	//
	// Upgrade
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// Business bandwidth (outside Mainland China). Unit: Mbps.
	//
	// example:
	//
	// 200
	NormalBandwidth *string `json:"NormalBandwidth,omitempty" xml:"NormalBandwidth,omitempty"`
	// Business QPS. Unit: Mbps.
	//
	// example:
	//
	// 100
	NormalQps *string `json:"NormalQps,omitempty" xml:"NormalQps,omitempty"`
	// Number of protected ports.
	//
	// example:
	//
	// 50
	PortCount *string `json:"PortCount,omitempty" xml:"PortCount,omitempty"`
	// Protection package (outside Mainland China). Values:
	//
	// - **0**: Indicates the DDoS High Defense (outside Mainland China) Insurance Edition.
	//
	// - **1**: Indicates the DDoS High Defense (outside Mainland China) Worry-Free Edition.
	//
	// - **2**: Indicates the DDoS High Defense (outside Mainland China) Acceleration Line.
	//
	// - **3**: Indicates the DDoS High Defense (outside Mainland China) Secure Acceleration Line.
	//
	// example:
	//
	// 0
	ProductPlan *string `json:"ProductPlan,omitempty" xml:"ProductPlan,omitempty"`
	// Product type.
	//
	// Values:
	//
	// - **ddoscoo**: Indicates that the DDoS High Defense (Mainland China) instance is being adjusted for a China site account.
	//
	// - **ddoscoo_intl**: Indicates that the DDoS High Defense (Mainland China) instance is being adjusted for an international site account.
	//
	// - **ddosDip**: Indicates that the DDoS High Defense (outside Mainland China) instance is being adjusted for either a China or international site account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// Business bandwidth (Mainland China). Unit: Mbps.
	//
	// example:
	//
	// 300
	ServiceBandwidth *string `json:"ServiceBandwidth,omitempty" xml:"ServiceBandwidth,omitempty"`
	// Line resources of the instance (Mainland China). Values:
	//
	// - **coop-line-001**: Indicates the DDoS High Defense (Mainland China) 8-line BGP line.
	//
	// example:
	//
	// coop-line-001
	ServicePartner *string `json:"ServicePartner,omitempty" xml:"ServicePartner,omitempty"`
}

func (s ModifyInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *ModifyInstanceRequest) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *ModifyInstanceRequest) GetBaseBandwidth() *string {
	return s.BaseBandwidth
}

func (s *ModifyInstanceRequest) GetDomainCount() *string {
	return s.DomainCount
}

func (s *ModifyInstanceRequest) GetEditionSale() *string {
	return s.EditionSale
}

func (s *ModifyInstanceRequest) GetFunctionVersion() *string {
	return s.FunctionVersion
}

func (s *ModifyInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *ModifyInstanceRequest) GetNormalBandwidth() *string {
	return s.NormalBandwidth
}

func (s *ModifyInstanceRequest) GetNormalQps() *string {
	return s.NormalQps
}

func (s *ModifyInstanceRequest) GetPortCount() *string {
	return s.PortCount
}

func (s *ModifyInstanceRequest) GetProductPlan() *string {
	return s.ProductPlan
}

func (s *ModifyInstanceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyInstanceRequest) GetServiceBandwidth() *string {
	return s.ServiceBandwidth
}

func (s *ModifyInstanceRequest) GetServicePartner() *string {
	return s.ServicePartner
}

func (s *ModifyInstanceRequest) SetAddressType(v string) *ModifyInstanceRequest {
	s.AddressType = &v
	return s
}

func (s *ModifyInstanceRequest) SetBandwidth(v string) *ModifyInstanceRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyInstanceRequest) SetBaseBandwidth(v string) *ModifyInstanceRequest {
	s.BaseBandwidth = &v
	return s
}

func (s *ModifyInstanceRequest) SetDomainCount(v string) *ModifyInstanceRequest {
	s.DomainCount = &v
	return s
}

func (s *ModifyInstanceRequest) SetEditionSale(v string) *ModifyInstanceRequest {
	s.EditionSale = &v
	return s
}

func (s *ModifyInstanceRequest) SetFunctionVersion(v string) *ModifyInstanceRequest {
	s.FunctionVersion = &v
	return s
}

func (s *ModifyInstanceRequest) SetInstanceId(v string) *ModifyInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceRequest) SetModifyType(v string) *ModifyInstanceRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyInstanceRequest) SetNormalBandwidth(v string) *ModifyInstanceRequest {
	s.NormalBandwidth = &v
	return s
}

func (s *ModifyInstanceRequest) SetNormalQps(v string) *ModifyInstanceRequest {
	s.NormalQps = &v
	return s
}

func (s *ModifyInstanceRequest) SetPortCount(v string) *ModifyInstanceRequest {
	s.PortCount = &v
	return s
}

func (s *ModifyInstanceRequest) SetProductPlan(v string) *ModifyInstanceRequest {
	s.ProductPlan = &v
	return s
}

func (s *ModifyInstanceRequest) SetProductType(v string) *ModifyInstanceRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyInstanceRequest) SetServiceBandwidth(v string) *ModifyInstanceRequest {
	s.ServiceBandwidth = &v
	return s
}

func (s *ModifyInstanceRequest) SetServicePartner(v string) *ModifyInstanceRequest {
	s.ServicePartner = &v
	return s
}

func (s *ModifyInstanceRequest) Validate() error {
	return dara.Validate(s)
}
