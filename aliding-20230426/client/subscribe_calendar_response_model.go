// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeCalendarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubscribeCalendarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubscribeCalendarResponse
	GetStatusCode() *int32
	SetBody(v *SubscribeCalendarResponseBody) *SubscribeCalendarResponse
	GetBody() *SubscribeCalendarResponseBody
}

type SubscribeCalendarResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubscribeCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubscribeCalendarResponse) String() string {
	return dara.Prettify(s)
}

func (s SubscribeCalendarResponse) GoString() string {
	return s.String()
}

func (s *SubscribeCalendarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubscribeCalendarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubscribeCalendarResponse) GetBody() *SubscribeCalendarResponseBody {
	return s.Body
}

func (s *SubscribeCalendarResponse) SetHeaders(v map[string]*string) *SubscribeCalendarResponse {
	s.Headers = v
	return s
}

func (s *SubscribeCalendarResponse) SetStatusCode(v int32) *SubscribeCalendarResponse {
	s.StatusCode = &v
	return s
}

func (s *SubscribeCalendarResponse) SetBody(v *SubscribeCalendarResponseBody) *SubscribeCalendarResponse {
	s.Body = v
	return s
}

func (s *SubscribeCalendarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
