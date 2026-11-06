// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResponseRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateResponseRuleRequest
	GetLang() *string
	SetMaxResults(v int32) *UpdateResponseRuleRequest
	GetMaxResults() *int32
	SetNextToken(v string) *UpdateResponseRuleRequest
	GetNextToken() *string
	SetRegionId(v string) *UpdateResponseRuleRequest
	GetRegionId() *string
	SetResponseActionConfig(v string) *UpdateResponseRuleRequest
	GetResponseActionConfig() *string
	SetResponseActionType(v string) *UpdateResponseRuleRequest
	GetResponseActionType() *string
	SetResponseExecutionCondition(v string) *UpdateResponseRuleRequest
	GetResponseExecutionCondition() *string
	SetResponseRuleId(v string) *UpdateResponseRuleRequest
	GetResponseRuleId() *string
	SetResponseRuleName(v string) *UpdateResponseRuleRequest
	GetResponseRuleName() *string
	SetResponseRulePriority(v int32) *UpdateResponseRuleRequest
	GetResponseRulePriority() *int32
	SetResponseRuleRemark(v string) *UpdateResponseRuleRequest
	GetResponseRuleRemark() *string
	SetResponseRuleStatus(v int32) *UpdateResponseRuleRequest
	GetResponseRuleStatus() *int32
	SetResponseTriggerType(v string) *UpdateResponseRuleRequest
	GetResponseTriggerType() *string
}

type UpdateResponseRuleRequest struct {
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of data records to read in the current request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that specifies the position from which to start reading. If this parameter is left empty, the reading starts from the beginning.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region where the data management center of Threat Analysis is located. You must select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: The assets are located in the Chinese mainland.
	//
	// - ap-southeast-1: The assets are located outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The action configuration of the automated response rule.
	//
	// example:
	//
	// [{"actionType":"doPlaybook","playbookName":"block waf IP","playbookUuid":"system_aliyun_waf_whole_process_book","disposeParam":{"period":"7d"}}]
	ResponseActionConfig *string `json:"ResponseActionConfig,omitempty" xml:"ResponseActionConfig,omitempty"`
	// The action type of the automated response rule. Valid values:
	//
	// - doPlaybook: Executes a playbook.
	//
	// - changeEventStatus: Updates the event status.
	//
	// - changeThreatLevel: Updates the threat level of the event.
	//
	// - addEventTag: Adds an event tag.
	//
	// - deleteEventTag: Deletes an event tag.
	//
	// - alertWhitelist: Adds an alert to the whitelist.
	//
	// example:
	//
	// alertWhitelist
	ResponseActionType *string `json:"ResponseActionType,omitempty" xml:"ResponseActionType,omitempty"`
	// The rule trigger condition configuration.
	//
	// example:
	//
	// [{"left":{"value":"threat_level"},"operator":"equals","right":{"value":"suspicious"}}]
	ResponseExecutionCondition *string `json:"ResponseExecutionCondition,omitempty" xml:"ResponseExecutionCondition,omitempty"`
	// The ID of the automated response rule.
	//
	// example:
	//
	// 440918
	ResponseRuleId *string `json:"ResponseRuleId,omitempty" xml:"ResponseRuleId,omitempty"`
	// The name of the automated response rule.
	//
	// example:
	//
	// Send Notification When Generating Urgent Incident
	ResponseRuleName *string `json:"ResponseRuleName,omitempty" xml:"ResponseRuleName,omitempty"`
	// The execution priority of the automated response rule.
	//
	// example:
	//
	// 1
	ResponseRulePriority *int32 `json:"ResponseRulePriority,omitempty" xml:"ResponseRulePriority,omitempty"`
	// The remarks of the response rule.
	//
	// example:
	//
	// Verified as a low-risk alert and scheduled for regular cleanup
	ResponseRuleRemark *string `json:"ResponseRuleRemark,omitempty" xml:"ResponseRuleRemark,omitempty"`
	// The status of the automated response rule. Valid values:
	//
	// - 0: Disabled.
	//
	// - 100: Enabled.
	//
	// example:
	//
	// 0
	ResponseRuleStatus *int32 `json:"ResponseRuleStatus,omitempty" xml:"ResponseRuleStatus,omitempty"`
	// The trigger type of the automated response rule. Valid values:
	//
	// - event: An event occurs.
	//
	// - event_update: An event is updated.
	//
	// - alert: An alert is triggered.
	//
	// example:
	//
	// event
	ResponseTriggerType *string `json:"ResponseTriggerType,omitempty" xml:"ResponseTriggerType,omitempty"`
}

