// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardBaseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeDashboardBaseInfoRequest
	GetRegionId() *string
}

type DescribeDashboardBaseInfoRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDashboardBaseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardBaseInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDashboardBaseInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDashboardBaseInfoRequest) SetRegionId(v string) *DescribeDashboardBaseInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDashboardBaseInfoRequest) Validate() error {
	return dara.Validate(s)
}
