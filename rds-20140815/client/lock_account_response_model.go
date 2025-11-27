// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LockAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LockAccountResponse
	GetStatusCode() *int32
	SetBody(v *LockAccountResponseBody) *LockAccountResponse
	GetBody() *LockAccountResponseBody
}

type LockAccountResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LockAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LockAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s LockAccountResponse) GoString() string {
	return s.String()
}

func (s *LockAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LockAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LockAccountResponse) GetBody() *LockAccountResponseBody {
	return s.Body
}

func (s *LockAccountResponse) SetHeaders(v map[string]*string) *LockAccountResponse {
	s.Headers = v
	return s
}

func (s *LockAccountResponse) SetStatusCode(v int32) *LockAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *LockAccountResponse) SetBody(v *LockAccountResponseBody) *LockAccountResponse {
	s.Body = v
	return s
}

func (s *LockAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
