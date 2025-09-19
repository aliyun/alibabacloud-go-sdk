// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateDingTalkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigList(v string) *CreateOrUpdateDingTalkRequest
	GetConfigList() *string
	SetDingTalkLang(v string) *CreateOrUpdateDingTalkRequest
	GetDingTalkLang() *string
	SetGroupIdList(v string) *CreateOrUpdateDingTalkRequest
	GetGroupIdList() *string
	SetId(v int64) *CreateOrUpdateDingTalkRequest
	GetId() *int64
	SetIntervalTime(v int64) *CreateOrUpdateDingTalkRequest
	GetIntervalTime() *int64
	SetRuleActionName(v string) *CreateOrUpdateDingTalkRequest
	GetRuleActionName() *string
	SetSendUrl(v string) *CreateOrUpdateDingTalkRequest
	GetSendUrl() *string
}

type CreateOrUpdateDingTalkRequest struct {
	// The alerts for which you want the chatbot to send notifications. The value is a JSON array that contains the following fields:
	//
	// 	- **type**: the types of alerts. The valid values are listed in the "Additional description of parameters" section in this topic.
	//
	// 	- **configItemList**: the list of check items. The value is a JSON array that contains the following fields:
	//
	//     	- **key**: the key of the check item.
	//
	//     	- **valueList**: the values of the check item. The value of valueList is a JSON array.
	//
	// > For more information about the value of this parameter, see the "Addition description of parameters" section in this topic.
	//
	// example:
	//
	// [{"type":"sas_analysis_online-sas-operation-log-sas-event-suspicious","configItemList":[{"key":"item_level","valueList":["all"]},{"key":"event_type","valueList":["all"]}]}]
	ConfigList *string `json:"ConfigList,omitempty" xml:"ConfigList,omitempty"`
	// The language of the notifications. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	DingTalkLang *string `json:"DingTalkLang,omitempty" xml:"DingTalkLang,omitempty"`
	// The IDs of asset groups for which you want the chatbot to send notifications. The value is a JSON array.
	//
	// > You can call the [DescribeGroupStruct](~~DescribeGroupStruct~~) operation to query the IDs of asset groups.
	//
	// example:
	//
	// ["10417151"]
	GroupIdList *string `json:"GroupIdList,omitempty" xml:"GroupIdList,omitempty"`
	// The ID of the chatbot.
	//
	// > You can call the [DescribeDingTalk](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-describedingtalk/?spm=a2c63.p38356.0.0.681e4360Qd1eb1) operation to query the IDs of chatbots.
	//
	// example:
	//
	// 1589
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time interval at which the chatbot sends notifications.
	//
	// > The value **0*	- indicates unlimited.
	//
	// example:
	//
	// 0
	IntervalTime *int64 `json:"IntervalTime,omitempty" xml:"IntervalTime,omitempty"`
	// The name of the chatbot.
	//
	// > The name of a chatbot must be 2 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testNotify
	RuleActionName *string `json:"RuleActionName,omitempty" xml:"RuleActionName,omitempty"`
	// The webhook URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=XXX
	SendUrl *string `json:"SendUrl,omitempty" xml:"SendUrl,omitempty"`
}

func (s CreateOrUpdateDingTalkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateDingTalkRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateDingTalkRequest) GetConfigList() *string {
	return s.ConfigList
}

func (s *CreateOrUpdateDingTalkRequest) GetDingTalkLang() *string {
	return s.DingTalkLang
}

func (s *CreateOrUpdateDingTalkRequest) GetGroupIdList() *string {
	return s.GroupIdList
}

func (s *CreateOrUpdateDingTalkRequest) GetId() *int64 {
	return s.Id
}

func (s *CreateOrUpdateDingTalkRequest) GetIntervalTime() *int64 {
	return s.IntervalTime
}

func (s *CreateOrUpdateDingTalkRequest) GetRuleActionName() *string {
	return s.RuleActionName
}

func (s *CreateOrUpdateDingTalkRequest) GetSendUrl() *string {
	return s.SendUrl
}

func (s *CreateOrUpdateDingTalkRequest) SetConfigList(v string) *CreateOrUpdateDingTalkRequest {
	s.ConfigList = &v
	return s
}

func (s *CreateOrUpdateDingTalkRequest) SetDingTalkLang(v string) *CreateOrUpdateDingTalkRequest {
	s.DingTalkLang = &v
	return s
}

func (s *CreateOrUpdateDingTalkRequest) SetGroupIdList(v string) *CreateOrUpdateDingTalkRequest {
	s.GroupIdList = &v
	return s
}

func (s *CreateOrUpdateDingTalkRequest) SetId(v int64) *CreateOrUpdateDingTalkRequest {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateDingTalkRequest) SetIntervalTime(v int64) *CreateOrUpdateDingTalkRequest {
	s.IntervalTime = &v
	return s
}

func (s *CreateOrUpdateDingTalkRequest) SetRuleActionName(v string) *CreateOrUpdateDingTalkRequest {
	s.RuleActionName = &v
	return s
}

func (s *CreateOrUpdateDingTalkRequest) SetSendUrl(v string) *CreateOrUpdateDingTalkRequest {
	s.SendUrl = &v
	return s
}

func (s *CreateOrUpdateDingTalkRequest) Validate() error {
	return dara.Validate(s)
}
