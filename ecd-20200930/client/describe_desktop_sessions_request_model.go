// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckOsSession(v bool) *DescribeDesktopSessionsRequest
	GetCheckOsSession() *bool
	SetDesktopId(v []*string) *DescribeDesktopSessionsRequest
	GetDesktopId() []*string
	SetDesktopName(v string) *DescribeDesktopSessionsRequest
	GetDesktopName() *string
	SetEndTime(v string) *DescribeDesktopSessionsRequest
	GetEndTime() *string
	SetEndUserId(v string) *DescribeDesktopSessionsRequest
	GetEndUserId() *string
	SetEndUserIdFilter(v string) *DescribeDesktopSessionsRequest
	GetEndUserIdFilter() *string
	SetEndUserIds(v []*string) *DescribeDesktopSessionsRequest
	GetEndUserIds() []*string
	SetFillHardwareInfo(v bool) *DescribeDesktopSessionsRequest
	GetFillHardwareInfo() *bool
	SetLanguage(v string) *DescribeDesktopSessionsRequest
	GetLanguage() *string
	SetOfficeSiteId(v string) *DescribeDesktopSessionsRequest
	GetOfficeSiteId() *string
	SetPageNumber(v int32) *DescribeDesktopSessionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDesktopSessionsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDesktopSessionsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDesktopSessionsRequest
	GetResourceGroupId() *string
	SetSessionStatus(v string) *DescribeDesktopSessionsRequest
	GetSessionStatus() *string
	SetStartTime(v string) *DescribeDesktopSessionsRequest
	GetStartTime() *string
	SetSubPayType(v string) *DescribeDesktopSessionsRequest
	GetSubPayType() *string
}

type DescribeDesktopSessionsRequest struct {
	// Specifies whether to check the session status within the cloud computer.
	//
	// example:
	//
	// true
	CheckOsSession *bool `json:"CheckOsSession,omitempty" xml:"CheckOsSession,omitempty"`
	// The cloud computer IDs. You can specify 1 to 100 IDs.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The cloud computer name.
	//
	// example:
	//
	// DemoComputer
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The end time of the query.
	//
	// example:
	//
	// 2023-02-13T02:51:43Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The end user ID.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The end user ID. This parameter is the same as EndUserId. You only need to specify one of them.
	//
	// example:
	//
	// alice
	EndUserIdFilter *string   `json:"EndUserIdFilter,omitempty" xml:"EndUserIdFilter,omitempty"`
	EndUserIds      []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// Specifies whether to include terminal information in the response.
	FillHardwareInfo *bool `json:"FillHardwareInfo,omitempty" xml:"FillHardwareInfo,omitempty"`
	// The language type of the response.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The page number of the current page in a paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of rows per page in a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The session connection status.
	//
	// example:
	//
	// Connected
	SessionStatus *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	// The start time of the query.
	//
	// example:
	//
	// 2023-01-28T02:31:43Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The purchase method of the cloud computer.
	//
	// example:
	//
	// monthPackage
	SubPayType *string `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
}

func (s DescribeDesktopSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopSessionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopSessionsRequest) GetCheckOsSession() *bool {
	return s.CheckOsSession
}

func (s *DescribeDesktopSessionsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *DescribeDesktopSessionsRequest) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeDesktopSessionsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDesktopSessionsRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeDesktopSessionsRequest) GetEndUserIdFilter() *string {
	return s.EndUserIdFilter
}

func (s *DescribeDesktopSessionsRequest) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *DescribeDesktopSessionsRequest) GetFillHardwareInfo() *bool {
	return s.FillHardwareInfo
}

func (s *DescribeDesktopSessionsRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeDesktopSessionsRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeDesktopSessionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDesktopSessionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDesktopSessionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDesktopSessionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDesktopSessionsRequest) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *DescribeDesktopSessionsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDesktopSessionsRequest) GetSubPayType() *string {
	return s.SubPayType
}

func (s *DescribeDesktopSessionsRequest) SetCheckOsSession(v bool) *DescribeDesktopSessionsRequest {
	s.CheckOsSession = &v
	return s
}

func (s *DescribeDesktopSessionsRequest) SetDesktopId(v []*string) *DescribeDesktopSessionsRequest {
	s.DesktopId = v
	return s
}

func (s *DescribeDesktopSessionsRequest) SetDesktopName(v string) *DescribeDesktopSessionsRequest {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopSessionsRequest) SetEndTime(v string) *DescribeDesktopSessionsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDesktopSessionsRequest) SetEndUserId(v string) *DescribeDesktopSessionsRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopSessionsRequest) SetEndUserIdFilter(v string) *DescribeDesktopSessionsRequest {
	s.EndUserIdFilter = &v
	return s
}

func (s *DescribeDesktopSessionsRequest) SetEndUserIds(v []*string) *DescribeDesktopSessionsRequest {
	s.EndUserIds = v
	return s
}

func (s *DescribeDesktopSessionsRequest) SetFillHardwareInfo(v bool) *DescribeDesktopSessionsRequest {
	s.FillHardwareInfo = &v
	return s
}

func (s *DescribeDesktopSessionsRequest) SetLanguage(v string) *DescribeDesktopSessionsRequest {
	s.Language = &v
	return s
}

func (s *DescribeDesktopSessionsRequest) SetOfficeSiteId(v string) *DescribeDesktopSessionsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopSessionsRequest) SetPageNumber(v int32) *DescribeDesktopSessionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDesktopSessionsRequest) SetPageSize(v int32) *DescribeDesktopSessionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDesktopSessionsRequest) SetRegionId(v string) *DescribeDesktopSessionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopSessionsRequest) SetResourceGroupId(v string) *DescribeDesktopSessionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDesktopSessionsRequest) SetSessionStatus(v string) *DescribeDesktopSessionsRequest {
	s.SessionStatus = &v
	return s
}

func (s *DescribeDesktopSessionsRequest) SetStartTime(v string) *DescribeDesktopSessionsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDesktopSessionsRequest) SetSubPayType(v string) *DescribeDesktopSessionsRequest {
	s.SubPayType = &v
	return s
}

func (s *DescribeDesktopSessionsRequest) Validate() error {
	return dara.Validate(s)
}
