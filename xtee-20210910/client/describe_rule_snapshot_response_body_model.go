// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRuleSnapshotResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeRuleSnapshotResponseBodyResultObject) *DescribeRuleSnapshotResponseBody
	GetResultObject() *DescribeRuleSnapshotResponseBodyResultObject
}

type DescribeRuleSnapshotResponseBody struct {
	// Request ID
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	ResultObject *DescribeRuleSnapshotResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeRuleSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRuleSnapshotResponseBody) GetResultObject() *DescribeRuleSnapshotResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeRuleSnapshotResponseBody) SetRequestId(v string) *DescribeRuleSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBody) SetResultObject(v *DescribeRuleSnapshotResponseBodyResultObject) *DescribeRuleSnapshotResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeRuleSnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRuleSnapshotResponseBodyResultObject struct {
	// Business version.
	//
	// example:
	//
	// 1
	BizVersion *string `json:"bizVersion,omitempty" xml:"bizVersion,omitempty"`
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
	// Creation time.
	//
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 1565701886000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Expression for analysis results.
	//
	// example:
	//
	// 1&2
	LogicExpression *string `json:"logicExpression,omitempty" xml:"logicExpression,omitempty"`
	// Memo.
	//
	// example:
	//
	// 备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// Rule actions.
	//
	// example:
	//
	// [{\\"inputs\\":[\\"LdShop\\"],\\"name\\":\\"__addDeTags__\\",\\"actionType\\":\\"TAG\\",\\"outputType\\":\\"const\\"}]
	RuleActions *string `json:"ruleActions,omitempty" xml:"ruleActions,omitempty"`
	// Expression.
	//
	// example:
	//
	// [{\\"expressionName\\":\\"代下单_记录日志\\",\\"itemId\\":1,\\"left\\":{\\"name\\":\\"dhcfX2v7670\\"},\\"operatorCode\\":\\"gte\\",\\"operatorName\\":\\"大于等于\\",\\"right\\":{\\"fieldValue\\":\\"2\\"}}]
	RuleExpressions *string `json:"ruleExpressions,omitempty" xml:"ruleExpressions,omitempty"`
	// Policy ID
	//
	// example:
	//
	// 101804
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Policy name
	//
	// example:
	//
	// 营销风险识别
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Policy status
	//
	// example:
	//
	// RUNNING
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
}

func (s DescribeRuleSnapshotResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleSnapshotResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetBizVersion() *string {
	return s.BizVersion
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetLogicExpression() *string {
	return s.LogicExpression
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetMemo() *string {
	return s.Memo
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetRuleActions() *string {
	return s.RuleActions
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetRuleExpressions() *string {
	return s.RuleExpressions
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetBizVersion(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.BizVersion = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetEventCode(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetEventName(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetGmtCreate(v int64) *DescribeRuleSnapshotResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetGmtModified(v int64) *DescribeRuleSnapshotResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetLogicExpression(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.LogicExpression = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetMemo(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.Memo = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetRuleActions(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.RuleActions = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetRuleExpressions(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.RuleExpressions = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetRuleId(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.RuleId = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetRuleName(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.RuleName = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetRuleStatus(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.RuleStatus = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
