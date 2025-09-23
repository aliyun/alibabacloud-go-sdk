// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *DescribeInstanceStatisticsRequest
	GetInstanceIds() *string
	SetSourceIp(v string) *DescribeInstanceStatisticsRequest
	GetSourceIp() *string
}

type DescribeInstanceStatisticsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"InstanceId":"ddoscoo-cn-XXXXX","InstanceId":"ddoscoo-cn-YYYYY"}]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeInstanceStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeInstanceStatisticsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInstanceStatisticsRequest) SetInstanceIds(v string) *DescribeInstanceStatisticsRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstanceStatisticsRequest) SetSourceIp(v string) *DescribeInstanceStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
