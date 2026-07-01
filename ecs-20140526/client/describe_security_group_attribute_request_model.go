// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttribute(v string) *DescribeSecurityGroupAttributeRequest
	GetAttribute() *string
	SetDirection(v string) *DescribeSecurityGroupAttributeRequest
	GetDirection() *string
	SetMaxResults(v int32) *DescribeSecurityGroupAttributeRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSecurityGroupAttributeRequest
	GetNextToken() *string
	SetNicType(v string) *DescribeSecurityGroupAttributeRequest
	GetNicType() *string
	SetOwnerAccount(v string) *DescribeSecurityGroupAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSecurityGroupAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSecurityGroupAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSecurityGroupAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSecurityGroupAttributeRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *DescribeSecurityGroupAttributeRequest
	GetSecurityGroupId() *string
}

type DescribeSecurityGroupAttributeRequest struct {
	// The security group attribute. Valid values:
	//
	// - snapshotPolicyIds: queries the snapshot policies associated with the security group.
	//
	// example:
	//
	// snapshotPolicyIds
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The direction of the security group rule. Valid values:
	//
	//
	//
	// - egress: outbound.
	//
	// - ingress: inbound.
	//
	// - all: both inbound and outbound.
	//
	// Default value: all.
	//
	// example:
	//
	// all
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The maximum number of entries per page for a paged query.
	//
	// - Minimum value: 10.
	//
	// - Maximum value: 1000.
	//
	// Default value: 500.
	//
	// example:
	//
	// 500
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous call. You do not need to set this parameter for the first request.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The network type of the security group rule.
	//
	// - For security groups in a VPC, the only valid value is intranet (default), which indicates internal network.
	//
	//     > If you set this parameter to internet or leave it empty, the value is automatically set to intranet.
	//
	// - Valid values for security groups in the classic network:
	//
	//     - internet (default): Internet.
	//
	//     - intranet: internal network.
	//
	//     > The classic network feature has been offline. For details, see [Retirement announcement](https://help.aliyun.com/document_detail/2833134.html).
	//
	// example:
	//
	// intranet
	NicType      *string `json:"NicType,omitempty" xml:"NicType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the security group. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The security group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp1gxw6bznjjvhu3****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DescribeSecurityGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeRequest) GetAttribute() *string {
	return s.Attribute
}

func (s *DescribeSecurityGroupAttributeRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeSecurityGroupAttributeRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSecurityGroupAttributeRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSecurityGroupAttributeRequest) GetNicType() *string {
	return s.NicType
}

func (s *DescribeSecurityGroupAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSecurityGroupAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSecurityGroupAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityGroupAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSecurityGroupAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSecurityGroupAttributeRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeSecurityGroupAttributeRequest) SetAttribute(v string) *DescribeSecurityGroupAttributeRequest {
	s.Attribute = &v
	return s
}

func (s *DescribeSecurityGroupAttributeRequest) SetDirection(v string) *DescribeSecurityGroupAttributeRequest {
	s.Direction = &v
	return s
}

func (s *DescribeSecurityGroupAttributeRequest) SetMaxResults(v int32) *DescribeSecurityGroupAttributeRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSecurityGroupAttributeRequest) SetNextToken(v string) *DescribeSecurityGroupAttributeRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSecurityGroupAttributeRequest) SetNicType(v string) *DescribeSecurityGroupAttributeRequest {
	s.NicType = &v
	return s
}

func (s *DescribeSecurityGroupAttributeRequest) SetOwnerAccount(v string) *DescribeSecurityGroupAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSecurityGroupAttributeRequest) SetOwnerId(v int64) *DescribeSecurityGroupAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeRequest) SetRegionId(v string) *DescribeSecurityGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeRequest) SetResourceOwnerAccount(v string) *DescribeSecurityGroupAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSecurityGroupAttributeRequest) SetResourceOwnerId(v int64) *DescribeSecurityGroupAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeRequest) SetSecurityGroupId(v string) *DescribeSecurityGroupAttributeRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeRequest) Validate() error {
	return dara.Validate(s)
}
