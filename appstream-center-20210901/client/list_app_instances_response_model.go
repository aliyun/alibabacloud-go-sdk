// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListAppInstancesResponseBody) *ListAppInstancesResponse
	GetBody() *ListAppInstancesResponseBody
}

type ListAppInstancesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListAppInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppInstancesResponse) GetBody() *ListAppInstancesResponseBody {
	return s.Body
}

func (s *ListAppInstancesResponse) SetHeaders(v map[string]*string) *ListAppInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListAppInstancesResponse) SetStatusCode(v int32) *ListAppInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppInstancesResponse) SetBody(v *ListAppInstancesResponseBody) *ListAppInstancesResponse {
	s.Body = v
	return s
}

func (s *ListAppInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
