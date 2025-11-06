// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExchangeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExchangeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExchangeResponse
	GetStatusCode() *int32
	SetBody(v *ListExchangeResponseBody) *ListExchangeResponse
	GetBody() *ListExchangeResponseBody
}

type ListExchangeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExchangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExchangeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeResponse) GoString() string {
	return s.String()
}

func (s *ListExchangeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExchangeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExchangeResponse) GetBody() *ListExchangeResponseBody {
	return s.Body
}

func (s *ListExchangeResponse) SetHeaders(v map[string]*string) *ListExchangeResponse {
	s.Headers = v
	return s
}

func (s *ListExchangeResponse) SetStatusCode(v int32) *ListExchangeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExchangeResponse) SetBody(v *ListExchangeResponseBody) *ListExchangeResponse {
	s.Body = v
	return s
}

func (s *ListExchangeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
