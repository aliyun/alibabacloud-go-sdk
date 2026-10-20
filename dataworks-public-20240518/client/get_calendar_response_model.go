// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCalendarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCalendarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCalendarResponse
	GetStatusCode() *int32
	SetBody(v *GetCalendarResponseBody) *GetCalendarResponse
	GetBody() *GetCalendarResponseBody
}

type GetCalendarResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCalendarResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCalendarResponse) GoString() string {
	return s.String()
}

func (s *GetCalendarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCalendarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCalendarResponse) GetBody() *GetCalendarResponseBody {
	return s.Body
}

func (s *GetCalendarResponse) SetHeaders(v map[string]*string) *GetCalendarResponse {
	s.Headers = v
	return s
}

func (s *GetCalendarResponse) SetStatusCode(v int32) *GetCalendarResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCalendarResponse) SetBody(v *GetCalendarResponseBody) *GetCalendarResponse {
	s.Body = v
	return s
}

func (s *GetCalendarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
