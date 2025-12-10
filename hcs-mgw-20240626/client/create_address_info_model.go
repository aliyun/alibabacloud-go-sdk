// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAddressInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAddressDetail(v *AddressDetail) *CreateAddressInfo
	GetAddressDetail() *AddressDetail
	SetName(v string) *CreateAddressInfo
	GetName() *string
	SetTags(v string) *CreateAddressInfo
	GetTags() *string
}

type CreateAddressInfo struct {
	// This parameter is required.
	AddressDetail *AddressDetail `json:"AddressDetail,omitempty" xml:"AddressDetail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// <your-address-name>
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// K1:V1,K2:V2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateAddressInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateAddressInfo) GoString() string {
	return s.String()
}

func (s *CreateAddressInfo) GetAddressDetail() *AddressDetail {
	return s.AddressDetail
}

func (s *CreateAddressInfo) GetName() *string {
	return s.Name
}

func (s *CreateAddressInfo) GetTags() *string {
	return s.Tags
}

func (s *CreateAddressInfo) SetAddressDetail(v *AddressDetail) *CreateAddressInfo {
	s.AddressDetail = v
	return s
}

func (s *CreateAddressInfo) SetName(v string) *CreateAddressInfo {
	s.Name = &v
	return s
}

func (s *CreateAddressInfo) SetTags(v string) *CreateAddressInfo {
	s.Tags = &v
	return s
}

func (s *CreateAddressInfo) Validate() error {
	if s.AddressDetail != nil {
		if err := s.AddressDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
