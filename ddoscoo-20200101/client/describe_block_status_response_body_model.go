// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBlockStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeBlockStatusResponseBody
	GetRequestId() *string
	SetStatusList(v []*DescribeBlockStatusResponseBodyStatusList) *DescribeBlockStatusResponseBody
	GetStatusList() []*DescribeBlockStatusResponseBodyStatusList
}

type DescribeBlockStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of details of the Diversion from Origin Server configurations of the instance.
	StatusList []*DescribeBlockStatusResponseBodyStatusList `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s DescribeBlockStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBlockStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBlockStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBlockStatusResponseBody) GetStatusList() []*DescribeBlockStatusResponseBodyStatusList {
	return s.StatusList
}

func (s *DescribeBlockStatusResponseBody) SetRequestId(v string) *DescribeBlockStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBlockStatusResponseBody) SetStatusList(v []*DescribeBlockStatusResponseBodyStatusList) *DescribeBlockStatusResponseBody {
	s.StatusList = v
	return s
}

func (s *DescribeBlockStatusResponseBody) Validate() error {
	if s.StatusList != nil {
		for _, item := range s.StatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBlockStatusResponseBodyStatusList struct {
	// An array that consists of details of the Diversion from Origin Server configuration.
	BlockStatusList []*DescribeBlockStatusResponseBodyStatusListBlockStatusList `json:"BlockStatusList,omitempty" xml:"BlockStatusList,omitempty" type:"Repeated"`
	// The IP address of the instance.
	//
	// example:
	//
	// 203.XX.XX.88
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeBlockStatusResponseBodyStatusList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBlockStatusResponseBodyStatusList) GoString() string {
	return s.String()
}

func (s *DescribeBlockStatusResponseBodyStatusList) GetBlockStatusList() []*DescribeBlockStatusResponseBodyStatusListBlockStatusList {
	return s.BlockStatusList
}

func (s *DescribeBlockStatusResponseBodyStatusList) GetIp() *string {
	return s.Ip
}

func (s *DescribeBlockStatusResponseBodyStatusList) SetBlockStatusList(v []*DescribeBlockStatusResponseBodyStatusListBlockStatusList) *DescribeBlockStatusResponseBodyStatusList {
	s.BlockStatusList = v
	return s
}

func (s *DescribeBlockStatusResponseBodyStatusList) SetIp(v string) *DescribeBlockStatusResponseBodyStatusList {
	s.Ip = &v
	return s
}

func (s *DescribeBlockStatusResponseBodyStatusList) Validate() error {
	if s.BlockStatusList != nil {
		for _, item := range s.BlockStatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBlockStatusResponseBodyStatusListBlockStatusList struct {
	// The blocking status of the network traffic. Valid values:
	//
	// 	- **areablock**: Network traffic is blocked.
	//
	// 	- **normal**: Network traffic is not blocked.
	//
	// example:
	//
	// areablock
	BlockStatus *string `json:"BlockStatus,omitempty" xml:"BlockStatus,omitempty"`
	// The end time of the blocking. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1540196323
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The Internet service provider (ISP) line from which the traffic is blocked. Valid values:
	//
	// 	- **ct**: China Telecom (International)
	//
	// 	- **cut**: China Unicom (International)
	//
	// example:
	//
	// cut
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The start time of the blocking. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1540195323
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBlockStatusResponseBodyStatusListBlockStatusList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBlockStatusResponseBodyStatusListBlockStatusList) GoString() string {
	return s.String()
}

func (s *DescribeBlockStatusResponseBodyStatusListBlockStatusList) GetBlockStatus() *string {
	return s.BlockStatus
}

func (s *DescribeBlockStatusResponseBodyStatusListBlockStatusList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeBlockStatusResponseBodyStatusListBlockStatusList) GetLine() *string {
	return s.Line
}

func (s *DescribeBlockStatusResponseBodyStatusListBlockStatusList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeBlockStatusResponseBodyStatusListBlockStatusList) SetBlockStatus(v string) *DescribeBlockStatusResponseBodyStatusListBlockStatusList {
	s.BlockStatus = &v
	return s
}

func (s *DescribeBlockStatusResponseBodyStatusListBlockStatusList) SetEndTime(v int64) *DescribeBlockStatusResponseBodyStatusListBlockStatusList {
	s.EndTime = &v
	return s
}

func (s *DescribeBlockStatusResponseBodyStatusListBlockStatusList) SetLine(v string) *DescribeBlockStatusResponseBodyStatusListBlockStatusList {
	s.Line = &v
	return s
}

func (s *DescribeBlockStatusResponseBodyStatusListBlockStatusList) SetStartTime(v int64) *DescribeBlockStatusResponseBodyStatusListBlockStatusList {
	s.StartTime = &v
	return s
}

func (s *DescribeBlockStatusResponseBodyStatusListBlockStatusList) Validate() error {
	return dara.Validate(s)
}
