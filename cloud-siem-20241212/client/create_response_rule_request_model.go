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
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// [{"actionType":"doPlaybook","playbookName":"block waf IP","playbookUuid":"system_aliyun_waf_whole_process_book","disposeParam":{"period":"7d"}}]
	ResponseActionConfig *string `json:"ResponseActionConfig,omitempty" xml:"ResponseActionConfig,omitempty"`
	// example:
	//
	// doPlaybook
	ResponseActionType *string `json:"ResponseActionType,omitempty" xml:"ResponseActionType,omitempty"`
	// example:
	//
	// [{"left":{"value":"threat_level"},"operator":"equals","right":{"value":"suspicious"}}]
	ResponseExecutionCondition *string `json:"ResponseExecutionCondition,omitempty" xml:"ResponseExecutionCondition,omitempty"`
	// example:
	//
	// Send Notification When Generating Urgent Incident
	ResponseRuleName *string `json:"ResponseRuleName,omitempty" xml:"ResponseRuleName,omitempty"`
	// example:
	//
	// 1
	ResponseRulePriority *string `json:"ResponseRulePriority,omitempty" xml:"ResponseRulePriority,omitempty"`
	// example:
	//
	// event
	ResponseTriggerType *string `json:"ResponseTriggerType,omitempty" xml:"ResponseTriggerType,omitempty"`
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
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
