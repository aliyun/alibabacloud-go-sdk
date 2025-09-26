// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRuntimeVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishRuntimeVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishRuntimeVersionResponse
	GetStatusCode() *int32
	SetBody(v *AgentRuntimeVersionResult) *PublishRuntimeVersionResponse
	GetBody() *AgentRuntimeVersionResult
}

type PublishRuntimeVersionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgentRuntimeVersionResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishRuntimeVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishRuntimeVersionResponse) GoString() string {
	return s.String()
}

func (s *PublishRuntimeVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishRuntimeVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishRuntimeVersionResponse) GetBody() *AgentRuntimeVersionResult {
	return s.Body
}

func (s *PublishRuntimeVersionResponse) SetHeaders(v map[string]*string) *PublishRuntimeVersionResponse {
	s.Headers = v
	return s
}

func (s *PublishRuntimeVersionResponse) SetStatusCode(v int32) *PublishRuntimeVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishRuntimeVersionResponse) SetBody(v *AgentRuntimeVersionResult) *PublishRuntimeVersionResponse {
	s.Body = v
	return s
}

func (s *PublishRuntimeVersionResponse) Validate() error {
	return dara.Validate(s)
}
