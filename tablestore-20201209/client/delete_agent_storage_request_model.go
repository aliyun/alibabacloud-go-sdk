// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentStorageName(v string) *DeleteAgentStorageRequest
	GetAgentStorageName() *string
}

type DeleteAgentStorageRequest struct {
	// agent storage name
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-test
	AgentStorageName *string `json:"AgentStorageName,omitempty" xml:"AgentStorageName,omitempty"`
}

func (s DeleteAgentStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentStorageRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentStorageRequest) GetAgentStorageName() *string {
	return s.AgentStorageName
}

func (s *DeleteAgentStorageRequest) SetAgentStorageName(v string) *DeleteAgentStorageRequest {
	s.AgentStorageName = &v
	return s
}

func (s *DeleteAgentStorageRequest) Validate() error {
	return dara.Validate(s)
}
