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
	// The queried alert logs.
	AlertLogList []*DescribeAlertLogListResponseBodyAlertLogList `json:"AlertLogList,omitempty" xml:"AlertLogList,omitempty" type:"Repeated"`
	// The HTTP status code.
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
	// The timestamp that was generated when the alert was triggered.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1610043776621
	AlertTime *string `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	// The details of the blacklist policy.
	//
	// example:
	//
	// BlackListDetail
	BlackListDetail *string `json:"BlackListDetail,omitempty" xml:"BlackListDetail,omitempty"`
	// The name of the blacklist policy.
	//
	// example:
	//
	// {"id":123,"metricProject":"acs_ecs_dashboard","userId":1736511134389110,"uuid":"8410dbbd-7d30-41c5-94cb-***","name":"alert-***","productCategory":"ecs","instances":[{"instanceId":"host-***"}],"metrics":null,"scopeType":"USER","scopeValue":"","startTime":"0001-01-01T00:00:00Z","endTime":"9999-12-31T23:59:59.999999999+08:00","effectiveTime":null,"isEnable":true,"status":1,"gmtCreate":"2021-11-02T16:35:59+08:00","gmtModified":"2021-11-02T16:35:59+08:00","loadTime":"2021-11-02T16:36:15.213072177+08:00"}
	BlackListName *string `json:"BlackListName,omitempty" xml:"BlackListName,omitempty"`
	// The ID of the blacklist policy.
	//
	// example:
	//
	// 8410dbbd-7d30-41c5-94cb-*****
	BlackListUUID     *string   `json:"BlackListUUID,omitempty" xml:"BlackListUUID,omitempty"`
	ContactALIIWWList []*string `json:"ContactALIIWWList,omitempty" xml:"ContactALIIWWList,omitempty" type:"Repeated"`
	ContactDingList   []*string `json:"ContactDingList,omitempty" xml:"ContactDingList,omitempty" type:"Repeated"`
	ContactGroups     []*string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Repeated"`
	ContactMailList   []*string `json:"ContactMailList,omitempty" xml:"ContactMailList,omitempty" type:"Repeated"`
	ContactOnCallList []*string `json:"ContactOnCallList,omitempty" xml:"ContactOnCallList,omitempty" type:"Repeated"`
	ContactSMSList    []*string `json:"ContactSMSList,omitempty" xml:"ContactSMSList,omitempty" type:"Repeated"`
	// The dimensions of the resource that triggered alerts.
	Dimensions          []*DescribeAlertLogListResponseBodyAlertLogListDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	DingdingWebhookList []*string                                                 `json:"DingdingWebhookList,omitempty" xml:"DingdingWebhookList,omitempty" type:"Repeated"`
	// The alert rule based on which the alert is triggered.
	Escalation *DescribeAlertLogListResponseBodyAlertLogListEscalation `json:"Escalation,omitempty" xml:"Escalation,omitempty" type:"Struct"`
	// The event name.
	//
	// example:
	//
	// IOHang
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The extended fields.
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
	// The resource ID.
	//
	// example:
	//
	// i-m5e1qg6uo38rztr4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// portalHost
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The alert level and the methods that are used to send alert notifications. Valid values:
	//
	// 	- P4: Alert notifications are sent by using emails and DingTalk chatbots.
	//
	// 	- OK: No alert is generated.
	//
	// example:
	//
	// P4
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// Indicates whether the alert level was changed. Valid values:
	//
	// 	- `P4->OK`: The alert level was changed from P4 to OK.
	//
	// 	- `P4->P4`: The alert level was still P4.
	//
	// example:
	//
	// P4->OK
	LevelChange *string `json:"LevelChange,omitempty" xml:"LevelChange,omitempty"`
	// The log ID.
	//
	// example:
	//
	// 7818361[1523]@1671593992[1]
	LogId *string `json:"LogId,omitempty" xml:"LogId,omitempty"`
	// The alert information in a JSON string.
	//
	// example:
	//
	// {"alertName":"e47aa0ac-4076-44db-a47d-d1083968****_Availability"}
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The metric name.
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
	// The identifier of the cloud service. Valid values:
	//
	// 	- If the cloud service is provided by Alibaba Cloud, the abbreviation of the service name is returned. Example: ECS.
	//
	// 	- If the cloud service is not provided by Alibaba Cloud, a value in the `acs_Service keyword` format is returned. Example: acs_networkmonitor.
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
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The details about the sending results of alert notifications.
	SendDetail *DescribeAlertLogListResponseBodyAlertLogListSendDetail `json:"SendDetail,omitempty" xml:"SendDetail,omitempty" type:"Struct"`
	// The sending results of alert notifications.
	SendResultList []*DescribeAlertLogListResponseBodyAlertLogListSendResultList `json:"SendResultList,omitempty" xml:"SendResultList,omitempty" type:"Repeated"`
	// The status of the alert. Valid values:
	//
	// 	- 0: The alert is triggered or cleared.
	//
	// 	- 1: The alert is ineffective.
	//
	// 	- 2: The alert is muted.
	//
	// 	- 3: The host is restarting.
	//
	// 	- 4: No alert notification is sent.
	//
	// If the value of the SendStatus parameter is 0, the value P4 of the Level parameter indicates a triggered alert and the value OK indicates a cleared alert.
	//
	// example:
	//
	// 0
	SendStatus *string `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	// The callback URLs.
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
	// The key of the dimension.
	//
	// example:
	//
	// instanceId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
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
	// The description of the alert rule.
	//
	// >  The content of the alert rule. This parameter indicates the conditions that trigger an alert.
	//
	// example:
	//
	// $Average<90
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The alert level and the methods that are used to send alert notifications. Valid values:
	//
	// 	- P4: Alert notifications are sent by using emails and DingTalk chatbots.
	//
	// 	- OK: No alert is generated.
	//
	// example:
	//
	// P4
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The consecutive number of times for which the metric value meets the alert condition before an alert is triggered.
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
	// The name of the extended field.
	//
	// example:
	//
	// userId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the extended field.
	//
	// example:
	//
	// 120886317861****
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
	// The list of sending results that are categorized by notification method.
	ChannelResultList []*DescribeAlertLogListResponseBodyAlertLogListSendDetailChannelResultList `json:"ChannelResultList,omitempty" xml:"ChannelResultList,omitempty" type:"Repeated"`
	// Indicates whether the alert notifications are sent.
	//
	// 	- If the alert notifications are sent, the value "success" is returned.
	//
	// 	- If the configuration is invalid, no alert notification is sent and an error code is returned.
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
	// The method that is used to send alert notifications. Valid values:
	//
	// 	- MAIL: email
	//
	// 	- SMS: text message
	//
	// 	- WEBHOOK: alert callback
	//
	// 	- SLS: Simple Log Service
	//
	// 	- ONCALL: phone call
	//
	// 	- FC: Function Compute
	//
	// 	- MNS: Message Service queue
	//
	// example:
	//
	// MAIL
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The sending results of alert notifications.
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
	// The HTTP status code.
	//
	// 	- If the value of the `Channel` parameter is `WEBHOOK`, the status code is 200 or 500.
	//
	// 	- If the value of the `Channel` parameter is `MAIL`, `SMS`, `SLS`, `ONCALL`, `FC`, or `MNS`, this parameter is empty or not returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the returned results.
	//
	// example:
	//
	// { }
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The request ID returned when CloudMonitor calls another cloud service.
	//
	// example:
	//
	// 0BDAF8A8-04DC-5F0C-90E4-724D42C4****
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
	Success          *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The category of the alert notification method. Valid values:
	//
	// 	- MAIL: email
	//
	// 	- ALIIM: TradeManager
	//
	// 	- SMS: text message
	//
	// 	- CALL: phone call
	//
	// 	- DING: DingTalk chatbot
	//
	// 	- Merged: alert merging
	//
	// example:
	//
	// Mail
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The notification object corresponding to the alert notification method.
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
	// The status code of the alert callback.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The message returned for the alert callback.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The callback URL.
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
