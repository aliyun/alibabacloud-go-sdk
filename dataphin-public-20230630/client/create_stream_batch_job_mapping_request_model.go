// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamBatchJobMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *CreateStreamBatchJobMappingRequest
	GetOpTenantId() *int64
	SetStreamBatchJobMappingCreateCommand(v *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) *CreateStreamBatchJobMappingRequest
	GetStreamBatchJobMappingCreateCommand() *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand
}

type CreateStreamBatchJobMappingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	StreamBatchJobMappingCreateCommand *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand `json:"StreamBatchJobMappingCreateCommand,omitempty" xml:"StreamBatchJobMappingCreateCommand,omitempty" type:"Struct"`
}

func (s CreateStreamBatchJobMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamBatchJobMappingRequest) GoString() string {
	return s.String()
}

func (s *CreateStreamBatchJobMappingRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateStreamBatchJobMappingRequest) GetStreamBatchJobMappingCreateCommand() *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand {
	return s.StreamBatchJobMappingCreateCommand
}

func (s *CreateStreamBatchJobMappingRequest) SetOpTenantId(v int64) *CreateStreamBatchJobMappingRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateStreamBatchJobMappingRequest) SetStreamBatchJobMappingCreateCommand(v *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) *CreateStreamBatchJobMappingRequest {
	s.StreamBatchJobMappingCreateCommand = v
	return s
}

func (s *CreateStreamBatchJobMappingRequest) Validate() error {
	return dara.Validate(s)
}

type CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 61187014-a3ba-4cdd-8609-1f0aa3df4a3d
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 这是一段任务的描述信息
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /karel
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vvr-8.0.9-flink-1.17
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// karel_hover_3
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FLINK_SQL
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7081229106458752
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default-queue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PREJOB
	VvpClusterType *string `json:"VvpClusterType,omitempty" xml:"VvpClusterType,omitempty"`
}

func (s CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) GetDirectory() *string {
	return s.Directory
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) GetEnv() *string {
	return s.Env
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) GetFileName() *string {
	return s.FileName
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) GetFileType() *string {
	return s.FileType
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) GetVvpClusterType() *string {
	return s.VvpClusterType
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) SetClusterId(v string) *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand {
	s.ClusterId = &v
	return s
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) SetDescription(v string) *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) SetDirectory(v string) *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand {
	s.Directory = &v
	return s
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) SetEngineVersion(v string) *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand {
	s.EngineVersion = &v
	return s
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) SetEnv(v string) *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand {
	s.Env = &v
	return s
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) SetFileName(v string) *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand {
	s.FileName = &v
	return s
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) SetFileType(v string) *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand {
	s.FileType = &v
	return s
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) SetProjectId(v int64) *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand {
	s.ProjectId = &v
	return s
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) SetQueueName(v string) *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand {
	s.QueueName = &v
	return s
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) SetVvpClusterType(v string) *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand {
	s.VvpClusterType = &v
	return s
}

func (s *CreateStreamBatchJobMappingRequestStreamBatchJobMappingCreateCommand) Validate() error {
	return dara.Validate(s)
}
