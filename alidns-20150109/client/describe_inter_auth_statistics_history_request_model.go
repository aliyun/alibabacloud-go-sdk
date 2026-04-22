// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInterAuthStatisticsHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeInterAuthStatisticsHistoryRequest
	GetDomainName() *string
	SetEndTimestamp(v int64) *DescribeInterAuthStatisticsHistoryRequest
	GetEndTimestamp() *int64
	SetRcode(v string) *DescribeInterAuthStatisticsHistoryRequest
	GetRcode() *string
	SetServerRegion(v string) *DescribeInterAuthStatisticsHistoryRequest
	GetServerRegion() *string
	SetStartTimestamp(v int64) *DescribeInterAuthStatisticsHistoryRequest
	GetStartTimestamp() *int64
	SetStatisticalType(v string) *DescribeInterAuthStatisticsHistoryRequest
	GetStatisticalType() *string
	SetZoneName(v string) *DescribeInterAuthStatisticsHistoryRequest
	GetZoneName() *string
}

type DescribeInterAuthStatisticsHistoryRequest struct {
	// example:
	//
	// dns-example.top
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 1741526400000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
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
	// 1474335170000
	StartTimestamp  *int64  `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	StatisticalType *string `json:"StatisticalType,omitempty" xml:"StatisticalType,omitempty"`
	// example:
	//
	// lisheng999.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeInterAuthStatisticsHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsHistoryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeInterAuthStatisticsHistoryRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *DescribeInterAuthStatisticsHistoryRequest) GetRcode() *string {
	return s.Rcode
}

func (s *DescribeInterAuthStatisticsHistoryRequest) GetServerRegion() *string {
	return s.ServerRegion
}

func (s *DescribeInterAuthStatisticsHistoryRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribeInterAuthStatisticsHistoryRequest) GetStatisticalType() *string {
	return s.StatisticalType
}

func (s *DescribeInterAuthStatisticsHistoryRequest) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeInterAuthStatisticsHistoryRequest) SetDomainName(v string) *DescribeInterAuthStatisticsHistoryRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryRequest) SetEndTimestamp(v int64) *DescribeInterAuthStatisticsHistoryRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryRequest) SetRcode(v string) *DescribeInterAuthStatisticsHistoryRequest {
	s.Rcode = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryRequest) SetServerRegion(v string) *DescribeInterAuthStatisticsHistoryRequest {
	s.ServerRegion = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryRequest) SetStartTimestamp(v int64) *DescribeInterAuthStatisticsHistoryRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryRequest) SetStatisticalType(v string) *DescribeInterAuthStatisticsHistoryRequest {
	s.StatisticalType = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryRequest) SetZoneName(v string) *DescribeInterAuthStatisticsHistoryRequest {
	s.ZoneName = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryRequest) Validate() error {
	return dara.Validate(s)
}
