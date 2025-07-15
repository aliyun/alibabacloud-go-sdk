// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpv4GatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIpv4GatewayResponseBody
	GetRequestId() *string
}

type DeleteIpv4GatewayResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 21133AC0-0636-521B-A400-253818691A56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpv4GatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpv4GatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpv4GatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIpv4GatewayResponseBody) SetRequestId(v string) *DeleteIpv4GatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIpv4GatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
