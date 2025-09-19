// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterVulStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterVulStatisticsRequest
	GetClusterId() *string
	SetTypes(v string) *DescribeClusterVulStatisticsRequest
	GetTypes() *string
}

type DescribeClusterVulStatisticsRequest struct {
	// The ID of the container cluster.
	//
	// example:
	//
	// c471f0f61b9c04f8380556e922cf1****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the vulnerabilities. Valid values:
	//
	// 	- **cve**: Linux software vulnerabilities
	//
	// 	- **app**: application vulnerabilities
	//
	// 	- **sca**: vulnerabilities that are detected based on software component analysis
	//
	// example:
	//
	// cve,app,sca
	Types *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s DescribeClusterVulStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterVulStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterVulStatisticsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterVulStatisticsRequest) GetTypes() *string {
	return s.Types
}

func (s *DescribeClusterVulStatisticsRequest) SetClusterId(v string) *DescribeClusterVulStatisticsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterVulStatisticsRequest) SetTypes(v string) *DescribeClusterVulStatisticsRequest {
	s.Types = &v
	return s
}

func (s *DescribeClusterVulStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
