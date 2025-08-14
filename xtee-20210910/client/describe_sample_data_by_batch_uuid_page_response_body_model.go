// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDataByBatchUUidPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSampleDataByBatchUUidPageResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeSampleDataByBatchUUidPageResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeSampleDataByBatchUUidPageResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeSampleDataByBatchUUidPageResponseBodyResultObject) *DescribeSampleDataByBatchUUidPageResponseBody
	GetResultObject() []*DescribeSampleDataByBatchUUidPageResponseBodyResultObject
	SetTotalItem(v int32) *DescribeSampleDataByBatchUUidPageResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeSampleDataByBatchUUidPageResponseBody
	GetTotalPage() *int32
}

type DescribeSampleDataByBatchUUidPageResponseBody struct {
	// ID of the request
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
	// Page size, default value is 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object
	ResultObject []*DescribeSampleDataByBatchUUidPageResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items
	//
	// example:
	//
	// 6
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeSampleDataByBatchUUidPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDataByBatchUUidPageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataByBatchUUidPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSampleDataByBatchUUidPageResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSampleDataByBatchUUidPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSampleDataByBatchUUidPageResponseBody) GetResultObject() []*DescribeSampleDataByBatchUUidPageResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSampleDataByBatchUUidPageResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeSampleDataByBatchUUidPageResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeSampleDataByBatchUUidPageResponseBody) SetRequestId(v string) *DescribeSampleDataByBatchUUidPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponseBody) SetCurrentPage(v int32) *DescribeSampleDataByBatchUUidPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponseBody) SetPageSize(v int32) *DescribeSampleDataByBatchUUidPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponseBody) SetResultObject(v []*DescribeSampleDataByBatchUUidPageResponseBodyResultObject) *DescribeSampleDataByBatchUUidPageResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponseBody) SetTotalItem(v int32) *DescribeSampleDataByBatchUUidPageResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponseBody) SetTotalPage(v int32) *DescribeSampleDataByBatchUUidPageResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSampleDataByBatchUUidPageResponseBodyResultObject struct {
	// Sample batch name
	//
	// example:
	//
	// 白样本
	BatchName *string `json:"batchName,omitempty" xml:"batchName,omitempty"`
	// Creator
	//
	// example:
	//
	// 1519714049632764
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// Sample type
	//
	// example:
	//
	// pass
	DataTagType *string `json:"dataTagType,omitempty" xml:"dataTagType,omitempty"`
	// Data content
	//
	// example:
	//
	// 177000001
	DataValue *string `json:"dataValue,omitempty" xml:"dataValue,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Modification time
	//
	// example:
	//
	// 1565701886000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Last source
	//
	// example:
	//
	// Console-Text
	LastSourceType *string `json:"lastSourceType,omitempty" xml:"lastSourceType,omitempty"`
	// Sample batch UUID
	//
	// example:
	//
	// 203f1ae65c0a41a49dc4f8a47946caa2
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// Version
	//
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeSampleDataByBatchUUidPageResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDataByBatchUUidPageResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) GetBatchName() *string {
	return s.BatchName
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) GetCreator() *string {
	return s.Creator
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) GetDataTagType() *string {
	return s.DataTagType
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) GetDataValue() *string {
	return s.DataValue
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) GetLastSourceType() *string {
	return s.LastSourceType
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) GetVersion() *int32 {
	return s.Version
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) SetBatchName(v string) *DescribeSampleDataByBatchUUidPageResponseBodyResultObject {
	s.BatchName = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) SetCreator(v string) *DescribeSampleDataByBatchUUidPageResponseBodyResultObject {
	s.Creator = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) SetDataTagType(v string) *DescribeSampleDataByBatchUUidPageResponseBodyResultObject {
	s.DataTagType = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) SetDataValue(v string) *DescribeSampleDataByBatchUUidPageResponseBodyResultObject {
	s.DataValue = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) SetGmtCreate(v int64) *DescribeSampleDataByBatchUUidPageResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) SetGmtModified(v int64) *DescribeSampleDataByBatchUUidPageResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) SetLastSourceType(v string) *DescribeSampleDataByBatchUUidPageResponseBodyResultObject {
	s.LastSourceType = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) SetUuid(v string) *DescribeSampleDataByBatchUUidPageResponseBodyResultObject {
	s.Uuid = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) SetVersion(v int32) *DescribeSampleDataByBatchUUidPageResponseBodyResultObject {
	s.Version = &v
	return s
}

func (s *DescribeSampleDataByBatchUUidPageResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
