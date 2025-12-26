// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlow interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *Flow
	GetAccessibility() *string
	SetCodeModeRunInfo(v string) *Flow
	GetCodeModeRunInfo() *string
	SetCreateFrom(v string) *Flow
	GetCreateFrom() *string
	SetCreator(v string) *Flow
	GetCreator() *string
	SetDescription(v string) *Flow
	GetDescription() *string
	SetFlowId(v string) *Flow
	GetFlowId() *string
	SetFlowName(v string) *Flow
	GetFlowName() *string
	SetFlowStoragePath(v string) *Flow
	GetFlowStoragePath() *string
	SetFlowTemplateId(v string) *Flow
	GetFlowTemplateId() *string
	SetFlowType(v string) *Flow
	GetFlowType() *string
	SetGmtCreateTime(v string) *Flow
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *Flow
	GetGmtModifiedTime() *string
	SetRuntime(v *FlowRuntime) *Flow
	GetRuntime() *FlowRuntime
	SetRuntimeId(v string) *Flow
	GetRuntimeId() *string
	SetSourceUri(v string) *Flow
	GetSourceUri() *string
	SetVersion(v string) *Flow
	GetVersion() *string
	SetWorkDir(v string) *Flow
	GetWorkDir() *string
	SetWorkspaceId(v string) *Flow
	GetWorkspaceId() *string
}

type Flow struct {
	Accessibility   *string      `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	CodeModeRunInfo *string      `json:"CodeModeRunInfo,omitempty" xml:"CodeModeRunInfo,omitempty"`
	CreateFrom      *string      `json:"CreateFrom,omitempty" xml:"CreateFrom,omitempty"`
	Creator         *string      `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description     *string      `json:"Description,omitempty" xml:"Description,omitempty"`
	FlowId          *string      `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	FlowName        *string      `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	FlowStoragePath *string      `json:"FlowStoragePath,omitempty" xml:"FlowStoragePath,omitempty"`
	FlowTemplateId  *string      `json:"FlowTemplateId,omitempty" xml:"FlowTemplateId,omitempty"`
	FlowType        *string      `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	GmtCreateTime   *string      `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string      `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Runtime         *FlowRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
	RuntimeId       *string      `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	SourceUri       *string      `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	Version         *string      `json:"Version,omitempty" xml:"Version,omitempty"`
	WorkDir         *string      `json:"WorkDir,omitempty" xml:"WorkDir,omitempty"`
	WorkspaceId     *string      `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Flow) String() string {
	return dara.Prettify(s)
}

func (s Flow) GoString() string {
	return s.String()
}

func (s *Flow) GetAccessibility() *string {
	return s.Accessibility
}

func (s *Flow) GetCodeModeRunInfo() *string {
	return s.CodeModeRunInfo
}

func (s *Flow) GetCreateFrom() *string {
	return s.CreateFrom
}

func (s *Flow) GetCreator() *string {
	return s.Creator
}

func (s *Flow) GetDescription() *string {
	return s.Description
}

func (s *Flow) GetFlowId() *string {
	return s.FlowId
}

func (s *Flow) GetFlowName() *string {
	return s.FlowName
}

func (s *Flow) GetFlowStoragePath() *string {
	return s.FlowStoragePath
}

func (s *Flow) GetFlowTemplateId() *string {
	return s.FlowTemplateId
}

func (s *Flow) GetFlowType() *string {
	return s.FlowType
}

func (s *Flow) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Flow) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *Flow) GetRuntime() *FlowRuntime {
	return s.Runtime
}

func (s *Flow) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *Flow) GetSourceUri() *string {
	return s.SourceUri
}

func (s *Flow) GetVersion() *string {
	return s.Version
}

func (s *Flow) GetWorkDir() *string {
	return s.WorkDir
}

func (s *Flow) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *Flow) SetAccessibility(v string) *Flow {
	s.Accessibility = &v
	return s
}

