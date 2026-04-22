// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePvtzStatisticsZoneOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribePvtzStatisticsZoneOverviewRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *DescribePvtzStatisticsZoneOverviewRequest
	GetClientToken() *string
	SetModule(v string) *DescribePvtzStatisticsZoneOverviewRequest
	GetModule() *string
	SetNetworkParams(v []*DescribePvtzStatisticsZoneOverviewRequestNetworkParams) *DescribePvtzStatisticsZoneOverviewRequest
	GetNetworkParams() []*DescribePvtzStatisticsZoneOverviewRequestNetworkParams
	SetOverviewPeriod(v string) *DescribePvtzStatisticsZoneOverviewRequest
	GetOverviewPeriod() *string
	SetServerRegion(v string) *DescribePvtzStatisticsZoneOverviewRequest
	GetServerRegion() *string
}

type DescribePvtzStatisticsZoneOverviewRequest struct {
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// 234534535432323...
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// AUTHORITY
	Module         *string                                                   `json:"Module,omitempty" xml:"Module,omitempty"`
	NetworkParams  []*DescribePvtzStatisticsZoneOverviewRequestNetworkParams `json:"NetworkParams,omitempty" xml:"NetworkParams,omitempty" type:"Repeated"`
	OverviewPeriod *string                                                   `json:"OverviewPeriod,omitempty" xml:"OverviewPeriod,omitempty"`
	// example:
	//
	// cn-hangzhou
	ServerRegion *string `json:"ServerRegion,omitempty" xml:"ServerRegion,omitempty"`
}

func (s DescribePvtzStatisticsZoneOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsZoneOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsZoneOverviewRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribePvtzStatisticsZoneOverviewRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribePvtzStatisticsZoneOverviewRequest) GetModule() *string {
	return s.Module
}

func (s *DescribePvtzStatisticsZoneOverviewRequest) GetNetworkParams() []*DescribePvtzStatisticsZoneOverviewRequestNetworkParams {
	return s.NetworkParams
}

func (s *DescribePvtzStatisticsZoneOverviewRequest) GetOverviewPeriod() *string {
	return s.OverviewPeriod
}

func (s *DescribePvtzStatisticsZoneOverviewRequest) GetServerRegion() *string {
	return s.ServerRegion
}

func (s *DescribePvtzStatisticsZoneOverviewRequest) SetAcceptLanguage(v string) *DescribePvtzStatisticsZoneOverviewRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewRequest) SetClientToken(v string) *DescribePvtzStatisticsZoneOverviewRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewRequest) SetModule(v string) *DescribePvtzStatisticsZoneOverviewRequest {
	s.Module = &v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewRequest) SetNetworkParams(v []*DescribePvtzStatisticsZoneOverviewRequestNetworkParams) *DescribePvtzStatisticsZoneOverviewRequest {
	s.NetworkParams = v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewRequest) SetOverviewPeriod(v string) *DescribePvtzStatisticsZoneOverviewRequest {
	s.OverviewPeriod = &v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewRequest) SetServerRegion(v string) *DescribePvtzStatisticsZoneOverviewRequest {
	s.ServerRegion = &v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewRequest) Validate() error {
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

type DescribePvtzStatisticsZoneOverviewRequestNetworkParams struct {
	// example:
	//
	// cn-hongkong
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-uf68l2l04nqoyg7ie1kaw
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// 5209821778041881
	VpcOwner *string `json:"VpcOwner,omitempty" xml:"VpcOwner,omitempty"`
	// example:
	//
	// STANDARD
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
}

func (s DescribePvtzStatisticsZoneOverviewRequestNetworkParams) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsZoneOverviewRequestNetworkParams) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsZoneOverviewRequestNetworkParams) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePvtzStatisticsZoneOverviewRequestNetworkParams) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribePvtzStatisticsZoneOverviewRequestNetworkParams) GetVpcOwner() *string {
	return s.VpcOwner
}

func (s *DescribePvtzStatisticsZoneOverviewRequestNetworkParams) GetVpcType() *string {
	return s.VpcType
}

func (s *DescribePvtzStatisticsZoneOverviewRequestNetworkParams) SetRegionId(v string) *DescribePvtzStatisticsZoneOverviewRequestNetworkParams {
	s.RegionId = &v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewRequestNetworkParams) SetVpcId(v string) *DescribePvtzStatisticsZoneOverviewRequestNetworkParams {
	s.VpcId = &v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewRequestNetworkParams) SetVpcOwner(v string) *DescribePvtzStatisticsZoneOverviewRequestNetworkParams {
	s.VpcOwner = &v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewRequestNetworkParams) SetVpcType(v string) *DescribePvtzStatisticsZoneOverviewRequestNetworkParams {
	s.VpcType = &v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewRequestNetworkParams) Validate() error {
	return dara.Validate(s)
}
