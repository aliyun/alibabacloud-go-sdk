// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInterAuthStatisticsGlobalOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeInterAuthStatisticsGlobalOverviewRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *DescribeInterAuthStatisticsGlobalOverviewRequest
	GetClientToken() *string
	SetOverviewPeriod(v string) *DescribeInterAuthStatisticsGlobalOverviewRequest
	GetOverviewPeriod() *string
	SetServerRegion(v string) *DescribeInterAuthStatisticsGlobalOverviewRequest
	GetServerRegion() *string
}

type DescribeInterAuthStatisticsGlobalOverviewRequest struct {
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// 23445411234395894586....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// DAY
	OverviewPeriod *string `json:"OverviewPeriod,omitempty" xml:"OverviewPeriod,omitempty"`
	// example:
	//
	// cn-beijing
	ServerRegion *string `json:"ServerRegion,omitempty" xml:"ServerRegion,omitempty"`
}

func (s DescribeInterAuthStatisticsGlobalOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsGlobalOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsGlobalOverviewRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeInterAuthStatisticsGlobalOverviewRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeInterAuthStatisticsGlobalOverviewRequest) GetOverviewPeriod() *string {
	return s.OverviewPeriod
}

func (s *DescribeInterAuthStatisticsGlobalOverviewRequest) GetServerRegion() *string {
	return s.ServerRegion
}

func (s *DescribeInterAuthStatisticsGlobalOverviewRequest) SetAcceptLanguage(v string) *DescribeInterAuthStatisticsGlobalOverviewRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeInterAuthStatisticsGlobalOverviewRequest) SetClientToken(v string) *DescribeInterAuthStatisticsGlobalOverviewRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeInterAuthStatisticsGlobalOverviewRequest) SetOverviewPeriod(v string) *DescribeInterAuthStatisticsGlobalOverviewRequest {
	s.OverviewPeriod = &v
	return s
}

func (s *DescribeInterAuthStatisticsGlobalOverviewRequest) SetServerRegion(v string) *DescribeInterAuthStatisticsGlobalOverviewRequest {
	s.ServerRegion = &v
	return s
}

func (s *DescribeInterAuthStatisticsGlobalOverviewRequest) Validate() error {
	return dara.Validate(s)
}
