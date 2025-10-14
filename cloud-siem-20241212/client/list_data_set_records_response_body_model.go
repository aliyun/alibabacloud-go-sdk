// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSetRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSetRecords(v []*ListDataSetRecordsResponseBodyDataSetRecords) *ListDataSetRecordsResponseBody
	GetDataSetRecords() []*ListDataSetRecordsResponseBodyDataSetRecords
	SetMaxResults(v int32) *ListDataSetRecordsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataSetRecordsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListDataSetRecordsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataSetRecordsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataSetRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDataSetRecordsResponseBody
	GetTotalCount() *int32
}

type ListDataSetRecordsResponseBody struct {
	DataSetRecords []*ListDataSetRecordsResponseBodyDataSetRecords `json:"DataSetRecords,omitempty" xml:"DataSetRecords,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 57
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataSetRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSetRecordsResponseBody) GetDataSetRecords() []*ListDataSetRecordsResponseBodyDataSetRecords {
	return s.DataSetRecords
}

func (s *ListDataSetRecordsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataSetRecordsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataSetRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataSetRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataSetRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSetRecordsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataSetRecordsResponseBody) SetDataSetRecords(v []*ListDataSetRecordsResponseBodyDataSetRecords) *ListDataSetRecordsResponseBody {
	s.DataSetRecords = v
	return s
}

func (s *ListDataSetRecordsResponseBody) SetMaxResults(v int32) *ListDataSetRecordsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataSetRecordsResponseBody) SetNextToken(v string) *ListDataSetRecordsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataSetRecordsResponseBody) SetPageNumber(v int32) *ListDataSetRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDataSetRecordsResponseBody) SetPageSize(v int32) *ListDataSetRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataSetRecordsResponseBody) SetRequestId(v string) *ListDataSetRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSetRecordsResponseBody) SetTotalCount(v int32) *ListDataSetRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataSetRecordsResponseBody) Validate() error {
	if s.DataSetRecords != nil {
		for _, item := range s.DataSetRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSetRecordsResponseBodyDataSetRecords struct {
	// example:
	//
	// 1658974643000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// dataset-t8ha6p7k61rmniqw****
	DataSetId *string `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
	// example:
	//
	// lmftest
	DataSetName *string `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	// example:
	//
	// 124566
	DataSetRecordId     *string `json:"DataSetRecordId,omitempty" xml:"DataSetRecordId,omitempty"`
	DataSetRecordValues *string `json:"DataSetRecordValues,omitempty" xml:"DataSetRecordValues,omitempty"`
	// example:
	//
	// 1658974643000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListDataSetRecordsResponseBodyDataSetRecords) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetRecordsResponseBodyDataSetRecords) GoString() string {
	return s.String()
}

func (s *ListDataSetRecordsResponseBodyDataSetRecords) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataSetRecordsResponseBodyDataSetRecords) GetDataSetId() *string {
	return s.DataSetId
}

func (s *ListDataSetRecordsResponseBodyDataSetRecords) GetDataSetName() *string {
	return s.DataSetName
}

func (s *ListDataSetRecordsResponseBodyDataSetRecords) GetDataSetRecordId() *string {
	return s.DataSetRecordId
}

func (s *ListDataSetRecordsResponseBodyDataSetRecords) GetDataSetRecordValues() *string {
	return s.DataSetRecordValues
}

func (s *ListDataSetRecordsResponseBodyDataSetRecords) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListDataSetRecordsResponseBodyDataSetRecords) SetCreateTime(v int64) *ListDataSetRecordsResponseBodyDataSetRecords {
	s.CreateTime = &v
	return s
}

func (s *ListDataSetRecordsResponseBodyDataSetRecords) SetDataSetId(v string) *ListDataSetRecordsResponseBodyDataSetRecords {
	s.DataSetId = &v
	return s
}

func (s *ListDataSetRecordsResponseBodyDataSetRecords) SetDataSetName(v string) *ListDataSetRecordsResponseBodyDataSetRecords {
	s.DataSetName = &v
	return s
}

func (s *ListDataSetRecordsResponseBodyDataSetRecords) SetDataSetRecordId(v string) *ListDataSetRecordsResponseBodyDataSetRecords {
	s.DataSetRecordId = &v
	return s
}

func (s *ListDataSetRecordsResponseBodyDataSetRecords) SetDataSetRecordValues(v string) *ListDataSetRecordsResponseBodyDataSetRecords {
	s.DataSetRecordValues = &v
	return s
}

func (s *ListDataSetRecordsResponseBodyDataSetRecords) SetUpdateTime(v int64) *ListDataSetRecordsResponseBodyDataSetRecords {
	s.UpdateTime = &v
	return s
}

func (s *ListDataSetRecordsResponseBodyDataSetRecords) Validate() error {
	return dara.Validate(s)
}
