// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkPackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInternetChargeType(v string) *DescribeNetworkPackagesRequest
	GetInternetChargeType() *string
	SetMaxResults(v int32) *DescribeNetworkPackagesRequest
	GetMaxResults() *int32
	SetNetworkPackageId(v []*string) *DescribeNetworkPackagesRequest
	GetNetworkPackageId() []*string
	SetNextToken(v string) *DescribeNetworkPackagesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeNetworkPackagesRequest
	GetRegionId() *string
}

type DescribeNetworkPackagesRequest struct {
	// The charge type of the pay-as-you-go premium bandwidth plan.
	//
	// Valid values:
	//
	// 	- PayByTraffic: charges by data transfer.
	//
	// 	- PayByBandwidth: charges by fixed bandwidth.
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The number of entries to return on each page.
	//
	// 	- Maximum value: 100
	//
	// 	- Default value: 10
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The ID of the premium bandwidth plan. You can specify 1 to 100 IDs.
	//
	// example:
	//
	// np-amtp8e8q1o9e4****
	NetworkPackageId []*string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty" type:"Repeated"`
	// The token that determines the start point of the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeNetworkPackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkPackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeNetworkPackagesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeNetworkPackagesRequest) GetNetworkPackageId() []*string {
	return s.NetworkPackageId
}

func (s *DescribeNetworkPackagesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNetworkPackagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNetworkPackagesRequest) SetInternetChargeType(v string) *DescribeNetworkPackagesRequest {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetMaxResults(v int32) *DescribeNetworkPackagesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetNetworkPackageId(v []*string) *DescribeNetworkPackagesRequest {
	s.NetworkPackageId = v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetNextToken(v string) *DescribeNetworkPackagesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetRegionId(v string) *DescribeNetworkPackagesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNetworkPackagesRequest) Validate() error {
	return dara.Validate(s)
}
