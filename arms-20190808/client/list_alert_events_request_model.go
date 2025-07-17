// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertName(v string) *ListAlertEventsRequest
	GetAlertName() *string
	SetEndTime(v string) *ListAlertEventsRequest
	GetEndTime() *string
	SetMatchingConditions(v string) *ListAlertEventsRequest
	GetMatchingConditions() *string
	SetPage(v int64) *ListAlertEventsRequest
	GetPage() *int64
	SetShowNotificationPolicies(v bool) *ListAlertEventsRequest
	GetShowNotificationPolicies() *bool
	SetSize(v int64) *ListAlertEventsRequest
	GetSize() *int64
	SetStartTime(v string) *ListAlertEventsRequest
	GetStartTime() *string
	SetStatus(v string) *ListAlertEventsRequest
	GetStatus() *string
}

type ListAlertEventsRequest struct {
	// The name of the alert.
	//
	// example:
	//
	// Test-triggered alert
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The end time of the alert events that you want to query. Specify the time in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2021-12-22 23:59:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of matching conditions.
	//
	// example:
	//
	// [         {           "value": "ARMS_NOTIFICATION",           "key": "clustername",           "operator": "eq"         }       ]     },{       "matchingConditions": [         {           "value": "test",           "key": "alertname",           "operator": "eq"         }       ]
	MatchingConditions *string `json:"MatchingConditions,omitempty" xml:"MatchingConditions,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// Specifies whether to show the associated notification policy.
	//
	// example:
	//
	// false
	ShowNotificationPolicies *bool `json:"ShowNotificationPolicies,omitempty" xml:"ShowNotificationPolicies,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The start time of the alert events that you want to query. Specify the time in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2021-12-19 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the alert events. Valid values:
	//
	// 	- Active
	//
	// 	- Silenced
	//
	// 	- Resolved
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAlertEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsRequest) GoString() string {
	return s.String()
}

func (s *ListAlertEventsRequest) GetAlertName() *string {
	return s.AlertName
}

func (s *ListAlertEventsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAlertEventsRequest) GetMatchingConditions() *string {
	return s.MatchingConditions
}

func (s *ListAlertEventsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListAlertEventsRequest) GetShowNotificationPolicies() *bool {
	return s.ShowNotificationPolicies
}

func (s *ListAlertEventsRequest) GetSize() *int64 {
	return s.Size
}

func (s *ListAlertEventsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAlertEventsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAlertEventsRequest) SetAlertName(v string) *ListAlertEventsRequest {
	s.AlertName = &v
	return s
}

func (s *ListAlertEventsRequest) SetEndTime(v string) *ListAlertEventsRequest {
	s.EndTime = &v
	return s
}

func (s *ListAlertEventsRequest) SetMatchingConditions(v string) *ListAlertEventsRequest {
	s.MatchingConditions = &v
	return s
}

func (s *ListAlertEventsRequest) SetPage(v int64) *ListAlertEventsRequest {
	s.Page = &v
	return s
}

func (s *ListAlertEventsRequest) SetShowNotificationPolicies(v bool) *ListAlertEventsRequest {
	s.ShowNotificationPolicies = &v
	return s
}

func (s *ListAlertEventsRequest) SetSize(v int64) *ListAlertEventsRequest {
	s.Size = &v
	return s
}

func (s *ListAlertEventsRequest) SetStartTime(v string) *ListAlertEventsRequest {
	s.StartTime = &v
	return s
}

func (s *ListAlertEventsRequest) SetStatus(v string) *ListAlertEventsRequest {
	s.Status = &v
	return s
}

func (s *ListAlertEventsRequest) Validate() error {
	return dara.Validate(s)
}
