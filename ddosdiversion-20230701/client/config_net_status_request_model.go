// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigNetStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNet(v string) *ConfigNetStatusRequest
	GetNet() *string
	SetRegions(v []*string) *ConfigNetStatusRequest
	GetRegions() []*string
	SetSaleId(v string) *ConfigNetStatusRequest
	GetSaleId() *string
	SetStatus(v string) *ConfigNetStatusRequest
	GetStatus() *string
	SetSubNets(v []*string) *ConfigNetStatusRequest
	GetSubNets() []*string
}

type ConfigNetStatusRequest struct {
	// The CIDR block of the anti-DDoS diversion instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX/22
	Net *string `json:"Net,omitempty" xml:"Net,omitempty"`
	// The regions in which the CIDR block needs to be advertised or withdrawn from advertising. If you leave this parameter empty, the CIDR blocks in all regions are configured.
	//
	// > You can call the [QueryNetList](https://help.aliyun.com/document_detail/2639086.html) operation to obtain the regions of the CIDR blocks.
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the anti-DDoS diversion instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddos_diversion_public_cn-xxxxxxxxxxxxx
	SaleId *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	// The status of the CIDR block. Valid values:
	//
	// - enable: advertises the CIDR block.
	//
	// - disable: withdraws the advertising of the CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subnet CIDR blocks of the CIDR block.
	SubNets []*string `json:"SubNets,omitempty" xml:"SubNets,omitempty" type:"Repeated"`
}

func (s ConfigNetStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigNetStatusRequest) GoString() string {
	return s.String()
}

func (s *ConfigNetStatusRequest) GetNet() *string {
	return s.Net
}

func (s *ConfigNetStatusRequest) GetRegions() []*string {
	return s.Regions
}

func (s *ConfigNetStatusRequest) GetSaleId() *string {
	return s.SaleId
}

func (s *ConfigNetStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *ConfigNetStatusRequest) GetSubNets() []*string {
	return s.SubNets
}

func (s *ConfigNetStatusRequest) SetNet(v string) *ConfigNetStatusRequest {
	s.Net = &v
	return s
}

func (s *ConfigNetStatusRequest) SetRegions(v []*string) *ConfigNetStatusRequest {
	s.Regions = v
	return s
}

func (s *ConfigNetStatusRequest) SetSaleId(v string) *ConfigNetStatusRequest {
	s.SaleId = &v
	return s
}

func (s *ConfigNetStatusRequest) SetStatus(v string) *ConfigNetStatusRequest {
	s.Status = &v
	return s
}

func (s *ConfigNetStatusRequest) SetSubNets(v []*string) *ConfigNetStatusRequest {
	s.SubNets = v
	return s
}

func (s *ConfigNetStatusRequest) Validate() error {
	return dara.Validate(s)
}
