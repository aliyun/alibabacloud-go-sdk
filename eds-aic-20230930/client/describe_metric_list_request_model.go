// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIds(v []*string) *DescribeMetricListRequest
	GetAndroidInstanceIds() []*string
	SetEndTime(v string) *DescribeMetricListRequest
	GetEndTime() *string
	SetInstanceIds(v []*string) *DescribeMetricListRequest
	GetInstanceIds() []*string
	SetLength(v string) *DescribeMetricListRequest
	GetLength() *string
	SetMetricNames(v []*string) *DescribeMetricListRequest
	GetMetricNames() []*string
	SetNextToken(v string) *DescribeMetricListRequest
	GetNextToken() *string
	SetPeriod(v int32) *DescribeMetricListRequest
	GetPeriod() *int32
	SetProcessInfos(v []*DescribeMetricListRequestProcessInfos) *DescribeMetricListRequest
	GetProcessInfos() []*DescribeMetricListRequestProcessInfos
	SetStartTime(v string) *DescribeMetricListRequest
	GetStartTime() *string
}

type DescribeMetricListRequest struct {
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2019-01-31 11:00:00
	EndTime     *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1000
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// This parameter is required.
	MetricNames []*string `json:"MetricNames,omitempty" xml:"MetricNames,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kw9dGL5jves2FS9RLq****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 60
	Period       *int32                                   `json:"Period,omitempty" xml:"Period,omitempty"`
	ProcessInfos []*DescribeMetricListRequestProcessInfos `json:"ProcessInfos,omitempty" xml:"ProcessInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 2019-01-31 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMetricListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricListRequest) GetAndroidInstanceIds() []*string {
	return s.AndroidInstanceIds
}

func (s *DescribeMetricListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMetricListRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeMetricListRequest) GetLength() *string {
	return s.Length
}

func (s *DescribeMetricListRequest) GetMetricNames() []*string {
	return s.MetricNames
}

func (s *DescribeMetricListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMetricListRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeMetricListRequest) GetProcessInfos() []*DescribeMetricListRequestProcessInfos {
	return s.ProcessInfos
}

func (s *DescribeMetricListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMetricListRequest) SetAndroidInstanceIds(v []*string) *DescribeMetricListRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *DescribeMetricListRequest) SetEndTime(v string) *DescribeMetricListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMetricListRequest) SetInstanceIds(v []*string) *DescribeMetricListRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeMetricListRequest) SetLength(v string) *DescribeMetricListRequest {
	s.Length = &v
	return s
}

func (s *DescribeMetricListRequest) SetMetricNames(v []*string) *DescribeMetricListRequest {
	s.MetricNames = v
	return s
}

func (s *DescribeMetricListRequest) SetNextToken(v string) *DescribeMetricListRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeMetricListRequest) SetPeriod(v int32) *DescribeMetricListRequest {
	s.Period = &v
	return s
}

func (s *DescribeMetricListRequest) SetProcessInfos(v []*DescribeMetricListRequestProcessInfos) *DescribeMetricListRequest {
	s.ProcessInfos = v
	return s
}

func (s *DescribeMetricListRequest) SetStartTime(v string) *DescribeMetricListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMetricListRequest) Validate() error {
	if s.ProcessInfos != nil {
		for _, item := range s.ProcessInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricListRequestProcessInfos struct {
	ProcessIds []*int32 `json:"ProcessIds,omitempty" xml:"ProcessIds,omitempty" type:"Repeated"`
	// example:
	//
	// php-fpm
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
}

func (s DescribeMetricListRequestProcessInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricListRequestProcessInfos) GoString() string {
	return s.String()
}

func (s *DescribeMetricListRequestProcessInfos) GetProcessIds() []*int32 {
	return s.ProcessIds
}

func (s *DescribeMetricListRequestProcessInfos) GetProcessName() *string {
	return s.ProcessName
}

func (s *DescribeMetricListRequestProcessInfos) SetProcessIds(v []*int32) *DescribeMetricListRequestProcessInfos {
	s.ProcessIds = v
	return s
}

func (s *DescribeMetricListRequestProcessInfos) SetProcessName(v string) *DescribeMetricListRequestProcessInfos {
	s.ProcessName = &v
	return s
}

func (s *DescribeMetricListRequestProcessInfos) Validate() error {
	return dara.Validate(s)
}
