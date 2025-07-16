// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListInstancesResponseBody) *ListInstancesResponse
	GetBody() *ListInstancesResponseBody
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstancesResponse) GetBody() *ListInstancesResponseBody {
	return s.Body
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

func (s *ListInstancesResponse) Validate() error {
	return dara.Validate(s)
}
