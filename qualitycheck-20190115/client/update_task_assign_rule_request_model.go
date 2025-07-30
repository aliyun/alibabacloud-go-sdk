// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskAssignRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateTaskAssignRuleRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateTaskAssignRuleRequest
	GetJsonStr() *string
}

type UpdateTaskAssignRuleRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateTaskAssignRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAssignRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskAssignRuleRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateTaskAssignRuleRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateTaskAssignRuleRequest) SetBaseMeAgentId(v int64) *UpdateTaskAssignRuleRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateTaskAssignRuleRequest) SetJsonStr(v string) *UpdateTaskAssignRuleRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateTaskAssignRuleRequest) Validate() error {
	return dara.Validate(s)
}
