// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVscsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeVscsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *DescribeVscsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeVscsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeVscsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVscsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeVscsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeVscsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeVscsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVscsRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *DescribeVscsRequest
	GetStatus() *string
	SetTag(v []*DescribeVscsRequestTag) *DescribeVscsRequest
	GetTag() []*DescribeVscsRequestTag
	SetVscIds(v []*string) *DescribeVscsRequest
	GetVscIds() []*string
}

type DescribeVscsRequest struct {
	// The instance ID of the cloud disk or local disk that is attached.
	//
	// example:
	//
	// i-bp67acfmxazb4ph****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return. Valid values: 10 to 500. Default value: If the value is not specified or is less than 10, the default value is 10. If the value is greater than 500, the default value is 500.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token returned by this call.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs. When you use this parameter to filter resources, the resource count cannot exceed 1000. Filtering by the default resource group is not supported.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the VSCs to query. If you do not specify this parameter, VSCs in all states are returned. Valid values:
	//
	// - In_use: in use.
	//
	// - Attaching: being attached.
	//
	// - Detaching: being detached.
	//
	// - AttachFailed: failed to attach.
	//
	// - DetachFailed: failed to detach.
	//
	// - All (default): all states.
	//
	// Default value: All.
	//
	// example:
	//
	// Successful
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tag []*DescribeVscsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The IDs of the VSCs to query.
	//
	// We recommend that you specify no more than 100 VSC IDs.
	VscIds []*string `json:"VscIds,omitempty" xml:"VscIds,omitempty" type:"Repeated"`
}

func (s DescribeVscsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVscsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVscsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeVscsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeVscsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVscsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVscsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVscsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVscsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVscsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVscsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeVscsRequest) GetTag() []*DescribeVscsRequestTag {
	return s.Tag
}

func (s *DescribeVscsRequest) GetVscIds() []*string {
	return s.VscIds
}

func (s *DescribeVscsRequest) SetInstanceId(v string) *DescribeVscsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeVscsRequest) SetMaxResults(v int32) *DescribeVscsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeVscsRequest) SetNextToken(v string) *DescribeVscsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeVscsRequest) SetOwnerAccount(v string) *DescribeVscsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVscsRequest) SetOwnerId(v int64) *DescribeVscsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVscsRequest) SetRegionId(v string) *DescribeVscsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVscsRequest) SetResourceGroupId(v string) *DescribeVscsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVscsRequest) SetResourceOwnerAccount(v string) *DescribeVscsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVscsRequest) SetResourceOwnerId(v int64) *DescribeVscsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVscsRequest) SetStatus(v string) *DescribeVscsRequest {
	s.Status = &v
	return s
}

func (s *DescribeVscsRequest) SetTag(v []*DescribeVscsRequestTag) *DescribeVscsRequest {
	s.Tag = v
	return s
}

func (s *DescribeVscsRequest) SetVscIds(v []*string) *DescribeVscsRequest {
	s.VscIds = v
	return s
}

func (s *DescribeVscsRequest) Validate() error {
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

type DescribeVscsRequestTag struct {
	// The tag key of the command. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// If you use a single tag to filter resources, the resource count with the specified tag cannot exceed 1000. If you use multiple tags to filter resources, the resource count of resources that are attached to all specified tags cannot exceed 1000. If the resource count exceeds 1000, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to query the resources.
	//
	// The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// Environment
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVscsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeVscsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVscsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVscsRequestTag) SetKey(v string) *DescribeVscsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeVscsRequestTag) SetValue(v string) *DescribeVscsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeVscsRequestTag) Validate() error {
	return dara.Validate(s)
}
