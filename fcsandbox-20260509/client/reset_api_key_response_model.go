// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *ResetApiKeyResponseBody) *ResetApiKeyResponse
	GetBody() *ResetApiKeyResponseBody
}

type ResetApiKeyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetApiKeyResponse) GoString() string {
	return s.String()
}

func (s *ResetApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetApiKeyResponse) GetBody() *ResetApiKeyResponseBody {
	return s.Body
}

func (s *ResetApiKeyResponse) SetHeaders(v map[string]*string) *ResetApiKeyResponse {
	s.Headers = v
	return s
}

func (s *ResetApiKeyResponse) SetStatusCode(v int32) *ResetApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetApiKeyResponse) SetBody(v *ResetApiKeyResponseBody) *ResetApiKeyResponse {
	s.Body = v
	return s
}

func (s *ResetApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
