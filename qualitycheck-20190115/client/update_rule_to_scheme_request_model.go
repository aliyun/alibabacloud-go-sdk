// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleToSchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateRuleToSchemeRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateRuleToSchemeRequest
	GetJsonStr() *string
}

type UpdateRuleToSchemeRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"schemeId":"10","schemeRules":[{"ruleId":229,"checkType":0}]}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s UpdateRuleToSchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleToSchemeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRuleToSchemeRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateRuleToSchemeRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateRuleToSchemeRequest) SetBaseMeAgentId(v int64) *UpdateRuleToSchemeRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateRuleToSchemeRequest) SetJsonStr(v string) *UpdateRuleToSchemeRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateRuleToSchemeRequest) Validate() error {
	return dara.Validate(s)
}
