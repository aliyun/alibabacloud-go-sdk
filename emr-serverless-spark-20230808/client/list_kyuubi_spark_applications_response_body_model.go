// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKyuubiSparkApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ListKyuubiSparkApplicationsResponseBodyApplications) *ListKyuubiSparkApplicationsResponseBody
	GetApplications() []*ListKyuubiSparkApplicationsResponseBodyApplications
	SetMaxResults(v int32) *ListKyuubiSparkApplicationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListKyuubiSparkApplicationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListKyuubiSparkApplicationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListKyuubiSparkApplicationsResponseBody
	GetTotalCount() *int32
}

type ListKyuubiSparkApplicationsResponseBody struct {
	// The details of the applications.
	Applications []*ListKyuubiSparkApplicationsResponseBodyApplications `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListKyuubiSparkApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKyuubiSparkApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKyuubiSparkApplicationsResponseBody) GetApplications() []*ListKyuubiSparkApplicationsResponseBodyApplications {
	return s.Applications
}

func (s *ListKyuubiSparkApplicationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKyuubiSparkApplicationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKyuubiSparkApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKyuubiSparkApplicationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListKyuubiSparkApplicationsResponseBody) SetApplications(v []*ListKyuubiSparkApplicationsResponseBodyApplications) *ListKyuubiSparkApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBody) SetMaxResults(v int32) *ListKyuubiSparkApplicationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBody) SetNextToken(v string) *ListKyuubiSparkApplicationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBody) SetRequestId(v string) *ListKyuubiSparkApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBody) SetTotalCount(v int32) *ListKyuubiSparkApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBody) Validate() error {
	if s.Applications != nil {
		for _, item := range s.Applications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKyuubiSparkApplicationsResponseBodyApplications struct {
	// The ID of the application that is submitted by using a Kyuubi gateway.
	//
	// example:
	//
	// spark-339f844005b6404c95f9f7c7a13b****
	ApplicationId *string `json:"applicationId,omitempty" xml:"applicationId,omitempty"`
	// The name of the Spark application that is submitted by using a Kyuubi gateway.
	//
	// example:
	//
	// kyuubi-connection-spark-sql-anonymous-fa9a5e73-b4b1-474a-b****
	ApplicationName *string `json:"applicationName,omitempty" xml:"applicationName,omitempty"`
	// The number of CUs consumed during a specified cycle of a task. The value is an estimated value. Refer to your Alibaba Cloud bill for the actual number of consumed CUs.
	//
	// example:
	//
	// 0.238302
	CuHours *float64 `json:"cuHours,omitempty" xml:"cuHours,omitempty"`
	// The time when the task ended.
	//
	// example:
	//
	// 2025-02-12 20:02:02
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// Exit Code: 137, Exit Reason: OOMKilled
	ExitReason *string `json:"exitReason,omitempty" xml:"exitReason,omitempty"`
	// example:
	//
	// kb-2b93ec*******c9440c
	KyuubiServiceId *string `json:"kyuubiServiceId,omitempty" xml:"kyuubiServiceId,omitempty"`
	// example:
	//
	// SUCCESS
	LatestSqlStatementStatus *string `json:"latestSqlStatementStatus,omitempty" xml:"latestSqlStatementStatus,omitempty"`
	// The total amount of memory allocated to the job multiplied by the running duration (seconds).
	//
	// example:
	//
	// 3513900
	MbSeconds *int64 `json:"mbSeconds,omitempty" xml:"mbSeconds,omitempty"`
	// The name of the resource queue on which the Spark jobs run.
	//
	// example:
	//
	// dev_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	RunLog          *RunLog `json:"runLog,omitempty" xml:"runLog,omitempty"`
	// The time when the task started.
	//
	// example:
	//
	// 2025-02-12 19:59:16
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The status of the Spark application.
	//
	// 	- STARTING
	//
	// 	- RUNNING
	//
	// 	- TERMINATED
	//
	// example:
	//
	// STARTING
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The total number of CPU cores allocated to the job multiplied by the running duration (seconds).
	//
	// example:
	//
	// 780
	VcoreSeconds *int64 `json:"vcoreSeconds,omitempty" xml:"vcoreSeconds,omitempty"`
	// The URL of the web UI for the Spark application.
	WebUI *string `json:"webUI,omitempty" xml:"webUI,omitempty"`
}

func (s ListKyuubiSparkApplicationsResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListKyuubiSparkApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) GetCuHours() *float64 {
	return s.CuHours
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) GetEndTime() *string {
	return s.EndTime
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) GetExitReason() *string {
	return s.ExitReason
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) GetKyuubiServiceId() *string {
	return s.KyuubiServiceId
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) GetLatestSqlStatementStatus() *string {
	return s.LatestSqlStatementStatus
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) GetMbSeconds() *int64 {
	return s.MbSeconds
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) GetResourceQueueId() *string {
	return s.ResourceQueueId
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) GetRunLog() *RunLog {
	return s.RunLog
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) GetStartTime() *string {
	return s.StartTime
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) GetState() *string {
	return s.State
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) GetVcoreSeconds() *int64 {
	return s.VcoreSeconds
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) GetWebUI() *string {
	return s.WebUI
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetApplicationId(v string) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetApplicationName(v string) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.ApplicationName = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetCuHours(v float64) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.CuHours = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetEndTime(v string) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.EndTime = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetExitReason(v string) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.ExitReason = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetKyuubiServiceId(v string) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.KyuubiServiceId = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetLatestSqlStatementStatus(v string) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.LatestSqlStatementStatus = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetMbSeconds(v int64) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.MbSeconds = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetResourceQueueId(v string) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.ResourceQueueId = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetRunLog(v *RunLog) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.RunLog = v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetStartTime(v string) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.StartTime = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetState(v string) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.State = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetVcoreSeconds(v int64) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.VcoreSeconds = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetWebUI(v string) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.WebUI = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) Validate() error {
	if s.RunLog != nil {
		if err := s.RunLog.Validate(); err != nil {
			return err
		}
	}
	return nil
}
