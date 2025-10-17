// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableAccessCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeTableAccessCountResponseBodyItems) *DescribeTableAccessCountResponseBody
	GetItems() []*DescribeTableAccessCountResponseBodyItems
	SetPageNumber(v int32) *DescribeTableAccessCountResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTableAccessCountResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeTableAccessCountResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeTableAccessCountResponseBody
	GetTotalCount() *int32
}

type DescribeTableAccessCountResponseBody struct {
	// The queried tables.
	Items []*DescribeTableAccessCountResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6B7D627B-DA23-572D-AD71-256F64698B7D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTableAccessCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableAccessCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableAccessCountResponseBody) GetItems() []*DescribeTableAccessCountResponseBodyItems {
	return s.Items
}

func (s *DescribeTableAccessCountResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTableAccessCountResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTableAccessCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTableAccessCountResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTableAccessCountResponseBody) SetItems(v []*DescribeTableAccessCountResponseBodyItems) *DescribeTableAccessCountResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTableAccessCountResponseBody) SetPageNumber(v int32) *DescribeTableAccessCountResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableAccessCountResponseBody) SetPageSize(v int32) *DescribeTableAccessCountResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTableAccessCountResponseBody) SetRequestId(v string) *DescribeTableAccessCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableAccessCountResponseBody) SetTotalCount(v int32) *DescribeTableAccessCountResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTableAccessCountResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTableAccessCountResponseBodyItems struct {
	// The number of accesses to the table.
	//
	// example:
	//
	// 6
	AccessCount *string `json:"AccessCount,omitempty" xml:"AccessCount,omitempty"`
	// The ID of the cluster to which the table belongs.
	//
	// example:
	//
	// amv-2ze627uzpkh8a8****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The date when the table was accessed.
	//
	// example:
	//
	// 2022-09-26
	ReportDate *string `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// CUSTOMER
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The database to which the table belongs.
	//
	// example:
	//
	// tpch
	TableSchema *string `json:"TableSchema,omitempty" xml:"TableSchema,omitempty"`
}

func (s DescribeTableAccessCountResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableAccessCountResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTableAccessCountResponseBodyItems) GetAccessCount() *string {
	return s.AccessCount
}

func (s *DescribeTableAccessCountResponseBodyItems) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeTableAccessCountResponseBodyItems) GetReportDate() *string {
	return s.ReportDate
}

func (s *DescribeTableAccessCountResponseBodyItems) GetTableName() *string {
	return s.TableName
}

func (s *DescribeTableAccessCountResponseBodyItems) GetTableSchema() *string {
	return s.TableSchema
}

func (s *DescribeTableAccessCountResponseBodyItems) SetAccessCount(v string) *DescribeTableAccessCountResponseBodyItems {
	s.AccessCount = &v
	return s
}

func (s *DescribeTableAccessCountResponseBodyItems) SetInstanceName(v string) *DescribeTableAccessCountResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeTableAccessCountResponseBodyItems) SetReportDate(v string) *DescribeTableAccessCountResponseBodyItems {
	s.ReportDate = &v
	return s
}

func (s *DescribeTableAccessCountResponseBodyItems) SetTableName(v string) *DescribeTableAccessCountResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeTableAccessCountResponseBodyItems) SetTableSchema(v string) *DescribeTableAccessCountResponseBodyItems {
	s.TableSchema = &v
	return s
}

func (s *DescribeTableAccessCountResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
