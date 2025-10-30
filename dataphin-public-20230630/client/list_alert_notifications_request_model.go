// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertNotificationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListAlertNotificationsRequestListQuery) *ListAlertNotificationsRequest
	GetListQuery() *ListAlertNotificationsRequestListQuery
	SetOpTenantId(v int64) *ListAlertNotificationsRequest
	GetOpTenantId() *int64
}

type ListAlertNotificationsRequest struct {
	ListQuery *ListAlertNotificationsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListAlertNotificationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotificationsRequest) GoString() string {
	return s.String()
}

func (s *ListAlertNotificationsRequest) GetListQuery() *ListAlertNotificationsRequestListQuery {
	return s.ListQuery
}

func (s *ListAlertNotificationsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListAlertNotificationsRequest) SetListQuery(v *ListAlertNotificationsRequestListQuery) *ListAlertNotificationsRequest {
	s.ListQuery = v
	return s
}

func (s *ListAlertNotificationsRequest) SetOpTenantId(v int64) *ListAlertNotificationsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListAlertNotificationsRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAlertNotificationsRequestListQuery struct {
	AlertReasonList     []*string `json:"AlertReasonList,omitempty" xml:"AlertReasonList,omitempty" type:"Repeated"`
	ChannelTypeList     []*string `json:"ChannelTypeList,omitempty" xml:"ChannelTypeList,omitempty" type:"Repeated"`
	CustomChannelIdList []*string `json:"CustomChannelIdList,omitempty" xml:"CustomChannelIdList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Keyword             *string   `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	MonitoredItemIdList []*string `json:"MonitoredItemIdList,omitempty" xml:"MonitoredItemIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-11-27 13:47:00
	NotifyEndTime *string `json:"NotifyEndTime,omitempty" xml:"NotifyEndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-11-16 00:00:00
	NotifyStartTime *string `json:"NotifyStartTime,omitempty" xml:"NotifyStartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALL
	SourceSystem *string   `json:"SourceSystem,omitempty" xml:"SourceSystem,omitempty"`
	StatusList   []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	UserIdList   []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s ListAlertNotificationsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotificationsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListAlertNotificationsRequestListQuery) GetAlertReasonList() []*string {
	return s.AlertReasonList
}

func (s *ListAlertNotificationsRequestListQuery) GetChannelTypeList() []*string {
	return s.ChannelTypeList
}

func (s *ListAlertNotificationsRequestListQuery) GetCustomChannelIdList() []*string {
	return s.CustomChannelIdList
}

func (s *ListAlertNotificationsRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAlertNotificationsRequestListQuery) GetMonitoredItemIdList() []*string {
	return s.MonitoredItemIdList
}

func (s *ListAlertNotificationsRequestListQuery) GetNotifyEndTime() *string {
	return s.NotifyEndTime
}

func (s *ListAlertNotificationsRequestListQuery) GetNotifyStartTime() *string {
	return s.NotifyStartTime
}

func (s *ListAlertNotificationsRequestListQuery) GetPage() *int32 {
	return s.Page
}

func (s *ListAlertNotificationsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlertNotificationsRequestListQuery) GetSourceSystem() *string {
	return s.SourceSystem
}

func (s *ListAlertNotificationsRequestListQuery) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListAlertNotificationsRequestListQuery) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *ListAlertNotificationsRequestListQuery) SetAlertReasonList(v []*string) *ListAlertNotificationsRequestListQuery {
	s.AlertReasonList = v
	return s
}

func (s *ListAlertNotificationsRequestListQuery) SetChannelTypeList(v []*string) *ListAlertNotificationsRequestListQuery {
	s.ChannelTypeList = v
	return s
}

func (s *ListAlertNotificationsRequestListQuery) SetCustomChannelIdList(v []*string) *ListAlertNotificationsRequestListQuery {
	s.CustomChannelIdList = v
	return s
}

func (s *ListAlertNotificationsRequestListQuery) SetKeyword(v string) *ListAlertNotificationsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListAlertNotificationsRequestListQuery) SetMonitoredItemIdList(v []*string) *ListAlertNotificationsRequestListQuery {
	s.MonitoredItemIdList = v
	return s
}

func (s *ListAlertNotificationsRequestListQuery) SetNotifyEndTime(v string) *ListAlertNotificationsRequestListQuery {
	s.NotifyEndTime = &v
	return s
}

func (s *ListAlertNotificationsRequestListQuery) SetNotifyStartTime(v string) *ListAlertNotificationsRequestListQuery {
	s.NotifyStartTime = &v
	return s
}

func (s *ListAlertNotificationsRequestListQuery) SetPage(v int32) *ListAlertNotificationsRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListAlertNotificationsRequestListQuery) SetPageSize(v int32) *ListAlertNotificationsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListAlertNotificationsRequestListQuery) SetSourceSystem(v string) *ListAlertNotificationsRequestListQuery {
	s.SourceSystem = &v
	return s
}

func (s *ListAlertNotificationsRequestListQuery) SetStatusList(v []*string) *ListAlertNotificationsRequestListQuery {
	s.StatusList = v
	return s
}

func (s *ListAlertNotificationsRequestListQuery) SetUserIdList(v []*string) *ListAlertNotificationsRequestListQuery {
	s.UserIdList = v
	return s
}

func (s *ListAlertNotificationsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
