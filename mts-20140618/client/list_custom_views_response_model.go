// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomViewsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomViewsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomViewsResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomViewsResponseBody) *ListCustomViewsResponse
	GetBody() *ListCustomViewsResponseBody
}

type ListCustomViewsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomViewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomViewsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomViewsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomViewsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomViewsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomViewsResponse) GetBody() *ListCustomViewsResponseBody {
	return s.Body
}

func (s *ListCustomViewsResponse) SetHeaders(v map[string]*string) *ListCustomViewsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomViewsResponse) SetStatusCode(v int32) *ListCustomViewsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomViewsResponse) SetBody(v *ListCustomViewsResponseBody) *ListCustomViewsResponse {
	s.Body = v
	return s
}

func (s *ListCustomViewsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
