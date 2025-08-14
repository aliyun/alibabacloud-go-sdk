// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePipelineResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePipelineResponseBody) *UpdatePipelineResponse
	GetBody() *UpdatePipelineResponseBody
}

type UpdatePipelineResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineResponse) GoString() string {
	return s.String()
}

func (s *UpdatePipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePipelineResponse) GetBody() *UpdatePipelineResponseBody {
	return s.Body
}

func (s *UpdatePipelineResponse) SetHeaders(v map[string]*string) *UpdatePipelineResponse {
	s.Headers = v
	return s
}

func (s *UpdatePipelineResponse) SetStatusCode(v int32) *UpdatePipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePipelineResponse) SetBody(v *UpdatePipelineResponseBody) *UpdatePipelineResponse {
	s.Body = v
	return s
}

func (s *UpdatePipelineResponse) Validate() error {
	return dara.Validate(s)
}
