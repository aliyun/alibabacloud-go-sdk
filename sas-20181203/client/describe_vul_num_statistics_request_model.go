// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulNumStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeVulNumStatisticsRequest
	GetFrom() *string
}

type DescribeVulNumStatisticsRequest struct {
	// The source of the request.
	//
	// 	- If you want to query Security Center-related data, set the value to **sas**.
	//
	// 	- If you want to query Server Guard-related data, you do not need to specify this parameter.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
}

func (s DescribeVulNumStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulNumStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulNumStatisticsRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeVulNumStatisticsRequest) SetFrom(v string) *DescribeVulNumStatisticsRequest {
	s.From = &v
	return s
}

func (s *DescribeVulNumStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
