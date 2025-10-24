// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomPersonsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomPersonsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomPersonsResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomPersonsResponseBody) *ListCustomPersonsResponse
	GetBody() *ListCustomPersonsResponseBody
}

type ListCustomPersonsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomPersonsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomPersonsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomPersonsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomPersonsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomPersonsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomPersonsResponse) GetBody() *ListCustomPersonsResponseBody {
	return s.Body
}

func (s *ListCustomPersonsResponse) SetHeaders(v map[string]*string) *ListCustomPersonsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomPersonsResponse) SetStatusCode(v int32) *ListCustomPersonsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomPersonsResponse) SetBody(v *ListCustomPersonsResponseBody) *ListCustomPersonsResponse {
	s.Body = v
	return s
}

func (s *ListCustomPersonsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
