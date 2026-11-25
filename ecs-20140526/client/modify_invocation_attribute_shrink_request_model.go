// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInvocationAttributeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyInvocationAttributeShrinkRequest
	GetClientToken() *string
	SetCommandContent(v string) *ModifyInvocationAttributeShrinkRequest
	GetCommandContent() *string
	SetContentEncoding(v string) *ModifyInvocationAttributeShrinkRequest
	GetContentEncoding() *string
	SetEnableParameter(v bool) *ModifyInvocationAttributeShrinkRequest
	GetEnableParameter() *bool
	SetFrequency(v string) *ModifyInvocationAttributeShrinkRequest
	GetFrequency() *string
	SetInstanceId(v []*string) *ModifyInvocationAttributeShrinkRequest
	GetInstanceId() []*string
	SetInvokeId(v string) *ModifyInvocationAttributeShrinkRequest
	GetInvokeId() *string
	SetOwnerAccount(v string) *ModifyInvocationAttributeShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInvocationAttributeShrinkRequest
	GetOwnerId() *int64
	SetParametersShrink(v string) *ModifyInvocationAttributeShrinkRequest
	GetParametersShrink() *string
	SetRegionId(v string) *ModifyInvocationAttributeShrinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyInvocationAttributeShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInvocationAttributeShrinkRequest
	GetResourceOwnerId() *int64
}

type ModifyInvocationAttributeShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **ClientToken*	- parameter can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The modified command content. The command content can be plaintext or Base64-encoded. Take note of the following items:
	//
	// - The size of the Base64-encoded command content cannot exceed 24 KB.
	//
	// - If the command content is Base64-encoded, you must set `ContentEncoding` to `Base64`.
	//
	// - If you set `EnableParameter` to `true`, you can use custom parameters in the command content:
	//
	//     - Custom parameters are enclosed in `{{}}`. Spaces and line breaks before and after the parameter name inside `{{}}` are ignored.
	//
	//     - You can specify a maximum of 20 custom parameters.
	//
	//     - A custom parameter name can contain letters, digits, hyphens (-), and underscores (_). The `acs::` prefix is not supported for specifying non-built-in environment variables. Other characters are not supported. Parameter names are not case-sensitive.
	//
	//     - A single custom parameter name cannot exceed 64 bytes in length.
	//
	// - You can specify built-in environment variables as custom parameters. When the command is executed, you do not need to manually assign values to the parameters. Cloud Assistant automatically replaces the parameters with the corresponding values in the environment. The following built-in environment variables are supported:
	//
	//     - `{{ACS::RegionId}}`: the region ID.
	//
	//     - `{{ACS::AccountId}}`: the UID of the Alibaba Cloud account.
	//
	//     - `{{ACS::InstanceId}}`: the instance ID. If the command is sent to multiple instances and you want to specify `{{ACS::InstanceId}}` as a built-in environment variable, make sure that the Cloud Assistant Agent version is not earlier than the following versions:
	//
	//         - Linux: 2.2.3.309
	//
	//         - Windows: 2.1.3.309
	//
	//     - `{{ACS::InstanceName}}`: the instance name. If the command is sent to multiple instances and you want to specify `{{ACS::InstanceName}}` as a built-in environment variable, make sure that the Cloud Assistant Agent version is not earlier than the following versions:
	//
	//         - Linux: 2.2.3.344
	//
	//         - Windows: 2.1.3.344
	//
	//     - `{{ACS::InvokeId}}`: the command execution ID. If you want to specify `{{ACS::InvokeId}}` as a built-in environment variable, make sure that the Cloud Assistant Agent version is not earlier than the following versions:
	//
	//         - Linux: 2.2.3.309
	//
	//         - Windows: 2.1.3.309
	//
	//     - `{{ACS::CommandId}}`: the command ID. If you call this operation to execute a command and want to specify `{{ACS::CommandId}}` as a built-in environment variable, make sure that the Cloud Assistant Agent version is not earlier than the following versions:
	//
	//         - Linux: 2.2.3.309
	//
	//         - Windows: 2.1.3.309
	//
	// example:
	//
	// ZWNobyAxMjM=
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The encoding method of the command content specified by `CommandContent`. Valid values (case-insensitive):
	//
	// - PlainText: The command content is not encoded and is transmitted in plaintext.
	//
	// - Base64: The command content is Base64-encoded.
	//
	// Default value: PlainText. If an invalid value is specified, PlainText is used.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// Specifies whether the modified command contains custom parameters.
	//
	// - If you enable custom parameters or modify the custom parameters specified by `Parameters`, set this parameter to `true`.
	//
	// - If you do not modify the custom parameters specified by `Parameters`, do not set this parameter or set it to `false`.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The modified scheduled execution frequency. This parameter takes effect only when `RepeatMode` is set to `Period`. Three scheduled execution modes are supported: execution at a fixed time interval (based on a rate expression), one-time execution at a specified time, and clock-based scheduled execution (based on a cron expression).
	//
	// - Execution at a fixed time interval: The command is executed at a specified time interval based on a rate expression. The time interval can be specified in seconds (s), minutes (m), hours (h), or days (d). This mode is suitable for scenarios where tasks are executed at fixed time intervals. The format is `rate(<interval value><interval unit>)`. For example, to execute the command every 5 minutes, use `rate(5m)`. The following limits apply to execution at a fixed time interval:
	//
	//     - The specified time interval cannot exceed 7 days or be less than 60 seconds, and must be greater than the timeout period specified when the scheduled task was created.
	//
	//     - The execution interval is based only on the fixed frequency and is independent of the actual time required to execute the task. For example, if the command is executed every 5 minutes and the task takes 2 minutes to complete, the next execution starts 3 minutes after the task is completed.
	//
	//     - The next execution time is calculated based on the task creation time (refer to [CreationTime](https://help.aliyun.com/document_detail/64840.html) returned by `DescribeInvocations`, not the modification time) and the modified execution interval.
	//
	// - One-time execution at a specified time: The command is executed once at a specified time in a specified time zone. The format is `at(yyyy-MM-dd HH:mm:ss <time zone>)`. If no time zone is specified, the UTC time zone is used by default. The time zone supports the following three formats:
	//
	//     - Time zone full name: such as `Asia/Shanghai` (China/Shanghai time) and `America/Los_Angeles` (US/Los Angeles time).
	//
	//     - Time zone offset from Greenwich Mean Time (GMT): such as `GMT+8:00` (East 8th zone) and `GMT-7:00` (West 7th zone). When the GMT format is used, leading zeros are not supported for the hour.
	//
	//     - Time zone abbreviation: Only UTC (Coordinated Universal Time) is supported.
	//
	//   For example, to execute the command once at 13:15:30 on June 6, 2022 in the China/Shanghai time zone, use `at(2022-06-06 13:15:30 Asia/Shanghai)`. To execute the command once at 13:15:30 on June 6, 2022 in the West 7th zone, use `at(2022-06-06 13:15:30 GMT-7:00)`.
	//
	// - Clock-based scheduled execution (based on a cron expression): The command is executed based on a cron expression. The format is `<seconds> <minutes> <hours> <day of month> <month> <day of week> <year (optional)> <time zone>`, which is `<cron expression> <time zone>`. The scheduled task execution time is calculated and executed based on the cron expression in the specified time zone. If no time zone is specified, the internal system time zone of the instance that executes the scheduled task is used by default. For more information about cron expressions, see [Cron expressions](https://help.aliyun.com/document_detail/64769.html). The time zone supports the following three formats:
	//
	//     - Time zone full name: such as `Asia/Shanghai` (China/Shanghai time) and `America/Los_Angeles` (US/Los Angeles time).
	//
	//     - Time zone offset from GMT: such as `GMT+8:00` (East 8th zone) and `GMT-7:00` (West 7th zone). When the GMT format is used, leading zeros are not supported for the hour.
	//
	//     - Time zone abbreviation: Only UTC is supported.
	//
	//   For example, to execute the command once at 10:15 every day in 2022 in the China/Shanghai time zone, use `0 15 10 ? 	- 	- 2022 Asia/Shanghai`. To execute the command every half hour from 10:00 to 11:30 every day in 2022 in the East 8th zone, use `0 0/30 10-11 	- 	- ? 2022 GMT+8:00`. To execute the command every 5 minutes from 14:00 to 14:55 every day in October every two years starting from 2022 in the UTC time zone, use `0 0/5 14 	- 10 ? 2022/2 UTC`.
	//
	//     > The minimum time interval must be greater than or equal to the timeout period specified when the scheduled task was created, and cannot be less than 10 seconds.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The ECS instance ID or managed instance ID to be added to the task.
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The execution ID of the command for the task to be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-hz0jdfwd9f****
	InvokeId     *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key-value pairs of the custom parameters passed in when the command is executed. This parameter is required when the command contains custom parameters.
	//
	// You can specify 0 to 10 custom parameters. Take note of the following items:
	//
	// - A key cannot be an empty string and can be up to 64 characters in length.
	//
	// - A value can be an empty string.
	//
	// - After the custom parameters and the original command content are Base64-encoded, the size of the Base64-encoded command content cannot exceed 24 KB.
	//
	// - The set of custom parameter names must be a subset of the parameter set defined when the command was created. For parameters that are not passed in, you can use empty strings.
	//
	// Default value: empty. This indicates that the key-value pairs of custom parameters are not modified.
	//
	// example:
	//
	// {"name":"Jack", "accessKey":"LTAI*************"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyInvocationAttributeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInvocationAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyInvocationAttributeShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInvocationAttributeShrinkRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *ModifyInvocationAttributeShrinkRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *ModifyInvocationAttributeShrinkRequest) GetEnableParameter() *bool {
	return s.EnableParameter
}

func (s *ModifyInvocationAttributeShrinkRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *ModifyInvocationAttributeShrinkRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *ModifyInvocationAttributeShrinkRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *ModifyInvocationAttributeShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInvocationAttributeShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInvocationAttributeShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *ModifyInvocationAttributeShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInvocationAttributeShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInvocationAttributeShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInvocationAttributeShrinkRequest) SetClientToken(v string) *ModifyInvocationAttributeShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInvocationAttributeShrinkRequest) SetCommandContent(v string) *ModifyInvocationAttributeShrinkRequest {
	s.CommandContent = &v
	return s
}

func (s *ModifyInvocationAttributeShrinkRequest) SetContentEncoding(v string) *ModifyInvocationAttributeShrinkRequest {
	s.ContentEncoding = &v
	return s
}

func (s *ModifyInvocationAttributeShrinkRequest) SetEnableParameter(v bool) *ModifyInvocationAttributeShrinkRequest {
	s.EnableParameter = &v
	return s
}

func (s *ModifyInvocationAttributeShrinkRequest) SetFrequency(v string) *ModifyInvocationAttributeShrinkRequest {
	s.Frequency = &v
	return s
}

func (s *ModifyInvocationAttributeShrinkRequest) SetInstanceId(v []*string) *ModifyInvocationAttributeShrinkRequest {
	s.InstanceId = v
	return s
}

func (s *ModifyInvocationAttributeShrinkRequest) SetInvokeId(v string) *ModifyInvocationAttributeShrinkRequest {
	s.InvokeId = &v
	return s
}

func (s *ModifyInvocationAttributeShrinkRequest) SetOwnerAccount(v string) *ModifyInvocationAttributeShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInvocationAttributeShrinkRequest) SetOwnerId(v int64) *ModifyInvocationAttributeShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInvocationAttributeShrinkRequest) SetParametersShrink(v string) *ModifyInvocationAttributeShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *ModifyInvocationAttributeShrinkRequest) SetRegionId(v string) *ModifyInvocationAttributeShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInvocationAttributeShrinkRequest) SetResourceOwnerAccount(v string) *ModifyInvocationAttributeShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInvocationAttributeShrinkRequest) SetResourceOwnerId(v int64) *ModifyInvocationAttributeShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInvocationAttributeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
