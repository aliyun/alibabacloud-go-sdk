// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPasskeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPasskeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPasskeysResponse
	GetStatusCode() *int32
	SetBody(v *ListPasskeysResponseBody) *ListPasskeysResponse
	GetBody() *ListPasskeysResponseBody
}

type ListPasskeysResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPasskeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPasskeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPasskeysResponse) GoString() string {
	return s.String()
}

func (s *ListPasskeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPasskeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPasskeysResponse) GetBody() *ListPasskeysResponseBody {
	return s.Body
}

func (s *ListPasskeysResponse) SetHeaders(v map[string]*string) *ListPasskeysResponse {
	s.Headers = v
	return s
}

func (s *ListPasskeysResponse) SetStatusCode(v int32) *ListPasskeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPasskeysResponse) SetBody(v *ListPasskeysResponseBody) *ListPasskeysResponse {
	s.Body = v
	return s
}

func (s *ListPasskeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
