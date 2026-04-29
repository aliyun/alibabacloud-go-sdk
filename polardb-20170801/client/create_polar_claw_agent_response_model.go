// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarClawAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePolarClawAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePolarClawAgentResponse
	GetStatusCode() *int32
	SetBody(v *CreatePolarClawAgentResponseBody) *CreatePolarClawAgentResponse
	GetBody() *CreatePolarClawAgentResponseBody
}

type CreatePolarClawAgentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolarClawAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolarClawAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawAgentResponse) GoString() string {
	return s.String()
}

func (s *CreatePolarClawAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePolarClawAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePolarClawAgentResponse) GetBody() *CreatePolarClawAgentResponseBody {
	return s.Body
}

func (s *CreatePolarClawAgentResponse) SetHeaders(v map[string]*string) *CreatePolarClawAgentResponse {
	s.Headers = v
	return s
}

func (s *CreatePolarClawAgentResponse) SetStatusCode(v int32) *CreatePolarClawAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolarClawAgentResponse) SetBody(v *CreatePolarClawAgentResponseBody) *CreatePolarClawAgentResponse {
	s.Body = v
	return s
}

func (s *CreatePolarClawAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
