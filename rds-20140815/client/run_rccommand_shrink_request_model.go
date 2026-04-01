// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunRCCommandShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RunRCCommandShrinkRequest
	GetClientToken() *string
	SetCommandContent(v string) *RunRCCommandShrinkRequest
	GetCommandContent() *string
	SetContainerId(v string) *RunRCCommandShrinkRequest
	GetContainerId() *string
	SetContainerName(v string) *RunRCCommandShrinkRequest
	GetContainerName() *string
	SetContentEncoding(v string) *RunRCCommandShrinkRequest
	GetContentEncoding() *string
	SetDescription(v string) *RunRCCommandShrinkRequest
	GetDescription() *string
	SetEnableParameter(v bool) *RunRCCommandShrinkRequest
	GetEnableParameter() *bool
	SetFrequency(v string) *RunRCCommandShrinkRequest
	GetFrequency() *string
	SetInstanceIdsShrink(v string) *RunRCCommandShrinkRequest
	GetInstanceIdsShrink() *string
	SetKeepCommand(v bool) *RunRCCommandShrinkRequest
	GetKeepCommand() *bool
	SetLauncher(v string) *RunRCCommandShrinkRequest
	GetLauncher() *string
	SetName(v string) *RunRCCommandShrinkRequest
	GetName() *string
	SetParametersShrink(v string) *RunRCCommandShrinkRequest
	GetParametersShrink() *string
	SetRegionId(v string) *RunRCCommandShrinkRequest
	GetRegionId() *string
	SetRepeatMode(v string) *RunRCCommandShrinkRequest
	GetRepeatMode() *string
	SetResourceGroupId(v string) *RunRCCommandShrinkRequest
	GetResourceGroupId() *string
	SetResourceTagsShrink(v string) *RunRCCommandShrinkRequest
	GetResourceTagsShrink() *string
	SetTagsShrink(v string) *RunRCCommandShrinkRequest
	GetTagsShrink() *string
	SetTerminationMode(v string) *RunRCCommandShrinkRequest
	GetTerminationMode() *string
	SetTimeout(v int64) *RunRCCommandShrinkRequest
	GetTimeout() *int64
	SetType(v string) *RunRCCommandShrinkRequest
	GetType() *string
	SetUsername(v string) *RunRCCommandShrinkRequest
	GetUsername() *string
	SetWindowsPasswordName(v string) *RunRCCommandShrinkRequest
	GetWindowsPasswordName() *string
	SetWorkingDir(v string) *RunRCCommandShrinkRequest
	GetWorkingDir() *string
}

type RunRCCommandShrinkRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOCziJZNwH****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ZWNobyAxMjM=
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// example:
	//
	// ab141ddfbacfe02d9dbc25966ed971536124527097398d419a6746873fea****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// example:
	//
	// test-container
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	EnableParameter   *bool   `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	Frequency         *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// false
	KeepCommand *bool `json:"KeepCommand,omitempty" xml:"KeepCommand,omitempty"`
	// example:
	//
	// python3 -u {{ACS::ScriptFileName|Ext(".py")}}
	Launcher *string `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {"name":"Jack", "accessKey":"LTAI*************"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId    *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceTagsShrink *string `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty"`
	TagsShrink         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// ProcessTree
	TerminationMode *string `json:"TerminationMode,omitempty" xml:"TerminationMode,omitempty"`
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RunShellScript
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// example:
	//
	// None
	WindowsPasswordName *string `json:"WindowsPasswordName,omitempty" xml:"WindowsPasswordName,omitempty"`
	// example:
	//
	// /home/user
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s RunRCCommandShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunRCCommandShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunRCCommandShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunRCCommandShrinkRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *RunRCCommandShrinkRequest) GetContainerId() *string {
	return s.ContainerId
}

func (s *RunRCCommandShrinkRequest) GetContainerName() *string {
	return s.ContainerName
}

func (s *RunRCCommandShrinkRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *RunRCCommandShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *RunRCCommandShrinkRequest) GetEnableParameter() *bool {
	return s.EnableParameter
}

func (s *RunRCCommandShrinkRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *RunRCCommandShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *RunRCCommandShrinkRequest) GetKeepCommand() *bool {
	return s.KeepCommand
}

func (s *RunRCCommandShrinkRequest) GetLauncher() *string {
	return s.Launcher
}

func (s *RunRCCommandShrinkRequest) GetName() *string {
	return s.Name
}

func (s *RunRCCommandShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *RunRCCommandShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunRCCommandShrinkRequest) GetRepeatMode() *string {
	return s.RepeatMode
}

func (s *RunRCCommandShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RunRCCommandShrinkRequest) GetResourceTagsShrink() *string {
	return s.ResourceTagsShrink
}

func (s *RunRCCommandShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *RunRCCommandShrinkRequest) GetTerminationMode() *string {
	return s.TerminationMode
}

func (s *RunRCCommandShrinkRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *RunRCCommandShrinkRequest) GetType() *string {
	return s.Type
}

func (s *RunRCCommandShrinkRequest) GetUsername() *string {
	return s.Username
}

func (s *RunRCCommandShrinkRequest) GetWindowsPasswordName() *string {
	return s.WindowsPasswordName
}

func (s *RunRCCommandShrinkRequest) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *RunRCCommandShrinkRequest) SetClientToken(v string) *RunRCCommandShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetCommandContent(v string) *RunRCCommandShrinkRequest {
	s.CommandContent = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetContainerId(v string) *RunRCCommandShrinkRequest {
	s.ContainerId = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetContainerName(v string) *RunRCCommandShrinkRequest {
	s.ContainerName = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetContentEncoding(v string) *RunRCCommandShrinkRequest {
	s.ContentEncoding = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetDescription(v string) *RunRCCommandShrinkRequest {
	s.Description = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetEnableParameter(v bool) *RunRCCommandShrinkRequest {
	s.EnableParameter = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetFrequency(v string) *RunRCCommandShrinkRequest {
	s.Frequency = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetInstanceIdsShrink(v string) *RunRCCommandShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetKeepCommand(v bool) *RunRCCommandShrinkRequest {
	s.KeepCommand = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetLauncher(v string) *RunRCCommandShrinkRequest {
	s.Launcher = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetName(v string) *RunRCCommandShrinkRequest {
	s.Name = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetParametersShrink(v string) *RunRCCommandShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetRegionId(v string) *RunRCCommandShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetRepeatMode(v string) *RunRCCommandShrinkRequest {
	s.RepeatMode = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetResourceGroupId(v string) *RunRCCommandShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetResourceTagsShrink(v string) *RunRCCommandShrinkRequest {
	s.ResourceTagsShrink = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetTagsShrink(v string) *RunRCCommandShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetTerminationMode(v string) *RunRCCommandShrinkRequest {
	s.TerminationMode = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetTimeout(v int64) *RunRCCommandShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetType(v string) *RunRCCommandShrinkRequest {
	s.Type = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetUsername(v string) *RunRCCommandShrinkRequest {
	s.Username = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetWindowsPasswordName(v string) *RunRCCommandShrinkRequest {
	s.WindowsPasswordName = &v
	return s
}

func (s *RunRCCommandShrinkRequest) SetWorkingDir(v string) *RunRCCommandShrinkRequest {
	s.WorkingDir = &v
	return s
}

func (s *RunRCCommandShrinkRequest) Validate() error {
	return dara.Validate(s)
}
