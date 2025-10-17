// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSparkWarehouseBatchSQLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListSparkWarehouseBatchSQLResponseBodyData) *ListSparkWarehouseBatchSQLResponseBody
	GetData() *ListSparkWarehouseBatchSQLResponseBodyData
	SetRequestId(v string) *ListSparkWarehouseBatchSQLResponseBody
	GetRequestId() *string
}

type ListSparkWarehouseBatchSQLResponseBody struct {
	// The returned data.
	Data *ListSparkWarehouseBatchSQLResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSparkWarehouseBatchSQLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSparkWarehouseBatchSQLResponseBody) GoString() string {
	return s.String()
}

func (s *ListSparkWarehouseBatchSQLResponseBody) GetData() *ListSparkWarehouseBatchSQLResponseBodyData {
	return s.Data
}

func (s *ListSparkWarehouseBatchSQLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSparkWarehouseBatchSQLResponseBody) SetData(v *ListSparkWarehouseBatchSQLResponseBodyData) *ListSparkWarehouseBatchSQLResponseBody {
	s.Data = v
	return s
}

func (s *ListSparkWarehouseBatchSQLResponseBody) SetRequestId(v string) *ListSparkWarehouseBatchSQLResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSparkWarehouseBatchSQLResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSparkWarehouseBatchSQLResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The queried Spark SQL statements.
	Queries []*SparkBatchSQL `json:"Queries,omitempty" xml:"Queries,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSparkWarehouseBatchSQLResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSparkWarehouseBatchSQLResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSparkWarehouseBatchSQLResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSparkWarehouseBatchSQLResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSparkWarehouseBatchSQLResponseBodyData) GetQueries() []*SparkBatchSQL {
	return s.Queries
}

func (s *ListSparkWarehouseBatchSQLResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListSparkWarehouseBatchSQLResponseBodyData) SetPageNumber(v int64) *ListSparkWarehouseBatchSQLResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSparkWarehouseBatchSQLResponseBodyData) SetPageSize(v int64) *ListSparkWarehouseBatchSQLResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSparkWarehouseBatchSQLResponseBodyData) SetQueries(v []*SparkBatchSQL) *ListSparkWarehouseBatchSQLResponseBodyData {
	s.Queries = v
	return s
}

func (s *ListSparkWarehouseBatchSQLResponseBodyData) SetTotal(v int64) *ListSparkWarehouseBatchSQLResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListSparkWarehouseBatchSQLResponseBodyData) Validate() error {
	if s.Queries != nil {
		for _, item := range s.Queries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
