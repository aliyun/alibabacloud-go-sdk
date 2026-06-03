// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmsSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSmsSignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSmsSignResponse
	GetStatusCode() *int32
	SetBody(v *ListSmsSignResponseBody) *ListSmsSignResponse
	GetBody() *ListSmsSignResponseBody
}

type ListSmsSignResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSmsSignResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSmsSignResponse) GoString() string {
	return s.String()
}

func (s *ListSmsSignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSmsSignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSmsSignResponse) GetBody() *ListSmsSignResponseBody {
	return s.Body
}

func (s *ListSmsSignResponse) SetHeaders(v map[string]*string) *ListSmsSignResponse {
	s.Headers = v
	return s
}

func (s *ListSmsSignResponse) SetStatusCode(v int32) *ListSmsSignResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSmsSignResponse) SetBody(v *ListSmsSignResponseBody) *ListSmsSignResponse {
	s.Body = v
	return s
}

func (s *ListSmsSignResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
