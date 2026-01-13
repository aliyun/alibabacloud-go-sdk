// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureConsistencyCheckJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureConsistencyCheckJobs(v []*ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) *ListFeatureConsistencyCheckJobsResponseBody
	GetFeatureConsistencyCheckJobs() []*ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs
	SetRequestId(v string) *ListFeatureConsistencyCheckJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListFeatureConsistencyCheckJobsResponseBody
	GetTotalCount() *string
}

type ListFeatureConsistencyCheckJobsResponseBody struct {
	FeatureConsistencyCheckJobs []*ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs `json:"FeatureConsistencyCheckJobs,omitempty" xml:"FeatureConsistencyCheckJobs,omitempty" type:"Repeated"`
	// example:
	//
	// A04CB8C0-E74A-5E83-BC61-64D153574EC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFeatureConsistencyCheckJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobsResponseBody) GetFeatureConsistencyCheckJobs() []*ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs {
	return s.FeatureConsistencyCheckJobs
}

func (s *ListFeatureConsistencyCheckJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFeatureConsistencyCheckJobsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListFeatureConsistencyCheckJobsResponseBody) SetFeatureConsistencyCheckJobs(v []*ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) *ListFeatureConsistencyCheckJobsResponseBody {
	s.FeatureConsistencyCheckJobs = v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBody) SetRequestId(v string) *ListFeatureConsistencyCheckJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBody) SetTotalCount(v string) *ListFeatureConsistencyCheckJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBody) Validate() error {
	if s.FeatureConsistencyCheckJobs != nil {
		for _, item := range s.FeatureConsistencyCheckJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs struct {
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 5
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	// example:
	//
	// feature_consistency_check_1
	FeatureConsistencyCheckJobConfigName *string `json:"FeatureConsistencyCheckJobConfigName,omitempty" xml:"FeatureConsistencyCheckJobConfigName,omitempty"`
	// example:
	//
	// 4
	FeatureConsistencyCheckJobId *string `json:"FeatureConsistencyCheckJobId,omitempty" xml:"FeatureConsistencyCheckJobId,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtEndTime *string `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtStartTime *string   `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	Logs         []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) GetConfig() *string {
	return s.Config
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) GetFeatureConsistencyCheckJobConfigId() *string {
	return s.FeatureConsistencyCheckJobConfigId
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) GetFeatureConsistencyCheckJobConfigName() *string {
	return s.FeatureConsistencyCheckJobConfigName
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) GetFeatureConsistencyCheckJobId() *string {
	return s.FeatureConsistencyCheckJobId
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) GetGmtEndTime() *string {
	return s.GmtEndTime
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) GetGmtStartTime() *string {
	return s.GmtStartTime
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) GetLogs() []*string {
	return s.Logs
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) GetStatus() *string {
	return s.Status
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) SetConfig(v string) *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs {
	s.Config = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) SetFeatureConsistencyCheckJobConfigId(v string) *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs {
	s.FeatureConsistencyCheckJobConfigId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) SetFeatureConsistencyCheckJobConfigName(v string) *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs {
	s.FeatureConsistencyCheckJobConfigName = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) SetFeatureConsistencyCheckJobId(v string) *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs {
	s.FeatureConsistencyCheckJobId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) SetGmtEndTime(v string) *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs {
	s.GmtEndTime = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) SetGmtStartTime(v string) *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs {
	s.GmtStartTime = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) SetLogs(v []*string) *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs {
	s.Logs = v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) SetStatus(v string) *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs {
	s.Status = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponseBodyFeatureConsistencyCheckJobs) Validate() error {
	return dara.Validate(s)
}
