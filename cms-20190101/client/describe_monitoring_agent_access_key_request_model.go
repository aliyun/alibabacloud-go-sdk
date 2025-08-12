// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringAgentAccessKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeMonitoringAgentAccessKeyRequest
	GetRegionId() *string
}

type DescribeMonitoringAgentAccessKeyRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMonitoringAgentAccessKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentAccessKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentAccessKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMonitoringAgentAccessKeyRequest) SetRegionId(v string) *DescribeMonitoringAgentAccessKeyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMonitoringAgentAccessKeyRequest) Validate() error {
	return dara.Validate(s)
}
