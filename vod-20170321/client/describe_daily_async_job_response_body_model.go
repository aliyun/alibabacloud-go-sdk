// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDailyAsyncJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDailyJobList(v []*DescribeDailyAsyncJobResponseBodyDailyJobList) *DescribeDailyAsyncJobResponseBody
	GetDailyJobList() []*DescribeDailyAsyncJobResponseBodyDailyJobList
	SetJobType(v string) *DescribeDailyAsyncJobResponseBody
	GetJobType() *string
	SetRequestId(v string) *DescribeDailyAsyncJobResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDailyAsyncJobResponseBody
	GetTotalCount() *int32
}

type DescribeDailyAsyncJobResponseBody struct {
	DailyJobList []*DescribeDailyAsyncJobResponseBodyDailyJobList `json:"DailyJobList,omitempty" xml:"DailyJobList,omitempty" type:"Repeated"`
	JobType      *string                                          `json:"JobType,omitempty" xml:"JobType,omitempty"`
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDailyAsyncJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDailyAsyncJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDailyAsyncJobResponseBody) GetDailyJobList() []*DescribeDailyAsyncJobResponseBodyDailyJobList {
	return s.DailyJobList
}

func (s *DescribeDailyAsyncJobResponseBody) GetJobType() *string {
	return s.JobType
}

func (s *DescribeDailyAsyncJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDailyAsyncJobResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDailyAsyncJobResponseBody) SetDailyJobList(v []*DescribeDailyAsyncJobResponseBodyDailyJobList) *DescribeDailyAsyncJobResponseBody {
	s.DailyJobList = v
	return s
}

func (s *DescribeDailyAsyncJobResponseBody) SetJobType(v string) *DescribeDailyAsyncJobResponseBody {
	s.JobType = &v
	return s
}

func (s *DescribeDailyAsyncJobResponseBody) SetRequestId(v string) *DescribeDailyAsyncJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDailyAsyncJobResponseBody) SetTotalCount(v int32) *DescribeDailyAsyncJobResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDailyAsyncJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDailyAsyncJobResponseBodyDailyJobList struct {
	Date     *string `json:"Date,omitempty" xml:"Date,omitempty"`
	JobCount *int32  `json:"JobCount,omitempty" xml:"JobCount,omitempty"`
}

func (s DescribeDailyAsyncJobResponseBodyDailyJobList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDailyAsyncJobResponseBodyDailyJobList) GoString() string {
	return s.String()
}

func (s *DescribeDailyAsyncJobResponseBodyDailyJobList) GetDate() *string {
	return s.Date
}

func (s *DescribeDailyAsyncJobResponseBodyDailyJobList) GetJobCount() *int32 {
	return s.JobCount
}

func (s *DescribeDailyAsyncJobResponseBodyDailyJobList) SetDate(v string) *DescribeDailyAsyncJobResponseBodyDailyJobList {
	s.Date = &v
	return s
}

func (s *DescribeDailyAsyncJobResponseBodyDailyJobList) SetJobCount(v int32) *DescribeDailyAsyncJobResponseBodyDailyJobList {
	s.JobCount = &v
	return s
}

func (s *DescribeDailyAsyncJobResponseBodyDailyJobList) Validate() error {
	return dara.Validate(s)
}
