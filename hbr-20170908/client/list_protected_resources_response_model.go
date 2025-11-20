// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProtectedResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProtectedResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProtectedResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListProtectedResourcesResponseBody) *ListProtectedResourcesResponse
	GetBody() *ListProtectedResourcesResponseBody
}

type ListProtectedResourcesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProtectedResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProtectedResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListProtectedResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProtectedResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProtectedResourcesResponse) GetBody() *ListProtectedResourcesResponseBody {
	return s.Body
}

func (s *ListProtectedResourcesResponse) SetHeaders(v map[string]*string) *ListProtectedResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListProtectedResourcesResponse) SetStatusCode(v int32) *ListProtectedResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProtectedResourcesResponse) SetBody(v *ListProtectedResourcesResponseBody) *ListProtectedResourcesResponse {
	s.Body = v
	return s
}

func (s *ListProtectedResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
