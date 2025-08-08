// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutPipelineStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutPipelineStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutPipelineStatusResponse
	GetStatusCode() *int32
	SetBody(v *Pipeline) *PutPipelineStatusResponse
	GetBody() *Pipeline
}

type PutPipelineStatusResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Pipeline          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutPipelineStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s PutPipelineStatusResponse) GoString() string {
	return s.String()
}

func (s *PutPipelineStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutPipelineStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutPipelineStatusResponse) GetBody() *Pipeline {
	return s.Body
}

func (s *PutPipelineStatusResponse) SetHeaders(v map[string]*string) *PutPipelineStatusResponse {
	s.Headers = v
	return s
}

func (s *PutPipelineStatusResponse) SetStatusCode(v int32) *PutPipelineStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *PutPipelineStatusResponse) SetBody(v *Pipeline) *PutPipelineStatusResponse {
	s.Body = v
	return s
}

func (s *PutPipelineStatusResponse) Validate() error {
	return dara.Validate(s)
}
