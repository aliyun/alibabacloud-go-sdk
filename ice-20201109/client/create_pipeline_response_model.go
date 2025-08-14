// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePipelineResponse
	GetStatusCode() *int32
	SetBody(v *CreatePipelineResponseBody) *CreatePipelineResponse
	GetBody() *CreatePipelineResponseBody
}

type CreatePipelineResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineResponse) GoString() string {
	return s.String()
}

func (s *CreatePipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePipelineResponse) GetBody() *CreatePipelineResponseBody {
	return s.Body
}

func (s *CreatePipelineResponse) SetHeaders(v map[string]*string) *CreatePipelineResponse {
	s.Headers = v
	return s
}

func (s *CreatePipelineResponse) SetStatusCode(v int32) *CreatePipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePipelineResponse) SetBody(v *CreatePipelineResponseBody) *CreatePipelineResponse {
	s.Body = v
	return s
}

func (s *CreatePipelineResponse) Validate() error {
	return dara.Validate(s)
}
