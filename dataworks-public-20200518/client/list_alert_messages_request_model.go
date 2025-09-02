// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertMethods(v string) *ListAlertMessagesRequest
	GetAlertMethods() *string
	SetAlertRuleTypes(v string) *ListAlertMessagesRequest
	GetAlertRuleTypes() *string
	SetAlertUser(v string) *ListAlertMessagesRequest
	GetAlertUser() *string
	SetBaselineId(v int64) *ListAlertMessagesRequest
	GetBaselineId() *int64
	SetBeginTime(v string) *ListAlertMessagesRequest
	GetBeginTime() *string
	SetEndTime(v string) *ListAlertMessagesRequest
	GetEndTime() *string
	SetPageNumber(v int32) *ListAlertMessagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAlertMessagesRequest
	GetPageSize() *int32
	SetRemindId(v int64) *ListAlertMessagesRequest
	GetRemindId() *int64
}

type ListAlertMessagesRequest struct {
	// The notification method. Valid values:
	//
	// 	- MAIL
	//
	// 	- SMS Alert notifications can be sent by text message only in the Singapore, Malaysia (Kuala Lumpur), and Germany (Frankfurt) regions.
	//
	// You can specify multiple notification methods. Separate them with commas (,).
	//
	// example:
	//
	// SMS,MAIL,PHONE
	AlertMethods *string `json:"AlertMethods,omitempty" xml:"AlertMethods,omitempty"`
	// The type of the alert rule. Valid values: GLOBAL, USER_DEFINE, and OTHER. The value GLOBAL indicates that the alert rule is a global alert rule. The value USER_DEFINE indicates that the alert rule is customized by a user. The value OTHER indicates that the alert rule is a rule of another type. You can specify multiple types. Separate them with commas (,).
	//
	// example:
	//
	// GLOBAL,USER_DEFINE,OTHER
	AlertRuleTypes *string `json:"AlertRuleTypes,omitempty" xml:"AlertRuleTypes,omitempty"`
	// The ID of the Alibaba Cloud account used by the alert recipient.
	//
	// example:
	//
	// 123456
	AlertUser *string `json:"AlertUser,omitempty" xml:"AlertUser,omitempty"`
	// The baseline ID. This parameter takes effect if the AlertRuleTypes parameter is set to GLOBAL. You can configure either this parameter or the RemindId parameter.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-dd\\"T\\"HH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-04-02T00:00:00+0800
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-dd\\"T\\"HH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-04-04T00:00:00+0800
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. Default value: 1. Minimum value: 1. Maximum value: 30.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The custom alert rule ID. This parameter takes effect if the AlertRuleTypes parameter is set to USER_DEFINE. You can configure either this parameter or the BaselineId parameter.
	//
	// example:
	//
	// 9527
	RemindId *int64 `json:"RemindId,omitempty" xml:"RemindId,omitempty"`
}

func (s ListAlertMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListAlertMessagesRequest) GetAlertMethods() *string {
	return s.AlertMethods
}

func (s *ListAlertMessagesRequest) GetAlertRuleTypes() *string {
	return s.AlertRuleTypes
}

func (s *ListAlertMessagesRequest) GetAlertUser() *string {
	return s.AlertUser
}

func (s *ListAlertMessagesRequest) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListAlertMessagesRequest) GetBeginTime() *string {
	return s.BeginTime
}

func (s *ListAlertMessagesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAlertMessagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAlertMessagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlertMessagesRequest) GetRemindId() *int64 {
	return s.RemindId
}

func (s *ListAlertMessagesRequest) SetAlertMethods(v string) *ListAlertMessagesRequest {
	s.AlertMethods = &v
	return s
}

func (s *ListAlertMessagesRequest) SetAlertRuleTypes(v string) *ListAlertMessagesRequest {
	s.AlertRuleTypes = &v
	return s
}

func (s *ListAlertMessagesRequest) SetAlertUser(v string) *ListAlertMessagesRequest {
	s.AlertUser = &v
	return s
}

func (s *ListAlertMessagesRequest) SetBaselineId(v int64) *ListAlertMessagesRequest {
	s.BaselineId = &v
	return s
}

func (s *ListAlertMessagesRequest) SetBeginTime(v string) *ListAlertMessagesRequest {
	s.BeginTime = &v
	return s
}

func (s *ListAlertMessagesRequest) SetEndTime(v string) *ListAlertMessagesRequest {
	s.EndTime = &v
	return s
}

func (s *ListAlertMessagesRequest) SetPageNumber(v int32) *ListAlertMessagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertMessagesRequest) SetPageSize(v int32) *ListAlertMessagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertMessagesRequest) SetRemindId(v int64) *ListAlertMessagesRequest {
	s.RemindId = &v
	return s
}

func (s *ListAlertMessagesRequest) Validate() error {
	return dara.Validate(s)
}
