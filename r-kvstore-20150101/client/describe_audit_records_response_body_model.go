// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeAuditRecordsResponseBody
	GetEndTime() *string
	SetInstanceName(v string) *DescribeAuditRecordsResponseBody
	GetInstanceName() *string
	SetItems(v *DescribeAuditRecordsResponseBodyItems) *DescribeAuditRecordsResponseBody
	GetItems() *DescribeAuditRecordsResponseBodyItems
	SetPageNumber(v int32) *DescribeAuditRecordsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAuditRecordsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAuditRecordsResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeAuditRecordsResponseBody
	GetStartTime() *string
	SetTotalRecordCount(v int32) *DescribeAuditRecordsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeAuditRecordsResponseBody struct {
	// The end time of the query.
	//
	// example:
	//
	// 2019-03-25T12:10:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The collection of returned audit log entries.
	Items *DescribeAuditRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9F5EB478-824E-4AC4-8D2B-58F31A02****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query.
	//
	// example:
	//
	// 2019-03-24T12:10:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 22222
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeAuditRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditRecordsResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAuditRecordsResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeAuditRecordsResponseBody) GetItems() *DescribeAuditRecordsResponseBodyItems {
	return s.Items
}

func (s *DescribeAuditRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAuditRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAuditRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAuditRecordsResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAuditRecordsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeAuditRecordsResponseBody) SetEndTime(v string) *DescribeAuditRecordsResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetInstanceName(v string) *DescribeAuditRecordsResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetItems(v *DescribeAuditRecordsResponseBodyItems) *DescribeAuditRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetPageNumber(v int32) *DescribeAuditRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetPageSize(v int32) *DescribeAuditRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetRequestId(v string) *DescribeAuditRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetStartTime(v string) *DescribeAuditRecordsResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetTotalRecordCount(v int32) *DescribeAuditRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAuditRecordsResponseBodyItems struct {
	SQL []*DescribeAuditRecordsResponseBodyItemsSQL `json:"SQL,omitempty" xml:"SQL,omitempty" type:"Repeated"`
}

func (s DescribeAuditRecordsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAuditRecordsResponseBodyItems) GetSQL() []*DescribeAuditRecordsResponseBodyItemsSQL {
	return s.SQL
}

func (s *DescribeAuditRecordsResponseBodyItems) SetSQL(v []*DescribeAuditRecordsResponseBodyItemsSQL) *DescribeAuditRecordsResponseBodyItems {
	s.SQL = v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItems) Validate() error {
	if s.SQL != nil {
		for _, item := range s.SQL {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAuditRecordsResponseBodyItemsSQL struct {
	// The username of the account.
	//
	// example:
	//
	// demo
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The database name.
	//
	// example:
	//
	// demo
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The time when the command was run.
	//
	// example:
	//
	// 2019-03-25T03:22:08Z
	ExecuteTime *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 127.0.0.1
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 192.16.100.***
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The ID of the node.
	//
	// > A specific node ID is returned only if the instance uses the cluster or read/write splitting architecture.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The command that was run.
	//
	// example:
	//
	// CONFIG GET maxmemory
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// The type of the command.
	//
	// example:
	//
	// non_read_write
	SQLType *string `json:"SQLType,omitempty" xml:"SQLType,omitempty"`
	// The amount of time consumed to run the command.
	//
	// example:
	//
	// 0
	TotalExecutionTimes *string `json:"TotalExecutionTimes,omitempty" xml:"TotalExecutionTimes,omitempty"`
}

func (s DescribeAuditRecordsResponseBodyItemsSQL) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditRecordsResponseBodyItemsSQL) GoString() string {
	return s.String()
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) GetExecuteTime() *string {
	return s.ExecuteTime
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) GetHostAddress() *string {
	return s.HostAddress
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) GetSQLType() *string {
	return s.SQLType
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) GetTotalExecutionTimes() *string {
	return s.TotalExecutionTimes
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) SetAccountName(v string) *DescribeAuditRecordsResponseBodyItemsSQL {
	s.AccountName = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) SetDatabaseName(v string) *DescribeAuditRecordsResponseBodyItemsSQL {
	s.DatabaseName = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) SetExecuteTime(v string) *DescribeAuditRecordsResponseBodyItemsSQL {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) SetHostAddress(v string) *DescribeAuditRecordsResponseBodyItemsSQL {
	s.HostAddress = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) SetIPAddress(v string) *DescribeAuditRecordsResponseBodyItemsSQL {
	s.IPAddress = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) SetNodeId(v string) *DescribeAuditRecordsResponseBodyItemsSQL {
	s.NodeId = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) SetSQLText(v string) *DescribeAuditRecordsResponseBodyItemsSQL {
	s.SQLText = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) SetSQLType(v string) *DescribeAuditRecordsResponseBodyItemsSQL {
	s.SQLType = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) SetTotalExecutionTimes(v string) *DescribeAuditRecordsResponseBodyItemsSQL {
	s.TotalExecutionTimes = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQL) Validate() error {
	return dara.Validate(s)
}
