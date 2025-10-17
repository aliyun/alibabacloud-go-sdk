// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobRunsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobRunsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobRunsResponse
	GetStatusCode() *int32
	SetBody(v *ListJobRunsResponseBody) *ListJobRunsResponse
	GetBody() *ListJobRunsResponseBody
}

type ListJobRunsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobRunsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobRunsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobRunsResponse) GoString() string {
	return s.String()
}

func (s *ListJobRunsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobRunsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobRunsResponse) GetBody() *ListJobRunsResponseBody {
	return s.Body
}

func (s *ListJobRunsResponse) SetHeaders(v map[string]*string) *ListJobRunsResponse {
	s.Headers = v
	return s
}

func (s *ListJobRunsResponse) SetStatusCode(v int32) *ListJobRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobRunsResponse) SetBody(v *ListJobRunsResponseBody) *ListJobRunsResponse {
	s.Body = v
	return s
}

func (s *ListJobRunsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
