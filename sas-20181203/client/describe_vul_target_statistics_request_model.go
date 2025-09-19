// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulTargetStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *DescribeVulTargetStatisticsRequest
	GetType() *string
}

type DescribeVulTargetStatisticsRequest struct {
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: Linux software vulnerability
	//
	// 	- **sys**: Windows system vulnerability
	//
	// 	- **cms**: Web-CMS vulnerability
	//
	// 	- **emg**: urgent vulnerability
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVulTargetStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulTargetStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulTargetStatisticsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeVulTargetStatisticsRequest) SetType(v string) *DescribeVulTargetStatisticsRequest {
	s.Type = &v
	return s
}

func (s *DescribeVulTargetStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
