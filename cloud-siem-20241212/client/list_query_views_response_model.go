// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueryViewsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQueryViewsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQueryViewsResponse
	GetStatusCode() *int32
	SetBody(v *ListQueryViewsResponseBody) *ListQueryViewsResponse
	GetBody() *ListQueryViewsResponseBody
}

type ListQueryViewsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueryViewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQueryViewsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQueryViewsResponse) GoString() string {
	return s.String()
}

func (s *ListQueryViewsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQueryViewsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQueryViewsResponse) GetBody() *ListQueryViewsResponseBody {
	return s.Body
}

func (s *ListQueryViewsResponse) SetHeaders(v map[string]*string) *ListQueryViewsResponse {
	s.Headers = v
	return s
}

func (s *ListQueryViewsResponse) SetStatusCode(v int32) *ListQueryViewsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueryViewsResponse) SetBody(v *ListQueryViewsResponseBody) *ListQueryViewsResponse {
	s.Body = v
	return s
}

func (s *ListQueryViewsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
