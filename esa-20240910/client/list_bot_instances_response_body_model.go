// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBotInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceInfo(v []*ListBotInstancesResponseBodyInstanceInfo) *ListBotInstancesResponseBody
	GetInstanceInfo() []*ListBotInstancesResponseBodyInstanceInfo
	SetPageNumber(v int32) *ListBotInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBotInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListBotInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListBotInstancesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListBotInstancesResponseBody
	GetTotalPage() *int32
}

type ListBotInstancesResponseBody struct {
	// The instances that match the specified conditions under the current account.
	InstanceInfo []*ListBotInstancesResponseBodyInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Repeated"`
	// The current page number, which is the same as the PageNumber request parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 509FD5AF-AB5B-55A9-9568-38D98668E3AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 0
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListBotInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBotInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBotInstancesResponseBody) GetInstanceInfo() []*ListBotInstancesResponseBodyInstanceInfo {
	return s.InstanceInfo
}

func (s *ListBotInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBotInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBotInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBotInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBotInstancesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListBotInstancesResponseBody) SetInstanceInfo(v []*ListBotInstancesResponseBodyInstanceInfo) *ListBotInstancesResponseBody {
	s.InstanceInfo = v
	return s
}

func (s *ListBotInstancesResponseBody) SetPageNumber(v int32) *ListBotInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBotInstancesResponseBody) SetPageSize(v int32) *ListBotInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBotInstancesResponseBody) SetRequestId(v string) *ListBotInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBotInstancesResponseBody) SetTotalCount(v int32) *ListBotInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBotInstancesResponseBody) SetTotalPage(v int32) *ListBotInstancesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListBotInstancesResponseBody) Validate() error {
	if s.InstanceInfo != nil {
		for _, item := range s.InstanceInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBotInstancesResponseBodyInstanceInfo struct {
	// The Bot protection instance level. If this parameter is empty, the plan does not include a Bot protection instance. If a value is returned, the plan includes a Bot protection instance. Valid values:
	//
	// - enterprise_bot: web edition.
	//
	// - enterprise_bot_with_app: app edition.
	//
	// example:
	//
	// enterprise_bot
	BotInstanceLevel *string `json:"BotInstanceLevel,omitempty" xml:"BotInstanceLevel,omitempty"`
	// The time when the instance was purchased. The time is in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-04-12T05:41:51Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// sp-xcdn-96wblslz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The scheduled release time. The time is in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2026-03-25T16:00:00Z
	ReserveReleaseTime *string `json:"ReserveReleaseTime,omitempty" xml:"ReserveReleaseTime,omitempty"`
	// The ID of the associated site plan instance.
	//
	// example:
	//
	// esa-site-b0s6kmx0r0n4
	SiteInstanceId *string `json:"SiteInstanceId,omitempty" xml:"SiteInstanceId,omitempty"`
	// The instance status. Valid values:
	//
	// - **online**: The instance is running normally.
	//
	// - **offline**: The instance has expired but has not exceeded the retention period and is unavailable.
	//
	// - **disable**: The instance has been released.
	//
	// - **overdue**: The instance has been stopped due to an overdue payment.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListBotInstancesResponseBodyInstanceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListBotInstancesResponseBodyInstanceInfo) GoString() string {
	return s.String()
}

func (s *ListBotInstancesResponseBodyInstanceInfo) GetBotInstanceLevel() *string {
	return s.BotInstanceLevel
}

func (s *ListBotInstancesResponseBodyInstanceInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListBotInstancesResponseBodyInstanceInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBotInstancesResponseBodyInstanceInfo) GetReserveReleaseTime() *string {
	return s.ReserveReleaseTime
}

func (s *ListBotInstancesResponseBodyInstanceInfo) GetSiteInstanceId() *string {
	return s.SiteInstanceId
}

func (s *ListBotInstancesResponseBodyInstanceInfo) GetStatus() *string {
	return s.Status
}

func (s *ListBotInstancesResponseBodyInstanceInfo) SetBotInstanceLevel(v string) *ListBotInstancesResponseBodyInstanceInfo {
	s.BotInstanceLevel = &v
	return s
}

func (s *ListBotInstancesResponseBodyInstanceInfo) SetCreateTime(v string) *ListBotInstancesResponseBodyInstanceInfo {
	s.CreateTime = &v
	return s
}

func (s *ListBotInstancesResponseBodyInstanceInfo) SetInstanceId(v string) *ListBotInstancesResponseBodyInstanceInfo {
	s.InstanceId = &v
	return s
}

func (s *ListBotInstancesResponseBodyInstanceInfo) SetReserveReleaseTime(v string) *ListBotInstancesResponseBodyInstanceInfo {
	s.ReserveReleaseTime = &v
	return s
}

func (s *ListBotInstancesResponseBodyInstanceInfo) SetSiteInstanceId(v string) *ListBotInstancesResponseBodyInstanceInfo {
	s.SiteInstanceId = &v
	return s
}

func (s *ListBotInstancesResponseBodyInstanceInfo) SetStatus(v string) *ListBotInstancesResponseBodyInstanceInfo {
	s.Status = &v
	return s
}

func (s *ListBotInstancesResponseBodyInstanceInfo) Validate() error {
	return dara.Validate(s)
}
