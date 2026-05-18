// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarms(v *DescribeMetricRuleListResponseBodyAlarms) *DescribeMetricRuleListResponseBody
	GetAlarms() *DescribeMetricRuleListResponseBodyAlarms
	SetCode(v int32) *DescribeMetricRuleListResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribeMetricRuleListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMetricRuleListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMetricRuleListResponseBody
	GetSuccess() *bool
	SetTotal(v string) *DescribeMetricRuleListResponseBody
	GetTotal() *string
}

type DescribeMetricRuleListResponseBody struct {
	Alarms *DescribeMetricRuleListResponseBodyAlarms `json:"Alarms,omitempty" xml:"Alarms,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 386C6712-335F-5054-930A-CC92B851ECBA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeMetricRuleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBody) GetAlarms() *DescribeMetricRuleListResponseBodyAlarms {
	return s.Alarms
}

func (s *DescribeMetricRuleListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeMetricRuleListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMetricRuleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricRuleListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMetricRuleListResponseBody) GetTotal() *string {
	return s.Total
}

func (s *DescribeMetricRuleListResponseBody) SetAlarms(v *DescribeMetricRuleListResponseBodyAlarms) *DescribeMetricRuleListResponseBody {
	s.Alarms = v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetCode(v int32) *DescribeMetricRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetMessage(v string) *DescribeMetricRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetRequestId(v string) *DescribeMetricRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetSuccess(v bool) *DescribeMetricRuleListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetTotal(v string) *DescribeMetricRuleListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) Validate() error {
	if s.Alarms != nil {
		if err := s.Alarms.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMetricRuleListResponseBodyAlarms struct {
	Alarm []*DescribeMetricRuleListResponseBodyAlarmsAlarm `json:"Alarm,omitempty" xml:"Alarm,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleListResponseBodyAlarms) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarms) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarms) GetAlarm() []*DescribeMetricRuleListResponseBodyAlarmsAlarm {
	return s.Alarm
}

