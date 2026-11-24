// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAgentDebugTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomAgentDebugTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomAgentDebugTokenResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomAgentDebugTokenResponseBody) *CreateCustomAgentDebugTokenResponse
	GetBody() *CreateCustomAgentDebugTokenResponseBody
}

type CreateCustomAgentDebugTokenResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomAgentDebugTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomAgentDebugTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentDebugTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentDebugTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomAgentDebugTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomAgentDebugTokenResponse) GetBody() *CreateCustomAgentDebugTokenResponseBody {
	return s.Body
}

func (s *CreateCustomAgentDebugTokenResponse) SetHeaders(v map[string]*string) *CreateCustomAgentDebugTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomAgentDebugTokenResponse) SetStatusCode(v int32) *CreateCustomAgentDebugTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomAgentDebugTokenResponse) SetBody(v *CreateCustomAgentDebugTokenResponseBody) *CreateCustomAgentDebugTokenResponse {
	s.Body = v
	return s
}

func (s *CreateCustomAgentDebugTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
