// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptFulfillmentDecisionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDecisionConclusion(v string) *AcceptFulfillmentDecisionRequest
	GetDecisionConclusion() *string
	SetDecisionType(v string) *AcceptFulfillmentDecisionRequest
	GetDecisionType() *string
	SetOrderId(v string) *AcceptFulfillmentDecisionRequest
	GetOrderId() *string
}

type AcceptFulfillmentDecisionRequest struct {
	DecisionConclusion *string `json:"DecisionConclusion,omitempty" xml:"DecisionConclusion,omitempty"`
	DecisionType       *string `json:"DecisionType,omitempty" xml:"DecisionType,omitempty"`
	OrderId            *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s AcceptFulfillmentDecisionRequest) String() string {
	return dara.Prettify(s)
}

func (s AcceptFulfillmentDecisionRequest) GoString() string {
	return s.String()
}

func (s *AcceptFulfillmentDecisionRequest) GetDecisionConclusion() *string {
	return s.DecisionConclusion
}

func (s *AcceptFulfillmentDecisionRequest) GetDecisionType() *string {
	return s.DecisionType
}

func (s *AcceptFulfillmentDecisionRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *AcceptFulfillmentDecisionRequest) SetDecisionConclusion(v string) *AcceptFulfillmentDecisionRequest {
	s.DecisionConclusion = &v
	return s
}

func (s *AcceptFulfillmentDecisionRequest) SetDecisionType(v string) *AcceptFulfillmentDecisionRequest {
	s.DecisionType = &v
	return s
}

func (s *AcceptFulfillmentDecisionRequest) SetOrderId(v string) *AcceptFulfillmentDecisionRequest {
	s.OrderId = &v
	return s
}

func (s *AcceptFulfillmentDecisionRequest) Validate() error {
	return dara.Validate(s)
}
