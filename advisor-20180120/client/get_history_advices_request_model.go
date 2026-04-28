// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoryAdvicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *GetHistoryAdvicesRequest
	GetEndDate() *string
	SetLanguage(v string) *GetHistoryAdvicesRequest
	GetLanguage() *string
	SetPageNum(v int32) *GetHistoryAdvicesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *GetHistoryAdvicesRequest
	GetPageSize() *int32
	SetProduct(v string) *GetHistoryAdvicesRequest
	GetProduct() *string
	SetReverse(v bool) *GetHistoryAdvicesRequest
	GetReverse() *bool
	SetSeverity(v string) *GetHistoryAdvicesRequest
	GetSeverity() *string
	SetStartDate(v string) *GetHistoryAdvicesRequest
	GetStartDate() *string
}

type GetHistoryAdvicesRequest struct {
	// example:
	//
	// 2023-07-01
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// example:
	//
	// 1
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// example:
	//
	// 2023-07-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetHistoryAdvicesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHistoryAdvicesRequest) GoString() string {
	return s.String()
}

func (s *GetHistoryAdvicesRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetHistoryAdvicesRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetHistoryAdvicesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetHistoryAdvicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetHistoryAdvicesRequest) GetProduct() *string {
	return s.Product
}

func (s *GetHistoryAdvicesRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *GetHistoryAdvicesRequest) GetSeverity() *string {
	return s.Severity
}

func (s *GetHistoryAdvicesRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetHistoryAdvicesRequest) SetEndDate(v string) *GetHistoryAdvicesRequest {
	s.EndDate = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetLanguage(v string) *GetHistoryAdvicesRequest {
	s.Language = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetPageNum(v int32) *GetHistoryAdvicesRequest {
	s.PageNum = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetPageSize(v int32) *GetHistoryAdvicesRequest {
	s.PageSize = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetProduct(v string) *GetHistoryAdvicesRequest {
	s.Product = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetReverse(v bool) *GetHistoryAdvicesRequest {
	s.Reverse = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetSeverity(v string) *GetHistoryAdvicesRequest {
	s.Severity = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetStartDate(v string) *GetHistoryAdvicesRequest {
	s.StartDate = &v
	return s
}

func (s *GetHistoryAdvicesRequest) Validate() error {
	return dara.Validate(s)
}
