// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndPulishAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAndPulishAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAndPulishAgentResponse
	GetStatusCode() *int32
	SetBody(v *CreateAndPulishAgentResponseBody) *CreateAndPulishAgentResponse
	GetBody() *CreateAndPulishAgentResponseBody
}

type CreateAndPulishAgentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAndPulishAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAndPulishAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPulishAgentResponse) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAndPulishAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAndPulishAgentResponse) GetBody() *CreateAndPulishAgentResponseBody {
	return s.Body
}

func (s *CreateAndPulishAgentResponse) SetHeaders(v map[string]*string) *CreateAndPulishAgentResponse {
	s.Headers = v
	return s
}

func (s *CreateAndPulishAgentResponse) SetStatusCode(v int32) *CreateAndPulishAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAndPulishAgentResponse) SetBody(v *CreateAndPulishAgentResponseBody) *CreateAndPulishAgentResponse {
	s.Body = v
	return s
}

func (s *CreateAndPulishAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
