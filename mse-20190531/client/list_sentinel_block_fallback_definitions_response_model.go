// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSentinelBlockFallbackDefinitionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSentinelBlockFallbackDefinitionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSentinelBlockFallbackDefinitionsResponse
	GetStatusCode() *int32
	SetBody(v *ListSentinelBlockFallbackDefinitionsResponseBody) *ListSentinelBlockFallbackDefinitionsResponse
	GetBody() *ListSentinelBlockFallbackDefinitionsResponseBody
}

type ListSentinelBlockFallbackDefinitionsResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSentinelBlockFallbackDefinitionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSentinelBlockFallbackDefinitionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSentinelBlockFallbackDefinitionsResponse) GoString() string {
	return s.String()
}

func (s *ListSentinelBlockFallbackDefinitionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSentinelBlockFallbackDefinitionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSentinelBlockFallbackDefinitionsResponse) GetBody() *ListSentinelBlockFallbackDefinitionsResponseBody {
	return s.Body
}

func (s *ListSentinelBlockFallbackDefinitionsResponse) SetHeaders(v map[string]*string) *ListSentinelBlockFallbackDefinitionsResponse {
	s.Headers = v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsResponse) SetStatusCode(v int32) *ListSentinelBlockFallbackDefinitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsResponse) SetBody(v *ListSentinelBlockFallbackDefinitionsResponseBody) *ListSentinelBlockFallbackDefinitionsResponse {
	s.Body = v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsResponse) Validate() error {
	return dara.Validate(s)
}
