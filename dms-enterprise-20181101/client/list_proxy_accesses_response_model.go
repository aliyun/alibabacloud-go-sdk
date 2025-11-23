// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProxyAccessesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProxyAccessesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProxyAccessesResponse
	GetStatusCode() *int32
	SetBody(v *ListProxyAccessesResponseBody) *ListProxyAccessesResponse
	GetBody() *ListProxyAccessesResponseBody
}

type ListProxyAccessesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProxyAccessesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProxyAccessesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProxyAccessesResponse) GoString() string {
	return s.String()
}

func (s *ListProxyAccessesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProxyAccessesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProxyAccessesResponse) GetBody() *ListProxyAccessesResponseBody {
	return s.Body
}

func (s *ListProxyAccessesResponse) SetHeaders(v map[string]*string) *ListProxyAccessesResponse {
	s.Headers = v
	return s
}

func (s *ListProxyAccessesResponse) SetStatusCode(v int32) *ListProxyAccessesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProxyAccessesResponse) SetBody(v *ListProxyAccessesResponseBody) *ListProxyAccessesResponse {
	s.Body = v
	return s
}

func (s *ListProxyAccessesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
