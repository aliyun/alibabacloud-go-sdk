// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDailyAsyncJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDailyAsyncJobRequest
	GetEndTime() *string
	SetJobState(v string) *DescribeDailyAsyncJobRequest
	GetJobState() *string
	SetJobType(v string) *DescribeDailyAsyncJobRequest
	GetJobType() *string
	SetStartTime(v string) *DescribeDailyAsyncJobRequest
	GetStartTime() *string
}

type DescribeDailyAsyncJobRequest struct {
	// This parameter is required.
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	JobState *string `json:"JobState,omitempty" xml:"JobState,omitempty"`
	// This parameter is required.
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDailyAsyncJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDailyAsyncJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeDailyAsyncJobRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDailyAsyncJobRequest) GetJobState() *string {
	return s.JobState
}

func (s *DescribeDailyAsyncJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *DescribeDailyAsyncJobRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDailyAsyncJobRequest) SetEndTime(v string) *DescribeDailyAsyncJobRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDailyAsyncJobRequest) SetJobState(v string) *DescribeDailyAsyncJobRequest {
	s.JobState = &v
	return s
}

func (s *DescribeDailyAsyncJobRequest) SetJobType(v string) *DescribeDailyAsyncJobRequest {
	s.JobType = &v
	return s
}

func (s *DescribeDailyAsyncJobRequest) SetStartTime(v string) *DescribeDailyAsyncJobRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDailyAsyncJobRequest) Validate() error {
	return dara.Validate(s)
}
