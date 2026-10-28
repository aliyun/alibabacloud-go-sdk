// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerQuotaRecordListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *CustomerQuotaRecordListRequest
	GetEndDate() *string
	SetEndUserPk(v int64) *CustomerQuotaRecordListRequest
	GetEndUserPk() *int64
	SetLanguage(v string) *CustomerQuotaRecordListRequest
	GetLanguage() *string
	SetOperationType(v string) *CustomerQuotaRecordListRequest
	GetOperationType() *string
	SetPageNo(v int32) *CustomerQuotaRecordListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *CustomerQuotaRecordListRequest
	GetPageSize() *int32
	SetStartDate(v string) *CustomerQuotaRecordListRequest
	GetStartDate() *string
}

type CustomerQuotaRecordListRequest struct {
	// The end date.
	//
	// Format: yyyy-MM-dd.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-09-24
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The UID of the customer.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5113766248601929
	EndUserPk *int64 `json:"EndUserPk,omitempty" xml:"EndUserPk,omitempty"`
	// The language. Default value: en. Valid values:
	//
	// - en: English.
	//
	// - zh: Chinese.
	//
	// - ja: Japanese.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The operation type. Valid values:
	//
	// - all: All.
	//
	// - quota_create: Create.
	//
	// - quota_amount_adjust: Quota adjustment.
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The page number. Pages start from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start date.
	//
	// Format: yyyy-MM-dd.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-02
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s CustomerQuotaRecordListRequest) String() string {
	return dara.Prettify(s)
}

func (s CustomerQuotaRecordListRequest) GoString() string {
	return s.String()
}

func (s *CustomerQuotaRecordListRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *CustomerQuotaRecordListRequest) GetEndUserPk() *int64 {
	return s.EndUserPk
}

func (s *CustomerQuotaRecordListRequest) GetLanguage() *string {
	return s.Language
}

func (s *CustomerQuotaRecordListRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *CustomerQuotaRecordListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *CustomerQuotaRecordListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CustomerQuotaRecordListRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *CustomerQuotaRecordListRequest) SetEndDate(v string) *CustomerQuotaRecordListRequest {
	s.EndDate = &v
	return s
}

func (s *CustomerQuotaRecordListRequest) SetEndUserPk(v int64) *CustomerQuotaRecordListRequest {
	s.EndUserPk = &v
	return s
}

func (s *CustomerQuotaRecordListRequest) SetLanguage(v string) *CustomerQuotaRecordListRequest {
	s.Language = &v
	return s
}

func (s *CustomerQuotaRecordListRequest) SetOperationType(v string) *CustomerQuotaRecordListRequest {
	s.OperationType = &v
	return s
}

func (s *CustomerQuotaRecordListRequest) SetPageNo(v int32) *CustomerQuotaRecordListRequest {
	s.PageNo = &v
	return s
}

func (s *CustomerQuotaRecordListRequest) SetPageSize(v int32) *CustomerQuotaRecordListRequest {
	s.PageSize = &v
	return s
}

func (s *CustomerQuotaRecordListRequest) SetStartDate(v string) *CustomerQuotaRecordListRequest {
	s.StartDate = &v
	return s
}

func (s *CustomerQuotaRecordListRequest) Validate() error {
	return dara.Validate(s)
}
