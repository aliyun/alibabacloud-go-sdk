// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDataByBatchUUidPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSampleDataByBatchUUidPageRequest
	GetLang() *string
	SetBatchUuid(v string) *DescribeSampleDataByBatchUUidPageRequest
	GetBatchUuid() *string
	SetCurrentPage(v int32) *DescribeSampleDataByBatchUUidPageRequest
	GetCurrentPage() *int32
	SetDataValue(v string) *DescribeSampleDataByBatchUUidPageRequest
	GetDataValue() *string
	SetPageSize(v int32) *DescribeSampleDataByBatchUUidPageRequest
	GetPageSize() *int32
	SetRegId(v string) *DescribeSampleDataByBatchUUidPageRequest
	GetRegId() *string
	SetUpdateBeginTime(v int64) *DescribeSampleDataByBatchUUidPageRequest
	GetUpdateBeginTime() *int64
	SetUpdateEndTime(v int64) *DescribeSampleDataByBatchUUidPageRequest
	GetUpdateEndTime() *int64
}

type DescribeSampleDataByBatchUUidPageRequest struct {
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
	// Sample batch UUID
	//
	// example:
	//
	// 203f1ae65c0a41a49dc4f8a47946caa2
	BatchUuid *string `json:"batchUuid,omitempty" xml:"batchUuid,omitempty"`
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
	// 1770000000,1770000001
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
	// 1724986526000
	UpdateBeginTime *int64 `json:"updateBeginTime,omitempty" xml:"updateBeginTime,omitempty"`
	// End time
	//
	// example:
	//
	// 1724986526000
	UpdateEndTime *int64 `json:"updateEndTime,omitempty" xml:"updateEndTime,omitempty"`
}

func (s DescribeSampleDataByBatchUUidPageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDataByBatchUUidPageRequest) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataByBatchUUidPageRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSampleDataByBatchUUidPageRequest) GetBatchUuid() *string {
	return s.BatchUuid
}

func (s *DescribeSampleDataByBatchUUidPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSampleDataByBatchUUidPageRequest) GetDataValue() *string {
	return s.DataValue
}

func (s *DescribeSampleDataByBatchUUidPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSampleDataByBatchUUidPageRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSampleDataByBatchUUidPageRequest) GetUpdateBeginTime() *int64 {
	return s.UpdateBeginTime
}

func (s *DescribeSampleDataByBatchUUidPageRequest) GetUpdateEndTime() *int64 {
	return s.UpdateEndTime
}

func (s *DescribeSampleDataByBatchUUidPageRequest) SetLang(v string) *DescribeSampleDataByBatchUUidPageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageRequest) SetBatchUuid(v string) *DescribeSampleDataByBatchUUidPageRequest {
	s.BatchUuid = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageRequest) SetCurrentPage(v int32) *DescribeSampleDataByBatchUUidPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageRequest) SetDataValue(v string) *DescribeSampleDataByBatchUUidPageRequest {
	s.DataValue = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageRequest) SetPageSize(v int32) *DescribeSampleDataByBatchUUidPageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageRequest) SetRegId(v string) *DescribeSampleDataByBatchUUidPageRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageRequest) SetUpdateBeginTime(v int64) *DescribeSampleDataByBatchUUidPageRequest {
	s.UpdateBeginTime = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageRequest) SetUpdateEndTime(v int64) *DescribeSampleDataByBatchUUidPageRequest {
	s.UpdateEndTime = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageRequest) Validate() error {
	return dara.Validate(s)
}
