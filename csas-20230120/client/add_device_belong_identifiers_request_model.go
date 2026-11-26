// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDeviceBelongIdentifiersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifiers(v []*AddDeviceBelongIdentifiersRequestIdentifiers) *AddDeviceBelongIdentifiersRequest
	GetIdentifiers() []*AddDeviceBelongIdentifiersRequestIdentifiers
}

type AddDeviceBelongIdentifiersRequest struct {
	// The list of device ownership identifiers to add. Each batch supports 1 to 1,000 items. Each item must include IdentifierType and IdentifierValue.
	//
	// This parameter is required.
	Identifiers []*AddDeviceBelongIdentifiersRequestIdentifiers `json:"Identifiers,omitempty" xml:"Identifiers,omitempty" type:"Repeated"`
}

func (s AddDeviceBelongIdentifiersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDeviceBelongIdentifiersRequest) GoString() string {
	return s.String()
}

func (s *AddDeviceBelongIdentifiersRequest) GetIdentifiers() []*AddDeviceBelongIdentifiersRequestIdentifiers {
	return s.Identifiers
}

func (s *AddDeviceBelongIdentifiersRequest) SetIdentifiers(v []*AddDeviceBelongIdentifiersRequestIdentifiers) *AddDeviceBelongIdentifiersRequest {
	s.Identifiers = v
	return s
}

func (s *AddDeviceBelongIdentifiersRequest) Validate() error {
	if s.Identifiers != nil {
		for _, item := range s.Identifiers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddDeviceBelongIdentifiersRequestIdentifiers struct {
	// The identifier type. MAC indicates a network interface controller (NIC) MAC address, SystemSN indicates a system serial number, and BIOSSN indicates a BIOS serial number.
	//
	// example:
	//
	// MAC
	IdentifierType *string `json:"IdentifierType,omitempty" xml:"IdentifierType,omitempty"`
	// The identifier value. A MAC address must be a 6-byte address. A system serial number accepts 1 to 64 characters, and a BIOS serial number accepts 1 to 128 characters. MAC addresses are converted to the lowercase colon-separated format when stored, while system serial numbers and BIOS serial numbers retain their original values.
	//
	// example:
	//
	// 02:00:00:00:00:01
	IdentifierValue *string `json:"IdentifierValue,omitempty" xml:"IdentifierValue,omitempty"`
}

func (s AddDeviceBelongIdentifiersRequestIdentifiers) String() string {
	return dara.Prettify(s)
}

func (s AddDeviceBelongIdentifiersRequestIdentifiers) GoString() string {
	return s.String()
}

func (s *AddDeviceBelongIdentifiersRequestIdentifiers) GetIdentifierType() *string {
	return s.IdentifierType
}

func (s *AddDeviceBelongIdentifiersRequestIdentifiers) GetIdentifierValue() *string {
	return s.IdentifierValue
}

func (s *AddDeviceBelongIdentifiersRequestIdentifiers) SetIdentifierType(v string) *AddDeviceBelongIdentifiersRequestIdentifiers {
	s.IdentifierType = &v
	return s
}

func (s *AddDeviceBelongIdentifiersRequestIdentifiers) SetIdentifierValue(v string) *AddDeviceBelongIdentifiersRequestIdentifiers {
	s.IdentifierValue = &v
	return s
}

func (s *AddDeviceBelongIdentifiersRequestIdentifiers) Validate() error {
	return dara.Validate(s)
}
