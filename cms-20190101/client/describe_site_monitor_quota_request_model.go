// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeSiteMonitorQuotaRequest
	GetRegionId() *string
}

type DescribeSiteMonitorQuotaRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSiteMonitorQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorQuotaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSiteMonitorQuotaRequest) SetRegionId(v string) *DescribeSiteMonitorQuotaRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSiteMonitorQuotaRequest) Validate() error {
	return dara.Validate(s)
}
