// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAsyncMobileCaptchaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *SendAsyncMobileCaptchaRequest
	GetAppName() *string
	SetContactInfo(v string) *SendAsyncMobileCaptchaRequest
	GetContactInfo() *string
	SetContactorId(v string) *SendAsyncMobileCaptchaRequest
	GetContactorId() *string
}

type SendAsyncMobileCaptchaRequest struct {
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

func (s SendAsyncMobileCaptchaRequest) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncMobileCaptchaRequest) GoString() string {
	return s.String()
}

func (s *SendAsyncMobileCaptchaRequest) GetAppName() *string {
	return s.AppName
}

func (s *SendAsyncMobileCaptchaRequest) GetContactInfo() *string {
	return s.ContactInfo
}

func (s *SendAsyncMobileCaptchaRequest) GetContactorId() *string {
	return s.ContactorId
}

func (s *SendAsyncMobileCaptchaRequest) SetAppName(v string) *SendAsyncMobileCaptchaRequest {
	s.AppName = &v
	return s
}

func (s *SendAsyncMobileCaptchaRequest) SetContactInfo(v string) *SendAsyncMobileCaptchaRequest {
	s.ContactInfo = &v
	return s
}

func (s *SendAsyncMobileCaptchaRequest) SetContactorId(v string) *SendAsyncMobileCaptchaRequest {
	s.ContactorId = &v
	return s
}

func (s *SendAsyncMobileCaptchaRequest) Validate() error {
	return dara.Validate(s)
}
