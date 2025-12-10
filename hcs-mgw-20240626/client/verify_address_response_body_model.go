// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetVerifyAddressResponse(v *VerifyAddressResp) *VerifyAddressResponseBody
	GetVerifyAddressResponse() *VerifyAddressResp
}

type VerifyAddressResponseBody struct {
	// The details for verifying the data address.
	VerifyAddressResponse *VerifyAddressResp `json:"VerifyAddressResponse,omitempty" xml:"VerifyAddressResponse,omitempty"`
}

func (s VerifyAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyAddressResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyAddressResponseBody) GetVerifyAddressResponse() *VerifyAddressResp {
	return s.VerifyAddressResponse
}

func (s *VerifyAddressResponseBody) SetVerifyAddressResponse(v *VerifyAddressResp) *VerifyAddressResponseBody {
	s.VerifyAddressResponse = v
	return s
}

func (s *VerifyAddressResponseBody) Validate() error {
	if s.VerifyAddressResponse != nil {
		if err := s.VerifyAddressResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}
