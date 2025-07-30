// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskAssignRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *CreateTaskAssignRuleRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *CreateTaskAssignRuleRequest
	GetJsonStr() *string
}

type CreateTaskAssignRuleRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateTaskAssignRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskAssignRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskAssignRuleRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *CreateTaskAssignRuleRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *CreateTaskAssignRuleRequest) SetBaseMeAgentId(v int64) *CreateTaskAssignRuleRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateTaskAssignRuleRequest) SetJsonStr(v string) *CreateTaskAssignRuleRequest {
	s.JsonStr = &v
	return s
}

func (s *CreateTaskAssignRuleRequest) Validate() error {
	return dara.Validate(s)
}
