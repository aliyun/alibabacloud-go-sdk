// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulDefendCountStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVulType(v string) *DescribeVulDefendCountStatisticsRequest
	GetVulType() *string
}

type DescribeVulDefendCountStatisticsRequest struct {
	// The type of the vulnerabilities. Valid values:
	//
	// 	- app: application vulnerabilities
	//
	// 	- emg: urgent vulnerabilities
	//
	// example:
	//
	// emg
	VulType *string `json:"VulType,omitempty" xml:"VulType,omitempty"`
}

func (s DescribeVulDefendCountStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulDefendCountStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulDefendCountStatisticsRequest) GetVulType() *string {
	return s.VulType
}

func (s *DescribeVulDefendCountStatisticsRequest) SetVulType(v string) *DescribeVulDefendCountStatisticsRequest {
	s.VulType = &v
	return s
}

func (s *DescribeVulDefendCountStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
