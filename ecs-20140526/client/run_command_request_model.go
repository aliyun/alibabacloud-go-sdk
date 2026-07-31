// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RunCommandRequest
	GetClientToken() *string
	SetCommandContent(v string) *RunCommandRequest
	GetCommandContent() *string
	SetContainerId(v string) *RunCommandRequest
	GetContainerId() *string
	SetContainerName(v string) *RunCommandRequest
	GetContainerName() *string
	SetContentEncoding(v string) *RunCommandRequest
	GetContentEncoding() *string
	SetDescription(v string) *RunCommandRequest
	GetDescription() *string
	SetEnableParameter(v bool) *RunCommandRequest
	GetEnableParameter() *bool
	SetFrequency(v string) *RunCommandRequest
	GetFrequency() *string
	SetInstanceId(v []*string) *RunCommandRequest
	GetInstanceId() []*string
	SetKeepCommand(v bool) *RunCommandRequest
	GetKeepCommand() *bool
	SetLauncher(v string) *RunCommandRequest
	GetLauncher() *string
	SetName(v string) *RunCommandRequest
	GetName() *string
	SetOssOutputDelivery(v string) *RunCommandRequest
	GetOssOutputDelivery() *string
	SetOwnerAccount(v string) *RunCommandRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RunCommandRequest
	GetOwnerId() *int64
	SetParameters(v map[string]interface{}) *RunCommandRequest
	GetParameters() map[string]interface{}
	SetRegionId(v string) *RunCommandRequest
	GetRegionId() *string
	SetRepeatMode(v string) *RunCommandRequest
	GetRepeatMode() *string
	SetResourceGroupId(v string) *RunCommandRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *RunCommandRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RunCommandRequest
	GetResourceOwnerId() *int64
	SetResourceTag(v []*RunCommandRequestResourceTag) *RunCommandRequest
	GetResourceTag() []*RunCommandRequestResourceTag
	SetTag(v []*RunCommandRequestTag) *RunCommandRequest
	GetTag() []*RunCommandRequestTag
	SetTerminationMode(v string) *RunCommandRequest
	GetTerminationMode() *string
	SetTimed(v bool) *RunCommandRequest
	GetTimed() *bool
	SetTimeout(v int64) *RunCommandRequest
	GetTimeout() *int64
	SetType(v string) *RunCommandRequest
	GetType() *string
	SetUsername(v string) *RunCommandRequest
	GetUsername() *string
	SetWindowsPasswordName(v string) *RunCommandRequest
	GetWindowsPasswordName() *string
	SetWorkingDir(v string) *RunCommandRequest
	GetWorkingDir() *string
}

type RunCommandRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The command content. The command content can be plaintext or Base64-encoded. Note the following items:
	//
	// - If you save the command, the Base64-encoded command content cannot exceed 18 KB. If you do not save the command, the Base64-encoded command content cannot exceed 24 KB. You can use `KeepCommand` to specify whether to save the command.
	//
	// - If the command content is Base64-encoded, set `ContentEncoding=Base64`.
	//
	// - Set `EnableParameter=true` to enable the custom parameter feature in the command content:
	//
	//     - Define custom parameters by enclosing them in `{{}}`. Spaces and line breaks before and after the parameter name within `{{}}` are ignored.
	//
	//     - The number of custom parameters cannot exceed 20.
	//
	//     - Custom parameter names can contain a-z, A-Z, 0-9, hyphens (-), and underscores (_). The acs:: prefix for specifying non-built-in environment parameters is not supported. Other characters are not supported. Parameter names are case-insensitive.
	//
	//     - Each custom parameter name cannot exceed 64 bytes.
	//
	// - You can specify built-in environment parameters as custom parameters. When the command is executed, Cloud Assistant automatically replaces them with the corresponding values without manual assignment. The following built-in environment parameters are supported:
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
	//     - `{{ACS::InvokeId}}`: The command execution ID. To use `{{ACS::InvokeId}}` as a built-in environment parameter, ensure that the Cloud Assistant Agent version is no earlier than:
	//
	//         - Linux: 2.2.3.309
	//
	//         - Windows: 2.1.3.309
	//
	//     - `{{ACS::CommandId}}`: The command ID. When you call this operation to run a command and want to use `{{ACS::CommandId}}` as a built-in environment parameter, ensure that the Cloud Assistant Agent version is no earlier than:
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
	// The container ID. Only 64-bit hexadecimal strings are supported. You can use the `docker://`, `containerd://`, or `cri-o://` prefix to explicitly specify the container runtime.
	//
	// Notes:
	//
	// - If this parameter is specified, Cloud Assistant executes the script in the specified container on the instance.
	//
	// - If this parameter is specified, only Linux instances with Cloud Assistant Agent version 2.2.3.344 or later are supported.
	//
	// - If this parameter is specified, the `Username` and `WorkingDir` parameters do not take effect. Commands can only be executed by the default container user in the default working directory of the container. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// > Only Shell scripts are supported in Linux containers. Specifying an interpreter at the beginning of the script (such as `#!/usr/bin/python`) is not supported. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// example:
	//
	// ab141ddfbacfe02d9dbc25966ed971536124527097398d419a6746873fea****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The container name.
	//
	// Notes:
	//
	// - If this parameter is specified, Cloud Assistant executes the script in the specified container on the instance.
	//
	// - If this parameter is specified, only Linux instances with Cloud Assistant Agent version 2.2.3.344 or later are supported.
	//
	// - If this parameter is specified, the `Username` and `WorkingDir` parameters do not take effect. Commands can only be executed by the default container user in the default working directory of the container. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// > Only Shell scripts are supported in Linux containers. Specifying an interpreter at the beginning of the script (such as `#!/usr/bin/python`) is not supported. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
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
	// Default value: PlainText. If an invalid value is specified, it is treated as PlainText.
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
	// The schedule for running the command. Three scheduling methods are supported: execution at fixed intervals (based on a Rate expression), one-time execution at a specified time, and clock-based scheduled execution (based on a Cron expression).
	//
	// - Execution at fixed intervals: Based on a Rate expression, the command is executed at the specified interval. The interval can be specified in seconds (s), minutes (m), hours (h), or days (d). This method is suitable for scenarios where tasks are executed at fixed intervals. Format: `rate(<interval value><interval unit>)`. For example, to execute every 5 minutes: `rate(5m)`. Limitations for fixed-interval execution:
	//
	//     - The interval cannot exceed 7 days or be less than 60 seconds, and must be greater than the timeout period of the scheduled task.
	//
	//     - The interval is based on a fixed frequency and is unrelated to the actual execution time of the task. For example, if the command is set to execute every 5 minutes and the task takes 2 minutes to complete, the next execution starts 3 minutes after the task completes.
	//
	//     - The task is not executed immediately upon creation. For example, if the command is set to execute every 5 minutes, it does not execute immediately when the task is created. Instead, execution begins 5 minutes after the task is created.
	//
	// - One-time execution at a specified time: The command is executed once at the specified time zone and time point. Format: `at(yyyy-MM-dd HH:mm:ss <time zone>)`. If no time zone is specified, the default is UTC. The time zone supports the following three formats:
	//
	//     - Full time zone name: such as `Asia/Shanghai` (China/Shanghai time) or `America/Los_Angeles` (US/Los Angeles time).
	//
	//     - GMT offset from Greenwich Mean Time: such as `GMT+8:00` (East 8th time zone) or `GMT-7:00` (West 7th time zone). When using GMT format, leading zeros are not supported in the hour field.
	//
	//     - Time zone abbreviation: Only UTC (Coordinated Universal Time) is supported.
	//
	//   For example, to execute once at 13:15:30 on June 6, 2022 in China/Shanghai time: `at(2022-06-06 13:15:30 Asia/Shanghai)`. To execute once at 13:15:30 on June 6, 2022 in GMT-7:00: `at(2022-06-06 13:15:30 GMT-7:00)`.
	//
	// - Clock-based scheduled execution (based on a Cron expression): Based on a Cron expression, the command is executed according to the scheduled task settings. Format: `<seconds> <minutes> <hours> <day of month> <month> <day of week> <year (optional)> <time zone>`, i.e., `<Cron expression> <time zone>`. The scheduled execution time is calculated based on the Cron expression in the specified time zone. If no time zone is specified, the system time zone of the instance running the scheduled task is used. For more information about Cron expressions, see [Cron expressions](https://help.aliyun.com/document_detail/64769.html). The time zone supports the following three formats:
	//
	//     - Full time zone name: such as `Asia/Shanghai` (China/Shanghai time) or `America/Los_Angeles` (US/Los Angeles time).
	//
	//     - GMT offset from Greenwich Mean Time: such as `GMT+8:00` (East 8th time zone) or `GMT-7:00` (West 7th time zone). When using GMT format, leading zeros are not supported in the hour field.
	//
	//     - Time zone abbreviation: Only UTC (Coordinated Universal Time) is supported.
	//
	//   For example, to execute once daily at 10:15 AM in China/Shanghai time in 2022: `0 15 10 ? 	- 	- 2022 Asia/Shanghai`. To execute every 30 minutes from 10:00 AM to 11:30 AM daily in GMT+8:00 in 2022: `0 0/30 10-11 	- 	- ? 2022 GMT+8:00`. To execute every 5 minutes from 2:00 PM to 2:55 PM every day in October every two years starting from 2022 in UTC: `0 0/5 14 	- 10 ? 2022/2 UTC`.
	//
	//     > The minimum interval must be greater than or equal to the timeout period of the scheduled task and no less than 10 seconds.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The instance ID list of ECS instances. Array length: 1 to 100.
	//
	// If one of the specified instances does not meet the execution conditions, you must reselect the instances.
	//
	// You can also request a quota increase in Quota Center (quota name: Maximum number of instances supported for command execute).
	//
	// example:
	//
	// i-bp185dy2o3o6neg****
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// Specifies whether to retain the command after execution. Valid values:
	//
	// - true: retains the command. You can run it again by calling InvokeCommand. This counts toward the Cloud Assistant command retention quota.
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
	// The key-value pairs of custom parameters to pass in when the command contains custom parameters. For example, if the command content is `echo {{name}}`, you can pass in the key-value pair `{"name":"Jack"}` through the `Parameter` parameter. The custom parameter automatically replaces the variable value `name`, resulting in a new command that actually executes `echo Jack`.
	//
	// The number of custom parameters ranges from 0 to 10. Note the following items:
	//
	// - Keys cannot be empty strings and can contain up to 64 characters.
	//
	// - Values can be empty strings.
	//
	// - If you save the command, the combined Base64-encoded size of custom parameters and original command content cannot exceed 18 KB. If you do not save the command, the size cannot exceed 24 KB. You can use `KeepCommand` to specify whether to save the command.
	//
	// - The set of custom parameter names must be a subset of the parameter set defined when the command was created. For parameters that are not passed in, you can use empty strings as substitutes.
	//
	// Default value: empty, which disables custom parameters.
	//
	// example:
	//
	// {"name":"Jack", "accessKey":"LTAI*************"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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
	// - Period: runs the command on a schedule. If you set this parameter to `Period`, you must also specify the `Frequency` parameter.
	//
	// - NextRebootOnly: automatically runs the command the next time the instance starts.
	//
	// - EveryReboot: automatically runs the command every time the instance starts.
	//
	// - DryRun: performs a dry run of the request only. The command is not actually executed. The check items include request parameters, instance execution environment, and Cloud Assistant Agent running status.
	//
	// Default value:
	//
	// - If the `Frequency` parameter is not specified, the default value is `Once`.
	//
	// - If the `Frequency` parameter is specified, the command is processed as `Period` regardless of whether this parameter is set.
	//
	// Notes:
	//
	// - You can call [StopInvocation](https://help.aliyun.com/document_detail/64838.html) to stop a pending or scheduled command.
	//
	// - If this parameter is set to `Period` or `EveryReboot`, you can call [DescribeInvocationResults](https://help.aliyun.com/document_detail/64845.html) and specify `IncludeHistory=true` to view the historical records of scheduled command executions.
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// The resource group ID for the command execution. If you specify this parameter:
	//
	// - If the ECS instance specified by InstanceId belongs to a non-default resource group, the ECS instance must belong to this resource group.
	//
	// - You can filter command execution results by specifying this parameter (by calling [DescribeInvocations](https://help.aliyun.com/document_detail/64840.html) or [DescribeInvocationResults](https://help.aliyun.com/document_detail/64845.html)).
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags used to filter instances. Array length: 0 to 20. You can run commands in batches on instances with the same tags without specifying InstanceId.
	ResourceTag []*RunCommandRequestResourceTag `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty" type:"Repeated"`
	// The tag pairs. Array length: 0 to 20.
	Tag []*RunCommandRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// > This parameter is deprecated and has no effect if specified.
	//
	// example:
	//
	// true
	Timed *bool `json:"Timed,omitempty" xml:"Timed,omitempty"`
	// The timeout period for command execution. Unit: seconds.
	//
	// A timeout occurs when a command cannot be run because the process does not exist, a module is missing, or Cloud Assistant Agent is unavailable. When a timeout occurs, the command process is forcefully terminated.
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
	// The username for executing the command on the ECS instance. The value cannot exceed 255 characters in length.
	//
	// - For Linux ECS instances, commands are executed as the root user by default.
	//
	// - For Windows ECS instances, commands are executed as the System user by default.
	//
	// You can also specify another existing user on the instance to execute the command. Executing Cloud Assistant commands as a regular user is more secure. For more information, see [Configure a regular user to run Cloud Assistant commands](https://help.aliyun.com/document_detail/203771.html).
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The name of the password for the user who executes the command on a Windows instance. The value cannot exceed 255 characters in length.
	//
	// If you want to execute a command as a non-default user (System) on a Windows instance, you must specify both `Username` and this parameter. To reduce the risk of password leakage, store the plaintext password in the parameter repository of CloudOps Orchestration Service and pass only the password name here. For more information, see [Encryption parameters](https://help.aliyun.com/document_detail/186828.html) and [Configure a regular user to run Cloud Assistant commands](https://help.aliyun.com/document_detail/203771.html).
	//
	// > This parameter is not required when you use the root user on a Linux instance or the System user on a Windows instance to execute commands.
	//
	// example:
	//
	// axtSecretPassword
	WindowsPasswordName *string `json:"WindowsPasswordName,omitempty" xml:"WindowsPasswordName,omitempty"`
	// The working directory of the command on the ECS instance. The value cannot exceed 200 characters in length.
	//
	// Default value:
	//
	// - For Linux instances, the default directory is the home directory of the administrator (root user): `/root`.
	//
	// - For Windows instances, the default directory is the directory where the Cloud Assistant Agent process is located, such as `C:\\Windows\\System32`.
	//
	// example:
	//
	// /home/user
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s RunCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s RunCommandRequest) GoString() string {
	return s.String()
}

