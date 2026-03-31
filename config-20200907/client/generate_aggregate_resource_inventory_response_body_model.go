// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAggregateResourceInventoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GenerateAggregateResourceInventoryResponseBody
	GetRequestId() *string
}

type GenerateAggregateResourceInventoryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5111CBA6-6485-57EB-BCDD-85D8BB31E7A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateAggregateResourceInventoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateAggregateResourceInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAggregateResourceInventoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateAggregateResourceInventoryResponseBody) SetRequestId(v string) *GenerateAggregateResourceInventoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateAggregateResourceInventoryResponseBody) Validate() error {
	return dara.Validate(s)
}
