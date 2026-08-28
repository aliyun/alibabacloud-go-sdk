// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentIMChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgentIMChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgentIMChannelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgentIMChannelResponseBody) *DeleteAgentIMChannelResponse
	GetBody() *DeleteAgentIMChannelResponseBody
}

type DeleteAgentIMChannelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgentIMChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgentIMChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentIMChannelResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentIMChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgentIMChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgentIMChannelResponse) GetBody() *DeleteAgentIMChannelResponseBody {
	return s.Body
}

func (s *DeleteAgentIMChannelResponse) SetHeaders(v map[string]*string) *DeleteAgentIMChannelResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentIMChannelResponse) SetStatusCode(v int32) *DeleteAgentIMChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentIMChannelResponse) SetBody(v *DeleteAgentIMChannelResponseBody) *DeleteAgentIMChannelResponse {
	s.Body = v
	return s
}

func (s *DeleteAgentIMChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