func (s *RunCommandRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunCommandRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *RunCommandRequest) GetContainerId() *string {
	return s.ContainerId
}

func (s *RunCommandRequest) GetContainerName() *string {
	return s.ContainerName
}

func (s *RunCommandRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *RunCommandRequest) GetDescription() *string {
	return s.Description
}

func (s *RunCommandRequest) GetEnableParameter() *bool {
	return s.EnableParameter
}

func (s *RunCommandRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *RunCommandRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *RunCommandRequest) GetKeepCommand() *bool {
	return s.KeepCommand
}

func (s *RunCommandRequest) GetLauncher() *string {
	return s.Launcher
}

func (s *RunCommandRequest) GetName() *string {
	return s.Name
}

func (s *RunCommandRequest) GetOssOutputDelivery() *string {
	return s.OssOutputDelivery
}

func (s *RunCommandRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RunCommandRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RunCommandRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *RunCommandRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunCommandRequest) GetRepeatMode() *string {
	return s.RepeatMode
}

func (s *RunCommandRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RunCommandRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RunCommandRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RunCommandRequest) GetResourceTag() []*RunCommandRequestResourceTag {
	return s.ResourceTag
}

func (s *RunCommandRequest) GetTag() []*RunCommandRequestTag {
	return s.Tag
}

