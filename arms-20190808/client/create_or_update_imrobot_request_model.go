// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateIMRobotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCardTemplate(v string) *CreateOrUpdateIMRobotRequest
	GetCardTemplate() *string
	SetDailyNoc(v bool) *CreateOrUpdateIMRobotRequest
	GetDailyNoc() *bool
	SetDailyNocTime(v string) *CreateOrUpdateIMRobotRequest
	GetDailyNocTime() *string
	SetDingSignKey(v string) *CreateOrUpdateIMRobotRequest
	GetDingSignKey() *string
	SetEnableOutgoing(v bool) *CreateOrUpdateIMRobotRequest
	GetEnableOutgoing() *bool
	SetRobotAddress(v string) *CreateOrUpdateIMRobotRequest
	GetRobotAddress() *string
	SetRobotId(v int64) *CreateOrUpdateIMRobotRequest
	GetRobotId() *int64
	SetRobotName(v string) *CreateOrUpdateIMRobotRequest
	GetRobotName() *string
	SetToken(v string) *CreateOrUpdateIMRobotRequest
	GetToken() *string
	SetType(v string) *CreateOrUpdateIMRobotRequest
	GetType() *string
}

type CreateOrUpdateIMRobotRequest struct {
	// The configurations of the alert card template. For more information about the parameters in the template, see the following section.
	//
	// example:
	//
	// {     "button": [         "claim",         "close",         "follow",         "send_itsm",         "block",         "unResolvedIncident"     ],     "field": [         {             "fieldName": "alarmName",             "visible": true         },         {             "fieldName": "notificationPolicy",             "visible": true         },         {             "fieldName": "alarmContent",             "visible": true         },         {             "fieldName": "alarmTime",             "visible": true         },         {             "fieldName": "seriesChart",             "visible": true         },         {             "fieldName": "includeEvent",             "visible": true         },         {             "fieldName": "assigned",             "visible": true         },         {             "fieldName": "similarAlarm",             "visible": true         },         {             "fieldName": "operator",             "visible": true         }     ] }
	CardTemplate *string `json:"CardTemplate,omitempty" xml:"CardTemplate,omitempty"`
	// Specifies whether to send daily statistics. Valid values:
	//
	// 	- `false` (default): Daily statistics are not sent.
	//
	// 	- `true`: Daily statistics are sent. If you set the value to `true`, the **DailyNocTime*	- parameter is required.
	//
	// example:
	//
	// true
	DailyNoc *bool `json:"DailyNoc,omitempty" xml:"DailyNoc,omitempty"`
	// The points in time at which the daily statistics are sent. Separate multiple points in time with commas (,). The points in time are in the HH:SS format. The information that ARMS sends at the specified points in time includes the total number of alerts generated on the current day, the number of cleared alerts, and the number of alerts to be cleared.
	//
	// example:
	//
	// 09:30,17:00
	DailyNocTime *string `json:"DailyNocTime,omitempty" xml:"DailyNocTime,omitempty"`
	// The signature key of DingTalk. If you specify a signature key, DingTalk authentication is performed by using the signature key. If you do not specify a signature key, a whitelist is used for authentication by default. The keyword of the whitelist is **Alert**.
	//
	// example:
	//
	// ******
	DingSignKey *string `json:"DingSignKey,omitempty" xml:"DingSignKey,omitempty"`
	// Specifies whether to enable the Outgoing feature.
	//
	// example:
	//
	// true
	EnableOutgoing *bool `json:"EnableOutgoing,omitempty" xml:"EnableOutgoing,omitempty"`
	// The webhook URL of the IM chatbot.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=e1a049121******
	RobotAddress *string `json:"RobotAddress,omitempty" xml:"RobotAddress,omitempty"`
	// The ID of the IM chatbot.
	//
	// > If you do not specify the parameter, a new IM chatbot is created.
	//
	// example:
	//
	// 123
	RobotId *int64 `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
	// The name of the IM chatbot.
	//
	// This parameter is required.
	//
	// example:
	//
	// Chatbot name
	RobotName *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	// The token required to enable the Outgoing feature.
	//
	// example:
	//
	// 1656558719183be1245ab44********
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The type of the IM chatbot. Valid values:
	//
	// 	- `dingding`: DingTalk chatbot
	//
	// 	- `wechat`: WeCom chatbot
	//
	// 	- `feishu`: Lark chatbot
	//
	// This parameter is required.
	//
	// example:
	//
	// dingding
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateOrUpdateIMRobotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateIMRobotRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateIMRobotRequest) GetCardTemplate() *string {
	return s.CardTemplate
}

func (s *CreateOrUpdateIMRobotRequest) GetDailyNoc() *bool {
	return s.DailyNoc
}

func (s *CreateOrUpdateIMRobotRequest) GetDailyNocTime() *string {
	return s.DailyNocTime
}

func (s *CreateOrUpdateIMRobotRequest) GetDingSignKey() *string {
	return s.DingSignKey
}

func (s *CreateOrUpdateIMRobotRequest) GetEnableOutgoing() *bool {
	return s.EnableOutgoing
}

func (s *CreateOrUpdateIMRobotRequest) GetRobotAddress() *string {
	return s.RobotAddress
}

func (s *CreateOrUpdateIMRobotRequest) GetRobotId() *int64 {
	return s.RobotId
}

func (s *CreateOrUpdateIMRobotRequest) GetRobotName() *string {
	return s.RobotName
}

func (s *CreateOrUpdateIMRobotRequest) GetToken() *string {
	return s.Token
}

func (s *CreateOrUpdateIMRobotRequest) GetType() *string {
	return s.Type
}

func (s *CreateOrUpdateIMRobotRequest) SetCardTemplate(v string) *CreateOrUpdateIMRobotRequest {
	s.CardTemplate = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) SetDailyNoc(v bool) *CreateOrUpdateIMRobotRequest {
	s.DailyNoc = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) SetDailyNocTime(v string) *CreateOrUpdateIMRobotRequest {
	s.DailyNocTime = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) SetDingSignKey(v string) *CreateOrUpdateIMRobotRequest {
	s.DingSignKey = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) SetEnableOutgoing(v bool) *CreateOrUpdateIMRobotRequest {
	s.EnableOutgoing = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) SetRobotAddress(v string) *CreateOrUpdateIMRobotRequest {
	s.RobotAddress = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) SetRobotId(v int64) *CreateOrUpdateIMRobotRequest {
	s.RobotId = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) SetRobotName(v string) *CreateOrUpdateIMRobotRequest {
	s.RobotName = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) SetToken(v string) *CreateOrUpdateIMRobotRequest {
	s.Token = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) SetType(v string) *CreateOrUpdateIMRobotRequest {
	s.Type = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) Validate() error {
	return dara.Validate(s)
}
