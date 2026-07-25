// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeJobCreditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreditStatus(v string) *GetYikeJobCreditResponseBody
	GetCreditStatus() *string
	SetJobCreditCost(v float64) *GetYikeJobCreditResponseBody
	GetJobCreditCost() *float64
	SetJobId(v string) *GetYikeJobCreditResponseBody
	GetJobId() *string
	SetRequestId(v string) *GetYikeJobCreditResponseBody
	GetRequestId() *string
}

type GetYikeJobCreditResponseBody struct {
	// example:
	//
	// success
	CreditStatus *string `json:"CreditStatus,omitempty" xml:"CreditStatus,omitempty"`
	// example:
	//
	// 20
	JobCreditCost *float64 `json:"JobCreditCost,omitempty" xml:"JobCreditCost,omitempty"`
	// example:
	//
	// ag_12412424****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetYikeJobCreditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetYikeJobCreditResponseBody) GoString() string {
	return s.String()
}

func (s *GetYikeJobCreditResponseBody) GetCreditStatus() *string {
	return s.CreditStatus
}

func (s *GetYikeJobCreditResponseBody) GetJobCreditCost() *float64 {
	return s.JobCreditCost
}

func (s *GetYikeJobCreditResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetYikeJobCreditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetYikeJobCreditResponseBody) SetCreditStatus(v string) *GetYikeJobCreditResponseBody {
	s.CreditStatus = &v
	return s
}

func (s *GetYikeJobCreditResponseBody) SetJobCreditCost(v float64) *GetYikeJobCreditResponseBody {
	s.JobCreditCost = &v
	return s
}

func (s *GetYikeJobCreditResponseBody) SetJobId(v string) *GetYikeJobCreditResponseBody {
	s.JobId = &v
	return s
}

func (s *GetYikeJobCreditResponseBody) SetRequestId(v string) *GetYikeJobCreditResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetYikeJobCreditResponseBody) Validate() error {
	return dara.Validate(s)
}
