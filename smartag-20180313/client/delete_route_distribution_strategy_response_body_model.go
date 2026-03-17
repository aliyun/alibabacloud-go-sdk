// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteDistributionStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRouteDistributionStrategyResponseBody
	GetRequestId() *string
}

type DeleteRouteDistributionStrategyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// BF62139B-D64A-4C95-A55F-6A2335C4417D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRouteDistributionStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteDistributionStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRouteDistributionStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRouteDistributionStrategyResponseBody) SetRequestId(v string) *DeleteRouteDistributionStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRouteDistributionStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
