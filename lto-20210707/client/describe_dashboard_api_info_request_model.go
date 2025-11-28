// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardApiInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeDashboardApiInfoRequest
	GetRegionId() *string
}

type DescribeDashboardApiInfoRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDashboardApiInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardApiInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDashboardApiInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDashboardApiInfoRequest) SetRegionId(v string) *DescribeDashboardApiInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDashboardApiInfoRequest) Validate() error {
	return dara.Validate(s)
}
