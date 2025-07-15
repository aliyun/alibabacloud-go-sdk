// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatIpCidrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNatIpCidrResponseBody
	GetRequestId() *string
}

type DeleteNatIpCidrResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7021BEB1-210F-48A9-AB82-BE9A9110BB89
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNatIpCidrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatIpCidrResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNatIpCidrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNatIpCidrResponseBody) SetRequestId(v string) *DeleteNatIpCidrResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNatIpCidrResponseBody) Validate() error {
	return dara.Validate(s)
}
