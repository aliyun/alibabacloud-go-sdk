// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppLlmApiKeyForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppLlmApiKeyForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppLlmApiKeyForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppLlmApiKeyForPartnerResponseBody) *CreateAppLlmApiKeyForPartnerResponse
	GetBody() *CreateAppLlmApiKeyForPartnerResponseBody
}

type CreateAppLlmApiKeyForPartnerResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppLlmApiKeyForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppLlmApiKeyForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppLlmApiKeyForPartnerResponse) GoString() string {
	return s.String()
}

func (s *CreateAppLlmApiKeyForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppLlmApiKeyForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppLlmApiKeyForPartnerResponse) GetBody() *CreateAppLlmApiKeyForPartnerResponseBody {
	return s.Body
}

func (s *CreateAppLlmApiKeyForPartnerResponse) SetHeaders(v map[string]*string) *CreateAppLlmApiKeyForPartnerResponse {
	s.Headers = v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerResponse) SetStatusCode(v int32) *CreateAppLlmApiKeyForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerResponse) SetBody(v *CreateAppLlmApiKeyForPartnerResponseBody) *CreateAppLlmApiKeyForPartnerResponse {
	s.Body = v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
