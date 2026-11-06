// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVendorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVendorResponseBody
	GetRequestId() *string
	SetVendorId(v string) *CreateVendorResponseBody
	GetVendorId() *string
}

type CreateVendorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VendorId  *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s CreateVendorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVendorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVendorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVendorResponseBody) GetVendorId() *string {
	return s.VendorId
}

func (s *CreateVendorResponseBody) SetRequestId(v string) *CreateVendorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVendorResponseBody) SetVendorId(v string) *CreateVendorResponseBody {
	s.VendorId = &v
	return s
}

func (s *CreateVendorResponseBody) Validate() error {
	return dara.Validate(s)
}
