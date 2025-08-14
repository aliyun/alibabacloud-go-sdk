// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouteTableAggregationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTransitRouteTableAggregationResponseBody
	GetRequestId() *string
}

type DeleteTransitRouteTableAggregationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0C2EE7A8-74D4-4081-8236-CEBDE3BBCF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouteTableAggregationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouteTableAggregationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouteTableAggregationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTransitRouteTableAggregationResponseBody) SetRequestId(v string) *DeleteTransitRouteTableAggregationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTransitRouteTableAggregationResponseBody) Validate() error {
	return dara.Validate(s)
}
