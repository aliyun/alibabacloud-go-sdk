// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHBaseInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHBaseInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHBaseInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListHBaseInstancesResponseBody) *ListHBaseInstancesResponse
	GetBody() *ListHBaseInstancesResponseBody
}

type ListHBaseInstancesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHBaseInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHBaseInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHBaseInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListHBaseInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHBaseInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHBaseInstancesResponse) GetBody() *ListHBaseInstancesResponseBody {
	return s.Body
}

func (s *ListHBaseInstancesResponse) SetHeaders(v map[string]*string) *ListHBaseInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListHBaseInstancesResponse) SetStatusCode(v int32) *ListHBaseInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHBaseInstancesResponse) SetBody(v *ListHBaseInstancesResponseBody) *ListHBaseInstancesResponse {
	s.Body = v
	return s
}

func (s *ListHBaseInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
