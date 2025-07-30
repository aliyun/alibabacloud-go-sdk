// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheAnalysisReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBigKeys(v []map[string]interface{}) *DescribeCacheAnalysisReportResponseBody
	GetBigKeys() []map[string]interface{}
	SetHotKeys(v []map[string]interface{}) *DescribeCacheAnalysisReportResponseBody
	GetHotKeys() []map[string]interface{}
	SetPageNumber(v int32) *DescribeCacheAnalysisReportResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeCacheAnalysisReportResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int32) *DescribeCacheAnalysisReportResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCacheAnalysisReportResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeCacheAnalysisReportResponseBody
	GetTotalRecordCount() *int32
}

type DescribeCacheAnalysisReportResponseBody struct {
	// Details of the large keys.
	BigKeys []map[string]interface{} `json:"BigKeys,omitempty" xml:"BigKeys,omitempty" type:"Repeated"`
	// Details of the hotkeys.
	//
	// > This parameter is not returned because Tair (Redis OSS-compatible) does not support hotkey analytics.
	HotKeys []map[string]interface{} `json:"HotKeys,omitempty" xml:"HotKeys,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A057C066-C3F5-4CC9-9FE4-A8D8B0DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 160
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeCacheAnalysisReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportResponseBody) GetBigKeys() []map[string]interface{} {
	return s.BigKeys
}

func (s *DescribeCacheAnalysisReportResponseBody) GetHotKeys() []map[string]interface{} {
	return s.HotKeys
}

func (s *DescribeCacheAnalysisReportResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCacheAnalysisReportResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeCacheAnalysisReportResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCacheAnalysisReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCacheAnalysisReportResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeCacheAnalysisReportResponseBody) SetBigKeys(v []map[string]interface{}) *DescribeCacheAnalysisReportResponseBody {
	s.BigKeys = v
	return s
}

func (s *DescribeCacheAnalysisReportResponseBody) SetHotKeys(v []map[string]interface{}) *DescribeCacheAnalysisReportResponseBody {
	s.HotKeys = v
	return s
}

func (s *DescribeCacheAnalysisReportResponseBody) SetPageNumber(v int32) *DescribeCacheAnalysisReportResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCacheAnalysisReportResponseBody) SetPageRecordCount(v int32) *DescribeCacheAnalysisReportResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeCacheAnalysisReportResponseBody) SetPageSize(v int32) *DescribeCacheAnalysisReportResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCacheAnalysisReportResponseBody) SetRequestId(v string) *DescribeCacheAnalysisReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCacheAnalysisReportResponseBody) SetTotalRecordCount(v int32) *DescribeCacheAnalysisReportResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeCacheAnalysisReportResponseBody) Validate() error {
	return dara.Validate(s)
}
