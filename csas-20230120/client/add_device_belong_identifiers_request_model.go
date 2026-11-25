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
	IdentifierType  *string `json:"IdentifierType,omitempty" xml:"IdentifierType,omitempty"`
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
