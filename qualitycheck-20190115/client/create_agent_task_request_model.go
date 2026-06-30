// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *CreateAgentTaskRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *CreateAgentTaskRequest
	GetJsonStr() *string
}

type CreateAgentTaskRequest struct {
	// The ID of the business space.
	//
	// example:
	//
	// 12345
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// The complete **JSON string*	- information. For details, see the following sections.
	//
	// example:
	//
	// {\\"\\":\\"\\"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateAgentTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentTaskRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *CreateAgentTaskRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *CreateAgentTaskRequest) SetBaseMeAgentId(v int64) *CreateAgentTaskRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateAgentTaskRequest) SetJsonStr(v string) *CreateAgentTaskRequest {
	s.JsonStr = &v
	return s
}

func (s *CreateAgentTaskRequest) Validate() error {
	return dara.Validate(s)
}
