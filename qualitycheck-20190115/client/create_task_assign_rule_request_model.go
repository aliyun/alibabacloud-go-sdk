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
	// Workspace ID.
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// A complete JSON string. For details, see the parameter descriptions below.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"callType":"0","durationMin":1,"durationMax":300,"agents":[{"agentId":"202526561358712105","agentName":"agent"}],"rules":[{"rid":15659}],"reviewers":[{"reviewerId":"255746168704895558","reviewerName":"0917质检员"},{"reviewerId":"268370362815185444","reviewerName":"0710质检员"}],"skillGroups":[{"skillName":"客服组"}],"priority":5}
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
