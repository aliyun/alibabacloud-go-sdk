// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDnsGtmAddressPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDnsGtmAddressPoolResponseBody
	GetRequestId() *string
}

type DeleteDnsGtmAddressPoolResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDnsGtmAddressPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDnsGtmAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDnsGtmAddressPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDnsGtmAddressPoolResponseBody) SetRequestId(v string) *DeleteDnsGtmAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDnsGtmAddressPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
