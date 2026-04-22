// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInterAuthStatisticsZoneOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOverviewPeriod(v string) *DescribeInterAuthStatisticsZoneOverviewRequest
	GetOverviewPeriod() *string
	SetServerRegion(v string) *DescribeInterAuthStatisticsZoneOverviewRequest
	GetServerRegion() *string
	SetZoneName(v string) *DescribeInterAuthStatisticsZoneOverviewRequest
	GetZoneName() *string
}

type DescribeInterAuthStatisticsZoneOverviewRequest struct {
	// example:
	//
	// DAY、WEEK、MONTH
	OverviewPeriod *string `json:"OverviewPeriod,omitempty" xml:"OverviewPeriod,omitempty"`
	// example:
	//
	// DescribeInterAuthStatisticsZoneOverview
	ServerRegion *string `json:"ServerRegion,omitempty" xml:"ServerRegion,omitempty"`
	// example:
	//
	// cheng.suow.cc
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeInterAuthStatisticsZoneOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsZoneOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsZoneOverviewRequest) GetOverviewPeriod() *string {
	return s.OverviewPeriod
}

func (s *DescribeInterAuthStatisticsZoneOverviewRequest) GetServerRegion() *string {
	return s.ServerRegion
}

func (s *DescribeInterAuthStatisticsZoneOverviewRequest) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeInterAuthStatisticsZoneOverviewRequest) SetOverviewPeriod(v string) *DescribeInterAuthStatisticsZoneOverviewRequest {
	s.OverviewPeriod = &v
	return s
}

func (s *DescribeInterAuthStatisticsZoneOverviewRequest) SetServerRegion(v string) *DescribeInterAuthStatisticsZoneOverviewRequest {
	s.ServerRegion = &v
	return s
}

func (s *DescribeInterAuthStatisticsZoneOverviewRequest) SetZoneName(v string) *DescribeInterAuthStatisticsZoneOverviewRequest {
	s.ZoneName = &v
	return s
}

func (s *DescribeInterAuthStatisticsZoneOverviewRequest) Validate() error {
	return dara.Validate(s)
}
