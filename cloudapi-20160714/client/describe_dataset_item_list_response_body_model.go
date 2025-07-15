// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatasetItemListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetItemInfoList(v []*DescribeDatasetItemListResponseBodyDatasetItemInfoList) *DescribeDatasetItemListResponseBody
	GetDatasetItemInfoList() []*DescribeDatasetItemListResponseBodyDatasetItemInfoList
	SetPageNumber(v int32) *DescribeDatasetItemListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDatasetItemListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDatasetItemListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDatasetItemListResponseBody
	GetTotalCount() *int32
}

type DescribeDatasetItemListResponseBody struct {
	// The Dataset information.
	DatasetItemInfoList []*DescribeDatasetItemListResponseBodyDatasetItemInfoList `json:"DatasetItemInfoList,omitempty" xml:"DatasetItemInfoList,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C6E9C6E4-608B-5C0F-9783-E288********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDatasetItemListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetItemListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatasetItemListResponseBody) GetDatasetItemInfoList() []*DescribeDatasetItemListResponseBodyDatasetItemInfoList {
	return s.DatasetItemInfoList
}

func (s *DescribeDatasetItemListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDatasetItemListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDatasetItemListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDatasetItemListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDatasetItemListResponseBody) SetDatasetItemInfoList(v []*DescribeDatasetItemListResponseBodyDatasetItemInfoList) *DescribeDatasetItemListResponseBody {
	s.DatasetItemInfoList = v
	return s
}

func (s *DescribeDatasetItemListResponseBody) SetPageNumber(v int32) *DescribeDatasetItemListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatasetItemListResponseBody) SetPageSize(v int32) *DescribeDatasetItemListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDatasetItemListResponseBody) SetRequestId(v string) *DescribeDatasetItemListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatasetItemListResponseBody) SetTotalCount(v int32) *DescribeDatasetItemListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDatasetItemListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDatasetItemListResponseBodyDatasetItemInfoList struct {
	// The time when the data entry was created.
	//
	// example:
	//
	// 2022-09-21T12:58:43Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The ID of the dataset.
	//
	// example:
	//
	// 4add6a61804e47858266883e********
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The ID of the data entry.
	//
	// example:
	//
	// 5045****
	DatasetItemId *string `json:"DatasetItemId,omitempty" xml:"DatasetItemId,omitempty"`
	// The description of the data entry.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time in UTC when the data entry expires. The time is in the **yyyy-MM-ddTHH:mm:ssZ*	- format. If this parameter is empty, the data entry does not expire.
	//
	// example:
	//
	// 2022-09-22T12:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The last modification time (UTC) of the data entry.
	//
	// example:
	//
	// 2022-09-21T12:58:43Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The value of the data entry.
	//
	// example:
	//
	// 106.43.XXX.XXX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDatasetItemListResponseBodyDatasetItemInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetItemListResponseBodyDatasetItemInfoList) GoString() string {
	return s.String()
}

func (s *DescribeDatasetItemListResponseBodyDatasetItemInfoList) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeDatasetItemListResponseBodyDatasetItemInfoList) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DescribeDatasetItemListResponseBodyDatasetItemInfoList) GetDatasetItemId() *string {
	return s.DatasetItemId
}

func (s *DescribeDatasetItemListResponseBodyDatasetItemInfoList) GetDescription() *string {
	return s.Description
}

func (s *DescribeDatasetItemListResponseBodyDatasetItemInfoList) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeDatasetItemListResponseBodyDatasetItemInfoList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeDatasetItemListResponseBodyDatasetItemInfoList) GetValue() *string {
	return s.Value
}

func (s *DescribeDatasetItemListResponseBodyDatasetItemInfoList) SetCreatedTime(v string) *DescribeDatasetItemListResponseBodyDatasetItemInfoList {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDatasetItemListResponseBodyDatasetItemInfoList) SetDatasetId(v string) *DescribeDatasetItemListResponseBodyDatasetItemInfoList {
	s.DatasetId = &v
	return s
}

func (s *DescribeDatasetItemListResponseBodyDatasetItemInfoList) SetDatasetItemId(v string) *DescribeDatasetItemListResponseBodyDatasetItemInfoList {
	s.DatasetItemId = &v
	return s
}

func (s *DescribeDatasetItemListResponseBodyDatasetItemInfoList) SetDescription(v string) *DescribeDatasetItemListResponseBodyDatasetItemInfoList {
	s.Description = &v
	return s
}

func (s *DescribeDatasetItemListResponseBodyDatasetItemInfoList) SetExpiredTime(v string) *DescribeDatasetItemListResponseBodyDatasetItemInfoList {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDatasetItemListResponseBodyDatasetItemInfoList) SetModifiedTime(v string) *DescribeDatasetItemListResponseBodyDatasetItemInfoList {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeDatasetItemListResponseBodyDatasetItemInfoList) SetValue(v string) *DescribeDatasetItemListResponseBodyDatasetItemInfoList {
	s.Value = &v
	return s
}

func (s *DescribeDatasetItemListResponseBodyDatasetItemInfoList) Validate() error {
	return dara.Validate(s)
}
