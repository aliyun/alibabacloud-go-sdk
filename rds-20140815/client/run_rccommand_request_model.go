// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunRCCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RunRCCommandRequest
	GetClientToken() *string
	SetCommandContent(v string) *RunRCCommandRequest
	GetCommandContent() *string
	SetContainerId(v string) *RunRCCommandRequest
	GetContainerId() *string
	SetContainerName(v string) *RunRCCommandRequest
	GetContainerName() *string
	SetContentEncoding(v string) *RunRCCommandRequest
	GetContentEncoding() *string
	SetDescription(v string) *RunRCCommandRequest
	GetDescription() *string
	SetEnableParameter(v bool) *RunRCCommandRequest
	GetEnableParameter() *bool
	SetFrequency(v string) *RunRCCommandRequest
	GetFrequency() *string
	SetInstanceIds(v []*string) *RunRCCommandRequest
	GetInstanceIds() []*string
	SetKeepCommand(v bool) *RunRCCommandRequest
	GetKeepCommand() *bool
	SetLauncher(v string) *RunRCCommandRequest
	GetLauncher() *string
	SetName(v string) *RunRCCommandRequest
	GetName() *string
	SetParameters(v map[string]interface{}) *RunRCCommandRequest
	GetParameters() map[string]interface{}
	SetRegionId(v string) *RunRCCommandRequest
	GetRegionId() *string
	SetRepeatMode(v string) *RunRCCommandRequest
	GetRepeatMode() *string
	SetResourceGroupId(v string) *RunRCCommandRequest
	GetResourceGroupId() *string
	SetResourceTags(v []*RunRCCommandRequestResourceTags) *RunRCCommandRequest
	GetResourceTags() []*RunRCCommandRequestResourceTags
	SetTags(v []*RunRCCommandRequestTags) *RunRCCommandRequest
	GetTags() []*RunRCCommandRequestTags
	SetTerminationMode(v string) *RunRCCommandRequest
	GetTerminationMode() *string
	SetTimeout(v int64) *RunRCCommandRequest
	GetTimeout() *int64
	SetType(v string) *RunRCCommandRequest
	GetType() *string
	SetUsername(v string) *RunRCCommandRequest
	GetUsername() *string
	SetWindowsPasswordName(v string) *RunRCCommandRequest
	GetWindowsPasswordName() *string
	SetWorkingDir(v string) *RunRCCommandRequest
	GetWorkingDir() *string
}

