// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalculationJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCalculationJobs(v []*ListCalculationJobsResponseBodyCalculationJobs) *ListCalculationJobsResponseBody
	GetCalculationJobs() []*ListCalculationJobsResponseBodyCalculationJobs
	SetRequestId(v string) *ListCalculationJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListCalculationJobsResponseBody
	GetTotalCount() *int64
}

type ListCalculationJobsResponseBody struct {
	CalculationJobs []*ListCalculationJobsResponseBodyCalculationJobs `json:"CalculationJobs,omitempty" xml:"CalculationJobs,omitempty" type:"Repeated"`
	// example:
	//
	// F7AC05FF-EDE7-5C2B-B9AE-33D6DF4178BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCalculationJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCalculationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCalculationJobsResponseBody) GetCalculationJobs() []*ListCalculationJobsResponseBodyCalculationJobs {
	return s.CalculationJobs
}

func (s *ListCalculationJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCalculationJobsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCalculationJobsResponseBody) SetCalculationJobs(v []*ListCalculationJobsResponseBodyCalculationJobs) *ListCalculationJobsResponseBody {
	s.CalculationJobs = v
	return s
}

func (s *ListCalculationJobsResponseBody) SetRequestId(v string) *ListCalculationJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCalculationJobsResponseBody) SetTotalCount(v int64) *ListCalculationJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCalculationJobsResponseBody) Validate() error {
	if s.CalculationJobs != nil {
		for _, item := range s.CalculationJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCalculationJobsResponseBodyCalculationJobs struct {
	// example:
	//
	// pv
	ABMetricName *string `json:"ABMetricName,omitempty" xml:"ABMetricName,omitempty"`
	// example:
	//
	// 2021-12-15
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// example:
	//
	// 2
	CalculationJobId *string `json:"CalculationJobId,omitempty" xml:"CalculationJobId,omitempty"`
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtRanTime *string   `json:"GmtRanTime,omitempty" xml:"GmtRanTime,omitempty"`
	JobMessage []*string `json:"JobMessage,omitempty" xml:"JobMessage,omitempty" type:"Repeated"`
	// example:
	//
	// CronOffline
	JobSource *string `json:"JobSource,omitempty" xml:"JobSource,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCalculationJobsResponseBodyCalculationJobs) String() string {
	return dara.Prettify(s)
}

func (s ListCalculationJobsResponseBodyCalculationJobs) GoString() string {
	return s.String()
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) GetABMetricName() *string {
	return s.ABMetricName
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) GetBizDate() *string {
	return s.BizDate
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) GetCalculationJobId() *string {
	return s.CalculationJobId
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) GetConfig() *string {
	return s.Config
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) GetGmtRanTime() *string {
	return s.GmtRanTime
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) GetJobMessage() []*string {
	return s.JobMessage
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) GetJobSource() *string {
	return s.JobSource
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) GetStatus() *string {
	return s.Status
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) SetABMetricName(v string) *ListCalculationJobsResponseBodyCalculationJobs {
	s.ABMetricName = &v
	return s
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) SetBizDate(v string) *ListCalculationJobsResponseBodyCalculationJobs {
	s.BizDate = &v
	return s
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) SetCalculationJobId(v string) *ListCalculationJobsResponseBodyCalculationJobs {
	s.CalculationJobId = &v
	return s
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) SetConfig(v string) *ListCalculationJobsResponseBodyCalculationJobs {
	s.Config = &v
	return s
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) SetGmtRanTime(v string) *ListCalculationJobsResponseBodyCalculationJobs {
	s.GmtRanTime = &v
	return s
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) SetJobMessage(v []*string) *ListCalculationJobsResponseBodyCalculationJobs {
	s.JobMessage = v
	return s
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) SetJobSource(v string) *ListCalculationJobsResponseBodyCalculationJobs {
	s.JobSource = &v
	return s
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) SetStatus(v string) *ListCalculationJobsResponseBodyCalculationJobs {
	s.Status = &v
	return s
}

func (s *ListCalculationJobsResponseBodyCalculationJobs) Validate() error {
	return dara.Validate(s)
}
