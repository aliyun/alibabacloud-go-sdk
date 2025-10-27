// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterHostSecuritySummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterHostEvent(v *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent) *DescribeClusterHostSecuritySummaryResponseBody
	GetClusterHostEvent() *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent
	SetRequestId(v string) *DescribeClusterHostSecuritySummaryResponseBody
	GetRequestId() *string
}

type DescribeClusterHostSecuritySummaryResponseBody struct {
	// The alert details of the hosts.
	ClusterHostEvent *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent `json:"ClusterHostEvent,omitempty" xml:"ClusterHostEvent,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterHostSecuritySummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterHostSecuritySummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterHostSecuritySummaryResponseBody) GetClusterHostEvent() *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent {
	return s.ClusterHostEvent
}

func (s *DescribeClusterHostSecuritySummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterHostSecuritySummaryResponseBody) SetClusterHostEvent(v *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent) *DescribeClusterHostSecuritySummaryResponseBody {
	s.ClusterHostEvent = v
	return s
}

func (s *DescribeClusterHostSecuritySummaryResponseBody) SetRequestId(v string) *DescribeClusterHostSecuritySummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterHostSecuritySummaryResponseBody) Validate() error {
	if s.ClusterHostEvent != nil {
		if err := s.ClusterHostEvent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent struct {
	// The alert details of the host.
	AlarmEvent []*DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventAlarmEvent `json:"AlarmEvent,omitempty" xml:"AlarmEvent,omitempty" type:"Repeated"`
	// The baseline details of the host.
	BaselineEvent []*DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventBaselineEvent `json:"BaselineEvent,omitempty" xml:"BaselineEvent,omitempty" type:"Repeated"`
	// The vulnerability details of the host.
	VulEvent []*DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventVulEvent `json:"VulEvent,omitempty" xml:"VulEvent,omitempty" type:"Repeated"`
}

func (s DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent) GoString() string {
	return s.String()
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent) GetAlarmEvent() []*DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventAlarmEvent {
	return s.AlarmEvent
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent) GetBaselineEvent() []*DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventBaselineEvent {
	return s.BaselineEvent
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent) GetVulEvent() []*DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventVulEvent {
	return s.VulEvent
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent) SetAlarmEvent(v []*DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventAlarmEvent) *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent {
	s.AlarmEvent = v
	return s
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent) SetBaselineEvent(v []*DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventBaselineEvent) *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent {
	s.BaselineEvent = v
	return s
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent) SetVulEvent(v []*DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventVulEvent) *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent {
	s.VulEvent = v
	return s
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEvent) Validate() error {
	if s.AlarmEvent != nil {
		for _, item := range s.AlarmEvent {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.BaselineEvent != nil {
		for _, item := range s.BaselineEvent {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VulEvent != nil {
		for _, item := range s.VulEvent {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventAlarmEvent struct {
	// The number of alerts.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The alert level. Valid values:
	//
	// 	- **serious**
	//
	// 	- **suspicious**
	//
	// 	- **remind**
	//
	// example:
	//
	// remind
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventAlarmEvent) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventAlarmEvent) GoString() string {
	return s.String()
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventAlarmEvent) GetCount() *int64 {
	return s.Count
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventAlarmEvent) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventAlarmEvent) SetCount(v int64) *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventAlarmEvent {
	s.Count = &v
	return s
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventAlarmEvent) SetRiskLevel(v string) *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventAlarmEvent {
	s.RiskLevel = &v
	return s
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventAlarmEvent) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventBaselineEvent struct {
	// The number of baselines.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The risk level of the baseline. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// medium
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventBaselineEvent) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventBaselineEvent) GoString() string {
	return s.String()
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventBaselineEvent) GetCount() *int64 {
	return s.Count
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventBaselineEvent) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventBaselineEvent) SetCount(v int64) *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventBaselineEvent {
	s.Count = &v
	return s
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventBaselineEvent) SetRiskLevel(v string) *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventBaselineEvent {
	s.RiskLevel = &v
	return s
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventBaselineEvent) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventVulEvent struct {
	// The number of vulnerabilities.
	//
	// example:
	//
	// 3
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The risk level of the vulnerability. Valid values:
	//
	// 	- **asap**: high. You must fix the vulnerability at the earliest opportunity.
	//
	// 	- **nntf**: medium. You can fix the vulnerability based on your business requirements.
	//
	// 	- **later**: low. You can ignore the vulnerability.
	//
	// example:
	//
	// later
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventVulEvent) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventVulEvent) GoString() string {
	return s.String()
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventVulEvent) GetCount() *int64 {
	return s.Count
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventVulEvent) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventVulEvent) SetCount(v int64) *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventVulEvent {
	s.Count = &v
	return s
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventVulEvent) SetRiskLevel(v string) *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventVulEvent {
	s.RiskLevel = &v
	return s
}

func (s *DescribeClusterHostSecuritySummaryResponseBodyClusterHostEventVulEvent) Validate() error {
	return dara.Validate(s)
}
