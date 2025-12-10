// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImportAddress(v *UpdateAddressInfo) *UpdateAddressRequest
	GetImportAddress() *UpdateAddressInfo
}

type UpdateAddressRequest struct {
	// The details for updating the data address.
	ImportAddress *UpdateAddressInfo `json:"ImportAddress,omitempty" xml:"ImportAddress,omitempty"`
}

func (s UpdateAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAddressRequest) GoString() string {
	return s.String()
}

func (s *UpdateAddressRequest) GetImportAddress() *UpdateAddressInfo {
	return s.ImportAddress
}

func (s *UpdateAddressRequest) SetImportAddress(v *UpdateAddressInfo) *UpdateAddressRequest {
	s.ImportAddress = v
	return s
}

func (s *UpdateAddressRequest) Validate() error {
	if s.ImportAddress != nil {
		if err := s.ImportAddress.Validate(); err != nil {
			return err
		}
	}
	return nil
}
