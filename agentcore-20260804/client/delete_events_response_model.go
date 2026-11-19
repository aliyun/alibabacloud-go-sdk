// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEventsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEventsResponseBody) *DeleteEventsResponse
	GetBody() *DeleteEventsResponseBody
}

type DeleteEventsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventsResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEventsResponse) GetBody() *DeleteEventsResponseBody {
	return s.Body
}

func (s *DeleteEventsResponse) SetHeaders(v map[string]*string) *DeleteEventsResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventsResponse) SetStatusCode(v int32) *DeleteEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventsResponse) SetBody(v *DeleteEventsResponseBody) *DeleteEventsResponse {
	s.Body = v
	return s
}

func (s *DeleteEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
