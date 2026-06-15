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
	// A client-generated token that is used to ensure the idempotence of the request. You must make sure that the token is unique among different requests. The `ClientToken` parameter can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The command content, which can be in plaintext or Base64-encoded. Note the following:
	//
	// - The size of the Base64-encoded command content is limited to 18 KB if `KeepCommand` is `true`, or 24 KB if `KeepCommand` is `false`.
	//
	// - If the command content is Base64-encoded, you must set `ContentEncoding` to `Base64`.
	//
	// - Set `EnableParameter` to `true` to enable the custom parameter feature in the command content.
	//
	//   - Define custom parameters by using the `{{}}` format. Spaces and line breaks before and after the parameter names within `{{}}` are ignored.
	//
	//   - You can define up to 20 custom parameters.
	//
	//   - A custom parameter name can contain only letters, digits, underscores (_), and hyphens (-). The name is case-insensitive and cannot start with `acs::`, which is reserved for built-in environment parameters.
	//
	//   - A custom parameter name can be up to 64 bytes long.
	//
	// - You can use built-in environment parameters, which Cloud Assistant automatically replaces with their corresponding values at runtime. The following built-in environment parameters are supported:
	//
	//   - `{{ACS::RegionId}}`: the region ID.
	//
	//   - `{{ACS::AccountId}}`: the UID of the Alibaba Cloud account.
	//
	//   - `{{ACS::InstanceId}}`: the instance ID. To use this parameter on multiple instances, the required Cloud Assistant Agent version is 2.2.3.309 or later for Linux instances, or 2.1.3.309 or later for Windows instances.
	//
	//     - Linux: 2.2.3.309
	//
	//     - Windows: 2.1.3.309
	//
	//   - `{{ACS::InstanceName}}`: the instance name. To use this parameter on multiple instances, the required Cloud Assistant Agent version is 2.2.3.344 or later for Linux instances, or 2.1.3.344 or later for Windows instances.
	//
	//     - Linux: 2.2.3.344
	//
	//     - Windows: 2.1.3.344
	//
	//   - `{{ACS::InvokeId}}`: the invocation ID. To use this parameter, the required Cloud Assistant Agent version is 2.2.3.309 or later for Linux instances, or 2.1.3.309 or later for Windows instances.
	//
	//     - Linux: 2.2.3.309
	//
	//     - Windows: 2.1.3.309
	//
	//   - `{{ACS::CommandId}}`: the command ID. To use this parameter, the required Cloud Assistant Agent version is 2.2.3.309 or later for Linux instances, or 2.1.3.309 or later for Windows instances.
	//
	//     - Linux: 2.2.3.309
	//
	//     - Windows: 2.1.3.309
	//
	// This parameter is required.
	//
	// example:
	//
	// ZWNobyAxMjM=
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The ID of the container. The ID must be a 64-bit hexadecimal string. You can add the `docker://`, `containerd://`, or `cri-o://` prefix to explicitly specify the container runtime.
	//
	// Notes:
	//
	// - If you specify this parameter, Cloud Assistant runs the script in the specified container of the instance.
	//
	// - This parameter is supported only on Linux instances with Cloud Assistant Agent version 2.2.3.344 or later.
	//
	// - If you specify this parameter, the specified `Username` and `WorkingDir` parameters are ignored. The command is run only by the default user in the default working directory of the container. For more information, see [Run commands in a container by using Cloud Assistant](https://help.aliyun.com/document_detail/456641.html).
	//
	// > In Linux containers, you can run only Shell scripts. You cannot use commands such as `#!/usr/bin/python` at the beginning of a script to specify an interpreter. For more information, see [Run commands in a container by using Cloud Assistant](https://help.aliyun.com/document_detail/456641.html).
	//
	// example:
	//
	// ab141ddfbacfe02d9dbc25966ed971536124527097398d419a6746873fea****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The name of the container.
	//
	// Notes:
	//
	// - If you specify this parameter, Cloud Assistant runs the script in the specified container of the instance.
	//
	// - This parameter is supported only on Linux instances with Cloud Assistant Agent version 2.2.3.344 or later.
	//
	// - If you specify this parameter, the specified `Username` and `WorkingDir` parameters are ignored. The command is run only by the default user in the default working directory of the container. For more information, see [Run commands in a container by using Cloud Assistant](https://help.aliyun.com/document_detail/456641.html).
	//
	// > In Linux containers, you can run only Shell scripts. You cannot use commands such as `#!/usr/bin/python` at the beginning of a script to specify an interpreter. For more information, see [Run commands in a container by using Cloud Assistant](https://help.aliyun.com/document_detail/456641.html).
	//
	// example:
	//
	// test-container
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The encoding mode of the command content (`CommandContent`). Valid values (case-insensitive):
	//
	// - `PlainText`: The command content is not encoded and is transmitted in plaintext.
	//
	// - `Base64`: The command content is Base64-encoded.
	//
	// Default value: `PlainText`. If you specify an invalid value, the value is automatically set to `PlainText`.
	//
	// example:
	//
	// Base64
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The description of the command. It can be up to 512 characters long and supports all character sets.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to use custom parameters in the command.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The schedule for the command. You can specify a rate expression, an at expression for one-time execution, or a cron expression.
	//
	// - **Fixed-interval execution**: Runs the command at fixed intervals defined by a rate expression. You can specify the interval in seconds (s), minutes (m), hours (h), or days (d). This method is suitable for tasks that must be run at fixed intervals. The format is `rate(<value><unit>)`. For example, to run a command every 5 minutes, use `rate(5m)`. The following limits apply to this method:
	//
	//   - The interval must be in the range of 60 seconds to 7 days, and must be longer than the timeout of the scheduled task.
	//
	//   - The interval is fixed and starts from the beginning of the previous execution, not from its completion.
	//
	//   - The task does not immediately run after it is created. For example, if you set an interval of 5 minutes, the first run starts 5 minutes after the task is created.
	//
	// - **One-time execution**: Run the command once at a specified time and in a specified time zone. The format is `at(yyyy-MM-dd HH:mm:ss <time_zone>)`. If you do not specify a time zone, UTC is used by default. The following time zone formats are supported:
	//
	//   - Full time zone name, such as `Asia/Shanghai` or `America/Los_Angeles`.
	//
	//   - Offset from GMT, such as `GMT+8:00` or `GMT-7:00`. When you use the GMT format, you cannot add a leading zero to the hour.
	//
	//   - Time zone abbreviation. Only `UTC` is supported.
	//
	//   Example 1: To run a task at 13:15:30 on June 6, 2022 in the `Asia/Shanghai` time zone, use `at(2022-06-06 13:15:30 Asia/Shanghai)`. Example 2: To run a task at 13:15:30 on June 6, 2022 in the `GMT-7:00` time zone, use `at(2022-06-06 13:15:30 GMT-7:00)`.
	//
	// - **Scheduled execution based on a cron expression**: Runs the command on a schedule defined by a cron expression. The format is `<second> <minute> <hour> <day_of_month> <month> <day_of_week> <year (optional)> <time_zone>`, or `<cron_expression> <time_zone>`. The task is run based on the cron expression in the specified time zone. If you do not specify a time zone, the system time zone of the instance where the task is run is used by default. For more information about cron expressions, see [Cron expressions](https://help.aliyun.com/document_detail/64769.html). The following time zone formats are supported:
	//
	//   - Full time zone name, such as `Asia/Shanghai` or `America/Los_Angeles`.
	//
	//   - Offset from GMT, such as `GMT+8:00` or `GMT-7:00`. When you use the GMT format, you cannot add a leading zero to the hour.
	//
	//   - Time zone abbreviation. Only `UTC` is supported.
	//
	//     For example, to run a command at 10:15 every day in 2022 in the `Asia/Shanghai` time zone, use `0 15 10 ? 	- 	- 2022 Asia/Shanghai`. To run a command every 30 minutes from 10:00 to 11:30 every day in 2022 in the `GMT+8:00` time zone, use `0 0/30 10-11 	- 	- ? 2022 GMT+8:00`. To run a command every 5 minutes from 14:00 to 14:55 every day in October of every two years starting from 2022 in `UTC`, use `0 0/5 14 	- 10 ? 2022/2 UTC`.
	//
	//   > The minimum interval must be greater than or equal to the timeout of the scheduled task, and cannot be less than 10 seconds.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The IDs of the ECS instances on which to run the command. You can specify from 1 to 100 instance IDs.
	//
	// If any specified instance does not meet the execution requirements, the entire operation fails.
	//
	// You can apply for a quota increase in Quota Center. The quota is named Maximum number of instances supported per command execution.
	//
	// example:
	//
	// i-bp185dy2o3o6neg****
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// Specifies whether to save the command after it is run. Valid values:
	//
	// - `true`: Saves the command. You can then re-run it by calling InvokeCommand. Saved commands count towards your Cloud Assistant command quota.
	//
	// - `false`: Does not save the command. The command is deleted after execution and does not count towards your quota.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	KeepCommand *bool `json:"KeepCommand,omitempty" xml:"KeepCommand,omitempty"`
	// The launcher that is used to run the script. The value can be up to 1 KB in length.
	//
	// example:
	//
	// python3 -u {{ACS::ScriptFileName|Ext(".py")}}
	Launcher *string `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	// The name of the command. It can be up to 128 characters long and supports all character sets.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The OSS delivery configuration for the command output.
	//
	// - Format: oss\\://${BucketName}/${Prefix}, where ${BucketName} is the name of the destination OSS bucket and ${Prefix} is the destination prefix.
	//
	// example:
	//
	// oss://testBucket/testPrefix
	OssOutputDelivery *string `json:"OssOutputDelivery,omitempty" xml:"OssOutputDelivery,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key-value pairs for custom parameters. For example, if `CommandContent` is `echo {{name}}`, setting `Parameters` to `{"name":"Jack"}` results in the command `echo Jack` being run.
	//
	// You can specify 0 to 10 key-value pairs. Note the following:
	//
	// - The key cannot be an empty string and can be up to 64 characters in length.
	//
	// - The value can be an empty string.
	//
	// - After Base64 encoding, the total size of the custom parameters and the original command content is limited to 18 KB if `KeepCommand` is `true`, or 24 KB if `KeepCommand` is `false`.
	//
	// - The set of custom parameter names that you specify must be a subset of the parameters defined in `CommandContent`. The value of an omitted parameter defaults to an empty string.
	//
	// By default, this parameter is empty, which indicates that no custom parameters are used.
	//
	// example:
	//
	// {"name":"Jack", "accessKey":"LTAI*************"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the latest Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The execution mode of the command. Valid values:
	//
	// - `Once`: The command is immediately run.
	//
	// - `Period`: Runs the command as a scheduled task. This mode requires the `Frequency` parameter.
	//
	// - `NextRebootOnly`: The command is automatically run the next time the instance starts.
	//
	// - `EveryReboot`: The command is automatically run every time the instance starts.
	//
	// - `DryRun`: Performs a dry run to check parameters and the environment without actually running the command.
	//
	// Default value:
	//
	// - If the `Frequency` parameter is not specified, the default value is `Once`.
	//
	// - If `Frequency` is specified, this parameter is automatically set to `Period`.
	//
	// Notes:
	//
	// - You can call the [StopInvocation](https://help.aliyun.com/document_detail/64838.html) operation to stop pending or scheduled commands.
	//
	// - If you set this parameter to `Period` or `EveryReboot`, you can call the [DescribeInvocationResults](https://help.aliyun.com/document_detail/64845.html) operation and set `IncludeHistory=true` to query the historical execution records of the scheduled command.
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// The ID of the resource group for the command execution. When you specify this parameter, the following rules apply:
	//
	// - If an ECS instance specified by `InstanceId` is in a non-default resource group, it must belong to the resource group specified by this parameter.
	//
	// - You can use this parameter to filter command execution results when you call the [DescribeInvocations](https://help.aliyun.com/document_detail/64840.html) or [DescribeInvocationResults](https://help.aliyun.com/document_detail/64845.html) operation.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Tags used to filter instances for command execution. This allows you to run the command on all instances with matching tags, as an alternative to specifying instance IDs. The array can contain 0 to 20 tags.
	ResourceTag []*RunCommandRequestResourceTag `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty" type:"Repeated"`
	// An array of tag pairs. The array can contain 0 to 20 tags.
	Tag []*RunCommandRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The mode for stopping the task when it is manually stopped or times out. Valid values:
	//
	// - `Process`: Stops the current script process.
	//
	// - `ProcessTree`: Stops the current process tree. A process tree includes the current script process and all of its subprocesses.
	//
	// example:
	//
	// ProcessTree
	TerminationMode *string `json:"TerminationMode,omitempty" xml:"TerminationMode,omitempty"`
	// > This parameter is deprecated and no longer has any effect.
	//
	// example:
	//
	// true
	Timed *bool `json:"Timed,omitempty" xml:"Timed,omitempty"`
	// The command execution timeout, in seconds.
	//
	// A timeout forcibly terminates the command process if the command fails to run due to exceptions, such as a process conflict, a missing module, or a disabled Cloud Assistant Agent.
	//
	// Default value: 60.
	//
	// example:
	//
	// 3600
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The type of the command. Valid values:
	//
	// - `RunBatScript`: Bat commands for Windows instances.
	//
	// - `RunPowerShellScript`: PowerShell commands for Windows instances.
	//
	// - `RunShellScript`: Shell commands for Linux instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// RunShellScript
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The name of the user that runs the command on the ECS instance. The name can be up to 255 characters in length.
	//
	// - Default on Linux: `root`.
	//
	// - Default on Windows: `System`.
	//
	// You can specify another existing user on the instance to run the command. Running Cloud Assistant commands as a standard user is more secure. For more information, see [Run Cloud Assistant commands as a standard user](https://help.aliyun.com/document_detail/203771.html).
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The name of the password of the user that runs the command on a Windows instance. The name can be up to 255 characters in length.
	//
	// To run a command as a non-default user on a Windows instance, you must specify both `Username` and `WindowsPasswordName`. To reduce the risk of password leaks, we recommend storing the password in OOS Parameter Store and providing the parameter name here. For more information, see [Encryption parameters](https://help.aliyun.com/document_detail/186828.html) and [Run Cloud Assistant commands as a standard user](https://help.aliyun.com/document_detail/203771.html).
	//
	// > You do not need to specify this parameter when you run a command as the `root` user on a Linux instance or as the `System` user on a Windows instance.
	//
	// example:
	//
	// axtSecretPassword
	WindowsPasswordName *string `json:"WindowsPasswordName,omitempty" xml:"WindowsPasswordName,omitempty"`
	// The working directory for the command on the instance. The path can be up to 200 characters long.
	//
	// Default values:
	//
	// - For Linux instances, the default is the home directory of the `root` user (`/root`).
	//
	// - For Windows instances, the default is the directory of the Cloud Assistant Agent process, such as `C:\\Windows\\System32`.
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
	// The tag key that is used to filter instances.
	//
	// Notes:
	//
	// - You cannot specify both this parameter and the InstanceId parameter.
	//
	// - The tag key cannot be an empty string.
	//
	// - The number of instances matching the specified tag cannot exceed the per-execution instance limit (100 by default). If the number of matching instances exceeds this limit, you can use additional tags, such as `batch:b1`, to refine the selection.
	//
	// - The value can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It also cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value that is used to filter instances.
	//
	// Notes:
	//
	// - The value can be an empty string.
	//
	// - The value can be up to 128 characters in length and cannot contain `http://` or `https://`.
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
	// The tag key for the command execution. The key cannot be an empty string.
	//
	// The key can be up to 64 characters long and cannot start with `aliyun` or `acs:`. It also cannot contain `http://` or `https://`.
	//
	// The value can be up to 64 characters long and cannot start with `aliyun` or `acs:` or contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value for the command execution. The value can be an empty string.
	//
	// The value can be up to 128 characters long and cannot contain `http://` or `https://`.
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
