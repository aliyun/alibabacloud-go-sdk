// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregatorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *UpdateAggregatorResponseBody
	GetAggregatorId() *string
	SetRequestId(v string) *UpdateAggregatorResponseBody
	GetRequestId() *string
}

type UpdateAggregatorResponseBody struct {
	// The ID of the account group.
	//
	// example:
	//
	// ca-dacf86d8314e00eb****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8195B664-9565-4685-89AC-8B5F04B44B92
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAggregatorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregatorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAggregatorResponseBody) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *UpdateAggregatorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAggregatorResponseBody) SetAggregatorId(v string) *UpdateAggregatorResponseBody {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregatorResponseBody) SetRequestId(v string) *UpdateAggregatorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAggregatorResponseBody) Validate() error {
	return dara.Validate(s)
}
