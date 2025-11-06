// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSentinelBlockFallbackDefinitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindSentinelBlockFallbackDefinitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindSentinelBlockFallbackDefinitionResponse
	GetStatusCode() *int32
	SetBody(v *BindSentinelBlockFallbackDefinitionResponseBody) *BindSentinelBlockFallbackDefinitionResponse
	GetBody() *BindSentinelBlockFallbackDefinitionResponseBody
}

type BindSentinelBlockFallbackDefinitionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindSentinelBlockFallbackDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindSentinelBlockFallbackDefinitionResponse) String() string {
	return dara.Prettify(s)
}

func (s BindSentinelBlockFallbackDefinitionResponse) GoString() string {
	return s.String()
}

func (s *BindSentinelBlockFallbackDefinitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindSentinelBlockFallbackDefinitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindSentinelBlockFallbackDefinitionResponse) GetBody() *BindSentinelBlockFallbackDefinitionResponseBody {
	return s.Body
}

func (s *BindSentinelBlockFallbackDefinitionResponse) SetHeaders(v map[string]*string) *BindSentinelBlockFallbackDefinitionResponse {
	s.Headers = v
	return s
}

func (s *BindSentinelBlockFallbackDefinitionResponse) SetStatusCode(v int32) *BindSentinelBlockFallbackDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *BindSentinelBlockFallbackDefinitionResponse) SetBody(v *BindSentinelBlockFallbackDefinitionResponseBody) *BindSentinelBlockFallbackDefinitionResponse {
	s.Body = v
	return s
}

func (s *BindSentinelBlockFallbackDefinitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
