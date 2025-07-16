// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnsubscribeCalendarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnsubscribeCalendarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnsubscribeCalendarResponse
	GetStatusCode() *int32
	SetBody(v *UnsubscribeCalendarResponseBody) *UnsubscribeCalendarResponse
	GetBody() *UnsubscribeCalendarResponseBody
}

type UnsubscribeCalendarResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnsubscribeCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnsubscribeCalendarResponse) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeCalendarResponse) GoString() string {
	return s.String()
}

func (s *UnsubscribeCalendarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnsubscribeCalendarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnsubscribeCalendarResponse) GetBody() *UnsubscribeCalendarResponseBody {
	return s.Body
}

func (s *UnsubscribeCalendarResponse) SetHeaders(v map[string]*string) *UnsubscribeCalendarResponse {
	s.Headers = v
	return s
}

func (s *UnsubscribeCalendarResponse) SetStatusCode(v int32) *UnsubscribeCalendarResponse {
	s.StatusCode = &v
	return s
}

func (s *UnsubscribeCalendarResponse) SetBody(v *UnsubscribeCalendarResponseBody) *UnsubscribeCalendarResponse {
	s.Body = v
	return s
}

func (s *UnsubscribeCalendarResponse) Validate() error {
	return dara.Validate(s)
}
