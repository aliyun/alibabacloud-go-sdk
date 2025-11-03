// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNatIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNatIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNatIpsResponse
	GetStatusCode() *int32
	SetBody(v *ListNatIpsResponseBody) *ListNatIpsResponse
	GetBody() *ListNatIpsResponseBody
}

type ListNatIpsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNatIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNatIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNatIpsResponse) GoString() string {
	return s.String()
}

func (s *ListNatIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNatIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNatIpsResponse) GetBody() *ListNatIpsResponseBody {
	return s.Body
}

func (s *ListNatIpsResponse) SetHeaders(v map[string]*string) *ListNatIpsResponse {
	s.Headers = v
	return s
}

func (s *ListNatIpsResponse) SetStatusCode(v int32) *ListNatIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNatIpsResponse) SetBody(v *ListNatIpsResponseBody) *ListNatIpsResponse {
	s.Body = v
	return s
}

func (s *ListNatIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
