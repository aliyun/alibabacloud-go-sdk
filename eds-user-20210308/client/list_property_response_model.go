// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPropertyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPropertyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPropertyResponse
	GetStatusCode() *int32
	SetBody(v *ListPropertyResponseBody) *ListPropertyResponse
	GetBody() *ListPropertyResponseBody
}

type ListPropertyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPropertyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPropertyResponse) GoString() string {
	return s.String()
}

func (s *ListPropertyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPropertyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPropertyResponse) GetBody() *ListPropertyResponseBody {
	return s.Body
}

func (s *ListPropertyResponse) SetHeaders(v map[string]*string) *ListPropertyResponse {
	s.Headers = v
	return s
}

func (s *ListPropertyResponse) SetStatusCode(v int32) *ListPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPropertyResponse) SetBody(v *ListPropertyResponseBody) *ListPropertyResponse {
	s.Body = v
	return s
}

func (s *ListPropertyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
