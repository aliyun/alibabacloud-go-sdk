// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressGroup interface {
	dara.Model
	String() string
	GoString() string
	SetAddresses(v []*string) *AddressGroup
	GetAddresses() []*string
	SetPorts(v []*AddressGroupPorts) *AddressGroup
	GetPorts() []*AddressGroupPorts
}

type AddressGroup struct {
	Addresses []*string            `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	Ports     []*AddressGroupPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
}

func (s AddressGroup) String() string {
	return dara.Prettify(s)
}

func (s AddressGroup) GoString() string {
	return s.String()
}

func (s *AddressGroup) GetAddresses() []*string {
	return s.Addresses
}

func (s *AddressGroup) GetPorts() []*AddressGroupPorts {
	return s.Ports
}

func (s *AddressGroup) SetAddresses(v []*string) *AddressGroup {
	s.Addresses = v
	return s
}

func (s *AddressGroup) SetPorts(v []*AddressGroupPorts) *AddressGroup {
	s.Ports = v
	return s
}

func (s *AddressGroup) Validate() error {
	if s.Ports != nil {
		for _, item := range s.Ports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddressGroupPorts struct {
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	End   *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s AddressGroupPorts) String() string {
	return dara.Prettify(s)
}

func (s AddressGroupPorts) GoString() string {
	return s.String()
}

func (s *AddressGroupPorts) GetBegin() *int32 {
	return s.Begin
}

func (s *AddressGroupPorts) GetEnd() *int32 {
	return s.End
}

func (s *AddressGroupPorts) SetBegin(v int32) *AddressGroupPorts {
	s.Begin = &v
	return s
}

func (s *AddressGroupPorts) SetEnd(v int32) *AddressGroupPorts {
	s.End = &v
	return s
}

func (s *AddressGroupPorts) Validate() error {
	return dara.Validate(s)
}
