// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVccRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVccRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVccRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *ListVccRouteEntriesResponseBody) *ListVccRouteEntriesResponse
	GetBody() *ListVccRouteEntriesResponseBody
}

type ListVccRouteEntriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVccRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVccRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVccRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListVccRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVccRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVccRouteEntriesResponse) GetBody() *ListVccRouteEntriesResponseBody {
	return s.Body
}

func (s *ListVccRouteEntriesResponse) SetHeaders(v map[string]*string) *ListVccRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListVccRouteEntriesResponse) SetStatusCode(v int32) *ListVccRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVccRouteEntriesResponse) SetBody(v *ListVccRouteEntriesResponseBody) *ListVccRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *ListVccRouteEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
