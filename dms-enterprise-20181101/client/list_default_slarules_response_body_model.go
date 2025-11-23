// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDefaultSLARulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListDefaultSLARulesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDefaultSLARulesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListDefaultSLARulesResponseBody
	GetRequestId() *string
	SetSLARuleList(v *ListDefaultSLARulesResponseBodySLARuleList) *ListDefaultSLARulesResponseBody
	GetSLARuleList() *ListDefaultSLARulesResponseBodySLARuleList
	SetSuccess(v bool) *ListDefaultSLARulesResponseBody
	GetSuccess() *bool
}

type ListDefaultSLARulesResponseBody struct {
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
	// 3E8AF4C3-A822-53A8-970C-059EE83BBD5A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of SLA rules.
	SLARuleList *ListDefaultSLARulesResponseBodySLARuleList `json:"SLARuleList,omitempty" xml:"SLARuleList,omitempty" type:"Struct"`
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

func (s ListDefaultSLARulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDefaultSLARulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDefaultSLARulesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDefaultSLARulesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDefaultSLARulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDefaultSLARulesResponseBody) GetSLARuleList() *ListDefaultSLARulesResponseBodySLARuleList {
	return s.SLARuleList
}

func (s *ListDefaultSLARulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDefaultSLARulesResponseBody) SetErrorCode(v string) *ListDefaultSLARulesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDefaultSLARulesResponseBody) SetErrorMessage(v string) *ListDefaultSLARulesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDefaultSLARulesResponseBody) SetRequestId(v string) *ListDefaultSLARulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDefaultSLARulesResponseBody) SetSLARuleList(v *ListDefaultSLARulesResponseBodySLARuleList) *ListDefaultSLARulesResponseBody {
	s.SLARuleList = v
	return s
}

func (s *ListDefaultSLARulesResponseBody) SetSuccess(v bool) *ListDefaultSLARulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListDefaultSLARulesResponseBody) Validate() error {
	if s.SLARuleList != nil {
		if err := s.SLARuleList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDefaultSLARulesResponseBodySLARuleList struct {
	SLARule []*ListDefaultSLARulesResponseBodySLARuleListSLARule `json:"SLARule,omitempty" xml:"SLARule,omitempty" type:"Repeated"`
}

func (s ListDefaultSLARulesResponseBodySLARuleList) String() string {
	return dara.Prettify(s)
}

func (s ListDefaultSLARulesResponseBodySLARuleList) GoString() string {
	return s.String()
}

func (s *ListDefaultSLARulesResponseBodySLARuleList) GetSLARule() []*ListDefaultSLARulesResponseBodySLARuleListSLARule {
	return s.SLARule
}

func (s *ListDefaultSLARulesResponseBodySLARuleList) SetSLARule(v []*ListDefaultSLARulesResponseBodySLARuleListSLARule) *ListDefaultSLARulesResponseBodySLARuleList {
	s.SLARule = v
	return s
}

func (s *ListDefaultSLARulesResponseBodySLARuleList) Validate() error {
	if s.SLARule != nil {
		for _, item := range s.SLARule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDefaultSLARulesResponseBodySLARuleListSLARule struct {
	// The ID of the task flow.
	//
	// example:
	//
	// 0
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the SLA rule.
	//
	// example:
	//
	// 1
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
	// 0
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

func (s ListDefaultSLARulesResponseBodySLARuleListSLARule) String() string {
	return dara.Prettify(s)
}

func (s ListDefaultSLARulesResponseBodySLARuleListSLARule) GoString() string {
	return s.String()
}

func (s *ListDefaultSLARulesResponseBodySLARuleListSLARule) GetDagId() *int64 {
	return s.DagId
}

func (s *ListDefaultSLARulesResponseBodySLARuleListSLARule) GetId() *int64 {
	return s.Id
}

func (s *ListDefaultSLARulesResponseBodySLARuleListSLARule) GetIntervalMinutes() *int32 {
	return s.IntervalMinutes
}

func (s *ListDefaultSLARulesResponseBodySLARuleListSLARule) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListDefaultSLARulesResponseBodySLARuleListSLARule) GetRuleType() *int32 {
	return s.RuleType
}

func (s *ListDefaultSLARulesResponseBodySLARuleListSLARule) SetDagId(v int64) *ListDefaultSLARulesResponseBodySLARuleListSLARule {
	s.DagId = &v
	return s
}

func (s *ListDefaultSLARulesResponseBodySLARuleListSLARule) SetId(v int64) *ListDefaultSLARulesResponseBodySLARuleListSLARule {
	s.Id = &v
	return s
}

func (s *ListDefaultSLARulesResponseBodySLARuleListSLARule) SetIntervalMinutes(v int32) *ListDefaultSLARulesResponseBodySLARuleListSLARule {
	s.IntervalMinutes = &v
	return s
}

func (s *ListDefaultSLARulesResponseBodySLARuleListSLARule) SetNodeId(v int64) *ListDefaultSLARulesResponseBodySLARuleListSLARule {
	s.NodeId = &v
	return s
}

func (s *ListDefaultSLARulesResponseBodySLARuleListSLARule) SetRuleType(v int32) *ListDefaultSLARulesResponseBodySLARuleListSLARule {
	s.RuleType = &v
	return s
}

func (s *ListDefaultSLARulesResponseBodySLARuleListSLARule) Validate() error {
	return dara.Validate(s)
}
