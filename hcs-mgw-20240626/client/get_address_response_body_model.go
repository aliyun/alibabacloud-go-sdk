// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImportAddress(v *GetAddressResp) *GetAddressResponseBody
	GetImportAddress() *GetAddressResp
}

type GetAddressResponseBody struct {
	// The details for obtaining the data address.
	//
	// Valid values:
	//
	// 	- 1:1
	ImportAddress *GetAddressResp `json:"ImportAddress,omitempty" xml:"ImportAddress,omitempty"`
}

func (s GetAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAddressResponseBody) GoString() string {
	return s.String()
}

func (s *GetAddressResponseBody) GetImportAddress() *GetAddressResp {
	return s.ImportAddress
}

func (s *GetAddressResponseBody) SetImportAddress(v *GetAddressResp) *GetAddressResponseBody {
	s.ImportAddress = v
	return s
}

func (s *GetAddressResponseBody) Validate() error {
	if s.ImportAddress != nil {
		if err := s.ImportAddress.Validate(); err != nil {
			return err
		}
	}
	return nil
}
