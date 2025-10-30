// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRecommendTaskDetailResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeRecommendTaskDetailResponseBodyResultObject) *DescribeRecommendTaskDetailResponseBody
	GetResultObject() *DescribeRecommendTaskDetailResponseBodyResultObject
}

type DescribeRecommendTaskDetailResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result
	ResultObject *DescribeRecommendTaskDetailResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeRecommendTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecommendTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecommendTaskDetailResponseBody) GetResultObject() *DescribeRecommendTaskDetailResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeRecommendTaskDetailResponseBody) SetRequestId(v string) *DescribeRecommendTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBody) SetResultObject(v *DescribeRecommendTaskDetailResponseBodyResultObject) *DescribeRecommendTaskDetailResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRecommendTaskDetailResponseBodyResultObject struct {
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
	// Comparison indicators
	ExpectVelocities []*string `json:"expectVelocities,omitempty" xml:"expectVelocities,omitempty" type:"Repeated"`
	// Creation time
	//
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Number of normal samples
	//
	// example:
	//
	// 100
	NormalSize *int64 `json:"normalSize,omitempty" xml:"normalSize,omitempty"`
	// Recommended strategy list
	RecommendRuleDTOs []*DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs `json:"recommendRuleDTOs,omitempty" xml:"recommendRuleDTOs,omitempty" type:"Repeated"`
	// Selected variable list
	RecommendVariableDTOs []*DescribeRecommendTaskDetailResponseBodyResultObjectRecommendVariableDTOs `json:"recommendVariableDTOs,omitempty" xml:"recommendVariableDTOs,omitempty" type:"Repeated"`
	// Number of risk samples
	//
	// example:
	//
	// 10
	RiskSize *int64 `json:"riskSize,omitempty" xml:"riskSize,omitempty"`
	// Sample name
	//
	// example:
	//
	// 白样本
	SampleName *string `json:"sampleName,omitempty" xml:"sampleName,omitempty"`
	// Sample scenario
	//
	// example:
	//
	// account_abuse_detection
	SampleScene *string `json:"sampleScene,omitempty" xml:"sampleScene,omitempty"`
	// Sample scenario name
	//
	// example:
	//
	// 防虚假账号
	SampleSceneName *string `json:"sampleSceneName,omitempty" xml:"sampleSceneName,omitempty"`
	// Task ID
	//
	// example:
	//
	// 887
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// Task name
	//
	// example:
	//
	// 策略推荐任务
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// Task status.
	//
	// example:
	//
	// CREATE
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s DescribeRecommendTaskDetailResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendTaskDetailResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) GetExpectVelocities() []*string {
	return s.ExpectVelocities
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) GetNormalSize() *int64 {
	return s.NormalSize
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) GetRecommendRuleDTOs() []*DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs {
	return s.RecommendRuleDTOs
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) GetRecommendVariableDTOs() []*DescribeRecommendTaskDetailResponseBodyResultObjectRecommendVariableDTOs {
	return s.RecommendVariableDTOs
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) GetRiskSize() *int64 {
	return s.RiskSize
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) GetSampleName() *string {
	return s.SampleName
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) GetSampleScene() *string {
	return s.SampleScene
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) GetSampleSceneName() *string {
	return s.SampleSceneName
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) SetEventCode(v string) *DescribeRecommendTaskDetailResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) SetEventName(v string) *DescribeRecommendTaskDetailResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) SetExpectVelocities(v []*string) *DescribeRecommendTaskDetailResponseBodyResultObject {
	s.ExpectVelocities = v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) SetGmtCreate(v int64) *DescribeRecommendTaskDetailResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) SetNormalSize(v int64) *DescribeRecommendTaskDetailResponseBodyResultObject {
	s.NormalSize = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) SetRecommendRuleDTOs(v []*DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) *DescribeRecommendTaskDetailResponseBodyResultObject {
	s.RecommendRuleDTOs = v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) SetRecommendVariableDTOs(v []*DescribeRecommendTaskDetailResponseBodyResultObjectRecommendVariableDTOs) *DescribeRecommendTaskDetailResponseBodyResultObject {
	s.RecommendVariableDTOs = v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) SetRiskSize(v int64) *DescribeRecommendTaskDetailResponseBodyResultObject {
	s.RiskSize = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) SetSampleName(v string) *DescribeRecommendTaskDetailResponseBodyResultObject {
	s.SampleName = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) SetSampleScene(v string) *DescribeRecommendTaskDetailResponseBodyResultObject {
	s.SampleScene = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) SetSampleSceneName(v string) *DescribeRecommendTaskDetailResponseBodyResultObject {
	s.SampleSceneName = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) SetTaskId(v int64) *DescribeRecommendTaskDetailResponseBodyResultObject {
	s.TaskId = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) SetTaskName(v string) *DescribeRecommendTaskDetailResponseBodyResultObject {
	s.TaskName = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) SetTaskStatus(v string) *DescribeRecommendTaskDetailResponseBodyResultObject {
	s.TaskStatus = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObject) Validate() error {
	if s.RecommendRuleDTOs != nil {
		for _, item := range s.RecommendRuleDTOs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RecommendVariableDTOs != nil {
		for _, item := range s.RecommendVariableDTOs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs struct {
	// Calculation path
	//
	// example:
	//
	// 1&2
	ComputeExpression *string `json:"computeExpression,omitempty" xml:"computeExpression,omitempty"`
	// Number of hit samples
	//
	// example:
	//
	// 99
	HitSample *int64 `json:"hitSample,omitempty" xml:"hitSample,omitempty"`
	// Primary key ID of the rule
	//
	// example:
	//
	// 30
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Number of unhit samples
	//
	// example:
	//
	// 1
	NotHitSample *int64 `json:"notHitSample,omitempty" xml:"notHitSample,omitempty"`
	// List of candidate rules
	RecommendRules []*DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOsRecommendRules `json:"recommendRules,omitempty" xml:"recommendRules,omitempty" type:"Repeated"`
	// Strategy ID
	//
	// example:
	//
	// 102224
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Rule name
	//
	// example:
	//
	// 营销风险识别评分_高风险_拒绝_副本
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Status
	//
	// example:
	//
	// NO_RULE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// List of indicators in JSON string format
	//
	// example:
	//
	// [{"key":"dxkkLw8fe18","value":"2"}]
	Velocities *string `json:"velocities,omitempty" xml:"velocities,omitempty"`
}

func (s DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) GoString() string {
	return s.String()
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) GetComputeExpression() *string {
	return s.ComputeExpression
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) GetHitSample() *int64 {
	return s.HitSample
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) GetId() *int64 {
	return s.Id
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) GetNotHitSample() *int64 {
	return s.NotHitSample
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) GetRecommendRules() []*DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOsRecommendRules {
	return s.RecommendRules
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) GetStatus() *string {
	return s.Status
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) GetVelocities() *string {
	return s.Velocities
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) SetComputeExpression(v string) *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs {
	s.ComputeExpression = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) SetHitSample(v int64) *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs {
	s.HitSample = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) SetId(v int64) *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs {
	s.Id = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) SetNotHitSample(v int64) *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs {
	s.NotHitSample = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) SetRecommendRules(v []*DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOsRecommendRules) *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs {
	s.RecommendRules = v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) SetRuleId(v string) *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs {
	s.RuleId = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) SetRuleName(v string) *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs {
	s.RuleName = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) SetStatus(v string) *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs {
	s.Status = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) SetVelocities(v string) *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs {
	s.Velocities = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOs) Validate() error {
	if s.RecommendRules != nil {
		for _, item := range s.RecommendRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOsRecommendRules struct {
	// Left variable
	//
	// example:
	//
	// age
	Left *string `json:"left,omitempty" xml:"left,omitempty"`
	// Operator
	//
	// example:
	//
	// equals
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// Right variable
	//
	// example:
	//
	// 20
	Right *string `json:"right,omitempty" xml:"right,omitempty"`
}

func (s DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOsRecommendRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOsRecommendRules) GoString() string {
	return s.String()
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOsRecommendRules) GetLeft() *string {
	return s.Left
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOsRecommendRules) GetOperator() *string {
	return s.Operator
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOsRecommendRules) GetRight() *string {
	return s.Right
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOsRecommendRules) SetLeft(v string) *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOsRecommendRules {
	s.Left = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOsRecommendRules) SetOperator(v string) *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOsRecommendRules {
	s.Operator = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOsRecommendRules) SetRight(v string) *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOsRecommendRules {
	s.Right = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendRuleDTOsRecommendRules) Validate() error {
	return dara.Validate(s)
}

type DescribeRecommendTaskDetailResponseBodyResultObjectRecommendVariableDTOs struct {
	// Primary key ID
	//
	// example:
	//
	// 234
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Title
	//
	// example:
	//
	// 手机号
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s DescribeRecommendTaskDetailResponseBodyResultObjectRecommendVariableDTOs) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendTaskDetailResponseBodyResultObjectRecommendVariableDTOs) GoString() string {
	return s.String()
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendVariableDTOs) GetId() *int64 {
	return s.Id
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendVariableDTOs) GetTitle() *string {
	return s.Title
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendVariableDTOs) SetId(v int64) *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendVariableDTOs {
	s.Id = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendVariableDTOs) SetTitle(v string) *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendVariableDTOs {
	s.Title = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponseBodyResultObjectRecommendVariableDTOs) Validate() error {
	return dara.Validate(s)
}
