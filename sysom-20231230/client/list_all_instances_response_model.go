// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListAllInstancesResponseBody) *ListAllInstancesResponse
	GetBody() *ListAllInstancesResponseBody
}

type ListAllInstancesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListAllInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllInstancesResponse) GetBody() *ListAllInstancesResponseBody {
	return s.Body
}

func (s *ListAllInstancesResponse) SetHeaders(v map[string]*string) *ListAllInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListAllInstancesResponse) SetStatusCode(v int32) *ListAllInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllInstancesResponse) SetBody(v *ListAllInstancesResponseBody) *ListAllInstancesResponse {
	s.Body = v
	return s
}

func (s *ListAllInstancesResponse) Validate() error {
	return dara.Validate(s)
}
