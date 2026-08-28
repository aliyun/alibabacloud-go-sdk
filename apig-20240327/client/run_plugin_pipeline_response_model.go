// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPluginPipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunPluginPipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunPluginPipelineResponse
	GetStatusCode() *int32
	SetBody(v *RunPluginPipelineResponseBody) *RunPluginPipelineResponse
	GetBody() *RunPluginPipelineResponseBody
}

type RunPluginPipelineResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunPluginPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunPluginPipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s RunPluginPipelineResponse) GoString() string {
	return s.String()
}

func (s *RunPluginPipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunPluginPipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunPluginPipelineResponse) GetBody() *RunPluginPipelineResponseBody {
	return s.Body
}

func (s *RunPluginPipelineResponse) SetHeaders(v map[string]*string) *RunPluginPipelineResponse {
	s.Headers = v
	return s
}

func (s *RunPluginPipelineResponse) SetStatusCode(v int32) *RunPluginPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *RunPluginPipelineResponse) SetBody(v *RunPluginPipelineResponseBody) *RunPluginPipelineResponse {
	s.Body = v
	return s
}

func (s *RunPluginPipelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
