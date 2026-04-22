// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePvtzStatisticsGlobalOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkParams(v []*DescribePvtzStatisticsGlobalOverviewRequestNetworkParams) *DescribePvtzStatisticsGlobalOverviewRequest
	GetNetworkParams() []*DescribePvtzStatisticsGlobalOverviewRequestNetworkParams
	SetOverviewPeriod(v string) *DescribePvtzStatisticsGlobalOverviewRequest
	GetOverviewPeriod() *string
	SetServerRegion(v string) *DescribePvtzStatisticsGlobalOverviewRequest
	GetServerRegion() *string
}

type DescribePvtzStatisticsGlobalOverviewRequest struct {
	NetworkParams []*DescribePvtzStatisticsGlobalOverviewRequestNetworkParams `json:"NetworkParams,omitempty" xml:"NetworkParams,omitempty" type:"Repeated"`
	// example:
	//
	// DAY, WEEK, MONTH
	OverviewPeriod *string `json:"OverviewPeriod,omitempty" xml:"OverviewPeriod,omitempty"`
	// example:
	//
	// cn-hangzhou
	ServerRegion *string `json:"ServerRegion,omitempty" xml:"ServerRegion,omitempty"`
}

func (s DescribePvtzStatisticsGlobalOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsGlobalOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsGlobalOverviewRequest) GetNetworkParams() []*DescribePvtzStatisticsGlobalOverviewRequestNetworkParams {
	return s.NetworkParams
}

func (s *DescribePvtzStatisticsGlobalOverviewRequest) GetOverviewPeriod() *string {
	return s.OverviewPeriod
}

func (s *DescribePvtzStatisticsGlobalOverviewRequest) GetServerRegion() *string {
	return s.ServerRegion
}

func (s *DescribePvtzStatisticsGlobalOverviewRequest) SetNetworkParams(v []*DescribePvtzStatisticsGlobalOverviewRequestNetworkParams) *DescribePvtzStatisticsGlobalOverviewRequest {
	s.NetworkParams = v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewRequest) SetOverviewPeriod(v string) *DescribePvtzStatisticsGlobalOverviewRequest {
	s.OverviewPeriod = &v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewRequest) SetServerRegion(v string) *DescribePvtzStatisticsGlobalOverviewRequest {
	s.ServerRegion = &v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewRequest) Validate() error {
	if s.NetworkParams != nil {
		for _, item := range s.NetworkParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePvtzStatisticsGlobalOverviewRequestNetworkParams struct {
	// example:
	//
	// cn-hongkong
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-bp1y5y4wk5810n50cx765
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// 5924158341670144
	VpcOwner *string `json:"VpcOwner,omitempty" xml:"VpcOwner,omitempty"`
	// example:
	//
	// STANDARD
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
}

func (s DescribePvtzStatisticsGlobalOverviewRequestNetworkParams) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsGlobalOverviewRequestNetworkParams) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsGlobalOverviewRequestNetworkParams) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePvtzStatisticsGlobalOverviewRequestNetworkParams) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribePvtzStatisticsGlobalOverviewRequestNetworkParams) GetVpcOwner() *string {
	return s.VpcOwner
}

func (s *DescribePvtzStatisticsGlobalOverviewRequestNetworkParams) GetVpcType() *string {
	return s.VpcType
}

func (s *DescribePvtzStatisticsGlobalOverviewRequestNetworkParams) SetRegionId(v string) *DescribePvtzStatisticsGlobalOverviewRequestNetworkParams {
	s.RegionId = &v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewRequestNetworkParams) SetVpcId(v string) *DescribePvtzStatisticsGlobalOverviewRequestNetworkParams {
	s.VpcId = &v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewRequestNetworkParams) SetVpcOwner(v string) *DescribePvtzStatisticsGlobalOverviewRequestNetworkParams {
	s.VpcOwner = &v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewRequestNetworkParams) SetVpcType(v string) *DescribePvtzStatisticsGlobalOverviewRequestNetworkParams {
	s.VpcType = &v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewRequestNetworkParams) Validate() error {
	return dara.Validate(s)
}
