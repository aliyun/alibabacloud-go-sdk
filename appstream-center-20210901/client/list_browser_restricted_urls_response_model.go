// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrowserRestrictedURLsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBrowserRestrictedURLsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBrowserRestrictedURLsResponse
	GetStatusCode() *int32
	SetBody(v *ListBrowserRestrictedURLsResponseBody) *ListBrowserRestrictedURLsResponse
	GetBody() *ListBrowserRestrictedURLsResponseBody
}

type ListBrowserRestrictedURLsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBrowserRestrictedURLsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBrowserRestrictedURLsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserRestrictedURLsResponse) GoString() string {
	return s.String()
}

func (s *ListBrowserRestrictedURLsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBrowserRestrictedURLsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBrowserRestrictedURLsResponse) GetBody() *ListBrowserRestrictedURLsResponseBody {
	return s.Body
}

func (s *ListBrowserRestrictedURLsResponse) SetHeaders(v map[string]*string) *ListBrowserRestrictedURLsResponse {
	s.Headers = v
	return s
}

func (s *ListBrowserRestrictedURLsResponse) SetStatusCode(v int32) *ListBrowserRestrictedURLsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBrowserRestrictedURLsResponse) SetBody(v *ListBrowserRestrictedURLsResponseBody) *ListBrowserRestrictedURLsResponse {
	s.Body = v
	return s
}

func (s *ListBrowserRestrictedURLsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
