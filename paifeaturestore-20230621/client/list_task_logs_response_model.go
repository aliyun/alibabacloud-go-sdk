// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTaskLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTaskLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListTaskLogsResponseBody) *ListTaskLogsResponse
	GetBody() *ListTaskLogsResponseBody
}

type ListTaskLogsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTaskLogsResponse) GoString() string {
	return s.String()
}

func (s *ListTaskLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTaskLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTaskLogsResponse) GetBody() *ListTaskLogsResponseBody {
	return s.Body
}

func (s *ListTaskLogsResponse) SetHeaders(v map[string]*string) *ListTaskLogsResponse {
	s.Headers = v
	return s
}

func (s *ListTaskLogsResponse) SetStatusCode(v int32) *ListTaskLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskLogsResponse) SetBody(v *ListTaskLogsResponseBody) *ListTaskLogsResponse {
	s.Body = v
	return s
}

func (s *ListTaskLogsResponse) Validate() error {
	return dara.Validate(s)
}
