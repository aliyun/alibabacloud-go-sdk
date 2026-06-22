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
	// The request source.
	//
	// 	- When querying Security Center data, set this parameter to **sas**.
	//
	// 	- When querying Server Guard data, you do not need to set this parameter.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The member account ID in the resource directory (Alibaba Cloud account).
	//
	// >You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
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
