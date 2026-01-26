// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAlertHistoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *SearchAlertHistoriesResponseBodyPageBean) *SearchAlertHistoriesResponseBody
	GetPageBean() *SearchAlertHistoriesResponseBodyPageBean
	SetRequestId(v string) *SearchAlertHistoriesResponseBody
	GetRequestId() *string
}

type SearchAlertHistoriesResponseBody struct {
	// The returned struct.
	PageBean *SearchAlertHistoriesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 2FC13182-B9AF-4E6B-BE51-72669B7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchAlertHistoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAlertHistoriesResponseBody) GetPageBean() *SearchAlertHistoriesResponseBodyPageBean {
	return s.PageBean
}

func (s *SearchAlertHistoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchAlertHistoriesResponseBody) SetPageBean(v *SearchAlertHistoriesResponseBodyPageBean) *SearchAlertHistoriesResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchAlertHistoriesResponseBody) SetRequestId(v string) *SearchAlertHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchAlertHistoriesResponseBody) Validate() error {
	if s.PageBean != nil {
		if err := s.PageBean.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchAlertHistoriesResponseBodyPageBean struct {
	// The information about alert records.
	AlarmHistories []*SearchAlertHistoriesResponseBodyPageBeanAlarmHistories `json:"AlarmHistories,omitempty" xml:"AlarmHistories,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchAlertHistoriesResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertHistoriesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchAlertHistoriesResponseBodyPageBean) GetAlarmHistories() []*SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	return s.AlarmHistories
}

func (s *SearchAlertHistoriesResponseBodyPageBean) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchAlertHistoriesResponseBodyPageBean) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchAlertHistoriesResponseBodyPageBean) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchAlertHistoriesResponseBodyPageBean) SetAlarmHistories(v []*SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) *SearchAlertHistoriesResponseBodyPageBean {
	s.AlarmHistories = v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBean) SetPageNumber(v int32) *SearchAlertHistoriesResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBean) SetPageSize(v int32) *SearchAlertHistoriesResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBean) SetTotalCount(v int32) *SearchAlertHistoriesResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBean) Validate() error {
	if s.AlarmHistories != nil {
		for _, item := range s.AlarmHistories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchAlertHistoriesResponseBodyPageBeanAlarmHistories struct {
	// The content of the alert notification.
	//
	// example:
	//
	// "Alert name: Alert1\\nAlert time: 2020-07-24 12:14:00\\nAlert content: A total of four alerts are triggered: \\*\\*\\*\\*"
	AlarmContent *string `json:"AlarmContent,omitempty" xml:"AlarmContent,omitempty"`
	// The response code returned after the alert notification was sent.
	//
	// example:
	//
	// 200
	AlarmResponseCode *int32 `json:"AlarmResponseCode,omitempty" xml:"AlarmResponseCode,omitempty"`
	// The webhook URL, such as the webhook URL of a DingTalk chatbot.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=91f2f65002fefe0ab9b71e6590c5ca504348cad742ff01e9c8ab204439ca****
	AlarmSources *string `json:"AlarmSources,omitempty" xml:"AlarmSources,omitempty"`
	// The time when the alert notification was sent.
	//
	// example:
	//
	// 1595564179000
	AlarmTime *int64 `json:"AlarmTime,omitempty" xml:"AlarmTime,omitempty"`
	// The type of the alert rule. Default value: 4. Valid values:
	//
	// 	- `1`: a custom alert rule that is used to monitor drill-down data sets
	//
	// 	- `3`: a custom alert rule that is used to monitor tiled data sets
	//
	// 	- `4`: an alert rule that is used to monitor web pages, including the default alert rule for browser monitoring
	//
	// 	- `5`: an alert rule that is used to monitor applications, including the default alert rule for application monitoring
	//
	// 	- `6`: the default alert rule for browser monitoring
	//
	// 	- `7`: the default alert rule for application monitoring
	//
	// 	- `8`: a Tracing Analysis alert rule
	//
	// 	- `101`: a Prometheus alert rule
	//
	// example:
	//
	// 4
	AlarmType *int32 `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	// The email address of the alert contact.
	//
	// example:
	//
	// someone@example.com
	Emails *string `json:"Emails,omitempty" xml:"Emails,omitempty"`
	// The ID of the alert notification.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The mobile phone number of the alert contact.
	//
	// example:
	//
	// 1381111****
	Phones *string `json:"Phones,omitempty" xml:"Phones,omitempty"`
	// The internal fields.
	//
	// example:
	//
	// ""
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The internal fields.
	//
	// example:
	//
	// ""
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 113197164949****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) GoString() string {
	return s.String()
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) GetAlarmContent() *string {
	return s.AlarmContent
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) GetAlarmResponseCode() *int32 {
	return s.AlarmResponseCode
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) GetAlarmSources() *string {
	return s.AlarmSources
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) GetAlarmTime() *int64 {
	return s.AlarmTime
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) GetAlarmType() *int32 {
	return s.AlarmType
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) GetEmails() *string {
	return s.Emails
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) GetId() *int64 {
	return s.Id
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) GetPhones() *string {
	return s.Phones
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) GetStrategyId() *string {
	return s.StrategyId
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) GetTarget() *string {
	return s.Target
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) GetUserId() *string {
	return s.UserId
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmContent(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmContent = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmResponseCode(v int32) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmResponseCode = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmSources(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmSources = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmTime(v int64) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmTime = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmType(v int32) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmType = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetEmails(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.Emails = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetId(v int64) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.Id = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetPhones(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.Phones = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetStrategyId(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.StrategyId = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetTarget(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.Target = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetUserId(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.UserId = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) Validate() error {
	return dara.Validate(s)
}
