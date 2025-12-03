// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPipelineJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPipelineJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListPipelineJobsResponseBody) *ListPipelineJobsResponse
	GetBody() *ListPipelineJobsResponseBody
}

type ListPipelineJobsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPipelineJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPipelineJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineJobsResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPipelineJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPipelineJobsResponse) GetBody() *ListPipelineJobsResponseBody {
	return s.Body
}

func (s *ListPipelineJobsResponse) SetHeaders(v map[string]*string) *ListPipelineJobsResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineJobsResponse) SetStatusCode(v int32) *ListPipelineJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineJobsResponse) SetBody(v *ListPipelineJobsResponseBody) *ListPipelineJobsResponse {
	s.Body = v
	return s
}

func (s *ListPipelineJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
