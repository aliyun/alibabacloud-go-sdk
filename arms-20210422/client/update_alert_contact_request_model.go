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
	// This parameter is required.
	ContactId           *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName         *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	DingRobotWebhookUrl *string `json:"DingRobotWebhookUrl,omitempty" xml:"DingRobotWebhookUrl,omitempty"`
	Email               *string `json:"Email,omitempty" xml:"Email,omitempty"`
	PhoneNum            *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	// This parameter is required.
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SystemNoc *bool   `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
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
