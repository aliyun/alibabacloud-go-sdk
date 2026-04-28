// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRevisionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRevisionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRevisionResponse
	GetStatusCode() *int32
	SetBody(v *ListRevisionResponseBody) *ListRevisionResponse
	GetBody() *ListRevisionResponseBody
}

type ListRevisionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRevisionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRevisionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRevisionResponse) GoString() string {
	return s.String()
}

func (s *ListRevisionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRevisionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRevisionResponse) GetBody() *ListRevisionResponseBody {
	return s.Body
}

func (s *ListRevisionResponse) SetHeaders(v map[string]*string) *ListRevisionResponse {
	s.Headers = v
	return s
}

func (s *ListRevisionResponse) SetStatusCode(v int32) *ListRevisionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRevisionResponse) SetBody(v *ListRevisionResponseBody) *ListRevisionResponse {
	s.Body = v
	return s
}

func (s *ListRevisionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
