// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLicenseInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *ListLicenseInfosRequest
	GetAccountId() *int64
	SetBusinessType(v string) *ListLicenseInfosRequest
	GetBusinessType() *string
	SetContractNo(v string) *ListLicenseInfosRequest
	GetContractNo() *string
	SetCustomerId(v int64) *ListLicenseInfosRequest
	GetCustomerId() *int64
	SetEndBeginTime(v string) *ListLicenseInfosRequest
	GetEndBeginTime() *string
	SetEndExpiredOn(v string) *ListLicenseInfosRequest
	GetEndExpiredOn() *string
	SetExtraInfo(v string) *ListLicenseInfosRequest
	GetExtraInfo() *string
	SetLicenseId(v string) *ListLicenseInfosRequest
	GetLicenseId() *string
	SetPageNo(v int64) *ListLicenseInfosRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListLicenseInfosRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListLicenseInfosRequest
	GetSortBy() *string
	SetStartBeginTime(v string) *ListLicenseInfosRequest
	GetStartBeginTime() *string
	SetStartExpiredOn(v string) *ListLicenseInfosRequest
	GetStartExpiredOn() *string
	SetStatus(v string) *ListLicenseInfosRequest
	GetStatus() *string
	SetType(v string) *ListLicenseInfosRequest
	GetType() *string
}

type ListLicenseInfosRequest struct {
	AccountId      *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	BusinessType   *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	ContractNo     *string `json:"ContractNo,omitempty" xml:"ContractNo,omitempty"`
	CustomerId     *int64  `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	EndBeginTime   *string `json:"EndBeginTime,omitempty" xml:"EndBeginTime,omitempty"`
	EndExpiredOn   *string `json:"EndExpiredOn,omitempty" xml:"EndExpiredOn,omitempty"`
	ExtraInfo      *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	LicenseId      *string `json:"LicenseId,omitempty" xml:"LicenseId,omitempty"`
	PageNo         *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy         *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	StartBeginTime *string `json:"StartBeginTime,omitempty" xml:"StartBeginTime,omitempty"`
	StartExpiredOn *string `json:"StartExpiredOn,omitempty" xml:"StartExpiredOn,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListLicenseInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLicenseInfosRequest) GoString() string {
	return s.String()
}

func (s *ListLicenseInfosRequest) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListLicenseInfosRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ListLicenseInfosRequest) GetContractNo() *string {
	return s.ContractNo
}

func (s *ListLicenseInfosRequest) GetCustomerId() *int64 {
	return s.CustomerId
}

func (s *ListLicenseInfosRequest) GetEndBeginTime() *string {
	return s.EndBeginTime
}

func (s *ListLicenseInfosRequest) GetEndExpiredOn() *string {
	return s.EndExpiredOn
}

func (s *ListLicenseInfosRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *ListLicenseInfosRequest) GetLicenseId() *string {
	return s.LicenseId
}

func (s *ListLicenseInfosRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListLicenseInfosRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListLicenseInfosRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLicenseInfosRequest) GetStartBeginTime() *string {
	return s.StartBeginTime
}

func (s *ListLicenseInfosRequest) GetStartExpiredOn() *string {
	return s.StartExpiredOn
}

func (s *ListLicenseInfosRequest) GetStatus() *string {
	return s.Status
}

func (s *ListLicenseInfosRequest) GetType() *string {
	return s.Type
}

func (s *ListLicenseInfosRequest) SetAccountId(v int64) *ListLicenseInfosRequest {
	s.AccountId = &v
	return s
}

func (s *ListLicenseInfosRequest) SetBusinessType(v string) *ListLicenseInfosRequest {
	s.BusinessType = &v
	return s
}

func (s *ListLicenseInfosRequest) SetContractNo(v string) *ListLicenseInfosRequest {
	s.ContractNo = &v
	return s
}

func (s *ListLicenseInfosRequest) SetCustomerId(v int64) *ListLicenseInfosRequest {
	s.CustomerId = &v
	return s
}

func (s *ListLicenseInfosRequest) SetEndBeginTime(v string) *ListLicenseInfosRequest {
	s.EndBeginTime = &v
	return s
}

func (s *ListLicenseInfosRequest) SetEndExpiredOn(v string) *ListLicenseInfosRequest {
	s.EndExpiredOn = &v
	return s
}

func (s *ListLicenseInfosRequest) SetExtraInfo(v string) *ListLicenseInfosRequest {
	s.ExtraInfo = &v
	return s
}

func (s *ListLicenseInfosRequest) SetLicenseId(v string) *ListLicenseInfosRequest {
	s.LicenseId = &v
	return s
}

func (s *ListLicenseInfosRequest) SetPageNo(v int64) *ListLicenseInfosRequest {
	s.PageNo = &v
	return s
}

func (s *ListLicenseInfosRequest) SetPageSize(v int64) *ListLicenseInfosRequest {
	s.PageSize = &v
	return s
}

func (s *ListLicenseInfosRequest) SetSortBy(v string) *ListLicenseInfosRequest {
	s.SortBy = &v
	return s
}

func (s *ListLicenseInfosRequest) SetStartBeginTime(v string) *ListLicenseInfosRequest {
	s.StartBeginTime = &v
	return s
}

func (s *ListLicenseInfosRequest) SetStartExpiredOn(v string) *ListLicenseInfosRequest {
	s.StartExpiredOn = &v
	return s
}

func (s *ListLicenseInfosRequest) SetStatus(v string) *ListLicenseInfosRequest {
	s.Status = &v
	return s
}

func (s *ListLicenseInfosRequest) SetType(v string) *ListLicenseInfosRequest {
	s.Type = &v
	return s
}

func (s *ListLicenseInfosRequest) Validate() error {
	return dara.Validate(s)
}
