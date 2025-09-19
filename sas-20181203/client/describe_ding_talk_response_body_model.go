// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDingTalkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActionList(v []*DescribeDingTalkResponseBodyActionList) *DescribeDingTalkResponseBody
	GetActionList() []*DescribeDingTalkResponseBodyActionList
	SetPageInfo(v *DescribeDingTalkResponseBodyPageInfo) *DescribeDingTalkResponseBody
	GetPageInfo() *DescribeDingTalkResponseBodyPageInfo
	SetRequestId(v string) *DescribeDingTalkResponseBody
	GetRequestId() *string
}

type DescribeDingTalkResponseBody struct {
	// An array that consists of details of notifications.
	ActionList []*DescribeDingTalkResponseBodyActionList `json:"ActionList,omitempty" xml:"ActionList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeDingTalkResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B256A525-7E42-4BB9-A27C-9017FDDFF1A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDingTalkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDingTalkResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDingTalkResponseBody) GetActionList() []*DescribeDingTalkResponseBodyActionList {
	return s.ActionList
}

func (s *DescribeDingTalkResponseBody) GetPageInfo() *DescribeDingTalkResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeDingTalkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDingTalkResponseBody) SetActionList(v []*DescribeDingTalkResponseBodyActionList) *DescribeDingTalkResponseBody {
	s.ActionList = v
	return s
}

func (s *DescribeDingTalkResponseBody) SetPageInfo(v *DescribeDingTalkResponseBodyPageInfo) *DescribeDingTalkResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeDingTalkResponseBody) SetRequestId(v string) *DescribeDingTalkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDingTalkResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDingTalkResponseBodyActionList struct {
	// The name of the notification.
	//
	// example:
	//
	// Alert notification
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// The UID of the user.
	//
	// example:
	//
	// 12312412341
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The list of notification settings.
	//
	// example:
	//
	// [{\\"type\\":\\"vul\\",\\"configItemList\\":[{\\"key\\":\\"key\\", \\"valueList\\":\\"123\\"}]}]
	ConfigList *string `json:"ConfigList,omitempty" xml:"ConfigList,omitempty"`
	// The language of the content within notifications. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	DingTalkLang *string `json:"DingTalkLang,omitempty" xml:"DingTalkLang,omitempty"`
	// The creation time. unit:millisecond.
	//
	// example:
	//
	// 1550828400000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1550828400000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The group IDs.
	//
	// example:
	//
	// "123,456"
	GroupIdList *string `json:"GroupIdList,omitempty" xml:"GroupIdList,omitempty"`
	// The ID of the notification.
	//
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The interval at which the notifications are sent.unit:minute.
	//
	// example:
	//
	// 1000
	IntervalTime *int32 `json:"IntervalTime,omitempty" xml:"IntervalTime,omitempty"`
	// The status of the notification. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The parameters of the notification.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeDingTalkResponseBodyActionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDingTalkResponseBodyActionList) GoString() string {
	return s.String()
}

func (s *DescribeDingTalkResponseBodyActionList) GetActionName() *string {
	return s.ActionName
}

func (s *DescribeDingTalkResponseBodyActionList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeDingTalkResponseBodyActionList) GetConfigList() *string {
	return s.ConfigList
}

func (s *DescribeDingTalkResponseBodyActionList) GetDingTalkLang() *string {
	return s.DingTalkLang
}

func (s *DescribeDingTalkResponseBodyActionList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeDingTalkResponseBodyActionList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeDingTalkResponseBodyActionList) GetGroupIdList() *string {
	return s.GroupIdList
}

func (s *DescribeDingTalkResponseBodyActionList) GetId() *int32 {
	return s.Id
}

func (s *DescribeDingTalkResponseBodyActionList) GetIntervalTime() *int32 {
	return s.IntervalTime
}

func (s *DescribeDingTalkResponseBodyActionList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeDingTalkResponseBodyActionList) GetUrl() *string {
	return s.Url
}

func (s *DescribeDingTalkResponseBodyActionList) SetActionName(v string) *DescribeDingTalkResponseBodyActionList {
	s.ActionName = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetAliUid(v int64) *DescribeDingTalkResponseBodyActionList {
	s.AliUid = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetConfigList(v string) *DescribeDingTalkResponseBodyActionList {
	s.ConfigList = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetDingTalkLang(v string) *DescribeDingTalkResponseBodyActionList {
	s.DingTalkLang = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetGmtCreate(v int64) *DescribeDingTalkResponseBodyActionList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetGmtModified(v int64) *DescribeDingTalkResponseBodyActionList {
	s.GmtModified = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetGroupIdList(v string) *DescribeDingTalkResponseBodyActionList {
	s.GroupIdList = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetId(v int32) *DescribeDingTalkResponseBodyActionList {
	s.Id = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetIntervalTime(v int32) *DescribeDingTalkResponseBodyActionList {
	s.IntervalTime = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetStatus(v int32) *DescribeDingTalkResponseBodyActionList {
	s.Status = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetUrl(v string) *DescribeDingTalkResponseBodyActionList {
	s.Url = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) Validate() error {
	return dara.Validate(s)
}

type DescribeDingTalkResponseBodyPageInfo struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of messages.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDingTalkResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDingTalkResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeDingTalkResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDingTalkResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDingTalkResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDingTalkResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeDingTalkResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDingTalkResponseBodyPageInfo) SetPageSize(v int32) *DescribeDingTalkResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeDingTalkResponseBodyPageInfo) SetTotalCount(v int32) *DescribeDingTalkResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeDingTalkResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
