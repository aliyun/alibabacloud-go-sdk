// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataPipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataPipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataPipelineResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataPipelineResponseBody) *CreateDataPipelineResponse
	GetBody() *CreateDataPipelineResponseBody
}

type CreateDataPipelineResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataPipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponse) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataPipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataPipelineResponse) GetBody() *CreateDataPipelineResponseBody {
	return s.Body
}

func (s *CreateDataPipelineResponse) SetHeaders(v map[string]*string) *CreateDataPipelineResponse {
	s.Headers = v
	return s
}

func (s *CreateDataPipelineResponse) SetStatusCode(v int32) *CreateDataPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataPipelineResponse) SetBody(v *CreateDataPipelineResponseBody) *CreateDataPipelineResponse {
	s.Body = v
	return s
}

func (s *CreateDataPipelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
