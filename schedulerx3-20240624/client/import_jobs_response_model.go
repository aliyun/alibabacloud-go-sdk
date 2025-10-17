// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportJobsResponse
	GetStatusCode() *int32
	SetBody(v *ImportJobsResponseBody) *ImportJobsResponse
	GetBody() *ImportJobsResponseBody
}

type ImportJobsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportJobsResponse) GoString() string {
	return s.String()
}

func (s *ImportJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportJobsResponse) GetBody() *ImportJobsResponseBody {
	return s.Body
}

func (s *ImportJobsResponse) SetHeaders(v map[string]*string) *ImportJobsResponse {
	s.Headers = v
	return s
}

func (s *ImportJobsResponse) SetStatusCode(v int32) *ImportJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportJobsResponse) SetBody(v *ImportJobsResponseBody) *ImportJobsResponse {
	s.Body = v
	return s
}

func (s *ImportJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
