// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVendorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateVendorResponseBody
	GetRequestId() *string
}

type UpdateVendorResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVendorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVendorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVendorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVendorResponseBody) SetRequestId(v string) *UpdateVendorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVendorResponseBody) Validate() error {
	return dara.Validate(s)
}
