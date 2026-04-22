// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInterAuthStatisticsSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetDirection() *string
	SetDomainName(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetDomainName() *string
	SetEndTimestamp(v int64) *DescribeInterAuthStatisticsSummaryRequest
	GetEndTimestamp() *int64
	SetGrowType(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetGrowType() *string
	SetOrderBy(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *DescribeInterAuthStatisticsSummaryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInterAuthStatisticsSummaryRequest
	GetPageSize() *int32
	SetPeriod(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetPeriod() *string
	SetRcode(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetRcode() *string
	SetServerRegion(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetServerRegion() *string
	SetSourceType(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetSourceType() *string
	SetStartTimestamp(v int64) *DescribeInterAuthStatisticsSummaryRequest
	GetStartTimestamp() *int64
	SetStatisticalType(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetStatisticalType() *string
	SetZoneName(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetZoneName() *string
}

type DescribeInterAuthStatisticsSummaryRequest struct {
	// example:
	//
	// DESC
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 1741526400000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// example:
	//
	// up
	//
	// down
	GrowType *string `json:"GrowType,omitempty" xml:"GrowType,omitempty"`
	// example:
	//
	// default
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 300
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// 0
	Rcode *string `json:"Rcode,omitempty" xml:"Rcode,omitempty"`
	// example:
	//
	// ap-southeast-3
	ServerRegion *string `json:"ServerRegion,omitempty" xml:"ServerRegion,omitempty"`
	// example:
	//
	// File
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 1516779348000
	StartTimestamp  *int64  `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	StatisticalType *string `json:"StatisticalType,omitempty" xml:"StatisticalType,omitempty"`
	// example:
	//
	// example.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeInterAuthStatisticsSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetGrowType() *string {
	return s.GrowType
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetRcode() *string {
	return s.Rcode
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetServerRegion() *string {
	return s.ServerRegion
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetStatisticalType() *string {
	return s.StatisticalType
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetDirection(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.Direction = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetDomainName(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetEndTimestamp(v int64) *DescribeInterAuthStatisticsSummaryRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetGrowType(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.GrowType = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetOrderBy(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetPageNumber(v int32) *DescribeInterAuthStatisticsSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetPageSize(v int32) *DescribeInterAuthStatisticsSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetPeriod(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.Period = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetRcode(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.Rcode = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetServerRegion(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.ServerRegion = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetSourceType(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.SourceType = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetStartTimestamp(v int64) *DescribeInterAuthStatisticsSummaryRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetStatisticalType(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.StatisticalType = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetZoneName(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.ZoneName = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) Validate() error {
	return dara.Validate(s)
}
