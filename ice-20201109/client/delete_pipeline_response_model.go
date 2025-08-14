// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePipelineResponse
	GetStatusCode() *int32
	SetBody(v *DeletePipelineResponseBody) *DeletePipelineResponse
	GetBody() *DeletePipelineResponseBody
}

type DeletePipelineResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePipelineResponse) GoString() string {
	return s.String()
}

func (s *DeletePipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePipelineResponse) GetBody() *DeletePipelineResponseBody {
	return s.Body
}

func (s *DeletePipelineResponse) SetHeaders(v map[string]*string) *DeletePipelineResponse {
	s.Headers = v
	return s
}

func (s *DeletePipelineResponse) SetStatusCode(v int32) *DeletePipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePipelineResponse) SetBody(v *DeletePipelineResponseBody) *DeletePipelineResponse {
	s.Body = v
	return s
}

func (s *DeletePipelineResponse) Validate() error {
	return dara.Validate(s)
}
