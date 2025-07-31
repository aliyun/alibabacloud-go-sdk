// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeCommandShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *InvokeCommandShrinkRequest
	GetClientToken() *string
	SetCommandId(v string) *InvokeCommandShrinkRequest
	GetCommandId() *string
	SetContainerId(v string) *InvokeCommandShrinkRequest
	GetContainerId() *string
	SetContainerName(v string) *InvokeCommandShrinkRequest
	GetContainerName() *string
	SetFrequency(v string) *InvokeCommandShrinkRequest
	GetFrequency() *string
	SetInstanceId(v []*string) *InvokeCommandShrinkRequest
	GetInstanceId() []*string
	SetLauncher(v string) *InvokeCommandShrinkRequest
	GetLauncher() *string
	SetOssOutputDelivery(v string) *InvokeCommandShrinkRequest
	GetOssOutputDelivery() *string
	SetOwnerAccount(v string) *InvokeCommandShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *InvokeCommandShrinkRequest
	GetOwnerId() *int64
	SetParametersShrink(v string) *InvokeCommandShrinkRequest
	GetParametersShrink() *string
	SetRegionId(v string) *InvokeCommandShrinkRequest
	GetRegionId() *string
	SetRepeatMode(v string) *InvokeCommandShrinkRequest
	GetRepeatMode() *string
	SetResourceGroupId(v string) *InvokeCommandShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *InvokeCommandShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *InvokeCommandShrinkRequest
	GetResourceOwnerId() *int64
	SetResourceTag(v []*InvokeCommandShrinkRequestResourceTag) *InvokeCommandShrinkRequest
	GetResourceTag() []*InvokeCommandShrinkRequestResourceTag
	SetTag(v []*InvokeCommandShrinkRequestTag) *InvokeCommandShrinkRequest
	GetTag() []*InvokeCommandShrinkRequestTag
	SetTerminationMode(v string) *InvokeCommandShrinkRequest
	GetTerminationMode() *string
	SetTimed(v bool) *InvokeCommandShrinkRequest
	GetTimed() *bool
	SetTimeout(v int64) *InvokeCommandShrinkRequest
	GetTimeout() *int64
	SetUsername(v string) *InvokeCommandShrinkRequest
	GetUsername() *string
	SetWindowsPasswordName(v string) *InvokeCommandShrinkRequest
	GetWindowsPasswordName() *string
	SetWorkingDir(v string) *InvokeCommandShrinkRequest
	GetWorkingDir() *string
}

type InvokeCommandShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The command ID. You can call the [DescribeCommands](https://help.aliyun.com/document_detail/64843.html) operation to query all available command IDs.
	//
	// >  Common Cloud Assistant commands can be run based on their names. For more information, see [View and run common Cloud Assistant commands](https://help.aliyun.com/document_detail/429635.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// c-e996287206324975b5fbe1d****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The ID of the container. Only 64-bit hexadecimal strings are supported. You can use container IDs that are prefixed with `docker://`, `containerd://`, or `cri-o://` to specify container runtimes.
	//
	// Take note of the following items:
	//
	// 	- If this parameter is specified, Cloud Assistant runs the command in the specified container of the instance.
	//
	// 	- If this parameter is specified, the command can run only on Linux instances on which Cloud Assistant Agent 2.2.3.344 or later is installed.
	//
	//     	- For information about how to query the version of Cloud Assistant Agent, see [Install Cloud Assistant Agent](https://help.aliyun.com/document_detail/64921.html).
	//
	//     	- For information about how to upgrade Cloud Assistant Agent, see [Upgrade or disable upgrades for Cloud Assistant Agent](https://help.aliyun.com/document_detail/134383.html).
	//
	// 	- If this parameter is specified, the `Username` parameter that is specified in a request to call this operation and the `WorkingDir` parameter that is specified in a request to call the [CreateCommand](https://help.aliyun.com/document_detail/64844.html) operation do not take effect. You can run the command only in the default working directory of the container by using the default user of the container. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// 	- If this parameter is specified, only shell scripts can be run in Linux containers. You cannot add a command in the format similar to `#!/usr/bin/python` at the beginning of a script to specify a script interpreter. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// example:
	//
	// ab141ddfbacfe02d9dbc25966ed971536124527097398d419a6746873fea****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The name of the container.
	//
	// Take note of the following items:
	//
	// 	- If this parameter is specified, Cloud Assistant runs the command in the specified container of the instance.
	//
	// 	- If this parameter is specified, the command can run only on Linux instances on which Cloud Assistant Agent 2.2.3.344 or later is installed.
	//
	//     	- For information about how to query the version of Cloud Assistant Agent, see [Install Cloud Assistant Agent](https://help.aliyun.com/document_detail/64921.html).
	//
	//     	- For information about how to upgrade Cloud Assistant Agent, see [Upgrade or disable upgrades for Cloud Assistant Agent](https://help.aliyun.com/document_detail/134383.html).
	//
	// 	- If this parameter is specified, the `Username` parameter that is specified in a request to call this operation and the `WorkingDir` parameter that is specified in a request to call the [CreateCommand](https://help.aliyun.com/document_detail/64844.html) operation do not take effect. You can run the command only in the default working directory of the container by using the default user of the container. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// 	- If this parameter is specified, only shell scripts can be run in Linux containers. You cannot add a command in the format similar to `#!/usr/bin/python` at the beginning of a script to specify a script interpreter. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// example:
	//
	// test-container
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The schedule on which to run the command. You can configure a command to run at a fixed interval based on a rate expression, run only once at a specific time, or run at specific times based on a cron expression.
	//
	// 	- To run a command at a fixed interval, use a rate expression to specify the interval. You can specify the interval in seconds, minutes, hours, or days. This option is suitable for scenarios in which tasks need to be executed at a fixed interval. Specify the interval in the following format: `rate(<Execution interval value><Execution interval unit>)`. For example, specify `rate(5m)` to run the command every 5 minutes. When you specify an interval, take note of the following limits:
	//
	//     	- The interval can be anywhere from 60 seconds to 7 days, but must be longer than the timeout period of the scheduled task.
	//
	//     	- The interval is the amount of time that elapses between two consecutive executions. The interval is irrelevant to the amount of time that is required to run the command once. For example, assume that you set the interval to 5 minutes and that it takes 2 minutes to run the command each time. Each time the command is run, the system waits 3 minutes before the system runs the command again.
	//
	//     	- A task is not immediately executed after the task is created. For example, assume that you set the interval to 5 minutes for a task. The task begins to be executed 5 minutes after it is created.
	//
	// 	- To run a command only once at a specific time, specify a point in time and a time zone. Specify the point in time in the `at(yyyy-MM-dd HH:mm:ss <Time zone>)` format, which indicates `at(Year-Month-Day Hour:Minute:Second <Time zone>)`. If you do not specify a time zone, the Coordinated Universal Time (UTC) time zone is used by default. You can specify a time zone in the following forms:
	//
	//     	- The time zone name. Examples: `Asia/Shanghai` and `America/Los_Angeles`.
	//
	//     	- The time offset from GMT. Examples: `GMT+8:00` (UTC+8) and `GMT-7:00` (UTC-7). If you use the GMT format, you cannot add leading zeros to the hour value.
	//
	//     	- The time zone abbreviation. Only UTC is supported.
	//
	//     For example, to configure a command to run only once at 13:15:30 on June 6, 2022 (Shanghai time), set the time to `at(2022-06-06 13:15:30 Asia/Shanghai)`. To configure a command to run only once at 13:15:30 on June 6, 2022 (UTC-7), set the time to `at(2022-06-06 13:15:30 GMT-7:00)`.
	//
	// 	- To run a command at specific times, use a cron expression to define the schedule. Specify a schedule in the `<Cron expression> <Time zone>` format. The cron expression is in the `<seconds> <minutes> <hours> <day of the month> <month> <day of the week> <year (optional)>` format. The system calculates the execution times of the command based on the specified cron expression and time zone and runs the command as scheduled. If you do not specify a time zone, the system time zone of the instance on which you want to run the command is used by default. For more information about cron expressions, see [Cron expressions](https://help.aliyun.com/document_detail/64769.html). You can specify a time zone in the following forms:
	//
	//     	- The time zone name. Examples: `Asia/Shanghai` and `America/Los_Angeles`.
	//
	//     	- The time offset from GMT. Examples: `GMT+8:00` (UTC+8) and `GMT-7:00` (UTC-7). If you use the GMT format, you cannot add leading zeros to the hour value.
	//
	//     	- The time zone abbreviation. Only UTC is supported. For example, to configure a command to run at 10:15:00 every day in 2022 (Shanghai time), set the schedule to `0 15 10 ? 	- 	- 2022 Asia/Shanghai`. To configure a command to run every half an hour from 10:00:00 to 11:30:00 every day in 2022 (UTC+8), set the schedule to `0 0/30 10-11 	- 	- ? 2022 GMT+8:00`. To configure a command to run every 5 minutes from 14:00:00 to 14:55:00 every October every two years from 2022 in UTC, set the schedule to `0 0/5 14 	- 10 ? 2022/2 UTC`.
	//
	//     **
	//
	//     **Note*	- The minimum interval must be 10 seconds or more and cannot be shorter than the timeout period of scheduled executions.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The IDs of instances on which you want to run the command. You can specify up to 100 instance IDs in each request. Valid values of N: 1 to 100.
	//
	// You can apply for a quota increase in the Quota Center console. The quota name is Maximum number of instances supported for command execution.
	//
	// example:
	//
	// i-bp185dy2o3o6n****
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The launcher for script execution. The value cannot exceed 1 KB in length.
	//
	// example:
	//
	// python3 -u {{ACS::ScriptFileName|Ext(".py")}}
	Launcher          *string `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	OssOutputDelivery *string `json:"OssOutputDelivery,omitempty" xml:"OssOutputDelivery,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key-value pairs of custom parameters to pass in when the custom parameter feature is enabled. You can specify up to 10 custom parameters.
	//
	// 	- Each key in a Map collection cannot be an empty string, and can be up to 64 characters in length.
	//
	// 	- Each value in a Map collection can be an empty string.
	//
	// 	- The size of the command after Base64 encoding, including the custom parameters and the original command content, cannot exceed 18 KB.
	//
	// 	- The custom parameter names that are specified by Parameters must be included in the custom parameter names that you specified when you created the command. You can use empty strings to represent the custom parameters that are not specified.
	//
	// If you want to disable the custom parameter feature, you can leave this parameter empty.
	//
	// example:
	//
	// {"name":"Jack", "accessKey":"LTAI************"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID of the command. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies how to run the command. Valid values:
	//
	// 	- Once: immediately runs the command.
	//
	// 	- Period: runs the command based on a schedule. If you set this parameter to `Period`, you must also configure the `Frequency` parameter.
	//
	// 	- NextRebootOnly: runs the command the next time the instance is started.
	//
	// 	- EveryReboot: The command is run every time the instances start.
	//
	// 	- DryRun: Specifies whether to perform only a dry run, without performing the actual request. The command does not take effect. The system checks the request, including the request parameters, instance execution environment, and Cloud Assistant Agent running status.
	//
	// Default value:
	//
	// 	- If you do not specify `Frequency`, the default value is `Once`.
	//
	// 	- If you specify the `Frequency` parameter, `Period` is used as the value of RepeatMode regardless of whether RepeatMode is set to Period.
	//
	// Take note of the following items when you specify this property:
	//
	// 	- You can call the [StopInvocation](https://help.aliyun.com/document_detail/64838.html) operation to stop the pending or scheduled executions of the command.
	//
	// 	- If you set this parameter to `Period` or `EveryReboot`, you can call the [DescribeInvocationResults](https://help.aliyun.com/document_detail/64845.html) operation with `IncludeHistory` set to true to query the results of historical scheduled executions.
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// The ID of the resource group to which to assign the command executions. When you set this parameter, take note of the following items:
	//
	// 	- The instances specified by InstanceId.N must belong to the specified resource group.
	//
	// 	- After the command is run, you can call the [DescribeInvocations](https://help.aliyun.com/document_detail/64840.html) or [DescribeInvocationResults](https://help.aliyun.com/document_detail/64845.html) operation with ResourceGroupId set to query the execution results in the specified resource group.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the instance. If you do not specify InstanceId.N, the command is run on the instances that have the specified tags.
	ResourceTag []*InvokeCommandShrinkRequestResourceTag `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty" type:"Repeated"`
	// The tags of the command.
	Tag []*InvokeCommandShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies how to stop the command task when a command execution is manually stopped or times out. Valid values:
	//
	// 	- Process: stops the process of the command.
	//
	// 	- ProcessTree: stops the process tree of the command. In this case, the process of the command and all subprocesses of the process are stopped.
	//
	// example:
	//
	// ProcessTree
	TerminationMode *string `json:"TerminationMode,omitempty" xml:"TerminationMode,omitempty"`
	// >  This parameter is no longer used and does not take effect.
	//
	// example:
	//
	// true
	Timed *bool `json:"Timed,omitempty" xml:"Timed,omitempty"`
	// The timeout period for the command execution. Unit: seconds.
	//
	// 	- The timeout period cannot be less than 10 seconds.
	//
	// 	- A timeout error occurs if the command cannot be run because the process slows down or because a specific module or Cloud Assistant Agent does not exist. When the specified timeout period ends, the command process is forcefully terminated.
	//
	// 	- If you do not specify this parameter, the timeout period that is specified when the command is created is used.
	//
	// 	- This timeout period is applicable only to this execution. The timeout period of the command is not modified.
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The username to use to run the command on the ECS instances. The username cannot exceed 255 characters in length.
	//
	// 	- For Linux instances, the root username is used by default.
	//
	// 	- For Windows instances, the System username is used by default.
	//
	// You can also specify other usernames that already exist in the instances to run the command. For security purposes, we recommend that you run Cloud Assistant commands as a regular user. For more information, see [Run Cloud Assistant commands as a regular user](https://help.aliyun.com/document_detail/203771.html).
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The name of the password to use to run the command on a Windows instance. The name cannot exceed 255 characters in length.
	//
	// If you do not want to use the default System user to run the command on Windows instances, specify both WindowsPasswordName and `Username`. To mitigate the risk of password leaks, the password is stored in plaintext in CloudOps Orchestration Service (OOS) Parameter Store, and only the name of the password is passed in by using WindowsPasswordName. For more information, see [Manage encryption parameters](https://help.aliyun.com/document_detail/186828.html) and [Run Cloud Assistant commands as a regular user](https://help.aliyun.com/document_detail/203771.html).
	//
	// >  If you use the root username for Linux instances or the System username for Windows instances to run the command, you do not need to specify WindowsPasswordName.
	//
	// example:
	//
	// axtSecretPassword
	WindowsPasswordName *string `json:"WindowsPasswordName,omitempty" xml:"WindowsPasswordName,omitempty"`
	// The execution path of the command on ECS instances. The value can be up to 200 characters in length.
	//
	// 	- If you do not specify this parameter, the execution path specified when the command is created is used.
	//
	// 	- This execution path is applicable only to this task. The execution path of the command is not changed.
	//
	// example:
	//
	// /home/user
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s InvokeCommandShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeCommandShrinkRequest) GoString() string {
	return s.String()
}

