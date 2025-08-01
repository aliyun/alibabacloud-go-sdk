// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSentinelBlockFallbackDefinitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSentinelBlockFallbackDefinitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSentinelBlockFallbackDefinitionResponse
	GetStatusCode() *int32
	SetBody(v *CreateSentinelBlockFallbackDefinitionResponseBody) *CreateSentinelBlockFallbackDefinitionResponse
	GetBody() *CreateSentinelBlockFallbackDefinitionResponseBody
}

type CreateSentinelBlockFallbackDefinitionResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSentinelBlockFallbackDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSentinelBlockFallbackDefinitionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSentinelBlockFallbackDefinitionResponse) GoString() string {
	return s.String()
}

func (s *CreateSentinelBlockFallbackDefinitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSentinelBlockFallbackDefinitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSentinelBlockFallbackDefinitionResponse) GetBody() *CreateSentinelBlockFallbackDefinitionResponseBody {
	return s.Body
}

func (s *CreateSentinelBlockFallbackDefinitionResponse) SetHeaders(v map[string]*string) *CreateSentinelBlockFallbackDefinitionResponse {
	s.Headers = v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionResponse) SetStatusCode(v int32) *CreateSentinelBlockFallbackDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionResponse) SetBody(v *CreateSentinelBlockFallbackDefinitionResponseBody) *CreateSentinelBlockFallbackDefinitionResponse {
	s.Body = v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionResponse) Validate() error {
	return dara.Validate(s)
}
