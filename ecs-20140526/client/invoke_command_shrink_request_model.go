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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The command ID. You can call the [DescribeCommands](https://help.aliyun.com/document_detail/64843.html) operation to query all available command IDs.
	//
	// > For public commands, you can run them by command name. For more information, see [View and run Cloud Assistant public commands](https://help.aliyun.com/document_detail/429635.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// c-e996287206324975b5fbe1d****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The container ID. Only 64-bit hexadecimal strings are supported. You can use the `docker://`, `containerd://`, or `cri-o://` prefix to specify the container runtime.
	//
	// Precautions:
	//
	// - If you specify this parameter, Cloud Assistant runs the script in the specified container on the instance.
	//
	// - If you specify this parameter, the command can only be run on Linux instances where the Cloud Assistant Agent version is 2.2.3.344 or later.
	//
	//     - For more information about how to view the Cloud Assistant Agent version, see [Install the Cloud Assistant Agent](https://help.aliyun.com/document_detail/64921.html).
	//
	//     - For more information about how to upgrade the Cloud Assistant Agent version, see [Upgrade or disable the upgrade of the Cloud Assistant Agent](https://help.aliyun.com/document_detail/134383.html).
	//
	// - If you specify this parameter, the `Username` parameter specified in this operation and the `WorkingDir` parameter specified in [CreateCommand](https://help.aliyun.com/document_detail/64844.html) do not take effect. The command can only be run by the default user of the container in the default working directory of the container. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// - If you specify this parameter, only shell scripts can be run in Linux containers. You cannot specify an interpreter for the script content by using a command such as `#!/usr/bin/python` at the beginning of the script. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
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
	// - If you specify this parameter, the command can only be run on Linux instances where the Cloud Assistant Agent version is 2.2.3.344 or later.
	//
	//     - For more information about how to view the Cloud Assistant Agent version, see [Install the Cloud Assistant Agent](https://help.aliyun.com/document_detail/64921.html).
	//
	//     - For more information about how to upgrade the Cloud Assistant Agent version, see [Upgrade or disable the upgrade of the Cloud Assistant Agent](https://help.aliyun.com/document_detail/134383.html).
	//
	// - If you specify this parameter, the `Username` parameter specified in this operation and the `WorkingDir` parameter specified in [CreateCommand](https://help.aliyun.com/document_detail/64844.html) do not take effect. The command can only be run by the default user of the container in the default working directory of the container. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// - If you specify this parameter, only shell scripts can be run in Linux containers. You cannot specify an interpreter for the script content by using a command such as `#!/usr/bin/python` at the beginning of the script. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// example:
	//
	// test-container
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The execution time of the scheduled command. Three scheduled execution modes are supported: execution at fixed intervals (based on rate expressions), one-time execution at a specified time, and clock-based scheduled execution (based on cron expressions).
	//
	// - Execution at fixed intervals: runs the command at a specified interval based on a rate expression. The interval can be set in seconds (s), minutes (m), hours (h), or days (d). This mode is suitable for scenarios where tasks are run at fixed intervals. The format is `rate(<interval value><interval unit>)`. For example, to run the command every 5 minutes, use `rate(5m)`. The following limits apply to execution at fixed intervals:
	//
	//     - The specified interval must not exceed 7 days or be less than 60 seconds, and must be greater than the timeout period of the scheduled task.
	//
	//     - The execution interval is based only on the fixed frequency and is independent of the actual time required to run the task. For example, if you set the command to run every 5 minutes and the task takes 2 minutes to complete, the next execution starts 3 minutes after the task is completed.
	//
	//     - The task is not run immediately upon creation. For example, if you set the command to run every 5 minutes, the command is not run immediately upon task creation, but starts 5 minutes after the task is created.
	//
	// - One-time execution at a specified time: runs the command once at the specified time zone and time point. The format is `at(yyyy-MM-dd HH:mm:ss <time zone>)`. If you do not specify a time zone, the UTC time zone is used by default. The time zone supports the following three formats:
	//
	//     - Time zone full name: for example, `Asia/Shanghai` (China/Shanghai time) and `America/Los_Angeles` (US/Los Angeles time).
	//
	//     - Time zone offset from Greenwich Mean Time: for example, `GMT+8:00` (East 8th zone) and `GMT-7:00` (West 7th zone). When you use the GMT format, leading zeros are not supported for the hour.
	//
	//     - Time zone abbreviation: only UTC (Coordinated Universal Time) is supported.
	//
	//   For example, to run the command once at 13:15:30 on June 6, 2022 in the China/Shanghai time zone, use `at(2022-06-06 13:15:30 Asia/Shanghai)`. To run the command once at 13:15:30 on June 6, 2022 in the West 7th zone, use `at(2022-06-06 13:15:30 GMT-7:00)`.
	//
	// - Clock-based scheduled execution (based on cron expressions): runs the command based on a cron expression. The format is `<seconds> <minutes> <hours> <day of month> <month> <day of week> <year (optional)> <time zone>`, which is `<cron expression> <time zone>`. The scheduled task execution time is calculated and executed based on the cron expression in the specified time zone. If you do not specify a time zone, the internal system time zone of the instance that runs the scheduled task is used by default. For more information about cron expressions, see [Cron expressions](https://help.aliyun.com/document_detail/64769.html). The time zone supports the following three formats:
	//
	//     - Time zone full name: for example, `Asia/Shanghai` (China/Shanghai time) and `America/Los_Angeles` (US/Los Angeles time).
	//
	//     - Time zone offset from Greenwich Mean Time: for example, `GMT+8:00` (East 8th zone) and `GMT-7:00` (West 7th zone). When you use the GMT format, leading zeros are not supported for the hour.
	//
	//     - Time zone abbreviation: only UTC (Coordinated Universal Time) is supported.
	//
	//   For example, to run the command once at 10:15 every day in 2022 in the China/Shanghai time zone, use `0 15 10 ? 	- 	- 2022 Asia/Shanghai`. To run the command every half hour from 10:00 to 11:30 every day in 2022 in the East 8th zone, use `0 0/30 10-11 	- 	- ? 2022 GMT+8:00`. To run the command every 5 minutes from 14:00 to 14:55 every day in October every two years starting from 2022 in the UTC time zone, use `0 0/5 14 	- 10 ? 2022/2 UTC`.
	//
	//     > The minimum specified time interval must be greater than or equal to the timeout period of the scheduled task and not less than 10 seconds.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The list of instance IDs on which the command is run. You can specify up to 100 instance IDs. Valid values of N: 1 to 100.
	//
	// You can also apply for a quota increase in Quota Center. The quota name is the maximum number of instances supported for command execution.
	//
	// example:
	//
	// ["i-bp185dy2o3o6n****"]
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The bootstrap program for script execution. The length cannot exceed 1 KB.
	//
	// example:
	//
	// python3 -u {{ACS::ScriptFileName|Ext(".py")}}
	Launcher *string `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	// The OSS data shipping configuration for the command execution output.
	//
	// - Format: `oss://${BucketName}/${Prefix}`. `${BucketName}` is the name of the OSS bucket to which the data is shipped, and `${Prefix}` is the directory prefix to which the data is shipped.
	//
	// example:
	//
	// oss://testBucket/testPrefix
	OssOutputDelivery *string `json:"OssOutputDelivery,omitempty" xml:"OssOutputDelivery,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key-value pairs of custom parameters passed when the command is run if the custom parameter feature is enabled. The number of custom parameters ranges from 0 to 10.
	//
	// - The key in the map cannot be an empty string and can be up to 64 characters in length.
	//
	// - The value in the map can be an empty string.
	//
	// - The total length of the custom parameters and the original command content cannot exceed 24 KB after Base64 encoding.
	//
	// - The set of custom parameter names must be a subset of the parameter set defined when the command was created. For parameters that are not passed, you can use an empty string instead.
	//
	// You can leave this parameter unspecified to disable custom parameters.
	//
	// example:
	//
	// {"name":"Jack", "accessKey":"LTAI************"}
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
	// - DryRun: only prechecks the request. The command is not actually run. The check items include the request parameters, instance execution environment, and Cloud Assistant Agent running status.
	//
	// Default value:
	//
	// - If you do not specify the `Frequency` parameter, the default value is `Once`.
	//
	// - If you specify the `Frequency` parameter, the value is processed as `Period` regardless of whether this parameter is set.
	//
	// Precautions:
	//
	// - You can call the [StopInvocation](https://help.aliyun.com/document_detail/64838.html) operation to stop a pending or scheduled command.
	//
	// - If you set this parameter to `Period` or `EveryReboot`, you can call the [DescribeInvocationResults](https://help.aliyun.com/document_detail/64845.html) operation and set `IncludeHistory` to `true` to view the historical records of the scheduled command execution.
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// The ID of the resource group in which the command is run. When you specify this parameter:
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
	// The list of tags used to filter instances. You can run the command on instances with the same tag in batches without specifying InstanceId.
	ResourceTag []*InvokeCommandShrinkRequestResourceTag `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty" type:"Repeated"`
	// The list of tags.
	Tag []*InvokeCommandShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The mode used when the task is stopped (manually stopped or interrupted due to execution timeout). Valid values:
	//
	// - Process: stops the current script process.
	//
	// - ProcessTree: stops the current process tree, which includes the script process and all its child processes.
	//
	// example:
	//
	// ProcessTree
	TerminationMode *string `json:"TerminationMode,omitempty" xml:"TerminationMode,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated and does not take effect even if it is specified.
	//
	// example:
	//
	// true
	Timed *bool `json:"Timed,omitempty" xml:"Timed,omitempty"`
	// The timeout period for running the command. Unit: seconds.
	//
	// - The value cannot be less than 10.
	//
	// - A timeout occurs when the command cannot be run due to process issues, missing modules, or a missing Cloud Assistant Agent. After a timeout, the command process is forcibly terminated.
	//
	// - If you do not specify this parameter, the timeout period specified when the command was created is used.
	//
	// - This value applies only to the current command execution and does not change the timeout period of the command itself.
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The name of the user that runs the command on the ECS instance. The name can be up to 255 characters in length.
	//
	// - For ECS instances that run Linux, the command is run by the root user by default.
	//
	// - For ECS instances that run Windows, the command is run by the System user by default.
	//
	// You can also specify another existing user on the instance to run the command. Running Cloud Assistant commands as a regular user is more secure. For more information, see [Configure a regular user to run Cloud Assistant commands](https://help.aliyun.com/document_detail/203771.html).
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The name of the password for the user that executes the command on a Windows instance. The name can be up to 255 characters in length.
	//
	// If you want to execute the command on a Windows instance as a non-default user (other than System), you must specify both `Username` and this parameter. To reduce the risk of password leakage, the plaintext password must be hosted in the parameter repository of System Operations Management. Here, you only pass the name of the password. For more information, see [Encrypt parameters](https://help.aliyun.com/document_detail/186828.html) and [Settings for regular users to execute Cloud Assistant commands](https://help.aliyun.com/document_detail/203771.html).
	//
	// > You do not need to pass this parameter when you run the command as the root user on a Linux instance or the System user on a Windows instance.
	//
	// example:
	//
	// axtSecretPassword
	WindowsPasswordName *string `json:"WindowsPasswordName,omitempty" xml:"WindowsPasswordName,omitempty"`
	// The directory in which the command is run on the ECS instance. The length cannot exceed 200 characters.
	//
	// - If you do not specify this parameter, the running directory specified when the command was created is used.
	//
	// - This value applies only to the current command execution and does not change the running directory of the command itself.
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

type InvokeCommandShrinkRequestResourceTag struct {
	// The tag key used to filter instances.
	//
	// Precautions:
	//
	// - This parameter conflicts with the InstanceId parameter and cannot be specified at the same time.
	//
	// - Valid values of N: 1 to 10. The value cannot be an empty string.
	//
	// - The number of instances under the tag cannot exceed the limit of InstanceId.N. If the number of instances exceeds the limit, control the number of instances by adding batch tags, such as batch: b1.
	//
	// - The value can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value used to filter instances.
	//
	// Precautions:
	//
	// - Valid values of N: 1 to 10.
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
	// The tag key for the command to execute. Valid values of N: 1 to 20. The value cannot be an empty string.
	//
	// If you use a single tag to filter resources, the resource count under the tag cannot exceed 1,000. If you use multiple tags to filter resources, the resource count of resources attached to all the specified tags cannot exceed 1,000. If the resource count exceeds 1,000, you must use the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to query the resources.
	//
	// The value can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the command execution. Valid values of N: 1 to 20. The value can be an empty string.
	//
	// The value can be up to 128 characters in length and cannot contain `http://` or `https://`.
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
