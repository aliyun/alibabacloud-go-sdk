// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnsInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAnsInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAnsInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListAnsInstancesResponseBody) *ListAnsInstancesResponse
	GetBody() *ListAnsInstancesResponseBody
}

type ListAnsInstancesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAnsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAnsInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAnsInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListAnsInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAnsInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAnsInstancesResponse) GetBody() *ListAnsInstancesResponseBody {
	return s.Body
}

func (s *ListAnsInstancesResponse) SetHeaders(v map[string]*string) *ListAnsInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListAnsInstancesResponse) SetStatusCode(v int32) *ListAnsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAnsInstancesResponse) SetBody(v *ListAnsInstancesResponseBody) *ListAnsInstancesResponse {
	s.Body = v
	return s
}

func (s *ListAnsInstancesResponse) Validate() error {
	return dara.Validate(s)
}
