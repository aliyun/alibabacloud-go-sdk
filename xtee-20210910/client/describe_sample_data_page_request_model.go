// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDataPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSampleDataPageRequest
	GetLang() *string
	SetCurrentPage(v int32) *DescribeSampleDataPageRequest
	GetCurrentPage() *int32
	SetDataValue(v string) *DescribeSampleDataPageRequest
	GetDataValue() *string
	SetPageSize(v int32) *DescribeSampleDataPageRequest
	GetPageSize() *int32
	SetRegId(v string) *DescribeSampleDataPageRequest
	GetRegId() *string
	SetUpdateBeginTime(v int64) *DescribeSampleDataPageRequest
	GetUpdateBeginTime() *int64
	SetUpdateEndTime(v int64) *DescribeSampleDataPageRequest
	GetUpdateEndTime() *int64
}

type DescribeSampleDataPageRequest struct {
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
	// Content of the list entered in the text box
	//
	// example:
	//
	// 1770000000
	DataValue *string `json:"dataValue,omitempty" xml:"dataValue,omitempty"`
	// Page size, default value is 10
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
	// Start time
	//
	// example:
	//
	// 1730429469000
	UpdateBeginTime *int64 `json:"updateBeginTime,omitempty" xml:"updateBeginTime,omitempty"`
	// End time
	//
	// example:
	//
	// 1730429469000
	UpdateEndTime *int64 `json:"updateEndTime,omitempty" xml:"updateEndTime,omitempty"`
}

func (s DescribeSampleDataPageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDataPageRequest) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataPageRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSampleDataPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSampleDataPageRequest) GetDataValue() *string {
	return s.DataValue
}

func (s *DescribeSampleDataPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSampleDataPageRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSampleDataPageRequest) GetUpdateBeginTime() *int64 {
	return s.UpdateBeginTime
}

func (s *DescribeSampleDataPageRequest) GetUpdateEndTime() *int64 {
	return s.UpdateEndTime
}

func (s *DescribeSampleDataPageRequest) SetLang(v string) *DescribeSampleDataPageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSampleDataPageRequest) SetCurrentPage(v int32) *DescribeSampleDataPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSampleDataPageRequest) SetDataValue(v string) *DescribeSampleDataPageRequest {
	s.DataValue = &v
	return s
}

func (s *DescribeSampleDataPageRequest) SetPageSize(v int32) *DescribeSampleDataPageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSampleDataPageRequest) SetRegId(v string) *DescribeSampleDataPageRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSampleDataPageRequest) SetUpdateBeginTime(v int64) *DescribeSampleDataPageRequest {
	s.UpdateBeginTime = &v
	return s
}

func (s *DescribeSampleDataPageRequest) SetUpdateEndTime(v int64) *DescribeSampleDataPageRequest {
	s.UpdateEndTime = &v
	return s
}

func (s *DescribeSampleDataPageRequest) Validate() error {
	return dara.Validate(s)
}
