// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreatePipelineNodeCommand(v *CreatePipelineNodeRequestCreatePipelineNodeCommand) *CreatePipelineNodeRequest
	GetCreatePipelineNodeCommand() *CreatePipelineNodeRequestCreatePipelineNodeCommand
	SetOpTenantId(v int64) *CreatePipelineNodeRequest
	GetOpTenantId() *int64
}

type CreatePipelineNodeRequest struct {
	// This parameter is required.
	CreatePipelineNodeCommand *CreatePipelineNodeRequestCreatePipelineNodeCommand `json:"CreatePipelineNodeCommand,omitempty" xml:"CreatePipelineNodeCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreatePipelineNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineNodeRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineNodeRequest) GetCreatePipelineNodeCommand() *CreatePipelineNodeRequestCreatePipelineNodeCommand {
	return s.CreatePipelineNodeCommand
}

func (s *CreatePipelineNodeRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreatePipelineNodeRequest) SetCreatePipelineNodeCommand(v *CreatePipelineNodeRequestCreatePipelineNodeCommand) *CreatePipelineNodeRequest {
	s.CreatePipelineNodeCommand = v
	return s
}

func (s *CreatePipelineNodeRequest) SetOpTenantId(v int64) *CreatePipelineNodeRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreatePipelineNodeRequest) Validate() error {
	if s.CreatePipelineNodeCommand != nil {
		if err := s.CreatePipelineNodeCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePipelineNodeRequestCreatePipelineNodeCommand struct {
	// This parameter is required.
	FileInfo *CreatePipelineNodeRequestCreatePipelineNodeCommandFileInfo `json:"FileInfo,omitempty" xml:"FileInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// REAL_TIME
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_pipeline
	PipelineName *string `json:"PipelineName,omitempty" xml:"PipelineName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// REAL_TIME_PIPELINE
	PipelineType *string `json:"PipelineType,omitempty" xml:"PipelineType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7091124176569088
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreatePipelineNodeRequestCreatePipelineNodeCommand) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineNodeRequestCreatePipelineNodeCommand) GoString() string {
	return s.String()
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommand) GetFileInfo() *CreatePipelineNodeRequestCreatePipelineNodeCommandFileInfo {
	return s.FileInfo
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommand) GetNodeType() *string {
	return s.NodeType
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommand) GetPipelineName() *string {
	return s.PipelineName
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommand) GetPipelineType() *string {
	return s.PipelineType
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommand) SetFileInfo(v *CreatePipelineNodeRequestCreatePipelineNodeCommandFileInfo) *CreatePipelineNodeRequestCreatePipelineNodeCommand {
	s.FileInfo = v
	return s
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommand) SetNodeType(v string) *CreatePipelineNodeRequestCreatePipelineNodeCommand {
	s.NodeType = &v
	return s
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommand) SetPipelineName(v string) *CreatePipelineNodeRequestCreatePipelineNodeCommand {
	s.PipelineName = &v
	return s
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommand) SetPipelineType(v string) *CreatePipelineNodeRequestCreatePipelineNodeCommand {
	s.PipelineType = &v
	return s
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommand) SetProjectId(v int64) *CreatePipelineNodeRequestCreatePipelineNodeCommand {
	s.ProjectId = &v
	return s
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommand) Validate() error {
	if s.FileInfo != nil {
		if err := s.FileInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePipelineNodeRequestCreatePipelineNodeCommandFileInfo struct {
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// test
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_pipeline
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s CreatePipelineNodeRequestCreatePipelineNodeCommandFileInfo) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineNodeRequestCreatePipelineNodeCommandFileInfo) GoString() string {
	return s.String()
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommandFileInfo) GetDescription() *string {
	return s.Description
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommandFileInfo) GetDirectory() *string {
	return s.Directory
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommandFileInfo) GetFileName() *string {
	return s.FileName
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommandFileInfo) SetDescription(v string) *CreatePipelineNodeRequestCreatePipelineNodeCommandFileInfo {
	s.Description = &v
	return s
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommandFileInfo) SetDirectory(v string) *CreatePipelineNodeRequestCreatePipelineNodeCommandFileInfo {
	s.Directory = &v
	return s
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommandFileInfo) SetFileName(v string) *CreatePipelineNodeRequestCreatePipelineNodeCommandFileInfo {
	s.FileName = &v
	return s
}

func (s *CreatePipelineNodeRequestCreatePipelineNodeCommandFileInfo) Validate() error {
	return dara.Validate(s)
}