func (s *Flow) SetCodeModeRunInfo(v string) *Flow {
	s.CodeModeRunInfo = &v
	return s
}

func (s *Flow) SetCreateFrom(v string) *Flow {
	s.CreateFrom = &v
	return s
}

func (s *Flow) SetCreator(v string) *Flow {
	s.Creator = &v
	return s
}

func (s *Flow) SetDescription(v string) *Flow {
	s.Description = &v
	return s
}

func (s *Flow) SetFlowId(v string) *Flow {
	s.FlowId = &v
	return s
}

func (s *Flow) SetFlowName(v string) *Flow {
	s.FlowName = &v
	return s
}

func (s *Flow) SetFlowStoragePath(v string) *Flow {
	s.FlowStoragePath = &v
	return s
}

func (s *Flow) SetFlowTemplateId(v string) *Flow {
	s.FlowTemplateId = &v
	return s
}

func (s *Flow) SetFlowType(v string) *Flow {
	s.FlowType = &v
	return s
}

func (s *Flow) SetGmtCreateTime(v string) *Flow {
	s.GmtCreateTime = &v
	return s
}

func (s *Flow) SetGmtModifiedTime(v string) *Flow {
	s.GmtModifiedTime = &v
	return s
}

func (s *Flow) SetRuntime(v *FlowRuntime) *Flow {
	s.Runtime = v
	return s
}

func (s *Flow) SetRuntimeId(v string) *Flow {
	s.RuntimeId = &v
	return s
}

func (s *Flow) SetSourceUri(v string) *Flow {
	s.SourceUri = &v
	return s
}

func (s *Flow) SetVersion(v string) *Flow {
	s.Version = &v
	return s
}

func (s *Flow) SetWorkDir(v string) *Flow {
	s.WorkDir = &v
	return s
}

func (s *Flow) SetWorkspaceId(v string) *Flow {
	s.WorkspaceId = &v
	return s
}

func (s *Flow) Validate() error {
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FlowRuntime struct {
	// 运行时ID
	RuntimeId *string `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	// 运行时名称
	RuntimeName *string `json:"RuntimeName,omitempty" xml:"RuntimeName,omitempty"`
	// 运行时状态
	RuntimeStatus *string `json:"RuntimeStatus,omitempty" xml:"RuntimeStatus,omitempty"`
	// 运行时类型
	RuntimeType *string `json:"RuntimeType,omitempty" xml:"RuntimeType,omitempty"`
	// 是否支持SSE
	SupportSSEStream *bool `json:"SupportSSEStream,omitempty" xml:"SupportSSEStream,omitempty"`
}

func (s FlowRuntime) String() string {
	return dara.Prettify(s)
}

func (s FlowRuntime) GoString() string {
	return s.String()
}

func (s *FlowRuntime) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *FlowRuntime) GetRuntimeName() *string {
	return s.RuntimeName
}

func (s *FlowRuntime) GetRuntimeStatus() *string {
	return s.RuntimeStatus
}

func (s *FlowRuntime) GetRuntimeType() *string {
	return s.RuntimeType
}

func (s *FlowRuntime) GetSupportSSEStream() *bool {
	return s.SupportSSEStream
}

func (s *FlowRuntime) SetRuntimeId(v string) *FlowRuntime {
	s.RuntimeId = &v
	return s
}

func (s *FlowRuntime) SetRuntimeName(v string) *FlowRuntime {
	s.RuntimeName = &v
	return s
}

func (s *FlowRuntime) SetRuntimeStatus(v string) *FlowRuntime {
	s.RuntimeStatus = &v
	return s
}

func (s *FlowRuntime) SetRuntimeType(v string) *FlowRuntime {
	s.RuntimeType = &v
	return s
}

func (s *FlowRuntime) SetSupportSSEStream(v bool) *FlowRuntime {
	s.SupportSSEStream = &v
	return s
}

func (s *FlowRuntime) Validate() error {
	return dara.Validate(s)
}
