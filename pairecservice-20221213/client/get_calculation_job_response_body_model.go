// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCalculationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetABMetricId(v string) *GetCalculationJobResponseBody
	GetABMetricId() *string
	SetABMetricName(v string) *GetCalculationJobResponseBody
	GetABMetricName() *string
	SetBizDate(v string) *GetCalculationJobResponseBody
	GetBizDate() *string
	SetConfig(v string) *GetCalculationJobResponseBody
	GetConfig() *string
	SetGmtRanTime(v string) *GetCalculationJobResponseBody
	GetGmtRanTime() *string
	SetJobMessage(v []*string) *GetCalculationJobResponseBody
	GetJobMessage() []*string
	SetJobSource(v string) *GetCalculationJobResponseBody
	GetJobSource() *string
	SetRequestId(v string) *GetCalculationJobResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetCalculationJobResponseBody
	GetStatus() *string
}

type GetCalculationJobResponseBody struct {
	// The AB metric ID.
	//
	// example:
	//
	// 1
	ABMetricId *string `json:"ABMetricId,omitempty" xml:"ABMetricId,omitempty"`
	// The AB metric name.
	//
	// example:
	//
	// pv
	ABMetricName *string `json:"ABMetricName,omitempty" xml:"ABMetricName,omitempty"`
	// The business date. The value is in YYYY-MM-DD format.
	//
	// example:
	//
	// 2021-12-15
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The task configuration.
	//
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The time when the task ran. The value is in ISO 8601 format and includes a time zone offset.
	//
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtRanTime *string `json:"GmtRanTime,omitempty" xml:"GmtRanTime,omitempty"`
	// The task information.
	JobMessage []*string `json:"JobMessage,omitempty" xml:"JobMessage,omitempty" type:"Repeated"`
	// The task source. Valid values:
	//
	// - CronOffline: a scheduled offline task.
	//
	// - DataRerun: a data rerun task.
	//
	// example:
	//
	// CronOffline
	JobSource *string `json:"JobSource,omitempty" xml:"JobSource,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7D59453C-48AA-5FC5-8848-2D373BD1A17F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the compute task. Valid values:
	//
	// - Success: The task succeeded.
	//
	// - Failure: The task failed.
	//
	// - Initializing: The task is being initialized.
	//
	// - Running: The task is running.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetCalculationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCalculationJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetCalculationJobResponseBody) GetABMetricId() *string {
	return s.ABMetricId
}

func (s *GetCalculationJobResponseBody) GetABMetricName() *string {
	return s.ABMetricName
}

func (s *GetCalculationJobResponseBody) GetBizDate() *string {
	return s.BizDate
}

func (s *GetCalculationJobResponseBody) GetConfig() *string {
	return s.Config
}

func (s *GetCalculationJobResponseBody) GetGmtRanTime() *string {
	return s.GmtRanTime
}

func (s *GetCalculationJobResponseBody) GetJobMessage() []*string {
	return s.JobMessage
}

func (s *GetCalculationJobResponseBody) GetJobSource() *string {
	return s.JobSource
}

func (s *GetCalculationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCalculationJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetCalculationJobResponseBody) SetABMetricId(v string) *GetCalculationJobResponseBody {
	s.ABMetricId = &v
	return s
}

func (s *GetCalculationJobResponseBody) SetABMetricName(v string) *GetCalculationJobResponseBody {
	s.ABMetricName = &v
	return s
}

func (s *GetCalculationJobResponseBody) SetBizDate(v string) *GetCalculationJobResponseBody {
	s.BizDate = &v
	return s
}

func (s *GetCalculationJobResponseBody) SetConfig(v string) *GetCalculationJobResponseBody {
	s.Config = &v
	return s
}

func (s *GetCalculationJobResponseBody) SetGmtRanTime(v string) *GetCalculationJobResponseBody {
	s.GmtRanTime = &v
	return s
}

func (s *GetCalculationJobResponseBody) SetJobMessage(v []*string) *GetCalculationJobResponseBody {
	s.JobMessage = v
	return s
}

func (s *GetCalculationJobResponseBody) SetJobSource(v string) *GetCalculationJobResponseBody {
	s.JobSource = &v
	return s
}

func (s *GetCalculationJobResponseBody) SetRequestId(v string) *GetCalculationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCalculationJobResponseBody) SetStatus(v string) *GetCalculationJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetCalculationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
