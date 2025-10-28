// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinuePipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ContinuePipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ContinuePipelineResponse
	GetStatusCode() *int32
	SetBody(v *ContinuePipelineResponseBody) *ContinuePipelineResponse
	GetBody() *ContinuePipelineResponseBody
}

type ContinuePipelineResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContinuePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContinuePipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s ContinuePipelineResponse) GoString() string {
	return s.String()
}

func (s *ContinuePipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ContinuePipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ContinuePipelineResponse) GetBody() *ContinuePipelineResponseBody {
	return s.Body
}

func (s *ContinuePipelineResponse) SetHeaders(v map[string]*string) *ContinuePipelineResponse {
	s.Headers = v
	return s
}

func (s *ContinuePipelineResponse) SetStatusCode(v int32) *ContinuePipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinuePipelineResponse) SetBody(v *ContinuePipelineResponseBody) *ContinuePipelineResponse {
	s.Body = v
	return s
}

func (s *ContinuePipelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
