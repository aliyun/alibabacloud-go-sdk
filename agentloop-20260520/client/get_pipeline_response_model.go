// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPipelineResponse
	GetStatusCode() *int32
	SetBody(v *GetPipelineResponseBody) *GetPipelineResponse
	GetBody() *GetPipelineResponseBody
}

type GetPipelineResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPipelineResponse) GetBody() *GetPipelineResponseBody {
	return s.Body
}

func (s *GetPipelineResponse) SetHeaders(v map[string]*string) *GetPipelineResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineResponse) SetStatusCode(v int32) *GetPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineResponse) SetBody(v *GetPipelineResponseBody) *GetPipelineResponse {
	s.Body = v
	return s
}

func (s *GetPipelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
