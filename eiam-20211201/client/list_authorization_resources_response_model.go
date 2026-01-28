// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizationResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizationResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizationResourcesResponseBody) *ListAuthorizationResourcesResponse
	GetBody() *ListAuthorizationResourcesResponseBody
}

type ListAuthorizationResourcesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizationResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizationResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizationResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizationResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizationResourcesResponse) GetBody() *ListAuthorizationResourcesResponseBody {
	return s.Body
}

func (s *ListAuthorizationResourcesResponse) SetHeaders(v map[string]*string) *ListAuthorizationResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizationResourcesResponse) SetStatusCode(v int32) *ListAuthorizationResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizationResourcesResponse) SetBody(v *ListAuthorizationResourcesResponseBody) *ListAuthorizationResourcesResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizationResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
