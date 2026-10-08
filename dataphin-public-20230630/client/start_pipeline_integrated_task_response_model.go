// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPipelineIntegratedTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartPipelineIntegratedTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartPipelineIntegratedTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartPipelineIntegratedTaskResponseBody) *StartPipelineIntegratedTaskResponse
	GetBody() *StartPipelineIntegratedTaskResponseBody
}

type StartPipelineIntegratedTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartPipelineIntegratedTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartPipelineIntegratedTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartPipelineIntegratedTaskResponse) GoString() string {
	return s.String()
}

func (s *StartPipelineIntegratedTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartPipelineIntegratedTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartPipelineIntegratedTaskResponse) GetBody() *StartPipelineIntegratedTaskResponseBody {
	return s.Body
}

func (s *StartPipelineIntegratedTaskResponse) SetHeaders(v map[string]*string) *StartPipelineIntegratedTaskResponse {
	s.Headers = v
	return s
}

func (s *StartPipelineIntegratedTaskResponse) SetStatusCode(v int32) *StartPipelineIntegratedTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartPipelineIntegratedTaskResponse) SetBody(v *StartPipelineIntegratedTaskResponseBody) *StartPipelineIntegratedTaskResponse {
	s.Body = v
	return s
}

func (s *StartPipelineIntegratedTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
