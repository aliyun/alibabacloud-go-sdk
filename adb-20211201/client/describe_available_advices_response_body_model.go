// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableAdvicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeAvailableAdvicesResponseBodyItems) *DescribeAvailableAdvicesResponseBody
	GetItems() []*DescribeAvailableAdvicesResponseBodyItems
	SetPageNumber(v int64) *DescribeAvailableAdvicesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeAvailableAdvicesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeAvailableAdvicesResponseBody
	GetRequestId() *string
	SetSchemaTableNames(v []*string) *DescribeAvailableAdvicesResponseBody
	GetSchemaTableNames() []*string
	SetTotalCount(v int64) *DescribeAvailableAdvicesResponseBody
	GetTotalCount() *int64
}

type DescribeAvailableAdvicesResponseBody struct {
	// The details.
	Items []*DescribeAvailableAdvicesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number. Valid values: greater than 0 and up to the maximum value of the Integer data type. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **30*	- (default)
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 96A55627-28E9-5E47-B8F6-D786BE551349
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The concatenated string of the database name and table name.
	SchemaTableNames []*string `json:"SchemaTableNames,omitempty" xml:"SchemaTableNames,omitempty" type:"Repeated"`
	// The total number of returned results. Valid values: greater than or equal to 0 and up to the maximum value of the Integer data type. Default value: 0.
	//
	// example:
	//
	// 10192
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAvailableAdvicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableAdvicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableAdvicesResponseBody) GetItems() []*DescribeAvailableAdvicesResponseBodyItems {
	return s.Items
}

func (s *DescribeAvailableAdvicesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAvailableAdvicesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAvailableAdvicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableAdvicesResponseBody) GetSchemaTableNames() []*string {
	return s.SchemaTableNames
}

func (s *DescribeAvailableAdvicesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAvailableAdvicesResponseBody) SetItems(v []*DescribeAvailableAdvicesResponseBodyItems) *DescribeAvailableAdvicesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAvailableAdvicesResponseBody) SetPageNumber(v int64) *DescribeAvailableAdvicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBody) SetPageSize(v int64) *DescribeAvailableAdvicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBody) SetRequestId(v string) *DescribeAvailableAdvicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBody) SetSchemaTableNames(v []*string) *DescribeAvailableAdvicesResponseBody {
	s.SchemaTableNames = v
	return s
}

func (s *DescribeAvailableAdvicesResponseBody) SetTotalCount(v int64) *DescribeAvailableAdvicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBody) Validate() error {
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

type DescribeAvailableAdvicesResponseBodyItems struct {
	// The time when the suggestion was generated, in the yyyyMMdd format.
	//
	// example:
	//
	// 20221124
	AdviceDate *string `json:"AdviceDate,omitempty" xml:"AdviceDate,omitempty"`
	// The suggestion ID.
	//
	// example:
	//
	// dcd04135-0925-4aed-a5a7-e7d92cb1****
	AdviceId *string `json:"AdviceId,omitempty" xml:"AdviceId,omitempty"`
	// The type of the suggestion. Valid values:
	//
	// - **INDEX**: Index optimization.
	//
	// - **TIERING**: Hot and cold data optimization.
	//
	// example:
	//
	// INDEX
	AdviceType *string `json:"AdviceType,omitempty" xml:"AdviceType,omitempty"`
	// The benefit of the suggestion.
	//
	// example:
	//
	// Saves 0.4 GB of storage space
	Benefit *string `json:"Benefit,omitempty" xml:"Benefit,omitempty"`
	// The index fields.
	//
	// example:
	//
	// message
	IndexFields *string `json:"IndexFields,omitempty" xml:"IndexFields,omitempty"`
	// The page number. Valid values: greater than 0 and up to the maximum value of the Integer data type. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **30*	- (default)
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The reason for generating optimization suggestions.
	//
	// example:
	//
	// Not used for 15 days, and the historical usage probability is less than 1%
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The SQL statement to apply the suggestion.
	//
	// example:
	//
	// alter table `schema1`.`table1` drop key col1_1_idx
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	// The database name.
	//
	// example:
	//
	// adb_demo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The table name.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The total number of returned results. Valid values: greater than or equal to 0 and up to the maximum value of the Integer data type. Default value: 0.
	//
	// example:
	//
	// 10192
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAvailableAdvicesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableAdvicesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAvailableAdvicesResponseBodyItems) GetAdviceDate() *string {
	return s.AdviceDate
}

func (s *DescribeAvailableAdvicesResponseBodyItems) GetAdviceId() *string {
	return s.AdviceId
}

func (s *DescribeAvailableAdvicesResponseBodyItems) GetAdviceType() *string {
	return s.AdviceType
}

func (s *DescribeAvailableAdvicesResponseBodyItems) GetBenefit() *string {
	return s.Benefit
}

func (s *DescribeAvailableAdvicesResponseBodyItems) GetIndexFields() *string {
	return s.IndexFields
}

func (s *DescribeAvailableAdvicesResponseBodyItems) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAvailableAdvicesResponseBodyItems) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAvailableAdvicesResponseBodyItems) GetReason() *string {
	return s.Reason
}

func (s *DescribeAvailableAdvicesResponseBodyItems) GetSQL() *string {
	return s.SQL
}

func (s *DescribeAvailableAdvicesResponseBodyItems) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeAvailableAdvicesResponseBodyItems) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAvailableAdvicesResponseBodyItems) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetAdviceDate(v string) *DescribeAvailableAdvicesResponseBodyItems {
	s.AdviceDate = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetAdviceId(v string) *DescribeAvailableAdvicesResponseBodyItems {
	s.AdviceId = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetAdviceType(v string) *DescribeAvailableAdvicesResponseBodyItems {
	s.AdviceType = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetBenefit(v string) *DescribeAvailableAdvicesResponseBodyItems {
	s.Benefit = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetIndexFields(v string) *DescribeAvailableAdvicesResponseBodyItems {
	s.IndexFields = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetPageNumber(v int64) *DescribeAvailableAdvicesResponseBodyItems {
	s.PageNumber = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetPageSize(v int64) *DescribeAvailableAdvicesResponseBodyItems {
	s.PageSize = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetReason(v string) *DescribeAvailableAdvicesResponseBodyItems {
	s.Reason = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetSQL(v string) *DescribeAvailableAdvicesResponseBodyItems {
	s.SQL = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetSchemaName(v string) *DescribeAvailableAdvicesResponseBodyItems {
	s.SchemaName = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetTableName(v string) *DescribeAvailableAdvicesResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetTotalCount(v int64) *DescribeAvailableAdvicesResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
