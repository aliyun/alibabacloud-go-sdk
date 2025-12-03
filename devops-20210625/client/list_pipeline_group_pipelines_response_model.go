// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineGroupPipelinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPipelineGroupPipelinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPipelineGroupPipelinesResponse
	GetStatusCode() *int32
	SetBody(v *ListPipelineGroupPipelinesResponseBody) *ListPipelineGroupPipelinesResponse
	GetBody() *ListPipelineGroupPipelinesResponseBody
}

type ListPipelineGroupPipelinesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPipelineGroupPipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPipelineGroupPipelinesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineGroupPipelinesResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineGroupPipelinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPipelineGroupPipelinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPipelineGroupPipelinesResponse) GetBody() *ListPipelineGroupPipelinesResponseBody {
	return s.Body
}

func (s *ListPipelineGroupPipelinesResponse) SetHeaders(v map[string]*string) *ListPipelineGroupPipelinesResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineGroupPipelinesResponse) SetStatusCode(v int32) *ListPipelineGroupPipelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponse) SetBody(v *ListPipelineGroupPipelinesResponseBody) *ListPipelineGroupPipelinesResponse {
	s.Body = v
	return s
}

func (s *ListPipelineGroupPipelinesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
