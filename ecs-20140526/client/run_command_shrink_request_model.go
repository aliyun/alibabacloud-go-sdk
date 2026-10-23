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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The command content. The command content can be plaintext or Base64-encoded. Take note of the following items:
	//
	// - The size of the Base64-encoded command content cannot exceed 24 KB. You can use the `KeepCommand` parameter to specify whether to retain the command.
	//
	// - If the command content is Base64-encoded, you must set `ContentEncoding` to Base64.
	//
	// - If you set `EnableParameter` to true, you can use custom parameters in the command content:
	//
	//     - Use `{{}}` to enclose custom parameters. Spaces and line breaks before and after the parameter names within `{{}}` are ignored.
	//
	//     - You can define up to 20 custom parameters.
	//
	//     - A custom parameter name can contain letters, digits, hyphens (-), and underscores (_). The `acs::` prefix cannot be used to specify non-built-in environment variables. Other characters are not supported. Parameter names are not case-sensitive.
	//
	//     - A single custom parameter name cannot exceed 64 bytes in length.
	//
	// - You can specify built-in environment variables as custom parameters. When the command is run, you do not need to manually assign values to the parameters. Cloud Assistant automatically replaces them with the corresponding values in the environment. The following built-in environment variables are supported:
	//
	//     - `{{ACS::RegionId}}`: the region ID.
	//
	//     - `{{ACS::AccountId}}`: the UID of the Alibaba Cloud account.
	//
	//     - `{{ACS::InstanceId}}`: the instance ID. If the command is delivered to multiple instances and you want to specify `{{ACS::InstanceId}}` as a built-in environment variable, make sure that the Cloud Assistant Agent is of the following versions or later:
	//
	//         - Linux: 2.2.3.309
	//
	//         - Windows: 2.1.3.309
	//
	//     - `{{ACS::InstanceName}}`: the instance name. If the command is delivered to multiple instances and you want to specify `{{ACS::InstanceName}}` as a built-in environment variable, make sure that the Cloud Assistant Agent is of the following versions or later:
	//
	//         - Linux: 2.2.3.344
	//
	//         - Windows: 2.1.3.344
	//
	//     - `{{ACS::InvokeId}}`: the invocation ID. If you want to specify `{{ACS::InvokeId}}` as a built-in environment variable, make sure that the Cloud Assistant Agent is of the following versions or later:
	//
	//         - Linux: 2.2.3.309
	//
	//         - Windows: 2.1.3.309
	//
	//     - `{{ACS::CommandId}}`: the command ID. If you call this operation to run a command and want to specify `{{ACS::CommandId}}` as a built-in environment variable, make sure that the Cloud Assistant Agent is of the following versions or later:
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
	// The container ID. Only 64-bit hexadecimal strings are supported. The `docker://`, `containerd://`, or `cri-o://` prefix is allowed to explicitly specify the container runtime.
	//
	// Precautions:
	//
	// - If you specify this parameter, Cloud Assistant runs the script in the specified container on the instance.
	//
	// - If you specify this parameter, it is only supported on Linux instances where the Cloud Assistant Agent version is 2.2.3.344 or later.
	//
	// - If you specify this parameter, the specified `Username` and `WorkingDir` parameters do not take effect. The command can only be run by the default user of the container in the default working directory of the container. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// > Only Shell scripts can be run in Linux containers. You cannot use a command similar to `#!/usr/bin/python` at the beginning of the script to specify the interpreter for the script content. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// example:
	//
	// ab141ddfbacfe02d9dbc25966ed971536124527097398d419a6746873fea****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The container name.
	//
	// Precautions:
	//
	// - If you specify this parameter, Cloud Assistant runs the script in the specified container on the instance.
	//
	// - If you specify this parameter, it is only supported on Linux instances where the Cloud Assistant Agent version is 2.2.3.344 or later.
	//
	// - If you specify this parameter, the specified `Username` and `WorkingDir` parameters do not take effect. The command can only be run by the default user of the container in the default working directory of the container. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// > Only Shell scripts can be run in Linux containers. You cannot use a command similar to `#!/usr/bin/python` at the beginning of the script to specify the interpreter for the script content. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// example:
	//
	// test-container
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The encoding mode of the command content (`CommandContent`). Valid values (case-insensitive):
	//
	// - PlainText: no encoding. The command content is transmitted in plaintext.
	//
	// - Base64: Base64 encoding.
	//
	// Default value: PlainText. If you enter an invalid value, it is processed as PlainText.
	//
	// example:
	//
	// Base64
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The command description. It supports all character sets and can be up to 512 characters in length.
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
	// The execution time of the scheduled command. Three scheduled execution modes are supported: execution at fixed time intervals (based on rate expressions), one-time execution at a specified time, and clock-based scheduled execution (based on cron expressions).
	//
	// - Execution at fixed time intervals: runs the command at the specified time interval based on a rate expression. The time interval can be specified in seconds (s), minutes (m), hours (h), or days (d). This mode is suitable for scenarios where tasks are run at fixed time intervals. The format is `rate(<interval value><interval unit>)`. For example, to run the command every 5 minutes, the format is `rate(5m)`. The following limits apply to execution at fixed time intervals:
	//
	//     - The specified time interval must not exceed 7 days, must not be less than 60 seconds, and must be greater than the timeout period of the scheduled task.
	//
	//     - The execution interval is based only on the fixed frequency and is independent of the actual time required to run the task. For example, if you set the command to run every 5 minutes and the task takes 2 minutes to complete, the next execution starts 3 minutes after the task is completed.
	//
	//     - The task is not run immediately when it is created. For example, if you set the command to run every 5 minutes, the command is not run immediately when the task is created. Instead, it starts to run 5 minutes after the task is created.
	//
	// - One-time execution at a specified time: runs the command once at the specified time zone and execution time. The format is `at(yyyy-MM-dd HH:mm:ss <time zone>)`, which is `at(year-month-day hour:minute:second <time zone>)`. If you do not specify a time zone, the UTC time zone is used by default. The time zone supports the following three formats:
	//
	//     - Time zone full name: such as `Asia/Shanghai` (China/Shanghai time) and `America/Los_Angeles` (US/Los Angeles time).
	//
	//     - Time zone offset from Greenwich Mean Time: such as `GMT+8:00` (East 8th time zone) and `GMT-7:00` (West 7th time zone). When you use the GMT format, leading zeros are not supported for the hour digit.
	//
	//     - Time zone abbreviation: only UTC (Coordinated Universal Time) is supported.
	//
	//   For example, to run the command once at 13:15:30 on June 6, 2022 in the China/Shanghai time zone, the format is `at(2022-06-06 13:15:30 Asia/Shanghai)`. To run the command once at 13:15:30 on June 6, 2022 in the West 7th time zone, the format is `at(2022-06-06 13:15:30 GMT-7:00)`.
	//
	// - Clock-based scheduled execution (based on cron expressions): runs the command based on a cron expression. The format is `<seconds> <minutes> <hours> <day of month> <month> <day of week> <year (optional)> <time zone>`, which is `<cron expression> <time zone>`. In the specified time zone, the scheduled task execution time is calculated based on the cron expression and the task is run. If you do not specify a time zone, the internal system time zone of the instance that runs the scheduled task is used by default. For more information about cron expressions, see [Cron expressions](https://help.aliyun.com/document_detail/64769.html). The time zone supports the following three formats:
	//
	//     - Time zone full name: such as `Asia/Shanghai` (China/Shanghai time) and `America/Los_Angeles` (US/Los Angeles time).
	//
	//     - Time zone offset from Greenwich Mean Time: such as `GMT+8:00` (East 8th time zone) and `GMT-7:00` (West 7th time zone). When you use the GMT format, leading zeros are not supported for the hour digit.
	//
	//     - Time zone abbreviation: only UTC (Coordinated Universal Time) is supported.
	//
	//   For example, to run the command once at 10:15 every day in 2022 in the China/Shanghai time zone, the format is `0 15 10 ? 	- 	- 2022 Asia/Shanghai`. To run the command every half hour from 10:00 to 11:30 every day in 2022 in the East 8th time zone, the format is `0 0/30 10-11 	- 	- ? 2022 GMT+8:00`. To run the command every 5 minutes from 14:00 to 14:55 every day in October every two years starting from 2022 in the UTC time zone, the format is `0 0/5 14 	- 10 ? 2022/2 UTC`.
	//
	//     > The minimum time interval must be greater than or equal to the timeout period of the scheduled task and must not be less than 10 seconds.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The array of ECS instance IDs. Array length: 1 to 100.
	//
	// If you specify multiple instances and one of the instances does not meet the execution conditions, you must reselect the instances.
	//
	// You can also apply for a quota increase in Quota Center. The quota name is the maximum number of instances that support command execution.
	//
	// example:
	//
	// i-bp185dy2o3o6neg****
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// Specifies whether to retain the command after it is run. Valid values:
	//
	// - true: retains the command. You can call the InvokeCommand operation to run the command again. This consumes the quota for retained Cloud Assistant commands.
	//
	// - false: does not retain the command. The command is automatically deleted after it is run and does not consume the quota for retained Cloud Assistant commands.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	KeepCommand *bool `json:"KeepCommand,omitempty" xml:"KeepCommand,omitempty"`
	// The bootstrap program for script execution. It cannot exceed 1 KB in length.
	//
	// example:
	//
	// python3 -u {{ACS::ScriptFileName|Ext(".py")}}
	Launcher *string `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	// The command name. It supports all character sets and can be up to 128 characters in length.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The OSS delivery configuration for the command execution output.
	//
	// - Format: oss://${BucketName}/${Prefix}. ${BucketName} is the name of the OSS bucket to which the output is delivered, and ${Prefix} is the directory prefix to which the output is delivered.
	//
	// example:
	//
	// oss://testBucket/testPrefix
	OssOutputDelivery *string `json:"OssOutputDelivery,omitempty" xml:"OssOutputDelivery,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key-value pairs of custom parameters that are passed in when the command is run, if the command contains custom parameters. For example, if the command content is `echo {{name}}`, you can pass in the key-value pair `{"name":"Jack"}` by using the `Parameter` parameter. The custom parameter automatically replaces the variable value `name`, and a new command is generated. The actual command that is run is `echo Jack`.
	//
	// The number of custom parameters ranges from 0 to 10. Take note of the following items:
	//
	// - The key cannot be an empty string and can be up to 64 characters in length.
	//
	// - The value can be an empty string.
	//
	// - After the custom parameters and the original command content are Base64-encoded, the size of the Base64-encoded command content cannot exceed 24 KB. You can use the `KeepCommand` parameter to specify whether to retain the command.
	//
	// - The set of custom parameter names must be a subset of the parameter set defined when the command is created. For parameters that are not passed in, you can use empty strings instead.
	//
	// The default value is empty, which indicates that this parameter is not set and the custom parameter feature is disabled.
	//
	// example:
	//
	// {"name":"Jack", "accessKey":"LTAI*************"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The execution mode of the command. Valid values:
	//
	// - Once: runs the command immediately.
	//
	// - Period: runs the command on a schedule. If you set this parameter to `Period`, you must also specify the `Frequency` parameter.
	//
	// - NextRebootOnly: automatically runs the command the next time the instance starts.
	//
	// - EveryReboot: automatically runs the command every time the instance starts.
	//
	// - DryRun: prechecks the request only. The command is not actually run. The check items include the request parameters, instance execution environment, and Cloud Assistant Agent running status.
	//
	// Default values:
	//
	// - If you do not specify the `Frequency` parameter, the default value is `Once`.
	//
	// - If you specify the `Frequency` parameter, the value is processed as `Period` regardless of whether this parameter is set.
	//
	// Precautions:
	//
	// - You can call the [StopInvocation](https://help.aliyun.com/document_detail/64838.html) operation to stop a pending command or a scheduled command.
	//
	// - If you set this parameter to `Period` or `EveryReboot`, you can call the [DescribeInvocationResults](https://help.aliyun.com/document_detail/64845.html) operation and set `IncludeHistory` to true to view the historical records of the scheduled command execution.
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// The ID of the resource group to which the command execution belongs. When you specify this parameter:
	//
	// - If the ECS instance specified by InstanceId belongs to a non-default resource group, the ECS instance must belong to this resource group.
	//
	// - You can specify this parameter to filter the corresponding command execution results by calling the [DescribeInvocations](https://help.aliyun.com/document_detail/64840.html) or [DescribeInvocationResults](https://help.aliyun.com/document_detail/64845.html) operation.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The array of tags used to filter instances. Array length: 0 to 20. You can run the command on multiple instances that have the same tags without specifying InstanceId.
	ResourceTag []*RunCommandShrinkRequestResourceTag `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty" type:"Repeated"`
	// The array of tags. Array length: 0 to 20.
	Tag []*RunCommandShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The mode used when the task is stopped (manually stopped or interrupted due to execution timeout). Valid values:
	//
	// - Process: stops the current script process.
	//
	// - ProcessTree: stops the current process tree (the set of the script process and all child processes created by it).
	//
	// example:
	//
	// ProcessTree
	TerminationMode *string `json:"TerminationMode,omitempty" xml:"TerminationMode,omitempty"`
	// > **[Deprecated]*	- This parameter is deprecated and does not take effect.
	//
	// example:
	//
	// true
	Timed *bool `json:"Timed,omitempty" xml:"Timed,omitempty"`
	// The timeout period for running the command. Unit: seconds.
	//
	// If the command cannot be run because of process issues, missing modules, or a missing Cloud Assistant Agent, a timeout occurs. After a timeout, the command process is forcefully terminated.
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
	// The username that is used to run the command on the ECS instance. It can be up to 255 characters in length.
	//
	// - For Linux ECS instances, the command is run by the root user by default.
	//
	// - For Windows ECS instances, the command is run by the System user by default.
	//
	// You can also specify another existing user on the instance to run the command. Running Cloud Assistant commands as a regular user is more secure. For more information, see [Configure a regular user to run Cloud Assistant commands](https://help.aliyun.com/document_detail/203771.html).
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The password name of the user that is used to execute the command on the Windows instance. It can be up to 255 characters in length.
	//
	// If you want to execute the command as a non-default user (System) on a Windows instance, you must specify both `Username` and this parameter. To reduce the risk of password leakage, you must host the plaintext password in the parameter repository of CloudOps Orchestration Service and pass in only the password name here. For more information, see [Encrypt parameters](https://help.aliyun.com/document_detail/186828.html) and [Settings for regular users to execute Cloud Assistant commands](https://help.aliyun.com/document_detail/203771.html).
	//
	// > You do not need to pass in this parameter when you execute the command as the root user on a Linux instance or the System user on a Windows instance.
	//
	// example:
	//
	// axtSecretPassword
	WindowsPasswordName *string `json:"WindowsPasswordName,omitempty" xml:"WindowsPasswordName,omitempty"`
	// The running directory of the command on the ECS instance. It can be up to 200 characters in length.
	//
	// Default values:
	//
	// - For Linux instances, the default running directory is the home directory of the root user, which is `/root`.
	//
	// - For Windows instances, the default running directory is the directory where the Cloud Assistant Agent process resides, such as `C:\\Windows\\System32`.
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
	// - This parameter conflicts with the InstanceId parameter and they cannot be specified at the same time.
	//
	// - It cannot be an empty string once specified.
	//
	// - The number of instances under the tag cannot exceed the limit on the number of InstanceId.N parameters. If the number of instances exceeds the limit, we recommend that you control the number of instances by adding batch tags, such as batch: b1.
	//
	// - It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value used to filter instances.
	//
	// Precautions:
	//
	// - It can be an empty string.
	//
	// - It can be up to 128 characters in length and cannot contain http:// or https://.
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
	// The tag key of the command to execute. It cannot be an empty string once specified.
	//
	// If you use a single tag to filter resources, the resource count under the tag cannot exceed 1,000. If you use multiple tags to filter resources, the resource count of resources that are attached to all the specified tags cannot exceed 1,000. If the resource count exceeds 1,000, you must use the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to query the resources.
	//
	// It can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the command execution. It can be an empty string.
	//
	// It can be up to 128 characters in length and cannot contain `http://` or `https://`.
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
