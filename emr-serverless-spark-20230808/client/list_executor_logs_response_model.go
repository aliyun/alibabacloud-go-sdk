// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutorLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExecutorLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExecutorLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListExecutorLogsResponseBody) *ListExecutorLogsResponse
	GetBody() *ListExecutorLogsResponseBody
}

type ListExecutorLogsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExecutorLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExecutorLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorLogsResponse) GoString() string {
	return s.String()
}

func (s *ListExecutorLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExecutorLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExecutorLogsResponse) GetBody() *ListExecutorLogsResponseBody {
	return s.Body
}

func (s *ListExecutorLogsResponse) SetHeaders(v map[string]*string) *ListExecutorLogsResponse {
	s.Headers = v
	return s
}

func (s *ListExecutorLogsResponse) SetStatusCode(v int32) *ListExecutorLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExecutorLogsResponse) SetBody(v *ListExecutorLogsResponseBody) *ListExecutorLogsResponse {
	s.Body = v
	return s
}

func (s *ListExecutorLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
