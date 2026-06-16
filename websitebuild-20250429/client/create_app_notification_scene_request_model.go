// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppNotificationSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *CreateAppNotificationSceneRequest
	GetBizId() *string
	SetChannelsJson(v string) *CreateAppNotificationSceneRequest
	GetChannelsJson() *string
	SetDescription(v string) *CreateAppNotificationSceneRequest
	GetDescription() *string
	SetEmailFieldsJson(v string) *CreateAppNotificationSceneRequest
	GetEmailFieldsJson() *string
	SetEmailLimitJson(v string) *CreateAppNotificationSceneRequest
	GetEmailLimitJson() *string
	SetEmailRecipientIdsJson(v string) *CreateAppNotificationSceneRequest
	GetEmailRecipientIdsJson() *string
	SetName(v string) *CreateAppNotificationSceneRequest
	GetName() *string
	SetPhoneRecipientIdsJson(v string) *CreateAppNotificationSceneRequest
	GetPhoneRecipientIdsJson() *string
	SetSmsFieldsJson(v string) *CreateAppNotificationSceneRequest
	GetSmsFieldsJson() *string
	SetSmsLimitJson(v string) *CreateAppNotificationSceneRequest
	GetSmsLimitJson() *string
	SetTableName(v string) *CreateAppNotificationSceneRequest
	GetTableName() *string
	SetTriggerEventsJson(v string) *CreateAppNotificationSceneRequest
	GetTriggerEventsJson() *string
}

type CreateAppNotificationSceneRequest struct {
	// The business ID.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The notification channels in a JSON array string, such as ["sms","email"].
	//
	// example:
	//
	// ["sms","email"]
	ChannelsJson *string `json:"ChannelsJson,omitempty" xml:"ChannelsJson,omitempty"`
	// The description of the scenario.
	//
	// example:
	//
	// cn_graph_prod
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The email notification fields in a JSON array string. A maximum of 10 fields are supported.
	//
	// example:
	//
	// {}
	EmailFieldsJson *string `json:"EmailFieldsJson,omitempty" xml:"EmailFieldsJson,omitempty"`
	// The email sending limit in a JSON string, including the minInterval and dailyLimit fields.
	//
	// example:
	//
	// {}
	EmailLimitJson *string `json:"EmailLimitJson,omitempty" xml:"EmailLimitJson,omitempty"`
	// The list of email recipient IDs in a JSON array string.
	//
	// example:
	//
	// {}
	EmailRecipientIdsJson *string `json:"EmailRecipientIdsJson,omitempty" xml:"EmailRecipientIdsJson,omitempty"`
	// The name of the scenario.
	//
	// example:
	//
	// 设备能力手册
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of SMS recipient IDs in a JSON array string.
	//
	// example:
	//
	// {}
	PhoneRecipientIdsJson *string `json:"PhoneRecipientIdsJson,omitempty" xml:"PhoneRecipientIdsJson,omitempty"`
	// The SMS notification fields in a JSON array string. A maximum of 3 fields are supported.
	//
	// example:
	//
	// {}
	SmsFieldsJson *string `json:"SmsFieldsJson,omitempty" xml:"SmsFieldsJson,omitempty"`
	// The SMS sending limit in a JSON string, including the minInterval and dailyLimit fields.
	//
	// example:
	//
	// {}
	SmsLimitJson *string `json:"SmsLimitJson,omitempty" xml:"SmsLimitJson,omitempty"`
	// The name of the monitored data table.
	//
	// example:
	//
	// default.ai_advertising_material_rec_train_v1103
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The trigger events in a JSON array string, such as ["INSERT","UPDATE","DELETE"].
	//
	// example:
	//
	// ["INSERT","UPDATE","DELETE"]
	TriggerEventsJson *string `json:"TriggerEventsJson,omitempty" xml:"TriggerEventsJson,omitempty"`
}

func (s CreateAppNotificationSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppNotificationSceneRequest) GoString() string {
	return s.String()
}

func (s *CreateAppNotificationSceneRequest) GetBizId() *string {
	return s.BizId
}

func (s *CreateAppNotificationSceneRequest) GetChannelsJson() *string {
	return s.ChannelsJson
}

func (s *CreateAppNotificationSceneRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAppNotificationSceneRequest) GetEmailFieldsJson() *string {
	return s.EmailFieldsJson
}

func (s *CreateAppNotificationSceneRequest) GetEmailLimitJson() *string {
	return s.EmailLimitJson
}

func (s *CreateAppNotificationSceneRequest) GetEmailRecipientIdsJson() *string {
	return s.EmailRecipientIdsJson
}

func (s *CreateAppNotificationSceneRequest) GetName() *string {
	return s.Name
}

func (s *CreateAppNotificationSceneRequest) GetPhoneRecipientIdsJson() *string {
	return s.PhoneRecipientIdsJson
}

func (s *CreateAppNotificationSceneRequest) GetSmsFieldsJson() *string {
	return s.SmsFieldsJson
}

func (s *CreateAppNotificationSceneRequest) GetSmsLimitJson() *string {
	return s.SmsLimitJson
}

func (s *CreateAppNotificationSceneRequest) GetTableName() *string {
	return s.TableName
}

func (s *CreateAppNotificationSceneRequest) GetTriggerEventsJson() *string {
	return s.TriggerEventsJson
}

func (s *CreateAppNotificationSceneRequest) SetBizId(v string) *CreateAppNotificationSceneRequest {
	s.BizId = &v
	return s
}

func (s *CreateAppNotificationSceneRequest) SetChannelsJson(v string) *CreateAppNotificationSceneRequest {
	s.ChannelsJson = &v
	return s
}

func (s *CreateAppNotificationSceneRequest) SetDescription(v string) *CreateAppNotificationSceneRequest {
	s.Description = &v
	return s
}

func (s *CreateAppNotificationSceneRequest) SetEmailFieldsJson(v string) *CreateAppNotificationSceneRequest {
	s.EmailFieldsJson = &v
	return s
}

func (s *CreateAppNotificationSceneRequest) SetEmailLimitJson(v string) *CreateAppNotificationSceneRequest {
	s.EmailLimitJson = &v
	return s
}

func (s *CreateAppNotificationSceneRequest) SetEmailRecipientIdsJson(v string) *CreateAppNotificationSceneRequest {
	s.EmailRecipientIdsJson = &v
	return s
}

func (s *CreateAppNotificationSceneRequest) SetName(v string) *CreateAppNotificationSceneRequest {
	s.Name = &v
	return s
}

func (s *CreateAppNotificationSceneRequest) SetPhoneRecipientIdsJson(v string) *CreateAppNotificationSceneRequest {
	s.PhoneRecipientIdsJson = &v
	return s
}

func (s *CreateAppNotificationSceneRequest) SetSmsFieldsJson(v string) *CreateAppNotificationSceneRequest {
	s.SmsFieldsJson = &v
	return s
}

func (s *CreateAppNotificationSceneRequest) SetSmsLimitJson(v string) *CreateAppNotificationSceneRequest {
	s.SmsLimitJson = &v
	return s
}

func (s *CreateAppNotificationSceneRequest) SetTableName(v string) *CreateAppNotificationSceneRequest {
	s.TableName = &v
	return s
}

func (s *CreateAppNotificationSceneRequest) SetTriggerEventsJson(v string) *CreateAppNotificationSceneRequest {
	s.TriggerEventsJson = &v
	return s
}

func (s *CreateAppNotificationSceneRequest) Validate() error {
	return dara.Validate(s)
}
