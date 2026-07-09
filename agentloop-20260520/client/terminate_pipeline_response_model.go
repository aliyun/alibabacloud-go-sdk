// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminatePipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TerminatePipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TerminatePipelineResponse
	GetStatusCode() *int32
	SetBody(v *TerminatePipelineResponseBody) *TerminatePipelineResponse
	GetBody() *TerminatePipelineResponseBody
}

type TerminatePipelineResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminatePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminatePipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s TerminatePipelineResponse) GoString() string {
	return s.String()
}

func (s *TerminatePipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TerminatePipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TerminatePipelineResponse) GetBody() *TerminatePipelineResponseBody {
	return s.Body
}

func (s *TerminatePipelineResponse) SetHeaders(v map[string]*string) *TerminatePipelineResponse {
	s.Headers = v
	return s
}

func (s *TerminatePipelineResponse) SetStatusCode(v int32) *TerminatePipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminatePipelineResponse) SetBody(v *TerminatePipelineResponseBody) *TerminatePipelineResponse {
	s.Body = v
	return s
}

func (s *TerminatePipelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
