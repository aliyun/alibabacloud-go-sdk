// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentVersion(v int32) *GetAgentRequest
	GetAgentVersion() *int32
	SetName(v string) *GetAgentRequest
	GetName() *string
}

type GetAgentRequest struct {
	// If omitted, the current definition is retrieved. If a published version is specified, the corresponding snapshot and the fixed versions of official skills and suite artifacts are retrieved. If the current unpublished version is specified, the draft is returned without fixed artifact versions. This is not the POP system parameter Version.
	//
	// example:
	//
	// 2
	AgentVersion *int32 `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// The name of the skill to retrieve.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-skill
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentRequest) GoString() string {
	return s.String()
}

func (s *GetAgentRequest) GetAgentVersion() *int32 {
	return s.AgentVersion
}

func (s *GetAgentRequest) GetName() *string {
	return s.Name
}

func (s *GetAgentRequest) SetAgentVersion(v int32) *GetAgentRequest {
	s.AgentVersion = &v
	return s
}

func (s *GetAgentRequest) SetName(v string) *GetAgentRequest {
	s.Name = &v
	return s
}

func (s *GetAgentRequest) Validate() error {
	return dara.Validate(s)
}
