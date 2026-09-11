// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobMonitorRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDelayRuleTime(v int64) *CreateJobMonitorRuleRequest
	GetDelayRuleTime() *int64
	SetDtsJobId(v string) *CreateJobMonitorRuleRequest
	GetDtsJobId() *string
	SetNoticeValue(v int32) *CreateJobMonitorRuleRequest
	GetNoticeValue() *int32
	SetPeriod(v int32) *CreateJobMonitorRuleRequest
	GetPeriod() *int32
	SetPhone(v string) *CreateJobMonitorRuleRequest
	GetPhone() *string
	SetRegionId(v string) *CreateJobMonitorRuleRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateJobMonitorRuleRequest
	GetResourceGroupId() *string
	SetState(v string) *CreateJobMonitorRuleRequest
	GetState() *string
	SetTimes(v int32) *CreateJobMonitorRuleRequest
	GetTimes() *int32
	SetType(v string) *CreateJobMonitorRuleRequest
	GetType() *string
}

type CreateJobMonitorRuleRequest struct {
	// The threshold for triggering an alert.
	//
	// - If **Type*	- is set to **delay**, the unit is seconds and the value must be an integer. Set the threshold based on your business requirements. A value of 10 or greater is recommended to avoid alert fluctuations caused by network issues or database loads.
	//
	// - If **Type*	- is set to **full_timeout**, the unit is hours and the value must be an integer.
	//
	// > This parameter is required when **Type*	- is set to **delay*	- or **full_timeout*	- and **State*	- is set to **Y**.
	//
	// example:
	//
	// 11
	DelayRuleTime *int64 `json:"DelayRuleTime,omitempty" xml:"DelayRuleTime,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking task. You can call [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) to obtain the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i03e3zty16i****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The alert threshold.
	//
	// example:
	//
	// 2
	NoticeValue *int32 `json:"NoticeValue,omitempty" xml:"NoticeValue,omitempty"`
	// The statistical period of the incremental verification task. Unit: minutes.
	//
	// > Valid values: 1, 5, 10, and 30.
	//
	// example:
	//
	// 5
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The mobile phone numbers of alert contacts, separated by commas (,).
	//
	// >-  This parameter is supported only on the China site (aliyun.com) and only for the Chinese mainland mobile phone numbers. A maximum of 10 mobile phone numbers can be specified.
	//
	// - The international site does not support SMS-based alerting. You can only [set alert rules for DTS tasks through the CloudMonitor monitoring platform](https://help.aliyun.com/document_detail/175876.html).
	//
	// example:
	//
	// 1361234****,1371234****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The region in which the DTS instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to enable the alert rule. Valid values:
	//
	// - **Y**: Enable the alert rule.
	//
	// - **N**: Disable the alert rule.
	//
	// Default value: **Y**.
	//
	// example:
	//
	// Y
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The number of statistical periods for the incremental verification task.
	//
	// example:
	//
	// 2
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
	// The type of the alert metric. Valid values:
	//
	// - **delay**: the **Latency*	- metric.
	//
	// - **error**: the **Migration Status*	- metric.
	//
	// - **full_timeout**: the **Full Migration Duration*	- metric.
	//
	// Default value: **error**. This parameter must be manually specified.
	//
	// example:
	//
	// delay
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateJobMonitorRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobMonitorRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateJobMonitorRuleRequest) GetDelayRuleTime() *int64 {
	return s.DelayRuleTime
}

func (s *CreateJobMonitorRuleRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *CreateJobMonitorRuleRequest) GetNoticeValue() *int32 {
	return s.NoticeValue
}

func (s *CreateJobMonitorRuleRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateJobMonitorRuleRequest) GetPhone() *string {
	return s.Phone
}

func (s *CreateJobMonitorRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateJobMonitorRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateJobMonitorRuleRequest) GetState() *string {
	return s.State
}

func (s *CreateJobMonitorRuleRequest) GetTimes() *int32 {
	return s.Times
}

func (s *CreateJobMonitorRuleRequest) GetType() *string {
	return s.Type
}

func (s *CreateJobMonitorRuleRequest) SetDelayRuleTime(v int64) *CreateJobMonitorRuleRequest {
	s.DelayRuleTime = &v
	return s
}

func (s *CreateJobMonitorRuleRequest) SetDtsJobId(v string) *CreateJobMonitorRuleRequest {
	s.DtsJobId = &v
	return s
}

func (s *CreateJobMonitorRuleRequest) SetNoticeValue(v int32) *CreateJobMonitorRuleRequest {
	s.NoticeValue = &v
	return s
}

func (s *CreateJobMonitorRuleRequest) SetPeriod(v int32) *CreateJobMonitorRuleRequest {
	s.Period = &v
	return s
}

func (s *CreateJobMonitorRuleRequest) SetPhone(v string) *CreateJobMonitorRuleRequest {
	s.Phone = &v
	return s
}

func (s *CreateJobMonitorRuleRequest) SetRegionId(v string) *CreateJobMonitorRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateJobMonitorRuleRequest) SetResourceGroupId(v string) *CreateJobMonitorRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateJobMonitorRuleRequest) SetState(v string) *CreateJobMonitorRuleRequest {
	s.State = &v
	return s
}

func (s *CreateJobMonitorRuleRequest) SetTimes(v int32) *CreateJobMonitorRuleRequest {
	s.Times = &v
	return s
}

func (s *CreateJobMonitorRuleRequest) SetType(v string) *CreateJobMonitorRuleRequest {
	s.Type = &v
	return s
}

func (s *CreateJobMonitorRuleRequest) Validate() error {
	return dara.Validate(s)
}
