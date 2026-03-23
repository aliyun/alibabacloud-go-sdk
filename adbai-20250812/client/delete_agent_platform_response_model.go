// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentPlatformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgentPlatformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgentPlatformResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgentPlatformResponseBody) *DeleteAgentPlatformResponse
	GetBody() *DeleteAgentPlatformResponseBody
}

type DeleteAgentPlatformResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgentPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgentPlatformResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentPlatformResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentPlatformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgentPlatformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgentPlatformResponse) GetBody() *DeleteAgentPlatformResponseBody {
	return s.Body
}

func (s *DeleteAgentPlatformResponse) SetHeaders(v map[string]*string) *DeleteAgentPlatformResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentPlatformResponse) SetStatusCode(v int32) *DeleteAgentPlatformResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentPlatformResponse) SetBody(v *DeleteAgentPlatformResponseBody) *DeleteAgentPlatformResponse {
	s.Body = v
	return s
}

func (s *DeleteAgentPlatformResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
