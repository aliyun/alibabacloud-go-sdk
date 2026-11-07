// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrowserBookmarksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBrowserBookmarksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBrowserBookmarksResponse
	GetStatusCode() *int32
	SetBody(v *ListBrowserBookmarksResponseBody) *ListBrowserBookmarksResponse
	GetBody() *ListBrowserBookmarksResponseBody
}

type ListBrowserBookmarksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBrowserBookmarksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBrowserBookmarksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserBookmarksResponse) GoString() string {
	return s.String()
}

func (s *ListBrowserBookmarksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBrowserBookmarksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBrowserBookmarksResponse) GetBody() *ListBrowserBookmarksResponseBody {
	return s.Body
}

func (s *ListBrowserBookmarksResponse) SetHeaders(v map[string]*string) *ListBrowserBookmarksResponse {
	s.Headers = v
	return s
}

func (s *ListBrowserBookmarksResponse) SetStatusCode(v int32) *ListBrowserBookmarksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBrowserBookmarksResponse) SetBody(v *ListBrowserBookmarksResponseBody) *ListBrowserBookmarksResponse {
	s.Body = v
	return s
}

func (s *ListBrowserBookmarksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
