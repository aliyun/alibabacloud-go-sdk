// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIstioGatewayDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIstioGatewayDomainsResponseBody
	GetRequestId() *string
}

type DeleteIstioGatewayDomainsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIstioGatewayDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIstioGatewayDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIstioGatewayDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIstioGatewayDomainsResponseBody) SetRequestId(v string) *DeleteIstioGatewayDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIstioGatewayDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}
