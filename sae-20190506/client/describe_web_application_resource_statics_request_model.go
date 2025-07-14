// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebApplicationResourceStaticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeWebApplicationResourceStaticsRequest
	GetEndTime() *int64
	SetNamespaceId(v string) *DescribeWebApplicationResourceStaticsRequest
	GetNamespaceId() *string
	SetRegionId(v string) *DescribeWebApplicationResourceStaticsRequest
	GetRegionId() *string
	SetStartTime(v int64) *DescribeWebApplicationResourceStaticsRequest
	GetStartTime() *int64
}

type DescribeWebApplicationResourceStaticsRequest struct {
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 1687832980387
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 1562831689704
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeWebApplicationResourceStaticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebApplicationResourceStaticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebApplicationResourceStaticsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeWebApplicationResourceStaticsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeWebApplicationResourceStaticsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeWebApplicationResourceStaticsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeWebApplicationResourceStaticsRequest) SetEndTime(v int64) *DescribeWebApplicationResourceStaticsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeWebApplicationResourceStaticsRequest) SetNamespaceId(v string) *DescribeWebApplicationResourceStaticsRequest {
	s.NamespaceId = &v
	return s
}

func (s *DescribeWebApplicationResourceStaticsRequest) SetRegionId(v string) *DescribeWebApplicationResourceStaticsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeWebApplicationResourceStaticsRequest) SetStartTime(v int64) *DescribeWebApplicationResourceStaticsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeWebApplicationResourceStaticsRequest) Validate() error {
	return dara.Validate(s)
}
