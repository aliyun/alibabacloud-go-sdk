// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateNetworkPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkPackageId(v string) *AssociateNetworkPackageRequest
	GetNetworkPackageId() *string
	SetOfficeSiteId(v string) *AssociateNetworkPackageRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *AssociateNetworkPackageRequest
	GetRegionId() *string
}

type AssociateNetworkPackageRequest struct {
	// The ID of the premium bandwidth plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// np-e0iodl3yzb62q****
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	// The ID of the office network. You can call the [DescribeNetworkPackages](https://help.aliyun.com/document_detail/216079.html) to obtain the ID of the office network to which a premium bandwidth plan is bound.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-*********
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AssociateNetworkPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateNetworkPackageRequest) GoString() string {
	return s.String()
}

func (s *AssociateNetworkPackageRequest) GetNetworkPackageId() *string {
	return s.NetworkPackageId
}

func (s *AssociateNetworkPackageRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *AssociateNetworkPackageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateNetworkPackageRequest) SetNetworkPackageId(v string) *AssociateNetworkPackageRequest {
	s.NetworkPackageId = &v
	return s
}

func (s *AssociateNetworkPackageRequest) SetOfficeSiteId(v string) *AssociateNetworkPackageRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *AssociateNetworkPackageRequest) SetRegionId(v string) *AssociateNetworkPackageRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateNetworkPackageRequest) Validate() error {
	return dara.Validate(s)
}
