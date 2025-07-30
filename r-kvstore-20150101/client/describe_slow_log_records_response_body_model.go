// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *DescribeSlowLogRecordsResponseBody
	GetEngine() *string
	SetInstanceId(v string) *DescribeSlowLogRecordsResponseBody
	GetInstanceId() *string
	SetItems(v *DescribeSlowLogRecordsResponseBodyItems) *DescribeSlowLogRecordsResponseBody
	GetItems() *DescribeSlowLogRecordsResponseBodyItems
	SetPageNumber(v int32) *DescribeSlowLogRecordsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeSlowLogRecordsResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int32) *DescribeSlowLogRecordsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSlowLogRecordsResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeSlowLogRecordsResponseBody
	GetStartTime() *string
	SetTotalRecordCount(v int32) *DescribeSlowLogRecordsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeSlowLogRecordsResponseBody struct {
	// The database engine that the instance runs.
	//
	// example:
	//
	// Redis
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// r-bp10n********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The slow query log entries.
	Items *DescribeSlowLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of log entries returned on the current page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The maximum number of log entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 686BB8A6-BBA5-47E5-8A75-D2ADE433****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query.
	//
	// example:
	//
	// 2019-03-10T13:11Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of returned log entries.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeSlowLogRecordsResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSlowLogRecordsResponseBody) GetItems() *DescribeSlowLogRecordsResponseBodyItems {
	return s.Items
}

func (s *DescribeSlowLogRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSlowLogRecordsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeSlowLogRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSlowLogRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlowLogRecordsResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSlowLogRecordsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeSlowLogRecordsResponseBody) SetEngine(v string) *DescribeSlowLogRecordsResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetInstanceId(v string) *DescribeSlowLogRecordsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetItems(v *DescribeSlowLogRecordsResponseBodyItems) *DescribeSlowLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageNumber(v int32) *DescribeSlowLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageRecordCount(v int32) *DescribeSlowLogRecordsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageSize(v int32) *DescribeSlowLogRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetStartTime(v string) *DescribeSlowLogRecordsResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetTotalRecordCount(v int32) *DescribeSlowLogRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogRecordsResponseBodyItems struct {
	LogRecords []*DescribeSlowLogRecordsResponseBodyItemsLogRecords `json:"LogRecords,omitempty" xml:"LogRecords,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogRecordsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetLogRecords() []*DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	return s.LogRecords
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetLogRecords(v []*DescribeSlowLogRecordsResponseBodyItemsLogRecords) *DescribeSlowLogRecordsResponseBodyItems {
	s.LogRecords = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogRecordsResponseBodyItemsLogRecords struct {
	// The ID of the account.
	//
	// example:
	//
	// 0
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// demo
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The slow query statement.
	//
	// example:
	//
	// KEYS *
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The database name.
	//
	// example:
	//
	// -1
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The database name. This parameter serves the same purpose as the **DBName*	- parameter. We recommend that you use the **DBName*	- parameter.
	//
	// example:
	//
	// -1
	DataBaseName *string `json:"DataBaseName,omitempty" xml:"DataBaseName,omitempty"`
	// The amount of time consumed to execute the slow query statement. Unit: microseconds.
	//
	// example:
	//
	// 248
	ElapsedTime *int64 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	// The start time when the slow query statement was executed. The time is displayed in the YYYY-MM-DDTHH:mm:ssZ format.
	//
	// example:
	//
	// 2019-03-20T09:18:41Z
	ExecuteTime *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 172.16.88.***
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The node ID.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyItemsLogRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItemsLogRecords) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetAccount() *string {
	return s.Account
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetCommand() *string {
	return s.Command
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetDataBaseName() *string {
	return s.DataBaseName
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetElapsedTime() *int64 {
	return s.ElapsedTime
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetExecuteTime() *string {
	return s.ExecuteTime
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetAccount(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.Account = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetAccountName(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.AccountName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetCommand(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.Command = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetDBName(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetDataBaseName(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.DataBaseName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetElapsedTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetExecuteTime(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetIPAddress(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.IPAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetNodeId(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.NodeId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) Validate() error {
	return dara.Validate(s)
}
