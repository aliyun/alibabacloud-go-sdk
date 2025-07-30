// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetJobEventsRequest
	GetEndTime() *string
	SetMaxEventsNum(v int32) *GetJobEventsRequest
	GetMaxEventsNum() *int32
	SetStartTime(v string) *GetJobEventsRequest
	GetStartTime() *string
}

type GetJobEventsRequest struct {
	// The end time (UTC) of the time range for querying events. The default value is the current time.
	//
	// example:
	//
	// 2020-11-08T18:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of events that can be returned. Default value: 2000.
	//
	// example:
	//
	// 100
	MaxEventsNum *int32 `json:"MaxEventsNum,omitempty" xml:"MaxEventsNum,omitempty"`
	// The start time (UTC) of the time range for querying events. The default value is 7 days ago.
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetJobEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobEventsRequest) GoString() string {
	return s.String()
}

func (s *GetJobEventsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetJobEventsRequest) GetMaxEventsNum() *int32 {
	return s.MaxEventsNum
}

func (s *GetJobEventsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetJobEventsRequest) SetEndTime(v string) *GetJobEventsRequest {
	s.EndTime = &v
	return s
}

func (s *GetJobEventsRequest) SetMaxEventsNum(v int32) *GetJobEventsRequest {
	s.MaxEventsNum = &v
	return s
}

func (s *GetJobEventsRequest) SetStartTime(v string) *GetJobEventsRequest {
	s.StartTime = &v
	return s
}

func (s *GetJobEventsRequest) Validate() error {
	return dara.Validate(s)
}
