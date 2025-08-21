// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOIDCProvidersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOIDCProvidersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOIDCProvidersResponse
	GetStatusCode() *int32
	SetBody(v *ListOIDCProvidersResponseBody) *ListOIDCProvidersResponse
	GetBody() *ListOIDCProvidersResponseBody
}

type ListOIDCProvidersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOIDCProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOIDCProvidersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOIDCProvidersResponse) GoString() string {
	return s.String()
}

func (s *ListOIDCProvidersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOIDCProvidersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOIDCProvidersResponse) GetBody() *ListOIDCProvidersResponseBody {
	return s.Body
}

func (s *ListOIDCProvidersResponse) SetHeaders(v map[string]*string) *ListOIDCProvidersResponse {
	s.Headers = v
	return s
}

func (s *ListOIDCProvidersResponse) SetStatusCode(v int32) *ListOIDCProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOIDCProvidersResponse) SetBody(v *ListOIDCProvidersResponseBody) *ListOIDCProvidersResponse {
	s.Body = v
	return s
}

func (s *ListOIDCProvidersResponse) Validate() error {
	return dara.Validate(s)
}
