// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataPipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataPipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataPipelineResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataPipelineResponseBody) *UpdateDataPipelineResponse
	GetBody() *UpdateDataPipelineResponseBody
}

type UpdateDataPipelineResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataPipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataPipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataPipelineResponse) GetBody() *UpdateDataPipelineResponseBody {
	return s.Body
}

func (s *UpdateDataPipelineResponse) SetHeaders(v map[string]*string) *UpdateDataPipelineResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataPipelineResponse) SetStatusCode(v int32) *UpdateDataPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataPipelineResponse) SetBody(v *UpdateDataPipelineResponseBody) *UpdateDataPipelineResponse {
	s.Body = v
	return s
}

func (s *UpdateDataPipelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
