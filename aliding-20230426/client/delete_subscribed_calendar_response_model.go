// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubscribedCalendarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSubscribedCalendarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSubscribedCalendarResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSubscribedCalendarResponseBody) *DeleteSubscribedCalendarResponse
	GetBody() *DeleteSubscribedCalendarResponseBody
}

type DeleteSubscribedCalendarResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSubscribedCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSubscribedCalendarResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubscribedCalendarResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubscribedCalendarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSubscribedCalendarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSubscribedCalendarResponse) GetBody() *DeleteSubscribedCalendarResponseBody {
	return s.Body
}

func (s *DeleteSubscribedCalendarResponse) SetHeaders(v map[string]*string) *DeleteSubscribedCalendarResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubscribedCalendarResponse) SetStatusCode(v int32) *DeleteSubscribedCalendarResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSubscribedCalendarResponse) SetBody(v *DeleteSubscribedCalendarResponseBody) *DeleteSubscribedCalendarResponse {
	s.Body = v
	return s
}

func (s *DeleteSubscribedCalendarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
