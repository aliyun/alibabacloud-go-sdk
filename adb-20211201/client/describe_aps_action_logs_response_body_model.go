// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsActionLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeApsActionLogsResponseBody
	GetAccessDeniedDetail() *string
	SetActionLogs(v []*DescribeApsActionLogsResponseBodyActionLogs) *DescribeApsActionLogsResponseBody
	GetActionLogs() []*DescribeApsActionLogsResponseBodyActionLogs
	SetDBClusterId(v string) *DescribeApsActionLogsResponseBody
	GetDBClusterId() *string
	SetPageNumber(v string) *DescribeApsActionLogsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeApsActionLogsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeApsActionLogsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeApsActionLogsResponseBody
	GetTotalCount() *string
	SetWorkloadId(v string) *DescribeApsActionLogsResponseBody
	GetWorkloadId() *string
}

type DescribeApsActionLogsResponseBody struct {
	// The information about the request denial.
	//
	// example:
	//
	// {
	//
	//   "AuthAction": "xxx",
	//
	//   "AuthPrincipalDisplayName": "sampleName",
	//
	//   "AuthPrincipalOwnerId": "111111111111111111",
	//
	//   "AuthPrincipalType": "SubUser",
	//
	//   "AuthResource": "xxx",
	//
	//   "NoPermissionType": "xxx",
	//
	//   "PolicyType": "xxx"
	//
	// }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The queried logs.
	ActionLogs []*DescribeApsActionLogsResponseBodyActionLogs `json:"ActionLogs,omitempty" xml:"ActionLogs,omitempty" type:"Repeated"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// example:
	//
	// amv-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
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
	// A5EDBA27-AF3E-5966-9503-FD1557E19167
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The ID of the real-time data ingestion job.
	//
	// example:
	//
	// aps-hz109vpvt4fg8528d****
	WorkloadId *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
}

func (s DescribeApsActionLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsActionLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApsActionLogsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeApsActionLogsResponseBody) GetActionLogs() []*DescribeApsActionLogsResponseBodyActionLogs {
	return s.ActionLogs
}

func (s *DescribeApsActionLogsResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeApsActionLogsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeApsActionLogsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeApsActionLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApsActionLogsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeApsActionLogsResponseBody) GetWorkloadId() *string {
	return s.WorkloadId
}

func (s *DescribeApsActionLogsResponseBody) SetAccessDeniedDetail(v string) *DescribeApsActionLogsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeApsActionLogsResponseBody) SetActionLogs(v []*DescribeApsActionLogsResponseBodyActionLogs) *DescribeApsActionLogsResponseBody {
	s.ActionLogs = v
	return s
}

func (s *DescribeApsActionLogsResponseBody) SetDBClusterId(v string) *DescribeApsActionLogsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeApsActionLogsResponseBody) SetPageNumber(v string) *DescribeApsActionLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApsActionLogsResponseBody) SetPageSize(v string) *DescribeApsActionLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApsActionLogsResponseBody) SetRequestId(v string) *DescribeApsActionLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApsActionLogsResponseBody) SetTotalCount(v string) *DescribeApsActionLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApsActionLogsResponseBody) SetWorkloadId(v string) *DescribeApsActionLogsResponseBody {
	s.WorkloadId = &v
	return s
}

func (s *DescribeApsActionLogsResponseBody) Validate() error {
	if s.ActionLogs != nil {
		for _, item := range s.ActionLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApsActionLogsResponseBodyActionLogs struct {
	// The content of the log.
	//
	// example:
	//
	// DDL migration job finished
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// The phase during which the log was generated. Valid values:
	//
	// 	- **StructureMigrate**: schema migration.
	//
	// 	- **FullDataSync**: full data synchronization.
	//
	// 	- **IncrementalSync**: incremental data synchronization.
	//
	// example:
	//
	// FullDataSync
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
	// The type of the log. Multiple log types are separated by commas (,). Valid values:
	//
	// 	- **INFO**
	//
	// 	- **WARN**
	//
	// 	- **ERROR**
	//
	// example:
	//
	// INFO,WARN,ERROR
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the log was generated. The time follows the ISO 8601 standard in the **yyyy-MM-ddTHH:mm:ssZ*	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-02-01T05:46:30Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeApsActionLogsResponseBodyActionLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsActionLogsResponseBodyActionLogs) GoString() string {
	return s.String()
}

func (s *DescribeApsActionLogsResponseBodyActionLogs) GetContext() *string {
	return s.Context
}

func (s *DescribeApsActionLogsResponseBodyActionLogs) GetStage() *string {
	return s.Stage
}

func (s *DescribeApsActionLogsResponseBodyActionLogs) GetState() *string {
	return s.State
}

func (s *DescribeApsActionLogsResponseBodyActionLogs) GetTime() *string {
	return s.Time
}

func (s *DescribeApsActionLogsResponseBodyActionLogs) SetContext(v string) *DescribeApsActionLogsResponseBodyActionLogs {
	s.Context = &v
	return s
}

func (s *DescribeApsActionLogsResponseBodyActionLogs) SetStage(v string) *DescribeApsActionLogsResponseBodyActionLogs {
	s.Stage = &v
	return s
}

func (s *DescribeApsActionLogsResponseBodyActionLogs) SetState(v string) *DescribeApsActionLogsResponseBodyActionLogs {
	s.State = &v
	return s
}

func (s *DescribeApsActionLogsResponseBodyActionLogs) SetTime(v string) *DescribeApsActionLogsResponseBodyActionLogs {
	s.Time = &v
	return s
}

func (s *DescribeApsActionLogsResponseBodyActionLogs) Validate() error {
	return dara.Validate(s)
}
