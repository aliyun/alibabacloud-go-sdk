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
	// The address type. Valid values:
	//
	// - **Ipv4**: IPv4.
	//
	// - **Ipv6**: IPv6.
	//
	// example:
	//
	// Ipv4
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The burstable protection bandwidth (the Chinese mainland). Unit: Gbit/s.
	//
	// example:
	//
	// 30
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The basic protection bandwidth (the Chinese mainland). Unit: Gbit/s.
	//
	// example:
	//
	// 30
	BaseBandwidth *string `json:"BaseBandwidth,omitempty" xml:"BaseBandwidth,omitempty"`
	// The number of protected domain names.
	//
	// example:
	//
	// 50
	DomainCount *string `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	// The mitigation plan (the Chinese mainland). Valid values:
	//
	// - **coop**: Professional Edition of Anti-DDoS Pro or Anti-DDoS Premium (Chinese mainland).
	//
	// - **advance**: Professional Edition of Anti-DDoS Pro or Anti-DDoS Premium (Chinese mainland).
	//
	// example:
	//
	// coop
	EditionSale *string `json:"EditionSale,omitempty" xml:"EditionSale,omitempty"`
	// The function version. Valid values:
	//
	// - **0**: standard features.
	//
	// - **1**: enhanced features.
	//
	// example:
	//
	// 0
	FunctionVersion *string `json:"FunctionVersion,omitempty" xml:"FunctionVersion,omitempty"`
	// The ID of the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// > You can invoke [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) to query the IDs of all Anti-DDoS Pro or Anti-DDoS Premium instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-6ja1y6p5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of specification change. Valid values:
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
	// The clean bandwidth (outside the Chinese mainland). Unit: Mbit/s.
	//
	// example:
	//
	// 200
	NormalBandwidth *string `json:"NormalBandwidth,omitempty" xml:"NormalBandwidth,omitempty"`
	// The clean queries per second (QPS). Unit: Mbit/s.
	//
	// example:
	//
	// 100
	NormalQps *string `json:"NormalQps,omitempty" xml:"NormalQps,omitempty"`
	// The number of protected ports.
	//
	// example:
	//
	// 50
	PortCount *string `json:"PortCount,omitempty" xml:"PortCount,omitempty"`
	// The mitigation plan (outside the Chinese mainland). Valid values:
	//
	// - **0**: Insurance mitigation plan for Anti-DDoS Pro or Anti-DDoS Premium (outside the Chinese mainland).
	//
	// - **1**: Unlimited mitigation plan for Anti-DDoS Pro or Anti-DDoS Premium (outside the Chinese mainland).
	//
	// - **2**: acceleration line for Anti-DDoS Pro or Anti-DDoS Premium (outside the Chinese mainland).
	//
	// - **3**: secure acceleration line for Anti-DDoS Pro or Anti-DDoS Premium (outside the Chinese mainland).
	//
	// example:
	//
	// 0
	ProductPlan *string `json:"ProductPlan,omitempty" xml:"ProductPlan,omitempty"`
	// The product type.
	//
	// Valid values:
	//
	// - **ddoscoo**: indicates that the China site account is used for specification changes of an Anti-DDoS Pro or Anti-DDoS Premium (Chinese mainland) instance.
	//
	// - **ddoscoo_intl**: indicates that the international site account is used for specification changes of an Anti-DDoS Pro or Anti-DDoS Premium (Chinese mainland) instance.
	//
	// - **ddosDip**: indicates that the China site account or international site account is used for specification changes of an Anti-DDoS Pro or Anti-DDoS Premium (outside the Chinese mainland) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The clean bandwidth (the Chinese mainland). Unit: Mbit/s.
	//
	// example:
	//
	// 300
	ServiceBandwidth *string `json:"ServiceBandwidth,omitempty" xml:"ServiceBandwidth,omitempty"`
	// The line resource of the instance (the Chinese mainland). Valid values:
	//
	// - **coop-line-001**: eight-line BGP line for Anti-DDoS Pro or Anti-DDoS Premium (Chinese mainland).
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
