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
	// The queried suggestions.
	Items []*DescribeAvailableAdvicesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page. The value must be an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Default value: 30. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 96A55627-28E9-5E47-B8F6-D786BE551349
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the table in the DatabaseName.TableName format.
	SchemaTableNames []*string `json:"SchemaTableNames,omitempty" xml:"SchemaTableNames,omitempty" type:"Repeated"`
	// The total number of entries returned. The value must be an integer that is greater than or equal to 0. Default value: 0.
	//
	// example:
	//
	// 30
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
	// The time when the suggestion was generated. The time follows the ISO 8601 standard in the yyyyMMdd format. The time is displayed in UTC.
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
	// 	- **Index**: index optimization.
	//
	// 	- **Tiering**: hot and cold data optimization.
	//
	// example:
	//
	// Index
	AdviceType *string `json:"AdviceType,omitempty" xml:"AdviceType,omitempty"`
	// The benefit of the suggestion.
	//
	// example:
	//
	// 0.4 GB of storage saved
	Benefit     *string `json:"Benefit,omitempty" xml:"Benefit,omitempty"`
	IndexFields *string `json:"IndexFields,omitempty" xml:"IndexFields,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The reason why the suggestion was generated.
	//
	// example:
	//
	// Unused for 15 days, historical usage less than 1%
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The SQL statement that is used to apply the suggestion.
	//
	// example:
	//
	// alter table `schema1`.`table1` drop key col1_1_idx
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	// The name of the schema.
	//
	// example:
	//
	// adb_demo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The total number of entries returned. Minimum value: 0. Default value: 0.
	//
	// example:
	//
	// 30
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
