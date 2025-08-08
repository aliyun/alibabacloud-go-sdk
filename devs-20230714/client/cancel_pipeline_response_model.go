// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelPipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelPipelineResponse
	GetStatusCode() *int32
	SetBody(v *Pipeline) *CancelPipelineResponse
	GetBody() *Pipeline
}

type CancelPipelineResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Pipeline          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelPipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelPipelineResponse) GoString() string {
	return s.String()
}

func (s *CancelPipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelPipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelPipelineResponse) GetBody() *Pipeline {
	return s.Body
}

func (s *CancelPipelineResponse) SetHeaders(v map[string]*string) *CancelPipelineResponse {
	s.Headers = v
	return s
}

func (s *CancelPipelineResponse) SetStatusCode(v int32) *CancelPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelPipelineResponse) SetBody(v *Pipeline) *CancelPipelineResponse {
	s.Body = v
	return s
}

func (s *CancelPipelineResponse) Validate() error {
	return dara.Validate(s)
}
