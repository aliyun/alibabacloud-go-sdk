// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScanRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetScanRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *GetScanRuleResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetScanRuleResponseBody
	GetRequestId() *string
	SetScanRule(v *GetScanRuleResponseBodyScanRule) *GetScanRuleResponseBody
	GetScanRule() *GetScanRuleResponseBodyScanRule
}

type GetScanRuleResponseBody struct {
	// Return value
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API call was successful, valid values:
	//
	// - `true`: The API call was successful
	//
	// - `false`: The API call failed
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C87993B5-7D61-5CAC-8D64-1AC732DD69FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scan rule.
	ScanRule *GetScanRuleResponseBodyScanRule `json:"ScanRule,omitempty" xml:"ScanRule,omitempty" type:"Struct"`
}

func (s GetScanRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScanRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetScanRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetScanRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetScanRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScanRuleResponseBody) GetScanRule() *GetScanRuleResponseBodyScanRule {
	return s.ScanRule
}

func (s *GetScanRuleResponseBody) SetCode(v string) *GetScanRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetScanRuleResponseBody) SetIsSuccess(v bool) *GetScanRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetScanRuleResponseBody) SetRequestId(v string) *GetScanRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScanRuleResponseBody) SetScanRule(v *GetScanRuleResponseBodyScanRule) *GetScanRuleResponseBody {
	s.ScanRule = v
	return s
}

func (s *GetScanRuleResponseBody) Validate() error {
	if s.ScanRule != nil {
		if err := s.ScanRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScanRuleResponseBodyScanRule struct {
	// The creation time.
	//
	// example:
	//
	// 2018-03-15T17:08Z
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-szw6f6bhrky0c8jk
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Names of namespaces where the event is effective.
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// Names of repositories where the event is effective.
	RepoNames []*string `json:"RepoNames,omitempty" xml:"RepoNames,omitempty" type:"Repeated"`
	// Tag filter pattern for event triggering.
	//
	// example:
	//
	// .*
	RepoTagFilterPattern *string `json:"RepoTagFilterPattern,omitempty" xml:"RepoTagFilterPattern,omitempty"`
	// The event rule name.
	//
	// example:
	//
	// protection
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The scan rule ID.
	//
	// example:
	//
	// crscnr-aemytkwad2h7fyhb
	ScanRuleId *string `json:"ScanRuleId,omitempty" xml:"ScanRuleId,omitempty"`
	// The scan scope.
	//
	// example:
	//
	// REPO
	ScanScope *string `json:"ScanScope,omitempty" xml:"ScanScope,omitempty"`
	// The vulnerability type:
	//
	// - `cve`: System vulnerability
	//
	// - `sca`: Application vulnerability
	//
	// example:
	//
	// SBOM
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The trigger type, valid values:
	//
	// - `ALL`: All triggers
	//
	// - `TAG_LISTTAG`: Trigger
	//
	// - `TAG_REG_EXP`: Expression trigger
	//
	// example:
	//
	// AUTO
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2025-08-28T20:07:33.164
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetScanRuleResponseBodyScanRule) String() string {
	return dara.Prettify(s)
}

func (s GetScanRuleResponseBodyScanRule) GoString() string {
	return s.String()
}

func (s *GetScanRuleResponseBodyScanRule) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetScanRuleResponseBodyScanRule) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetScanRuleResponseBodyScanRule) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *GetScanRuleResponseBodyScanRule) GetRepoNames() []*string {
	return s.RepoNames
}

func (s *GetScanRuleResponseBodyScanRule) GetRepoTagFilterPattern() *string {
	return s.RepoTagFilterPattern
}

func (s *GetScanRuleResponseBodyScanRule) GetRuleName() *string {
	return s.RuleName
}

func (s *GetScanRuleResponseBodyScanRule) GetScanRuleId() *string {
	return s.ScanRuleId
}

func (s *GetScanRuleResponseBodyScanRule) GetScanScope() *string {
	return s.ScanScope
}

func (s *GetScanRuleResponseBodyScanRule) GetScanType() *string {
	return s.ScanType
}

func (s *GetScanRuleResponseBodyScanRule) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetScanRuleResponseBodyScanRule) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetScanRuleResponseBodyScanRule) SetCreateTime(v int64) *GetScanRuleResponseBodyScanRule {
	s.CreateTime = &v
	return s
}

func (s *GetScanRuleResponseBodyScanRule) SetInstanceId(v string) *GetScanRuleResponseBodyScanRule {
	s.InstanceId = &v
	return s
}

func (s *GetScanRuleResponseBodyScanRule) SetNamespaces(v []*string) *GetScanRuleResponseBodyScanRule {
	s.Namespaces = v
	return s
}

func (s *GetScanRuleResponseBodyScanRule) SetRepoNames(v []*string) *GetScanRuleResponseBodyScanRule {
	s.RepoNames = v
	return s
}

func (s *GetScanRuleResponseBodyScanRule) SetRepoTagFilterPattern(v string) *GetScanRuleResponseBodyScanRule {
	s.RepoTagFilterPattern = &v
	return s
}

func (s *GetScanRuleResponseBodyScanRule) SetRuleName(v string) *GetScanRuleResponseBodyScanRule {
	s.RuleName = &v
	return s
}

func (s *GetScanRuleResponseBodyScanRule) SetScanRuleId(v string) *GetScanRuleResponseBodyScanRule {
	s.ScanRuleId = &v
	return s
}

func (s *GetScanRuleResponseBodyScanRule) SetScanScope(v string) *GetScanRuleResponseBodyScanRule {
	s.ScanScope = &v
	return s
}

func (s *GetScanRuleResponseBodyScanRule) SetScanType(v string) *GetScanRuleResponseBodyScanRule {
	s.ScanType = &v
	return s
}

func (s *GetScanRuleResponseBodyScanRule) SetTriggerType(v string) *GetScanRuleResponseBodyScanRule {
	s.TriggerType = &v
	return s
}

func (s *GetScanRuleResponseBodyScanRule) SetUpdateTime(v int64) *GetScanRuleResponseBodyScanRule {
	s.UpdateTime = &v
	return s
}

func (s *GetScanRuleResponseBodyScanRule) Validate() error {
	return dara.Validate(s)
}
