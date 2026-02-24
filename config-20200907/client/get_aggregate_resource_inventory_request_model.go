// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceInventoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateResourceInventoryRequest
	GetAggregatorId() *string
}

type GetAggregateResourceInventoryRequest struct {
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
}

func (s GetAggregateResourceInventoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceInventoryRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceInventoryRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateResourceInventoryRequest) SetAggregatorId(v string) *GetAggregateResourceInventoryRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateResourceInventoryRequest) Validate() error {
	return dara.Validate(s)
}
