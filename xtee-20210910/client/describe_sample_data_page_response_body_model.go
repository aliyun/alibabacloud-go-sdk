// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDataPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSampleDataPageResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeSampleDataPageResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeSampleDataPageResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeSampleDataPageResponseBodyResultObject) *DescribeSampleDataPageResponseBody
	GetResultObject() []*DescribeSampleDataPageResponseBodyResultObject
	SetTotalItem(v int32) *DescribeSampleDataPageResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeSampleDataPageResponseBody
	GetTotalPage() *int32
}

type DescribeSampleDataPageResponseBody struct {
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
	// Number of items per page, default is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object
	ResultObject []*DescribeSampleDataPageResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items
	//
	// example:
	//
	// 3
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeSampleDataPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDataPageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSampleDataPageResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSampleDataPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSampleDataPageResponseBody) GetResultObject() []*DescribeSampleDataPageResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSampleDataPageResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeSampleDataPageResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeSampleDataPageResponseBody) SetRequestId(v string) *DescribeSampleDataPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSampleDataPageResponseBody) SetCurrentPage(v int32) *DescribeSampleDataPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSampleDataPageResponseBody) SetPageSize(v int32) *DescribeSampleDataPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSampleDataPageResponseBody) SetResultObject(v []*DescribeSampleDataPageResponseBodyResultObject) *DescribeSampleDataPageResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSampleDataPageResponseBody) SetTotalItem(v int32) *DescribeSampleDataPageResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeSampleDataPageResponseBody) SetTotalPage(v int32) *DescribeSampleDataPageResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeSampleDataPageResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSampleDataPageResponseBodyResultObject struct {
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
	// Content of the list entered in the text box
	//
	// example:
	//
	// 1770000000
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
	// Modifier
	//
	// example:
	//
	// 1519714049632764
	Updator *string `json:"updator,omitempty" xml:"updator,omitempty"`
	// UUID of the sample batch
	//
	// example:
	//
	// 48653f1372444c078f7b3d1c317d37dc
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// Version number
	//
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeSampleDataPageResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDataPageResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataPageResponseBodyResultObject) GetCreator() *string {
	return s.Creator
}

func (s *DescribeSampleDataPageResponseBodyResultObject) GetDataTagType() *string {
	return s.DataTagType
}

func (s *DescribeSampleDataPageResponseBodyResultObject) GetDataValue() *string {
	return s.DataValue
}

func (s *DescribeSampleDataPageResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeSampleDataPageResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeSampleDataPageResponseBodyResultObject) GetLastSourceType() *string {
	return s.LastSourceType
}

func (s *DescribeSampleDataPageResponseBodyResultObject) GetUpdator() *string {
	return s.Updator
}

func (s *DescribeSampleDataPageResponseBodyResultObject) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSampleDataPageResponseBodyResultObject) GetVersion() *int32 {
	return s.Version
}

func (s *DescribeSampleDataPageResponseBodyResultObject) SetCreator(v string) *DescribeSampleDataPageResponseBodyResultObject {
	s.Creator = &v
	return s
}

func (s *DescribeSampleDataPageResponseBodyResultObject) SetDataTagType(v string) *DescribeSampleDataPageResponseBodyResultObject {
	s.DataTagType = &v
	return s
}

func (s *DescribeSampleDataPageResponseBodyResultObject) SetDataValue(v string) *DescribeSampleDataPageResponseBodyResultObject {
	s.DataValue = &v
	return s
}

func (s *DescribeSampleDataPageResponseBodyResultObject) SetGmtCreate(v int64) *DescribeSampleDataPageResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSampleDataPageResponseBodyResultObject) SetGmtModified(v int64) *DescribeSampleDataPageResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeSampleDataPageResponseBodyResultObject) SetLastSourceType(v string) *DescribeSampleDataPageResponseBodyResultObject {
	s.LastSourceType = &v
	return s
}

func (s *DescribeSampleDataPageResponseBodyResultObject) SetUpdator(v string) *DescribeSampleDataPageResponseBodyResultObject {
	s.Updator = &v
	return s
}

func (s *DescribeSampleDataPageResponseBodyResultObject) SetUuid(v string) *DescribeSampleDataPageResponseBodyResultObject {
	s.Uuid = &v
	return s
}

func (s *DescribeSampleDataPageResponseBodyResultObject) SetVersion(v int32) *DescribeSampleDataPageResponseBodyResultObject {
	s.Version = &v
	return s
}

func (s *DescribeSampleDataPageResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
