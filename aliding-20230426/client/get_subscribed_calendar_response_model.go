// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscribedCalendarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSubscribedCalendarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSubscribedCalendarResponse
	GetStatusCode() *int32
	SetBody(v *GetSubscribedCalendarResponseBody) *GetSubscribedCalendarResponse
	GetBody() *GetSubscribedCalendarResponseBody
}

type GetSubscribedCalendarResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubscribedCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubscribedCalendarResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSubscribedCalendarResponse) GoString() string {
	return s.String()
}

func (s *GetSubscribedCalendarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSubscribedCalendarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSubscribedCalendarResponse) GetBody() *GetSubscribedCalendarResponseBody {
	return s.Body
}

func (s *GetSubscribedCalendarResponse) SetHeaders(v map[string]*string) *GetSubscribedCalendarResponse {
	s.Headers = v
	return s
}

func (s *GetSubscribedCalendarResponse) SetStatusCode(v int32) *GetSubscribedCalendarResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubscribedCalendarResponse) SetBody(v *GetSubscribedCalendarResponseBody) *GetSubscribedCalendarResponse {
	s.Body = v
	return s
}

func (s *GetSubscribedCalendarResponse) Validate() error {
	return dara.Validate(s)
}
