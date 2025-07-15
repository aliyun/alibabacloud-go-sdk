// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateNetworkPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkPackageId(v string) *DissociateNetworkPackageRequest
	GetNetworkPackageId() *string
	SetRegionId(v string) *DissociateNetworkPackageRequest
	GetRegionId() *string
}

type DissociateNetworkPackageRequest struct {
	// The ID of the premium bandwidth plan. You can call the [DescribeNetworkPackages](https://help.aliyun.com/document_detail/216079.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// np-*********
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DissociateNetworkPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateNetworkPackageRequest) GoString() string {
	return s.String()
}

func (s *DissociateNetworkPackageRequest) GetNetworkPackageId() *string {
	return s.NetworkPackageId
}

func (s *DissociateNetworkPackageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DissociateNetworkPackageRequest) SetNetworkPackageId(v string) *DissociateNetworkPackageRequest {
	s.NetworkPackageId = &v
	return s
}

func (s *DissociateNetworkPackageRequest) SetRegionId(v string) *DissociateNetworkPackageRequest {
	s.RegionId = &v
	return s
}

func (s *DissociateNetworkPackageRequest) Validate() error {
	return dara.Validate(s)
}
