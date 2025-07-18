// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeSQLLogsResponseBodyItems) *DescribeSQLLogsResponseBody
	GetItems() []*DescribeSQLLogsResponseBodyItems
	SetPageNumber(v int32) *DescribeSQLLogsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeSQLLogsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeSQLLogsResponseBody
	GetRequestId() *string
}

type DescribeSQLLogsResponseBody struct {
	// The queried SQL execution logs.
	Items []*DescribeSQLLogsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A7941C94-B92F-46A0-BD3E-2D**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSQLLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsResponseBody) GetItems() []*DescribeSQLLogsResponseBodyItems {
	return s.Items
}

func (s *DescribeSQLLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSQLLogsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeSQLLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSQLLogsResponseBody) SetItems(v []*DescribeSQLLogsResponseBodyItems) *DescribeSQLLogsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSQLLogsResponseBody) SetPageNumber(v int32) *DescribeSQLLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsResponseBody) SetPageRecordCount(v int32) *DescribeSQLLogsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogsResponseBody) SetRequestId(v string) *DescribeSQLLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSQLLogsResponseBodyItems struct {
	// The database account that executes the SQL statement.
	//
	// example:
	//
	// testadmin
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adbpgadmin
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The role of the database.
	//
	// example:
	//
	// master
	DBRole *string `json:"DBRole,omitempty" xml:"DBRole,omitempty"`
	// The execution duration of the SQL statement.
	//
	// example:
	//
	// 2
	ExecuteCost *float32 `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	// The execution status of the SQL statement. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	ExecuteState *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	// The type of the query language.
	//
	// example:
	//
	// DQL
	OperationClass *string `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	// The time when the SQL statement was executed.
	//
	// example:
	//
	// 2021-03-15T17:02:32Z
	OperationExecuteTime *string `json:"OperationExecuteTime,omitempty" xml:"OperationExecuteTime,omitempty"`
	// The type of the SQL statement.
	//
	// example:
	//
	// SELECT
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	ReturnRowCounts *int64 `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	// The SQL execution plan.
	//
	// example:
	//
	// ""
	SQLPlan *string `json:"SQLPlan,omitempty" xml:"SQLPlan,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// select 1
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// The number of entries scanned.
	//
	// example:
	//
	// 1
	ScanRowCounts *int64 `json:"ScanRowCounts,omitempty" xml:"ScanRowCounts,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 100.**.**.90
	SourceIP *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	// The number of the source port.
	//
	// example:
	//
	// 50514
	SourcePort *int32 `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
}

func (s DescribeSQLLogsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsResponseBodyItems) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeSQLLogsResponseBodyItems) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSQLLogsResponseBodyItems) GetDBRole() *string {
	return s.DBRole
}

func (s *DescribeSQLLogsResponseBodyItems) GetExecuteCost() *float32 {
	return s.ExecuteCost
}

func (s *DescribeSQLLogsResponseBodyItems) GetExecuteState() *string {
	return s.ExecuteState
}

func (s *DescribeSQLLogsResponseBodyItems) GetOperationClass() *string {
	return s.OperationClass
}

func (s *DescribeSQLLogsResponseBodyItems) GetOperationExecuteTime() *string {
	return s.OperationExecuteTime
}

func (s *DescribeSQLLogsResponseBodyItems) GetOperationType() *string {
	return s.OperationType
}

func (s *DescribeSQLLogsResponseBodyItems) GetReturnRowCounts() *int64 {
	return s.ReturnRowCounts
}

func (s *DescribeSQLLogsResponseBodyItems) GetSQLPlan() *string {
	return s.SQLPlan
}

func (s *DescribeSQLLogsResponseBodyItems) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeSQLLogsResponseBodyItems) GetScanRowCounts() *int64 {
	return s.ScanRowCounts
}

func (s *DescribeSQLLogsResponseBodyItems) GetSourceIP() *string {
	return s.SourceIP
}

func (s *DescribeSQLLogsResponseBodyItems) GetSourcePort() *int32 {
	return s.SourcePort
}

func (s *DescribeSQLLogsResponseBodyItems) SetAccountName(v string) *DescribeSQLLogsResponseBodyItems {
	s.AccountName = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetDBName(v string) *DescribeSQLLogsResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetDBRole(v string) *DescribeSQLLogsResponseBodyItems {
	s.DBRole = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetExecuteCost(v float32) *DescribeSQLLogsResponseBodyItems {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetExecuteState(v string) *DescribeSQLLogsResponseBodyItems {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetOperationClass(v string) *DescribeSQLLogsResponseBodyItems {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetOperationExecuteTime(v string) *DescribeSQLLogsResponseBodyItems {
	s.OperationExecuteTime = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetOperationType(v string) *DescribeSQLLogsResponseBodyItems {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetReturnRowCounts(v int64) *DescribeSQLLogsResponseBodyItems {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetSQLPlan(v string) *DescribeSQLLogsResponseBodyItems {
	s.SQLPlan = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetSQLText(v string) *DescribeSQLLogsResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetScanRowCounts(v int64) *DescribeSQLLogsResponseBodyItems {
	s.ScanRowCounts = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetSourceIP(v string) *DescribeSQLLogsResponseBodyItems {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) SetSourcePort(v int32) *DescribeSQLLogsResponseBodyItems {
	s.SourcePort = &v
	return s
}

func (s *DescribeSQLLogsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
