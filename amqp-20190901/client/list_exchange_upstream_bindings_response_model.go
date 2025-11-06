// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExchangeUpstreamBindingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExchangeUpstreamBindingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExchangeUpstreamBindingsResponse
	GetStatusCode() *int32
	SetBody(v *ListExchangeUpstreamBindingsResponseBody) *ListExchangeUpstreamBindingsResponse
	GetBody() *ListExchangeUpstreamBindingsResponseBody
}

type ListExchangeUpstreamBindingsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExchangeUpstreamBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExchangeUpstreamBindingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeUpstreamBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListExchangeUpstreamBindingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExchangeUpstreamBindingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExchangeUpstreamBindingsResponse) GetBody() *ListExchangeUpstreamBindingsResponseBody {
	return s.Body
}

func (s *ListExchangeUpstreamBindingsResponse) SetHeaders(v map[string]*string) *ListExchangeUpstreamBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListExchangeUpstreamBindingsResponse) SetStatusCode(v int32) *ListExchangeUpstreamBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExchangeUpstreamBindingsResponse) SetBody(v *ListExchangeUpstreamBindingsResponseBody) *ListExchangeUpstreamBindingsResponse {
	s.Body = v
	return s
}

func (s *ListExchangeUpstreamBindingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
