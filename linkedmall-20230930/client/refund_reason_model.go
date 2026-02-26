// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundReason interface {
	dara.Model
	String() string
	GoString() string
	SetProofRequired(v bool) *RefundReason
	GetProofRequired() *bool
	SetReasonTextId(v string) *RefundReason
	GetReasonTextId() *string
	SetReasonTips(v string) *RefundReason
	GetReasonTips() *string
	SetRefundDescRequired(v bool) *RefundReason
	GetRefundDescRequired() *bool
}

type RefundReason struct {
	// example:
	//
	// true
	ProofRequired *bool `json:"proofRequired,omitempty" xml:"proofRequired,omitempty"`
	// example:
	//
	// 47683
	ReasonTextId *string `json:"reasonTextId,omitempty" xml:"reasonTextId,omitempty"`
	ReasonTips   *string `json:"reasonTips,omitempty" xml:"reasonTips,omitempty"`
	// example:
	//
	// true
	RefundDescRequired *bool `json:"refundDescRequired,omitempty" xml:"refundDescRequired,omitempty"`
}

func (s RefundReason) String() string {
	return dara.Prettify(s)
}

func (s RefundReason) GoString() string {
	return s.String()
}

func (s *RefundReason) GetProofRequired() *bool {
	return s.ProofRequired
}

func (s *RefundReason) GetReasonTextId() *string {
	return s.ReasonTextId
}

func (s *RefundReason) GetReasonTips() *string {
	return s.ReasonTips
}

func (s *RefundReason) GetRefundDescRequired() *bool {
	return s.RefundDescRequired
}

func (s *RefundReason) SetProofRequired(v bool) *RefundReason {
	s.ProofRequired = &v
	return s
}

func (s *RefundReason) SetReasonTextId(v string) *RefundReason {
	s.ReasonTextId = &v
	return s
}

func (s *RefundReason) SetReasonTips(v string) *RefundReason {
	s.ReasonTips = &v
	return s
}

func (s *RefundReason) SetRefundDescRequired(v bool) *RefundReason {
	s.RefundDescRequired = &v
	return s
}

func (s *RefundReason) Validate() error {
	return dara.Validate(s)
}
