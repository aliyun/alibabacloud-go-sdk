// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *RunWorkflowResponseBody) *RunWorkflowResponse
	GetBody() *RunWorkflowResponseBody
}

type RunWorkflowResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s RunWorkflowResponse) GoString() string {
	return s.String()
}

func (s *RunWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunWorkflowResponse) GetBody() *RunWorkflowResponseBody {
	return s.Body
}

func (s *RunWorkflowResponse) SetHeaders(v map[string]*string) *RunWorkflowResponse {
	s.Headers = v
	return s
}

func (s *RunWorkflowResponse) SetStatusCode(v int32) *RunWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *RunWorkflowResponse) SetBody(v *RunWorkflowResponseBody) *RunWorkflowResponse {
	s.Body = v
	return s
}

func (s *RunWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
