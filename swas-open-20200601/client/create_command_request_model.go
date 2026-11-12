// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandContent(v string) *CreateCommandRequest
	GetCommandContent() *string
	SetDescription(v string) *CreateCommandRequest
	GetDescription() *string
	SetEnableParameter(v bool) *CreateCommandRequest
	GetEnableParameter() *bool
	SetName(v string) *CreateCommandRequest
	GetName() *string
	SetRegionId(v string) *CreateCommandRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateCommandRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateCommandRequestTag) *CreateCommandRequest
	GetTag() []*CreateCommandRequestTag
	SetTimeout(v int64) *CreateCommandRequest
	GetTimeout() *int64
	SetType(v string) *CreateCommandRequest
	GetType() *string
	SetWorkingDir(v string) *CreateCommandRequest
	GetWorkingDir() *string
}

type CreateCommandRequest struct {
	// The content of the command. Take note of the following items:
	//
	// - If you set the EnableParameter parameter to true, you can use custom parameters in the command content.
	//
	// - Custom parameters are enclosed in braces ({{}}). Spaces and line breaks before and after the parameter names within the braces are ignored.
	//
	// - You can specify a maximum of 20 custom parameters.
	//
	// - A custom parameter name can contain only letters, digits, hyphens (-), and underscores (_). Parameter names are not case-sensitive.
	//
	// - A custom parameter name cannot exceed 64 bytes in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// ifconfig -s
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The description of the command. All characters are supported. The description cannot exceed 512 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the created command uses custom parameters.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The name of the command. All characters are supported. The name cannot exceed 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. You can call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags. You can specify up to 20 tags.
	Tag []*CreateCommandRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The timeout period for running the command on the instance, in seconds.
	//
	// If the task that runs the command times out, Cloud Assistant forcibly terminates the task process. Valid values: 10 to 86400 (24 hours).
	//
	// Default value: 60.
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The language type of the command. Valid values:
	//
	// - RunBatScript: bat commands for Windows instances
	//
	// - RunPowerShellScript: PowerShell commands for Windows instances
	//
	// - RunShellScript: shell commands for Linux instances
	//
	// This parameter is required.
	//
	// example:
	//
	// RunShellScript
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The directory in which the created command is run on the ECS instance.
	//
	// Default value:
	//
	// - Linux instances: the home directory of the root user, which is /root.
	//
	// - Windows instances: the directory where the Cloud Assistant client process is located, such as C:\\Windows\\System32.
	//
	// example:
	//
	// /root/
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s CreateCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCommandRequest) GoString() string {
	return s.String()
}

func (s *CreateCommandRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *CreateCommandRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCommandRequest) GetEnableParameter() *bool {
	return s.EnableParameter
}

func (s *CreateCommandRequest) GetName() *string {
	return s.Name
}

func (s *CreateCommandRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCommandRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateCommandRequest) GetTag() []*CreateCommandRequestTag {
	return s.Tag
}

func (s *CreateCommandRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *CreateCommandRequest) GetType() *string {
	return s.Type
}

func (s *CreateCommandRequest) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *CreateCommandRequest) SetCommandContent(v string) *CreateCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *CreateCommandRequest) SetDescription(v string) *CreateCommandRequest {
	s.Description = &v
	return s
}

func (s *CreateCommandRequest) SetEnableParameter(v bool) *CreateCommandRequest {
	s.EnableParameter = &v
	return s
}

func (s *CreateCommandRequest) SetName(v string) *CreateCommandRequest {
	s.Name = &v
	return s
}

func (s *CreateCommandRequest) SetRegionId(v string) *CreateCommandRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCommandRequest) SetResourceGroupId(v string) *CreateCommandRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateCommandRequest) SetTag(v []*CreateCommandRequestTag) *CreateCommandRequest {
	s.Tag = v
	return s
}

func (s *CreateCommandRequest) SetTimeout(v int64) *CreateCommandRequest {
	s.Timeout = &v
	return s
}

func (s *CreateCommandRequest) SetType(v string) *CreateCommandRequest {
	s.Type = &v
	return s
}

func (s *CreateCommandRequest) SetWorkingDir(v string) *CreateCommandRequest {
	s.WorkingDir = &v
	return s
}

func (s *CreateCommandRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCommandRequestTag struct {
	// The key of tag N that is added to the command. Valid values of N: 1 to 20.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length, and cannot start with aliyun or acs:, or contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is added to the command. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. The tag value can be up to 64 characters in length, and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCommandRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateCommandRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCommandRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateCommandRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateCommandRequestTag) SetKey(v string) *CreateCommandRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCommandRequestTag) SetValue(v string) *CreateCommandRequestTag {
	s.Value = &v
	return s
}

func (s *CreateCommandRequestTag) Validate() error {
	return dara.Validate(s)
}
