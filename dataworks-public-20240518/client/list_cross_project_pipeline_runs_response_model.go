// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrossProjectPipelineRunsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCrossProjectPipelineRunsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCrossProjectPipelineRunsResponse
	GetStatusCode() *int32
	SetBody(v *ListCrossProjectPipelineRunsResponseBody) *ListCrossProjectPipelineRunsResponse
	GetBody() *ListCrossProjectPipelineRunsResponseBody
}

type ListCrossProjectPipelineRunsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCrossProjectPipelineRunsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCrossProjectPipelineRunsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectPipelineRunsResponse) GoString() string {
	return s.String()
}

func (s *ListCrossProjectPipelineRunsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCrossProjectPipelineRunsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCrossProjectPipelineRunsResponse) GetBody() *ListCrossProjectPipelineRunsResponseBody {
	return s.Body
}

func (s *ListCrossProjectPipelineRunsResponse) SetHeaders(v map[string]*string) *ListCrossProjectPipelineRunsResponse {
	s.Headers = v
	return s
}

func (s *ListCrossProjectPipelineRunsResponse) SetStatusCode(v int32) *ListCrossProjectPipelineRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponse) SetBody(v *ListCrossProjectPipelineRunsResponseBody) *ListCrossProjectPipelineRunsResponse {
	s.Body = v
	return s
}

func (s *ListCrossProjectPipelineRunsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
