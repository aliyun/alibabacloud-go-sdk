// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListLogsResponseBody) *ListLogsResponse
	GetBody() *ListLogsResponseBody
}

type ListLogsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLogsResponse) GoString() string {
	return s.String()
}

func (s *ListLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLogsResponse) GetBody() *ListLogsResponseBody {
	return s.Body
}

func (s *ListLogsResponse) SetHeaders(v map[string]*string) *ListLogsResponse {
	s.Headers = v
	return s
}

func (s *ListLogsResponse) SetStatusCode(v int32) *ListLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogsResponse) SetBody(v *ListLogsResponseBody) *ListLogsResponse {
	s.Body = v
	return s
}

func (s *ListLogsResponse) Validate() error {
	return dara.Validate(s)
}
