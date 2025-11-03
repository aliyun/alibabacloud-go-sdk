// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExchangeUpStreamBindingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExchangeUpStreamBindingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExchangeUpStreamBindingsResponse
	GetStatusCode() *int32
	SetBody(v *ListExchangeUpStreamBindingsResponseBody) *ListExchangeUpStreamBindingsResponse
	GetBody() *ListExchangeUpStreamBindingsResponseBody
}

type ListExchangeUpStreamBindingsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExchangeUpStreamBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExchangeUpStreamBindingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeUpStreamBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExchangeUpStreamBindingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExchangeUpStreamBindingsResponse) GetBody() *ListExchangeUpStreamBindingsResponseBody {
	return s.Body
}

func (s *ListExchangeUpStreamBindingsResponse) SetHeaders(v map[string]*string) *ListExchangeUpStreamBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListExchangeUpStreamBindingsResponse) SetStatusCode(v int32) *ListExchangeUpStreamBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponse) SetBody(v *ListExchangeUpStreamBindingsResponseBody) *ListExchangeUpStreamBindingsResponse {
	s.Body = v
	return s
}

func (s *ListExchangeUpStreamBindingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
