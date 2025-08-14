// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshTransitRouteTableAggregationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RefreshTransitRouteTableAggregationResponseBody
	GetRequestId() *string
}

type RefreshTransitRouteTableAggregationResponseBody struct {
	// example:
	//
	// 0C2EE7A8-74D4-4081-8236-CEBDE3BBCF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshTransitRouteTableAggregationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshTransitRouteTableAggregationResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshTransitRouteTableAggregationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshTransitRouteTableAggregationResponseBody) SetRequestId(v string) *RefreshTransitRouteTableAggregationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshTransitRouteTableAggregationResponseBody) Validate() error {
	return dara.Validate(s)
}