func (s *RunCommandRequest) GetTerminationMode() *string {
	return s.TerminationMode
}

func (s *RunCommandRequest) GetTimed() *bool {
	return s.Timed
}

func (s *RunCommandRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *RunCommandRequest) GetType() *string {
	return s.Type
}

func (s *RunCommandRequest) GetUsername() *string {
	return s.Username
}

func (s *RunCommandRequest) GetWindowsPasswordName() *string {
	return s.WindowsPasswordName
}

func (s *RunCommandRequest) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *RunCommandRequest) SetClientToken(v string) *RunCommandRequest {
	s.ClientToken = &v
	return s
}

func (s *RunCommandRequest) SetCommandContent(v string) *RunCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandRequest) SetContainerId(v string) *RunCommandRequest {
	s.ContainerId = &v
	return s
}

func (s *RunCommandRequest) SetContainerName(v string) *RunCommandRequest {
	s.ContainerName = &v
	return s
}

func (s *RunCommandRequest) SetContentEncoding(v string) *RunCommandRequest {
	s.ContentEncoding = &v
	return s
}

func (s *RunCommandRequest) SetDescription(v string) *RunCommandRequest {
	s.Description = &v
	return s
}

func (s *RunCommandRequest) SetEnableParameter(v bool) *RunCommandRequest {
	s.EnableParameter = &v
	return s
}

func (s *RunCommandRequest) SetFrequency(v string) *RunCommandRequest {
	s.Frequency = &v
	return s
}

func (s *RunCommandRequest) SetInstanceId(v []*string) *RunCommandRequest {
	s.InstanceId = v
	return s
}

