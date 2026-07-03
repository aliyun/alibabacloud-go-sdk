// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResponseRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateResponseRuleRequest
	GetLang() *string
	SetMaxResults(v int32) *CreateResponseRuleRequest
	GetMaxResults() *int32
	SetNextToken(v string) *CreateResponseRuleRequest
	GetNextToken() *string
	SetRegionId(v string) *CreateResponseRuleRequest
	GetRegionId() *string
	SetResponseActionConfig(v string) *CreateResponseRuleRequest
	GetResponseActionConfig() *string
	SetResponseActionType(v string) *CreateResponseRuleRequest
	GetResponseActionType() *string
	SetResponseExecutionCondition(v string) *CreateResponseRuleRequest
	GetResponseExecutionCondition() *string
	SetResponseRuleName(v string) *CreateResponseRuleRequest
	GetResponseRuleName() *string
	SetResponseRulePriority(v string) *CreateResponseRuleRequest
	GetResponseRulePriority() *string
	SetResponseTriggerType(v string) *CreateResponseRuleRequest
	GetResponseTriggerType() *string
	SetRoleFor(v int64) *CreateResponseRuleRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *CreateResponseRuleRequest
	GetRoleType() *int32
}

type CreateResponseRuleRequest struct {
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of results to return.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that specifies the position from which to start the query. If you do not specify this parameter, the query starts from the beginning.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The deployment region of the data management center for threat analysis. You must select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: Your assets are in regions outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The action configuration, specified as a JSON string.
	//
	// example:
	//
	// [{"actionType":"doPlaybook","playbookName":"block waf IP","playbookUuid":"system_aliyun_waf_whole_process_book","disposeParam":{"period":"7d"}}]
	ResponseActionConfig *string `json:"ResponseActionConfig,omitempty" xml:"ResponseActionConfig,omitempty"`
	// The action type for the automatic response rule. Valid values:
	//
	// - doPlaybook: Runs a playbook.
	//
	// - changeEventStatus: Changes the status of an event.
	//
	// - changeThreatLevel: Changes the threat level of an event.
	//
	// - addEventTag: Adds a tag to an event.
	//
	// - deleteEventTag: Deletes a tag from an event.
	//
	// - alertWhitelist: Adds an alert to the allowlist.
	//
	// example:
	//
	// doPlaybook
	ResponseActionType *string `json:"ResponseActionType,omitempty" xml:"ResponseActionType,omitempty"`
	// The trigger conditions for the rule, specified as a JSON string.
	//
	// example:
	//
	// [{"left":{"value":"threat_level"},"operator":"equals","right":{"value":"suspicious"}}]
	ResponseExecutionCondition *string `json:"ResponseExecutionCondition,omitempty" xml:"ResponseExecutionCondition,omitempty"`
	// The name of the automatic response rule.
	//
	// example:
	//
	// Send Notification When Generating Urgent Incident
	ResponseRuleName *string `json:"ResponseRuleName,omitempty" xml:"ResponseRuleName,omitempty"`
	// The execution priority of the automatic response rule.
	//
	// example:
	//
	// 1
	ResponseRulePriority *string `json:"ResponseRulePriority,omitempty" xml:"ResponseRulePriority,omitempty"`
	// The trigger type for the automatic response rule. Valid values:
	//
	// - event: An event is generated.
	//
	// - event_update: An event is updated.
	//
	// - alert: An alert is generated.
	//
	// example:
	//
	// event
	ResponseTriggerType *string `json:"ResponseTriggerType,omitempty" xml:"ResponseTriggerType,omitempty"`
	// The ID of the member account. An administrator uses this parameter to operate on behalf of the specified member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The operational scope. Valid values:
	//
	// - 0: Sets the scope to the current Alibaba Cloud account.
	//
	// - 1: Sets the scope to all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s CreateResponseRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResponseRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateResponseRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateResponseRuleRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *CreateResponseRuleRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *CreateResponseRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateResponseRuleRequest) GetResponseActionConfig() *string {
	return s.ResponseActionConfig
}

func (s *CreateResponseRuleRequest) GetResponseActionType() *string {
	return s.ResponseActionType
}

func (s *CreateResponseRuleRequest) GetResponseExecutionCondition() *string {
	return s.ResponseExecutionCondition
}

func (s *CreateResponseRuleRequest) GetResponseRuleName() *string {
	return s.ResponseRuleName
}

func (s *CreateResponseRuleRequest) GetResponseRulePriority() *string {
	return s.ResponseRulePriority
}

func (s *CreateResponseRuleRequest) GetResponseTriggerType() *string {
	return s.ResponseTriggerType
}

func (s *CreateResponseRuleRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *CreateResponseRuleRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *CreateResponseRuleRequest) SetLang(v string) *CreateResponseRuleRequest {
	s.Lang = &v
	return s
}

func (s *CreateResponseRuleRequest) SetMaxResults(v int32) *CreateResponseRuleRequest {
	s.MaxResults = &v
	return s
}

func (s *CreateResponseRuleRequest) SetNextToken(v string) *CreateResponseRuleRequest {
	s.NextToken = &v
	return s
}

func (s *CreateResponseRuleRequest) SetRegionId(v string) *CreateResponseRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateResponseRuleRequest) SetResponseActionConfig(v string) *CreateResponseRuleRequest {
	s.ResponseActionConfig = &v
	return s
}

func (s *CreateResponseRuleRequest) SetResponseActionType(v string) *CreateResponseRuleRequest {
	s.ResponseActionType = &v
	return s
}

func (s *CreateResponseRuleRequest) SetResponseExecutionCondition(v string) *CreateResponseRuleRequest {
	s.ResponseExecutionCondition = &v
	return s
}

func (s *CreateResponseRuleRequest) SetResponseRuleName(v string) *CreateResponseRuleRequest {
	s.ResponseRuleName = &v
	return s
}

func (s *CreateResponseRuleRequest) SetResponseRulePriority(v string) *CreateResponseRuleRequest {
	s.ResponseRulePriority = &v
	return s
}

func (s *CreateResponseRuleRequest) SetResponseTriggerType(v string) *CreateResponseRuleRequest {
	s.ResponseTriggerType = &v
	return s
}

func (s *CreateResponseRuleRequest) SetRoleFor(v int64) *CreateResponseRuleRequest {
	s.RoleFor = &v
	return s
}

func (s *CreateResponseRuleRequest) SetRoleType(v int32) *CreateResponseRuleRequest {
	s.RoleType = &v
	return s
}

func (s *CreateResponseRuleRequest) Validate() error {
	return dara.Validate(s)
}
