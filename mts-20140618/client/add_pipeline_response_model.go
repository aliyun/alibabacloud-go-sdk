// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPipelineResponse
	GetStatusCode() *int32
	SetBody(v *AddPipelineResponseBody) *AddPipelineResponse
	GetBody() *AddPipelineResponseBody
}

type AddPipelineResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPipelineResponse) GoString() string {
	return s.String()
}

func (s *AddPipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPipelineResponse) GetBody() *AddPipelineResponseBody {
	return s.Body
}

func (s *AddPipelineResponse) SetHeaders(v map[string]*string) *AddPipelineResponse {
	s.Headers = v
	return s
}

func (s *AddPipelineResponse) SetStatusCode(v int32) *AddPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPipelineResponse) SetBody(v *AddPipelineResponseBody) *AddPipelineResponse {
	s.Body = v
	return s
}

func (s *AddPipelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
