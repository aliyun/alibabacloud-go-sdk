// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserOwnedResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserOwnedResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserOwnedResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserOwnedResourcesResponseBody) *ListUserOwnedResourcesResponse
	GetBody() *ListUserOwnedResourcesResponseBody
}

type ListUserOwnedResourcesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserOwnedResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserOwnedResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserOwnedResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListUserOwnedResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserOwnedResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserOwnedResourcesResponse) GetBody() *ListUserOwnedResourcesResponseBody {
	return s.Body
}

func (s *ListUserOwnedResourcesResponse) SetHeaders(v map[string]*string) *ListUserOwnedResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListUserOwnedResourcesResponse) SetStatusCode(v int32) *ListUserOwnedResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserOwnedResourcesResponse) SetBody(v *ListUserOwnedResourcesResponseBody) *ListUserOwnedResourcesResponse {
	s.Body = v
	return s
}

func (s *ListUserOwnedResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
