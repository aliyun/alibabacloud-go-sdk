// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgentSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgentSpaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgentSpaceResponseBody) *DeleteAgentSpaceResponse
	GetBody() *DeleteAgentSpaceResponseBody
}

type DeleteAgentSpaceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgentSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgentSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSpaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgentSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgentSpaceResponse) GetBody() *DeleteAgentSpaceResponseBody {
	return s.Body
}

func (s *DeleteAgentSpaceResponse) SetHeaders(v map[string]*string) *DeleteAgentSpaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentSpaceResponse) SetStatusCode(v int32) *DeleteAgentSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentSpaceResponse) SetBody(v *DeleteAgentSpaceResponseBody) *DeleteAgentSpaceResponse {
	s.Body = v
	return s
}

func (s *DeleteAgentSpaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
