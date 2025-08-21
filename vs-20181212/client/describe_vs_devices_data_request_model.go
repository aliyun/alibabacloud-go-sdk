// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDevicesDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeVsDevicesDataRequest
	GetEndTime() *string
	SetGroupId(v string) *DescribeVsDevicesDataRequest
	GetGroupId() *string
	SetOwnerId(v int64) *DescribeVsDevicesDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVsDevicesDataRequest
	GetStartTime() *string
}

type DescribeVsDevicesDataRequest struct {
	// example:
	//
	// 2022-01-30T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 348*****174-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 2022-01-04T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDevicesDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDevicesDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDevicesDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDevicesDataRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeVsDevicesDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsDevicesDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDevicesDataRequest) SetEndTime(v string) *DescribeVsDevicesDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDevicesDataRequest) SetGroupId(v string) *DescribeVsDevicesDataRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeVsDevicesDataRequest) SetOwnerId(v int64) *DescribeVsDevicesDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDevicesDataRequest) SetStartTime(v string) *DescribeVsDevicesDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDevicesDataRequest) Validate() error {
	return dara.Validate(s)
}
