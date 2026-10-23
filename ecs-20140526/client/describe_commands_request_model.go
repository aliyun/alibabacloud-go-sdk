// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommandsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *DescribeCommandsRequest
	GetCommandId() *string
	SetContentEncoding(v string) *DescribeCommandsRequest
	GetContentEncoding() *string
	SetDescription(v string) *DescribeCommandsRequest
	GetDescription() *string
	SetLatest(v bool) *DescribeCommandsRequest
	GetLatest() *bool
	SetMaxResults(v int32) *DescribeCommandsRequest
	GetMaxResults() *int32
	SetName(v string) *DescribeCommandsRequest
	GetName() *string
	SetNextToken(v string) *DescribeCommandsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeCommandsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCommandsRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeCommandsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCommandsRequest
	GetPageSize() *int64
	SetProvider(v string) *DescribeCommandsRequest
	GetProvider() *string
	SetRegionId(v string) *DescribeCommandsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeCommandsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeCommandsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCommandsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeCommandsRequestTag) *DescribeCommandsRequest
	GetTag() []*DescribeCommandsRequestTag
	SetType(v string) *DescribeCommandsRequest
	GetType() *string
}

type DescribeCommandsRequest struct {
	// The command ID.
	//
	// example:
	//
	// c-hz01272yr52****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The encoding mode of the CommandContent and Output fields in the response. Valid values:
	//
	// - PlainText: Returns the original script content and output information.
	//
	// - Base64: Returns the Base64-encoded script content and output information.
	//
	// Default value: Base64.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The description of the command.
	//
	// - When you specify the Provider parameter parameter query public commands, fuzzy search is supported by default.
	//
	// - When you do not specify the Provider parameter parameter query private commands, fuzzy search is supported. For example, if you enter test*, the system searches for all commands whose descriptions start with test. If you enter *test, the system searches for all commands whose descriptions end with test. If you enter *test*, the system searches for all commands whose descriptions contain test.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to query only the latest version of public commands if the query result includes public commands. This parameter does not affect the query of private commands. Valid values:
	//
	// - true: Queries only the latest version of public commands.
	//
	// - false: Queries all versions of public commands.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	Latest *bool `json:"Latest,omitempty" xml:"Latest,omitempty"`
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
	// The name of the command.
	//
	// - When you specify the Provider parameter parameter query public commands, fuzzy search is supported by default.
	//
	// - When you do not specify the Provider parameter parameter query private commands, fuzzy search is supported. For example, if you enter command*, the system searches for all commands whose names start with command. If you enter *command, the system searches for all commands whose names end with command. If you enter *command*, the system searches for all commands whose names contain command.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token. Set the value to the NextToken value returned from the previous API call.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// > This parameter will be deprecated soon. Use the NextToken and MaxResults parameters for paging.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// > This parameter will be deprecated soon. Use the NextToken and MaxResults parameters for paging.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The provider of the public command. Valid values:
	//
	// - If you do not specify this parameter, the system queries all Cloud Assistant commands that you manually created by default.
	//
	// - If you set this parameter to AlibabaCloud, the system queries all public commands provided by Alibaba Cloud.
	//
	// - If you set this parameter to a specific public command provider, the system queries all public commands provided by the provider. Examples:
	//
	//     - If you set this parameter to AlibabaCloud.ECS.GuestOS, the system queries the public commands provided by AlibabaCloud.ECS.GuestOS.
	//
	//     - If you set this parameter to AlibabaCloud.ECS.GuestOSDiagnose, the system queries the public commands provided by AlibabaCloud.ECS.GuestOSDiagnose.
	//
	// example:
	//
	// AlibabaCloud
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the command belongs.
	//
	// example:
	//
	// rg-123******
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of tags.
	Tag []*DescribeCommandsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type of the command. Valid values:
	//
	// - RunBatScript: a Bat script that runs on a Windows instance.
	//
	// - RunPowerShellScript: a PowerShell script that runs on a Windows instance.
	//
	// - RunShellScript: a shell script that runs on a Linux instance.
	//
	// example:
	//
	// RunShellScript
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCommandsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCommandsRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *DescribeCommandsRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *DescribeCommandsRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeCommandsRequest) GetLatest() *bool {
	return s.Latest
}

func (s *DescribeCommandsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCommandsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeCommandsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCommandsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCommandsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCommandsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCommandsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCommandsRequest) GetProvider() *string {
	return s.Provider
}

func (s *DescribeCommandsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCommandsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCommandsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCommandsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCommandsRequest) GetTag() []*DescribeCommandsRequestTag {
	return s.Tag
}

func (s *DescribeCommandsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeCommandsRequest) SetCommandId(v string) *DescribeCommandsRequest {
	s.CommandId = &v
	return s
}

func (s *DescribeCommandsRequest) SetContentEncoding(v string) *DescribeCommandsRequest {
	s.ContentEncoding = &v
	return s
}

func (s *DescribeCommandsRequest) SetDescription(v string) *DescribeCommandsRequest {
	s.Description = &v
	return s
}

func (s *DescribeCommandsRequest) SetLatest(v bool) *DescribeCommandsRequest {
	s.Latest = &v
	return s
}

func (s *DescribeCommandsRequest) SetMaxResults(v int32) *DescribeCommandsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCommandsRequest) SetName(v string) *DescribeCommandsRequest {
	s.Name = &v
	return s
}

func (s *DescribeCommandsRequest) SetNextToken(v string) *DescribeCommandsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCommandsRequest) SetOwnerAccount(v string) *DescribeCommandsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCommandsRequest) SetOwnerId(v int64) *DescribeCommandsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCommandsRequest) SetPageNumber(v int64) *DescribeCommandsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCommandsRequest) SetPageSize(v int64) *DescribeCommandsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCommandsRequest) SetProvider(v string) *DescribeCommandsRequest {
	s.Provider = &v
	return s
}

func (s *DescribeCommandsRequest) SetRegionId(v string) *DescribeCommandsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCommandsRequest) SetResourceGroupId(v string) *DescribeCommandsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCommandsRequest) SetResourceOwnerAccount(v string) *DescribeCommandsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCommandsRequest) SetResourceOwnerId(v int64) *DescribeCommandsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCommandsRequest) SetTag(v []*DescribeCommandsRequestTag) *DescribeCommandsRequest {
	s.Tag = v
	return s
}

func (s *DescribeCommandsRequest) SetType(v string) *DescribeCommandsRequest {
	s.Type = &v
	return s
}

func (s *DescribeCommandsRequest) Validate() error {
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

type DescribeCommandsRequestTag struct {
	// The tag key of the command. Valid values of N: 1 to 20. The value cannot be an empty string.
	//
	// When you use a single tag to filter resources, the resource count for that tag cannot exceed 1,000. When you use multiple tags to filter resources, the resource count for resources attached to all the specified tags cannot exceed 1,000. If the resource count exceeds 1,000, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to query the resources.
	//
	// The tag key can be up to 64 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the command. Valid values of N: 1 to 20. The value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCommandsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeCommandsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeCommandsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeCommandsRequestTag) SetKey(v string) *DescribeCommandsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeCommandsRequestTag) SetValue(v string) *DescribeCommandsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeCommandsRequestTag) Validate() error {
	return dara.Validate(s)
}
