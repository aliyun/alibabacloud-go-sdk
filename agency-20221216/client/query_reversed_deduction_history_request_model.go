// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReversedDeductionHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *QueryReversedDeductionHistoryRequest
	GetEndDate() *string
	SetLanguage(v string) *QueryReversedDeductionHistoryRequest
	GetLanguage() *string
	SetPageNo(v int32) *QueryReversedDeductionHistoryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryReversedDeductionHistoryRequest
	GetPageSize() *int32
	SetStartDate(v string) *QueryReversedDeductionHistoryRequest
	GetStartDate() *string
	SetUid(v int64) *QueryReversedDeductionHistoryRequest
	GetUid() *int64
}

type QueryReversedDeductionHistoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2023-12-31
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-07-31
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s QueryReversedDeductionHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryReversedDeductionHistoryRequest) GoString() string {
	return s.String()
}

func (s *QueryReversedDeductionHistoryRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryReversedDeductionHistoryRequest) GetLanguage() *string {
	return s.Language
}

func (s *QueryReversedDeductionHistoryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryReversedDeductionHistoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryReversedDeductionHistoryRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryReversedDeductionHistoryRequest) GetUid() *int64 {
	return s.Uid
}

func (s *QueryReversedDeductionHistoryRequest) SetEndDate(v string) *QueryReversedDeductionHistoryRequest {
	s.EndDate = &v
	return s
}

func (s *QueryReversedDeductionHistoryRequest) SetLanguage(v string) *QueryReversedDeductionHistoryRequest {
	s.Language = &v
	return s
}

func (s *QueryReversedDeductionHistoryRequest) SetPageNo(v int32) *QueryReversedDeductionHistoryRequest {
	s.PageNo = &v
	return s
}

func (s *QueryReversedDeductionHistoryRequest) SetPageSize(v int32) *QueryReversedDeductionHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *QueryReversedDeductionHistoryRequest) SetStartDate(v string) *QueryReversedDeductionHistoryRequest {
	s.StartDate = &v
	return s
}

func (s *QueryReversedDeductionHistoryRequest) SetUid(v int64) *QueryReversedDeductionHistoryRequest {
	s.Uid = &v
	return s
}

func (s *QueryReversedDeductionHistoryRequest) Validate() error {
	return dara.Validate(s)
}