func (s UpdateResponseRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResponseRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateResponseRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateResponseRuleRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *UpdateResponseRuleRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *UpdateResponseRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateResponseRuleRequest) GetResponseActionConfig() *string {
	return s.ResponseActionConfig
}

func (s *UpdateResponseRuleRequest) GetResponseActionType() *string {
	return s.ResponseActionType
}

func (s *UpdateResponseRuleRequest) GetResponseExecutionCondition() *string {
	return s.ResponseExecutionCondition
}

func (s *UpdateResponseRuleRequest) GetResponseRuleId() *string {
	return s.ResponseRuleId
}

func (s *UpdateResponseRuleRequest) GetResponseRuleName() *string {
	return s.ResponseRuleName
}

func (s *UpdateResponseRuleRequest) GetResponseRulePriority() *int32 {
	return s.ResponseRulePriority
}

func (s *UpdateResponseRuleRequest) GetResponseRuleRemark() *string {
	return s.ResponseRuleRemark
}

func (s *UpdateResponseRuleRequest) GetResponseRuleStatus() *int32 {
	return s.ResponseRuleStatus
}

func (s *UpdateResponseRuleRequest) GetResponseTriggerType() *string {
	return s.ResponseTriggerType
}

func (s *UpdateResponseRuleRequest) SetLang(v string) *UpdateResponseRuleRequest {
	s.Lang = &v
	return s
}

func (s *UpdateResponseRuleRequest) SetMaxResults(v int32) *UpdateResponseRuleRequest {
	s.MaxResults = &v
	return s
}

func (s *UpdateResponseRuleRequest) SetNextToken(v string) *UpdateResponseRuleRequest {
	s.NextToken = &v
	return s
}

func (s *UpdateResponseRuleRequest) SetRegionId(v string) *UpdateResponseRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateResponseRuleRequest) SetResponseActionConfig(v string) *UpdateResponseRuleRequest {
	s.ResponseActionConfig = &v
	return s
}

func (s *UpdateResponseRuleRequest) SetResponseActionType(v string) *UpdateResponseRuleRequest {
	s.ResponseActionType = &v
	return s
}

func (s *UpdateResponseRuleRequest) SetResponseExecutionCondition(v string) *UpdateResponseRuleRequest {
	s.ResponseExecutionCondition = &v
	return s
}

func (s *UpdateResponseRuleRequest) SetResponseRuleId(v string) *UpdateResponseRuleRequest {
	s.ResponseRuleId = &v
	return s
}

func (s *UpdateResponseRuleRequest) SetResponseRuleName(v string) *UpdateResponseRuleRequest {
	s.ResponseRuleName = &v
	return s
}

func (s *UpdateResponseRuleRequest) SetResponseRulePriority(v int32) *UpdateResponseRuleRequest {
	s.ResponseRulePriority = &v
	return s
}

func (s *UpdateResponseRuleRequest) SetResponseRuleRemark(v string) *UpdateResponseRuleRequest {
	s.ResponseRuleRemark = &v
	return s
}

func (s *UpdateResponseRuleRequest) SetResponseRuleStatus(v int32) *UpdateResponseRuleRequest {
	s.ResponseRuleStatus = &v
	return s
}

func (s *UpdateResponseRuleRequest) SetResponseTriggerType(v string) *UpdateResponseRuleRequest {
	s.ResponseTriggerType = &v
	return s
}

func (s *UpdateResponseRuleRequest) Validate() error {
	return dara.Validate(s)
}
