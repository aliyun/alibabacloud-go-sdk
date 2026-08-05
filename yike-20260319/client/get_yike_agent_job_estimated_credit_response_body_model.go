// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeAgentJobEstimatedCreditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEstimatedCreditCost(v float64) *GetYikeAgentJobEstimatedCreditResponseBody
	GetEstimatedCreditCost() *float64
	SetRequestId(v string) *GetYikeAgentJobEstimatedCreditResponseBody
	GetRequestId() *string
}

type GetYikeAgentJobEstimatedCreditResponseBody struct {
	// The estimated credits to be deducted.
	//
	// example:
	//
	// 20.1
	EstimatedCreditCost *float64 `json:"EstimatedCreditCost,omitempty" xml:"EstimatedCreditCost,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetYikeAgentJobEstimatedCreditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAgentJobEstimatedCreditResponseBody) GoString() string {
	return s.String()
}

func (s *GetYikeAgentJobEstimatedCreditResponseBody) GetEstimatedCreditCost() *float64 {
	return s.EstimatedCreditCost
}

func (s *GetYikeAgentJobEstimatedCreditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetYikeAgentJobEstimatedCreditResponseBody) SetEstimatedCreditCost(v float64) *GetYikeAgentJobEstimatedCreditResponseBody {
	s.EstimatedCreditCost = &v
	return s
}

func (s *GetYikeAgentJobEstimatedCreditResponseBody) SetRequestId(v string) *GetYikeAgentJobEstimatedCreditResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetYikeAgentJobEstimatedCreditResponseBody) Validate() error {
	return dara.Validate(s)
}
