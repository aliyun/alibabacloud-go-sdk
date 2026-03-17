// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserOnlineClientStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUserOnlineClientStatisticsResponseBody
	GetRequestId() *string
	SetUserStatistics(v *DescribeUserOnlineClientStatisticsResponseBodyUserStatistics) *DescribeUserOnlineClientStatisticsResponseBody
	GetUserStatistics() *DescribeUserOnlineClientStatisticsResponseBodyUserStatistics
}

type DescribeUserOnlineClientStatisticsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 14846A6A-2192-4F6A-B272-B8BD68EBC89B
	RequestId      *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserStatistics *DescribeUserOnlineClientStatisticsResponseBodyUserStatistics `json:"UserStatistics,omitempty" xml:"UserStatistics,omitempty" type:"Struct"`
}

func (s DescribeUserOnlineClientStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserOnlineClientStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserOnlineClientStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserOnlineClientStatisticsResponseBody) GetUserStatistics() *DescribeUserOnlineClientStatisticsResponseBodyUserStatistics {
	return s.UserStatistics
}

func (s *DescribeUserOnlineClientStatisticsResponseBody) SetRequestId(v string) *DescribeUserOnlineClientStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserOnlineClientStatisticsResponseBody) SetUserStatistics(v *DescribeUserOnlineClientStatisticsResponseBodyUserStatistics) *DescribeUserOnlineClientStatisticsResponseBody {
	s.UserStatistics = v
	return s
}

func (s *DescribeUserOnlineClientStatisticsResponseBody) Validate() error {
	if s.UserStatistics != nil {
		if err := s.UserStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUserOnlineClientStatisticsResponseBodyUserStatistics struct {
	Statistics []*DescribeUserOnlineClientStatisticsResponseBodyUserStatisticsStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s DescribeUserOnlineClientStatisticsResponseBodyUserStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserOnlineClientStatisticsResponseBodyUserStatistics) GoString() string {
	return s.String()
}

func (s *DescribeUserOnlineClientStatisticsResponseBodyUserStatistics) GetStatistics() []*DescribeUserOnlineClientStatisticsResponseBodyUserStatisticsStatistics {
	return s.Statistics
}

func (s *DescribeUserOnlineClientStatisticsResponseBodyUserStatistics) SetStatistics(v []*DescribeUserOnlineClientStatisticsResponseBodyUserStatisticsStatistics) *DescribeUserOnlineClientStatisticsResponseBodyUserStatistics {
	s.Statistics = v
	return s
}

func (s *DescribeUserOnlineClientStatisticsResponseBodyUserStatistics) Validate() error {
	if s.Statistics != nil {
		for _, item := range s.Statistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserOnlineClientStatisticsResponseBodyUserStatisticsStatistics struct {
	OnlineCount *string `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeUserOnlineClientStatisticsResponseBodyUserStatisticsStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserOnlineClientStatisticsResponseBodyUserStatisticsStatistics) GoString() string {
	return s.String()
}

func (s *DescribeUserOnlineClientStatisticsResponseBodyUserStatisticsStatistics) GetOnlineCount() *string {
	return s.OnlineCount
}

func (s *DescribeUserOnlineClientStatisticsResponseBodyUserStatisticsStatistics) GetUserName() *string {
	return s.UserName
}

func (s *DescribeUserOnlineClientStatisticsResponseBodyUserStatisticsStatistics) SetOnlineCount(v string) *DescribeUserOnlineClientStatisticsResponseBodyUserStatisticsStatistics {
	s.OnlineCount = &v
	return s
}

func (s *DescribeUserOnlineClientStatisticsResponseBodyUserStatisticsStatistics) SetUserName(v string) *DescribeUserOnlineClientStatisticsResponseBodyUserStatisticsStatistics {
	s.UserName = &v
	return s
}

func (s *DescribeUserOnlineClientStatisticsResponseBodyUserStatisticsStatistics) Validate() error {
	return dara.Validate(s)
}
