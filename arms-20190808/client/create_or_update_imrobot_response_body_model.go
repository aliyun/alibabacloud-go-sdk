// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateIMRobotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertRobot(v *CreateOrUpdateIMRobotResponseBodyAlertRobot) *CreateOrUpdateIMRobotResponseBody
	GetAlertRobot() *CreateOrUpdateIMRobotResponseBodyAlertRobot
	SetRequestId(v string) *CreateOrUpdateIMRobotResponseBody
	GetRequestId() *string
}

type CreateOrUpdateIMRobotResponseBody struct {
	// The information about the IM chatbot.
	AlertRobot *CreateOrUpdateIMRobotResponseBodyAlertRobot `json:"AlertRobot,omitempty" xml:"AlertRobot,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 16AF921B-8187-489F-9913-43C808B4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrUpdateIMRobotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateIMRobotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateIMRobotResponseBody) GetAlertRobot() *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	return s.AlertRobot
}

func (s *CreateOrUpdateIMRobotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrUpdateIMRobotResponseBody) SetAlertRobot(v *CreateOrUpdateIMRobotResponseBodyAlertRobot) *CreateOrUpdateIMRobotResponseBody {
	s.AlertRobot = v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBody) SetRequestId(v string) *CreateOrUpdateIMRobotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateIMRobotResponseBodyAlertRobot struct {
	// The configurations of the alert card template.
	//
	// example:
	//
	// {     "button": [         "claim",         "close",         "follow",         "send_itsm",         "block",         "unResolvedIncident"     ],     "field": [         {             "fieldName": "alarmName",             "visible": true         },         {             "fieldName": "notificationPolicy",             "visible": true         },         {             "fieldName": "alarmContent",             "visible": true         },         {             "fieldName": "alarmTime",             "visible": true         },         {             "fieldName": "seriesChart",             "visible": true         },         {             "fieldName": "includeEvent",             "visible": true         },         {             "fieldName": "assigned",             "visible": true         },         {             "fieldName": "similarAlarm",             "visible": true         },         {             "fieldName": "operator",             "visible": true         }     ] }
	CardTemplate *string `json:"CardTemplate,omitempty" xml:"CardTemplate,omitempty"`
	// Indicates whether daily statistics are sent. Valid values:
	//
	// 	- `false` (default): Daily statistics are not sent.
	//
	// 	- `true`: Daily statistics are sent.
	//
	// example:
	//
	// true
	DailyNoc *bool `json:"DailyNoc,omitempty" xml:"DailyNoc,omitempty"`
	// The point in time at which the daily statistics are sent. The information that ARMS sends at the specified points in time includes the total number of alerts generated on the current day, the number of cleared alerts, and the number of alerts to be cleared.
	//
	// example:
	//
	// 09:30,17:00
	DailyNocTime *string `json:"DailyNocTime,omitempty" xml:"DailyNocTime,omitempty"`
	// Indicates whether the Outgoing feature is enabled.
	//
	// example:
	//
	// true
	EnableOutgoing *bool `json:"EnableOutgoing,omitempty" xml:"EnableOutgoing,omitempty"`
	// The webhook URL of the IM chatbot.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=e1a049121******
	RobotAddress *string `json:"RobotAddress,omitempty" xml:"RobotAddress,omitempty"`
	// The ID of the IM chatbot.
	//
	// example:
	//
	// 123
	RobotId *float32 `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
	// The name of the IM chatbot.
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
	// example:
	//
	// dingding
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateOrUpdateIMRobotResponseBodyAlertRobot) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateIMRobotResponseBodyAlertRobot) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) GetCardTemplate() *string {
	return s.CardTemplate
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) GetDailyNoc() *bool {
	return s.DailyNoc
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) GetDailyNocTime() *string {
	return s.DailyNocTime
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) GetEnableOutgoing() *bool {
	return s.EnableOutgoing
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) GetRobotAddress() *string {
	return s.RobotAddress
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) GetRobotId() *float32 {
	return s.RobotId
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) GetRobotName() *string {
	return s.RobotName
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) GetToken() *string {
	return s.Token
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) GetType() *string {
	return s.Type
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetCardTemplate(v string) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.CardTemplate = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetDailyNoc(v bool) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.DailyNoc = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetDailyNocTime(v string) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.DailyNocTime = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetEnableOutgoing(v bool) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.EnableOutgoing = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetRobotAddress(v string) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.RobotAddress = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetRobotId(v float32) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.RobotId = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetRobotName(v string) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.RobotName = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetToken(v string) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.Token = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetType(v string) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.Type = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) Validate() error {
	return dara.Validate(s)
}
