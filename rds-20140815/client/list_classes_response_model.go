// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClassesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClassesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClassesResponse
	GetStatusCode() *int32
	SetBody(v *ListClassesResponseBody) *ListClassesResponse
	GetBody() *ListClassesResponseBody
}

type ListClassesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClassesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClassesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClassesResponse) GoString() string {
	return s.String()
}

func (s *ListClassesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClassesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClassesResponse) GetBody() *ListClassesResponseBody {
	return s.Body
}

func (s *ListClassesResponse) SetHeaders(v map[string]*string) *ListClassesResponse {
	s.Headers = v
	return s
}

func (s *ListClassesResponse) SetStatusCode(v int32) *ListClassesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClassesResponse) SetBody(v *ListClassesResponseBody) *ListClassesResponse {
	s.Body = v
	return s
}

func (s *ListClassesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
