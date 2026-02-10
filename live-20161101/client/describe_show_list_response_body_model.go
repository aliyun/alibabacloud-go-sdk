// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeShowListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeShowListResponseBody
	GetRequestId() *string
	SetShowList(v string) *DescribeShowListResponseBody
	GetShowList() *string
	SetShowListInfo(v *DescribeShowListResponseBodyShowListInfo) *DescribeShowListResponseBody
	GetShowListInfo() *DescribeShowListResponseBodyShowListInfo
}

type DescribeShowListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// ****Details of the episode list.
	//
	// Show indicates the information about a specific episode. For more information, see the **Show*	- parameter.
	//
	// example:
	//
	// ShowList[Show1, Show2, Show3...]
	ShowList *string `json:"ShowList,omitempty" xml:"ShowList,omitempty"`
	// The information about the episode list.
	ShowListInfo *DescribeShowListResponseBodyShowListInfo `json:"ShowListInfo,omitempty" xml:"ShowListInfo,omitempty" type:"Struct"`
}

func (s DescribeShowListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeShowListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeShowListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeShowListResponseBody) GetShowList() *string {
	return s.ShowList
}

func (s *DescribeShowListResponseBody) GetShowListInfo() *DescribeShowListResponseBodyShowListInfo {
	return s.ShowListInfo
}

func (s *DescribeShowListResponseBody) SetRequestId(v string) *DescribeShowListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeShowListResponseBody) SetShowList(v string) *DescribeShowListResponseBody {
	s.ShowList = &v
	return s
}

func (s *DescribeShowListResponseBody) SetShowListInfo(v *DescribeShowListResponseBodyShowListInfo) *DescribeShowListResponseBody {
	s.ShowListInfo = v
	return s
}

