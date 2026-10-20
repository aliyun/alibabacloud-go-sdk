// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalendarVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCalendarVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCalendarVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListCalendarVersionsResponseBody) *ListCalendarVersionsResponse
	GetBody() *ListCalendarVersionsResponseBody
}

type ListCalendarVersionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCalendarVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCalendarVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListCalendarVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCalendarVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCalendarVersionsResponse) GetBody() *ListCalendarVersionsResponseBody {
	return s.Body
}

func (s *ListCalendarVersionsResponse) SetHeaders(v map[string]*string) *ListCalendarVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListCalendarVersionsResponse) SetStatusCode(v int32) *ListCalendarVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCalendarVersionsResponse) SetBody(v *ListCalendarVersionsResponseBody) *ListCalendarVersionsResponse {
	s.Body = v
	return s
}

func (s *ListCalendarVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
