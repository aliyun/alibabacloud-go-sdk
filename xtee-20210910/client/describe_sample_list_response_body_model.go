// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSampleListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeSampleListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeSampleListResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeSampleListResponseBodyResultObject) *DescribeSampleListResponseBody
	GetResultObject() []*DescribeSampleListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeSampleListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeSampleListResponseBody
	GetTotalPage() *int32
}

type DescribeSampleListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// Returned object
	ResultObject []*DescribeSampleListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total count.
	//
	// example:
	//
	// 6
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total pages
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeSampleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSampleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSampleListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSampleListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSampleListResponseBody) GetResultObject() []*DescribeSampleListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSampleListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeSampleListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeSampleListResponseBody) SetRequestId(v string) *DescribeSampleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSampleListResponseBody) SetCurrentPage(v int32) *DescribeSampleListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSampleListResponseBody) SetPageSize(v int32) *DescribeSampleListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSampleListResponseBody) SetResultObject(v []*DescribeSampleListResponseBodyResultObject) *DescribeSampleListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSampleListResponseBody) SetTotalItem(v int32) *DescribeSampleListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeSampleListResponseBody) SetTotalPage(v int32) *DescribeSampleListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeSampleListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSampleListResponseBodyResultObject struct {
	// Database ID.
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Sample tags.
	//
	// example:
	//
	// rm0102
	SampleTags *string `json:"sampleTags,omitempty" xml:"sampleTags,omitempty"`
	// Sample type
	//
	// example:
	//
	// PHONE
	SampleType *int32 `json:"sampleType,omitempty" xml:"sampleType,omitempty"`
	// Sample value.
	//
	// example:
	//
	// 1770000000
	SampleValue *string `json:"sampleValue,omitempty" xml:"sampleValue,omitempty"`
	// Update time.
	//
	// example:
	//
	// 1699450018265
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// Version number.
	//
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeSampleListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSampleListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeSampleListResponseBodyResultObject) GetSampleTags() *string {
	return s.SampleTags
}

func (s *DescribeSampleListResponseBodyResultObject) GetSampleType() *int32 {
	return s.SampleType
}

func (s *DescribeSampleListResponseBodyResultObject) GetSampleValue() *string {
	return s.SampleValue
}

func (s *DescribeSampleListResponseBodyResultObject) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeSampleListResponseBodyResultObject) GetVersion() *int32 {
	return s.Version
}

func (s *DescribeSampleListResponseBodyResultObject) SetId(v int64) *DescribeSampleListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeSampleListResponseBodyResultObject) SetSampleTags(v string) *DescribeSampleListResponseBodyResultObject {
	s.SampleTags = &v
	return s
}

func (s *DescribeSampleListResponseBodyResultObject) SetSampleType(v int32) *DescribeSampleListResponseBodyResultObject {
	s.SampleType = &v
	return s
}

func (s *DescribeSampleListResponseBodyResultObject) SetSampleValue(v string) *DescribeSampleListResponseBodyResultObject {
	s.SampleValue = &v
	return s
}

func (s *DescribeSampleListResponseBodyResultObject) SetUpdateTime(v int64) *DescribeSampleListResponseBodyResultObject {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSampleListResponseBodyResultObject) SetVersion(v int32) *DescribeSampleListResponseBodyResultObject {
	s.Version = &v
	return s
}

func (s *DescribeSampleListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
