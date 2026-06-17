// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertLogListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLogList(v []*DescribeAlertLogListResponseBodyAlertLogList) *DescribeAlertLogListResponseBody
	GetAlertLogList() []*DescribeAlertLogListResponseBodyAlertLogList
	SetCode(v string) *DescribeAlertLogListResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeAlertLogListResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeAlertLogListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAlertLogListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAlertLogListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAlertLogListResponseBody
	GetSuccess() *bool
}

type DescribeAlertLogListResponseBody struct {
	// The list of alert history entries.
	AlertLogList []*DescribeAlertLogListResponseBodyAlertLogList `json:"AlertLogList,omitempty" xml:"AlertLogList,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// > The status code 200 indicates that the call was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1C4A3709-BF52-42EE-87B5-7435F0929585
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertLogListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListResponseBody) GetAlertLogList() []*DescribeAlertLogListResponseBodyAlertLogList {
	return s.AlertLogList
}

func (s *DescribeAlertLogListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAlertLogListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertLogListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAlertLogListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAlertLogListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertLogListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAlertLogListResponseBody) SetAlertLogList(v []*DescribeAlertLogListResponseBodyAlertLogList) *DescribeAlertLogListResponseBody {
	s.AlertLogList = v
	return s
}

func (s *DescribeAlertLogListResponseBody) SetCode(v string) *DescribeAlertLogListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertLogListResponseBody) SetMessage(v string) *DescribeAlertLogListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertLogListResponseBody) SetPageNumber(v int32) *DescribeAlertLogListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAlertLogListResponseBody) SetPageSize(v int32) *DescribeAlertLogListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertLogListResponseBody) SetRequestId(v string) *DescribeAlertLogListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertLogListResponseBody) SetSuccess(v bool) *DescribeAlertLogListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlertLogListResponseBody) Validate() error {
	if s.AlertLogList != nil {
		for _, item := range s.AlertLogList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlertLogListResponseBodyAlertLogList struct {
	// The timestamp when the alert was triggered.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1610043776621
	AlertTime *string `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	// The details of the matched alert blacklist.
	//
	// example:
	//
	// {"id":12****,"metricProject":"acs_ecs_dashboard","userId":173651113438****,"uuid":"8410dbbd-7d30-41c5-94cb-****","name":"alert-****","productCategory":"ecs","instances":[{"instanceId":"i-m5e1qg6uo38rztr4****"}],"metrics":null,"scopeType":"USER","scopeValue":"","startTime":"0001-01-01T00:00:00Z","endTime":"9999-12-31T23:59:59.999999999+08:00","effectiveTime":null,"isEnable":true,"status":1,"gmtCreate":"2021-11-02T16:35:59+08:00","gmtModified":"2021-11-02T16:35:59+08:00","loadTime":"2021-11-02T16:36:15.213072177+08:00"}
	BlackListDetail *string `json:"BlackListDetail,omitempty" xml:"BlackListDetail,omitempty"`
	// The name of the matched alert blacklist.
	//
	// example:
	//
	// Black_Test
	BlackListName *string `json:"BlackListName,omitempty" xml:"BlackListName,omitempty"`
	// The UUID of the matched alert blacklist.
	//
	// example:
	//
	// 8410dbbd-7d30-41c5-94cb-****
	BlackListUUID *string `json:"BlackListUUID,omitempty" xml:"BlackListUUID,omitempty"`
	// The list of Wangwang IDs of the alert contact.
	ContactALIIWWList []*string `json:"ContactALIIWWList,omitempty" xml:"ContactALIIWWList,omitempty" type:"Repeated"`
	// The list of DingTalk accounts of the alert contact.
	ContactDingList []*string `json:"ContactDingList,omitempty" xml:"ContactDingList,omitempty" type:"Repeated"`
	// The list of alert contact groups.
	ContactGroups []*string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Repeated"`
	// The list of email addresses of the alert contact.
	ContactMailList []*string `json:"ContactMailList,omitempty" xml:"ContactMailList,omitempty" type:"Repeated"`
	// The list of phone numbers of the alert contact.
	ContactOnCallList []*string `json:"ContactOnCallList,omitempty" xml:"ContactOnCallList,omitempty" type:"Repeated"`
	// The list of phone numbers that receive text messages of the alert contact.
	ContactSMSList []*string `json:"ContactSMSList,omitempty" xml:"ContactSMSList,omitempty" type:"Repeated"`
	// The dimensions of the resource for which the alert is triggered.
	Dimensions []*DescribeAlertLogListResponseBodyAlertLogListDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The list of webhook URLs of DingTalk chatbots for the alert contact.
	DingdingWebhookList []*string `json:"DingdingWebhookList,omitempty" xml:"DingdingWebhookList,omitempty" type:"Repeated"`
	// The rule that triggers the alert.
	Escalation *DescribeAlertLogListResponseBodyAlertLogListEscalation `json:"Escalation,omitempty" xml:"Escalation,omitempty" type:"Struct"`
	// The name of the event.
	//
	// example:
	//
	// IOHang
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The extended information of the alert.
	ExtendedInfo []*DescribeAlertLogListResponseBodyAlertLogListExtendedInfo `json:"ExtendedInfo,omitempty" xml:"ExtendedInfo,omitempty" type:"Repeated"`
	// The ID of the application group.
	//
	// example:
	//
	// 7301****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the application group.
	//
	// example:
	//
	// ECS_Instances
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// i-m5e1qg6uo38rztr4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// portalHost
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The alert level and notification methods. Valid values:
	//
	// <props="china">- P2: phone calls, text messages, emails, and DingTalk chatbots.
	//
	// <props="china">- P3: text messages, emails, and DingTalk chatbots.
	//
	// <props="china">- P4: emails and DingTalk chatbots.
	//
	// <props="china">- OK: no alerts.
	//
	// <props="intl">- P4: emails and DingTalk chatbots.
	//
	// <props="intl">- OK: no alerts.
	//
	// <props="partner">- P4: emails and DingTalk chatbots.
	//
	// <props="partner">- OK: no alerts.
	//
	// example:
	//
	// P4
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The change of the alert level. Valid values:
	//
	// - `P4->OK`: The alert level changes from P4 to OK, which indicates that the alert is cleared.
	//
	// - `P4->P4`: indicates a P4-level alert.
	//
	// example:
	//
	// P4->OK
	LevelChange *string `json:"LevelChange,omitempty" xml:"LevelChange,omitempty"`
	// The log ID.
	//
	// example:
	//
	// 7510****::e8a472a0-46ae-4ac0-84b1-e46be368****
	LogId *string `json:"LogId,omitempty" xml:"LogId,omitempty"`
	// The alert-related information, which is a JSON string.
	//
	// example:
	//
	// {"alertName":"e47aa0ac-4076-44db-a47d-d1083968****_Availability"}
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The cloud service identifier. Valid values:
	//
	// - For an Alibaba Cloud service, the value is the abbreviation of the cloud service name. Example: ECS.
	//
	// - For a non-Alibaba Cloud service, the value is in the format of `acs_Product keyword`. Example: acs_networkmonitor.
	//
	// example:
	//
	// ECS
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The ID of the alert rule.
	//
	// example:
	//
	// d582b9e9-b1c1-4f17-9279-0fe7333a****_ResponseTime
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// CPU utilization
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The details of the alert pushing result.
	SendDetail *DescribeAlertLogListResponseBodyAlertLogListSendDetail `json:"SendDetail,omitempty" xml:"SendDetail,omitempty" type:"Struct"`
	// The list of alert sending results.
	SendResultList []*DescribeAlertLogListResponseBodyAlertLogListSendResultList `json:"SendResultList,omitempty" xml:"SendResultList,omitempty" type:"Repeated"`
	// The alert status. Valid values:
	//
	// - 0: An alert is triggered or cleared.
	//
	// - 1: The current time is not within the effective period of the alert.
	//
	// - 2: The current time is within the channel silence period.
	//
	// - 3: The host is being restarted.
	//
	// - 4: No alerts are sent.
	//
	// <props="china">When the alert status is 0, an alert is triggered if Level is set to P2, P3, or P4; the alert is cleared if Level is set to OK.
	//
	// <props="intl">When the alert status is 0, an alert is triggered if Level is set to P4; the alert is cleared if Level is set to OK.
	//
	// <props="partner">When the alert status is 0, an alert is triggered if Level is set to P4; the alert is cleared if Level is set to OK.
	//
	// example:
	//
	// 0
	SendStatus *string `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	// The list of URLs that are called back when the alert is triggered.
	WebhookList []*DescribeAlertLogListResponseBodyAlertLogListWebhookList `json:"WebhookList,omitempty" xml:"WebhookList,omitempty" type:"Repeated"`
}

func (s DescribeAlertLogListResponseBodyAlertLogList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogListResponseBodyAlertLogList) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetAlertTime() *string {
	return s.AlertTime
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetBlackListDetail() *string {
	return s.BlackListDetail
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetBlackListName() *string {
	return s.BlackListName
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetBlackListUUID() *string {
	return s.BlackListUUID
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetContactALIIWWList() []*string {
	return s.ContactALIIWWList
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetContactDingList() []*string {
	return s.ContactDingList
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetContactGroups() []*string {
	return s.ContactGroups
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetContactMailList() []*string {
	return s.ContactMailList
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetContactOnCallList() []*string {
	return s.ContactOnCallList
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetContactSMSList() []*string {
	return s.ContactSMSList
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetDimensions() []*DescribeAlertLogListResponseBodyAlertLogListDimensions {
	return s.Dimensions
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetDingdingWebhookList() []*string {
	return s.DingdingWebhookList
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetEscalation() *DescribeAlertLogListResponseBodyAlertLogListEscalation {
	return s.Escalation
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetEventName() *string {
	return s.EventName
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetExtendedInfo() []*DescribeAlertLogListResponseBodyAlertLogListExtendedInfo {
	return s.ExtendedInfo
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetLevel() *string {
	return s.Level
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetLevelChange() *string {
	return s.LevelChange
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetLogId() *string {
	return s.LogId
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetProduct() *string {
	return s.Product
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetSendDetail() *DescribeAlertLogListResponseBodyAlertLogListSendDetail {
	return s.SendDetail
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetSendResultList() []*DescribeAlertLogListResponseBodyAlertLogListSendResultList {
	return s.SendResultList
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetSendStatus() *string {
	return s.SendStatus
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) GetWebhookList() []*DescribeAlertLogListResponseBodyAlertLogListWebhookList {
	return s.WebhookList
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetAlertTime(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.AlertTime = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetBlackListDetail(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.BlackListDetail = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetBlackListName(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.BlackListName = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetBlackListUUID(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.BlackListUUID = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetContactALIIWWList(v []*string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.ContactALIIWWList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetContactDingList(v []*string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.ContactDingList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetContactGroups(v []*string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.ContactGroups = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetContactMailList(v []*string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.ContactMailList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetContactOnCallList(v []*string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.ContactOnCallList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetContactSMSList(v []*string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.ContactSMSList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetDimensions(v []*DescribeAlertLogListResponseBodyAlertLogListDimensions) *DescribeAlertLogListResponseBodyAlertLogList {
	s.Dimensions = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetDingdingWebhookList(v []*string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.DingdingWebhookList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetEscalation(v *DescribeAlertLogListResponseBodyAlertLogListEscalation) *DescribeAlertLogListResponseBodyAlertLogList {
	s.Escalation = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetEventName(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.EventName = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetExtendedInfo(v []*DescribeAlertLogListResponseBodyAlertLogListExtendedInfo) *DescribeAlertLogListResponseBodyAlertLogList {
	s.ExtendedInfo = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetGroupId(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.GroupId = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetGroupName(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.GroupName = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetInstanceId(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.InstanceId = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetInstanceName(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.InstanceName = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetLevel(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.Level = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetLevelChange(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.LevelChange = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetLogId(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.LogId = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetMessage(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.Message = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetMetricName(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.MetricName = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetNamespace(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.Namespace = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetProduct(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.Product = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetRuleId(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.RuleId = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetRuleName(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.RuleName = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetSendDetail(v *DescribeAlertLogListResponseBodyAlertLogListSendDetail) *DescribeAlertLogListResponseBodyAlertLogList {
	s.SendDetail = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetSendResultList(v []*DescribeAlertLogListResponseBodyAlertLogListSendResultList) *DescribeAlertLogListResponseBodyAlertLogList {
	s.SendResultList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetSendStatus(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.SendStatus = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetWebhookList(v []*DescribeAlertLogListResponseBodyAlertLogListWebhookList) *DescribeAlertLogListResponseBodyAlertLogList {
	s.WebhookList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) Validate() error {
	if s.Dimensions != nil {
		for _, item := range s.Dimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Escalation != nil {
		if err := s.Escalation.Validate(); err != nil {
			return err
		}
	}
	if s.ExtendedInfo != nil {
		for _, item := range s.ExtendedInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SendDetail != nil {
		if err := s.SendDetail.Validate(); err != nil {
			return err
		}
	}
	if s.SendResultList != nil {
		for _, item := range s.SendResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WebhookList != nil {
		for _, item := range s.WebhookList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlertLogListResponseBodyAlertLogListDimensions struct {
	// The key of the alerting resource.
	//
	// example:
	//
	// instanceId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the alerting resource.
	//
	// example:
	//
	// i-m5e1qg6uo38rztr4****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAlertLogListResponseBodyAlertLogListDimensions) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogListResponseBodyAlertLogListDimensions) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListResponseBodyAlertLogListDimensions) GetKey() *string {
	return s.Key
}

func (s *DescribeAlertLogListResponseBodyAlertLogListDimensions) GetValue() *string {
	return s.Value
}

func (s *DescribeAlertLogListResponseBodyAlertLogListDimensions) SetKey(v string) *DescribeAlertLogListResponseBodyAlertLogListDimensions {
	s.Key = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListDimensions) SetValue(v string) *DescribeAlertLogListResponseBodyAlertLogListDimensions {
	s.Value = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListDimensions) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertLogListResponseBodyAlertLogListEscalation struct {
	// The description of the rule that triggers the alert.
	//
	// > The body of the alert rule. An alert rule is triggered when the monitoring data meets the alert conditions.
	//
	// example:
	//
	// $Average<90
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The alert level and notification methods. Valid values:
	//
	// <props="china">- P2: phone calls, text messages, emails, and DingTalk chatbots.
	//
	// <props="china">- P3: text messages, emails, and DingTalk chatbots.
	//
	// <props="china">- P4: emails and DingTalk chatbots.
	//
	// <props="china">- OK: no alerts.
	//
	// <props="intl">- P4: emails and DingTalk chatbots.
	//
	// <props="intl">- OK: no alerts.
	//
	// <props="partner">- P4: emails and DingTalk chatbots.
	//
	// <props="partner">- OK: no alerts.
	//
	// example:
	//
	// P4
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The number of times that the alert is retried.
	//
	// example:
	//
	// 1
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeAlertLogListResponseBodyAlertLogListEscalation) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogListResponseBodyAlertLogListEscalation) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListResponseBodyAlertLogListEscalation) GetExpression() *string {
	return s.Expression
}

func (s *DescribeAlertLogListResponseBodyAlertLogListEscalation) GetLevel() *string {
	return s.Level
}

func (s *DescribeAlertLogListResponseBodyAlertLogListEscalation) GetTimes() *int32 {
	return s.Times
}

func (s *DescribeAlertLogListResponseBodyAlertLogListEscalation) SetExpression(v string) *DescribeAlertLogListResponseBodyAlertLogListEscalation {
	s.Expression = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListEscalation) SetLevel(v string) *DescribeAlertLogListResponseBodyAlertLogListEscalation {
	s.Level = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListEscalation) SetTimes(v int32) *DescribeAlertLogListResponseBodyAlertLogListEscalation {
	s.Times = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListEscalation) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertLogListResponseBodyAlertLogListExtendedInfo struct {
	// The name of the extension field.
	//
	// example:
	//
	// userId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the extension field.
	//
	// example:
	//
	// 100931896542****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAlertLogListResponseBodyAlertLogListExtendedInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogListResponseBodyAlertLogListExtendedInfo) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListResponseBodyAlertLogListExtendedInfo) GetName() *string {
	return s.Name
}

func (s *DescribeAlertLogListResponseBodyAlertLogListExtendedInfo) GetValue() *string {
	return s.Value
}

func (s *DescribeAlertLogListResponseBodyAlertLogListExtendedInfo) SetName(v string) *DescribeAlertLogListResponseBodyAlertLogListExtendedInfo {
	s.Name = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListExtendedInfo) SetValue(v string) *DescribeAlertLogListResponseBodyAlertLogListExtendedInfo {
	s.Value = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListExtendedInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertLogListResponseBodyAlertLogListSendDetail struct {
	// The list of alert pushing results by alert channel.
	ChannelResultList []*DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultList `json:"ChannelResultList,omitempty" xml:"ChannelResultList,omitempty" type:"Repeated"`
	// The pushing status of the alert information.
	//
	// - success: The alert was pushed.
	//
	// - error code: If a configuration error occurs and the pushing list is empty, an error code is displayed.
	//
	// example:
	//
	// success
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
}

func (s DescribeAlertLogListResponseBodyAlertLogListSendDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogListResponseBodyAlertLogListSendDetail) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetail) GetChannelResultList() []*DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultList {
	return s.ChannelResultList
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetail) GetResultCode() *string {
	return s.ResultCode
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetail) SetChannelResultList(v []*DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultList) *DescribeAlertLogListResponseBodyAlertLogListSendDetail {
	s.ChannelResultList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetail) SetResultCode(v string) *DescribeAlertLogListResponseBodyAlertLogListSendDetail {
	s.ResultCode = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetail) Validate() error {
	if s.ChannelResultList != nil {
		for _, item := range s.ChannelResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultList struct {
	// The alert pushing channel. Valid values:
	//
	// - MAIL: email.
	//
	// - SMS: text message.
	//
	// - WEBHOOK: alert callback.
	//
	// - SLS: Log Service.
	//
	// - ONCALL: phone call.
	//
	// - FC: Function Compute.
	//
	// - MNS: Message Service (MNS).
	//
	// example:
	//
	// MAIL
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The list of alert information results that CloudMonitor sends to the alert channel.
	ResultList []*DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
}

func (s DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultList) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultList) GetChannel() *string {
	return s.Channel
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultList) GetResultList() []*DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList {
	return s.ResultList
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultList) SetChannel(v string) *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultList {
	s.Channel = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultList) SetResultList(v []*DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList) *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultList {
	s.ResultList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultList) Validate() error {
	if s.ResultList != nil {
		for _, item := range s.ResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList struct {
	// The status code.
	//
	// - If `Channel` is set to `WEBHOOK`, the status code is 200 or 500.
	//
	// - If `Channel` is set to `MAIL`, `SMS`, `SLS`, `ONCALL`, `FC`, or `MNS`, this parameter is unavailable or empty.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the returned result.
	//
	// example:
	//
	// { }
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The request ID returned by calling another cloud service.
	//
	// example:
	//
	// 0BDAF8A8-04DC-5F0C-90E4-724D42C4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of calling the target.
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The list of channel notifications.
	NotifyTargetList []*string `json:"notifyTargetList,omitempty" xml:"notifyTargetList,omitempty" type:"Repeated"`
}

func (s DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList) GetCode() *string {
	return s.Code
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList) GetDetail() *string {
	return s.Detail
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList) GetNotifyTargetList() []*string {
	return s.NotifyTargetList
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList) SetCode(v string) *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList {
	s.Code = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList) SetDetail(v string) *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList {
	s.Detail = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList) SetRequestId(v string) *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList) SetSuccess(v bool) *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList {
	s.Success = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList) SetNotifyTargetList(v []*string) *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList {
	s.NotifyTargetList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultListResultList) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertLogListResponseBodyAlertLogListSendResultList struct {
	// The channel that sends the alert. Valid values:
	//
	// - MAIL: email.
	//
	// - ALIIM: Wangwang.
	//
	// - SMS: text message.
	//
	// - CALL: phone call.
	//
	// - DING: DingTalk chatbot.
	//
	// - Merged: alert combination.
	//
	// example:
	//
	// MAIL
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The notification target that corresponds to the alert channel.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeAlertLogListResponseBodyAlertLogListSendResultList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogListResponseBodyAlertLogListSendResultList) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendResultList) GetKey() *string {
	return s.Key
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendResultList) GetValue() []*string {
	return s.Value
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendResultList) SetKey(v string) *DescribeAlertLogListResponseBodyAlertLogListSendResultList {
	s.Key = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendResultList) SetValue(v []*string) *DescribeAlertLogListResponseBodyAlertLogListSendResultList {
	s.Value = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListSendResultList) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertLogListResponseBodyAlertLogListWebhookList struct {
	// The status code returned for the alert callback.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The information returned for the alert callback.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The URL that is called back when the alert is triggered.
	//
	// example:
	//
	// https://www.aliyun.com/webhook.html
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s DescribeAlertLogListResponseBodyAlertLogListWebhookList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogListResponseBodyAlertLogListWebhookList) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListResponseBodyAlertLogListWebhookList) GetCode() *string {
	return s.Code
}

func (s *DescribeAlertLogListResponseBodyAlertLogListWebhookList) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertLogListResponseBodyAlertLogListWebhookList) GetUrl() *string {
	return s.Url
}

func (s *DescribeAlertLogListResponseBodyAlertLogListWebhookList) SetCode(v string) *DescribeAlertLogListResponseBodyAlertLogListWebhookList {
	s.Code = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListWebhookList) SetMessage(v string) *DescribeAlertLogListResponseBodyAlertLogListWebhookList {
	s.Message = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListWebhookList) SetUrl(v string) *DescribeAlertLogListResponseBodyAlertLogListWebhookList {
	s.Url = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListWebhookList) Validate() error {
	return dara.Validate(s)
}
