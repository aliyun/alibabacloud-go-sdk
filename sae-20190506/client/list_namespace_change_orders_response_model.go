// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespaceChangeOrdersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNamespaceChangeOrdersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNamespaceChangeOrdersResponse
	GetStatusCode() *int32
	SetBody(v *ListNamespaceChangeOrdersResponseBody) *ListNamespaceChangeOrdersResponse
	GetBody() *ListNamespaceChangeOrdersResponseBody
}

type ListNamespaceChangeOrdersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNamespaceChangeOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNamespaceChangeOrdersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNamespaceChangeOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListNamespaceChangeOrdersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNamespaceChangeOrdersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNamespaceChangeOrdersResponse) GetBody() *ListNamespaceChangeOrdersResponseBody {
	return s.Body
}

func (s *ListNamespaceChangeOrdersResponse) SetHeaders(v map[string]*string) *ListNamespaceChangeOrdersResponse {
	s.Headers = v
	return s
}

func (s *ListNamespaceChangeOrdersResponse) SetStatusCode(v int32) *ListNamespaceChangeOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponse) SetBody(v *ListNamespaceChangeOrdersResponseBody) *ListNamespaceChangeOrdersResponse {
	s.Body = v
	return s
}

func (s *ListNamespaceChangeOrdersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
