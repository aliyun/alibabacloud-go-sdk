// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHoneypotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHoneypotResponse
	GetStatusCode() *int32
	SetBody(v *ListHoneypotResponseBody) *ListHoneypotResponse
	GetBody() *ListHoneypotResponseBody
}

type ListHoneypotResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHoneypotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHoneypotResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotResponse) GoString() string {
	return s.String()
}

func (s *ListHoneypotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHoneypotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHoneypotResponse) GetBody() *ListHoneypotResponseBody {
	return s.Body
}

func (s *ListHoneypotResponse) SetHeaders(v map[string]*string) *ListHoneypotResponse {
	s.Headers = v
	return s
}

func (s *ListHoneypotResponse) SetStatusCode(v int32) *ListHoneypotResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHoneypotResponse) SetBody(v *ListHoneypotResponseBody) *ListHoneypotResponse {
	s.Body = v
	return s
}

func (s *ListHoneypotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
