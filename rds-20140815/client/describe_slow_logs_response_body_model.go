// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSlowLogsResponseBody
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeSlowLogsResponseBody
	GetEndTime() *string
	SetEngine(v string) *DescribeSlowLogsResponseBody
	GetEngine() *string
	SetItems(v *DescribeSlowLogsResponseBodyItems) *DescribeSlowLogsResponseBody
	GetItems() *DescribeSlowLogsResponseBodyItems
	SetPageNumber(v int32) *DescribeSlowLogsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeSlowLogsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeSlowLogsResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeSlowLogsResponseBody
	GetStartTime() *string
	SetTotalRecordCount(v int32) *DescribeSlowLogsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeSlowLogsResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end date of the query.
	//
	// example:
	//
	// 2011-05-30Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The database engine of the instance.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// An array that consists of the information about each slow query log.
	Items *DescribeSlowLogsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The number of the page returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of SQL statements that are returned on the current page.
	//
	// example:
	//
	// 10
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2553A660-E4EB-4AF4-A402-8AFF70A49143
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start date of the query.
	//
	// example:
	//
	// 2011-05-30Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 5
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSlowLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogsResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSlowLogsResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSlowLogsResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeSlowLogsResponseBody) GetItems() *DescribeSlowLogsResponseBodyItems {
	return s.Items
}

