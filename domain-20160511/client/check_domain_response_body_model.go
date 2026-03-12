// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvail(v int32) *CheckDomainResponseBody
	GetAvail() *int32
	SetFeeCommand(v string) *CheckDomainResponseBody
	GetFeeCommand() *string
	SetFeeCurrency(v string) *CheckDomainResponseBody
	GetFeeCurrency() *string
	SetFeeFee(v string) *CheckDomainResponseBody
	GetFeeFee() *string
	SetFeePeriod(v int32) *CheckDomainResponseBody
	GetFeePeriod() *int32
	SetName(v string) *CheckDomainResponseBody
	GetName() *string
	SetReason(v string) *CheckDomainResponseBody
	GetReason() *string
	SetRequestId(v string) *CheckDomainResponseBody
	GetRequestId() *string
	SetRmbFee(v string) *CheckDomainResponseBody
	GetRmbFee() *string
}

type CheckDomainResponseBody struct {
	Avail       *int32  `json:"Avail,omitempty" xml:"Avail,omitempty"`
	FeeCommand  *string `json:"FeeCommand,omitempty" xml:"FeeCommand,omitempty"`
	FeeCurrency *string `json:"FeeCurrency,omitempty" xml:"FeeCurrency,omitempty"`
	FeeFee      *string `json:"FeeFee,omitempty" xml:"FeeFee,omitempty"`
	FeePeriod   *int32  `json:"FeePeriod,omitempty" xml:"FeePeriod,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RmbFee      *string `json:"RmbFee,omitempty" xml:"RmbFee,omitempty"`
}

func (s CheckDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDomainResponseBody) GetAvail() *int32 {
	return s.Avail
}

func (s *CheckDomainResponseBody) GetFeeCommand() *string {
	return s.FeeCommand
}

func (s *CheckDomainResponseBody) GetFeeCurrency() *string {
	return s.FeeCurrency
}

func (s *CheckDomainResponseBody) GetFeeFee() *string {
	return s.FeeFee
}

func (s *CheckDomainResponseBody) GetFeePeriod() *int32 {
	return s.FeePeriod
}

func (s *CheckDomainResponseBody) GetName() *string {
	return s.Name
}

func (s *CheckDomainResponseBody) GetReason() *string {
	return s.Reason
}

func (s *CheckDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDomainResponseBody) GetRmbFee() *string {
	return s.RmbFee
}

func (s *CheckDomainResponseBody) SetAvail(v int32) *CheckDomainResponseBody {
	s.Avail = &v
	return s
}

func (s *CheckDomainResponseBody) SetFeeCommand(v string) *CheckDomainResponseBody {
	s.FeeCommand = &v
	return s
}

func (s *CheckDomainResponseBody) SetFeeCurrency(v string) *CheckDomainResponseBody {
	s.FeeCurrency = &v
	return s
}

func (s *CheckDomainResponseBody) SetFeeFee(v string) *CheckDomainResponseBody {
	s.FeeFee = &v
	return s
}

func (s *CheckDomainResponseBody) SetFeePeriod(v int32) *CheckDomainResponseBody {
	s.FeePeriod = &v
	return s
}

func (s *CheckDomainResponseBody) SetName(v string) *CheckDomainResponseBody {
	s.Name = &v
	return s
}

func (s *CheckDomainResponseBody) SetReason(v string) *CheckDomainResponseBody {
	s.Reason = &v
	return s
}

func (s *CheckDomainResponseBody) SetRequestId(v string) *CheckDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDomainResponseBody) SetRmbFee(v string) *CheckDomainResponseBody {
	s.RmbFee = &v
	return s
}

func (s *CheckDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
