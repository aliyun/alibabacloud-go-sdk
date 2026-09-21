// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApiKeyResponseBody) *UpdateApiKeyResponse
	GetBody() *UpdateApiKeyResponseBody
}

type UpdateApiKeyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiKeyResponse) GoString() string {
	return s.String()
}

func (s *UpdateApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApiKeyResponse) GetBody() *UpdateApiKeyResponseBody {
	return s.Body
}

func (s *UpdateApiKeyResponse) SetHeaders(v map[string]*string) *UpdateApiKeyResponse {
	s.Headers = v
	return s
}

func (s *UpdateApiKeyResponse) SetStatusCode(v int32) *UpdateApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApiKeyResponse) SetBody(v *UpdateApiKeyResponseBody) *UpdateApiKeyResponse {
	s.Body = v
	return s
}

func (s *UpdateApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
