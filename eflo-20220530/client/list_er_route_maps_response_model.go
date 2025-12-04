// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListErRouteMapsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListErRouteMapsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListErRouteMapsResponse
	GetStatusCode() *int32
	SetBody(v *ListErRouteMapsResponseBody) *ListErRouteMapsResponse
	GetBody() *ListErRouteMapsResponseBody
}

type ListErRouteMapsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListErRouteMapsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListErRouteMapsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListErRouteMapsResponse) GoString() string {
	return s.String()
}

func (s *ListErRouteMapsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListErRouteMapsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListErRouteMapsResponse) GetBody() *ListErRouteMapsResponseBody {
	return s.Body
}

func (s *ListErRouteMapsResponse) SetHeaders(v map[string]*string) *ListErRouteMapsResponse {
	s.Headers = v
	return s
}

func (s *ListErRouteMapsResponse) SetStatusCode(v int32) *ListErRouteMapsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListErRouteMapsResponse) SetBody(v *ListErRouteMapsResponseBody) *ListErRouteMapsResponse {
	s.Body = v
	return s
}

func (s *ListErRouteMapsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
