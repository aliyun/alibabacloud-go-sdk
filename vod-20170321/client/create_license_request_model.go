// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *CreateLicenseRequest
	GetAccountId() *int64
	SetBeginTime(v string) *CreateLicenseRequest
	GetBeginTime() *string
	SetBusinessType(v string) *CreateLicenseRequest
	GetBusinessType() *string
	SetContractNo(v string) *CreateLicenseRequest
	GetContractNo() *string
	SetCustomerId(v int64) *CreateLicenseRequest
	GetCustomerId() *int64
	SetExpiredOn(v string) *CreateLicenseRequest
	GetExpiredOn() *string
	SetExtraInfo(v string) *CreateLicenseRequest
	GetExtraInfo() *string
	SetOperator(v string) *CreateLicenseRequest
	GetOperator() *string
	SetType(v string) *CreateLicenseRequest
	GetType() *string
}

type CreateLicenseRequest struct {
	AccountId *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// This parameter is required.
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	ContractNo   *string `json:"ContractNo,omitempty" xml:"ContractNo,omitempty"`
	CustomerId   *int64  `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	ExpiredOn    *string `json:"ExpiredOn,omitempty" xml:"ExpiredOn,omitempty"`
	// This parameter is required.
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	Operator  *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLicenseRequest) GoString() string {
	return s.String()
}

func (s *CreateLicenseRequest) GetAccountId() *int64 {
	return s.AccountId
}

func (s *CreateLicenseRequest) GetBeginTime() *string {
	return s.BeginTime
}

func (s *CreateLicenseRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CreateLicenseRequest) GetContractNo() *string {
	return s.ContractNo
}

func (s *CreateLicenseRequest) GetCustomerId() *int64 {
	return s.CustomerId
}

func (s *CreateLicenseRequest) GetExpiredOn() *string {
	return s.ExpiredOn
}

func (s *CreateLicenseRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *CreateLicenseRequest) GetOperator() *string {
	return s.Operator
}

func (s *CreateLicenseRequest) GetType() *string {
	return s.Type
}

func (s *CreateLicenseRequest) SetAccountId(v int64) *CreateLicenseRequest {
	s.AccountId = &v
	return s
}

func (s *CreateLicenseRequest) SetBeginTime(v string) *CreateLicenseRequest {
	s.BeginTime = &v
	return s
}

func (s *CreateLicenseRequest) SetBusinessType(v string) *CreateLicenseRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateLicenseRequest) SetContractNo(v string) *CreateLicenseRequest {
	s.ContractNo = &v
	return s
}

func (s *CreateLicenseRequest) SetCustomerId(v int64) *CreateLicenseRequest {
	s.CustomerId = &v
	return s
}

func (s *CreateLicenseRequest) SetExpiredOn(v string) *CreateLicenseRequest {
	s.ExpiredOn = &v
	return s
}

func (s *CreateLicenseRequest) SetExtraInfo(v string) *CreateLicenseRequest {
	s.ExtraInfo = &v
	return s
}

func (s *CreateLicenseRequest) SetOperator(v string) *CreateLicenseRequest {
	s.Operator = &v
	return s
}

func (s *CreateLicenseRequest) SetType(v string) *CreateLicenseRequest {
	s.Type = &v
	return s
}

func (s *CreateLicenseRequest) Validate() error {
	return dara.Validate(s)
}
