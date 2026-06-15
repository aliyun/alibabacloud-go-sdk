// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvocationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *DescribeInvocationsRequest
	GetCommandId() *string
	SetCommandName(v string) *DescribeInvocationsRequest
	GetCommandName() *string
	SetCommandType(v string) *DescribeInvocationsRequest
	GetCommandType() *string
	SetContentEncoding(v string) *DescribeInvocationsRequest
	GetContentEncoding() *string
	SetIncludeOutput(v bool) *DescribeInvocationsRequest
	GetIncludeOutput() *bool
	SetInstanceId(v string) *DescribeInvocationsRequest
	GetInstanceId() *string
	SetInvokeId(v string) *DescribeInvocationsRequest
	GetInvokeId() *string
	SetInvokeStatus(v string) *DescribeInvocationsRequest
	GetInvokeStatus() *string
	SetMaxResults(v int32) *DescribeInvocationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInvocationsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeInvocationsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInvocationsRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeInvocationsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeInvocationsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeInvocationsRequest
	GetRegionId() *string
	SetRepeatMode(v string) *DescribeInvocationsRequest
	GetRepeatMode() *string
	SetResourceGroupId(v string) *DescribeInvocationsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeInvocationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInvocationsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeInvocationsRequestTag) *DescribeInvocationsRequest
	GetTag() []*DescribeInvocationsRequestTag
	SetTimed(v bool) *DescribeInvocationsRequest
	GetTimed() *bool
}

type DescribeInvocationsRequest struct {
	// The command ID. You can call the [DescribeCommands](https://help.aliyun.com/document_detail/64843.html) API to query all available CommandId values.
	//
	// example:
	//
	// c-hz0jdfwcsr****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The command name. If the `InstanceId` parameter is also specified, this parameter does not take effect.
	//
	// example:
	//
	// CommandTestName
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	// The command type. Valid values:
	//
	// - RunBatScript: The command is a Bat script that runs on a Windows instance.
	//
	// - RunPowerShellScript: The command is a PowerShell script that runs on a Windows instance.
	//
	// - RunShellScript: The command is a Shell script that runs on a Linux instance.
	//
	// example:
	//
	// RunShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The codec for the `CommandContent` and `Output` fields in the returned data. Valid values:
	//
	// - PlainText: Returns the original command content and output information.
	//
	// - Base64: Returns Base64-encoded command content and output information.
	//
	// Default value: Base64.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// Indicates whether to return the command execution output in the results.
	//
	// - true: Returns the output. In this case, you must specify at least one of the `InvokeId` or `InstanceId` parameters.
	//
	// - false: Does not return the output.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IncludeOutput *bool `json:"IncludeOutput,omitempty" xml:"IncludeOutput,omitempty"`
	// The instance ID. If you specify this parameter, all command execution records for this instance will be queried.
	//
	// example:
	//
	// i-bp1i7gg30r52z2em****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The command execution ID.
	//
	// example:
	//
	// t-hz0jdfwd9f****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The overall execution status of the command. The overall status depends on the combined execution statuses of one or more instances involved in the command execution. Valid values:
	//
	// - Running:
	//
	//   - Periodic execution: The status remains Running until the periodic execution is manually stopped.
	//
	//   - One-time execution: The overall status is Running as long as any instance has a running command process.
	//
	// - Finished:
	//
	//   - Periodic execution: The command process cannot reach a Finished state.
	//
	//   - One-time execution: All instances have completed execution, or some instances were manually stopped while the rest completed execution.
	//
	// - Success: All instances have a command execution status of either Stopped or Success, and at least one instance has a status of Success. Specifically:
	//
	//   - For immediate jobs: The command completed successfully with an exit code of 0.
	//
	//   - For scheduled jobs: The most recent execution succeeded with an exit code of 0, and all scheduled times have been completed.
	//
	// - Failed:
	//
	//   - Periodic execution: The command process cannot reach a Failed state.
	//
	//   - One-time execution: All instances failed execution.
	//
	// - Stopped: The command was stopped.
	//
	// - Stopping: The command is being stopped.
	//
	// - PartialFailed: Partial failure. This value does not take effect if the `InstanceId` parameter is also specified.
	//
	// - Pending: The system is validating or sending the command. If at least one instance has a Pending execution status, the overall status is Pending.
	//
	// - Scheduled: The scheduled command has been sent and is waiting to run. If at least one instance has a Scheduled execution status, the overall status is Scheduled.
	//
	// example:
	//
	// Finished
	InvokeStatus *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	// The maximum number of entries per page for paged queries.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query credential (Token). Set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// > This parameter will be unpublished soon. We recommend that you use NextToken and MaxResults to perform paged queries.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// > This parameter will be unpublished soon. We recommend that you use NextToken and MaxResults to perform paged query operations.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The Region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to view the latest Alibaba Cloud region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The execution mode of the command. This parameter does not take effect if the `InstanceId` parameter is also specified. Valid values:
	//
	// - Once: Executes the command immediately.
	//
	// - Period: Executes the command periodically.
	//
	// - NextRebootOnly: Automatically executes the command the next time the instance starts.
	//
	// - EveryReboot: Automatically executes the command every time the instance starts.
	//
	// Default value: empty, which indicates that all modes are queried.
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// The resource group ID of the command execution. After you specify this parameter, you must also specify ResourceGroupId when executing the command. This enables filtering to retrieve the corresponding command execution results.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag list.
	Tag []*DescribeInvocationsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Indicates whether the queried command will be automatically executed in the future. Valid values:
	//
	// - true: The command is queried when the `RepeatMode` parameter is set to `Period`, `NextRebootOnly`, or `EveryReboot` during a call to `RunCommand` or `InvokeCommand`.
	//
	// - false: The command is queried under either of the following conditions:
	//
	//   - The `RepeatMode` parameter is set to `Once` during a call to `RunCommand` or `InvokeCommand`.
	//
	//   - The command has been canceled, stopped, or has finished execution.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	Timed *bool `json:"Timed,omitempty" xml:"Timed,omitempty"`
}

