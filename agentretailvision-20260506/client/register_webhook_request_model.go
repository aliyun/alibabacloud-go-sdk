// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterWebhookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackSecret(v string) *RegisterWebhookRequest
	GetCallbackSecret() *string
	SetCallbackUrl(v string) *RegisterWebhookRequest
	GetCallbackUrl() *string
}

type RegisterWebhookRequest struct {
	// example:
	//
	// your_secret_key
	CallbackSecret *string `json:"CallbackSecret,omitempty" xml:"CallbackSecret,omitempty"`
	// example:
	//
	// https://example.com/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
}

func (s RegisterWebhookRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterWebhookRequest) GoString() string {
	return s.String()
}

func (s *RegisterWebhookRequest) GetCallbackSecret() *string {
	return s.CallbackSecret
}

func (s *RegisterWebhookRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *RegisterWebhookRequest) SetCallbackSecret(v string) *RegisterWebhookRequest {
	s.CallbackSecret = &v
	return s
}

func (s *RegisterWebhookRequest) SetCallbackUrl(v string) *RegisterWebhookRequest {
	s.CallbackUrl = &v
	return s
}

func (s *RegisterWebhookRequest) Validate() error {
	return dara.Validate(s)
}
