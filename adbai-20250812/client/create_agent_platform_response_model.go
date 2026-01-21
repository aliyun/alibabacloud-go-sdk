// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentPlatformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentPlatformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentPlatformResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgentPlatformResponseBody) *CreateAgentPlatformResponse
	GetBody() *CreateAgentPlatformResponseBody
}

type CreateAgentPlatformResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgentPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentPlatformResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentPlatformResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentPlatformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentPlatformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentPlatformResponse) GetBody() *CreateAgentPlatformResponseBody {
	return s.Body
}

func (s *CreateAgentPlatformResponse) SetHeaders(v map[string]*string) *CreateAgentPlatformResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentPlatformResponse) SetStatusCode(v int32) *CreateAgentPlatformResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentPlatformResponse) SetBody(v *CreateAgentPlatformResponseBody) *CreateAgentPlatformResponse {
	s.Body = v
	return s
}

func (s *CreateAgentPlatformResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