func (s *InvokeCommandShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *InvokeCommandShrinkRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *InvokeCommandShrinkRequest) GetContainerId() *string {
	return s.ContainerId
}

func (s *InvokeCommandShrinkRequest) GetContainerName() *string {
	return s.ContainerName
}

func (s *InvokeCommandShrinkRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *InvokeCommandShrinkRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *InvokeCommandShrinkRequest) GetLauncher() *string {
	return s.Launcher
}

func (s *InvokeCommandShrinkRequest) GetOssOutputDelivery() *string {
	return s.OssOutputDelivery
}

func (s *InvokeCommandShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *InvokeCommandShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *InvokeCommandShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *InvokeCommandShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InvokeCommandShrinkRequest) GetRepeatMode() *string {
	return s.RepeatMode
}

func (s *InvokeCommandShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *InvokeCommandShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *InvokeCommandShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *InvokeCommandShrinkRequest) GetResourceTag() []*InvokeCommandShrinkRequestResourceTag {
	return s.ResourceTag
}

func (s *InvokeCommandShrinkRequest) GetTag() []*InvokeCommandShrinkRequestTag {
	return s.Tag
}

func (s *InvokeCommandShrinkRequest) GetTerminationMode() *string {
	return s.TerminationMode
}

