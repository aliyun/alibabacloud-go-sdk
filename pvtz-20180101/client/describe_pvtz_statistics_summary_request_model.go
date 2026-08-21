// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePvtzStatisticsSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *DescribePvtzStatisticsSummaryRequest
	GetDirection() *string
	SetDomainName(v string) *DescribePvtzStatisticsSummaryRequest
	GetDomainName() *string
	SetEndTimestamp(v string) *DescribePvtzStatisticsSummaryRequest
	GetEndTimestamp() *string
	SetGrowType(v string) *DescribePvtzStatisticsSummaryRequest
	GetGrowType() *string
	SetModule(v string) *DescribePvtzStatisticsSummaryRequest
	GetModule() *string
	SetNetworkParams(v []*DescribePvtzStatisticsSummaryRequestNetworkParams) *DescribePvtzStatisticsSummaryRequest
	GetNetworkParams() []*DescribePvtzStatisticsSummaryRequestNetworkParams
	SetOrderBy(v string) *DescribePvtzStatisticsSummaryRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *DescribePvtzStatisticsSummaryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePvtzStatisticsSummaryRequest
	GetPageSize() *int32
	SetPeriod(v string) *DescribePvtzStatisticsSummaryRequest
	GetPeriod() *string
	SetRcode(v string) *DescribePvtzStatisticsSummaryRequest
	GetRcode() *string
	SetServerRegion(v string) *DescribePvtzStatisticsSummaryRequest
	GetServerRegion() *string
	SetStartTimestamp(v string) *DescribePvtzStatisticsSummaryRequest
	GetStartTimestamp() *string
	SetStatisticalType(v string) *DescribePvtzStatisticsSummaryRequest
	GetStatisticalType() *string
	SetZoneName(v string) *DescribePvtzStatisticsSummaryRequest
	GetZoneName() *string
}

