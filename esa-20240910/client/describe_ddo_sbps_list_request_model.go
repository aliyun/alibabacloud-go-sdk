// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSBpsListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoverage(v string) *DescribeDDoSBpsListRequest
	GetCoverage() *string
	SetEndTime(v string) *DescribeDDoSBpsListRequest
	GetEndTime() *string
	SetSiteId(v int64) *DescribeDDoSBpsListRequest
	GetSiteId() *int64
	SetStartTime(v string) *DescribeDDoSBpsListRequest
	GetStartTime() *string
}

type DescribeDDoSBpsListRequest struct {
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSBpsListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSBpsListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSBpsListRequest) GetCoverage() *string {
	return s.Coverage
}

func (s *DescribeDDoSBpsListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDDoSBpsListRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeDDoSBpsListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDDoSBpsListRequest) SetCoverage(v string) *DescribeDDoSBpsListRequest {
	s.Coverage = &v
	return s
}

func (s *DescribeDDoSBpsListRequest) SetEndTime(v string) *DescribeDDoSBpsListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSBpsListRequest) SetSiteId(v int64) *DescribeDDoSBpsListRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeDDoSBpsListRequest) SetStartTime(v string) *DescribeDDoSBpsListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSBpsListRequest) Validate() error {
	return dara.Validate(s)
}
