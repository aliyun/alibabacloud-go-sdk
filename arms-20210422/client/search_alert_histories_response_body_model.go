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
	PageBean  *SearchAlertHistoriesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AlarmHistories []*SearchAlertHistoriesResponseBodyPageBeanAlarmHistories `json:"AlarmHistories,omitempty" xml:"AlarmHistories,omitempty" type:"Repeated"`
	PageNumber     *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AlarmContent      *string `json:"AlarmContent,omitempty" xml:"AlarmContent,omitempty"`
	AlarmResponseCode *int32  `json:"AlarmResponseCode,omitempty" xml:"AlarmResponseCode,omitempty"`
	AlarmSources      *string `json:"AlarmSources,omitempty" xml:"AlarmSources,omitempty"`
	AlarmTime         *int64  `json:"AlarmTime,omitempty" xml:"AlarmTime,omitempty"`
	AlarmType         *int32  `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	Emails            *string `json:"Emails,omitempty" xml:"Emails,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Phones            *string `json:"Phones,omitempty" xml:"Phones,omitempty"`
	StrategyId        *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	Target            *string `json:"Target,omitempty" xml:"Target,omitempty"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
