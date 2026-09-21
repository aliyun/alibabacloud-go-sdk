// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAgenticApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetAgenticApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetAgenticApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *ResetAgenticApiKeyResponseBody) *ResetAgenticApiKeyResponse
	GetBody() *ResetAgenticApiKeyResponseBody
}

type ResetAgenticApiKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAgenticApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAgenticApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetAgenticApiKeyResponse) GoString() string {
	return s.String()
}

func (s *ResetAgenticApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetAgenticApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetAgenticApiKeyResponse) GetBody() *ResetAgenticApiKeyResponseBody {
	return s.Body
}

func (s *ResetAgenticApiKeyResponse) SetHeaders(v map[string]*string) *ResetAgenticApiKeyResponse {
	s.Headers = v
	return s
}

func (s *ResetAgenticApiKeyResponse) SetStatusCode(v int32) *ResetAgenticApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAgenticApiKeyResponse) SetBody(v *ResetAgenticApiKeyResponseBody) *ResetAgenticApiKeyResponse {
	s.Body = v
	return s
}

func (s *ResetAgenticApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
