// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAutomaticWriteOffChangeRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerUid(v int64) *QueryAutomaticWriteOffChangeRecordsRequest
	GetCustomerUid() *int64
	SetEndDate(v string) *QueryAutomaticWriteOffChangeRecordsRequest
	GetEndDate() *string
	SetLanguage(v string) *QueryAutomaticWriteOffChangeRecordsRequest
	GetLanguage() *string
	SetPageNo(v int32) *QueryAutomaticWriteOffChangeRecordsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryAutomaticWriteOffChangeRecordsRequest
	GetPageSize() *int32
	SetStartDate(v string) *QueryAutomaticWriteOffChangeRecordsRequest
	GetStartDate() *string
}

type QueryAutomaticWriteOffChangeRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	CustomerUid *int64 `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	// example:
	//
	// 2026-05-20
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2026-06-20
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryAutomaticWriteOffChangeRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAutomaticWriteOffChangeRecordsRequest) GoString() string {
	return s.String()
}

func (s *QueryAutomaticWriteOffChangeRecordsRequest) GetCustomerUid() *int64 {
	return s.CustomerUid
}

func (s *QueryAutomaticWriteOffChangeRecordsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryAutomaticWriteOffChangeRecordsRequest) GetLanguage() *string {
	return s.Language
}

func (s *QueryAutomaticWriteOffChangeRecordsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryAutomaticWriteOffChangeRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryAutomaticWriteOffChangeRecordsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryAutomaticWriteOffChangeRecordsRequest) SetCustomerUid(v int64) *QueryAutomaticWriteOffChangeRecordsRequest {
	s.CustomerUid = &v
	return s
}

func (s *QueryAutomaticWriteOffChangeRecordsRequest) SetEndDate(v string) *QueryAutomaticWriteOffChangeRecordsRequest {
	s.EndDate = &v
	return s
}

func (s *QueryAutomaticWriteOffChangeRecordsRequest) SetLanguage(v string) *QueryAutomaticWriteOffChangeRecordsRequest {
	s.Language = &v
	return s
}

func (s *QueryAutomaticWriteOffChangeRecordsRequest) SetPageNo(v int32) *QueryAutomaticWriteOffChangeRecordsRequest {
	s.PageNo = &v
	return s
}

func (s *QueryAutomaticWriteOffChangeRecordsRequest) SetPageSize(v int32) *QueryAutomaticWriteOffChangeRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAutomaticWriteOffChangeRecordsRequest) SetStartDate(v string) *QueryAutomaticWriteOffChangeRecordsRequest {
	s.StartDate = &v
	return s
}

func (s *QueryAutomaticWriteOffChangeRecordsRequest) Validate() error {
	return dara.Validate(s)
}