func (s DescribeInvocationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *DescribeInvocationsRequest) GetCommandName() *string {
	return s.CommandName
}

func (s *DescribeInvocationsRequest) GetCommandType() *string {
	return s.CommandType
}

func (s *DescribeInvocationsRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *DescribeInvocationsRequest) GetIncludeOutput() *bool {
	return s.IncludeOutput
}

func (s *DescribeInvocationsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInvocationsRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeInvocationsRequest) GetInvokeStatus() *string {
	return s.InvokeStatus
}

func (s *DescribeInvocationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInvocationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInvocationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInvocationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInvocationsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeInvocationsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInvocationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInvocationsRequest) GetRepeatMode() *string {
	return s.RepeatMode
}

func (s *DescribeInvocationsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInvocationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInvocationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInvocationsRequest) GetTag() []*DescribeInvocationsRequestTag {
	return s.Tag
}

func (s *DescribeInvocationsRequest) GetTimed() *bool {
	return s.Timed
}

func (s *DescribeInvocationsRequest) SetCommandId(v string) *DescribeInvocationsRequest {
	s.CommandId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetCommandName(v string) *DescribeInvocationsRequest {
	s.CommandName = &v
	return s
}

func (s *DescribeInvocationsRequest) SetCommandType(v string) *DescribeInvocationsRequest {
	s.CommandType = &v
	return s
}

func (s *DescribeInvocationsRequest) SetContentEncoding(v string) *DescribeInvocationsRequest {
	s.ContentEncoding = &v
	return s
}

func (s *DescribeInvocationsRequest) SetIncludeOutput(v bool) *DescribeInvocationsRequest {
	s.IncludeOutput = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInstanceId(v string) *DescribeInvocationsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeId(v string) *DescribeInvocationsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeStatus(v string) *DescribeInvocationsRequest {
	s.InvokeStatus = &v
	return s
}

func (s *DescribeInvocationsRequest) SetMaxResults(v int32) *DescribeInvocationsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInvocationsRequest) SetNextToken(v string) *DescribeInvocationsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInvocationsRequest) SetOwnerAccount(v string) *DescribeInvocationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInvocationsRequest) SetOwnerId(v int64) *DescribeInvocationsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetPageNumber(v int64) *DescribeInvocationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInvocationsRequest) SetPageSize(v int64) *DescribeInvocationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInvocationsRequest) SetRegionId(v string) *DescribeInvocationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetRepeatMode(v string) *DescribeInvocationsRequest {
	s.RepeatMode = &v
	return s
}

func (s *DescribeInvocationsRequest) SetResourceGroupId(v string) *DescribeInvocationsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetResourceOwnerAccount(v string) *DescribeInvocationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInvocationsRequest) SetResourceOwnerId(v int64) *DescribeInvocationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetTag(v []*DescribeInvocationsRequestTag) *DescribeInvocationsRequest {
	s.Tag = v
	return s
}

func (s *DescribeInvocationsRequest) SetTimed(v bool) *DescribeInvocationsRequest {
	s.Timed = &v
	return s
}

func (s *DescribeInvocationsRequest) Validate() error {
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

type DescribeInvocationsRequestTag struct {
	// The tag key for command execution. Valid values for N are 1 to 20. If this value is specified, it cannot be an empty string.
	//
	// When you use one tag to filter resources, the number of resources under this tag cannot exceed 1,000. When you use multiple tags to filter resources, the number of resources bound to all specified tags simultaneously cannot exceed 1,000. If the number of resources exceeds 1,000, you must use the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) API to query them.
	//
	// The key can contain up to 64 characters, must not start with `aliyun` or `acs:`, and must not contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value for command execution. Valid values for N are 1 to 20. This value can be an empty string.
	//
	// The value can contain up to 128 characters and must not contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInvocationsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeInvocationsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeInvocationsRequestTag) SetKey(v string) *DescribeInvocationsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInvocationsRequestTag) SetValue(v string) *DescribeInvocationsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeInvocationsRequestTag) Validate() error {
	return dara.Validate(s)
}
