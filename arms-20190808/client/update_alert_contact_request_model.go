// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v int64) *UpdateAlertContactRequest
	GetContactId() *int64
	SetContactName(v string) *UpdateAlertContactRequest
	GetContactName() *string
	SetDingRobotWebhookUrl(v string) *UpdateAlertContactRequest
	GetDingRobotWebhookUrl() *string
	SetEmail(v string) *UpdateAlertContactRequest
	GetEmail() *string
	SetPhoneNum(v string) *UpdateAlertContactRequest
	GetPhoneNum() *string
	SetRegionId(v string) *UpdateAlertContactRequest
	GetRegionId() *string
	SetSystemNoc(v bool) *UpdateAlertContactRequest
	GetSystemNoc() *bool
}

type UpdateAlertContactRequest struct {
	// The ID of the alert contact to be updated. You can call the SearchAlertContact operation to query the contact ID. For more information, see [SearchAlertContact](https://help.aliyun.com/document_detail/130703.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The new name of the alert contact.
	//
	// example:
	//
	// John Doe
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The new webhook URL of the DingTalk chatbot. For more information, see [Configure a DingTalk chatbot to send alert notifications](https://help.aliyun.com/document_detail/106247.html). You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
	//
	// >  If you do not specify this parameter, the original parameter value is deleted. If you specify this parameter, the original parameter value is updated.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=91f2f6****
	DingRobotWebhookUrl *string `json:"DingRobotWebhookUrl,omitempty" xml:"DingRobotWebhookUrl,omitempty"`
	// The new email address of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
	//
	// >  If you do not specify this parameter, the original parameter value is deleted. If you specify this parameter, the original parameter value is updated.
	//
	// example:
	//
	// someone@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The new mobile phone number of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
	//
	// >  If you do not specify this parameter, the original parameter value is deleted. If you specify this parameter, the original parameter value is updated.
	//
	// example:
	//
	// 1381111****
	PhoneNum *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	// The ID of the region. Set the value to `cn-hangzhou`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether the alert contact receives system notifications. Valid values:
	//
	// 	- `true`: The alert contact receives system notifications.
	//
	// 	- `false`: The alert contact does not receive system notifications.
	//
	// example:
	//
	// true
	SystemNoc *bool `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
}

func (s UpdateAlertContactRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertContactRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *UpdateAlertContactRequest) GetContactName() *string {
	return s.ContactName
}

func (s *UpdateAlertContactRequest) GetDingRobotWebhookUrl() *string {
	return s.DingRobotWebhookUrl
}

func (s *UpdateAlertContactRequest) GetEmail() *string {
	return s.Email
}

func (s *UpdateAlertContactRequest) GetPhoneNum() *string {
	return s.PhoneNum
}

func (s *UpdateAlertContactRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAlertContactRequest) GetSystemNoc() *bool {
	return s.SystemNoc
}

func (s *UpdateAlertContactRequest) SetContactId(v int64) *UpdateAlertContactRequest {
	s.ContactId = &v
	return s
}

func (s *UpdateAlertContactRequest) SetContactName(v string) *UpdateAlertContactRequest {
	s.ContactName = &v
	return s
}

func (s *UpdateAlertContactRequest) SetDingRobotWebhookUrl(v string) *UpdateAlertContactRequest {
	s.DingRobotWebhookUrl = &v
	return s
}

func (s *UpdateAlertContactRequest) SetEmail(v string) *UpdateAlertContactRequest {
	s.Email = &v
	return s
}

func (s *UpdateAlertContactRequest) SetPhoneNum(v string) *UpdateAlertContactRequest {
	s.PhoneNum = &v
	return s
}

func (s *UpdateAlertContactRequest) SetRegionId(v string) *UpdateAlertContactRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertContactRequest) SetSystemNoc(v bool) *UpdateAlertContactRequest {
	s.SystemNoc = &v
	return s
}

func (s *UpdateAlertContactRequest) Validate() error {
	return dara.Validate(s)
}
