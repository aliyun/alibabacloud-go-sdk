// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentIMChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentIMChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentIMChannelResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgentIMChannelResponseBody) *CreateAgentIMChannelResponse
	GetBody() *CreateAgentIMChannelResponseBody
}

type CreateAgentIMChannelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgentIMChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentIMChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentIMChannelResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentIMChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentIMChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentIMChannelResponse) GetBody() *CreateAgentIMChannelResponseBody {
	return s.Body
}

func (s *CreateAgentIMChannelResponse) SetHeaders(v map[string]*string) *CreateAgentIMChannelResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentIMChannelResponse) SetStatusCode(v int32) *CreateAgentIMChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentIMChannelResponse) SetBody(v *CreateAgentIMChannelResponseBody) *CreateAgentIMChannelResponse {
	s.Body = v
	return s
}

func (s *CreateAgentIMChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
