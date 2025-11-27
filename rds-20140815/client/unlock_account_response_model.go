// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnlockAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnlockAccountResponse
	GetStatusCode() *int32
	SetBody(v *UnlockAccountResponseBody) *UnlockAccountResponse
	GetBody() *UnlockAccountResponseBody
}

type UnlockAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnlockAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnlockAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s UnlockAccountResponse) GoString() string {
	return s.String()
}

func (s *UnlockAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnlockAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnlockAccountResponse) GetBody() *UnlockAccountResponseBody {
	return s.Body
}

func (s *UnlockAccountResponse) SetHeaders(v map[string]*string) *UnlockAccountResponse {
	s.Headers = v
	return s
}

func (s *UnlockAccountResponse) SetStatusCode(v int32) *UnlockAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockAccountResponse) SetBody(v *UnlockAccountResponseBody) *UnlockAccountResponse {
	s.Body = v
	return s
}

func (s *UnlockAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
