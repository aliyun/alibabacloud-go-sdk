// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAppSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetAppSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetAppSecretResponse
	GetStatusCode() *int32
	SetBody(v *ResetAppSecretResponseBody) *ResetAppSecretResponse
	GetBody() *ResetAppSecretResponseBody
}

type ResetAppSecretResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAppSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAppSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetAppSecretResponse) GoString() string {
	return s.String()
}

func (s *ResetAppSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetAppSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetAppSecretResponse) GetBody() *ResetAppSecretResponseBody {
	return s.Body
}

func (s *ResetAppSecretResponse) SetHeaders(v map[string]*string) *ResetAppSecretResponse {
	s.Headers = v
	return s
}

func (s *ResetAppSecretResponse) SetStatusCode(v int32) *ResetAppSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAppSecretResponse) SetBody(v *ResetAppSecretResponseBody) *ResetAppSecretResponse {
	s.Body = v
	return s
}

func (s *ResetAppSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
