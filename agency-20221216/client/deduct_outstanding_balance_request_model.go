// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeductOutstandingBalanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeductAmount(v string) *DeductOutstandingBalanceRequest
	GetDeductAmount() *string
	SetUid(v int64) *DeductOutstandingBalanceRequest
	GetUid() *int64
}

type DeductOutstandingBalanceRequest struct {
	// The Deducted Credit to be offset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 300
	DeductAmount *string `json:"DeductAmount,omitempty" xml:"DeductAmount,omitempty"`
	// Account UID of Distribution Customer.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1133166938931507
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DeductOutstandingBalanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeductOutstandingBalanceRequest) GoString() string {
	return s.String()
}

func (s *DeductOutstandingBalanceRequest) GetDeductAmount() *string {
	return s.DeductAmount
}

func (s *DeductOutstandingBalanceRequest) GetUid() *int64 {
	return s.Uid
}

func (s *DeductOutstandingBalanceRequest) SetDeductAmount(v string) *DeductOutstandingBalanceRequest {
	s.DeductAmount = &v
	return s
}

func (s *DeductOutstandingBalanceRequest) SetUid(v int64) *DeductOutstandingBalanceRequest {
	s.Uid = &v
	return s
}

func (s *DeductOutstandingBalanceRequest) Validate() error {
	return dara.Validate(s)
}
