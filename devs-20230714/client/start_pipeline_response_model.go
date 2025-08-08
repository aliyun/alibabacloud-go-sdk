// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartPipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartPipelineResponse
	GetStatusCode() *int32
	SetBody(v *Pipeline) *StartPipelineResponse
	GetBody() *Pipeline
}

type StartPipelineResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Pipeline          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartPipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s StartPipelineResponse) GoString() string {
	return s.String()
}

func (s *StartPipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartPipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartPipelineResponse) GetBody() *Pipeline {
	return s.Body
}

func (s *StartPipelineResponse) SetHeaders(v map[string]*string) *StartPipelineResponse {
	s.Headers = v
	return s
}

func (s *StartPipelineResponse) SetStatusCode(v int32) *StartPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *StartPipelineResponse) SetBody(v *Pipeline) *StartPipelineResponse {
	s.Body = v
	return s
}

func (s *StartPipelineResponse) Validate() error {
	return dara.Validate(s)
}