func (s *RunCommandRequest) SetKeepCommand(v bool) *RunCommandRequest {
	s.KeepCommand = &v
	return s
}

func (s *RunCommandRequest) SetLauncher(v string) *RunCommandRequest {
	s.Launcher = &v
	return s
}

func (s *RunCommandRequest) SetName(v string) *RunCommandRequest {
	s.Name = &v
	return s
}

func (s *RunCommandRequest) SetOssOutputDelivery(v string) *RunCommandRequest {
	s.OssOutputDelivery = &v
	return s
}

func (s *RunCommandRequest) SetOwnerAccount(v string) *RunCommandRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RunCommandRequest) SetOwnerId(v int64) *RunCommandRequest {
	s.OwnerId = &v
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

func (s *RunCommandRequest) SetRepeatMode(v string) *RunCommandRequest {
	s.RepeatMode = &v
	return s
}

func (s *RunCommandRequest) SetResourceGroupId(v string) *RunCommandRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RunCommandRequest) SetResourceOwnerAccount(v string) *RunCommandRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RunCommandRequest) SetResourceOwnerId(v int64) *RunCommandRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RunCommandRequest) SetResourceTag(v []*RunCommandRequestResourceTag) *RunCommandRequest {
	s.ResourceTag = v
	return s
}

func (s *RunCommandRequest) SetTag(v []*RunCommandRequestTag) *RunCommandRequest {
	s.Tag = v
	return s
}

func (s *RunCommandRequest) SetTerminationMode(v string) *RunCommandRequest {
	s.TerminationMode = &v
	return s
}

func (s *RunCommandRequest) SetTimed(v bool) *RunCommandRequest {
	s.Timed = &v
	return s
}

func (s *RunCommandRequest) SetTimeout(v int64) *RunCommandRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandRequest) SetType(v string) *RunCommandRequest {
	s.Type = &v
	return s
}

func (s *RunCommandRequest) SetUsername(v string) *RunCommandRequest {
	s.Username = &v
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

func (s *RunCommandRequest) Validate() error {
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

type RunCommandRequestResourceTag struct {
	// The tag key used to filter instances.
	//
	// Notes:
	//
	// - This parameter conflicts with the InstanceId parameter. You cannot specify both.
	//
	// - If you specify this parameter, it cannot be an empty string.
	//
	// - The number of instances with the specified tag cannot exceed the limit of InstanceId.N. If the number of instances exceeds the limit, control the number of instances by adding batch tags, such as batch: b1.
	//
	// - The key can be up to 64 characters in length and cannot start with aliyun or acs:, or contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value used to filter instances.
	//
	// Notes:
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

func (s RunCommandRequestResourceTag) String() string {
	return dara.Prettify(s)
}

func (s RunCommandRequestResourceTag) GoString() string {
	return s.String()
}

func (s *RunCommandRequestResourceTag) GetKey() *string {
	return s.Key
}

func (s *RunCommandRequestResourceTag) GetValue() *string {
	return s.Value
}

func (s *RunCommandRequestResourceTag) SetKey(v string) *RunCommandRequestResourceTag {
	s.Key = &v
	return s
}

func (s *RunCommandRequestResourceTag) SetValue(v string) *RunCommandRequestResourceTag {
	s.Value = &v
	return s
}

func (s *RunCommandRequestResourceTag) Validate() error {
	return dara.Validate(s)
}

type RunCommandRequestTag struct {
	// The tag key of the command execution. If you specify this parameter, it cannot be an empty string.
	//
	// If you use a single tag to filter resources, the number of resources with this tag cannot exceed 1,000. If you use multiple tags to filter resources, the number of resources with all specified tags attached cannot exceed 1,000. If the number of resources exceeds 1,000, use the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to execute the query.
	//
	// The key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`, or contain `http://` or `https://`.
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

func (s RunCommandRequestTag) String() string {
	return dara.Prettify(s)
}

func (s RunCommandRequestTag) GoString() string {
	return s.String()
}

func (s *RunCommandRequestTag) GetKey() *string {
	return s.Key
}

func (s *RunCommandRequestTag) GetValue() *string {
	return s.Value
}

func (s *RunCommandRequestTag) SetKey(v string) *RunCommandRequestTag {
	s.Key = &v
	return s
}

func (s *RunCommandRequestTag) SetValue(v string) *RunCommandRequestTag {
	s.Value = &v
	return s
}

func (s *RunCommandRequestTag) Validate() error {
	return dara.Validate(s)
}
