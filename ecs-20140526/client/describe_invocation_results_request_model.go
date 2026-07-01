// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvocationResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *DescribeInvocationResultsRequest
	GetCommandId() *string
	SetContentEncoding(v string) *DescribeInvocationResultsRequest
	GetContentEncoding() *string
	SetIncludeHistory(v bool) *DescribeInvocationResultsRequest
	GetIncludeHistory() *bool
	SetInstanceId(v string) *DescribeInvocationResultsRequest
	GetInstanceId() *string
	SetInvokeId(v string) *DescribeInvocationResultsRequest
	GetInvokeId() *string
	SetInvokeRecordStatus(v string) *DescribeInvocationResultsRequest
	GetInvokeRecordStatus() *string
	SetMaxResults(v int32) *DescribeInvocationResultsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInvocationResultsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeInvocationResultsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInvocationResultsRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeInvocationResultsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeInvocationResultsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeInvocationResultsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeInvocationResultsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeInvocationResultsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInvocationResultsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeInvocationResultsRequestTag) *DescribeInvocationResultsRequest
	GetTag() []*DescribeInvocationResultsRequestTag
}

type DescribeInvocationResultsRequest struct {
	// The command ID.
	//
	// example:
	//
	// c-hz0jdfwcsr****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The encoding method of the `CommandContent` and `Output` fields in the response. Valid values:
	//
	// - PlainText: Returns the original command content and output.
	//
	// - Base64: Returns the Base64-encoded command content and output.
	//
	// Default value: Base64.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// Specifies whether to return the execution history of scheduled commands. Valid values:
	//
	//  - true: Returns the execution results of scheduled commands. When this parameter is set to true, the InvokeId parameter is required and must be the execution ID of a scheduled command (RepeatMode is Period) or a command that runs at each system startup (RepeatMode is EveryReboot).
	//
	//  - false: Does not return the execution history.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IncludeHistory *bool `json:"IncludeHistory,omitempty" xml:"IncludeHistory,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp1i7gg30r52z2em****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The command execution ID. You can call [DescribeInvocations](https://help.aliyun.com/document_detail/64840.html) to query the InvokeId.
	//
	// example:
	//
	// t-hz0jdfwd9f****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The execution status of the command. Valid values:
	//
	// - Running: The command is running.
	//
	//     - Scheduled execution: The execution status remains running until you manually stop the scheduled command.
	//
	//     - One-time execution: The overall execution status is running as long as any command process is running.
	//
	// - Finished: The command execution is complete.
	//
	//     - Scheduled execution: The command process cannot have a status of finished.
	//
	//     - One-time execution: All instances have completed execution, or you manually stopped the command process on some instances and the remaining instances have completed execution.
	//
	// - Success:
	//
	//     - One-time execution: The command execution is complete and the exit code is 0.
	//
	//     - Scheduled execution: The last execution succeeded with an exit code of 0, and the specified execution time has ended.
	//
	// - Failed: The command execution failed.
	//
	//     - Scheduled execution: The command process cannot have a status of failed.
	//
	//     - One-time execution: The command execution failed on all instances.
	//
	// - PartialFailed: The command execution partially failed.
	//
	//     - Scheduled execution: The command process cannot have a status of partially failed.
	//
	//     - One-time execution: The command execution failed on some instances, so the overall execution status is partially failed.
	//
	// - Stopped: The command execution has been stopped.
	//
	// - Stopping: The command execution is being stopped.
	//
	// example:
	//
	// Running
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	// The maximum number of entries per page in a paging query.
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
	// > This parameter is about to go offline. Use NextToken and MaxResults to complete paging query operations.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// > This parameter is about to go offline. Use NextToken and MaxResults to complete paging query operations.
	//
	// example:
	//
	// 1
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID of the command execution. After you specify this parameter, the resource group ID must also be specified when you run the command. This parameter filters the corresponding command execution results.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*DescribeInvocationResultsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInvocationResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationResultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultsRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *DescribeInvocationResultsRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *DescribeInvocationResultsRequest) GetIncludeHistory() *bool {
	return s.IncludeHistory
}

func (s *DescribeInvocationResultsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInvocationResultsRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeInvocationResultsRequest) GetInvokeRecordStatus() *string {
	return s.InvokeRecordStatus
}

func (s *DescribeInvocationResultsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInvocationResultsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInvocationResultsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInvocationResultsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInvocationResultsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeInvocationResultsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInvocationResultsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInvocationResultsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInvocationResultsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInvocationResultsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInvocationResultsRequest) GetTag() []*DescribeInvocationResultsRequestTag {
	return s.Tag
}

func (s *DescribeInvocationResultsRequest) SetCommandId(v string) *DescribeInvocationResultsRequest {
	s.CommandId = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetContentEncoding(v string) *DescribeInvocationResultsRequest {
	s.ContentEncoding = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetIncludeHistory(v bool) *DescribeInvocationResultsRequest {
	s.IncludeHistory = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetInstanceId(v string) *DescribeInvocationResultsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetInvokeId(v string) *DescribeInvocationResultsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetInvokeRecordStatus(v string) *DescribeInvocationResultsRequest {
	s.InvokeRecordStatus = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetMaxResults(v int32) *DescribeInvocationResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetNextToken(v string) *DescribeInvocationResultsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetOwnerAccount(v string) *DescribeInvocationResultsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetOwnerId(v int64) *DescribeInvocationResultsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetPageNumber(v int64) *DescribeInvocationResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetPageSize(v int64) *DescribeInvocationResultsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetRegionId(v string) *DescribeInvocationResultsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetResourceGroupId(v string) *DescribeInvocationResultsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetResourceOwnerAccount(v string) *DescribeInvocationResultsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetResourceOwnerId(v int64) *DescribeInvocationResultsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetTag(v []*DescribeInvocationResultsRequestTag) *DescribeInvocationResultsRequest {
	s.Tag = v
	return s
}

func (s *DescribeInvocationResultsRequest) Validate() error {
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

type DescribeInvocationResultsRequestTag struct {
	// The tag key of the command execution. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// If you use a single tag to filter resources, the number of resources with this tag cannot exceed 1,000. If you use multiple tags to filter resources, the number of resources with all specified tags attached cannot exceed 1,000. If the number of resources exceeds 1,000, call [ListTagResources](https://help.aliyun.com/document_detail/110425.html) to execute the query.
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

func (s DescribeInvocationResultsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationResultsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeInvocationResultsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeInvocationResultsRequestTag) SetKey(v string) *DescribeInvocationResultsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInvocationResultsRequestTag) SetValue(v string) *DescribeInvocationResultsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeInvocationResultsRequestTag) Validate() error {
	return dara.Validate(s)
}
