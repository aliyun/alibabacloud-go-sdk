// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSBpsMaxRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoverage(v string) *DescribeDDoSBpsMaxRequest
	GetCoverage() *string
	SetEndTime(v string) *DescribeDDoSBpsMaxRequest
	GetEndTime() *string
	SetSiteId(v int64) *DescribeDDoSBpsMaxRequest
	GetSiteId() *int64
	SetStartTime(v string) *DescribeDDoSBpsMaxRequest
	GetStartTime() *string
}

type DescribeDDoSBpsMaxRequest struct {
	// example:
	//
	// global
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// example:
	//
	// 2023-04-07T02:34:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 709662109****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-14T17:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSBpsMaxRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSBpsMaxRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSBpsMaxRequest) GetCoverage() *string {
	return s.Coverage
}

func (s *DescribeDDoSBpsMaxRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDDoSBpsMaxRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeDDoSBpsMaxRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDDoSBpsMaxRequest) SetCoverage(v string) *DescribeDDoSBpsMaxRequest {
	s.Coverage = &v
	return s
}

func (s *DescribeDDoSBpsMaxRequest) SetEndTime(v string) *DescribeDDoSBpsMaxRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSBpsMaxRequest) SetSiteId(v int64) *DescribeDDoSBpsMaxRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeDDoSBpsMaxRequest) SetStartTime(v string) *DescribeDDoSBpsMaxRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSBpsMaxRequest) Validate() error {
	return dara.Validate(s)
}
