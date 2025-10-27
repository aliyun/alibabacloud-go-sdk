// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditLogRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAuditLogRecordsResponseBody
	GetDBClusterId() *string
	SetItems(v []*DescribeAuditLogRecordsResponseBodyItems) *DescribeAuditLogRecordsResponseBody
	GetItems() []*DescribeAuditLogRecordsResponseBodyItems
	SetPageNumber(v string) *DescribeAuditLogRecordsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeAuditLogRecordsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeAuditLogRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeAuditLogRecordsResponseBody
	GetTotalCount() *string
}

type DescribeAuditLogRecordsResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// example:
	//
	// amv-t4nj8619bz2w3****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The queried SQL audit logs.
	Items []*DescribeAuditLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8A564B7F-8C00-43C0-8EC5-919FBB70573
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 6974
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAuditLogRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAuditLogRecordsResponseBody) GetItems() []*DescribeAuditLogRecordsResponseBodyItems {
	return s.Items
}

func (s *DescribeAuditLogRecordsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeAuditLogRecordsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAuditLogRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAuditLogRecordsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeAuditLogRecordsResponseBody) SetDBClusterId(v string) *DescribeAuditLogRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetItems(v []*DescribeAuditLogRecordsResponseBodyItems) *DescribeAuditLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetPageNumber(v string) *DescribeAuditLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetPageSize(v string) *DescribeAuditLogRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetRequestId(v string) *DescribeAuditLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetTotalCount(v string) *DescribeAuditLogRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) Validate() error {
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

type DescribeAuditLogRecordsResponseBodyItems struct {
	// The connection ID.
	//
	// example:
	//
	// 14356****
	ConnId *string `json:"ConnId,omitempty" xml:"ConnId,omitempty"`
	// The name of the database on which the SQL statement was executed.
	//
	// example:
	//
	// adb_demo
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The start time of the execution of the SQL statement. The time is displayed in the ISO 8601 standard in the yyyy-MM-dd HH:mm:ss format. The time must be in UTC.
	//
	// example:
	//
	// 2022-08-12 10:10:00
	ExecuteTime      *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	ExecuteTimestamp *int64  `json:"ExecuteTimestamp,omitempty" xml:"ExecuteTimestamp,omitempty"`
	// The IP address and port number of the client that is used to execute the SQL statement.
	//
	// example:
	//
	// 100.104.XX.XX:43908
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 202106081752021720161662490345362390
	ProcessID *string `json:"ProcessID,omitempty" xml:"ProcessID,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// SELECT 	- FROM adb_hdfs_import_source
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// The type of the SQL statement.
	//
	// example:
	//
	// SELECT
	SQLType *string `json:"SQLType,omitempty" xml:"SQLType,omitempty"`
	// Indicates whether the SQL statement was successfully executed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Succeed *string `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	// The amount of time that is consumed to execute the SQL statement. Unit: milliseconds.
	//
	// example:
	//
	// 216
	TotalTime *string `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	// The username that is used to execute the SQL statement.
	//
	// example:
	//
	// test
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeAuditLogRecordsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsResponseBodyItems) GetConnId() *string {
	return s.ConnId
}

func (s *DescribeAuditLogRecordsResponseBodyItems) GetDBName() *string {
	return s.DBName
}

func (s *DescribeAuditLogRecordsResponseBodyItems) GetExecuteTime() *string {
	return s.ExecuteTime
}

func (s *DescribeAuditLogRecordsResponseBodyItems) GetExecuteTimestamp() *int64 {
	return s.ExecuteTimestamp
}

func (s *DescribeAuditLogRecordsResponseBodyItems) GetHostAddress() *string {
	return s.HostAddress
}

func (s *DescribeAuditLogRecordsResponseBodyItems) GetProcessID() *string {
	return s.ProcessID
}

func (s *DescribeAuditLogRecordsResponseBodyItems) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeAuditLogRecordsResponseBodyItems) GetSQLType() *string {
	return s.SQLType
}

func (s *DescribeAuditLogRecordsResponseBodyItems) GetSucceed() *string {
	return s.Succeed
}

func (s *DescribeAuditLogRecordsResponseBodyItems) GetTotalTime() *string {
	return s.TotalTime
}

func (s *DescribeAuditLogRecordsResponseBodyItems) GetUser() *string {
	return s.User
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetConnId(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.ConnId = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetDBName(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetExecuteTime(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetExecuteTimestamp(v int64) *DescribeAuditLogRecordsResponseBodyItems {
	s.ExecuteTimestamp = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetHostAddress(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.HostAddress = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetProcessID(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.ProcessID = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetSQLText(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetSQLType(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.SQLType = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetSucceed(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.Succeed = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetTotalTime(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.TotalTime = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetUser(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.User = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
