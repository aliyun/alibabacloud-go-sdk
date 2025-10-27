// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppliedAdvicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeAppliedAdvicesResponseBodyItems) *DescribeAppliedAdvicesResponseBody
	GetItems() []*DescribeAppliedAdvicesResponseBodyItems
	SetPageNumber(v int64) *DescribeAppliedAdvicesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeAppliedAdvicesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeAppliedAdvicesResponseBody
	GetRequestId() *string
	SetSchemaTableNames(v []*string) *DescribeAppliedAdvicesResponseBody
	GetSchemaTableNames() []*string
	SetTotalCount(v int64) *DescribeAppliedAdvicesResponseBody
	GetTotalCount() *int64
}

type DescribeAppliedAdvicesResponseBody struct {
	// The queried suggestions.
	Items []*DescribeAppliedAdvicesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number. Pages start from page 1. Default value: 1.
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
	// The request ID.
	//
	// example:
	//
	// 84489769-3065-5A28-A4CB-977CD426F1C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The names of the tables in the DatabaseName.TableName format.
	SchemaTableNames []*string `json:"SchemaTableNames,omitempty" xml:"SchemaTableNames,omitempty" type:"Repeated"`
	// The total number of entries returned. The value is an integer that is greater than or equal to 0. Default value: 0.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAppliedAdvicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppliedAdvicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppliedAdvicesResponseBody) GetItems() []*DescribeAppliedAdvicesResponseBodyItems {
	return s.Items
}

func (s *DescribeAppliedAdvicesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAppliedAdvicesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAppliedAdvicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppliedAdvicesResponseBody) GetSchemaTableNames() []*string {
	return s.SchemaTableNames
}

func (s *DescribeAppliedAdvicesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAppliedAdvicesResponseBody) SetItems(v []*DescribeAppliedAdvicesResponseBodyItems) *DescribeAppliedAdvicesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAppliedAdvicesResponseBody) SetPageNumber(v int64) *DescribeAppliedAdvicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBody) SetPageSize(v int64) *DescribeAppliedAdvicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBody) SetRequestId(v string) *DescribeAppliedAdvicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBody) SetSchemaTableNames(v []*string) *DescribeAppliedAdvicesResponseBody {
	s.SchemaTableNames = v
	return s
}

func (s *DescribeAppliedAdvicesResponseBody) SetTotalCount(v int64) *DescribeAppliedAdvicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBody) Validate() error {
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

type DescribeAppliedAdvicesResponseBodyItems struct {
	// The suggestion ID.
	//
	// example:
	//
	// 7417db9c-914d-43f3-a123-4d0e448f****
	AdviceId *string `json:"AdviceId,omitempty" xml:"AdviceId,omitempty"`
	// The benefit of the suggestion.
	//
	// example:
	//
	// 0.4 GB of storage saved
	Benefit *string `json:"Benefit,omitempty" xml:"Benefit,omitempty"`
	// The SQL statement that is used to execute the BUILD job.
	//
	// example:
	//
	// build table `schema1`.`table1`
	BuildSQL    *string `json:"BuildSQL,omitempty" xml:"BuildSQL,omitempty"`
	IndexFields *string `json:"IndexFields,omitempty" xml:"IndexFields,omitempty"`
	// The state of the suggestion execution job. Valid values:
	//
	// 	- **SUCCEED**
	//
	// 	- **FAILED**
	//
	// example:
	//
	// SUCCEED
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
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
	// The SQL statement that is used to roll back the suggestion.
	//
	// example:
	//
	// alter table `schema1`.`table1` add key col1_1_idx(col1)
	RollbackSQL *string `json:"RollbackSQL,omitempty" xml:"RollbackSQL,omitempty"`
	// The SQL statement that is used to apply the suggestion.
	//
	// example:
	//
	// alter table `schema1`.`table1` drop key col1_1_idx
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adb_demo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The submission state of the suggestion. Valid values:
	//
	// 	- **SUCCEED**
	//
	// 	- **FAILED**
	//
	// example:
	//
	// SUCCEED
	SubmitStatus *string `json:"SubmitStatus,omitempty" xml:"SubmitStatus,omitempty"`
	// The time when the suggestion was submitted. The time follows the ISO 8601 standard in the yyMMddHHmm format. The time is displayed in UTC.
	//
	// example:
	//
	// 2208131600
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The total number of entries returned. Minimum value: 0. Default value: 0.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAppliedAdvicesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppliedAdvicesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAppliedAdvicesResponseBodyItems) GetAdviceId() *string {
	return s.AdviceId
}

func (s *DescribeAppliedAdvicesResponseBodyItems) GetBenefit() *string {
	return s.Benefit
}

func (s *DescribeAppliedAdvicesResponseBodyItems) GetBuildSQL() *string {
	return s.BuildSQL
}

func (s *DescribeAppliedAdvicesResponseBodyItems) GetIndexFields() *string {
	return s.IndexFields
}

func (s *DescribeAppliedAdvicesResponseBodyItems) GetJobStatus() *string {
	return s.JobStatus
}

func (s *DescribeAppliedAdvicesResponseBodyItems) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAppliedAdvicesResponseBodyItems) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAppliedAdvicesResponseBodyItems) GetRollbackSQL() *string {
	return s.RollbackSQL
}

func (s *DescribeAppliedAdvicesResponseBodyItems) GetSQL() *string {
	return s.SQL
}

func (s *DescribeAppliedAdvicesResponseBodyItems) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeAppliedAdvicesResponseBodyItems) GetSubmitStatus() *string {
	return s.SubmitStatus
}

func (s *DescribeAppliedAdvicesResponseBodyItems) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *DescribeAppliedAdvicesResponseBodyItems) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAppliedAdvicesResponseBodyItems) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetAdviceId(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.AdviceId = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetBenefit(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.Benefit = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetBuildSQL(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.BuildSQL = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetIndexFields(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.IndexFields = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetJobStatus(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.JobStatus = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetPageNumber(v int64) *DescribeAppliedAdvicesResponseBodyItems {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetPageSize(v int64) *DescribeAppliedAdvicesResponseBodyItems {
	s.PageSize = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetRollbackSQL(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.RollbackSQL = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetSQL(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.SQL = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetSchemaName(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.SchemaName = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetSubmitStatus(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.SubmitStatus = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetSubmitTime(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.SubmitTime = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetTableName(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetTotalCount(v int64) *DescribeAppliedAdvicesResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