type RunRCCommandRequest struct {
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
	EnableParameter *bool     `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	Frequency       *string   `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
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
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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
	ResourceGroupId *string                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceTags    []*RunRCCommandRequestResourceTags `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty" type:"Repeated"`
	Tags            []*RunRCCommandRequestTags         `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s RunRCCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s RunRCCommandRequest) GoString() string {
	return s.String()
}

func (s *RunRCCommandRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunRCCommandRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *RunRCCommandRequest) GetContainerId() *string {
	return s.ContainerId
}

func (s *RunRCCommandRequest) GetContainerName() *string {
	return s.ContainerName
}

func (s *RunRCCommandRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *RunRCCommandRequest) GetDescription() *string {
	return s.Description
}

func (s *RunRCCommandRequest) GetEnableParameter() *bool {
	return s.EnableParameter
}

func (s *RunRCCommandRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *RunRCCommandRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RunRCCommandRequest) GetKeepCommand() *bool {
	return s.KeepCommand
}

func (s *RunRCCommandRequest) GetLauncher() *string {
	return s.Launcher
}

func (s *RunRCCommandRequest) GetName() *string {
	return s.Name
}

func (s *RunRCCommandRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *RunRCCommandRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunRCCommandRequest) GetRepeatMode() *string {
	return s.RepeatMode
}

func (s *RunRCCommandRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RunRCCommandRequest) GetResourceTags() []*RunRCCommandRequestResourceTags {
	return s.ResourceTags
}

func (s *RunRCCommandRequest) GetTags() []*RunRCCommandRequestTags {
	return s.Tags
}

func (s *RunRCCommandRequest) GetTerminationMode() *string {
	return s.TerminationMode
}

func (s *RunRCCommandRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *RunRCCommandRequest) GetType() *string {
	return s.Type
}

func (s *RunRCCommandRequest) GetUsername() *string {
	return s.Username
}

func (s *RunRCCommandRequest) GetWindowsPasswordName() *string {
	return s.WindowsPasswordName
}

func (s *RunRCCommandRequest) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *RunRCCommandRequest) SetClientToken(v string) *RunRCCommandRequest {
	s.ClientToken = &v
	return s
}

func (s *RunRCCommandRequest) SetCommandContent(v string) *RunRCCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *RunRCCommandRequest) SetContainerId(v string) *RunRCCommandRequest {
	s.ContainerId = &v
	return s
}

func (s *RunRCCommandRequest) SetContainerName(v string) *RunRCCommandRequest {
	s.ContainerName = &v
	return s
}

func (s *RunRCCommandRequest) SetContentEncoding(v string) *RunRCCommandRequest {
	s.ContentEncoding = &v
	return s
}

func (s *RunRCCommandRequest) SetDescription(v string) *RunRCCommandRequest {
	s.Description = &v
	return s
}

func (s *RunRCCommandRequest) SetEnableParameter(v bool) *RunRCCommandRequest {
	s.EnableParameter = &v
	return s
}

func (s *RunRCCommandRequest) SetFrequency(v string) *RunRCCommandRequest {
	s.Frequency = &v
	return s
}

func (s *RunRCCommandRequest) SetInstanceIds(v []*string) *RunRCCommandRequest {
	s.InstanceIds = v
	return s
}

func (s *RunRCCommandRequest) SetKeepCommand(v bool) *RunRCCommandRequest {
	s.KeepCommand = &v
	return s
}

func (s *RunRCCommandRequest) SetLauncher(v string) *RunRCCommandRequest {
	s.Launcher = &v
	return s
}

func (s *RunRCCommandRequest) SetName(v string) *RunRCCommandRequest {
	s.Name = &v
	return s
}

func (s *RunRCCommandRequest) SetParameters(v map[string]interface{}) *RunRCCommandRequest {
	s.Parameters = v
	return s
}

func (s *RunRCCommandRequest) SetRegionId(v string) *RunRCCommandRequest {
	s.RegionId = &v
	return s
}

func (s *RunRCCommandRequest) SetRepeatMode(v string) *RunRCCommandRequest {
	s.RepeatMode = &v
	return s
}

func (s *RunRCCommandRequest) SetResourceGroupId(v string) *RunRCCommandRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RunRCCommandRequest) SetResourceTags(v []*RunRCCommandRequestResourceTags) *RunRCCommandRequest {
	s.ResourceTags = v
	return s
}

func (s *RunRCCommandRequest) SetTags(v []*RunRCCommandRequestTags) *RunRCCommandRequest {
	s.Tags = v
	return s
}

func (s *RunRCCommandRequest) SetTerminationMode(v string) *RunRCCommandRequest {
	s.TerminationMode = &v
	return s
}

func (s *RunRCCommandRequest) SetTimeout(v int64) *RunRCCommandRequest {
	s.Timeout = &v
	return s
}

func (s *RunRCCommandRequest) SetType(v string) *RunRCCommandRequest {
	s.Type = &v
	return s
}

func (s *RunRCCommandRequest) SetUsername(v string) *RunRCCommandRequest {
	s.Username = &v
	return s
}

func (s *RunRCCommandRequest) SetWindowsPasswordName(v string) *RunRCCommandRequest {
	s.WindowsPasswordName = &v
	return s
}

func (s *RunRCCommandRequest) SetWorkingDir(v string) *RunRCCommandRequest {
	s.WorkingDir = &v
	return s
}

func (s *RunRCCommandRequest) Validate() error {
	if s.ResourceTags != nil {
		for _, item := range s.ResourceTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunRCCommandRequestResourceTags struct {
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunRCCommandRequestResourceTags) String() string {
	return dara.Prettify(s)
}

func (s RunRCCommandRequestResourceTags) GoString() string {
	return s.String()
}

func (s *RunRCCommandRequestResourceTags) GetKey() *string {
	return s.Key
}

func (s *RunRCCommandRequestResourceTags) GetValue() *string {
	return s.Value
}

func (s *RunRCCommandRequestResourceTags) SetKey(v string) *RunRCCommandRequestResourceTags {
	s.Key = &v
	return s
}

func (s *RunRCCommandRequestResourceTags) SetValue(v string) *RunRCCommandRequestResourceTags {
	s.Value = &v
	return s
}

func (s *RunRCCommandRequestResourceTags) Validate() error {
	return dara.Validate(s)
}

type RunRCCommandRequestTags struct {
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunRCCommandRequestTags) String() string {
	return dara.Prettify(s)
}

func (s RunRCCommandRequestTags) GoString() string {
	return s.String()
}

func (s *RunRCCommandRequestTags) GetKey() *string {
	return s.Key
}

func (s *RunRCCommandRequestTags) GetValue() *string {
	return s.Value
}

func (s *RunRCCommandRequestTags) SetKey(v string) *RunRCCommandRequestTags {
	s.Key = &v
	return s
}

func (s *RunRCCommandRequestTags) SetValue(v string) *RunRCCommandRequestTags {
	s.Value = &v
	return s
}

func (s *RunRCCommandRequestTags) Validate() error {
	return dara.Validate(s)
}
