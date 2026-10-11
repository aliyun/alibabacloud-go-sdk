// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomAgentResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomAgentResponseBody) *GetCustomAgentResponse
	GetBody() *GetCustomAgentResponseBody
}

type GetCustomAgentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponse) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomAgentResponse) GetBody() *GetCustomAgentResponseBody {
	return s.Body
}

func (s *GetCustomAgentResponse) SetHeaders(v map[string]*string) *GetCustomAgentResponse {
	s.Headers = v
	return s
}

func (s *GetCustomAgentResponse) SetStatusCode(v int32) *GetCustomAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomAgentResponse) SetBody(v *GetCustomAgentResponseBody) *GetCustomAgentResponse {
	s.Body = v
	return s
}

func (s *GetCustomAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
