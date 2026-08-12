// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataPipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataPipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataPipelineResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataPipelineResponseBody) *DeleteDataPipelineResponse
	GetBody() *DeleteDataPipelineResponseBody
}

type DeleteDataPipelineResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataPipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataPipelineResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataPipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataPipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataPipelineResponse) GetBody() *DeleteDataPipelineResponseBody {
	return s.Body
}

func (s *DeleteDataPipelineResponse) SetHeaders(v map[string]*string) *DeleteDataPipelineResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataPipelineResponse) SetStatusCode(v int32) *DeleteDataPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataPipelineResponse) SetBody(v *DeleteDataPipelineResponseBody) *DeleteDataPipelineResponse {
	s.Body = v
	return s
}

func (s *DeleteDataPipelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