func (s *InvokeCommandShrinkRequest) GetTimed() *bool {
	return s.Timed
}

func (s *InvokeCommandShrinkRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *InvokeCommandShrinkRequest) GetUsername() *string {
	return s.Username
}

func (s *InvokeCommandShrinkRequest) GetWindowsPasswordName() *string {
	return s.WindowsPasswordName
}

func (s *InvokeCommandShrinkRequest) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *InvokeCommandShrinkRequest) SetClientToken(v string) *InvokeCommandShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetCommandId(v string) *InvokeCommandShrinkRequest {
	s.CommandId = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetContainerId(v string) *InvokeCommandShrinkRequest {
	s.ContainerId = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetContainerName(v string) *InvokeCommandShrinkRequest {
	s.ContainerName = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetFrequency(v string) *InvokeCommandShrinkRequest {
	s.Frequency = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetInstanceId(v []*string) *InvokeCommandShrinkRequest {
	s.InstanceId = v
	return s
}

func (s *InvokeCommandShrinkRequest) SetLauncher(v string) *InvokeCommandShrinkRequest {
	s.Launcher = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetOssOutputDelivery(v string) *InvokeCommandShrinkRequest {
	s.OssOutputDelivery = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetOwnerAccount(v string) *InvokeCommandShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetOwnerId(v int64) *InvokeCommandShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetParametersShrink(v string) *InvokeCommandShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetRegionId(v string) *InvokeCommandShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetRepeatMode(v string) *InvokeCommandShrinkRequest {
	s.RepeatMode = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetResourceGroupId(v string) *InvokeCommandShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetResourceOwnerAccount(v string) *InvokeCommandShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetResourceOwnerId(v int64) *InvokeCommandShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetResourceTag(v []*InvokeCommandShrinkRequestResourceTag) *InvokeCommandShrinkRequest {
	s.ResourceTag = v
	return s
}

func (s *InvokeCommandShrinkRequest) SetTag(v []*InvokeCommandShrinkRequestTag) *InvokeCommandShrinkRequest {
	s.Tag = v
	return s
}

func (s *InvokeCommandShrinkRequest) SetTerminationMode(v string) *InvokeCommandShrinkRequest {
	s.TerminationMode = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetTimed(v bool) *InvokeCommandShrinkRequest {
	s.Timed = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetTimeout(v int64) *InvokeCommandShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetUsername(v string) *InvokeCommandShrinkRequest {
	s.Username = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetWindowsPasswordName(v string) *InvokeCommandShrinkRequest {
	s.WindowsPasswordName = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetWorkingDir(v string) *InvokeCommandShrinkRequest {
	s.WorkingDir = &v
	return s
}

func (s *InvokeCommandShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type InvokeCommandShrinkRequestResourceTag struct {
	// The key of tag N of the instance.
	//
	// Take note of the following items:
	//
	// 	- This parameter and InstanceId.N are mutually exclusive.
	//
	// 	- Valid values of N: 1 to 10. The tag key cannot be an empty string.
	//
	// 	- The number of instances that have the specified tags cannot exceed 100. If more than 100 instances have the specified tags, we recommend that you use batch tags such as batch: b1 to group the instances into batches of up to 100 instances.
	//
	// 	- The tag key can be up to 64 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the instance.
	//
	// Take note of the following items:
	//
	// 	- Valid values of N: 1 to 10.
	//
	// 	- The tag value can be an empty string.
	//
	// 	- The tag value can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s InvokeCommandShrinkRequestResourceTag) String() string {
	return dara.Prettify(s)
}

func (s InvokeCommandShrinkRequestResourceTag) GoString() string {
	return s.String()
}

func (s *InvokeCommandShrinkRequestResourceTag) GetKey() *string {
	return s.Key
}

func (s *InvokeCommandShrinkRequestResourceTag) GetValue() *string {
	return s.Value
}

func (s *InvokeCommandShrinkRequestResourceTag) SetKey(v string) *InvokeCommandShrinkRequestResourceTag {
	s.Key = &v
	return s
}

func (s *InvokeCommandShrinkRequestResourceTag) SetValue(v string) *InvokeCommandShrinkRequestResourceTag {
	s.Value = &v
	return s
}

func (s *InvokeCommandShrinkRequestResourceTag) Validate() error {
	return dara.Validate(s)
}

type InvokeCommandShrinkRequestTag struct {
	// The key of tag N to add to the command task. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// If a single tag is specified to query resources, up to 1,000 resources that have this tag added can be displayed in the response. If multiple tags are specified to query resources, up to 1,000 resources that have all these tags added can be displayed in the response. To query more than 1,000 resources that have specified tags added, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// The tag key can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the command task. Valid values of N: 1 to 20. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s InvokeCommandShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s InvokeCommandShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *InvokeCommandShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *InvokeCommandShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *InvokeCommandShrinkRequestTag) SetKey(v string) *InvokeCommandShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *InvokeCommandShrinkRequestTag) SetValue(v string) *InvokeCommandShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *InvokeCommandShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
