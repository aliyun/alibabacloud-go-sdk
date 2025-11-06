// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExchangeDownstreamBindingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExchangeDownstreamBindingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExchangeDownstreamBindingsResponse
	GetStatusCode() *int32
	SetBody(v *ListExchangeDownstreamBindingsResponseBody) *ListExchangeDownstreamBindingsResponse
	GetBody() *ListExchangeDownstreamBindingsResponseBody
}

type ListExchangeDownstreamBindingsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExchangeDownstreamBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExchangeDownstreamBindingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeDownstreamBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListExchangeDownstreamBindingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExchangeDownstreamBindingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExchangeDownstreamBindingsResponse) GetBody() *ListExchangeDownstreamBindingsResponseBody {
	return s.Body
}

func (s *ListExchangeDownstreamBindingsResponse) SetHeaders(v map[string]*string) *ListExchangeDownstreamBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListExchangeDownstreamBindingsResponse) SetStatusCode(v int32) *ListExchangeDownstreamBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExchangeDownstreamBindingsResponse) SetBody(v *ListExchangeDownstreamBindingsResponseBody) *ListExchangeDownstreamBindingsResponse {
	s.Body = v
	return s
}

func (s *ListExchangeDownstreamBindingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
