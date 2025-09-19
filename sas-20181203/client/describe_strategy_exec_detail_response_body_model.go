// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStrategyExecDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeStrategyExecDetailResponseBody
	GetEndTime() *string
	SetFailCount(v int32) *DescribeStrategyExecDetailResponseBody
	GetFailCount() *int32
	SetFailedEcsList(v []*DescribeStrategyExecDetailResponseBodyFailedEcsList) *DescribeStrategyExecDetailResponseBody
	GetFailedEcsList() []*DescribeStrategyExecDetailResponseBodyFailedEcsList
	SetInProcessCount(v int32) *DescribeStrategyExecDetailResponseBody
	GetInProcessCount() *int32
	SetPercent(v string) *DescribeStrategyExecDetailResponseBody
	GetPercent() *string
	SetRequestId(v string) *DescribeStrategyExecDetailResponseBody
	GetRequestId() *string
	SetSource(v string) *DescribeStrategyExecDetailResponseBody
	GetSource() *string
	SetStartTime(v string) *DescribeStrategyExecDetailResponseBody
	GetStartTime() *string
	SetSuccessCount(v int32) *DescribeStrategyExecDetailResponseBody
	GetSuccessCount() *int32
}

type DescribeStrategyExecDetailResponseBody struct {
	// The time when the baseline check ends.
	//
	// example:
	//
	// 2021-12-05 00:12:46
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of check items that failed to pass the baseline check. This type of check item is considered risk items.
	//
	// example:
	//
	// 94
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The servers on which risk items were detected.
	FailedEcsList []*DescribeStrategyExecDetailResponseBodyFailedEcsList `json:"FailedEcsList,omitempty" xml:"FailedEcsList,omitempty" type:"Repeated"`
	// The number of tasks that are **running **based on the baseline check policy.
	//
	// example:
	//
	// 0
	InProcessCount *int32 `json:"InProcessCount,omitempty" xml:"InProcessCount,omitempty"`
	// The execution progress of the baseline check policy.
	//
	// example:
	//
	// 100%
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 01611D72-1E33-53F5-A9A5-C81B5561970F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the baseline check. Valid values:
	//
	// 	- **Schedule**: automatic check that periodically runs
	//
	// 	- **Manual**: intermediate check that is manually performed
	//
	// example:
	//
	// Manual
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The time when the baseline check starts.
	//
	// example:
	//
	// 2021-12-05 00:12:16
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The number of check items that **passed*	- the baseline check.
	//
	// example:
	//
	// 81
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s DescribeStrategyExecDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyExecDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStrategyExecDetailResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeStrategyExecDetailResponseBody) GetFailCount() *int32 {
	return s.FailCount
}

func (s *DescribeStrategyExecDetailResponseBody) GetFailedEcsList() []*DescribeStrategyExecDetailResponseBodyFailedEcsList {
	return s.FailedEcsList
}

func (s *DescribeStrategyExecDetailResponseBody) GetInProcessCount() *int32 {
	return s.InProcessCount
}

func (s *DescribeStrategyExecDetailResponseBody) GetPercent() *string {
	return s.Percent
}

func (s *DescribeStrategyExecDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStrategyExecDetailResponseBody) GetSource() *string {
	return s.Source
}

func (s *DescribeStrategyExecDetailResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeStrategyExecDetailResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *DescribeStrategyExecDetailResponseBody) SetEndTime(v string) *DescribeStrategyExecDetailResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetFailCount(v int32) *DescribeStrategyExecDetailResponseBody {
	s.FailCount = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetFailedEcsList(v []*DescribeStrategyExecDetailResponseBodyFailedEcsList) *DescribeStrategyExecDetailResponseBody {
	s.FailedEcsList = v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetInProcessCount(v int32) *DescribeStrategyExecDetailResponseBody {
	s.InProcessCount = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetPercent(v string) *DescribeStrategyExecDetailResponseBody {
	s.Percent = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetRequestId(v string) *DescribeStrategyExecDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetSource(v string) *DescribeStrategyExecDetailResponseBody {
	s.Source = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetStartTime(v string) *DescribeStrategyExecDetailResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetSuccessCount(v int32) *DescribeStrategyExecDetailResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeStrategyExecDetailResponseBodyFailedEcsList struct {
	// The IP address of the server on which the baseline check was performed.
	//
	// example:
	//
	// 115.28.XX.XX
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// worker-k8s-for-cs-cf5741fb95c3a45d0864f0e4544dc****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 115.28.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 172.31.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The failure cause for the check item.
	//
	// example:
	//
	// Detect timeout
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DescribeStrategyExecDetailResponseBodyFailedEcsList) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyExecDetailResponseBodyFailedEcsList) GoString() string {
	return s.String()
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) GetIP() *string {
	return s.IP
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) GetReason() *string {
	return s.Reason
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) SetIP(v string) *DescribeStrategyExecDetailResponseBodyFailedEcsList {
	s.IP = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) SetInstanceName(v string) *DescribeStrategyExecDetailResponseBodyFailedEcsList {
	s.InstanceName = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) SetInternetIp(v string) *DescribeStrategyExecDetailResponseBodyFailedEcsList {
	s.InternetIp = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) SetIntranetIp(v string) *DescribeStrategyExecDetailResponseBodyFailedEcsList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) SetReason(v string) *DescribeStrategyExecDetailResponseBodyFailedEcsList {
	s.Reason = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) Validate() error {
	return dara.Validate(s)
}
