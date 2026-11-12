// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCommandShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandContent(v string) *RunCommandShrinkRequest
	GetCommandContent() *string
	SetEnableParameter(v bool) *RunCommandShrinkRequest
	GetEnableParameter() *bool
	SetInstanceId(v string) *RunCommandShrinkRequest
	GetInstanceId() *string
	SetName(v string) *RunCommandShrinkRequest
	GetName() *string
	SetParametersShrink(v string) *RunCommandShrinkRequest
	GetParametersShrink() *string
	SetRegionId(v string) *RunCommandShrinkRequest
	GetRegionId() *string
	SetTimeout(v int32) *RunCommandShrinkRequest
	GetTimeout() *int32
	SetType(v string) *RunCommandShrinkRequest
	GetType() *string
	SetWindowsPasswordName(v string) *RunCommandShrinkRequest
	GetWindowsPasswordName() *string
	SetWorkingDir(v string) *RunCommandShrinkRequest
	GetWorkingDir() *string
	SetWorkingUser(v string) *RunCommandShrinkRequest
	GetWorkingUser() *string
}

type RunCommandShrinkRequest struct {
	// The command content. Note the following items:
	//
	// - Set `EnableParameter=true` to enable the custom parameter feature in the command content:
	//
	// - Define custom parameters by enclosing them in `{{}}`. Spaces and line breaks before and after the parameter name within `{{}}` are ignored.
	//
	// - The number of custom parameters cannot exceed 20.
	//
	// - Custom parameter names can contain only letters (a-z, A-Z), digits (0-9), hyphens (-), and underscores (_). Other characters are not supported. Parameter names are case-insensitive.
	//
	// - A single custom parameter name cannot exceed 64 bytes.
	//
	// This parameter is required.
	//
	// example:
	//
	// ifconfig -s
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// Specifies whether the command contains custom parameters.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The instance ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the command execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The key-value pairs of custom parameters to pass in when the command contains custom parameters. For example, if the command content is `echo {{name}}`, you can use the `Parameters` parameter to pass in the key-value pair `{"name":"Jack"}`. The custom parameter automatically replaces the variable value `name`, and the actual command that runs is `echo Jack`.
	//
	// The number of custom parameters ranges from 0 to 20. Note the following items:
	//
	// - Keys cannot be empty strings and can contain up to 64 characters.
	//
	// - Values can be empty strings.
	//
	// - The set of custom parameter names must be a subset of the parameter set defined when the command was created. For parameters that are not passed in, you can use empty strings as substitutes.
	//
	// Default value: empty, which indicates that this parameter is unset to disable custom parameters.
	//
	// example:
	//
	// {"name":"Jack", "accessKey":"LTAI****************"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. You can call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to query the latest region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The timeout period for the command to run on the server instance. Unit: seconds.
	//
	// When a command execution task times out, Command Assistant forcefully terminates the task process. Valid values: 10 to 86400 (24 hours).
	//
	// Default value: 60.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The language type of the command. Valid values:
	//
	// - RunBatScript: Bat commands for Windows instances.
	//
	// - RunPowerShellScript: PowerShell commands for Windows instances.
	//
	// - RunShellScript: Shell commands for Linux instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// RunShellScript
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The password name of the user that runs the command on a Windows instance.
	//
	// If you want to run a command as a non-default user (System) on a Windows instance, you must specify both WorkingUser and this parameter. To reduce the risk of password leaks, store the plaintext password in the parameter repository of CloudOps Orchestration Service and pass in only the password name here.
	//
	// example:
	//
	// axtSecretPassword
	WindowsPasswordName *string `json:"WindowsPasswordName,omitempty" xml:"WindowsPasswordName,omitempty"`
	// The custom execution path of the command. Default paths:
	//
	// - Linux instances: The default execution path is the /root directory of the root user.
	//
	// - Windows instances: The default execution path is the C:\\Windows\\system32 directory.
	//
	// example:
	//
	// /root
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	// The user that you want to use to run the command on the simple application server instance. We recommend that you use a regular user to run commands to reduce security risks.
	//
	// Default values:
	//
	// - root: On Linux instances, commands are run as the root user by default.
	//
	// - system: On Windows instances, commands are run as the system user by default.
	//
	// example:
	//
	// root
	WorkingUser *string `json:"WorkingUser,omitempty" xml:"WorkingUser,omitempty"`
}

func (s RunCommandShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunCommandShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunCommandShrinkRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *RunCommandShrinkRequest) GetEnableParameter() *bool {
	return s.EnableParameter
}

func (s *RunCommandShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RunCommandShrinkRequest) GetName() *string {
	return s.Name
}

func (s *RunCommandShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *RunCommandShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunCommandShrinkRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *RunCommandShrinkRequest) GetType() *string {
	return s.Type
}

func (s *RunCommandShrinkRequest) GetWindowsPasswordName() *string {
	return s.WindowsPasswordName
}

func (s *RunCommandShrinkRequest) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *RunCommandShrinkRequest) GetWorkingUser() *string {
	return s.WorkingUser
}

func (s *RunCommandShrinkRequest) SetCommandContent(v string) *RunCommandShrinkRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandShrinkRequest) SetEnableParameter(v bool) *RunCommandShrinkRequest {
	s.EnableParameter = &v
	return s
}

func (s *RunCommandShrinkRequest) SetInstanceId(v string) *RunCommandShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *RunCommandShrinkRequest) SetName(v string) *RunCommandShrinkRequest {
	s.Name = &v
	return s
}

func (s *RunCommandShrinkRequest) SetParametersShrink(v string) *RunCommandShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *RunCommandShrinkRequest) SetRegionId(v string) *RunCommandShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RunCommandShrinkRequest) SetTimeout(v int32) *RunCommandShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandShrinkRequest) SetType(v string) *RunCommandShrinkRequest {
	s.Type = &v
	return s
}

func (s *RunCommandShrinkRequest) SetWindowsPasswordName(v string) *RunCommandShrinkRequest {
	s.WindowsPasswordName = &v
	return s
}

func (s *RunCommandShrinkRequest) SetWorkingDir(v string) *RunCommandShrinkRequest {
	s.WorkingDir = &v
	return s
}

func (s *RunCommandShrinkRequest) SetWorkingUser(v string) *RunCommandShrinkRequest {
	s.WorkingUser = &v
	return s
}

func (s *RunCommandShrinkRequest) Validate() error {
	return dara.Validate(s)
}
