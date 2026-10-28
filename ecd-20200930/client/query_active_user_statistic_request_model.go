// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryActiveUserStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v int32) *QueryActiveUserStatisticRequest
	GetBizType() *int32
	SetEndTime(v string) *QueryActiveUserStatisticRequest
	GetEndTime() *string
	SetOfficeSiteId(v string) *QueryActiveUserStatisticRequest
	GetOfficeSiteId() *string
	SetPeriod(v string) *QueryActiveUserStatisticRequest
	GetPeriod() *string
	SetStartTime(v string) *QueryActiveUserStatisticRequest
	GetStartTime() *string
}

type QueryActiveUserStatisticRequest struct {
	// The business channel type code. Valid values:
	//
	// - 1: Enterprise Edition. This is the default value.
	//
	// - 3: Cloud Office.
	//
	// - 10: Standard Edition.
	//
	// - 20: Business Edition.
	//
	// - 30: Education Business Edition.
	//
	// - 40: Cloud phone isolated resources.
	//
	// - 50: AgentBay.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The end time of the query. The format is the same as that of StartTime. If the specified time is later than the current time, it is automatically truncated to the current time.
	//
	// example:
	//
	// 2020-12-01T06:32:31Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The workspace ID. If specified, only active users of desktops in this workspace are counted.
	//
	// example:
	//
	// cn-hangzhou+dir-885351****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The aggregation granularity of the statistics. Valid values:
	//
	// - ONE_MINUTE: 1 minute.
	//
	// - TWO_MINUTE: 2 minutes.
	//
	// - FIVE_MINUTE: 5 minutes. This is the default value.
	//
	// - ONE_HOUR: 1 hour.
	//
	// - ONE_DAY: 1 day.
	//
	// example:
	//
	// FIVE_MINUTE
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The start time of the query. Two formats are supported:
	//
	// - UTC format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// - Standard format: yyyy-MM-dd HH:mm:ss.
	//
	// The start time cannot be earlier than six months ago or later than the end time.
	//
	// example:
	//
	// 2020-11-30T06:32:31Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryActiveUserStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryActiveUserStatisticRequest) GoString() string {
	return s.String()
}

func (s *QueryActiveUserStatisticRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *QueryActiveUserStatisticRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryActiveUserStatisticRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *QueryActiveUserStatisticRequest) GetPeriod() *string {
	return s.Period
}

func (s *QueryActiveUserStatisticRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryActiveUserStatisticRequest) SetBizType(v int32) *QueryActiveUserStatisticRequest {
	s.BizType = &v
	return s
}

func (s *QueryActiveUserStatisticRequest) SetEndTime(v string) *QueryActiveUserStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *QueryActiveUserStatisticRequest) SetOfficeSiteId(v string) *QueryActiveUserStatisticRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *QueryActiveUserStatisticRequest) SetPeriod(v string) *QueryActiveUserStatisticRequest {
	s.Period = &v
	return s
}

func (s *QueryActiveUserStatisticRequest) SetStartTime(v string) *QueryActiveUserStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *QueryActiveUserStatisticRequest) Validate() error {
	return dara.Validate(s)
}
