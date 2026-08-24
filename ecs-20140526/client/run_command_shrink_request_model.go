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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. **ClientToken*	- supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The command content. The command content can be plaintext or Base64-encoded. Note the following:
	//
	// - The size of the command content after Base64 encoding cannot exceed 24 KB. You can use `KeepCommand` to specify whether to retain the command.
	//
	// - If the command content is Base64-encoded, set `ContentEncoding=Base64`.
	//
	// - Set `EnableParameter=true` to enable custom parameters in the command content:
	//
	//     - Custom parameters are defined by enclosing them in `{{}}`. Spaces and line breaks before and after the parameter name within `{{}}` are ignored.
	//
	//     - A maximum of 20 custom parameters are supported.
	//
	//     - Custom parameter names can contain only a-z, A-Z, 0-9, hyphens (-), and underscores (_). The acs:: prefix for specifying non-built-in environment parameters is not supported. Other characters are not supported. Parameter names are case-insensitive.
	//
	//     - Each custom parameter name cannot exceed 64 bytes.
	//
	// - You can specify built-in environment parameters as custom parameters. When the command is run, Cloud Assistant automatically replaces the parameters with the corresponding values in the environment without manual assignment. The following built-in environment parameters are supported:
	//
	//     - `{{ACS::RegionId}}`: The region ID.
	//
	//     - `{{ACS::AccountId}}`: The UID of the Alibaba Cloud account.
	//
	//     - `{{ACS::InstanceId}}`: The instance ID. When the command is sent to multiple instances and you want to use `{{ACS::InstanceId}}` as a built-in environment parameter, make sure that the Cloud Assistant Agent version is not earlier than the following:
	//
	//         - Linux: 2.2.3.309
	//
	//         - Windows: 2.1.3.309
	//
	//     - `{{ACS::InstanceName}}`: The instance name. When the command is sent to multiple instances and you want to use `{{ACS::InstanceName}}` as a built-in environment parameter, make sure that the Cloud Assistant Agent version is not earlier than the following:
	//
	//         - Linux: 2.2.3.344
	//
	//         - Windows: 2.1.3.344
	//
	//     - `{{ACS::InvokeId}}`: The command execution ID. To use `{{ACS::InvokeId}}` as a built-in environment parameter, make sure that the Cloud Assistant Agent version is not earlier than the following:
	//
	//         - Linux: 2.2.3.309
	//
	//         - Windows: 2.1.3.309
	//
	//     - `{{ACS::CommandId}}`: The command ID. When you call this operation to run a command and want to use `{{ACS::CommandId}}` as a built-in environment parameter, make sure that the Cloud Assistant Agent version is not earlier than the following:
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
	// The container ID. Only 64-bit hexadecimal strings are supported. The `docker://`, `containerd://`, or `cri-o://` prefix can be added to specify the container runtime.
	//
	// Precautions:
	//
	// - If this parameter is specified, Cloud Assistant runs the script in the specified container of the instance.
	//
	// - If this parameter is specified, the script can run only on Linux instances whose Cloud Assistant Agent version is 2.2.3.344 or later.
	//
	// - If this parameter is specified, the specified `Username` and `WorkingDir` parameters do not take effect. Commands can be run only by using the default user of the container in the default working directory of the container. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// > In Linux containers, only shell scripts are supported. You cannot use a command such as `#!/usr/bin/python` at the beginning of a script to specify the interpreter. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// example:
	//
	// ab141ddfbacfe02d9dbc25966ed971536124527097398d419a6746873fea****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The container name.
	//
	// Precautions:
	//
	// - If this parameter is specified, Cloud Assistant runs the script in the specified container of the instance.
	//
	// - If this parameter is specified, the script can run only on Linux instances whose Cloud Assistant Agent version is 2.2.3.344 or later.
	//
	// - If this parameter is specified, the Username and WorkingDir parameters do not take effect. Commands can be run only by using the default user in the default working directory of the container. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// > In Linux containers, only shell scripts can be run. You cannot specify an interpreter for the script content by adding a command such as `#!/usr/bin/python` to the beginning of the script. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// example:
	//
	// test-container
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The encoding mode of the command content (`CommandContent`). Valid values (case-insensitive):
	//
	// example:
	//
	// Base64
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The command description. The description supports all character sets and can be up to 512 characters in length.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the command contains custom parameters.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The execution time for scheduled command execution. Three scheduling methods are supported: fixed interval execution (based on Rate expressions), one-time execution at a specified time, and clock-based scheduled execution (based on Cron expressions).
	//
	// - Fixed interval execution: Based on Rate expressions, commands are executed at the specified time interval. The time interval can be specified in seconds (s), minutes (m), hours (h), or days (d). This method is suitable for scenarios that require task execution at fixed intervals. The format is `rate(<interval value><interval unit>)`. For example, to execute a command every 5 minutes, use `rate(5m)`. The following limits apply to fixed interval execution:
	//
	//     - The specified interval cannot exceed 7 days or be less than 60 seconds, and must be greater than the timeout period of the scheduled task.
	//
	//     - The execution interval is based on a fixed frequency and is independent of the actual execution time of the task. For example, if a command is set to execute every 5 minutes and the task takes 2 minutes to complete, the next execution starts 3 minutes after the task is completed.
	//
	//     - The task is not executed immediately upon creation. For example, if a command is set to execute every 5 minutes, the command is not executed immediately when the task is created. Instead, execution starts 5 minutes after the task is created.
	//
	// - One-time execution at a specified time: The command is executed once at the specified time and time zone. The format is `at(yyyy-MM-dd HH:mm:ss <time zone>)`. If no time zone is specified, the default is UTC. The following three time zone formats are supported:
	//
	//     - Full time zone name: For example, `Asia/Shanghai` (China/Shanghai time) or `America/Los_Angeles` (US/Los Angeles time).
	//
	//     - GMT offset from Greenwich Mean Time: For example, `GMT+8:00` (UTC+8) or `GMT-7:00` (UTC-7). When using the GMT format, leading zeros are not supported for the hour value.
	//
	//     - Time zone abbreviation: Only UTC (Coordinated Universal Time) is supported.
	//
	//   For example, to execute a command once at 13:15:30 on June 6, 2022 in China/Shanghai time, use `at(2022-06-06 13:15:30 Asia/Shanghai)`. To execute a command once at 13:15:30 on June 6, 2022 in UTC-7, use `at(2022-06-06 13:15:30 GMT-7:00)`.
	//
	// - Clock-based scheduled execution (based on Cron expressions): Based on Cron expressions, commands are executed according to the scheduled task settings. The format is `<seconds> <minutes> <hours> <day of month> <month> <day of week> <year (optional)> <time zone>`, which is `<Cron expression> <time zone>`. The scheduled task execution time is calculated based on the Cron expression in the specified time zone. If no time zone is specified, the default is the internal system time zone of the instance that runs the scheduled task. For more information about Cron expressions, see [Cron expressions](https://help.aliyun.com/document_detail/64769.html). The following three time zone formats are supported:
	//
	//     - Full time zone name: For example, `Asia/Shanghai` (China/Shanghai time) or `America/Los_Angeles` (US/Los Angeles time).
	//
	//     - GMT offset from Greenwich Mean Time: For example, `GMT+8:00` (UTC+8) or `GMT-7:00` (UTC-7). When using the GMT format, leading zeros are not supported for the hour value.
	//
	//     - Time zone abbreviation: Only UTC (Coordinated Universal Time) is supported.
	//
	//   For example, to execute a command once at 10:15 every day in 2022 in China/Shanghai time, use `0 15 10 ? 	- 	- 2022 Asia/Shanghai`. To execute a command every 30 minutes from 10:00 to 11:30 every day in 2022 in UTC+8, use `0 0/30 10-11 	- 	- ? 2022 GMT+8:00`. To execute a command every 5 minutes from 14:00 to 14:55 every day in October every two years starting from 2022 in UTC, use `0 0/5 14 	- 10 ? 2022/2 UTC`.
	//
	//     >The minimum time interval must be greater than or equal to the timeout period of the scheduled task and no less than 10 seconds.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The ECS instance ID array. Array length: 1 to 100.
	//
	// example:
	//
	// i-bp185dy2o3o6neg****
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// Specifies whether to retain the command after it is run. Valid values:
	//
	// - true: The command is retained. You can run it again by calling InvokeCommand. The command counts against the Cloud Assistant command quota.
	//
	// - false: The command is not retained. It is automatically deleted after execution and does not count against the Cloud Assistant command quota.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	KeepCommand *bool `json:"KeepCommand,omitempty" xml:"KeepCommand,omitempty"`
	// The bootstrap program for script execution. The value can be up to 1 KB in length.
	//
	// example:
	//
	// python3 -u {{ACS::ScriptFileName|Ext(".py")}}
	Launcher *string `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	// The command name. The name supports all character sets and can be up to 128 characters in length.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The OSS delivery configuration for command execution output.
	//
	// - Format: oss://${BucketName}/${Prefix}, where ${BucketName} is the name of the destination OSS bucket and ${Prefix} is the directory prefix of the destination.
	//
	// example:
	//
	// oss://testBucket/testPrefix
	OssOutputDelivery *string `json:"OssOutputDelivery,omitempty" xml:"OssOutputDelivery,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key-value pairs of custom parameters to pass in when the command contains custom parameters. For example, if the command content is `echo {{name}}`, you can use the `Parameter` parameter to pass in the key-value pair `{"name":"Jack"}`. The custom parameter automatically replaces the variable value `name`, and the command that is actually run is `echo Jack`.
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
	// - Once: immediately executes the command.
	//
	// - Period: executes the command on a schedule. If you set this parameter to `Period`, you must also specify the `Frequency` parameter.
	//
	// - NextRebootOnly: automatically executes the command the next time the instance starts.
	//
	// - EveryReboot: automatically executes the command every time the instance starts.
	//
	// - DryRun: performs only a dry run of the request without actually executing the command. The dry run checks items such as request parameters, instance execution environment, and Cloud Assistant Agent status.
	//
	// Default value:
	//
	// - If the `Frequency` parameter is not specified, the default value is `Once`.
	//
	// - If the `Frequency` parameter is specified, the command is executed as `Period` regardless of whether this parameter is set.
	//
	// Precautions:
	//
	// - You can call [StopInvocation](https://help.aliyun.com/document_detail/64838.html) to stop a pending or scheduled command.
	//
	// - If this parameter is set to `Period` or `EveryReboot`, you can call [DescribeInvocationResults](https://help.aliyun.com/document_detail/64845.html) and specify `IncludeHistory=true` to view the historical records of scheduled command executions.
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// The ID of the resource group for the command execution. When you specify this parameter:
	//
	// - If the ECS instance specified by InstanceId belongs to a non-default resource group, the ECS instance must belong to this resource group.
	//
	// - You can filter command execution results by specifying this parameter when you call [DescribeInvocations](https://help.aliyun.com/document_detail/64840.html) or [DescribeInvocationResults](https://help.aliyun.com/document_detail/64845.html).
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags that are used to filter instances. Array length: 0 to 20. You can run commands on instances that have the same tags in batches without specifying InstanceId.
	ResourceTag []*RunCommandShrinkRequestResourceTag `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty" type:"Repeated"`
	// The tags. Array length: 0 to 20.
	Tag []*RunCommandShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The mode in which the task is stopped (manually stopped or interrupted due to timeout). Valid values:
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
	// The timeout period for the command execution. Unit: seconds.
	//
	// A timeout occurs when the command cannot run due to process issues, missing modules, or missing Cloud Assistant Agent. After a timeout, the command process is forcefully terminated.
	//
	// Default value: 60.
	//
	// example:
	//
	// 3600
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The type of the command. Valid values:
	//
	// - RunBatScript: Bat command for Windows instances.
	//
	// - RunPowerShellScript: PowerShell command for Windows instances.
	//
	// - RunShellScript: Shell command for Linux instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// RunShellScript
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The username that is used to run the command on the ECS instance. The username can be up to 255 characters in length.
	//
	// - For Linux ECS instances, the command is run by the root user by default.
	//
	// - For Windows ECS instances, the command is run by the System user by default.
	//
	// You can also specify another existing user of the instance to run the command. Running Cloud Assistant commands as a regular user is more secure. For more information, see [Run Cloud Assistant commands as a regular user](https://help.aliyun.com/document_detail/203771.html).
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The name of the password for the user who runs the command on a Windows instance. The value can be up to 255 characters in length.
	//
	// example:
	//
	// axtSecretPassword
	WindowsPasswordName *string `json:"WindowsPasswordName,omitempty" xml:"WindowsPasswordName,omitempty"`
	// The working directory of the command on the ECS instance. Maximum length: 200 characters.
	//
	// Default value:
	//
	// - For Linux instances, the default directory is the home directory of the root user, which is `/root`.
	//
	// - For Windows instances, the default directory is the directory where the Cloud Assistant Agent process resides, such as `C:\\Windows\\System32`.
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
	// The tag key that is used to filter instances.
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
	// The tag key of the command execution. If this value is specified, it cannot be an empty string.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the command execution. The value can be an empty string.
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
