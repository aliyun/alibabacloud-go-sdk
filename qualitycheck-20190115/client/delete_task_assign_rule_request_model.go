// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskAssignRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DeleteTaskAssignRuleRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *DeleteTaskAssignRuleRequest
	GetJsonStr() *string
}

type DeleteTaskAssignRuleRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"ruleId": 24}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteTaskAssignRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskAssignRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteTaskAssignRuleRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DeleteTaskAssignRuleRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *DeleteTaskAssignRuleRequest) SetBaseMeAgentId(v int64) *DeleteTaskAssignRuleRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteTaskAssignRuleRequest) SetJsonStr(v string) *DeleteTaskAssignRuleRequest {
	s.JsonStr = &v
	return s
}

func (s *DeleteTaskAssignRuleRequest) Validate() error {
	return dara.Validate(s)
}
