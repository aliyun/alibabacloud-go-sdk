// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSettledsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ListSettledsRequest
	GetAccountId() *string
	SetApplicant(v string) *ListSettledsRequest
	GetApplicant() *string
	SetEnterpriseName(v string) *ListSettledsRequest
	GetEnterpriseName() *string
	SetStatus(v string) *ListSettledsRequest
	GetStatus() *string
}

type ListSettledsRequest struct {
	// example:
	//
	// 273803534812230281
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 223352752411587433
	Applicant      *string `json:"applicant,omitempty" xml:"applicant,omitempty"`
	EnterpriseName *string `json:"enterpriseName,omitempty" xml:"enterpriseName,omitempty"`
	// example:
	//
	// 3
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListSettledsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSettledsRequest) GoString() string {
	return s.String()
}

func (s *ListSettledsRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ListSettledsRequest) GetApplicant() *string {
	return s.Applicant
}

func (s *ListSettledsRequest) GetEnterpriseName() *string {
	return s.EnterpriseName
}

func (s *ListSettledsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListSettledsRequest) SetAccountId(v string) *ListSettledsRequest {
	s.AccountId = &v
	return s
}

func (s *ListSettledsRequest) SetApplicant(v string) *ListSettledsRequest {
	s.Applicant = &v
	return s
}

func (s *ListSettledsRequest) SetEnterpriseName(v string) *ListSettledsRequest {
	s.EnterpriseName = &v
	return s
}

func (s *ListSettledsRequest) SetStatus(v string) *ListSettledsRequest {
	s.Status = &v
	return s
}

func (s *ListSettledsRequest) Validate() error {
	return dara.Validate(s)
}
