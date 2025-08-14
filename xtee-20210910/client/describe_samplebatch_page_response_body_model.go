// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSamplebatchPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSamplebatchPageResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeSamplebatchPageResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeSamplebatchPageResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeSamplebatchPageResponseBodyResultObject) *DescribeSamplebatchPageResponseBody
	GetResultObject() []*DescribeSamplebatchPageResponseBodyResultObject
	SetTotalItem(v int32) *DescribeSamplebatchPageResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeSamplebatchPageResponseBody
	GetTotalPage() *int32
}

type DescribeSamplebatchPageResponseBody struct {
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
	// Page size, with a default value of 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object
	ResultObject []*DescribeSamplebatchPageResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items
	//
	// example:
	//
	// 3
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeSamplebatchPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSamplebatchPageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSamplebatchPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSamplebatchPageResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSamplebatchPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSamplebatchPageResponseBody) GetResultObject() []*DescribeSamplebatchPageResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSamplebatchPageResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeSamplebatchPageResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeSamplebatchPageResponseBody) SetRequestId(v string) *DescribeSamplebatchPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSamplebatchPageResponseBody) SetCurrentPage(v int32) *DescribeSamplebatchPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSamplebatchPageResponseBody) SetPageSize(v int32) *DescribeSamplebatchPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSamplebatchPageResponseBody) SetResultObject(v []*DescribeSamplebatchPageResponseBodyResultObject) *DescribeSamplebatchPageResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSamplebatchPageResponseBody) SetTotalItem(v int32) *DescribeSamplebatchPageResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeSamplebatchPageResponseBody) SetTotalPage(v int32) *DescribeSamplebatchPageResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeSamplebatchPageResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSamplebatchPageResponseBodyResultObject struct {
	// Sample batch name
	//
	// example:
	//
	// 白样本
	BatchName *string `json:"batchName,omitempty" xml:"batchName,omitempty"`
	// Creator.
	//
	// example:
	//
	// 1519714049632764
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// Data type
	//
	// example:
	//
	// mobile
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// Description.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
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
	// Valid sample content data
	//
	// example:
	//
	// 2
	InitValidFileRow *int64 `json:"initValidFileRow,omitempty" xml:"initValidFileRow,omitempty"`
	// Specific type of the sample list
	//
	// example:
	//
	// pass
	SampleBatchType *string `json:"sampleBatchType,omitempty" xml:"sampleBatchType,omitempty"`
	// Service ID
	//
	// example:
	//
	// account_abuse
	Services *string `json:"services,omitempty" xml:"services,omitempty"`
	// Modifier
	//
	// example:
	//
	// 1519714049632764
	Updator *string `json:"updator,omitempty" xml:"updator,omitempty"`
	// Sample batch UUID
	//
	// example:
	//
	// 203f1ae65c0a41a49dc4f8a47946caa2
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s DescribeSamplebatchPageResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSamplebatchPageResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) GetBatchName() *string {
	return s.BatchName
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) GetCreator() *string {
	return s.Creator
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) GetDataType() *string {
	return s.DataType
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) GetInitValidFileRow() *int64 {
	return s.InitValidFileRow
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) GetSampleBatchType() *string {
	return s.SampleBatchType
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) GetServices() *string {
	return s.Services
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) GetUpdator() *string {
	return s.Updator
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) SetBatchName(v string) *DescribeSamplebatchPageResponseBodyResultObject {
	s.BatchName = &v
	return s
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) SetCreator(v string) *DescribeSamplebatchPageResponseBodyResultObject {
	s.Creator = &v
	return s
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) SetDataType(v string) *DescribeSamplebatchPageResponseBodyResultObject {
	s.DataType = &v
	return s
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) SetDescription(v string) *DescribeSamplebatchPageResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) SetGmtCreate(v int64) *DescribeSamplebatchPageResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) SetGmtModified(v int64) *DescribeSamplebatchPageResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) SetInitValidFileRow(v int64) *DescribeSamplebatchPageResponseBodyResultObject {
	s.InitValidFileRow = &v
	return s
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) SetSampleBatchType(v string) *DescribeSamplebatchPageResponseBodyResultObject {
	s.SampleBatchType = &v
	return s
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) SetServices(v string) *DescribeSamplebatchPageResponseBodyResultObject {
	s.Services = &v
	return s
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) SetUpdator(v string) *DescribeSamplebatchPageResponseBodyResultObject {
	s.Updator = &v
	return s
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) SetUuid(v string) *DescribeSamplebatchPageResponseBodyResultObject {
	s.Uuid = &v
	return s
}

func (s *DescribeSamplebatchPageResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
