// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcHoneyPotListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeVpcHoneyPotListRequest
	GetCurrentPage() *int32
	SetHoneyPotExistence(v bool) *DescribeVpcHoneyPotListRequest
	GetHoneyPotExistence() *bool
	SetPageSize(v int32) *DescribeVpcHoneyPotListRequest
	GetPageSize() *int32
	SetVpcId(v string) *DescribeVpcHoneyPotListRequest
	GetVpcId() *string
	SetVpcName(v string) *DescribeVpcHoneyPotListRequest
	GetVpcName() *string
	SetVpcRegionId(v string) *DescribeVpcHoneyPotListRequest
	GetVpcRegionId() *string
}

type DescribeVpcHoneyPotListRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 2
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether the cloud honeypot feature is enabled for the VPCs. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	HoneyPotExistence *bool `json:"HoneyPotExistence,omitempty" xml:"HoneyPotExistence,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the VPC on which the honeypot is deployed.
	//
	// > You can call the [DescribeVpcList](~~DescribeVpcList~~) operation to query the IDs of VPCs.
	//
	// example:
	//
	// vpc-d7o009q63fqy21r8u****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// > You can call the [DescribeVpcList](~~DescribeVpcList~~) operation to query the names of VPCs.
	//
	// example:
	//
	// abcnet
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	// The region ID of the VPC.
	//
	// > You can call the [DescribeVpcList](~~DescribeVpcList~~) operation to query the region IDs of VPCs.
	//
	// example:
	//
	// ap-southeast-2
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
}

func (s DescribeVpcHoneyPotListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcHoneyPotListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeVpcHoneyPotListRequest) GetHoneyPotExistence() *bool {
	return s.HoneyPotExistence
}

func (s *DescribeVpcHoneyPotListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpcHoneyPotListRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcHoneyPotListRequest) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcHoneyPotListRequest) GetVpcRegionId() *string {
	return s.VpcRegionId
}

func (s *DescribeVpcHoneyPotListRequest) SetCurrentPage(v int32) *DescribeVpcHoneyPotListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcHoneyPotListRequest) SetHoneyPotExistence(v bool) *DescribeVpcHoneyPotListRequest {
	s.HoneyPotExistence = &v
	return s
}

func (s *DescribeVpcHoneyPotListRequest) SetPageSize(v int32) *DescribeVpcHoneyPotListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcHoneyPotListRequest) SetVpcId(v string) *DescribeVpcHoneyPotListRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcHoneyPotListRequest) SetVpcName(v string) *DescribeVpcHoneyPotListRequest {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcHoneyPotListRequest) SetVpcRegionId(v string) *DescribeVpcHoneyPotListRequest {
	s.VpcRegionId = &v
	return s
}

func (s *DescribeVpcHoneyPotListRequest) Validate() error {
	return dara.Validate(s)
}
