// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSystemLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSystemLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListSystemLogsResponseBody) *ListSystemLogsResponse
	GetBody() *ListSystemLogsResponseBody
}

type ListSystemLogsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSystemLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSystemLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSystemLogsResponse) GoString() string {
	return s.String()
}

func (s *ListSystemLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSystemLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSystemLogsResponse) GetBody() *ListSystemLogsResponseBody {
	return s.Body
}

func (s *ListSystemLogsResponse) SetHeaders(v map[string]*string) *ListSystemLogsResponse {
	s.Headers = v
	return s
}

func (s *ListSystemLogsResponse) SetStatusCode(v int32) *ListSystemLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSystemLogsResponse) SetBody(v *ListSystemLogsResponseBody) *ListSystemLogsResponse {
	s.Body = v
	return s
}

func (s *ListSystemLogsResponse) Validate() error {
	return dara.Validate(s)
}
