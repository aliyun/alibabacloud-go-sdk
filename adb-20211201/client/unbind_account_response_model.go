// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindAccountResponse
	GetStatusCode() *int32
	SetBody(v *UnbindAccountResponseBody) *UnbindAccountResponse
	GetBody() *UnbindAccountResponseBody
}

type UnbindAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindAccountResponse) GoString() string {
	return s.String()
}

func (s *UnbindAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindAccountResponse) GetBody() *UnbindAccountResponseBody {
	return s.Body
}

func (s *UnbindAccountResponse) SetHeaders(v map[string]*string) *UnbindAccountResponse {
	s.Headers = v
	return s
}

func (s *UnbindAccountResponse) SetStatusCode(v int32) *UnbindAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindAccountResponse) SetBody(v *UnbindAccountResponseBody) *UnbindAccountResponse {
	s.Body = v
	return s
}

func (s *UnbindAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
