// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNamespacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNamespacesResponse
	GetStatusCode() *int32
	SetBody(v *ListNamespacesResponseBody) *ListNamespacesResponse
	GetBody() *ListNamespacesResponseBody
}

type ListNamespacesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNamespacesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponse) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNamespacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNamespacesResponse) GetBody() *ListNamespacesResponseBody {
	return s.Body
}

func (s *ListNamespacesResponse) SetHeaders(v map[string]*string) *ListNamespacesResponse {
	s.Headers = v
	return s
}

func (s *ListNamespacesResponse) SetStatusCode(v int32) *ListNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNamespacesResponse) SetBody(v *ListNamespacesResponseBody) *ListNamespacesResponse {
	s.Body = v
	return s
}

func (s *ListNamespacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
