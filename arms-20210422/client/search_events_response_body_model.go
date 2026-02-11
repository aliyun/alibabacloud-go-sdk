// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTrigger(v int32) *SearchEventsResponseBody
	GetIsTrigger() *int32
	SetPageBean(v *SearchEventsResponseBodyPageBean) *SearchEventsResponseBody
	GetPageBean() *SearchEventsResponseBodyPageBean
	SetRequestId(v string) *SearchEventsResponseBody
	GetRequestId() *string
}

type SearchEventsResponseBody struct {
	IsTrigger *int32                            `json:"IsTrigger,omitempty" xml:"IsTrigger,omitempty"`
	PageBean  *SearchEventsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchEventsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchEventsResponseBody) GetIsTrigger() *int32 {
	return s.IsTrigger
}

func (s *SearchEventsResponseBody) GetPageBean() *SearchEventsResponseBodyPageBean {
	return s.PageBean
}

func (s *SearchEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchEventsResponseBody) SetIsTrigger(v int32) *SearchEventsResponseBody {
	s.IsTrigger = &v
	return s
}

func (s *SearchEventsResponseBody) SetPageBean(v *SearchEventsResponseBodyPageBean) *SearchEventsResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchEventsResponseBody) SetRequestId(v string) *SearchEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchEventsResponseBody) Validate() error {
	if s.PageBean != nil {
		if err := s.PageBean.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchEventsResponseBodyPageBean struct {
	Event      []*SearchEventsResponseBodyPageBeanEvent `json:"Event,omitempty" xml:"Event,omitempty" type:"Repeated"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchEventsResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s SearchEventsResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchEventsResponseBodyPageBean) GetEvent() []*SearchEventsResponseBodyPageBeanEvent {
	return s.Event
}

func (s *SearchEventsResponseBodyPageBean) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchEventsResponseBodyPageBean) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchEventsResponseBodyPageBean) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchEventsResponseBodyPageBean) SetEvent(v []*SearchEventsResponseBodyPageBeanEvent) *SearchEventsResponseBodyPageBean {
	s.Event = v
	return s
}

func (s *SearchEventsResponseBodyPageBean) SetPageNumber(v int32) *SearchEventsResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchEventsResponseBodyPageBean) SetPageSize(v int32) *SearchEventsResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchEventsResponseBodyPageBean) SetTotalCount(v int32) *SearchEventsResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

func (s *SearchEventsResponseBodyPageBean) Validate() error {
	if s.Event != nil {
		for _, item := range s.Event {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchEventsResponseBodyPageBeanEvent struct {
	AlertId    *int64    `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName  *string   `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	AlertRule  *string   `json:"AlertRule,omitempty" xml:"AlertRule,omitempty"`
	AlertType  *int32    `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	EventLevel *string   `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	EventTime  *int64    `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	Id         *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Links      []*string `json:"Links,omitempty" xml:"Links,omitempty" type:"Repeated"`
	Message    *string   `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SearchEventsResponseBodyPageBeanEvent) String() string {
	return dara.Prettify(s)
}

func (s SearchEventsResponseBodyPageBeanEvent) GoString() string {
	return s.String()
}

func (s *SearchEventsResponseBodyPageBeanEvent) GetAlertId() *int64 {
	return s.AlertId
}

func (s *SearchEventsResponseBodyPageBeanEvent) GetAlertName() *string {
	return s.AlertName
}

func (s *SearchEventsResponseBodyPageBeanEvent) GetAlertRule() *string {
	return s.AlertRule
}

func (s *SearchEventsResponseBodyPageBeanEvent) GetAlertType() *int32 {
	return s.AlertType
}

func (s *SearchEventsResponseBodyPageBeanEvent) GetEventLevel() *string {
	return s.EventLevel
}

func (s *SearchEventsResponseBodyPageBeanEvent) GetEventTime() *int64 {
	return s.EventTime
}

func (s *SearchEventsResponseBodyPageBeanEvent) GetId() *int64 {
	return s.Id
}

func (s *SearchEventsResponseBodyPageBeanEvent) GetLinks() []*string {
	return s.Links
}

func (s *SearchEventsResponseBodyPageBeanEvent) GetMessage() *string {
	return s.Message
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetAlertId(v int64) *SearchEventsResponseBodyPageBeanEvent {
	s.AlertId = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetAlertName(v string) *SearchEventsResponseBodyPageBeanEvent {
	s.AlertName = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetAlertRule(v string) *SearchEventsResponseBodyPageBeanEvent {
	s.AlertRule = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetAlertType(v int32) *SearchEventsResponseBodyPageBeanEvent {
	s.AlertType = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetEventLevel(v string) *SearchEventsResponseBodyPageBeanEvent {
	s.EventLevel = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetEventTime(v int64) *SearchEventsResponseBodyPageBeanEvent {
	s.EventTime = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetId(v int64) *SearchEventsResponseBodyPageBeanEvent {
	s.Id = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetLinks(v []*string) *SearchEventsResponseBodyPageBeanEvent {
	s.Links = v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetMessage(v string) *SearchEventsResponseBodyPageBeanEvent {
	s.Message = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) Validate() error {
	return dara.Validate(s)
}
