// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertHistoryListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmHistoryList(v *DescribeAlertHistoryListResponseBodyAlarmHistoryList) *DescribeAlertHistoryListResponseBody
	GetAlarmHistoryList() *DescribeAlertHistoryListResponseBodyAlarmHistoryList
	SetCode(v string) *DescribeAlertHistoryListResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeAlertHistoryListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAlertHistoryListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAlertHistoryListResponseBody
	GetSuccess() *bool
	SetTotal(v string) *DescribeAlertHistoryListResponseBody
	GetTotal() *string
}

type DescribeAlertHistoryListResponseBody struct {
	AlarmHistoryList *DescribeAlertHistoryListResponseBodyAlarmHistoryList `json:"AlarmHistoryList,omitempty" xml:"AlarmHistoryList,omitempty" type:"Struct"`
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C3C69FBE-2262-541F-A409-C52F380BAE63
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
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAlertHistoryListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBody) GetAlarmHistoryList() *DescribeAlertHistoryListResponseBodyAlarmHistoryList {
	return s.AlarmHistoryList
}

func (s *DescribeAlertHistoryListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAlertHistoryListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertHistoryListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertHistoryListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAlertHistoryListResponseBody) GetTotal() *string {
	return s.Total
}

func (s *DescribeAlertHistoryListResponseBody) SetAlarmHistoryList(v *DescribeAlertHistoryListResponseBodyAlarmHistoryList) *DescribeAlertHistoryListResponseBody {
	s.AlarmHistoryList = v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetCode(v string) *DescribeAlertHistoryListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetMessage(v string) *DescribeAlertHistoryListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetRequestId(v string) *DescribeAlertHistoryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetSuccess(v bool) *DescribeAlertHistoryListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetTotal(v string) *DescribeAlertHistoryListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) Validate() error {
	if s.AlarmHistoryList != nil {
		if err := s.AlarmHistoryList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryList struct {
	AlarmHistory []*DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory `json:"AlarmHistory,omitempty" xml:"AlarmHistory,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryList) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryList) GetAlarmHistory() []*DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	return s.AlarmHistory
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryList) SetAlarmHistory(v []*DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) *DescribeAlertHistoryListResponseBodyAlarmHistoryList {
	s.AlarmHistory = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryList) Validate() error {
	if s.AlarmHistory != nil {
		for _, item := range s.AlarmHistory {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory struct {
	AlertTime       *int64                                                                         `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	ContactALIIMs   *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs `json:"ContactALIIMs,omitempty" xml:"ContactALIIMs,omitempty" type:"Struct"`
	ContactGroups   *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Struct"`
	ContactMails    *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails  `json:"ContactMails,omitempty" xml:"ContactMails,omitempty" type:"Struct"`
	ContactSmses    *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses  `json:"ContactSmses,omitempty" xml:"ContactSmses,omitempty" type:"Struct"`
	Contacts        *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts      `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Struct"`
	Dimensions      *string                                                                        `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EvaluationCount *int32                                                                         `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Expression      *string                                                                        `json:"Expression,omitempty" xml:"Expression,omitempty"`
	GroupId         *string                                                                        `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceName    *string                                                                        `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	LastTime        *int64                                                                         `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	Level           *string                                                                        `json:"Level,omitempty" xml:"Level,omitempty"`
	MetricName      *string                                                                        `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Namespace       *string                                                                        `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RuleId          *string                                                                        `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName        *string                                                                        `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	State           *string                                                                        `json:"State,omitempty" xml:"State,omitempty"`
	Status          *int32                                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Value           *string                                                                        `json:"Value,omitempty" xml:"Value,omitempty"`
	Webhooks        *string                                                                        `json:"Webhooks,omitempty" xml:"Webhooks,omitempty"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetAlertTime() *int64 {
	return s.AlertTime
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetContactALIIMs() *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs {
	return s.ContactALIIMs
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetContactGroups() *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups {
	return s.ContactGroups
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetContactMails() *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails {
	return s.ContactMails
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetContactSmses() *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses {
	return s.ContactSmses
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetContacts() *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts {
	return s.Contacts
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetDimensions() *string {
	return s.Dimensions
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetExpression() *string {
	return s.Expression
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetLastTime() *int64 {
	return s.LastTime
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetLevel() *string {
	return s.Level
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetState() *string {
	return s.State
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetValue() *string {
	return s.Value
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetWebhooks() *string {
	return s.Webhooks
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetAlertTime(v int64) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.AlertTime = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContactALIIMs(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactALIIMs = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContactGroups(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactGroups = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContactMails(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactMails = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContactSmses(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactSmses = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContacts(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Contacts = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetDimensions(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Dimensions = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetEvaluationCount(v int32) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetExpression(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Expression = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetGroupId(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.GroupId = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetInstanceName(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.InstanceName = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetLastTime(v int64) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.LastTime = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetLevel(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Level = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetMetricName(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.MetricName = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetNamespace(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Namespace = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetRuleId(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.RuleId = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetRuleName(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.RuleName = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetState(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.State = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetStatus(v int32) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Status = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetValue(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Value = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetWebhooks(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Webhooks = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) Validate() error {
	if s.ContactALIIMs != nil {
		if err := s.ContactALIIMs.Validate(); err != nil {
			return err
		}
	}
	if s.ContactGroups != nil {
		if err := s.ContactGroups.Validate(); err != nil {
			return err
		}
	}
	if s.ContactMails != nil {
		if err := s.ContactMails.Validate(); err != nil {
			return err
		}
	}
	if s.ContactSmses != nil {
		if err := s.ContactSmses.Validate(); err != nil {
			return err
		}
	}
	if s.Contacts != nil {
		if err := s.Contacts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs struct {
	ContactALIIM []*string `json:"ContactALIIM,omitempty" xml:"ContactALIIM,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) GetContactALIIM() []*string {
	return s.ContactALIIM
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) SetContactALIIM(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs {
	s.ContactALIIM = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups struct {
	ContactGroup []*string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) GetContactGroup() []*string {
	return s.ContactGroup
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) SetContactGroup(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups {
	s.ContactGroup = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails struct {
	ContactMail []*string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) GetContactMail() []*string {
	return s.ContactMail
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) SetContactMail(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails {
	s.ContactMail = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses struct {
	ContactSms []*string `json:"ContactSms,omitempty" xml:"ContactSms,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) GetContactSms() []*string {
	return s.ContactSms
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) SetContactSms(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses {
	s.ContactSms = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts struct {
	Contact []*string `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) GetContact() []*string {
	return s.Contact
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) SetContact(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts {
	s.Contact = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) Validate() error {
	return dara.Validate(s)
}
