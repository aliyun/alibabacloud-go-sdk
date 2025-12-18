// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateComplianceSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateComplianceSummaryRequest
	GetAggregatorId() *string
}

type GetAggregateComplianceSummaryRequest struct {
	// The ID of the account group.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// ca-a91d626622af0035****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
}

func (s GetAggregateComplianceSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateComplianceSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateComplianceSummaryRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateComplianceSummaryRequest) SetAggregatorId(v string) *GetAggregateComplianceSummaryRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateComplianceSummaryRequest) Validate() error {
	return dara.Validate(s)
}
