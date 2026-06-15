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
	// The ID of the command.
	//
	// example:
	//
	// c-hz01272yr52****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The encoding format for the `CommandContent` and `Output` values in the response. Valid values:
	//
	// - PlainText: returns the raw script content and output.
	//
	// - Base64: returns the Base64-encoded script content and output.
	//
	// Default value: Base64.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The description of the command.
	//
	// - If you specify the `Provider` parameter to query public commands, fuzzy search is supported by default.
	//
	// - If you do not specify the `Provider` parameter to query private commands, fuzzy search is supported. You can use an asterisk (\\*) as a wildcard. For example, `test*` returns all commands whose descriptions start with `test`, `*test` returns all commands whose descriptions end with `test`, and `*test*` returns all commands whose descriptions contain `test`.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to return only the latest version of public commands. This parameter does not affect private commands.
	//
	// - true: returns only the latest version of public commands.
	//
	// - false: returns all versions of public commands.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	Latest *bool `json:"Latest,omitempty" xml:"Latest,omitempty"`
	// The maximum number of entries to return per page.
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
	// - If you specify the `Provider` parameter to query public commands, fuzzy search is supported by default.
	//
	// - If you do not specify the `Provider` parameter to query private commands, fuzzy search is supported. You can use an asterisk (\\*) as a wildcard. For example, `command*` returns all commands whose names start with `command`, `*command` returns all commands whose names end with `command`, and `*command*` returns all commands whose names contain `command`.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token for the next page of results. To retrieve the next page, set this parameter to the `NextToken` value from a previous call.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// > This parameter is being deprecated. We recommend using NextToken and MaxResults for pagination instead.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// > This parameter is being deprecated. We recommend using NextToken and MaxResults for pagination instead.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The provider of the public command.
	//
	// - If you omit this parameter, the operation queries your private commands by default.
	//
	// - Set this parameter to `AlibabaCloud` to query all public commands from Alibaba Cloud.
	//
	// - If you set the value to a specific provider, the public commands from that provider are queried. For example:
	//
	//   - If you set `Provider` to `AlibabaCloud.ECS.GuestOS`, the public commands provided by AlibabaCloud.ECS.GuestOS are queried.
	//
	//   - If you set `Provider` to `AlibabaCloud.ECS.GuestOSDiagnose`, the public commands provided by AlibabaCloud.ECS.GuestOSDiagnose are queried.
	//
	// example:
	//
	// AlibabaCloud
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The ID of the region. To view the latest list of regions, call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation.
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
	// The tags used to filter commands.
	Tag []*DescribeCommandsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type of the command. Valid values:
	//
	// - RunBatScript: A Bat script for Windows instances.
	//
	// - RunPowerShellScript: A PowerShell script for Windows instances.
	//
	// - RunShellScript: A Shell script for Linux instances.
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
	// The key of the tag. You can specify up to 20 tags. The tag key cannot be an empty string.
	//
	// A query can return a maximum of 1,000 resources that match the specified tags. If more than 1,000 resources match, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to query all matching resources.
	//
	// The key can be up to 64 characters in length, cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. You can specify up to 20 tags. The tag value can be an empty string.
	//
	// The value can be up to 128 characters in length and cannot contain `http://` or `https://`.
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
