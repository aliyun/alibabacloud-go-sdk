// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImportAddressList(v *ListAddressResp) *ListAddressResponseBody
	GetImportAddressList() *ListAddressResp
}

type ListAddressResponseBody struct {
	// The details of migration addresses.
	ImportAddressList *ListAddressResp `json:"ImportAddressList,omitempty" xml:"ImportAddressList,omitempty"`
}

func (s ListAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddressResponseBody) GetImportAddressList() *ListAddressResp {
	return s.ImportAddressList
}

func (s *ListAddressResponseBody) SetImportAddressList(v *ListAddressResp) *ListAddressResponseBody {
	s.ImportAddressList = v
	return s
}

func (s *ListAddressResponseBody) Validate() error {
	if s.ImportAddressList != nil {
		if err := s.ImportAddressList.Validate(); err != nil {
			return err
		}
	}
	return nil
}
