// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmAlertLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v string) *ListCloudGtmAlertLogsRequest
	GetActionType() *string
	SetEndTimestamp(v int64) *ListCloudGtmAlertLogsRequest
	GetEndTimestamp() *int64
	SetEntityType(v string) *ListCloudGtmAlertLogsRequest
	GetEntityType() *string
	SetKeyword(v string) *ListCloudGtmAlertLogsRequest
	GetKeyword() *string
	SetLang(v string) *ListCloudGtmAlertLogsRequest
	GetLang() *string
	SetPageNumber(v int32) *ListCloudGtmAlertLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCloudGtmAlertLogsRequest
	GetPageSize() *int32
	SetStartTimestamp(v int64) *ListCloudGtmAlertLogsRequest
	GetStartTimestamp() *int64
}

type ListCloudGtmAlertLogsRequest struct {
	// The alert type.
	//
	// - ALERT: An alert is triggered.
	//
	// - RESUME: The service has recovered.
	//
	// example:
	//
	// ALERT
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The end of the time range to query. This is a UNIX timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1711328826977
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The alert object.
	//
	// - GTM_ADDRESS: Address
	//
	// - GTM_ADDRESS_POOL: Address pool
	//
	// - GTM_INSTANCE: Instance
	//
	// - GTM_MONITOR_TEMPLATE: Health check template
	//
	// example:
	//
	// GTM_ADDRESS
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The keyword for the search. This is usually an address ID, address pool ID, or domain name.
	//
	// example:
	//
	// pool-895280232254422016
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language of the response.
	//
	// - zh-CN: Chinese
	//
	// - en-US: English
	//
	// example:
	//
	// zh-CN
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The current page number. The value starts from **1**. The default value is **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for a paged query. The maximum value is 100. The default value is 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start of the time range to query. This is a UNIX timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1611328826977
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s ListCloudGtmAlertLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAlertLogsRequest) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAlertLogsRequest) GetActionType() *string {
	return s.ActionType
}

func (s *ListCloudGtmAlertLogsRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *ListCloudGtmAlertLogsRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *ListCloudGtmAlertLogsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListCloudGtmAlertLogsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListCloudGtmAlertLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloudGtmAlertLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudGtmAlertLogsRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *ListCloudGtmAlertLogsRequest) SetActionType(v string) *ListCloudGtmAlertLogsRequest {
	s.ActionType = &v
	return s
}

func (s *ListCloudGtmAlertLogsRequest) SetEndTimestamp(v int64) *ListCloudGtmAlertLogsRequest {
	s.EndTimestamp = &v
	return s
}

func (s *ListCloudGtmAlertLogsRequest) SetEntityType(v string) *ListCloudGtmAlertLogsRequest {
	s.EntityType = &v
	return s
}

func (s *ListCloudGtmAlertLogsRequest) SetKeyword(v string) *ListCloudGtmAlertLogsRequest {
	s.Keyword = &v
	return s
}

func (s *ListCloudGtmAlertLogsRequest) SetLang(v string) *ListCloudGtmAlertLogsRequest {
	s.Lang = &v
	return s
}

func (s *ListCloudGtmAlertLogsRequest) SetPageNumber(v int32) *ListCloudGtmAlertLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCloudGtmAlertLogsRequest) SetPageSize(v int32) *ListCloudGtmAlertLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudGtmAlertLogsRequest) SetStartTimestamp(v int64) *ListCloudGtmAlertLogsRequest {
	s.StartTimestamp = &v
	return s
}

func (s *ListCloudGtmAlertLogsRequest) Validate() error {
	return dara.Validate(s)
}
