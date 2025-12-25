// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCalendarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCalendarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCalendarResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCalendarResponseBody) *DeleteCalendarResponse
	GetBody() *DeleteCalendarResponseBody
}

type DeleteCalendarResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCalendarResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCalendarResponse) GoString() string {
	return s.String()
}

func (s *DeleteCalendarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCalendarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCalendarResponse) GetBody() *DeleteCalendarResponseBody {
	return s.Body
}

func (s *DeleteCalendarResponse) SetHeaders(v map[string]*string) *DeleteCalendarResponse {
	s.Headers = v
	return s
}

func (s *DeleteCalendarResponse) SetStatusCode(v int32) *DeleteCalendarResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCalendarResponse) SetBody(v *DeleteCalendarResponseBody) *DeleteCalendarResponse {
	s.Body = v
	return s
}

func (s *DeleteCalendarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
