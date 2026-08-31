// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCommandShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RunCommandShrinkRequest
	GetClientToken() *string
	SetCommandContent(v string) *RunCommandShrinkRequest
	GetCommandContent() *string
	SetContainerId(v string) *RunCommandShrinkRequest
	GetContainerId() *string
	SetContainerName(v string) *RunCommandShrinkRequest
	GetContainerName() *string
	SetContentEncoding(v string) *RunCommandShrinkRequest
	GetContentEncoding() *string
	SetDescription(v string) *RunCommandShrinkRequest
	GetDescription() *string
	SetEnableParameter(v bool) *RunCommandShrinkRequest
	GetEnableParameter() *bool
	SetFrequency(v string) *RunCommandShrinkRequest
	GetFrequency() *string
	SetInstanceId(v []*string) *RunCommandShrinkRequest
	GetInstanceId() []*string
	SetKeepCommand(v bool) *RunCommandShrinkRequest
	GetKeepCommand() *bool
	SetLauncher(v string) *RunCommandShrinkRequest
	GetLauncher() *string
	SetName(v string) *RunCommandShrinkRequest
	GetName() *string
	SetOssOutputDelivery(v string) *RunCommandShrinkRequest
	GetOssOutputDelivery() *string
	SetOwnerAccount(v string) *RunCommandShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RunCommandShrinkRequest
	GetOwnerId() *int64
	SetParametersShrink(v string) *RunCommandShrinkRequest
	GetParametersShrink() *string
	SetRegionId(v string) *RunCommandShrinkRequest
	GetRegionId() *string
	SetRepeatMode(v string) *RunCommandShrinkRequest
	GetRepeatMode() *string
	SetResourceGroupId(v string) *RunCommandShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *RunCommandShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RunCommandShrinkRequest
	GetResourceOwnerId() *int64
	SetResourceTag(v []*RunCommandShrinkRequestResourceTag) *RunCommandShrinkRequest
	GetResourceTag() []*RunCommandShrinkRequestResourceTag
	SetTag(v []*RunCommandShrinkRequestTag) *RunCommandShrinkRequest
	GetTag() []*RunCommandShrinkRequestTag
	SetTerminationMode(v string) *RunCommandShrinkRequest
	GetTerminationMode() *string
	SetTimed(v bool) *RunCommandShrinkRequest
	GetTimed() *bool
	SetTimeout(v int64) *RunCommandShrinkRequest
	GetTimeout() *int64
	SetType(v string) *RunCommandShrinkRequest
	GetType() *string
	SetUsername(v string) *RunCommandShrinkRequest
	GetUsername() *string
	SetWindowsPasswordName(v string) *RunCommandShrinkRequest
	GetWindowsPasswordName() *string
	SetWorkingDir(v string) *RunCommandShrinkRequest
	GetWorkingDir() *string
}

type RunCommandShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The command content. The command content can be plaintext or Base64-encoded. Note the following items:
	//
	// - The command content cannot exceed 24 KB after Base64 encoding. You can use `KeepCommand` to specify whether to retain the command.
	//
	// - If the command content is Base64-encoded, you must set `ContentEncoding=Base64`.
	//
	// - When `EnableParameter=true` is specified, the custom parameter feature is enabled in the command content:
	//
	//     - Define custom parameters by enclosing them in `{{}}`. Spaces and line breaks before and after the parameter name within `{{}}` are ignored.
	//
	//     - The number of custom parameters cannot exceed 20.
	//
	//     - Custom parameter names can contain a-zA-Z0-9-_ combinations. The acs:: prefix for specifying non-built-in environment parameters is not supported. Other characters are not supported. Parameter names are case-insensitive.
	//
	//     - Each custom parameter name cannot exceed 64 bytes.
	//
	// - You can specify built-in environment parameters as custom parameters. When running the command, you do not need to manually assign values to these parameters because Cloud Assistant automatically replaces them with the corresponding values. The following built-in environment parameters are supported:
	//
	//     - `{{ACS::RegionId}}`: The region ID.
	//
	//     - `{{ACS::AccountId}}`: The Alibaba Cloud account ID.
	//
	//     - `{{ACS::InstanceId}}`: The instance ID. When a command is sent to multiple instances and you want to use `{{ACS::InstanceId}}` as a built-in environment parameter, ensure that the Cloud Assistant Agent version is no earlier than:
	//
	//         - Linux: 2.2.3.309
	//
	//         - Windows: 2.1.3.309
	//
	//     - `{{ACS::InstanceName}}`: The instance name. When a command is sent to multiple instances and you want to use `{{ACS::InstanceName}}` as a built-in environment parameter, ensure that the Cloud Assistant Agent version is no earlier than:
	//
	//         - Linux: 2.2.3.344
	//
	//         - Windows: 2.1.3.344
	//
	//     - `{{ACS::InvokeId}}`: The invocation ID. To use `{{ACS::InvokeId}}` as a built-in environment parameter, ensure that the Cloud Assistant Agent version is no earlier than:
	//
	//         - Linux: 2.2.3.309
	//
	//         - Windows: 2.1.3.309
	//
	//     - `{{ACS::CommandId}}`: The command ID. When running a command by calling this operation and you want to use `{{ACS::CommandId}}` as a built-in environment parameter, ensure that the Cloud Assistant Agent version is no earlier than:
	//
	//         - Linux: 2.2.3.309
	//
	//         - Windows: 2.1.3.309
	//
	// This parameter is required.
	//
	// example:
	//
	// ZWNobyAxMjM=
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The container ID. Only 64-bit hexadecimal strings are supported. The `docker://`, `containerd://`, or `cri-o://` prefix can be used to explicitly specify the container runtime.
	//
	// Precautions:
	//
	// - If this parameter is specified, Cloud Assistant runs the script in the specified container on the instance.
	//
	// - If this parameter is specified, the command can only be run on Linux instances with Cloud Assistant Agent version 2.2.3.344 or later.
	//
	// - If this parameter is specified, the `Username` and `WorkingDir` parameters do not take effect. The command is run only as the default container user in the default working directory of the container. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// > Only Shell scripts are supported in Linux containers. Specifying an interpreter at the beginning of the script in the format of `#!/usr/bin/python` is not supported. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// example:
	//
	// ab141ddfbacfe02d9dbc25966ed971536124527097398d419a6746873fea****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The container name.
	//
	// Precautions:
	//
	// - If this parameter is specified, Cloud Assistant runs the script in the specified container on the instance.
	//
	// - If this parameter is specified, the command can only be run on Linux instances with Cloud Assistant Agent version 2.2.3.344 or later.
	//
	// - If this parameter is specified, the `Username` and `WorkingDir` parameters do not take effect. The command is run only as the default container user in the default working directory of the container. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// > Only Shell scripts are supported in Linux containers. Specifying an interpreter at the beginning of the script in the format of `#!/usr/bin/python` is not supported. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// example:
	//
	// test-container
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The encoding method of the command content (`CommandContent`). Valid values (case-insensitive):
	//
	// - PlainText: no encoding. The content is transmitted in plaintext.
	//
	// - Base64: Base64 encoding.
	//
	// Default value: PlainText. Invalid values are treated as PlainText.
	//
	// example:
	//
	// Base64
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The command description. All character sets are supported. The description cannot exceed 512 characters in length.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the command contains custom parameters.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The schedule for running the command. Three scheduling methods are supported: execution at fixed intervals (based on Rate expressions), one-time execution at a specified time, and clock-based scheduled execution (based on Cron expressions).
	//
	// - Execution at fixed intervals: Based on Rate expressions, the command is run at the specified interval. The interval can be specified in seconds (s), minutes (m), hours (h), or days (d). This method is applicable to scenarios where tasks are run at fixed intervals. Format: `rate(<interval value><interval unit>)`. For example, to run a command every 5 minutes, use `rate(5m)`. The following limits apply to fixed-interval execution:
	//
	//     - The interval must be no greater than 7 days and no less than 60 seconds, and must be greater than the timeout period of the scheduled task.
	//
	//     - The interval is based on a fixed frequency and is not related to the actual execution time of the task. For example, if a command is set to run every 5 minutes and the task takes 2 minutes to complete, the next round starts 3 minutes after the task is completed.
	//
	//     - The task is not run immediately upon creation. For example, if a command is set to run every 5 minutes, the command is not run immediately when the task is created. Instead, execution starts 5 minutes after the task is created.
	//
	// - One-time execution at a specified time: The command is run once at the specified time zone and time point. Format: `at(yyyy-MM-dd HH:mm:ss <time zone>)`. If no time zone is specified, UTC is used by default. The time zone supports the following three formats:
	//
	//     - Full time zone name: such as `Asia/Shanghai` (China/Shanghai time) or `America/Los_Angeles` (US/Los Angeles time).
	//
	//     - Time zone offset from Greenwich Mean Time: such as `GMT+8:00` (East 8th time zone) or `GMT-7:00` (West 7th time zone). When using the GMT format, leading zeros are not supported in the hour field.
	//
	//     - Time zone abbreviation: Only UTC (Coordinated Universal Time) is supported.
	//
	//   For example, to run a command once at 13:15:30 on June 6, 2022 in China/Shanghai time, use: `at(2022-06-06 13:15:30 Asia/Shanghai)`. To run a command once at 13:15:30 on June 6, 2022 in the West 7th time zone, use: `at(2022-06-06 13:15:30 GMT-7:00)`.
	//
	// - Clock-based scheduled execution (based on Cron expressions): Based on Cron expressions, the command is run according to the scheduled task settings. Format: `<seconds> <minutes> <hours> <day of month> <month> <day of week> <year (optional)> <time zone>`, i.e., `<Cron expression> <time zone>`. The scheduled task execution time is calculated based on the Cron expression in the specified time zone. If no time zone is specified, the system time zone of the instance running the scheduled task is used by default. For more information about Cron expressions, see [Cron expressions](https://help.aliyun.com/document_detail/64769.html). The time zone supports the following three formats:
	//
	//     - Full time zone name: such as `Asia/Shanghai` (China/Shanghai time) or `America/Los_Angeles` (US/Los Angeles time).
	//
	//     - Time zone offset from Greenwich Mean Time: such as `GMT+8:00` (East 8th time zone) or `GMT-7:00` (West 7th time zone). When using the GMT format, leading zeros are not supported in the hour field.
	//
	//     - Time zone abbreviation: Only UTC (Coordinated Universal Time) is supported.
	//
	//   For example, to run a command once every day at 10:15 AM in China/Shanghai time in 2022, use `0 15 10 ? 	- 	- 2022 Asia/Shanghai`. To run a command every 30 minutes from 10:00 AM to 11:30 AM every day in the East 8th time zone in 2022, use `0 0/30 10-11 	- 	- ? 2022 GMT+8:00`. To run a command every 5 minutes from 2:00 PM to 2:55 PM every day in October every two years starting from 2022 in UTC, use `0 0/5 14 	- 10 ? 2022/2 UTC`.
	//
	//     > The minimum interval must be greater than or equal to the timeout period of the scheduled task and no less than 10 seconds.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The instance ID array of ECS instances. Array length: 1 to 100.
	//
	// If any of the specified instances does not meet the execution conditions, you must reselect the instances.
	//
	// You can also request a quota increase in Quota Center (quota name: Maximum number of instances supported for command execute).
	//
	// example:
	//
	// i-bp185dy2o3o6neg****
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// Specifies whether to retain the command after execution. Valid values:
	//
	// - true: retains the command. The command can be run again by calling InvokeCommand. This counts toward the Cloud Assistant command retention quota.
	//
	// - false: does not retain the command. The command is automatically deleted after execution and does not count toward the Cloud Assistant command retention quota.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	KeepCommand *bool `json:"KeepCommand,omitempty" xml:"KeepCommand,omitempty"`
	// The bootstrap program for script execution. The value cannot exceed 1 KB in length.
	//
	// example:
	//
	// python3 -u {{ACS::ScriptFileName|Ext(".py")}}
	Launcher *string `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	// The command name. All character sets are supported. The name cannot exceed 128 characters in length.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The OSS delivery configuration for command execution output.
	//
	// - Format: oss://${BucketName}/${Prefix}, where ${BucketName} is the name of the destination OSS bucket and ${Prefix} is the directory prefix for delivery.
	//
	// example:
	//
	// oss://testBucket/testPrefix
	OssOutputDelivery *string `json:"OssOutputDelivery,omitempty" xml:"OssOutputDelivery,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key-value pairs of custom parameters to pass in when running a command that contains custom parameters. For example, if the command content is `echo {{name}}`, you can pass in the key-value pair `{"name":"Jack"}` through the Parameter parameter. The custom parameter automatically replaces the variable value `name`, and the actual command executed is `echo Jack`.
	//
	// The number of custom parameters ranges from 0 to 10. Note the following items:
	//
	// - Keys cannot be empty strings and can contain up to 64 characters.
	//
	// - Values can be empty strings.
	//
	// - After custom parameters and the original command content are Base64-encoded, the total size cannot exceed 24 KB. You can use `KeepCommand` to specify whether to retain the command.
	//
	// - The set of custom parameter names must be a subset of the parameter set defined when the command was created. For parameters that are not passed in, you can use empty strings as substitutes.
	//
	// Default value: empty, which disables custom parameters.
	//
	// example:
	//
	// {"name":"Jack", "accessKey":"LTAI*************"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The execution mode of the command. Valid values:
	//
	// - Once: immediately runs the command.
	//
	// - Period: runs the command on a schedule. When this parameter is set to `Period`, you must also specify the `Frequency` parameter.
	//
	// - NextRebootOnly: automatically runs the command the next time the instance starts.
	//
	// - EveryReboot: automatically runs the command every time the instance starts.
	//
	// - DryRun: performs a dry run of the request without actually running the command. Checks include request parameters, instance execution environment, and Cloud Assistant Agent running status.
	//
	// Default values:
	//
	// - When the `Frequency` parameter is not specified, the default value is `Once`.
	//
	// - When the `Frequency` parameter is specified, the command is processed as `Period` regardless of whether this parameter is set.
	//
	// Precautions:
	//
	// - You can call [StopInvocation](https://help.aliyun.com/document_detail/64838.html) to stop a pending or scheduled command.
	//
	// - When this parameter is set to `Period` or `EveryReboot`, you can call [DescribeInvocationResults](https://help.aliyun.com/document_detail/64845.html) and specify `IncludeHistory=true` to view the execution history of the scheduled command.
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// The resource group ID for the command execution. When this parameter is specified:
	//
	// - If the ECS instance corresponding to InstanceId belongs to a non-default resource group, the ECS instance must belong to this resource group.
	//
	// - You can filter the corresponding command execution results by specifying this parameter (by calling [DescribeInvocations](https://help.aliyun.com/document_detail/64840.html) or [DescribeInvocationResults](https://help.aliyun.com/document_detail/64845.html)).
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags used to filter instances. Array length: 0 to 20. You can run commands in batches on instances with the same tags without specifying InstanceId.
	ResourceTag []*RunCommandShrinkRequestResourceTag `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty" type:"Repeated"`
	// The tag pairs. Array length: 0 to 20.
	Tag []*RunCommandShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The mode for stopping the task (manual stop or timeout interruption). Valid values:
	//
	// - Process: stops the current script process.
	//
	// - ProcessTree: stops the current process tree (the collection of the script process and all child processes it created).
	//
	// example:
	//
	// ProcessTree
	TerminationMode *string `json:"TerminationMode,omitempty" xml:"TerminationMode,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. Passing in this parameter has no effect.
	//
	// example:
	//
	// true
	Timed *bool `json:"Timed,omitempty" xml:"Timed,omitempty"`
	// The timeout period for command execution. Unit: seconds.
	//
	// A timeout occurs when a command cannot be run because of process issues, missing modules, or missing Cloud Assistant Agent. When a timeout occurs, the command process is forcefully terminated.
	//
	// Default value: 60.
	//
	// example:
	//
	// 3600
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The command type. Valid values:
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
	// The username for running the command on the ECS instance. The value cannot exceed 255 characters in length.
	//
	// - For Linux ECS instances, commands are run as the root user by default.
	//
	// - For Windows ECS instances, commands are run as the System user by default.
	//
	// You can also specify another existing user on the instance to run the command. Running Cloud Assistant commands as a regular user is more secure. For more information, see [Configure a regular user to run Cloud Assistant commands](https://help.aliyun.com/document_detail/203771.html).
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The name of the password for the user who executes the command on a Windows instance. The value cannot exceed 255 characters in length.
	//
	// When you want to execute a command as a non-default user (System) on a Windows instance, you must specify both `Username` and this parameter. To reduce the risk of password leaks, store the plaintext password in the parameter repository of operations management, and pass in only the password name here. For more information, see [Encryption parameters](https://help.aliyun.com/document_detail/186828.html) and [Settings for a regular user to execute Cloud Assistant commands](https://help.aliyun.com/document_detail/203771.html).
	//
	// > This parameter is not required when you execute commands as the root user on a Linux instance or the System user on a Windows instance.
	//
	// example:
	//
	// axtSecretPassword
	WindowsPasswordName *string `json:"WindowsPasswordName,omitempty" xml:"WindowsPasswordName,omitempty"`
	// The working directory of the command on the ECS instance. The value cannot exceed 200 characters in length.
	//
	// Default values:
	//
	// - For Linux instances, the default directory is the home directory of the root user, which is `/root`.
	//
	// - For Windows instances, the default directory is the directory where the Cloud Assistant Agent process is located, such as `C:\\Windows\\System32`.
	//
	// example:
	//
	// /home/user
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s RunCommandShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunCommandShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunCommandShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunCommandShrinkRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *RunCommandShrinkRequest) GetContainerId() *string {
	return s.ContainerId
}

func (s *RunCommandShrinkRequest) GetContainerName() *string {
	return s.ContainerName
}

func (s *RunCommandShrinkRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *RunCommandShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *RunCommandShrinkRequest) GetEnableParameter() *bool {
	return s.EnableParameter
}

func (s *RunCommandShrinkRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *RunCommandShrinkRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *RunCommandShrinkRequest) GetKeepCommand() *bool {
	return s.KeepCommand
}

func (s *RunCommandShrinkRequest) GetLauncher() *string {
	return s.Launcher
}

func (s *RunCommandShrinkRequest) GetName() *string {
	return s.Name
}

func (s *RunCommandShrinkRequest) GetOssOutputDelivery() *string {
	return s.OssOutputDelivery
}

func (s *RunCommandShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RunCommandShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RunCommandShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *RunCommandShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunCommandShrinkRequest) GetRepeatMode() *string {
	return s.RepeatMode
}

func (s *RunCommandShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RunCommandShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RunCommandShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RunCommandShrinkRequest) GetResourceTag() []*RunCommandShrinkRequestResourceTag {
	return s.ResourceTag
}

func (s *RunCommandShrinkRequest) GetTag() []*RunCommandShrinkRequestTag {
	return s.Tag
}

func (s *RunCommandShrinkRequest) GetTerminationMode() *string {
	return s.TerminationMode
}

func (s *RunCommandShrinkRequest) GetTimed() *bool {
	return s.Timed
}

func (s *RunCommandShrinkRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *RunCommandShrinkRequest) GetType() *string {
	return s.Type
}

func (s *RunCommandShrinkRequest) GetUsername() *string {
	return s.Username
}

func (s *RunCommandShrinkRequest) GetWindowsPasswordName() *string {
	return s.WindowsPasswordName
}

func (s *RunCommandShrinkRequest) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *RunCommandShrinkRequest) SetClientToken(v string) *RunCommandShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *RunCommandShrinkRequest) SetCommandContent(v string) *RunCommandShrinkRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandShrinkRequest) SetContainerId(v string) *RunCommandShrinkRequest {
	s.ContainerId = &v
	return s
}

func (s *RunCommandShrinkRequest) SetContainerName(v string) *RunCommandShrinkRequest {
	s.ContainerName = &v
	return s
}

func (s *RunCommandShrinkRequest) SetContentEncoding(v string) *RunCommandShrinkRequest {
	s.ContentEncoding = &v
	return s
}

func (s *RunCommandShrinkRequest) SetDescription(v string) *RunCommandShrinkRequest {
	s.Description = &v
	return s
}

func (s *RunCommandShrinkRequest) SetEnableParameter(v bool) *RunCommandShrinkRequest {
	s.EnableParameter = &v
	return s
}

func (s *RunCommandShrinkRequest) SetFrequency(v string) *RunCommandShrinkRequest {
	s.Frequency = &v
	return s
}

func (s *RunCommandShrinkRequest) SetInstanceId(v []*string) *RunCommandShrinkRequest {
	s.InstanceId = v
	return s
}

func (s *RunCommandShrinkRequest) SetKeepCommand(v bool) *RunCommandShrinkRequest {
	s.KeepCommand = &v
	return s
}

func (s *RunCommandShrinkRequest) SetLauncher(v string) *RunCommandShrinkRequest {
	s.Launcher = &v
	return s
}

func (s *RunCommandShrinkRequest) SetName(v string) *RunCommandShrinkRequest {
	s.Name = &v
	return s
}

func (s *RunCommandShrinkRequest) SetOssOutputDelivery(v string) *RunCommandShrinkRequest {
	s.OssOutputDelivery = &v
	return s
}

func (s *RunCommandShrinkRequest) SetOwnerAccount(v string) *RunCommandShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RunCommandShrinkRequest) SetOwnerId(v int64) *RunCommandShrinkRequest {
	s.OwnerId = &v
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

func (s *RunCommandShrinkRequest) SetRepeatMode(v string) *RunCommandShrinkRequest {
	s.RepeatMode = &v
	return s
}

func (s *RunCommandShrinkRequest) SetResourceGroupId(v string) *RunCommandShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RunCommandShrinkRequest) SetResourceOwnerAccount(v string) *RunCommandShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RunCommandShrinkRequest) SetResourceOwnerId(v int64) *RunCommandShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RunCommandShrinkRequest) SetResourceTag(v []*RunCommandShrinkRequestResourceTag) *RunCommandShrinkRequest {
	s.ResourceTag = v
	return s
}

func (s *RunCommandShrinkRequest) SetTag(v []*RunCommandShrinkRequestTag) *RunCommandShrinkRequest {
	s.Tag = v
	return s
}

func (s *RunCommandShrinkRequest) SetTerminationMode(v string) *RunCommandShrinkRequest {
	s.TerminationMode = &v
	return s
}

func (s *RunCommandShrinkRequest) SetTimed(v bool) *RunCommandShrinkRequest {
	s.Timed = &v
	return s
}

func (s *RunCommandShrinkRequest) SetTimeout(v int64) *RunCommandShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandShrinkRequest) SetType(v string) *RunCommandShrinkRequest {
	s.Type = &v
	return s
}

func (s *RunCommandShrinkRequest) SetUsername(v string) *RunCommandShrinkRequest {
	s.Username = &v
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

func (s *RunCommandShrinkRequest) Validate() error {
	if s.ResourceTag != nil {
		for _, item := range s.ResourceTag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type RunCommandShrinkRequestResourceTag struct {
	// The tag key used to filter instances.
	//
	// Precautions:
	//
	// - This parameter conflicts with the InstanceId parameter. They cannot be specified at the same time.
	//
	// - If this value is specified, it cannot be an empty string.
	//
	// - The number of instances under the tag cannot exceed the quantity limit of InstanceId.N. If the number of instances exceeds the limit, control the number of instances by adding batch tags, such as batch: b1.
	//
	// - The key can be up to 64 characters in length and cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value used to filter instances.
	//
	// Precautions:
	//
	// - The value can be an empty string.
	//
	// - The value can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunCommandShrinkRequestResourceTag) String() string {
	return dara.Prettify(s)
}

func (s RunCommandShrinkRequestResourceTag) GoString() string {
	return s.String()
}

func (s *RunCommandShrinkRequestResourceTag) GetKey() *string {
	return s.Key
}

func (s *RunCommandShrinkRequestResourceTag) GetValue() *string {
	return s.Value
}

func (s *RunCommandShrinkRequestResourceTag) SetKey(v string) *RunCommandShrinkRequestResourceTag {
	s.Key = &v
	return s
}

func (s *RunCommandShrinkRequestResourceTag) SetValue(v string) *RunCommandShrinkRequestResourceTag {
	s.Value = &v
	return s
}

func (s *RunCommandShrinkRequestResourceTag) Validate() error {
	return dara.Validate(s)
}

type RunCommandShrinkRequestTag struct {
	// The tag key of the command execute. If this value is specified, it cannot be an empty string.
	//
	// When you use a single tag to filter resources, the resource count under that tag cannot exceed 1,000. When you use multiple tags to filter resources, the resource count of resources that are attached to all specified tags cannot exceed 1,000. If the resource count exceeds 1,000, use the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to query resources.
	//
	// The key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the command execution. The value can be an empty string.
	//
	// The value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunCommandShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s RunCommandShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *RunCommandShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *RunCommandShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *RunCommandShrinkRequestTag) SetKey(v string) *RunCommandShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *RunCommandShrinkRequestTag) SetValue(v string) *RunCommandShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *RunCommandShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
