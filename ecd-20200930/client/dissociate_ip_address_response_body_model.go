// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DissociateIpAddressResponseBody
	GetRequestId() *string
}

type DissociateIpAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DissociateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateIpAddressResponseBody) SetRequestId(v string) *DissociateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateIpAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
