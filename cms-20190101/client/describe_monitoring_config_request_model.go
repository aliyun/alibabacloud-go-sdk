// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeMonitoringConfigRequest
	GetRegionId() *string
}

type DescribeMonitoringConfigRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMonitoringConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMonitoringConfigRequest) SetRegionId(v string) *DescribeMonitoringConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMonitoringConfigRequest) Validate() error {
	return dara.Validate(s)
}
