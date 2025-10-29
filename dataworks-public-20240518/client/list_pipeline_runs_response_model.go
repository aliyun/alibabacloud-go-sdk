// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineRunsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPipelineRunsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPipelineRunsResponse
	GetStatusCode() *int32
	SetBody(v *ListPipelineRunsResponseBody) *ListPipelineRunsResponse
	GetBody() *ListPipelineRunsResponseBody
}

type ListPipelineRunsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPipelineRunsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPipelineRunsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRunsResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPipelineRunsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPipelineRunsResponse) GetBody() *ListPipelineRunsResponseBody {
	return s.Body
}

func (s *ListPipelineRunsResponse) SetHeaders(v map[string]*string) *ListPipelineRunsResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineRunsResponse) SetStatusCode(v int32) *ListPipelineRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineRunsResponse) SetBody(v *ListPipelineRunsResponseBody) *ListPipelineRunsResponse {
	s.Body = v
	return s
}

func (s *ListPipelineRunsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
