// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetMcpRequest
	GetId() *string
	SetInstanceId(v string) *GetMcpRequest
	GetInstanceId() *string
}

type GetMcpRequest struct {
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AgentTeams
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetMcpRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcpRequest) GoString() string {
	return s.String()
}

func (s *GetMcpRequest) GetId() *string {
	return s.Id
}

func (s *GetMcpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetMcpRequest) SetId(v string) *GetMcpRequest {
	s.Id = &v
	return s
}

func (s *GetMcpRequest) SetInstanceId(v string) *GetMcpRequest {
	s.InstanceId = &v
	return s
}

func (s *GetMcpRequest) Validate() error {
	return dara.Validate(s)
}
