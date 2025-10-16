// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrowsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBrowsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBrowsersResponse
	GetStatusCode() *int32
	SetBody(v *ListBrowsersResult) *ListBrowsersResponse
	GetBody() *ListBrowsersResult
}

type ListBrowsersResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBrowsersResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBrowsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBrowsersResponse) GoString() string {
	return s.String()
}

func (s *ListBrowsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBrowsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBrowsersResponse) GetBody() *ListBrowsersResult {
	return s.Body
}

func (s *ListBrowsersResponse) SetHeaders(v map[string]*string) *ListBrowsersResponse {
	s.Headers = v
	return s
}

func (s *ListBrowsersResponse) SetStatusCode(v int32) *ListBrowsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBrowsersResponse) SetBody(v *ListBrowsersResult) *ListBrowsersResponse {
	s.Body = v
	return s
}

func (s *ListBrowsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
