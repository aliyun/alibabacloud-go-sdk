// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceVulStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTypes(v string) *DescribeInstanceVulStatisticsRequest
	GetTypes() *string
	SetUuid(v string) *DescribeInstanceVulStatisticsRequest
	GetUuid() *string
}

type DescribeInstanceVulStatisticsRequest struct {
	// The vulnerability type of the serverless instance. Valid values:
	//
	// 	- **sca**: middleware vulnerabilities
	//
	// 	- **app**: application vulnerabilities detected by using a scanner
	//
	// >  Serverless instances allow you to detect only application vulnerabilities by using a scanner.
	//
	// example:
	//
	// sca,app
	Types *string `json:"Types,omitempty" xml:"Types,omitempty"`
	// The UUID of the instance to query.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5b268326-273e-44fc-a0e3-9482435c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeInstanceVulStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceVulStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceVulStatisticsRequest) GetTypes() *string {
	return s.Types
}

func (s *DescribeInstanceVulStatisticsRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeInstanceVulStatisticsRequest) SetTypes(v string) *DescribeInstanceVulStatisticsRequest {
	s.Types = &v
	return s
}

func (s *DescribeInstanceVulStatisticsRequest) SetUuid(v string) *DescribeInstanceVulStatisticsRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeInstanceVulStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
