// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteMcpRequest
	GetId() *string
	SetInstanceId(v string) *DeleteMcpRequest
	GetInstanceId() *string
}

type DeleteMcpRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// AgentTeams
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteMcpRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcpRequest) GetId() *string {
	return s.Id
}

func (s *DeleteMcpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteMcpRequest) SetId(v string) *DeleteMcpRequest {
	s.Id = &v
	return s
}

func (s *DeleteMcpRequest) SetInstanceId(v string) *DeleteMcpRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteMcpRequest) Validate() error {
	return dara.Validate(s)
}
