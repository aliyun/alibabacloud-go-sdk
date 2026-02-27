// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBgpNetworksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBgpNetworks(v *DescribeBgpNetworksResponseBodyBgpNetworks) *DescribeBgpNetworksResponseBody
	GetBgpNetworks() *DescribeBgpNetworksResponseBodyBgpNetworks
	SetPageNumber(v int32) *DescribeBgpNetworksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBgpNetworksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBgpNetworksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeBgpNetworksResponseBody
	GetTotalCount() *int32
}

type DescribeBgpNetworksResponseBody struct {
	BgpNetworks *DescribeBgpNetworksResponseBodyBgpNetworks `json:"BgpNetworks,omitempty" xml:"BgpNetworks,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6F513A15-669F-419D-B511-08A85292059B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of advertised BGP networks.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBgpNetworksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpNetworksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBgpNetworksResponseBody) GetBgpNetworks() *DescribeBgpNetworksResponseBodyBgpNetworks {
	return s.BgpNetworks
}

func (s *DescribeBgpNetworksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBgpNetworksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBgpNetworksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBgpNetworksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBgpNetworksResponseBody) SetBgpNetworks(v *DescribeBgpNetworksResponseBodyBgpNetworks) *DescribeBgpNetworksResponseBody {
	s.BgpNetworks = v
	return s
}

func (s *DescribeBgpNetworksResponseBody) SetPageNumber(v int32) *DescribeBgpNetworksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBgpNetworksResponseBody) SetPageSize(v int32) *DescribeBgpNetworksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBgpNetworksResponseBody) SetRequestId(v string) *DescribeBgpNetworksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBgpNetworksResponseBody) SetTotalCount(v int32) *DescribeBgpNetworksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBgpNetworksResponseBody) Validate() error {
	if s.BgpNetworks != nil {
		if err := s.BgpNetworks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBgpNetworksResponseBodyBgpNetworks struct {
	BgpNetwork []*DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork `json:"BgpNetwork,omitempty" xml:"BgpNetwork,omitempty" type:"Repeated"`
}

func (s DescribeBgpNetworksResponseBodyBgpNetworks) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpNetworksResponseBodyBgpNetworks) GoString() string {
	return s.String()
}

func (s *DescribeBgpNetworksResponseBodyBgpNetworks) GetBgpNetwork() []*DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork {
	return s.BgpNetwork
}

func (s *DescribeBgpNetworksResponseBodyBgpNetworks) SetBgpNetwork(v []*DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork) *DescribeBgpNetworksResponseBodyBgpNetworks {
	s.BgpNetwork = v
	return s
}

func (s *DescribeBgpNetworksResponseBodyBgpNetworks) Validate() error {
	if s.BgpNetwork != nil {
		for _, item := range s.BgpNetwork {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork struct {
	DstCidrBlock *string `json:"DstCidrBlock,omitempty" xml:"DstCidrBlock,omitempty"`
	RouterId     *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork) GoString() string {
	return s.String()
}

func (s *DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork) GetDstCidrBlock() *string {
	return s.DstCidrBlock
}

func (s *DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork) GetRouterId() *string {
	return s.RouterId
}

func (s *DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork) GetStatus() *string {
	return s.Status
}

func (s *DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork) SetDstCidrBlock(v string) *DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork {
	s.DstCidrBlock = &v
	return s
}

func (s *DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork) SetRouterId(v string) *DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork {
	s.RouterId = &v
	return s
}

func (s *DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork) SetStatus(v string) *DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork {
	s.Status = &v
	return s
}

func (s *DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork) SetVpcId(v string) *DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork {
	s.VpcId = &v
	return s
}

func (s *DescribeBgpNetworksResponseBodyBgpNetworksBgpNetwork) Validate() error {
	return dara.Validate(s)
}
