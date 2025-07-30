// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvalidRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *InvalidRuleRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *InvalidRuleRequest
	GetJsonStr() *string
}

type InvalidRuleRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"ruleIds":[3,4]}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s InvalidRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s InvalidRuleRequest) GoString() string {
	return s.String()
}

func (s *InvalidRuleRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *InvalidRuleRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *InvalidRuleRequest) SetBaseMeAgentId(v int64) *InvalidRuleRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *InvalidRuleRequest) SetJsonStr(v string) *InvalidRuleRequest {
	s.JsonStr = &v
	return s
}

func (s *InvalidRuleRequest) Validate() error {
	return dara.Validate(s)
}
