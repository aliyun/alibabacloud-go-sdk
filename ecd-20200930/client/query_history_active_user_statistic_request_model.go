// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoryActiveUserStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v int32) *QueryHistoryActiveUserStatisticRequest
	GetBizType() *int32
	SetEndDate(v string) *QueryHistoryActiveUserStatisticRequest
	GetEndDate() *string
	SetOfficeSiteId(v string) *QueryHistoryActiveUserStatisticRequest
	GetOfficeSiteId() *string
	SetPeriod(v string) *QueryHistoryActiveUserStatisticRequest
	GetPeriod() *string
	SetStartDate(v string) *QueryHistoryActiveUserStatisticRequest
	GetStartDate() *string
	SetUserGroupId(v string) *QueryHistoryActiveUserStatisticRequest
	GetUserGroupId() *string
}

type QueryHistoryActiveUserStatisticRequest struct {
	// The business channel type code.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The end date of the query. The date is in the yyyy-MM-dd format. The maximum value is yesterday (N-1 data).
	//
	// example:
	//
	// 2024-12-31
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The workspace ID. If specified, only active users within the specified workspace are counted.
	//
	// example:
	//
	// cn-hangzhou+dir-467671****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The statistical period.
	//
	// example:
	//
	// day
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The start date of the query. The date is in the yyyy-MM-dd format. The value cannot be earlier than 6 months ago or later than EndDate.
	//
	// example:
	//
	// 2024-12-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The user group ID. If specified, only active users within the specified user group are counted.
	//
	// example:
	//
	// ug-12345678
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s QueryHistoryActiveUserStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryActiveUserStatisticRequest) GoString() string {
	return s.String()
}

func (s *QueryHistoryActiveUserStatisticRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *QueryHistoryActiveUserStatisticRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryHistoryActiveUserStatisticRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *QueryHistoryActiveUserStatisticRequest) GetPeriod() *string {
	return s.Period
}

func (s *QueryHistoryActiveUserStatisticRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryHistoryActiveUserStatisticRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *QueryHistoryActiveUserStatisticRequest) SetBizType(v int32) *QueryHistoryActiveUserStatisticRequest {
	s.BizType = &v
	return s
}

func (s *QueryHistoryActiveUserStatisticRequest) SetEndDate(v string) *QueryHistoryActiveUserStatisticRequest {
	s.EndDate = &v
	return s
}

func (s *QueryHistoryActiveUserStatisticRequest) SetOfficeSiteId(v string) *QueryHistoryActiveUserStatisticRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *QueryHistoryActiveUserStatisticRequest) SetPeriod(v string) *QueryHistoryActiveUserStatisticRequest {
	s.Period = &v
	return s
}

func (s *QueryHistoryActiveUserStatisticRequest) SetStartDate(v string) *QueryHistoryActiveUserStatisticRequest {
	s.StartDate = &v
	return s
}

func (s *QueryHistoryActiveUserStatisticRequest) SetUserGroupId(v string) *QueryHistoryActiveUserStatisticRequest {
	s.UserGroupId = &v
	return s
}

func (s *QueryHistoryActiveUserStatisticRequest) Validate() error {
	return dara.Validate(s)
}
