// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutionLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExecutionLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExecutionLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListExecutionLogsResponseBody) *ListExecutionLogsResponse
	GetBody() *ListExecutionLogsResponseBody
}

type ListExecutionLogsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExecutionLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExecutionLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionLogsResponse) GoString() string {
	return s.String()
}

func (s *ListExecutionLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExecutionLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExecutionLogsResponse) GetBody() *ListExecutionLogsResponseBody {
	return s.Body
}

func (s *ListExecutionLogsResponse) SetHeaders(v map[string]*string) *ListExecutionLogsResponse {
	s.Headers = v
	return s
}

func (s *ListExecutionLogsResponse) SetStatusCode(v int32) *ListExecutionLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExecutionLogsResponse) SetBody(v *ListExecutionLogsResponseBody) *ListExecutionLogsResponse {
	s.Body = v
	return s
}

func (s *ListExecutionLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
