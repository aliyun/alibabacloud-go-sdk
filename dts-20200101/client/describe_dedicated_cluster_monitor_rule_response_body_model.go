// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedClusterMonitorRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCpuAlarmThreshold(v string) *DescribeDedicatedClusterMonitorRuleResponseBody
	GetCpuAlarmThreshold() *string
	SetDedicatedClusterId(v string) *DescribeDedicatedClusterMonitorRuleResponseBody
	GetDedicatedClusterId() *string
	SetDiskAlarmThreshold(v string) *DescribeDedicatedClusterMonitorRuleResponseBody
	GetDiskAlarmThreshold() *string
	SetDuAlarmThreshold(v string) *DescribeDedicatedClusterMonitorRuleResponseBody
	GetDuAlarmThreshold() *string
	SetErrCode(v string) *DescribeDedicatedClusterMonitorRuleResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDedicatedClusterMonitorRuleResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *DescribeDedicatedClusterMonitorRuleResponseBody
	GetHttpStatusCode() *string
	SetMemAlarmThreshold(v string) *DescribeDedicatedClusterMonitorRuleResponseBody
	GetMemAlarmThreshold() *string
	SetNoticeSwitch(v string) *DescribeDedicatedClusterMonitorRuleResponseBody
	GetNoticeSwitch() *string
	SetPhones(v string) *DescribeDedicatedClusterMonitorRuleResponseBody
	GetPhones() *string
	SetRequestId(v string) *DescribeDedicatedClusterMonitorRuleResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeDedicatedClusterMonitorRuleResponseBody
	GetSuccess() *string
}

type DescribeDedicatedClusterMonitorRuleResponseBody struct {
	// The CPU alert threshold, in percentage.
	//
	// example:
	//
	// 80
	CpuAlarmThreshold *string `json:"CpuAlarmThreshold,omitempty" xml:"CpuAlarmThreshold,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// dtsClustervcwn1oeyu5fx4yf
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// The disk alert threshold, in percentage.
	//
	// example:
	//
	// 80
	DiskAlarmThreshold *string `json:"DiskAlarmThreshold,omitempty" xml:"DiskAlarmThreshold,omitempty"`
	// The DU alert threshold, in percentage.
	//
	// example:
	//
	// 46
	DuAlarmThreshold *string `json:"DuAlarmThreshold,omitempty" xml:"DuAlarmThreshold,omitempty"`
	// The error code returned if the call fails.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The memory alert threshold, in percentage.
	//
	// example:
	//
	// 80
	MemAlarmThreshold *string `json:"MemAlarmThreshold,omitempty" xml:"MemAlarmThreshold,omitempty"`
	// The alert switch. Valid values:
	//
	// - **1**: enabled.
	//
	// - **0**: disabled.
	//
	// example:
	//
	// 1
	NoticeSwitch *string `json:"NoticeSwitch,omitempty" xml:"NoticeSwitch,omitempty"`
	// The phone numbers that receive alerts. If there are multiple phone numbers, they are separated by commas (,).
	//
	// example:
	//
	// 186****7653
	Phones *string `json:"Phones,omitempty" xml:"Phones,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 621BB4F8-3016-4FAA-8D5A-5D3163CC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDedicatedClusterMonitorRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedClusterMonitorRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) GetCpuAlarmThreshold() *string {
	return s.CpuAlarmThreshold
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) GetDiskAlarmThreshold() *string {
	return s.DiskAlarmThreshold
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) GetDuAlarmThreshold() *string {
	return s.DuAlarmThreshold
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) GetMemAlarmThreshold() *string {
	return s.MemAlarmThreshold
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) GetNoticeSwitch() *string {
	return s.NoticeSwitch
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) GetPhones() *string {
	return s.Phones
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) SetCpuAlarmThreshold(v string) *DescribeDedicatedClusterMonitorRuleResponseBody {
	s.CpuAlarmThreshold = &v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) SetDedicatedClusterId(v string) *DescribeDedicatedClusterMonitorRuleResponseBody {
	s.DedicatedClusterId = &v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) SetDiskAlarmThreshold(v string) *DescribeDedicatedClusterMonitorRuleResponseBody {
	s.DiskAlarmThreshold = &v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) SetDuAlarmThreshold(v string) *DescribeDedicatedClusterMonitorRuleResponseBody {
	s.DuAlarmThreshold = &v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) SetErrCode(v string) *DescribeDedicatedClusterMonitorRuleResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) SetErrMessage(v string) *DescribeDedicatedClusterMonitorRuleResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) SetHttpStatusCode(v string) *DescribeDedicatedClusterMonitorRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) SetMemAlarmThreshold(v string) *DescribeDedicatedClusterMonitorRuleResponseBody {
	s.MemAlarmThreshold = &v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) SetNoticeSwitch(v string) *DescribeDedicatedClusterMonitorRuleResponseBody {
	s.NoticeSwitch = &v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) SetPhones(v string) *DescribeDedicatedClusterMonitorRuleResponseBody {
	s.Phones = &v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) SetRequestId(v string) *DescribeDedicatedClusterMonitorRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) SetSuccess(v string) *DescribeDedicatedClusterMonitorRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
