// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserAppUsageTimeRankRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *QueryUserAppUsageTimeRankRequest
	GetAppInstanceGroupId() *string
	SetBizRegionId(v string) *QueryUserAppUsageTimeRankRequest
	GetBizRegionId() *string
	SetEndDate(v string) *QueryUserAppUsageTimeRankRequest
	GetEndDate() *string
	SetRegionId(v string) *QueryUserAppUsageTimeRankRequest
	GetRegionId() *string
	SetStartDate(v string) *QueryUserAppUsageTimeRankRequest
	GetStartDate() *string
	SetUserId(v string) *QueryUserAppUsageTimeRankRequest
	GetUserId() *string
}

type QueryUserAppUsageTimeRankRequest struct {
	// The app instance group ID, which is used to filter data for a specified app instance group. If this parameter is not provided, data is not filtered by app instance group.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The business region ID, which is used to filter application usage data. If this parameter is not provided, data is not filtered by business region. You can call the [ListRegions](~~ListRegions~~) operation to obtain the ID.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The end date of the statistical period, in `yyyyMMdd` format. If this parameter is not provided, it defaults to yesterday. The date cannot be earlier than `StartDate`, and cannot be earlier than 90 days before the call date. The query range includes this date. The date is interpreted as a calendar date in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 20260907
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The region ID for OpenAPI access, which is used for request authentication and routing. To filter business data by region, use `BizRegionId`.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start date of the statistical period, in `yyyyMMdd` format. If this parameter is not provided, it defaults to yesterday. The date cannot be earlier than 90 days before the call date. The query range includes this date. The date is interpreted as a calendar date in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 20260901
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The ID of the end user to be queried. This parameter must be explicitly provided. If it is not provided, no matching records are usually returned.
	//
	// example:
	//
	// user01
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryUserAppUsageTimeRankRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserAppUsageTimeRankRequest) GoString() string {
	return s.String()
}

func (s *QueryUserAppUsageTimeRankRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *QueryUserAppUsageTimeRankRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *QueryUserAppUsageTimeRankRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryUserAppUsageTimeRankRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryUserAppUsageTimeRankRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryUserAppUsageTimeRankRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryUserAppUsageTimeRankRequest) SetAppInstanceGroupId(v string) *QueryUserAppUsageTimeRankRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *QueryUserAppUsageTimeRankRequest) SetBizRegionId(v string) *QueryUserAppUsageTimeRankRequest {
	s.BizRegionId = &v
	return s
}

func (s *QueryUserAppUsageTimeRankRequest) SetEndDate(v string) *QueryUserAppUsageTimeRankRequest {
	s.EndDate = &v
	return s
}

func (s *QueryUserAppUsageTimeRankRequest) SetRegionId(v string) *QueryUserAppUsageTimeRankRequest {
	s.RegionId = &v
	return s
}

func (s *QueryUserAppUsageTimeRankRequest) SetStartDate(v string) *QueryUserAppUsageTimeRankRequest {
	s.StartDate = &v
	return s
}

func (s *QueryUserAppUsageTimeRankRequest) SetUserId(v string) *QueryUserAppUsageTimeRankRequest {
	s.UserId = &v
	return s
}

func (s *QueryUserAppUsageTimeRankRequest) Validate() error {
	return dara.Validate(s)
}
