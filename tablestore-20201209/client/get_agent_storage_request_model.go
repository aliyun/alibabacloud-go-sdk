// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentStorageName(v string) *GetAgentStorageRequest
	GetAgentStorageName() *string
}

type GetAgentStorageRequest struct {
	// The name of the agent storage.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-test
	AgentStorageName *string `json:"AgentStorageName,omitempty" xml:"AgentStorageName,omitempty"`
}

func (s GetAgentStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentStorageRequest) GoString() string {
	return s.String()
}

func (s *GetAgentStorageRequest) GetAgentStorageName() *string {
	return s.AgentStorageName
}

func (s *GetAgentStorageRequest) SetAgentStorageName(v string) *GetAgentStorageRequest {
	s.AgentStorageName = &v
	return s
}

func (s *GetAgentStorageRequest) Validate() error {
	return dara.Validate(s)
}
