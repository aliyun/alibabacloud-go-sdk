// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensions(v string) *GetMetricsRequest
	GetDimensions() *string
	SetEndTime(v string) *GetMetricsRequest
	GetEndTime() *string
	SetJobId(v string) *GetMetricsRequest
	GetJobId() *string
	SetLength(v string) *GetMetricsRequest
	GetLength() *string
	SetMetricName(v string) *GetMetricsRequest
	GetMetricName() *string
	SetNamespace(v string) *GetMetricsRequest
	GetNamespace() *string
	SetNextToken(v string) *GetMetricsRequest
	GetNextToken() *string
	SetPeriod(v string) *GetMetricsRequest
	GetPeriod() *string
	SetStartTime(v string) *GetMetricsRequest
	GetStartTime() *string
	SetToken(v string) *GetMetricsRequest
	GetToken() *string
}

type GetMetricsRequest struct {
	// example:
	//
	// [{\\"jobId\\":\\"dlcdpfpc96mh63mg\\",\\"pod\\":\\"dlcdpfpc96mh63mg-worker-748\\",\\"regionId\\":\\"cn-wulanchabu\\",\\"userId\\":\\"1458867964644701\\"}]
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// example:
	//
	// 2020-11-09T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// dlckjd5hm84tmjec
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 5000
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// example:
	//
	// JOB_MEMORY_FREE
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// example:
	//
	// acs_pai_dlc
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 2c6b65b6f9d625d4716568ca19b2064be0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 5
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// eyXXXX-XXXX.XXXXX
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetMetricsRequest) GetDimensions() *string {
	return s.Dimensions
}

func (s *GetMetricsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetMetricsRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetMetricsRequest) GetLength() *string {
	return s.Length
}

func (s *GetMetricsRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *GetMetricsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetMetricsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetMetricsRequest) GetPeriod() *string {
	return s.Period
}

func (s *GetMetricsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMetricsRequest) GetToken() *string {
	return s.Token
}

func (s *GetMetricsRequest) SetDimensions(v string) *GetMetricsRequest {
	s.Dimensions = &v
	return s
}

func (s *GetMetricsRequest) SetEndTime(v string) *GetMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetMetricsRequest) SetJobId(v string) *GetMetricsRequest {
	s.JobId = &v
	return s
}

func (s *GetMetricsRequest) SetLength(v string) *GetMetricsRequest {
	s.Length = &v
	return s
}

func (s *GetMetricsRequest) SetMetricName(v string) *GetMetricsRequest {
	s.MetricName = &v
	return s
}

func (s *GetMetricsRequest) SetNamespace(v string) *GetMetricsRequest {
	s.Namespace = &v
	return s
}

func (s *GetMetricsRequest) SetNextToken(v string) *GetMetricsRequest {
	s.NextToken = &v
	return s
}

func (s *GetMetricsRequest) SetPeriod(v string) *GetMetricsRequest {
	s.Period = &v
	return s
}

func (s *GetMetricsRequest) SetStartTime(v string) *GetMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *GetMetricsRequest) SetToken(v string) *GetMetricsRequest {
	s.Token = &v
	return s
}

func (s *GetMetricsRequest) Validate() error {
	return dara.Validate(s)
}
