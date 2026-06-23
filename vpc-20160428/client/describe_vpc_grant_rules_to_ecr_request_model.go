// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcGrantRulesToEcrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcrInstanceId(v string) *DescribeVpcGrantRulesToEcrRequest
	GetEcrInstanceId() *string
	SetEcrOwnerId(v int64) *DescribeVpcGrantRulesToEcrRequest
	GetEcrOwnerId() *int64
	SetInstanceId(v string) *DescribeVpcGrantRulesToEcrRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeVpcGrantRulesToEcrRequest
	GetInstanceType() *string
	SetMaxResults(v int32) *DescribeVpcGrantRulesToEcrRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeVpcGrantRulesToEcrRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeVpcGrantRulesToEcrRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVpcGrantRulesToEcrRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeVpcGrantRulesToEcrRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeVpcGrantRulesToEcrRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeVpcGrantRulesToEcrRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVpcGrantRulesToEcrRequest
	GetResourceOwnerId() *int64
	SetTags(v []*DescribeVpcGrantRulesToEcrRequestTags) *DescribeVpcGrantRulesToEcrRequest
	GetTags() []*DescribeVpcGrantRulesToEcrRequestTags
}

type DescribeVpcGrantRulesToEcrRequest struct {
	// The ID of the Express Connect Router.
	//
	// example:
	//
	// ecr-ncxadcujadncsa****
	EcrInstanceId *string `json:"EcrInstanceId,omitempty" xml:"EcrInstanceId,omitempty"`
	// The ID of the Alibaba Cloud account (main account) that owns the Express Connect Router.
	//
	// > This parameter is required when querying a cross-account network instance.
	//
	// example:
	//
	// 192732132151****
	EcrOwnerId *int64 `json:"EcrOwnerId,omitempty" xml:"EcrOwnerId,omitempty"`
	// The ID of the network instance.
	//
	// example:
	//
	// vpc-wz9ek66wd7tl5xqpy****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of instance whose authorization rules you want to query. Valid values:
	//
	// - **VBR**: Set the value to **VBR*	- to query the Virtual Private Cloud (VPC) instances authorized to connect to the specified virtual border router (VBR).
	//
	// - **VPC**: Set the value to **VPC*	- to query the VBRs to which the specified VPC has granted authorization.
	//
	// example:
	//
	// VPC
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of entries to return per page. Valid values: **1*	- to **100**. Default value: **100**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results. Valid values:
	//
	// - Omit this parameter for the first request.
	//
	// - For subsequent requests, set this to the **NextToken*	- value from the previous response.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the network instance is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the network instance belongs.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags. You can specify up to 20 tags.
	Tags []*DescribeVpcGrantRulesToEcrRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeVpcGrantRulesToEcrRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcGrantRulesToEcrRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcGrantRulesToEcrRequest) GetEcrInstanceId() *string {
	return s.EcrInstanceId
}

func (s *DescribeVpcGrantRulesToEcrRequest) GetEcrOwnerId() *int64 {
	return s.EcrOwnerId
}

func (s *DescribeVpcGrantRulesToEcrRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVpcGrantRulesToEcrRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeVpcGrantRulesToEcrRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeVpcGrantRulesToEcrRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeVpcGrantRulesToEcrRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVpcGrantRulesToEcrRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpcGrantRulesToEcrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpcGrantRulesToEcrRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVpcGrantRulesToEcrRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVpcGrantRulesToEcrRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVpcGrantRulesToEcrRequest) GetTags() []*DescribeVpcGrantRulesToEcrRequestTags {
	return s.Tags
}

func (s *DescribeVpcGrantRulesToEcrRequest) SetEcrInstanceId(v string) *DescribeVpcGrantRulesToEcrRequest {
	s.EcrInstanceId = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrRequest) SetEcrOwnerId(v int64) *DescribeVpcGrantRulesToEcrRequest {
	s.EcrOwnerId = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrRequest) SetInstanceId(v string) *DescribeVpcGrantRulesToEcrRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrRequest) SetInstanceType(v string) *DescribeVpcGrantRulesToEcrRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrRequest) SetMaxResults(v int32) *DescribeVpcGrantRulesToEcrRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrRequest) SetNextToken(v string) *DescribeVpcGrantRulesToEcrRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrRequest) SetOwnerAccount(v string) *DescribeVpcGrantRulesToEcrRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrRequest) SetOwnerId(v int64) *DescribeVpcGrantRulesToEcrRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrRequest) SetRegionId(v string) *DescribeVpcGrantRulesToEcrRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrRequest) SetResourceGroupId(v string) *DescribeVpcGrantRulesToEcrRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrRequest) SetResourceOwnerAccount(v string) *DescribeVpcGrantRulesToEcrRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrRequest) SetResourceOwnerId(v int64) *DescribeVpcGrantRulesToEcrRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrRequest) SetTags(v []*DescribeVpcGrantRulesToEcrRequestTags) *DescribeVpcGrantRulesToEcrRequest {
	s.Tags = v
	return s
}

func (s *DescribeVpcGrantRulesToEcrRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcGrantRulesToEcrRequestTags struct {
	// The tag key. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVpcGrantRulesToEcrRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcGrantRulesToEcrRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeVpcGrantRulesToEcrRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeVpcGrantRulesToEcrRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeVpcGrantRulesToEcrRequestTags) SetKey(v string) *DescribeVpcGrantRulesToEcrRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrRequestTags) SetValue(v string) *DescribeVpcGrantRulesToEcrRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrRequestTags) Validate() error {
	return dara.Validate(s)
}
