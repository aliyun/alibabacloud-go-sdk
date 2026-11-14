// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgentInstanceMemoryConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAgentInstanceMemoryConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAgentInstanceMemoryConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAgentInstanceMemoryConfigResponseBody) *ModifyAgentInstanceMemoryConfigResponse
	GetBody() *ModifyAgentInstanceMemoryConfigResponseBody
}

type ModifyAgentInstanceMemoryConfigResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAgentInstanceMemoryConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAgentInstanceMemoryConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgentInstanceMemoryConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyAgentInstanceMemoryConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAgentInstanceMemoryConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAgentInstanceMemoryConfigResponse) GetBody() *ModifyAgentInstanceMemoryConfigResponseBody {
	return s.Body
}

func (s *ModifyAgentInstanceMemoryConfigResponse) SetHeaders(v map[string]*string) *ModifyAgentInstanceMemoryConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyAgentInstanceMemoryConfigResponse) SetStatusCode(v int32) *ModifyAgentInstanceMemoryConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAgentInstanceMemoryConfigResponse) SetBody(v *ModifyAgentInstanceMemoryConfigResponseBody) *ModifyAgentInstanceMemoryConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyAgentInstanceMemoryConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
