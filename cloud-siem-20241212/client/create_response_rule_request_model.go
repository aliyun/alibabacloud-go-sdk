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
	SetResponseRuleRemark(v string) *CreateResponseRuleRequest
	GetResponseRuleRemark() *string
	SetResponseTriggerType(v string) *CreateResponseRuleRequest
	GetResponseTriggerType() *string
	SetRoleFor(v int64) *CreateResponseRuleRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *CreateResponseRuleRequest
	GetRoleType() *int32
}

type CreateResponseRuleRequest struct {
	// The language of the response messages. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of data records to read in this request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the start position of the current read. Leave this parameter empty to start from the beginning.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region where the data management center for Cloud Threat Detection and Response (CTDR) is located. Select the management center region based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: Assets are located in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The action configuration of the automatic response rule.
	//
	// example:
	//
	// [{"actionType":"doPlaybook","playbookName":"block waf IP","playbookUuid":"system_aliyun_waf_whole_process_book","disposeParam":{"period":"7d"}}]
	ResponseActionConfig *string `json:"ResponseActionConfig,omitempty" xml:"ResponseActionConfig,omitempty"`
	// The action type of the automatic response rule. Valid values:
	//
	// - doPlaybook: runs a playbook.
	//
	// - changeEventStatus: updates the event status.
	//
	// - changeThreatLevel: updates the threat level of an event.
	//
	// - addEventTag: adds a tag to an event.
	//
	// - deleteEventTag: deletes a tag from an event.
	//
	// - alertWhitelist: adds an alert to the whitelist.
	//
	// example:
	//
	// doPlaybook
	ResponseActionType *string `json:"ResponseActionType,omitempty" xml:"ResponseActionType,omitempty"`
	// The trigger condition configuration of the rule.
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
	// The remarks for the response rule.
	ResponseRuleRemark *string `json:"ResponseRuleRemark,omitempty" xml:"ResponseRuleRemark,omitempty"`
	// The trigger type of the automatic response rule. Valid values:
	//
	// - event: an event occurs.
	//
	// - event_update: an event is updated.
	//
	// - alert: an alert occurs.
	//
	// example:
	//
	// event
	ResponseTriggerType *string `json:"ResponseTriggerType,omitempty" xml:"ResponseTriggerType,omitempty"`
	// The user ID that the administrator uses to switch to the view of another member account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts in the enterprise.
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

func (s *CreateResponseRuleRequest) GetResponseRuleRemark() *string {
	return s.ResponseRuleRemark
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

func (s *CreateResponseRuleRequest) SetResponseRuleRemark(v string) *CreateResponseRuleRequest {
	s.ResponseRuleRemark = &v
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
