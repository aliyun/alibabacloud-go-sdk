// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsGtmAddressPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDnsGtmAddressPoolResponseBody
	GetRequestId() *string
}

type UpdateDnsGtmAddressPoolResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDnsGtmAddressPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsGtmAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAddressPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDnsGtmAddressPoolResponseBody) SetRequestId(v string) *UpdateDnsGtmAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
