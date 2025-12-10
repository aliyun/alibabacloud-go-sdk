// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImportAddress(v *CreateAddressInfo) *CreateAddressRequest
	GetImportAddress() *CreateAddressInfo
}

type CreateAddressRequest struct {
	// The details for creating the data address.
	ImportAddress *CreateAddressInfo `json:"ImportAddress,omitempty" xml:"ImportAddress,omitempty"`
}

func (s CreateAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateAddressRequest) GetImportAddress() *CreateAddressInfo {
	return s.ImportAddress
}

func (s *CreateAddressRequest) SetImportAddress(v *CreateAddressInfo) *CreateAddressRequest {
	s.ImportAddress = v
	return s
}

func (s *CreateAddressRequest) Validate() error {
	if s.ImportAddress != nil {
		if err := s.ImportAddress.Validate(); err != nil {
			return err
		}
	}
	return nil
}
