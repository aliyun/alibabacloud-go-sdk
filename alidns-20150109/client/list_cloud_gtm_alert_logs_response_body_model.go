// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmAlertLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v *ListCloudGtmAlertLogsResponseBodyLogs) *ListCloudGtmAlertLogsResponseBody
	GetLogs() *ListCloudGtmAlertLogsResponseBodyLogs
	SetPageNumber(v int32) *ListCloudGtmAlertLogsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCloudGtmAlertLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCloudGtmAlertLogsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *ListCloudGtmAlertLogsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *ListCloudGtmAlertLogsResponseBody
	GetTotalPages() *int32
}

type ListCloudGtmAlertLogsResponseBody struct {
	// The alert logs.
	Logs *ListCloudGtmAlertLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Struct"`
	// Current page number, starting from 1, default is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of rows per page when paginating queries, with a maximum value of 100 and a default of 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Unique request identification code.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of alarm log entries.
	//
	// example:
	//
	// 15
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListCloudGtmAlertLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAlertLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAlertLogsResponseBody) GetLogs() *ListCloudGtmAlertLogsResponseBodyLogs {
	return s.Logs
}

func (s *ListCloudGtmAlertLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloudGtmAlertLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudGtmAlertLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudGtmAlertLogsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *ListCloudGtmAlertLogsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListCloudGtmAlertLogsResponseBody) SetLogs(v *ListCloudGtmAlertLogsResponseBodyLogs) *ListCloudGtmAlertLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListCloudGtmAlertLogsResponseBody) SetPageNumber(v int32) *ListCloudGtmAlertLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCloudGtmAlertLogsResponseBody) SetPageSize(v int32) *ListCloudGtmAlertLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCloudGtmAlertLogsResponseBody) SetRequestId(v string) *ListCloudGtmAlertLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudGtmAlertLogsResponseBody) SetTotalItems(v int32) *ListCloudGtmAlertLogsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *ListCloudGtmAlertLogsResponseBody) SetTotalPages(v int32) *ListCloudGtmAlertLogsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *ListCloudGtmAlertLogsResponseBody) Validate() error {
	if s.Logs != nil {
		if err := s.Logs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudGtmAlertLogsResponseBodyLogs struct {
	Log []*ListCloudGtmAlertLogsResponseBodyLogsLog `json:"Log,omitempty" xml:"Log,omitempty" type:"Repeated"`
}

func (s ListCloudGtmAlertLogsResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAlertLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAlertLogsResponseBodyLogs) GetLog() []*ListCloudGtmAlertLogsResponseBodyLogsLog {
	return s.Log
}

func (s *ListCloudGtmAlertLogsResponseBodyLogs) SetLog(v []*ListCloudGtmAlertLogsResponseBodyLogsLog) *ListCloudGtmAlertLogsResponseBodyLogs {
	s.Log = v
	return s
}

func (s *ListCloudGtmAlertLogsResponseBodyLogs) Validate() error {
	if s.Log != nil {
		for _, item := range s.Log {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCloudGtmAlertLogsResponseBodyLogsLog struct {
	// Alert type:
	//
	// - ALERT
	//
	// - RESUME
	//
	// example:
	//
	// ALERT
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The alert content.
	//
	// example:
	//
	// The alert content.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Alarm object types:
	//
	// - GTM_ADDRESS: Address
	//
	// - GTM_ADDRESS_POOL: Address Pool
	//
	// - GTM_INSTANCE: Instance
	//
	// - GTM_MONITOR_TEMPLATE: Health Check Template
	//
	// example:
	//
	// GTM_ADDRESS
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// Alert log time (timestamp).
	//
	// example:
	//
	// 1711328826977
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListCloudGtmAlertLogsResponseBodyLogsLog) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAlertLogsResponseBodyLogsLog) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAlertLogsResponseBodyLogsLog) GetActionType() *string {
	return s.ActionType
}

func (s *ListCloudGtmAlertLogsResponseBodyLogsLog) GetContent() *string {
	return s.Content
}

func (s *ListCloudGtmAlertLogsResponseBodyLogsLog) GetEntityType() *string {
	return s.EntityType
}

func (s *ListCloudGtmAlertLogsResponseBodyLogsLog) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ListCloudGtmAlertLogsResponseBodyLogsLog) SetActionType(v string) *ListCloudGtmAlertLogsResponseBodyLogsLog {
	s.ActionType = &v
	return s
}

func (s *ListCloudGtmAlertLogsResponseBodyLogsLog) SetContent(v string) *ListCloudGtmAlertLogsResponseBodyLogsLog {
	s.Content = &v
	return s
}

func (s *ListCloudGtmAlertLogsResponseBodyLogsLog) SetEntityType(v string) *ListCloudGtmAlertLogsResponseBodyLogsLog {
	s.EntityType = &v
	return s
}

func (s *ListCloudGtmAlertLogsResponseBodyLogsLog) SetTimestamp(v int64) *ListCloudGtmAlertLogsResponseBodyLogsLog {
	s.Timestamp = &v
	return s
}

func (s *ListCloudGtmAlertLogsResponseBodyLogsLog) Validate() error {
	return dara.Validate(s)
}
