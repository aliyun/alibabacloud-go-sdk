// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DescribeVpcAttributeRequest
	GetDryRun() *bool
	SetIsDefault(v bool) *DescribeVpcAttributeRequest
	GetIsDefault() *bool
	SetOwnerAccount(v string) *DescribeVpcAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVpcAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeVpcAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVpcAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVpcAttributeRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *DescribeVpcAttributeRequest
	GetVpcId() *string
}

type DescribeVpcAttributeRequest struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether the VPC is the default VPC. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// false
	IsDefault    *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the VPC is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPC that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp18sth14qii3pnv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVpcAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeVpcAttributeRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVpcAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVpcAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpcAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpcAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVpcAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVpcAttributeRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcAttributeRequest) SetDryRun(v bool) *DescribeVpcAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeVpcAttributeRequest) SetIsDefault(v bool) *DescribeVpcAttributeRequest {
	s.IsDefault = &v
	return s
}

func (s *DescribeVpcAttributeRequest) SetOwnerAccount(v string) *DescribeVpcAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVpcAttributeRequest) SetOwnerId(v int64) *DescribeVpcAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcAttributeRequest) SetRegionId(v string) *DescribeVpcAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpcAttributeRequest) SetResourceOwnerAccount(v string) *DescribeVpcAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVpcAttributeRequest) SetResourceOwnerId(v int64) *DescribeVpcAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVpcAttributeRequest) SetVpcId(v string) *DescribeVpcAttributeRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcAttributeRequest) Validate() error {
	return dara.Validate(s)
}
