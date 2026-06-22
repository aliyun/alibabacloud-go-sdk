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
	// The vulnerability type of the Serverless asset. Valid values:
	//
	// - **sca**: middleware vulnerability
	//
	// - **app**: scanner vulnerability
	//
	// >Serverless assets currently support only application vulnerability scanning.
	//
	// example:
	//
	// sca,app
	Types *string `json:"Types,omitempty" xml:"Types,omitempty"`
	// The UUID of the asset instance to query.
	//
	// >You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to obtain this parameter.
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
