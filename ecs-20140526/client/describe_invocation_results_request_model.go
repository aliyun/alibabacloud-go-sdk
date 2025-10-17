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
	// The ID of the command.
	//
	// example:
	//
	// c-hz0jdfwcsr****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The encoding mode of the `CommandContent` and `Output` values in the response. Valid values:
	//
	// 	- PlainText: returns the original command content and command output.
	//
	// 	- Base64: returns the Base64-encoded command content and command output.
	//
	// Default value: Base64.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// Specifies whether to return the results of historical scheduled executions. Valid values:
	//
	// 	- true: returns the results of historical scheduled executions. If you set this parameter to true, you must set InvokeId to the ID of a task that is run on a schedule (RepeatMode set to Period) or on each system startup (RepeatMode set to EveryReboot).
	//
	// 	- false: does not return the results of historical scheduled executions.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IncludeHistory *bool `json:"IncludeHistory,omitempty" xml:"IncludeHistory,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-bp1i7gg30r52z2em****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the command task. You can call the [DescribeInvocations](https://help.aliyun.com/document_detail/64840.html) operation to query the IDs of all command tasks.
	//
	// example:
	//
	// t-hz0jdfwd9f****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The execution status of the command task. Valid values:
	//
	// 	- Running:
	//
	//     	- Scheduled task: Before you stop the scheduled execution of the command, the execution state is always Running.
	//
	//     	- One-time task: If the command is being run on instances, the execution state is Running.
	//
	// 	- Finished:
	//
	//     	- Scheduled task: The execution state can never be Finished.
	//
	//     	- One-time task: The execution is complete on all instances, or the execution is stopped on some instances and is complete on the other instances.
	//
	// 	- Success:
	//
	//     	- One-time task: The execution is complete, and the exit code is 0.
	//
	//     	- Scheduled task: The last execution is complete, the exit code is 0, and the specified period ends.
	//
	// 	- Failed:
	//
	//     	- Scheduled task: The execution state can never be Failed.
	//
	//     	- One-time task: The execution fails on all instances.
	//
	// 	- PartialFailed:
	//
	//     	- Scheduled task: The execution state can never be PartialFailed.
	//
	//     	- One-time task: The execution fails on some instances.
	//
	// 	- Stopped: The task is stopped.
	//
	// 	- Stopping: The task is being stopped.
	//
	// example:
	//
	// Running
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// >  This parameter will be removed in the future. We recommend that you use NextToken and MaxResults for a paged query.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// >  This parameter will be removed in the future. We recommend that you use NextToken and MaxResults for a paged query.
	//
	// example:
	//
	// 1
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the command. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. After you set this parameter, command execution results in the specified resource group are queried.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the command task.
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
	// The key of tag N of the command task. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// If a single tag is specified to query resources, up to 1,000 resources that have this tag added can be displayed in the response. If multiple tags are specified to query resources, up to 1,000 resources that have all these tags added can be displayed in the response. To query more than 1,000 resources that have specified tags added, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// The tag key can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the command task. Valid values of N: 1 to 20. The tag value can be an empty string.
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
