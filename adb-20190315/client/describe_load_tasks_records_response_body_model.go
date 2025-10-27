// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadTasksRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeLoadTasksRecordsResponseBody
	GetDBClusterId() *string
	SetLoadTasksRecords(v []*DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) *DescribeLoadTasksRecordsResponseBody
	GetLoadTasksRecords() []*DescribeLoadTasksRecordsResponseBodyLoadTasksRecords
	SetPageNumber(v string) *DescribeLoadTasksRecordsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeLoadTasksRecordsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeLoadTasksRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeLoadTasksRecordsResponseBody
	GetTotalCount() *string
}

type DescribeLoadTasksRecordsResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// am-bp2590j****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The queried asynchronous import and export tasks.
	LoadTasksRecords []*DescribeLoadTasksRecordsResponseBodyLoadTasksRecords `json:"LoadTasksRecords,omitempty" xml:"LoadTasksRecords,omitempty" type:"Repeated"`
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
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C60B05DB-2B77-421A-98E9-92C20E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLoadTasksRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadTasksRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadTasksRecordsResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeLoadTasksRecordsResponseBody) GetLoadTasksRecords() []*DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	return s.LoadTasksRecords
}

func (s *DescribeLoadTasksRecordsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeLoadTasksRecordsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeLoadTasksRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadTasksRecordsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeLoadTasksRecordsResponseBody) SetDBClusterId(v string) *DescribeLoadTasksRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) SetLoadTasksRecords(v []*DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) *DescribeLoadTasksRecordsResponseBody {
	s.LoadTasksRecords = v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) SetPageNumber(v string) *DescribeLoadTasksRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) SetPageSize(v string) *DescribeLoadTasksRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) SetRequestId(v string) *DescribeLoadTasksRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) SetTotalCount(v string) *DescribeLoadTasksRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) Validate() error {
	if s.LoadTasksRecords != nil {
		for _, item := range s.LoadTasksRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadTasksRecordsResponseBodyLoadTasksRecords struct {
	// The start time of the task. The time is accurate to milliseconds. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ss.SSSZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-05-18 18:47:27.0
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the database that is involved in the import or export task.
	//
	// example:
	//
	// adb_demo
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 2021051818472717201616624903453******
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The process ID.
	//
	// example:
	//
	// 2021051818472717201616624903453******
	ProcessID *string `json:"ProcessID,omitempty" xml:"ProcessID,omitempty"`
	// The number of rows that are processed in the asynchronous import or export task.
	//
	// example:
	//
	// 6
	ProcessRows *int64 `json:"ProcessRows,omitempty" xml:"ProcessRows,omitempty"`
	// The SQL statement that is used in the asynchronous import or export task.
	//
	// example:
	//
	// insert overwrite into courses_external_table\\nselect 	- from courses
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// The state of the task.
	//
	// example:
	//
	// FINISH
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the task state was updated. The time is accurate to milliseconds. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ss.SSSZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-05-18 18:47:31.0
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) GoString() string {
	return s.String()
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) GetDBName() *string {
	return s.DBName
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) GetJobName() *string {
	return s.JobName
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) GetProcessID() *string {
	return s.ProcessID
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) GetProcessRows() *int64 {
	return s.ProcessRows
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) GetSql() *string {
	return s.Sql
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) GetState() *string {
	return s.State
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetCreateTime(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.CreateTime = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetDBName(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.DBName = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetJobName(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.JobName = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetProcessID(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.ProcessID = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetProcessRows(v int64) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.ProcessRows = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetSql(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.Sql = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetState(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.State = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetUpdateTime(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.UpdateTime = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) Validate() error {
	return dara.Validate(s)
}
