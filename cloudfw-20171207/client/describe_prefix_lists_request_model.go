// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrefixListsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionNo(v string) *DescribePrefixListsRequest
	GetRegionNo() *string
	SetSourceIp(v string) *DescribePrefixListsRequest
	GetSourceIp() *string
}

type DescribePrefixListsRequest struct {
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 47.100.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribePrefixListsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListsRequest) GoString() string {
	return s.String()
}

func (s *DescribePrefixListsRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribePrefixListsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribePrefixListsRequest) SetRegionNo(v string) *DescribePrefixListsRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribePrefixListsRequest) SetSourceIp(v string) *DescribePrefixListsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePrefixListsRequest) Validate() error {
	return dara.Validate(s)
}
