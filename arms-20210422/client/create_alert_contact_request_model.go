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
	SetSystemNoc(v bool) *CreateAlertContactRequest
	GetSystemNoc() *bool
}

type CreateAlertContactRequest struct {
	ContactName         *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	DingRobotWebhookUrl *string `json:"DingRobotWebhookUrl,omitempty" xml:"DingRobotWebhookUrl,omitempty"`
	Email               *string `json:"Email,omitempty" xml:"Email,omitempty"`
	PhoneNum            *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	// This parameter is required.
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SystemNoc *bool   `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
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

func (s *CreateAlertContactRequest) SetSystemNoc(v bool) *CreateAlertContactRequest {
	s.SystemNoc = &v
	return s
}

func (s *CreateAlertContactRequest) Validate() error {
	return dara.Validate(s)
}
