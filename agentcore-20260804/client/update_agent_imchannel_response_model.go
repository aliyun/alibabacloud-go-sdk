// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentIMChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgentIMChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgentIMChannelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAgentIMChannelResponseBody) *UpdateAgentIMChannelResponse
	GetBody() *UpdateAgentIMChannelResponseBody
}

type UpdateAgentIMChannelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgentIMChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentIMChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentIMChannelResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentIMChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgentIMChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgentIMChannelResponse) GetBody() *UpdateAgentIMChannelResponseBody {
	return s.Body
}

func (s *UpdateAgentIMChannelResponse) SetHeaders(v map[string]*string) *UpdateAgentIMChannelResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentIMChannelResponse) SetStatusCode(v int32) *UpdateAgentIMChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentIMChannelResponse) SetBody(v *UpdateAgentIMChannelResponseBody) *UpdateAgentIMChannelResponse {
	s.Body = v
	return s
}

func (s *UpdateAgentIMChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
