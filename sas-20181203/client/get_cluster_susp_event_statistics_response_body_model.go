// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterSuspEventStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetClusterSuspEventStatisticsResponseBody
	GetRequestId() *string
	SetSuspStatistics(v *GetClusterSuspEventStatisticsResponseBodySuspStatistics) *GetClusterSuspEventStatisticsResponseBody
	GetSuspStatistics() *GetClusterSuspEventStatisticsResponseBodySuspStatistics
}

type GetClusterSuspEventStatisticsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// ACF97412-FD09-4D1F-994F-34DF12BR****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of alerts by risk level.
	SuspStatistics *GetClusterSuspEventStatisticsResponseBodySuspStatistics `json:"SuspStatistics,omitempty" xml:"SuspStatistics,omitempty" type:"Struct"`
}

func (s GetClusterSuspEventStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterSuspEventStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterSuspEventStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterSuspEventStatisticsResponseBody) GetSuspStatistics() *GetClusterSuspEventStatisticsResponseBodySuspStatistics {
	return s.SuspStatistics
}

func (s *GetClusterSuspEventStatisticsResponseBody) SetRequestId(v string) *GetClusterSuspEventStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterSuspEventStatisticsResponseBody) SetSuspStatistics(v *GetClusterSuspEventStatisticsResponseBodySuspStatistics) *GetClusterSuspEventStatisticsResponseBody {
	s.SuspStatistics = v
	return s
}

func (s *GetClusterSuspEventStatisticsResponseBody) Validate() error {
	if s.SuspStatistics != nil {
		if err := s.SuspStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetClusterSuspEventStatisticsResponseBodySuspStatistics struct {
	// The number of alerts whose Emergency level is Reminder.
	//
	// example:
	//
	// 1
	Remind *int32 `json:"Remind,omitempty" xml:"Remind,omitempty"`
	// The number of alerts whose Emergency level is Urgent.
	//
	// example:
	//
	// 1
	Serious *int32 `json:"Serious,omitempty" xml:"Serious,omitempty"`
	// The number of alerts whose Emergency level is Suspicious.
	//
	// example:
	//
	// 2
	Suspicious *int32 `json:"Suspicious,omitempty" xml:"Suspicious,omitempty"`
}

func (s GetClusterSuspEventStatisticsResponseBodySuspStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetClusterSuspEventStatisticsResponseBodySuspStatistics) GoString() string {
	return s.String()
}

func (s *GetClusterSuspEventStatisticsResponseBodySuspStatistics) GetRemind() *int32 {
	return s.Remind
}

func (s *GetClusterSuspEventStatisticsResponseBodySuspStatistics) GetSerious() *int32 {
	return s.Serious
}

func (s *GetClusterSuspEventStatisticsResponseBodySuspStatistics) GetSuspicious() *int32 {
	return s.Suspicious
}

func (s *GetClusterSuspEventStatisticsResponseBodySuspStatistics) SetRemind(v int32) *GetClusterSuspEventStatisticsResponseBodySuspStatistics {
	s.Remind = &v
	return s
}

func (s *GetClusterSuspEventStatisticsResponseBodySuspStatistics) SetSerious(v int32) *GetClusterSuspEventStatisticsResponseBodySuspStatistics {
	s.Serious = &v
	return s
}

func (s *GetClusterSuspEventStatisticsResponseBodySuspStatistics) SetSuspicious(v int32) *GetClusterSuspEventStatisticsResponseBodySuspStatistics {
	s.Suspicious = &v
	return s
}

func (s *GetClusterSuspEventStatisticsResponseBodySuspStatistics) Validate() error {
	return dara.Validate(s)
}
