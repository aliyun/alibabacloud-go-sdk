// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHbaseInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHbaseInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHbaseInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListHbaseInstancesResponseBody) *ListHbaseInstancesResponse
	GetBody() *ListHbaseInstancesResponseBody
}

type ListHbaseInstancesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHbaseInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHbaseInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHbaseInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListHbaseInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHbaseInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHbaseInstancesResponse) GetBody() *ListHbaseInstancesResponseBody {
	return s.Body
}

func (s *ListHbaseInstancesResponse) SetHeaders(v map[string]*string) *ListHbaseInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListHbaseInstancesResponse) SetStatusCode(v int32) *ListHbaseInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHbaseInstancesResponse) SetBody(v *ListHbaseInstancesResponseBody) *ListHbaseInstancesResponse {
	s.Body = v
	return s
}

func (s *ListHbaseInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
