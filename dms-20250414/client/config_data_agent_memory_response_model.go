// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigDataAgentMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigDataAgentMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigDataAgentMemoryResponse
	GetStatusCode() *int32
	SetBody(v *ConfigDataAgentMemoryResponseBody) *ConfigDataAgentMemoryResponse
	GetBody() *ConfigDataAgentMemoryResponseBody
}

type ConfigDataAgentMemoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigDataAgentMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigDataAgentMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigDataAgentMemoryResponse) GoString() string {
	return s.String()
}

func (s *ConfigDataAgentMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigDataAgentMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigDataAgentMemoryResponse) GetBody() *ConfigDataAgentMemoryResponseBody {
	return s.Body
}

func (s *ConfigDataAgentMemoryResponse) SetHeaders(v map[string]*string) *ConfigDataAgentMemoryResponse {
	s.Headers = v
	return s
}

func (s *ConfigDataAgentMemoryResponse) SetStatusCode(v int32) *ConfigDataAgentMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigDataAgentMemoryResponse) SetBody(v *ConfigDataAgentMemoryResponseBody) *ConfigDataAgentMemoryResponse {
	s.Body = v
	return s
}

func (s *ConfigDataAgentMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
