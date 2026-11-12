// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandContent(v string) *RunCommandRequest
	GetCommandContent() *string
	SetEnableParameter(v bool) *RunCommandRequest
	GetEnableParameter() *bool
	SetInstanceId(v string) *RunCommandRequest
	GetInstanceId() *string
	SetName(v string) *RunCommandRequest
	GetName() *string
	SetParameters(v map[string]interface{}) *RunCommandRequest
	GetParameters() map[string]interface{}
	SetRegionId(v string) *RunCommandRequest
	GetRegionId() *string
	SetTimeout(v int32) *RunCommandRequest
	GetTimeout() *int32
	SetType(v string) *RunCommandRequest
	GetType() *string
	SetWindowsPasswordName(v string) *RunCommandRequest
	GetWindowsPasswordName() *string
	SetWorkingDir(v string) *RunCommandRequest
	GetWorkingDir() *string
	SetWorkingUser(v string) *RunCommandRequest
	GetWorkingUser() *string
}

type RunCommandRequest struct {
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
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

func (s RunCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s RunCommandRequest) GoString() string {
	return s.String()
}

func (s *RunCommandRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *RunCommandRequest) GetEnableParameter() *bool {
	return s.EnableParameter
}

func (s *RunCommandRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RunCommandRequest) GetName() *string {
	return s.Name
}

func (s *RunCommandRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *RunCommandRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunCommandRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *RunCommandRequest) GetType() *string {
	return s.Type
}

func (s *RunCommandRequest) GetWindowsPasswordName() *string {
	return s.WindowsPasswordName
}

func (s *RunCommandRequest) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *RunCommandRequest) GetWorkingUser() *string {
	return s.WorkingUser
}

func (s *RunCommandRequest) SetCommandContent(v string) *RunCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandRequest) SetEnableParameter(v bool) *RunCommandRequest {
	s.EnableParameter = &v
	return s
}

func (s *RunCommandRequest) SetInstanceId(v string) *RunCommandRequest {
	s.InstanceId = &v
	return s
}

func (s *RunCommandRequest) SetName(v string) *RunCommandRequest {
	s.Name = &v
	return s
}

func (s *RunCommandRequest) SetParameters(v map[string]interface{}) *RunCommandRequest {
	s.Parameters = v
	return s
}

func (s *RunCommandRequest) SetRegionId(v string) *RunCommandRequest {
	s.RegionId = &v
	return s
}

func (s *RunCommandRequest) SetTimeout(v int32) *RunCommandRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandRequest) SetType(v string) *RunCommandRequest {
	s.Type = &v
	return s
}

func (s *RunCommandRequest) SetWindowsPasswordName(v string) *RunCommandRequest {
	s.WindowsPasswordName = &v
	return s
}

func (s *RunCommandRequest) SetWorkingDir(v string) *RunCommandRequest {
	s.WorkingDir = &v
	return s
}

func (s *RunCommandRequest) SetWorkingUser(v string) *RunCommandRequest {
	s.WorkingUser = &v
	return s
}

func (s *RunCommandRequest) Validate() error {
	return dara.Validate(s)
}
