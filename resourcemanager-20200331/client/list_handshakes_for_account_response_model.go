// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHandshakesForAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHandshakesForAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHandshakesForAccountResponse
	GetStatusCode() *int32
	SetBody(v *ListHandshakesForAccountResponseBody) *ListHandshakesForAccountResponse
	GetBody() *ListHandshakesForAccountResponseBody
}

type ListHandshakesForAccountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHandshakesForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHandshakesForAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHandshakesForAccountResponse) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHandshakesForAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHandshakesForAccountResponse) GetBody() *ListHandshakesForAccountResponseBody {
	return s.Body
}

func (s *ListHandshakesForAccountResponse) SetHeaders(v map[string]*string) *ListHandshakesForAccountResponse {
	s.Headers = v
	return s
}

func (s *ListHandshakesForAccountResponse) SetStatusCode(v int32) *ListHandshakesForAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHandshakesForAccountResponse) SetBody(v *ListHandshakesForAccountResponseBody) *ListHandshakesForAccountResponse {
	s.Body = v
	return s
}

func (s *ListHandshakesForAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
