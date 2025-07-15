// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePublicIpAddressPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePublicIpAddressPoolResponseBody
	GetRequestId() *string
}

type DeletePublicIpAddressPoolResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePublicIpAddressPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePublicIpAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePublicIpAddressPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePublicIpAddressPoolResponseBody) SetRequestId(v string) *DeletePublicIpAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePublicIpAddressPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
