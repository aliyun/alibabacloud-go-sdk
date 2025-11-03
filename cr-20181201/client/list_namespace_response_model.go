// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *ListNamespaceResponseBody) *ListNamespaceResponse
	GetBody() *ListNamespaceResponseBody
}

type ListNamespaceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNamespaceResponse) GoString() string {
	return s.String()
}

func (s *ListNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNamespaceResponse) GetBody() *ListNamespaceResponseBody {
	return s.Body
}

func (s *ListNamespaceResponse) SetHeaders(v map[string]*string) *ListNamespaceResponse {
	s.Headers = v
	return s
}

func (s *ListNamespaceResponse) SetStatusCode(v int32) *ListNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNamespaceResponse) SetBody(v *ListNamespaceResponseBody) *ListNamespaceResponse {
	s.Body = v
	return s
}

func (s *ListNamespaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
