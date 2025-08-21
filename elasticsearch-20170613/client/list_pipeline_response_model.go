// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPipelineResponse
	GetStatusCode() *int32
	SetBody(v *ListPipelineResponseBody) *ListPipelineResponse
	GetBody() *ListPipelineResponseBody
}

type ListPipelineResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPipelineResponse) GetBody() *ListPipelineResponseBody {
	return s.Body
}

func (s *ListPipelineResponse) SetHeaders(v map[string]*string) *ListPipelineResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineResponse) SetStatusCode(v int32) *ListPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineResponse) SetBody(v *ListPipelineResponseBody) *ListPipelineResponse {
	s.Body = v
	return s
}

func (s *ListPipelineResponse) Validate() error {
	return dara.Validate(s)
}
