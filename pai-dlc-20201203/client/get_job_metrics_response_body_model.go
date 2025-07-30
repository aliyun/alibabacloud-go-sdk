// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetJobMetricsResponseBody
	GetJobId() *string
	SetPodMetrics(v []*PodMetric) *GetJobMetricsResponseBody
	GetPodMetrics() []*PodMetric
	SetRequestId(v string) *GetJobMetricsResponseBody
	GetRequestId() *string
}

type GetJobMetricsResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// dlc-20210126170216-*******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The monitoring metrics of the job.
	PodMetrics []*PodMetric `json:"PodMetrics,omitempty" xml:"PodMetrics,omitempty" type:"Repeated"`
	// The request ID. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobMetricsResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetJobMetricsResponseBody) GetPodMetrics() []*PodMetric {
	return s.PodMetrics
}

func (s *GetJobMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobMetricsResponseBody) SetJobId(v string) *GetJobMetricsResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobMetricsResponseBody) SetPodMetrics(v []*PodMetric) *GetJobMetricsResponseBody {
	s.PodMetrics = v
	return s
}

func (s *GetJobMetricsResponseBody) SetRequestId(v string) *GetJobMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobMetricsResponseBody) Validate() error {
	return dara.Validate(s)
}
