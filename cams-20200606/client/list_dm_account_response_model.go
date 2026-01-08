// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDmAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDmAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDmAccountResponse
	GetStatusCode() *int32
	SetBody(v *ListDmAccountResponseBody) *ListDmAccountResponse
	GetBody() *ListDmAccountResponseBody
}

type ListDmAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDmAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDmAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDmAccountResponse) GoString() string {
	return s.String()
}

func (s *ListDmAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDmAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDmAccountResponse) GetBody() *ListDmAccountResponseBody {
	return s.Body
}

func (s *ListDmAccountResponse) SetHeaders(v map[string]*string) *ListDmAccountResponse {
	s.Headers = v
	return s
}

func (s *ListDmAccountResponse) SetStatusCode(v int32) *ListDmAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDmAccountResponse) SetBody(v *ListDmAccountResponseBody) *ListDmAccountResponse {
	s.Body = v
	return s
}

func (s *ListDmAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
