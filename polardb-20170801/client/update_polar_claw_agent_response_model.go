// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePolarClawAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePolarClawAgentResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePolarClawAgentResponseBody) *UpdatePolarClawAgentResponse
	GetBody() *UpdatePolarClawAgentResponseBody
}

type UpdatePolarClawAgentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePolarClawAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePolarClawAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawAgentResponse) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePolarClawAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePolarClawAgentResponse) GetBody() *UpdatePolarClawAgentResponseBody {
	return s.Body
}

func (s *UpdatePolarClawAgentResponse) SetHeaders(v map[string]*string) *UpdatePolarClawAgentResponse {
	s.Headers = v
	return s
}

func (s *UpdatePolarClawAgentResponse) SetStatusCode(v int32) *UpdatePolarClawAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePolarClawAgentResponse) SetBody(v *UpdatePolarClawAgentResponseBody) *UpdatePolarClawAgentResponse {
	s.Body = v
	return s
}

func (s *UpdatePolarClawAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
