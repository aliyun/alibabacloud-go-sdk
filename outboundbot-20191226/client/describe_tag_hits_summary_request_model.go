// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagHitsSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeTagHitsSummaryRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *DescribeTagHitsSummaryRequest
	GetJobGroupId() *string
}

type DescribeTagHitsSummaryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 8fa1953f-4a84-46d8-b80c-8ce9cf684fb3
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8a4c6d3d-5ed6-44ca-b779-16c20f8862be
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
}

func (s DescribeTagHitsSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagHitsSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagHitsSummaryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTagHitsSummaryRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *DescribeTagHitsSummaryRequest) SetInstanceId(v string) *DescribeTagHitsSummaryRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTagHitsSummaryRequest) SetJobGroupId(v string) *DescribeTagHitsSummaryRequest {
	s.JobGroupId = &v
	return s
}

func (s *DescribeTagHitsSummaryRequest) Validate() error {
	return dara.Validate(s)
}
