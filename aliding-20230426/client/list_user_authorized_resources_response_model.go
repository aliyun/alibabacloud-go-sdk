// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAuthorizedResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserAuthorizedResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserAuthorizedResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserAuthorizedResourcesResponseBody) *ListUserAuthorizedResourcesResponse
	GetBody() *ListUserAuthorizedResourcesResponseBody
}

type ListUserAuthorizedResourcesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserAuthorizedResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserAuthorizedResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserAuthorizedResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListUserAuthorizedResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserAuthorizedResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserAuthorizedResourcesResponse) GetBody() *ListUserAuthorizedResourcesResponseBody {
	return s.Body
}

func (s *ListUserAuthorizedResourcesResponse) SetHeaders(v map[string]*string) *ListUserAuthorizedResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListUserAuthorizedResourcesResponse) SetStatusCode(v int32) *ListUserAuthorizedResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserAuthorizedResourcesResponse) SetBody(v *ListUserAuthorizedResourcesResponseBody) *ListUserAuthorizedResourcesResponse {
	s.Body = v
	return s
}

func (s *ListUserAuthorizedResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
