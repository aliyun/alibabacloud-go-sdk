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
	// The details of the data address.
	//
	// This parameter is required.
	AddressDetail *AddressDetail `json:"AddressDetail,omitempty" xml:"AddressDetail,omitempty"`
	// The name of the data address.\\
	//
	// The name can contain lowercase letters, digits, hyphens (-), and underscores (_). The name must be 3 to 63 characters in length. The name is case-sensitive and encoded in UTF-8. The name cannot start with a hyphen (-) or an underscore (_). You must specify a name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tags in the key:value format.\\
	//
	// The value can contain letters, digits, hyphens (-), underscores (_), and commas (,). The value can be up to 1,024 characters in length.
	//
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
