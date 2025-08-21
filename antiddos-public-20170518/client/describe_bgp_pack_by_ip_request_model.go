// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBgpPackByIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDdosRegionId(v string) *DescribeBgpPackByIpRequest
	GetDdosRegionId() *string
	SetIp(v string) *DescribeBgpPackByIpRequest
	GetIp() *string
}

type DescribeBgpPackByIpRequest struct {
	// The region ID of the asset to query.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/353250.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// The IP address of the asset to query.
	//
	// >  You can call the [DescribeInstanceIpAddress](https://help.aliyun.com/document_detail/472620.html) operation to query the IDs of Elastic Compute Service (ECS) instances, Server Load Balancer (SLB) instances, and elastic IP addresses (EIPs) within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 118.31.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeBgpPackByIpRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpPackByIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeBgpPackByIpRequest) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *DescribeBgpPackByIpRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeBgpPackByIpRequest) SetDdosRegionId(v string) *DescribeBgpPackByIpRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeBgpPackByIpRequest) SetIp(v string) *DescribeBgpPackByIpRequest {
	s.Ip = &v
	return s
}

func (s *DescribeBgpPackByIpRequest) Validate() error {
	return dara.Validate(s)
}
