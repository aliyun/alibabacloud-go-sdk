// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLicenseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *ModifyLicenseInfoRequest
	GetAccountId() *int64
	SetBeginTime(v string) *ModifyLicenseInfoRequest
	GetBeginTime() *string
	SetBusinessType(v string) *ModifyLicenseInfoRequest
	GetBusinessType() *string
	SetContractNo(v string) *ModifyLicenseInfoRequest
	GetContractNo() *string
	SetCustomerId(v int64) *ModifyLicenseInfoRequest
	GetCustomerId() *int64
	SetExpiredOn(v string) *ModifyLicenseInfoRequest
	GetExpiredOn() *string
	SetExtraInfo(v string) *ModifyLicenseInfoRequest
	GetExtraInfo() *string
	SetLicenseId(v string) *ModifyLicenseInfoRequest
	GetLicenseId() *string
	SetOperator(v string) *ModifyLicenseInfoRequest
	GetOperator() *string
	SetType(v string) *ModifyLicenseInfoRequest
	GetType() *string
}

type ModifyLicenseInfoRequest struct {
	AccountId *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// This parameter is required.
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	ContractNo   *string `json:"ContractNo,omitempty" xml:"ContractNo,omitempty"`
	CustomerId   *int64  `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	ExpiredOn    *string `json:"ExpiredOn,omitempty" xml:"ExpiredOn,omitempty"`
	// This parameter is required.
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// This parameter is required.
	LicenseId *string `json:"LicenseId,omitempty" xml:"LicenseId,omitempty"`
	Operator  *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyLicenseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLicenseInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyLicenseInfoRequest) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ModifyLicenseInfoRequest) GetBeginTime() *string {
	return s.BeginTime
}

func (s *ModifyLicenseInfoRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ModifyLicenseInfoRequest) GetContractNo() *string {
	return s.ContractNo
}

func (s *ModifyLicenseInfoRequest) GetCustomerId() *int64 {
	return s.CustomerId
}

func (s *ModifyLicenseInfoRequest) GetExpiredOn() *string {
	return s.ExpiredOn
}

func (s *ModifyLicenseInfoRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *ModifyLicenseInfoRequest) GetLicenseId() *string {
	return s.LicenseId
}

func (s *ModifyLicenseInfoRequest) GetOperator() *string {
	return s.Operator
}

func (s *ModifyLicenseInfoRequest) GetType() *string {
	return s.Type
}

func (s *ModifyLicenseInfoRequest) SetAccountId(v int64) *ModifyLicenseInfoRequest {
	s.AccountId = &v
	return s
}

func (s *ModifyLicenseInfoRequest) SetBeginTime(v string) *ModifyLicenseInfoRequest {
	s.BeginTime = &v
	return s
}

func (s *ModifyLicenseInfoRequest) SetBusinessType(v string) *ModifyLicenseInfoRequest {
	s.BusinessType = &v
	return s
}

func (s *ModifyLicenseInfoRequest) SetContractNo(v string) *ModifyLicenseInfoRequest {
	s.ContractNo = &v
	return s
}

func (s *ModifyLicenseInfoRequest) SetCustomerId(v int64) *ModifyLicenseInfoRequest {
	s.CustomerId = &v
	return s
}

func (s *ModifyLicenseInfoRequest) SetExpiredOn(v string) *ModifyLicenseInfoRequest {
	s.ExpiredOn = &v
	return s
}

func (s *ModifyLicenseInfoRequest) SetExtraInfo(v string) *ModifyLicenseInfoRequest {
	s.ExtraInfo = &v
	return s
}

func (s *ModifyLicenseInfoRequest) SetLicenseId(v string) *ModifyLicenseInfoRequest {
	s.LicenseId = &v
	return s
}

func (s *ModifyLicenseInfoRequest) SetOperator(v string) *ModifyLicenseInfoRequest {
	s.Operator = &v
	return s
}

func (s *ModifyLicenseInfoRequest) SetType(v string) *ModifyLicenseInfoRequest {
	s.Type = &v
	return s
}

func (s *ModifyLicenseInfoRequest) Validate() error {
	return dara.Validate(s)
}
