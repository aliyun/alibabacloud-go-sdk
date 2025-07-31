// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIcebergNamespaceDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIcebergNamespaceDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIcebergNamespaceDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListIcebergNamespaceDetailsResponseBody) *ListIcebergNamespaceDetailsResponse
	GetBody() *ListIcebergNamespaceDetailsResponseBody
}

type ListIcebergNamespaceDetailsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIcebergNamespaceDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIcebergNamespaceDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIcebergNamespaceDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListIcebergNamespaceDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIcebergNamespaceDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIcebergNamespaceDetailsResponse) GetBody() *ListIcebergNamespaceDetailsResponseBody {
	return s.Body
}

func (s *ListIcebergNamespaceDetailsResponse) SetHeaders(v map[string]*string) *ListIcebergNamespaceDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListIcebergNamespaceDetailsResponse) SetStatusCode(v int32) *ListIcebergNamespaceDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIcebergNamespaceDetailsResponse) SetBody(v *ListIcebergNamespaceDetailsResponseBody) *ListIcebergNamespaceDetailsResponse {
	s.Body = v
	return s
}

func (s *ListIcebergNamespaceDetailsResponse) Validate() error {
	return dara.Validate(s)
}