func (s *DescribeShowListResponseBody) Validate() error {
	if s.ShowListInfo != nil {
		if err := s.ShowListInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeShowListResponseBodyShowListInfo struct {
	// The background of the episode list.
	//
	// example:
	//
	// {\\"MaterialId\\":\\"a2b8e671-2fe5-4642-a2ec-bf93880e****\\",\\"resourceType\\":\\"VOD\\"}
	Background *string `json:"Background,omitempty" xml:"Background,omitempty"`
	// The ID of the episode that is playing.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CurrentShowId *string `json:"CurrentShowId,omitempty" xml:"CurrentShowId,omitempty"`
	// The episode of the highest priority.
	//
	// > You can configure this parameter only before the episode list starts playing.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	HighPriorityShowId *string `json:"HighPriorityShowId,omitempty" xml:"HighPriorityShowId,omitempty"`
	// The time at which the episode of the highest priority is played. Format: yyyy-MM-dd\\"T\\"HH:mm:ss.
	//
	// > You can configure this parameter only before the episode list starts playing. After you configure this parameter, when the specified point in time is reached, any episode that is playing stops and the episode of the highest priority in the episode list starts to play.
	//
	// example:
	//
	// 2021-11-23T12:30:00
	HighPriorityShowStartTime *string                                           `json:"HighPriorityShowStartTime,omitempty" xml:"HighPriorityShowStartTime,omitempty"`
	ShowList                  *DescribeShowListResponseBodyShowListInfoShowList `json:"ShowList,omitempty" xml:"ShowList,omitempty" type:"Struct"`
	// The number of additional times the episode list is played by default. The value is 0.
	//
	// example:
	//
	// 0
	ShowListRepeatTimes *int32 `json:"ShowListRepeatTimes,omitempty" xml:"ShowListRepeatTimes,omitempty"`
	// The number of additional times the episode list is played.
	//
	// example:
	//
	// 1
	TotalShowListRepeatTimes *int32 `json:"TotalShowListRepeatTimes,omitempty" xml:"TotalShowListRepeatTimes,omitempty"`
}

func (s DescribeShowListResponseBodyShowListInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeShowListResponseBodyShowListInfo) GoString() string {
	return s.String()
}

func (s *DescribeShowListResponseBodyShowListInfo) GetBackground() *string {
	return s.Background
}

func (s *DescribeShowListResponseBodyShowListInfo) GetCurrentShowId() *string {
	return s.CurrentShowId
}

func (s *DescribeShowListResponseBodyShowListInfo) GetHighPriorityShowId() *string {
	return s.HighPriorityShowId
}

func (s *DescribeShowListResponseBodyShowListInfo) GetHighPriorityShowStartTime() *string {
	return s.HighPriorityShowStartTime
}

func (s *DescribeShowListResponseBodyShowListInfo) GetShowList() *DescribeShowListResponseBodyShowListInfoShowList {
	return s.ShowList
}

func (s *DescribeShowListResponseBodyShowListInfo) GetShowListRepeatTimes() *int32 {
	return s.ShowListRepeatTimes
}

func (s *DescribeShowListResponseBodyShowListInfo) GetTotalShowListRepeatTimes() *int32 {
	return s.TotalShowListRepeatTimes
}

func (s *DescribeShowListResponseBodyShowListInfo) SetBackground(v string) *DescribeShowListResponseBodyShowListInfo {
	s.Background = &v
	return s
}

func (s *DescribeShowListResponseBodyShowListInfo) SetCurrentShowId(v string) *DescribeShowListResponseBodyShowListInfo {
	s.CurrentShowId = &v
	return s
}

func (s *DescribeShowListResponseBodyShowListInfo) SetHighPriorityShowId(v string) *DescribeShowListResponseBodyShowListInfo {
	s.HighPriorityShowId = &v
	return s
}

func (s *DescribeShowListResponseBodyShowListInfo) SetHighPriorityShowStartTime(v string) *DescribeShowListResponseBodyShowListInfo {
	s.HighPriorityShowStartTime = &v
	return s
}

func (s *DescribeShowListResponseBodyShowListInfo) SetShowList(v *DescribeShowListResponseBodyShowListInfoShowList) *DescribeShowListResponseBodyShowListInfo {
	s.ShowList = v
	return s
}

func (s *DescribeShowListResponseBodyShowListInfo) SetShowListRepeatTimes(v int32) *DescribeShowListResponseBodyShowListInfo {
	s.ShowListRepeatTimes = &v
	return s
}

func (s *DescribeShowListResponseBodyShowListInfo) SetTotalShowListRepeatTimes(v int32) *DescribeShowListResponseBodyShowListInfo {
	s.TotalShowListRepeatTimes = &v
	return s
}

func (s *DescribeShowListResponseBodyShowListInfo) Validate() error {
	if s.ShowList != nil {
		if err := s.ShowList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeShowListResponseBodyShowListInfoShowList struct {
	Show []*DescribeShowListResponseBodyShowListInfoShowListShow `json:"Show,omitempty" xml:"Show,omitempty" type:"Repeated"`
}

func (s DescribeShowListResponseBodyShowListInfoShowList) String() string {
	return dara.Prettify(s)
}

func (s DescribeShowListResponseBodyShowListInfoShowList) GoString() string {
	return s.String()
}

func (s *DescribeShowListResponseBodyShowListInfoShowList) GetShow() []*DescribeShowListResponseBodyShowListInfoShowListShow {
	return s.Show
}

func (s *DescribeShowListResponseBodyShowListInfoShowList) SetShow(v []*DescribeShowListResponseBodyShowListInfoShowListShow) *DescribeShowListResponseBodyShowListInfoShowList {
	s.Show = v
	return s
}

func (s *DescribeShowListResponseBodyShowListInfoShowList) Validate() error {
	if s.Show != nil {
		for _, item := range s.Show {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeShowListResponseBodyShowListInfoShowListShow struct {
	Duration     *int64                                                            `json:"Duration,omitempty" xml:"Duration,omitempty"`
	RepeatTimes  *int32                                                            `json:"RepeatTimes,omitempty" xml:"RepeatTimes,omitempty"`
	ResourceInfo *DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo `json:"ResourceInfo,omitempty" xml:"ResourceInfo,omitempty" type:"Struct"`
	ShowId       *string                                                           `json:"ShowId,omitempty" xml:"ShowId,omitempty"`
	ShowName     *string                                                           `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
}

func (s DescribeShowListResponseBodyShowListInfoShowListShow) String() string {
	return dara.Prettify(s)
}

func (s DescribeShowListResponseBodyShowListInfoShowListShow) GoString() string {
	return s.String()
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShow) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShow) GetRepeatTimes() *int32 {
	return s.RepeatTimes
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShow) GetResourceInfo() *DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo {
	return s.ResourceInfo
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShow) GetShowId() *string {
	return s.ShowId
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShow) GetShowName() *string {
	return s.ShowName
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShow) SetDuration(v int64) *DescribeShowListResponseBodyShowListInfoShowListShow {
	s.Duration = &v
	return s
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShow) SetRepeatTimes(v int32) *DescribeShowListResponseBodyShowListInfoShowListShow {
	s.RepeatTimes = &v
	return s
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShow) SetResourceInfo(v *DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo) *DescribeShowListResponseBodyShowListInfoShowListShow {
	s.ResourceInfo = v
	return s
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShow) SetShowId(v string) *DescribeShowListResponseBodyShowListInfoShowListShow {
	s.ShowId = &v
	return s
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShow) SetShowName(v string) *DescribeShowListResponseBodyShowListInfoShowListShow {
	s.ShowName = &v
	return s
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShow) Validate() error {
	if s.ResourceInfo != nil {
		if err := s.ResourceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo struct {
	LiveInputType *int32  `json:"LiveInputType,omitempty" xml:"LiveInputType,omitempty"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceUrl   *string `json:"ResourceUrl,omitempty" xml:"ResourceUrl,omitempty"`
}

func (s DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo) GoString() string {
	return s.String()
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo) GetLiveInputType() *int32 {
	return s.LiveInputType
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo) GetResourceUrl() *string {
	return s.ResourceUrl
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo) SetLiveInputType(v int32) *DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo {
	s.LiveInputType = &v
	return s
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo) SetResourceId(v string) *DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo {
	s.ResourceId = &v
	return s
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo) SetResourceType(v string) *DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo {
	s.ResourceType = &v
	return s
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo) SetResourceUrl(v string) *DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo {
	s.ResourceUrl = &v
	return s
}

func (s *DescribeShowListResponseBodyShowListInfoShowListShowResourceInfo) Validate() error {
	return dara.Validate(s)
}
