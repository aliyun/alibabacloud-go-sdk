// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePvtzStatisticsHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribePvtzStatisticsHistoryRequest
	GetDomainName() *string
	SetEndTimestamp(v string) *DescribePvtzStatisticsHistoryRequest
	GetEndTimestamp() *string
	SetModule(v string) *DescribePvtzStatisticsHistoryRequest
	GetModule() *string
	SetNetworkParams(v []*DescribePvtzStatisticsHistoryRequestNetworkParams) *DescribePvtzStatisticsHistoryRequest
	GetNetworkParams() []*DescribePvtzStatisticsHistoryRequestNetworkParams
	SetRcode(v string) *DescribePvtzStatisticsHistoryRequest
	GetRcode() *string
	SetServerRegion(v string) *DescribePvtzStatisticsHistoryRequest
	GetServerRegion() *string
	SetStartTimestamp(v string) *DescribePvtzStatisticsHistoryRequest
	GetStartTimestamp() *string
	SetStatisticalType(v string) *DescribePvtzStatisticsHistoryRequest
	GetStatisticalType() *string
	SetZoneName(v string) *DescribePvtzStatisticsHistoryRequest
	GetZoneName() *string
}

type DescribePvtzStatisticsHistoryRequest struct {
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 1571673600000
	EndTimestamp  *string                                              `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	Module        *string                                              `json:"Module,omitempty" xml:"Module,omitempty"`
	NetworkParams []*DescribePvtzStatisticsHistoryRequestNetworkParams `json:"NetworkParams,omitempty" xml:"NetworkParams,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Rcode *string `json:"Rcode,omitempty" xml:"Rcode,omitempty"`
	// example:
	//
	// cn-hangzhou
	ServerRegion *string `json:"ServerRegion,omitempty" xml:"ServerRegion,omitempty"`
	// example:
	//
	// 1516779348000
	StartTimestamp  *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	StatisticalType *string `json:"StatisticalType,omitempty" xml:"StatisticalType,omitempty"`
	// example:
	//
	// host.local
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribePvtzStatisticsHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsHistoryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribePvtzStatisticsHistoryRequest) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribePvtzStatisticsHistoryRequest) GetModule() *string {
	return s.Module
}

func (s *DescribePvtzStatisticsHistoryRequest) GetNetworkParams() []*DescribePvtzStatisticsHistoryRequestNetworkParams {
	return s.NetworkParams
}

func (s *DescribePvtzStatisticsHistoryRequest) GetRcode() *string {
	return s.Rcode
}

func (s *DescribePvtzStatisticsHistoryRequest) GetServerRegion() *string {
	return s.ServerRegion
}

func (s *DescribePvtzStatisticsHistoryRequest) GetStartTimestamp() *string {
	return s.StartTimestamp
}

func (s *DescribePvtzStatisticsHistoryRequest) GetStatisticalType() *string {
	return s.StatisticalType
}

func (s *DescribePvtzStatisticsHistoryRequest) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribePvtzStatisticsHistoryRequest) SetDomainName(v string) *DescribePvtzStatisticsHistoryRequest {
	s.DomainName = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryRequest) SetEndTimestamp(v string) *DescribePvtzStatisticsHistoryRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryRequest) SetModule(v string) *DescribePvtzStatisticsHistoryRequest {
	s.Module = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryRequest) SetNetworkParams(v []*DescribePvtzStatisticsHistoryRequestNetworkParams) *DescribePvtzStatisticsHistoryRequest {
	s.NetworkParams = v
	return s
}

func (s *DescribePvtzStatisticsHistoryRequest) SetRcode(v string) *DescribePvtzStatisticsHistoryRequest {
	s.Rcode = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryRequest) SetServerRegion(v string) *DescribePvtzStatisticsHistoryRequest {
	s.ServerRegion = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryRequest) SetStartTimestamp(v string) *DescribePvtzStatisticsHistoryRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryRequest) SetStatisticalType(v string) *DescribePvtzStatisticsHistoryRequest {
	s.StatisticalType = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryRequest) SetZoneName(v string) *DescribePvtzStatisticsHistoryRequest {
	s.ZoneName = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryRequest) Validate() error {
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

type DescribePvtzStatisticsHistoryRequestNetworkParams struct {
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// vpc-bp1hneq5pcy2gv87op0uf
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// 1634808626233933
	VpcOwner *string `json:"VpcOwner,omitempty" xml:"VpcOwner,omitempty"`
	// example:
	//
	// STANDARD
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
}

func (s DescribePvtzStatisticsHistoryRequestNetworkParams) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsHistoryRequestNetworkParams) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsHistoryRequestNetworkParams) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePvtzStatisticsHistoryRequestNetworkParams) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribePvtzStatisticsHistoryRequestNetworkParams) GetVpcOwner() *string {
	return s.VpcOwner
}

func (s *DescribePvtzStatisticsHistoryRequestNetworkParams) GetVpcType() *string {
	return s.VpcType
}

func (s *DescribePvtzStatisticsHistoryRequestNetworkParams) SetRegionId(v string) *DescribePvtzStatisticsHistoryRequestNetworkParams {
	s.RegionId = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryRequestNetworkParams) SetVpcId(v string) *DescribePvtzStatisticsHistoryRequestNetworkParams {
	s.VpcId = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryRequestNetworkParams) SetVpcOwner(v string) *DescribePvtzStatisticsHistoryRequestNetworkParams {
	s.VpcOwner = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryRequestNetworkParams) SetVpcType(v string) *DescribePvtzStatisticsHistoryRequestNetworkParams {
	s.VpcType = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryRequestNetworkParams) Validate() error {
	return dara.Validate(s)
}
