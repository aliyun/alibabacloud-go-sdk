// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataDiagnosisJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataDiagnosisJobs(v []*ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) *ListDataDiagnosisJobsResponseBody
	GetDataDiagnosisJobs() []*ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs
	SetRequestId(v string) *ListDataDiagnosisJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListDataDiagnosisJobsResponseBody
	GetTotalCount() *int64
}

type ListDataDiagnosisJobsResponseBody struct {
	// The returned data diagnosis jobs.
	DataDiagnosisJobs []*ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs `json:"DataDiagnosisJobs,omitempty" xml:"DataDiagnosisJobs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataDiagnosisJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosisJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosisJobsResponseBody) GetDataDiagnosisJobs() []*ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs {
	return s.DataDiagnosisJobs
}

func (s *ListDataDiagnosisJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataDiagnosisJobsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDataDiagnosisJobsResponseBody) SetDataDiagnosisJobs(v []*ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) *ListDataDiagnosisJobsResponseBody {
	s.DataDiagnosisJobs = v
	return s
}

func (s *ListDataDiagnosisJobsResponseBody) SetRequestId(v string) *ListDataDiagnosisJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataDiagnosisJobsResponseBody) SetTotalCount(v int64) *ListDataDiagnosisJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataDiagnosisJobsResponseBody) Validate() error {
	if s.DataDiagnosisJobs != nil {
		for _, item := range s.DataDiagnosisJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs struct {
	// The business date.
	//
	// example:
	//
	// 2023-05-28
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The job configuration.
	//
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The data diagnosis configuration ID.
	//
	// example:
	//
	// 3
	DataDiagnosisConfigId *string `json:"DataDiagnosisConfigId,omitempty" xml:"DataDiagnosisConfigId,omitempty"`
	// The data diagnosis configuration name.
	//
	// example:
	//
	// 异常行为分析-xxx
	DataDiagnosisConfigName *string `json:"DataDiagnosisConfigName,omitempty" xml:"DataDiagnosisConfigName,omitempty"`
	// The data diagnosis job ID.
	//
	// example:
	//
	// 3
	DataDiagnosisJobId *string `json:"DataDiagnosisJobId,omitempty" xml:"DataDiagnosisJobId,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the job started.
	//
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtStartTime *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	// The job source. Valid values:
	//
	// - Rerun: A manual rerun.
	//
	// - Cycle: A periodic run.
	//
	// - ByInitializing: Generated during initialization.
	//
	// example:
	//
	// Cycle
	JobSource *string `json:"JobSource,omitempty" xml:"JobSource,omitempty"`
	// The log messages.
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The task status.
	//
	// example:
	//
	// Initializing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The data diagnosis type.
	//
	// example:
	//
	// ChangeRate
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) GetBizDate() *string {
	return s.BizDate
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) GetConfig() *string {
	return s.Config
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) GetDataDiagnosisConfigId() *string {
	return s.DataDiagnosisConfigId
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) GetDataDiagnosisConfigName() *string {
	return s.DataDiagnosisConfigName
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) GetDataDiagnosisJobId() *string {
	return s.DataDiagnosisJobId
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) GetGmtStartTime() *string {
	return s.GmtStartTime
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) GetJobSource() *string {
	return s.JobSource
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) GetLogs() []*string {
	return s.Logs
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) GetStatus() *string {
	return s.Status
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) GetType() *string {
	return s.Type
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) SetBizDate(v string) *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs {
	s.BizDate = &v
	return s
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) SetConfig(v string) *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs {
	s.Config = &v
	return s
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) SetDataDiagnosisConfigId(v string) *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs {
	s.DataDiagnosisConfigId = &v
	return s
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) SetDataDiagnosisConfigName(v string) *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs {
	s.DataDiagnosisConfigName = &v
	return s
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) SetDataDiagnosisJobId(v string) *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs {
	s.DataDiagnosisJobId = &v
	return s
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) SetGmtCreateTime(v string) *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs {
	s.GmtCreateTime = &v
	return s
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) SetGmtStartTime(v string) *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs {
	s.GmtStartTime = &v
	return s
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) SetJobSource(v string) *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs {
	s.JobSource = &v
	return s
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) SetLogs(v []*string) *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs {
	s.Logs = v
	return s
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) SetStatus(v string) *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs {
	s.Status = &v
	return s
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) SetType(v string) *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs {
	s.Type = &v
	return s
}

func (s *ListDataDiagnosisJobsResponseBodyDataDiagnosisJobs) Validate() error {
	return dara.Validate(s)
}
