// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *GetAgentByIdRequest
	GetAgentId() *int64
	SetInstanceId(v string) *GetAgentByIdRequest
	GetInstanceId() *string
}

type GetAgentByIdRequest struct {
	// This parameter is required.
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAgentByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentByIdRequest) GoString() string {
	return s.String()
}

func (s *GetAgentByIdRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *GetAgentByIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAgentByIdRequest) SetAgentId(v int64) *GetAgentByIdRequest {
	s.AgentId = &v
	return s
}

func (s *GetAgentByIdRequest) SetInstanceId(v string) *GetAgentByIdRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentByIdRequest) Validate() error {
	return dara.Validate(s)
}
