// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsGtmAddressPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddr(v []*UpdateDnsGtmAddressPoolRequestAddr) *UpdateDnsGtmAddressPoolRequest
	GetAddr() []*UpdateDnsGtmAddressPoolRequestAddr
	SetAddrPoolId(v string) *UpdateDnsGtmAddressPoolRequest
	GetAddrPoolId() *string
	SetLang(v string) *UpdateDnsGtmAddressPoolRequest
	GetLang() *string
	SetLbaStrategy(v string) *UpdateDnsGtmAddressPoolRequest
	GetLbaStrategy() *string
	SetName(v string) *UpdateDnsGtmAddressPoolRequest
	GetName() *string
}

type UpdateDnsGtmAddressPoolRequest struct {
	// A list of addresses in the address pool.
	//
	// This parameter is required.
	Addr []*UpdateDnsGtmAddressPoolRequestAddr `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Repeated"`
	// The ID of the address pool. For more information, see [DescribeDnsGtmInstanceAddressPools](https://www.alibabacloud.com/help/en/dns/api-alidns-2015-01-09-describednsgtminstanceaddresspools).
	//
	// This parameter is required.
	//
	// example:
	//
	// testp******
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The language of the response. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The load balancing policy. Valid values:
	//
	// - ALL_RR: Returns all addresses.
	//
	// - RATIO: Returns addresses by weight.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL_RR
	LbaStrategy *string `json:"LbaStrategy,omitempty" xml:"LbaStrategy,omitempty"`
	// The name of the address pool.
	//
	// example:
	//
	// testpoolname
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateDnsGtmAddressPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsGtmAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAddressPoolRequest) GetAddr() []*UpdateDnsGtmAddressPoolRequestAddr {
	return s.Addr
}

func (s *UpdateDnsGtmAddressPoolRequest) GetAddrPoolId() *string {
	return s.AddrPoolId
}

func (s *UpdateDnsGtmAddressPoolRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDnsGtmAddressPoolRequest) GetLbaStrategy() *string {
	return s.LbaStrategy
}

func (s *UpdateDnsGtmAddressPoolRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDnsGtmAddressPoolRequest) SetAddr(v []*UpdateDnsGtmAddressPoolRequestAddr) *UpdateDnsGtmAddressPoolRequest {
	s.Addr = v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequest) SetAddrPoolId(v string) *UpdateDnsGtmAddressPoolRequest {
	s.AddrPoolId = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequest) SetLang(v string) *UpdateDnsGtmAddressPoolRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequest) SetLbaStrategy(v string) *UpdateDnsGtmAddressPoolRequest {
	s.LbaStrategy = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequest) SetName(v string) *UpdateDnsGtmAddressPoolRequest {
	s.Name = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequest) Validate() error {
	if s.Addr != nil {
		for _, item := range s.Addr {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDnsGtmAddressPoolRequestAddr struct {
	// The address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.1.XX.XX
	Addr *string `json:"Addr,omitempty" xml:"Addr,omitempty"`
	// The source region of the address. This parameter is a JSON string.
	//
	// - LineCode: The line code of the source region. This parameter is deprecated. Use lineCodes instead.
	//
	// - lineCodes: A list of line codes for the source regions.
	//
	// - lineCodeRectifyType: The type of line code rectification. Default value: AUTO. Valid values:
	//
	//   - NO_NEED: No rectification is required.
	//
	//   - RECTIFIED: The line code is rectified.
	//
	//   - AUTO: The line code is automatically rectified.
	//
	// example:
	//
	// Linecode: default
	//
	// lineCodes： ["default"]
	//
	// lineCodeRectifyType: AUTO
	AttributeInfo *string `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	// The weight.
	//
	// example:
	//
	// 1
	LbaWeight *int32 `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	// The mode. Valid values:
	//
	// - SMART: Smart return
	//
	// - ONLINE: Always online
	//
	// - OFFLINE: Always offline
	//
	// This parameter is required.
	//
	// example:
	//
	// SMART
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateDnsGtmAddressPoolRequestAddr) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsGtmAddressPoolRequestAddr) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAddressPoolRequestAddr) GetAddr() *string {
	return s.Addr
}

func (s *UpdateDnsGtmAddressPoolRequestAddr) GetAttributeInfo() *string {
	return s.AttributeInfo
}

func (s *UpdateDnsGtmAddressPoolRequestAddr) GetLbaWeight() *int32 {
	return s.LbaWeight
}

func (s *UpdateDnsGtmAddressPoolRequestAddr) GetMode() *string {
	return s.Mode
}

func (s *UpdateDnsGtmAddressPoolRequestAddr) GetRemark() *string {
	return s.Remark
}

func (s *UpdateDnsGtmAddressPoolRequestAddr) SetAddr(v string) *UpdateDnsGtmAddressPoolRequestAddr {
	s.Addr = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequestAddr) SetAttributeInfo(v string) *UpdateDnsGtmAddressPoolRequestAddr {
	s.AttributeInfo = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequestAddr) SetLbaWeight(v int32) *UpdateDnsGtmAddressPoolRequestAddr {
	s.LbaWeight = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequestAddr) SetMode(v string) *UpdateDnsGtmAddressPoolRequestAddr {
	s.Mode = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequestAddr) SetRemark(v string) *UpdateDnsGtmAddressPoolRequestAddr {
	s.Remark = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolRequestAddr) Validate() error {
	return dara.Validate(s)
}
