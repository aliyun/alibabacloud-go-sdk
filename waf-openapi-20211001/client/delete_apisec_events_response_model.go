// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApisecEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApisecEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApisecEventsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApisecEventsResponseBody) *DeleteApisecEventsResponse
	GetBody() *DeleteApisecEventsResponseBody
}

type DeleteApisecEventsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApisecEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApisecEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApisecEventsResponse) GoString() string {
	return s.String()
}

func (s *DeleteApisecEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApisecEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApisecEventsResponse) GetBody() *DeleteApisecEventsResponseBody {
	return s.Body
}

func (s *DeleteApisecEventsResponse) SetHeaders(v map[string]*string) *DeleteApisecEventsResponse {
	s.Headers = v
	return s
}

func (s *DeleteApisecEventsResponse) SetStatusCode(v int32) *DeleteApisecEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApisecEventsResponse) SetBody(v *DeleteApisecEventsResponseBody) *DeleteApisecEventsResponse {
	s.Body = v
	return s
}

func (s *DeleteApisecEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
