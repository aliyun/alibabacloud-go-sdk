// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactName(v string) *CreateAlertContactRequest
	GetContactName() *string
	SetDingRobotWebhookUrl(v string) *CreateAlertContactRequest
	GetDingRobotWebhookUrl() *string
	SetEmail(v string) *CreateAlertContactRequest
	GetEmail() *string
	SetPhoneNum(v string) *CreateAlertContactRequest
	GetPhoneNum() *string
	SetRegionId(v string) *CreateAlertContactRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateAlertContactRequest
	GetResourceGroupId() *string
	SetSystemNoc(v bool) *CreateAlertContactRequest
	GetSystemNoc() *bool
}

type CreateAlertContactRequest struct {
	// The name of the alert contact.
	//
	// example:
	//
	// JohnDoe
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The webhook URL of the DingTalk chatbot. For more information about how to obtain the URL, see [Configure a DingTalk chatbot to send alert notifications](https://www.alibabacloud.com/help/zh/doc-detail/106247.htm). You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
	//
	// >  Enter `alert` in the custom keyword field of DingTalk chatbot security settings.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=91f2f6****
	DingRobotWebhookUrl *string `json:"DingRobotWebhookUrl,omitempty" xml:"DingRobotWebhookUrl,omitempty"`
	// The email address of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
	//
	// example:
	//
	// someone@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The mobile number of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
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
	// The ID of the resource group. You can obtain the resource group ID in the **Resource Management*	- console.
	//
	// example:
	//
	// rg-aek2eq4peca****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s CreateAlertContactRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertContactRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertContactRequest) GetContactName() *string {
	return s.ContactName
}

func (s *CreateAlertContactRequest) GetDingRobotWebhookUrl() *string {
	return s.DingRobotWebhookUrl
}

func (s *CreateAlertContactRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateAlertContactRequest) GetPhoneNum() *string {
	return s.PhoneNum
}

func (s *CreateAlertContactRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAlertContactRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAlertContactRequest) GetSystemNoc() *bool {
	return s.SystemNoc
}

func (s *CreateAlertContactRequest) SetContactName(v string) *CreateAlertContactRequest {
	s.ContactName = &v
	return s
}

func (s *CreateAlertContactRequest) SetDingRobotWebhookUrl(v string) *CreateAlertContactRequest {
	s.DingRobotWebhookUrl = &v
	return s
}

func (s *CreateAlertContactRequest) SetEmail(v string) *CreateAlertContactRequest {
	s.Email = &v
	return s
}

func (s *CreateAlertContactRequest) SetPhoneNum(v string) *CreateAlertContactRequest {
	s.PhoneNum = &v
	return s
}

func (s *CreateAlertContactRequest) SetRegionId(v string) *CreateAlertContactRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAlertContactRequest) SetResourceGroupId(v string) *CreateAlertContactRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAlertContactRequest) SetSystemNoc(v bool) *CreateAlertContactRequest {
	s.SystemNoc = &v
	return s
}

func (s *CreateAlertContactRequest) Validate() error {
	return dara.Validate(s)
}