type DescribePvtzStatisticsSummaryRequest struct {
	// The sort order. Valid values: ASC and DESC.
	//
	// example:
	//
	// ASC
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query, specified as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1571673600000
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// Specifies the type of change in request volume to report. Valid values: up (request spikes) and down (request drops).
	//
	// example:
	//
	// up
	GrowType *string `json:"GrowType,omitempty" xml:"GrowType,omitempty"`
	// The statistics module. Valid values: AUTHORITY, AUTH_FAST, AUTH_SLOW, GLOBAL, CACHE, FORWARD, and RECURSION.
	//
	// example:
	//
	// AUTH_FAST
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The network parameters.
	NetworkParams []*DescribePvtzStatisticsSummaryRequestNetworkParams `json:"NetworkParams,omitempty" xml:"NetworkParams,omitempty" type:"Repeated"`
	// The field by which to sort the results. To sort by fluctuation ratio, set this parameter to fluctuation_ratio.
	//
	// example:
	//
	// fluctuation_ratio
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number to return.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The time granularity for the statistics.
	//
	// example:
	//
	// day
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The DNS response code.
	//
	// example:
	//
	// 0
	Rcode *string `json:"Rcode,omitempty" xml:"Rcode,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-hangzhou
	ServerRegion *string `json:"ServerRegion,omitempty" xml:"ServerRegion,omitempty"`
	// The beginning of the time range to query, specified as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1571587200000
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The type of statistics to collect. Valid values: QTYPE_RATIO: The ratio of DNS request record types. NEGATIVE_RCODE: Statistics on negative DNS responses. RCODE_DOMAIN: The top domain names that receive negative responses. RCODE_SOURCE_IP: The top source IP addresses for requests to a specified domain name that receive negative responses. REQUEST_ZONE: The top zones by request volume. REQUEST_DOMAIN: The top subdomains by request volume. VPC_RATIO: The ratio of request sources. SOURCE_VPC: Analysis of high-traffic source networks. SOURCE_IP: Statistics on source IP addresses. LINE_RATIO: The percentage of traffic per resolution line. LINE_HIT: Details about resolution line hits.
	//
	// example:
	//
	// - QTYPE_RATIO: Distribution ratio of DNS request record types.
	//
	// - NEGATIVE_RCODE: Negative DNS response type statistics.
	//
	// - RCODE_DOMAIN: Top domains by negative response count
	//
	// - RCODE_SOURCE_IP: Top source IP addresses requesting a specific domain with negative responses.
	//
	// - REQUEST_ZONE: Domain request volume ranking (zone level).
	//
	// - REQUEST_DOMAIN: Subdomain request volume ranking (domain name level).
	//
	// - VPC_RATIO: Request source distribution ratio.
	//
	// - SOURCE_VPC: Hot source network analysis for requests.
	//
	// - SOURCE_IP: Request source IP address statistics.
	//
	// - LINE_RATIO: Traffic distribution ratio by resolution line.
	//
	// - LINE_HIT: Resolution line hit details.
	StatisticalType *string `json:"StatisticalType,omitempty" xml:"StatisticalType,omitempty"`
	// The zone name.
	//
	// example:
	//
	// host.local
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribePvtzStatisticsSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsSummaryRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribePvtzStatisticsSummaryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribePvtzStatisticsSummaryRequest) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribePvtzStatisticsSummaryRequest) GetGrowType() *string {
	return s.GrowType
}

func (s *DescribePvtzStatisticsSummaryRequest) GetModule() *string {
	return s.Module
}

func (s *DescribePvtzStatisticsSummaryRequest) GetNetworkParams() []*DescribePvtzStatisticsSummaryRequestNetworkParams {
	return s.NetworkParams
}

func (s *DescribePvtzStatisticsSummaryRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribePvtzStatisticsSummaryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePvtzStatisticsSummaryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePvtzStatisticsSummaryRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribePvtzStatisticsSummaryRequest) GetRcode() *string {
	return s.Rcode
}

func (s *DescribePvtzStatisticsSummaryRequest) GetServerRegion() *string {
	return s.ServerRegion
}

func (s *DescribePvtzStatisticsSummaryRequest) GetStartTimestamp() *string {
	return s.StartTimestamp
}

func (s *DescribePvtzStatisticsSummaryRequest) GetStatisticalType() *string {
	return s.StatisticalType
}

func (s *DescribePvtzStatisticsSummaryRequest) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribePvtzStatisticsSummaryRequest) SetDirection(v string) *DescribePvtzStatisticsSummaryRequest {
	s.Direction = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequest) SetDomainName(v string) *DescribePvtzStatisticsSummaryRequest {
	s.DomainName = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequest) SetEndTimestamp(v string) *DescribePvtzStatisticsSummaryRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequest) SetGrowType(v string) *DescribePvtzStatisticsSummaryRequest {
	s.GrowType = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequest) SetModule(v string) *DescribePvtzStatisticsSummaryRequest {
	s.Module = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequest) SetNetworkParams(v []*DescribePvtzStatisticsSummaryRequestNetworkParams) *DescribePvtzStatisticsSummaryRequest {
	s.NetworkParams = v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequest) SetOrderBy(v string) *DescribePvtzStatisticsSummaryRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequest) SetPageNumber(v int32) *DescribePvtzStatisticsSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequest) SetPageSize(v int32) *DescribePvtzStatisticsSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequest) SetPeriod(v string) *DescribePvtzStatisticsSummaryRequest {
	s.Period = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequest) SetRcode(v string) *DescribePvtzStatisticsSummaryRequest {
	s.Rcode = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequest) SetServerRegion(v string) *DescribePvtzStatisticsSummaryRequest {
	s.ServerRegion = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequest) SetStartTimestamp(v string) *DescribePvtzStatisticsSummaryRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequest) SetStatisticalType(v string) *DescribePvtzStatisticsSummaryRequest {
	s.StatisticalType = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequest) SetZoneName(v string) *DescribePvtzStatisticsSummaryRequest {
	s.ZoneName = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequest) Validate() error {
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

type DescribePvtzStatisticsSummaryRequestNetworkParams struct {
	// The ID of the region where the VPC is deployed.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp1uvv79h1t8unnzdh3nq
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the VPC.
	//
	// example:
	//
	// 1256177436790486
	VpcOwner *string `json:"VpcOwner,omitempty" xml:"VpcOwner,omitempty"`
	// The type of the VPC.
	//
	// - STANDARD: A standard VPC.
	//
	// - EDS: A VPC for Elastic Desktop Service (EDS).
	//
	// example:
	//
	// STANDARD
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
}

func (s DescribePvtzStatisticsSummaryRequestNetworkParams) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsSummaryRequestNetworkParams) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsSummaryRequestNetworkParams) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePvtzStatisticsSummaryRequestNetworkParams) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribePvtzStatisticsSummaryRequestNetworkParams) GetVpcOwner() *string {
	return s.VpcOwner
}

func (s *DescribePvtzStatisticsSummaryRequestNetworkParams) GetVpcType() *string {
	return s.VpcType
}

func (s *DescribePvtzStatisticsSummaryRequestNetworkParams) SetRegionId(v string) *DescribePvtzStatisticsSummaryRequestNetworkParams {
	s.RegionId = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequestNetworkParams) SetVpcId(v string) *DescribePvtzStatisticsSummaryRequestNetworkParams {
	s.VpcId = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequestNetworkParams) SetVpcOwner(v string) *DescribePvtzStatisticsSummaryRequestNetworkParams {
	s.VpcOwner = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequestNetworkParams) SetVpcType(v string) *DescribePvtzStatisticsSummaryRequestNetworkParams {
	s.VpcType = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryRequestNetworkParams) Validate() error {
	return dara.Validate(s)
}
