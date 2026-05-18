// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveMetricRuleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertList(v *DescribeActiveMetricRuleListResponseBodyAlertList) *DescribeActiveMetricRuleListResponseBody
	GetAlertList() *DescribeActiveMetricRuleListResponseBodyAlertList
	SetCode(v string) *DescribeActiveMetricRuleListResponseBody
	GetCode() *string
	SetDatapoints(v *DescribeActiveMetricRuleListResponseBodyDatapoints) *DescribeActiveMetricRuleListResponseBody
	GetDatapoints() *DescribeActiveMetricRuleListResponseBodyDatapoints
	SetMessage(v string) *DescribeActiveMetricRuleListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeActiveMetricRuleListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeActiveMetricRuleListResponseBody
	GetSuccess() *bool
}

type DescribeActiveMetricRuleListResponseBody struct {
	AlertList *DescribeActiveMetricRuleListResponseBodyAlertList `json:"AlertList,omitempty" xml:"AlertList,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code       *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Datapoints *DescribeActiveMetricRuleListResponseBodyDatapoints `json:"Datapoints,omitempty" xml:"Datapoints,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F82E6667-7811-4BA0-842F-5B2DC42BBAAD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeActiveMetricRuleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBody) GetAlertList() *DescribeActiveMetricRuleListResponseBodyAlertList {
	return s.AlertList
}

func (s *DescribeActiveMetricRuleListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeActiveMetricRuleListResponseBody) GetDatapoints() *DescribeActiveMetricRuleListResponseBodyDatapoints {
	return s.Datapoints
}

func (s *DescribeActiveMetricRuleListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeActiveMetricRuleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeActiveMetricRuleListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeActiveMetricRuleListResponseBody) SetAlertList(v *DescribeActiveMetricRuleListResponseBodyAlertList) *DescribeActiveMetricRuleListResponseBody {
	s.AlertList = v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBody) SetCode(v string) *DescribeActiveMetricRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBody) SetDatapoints(v *DescribeActiveMetricRuleListResponseBodyDatapoints) *DescribeActiveMetricRuleListResponseBody {
	s.Datapoints = v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBody) SetMessage(v string) *DescribeActiveMetricRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBody) SetRequestId(v string) *DescribeActiveMetricRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBody) SetSuccess(v bool) *DescribeActiveMetricRuleListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBody) Validate() error {
	if s.AlertList != nil {
		if err := s.AlertList.Validate(); err != nil {
			return err
		}
	}
	if s.Datapoints != nil {
		if err := s.Datapoints.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeActiveMetricRuleListResponseBodyAlertList struct {
	Alert []*DescribeActiveMetricRuleListResponseBodyAlertListAlert `json:"Alert,omitempty" xml:"Alert,omitempty" type:"Repeated"`
}

func (s DescribeActiveMetricRuleListResponseBodyAlertList) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBodyAlertList) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertList) GetAlert() []*DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	return s.Alert
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertList) SetAlert(v []*DescribeActiveMetricRuleListResponseBodyAlertListAlert) *DescribeActiveMetricRuleListResponseBodyAlertList {
	s.Alert = v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertList) Validate() error {
	if s.Alert != nil {
		for _, item := range s.Alert {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeActiveMetricRuleListResponseBodyAlertListAlert struct {
	AlertState          *string                                                            `json:"AlertState,omitempty" xml:"AlertState,omitempty"`
	ContactGroups       *string                                                            `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	Dimensions          *string                                                            `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EffectiveInterval   *string                                                            `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	EnableState         *bool                                                              `json:"EnableState,omitempty" xml:"EnableState,omitempty"`
	Escalations         *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	MailSubject         *string                                                            `json:"MailSubject,omitempty" xml:"MailSubject,omitempty"`
	MetricName          *string                                                            `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Namespace           *string                                                            `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NoEffectiveInterval *string                                                            `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	Period              *string                                                            `json:"Period,omitempty" xml:"Period,omitempty"`
	Resources           *string                                                            `json:"Resources,omitempty" xml:"Resources,omitempty"`
	RuleId              *string                                                            `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName            *string                                                            `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	SilenceTime         *string                                                            `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	Webhook             *string                                                            `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlert) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlert) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) GetAlertState() *string {
	return s.AlertState
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) GetContactGroups() *string {
	return s.ContactGroups
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) GetDimensions() *string {
	return s.Dimensions
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) GetEffectiveInterval() *string {
	return s.EffectiveInterval
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) GetEnableState() *bool {
	return s.EnableState
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) GetEscalations() *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations {
	return s.Escalations
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) GetMailSubject() *string {
	return s.MailSubject
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) GetNoEffectiveInterval() *string {
	return s.NoEffectiveInterval
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) GetPeriod() *string {
	return s.Period
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) GetResources() *string {
	return s.Resources
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) GetSilenceTime() *string {
	return s.SilenceTime
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) GetWebhook() *string {
	return s.Webhook
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetAlertState(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.AlertState = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetContactGroups(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.ContactGroups = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetDimensions(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.Dimensions = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetEffectiveInterval(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.EffectiveInterval = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetEnableState(v bool) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.EnableState = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetEscalations(v *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.Escalations = v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetMailSubject(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.MailSubject = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetMetricName(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.MetricName = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetNamespace(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.Namespace = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetNoEffectiveInterval(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.NoEffectiveInterval = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetPeriod(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.Period = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetResources(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.Resources = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetRuleId(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.RuleId = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetRuleName(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.RuleName = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetSilenceTime(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.SilenceTime = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetWebhook(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.Webhook = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) Validate() error {
	if s.Escalations != nil {
		if err := s.Escalations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations struct {
	Critical *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations) GetCritical() *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical {
	return s.Critical
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations) GetInfo() *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo {
	return s.Info
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations) GetWarn() *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn {
	return s.Warn
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations) SetCritical(v *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations {
	s.Critical = v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations) SetInfo(v *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations {
	s.Info = v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations) SetWarn(v *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations {
	s.Warn = v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations) Validate() error {
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

type DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *string `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) GetTimes() *string {
	return s.Times
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) SetComparisonOperator(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) SetStatistics(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) SetThreshold(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) SetTimes(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical {
	s.Times = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) Validate() error {
	return dara.Validate(s)
}

type DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *string `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) GetTimes() *string {
	return s.Times
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) SetComparisonOperator(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) SetStatistics(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) SetThreshold(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) SetTimes(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo {
	s.Times = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *string `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) GetTimes() *string {
	return s.Times
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) SetComparisonOperator(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) SetStatistics(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) SetThreshold(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) SetTimes(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn {
	s.Times = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) Validate() error {
	return dara.Validate(s)
}

type DescribeActiveMetricRuleListResponseBodyDatapoints struct {
	Alarm []*DescribeActiveMetricRuleListResponseBodyDatapointsAlarm `json:"Alarm,omitempty" xml:"Alarm,omitempty" type:"Repeated"`
}

func (s DescribeActiveMetricRuleListResponseBodyDatapoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBodyDatapoints) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapoints) GetAlarm() []*DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	return s.Alarm
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapoints) SetAlarm(v []*DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) *DescribeActiveMetricRuleListResponseBodyDatapoints {
	s.Alarm = v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapoints) Validate() error {
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

type DescribeActiveMetricRuleListResponseBodyDatapointsAlarm struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	ContactGroups      *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	Enable             *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	EndTime            *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EvaluationCount    *string `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	MetricName         *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Namespace          *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Period             *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RuleId             *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName           *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	SilenceTime        *string `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	StartTime          *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State              *string `json:"State,omitempty" xml:"State,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Webhook            *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GetContactGroups() *string {
	return s.ContactGroups
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GetEnable() *string {
	return s.Enable
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GetEvaluationCount() *string {
	return s.EvaluationCount
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GetPeriod() *string {
	return s.Period
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GetSilenceTime() *string {
	return s.SilenceTime
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GetState() *string {
	return s.State
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GetWebhook() *string {
	return s.Webhook
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetComparisonOperator(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetContactGroups(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.ContactGroups = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetEnable(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.Enable = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetEndTime(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.EndTime = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetEvaluationCount(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetMetricName(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.MetricName = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetNamespace(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.Namespace = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetPeriod(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.Period = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetRuleId(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.RuleId = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetRuleName(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.RuleName = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetSilenceTime(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.SilenceTime = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetStartTime(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.StartTime = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetState(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.State = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetStatistics(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.Statistics = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetThreshold(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.Threshold = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetWebhook(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.Webhook = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) Validate() error {
	return dara.Validate(s)
}
