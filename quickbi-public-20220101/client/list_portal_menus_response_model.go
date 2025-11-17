// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPortalMenusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPortalMenusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPortalMenusResponse
	GetStatusCode() *int32
	SetBody(v *ListPortalMenusResponseBody) *ListPortalMenusResponse
	GetBody() *ListPortalMenusResponseBody
}

type ListPortalMenusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPortalMenusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPortalMenusResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPortalMenusResponse) GoString() string {
	return s.String()
}

func (s *ListPortalMenusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPortalMenusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPortalMenusResponse) GetBody() *ListPortalMenusResponseBody {
	return s.Body
}

func (s *ListPortalMenusResponse) SetHeaders(v map[string]*string) *ListPortalMenusResponse {
	s.Headers = v
	return s
}

func (s *ListPortalMenusResponse) SetStatusCode(v int32) *ListPortalMenusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPortalMenusResponse) SetBody(v *ListPortalMenusResponseBody) *ListPortalMenusResponse {
	s.Body = v
	return s
}

func (s *ListPortalMenusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
