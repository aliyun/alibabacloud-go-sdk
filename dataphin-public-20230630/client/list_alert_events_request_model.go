// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListAlertEventsRequestListQuery) *ListAlertEventsRequest
	GetListQuery() *ListAlertEventsRequestListQuery
	SetOpTenantId(v int64) *ListAlertEventsRequest
	GetOpTenantId() *int64
}

type ListAlertEventsRequest struct {
	// The query request.
	//
	// This parameter is required.
	ListQuery *ListAlertEventsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListAlertEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsRequest) GoString() string {
	return s.String()
}

func (s *ListAlertEventsRequest) GetListQuery() *ListAlertEventsRequestListQuery {
	return s.ListQuery
}

func (s *ListAlertEventsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListAlertEventsRequest) SetListQuery(v *ListAlertEventsRequestListQuery) *ListAlertEventsRequest {
	s.ListQuery = v
	return s
}

func (s *ListAlertEventsRequest) SetOpTenantId(v int64) *ListAlertEventsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListAlertEventsRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAlertEventsRequestListQuery struct {
	// The end date of the alert.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-11-27 13:47:00
	AlertEndTime *string `json:"AlertEndTime,omitempty" xml:"AlertEndTime,omitempty"`
	// The list of alert object types.
	AlertObjectTypeList []*string `json:"AlertObjectTypeList,omitempty" xml:"AlertObjectTypeList,omitempty" type:"Repeated"`
	// The list of alert reasons.
	AlertReasonList []*string `json:"AlertReasonList,omitempty" xml:"AlertReasonList,omitempty" type:"Repeated"`
	// The start date of the alert.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-11-16 00:00:00
	AlertStartTime *string `json:"AlertStartTime,omitempty" xml:"AlertStartTime,omitempty"`
	// The list of business units.
	BizNameList []*string `json:"BizNameList,omitempty" xml:"BizNameList,omitempty" type:"Repeated"`
	// The keyword for the query.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The list of monitored item IDs.
	MonitoredItemIdList []*string `json:"MonitoredItemIdList,omitempty" xml:"MonitoredItemIdList,omitempty" type:"Repeated"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of projects.
	ProjectNameList []*string `json:"ProjectNameList,omitempty" xml:"ProjectNameList,omitempty" type:"Repeated"`
	// The source system. Valid values:
	//
	// - ALL: all
	//
	// - DQE: data quality
	//
	// - OS: data service
	//
	// - STREAM: real-time computing
	//
	// - VDM_BATCH: offline computing
	//
	// - SOP: O&M platform
	//
	// - REAL_TIME_PIPELINE: real-time integration
	//
	// - KGB: baseline monitoring
	//
	// and more.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	SourceSystem *string `json:"SourceSystem,omitempty" xml:"SourceSystem,omitempty"`
	// The list of alert statuses.
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// The list of alert contact IDs.
	UserIdList []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s ListAlertEventsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListAlertEventsRequestListQuery) GetAlertEndTime() *string {
	return s.AlertEndTime
}

func (s *ListAlertEventsRequestListQuery) GetAlertObjectTypeList() []*string {
	return s.AlertObjectTypeList
}

func (s *ListAlertEventsRequestListQuery) GetAlertReasonList() []*string {
	return s.AlertReasonList
}

func (s *ListAlertEventsRequestListQuery) GetAlertStartTime() *string {
	return s.AlertStartTime
}

func (s *ListAlertEventsRequestListQuery) GetBizNameList() []*string {
	return s.BizNameList
}

func (s *ListAlertEventsRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAlertEventsRequestListQuery) GetMonitoredItemIdList() []*string {
	return s.MonitoredItemIdList
}

func (s *ListAlertEventsRequestListQuery) GetPage() *int32 {
	return s.Page
}

func (s *ListAlertEventsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlertEventsRequestListQuery) GetProjectNameList() []*string {
	return s.ProjectNameList
}

func (s *ListAlertEventsRequestListQuery) GetSourceSystem() *string {
	return s.SourceSystem
}

func (s *ListAlertEventsRequestListQuery) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListAlertEventsRequestListQuery) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *ListAlertEventsRequestListQuery) SetAlertEndTime(v string) *ListAlertEventsRequestListQuery {
	s.AlertEndTime = &v
	return s
}

func (s *ListAlertEventsRequestListQuery) SetAlertObjectTypeList(v []*string) *ListAlertEventsRequestListQuery {
	s.AlertObjectTypeList = v
	return s
}

func (s *ListAlertEventsRequestListQuery) SetAlertReasonList(v []*string) *ListAlertEventsRequestListQuery {
	s.AlertReasonList = v
	return s
}

func (s *ListAlertEventsRequestListQuery) SetAlertStartTime(v string) *ListAlertEventsRequestListQuery {
	s.AlertStartTime = &v
	return s
}

func (s *ListAlertEventsRequestListQuery) SetBizNameList(v []*string) *ListAlertEventsRequestListQuery {
	s.BizNameList = v
	return s
}

func (s *ListAlertEventsRequestListQuery) SetKeyword(v string) *ListAlertEventsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListAlertEventsRequestListQuery) SetMonitoredItemIdList(v []*string) *ListAlertEventsRequestListQuery {
	s.MonitoredItemIdList = v
	return s
}

func (s *ListAlertEventsRequestListQuery) SetPage(v int32) *ListAlertEventsRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListAlertEventsRequestListQuery) SetPageSize(v int32) *ListAlertEventsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListAlertEventsRequestListQuery) SetProjectNameList(v []*string) *ListAlertEventsRequestListQuery {
	s.ProjectNameList = v
	return s
}

func (s *ListAlertEventsRequestListQuery) SetSourceSystem(v string) *ListAlertEventsRequestListQuery {
	s.SourceSystem = &v
	return s
}

func (s *ListAlertEventsRequestListQuery) SetStatusList(v []*string) *ListAlertEventsRequestListQuery {
	s.StatusList = v
	return s
}

func (s *ListAlertEventsRequestListQuery) SetUserIdList(v []*string) *ListAlertEventsRequestListQuery {
	s.UserIdList = v
	return s
}

func (s *ListAlertEventsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
