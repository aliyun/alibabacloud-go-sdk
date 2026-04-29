// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindPolarClawAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindPolarClawAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindPolarClawAgentResponse
	GetStatusCode() *int32
	SetBody(v *UnbindPolarClawAgentResponseBody) *UnbindPolarClawAgentResponse
	GetBody() *UnbindPolarClawAgentResponseBody
}

type UnbindPolarClawAgentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindPolarClawAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindPolarClawAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindPolarClawAgentResponse) GoString() string {
	return s.String()
}

func (s *UnbindPolarClawAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindPolarClawAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindPolarClawAgentResponse) GetBody() *UnbindPolarClawAgentResponseBody {
	return s.Body
}

func (s *UnbindPolarClawAgentResponse) SetHeaders(v map[string]*string) *UnbindPolarClawAgentResponse {
	s.Headers = v
	return s
}

func (s *UnbindPolarClawAgentResponse) SetStatusCode(v int32) *UnbindPolarClawAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindPolarClawAgentResponse) SetBody(v *UnbindPolarClawAgentResponseBody) *UnbindPolarClawAgentResponse {
	s.Body = v
	return s
}

func (s *UnbindPolarClawAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
