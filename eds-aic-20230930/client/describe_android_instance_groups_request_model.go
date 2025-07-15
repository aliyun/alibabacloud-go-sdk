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
}

type DescribeAndroidInstanceGroupsRequest struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// 	- PrePaid: subscription
	//
	// 	- PostPaid: pay-as-you-go
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The IDs of the instance groups.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// The name of the instance group. Instance groups support fuzzy search by name.
	//
	// example:
	//
	// defaultInstanceGroup
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" xml:"InstanceGroupName,omitempty"`
	// The ID of the key pair.
	//
	// example:
	//
	// kp-5htf0ymsrnb7q****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The maximum number of entries per page. Value range: 0 to 100. Default value: 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uONHqPtDLM2U8s****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// pg-1b77w6xrqfubi****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The purchase mode of cloud phone instances.
	//
	// Valid values:
	//
	// 	- Instance (default): the instance group mode.
	//
	// 	- Node: the matrix mode [whitelisted].
	//
	// example:
	//
	// standard
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	// The status of the instance group.
	//
	// Valid values:
	//
	// 	- UPDATING_FAILED: The image update for the instance group failed.
	//
	// 	- FAILED: The instance group failed to be created.
	//
	// 	- RUNNING: The instance group is available.
	//
	// 	- EXPIRED: The instance group expired.
	//
	// 	- DELETING: The instance group is being deleted.
	//
	// 	- DELETED: The instance group is deleted.
	//
	// 	- UPDATING: The instance group is undergoing an image update.
	//
	// 	- CREATING: The instance group is being created.
	//
	// example:
	//
	// CREATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *DescribeAndroidInstanceGroupsRequest) Validate() error {
	return dara.Validate(s)
}
