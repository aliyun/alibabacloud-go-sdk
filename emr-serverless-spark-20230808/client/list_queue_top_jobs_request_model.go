// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueueTopJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListQueueTopJobsRequest
	GetRegionId() *string
	SetTopN(v int32) *ListQueueTopJobsRequest
	GetTopN() *int32
}

type ListQueueTopJobsRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The maximum number of jobs to return. The maximum value is 10.
	//
	// example:
	//
	// 10
	TopN *int32 `json:"topN,omitempty" xml:"topN,omitempty"`
}

func (s ListQueueTopJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQueueTopJobsRequest) GoString() string {
	return s.String()
}

func (s *ListQueueTopJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListQueueTopJobsRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *ListQueueTopJobsRequest) SetRegionId(v string) *ListQueueTopJobsRequest {
	s.RegionId = &v
	return s
}

func (s *ListQueueTopJobsRequest) SetTopN(v int32) *ListQueueTopJobsRequest {
	s.TopN = &v
	return s
}

func (s *ListQueueTopJobsRequest) Validate() error {
	return dara.Validate(s)
}
