// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDmAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindDmAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindDmAccountResponse
	GetStatusCode() *int32
	SetBody(v *UnbindDmAccountResponseBody) *UnbindDmAccountResponse
	GetBody() *UnbindDmAccountResponseBody
}

type UnbindDmAccountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindDmAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindDmAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindDmAccountResponse) GoString() string {
	return s.String()
}

func (s *UnbindDmAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindDmAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindDmAccountResponse) GetBody() *UnbindDmAccountResponseBody {
	return s.Body
}

func (s *UnbindDmAccountResponse) SetHeaders(v map[string]*string) *UnbindDmAccountResponse {
	s.Headers = v
	return s
}

func (s *UnbindDmAccountResponse) SetStatusCode(v int32) *UnbindDmAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindDmAccountResponse) SetBody(v *UnbindDmAccountResponseBody) *UnbindDmAccountResponse {
	s.Body = v
	return s
}

func (s *UnbindDmAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
