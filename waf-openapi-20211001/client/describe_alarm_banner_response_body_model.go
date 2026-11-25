// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlarmBannerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBannerStatus(v *DescribeAlarmBannerResponseBodyBannerStatus) *DescribeAlarmBannerResponseBody
	GetBannerStatus() *DescribeAlarmBannerResponseBodyBannerStatus
	SetRequestId(v string) *DescribeAlarmBannerResponseBody
	GetRequestId() *string
}

type DescribeAlarmBannerResponseBody struct {
	// The alert banner status information.
	BannerStatus *DescribeAlarmBannerResponseBodyBannerStatus `json:"BannerStatus,omitempty" xml:"BannerStatus,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 5555DC36-0CF2-5AA3-B1C7-D6BD8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAlarmBannerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmBannerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlarmBannerResponseBody) GetBannerStatus() *DescribeAlarmBannerResponseBodyBannerStatus {
	return s.BannerStatus
}

func (s *DescribeAlarmBannerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlarmBannerResponseBody) SetBannerStatus(v *DescribeAlarmBannerResponseBodyBannerStatus) *DescribeAlarmBannerResponseBody {
	s.BannerStatus = v
	return s
}

func (s *DescribeAlarmBannerResponseBody) SetRequestId(v string) *DescribeAlarmBannerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlarmBannerResponseBody) Validate() error {
	if s.BannerStatus != nil {
		if err := s.BannerStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAlarmBannerResponseBodyBannerStatus struct {
	// The cause of the alerting. If the Alarm Metric is sandbox (**Type*	- is set to **sandbox**), valid values are:
	//
	// - **fivefold**: The clean queries per second (QPS) exceeds the current specifications.
	//
	// - **4count**: The accumulated actual clean QPS exceeds the current specifications.
	//
	// - **exceed10w**: The actual peak clean QPS exceeds 100,000.
	//
	// - **costProtection**: Billing protection is triggered.
	//
	// example:
	//
	// 4count
	Cause *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	// The usage when alerting is triggered.
	//
	// - If the Alarm Metric is sandbox (**Type*	- is set to **sandbox**), this parameter indicates the number of days the quota is exceeded.
	//
	// example:
	//
	// 9008
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The alerting status. Valid values:
	//
	// - **true**: Alerting is triggered. If the Alarm Metric is sandbox (**Type*	- is set to **sandbox**), the instance is in the sandbox.
	//
	// - **false**: No alerting is triggered. If the Alarm Metric is sandbox (**Type*	- is set to **sandbox**), the instance is not in the sandbox.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// The Alarm Metric. Valid values:
	//
	// - **sandbox**: sandbox alerting.
	//
	// example:
	//
	// sandbox
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAlarmBannerResponseBodyBannerStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmBannerResponseBodyBannerStatus) GoString() string {
	return s.String()
}

func (s *DescribeAlarmBannerResponseBodyBannerStatus) GetCause() *string {
	return s.Cause
}

func (s *DescribeAlarmBannerResponseBodyBannerStatus) GetCount() *int32 {
	return s.Count
}

func (s *DescribeAlarmBannerResponseBodyBannerStatus) GetStatus() *bool {
	return s.Status
}

func (s *DescribeAlarmBannerResponseBodyBannerStatus) GetType() *string {
	return s.Type
}

func (s *DescribeAlarmBannerResponseBodyBannerStatus) SetCause(v string) *DescribeAlarmBannerResponseBodyBannerStatus {
	s.Cause = &v
	return s
}

func (s *DescribeAlarmBannerResponseBodyBannerStatus) SetCount(v int32) *DescribeAlarmBannerResponseBodyBannerStatus {
	s.Count = &v
	return s
}

func (s *DescribeAlarmBannerResponseBodyBannerStatus) SetStatus(v bool) *DescribeAlarmBannerResponseBodyBannerStatus {
	s.Status = &v
	return s
}

func (s *DescribeAlarmBannerResponseBodyBannerStatus) SetType(v string) *DescribeAlarmBannerResponseBodyBannerStatus {
	s.Type = &v
	return s
}

func (s *DescribeAlarmBannerResponseBodyBannerStatus) Validate() error {
	return dara.Validate(s)
}
