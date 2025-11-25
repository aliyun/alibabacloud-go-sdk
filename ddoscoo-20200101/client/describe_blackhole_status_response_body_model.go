// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBlackholeStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBlackholeStatus(v []*DescribeBlackholeStatusResponseBodyBlackholeStatus) *DescribeBlackholeStatusResponseBody
	GetBlackholeStatus() []*DescribeBlackholeStatusResponseBodyBlackholeStatus
	SetRequestId(v string) *DescribeBlackholeStatusResponseBody
	GetRequestId() *string
}

type DescribeBlackholeStatusResponseBody struct {
	// An array that consists of the blackhole filtering status of the instance.
	BlackholeStatus []*DescribeBlackholeStatusResponseBodyBlackholeStatus `json:"BlackholeStatus,omitempty" xml:"BlackholeStatus,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBlackholeStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBlackholeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBlackholeStatusResponseBody) GetBlackholeStatus() []*DescribeBlackholeStatusResponseBodyBlackholeStatus {
	return s.BlackholeStatus
}

func (s *DescribeBlackholeStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBlackholeStatusResponseBody) SetBlackholeStatus(v []*DescribeBlackholeStatusResponseBodyBlackholeStatus) *DescribeBlackholeStatusResponseBody {
	s.BlackholeStatus = v
	return s
}

func (s *DescribeBlackholeStatusResponseBody) SetRequestId(v string) *DescribeBlackholeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBlackholeStatusResponseBody) Validate() error {
	if s.BlackholeStatus != nil {
		for _, item := range s.BlackholeStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBlackholeStatusResponseBodyBlackholeStatus struct {
	// Indicates whether blackhole filtering is triggered for the instance. Valid values:
	//
	// 	- **blackhole**: yes
	//
	// 	- **normal**: no
	//
	// example:
	//
	// blackhole
	BlackStatus *string `json:"BlackStatus,omitempty" xml:"BlackStatus,omitempty"`
	// The end time of blackhole filtering. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1540196323
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 203.***.***.132
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The start time of blackhole filtering. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1540195323
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBlackholeStatusResponseBodyBlackholeStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeBlackholeStatusResponseBodyBlackholeStatus) GoString() string {
	return s.String()
}

func (s *DescribeBlackholeStatusResponseBodyBlackholeStatus) GetBlackStatus() *string {
	return s.BlackStatus
}

func (s *DescribeBlackholeStatusResponseBodyBlackholeStatus) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeBlackholeStatusResponseBodyBlackholeStatus) GetIp() *string {
	return s.Ip
}

func (s *DescribeBlackholeStatusResponseBodyBlackholeStatus) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeBlackholeStatusResponseBodyBlackholeStatus) SetBlackStatus(v string) *DescribeBlackholeStatusResponseBodyBlackholeStatus {
	s.BlackStatus = &v
	return s
}

func (s *DescribeBlackholeStatusResponseBodyBlackholeStatus) SetEndTime(v int64) *DescribeBlackholeStatusResponseBodyBlackholeStatus {
	s.EndTime = &v
	return s
}

func (s *DescribeBlackholeStatusResponseBodyBlackholeStatus) SetIp(v string) *DescribeBlackholeStatusResponseBodyBlackholeStatus {
	s.Ip = &v
	return s
}

func (s *DescribeBlackholeStatusResponseBodyBlackholeStatus) SetStartTime(v int64) *DescribeBlackholeStatusResponseBodyBlackholeStatus {
	s.StartTime = &v
	return s
}

func (s *DescribeBlackholeStatusResponseBodyBlackholeStatus) Validate() error {
	return dara.Validate(s)
}
