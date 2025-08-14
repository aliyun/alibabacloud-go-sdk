// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSampleListRequest
	GetLang() *string
	SetCurrentPage(v int32) *DescribeSampleListRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeSampleListRequest
	GetPageSize() *int32
	SetRegId(v string) *DescribeSampleListRequest
	GetRegId() *string
	SetSampleType(v string) *DescribeSampleListRequest
	GetSampleType() *string
	SetSampleValue(v string) *DescribeSampleListRequest
	GetSampleValue() *string
}

type DescribeSampleListRequest struct {
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
	// Sample type
	//
	// example:
	//
	// PHONE
	SampleType *string `json:"sampleType,omitempty" xml:"sampleType,omitempty"`
	// Sample data value.
	//
	// example:
	//
	// 1770000000
	SampleValue *string `json:"sampleValue,omitempty" xml:"sampleValue,omitempty"`
}

func (s DescribeSampleListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSampleListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSampleListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSampleListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSampleListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSampleListRequest) GetSampleType() *string {
	return s.SampleType
}

func (s *DescribeSampleListRequest) GetSampleValue() *string {
	return s.SampleValue
}

func (s *DescribeSampleListRequest) SetLang(v string) *DescribeSampleListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSampleListRequest) SetCurrentPage(v int32) *DescribeSampleListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSampleListRequest) SetPageSize(v int32) *DescribeSampleListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSampleListRequest) SetRegId(v string) *DescribeSampleListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSampleListRequest) SetSampleType(v string) *DescribeSampleListRequest {
	s.SampleType = &v
	return s
}

func (s *DescribeSampleListRequest) SetSampleValue(v string) *DescribeSampleListRequest {
	s.SampleValue = &v
	return s
}

func (s *DescribeSampleListRequest) Validate() error {
	return dara.Validate(s)
}
