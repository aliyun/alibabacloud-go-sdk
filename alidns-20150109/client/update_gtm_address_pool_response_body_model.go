// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGtmAddressPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateGtmAddressPoolResponseBody
	GetRequestId() *string
}

type UpdateGtmAddressPoolResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGtmAddressPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGtmAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGtmAddressPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGtmAddressPoolResponseBody) SetRequestId(v string) *UpdateGtmAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGtmAddressPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
