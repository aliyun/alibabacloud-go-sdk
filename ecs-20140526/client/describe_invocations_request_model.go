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
	// The command ID. You can call [DescribeCommands](https://help.aliyun.com/document_detail/64843.html) to query all available command IDs.
	//
	// example:
	//
	// c-hz0jdfwcsr****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The command name. This parameter does not take effect if you also specify InstanceId.
	//
	// example:
	//
	// CommandTestName
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	// The command type. Valid values:
	//
	// - RunBatScript: Bat script that runs on Windows instances.
	//
	// - RunPowerShellScript: PowerShell script that runs on Windows instances.
	//
	// - RunShellScript: Shell script that runs on Linux instances.
	//
	// example:
	//
	// RunShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The encoding mode of the CommandContent and Output fields in the response. Valid values:
	//
	// - PlainText: Returns the original command content and output.
	//
	// - Base64: Returns Base64-encoded command content and output.
	//
	// Default value: Base64.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// Specifies whether to return the command output in the results.
	//
	// - true: The command output is returned. You must specify at least InvokeId or InstanceId.
	//
	// - false: The command output is not returned.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IncludeOutput *bool `json:"IncludeOutput,omitempty" xml:"IncludeOutput,omitempty"`
	// The instance ID. If you specify this parameter, all command execution records for the instance are queried.
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
	// The overall execution status of the command. The overall execution status is determined by the combined execution status across one or more instances. Valid values:
	//
	//
	//
	// - Running:
	//
	//     - Scheduled execution: The execution status remains Running until you manually stop the scheduled command.
	//
	//     - One-time execution: The overall status is Running if any command process is in progress.
	//
	// - Finished:
	//
	//     - Scheduled execution: The status can never be Finished.
	//
	//     - One-time execution: All instances have completed execution, or you manually stopped the command process on some instances while the remaining instances completed execution.
	//
	// - Success: The command execution status on each instance is Stopped or Success, and at least one instance has a status of Success. The overall status is then Success.
	//
	//     - Immediate task: The command execution is complete and the exit code is 0.
	//
	//     - Scheduled task: The most recent execution succeeded with an exit code of 0, and all specified execution times have elapsed.
	//
	// - Failed:
	//
	//     - Scheduled execution: The status can never be Failed.
	//
	//     - One-time execution: All instances failed to run the command.
	//
	// - Stopped: The command has been stopped.
	//
	// - Stopping: The command is being stopped.
	//
	// - PartialFailed: Some instances succeeded while others failed. This value does not take effect if you also specify InstanceId.
	//
	// - Pending: The system is validating or sending the command. The overall status is Pending if at least one instance has a status of Pending.
	//
	// - Scheduled: The scheduled command has been sent and is waiting to run. The overall status is Scheduled if at least one instance has a status of Scheduled.
	//
	// example:
	//
	// Finished
	InvokeStatus *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	// The maximum number of entries per page for paging.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// > This parameter is about to be deprecated. Use NextToken and MaxResults to complete paging operations.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// > This parameter is about to be deprecated. Use NextToken and MaxResults to complete paging operations.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The execution mode of the command. This parameter does not take effect if you also specify InstanceId. Valid values:
	//
	// - Once: immediately runs the command.
	//
	// - Period: runs the command on a schedule.
	//
	// - NextRebootOnly: automatically runs the command the next time the instance starts.
	//
	// - EveryReboot: automatically runs the command every time the instance starts.
	//
	// Default value: empty, which indicates that all execution modes are queried.
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// The resource group ID for command execution. After you specify this parameter, the resource group ID must also be specified when you run the command. This parameter filters the corresponding command execution results.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*DescribeInvocationsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether the queried commands will be automatically run in the future. Valid values:
	//
	// - true: queries commands whose RepeatMode parameter is set to Period, NextRebootOnly, or EveryReboot when RunCommand or InvokeCommand is called.
	//
	// - false: queries commands in the following states:
	//
	//     - Commands whose RepeatMode parameter is set to Once when RunCommand or InvokeCommand is called.
	//
	//     - Commands that have been canceled, stopped, or completed.
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
	// The tag key of the command execution. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// If you use a single tag to filter resources, the resource count with the specified tag cannot exceed 1,000. If you use multiple tags to filter resources, the resource count with all specified tags attached cannot exceed 1,000. If the resource count exceeds 1,000, call [ListTagResources](https://help.aliyun.com/document_detail/110425.html) to execute the query.
	//
	// The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the command execution. Valid values of N: 1 to 20. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
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
