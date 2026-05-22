// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *DisableApiKeyResponseBody) *DisableApiKeyResponse
	GetBody() *DisableApiKeyResponseBody
}

type DisableApiKeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableApiKeyResponse) GoString() string {
	return s.String()
}

func (s *DisableApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableApiKeyResponse) GetBody() *DisableApiKeyResponseBody {
	return s.Body
}

func (s *DisableApiKeyResponse) SetHeaders(v map[string]*string) *DisableApiKeyResponse {
	s.Headers = v
	return s
}

func (s *DisableApiKeyResponse) SetStatusCode(v int32) *DisableApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApiKeyResponse) SetBody(v *DisableApiKeyResponseBody) *DisableApiKeyResponse {
	s.Body = v
	return s
}

func (s *DisableApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
