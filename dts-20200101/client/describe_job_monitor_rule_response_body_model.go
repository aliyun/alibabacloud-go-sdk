// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobMonitorRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeJobMonitorRuleResponseBody
	GetCode() *string
	SetDtsJobId(v string) *DescribeJobMonitorRuleResponseBody
	GetDtsJobId() *string
	SetDynamicMessage(v string) *DescribeJobMonitorRuleResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeJobMonitorRuleResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeJobMonitorRuleResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeJobMonitorRuleResponseBody
	GetHttpStatusCode() *int32
	SetMonitorRules(v []*DescribeJobMonitorRuleResponseBodyMonitorRules) *DescribeJobMonitorRuleResponseBody
	GetMonitorRules() []*DescribeJobMonitorRuleResponseBodyMonitorRules
	SetRequestId(v string) *DescribeJobMonitorRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeJobMonitorRuleResponseBody
	GetSuccess() *bool
	SetTopics(v []*string) *DescribeJobMonitorRuleResponseBody
	GetTopics() []*string
}

type DescribeJobMonitorRuleResponseBody struct {
	// The error code. This parameter will be removed in the future.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking task.
	//
	// example:
	//
	// ta7w132u12h****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the **%s*	- variable in the **ErrMessage*	- parameter.
	//
	// >  If the specified **DtsJobId*	- parameter is invalid, **The Value of Input Parameter %s is not valid*	- is returned for **ErrMessage*	- and **DtsJobId*	- is returned for **DynamicMessage**.
	//
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// 403
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
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
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The monitoring rules of the DTS task.
	MonitorRules []*DescribeJobMonitorRuleResponseBodyMonitorRules `json:"MonitorRules,omitempty" xml:"MonitorRules,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0CA14388-DD89-4A7B-8CDD-884A10CE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- **true**: The call was successful.
	//
	// 	- **false**:The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The topics of all subtasks in the distributed change tracking task.
	Topics []*string `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
}

func (s DescribeJobMonitorRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobMonitorRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobMonitorRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeJobMonitorRuleResponseBody) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeJobMonitorRuleResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeJobMonitorRuleResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeJobMonitorRuleResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeJobMonitorRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeJobMonitorRuleResponseBody) GetMonitorRules() []*DescribeJobMonitorRuleResponseBodyMonitorRules {
	return s.MonitorRules
}

func (s *DescribeJobMonitorRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeJobMonitorRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeJobMonitorRuleResponseBody) GetTopics() []*string {
	return s.Topics
}

func (s *DescribeJobMonitorRuleResponseBody) SetCode(v string) *DescribeJobMonitorRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) SetDtsJobId(v string) *DescribeJobMonitorRuleResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) SetDynamicMessage(v string) *DescribeJobMonitorRuleResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) SetErrCode(v string) *DescribeJobMonitorRuleResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) SetErrMessage(v string) *DescribeJobMonitorRuleResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) SetHttpStatusCode(v int32) *DescribeJobMonitorRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) SetMonitorRules(v []*DescribeJobMonitorRuleResponseBodyMonitorRules) *DescribeJobMonitorRuleResponseBody {
	s.MonitorRules = v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) SetRequestId(v string) *DescribeJobMonitorRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) SetSuccess(v bool) *DescribeJobMonitorRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) SetTopics(v []*string) *DescribeJobMonitorRuleResponseBody {
	s.Topics = v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeJobMonitorRuleResponseBodyMonitorRules struct {
	// The threshold that triggers the alert.
	//
	// 	- If the request parameter **Type*	- of the [CreateJobMonitorRule](https://help.aliyun.com/document_detail/212332.html) operation is set to **delay**, the unit of DelayRuleTime is seconds.
	//
	// 	- If the request parameter **Type*	- of the [CreateJobMonitorRule](https://help.aliyun.com/document_detail/212332.html) operation is set to **full_timeout**, the unit of DelayRuleTime is hours.
	//
	// example:
	//
	// 11
	DelayRuleTime *int64 `json:"DelayRuleTime,omitempty" xml:"DelayRuleTime,omitempty"`
	// Task ID.
	//
	// example:
	//
	// bi6e22ay243****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The task type of the DTS instance, with values: - **normal**: Migration or synchronization task. - **full_check**: Associated full check task. - **etl_check**: Associated incremental check task.
	//
	// example:
	//
	// normal
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// Alarm threshold.
	//
	// example:
	//
	// 2
	NoticeValue *int32 `json:"NoticeValue,omitempty" xml:"NoticeValue,omitempty"`
	// The statistical period for incremental validation tasks, in minutes.
	//
	// > Currently supported values are 1 minute, 5 minutes, 10 minutes, and 30 minutes.
	//
	// example:
	//
	// 5
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The mobile phone numbers that receive alert notifications. Multiple mobile numbers are separated by commas (,).
	//
	// example:
	//
	// 1361234****,1371234****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// Indicates whether the monitoring rule is enabled. Valid values:
	//
	// 	- **Y**: The monitoring rule is enabled.
	//
	// 	- **N**: The monitoring rule is disabled.
	//
	// example:
	//
	// Y
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The number of cycles for the incremental validation task.
	//
	// example:
	//
	// 2
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
	// The type of the monitoring rule. Valid values:
	//
	// 	- **delay**: If the task latency reaches the threshold, an alert is triggered.
	//
	// 	- **error**: If an exception occurs, an alert is triggered.
	//
	// example:
	//
	// delay
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeJobMonitorRuleResponseBodyMonitorRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobMonitorRuleResponseBodyMonitorRules) GoString() string {
	return s.String()
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) GetDelayRuleTime() *int64 {
	return s.DelayRuleTime
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) GetJobType() *string {
	return s.JobType
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) GetNoticeValue() *int32 {
	return s.NoticeValue
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) GetPhone() *string {
	return s.Phone
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) GetState() *string {
	return s.State
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) GetTimes() *int32 {
	return s.Times
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) GetType() *string {
	return s.Type
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) SetDelayRuleTime(v int64) *DescribeJobMonitorRuleResponseBodyMonitorRules {
	s.DelayRuleTime = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) SetJobId(v string) *DescribeJobMonitorRuleResponseBodyMonitorRules {
	s.JobId = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) SetJobType(v string) *DescribeJobMonitorRuleResponseBodyMonitorRules {
	s.JobType = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) SetNoticeValue(v int32) *DescribeJobMonitorRuleResponseBodyMonitorRules {
	s.NoticeValue = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) SetPeriod(v int32) *DescribeJobMonitorRuleResponseBodyMonitorRules {
	s.Period = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) SetPhone(v string) *DescribeJobMonitorRuleResponseBodyMonitorRules {
	s.Phone = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) SetState(v string) *DescribeJobMonitorRuleResponseBodyMonitorRules {
	s.State = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) SetTimes(v int32) *DescribeJobMonitorRuleResponseBodyMonitorRules {
	s.Times = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) SetType(v string) *DescribeJobMonitorRuleResponseBodyMonitorRules {
	s.Type = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) Validate() error {
	return dara.Validate(s)
}
