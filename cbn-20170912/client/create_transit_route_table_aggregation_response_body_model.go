// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouteTableAggregationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTransitRouteTableAggregationResponseBody
	GetRequestId() *string
}

type CreateTransitRouteTableAggregationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0C2EE7A8-74D4-4081-8236-CEBDE3BBCF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTransitRouteTableAggregationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouteTableAggregationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouteTableAggregationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTransitRouteTableAggregationResponseBody) SetRequestId(v string) *CreateTransitRouteTableAggregationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouteTableAggregationResponseBody) Validate() error {
	return dara.Validate(s)
}
