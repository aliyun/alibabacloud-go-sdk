// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumePhysicalNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *ResumePhysicalNodeRequest
	GetEnv() *string
	SetOpTenantId(v int64) *ResumePhysicalNodeRequest
	GetOpTenantId() *int64
	SetResumeCommand(v *ResumePhysicalNodeRequestResumeCommand) *ResumePhysicalNodeRequest
	GetResumeCommand() *ResumePhysicalNodeRequestResumeCommand
}

type ResumePhysicalNodeRequest struct {
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
	ResumeCommand *ResumePhysicalNodeRequestResumeCommand `json:"ResumeCommand,omitempty" xml:"ResumeCommand,omitempty" type:"Struct"`
}

func (s ResumePhysicalNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumePhysicalNodeRequest) GoString() string {
	return s.String()
}

func (s *ResumePhysicalNodeRequest) GetEnv() *string {
	return s.Env
}

func (s *ResumePhysicalNodeRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ResumePhysicalNodeRequest) GetResumeCommand() *ResumePhysicalNodeRequestResumeCommand {
	return s.ResumeCommand
}

func (s *ResumePhysicalNodeRequest) SetEnv(v string) *ResumePhysicalNodeRequest {
	s.Env = &v
	return s
}

func (s *ResumePhysicalNodeRequest) SetOpTenantId(v int64) *ResumePhysicalNodeRequest {
	s.OpTenantId = &v
	return s
}

func (s *ResumePhysicalNodeRequest) SetResumeCommand(v *ResumePhysicalNodeRequestResumeCommand) *ResumePhysicalNodeRequest {
	s.ResumeCommand = v
	return s
}

func (s *ResumePhysicalNodeRequest) Validate() error {
	if s.ResumeCommand != nil {
		if err := s.ResumeCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResumePhysicalNodeRequestResumeCommand struct {
	// This parameter is required.
	NodeIdList []*string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 102011
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ResumePhysicalNodeRequestResumeCommand) String() string {
	return dara.Prettify(s)
}

func (s ResumePhysicalNodeRequestResumeCommand) GoString() string {
	return s.String()
}

func (s *ResumePhysicalNodeRequestResumeCommand) GetNodeIdList() []*string {
	return s.NodeIdList
}

func (s *ResumePhysicalNodeRequestResumeCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ResumePhysicalNodeRequestResumeCommand) SetNodeIdList(v []*string) *ResumePhysicalNodeRequestResumeCommand {
	s.NodeIdList = v
	return s
}

func (s *ResumePhysicalNodeRequestResumeCommand) SetProjectId(v int64) *ResumePhysicalNodeRequestResumeCommand {
	s.ProjectId = &v
	return s
}

func (s *ResumePhysicalNodeRequestResumeCommand) Validate() error {
	return dara.Validate(s)
}
