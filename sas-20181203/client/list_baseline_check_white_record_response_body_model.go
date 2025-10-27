// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaselineCheckWhiteRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*ListBaselineCheckWhiteRecordResponseBodyList) *ListBaselineCheckWhiteRecordResponseBody
	GetList() []*ListBaselineCheckWhiteRecordResponseBodyList
	SetPageInfo(v *ListBaselineCheckWhiteRecordResponseBodyPageInfo) *ListBaselineCheckWhiteRecordResponseBody
	GetPageInfo() *ListBaselineCheckWhiteRecordResponseBodyPageInfo
	SetRequestId(v string) *ListBaselineCheckWhiteRecordResponseBody
	GetRequestId() *string
}

type ListBaselineCheckWhiteRecordResponseBody struct {
	// The whitelist rules.
	List []*ListBaselineCheckWhiteRecordResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListBaselineCheckWhiteRecordResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9F4E6157-9600-5588-86B9-38F09067****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBaselineCheckWhiteRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBaselineCheckWhiteRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListBaselineCheckWhiteRecordResponseBody) GetList() []*ListBaselineCheckWhiteRecordResponseBodyList {
	return s.List
}

func (s *ListBaselineCheckWhiteRecordResponseBody) GetPageInfo() *ListBaselineCheckWhiteRecordResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListBaselineCheckWhiteRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBaselineCheckWhiteRecordResponseBody) SetList(v []*ListBaselineCheckWhiteRecordResponseBodyList) *ListBaselineCheckWhiteRecordResponseBody {
	s.List = v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBody) SetPageInfo(v *ListBaselineCheckWhiteRecordResponseBodyPageInfo) *ListBaselineCheckWhiteRecordResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBody) SetRequestId(v string) *ListBaselineCheckWhiteRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBaselineCheckWhiteRecordResponseBodyList struct {
	// The ID of the check item.
	//
	// example:
	//
	// 696
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The description of the check item.
	//
	// example:
	//
	// Config the Event Audit policys
	CheckItem *string `json:"CheckItem,omitempty" xml:"CheckItem,omitempty"`
	// The type of the check item.
	//
	// example:
	//
	// Security audit
	CheckType *string `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The display name of the check item type.
	//
	// example:
	//
	// Security audit
	CheckTypeDisName *string `json:"CheckTypeDisName,omitempty" xml:"CheckTypeDisName,omitempty"`
	// List of whitelisted container names in the check item.
	ContainerItems []*ListBaselineCheckWhiteRecordResponseBodyListContainerItems `json:"ContainerItems,omitempty" xml:"ContainerItems,omitempty" type:"Repeated"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The reason why the check item is added to the whitelist.
	//
	// example:
	//
	// AutoTest
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The ID of the whitelist rule.
	//
	// example:
	//
	// 79412
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The data source. Valid values:
	//
	// 	- **default**: server
	//
	// 	- **agentless**: agentless detection
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The object that is added to the whitelist.
	//
	// example:
	//
	// HOST_BASELINE_WHITE_LIST_21
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The type of the assets on which the whitelist rule takes effect. Valid values:
	//
	// 	- **all_instance**: all servers
	//
	// 	- **instance**: specific servers
	//
	// example:
	//
	// instance
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListBaselineCheckWhiteRecordResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListBaselineCheckWhiteRecordResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) GetCheckItem() *string {
	return s.CheckItem
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) GetCheckType() *string {
	return s.CheckType
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) GetCheckTypeDisName() *string {
	return s.CheckTypeDisName
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) GetContainerItems() []*ListBaselineCheckWhiteRecordResponseBodyListContainerItems {
	return s.ContainerItems
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) GetLang() *string {
	return s.Lang
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) GetReason() *string {
	return s.Reason
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) GetRecordId() *int64 {
	return s.RecordId
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) GetSource() *string {
	return s.Source
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) GetTarget() *string {
	return s.Target
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) GetTargetType() *string {
	return s.TargetType
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) SetCheckId(v int64) *ListBaselineCheckWhiteRecordResponseBodyList {
	s.CheckId = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) SetCheckItem(v string) *ListBaselineCheckWhiteRecordResponseBodyList {
	s.CheckItem = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) SetCheckType(v string) *ListBaselineCheckWhiteRecordResponseBodyList {
	s.CheckType = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) SetCheckTypeDisName(v string) *ListBaselineCheckWhiteRecordResponseBodyList {
	s.CheckTypeDisName = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) SetContainerItems(v []*ListBaselineCheckWhiteRecordResponseBodyListContainerItems) *ListBaselineCheckWhiteRecordResponseBodyList {
	s.ContainerItems = v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) SetLang(v string) *ListBaselineCheckWhiteRecordResponseBodyList {
	s.Lang = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) SetReason(v string) *ListBaselineCheckWhiteRecordResponseBodyList {
	s.Reason = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) SetRecordId(v int64) *ListBaselineCheckWhiteRecordResponseBodyList {
	s.RecordId = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) SetSource(v string) *ListBaselineCheckWhiteRecordResponseBodyList {
	s.Source = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) SetTarget(v string) *ListBaselineCheckWhiteRecordResponseBodyList {
	s.Target = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) SetTargetType(v string) *ListBaselineCheckWhiteRecordResponseBodyList {
	s.TargetType = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBodyList) Validate() error {
	if s.ContainerItems != nil {
		for _, item := range s.ContainerItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBaselineCheckWhiteRecordResponseBodyListContainerItems struct {
	// Names of the whitelisted containers for the current asset, separated by English commas.
	//
	// example:
	//
	// "anythingllm,ChuanhuChat"
	ContainerNames *string `json:"ContainerNames,omitempty" xml:"ContainerNames,omitempty"`
	// Server UUID.
	//
	// example:
	//
	// beeea5c2-1857-4b2b-a794-7d21eae*****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListBaselineCheckWhiteRecordResponseBodyListContainerItems) String() string {
	return dara.Prettify(s)
}

func (s ListBaselineCheckWhiteRecordResponseBodyListContainerItems) GoString() string {
	return s.String()
}

func (s *ListBaselineCheckWhiteRecordResponseBodyListContainerItems) GetContainerNames() *string {
	return s.ContainerNames
}

func (s *ListBaselineCheckWhiteRecordResponseBodyListContainerItems) GetUuid() *string {
	return s.Uuid
}

func (s *ListBaselineCheckWhiteRecordResponseBodyListContainerItems) SetContainerNames(v string) *ListBaselineCheckWhiteRecordResponseBodyListContainerItems {
	s.ContainerNames = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBodyListContainerItems) SetUuid(v string) *ListBaselineCheckWhiteRecordResponseBodyListContainerItems {
	s.Uuid = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBodyListContainerItems) Validate() error {
	return dara.Validate(s)
}

type ListBaselineCheckWhiteRecordResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 45
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBaselineCheckWhiteRecordResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListBaselineCheckWhiteRecordResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListBaselineCheckWhiteRecordResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListBaselineCheckWhiteRecordResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListBaselineCheckWhiteRecordResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBaselineCheckWhiteRecordResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBaselineCheckWhiteRecordResponseBodyPageInfo) SetCount(v int32) *ListBaselineCheckWhiteRecordResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBodyPageInfo) SetCurrentPage(v int32) *ListBaselineCheckWhiteRecordResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBodyPageInfo) SetPageSize(v int32) *ListBaselineCheckWhiteRecordResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBodyPageInfo) SetTotalCount(v int32) *ListBaselineCheckWhiteRecordResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
