// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIncidentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIncidentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIncidentsResponse
	GetStatusCode() *int32
	SetBody(v *ListIncidentsResponseBody) *ListIncidentsResponse
	GetBody() *ListIncidentsResponseBody
}

type ListIncidentsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIncidentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIncidentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIncidentsResponse) GoString() string {
	return s.String()
}

func (s *ListIncidentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIncidentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIncidentsResponse) GetBody() *ListIncidentsResponseBody {
	return s.Body
}

func (s *ListIncidentsResponse) SetHeaders(v map[string]*string) *ListIncidentsResponse {
	s.Headers = v
	return s
}

func (s *ListIncidentsResponse) SetStatusCode(v int32) *ListIncidentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIncidentsResponse) SetBody(v *ListIncidentsResponseBody) *ListIncidentsResponse {
	s.Body = v
	return s
}

func (s *ListIncidentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
