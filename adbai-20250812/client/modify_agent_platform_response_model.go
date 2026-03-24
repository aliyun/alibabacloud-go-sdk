// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgentPlatformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAgentPlatformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAgentPlatformResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAgentPlatformResponseBody) *ModifyAgentPlatformResponse
	GetBody() *ModifyAgentPlatformResponseBody
}

type ModifyAgentPlatformResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAgentPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAgentPlatformResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgentPlatformResponse) GoString() string {
	return s.String()
}

func (s *ModifyAgentPlatformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAgentPlatformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAgentPlatformResponse) GetBody() *ModifyAgentPlatformResponseBody {
	return s.Body
}

func (s *ModifyAgentPlatformResponse) SetHeaders(v map[string]*string) *ModifyAgentPlatformResponse {
	s.Headers = v
	return s
}

func (s *ModifyAgentPlatformResponse) SetStatusCode(v int32) *ModifyAgentPlatformResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAgentPlatformResponse) SetBody(v *ModifyAgentPlatformResponseBody) *ModifyAgentPlatformResponse {
	s.Body = v
	return s
}

func (s *ModifyAgentPlatformResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