func (s *DescribeMetricRuleListResponseBodyAlarms) SetAlarm(v []*DescribeMetricRuleListResponseBodyAlarmsAlarm) *DescribeMetricRuleListResponseBodyAlarms {
	s.Alarm = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarms) Validate() error {
	if s.Alarm != nil {
		for _, item := range s.Alarm {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricRuleListResponseBodyAlarmsAlarm struct {
	AlertState          *string                                                           `json:"AlertState,omitempty" xml:"AlertState,omitempty"`
	CompositeExpression *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression `json:"CompositeExpression,omitempty" xml:"CompositeExpression,omitempty" type:"Struct"`
	ContactGroups       *string                                                           `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	Dimensions          *string                                                           `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EffectiveInterval   *string                                                           `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	EnableState         *bool                                                             `json:"EnableState,omitempty" xml:"EnableState,omitempty"`
	Escalations         *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations         `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	GmtCreate           *int64                                                            `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtUpdate           *string                                                           `json:"GmtUpdate,omitempty" xml:"GmtUpdate,omitempty"`
	GroupId             *string                                                           `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName           *string                                                           `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Labels              *DescribeMetricRuleListResponseBodyAlarmsAlarmLabels              `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Struct"`
	MailSubject         *string                                                           `json:"MailSubject,omitempty" xml:"MailSubject,omitempty"`
	MetricName          *string                                                           `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Namespace           *string                                                           `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NoDataPolicy        *string                                                           `json:"NoDataPolicy,omitempty" xml:"NoDataPolicy,omitempty"`
	NoEffectiveInterval *string                                                           `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	Period              *string                                                           `json:"Period,omitempty" xml:"Period,omitempty"`
	ProductCategory     *string                                                           `json:"ProductCategory,omitempty" xml:"ProductCategory,omitempty"`
	Prometheus          *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus          `json:"Prometheus,omitempty" xml:"Prometheus,omitempty" type:"Struct"`
	Resources           *string                                                           `json:"Resources,omitempty" xml:"Resources,omitempty"`
	RuleId              *string                                                           `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName            *string                                                           `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	SilenceTime         *int32                                                            `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	SourceType          *string                                                           `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Webhook             *string                                                           `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarm) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarm) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetAlertState() *string {
	return s.AlertState
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetCompositeExpression() *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression {
	return s.CompositeExpression
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetContactGroups() *string {
	return s.ContactGroups
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetDimensions() *string {
	return s.Dimensions
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetEffectiveInterval() *string {
	return s.EffectiveInterval
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetEnableState() *bool {
	return s.EnableState
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetEscalations() *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations {
	return s.Escalations
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetGmtUpdate() *string {
	return s.GmtUpdate
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetLabels() *DescribeMetricRuleListResponseBodyAlarmsAlarmLabels {
	return s.Labels
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetMailSubject() *string {
	return s.MailSubject
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetNoDataPolicy() *string {
	return s.NoDataPolicy
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetNoEffectiveInterval() *string {
	return s.NoEffectiveInterval
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetPeriod() *string {
	return s.Period
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetProductCategory() *string {
	return s.ProductCategory
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetPrometheus() *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus {
	return s.Prometheus
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetResources() *string {
	return s.Resources
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetWebhook() *string {
	return s.Webhook
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetAlertState(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.AlertState = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetCompositeExpression(v *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.CompositeExpression = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetContactGroups(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.ContactGroups = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetDimensions(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetEffectiveInterval(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.EffectiveInterval = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetEnableState(v bool) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.EnableState = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetEscalations(v *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Escalations = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetGmtCreate(v int64) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.GmtCreate = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetGmtUpdate(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.GmtUpdate = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetGroupId(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.GroupId = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetGroupName(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.GroupName = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetLabels(v *DescribeMetricRuleListResponseBodyAlarmsAlarmLabels) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Labels = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetMailSubject(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.MailSubject = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetMetricName(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetNamespace(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetNoDataPolicy(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.NoDataPolicy = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetNoEffectiveInterval(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.NoEffectiveInterval = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetPeriod(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Period = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetProductCategory(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.ProductCategory = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetPrometheus(v *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Prometheus = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetResources(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Resources = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetRuleId(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.RuleId = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetRuleName(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.RuleName = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetSilenceTime(v int32) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.SilenceTime = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetSourceType(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.SourceType = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetWebhook(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Webhook = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) Validate() error {
	if s.CompositeExpression != nil {
		if err := s.CompositeExpression.Validate(); err != nil {
			return err
		}
	}
	if s.Escalations != nil {
		if err := s.Escalations.Validate(); err != nil {
			return err
		}
	}
	if s.Labels != nil {
		if err := s.Labels.Validate(); err != nil {
			return err
		}
	}
	if s.Prometheus != nil {
		if err := s.Prometheus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression struct {
	ExpressionList     *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList `json:"ExpressionList,omitempty" xml:"ExpressionList,omitempty" type:"Struct"`
	ExpressionListJoin *string                                                                         `json:"ExpressionListJoin,omitempty" xml:"ExpressionListJoin,omitempty"`
	ExpressionRaw      *string                                                                         `json:"ExpressionRaw,omitempty" xml:"ExpressionRaw,omitempty"`
	Level              *string                                                                         `json:"Level,omitempty" xml:"Level,omitempty"`
	Times              *int32                                                                          `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) GetExpressionList() *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList {
	return s.ExpressionList
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) GetExpressionListJoin() *string {
	return s.ExpressionListJoin
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) GetExpressionRaw() *string {
	return s.ExpressionRaw
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) GetLevel() *string {
	return s.Level
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) GetTimes() *int32 {
	return s.Times
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) SetExpressionList(v *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression {
	s.ExpressionList = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) SetExpressionListJoin(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression {
	s.ExpressionListJoin = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) SetExpressionRaw(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression {
	s.ExpressionRaw = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) SetLevel(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression {
	s.Level = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) SetTimes(v int32) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) Validate() error {
	if s.ExpressionList != nil {
		if err := s.ExpressionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList struct {
	ExpressionList []*DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList `json:"ExpressionList,omitempty" xml:"ExpressionList,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList) GetExpressionList() []*DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList {
	return s.ExpressionList
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList) SetExpressionList(v []*DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList {
	s.ExpressionList = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList) Validate() error {
	if s.ExpressionList != nil {
		for _, item := range s.ExpressionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	MetricName         *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Period             *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) SetComparisonOperator(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) SetMetricName(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) SetPeriod(v int32) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList {
	s.Period = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) SetStatistics(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList {
	s.Statistics = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) SetThreshold(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations struct {
	Critical *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) GetCritical() *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	return s.Critical
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) GetInfo() *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	return s.Info
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) GetWarn() *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	return s.Warn
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) SetCritical(v *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations {
	s.Critical = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) SetInfo(v *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations {
	s.Info = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) SetWarn(v *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations {
	s.Warn = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) Validate() error {
	if s.Critical != nil {
		if err := s.Critical.Validate(); err != nil {
			return err
		}
	}
	if s.Info != nil {
		if err := s.Info.Validate(); err != nil {
			return err
		}
	}
	if s.Warn != nil {
		if err := s.Warn.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) GetPreCondition() *string {
	return s.PreCondition
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) GetTimes() *int32 {
	return s.Times
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetComparisonOperator(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetPreCondition(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.PreCondition = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetStatistics(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetThreshold(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetTimes(v int32) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) GetPreCondition() *string {
	return s.PreCondition
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) GetTimes() *int32 {
	return s.Times
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetComparisonOperator(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetPreCondition(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.PreCondition = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetStatistics(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetThreshold(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetTimes(v int32) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) GetPreCondition() *string {
	return s.PreCondition
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) GetTimes() *int32 {
	return s.Times
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetComparisonOperator(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetPreCondition(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.PreCondition = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetStatistics(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetThreshold(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetTimes(v int32) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmLabels struct {
	Labels []*DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmLabels) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmLabels) GetLabels() []*DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels {
	return s.Labels
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmLabels) SetLabels(v []*DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels) *DescribeMetricRuleListResponseBodyAlarmsAlarmLabels {
	s.Labels = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmLabels) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels) GetKey() *string {
	return s.Key
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels) GetValue() *string {
	return s.Value
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels) SetKey(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels {
	s.Key = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels) SetValue(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels {
	s.Value = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus struct {
	Annotations *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Struct"`
	Level       *string                                                             `json:"Level,omitempty" xml:"Level,omitempty"`
	PromQL      *string                                                             `json:"PromQL,omitempty" xml:"PromQL,omitempty"`
	Times       *int64                                                              `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) GetAnnotations() *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations {
	return s.Annotations
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) GetLevel() *string {
	return s.Level
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) GetPromQL() *string {
	return s.PromQL
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) GetTimes() *int64 {
	return s.Times
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) SetAnnotations(v *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations) *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus {
	s.Annotations = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) SetLevel(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus {
	s.Level = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) SetPromQL(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus {
	s.PromQL = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) SetTimes(v int64) *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) Validate() error {
	if s.Annotations != nil {
		if err := s.Annotations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations struct {
	Annotations []*DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations) GetAnnotations() []*DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations {
	return s.Annotations
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations) SetAnnotations(v []*DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations) *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations {
	s.Annotations = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations) Validate() error {
	if s.Annotations != nil {
		for _, item := range s.Annotations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations) GetKey() *string {
	return s.Key
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations) GetValue() *string {
	return s.Value
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations) SetKey(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations {
	s.Key = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations) SetValue(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations {
	s.Value = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations) Validate() error {
	return dara.Validate(s)
}
