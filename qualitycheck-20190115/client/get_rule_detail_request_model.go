// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetRuleDetailRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetRuleDetailRequest
	GetJsonStr() *string
}

type GetRuleDetailRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"ruleIds":"123"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetRuleDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailRequest) GoString() string {
	return s.String()
}

func (s *GetRuleDetailRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetRuleDetailRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetRuleDetailRequest) SetBaseMeAgentId(v int64) *GetRuleDetailRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetRuleDetailRequest) SetJsonStr(v string) *GetRuleDetailRequest {
	s.JsonStr = &v
	return s
}

func (s *GetRuleDetailRequest) Validate() error {
	return dara.Validate(s)
}
