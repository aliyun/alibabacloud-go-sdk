// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJVSInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DescribeJVSInstanceRequest
	GetInstanceIds() []*string
	SetMaxResults(v int32) *DescribeJVSInstanceRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeJVSInstanceRequest
	GetNextToken() *string
}

type DescribeJVSInstanceRequest struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// 分页大小，最大值100，默认值10
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 分页游标，首次查询无需传入
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeJVSInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJVSInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeJVSInstanceRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeJVSInstanceRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeJVSInstanceRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeJVSInstanceRequest) SetInstanceIds(v []*string) *DescribeJVSInstanceRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeJVSInstanceRequest) SetMaxResults(v int32) *DescribeJVSInstanceRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeJVSInstanceRequest) SetNextToken(v string) *DescribeJVSInstanceRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeJVSInstanceRequest) Validate() error {
	return dara.Validate(s)
}
