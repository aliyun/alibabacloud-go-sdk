// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTransitRouteTableAggregationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTransitRouteTableAggregationResponseBody
	GetRequestId() *string
}

type ModifyTransitRouteTableAggregationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0C2EE7A8-74D4-4081-8236-CEBDE3BBCF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTransitRouteTableAggregationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouteTableAggregationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouteTableAggregationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTransitRouteTableAggregationResponseBody) SetRequestId(v string) *ModifyTransitRouteTableAggregationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationResponseBody) Validate() error {
	return dara.Validate(s)
}
