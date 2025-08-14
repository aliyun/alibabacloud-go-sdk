// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSamplebatchPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSamplebatchPageRequest
	GetLang() *string
	SetCurrentPage(v int32) *DescribeSamplebatchPageRequest
	GetCurrentPage() *int32
	SetDataValue(v string) *DescribeSamplebatchPageRequest
	GetDataValue() *string
	SetPageSize(v int32) *DescribeSamplebatchPageRequest
	GetPageSize() *int32
	SetRegId(v string) *DescribeSamplebatchPageRequest
	GetRegId() *string
}

type DescribeSamplebatchPageRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
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
	// Content of the list entered in the text box
	//
	// example:
	//
	// 1770000000
	DataValue *string `json:"dataValue,omitempty" xml:"dataValue,omitempty"`
	// Page size, with a default value of 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeSamplebatchPageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSamplebatchPageRequest) GoString() string {
	return s.String()
}

func (s *DescribeSamplebatchPageRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSamplebatchPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSamplebatchPageRequest) GetDataValue() *string {
	return s.DataValue
}

func (s *DescribeSamplebatchPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSamplebatchPageRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSamplebatchPageRequest) SetLang(v string) *DescribeSamplebatchPageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSamplebatchPageRequest) SetCurrentPage(v int32) *DescribeSamplebatchPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSamplebatchPageRequest) SetDataValue(v string) *DescribeSamplebatchPageRequest {
	s.DataValue = &v
	return s
}

func (s *DescribeSamplebatchPageRequest) SetPageSize(v int32) *DescribeSamplebatchPageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSamplebatchPageRequest) SetRegId(v string) *DescribeSamplebatchPageRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSamplebatchPageRequest) Validate() error {
	return dara.Validate(s)
}
