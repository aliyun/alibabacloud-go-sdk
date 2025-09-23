// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSpecsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *DescribeInstanceSpecsRequest
	GetInstanceIds() *string
	SetSourceIp(v string) *DescribeInstanceSpecsRequest
	GetSourceIp() *string
}

type DescribeInstanceSpecsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ["ddoscoo-cn-XXXXX"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeInstanceSpecsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSpecsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeInstanceSpecsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInstanceSpecsRequest) SetInstanceIds(v string) *DescribeInstanceSpecsRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstanceSpecsRequest) SetSourceIp(v string) *DescribeInstanceSpecsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceSpecsRequest) Validate() error {
	return dara.Validate(s)
}
