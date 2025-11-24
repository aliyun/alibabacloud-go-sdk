// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIstioGatewayDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateIstioGatewayDomainsResponseBody
	GetRequestId() *string
}

type CreateIstioGatewayDomainsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIstioGatewayDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIstioGatewayDomainsResponseBody) SetRequestId(v string) *CreateIstioGatewayDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIstioGatewayDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}
