// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserAppUsageAggregateByHourRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *QueryUserAppUsageAggregateByHourRequest
	GetAppInstanceGroupId() *string
	SetBizRegionId(v string) *QueryUserAppUsageAggregateByHourRequest
	GetBizRegionId() *string
	SetEndDate(v string) *QueryUserAppUsageAggregateByHourRequest
	GetEndDate() *string
	SetRegionId(v string) *QueryUserAppUsageAggregateByHourRequest
	GetRegionId() *string
	SetStartDate(v string) *QueryUserAppUsageAggregateByHourRequest
	GetStartDate() *string
	SetUserId(v string) *QueryUserAppUsageAggregateByHourRequest
	GetUserId() *string
}

type QueryUserAppUsageAggregateByHourRequest struct {
	// The ID of the app instance group, which is used to filter data for a specified app instance group. If this parameter is not provided, data is not filtered by app instance group.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The business region ID, which is used to filter application usage data. If this parameter is not provided, data is not filtered by business region. You can call [ListRegions](~~ListRegions~~) to obtain the value.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The end date for statistics in the yyyyMMdd format. If this parameter is not provided, it defaults to yesterday. The date cannot be earlier than StartDate or 90 days before the call date. The query range includes this date. The date is interpreted as a calendar date in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 20260907
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The region ID for OpenAPI access, which is used for request authentication and routing. To filter business data by region, use BizRegionId.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start date for statistics in the yyyyMMdd format. If this parameter is not provided, it defaults to yesterday. The date cannot be earlier than 90 days before the call date. The query range includes this date. The date is interpreted as a calendar date in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 20260901
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The ID of the end user to query. This parameter must be explicitly provided. If it is not provided, no matching records are usually returned.
	//
	// example:
	//
	// user01
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryUserAppUsageAggregateByHourRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserAppUsageAggregateByHourRequest) GoString() string {
	return s.String()
}

func (s *QueryUserAppUsageAggregateByHourRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *QueryUserAppUsageAggregateByHourRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *QueryUserAppUsageAggregateByHourRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryUserAppUsageAggregateByHourRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryUserAppUsageAggregateByHourRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryUserAppUsageAggregateByHourRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryUserAppUsageAggregateByHourRequest) SetAppInstanceGroupId(v string) *QueryUserAppUsageAggregateByHourRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *QueryUserAppUsageAggregateByHourRequest) SetBizRegionId(v string) *QueryUserAppUsageAggregateByHourRequest {
	s.BizRegionId = &v
	return s
}

func (s *QueryUserAppUsageAggregateByHourRequest) SetEndDate(v string) *QueryUserAppUsageAggregateByHourRequest {
	s.EndDate = &v
	return s
}

func (s *QueryUserAppUsageAggregateByHourRequest) SetRegionId(v string) *QueryUserAppUsageAggregateByHourRequest {
	s.RegionId = &v
	return s
}

func (s *QueryUserAppUsageAggregateByHourRequest) SetStartDate(v string) *QueryUserAppUsageAggregateByHourRequest {
	s.StartDate = &v
	return s
}

func (s *QueryUserAppUsageAggregateByHourRequest) SetUserId(v string) *QueryUserAppUsageAggregateByHourRequest {
	s.UserId = &v
	return s
}

func (s *QueryUserAppUsageAggregateByHourRequest) Validate() error {
	return dara.Validate(s)
}
