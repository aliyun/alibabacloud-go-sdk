// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIngressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIngressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIngressesResponse
	GetStatusCode() *int32
	SetBody(v *ListIngressesResponseBody) *ListIngressesResponse
	GetBody() *ListIngressesResponseBody
}

type ListIngressesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIngressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIngressesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIngressesResponse) GoString() string {
	return s.String()
}

func (s *ListIngressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIngressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIngressesResponse) GetBody() *ListIngressesResponseBody {
	return s.Body
}

func (s *ListIngressesResponse) SetHeaders(v map[string]*string) *ListIngressesResponse {
	s.Headers = v
	return s
}

func (s *ListIngressesResponse) SetStatusCode(v int32) *ListIngressesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIngressesResponse) SetBody(v *ListIngressesResponseBody) *ListIngressesResponse {
	s.Body = v
	return s
}

func (s *ListIngressesResponse) Validate() error {
	return dara.Validate(s)
}
