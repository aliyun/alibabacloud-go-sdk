// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExchangesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExchangesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExchangesResponse
	GetStatusCode() *int32
	SetBody(v *ListExchangesResponseBody) *ListExchangesResponse
	GetBody() *ListExchangesResponseBody
}

type ListExchangesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExchangesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExchangesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExchangesResponse) GoString() string {
	return s.String()
}

func (s *ListExchangesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExchangesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExchangesResponse) GetBody() *ListExchangesResponseBody {
	return s.Body
}

func (s *ListExchangesResponse) SetHeaders(v map[string]*string) *ListExchangesResponse {
	s.Headers = v
	return s
}

func (s *ListExchangesResponse) SetStatusCode(v int32) *ListExchangesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExchangesResponse) SetBody(v *ListExchangesResponseBody) *ListExchangesResponse {
	s.Body = v
	return s
}

func (s *ListExchangesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
