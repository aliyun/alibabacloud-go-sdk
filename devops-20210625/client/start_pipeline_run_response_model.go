// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPipelineRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartPipelineRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartPipelineRunResponse
	GetStatusCode() *int32
	SetBody(v *StartPipelineRunResponseBody) *StartPipelineRunResponse
	GetBody() *StartPipelineRunResponseBody
}

type StartPipelineRunResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartPipelineRunResponse) String() string {
	return dara.Prettify(s)
}

func (s StartPipelineRunResponse) GoString() string {
	return s.String()
}

func (s *StartPipelineRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartPipelineRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartPipelineRunResponse) GetBody() *StartPipelineRunResponseBody {
	return s.Body
}

func (s *StartPipelineRunResponse) SetHeaders(v map[string]*string) *StartPipelineRunResponse {
	s.Headers = v
	return s
}

func (s *StartPipelineRunResponse) SetStatusCode(v int32) *StartPipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *StartPipelineRunResponse) SetBody(v *StartPipelineRunResponseBody) *StartPipelineRunResponse {
	s.Body = v
	return s
}

func (s *StartPipelineRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
