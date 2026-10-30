// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRouteMapsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRouteMapsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRouteMapsResponse
	GetStatusCode() *int32
	SetBody(v *ListRouteMapsResponseBody) *ListRouteMapsResponse
	GetBody() *ListRouteMapsResponseBody
}

type ListRouteMapsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRouteMapsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRouteMapsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRouteMapsResponse) GoString() string {
	return s.String()
}

func (s *ListRouteMapsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRouteMapsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRouteMapsResponse) GetBody() *ListRouteMapsResponseBody {
	return s.Body
}

func (s *ListRouteMapsResponse) SetHeaders(v map[string]*string) *ListRouteMapsResponse {
	s.Headers = v
	return s
}

func (s *ListRouteMapsResponse) SetStatusCode(v int32) *ListRouteMapsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRouteMapsResponse) SetBody(v *ListRouteMapsResponseBody) *ListRouteMapsResponse {
	s.Body = v
	return s
}

func (s *ListRouteMapsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
