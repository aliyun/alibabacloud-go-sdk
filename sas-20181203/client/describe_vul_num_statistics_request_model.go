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
	SetResourceDirectoryAccountId(v int64) *DescribeVulNumStatisticsRequest
	GetResourceDirectoryAccountId() *int64
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
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the IDs.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
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

func (s *DescribeVulNumStatisticsRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeVulNumStatisticsRequest) SetFrom(v string) *DescribeVulNumStatisticsRequest {
	s.From = &v
	return s
}

func (s *DescribeVulNumStatisticsRequest) SetResourceDirectoryAccountId(v int64) *DescribeVulNumStatisticsRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeVulNumStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
