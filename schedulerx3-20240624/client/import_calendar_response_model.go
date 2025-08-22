// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportCalendarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportCalendarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportCalendarResponse
	GetStatusCode() *int32
	SetBody(v *ImportCalendarResponseBody) *ImportCalendarResponse
	GetBody() *ImportCalendarResponseBody
}

type ImportCalendarResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportCalendarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportCalendarResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportCalendarResponse) GoString() string {
	return s.String()
}

func (s *ImportCalendarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportCalendarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportCalendarResponse) GetBody() *ImportCalendarResponseBody {
	return s.Body
}

func (s *ImportCalendarResponse) SetHeaders(v map[string]*string) *ImportCalendarResponse {
	s.Headers = v
	return s
}

func (s *ImportCalendarResponse) SetStatusCode(v int32) *ImportCalendarResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportCalendarResponse) SetBody(v *ImportCalendarResponseBody) *ImportCalendarResponse {
	s.Body = v
	return s
}

func (s *ImportCalendarResponse) Validate() error {
	return dara.Validate(s)
}
