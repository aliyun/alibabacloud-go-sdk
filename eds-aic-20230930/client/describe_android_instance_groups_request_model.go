// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAndroidInstanceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *DescribeAndroidInstanceGroupsRequest
	GetBizRegionId() *string
	SetChargeType(v string) *DescribeAndroidInstanceGroupsRequest
	GetChargeType() *string
	SetInstanceGroupIds(v []*string) *DescribeAndroidInstanceGroupsRequest
	GetInstanceGroupIds() []*string
	SetInstanceGroupName(v string) *DescribeAndroidInstanceGroupsRequest
	GetInstanceGroupName() *string
	SetInstanceVersion(v string) *DescribeAndroidInstanceGroupsRequest
	GetInstanceVersion() *string
	SetKeyPairId(v string) *DescribeAndroidInstanceGroupsRequest
	GetKeyPairId() *string
	SetMaxResults(v int32) *DescribeAndroidInstanceGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAndroidInstanceGroupsRequest
	GetNextToken() *string
	SetPolicyGroupId(v string) *DescribeAndroidInstanceGroupsRequest
	GetPolicyGroupId() *string
	SetSaleMode(v string) *DescribeAndroidInstanceGroupsRequest
	GetSaleMode() *string
	SetStatus(v string) *DescribeAndroidInstanceGroupsRequest
	GetStatus() *string
	SetTags(v []*DescribeAndroidInstanceGroupsRequestTags) *DescribeAndroidInstanceGroupsRequest
	GetTags() []*DescribeAndroidInstanceGroupsRequestTags
}

type DescribeAndroidInstanceGroupsRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing type.
	//
	// [_single.params.ChargeType.enum. PrePaid]Subscription.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The list of instance group IDs.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// The instance group name. Fuzzy match is supported.
	//
	// example:
	//
	// Cloud phone
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" xml:"InstanceGroupName,omitempty"`
	InstanceVersion   *string `json:"InstanceVersion,omitempty" xml:"InstanceVersion,omitempty"`
	// The ID of the key pair.
	//
	// example:
	//
	// kp-5htf0ymsrnb7q****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The maximum number of entries per page for a paged query. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that indicates the position from which the current read operation starts. Leave this parameter empty to read from the beginning.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uONHqPtDLM2U8s****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// pg-1b77w6xrqfubi****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The purchase mode of the cloud phone.
	//
	// example:
	//
	// standard
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	// The instance group status.
	//
	// example:
	//
	// CREATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the instance group. You can bind up to 20 tags to each instance.
	Tags []*DescribeAndroidInstanceGroupsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeAndroidInstanceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstanceGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstanceGroupsRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *DescribeAndroidInstanceGroupsRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeAndroidInstanceGroupsRequest) GetInstanceGroupIds() []*string {
	return s.InstanceGroupIds
}

func (s *DescribeAndroidInstanceGroupsRequest) GetInstanceGroupName() *string {
	return s.InstanceGroupName
}

func (s *DescribeAndroidInstanceGroupsRequest) GetInstanceVersion() *string {
	return s.InstanceVersion
}

func (s *DescribeAndroidInstanceGroupsRequest) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *DescribeAndroidInstanceGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAndroidInstanceGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAndroidInstanceGroupsRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *DescribeAndroidInstanceGroupsRequest) GetSaleMode() *string {
	return s.SaleMode
}

func (s *DescribeAndroidInstanceGroupsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeAndroidInstanceGroupsRequest) GetTags() []*DescribeAndroidInstanceGroupsRequestTags {
	return s.Tags
}

func (s *DescribeAndroidInstanceGroupsRequest) SetBizRegionId(v string) *DescribeAndroidInstanceGroupsRequest {
	s.BizRegionId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetChargeType(v string) *DescribeAndroidInstanceGroupsRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetInstanceGroupIds(v []*string) *DescribeAndroidInstanceGroupsRequest {
	s.InstanceGroupIds = v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetInstanceGroupName(v string) *DescribeAndroidInstanceGroupsRequest {
	s.InstanceGroupName = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetInstanceVersion(v string) *DescribeAndroidInstanceGroupsRequest {
	s.InstanceVersion = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetKeyPairId(v string) *DescribeAndroidInstanceGroupsRequest {
	s.KeyPairId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetMaxResults(v int32) *DescribeAndroidInstanceGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetNextToken(v string) *DescribeAndroidInstanceGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetPolicyGroupId(v string) *DescribeAndroidInstanceGroupsRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetSaleMode(v string) *DescribeAndroidInstanceGroupsRequest {
	s.SaleMode = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetStatus(v string) *DescribeAndroidInstanceGroupsRequest {
	s.Status = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetTags(v []*DescribeAndroidInstanceGroupsRequestTags) *DescribeAndroidInstanceGroupsRequest {
	s.Tags = v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) Validate() error {
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

type DescribeAndroidInstanceGroupsRequestTags struct {
	// The tag key. You can specify 1 to 20 tag keys.
	//
	// 	Notice: The tag key can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://..
	//
	// example:
	//
	// phone
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// 	Notice: The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`..
	//
	// example:
	//
	// 2025
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAndroidInstanceGroupsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstanceGroupsRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstanceGroupsRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeAndroidInstanceGroupsRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeAndroidInstanceGroupsRequestTags) SetKey(v string) *DescribeAndroidInstanceGroupsRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequestTags) SetValue(v string) *DescribeAndroidInstanceGroupsRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequestTags) Validate() error {
	return dara.Validate(s)
}
