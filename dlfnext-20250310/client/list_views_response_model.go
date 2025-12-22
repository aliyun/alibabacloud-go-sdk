// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListViewsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListViewsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListViewsResponse
	GetStatusCode() *int32
	SetBody(v *ListViewsResponseBody) *ListViewsResponse
	GetBody() *ListViewsResponseBody
}

type ListViewsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListViewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListViewsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListViewsResponse) GoString() string {
	return s.String()
}

func (s *ListViewsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListViewsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListViewsResponse) GetBody() *ListViewsResponseBody {
	return s.Body
}

func (s *ListViewsResponse) SetHeaders(v map[string]*string) *ListViewsResponse {
	s.Headers = v
	return s
}

func (s *ListViewsResponse) SetStatusCode(v int32) *ListViewsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListViewsResponse) SetBody(v *ListViewsResponseBody) *ListViewsResponse {
	s.Body = v
	return s
}

func (s *ListViewsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
