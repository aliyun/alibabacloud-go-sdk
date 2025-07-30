// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetRuleRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetRuleRequest
	GetJsonStr() *string
}

type GetRuleRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"ruleIds":"123"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRuleRequest) GoString() string {
	return s.String()
}

func (s *GetRuleRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetRuleRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetRuleRequest) SetBaseMeAgentId(v int64) *GetRuleRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetRuleRequest) SetJsonStr(v string) *GetRuleRequest {
	s.JsonStr = &v
	return s
}

func (s *GetRuleRequest) Validate() error {
	return dara.Validate(s)
}
