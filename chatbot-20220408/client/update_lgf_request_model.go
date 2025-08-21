// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLgfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateLgfRequest
	GetAgentKey() *string
	SetInstanceId(v string) *UpdateLgfRequest
	GetInstanceId() *string
	SetLgfDefinition(v *UpdateLgfRequestLgfDefinition) *UpdateLgfRequest
	GetLgfDefinition() *UpdateLgfRequestLgfDefinition
	SetLgfId(v int64) *UpdateLgfRequest
	GetLgfId() *int64
}

type UpdateLgfRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	LgfDefinition *UpdateLgfRequestLgfDefinition `json:"LgfDefinition,omitempty" xml:"LgfDefinition,omitempty" type:"Struct"`
	// LGF ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 12121
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
}

func (s UpdateLgfRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLgfRequest) GoString() string {
	return s.String()
}

func (s *UpdateLgfRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateLgfRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateLgfRequest) GetLgfDefinition() *UpdateLgfRequestLgfDefinition {
	return s.LgfDefinition
}

func (s *UpdateLgfRequest) GetLgfId() *int64 {
	return s.LgfId
}

func (s *UpdateLgfRequest) SetAgentKey(v string) *UpdateLgfRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateLgfRequest) SetInstanceId(v string) *UpdateLgfRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLgfRequest) SetLgfDefinition(v *UpdateLgfRequestLgfDefinition) *UpdateLgfRequest {
	s.LgfDefinition = v
	return s
}

func (s *UpdateLgfRequest) SetLgfId(v int64) *UpdateLgfRequest {
	s.LgfId = &v
	return s
}

func (s *UpdateLgfRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateLgfRequestLgfDefinition struct {
	// This parameter is required.
	//
	// example:
	//
	// 23234523522
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// This parameter is required.
	RuleText *string `json:"RuleText,omitempty" xml:"RuleText,omitempty"`
}

func (s UpdateLgfRequestLgfDefinition) String() string {
	return dara.Prettify(s)
}

func (s UpdateLgfRequestLgfDefinition) GoString() string {
	return s.String()
}

func (s *UpdateLgfRequestLgfDefinition) GetIntentId() *int64 {
	return s.IntentId
}

func (s *UpdateLgfRequestLgfDefinition) GetRuleText() *string {
	return s.RuleText
}

func (s *UpdateLgfRequestLgfDefinition) SetIntentId(v int64) *UpdateLgfRequestLgfDefinition {
	s.IntentId = &v
	return s
}

func (s *UpdateLgfRequestLgfDefinition) SetRuleText(v string) *UpdateLgfRequestLgfDefinition {
	s.RuleText = &v
	return s
}

func (s *UpdateLgfRequestLgfDefinition) Validate() error {
	return dara.Validate(s)
}
