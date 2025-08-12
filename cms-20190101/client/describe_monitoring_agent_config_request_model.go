// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringAgentConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeMonitoringAgentConfigRequest
	GetRegionId() *string
}

type DescribeMonitoringAgentConfigRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMonitoringAgentConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMonitoringAgentConfigRequest) SetRegionId(v string) *DescribeMonitoringAgentConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMonitoringAgentConfigRequest) Validate() error {
	return dara.Validate(s)
}
