// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRequestHitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRequestHitResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeRequestHitResponseBodyResultObject) *DescribeRequestHitResponseBody
	GetResultObject() *DescribeRequestHitResponseBodyResultObject
}

type DescribeRequestHitResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	ResultObject *DescribeRequestHitResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeRequestHitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRequestHitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRequestHitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRequestHitResponseBody) GetResultObject() *DescribeRequestHitResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeRequestHitResponseBody) SetRequestId(v string) *DescribeRequestHitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRequestHitResponseBody) SetResultObject(v *DescribeRequestHitResponseBodyResultObject) *DescribeRequestHitResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeRequestHitResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRequestHitResponseBodyResultObject struct {
	// Event code
	//
	// example:
	//
	// de_aszbjb7236
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册风险
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Input parameters.
	//
	// example:
	//
	// {\\"eventCode\\":\\"de_afghcf6411\\",\\"ip\\":\\"196.168.0.1\\",\\"DEtest222\\":9007199254740999,\\"age\\":20}
	Inputs *string `json:"inputs,omitempty" xml:"inputs,omitempty"`
	// Output parameters
	//
	// example:
	//
	// {\\"tags\\":\\"rm0102,test_tag,age\\",\\"score\\":\\"30.0\\",\\"extend\\":\\"{\\\\\\"OUT_V01\\\\\\":\\\\\\"Maritime\\\\\\",\\\\\\"OUT_V02\\\\\\":\\\\\\"Lome\\\\\\",\\\\\\"OUT_V03\\\\\\":\\\\\\"196.168.0.1_A\\\\\\"}\\",\\"finalDecision\\":\\"REJECT\\"}
	Outputs *string `json:"outputs,omitempty" xml:"outputs,omitempty"`
	// Timestamp of the request.
	//
	// example:
	//
	// 1752571330000
	RequestTime *int64 `json:"requestTime,omitempty" xml:"requestTime,omitempty"`
	// Details of the executed rules.
	RuleHitRecords []*DescribeRequestHitResponseBodyResultObjectRuleHitRecords `json:"ruleHitRecords,omitempty" xml:"ruleHitRecords,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// 60C97040-D5D5-4906-9522-B9B413730CAA
	SRequestId *string `json:"sRequestId,omitempty" xml:"sRequestId,omitempty"`
	// Total amount of the request
	//
	// example:
	//
	// 4
	TotalCost *int64 `json:"totalCost,omitempty" xml:"totalCost,omitempty"`
}

func (s DescribeRequestHitResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeRequestHitResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeRequestHitResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeRequestHitResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeRequestHitResponseBodyResultObject) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeRequestHitResponseBodyResultObject) GetOutputs() *string {
	return s.Outputs
}

func (s *DescribeRequestHitResponseBodyResultObject) GetRequestTime() *int64 {
	return s.RequestTime
}

func (s *DescribeRequestHitResponseBodyResultObject) GetRuleHitRecords() []*DescribeRequestHitResponseBodyResultObjectRuleHitRecords {
	return s.RuleHitRecords
}

func (s *DescribeRequestHitResponseBodyResultObject) GetSRequestId() *string {
	return s.SRequestId
}

func (s *DescribeRequestHitResponseBodyResultObject) GetTotalCost() *int64 {
	return s.TotalCost
}

func (s *DescribeRequestHitResponseBodyResultObject) SetEventCode(v string) *DescribeRequestHitResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeRequestHitResponseBodyResultObject) SetEventName(v string) *DescribeRequestHitResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeRequestHitResponseBodyResultObject) SetInputs(v string) *DescribeRequestHitResponseBodyResultObject {
	s.Inputs = &v
	return s
}

func (s *DescribeRequestHitResponseBodyResultObject) SetOutputs(v string) *DescribeRequestHitResponseBodyResultObject {
	s.Outputs = &v
	return s
}

func (s *DescribeRequestHitResponseBodyResultObject) SetRequestTime(v int64) *DescribeRequestHitResponseBodyResultObject {
	s.RequestTime = &v
	return s
}

func (s *DescribeRequestHitResponseBodyResultObject) SetRuleHitRecords(v []*DescribeRequestHitResponseBodyResultObjectRuleHitRecords) *DescribeRequestHitResponseBodyResultObject {
	s.RuleHitRecords = v
	return s
}

func (s *DescribeRequestHitResponseBodyResultObject) SetSRequestId(v string) *DescribeRequestHitResponseBodyResultObject {
	s.SRequestId = &v
	return s
}

func (s *DescribeRequestHitResponseBodyResultObject) SetTotalCost(v int64) *DescribeRequestHitResponseBodyResultObject {
	s.TotalCost = &v
	return s
}

func (s *DescribeRequestHitResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeRequestHitResponseBodyResultObjectRuleHitRecords struct {
	// Duration
	//
	// example:
	//
	// 1
	Cost *int32 `json:"cost,omitempty" xml:"cost,omitempty"`
	// Whether the rule was hit.
	//
	// example:
	//
	// true
	HitSuccessful *bool `json:"hitSuccessful,omitempty" xml:"hitSuccessful,omitempty"`
	// Whether to show details
	//
	// example:
	//
	// true
	IsShowDetail *bool `json:"isShowDetail,omitempty" xml:"isShowDetail,omitempty"`
	// Order.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// Policy ID
	//
	// example:
	//
	// 101544
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Policy name
	//
	// example:
	//
	// 注册手机号是11位数字
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Rule snapshot ID
	//
	// example:
	//
	// 27
	RuleSnapshotId *string `json:"ruleSnapshotId,omitempty" xml:"ruleSnapshotId,omitempty"`
	// Policy status
	//
	// example:
	//
	// RUNNING
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
}

func (s DescribeRequestHitResponseBodyResultObjectRuleHitRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeRequestHitResponseBodyResultObjectRuleHitRecords) GoString() string {
	return s.String()
}

func (s *DescribeRequestHitResponseBodyResultObjectRuleHitRecords) GetCost() *int32 {
	return s.Cost
}

func (s *DescribeRequestHitResponseBodyResultObjectRuleHitRecords) GetHitSuccessful() *bool {
	return s.HitSuccessful
}

func (s *DescribeRequestHitResponseBodyResultObjectRuleHitRecords) GetIsShowDetail() *bool {
	return s.IsShowDetail
}

func (s *DescribeRequestHitResponseBodyResultObjectRuleHitRecords) GetOrder() *int32 {
	return s.Order
}

func (s *DescribeRequestHitResponseBodyResultObjectRuleHitRecords) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRequestHitResponseBodyResultObjectRuleHitRecords) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeRequestHitResponseBodyResultObjectRuleHitRecords) GetRuleSnapshotId() *string {
	return s.RuleSnapshotId
}

func (s *DescribeRequestHitResponseBodyResultObjectRuleHitRecords) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *DescribeRequestHitResponseBodyResultObjectRuleHitRecords) SetCost(v int32) *DescribeRequestHitResponseBodyResultObjectRuleHitRecords {
	s.Cost = &v
	return s
}

func (s *DescribeRequestHitResponseBodyResultObjectRuleHitRecords) SetHitSuccessful(v bool) *DescribeRequestHitResponseBodyResultObjectRuleHitRecords {
	s.HitSuccessful = &v
	return s
}

func (s *DescribeRequestHitResponseBodyResultObjectRuleHitRecords) SetIsShowDetail(v bool) *DescribeRequestHitResponseBodyResultObjectRuleHitRecords {
	s.IsShowDetail = &v
	return s
}

func (s *DescribeRequestHitResponseBodyResultObjectRuleHitRecords) SetOrder(v int32) *DescribeRequestHitResponseBodyResultObjectRuleHitRecords {
	s.Order = &v
	return s
}

func (s *DescribeRequestHitResponseBodyResultObjectRuleHitRecords) SetRuleId(v string) *DescribeRequestHitResponseBodyResultObjectRuleHitRecords {
	s.RuleId = &v
	return s
}

func (s *DescribeRequestHitResponseBodyResultObjectRuleHitRecords) SetRuleName(v string) *DescribeRequestHitResponseBodyResultObjectRuleHitRecords {
	s.RuleName = &v
	return s
}

func (s *DescribeRequestHitResponseBodyResultObjectRuleHitRecords) SetRuleSnapshotId(v string) *DescribeRequestHitResponseBodyResultObjectRuleHitRecords {
	s.RuleSnapshotId = &v
	return s
}

func (s *DescribeRequestHitResponseBodyResultObjectRuleHitRecords) SetRuleStatus(v string) *DescribeRequestHitResponseBodyResultObjectRuleHitRecords {
	s.RuleStatus = &v
	return s
}

func (s *DescribeRequestHitResponseBodyResultObjectRuleHitRecords) Validate() error {
	return dara.Validate(s)
}
