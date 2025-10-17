// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLgfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateLgfRequest
	GetAgentKey() *string
	SetInstanceId(v string) *CreateLgfRequest
	GetInstanceId() *string
	SetLgfDefinition(v *CreateLgfRequestLgfDefinition) *CreateLgfRequest
	GetLgfDefinition() *CreateLgfRequestLgfDefinition
}

type CreateLgfRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId    *string                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LgfDefinition *CreateLgfRequestLgfDefinition `json:"LgfDefinition,omitempty" xml:"LgfDefinition,omitempty" type:"Struct"`
}

func (s CreateLgfRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLgfRequest) GoString() string {
	return s.String()
}

func (s *CreateLgfRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateLgfRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateLgfRequest) GetLgfDefinition() *CreateLgfRequestLgfDefinition {
	return s.LgfDefinition
}

func (s *CreateLgfRequest) SetAgentKey(v string) *CreateLgfRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateLgfRequest) SetInstanceId(v string) *CreateLgfRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateLgfRequest) SetLgfDefinition(v *CreateLgfRequestLgfDefinition) *CreateLgfRequest {
	s.LgfDefinition = v
	return s
}

func (s *CreateLgfRequest) Validate() error {
	if s.LgfDefinition != nil {
		if err := s.LgfDefinition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLgfRequestLgfDefinition struct {
	// This parameter is required.
	//
	// example:
	//
	// 4675678567
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// This parameter is required.
	RuleText *string `json:"RuleText,omitempty" xml:"RuleText,omitempty"`
}

func (s CreateLgfRequestLgfDefinition) String() string {
	return dara.Prettify(s)
}

func (s CreateLgfRequestLgfDefinition) GoString() string {
	return s.String()
}

func (s *CreateLgfRequestLgfDefinition) GetIntentId() *int64 {
	return s.IntentId
}

func (s *CreateLgfRequestLgfDefinition) GetRuleText() *string {
	return s.RuleText
}

func (s *CreateLgfRequestLgfDefinition) SetIntentId(v int64) *CreateLgfRequestLgfDefinition {
	s.IntentId = &v
	return s
}

func (s *CreateLgfRequestLgfDefinition) SetRuleText(v string) *CreateLgfRequestLgfDefinition {
	s.RuleText = &v
	return s
}

func (s *CreateLgfRequestLgfDefinition) Validate() error {
	return dara.Validate(s)
}