func (s *DescribeSlowLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSlowLogsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeSlowLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlowLogsResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSlowLogsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeSlowLogsResponseBody) SetDBInstanceId(v string) *DescribeSlowLogsResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSlowLogsResponseBody) SetEndTime(v string) *DescribeSlowLogsResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBody) SetEngine(v string) *DescribeSlowLogsResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeSlowLogsResponseBody) SetItems(v *DescribeSlowLogsResponseBodyItems) *DescribeSlowLogsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSlowLogsResponseBody) SetPageNumber(v int32) *DescribeSlowLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogsResponseBody) SetPageRecordCount(v int32) *DescribeSlowLogsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSlowLogsResponseBody) SetRequestId(v string) *DescribeSlowLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogsResponseBody) SetStartTime(v string) *DescribeSlowLogsResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBody) SetTotalRecordCount(v int32) *DescribeSlowLogsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSlowLogsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSlowLogsResponseBodyItems struct {
	SQLSlowLog []*DescribeSlowLogsResponseBodyItemsSQLSlowLog `json:"SQLSlowLog,omitempty" xml:"SQLSlowLog,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogsResponseBodyItems) GetSQLSlowLog() []*DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	return s.SQLSlowLog
}

func (s *DescribeSlowLogsResponseBodyItems) SetSQLSlowLog(v []*DescribeSlowLogsResponseBodyItemsSQLSlowLog) *DescribeSlowLogsResponseBodyItems {
	s.SQLSlowLog = v
	return s
}

func (s *DescribeSlowLogsResponseBodyItems) Validate() error {
	if s.SQLSlowLog != nil {
		for _, item := range s.SQLSlowLog {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSlowLogsResponseBodyItemsSQLSlowLog struct {
	// The average execution duration per SQL statement in the query. Unit: seconds.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 1
	AvgExecutionTime *int64 `json:"AvgExecutionTime,omitempty" xml:"AvgExecutionTime,omitempty"`
	// The average number of I/O writes per SQL statement in the query.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	AvgIOWriteCounts *int64 `json:"AvgIOWriteCounts,omitempty" xml:"AvgIOWriteCounts,omitempty"`
	// The average number of rows that were affected by the last SQL statement in the query.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	AvgLastRowsAffectedCounts *int64 `json:"AvgLastRowsAffectedCounts,omitempty" xml:"AvgLastRowsAffectedCounts,omitempty"`
	// The average number of logical reads per SQL statement in the query.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	AvgLogicalReadCounts *int64 `json:"AvgLogicalReadCounts,omitempty" xml:"AvgLogicalReadCounts,omitempty"`
	// The average number of physical reads per SQL statement in the query.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	AvgPhysicalReadCounts *int64 `json:"AvgPhysicalReadCounts,omitempty" xml:"AvgPhysicalReadCounts,omitempty"`
	// The average number of rows that were affected per SQL statement in the query.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	AvgRowsAffectedCounts *int64 `json:"AvgRowsAffectedCounts,omitempty" xml:"AvgRowsAffectedCounts,omitempty"`
	// The date when the data was generated.
	//
	// example:
	//
	// 2011-05-30Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// RDS_MySQL
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The longest execution duration of a specific SQL statement in the query. Unit: seconds.
	//
	// example:
	//
	// 60
	MaxExecutionTime *int64 `json:"MaxExecutionTime,omitempty" xml:"MaxExecutionTime,omitempty"`
	// The longest execution duration of a specific SQL statement in the query. Unit: milliseconds.
	//
	// example:
	//
	// 60000
	MaxExecutionTimeMS *int64 `json:"MaxExecutionTimeMS,omitempty" xml:"MaxExecutionTimeMS,omitempty"`
	// The largest number of I/O writes that were performed by a specific SQL statement in the query.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	MaxIOWriteCounts *int64 `json:"MaxIOWriteCounts,omitempty" xml:"MaxIOWriteCounts,omitempty"`
	// The largest number of rows that were affected by the last SQL statement in the query.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	MaxLastRowsAffectedCounts *int64 `json:"MaxLastRowsAffectedCounts,omitempty" xml:"MaxLastRowsAffectedCounts,omitempty"`
	// The longest lock duration that was caused by a specific SQL statement in the query. Unit: seconds.
	//
	// example:
	//
	// 0
	MaxLockTime *int64 `json:"MaxLockTime,omitempty" xml:"MaxLockTime,omitempty"`
	// The longest lock duration that was caused by a specific SQL statement in the query. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	MaxLockTimeMS *int64 `json:"MaxLockTimeMS,omitempty" xml:"MaxLockTimeMS,omitempty"`
	// The largest number of logical reads that were performed by a specific SQL statement in the query.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	MaxLogicalReadCounts *int64 `json:"MaxLogicalReadCounts,omitempty" xml:"MaxLogicalReadCounts,omitempty"`
	// The largest number of physical reads that were performed by a specific SQL statement in the query.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	MaxPhysicalReadCounts *int64 `json:"MaxPhysicalReadCounts,omitempty" xml:"MaxPhysicalReadCounts,omitempty"`
	// The largest number of rows that were affected by a specific SQL statement in the query.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	MaxRowsAffectedCounts *int64 `json:"MaxRowsAffectedCounts,omitempty" xml:"MaxRowsAffectedCounts,omitempty"`
	// The smallest number of I/O writes that were performed by a specific SQL statement in the query.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	MinIOWriteCounts *int64 `json:"MinIOWriteCounts,omitempty" xml:"MinIOWriteCounts,omitempty"`
	// The smallest number of rows that were affected by the last SQL statement in the query.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	MinLastRowsAffectedCounts *int64 `json:"MinLastRowsAffectedCounts,omitempty" xml:"MinLastRowsAffectedCounts,omitempty"`
	// The smallest number of logical reads that were performed by a specific SQL statement in the query.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	MinLogicalReadCounts *int64 `json:"MinLogicalReadCounts,omitempty" xml:"MinLogicalReadCounts,omitempty"`
	// The smallest number of physical reads that were performed by a specific SQL statement in the query.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	MinPhysicalReadCounts *int64 `json:"MinPhysicalReadCounts,omitempty" xml:"MinPhysicalReadCounts,omitempty"`
	// The smallest number of rows that were affected by a specific SQL statement in the query.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	MinRowsAffectedCounts *int64 `json:"MinRowsAffectedCounts,omitempty" xml:"MinRowsAffectedCounts,omitempty"`
	// The total number of SQL statements that were executed in the query. This parameter is returned only for instances that run MySQL.
	//
	// example:
	//
	// 1
	MySQLTotalExecutionCounts *int64 `json:"MySQLTotalExecutionCounts,omitempty" xml:"MySQLTotalExecutionCounts,omitempty"`
	// The total execution duration of all SQL statements in the query. Unit: seconds. This parameter is returned only for instances that run MySQL.
	//
	// example:
	//
	// 1
	MySQLTotalExecutionTimes *int64 `json:"MySQLTotalExecutionTimes,omitempty" xml:"MySQLTotalExecutionTimes,omitempty"`
	// The largest number of rows that were parsed by a specific SQL statement in the query.
	//
	// example:
	//
	// 1
	ParseMaxRowCount *int64 `json:"ParseMaxRowCount,omitempty" xml:"ParseMaxRowCount,omitempty"`
	// The total number of rows that were parsed by all SQL statements in the query.
	//
	// example:
	//
	// 1
	ParseTotalRowCounts *int64 `json:"ParseTotalRowCounts,omitempty" xml:"ParseTotalRowCounts,omitempty"`
	// The date on which the data report was generated.
	//
	// example:
	//
	// 2011-05-30Z
	ReportTime *string `json:"ReportTime,omitempty" xml:"ReportTime,omitempty"`
	// The largest number of rows that were returned by a specific SQL statement in the query.
	//
	// example:
	//
	// 1
	ReturnMaxRowCount *int64 `json:"ReturnMaxRowCount,omitempty" xml:"ReturnMaxRowCount,omitempty"`
	// The total number of rows that were returned by all SQL statements in the query.
	//
	// example:
	//
	// 1
	ReturnTotalRowCounts *int64 `json:"ReturnTotalRowCounts,omitempty" xml:"ReturnTotalRowCounts,omitempty"`
	// The unique ID of the SQL statement. The ID is used to obtain the slow query logs of the SQL statement.
	//
	// example:
	//
	// U2FsdGVkxxxx
	SQLHASH *string `json:"SQLHASH,omitempty" xml:"SQLHASH,omitempty"`
	// The ID of the SQL statement in the statistical template of slow query logs. This parameter is replaced by the **SQLHASH*	- parameter.
	//
	// example:
	//
	// 521584
	SQLIdStr *string `json:"SQLIdStr,omitempty" xml:"SQLIdStr,omitempty"`
	// The average amount of CPU time per SQL statement in the query. Unit: seconds.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	SQLServerAvgCpuTime *int64 `json:"SQLServerAvgCpuTime,omitempty" xml:"SQLServerAvgCpuTime,omitempty"`
	// The average execution duration per SQL statement in the query. Unit: seconds.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	SQLServerAvgExecutionTime *int64 `json:"SQLServerAvgExecutionTime,omitempty" xml:"SQLServerAvgExecutionTime,omitempty"`
	// The largest amount of CPU time that was used by a specific SQL statement in the query. Unit: seconds.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	SQLServerMaxCpuTime *int64 `json:"SQLServerMaxCpuTime,omitempty" xml:"SQLServerMaxCpuTime,omitempty"`
	// The smallest amount of CPU time that was used by a specific SQL statement in the query. Unit: seconds.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	SQLServerMinCpuTime *int64 `json:"SQLServerMinCpuTime,omitempty" xml:"SQLServerMinCpuTime,omitempty"`
	// The smallest execution duration of a specific SQL statement in the query. Unit: seconds.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	SQLServerMinExecutionTime *int64 `json:"SQLServerMinExecutionTime,omitempty" xml:"SQLServerMinExecutionTime,omitempty"`
	// The total amount of CPU time that was used by all SQL statements in the query. Unit: seconds.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	SQLServerTotalCpuTime *int64 `json:"SQLServerTotalCpuTime,omitempty" xml:"SQLServerTotalCpuTime,omitempty"`
	// The total number of SQL statements that were executed in the query. This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 1
	SQLServerTotalExecutionCounts *int64 `json:"SQLServerTotalExecutionCounts,omitempty" xml:"SQLServerTotalExecutionCounts,omitempty"`
	// The total execution duration of all SQL statements in the query. This parameter is returned only for instances that run SQL Server. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	SQLServerTotalExecutionTimes *int64 `json:"SQLServerTotalExecutionTimes,omitempty" xml:"SQLServerTotalExecutionTimes,omitempty"`
	// The SQL statement that was executed in the query.
	//
	// example:
	//
	// select id,name from tb_table
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// The ID of the slow query log summary.
	//
	// example:
	//
	// 26584213
	SlowLogId *int64 `json:"SlowLogId,omitempty" xml:"SlowLogId,omitempty"`
	// The total number of I/O writes that were performed by all SQL statements in the query.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	TotalIOWriteCounts *int64 `json:"TotalIOWriteCounts,omitempty" xml:"TotalIOWriteCounts,omitempty"`
	// The total number of rows that were affected by the last SQL statement in the query.
	//
	// >  This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0
	TotalLastRowsAffectedCounts *int64 `json:"TotalLastRowsAffectedCounts,omitempty" xml:"TotalLastRowsAffectedCounts,omitempty"`
	// The total lock duration that was caused by all SQL statements in the query. Unit: seconds.
	//
	// example:
	//
	// 0
	TotalLockTimes *int64 `json:"TotalLockTimes,omitempty" xml:"TotalLockTimes,omitempty"`
	// The total number of logical reads that were performed by all SQL statements in the query.
	//
	// example:
	//
	// 1
	TotalLogicalReadCounts *int64 `json:"TotalLogicalReadCounts,omitempty" xml:"TotalLogicalReadCounts,omitempty"`
	// The total number of physical reads that were performed by all SQL statements in the query.
	//
	// example:
	//
	// 1
	TotalPhysicalReadCounts *int64 `json:"TotalPhysicalReadCounts,omitempty" xml:"TotalPhysicalReadCounts,omitempty"`
	// The total number of rows that were affected by all SQL statements in the query.
	//
	// example:
	//
	// 0
	TotalRowsAffectedCounts *int64 `json:"TotalRowsAffectedCounts,omitempty" xml:"TotalRowsAffectedCounts,omitempty"`
}

func (s DescribeSlowLogsResponseBodyItemsSQLSlowLog) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogsResponseBodyItemsSQLSlowLog) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetAvgExecutionTime() *int64 {
	return s.AvgExecutionTime
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetAvgIOWriteCounts() *int64 {
	return s.AvgIOWriteCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetAvgLastRowsAffectedCounts() *int64 {
	return s.AvgLastRowsAffectedCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetAvgLogicalReadCounts() *int64 {
	return s.AvgLogicalReadCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetAvgPhysicalReadCounts() *int64 {
	return s.AvgPhysicalReadCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetAvgRowsAffectedCounts() *int64 {
	return s.AvgRowsAffectedCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMaxExecutionTime() *int64 {
	return s.MaxExecutionTime
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMaxExecutionTimeMS() *int64 {
	return s.MaxExecutionTimeMS
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMaxIOWriteCounts() *int64 {
	return s.MaxIOWriteCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMaxLastRowsAffectedCounts() *int64 {
	return s.MaxLastRowsAffectedCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMaxLockTime() *int64 {
	return s.MaxLockTime
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMaxLockTimeMS() *int64 {
	return s.MaxLockTimeMS
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMaxLogicalReadCounts() *int64 {
	return s.MaxLogicalReadCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMaxPhysicalReadCounts() *int64 {
	return s.MaxPhysicalReadCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMaxRowsAffectedCounts() *int64 {
	return s.MaxRowsAffectedCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMinIOWriteCounts() *int64 {
	return s.MinIOWriteCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMinLastRowsAffectedCounts() *int64 {
	return s.MinLastRowsAffectedCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMinLogicalReadCounts() *int64 {
	return s.MinLogicalReadCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMinPhysicalReadCounts() *int64 {
	return s.MinPhysicalReadCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMinRowsAffectedCounts() *int64 {
	return s.MinRowsAffectedCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMySQLTotalExecutionCounts() *int64 {
	return s.MySQLTotalExecutionCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMySQLTotalExecutionTimes() *int64 {
	return s.MySQLTotalExecutionTimes
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetParseMaxRowCount() *int64 {
	return s.ParseMaxRowCount
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetParseTotalRowCounts() *int64 {
	return s.ParseTotalRowCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetReportTime() *string {
	return s.ReportTime
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetReturnMaxRowCount() *int64 {
	return s.ReturnMaxRowCount
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetReturnTotalRowCounts() *int64 {
	return s.ReturnTotalRowCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetSQLHASH() *string {
	return s.SQLHASH
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetSQLIdStr() *string {
	return s.SQLIdStr
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetSQLServerAvgCpuTime() *int64 {
	return s.SQLServerAvgCpuTime
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetSQLServerAvgExecutionTime() *int64 {
	return s.SQLServerAvgExecutionTime
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetSQLServerMaxCpuTime() *int64 {
	return s.SQLServerMaxCpuTime
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetSQLServerMinCpuTime() *int64 {
	return s.SQLServerMinCpuTime
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetSQLServerMinExecutionTime() *int64 {
	return s.SQLServerMinExecutionTime
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetSQLServerTotalCpuTime() *int64 {
	return s.SQLServerTotalCpuTime
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetSQLServerTotalExecutionCounts() *int64 {
	return s.SQLServerTotalExecutionCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetSQLServerTotalExecutionTimes() *int64 {
	return s.SQLServerTotalExecutionTimes
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetSlowLogId() *int64 {
	return s.SlowLogId
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetTotalIOWriteCounts() *int64 {
	return s.TotalIOWriteCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetTotalLastRowsAffectedCounts() *int64 {
	return s.TotalLastRowsAffectedCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetTotalLockTimes() *int64 {
	return s.TotalLockTimes
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetTotalLogicalReadCounts() *int64 {
	return s.TotalLogicalReadCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetTotalPhysicalReadCounts() *int64 {
	return s.TotalPhysicalReadCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetTotalRowsAffectedCounts() *int64 {
	return s.TotalRowsAffectedCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetAvgExecutionTime(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.AvgExecutionTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetAvgIOWriteCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.AvgIOWriteCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetAvgLastRowsAffectedCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.AvgLastRowsAffectedCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetAvgLogicalReadCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.AvgLogicalReadCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetAvgPhysicalReadCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.AvgPhysicalReadCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetAvgRowsAffectedCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.AvgRowsAffectedCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetCreateTime(v string) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.CreateTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetDBName(v string) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMaxExecutionTime(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MaxExecutionTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMaxExecutionTimeMS(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MaxExecutionTimeMS = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMaxIOWriteCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MaxIOWriteCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMaxLastRowsAffectedCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MaxLastRowsAffectedCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMaxLockTime(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MaxLockTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMaxLockTimeMS(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MaxLockTimeMS = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMaxLogicalReadCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MaxLogicalReadCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMaxPhysicalReadCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MaxPhysicalReadCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMaxRowsAffectedCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MaxRowsAffectedCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMinIOWriteCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MinIOWriteCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMinLastRowsAffectedCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MinLastRowsAffectedCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMinLogicalReadCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MinLogicalReadCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMinPhysicalReadCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MinPhysicalReadCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMinRowsAffectedCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MinRowsAffectedCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMySQLTotalExecutionCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MySQLTotalExecutionCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMySQLTotalExecutionTimes(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MySQLTotalExecutionTimes = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetParseMaxRowCount(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.ParseMaxRowCount = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetParseTotalRowCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.ParseTotalRowCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetReportTime(v string) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.ReportTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetReturnMaxRowCount(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.ReturnMaxRowCount = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetReturnTotalRowCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.ReturnTotalRowCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetSQLHASH(v string) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.SQLHASH = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetSQLIdStr(v string) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.SQLIdStr = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetSQLServerAvgCpuTime(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.SQLServerAvgCpuTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetSQLServerAvgExecutionTime(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.SQLServerAvgExecutionTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetSQLServerMaxCpuTime(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.SQLServerMaxCpuTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetSQLServerMinCpuTime(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.SQLServerMinCpuTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetSQLServerMinExecutionTime(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.SQLServerMinExecutionTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetSQLServerTotalCpuTime(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.SQLServerTotalCpuTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetSQLServerTotalExecutionCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.SQLServerTotalExecutionCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetSQLServerTotalExecutionTimes(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.SQLServerTotalExecutionTimes = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetSQLText(v string) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetSlowLogId(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.SlowLogId = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetTotalIOWriteCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.TotalIOWriteCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetTotalLastRowsAffectedCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.TotalLastRowsAffectedCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetTotalLockTimes(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.TotalLockTimes = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetTotalLogicalReadCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.TotalLogicalReadCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetTotalPhysicalReadCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.TotalPhysicalReadCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetTotalRowsAffectedCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.TotalRowsAffectedCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) Validate() error {
	return dara.Validate(s)
}
