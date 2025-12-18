// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregatorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *CreateAggregatorResponseBody
	GetAggregatorId() *string
	SetRequestId(v string) *CreateAggregatorResponseBody
	GetRequestId() *string
}

type CreateAggregatorResponseBody struct {
	// The account group ID.
	//
	// example:
	//
	// ca-dacf86d8314e00eb****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8195B664-9565-4685-89AC-8B5F04B44B92
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAggregatorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregatorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAggregatorResponseBody) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *CreateAggregatorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAggregatorResponseBody) SetAggregatorId(v string) *CreateAggregatorResponseBody {
	s.AggregatorId = &v
	return s
}

func (s *CreateAggregatorResponseBody) SetRequestId(v string) *CreateAggregatorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAggregatorResponseBody) Validate() error {
	return dara.Validate(s)
}
