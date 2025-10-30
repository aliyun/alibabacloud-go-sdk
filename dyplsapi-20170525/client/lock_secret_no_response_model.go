// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockSecretNoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LockSecretNoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LockSecretNoResponse
	GetStatusCode() *int32
	SetBody(v *LockSecretNoResponseBody) *LockSecretNoResponse
	GetBody() *LockSecretNoResponseBody
}

type LockSecretNoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LockSecretNoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LockSecretNoResponse) String() string {
	return dara.Prettify(s)
}

func (s LockSecretNoResponse) GoString() string {
	return s.String()
}

func (s *LockSecretNoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LockSecretNoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LockSecretNoResponse) GetBody() *LockSecretNoResponseBody {
	return s.Body
}

func (s *LockSecretNoResponse) SetHeaders(v map[string]*string) *LockSecretNoResponse {
	s.Headers = v
	return s
}

func (s *LockSecretNoResponse) SetStatusCode(v int32) *LockSecretNoResponse {
	s.StatusCode = &v
	return s
}

func (s *LockSecretNoResponse) SetBody(v *LockSecretNoResponseBody) *LockSecretNoResponse {
	s.Body = v
	return s
}

func (s *LockSecretNoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
