// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockSecretNoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnlockSecretNoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnlockSecretNoResponse
	GetStatusCode() *int32
	SetBody(v *UnlockSecretNoResponseBody) *UnlockSecretNoResponse
	GetBody() *UnlockSecretNoResponseBody
}

type UnlockSecretNoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnlockSecretNoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnlockSecretNoResponse) String() string {
	return dara.Prettify(s)
}

func (s UnlockSecretNoResponse) GoString() string {
	return s.String()
}

func (s *UnlockSecretNoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnlockSecretNoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnlockSecretNoResponse) GetBody() *UnlockSecretNoResponseBody {
	return s.Body
}

func (s *UnlockSecretNoResponse) SetHeaders(v map[string]*string) *UnlockSecretNoResponse {
	s.Headers = v
	return s
}

func (s *UnlockSecretNoResponse) SetStatusCode(v int32) *UnlockSecretNoResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockSecretNoResponse) SetBody(v *UnlockSecretNoResponseBody) *UnlockSecretNoResponse {
	s.Body = v
	return s
}

func (s *UnlockSecretNoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
