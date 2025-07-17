// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendTTSVerifyLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v int64) *SendTTSVerifyLinkRequest
	GetContactId() *int64
	SetPhone(v string) *SendTTSVerifyLinkRequest
	GetPhone() *string
}

type SendTTSVerifyLinkRequest struct {
	// The ID of the alert contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The mobile number of the alert contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s SendTTSVerifyLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendTTSVerifyLinkRequest) GoString() string {
	return s.String()
}

func (s *SendTTSVerifyLinkRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *SendTTSVerifyLinkRequest) GetPhone() *string {
	return s.Phone
}

func (s *SendTTSVerifyLinkRequest) SetContactId(v int64) *SendTTSVerifyLinkRequest {
	s.ContactId = &v
	return s
}

func (s *SendTTSVerifyLinkRequest) SetPhone(v string) *SendTTSVerifyLinkRequest {
	s.Phone = &v
	return s
}

func (s *SendTTSVerifyLinkRequest) Validate() error {
	return dara.Validate(s)
}
