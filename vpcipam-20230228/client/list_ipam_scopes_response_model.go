// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamScopesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIpamScopesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIpamScopesResponse
	GetStatusCode() *int32
	SetBody(v *ListIpamScopesResponseBody) *ListIpamScopesResponse
	GetBody() *ListIpamScopesResponseBody
}

type ListIpamScopesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamScopesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamScopesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIpamScopesResponse) GoString() string {
	return s.String()
}

func (s *ListIpamScopesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIpamScopesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIpamScopesResponse) GetBody() *ListIpamScopesResponseBody {
	return s.Body
}

func (s *ListIpamScopesResponse) SetHeaders(v map[string]*string) *ListIpamScopesResponse {
	s.Headers = v
	return s
}

func (s *ListIpamScopesResponse) SetStatusCode(v int32) *ListIpamScopesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamScopesResponse) SetBody(v *ListIpamScopesResponseBody) *ListIpamScopesResponse {
	s.Body = v
	return s
}

func (s *ListIpamScopesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
