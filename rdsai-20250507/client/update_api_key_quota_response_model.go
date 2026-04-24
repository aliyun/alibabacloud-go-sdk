// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiKeyQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApiKeyQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApiKeyQuotaResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApiKeyQuotaResponseBody) *UpdateApiKeyQuotaResponse
	GetBody() *UpdateApiKeyQuotaResponseBody
}

type UpdateApiKeyQuotaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApiKeyQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApiKeyQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiKeyQuotaResponse) GoString() string {
	return s.String()
}

func (s *UpdateApiKeyQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApiKeyQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApiKeyQuotaResponse) GetBody() *UpdateApiKeyQuotaResponseBody {
	return s.Body
}

func (s *UpdateApiKeyQuotaResponse) SetHeaders(v map[string]*string) *UpdateApiKeyQuotaResponse {
	s.Headers = v
	return s
}

func (s *UpdateApiKeyQuotaResponse) SetStatusCode(v int32) *UpdateApiKeyQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApiKeyQuotaResponse) SetBody(v *UpdateApiKeyQuotaResponseBody) *UpdateApiKeyQuotaResponse {
	s.Body = v
	return s
}

func (s *UpdateApiKeyQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
