// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePipelineNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePipelineNodeResponse
	GetStatusCode() *int32
	SetBody(v *CreatePipelineNodeResponseBody) *CreatePipelineNodeResponse
	GetBody() *CreatePipelineNodeResponseBody
}

type CreatePipelineNodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePipelineNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePipelineNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineNodeResponse) GoString() string {
	return s.String()
}

func (s *CreatePipelineNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePipelineNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePipelineNodeResponse) GetBody() *CreatePipelineNodeResponseBody {
	return s.Body
}

func (s *CreatePipelineNodeResponse) SetHeaders(v map[string]*string) *CreatePipelineNodeResponse {
	s.Headers = v
	return s
}

func (s *CreatePipelineNodeResponse) SetStatusCode(v int32) *CreatePipelineNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePipelineNodeResponse) SetBody(v *CreatePipelineNodeResponseBody) *CreatePipelineNodeResponse {
	s.Body = v
	return s
}

func (s *CreatePipelineNodeResponse) Validate() error {
	return dara.Validate(s)
}
