// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPausePipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PausePipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PausePipelineResponse
	GetStatusCode() *int32
	SetBody(v *PausePipelineResponseBody) *PausePipelineResponse
	GetBody() *PausePipelineResponseBody
}

type PausePipelineResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PausePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PausePipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s PausePipelineResponse) GoString() string {
	return s.String()
}

func (s *PausePipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PausePipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PausePipelineResponse) GetBody() *PausePipelineResponseBody {
	return s.Body
}

func (s *PausePipelineResponse) SetHeaders(v map[string]*string) *PausePipelineResponse {
	s.Headers = v
	return s
}

func (s *PausePipelineResponse) SetStatusCode(v int32) *PausePipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *PausePipelineResponse) SetBody(v *PausePipelineResponseBody) *PausePipelineResponse {
	s.Body = v
	return s
}

func (s *PausePipelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
