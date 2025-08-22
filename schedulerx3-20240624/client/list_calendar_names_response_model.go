// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalendarNamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCalendarNamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCalendarNamesResponse
	GetStatusCode() *int32
	SetBody(v *ListCalendarNamesResponseBody) *ListCalendarNamesResponse
	GetBody() *ListCalendarNamesResponseBody
}

type ListCalendarNamesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCalendarNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCalendarNamesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarNamesResponse) GoString() string {
	return s.String()
}

func (s *ListCalendarNamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCalendarNamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCalendarNamesResponse) GetBody() *ListCalendarNamesResponseBody {
	return s.Body
}

func (s *ListCalendarNamesResponse) SetHeaders(v map[string]*string) *ListCalendarNamesResponse {
	s.Headers = v
	return s
}

func (s *ListCalendarNamesResponse) SetStatusCode(v int32) *ListCalendarNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCalendarNamesResponse) SetBody(v *ListCalendarNamesResponseBody) *ListCalendarNamesResponse {
	s.Body = v
	return s
}

func (s *ListCalendarNamesResponse) Validate() error {
	return dara.Validate(s)
}
