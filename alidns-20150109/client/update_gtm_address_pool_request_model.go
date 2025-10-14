// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGtmAddressPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddr(v []*UpdateGtmAddressPoolRequestAddr) *UpdateGtmAddressPoolRequest
	GetAddr() []*UpdateGtmAddressPoolRequestAddr
	SetAddrPoolId(v string) *UpdateGtmAddressPoolRequest
	GetAddrPoolId() *string
	SetLang(v string) *UpdateGtmAddressPoolRequest
	GetLang() *string
	SetMinAvailableAddrNum(v int32) *UpdateGtmAddressPoolRequest
	GetMinAvailableAddrNum() *int32
	SetName(v string) *UpdateGtmAddressPoolRequest
	GetName() *string
	SetType(v string) *UpdateGtmAddressPoolRequest
	GetType() *string
}

type UpdateGtmAddressPoolRequest struct {
	// This parameter is required.
	Addr []*UpdateGtmAddressPoolRequestAddr `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Repeated"`
	// The ID of the address pool that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abc
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The language used by the user.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The minimum number of available addresses in the address pool.
	//
	// example:
	//
	// 2
	MinAvailableAddrNum *int32 `json:"MinAvailableAddrNum,omitempty" xml:"MinAvailableAddrNum,omitempty"`
	// The name of the address pool that you want to modify.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the address pool that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// IP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateGtmAddressPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGtmAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *UpdateGtmAddressPoolRequest) GetAddr() []*UpdateGtmAddressPoolRequestAddr {
	return s.Addr
}

func (s *UpdateGtmAddressPoolRequest) GetAddrPoolId() *string {
	return s.AddrPoolId
}

func (s *UpdateGtmAddressPoolRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateGtmAddressPoolRequest) GetMinAvailableAddrNum() *int32 {
	return s.MinAvailableAddrNum
}

func (s *UpdateGtmAddressPoolRequest) GetName() *string {
	return s.Name
}

func (s *UpdateGtmAddressPoolRequest) GetType() *string {
	return s.Type
}

func (s *UpdateGtmAddressPoolRequest) SetAddr(v []*UpdateGtmAddressPoolRequestAddr) *UpdateGtmAddressPoolRequest {
	s.Addr = v
	return s
}

func (s *UpdateGtmAddressPoolRequest) SetAddrPoolId(v string) *UpdateGtmAddressPoolRequest {
	s.AddrPoolId = &v
	return s
}

func (s *UpdateGtmAddressPoolRequest) SetLang(v string) *UpdateGtmAddressPoolRequest {
	s.Lang = &v
	return s
}

func (s *UpdateGtmAddressPoolRequest) SetMinAvailableAddrNum(v int32) *UpdateGtmAddressPoolRequest {
	s.MinAvailableAddrNum = &v
	return s
}

func (s *UpdateGtmAddressPoolRequest) SetName(v string) *UpdateGtmAddressPoolRequest {
	s.Name = &v
	return s
}

func (s *UpdateGtmAddressPoolRequest) SetType(v string) *UpdateGtmAddressPoolRequest {
	s.Type = &v
	return s
}

func (s *UpdateGtmAddressPoolRequest) Validate() error {
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

type UpdateGtmAddressPoolRequestAddr struct {
	// The weight of the address pool that you want to modify.
	//
	// example:
	//
	// 1
	LbaWeight *int32 `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	// The mode of the address pool that you want to modify.
	//
	// 	- **SMART**: Intelligent return
	//
	// 	- **ONLINE**: Always online
	//
	// 	- **OFFLINE**: Always offline
	//
	// example:
	//
	// SMART
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The addresses in the address pool.
	//
	// example:
	//
	// 1.1.1.1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateGtmAddressPoolRequestAddr) String() string {
	return dara.Prettify(s)
}

func (s UpdateGtmAddressPoolRequestAddr) GoString() string {
	return s.String()
}

func (s *UpdateGtmAddressPoolRequestAddr) GetLbaWeight() *int32 {
	return s.LbaWeight
}

func (s *UpdateGtmAddressPoolRequestAddr) GetMode() *string {
	return s.Mode
}

func (s *UpdateGtmAddressPoolRequestAddr) GetValue() *string {
	return s.Value
}

func (s *UpdateGtmAddressPoolRequestAddr) SetLbaWeight(v int32) *UpdateGtmAddressPoolRequestAddr {
	s.LbaWeight = &v
	return s
}

func (s *UpdateGtmAddressPoolRequestAddr) SetMode(v string) *UpdateGtmAddressPoolRequestAddr {
	s.Mode = &v
	return s
}

func (s *UpdateGtmAddressPoolRequestAddr) SetValue(v string) *UpdateGtmAddressPoolRequestAddr {
	s.Value = &v
	return s
}

func (s *UpdateGtmAddressPoolRequestAddr) Validate() error {
	return dara.Validate(s)
}
