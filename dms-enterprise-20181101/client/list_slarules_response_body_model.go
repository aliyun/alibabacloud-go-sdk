// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSLARulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListSLARulesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListSLARulesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListSLARulesResponseBody
	GetRequestId() *string
	SetSLARuleList(v *ListSLARulesResponseBodySLARuleList) *ListSLARulesResponseBody
	GetSLARuleList() *ListSLARulesResponseBodySLARuleList
	SetSuccess(v bool) *ListSLARulesResponseBody
	GetSuccess() *bool
}

type ListSLARulesResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 3D1A59F4-EB2B-5D24-80A5-90C446A00DE2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of SLA rules.
	SLARuleList *ListSLARulesResponseBodySLARuleList `json:"SLARuleList,omitempty" xml:"SLARuleList,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSLARulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSLARulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSLARulesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSLARulesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListSLARulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSLARulesResponseBody) GetSLARuleList() *ListSLARulesResponseBodySLARuleList {
	return s.SLARuleList
}

func (s *ListSLARulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSLARulesResponseBody) SetErrorCode(v string) *ListSLARulesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSLARulesResponseBody) SetErrorMessage(v string) *ListSLARulesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSLARulesResponseBody) SetRequestId(v string) *ListSLARulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSLARulesResponseBody) SetSLARuleList(v *ListSLARulesResponseBodySLARuleList) *ListSLARulesResponseBody {
	s.SLARuleList = v
	return s
}

func (s *ListSLARulesResponseBody) SetSuccess(v bool) *ListSLARulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListSLARulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSLARulesResponseBodySLARuleList struct {
	SLARule []*ListSLARulesResponseBodySLARuleListSLARule `json:"SLARule,omitempty" xml:"SLARule,omitempty" type:"Repeated"`
}

func (s ListSLARulesResponseBodySLARuleList) String() string {
	return dara.Prettify(s)
}

func (s ListSLARulesResponseBodySLARuleList) GoString() string {
	return s.String()
}

func (s *ListSLARulesResponseBodySLARuleList) GetSLARule() []*ListSLARulesResponseBodySLARuleListSLARule {
	return s.SLARule
}

func (s *ListSLARulesResponseBodySLARuleList) SetSLARule(v []*ListSLARulesResponseBodySLARuleListSLARule) *ListSLARulesResponseBodySLARuleList {
	s.SLARule = v
	return s
}

func (s *ListSLARulesResponseBodySLARuleList) Validate() error {
	return dara.Validate(s)
}

type ListSLARulesResponseBodySLARuleListSLARule struct {
	// The ID of the task flow.
	//
	// example:
	//
	// 11****
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the SLA rule.
	//
	// example:
	//
	// 2
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The timeout period. Unit: minutes.
	//
	// example:
	//
	// 1080
	IntervalMinutes *int32 `json:"IntervalMinutes,omitempty" xml:"IntervalMinutes,omitempty"`
	// The ID of the task node.
	//
	// example:
	//
	// 1
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **0**: an SLA rule for a task flow
	//
	// 	- **1**: an SLA rule for a task node
	//
	// example:
	//
	// 0
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s ListSLARulesResponseBodySLARuleListSLARule) String() string {
	return dara.Prettify(s)
}

func (s ListSLARulesResponseBodySLARuleListSLARule) GoString() string {
	return s.String()
}

func (s *ListSLARulesResponseBodySLARuleListSLARule) GetDagId() *int64 {
	return s.DagId
}

func (s *ListSLARulesResponseBodySLARuleListSLARule) GetId() *int64 {
	return s.Id
}

func (s *ListSLARulesResponseBodySLARuleListSLARule) GetIntervalMinutes() *int32 {
	return s.IntervalMinutes
}

func (s *ListSLARulesResponseBodySLARuleListSLARule) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListSLARulesResponseBodySLARuleListSLARule) GetRuleType() *int32 {
	return s.RuleType
}

func (s *ListSLARulesResponseBodySLARuleListSLARule) SetDagId(v int64) *ListSLARulesResponseBodySLARuleListSLARule {
	s.DagId = &v
	return s
}

func (s *ListSLARulesResponseBodySLARuleListSLARule) SetId(v int64) *ListSLARulesResponseBodySLARuleListSLARule {
	s.Id = &v
	return s
}

func (s *ListSLARulesResponseBodySLARuleListSLARule) SetIntervalMinutes(v int32) *ListSLARulesResponseBodySLARuleListSLARule {
	s.IntervalMinutes = &v
	return s
}

func (s *ListSLARulesResponseBodySLARuleListSLARule) SetNodeId(v int64) *ListSLARulesResponseBodySLARuleListSLARule {
	s.NodeId = &v
	return s
}

func (s *ListSLARulesResponseBodySLARuleListSLARule) SetRuleType(v int32) *ListSLARulesResponseBodySLARuleListSLARule {
	s.RuleType = &v
	return s
}

func (s *ListSLARulesResponseBodySLARuleListSLARule) Validate() error {
	return dara.Validate(s)
}
