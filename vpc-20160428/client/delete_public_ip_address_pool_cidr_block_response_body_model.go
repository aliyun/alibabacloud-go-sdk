// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePublicIpAddressPoolCidrBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePublicIpAddressPoolCidrBlockResponseBody
	GetRequestId() *string
}

type DeletePublicIpAddressPoolCidrBlockResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePublicIpAddressPoolCidrBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePublicIpAddressPoolCidrBlockResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePublicIpAddressPoolCidrBlockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePublicIpAddressPoolCidrBlockResponseBody) SetRequestId(v string) *DeletePublicIpAddressPoolCidrBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePublicIpAddressPoolCidrBlockResponseBody) Validate() error {
	return dara.Validate(s)
}
