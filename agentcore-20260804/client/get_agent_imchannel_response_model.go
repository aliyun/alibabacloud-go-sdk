// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentIMChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentIMChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentIMChannelResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentIMChannelResponseBody) *GetAgentIMChannelResponse
	GetBody() *GetAgentIMChannelResponseBody
}

type GetAgentIMChannelResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentIMChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentIMChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentIMChannelResponse) GoString() string {
	return s.String()
}

func (s *GetAgentIMChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentIMChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentIMChannelResponse) GetBody() *GetAgentIMChannelResponseBody {
	return s.Body
}

func (s *GetAgentIMChannelResponse) SetHeaders(v map[string]*string) *GetAgentIMChannelResponse {
	s.Headers = v
	return s
}

func (s *GetAgentIMChannelResponse) SetStatusCode(v int32) *GetAgentIMChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentIMChannelResponse) SetBody(v *GetAgentIMChannelResponseBody) *GetAgentIMChannelResponse {
	s.Body = v
	return s
}

func (s *GetAgentIMChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
