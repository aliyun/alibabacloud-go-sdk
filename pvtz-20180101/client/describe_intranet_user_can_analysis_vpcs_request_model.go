// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIntranetUserCanAnalysisVpcsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkType(v string) *DescribeIntranetUserCanAnalysisVpcsRequest
	GetNetworkType() *string
	SetPageNumber(v int32) *DescribeIntranetUserCanAnalysisVpcsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIntranetUserCanAnalysisVpcsRequest
	GetPageSize() *int32
	SetQueryTreeLevel(v int32) *DescribeIntranetUserCanAnalysisVpcsRequest
	GetQueryTreeLevel() *int32
	SetRegionId(v string) *DescribeIntranetUserCanAnalysisVpcsRequest
	GetRegionId() *string
	SetVpcOwner(v int64) *DescribeIntranetUserCanAnalysisVpcsRequest
	GetVpcOwner() *int64
	SetVpcType(v string) *DescribeIntranetUserCanAnalysisVpcsRequest
	GetVpcType() *string
}

type DescribeIntranetUserCanAnalysisVpcsRequest struct {
	// example:
	//
	// CLOUD
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryTreeLevel *int32 `json:"QueryTreeLevel,omitempty" xml:"QueryTreeLevel,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1204673843521685
	VpcOwner *int64 `json:"VpcOwner,omitempty" xml:"VpcOwner,omitempty"`
	// example:
	//
	// STANDARD
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
}

func (s DescribeIntranetUserCanAnalysisVpcsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntranetUserCanAnalysisVpcsRequest) GoString() string {
	return s.String()
}

func (s *DescribeIntranetUserCanAnalysisVpcsRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeIntranetUserCanAnalysisVpcsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIntranetUserCanAnalysisVpcsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIntranetUserCanAnalysisVpcsRequest) GetQueryTreeLevel() *int32 {
	return s.QueryTreeLevel
}

func (s *DescribeIntranetUserCanAnalysisVpcsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIntranetUserCanAnalysisVpcsRequest) GetVpcOwner() *int64 {
	return s.VpcOwner
}

func (s *DescribeIntranetUserCanAnalysisVpcsRequest) GetVpcType() *string {
	return s.VpcType
}

func (s *DescribeIntranetUserCanAnalysisVpcsRequest) SetNetworkType(v string) *DescribeIntranetUserCanAnalysisVpcsRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsRequest) SetPageNumber(v int32) *DescribeIntranetUserCanAnalysisVpcsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsRequest) SetPageSize(v int32) *DescribeIntranetUserCanAnalysisVpcsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsRequest) SetQueryTreeLevel(v int32) *DescribeIntranetUserCanAnalysisVpcsRequest {
	s.QueryTreeLevel = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsRequest) SetRegionId(v string) *DescribeIntranetUserCanAnalysisVpcsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsRequest) SetVpcOwner(v int64) *DescribeIntranetUserCanAnalysisVpcsRequest {
	s.VpcOwner = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsRequest) SetVpcType(v string) *DescribeIntranetUserCanAnalysisVpcsRequest {
	s.VpcType = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsRequest) Validate() error {
	return dara.Validate(s)
}
