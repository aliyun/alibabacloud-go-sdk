// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *DescribeInstanceDetailsRequest
	GetInstanceIds() *string
	SetSourceIp(v string) *DescribeInstanceDetailsRequest
	GetSourceIp() *string
}

type DescribeInstanceDetailsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ["ddoscoo-cn-XXXX1", "ddoscoo-cn-XXXX2"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeInstanceDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeInstanceDetailsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInstanceDetailsRequest) SetInstanceIds(v string) *DescribeInstanceDetailsRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstanceDetailsRequest) SetSourceIp(v string) *DescribeInstanceDetailsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceDetailsRequest) Validate() error {
	return dara.Validate(s)
}
