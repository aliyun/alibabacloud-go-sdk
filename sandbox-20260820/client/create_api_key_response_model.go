// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateApiKeyResponseBody) *CreateApiKeyResponse
	GetBody() *CreateApiKeyResponseBody
}

type CreateApiKeyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApiKeyResponse) GetBody() *CreateApiKeyResponseBody {
	return s.Body
}

func (s *CreateApiKeyResponse) SetHeaders(v map[string]*string) *CreateApiKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateApiKeyResponse) SetStatusCode(v int32) *CreateApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApiKeyResponse) SetBody(v *CreateApiKeyResponseBody) *CreateApiKeyResponse {
	s.Body = v
	return s
}

func (s *CreateApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
