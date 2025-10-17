// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLablesResponse
	GetStatusCode() *int32
	SetBody(v *ListLablesResponseBody) *ListLablesResponse
	GetBody() *ListLablesResponseBody
}

type ListLablesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLablesResponse) GoString() string {
	return s.String()
}

func (s *ListLablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLablesResponse) GetBody() *ListLablesResponseBody {
	return s.Body
}

func (s *ListLablesResponse) SetHeaders(v map[string]*string) *ListLablesResponse {
	s.Headers = v
	return s
}

func (s *ListLablesResponse) SetStatusCode(v int32) *ListLablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLablesResponse) SetBody(v *ListLablesResponseBody) *ListLablesResponse {
	s.Body = v
	return s
}

func (s *ListLablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
