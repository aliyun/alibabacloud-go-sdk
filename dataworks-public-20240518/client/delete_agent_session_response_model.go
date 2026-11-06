// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgentSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgentSessionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgentSessionResponseBody) *DeleteAgentSessionResponse
	GetBody() *DeleteAgentSessionResponseBody
}

type DeleteAgentSessionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgentSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgentSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSessionResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgentSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgentSessionResponse) GetBody() *DeleteAgentSessionResponseBody {
	return s.Body
}

func (s *DeleteAgentSessionResponse) SetHeaders(v map[string]*string) *DeleteAgentSessionResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentSessionResponse) SetStatusCode(v int32) *DeleteAgentSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentSessionResponse) SetBody(v *DeleteAgentSessionResponseBody) *DeleteAgentSessionResponse {
	s.Body = v
	return s
}

func (s *DeleteAgentSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
