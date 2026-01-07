// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAsyncEmailCaptchaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *SendAsyncEmailCaptchaRequest
	GetAppName() *string
	SetContactInfo(v string) *SendAsyncEmailCaptchaRequest
	GetContactInfo() *string
	SetContactorId(v string) *SendAsyncEmailCaptchaRequest
	GetContactorId() *string
}

type SendAsyncEmailCaptchaRequest struct {
	// example:
	//
	// xxx
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// xxx
	ContactInfo *string `json:"ContactInfo,omitempty" xml:"ContactInfo,omitempty"`
	// example:
	//
	// xxx
	ContactorId *string `json:"ContactorId,omitempty" xml:"ContactorId,omitempty"`
}

func (s SendAsyncEmailCaptchaRequest) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncEmailCaptchaRequest) GoString() string {
	return s.String()
}

func (s *SendAsyncEmailCaptchaRequest) GetAppName() *string {
	return s.AppName
}

func (s *SendAsyncEmailCaptchaRequest) GetContactInfo() *string {
	return s.ContactInfo
}

func (s *SendAsyncEmailCaptchaRequest) GetContactorId() *string {
	return s.ContactorId
}

func (s *SendAsyncEmailCaptchaRequest) SetAppName(v string) *SendAsyncEmailCaptchaRequest {
	s.AppName = &v
	return s
}

func (s *SendAsyncEmailCaptchaRequest) SetContactInfo(v string) *SendAsyncEmailCaptchaRequest {
	s.ContactInfo = &v
	return s
}

func (s *SendAsyncEmailCaptchaRequest) SetContactorId(v string) *SendAsyncEmailCaptchaRequest {
	s.ContactorId = &v
	return s
}

func (s *SendAsyncEmailCaptchaRequest) Validate() error {
	return dara.Validate(s)
}
