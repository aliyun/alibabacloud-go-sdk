// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPausePhysicalNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *PausePhysicalNodeRequest
	GetEnv() *string
	SetOpTenantId(v int64) *PausePhysicalNodeRequest
	GetOpTenantId() *int64
	SetPauseCommand(v *PausePhysicalNodeRequestPauseCommand) *PausePhysicalNodeRequest
	GetPauseCommand() *PausePhysicalNodeRequestPauseCommand
}

type PausePhysicalNodeRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	PauseCommand *PausePhysicalNodeRequestPauseCommand `json:"PauseCommand,omitempty" xml:"PauseCommand,omitempty" type:"Struct"`
}

func (s PausePhysicalNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s PausePhysicalNodeRequest) GoString() string {
	return s.String()
}

func (s *PausePhysicalNodeRequest) GetEnv() *string {
	return s.Env
}

func (s *PausePhysicalNodeRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *PausePhysicalNodeRequest) GetPauseCommand() *PausePhysicalNodeRequestPauseCommand {
	return s.PauseCommand
}

func (s *PausePhysicalNodeRequest) SetEnv(v string) *PausePhysicalNodeRequest {
	s.Env = &v
	return s
}

func (s *PausePhysicalNodeRequest) SetOpTenantId(v int64) *PausePhysicalNodeRequest {
	s.OpTenantId = &v
	return s
}

func (s *PausePhysicalNodeRequest) SetPauseCommand(v *PausePhysicalNodeRequestPauseCommand) *PausePhysicalNodeRequest {
	s.PauseCommand = v
	return s
}

func (s *PausePhysicalNodeRequest) Validate() error {
	if s.PauseCommand != nil {
		if err := s.PauseCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PausePhysicalNodeRequestPauseCommand struct {
	// This parameter is required.
	NodeIdList []*string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 13222210
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PausePhysicalNodeRequestPauseCommand) String() string {
	return dara.Prettify(s)
}

func (s PausePhysicalNodeRequestPauseCommand) GoString() string {
	return s.String()
}

func (s *PausePhysicalNodeRequestPauseCommand) GetNodeIdList() []*string {
	return s.NodeIdList
}

func (s *PausePhysicalNodeRequestPauseCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *PausePhysicalNodeRequestPauseCommand) SetNodeIdList(v []*string) *PausePhysicalNodeRequestPauseCommand {
	s.NodeIdList = v
	return s
}

func (s *PausePhysicalNodeRequestPauseCommand) SetProjectId(v int64) *PausePhysicalNodeRequestPauseCommand {
	s.ProjectId = &v
	return s
}

func (s *PausePhysicalNodeRequestPauseCommand) Validate() error {
	return dara.Validate(s)
}
