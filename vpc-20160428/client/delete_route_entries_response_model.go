// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRouteEntriesResponseBody) *DeleteRouteEntriesResponse
	GetBody() *DeleteRouteEntriesResponseBody
}

type DeleteRouteEntriesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DeleteRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRouteEntriesResponse) GetBody() *DeleteRouteEntriesResponseBody {
	return s.Body
}

func (s *DeleteRouteEntriesResponse) SetHeaders(v map[string]*string) *DeleteRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DeleteRouteEntriesResponse) SetStatusCode(v int32) *DeleteRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRouteEntriesResponse) SetBody(v *DeleteRouteEntriesResponseBody) *DeleteRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *DeleteRouteEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
