// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DescribeVSwitchAttributesRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DescribeVSwitchAttributesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVSwitchAttributesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeVSwitchAttributesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVSwitchAttributesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVSwitchAttributesRequest
	GetResourceOwnerId() *int64
	SetVSwitchId(v string) *DescribeVSwitchAttributesRequest
	GetVSwitchId() *string
}

type DescribeVSwitchAttributesRequest struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends a request and performs a dry run. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the virtual private cloud (VPC) to which the custom route table belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-25naue4g****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeVSwitchAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchAttributesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeVSwitchAttributesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVSwitchAttributesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVSwitchAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVSwitchAttributesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVSwitchAttributesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVSwitchAttributesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchAttributesRequest) SetDryRun(v bool) *DescribeVSwitchAttributesRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeVSwitchAttributesRequest) SetOwnerAccount(v string) *DescribeVSwitchAttributesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVSwitchAttributesRequest) SetOwnerId(v int64) *DescribeVSwitchAttributesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVSwitchAttributesRequest) SetRegionId(v string) *DescribeVSwitchAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVSwitchAttributesRequest) SetResourceOwnerAccount(v string) *DescribeVSwitchAttributesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVSwitchAttributesRequest) SetResourceOwnerId(v int64) *DescribeVSwitchAttributesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVSwitchAttributesRequest) SetVSwitchId(v string) *DescribeVSwitchAttributesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchAttributesRequest) Validate() error {
	return dara.Validate(s)
}
