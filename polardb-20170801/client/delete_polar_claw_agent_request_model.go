// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarClawAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *DeletePolarClawAgentRequest
	GetAgentId() *string
	SetApplicationId(v string) *DeletePolarClawAgentRequest
	GetApplicationId() *string
	SetDeleteFiles(v bool) *DeletePolarClawAgentRequest
	GetDeleteFiles() *bool
}

type DeletePolarClawAgentRequest struct {
	// The ID of the agent to delete. This parameter cannot be set to `main`.
	//
	// This parameter is required.
	//
	// example:
	//
	// work
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Specifies whether to delete the working directory and session files. Default value: `true`.
	//
	// example:
	//
	// true
	DeleteFiles *bool `json:"DeleteFiles,omitempty" xml:"DeleteFiles,omitempty"`
}

func (s DeletePolarClawAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarClawAgentRequest) GoString() string {
	return s.String()
}

func (s *DeletePolarClawAgentRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *DeletePolarClawAgentRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeletePolarClawAgentRequest) GetDeleteFiles() *bool {
	return s.DeleteFiles
}

func (s *DeletePolarClawAgentRequest) SetAgentId(v string) *DeletePolarClawAgentRequest {
	s.AgentId = &v
	return s
}

func (s *DeletePolarClawAgentRequest) SetApplicationId(v string) *DeletePolarClawAgentRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeletePolarClawAgentRequest) SetDeleteFiles(v bool) *DeletePolarClawAgentRequest {
	s.DeleteFiles = &v
	return s
}

func (s *DeletePolarClawAgentRequest) Validate() error {
	return dara.Validate(s)
}
