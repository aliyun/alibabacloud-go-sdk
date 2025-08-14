// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePocTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribePocTaskListRequest
	GetLang() *string
	SetCurrentPage(v int32) *DescribePocTaskListRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribePocTaskListRequest
	GetPageSize() *int32
	SetRegId(v string) *DescribePocTaskListRequest
	GetRegId() *string
	SetType(v string) *DescribePocTaskListRequest
	GetType() *string
}

type DescribePocTaskListRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, with a default value of 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Query type.
	//
	// example:
	//
	// SAF_CONSOLE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribePocTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePocTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribePocTaskListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePocTaskListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePocTaskListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePocTaskListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribePocTaskListRequest) GetType() *string {
	return s.Type
}

func (s *DescribePocTaskListRequest) SetLang(v string) *DescribePocTaskListRequest {
	s.Lang = &v
	return s
}

func (s *DescribePocTaskListRequest) SetCurrentPage(v int32) *DescribePocTaskListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePocTaskListRequest) SetPageSize(v int32) *DescribePocTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePocTaskListRequest) SetRegId(v string) *DescribePocTaskListRequest {
	s.RegId = &v
	return s
}

func (s *DescribePocTaskListRequest) SetType(v string) *DescribePocTaskListRequest {
	s.Type = &v
	return s
}

func (s *DescribePocTaskListRequest) Validate() error {
	return dara.Validate(s)
}
