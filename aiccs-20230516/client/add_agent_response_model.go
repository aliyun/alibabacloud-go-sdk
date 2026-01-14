// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAgentResponse
	GetStatusCode() *int32
	SetBody(v *AddAgentResponseBody) *AddAgentResponse
	GetBody() *AddAgentResponseBody
}

type AddAgentResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAgentResponse) GoString() string {
	return s.String()
}

func (s *AddAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAgentResponse) GetBody() *AddAgentResponseBody {
	return s.Body
}

func (s *AddAgentResponse) SetHeaders(v map[string]*string) *AddAgentResponse {
	s.Headers = v
	return s
}

func (s *AddAgentResponse) SetStatusCode(v int32) *AddAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAgentResponse) SetBody(v *AddAgentResponseBody) *AddAgentResponse {
	s.Body = v
	return s
}

func (s *AddAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
