// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v int64) *CreateOrUpdateContactRequest
	GetContactId() *int64
	SetContactName(v string) *CreateOrUpdateContactRequest
	GetContactName() *string
	SetCorpUserId(v string) *CreateOrUpdateContactRequest
	GetCorpUserId() *string
	SetDingRobotUrl(v string) *CreateOrUpdateContactRequest
	GetDingRobotUrl() *string
	SetEmail(v string) *CreateOrUpdateContactRequest
	GetEmail() *string
	SetIsEmailVerify(v bool) *CreateOrUpdateContactRequest
	GetIsEmailVerify() *bool
	SetPhone(v string) *CreateOrUpdateContactRequest
	GetPhone() *string
	SetReissueSendNotice(v int64) *CreateOrUpdateContactRequest
	GetReissueSendNotice() *int64
	SetResourceGroupId(v string) *CreateOrUpdateContactRequest
	GetResourceGroupId() *string
}

type CreateOrUpdateContactRequest struct {
	// The ID of the alert contact.
	//
	// 	- If you do not specify this parameter, a new alert contact is created.
	//
	// 	- If you specify this parameter, the specified alert contact is modified.
	//
	// example:
	//
	// 123
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The name of the alert contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// JohnDoe
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The ID of the alert contact that is shown to the enterprise when the contact is mentioned with the at sign (@) by a third-party instant messaging (IM) tool.
	//
	// example:
	//
	// A123221
	CorpUserId *string `json:"CorpUserId,omitempty" xml:"CorpUserId,omitempty"`
	// The webhook URL of the DingTalk chatbot.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=69d4e0******
	DingRobotUrl *string `json:"DingRobotUrl,omitempty" xml:"DingRobotUrl,omitempty"`
	// The email address of the alert contact.
	//
	// > You must specify at least one of the **Phone*	- and **Email*	- parameters. Each mobile number or email address can be used for only one alert contact.
	//
	// example:
	//
	// someone@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Specifies whether the email address is verified.
	//
	// example:
	//
	// true
	IsEmailVerify *bool `json:"IsEmailVerify,omitempty" xml:"IsEmailVerify,omitempty"`
	// The mobile number of the alert contact.
	//
	// > You must specify at least one of the **Phone*	- and **Email*	- parameters. Each mobile number or email address can be used for only one alert contact.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The operation that you want to perform if phone calls fail to be answered. Valid values:
	//
	// 	- 0: No operation is performed.
	//
	// 	- 1: A phone call is made again.
	//
	// 	- 2: A text message is sent.
	//
	// 	- 3 (default value): The global default value is used.
	//
	// example:
	//
	// 3
	ReissueSendNotice *int64 `json:"ReissueSendNotice,omitempty" xml:"ReissueSendNotice,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateOrUpdateContactRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateContactRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateContactRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *CreateOrUpdateContactRequest) GetContactName() *string {
	return s.ContactName
}

func (s *CreateOrUpdateContactRequest) GetCorpUserId() *string {
	return s.CorpUserId
}

func (s *CreateOrUpdateContactRequest) GetDingRobotUrl() *string {
	return s.DingRobotUrl
}

func (s *CreateOrUpdateContactRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateOrUpdateContactRequest) GetIsEmailVerify() *bool {
	return s.IsEmailVerify
}

func (s *CreateOrUpdateContactRequest) GetPhone() *string {
	return s.Phone
}

func (s *CreateOrUpdateContactRequest) GetReissueSendNotice() *int64 {
	return s.ReissueSendNotice
}

func (s *CreateOrUpdateContactRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateOrUpdateContactRequest) SetContactId(v int64) *CreateOrUpdateContactRequest {
	s.ContactId = &v
	return s
}

func (s *CreateOrUpdateContactRequest) SetContactName(v string) *CreateOrUpdateContactRequest {
	s.ContactName = &v
	return s
}

func (s *CreateOrUpdateContactRequest) SetCorpUserId(v string) *CreateOrUpdateContactRequest {
	s.CorpUserId = &v
	return s
}

func (s *CreateOrUpdateContactRequest) SetDingRobotUrl(v string) *CreateOrUpdateContactRequest {
	s.DingRobotUrl = &v
	return s
}

func (s *CreateOrUpdateContactRequest) SetEmail(v string) *CreateOrUpdateContactRequest {
	s.Email = &v
	return s
}

func (s *CreateOrUpdateContactRequest) SetIsEmailVerify(v bool) *CreateOrUpdateContactRequest {
	s.IsEmailVerify = &v
	return s
}

func (s *CreateOrUpdateContactRequest) SetPhone(v string) *CreateOrUpdateContactRequest {
	s.Phone = &v
	return s
}

func (s *CreateOrUpdateContactRequest) SetReissueSendNotice(v int64) *CreateOrUpdateContactRequest {
	s.ReissueSendNotice = &v
	return s
}

func (s *CreateOrUpdateContactRequest) SetResourceGroupId(v string) *CreateOrUpdateContactRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateOrUpdateContactRequest) Validate() error {
	return dara.Validate(s)
}
