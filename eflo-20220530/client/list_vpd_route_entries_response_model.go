// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpdRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpdRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpdRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *ListVpdRouteEntriesResponseBody) *ListVpdRouteEntriesResponse
	GetBody() *ListVpdRouteEntriesResponseBody
}

type ListVpdRouteEntriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpdRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpdRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpdRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListVpdRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpdRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpdRouteEntriesResponse) GetBody() *ListVpdRouteEntriesResponseBody {
	return s.Body
}

func (s *ListVpdRouteEntriesResponse) SetHeaders(v map[string]*string) *ListVpdRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListVpdRouteEntriesResponse) SetStatusCode(v int32) *ListVpdRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpdRouteEntriesResponse) SetBody(v *ListVpdRouteEntriesResponseBody) *ListVpdRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *ListVpdRouteEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
