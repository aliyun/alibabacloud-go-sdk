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
	// Specifies whether the alert event is triggered. If you do not set this parameter, all alert events are queried. Valid values:
	//
	// 	- `1`: The event is triggered.
	//
	// 	- `0`: The event is not triggered.
	//
	// example:
	//
	// 0
	IsTrigger *int32 `json:"IsTrigger,omitempty" xml:"IsTrigger,omitempty"`
	// The struct returned.
	PageBean *SearchEventsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 32940175-181B-4B93-966E-4BB69176****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	return dara.Validate(s)
}

type SearchEventsResponseBodyPageBean struct {
	// The information about the alert events.
	Event []*SearchEventsResponseBodyPageBeanEvent `json:"Event,omitempty" xml:"Event,omitempty" type:"Repeated"`
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
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	return dara.Validate(s)
}

type SearchEventsResponseBodyPageBeanEvent struct {
	// The ID of the alert rule that is associated with the event.
	//
	// example:
	//
	// 123
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The name of the alert rule that is associated with the event.
	//
	// example:
	//
	// alertName
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The condition of the alert rule.
	//
	// example:
	//
	// {\\"operator\\":\\"&\\",\\"rules\\":[{\\"aggregates\\":\\"AVG\\",\\"alias\\":\\"JVM_线程总数\\",\\"measure\\":\\"appstat.jvm.ThreadCount\\",\\"nValue\\":1,\\"operator\\":\\"HOH_DOWN\\",\\"value\\":50.0}]}
	AlertRule *string `json:"AlertRule,omitempty" xml:"AlertRule,omitempty"`
	// The type of the alert rule. This parameter is not returned. Valid values:
	//
	// 	- `1`: custom alert rules to monitor drill-down data sets
	//
	// 	- `3`: custom alert rules to monitor tiled data sets
	//
	// 	- `4`: alert rules to monitor the frontend, including the default frontend alert rules
	//
	// 	- `5`: alert rules to monitor applications, including the default application alert rules
	//
	// 	- `6`: the default frontend alert rules
	//
	// 	- `7`: the default application alert rules
	//
	// 	- `8`: Tracing Analysis alert rules
	//
	// 	- `101`: Prometheus alert rules
	//
	// example:
	//
	// 4
	AlertType *int32 `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The severity of the event.
	//
	// example:
	//
	// 1
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The timestamp when the event occurred.
	//
	// example:
	//
	// 1595569020000
	EventTime *int64 `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// The ID of the event record.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The list of event URLs.
	Links []*string `json:"Links,omitempty" xml:"Links,omitempty" type:"Repeated"`
	// The event content. The parameter value is a JSON string. Each key indicates a dimension and each value indicates the alert content in the dimension.
	//
	// example:
	//
	// unknow紧急报警\\nip：172.27.XX.XX\\n应用名 = test\\nRegion = cn-shenzhen\\n异常信息 = {\\"timestamp\\"：\\"1615447972235\\"}
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
